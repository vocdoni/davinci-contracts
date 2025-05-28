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
        hex"078ae3e6c6c404cb161927a4dc711893798845cccc3b226ec56d0da22e16e1cb2ea60a232e5028dd56fbbf1beb512cfa933552d76ce4eb481e643b645c641e022226f0c2536331c01e39e71dc4ed1d8d761408d5aa94206174a9af6a0405589013031e68c4a1de5007d3915249f7f685d90d9a4235943c90785675a2e8d7d35a29be5d62953b0ee9629e5535077c9fd46e5bd2989922308ee5dc0ad7dab48d8c21537b7a5eb721cb271622974c96c000992c53b7d5872c16082eacee81bbdd9c17d329bd56549866e695a25cbcab95d1a61f7124a751b7dbc1dc1fc9b6ea92af2727e2440e22d8f263298d520bd6059094e50cefcb4d4dc7dc26c99a62deb4bc12d7bf39651303fdb7d56edd404c159b68baa249ddac050b220dedd75a0a736816b60bcf0e6ee111aef390daf2a76574d00a3ecc49912e7626c4a31b0dcc01921dd5a77472474f9df1dea6620d54e6ca4b366178995c911ff03cdececfc344841048e3d00072b27542d626d1326d1e51f039b182f31edcbd066434eeb4d838a0";

    bytes public encodedInputs =
        hex"17827599204d66e404f804ab4a1d7f1f8a86af5196d5d2a718b2a3e13ffbeb0f15ca9437fcfa27637465824011e905ff484485742595c25132d2c027994bb77000000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256 initStateRoot = 10633697268432566100547111228192675718879714909099079783139311163602236074767;

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
        assertEq(process.latestStateRoot, 9856495895231824317321181173471266818915032605823360987155589090285852538736);
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
