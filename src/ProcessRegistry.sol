// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { IProcessRegistry } from "./interfaces/IProcessRegistry.sol";
import { IZKVerifier } from "./interfaces/IZKVerifier.sol";
import { ProcessIdLib } from "./libraries/ProcessIdLib.sol";
import { BlobsLib } from "./libraries/BlobsLib.sol";
import "forge-std/console.sol";

/**
 * @title ProcessRegistry
 * @notice This contract is responsible for storing processes data and managing their lifecycle.
 */
contract ProcessRegistry is IProcessRegistry {
    using ProcessIdLib for bytes32;
    using BlobsLib for bytes;

    /**
     * @notice The maximum value of the census origin.
     */
    uint8 public constant MAX_CENSUS_ORIGIN = 9;
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
    mapping(bytes32 => Process) public processes;
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

    /// @inheritdoc IProcessRegistry
    function getProcess(bytes32 processId) external view override returns (Process memory) {
        return processes[processId];
    }

    /// @inheritdoc IProcessRegistry
    function getProcessEndTime(bytes32 processId) external view returns (uint256) {
        Process memory p = processes[processId];
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
    function getNextProcessId(address organizationId) external view override returns (bytes32) {
        return ProcessIdLib.computeProcessId(pidPrefix, organizationId, processNonce[organizationId]);
    }

    /// @inheritdoc IProcessRegistry
    function newProcess(
        ProcessStatus status,
        uint256 startTime,
        uint256 duration,
        BallotMode calldata ballotMode,
        Census calldata census,
        string calldata metadata,
        EncryptionKey calldata encryptionKey,
        uint256 initStateRoot
    ) external override returns (bytes32) {
        address sender = msg.sender;
        bytes32 processId = ProcessIdLib.computeProcessId(pidPrefix, sender, processNonce[sender]);

        // Validate process doesn't exist and validate inputs
        _validateNewProcess(processId, sender, status, startTime, duration, ballotMode, census);

        Process storage p = processes[processId];

        p.status = status;
        p.startTime = startTime;
        p.duration = duration;
        p.organizationId = sender;
        p.encryptionKey = encryptionKey;
        p.latestStateRoot = initStateRoot;
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
    function setProcessStatus(bytes32 processId, ProcessStatus newStatus) external override {
        if (processId == bytes32(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        if (uint8(newStatus) > MAX_STATUS) revert InvalidStatus();

        Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.organizationId != msg.sender) revert Unauthorized();

        // validate status transition
        ProcessStatus oldStatus = p.status;
        if (!_validateStatusTransition(oldStatus, newStatus)) revert InvalidStatus();

        p.status = newStatus;
        // if newStatus is ENDED, update duration to the time difference between current time and start time
        if (newStatus == ProcessStatus.ENDED) {
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
    function setProcessCensus(bytes32 processId, Census calldata census) external override {
        if (processId == bytes32(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.organizationId != msg.sender) revert Unauthorized();

        // check census
        if (p.census.censusOrigin != census.censusOrigin) revert InvalidCensusOrigin();
        if (p.census.maxVotes > census.maxVotes) revert InvalidMaxVotes();
        if (census.censusRoot == bytes32(0)) revert InvalidCensusRoot();
        if (bytes(census.censusURI).length == 0) revert InvalidCensusURI();

        // check ongoing process
        if (p.status != ProcessStatus.READY && p.status != ProcessStatus.PAUSED) revert InvalidStatus();
        if (p.startTime + p.duration <= block.timestamp) revert InvalidTimeBounds();

        p.census.maxVotes = census.maxVotes;
        p.census.censusRoot = census.censusRoot;
        p.census.censusURI = census.censusURI;

        emit CensusUpdated(processId, census.censusRoot, census.censusURI, census.maxVotes);
    }

    /// @inheritdoc IProcessRegistry
    /// @dev Note that the end time of the process is startTime + duration.
    function setProcessDuration(bytes32 processId, uint256 _duration) external override {
        if (processId == bytes32(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.organizationId != msg.sender) revert Unauthorized();

        // check ongoing process
        ProcessStatus status = p.status;
        if (status != ProcessStatus.READY && status != ProcessStatus.PAUSED) revert InvalidStatus();

        // check valid duration
        uint256 startTime = p.startTime;
        uint256 oldDuration = p.duration;
        if (
            _duration == 0 ||
            startTime + _duration <= block.timestamp ||
            startTime + _duration <= startTime + oldDuration
        ) revert InvalidDuration();

        p.duration = _duration;

        emit ProcessDurationChanged(processId, _duration);
    }

    /// @inheritdoc IProcessRegistry
    function submitStateTransition(bytes32 processId, bytes calldata proof, bytes calldata input) external override {
        if (processId == bytes32(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.status != ProcessStatus.READY) revert InvalidStatus();
        if (p.startTime + p.duration <= block.timestamp) revert InvalidTimeBounds();

        (
            uint256[9] memory decompressedInput,
            bytes memory blobCommitment,
            bytes memory blobProof
        ) = abi.decode(input, (uint256[9], bytes, bytes));

        if (decompressedInput[0] != p.latestStateRoot) {
            revert InvalidStateRoot();
        }

        if (blobsDA) {
            bytes32 versionedHash = BlobsLib.calcBlobHashV1(blobCommitment);
            bytes32 expectedHash = BlobsLib.blobHash(0);

            console.log("versionedHash:");
            console.logBytes32(versionedHash);
            console.log("expectedHash (blobHash(0)):");
            console.logBytes32(expectedHash);

            if (versionedHash != BlobsLib.blobHash(0)) revert InvalidBlobHash();

            bytes32 z = bytes32(decompressedInput[4]);

            bytes32 y;
            unchecked {
                uint256 MASK = (uint256(1) << 64) - 1; // 0xffffffffffffffff
                y = bytes32(
                    ((decompressedInput[5] & MASK) << 192) |
                    ((decompressedInput[6] & MASK) << 128) |
                    ((decompressedInput[7] & MASK) << 64) |
                    (decompressedInput[8] & MASK)
                );
            }

            bytes memory kzgInput = BlobsLib.buildKZGInput(
                versionedHash,
                z,
                y,
                blobCommitment,
                blobProof
            );

            if (!BlobsLib.verifyKZG(kzgInput)) revert BlobVerificationFailed();
        }

        IZKVerifier(stVerifier).verifyProof(proof, input);

        p.latestStateRoot = decompressedInput[1];
        p.voteCount += decompressedInput[2];
        p.voteOverwriteCount += decompressedInput[3];
        p.batchNumber++;

        emit ProcessStateRootUpdated(processId, msg.sender, decompressedInput[1]);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessResults(bytes32 processId, bytes calldata proof, bytes calldata input) external override {
        if (processId == bytes32(0)) revert InvalidProcessId();
        if (!ProcessIdLib.hasPrefix(processId, pidPrefix)) revert UnknownProcessIdPrefix();
        Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        
        // Cannot set results on CANCELLED or RESULTS processes
        if (p.status == ProcessStatus.CANCELED || p.status == ProcessStatus.RESULTS) revert InvalidStatus();
        
        // Require that the process has ended, either by status or by time
        if (p.status != ProcessStatus.ENDED && p.startTime + p.duration > block.timestamp) {
            revert InvalidTimeBounds();
        }
        
        // Store the old status for the event
        ProcessStatus oldStatus = p.status;

        IZKVerifier(rVerifier).verifyProof(proof, input);

        uint256[9] memory decompressedInput = abi.decode(input, (uint256[9]));

        if (decompressedInput[0] != p.latestStateRoot) {
            revert InvalidStateRoot();
        }

        uint256[] memory result = new uint256[](decompressedInput.length - 1);
        for (uint256 i = 1; i < decompressedInput.length; i++) {
            result[i - 1] = decompressedInput[i];
        }

        p.status = ProcessStatus.RESULTS;
        p.result = result;

        emit ProcessStatusChanged(processId, oldStatus, ProcessStatus.RESULTS);
        emit ProcessResultsSet(processId, msg.sender, result);
    }

    /**
     * @dev Validates inputs for a new process
     */
    function _validateNewProcess(
        bytes32 processId,
        address sender,
        ProcessStatus status,
        uint256 startTime,
        uint256 duration,
        BallotMode calldata ballotMode,
        Census calldata census
    ) private view {
        if (processes[processId].organizationId == sender) revert ProcessAlreadyExists();

        // validate ballot mode
        if (ballotMode.numFields == 0 || ballotMode.numFields > 8) revert InvalidMaxCount();
        if (ballotMode.maxValueSum > 65535) revert InvalidMaxValueSum();
        if (ballotMode.minValue > ballotMode.maxValue) revert InvalidMaxMinValueBounds();
        if (ballotMode.minValueSum > ballotMode.maxValueSum) revert InvalidValueSumBounds();

        // validate census
        if (uint8(census.censusOrigin) > MAX_CENSUS_ORIGIN) revert InvalidCensusOrigin();
        if (census.maxVotes == 0) revert InvalidMaxVotes();
        if (census.censusRoot == bytes32(0)) revert InvalidCensusRoot();
        if (bytes(census.censusURI).length == 0) revert InvalidCensusURI();

        // validate status
        if (uint8(status) > MAX_STATUS || (status != ProcessStatus.READY && status != ProcessStatus.PAUSED))
            revert InvalidStatus();

        // validate start time and duration
        uint256 currentTimestamp = block.timestamp;
        if (startTime == 0) {
            startTime = currentTimestamp;
        }
        if (startTime < currentTimestamp) revert InvalidStartTime();
        if (startTime + duration <= currentTimestamp) revert InvalidDuration();
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
    function _validateStatusTransition(
        ProcessStatus currentStatus,
        ProcessStatus newStatus
    ) internal pure returns (bool) {
        if (newStatus == currentStatus) return false;

        if (
            currentStatus == ProcessStatus.CANCELED ||
            currentStatus == ProcessStatus.RESULTS ||
            currentStatus == ProcessStatus.ENDED
        ) return false;

        if (currentStatus == ProcessStatus.READY) {
            return
                newStatus == ProcessStatus.PAUSED ||
                newStatus == ProcessStatus.CANCELED ||
                newStatus == ProcessStatus.ENDED;
        }

        if (currentStatus == ProcessStatus.PAUSED) {
            return
                newStatus == ProcessStatus.READY ||
                newStatus == ProcessStatus.CANCELED ||
                newStatus == ProcessStatus.ENDED;
        }

        return false;
    }
}
