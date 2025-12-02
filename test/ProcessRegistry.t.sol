// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { TestHelpers } from "test/TestHelpers.t.sol";
import { ProcessRegistry } from "../src/ProcessRegistry.sol";
import { OrganizationRegistry } from "../src/OrganizationRegistry.sol";
import { ProcessIdLib } from "../src/libraries/ProcessIdLib.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";
import { IProcessRegistry } from "../src/interfaces/IProcessRegistry.sol";
import { BlobsLib } from "../src/libraries/BlobsLib.sol";

/**
 * @title ProcessRegistryMock
 * @notice Test contract that allows overriding blob hash behavior for testing
 * @dev This demonstrates how to override the _blobHash function from Blobhashable
 */
contract ProcessRegistryMock is ProcessRegistry, TestHelpers {
    mapping(bytes32 => bool) public availableBlobs;

    error BlobNotFoundInTx(bytes32 versionedHash);

    constructor(
        uint32 _chainID,
        address _stVerifier,
        address _rVerifier,
        bool _blobsDA
    ) ProcessRegistry(_chainID, _stVerifier, _rVerifier, _blobsDA) {}

    // there's no way to verify blob data availability in tests
    function _verifyBlobDataIsAvailable(bytes32 versionedHash) internal view override {
        if (!availableBlobs[versionedHash]) revert BlobNotFoundInTx(versionedHash);
    }

    function setMockBlobDataAvailable(bytes32 versionedHash, bool available) external {
        availableBlobs[versionedHash] = available;
    }
}

