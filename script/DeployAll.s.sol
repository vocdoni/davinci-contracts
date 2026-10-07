// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {ProcessRegistry} from "../src/ProcessRegistry.sol";
import {ZiskVerifier} from "../src/verifiers/ZiskVerifier.sol";
import {IDKGManager} from "../src/interfaces/dkg/IDKGManager.sol";

/// @notice Deploys the ZisK PLONK verifier (ZiskVerifier extends PlonkVerifier) and the
///         ProcessRegistry pinned to the davinci-zkvm program vks.
/// @dev Env: PRIVATE_KEY, CHAIN_ID, BATCH_PROGRAM_VK, RESULTS_PROGRAM_VK, ROOT_C_VADCOP_FINAL,
///      BALLOT_VK_HASH (bytes32 hex). ROOT_C_VADCOP_FINAL must match the vendored verifier setup.
///      Optional: ZISK_VERIFIER (reuse a deployed verifier instead of deploying one; it must
///      report ROOT_C_VADCOP_FINAL, and its code hash is logged for the release pin),
///      DKG_MANAGER, COUNCIL_MANAGER, and in seconds the grace window GRACE_DEFAULT
///      (180), GRACE_FLOOR (150), GRACE_CEIL (600), GRACE_MAX_TOTAL (1800) and the shorten notice
///      NOTICE_MIN (60).
///      PINS_FROM_REGISTRY, a live ProcessRegistry on the same chain, replaces all of the
///      optional verifier, pin and grace variables above: the new registry reuses its verifier
///      and copies its pins and grace settings. Any of them also set in the environment must
///      equal the inherited value, so a stale .env aborts the deployment instead of winning.
contract DeployAllScript is Script {
    struct Config {
        uint32 chainId;
        address verifier; // zero deploys a new ZiskVerifier
        bytes32 batchProgramVK;
        bytes32 resultsProgramVK;
        bytes32 rootCVadcopFinal;
        bytes32 ballotVKHash;
        address dkgManager;
        address councilManager;
        uint32 defaultGrace;
        uint32 graceFloor;
        uint32 graceCeil;
        uint32 graceMaxTotal;
        uint32 noticeMin;
    }

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployerAddress = vm.addr(deployerPrivateKey);
        console.log("Deployer address:", deployerAddress);

        Config memory c = _config();
        _checkManagers(c);

        vm.startBroadcast(deployerPrivateKey);

        ZiskVerifier zisk;
        if (c.verifier == address(0)) {
            zisk = new ZiskVerifier();
            console.log("ZiskVerifier deployed at:", address(zisk));
        } else {
            require(c.verifier.code.length > 0, "ZISK_VERIFIER has no code");
            zisk = ZiskVerifier(c.verifier);
            console.log("ZiskVerifier reused at:", address(zisk));
        }
        console.log("ZiskVerifier code hash:", vm.toString(address(zisk).codehash));
        require(zisk.getRootCVadcopFinal() == c.rootCVadcopFinal, "ROOT_C_VADCOP_FINAL differs from the verifier setup");

        ProcessRegistry processRegistry = new ProcessRegistry(
            c.chainId,
            address(zisk),
            c.batchProgramVK,
            c.resultsProgramVK,
            c.rootCVadcopFinal,
            c.ballotVKHash,
            c.dkgManager,
            c.councilManager,
            c.defaultGrace,
            c.graceFloor,
            c.graceCeil,
            c.graceMaxTotal,
            c.noticeMin
        );
        console.log("ProcessRegistry deployed at:", address(processRegistry));
        console.log("DavinciDKGAdapter deployed at:", processRegistry.dkgAdapter());
        console.log("  DKG manager:", c.dkgManager);
        console.log("CouncilAdapter deployed at:", processRegistry.councilAdapter());
        console.log("  Council manager:", c.councilManager);
        console.log("Batch program vk:", vm.toString(c.batchProgramVK));
        console.log("Results program vk:", vm.toString(c.resultsProgramVK));
        console.log("rootCVadcopFinal:", vm.toString(c.rootCVadcopFinal));
        console.log("Ballot vk hash:", vm.toString(c.ballotVKHash));
        console.log(
            string.concat(
                "Grace (s): default ",
                vm.toString(uint256(c.defaultGrace)),
                ", floor ",
                vm.toString(uint256(c.graceFloor)),
                ", ceil ",
                vm.toString(uint256(c.graceCeil)),
                ", max total ",
                vm.toString(uint256(c.graceMaxTotal)),
                "; notice min ",
                vm.toString(uint256(c.noticeMin))
            )
        );

        vm.stopBroadcast();
    }

    function _config() internal view returns (Config memory c) {
        uint256 chainId = vm.envUint("CHAIN_ID");
        require(chainId <= type(uint32).max, "CHAIN_ID exceeds uint32");
        // forge-lint: disable-next-line(unsafe-typecast)
        c.chainId = uint32(chainId);
        // Optional davinci-dkg manager; zero (the default) disables the DKG key modes.
        c.dkgManager = _envAddress("DKG_MANAGER");
        // Optional Council manager; zero (the default) disables the COUNCIL key mode.
        c.councilManager = _envAddress("COUNCIL_MANAGER");

        ProcessRegistry source = ProcessRegistry(_envAddress("PINS_FROM_REGISTRY"));
        if (address(source) == address(0)) {
            // Optional deployed ZiskVerifier to reuse; zero (the default) deploys a new one.
            c.verifier = _envAddress("ZISK_VERIFIER");
            c.batchProgramVK = vm.envBytes32("BATCH_PROGRAM_VK");
            c.resultsProgramVK = vm.envBytes32("RESULTS_PROGRAM_VK");
            c.rootCVadcopFinal = vm.envBytes32("ROOT_C_VADCOP_FINAL");
            c.ballotVKHash = vm.envBytes32("BALLOT_VK_HASH");
            // The registry checks 0 < floor <= default <= ceil <= maxTotal and noticeMin > 0.
            c.defaultGrace = _envSeconds("GRACE_DEFAULT", 180);
            c.graceFloor = _envSeconds("GRACE_FLOOR", 150);
            c.graceCeil = _envSeconds("GRACE_CEIL", 600);
            c.graceMaxTotal = _envSeconds("GRACE_MAX_TOTAL", 1800);
            c.noticeMin = _envSeconds("NOTICE_MIN", 60);
            return c;
        }

        require(address(source).code.length > 0, "PINS_FROM_REGISTRY has no code");
        require(source.chainID() == c.chainId, "PINS_FROM_REGISTRY is on another CHAIN_ID");
        console.log("Pins and verifier inherited from registry:", address(source));
        c.verifier = _inherit("ZISK_VERIFIER", address(source.ziskVerifier()));
        c.batchProgramVK = _inherit("BATCH_PROGRAM_VK", source.batchProgramVK());
        c.resultsProgramVK = _inherit("RESULTS_PROGRAM_VK", source.resultsProgramVK());
        c.rootCVadcopFinal = _inherit("ROOT_C_VADCOP_FINAL", source.rootCVadcopFinal());
        c.ballotVKHash = _inherit("BALLOT_VK_HASH", source.ballotVKHash());
        c.defaultGrace = _inherit("GRACE_DEFAULT", source.defaultGrace());
        c.graceFloor = _inherit("GRACE_FLOOR", source.graceFloor());
        c.graceCeil = _inherit("GRACE_CEIL", source.graceCeil());
        c.graceMaxTotal = _inherit("GRACE_MAX_TOTAL", source.graceMaxTotal());
        c.noticeMin = _inherit("NOTICE_MIN", source.noticeMin());
    }

    /// @dev Both managers must be deployed contracts; the DKG adapter also reads
    ///      appManager() from its manager at construction.
    function _checkManagers(Config memory c) internal view {
        if (c.dkgManager != address(0)) {
            require(c.dkgManager.code.length > 0, "DKG_MANAGER has no code");
            require(IDKGManager(c.dkgManager).appManager().code.length > 0, "DKG_MANAGER has no app manager");
        }
        if (c.councilManager != address(0)) {
            require(c.councilManager.code.length > 0, "COUNCIL_MANAGER has no code");
        }
        require(
            c.dkgManager == address(0) || c.dkgManager != c.councilManager, "DKG_MANAGER and COUNCIL_MANAGER are equal"
        );
    }

    /// @dev An optional address: unset or empty is zero, anything else must parse, so a typo or a
    ///      placeholder aborts the deployment instead of silently disabling what it names.
    function _envAddress(string memory name) internal view returns (address) {
        if (!vm.envExists(name) || bytes(vm.envString(name)).length == 0) return address(0);
        return vm.envAddress(name);
    }

    /// @dev A seconds value from the environment, or def when unset.
    function _envSeconds(string memory name, uint256 def) internal view returns (uint32) {
        uint256 v = vm.envOr(name, def);
        require(v <= type(uint32).max, string.concat(name, " exceeds uint32"));
        // forge-lint: disable-next-line(unsafe-typecast)
        return uint32(v);
    }

    /// @dev The inherited value; name, when set in the environment, must equal it.
    function _inherit(string memory name, bytes32 inherited) internal view returns (bytes32) {
        require(vm.envOr(name, inherited) == inherited, string.concat(name, " differs from PINS_FROM_REGISTRY"));
        return inherited;
    }

    function _inherit(string memory name, address inherited) internal view returns (address) {
        address v = _envAddress(name);
        require(v == address(0) || v == inherited, string.concat(name, " differs from PINS_FROM_REGISTRY"));
        return inherited;
    }

    function _inherit(string memory name, uint32 inherited) internal view returns (uint32) {
        require(
            vm.envOr(name, uint256(inherited)) == inherited, string.concat(name, " differs from PINS_FROM_REGISTRY")
        );
        return inherited;
    }
}
