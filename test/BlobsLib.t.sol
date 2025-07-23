// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { BlobsLib } from "../src/libraries/BlobsLib.sol";

contract BlobsLibTest is Test {
    using BlobsLib for *;

    /*//////////////////////////////////////////////////////////////
                            BLOB HASH TESTS
    //////////////////////////////////////////////////////////////*/

    function test_BlobHash_ReturnsZeroForNonExistentIndex() public view {
        // In a non-blob transaction, all indices should return 0
        bytes32 hash = BlobsLib.blobHash(0);
        assertEq(hash, bytes32(0));

        hash = BlobsLib.blobHash(1);
        assertEq(hash, bytes32(0));

        hash = BlobsLib.blobHash(5);
        assertEq(hash, bytes32(0));
    }

    function test_BlobHash_ConsistentResults() public view {
        // Multiple calls should return the same result
        bytes32 hash1 = BlobsLib.blobHash(0);
        bytes32 hash2 = BlobsLib.blobHash(0);
        assertEq(hash1, hash2);
    }

    /*//////////////////////////////////////////////////////////////
                            BLOB BASE FEE TESTS
    //////////////////////////////////////////////////////////////*/

    function test_BlobBaseFee_ReturnsValue() public view {
        // Should return a value (likely 0 in test environment)
        uint256 fee = BlobsLib.blobBaseFee();
        // In test environment, this will likely be 0, but function should not revert
        assertTrue(fee >= 0); // Always true, but ensures no revert
    }

    function test_BlobBaseFee_ConsistentResults() public view {
        // Multiple calls should return the same result
        uint256 fee1 = BlobsLib.blobBaseFee();
        uint256 fee2 = BlobsLib.blobBaseFee();
        assertEq(fee1, fee2);
    }

    /*//////////////////////////////////////////////////////////////
                        BLOB FEE CALCULATION TESTS
    //////////////////////////////////////////////////////////////*/

    function test_CalculateBlobFee_ZeroBlobs() public view {
        uint256 fee = BlobsLib.calculateBlobFee(0);
        assertEq(fee, 0);
    }

    function test_CalculateBlobFee_SingleBlob() public view {
        uint256 fee = BlobsLib.calculateBlobFee(1);
        uint256 expected = 131072; // baseFee * GAS_PER_BLOB
        assertEq(fee, expected);
    }

    function test_CalculateBlobFee_MultipleBlobs() public view {
        uint256 fee = BlobsLib.calculateBlobFee(3);
        uint256 baseFee = BlobsLib.blobBaseFee();
        uint256 expected = baseFee * 3 * 131072;
        assertEq(fee, expected);
    }

    /*//////////////////////////////////////////////////////////////
                            KZG VERIFICATION TESTS
    //////////////////////////////////////////////////////////////*/

    function test_VerifyKZG_InvalidInputLength_TooShort() public {
        bytes memory shortInputMem = new bytes(191); // One byte short

        vm.expectRevert(abi.encodeWithSelector(BlobsLib.InvalidKZGInputLength.selector, 191, 192));
        BlobsLib.verifyKZG(shortInputMem);
    }

    function test_VerifyKZG_InvalidInputLength_TooLong() public {
        bytes memory longInputMem = new bytes(193); // One byte too long

        vm.expectRevert(abi.encodeWithSelector(BlobsLib.InvalidKZGInputLength.selector, 193, 192));
        BlobsLib.verifyKZG(longInputMem);
    }

    function test_VerifyKZG_ValidInputLength_PrecompileFails() public {
        bytes memory validInputMem = new bytes(192);

        // Mock the precompile to fail
        vm.mockCall(address(0x0A), validInputMem, abi.encode(false, ""));

        vm.expectRevert(BlobsLib.KZGVerificationFailed.selector);
        BlobsLib.verifyKZG(validInputMem);
    }

    function test_VerifyKZG_InvalidOutputLength() public {
        bytes memory validInputMem = new bytes(192);
        bytes memory invalidOutput = new bytes(32); // Should be 64 bytes

        // Mock the precompile to return invalid output length
        vm.mockCall(address(0x0A), validInputMem, abi.encode(true, invalidOutput));

        vm.expectRevert(abi.encodeWithSelector(BlobsLib.InvalidKZGOutputLength.selector, 32));
        BlobsLib.verifyKZG(validInputMem);
    }

    function test_VerifyKZG_ValidProof() public {
        bytes memory validInputMem = new bytes(192);
        bytes memory validOutput = new bytes(64);

        // Set first 32 bytes to 1 (successful verification)
        assembly {
            mstore(add(validOutput, 0x20), 1)
        }

        // Mock the precompile to return success
        vm.mockCall(address(0x0A), validInputMem, abi.encode(true, validOutput));

        bool result = BlobsLib.verifyKZG(validInputMem);
        assertTrue(result);
    }

    function test_VerifyKZG_InvalidProof() public {
        bytes memory validInputMem = new bytes(192);
        bytes memory validOutput = new bytes(64);

        // Set first 32 bytes to 0 (failed verification)
        assembly {
            mstore(add(validOutput, 0x20), 0)
        }

        // Mock the precompile to return failure
        vm.mockCall(address(0x0A), validInputMem, abi.encode(true, validOutput));

        bool result = BlobsLib.verifyKZG(validInputMem);
        assertFalse(result);
    }

    /*//////////////////////////////////////////////////////////////
                        TRY VERIFY KZG TESTS
    //////////////////////////////////////////////////////////////*/

    function test_TryVerifyKZG_InvalidInputLength() public view {
        bytes memory shortInput = new bytes(191);
        bool result = BlobsLib.tryVerifyKZG(shortInput);
        assertFalse(result);
    }

    function test_TryVerifyKZG_PrecompileFails() public {
        bytes memory validInput = new bytes(192);

        // Mock the precompile to fail
        vm.mockCall(address(0x0A), validInput, abi.encode(false, ""));

        bool result = BlobsLib.tryVerifyKZG(validInput);
        assertFalse(result);
    }

    function test_TryVerifyKZG_ValidProof() public {
        bytes memory validInput = new bytes(192);
        bytes memory validOutput = new bytes(64);

        // Set first 32 bytes to 1 (successful verification)
        assembly {
            mstore(add(validOutput, 0x20), 1)
        }

        // Mock the precompile to return success
        vm.mockCall(address(0x0A), validInput, abi.encode(true, validOutput));

        bool result = BlobsLib.tryVerifyKZG(validInput);
        assertTrue(result);
    }

    /*//////////////////////////////////////////////////////////////
                        UTILITY FUNCTION TESTS
    //////////////////////////////////////////////////////////////*/

    function test_IsBlobDataAvailable_NoBlobs() public view {
        // In test environment, no blobs are available
        bool available = BlobsLib.isBlobDataAvailable();
        assertFalse(available);
    }

    function test_GetBlobCount_NoBlobs() public view {
        // In test environment, blob count should be 0
        uint256 count = BlobsLib.getBlobCount();
        assertEq(count, 0);
    }

    function test_GetBlobCount_SafetyLimit() public {
        // Mock blob hashes to test safety limit
        for (uint256 i = 0; i < 10; i++) {
            vm.mockCall(
                address(this),
                abi.encodeWithSignature("blobhashes(uint256)", i),
                abi.encode(bytes32(uint256(i + 1)))
            );
        }

        // Should stop at 6 due to safety limit
        uint256 count = BlobsLib.getBlobCount();
        assertEq(count, 6);
    }

    function test_GetAllBlobHashes_NoBlobs() public view {
        bytes32[] memory hashes = BlobsLib.getAllBlobHashes();
        assertEq(hashes.length, 0);
    }

    function test_GetAllBlobHashes_WithBlobs() public {
        // Mock 3 blob hashes
        bytes32 hash1 = keccak256("blob1");
        bytes32 hash2 = keccak256("blob2");
        bytes32 hash3 = keccak256("blob3");

        vm.mockCall(address(this), abi.encodeWithSignature("blobhashes(uint256)", 0), abi.encode(hash1));
        vm.mockCall(address(this), abi.encodeWithSignature("blobhashes(uint256)", 1), abi.encode(hash2));
        vm.mockCall(address(this), abi.encodeWithSignature("blobhashes(uint256)", 2), abi.encode(hash3));
        vm.mockCall(address(this), abi.encodeWithSignature("blobhashes(uint256)", 3), abi.encode(bytes32(0)));

        bytes32[] memory hashes = BlobsLib.getAllBlobHashes();
        assertEq(hashes.length, 3);
        assertEq(hashes[0], hash1);
        assertEq(hashes[1], hash2);
        assertEq(hashes[2], hash3);
    }

    /*//////////////////////////////////////////////////////////////
                            INTEGRATION TESTS
    //////////////////////////////////////////////////////////////*/

    function test_Integration_BlobOperationsFlow() public view {
        // Test the typical flow of blob operations

        // 1. Check if blob data is available
        bool available = BlobsLib.isBlobDataAvailable();

        // 2. Get blob count
        uint256 count = BlobsLib.getBlobCount();

        // 3. If no blobs, both should be consistent
        if (!available) {
            assertEq(count, 0);
        }

        // 4. Get all blob hashes
        bytes32[] memory hashes = BlobsLib.getAllBlobHashes();
        assertEq(hashes.length, count);

        // 5. Calculate blob fee
        uint256 fee = BlobsLib.calculateBlobFee(count);
        if (count == 0) {
            assertEq(fee, 0);
        }
    }

    function test_Integration_KZGVerificationFlow() public {
        // Test both verification methods with invalid input
        bytes memory invalidInput = new bytes(100);

        // tryVerifyKZG should return false for invalid input
        bool tryResult = BlobsLib.tryVerifyKZG(invalidInput);
        assertFalse(tryResult);

        // verifyKZG should revert for invalid input
        vm.expectRevert();
        BlobsLib.verifyKZG(invalidInput);
    }
}
