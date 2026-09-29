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
     * @notice Emitted when the organizer replaces the census of an origin-2 process.
     * @param processId The ID of the process.
     * @param censusRoot The new census root (big-endian integer).
     * @param censusURI The URI of the new census.
     */
    event CensusUpdated(bytes31 indexed processId, bytes32 censusRoot, string censusURI);
    /**
     * @notice Emitted with a process's metadata at creation and on every setProcessMetadata,
     *         so the log alone holds the full metadata history.
     * @param processId The ID of the process.
     * @param metadataURI The URI of the metadata document.
     * @param metadataHash SHA-256 of the exact bytes served at metadataURI (no JSON
     *        canonicalisation).
     */
    event ProcessMetadataUpdated(bytes31 indexed processId, string metadataURI, bytes32 metadataHash);
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
    /**
     * @notice Emitted when the organizer changes the grace window of a process.
     * @param processId The ID of the process.
     * @param grace The new idle window past the end, in seconds.
     */
    event ProcessGraceChanged(bytes31 indexed processId, uint32 grace);
    /**
     * @notice Emitted when the final accumulator of a DKG-mode process is bound to its
     *         state root and its active ciphertexts are submitted for threshold decryption.
     * @param processId The ID of the process.
     * @param epochId The DKG epoch of the process's application.
     * @param aid The DKG application id.
     * @param firstIndex The DKG index of the first submitted ciphertext (0 when none).
     * @param count The number of submitted ciphertexts (identity fields are skipped).
     */
    event ResultsDecryptionRequested(
        bytes31 indexed processId, bytes12 epochId, bytes32 aid, uint16 firstIndex, uint8 count
    );

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
     * @notice InvalidMetadata error is emitted when the metadata URI is empty or its hash is zero.
     */
    error InvalidMetadata();
    /**
     * @notice InvalidCensusOrigin error is emitted when the census origin is invalid.
     */
    error InvalidCensusOrigin();
    /**
     * @notice InvalidCensusConfig error is a more generic error emitted when a census configuration is invalid.
     */
    error InvalidCensusConfig();
    /**
     * @notice InvalidCensusAddress error is emitted when an on-chain census address has no code
     *         or does not answer getCensusRoot(), or when another census origin sets an address.
     */
    error InvalidCensusAddress();
    /**
     * @notice CensusNotUpdatable error is emitted when setProcessCensus targets a process whose
     *         census origin is not MERKLE_TREE_OFFCHAIN_DYNAMIC_V1.
     */
    error CensusNotUpdatable();
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
     * @notice Thrown when the blob arrays do not all have n_blobs entries, or the transaction
     *         carries a blob beyond the first n_blobs.
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
    /**
     * @notice Thrown when newProcess carries DKG params that do not match its key mode
     *         (non-zero fields in SEQUENCER mode).
     */
    error InvalidDKGParams();
    /**
     * @notice Thrown when a DKG key mode is used on a registry deployed without a DKG manager.
     */
    error DKGDisabled();
    /**
     * @notice Thrown when the call is not valid for the process's key mode.
     */
    error InvalidKeyMode();
    /**
     * @notice Thrown when the accumulator's SMT inclusion proof does not verify under the
     *         process's latest state root.
     */
    error InvalidInclusionProof();
    /**
     * @notice Thrown when an accumulator coordinate is out of range or a field has an
     *         identity C1 with a non-identity C2.
     */
    error InvalidAccumulator();
    /**
     * @notice Thrown when finalizeResultsFromDKG runs before the request or before every
     *         submitted ciphertext has a completed combine.
     */
    error ResultsNotReady();
    /**
     * @notice Thrown when requestResultsDecryption runs twice for a process.
     */
    error ResultsAlreadyRequested();
    /**
     * @notice Thrown when a grace value is outside [graceFloor, graceCeil], or when the
     *         constructor's grace bounds are not 0 < floor <= default <= ceil <= maxTotal.
     */
    error InvalidGrace();
    /**
     * @notice Thrown when results are set or requested before the grace window has closed.
     */
    error GraceOpen();
    /**
     * @notice Thrown when a state transition adds no vote (no new voter and no overwrite).
     */
    error EmptyTransition();

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

    /**
     * @notice Returns when the grace window of a process closes: transitions settle while
     *         block.timestamp is below it and the results calls open at it. With
     *         end = startTime + duration it is
     *         min(end + graceMaxTotal, max(end, lastVoteAt) + grace). Every settled
     *         transition moves lastVoteAt, so the window extends while votes keep landing,
     *         up to the cap; once it passes nothing can settle and it never reopens.
     * @param processId The ID of the process.
     * @return The grace end, a unix timestamp in seconds.
     */
    function getProcessGraceEnd(bytes31 processId) external view returns (uint256);

    /**
     * @notice The grace window a new process starts with, in seconds.
     */
    function defaultGrace() external view returns (uint32);

    /**
     * @notice The shortest grace window setProcessGrace accepts, in seconds.
     */
    function graceFloor() external view returns (uint32);

    /**
     * @notice The longest grace window setProcessGrace accepts, in seconds.
     */
    function graceCeil() external view returns (uint32);

    /**
     * @notice The cap on the grace window past the end time, in seconds.
     */
    function graceMaxTotal() external view returns (uint32);

    /**
     * @notice The minimum notice, in seconds, for shortening a process with setProcessDuration.
     */
    function noticeMin() external view returns (uint32);

    /**
     * @notice The DavinciDKGAdapter created at deploy, or address(0) when DKG modes are
     *         disabled. Clients read registrationEpoch() from it before a DKG_LOCKED
     *         newProcess.
     */
    function dkgAdapter() external view returns (address);

    /**
     * @notice The DKG application id a process registers under. Reverts DKGDisabled when
     *         no adapter is configured.
     * @param processId The ID of the process (existing or upcoming, see getNextProcessId).
     */
    function aidFor(bytes31 processId) external view returns (bytes32);

    /// SETTERS ///

    /**
     * @notice Creates a new process.
     * @param status The initial status of the process.
     * @param startTime The start time of the process.
     * @param duration The duration of the process.
     * @param maxVoters The maximum number of voters allowed.
     * @param ballotMode The ballot mode of the process.
     * @param census The census of the process.
     * @param metadataURI The URI of the metadata document, non-empty.
     * @param metadataHash SHA-256 of the exact bytes served at metadataURI (no JSON
     *        canonicalisation), non-zero. Emitted in ProcessMetadataUpdated.
     * @param encryptionKey The public key used for vote encryption. Must be (0, 0) in the
     *        DKG key modes, where the registry takes the key from the DKG committee.
     * @param dkg The key mode and DKG registration arguments (all zero for SEQUENCER).
     */
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
    ) external returns (bytes31);

    /**
     * @notice Sets the status of a process.
     * @param processId The ID of the process.
     * @param newStatus The new status of the process.
     */
    function setProcessStatus(bytes31 processId, DAVINCITypes.ProcessStatus newStatus) external;

    /**
     * @notice Replaces the census root and URI of a MERKLE_TREE_OFFCHAIN_DYNAMIC_V1 process.
     *         Only the organizer, while READY or PAUSED and before the end time. Batches
     *         proven against the previous root no longer settle.
     * @param processId The ID of the process.
     * @param census The new census: the process's origin, a non-zero root, a non-empty URI and
     *        no contract address.
     */
    function setProcessCensus(bytes31 processId, DAVINCITypes.Census calldata census) external;

    /**
     * @notice Replaces the metadata URI and hash of a process. Only the organizer, while
     *         READY or PAUSED and before the end time, so what a ballot field means is
     *         frozen once voting closes.
     * @param processId The ID of the process.
     * @param metadataURI The URI of the new metadata document, non-empty.
     * @param metadataHash SHA-256 of the exact bytes served at metadataURI (no JSON
     *        canonicalisation), non-zero.
     */
    function setProcessMetadata(bytes31 processId, string calldata metadataURI, bytes32 metadataHash) external;

    /**
     * @notice Sets the duration of a process. Only the organizer, while READY or PAUSED and
     *         before its current end: past it the tally may already be public, so the election
     *         cannot be reopened. The end can move later freely, or earlier with notice: the
     *         new end must be at least noticeMin seconds away. The grace window follows the
     *         new end.
     * @param processId The ID of the process.
     * @param duration The new duration of the process, non-zero.
     */
    function setProcessDuration(bytes31 processId, uint256 duration) external;

    /**
     * @notice Sets the maximum number of voters allowed in a process.
     * @param processId The ID of the process.
     * @param maxVoters The new maximum number of voters.
     */
    function setProcessMaxVoters(bytes31 processId, uint256 maxVoters) external;

    /**
     * @notice Sets the grace window of a process: how long past the end, or past the last
     *         settled transition, transitions keep settling. Only the organizer, while
     *         READY or PAUSED and before the end time.
     * @param processId The ID of the process.
     * @param grace The idle window in seconds, within [graceFloor, graceCeil].
     */
    function setProcessGrace(bytes31 processId, uint32 grace) external;

    /**
     * @notice Sets the results of a process from a results-guest proof over its final state root.
     *         Only once the grace window has closed (getProcessGraceEnd), so the root is final.
     * @param processId The ID of the process.
     * @param publicValues The 512-byte ZisK public values.
     * @param proofBytes The PLONK proof, abi-encoded uint256[24].
     */
    function setProcessResults(bytes31 processId, bytes calldata publicValues, bytes calldata proofBytes) external;

    /**
     * @notice Binds the final accumulator of an ended DKG-mode process to its latest state
     *         root (SMT inclusion of key 0x04) and submits every active field's ciphertext
     *         to the DKG committee for threshold decryption. Permissionless, at most once
     *         per process, and only once the grace window has closed (getProcessGraceEnd).
     *         With no active field (all identity) the results are finalized to zero
     *         immediately; otherwise the process is moved to ENDED, so the tally cannot be
     *         read off the DKG and then canceled.
     * @dev DKG liveness is election liveness: once requested there is no un-request and
     *      no fallback to setProcessResults, so if more than n - t committee members of
     *      the registration epoch are gone the combines never complete and the results
     *      are lost.
     * @param processId The ID of the process.
     * @param accumulator The 64 BE coordinates of the results accumulator (16 ElGamal
     *        ciphertexts, circomlib form), whose sha256 is the value of state leaf 0x04.
     * @param siblings The SMT inclusion siblings, root to leaf, zero-padded.
     */
    function requestResultsDecryption(bytes31 processId, uint256[64] calldata accumulator, bytes32[] calldata siblings)
        external;

    /**
     * @notice Reads the DKG's combined plaintexts once every submitted ciphertext is
     *         decrypted, stores the results and sets the process to RESULTS.
     *         Permissionless; reverts ResultsNotReady until the combines are complete.
     * @dev A DKG committee that never completes the combines strands the process short
     *      of RESULTS forever (see requestResultsDecryption); that is the trust model.
     * @param processId The ID of the process.
     */
    function finalizeResultsFromDKG(bytes31 processId) external;

    /**
     * @notice Publishes the organizer secret of a DKG_LOCKED process, after which the DKG
     *         committee can combine decryptions. Permissionless: the DKG checks
     *         sk·G == PK_org, so only the real secret is accepted.
     * @param processId The ID of the process.
     * @param sk The organizer secret key (scalar of the reduced-form organizer key).
     */
    function revealProcessKey(bytes31 processId, uint256 sk) external;

    /**
     * @notice Settles a state transition proven by the vote-batch guest. Must be sent as a blob
     *         transaction carrying the transition's blobs in order. Accepted while the process
     *         is READY or ENDED, from startTime until getProcessGraceEnd, and only for a batch
     *         with at least one vote. Sets lastVoteAt, which extends the grace window.
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
