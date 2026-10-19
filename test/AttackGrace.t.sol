// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {AttackBase, AttackLib} from "./AttackBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {MockZiskVerifier} from "./mocks/MockZiskVerifier.sol";

/// @dev Grace window arithmetic under attack: exact boundaries, extension rounds up to the
///      cap, a window that must never reopen, end-time changes with batches pending, and
///      values near 2^256.
contract AttackGraceTest is AttackBase {
    uint256 internal salt;

    function setUp() public override {
        super.setUp();
        _mockOpenings();
    }

    function _land(bytes31 pid) internal returns (bool) {
        return _trySubmit(pid, _next(pid, 1, 0, ++salt));
    }

    /// The window is half-open: a batch at graceEnd - 1 lands, one at graceEnd does not.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_ExactBoundary(uint32 grace, uint256 lastOff, bool vote) public {
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        grace = uint32(bound(grace, GRACE_FLOOR, GRACE_CEIL));
        vm.prank(ORGANIZER);
        registry.setProcessGrace(pid, grace);
        uint256 end = registry.getProcessEndTime(pid);
        if (vote) {
            // A landing inside the window moves it.
            vm.warp(end + bound(lastOff, 0, grace - 1));
            assertTrue(_land(pid));
        }
        uint256 ge = registry.getProcessGraceEnd(pid);
        uint256 last = registry.getProcess(pid).lastVoteAt;
        uint256 idle = (last > end ? last : end) + grace;
        assertEq(ge, idle < end + GRACE_MAX_TOTAL ? idle : end + GRACE_MAX_TOTAL);

        uint256 snap = vm.snapshot();
        vm.warp(ge - 1);
        assertTrue(_land(pid), "batch at graceEnd - 1 refused");
        vm.revertTo(snap);
        vm.warp(ge);
        assertFalse(_land(pid), "batch at graceEnd settled");
    }

    /// Batches keep the window open while they land less than grace apart, never past
    /// end + graceMaxTotal; once a gap reaches the window it is closed for good.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_ExtensionRounds(uint256[24] memory gaps) public {
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        uint256 end = registry.getProcessEndTime(pid);
        vm.warp(end);
        bool closed;
        for (uint256 i = 0; i < gaps.length; ++i) {
            uint256 ge = registry.getProcessGraceEnd(pid);
            assertLe(ge, end + GRACE_MAX_TOTAL, "window past the cap");
            vm.warp(vm.getBlockTimestamp() + bound(gaps[i], 0, GRACE + 20));
            bool want = !closed && vm.getBlockTimestamp() < ge;
            bool ok = _land(pid);
            assertEq(ok, want, "landing vs window");
            if (!ok && vm.getBlockTimestamp() >= ge) closed = true;
            if (ok) assertEq(registry.getProcess(pid).lastVoteAt, vm.getBlockTimestamp());
        }
        // Whatever happened, at the cap nothing lands.
        vm.warp(end + GRACE_MAX_TOTAL);
        assertFalse(_land(pid));
    }

    /// A batch every grace - 1 seconds hits the cap exactly: the last landing is at
    /// end + graceMaxTotal - 1 at most.
    function test_CapReachedByBackToBackBatches() public {
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        uint256 end = registry.getProcessEndTime(pid);
        vm.warp(end);
        uint256 landed;
        while (_land(pid)) {
            landed = vm.getBlockTimestamp();
            vm.warp(vm.getBlockTimestamp() + GRACE - 1);
        }
        assertLt(landed, end + GRACE_MAX_TOTAL);
        assertGe(landed + GRACE - 1, end + GRACE_MAX_TOTAL);
        assertEq(registry.getProcessGraceEnd(pid), end + GRACE_MAX_TOTAL);
    }

    /// After the window closes no organizer action reopens it and no batch lands, at any
    /// later time.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_NeverReopens(uint256 later, uint256 d, uint32 g, bool landFirst) public {
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        uint256 end = registry.getProcessEndTime(pid);
        if (landFirst) {
            vm.warp(end + 10);
            assertTrue(_land(pid));
        }
        uint256 ge = registry.getProcessGraceEnd(pid);
        vm.warp(ge + bound(later, 0, 365 days));

        vm.startPrank(ORGANIZER);
        (bool ok,) = address(registry).call(abi.encodeCall(registry.setProcessDuration, (pid, bound(d, 1, 1000 days))));
        assertFalse(ok, "duration moved after the window");
        (ok,) = address(registry)
            .call(abi.encodeCall(registry.setProcessGrace, (pid, uint32(bound(g, GRACE_FLOOR, GRACE_CEIL)))));
        assertFalse(ok, "grace moved after the window");
        (ok,) = address(registry).call(abi.encodeCall(registry.setProcessMaxVoters, (pid, MAX_VOTERS + 1)));
        assertFalse(ok, "maxVoters moved after the window");
        (ok,) =
            address(registry).call(abi.encodeCall(registry.setProcessStatus, (pid, DAVINCITypes.ProcessStatus.PAUSED)));
        assertFalse(ok, "paused after the window");
        (ok,) =
            address(registry).call(abi.encodeCall(registry.setProcessStatus, (pid, DAVINCITypes.ProcessStatus.ENDED)));
        assertTrue(ok, "ENDED past the end is a no-op move");
        vm.stopPrank();

        assertEq(registry.getProcessGraceEnd(pid), ge, "window moved");
        assertEq(registry.getProcessEndTime(pid), end, "end moved");
        assertFalse(_land(pid), "batch after the window");
    }

    /// Shortening with exactly noticeMin while a batch is pending: it lands until the new
    /// window closes, and the old end no longer counts.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_ShortenWhilePending(uint256 callAt, uint256 landAt) public {
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        uint256 start = vm.getBlockTimestamp();
        uint256 oldEnd = start + DURATION;
        uint256 tc = start + bound(callAt, 0, DURATION - NOTICE_MIN - 1);
        vm.warp(tc);
        uint256 newEnd = tc + NOTICE_MIN;
        vm.prank(ORGANIZER);
        registry.setProcessDuration(pid, newEnd - start);
        // one second less notice is refused
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        registry.setProcessDuration(pid, newEnd - start - 1);

        uint256 tl = tc + bound(landAt, 0, oldEnd + GRACE - tc);
        vm.warp(tl);
        assertEq(_land(pid), tl < newEnd + GRACE, "pending batch vs shortened window");
    }

    /// Repeated shortening cannot walk the end below now + noticeMin, however it is split.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_RepeatedShortening(uint256[6] memory steps) public {
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        uint256 start = vm.getBlockTimestamp();
        for (uint256 i = 0; i < steps.length; ++i) {
            vm.warp(vm.getBlockTimestamp() + bound(steps[i], 0, 30));
            uint256 end = registry.getProcessEndTime(pid);
            if (end <= vm.getBlockTimestamp()) break;
            uint256 target = vm.getBlockTimestamp() + bound(steps[i] >> 8, 0, 2 * NOTICE_MIN);
            if (target <= start || target == end) continue;
            vm.prank(ORGANIZER);
            (bool ok,) = address(registry).call(abi.encodeCall(registry.setProcessDuration, (pid, target - start)));
            bool want =
                target > vm.getBlockTimestamp() && (target > end || target >= vm.getBlockTimestamp() + NOTICE_MIN);
            assertEq(ok, want, "shortening rule");
            assertGe(registry.getProcessEndTime(pid), vm.getBlockTimestamp() + (ok && target < end ? NOTICE_MIN : 1));
        }
    }

    /// Ends near 2^256: the window saturates instead of overflowing, and nothing below the
    /// saturation point overflows either.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_NearOverflow(uint256 dur, bool landing) public {
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        uint256 start = vm.getBlockTimestamp();
        if (landing) assertTrue(_land(pid));
        dur = bound(dur, type(uint256).max - start - 2 * GRACE_MAX_TOTAL, type(uint256).max - start);
        vm.prank(ORGANIZER);
        registry.setProcessDuration(pid, dur);
        uint256 end = start + dur;
        uint256 ge = registry.getProcessGraceEnd(pid);
        if (end > type(uint256).max - GRACE_MAX_TOTAL) assertEq(ge, type(uint256).max);
        else assertEq(ge, end + GRACE);
        // Still open, and still settles.
        assertTrue(_land(pid));
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessResults(pid, new bytes(512), hex"");
    }

    /// The constructor accepts exactly 0 < floor <= default <= ceil <= maxTotal, notice > 0.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_ConstructorBounds(uint32 dflt, uint32 floor_, uint32 ceil_, uint32 maxTotal, uint32 notice)
        public
    {
        bool want = floor_ != 0 && floor_ <= dflt && dflt <= ceil_ && ceil_ <= maxTotal && notice != 0;
        address v = address(new MockZiskVerifier());
        if (!want) vm.expectRevert(IProcessRegistry.InvalidGrace.selector);
        ProcessRegistry r = new ProcessRegistry(
            CHAIN_ID,
            v,
            BATCH_VK,
            RESULTS_VK,
            ROOT_C,
            keccak256("vk"),
            address(0),
            address(0),
            dflt,
            floor_,
            ceil_,
            maxTotal,
            notice,
            address(0)
        );
        if (want) assertEq(r.graceMaxTotal(), maxTotal);
    }

    // BUG: newProcess accepts duration 0 with a future start, which setProcessDuration refuses:
    // an election with an empty voting window whose batches still settle in its grace window.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_NewProcessRefusesZeroDuration(uint256 startOff, bool nowStart) public {
        uint256 start = nowStart ? 0 : vm.getBlockTimestamp() + bound(startOff, 0, 365 days);
        DAVINCITypes.BallotMode memory mode = _ballotMode();
        DAVINCITypes.Census memory census = _census();
        DAVINCITypes.EncryptionKey memory key = _encKey();
        DAVINCITypes.DKGParams memory dkg = _noDkg();
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        registry.newProcess(
            DAVINCITypes.ProcessStatus.READY, start, 0, MAX_VOTERS, mode, census, METADATA_URI, METADATA_HASH, key, dkg
        );
    }

    /// A process ended by hand in its last second still gets the full window from then.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_EndedByHandWindow(uint256 at) public {
        bytes31 pid = _process(vm.getBlockTimestamp(), DURATION, MAX_VOTERS);
        uint256 start = vm.getBlockTimestamp();
        uint256 t = start + bound(at, 0, DURATION - 1);
        vm.warp(t);
        _setStatus(pid, DAVINCITypes.ProcessStatus.ENDED);
        assertEq(registry.getProcessEndTime(pid), t);
        assertEq(registry.getProcessGraceEnd(pid), t + GRACE);
        vm.warp(t + GRACE - 1);
        assertTrue(_land(pid));
        assertEq(registry.getProcessGraceEnd(pid), t + GRACE - 1 + GRACE);
    }
}
