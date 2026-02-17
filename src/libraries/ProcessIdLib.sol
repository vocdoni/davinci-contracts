// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

library ProcessIdLib {
    /**
     * @notice Computes a processId with a hashed 4-byte prefix of chainId and contractAddr.
     * Layout (big-endian where applicable):
     * - [0..19]  : creatorAddr (20 bytes)
     * - [20..23] : hashed prefix (uint32, big endian)
     * - [24..30] : nonce (uint56, big endian)
     *
     * @dev nonce is passed as uint64 and truncated to uint56.
     */
    function computeProcessId(
        uint32 prefix,
        address creatorAddr,
        uint64 nonce
    ) external pure returns (bytes31 processId) {
        // Build the 31-byte value:
        // - creatorAddr in the top 20 bytes (<< 88)
        // - prefix in bytes 20-23 (<< 56)
        // - nonce in the last 7 bytes (least significant 56 bits)
        processId = bytes31(
            uint248((uint248(uint160(creatorAddr)) << 88) | (uint248(prefix) << 56) | uint248(uint56(nonce)))
        );
    }

    /**
     * @notice Computes the 4-byte prefix from chainId and contractAddr.
     * Prefix is the last 4 bytes of keccak256(abi.encodePacked(chainId, contractAddr)).
     */
    function getPrefix(uint32 chainId, address contractAddr) external pure returns (uint32) {
        // keccak(chainId, otherAddr) and take the LAST 4 bytes (least significant 32 bits)
        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        return uint32(uint256(h)); // last 4 bytes
    }

    /**
     * @notice Checks if the given processId has the expected prefix.
     */
    function hasPrefix(bytes31 processId, uint32 expectedPrefix) external pure returns (bool) {
        return uint32(uint248(processId >> 56)) == expectedPrefix;
    }
}
