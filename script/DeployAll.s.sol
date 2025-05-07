// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import { Script, console } from "forge-std/Script.sol";
import { OrganizationRegistry } from "../src/OrganizationRegistry.sol";
import { ProcessRegistry } from "../src/ProcessRegistry.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract TestDeployAllScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("SEPOLIA_PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address proxy = Upgrades.deployUUPSProxy(
            "OrganizationRegistry.sol",
            abi.encodeCall(OrganizationRegistry.initialize, ())
        );
        console.log("OrganizationRegistry deployed at:", proxy);
        address proxy2 = Upgrades.deployUUPSProxy(
            "ProcessRegistry.sol",
            abi.encodeCall(ProcessRegistry.initialize, ("11155111", proxy, address(0)))
        );
        console.log("ProcessRegistry deployed at:", proxy2);
        vm.stopBroadcast();
    }
}
