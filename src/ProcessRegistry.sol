// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { IProcessRegistry } from "./interfaces/IProcessRegistry.sol";
import { IZKVerifier } from "./interfaces/IZKVerifier.sol";
import { ProcessIdLib } from "./libraries/ProcessIdLib.sol";
import { BlobsLib } from "./libraries/BlobsLib.sol";

/**
 * @title ProcessRegistry
 * @notice This contract is responsible for storing processes data and managing their lifecycle.
 */
contract ProcessRegistry is IProcessRegistry {
    using ProcessIdLib for bytes32;

    /**
     * @notice The maximum value of the census origin.
     */
    uint8 public constant MAX_CENSUS_ORIGIN = 9;
    /**
     * @notice The maximum value of the process status.
     */
    uint8 public constant MAX_STATUS = 4;
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
     * @notice Initializes the contract.
     * @param _chainID The ID of the chain.
     * @param _stVerifier The address of the state transition ZK verifier contract.
     * @param _rVerifier The address of the results ZK verifier contract.
     */
    constructor(uint32 _chainID, address _stVerifier, address _rVerifier) {
        stVerifier = _stVerifier;
        rVerifier = _rVerifier;
        chainID = _chainID;
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
        return ProcessIdLib.computeProcessId(chainID, organizationId, processNonce[organizationId]);
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
        bytes32 processId = ProcessIdLib.computeProcessId(chainID, sender, processNonce[sender]);

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

        processCount++;
        processNonce[sender]++;

        emit ProcessCreated(processId, sender);
        return processId;
    }

    /// @inheritdoc IProcessRegistry
    function setProcessStatus(bytes32 processId, ProcessStatus newStatus) external override {
        if (processId == bytes32(0)) revert InvalidProcessId();
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
            uint256 newDuration = block.timestamp - p.startTime;
            p.duration = newDuration;
            emit ProcessDurationChanged(processId, newDuration);
        }

        emit ProcessStatusChanged(processId, oldStatus, newStatus);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessCensus(bytes32 processId, Census calldata census) external override {
        if (processId == bytes32(0)) revert InvalidProcessId();
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

    ///@inheritdoc IProcessRegistry
    function submitStateTransition(
        bytes32 processId,
        bytes calldata proof,
        bytes calldata input,
        bytes32 expectedBlobHash
    ) external override {
        if (processId == bytes32(0)) revert InvalidProcessId();
        Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.status != ProcessStatus.READY) revert InvalidStatus();
        if (p.startTime + p.duration <= block.timestamp) revert InvalidTimeBounds();

        // The tx **must** be type-3 and contain *one* blob at one of the available indexes.
        // We do not attempt to verify the full KZG proof on-chain – the
        // circuit already guarantees consistency between the votes and
        // the commitment, we only check availability.

        // TODO: Check all available blobs ?
        for (uint256 i = 0; i < 6; i++) {
            bytes32 actual = BlobsLib.blobHash(i);
            if (actual == expectedBlobHash) {
                break;
            }
            if (actual == bytes32(0)) {
                revert InvalidBlob();
            }
        }

        IZKVerifier(stVerifier).verifyProof(proof, input);

        uint256[4] memory decompressedInput = abi.decode(input, (uint256[4]));
        if (decompressedInput[0] != p.latestStateRoot) {
            revert InvalidStateRoot();
        }

        p.latestStateRoot = decompressedInput[1];
        p.voteCount += decompressedInput[2];
        p.voteOverwriteCount += decompressedInput[3];
        p.batchNumber++;

        emit ProcessStateRootUpdated(processId, msg.sender, decompressedInput[1], p.batchNumber);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessResults(bytes32 processId, bytes calldata proof, bytes calldata input) external override {
        if (processId == bytes32(0)) revert InvalidProcessId();
        Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.status != ProcessStatus.ENDED) revert InvalidStatus();
        if (p.startTime + p.duration > block.timestamp) revert InvalidTimeBounds();

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

        emit ProcessStatusChanged(processId, ProcessStatus.ENDED, ProcessStatus.RESULTS);
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
        if (ballotMode.maxCount == 0 || ballotMode.maxCount > 8) revert InvalidMaxCount();
        if (ballotMode.maxTotalCost > 65535) revert InvalidMaxTotalCost();
        if (ballotMode.minValue > ballotMode.maxValue) revert InvalidMaxMinValueBounds();
        if (ballotMode.minTotalCost > ballotMode.maxTotalCost) revert InvalidTotalCostBounds();

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
