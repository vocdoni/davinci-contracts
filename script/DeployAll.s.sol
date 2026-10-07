// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {ZiskVerifier} from "../src/verifiers/ZiskVerifier.sol";

/// @notice Deploys the ZisK PLONK verifier (ZiskVerifier extends PlonkVerifier) and the
///         ProcessRegistry pinned to the davinci-zkvm program vks.
/// @dev Env: PRIVATE_KEY, CHAIN_ID, BATCH_PROGRAM_VK, RESULTS_PROGRAM_VK, ROOT_C_VADCOP_FINAL,
///      BALLOT_VK_HASH (bytes32 hex). ROOT_C_VADCOP_FINAL must match the vendored verifier setup.
///      Optional: ZISK_VERIFIER (reuse a deployed verifier instead of deploying one; it must
///      report ROOT_C_VADCOP_FINAL, and its code hash is logged for the release pin),
///      DKG_MANAGER, COUNCIL_MANAGER, and in seconds the grace window GRACE_DEFAULT
///      (180), GRACE_FLOOR (150), GRACE_CEIL (600), GRACE_MAX_TOTAL (1800) and the shorten notice
///      NOTICE_MIN (60).
contract DeployAllScript is Script {
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deployer address:", deployerAddress);

        uint256 chainId = vm.envUint("CHAIN_ID");
        require(chainId <= type(uint32).max, "CHAIN_ID exceeds uint32");
        // forge-lint: disable-next-line(unsafe-typecast)
        uint32 chainId32 = uint32(chainId);
        bytes32 batchProgramVK = vm.envBytes32("BATCH_PROGRAM_VK");
        bytes32 resultsProgramVK = vm.envBytes32("RESULTS_PROGRAM_VK");
        bytes32 rootCVadcopFinal = vm.envBytes32("ROOT_C_VADCOP_FINAL");
        bytes32 ballotVKHash = vm.envBytes32("BALLOT_VK_HASH");
        // Optional deployed ZiskVerifier to reuse; zero (the default) deploys a new one.
        address existingVerifier = vm.envOr("ZISK_VERIFIER", address(0));
        // Optional davinci-dkg manager; zero (the default) disables the DKG key modes.
        address dkgManager = vm.envOr("DKG_MANAGER", address(0));
        // Optional Council manager; zero (the default) disables the COUNCIL key mode.
        address councilManager = vm.envOr("COUNCIL_MANAGER", address(0));
        // The registry checks 0 < floor <= default <= ceil <= maxTotal and noticeMin > 0.
        uint32 defaultGrace = _envSeconds("GRACE_DEFAULT", 180);
        uint32 graceFloor = _envSeconds("GRACE_FLOOR", 150);
        uint32 graceCeil = _envSeconds("GRACE_CEIL", 600);
        uint32 graceMaxTotal = _envSeconds("GRACE_MAX_TOTAL", 1800);
        uint32 noticeMin = _envSeconds("NOTICE_MIN", 60);

        vm.startBroadcast(deployerPrivateKey);

        ZiskVerifier zisk;
        if (existingVerifier == address(0)) {
            zisk = new ZiskVerifier();
            console.log("ZiskVerifier deployed at:", address(zisk));
        } else {
            require(existingVerifier.code.length > 0, "ZISK_VERIFIER has no code");
            zisk = ZiskVerifier(existingVerifier);
            console.log("ZiskVerifier reused at:", address(zisk));
        }
        console.log("ZiskVerifier code hash:", vm.toString(address(zisk).codehash));
        require(zisk.getRootCVadcopFinal() == rootCVadcopFinal, "ROOT_C_VADCOP_FINAL differs from the verifier setup");

        ProcessRegistry processRegistry = new ProcessRegistry(
            chainId32,
            address(zisk),
            batchProgramVK,
            resultsProgramVK,
            rootCVadcopFinal,
            ballotVKHash,
            dkgManager,
            councilManager,
            defaultGrace,
            graceFloor,
            graceCeil,
            graceMaxTotal,
            noticeMin
        );
        console.log("ProcessRegistry deployed at:", address(processRegistry));
        console.log("DavinciDKGAdapter deployed at:", processRegistry.dkgAdapter());
        console.log("CouncilAdapter deployed at:", processRegistry.councilAdapter());
        console.log(
            string.concat(
                "Grace (s): default ",
                vm.toString(uint256(defaultGrace)),
                ", floor ",
                vm.toString(uint256(graceFloor)),
                ", ceil ",
                vm.toString(uint256(graceCeil)),
                ", max total ",
                vm.toString(uint256(graceMaxTotal)),
                "; notice min ",
                vm.toString(uint256(noticeMin))
            )
        );

        vm.stopBroadcast();
    }

    /// @dev A seconds value from the environment, or def when unset.
    function _envSeconds(string memory name, uint256 def) internal view returns (uint32) {
        uint256 v = vm.envOr(name, def);
        require(v <= type(uint32).max, string.concat(name, " exceeds uint32"));
        // forge-lint: disable-next-line(unsafe-typecast)
        return uint32(v);
    }
}
