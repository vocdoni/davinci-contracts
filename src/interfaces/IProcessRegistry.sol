// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/**
 * @title IProcessRegistry
 * @author Vocdoni Association
 * @notice The Process Registry contract interface.
 */
interface IProcessRegistry {
    /// EVENTS ///

    /*
     * @notice Emitted when a new process is created.
     * @param processId The ID of the process.
     * @param creator The address of the creator of the process.
     */
    event ProcessCreated(bytes32 indexed processId, address indexed creator);
    /*
     * @notice Emitted when the census of a process is updated.
     * @param processId The ID of the process.
     * @param censusRoot The new root of the census.
     * @param censusURI The URI of the census.
     */
    event CensusUpdated(bytes32 indexed processId, bytes32 censusRoot, string censusURI);
    /*
     * @notice Emitted when the duration of a process is modified.
     * @param processId The ID of the process.
     * @param duration The new duration of the process.
     */
    event ProcessDurationChanged(bytes32 indexed processId, uint256 duration);
    /*
     * @notice Emitted when the state root of a process is updated.
     * @param processId The ID of the process.
     * @param sender The address of the sender.
     * @param newStateRoot The new state root of the process.
     */
    event ProcessStateTransitioned(
        bytes32 indexed processId,
        address indexed sender,
        uint256 oldStateRoot,
        uint256 newStateRoot,
        uint256 newVotersCount,
        uint256 newOverwrittenVotesCount
    );

    /*
     * @notice Emitted when the results of a process are set.
     * @param processId The ID of the process.
     * @param sender The address of the sender.
     * @param result The result of the process.
     */
    event ProcessResultsSet(bytes32 indexed processId, address indexed sender, uint256[] result);
    /**
     * @notice Emitted when a process status is modified
     * @param processId The ID of the process
     * @param oldStatus The previous status of the process
     * @param newStatus The new status of the process
     */
    event ProcessStatusChanged(bytes32 indexed processId, ProcessStatus oldStatus, ProcessStatus newStatus);
    /**
     * @notice Emitted when the max voters of a process is modified
     * @param processId The ID of the process
     * @param maxVoters The new max voters of the process
     */
    event ProcessMaxVotersChanged(bytes32 indexed processId, uint256 maxVoters);

    /// ERRORS ///

