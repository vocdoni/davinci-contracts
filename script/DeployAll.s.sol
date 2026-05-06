// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {StateTransitionVerifierGroth16} from "../src/verifiers/StateTransitionVerifierGroth16.sol";
import {ResultsVerifierGroth16} from "../src/verifiers/ResultsVerifierGroth16.sol";

contract DeployAllScript is Script {
    bytes32 private constant STATE_TRANSITION_VERIFIER_SALT =
        keccak256("davinci-contracts.state-transition-verifier-groth16");
    bytes32 private constant RESULTS_VERIFIER_SALT =
        keccak256("davinci-contracts.results-verifier-groth16");
    bytes32 private constant PROCESS_REGISTRY_SALT = keccak256("davinci-contracts.process-registry");

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deployer address:", deployerAddress);
        vm.startBroadcast(deployerPrivateKey);

        StateTransitionVerifierGroth16 stv =
            new StateTransitionVerifierGroth16{salt: STATE_TRANSITION_VERIFIER_SALT}();
        console.log("StateTransitionVerifierGroth16 deployed at:", address(stv));

        ResultsVerifierGroth16 rv = new ResultsVerifierGroth16{salt: RESULTS_VERIFIER_SALT}();
        console.log("ResultsVerifierGroth16 deployed at:", address(rv));

        ProcessRegistry processRegistry =
            new ProcessRegistry{salt: PROCESS_REGISTRY_SALT}(address(stv), address(rv));
        console.log("ProcessRegistry deployed at:", address(processRegistry));

        vm.stopBroadcast();
    }
}
