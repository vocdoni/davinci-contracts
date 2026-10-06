// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {Vm} from "forge-std/Vm.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {DKGTest} from "./DKG.t.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {ICouncilManager} from "../src/interfaces/council/ICouncilManager.sol";
import {ICouncilManagerErrors} from "../src/interfaces/council/ICouncilManagerErrors.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {Sha256SmtLib} from "../src/libraries/Sha256SmtLib.sol";
import {CouncilAdapter} from "../src/CouncilAdapter.sol";
import {MockCouncilManager} from "./mocks/MockCouncilManager.sol";
import {MockDKG} from "./mocks/MockDKG.sol";

/// @dev A registry with a Council manager: a Live ceremony that allows the adapter and
///      authorizes ORGANIZER, keyed with a fixture TE point.
abstract contract CouncilTestBase is RegistryTestBase {
    using stdJson for string;

    bytes12 internal constant CID = bytes12(0x00c0c1a7e0000000000000a1);

    MockCouncilManager internal council;
    CouncilAdapter internal cadapter;
    string internal dkg;
    uint256 internal pkX;
    uint256 internal pkY;

    function _councilManager() internal override returns (address) {
        council = new MockCouncilManager();
        return address(council);
    }

    function setUp() public virtual override {
        super.setUp();
        dkg = vm.readFile("test/vectors/dkg.json");
        cadapter = CouncilAdapter(registry.councilAdapter());
        pkX = vm.parseUint(dkg.readString(".pool_keys[1].te_x"));
        pkY = vm.parseUint(dkg.readString(".pool_keys[1].te_y"));
        council.newCeremony(CID, pkX, pkY);
        council.allowAdapter(CID, address(cadapter));
        council.authorizeCreator(CID, ORGANIZER);
    }

    function _dkgBallotMode() internal view returns (DAVINCITypes.BallotMode memory m) {
        m.numFields = uint8(dkg.readUint(".ballot_mode.num_fields"));
        m.costExponent = uint8(dkg.readUint(".ballot_mode.cost_exponent"));
        m.maxValue = dkg.readUint(".ballot_mode.max_value");
        m.maxValueSum = dkg.readUint(".ballot_mode.max_value_sum");
    }

    function _dkgCensus() internal view returns (DAVINCITypes.Census memory c) {
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1;
        c.censusRoot = dkg.readBytes32(".census_root");
        c.censusURI = "https://example.com/census";
    }

    function _councilParams(bytes12 cid) internal pure returns (DAVINCITypes.DKGParams memory d) {
        d.mode = DAVINCITypes.KeyMode.COUNCIL;
        d.epochId = cid;
    }

    function _newKeyedProcess(address creator, DAVINCITypes.DKGParams memory d) internal returns (bytes31) {
        vm.prank(creator);
        return registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            0,
            DURATION,
            MAX_VOTERS,
            _dkgBallotMode(),
            _dkgCensus(),
            METADATA_URI,
            METADATA_HASH,
            DAVINCITypes.EncryptionKey(0, 0),
            d
        );
    }

    function _councilProcess() internal returns (bytes31) {
        return _newKeyedProcess(ORGANIZER, _councilParams(CID));
    }

    /// @dev The fixture's settled accumulator: fields 0 and 2 active, 1 and 3 identity.
    function _settledAccumulator() internal view returns (uint256[64] memory acc) {
        bytes32[] memory words = dkg.readBytes32Array(".settled.accumulator");
        for (uint256 i = 0; i < 64; ++i) {
            acc[i] = uint256(words[i]);
        }
    }

    /// @dev Every field the identity ciphertext (0, 1, 0, 1).
    function _identityAccumulator() internal pure returns (uint256[64] memory acc) {
        for (uint256 i = 0; i < 16; ++i) {
            acc[4 * i + 1] = 1;
            acc[4 * i + 3] = 1;
        }
    }

    /// @dev Field `i` of acc as [c1x, c1y, c2x, c2y].
    function _field(uint256[64] memory acc, uint256 i) internal pure returns (uint256[4] memory ct) {
        for (uint256 w = 0; w < 4; ++w) {
            ct[w] = acc[4 * i + w];
        }
    }

    /// @dev Settles `acc` as the process's only leaf (0x04): root = its leaf hash, proven by
    ///      one zero sibling. Overwrites p.latestStateRoot (struct slot 3, processes at slot 1).
    function _settle(bytes31 pid, uint256[64] memory acc) internal returns (bytes32[] memory siblings) {
        bytes32 root = Sha256SmtLib.leafHash(4, uint256(sha256(abi.encode(acc))));
        bytes32 base = keccak256(abi.encode(pid, uint256(1)));
        vm.store(address(registry), bytes32(uint256(base) + 3), root);
        assertEq(registry.getProcess(pid).latestStateRoot, root, "latestStateRoot slot moved");
        siblings = new bytes32[](1);
    }

    function _assertResults(bytes31 pid, uint256[4] memory want) internal view {
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(uint8(p.status), uint8(DAVINCITypes.ProcessStatus.RESULTS));
        assertEq(p.result.length, 4);
        for (uint256 i = 0; i < 4; ++i) {
            assertEq(p.result[i], want[i]);
        }
    }
}

