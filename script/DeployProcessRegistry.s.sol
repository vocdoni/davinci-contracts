// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import { Script, console } from "forge-std/Script.sol";
import { ProcessRegistry } from "../src/ProcessRegistry.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract DeployProcessRegistryScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address organizationRegistryAddress = vm.envAddress("ORGANIZATION_REGISTRY");
        string memory chainID = vm.envString("CHAIN_ID");
        vm.startBroadcast(deployerPrivateKey);
        address proxy = Upgrades.deployUUPSProxy(
            "ProcessRegistry.sol",
            abi.encodeCall(ProcessRegistry.initialize, (chainID, organizationRegistryAddress, address(0), address(0)))
        );
        console.log("ProcessRegistry deployed at:", proxy);
        vm.stopBroadcast();
    }
}
