// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

/// @title BlobsLib - Low-level helpers for EIP-4844 blob opcodes & precompile
/// @notice This library provides secure and efficient access to EIP-4844 blob functionality
/// @dev All functions are marked as internal view for gas efficiency and security
library BlobsLib {
    /*//////////////////////////////////////////////////////////////
                                CONSTANTS
    //////////////////////////////////////////////////////////////*/

    /// @dev KZG precompile address as per EIP-4844
    address private constant KZG_PRECOMPILE = address(0x0A);

    /// @dev Expected input length for KZG proof verification (192 bytes)
    /// 48 bytes (versioned_hash) + 48 bytes (z) + 48 bytes (y) + 48 bytes (commitment) + 48 bytes (proof)
    uint256 private constant KZG_INPUT_LENGTH = 192;

    /// @dev Expected output length for successful KZG verification (64 bytes)
    uint256 private constant KZG_OUTPUT_LENGTH = 64;

    /*//////////////////////////////////////////////////////////////
                            BLOB OPERATIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the versioned hash of the blob at the specified index
    /// @dev Uses the BLOBHASH opcode (0x49) introduced in EIP-4844
    /// @param idx The index of the blob hash to retrieve
    /// @return h The versioned hash of the blob, or 0x0 if idx >= blob_versioned_hashes.length
    function blobHash(uint256 idx) internal view returns (bytes32 h) {
        assembly ("memory-safe") {
            h := blobhash(idx)
        }
    }

    /// @notice Returns the current blob base fee
    /// @dev Uses the BLOBBASEFEE opcode (0x4A) introduced in EIP-7516
    /// @dev Useful for calculating blob fees and implementing in-protocol refunds
    /// @return fee The current blob base fee in wei
    function blobBaseFee() internal view returns (uint256 fee) {
        assembly ("memory-safe") {
            fee := blobbasefee()
        }
    }

    /// @notice Calculates the total blob fee for a given number of blobs
    /// @dev Helper function to calculate blob fees using current base fee
    /// @param blobCount The number of blobs to calculate fee for
    /// @return totalFee The total blob fee in wei
    function calculateBlobFee(uint256 blobCount) internal view returns (uint256 totalFee) {
        if (blobCount == 0) return 0;

        uint256 baseFee = blobBaseFee();
        // Each blob costs baseFee * GAS_PER_BLOB (131072 gas per blob as per EIP-4844)
        unchecked {
            totalFee = baseFee * blobCount * 131072;
        }
    }

    /*//////////////////////////////////////////////////////////////
                            KZG OPERATIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Verifies a KZG point-evaluation proof using the precompile
    /// @dev Calls the KZG point evaluation precompile at address 0x0A
    /// @dev Input format: versioned_hash (32) + z (32) + y (32) + commitment (48) + proof (48) = 192 bytes
    /// @param input The KZG proof data (must be exactly 192 bytes)
    /// @return success True if the proof is valid, false otherwise
    function verifyKZG(bytes memory input) internal view returns (bool success) {
        if (input.length != KZG_INPUT_LENGTH) {
            return false;
        }

        (bool callSuccess, bytes memory result) = KZG_PRECOMPILE.staticcall(input);

        if (!callSuccess) {
           return false;
        }

        if (result.length != KZG_OUTPUT_LENGTH) {
            return false;
        }

        // The precompile returns 64 bytes on success, with the first 32 bytes being the result
        // A successful verification returns 0x000...001 in the first 32 bytes
        bytes32 resultValue;
        assembly {
            resultValue := mload(add(result, 0x20))
        }
        success = resultValue == bytes32(uint256(1));
    }

    /*//////////////////////////////////////////////////////////////
                            UTILITY FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Checks if blob data is available for the current transaction
    /// @dev Verifies that at least one blob hash exists
    /// @return available True if blob data is available
    function isBlobDataAvailable() internal view returns (bool available) {
        return blobHash(0) != bytes32(0);
    }

    /// @notice Gets the number of blobs in the current transaction
    /// @dev Iterates through blob hashes until finding a zero hash
    /// @return count The number of blobs in the transaction
    function getBlobCount() internal view returns (uint256 count) {
        while (blobHash(count) != bytes32(0)) {
            unchecked {
                ++count;
            }
            // Safety check to prevent infinite loops (max 6 blobs per tx as per EIP-4844)
            if (count >= 6) break;
        }
    }

    /// @notice Retrieves all blob hashes for the current transaction
    /// @dev Returns an array of all non-zero blob hashes
    /// @return hashes Array of blob versioned hashes
    function getAllBlobHashes() internal view returns (bytes32[] memory hashes) {
        uint256 count = getBlobCount();
        hashes = new bytes32[](count);

        for (uint256 i = 0; i < count; ) {
            hashes[i] = blobHash(i);
            unchecked {
                ++i;
            }
        }
    }
}
