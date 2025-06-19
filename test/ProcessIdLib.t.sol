// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { ProcessIdLib } from "../src/libraries/ProcessIdLib.sol";

contract ProcessIdLibTest is Test {
    using ProcessIdLib for bytes32;

    function test_ComputeProcessId_Basic() public {
        uint32 chainId = 1;
        address addr = address(0x1234567890123456789012345678901234567890);
        uint64 nonce = 42;

        bytes32 expectedProcessId = 0x000000011234567890123456789012345678901234567890000000000000002a;

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, nonce);

        assertEq(uint32(uint256(processId >> 224)), chainId);
        assertEq(address(uint160(uint256(processId >> 64))), addr);
        assertEq(uint64(uint256(processId)), nonce);
        assertEq(processId, expectedProcessId);
    }

    function test_ComputeProcessId_MaxValues() public {
        uint32 maxChainId = type(uint32).max;
        address maxAddr = address(0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF);
        uint64 maxNonce = type(uint64).max;

        bytes32 processId = ProcessIdLib.computeProcessId(maxChainId, maxAddr, maxNonce);

        assertEq(uint32(uint256(processId >> 224)), maxChainId);
        assertEq(address(uint160(uint256(processId >> 64))), maxAddr);
        assertEq(uint64(uint256(processId)), maxNonce);
    }

    function test_ComputeProcessId_ZeroValues() public {
        uint32 chainId = 0;
        address addr = address(0);
        uint64 nonce = 0;

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, nonce);

        assertEq(uint32(uint256(processId >> 224)), chainId);
        assertEq(address(uint160(uint256(processId >> 64))), addr);
        assertEq(uint64(uint256(processId)), nonce);
    }

    function test_ComputeProcessId_NonceTruncation() public {
        uint32 chainId = 1;
        address addr = address(0x1234567890123456789012345678901234567890);
        uint256 largeNonce = type(uint256).max; // This will be truncated to uint64

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, uint64(largeNonce));
        assertEq(uint64(uint256(processId)), uint64(largeNonce));
    }

    function test_ComputeProcessId_ChainIdOverflow() public {
        uint32 chainId = type(uint32).max;
        address addr = address(0x1234567890123456789012345678901234567890);
        uint64 nonce = 42;

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, nonce);
        assertEq(uint32(uint256(processId >> 224)), chainId);
    }

    function test_ComputeProcessId_AddressOverflow() public {
        uint32 chainId = 1;
        address addr = address(0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF);
        uint64 nonce = 42;

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, nonce);
        assertEq(address(uint160(uint256(processId >> 64))), addr);
    }

    function test_ComputeProcessId_NonceOverflow() public {
        uint32 chainId = 1;
        address addr = address(0x1234567890123456789012345678901234567890);
        uint64 nonce = type(uint64).max;

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, nonce);
        assertEq(uint64(uint256(processId)), nonce);
    }

    function test_ComputeProcessId_BitManipulation() public {
        uint32 chainId = 0x12345678;
        address addr = address(0x1234567890123456789012345678901234567890);
        uint64 nonce = 0x1234567890ABCDEF;

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, nonce);

        assertEq(uint8(uint256(processId >> 248)), 0x12);
        assertEq(uint8(uint256(processId >> 240)), 0x34);
        assertEq(uint8(uint256(processId >> 232)), 0x56);
        assertEq(uint8(uint256(processId >> 224)), 0x78);

        assertEq(uint8(uint256(processId)), 0xEF);
        assertEq(uint8(uint256(processId >> 8)), 0xCD);
        assertEq(uint8(uint256(processId >> 16)), 0xAB);
        assertEq(uint8(uint256(processId >> 24)), 0x90);
        assertEq(uint8(uint256(processId >> 32)), 0x78);
        assertEq(uint8(uint256(processId >> 40)), 0x56);
        assertEq(uint8(uint256(processId >> 48)), 0x34);
        assertEq(uint8(uint256(processId >> 56)), 0x12);
    }
}
