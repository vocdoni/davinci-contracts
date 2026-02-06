// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {PoseidonT3} from "poseidon-solidity/PoseidonT3.sol";
import {PoseidonT4} from "poseidon-solidity/PoseidonT4.sol";
import {DAVINCITypes} from "./DAVINCITypes.sol";

/**
 * @title StateRootLib
 * @author Vocdoni Association
 * @notice Computes and verifies the initial state root for a DAVINCI process.
 */
library StateRootLib {
    uint256 private constant _MAX_48 = (1 << 48) - 1;
    uint256 private constant _MAX_63 = (1 << 63) - 1;

    /// @dev Field key for the process identifier leaf.
    uint8 private constant KEY_PROCESS_ID = 0;
    /// @dev Field key for the packed ballot mode leaf.
    uint8 private constant KEY_BALLOT_MODE = 2;
    /// @dev Field key for the encryption key hash leaf.
    uint8 private constant KEY_ENCRYPTION_KEY = 3;
    /// @dev Field key for the results-add accumulator leaf.
    uint8 private constant KEY_RESULTS_ADD = 4;
    /// @dev Field key for the results-sub accumulator leaf.
    uint8 private constant KEY_RESULTS_SUB = 5;
    /// @dev Field key for the census origin leaf.
    uint8 private constant KEY_CENSUS_ORIGIN = 6;
    /// @dev Domain separator used for all leaf hashes in this tree.
    uint8 private constant LEAF_DOMAIN = 1;

    /// @dev Precomputed leaf hash for the results-add accumulator branch.
    uint256 private constant LEAF_RESULTS_ADD = 0x1f72c52b6e5dedca4f99ecfa24f2776732431e8d544e14c6f78f5042727c4657;
    /// @dev Precomputed leaf hash for the results-sub accumulator branch.
    uint256 private constant LEAF_RESULTS_SUB = 0x2b853c511fba705a6030f80ce83d6ee8cbf4a1273724368728c11682eae4c51a;

    error BallotModeGroupSizeExceedsNumFields();
    error BallotModeMaxValueTooLarge();
    error BallotModeMinValueTooLarge();
    error BallotModeMaxValueSumTooLarge();
    error BallotModeMinValueSumTooLarge();

    /**
     * @notice Computes the initial state root for a process.
     * @param processId Process identifier.
     * @param censusOrigin Census origin value used in the state root.
     * @param ballotMode The ballot mode.
     * @param encryptionKey The encryption key coordinates.
     * @return stateRoot Root of the initial process state tree.
     */
    function computeStateRoot(
        bytes32 processId,
        DAVINCITypes.CensusOrigin censusOrigin,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.EncryptionKey calldata encryptionKey
    ) external pure returns (uint256) {
        return _computeStateRoot(processId, censusOrigin, ballotMode, encryptionKey);
    }

    /**
     * @notice Verifies that a given initial process state root.
     * @param expectedRoot Root expected by the caller.
     * @param processId Process identifier.
     * @param censusOrigin Census origin value used in the state root.
     * @param ballotMode The ballot mode.
     * @param encryptionKey The encryption key coordinates.
     * @return isValid True if `expectedRoot` matches the computed root.
     */
    function verifyStateRoot(
        uint256 expectedRoot,
        bytes32 processId,
        DAVINCITypes.CensusOrigin censusOrigin,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.EncryptionKey calldata encryptionKey
    ) external pure returns (bool) {
        return expectedRoot == _computeStateRoot(processId, censusOrigin, ballotMode, encryptionKey);
    }

    /**
     * @dev Internal state-root computation shared by all public entrypoints.
     */
    function _computeStateRoot(
        bytes32 processId,
        DAVINCITypes.CensusOrigin censusOrigin,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.EncryptionKey calldata encryptionKey
    ) private pure returns (uint256) {
        uint256 packedBallotMode = _packBallotMode(ballotMode);
        uint256 encKeyHash = PoseidonT3.hash([encryptionKey.x, encryptionKey.y]);

        uint256 leafProcess = PoseidonT4.hash([KEY_PROCESS_ID, uint256(processId), LEAF_DOMAIN]);
        uint256 leafBallotMode = PoseidonT4.hash([KEY_BALLOT_MODE, packedBallotMode, LEAF_DOMAIN]);
        uint256 leafEncKey = PoseidonT4.hash([KEY_ENCRYPTION_KEY, encKeyHash, LEAF_DOMAIN]);
        uint256 leafResultsAdd = LEAF_RESULTS_ADD;
        uint256 leafResultsSub = LEAF_RESULTS_SUB;
        uint256 leafCensus = PoseidonT4.hash([KEY_CENSUS_ORIGIN, uint256(censusOrigin), LEAF_DOMAIN]);

        uint256 nodeA0 = PoseidonT3.hash([leafProcess, leafResultsAdd]);
        uint256 nodeA1 = PoseidonT3.hash([leafBallotMode, leafCensus]);
        uint256 nodeA = PoseidonT3.hash([nodeA0, nodeA1]);
        uint256 nodeB = PoseidonT3.hash([leafResultsSub, leafEncKey]);

        return PoseidonT3.hash([nodeA, nodeB]);
    }

    /**
     * @dev Packs the 9-element ballot mode vector into a single field element using fixed bit offsets.
     */
    function _packBallotMode(DAVINCITypes.BallotMode calldata ballotMode) private pure returns (uint256 packed) {
        if (ballotMode.groupSize > ballotMode.numFields) revert BallotModeGroupSizeExceedsNumFields();

        packed = uint256(ballotMode.numFields);
        packed |= uint256(ballotMode.groupSize) << 8;
        packed |= uint256(ballotMode.uniqueValues ? 1 : 0) << 16;
        packed |= uint256(ballotMode.costFromWeight ? 1 : 0) << 17;
        packed |= uint256(ballotMode.costExponent) << 18;

        // check potential bit bleeding
        if (ballotMode.maxValue > _MAX_48) revert BallotModeMaxValueTooLarge();
        if (ballotMode.minValue > _MAX_48) revert BallotModeMinValueTooLarge();
        if (ballotMode.maxValueSum > _MAX_63) revert BallotModeMaxValueSumTooLarge();
        if (ballotMode.minValueSum > _MAX_63) revert BallotModeMinValueSumTooLarge();

        packed |= ballotMode.maxValue << 26;
        packed |= ballotMode.minValue << 74;
        packed |= ballotMode.maxValueSum << 122;
        packed |= ballotMode.minValueSum << 185;
    }
}
