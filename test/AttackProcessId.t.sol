// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {AttackBase} from "./AttackBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {ProcessIdLib} from "../src/libraries/ProcessIdLib.sol";
import {DavinciDKGAdapter} from "../src/DavinciDKGAdapter.sol";
import {CouncilAdapter} from "../src/CouncilAdapter.sol";
import {DKGTypes} from "../src/interfaces/dkg/DKGTypes.sol";
import {MockDKG} from "./mocks/MockDKG.sol";
import {MockCouncilManager} from "./mocks/MockCouncilManager.sol";
import {MockZiskVerifier} from "./mocks/MockZiskVerifier.sol";

/// @dev Process ids: layout, who can mint which id, and ids of other registries.
contract AttackProcessIdTest is AttackBase {
    function testFuzz_IdLayout(address org, uint64 nonce, uint32 prefix) public pure {
        bytes31 pid = ProcessIdLib.computeProcessId(prefix, org, nonce);
        assertEq(address(uint160(uint248(pid) >> 88)), org);
        assertEq(uint32(uint248(pid) >> 56), prefix);
        assertEq(uint56(uint248(pid)), uint56(nonce));
        assertTrue(ProcessIdLib.hasPrefix(pid, prefix));
        if (prefix != type(uint32).max) assertFalse(ProcessIdLib.hasPrefix(pid, prefix + 1));
    }

    /// Ids are per creator: other accounts creating processes never consume or collide with
    /// an organizer's next id, and every id matches getNextProcessId at its creation.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_IdsPerCreator(address[4] memory who, uint8[8] memory order) public {
        for (uint256 i = 0; i < 4; ++i) {
            if (who[i] == address(0) || who[i].code.length > 0 || uint160(who[i]) < 0x100) {
                who[i] = address(uint160(0x1000 + i));
            }
        }
        bytes31[] memory seen = new bytes31[](8);
        for (uint256 i = 0; i < 8; ++i) {
            address a = who[order[i] % 4];
            bytes31 next = registry.getNextProcessId(a);
            vm.prank(a);
            bytes31 pid = registry.newProcess(
                DAVINCITypes.ProcessStatus.READY,
                0,
                DURATION,
                MAX_VOTERS,
                _ballotMode(),
                _census(),
                METADATA_URI,
                METADATA_HASH,
                _encKey(),
                _noDkg()
            );
            assertEq(pid, next);
            assertEq(registry.getProcess(pid).organizationId, a);
            for (uint256 j = 0; j < i; ++j) {
                assertTrue(seen[j] != pid, "duplicate id");
            }
            seen[i] = pid;
        }
    }

    /// An id with another registry's prefix (or zero) is refused by every entry point.
    function testFuzz_ForeignPrefixRefused(uint248 raw) public {
        bytes31 pid = bytes31(raw);
        vm.assume(!ProcessIdLib.hasPrefix(pid, registry.pidPrefix()));
        bytes4 want = pid == bytes31(0)
            ? IProcessRegistry.InvalidProcessId.selector
            : IProcessRegistry.UnknownProcessIdPrefix.selector;
        vm.startPrank(ORGANIZER);
        vm.expectRevert(want);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.ENDED);
        vm.expectRevert(want);
        registry.setProcessDuration(pid, 1);
        vm.expectRevert(want);
        registry.setProcessMaxVoters(pid, 1);
        vm.expectRevert(want);
        registry.setProcessGrace(pid, GRACE);
        vm.expectRevert(want);
        registry.setProcessMetadata(pid, METADATA_URI, METADATA_HASH);
        vm.expectRevert(want);
        registry.setProcessResults(pid, new bytes(512), hex"");
        vm.expectRevert(want);
        registry.finalizeResultsFromDKG(pid);
        vm.expectRevert(want);
        registry.submitStateTransition(pid, new bytes(512), hex"", new bytes[](0), new bytes32[](0), new bytes[](0));
        vm.stopPrank();
    }

    /// The same creator and nonce on a second registry of the same chain gets another id,
    /// and neither registry accepts the other's.
    function test_SecondRegistryIdsAreDisjoint() public {
        bytes31 a = _fixtureProcess();
        ProcessRegistry r2 = new ProcessRegistry(
            CHAIN_ID,
            address(new MockZiskVerifier()),
            BATCH_VK,
            RESULTS_VK,
            ROOT_C,
            keccak256("vk"),
            address(0),
            address(0),
            GRACE,
            GRACE_FLOOR,
            GRACE_CEIL,
            GRACE_MAX_TOTAL,
            NOTICE_MIN,
            address(0)
        );
        assertTrue(r2.pidPrefix() != registry.pidPrefix());
        assertTrue(r2.getNextProcessId(ORGANIZER) != a);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.UnknownProcessIdPrefix.selector);
        r2.setProcessStatus(a, DAVINCITypes.ProcessStatus.CANCELED);
    }
}