/// @dev Both key adapters enabled: the DKG modes against MockDKG, COUNCIL against
///      MockCouncilManager.
contract CouncilTest is CouncilTestBase {
    using stdJson for string;

    address internal constant CREATOR2 = address(0xC0FFEE);
    address internal constant STRANGER = address(0xBAD);

    MockDKG internal mock;
    bytes12 internal eid1;

    function _dkgManager() internal override returns (address) {
        mock = new MockDKG();
        return address(mock);
    }

    function setUp() public override {
        super.setUp();
        uint256[2][] memory keys = new uint256[2][](1);
        keys[0][0] = vm.parseUint(dkg.readString(".pool_keys[0].rte_x"));
        keys[0][1] = vm.parseUint(dkg.readString(".pool_keys[0].rte_y"));
        eid1 = mock.newEpoch(true, keys);
    }

    // --- wiring and ABI --------------------------------------------------------------

    function test_Wiring() public view {
        assertEq(uint8(DAVINCITypes.KeyMode.COUNCIL), 3);
        assertEq(IProcessRegistry.newProcess.selector, bytes4(0x08c0fdd3), "newProcess selector moved");
        // The DKG adapter keeps CREATE(registry, 1); the Council adapter comes second.
        assertEq(registry.dkgAdapter(), vm.computeCreateAddress(address(registry), 1));
        assertEq(registry.councilAdapter(), vm.computeCreateAddress(address(registry), 2));
        assertEq(cadapter.registry(), address(registry));
        assertEq(address(cadapter.manager()), address(council));
        // Adapter reverts that surface through newProcess decode against the registry ABI.
        assertEq(CouncilAdapter.InvalidKeyMode.selector, IProcessRegistry.InvalidKeyMode.selector);
        assertEq(CouncilAdapter.InvalidDKGParams.selector, IProcessRegistry.InvalidDKGParams.selector);
        assertEq(CouncilAdapter.UnknownRequest.selector, ICouncilManagerErrors.UnknownRequest.selector);
        // The registry's gate revert is the manager's.
        assertEq(IProcessRegistry.DecryptionNotOpen.selector, ICouncilManagerErrors.DecryptionNotOpen.selector);
    }

    function test_Adapter_OnlyRegistry() public {
        vm.expectRevert(CouncilAdapter.NotRegistry.selector);
        cadapter.register(bytes31(uint248(1)), ORGANIZER, _councilParams(CID));
        vm.expectRevert(CouncilAdapter.NotRegistry.selector);
        cadapter.submit(CID, bytes32(uint256(1)), new uint256[4][](1));

        DAVINCITypes.DKGParams memory d = _councilParams(CID);
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        vm.prank(address(registry));
        vm.expectRevert(CouncilAdapter.InvalidKeyMode.selector);
        cadapter.register(bytes31(uint248(1)), ORGANIZER, d);

        vm.expectRevert(CouncilAdapter.UnsupportedKeyMode.selector);
        cadapter.reveal(CID, bytes32(uint256(1)), 1);
    }

    // --- newProcess --------------------------------------------------------------------

    function test_NewProcess_Council() public {
        bytes31 pid = registry.getNextProcessId(ORGANIZER);
        bytes32 rid = council.requestIdFor(CID, address(cadapter), pid);
        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessCreated(pid, ORGANIZER);
        assertEq(_councilProcess(), pid);

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(uint8(p.keyMode), uint8(DAVINCITypes.KeyMode.COUNCIL));
        assertEq(p.dkgEpochId, CID);
        assertEq(p.dkgAid, rid, "dkgAid is the request id");
        // The ceremony key, TE as Council returns it: no chart conversion.
        assertEq(p.encryptionKey.x, pkX);
        assertEq(p.encryptionKey.y, pkY);
        assertEq(
            p.latestStateRoot,
            registry.genesisRoot(
                pid, _dkgBallotMode(), p.encryptionKey, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
            )
        );

        (bytes12 bcid, uint8 fieldCount, bytes31 bpid) = cadapter.bindings(rid);
        assertEq(bcid, CID);
        assertEq(fieldCount, 0);
        assertEq(bytes32(bpid), bytes32(pid), "adapter maps the request back to the process");
        (bytes12 mcid, bytes32 mrid, bool requested) = council.getBinding(address(cadapter), pid);
        assertEq(mcid, CID);
        assertEq(mrid, rid);
        assertFalse(requested);
        assertEq(council.bindingCreator(address(cadapter), pid), ORGANIZER);
    }

    /// @dev The registry hands the adapter its own msg.sender: neither itself, tx.origin nor
    ///      the adapter counts as the creator.
    function test_NewProcess_CreatorIsTheCaller() public {
        council.authorizeCreator(CID, CREATOR2);
        bytes31 pid = _newKeyedProcess(CREATOR2, _councilParams(CID));
        assertEq(council.bindingCreator(address(cadapter), pid), CREATOR2);
        assertEq(registry.getProcess(pid).organizationId, CREATOR2);

        council.authorizeCreator(CID, address(registry));
        council.authorizeCreator(CID, address(cadapter));
        DAVINCITypes.DKGParams memory d = _councilParams(CID);
        vm.expectRevert(ICouncilManagerErrors.NotAuthorizedCreator.selector);
        vm.prank(STRANGER, ORGANIZER); // tx.origin is authorized, msg.sender is not
        registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            0,
            DURATION,
            MAX_VOTERS,
            _dkgBallotMode(),
            _dkgCensus(),
            METADATA_URI,
            METADATA_HASH,
            DAVINCITypes.EncryptionKey(0, 0),
            d
        );
    }

    function test_NewProcess_UnauthorizedCreator() public {
        DAVINCITypes.DKGParams memory d = _councilParams(CID);
        vm.expectRevert(ICouncilManagerErrors.NotAuthorizedCreator.selector);
        this.createAs(STRANGER, d);
        assertEq(registry.processNonce(STRANGER), 0, "nothing created");
    }

    function test_NewProcess_AdapterNotAllowed() public {
        bytes12 cid2 = bytes12(uint96(2));
        council.newCeremony(cid2, pkX, pkY);
        council.authorizeCreator(cid2, ORGANIZER);
        vm.expectRevert(ICouncilManagerErrors.NotAllowedAdapter.selector);
        this.createAs(ORGANIZER, _councilParams(cid2));
    }

    function test_NewProcess_CeremonyNotLive() public {
        council.setPhase(CID, MockCouncilManager.Phase.Dealing);
        vm.expectRevert(ICouncilManagerErrors.WrongPhase.selector);
        this.createAs(ORGANIZER, _councilParams(CID));

        vm.expectRevert(ICouncilManagerErrors.UnknownCeremony.selector);
        this.createAs(ORGANIZER, _councilParams(bytes12(uint96(3))));
    }

    function test_NewProcess_RejectsOrganizerFields() public {
        for (uint256 i = 0; i < 6; ++i) {
            DAVINCITypes.DKGParams memory d = _councilParams(CID);
            if (i == 0) d.orgPKx = 1;
            else if (i == 1) d.orgPKy = 1;
            else if (i == 2) d.popAx = 1;
            else if (i == 3) d.popAy = 1;
            else if (i == 4) d.popZ = 1;
            else d.epochId = bytes12(0);
            vm.expectRevert(IProcessRegistry.InvalidDKGParams.selector);
            this.createAs(ORGANIZER, d);
        }
    }

    function test_NewProcess_RejectsEncryptionKey() public {
        vm.expectRevert(IProcessRegistry.InvalidEncryptionKey.selector);
        vm.prank(ORGANIZER);
        registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            0,
            DURATION,
            MAX_VOTERS,
            _dkgBallotMode(),
            _dkgCensus(),
            METADATA_URI,
            METADATA_HASH,
            DAVINCITypes.EncryptionKey(pkX, pkY),
            _councilParams(CID)
        );
    }

    /// @dev External, so the revert of a nested newProcess is what vm.expectRevert sees.
    function createAs(address creator, DAVINCITypes.DKGParams memory d) external returns (bytes31) {
        return _newKeyedProcess(creator, d);
    }

    // --- results ---------------------------------------------------------------------

    function test_Request_RoutesToCouncil() public {
        bytes31 pid = _councilProcess();
        bytes32 rid = registry.getProcess(pid).dkgAid;
        uint256[64] memory acc = _settledAccumulator();
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);

        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessStatusChanged(
            pid, DAVINCITypes.ProcessStatus.READY, DAVINCITypes.ProcessStatus.ENDED
        );
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ResultsDecryptionRequested(pid, CID, rid, 0, 2);
        registry.requestResultsDecryption(pid, acc, siblings);

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(uint8(p.status), uint8(DAVINCITypes.ProcessStatus.ENDED));
        assertTrue(p.dkgResultsRequested);
        assertEq(p.dkgFirstIndex, 0);
        assertEq(p.dkgCount, 2);
        assertEq(p.dkgZeroSkipped, (1 << 1) | (1 << 3));

        // One request, the active fields in order, TE words untouched.
        (bytes12 rcid, uint8 fieldCount,,) = council.getRequestMeta(rid);
        uint256[4][] memory cts = council.requestCts(rid);
        assertEq(rcid, CID);
        assertEq(fieldCount, 2);
        assertEq(cts.length, 2);
        for (uint256 w = 0; w < 4; ++w) {
            assertEq(cts[0][w], acc[w]);
            assertEq(cts[1][w], acc[8 + w]);
        }
        (,, bool requested) = council.getBinding(address(cadapter), pid);
        assertTrue(requested);
        (, uint8 adapterCount,) = cadapter.bindings(rid);
        assertEq(adapterCount, 2);

        vm.expectRevert(IProcessRegistry.ResultsAlreadyRequested.selector);
        registry.requestResultsDecryption(pid, acc, siblings);

        // Not ready until every field is combined.
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pid);
        council.setPlaintext(rid, 0, 77);
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pid);
        council.setPlaintext(rid, 1, 99);

        uint256[] memory expected = new uint256[](4);
        expected[0] = 77;
        expected[2] = 99;
        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessResultsSet(pid, address(this), expected);
        registry.finalizeResultsFromDKG(pid);
        _assertResults(pid, [uint256(77), 0, 99, 0]);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.finalizeResultsFromDKG(pid);
    }

    /// @dev Identity fields first and in between: request field k maps to the k-th active
    ///      field, the skipped ones read 0.
    function test_Request_PlaintextsSkipZeroFields() public {
        bytes31 pid = _councilProcess();
        bytes32 rid = registry.getProcess(pid).dkgAid;
        uint256[64] memory settled = _settledAccumulator();
        uint256[64] memory acc = _identityAccumulator();
        for (uint256 w = 0; w < 4; ++w) {
            acc[4 + w] = settled[w]; // field 1
            acc[12 + w] = settled[8 + w]; // field 3
        }
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);
        registry.requestResultsDecryption(pid, acc, siblings);

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.dkgZeroSkipped, (1 << 0) | (1 << 2));
        uint256[4][] memory cts = council.requestCts(rid);
        assertEq(cts.length, 2);
        assertEq(keccak256(abi.encode(cts[0])), keccak256(abi.encode(_field(acc, 1))));
        assertEq(keccak256(abi.encode(cts[1])), keccak256(abi.encode(_field(acc, 3))));

        council.setPlaintext(rid, 0, 5);
        council.setPlaintext(rid, 1, 6);
        registry.finalizeResultsFromDKG(pid);
        _assertResults(pid, [uint256(0), 5, 0, 6]);
    }

    /// @dev The manager's admission checks bubble up and revert the whole request: the process
    ///      stays as it was and can be requested again.
    function test_Request_ManagerRejectsAnInvalidCiphertext() public {
        bytes31 pid = _councilProcess();
        uint256[64] memory acc = _settledAccumulator();
        acc[0] = addmod(acc[0], 1, BN254_P); // C1 of field 0 off the curve, still canonical
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);
        vm.expectRevert(ICouncilManagerErrors.InvalidPoint.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertFalse(p.dkgResultsRequested);
        assertEq(uint8(p.status), uint8(DAVINCITypes.ProcessStatus.READY));
    }

    function test_Request_AllIdentityFinalizesWithoutRequest() public {
        bytes31 pid = _councilProcess();
        bytes32 rid = registry.getProcess(pid).dkgAid;
        uint256[64] memory acc = _identityAccumulator();
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);

        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ResultsDecryptionRequested(pid, CID, rid, 0, 0);
        registry.requestResultsDecryption(pid, acc, siblings);
        _assertResults(pid, [uint256(0), 0, 0, 0]);

        (,, bool requested) = council.getBinding(address(cadapter), pid);
        assertFalse(requested, "nothing sent to the manager");
        (, uint8 fieldCount,,) = council.getRequestMeta(rid);
        assertEq(fieldCount, 0, "bound, never submitted");
    }

    function test_Plaintexts_WholeRequestOnly() public {
        bytes31 pid = _councilProcess();
        bytes32 rid = registry.getProcess(pid).dkgAid;

        // Bound, not yet requested.
        vm.expectRevert(CouncilAdapter.UnknownRequest.selector);
        cadapter.plaintexts(CID, rid, 0, 2);

        uint256[64] memory acc = _settledAccumulator();
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);
        registry.requestResultsDecryption(pid, acc, siblings);

        vm.expectRevert(CouncilAdapter.InvalidFieldRange.selector);
        cadapter.plaintexts(CID, rid, 1, 2);
        vm.expectRevert(CouncilAdapter.InvalidFieldRange.selector);
        cadapter.plaintexts(CID, rid, 0, 1);
        vm.expectRevert(CouncilAdapter.InvalidFieldRange.selector);
        cadapter.plaintexts(CID, rid, 0, 3);
        vm.expectRevert(CouncilAdapter.UnknownRequest.selector);
        cadapter.plaintexts(bytes12(uint96(9)), rid, 0, 2);
        vm.expectRevert(CouncilAdapter.UnknownRequest.selector);
        cadapter.plaintexts(CID, keccak256("other"), 0, 2);

        (bool ready, uint256[] memory values) = cadapter.plaintexts(CID, rid, 0, 2);
        assertFalse(ready);
        assertEq(values.length, 2);
        council.setPlaintext(rid, 0, 1);
        council.setPlaintext(rid, 1, 2);
        (ready, values) = cadapter.plaintexts(CID, rid, 0, 2);
        assertTrue(ready);
        assertEq(values[0], 1);
        assertEq(values[1], 2);
    }

    function test_TwoProcessesOnOneCeremony() public {
        bytes31 pidA = _councilProcess();
        bytes31 pidB = _councilProcess();
        DAVINCITypes.Process memory a = registry.getProcess(pidA);
        DAVINCITypes.Process memory b = registry.getProcess(pidB);

        // One ceremony key, two requests.
        assertEq(a.encryptionKey.x, b.encryptionKey.x);
        assertEq(a.encryptionKey.y, b.encryptionKey.y);
        assertEq(a.dkgEpochId, CID);
        assertEq(b.dkgEpochId, CID);
        assertTrue(a.dkgAid != b.dkgAid);
        bytes32[] memory ids = council.getRequestIds(CID);
        assertEq(ids.length, 2);
        assertEq(ids[0], a.dkgAid);
        assertEq(ids[1], b.dkgAid);
        (,, bytes31 mappedA) = cadapter.bindings(a.dkgAid);
        (,, bytes31 mappedB) = cadapter.bindings(b.dkgAid);
        assertEq(bytes32(mappedA), bytes32(pidA));
        assertEq(bytes32(mappedB), bytes32(pidB));

        uint256[64] memory accA = _settledAccumulator();
        uint256[64] memory accB = _identityAccumulator();
        for (uint256 w = 0; w < 4; ++w) {
            accB[4 + w] = accA[w];
            accB[12 + w] = accA[8 + w];
        }
        bytes32[] memory sibA = _settle(pidA, accA);
        bytes32[] memory sibB = _settle(pidB, accB);
        _warpToGraceEnd(pidB);
        registry.requestResultsDecryption(pidA, accA, sibA);
        registry.requestResultsDecryption(pidB, accB, sibB);

        council.setPlaintext(a.dkgAid, 0, 1);
        council.setPlaintext(a.dkgAid, 1, 2);
        registry.finalizeResultsFromDKG(pidA);
        _assertResults(pidA, [uint256(1), 0, 2, 0]);

        // B's request is untouched by A's combines.
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pidB);
        council.setPlaintext(b.dkgAid, 0, 3);
        council.setPlaintext(b.dkgAid, 1, 4);
        registry.finalizeResultsFromDKG(pidB);
        _assertResults(pidB, [uint256(0), 3, 0, 4]);
    }

    function test_KeyModeGates() public {
        bytes31 pid = _councilProcess();
        vm.expectRevert(IProcessRegistry.InvalidKeyMode.selector);
        registry.revealProcessKey(pid, 12345);

        _warpToGraceEnd(pid);
        uint64[16] memory values;
        bytes memory pv = _resultsPublics(registry.getProcess(pid).latestStateRoot, values);
        vm.expectRevert(IProcessRegistry.InvalidKeyMode.selector);
        registry.setProcessResults(pid, pv, new bytes(768));
    }

    /// @dev A DKG_AUTOMATIC and a COUNCIL process side by side: each request and each
    ///      finalize reaches only its own adapter.
    function test_ModesRouteToTheirOwnAdapter() public {
        DAVINCITypes.DKGParams memory auto_;
        auto_.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        bytes31 dkgPid = _newKeyedProcess(ORGANIZER, auto_);
        bytes31 cPid = _councilProcess();
        DAVINCITypes.Process memory dp = registry.getProcess(dkgPid);
        assertEq(uint8(dp.keyMode), uint8(DAVINCITypes.KeyMode.DKG_AUTOMATIC));
        assertEq(dp.dkgEpochId, eid1);
        assertEq(dp.dkgAid, registry.aidFor(dkgPid));
        vm.expectRevert(ICouncilManagerErrors.UnknownBinding.selector);
        council.getBinding(address(cadapter), dkgPid); // the DKG process is not bound on Council

        uint256[64] memory acc = _settledAccumulator();
        bytes32[] memory sibD = _settle(dkgPid, acc);
        bytes32[] memory sibC = _settle(cPid, acc);
        _warpToGraceEnd(cPid);
        registry.requestResultsDecryption(dkgPid, acc, sibD);
        registry.requestResultsDecryption(cPid, acc, sibC);

        dp = registry.getProcess(dkgPid);
        assertEq(dp.dkgFirstIndex, 1, "MockDKG indices are 1-based");
        assertEq(mock.ctCount(eid1, dp.dkgAid), 2);
        bytes32 rid = registry.getProcess(cPid).dkgAid;
        (, uint8 fieldCount,,) = council.getRequestMeta(rid);
        assertEq(fieldCount, 2);

        mock.setPlaintext(eid1, dp.dkgAid, 1, 10);
        mock.setPlaintext(eid1, dp.dkgAid, 2, 20);
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(cPid);
        registry.finalizeResultsFromDKG(dkgPid);
        _assertResults(dkgPid, [uint256(10), 0, 20, 0]);

        council.setPlaintext(rid, 0, 30);
        council.setPlaintext(rid, 1, 40);
        registry.finalizeResultsFromDKG(cPid);
        _assertResults(cPid, [uint256(30), 0, 40, 0]);
    }

    // --- decryption gate (Council protocol §8.7) ---------------------------------------

    uint64 internal constant SIX_MONTHS = 182 days;

    /// @dev A COUNCIL process with an all-identity accumulator, requested at its grace end.
    function _requestZeroResults() internal returns (bytes31 pid) {
        pid = _councilProcess();
        uint256[64] memory acc = _identityAccumulator();
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);
        registry.requestResultsDecryption(pid, acc, siblings);
    }

    /// @dev Requested, ENDED, no results.
    function _assertWaiting(bytes31 pid) internal view {
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(uint8(p.status), uint8(DAVINCITypes.ProcessStatus.ENDED));
        assertTrue(p.dkgResultsRequested);
        assertEq(p.result.length, 0);
    }

    function test_Gate_AdapterReadsTheManager() public {
        assertTrue(cadapter.isDecryptionOpen(CID));
        uint64 openAt = uint64(block.timestamp + 1 days);
        council.setDecryptionPolicy(CID, MockCouncilManager.PhaseMode.Scheduled, openAt, 0);
        assertFalse(cadapter.isDecryptionOpen(CID));
        vm.warp(openAt - 1);
        assertFalse(cadapter.isDecryptionOpen(CID));
        vm.warp(openAt);
        assertTrue(cadapter.isDecryptionOpen(CID));
        vm.expectRevert(ICouncilManagerErrors.UnknownCeremony.selector);
        cadapter.isDecryptionOpen(bytes12(uint96(9)));
    }

    /// @dev An all-zero tally ends six months before a scheduled opening: the request records
    ///      ENDED and nothing else, nobody can publish or cancel meanwhile, and anyone publishes
    ///      the zero vector from the opening on.
    function test_Gate_ZeroResultsWaitForScheduledOpening() public {
        uint64 openAt = uint64(block.timestamp + SIX_MONTHS);
        council.setDecryptionPolicy(CID, MockCouncilManager.PhaseMode.Scheduled, openAt, 0);
        bytes31 pid = _councilProcess();
        bytes32 rid = registry.getProcess(pid).dkgAid;
        uint256[64] memory acc = _identityAccumulator();
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);

        vm.recordLogs();
        registry.requestResultsDecryption(pid, acc, siblings);
        Vm.Log[] memory logs = vm.getRecordedLogs();
        assertEq(logs.length, 2, "no ProcessResultsSet");
        assertEq(logs[0].topics[0], IProcessRegistry.ProcessStatusChanged.selector);
        assertEq(
            keccak256(logs[0].data),
            keccak256(abi.encode(DAVINCITypes.ProcessStatus.READY, DAVINCITypes.ProcessStatus.ENDED))
        );
        assertEq(logs[1].topics[0], IProcessRegistry.ResultsDecryptionRequested.selector);
        assertEq(keccak256(logs[1].data), keccak256(abi.encode(CID, rid, uint16(0), uint8(0))));
        _assertWaiting(pid);
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.dkgCount, 0);
        assertEq(p.dkgZeroSkipped, 0xf);
        (,, bool requested) = council.getBinding(address(cadapter), pid);
        assertFalse(requested, "nothing sent to the manager");

        vm.expectRevert(IProcessRegistry.DecryptionNotOpen.selector);
        registry.finalizeResultsFromDKG(pid);
        vm.expectRevert(IProcessRegistry.ResultsAlreadyRequested.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);
        vm.warp(openAt - 1);
        vm.expectRevert(IProcessRegistry.DecryptionNotOpen.selector);
        registry.finalizeResultsFromDKG(pid);
        _assertWaiting(pid);

        vm.warp(openAt);
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessStatusChanged(
            pid, DAVINCITypes.ProcessStatus.ENDED, DAVINCITypes.ProcessStatus.RESULTS
        );
        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessResultsSet(pid, STRANGER, new uint256[](4));
        vm.prank(STRANGER);
        registry.finalizeResultsFromDKG(pid);
        _assertResults(pid, [uint256(0), 0, 0, 0]);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.finalizeResultsFromDKG(pid);
    }

    /// @dev Manual opening without a fallback: the zero results wait for the organizer, for
    ///      years if need be.
    function test_Gate_ZeroResultsWaitForManualOpening() public {
        council.setDecryptionPolicy(CID, MockCouncilManager.PhaseMode.Manual, 0, 0);
        bytes31 pid = _requestZeroResults();
        _assertWaiting(pid);
        vm.warp(block.timestamp + 10 * 365 days);
        vm.expectRevert(IProcessRegistry.DecryptionNotOpen.selector);
        registry.finalizeResultsFromDKG(pid);

        council.openDecryption(CID);
        registry.finalizeResultsFromDKG(pid);
        _assertResults(pid, [uint256(0), 0, 0, 0]);
    }

    /// @dev Manual opening with a fallback date and an absent organizer: the date alone opens
    ///      the gate, with no transaction on the manager.
    function test_Gate_ZeroResultsOpenAtTheManualFallback() public {
        uint64 fallbackAt = uint64(block.timestamp + SIX_MONTHS);
        council.setDecryptionPolicy(CID, MockCouncilManager.PhaseMode.Manual, 0, fallbackAt);
        bytes31 pid = _requestZeroResults();
        vm.warp(fallbackAt - 1);
        vm.expectRevert(IProcessRegistry.DecryptionNotOpen.selector);
        registry.finalizeResultsFromDKG(pid);
        _assertWaiting(pid);
        vm.warp(fallbackAt);
        registry.finalizeResultsFromDKG(pid);
        _assertResults(pid, [uint256(0), 0, 0, 0]);
    }

    /// @dev A request made after the opening takes the fast path, as before.
    function test_Gate_ZeroResultsFinalizeAtRequestOnceOpen() public {
        uint64 openAt = uint64(block.timestamp + 1);
        council.setDecryptionPolicy(CID, MockCouncilManager.PhaseMode.Scheduled, openAt, 0);
        assertFalse(cadapter.isDecryptionOpen(CID));
        bytes31 pid = _councilProcess();
        uint256[64] memory acc = _identityAccumulator();
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);
        assertTrue(cadapter.isDecryptionOpen(CID));
        registry.requestResultsDecryption(pid, acc, siblings);
        _assertResults(pid, [uint256(0), 0, 0, 0]);
    }

    /// @dev Nonzero results: the request is admitted while the gate is closed, no combine
    ///      lands and the registry refuses to finalize until the opening.
    function test_Gate_NonzeroResultsWaitForOpening() public {
        uint64 openAt = uint64(block.timestamp + SIX_MONTHS);
        council.setDecryptionPolicy(CID, MockCouncilManager.PhaseMode.Scheduled, openAt, 0);
        bytes31 pid = _councilProcess();
        bytes32 rid = registry.getProcess(pid).dkgAid;
        uint256[64] memory acc = _settledAccumulator();
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);
        registry.requestResultsDecryption(pid, acc, siblings);
        (, uint8 fieldCount,,) = council.getRequestMeta(rid);
        assertEq(fieldCount, 2, "admitted while closed");
        _assertWaiting(pid);

        vm.expectRevert(ICouncilManagerErrors.DecryptionNotOpen.selector);
        council.setPlaintext(rid, 0, 77);
        vm.expectRevert(IProcessRegistry.DecryptionNotOpen.selector);
        registry.finalizeResultsFromDKG(pid);

        vm.warp(openAt);
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pid);
        council.setPlaintext(rid, 0, 77);
        council.setPlaintext(rid, 1, 99);
        registry.finalizeResultsFromDKG(pid);
        _assertResults(pid, [uint256(77), 0, 99, 0]);
    }

    /// @dev The registry checks the gate itself on the nonzero path too: complete plaintexts
    ///      under a gate the manager reports closed (which the real manager cannot produce)
    ///      are not published.
    function test_Gate_RegistryChecksTheNonzeroPathItself() public {
        bytes31 pid = _councilProcess();
        bytes32 rid = registry.getProcess(pid).dkgAid;
        uint256[64] memory acc = _settledAccumulator();
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);
        registry.requestResultsDecryption(pid, acc, siblings);
        council.setPlaintext(rid, 0, 1);
        council.setPlaintext(rid, 1, 2);
        (bool ready,) = cadapter.plaintexts(CID, rid, 0, 2);
        assertTrue(ready);

        vm.mockCall(address(council), abi.encodeCall(ICouncilManager.isDecryptionOpen, (CID)), abi.encode(false));
        vm.expectRevert(IProcessRegistry.DecryptionNotOpen.selector);
        registry.finalizeResultsFromDKG(pid);
        vm.clearMockedCalls();
        registry.finalizeResultsFromDKG(pid);
        _assertResults(pid, [uint256(1), 0, 2, 0]);
    }

    /// @dev The gate is COUNCIL-only: with the ceremony closed, a DKG-mode all-zero tally
    ///      still finalizes at the request, and the Council side is never asked.
    function test_Gate_DKGZeroResultsIgnoreIt() public {
        council.setDecryptionPolicy(CID, MockCouncilManager.PhaseMode.Manual, 0, 0);
        DAVINCITypes.DKGParams memory auto_;
        auto_.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        bytes31 pid = _newKeyedProcess(ORGANIZER, auto_);
        uint256[64] memory acc = _identityAccumulator();
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);

        vm.expectCall(address(cadapter), abi.encodeWithSelector(CouncilAdapter.isDecryptionOpen.selector), 0);
        vm.expectCall(address(council), abi.encodeWithSelector(ICouncilManager.isDecryptionOpen.selector), 0);
        registry.requestResultsDecryption(pid, acc, siblings);
        _assertResults(pid, [uint256(0), 0, 0, 0]);
    }

    /// @dev Same for a nonzero DKG-mode tally: finalized once combined, gate or not.
    function test_Gate_DKGNonzeroResultsIgnoreIt() public {
        council.setDecryptionPolicy(CID, MockCouncilManager.PhaseMode.Manual, 0, 0);
        DAVINCITypes.DKGParams memory auto_;
        auto_.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        bytes31 pid = _newKeyedProcess(ORGANIZER, auto_);
        uint256[64] memory acc = _settledAccumulator();
        bytes32[] memory siblings = _settle(pid, acc);
        _warpToGraceEnd(pid);

        vm.expectCall(address(cadapter), abi.encodeWithSelector(CouncilAdapter.isDecryptionOpen.selector), 0);
        vm.expectCall(address(council), abi.encodeWithSelector(ICouncilManager.isDecryptionOpen.selector), 0);
        registry.requestResultsDecryption(pid, acc, siblings);
        DAVINCITypes.Process memory dp = registry.getProcess(pid);
        mock.setPlaintext(eid1, dp.dkgAid, dp.dkgFirstIndex, 10);
        mock.setPlaintext(eid1, dp.dkgAid, dp.dkgFirstIndex + 1, 20);
        registry.finalizeResultsFromDKG(pid);
        _assertResults(pid, [uint256(10), 0, 20, 0]);
    }
}

