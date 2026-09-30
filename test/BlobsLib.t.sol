// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.24;

import {Test, console2} from "forge-std/Test.sol";
import {BlobsLib} from "../src/libraries/BlobsLib.sol";

/// @title Helper contract to test BlobsLib reverts
contract BlobsLibTestHelper {
    function buildKZGInput(bytes32 versionedHash, bytes32 z, bytes32 y, bytes memory commitment, bytes memory proof)
        external
        pure
        returns (bytes memory)
    {
        return BlobsLib.buildKZGInput(versionedHash, z, y, commitment, proof);
    }

    function decodeKZGInput(bytes memory input) external pure returns (BlobsLib.KZGProof memory) {
        return BlobsLib.decodeKZGInput(input);
    }

    function verifyKZG(bytes memory input) external view {
        BlobsLib.verifyKZG(input);
    }
}

/// @title BlobsLib Test Suite
/// @notice The tests run outside a blob transaction, so every BLOBHASH reads zero.
contract BlobsLibTest is Test {
    using BlobsLib for *;

    BlobsLibTestHelper helper;

    /*//////////////////////////////////////////////////////////////
                            TEST CONSTANTS
    //////////////////////////////////////////////////////////////*/

    bytes32 constant TEST_VERSIONED_HASH = 0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef;
    bytes32 constant TEST_Z = 0x1111111111111111111111111111111111111111111111111111111111111111;
    bytes32 constant TEST_Y = 0x2222222222222222222222222222222222222222222222222222222222222222;

    // Arbitrary bytes with the length of a KZG commitment and proof; not curve points.
    bytes constant TEST_COMMITMENT =
        hex"123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0";

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
        vm.label(address(this), "BlobsLibTest");
        helper = new BlobsLibTestHelper();
        vm.label(address(helper), "BlobsLibTestHelper");
    }

    /*//////////////////////////////////////////////////////////////
                        BLOB OPERATIONS TESTS
    //////////////////////////////////////////////////////////////*/

    /// @notice Test blobHash function with various indices
    function test_blobHash() public view {
        bytes32 hash0 = BlobsLib.blobHash(0);
        assertEq(hash0, bytes32(0), "blobHash(0) should return 0 in test environment");

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
        // Only checks that the call does not revert.
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

        // Only checks that the calls do not revert and the fee does not fall with the count.
        assertTrue(fee2 >= fee1, "Fee for 2 blobs should be >= fee for 1 blob");
        assertTrue(fee6 >= fee2, "Fee for 6 blobs should be >= fee for 2 blobs");
    }

    /// @notice Test calculateBlobFee with maximum blob count
    function test_calculateBlobFee_MaxBlobs() public view {
        uint256 fee = BlobsLib.calculateBlobFee(6);
        assertTrue(fee >= 0, "Fee calculation should not revert for max blobs");
    }

    /*//////////////////////////////////////////////////////////////
                        KZG OPERATIONS TESTS
    //////////////////////////////////////////////////////////////*/

    /// @notice Test verifyKZG with invalid input length
    function test_verifyKZG_ShortInputLength() public {
        uint256 badLength = 100;
        vm.expectRevert(
            abi.encodeWithSelector(BlobsLib.BlobVerificationInvalidInputLength.selector, badLength, KZG_INPUT_LENGTH)
        );
        helper.verifyKZG(new bytes(badLength));
    }

    /// @notice Test verifyKZG with invalid input length
    function test_verifyKZG_LongInputLength() public {
        uint256 badLength = 300;
        vm.expectRevert(
            abi.encodeWithSelector(BlobsLib.BlobVerificationInvalidInputLength.selector, badLength, KZG_INPUT_LENGTH)
        );
        helper.verifyKZG(new bytes(badLength));
    }

    /// @notice Test verifyKZG with correct input length
    function test_verifyKZG_CorrectInputLength() public {
        bytes memory validInput = new bytes(KZG_INPUT_LENGTH);
        for (uint8 i = 0; i < 192; i++) {
            validInput[i] = bytes1(i);
        }

        vm.expectRevert(BlobsLib.BlobVerificationPointEvaluationFailed.selector);
        helper.verifyKZG(validInput);
    }

    /// @notice Test verifyKZG with empty input
    function test_verifyKZG_EmptyInput() public {
        uint256 badLength = 0;
        vm.expectRevert(
            abi.encodeWithSelector(BlobsLib.BlobVerificationInvalidInputLength.selector, badLength, KZG_INPUT_LENGTH)
        );
        helper.verifyKZG(new bytes(badLength));
    }

    /*//////////////////////////////////////////////////////////////
                        UTILITY FUNCTIONS TESTS
    //////////////////////////////////////////////////////////////*/

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

        uint8 versionByte = uint8(uint256(hash) >> 248);
        assertEq(versionByte, 0x01, "Version byte should be 0x01");

        assertTrue(hash != bytes32(0), "Hash should not be zero for valid commitment");
    }

    /// @notice Test calcBlobHashV1 with various commitment sizes
    function test_calcBlobHashV1_VariousSizes() public pure {
        bytes memory smallCommitment = new bytes(1);
        smallCommitment[0] = 0xFF;
        bytes32 hash1 = BlobsLib.calcBlobHashV1(smallCommitment);
        assertTrue(hash1 != bytes32(0), "Should handle small commitments");

        bytes memory mediumCommitment = new bytes(32);
        for (uint8 i = 0; i < 32; i++) {
            mediumCommitment[i] = bytes1(i);
        }
        bytes32 hash2 = BlobsLib.calcBlobHashV1(mediumCommitment);
        assertTrue(hash2 != bytes32(0), "Should handle medium commitments");

        bytes32 hash3 = BlobsLib.calcBlobHashV1(TEST_COMMITMENT);
        assertTrue(hash3 != bytes32(0), "Should handle standard commitments");

        assertTrue(hash1 != hash2, "Different commitments should produce different hashes");
        assertTrue(hash2 != hash3, "Different commitments should produce different hashes");
        assertTrue(hash1 != hash3, "Different commitments should produce different hashes");
    }

    /// @notice Test buildKZGInput with valid parameters
    function test_buildKZGInput_ValidParameters() public pure {
        bytes memory input = BlobsLib.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, TEST_PROOF);

        assertEq(input.length, KZG_INPUT_LENGTH, "Input should be 192 bytes");

        bytes32 extractedHash;
        assembly {
            extractedHash := mload(add(input, 0x20))
        }
        assertEq(extractedHash, TEST_VERSIONED_HASH, "Versioned hash should match");
    }

    /// @notice Test buildKZGInput with invalid commitment length
    function test_buildKZGInput_InvalidCommitmentLength() public {
        uint256 badLength = 47;
        vm.expectRevert(
            abi.encodeWithSelector(BlobsLib.KZGInputBadCommitmentLength.selector, badLength, KZG_COMMITMENT_LENGTH)
        );
        helper.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, new bytes(badLength), TEST_PROOF);
    }

    /// @notice Test buildKZGInput with invalid proof length
    function test_buildKZGInput_InvalidProofLength() public {
        uint256 badLength = 47;
        vm.expectRevert(
            abi.encodeWithSelector(BlobsLib.KZGInputBadProofLength.selector, badLength, KZG_COMMITMENT_LENGTH)
        );
        helper.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, new bytes(badLength));
    }

    /// @notice Test buildKZGInput with zero-length commitment
    function test_buildKZGInput_ZeroLengthCommitment() public {
        vm.expectRevert(abi.encodeWithSelector(BlobsLib.KZGInputBadCommitmentLength.selector, 0, KZG_COMMITMENT_LENGTH));
        helper.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, new bytes(0), TEST_PROOF);
    }

    /// @notice Test buildKZGInput with oversized commitment
    function test_buildKZGInput_OversizedCommitment() public {
        uint256 badLength = 49;
        vm.expectRevert(
            abi.encodeWithSelector(BlobsLib.KZGInputBadCommitmentLength.selector, badLength, KZG_COMMITMENT_LENGTH)
        );
        helper.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, new bytes(badLength), TEST_PROOF);
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

        for (uint256 i = 0; i < KZG_COMMITMENT_LENGTH; i++) {
            assertEq(proof.commitment[i], TEST_COMMITMENT[i], "Commitment bytes should match");
        }

        for (uint256 i = 0; i < KZG_PROOF_LENGTH; i++) {
            assertEq(proof.proof[i], TEST_PROOF[i], "Proof bytes should match");
        }
    }

    /// @notice Test decodeKZGInput with invalid input length
    function test_decodeKZGInput_InvalidInputLength() public {
        uint256 badLength = 100;
        vm.expectRevert(abi.encodeWithSelector(BlobsLib.KZGInputBadInputLength.selector, badLength, KZG_INPUT_LENGTH));
        helper.decodeKZGInput(new bytes(badLength));
    }

    /// @notice Test decodeKZGInput with empty input
    function test_decodeKZGInput_EmptyInput() public {
        uint256 badLength = 0;
        vm.expectRevert(abi.encodeWithSelector(BlobsLib.KZGInputBadInputLength.selector, badLength, KZG_INPUT_LENGTH));
        helper.decodeKZGInput(new bytes(badLength));
    }

    /// @notice Test decodeKZGInput with oversized input
    function test_decodeKZGInput_OversizedInput() public {
        uint256 badLength = 300;
        vm.expectRevert(abi.encodeWithSelector(BlobsLib.KZGInputBadInputLength.selector, badLength, KZG_INPUT_LENGTH));
        helper.decodeKZGInput(new bytes(badLength));
    }

    /*//////////////////////////////////////////////////////////////
                        INTEGRATION TESTS
    //////////////////////////////////////////////////////////////*/

    /// @notice Test round-trip encoding/decoding of KZG input
    function test_KZGInput_RoundTrip() public pure {
        bytes memory input = BlobsLib.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, TEST_PROOF);

        BlobsLib.KZGProof memory proof = BlobsLib.decodeKZGInput(input);

        bytes memory rebuiltInput =
            BlobsLib.buildKZGInput(proof.versionedHash, proof.z, proof.y, proof.commitment, proof.proof);

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
        idx = bound(idx, 0, 1000);

        bytes32 hash = BlobsLib.blobHash(idx);
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

        assertTrue(gasUsed < 10000, "calcBlobHashV1 should use reasonable gas");
    }

    /// @notice Test gas usage for buildKZGInput
    function test_gas_buildKZGInput() public view {
        uint256 gasBefore = gasleft();
        BlobsLib.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, TEST_PROOF);
        uint256 gasUsed = gasBefore - gasleft();

        assertTrue(gasUsed < 20000, "buildKZGInput should use reasonable gas");
    }

    /// @notice Test gas usage for decodeKZGInput
    function test_gas_decodeKZGInput() public view {
        bytes memory input = BlobsLib.buildKZGInput(TEST_VERSIONED_HASH, TEST_Z, TEST_Y, TEST_COMMITMENT, TEST_PROOF);

        uint256 gasBefore = gasleft();
        BlobsLib.decodeKZGInput(input);
        uint256 gasUsed = gasBefore - gasleft();

        assertTrue(gasUsed < 30000, "decodeKZGInput should use reasonable gas");
    }
}
