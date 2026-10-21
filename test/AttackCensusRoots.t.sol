// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {AttackBase, AttackLib} from "./AttackBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {MockCensusValidator} from "./mocks/MockCensusValidator.sol";

/// @dev Census roots for every origin, with the Keccak roots of protocol v5 in mind: they
///      are full 256-bit values, most of them above the BN254 modulus, and nothing may
///      reduce, truncate or byte-swap them on the way from register 20 to the comparison.
contract AttackCensusRootsTest is AttackBase {
    uint256 internal constant P = BN254_P;

    /// @dev Edge roots: around the BN254 modulus, the top bit, the maximum, 1, and one
    ///      fuzzed value.
    function _edgeRoot(uint256 sel, uint256 raw) internal pure returns (bytes32) {
        uint256[9] memory edges =
            [uint256(1), P - 1, P, P + 1, 2 * P, uint256(1) << 255, type(uint256).max, type(uint256).max - P, raw];
        uint256 r = edges[sel % edges.length];
        return bytes32(r == 0 ? 1 : r);
    }

    function _cspRoot(uint256 sel, uint160 raw) internal pure returns (bytes32) {
        uint160[4] memory edges = [uint160(1), type(uint160).max, uint160(1) << 159, raw];
        uint160 r = edges[sel % edges.length];
        return bytes32(uint256(r == 0 ? 1 : r));
    }

    function _censusOf(DAVINCITypes.CensusOrigin origin, bytes32 root)
        internal
        pure
        returns (DAVINCITypes.Census memory c)
    {
        c.censusOrigin = origin;
        c.censusRoot = root;
        c.censusURI = "https://example.com/census";
    }

    function _create(DAVINCITypes.Census memory c) internal returns (bytes31) {
        return _newProcess(vm.getBlockTimestamp(), DURATION, MAX_VOTERS, _ballotMode(), c);
    }

    /// @dev The process's next batch, proven against census root r (BE integer).
    function _batch(bytes31 pid, bytes32 r, uint256 salt) internal view returns (AttackLib.Synth memory s) {
        s = _next(pid, 2, 0, salt);
        AttackLib.setReg32(s.pv, 20, AttackLib.rev(r));
    }

    function _origin(uint256 sel) internal pure returns (DAVINCITypes.CensusOrigin) {
        uint256 k = sel % 3;
        if (k == 0) return DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1;
        if (k == 1) return DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1;
        return DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1;
    }

    // --- origins 1, 2 and 4: the stored root, exactly ---------------------------------

    /// Origins 1, 2 and 4 settle a batch only on the stored root, bit for bit: never on the
    /// root reduced or lifted by the BN254 modulus, its byte reverse, or any other value.
    /// forge-config: default.fuzz.runs = 2000
    function testFuzz_FixedOriginsSettleOnlyTheStoredRoot(
        uint256 originSel,
        uint256 rootSel,
        uint256 raw,
        uint256 attackSel,
        bytes32 other
    ) public {
        _mockOpenings();
        DAVINCITypes.CensusOrigin origin = _origin(originSel);
        bool csp = origin == DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1;
        bytes32 stored = csp ? _cspRoot(rootSel, uint160(raw)) : _edgeRoot(rootSel, raw);
        bytes31 pid = _create(_censusOf(origin, stored));
        assertEq(registry.getProcess(pid).census.censusRoot, stored, "stored root changed at creation");

        uint256 s = uint256(stored);
        bytes32[6] memory attacks = [
            bytes32(s % P),
            bytes32(s >= P ? s - P : s + P),
            _rev(stored),
            bytes32(s ^ 1),
            bytes32(s ^ (uint256(1) << 255)),
            other
        ];
        bytes32 r = attacks[attackSel % attacks.length];
        if (r != stored) {
            AttackLib.Synth memory bad = _batch(pid, r, 1);
            vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
            _submitSynth(pid, bad);
        }
        _submitSynth(pid, _batch(pid, stored, 2));
        assertEq(registry.getProcess(pid).batchNumber, 1);
    }

    /// A CSP root is a uint160 address: 2^160 - 1 is accepted, anything with a bit at or
    /// above 160 is refused at creation, and so is 0.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_CspRootBounds(uint256 root) public {
        DAVINCITypes.Census memory c = _censusOf(DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1, bytes32(root));
        bool ok = root != 0 && root >> 160 == 0;
        if (!ok) vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _create(c);
    }

    function test_CspRootEdges() public {
        uint256[5] memory refused = [uint256(0), uint256(1) << 160, (uint256(1) << 160) + 1, P, type(uint256).max];
        for (uint256 i = 0; i < refused.length; ++i) {
            DAVINCITypes.Census memory c =
                _censusOf(DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1, bytes32(refused[i]));
            vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
            _create(c);
        }
        _create(_censusOf(DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1, bytes32(uint256(type(uint160).max))));
    }

    /// Origin 2: a sequence of updates to roots above the modulus. After each update only the
    /// newest root settles, and an earlier root (even one that settled before) is refused.
    /// forge-config: default.fuzz.runs = 500
    function testFuzz_Origin2UpdatesAboveTheModulus(uint256[6] memory raws) public {
        _mockOpenings();
        bytes32 first = bytes32(P + 7);
        bytes31 pid = _create(_censusOf(DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1, first));
        bytes32 prev = first;
        _submitSynth(pid, _batch(pid, first, 100));
        for (uint256 i = 0; i < raws.length; ++i) {
            bytes32 next = bytes32(P + bound(raws[i], 1, type(uint256).max - P));
            DAVINCITypes.Census memory c = _censusOf(DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1, next);
            vm.prank(ORGANIZER);
            registry.setProcessCensus(pid, c);
            assertEq(registry.getProcess(pid).census.censusRoot, next);
            if (prev != next) {
                AttackLib.Synth memory stale = _batch(pid, prev, i);
                vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
                _submitSynth(pid, stale);
            }
            AttackLib.Synth memory reduced = _batch(pid, bytes32(uint256(next) % P), i);
            vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
            _submitSynth(pid, reduced);
            _submitSynth(pid, _batch(pid, next, 200 + i));
            prev = next;
        }
        assertEq(registry.getProcess(pid).batchNumber, 1 + raws.length);
    }

    // --- origin 3: the census contract's root history -----------------------------------

    /// Origin 3 against a census contract with OnchainCensus semantics: roots set before
    /// and after creation, in a random block schedule. A batch settles on root R exactly
    /// when the contract held R at some block in [creationBlock, block.number]; R mod p,
    /// R - p and never-held roots are refused. Each root is tried on a snapshot.
    /// forge-config: default.fuzz.runs = 300
    function testFuzz_Origin3RootWindow(uint256[10] memory raws, uint8[10] memory gaps, uint8 createAt) public {
        _mockOpenings();
        vm.roll(1000);
        MockCensusValidator census = new MockCensusValidator();
        createAt = uint8(bound(createAt, 0, 9));
        uint256[] memory roots = new uint256[](10);
        bytes31 pid;
        uint256 creation;
        for (uint256 i = 0; i < 10; ++i) {
            vm.roll(vm.getBlockNumber() + (gaps[i] % 3));
            // Mostly above the modulus, as Keccak roots are.
            roots[i] = raws[i] % 4 == 0 ? bound(raws[i], 1, P - 1) : P + bound(raws[i], 1, type(uint256).max - P);
            census.setRoot(roots[i]);
            if (i == createAt) {
                DAVINCITypes.Census memory c;
                c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1;
                c.contractAddress = address(census);
                c.censusURI = "onchain://census";
                pid = _create(c);
                creation = vm.getBlockNumber();
                assertEq(uint256(registry.getProcess(pid).census.censusRoot), roots[i], "stored informational root");
            }
        }
        vm.roll(vm.getBlockNumber() + 3);

        uint256 held;
        for (uint256 i = 0; i < 10; ++i) {
            // The answer the census gives is the spec input; the registry must follow it.
            uint256 rbn = census.getRootBlockNumber(roots[i]);
            bool want = rbn != 0 && rbn >= creation && rbn <= vm.getBlockNumber();
            if (want) ++held;
            uint256 snap = vm.snapshot();
            bool ok = _trySubmit(pid, _batch(pid, bytes32(roots[i]), i));
            assertEq(ok, want, string.concat("root ", vm.toString(i)));
            vm.revertTo(snap);
            _mockOpenings();

            uint256 alias_ = roots[i] >= P ? roots[i] - P : roots[i] + P;
            if (census.getRootBlockNumber(alias_) == 0) {
                assertFalse(_trySubmit(pid, _batch(pid, bytes32(alias_), 50 + i)), "aliased root settled");
            }
            if (roots[i] >= P && census.getRootBlockNumber(roots[i] % P) == 0) {
                assertFalse(_trySubmit(pid, _batch(pid, bytes32(roots[i] % P), 70 + i)), "reduced root settled");
            }
        }
        assertFalse(_trySubmit(pid, _batch(pid, bytes32(0), 99)), "root 0 settled");
        // The root current at creation and every later one are usable.
        assertGe(held, 10 - createAt);
    }

    /// The census moves in the same block, right before the settlement: the batch proven
    /// on the root that was just replaced still lands (its last valid block is this one),
    /// and so does the next batch on the new root.
    function test_Origin3RootReplacedInTheSettlementBlock() public {
        _mockOpenings();
        vm.roll(500);
        MockCensusValidator census = new MockCensusValidator();
        uint256 r1 = P + 1;
        uint256 r2 = type(uint256).max;
        census.setRoot(r1);
        DAVINCITypes.Census memory c;
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1;
        c.contractAddress = address(census);
        c.censusURI = "onchain://census";
        bytes31 pid = _create(c);
        vm.roll(510);
        census.setRoot(r2);
        _submitSynth(pid, _batch(pid, bytes32(r1), 1));
        _submitSynth(pid, _batch(pid, bytes32(r2), 2));
        vm.roll(600);
        _submitSynth(pid, _batch(pid, bytes32(r1), 3));
        assertEq(registry.getProcess(pid).batchNumber, 3);
    }
}
