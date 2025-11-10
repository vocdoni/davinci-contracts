// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

/// @title BlobsLib - Low-level helpers for EIP-4844 blob opcodes & precompile
/// @notice This library provides secure and efficient access to EIP-4844 blob functionality
library BlobsLib {
    /*//////////////////////////////////////////////////////////////
                                STRUCTS
    //////////////////////////////////////////////////////////////*/

    /// @notice KZG proof structure containing all components for verification
    /// @dev Matches the input format required by the KZG precompile
    /// @param versionedHash  32 bytes (0x01‖sha256(commitment))
    /// @param z              32 bytes challenge point   (BLS12‑381 Fr
    /// @param y              32 bytes evaluation value  (BLS12‑381 Fr)
    /// @param commitment     48 bytes G1 (BLS12‑381)
    /// @param proof          48 bytes G1 (BLS12‑381)
    struct KZGProof {
        bytes32 versionedHash;
        bytes32 z;
        bytes32 y;
        bytes commitment;
        bytes proof;
    }

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

    /// @dev Length of the KZG commitment (48 bytes, G1 compressed)
    uint256 private constant KZG_COMMITMENT_LENGTH = 48;

    /// @dev Length of the KZG proof (48 bytes, G1 compressed)
    uint256 private constant KZG_PROOF_LENGTH = 48;

    /// @dev Number of field elements per blob (as per EIP-4844)
    uint256 private constant FIELD_ELEMENTS_PER_BLOB = 4096;

    /// @dev Mask to keep the lower 31 bytes, clearing the MSB
    uint256 private constant MASK_LOW_31_BYTES =
        0x00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;

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


    /// @notice Verifies a KZG point‑evaluation proof through the
    ///         EIP‑4844 precompile at address 0x0A.
    ///
    /// @dev  Input  (192 bytes) = versionedHash ‖ z ‖ y ‖ commitment ‖ proof
    ///       Return (64 bytes)  = FIELD_ELEMENTS_PER_BLOB (4096) ‖
    ///                            BLS_MODULUS
    ///
    ///       A correct proof *always* returns 64 bytes; the call reverts
    ///       or returns empty data on failure. See EIP‑4844.
    ///
    /// @param input  Exactly 192 bytes, formatted as above.
    /// @return success  True if the precompile accepted the proof.
    function verifyKZG(bytes memory input) internal view returns (bool success) {
        if (input.length != KZG_INPUT_LENGTH) return false; // 192 bytes

        (bool ok, bytes memory out) = KZG_PRECOMPILE.staticcall(input);

        // call did not revert and returned the canonical 64‑byte payload
        if (!ok || out.length != KZG_OUTPUT_LENGTH) return false; 

        uint256 resultValue;
        assembly {
            resultValue := mload(add(out, 0x20))
        }
        // the first 32 bytes of the output should be FIELD_ELEMENTS_PER_BLOB
        if (resultValue != FIELD_ELEMENTS_PER_BLOB) return false;

        return true;
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

    /// @param commitment  KZG commitment (48 bytes).  Zero‑length is allowed and
    ///                    returns 0x00…00 (same as the Go nil‑check).
    /// @return vh         32‑byte versioned blob hash whose first byte is 0x01
    ///                    and the remaining 31 bytes are SHA‑256(commitment).
    function calcBlobHashV1(bytes memory commitment) internal pure returns (bytes32 vh) {
        if (commitment.length == 0) {
            return bytes32(0);
        }

        vh = sha256(commitment);

        unchecked {
            uint256 v = (uint256(vh) & MASK_LOW_31_BYTES) | (uint256(0x01) << 248);
            vh = bytes32(v);
        }
    }


    /// @notice Builds the input for the KZG precompile
    /// @param versionedHash  32 bytes (0x01‖sha256(commitment))
    /// @param z              32 bytes challenge point   (BLS12‑381 Fr, big‑endian)
    /// @param y              32 bytes evaluation value  (BLS12‑381 Fr, big‑endian)
    /// @param commitment     48 bytes G1 (BLS12‑381)
    /// @param proof          48 bytes G1 (BLS12‑381)
    /// @return input         192‑byte input
    function buildKZGInput(
        bytes32 versionedHash,
        bytes32 z,
        bytes32 y,
        bytes memory commitment,
        bytes memory proof
    ) internal pure returns (bytes memory input) {
        if (commitment.length != KZG_COMMITMENT_LENGTH || proof.length != KZG_PROOF_LENGTH) {
            revert("KZGInput: bad G1 length");
        }

        input = abi.encodePacked(versionedHash, z, y, commitment, proof);

        assert(input.length == KZG_INPUT_LENGTH);
    }

    /// @notice Decodes the KZG input data back into its components
    /// @param input 192-byte KZG input data
    /// @return kzgProof The decoded KZG proof components
    function decodeKZGInput(bytes memory input) internal pure returns (KZGProof memory kzgProof) {
        if (input.length != KZG_INPUT_LENGTH) {
            revert("KZGInput: invalid input length");
        }

        // Extract versionedHash (bytes 0-31)
        bytes32 versionedHash;
        assembly {
            versionedHash := mload(add(input, 0x20))
        }
        kzgProof.versionedHash = versionedHash;

        // Extract z (bytes 32-63)
        bytes32 z;
        assembly {
            z := mload(add(input, 0x40))
        }
        kzgProof.z = z;

        // Extract y (bytes 64-95)
        bytes32 y;
        assembly {
            y := mload(add(input, 0x60))
        }
        kzgProof.y = y;

        // Extract commitment (bytes 96-143, 48 bytes)
        kzgProof.commitment = new bytes(KZG_COMMITMENT_LENGTH);
        for (uint256 i = 0; i < KZG_COMMITMENT_LENGTH; i++) {
            kzgProof.commitment[i] = input[96 + i];
        }

        // Extract proof (bytes 144-191, 48 bytes)
        kzgProof.proof = new bytes(KZG_PROOF_LENGTH);
        for (uint256 i = 0; i < KZG_PROOF_LENGTH; i++) {
            kzgProof.proof[i] = input[144 + i];
        }
    }
}
