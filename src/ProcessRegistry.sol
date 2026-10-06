// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {IProcessRegistry} from "./interfaces/IProcessRegistry.sol";
import {IZiskVerifier} from "./interfaces/IZiskVerifier.sol";
import {ICensusValidator} from "./interfaces/ICensusValidator.sol";
import {DAVINCITypes} from "./libraries/DAVINCITypes.sol";
import {ProcessIdLib} from "./libraries/ProcessIdLib.sol";
import {BlobsLib} from "./libraries/BlobsLib.sol";
import {GenesisLib} from "./libraries/GenesisLib.sol";
import {PublicsLib} from "./libraries/PublicsLib.sol";
import {Sha256SmtLib} from "./libraries/Sha256SmtLib.sol";
import {BjjFormLib} from "./libraries/BjjFormLib.sol";
import {DavinciDKGAdapter} from "./DavinciDKGAdapter.sol";
import {CouncilAdapter} from "./CouncilAdapter.sol";
import {IDkgResultsAdapter} from "./interfaces/IDkgResultsAdapter.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title ProcessRegistry
 * @notice Stores DAVINCI processes, manages their lifecycle and settles davinci-zkvm proofs:
 *         one ZisK PLONK per state transition, bound to its EIP-4844 blobs, and one for the
 *         results.
 * @dev Root encodings: state roots are the raw SHA-256 digest of the arbo root (reg32 of the
 *      publics); census roots are big-endian integers (the byte reverse of reg32).
 */
