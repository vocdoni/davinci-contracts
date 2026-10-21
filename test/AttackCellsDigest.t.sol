// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {IL1BlobVerifier} from "../src/interfaces/IL1BlobVerifier.sol";
import {MockL1BlobVerifier} from "./mocks/MockL1BlobVerifier.sol";

/// @dev The cell digest D (registers 46..53) seeds every blob's evaluation point. With the
///      fixture's real KZG openings and the accept-all PLONK, any change to a byte of D the
///      registry reads moves every point and fails the first opening; the bytes it does not
///      read (the high halves, and registers 37-41, 43-45, 54-63) are left to the SNARK.
contract AttackCellsDigestTest is RegistryTestBase {
    using stdJson for string;

    uint256 internal constant REG_D = 46;

    /// @dev Transition i of the fixture process, after settling the ones before it.
    function _ready(uint256 i) internal returns (bytes31 pid, Transition memory t) {
        pid = _fixtureProcess();
        for (uint256 j = 0; j < i; ++j) {
            _submit(pid, _transition(j));
        }
        t = _transition(i);
    }

    function _try(bytes31 pid, Transition memory t) internal returns (bool ok, bytes memory ret) {
        vm.blobhashes(t.versionedHashes);
        (ok, ret) = address(registry)
            .call(
                abi.encodeCall(
                    registry.submitStateTransition,
                    (pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs)
                )
            );
    }

    /// Every bit of every byte of D the registry reads is bound by the openings, for the
    /// one-blob and the two-blob transition.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_EveryReadByteOfDIsBound(bool second, uint256 k, uint256 b, uint8 x) public {
        (bytes31 pid, Transition memory t) = _ready(second ? 1 : 0);
        k = bound(k, REG_D, REG_D + 7);
        b = bound(b, 0, 3);
        x = uint8(bound(x, 1, 255));
        t.publicValues[8 * k + b] = bytes1(uint8(t.publicValues[8 * k + b]) ^ x);
        (bool ok, bytes memory ret) = _try(pid, t);
        assertFalse(ok, "tampered D settled");
        assertEq(ret, abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 0));
        assertEq(registry.getProcess(pid).batchNumber, second ? 1 : 0);
    }

    /// D cannot be borrowed from another batch: transition 1's D on transition 0 (and back)
    /// fails, though both are genuine guest outputs.
    function test_DOfAnotherBatchRefused() public {
        bytes32 d0 = fixture.readBytes32(".transitions[0].cells_digest");
        bytes32 d1 = fixture.readBytes32(".transitions[1].cells_digest");
        assertTrue(d0 != d1);

        (bytes31 pid, Transition memory t) = _ready(0);
        _setReg32(t.publicValues, REG_D, d1);
        (bool ok,) = _try(pid, t);
        assertFalse(ok, "transition 0 with transition 1's D");

        _submit(pid, _transition(0));
        t = _transition(1);
        _setReg32(t.publicValues, REG_D, d0);
        (ok,) = _try(pid, t);
        assertFalse(ok, "transition 1 with transition 0's D");
    }

    /// The high halves of D's words are not read (the guest writes u32 limbs); only the
    /// SNARK binds them. With the accept-all PLONK the batch settles unchanged.
    /// forge-config: default.fuzz.runs = 500
    function testFuzz_DHighHalvesLeftToTheSnark(uint256 k, uint256 b, uint8 x) public {
        (bytes31 pid, Transition memory t) = _ready(0);
        k = bound(k, REG_D, REG_D + 7);
        b = bound(b, 4, 7);
        x = uint8(bound(x, 1, 255));
        t.publicValues[8 * k + b] = bytes1(uint8(t.publicValues[8 * k + b]) ^ x);
        (bool ok,) = _try(pid, t);
        assertTrue(ok, "a high half of D is read");
        assertEq(registry.getProcess(pid).latestStateRoot, t.rootAfter);
    }

    /// Registers 37-41, 43-45 and 54-63 are diagnostics or zero padding; the registry does
    /// not read them, so a batch settles whatever they hold once the SNARK accepts it.
    /// forge-config: default.fuzz.runs = 500
    function testFuzz_UnreadRegistersLeftToTheSnark(uint256 k, uint256 b, uint8 x) public {
        (bytes31 pid, Transition memory t) = _ready(0);
        k = bound(k, 37, 63);
        vm.assume(k != 42 && (k < REG_D || k > REG_D + 7));
        b = bound(b, 0, 7);
        x = uint8(bound(x, 1, 255));
        t.publicValues[8 * k + b] = bytes1(uint8(t.publicValues[8 * k + b]) ^ x);
        (bool ok,) = _try(pid, t);
        assertTrue(ok, "an unread register is read");
    }
}

/// @dev The same binding in L1-blob mode, where the blob hashes come from the L1 verifier.
contract AttackCellsDigestL1Test is RegistryTestBase {
    MockL1BlobVerifier internal mock;

    function _l1BlobVerifier() internal override returns (address) {
        mock = new MockL1BlobVerifier();
        return address(mock);
    }

    function _arm(Transition memory t) internal {
        bytes32[] memory vhs = new bytes32[](t.commitments.length);
        uint256[] memory blocks = new uint256[](vhs.length);
        for (uint256 i = 0; i < vhs.length; ++i) {
            vhs[i] = bytes32((uint256(sha256(t.commitments[i])) & type(uint256).max >> 8) | (uint256(1) << 248));
            blocks[i] = 1000 + i;
        }
        mock.set(vhs, blocks, false);
    }

    function _tryL1(bytes31 pid, Transition memory t) internal returns (bool ok, bytes memory ret) {
        IL1BlobVerifier.L1BlobProof memory l1;
        (ok, ret) = address(registry)
            .call(
                abi.encodeCall(
                    registry.submitStateTransitionL1,
                    (pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs, l1)
                )
            );
    }

    /// forge-config: default.fuzz.runs = 500
    function testFuzz_L1EveryReadByteOfDIsBound(bool second, uint256 k, uint256 b, uint8 x) public {
        bytes31 pid = _fixtureProcess();
        if (second) {
            Transition memory t0 = _transition(0);
            _arm(t0);
            (bool ok0,) = _tryL1(pid, t0);
            assertTrue(ok0, "honest first batch");
        }
        Transition memory t = _transition(second ? 1 : 0);
        _arm(t);
        k = bound(k, 46, 53);
        b = bound(b, 0, 3);
        x = uint8(bound(x, 1, 255));
        t.publicValues[8 * k + b] = bytes1(uint8(t.publicValues[8 * k + b]) ^ x);
        (bool ok, bytes memory ret) = _tryL1(pid, t);
        assertFalse(ok, "tampered D settled in L1 mode");
        assertEq(ret, abi.encodeWithSelector(IProcessRegistry.InvalidBlobOpening.selector, 0));
    }
}
