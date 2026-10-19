// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {MockZiskVerifier} from "./mocks/MockZiskVerifier.sol";

/// @dev Exposes ProcessRegistry._point.
contract PointHarness is ProcessRegistry {
    constructor(address verifier, bytes32 ballotVKHash)
        ProcessRegistry(
            1337,
            verifier,
            keccak256("batch"),
            keccak256("results"),
            keccak256("rootC"),
            ballotVKHash,
            address(0),
            address(0),
            180,
            150,
            600,
            1800,
            60,
            address(0)
        )
    {}

    function point(bytes31 pid, bytes32 rootBefore, bytes32 cellsDigest, uint256 n, uint256 i, bytes calldata com)
        external
        pure
        returns (bytes32)
    {
        return _point(pid, rootBefore, cellsDigest, n, i, com);
    }
}

/// @dev The blob evaluation points against the go-sdk (test/vectors/point.json, the same known
///      answers rust-sdk and go-sdk pin) and the fixture's per-blob zs, and that a cells digest
///      other than the proven one fails the openings.
contract PointTest is RegistryTestBase {
    using stdJson for string;

    PointHarness internal harness;
    string internal points;

    function setUp() public override {
        super.setUp();
        points = vm.readFile("test/vectors/point.json");
        harness = new PointHarness(address(new MockZiskVerifier()), fixture.readBytes32(".ballot_vk_hash"));
    }

    function test_Point_MatchesSdkVectors() public view {
        uint256 k;
        for (; vm.keyExistsJson(points, string.concat(".cases[", vm.toString(k), "]")); ++k) {
            string memory p = string.concat(".cases[", vm.toString(k), "]");
            bytes32 z = harness.point(
                bytes31(points.readBytes(string.concat(p, ".process_id"))),
                points.readBytes32(string.concat(p, ".root_before")),
                points.readBytes32(string.concat(p, ".cells_digest")),
                points.readUint(string.concat(p, ".n_blobs")),
                points.readUint(string.concat(p, ".index")),
                points.readBytes(string.concat(p, ".commitment"))
            );
            assertEq(z, points.readBytes32(string.concat(p, ".z")), p);
        }
        assertEq(k, 4);
    }

    function test_Point_MatchesFixtureZs() public view {
        bytes31 pid = bytes31(fixture.readBytes(".process_id"));
        for (uint256 i = 0; i < 2; ++i) {
            string memory p = string.concat(".transitions[", vm.toString(i), "]");
            Transition memory t = _transition(i);
            bytes32 d = fixture.readBytes32(string.concat(p, ".cells_digest"));
            bytes32[] memory zs = fixture.readBytes32Array(string.concat(p, ".zs"));
            assertEq(zs.length, t.nBlobs);
            for (uint256 b = 0; b < t.nBlobs; ++b) {
                assertEq(harness.point(pid, t.rootBefore, d, t.nBlobs, b, t.commitments[b]), zs[b]);
            }
        }
    }

    function test_Point_BindsCellsDigestIndexAndCount() public view {
        // Same process, root and commitment: the content digest, the index and the count each
        // move the point, so content cannot be chosen after the point is known.
        bytes31 pid = bytes31(fixture.readBytes(".process_id"));
        Transition memory t = _transition(1);
        bytes32 d = fixture.readBytes32(".transitions[1].cells_digest");
        bytes memory c = t.commitments[0];
        bytes32 z = harness.point(pid, t.rootBefore, d, 2, 0, c);
        assertTrue(z != harness.point(pid, t.rootBefore, d ^ bytes32(uint256(1)), 2, 0, c));
        assertTrue(z != harness.point(pid, t.rootBefore, d, 2, 1, c));
        assertTrue(z != harness.point(pid, t.rootBefore, d, 3, 0, c));
    }

    function test_RevertWhen_CellsDigestTampered() public {
        // The pair digest still matches; only D differs from the one the guest proved.
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        bytes32 d = fixture.readBytes32(".transitions[0].cells_digest");
        _setReg32(t.publicValues, 46, d ^ bytes32(uint256(1) << 255));
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 0));
        _submit(pid, t);
    }

    function test_RevertWhen_CellsDigestZero() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _setReg32(t.publicValues, 46, bytes32(0));
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 0));
        _submit(pid, t);
    }
}
