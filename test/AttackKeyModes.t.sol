// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {StdInvariant} from "forge-std/StdInvariant.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {AttackBase, AttackLib} from "./AttackBase.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {Sha256SmtLib} from "../src/libraries/Sha256SmtLib.sol";
import {BjjFormLib} from "../src/libraries/BjjFormLib.sol";
import {CouncilAdapter} from "../src/CouncilAdapter.sol";
import {MockDKG} from "./mocks/MockDKG.sol";
import {MockCouncilManager} from "./mocks/MockCouncilManager.sol";

/// @dev Random organizer, stranger, sequencer, requester, committee and finalizer actions on
///      DKG_AUTOMATIC, DKG_LOCKED and COUNCIL processes, against a reference model of the
///      lifecycle, the grace window, the results request and the results vector. Batches
///      settle state roots that commit to a model accumulator (leaf 0x04 under one sibling),
///      so requestResultsDecryption can be driven with honest and tampered arguments.
contract KeyModesHandler is Test {
    uint8 internal constant READY = 0;
    uint8 internal constant ENDED = 1;
    uint8 internal constant CANCELED = 2;
    uint8 internal constant PAUSED = 3;
    uint8 internal constant RESULTS = 4;
    uint256 internal constant Q = BjjFormLib.Q;
    address internal constant STRANGER = address(0xBAD);
    uint256 internal constant ORG_SK = 12345;

    struct M {
        address org;
        uint8 mode; // KeyMode
        bytes12 epoch; // dkgEpochId: the DKG epoch or the Council ceremony
        uint8 status;
        uint8 numFields;
        uint256 start;
        uint256 duration;
        uint256 maxVoters;
        uint256 votersCount;
        uint256 batches;
        uint64 last;
        uint32 grace;
        bytes32 root;
        bytes32 census;
        bytes32 sib; // the one sibling of leaf 0x04 under root (0 at genesis: no proof built)
        uint256[64] acc;
        bool requested;
        bytes32 requestedRoot;
        uint16 zeroSkipped;
        uint8 count;
        bool revealed;
        bool gateWasOpenAtResults;
        uint256[16] plain; // plaintext of active field j (in submission order)
        uint16 plainSet; // bit j: active field j combined
    }

    ProcessRegistry public immutable registry;
    MockDKG public immutable dkg;
    MockCouncilManager public immutable council;
    bytes12 public immutable manualCid;
    uint256 internal immutable maxTotal;

    uint256[2][] internal points; // circomlib TE points, on curve, not the identity
    bytes31[] public pids;
    mapping(bytes31 => M) internal model;
    bool public gateOpened;

    uint256 public violations;
    string public firstViolation;
    mapping(string => uint256) public calls;
    mapping(string => uint256) public successes;

    constructor(
        ProcessRegistry r,
        MockDKG d,
        MockCouncilManager c,
        bytes12 cid,
        bytes31[] memory ps,
        uint256[2][] memory pts
    ) {
        registry = r;
        dkg = d;
        council = c;
        manualCid = cid;
        maxTotal = r.graceMaxTotal();
        for (uint256 i = 0; i < pts.length; ++i) {
            points.push(pts[i]);
        }
        for (uint256 i = 0; i < ps.length; ++i) {
            pids.push(ps[i]);
            DAVINCITypes.Process memory p = r.getProcess(ps[i]);
            M storage m = model[ps[i]];
            m.org = p.organizationId;
            m.mode = uint8(p.keyMode);
            m.epoch = p.dkgEpochId;
            m.status = uint8(p.status);
            m.numFields = p.ballotMode.numFields;
            m.start = p.startTime;
            m.duration = p.duration;
            m.maxVoters = p.maxVoters;
            m.grace = p.grace;
            m.root = p.latestStateRoot;
            m.census = p.census.censusRoot;
            for (uint256 f = 0; f < 16; ++f) {
                m.acc[4 * f + 1] = 1;
                m.acc[4 * f + 3] = 1;
            }
        }
    }

    // --- model ---------------------------------------------------------------------

    function pidCount() external view returns (uint256) {
        return pids.length;
    }

    function accOf(bytes31 pid) external view returns (uint256[64] memory) {
        return model[pid].acc;
    }

    function modelOf(bytes31 pid)
        external
        view
        returns (uint8 status, bytes32 root, bool requested, uint8 count, uint16 zeroSkipped, uint256 voters)
    {
        M storage m = model[pid];
        return (m.status, m.root, m.requested, m.count, m.zeroSkipped, m.votersCount);
    }

    /// @dev The results vector the model expects: 0 for skipped fields, else the plaintext
    ///      of the next active field.
    function expectedResult(bytes31 pid) public view returns (uint256[] memory r) {
        M storage m = model[pid];
        r = new uint256[](m.numFields);
        uint256 j;
        for (uint256 i = 0; i < m.numFields; ++i) {
            if ((m.zeroSkipped >> i) & 1 == 1) continue;
            r[i] = m.plain[j++];
        }
    }

    function gateWasOpenAtResults(bytes31 pid) external view returns (bool) {
        return model[pid].gateWasOpenAtResults;
    }

    function requestedRoot(bytes31 pid) external view returns (bytes32) {
        return model[pid].requestedRoot;
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

    function _gateOpen(M storage m) internal view returns (bool) {
        if (m.mode != uint8(DAVINCITypes.KeyMode.COUNCIL)) return true;
        return council.isDecryptionOpen(m.epoch);
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

    // --- time ----------------------------------------------------------------------

    function warp(uint256 dt) external {
        dt = dt % 16 == 0 ? bound(dt, 0, 1 days) : bound(dt, 0, 900);
        vm.warp(vm.getBlockTimestamp() + dt);
        vm.roll(vm.getBlockNumber() + 1);
    }

    // --- organizer and stranger ------------------------------------------------------

    function _statusMoveOk(uint8 from, uint8 to) internal pure returns (bool) {
        if (from == to || from == CANCELED || from == RESULTS || from == ENDED) return false;
        if (from == READY) return to == PAUSED || to == CANCELED || to == ENDED;
        return to == READY || to == CANCELED || to == ENDED;
    }

    function setStatus(uint256 pi, uint8 s, bool stranger) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        s = uint8(bound(s, 0, 4));
        bool timeOk = true;
        if (s == ENDED) timeOk = vm.getBlockTimestamp() >= m.start;
        if (s == PAUSED) timeOk = vm.getBlockTimestamp() < _end(m);
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

    // --- sequencer -------------------------------------------------------------------

    /// @dev A pseudo-random accumulator: each declared field identity or active (a pair of
    ///      curve points), and now and then a half-identity field the guest never produces.
    ///      Fields past numFields stay identity, as the guest leaves them.
    function _randomAcc(uint256 seed, uint8 numFields) internal view returns (uint256[64] memory acc) {
        for (uint256 f = 0; f < 16; ++f) {
            uint256 r = uint256(keccak256(abi.encode(seed, f)));
            uint256 kind = f < numFields ? r % 10 : 0;
            if (kind < 4) {
                (acc[4 * f], acc[4 * f + 1], acc[4 * f + 2], acc[4 * f + 3]) = (0, 1, 0, 1);
            } else if (kind == 9 && r % 7 == 0) {
                uint256[2] memory pt = points[(r >> 8) % points.length];
                (acc[4 * f], acc[4 * f + 1], acc[4 * f + 2], acc[4 * f + 3]) = (pt[0], pt[1], 0, 1);
            } else {
                uint256[2] memory c1 = points[(r >> 8) % points.length];
                uint256[2] memory c2 = points[(r >> 16) % points.length];
                (acc[4 * f], acc[4 * f + 1], acc[4 * f + 2], acc[4 * f + 3]) = (c1[0], c1[1], c2[0], c2[1]);
            }
        }
    }

    /// @dev Root of a tree whose results leaf is acc, with sibling sib at level 0 (bit 0 of
    ///      key 4 is 0, so the leaf is the left child).
    function _rootOf(uint256[64] memory acc, bytes32 sib) internal pure returns (bytes32) {
        bytes32 leaf = Sha256SmtLib.leafHash(4, uint256(sha256(abi.encode(acc))));
        return sha256(abi.encodePacked(leaf, sib));
    }

    function submit(uint256 pi, uint64 voters, uint256 accSeed) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        voters = uint64(bound(voters, 0, 30));
        uint256[64] memory acc = _randomAcc(accSeed, m.numFields);
        bytes32 sib = bytes32(uint256(keccak256(abi.encode("sib", pid, m.batches, accSeed))) | 1);
        bytes32 rootAfter = _rootOf(acc, sib);
        AttackLib.Synth memory s =
            AttackLib.synth(m.root, rootAfter, voters, 0, m.census, uint64(m.votersCount), 1 + (accSeed % 2), accSeed);

        bool statusOk =
            m.status == READY || m.status == ENDED || (m.status == PAUSED && vm.getBlockTimestamp() >= _end(m));
        bool want = statusOk && vm.getBlockTimestamp() >= m.start && vm.getBlockTimestamp() < graceEndOf(pid)
            && voters > 0 && m.votersCount + voters <= m.maxVoters;

        vm.mockCall(address(0x0A), bytes(""), abi.encode(uint256(4096), AttackLib.BLS_MODULUS));
        vm.blobhashes(s.vhs);
        (bool ok,) = address(registry)
            .call(abi.encodeCall(registry.submitStateTransition, (pid, s.pv, hex"", s.commitments, s.ys, s.kzgProofs)));
        _check(ok, want, "submitStateTransition");
        if (ok) {
            if (m.requested) _check(true, false, "settled after the results request");
            m.root = rootAfter;
            m.sib = sib;
            m.acc = acc;
            m.votersCount += voters;
            m.batches++;
            m.last = uint64(vm.getBlockTimestamp());
        }
    }

    // --- results request -------------------------------------------------------------

    function _shape(M storage m) internal view returns (bool ok, uint16 zeroSkipped, uint8 n) {
        for (uint256 i = 0; i < m.numFields; ++i) {
            bool c1Id = m.acc[4 * i] == 0 && m.acc[4 * i + 1] == 1;
            bool c2Id = m.acc[4 * i + 2] == 0 && m.acc[4 * i + 3] == 1;
            if (c1Id != c2Id) return (false, 0, 0);
            if (c1Id) zeroSkipped |= uint16(1 << i);
            else ++n;
        }
        ok = true;
    }

    /// variant: 0 honest, 1 honest with zero-padded siblings, 2 a coordinate + 1, 3 a
    /// coordinate + Q, 4 a tampered sibling, 5 no siblings, 6 a non-zero last sibling,
    /// 7 66 siblings, 8 a field past numFields changed, 9 two fields swapped, 10 honest
    /// with 65 siblings (the most a proof carries).
    function request(uint256 pi, uint256 variant, uint256 where) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        variant = bound(variant, 0, 10);
        uint256[64] memory acc = m.acc;
        bytes32[] memory sibs =
            new bytes32[](variant == 1 ? 5 : variant == 7 ? 66 : variant == 10 ? 65 : variant == 5 ? 0 : 2);
        if (sibs.length > 0) sibs[0] = m.sib;
        if (variant == 7) sibs[63] = m.sib;
        uint256 w = bound(where, 0, 63);
        if (variant == 2) acc[w] = acc[w] + 1 < Q ? acc[w] + 1 : 0;
        if (variant == 3) acc[w] = acc[w] + Q;
        if (variant == 4) sibs[0] = m.sib ^ bytes32(uint256(1) << (w % 256));
        if (variant == 6) sibs[1] = m.sib;
        if (variant == 8) {
            if (m.numFields == 16) variant = 0; // no field past numFields: honest
            else acc[4 * m.numFields] = acc[4 * m.numFields] + 1;
        }
        if (variant == 9) {
            uint256 a = w % 16;
            uint256 b = (a + 1) % 16;
            bool same = true;
            for (uint256 k = 0; k < 4; ++k) {
                (acc[4 * a + k], acc[4 * b + k]) = (acc[4 * b + k], acc[4 * a + k]);
                if (acc[4 * a + k] != acc[4 * b + k]) same = false;
            }
            if (same) variant = 0;
        }

        (bool shapeOk, uint16 zs, uint8 n) = _shape(m);
        bool honest = (variant <= 1 || variant == 10) && m.sib != bytes32(0);
        bool want = !m.requested && m.status != CANCELED && m.status != RESULTS
            && (m.status == ENDED || vm.getBlockTimestamp() >= _end(m)) && vm.getBlockTimestamp() >= graceEndOf(pid)
            && honest && shapeOk;

        (bool ok,) = address(registry).call(abi.encodeCall(registry.requestResultsDecryption, (pid, acc, sibs)));
        _check(ok, want, "requestResultsDecryption");
        if (ok) {
            m.requested = true;
            m.requestedRoot = m.root;
            m.status = ENDED;
            m.zeroSkipped = zs;
            m.count = n;
            if (n == 0 && _gateOpen(m)) _resultsLanded(m);
        }
    }

    function _resultsLanded(M storage m) internal {
        m.status = RESULTS;
        m.gateWasOpenAtResults = _gateOpen(m);
    }

    // --- committee -------------------------------------------------------------------

    /// Combines some of the request's ciphertexts with random plaintexts. The Council
    /// manager refuses a combine while its gate is closed.
    function combine(uint256 pi, uint256 mask, uint256 seed) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        if (!m.requested || m.count == 0) return;
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        for (uint256 j = 0; j < m.count; ++j) {
            if ((mask >> j) & 1 == 0 || (m.plainSet >> j) & 1 == 1) continue;
            uint64 value = uint64(uint256(keccak256(abi.encode(seed, j))) % 1_000_000);
            bool ok;
            if (m.mode == uint8(DAVINCITypes.KeyMode.COUNCIL)) {
                (ok,) = address(council).call(abi.encodeCall(council.setPlaintext, (p.dkgAid, uint8(j), value)));
                _check(ok, _gateOpen(m), "council combine");
            } else {
                // MockDKG indices are 1-based per application; the request used [first, first + count).
                dkg.setPlaintext(p.dkgEpochId, p.dkgAid, uint16(p.dkgFirstIndex + j), value);
                ok = true;
            }
            if (ok) {
                m.plain[j] = value;
                m.plainSet |= uint16(1 << j);
            }
        }
    }

    function openGate() external {
        if (gateOpened) return;
        council.openDecryption(manualCid);
        gateOpened = true;
    }

    function reveal(uint256 pi, bool wrong) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        uint256 sk = wrong ? ORG_SK + 1 : ORG_SK;
        bool want = m.mode == uint8(DAVINCITypes.KeyMode.DKG_LOCKED) && !wrong && !m.revealed;
        (bool ok,) = address(registry).call(abi.encodeCall(registry.revealProcessKey, (pid, sk)));
        _check(ok, want, "revealProcessKey");
        if (ok) m.revealed = true;
    }

    // --- finalize --------------------------------------------------------------------

    function finalize(uint256 pi, bool stranger) external {
        bytes31 pid = _pid(pi);
        M storage m = model[pid];
        uint16 all = m.count == 0 ? 0 : uint16((uint256(1) << m.count) - 1);
        bool want = m.status != CANCELED && m.status != RESULTS && vm.getBlockTimestamp() >= graceEndOf(pid)
            && m.requested && _gateOpen(m) && (m.plainSet & all) == all;
        if (stranger) vm.prank(STRANGER);
        (bool ok,) = address(registry).call(abi.encodeCall(registry.finalizeResultsFromDKG, (pid)));
        _check(ok, want, "finalizeResultsFromDKG");
        if (ok) _resultsLanded(m);
    }

    /// Sequencer-mode results never apply to a key-mode process.
    function sequencerResults(uint256 pi, uint64[16] memory values) external {
        bytes31 pid = _pid(pi);
        (bool ok,) = address(registry)
            .call(
                abi.encodeCall(
                    registry.setProcessResults, (pid, AttackLib.resultsPublics(model[pid].root, values), hex"")
                )
            );
        _check(ok, false, "setProcessResults");
    }
}

