// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {BjjFormLib} from "../src/libraries/BjjFormLib.sol";
import {Sha256SmtLib} from "../src/libraries/Sha256SmtLib.sol";
import {DavinciDKGAdapter} from "../src/DavinciDKGAdapter.sol";
import {MockDKG, MockBjj} from "./mocks/MockDKG.sol";

/// @dev External wrapper so tests can call the internal library with calldata arrays.
contract SmtHarness {
    function verify(bytes32 root, uint64 key, uint256 value, bytes32[] calldata siblings) external pure returns (bool) {
        return Sha256SmtLib.verifyInclusion(root, key, value, siblings);
    }

    function leaf(uint64 key, uint256 value) external pure returns (bytes32) {
        return Sha256SmtLib.leafHash(key, value);
    }
}

/// @dev Registry deployed without a DKG manager: the DKG surface must be disabled.
contract DKGDisabledTest is RegistryTestBase {
    function test_DKGDisabled() public {
        assertEq(registry.dkgAdapter(), address(0));

        vm.expectRevert(IProcessRegistry.DKGDisabled.selector);
        registry.aidFor(bytes31(bytes32(uint256(1) << 8)));

        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        vm.expectRevert(IProcessRegistry.DKGDisabled.selector);
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

contract DKGTest is RegistryTestBase {
    using stdJson for string;

    // The davinci-dkg test vectors: P_1 = 1000004·G plus PK_org = 12345·G, reduced form.
    uint256 internal constant DKG_VECTOR_LOCKED_X =
        13419422839618223618338033640401495486422796103482975513089995370826299959057;
    uint256 internal constant DKG_VECTOR_LOCKED_Y =
        3538309003386952231832535335923184813316582788885500404757678728572039432940;
    // Circomlib base point B8 (x differs from the reduced GENERATOR_X, y is shared).
    uint256 internal constant B8_X = 5299619240641551281634865583518297030282874472190772894086521144482721001553;

    MockDKG internal mock;
    DavinciDKGAdapter internal adapter;
    bytes12 internal eid1;
    string internal dkg;

    function _dkgManager() internal override returns (address) {
        mock = new MockDKG();
        return address(mock);
    }

    function setUp() public override {
        super.setUp();
        dkg = vm.readFile("test/vectors/dkg.json");
        adapter = DavinciDKGAdapter(registry.dkgAdapter());
        uint256[2][] memory keys = new uint256[2][](2);
        keys[0] = _rtePoint(".pool_keys[0]");
        keys[1] = _rtePoint(".pool_keys[1]");
        eid1 = mock.newEpoch(true, keys);
    }

    // --- fixture readers -------------------------------------------------------

    function _rtePoint(string memory p) internal view returns (uint256[2] memory pt) {
        pt[0] = vm.parseUint(dkg.readString(string.concat(p, ".rte_x")));
        pt[1] = vm.parseUint(dkg.readString(string.concat(p, ".rte_y")));
    }

    function _tePoint(string memory p) internal view returns (uint256 x, uint256 y) {
        x = vm.parseUint(dkg.readString(string.concat(p, ".te_x")));
        y = vm.parseUint(dkg.readString(string.concat(p, ".te_y")));
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

    /// @dev The accumulator and siblings of inclusion case p (".genesis", ".settled", ...).
    function _inclusion(string memory p)
        internal
        view
        returns (bytes32 root, uint256[64] memory acc, bytes32[] memory siblings)
    {
        root = dkg.readBytes32(string.concat(p, ".root"));
        bytes32[] memory words = dkg.readBytes32Array(string.concat(p, ".accumulator"));
        assertEq(words.length, 64, "accumulator length");
        for (uint256 i = 0; i < 64; ++i) {
            acc[i] = uint256(words[i]);
        }
        siblings = dkg.readBytes32Array(string.concat(p, ".siblings"));
    }

    // --- helpers -----------------------------------------------------------------

    function _newDkgProcess(DAVINCITypes.DKGParams memory d) internal returns (bytes31) {
        vm.prank(ORGANIZER);
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

    function _automaticProcess() internal returns (bytes31) {
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        return _newDkgProcess(d);
    }

    function _lockedProcess() internal returns (bytes31) {
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_LOCKED;
        d.epochId = eid1;
        uint256[2] memory org = _rtePoint(".org_pk");
        (d.orgPKx, d.orgPKy) = (org[0], org[1]);
        return _newDkgProcess(d);
    }

    /// @dev Overwrites p.latestStateRoot (struct slot 3, processes mapping at slot 1) to
    ///      simulate settled transitions; the readback assert guards the layout.
    function _forceRoot(bytes31 pid, bytes32 root) internal {
        bytes32 base = keccak256(abi.encode(pid, uint256(1)));
        vm.store(address(registry), bytes32(uint256(base) + 3), root);
        assertEq(registry.getProcess(pid).latestStateRoot, root, "latestStateRoot slot moved");
    }

    // --- BjjFormLib and the mock's curve math -------------------------------------

    function test_BjjFormLib_Constants() public pure {
        assertEq(mulmod(BjjFormLib.K, BjjFormLib.K_INV, BjjFormLib.Q), 1);
        assertEq(BjjFormLib.toRTE(B8_X), MockBjj.GX);
        assertEq(BjjFormLib.toTE(MockBjj.GX), B8_X);
    }

    function test_MockBjj_MatchesGoVectors() public view {
        // 1000003·G and 12345·G recomputed with the mock's affine math must match the
        // go-sdk gnark values in the fixture.
        (uint256 x, uint256 y) = MockBjj.mulBase(1000003);
        uint256[2] memory p0 = _rtePoint(".pool_keys[0]");
        assertEq(x, p0[0]);
        assertEq(y, p0[1]);
        (x, y) = MockBjj.mulBase(12345);
        uint256[2] memory org = _rtePoint(".org_pk");
        assertEq(x, org[0]);
        assertEq(y, org[1]);
    }

    // --- newProcess --------------------------------------------------------------

    function test_NewProcess_SequencerRejectsDKGParams() public {
        DAVINCITypes.DKGParams memory d;
        d.epochId = eid1;
        vm.expectRevert(IProcessRegistry.InvalidDKGParams.selector);
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
            _encKey(),
            d
        );
    }

    function test_NewProcess_DKGRejectsNonZeroKey() public {
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
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
            _encKey(),
            d
        );
    }

    function test_NewProcess_DKGAutomatic() public {
        bytes31 pid = registry.getNextProcessId(ORGANIZER);
        // The metadata is logged after the adapter's registration, as in sequencer mode.
        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessCreated(pid, ORGANIZER);
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessMetadataUpdated(pid, METADATA_URI, METADATA_HASH);
        assertEq(_automaticProcess(), pid);
        assertEq(bytes32(pid), bytes32(dkg.readBytes(".process_id")), "fixture pid");

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.metadataURI, METADATA_URI);
        assertEq(p.metadataHash, METADATA_HASH);
        assertEq(uint8(p.keyMode), uint8(DAVINCITypes.KeyMode.DKG_AUTOMATIC));
        assertEq(p.dkgEpochId, eid1);
        assertEq(p.dkgAid, registry.aidFor(pid));

        // The stored key is pool key 0 converted to circomlib form, and the genesis
        // root was computed with it (pinned by the Go fixture).
        (uint256 teX, uint256 teY) = _tePoint(".pool_keys[0]");
        assertEq(p.encryptionKey.x, teX);
        assertEq(p.encryptionKey.y, teY);
        assertEq(p.latestStateRoot, dkg.readBytes32(".genesis.root"));
        assertEq(
            p.latestStateRoot,
            registry.genesisRoot(
                pid, _dkgBallotMode(), p.encryptionKey, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
            )
        );
    }

    function test_NewProcess_DKGLocked() public {
        _automaticProcess(); // claims pool key 0
        bytes31 pid = _lockedProcess(); // claims pool key 1
        assertEq(bytes32(pid), bytes32(dkg.readBytes(".process_id_locked")), "fixture locked pid");

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(uint8(p.keyMode), uint8(DAVINCITypes.KeyMode.DKG_LOCKED));

        // PK_aid = P_1 + PK_org: the mock's real point addition must match both the
        // Go fixture and the davinci-dkg test vector.
        (uint256 rx, uint256 ry) = mock.getApplicationKey(eid1, p.dkgAid);
        assertEq(rx, DKG_VECTOR_LOCKED_X);
        assertEq(ry, DKG_VECTOR_LOCKED_Y);
        (uint256 teX, uint256 teY) = _tePoint(".locked_key");
        assertEq(p.encryptionKey.x, teX);
        assertEq(p.encryptionKey.y, teY);
        assertEq(
            p.latestStateRoot,
            registry.genesisRoot(
                pid, _dkgBallotMode(), p.encryptionKey, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
            )
        );
    }

    function test_NewProcess_DKGLockedNeedsLiveEpoch() public {
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_LOCKED;
        d.epochId = mock.newEpoch(false, new uint256[2][](0)); // not Live
        vm.expectRevert(MockDKG.InvalidPhase.selector);
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
            DAVINCITypes.EncryptionKey(0, 0),
            d
        );
    }

    // --- adapter views and access control ------------------------------------------

    function test_Adapter_Wiring() public view {
        assertEq(adapter.registry(), address(registry));
        assertEq(address(adapter.manager()), address(mock));
        assertEq(address(adapter.appManager()), address(mock));
    }

    function test_AidFor() public view {
        bytes31 pid = bytes31(dkg.readBytes(".process_id"));
        bytes32 aid = registry.aidFor(pid);
        uint256 expected = uint256(keccak256(abi.encode(block.chainid, address(registry), pid))) % BjjFormLib.Q;
        if (expected == 0) expected = 1;
        assertEq(aid, bytes32(expected));
        assertTrue(uint256(aid) != 0 && uint256(aid) < BjjFormLib.Q);
        assertEq(aid, adapter.aidFor(pid));
    }

    function test_Adapter_OnlyRegistry() public {
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        vm.expectRevert(DavinciDKGAdapter.NotRegistry.selector);
        adapter.register(bytes31(uint248(1)), d);
        vm.expectRevert(DavinciDKGAdapter.NotRegistry.selector);
        adapter.submit(eid1, bytes32(uint256(1)), new uint256[4][](0));
        vm.expectRevert(DavinciDKGAdapter.NotRegistry.selector);
        adapter.reveal(eid1, bytes32(uint256(1)), 1);
    }

    function test_RegistrationEpoch_ScansBackAndExhausts() public {
        assertEq(adapter.registrationEpoch(), eid1);

        // A newer, not-yet-Live epoch is skipped.
        bytes12 eid2 = mock.newEpoch(false, new uint256[2][](0));
        assertEq(adapter.registrationEpoch(), eid1);

        // Once Live with a free pool key, the newest epoch wins.
        uint256[2][] memory keys = new uint256[2][](1);
        keys[0] = _rtePoint(".pool_keys[1]");
        bytes12 eid3 = mock.newEpoch(true, keys);
        assertEq(adapter.registrationEpoch(), eid3);
        assertEq(registry.getProcess(_automaticProcess()).dkgEpochId, eid3);
        eid2; // silence

        // Pool of eid3 is now spent (1 key); eid1 still has free keys.
        assertEq(adapter.registrationEpoch(), eid1);

        // With every pool spent there is nowhere to register.
        mock.setPoolNext(eid1, 16);
        vm.expectRevert(DavinciDKGAdapter.NoLiveEpoch.selector);
        adapter.registrationEpoch();
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        vm.expectRevert(DavinciDKGAdapter.NoLiveEpoch.selector);
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
            DAVINCITypes.EncryptionKey(0, 0),
            d
        );
    }

    // --- reveal ----------------------------------------------------------------------

    function test_RevealProcessKey() public {
        _automaticProcess();
        bytes31 pid = _lockedProcess();
        bytes32 aid = registry.aidFor(pid);

        vm.expectRevert(MockDKG.InvalidOrganizerSecret.selector);
        registry.revealProcessKey(pid, 4242);

        registry.revealProcessKey(pid, dkg.readUint(".org_sk"));
        assertTrue(mock.revealed(eid1, aid));

        vm.expectRevert(MockDKG.AlreadyRevealed.selector);
        registry.revealProcessKey(pid, dkg.readUint(".org_sk"));
    }

    function test_RevealProcessKey_WrongMode() public {
        bytes31 pid = _automaticProcess();
        vm.expectRevert(IProcessRegistry.InvalidKeyMode.selector);
        registry.revealProcessKey(pid, 12345);

        bytes31 seqPid = _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), _census());
        vm.expectRevert(IProcessRegistry.InvalidKeyMode.selector);
        registry.revealProcessKey(seqPid, 12345);
    }

    // --- requestResultsDecryption -------------------------------------------------------

    function test_Request_ZeroVotesFinalizesImmediately() public {
        bytes31 pid = _automaticProcess();
        _warpToGraceEnd(pid);
        (, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".genesis");

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessStatusChanged(
            pid, DAVINCITypes.ProcessStatus.READY, DAVINCITypes.ProcessStatus.ENDED
        );
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ResultsDecryptionRequested(pid, eid1, p.dkgAid, 0, 0);
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessStatusChanged(
            pid, DAVINCITypes.ProcessStatus.ENDED, DAVINCITypes.ProcessStatus.RESULTS
        );
        registry.requestResultsDecryption(pid, acc, siblings);

        p = registry.getProcess(pid);
        assertEq(uint8(p.status), uint8(DAVINCITypes.ProcessStatus.RESULTS));
        assertEq(p.result.length, 4); // numFields entries, like setProcessResults
        for (uint256 i = 0; i < 4; ++i) {
            assertEq(p.result[i], 0);
        }
        assertEq(mock.ctCount(eid1, p.dkgAid), 0);
    }

    function test_Request_SettledAccumulatorAndFinalize() public {
        bytes31 pid = _automaticProcess();
        (bytes32 root, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".settled");
        _forceRoot(pid, root);
        _warpToGraceEnd(pid);

        bytes32 aid = registry.aidFor(pid);
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessStatusChanged(
            pid, DAVINCITypes.ProcessStatus.READY, DAVINCITypes.ProcessStatus.ENDED
        );
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ResultsDecryptionRequested(pid, eid1, aid, 1, 2);
        registry.requestResultsDecryption(pid, acc, siblings);

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.dkgFirstIndex, 1); // the DKG assigns 1-based indices
        assertEq(p.dkgCount, 2); // fields 0 and 2
        assertEq(p.dkgZeroSkipped, (1 << 1) | (1 << 3)); // fields 1 and 3 are identity
        assertTrue(p.dkgResultsRequested);
        // ENDED, not finalized: the plaintexts are public on the DKG from here, so the
        // organizer must not be able to read the tally and cancel.
        assertEq(uint8(p.status), uint8(DAVINCITypes.ProcessStatus.ENDED));
        assertEq(mock.ctCount(eid1, aid), 2);

        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);

        vm.expectRevert(IProcessRegistry.ResultsAlreadyRequested.selector);
        registry.requestResultsDecryption(pid, acc, siblings);

        // Combines incomplete: finalize refuses, even with one of two plaintexts set.
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pid);
        mock.setPlaintext(eid1, aid, 1, 77);
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pid);

        mock.setPlaintext(eid1, aid, 2, 99);
        uint256[] memory expected = new uint256[](4); // numFields entries, like setProcessResults
        expected[0] = 77;
        expected[2] = 99;
        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessResultsSet(pid, address(this), expected);
        registry.finalizeResultsFromDKG(pid);

        p = registry.getProcess(pid);
        assertEq(uint8(p.status), uint8(DAVINCITypes.ProcessStatus.RESULTS));
        assertEq(p.result.length, 4);
        assertEq(p.result[0], 77);
        assertEq(p.result[1], 0);
        assertEq(p.result[2], 99);
        assertEq(p.result[3], 0);

        // Terminal: nothing else can touch the results.
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.finalizeResultsFromDKG(pid);
        vm.expectRevert(IProcessRegistry.ResultsAlreadyRequested.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
    }

    function test_Request_RejectsWrongProofs() public {
        bytes31 pid = _automaticProcess();
        _warpToGraceEnd(pid);
        (, uint256[64] memory genesisAcc, bytes32[] memory genesisSiblings) = _inclusion(".genesis");
        (, uint256[64] memory settledAcc, bytes32[] memory settledSiblings) = _inclusion(".settled");

        // Wrong accumulator for the stored (genesis) root.
        vm.expectRevert(IProcessRegistry.InvalidInclusionProof.selector);
        registry.requestResultsDecryption(pid, settledAcc, genesisSiblings);

        // Right accumulator and siblings, wrong root.
        vm.expectRevert(IProcessRegistry.InvalidInclusionProof.selector);
        registry.requestResultsDecryption(pid, settledAcc, settledSiblings);

        // Tampered sibling.
        bytes32[] memory tampered = genesisSiblings;
        tampered[0] = tampered[0] ^ bytes32(uint256(1));
        vm.expectRevert(IProcessRegistry.InvalidInclusionProof.selector);
        registry.requestResultsDecryption(pid, genesisAcc, tampered);

        // Tampered accumulator coordinate.
        tampered[0] = tampered[0] ^ bytes32(uint256(1));
        genesisAcc[5] = 2;
        vm.expectRevert(IProcessRegistry.InvalidInclusionProof.selector);
        registry.requestResultsDecryption(pid, genesisAcc, genesisSiblings);
    }

    function test_Request_RejectsNonCanonicalCoordinate() public {
        bytes31 pid = _automaticProcess();
        _warpToGraceEnd(pid);
        (, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".genesis");
        acc[1] += BjjFormLib.Q; // y + Q encodes the same point; the leaf binds raw bytes
        vm.expectRevert(IProcessRegistry.InvalidAccumulator.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
    }

    function test_Request_RejectsIdentityC1WithNonIdentityC2() public {
        bytes31 pid = _automaticProcess();
        (bytes32 root, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".bad_c1_zero");
        _forceRoot(pid, root);
        _warpToGraceEnd(pid);
        vm.expectRevert(IProcessRegistry.InvalidAccumulator.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
    }

    function test_Request_RejectsIdentityC2WithNonIdentityC1() public {
        bytes31 pid = _automaticProcess();
        (bytes32 root, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".bad_c2_zero");
        _forceRoot(pid, root);
        _warpToGraceEnd(pid);
        vm.expectRevert(IProcessRegistry.InvalidAccumulator.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
    }

    function test_Request_RejectsNonContiguousIndices() public {
        bytes31 pid = _automaticProcess();
        (bytes32 root, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".settled");
        _forceRoot(pid, root);
        _warpToGraceEnd(pid);
        mock.setSkipAtCall(2); // the second submit lands after someone else's ciphertext
        vm.expectRevert(DavinciDKGAdapter.NonContiguousIndex.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
    }

    function test_Request_EndedBeforeEndTime() public {
        bytes31 pid = _automaticProcess();
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.ENDED); // organizer ends early
        (, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".genesis");
        // ENDED moved the end to now, and the grace window runs from there.
        uint256 graceEnd = registry.getProcessGraceEnd(pid);
        assertEq(graceEnd, registry.getProcessEndTime(pid) + GRACE);
        vm.expectRevert(IProcessRegistry.GraceOpen.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
        vm.warp(graceEnd);
        registry.requestResultsDecryption(pid, acc, siblings);
        assertEq(uint8(registry.getProcess(pid).status), uint8(DAVINCITypes.ProcessStatus.RESULTS));
    }

    /// @dev Past the end time the request waits for the grace window, and so does finalize
    ///      (which otherwise answers ResultsNotReady before any request).
    function test_RequestAndFinalize_GraceOpen() public {
        bytes31 pid = _automaticProcess();
        (bytes32 root, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".settled");
        _forceRoot(pid, root);
        uint256 graceEnd = registry.getProcessGraceEnd(pid);
        assertEq(graceEnd, registry.getProcessEndTime(pid) + GRACE);

        vm.warp(registry.getProcessEndTime(pid));
        vm.expectRevert(IProcessRegistry.GraceOpen.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
        vm.expectRevert(IProcessRegistry.GraceOpen.selector);
        registry.finalizeResultsFromDKG(pid);

        vm.warp(graceEnd - 1);
        vm.expectRevert(IProcessRegistry.GraceOpen.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
        vm.expectRevert(IProcessRegistry.GraceOpen.selector);
        registry.finalizeResultsFromDKG(pid);

        vm.warp(graceEnd);
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pid);
        registry.requestResultsDecryption(pid, acc, siblings);
        assertTrue(registry.getProcess(pid).dkgResultsRequested);
    }

    function test_Request_TimeAndStatusRules() public {
        bytes31 pid = _automaticProcess();
        (, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".genesis");

        // Still running.
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.requestResultsDecryption(pid, acc, siblings);

        // Canceled processes cannot be tallied.
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);
        _warpToGraceEnd(pid);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
    }

    function test_Request_SequencerMode() public {
        bytes31 pid = _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), _census());
        _warpToGraceEnd(pid);
        (, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".genesis");
        vm.expectRevert(IProcessRegistry.InvalidKeyMode.selector);
        registry.requestResultsDecryption(pid, acc, siblings);
        vm.expectRevert(IProcessRegistry.InvalidKeyMode.selector);
        registry.finalizeResultsFromDKG(pid);
    }

    function test_Finalize_BeforeRequest() public {
        bytes31 pid = _automaticProcess();
        _warpToGraceEnd(pid);
        vm.expectRevert(IProcessRegistry.ResultsNotReady.selector);
        registry.finalizeResultsFromDKG(pid);
    }

    function test_SetProcessResults_DKGMode() public {
        bytes31 pid = _automaticProcess();
        _warpToGraceEnd(pid);
        uint64[16] memory values;
        bytes memory pv = _resultsPublics(registry.getProcess(pid).latestStateRoot, values);
        vm.expectRevert(IProcessRegistry.InvalidKeyMode.selector);
        registry.setProcessResults(pid, pv, new bytes(768));
    }

    // --- Sha256SmtLib.verifyInclusion edges ----------------------------------------------

    function test_VerifyInclusion_Edges() public {
        SmtHarness h = new SmtHarness();
        (bytes32 root, uint256[64] memory acc, bytes32[] memory siblings) = _inclusion(".genesis");
        uint256 leaf = uint256(sha256(abi.encode(acc)));

        assertTrue(h.verify(root, 4, leaf, siblings));
        assertFalse(h.verify(root, 5, leaf, siblings)); // pinned key
        assertFalse(h.verify(root, 4, leaf ^ 1, siblings));
        assertFalse(h.verify(root, 4, leaf, new bytes32[](0))); // empty

        // Truncating the trailing zero leaves the last sibling non-zero: rejected.
        uint256 n = siblings.length;
        bytes32[] memory noPad = new bytes32[](n - 1);
        for (uint256 i = 0; i < n - 1; ++i) {
            noPad[i] = siblings[i];
        }
        assertFalse(h.verify(root, 4, leaf, noPad));

        // Extra zero padding is harmless; 65 entries are not.
        bytes32[] memory padded = new bytes32[](64);
        for (uint256 i = 0; i < n; ++i) {
            padded[i] = siblings[i];
        }
        assertTrue(h.verify(root, 4, leaf, padded));
        bytes32[] memory tooLong = new bytes32[](65);
        for (uint256 i = 0; i < n; ++i) {
            tooLong[i] = siblings[i];
        }
        assertFalse(h.verify(root, 4, leaf, tooLong));

        // A single-leaf tree: the root is the leaf hash, proven with one zero sibling.
        bytes32 single = h.leaf(7, 42);
        bytes32[] memory zero = new bytes32[](1);
        assertTrue(h.verify(single, 7, 42, zero));
        assertFalse(h.verify(single, 7, 43, zero));
    }
}
