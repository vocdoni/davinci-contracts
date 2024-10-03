// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.24;

import {Script, console} from "forge-std/Script.sol";
import {ProposalRegistry} from "../src/ProposalRegistry.sol";
import {Upgrades} from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract DeployProposalRegistryScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address organizationRegistryAddress = vm.envAddress("ORGANIZATION_REGISTRY");
        string memory chainID = vm.envString("CHAIN_ID");
        vm.startBroadcast(deployerPrivateKey);
        address proxy = Upgrades.deployUUPSProxy("ProposalRegistry.sol", abi.encodeCall(ProposalRegistry.initialize, (chainID, organizationRegistryAddress)));
        console.log("ProposalRegistry deployed at:", proxy);
        vm.stopBroadcast();
    }
}
