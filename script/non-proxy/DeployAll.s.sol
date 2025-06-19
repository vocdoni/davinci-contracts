// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import { Script, console } from "forge-std/Script.sol";
import { OrganizationRegistry } from "../../src/non-proxy/OrganizationRegistry.sol";
import { ProcessRegistry } from "../../src/non-proxy/ProcessRegistry.sol";
import { StateTransitionVerifierGroth16 } from "../../src/verifiers/StateTransitionVerifierGroth16.sol";
import { ResultsVerifierGroth16 } from "../../src/verifiers/ResultsVerifierGroth16.sol";

contract DeployAllScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("UZH_PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deployer address:", deployerAddress);
        vm.startBroadcast(deployerPrivateKey);

        OrganizationRegistry organizationRegistry = new OrganizationRegistry();
        console.log("OrganizationRegistry deployed at:", address(organizationRegistry));

        StateTransitionVerifierGroth16 stv = new StateTransitionVerifierGroth16();
        console.log("StateTransitionVerifierGroth16 deployed at:", address(stv));

        ResultsVerifierGroth16 rv = new ResultsVerifierGroth16();
        console.log("ResultsVerifierGroth16 deployed at:", address(rv));

        uint256 chainId = vm.envUint("UZH_CHAIN_ID");
        ProcessRegistry processRegistry = new ProcessRegistry(
            uint32(chainId), // change this to the desired chain ID
            address(stv),
            address(rv)
        );
        console.log("ProcessRegistry deployed at:", address(processRegistry));

        vm.stopBroadcast();
    }
}
