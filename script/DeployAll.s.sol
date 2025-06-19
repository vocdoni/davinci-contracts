// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import { Script, console } from "forge-std/Script.sol";
import { OrganizationRegistry } from "../src/OrganizationRegistry.sol";
import { ProcessRegistry } from "../src/ProcessRegistry.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract DeployAllScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        address proxy = Upgrades.deployUUPSProxy(
            "OrganizationRegistry.sol:OrganizationRegistry",
            abi.encodeCall(OrganizationRegistry.initialize, ())
        );
        console.log("OrganizationRegistry deployed at:", proxy);

        StateTransitionVerifierGroth16 stv = new StateTransitionVerifierGroth16();
        console.log("StateTransitionVerifierGroth16 deployed at:", address(stv));

        ResultsVerifierGroth16 rv = new ResultsVerifierGroth16();
        console.log("ResultsVerifierGroth16 deployed at:", address(rv));

        uint256 chainId = vm.envUint("CHAIN_ID");
        address proxy2 = Upgrades.deployUUPSProxy(
            "ProcessRegistry:ProcessRegistry.sol",
            abi.encodeCall(ProcessRegistry.initialize, (uint32(chainId), address(stv), address(rv)))
        );
        console.log("ProcessRegistry deployed at:", proxy2);
        vm.stopBroadcast();
    }
}
