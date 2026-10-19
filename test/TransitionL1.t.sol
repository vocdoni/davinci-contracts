// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {IL1BlobVerifier} from "../src/interfaces/IL1BlobVerifier.sol";
import {L1BlobVerifier} from "../src/L1BlobVerifier.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {MockL1BlobVerifier} from "./mocks/MockL1BlobVerifier.sol";
import {MockL1Block} from "./mocks/MockL1Block.sol";

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

    function test_RevertWhen_L1_CellsDigest() public {
        // L1 mode derives the points from the same proven D.
        bytes31 pid = _fixtureProcess();
        Transition memory t = _transition(0);
        _arm(t);
        _setReg32(t.publicValues, 46, bytes32(0));
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

/// @dev L1-blob mode with the real L1BlobVerifier: the fixture transitions' blobs in synthetic
///      Ethereum blob transactions (test/vectors/l1blob, written by `go run .`). The second
///      transition's two blobs went out in two transactions, as a node posts a transition
///      above one transaction's blob limit; the proof lists them in blob order.
contract TransitionL1TwoTxTest is RegistryTestBase {
    using stdJson for string;

    MockL1Block internal l1;

    function _l1BlobVerifier() internal override returns (address) {
        l1 = new MockL1Block();
        return address(new L1BlobVerifier(address(l1)));
    }

    /// @dev The L1 proof of fixture `name`, with OP's origin set at its top header.
    function _proof(string memory name)
        internal
        returns (IL1BlobVerifier.L1BlobProof memory p, uint256[] memory blocks)
    {
        string memory j = vm.readFile(string.concat("test/vectors/l1blob/", name, ".json"));
        p.headers = j.readBytesArray(".headers");
        uint256 nt = abi.decode(vm.parseJson(j, ".txs"), (bytes[])).length;
        p.txs = new IL1BlobVerifier.L1TxProof[](nt);
        for (uint256 i = 0; i < nt; ++i) {
            string memory b = string.concat(".txs[", vm.toString(i), "]");
            p.txs[i].headerIndex = j.readUint(string.concat(b, ".header_index"));
            p.txs[i].key = j.readBytes(string.concat(b, ".key"));
            p.txs[i].proof = j.readBytesArray(string.concat(b, ".proof"));
        }
        blocks = j.readUintArray(".expected_tx_blocks");
        l1.set(uint64(j.readUintArray(".header_numbers")[0]), j.readBytes32Array(".header_hashes")[0]);
    }

    function _submitL1(bytes31 pid, Transition memory t, IL1BlobVerifier.L1BlobProof memory p) internal {
        registry.submitStateTransitionL1(pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs, p);
    }

    /// @dev The first transition, from its one-transaction proof.
    function _first() internal returns (bytes31 pid) {
        pid = _fixtureProcess();
        (IL1BlobVerifier.L1BlobProof memory p,) = _proof("one_tx");
        _submitL1(pid, _transition(0), p);
    }

    function _settledTwoTx(bytes31 pid, string memory name, uint256 b0, uint256 b1) internal {
        Transition memory t = _transition(1);
        (IL1BlobVerifier.L1BlobProof memory p, uint256[] memory blocks) = _proof(name);
        assertEq(t.commitments.length, 2);
        assertEq(blocks[0], b0);
        assertEq(blocks[1], b1);
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessStateTransitionedL1(pid, blocks, t.versionedHashes);
        _submitL1(pid, t, p);
        DAVINCITypes.Process memory proc = registry.getProcess(pid);
        assertEq(proc.latestStateRoot, t.rootAfter);
        assertEq(proc.batchNumber, 2);
        assertEq(proc.votersCount, 1000);
    }

    function test_SubmitL1_TwoTransactionsInChunkOrder() public {
        _settledTwoTx(_first(), "two_tx", 2000, 2002);
    }

    function test_SubmitL1_RepostedFirstChunkInTheNewerBlock() public {
        _settledTwoTx(_first(), "two_tx_reposted", 2002, 2000);
    }

    /// The same two transactions listed newest blob first: the openings follow the proof order,
    /// so commitment 0 meets the second blob's hash.
    function test_RevertWhen_L1_ChunksOutOfOrder() public {
        bytes31 pid = _first();
        Transition memory t = _transition(1);
        (IL1BlobVerifier.L1BlobProof memory p,) = _proof("two_tx");
        (p.txs[0], p.txs[1]) = (p.txs[1], p.txs[0]);
        vm.expectRevert(abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 0));
        _submitL1(pid, t, p);
    }
}
