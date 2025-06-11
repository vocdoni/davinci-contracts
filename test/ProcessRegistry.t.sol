// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { ProcessRegistry } from "../src/non-proxy/ProcessRegistry.sol";
import { OrganizationRegistry } from "../src/non-proxy/OrganizationRegistry.sol";
import { StateTransitionVerifierGroth16 } from "../src/StateTransitionVerifierGroth16.sol";
import { ResultsVerifierGroth16 } from "../src/ResultsVerifierGroth16.sol";
import { IProcessRegistry } from "../src/IProcessRegistry.sol";

contract ProcessRegistryTest is Test {
    ProcessRegistry public processRegistry;
    OrganizationRegistry public organizationRegistry;
    StateTransitionVerifierGroth16 public stv;
    ResultsVerifierGroth16 public rv;

    bytes public zkp =
        hex"2e688a61c5924ca3464fce45d4b1b1639b31e3c7fc6e4e74244fd6af74ea57522e9aea56c395612a89f15176942a5f5a264d4125e00c024f725e00fe9d52765a226cd50609972adb62b8c71a5b1aa60b6dc298075b44203d1cecdf3fe43c7a171e586cc471d3934312167ded54fab40cdf81e61abd3608c830ae730759e2c87f0728ad0f1277c067e26ea451d337abc81abfb83cf1c972f2d01f14a87ffbe27a1f5253ae069b804783f59ef87db09ef7df1a45b93adc28f0dcc3d8cddcd90ab20ef68f7c1c26efedcfc7f897abc69a34a0dde812ac895baf2c69b1af02493ecb2e5e1f8d4ce0f18bc3e863970eaef46154e94c8f4f2b112a7a636c5283a2fe790741f6373905e5c54bfa57f0b70104f485f24b694baa249491f0a308734548bb02c9b350371a51c8db2b1abe0e01da678da7e826474481ffef66b1dcebad58a82d56fcfbcdaeca5a963a3b510a09bd6cd17afce730f213b729bba8d1da84d982209d534b8a7b2e19f402935ffc64a3897273f82379a8a670a42bc0b00ae157cc";

    bytes public encodedInputs =
        hex"23ffb93b6e4fb1b2098ceff1e0f20ec383761619d80f1d80ea21c652c738d99a1ef33e4eac84b3192df54859cf12c0a632f66912dbd1eee0ebc9d1dcb8bde79a00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256 initStateRoot = 16282774127260957557439015348635445656767673488688041267615841575933035731354;

    address orgId = 0x0056d6eD515B2e0af39bb068F587D0de83fACD1B;

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
        processRegistry = new ProcessRegistry("11155111", address(organizationRegistry), address(stv), address(rv));

        createTestOrganization();
    }

    function test_Process_State_Transition() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);
        IProcessRegistry.Process memory process = processRegistry.getProcess(processId);
        assertEq(process.latestStateRoot, initStateRoot);
        processRegistry.submitStateTransition(processId, zkp, encodedInputs);
        process = processRegistry.getProcess(processId);
        assertEq(
            process.latestStateRoot,
            13999159323556783713340982889683090407637562075598756749526547570965285300122
        );
    }

    function createTestProcess(
        address organizationId,
        IProcessRegistry.BallotMode memory ballotMode
    ) internal returns (bytes32) {
        // Create a random process ID using organizationId and a random nonce
        uint32 processCount = processRegistry.processCount();
        bytes32 processId = bytes32(
            abi.encodePacked(
                uint32(block.timestamp), // timestamp as prefix
                organizationId, // organization address
                uint32(processCount) // nonce
            )
        );

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

        processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            block.timestamp, // current time
            1000000,
            ballotMode,
            cen,
            "https://example.com/metadata/",
            organizationId,
            processId,
            key,
            initStateRoot
        );

        return processId;
    }

    function createTestOrganization() internal {
        address[] memory admins = new address[](1);
        admins[0] = address(0x1804c8AB1F12E6bbf3894d4083f33e07309d1f38);
        organizationRegistry.createOrganization(
            0x0056d6eD515B2e0af39bb068F587D0de83fACD1B,
            "testOrganization",
            "https://example.com/metadata",
            admins
        );
    }

    // ========== Process Status Tests ==========

    function test_SetProcessStatus_NonExistentProcess() public {
        vm.expectRevert(IProcessRegistry.ProcessNotFound.selector);
        processRegistry.setProcessStatus(bytes32(0), IProcessRegistry.ProcessStatus.ENDED);
    }

    function test_SetProcessStatus_NotAdmin() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);
        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.NotOrganizationAdministrator.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        vm.stopPrank();
    }

    function test_SetProcessStatus_SameStatus() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);
    }

    function test_SetProcessStatus_FromReady() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);
        // READY -> PAUSED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.PAUSED));

        // Reset to READY
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);

        // READY -> CANCELED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.CANCELED));

        // Reset process for next test
        processId = createTestProcess(orgId, defaultBallotMode);

        // READY -> ENDED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.ENDED));

        // Reset process for next test
        processId = createTestProcess(orgId, defaultBallotMode);

        // READY -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromPaused() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

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
        processId = createTestProcess(orgId, defaultBallotMode);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // PAUSED -> ENDED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.ENDED));

        // Reset process for next test
        processId = createTestProcess(orgId, defaultBallotMode);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        // PAUSED -> RESULTS (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);
    }

    function test_SetProcessStatus_FromEnded() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

        // Set initial state to ENDED
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // ENDED -> RESULTS (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.RESULTS));

        // Reset process for next test
        processId = createTestProcess(orgId, defaultBallotMode);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // ENDED -> CANCELED (valid)
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.CANCELED);
        assertEq(uint(processRegistry.getProcess(processId).status), uint(IProcessRegistry.ProcessStatus.CANCELED));

        // Reset process for next test
        processId = createTestProcess(orgId, defaultBallotMode);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);

        // ENDED -> READY (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);

        // ENDED -> PAUSED (invalid)
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);
    }

    function test_SetProcessStatus_FromCanceled() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

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
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

        // Set initial state to ENDED then RESULTS
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);

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
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

        emit IProcessRegistry.ProcessStatusChanged(processId, IProcessRegistry.ProcessStatus.PAUSED);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.PAUSED);

        emit IProcessRegistry.ProcessStatusChanged(processId, IProcessRegistry.ProcessStatus.READY);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.READY);

        emit IProcessRegistry.ProcessStatusChanged(processId, IProcessRegistry.ProcessStatus.ENDED);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
    }

    // ========== Process Census Tests ==========

    function test_SetProcessCensus_Success() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

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

        vm.expectRevert(IProcessRegistry.ProcessNotFound.selector);
        processRegistry.setProcessCensus(bytes32(0), newCensus);
    }

    function test_SetProcessCensus_NotAdmin() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 2000,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.NotOrganizationAdministrator.selector);
        processRegistry.setProcessCensus(processId, newCensus);
        vm.stopPrank();
    }

    function test_SetProcessCensus_InvalidCensus_EmptyURI() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

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
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

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
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 500, // Less than initial 1000
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        vm.expectRevert(IProcessRegistry.InvalidMaxVotes.selector);
        processRegistry.setProcessCensus(processId, newCensus);
    }

    function test_SetProcessCensus_InvalidStatus_Results() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

        // Set process to ENDED then RESULTS
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.ENDED);
        processRegistry.setProcessStatus(processId, IProcessRegistry.ProcessStatus.RESULTS);

        IProcessRegistry.Census memory newCensus = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 2000,
            censusRoot: 0x123400000000000000000000000000000000000000000000000000000000abcd,
            censusURI: "https://example.com/new-census"
        });

        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        processRegistry.setProcessCensus(processId, newCensus);
    }

    function test_SetProcessCensus_InvalidStatus_Canceled() public {
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

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
        bytes32 processId = createTestProcess(orgId, defaultBallotMode);

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
        bytes32 processId1 = createTestProcess(orgId, validBallotMode1);
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
        bytes32 processId2 = createTestProcess(orgId, validBallotMode2);
        assertTrue(processId2 != bytes32(0));

        // Test case 3: Edge case - maxValue equals minValue
        IProcessRegistry.BallotMode memory validBallotMode3 = IProcessRegistry.BallotMode({
            costFromWeight: false,
            forceUniqueness: false,
            maxCount: 10,
            costExponent: 1,
            maxValue: 5,
            minValue: 5,
            maxTotalCost: 50,
            minTotalCost: 50
        });
        bytes32 processId3 = createTestProcess(orgId, validBallotMode3);
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
        vm.expectRevert(IProcessRegistry.InvalidMaxCount.selector);
        createTestProcess(orgId, invalidBallotMode);
    }

    // function test_ValidateBallotMode_InvalidValueRange() public {
    //     IProcessRegistry.BallotMode memory invalidBallot = IProcessRegistry.BallotMode({
    //         costFromWeight: false,
    //         forceUniqueness: false,
    //         maxCount: 1,
    //         costExponent: 1,
    //         maxValue: 5,
    //         minValue: 10, // Invalid: maxValue < minValue
    //         maxTotalCost: 100,
    //         minTotalCost: 50
    //     });
    //     vm.expectRevert(IProcessRegistry.InvalidMinValue.selector);
    //     createTestProcess(orgId, invalidBallot);
    // }

    // function test_ValidateBallotMode_InvalidMaxTotalCost_NonWeighted() public {
    //     IProcessRegistry.BallotMode memory invalidBallot = IProcessRegistry.BallotMode({
    //         costFromWeight: false,
    //         forceUniqueness: false,
    //         maxCount: 1,
    //         costExponent: 1,
    //         maxValue: 10,
    //         minValue: 0,
    //         maxTotalCost: 0, // Invalid when costFromWeight is false
    //         minTotalCost: 0
    //     });
    //     //vm.expectRevert(IProcessRegistry.InvalidMaxTotalCost.selector);
    //     createTestProcess(orgId, invalidBallot);
    // }

    // function test_ValidateBallotMode_InvalidTotalCostBounds() public {
    //     IProcessRegistry.BallotMode memory invalidBallot = IProcessRegistry.BallotMode({
    //         costFromWeight: false,
    //         forceUniqueness: false,
    //         maxCount: 1,
    //         costExponent: 1,
    //         maxValue: 10,
    //         minValue: 0,
    //         maxTotalCost: 50,
    //         minTotalCost: 100 // Invalid: maxTotalCost < minTotalCost
    //     });
    //     //vm.expectRevert(IProcessRegistry.InvalidTotalCostBounds.selector);
    //     createTestProcess(orgId, invalidBallot);
    // }
}
