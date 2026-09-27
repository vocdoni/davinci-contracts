// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {DAVINCITypes} from "./DAVINCITypes.sol";
import {Sha256SmtLib} from "./Sha256SmtLib.sol";
import {IProcessRegistry} from "../interfaces/IProcessRegistry.sol";

/**
 * @title GenesisLib
 * @notice The genesis state tree of a davinci-zkvm process: six config leaves in the
 *         SHA-256 state tree (go-sdk chain.NewState). Leaf values are integers; values
 *         that are hashes are the SHA-256 digest read big-endian.
 */
library GenesisLib {
    uint64 internal constant KEY_PROCESS_ID = 0x00;
    uint64 internal constant KEY_BALLOT_MODE = 0x02;
    uint64 internal constant KEY_ENCRYPTION_KEY = 0x03;
    uint64 internal constant KEY_RESULTS = 0x04;
    uint64 internal constant KEY_CENSUS_ORIGIN = 0x06;
    uint64 internal constant KEY_BALLOT_VK = 0x07;

    /// @dev Results leaf at genesis: sha256 of the identity accumulator [0,1,0,1,...]
    ///      (64 BE32 coordinates). Checked in test/Genesis.t.sol.
    uint256 internal constant IDENTITY_ACC_LEAF = 0x1df400d68944aa728f663b60bd3949ede2595ea1ee2dcc95b83a5f36facae278;

    /// @dev BabyJubJub (circomlib twisted Edwards) over the BN254 scalar field.
    uint256 private constant BJJ_P = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    uint256 private constant BJJ_A = 168700;
    uint256 private constant BJJ_D = 168696;

    uint256 private constant MAX_48 = (1 << 48) - 1;
    uint256 private constant MAX_63 = (1 << 63) - 1;

    /// @notice Genesis state root of a process.
    function root(
        bytes31 processId,
        DAVINCITypes.BallotMode calldata ballotMode,
        DAVINCITypes.EncryptionKey calldata encryptionKey,
        DAVINCITypes.CensusOrigin censusOrigin,
        bytes32 ballotVKHash
    ) internal pure returns (bytes32) {
        uint64[] memory keys = new uint64[](6);
        uint256[] memory values = new uint256[](6);
        (keys[0], values[0]) = (KEY_PROCESS_ID, uint256(uint248(processId)));
        (keys[1], values[1]) = (KEY_BALLOT_MODE, packBallotMode(ballotMode));
        (keys[2], values[2]) = (KEY_ENCRYPTION_KEY, encKeyLeaf(encryptionKey));
        (keys[3], values[3]) = (KEY_RESULTS, IDENTITY_ACC_LEAF);
        (keys[4], values[4]) = (KEY_CENSUS_ORIGIN, uint256(censusOrigin));
        (keys[5], values[5]) = (KEY_BALLOT_VK, uint256(ballotVKHash));
        return Sha256SmtLib.root(keys, values);
    }

    /// @notice Packs the ballot mode into one field element (spec.BallotMode.Pack):
    ///         numFields[0:8] groupSize[8:16] uniqueValues[16] costExponent[17:25]
    ///         maxValue[25:73] minValue[73:121] maxValueSum[121:184] minValueSum[184:247].
    function packBallotMode(DAVINCITypes.BallotMode calldata m) internal pure returns (uint256 packed) {
        if (m.groupSize > m.numFields) revert IProcessRegistry.InvalidGroupSize();
        if (m.maxValue > MAX_48) revert IProcessRegistry.BallotModeMaxValueTooLarge();
        if (m.minValue > MAX_48) revert IProcessRegistry.BallotModeMinValueTooLarge();
        if (m.maxValueSum > MAX_63) revert IProcessRegistry.BallotModeMaxValueSumTooLarge();
        if (m.minValueSum > MAX_63) revert IProcessRegistry.BallotModeMinValueSumTooLarge();

        packed = uint256(m.numFields);
        packed |= uint256(m.groupSize) << 8;
        packed |= uint256(m.uniqueValues ? 1 : 0) << 16;
        packed |= uint256(m.costExponent) << 17;
        packed |= m.maxValue << 25;
        packed |= m.minValue << 73;
        packed |= m.maxValueSum << 121;
        packed |= m.minValueSum << 184;
    }

    /// @notice Encryption key leaf: sha256(x_BE32 ‖ y_BE32) over the TE coordinates.
    function encKeyLeaf(DAVINCITypes.EncryptionKey calldata k) internal pure returns (uint256) {
        return uint256(sha256(abi.encodePacked(k.x, k.y)));
    }

    /// @notice True for a canonical point on the circomlib BabyJubJub curve with x != 0,
    ///         which rules out the identity, the order-2 point and gnark's reduced-form
    ///         coordinates. Subgroup membership is left to the batch guest.
    function isValidEncryptionKey(DAVINCITypes.EncryptionKey calldata k) internal pure returns (bool) {
        uint256 x = k.x;
        uint256 y = k.y;
        if (x == 0 || x >= BJJ_P || y >= BJJ_P) return false;
        uint256 x2 = mulmod(x, x, BJJ_P);
        uint256 y2 = mulmod(y, y, BJJ_P);
        uint256 lhs = addmod(mulmod(BJJ_A, x2, BJJ_P), y2, BJJ_P);
        uint256 rhs = addmod(1, mulmod(BJJ_D, mulmod(x2, y2, BJJ_P), BJJ_P), BJJ_P);
        return lhs == rhs;
    }
}
