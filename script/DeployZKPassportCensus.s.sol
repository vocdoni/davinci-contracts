// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {ZKPassportCensus} from "../src/ZKPassportCensus.sol";
import {IZKPassportVerifier} from "../src/interfaces/IZKPassportVerifier.sol";

/// @notice Deploy ZKPassportCensus.
///
/// Required env vars:
///   PRIVATE_KEY  — deployer private key
///
/// Optional env vars:
///   VERIFIER_ADDRESS  — ZKPassport RootVerifier (default: 0x1D000001000EFD9a6371f4d90bB8920D5431c0D8)
///   SCOPE             — service scope string (default: "vocdoni")
contract DeployZKPassportCensus is Script {
    address constant DEFAULT_VERIFIER = 0x1D000001000EFD9a6371f4d90bB8920D5431c0D8;

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address verifierAddress = vm.envOr("VERIFIER_ADDRESS", DEFAULT_VERIFIER);
        string memory scope = vm.envOr("SCOPE", string("vocdoni"));

        vm.startBroadcast(deployerPrivateKey);

        console.log("Verifier:", verifierAddress);
        console.log("Scope:", scope);

        ZKPassportCensus census = new ZKPassportCensus(
            IZKPassportVerifier(verifierAddress),
            scope
        );
        console.log("ZKPassportCensus deployed at:", address(census));

        vm.stopBroadcast();
    }
}
