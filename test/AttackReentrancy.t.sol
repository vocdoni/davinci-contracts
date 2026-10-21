// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {AttackBase} from "./AttackBase.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {DKGTypes} from "../src/interfaces/dkg/DKGTypes.sol";
import {Sha256SmtLib} from "../src/libraries/Sha256SmtLib.sol";

/// @dev A DKG manager (and app manager) that calls back into the registry from every
///      state-changing hook and records how each call-back ended.
contract ReentrantDKG {
    uint32 public constant EPOCH_PREFIX = 0x00444b47;
    uint64 public epochNonce = 1;

    ProcessRegistry public registry;
    uint256 public keyX;
    uint256 public keyY;
    uint16 internal next;
    bytes4[] public caught;
    bytes public reentryCall;

    function setRegistry(ProcessRegistry r, uint256 x, uint256 y) external {
        registry = r;
        (keyX, keyY) = (x, y);
    }

    function setReentry(bytes calldata data) external {
        reentryCall = data;
    }

    function caughtCount() external view returns (uint256) {
        return caught.length;
    }

    function _reenter() internal {
        if (reentryCall.length == 0) return;
        (bool ok, bytes memory err) = address(registry).call(reentryCall);
        caught.push(ok ? bytes4(0xffffffff) : bytes4(err));
    }

    function appManager() external view returns (address) {
        return address(this);
    }

    function getPoolStatus(bytes12) external pure returns (uint8) {
        return 0;
    }

    function getPoolKey(bytes12, uint8) external view returns (uint256, uint256) {
        return (keyX, keyY);
    }

    function registerApplication(
        bytes12,
        bytes32,
        DKGTypes.AppPolicy calldata,
        uint256,
        uint256,
        uint256,
        uint256,
        uint256
    ) external {
        _reenter();
    }

    function getApplicationKey(bytes12, bytes32) external view returns (uint256, uint256) {
        return (keyX, keyY);
    }

    function submitCiphertext(bytes12, bytes32, uint256, uint256, uint256, uint256) external returns (uint16) {
        _reenter();
        return ++next;
    }

    function getCombinedDecryption(bytes12, bytes32, uint16 index)
        external
        pure
        returns (DKGTypes.CombinedDecryptionRecord memory r)
    {
        r.ciphertextIndex = index;
        r.completed = true;
        r.plaintext = index;
    }

    function revealOrganizerSecret(bytes12, bytes32, uint256) external {
        _reenter();
    }
}

/// @dev Re-entrancy from the DKG side: every guarded entry point refuses a nested call, and
///      the unguarded organizer calls see the process in a state they cannot exploit.
contract AttackReentrancyTest is RegistryTestBase {
    using stdJson for string;

    ReentrantDKG internal evil;
    string internal dkg;

    function _dkgManager() internal override returns (address) {
        evil = new ReentrantDKG();
        return address(evil);
    }

    function setUp() public override {
        super.setUp();
        dkg = vm.readFile("test/vectors/dkg.json");
        evil.setRegistry(
            registry,
            vm.parseUint(dkg.readString(".pool_keys[0].rte_x")),
            vm.parseUint(dkg.readString(".pool_keys[0].rte_y"))
        );
    }

    function _mode() internal view returns (DAVINCITypes.BallotMode memory m) {
        m.numFields = uint8(dkg.readUint(".ballot_mode.num_fields"));
        m.costExponent = uint8(dkg.readUint(".ballot_mode.cost_exponent"));
        m.maxValue = dkg.readUint(".ballot_mode.max_value");
        m.maxValueSum = dkg.readUint(".ballot_mode.max_value_sum");
    }

    function _censusD() internal view returns (DAVINCITypes.Census memory c) {
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1;
        c.censusRoot = dkg.readBytes32(".census_root");
        c.censusURI = "https://example.com/census";
    }

    function _newCall() internal view returns (bytes memory) {
        DAVINCITypes.DKGParams memory d;
        d.mode = DAVINCITypes.KeyMode.DKG_AUTOMATIC;
        return abi.encodeCall(
            registry.newProcess,
            (
                DAVINCITypes.ProcessStatus.READY,
                0,
                DURATION,
                MAX_VOTERS,
                _mode(),
                _censusD(),
                METADATA_URI,
                METADATA_HASH,
                DAVINCITypes.EncryptionKey(0, 0),
                d
            )
        );
    }

    function _create() internal returns (bytes31 pid) {
        vm.prank(ORGANIZER);
        (bool ok, bytes memory out) = address(registry).call(_newCall());
        assertTrue(ok, "newProcess");
        pid = abi.decode(out, (bytes31));
    }

    /// A settled accumulator as the process's only leaf, its request ready to send.
    function _settle(bytes31 pid) internal returns (uint256[64] memory acc, bytes32[] memory s) {
        bytes32[] memory words = dkg.readBytes32Array(".settled.accumulator");
        for (uint256 i = 0; i < 64; ++i) {
            acc[i] = uint256(words[i]);
        }
        bytes32 root = Sha256SmtLib.leafHash(4, uint256(sha256(abi.encode(acc))));
        bytes32 base = keccak256(abi.encode(pid, uint256(1)));
        vm.store(address(registry), bytes32(uint256(base) + 3), root);
        assertEq(registry.getProcess(pid).latestStateRoot, root);
        s = new bytes32[](1);
        vm.warp(registry.getProcessGraceEnd(pid));
    }

    function test_NewProcessReentryRefused() public {
        evil.setReentry(_newCall());
        _create();
        assertEq(evil.caughtCount(), 1);
        assertEq(evil.caught(0), ReentrancyGuard.ReentrancyGuardReentrantCall.selector);
        // Only the outer process exists.
        assertEq(registry.processCount(), 1);
    }

    function test_RequestReentryRefused() public {
        bytes31 pid = _create();
        (uint256[64] memory acc, bytes32[] memory s) = _settle(pid);
        bytes[4] memory calls_ = [
            abi.encodeCall(registry.finalizeResultsFromDKG, (pid)),
            abi.encodeCall(registry.requestResultsDecryption, (pid, acc, s)),
            abi.encodeCall(registry.setProcessResults, (pid, new bytes(512), hex"")),
            abi.encodeCall(registry.revealProcessKey, (pid, 1))
        ];
        for (uint256 i = 0; i < calls_.length; ++i) {
            uint256 snap = vm.snapshot();
            evil.setReentry(calls_[i]);
            registry.requestResultsDecryption(pid, acc, s);
            assertGt(evil.caughtCount(), 0);
            for (uint256 j = 0; j < evil.caughtCount(); ++j) {
                assertEq(evil.caught(j), ReentrancyGuard.ReentrancyGuardReentrantCall.selector);
            }
            DAVINCITypes.Process memory p = registry.getProcess(pid);
            assertEq(uint8(p.status), uint8(DAVINCITypes.ProcessStatus.ENDED));
            assertEq(p.dkgCount, 2);
            vm.revertTo(snap);
        }
    }

    /// The unguarded organizer calls during the submission see a process already ENDED and
    /// requested: the organizer cannot cancel it from inside the DKG call either.
    function test_OrganizerCallsDuringSubmissionSeeEnded() public {
        bytes31 pid = _create();
        (uint256[64] memory acc, bytes32[] memory s) = _settle(pid);
        // The manager cannot act as the organizer, but even the status move it tries fails
        // on the status, not only on the caller.
        evil.setReentry(abi.encodeCall(registry.setProcessStatus, (pid, DAVINCITypes.ProcessStatus.CANCELED)));
        registry.requestResultsDecryption(pid, acc, s);
        assertEq(evil.caught(0), IProcessRegistry.Unauthorized.selector);
        assertEq(uint8(registry.getProcess(pid).status), uint8(DAVINCITypes.ProcessStatus.ENDED));
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);
    }
}