/// @dev Application ids of the DKG adapter: nobody can take a future process's id, across
///      registries they are disjoint, and a squatter on a look-alike id blocks nothing.
contract AttackAidTest is RegistryTestBase {
    using stdJson for string;

    MockDKG internal mock;
    bytes12 internal eid1;
    string internal dkg;

    function _dkgManager() internal override returns (address) {
        mock = new MockDKG();
        return address(mock);
    }

    function setUp() public override {
        super.setUp();
        dkg = vm.readFile("test/vectors/dkg.json");
        uint256[2][] memory keys = new uint256[2][](4);
        for (uint256 i = 0; i < 4; ++i) {
            keys[i][0] = vm.parseUint(dkg.readString(string.concat(".pool_keys[", vm.toString(i % 2), "].rte_x")));
            keys[i][1] = vm.parseUint(dkg.readString(string.concat(".pool_keys[", vm.toString(i % 2), "].rte_y")));
        }
        eid1 = mock.newEpoch(true, keys);
    }

    function _policy(address submitter) internal pure returns (DKGTypes.AppPolicy memory p) {
        address[] memory s = new address[](1);
        s[0] = submitter;
        p.mode = DKGTypes.AppMode.Automatic;
        p.submitters = s;
        p.maxCiphertexts = 16;
    }

    function _automatic(ProcessRegistry r) internal returns (bytes31) {
        DAVINCITypes.BallotMode memory m;
        m.numFields = uint8(dkg.readUint(".ballot_mode.num_fields"));
        m.costExponent = uint8(dkg.readUint(".ballot_mode.cost_exponent"));
        m.maxValue = dkg.readUint(".ballot_mode.max_value");
        m.maxValueSum = dkg.readUint(".ballot_mode.max_value_sum");
        DAVINCITypes.Census memory c;
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1;
        c.censusRoot = dkg.readBytes32(".census_root");
        c.censusURI = "https://example.com/census";
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        vm.prank(ORGANIZER);
        return r.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            0,
            DURATION,
            MAX_VOTERS,
            m,
            c,
            METADATA_URI,
            METADATA_HASH,
            DAVINCITypes.EncryptionKey(0, 0),
            d
        );
    }

    /// A squatter registers every look-alike of the victim's future id it can (its own
    /// address under the same salt, and the exact id): the exact id is refused, the
    /// look-alike lands in the squatter's namespace and the victim's process still binds.
    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_SquatterCannotBlockAFutureProcess(address squatter, uint8 ahead) public {
        vm.assume(squatter != address(0) && squatter != registry.dkgAdapter() && squatter.code.length == 0);
        // Burn a few of the organizer's ids first so the target is a later one.
        ahead = uint8(bound(ahead, 0, 2));
        for (uint256 i = 0; i < ahead; ++i) {
            _automatic(registry);
        }
        bytes31 pid = registry.getNextProcessId(ORGANIZER);
        bytes32 aid = registry.aidFor(pid);
        bytes32 lookAlike = bytes32((uint256(aid) >> 160 << 160) | uint256(uint160(squatter)));

        vm.startPrank(squatter);
        vm.expectRevert(MockDKG.InvalidApplication.selector);
        mock.registerApplication(eid1, aid, _policy(squatter), 0, 1, 0, 0, 0);
        mock.registerApplication(eid1, lookAlike, _policy(squatter), 0, 1, 0, 0, 0);
        vm.stopPrank();

        assertEq(_automatic(registry), pid);
        assertEq(registry.getProcess(pid).dkgAid, aid);
    }

    /// Two registries sharing one DKG manager have disjoint application ids for the same
    /// process id bytes (salt and adapter address both differ).
    function test_AidsDisjointAcrossRegistries() public {
        ProcessRegistry r2 = new ProcessRegistry(
            CHAIN_ID,
            address(new MockZiskVerifier()),
            BATCH_VK,
            RESULTS_VK,
            ROOT_C,
            fixture.readBytes32(".ballot_vk_hash"),
            address(mock),
            address(0),
            GRACE,
            GRACE_FLOOR,
            GRACE_CEIL,
            GRACE_MAX_TOTAL,
            NOTICE_MIN,
            address(0)
        );
        bytes31 pid = registry.getNextProcessId(ORGANIZER);
        assertTrue(registry.aidFor(pid) != r2.aidFor(pid));
        bytes31 a = _automatic(registry);
        bytes31 b = _automatic(r2);
        assertTrue(registry.getProcess(a).dkgAid != r2.getProcess(b).dkgAid);
        // Each adapter is the only submitter of its own applications.
        assertEq(address(uint160(uint256(registry.getProcess(a).dkgAid))), registry.dkgAdapter());
        assertEq(address(uint160(uint256(r2.getProcess(b).dkgAid))), r2.dkgAdapter());
    }
}

