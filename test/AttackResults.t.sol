// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {AttackBase, AttackLib} from "./AttackBase.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {PinnedMockZiskVerifier} from "./mocks/MockZiskVerifier.sol";
import {MockDKG} from "./mocks/MockDKG.sol";

/// @dev Results in sequencer mode: who can stop them, when they open, what they read.
contract AttackResultsTest is AttackBase {
    using stdJson for string;

    function _fixtureResults() internal view returns (bytes memory pv, bytes memory proof) {
        pv = fixture.readBytes(".results.public_values");
        proof = fixture.readBytes(".results.proof_bytes");
    }

    function _settled() internal returns (bytes31 pid) {
        pid = _fixtureProcess();
        _settleBoth(pid);
    }

    // BUG: once the grace window has closed a sequencer-mode setProcessResults is valid
    // and carries the tally in calldata, yet the organizer can still cancel: with a
    // public mempool it reads the tally and front-runs the results with a veto.
    function test_OrganizerCannotVetoPublishableResults() public {
        bytes31 pid = _settled();
        _warpToGraceEnd(pid);
        (bytes memory pv, bytes memory proof) = _fixtureResults();

        // The results transaction is pending; the organizer front-runs it.
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);

        registry.setProcessResults(pid, pv, proof);
        assertEq(uint8(_status(pid)), uint8(DAVINCITypes.ProcessStatus.RESULTS));
    }

    // BUG: same veto from PAUSED, at any time from the grace end on; before it the
    // organizer keeps its power to cancel without having seen a tally.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_NoCancelFromTheGraceEnd(bool paused, uint256 at) public {
        bytes31 pid = _settled();
        if (paused) _setStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        uint256 graceEnd = registry.getProcessGraceEnd(pid);
        uint256 t = vm.getBlockTimestamp() + bound(at, 0, graceEnd + 30 days - vm.getBlockTimestamp());
        vm.warp(t);
        vm.prank(ORGANIZER);
        if (t >= graceEnd) vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);
    }

    /// The other status moves past the grace end cannot hide or change a tally.
    function test_StatusMovesPastTheGraceEnd() public {
        bytes31 pid = _settled();
        _setStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        _warpToGraceEnd(pid);
        bytes32 root = registry.getProcess(pid).latestStateRoot;
        _setStatus(pid, DAVINCITypes.ProcessStatus.READY);
        _setStatus(pid, DAVINCITypes.ProcessStatus.ENDED);
        assertEq(registry.getProcess(pid).latestStateRoot, root);
        (bytes memory pv, bytes memory proof) = _fixtureResults();
        registry.setProcessResults(pid, pv, proof);
        assertEq(uint8(_status(pid)), uint8(DAVINCITypes.ProcessStatus.RESULTS));
    }

    /// Before the grace end every results call fails: before the end by time, after it by
    /// the open window.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_ResultsPremature(uint256 at) public {
        bytes31 pid = _settled();
        uint256 end = registry.getProcessEndTime(pid);
        uint256 graceEnd = registry.getProcessGraceEnd(pid);
        uint256 t = vm.getBlockTimestamp() + bound(at, 0, graceEnd - 1 - vm.getBlockTimestamp());
        vm.warp(t);
        (bytes memory pv, bytes memory proof) = _fixtureResults();
        vm.expectRevert(t < end ? IProcessRegistry.InvalidTimeBounds.selector : IProcessRegistry.GraceOpen.selector);
        registry.setProcessResults(pid, pv, proof);
    }

    /// A process with no transitions publishes the zero tally of its genesis root.
    function test_ResultsOfAProcessWithNoTransitions() public {
        bytes31 pid = _fixtureProcess();
        _warpToGraceEnd(pid);
        uint64[16] memory v;
        registry.setProcessResults(pid, _resultsPublics(registry.getProcess(pid).latestStateRoot, v), hex"");
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(uint8(p.status), uint8(DAVINCITypes.ProcessStatus.RESULTS));
        for (uint256 i = 0; i < p.result.length; ++i) {
            assertEq(p.result[i], 0);
        }
    }

    /// Results are only over the latest root: an intermediate or foreign root fails.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_ResultsRootMustBeLatest(bytes32 root) public {
        bytes31 pid = _settled();
        vm.assume(root != registry.getProcess(pid).latestStateRoot);
        _warpToGraceEnd(pid);
        uint64[16] memory v;
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        registry.setProcessResults(pid, _resultsPublics(root, v), hex"");
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        registry.setProcessResults(pid, _resultsPublics(_transition(1).rootBefore, v), hex"");
    }

    /// Once set, results are final: a second proof, a status move or a batch changes nothing.
    function test_ResultsAreFinal() public {
        bytes31 pid = _settled();
        _warpToGraceEnd(pid);
        (bytes memory pv, bytes memory proof) = _fixtureResults();
        registry.setProcessResults(pid, pv, proof);
        uint256[] memory before = registry.getProcess(pid).result;

        uint64[16] memory v;
        v[0] = 999;
        bytes memory other = _resultsPublics(registry.getProcess(pid).latestStateRoot, v);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessResults(pid, other, hex"");
        for (uint8 s = 0; s <= 4; ++s) {
            vm.prank(ORGANIZER);
            vm.expectRevert();
            registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus(s));
        }
        vm.expectRevert(IProcessRegistry.InvalidKeyMode.selector);
        registry.finalizeResultsFromDKG(pid);
        assertEq(registry.getProcess(pid).result, before);
    }

    /// result[i] = lo | hi << 32 over numFields fields; registers past them are ignored.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_ResultsPacking(uint64[16] memory values) public {
        bytes31 pid = _fixtureProcess();
        _warpToGraceEnd(pid);
        registry.setProcessResults(pid, _resultsPublics(registry.getProcess(pid).latestStateRoot, values), hex"");
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.result.length, p.ballotMode.numFields);
        for (uint256 i = 0; i < p.result.length; ++i) {
            assertEq(p.result[i], values[i]);
        }
    }

    // BUG: a registry deployed with the same program key for batches and results accepts
    // any unsettled batch proof as a results proof: with ok = 1 and its root before equal
    // to the latest root, setProcessResults publishes the batch's root-after limbs as the
    // tally. The constructor must refuse the configuration.
    function test_ConstructorRejectsEqualProgramKeys() public {
        PinnedMockZiskVerifier v = new PinnedMockZiskVerifier(BATCH_VK, ROOT_C);
        bytes memory init = abi.encodePacked(
            type(ProcessRegistry).creationCode,
            abi.encode(
                CHAIN_ID,
                address(v),
                BATCH_VK,
                BATCH_VK,
                ROOT_C,
                fixture.readBytes32(".ballot_vk_hash"),
                address(0),
                address(0),
                GRACE,
                GRACE_FLOOR,
                GRACE_CEIL,
                GRACE_MAX_TOTAL,
                NOTICE_MIN,
                address(0)
            )
        );
        address r;
        assembly {
            r := create(0, add(init, 0x20), mload(init))
        }
        if (r == address(0)) return; // refused, as it must be
        _mockOpenings();
        registry = ProcessRegistry(r);
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        _submitSynth(pid, _next(pid, 3, 0, 1));
        // A batch proven on the latest root but never settled, then the window closes.
        AttackLib.Synth memory pending = _next(pid, 5, 0, 2);
        _warpToGraceEnd(pid);
        registry.setProcessResults(pid, pending.pv, hex"");
        assertEq(uint8(_status(pid)), uint8(DAVINCITypes.ProcessStatus.RESULTS));
        assertTrue(false, "a batch proof published the tally");
    }
}