/// @dev An origin-3 census contract is the organizer's choice and fully untrusted: it is
///      STATICCALLed with a gas cap and only its first word is read.
contract HostileCensus {
    ProcessRegistry public registry;
    uint256 public mode;
    uint256 public answer;

    constructor(ProcessRegistry r) {
        registry = r;
    }

    function set(uint256 m, uint256 a) external {
        (mode, answer) = (m, a);
    }

    function getCensusRoot() external view returns (uint256) {
        return 1;
    }

    /// @dev Not view on purpose: the registry calls it with STATICCALL anyway.
    function getRootBlockNumber(uint256) external returns (uint256) {
        if (mode == 1) {
            // Try to settle something from inside the check: a STATICCALL forbids it.
            (bool ok,) = address(registry).call(abi.encodeWithSignature("processCount()"));
            ok;
            assembly {
                sstore(0, 1)
            }
        }
        if (mode == 2) {
            // Return 64 KiB; only the first word may be copied.
            uint256 a = answer;
            assembly {
                mstore(0, a)
                return(0, 0x10000)
            }
        }
        return answer;
    }
}

contract AttackHostileCensusTest is AttackBase {
    HostileCensus internal census;
    bytes31 internal pid;

    function setUp() public override {
        super.setUp();
        _mockOpenings();
        census = new HostileCensus(registry);
        DAVINCITypes.Census memory c;
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1;
        c.contractAddress = address(census);
        c.censusURI = "https://example.com/census";
        pid = _newProcess(vm.getBlockTimestamp(), DURATION, MAX_VOTERS, _ballotMode(), c);
        vm.roll(vm.getBlockNumber() + 5);
    }

    function test_StateChangeInsideTheCheckFails() public {
        census.set(1, vm.getBlockNumber());
        assertFalse(_trySubmit(pid, _next(pid, 1, 0, 1)));
    }

    function test_HugeReturnDataIsNotCopied() public {
        census.set(2, vm.getBlockNumber());
        uint256 g = gasleft();
        assertTrue(_trySubmit(pid, _next(pid, 1, 0, 2)));
        // Copying 64 KiB into the caller's memory would show up here.
        assertLt(g - gasleft(), 1_000_000);
    }

    /// forge-config: default.fuzz.runs = 1000
    function testFuzz_AnswerWindow(uint256 answer) public {
        uint256 created = registry.getProcess(pid).creationBlock;
        answer = answer % 3 == 0 ? bound(answer, 0, vm.getBlockNumber() + 5) : answer;
        census.set(0, answer);
        bool want = answer != 0 && answer <= vm.getBlockNumber() && answer >= created;
        assertEq(_trySubmit(pid, _next(pid, 1, 0, 3)), want);
    }
}
