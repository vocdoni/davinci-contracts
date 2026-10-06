// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Vm} from "forge-std/Vm.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";

/// @dev The grace window past the end (settlement, the idle extension, its cap and freeze,
///      setProcessGrace) and shortening a process with notice. Results gating lives in
///      Results.t.sol and DKG.t.sol.
contract GraceTest is RegistryTestBase {
    /// @dev Struct slot of grace (offset 0) and lastVoteAt (offset 4); processes is at slot 1.
    uint256 internal constant GRACE_SLOT = 26;

    function _end(bytes31 pid) internal view returns (uint256) {
        return registry.getProcessEndTime(pid);
    }

    function _organizer(bytes31 pid, DAVINCITypes.ProcessStatus s) internal {
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, s);
    }

    function _setGrace(bytes31 pid, uint32 g) internal {
        vm.prank(ORGANIZER);
        registry.setProcessGrace(pid, g);
    }

    function _setDuration(bytes31 pid, uint256 d) internal {
        vm.prank(ORGANIZER);
        registry.setProcessDuration(pid, d);
    }

    /// @dev Whether the last recorded logs hold a ProcessDurationChanged.
    function _durationChangedLogged() internal returns (bool) {
        Vm.Log[] memory logs = vm.getRecordedLogs();
        for (uint256 i = 0; i < logs.length; i++) {
            if (logs[i].topics[0] == IProcessRegistry.ProcessDurationChanged.selector) return true;
        }
        return false;
    }

    // --- constructor -------------------------------------------------------------

    function test_Constructor_ExposesTimeConfig() public view {
        assertEq(registry.defaultGrace(), GRACE);
        assertEq(registry.graceFloor(), GRACE_FLOOR);
        assertEq(registry.graceCeil(), GRACE_CEIL);
        assertEq(registry.graceMaxTotal(), GRACE_MAX_TOTAL);
        assertEq(registry.noticeMin(), NOTICE_MIN);
    }

    /// @dev External, so vm.expectRevert has a call to expect: forge compiles a test's `new`
    ///      into a deployCode cheatcode, whose revert ends the test instead.
    function deploy(uint32 d, uint32 f, uint32 c, uint32 m, uint32 n) external returns (ProcessRegistry) {
        return new ProcessRegistry(
            CHAIN_ID,
            address(1),
            BATCH_VK,
            RESULTS_VK,
            ROOT_C,
            bytes32(uint256(1)),
            address(0),
            address(0),
            d,
            f,
            c,
            m,
            n
        );
    }

    function test_Constructor_RejectsBadGrace() public {
        vm.expectRevert(IProcessRegistry.InvalidGrace.selector);
        this.deploy(180, 0, 600, 1800, 60); // floor 0
        vm.expectRevert(IProcessRegistry.InvalidGrace.selector);
        this.deploy(149, 150, 600, 1800, 60); // default below floor
        vm.expectRevert(IProcessRegistry.InvalidGrace.selector);
        this.deploy(601, 150, 600, 1800, 60); // default above ceil
        vm.expectRevert(IProcessRegistry.InvalidGrace.selector);
        this.deploy(180, 150, 1801, 1800, 60); // ceil above the cap
        vm.expectRevert(IProcessRegistry.InvalidGrace.selector);
        this.deploy(180, 150, 600, 1800, 0); // no notice

        // Equal bounds are a fixed grace.
        ProcessRegistry r = this.deploy(10, 10, 10, 10, 1);
        assertEq(r.defaultGrace(), 10);
        assertEq(r.graceMaxTotal(), 10);
        assertEq(r.noticeMin(), 1);
    }

    // --- the window ----------------------------------------------------------------

    function test_NewProcess_StartsWithDefaultGrace() public {
        bytes31 pid = _fixtureProcess();
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.grace, GRACE);
        assertEq(p.lastVoteAt, 0);
        assertEq(registry.getProcessGraceEnd(pid), _end(pid) + GRACE);
    }

    /// @dev A landing before the end leaves the window at end + grace.
    function test_GraceEnd_TransitionBeforeTheEnd() public {
        bytes31 pid = _fixtureProcess();
        vm.warp(_end(pid) - 1000);
        _submit(pid, _transition(0));
        assertEq(registry.getProcess(pid).lastVoteAt, _end(pid) - 1000);
        assertEq(registry.getProcessGraceEnd(pid), _end(pid) + GRACE);
    }

    /// @dev READY past the end settles inside the window, and every landing restarts it.
    function test_GraceEnd_IdleExtension() public {
        bytes31 pid = _fixtureProcess();
        uint256 end = _end(pid);

        vm.warp(end + 100);
        _submit(pid, _transition(0));
        assertEq(registry.getProcess(pid).lastVoteAt, end + 100);
        assertEq(registry.getProcessGraceEnd(pid), end + 100 + GRACE);

        // Past the first window, inside the extended one.
        vm.warp(end + GRACE + 50);
        _submit(pid, _transition(1));
        assertEq(registry.getProcess(pid).lastVoteAt, end + GRACE + 50);
        assertEq(registry.getProcessGraceEnd(pid), end + GRACE + 50 + GRACE);
        assertEq(uint8(registry.getProcess(pid).status), uint8(DAVINCITypes.ProcessStatus.READY));
    }

    /// @dev The window is [start, graceEnd): a valid batch lands at graceEnd - 1, not at graceEnd.
    function test_GraceEnd_Boundary() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        uint256 graceEnd = registry.getProcessGraceEnd(pid);

        vm.warp(graceEnd);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        _submit(pid, _transition(1));

        vm.warp(graceEnd - 1);
        _submit(pid, _transition(1));
        assertEq(registry.getProcess(pid).batchNumber, 2);
    }

    /// @dev Once graceEnd passes nothing moves it again: no landing, no duration, no grace, no
    ///      cap and no status change reopens settlement.
    function test_GraceEnd_NeverReopens() public {
        bytes31 pid = _fixtureProcess();
        uint256 end = _end(pid);
        uint256 graceEnd = registry.getProcessGraceEnd(pid);
        vm.warp(graceEnd + 1000);

        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        _submit(pid, _transition(0));
        vm.startPrank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessDuration(pid, 2 * DURATION);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessGrace(pid, GRACE_CEIL);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessMaxVoters(pid, 1);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        vm.stopPrank();

        // ENDED past the end keeps the end where it was.
        vm.recordLogs();
        _organizer(pid, DAVINCITypes.ProcessStatus.ENDED);
        assertFalse(_durationChangedLogged(), "ENDED past the end moved the end");
        assertEq(_end(pid), end);
        assertEq(registry.getProcessGraceEnd(pid), graceEnd);

        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        _submit(pid, _transition(0));
    }

    /// @dev END inside the window does not move the end either, so it cannot stretch the window.
    function test_GraceEnd_EndedInGraceKeepsTheWindow() public {
        bytes31 pid = _fixtureProcess();
        uint256 end = _end(pid);
        vm.warp(end + 50);
        vm.recordLogs();
        _organizer(pid, DAVINCITypes.ProcessStatus.ENDED);
        assertFalse(_durationChangedLogged(), "ENDED past the end moved the end");
        assertEq(_end(pid), end);
        assertEq(registry.getProcessGraceEnd(pid), end + GRACE);

        // ENDED settles inside the window like READY.
        vm.warp(end + 60);
        _submit(pid, _transition(0));
        assertEq(registry.getProcessGraceEnd(pid), end + 60 + GRACE);
    }

    /// @dev Manual END moves the end to now; the window then runs as for a time end.
    function test_GraceEnd_ManualEnd() public {
        bytes31 pid = _fixtureProcess();
        uint256 endedAt = vm.getBlockTimestamp() + 1000;
        vm.warp(endedAt);
        _organizer(pid, DAVINCITypes.ProcessStatus.ENDED);
        assertEq(_end(pid), endedAt);
        assertEq(registry.getProcessGraceEnd(pid), endedAt + GRACE);

        vm.warp(endedAt + 100);
        _submit(pid, _transition(0));
        uint256 graceEnd = registry.getProcessGraceEnd(pid);
        assertEq(graceEnd, endedAt + 100 + GRACE);

        vm.warp(graceEnd);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        _submit(pid, _transition(1));
    }

    /// @dev A pause blocks settlement only while voting is open: a process paused at the end
    ///      settles through the window like READY, and the window runs as usual.
    function test_GraceEnd_PausedAtTheEndSettles() public {
        bytes31 pid = _fixtureProcess();
        uint256 end = _end(pid);
        _organizer(pid, DAVINCITypes.ProcessStatus.PAUSED);
        vm.warp(end - 1);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        _submit(pid, _transition(0));

        vm.warp(end + 10);
        _submit(pid, _transition(0));
        assertEq(uint8(registry.getProcess(pid).status), uint8(DAVINCITypes.ProcessStatus.PAUSED));
        uint256 graceEnd = registry.getProcessGraceEnd(pid);
        assertEq(graceEnd, end + 10 + GRACE);

        // Resuming past the end is allowed and changes nothing; pausing again is refused.
        _organizer(pid, DAVINCITypes.ProcessStatus.READY);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        assertEq(registry.getProcessGraceEnd(pid), graceEnd);

        vm.warp(graceEnd);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        _submit(pid, _transition(1));
    }

    /// @dev PAUSED to ENDED past the end is allowed, keeps the end and keeps settling.
    function test_GraceEnd_PausedThenEndedPastTheEnd() public {
        bytes31 pid = _fixtureProcess();
        uint256 end = _end(pid);
        _organizer(pid, DAVINCITypes.ProcessStatus.PAUSED);
        vm.warp(end + 10);
        _organizer(pid, DAVINCITypes.ProcessStatus.ENDED);
        assertEq(_end(pid), end);
        assertEq(registry.getProcessGraceEnd(pid), end + GRACE);
        _submit(pid, _transition(0));
        assertEq(registry.getProcessGraceEnd(pid), end + 10 + GRACE);
    }

    /// @dev READY to PAUSED works up to the last second before the end, never from the end on.
    function test_SetProcessStatus_RevertWhen_PausedFromTheEnd() public {
        bytes31 pid = _fixtureProcess();
        uint256 end = _end(pid);
        vm.warp(end - 1);
        _organizer(pid, DAVINCITypes.ProcessStatus.PAUSED);
        _organizer(pid, DAVINCITypes.ProcessStatus.READY);

        vm.warp(end);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        vm.warp(end + GRACE + 1);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        assertEq(uint8(registry.getProcess(pid).status), uint8(DAVINCITypes.ProcessStatus.READY));
    }

    /// @dev A batch of overwrites only (no new voter) is a vote: it passes EmptyTransition,
    ///      lands at a full maxVoters and extends the window like any other.
    function test_GraceEnd_OverwriteOnlyBatch() public {
        bytes31 pid = _fixtureProcess(3);
        _submit(pid, _transition(0)); // votersCount == maxVoters
        uint256 end = _end(pid);
        vm.warp(end + 50);

        Transition memory t = _transition(1);
        vm.expectRevert(IProcessRegistry.MaxVotersReached.selector);
        _submit(pid, t);

        // Relabel every vote of the batch as an overwrite (the PLONK verifier is mocked).
        _setWord(t.publicValues, 19, uint64(t.voters));
        _submit(pid, t);
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.votersCount, 3);
        assertEq(p.overwrittenVotesCount, t.voters);
        assertEq(p.lastVoteAt, end + 50);
        assertEq(registry.getProcessGraceEnd(pid), end + 50 + GRACE);
    }

    /// @dev The cap is fixed at the end: lowering it would strand queued new-voter batches,
    ///      raising it would admit late ones.
    function test_SetProcessMaxVoters_RevertWhen_PastTheEnd() public {
        bytes31 pid = _fixtureProcess();
        _submit(pid, _transition(0));
        uint256 end = _end(pid);
        vm.warp(end - 1);
        vm.prank(ORGANIZER);
        registry.setProcessMaxVoters(pid, MAX_VOTERS - 1);

        vm.warp(end + 10);
        vm.startPrank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessMaxVoters(pid, 3);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessMaxVoters(pid, MAX_VOTERS);
        vm.stopPrank();
        assertEq(registry.getProcess(pid).maxVoters, MAX_VOTERS - 1);

        // The queued new-voter batch lands under the cap voters saw at the close.
        _submit(pid, _transition(1));
        assertEq(registry.getProcess(pid).votersCount, 820);
    }

    function test_GetProcessGraceEnd_UnknownProcess() public view {
        assertEq(registry.getProcessGraceEnd(registry.getNextProcessId(ORGANIZER)), 0);
        assertEq(registry.getProcessGraceEnd(bytes31(0)), 0);
    }

    function test_GraceEnd_CanceledBlocksSettlement() public {
        bytes31 pid = _fixtureProcess();
        _organizer(pid, DAVINCITypes.ProcessStatus.CANCELED);
        vm.warp(_end(pid) + 10);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        _submit(pid, _transition(0));
    }

    /// @dev The formula against a model, over stored grace and lastVoteAt values.
    function testFuzz_GraceEnd_Formula(uint256 duration, uint32 grace, uint64 lastVoteAt) public {
        duration = bound(duration, 1, 3650 days);
        grace = uint32(bound(grace, GRACE_FLOOR, GRACE_CEIL));
        bytes31 pid = _newProcess(block.timestamp, duration, MAX_VOTERS, _ballotMode(), _census());
        uint256 end = _end(pid);
        lastVoteAt = uint64(bound(lastVoteAt, 0, end + GRACE_MAX_TOTAL));

        bytes32 slot = bytes32(uint256(keccak256(abi.encode(pid, uint256(1)))) + GRACE_SLOT);
        vm.store(address(registry), slot, bytes32(uint256(grace) | (uint256(lastVoteAt) << 32)));
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.grace, grace, "grace slot moved");
        assertEq(p.lastVoteAt, lastVoteAt, "lastVoteAt slot moved");

        uint256 idle = (lastVoteAt > end ? lastVoteAt : end) + grace;
        uint256 cap = end + GRACE_MAX_TOTAL;
        assertEq(registry.getProcessGraceEnd(pid), idle < cap ? idle : cap);
    }

    /// @dev setProcessDuration allows an end at 2^256 - 1; the window must not overflow.
    function test_GraceEnd_MaxEndDoesNotOverflow() public {
        bytes31 pid = _fixtureProcess();
        _setDuration(pid, type(uint256).max - vm.getBlockTimestamp());
        assertEq(registry.getProcessGraceEnd(pid), type(uint256).max);
        _submit(pid, _transition(0));
        assertEq(registry.getProcessGraceEnd(pid), type(uint256).max);
    }

    // --- setProcessGrace -----------------------------------------------------------

    function test_SetProcessGrace() public {
        bytes31 pid = _fixtureProcess();
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessGraceChanged(pid, GRACE_FLOOR);
        _setGrace(pid, GRACE_FLOOR);
        assertEq(registry.getProcess(pid).grace, GRACE_FLOOR);
        assertEq(registry.getProcessGraceEnd(pid), _end(pid) + GRACE_FLOOR);

        // PAUSED counts as open too.
        _organizer(pid, DAVINCITypes.ProcessStatus.PAUSED);
        _setGrace(pid, GRACE_CEIL);
        assertEq(registry.getProcessGraceEnd(pid), _end(pid) + GRACE_CEIL);

        // Up to the last second before the end.
        vm.warp(_end(pid) - 1);
        _setGrace(pid, GRACE);
        assertEq(registry.getProcess(pid).grace, GRACE);
    }

    function test_SetProcessGrace_RevertWhen_OutOfBounds() public {
        bytes31 pid = _fixtureProcess();
        vm.startPrank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidGrace.selector);
        registry.setProcessGrace(pid, 0);
        vm.expectRevert(IProcessRegistry.InvalidGrace.selector);
        registry.setProcessGrace(pid, GRACE_FLOOR - 1);
        vm.expectRevert(IProcessRegistry.InvalidGrace.selector);
        registry.setProcessGrace(pid, GRACE_CEIL + 1);
        vm.stopPrank();
        assertEq(registry.getProcess(pid).grace, GRACE);
    }

    function test_SetProcessGrace_RevertWhen_NotOrganizer() public {
        bytes31 pid = _fixtureProcess();
        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        registry.setProcessGrace(pid, GRACE_FLOOR);
    }

    /// @dev Only while the process is open: past the end the window is already running.
    function test_SetProcessGrace_RevertWhen_NotOpen() public {
        bytes31 pid = _fixtureProcess();
        bytes31 paused = _newProcess(0, DURATION, MAX_VOTERS, _ballotMode(), _census());
        _organizer(paused, DAVINCITypes.ProcessStatus.PAUSED);
        vm.warp(_end(pid));
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessGrace(pid, GRACE_FLOOR);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessGrace(paused, GRACE_FLOOR);

        _organizer(pid, DAVINCITypes.ProcessStatus.ENDED);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessGrace(pid, GRACE_FLOOR);

        bytes31 canceled = _newProcess(0, DURATION, MAX_VOTERS, _ballotMode(), _census());
        _organizer(canceled, DAVINCITypes.ProcessStatus.CANCELED);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessGrace(canceled, GRACE_FLOOR);

        bytes31 unknown = registry.getNextProcessId(ORGANIZER);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.ProcessNotFound.selector);
        registry.setProcessGrace(unknown, GRACE_FLOOR);
    }

    // --- shortening with notice ------------------------------------------------------

    /// @dev The chair's "voting closes in one minute": the end moves to now + noticeMin and
    ///      the grace window follows it.
    function test_SetProcessDuration_ShortenWithNotice() public {
        bytes31 pid = _fixtureProcess();
        uint256 start = registry.getProcess(pid).startTime;
        uint256 now_ = start + 1000;
        vm.warp(now_);
        uint256 newEnd = now_ + NOTICE_MIN;

        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessDurationChanged(pid, newEnd - start);
        _setDuration(pid, newEnd - start);
        assertEq(_end(pid), newEnd);
        assertEq(registry.getProcessGraceEnd(pid), newEnd + GRACE);

        // Settles past the new end, inside its window, and not after it.
        vm.warp(newEnd + 10);
        _submit(pid, _transition(0));
        uint256 graceEnd = registry.getProcessGraceEnd(pid);
        assertEq(graceEnd, newEnd + 10 + GRACE);
        vm.warp(graceEnd);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        _submit(pid, _transition(1));
    }

    function test_SetProcessDuration_RevertWhen_ShortNotice() public {
        bytes31 pid = _fixtureProcess();
        uint256 start = registry.getProcess(pid).startTime;
        uint256 now_ = start + 1000;
        vm.warp(now_);

        vm.startPrank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        registry.setProcessDuration(pid, now_ + NOTICE_MIN - 1 - start);
        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        registry.setProcessDuration(pid, now_ - start); // ends now
        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        registry.setProcessDuration(pid, 1); // in the past
        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        registry.setProcessDuration(pid, 0);
        // PAUSED shortens with the same notice.
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        registry.setProcessDuration(pid, now_ + NOTICE_MIN - 1 - start);
        registry.setProcessDuration(pid, now_ + NOTICE_MIN - start);
        vm.stopPrank();
        assertEq(_end(pid), now_ + NOTICE_MIN);
    }

    /// @dev Before the start the new end only has to be noticeMin away (and after the start).
    function test_SetProcessDuration_ShortenBeforeStart() public {
        uint256 now_ = vm.getBlockTimestamp();
        bytes31 later = _newProcess(now_ + 1000, DURATION, MAX_VOTERS, _ballotMode(), _census());
        _setDuration(later, 1);
        assertEq(_end(later), now_ + 1001);

        bytes31 soon = _newProcess(now_ + 30, DURATION, MAX_VOTERS, _ballotMode(), _census());
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        registry.setProcessDuration(soon, 1); // now + 31 < now + noticeMin
        _setDuration(soon, NOTICE_MIN - 30);
        assertEq(_end(soon), now_ + NOTICE_MIN);
    }

    function test_SetProcessDuration_RevertWhen_ShortenAfterEnd() public {
        bytes31 pid = _fixtureProcess();
        uint256 end = _end(pid);
        vm.warp(end + 10); // inside the grace window
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessDuration(pid, DURATION - 1000);
        assertEq(_end(pid), end);
    }

    /// @dev Extending needs no notice, up to the last second; the same end is no change.
    function test_SetProcessDuration_ExtendUnchanged() public {
        bytes31 pid = _fixtureProcess();
        vm.warp(_end(pid) - 1);
        _setDuration(pid, DURATION + 1);
        assertEq(registry.getProcess(pid).duration, DURATION + 1);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        registry.setProcessDuration(pid, DURATION + 1);
    }
}

