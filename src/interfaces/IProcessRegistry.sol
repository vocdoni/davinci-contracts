// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {DAVINCITypes} from "../libraries/DAVINCITypes.sol";

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
    event ProcessCreated(bytes31 indexed processId, address indexed creator);
    /*
     * @notice Emitted when the duration of a process is modified.
     * @param processId The ID of the process.
     * @param duration The new duration of the process.
     */
    event ProcessDurationChanged(bytes31 indexed processId, uint256 duration);
    /*
     * @notice Emitted when a state transition is settled.
     * @param processId The ID of the process.
     * @param sender The address of the sender.
     * @param oldStateRoot The state root before the transition (raw digest).
     * @param newStateRoot The state root after the transition (raw digest).
     * @param newVotersCount The process votersCount after the transition.
     * @param newOverwrittenVotesCount The process overwrittenVotesCount after the transition.
     * @param nBlobs The number of blobs the transition was published in.
     */
    event ProcessStateTransitioned(
        bytes31 indexed processId,
        address indexed sender,
        bytes32 oldStateRoot,
        bytes32 newStateRoot,
        uint256 newVotersCount,
        uint256 newOverwrittenVotesCount,
        uint256 nBlobs
    );

    /*
     * @notice Emitted when the results of a process are set.
     * @param processId The ID of the process.
     * @param sender The address of the sender.
     * @param result The result of the process.
     */
    event ProcessResultsSet(bytes31 indexed processId, address indexed sender, uint256[] result);
    /**
     * @notice Emitted when a process status is modified
     * @param processId The ID of the process
     * @param oldStatus The previous status of the process
     * @param newStatus The new status of the process
     */
    event ProcessStatusChanged(
        bytes31 indexed processId, DAVINCITypes.ProcessStatus oldStatus, DAVINCITypes.ProcessStatus newStatus
    );
    /**
     * @notice Emitted when the max voters of a process is modified
     * @param processId The ID of the process
     * @param maxVoters The new max voters of the process
     */
    event ProcessMaxVotersChanged(bytes31 indexed processId, uint256 maxVoters);

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
     * InvalidBlockNumber error is emitted when a block number is invalid.
     */
    error InvalidBlockNumber();
    /**
     * @notice InvalidDuration error is emitted when the duration of the process is invalid.
     */
    error InvalidDuration();
    /**
     * @notice InvalidMaxVoters error is emitted when the maximum number of voters is invalid.
     */
    error InvalidMaxVoters();
    /**
     * @notice MaxPossibleResultCapExceeded error is emitted when the process can exceed the decryption cap.
     */
    error MaxPossibleResultCapExceeded();
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
     * @notice InvalidGroupSize error is emitted when the grup size value is invalid.
     */
    error InvalidGroupSize();
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
     * @notice InvalidCensusConfig error is a more generic error emitted when a census configuration is invalid.
     */
    error InvalidCensusConfig();
    /**
     * @notice InvalidEncryptionKey error is emitted when the key is not a canonical BabyJubJub point
     *         (circomlib twisted Edwards) with x != 0.
     */
    error InvalidEncryptionKey();
    /**
     * @notice Ballot mode fields that do not fit their packed bit width.
     */
    error BallotModeMaxValueTooLarge();
    error BallotModeMinValueTooLarge();
    error BallotModeMaxValueSumTooLarge();
    error BallotModeMinValueSumTooLarge();
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
     * @notice Thrown when a constructor address or key is zero.
     */
    error InvalidVerifierConfig();
    /**
     * @notice Thrown when publicValues is not 512 bytes.
     */
    error InvalidPublicValues();
    /**
     * @notice Thrown when the guest reports a failed check (ok != 1 or fail_mask != 0).
     */
    error CircuitFailed();
    /**
     * @notice Thrown when the proof's occupied_before differs from the process votersCount.
     */
    error InvalidOccupiedBefore();
    /**
     * @notice Thrown when the proof publishes no blobs.
     */
    error NoBlobs();
    /**
     * @notice Thrown when the blob arrays do not all have n_blobs entries.
     */
    error BlobCountMismatch();
    /**
     * @notice Thrown when a blob commitment is not 48 bytes.
     */
    error InvalidBlobCommitmentLength();
    /**
     * @notice Thrown when a KZG opening proof is not 48 bytes.
     */
    error InvalidKZGProofLength();
    /**
     * @notice Thrown when sha256(commitment_0 ‖ y_0 ‖ ...) differs from the digest the guest published.
     */
    error InvalidBlobsDigest();
    /**
     * @notice Thrown when the transaction carries no blob at the given index.
     */
    error MissingBlob(uint256 index);
    /**
     * @notice Thrown when the point-evaluation precompile rejects the opening of blob index.
     */
    error InvalidBlobOpening(uint256 index);
    /**
     * @notice Thrown when the sender is not authorized to perform the action.
     */
    error Unauthorized();

    /// GETTERS ///

    /**
     * @notice Returns the process data.
     * @param processId The ID of the process.
     * @return process The process struct.
     */
    function getProcess(bytes31 processId) external view returns (DAVINCITypes.Process memory process);

    /**
     * @notice Returns the next process ID.
     * @return The next process ID.
     * @param organizationId The ID of the organization.
     */
    function getNextProcessId(address organizationId) external view returns (bytes31);

    /**
     * @notice Returns the program vk of the vote-batch guest that state transitions are proven with.
     */
    function getSTVerifierVKeyHash() external view returns (bytes32);

    /**
     * @notice Returns the program vk of the results guest that results are proven with.
     */
    function getRVerifierVKeyHash() external view returns (bytes32);

    /**
     * @notice Returns the genesis state root a process with this configuration starts from.
     * @param processId The ID of the process.
     * @param ballotMode The ballot mode of the process.
     * @param encryptionKey The encryption key (twisted Edwards coordinates).
     * @param censusOrigin The census origin.
     * @return The raw SHA-256 digest of the arbo root.
     */
    function genesisRoot(
        bytes31 processId,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.EncryptionKey calldata encryptionKey,
        DAVINCITypes.CensusOrigin censusOrigin
    ) external view returns (bytes32);

    /**
     * @notice Returns the end time of a process.
     * @param processId The ID of the process.
     * @return The end time of the process.
     */
    function getProcessEndTime(bytes31 processId) external view returns (uint256);

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
     */
    function newProcess(
        DAVINCITypes.ProcessStatus status,
        uint256 startTime,
        uint256 duration,
        uint256 maxVoters,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.Census calldata census,
        string calldata metadata,
        DAVINCITypes.EncryptionKey calldata encryptionKey
    ) external returns (bytes31);

    /**
     * @notice Sets the status of a process.
     * @param processId The ID of the process.
     * @param newStatus The new status of the process.
     */
    function setProcessStatus(bytes31 processId, DAVINCITypes.ProcessStatus newStatus) external;

    /**
     * @notice Sets the duration of a process.
     * @param processId The ID of the process.
     * @param duration The new duration of the process.
     */
    function setProcessDuration(bytes31 processId, uint256 duration) external;

    /**
     * @notice Sets the maximum number of voters allowed in a process.
     * @param processId The ID of the process.
     * @param maxVoters The new maximum number of voters.
     */
    function setProcessMaxVoters(bytes31 processId, uint256 maxVoters) external;

    /**
     * @notice Sets the results of a process from a results-guest proof over its final state root.
     * @param processId The ID of the process.
     * @param publicValues The 512-byte ZisK public values.
     * @param proofBytes The PLONK proof, abi-encoded uint256[24].
     */
    function setProcessResults(bytes31 processId, bytes calldata publicValues, bytes calldata proofBytes) external;

    /**
     * @notice Settles a state transition proven by the vote-batch guest. Must be sent as a blob
     *         transaction carrying the transition's blobs in order.
     * @param processId The ID of the process.
     * @param publicValues The 512-byte ZisK public values.
     * @param proofBytes The PLONK proof, abi-encoded uint256[24].
     * @param commitments The 48-byte KZG commitment of each blob.
     * @param ys The big-endian evaluation of each blob at its bound point.
     * @param kzgProofs The 48-byte opening proof of each blob at its bound point.
     */
    function submitStateTransition(
        bytes31 processId,
        bytes calldata publicValues,
        bytes calldata proofBytes,
        bytes[] calldata commitments,
        bytes32[] calldata ys,
        bytes[] calldata kzgProofs
    ) external;
}
