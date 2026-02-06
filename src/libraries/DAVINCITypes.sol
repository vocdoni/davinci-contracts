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
        CENSUS_UNKNOWN,
        MERKLE_TREE_OFFCHAIN_STATIC_V1,
        MERKLE_TREE_OFFCHAIN_DYNAMIC_V1,
        MERKLE_TREE_ONCHAIN_DYNAMIC_V1,
        CSP_EDDSA_BABYJUBJUB_V1
    }

    /**
     * @notice The ballot mode define the parameters of the vote.
     * @param costFromWeight If weighted census, the ballot weight is used as maxValueSum.
     * @param uniqueValues Choices cannot appear twice or more.
     * @param numFields The maximum number of fields per ballot.
     * @param groupSize Used for multiquestion patterns.
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
     * @param censusRoot The root of the census. CSP -> A PublicKey, MerkleTree OffchainStatic, OffchainDynamic -> A Hash, MerkleTree Onchain -> A Contract address
     * @param contractAddress An EVM contract address (optional). Ideally this contract returns census information and/or data.
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
}
