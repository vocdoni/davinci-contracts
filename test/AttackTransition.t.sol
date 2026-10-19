// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {AttackBase, AttackLib} from "./AttackBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {ZiskVerifier} from "../src/verifiers/ZiskVerifier.sol";
import {MockZiskVerifier} from "./mocks/MockZiskVerifier.sol";

/// @dev Attacks on submitStateTransition: tampered publics against real KZG openings,
///      replays across processes and registries, blob binding, and a differential fuzz of
///      the settlement window against a reference model.
contract AttackTransitionTest is AttackBase {
    using stdJson for string;

    // --- publics the registry validates -----------------------------------------

    /// @dev Registers the registry checks against its own state or the blob arrays (as
    ///      opposed to rootAfter, voters and overwrites, which only the SNARK binds).
    function _validated(uint256 k) internal pure returns (bool) {
        return k <= 9 || (k >= 20 && k <= 36) || k == 42;
    }

    /// @dev reg32 registers read only the low 4 bytes of each 8-byte word.
    function _isReg32(uint256 k) internal pure returns (bool) {
        return (k >= 2 && k <= 9) || (k >= 20 && k <= 35);
    }

    /// Flipping any byte the registry reads in a validated register of the honest first
    /// batch (real openings, accept-all PLONK) must fail settlement.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_ValidatedRegisterTamperReverts(uint256 k, uint256 b, uint8 x) public {
        k = bound(k, 0, 42);
        vm.assume(_validated(k));
        b = _isReg32(k) ? bound(b, 0, 3) : bound(b, 0, 7);
        x = uint8(bound(x, 1, 255));
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        t.publicValues[8 * k + b] = bytes1(uint8(t.publicValues[8 * k + b]) ^ x);
        vm.blobhashes(t.versionedHashes);
        (bool ok,) = address(registry)
            .call(
                abi.encodeCall(
                    registry.submitStateTransition,
                    (pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs)
                )
            );
        assertFalse(ok, "tampered register settled");
        assertEq(registry.getProcess(pid).batchNumber, 0);
    }

    /// The high halves of reg32 words are not read: the guest writes u32 limbs, so only
    /// the SNARK (which hashes all 512 bytes) rejects a nonzero high half. With the
    /// accept-all mock the batch settles with the honest values.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_Reg32HighHalfIgnored(uint256 k, uint256 b, uint8 x) public {
        k = bound(k, 2, 35);
        vm.assume(_isReg32(k));
        b = bound(b, 4, 7);
        x = uint8(bound(x, 1, 255));
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        t.publicValues[8 * k + b] = bytes1(uint8(t.publicValues[8 * k + b]) ^ x);
        _submit(pid, t);
        assertEq(registry.getProcess(pid).latestStateRoot, t.rootAfter);
    }

    /// The batch publics are 512 bytes exactly, whatever is appended or cut.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_PublicsLength(uint256 len) public {
        len = bound(len, 0, 2048);
        vm.assume(len != 512);
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        bytes memory pv = new bytes(len);
        for (uint256 i = 0; i < len && i < 512; ++i) {
            pv[i] = t.publicValues[i];
        }
        vm.blobhashes(t.versionedHashes);
        vm.expectRevert(IProcessRegistry.InvalidPublicValues.selector);
        registry.submitStateTransition(pid, pv, t.proofBytes, t.commitments, t.ys, t.kzgProofs);
    }

    // --- the SNARK binding -------------------------------------------------------

    /// Every byte of the public values is bound by the PLONK public input
    /// sha256(programVK ‖ publicValues ‖ rootC): one flipped byte anywhere fails the real
    /// verifier with the recorded batch proof.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_RealVerifierBindsEveryPublicByte(uint256 i, uint8 x) public {
        (ZiskVerifier v, bytes32 vk, bytes32 rootC, bytes memory pv, bytes memory proof) = _recorded();
        i = bound(i, 0, pv.length - 1);
        x = uint8(bound(x, 1, 255));
        pv[i] = bytes1(uint8(pv[i]) ^ x);
        vm.expectRevert(ZiskVerifier.InvalidProof.selector);
        v.verifySnarkProof(vk, rootC, pv, proof);
    }

    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_RealVerifierBindsProofBytes(uint256 i, uint8 x) public {
        (ZiskVerifier v, bytes32 vk, bytes32 rootC, bytes memory pv, bytes memory proof) = _recorded();
        i = bound(i, 0, 767);
        x = uint8(bound(x, 1, 255));
        proof[i] = bytes1(uint8(proof[i]) ^ x);
        // An invalid point or field element makes the inner verifier return false or
        // revert; either way the call must not succeed.
        (bool ok,) = address(v).staticcall(abi.encodeCall(v.verifySnarkProof, (vk, rootC, pv, proof)));
        assertFalse(ok, "mutated proof verified");
    }

    /// A scalar of the proof shifted by the field modulus is refused, not reduced.
    function test_RealVerifierRejectsNonCanonicalEvaluation() public {
        (ZiskVerifier v, bytes32 vk, bytes32 rootC, bytes memory pv, bytes memory proof) = _recorded();
        uint256[24] memory p = abi.decode(proof, (uint256[24]));
        uint256 q = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
        for (uint256 j = 18; j < 24; ++j) {
            uint256[24] memory m = p;
            if (m[j] > type(uint256).max - q) continue;
            m[j] += q;
            (bool ok,) = address(v).staticcall(abi.encodeCall(v.verifySnarkProof, (vk, rootC, pv, abi.encode(m))));
            assertFalse(ok, "non-canonical evaluation verified");
        }
    }

    /// The program key and the setup root are both bound.
    function testFuzz_RealVerifierBindsKeys(bytes32 vk2, bytes32 rootC2) public {
        (ZiskVerifier v, bytes32 vk, bytes32 rootC, bytes memory pv, bytes memory proof) = _recorded();
        vm.assume(vk2 != vk || rootC2 != rootC);
        vm.expectRevert(ZiskVerifier.InvalidProof.selector);
        v.verifySnarkProof(vk2, rootC2, pv, proof);
    }

    function _recorded()
        internal
        returns (ZiskVerifier v, bytes32 vk, bytes32 rootC, bytes memory pv, bytes memory proof)
    {
        v = new ZiskVerifier();
        string memory s = vm.readFile("test/vectors/recorded_batch_snark.json");
        vk = s.readBytes32(".program_vk");
        rootC = s.readBytes32(".root_c_vadcop_final");
        pv = s.readBytes(".public_values");
        proof = s.readBytes(".proof_bytes");
        v.verifySnarkProof(vk, rootC, pv, proof);
    }

    // --- replay across processes and registries --------------------------------

    /// A transition of one process cannot settle on another with the same configuration:
    /// the genesis root binds the process id, so the roots never coincide.
    function test_TransitionOfOneProcessRefusedOnAnother() public {
        bytes31 a = _fixtureProcess();
        bytes31 b = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        assertTrue(registry.getProcess(a).latestStateRoot != registry.getProcess(b).latestStateRoot);
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        _submit(b, _transition(0));
        _submit(a, _transition(0));
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        _submit(b, _transition(0));
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        _submit(b, _transition(1));
    }

    /// The same organizer's first process on a second registry (same chain id, keys and
    /// config) has another id and genesis, so the fixture transitions replay on neither.
    function test_TransitionRefusedOnAnotherRegistry() public {
        bytes31 a = _fixtureProcess();
        ProcessRegistry r2 = new ProcessRegistry(
            CHAIN_ID,
            address(new MockZiskVerifier()),
            BATCH_VK,
            RESULTS_VK,
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
        );
        vm.prank(ORGANIZER);
        bytes31 b = r2.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            0,
            DURATION,
            MAX_VOTERS,
            _ballotMode(),
            _census(),
            METADATA_URI,
            METADATA_HASH,
            _encKey(),
            _noDkg()
        );
        assertTrue(a != b, "same id on two registries");
        assertTrue(r2.getProcess(b).latestStateRoot != registry.getProcess(a).latestStateRoot);

        Transition memory t = _transition(0);
        vm.blobhashes(t.versionedHashes);
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        r2.submitStateTransition(b, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs);
        vm.expectRevert(IProcessRegistry.UnknownProcessIdPrefix.selector);
        r2.submitStateTransition(a, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs);
    }

    /// Once a root is passed, no batch built on it settles again (replay and fork).
    function test_RefusedOnAnySupersededRoot() public {
        bytes31 pid = _fixtureProcess();
        _settleBoth(pid);
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        _submit(pid, _transition(0));
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        _submit(pid, _transition(1));
    }

    // --- blob binding (real openings) -------------------------------------------

    /// The blobs of the transaction are the transition's, in order: swapping the two
    /// versioned hashes of the second batch fails the opening of blob 0.
    function test_BlobHashesSwapped() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        Transition memory t = _transition(1);
        (t.versionedHashes[0], t.versionedHashes[1]) = (t.versionedHashes[1], t.versionedHashes[0]);
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 0));
        _submit(pid, t);
    }

    /// The same blob twice in the transaction cannot stand for both.
    function test_BlobHashDuplicated() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        Transition memory t = _transition(1);
        t.versionedHashes[1] = t.versionedHashes[0];
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 1));
        _submit(pid, t);
    }

    /// A blob of the transaction with a commitment the guest did not commit to: the pair
    /// digest fails before any opening.
    function test_ForeignCommitmentWithItsOwnOpening() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        Transition memory t = _transition(1);
        Transition memory t0 = _transition(0);
        t.commitments[1] = t0.commitments[0];
        t.ys[1] = t0.ys[0];
        t.kzgProofs[1] = t0.kzgProofs[0];
        t.versionedHashes[1] = t0.versionedHashes[0];
        vm.expectRevert(IProcessRegistry.InvalidBlobsDigest.selector);
        _submit(pid, t);
    }

    /// The one-blob batch with its array entries duplicated (n_blobs still 1).
    function test_ArraysLongerThanNBlobs() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        bytes[] memory c = new bytes[](2);
        (c[0], c[1]) = (t.commitments[0], t.commitments[0]);
        bytes32[] memory y = new bytes32[](2);
        (y[0], y[1]) = (t.ys[0], t.ys[0]);
        bytes[] memory k = new bytes[](2);
        (k[0], k[1]) = (t.kzgProofs[0], t.kzgProofs[0]);
        vm.blobhashes(t.versionedHashes);
        vm.expectRevert(IProcessRegistry.BlobCountMismatch.selector);
        registry.submitStateTransition(pid, t.publicValues, t.proofBytes, c, y, k);
    }

    // --- voters and maxVoters (synthetic) ----------------------------------------

    /// votersCount + votes - overwrites <= maxVoters, exactly at the boundary, for any split.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_MaxVotersBoundary(uint64 maxVoters, uint64 first, uint64 voters, uint64 overwrites) public {
        _mockOpenings();
        maxVoters = uint64(bound(maxVoters, 1, 1_000_000));
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, maxVoters);
        first = uint64(bound(first, 1, maxVoters));
        _submitSynth(pid, _next(pid, first, 0, 1));
        voters = uint64(bound(voters, 1, 2_000_000));
        overwrites = uint64(bound(overwrites, 0, voters));
        // Overwrites can only touch occupied slots; the guest enforces it, not the registry.
        bool fits = uint256(first) + voters - overwrites <= maxVoters;
        bool ok = _trySubmit(pid, _next(pid, voters, overwrites, 2));
        assertEq(ok, fits, "maxVoters rule");
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertLe(p.votersCount, p.maxVoters);
        assertEq(p.votersCount, fits ? uint256(first) + voters - overwrites : first);
    }

    /// More overwrites than votes never settles (no underflowed voter count).
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_OverwritesAboveVotes(uint64 voters, uint64 overwrites) public {
        _mockOpenings();
        voters = uint64(bound(voters, 0, 1000));
        overwrites = uint64(bound(overwrites, uint256(voters) + 1, type(uint64).max));
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        assertFalse(_trySubmit(pid, _next(pid, voters, overwrites, 3)));
    }

    // --- the settlement window, differential ------------------------------------

    /// Model: settles iff status is READY or ENDED, or PAUSED from the end on, and
    /// startTime <= now < end + grace (no transition landed yet, so lastVoteAt = 0).
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_SettlementWindow(uint256 startOff, uint256 dur, uint8 action, uint256 actAt, uint256 at) public {
        _mockOpenings();
        uint256 t0 = vm.getBlockTimestamp();
        uint256 start = t0 + bound(startOff, 0, 2 days);
        uint256 end = start + bound(dur, 1, 3 days);
        bytes31 pid = _process(start, end - start, MAX_VOTERS);
        DAVINCITypes.ProcessStatus st = DAVINCITypes.ProcessStatus.READY;

        // At most one organizer action, at a random time, applied only if it succeeds.
        action = uint8(bound(action, 0, 3));
        uint256 tAct = t0 + bound(actAt, 0, end + 2 * GRACE - t0);
        vm.warp(tAct);
        if (action == 1) {
            vm.prank(ORGANIZER);
            try registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.PAUSED) {
                st = DAVINCITypes.ProcessStatus.PAUSED;
            } catch {}
        } else if (action == 2) {
            vm.prank(ORGANIZER);
            try registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.ENDED) {
                st = DAVINCITypes.ProcessStatus.ENDED;
                if (tAct < end) end = tAct;
            } catch {}
        } else if (action == 3) {
            vm.prank(ORGANIZER);
            try registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED) {
                st = DAVINCITypes.ProcessStatus.CANCELED;
            } catch {}
        }
        assertEq(uint8(_status(pid)), uint8(st), "status model");
        assertEq(registry.getProcessEndTime(pid), end, "end model");
        assertEq(registry.getProcessGraceEnd(pid), end + GRACE, "grace model");

        uint256 t = tAct + bound(at, 0, end + 2 * GRACE - tAct);
        vm.warp(t);
        bool statusOk = st == DAVINCITypes.ProcessStatus.READY || st == DAVINCITypes.ProcessStatus.ENDED
            || (st == DAVINCITypes.ProcessStatus.PAUSED && t >= end);
        bool want = statusOk && t >= start && t < end + GRACE;
        assertEq(_trySubmit(pid, _next(pid, 1, 0, 7)), want, "settlement window");
        if (want) assertEq(registry.getProcess(pid).lastVoteAt, t);
    }

    /// No status or time lets a batch settle on a CANCELED or RESULTS process.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_NothingSettlesAfterCancelOrResults(bool results, uint256 at) public {
        _mockOpenings();
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        if (results) {
            _warpToGraceEnd(pid);
            uint64[16] memory v;
            registry.setProcessResults(pid, _resultsPublics(registry.getProcess(pid).latestStateRoot, v), hex"");
        } else {
            _setStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);
        }
        vm.warp(vm.getBlockTimestamp() + bound(at, 0, 30 days));
        assertFalse(_trySubmit(pid, _next(pid, 1, 0, 9)));
    }
}
