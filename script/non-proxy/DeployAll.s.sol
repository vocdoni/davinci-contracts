// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import { Script, console } from "forge-std/Script.sol";
import { OrganizationRegistry } from "../../src/non-proxy/OrganizationRegistry.sol";
import { ProcessRegistry } from "../../src/non-proxy/ProcessRegistry.sol";
import { Groth16Verifier } from "../../src/Groth16Verifier.sol";

contract TestDeployAllScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("SEPOLIA_PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        OrganizationRegistry organizationRegistry = new OrganizationRegistry();
        console.log("OrganizationRegistry deployed at:", address(organizationRegistry));

        Groth16Verifier groth16Verifier = new Groth16Verifier();
        console.log("Groth16Verifier deployed at:", address(groth16Verifier));

        ProcessRegistry processRegistry = new ProcessRegistry(
            "11155111",
            address(organizationRegistry),
            address(groth16Verifier)
        );
        console.log("ProcessRegistry deployed at:", address(processRegistry));

        vm.stopBroadcast();
    }
}
