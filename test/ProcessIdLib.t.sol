// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {ProcessIdLib} from "../src/libraries/ProcessIdLib.sol";

contract ProcessIdLibTest is Test {
    function test_ComputeProcessId_Basic() public pure {
        uint32 chainId = 1;
        address creatorAddr = address(0x1234567890123456789012345678901234567890);
        address contractAddr = address(0x000000000000000000000000000000000000dEaD);
        uint64 nonce = 42;

        // Expected prefix = last 4 bytes of keccak256(chainId, contractAddr)
        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 prefix = uint32(uint256(h));

        bytes31 expectedProcessId =
            bytes31(uint248((uint248(uint160(creatorAddr)) << 88) | (uint248(prefix) << 56) | uint248(uint56(nonce))));

        bytes31 processId = ProcessIdLib.computeProcessId(prefix, creatorAddr, nonce);

        // Assert layout
        assertEq(address(uint160(uint248(processId >> 88))), creatorAddr);
        assertEq(uint32(uint248(processId >> 56)), prefix);
        assertEq(uint56(uint248(processId)), uint56(nonce));
        assertEq(processId, expectedProcessId);
    }

    function test_ComputeProcessId_MaxValues() public pure {
        uint32 chainId = type(uint32).max;
        address creatorAddr = address(0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF);
        address contractAddr = address(0x7777777777777777777777777777777777777777);
        uint64 nonce = type(uint64).max;

        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 prefix = uint32(uint256(h));

        bytes31 processId = ProcessIdLib.computeProcessId(prefix, creatorAddr, nonce);

        assertEq(address(uint160(uint248(processId >> 88))), creatorAddr);
        assertEq(uint32(uint248(processId >> 56)), prefix);
        assertEq(uint56(uint248(processId)), type(uint56).max);
    }

    function test_ComputeProcessId_ZeroValues() public pure {
        uint32 chainId = 0;
        address creatorAddr = address(0);
        address contractAddr = address(0);
        uint64 nonce = 0;

        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 prefix = uint32(uint256(h));

        bytes31 processId = ProcessIdLib.computeProcessId(prefix, creatorAddr, nonce);

        assertEq(address(uint160(uint248(processId >> 88))), creatorAddr);
        assertEq(uint32(uint248(processId >> 56)), prefix);
        assertEq(uint56(uint248(processId)), uint56(nonce));
    }

    function test_ComputeProcessId_NonceTruncation() public pure {
        uint32 chainId = 1;
        address creatorAddr = address(0x1234567890123456789012345678901234567890);
        address contractAddr = address(0x7777777777777777777777777777777777777777);
        uint256 largeNonce = type(uint256).max; // low 64 bits are all ones
        uint64 truncatedNonce = type(uint64).max;

        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 prefix = uint32(uint256(h));

        assertEq(largeNonce & type(uint64).max, uint256(truncatedNonce));
        bytes31 processId = ProcessIdLib.computeProcessId(prefix, creatorAddr, truncatedNonce);
        assertEq(uint56(uint248(processId)), type(uint56).max);
    }

    function test_ComputeProcessId_PrefixMatchesHashTail() public pure {
        uint32 chainId = type(uint32).max;
        address creatorAddr = address(0x1234567890123456789012345678901234567890);
        address contractAddr = address(0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF);
        uint64 nonce = 42;

        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 expectedPrefix = uint32(uint256(h));

        bytes31 processId = ProcessIdLib.computeProcessId(expectedPrefix, creatorAddr, nonce);
        assertEq(uint32(uint248(processId >> 56)), expectedPrefix);
    }

    function test_ComputeProcessId_AddressPlacement() public pure {
        uint32 chainId = 1;
        address creatorAddr = address(0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF);
        address contractAddr = address(0x0000000000000000000000000000000000000001);
        uint64 nonce = 42;

        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 prefix = uint32(uint256(h));

        bytes31 processId = ProcessIdLib.computeProcessId(prefix, creatorAddr, nonce);
        assertEq(address(uint160(uint248(processId >> 88))), creatorAddr);
    }

    function test_ComputeProcessId_NoncePlacement() public pure {
        uint32 chainId = 1;
        address creatorAddr = address(0x1234567890123456789012345678901234567890);
        address contractAddr = address(0x0000000000000000000000000000000000000002);
        uint64 nonce = type(uint64).max;

        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 prefix = uint32(uint256(h));

        bytes31 processId = ProcessIdLib.computeProcessId(prefix, creatorAddr, nonce);
        assertEq(uint56(uint248(processId)), type(uint56).max);
    }

    function test_ComputeProcessId_BitManipulation() public pure {
        uint32 chainId = 0x12345678;
        address creatorAddr = address(0x1234567890123456789012345678901234567890);
        address contractAddr = address(0x7777777777777777777777777777777777777777);
        uint64 nonce = 0x1234567890ABCDEF;

        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 prefix = uint32(uint256(h));

        bytes31 processId = ProcessIdLib.computeProcessId(prefix, creatorAddr, nonce);

        // Check the prefix is in bytes 20-23 (bits 95-64)
        // Byte 20 (MSB of prefix) lives at bits [95..88] of processId, etc.
        assertEq(uint248(processId >> 80) & 0xFF, (uint256(prefix) >> 24) & 0xFF);
        assertEq(uint248(processId >> 72) & 0xFF, (uint256(prefix) >> 16) & 0xFF);
        assertEq(uint248(processId >> 64) & 0xFF, (uint256(prefix) >> 8) & 0xFF);
        assertEq(uint248(processId >> 56) & 0xFF, uint256(prefix) & 0xFF);

        // Check the last 7 bytes (nonce) little→big order via shifts
        assertEq(uint8(uint248(processId)), 0xEF);
        assertEq(uint8(uint248(processId >> 8)), 0xCD);
        assertEq(uint8(uint248(processId >> 16)), 0xAB);
        assertEq(uint8(uint248(processId >> 24)), 0x90);
        assertEq(uint8(uint248(processId >> 32)), 0x78);
        assertEq(uint8(uint248(processId >> 40)), 0x56);
        assertEq(uint8(uint248(processId >> 48)), 0x34);
    }

    function test_HasPrefix_True() public pure {
        uint32 chainId = 1;
        address creatorAddr = address(0x1234567890123456789012345678901234567890);
        address contractAddr = address(0x000000000000000000000000000000000000dEaD);
        uint64 nonce = 42;

        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 prefix = uint32(uint256(h));

        bytes31 processId = ProcessIdLib.computeProcessId(prefix, creatorAddr, nonce);
        assertTrue(ProcessIdLib.hasPrefix(processId, ProcessIdLib.getPrefix(chainId, contractAddr)));
    }

    function test_HasPrefix_False() public pure {
        uint32 chainId = 1;
        address creatorAddr = address(0x1234567890123456789012345678901234567890);
        address contractAddr = address(0x000000000000000000000000000000000000dEaD);
        uint64 nonce = 42;

        bytes32 h = keccak256(abi.encodePacked(chainId, contractAddr));
        uint32 prefix = uint32(uint256(h));

        bytes31 processId = ProcessIdLib.computeProcessId(prefix, creatorAddr, nonce);
        assertFalse(ProcessIdLib.hasPrefix(processId, 0xDEADBEEF));
    }
}