/// @dev Council bindings are per (adapter, process id): another allowed adapter binding the
///      victim's future id first does not block the victim's own process.
contract AttackCouncilBindingTest is RegistryTestBase {
    using stdJson for string;

    bytes12 internal constant CID = bytes12(0x00c0c1a7e0000000000000a1);
    MockCouncilManager internal council;

    function _councilManager() internal override returns (address) {
        council = new MockCouncilManager();
        return address(council);
    }

    function setUp() public override {
        super.setUp();
        string memory dkg = vm.readFile("test/vectors/dkg.json");
        council.newCeremony(
            CID, vm.parseUint(dkg.readString(".pool_keys[1].te_x")), vm.parseUint(dkg.readString(".pool_keys[1].te_y"))
        );
        council.allowAdapter(CID, registry.councilAdapter());
        council.authorizeCreator(CID, ORGANIZER);
    }

    function test_OtherAdapterCannotPreBindTheVictim() public {
        address rogue = address(0xBAD);
        council.allowAdapter(CID, rogue);
        bytes31 pid = registry.getNextProcessId(ORGANIZER);
        vm.prank(rogue);
        (bytes32 rogueReq,,) = council.bindProcess(CID, pid, ORGANIZER);

        DAVINCITypes.BallotMode memory m = _ballotMode();
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.COUNCIL;
        d.epochId = CID;
        vm.prank(ORGANIZER);
        bytes31 got = registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            0,
            DURATION,
            MAX_VOTERS,
            m,
            _census(),
            METADATA_URI,
            METADATA_HASH,
            DAVINCITypes.EncryptionKey(0, 0),
            d
        );
        assertEq(got, pid);
        bytes32 req = registry.getProcess(pid).dkgAid;
        assertTrue(req != rogueReq);
        (bytes12 cid,, bytes31 bound_) = CouncilAdapter(registry.councilAdapter()).bindings(req);
        assertEq(cid, CID);
        assertEq(bound_, pid);
    }
}
