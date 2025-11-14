// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.24;

import { Test, console2 } from "forge-std/Test.sol";
import { BlobsLib } from "../src/libraries/BlobsLib.sol";

/// @title Helper contract to test BlobsLib reverts
contract BlobsLibTestHelper {
    function buildKZGInput(
        bytes32 versionedHash,
        bytes32 z,
        bytes32 y,
        bytes memory commitment,
        bytes memory proof
    ) external pure returns (bytes memory) {
        return BlobsLib.buildKZGInput(versionedHash, z, y, commitment, proof);
    }

    function decodeKZGInput(bytes memory input) external pure returns (BlobsLib.KZGProof memory) {
        return BlobsLib.decodeKZGInput(input);
    }
}

/// @title BlobsLib Test Suite
/// @notice Comprehensive tests for the BlobsLib library covering all functions and edge cases
contract BlobsLibTest is Test {
    using BlobsLib for *;

    BlobsLibTestHelper helper;

    /*//////////////////////////////////////////////////////////////
                            TEST CONSTANTS
    //////////////////////////////////////////////////////////////*/

    // Test data for KZG operations
    bytes32 constant TEST_VERSIONED_HASH = 0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef;
    bytes32 constant TEST_Z = 0x1111111111111111111111111111111111111111111111111111111111111111;
    bytes32 constant TEST_Y = 0x2222222222222222222222222222222222222222222222222222222222222222;

    // 48-byte test commitment (G1 compressed point)
    bytes constant TEST_COMMITMENT =
        hex"123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0";

    // 48-byte test proof (G1 compressed point)
    bytes constant TEST_PROOF =
        hex"abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef012345678a";

    // Expected constants from BlobsLib
    uint256 constant FIELD_ELEMENTS_PER_BLOB = 4096;
    uint256 constant KZG_INPUT_LENGTH = 192;
    uint256 constant KZG_OUTPUT_LENGTH = 64;
    uint256 constant KZG_COMMITMENT_LENGTH = 48;
    uint256 constant KZG_PROOF_LENGTH = 48;

    /*//////////////////////////////////////////////////////////////
                            SETUP
    //////////////////////////////////////////////////////////////*/

    function setUp() public {
        // Setup test environment
        vm.label(address(this), "BlobsLibTest");
        helper = new BlobsLibTestHelper();
        vm.label(address(helper), "BlobsLibTestHelper");
    }

    /*//////////////////////////////////////////////////////////////
                        BLOB OPERATIONS TESTS
    //////////////////////////////////////////////////////////////*/

    /// @notice Test blobHash function with various indices
    function test_blobHash() public view {
        // Test with index 0 (should return 0 in normal test environment)
        bytes32 hash0 = BlobsLib.blobHash(0);
        assertEq(hash0, bytes32(0), "blobHash(0) should return 0 in test environment");

        // Test with higher indices
        bytes32 hash1 = BlobsLib.blobHash(1);
        assertEq(hash1, bytes32(0), "blobHash(1) should return 0 in test environment");

        bytes32 hash5 = BlobsLib.blobHash(5);
        assertEq(hash5, bytes32(0), "blobHash(5) should return 0 in test environment");
    }

    /// @notice Test blobHash with maximum index
    function test_blobHash_MaxIndex() public view {
        bytes32 hashMax = BlobsLib.blobHash(type(uint256).max);
        assertEq(hashMax, bytes32(0), "blobHash with max index should return 0");
    }

    /// @notice Test blobBaseFee function
    function test_blobBaseFee() public view {
        uint256 baseFee = BlobsLib.blobBaseFee();
        // In test environment, this should return 0 or a default value
        // We just verify it doesn't revert
        assertTrue(baseFee >= 0, "blobBaseFee should not revert");
    }

    /// @notice Test calculateBlobFee with zero blobs
    function test_calculateBlobFee_ZeroBlobs() public view {
        uint256 fee = BlobsLib.calculateBlobFee(0);
        assertEq(fee, 0, "Fee for 0 blobs should be 0");
    }

    /// @notice Test calculateBlobFee with various blob counts
    function test_calculateBlobFee_VariousCounts() public view {
        uint256 fee1 = BlobsLib.calculateBlobFee(1);
        uint256 fee2 = BlobsLib.calculateBlobFee(2);
        uint256 fee6 = BlobsLib.calculateBlobFee(6);

        // Verify fees scale correctly (in test environment, base fee is 0, so all fees will be 0)
        // But we can still verify the function doesn't revert and returns consistent results
        assertTrue(fee2 >= fee1, "Fee for 2 blobs should be >= fee for 1 blob");
        assertTrue(fee6 >= fee2, "Fee for 6 blobs should be >= fee for 2 blobs");
    }

    /// @notice Test calculateBlobFee with maximum blob count
    function test_calculateBlobFee_MaxBlobs() public view {
        // Test with maximum reasonable blob count (6 as per EIP-4844)
        uint256 fee = BlobsLib.calculateBlobFee(6);
        assertTrue(fee >= 0, "Fee calculation should not revert for max blobs");
    }

    /*//////////////////////////////////////////////////////////////
                        KZG OPERATIONS TESTS
    //////////////////////////////////////////////////////////////*/

    /// @notice Test verifyKZG with invalid input length
    function test_verifyKZG_InvalidInputLength() public view {
        bytes memory shortInput = new bytes(100);
        bool result = BlobsLib.verifyKZG(shortInput);
        assertFalse(result, "verifyKZG should return false for invalid input length");

        bytes memory longInput = new bytes(300);
        result = BlobsLib.verifyKZG(longInput);
        assertFalse(result, "verifyKZG should return false for too long input");
    }

    /// @notice Test verifyKZG with correct input length
    function test_verifyKZG_CorrectInputLength() public view {
        bytes memory validInput = new bytes(KZG_INPUT_LENGTH);
        // Fill with test data
        for (uint256 i = 0; i < KZG_INPUT_LENGTH; i++) {
            validInput[i] = bytes1(uint8(i % 256));
        }

        // In test environment, KZG precompile is not available, so this will return false
        bool result = BlobsLib.verifyKZG(validInput);
        assertFalse(result, "verifyKZG should return false in test environment");
    }

    /// @notice Test verifyKZG with empty input
    function test_verifyKZG_EmptyInput() public view {
        bytes memory emptyInput = new bytes(0);
        bool result = BlobsLib.verifyKZG(emptyInput);
        assertFalse(result, "verifyKZG should return false for empty input");
    }

    /*//////////////////////////////////////////////////////////////
                        UTILITY FUNCTIONS TESTS
    //////////////////////////////////////////////////////////////*/

    /// @notice Test isBlobDataAvailable function
    function test_isBlobDataAvailable() public view {
        bool available = BlobsLib.isBlobDataAvailable();
        assertFalse(available, "Blob data should not be available in test environment");
    }

    /// @notice Test getBlobCount function
    function test_getBlobCount() public view {
        uint256 count = BlobsLib.getBlobCount();
        assertEq(count, 0, "Blob count should be 0 in test environment");
    }

    /// @notice Test getAllBlobHashes function
    function test_getAllBlobHashes() public view {
        bytes32[] memory hashes = BlobsLib.getAllBlobHashes();
        assertEq(hashes.length, 0, "Should return empty array in test environment");
    }

    /// @notice Test calcBlobHashV1 with empty commitment
    function test_calcBlobHashV1_EmptyCommitment() public pure {
        bytes memory emptyCommitment = new bytes(0);
        bytes32 hash = BlobsLib.calcBlobHashV1(emptyCommitment);
        assertEq(hash, bytes32(0), "Empty commitment should return zero hash");
    }

    /// @notice Test calcBlobHashV1 with valid commitment
    function test_calcBlobHashV1_ValidCommitment() public pure {
        bytes32 hash = BlobsLib.calcBlobHashV1(TEST_COMMITMENT);

        // Verify the hash has version byte 0x01
        uint8 versionByte = uint8(uint256(hash) >> 248);
        assertEq(versionByte, 0x01, "Version byte should be 0x01");

        // Verify it's not zero
        assertTrue(hash != bytes32(0), "Hash should not be zero for valid commitment");
    }

    /// @notice Test calcBlobHashV1 with various commitment sizes
    function test_calcBlobHashV1_VariousSizes() public pure {
        // Test with 1 byte
        bytes memory smallCommitment = new bytes(1);
        smallCommitment[0] = 0xFF;
        bytes32 hash1 = BlobsLib.calcBlobHashV1(smallCommitment);
        assertTrue(hash1 != bytes32(0), "Should handle small commitments");

        // Test with 32 bytes
        bytes memory mediumCommitment = new bytes(32);
        for (uint256 i = 0; i < 32; i++) {
            mediumCommitment[i] = bytes1(uint8(i));
        }
        bytes32 hash2 = BlobsLib.calcBlobHashV1(mediumCommitment);
        assertTrue(hash2 != bytes32(0), "Should handle medium commitments");

        // Test with 48 bytes (standard KZG commitment size)
        bytes32 hash3 = BlobsLib.calcBlobHashV1(TEST_COMMITMENT);
        assertTrue(hash3 != bytes32(0), "Should handle standard commitments");

        // All hashes should be different
        assertTrue(hash1 != hash2, "Different commitments should produce different hashes");
        assertTrue(hash2 != hash3, "Different commitments should produce different hashes");
        assertTrue(hash1 != hash3, "Different commitments should produce different hashes");
    }

    /// @notice Test buildKZGInput with valid parameters
    function test_buildKZGInput_ValidParameters() public pure {
        bytes memory input = BlobsLib.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, TEST_PROOF);

        assertEq(input.length, KZG_INPUT_LENGTH, "Input should be 192 bytes");

        // Verify the structure
        bytes32 extractedHash;
        assembly {
            extractedHash := mload(add(input, 0x20))
        }
        assertEq(extractedHash, TEST_VERSIONED_HASH, "Versioned hash should match");
    }

    /// @notice Test buildKZGInput with invalid commitment length
    function test_buildKZGInput_InvalidCommitmentLength() public {
        bytes memory shortCommitment = new bytes(47);

        vm.expectRevert("KZGInput: bad G1 length");
        helper.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, shortCommitment, TEST_PROOF);
    }

    /// @notice Test buildKZGInput with invalid proof length
    function test_buildKZGInput_InvalidProofLength() public {
        bytes memory shortProof = new bytes(47);

        vm.expectRevert("KZGInput: bad G1 length");
        helper.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, shortProof);
    }

    /// @notice Test buildKZGInput with zero-length commitment
    function test_buildKZGInput_ZeroLengthCommitment() public {
        bytes memory emptyCommitment = new bytes(0);

        vm.expectRevert("KZGInput: bad G1 length");
        helper.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, emptyCommitment, TEST_PROOF);
    }

    /// @notice Test buildKZGInput with oversized commitment
    function test_buildKZGInput_OversizedCommitment() public {
        bytes memory longCommitment = new bytes(49);

        vm.expectRevert("KZGInput: bad G1 length");
        helper.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, longCommitment, TEST_PROOF);
    }

    /// @notice Test decodeKZGInput with valid input
    function test_decodeKZGInput_ValidInput() public pure {
        bytes memory input = BlobsLib.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, TEST_PROOF);

        BlobsLib.KZGProof memory proof = BlobsLib.decodeKZGInput(input);

        assertEq(proof.versionedHash, TEST_VERSIONED_HASH, "Versioned hash should match");
        assertEq(proof.z, TEST_Z, "Z should match");
        assertEq(proof.y, TEST_Y, "Y should match");
        assertEq(proof.commitment.length, KZG_COMMITMENT_LENGTH, "Commitment length should be 48");
        assertEq(proof.proof.length, KZG_PROOF_LENGTH, "Proof length should be 48");

        // Verify commitment bytes
        for (uint256 i = 0; i < KZG_COMMITMENT_LENGTH; i++) {
            assertEq(proof.commitment[i], TEST_COMMITMENT[i], "Commitment bytes should match");
        }

        // Verify proof bytes
        for (uint256 i = 0; i < KZG_PROOF_LENGTH; i++) {
            assertEq(proof.proof[i], TEST_PROOF[i], "Proof bytes should match");
        }
    }

    /// @notice Test decodeKZGInput with invalid input length
    function test_decodeKZGInput_InvalidInputLength() public {
        bytes memory shortInput = new bytes(100);

        vm.expectRevert("KZGInput: invalid input length");
        helper.decodeKZGInput(shortInput);
    }

    /// @notice Test decodeKZGInput with empty input
    function test_decodeKZGInput_EmptyInput() public {
        bytes memory emptyInput = new bytes(0);

        vm.expectRevert("KZGInput: invalid input length");
        helper.decodeKZGInput(emptyInput);
    }

    /// @notice Test decodeKZGInput with oversized input
    function test_decodeKZGInput_OversizedInput() public {
        bytes memory longInput = new bytes(300);

        vm.expectRevert("KZGInput: invalid input length");
        helper.decodeKZGInput(longInput);
    }

    /*//////////////////////////////////////////////////////////////
                        INTEGRATION TESTS
    //////////////////////////////////////////////////////////////*/

    /// @notice Test round-trip encoding/decoding of KZG input
    function test_KZGInput_RoundTrip() public pure {
        // Build input
        bytes memory input = BlobsLib.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, TEST_PROOF);

        // Decode input
        BlobsLib.KZGProof memory proof = BlobsLib.decodeKZGInput(input);

        // Rebuild input from decoded proof
        bytes memory rebuiltInput = BlobsLib.buildKZGInput(
            proof.versionedHash,
            proof.z,
            proof.y,
            proof.commitment,
            proof.proof
        );

        // Verify they match
        assertEq(input.length, rebuiltInput.length, "Input lengths should match");
        for (uint256 i = 0; i < input.length; i++) {
            assertEq(input[i], rebuiltInput[i], "Input bytes should match");
        }
    }

    /// @notice Test calcBlobHashV1 with commitment from buildKZGInput
    function test_calcBlobHashV1_Integration() public pure {
        bytes32 calculatedHash = BlobsLib.calcBlobHashV1(TEST_COMMITMENT);

        bytes memory input = BlobsLib.buildKZGInput(calculatedHash, TEST_Z, TEST_Y, TEST_COMMITMENT, TEST_PROOF);

        BlobsLib.KZGProof memory proof = BlobsLib.decodeKZGInput(input);
        assertEq(proof.versionedHash, calculatedHash, "Calculated hash should match");
    }

    /*//////////////////////////////////////////////////////////////
                        FUZZ TESTS
    //////////////////////////////////////////////////////////////*/

    /// @notice Fuzz test for calcBlobHashV1
    function testFuzz_calcBlobHashV1(bytes memory commitment) public pure {
        bytes32 hash = BlobsLib.calcBlobHashV1(commitment);

        if (commitment.length == 0) {
            assertEq(hash, bytes32(0), "Empty commitment should return zero");
        } else {
            // Version byte should always be 0x01
            uint8 versionByte = uint8(uint256(hash) >> 248);
            assertEq(versionByte, 0x01, "Version byte should be 0x01");
        }
    }

    /// @notice Fuzz test for calculateBlobFee
    function testFuzz_calculateBlobFee(uint256 blobCount) public view {
        // Limit blob count to reasonable range to avoid overflow
        blobCount = bound(blobCount, 0, 100);

        uint256 fee = BlobsLib.calculateBlobFee(blobCount);

        if (blobCount == 0) {
            assertEq(fee, 0, "Zero blobs should have zero fee");
        } else {
            assertTrue(fee >= 0, "Fee should be non-negative");
        }
    }

    /// @notice Fuzz test for blobHash
    function testFuzz_blobHash(uint256 idx) public view {
        // Limit index to reasonable range
        idx = bound(idx, 0, 1000);

        bytes32 hash = BlobsLib.blobHash(idx);
        // In test environment, should always return 0
        assertEq(hash, bytes32(0), "Should return zero in test environment");
    }

    /*//////////////////////////////////////////////////////////////
                        GAS OPTIMIZATION TESTS
    //////////////////////////////////////////////////////////////*/

    /// @notice Test gas usage for calcBlobHashV1
    function test_gas_calcBlobHashV1() public view {
        uint256 gasBefore = gasleft();
        BlobsLib.calcBlobHashV1(TEST_COMMITMENT);
        uint256 gasUsed = gasBefore - gasleft();

        // Verify reasonable gas usage (should be less than 10k gas)
        assertTrue(gasUsed < 10000, "calcBlobHashV1 should use reasonable gas");
    }

    /// @notice Test gas usage for buildKZGInput
    function test_gas_buildKZGInput() public view {
        uint256 gasBefore = gasleft();
        BlobsLib.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, TEST_PROOF);
        uint256 gasUsed = gasBefore - gasleft();

        // Verify reasonable gas usage
        assertTrue(gasUsed < 20000, "buildKZGInput should use reasonable gas");
    }

    /// @notice Test gas usage for decodeKZGInput
    function test_gas_decodeKZGInput() public view {
        bytes memory input = BlobsLib.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, TEST_PROOF);

        uint256 gasBefore = gasleft();
        BlobsLib.decodeKZGInput(input);
        uint256 gasUsed = gasBefore - gasleft();

        // Verify reasonable gas usage
        assertTrue(gasUsed < 30000, "decodeKZGInput should use reasonable gas");
    }
}
