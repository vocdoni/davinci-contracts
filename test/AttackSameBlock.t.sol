// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {AttackBase} from "./AttackBase.sol";
import {LifecycleHandler} from "./AttackLifecycle.t.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";

/// @dev Same-block races at the lifecycle boundaries. The lifecycle model of
///      AttackLifecycle.t.sol is driven to one boundary of one process (its start, its end, its
///      grace end, a second either side) and then runs a burst of organizer, stranger,
///      sequencer and results actions with no time passing between them: a settlement and a
///      status change in one block, a late settlement extending the window right before a
///      cancel or a results call, the results call racing the cancel at the grace end.
///      Every outcome must match the model.
contract AttackSameBlockTest is AttackBase {
    LifecycleHandler internal handler;
    bytes31[] internal ps;

    function setUp() public override {
        super.setUp();
        ps.push(_process(vm.getBlockTimestamp(), 2 hours, 200));
        DAVINCITypes.Census memory dyn = _census();
        dyn.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1;
        ps.push(_newProcess(vm.getBlockTimestamp() + 10 minutes, 1 hours, 100, _ballotMode(), dyn));
        handler = new LifecycleHandler(registry, ps);
    }

    /// @dev The boundary timestamp `kind` of process pid, from the model.
    function _boundary(bytes31 pid, uint256 kind) internal view returns (uint256 t) {
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        uint256 end = p.startTime + p.duration;
        uint256 graceEnd = handler.graceEndOf(pid);
        uint256[6] memory ts = [p.startTime, end - 1, end, graceEnd - 1, graceEnd, graceEnd + 1];
        t = ts[kind % ts.length];
    }

    function _act(uint8 op, uint256 a, uint256 b, uint64[16] memory values) internal {
        op = op % 8;
        uint256 pi = a % 2;
        bool stranger = b % 9 == 0;
        if (op == 0) handler.submit(pi, uint64(1 + b % 5), uint64(b % 2), b % 7 == 0 ? b % 15 : 0, b);
        else if (op == 1) handler.setStatus(pi, uint8(b % 5), stranger);
        else if (op == 2) handler.results(pi, b % 11 == 0, values);
        else if (op == 3) handler.setDuration(pi, b, stranger);
        else if (op == 4) handler.setGrace(pi, uint32(b), stranger);
        else if (op == 5) handler.setMaxVoters(pi, b, stranger);
        else if (op == 6) handler.setCensus(pi, bytes32(b), false, stranger);
        else handler.submit(pi, uint64(1 + b % 3), 0, 0, b ^ 0x5eed);
    }

    /// Before the burst, batches land at fuzzed times so lastVoteAt (and the grace end) vary;
    /// then a burst at one boundary, a step of one second, and a second burst.
    /// forge-config: default.fuzz.runs = 1500
    function testFuzz_SameBlockBurstsAtBoundaries(
        uint256 who,
        uint256 kind,
        uint256[3] memory pre,
        uint8[10] memory ops,
        uint256[10] memory args,
        uint64[16] memory values
    ) public {
        _mockOpenings();
        for (uint256 i = 0; i < pre.length; ++i) {
            handler.warp(pre[i]);
            handler.submit(pre[i] % 2, 2, 0, 0, pre[i]);
        }
        bytes31 pid = ps[who % ps.length];
        uint256 t = _boundary(pid, kind);
        if (t > vm.getBlockTimestamp()) vm.warp(t);
        handler.warp(0); // next block, same second; lets the model see a closed window
        for (uint256 i = 0; i < 5; ++i) {
            _act(ops[i], args[i], uint256(keccak256(abi.encode(args[i], i))), values);
        }
        handler.warp(1);
        for (uint256 i = 5; i < 10; ++i) {
            _act(ops[i], args[i], uint256(keccak256(abi.encode(args[i], i))), values);
        }
        assertEq(handler.violations(), 0, handler.firstViolation());
    }

    /// The results race at the grace end, spelled out: a cancel one second before it wins
    /// and voids the results; at the grace end the cancel is refused and the results land,
    /// in either order within the block.
    /// forge-config: default.fuzz.runs = 500
    function testFuzz_ResultsVersusCancelAtTheGraceEnd(bool resultsFirst, bool lateBatch, uint64[16] memory values)
        public
    {
        _mockOpenings();
        bytes31 pid = ps[0];
        handler.submit(0, 3, 0, 0, 1);
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        vm.warp(p.startTime + p.duration + 10);
        if (lateBatch) handler.submit(0, 1, 0, 0, 2); // moves the grace end
        uint256 graceEnd = handler.graceEndOf(pid);
        assertEq(registry.getProcessGraceEnd(pid), graceEnd);

        uint256 snap = vm.snapshot();
        vm.warp(graceEnd - 1);
        handler.results(0, false, values); // GraceOpen
        handler.setStatus(0, 2, false); // cancel lands
        handler.results(0, false, values); // refused: canceled
        assertEq(uint8(registry.getProcess(pid).status), 2);
        assertEq(handler.violations(), 0, handler.firstViolation());
        vm.revertTo(snap);

        vm.warp(graceEnd);
        if (resultsFirst) {
            handler.results(0, false, values);
            handler.setStatus(0, 2, false);
        } else {
            handler.setStatus(0, 2, false);
            handler.results(0, false, values);
        }
        handler.submit(0, 1, 0, 0, 3); // too late
        assertEq(uint8(registry.getProcess(pid).status), 4);
        assertEq(handler.violations(), 0, handler.firstViolation());
    }
}