contract ProcessRegistryTest is Test, TestHelpers {
    ProcessRegistryMock public processRegistry;
    OrganizationRegistry public organizationRegistry;
    StateTransitionVerifierGroth16 public stv;
    ResultsVerifierGroth16 public rv;

    IProcessRegistry.BallotMode public defaultBallotMode =
        IProcessRegistry.BallotMode({
            costFromWeight: false,
            uniqueValues: false,
            numFields: 5,
            costExponent: 2,
            maxValue: 16,
            minValue: 0,
            maxValueSum: 1280,
            minValueSum: 5
        });

    function setUp() public {
        organizationRegistry = new OrganizationRegistry();
        stv = new StateTransitionVerifierGroth16();
        rv = new ResultsVerifierGroth16();
        processRegistry = new ProcessRegistryMock(11155111, address(stv), address(rv), true);

        createTestOrganization();
    }

    function createTestOrganization() internal {
        address[] memory admins = new address[](1);
        admins[0] = ORGANIZATION_ADDRESS;
        organizationRegistry.createOrganization("testOrganization", "https://example.com/metadata", admins);
    }

    function createTestProcess(
        IProcessRegistry.BallotMode memory ballotMode,
        uint256 initStateRoot,
        IProcessRegistry.CensusOrigin censusOrigin
    ) internal returns (bytes32) {
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: censusOrigin,
            censusRoot: bytes32(CENSUS_ROOT),
            censusURI: "https://example.com/census"
        });

        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({
            x: uint256(keccak256(abi.encodePacked(block.timestamp, "x"))),
            y: uint256(keccak256(abi.encodePacked(block.timestamp, "y")))
        });

        bytes32 processId = processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp, // current time
            1000,
            ballotMode,
            cen,
            "https://example.com/metadata/",
            key,
            initStateRoot
        );

        return processId;
    }

    // ========== Process Status Tests ==========

    function test_SetProcessStatus_NonExistentProcess() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.setProcessStatus(bytes32(0), IProcessRegistry.ProcessStatus.ENDED);

        vm.expectRevert(IProcessRegistry.ProcessNotFound.selector);

        bytes32 h = keccak256(abi.encodePacked(uint32(11155111), address(processRegistry)));
        uint32 prefix = uint32(uint256(h));
        bytes32 invalidProcessId = ProcessIdLib.computeProcessId(
            prefix,
            address(0x1234567890123456789012345678901234567890),
            1
        );
        processRegistry.setProcessStatus(invalidProcessId, IProcessRegistry.ProcessStatus.ENDED);
    }

    function test_SetProcessStatus_UnknownProcessIdPrefix() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.setProcessStatus(bytes32(0), IProcessRegistry.ProcessStatus.ENDED);

        vm.expectRevert(IProcessRegistry.UnknownProcessIdPrefix.selector);
        processRegistry.setProcessStatus(
            bytes32(0x1000000000000000000000000000000000000000000000000000000000000001),
            IProcessRegistry.ProcessStatus.ENDED
        );
    }

    function test_SetProcessStatus_NotAdmin() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        vm.stopPrank();
    }

    function test_SetProcessStatus_SameStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);
    }

    function test_SetProcessStatus_FromReady() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        // READY -> PAUSED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.PAUSED));

        // Reset to READY
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);

        // READY -> CANCELED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.CANCELED));

        // Reset process for next test
        processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // READY -> ENDED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.ENDED));

        // Reset process for next test
        processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // READY -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromPaused() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set initial state to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // PAUSED -> READY (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.READY));

        // Reset to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // PAUSED -> CANCELED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.CANCELED));

        // Reset process for next test
        processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // PAUSED -> ENDED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.ENDED));

        // Reset process for next test
        processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // PAUSED -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromEnded() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set initial state to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        // ENDED -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);

        // Reset process for next test
        processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        vm.warp(block.timestamp + 1001);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // ENDED -> CANCELED (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);

        // Reset process for next test
        processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // ENDED -> READY (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);

        // ENDED -> PAUSED (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);
    }

    function test_SetProcessStatus_FromCanceled() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set initial state to CANCELED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);

        // Try all transitions from CANCELED (all should fail)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromResults() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set initial state to RESULTS
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
        // Try all transitions from RESULTS (all should fail)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);
    }

    function test_SetProcessStatus_Events() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        emit IProcessRegistry.ProcessStatusChanged(
            processId,
            IProcessRegistry.ProcessStatus.READY,
            IProcessRegistry.ProcessStatus.PAUSED
        );
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        emit IProcessRegistry.ProcessStatusChanged(
            processId,
            IProcessRegistry.ProcessStatus.PAUSED,
            IProcessRegistry.ProcessStatus.READY
        );
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);

        emit IProcessRegistry.ProcessStatusChanged(
            processId,
            IProcessRegistry.ProcessStatus.READY,
            IProcessRegistry.ProcessStatus.ENDED
        );
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
    }

    function test_SetProcessStatus_EndedBeforeStart_FromReady() public {
        // Create a process with start time in the future
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census"
        });

        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({
            x: uint256(keccak256(abi.encodePacked(block.timestamp, "x"))),
            y: uint256(keccak256(abi.encodePacked(block.timestamp, "y")))
        });

        uint256 futureStartTime = block.timestamp + 1000; // Start in 1000 seconds

        bytes32 processId = processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            futureStartTime,
            2000, // Duration of 2000 seconds
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            ROOT_HASH_BEFORE
        );

        // Verify initial state
        IProcessRegistry.Process memory processBefore = processRegistry.getProcess(processId);
        assertEq(processBefore.startTime, futureStartTime);
        assertEq(processBefore.duration, 2000);
        assertEq(uint(processBefore.status), uint(IProcessRegistry.ProcessStatus.READY));

        // Set status to ENDED before start time
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Verify the process is ENDED and duration is 0
        IProcessRegistry.Process memory processAfter = processRegistry.getProcess(processId);
        assertEq(uint(processAfter.status), uint(IProcessRegistry.ProcessStatus.ENDED));
        assertEq(processAfter.duration, 0, "Duration should be 0 when ended before start");
        assertEq(processAfter.startTime, futureStartTime, "Start time should remain unchanged");
    }

    function test_SetProcessStatus_EndedBeforeStart_FromPaused() public {
        // Create a process with start time in the future
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census"
        });

        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({
            x: uint256(keccak256(abi.encodePacked(block.timestamp, "x"))),
            y: uint256(keccak256(abi.encodePacked(block.timestamp, "y")))
        });

        uint256 futureStartTime = block.timestamp + 500;

        bytes32 processId = processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.PAUSED, // Start in PAUSED state
            futureStartTime,
            1000,
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            ROOT_HASH_BEFORE
        );

        // Verify initial state
        IProcessRegistry.Process memory processBefore = processRegistry.getProcess(processId);
        assertEq(uint(processBefore.status), uint(IProcessRegistry.ProcessStatus.PAUSED));

        // Set status to ENDED before start time (from PAUSED)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Verify the process is ENDED and duration is 0
        IProcessRegistry.Process memory processAfter = processRegistry.getProcess(processId);
        assertEq(uint(processAfter.status), uint(IProcessRegistry.ProcessStatus.ENDED));
        assertEq(processAfter.duration, 0, "Duration should be 0 when ended before start from PAUSED");
    }

    function test_SetProcessStatus_EndedBeforeStart_EventEmitted() public {
        // Create a process with start time in the future
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census"
        });

        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({
            x: uint256(keccak256(abi.encodePacked(block.timestamp, "x"))),
            y: uint256(keccak256(abi.encodePacked(block.timestamp, "y")))
        });

        uint256 futureStartTime = block.timestamp + 1000;

        bytes32 processId = processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            futureStartTime,
            2000,
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            ROOT_HASH_BEFORE
        );

        // Expect both status change and duration change events
        vm.expectEmit(true, true, true, true);
        emit IProcessRegistry.ProcessDurationChanged(processId, 0);

        vm.expectEmit(true, true, true, true);
        emit IProcessRegistry.ProcessStatusChanged(
            processId,
            IProcessRegistry.ProcessStatus.READY,
            IProcessRegistry.ProcessStatus.ENDED
        );

        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
    }

    function test_SetProcessStatus_EndedAfterStart_NormalDuration() public {
        // Create a process that starts immediately
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Warp time forward 500 seconds (half the duration)
        vm.warp(block.timestamp + 500);

        // Set status to ENDED after process has started
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Verify duration is calculated correctly (500 seconds elapsed)
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.ENDED));
        assertEq(process.duration, 500, "Duration should be time elapsed since start");
    }

    function test_SetProcessStatus_EndedExactlyAtStartTime() public {
        // Create a process with start time equal to current time
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census"
        });

        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({
            x: uint256(keccak256(abi.encodePacked(block.timestamp, "x"))),
            y: uint256(keccak256(abi.encodePacked(block.timestamp, "y")))
        });

        bytes32 processId = processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp, // Start now
            1000,
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            ROOT_HASH_BEFORE
        );

        // End process at exact start time (same block)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Duration should be 0 (no time elapsed)
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.duration, 0, "Duration should be 0 when ended at start time");
    }

    function test_GetProcessEndTime_WhenEndedBeforeStart() public {
        // Create a process with start time in the future
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census"
        });

        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({
            x: uint256(keccak256(abi.encodePacked(block.timestamp, "x"))),
            y: uint256(keccak256(abi.encodePacked(block.timestamp, "y")))
        });

        uint256 futureStartTime = block.timestamp + 1000;

        bytes32 processId = processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            futureStartTime,
            2000,
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            ROOT_HASH_BEFORE
        );

        // End process before start time
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Get process end time - should equal start time (since duration is 0)
        uint256 endTime = processRegistry.getProcessEndTime(processId);
        assertEq(endTime, futureStartTime, "End time should equal start time when ended before start");
    }

    function test_SetProcessStatus_EndedBeforeStart_MultipleTimes() public {
        // Create process 1 - ended immediately
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census"
        });

        IProcessRegistry.EncryptionKey memory key1 = IProcessRegistry.EncryptionKey({
            x: uint256(keccak256(abi.encodePacked(block.timestamp, "x1"))),
            y: uint256(keccak256(abi.encodePacked(block.timestamp, "y1")))
        });

        bytes32 processId1 = processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp + 1000,
            2000,
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key1,
            ROOT_HASH_BEFORE
        );

        processRegistry.setProcessStatus(processId1, IProcessRegistry.ProcessStatus.ENDED);

        // Create process 2 - also ended before start
        IProcessRegistry.EncryptionKey memory key2 = IProcessRegistry.EncryptionKey({
            x: uint256(keccak256(abi.encodePacked(block.timestamp, "x2"))),
            y: uint256(keccak256(abi.encodePacked(block.timestamp, "y2")))
        });

        bytes32 processId2 = processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp + 500,
            1000,
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key2,
            ROOT_HASH_BEFORE
        );

        processRegistry.setProcessStatus(processId2, IProcessRegistry.ProcessStatus.ENDED);

        // Verify both processes
        IProcessRegistry.Process memory process1 = processRegistry.getProcess(processId1);
        IProcessRegistry.Process memory process2 = processRegistry.getProcess(processId2);

        assertEq(process1.duration, 0, "Process 1 duration should be 0");
        assertEq(process2.duration, 0, "Process 2 duration should be 0");
        assertEq(uint(process1.status), uint(IProcessRegistry.ProcessStatus.ENDED));
        assertEq(uint(process2.status), uint(IProcessRegistry.ProcessStatus.ENDED));
    }

    // ========== Process Census Tests ==========

    function test_SetProcessCensus_Success() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        processRegistry.setProcessCensus(processId, newCensus);

        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.census.censusRoot, newCensus.censusRoot);
        assertEq(process.census.censusURI, newCensus.censusURI);
    }

    function test_SetProcessCensus_NonExistentProcess() public {
        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.setProcessCensus(bytes32(0), newCensus);
    }

    function test_SetProcessCensus_NotAdmin() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        processRegistry.setProcessCensus(processId, newCensus);
        vm.stopPrank();
    }

    function test_SetProcessCensus_InvalidCensus_EmptyURI() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: ""
        });

        vm.expectRevert(IProcessRegistry.InvalidCensusURI.selector);
        processRegistry.setProcessCensus(processId, newCensus);
    }

    function test_SetProcessCensus_InvalidCensus_ZeroCensusRoot() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0,
            censusURI: "https://example.com/new-census"
        });

        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        processRegistry.setProcessCensus(processId, newCensus);
    }

    function test_SetProcessCensus_InvalidStatus_Canceled() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to CANCELED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessCensus(processId, newCensus);
    }

    function test_SetProcessCensus_ValidStatus_Paused() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        // Should succeed
        processRegistry.setProcessCensus(processId, newCensus);

        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.census.censusRoot, newCensus.censusRoot);
        assertEq(process.census.censusURI, newCensus.censusURI);
    }

    // ========== Ballot Mode Tests ==========

    function test_ValidateBallotMode_ValidCases() public {
        // Test case 1: Basic valid ballot mode
        IProcessRegistry.BallotMode memory validBallotMode1 = IProcessRegistry.BallotMode({
            costFromWeight: false,
            uniqueValues: false,
            numFields: 1,
            costExponent: 1,
            maxValue: 10,
            minValue: 0,
            maxValueSum: 100,
            minValueSum: 50
        });
        bytes32 processId1 = createTestProcess(
            validBallotMode1,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        assertTrue(processId1 != bytes32(0));

        // Test case 2: Valid ballot with costFromWeight true
        IProcessRegistry.BallotMode memory validBallotMode2 = IProcessRegistry.BallotMode({
            costFromWeight: true,
            uniqueValues: true,
            numFields: 5,
            costExponent: 2,
            maxValue: 100,
            minValue: 1,
            maxValueSum: 0, // This is valid when costFromWeight is true
            minValueSum: 0
        });
        bytes32 processId2 = createTestProcess(
            validBallotMode2,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        assertTrue(processId2 != bytes32(0));

        // Test case 3: Edge case - maxValue equals minValue
        IProcessRegistry.BallotMode memory validBallotMode3 = IProcessRegistry.BallotMode({
            costFromWeight: false,
            uniqueValues: false,
            numFields: 8,
            costExponent: 1,
            maxValue: 5,
            minValue: 5,
            maxValueSum: 50,
            minValueSum: 50
        });
        bytes32 processId3 = createTestProcess(
            validBallotMode3,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        assertTrue(processId3 != bytes32(0));
    }

    function test_ValidateBallotMode_InvalidMaxCount() public {
        IProcessRegistry.BallotMode memory invalidBallotMode = IProcessRegistry.BallotMode({
            costFromWeight: false,
            uniqueValues: false,
            numFields: 0, // Invalid: must be >= 1
            costExponent: 1,
            maxValue: 10,
            minValue: 0,
            maxValueSum: 100,
            minValueSum: 50
        });

        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census"
        });

        // Generate random but valid EC point coordinates
        uint256 keyX = uint256(keccak256(abi.encodePacked(block.timestamp, "x")));
        uint256 keyY = uint256(keccak256(abi.encodePacked(block.timestamp, "y")));
        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({ x: keyX, y: keyY });

        vm.expectRevert(IProcessRegistry.InvalidMaxCount.selector);
        processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp, // current time
            1000000,
            invalidBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            ROOT_HASH_BEFORE
        );
    }

    function test_ValidateBallotMode_InvalidValueRange() public {
        IProcessRegistry.BallotMode memory invalidBallotMode = IProcessRegistry.BallotMode({
            costFromWeight: false,
            uniqueValues: false,
            numFields: 1,
            costExponent: 1,
            maxValue: 5,
            minValue: 10, // Invalid: maxValue < minValue
            maxValueSum: 100,
            minValueSum: 50
        });

        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census"
        });

        // Generate random but valid EC point coordinates
        uint256 keyX = uint256(keccak256(abi.encodePacked(block.timestamp, "x")));
        uint256 keyY = uint256(keccak256(abi.encodePacked(block.timestamp, "y")));
        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({ x: keyX, y: keyY });

        vm.expectRevert(IProcessRegistry.InvalidMaxMinValueBounds.selector);
        processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp, // current time
            1000000,
            invalidBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            ROOT_HASH_BEFORE
        );
    }

    // ========== Process Duration Tests ==========

    function test_SetProcessDuration_Success() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        uint256 newDuration = 2000000;

        emit IProcessRegistry.ProcessDurationChanged(processId, newDuration);
        processRegistry.setProcessDuration(processId, newDuration);

        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.duration, newDuration);
    }

    function test_SetProcessDuration_NonExistentProcess() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.setProcessDuration(bytes32(0), 1000000);
    }

    function test_SetProcessDuration_NotAdmin() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        processRegistry.setProcessDuration(processId, 2000000);
        vm.stopPrank();
    }

    function test_SetProcessDuration_InvalidStatus_Canceled() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to CANCELED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessDuration(processId, 2000000);
    }

    function test_SetProcessDuration_ValidStatus_Paused() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        uint256 newDuration = 2000000;

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        emit IProcessRegistry.ProcessDurationChanged(processId, newDuration);
        processRegistry.setProcessDuration(processId, newDuration);

        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.duration, newDuration);
    }

    function test_SetProcessDuration_InvalidDuration_Zero() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        processRegistry.setProcessDuration(processId, 0);
    }

    function test_SetProcessDuration_InvalidDuration_PastEndTime() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Try to set a duration that would make the process end in the past
        uint256 invalidDuration = 1; // Very short duration that will definitely be in the past

        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        processRegistry.setProcessDuration(processId, invalidDuration);
    }

    function test_SetProcessDuration_MaxDuration() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        uint256 maxDuration = type(uint256).max - block.timestamp;

        emit IProcessRegistry.ProcessDurationChanged(processId, maxDuration);
        processRegistry.setProcessDuration(processId, maxDuration);

        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.duration, maxDuration);
    }

    // ========== State Transition Tests ==========

    function test_SubmitStateTransition_Success() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);

        // Verify initial state
        assertEq(process.latestStateRoot, ROOT_HASH_BEFORE);
        assertEq(process.votersCount, 0);
        assertEq(process.overwrittenVotesCount, 0);

        processRegistry.setMockBlobDataAvailable(BLOB_VERSIONEDHASH, true);

        vm.mockCall(KZG_PRECOMPILE, "", abi.encode(FIELD_ELEMENTS_PER_BLOB, BLS_MODULUS));

        // Submit state transition
        emit IProcessRegistry.ProcessStateRootUpdated(processId, address(this), ROOT_HASH_BEFORE);
        processRegistry.submitStateTransition(processId, STATETRANSITION_ABI_PROOF, stateTransitionInputs());

        // Verify state after transition
        process = processRegistry.getProcess(processId);
        assertEq(process.latestStateRoot, ROOT_HASH_AFTER);
        assertEq(process.votersCount, VOTERS_COUNT - OVERWRITTEN_VOTES_COUNT);
        assertEq(process.overwrittenVotesCount, OVERWRITTEN_VOTES_COUNT);
    }

    function test_SubmitStateTransition_NonExistentProcess() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.submitStateTransition(bytes32(0), STATETRANSITION_ABI_PROOF, STATETRANSITION_ABI_INPUTS);
    }

    function test_SubmitStateTransition_InvalidStatus_Paused() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, STATETRANSITION_ABI_PROOF, STATETRANSITION_ABI_INPUTS);
    }

    function test_SubmitStateTransition_InvalidStatus_Ended() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, STATETRANSITION_ABI_PROOF, STATETRANSITION_ABI_INPUTS);
    }

    function test_SubmitStateTransition_InvalidStatus_Canceled() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to CANCELED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, STATETRANSITION_ABI_PROOF, STATETRANSITION_ABI_INPUTS);
    }

    function test_SubmitStateTransition_InvalidStatus_Results() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED then RESULTS
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, STATETRANSITION_ABI_PROOF, STATETRANSITION_ABI_INPUTS);
    }

    function test_SubmitStateTransition_ProofInvalid() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        processRegistry.setMockBlobDataAvailable(BLOB_VERSIONEDHASH, true);

        vm.expectRevert(IProcessRegistry.ProofInvalid.selector);
        processRegistry.submitStateTransition(processId, STATETRANSITION_ABI_PROOF_INVALID, stateTransitionInputs());
    }

    function test_SubmitStateTransition_InvalidTimeBounds() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Fast forward time beyond process duration
        vm.warp(block.timestamp + 1000001);

        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.submitStateTransition(processId, STATETRANSITION_ABI_PROOF, STATETRANSITION_ABI_INPUTS);
    }

    // ========== Process Results Tests ==========

    function test_SetProcessResults_Success() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Expect events to be emitted
        emit IProcessRegistry.ProcessStatusChanged(
            processId,
            IProcessRegistry.ProcessStatus.ENDED,
            IProcessRegistry.ProcessStatus.RESULTS
        );
        uint256[] memory mresults = new uint256[](FINAL_RESULTS.length);
        for (uint256 i = 0; i < FINAL_RESULTS.length; i++) {
            mresults[i] = FINAL_RESULTS[i];
        }
        emit IProcessRegistry.ProcessResultsSet(processId, address(this), mresults);

        // Set results
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify process state
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
        assertEq(process.result.length, FINAL_RESULTS.length);
        for (uint i = 0; i < FINAL_RESULTS.length; i++) {
            assertEq(process.result[i], FINAL_RESULTS[i]);
        }
    }

    function test_SetProcessResults_NonExistentProcess() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.setProcessResults(bytes32(0), RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_NotEndedStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Try to set results when process is in READY state (time not expired)
        // Should fail with InvalidTimeBounds because time check happens after status check
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Set to PAUSED and try (time not expired)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Set to CANCELED and try - this should failRESULTS_ABI_PROOFatus
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_ProcessNotEndedByTime() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Try to set results before process duration has passed (should work)
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify process state
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
        assertEq(process.result.length, FINAL_RESULTS.length);
        for (uint i = 0; i < FINAL_RESULTS.length; i++) {
            assertEq(process.result[i], FINAL_RESULTS[i]);
        }
    }

    function test_SetProcessResults_InvalidStateRoot() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        ); // Using different state root

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Try to set results with mismatched state root
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_CannotSetTwice() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Set results first time
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Try to set results again
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_InvalidProof() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Try to set results with invalid proof
        bytes memory invalidProof = hex"deadbeef";
        vm.expectRevert(); // The ZK verifier should revert with invalid proof
        processRegistry.setProcessResults(processId, invalidProof, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_InvalidInput() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Try to set results with invalid input
        bytes memory invalidInput = hex"deadbeef";
        vm.expectRevert(); // Should revert on abi.decode
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, invalidInput);
    }

    function test_SetProcessResults_EmptyResults() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Create empty results input
        bytes memory emptyInput = abi.encode([ROOT_HASH_AFTER]); // Only state root, no results

        vm.expectRevert(); // Should revert as input must have at least one result
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, emptyInput);
    }

    function test_SetProcessResults_UnknownProcessIdPrefix() public {
        // Create a process ID with invalid prefix
        bytes32 invalidProcessId = bytes32(0x1000000000000000000000000000000000000000000000000000000000000001);

        vm.expectRevert(IProcessRegistry.UnknownProcessIdPrefix.selector);
        processRegistry.setProcessResults(invalidProcessId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_TimeExpired_ReadyStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Process is in READY status, fast forward time beyond duration
        vm.warp(block.timestamp + 1001);

        // Should succeed because time has expired
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify process state
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
        assertEq(process.result.length, FINAL_RESULTS.length);
        for (uint i = 0; i < FINAL_RESULTS.length; i++) {
            assertEq(process.result[i], FINAL_RESULTS[i]);
        }
    }

    function test_SetProcessResults_TimeExpired_PausedStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // Fast forward time beyond duration
        vm.warp(block.timestamp + 1001);

        // Should succeed because time has expired
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify process state
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
        assertEq(process.result.length, FINAL_RESULTS.length);
    }

    function test_SetProcessResults_TimeNotExpired_ReadyStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Process is in READY status and time has NOT expired
        // Should fail because neither ENDED status nor time expired
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_TimeNotExpired_PausedStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // Time has NOT expired
        // Should fail because neither ENDED status nor time expired
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_ExactlyAtExpiration() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Get process to check start time and duration
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        uint256 expirationTime = process.startTime + process.duration;

        // Warp to exact expiration time
        vm.warp(expirationTime);

        // Should succeed at exact expiration boundary
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify results were set
        process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
    }

    function test_SetProcessResults_OneSecondBeforeExpiration() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Get process to check start time and duration
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        uint256 expirationTime = process.startTime + process.duration;

        // Warp to 1 second before expiration
        vm.warp(expirationTime - 1);

        // Should fail because time has not expired
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_OneSecondAfterExpiration() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Get process to check start time and duration
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        uint256 expirationTime = process.startTime + process.duration;

        // Warp to 1 second after expiration
        vm.warp(expirationTime + 1);

        // Should succeed
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify results were set
        process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
    }

    function test_SetProcessResults_ByNonAdmin() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Submit results from different address (not the process admin)
        address randomUser = address(0xbeef);
        vm.prank(randomUser);
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
        vm.stopPrank();

        // Verify results were set successfully (permissionless)
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
    }

    function test_SetProcessResults_EventsEmitted() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Prepare expected results array
        uint256[] memory expectedResults = new uint256[](FINAL_RESULTS.length);
        for (uint256 i = 0; i < FINAL_RESULTS.length; i++) {
            expectedResults[i] = FINAL_RESULTS[i];
        }

        // Expect both events
        vm.expectEmit(true, true, true, true);
        emit IProcessRegistry.ProcessStatusChanged(
            processId,
            IProcessRegistry.ProcessStatus.ENDED,
            IProcessRegistry.ProcessStatus.RESULTS
        );

        vm.expectEmit(true, true, true, true);
        emit IProcessRegistry.ProcessResultsSet(processId, address(this), expectedResults);

        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_OldStatusCaptured_FromReady() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Fast forward time to expiration
        vm.warp(block.timestamp + 1001);

        // Expect event with READY as old status
        vm.expectEmit(true, true, true, true);
        emit IProcessRegistry.ProcessStatusChanged(
            processId,
            IProcessRegistry.ProcessStatus.READY,
            IProcessRegistry.ProcessStatus.RESULTS
        );

        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_OldStatusCaptured_FromPaused() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // Fast forward time to expiration
        vm.warp(block.timestamp + 1001);

        // Expect event with PAUSED as old status
        vm.expectEmit(true, true, true, true);
        emit IProcessRegistry.ProcessStatusChanged(
            processId,
            IProcessRegistry.ProcessStatus.PAUSED,
            IProcessRegistry.ProcessStatus.RESULTS
        );

        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_SetProcessResults_ShortDurationProcess() public {
        // Create a process with very short duration (1 second)
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census"
        });

        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({
            x: uint256(keccak256(abi.encodePacked(block.timestamp, "x"))),
            y: uint256(keccak256(abi.encodePacked(block.timestamp, "y")))
        });

        bytes32 processId = processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp,
            1, // 1 second duration
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            ROOT_HASH_AFTER
        );

        // Wait for 1 second
        vm.warp(block.timestamp + 1);

        // Should succeed
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify results
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
    }

    function test_SetProcessResults_SingleResultValue() public {
        // The rei constant already has valid proof for 8 results, but we're testing decode behavior
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Set results (will extract all values after state root)
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify results array is created correctly
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
        assertEq(process.result.length, 8); // 9 inputs - 1 state root = 8 results
    }

    function test_SetProcessResults_VerifyStateRootMatch() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Verify the process has the correct state root before setting results
        IProcessRegistry.Process memory processBefore = processRegistry.getProcess(processId);
        assertEq(processBefore.latestStateRoot, ROOT_HASH_AFTER);

        // Set results
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify state root remains the same after setting results
        IProcessRegistry.Process memory processAfter = processRegistry.getProcess(processId);
        assertEq(processAfter.latestStateRoot, ROOT_HASH_AFTER);
    }

    function test_SetProcessResults_ResultsArrayLengthCorrect() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Set results
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify results array length matches input (9 inputs - 1 state root = 8 results)
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.result.length, 8);
    }

    function test_SetProcessResults_MultipleProcessesSameOrganization() public {
        // Create first process
        bytes32 processId1 = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        processRegistry.setProcessStatus(processId1, IProcessRegistry.ProcessStatus.ENDED);

        // Create second process with different state root
        bytes32 processId2 = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        processRegistry.setProcessStatus(processId2, IProcessRegistry.ProcessStatus.ENDED);

        // Set results for first process
        processRegistry.setProcessResults(processId1, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify first process has results
        IProcessRegistry.Process memory process1 = processRegistry.getProcess(processId1);
        assertEq(uint(process1.status), uint(IProcessRegistry.ProcessStatus.RESULTS));

        // Verify second process is still in ENDED state
        IProcessRegistry.Process memory process2 = processRegistry.getProcess(processId2);
        assertEq(uint(process2.status), uint(IProcessRegistry.ProcessStatus.ENDED));
    }

    function test_SetProcessResults_ProcessCountersUnaffected() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Get initial counters
        IProcessRegistry.Process memory processBefore = processRegistry.getProcess(processId);
        uint256 votersCountBefore = processBefore.votersCount;
        uint256 overwrittenVotesCountBefore = processBefore.overwrittenVotesCount;
        uint256 batchNumberBefore = processBefore.batchNumber;

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Set results
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify counters remain unchanged
        IProcessRegistry.Process memory processAfter = processRegistry.getProcess(processId);
        assertEq(processAfter.votersCount, votersCountBefore);
        assertEq(processAfter.overwrittenVotesCount, overwrittenVotesCountBefore);
        assertEq(processAfter.batchNumber, batchNumberBefore);
    }

    function test_SetProcessResults_ResultsMatchInput() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_AFTER,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Set results
        processRegistry.setProcessResults(processId, RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);

        // Verify each result value matches the expected values
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        for (uint256 i = 0; i < FINAL_RESULTS.length; i++) {
            assertEq(process.result[i], FINAL_RESULTS[i], "Result mismatch at index");
        }
    }

    // ========== Process Getters Tests ==========
    function test_GetNextProcessId_Basic() public {
        // Get the next process ID for the organization
        vm.startPrank(ORGANIZATION_ADDRESS);
        bytes32 nextProcessId = processRegistry.getNextProcessId(ORGANIZATION_ADDRESS);
        uint64 currentNonce = processRegistry.processNonce(ORGANIZATION_ADDRESS);
        assertEq(currentNonce, uint64(0));
        // Create a new process
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Verify that the next process ID matches the created process ID
        assertEq(nextProcessId, processId);
        uint64 nextCurrentNonce = processRegistry.processNonce(ORGANIZATION_ADDRESS);
        assertEq(nextCurrentNonce, uint64(1));

        // Create another process
        bytes32 otherNextProcessId = processRegistry.getNextProcessId(ORGANIZATION_ADDRESS);
        bytes32 otherProcessId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        assertEq(otherNextProcessId, otherProcessId);
        vm.stopPrank();
    }

    struct CensusOriginTestCase {
        IProcessRegistry.CensusOrigin censusOrigin;
        bytes32 censusRoot;
        string censusURI;
        bytes4 revertData;
    }

    function test_newProcess_CensusOrigin() public {
        // 6 test cases
        CensusOriginTestCase[] memory testCases = new CensusOriginTestCase[](6);

        testCases[0] = CensusOriginTestCase({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: bytes32(0),
            censusURI: "https://example.com/census",
            revertData: IProcessRegistry.InvalidCensusRoot.selector
        });

        testCases[1] = CensusOriginTestCase({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: bytes32(CENSUS_ROOT),
            censusURI: "",
            revertData: IProcessRegistry.InvalidCensusURI.selector
        });

        testCases[2] = CensusOriginTestCase({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: bytes32(CENSUS_ROOT),
            censusURI: "https://example.com/census",
            revertData: bytes4(0)
        });

        testCases[3] = CensusOriginTestCase({
            censusOrigin: IProcessRegistry.CensusOrigin.CSP_EDDSA_BN254_V1,
            censusRoot: bytes32(0),
            censusURI: "https://example.com/census",
            revertData: IProcessRegistry.InvalidCensusRoot.selector
        });

        testCases[4] = CensusOriginTestCase({
            censusOrigin: IProcessRegistry.CensusOrigin.CSP_EDDSA_BN254_V1,
            censusRoot: bytes32(CENSUS_ROOT),
            censusURI: "",
            revertData: IProcessRegistry.InvalidCensusURI.selector
        });

        testCases[5] = CensusOriginTestCase({
            censusOrigin: IProcessRegistry.CensusOrigin.CSP_EDDSA_BN254_V1,
            censusRoot: bytes32(CENSUS_ROOT),
            censusURI: "https://example.com/census",
            revertData: bytes4(0)
        });

        // Iterate over test cases
        for (uint256 i = 0; i < testCases.length; i++) {
            CensusOriginTestCase memory tc = testCases[i];

            IProcessRegistry.Census memory cen = IProcessRegistry.Census({
                censusOrigin: tc.censusOrigin,
                censusRoot: tc.censusRoot,
                censusURI: tc.censusURI
            });

            IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({
                x: uint256(keccak256(abi.encodePacked(block.timestamp, "x", i))),
                y: uint256(keccak256(abi.encodePacked(block.timestamp, "y", i)))
            });

            if (tc.revertData != bytes4(0)) {
                vm.expectRevert(tc.revertData);
            }

            processRegistry.newProcess(
                IProcessRegistry.ProcessStatus.READY,
                block.timestamp,
                1000,
                defaultBallotMode,
                cen,
                "https://example.com/metadata/",
                key,
                ROOT_HASH_BEFORE
            );
        }
    }
}
