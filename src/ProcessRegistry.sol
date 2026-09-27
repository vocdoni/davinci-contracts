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
     * @notice Initializes the contract.
     * @param _chainID The ID of the chain.
     * @param _ziskVerifier The ZisK PLONK verifier.
     * @param _batchProgramVK Program vk of the vote-batch guest.
     * @param _resultsProgramVK Program vk of the results guest.
     * @param _rootCVadcopFinal Root of the ZisK vadcop-final setup.
     * @param _ballotVKHash sha256 digest of the ballot proof VK (davinci.BallotVKLeaf).
     */
    constructor(
        uint32 _chainID,
        address _ziskVerifier,
        bytes32 _batchProgramVK,
        bytes32 _resultsProgramVK,
        bytes32 _rootCVadcopFinal,
        bytes32 _ballotVKHash
    ) {
        if (
            _ziskVerifier == address(0) || _batchProgramVK == bytes32(0) || _resultsProgramVK == bytes32(0)
                || _rootCVadcopFinal == bytes32(0) || _ballotVKHash == bytes32(0)
        ) revert InvalidVerifierConfig();
        ziskVerifier = IZiskVerifier(_ziskVerifier);
        batchProgramVK = _batchProgramVK;
        resultsProgramVK = _resultsProgramVK;
        rootCVadcopFinal = _rootCVadcopFinal;
        ballotVKHash = _ballotVKHash;
        chainID = _chainID;
        pidPrefix = ProcessIdLib.getPrefix(_chainID, address(this));
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
        string calldata metadata,
        DAVINCITypes.EncryptionKey calldata encryptionKey
    ) external override returns (bytes31) {
        address sender = msg.sender;
        bytes31 processId = ProcessIdLib.computeProcessId(pidPrefix, sender, processNonce[sender]);

        // Validate process doesn't exist and validate inputs
        bytes32 censusRoot =
            _validateNewProcess(processId, sender, status, maxVoters, ballotMode, census, encryptionKey);

        // validate start time, block and duration
        uint256 currentTimestamp = block.timestamp;
        if (startTime == 0) {
            startTime = currentTimestamp;
        }
        if (startTime < currentTimestamp) revert InvalidStartTime();
        if (startTime + duration <= currentTimestamp) revert InvalidDuration();

        DAVINCITypes.Process storage p = processes[processId];

        p.status = status;
        p.startTime = startTime;
        p.duration = duration;
        p.maxVoters = maxVoters;
        p.organizationId = sender;
        p.encryptionKey = encryptionKey;
        p.latestStateRoot = GenesisLib.root(processId, ballotMode, encryptionKey, census.censusOrigin, ballotVKHash);
        p.metadataURI = metadata;
        p.ballotMode = ballotMode;
        p.census = census;
        p.census.censusRoot = censusRoot;
        p.creationBlock = block.number;

        processCount++;
        processNonce[sender]++;

        emit ProcessCreated(processId, sender);
        return processId;
    }

    /// @inheritdoc IProcessRegistry
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

        p.status = newStatus;
        // if newStatus is ENDED, update duration to the time difference between current time and start time
        if (newStatus == DAVINCITypes.ProcessStatus.ENDED) {
            uint256 newDuration;
            if (block.timestamp >= p.startTime) {
                newDuration = block.timestamp - p.startTime;
            } else {
                newDuration = 0; // Process never started so duration is 0
            }
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
    /// @dev Note that the end time of the process is startTime + duration.
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
        uint256 oldDuration = p.duration;
        // Past the end the tally may already be public (results tx in the mempool): no reopening.
        if (startTime + oldDuration <= block.timestamp) revert InvalidTimeBounds();
        if (
            _duration == 0 || startTime + _duration <= block.timestamp
                || startTime + _duration <= startTime + oldDuration
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

        // check valid maxVoters
        if (_maxVoters == 0 || _maxVoters < p.votersCount) revert InvalidMaxVoters();
        _validateMaxPossibleResultCap(_maxVoters, p.ballotMode.maxValue);

        p.maxVoters = _maxVoters;

        emit ProcessMaxVotersChanged(processId, _maxVoters);
    }

    /// @inheritdoc IProcessRegistry
    /// @dev Permissionless: the proof, the blob openings and root continuity authenticate it.
    function submitStateTransition(
        bytes31 processId,
        bytes calldata publicValues,
        bytes calldata proofBytes,
        bytes[] calldata commitments,
        bytes32[] calldata ys,
        bytes[] calldata kzgProofs
    ) external override nonReentrant {
        DAVINCITypes.Process storage p = _existingProcess(processId);
        if (p.status != DAVINCITypes.ProcessStatus.READY) revert InvalidStatus();
        if (p.startTime + p.duration <= block.timestamp) revert InvalidTimeBounds();
        if (block.timestamp < p.startTime) revert InvalidTimeBounds();

        _checkGuestOk(publicValues);
        bytes32 rootBefore = PublicsLib.reg32(publicValues, REG_ROOT_BEFORE);
        if (rootBefore != p.latestStateRoot) revert InvalidStateRoot();
        _checkCensusRoot(p, PublicsLib.reverse32(PublicsLib.reg32(publicValues, REG_CENSUS_ROOT)));
        // occupied_before counts the distinct slots written so far, which is votersCount.
        uint256 votersCount = p.votersCount;
        if (PublicsLib.word(publicValues, REG_OCCUPIED_BEFORE) != votersCount) revert InvalidOccupiedBefore();

        uint256 overwrites = PublicsLib.word(publicValues, REG_OVERWRITES);
        uint256 newVoters = PublicsLib.word(publicValues, REG_VOTERS) - overwrites;
        if (votersCount + newVoters > p.maxVoters) revert MaxVotersReached();

        uint256 nBlobs = _checkBlobsDigest(publicValues, commitments, ys, kzgProofs);

        ziskVerifier.verifySnarkProof(batchProgramVK, rootCVadcopFinal, publicValues, proofBytes);

        _verifyOpenings(processId, rootBefore, commitments, ys, kzgProofs);

        bytes32 rootAfter = PublicsLib.reg32(publicValues, REG_ROOT_AFTER);
        p.latestStateRoot = rootAfter;
        p.votersCount = votersCount + newVoters;
        p.overwrittenVotesCount += overwrites;
        ++p.batchNumber;

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

        // Cannot set results on CANCELLED or RESULTS processes
        DAVINCITypes.ProcessStatus oldStatus = p.status;
        if (oldStatus == DAVINCITypes.ProcessStatus.CANCELED || oldStatus == DAVINCITypes.ProcessStatus.RESULTS) {
            revert InvalidStatus();
        }

        // Require that the process has ended, either by status or by time
        if (oldStatus != DAVINCITypes.ProcessStatus.ENDED && p.startTime + p.duration > block.timestamp) {
            revert InvalidTimeBounds();
        }

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

    /**
     * @dev Validates inputs for a new process and returns the census root to store.
     */
    function _validateNewProcess(
        bytes31 processId,
        address sender,
        DAVINCITypes.ProcessStatus status,
        uint256 maxVoters,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.Census calldata census,
        DAVINCITypes.EncryptionKey calldata encryptionKey
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

        if (!GenesisLib.isValidEncryptionKey(encryptionKey)) revert InvalidEncryptionKey();
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
