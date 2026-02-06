// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

library ProcessIdLib {
    /**
     * @notice Computes a processId with a hashed 4-byte prefix of chainId and contractAddr.
     * Layout (big-endian where applicable):
     * - [0..19]  : creatorAddr (20 bytes)
     * - [20..23] : hashed prefix (uint32, big endian)
     * - [24..31] : nonce (uint64, big endian)
     *
     * @dev nonce is limited to uint64 (truncates high bits of uint256 if any).
     */
    function computeProcessId(uint32 prefix, address creatorAddr, uint64 nonce)
        internal
        pure
        returns (bytes32 processId)
    {
        // Build the 32-byte value:
        // - creatorAddr in the top 20 bytes (<< 96)
        // - prefix in bytes 20-23 (<< 64)
        // - nonce in the last 8 bytes (no shift; goes to least significant 64 bits)
        processId = bytes32((uint256(uint160(creatorAddr)) << 96) | (uint256(prefix) << 64) | uint256(nonce));
    }

    /**
     * @notice Computes the 4-byte prefix from chainId and contractAddr.
     * Prefix is the last 4 bytes of keccak256(abi.encodePacked(chainId, contractAddr)).
     */
    function getPrefix(uint32 chainId, address contractAddr) internal pure returns (uint32) {
        // keccak(chainId, otherAddr) and take the LAST 4 bytes (least significant 32 bits)
        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        return uint32(uint256(h)); // last 4 bytes
    }

    /**
     * @notice Checks if the given processId has the expected prefix.
     */
    function hasPrefix(bytes32 processId, uint32 expectedPrefix) internal pure returns (bool) {
        return uint32(uint256(processId >> 64)) == expectedPrefix;
    }
}
