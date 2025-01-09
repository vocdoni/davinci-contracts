// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.24;

import {Script, console} from "forge-std/Script.sol";
import {OrganizationRegistry} from "../src/OrganizationRegistry.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract TestDeployAllScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address proxy = Upgrades.deployUUPSProxy("OrganizationRegistry.sol", abi.encodeCall(OrganizationRegistry.initialize, ()));
        console.log("OrganizationRegistry deployed at:", proxy);
        address proxy2 = Upgrades.deployUUPSProxy("ProcessRegistry.sol", abi.encodeCall(ProcessRegistry.initialize, ("31337", proxy)));
        console.log("ProcessRegistry deployed at:", proxy2);
        vm.stopBroadcast();
    }
}
