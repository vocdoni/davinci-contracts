// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {StdInvariant} from "forge-std/StdInvariant.sol";
import {AttackBase, AttackLib} from "./AttackBase.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";

/// @dev Random sequences of organizer, stranger, sequencer and results actions against a
///      reference model of the lifecycle, the grace window and settlement. Every call's
///      outcome must match the model's prediction; a mismatch is counted, not reverted, so
///      the invariant reports it.
contract LifecycleHandler is Test {
    uint8 internal constant READY = 0;
    uint8 internal constant ENDED = 1;
    uint8 internal constant CANCELED = 2;
    uint8 internal constant PAUSED = 3;
    uint8 internal constant RESULTS = 4;
    uint256 internal constant CAP = 1_000_000_000_000;
    address internal constant STRANGER = address(0xBAD);

    struct M {
        address org;
        uint8 status;
        uint256 start;
        uint256 duration;
        uint256 maxVoters;
        uint256 maxValue;
        uint256 votersCount;
        uint256 overwritten;
        uint256 batches;
        uint64 last;
        uint32 grace;
        bytes32 root;
        bytes32 census;
        bool dynamic;
        bool everClosed;
        bytes32 closedRoot;
    }

    ProcessRegistry public immutable registry;
    uint32 internal immutable floor_;
    uint32 internal immutable ceil_;
    uint32 internal immutable maxTotal;
    uint32 internal immutable notice;

    bytes31[] public pids;
    mapping(bytes31 => M) internal model;

    uint256 public violations;
    string public firstViolation;
    mapping(string => uint256) public calls;
    mapping(string => uint256) public successes;

    constructor(ProcessRegistry r, bytes31[] memory ps) {
        registry = r;
        floor_ = r.graceFloor();
        ceil_ = r.graceCeil();
        maxTotal = r.graceMaxTotal();
        notice = r.noticeMin();
        for (uint256 i = 0; i < ps.length; ++i) {
            pids.push(ps[i]);
            DAVINCITypes.Process memory p = r.getProcess(ps[i]);
            M storage m = model[ps[i]];
            m.org = p.organizationId;
            m.status = uint8(p.status);
            m.start = p.startTime;
            m.duration = p.duration;
            m.maxVoters = p.maxVoters;
            m.maxValue = p.ballotMode.maxValue;
            m.grace = p.grace;
            m.root = p.latestStateRoot;
            m.census = p.census.censusRoot;
            m.dynamic = p.census.censusOrigin == DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1;
        }
    }

    // --- model ---------------------------------------------------------------------

    function modelOf(bytes31 pid) external view returns (M memory) {
        return model[pid];
    }

    function pidCount() external view returns (uint256) {
        return pids.length;
    }

    function _end(M storage m) internal view returns (uint256) {
        return m.start + m.duration;
    }

    function graceEndOf(bytes31 pid) public view returns (uint256) {
        M storage m = model[pid];
        uint256 end = _end(m);
        if (end > type(uint256).max - maxTotal) return type(uint256).max;
        uint256 idle = (m.last > end ? m.last : end) + m.grace;
        uint256 cap = end + maxTotal;
        return idle < cap ? idle : cap;
    }

    function _open(M storage m) internal view returns (bool) {
        return (m.status == READY || m.status == PAUSED) && vm.getBlockTimestamp() < _end(m);
    }

    function _check(bool ok, bool want, string memory what) internal {
        calls[what]++;
        if (ok) successes[what]++;
        if (ok != want) {
            if (violations == 0) {
                firstViolation = string.concat(what, ok ? ": succeeded, model refuses" : ": failed, model accepts");
            }
            violations++;
        }
    }

    function _pid(uint256 i) internal view returns (bytes31) {
        return pids[i % pids.length];
    }

    function _observe() internal {
        for (uint256 i = 0; i < pids.length; ++i) {
            M storage m = model[pids[i]];
            if (!m.everClosed && vm.getBlockTimestamp() >= graceEndOf(pids[i])) {
                m.everClosed = true;
                m.closedRoot = m.root;
            }
        }
    }

    // --- time ------------------------------------------------------------------------

    function warp(uint256 dt) external {
        dt = dt % 16 == 0 ? bound(dt, 0, 1 days) : bound(dt, 0, 900);
        vm.warp(vm.getBlockTimestamp() + dt);
        vm.roll(vm.getBlockNumber() + 1);
        _observe();
    }

    // --- organizer and stranger ------------------------------------------------------

    function _statusMoveOk(uint8 from, uint8 to) internal pure returns (bool) {
        if (from == to || from == CANCELED || from == RESULTS || from == ENDED) return false;
        if (from == READY) return to == PAUSED || to == CANCELED || to == ENDED;
        return to == READY || to == CANCELED || to == ENDED; // PAUSED
    }

    function setStatus(uint256 pi, uint8 s, bool stranger) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        s = uint8(bound(s, 0, 4));
        bool timeOk = true;
        if (s == ENDED) timeOk = vm.getBlockTimestamp() >= m.start;
        if (s == PAUSED) timeOk = vm.getBlockTimestamp() < _end(m);
        // A cancel must not be able to veto a tally that can already be published.
        if (s == CANCELED) timeOk = vm.getBlockTimestamp() < graceEndOf(pid);
        bool want = !stranger && _statusMoveOk(m.status, s) && timeOk;
        vm.prank(stranger ? STRANGER : m.org);
        (bool ok,) =
            address(registry).call(abi.encodeCall(registry.setProcessStatus, (pid, DAVINCITypes.ProcessStatus(s))));
        _check(ok, want, "setProcessStatus");
        if (ok) {
            if (s == ENDED && vm.getBlockTimestamp() < _end(m)) m.duration = vm.getBlockTimestamp() - m.start;
            m.status = s;
        }
    }

    function setDuration(uint256 pi, uint256 d, bool stranger) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        if (d % 50 == 0) d = type(uint256).max - bound(d, 0, 10);
        else if (d % 7 == 0) d = 0;
        else d = bound(d, 1, 4 hours);
        bool want = !stranger && _open(m) && d != 0 && d <= type(uint256).max - m.start;
        if (want) {
            uint256 newEnd = m.start + d;
            uint256 end = _end(m);
            want = newEnd > vm.getBlockTimestamp() && newEnd != end
                && !(newEnd < end && newEnd < vm.getBlockTimestamp() + notice);
        }
        vm.prank(stranger ? STRANGER : m.org);
        (bool ok,) = address(registry).call(abi.encodeCall(registry.setProcessDuration, (pid, d)));
        _check(ok, want, "setProcessDuration");
        if (ok) m.duration = d;
    }

    function setGrace(uint256 pi, uint32 g, bool stranger) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        g = uint32(bound(g, 0, uint256(ceil_) + 60));
        bool want = !stranger && _open(m) && g >= floor_ && g <= ceil_;
        vm.prank(stranger ? STRANGER : m.org);
        (bool ok,) = address(registry).call(abi.encodeCall(registry.setProcessGrace, (pid, g)));
        _check(ok, want, "setProcessGrace");
        if (ok) m.grace = g;
    }

    function setMaxVoters(uint256 pi, uint256 mv, bool stranger) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        mv = mv % 20 == 0 ? bound(mv, 0, type(uint256).max) : bound(mv, 0, m.votersCount + 60);
        bool want = !stranger && _open(m) && mv != 0 && mv >= m.votersCount && m.maxValue <= CAP / mv;
        vm.prank(stranger ? STRANGER : m.org);
        (bool ok,) = address(registry).call(abi.encodeCall(registry.setProcessMaxVoters, (pid, mv)));
        _check(ok, want, "setProcessMaxVoters");
        if (ok) m.maxVoters = mv;
    }

    function setCensus(uint256 pi, bytes32 root, bool otherOrigin, bool stranger) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        if (uint256(root) % 9 == 0) root = bytes32(0);
        DAVINCITypes.Census memory c;
        c.censusOrigin = otherOrigin
            ? DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
            : DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1;
        c.censusRoot = root;
        c.censusURI = "https://example.com/census2";
        bool want = !stranger && m.dynamic && !otherOrigin && root != 0 && _open(m);
        vm.prank(stranger ? STRANGER : m.org);
        (bool ok,) = address(registry).call(abi.encodeCall(registry.setProcessCensus, (pid, c)));
        _check(ok, want, "setProcessCensus");
        if (ok) m.census = root;
    }

    // --- sequencer -------------------------------------------------------------------

    function submit(uint256 pi, uint64 voters, uint64 overwrites, uint256 tamper, uint256 salt) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        voters = uint64(bound(voters, 0, 40));
        uint256 maxOw = voters < m.votersCount ? voters : m.votersCount;
        overwrites = uint64(bound(overwrites, 0, maxOw));
        tamper = bound(tamper, 0, 14);
        bytes32 rootAfter = keccak256(abi.encode(pid, m.batches, salt));
        AttackLib.Synth memory s = AttackLib.synth(
            m.root, rootAfter, voters, overwrites, m.census, uint64(m.votersCount), 1 + (salt % 3), salt
        );
        if (tamper == 5) AttackLib.setReg32(s.pv, 2, keccak256(abi.encode("wrong root", salt)));
        if (tamper == 6) AttackLib.setWord(s.pv, 42, uint64(m.votersCount + 1));
        if (tamper == 7) AttackLib.setReg32(s.pv, 20, AttackLib.rev(bytes32(uint256(m.census) ^ 1)));
        if (tamper == 8) {
            AttackLib.setWord(s.pv, 18, 0);
            AttackLib.setWord(s.pv, 19, 0);
        }
        if (tamper == 9) AttackLib.setWord(s.pv, 0, 0);
        if (tamper == 10) AttackLib.setWord(s.pv, 1, 1);
        bytes32[] memory vhs = s.vhs;
        if (tamper == 11) {
            vhs = new bytes32[](s.vhs.length + 1);
            for (uint256 i = 0; i < s.vhs.length; ++i) {
                vhs[i] = s.vhs[i];
            }
            vhs[s.vhs.length] = keccak256("extra blob");
        }
        if (tamper == 12) {
            vhs = new bytes32[](s.vhs.length - 1);
            for (uint256 i = 0; i < vhs.length; ++i) {
                vhs[i] = s.vhs[i];
            }
        }
        if (tamper == 13) s.ys[0] = bytes32(uint256(s.ys[0]) ^ 1);
        bool cross;
        if (tamper == 14) {
            bytes32 other = model[_pid(pi % pids.length + 1)].root;
            AttackLib.setReg32(s.pv, 2, other);
            cross = other == m.root;
        }

        uint256 end = _end(m);
        bool statusOk = m.status == READY || m.status == ENDED || (m.status == PAUSED && vm.getBlockTimestamp() >= end);
        bool want = (tamper < 5 || cross) && statusOk && vm.getBlockTimestamp() >= m.start
            && vm.getBlockTimestamp() < graceEndOf(pid) && voters > 0
            && m.votersCount + voters - overwrites <= m.maxVoters;

        vm.mockCall(address(0x0A), bytes(""), abi.encode(uint256(4096), AttackLib.BLS_MODULUS));
        vm.blobhashes(vhs);
        (bool ok,) = address(registry)
            .call(abi.encodeCall(registry.submitStateTransition, (pid, s.pv, hex"", s.commitments, s.ys, s.kzgProofs)));
        _check(ok, want, "submitStateTransition");
        if (ok) {
            if (m.everClosed) _check(true, false, "settled after the window closed");
            m.root = rootAfter;
            m.votersCount += voters - overwrites;
            m.overwritten += overwrites;
            m.batches++;
            m.last = uint64(vm.getBlockTimestamp());
        }
    }

    // --- results ---------------------------------------------------------------------

    function results(uint256 pi, bool wrongRoot, uint64[16] memory values) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        bytes32 root = wrongRoot ? keccak256(abi.encode(m.root)) : m.root;
        bool want = !wrongRoot && m.status != CANCELED && m.status != RESULTS
            && (m.status == ENDED || vm.getBlockTimestamp() >= _end(m)) && vm.getBlockTimestamp() >= graceEndOf(pid);
        (bool ok,) = address(registry)
            .call(abi.encodeCall(registry.setProcessResults, (pid, AttackLib.resultsPublics(root, values), hex"")));
        _check(ok, want, "setProcessResults");
        if (ok) m.status = RESULTS;
    }

    /// The DKG entry points never apply to a sequencer-key process.
    function dkgCalls(uint256 pi, uint256 sk) external {
        bytes31 pid = _pid(pi);
        uint256[64] memory acc;
        (bool ok,) =
            address(registry).call(abi.encodeCall(registry.requestResultsDecryption, (pid, acc, new bytes32[](1))));
        _check(ok, false, "requestResultsDecryption");
        (ok,) = address(registry).call(abi.encodeCall(registry.finalizeResultsFromDKG, (pid)));
        _check(ok, false, "finalizeResultsFromDKG");
        (ok,) = address(registry).call(abi.encodeCall(registry.revealProcessKey, (pid, sk)));
        _check(ok, false, "revealProcessKey");
    }
}

