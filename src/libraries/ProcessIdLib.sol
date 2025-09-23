// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

library ProcessIdLib {
    /**
     * @notice Computes a processId with a hashed 4-byte prefix.
     * Prefix is the last 4 bytes of keccak256(abi.encodePacked(chainId, otherAddr)).
     * Layout (big-endian where applicable):
     * - [0..3]   : hashed prefix (uint32, big endian)
     * - [4..23]  : address (20 bytes)
     * - [24..31] : nonce (uint64, big endian)
     *
     * @dev nonce is limited to uint64 (truncates high bits of uint256 if any).
     */
    function computeProcessId(
        uint32 chainId,
        address contractAddr,
        address creatorAddr,
        uint64 nonce
    ) internal pure returns (bytes32 processId) {
        // keccak(chainId, otherAddr) and take the LAST 4 bytes (least significant 32 bits)
        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 prefix = uint32(uint256(h)); // last 4 bytes

        // Build the 32-byte value:
        // - prefix in the top 4 bytes (<< 224)
        // - mainAddr in the middle 20 bytes (<< 64)
        // - nonce in the last 8 bytes (no shift; goes to least significant 64 bits)
        processId = bytes32(
            (uint256(prefix) << 224) |
            (uint256(uint160(creatorAddr)) << 64) |
            uint256(nonce)
        );
    }
}