/// @dev Results in the DKG modes against MockDKG: once-only calls, timing and finality.
contract AttackDKGResultsTest is RegistryTestBase {
    using stdJson for string;

    MockDKG internal mock;
    bytes12 internal eid1;
    string internal dkg;

    function _dkgManager() internal override returns (address) {
        mock = new MockDKG();
        return address(mock);
    }

    function setUp() public override {
        super.setUp();
        dkg = vm.readFile("test/vectors/dkg.json");
        uint256[2][] memory keys = new uint256[2][](2);
        keys[0] = _rte(".pool_keys[0]");
        keys[1] = _rte(".pool_keys[1]");
        eid1 = mock.newEpoch(true, keys);
    }

    function _rte(string memory p) internal view returns (uint256[2] memory pt) {
        pt[0] = vm.parseUint(dkg.readString(string.concat(p, ".rte_x")));
        pt[1] = vm.parseUint(dkg.readString(string.concat(p, ".rte_y")));
    }

    function _automatic() internal returns (bytes31) {
        DAVINCITypes.BallotMode memory m;
        m.numFields = uint8(dkg.readUint(".ballot_mode.num_fields"));
        m.costExponent = uint8(dkg.readUint(".ballot_mode.cost_exponent"));
        m.maxValue = dkg.readUint(".ballot_mode.max_value");
        m.maxValueSum = dkg.readUint(".ballot_mode.max_value_sum");
        DAVINCITypes.Census memory c;
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1;
        c.censusRoot = dkg.readBytes32(".census_root");
        c.censusURI = "https://example.com/census";
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        vm.prank(ORGANIZER);
        return registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            0,
            DURATION,
            MAX_VOTERS,
            m,
            c,
            METADATA_URI,
            METADATA_HASH,
            DAVINCITypes.EncryptionKey(0, 0),
            d
        );
    }

    function _inclusion(string memory p)
        internal
        view
        returns (bytes32 root, uint256[64] memory acc, bytes32[] memory s)
    {
        root = dkg.readBytes32(string.concat(p, ".root"));
        bytes32[] memory words = dkg.readBytes32Array(string.concat(p, ".accumulator"));
        for (uint256 i = 0; i < 64; ++i) {
            acc[i] = uint256(words[i]);
        }
        s = dkg.readBytes32Array(string.concat(p, ".siblings"));
    }

    function _forceRoot(bytes31 pid, bytes32 root) internal {
        bytes32 base = keccak256(abi.encode(pid, uint256(1)));
        vm.store(address(registry), bytes32(uint256(base) + 3), root);
        assertEq(registry.getProcess(pid).latestStateRoot, root, "latestStateRoot slot moved");
    }

    /// A settled process with its ciphertexts submitted.
    function _requested() internal returns (bytes31 pid, uint256[64] memory acc, bytes32[] memory s) {
        pid = _automatic();
        bytes32 root;
        (root, acc, s) = _inclusion(".settled");
        _forceRoot(pid, root);
        vm.warp(registry.getProcessGraceEnd(pid));
        registry.requestResultsDecryption(pid, acc, s);
    }

    function test_RequestOnlyOnce() public {
        (bytes31 pid, uint256[64] memory acc, bytes32[] memory s) = _requested();
        vm.expectRevert(IProcessRegistry.ResultsAlreadyRequested.selector);
        registry.requestResultsDecryption(pid, acc, s);
    }

    function test_FinalizeOnlyOnceAndOnlyWhenCombined() public {
        (bytes31 pid,,) = _requested();
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pid);
        mock.setPlaintext(p.dkgEpochId, p.dkgAid, p.dkgFirstIndex, 7);
        // One of two combines done: still not ready.
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pid);
        mock.setPlaintext(p.dkgEpochId, p.dkgAid, p.dkgFirstIndex + 1, 9);
        registry.finalizeResultsFromDKG(pid);
        p = registry.getProcess(pid);
        assertEq(uint8(p.status), uint8(DAVINCITypes.ProcessStatus.RESULTS));
        assertEq(p.result[0], 7);
        assertEq(p.result[1], 0);
        assertEq(p.result[2], 9);
        assertEq(p.result[3], 0);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.finalizeResultsFromDKG(pid);
    }

    function test_FinalizeBeforeRequest() public {
        bytes31 pid = _automatic();
        vm.warp(registry.getProcessGraceEnd(pid));
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pid);
    }

    /// Before the grace end the request fails: by time before the end, by the window after.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_RequestPremature(uint256 at) public {
        bytes31 pid = _automatic();
        (bytes32 root, uint256[64] memory acc, bytes32[] memory s) = _inclusion(".settled");
        _forceRoot(pid, root);
        uint256 end = registry.getProcessEndTime(pid);
        uint256 graceEnd = registry.getProcessGraceEnd(pid);
        uint256 t = vm.getBlockTimestamp() + bound(at, 0, graceEnd - 1 - vm.getBlockTimestamp());
        vm.warp(t);
        vm.expectRevert(t < end ? IProcessRegistry.InvalidTimeBounds.selector : IProcessRegistry.GraceOpen.selector);
        registry.requestResultsDecryption(pid, acc, s);
    }

    /// After the request the process is out of the organizer's hands and closed to batches.
    function test_RequestLocksTheProcess() public {
        (bytes31 pid,,) = _requested();
        assertEq(uint8(registry.getProcess(pid).status), uint8(DAVINCITypes.ProcessStatus.ENDED));
        for (uint8 s = 0; s <= 4; ++s) {
            vm.prank(ORGANIZER);
            vm.expectRevert();
            registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus(s));
        }
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessDuration(pid, DURATION * 2);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.submitStateTransition(pid, new bytes(512), hex"", new bytes[](0), new bytes32[](0), new bytes[](0));
    }

    /// The accumulator of another state root (here genesis, after a settled batch) is refused.
    function test_RequestWithAStaleAccumulator() public {
        bytes31 pid = _automatic();
        (bytes32 root,,) = _inclusion(".settled");
        (, uint256[64] memory acc0, bytes32[] memory s0) = _inclusion(".genesis");
        _forceRoot(pid, root);
        vm.warp(registry.getProcessGraceEnd(pid));
        vm.expectRevert(IProcessRegistry.InvalidInclusionProof.selector);
        registry.requestResultsDecryption(pid, acc0, s0);
    }

    /// Any change to the accumulator or its proof fails the inclusion (or the range check).
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_RequestTamperedInclusion(uint256 w, uint256 delta, uint256 sib, bool tamperSibling) public {
        bytes31 pid = _automatic();
        (bytes32 root, uint256[64] memory acc, bytes32[] memory s) = _inclusion(".settled");
        _forceRoot(pid, root);
        vm.warp(registry.getProcessGraceEnd(pid));
        if (tamperSibling) {
            sib = bound(sib, 0, s.length - 1);
            s[sib] = bytes32(uint256(s[sib]) ^ (1 + bound(delta, 0, type(uint256).max - 1)));
        } else {
            w = bound(w, 0, 63);
            acc[w] = acc[w] ^ (1 + bound(delta, 0, type(uint128).max));
        }
        (bool ok,) = address(registry).call(abi.encodeCall(registry.requestResultsDecryption, (pid, acc, s)));
        assertFalse(ok, "tampered accumulator accepted");
    }
}