    /**
     * @notice InvalidStatus error is emitted when the status of the process is invalid.
     */
    error InvalidStatus();
    /**
     * @notice InvalidStartTime error is emitted when the start time of the process is invalid.
     */
    error InvalidStartTime();
    /**
     * @notice InvalidDuration error is emitted when the duration of the process is invalid.
     */
    error InvalidDuration();
    /**
     * @notice InvalidMaxVoters error is emitted when the maximum number of voters is invalid.
     */
    error InvalidMaxVoters();
    /**
     * @notice MaxVotersReached error is emitted when the maximum number of voters has been reached.
     */
    error MaxVotersReached();
    /**
     * @notice InvalidMaxCount error is emitted when the maximum count of the ballot mode is invalid.
     */
    error InvalidMaxCount();
    /**
     * @notice InvalidMaxValue error is emitted when the maximum value of the ballot mode is invalid.
     */
    error InvalidMaxValue();
    /**
     * @notice InvalidMinValue error is emitted when the minimum value of the ballot mode is invalid.
     */
    error InvalidMinValue();
    /**
     * @notice InvalidMaxValueSum error is emitted when the maximum value sum of the ballot mode is invalid.
     */
    error InvalidMaxValueSum();
    /**
     * @notice InvalidMinTotalCost error is emitted when the minimum total cost of the ballot mode is invalid.
     */
    error InvalidMinTotalCost();
    /**
     * @notice InvalidValueSumBounds error is emitted when the total cost bounds of the ballot mode are invalid.
     */
    error InvalidValueSumBounds();
    /**
     * @notice InvalidMaxMinValueBounds error is emitted when the maximum and minimum value bounds are invalid.
     */
    error InvalidMaxMinValueBounds();
    /**
     * @notice InvalidUniqueValues error is emitted when the unique values are invalid.
     */
    error InvalidUniqueValues();
    /**
     * @notice InvalidCensusRoot error is emitted when the census root is invalid.
     */
    error InvalidCensusRoot();
    /**
     * @notice InvalidCensusURI error is emitted when the census URI is invalid.
     */
    error InvalidCensusURI();
    /**
     * @notice InvalidCensusOrigin error is emitted when the census origin is invalid.
     */
    error InvalidCensusOrigin();
    /**
     * @notice InvalidBlobCommitmentLimb error is emitted when a blob commitment limb exceeds 16 bytes.
     */
    error InvalidBlobCommitmentLimb(uint8 limbIndex);
    /**
     * @notice InvalidStateRoot error is emitted when a state root is invalid.
     */
    error InvalidStateRoot();
    /**
     * @notice ProcessAlreadyExists error is emitted when the process already exists.
     */
    error ProcessAlreadyExists();
    /**
     * @notice ProcessNotFound error is emitted when a process is not found
     */
    error ProcessNotFound();
    /**
     * @notice CannotAcceptResult error is emitted when a process cannot allow the results to be set.
     */
    error CannotAcceptResult();
    /**
     * @notice Thrown when the process ID is invalid (zero)
     */
    error InvalidProcessId();
    /**
     * @notice Thrown when the process ID prefix is unknown (does not match this contract)
     */
    error UnknownProcessIdPrefix();
    /**
     * @notice Thrown when attempting to transition to RESULTS state before process has ended
     */
    error ProcessNotEnded();
    /**
     * @notice Thrown when the process time bounds are invalid
     */
    error InvalidTimeBounds();
    /**
     * @notice Thrown when the proof is invalid.
     */
    error ProofInvalid();
    /**
     * @notice Thrown when the census is not updatable.
     */
    error CensusNotUpdatable();
    /**
     * @notice Thrown when the sender is not authorized to perform the action.
     */
    error Unauthorized();

    /// ENUMS ///

    /**
     * @notice The process status defines the state of a process.
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
        MERKLE_TREE_OFFCHAIN_STATIC_V1,
        MERKLE_TREE_OFFCHAIN_DYNAMIC_V1,
        MERKLE_TREE_ONCHAIN_V1,
        CSP_EDDSA_BN254_V1
    }

    /// STRUCTS ///

    /**
     * @notice The ballot mode define the parameters of the vote.
     * @param costFromWeight If weighted census, the ballot weight is used as maxValueSum.
     * @param uniqueValues Choices cannot appear twice or more.
     * @param numFields The maximum number of fields per ballot.
     * @param costExponent The exponent that will be used to compute the "cost" of the field values.
     * @param maxValue The maximum value for all fields.
     * @param minValue The minimum value for all fields.
     * @param maxValueSum Maximum limit on the total sum of all ballot fields' values. 0 => Not applicable.
     * @param minValueSum Minimum limit on the total sum of all ballot fields' values. 0 => Not applicable.
     */
    struct BallotMode {
        bool costFromWeight;
        bool uniqueValues;
        uint8 numFields;
        uint8 costExponent;
        uint256 maxValue;
        uint256 minValue;
        uint256 maxValueSum;
        uint256 minValueSum;
    }

    /**
     * @notice The census defines the parameters of the census.
     * @param censusOrigin The origin of the census.
     * @param censusRoot The root of the census.
     * @param censusURI The URI of the census.
     */
    struct Census {
        CensusOrigin censusOrigin;
        bytes32 censusRoot;
        string censusURI;
    }

