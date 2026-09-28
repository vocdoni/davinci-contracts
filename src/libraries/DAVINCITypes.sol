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
     * @dev The zkVM ProcessRegistry accepts origins 1-4. Origins 1-3 are lean-IMT censuses: a fixed
     *      root, a root the organizer replaces with setProcessCensus, and the root history of an
     *      ICensusValidator contract. CSP_EDDSA_BABYJUBJUB_V1 is verified by the davinci-zkvm guest
     *      as an ECDSA/secp256k1 CSP (the value 4 is kept; the name is historical).
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
     * @param uniqueValues Choices cannot appear twice or more.
     * @param numFields The maximum number of fields per ballot (1..16).
     * @param groupSize Used for multiquestion patterns.
     * @param costExponent The exponent that will be used to compute the "cost" of the field values.
     * @param maxValue The maximum value for all fields.
     * @param minValue The minimum value for all fields.
     * @param maxValueSum Maximum limit on the total sum of all ballot fields' values. 0 => Limit will be each voter weight.
     * @param minValueSum Minimum limit on the total sum of all ballot fields' values. 0 => No lower limit
     */
    struct BallotMode {
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
     * @param censusRoot The root of the census as a big-endian integer: the lean-IMT root for a Merkle census,
     *        bytes32(uint256(uint160(cspAddress))) for a CSP census. For MERKLE_TREE_ONCHAIN_DYNAMIC_V1 the
     *        registry stores the contract's root at creation, for information only.
     * @param contractAddress The ICensusValidator census contract for MERKLE_TREE_ONCHAIN_DYNAMIC_V1, which must
     *        have code and answer getCensusRoot(). Must be address(0) for every other origin.
     * @param censusURI The URI of the census.
     * @param onchainAllowAnyValidRoot Must be false: the zkVM registry only settles on-chain census roots held at or
     *        after process creation.
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
     * @notice Where a process's encryption key comes from.
     *         SEQUENCER: the organizer/sequencer supplies it (results via the zkVM results
     *         guest). DKG_AUTOMATIC: a davinci-dkg committee pool key; the committee alone
     *         decrypts the final accumulator. DKG_LOCKED: pool key plus an organizer key;
     *         results wait for revealProcessKey.
     */
    enum KeyMode {
        SEQUENCER,
        DKG_AUTOMATIC,
        DKG_LOCKED
    }

    /**
     * @notice DKG arguments of newProcess. All fields must be zero in SEQUENCER mode.
     *         For DKG_LOCKED, epochId picks the epoch (the PoP binds it) and the organizer
     *         key plus Schnorr PoP are in the DKG's reduced (a = -1) form, exactly as
     *         DKGAppManager.registerApplication takes them. For DKG_AUTOMATIC only mode
     *         is read; the adapter picks the epoch.
     */
    struct DKGParams {
        KeyMode mode;
        bytes12 epochId;
        uint256 orgPKx;
        uint256 orgPKy;
        uint256 popAx;
        uint256 popAy;
        uint256 popZ;
    }

    /**
     * @notice The process defines the parameters of the process.
     * @param status The status of the process.
     * @param organizationId The organizationId of the process.
     * @param encryptionKey The encryption key of the process.
     * @param latestStateRoot The latest state root of the process: the raw SHA-256 digest of the arbo root.
     * @param result The result of the process.
     * @param startTime The start time of the process.
     * @param duration The duration of the process.
     * @param maxVoters The maximum number of voters allowed.
     * @param votersCount The number of distinct ballot slots written (votes minus overwrites).
     * @param overwrittenVotesCount The number of times votes were overwritten in the state.
     * @param creationBlock The block number when the process was created.
     * @param batchNumber The batch number of the process that increments with each state transition.
     * @param metadataURI The URI of the metadata.
     * @param ballotMode The ballot mode.
     * @param census The census of the process.
     * @param keyMode Where the encryption key comes from.
     * @param dkgEpochId The DKG epoch the process registered against (DKG modes).
     * @param dkgFirstIndex Index of the first ciphertext submitted by requestResultsDecryption.
     * @param dkgCount Number of ciphertexts submitted; > 0 also means results were requested.
     * @param dkgZeroSkipped Bitmask of fields recorded as 0 without a DKG submission
     *        (identity ciphertexts), bit i = field i.
     * @param dkgResultsRequested Whether requestResultsDecryption ran (covers the all-identity
     *        case where dkgCount stays 0).
     * @param dkgAid The DKG application id (keccak(chainid, registry, pid) mod Q).
     */
    struct Process {
        ProcessStatus status;
        address organizationId;
        EncryptionKey encryptionKey;
        bytes32 latestStateRoot;
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
        KeyMode keyMode;
        bytes12 dkgEpochId;
        uint16 dkgFirstIndex;
        uint8 dkgCount;
        uint16 dkgZeroSkipped;
        bool dkgResultsRequested;
        bytes32 dkgAid;
    }
}