contract ProcessRegistry is IProcessRegistry, ReentrancyGuard {
    using ProcessIdLib for bytes31;

    /// @dev Upper bound for the maximum possible decrypted result.
    uint256 private constant MAX_POSSIBLE_RESULT_CAP = 1_000_000_000_000;
    /// @dev Ballot capacity of the zkVM guest (NUM_FIELDS).
    uint8 private constant MAX_NUM_FIELDS = 16;
    /// @dev EIP-4844 point-evaluation precompile and the output it returns on success.
    address private constant POINT_EVALUATION = address(0x0A);
    uint256 private constant FIELD_ELEMENTS_PER_BLOB = 4096;
    uint256 private constant BLS_MODULUS =
        52435875175126190479447740508185965837690552500527637822603658699938581184513;
    /// @dev Gas for each call into an origin-3 census contract (OnchainCensus needs < 5k).
    uint256 private constant CENSUS_CALL_GAS = 100_000;

    // Batch guest output registers (circuit/CIRCUIT.md §3).
    uint256 private constant REG_OK = 0;
    uint256 private constant REG_FAIL_MASK = 1;
    uint256 private constant REG_ROOT_BEFORE = 2;
    uint256 private constant REG_ROOT_AFTER = 10;
    uint256 private constant REG_VOTERS = 18;
    uint256 private constant REG_OVERWRITES = 19;
    uint256 private constant REG_CENSUS_ROOT = 20;
    uint256 private constant REG_BLOBS_DIGEST = 28;
    uint256 private constant REG_N_BLOBS = 36;
    uint256 private constant REG_OCCUPIED_BEFORE = 42;
    // Results guest output registers: the state root, then results[i] as (lo, hi) words.
    uint256 private constant REG_RESULTS_STATE_ROOT = 2;
    uint256 private constant REG_RESULTS = 10;

    /**
     * @notice The maximum value of the process status.
     */
    uint8 public constant MAX_STATUS = 4;
    /**
     * @notice The process mapping is a mapping of process IDs to processes.
     */
    mapping(bytes31 => DAVINCITypes.Process) public processes;
    /**
     * @notice The process nonce mapping is a mapping of addresses to process nonces.
     */
    mapping(address => uint64) public processNonce;
    /**
     * @notice The process count is the number of processes created.
     */
    uint32 public processCount;
    /**
     * @notice The chain ID is the ID of the chain.
     */
    uint32 public chainID;
    /**
     * @notice The pidPrefix is the 4-byte prefix used in process IDs.
     * This is computed as the last 4 bytes of keccak256(abi.encodePacked(chainID, address(this))).
     */
    uint32 public pidPrefix;
    /**
     * @notice The ZisK PLONK verifier.
     */
    IZiskVerifier public immutable ziskVerifier;
    /**
     * @notice Program vk of the vote-batch guest.
     */
    bytes32 public immutable batchProgramVK;
    /**
     * @notice Program vk of the results guest.
     */
    bytes32 public immutable resultsProgramVK;
    /**
     * @notice Root of the ZisK vadcop-final setup both proofs are wrapped under.
     */
    bytes32 public immutable rootCVadcopFinal;
    /**
     * @notice sha256 of the ballot Groth16 VK wire bytes (state leaf 0x07), as the digest.
     */
    bytes32 public immutable ballotVKHash;
    /**
     * @notice The DavinciDKGAdapter this registry created at deploy, or address(0) when
     *         the DKG key modes are disabled.
     */
    address public immutable dkgAdapter;
    /**
     * @notice The CouncilAdapter this registry created at deploy, or address(0) when the
     *         COUNCIL key mode is disabled.
     */
    address public immutable councilAdapter;
    /**
     * @notice The grace window every new process starts with, in seconds.
     */
    uint32 public immutable defaultGrace;
    /**
     * @notice The shortest grace window setProcessGrace accepts, in seconds.
     */
    uint32 public immutable graceFloor;
    /**
     * @notice The longest grace window setProcessGrace accepts, in seconds.
     */
    uint32 public immutable graceCeil;
    /**
     * @notice Hard cap on the grace window past the end time, in seconds, however many
     *         transitions keep extending it.
     */
    uint32 public immutable graceMaxTotal;
    /**
     * @notice The shortest notice setProcessDuration gives when it shortens a process: the new
     *         end is at least this many seconds after the call.
     */
    uint32 public immutable noticeMin;

    /**
     * @notice Initializes the contract.
     * @param _chainID The ID of the chain.
     * @param _ziskVerifier The ZisK PLONK verifier.
     * @param _batchProgramVK Program vk of the vote-batch guest.
     * @param _resultsProgramVK Program vk of the results guest.
     * @param _rootCVadcopFinal Root of the ZisK vadcop-final setup.
     * @param _ballotVKHash sha256 digest of the ballot proof VK (davinci.BallotVKLeaf).
     * @param _dkgManager The davinci-dkg DKGManager, or address(0) to disable DKG modes.
     *        When set, the constructor creates the DavinciDKGAdapter.
     * @param _councilManager The Council manager, or address(0) to disable the COUNCIL
     *        mode. When set, the constructor creates the CouncilAdapter.
     * @param _defaultGrace The grace window of a new process, in seconds.
     * @param _graceFloor The minimum for setProcessGrace, non-zero and at most _defaultGrace.
     * @param _graceCeil The maximum for setProcessGrace, at least _defaultGrace.
     * @param _graceMaxTotal The cap on the window past the end time, at least _graceCeil.
     * @param _noticeMin The minimum notice, in seconds, for shortening a process, non-zero.
     */
    constructor(
        uint32 _chainID,
        address _ziskVerifier,
        bytes32 _batchProgramVK,
        bytes32 _resultsProgramVK,
        bytes32 _rootCVadcopFinal,
        bytes32 _ballotVKHash,
        address _dkgManager,
        address _councilManager,
        uint32 _defaultGrace,
        uint32 _graceFloor,
        uint32 _graceCeil,
        uint32 _graceMaxTotal,
        uint32 _noticeMin
    ) {
        if (
            _ziskVerifier == address(0) || _batchProgramVK == bytes32(0) || _resultsProgramVK == bytes32(0)
                || _rootCVadcopFinal == bytes32(0) || _ballotVKHash == bytes32(0)
        ) revert InvalidVerifierConfig();
        if (
            _graceFloor == 0 || _graceFloor > _defaultGrace || _defaultGrace > _graceCeil || _graceCeil > _graceMaxTotal
                || _noticeMin == 0
        ) revert InvalidGrace();
        defaultGrace = _defaultGrace;
        graceFloor = _graceFloor;
        graceCeil = _graceCeil;
        graceMaxTotal = _graceMaxTotal;
        noticeMin = _noticeMin;
        ziskVerifier = IZiskVerifier(_ziskVerifier);
        batchProgramVK = _batchProgramVK;
        resultsProgramVK = _resultsProgramVK;
        rootCVadcopFinal = _rootCVadcopFinal;
        ballotVKHash = _ballotVKHash;
        chainID = _chainID;
        pidPrefix = ProcessIdLib.getPrefix(_chainID, address(this));
        dkgAdapter = _dkgManager == address(0) ? address(0) : address(new DavinciDKGAdapter(_dkgManager));
        // Created second, so the DKG adapter keeps its address (CREATE(registry, 1)).
        councilAdapter = _councilManager == address(0) ? address(0) : address(new CouncilAdapter(_councilManager));
    }

    /// @inheritdoc IProcessRegistry
    function getProcess(bytes31 processId) external view override returns (DAVINCITypes.Process memory) {
        return processes[processId];
    }

    /// @inheritdoc IProcessRegistry
    function getProcessEndTime(bytes31 processId) external view returns (uint256) {
        DAVINCITypes.Process memory p = processes[processId];
        return p.startTime + p.duration;
    }

    /// @inheritdoc IProcessRegistry
    function getProcessGraceEnd(bytes31 processId) external view override returns (uint256) {
        return _graceEnd(processes[processId]);
    }

    /// @inheritdoc IProcessRegistry
    function getSTVerifierVKeyHash() external view override returns (bytes32) {
        return batchProgramVK;
    }

    /// @inheritdoc IProcessRegistry
    function getRVerifierVKeyHash() external view override returns (bytes32) {
        return resultsProgramVK;
    }

    /// @inheritdoc IProcessRegistry
    function getNextProcessId(address organizationId) external view override returns (bytes31) {
        return ProcessIdLib.computeProcessId(pidPrefix, organizationId, processNonce[organizationId]);
    }

    /// @inheritdoc IProcessRegistry
    function genesisRoot(
        bytes31 processId,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.EncryptionKey calldata encryptionKey,
        DAVINCITypes.CensusOrigin censusOrigin
    ) external view override returns (bytes32) {
        return GenesisLib.root(processId, ballotMode, encryptionKey, censusOrigin, ballotVKHash);
    }

    /// @inheritdoc IProcessRegistry
    function newProcess(
        DAVINCITypes.ProcessStatus status,
        uint256 startTime,
        uint256 duration,
        uint256 maxVoters,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.Census calldata census,
        string calldata metadataURI,
        bytes32 metadataHash,
        DAVINCITypes.EncryptionKey calldata encryptionKey,
        DAVINCITypes.DKGParams calldata dkg
    ) external override nonReentrant returns (bytes31) {
        address sender = msg.sender;
        bytes31 processId = ProcessIdLib.computeProcessId(pidPrefix, sender, processNonce[sender]);

        // Validate process doesn't exist and validate inputs
        bytes32 censusRoot = _validateNewProcess(processId, sender, status, maxVoters, ballotMode, census);
        _checkMetadata(metadataURI, metadataHash);

        // validate start time, block and duration
        uint256 currentTimestamp = block.timestamp;
        if (startTime == 0) {
            startTime = currentTimestamp;
        }
        if (startTime < currentTimestamp) revert InvalidStartTime();
        if (startTime + duration <= currentTimestamp) revert InvalidDuration();

        DAVINCITypes.Process storage p = processes[processId];

        // Resolve the encryption key: the caller's in SEQUENCER mode, the DKG committee's
        // (converted to circomlib form) in the DKG modes, the Council ceremony's in COUNCIL.
        DAVINCITypes.EncryptionKey memory key = encryptionKey;
        if (dkg.mode == DAVINCITypes.KeyMode.SEQUENCER) {
            if (
                dkg.epochId != bytes12(0) || dkg.orgPKx != 0 || dkg.orgPKy != 0 || dkg.popAx != 0 || dkg.popAy != 0
                    || dkg.popZ != 0
            ) revert InvalidDKGParams();
        } else {
            if (encryptionKey.x != 0 || encryptionKey.y != 0) revert InvalidEncryptionKey();
            bytes12 eid;
            bytes32 aid;
            uint256 teX;
            uint256 teY;
            if (dkg.mode == DAVINCITypes.KeyMode.COUNCIL) {
                if (councilAdapter == address(0)) revert CouncilDisabled();
                // The ceremony authorizes the creator, so the adapter must learn it.
                (eid, aid, teX, teY) = CouncilAdapter(councilAdapter).register(processId, sender, dkg);
            } else {
                if (dkgAdapter == address(0)) revert DKGDisabled();
                (eid, aid, teX, teY) = DavinciDKGAdapter(dkgAdapter).register(processId, dkg);
            }
            key = DAVINCITypes.EncryptionKey(teX, teY);
            p.keyMode = dkg.mode;
            p.dkgEpochId = eid;
            p.dkgAid = aid;
        }
        if (!GenesisLib.isValidEncryptionKey(key)) revert InvalidEncryptionKey();

        p.status = status;
        p.startTime = startTime;
        p.duration = duration;
        p.maxVoters = maxVoters;
        p.organizationId = sender;
        p.encryptionKey = key;
        p.latestStateRoot = GenesisLib.root(processId, ballotMode, key, census.censusOrigin, ballotVKHash);
        p.ballotMode = ballotMode;
        p.census = census;
        p.census.censusRoot = censusRoot;
        p.creationBlock = block.number;
        p.grace = defaultGrace;

        processCount++;
        processNonce[sender]++;

        emit ProcessCreated(processId, sender);
        _storeMetadata(processId, p, metadataURI, metadataHash);
        return processId;
    }

    /// @inheritdoc IProcessRegistry
    /// @dev ENDED before the end time moves the end to now. Past it the end stays put: moving
    ///      it forward would reopen a grace window that has already closed. Two transitions the
    ///      matrix allows are refused by time: ENDED before startTime (CANCELED is how a process
    ///      that never opened is voided) and PAUSED from the end on (a pause cannot hold the
    ///      grace window shut).
    function setProcessStatus(bytes31 processId, DAVINCITypes.ProcessStatus newStatus) external override {
        if (processId == bytes31(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        if (uint8(newStatus) > MAX_STATUS) revert InvalidStatus();

        DAVINCITypes.Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.organizationId != msg.sender) revert Unauthorized();

        // validate status transition
        DAVINCITypes.ProcessStatus oldStatus = p.status;
        if (!_validateStatusTransition(oldStatus, newStatus)) revert InvalidStatus();
        uint256 startTime = p.startTime;
        uint256 end = startTime + p.duration;
        if (newStatus == DAVINCITypes.ProcessStatus.ENDED && block.timestamp < startTime) revert InvalidTimeBounds();
        if (newStatus == DAVINCITypes.ProcessStatus.PAUSED && block.timestamp >= end) revert InvalidTimeBounds();

        p.status = newStatus;
        // if newStatus is ENDED before the end, set duration to the time elapsed since start
        if (newStatus == DAVINCITypes.ProcessStatus.ENDED && block.timestamp < end) {
            uint256 newDuration = block.timestamp - startTime;
            p.duration = newDuration;
            emit ProcessDurationChanged(processId, newDuration);
        }

        emit ProcessStatusChanged(processId, oldStatus, newStatus);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessCensus(bytes31 processId, DAVINCITypes.Census calldata census) external override {
        DAVINCITypes.Process storage p = _existingProcess(processId);
        if (p.organizationId != msg.sender) revert Unauthorized();
        if (p.census.censusOrigin != DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1) {
            revert CensusNotUpdatable();
        }
        if (census.censusOrigin != DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1) {
            revert InvalidCensusOrigin();
        }
        if (census.contractAddress != address(0)) revert InvalidCensusAddress();
        if (census.censusRoot == bytes32(0)) revert InvalidCensusRoot();
        if (bytes(census.censusURI).length == 0) revert InvalidCensusURI();

        DAVINCITypes.ProcessStatus status = p.status;
        if (status != DAVINCITypes.ProcessStatus.READY && status != DAVINCITypes.ProcessStatus.PAUSED) {
            revert InvalidStatus();
        }
        if (p.startTime + p.duration <= block.timestamp) revert InvalidTimeBounds();

        p.census.censusRoot = census.censusRoot;
        p.census.censusURI = census.censusURI;

        emit CensusUpdated(processId, census.censusRoot, census.censusURI);
    }

    /// @inheritdoc IProcessRegistry
    /// @dev Same window as setProcessCensus: once the end passes, the meaning of every ballot
    ///      field is fixed.
    function setProcessMetadata(bytes31 processId, string calldata metadataURI, bytes32 metadataHash)
        external
        override
    {
        DAVINCITypes.Process storage p = _existingProcess(processId);
        if (p.organizationId != msg.sender) revert Unauthorized();
        _checkMetadata(metadataURI, metadataHash);

        DAVINCITypes.ProcessStatus status = p.status;
        if (status != DAVINCITypes.ProcessStatus.READY && status != DAVINCITypes.ProcessStatus.PAUSED) {
            revert InvalidStatus();
        }
        if (p.startTime + p.duration <= block.timestamp) revert InvalidTimeBounds();

        _storeMetadata(processId, p, metadataURI, metadataHash);
    }

    /// @inheritdoc IProcessRegistry
    /// @dev Note that the end time of the process is startTime + duration. A shorter duration
    ///      needs noticeMin: every node sees the new end before it bites, and no vote admitted
    ///      under the old one turns late.
    function setProcessDuration(bytes31 processId, uint256 _duration) external override {
        if (processId == bytes31(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        DAVINCITypes.Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.organizationId != msg.sender) revert Unauthorized();

        // check ongoing process
        DAVINCITypes.ProcessStatus status = p.status;
        if (status != DAVINCITypes.ProcessStatus.READY && status != DAVINCITypes.ProcessStatus.PAUSED) {
            revert InvalidStatus();
        }

        // check valid duration
        uint256 startTime = p.startTime;
        uint256 oldEnd = startTime + p.duration;
        // Past the end the tally may already be public (results tx in the mempool): no reopening.
        if (oldEnd <= block.timestamp) revert InvalidTimeBounds();
        uint256 newEnd = startTime + _duration;
        if (
            _duration == 0 || newEnd <= block.timestamp || newEnd == oldEnd
                || (newEnd < oldEnd && newEnd < block.timestamp + noticeMin)
        ) revert InvalidDuration();

        p.duration = _duration;

        emit ProcessDurationChanged(processId, _duration);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessMaxVoters(bytes31 processId, uint256 _maxVoters) external override {
        if (processId == bytes31(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        DAVINCITypes.Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.organizationId != msg.sender) revert Unauthorized();

        // check ongoing process
        DAVINCITypes.ProcessStatus status = p.status;
        if (status != DAVINCITypes.ProcessStatus.READY && status != DAVINCITypes.ProcessStatus.PAUSED) {
            revert InvalidStatus();
        }
        // Past the end the cap would pick which queued batches still land in the grace window.
        if (p.startTime + p.duration <= block.timestamp) revert InvalidTimeBounds();

        // check valid maxVoters
        if (_maxVoters == 0 || _maxVoters < p.votersCount) revert InvalidMaxVoters();
        _validateMaxPossibleResultCap(_maxVoters, p.ballotMode.maxValue);

        p.maxVoters = _maxVoters;

        emit ProcessMaxVotersChanged(processId, _maxVoters);
    }

    /// @inheritdoc IProcessRegistry
    /// @dev Same window as setProcessDuration: past the end the grace is already running.
    function setProcessGrace(bytes31 processId, uint32 grace) external override {
        DAVINCITypes.Process storage p = _existingProcess(processId);
        if (p.organizationId != msg.sender) revert Unauthorized();

        DAVINCITypes.ProcessStatus status = p.status;
        if (status != DAVINCITypes.ProcessStatus.READY && status != DAVINCITypes.ProcessStatus.PAUSED) {
            revert InvalidStatus();
        }
        if (p.startTime + p.duration <= block.timestamp) revert InvalidTimeBounds();
        if (grace < graceFloor || grace > graceCeil) revert InvalidGrace();

        p.grace = grace;

        emit ProcessGraceChanged(processId, grace);
    }

    /// @inheritdoc IProcessRegistry
    /// @dev Permissionless: the proof, the blob openings and root continuity authenticate it.
    ///      Settles through the grace window past the end, so batches still queued or proving
    ///      at the end land. PAUSED blocks settlement only while voting is open: a process
    ///      paused at the end settles through the window like READY or ENDED.
    function submitStateTransition(
        bytes31 processId,
        bytes calldata publicValues,
        bytes calldata proofBytes,
        bytes[] calldata commitments,
        bytes32[] calldata ys,
        bytes[] calldata kzgProofs
    ) external override nonReentrant {
        DAVINCITypes.Process storage p = _existingProcess(processId);
        DAVINCITypes.ProcessStatus status = p.status;
        if (
            status != DAVINCITypes.ProcessStatus.READY && status != DAVINCITypes.ProcessStatus.ENDED
                && !(status == DAVINCITypes.ProcessStatus.PAUSED && block.timestamp >= p.startTime + p.duration)
        ) revert InvalidStatus();
        if (block.timestamp < p.startTime) revert InvalidTimeBounds();
        if (block.timestamp >= _graceEnd(p)) revert InvalidTimeBounds();

        _checkGuestOk(publicValues);
        bytes32 rootBefore = PublicsLib.reg32(publicValues, REG_ROOT_BEFORE);
        if (rootBefore != p.latestStateRoot) revert InvalidStateRoot();
        _checkCensusRoot(p, PublicsLib.reverse32(PublicsLib.reg32(publicValues, REG_CENSUS_ROOT)));
        // occupied_before counts the distinct slots written so far, which is votersCount.
        uint256 votersCount = p.votersCount;
        if (PublicsLib.word(publicValues, REG_OCCUPIED_BEFORE) != votersCount) revert InvalidOccupiedBefore();

        uint256 overwrites = PublicsLib.word(publicValues, REG_OVERWRITES);
        uint256 newVoters = PublicsLib.word(publicValues, REG_VOTERS) - overwrites;
        // A refresh-only batch would extend the grace window without carrying a vote.
        if (newVoters + overwrites == 0) revert EmptyTransition();
        if (votersCount + newVoters > p.maxVoters) revert MaxVotersReached();

        uint256 nBlobs = _checkBlobsDigest(publicValues, commitments, ys, kzgProofs);

        ziskVerifier.verifySnarkProof(batchProgramVK, rootCVadcopFinal, publicValues, proofBytes);

        _verifyOpenings(processId, rootBefore, commitments, ys, kzgProofs);

        bytes32 rootAfter = PublicsLib.reg32(publicValues, REG_ROOT_AFTER);
        p.latestStateRoot = rootAfter;
        p.votersCount = votersCount + newVoters;
        p.overwrittenVotesCount += overwrites;
        ++p.batchNumber;
        p.lastVoteAt = uint64(block.timestamp);

        emit ProcessStateTransitioned(
            processId, msg.sender, rootBefore, rootAfter, p.votersCount, p.overwrittenVotesCount, nBlobs
        );
    }

    /// @inheritdoc IProcessRegistry
    /// @dev Permissionless: the results guest proves the tally against the final state root.
    function setProcessResults(bytes31 processId, bytes calldata publicValues, bytes calldata proofBytes)
        external
        override
        nonReentrant
    {
        DAVINCITypes.Process storage p = _existingProcess(processId);
        // DKG-mode results come from the committee, not the results guest.
        if (p.keyMode != DAVINCITypes.KeyMode.SEQUENCER) revert InvalidKeyMode();

        // Cannot set results on CANCELLED or RESULTS processes
        DAVINCITypes.ProcessStatus oldStatus = p.status;
        if (oldStatus == DAVINCITypes.ProcessStatus.CANCELED || oldStatus == DAVINCITypes.ProcessStatus.RESULTS) {
            revert InvalidStatus();
        }

        // Require that the process has ended, either by status or by time
        if (oldStatus != DAVINCITypes.ProcessStatus.ENDED && p.startTime + p.duration > block.timestamp) {
            revert InvalidTimeBounds();
        }
        // and that the grace window has closed, so latestStateRoot is final.
        if (block.timestamp < _graceEnd(p)) revert GraceOpen();

        _checkGuestOk(publicValues);
        if (PublicsLib.reg32(publicValues, REG_RESULTS_STATE_ROOT) != p.latestStateRoot) revert InvalidStateRoot();

        ziskVerifier.verifySnarkProof(resultsProgramVK, rootCVadcopFinal, publicValues, proofBytes);

        uint256 numFields = p.ballotMode.numFields;
        uint256[] memory result = new uint256[](numFields);
        for (uint256 i = 0; i < numFields; ++i) {
            uint256 lo = PublicsLib.word(publicValues, REG_RESULTS + 2 * i);
            uint256 hi = PublicsLib.word(publicValues, REG_RESULTS + 2 * i + 1);
            result[i] = lo | (hi << 32);
        }

        p.status = DAVINCITypes.ProcessStatus.RESULTS;
        p.result = result;

        emit ProcessStatusChanged(processId, oldStatus, DAVINCITypes.ProcessStatus.RESULTS);
        emit ProcessResultsSet(processId, msg.sender, result);
    }

    /// @inheritdoc IProcessRegistry
    function aidFor(bytes31 processId) external view override returns (bytes32) {
        if (dkgAdapter == address(0)) revert DKGDisabled();
        return DavinciDKGAdapter(dkgAdapter).aidFor(processId);
    }

    /// @inheritdoc IProcessRegistry
    /// @dev Permissionless: the SMT inclusion proof binds the accumulator to the settled
    ///      state root, so a wrong accumulator cannot be submitted for decryption.
    function requestResultsDecryption(bytes31 processId, uint256[64] calldata accumulator, bytes32[] calldata siblings)
        external
        override
        nonReentrant
    {
        DAVINCITypes.Process storage p = _existingProcess(processId);
        if (p.keyMode == DAVINCITypes.KeyMode.SEQUENCER) revert InvalidKeyMode();
        if (p.dkgResultsRequested) revert ResultsAlreadyRequested();

        // Same end rule as setProcessResults: ENDED, or READY/PAUSED past the end, and the
        // grace window closed.
        DAVINCITypes.ProcessStatus oldStatus = p.status;
        if (oldStatus == DAVINCITypes.ProcessStatus.CANCELED || oldStatus == DAVINCITypes.ProcessStatus.RESULTS) {
            revert InvalidStatus();
        }
        if (oldStatus != DAVINCITypes.ProcessStatus.ENDED && p.startTime + p.duration > block.timestamp) {
            revert InvalidTimeBounds();
        }
        if (block.timestamp < _graceEnd(p)) revert GraceOpen();

        // The leaf hash binds raw bytes; range-checking every coordinate leaves (0, 1)
        // as the unique identity encoding.
        for (uint256 i = 0; i < 64; ++i) {
            if (accumulator[i] >= BjjFormLib.Q) revert InvalidAccumulator();
        }
        // Leaf 0x04 value: sha256 of the 64 coordinates as BE32 words.
        uint256 leaf = uint256(sha256(abi.encode(accumulator)));
        if (!Sha256SmtLib.verifyInclusion(p.latestStateRoot, GenesisLib.KEY_RESULTS, leaf, siblings)) {
            revert InvalidInclusionProof();
        }

        p.dkgResultsRequested = true;

        // The combined plaintexts become public on the DKG before finalize runs, so the
        // process must leave organizer control here: move it to ENDED (terminal except
        // for the internal step to RESULTS), or a canceling organizer could read the
        // tally first and veto it. duration is left alone — the end has already passed
        // or the organizer ended it.
        if (oldStatus != DAVINCITypes.ProcessStatus.ENDED) {
            p.status = DAVINCITypes.ProcessStatus.ENDED;
            emit ProcessStatusChanged(processId, oldStatus, DAVINCITypes.ProcessStatus.ENDED);
        }

        // Collect the active (non-identity) ciphertexts of the declared fields.
        uint256 numFields = p.ballotMode.numFields;
        uint256[4][] memory cts = new uint256[4][](numFields);
        uint16 zeroSkipped;
        uint256 n;
        for (uint256 i = 0; i < numFields; ++i) {
            uint256 c1x = accumulator[4 * i];
            uint256 c1y = accumulator[4 * i + 1];
            bool c2Identity = accumulator[4 * i + 2] == 0 && accumulator[4 * i + 3] == 1;
            // A field is either fully identity or fully active: a half-identity
            // ciphertext cannot happen honestly and the DKG would reject it anyway.
            if (c1x == 0 && c1y == 1) {
                if (!c2Identity) revert InvalidAccumulator();
                zeroSkipped |= uint16(1 << i);
                continue;
            }
            if (c2Identity) revert InvalidAccumulator();
            cts[n] = [c1x, c1y, accumulator[4 * i + 2], accumulator[4 * i + 3]];
            ++n;
        }

        uint16 firstIndex;
        if (n > 0) {
            assembly ("memory-safe") {
                mstore(cts, n) // shrink to the active count
            }
            firstIndex = _adapterFor(p.keyMode).submit(p.dkgEpochId, p.dkgAid, cts);
        }
        p.dkgFirstIndex = firstIndex;
        p.dkgCount = uint8(n);
        p.dkgZeroSkipped = zeroSkipped;

        emit ResultsDecryptionRequested(processId, p.dkgEpochId, p.dkgAid, firstIndex, uint8(n));

        // Nothing to decrypt: every field is zero, finalize right away.
        if (n == 0) _finalizeDKGResults(processId, p);
    }

    /// @inheritdoc IProcessRegistry
    function finalizeResultsFromDKG(bytes31 processId) external override nonReentrant {
        DAVINCITypes.Process storage p = _existingProcess(processId);
        if (p.keyMode == DAVINCITypes.KeyMode.SEQUENCER) revert InvalidKeyMode();
        DAVINCITypes.ProcessStatus status = p.status;
        if (status == DAVINCITypes.ProcessStatus.CANCELED || status == DAVINCITypes.ProcessStatus.RESULTS) {
            revert InvalidStatus();
        }
        if (block.timestamp < _graceEnd(p)) revert GraceOpen();
        if (!p.dkgResultsRequested) revert ResultsNotReady();
        _finalizeDKGResults(processId, p);
    }

    /// @inheritdoc IProcessRegistry
    /// @dev COUNCIL processes have no organizer key and revert InvalidKeyMode here too.
    function revealProcessKey(bytes31 processId, uint256 sk) external override nonReentrant {
        DAVINCITypes.Process storage p = _existingProcess(processId);
        if (p.keyMode != DAVINCITypes.KeyMode.DKG_LOCKED) revert InvalidKeyMode();
        DavinciDKGAdapter(dkgAdapter).reveal(p.dkgEpochId, p.dkgAid, sk);
    }

    /// @dev Reads the combined plaintexts, fills result[] (numFields entries, the same
    ///      shape setProcessResults stores: DKG plaintexts in field order, 0 for identity
    ///      fields) and flips the process to RESULTS.
    function _finalizeDKGResults(bytes31 processId, DAVINCITypes.Process storage p) private {
        uint256 numFields = p.ballotMode.numFields;
        uint256[] memory result = new uint256[](numFields);
        uint256 count = p.dkgCount;
        if (count > 0) {
            (bool ready, uint256[] memory values) =
                _adapterFor(p.keyMode).plaintexts(p.dkgEpochId, p.dkgAid, p.dkgFirstIndex, uint16(count));
            if (!ready) revert ResultsNotReady();
            uint256 zeroSkipped = p.dkgZeroSkipped;
            uint256 j;
            for (uint256 i = 0; i < numFields; ++i) {
                if ((zeroSkipped >> i) & 1 == 0) {
                    result[i] = values[j];
                    ++j;
                }
            }
        }

        DAVINCITypes.ProcessStatus oldStatus = p.status;
        p.status = DAVINCITypes.ProcessStatus.RESULTS;
        p.result = result;

        emit ProcessStatusChanged(processId, oldStatus, DAVINCITypes.ProcessStatus.RESULTS);
        emit ProcessResultsSet(processId, msg.sender, result);
    }

    /// @dev The results adapter of a non-SEQUENCER process: the CouncilAdapter for
    ///      COUNCIL, the DavinciDKGAdapter for the DKG modes.
    function _adapterFor(DAVINCITypes.KeyMode mode) private view returns (IDkgResultsAdapter) {
        return IDkgResultsAdapter(mode == DAVINCITypes.KeyMode.COUNCIL ? councilAdapter : dkgAdapter);
    }

    /**
     * @dev Validates inputs for a new process and returns the census root to store.
     */
    function _validateNewProcess(
        bytes31 processId,
        address sender,
        DAVINCITypes.ProcessStatus status,
        uint256 maxVoters,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.Census calldata census
    ) private view returns (bytes32 censusRoot) {
        if (processes[processId].organizationId == sender) {
            revert ProcessAlreadyExists();
        }

        // validate ballot mode
        if (ballotMode.numFields == 0 || ballotMode.numFields > MAX_NUM_FIELDS) revert InvalidMaxCount();
        if (ballotMode.groupSize > ballotMode.numFields) revert InvalidGroupSize();
        if (ballotMode.minValue > ballotMode.maxValue) revert InvalidMaxMinValueBounds();
        if (ballotMode.minValueSum > ballotMode.maxValueSum) revert InvalidValueSumBounds();

        // validate max voters
        if (maxVoters == 0) revert InvalidMaxVoters();
        _validateMaxPossibleResultCap(maxVoters, ballotMode.maxValue);

        // validate census:
        //  - MERKLE_TREE_OFFCHAIN_STATIC_V1 -> lean-IMT root (fixed)
        //  - MERKLE_TREE_OFFCHAIN_DYNAMIC_V1 -> lean-IMT root, replaced by setProcessCensus
        //  - MERKLE_TREE_ONCHAIN_DYNAMIC_V1 -> census contract, asked about each batch root; the
        //    root stored here is its current one, for information only
        //  - CSP_EDDSA_BABYJUBJUB_V1 -> the CSP signer address as uint160 (ECDSA/secp256k1 in the guest)
        DAVINCITypes.CensusOrigin origin = census.censusOrigin;
        if (origin == DAVINCITypes.CensusOrigin.CENSUS_UNKNOWN) revert InvalidCensusOrigin();
        if (census.onchainAllowAnyValidRoot) revert InvalidCensusConfig();
        censusRoot = census.censusRoot;
        if (origin == DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1) {
            address censusContract = census.contractAddress;
            if (censusContract.code.length == 0) revert InvalidCensusAddress();
            (bool ok, uint256 current) = _censusCall(censusContract, abi.encodeCall(ICensusValidator.getCensusRoot, ()));
            if (!ok) revert InvalidCensusAddress();
            censusRoot = bytes32(current);
        } else {
            if (census.contractAddress != address(0)) revert InvalidCensusAddress();
            if (censusRoot == bytes32(0)) revert InvalidCensusRoot();
            if (origin == DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1 && uint256(censusRoot) >> 160 != 0) {
                revert InvalidCensusRoot();
            }
        }
        // CensusURI: where the sequencer downloads the census (Merkle) or voters get signatures (CSP)
        if (bytes(census.censusURI).length == 0) revert InvalidCensusURI();

        // validate status
        if (
            uint8(status) > MAX_STATUS
                || (status != DAVINCITypes.ProcessStatus.READY && status != DAVINCITypes.ProcessStatus.PAUSED)
        ) {
            revert InvalidStatus();
        }
    }

    /// @dev Stores a process's metadata and emits ProcessMetadataUpdated.
    function _storeMetadata(
        bytes31 processId,
        DAVINCITypes.Process storage p,
        string calldata metadataURI,
        bytes32 metadataHash
    ) private {
        p.metadataURI = metadataURI;
        p.metadataHash = metadataHash;
        emit ProcessMetadataUpdated(processId, metadataURI, metadataHash);
    }

    /// @dev A metadata URI must be non-empty and its SHA-256 non-zero.
    function _checkMetadata(string calldata metadataURI, bytes32 metadataHash) private pure {
        if (bytes(metadataURI).length == 0 || metadataHash == bytes32(0)) revert InvalidMetadata();
    }

    /**
     * @dev Rejects processes whose worst-case result could exceed the sequencer's
     *      bounded decryption search window.
     */
    function _validateMaxPossibleResultCap(uint256 maxVoters, uint256 maxValue) private pure {
        if (maxValue > MAX_POSSIBLE_RESULT_CAP / maxVoters) {
            revert MaxPossibleResultCapExceeded();
        }
    }

    /**
     * @notice Validates if a status transition is allowed
     * @param currentStatus The current status of the process
     * @param newStatus The new status to transition to
     * @return bool True if the transition is valid, false otherwise
     * @dev Status transition rules:
     * - READY -> PAUSED, CANCELED, ENDED
     * - PAUSED -> READY, CANCELED, ENDED
     * - ENDED -> RESULTS (automatically set when calling setProcessResults)
     * - CANCELED -> No transitions allowed
     * - RESULTS -> No transitions allowed
     */
    function _validateStatusTransition(DAVINCITypes.ProcessStatus currentStatus, DAVINCITypes.ProcessStatus newStatus)
        internal
        pure
        returns (bool)
    {
        if (newStatus == currentStatus) return false;

        if (
            currentStatus == DAVINCITypes.ProcessStatus.CANCELED || currentStatus == DAVINCITypes.ProcessStatus.RESULTS
                || currentStatus == DAVINCITypes.ProcessStatus.ENDED
        ) return false;

        if (currentStatus == DAVINCITypes.ProcessStatus.READY) {
            return newStatus == DAVINCITypes.ProcessStatus.PAUSED || newStatus == DAVINCITypes.ProcessStatus.CANCELED
                || newStatus == DAVINCITypes.ProcessStatus.ENDED;
        }

        if (currentStatus == DAVINCITypes.ProcessStatus.PAUSED) {
            return newStatus == DAVINCITypes.ProcessStatus.READY || newStatus == DAVINCITypes.ProcessStatus.CANCELED
                || newStatus == DAVINCITypes.ProcessStatus.ENDED;
        }

        return false;
    }

    /// @dev Loads a process, checking the id is well formed and the process exists.
    function _existingProcess(bytes31 processId) private view returns (DAVINCITypes.Process storage p) {
        if (processId == bytes31(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
    }

    /// @dev min(end + graceMaxTotal, max(end, lastVoteAt) + grace). An end within graceMaxTotal
    ///      of 2^256 (setProcessDuration allows it) never closes; below that nothing overflows,
    ///      as grace never exceeds graceMaxTotal.
    function _graceEnd(DAVINCITypes.Process storage p) private view returns (uint256) {
        uint256 end = p.startTime + p.duration;
        uint256 maxTotal = graceMaxTotal;
        if (end > type(uint256).max - maxTotal) return type(uint256).max;
        uint256 last = p.lastVoteAt;
        uint256 idleEnd = (last > end ? last : end) + p.grace;
        uint256 cap = end + maxTotal;
        return idleEnd < cap ? idleEnd : cap;
    }

    /// @dev Checks the census root a batch was proven against (BE integer). Origins 1, 2 and 4
    ///      must match the stored root. Origin 3 must be a root the census contract held at some
    ///      block in [creationBlock, block.number]: getRootBlockNumber returns the last block a
    ///      root was current, and 0 for a root it never held (rejected, unlike v0.0.49). A failed
    ///      or short answer is rejected too.
    function _checkCensusRoot(DAVINCITypes.Process storage p, bytes32 censusRoot) private view {
        if (p.census.censusOrigin != DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1) {
            if (censusRoot != p.census.censusRoot) revert InvalidCensusRoot();
            return;
        }
        (bool ok, uint256 rbn) = _censusCall(
            p.census.contractAddress, abi.encodeCall(ICensusValidator.getRootBlockNumber, (uint256(censusRoot)))
        );
        if (!ok || rbn == 0 || rbn > block.number || rbn < p.creationBlock) revert InvalidCensusRoot();
    }

    /// @dev STATICCALLs an origin-3 census contract with CENSUS_CALL_GAS and reads one word.
    ///      ok is false if the call fails or returns fewer than 32 bytes; nothing past the
    ///      first word is copied.
    function _censusCall(address census, bytes memory data) private view returns (bool ok, uint256 value) {
        assembly ("memory-safe") {
            ok := staticcall(CENSUS_CALL_GAS, census, add(data, 0x20), mload(data), 0x00, 0x20)
            ok := and(ok, gt(returndatasize(), 0x1f))
            value := mload(0x00)
        }
    }

    /// @dev The publics must be the 512-byte layout and report every guest check passed.
    function _checkGuestOk(bytes calldata publicValues) private pure {
        if (publicValues.length != PublicsLib.LENGTH) revert InvalidPublicValues();
        if (PublicsLib.word(publicValues, REG_OK) != 1 || PublicsLib.word(publicValues, REG_FAIL_MASK) != 0) {
            revert CircuitFailed();
        }
    }

    /// @dev Checks the blob arrays and the transaction's blobs against n_blobs, and the pair
    ///      digest the guest published.
    function _checkBlobsDigest(
        bytes calldata publicValues,
        bytes[] calldata commitments,
        bytes32[] calldata ys,
        bytes[] calldata kzgProofs
    ) private view returns (uint256 n) {
        n = PublicsLib.word(publicValues, REG_N_BLOBS);
        if (n == 0) revert NoBlobs();
        if (commitments.length != n || ys.length != n || kzgProofs.length != n) revert BlobCountMismatch();
        // No blob may follow the n the guest committed to: nodes decode every blob of the tx.
        if (BlobsLib.blobHash(n) != bytes32(0)) revert BlobCountMismatch();

        bytes memory pairs = new bytes(n * 80);
        for (uint256 i = 0; i < n; ++i) {
            bytes calldata c = commitments[i];
            if (c.length != 48) revert InvalidBlobCommitmentLength();
            if (kzgProofs[i].length != 48) revert InvalidKZGProofLength();
            bytes32 y = ys[i];
            assembly ("memory-safe") {
                let dst := add(add(pairs, 32), mul(i, 80))
                calldatacopy(dst, c.offset, 48)
                mstore(add(dst, 48), y)
            }
        }
        if (sha256(pairs) != PublicsLib.reg32(publicValues, REG_BLOBS_DIGEST)) revert InvalidBlobsDigest();
    }

    /// @dev Opens blob i of this transaction at z_i = sha256(BE32(pid) ‖ BE32(root) ‖ commitment_i)
    ///      mod r_bls, the point the guest evaluated it at, and checks y_i.
    function _verifyOpenings(
        bytes31 processId,
        bytes32 rootBefore,
        bytes[] calldata commitments,
        bytes32[] calldata ys,
        bytes[] calldata kzgProofs
    ) private view {
        bytes32 pidBE = bytes32(uint256(uint248(processId)));
        bytes32 rootBE = PublicsLib.reverse32(rootBefore);
        for (uint256 i = 0; i < commitments.length; ++i) {
            bytes32 versionedHash = BlobsLib.blobHash(i);
            if (versionedHash == bytes32(0)) revert MissingBlob(i);
            bytes32 z = bytes32(uint256(sha256(abi.encodePacked(pidBE, rootBE, commitments[i]))) % BLS_MODULUS);
            (bool ok, bytes memory out) =
                POINT_EVALUATION.staticcall(abi.encodePacked(versionedHash, z, ys[i], commitments[i], kzgProofs[i]));
            // Also rejects a chain where 0x0A is not the precompile (empty return data).
            if (!ok || out.length != 64) revert InvalidBlobOpening(i);
            (uint256 fieldElements, uint256 modulus) = abi.decode(out, (uint256, uint256));
            if (fieldElements != FIELD_ELEMENTS_PER_BLOB || modulus != BLS_MODULUS) revert InvalidBlobOpening(i);
        }
    }
}
