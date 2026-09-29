// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {RegistryTestBase} from "./RegistryTestBase.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {ProcessIdLib} from "../src/libraries/ProcessIdLib.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";

/// @dev Process lifecycle: status machine, duration, max voters, ballot mode and census
///      validation. Settlement lives in Transition.t.sol and Results.t.sol.
contract ProcessRegistryTest is RegistryTestBase {
    ProcessRegistry public processRegistry;

    bytes32 internal constant CENSUS_ROOT = 0x2bc6f255d02b18329662a71d7d66c8ce06fe984607fd4fe26ef33ab93a78f683;
    bytes32 internal constant CSP_ROOT = bytes32(uint256(0xC5C5C));

    DAVINCITypes.BallotMode public defaultBallotMode = DAVINCITypes.BallotMode({
        uniqueValues: false,
        numFields: 5,
        groupSize: 0,
        costExponent: 2,
        maxValue: 16,
        minValue: 0,
        maxValueSum: 1280,
        minValueSum: 5
    });

    function setUp() public override {
        super.setUp();
        processRegistry = registry;
    }

    function createTestProcess(DAVINCITypes.BallotMode memory ballotMode, DAVINCITypes.CensusOrigin censusOrigin)
        internal
        returns (bytes31)
    {
        DAVINCITypes.Census memory cen = DAVINCITypes.Census({
            onchainAllowAnyValidRoot: false,
            censusOrigin: censusOrigin,
            censusRoot: censusOrigin == DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1 ? CSP_ROOT : CENSUS_ROOT,
            censusURI: "https://example.com/census",
            contractAddress: address(0)
        });

        return processRegistry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            block.timestamp, // current time
            1000,
            10000,
            ballotMode,
            cen,
            METADATA_URI,
            METADATA_HASH,
            _encKey(),
            _noDkg()
        );
    }

    // ========== Process Status Tests ==========

    function test_SetProcessStatus_NonExistentProcess() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.setProcessStatus(bytes31(0), DAVINCITypes.ProcessStatus.ENDED);

        bytes32 h = keccak256(abi.encodePacked(CHAIN_ID, address(processRegistry)));
        uint32 prefix = uint32(uint256(h));
        bytes31 invalidProcessId =
            ProcessIdLib.computeProcessId(prefix, address(0x1234567890123456789012345678901234567890), 1);

        vm.expectRevert(IProcessRegistry.ProcessNotFound.selector);
        processRegistry.setProcessStatus(invalidProcessId, DAVINCITypes.ProcessStatus.ENDED);
    }

    function test_SetProcessStatus_UnknownProcessIdPrefix() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.setProcessStatus(bytes31(0), DAVINCITypes.ProcessStatus.ENDED);

        vm.expectRevert(IProcessRegistry.UnknownProcessIdPrefix.selector);
        processRegistry.setProcessStatus(bytes31(uint248((uint248(1) << 247) | 1)), DAVINCITypes.ProcessStatus.ENDED);
    }

    function test_SetProcessStatus_NotAdmin() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);
        vm.stopPrank();
    }

    function test_SetProcessStatus_SameStatus() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.READY);
    }

    function test_SetProcessStatus_FromReady() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        // READY -> PAUSED (valid)
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.PAUSED);
        assertEq(uint256(processRegistry.getProcess(processId).status), uint256(DAVINCITypes.ProcessStatus.PAUSED));

        // Reset to READY
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.READY);

        // READY -> CANCELED (valid)
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.CANCELED);
        assertEq(uint256(processRegistry.getProcess(processId).status), uint256(DAVINCITypes.ProcessStatus.CANCELED));

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        // READY -> ENDED (valid)
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);
        assertEq(uint256(processRegistry.getProcess(processId).status), uint256(DAVINCITypes.ProcessStatus.ENDED));

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        // READY -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromPaused() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        // Set initial state to PAUSED
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.PAUSED);

        // PAUSED -> READY (valid)
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.READY);
        assertEq(uint256(processRegistry.getProcess(processId).status), uint256(DAVINCITypes.ProcessStatus.READY));

        // Reset to PAUSED
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.PAUSED);

        // PAUSED -> CANCELED (valid)
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.CANCELED);
        assertEq(uint256(processRegistry.getProcess(processId).status), uint256(DAVINCITypes.ProcessStatus.CANCELED));

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.PAUSED);

        // PAUSED -> ENDED (valid)
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);
        assertEq(uint256(processRegistry.getProcess(processId).status), uint256(DAVINCITypes.ProcessStatus.ENDED));

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.PAUSED);

        // PAUSED -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromEnded() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        // Set initial state to ENDED
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);
        // ENDED -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.RESULTS);

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        vm.warp(block.timestamp + 1001);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);

        // ENDED -> CANCELED (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.CANCELED);

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);

        // ENDED -> READY (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.READY);

        // ENDED -> PAUSED (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.PAUSED);
    }

    function test_SetProcessStatus_FromCanceled() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        // Set initial state to CANCELED
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.CANCELED);

        // Try all transitions from CANCELED (all should fail)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.READY);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.PAUSED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromResults() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        // Set initial state to RESULTS
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);
        _warpToGraceEnd(processId);
        uint64[16] memory values;
        bytes32 root = processRegistry.getProcess(processId).latestStateRoot;
        processRegistry.setProcessResults(processId, _resultsPublics(root, values), "");
        // Try all transitions from RESULTS (all should fail)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.READY);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.PAUSED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.CANCELED);
    }

    function test_SetProcessStatus_Events() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        emit IProcessRegistry.ProcessStatusChanged(
            processId, DAVINCITypes.ProcessStatus.READY, DAVINCITypes.ProcessStatus.PAUSED
        );
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.PAUSED);

        emit IProcessRegistry.ProcessStatusChanged(
            processId, DAVINCITypes.ProcessStatus.PAUSED, DAVINCITypes.ProcessStatus.READY
        );
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.READY);

        emit IProcessRegistry.ProcessStatusChanged(
            processId, DAVINCITypes.ProcessStatus.READY, DAVINCITypes.ProcessStatus.ENDED
        );
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);
    }

    /// @dev A process starting in startIn seconds, 2000 seconds long.
    function _futureProcess(DAVINCITypes.ProcessStatus status, uint256 startIn) internal returns (bytes31) {
        DAVINCITypes.Census memory cen = DAVINCITypes.Census({
            onchainAllowAnyValidRoot: false,
            censusOrigin: DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census",
            contractAddress: address(0)
        });
        return processRegistry.newProcess(
            status,
            vm.getBlockTimestamp() + startIn,
            2000,
            10000,
            defaultBallotMode,
            cen,
            METADATA_URI,
            METADATA_HASH,
            _encKey(),
            _noDkg()
        );
    }

    /// @dev ENDED before startTime would open a grace window on an election that never ran;
    ///      CANCELED is how such a process is voided.
    function test_SetProcessStatus_RevertWhen_EndedBeforeStart() public {
        uint256 start = vm.getBlockTimestamp() + 1000;
        bytes31 ready = _futureProcess(DAVINCITypes.ProcessStatus.READY, 1000);
        bytes31 paused = _futureProcess(DAVINCITypes.ProcessStatus.PAUSED, 1000);

        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessStatus(ready, DAVINCITypes.ProcessStatus.ENDED);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessStatus(paused, DAVINCITypes.ProcessStatus.ENDED);
        vm.warp(start - 1);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessStatus(ready, DAVINCITypes.ProcessStatus.ENDED);

        DAVINCITypes.Process memory p = processRegistry.getProcess(ready);
        assertEq(uint256(p.status), uint256(DAVINCITypes.ProcessStatus.READY));
        assertEq(p.duration, 2000);
        assertEq(processRegistry.getProcessEndTime(ready), start + 2000);

        // From startTime on it ends as usual: duration 0 at the very start.
        vm.warp(start);
        processRegistry.setProcessStatus(ready, DAVINCITypes.ProcessStatus.ENDED);
        assertEq(processRegistry.getProcess(ready).duration, 0);
        assertEq(processRegistry.getProcessGraceEnd(ready), start + GRACE);
    }

    function test_SetProcessStatus_CanceledBeforeStart() public {
        bytes31 pid = _futureProcess(DAVINCITypes.ProcessStatus.READY, 1000);
        vm.expectEmit(true, true, true, true);
        emit IProcessRegistry.ProcessStatusChanged(
            pid, DAVINCITypes.ProcessStatus.READY, DAVINCITypes.ProcessStatus.CANCELED
        );
        processRegistry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);
        DAVINCITypes.Process memory p = processRegistry.getProcess(pid);
        assertEq(uint256(p.status), uint256(DAVINCITypes.ProcessStatus.CANCELED));
        assertEq(p.duration, 2000);
    }

    function test_SetProcessStatus_EndedAfterStart_NormalDuration() public {
        // Create a process that starts immediately
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        // Warp time forward 500 seconds (half the duration)
        vm.warp(block.timestamp + 500);

        // Set status to ENDED after process has started
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);

        // Verify duration is calculated correctly (500 seconds elapsed)
        DAVINCITypes.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint256(process.status), uint256(DAVINCITypes.ProcessStatus.ENDED));
        assertEq(process.duration, 500, "Duration should be time elapsed since start");
    }

    function test_SetProcessStatus_EndedExactlyAtStartTime() public {
        // Create a process with start time equal to current time
        DAVINCITypes.Census memory cen = DAVINCITypes.Census({
            onchainAllowAnyValidRoot: false,
            censusOrigin: DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census",
            contractAddress: address(0)
        });

        DAVINCITypes.EncryptionKey memory key = _encKey();

        bytes31 processId = processRegistry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            block.timestamp, // Start now
            1000,
            10000,
            defaultBallotMode,
            cen,
            METADATA_URI,
            METADATA_HASH,
            key,
            _noDkg()
        );

        // End process at exact start time (same block)
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.ENDED);

        // Duration should be 0 (no time elapsed)
        DAVINCITypes.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.duration, 0, "Duration should be 0 when ended at start time");
    }

    // ========== Ballot Mode Tests ==========

    function test_ValidateBallotMode_ValidCases() public {
        // Test case 1: Basic valid ballot mode
        DAVINCITypes.BallotMode memory validBallotMode1 = DAVINCITypes.BallotMode({
            uniqueValues: false,
            numFields: 1,
            groupSize: 0,
            costExponent: 1,
            maxValue: 10,
            minValue: 0,
            maxValueSum: 100,
            minValueSum: 50
        });
        bytes31 processId1 =
            createTestProcess(validBallotMode1, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        assertTrue(processId1 != bytes31(0));

        // Test case 2: Valid ballot with zero maxValueSum
        DAVINCITypes.BallotMode memory validBallotMode2 = DAVINCITypes.BallotMode({
            uniqueValues: true,
            numFields: 5,
            groupSize: 0,
            costExponent: 2,
            maxValue: 100,
            minValue: 1,
            maxValueSum: 0,
            minValueSum: 0
        });
        bytes31 processId2 =
            createTestProcess(validBallotMode2, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        assertTrue(processId2 != bytes31(0));

        // Test case 3: Edge case - maxValue equals minValue
        DAVINCITypes.BallotMode memory validBallotMode3 = DAVINCITypes.BallotMode({
            uniqueValues: false,
            numFields: 8,
            groupSize: 0,
            costExponent: 1,
            maxValue: 5,
            minValue: 5,
            maxValueSum: 50,
            minValueSum: 50
        });
        bytes31 processId3 =
            createTestProcess(validBallotMode3, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        assertTrue(processId3 != bytes31(0));
    }

    function test_ValidateBallotMode_InvalidMaxCount() public {
        DAVINCITypes.BallotMode memory invalidBallotMode = DAVINCITypes.BallotMode({
            uniqueValues: false,
            numFields: 0, // Invalid: must be >= 1
            groupSize: 0,
            costExponent: 1,
            maxValue: 10,
            minValue: 0,
            maxValueSum: 100,
            minValueSum: 50
        });

        DAVINCITypes.Census memory cen = DAVINCITypes.Census({
            onchainAllowAnyValidRoot: false,
            censusOrigin: DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census",
            contractAddress: address(0)
        });

        DAVINCITypes.EncryptionKey memory key = _encKey();

        vm.expectRevert(IProcessRegistry.InvalidMaxCount.selector);
        processRegistry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            block.timestamp, // current time
            1000000,
            10000,
            invalidBallotMode,
            cen,
            METADATA_URI,
            METADATA_HASH,
            key,
            _noDkg()
        );
    }

    function test_ValidateBallotMode_InvalidValueRange() public {
        DAVINCITypes.BallotMode memory invalidBallotMode = DAVINCITypes.BallotMode({
            uniqueValues: false,
            numFields: 1,
            groupSize: 0,
            costExponent: 1,
            maxValue: 5,
            minValue: 10, // Invalid: maxValue < minValue
            maxValueSum: 100,
            minValueSum: 50
        });

        DAVINCITypes.Census memory cen = DAVINCITypes.Census({
            onchainAllowAnyValidRoot: false,
            censusOrigin: DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census",
            contractAddress: address(0)
        });

        DAVINCITypes.EncryptionKey memory key = _encKey();

        vm.expectRevert(IProcessRegistry.InvalidMaxMinValueBounds.selector);
        processRegistry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            block.timestamp, // current time
            1000000,
            10000,
            invalidBallotMode,
            cen,
            METADATA_URI,
            METADATA_HASH,
            key,
            _noDkg()
        );
    }

    function test_NewProcess_RevertsWhenMaxPossibleResultExceedsCap() public {
        DAVINCITypes.BallotMode memory oversizedBallotMode = DAVINCITypes.BallotMode({
            uniqueValues: false,
            numFields: 1,
            groupSize: 0,
            costExponent: 1,
            maxValue: 16,
            minValue: 0,
            maxValueSum: 16,
            minValueSum: 0
        });

        DAVINCITypes.Census memory cen = DAVINCITypes.Census({
            onchainAllowAnyValidRoot: false,
            censusOrigin: DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census",
            contractAddress: address(0)
        });

        DAVINCITypes.EncryptionKey memory key = _encKey();

        vm.expectRevert(IProcessRegistry.MaxPossibleResultCapExceeded.selector);
        processRegistry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            block.timestamp,
            1000000,
            62_500_000_001,
            oversizedBallotMode,
            cen,
            METADATA_URI,
            METADATA_HASH,
            key,
            _noDkg()
        );
    }

    function test_SetProcessMaxVoters_RevertsWhenMaxPossibleResultExceedsCap() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        vm.expectRevert(IProcessRegistry.MaxPossibleResultCapExceeded.selector);
        processRegistry.setProcessMaxVoters(processId, 62_500_000_001);
    }

    // ========== Process Duration Tests ==========

    function test_SetProcessDuration_Success() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        uint256 newDuration = 2000000;

        emit IProcessRegistry.ProcessDurationChanged(processId, newDuration);
        processRegistry.setProcessDuration(processId, newDuration);

        DAVINCITypes.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.duration, newDuration);
    }

    function test_SetProcessDuration_NonExistentProcess() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.setProcessDuration(bytes31(0), 1000000);
    }

    function test_SetProcessDuration_NotAdmin() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        processRegistry.setProcessDuration(processId, 2000000);
        vm.stopPrank();
    }

    function test_SetProcessDuration_InvalidStatus_Canceled() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        // Set process to CANCELED
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.CANCELED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessDuration(processId, 2000000);
    }

    function test_SetProcessDuration_ValidStatus_Paused() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        uint256 newDuration = 2000000;

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, DAVINCITypes.ProcessStatus.PAUSED);

        emit IProcessRegistry.ProcessDurationChanged(processId, newDuration);
        processRegistry.setProcessDuration(processId, newDuration);

        DAVINCITypes.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.duration, newDuration);
    }

    function test_SetProcessDuration_InvalidDuration_Zero() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        processRegistry.setProcessDuration(processId, 0);
    }

    function test_SetProcessDuration_InvalidDuration_PastEndTime() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        // Try to set a duration that would make the process end in the past
        uint256 invalidDuration = 1; // Very short duration that will definitely be in the past

        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        processRegistry.setProcessDuration(processId, invalidDuration);
    }

    /// @dev Once the end has passed the tally may already be public (results tx in the
    ///      mempool), so the election cannot be extended and reopened.
    function test_SetProcessDuration_RevertWhen_AfterEnd() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        bytes31 paused = createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        processRegistry.setProcessStatus(paused, DAVINCITypes.ProcessStatus.PAUSED);
        uint256 end = block.timestamp + 1000;

        vm.warp(end);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessDuration(processId, 2000);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessDuration(paused, 2000);

        vm.warp(end + 500);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessDuration(processId, 2000);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessDuration(paused, 2000);
        assertEq(processRegistry.getProcess(processId).duration, 1000);
        assertEq(processRegistry.getProcess(paused).duration, 1000);
    }

    function test_SetProcessDuration_ExtendJustBeforeEnd() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        vm.warp(block.timestamp + 999);
        processRegistry.setProcessDuration(processId, 2000);
        assertEq(processRegistry.getProcess(processId).duration, 2000);
    }

    function test_SetProcessDuration_MaxDuration() public {
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        uint256 maxDuration = type(uint256).max - block.timestamp;

        emit IProcessRegistry.ProcessDurationChanged(processId, maxDuration);
        processRegistry.setProcessDuration(processId, maxDuration);

        DAVINCITypes.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.duration, maxDuration);
    }

    // ========== Process Getters Tests ==========
    function test_GetNextProcessId_Basic() public {
        // Get the next process ID for the organization
        vm.startPrank(ORGANIZER);
        bytes32 nextProcessId = processRegistry.getNextProcessId(ORGANIZER);
        uint64 currentNonce = processRegistry.processNonce(ORGANIZER);
        assertEq(currentNonce, uint64(0));
        // Create a new process
        bytes31 processId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);

        // Verify that the next process ID matches the created process ID
        assertEq(nextProcessId, processId);
        uint64 nextCurrentNonce = processRegistry.processNonce(ORGANIZER);
        assertEq(nextCurrentNonce, uint64(1));

        // Create another process
        bytes32 otherNextProcessId = processRegistry.getNextProcessId(ORGANIZER);
        bytes32 otherProcessId =
            createTestProcess(defaultBallotMode, DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1);
        assertEq(otherNextProcessId, otherProcessId);
        vm.stopPrank();
    }

    struct CensusOriginTestCase {
        DAVINCITypes.CensusOrigin censusOrigin;
        bytes32 censusRoot;
        string censusURI;
        address contractAddress;
        bytes4 revertData;
    }

    function test_newProcess_CensusOrigin() public {
        // 6 test cases
        CensusOriginTestCase[] memory testCases = new CensusOriginTestCase[](6);

        testCases[0] = CensusOriginTestCase({
            censusOrigin: DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: bytes32(0),
            censusURI: "https://example.com/census",
            contractAddress: address(0),
            revertData: IProcessRegistry.InvalidCensusRoot.selector
        });

        testCases[1] = CensusOriginTestCase({
            censusOrigin: DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: CENSUS_ROOT,
            contractAddress: address(0),
            censusURI: "",
            revertData: IProcessRegistry.InvalidCensusURI.selector
        });

        testCases[2] = CensusOriginTestCase({
            censusOrigin: DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: CENSUS_ROOT,
            censusURI: "https://example.com/census",
            contractAddress: address(0),
            revertData: bytes4(0)
        });

        testCases[3] = CensusOriginTestCase({
            censusOrigin: DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1,
            censusRoot: bytes32(0),
            censusURI: "https://example.com/census",
            contractAddress: address(0),
            revertData: IProcessRegistry.InvalidCensusRoot.selector
        });

        testCases[4] = CensusOriginTestCase({
            censusOrigin: DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1,
            censusRoot: CSP_ROOT,
            censusURI: "",
            contractAddress: address(0),
            revertData: IProcessRegistry.InvalidCensusURI.selector
        });

        testCases[5] = CensusOriginTestCase({
            censusOrigin: DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1,
            censusRoot: CSP_ROOT,
            censusURI: "https://example.com/census",
            contractAddress: address(0),
            revertData: bytes4(0)
        });

        // Iterate over test cases
        for (uint256 i = 0; i < testCases.length; i++) {
            CensusOriginTestCase memory tc = testCases[i];

            DAVINCITypes.Census memory cen = DAVINCITypes.Census({
                onchainAllowAnyValidRoot: false,
                censusOrigin: tc.censusOrigin,
                censusRoot: tc.censusRoot,
                censusURI: tc.censusURI,
                contractAddress: tc.contractAddress
            });

            DAVINCITypes.EncryptionKey memory key = _encKey();

            if (tc.revertData != bytes4(0)) {
                vm.expectRevert(tc.revertData);
            }

            processRegistry.newProcess(
                DAVINCITypes.ProcessStatus.READY,
                block.timestamp,
                1000,
                10000,
                defaultBallotMode,
                cen,
                METADATA_URI,
                METADATA_HASH,
                key,
                _noDkg()
            );
        }
    }
}
