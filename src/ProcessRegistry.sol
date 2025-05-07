// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import "./IProcessRegistry.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "./OrganizationRegistry.sol";
import "./IZKVerifier.sol";

/**
 * @title ProcessRegistry
 * @notice This contract is responsible for storing processes data and managing their lifecycle.
 */
contract ProcessRegistry is IProcessRegistry, Initializable, UUPSUpgradeable, OwnableUpgradeable {
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
     * @notice The process count is the number of processes created.
     */
    uint32 public processCount;
    /**
     * @notice The chain ID is the ID of the chain.
     */
    string public chainID;
    /**
     * @notice The verifier is the address of the ZK verifier contract.
     */
    address public verifier;

    /**
     * @notice Initializes the contract.
     * @param _chainID The ID of the chain.
     * @param _organizationRegistryAddress The address of the organization registry.
     * @param _verifier The address of the ZK verifier contract.
     */
    function initialize(
        string calldata _chainID,
        address _organizationRegistryAddress,
        address _verifier
    ) public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
        chainID = _chainID;
        organizationRegistryAddress = _organizationRegistryAddress;
        verifier = _verifier;
    }

    /// @inheritdoc IProcessRegistry
    function newProcess(
        ProcessStatus status,
        uint256 startTime,
        uint256 duration,
        BallotMode calldata ballotMode,
        Census calldata census,
        string calldata metadata,
        address organizationID,
        bytes32 processID,
        EncryptionKey calldata encryptionKey,
        uint256 initStateRoot
    ) external override {
        // check if the caller is an administrator
        if (!OrganizationRegistry(organizationRegistryAddress).isAdministrator(organizationID, msg.sender))
            revert NotOrganizationAdministrator();

        Process storage p = processes[processID];
        // check process does not exist
        if (p.organizationID != address(0)) revert ProcessAlreadyExists();
        // validate ballot mode parameters
        if (ballotMode.maxCount == 0) revert InvalidMaxCount();
        if (ballotMode.maxValue <= ballotMode.maxCount) revert InvalidMaxValue();
        // validate status
        if (uint8(status) > MAX_STATUS || (status != ProcessStatus.READY && status != ProcessStatus.PAUSED))
            revert InvalidStatus();
        // validate start time and duration
        uint256 currentTimestamp = block.timestamp;
        if (startTime <= currentTimestamp) revert InvalidStartTime();
        if (startTime + duration <= currentTimestamp) revert InvalidDuration();
        // validate census
        if (uint8(census.censusOrigin) > MAX_CENSUS_ORIGIN) revert InvalidCensus();

        p.status = status;
        p.startTime = startTime;
        p.duration = duration;
        p.organizationID = organizationID;
        p.encryptionKey = encryptionKey;
        p.latestStateRoot = initStateRoot;
        p.metadataURI = metadata;
        p.ballotMode = ballotMode;
        p.census = census;

        processCount++;

        emit ProcessCreated(processID, msg.sender);
    }

    /// @inheritdoc IProcessRegistry
    function getProcess(
        bytes32 processID
    )
        external
        view
        override
        returns (
            ProcessStatus status,
            address organizationID,
            EncryptionKey memory encryptionKey,
            uint256 latestStateRoot,
            uint256[] memory result,
            uint256 startTime,
            uint256 duration,
            uint256 voteCount,
            uint256 voteOverwriteCount,
            string memory metadataURI,
            BallotMode memory ballotMode,
            Census memory census
        )
    {
        Process memory p = processes[processID];
        return (
            p.status,
            p.organizationID,
            p.encryptionKey,
            p.latestStateRoot,
            p.result,
            p.startTime,
            p.duration,
            p.voteCount,
            p.voteOverwriteCount,
            p.metadataURI,
            p.ballotMode,
            p.census
        );
    }

    /// @inheritdoc IProcessRegistry
    function setProcessStatus(bytes32 processID, ProcessStatus newStatus) external override {
        Process storage p = processes[processID];
        address orgID = p.organizationID;
        // check process exist
        if (orgID == address(0)) revert ProcessNotFound();
        // check status
        ProcessStatus currentStatus = p.status;
        // if current status is READY => Can go to [ENDED, CANCELED, PAUSED].
        // if current status is PAUSED => Can go to [READY, ENDED, CANCELED].
        if (
            newStatus == currentStatus ||
            uint8(newStatus) > MAX_STATUS ||
            (currentStatus != ProcessStatus.READY && currentStatus != ProcessStatus.PAUSED)
        ) revert InvalidStatus();
        // check if the caller is an administrator
        if (!OrganizationRegistry(organizationRegistryAddress).isAdministrator(orgID, msg.sender))
            revert NotOrganizationAdministrator();

        p.status = newStatus;

        emit ProcessStatusChanged(processID, newStatus);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessCensus(bytes32 processID, Census calldata census) external override {
        // check census
        if (bytes(census.censusURI).length == 0) revert InvalidCensus();
        if (census.censusRoot == 0) revert InvalidCensus();

        Process storage p = processes[processID];
        address orgID = p.organizationID;
        // check process exists
        if (orgID == address(0)) revert ProcessNotFound();
        // check if sender is administrator
        if (!OrganizationRegistry(organizationRegistryAddress).isAdministrator(orgID, msg.sender)) {
            revert NotOrganizationAdministrator();
        }

        if (p.census.maxVotes > census.maxVotes) revert InvalidCensus();
        // check ongoing process
        if (p.status != ProcessStatus.READY && p.status != ProcessStatus.PAUSED) revert InvalidStatus();

        p.census.maxVotes = census.maxVotes;
        p.census.censusRoot = census.censusRoot;
        p.census.censusURI = census.censusURI;

        emit CensusUpdated(processID, census.censusRoot, census.censusURI, census.maxVotes);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessDuration(bytes32 processID, uint256 _duration) external override {
        // check valid duration
        if (_duration <= block.timestamp) revert InvalidDuration();

        Process storage p = processes[processID];
        address orgID = p.organizationID;
        // check process exists
        if (orgID == address(0)) revert ProcessNotFound();
        // check admin
        if (!OrganizationRegistry(organizationRegistryAddress).isAdministrator(orgID, msg.sender))
            revert NotOrganizationAdministrator();
        // check ongoing process
        if (p.status != ProcessStatus.READY && p.status != ProcessStatus.PAUSED) revert InvalidStatus();

        p.duration = _duration;

        emit ProcessDurationChanged(processID, _duration);
    }

    /// @inheritdoc IProcessRegistry
    function submitStateTransition(bytes32 processID, bytes calldata proof, bytes calldata input) external override {
        // check process exists
        Process storage p = processes[processID];
        if (p.organizationID == address(0)) revert ProcessNotFound();
        // check process is ongoing
        if (p.status != ProcessStatus.READY) revert InvalidStatus();
        // verify proof
        IZKVerifier(verifier).verifyProof(proof, input);
        // decompress data
        uint256[4] memory decompressedInput = abi.decode(input, (uint256[4]));
        // update process
        if (decompressedInput[0] != p.latestStateRoot) {
            revert InvalidStateRoot();
        }
        p.latestStateRoot = decompressedInput[1];
        p.voteCount = decompressedInput[2];
        p.voteOverwriteCount = decompressedInput[3];
        emit ProcessStateRootUpdated(processID, decompressedInput[1]);
    }

    /// @inheritdoc IProcessRegistry
    function setProcessResults(bytes32 processID, uint256[] calldata results) external override {}

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
