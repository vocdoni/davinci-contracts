// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {PublicsLib} from "../src/libraries/PublicsLib.sol";

contract PublicsHarness {
    function word(bytes calldata pv, uint256 k) external pure returns (uint256) {
        return PublicsLib.word(pv, k);
    }

    function reg32(bytes calldata pv, uint256 k) external pure returns (bytes32) {
        return PublicsLib.reg32(pv, k);
    }

    function reverse32(bytes32 x) external pure returns (bytes32) {
        return PublicsLib.reverse32(x);
    }
}

/// @dev Register decoding of the 512-byte publicValues, checked against a batch PLONK
///      recorded from the prover (test/vectors/recorded_batch_snark.json) and decoded
///      independently by the Go vectors program (publics.json).
contract PublicsTest is Test {
    using stdJson for string;

    PublicsHarness internal h;
    string internal want;
    bytes internal pv;

    function setUp() public {
        h = new PublicsHarness();
        want = vm.readFile("test/vectors/publics.json");
        pv = vm.readFile("test/vectors/recorded_batch_snark.json").readBytes(".public_values");
    }

    function test_RecordedPublics_Words() public view {
        assertEq(pv.length, 512);
        uint256[] memory words = want.readUintArray(".words");
        assertEq(words.length, 64);
        for (uint256 k = 0; k < 64; k++) {
            assertEq(h.word(pv, k), words[k], string.concat("word ", vm.toString(k)));
        }
        assertEq(h.word(pv, 0), 1, "ok");
        assertEq(h.word(pv, 1), 0, "fail_mask");
        assertEq(h.word(pv, 18), want.readUint(".voters"));
        assertEq(h.word(pv, 19), want.readUint(".overwrites"));
        assertEq(h.word(pv, 36), want.readUint(".n_blobs"));
        assertEq(h.word(pv, 42), want.readUint(".occupied_before"));
    }

    function test_RecordedPublics_Reg32() public view {
        assertEq(h.reg32(pv, 2), want.readBytes32(".root_before"), "root before");
        assertEq(h.reg32(pv, 10), want.readBytes32(".root_after"), "root after");
        assertEq(h.reg32(pv, 20), want.readBytes32(".census_root_le"), "census root limbs");
        assertEq(h.reverse32(h.reg32(pv, 20)), want.readBytes32(".census_root"), "census root integer");
        assertEq(h.reg32(pv, 28), want.readBytes32(".blobs_digest"), "blobs digest");
    }

    function test_FixturePublics() public view {
        string memory fx = vm.readFile("test/vectors/transition.json");
        bytes memory t = fx.readBytes(".transitions[1].public_values");
        assertEq(h.reg32(t, 2), fx.readBytes32(".transitions[1].root_before"));
        assertEq(h.reg32(t, 10), fx.readBytes32(".transitions[1].root_after"));
        assertEq(h.reverse32(h.reg32(t, 20)), fx.readBytes32(".census_root"));
        assertEq(h.reg32(t, 28), fx.readBytes32(".transitions[1].blobs_digest"));
        assertEq(h.word(t, 18), fx.readUint(".transitions[1].voters"));
        assertEq(h.word(t, 19), fx.readUint(".transitions[1].overwrites"));
        assertEq(h.word(t, 36), fx.readUint(".transitions[1].n_blobs"));
        assertEq(h.word(t, 42), fx.readUint(".transitions[1].occupied_before"));
    }

    function testFuzz_Word(bytes32 a, bytes32 noise, uint8 k) public view {
        // Word kk holds a; every other byte is noise that must not leak in.
        bytes memory buf = new bytes(512);
        for (uint256 i = 0; i < 512; i++) {
            buf[i] = noise[i % 32];
        }
        uint256 kk = uint256(k) % 64;
        for (uint256 i = 0; i < 8; i++) {
            buf[8 * kk + i] = a[i];
        }
        uint256 naive;
        for (uint256 i = 0; i < 8; i++) {
            naive |= uint256(uint8(a[i])) << (8 * i);
        }
        assertEq(h.word(buf, kk), naive);
    }

    function testFuzz_Reg32(bytes calldata seed, uint8 k) public view {
        uint256 kk = uint256(k) % 57;
        bytes memory buf = new bytes(512);
        for (uint256 i = 0; i < 512 && i < seed.length; i++) {
            buf[i] = seed[i];
        }
        bytes32 naive;
        for (uint256 j = 0; j < 8; j++) {
            for (uint256 t = 0; t < 4; t++) {
                naive |= bytes32(uint256(uint8(buf[8 * (kk + j) + t])) << (8 * (31 - (4 * j + t))));
            }
        }
        assertEq(h.reg32(buf, kk), naive);
    }

    function testFuzz_Reverse32(bytes32 x) public view {
        bytes32 naive;
        for (uint256 i = 0; i < 32; i++) {
            naive |= bytes32(uint256(uint8(x[i])) << (8 * i));
        }
        assertEq(h.reverse32(x), naive);
        assertEq(h.reverse32(h.reverse32(x)), x);
    }
}
