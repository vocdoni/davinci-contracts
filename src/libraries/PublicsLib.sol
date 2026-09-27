// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/**
 * @title PublicsLib
 * @notice Reads the publicValues of a davinci-zkvm PLONK (ZisK 1.3): 64 registers, each
 *         the guest's u32 output as an 8-byte little-endian word, 512 bytes in total.
 * @dev Callers check the length first; register indices are constants.
 */
library PublicsLib {
    /// @dev publicValues length: 64 registers of 8 bytes.
    uint256 internal constant LENGTH = 512;

    /// @notice Register k as an integer.
    function word(bytes calldata pv, uint256 k) internal pure returns (uint256 w) {
        assembly ("memory-safe") {
            w := shr(192, calldataload(add(pv.offset, shl(3, k))))
        }
        w = reverse64(w);
    }

    /// @notice A 256-bit output spread over registers k..k+7: the low 4 bytes of each word,
    ///         in order. The guest writes a value as its little-endian limbs, so this is the
    ///         value's LE encoding; for a state root, its raw SHA-256 digest.
    function reg32(bytes calldata pv, uint256 k) internal pure returns (bytes32) {
        uint256 acc;
        unchecked {
            for (uint256 j = 0; j < 8; ++j) {
                uint256 w;
                assembly ("memory-safe") {
                    w := calldataload(add(pv.offset, shl(3, add(k, j))))
                }
                acc |= (w >> 224) << (224 - 32 * j);
            }
        }
        return bytes32(acc);
    }

    /// @notice Byte-reverses the low 8 bytes of x.
    function reverse64(uint256 x) internal pure returns (uint256) {
        x = ((x & 0xFF00FF00FF00FF00) >> 8) | ((x & 0x00FF00FF00FF00FF) << 8);
        x = ((x & 0xFFFF0000FFFF0000) >> 16) | ((x & 0x0000FFFF0000FFFF) << 16);
        return (x >> 32) | ((x & 0xFFFFFFFF) << 32);
    }

    /// @notice Byte-reverses 32 bytes: turns an LE encoding into the BE one and back.
    function reverse32(bytes32 x) internal pure returns (bytes32) {
        uint256 v = uint256(x);
        v = ((v >> 8) & 0x00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF)
            | ((v & 0x00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF00FF) << 8);
        v = ((v >> 16) & 0x0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF)
            | ((v & 0x0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF0000FFFF) << 16);
        v = ((v >> 32) & 0x00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF)
            | ((v & 0x00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF00000000FFFFFFFF) << 32);
        v = ((v >> 64) & 0x0000000000000000FFFFFFFFFFFFFFFF0000000000000000FFFFFFFFFFFFFFFF)
            | ((v & 0x0000000000000000FFFFFFFFFFFFFFFF0000000000000000FFFFFFFFFFFFFFFF) << 64);
        return bytes32((v >> 128) | (v << 128));
    }
}