/// @dev The anvil/e2e bounds (default 10 s, floor 2, ceil 60, cap 60, notice 5): short enough
///      for landings to reach the cap.
contract GraceCapTest is RegistryTestBase {
    using stdJson for string;

    function _timeConfig() internal pure override returns (uint32, uint32, uint32, uint32, uint32) {
        return (10, 2, 60, 60, 5);
    }

    /// @dev A sequencer trickling batches cannot hold the window open past end + graceMaxTotal,
    ///      and the results open there.
    function test_GraceEnd_CappedAtMaxTotal() public {
        bytes31 pid = _fixtureProcess();
        uint256 end = registry.getProcessEndTime(pid);
        vm.prank(ORGANIZER);
        registry.setProcessGrace(pid, 60);

        vm.warp(end + 30);
        _submit(pid, _transition(0));
        assertEq(registry.getProcessGraceEnd(pid), end + 60); // not end + 90
        vm.warp(end + 59);
        _submit(pid, _transition(1));
        assertEq(registry.getProcessGraceEnd(pid), end + 60); // not end + 119

        bytes memory pv = fixture.readBytes(".results.public_values");
        bytes memory proof = fixture.readBytes(".results.proof_bytes");
        vm.expectRevert(IProcessRegistry.GraceOpen.selector);
        registry.setProcessResults(pid, pv, proof);
        vm.warp(end + 60);
        registry.setProcessResults(pid, pv, proof);
        assertEq(uint8(registry.getProcess(pid).status), uint8(DAVINCITypes.ProcessStatus.RESULTS));
    }

    function test_ShortenWithNotice_AnvilNotice() public {
        bytes31 pid = _fixtureProcess();
        uint256 start = registry.getProcess(pid).startTime;
        vm.warp(start + 100);
        vm.startPrank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        registry.setProcessDuration(pid, 104);
        registry.setProcessDuration(pid, 105);
        vm.stopPrank();
        assertEq(registry.getProcessGraceEnd(pid), start + 105 + 10);
    }
}
