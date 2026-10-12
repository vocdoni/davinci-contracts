// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {RegistryTestBase} from "./RegistryTestBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {IL1BlobVerifier} from "../src/interfaces/IL1BlobVerifier.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {MockL1BlobVerifier} from "./mocks/MockL1BlobVerifier.sol";

/// @dev L1-blob mode: the blob hashes come from the (mocked) L1 verifier instead of BLOBHASH.
///      The registry address does not depend on the mock: the mock is deployed by the test
///      contract, and the registry's CREATE address comes from DEPLOYER's nonce.
contract TransitionL1Test is RegistryTestBase {
    MockL1BlobVerifier internal mock;

    function _l1BlobVerifier() internal override returns (address) {
        mock = new MockL1BlobVerifier();
        return address(mock);
    }

    /// @dev 0x01 ++ sha256(commitment)[1:], the EIP-4844 versioned hash.
    function _vhs(Transition memory t) internal pure returns (bytes32[] memory vhs) {
        vhs = new bytes32[](t.commitments.length);
        for (uint256 i = 0; i < vhs.length; ++i) {
            vhs[i] = bytes32((uint256(sha256(t.commitments[i])) & type(uint256).max >> 8) | (uint256(1) << 248));
        }
    }

    function _blocks(uint256 n) internal pure returns (uint256[] memory b) {
        b = new uint256[](n);
        for (uint256 i = 0; i < n; ++i) {
            b[i] = 1000 + i;
        }
    }

    function _arm(Transition memory t) internal {
        mock.set(_vhs(t), _blocks(t.commitments.length), false);
    }

    function _submitL1(bytes31 pid, Transition memory t) internal {
        IL1BlobVerifier.L1BlobProof memory l1;
        registry.submitStateTransitionL1(pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs, l1);
    }

    function test_SubmitL1_OneBlob() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _arm(t);

        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessStateTransitioned(pid, address(this), t.rootBefore, t.rootAfter, 3, 0, 1);
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessStateTransitionedL1(pid, _blocks(1), _vhs(t));
        uint256 g = gasleft();
        _submitL1(pid, t);
        emit log_named_uint("submitStateTransitionL1 gas, 1 blob (mock verifier):", g - gasleft());

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.latestStateRoot, t.rootAfter);
        assertEq(p.votersCount, 3);
        assertEq(p.overwrittenVotesCount, 0);
        assertEq(p.batchNumber, 1);
        assertEq(p.lastVoteAt, vm.getBlockTimestamp());
    }

    function test_RevertWhen_Blob_WrongMode() public {
        bytes31 pid = _fixtureProcess();
        vm.expectRevert(IProcessRegistry.WrongDAMode.selector);
        _submit(pid, _transition(0));
    }

    function test_RevertWhen_L1_CountMismatch() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        bytes32[] memory v = new bytes32[](2);
        v[0] = _vhs(t)[0];
        v[1] = v[0];
        mock.set(v, _blocks(2), false);
        vm.expectRevert(IProcessRegistry.BlobCountMismatch.selector);
        _submitL1(pid, t);
    }

    function test_RevertWhen_L1_WrongHash() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        bytes32[] memory v = _vhs(t);
        v[0] = bytes32(uint256(v[0]) ^ 1);
        mock.set(v, _blocks(1), false);
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 0));
        _submitL1(pid, t);
    }

    function test_RevertWhen_L1_Digest() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _arm(t);
        t.ys[0] = bytes32(uint256(t.ys[0]) ^ 1);
        vm.expectRevert(IProcessRegistry.InvalidBlobsDigest.selector);
        _submitL1(pid, t);
    }

    function test_RevertWhen_L1_VerifierFails() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        mock.set(_vhs(t), _blocks(1), true);
        vm.expectRevert(IL1BlobVerifier.L1OriginMismatch.selector);
        _submitL1(pid, t);
    }

    function test_RevertWhen_L1_StaleRoot() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _arm(t);
        _submitL1(pid, t);
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        _submitL1(pid, t);
    }
}

/// @dev Blob-mode registry rejects the L1 entry point.
contract TransitionL1WrongModeTest is RegistryTestBase {
    function test_RevertWhen_L1_WrongMode() public {
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        IL1BlobVerifier.L1BlobProof memory l1;
        vm.expectRevert(IProcessRegistry.WrongDAMode.selector);
        registry.submitStateTransitionL1(pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs, l1);
    }
}
