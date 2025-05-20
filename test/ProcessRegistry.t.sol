// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { ProcessRegistry } from "../src/non-proxy/ProcessRegistry.sol";
import { OrganizationRegistry } from "../src/non-proxy/OrganizationRegistry.sol";
import { Groth16Verifier } from "../src/Groth16Verifier.sol";
import { IProcessRegistry } from "../src/IProcessRegistry.sol";

contract ProcessRegistryTest is Test {
    ProcessRegistry public processRegistry;
    OrganizationRegistry public organizationRegistry;
    Groth16Verifier public groth16Verifier;

    bytes public zkp =
        hex"0ecab3cb35e0fe048245884fd3714a38d5f252a635eee1daaefcecedbe28557d04ca7508c53b20f185cf07cee7769eb259d6b95862cd6b3bb283828e1f2aad3115d502cc34fb6f67960d0ba635f26ca472e92507b34269ab274d5b32942b53e20983e9e55869b7a6894d5741a8eb92616910cb07faffc943309c1a75b60e41652ed5cb3a891b992a9e65b617f92d90fe27df72b22026f64973537d84b759aa2c2f96d7d7cf3a22184b05c7a697ef2539cff420956944da825a50f2017fa9362d05cb4d8a84499313b004bfc497bd180bf614c1e35e96fcae757df780c7e081db1c0ac6c54f917aee5df7532b681d39d2d584f6fd4c306f04870cb070db3eddb523cfbb5ec47f2ba67e73558f057f5fa22da5e1bdea9858a0553c159d3edee8f50779f592936e47a69b718f08fdaed9a34b2660fe110e1c38607cf3fc3bd0bc4612a36b159967b420a0c0860efd43eae410620b58e2bf639c795100e899baf22926a7512d00d5f9668fb943cb61ee6ee12a036d5d565e6c57ed252a2c1edf0d9b";

    bytes public encodedInputs =
        hex"224ee39d8852a0a5b414cb2a312f9cc812b767a6d676c583b0fe17b82517ca0926131bc948cf601e6038c00507492262a56a511c85c0cb0a19ad411639b1384c00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256 initStateRoot = 15518021866372582043604022785954659526716738998075185570627351931444151175689;

    function setUp() public {
        organizationRegistry = new OrganizationRegistry();
        groth16Verifier = new Groth16Verifier();
        processRegistry = new ProcessRegistry("11155111", address(organizationRegistry), address(groth16Verifier));

        createData();
    }

    function createData() internal {
        createTestOrganization();
        createTestProcess();
    }

    function test_Process_State_Transition() public {
        IProcessRegistry.Process memory process = processRegistry.getProcess(
            hex"000005390056d6ed515b2e0af39bb068f587d0de83facd1b0000000000000004"
        );
        assertEq(process.latestStateRoot, initStateRoot);
        processRegistry.submitStateTransition(
            hex"000005390056d6ed515b2e0af39bb068f587d0de83facd1b0000000000000004",
            zkp,
            encodedInputs
        );
        process = processRegistry.getProcess(hex"000005390056d6ed515b2e0af39bb068f587d0de83facd1b0000000000000004");
        assertEq(
            process.latestStateRoot,
            17221650114163349820651399714058649394077581872287092757887500660867278649420
        );
    }

    function createTestProcess() internal {
        IProcessRegistry.BallotMode memory bm = IProcessRegistry.BallotMode({
            costFromWeight: false,
            forceUniqueness: false,
            maxCount: 5,
            costExponent: 2,
            maxValue: 16,
            minValue: 0,
            maxTotalCost: 1280,
            minTotalCost: 5
        });
        IProcessRegistry.Census memory cen = IProcessRegistry.Census({
            censusOrigin: IProcessRegistry.CensusOrigin.OFF_CHAIN_TREE,
            maxVotes: 1000,
            censusRoot: 0x59a5002406c534a8f713bd96d6ff0fb8d84828aceeba5e26808a0f2df0cc9c03,
            censusURI: "https://example.com/census"
        });
        IProcessRegistry.EncryptionKey memory key = IProcessRegistry.EncryptionKey({
            x: 9008594880724679043974778531972871007153582286259636428553542667152129354466,
            y: 11390364202388007336065472767767585257356174356794824068569004670923989712381
        });
        processRegistry.newProcess(
            IProcessRegistry.ProcessStatus.READY,
            0,
            1000000,
            bm,
            cen,
            "https://example.com/metadata",
            0x0056d6eD515B2e0af39bb068F587D0de83fACD1B,
            hex"000005390056d6ed515b2e0af39bb068f587d0de83facd1b0000000000000004",
            key,
            initStateRoot
        );
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
}
