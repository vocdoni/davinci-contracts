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
        hex"2ec28afc568946445b8188ba7a9eed8c3503b57acf585e6803690a02f02da72f2ab7fe9a2ed0fa8158997b6c61771738727e8529242bec80448300ab29b679080f714ddabca5e637b0ea857ed01d6d31e34f65f47b9ea9c9caebaa5808f7e2b627b3035936c38b340889f60434f4ba6eb1208f68f85441afef15932785bca8b42475061108d4fed129aaf7654a1325f7351b0e9c6819f62d97f922964df7aa9327363d908548e015634ee6f95ed824be52b742a8d55e31255d152bca6cb10fb00305489c14f01784fb1d98eb4864ba6bc24dd245f06baffce8d597ea1e5a21be248ccce1880be3274f1059d05dc5f6f7bf53e52f6896ad498608e0e60ebef972276db33e5a454c61ec2fde9ceb0d33862f43e93c9ed753fa6ea8e16f3e192b8c2add78e00adac6c0ffd91bfe8e20b7e0e5cfb4a3bc566799ad371ba6e2a2cd2f05bb9fcf7ff49f2380559831d4d0f092d0e76032d26b6de16e3631dcce9d1e782c34b777dab451675bed9512ff95a34e59bf2f036319aeba17fcc61c0320c341";
    bytes public constant rzkp =
        hex"126448c676044eb81e2cd15cb9c0383596aa3450990e57744c17eb475f2bdd502f8f66a9f7610a9b97743d892fc3f79b632c9a69d7dcc08d4a1239f68fca226c08206218068b9bc835966fbfe081fa44b61ab29aa00529597a9d57d6e3e866a30dea0313b01990038b2dcb83f07fadd555dbdb73cb3a20bde2568e3858bd8a3f232bedae35631c49aee01b02806893ac86bbf8b81542b79029a374076482ed2c07c4d94adcff5b44b99ccbe083f10c87915e2cf5e932d11ae20659cd71c84e1d0c8c853624b49262a60f1c1701015769ac0d7292c36aea95f436ffc3b5e8509a0cf80390c9cef9bf7af8dd5118abbce4dcfac7fd4ad9b8a877e30e6dc9a43a0d0c40b2eb637ebbec1ddc9e679ebb3b604d3fb31ed19e1206c340484e13af263603cec4f52a87f7f4b9d8ebe2b7f455e8b1c112f82f6b267babb8c9d5e852e13a224fef3c8e634797bc3571e232466f7c27d2750b56007e764652f930a1fc51d2267da98091f93f68eb1d5a8b29bcb195a0d5dacc60b0821ecd3986e6b0d2fa72";

    bytes public constant stei =
        hex"1983b53c0e048008bd702af3060684911a3866f56e08f74881143cdf854cca160452e4c03c0464b0d490346235926028e4f8708b358f059229596a7feceac93b00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000005";
    bytes public constant rei =
        hex"0452e4c03c0464b0d490346235926028e4f8708b358f059229596a7feceac93b0000000000000000000000000000000000000000000000000000000000000027000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000001d0000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256 public constant stInitStateRoot =
        11540529015229420508441040255455111908497518547236688677125823343202283604502;
    uint256 public constant rInitStateRoot =
        1955711634442266218442623404753454420621019669225137434447114896701269920059;

    address orgId = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
    uint256[8] public results = [39,46,29,22,42,0,0,0];

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
            uniqueValues: false,
            numFields: 1,
            costExponent: 1,
            maxValue: 10,
            minValue: 0,
            maxValueSum: 100,
            minValueSum: 50
        });
        bytes32 processId1 = createTestProcess(validBallotMode1, stInitStateRoot);
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
        bytes32 processId2 = createTestProcess(validBallotMode2, stInitStateRoot);
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
        bytes32 processId3 = createTestProcess(validBallotMode3, stInitStateRoot);
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
            uniqueValues: false,
            numFields: 1,
            costExponent: 1,
            maxValue: 5,
            minValue: 10, // Invalid: maxValue < minValue
            maxValueSum: 100,
            minValueSum: 50
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
            rInitStateRoot
        );
        processRegistry.submitStateTransition(processId, stzkp, stei);

        // Verify state after transition
        process = processRegistry.getProcess(processId);
        assertEq(
            process.latestStateRoot,
            rInitStateRoot
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
