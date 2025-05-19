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
        hex"0c07d1fa200162652c1432be2b64524fae8ea94aa523a02e04fabadb3706c7c60e3ef6ab9394aee502c94bf66fd6013c78614b80194ddb9738498b0c0fd589ef213ff6f9546e83f9ce31c051e33d8f493c3faee50641bf7bfa0c0293e6850b6428acb11796557f70555af9dd68fe492c325f3553bf4af89eb43a621440e87c4c1f165c6f9e2d9ffe9180e06938ed13dda476c139eda6e0226a3251d1f8564d32201ceee55155227d5a36d432cc6f8af6c87c92b89a7bb4804a837ac50e67930e230cdbdfb0a4d7e6e2752f53de9ac04449b862897b5db0c003a73e1b2ebab4371d7f943089cef554ed58fc50934d2197c924fdeb6b4f77e87d7ef0bcd8d3082e25833c49c60712d1955fc055a6f302769a784ce8adf719fd4aa966419c686bd82b37c3a5016405db96e644e5590ac98cbd92a4931cf444c269e3a80bbf1135ba1de48f00a7060484fdbe01dba1ba5797994b0170512cce78ed4482fa930007ba1e8fb338c102998fc2108fb434a3499dc27e8f0b718bc9ce974adee38635e571";

    bytes public encodedInputs =
        hex"2049d2a4bdcb60f7f241e4115494be9ea41e0ffa4b12f4f204c252c6b37c1df3284be1b319af8f126a739ea5abb992545b2b3238a93c969e26b91ceb5849b52d00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256 initStateRoot = 14604444798545114219184112291507703832081104487914651663816669607190240959987;

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
            18226585194652543680303489095944698048129925574222676985726996869319177647405
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
