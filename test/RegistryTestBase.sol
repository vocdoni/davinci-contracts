// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {MockZiskVerifier} from "./mocks/MockZiskVerifier.sol";

/// @dev Deploys the registry at the address test/vectors/transition.json was built for
///      (CREATE(DEPLOYER, 0)) and reads the fixture. Regenerate the vectors with
///      `cd test/vectors && go run .`.
abstract contract RegistryTestBase is Test {
    using stdJson for string;

    struct Transition {
        bytes publicValues;
        bytes proofBytes;
        bytes32 rootBefore;
        bytes32 rootAfter;
        uint256 voters;
        uint256 overwrites;
        uint256 nBlobs;
        bytes[] commitments;
        bytes32[] ys;
        bytes[] kzgProofs;
        bytes32[] versionedHashes;
    }

    // Anvil account 0: deploys the registry with nonce 0 and organizes the fixture process.
    address internal constant DEPLOYER = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
    address internal constant ORGANIZER = DEPLOYER;
    uint32 internal constant CHAIN_ID = 1337;
    bytes32 internal constant BATCH_VK = keccak256("davinci batch program vk");
    bytes32 internal constant RESULTS_VK = keccak256("davinci results program vk");
    bytes32 internal constant ROOT_C = 0x05006517b6ccde5da4d890587ba62845b5af8a307c00e87d4b9d05099b16dc80;
    uint256 internal constant DURATION = 1 days;
    // The production grace bounds (DeployAll defaults).
    uint32 internal constant GRACE = 180;
    uint32 internal constant GRACE_FLOOR = 150;
    uint32 internal constant GRACE_CEIL = 600;
    uint32 internal constant GRACE_MAX_TOTAL = 1800;
    uint32 internal constant NOTICE_MIN = 60;
    uint256 internal constant MAX_VOTERS = 10_000;
    uint256 internal constant BN254_P = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    string internal constant METADATA_URI = "ipfs://metadata";
    /// @dev sha256("metadata document"), the document behind METADATA_URI. A literal: a
    ///      sha256() constant is a precompile call at each use and would consume vm.prank.
    bytes32 internal constant METADATA_HASH = 0xad1b97e9a66b961ecf9c5eb649f5d01b6f01316a6db0b4f8632db058ea1bc8d4;

    ProcessRegistry internal registry;
    string internal fixture;

    function setUp() public virtual {
        fixture = vm.readFile("test/vectors/transition.json");
        address verifier = _verifier();
        bytes32 ballotVKHash = fixture.readBytes32(".ballot_vk_hash");
        address dkgManager = _dkgManager();
        (uint32 grace, uint32 floor, uint32 ceil, uint32 maxTotal, uint32 notice) = _timeConfig();
        vm.prank(DEPLOYER);
        registry = new ProcessRegistry(
            CHAIN_ID,
            verifier,
            BATCH_VK,
            RESULTS_VK,
            ROOT_C,
            ballotVKHash,
            dkgManager,
            grace,
            floor,
            ceil,
            maxTotal,
            notice
        );
        assertEq(address(registry), fixture.readAddress(".registry"), "registry address differs from the fixture");
    }

    /// @dev The verifier the registry is deployed with.
    function _verifier() internal virtual returns (address) {
        return address(new MockZiskVerifier());
    }

    /// @dev The DKG manager the registry is deployed with (0 = DKG modes disabled).
    function _dkgManager() internal virtual returns (address) {
        return address(0);
    }

    /// @dev The registry's grace default, floor, ceil and max total, and its notice minimum.
    function _timeConfig() internal pure virtual returns (uint32, uint32, uint32, uint32, uint32) {
        return (GRACE, GRACE_FLOOR, GRACE_CEIL, GRACE_MAX_TOTAL, NOTICE_MIN);
    }

    /// @dev Warps to the grace end of a process, where transitions stop and results open.
    function _warpToGraceEnd(bytes31 pid) internal {
        vm.warp(registry.getProcessGraceEnd(pid));
    }

    /// @dev All-zero DKGParams: SEQUENCER mode.
    function _noDkg() internal pure returns (DAVINCITypes.DKGParams memory d) {}

    function _ballotMode() internal view returns (DAVINCITypes.BallotMode memory m) {
        m.numFields = uint8(fixture.readUint(".ballot_mode.num_fields"));
        m.groupSize = uint8(fixture.readUint(".ballot_mode.group_size"));
        m.uniqueValues = fixture.readBool(".ballot_mode.unique_values");
        m.costExponent = uint8(fixture.readUint(".ballot_mode.cost_exponent"));
        m.maxValue = fixture.readUint(".ballot_mode.max_value");
        m.minValue = fixture.readUint(".ballot_mode.min_value");
        m.maxValueSum = fixture.readUint(".ballot_mode.max_value_sum");
        m.minValueSum = fixture.readUint(".ballot_mode.min_value_sum");
    }

    function _encKey() internal view returns (DAVINCITypes.EncryptionKey memory k) {
        k.x = vm.parseUint(fixture.readString(".enc_key.x"));
        k.y = vm.parseUint(fixture.readString(".enc_key.y"));
    }

    function _census() internal view returns (DAVINCITypes.Census memory c) {
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1;
        c.censusRoot = fixture.readBytes32(".census_root");
        c.censusURI = "https://example.com/census";
    }

    function _newProcess(
        uint256 startTime,
        uint256 duration,
        uint256 maxVoters,
        DAVINCITypes.BallotMode memory mode,
        DAVINCITypes.Census memory census
    ) internal returns (bytes31 pid) {
        DAVINCITypes.EncryptionKey memory key = _encKey();
        vm.prank(ORGANIZER);
        pid = registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            startTime,
            duration,
            maxVoters,
            mode,
            census,
            METADATA_URI,
            METADATA_HASH,
            key,
            _noDkg()
        );
    }

    /// @dev The organizer's first process, which the fixture transitions were built for.
    function _fixtureProcess(uint256 maxVoters) internal returns (bytes31 pid) {
        pid = _newProcess(block.timestamp, DURATION, maxVoters, _ballotMode(), _census());
        assertEq(bytes32(pid), bytes32(fixture.readBytes(".process_id")), "fixture process id");
    }

    function _fixtureProcess() internal returns (bytes31) {
        return _fixtureProcess(MAX_VOTERS);
    }

    function _transition(uint256 i) internal view returns (Transition memory t) {
        return _transitionAt(string.concat(".transitions[", vm.toString(i), "]"));
    }

    /// @dev The transition at JSON path p of the fixture.
    function _transitionAt(string memory p) internal view returns (Transition memory t) {
        t.publicValues = fixture.readBytes(string.concat(p, ".public_values"));
        t.proofBytes = fixture.readBytes(string.concat(p, ".proof_bytes"));
        t.rootBefore = fixture.readBytes32(string.concat(p, ".root_before"));
        t.rootAfter = fixture.readBytes32(string.concat(p, ".root_after"));
        t.voters = fixture.readUint(string.concat(p, ".voters"));
        t.overwrites = fixture.readUint(string.concat(p, ".overwrites"));
        t.nBlobs = fixture.readUint(string.concat(p, ".n_blobs"));
        t.commitments = fixture.readBytesArray(string.concat(p, ".commitments"));
        t.ys = fixture.readBytes32Array(string.concat(p, ".ys"));
        t.kzgProofs = fixture.readBytesArray(string.concat(p, ".kzg_proofs"));
        t.versionedHashes = fixture.readBytes32Array(string.concat(p, ".versioned_hashes"));
    }

    /// @dev Attaches the transition's blobs to the tx and submits it.
    function _submit(bytes31 pid, Transition memory t) internal {
        vm.blobhashes(t.versionedHashes);
        registry.submitStateTransition(pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs);
    }

    function _settleBoth(bytes31 pid) internal {
        _submit(pid, _transition(0));
        _submit(pid, _transition(1));
    }

    /// @dev Sets register k to v (8-byte LE word).
    function _setWord(bytes memory pv, uint256 k, uint64 v) internal pure {
        for (uint256 i = 0; i < 8; i++) {
            pv[8 * k + i] = bytes1(uint8(v >> (8 * i)));
        }
    }

    /// @dev Sets registers k..k+7 so that reg32(k) == v.
    function _setReg32(bytes memory pv, uint256 k, bytes32 v) internal pure {
        for (uint256 j = 0; j < 8; j++) {
            for (uint256 t = 0; t < 4; t++) {
                pv[8 * (k + j) + t] = v[4 * j + t];
            }
        }
    }

    /// @dev Results-guest publics: ok, fail_mask, state root, 16 results as (lo, hi) registers.
    function _resultsPublics(bytes32 root, uint64[16] memory values) internal pure returns (bytes memory pv) {
        pv = new bytes(512);
        _setWord(pv, 0, 1);
        _setReg32(pv, 2, root);
        for (uint256 i = 0; i < 16; i++) {
            _setWord(pv, 10 + 2 * i, uint64(uint32(values[i])));
            _setWord(pv, 11 + 2 * i, values[i] >> 32);
        }
        _setWord(pv, 42, 0xFFFFFFFF);
    }

    function _copy(bytes memory b) internal pure returns (bytes memory) {
        return bytes.concat(b);
    }
}
