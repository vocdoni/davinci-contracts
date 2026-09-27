// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {PinnedMockZiskVerifier} from "./mocks/MockZiskVerifier.sol";

/// @dev Settlement of zkVM state transitions against real KZG openings; the PLONK
///      verifier is mocked.
contract TransitionTest is RegistryTestBase {
    using stdJson for string;

    function test_FixtureRootBeforeIsGenesis() public {
        bytes31 pid = _fixtureProcess();
        bytes32 genesis = fixture.readBytes32(".genesis_root");
        assertEq(registry.getProcess(pid).latestStateRoot, genesis);
        assertEq(_transition(0).rootBefore, genesis);
    }

    function test_SubmitStateTransition_OneBlob() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        assertEq(t.nBlobs, 1);

        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessStateTransitioned(pid, address(this), t.rootBefore, t.rootAfter, 3, 0, 1);
        vm.blobhashes(t.versionedHashes);
        uint256 g = gasleft();
        registry.submitStateTransition(pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs);
        emit log_named_uint("submitStateTransition gas, 1 blob (mock verifier):", g - gasleft());

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.latestStateRoot, t.rootAfter);
        assertEq(p.votersCount, 3);
        assertEq(p.overwrittenVotesCount, 0);
        assertEq(p.batchNumber, 1);
    }

    function test_SubmitStateTransition_TwoBatchesTwoBlobs() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));

        Transition memory t = _transition(1);
        assertEq(t.nBlobs, 2);
        assertEq(t.overwrites, 3);
        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessStateTransitioned(pid, address(this), t.rootBefore, t.rootAfter, 820, 3, 2);
        vm.blobhashes(t.versionedHashes);
        uint256 g = gasleft();
        registry.submitStateTransition(pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs);
        emit log_named_uint("submitStateTransition gas, 2 blobs (mock verifier):", g - gasleft());

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.latestStateRoot, t.rootAfter);
        // 3 fresh votes, then 820 votes of which 3 overwrite the first batch.
        assertEq(p.votersCount, 820);
        assertEq(p.overwrittenVotesCount, 3);
        assertEq(p.batchNumber, 2);
    }

    function test_SubmitStateTransition_Permissionless() public {
        bytes31 pid = _fixtureProcess();
        vm.prank(address(0xBEEF));
        _submit(pid, _transition(0));
        assertEq(registry.getProcess(pid).batchNumber, 1);
    }

    function test_SubmitStateTransition_MaxVotersBoundary() public {
        bytes31 pid = _fixtureProcess(3);
        _submit(pid, _transition(0));
        assertEq(registry.getProcess(pid).votersCount, 3);
    }

    // --- reverts ---------------------------------------------------------------

    function test_RevertWhen_NotOk() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _setWord(t.publicValues, 0, 0);
        vm.expectRevert(IProcessRegistry.CircuitFailed.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_FailMaskSet() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _setWord(t.publicValues, 1, uint64(1) << 18);
        vm.expectRevert(IProcessRegistry.CircuitFailed.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_PublicValuesLength() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        t.publicValues = bytes.concat(t.publicValues, hex"00");
        vm.expectRevert(IProcessRegistry.InvalidPublicValues.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_WrongRootBefore() public {
        bytes31 pid = _fixtureProcess();
        // The second batch starts from the first batch's root.
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        _submit(pid, _transition(1));
    }

    function test_RevertWhen_Replayed() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        _submit(pid, _transition(0));
    }

    function test_RevertWhen_WrongCensusRoot() public {
        DAVINCITypes.Census memory c = _census();
        c.censusRoot = bytes32(uint256(c.censusRoot) + 1);
        bytes31 pid = _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), c);
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, _transition(0));
    }

    function test_RevertWhen_CensusRootByteOrderSwapped() public {
        // The guest publishes the census root as LE limbs; the registry stores the BE integer.
        DAVINCITypes.Census memory c = _census();
        c.censusRoot = _reverse(c.censusRoot);
        bytes31 pid = _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), c);
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, _transition(0));
    }

    function test_RevertWhen_WrongOccupiedBefore() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _setWord(t.publicValues, 42, 1);
        vm.expectRevert(IProcessRegistry.InvalidOccupiedBefore.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_OccupiedBeforeCountsOverwrites() public {
        // occupied_before is the distinct slot count (votersCount), not votes + overwrites.
        bytes31 pid = _fixtureProcess();
        _settleBoth(pid);
        Transition memory t = _transition(1);
        _setReg32(t.publicValues, 2, t.rootAfter);
        _setWord(t.publicValues, 42, 823);
        vm.expectRevert(IProcessRegistry.InvalidOccupiedBefore.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_DigestMismatch_Y() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        t.ys[0] = bytes32(uint256(t.ys[0]) ^ 1);
        vm.expectRevert(IProcessRegistry.InvalidBlobsDigest.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_DigestMismatch_BlobOrder() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        Transition memory t = _transition(1);
        (t.commitments[0], t.commitments[1]) = (t.commitments[1], t.commitments[0]);
        (t.ys[0], t.ys[1]) = (t.ys[1], t.ys[0]);
        (t.kzgProofs[0], t.kzgProofs[1]) = (t.kzgProofs[1], t.kzgProofs[0]);
        vm.expectRevert(IProcessRegistry.InvalidBlobsDigest.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_DigestMismatch_Register() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _setWord(t.publicValues, 28, 0);
        vm.expectRevert(IProcessRegistry.InvalidBlobsDigest.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_BadOpening() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        Transition memory t = _transition(1);
        // A valid G1 point, but the opening proof of the other blob.
        t.kzgProofs[1] = t.kzgProofs[0];
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 1));
        _submit(pid, t);
    }

    function test_RevertWhen_OpeningAtAnotherPoint() public {
        // Same blob and digest replayed on top of the first batch: z binds the root
        // before, so the recorded openings no longer verify.
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        Transition memory t = _transition(0);
        _setReg32(t.publicValues, 2, t.rootAfter);
        _setWord(t.publicValues, 42, 3);
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 0));
        _submit(pid, t);
    }

    function test_RevertWhen_WrongBlobInTx() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        t.versionedHashes[0] = _transition(1).versionedHashes[0];
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 0));
        _submit(pid, t);
    }

    function test_RevertWhen_MissingBlob() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        t.versionedHashes = new bytes32[](0);
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.MissingBlob.selector, 0));
        _submit(pid, t);
    }

    function test_RevertWhen_SecondBlobMissing() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        Transition memory t = _transition(1);
        bytes32[] memory only = new bytes32[](1);
        only[0] = t.versionedHashes[0];
        t.versionedHashes = only;
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.MissingBlob.selector, 1));
        _submit(pid, t);
    }

    function test_RevertWhen_NoBlobs() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _setWord(t.publicValues, 36, 0);
        t.commitments = new bytes[](0);
        t.ys = new bytes32[](0);
        t.kzgProofs = new bytes[](0);
        vm.expectRevert(IProcessRegistry.NoBlobs.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_ArrayLengthsDiffer() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        t.ys = new bytes32[](0);
        vm.expectRevert(IProcessRegistry.BlobCountMismatch.selector);
        _submit(pid, t);

        t = _transition(0);
        t.kzgProofs = new bytes[](2);
        vm.expectRevert(IProcessRegistry.BlobCountMismatch.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_BlobCountDiffersFromRegister() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _setWord(t.publicValues, 36, 2);
        vm.expectRevert(IProcessRegistry.BlobCountMismatch.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_CommitmentLength() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        t.commitments[0] = bytes.concat(t.commitments[0], hex"00");
        vm.expectRevert(IProcessRegistry.InvalidBlobCommitmentLength.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_KZGProofLength() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        t.kzgProofs[0] = hex"00";
        vm.expectRevert(IProcessRegistry.InvalidKZGProofLength.selector);
        _submit(pid, t);
    }

    function test_RevertWhen_BeforeStart() public {
        bytes31 pid = _newProcess(block.timestamp + 100, DURATION, MAX_VOTERS, _ballotMode(), _census());
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        _submit(pid, _transition(0));
    }

    function test_RevertWhen_AfterEnd() public {
        bytes31 pid = _fixtureProcess();
        vm.warp(block.timestamp + DURATION);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        _submit(pid, _transition(0));
    }

    function test_RevertWhen_Paused() public {
        bytes31 pid = _fixtureProcess();
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        _submit(pid, _transition(0));
    }

    function test_RevertWhen_Ended() public {
        bytes31 pid = _fixtureProcess();
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.ENDED);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        _submit(pid, _transition(0));
    }

    function test_RevertWhen_Canceled() public {
        bytes31 pid = _fixtureProcess();
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        _submit(pid, _transition(0));
    }

    function test_RevertWhen_MaxVotersExceeded() public {
        bytes31 pid = _fixtureProcess(2);
        vm.expectRevert(IProcessRegistry.MaxVotersReached.selector);
        _submit(pid, _transition(0));
    }

    function test_RevertWhen_MaxVotersExceededBySecondBatch() public {
        bytes31 pid = _fixtureProcess(819);
        _submit(pid, _transition(0));
        vm.expectRevert(IProcessRegistry.MaxVotersReached.selector);
        _submit(pid, _transition(1));
    }

    function test_RevertWhen_UnknownProcess() public {
        Transition memory t = _transition(0);
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        registry.submitStateTransition(bytes31(0), t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs);

        vm.expectRevert(IProcessRegistry.UnknownProcessIdPrefix.selector);
        registry.submitStateTransition(
            bytes31(uint248(1)), t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs
        );

        // Right prefix, never created.
        bytes31 pid = bytes31(fixture.readBytes(".process_id"));
        vm.expectRevert(IProcessRegistry.ProcessNotFound.selector);
        registry.submitStateTransition(pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs);
    }

    function _reverse(bytes32 x) internal pure returns (bytes32 out) {
        for (uint256 i = 0; i < 32; i++) {
            out |= bytes32(uint256(uint8(x[i])) << (8 * i));
        }
    }
}

/// @dev The registry must hand the verifier its pinned batch vk and rootC.
contract TransitionPinnedVerifierTest is RegistryTestBase {
    function _verifier() internal override returns (address) {
        return address(new PinnedMockZiskVerifier(BATCH_VK, ROOT_C));
    }

    function test_SubmitStateTransition_UsesBatchProgramVK() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        assertEq(registry.getProcess(pid).batchNumber, 1);
    }
}

/// @dev A proof of another program (here the results guest) must not settle a transition.
contract TransitionWrongProgramTest is RegistryTestBase {
    function _verifier() internal override returns (address) {
        return address(new PinnedMockZiskVerifier(RESULTS_VK, ROOT_C));
    }

    function test_RevertWhen_WrongProgramVK() public {
        bytes31 pid = _fixtureProcess();
        vm.expectRevert(PinnedMockZiskVerifier.InvalidProof.selector);
        _submit(pid, _transition(0));
    }
}
