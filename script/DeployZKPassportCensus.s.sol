// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {ZKPassportCensus} from "../src/ZKPassportCensus.sol";
import {IProofVerifier} from "../src/verifiers/IProofVerifier.sol";
// OuterCount4 is NOT imported here — it uses assembly incompatible with
// via_ir=true (which ZKPassportCensus/InternalLeanIMT require).
// Deploy OuterCount4 separately (e.g. forge create from its own project)
// and pass its address via OUTER_COUNT4_ADDRESS.

/// @notice Deploy script for ZKPassportCensus.
///
/// Required env vars:
///   PRIVATE_KEY        — deployer private key
///   BACKEND_ADDRESS    — trusted backend address passed to ZKPassportCensus
///
/// Optional env vars:
///   OUTER_COUNT4_ADDRESS — address of a pre-deployed OuterCount4 (HonkVerifier).
///                          When set, the script skips deploying the verifier.
///                          When unset, the script deploys a new HonkVerifier.
///
/// Note: OuterCount4.sol contains a large Barretenberg UltraHonk verifier that uses
/// inline assembly without a memoryguard annotation.  This prevents the Yul optimizer
/// (required by via_ir=true) from spilling stack variables to memory, causing a
/// "stack too deep" error in the Foundry default profile.  The verifier must be
/// compiled separately (e.g. via `solc --optimize` without --via-ir) and deployed
/// independently; its address is then passed here via OUTER_COUNT4_ADDRESS.
contract DeployZKPassportCensus is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address backendAddress = vm.envAddress("BACKEND_ADDRESS");
        address outerCount4Address = vm.envOr("OUTER_COUNT4_ADDRESS", address(0));

        vm.startBroadcast(deployerPrivateKey);

        if (outerCount4Address == address(0)) {
            // Deploy verifier inline — only works if the project compiles without via_ir.
            // For normal Foundry profiles (via_ir=true), set OUTER_COUNT4_ADDRESS instead.
            revert("OUTER_COUNT4_ADDRESS not set. Deploy OuterCount4 first and set the env var.");
        }

        console.log("OuterCount4 at:", outerCount4Address);

        ZKPassportCensus zkPassportCensus = new ZKPassportCensus(
            IProofVerifier(outerCount4Address),
            backendAddress
        );
        console.log("ZKPassportCensus deployed at:", address(zkPassportCensus));

        vm.stopBroadcast();
    }
}
