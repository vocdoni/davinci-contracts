// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import { Script, console } from "forge-std/Script.sol";
import { OrganizationRegistry } from "../../src/non-proxy/OrganizationRegistry.sol";
import { ProcessRegistry } from "../../src/non-proxy/ProcessRegistry.sol";
import { StateTransitionVerifierGroth16 } from "../../src/StateTransitionVerifierGroth16.sol";

contract TestDeployAllScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("SEPOLIA_PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        OrganizationRegistry organizationRegistry = new OrganizationRegistry();
        console.log("OrganizationRegistry deployed at:", address(organizationRegistry));

        StateTransitionVerifierGroth16 stv = new StateTransitionVerifierGroth16();
        console.log("StateTransitionVerifierGroth16 deployed at:", address(stv));

        ProcessRegistry processRegistry = new ProcessRegistry("11155111", address(organizationRegistry), address(stv));
        console.log("ProcessRegistry deployed at:", address(processRegistry));

        vm.stopBroadcast();
    }
}
