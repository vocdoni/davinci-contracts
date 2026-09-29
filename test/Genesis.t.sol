// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {GenesisLib} from "../src/libraries/GenesisLib.sol";
import {Sha256SmtLib} from "../src/libraries/Sha256SmtLib.sol";
import {MockZiskVerifier} from "./mocks/MockZiskVerifier.sol";

contract GenesisHarness {
    function root(uint64[] calldata keys, uint256[] calldata values) external pure returns (bytes32) {
        return Sha256SmtLib.root(keys, values);
    }

    function leafHash(uint64 key, uint256 value) external pure returns (bytes32) {
        return Sha256SmtLib.leafHash(key, value);
    }

    function packBallotMode(DAVINCITypes.BallotMode calldata m) external pure returns (uint256) {
        return GenesisLib.packBallotMode(m);
    }

    function encKeyLeaf(DAVINCITypes.EncryptionKey calldata k) external pure returns (uint256) {
        return GenesisLib.encKeyLeaf(k);
    }
}

/// @dev genesisRoot against test/vectors/genesis.json: go-sdk chain.NewState roots for
///      registry-style pids (three ballot modes, origins 1 and 4) plus the rust-sdk
///      vector, which the generator recomputes and checks before including it, and
///      the dynamic origins 2 and 3 last.
contract GenesisTest is Test {
    using stdJson for string;

    ProcessRegistry internal registry;
    GenesisHarness internal harness;
    string internal vectors;

    function setUp() public {
        vectors = vm.readFile("test/vectors/genesis.json");
        registry = new ProcessRegistry(
            1337,
            address(new MockZiskVerifier()),
            bytes32(uint256(1)),
            bytes32(uint256(2)),
            bytes32(uint256(3)),
            vectors.readBytes32(".ballot_vk_hash"),
            address(0),
            180,
            150,
            600,
            1800,
            60
        );
        harness = new GenesisHarness();
    }

    function _case(uint256 i)
        internal
        view
        returns (
            bytes31 pid,
            DAVINCITypes.BallotMode memory m,
            DAVINCITypes.EncryptionKey memory k,
            DAVINCITypes.CensusOrigin origin,
            string memory p
        )
    {
        p = string.concat(".cases[", vm.toString(i), "]");
        pid = bytes31(vectors.readBytes(string.concat(p, ".process_id")));
        string memory bm = string.concat(p, ".ballot_mode");
        m.numFields = uint8(vectors.readUint(string.concat(bm, ".num_fields")));
        m.groupSize = uint8(vectors.readUint(string.concat(bm, ".group_size")));
        m.uniqueValues = vectors.readBool(string.concat(bm, ".unique_values"));
        m.costExponent = uint8(vectors.readUint(string.concat(bm, ".cost_exponent")));
        m.maxValue = vectors.readUint(string.concat(bm, ".max_value"));
        m.minValue = vectors.readUint(string.concat(bm, ".min_value"));
        m.maxValueSum = vectors.readUint(string.concat(bm, ".max_value_sum"));
        m.minValueSum = vectors.readUint(string.concat(bm, ".min_value_sum"));
        k.x = vm.parseUint(vectors.readString(string.concat(p, ".enc_key.x")));
        k.y = vm.parseUint(vectors.readString(string.concat(p, ".enc_key.y")));
        origin = DAVINCITypes.CensusOrigin(vectors.readUint(string.concat(p, ".census_origin")));
    }

    function test_GenesisRoot_MatchesGoVectors() public view {
        uint256 n;
        uint256 origins;
        for (uint256 i = 0; vm.keyExistsJson(vectors, string.concat(".cases[", vm.toString(i), "]")); i++) {
            (
                bytes31 pid,
                DAVINCITypes.BallotMode memory m,
                DAVINCITypes.EncryptionKey memory k,
                DAVINCITypes.CensusOrigin origin,
                string memory p
            ) = _case(i);
            assertEq(
                registry.genesisRoot(pid, m, k, origin),
                vectors.readBytes32(string.concat(p, ".root")),
                string.concat("genesis root, case ", vm.toString(i))
            );
            origins |= 1 << uint8(origin);
            n++;
        }
        assertEq(n, 18, "case count");
        assertEq(origins, (1 << 1) | (1 << 2) | (1 << 3) | (1 << 4), "every origin covered");
    }

    function test_GenesisLeaves_MatchGoVectors() public view {
        for (uint256 i = 0; vm.keyExistsJson(vectors, string.concat(".cases[", vm.toString(i), "]")); i++) {
            (
                bytes31 pid,
                DAVINCITypes.BallotMode memory m,
                DAVINCITypes.EncryptionKey memory k,
                DAVINCITypes.CensusOrigin origin,
                string memory p
            ) = _case(i);
            string memory leaves = string.concat(p, ".leaves");
            assertEq(
                harness.packBallotMode(m), vm.parseUint(vectors.readString(string.concat(p, ".ballot_mode.packed")))
            );
            assertEq(_le(uint256(uint248(pid))), vectors.readBytes32(string.concat(leaves, ".0x00")), "pid leaf");
            assertEq(_le(harness.packBallotMode(m)), vectors.readBytes32(string.concat(leaves, ".0x02")), "mode leaf");
            assertEq(_le(harness.encKeyLeaf(k)), vectors.readBytes32(string.concat(leaves, ".0x03")), "key leaf");
            assertEq(
                _le(GenesisLib.IDENTITY_ACC_LEAF), vectors.readBytes32(string.concat(leaves, ".0x04")), "results leaf"
            );
            assertEq(_le(uint256(origin)), vectors.readBytes32(string.concat(leaves, ".0x06")), "origin leaf");
            assertEq(
                _le(uint256(registry.ballotVKHash())), vectors.readBytes32(string.concat(leaves, ".0x07")), "vk leaf"
            );
        }
    }

    function test_IdentityAccumulatorLeaf() public view {
        // sha256 of the identity ballot [0,1,0,1,...] as 64 BE32 words, read BE.
        bytes memory buf;
        for (uint256 i = 0; i < 64; i++) {
            buf = bytes.concat(buf, bytes32(i % 2));
        }
        assertEq(GenesisLib.IDENTITY_ACC_LEAF, uint256(sha256(buf)));
        assertEq(bytes32(GenesisLib.IDENTITY_ACC_LEAF), vectors.readBytes32(".identity_acc_leaf"));
    }

    function test_Sha256Smt_MatchesArbo() public {
        string memory smt = vm.readFile("test/vectors/smt.json");
        uint256 n;
        for (uint256 i = 0; vm.keyExistsJson(smt, string.concat(".cases[", vm.toString(i), "]")); i++) {
            string memory p = string.concat(".cases[", vm.toString(i), "]");
            uint256[] memory rawKeys = smt.readUintArray(string.concat(p, ".keys"));
            bytes32[] memory rawValues = smt.readBytes32Array(string.concat(p, ".values"));
            uint64[] memory keys = new uint64[](rawKeys.length);
            uint256[] memory values = new uint256[](rawValues.length);
            for (uint256 j = 0; j < keys.length; j++) {
                keys[j] = uint64(rawKeys[j]);
                values[j] = uint256(rawValues[j]);
            }
            uint256 g = gasleft();
            bytes32 got = harness.root(keys, values);
            emit log_named_uint(
                string.concat("smt root gas, ", smt.readString(string.concat(p, ".name"))), g - gasleft()
            );
            assertEq(got, smt.readBytes32(string.concat(p, ".root")), smt.readString(string.concat(p, ".name")));
            n++;
        }
        assertEq(n, 6, "case count");
    }

    function test_Sha256Smt_SingleLeafIsLeafHash() public view {
        uint64[] memory keys = new uint64[](1);
        uint256[] memory values = new uint256[](1);
        keys[0] = 0x10;
        values[0] = 42;
        assertEq(harness.root(keys, values), harness.leafHash(0x10, 42));
        // sha256(key_le8 ‖ value_le32 ‖ 0x01)
        assertEq(harness.leafHash(0x10, 42), sha256(abi.encodePacked(bytes8(0x1000000000000000), _le(42), hex"01")));
    }

    function test_RevertWhen_DuplicateSmtKeys() public {
        uint64[] memory keys = new uint64[](2);
        uint256[] memory values = new uint256[](2);
        keys[0] = 7;
        keys[1] = 7;
        vm.expectRevert(Sha256SmtLib.SmtMaxLevelsReached.selector);
        harness.root(keys, values);
    }

    function _le(uint256 v) internal pure returns (bytes32 out) {
        for (uint256 i = 0; i < 32; i++) {
            out |= bytes32(((v >> (8 * i)) & 0xff) << (8 * (31 - i)));
        }
    }
}
