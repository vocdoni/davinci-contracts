// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {PublicsLib} from "./PublicsLib.sol";

/**
 * @title Sha256SmtLib
 * @notice Root of a vocdoni/arbo SHA-256 sparse Merkle tree with 64 levels and 8-byte
 *         keys, the davinci state tree:
 *         leaf = sha256(key_le8 ‖ value_le32 ‖ 0x01), node = sha256(left ‖ right),
 *         empty = 32 zero bytes. Bit i of the key (LSB first) picks the child at depth i,
 *         and a subtree holding one leaf is that leaf's hash.
 */
library Sha256SmtLib {
    uint256 internal constant MAX_LEVELS = 64;

    error SmtMaxLevelsReached();
    error SmtLengthMismatch();

    function leafHash(uint64 key, uint256 value) internal pure returns (bytes32) {
        return sha256(
            abi.encodePacked(bytes8(uint64(PublicsLib.reverse64(key))), PublicsLib.reverse32(bytes32(value)), hex"01")
        );
    }

    function nodeHash(bytes32 left, bytes32 right) internal pure returns (bytes32) {
        return sha256(abi.encodePacked(left, right));
    }

    /// @notice Root of the tree holding (keys[i], values[i]). Keys must be distinct.
    function root(uint64[] memory keys, uint256[] memory values) internal pure returns (bytes32) {
        if (keys.length != values.length) revert SmtLengthMismatch();
        bytes32[] memory leaves = new bytes32[](keys.length);
        for (uint256 i = 0; i < keys.length; ++i) {
            leaves[i] = leafHash(keys[i], values[i]);
        }
        return _subtree(keys, leaves, 0);
    }

    function _subtree(uint64[] memory keys, bytes32[] memory leaves, uint256 level) private pure returns (bytes32) {
        uint256 n = keys.length;
        if (n == 0) return bytes32(0);
        if (n == 1) return leaves[0];
        // Distinct u64 keys always split before level 64.
        if (level == MAX_LEVELS) revert SmtMaxLevelsReached();

        uint256 nRight;
        for (uint256 i = 0; i < n; ++i) {
            nRight += (keys[i] >> level) & 1;
        }
        uint64[] memory lk = new uint64[](n - nRight);
        bytes32[] memory ll = new bytes32[](n - nRight);
        uint64[] memory rk = new uint64[](nRight);
        bytes32[] memory rl = new bytes32[](nRight);
        uint256 a;
        uint256 b;
        for (uint256 i = 0; i < n; ++i) {
            if (((keys[i] >> level) & 1) == 0) {
                (lk[a], ll[a]) = (keys[i], leaves[i]);
                ++a;
            } else {
                (rk[b], rl[b]) = (keys[i], leaves[i]);
                ++b;
            }
        }
        return nodeHash(_subtree(lk, ll, level + 1), _subtree(rk, rl, level + 1));
    }
}
