// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { ProcessRegistry } from "../src/ProcessRegistry.sol";
import { OrganizationRegistry } from "../src/OrganizationRegistry.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";
import { IProcessRegistry } from "../src/interfaces/IProcessRegistry.sol";

contract ProcessRegistryTest is Test {
    ProcessRegistry public processRegistry;
    OrganizationRegistry public organizationRegistry;
    StateTransitionVerifierGroth16 public stv;
    ResultsVerifierGroth16 public rv;

    bytes public constant stzkp =
        hex"2fae5436ab7f722800bf8f04d0141f8bd8bb45fe6f697fb17de992fc213c38990ab27a48bf7d7707b87043ecc0e3c49f1cbe05437145285f7dec429d5a923506053e36ba3d95b0d8854d1cc8c61d2e26484e8c127e84c07edb8e7fc5ded4e73e05df7d9f586322f7be883b3cd98efd5ebe2d976269f67a50ccc83629983c61b70796813e0681f073711a783b65ea906ce5a091f99756f1dbfc36448e4b957a4f1fef874f0919eef257a3e20065ccbeb54f93353b4e4697b59f836fcd10329f9218a4b5117979d61dec2526b7a2c65e6f24d687140c42224dffbb3af67230d12b28c200ca8db1fbf553ca5100aa717e21be7fd01018797c98222b0193ae342c532d3064567ee8d80219a2d9f8f9f5ae6d9c0ab2fe0b3bf2fdecca83059d9096c9286122c449fbf3746250995ffaa4a76bb027db97ae89cab1cd92b933c3e8470a0b1538acb5269c96454f9f5ce37661596f332f1838ab051bbc309124a81b136a2e8029dba69e96d8d9e6546f1160425a62d2b280bb7dd45912a17863f8d9680c";
    bytes public constant rzkp =
        hex"0487ebaa3d2f48b09a97fa9be27d15cd945624ab328056840e844fb8458b6f0d0305178df1ffbb419bb26d05302ebf8a2161b5868f99d153491456456bfae3821ba68ec7a1c8b968a1322eb402f9f8d0d6776f16ee557f560f86efaa78c3a00a2f1d11ef7bd8492f28ce881360c70f227ce55d8eec55cc9b5ecd0be024a976332ff734345491ec7f98ef341c29b0f65ddb31d1e9f8a2961f649056f383b6c36302bd58fb2414dc5c911c215f740fe34d3349d21e9a4dfc35944d5be1dedeb1e02c66410e6938b9b64a6170b25e0a21db9df7448cdda519290c86e59462f606f90210bcfb97de66bda8af8db0e1d9e4de84892e1cb01cf7661d147977543c65142a7dd4d5a23eaf0f75ca070c83822673c71d7ce76e5cbb8f7b7e56a29b96519005ad26c999fb1bc370771e1d43d4ba1261f569d0bdad74ff48db02ff12e56f3214e341fae2089cedbace782230ea1749af3060e36e6f139331a52d63402203bd0ac45b18912576cec84274054d72f3a9553df038a1e8fb4ab9d5f5036fe79cc1";

    bytes public constant stei =
        hex"1fe3b71818c0557f2f40ec198c8cf176b54d37e9b6af21f71ae2fc13ca2578850b987778b1b51747b94a15831946b0eb2431ad7c2a3608c733419986189553f900000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000005";
    bytes public constant rei =
        hex"0b987778b1b51747b94a15831946b0eb2431ad7c2a3608c733419986189553f9000000000000000000000000000000000000000000000000000000000000003400000000000000000000000000000000000000000000000000000000000000390000000000000000000000000000000000000000000000000000000000000017000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256 public constant stInitStateRoot =
        14424036259012750574058040874268427628643176097624202352033583206326205773957;
    uint256 public constant rInitStateRoot =
        5244826649985884941716637507837432287554322849727016690652948098167443117049;

    address orgId = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
    uint256[8] public results = [52,57,23,42,46,0,0,0];

    IProcessRegistry.BallotMode public defaultBallotMode =
        IProcessRegistry.BallotMode({
            costFromWeight: false,
            forceUniqueness: false,
            maxCount: 5,
            costExponent: 2,
            maxValue: 16,
            minValue: 0,
            maxTotalCost: 1280,
            minTotalCost: 5
        });

    function setUp() public {
        organizationRegistry = new OrganizationRegistry();
        stv = new StateTransitionVerifierGroth16();
        rv = new ResultsVerifierGroth16();
        processRegistry = new ProcessRegistry(11155111, address(stv), address(rv));

        createTestOrganization();
    }

    function createTestOrganization() internal {
        address[] memory admins = new address[](1);
        admins[0] = orgId;
        organizationRegistry.createOrganization("testOrganization", "https://example.com/metadata", admins);
    }

    function createTestProcess(
        IProcessRegistry.BallotMode memory ballotMode,
        uint256 initStateRoot
    ) internal returns (bytes32) {
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 1000,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
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
        processRegistry.setProcessStatus(
            bytes32(0x1000000000000000000000000000000000000000000000000000000000000001),
            IProcessRegistry.ProcessStatus.ENDED
        );
    }

    function test_SetProcessStatus_NotAdmin() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);
        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        vm.stopPrank();
    }

    function test_SetProcessStatus_SameStatus() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);
    }

    function test_SetProcessStatus_FromReady() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);
        // READY -> PAUSED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.PAUSED));

        // Reset to READY
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);

        // READY -> CANCELED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.CANCELED));

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // READY -> ENDED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.ENDED));

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // READY -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromPaused() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

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
        processId = createTestProcess(defaultBallotMode, stInitStateRoot);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // PAUSED -> ENDED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.ENDED));

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, stInitStateRoot);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // PAUSED -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromEnded() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Set initial state to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        // ENDED -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, stInitStateRoot);
        vm.warp(block.timestamp + 1001);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // ENDED -> CANCELED (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);

        // Reset process for next test
        processId = createTestProcess(defaultBallotMode, stInitStateRoot);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // ENDED -> READY (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);

        // ENDED -> PAUSED (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);
    }

    function test_SetProcessStatus_FromCanceled() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

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
        bytes32 processId = createTestProcess(defaultBallotMode, rInitStateRoot);

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
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

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

    // ========== Process Census Tests ==========

    function test_SetProcessCensus_Success() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 2000, // Increased from initial 1000
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        processRegistry.setProcessCensus(processId, newCensus);

        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.census.maxVotes, newCensus.maxVotes);
        assertEq(process.census.censusRoot, newCensus.censusRoot);
        assertEq(process.census.censusURI, newCensus.censusURI);
    }

    function test_SetProcessCensus_NonExistentProcess() public {
        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 2000,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.setProcessCensus(bytes32(0), newCensus);
    }

    function test_SetProcessCensus_NotAdmin() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 2000,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        processRegistry.setProcessCensus(processId, newCensus);
        vm.stopPrank();
    }

    function test_SetProcessCensus_InvalidCensus_EmptyURI() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 2000,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: ""
        });

        vm.expectRevert(IProcessRegistry.InvalidCensusURI.selector);
        processRegistry.setProcessCensus(processId, newCensus);
    }

    function test_SetProcessCensus_InvalidCensus_ZeroCensusRoot() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 2000,
            censusRoot: 0,
            censusURI: "https://example.com/new-census"
        });

        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        processRegistry.setProcessCensus(processId, newCensus);
    }

    function test_SetProcessCensus_InvalidCensus_DecreaseMaxVotes() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 500, // Less than initial 1000
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        vm.expectRevert(IProcessRegistry.InvalidMaxVotes.selector);
        processRegistry.setProcessCensus(processId, newCensus);
    }

    function test_SetProcessCensus_InvalidStatus_Canceled() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Set process to CANCELED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 2000,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessCensus(processId, newCensus);
    }

    function test_SetProcessCensus_ValidStatus_Paused() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 2000,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        // Should succeed
        processRegistry.setProcessCensus(processId, newCensus);

        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.census.maxVotes, newCensus.maxVotes);
        assertEq(process.census.censusRoot, newCensus.censusRoot);
        assertEq(process.census.censusURI, newCensus.censusURI);
    }

    // ========== Ballot Mode Tests ==========

    function test_ValidateBallotMode_ValidCases() public {
        // Test case 1: Basic valid ballot mode
        IProcessRegistry.BallotMode memory validBallotMode1 = IProcessRegistry.BallotMode({
            costFromWeight: false,
            forceUniqueness: false,
            maxCount: 1,
            costExponent: 1,
            maxValue: 10,
            minValue: 0,
            maxTotalCost: 100,
            minTotalCost: 50
        });
        bytes32 processId1 = createTestProcess(validBallotMode1, stInitStateRoot);
        assertTrue(processId1 != bytes32(0));

        // Test case 2: Valid ballot with costFromWeight true
        IProcessRegistry.BallotMode memory validBallotMode2 = IProcessRegistry.BallotMode({
            costFromWeight: true,
            forceUniqueness: true,
            maxCount: 5,
            costExponent: 2,
            maxValue: 100,
            minValue: 1,
            maxTotalCost: 0, // This is valid when costFromWeight is true
            minTotalCost: 0
        });
        bytes32 processId2 = createTestProcess(validBallotMode2, stInitStateRoot);
        assertTrue(processId2 != bytes32(0));

        // Test case 3: Edge case - maxValue equals minValue
        IProcessRegistry.BallotMode memory validBallotMode3 = IProcessRegistry.BallotMode({
            costFromWeight: false,
            forceUniqueness: false,
            maxCount: 8,
            costExponent: 1,
            maxValue: 5,
            minValue: 5,
            maxTotalCost: 50,
            minTotalCost: 50
        });
        bytes32 processId3 = createTestProcess(validBallotMode3, stInitStateRoot);
        assertTrue(processId3 != bytes32(0));
    }

    function test_ValidateBallotMode_InvalidMaxCount() public {
        IProcessRegistry.BallotMode memory invalidBallotMode = IProcessRegistry.BallotMode({
            costFromWeight: false,
            forceUniqueness: false,
            maxCount: 0, // Invalid: must be >= 1
            costExponent: 1,
            maxValue: 10,
            minValue: 0,
            maxTotalCost: 100,
            minTotalCost: 50
        });

        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 1000,
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
            forceUniqueness: false,
            maxCount: 1,
            costExponent: 1,
            maxValue: 5,
            minValue: 10, // Invalid: maxValue < minValue
            maxTotalCost: 100,
            minTotalCost: 50
        });

        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 1000,
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
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);
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
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        processRegistry.setProcessDuration(processId, 2000000);
        vm.stopPrank();
    }

    function test_SetProcessDuration_InvalidStatus_Canceled() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Set process to CANCELED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessDuration(processId, 2000000);
    }

    function test_SetProcessDuration_ValidStatus_Paused() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);
        uint256 newDuration = 2000000;

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        emit IProcessRegistry.ProcessDurationChanged(processId, newDuration);
        processRegistry.setProcessDuration(processId, newDuration);

        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.duration, newDuration);
    }

    function test_SetProcessDuration_InvalidDuration_Zero() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        processRegistry.setProcessDuration(processId, 0);
    }

    function test_SetProcessDuration_InvalidDuration_PastEndTime() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Try to set a duration that would make the process end in the past
        uint256 invalidDuration = 1; // Very short duration that will definitely be in the past

        vm.expectRevert(IProcessRegistry.InvalidDuration.selector);
        processRegistry.setProcessDuration(processId, invalidDuration);
    }

    function test_SetProcessDuration_MaxDuration() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);
        uint256 maxDuration = type(uint256).max - block.timestamp;

        emit IProcessRegistry.ProcessDurationChanged(processId, maxDuration);
        processRegistry.setProcessDuration(processId, maxDuration);

        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.duration, maxDuration);
    }

    // ========== State Transition Tests ==========

    function test_SubmitStateTransition_Success() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);

        // Verify initial state
        assertEq(process.latestStateRoot, stInitStateRoot);
        assertEq(process.voteCount, 0);
        assertEq(process.voteOverwriteCount, 0);

        // Submit state transition
        emit IProcessRegistry.ProcessStateRootUpdated(
            processId,
            address(this),
            5244826649985884941716637507837432287554322849727016690652948098167443117049
        );
        processRegistry.submitStateTransition(processId, stzkp, stei);

        // Verify state after transition
        process = processRegistry.getProcess(processId);
        assertEq(
            process.latestStateRoot,
            5244826649985884941716637507837432287554322849727016690652948098167443117049
        );
        assertEq(process.voteCount, uint256(abi.decode(stei, (uint256[4]))[2]));
        assertEq(process.voteOverwriteCount, uint256(abi.decode(stei, (uint256[4]))[3]));
    }

    function test_SubmitStateTransition_NonExistentProcess() public {
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        processRegistry.submitStateTransition(bytes32(0), stzkp, stei);
    }

    function test_SubmitStateTransition_InvalidStatus_Paused() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Set process to PAUSED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, stzkp, stei);
    }

    function test_SubmitStateTransition_InvalidStatus_Ended() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, stzkp, stei);
    }

    function test_SubmitStateTransition_InvalidStatus_Canceled() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Set process to CANCELED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, stzkp, stei);
    }

    function test_SubmitStateTransition_InvalidStatus_Results() public {
        bytes32 processId = createTestProcess(defaultBallotMode, rInitStateRoot);

        // Set process to ENDED then RESULTS
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        processRegistry.setProcessResults(processId, rzkp, rei);

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.submitStateTransition(processId, stzkp, stei);
    }

    function test_SubmitStateTransition_ProofInvalid() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Create invalid encoded inputs with wrong state root
        bytes memory invalidEncodedInputs = abi.encode(
            [
                uint256(123), // Wrong state root
                uint256(13999159323556783713340982889683090407637562075598756749526547570965285300122),
                uint256(5),
                uint256(0)
            ]
        );

        vm.expectRevert(IProcessRegistry.ProofInvalid.selector);
        processRegistry.submitStateTransition(processId, stzkp, invalidEncodedInputs);
    }

    function test_SubmitStateTransition_InvalidTimeBounds() public {
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Fast forward time beyond process duration
        vm.warp(block.timestamp + 1000001);

        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        processRegistry.submitStateTransition(processId, stzkp, stei);
    }

    // ========== Process Results Tests ==========

    function test_SetProcessResults_Success() public {
        bytes32 processId = createTestProcess(defaultBallotMode, rInitStateRoot);

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
        bytes32 processId = createTestProcess(defaultBallotMode, rInitStateRoot);

        // Try to set results when process is in READY state
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Set to PAUSED and try
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);

        // Set to CANCELED and try
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);
    }

    function test_SetProcessResults_ProcessNotEndedByTime() public {
        bytes32 processId = createTestProcess(defaultBallotMode, rInitStateRoot);

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
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot); // Using different state root

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Try to set results with mismatched state root
        vm.expectRevert(IProcessRegistry.InvalidStateRoot.selector);
        processRegistry.setProcessResults(processId, rzkp, rei);
    }

    function test_SetProcessResults_CannotSetTwice() public {
        bytes32 processId = createTestProcess(defaultBallotMode, rInitStateRoot);

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
        bytes32 processId = createTestProcess(defaultBallotMode, rInitStateRoot);

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
        bytes32 processId = createTestProcess(defaultBallotMode, rInitStateRoot);

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
        bytes32 processId = createTestProcess(defaultBallotMode, rInitStateRoot);

        // Set process to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // Fast forward time to ensure process has ended
        vm.warp(block.timestamp + 1001);

        // Create empty results input
        bytes memory emptyInput = abi.encode([rInitStateRoot]); // Only state root, no results

        vm.expectRevert(); // Should revert as input must have at least one result
        processRegistry.setProcessResults(processId, rzkp, emptyInput);
    }

    // ========== Process Getters Tests ==========
    function test_GetNextProcessId_Basic() public {
        // Get the next process ID for the organization
        vm.startPrank(orgId);
        bytes32 nextProcessId = processRegistry.getNextProcessId(orgId);
        uint64 currentNonce = processRegistry.processNonce(orgId);
        assertEq(currentNonce, uint64(0));
        // Create a new process
        bytes32 processId = createTestProcess(defaultBallotMode, stInitStateRoot);

        // Verify that the next process ID matches the created process ID
        assertEq(nextProcessId, processId);
        uint64 nextCurrentNonce = processRegistry.processNonce(orgId);
        assertEq(nextCurrentNonce, uint64(1));

        // Create another process
        bytes32 otherNextProcessId = processRegistry.getNextProcessId(orgId);
        bytes32 otherProcessId = createTestProcess(defaultBallotMode, stInitStateRoot);
        assertEq(otherNextProcessId, otherProcessId);
        vm.stopPrank();
    }
}
