// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import { Script, console } from "forge-std/Script.sol";
import { ProcessRegistry } from "../src/ProcessRegistry.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract DeployProcessRegistryScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        uint256 chainID = vm.envUint("CHAIN_ID");
        address proxy = Upgrades.deployUUPSProxy(
            "ProcessRegistry.sol",
            abi.encodeCall(ProcessRegistry.initialize, (uint32(chainID), address(0), address(0))) // TODO: replace with actual verifier addresses
        );
        console.log("ProcessRegistry deployed at:", proxy);
        vm.stopBroadcast();
    }
}