    /**
     * @notice The process ID is a unique identifier for a process.
     * @param organizationId The organizationId of the process.
     * @param chainID The ID of the chain.
     * @param nonce The nonce of the process.
     */
    struct ProcessId {
        address organizationId;
        uint32 chainID;
        uint64 nonce;
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
     * @param organizationId The organizationId of the process.
     * @param encryptionKey The encryption key of the process.
     * @param latestStateRoot The latest state root of the process.
     * @param result The result of the process.
     * @param startTime The start time of the process.
     * @param duration The duration of the process.
     * @param maxVoters The maximum number of voters allowed.
     * @param votersCount The total number of voters that participated.
     * @param overwrittenVotesCount The number of times votes were overwritten in the state.
     * @param creationBlock The block number when the process was created.
     * @param batchNumber The batch number of the process that increments with each state transition.
     * @param metadataURI The URI of the metadata.
     * @param ballotMode The ballot mode.
     * @param census The census of the process.
     */
    struct Process {
        ProcessStatus status;
        address organizationId;
        EncryptionKey encryptionKey;
        uint256 latestStateRoot;
        uint256[] result;
        uint256 startTime;
        uint256 duration;
        uint256 maxVoters;
        uint256 votersCount;
        uint256 overwrittenVotesCount;
        uint256 creationBlock;
        uint256 batchNumber;
        string metadataURI;
        BallotMode ballotMode;
        Census census;
    }

    /// GETTERS ///

    /**
     * @notice Returns the process data.
     * @param processId The ID of the process.
     * @return process The process struct.
     */
    function getProcess(bytes32 processId) external view returns (Process memory process);

    /**
     * @notice Returns the next process ID.
     * @return The next process ID.
     * @param organizationId The ID of the organization.
     */
    function getNextProcessId(address organizationId) external view returns (bytes32);

    /**
     * @notice Returns the hash of the state transition ZK verifier proving key.
     * @return The hash of the state transition ZK verifier proving key.
     */
    function getSTVerifierVKeyHash() external view returns (bytes32);

    /**
     * @notice Returns the hash of the results ZK verifier proving key.
     * @return The hash of the results ZK verifier proving key.
     */
    function getRVerifierVKeyHash() external view returns (bytes32);

    /**
     * @notice Returns the end time of a process.
     * @param processId The ID of the process.
     * @return The end time of the process.
     */
    function getProcessEndTime(bytes32 processId) external view returns (uint256);

    /// SETTERS ///

    /**
     * @notice Creates a new process.
     * @param status The initial status of the process.
     * @param startTime The start time of the process.
     * @param duration The duration of the process.
     * @param maxVoters The maximum number of voters allowed.
     * @param ballotMode The ballot mode of the process.
     * @param census The census of the process.
     * @param metadata The URI of the metadata.
     * @param encryptionKey The public key used for vote encryption.
     * @param initStateRoot The initial state root.
     */
    function newProcess(
        ProcessStatus status,
        uint256 startTime,
        uint256 duration,
        uint256 maxVoters,
        BallotMode calldata ballotMode,
        Census calldata census,
        string calldata metadata,
        EncryptionKey calldata encryptionKey,
        uint256 initStateRoot
    ) external returns (bytes32);

    /**
     * @notice Sets the status of a process.
     * @param processId The ID of the process.
     * @param newStatus The new status of the process.
     */
    function setProcessStatus(bytes32 processId, ProcessStatus newStatus) external;

    /**
     * @notice Sets the census of a process.
     * @param processId The ID of the process.
     * @param census The census of the process.
     */
    function setProcessCensus(bytes32 processId, Census calldata census) external;

    /**
     * @notice Sets the duration of a process.
     * @param processId The ID of the process.
     * @param duration The new duration of the process.
     */
    function setProcessDuration(bytes32 processId, uint256 duration) external;

    /**
     * @notice Sets the maximum number of voters allowed in a process.
     * @param processId The ID of the process.
     * @param maxVoters The new maximum number of voters.
     */
    function setProcessMaxVoters(bytes32 processId, uint256 maxVoters) external;

    /**
     * @notice Sets the results of a process.
     * @param processId The ID of the process.
     * @param proof The proof for validating the process results.
     * @param input The public inputs data for the results.
     */
    function setProcessResults(bytes32 processId, bytes calldata proof, bytes calldata input) external;

    /**
     * @notice Submits a process state transition.
     * @param processId The ID of the process.
     * @param proof The proof for validating the process state transition.
     * @param input The public inputs data for the state transition.
     */
    function submitStateTransition(bytes32 processId, bytes calldata proof, bytes calldata input) external;
}
