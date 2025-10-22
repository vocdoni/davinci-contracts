// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Script, console } from "forge-std/Script.sol";
import { OrganizationRegistry } from "../src/OrganizationRegistry.sol";
import { ProcessRegistry } from "../src/ProcessRegistry.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";

contract DeployAllScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deployer address:", deployerAddress);
        vm.startBroadcast(deployerPrivateKey);

        OrganizationRegistry organizationRegistry = new OrganizationRegistry();
        console.log("OrganizationRegistry deployed at:", address(organizationRegistry));

        StateTransitionVerifierGroth16 stv = new StateTransitionVerifierGroth16();
        console.log("StateTransitionVerifierGroth16 deployed at:", address(stv));

        ResultsVerifierGroth16 rv = new ResultsVerifierGroth16();
        console.log("ResultsVerifierGroth16 deployed at:", address(rv));

        uint256 chainId = vm.envUint("CHAIN_ID");
        bool blobs = vm.envBool("ACTIVATE_BLOBS");
        ProcessRegistry processRegistry = new ProcessRegistry(
            uint32(chainId), // change this to the desired chain ID
            address(stv),
            address(rv),
            bool(blobs)
        );
        console.log("ProcessRegistry deployed at:", address(processRegistry));

        vm.stopBroadcast();
    }
}
