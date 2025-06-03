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

    function setUp() public {
        organizationRegistry = new OrganizationRegistry();
        stv = new StateTransitionVerifierGroth16();
        rv = new ResultsVerifierGroth16();
        processRegistry = new ProcessRegistry("11155111", address(organizationRegistry), address(stv), address(rv));

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
            13999159323556783713340982889683090407637562075598756749526547570965285300122
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