/// @dev Council enabled, DKG disabled: the adapters are independent.
contract CouncilOnlyTest is CouncilTestBase {
    function test_DKGModesStayDisabled() public {
        assertEq(registry.dkgAdapter(), address(0));
        assertEq(registry.councilAdapter(), vm.computeCreateAddress(address(registry), 1));

        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        vm.expectRevert(IProcessRegistry.DKGDisabled.selector);
        this.createAs(ORGANIZER, d);
        vm.expectRevert(IProcessRegistry.DKGDisabled.selector);
        registry.aidFor(bytes31(uint248(1)));

        bytes31 pid = _councilProcess();
        assertEq(uint8(registry.getProcess(pid).keyMode), uint8(DAVINCITypes.KeyMode.COUNCIL));
    }

    function createAs(address creator, DAVINCITypes.DKGParams memory d) external returns (bytes31) {
        return _newKeyedProcess(creator, d);
    }
}

/// @dev DKG enabled, Council disabled: COUNCIL never falls through to the DKG adapter.
contract CouncilDisabledTest is RegistryTestBase {
    function _dkgManager() internal override returns (address) {
        return address(new MockDKG());
    }

    function test_CouncilDisabled() public {
        assertEq(registry.councilAdapter(), address(0));
        assertTrue(registry.dkgAdapter() != address(0));

        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.COUNCIL;
        d.epochId = bytes12(uint96(1));
        vm.expectRevert(IProcessRegistry.CouncilDisabled.selector);
        vm.prank(ORGANIZER);
        registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            0,
            DURATION,
            MAX_VOTERS,
            _ballotMode(),
            _census(),
            METADATA_URI,
            METADATA_HASH,
            DAVINCITypes.EncryptionKey(0, 0),
            d
        );
    }
}

/// @dev The whole DKG suite again, on a registry that also has a Council adapter: the DKG
///      modes behave exactly as without it.
contract DKGWithCouncilTest is DKGTest {
    function _councilManager() internal override returns (address) {
        return address(new MockCouncilManager());
    }

    function test_CouncilAdapterPresent() public view {
        assertEq(registry.councilAdapter(), vm.computeCreateAddress(address(registry), 2));
        assertEq(registry.dkgAdapter(), vm.computeCreateAddress(address(registry), 1));
    }
}
