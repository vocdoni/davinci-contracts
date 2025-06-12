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

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, nonce);

        assertEq(uint32(uint256(processId >> 224)), chainId);
        assertEq(address(uint160(uint256(processId >> 64))), addr);
        assertEq(uint64(uint256(processId)), nonce);
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

    function test_ExtractNonce_Basic() public {
        uint64 nonce = 42;
        bytes32 processId = bytes32(uint256(nonce));

        assertEq(ProcessIdLib.extractNonce(processId), nonce);
    }

    function test_ExtractNonce_MaxValue() public {
        uint64 maxNonce = type(uint64).max;
        bytes32 processId = bytes32(uint256(maxNonce));

        assertEq(ProcessIdLib.extractNonce(processId), maxNonce);
    }

    function test_ExtractNonce_Zero() public {
        uint64 nonce = 0;
        bytes32 processId = bytes32(uint256(nonce));

        assertEq(ProcessIdLib.extractNonce(processId), nonce);
    }

    function test_ExtractNonce_WithOtherData() public {
        uint32 chainId = 1;
        address addr = address(0x1234567890123456789012345678901234567890);
        uint64 nonce = 42;

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, nonce);
        assertEq(ProcessIdLib.extractNonce(processId), nonce);
    }

    function test_ExtractAddress_Basic() public {
        address addr = address(0x1234567890123456789012345678901234567890);
        bytes32 processId = bytes32(uint256(uint160(addr)) << 64);

        assertEq(ProcessIdLib.extractAddress(processId), addr);
    }

    function test_ExtractAddress_MaxValue() public {
        address maxAddr = address(0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF);
        bytes32 processId = bytes32(uint256(uint160(maxAddr)) << 64);

        assertEq(ProcessIdLib.extractAddress(processId), maxAddr);
    }

    function test_ExtractAddress_Zero() public {
        address addr = address(0);
        bytes32 processId = bytes32(uint256(uint160(addr)) << 64);

        assertEq(ProcessIdLib.extractAddress(processId), addr);
    }

    function test_ExtractAddress_WithOtherData() public {
        uint32 chainId = 1;
        address addr = address(0x1234567890123456789012345678901234567890);
        uint64 nonce = 42;

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, nonce);
        assertEq(ProcessIdLib.extractAddress(processId), addr);
    }

    function test_Roundtrip_AllComponents() public {
        uint32 chainId = 1;
        address addr = address(0x1234567890123456789012345678901234567890);
        uint64 nonce = 42;

        bytes32 processId = ProcessIdLib.computeProcessId(chainId, addr, nonce);

        assertEq(uint32(uint256(processId >> 224)), chainId);
        assertEq(ProcessIdLib.extractAddress(processId), addr);
        assertEq(ProcessIdLib.extractNonce(processId), nonce);
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
