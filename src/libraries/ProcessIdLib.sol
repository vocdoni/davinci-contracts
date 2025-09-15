// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

library ProcessIdLib {
    /**
     * @notice Computes a deterministic processId (bytes32) from chainId, address, and its current blockchain nonce.
     * Format:
     * - [0..3]   : chainId (uint32, big endian)
     * - [4..23]  : address (20 bytes)
     * - [24..31] : nonce (uint64, big endian) pulled from address custom nonce
     * @dev Note: nonce is limited to uint64 (truncates high bits of uint256)
     */
    function computeProcessId(
        uint32 chainId,
        address _address,
        uint64 nonce
    ) internal pure returns (bytes32 processId) {
        processId = bytes32(0);

        // encode chainId (big endian, 4 bytes)
        processId |= bytes32(uint256(uint8(chainId >> 24)) << 248);
        processId |= bytes32(uint256(uint8(chainId >> 16)) << 240);
        processId |= bytes32(uint256(uint8(chainId >> 8)) << 232);
        processId |= bytes32(uint256(uint8(chainId)) << 224);

        // encode address (middle 20 bytes)
        processId |= bytes32(uint256(uint160(_address)) << 64);

        // encode nonce (big endian, last 8 bytes)
        processId |= bytes32(uint256(uint8(nonce >> 56)) << 56);
        processId |= bytes32(uint256(uint8(nonce >> 48)) << 48);
        processId |= bytes32(uint256(uint8(nonce >> 40)) << 40);
        processId |= bytes32(uint256(uint8(nonce >> 32)) << 32);
        processId |= bytes32(uint256(uint8(nonce >> 24)) << 24);
        processId |= bytes32(uint256(uint8(nonce >> 16)) << 16);
        processId |= bytes32(uint256(uint8(nonce >> 8)) << 8);
        processId |= bytes32(uint256(uint8(nonce)));
    }
}
