// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { TestHelpers } from "./TestHelpers.t.sol";
import { ProcessRegistry } from "../src/ProcessRegistry.sol";
import { OrganizationRegistry } from "../src/OrganizationRegistry.sol";
import { ProcessIdLib } from "../src/libraries/ProcessIdLib.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";
import { IProcessRegistry } from "../src/interfaces/IProcessRegistry.sol";

contract ProcessRegistryTest is Test, TestHelpers {
    ProcessRegistry public processRegistry;
    OrganizationRegistry public organizationRegistry;
    StateTransitionVerifierGroth16 public stv;
    ResultsVerifierGroth16 public rv;

    bytes public constant stzkp =
        hex"00e53c8167049f28eab4d9352b59e1379e357e4a8cd7b070ddee85f4c767baa1001a7110fcdbf19e2ba7bc794c0504fd21f95ebe2a8a97bda8eb6ddf653cf226267ccd3f08fefb8b90c5621753b2f7ef403172663984c5c4bb27bbf6b971cde5220db7bd2efd0de9c3208f439aa96f52c3bda840f0fb85a3ac676e4fe1a74a001b348ca610ad440bbdfda60d0e803233a001f6ee07baee0444124e9f082e10641bf6a437676dc5eee04391e2d4a1e582c2e4d44047826a254193daedaba3233b2f88cc0117eb857fc39484b658336f8d0578270b8bd8f32bc06e4245243f8e3e08860b64057c5f0da39cae5700cbab4dc785c86a334d21ebec37c860f20ffd7a249f1c54ff9e5127c74e89e799e770f85a17cf3072f4e7c04f7bf592a1b109ad19114e6304e3adc1663569832a251ffa2bf3b5f340d388af2dc7a7790f1a8c6d173c081ceb82feda29e2c4feb7c7019d8bd56c4bea503658f5f87711aaa2612812bfca2a38897d7a1c87f7cf95d78f9bb003d61e8c2a6499d21becc216399d85";
    bytes public constant rzkp =
        hex"30171a422b25e70e0707be4fccffa7638a4c4cefc7e8c2e5625182aed3bda8d62e522daaafba69d96b0939a7229eb2bccf92696759a918ab388db28da665749904c0791c56532e7f8bf8e0dfe9ac16abfef3c476700a7c7cd137426c1c34976013e912297ba7f085d18e257891dca0960883f037509ccd93ff1a963c3d542dbd24d0ffc967592007d2e055cb3a9391802e65f8574678ecd9583045cfdc2881e22316ea8aa41936ada7d834546fdb274e92bbad1d3c640b4190f62126b6d5cceb0c31d3592f9fa60e51e1ad2e01ea70a2aef1b87923dd92cbe404cc19f9c76c8a2cf50262055abac1a0f624f5b8f958ea2c51c0b9e5820bb143840d6cbc4c0b5a22f679979ac66b13134ed83a22a5b25a052bbd72843e6450d91c4a639645514402920440fb98ae4c80851fe5f43983e759601749d7ffd2c7a408b886f45587410cee512815a3168796b2f6e8f15cb58a2a79e0426a8c6de0dcc2987727781d461d6f7c167707e8e9d855e1d024a5cf82042cc0df7fa5e6e9601c1c5e1920a716";

    bytes public constant stei =
        hex"23b39023b05670b89f104beeed738d1fafa3ca1aa879f899732ebd1baeab50eb2b9b48c4c4fcec686d3144c6c3a6ea1172251bd6a43d915596485f19a4c57f210000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000000514944e45dbe85232e2c2e6c2d68d4636c7c1e18e21f115ade515a3d73d4a638000927e5a75c2bfbd43932609df1ed16b583af2d606ce9a01e82439010c8b0e870000000000000000000000000000000000000000000000001b609d0d33d8330a000000000000000000000000000000000000000000000000e5d09a0758fdf60600000000000000000000000000000000000000000000000071376b5766880ef80000000000000000000000000000000000000000000000001ed352ee8177b369000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000030a8a4d78fe15776833f87428bc093a0f1c3684605fe64cbe791a372b2bfe9f349c92a3d9719688a0ecafeb49ed9edeebd000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030a8bc33f0adb47d3f9e3273f9f68ef96e414055d094c104b2290f47803581bd04fd33b52a7fc920985aa3af1fec58c49c00000000000000000000000000000000";
    bytes public constant rei =
        hex"2b9b48c4c4fcec686d3144c6c3a6ea1172251bd6a43d915596485f19a4c57f21000000000000000000000000000000000000000000000000000000000000001f000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002a00000000000000000000000000000000000000000000000000000000000000290000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256 public constant stInitStateRoot =
        10097873574587648384391095485391824303192352374242914629845966949633379821128;
    uint256 public constant rInitStateRoot =
        19723816014752875800279306456496578536046296876388232350336782237768121417505;

    address orgId = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;

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
        processRegistry = new ProcessRegistry(11155111, address(stv), address(rv), false); // TODO SET TRUE FOR BLOB TEST

        createTestOrganization();
    }

    function createTestOrganization() internal {
        address[] memory admins = new address[](1);
        admins[0] = orgId;
        organizationRegistry.createOrganization("testOrganization", "https://example.com/metadata", admins);
    }

    function createTestProcess(
        IProcessRegistry.BallotMode memory ballotMode,
        uint256 initStateRoot,
        IProcessRegistry.CensusOrigin censusOrigin
    ) internal returns (bytes32) {
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: censusOrigin,
            censusRoot: bytes32(TestHelpers.CENSUS_ROOT),
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
            stInitStateRoot,
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
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);
    }

    function test_SetProcessStatus_FromReady() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
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
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // READY -> ENDED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.ENDED));

        // Reset process for next test
        processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // READY -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromPaused() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
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
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // PAUSED -> ENDED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.ENDED));

        // Reset process for next test
        processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set initial state to RESULTS
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        processRegistry.setProcessResults(processId, rzkp, rei);
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
            stInitStateRoot,
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
            stInitStateRoot
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
            stInitStateRoot
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
            stInitStateRoot
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
            stInitStateRoot,
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
            stInitStateRoot
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
            stInitStateRoot
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
            stInitStateRoot
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
            stInitStateRoot
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot
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
            stInitStateRoot
        );
    }

    // ========== Process Duration Tests ==========

    function test_SetProcessDuration_Success() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
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
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        processRegistry.setProcessDuration(processId, 0);
    }

    function test_SetProcessDuration_InvalidDuration_PastEndTime() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
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
            stInitStateRoot,
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
        assertEq(process.voteCount, 0);
        assertEq(process.voteOverwriteCount, 0);

        // Submit state transition
        emit IProcessRegistry.ProcessStateRootUpdated(processId, address(this), ROOT_HASH_BEFORE);
        processRegistry.submitStateTransition(processId, stateTransitionZKProof, stateTransitionInputs());

        // Verify state after transition
        process = processRegistry.getProcess(processId);
        assertEq(process.latestStateRoot, ROOT_HASH_AFTER);
        assertEq(process.voteCount, NUM_NEW_VOTES);
        assertEq(process.voteOverwriteCount, NUM_OVERWRITES);
    }

    function test_SubmitStateTransition_NonExistentProcess() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.submitStateTransition(bytes32(0), stzkp, stei);
    }

    function test_SubmitStateTransition_InvalidStatus_Paused() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, stzkp, stei);
    }

    function test_SubmitStateTransition_InvalidStatus_Ended() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, stzkp, stei);
    }

    function test_SubmitStateTransition_InvalidStatus_Canceled() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to CANCELED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, stzkp, stei);
    }

    function test_SubmitStateTransition_InvalidStatus_Results() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED then RESULTS
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        processRegistry.setProcessResults(processId, rzkp, rei);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, stzkp, stei);
    }

    function test_SubmitStateTransition_ProofInvalid() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            ROOT_HASH_BEFORE,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        vm.expectRevert(IProcessRegistry.ProofInvalid.selector);
        processRegistry.submitStateTransition(processId, stateTransitionZKProof_Invalid, stateTransitionInputs());
    }

    function test_SubmitStateTransition_InvalidTimeBounds() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Fast forward time beyond process duration
        vm.warp(block.timestamp + 1000001);

        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.submitStateTransition(processId, stzkp, stei);
    }

    // ========== Process Results Tests ==========

    function test_SetProcessResults_Success() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
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
        uint256[] memory mresults = new uint256[](results.length);
        for (uint256 i = 0; i < results.length; i++) {
            mresults[i] = results[i];
        }
        emit IProcessRegistry.ProcessResultsSet(processId, address(this), mresults);

        // Set results
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify process state
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
        assertEq(process.result.length, results.length);
        for (uint i = 0; i < results.length; i++) {
            assertEq(process.result[i], results[i]);
        }
    }

    function test_SetProcessResults_NonExistentProcess() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.setProcessResults(bytes32(0), rzkp, rei);
    }

    function test_SetProcessResults_NotEndedStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Try to set results when process is in READY state (time not expired)
        // Should fail with InvalidTimeBounds because time check happens after status check
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Set to PAUSED and try (time not expired)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Set to CANCELED and try - this should fail with InvalidStatus
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);
    }

    function test_SetProcessResults_ProcessNotEndedByTime() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Try to set results before process duration has passed (should work)
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify process state
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
        assertEq(process.result.length, results.length);
        for (uint i = 0; i < results.length; i++) {
            assertEq(process.result[i], results[i]);
        }
    }

    function test_SetProcessResults_InvalidStateRoot() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        ); // Using different state root

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Try to set results with mismatched state root
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);
    }

    function test_SetProcessResults_CannotSetTwice() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Set results first time
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Try to set results again
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);
    }

    function test_SetProcessResults_InvalidProof() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Try to set results with invalid proof
        bytes memory invalidProof = hex"deadbeef";
        vm.expectRevert(); // The ZK verifier should revert with invalid proof
        processRegistry.setProcessResults(processId, invalidProof, rei);
    }

    function test_SetProcessResults_InvalidInput() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Try to set results with invalid input
        bytes memory invalidInput = hex"deadbeef";
        vm.expectRevert(); // Should revert on abi.decode
        processRegistry.setProcessResults(processId, rzkp, invalidInput);
    }

    function test_SetProcessResults_EmptyResults() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Create empty results input
        bytes memory emptyInput = abi.encode([rInitStateRoot]); // Only state root, no results

        vm.expectRevert(); // Should revert as input must have at least one result
        processRegistry.setProcessResults(processId, rzkp, emptyInput);
    }

    function test_SetProcessResults_UnknownProcessIdPrefix() public {
        // Create a process ID with invalid prefix
        bytes32 invalidProcessId = bytes32(0x1000000000000000000000000000000000000000000000000000000000000001);

        vm.expectRevert(IProcessRegistry.UnknownProcessIdPrefix.selector);
        processRegistry.setProcessResults(invalidProcessId, rzkp, rei);
    }

    function test_SetProcessResults_TimeExpired_ReadyStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Process is in READY status, fast forward time beyond duration
        vm.warp(block.timestamp + 1001);

        // Should succeed because time has expired
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify process state
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
        assertEq(process.result.length, results.length);
        for (uint i = 0; i < results.length; i++) {
            assertEq(process.result[i], results[i]);
        }
    }

    function test_SetProcessResults_TimeExpired_PausedStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // Fast forward time beyond duration
        vm.warp(block.timestamp + 1001);

        // Should succeed because time has expired
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify process state
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
        assertEq(process.result.length, results.length);
    }

    function test_SetProcessResults_TimeNotExpired_ReadyStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Process is in READY status and time has NOT expired
        // Should fail because neither ENDED status nor time expired
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);
    }

    function test_SetProcessResults_TimeNotExpired_PausedStatus() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // Time has NOT expired
        // Should fail because neither ENDED status nor time expired
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);
    }

    function test_SetProcessResults_ExactlyAtExpiration() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Get process to check start time and duration
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        uint256 expirationTime = process.startTime + process.duration;

        // Warp to exact expiration time
        vm.warp(expirationTime);

        // Should succeed at exact expiration boundary
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify results were set
        process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
    }

    function test_SetProcessResults_OneSecondBeforeExpiration() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Get process to check start time and duration
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        uint256 expirationTime = process.startTime + process.duration;

        // Warp to 1 second before expiration
        vm.warp(expirationTime - 1);

        // Should fail because time has not expired
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);
    }

    function test_SetProcessResults_OneSecondAfterExpiration() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Get process to check start time and duration
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        uint256 expirationTime = process.startTime + process.duration;

        // Warp to 1 second after expiration
        vm.warp(expirationTime + 1);

        // Should succeed
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify results were set
        process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
    }

    function test_SetProcessResults_ByNonAdmin() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Submit results from different address (not the process admin)
        address randomUser = address(0xbeef);
        vm.prank(randomUser);
        processRegistry.setProcessResults(processId, rzkp, rei);
        vm.stopPrank();

        // Verify results were set successfully (permissionless)
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
    }

    function test_SetProcessResults_EventsEmitted() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Prepare expected results array
        uint256[] memory expectedResults = new uint256[](results.length);
        for (uint256 i = 0; i < results.length; i++) {
            expectedResults[i] = results[i];
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

        processRegistry.setProcessResults(processId, rzkp, rei);
    }

    function test_SetProcessResults_OldStatusCaptured_FromReady() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
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

        processRegistry.setProcessResults(processId, rzkp, rei);
    }

    function test_SetProcessResults_OldStatusCaptured_FromPaused() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
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

        processRegistry.setProcessResults(processId, rzkp, rei);
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
            rInitStateRoot
        );

        // Wait for 1 second
        vm.warp(block.timestamp + 1);

        // Should succeed
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify results
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
    }

    function test_SetProcessResults_SingleResultValue() public {
        // The rei constant already has valid proof for 8 results, but we're testing decode behavior
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Set results (will extract all values after state root)
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify results array is created correctly
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(uint(process.status), uint(IProcessRegistry.ProcessStatus.RESULTS));
        assertEq(process.result.length, 8); // 9 inputs - 1 state root = 8 results
    }

    function test_SetProcessResults_VerifyStateRootMatch() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Verify the process has the correct state root before setting results
        IProcessRegistry.Process memory processBefore = processRegistry.getProcess(processId);
        assertEq(processBefore.latestStateRoot, rInitStateRoot);

        // Set results
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify state root remains the same after setting results
        IProcessRegistry.Process memory processAfter = processRegistry.getProcess(processId);
        assertEq(processAfter.latestStateRoot, rInitStateRoot);
    }

    function test_SetProcessResults_ResultsArrayLengthCorrect() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Set results
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify results array length matches input (9 inputs - 1 state root = 8 results)
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.result.length, 8);
    }

    function test_SetProcessResults_MultipleProcessesSameOrganization() public {
        // Create first process
        bytes32 processId1 = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        processRegistry.setProcessStatus(processId1, IProcessRegistry.ProcessStatus.ENDED);

        // Create second process with different state root
        bytes32 processId2 = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        processRegistry.setProcessStatus(processId2, IProcessRegistry.ProcessStatus.ENDED);

        // Set results for first process
        processRegistry.setProcessResults(processId1, rzkp, rei);

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
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Get initial counters
        IProcessRegistry.Process memory processBefore = processRegistry.getProcess(processId);
        uint256 voteCountBefore = processBefore.voteCount;
        uint256 voteOverwriteCountBefore = processBefore.voteOverwriteCount;
        uint256 batchNumberBefore = processBefore.batchNumber;

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Set results
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify counters remain unchanged
        IProcessRegistry.Process memory processAfter = processRegistry.getProcess(processId);
        assertEq(processAfter.voteCount, voteCountBefore);
        assertEq(processAfter.voteOverwriteCount, voteOverwriteCountBefore);
        assertEq(processAfter.batchNumber, batchNumberBefore);
    }

    function test_SetProcessResults_ResultsMatchInput() public {
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            rInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Set results
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Verify each result value matches the expected values
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        for (uint256 i = 0; i < results.length; i++) {
            assertEq(process.result[i], results[i], "Result mismatch at index");
        }
    }

    // ========== Process Getters Tests ==========
    function test_GetNextProcessId_Basic() public {
        // Get the next process ID for the organization
        vm.startPrank(orgId);
        bytes32 nextProcessId = processRegistry.getNextProcessId(orgId);
        uint64 currentNonce = processRegistry.processNonce(orgId);
        assertEq(currentNonce, uint64(0));
        // Create a new process
        bytes32 processId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );

        // Verify that the next process ID matches the created process ID
        assertEq(nextProcessId, processId);
        uint64 nextCurrentNonce = processRegistry.processNonce(orgId);
        assertEq(nextCurrentNonce, uint64(1));

        // Create another process
        bytes32 otherNextProcessId = processRegistry.getNextProcessId(orgId);
        bytes32 otherProcessId = createTestProcess(
            defaultBallotMode,
            stInitStateRoot,
            IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1
        );
        assertEq(otherNextProcessId, otherProcessId);
        vm.stopPrank();
    }

    function test_newProcess_CensusOrigin() public {
        // Create a process with very short duration (1 second)
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1,
            censusRoot: 0x0,
            censusURI: "https://example.com/census"
        });
        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({
            x: uint256(keccak256(abi.encodePacked(block.timestamp, "x"))),
            y: uint256(keccak256(abi.encodePacked(block.timestamp, "y")))
        });

        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp, // current time
            1000,
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            stInitStateRoot
        );
        cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.CSP_EDDSA_BN254_V1,
            censusRoot: 0x0,
            censusURI: "https://example.com/census"
        });

        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp, // current time
            1000,
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            stInitStateRoot
        );

        cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.CSP_EDDSA_BN254_V1,
            censusRoot: bytes32(TestHelpers.CENSUS_ROOT),
            censusURI: "https://example.com/census"
        });

        processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp, // current time
            1000,
            defaultBallotMode,
            cen,
            "https://example.com/metadata/",
            key,
            stInitStateRoot
        );
        vm.stopPrank();
    }
}