contract AttackLifecycleInvariantTest is AttackBase {
    address internal constant ORG2 = address(0xA11CE);

    LifecycleHandler internal handler;

    function setUp() public override {
        super.setUp();
        bytes31[] memory ps = new bytes31[](3);
        ps[0] = _process(vm.getBlockTimestamp(), 2 hours, 200);

        DAVINCITypes.Census memory dyn = _census();
        dyn.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1;
        vm.prank(ORGANIZER);
        ps[1] = registry.newProcess(
            DAVINCITypes.ProcessStatus.PAUSED,
            vm.getBlockTimestamp() + 30 minutes,
            1 hours,
            100,
            _ballotMode(),
            dyn,
            METADATA_URI,
            METADATA_HASH,
            _encKey(),
            _noDkg()
        );

        vm.prank(ORG2);
        ps[2] = registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            vm.getBlockTimestamp() + 10 minutes,
            30 minutes,
            50,
            _ballotMode(),
            _census(),
            METADATA_URI,
            METADATA_HASH,
            _encKey(),
            _noDkg()
        );

        handler = new LifecycleHandler(registry, ps);
        bytes4[] memory sel = new bytes4[](10);
        sel[0] = LifecycleHandler.warp.selector;
        sel[1] = LifecycleHandler.warp.selector;
        sel[2] = LifecycleHandler.setStatus.selector;
        sel[3] = LifecycleHandler.setDuration.selector;
        sel[4] = LifecycleHandler.setGrace.selector;
        sel[5] = LifecycleHandler.setMaxVoters.selector;
        sel[6] = LifecycleHandler.setCensus.selector;
        sel[7] = LifecycleHandler.submit.selector;
        sel[8] = LifecycleHandler.results.selector;
        sel[9] = LifecycleHandler.dkgCalls.selector;
        targetSelector(StdInvariant.FuzzSelector({addr: address(handler), selectors: sel}));
        targetContract(address(handler));
    }

    /// forge-config: default.invariant.runs = 128
    /// forge-config: default.invariant.depth = 120
    /// forge-config: default.invariant.fail-on-revert = true
    function invariant_EveryCallMatchesTheModel() public view {
        assertEq(handler.violations(), 0, handler.firstViolation());
    }

    /// forge-config: default.invariant.runs = 128
    /// forge-config: default.invariant.depth = 120
    /// forge-config: default.invariant.fail-on-revert = true
    function invariant_StateMatchesTheModel() public view {
        for (uint256 i = 0; i < handler.pidCount(); ++i) {
            bytes31 pid = handler.pids(i);
            LifecycleHandler.M memory m = handler.modelOf(pid);
            DAVINCITypes.Process memory p = registry.getProcess(pid);
            assertEq(uint8(p.status), m.status, "status");
            assertEq(p.duration, m.duration, "duration");
            assertEq(p.maxVoters, m.maxVoters, "maxVoters");
            assertEq(p.votersCount, m.votersCount, "votersCount");
            assertEq(p.overwrittenVotesCount, m.overwritten, "overwrites");
            assertEq(p.batchNumber, m.batches, "batchNumber");
            assertEq(p.lastVoteAt, m.last, "lastVoteAt");
            assertEq(p.grace, m.grace, "grace");
            assertEq(p.latestStateRoot, m.root, "root");
            assertEq(p.census.censusRoot, m.census, "census");
            assertEq(registry.getProcessGraceEnd(pid), handler.graceEndOf(pid), "grace end");
            assertLe(p.votersCount, p.maxVoters, "voters above the cap");
            if (m.everClosed) assertEq(p.latestStateRoot, m.closedRoot, "root moved after the window closed");
            uint256 end = p.startTime + p.duration;
            if (end <= type(uint256).max - registry.graceMaxTotal()) {
                assertLe(registry.getProcessGraceEnd(pid), end + registry.graceMaxTotal(), "grace past the cap");
            }
        }
    }
}
