// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {ZKPassportCensus} from "../src/ZKPassportCensus.sol";
import {IProofVerifier} from "../src/verifiers/IProofVerifier.sol";
// OuterCountN verifiers use assembly incompatible with via_ir=true.
// Deploy them separately and pass their address via OUTER_COUNT_ADDRESS.

/// @notice Deploy script for ZKPassportCensus.
///
/// Required env vars:
///   PRIVATE_KEY          — deployer private key
///   BACKEND_ADDRESS      — trusted backend address passed to ZKPassportCensus
///   OUTER_COUNT_ADDRESS  — address of a pre-deployed OuterCountN verifier
///
/// Optional env vars (default to outer_evm_count_5 values):
///   EXPECTED_PUBLIC_INPUTS  — total public input count (default 10 for OuterCount5)
///   NULLIFIER_INPUT_INDEX   — index of scoped_nullifier (default 8 = 10-2)
contract DeployZKPassportCensus is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address backendAddress = vm.envAddress("BACKEND_ADDRESS");
        address outerCountAddress = vm.envAddress("OUTER_COUNT_ADDRESS");
        uint256 expectedPublicInputs = vm.envOr("EXPECTED_PUBLIC_INPUTS", uint256(10));
        uint256 nullifierInputIndex = vm.envOr("NULLIFIER_INPUT_INDEX", uint256(8));

        vm.startBroadcast(deployerPrivateKey);

        console.log("Verifier at:", outerCountAddress);
        console.log("Expected public inputs:", expectedPublicInputs);
        console.log("Nullifier input index:", nullifierInputIndex);

        ZKPassportCensus zkPassportCensus = new ZKPassportCensus(
            IProofVerifier(outerCountAddress),
            backendAddress,
            expectedPublicInputs,
            nullifierInputIndex
        );
        console.log("ZKPassportCensus deployed at:", address(zkPassportCensus));

        vm.stopBroadcast();
    }
}
