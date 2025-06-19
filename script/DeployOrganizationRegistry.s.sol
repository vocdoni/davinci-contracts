// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import { Script, console } from "forge-std/Script.sol";
import { OrganizationRegistry } from "../src/OrganizationRegistry.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract DeployOrganizationRegistryScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address proxy = Upgrades.deployUUPSProxy(
            "OrganizationRegistry.sol",
            abi.encodeCall(OrganizationRegistry.initialize, ())
        );
        console.log("OrganizationRegistry deployed at:", proxy);
        vm.stopBroadcast();
    }
}
