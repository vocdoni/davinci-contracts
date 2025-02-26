// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "./OrganizationRegistry.sol";
import "./Groth16Verifier.sol";

/**
 * @title ProcessRegistry
 * @notice This contract is responsible for storing processes data and managing their lifecycle.
 */
contract ProcessRegistry is Initializable, UUPSUpgradeable, OwnableUpgradeable {
    /*
     * @notice Emitted when a new process is created.
     * @param processID The ID of the process.
     * @param creator The address of the creator of the process.
     */
    event ProcessCreated(bytes32 indexed processID, address indexed creator);
    /*
     * @notice Emitted when the status of a process is modified.
     * @param processID The ID of the process.
     * @param newStatus The new status of the process.
     */
    event ProcessStatusChanged(bytes32 indexed processID, ProcessStatus newStatus);
    /*
     * @notice Emitted when the census of a process is updated.
     * @param processID The ID of the process.
     * @param censusRoot The new root of the census.
     * @param censusURI The URI of the census.
     * @param maxVotes The maximum number of votes.
     */
    event CensusUpdated(bytes32 indexed processID, bytes32 censusRoot, string censusURI, uint256 maxVotes);
    /*
     * @notice Emitted when the duration of a process is modified.
     * @param processID The ID of the process.
     * @param duration The new duration of the process.
     */
    event ProcessDurationChanged(bytes32 indexed processID, uint256 duration);
    /*
     * @notice Emitted when the state root of a process is updated.
     * @param processID The ID of the process.
     * @param newStateRoot The new state root of the process.
     */
    event ProcessStateRootUpdated(bytes32 indexed processID, bytes32 newStateRoot);

    /**
     * @notice The process status defines the current state of the process.
     */
    enum ProcessStatus {
        READY,
        ENDED,
        CANCELED,
        PAUSED,
        RESULTS
    }

    /**
     * @notice The census origin defines the origin of the census data. It affects the way the census is handled.
     */
    enum CensusOrigin {
        CENSUS_UNKNOWN,
        OFF_CHAIN_TREE,
        OFF_CHAIN_TREE_WEIGHTED,
        OFF_CHAIN_CA,
        ERC20,
        ERC721,
        ERC1155,
        ERC777,
        MINI_ME,
        FARCASTER_FRAME
    }

    /**
     * @notice The ballot mode define the parameters of the vote.
     * @param costFromWeight If weighted census, the ballot weight is used as maxTotalCost.
     * @param forceUniqueness Choices cannot appear twice or more.
     * @param maxCount The maximum number of field per ballot.
     * @param costExponent The exponent that will be used to compute the "cost" of the field values.
     * @param maxValue The maximum value for all fields.
     * @param minValue The minimum value for all fields.
     * @param maxTotalCost Maximum limit on the total sum of all ballot fields' values. 0 => Not applicable.
     * @param minTotalCost Minimum limit on the total sum of all ballot fields' values. 0 => Not applicable.
     */
    struct BallotMode {
        bool costFromWeight;
        bool forceUniqueness;
        uint8 maxCount;
        uint8 costExponent;
        uint256 maxValue;
        uint256 minValue;
        uint256 maxTotalCost;
        uint256 minTotalCost;
    }

    /**
     * @notice The census defines the parameters of the census.
     * @param censusOrigin The origin of the census.
     * @param maxVotes The maximum number of votes.
     * @param censusRoot The root of the census.
     * @param censusURI The URI of the census.
     */
    struct Census {
        CensusOrigin censusOrigin;
        uint256 maxVotes;
        bytes32 censusRoot;
        string censusURI;
    }

    /**
     * @notice The process ID is a unique identifier for a process.
     * @param nonce The nonce of the process.
     * @param organizationID The ID of the organization.
     * @param chainID The ID of the chain.
     */
    struct ProcessID {
        uint256 nonce;
        address organizationID;
        string chainID;
    }

    /**
     * @notice EcryptionKey of a process
     * @param x value of the X coordinate on the curve
     * @param y value of the Y coordinate on the curve
     */
    struct EncryptionKey {
        uint256 x;
        uint256 y;
    }

    /**
     * @notice The process defines the parameters of the process.
     * @param status The status of the process.
     * @param organizationId The ID of the organization.
     * @param encryptionKey The encryption key of the process.
     * @param latestStateRoot The latest state root of the process.
     * @param result The result of the process.
     * @param startTime The start time of the process.
     * @param duration The duration of the process.
     * @param metadataURI The URI of the metadata.
     * @param ballotMode The ballot mode.
     * @param census The census of the process.
     */
    struct Process {
        ProcessStatus status;
        address organizationId;
        EncryptionKey encryptionKey;
        bytes32 latestStateRoot;
        uint256[] result;
        uint256 startTime;
        uint256 duration;
        string metadataURI;
        BallotMode ballotMode;
        Census census;
    }

    /**
     * @notice The process mapping is a mapping of process IDs to processes.
     */
    mapping(bytes32 => Process) public processes;
    /**
     * @notice The organization registry is the contract address of the organization registry.
     */
    address public organizationRegistry;
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
     * @param _organizationRegistry The address of the organization registry.
     */
    function initialize(string calldata _chainID, address _organizationRegistry, address _verifier) public initializer {
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();
        chainID = _chainID;
        organizationRegistry = _organizationRegistry;
        verifier = _verifier;
    }

    /**
     * @notice Creates a new process.
     * @param _status The initial status of the process.
     * @param _startTime The start time of the process.
     * @param _duration The duration of the process.
     * @param _ballotMode The ballot mode of the process.
     * @param _census The census of the process.
     * @param _metadata The URI of the metadata.
     * @param _organizationID The ID of the organization.
     * @param _processID The ID of the process.
     * @param _encryptionKey The public key of the encryption.
     * @param _initStateRoot The initial state root.
     */
    function newProcess(
        ProcessStatus _status,
        uint256 _startTime,
        uint256 _duration,
        BallotMode calldata _ballotMode,
        Census calldata _census,
        string calldata _metadata,
        address _organizationID,
        bytes32 _processID,
        EncryptionKey calldata _encryptionKey,
        bytes32 _initStateRoot
    ) public {
        require(_ballotMode.maxCount > 0, "NewProcess: invalid maxCount");
        require(_ballotMode.maxValue > _ballotMode.maxCount, "NewProcess: maxCount > maxValue");
        require(
            _status == ProcessStatus.READY || _status == ProcessStatus.PAUSED,
            "NewProcess: invalid status"
        );
        require(_startTime > block.timestamp, "NewProcess: invalid startTime");
        require(_startTime + _duration > block.timestamp, "NewProcess: invalid duration");
        require(
            OrganizationRegistry(organizationRegistry).isAdministrator(_organizationID, msg.sender),
            "NewProcess: not an administrator"
        );

        if (processes[_processID].organizationId != address(0)) {
            revert("NewProcess: process already exists");
        }

        Process memory p = Process({
            status: _status,
            startTime: _startTime,
            duration: _duration,
            organizationId: _organizationID,
            encryptionKey: _encryptionKey,
            latestStateRoot: _initStateRoot,
            result: new uint256[](0),
            metadataURI: _metadata,
            ballotMode: _ballotMode,
            census: _census
        });

        processes[_processID] = p;
        processCount++;

        emit ProcessCreated(_processID, msg.sender);
    }

    /**
     * @notice Returns the process data.
     * @param _processID The ID of the process.
     * @return The process data.
     */
    function getProcess(bytes32 _processID) public view returns (Process memory) {
        return processes[_processID];
    }

    /**
     * @notice Sets the status of a process.
     * @param _processID The ID of the process.
     * @param _newStatus The new status of the process.
     */
    function setProcessStatus(bytes32 _processID, ProcessStatus _newStatus) public {
        require(
            OrganizationRegistry(organizationRegistry).isAdministrator(processes[_processID].organizationId, msg.sender),
            "SetProcessStatus: not an administrator"
        );

        ProcessStatus currentStatus = processes[_processID].status;
        if (currentStatus != ProcessStatus.READY && currentStatus != ProcessStatus.PAUSED) {
            // When currentStatus is [ENDED, CANCELED, RESULTS], no update is allowed
            revert("Process terminated");
        }

        // If currentStatus is READY => Can go to [ENDED, CANCELED, PAUSED].
        // If currentStatus is PAUSED => Can go to [READY, ENDED, CANCELED].
        require(_newStatus != currentStatus, "Must differ");

        processes[_processID].status = _newStatus;

        emit ProcessStatusChanged(_processID, _newStatus);
    }

    /**
     * @notice Sets the census of a process.
     * @param _processID The ID of the process.
     * @param _census The census of the process.
     */
    function setProcessCensus(bytes32 _processID, Census calldata _census) public {
        require(
            OrganizationRegistry(organizationRegistry).isAdministrator(processes[_processID].organizationId, msg.sender),
            "SetProcessCensus: not an administrator"
        );

        // check census URI is not empty
        require(bytes(_census.censusURI).length > 0, "Empty URI");
        // check census root is not empty
        require(_census.censusRoot != 0, "Empty root");

        // check if the process exists
        require(processes[_processID].organizationId != address(0), "Process not found");

        // Only if the process is ongoing
        require(
            processes[_processID].status == ProcessStatus.READY || processes[_processID].status == ProcessStatus.PAUSED,
            "Process terminated"
        );

        if (processes[_processID].census.maxVotes < _census.maxVotes) {
            processes[_processID].census.maxVotes = _census.maxVotes;
        }

        processes[_processID].census.censusRoot = _census.censusRoot;
        processes[_processID].census.censusURI = _census.censusURI;

        emit CensusUpdated(_processID, _census.censusRoot, _census.censusURI, _census.maxVotes);
    }

    /**
     * @notice Sets the duration of a process.
     * @param _processID The ID of the process.
     * @param _duration The new duration of the process.
     */
    function setProcessDuration(bytes32 _processID, uint256 _duration) public {
        require(
            OrganizationRegistry(organizationRegistry).isAdministrator(processes[_processID].organizationId, msg.sender),
            "SetProcessDuration: not an administrator"
        );

        // check if the process exists
        require(processes[_processID].organizationId != address(0), "Process not found");

        // Only if the process is ongoing
        require(
            processes[_processID].status == ProcessStatus.READY || processes[_processID].status == ProcessStatus.PAUSED,
            "Process terminated"
        );

        require(_duration > block.timestamp, "Invalid duration");
        processes[_processID].duration = _duration;

        emit ProcessDurationChanged(_processID, _duration);
    }

    /**
     * @notice Ends a process.
     * @param _processID The ID of the process.
     */
    function endProcess(bytes32 _processID) public {
        require(
            OrganizationRegistry(organizationRegistry).isAdministrator(processes[_processID].organizationId, msg.sender),
            "endProcess: not an administrator"
        );
        require(
            processes[_processID].status == ProcessStatus.READY || processes[_processID].status == ProcessStatus.PAUSED,
            "Process terminated"
        );
        processes[_processID].status = ProcessStatus.ENDED;

        emit ProcessStatusChanged(_processID, ProcessStatus.ENDED);
    }

    /**
     * @notice Used to submit a state transition.
     * @param _processID The ID of the process.
     * @param _oldRoot The old state root.
     * @param _newRoot The new state root.
     * @param _proof The proof of the state transition.
     */
    function submitStateTransition(bytes32 _processID, bytes32 _oldRoot, bytes32 _newRoot, bytes calldata _proof)
        public
    {
        require(processes[_processID].organizationId != address(0), "Process not found");
        require(
            processes[_processID].status != ProcessStatus.RESULTS
                && processes[_processID].status != ProcessStatus.CANCELED,
            "Invalid status for submitting state transition"
        );
        require(processes[_processID].latestStateRoot == _oldRoot, "Invalid old root");
        // TODO verify proof
        // update state root
        processes[_processID].latestStateRoot = _newRoot;
        emit ProcessStateRootUpdated(_processID, _newRoot);
    }

    /**
     * @notice Sets the result of a process.
     * @param _processID The ID of the process.
     * @param _result The result of the process.
     * @param _proof The proof of the result.
     */
    function setProcessResult(bytes32 _processID, uint256[] calldata _result, bytes calldata _proof) public {
        // require sequencer from sequencer registry
        // TODO

        require(processes[_processID].organizationId != address(0), "Process not found");
        require(processes[_processID].status == ProcessStatus.ENDED, "Process not ended");

        // TODO verify proof

        processes[_processID].result = _result;
        processes[_processID].status = ProcessStatus.RESULTS;
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
