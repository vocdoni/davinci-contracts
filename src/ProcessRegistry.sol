// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {IProcessRegistry} from "./interfaces/IProcessRegistry.sol";
import {IZKVerifier} from "./interfaces/IZKVerifier.sol";
import {DAVINCITypes} from "./libraries/DAVINCITypes.sol";
import {ProcessIdLib} from "./libraries/ProcessIdLib.sol";
import {BlobsLib} from "./libraries/BlobsLib.sol";
import {StateRootLib} from "./libraries/StateRootLib.sol";
import {ICensusValidator} from "./interfaces/ICensusValidator.sol";

/**
 * @title ProcessRegistry
 * @notice This contract is responsible for storing processes data and managing their lifecycle.
 */
contract ProcessRegistry is IProcessRegistry {
    using ProcessIdLib for bytes31;
    using BlobsLib for bytes;

    /**
     * @notice The maximum value of the census origin.
     */
    uint8 public constant MAX_CENSUS_ORIGIN = 5;
    /**
     * @notice The maximum value of the process status.
     */
    uint8 public constant MAX_STATUS = 4;
    /**
     * @notice The index of the blob in the blob transaction.
     */
    uint8 public constant BLOB_INDEX = 0;
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
     * @notice The stVerifier is the address of the state transition ZK verifier contract.
     */
    address public stVerifier;
    /**
     * @notice The rVerifier is the address of the results ZK verifier contract.
     */
    address public rVerifier;
    /**
     * @notice The pidPrefix is the 4-byte prefix used in process IDs.
     * This is computed as the last 4 bytes of keccak256(abi.encodePacked(chainID, address(this))).
     */
    uint32 public pidPrefix;
    /**
     * @notice blobsDA is a boolean indicating if the contract uses EIP4844 blobs.
     * If true, the contract expects blob-related data in proofs and verifies them.
     * @dev This is required to ensure compatibility with networks that uses other DA mechanisms.
     */
    bool public blobsDA;

    /**
     * @notice Initializes the contract.
     * @param _chainID The ID of the chain.
     * @param _stVerifier The address of the state transition ZK verifier contract.
     * @param _rVerifier The address of the results ZK verifier contract.
     */
    constructor(uint32 _chainID, address _stVerifier, address _rVerifier, bool _blobsDA) {
        stVerifier = _stVerifier;
        rVerifier = _rVerifier;
        chainID = _chainID;
        blobsDA = _blobsDA;
        pidPrefix = ProcessIdLib.getPrefix(_chainID, address(this));
    }

    struct StateTransitionBatchProofInputs {
        uint256 rootHashBefore;
        uint256 rootHashAfter;
        uint256 votersCount;
        uint256 overwrittenVotesCount;
        uint256 censusRoot;
        uint256 blobCommitmentLimb0;
        uint256 blobCommitmentLimb1;
        uint256 blobCommitmentLimb2;
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
        return IZKVerifier(stVerifier).provingKeyHash();
    }

    /// @inheritdoc IProcessRegistry
    function getRVerifierVKeyHash() external view override returns (bytes32) {
        return IZKVerifier(rVerifier).provingKeyHash();
    }

    /// @inheritdoc IProcessRegistry
    function getNextProcessId(address organizationId) external view override returns (bytes31) {
        return ProcessIdLib.computeProcessId(pidPrefix, organizationId, processNonce[organizationId]);
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
        _validateNewProcess(processId, sender, status, maxVoters, ballotMode, census);

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
        p.latestStateRoot = StateRootLib.computeStateRoot(processId, census.censusOrigin, ballotMode, encryptionKey);
        p.metadataURI = metadata;
        p.ballotMode = ballotMode;
        p.census = census;
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
        if (processId == bytes31(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        DAVINCITypes.Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.organizationId != msg.sender) revert Unauthorized();
        if (p.census.censusOrigin != DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1) {
            revert CensusNotUpdatable();
        }

        // check census
        if (p.census.censusOrigin != census.censusOrigin) revert InvalidCensusOrigin();
        if (census.censusRoot == bytes32(0)) revert InvalidCensusRoot();
        if (bytes(census.censusURI).length == 0) revert InvalidCensusURI();
        if (
            p.census.censusOrigin == DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1
                && census.contractAddress == address(0)
        ) revert InvalidCensusAddress();

        // check ongoing process
        if (p.status != DAVINCITypes.ProcessStatus.READY && p.status != DAVINCITypes.ProcessStatus.PAUSED) {
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

        p.maxVoters = _maxVoters;

        emit ProcessMaxVotersChanged(processId, _maxVoters);
    }

    /// @inheritdoc IProcessRegistry
    function submitStateTransition(bytes31 processId, bytes calldata proof, bytes calldata input) external override {
        if (processId == bytes31(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        DAVINCITypes.Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.status != DAVINCITypes.ProcessStatus.READY) revert InvalidStatus();
        if (p.startTime + p.duration <= block.timestamp) revert InvalidTimeBounds();

        StateTransitionBatchProofInputs memory st = _decodeStateTransitionBatchProofInputs(input);
        if (p.census.censusOrigin == DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1) {
            uint256 rootBlockNumber = ICensusValidator(p.census.contractAddress).getRootBlockNumber(st.censusRoot);
            if (
                (!p.census.onchainAllowAnyValidRoot && rootBlockNumber < p.creationBlock)
                    || rootBlockNumber > block.number
            ) {
                revert InvalidCensusRoot();
            }
        } else {
            if (st.censusRoot != uint256(p.census.censusRoot)) revert InvalidCensusRoot();
        }

        // Validate state root before matches latest state root
        if (st.rootHashBefore != p.latestStateRoot) revert InvalidStateRoot();

        // Validate max votes not exceeded after including the new votes in the batch
        // but only if the batch contains new votes (votersCount - overwrittenVotesCount > 0)
        // Subsequent calls to this function will fail if maxVoters is not updated after surpassing the threshold.
        uint256 newVoters = st.votersCount - st.overwrittenVotesCount;
        if (newVoters > 0 && p.votersCount >= p.maxVoters) revert MaxVotersReached();

        if (blobsDA) {
            bytes memory blobCommitment =
                _blobCommitmentFromLimbs(st.blobCommitmentLimb0, st.blobCommitmentLimb1, st.blobCommitmentLimb2);
            bytes32 versionedHash = BlobsLib.calcBlobHashV1(blobCommitment);
            _verifyBlobDataIsAvailable(versionedHash);
        }

        IZKVerifier(stVerifier).verifyProof(proof, input);

        p.latestStateRoot = st.rootHashAfter;
        p.votersCount += newVoters;
        p.overwrittenVotesCount += st.overwrittenVotesCount;
        ++p.batchNumber;

        emit ProcessStateTransitioned(
            processId, msg.sender, st.rootHashBefore, st.rootHashAfter, p.votersCount, p.overwrittenVotesCount
        );
    }

    /// @inheritdoc IProcessRegistry
    function setProcessResults(bytes31 processId, bytes calldata proof, bytes calldata input) external override {
        if (processId == bytes31(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        DAVINCITypes.Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();

        // Cannot set results on CANCELLED or RESULTS processes
        if (p.status == DAVINCITypes.ProcessStatus.CANCELED || p.status == DAVINCITypes.ProcessStatus.RESULTS) {
            revert InvalidStatus();
        }

        // Require that the process has ended, either by status or by time
        if (p.status != DAVINCITypes.ProcessStatus.ENDED && p.startTime + p.duration > block.timestamp) {
            revert InvalidTimeBounds();
        }

        // Store the old status for the event
        DAVINCITypes.ProcessStatus oldStatus = p.status;

        IZKVerifier(rVerifier).verifyProof(proof, input);

        uint256[9] memory decompressedInput = abi.decode(input, (uint256[9]));

        if (decompressedInput[0] != p.latestStateRoot) {
            revert InvalidStateRoot();
        }

        uint256[] memory result = new uint256[](decompressedInput.length - 1);
        for (uint256 i = 1; i < decompressedInput.length; i++) {
            result[i - 1] = decompressedInput[i];
        }

        p.status = DAVINCITypes.ProcessStatus.RESULTS;
        p.result = result;

        emit ProcessStatusChanged(processId, oldStatus, DAVINCITypes.ProcessStatus.RESULTS);
        emit ProcessResultsSet(processId, msg.sender, result);
    }

    /**
     * @dev Validates inputs for a new process
     */
    function _validateNewProcess(
        bytes31 processId,
        address sender,
        DAVINCITypes.ProcessStatus status,
        uint256 maxVoters,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.Census calldata census
    ) private view {
        if (processes[processId].organizationId == sender) {
            revert ProcessAlreadyExists();
        }

        // validate ballot mode
        if (ballotMode.numFields == 0 || ballotMode.numFields > 8) revert InvalidMaxCount();
        if (ballotMode.groupSize > ballotMode.numFields) revert InvalidGroupSize();
        if (ballotMode.maxValueSum > 65535) revert InvalidMaxValueSum();
        if (ballotMode.minValue > ballotMode.maxValue) revert InvalidMaxMinValueBounds();
        if (ballotMode.minValueSum > ballotMode.maxValueSum) revert InvalidValueSumBounds();

        // validate max voters
        if (maxVoters == 0) revert InvalidMaxVoters();

        // validate census
        if (uint8(census.censusOrigin) > MAX_CENSUS_ORIGIN) revert InvalidCensusOrigin();
        if (
            census.censusOrigin != DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1
                && census.onchainAllowAnyValidRoot == true
        ) {
            revert InvalidCensusConfig();
        }
        // CensusRoot based on census origin:
        //  - MERKLE_TREE_OFFCHAIN_STATIC_V1 -> Merkle Root (fixed)
        //  - MERKLE_TREE_OFFCHAIN_DYNAMIC_V1 -> Merkle Root (could change via tx)
        //  - MERKLE_TREE_ONCHAIN_DYNAMIC_V1 -> Address of census manager contract (queried on each transition)
        //  - CSP_EDDSA_BABYJUBJUB_V1 -> CSP PubKey (fixed)
        if (census.censusOrigin != DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1) {
            if (census.censusRoot == bytes32(0)) revert InvalidCensusRoot();
        } else {
            if (census.contractAddress == address(0)) revert InvalidCensusAddress();
        }
        // CensusURI based on census origin:
        //  - MERKLE_TREE_OFFCHAIN_STATIC_V1 ──┬> URL where the sequencer can download the census snapshot used to compute the Merkle Proofs
        //  - MERKLE_TREE_OFFCHAIN_DYNAMIC_V1 ─┤
        //  - MERKLE_TREE_ONCHAIN_DYNAMIC_V1 ──┘
        //  - CSP_EDDSA_BABYJUBJUB_V1 > URL where the voters can generate their signatures
        if (bytes(census.censusURI).length == 0) revert InvalidCensusURI();

        // validate status
        if (
            uint8(status) > MAX_STATUS
                || (status != DAVINCITypes.ProcessStatus.READY && status != DAVINCITypes.ProcessStatus.PAUSED)
        ) {
            revert InvalidStatus();
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

    /// @notice Checks that blob data is available for the current transaction
    /// @dev Wrapper for BlobsLib.verifyBlobDataIsAvailable, that can be overridden in tests
    /// @param versionedHash The blob versioned hash
    function _verifyBlobDataIsAvailable(bytes32 versionedHash) internal view virtual {
        BlobsLib.verifyBlobDataIsAvailable(versionedHash);
    }

    /// @notice Decodes state transition batch proof inputs
    /// @dev Wrapper around abi.decode for (uint256[8]) produced by the sequencer.
    ///      Returns a named struct for readability and to avoid magic indices.
    /// @param input ABI-encoded batch inputs: (uint256[8])
    /// @return st The decoded inputs as StateTransitionBatchProofInputs
    function _decodeStateTransitionBatchProofInputs(bytes calldata input)
        internal
        pure
        returns (StateTransitionBatchProofInputs memory st)
    {
        uint256[8] memory d = abi.decode(input, (uint256[8]));

        st = StateTransitionBatchProofInputs({
            rootHashBefore: d[0],
            rootHashAfter: d[1],
            votersCount: d[2],
            overwrittenVotesCount: d[3],
            censusRoot: d[4],
            blobCommitmentLimb0: d[5],
            blobCommitmentLimb1: d[6],
            blobCommitmentLimb2: d[7]
        });
    }

    function _blobCommitmentFromLimbs(uint256 limb0, uint256 limb1, uint256 limb2)
        internal
        pure
        returns (bytes memory commitment)
    {
        commitment = new bytes(48);
        _writeCommitmentLimb(commitment, 0, limb0, 0);
        _writeCommitmentLimb(commitment, 16, limb1, 1);
        _writeCommitmentLimb(commitment, 32, limb2, 2);
    }

    function _writeCommitmentLimb(bytes memory commitment, uint256 offset, uint256 limb, uint8 idx) private pure {
        if (limb >> 128 != 0) revert InvalidBlobCommitmentLimb(idx);
        assembly ("memory-safe") {
            mstore(add(add(commitment, 32), offset), shl(128, limb))
        }
    }
}