contract AttackKeyModesInvariantTest is AttackBase {
    using stdJson for string;

    bytes12 internal constant CID_OPEN = bytes12(0x00c0c1a7e0000000000000a1);
    bytes12 internal constant CID_MANUAL = bytes12(0x00c0c1a7e0000000000000b2);

    MockDKG internal mock;
    MockCouncilManager internal council;
    KeyModesHandler internal handler;
    string internal dkgv;

    function _dkgManager() internal override returns (address) {
        mock = new MockDKG();
        return address(mock);
    }

    function _councilManager() internal override returns (address) {
        council = new MockCouncilManager();
        return address(council);
    }

    function _te(string memory p) internal view returns (uint256[2] memory pt) {
        pt[0] = vm.parseUint(dkgv.readString(string.concat(p, ".te_x")));
        pt[1] = vm.parseUint(dkgv.readString(string.concat(p, ".te_y")));
    }

    function _rte(string memory p) internal view returns (uint256[2] memory pt) {
        pt[0] = vm.parseUint(dkgv.readString(string.concat(p, ".rte_x")));
        pt[1] = vm.parseUint(dkgv.readString(string.concat(p, ".rte_y")));
    }

    function _keyed(DAVINCITypes.DKGParams memory d, uint8 numFields, uint256 startOff, uint256 duration)
        internal
        returns (bytes31)
    {
        DAVINCITypes.BallotMode memory bm;
        bm.numFields = numFields;
        bm.costExponent = 1;
        bm.maxValue = 100;
        bm.maxValueSum = 1600;
        DAVINCITypes.Census memory c;
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1;
        c.censusRoot = bytes32(uint256(0xC0FFEE));
        c.censusURI = "https://example.com/census";
        vm.prank(ORGANIZER);
        return registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            vm.getBlockTimestamp() + startOff,
            duration,
            300,
            bm,
            c,
            METADATA_URI,
            METADATA_HASH,
            DAVINCITypes.EncryptionKey(0, 0),
            d
        );
    }

    function setUp() public override {
        super.setUp();
        dkgv = vm.readFile("test/vectors/dkg.json");

        uint256[2][] memory keys = new uint256[2][](4);
        for (uint256 i = 0; i < 4; ++i) {
            keys[i] = _rte(string.concat(".pool_keys[", vm.toString(i % 2), "]"));
        }
        bytes12 eid = mock.newEpoch(true, keys);

        uint256[2] memory ck = _te(".pool_keys[1]");
        council.newCeremony(CID_OPEN, ck[0], ck[1]);
        council.newCeremony(CID_MANUAL, ck[0], ck[1]);
        council.setDecryptionPolicy(CID_MANUAL, MockCouncilManager.PhaseMode.Manual, 0, 0);
        address cadapter = registry.councilAdapter();
        council.allowAdapter(CID_OPEN, cadapter);
        council.allowAdapter(CID_MANUAL, cadapter);
        council.authorizeCreator(CID_OPEN, ORGANIZER);
        council.authorizeCreator(CID_MANUAL, ORGANIZER);

        bytes31[] memory ps = new bytes31[](4);
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        ps[0] = _keyed(d, 16, 0, 2 hours);

        d.mode = DAVINCITypes.KeyMode.DKG_LOCKED;
        d.epochId = eid;
        uint256[2] memory org = _rte(".org_pk");
        (d.orgPKx, d.orgPKy) = (org[0], org[1]);
        ps[1] = _keyed(d, 4, 10 minutes, 1 hours);

        DAVINCITypes.DKGParams memory cd;
        cd.mode = DAVINCITypes.KeyMode.COUNCIL;
        cd.epochId = CID_OPEN;
        ps[2] = _keyed(cd, 1, 0, 30 minutes);
        cd.epochId = CID_MANUAL;
        ps[3] = _keyed(cd, 16, 5 minutes, 90 minutes);

        // Curve points for the accumulators: the fixture's settled ciphertexts and keys.
        uint256[2][] memory pts = new uint256[2][](8);
        bytes32[] memory words = dkgv.readBytes32Array(".settled.accumulator");
        pts[0] = [uint256(words[0]), uint256(words[1])];
        pts[1] = [uint256(words[2]), uint256(words[3])];
        pts[2] = [uint256(words[8]), uint256(words[9])];
        pts[3] = [uint256(words[10]), uint256(words[11])];
        pts[4] = _te(".pool_keys[0]");
        pts[5] = _te(".pool_keys[1]");
        pts[6] = _te(".org_pk");
        pts[7] = _te(".locked_key");

        handler = new KeyModesHandler(registry, mock, council, CID_MANUAL, ps, pts);
        bytes4[] memory sel = new bytes4[](14);
        sel[0] = KeyModesHandler.warp.selector;
        sel[1] = KeyModesHandler.warp.selector;
        sel[2] = KeyModesHandler.warp.selector;
        sel[3] = KeyModesHandler.setStatus.selector;
        sel[4] = KeyModesHandler.submit.selector;
        sel[5] = KeyModesHandler.submit.selector;
        sel[6] = KeyModesHandler.request.selector;
        sel[7] = KeyModesHandler.request.selector;
        sel[8] = KeyModesHandler.combine.selector;
        sel[9] = KeyModesHandler.openGate.selector;
        sel[10] = KeyModesHandler.reveal.selector;
        sel[11] = KeyModesHandler.finalize.selector;
        sel[12] = KeyModesHandler.finalize.selector;
        sel[13] = KeyModesHandler.sequencerResults.selector;
        targetSelector(StdInvariant.FuzzSelector({addr: address(handler), selectors: sel}));
        targetContract(address(handler));
    }

    /// The handler is not vacuous: a scripted run through it settles, requests with every
    /// argument variant, combines, opens the gate and finalizes every process.
    function test_HandlerReachesResultsInEveryMode() public {
        handler.warp(660); // past every start
        for (uint256 i = 0; i < 4; ++i) {
            handler.submit(i, 3, 1000 + i);
        }
        handler.warp(16 * 6000); // a multiple of 16: up to a day
        handler.warp(16 * 6000);
        // Variants 2-7 always tamper; 8 and 9 fall back to honest when they would not
        // change anything (no field past numFields, two equal fields), so skip them here.
        for (uint256 v = 2; v <= 7; ++v) {
            for (uint256 i = 0; i < 4; ++i) {
                handler.request(i, v, 7 * v + i);
            }
        }
        assertEq(handler.successes("requestResultsDecryption"), 0, "a tampered request landed");
        handler.request(1, 8, 0); // numFields 4: field 4 changed
        handler.request(2, 8, 0); // numFields 1: field 1 changed
        assertEq(handler.successes("requestResultsDecryption"), 0, "a tampered request landed");
        for (uint256 i = 0; i < 4; ++i) {
            handler.request(i, i % 2, 0);
        }
        assertEq(handler.successes("submitStateTransition"), 4);
        assertEq(handler.successes("requestResultsDecryption"), 4);
        handler.openGate();
        for (uint256 i = 0; i < 4; ++i) {
            handler.combine(i, type(uint256).max, 77);
            handler.finalize(i, true);
        }
        assertEq(handler.violations(), 0, handler.firstViolation());
        for (uint256 i = 0; i < 4; ++i) {
            assertEq(uint8(registry.getProcess(handler.pids(i)).status), 4, "no results");
        }
        invariant_StateAndResultsMatchTheModel();
    }

    /// forge-config: default.invariant.runs = 96
    /// forge-config: default.invariant.depth = 100
    /// forge-config: default.invariant.fail-on-revert = true
    function invariant_EveryCallMatchesTheModel() public view {
        assertEq(handler.violations(), 0, handler.firstViolation());
    }

    /// forge-config: default.invariant.runs = 96
    /// forge-config: default.invariant.depth = 100
    /// forge-config: default.invariant.fail-on-revert = true
    function invariant_StateAndResultsMatchTheModel() public view {
        for (uint256 i = 0; i < handler.pidCount(); ++i) {
            bytes31 pid = handler.pids(i);
            (uint8 status, bytes32 root, bool requested, uint8 count, uint16 zs, uint256 voters) = handler.modelOf(pid);
            DAVINCITypes.Process memory p = registry.getProcess(pid);
            assertEq(uint8(p.status), status, "status");
            assertEq(p.latestStateRoot, root, "root");
            assertEq(p.votersCount, voters, "voters");
            assertEq(p.dkgResultsRequested, requested, "requested");
            assertEq(registry.getProcessGraceEnd(pid), handler.graceEndOf(pid), "grace end");
            if (!requested) continue;
            assertEq(p.dkgCount, count, "dkgCount");
            assertEq(p.dkgZeroSkipped, zs, "dkgZeroSkipped");
            // The tally is over the root the request proved, and nothing settled after it.
            assertEq(p.latestStateRoot, handler.requestedRoot(pid), "root moved after the request");
            if (status != 4) continue;
            uint256[] memory want = handler.expectedResult(pid);
            assertEq(p.result.length, want.length, "result length");
            for (uint256 f = 0; f < want.length; ++f) {
                assertEq(p.result[f], want[f], "result field");
            }
            assertTrue(handler.gateWasOpenAtResults(pid), "results published with the gate closed");
        }
    }
}
