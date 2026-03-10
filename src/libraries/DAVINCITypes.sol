// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/**
 * @title DAVINCITypes
 * @notice Shared DAVINCI enums and structs.
 */
library DAVINCITypes {
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
        MERKLE_TREE_OFFCHAIN,
        MERKLE_TREE_ONCHAIN,
        CSP_EDDSA_BABYJUBJUB_V1
    }

    /**
     * @notice The ballot mode define the parameters of the vote.
     * @param uniqueValues Choices cannot appear twice or more.
     * @param costFromWeight If weighted census, the ballot weight is used as maxValueSum.
     * @param numFields The maximum number of fields per ballot.
     * @param groupSize Used for multiquestion patterns.
     * @param costExponent The exponent that will be used to compute the "cost" of the field values.
     * @param maxValue The maximum value for all fields.
     * @param minValue The minimum value for all fields.
     * @param maxValueSum Maximum limit on the total sum of all ballot fields' values. 0 => Not applicable.
     * @param minValueSum Minimum limit on the total sum of all ballot fields' values. 0 => Not applicable.
     */
    struct BallotMode {
        bool uniqueValues;
        bool costFromWeight;
        uint8 numFields;
        uint8 groupSize;
        uint8 costExponent;
        uint256 maxValue;
        uint256 minValue;
        uint256 maxValueSum;
        uint256 minValueSum;
    }

    /**
     * @notice The census defines the parameters of the census.
     * @param censusOrigin The origin of the census.
     * @param censusRoot The committed census value. CSP -> A public key, MerkleTree -> A hash.
     * Dynamic onchain censuses validate roots through contractAddress on each state transition.
     * @param contractAddress An EVM contract address (optional). Required for onchain censuses to validate accepted roots.
     * @param censusURI The URI of the census.
     * @param onchainAllowAnyValidRoot Used for onchain censuses. If true allows to skip the census startBlock check in the state transition function.
     */
    struct Census {
        CensusOrigin censusOrigin;
        bytes32 censusRoot;
        address contractAddress;
        string censusURI;
        bool onchainAllowAnyValidRoot;
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
     * @notice EncryptionKey of a process
     * @param x value of the X coordinate on the curve
     * @param y value of the Y coordinate on the curve
     */
    struct EncryptionKey {
        uint256 x;
        uint256 y;
    }

    /**
     * @notice ParamsMod can be used to restrict or allow certain changes on process data.
     * @dev The ballot mode cannot be modified.
     * @param census If true, allows changes on the Census.
     * @param status If true, allows non-contract automated changes.
     * @param duration If true, allows changes on the duration.
     * @param metadata If true, allows changes on the metadata.
     * @param maxVoters If true, allows to change the maximum number of voters.
     */
    struct ParamsMod {
        bool census;
        bool status;
        bool duration;
        bool metadata;
        bool maxVoters;
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
     * @param paramsMod The paramsMod of the process.
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
        ParamsMod paramsMod;
    }

    /**
     * @notice StateTransitionBatchProofInputs wraps all the inputs required for a state transition operation
     * @param rootHashBefore The old root hash of the state.
     * @param rootHashAfter The new root hash of the state.
     * @param votersCount The number of voters included in the specific state transition.
     * @param overwrittenVotesCount The number of vote overwrittes included in the specific state transition.
     * @param censusRoot The census root used in the state transition.
     * @param blobCommitmentLimb0 The first limb of the blob commitment.
     * @param blobCommitmentLimb1 The second limb of the blob commitment.
     * @param blobCommitmentLimb2 The third limb of the blob commitment
     * @dev The total number of voters and overwrites are stored in the Process struct.
     * @dev The blob commitment limbs are only used if the ProcessRegistry contract have the blobs feature activated.
     */
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
}
