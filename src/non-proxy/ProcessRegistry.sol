// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import "../IProcessRegistry.sol";
import "./OrganizationRegistry.sol";
import "../IZKVerifier.sol";

/**
 * @title ProcessRegistry
 * @notice This contract is responsible for storing processes data and managing their lifecycle.
 */
contract ProcessRegistry is IProcessRegistry {
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
     * @notice The organization registry address is the contract address of the organization registry.
     */
    address public organizationRegistryAddress;
    /**
     * @notice The sequencer registry address is the sequencer registry contract that manages sequencers.
     */
    address public sequencerRegistryAddress;
    /**
     * @notice The process count is the number of processes created.
     */
    uint32 public processCount;
    /**
     * @notice The chain ID is the ID of the chain.
     */
    string public chainID;
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
     * @param _organizationRegistryAddress The address of the organization registry.
     * @param _stVerifier The address of the state transition ZK verifier contract.
     * @param _rVerifier The address of the results ZK verifier contract.
     */
    constructor(string memory _chainID, address _organizationRegistryAddress, address _stVerifier, address _rVerifier) {
        chainID = _chainID;
        organizationRegistryAddress = _organizationRegistryAddress;
        stVerifier = _stVerifier;
        rVerifier = _rVerifier;
    }

    /// @inheritdoc IProcessRegistry
    function newProcess(
        ProcessStatus status,
        uint256 startTime,
        uint256 duration,
        BallotMode calldata ballotMode,
        Census calldata census,
        string calldata metadata,
        address organizationId,
        bytes32 processId,
        EncryptionKey calldata encryptionKey,
        uint256 initStateRoot
    ) external override {
        // check if the caller is an administrator
        if (!OrganizationRegistry(organizationRegistryAddress).isAdministrator(organizationId, msg.sender))
            revert NotOrganizationAdministrator();

        Process storage p = processes[processId];
        // check process does not exist
        if (p.organizationId != address(0)) revert ProcessAlreadyExists();

        // validate ballot mode parameters
        if (ballotMode.maxCount == 0) revert InvalidMaxCount();
        if (ballotMode.maxValue < ballotMode.minValue) revert InvalidMinValue();
        if (!ballotMode.costFromWeight && ballotMode.maxTotalCost == 0) revert InvalidMaxTotalCost();
        if (ballotMode.maxTotalCost > 0 && (ballotMode.maxTotalCost < ballotMode.minTotalCost))
            revert InvalidTotalCostBounds();
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

        p.status = status;
        p.startTime = startTime;
        p.duration = duration;
        p.organizationId = organizationId;
        p.encryptionKey = encryptionKey;
        p.latestStateRoot = initStateRoot;
        p.metadataURI = metadata;
        p.ballotMode = ballotMode;
        p.census = census;

        processCount++;

        emit ProcessCreated(processId, msg.sender);
    }

    /// @inheritdoc IProcessRegistry
    function getProcess(bytes32 processId) external view override returns (Process memory) {
        return processes[processId];
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
    function setProcessStatus(bytes32 processId, ProcessStatus newStatus) external override {
        Process storage p = processes[processId];
        address orgID = p.organizationId;
        // check process exist
        if (orgID == address(0)) revert ProcessNotFound();

        // check if the caller is an administrator
        if (!OrganizationRegistry(organizationRegistryAddress).isAdministrator(orgID, msg.sender))
            revert NotOrganizationAdministrator();

        // Validate status transition
        if (!__validateStatusTransition(p.status, newStatus)) revert InvalidStatus();

        p.status = newStatus;

        emit ProcessStatusChanged(processId, newStatus);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessCensus(bytes32 processId, Census calldata census) external override {
        Process storage p = processes[processId];
        address orgID = p.organizationId;
        // check process exists
        if (orgID == address(0)) revert ProcessNotFound();
        // check if sender is administrator
        if (!OrganizationRegistry(organizationRegistryAddress).isAdministrator(orgID, msg.sender)) {
            revert NotOrganizationAdministrator();
        }

        if (p.census.censusOrigin != census.censusOrigin) revert InvalidCensusOrigin(); // enforce census origin cannot be changed
        if (census.maxVotes == 0 || p.census.maxVotes > census.maxVotes) revert InvalidMaxVotes();
        if (census.censusRoot == bytes32(0)) revert InvalidCensusRoot();
        if (bytes(census.censusURI).length == 0) revert InvalidCensusURI();

        // check ongoing process
        if (p.status != ProcessStatus.READY && p.status != ProcessStatus.PAUSED) revert InvalidStatus();

        p.census.maxVotes = census.maxVotes;
        p.census.censusRoot = census.censusRoot;
        p.census.censusURI = census.censusURI;

        emit CensusUpdated(processId, census.censusRoot, census.censusURI, census.maxVotes);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessDuration(bytes32 processId, uint256 _duration) external override {
        // check valid duration
        if (_duration <= block.timestamp) revert InvalidDuration();

        Process storage p = processes[processId];
        address orgID = p.organizationId;
        // check process exists
        if (orgID == address(0)) revert ProcessNotFound();
        // check admin
        if (!OrganizationRegistry(organizationRegistryAddress).isAdministrator(orgID, msg.sender))
            revert NotOrganizationAdministrator();
        // check ongoing process
        if (p.status != ProcessStatus.READY && p.status != ProcessStatus.PAUSED) revert InvalidStatus();

        p.duration = _duration;

        emit ProcessDurationChanged(processId, _duration);
    }

    /// @inheritdoc IProcessRegistry
    function submitStateTransition(bytes32 processId, bytes calldata proof, bytes calldata input) external override {
        // check process exists
        Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        // check process is ongoing
        if (p.status != ProcessStatus.READY) revert InvalidStatus();
        // verify proof
        IZKVerifier(stVerifier).verifyProof(proof, input);
        // decompress data
        uint256[4] memory decompressedInput = abi.decode(input, (uint256[4]));
        // update process
        if (decompressedInput[0] != p.latestStateRoot) {
            revert InvalidStateRoot();
        }
        p.latestStateRoot = decompressedInput[1];
        p.voteCount += decompressedInput[2];
        p.voteOverwriteCount += decompressedInput[3];
        emit ProcessStateRootUpdated(processId, decompressedInput[1]);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessResults(bytes32 processId, bytes calldata proof, bytes calldata input) external override {
        // check process exists
        Process storage p = processes[processId];
        if (p.organizationId == address(0)) revert ProcessNotFound();
        if (p.status == ProcessStatus.RESULTS || p.status == ProcessStatus.CANCELED) revert InvalidStatus();
        if (!_checkProcessEnded(p)) revert CannotAcceptResult();
        // verify proof
        IZKVerifier(rVerifier).verifyProof(proof, input);
        // decompress data
        uint256[9] memory decompressedInput = abi.decode(input, (uint256[9]));
        // update process
        if (decompressedInput[0] != p.latestStateRoot) {
            revert InvalidStateRoot();
        }
        // loop through the result and update the process
        uint256[] memory result = new uint256[](decompressedInput.length - 1);
        for (uint256 i = 1; i < decompressedInput.length; i++) {
            result[i - 1] = decompressedInput[i];
        }
        p.status = ProcessStatus.RESULTS;
        p.result = result;

        emit ProcessStatusChanged(processId, ProcessStatus.RESULTS);
        emit ProcessResultsSet(processId, result);
    }

    // internal functions

    /**
     * @notice Validates if a status transition is allowed
     * @param currentStatus The current status of the process
     * @param newStatus The new status to transition to
     * @return bool True if the transition is valid, false otherwise
     */
    function __validateStatusTransition(
        ProcessStatus currentStatus,
        ProcessStatus newStatus
    ) internal pure returns (bool) {
        // Cannot transition to the same status
        if (newStatus == currentStatus) return false;

        // CANCELED and RESULTS states cannot be changed
        if (currentStatus == ProcessStatus.CANCELED || currentStatus == ProcessStatus.RESULTS) return false;

        // Validate transitions based on current status
        if (currentStatus == ProcessStatus.READY) {
            // READY can only go to PAUSED, CANCELED, or ENDED
            return (newStatus == ProcessStatus.PAUSED ||
                newStatus == ProcessStatus.CANCELED ||
                newStatus == ProcessStatus.ENDED);
        } else if (currentStatus == ProcessStatus.PAUSED) {
            // PAUSED can only go to READY, CANCELED, or ENDED
            return (newStatus == ProcessStatus.READY ||
                newStatus == ProcessStatus.CANCELED ||
                newStatus == ProcessStatus.ENDED);
        } else if (currentStatus == ProcessStatus.ENDED) {
            // ENDED can only go to RESULTS or CANCELED
            return (newStatus == ProcessStatus.RESULTS || newStatus == ProcessStatus.CANCELED);
        }

        return false;
    }

    function _checkProcessEnded(Process memory p) internal view returns (bool) {
        if (p.status == ProcessStatus.ENDED || (p.startTime + p.duration) < block.timestamp) {
            return true;
        }
        return false;
    }
}
