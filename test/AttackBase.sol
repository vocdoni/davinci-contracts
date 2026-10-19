// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {RegistryTestBase} from "./RegistryTestBase.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";

/// @dev Builders for synthetic batch publics and blob arrays.
library AttackLib {
    uint256 internal constant BLS_MODULUS =
        52435875175126190479447740508185965837690552500527637822603658699938581184513;

    struct Synth {
        bytes pv;
        bytes[] commitments;
        bytes32[] ys;
        bytes[] kzgProofs;
        bytes32[] vhs;
    }

    /// @dev 0x01 ++ sha256(commitment)[1:].
    function vh(bytes memory commitment) internal pure returns (bytes32) {
        return bytes32((uint256(sha256(commitment)) & (type(uint256).max >> 8)) | (uint256(1) << 248));
    }

    /// @dev Register k as an 8-byte LE word.
    function setWord(bytes memory pv, uint256 k, uint64 v) internal pure {
        for (uint256 i = 0; i < 8; i++) {
            pv[8 * k + i] = bytes1(uint8(v >> (8 * i)));
        }
    }

    /// @dev Registers k..k+7 so that PublicsLib.reg32(k) == v.
    function setReg32(bytes memory pv, uint256 k, bytes32 v) internal pure {
        for (uint256 j = 0; j < 8; j++) {
            for (uint256 t = 0; t < 4; t++) {
                pv[8 * (k + j) + t] = v[4 * j + t];
            }
        }
    }

    function rev(bytes32 x) internal pure returns (bytes32 r) {
        for (uint256 i = 0; i < 32; ++i) {
            r |= bytes32(uint256(uint8(x[i])) << (8 * i));
        }
    }

    /// @dev A well-formed batch publics block and blob arrays. censusBE is the census root
    ///      as the registry stores it (big-endian integer).
    function synth(
        bytes32 rootBefore,
        bytes32 rootAfter,
        uint64 voters,
        uint64 overwrites,
        bytes32 censusBE,
        uint64 occupied,
        uint256 nBlobs,
        uint256 salt
    ) internal pure returns (Synth memory s) {
        s.commitments = new bytes[](nBlobs);
        s.ys = new bytes32[](nBlobs);
        s.kzgProofs = new bytes[](nBlobs);
        s.vhs = new bytes32[](nBlobs);
        bytes memory pairs;
        for (uint256 i = 0; i < nBlobs; ++i) {
            bytes32 a = keccak256(abi.encode(salt, i, "c"));
            s.commitments[i] = bytes.concat(a, bytes16(keccak256(abi.encode(a))));
            s.ys[i] = bytes32(uint256(keccak256(abi.encode(salt, i, "y"))) % BLS_MODULUS);
            s.kzgProofs[i] = bytes.concat(keccak256(abi.encode(salt, i, "p")), bytes16(0));
            s.vhs[i] = vh(s.commitments[i]);
            pairs = bytes.concat(pairs, s.commitments[i], s.ys[i]);
        }
        s.pv = new bytes(512);
        setWord(s.pv, 0, 1);
        setReg32(s.pv, 2, rootBefore);
        setReg32(s.pv, 10, rootAfter);
        setWord(s.pv, 18, voters);
        setWord(s.pv, 19, overwrites);
        setReg32(s.pv, 20, rev(censusBE));
        setReg32(s.pv, 28, sha256(pairs));
        setWord(s.pv, 36, uint64(nBlobs));
        setWord(s.pv, 42, occupied);
    }

    /// @dev Results-guest publics over root: ok, the root, 16 results as (lo, hi).
    function resultsPublics(bytes32 root, uint64[16] memory values) internal pure returns (bytes memory pv) {
        pv = new bytes(512);
        setWord(pv, 0, 1);
        setReg32(pv, 2, root);
        for (uint256 i = 0; i < 16; i++) {
            setWord(pv, 10 + 2 * i, uint64(uint32(values[i])));
            setWord(pv, 11 + 2 * i, values[i] >> 32);
        }
    }
}

/// @dev Synthetic transitions for the attack suites. The PLONK verifier is the accept-all
///      mock and the point-evaluation precompile is mocked to accept every opening, so any
///      register layout can be settled; what is left is exactly the registry's own checks.
///      Suites that need real openings use the fixture instead.
abstract contract AttackBase is RegistryTestBase {
    uint256 internal constant BLS_MODULUS = AttackLib.BLS_MODULUS;

    /// @dev Every call to 0x0A answers like a passing point evaluation.
    function _mockOpenings() internal {
        vm.mockCall(address(0x0A), bytes(""), abi.encode(uint256(4096), BLS_MODULUS));
    }

    function _synth(
        bytes32 rootBefore,
        bytes32 rootAfter,
        uint64 voters,
        uint64 overwrites,
        bytes32 censusBE,
        uint64 occupied,
        uint256 nBlobs,
        uint256 salt
    ) internal pure returns (AttackLib.Synth memory) {
        return AttackLib.synth(rootBefore, rootAfter, voters, overwrites, censusBE, occupied, nBlobs, salt);
    }

    /// @dev An honest next transition of pid: its root, census and occupancy.
    function _next(bytes31 pid, uint64 voters, uint64 overwrites, uint256 salt)
        internal
        view
        returns (AttackLib.Synth memory)
    {
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        return AttackLib.synth(
            p.latestStateRoot,
            keccak256(abi.encode(pid, p.batchNumber, salt)),
            voters,
            overwrites,
            p.census.censusRoot,
            uint64(p.votersCount),
            1 + (salt % 3),
            salt
        );
    }

    function _submitSynth(bytes31 pid, AttackLib.Synth memory s) internal {
        vm.blobhashes(s.vhs);
        registry.submitStateTransition(pid, s.pv, hex"", s.commitments, s.ys, s.kzgProofs);
    }

    /// @dev Low-level submit: whether it succeeded, without bubbling the revert.
    function _trySubmit(bytes31 pid, AttackLib.Synth memory s) internal returns (bool ok) {
        vm.blobhashes(s.vhs);
        (ok,) = address(registry)
            .call(abi.encodeCall(registry.submitStateTransition, (pid, s.pv, hex"", s.commitments, s.ys, s.kzgProofs)));
    }

    function _rev(bytes32 x) internal pure returns (bytes32) {
        return AttackLib.rev(x);
    }

    /// @dev A static-census process of ORGANIZER with the fixture's ballot mode and key.
    function _process(uint256 startTime, uint256 duration, uint256 maxVoters) internal returns (bytes31) {
        return _newProcess(startTime, duration, maxVoters, _ballotMode(), _census());
    }

    function _status(bytes31 pid) internal view returns (DAVINCITypes.ProcessStatus) {
        return registry.getProcess(pid).status;
    }

    function _setStatus(bytes31 pid, DAVINCITypes.ProcessStatus s) internal {
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, s);
    }
}
