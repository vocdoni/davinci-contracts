// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {stdJson} from "forge-std/StdJson.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";
import {PublicsLib} from "../src/libraries/PublicsLib.sol";
import {MockCensusValidator, BadCensusValidator} from "./mocks/MockCensusValidator.sol";
import {MockZiskVerifier} from "./mocks/MockZiskVerifier.sol";

/// @dev Dynamic Merkle censuses: origin 2 (organizer-updated root) and origin 3 (root
///      history of an on-chain census contract). The transitions are the fixture's first
///      batch rebuilt on the origin-2 and origin-3 genesis (test/vectors `.dynamic`).
contract DynamicCensusTest is RegistryTestBase {
    using stdJson for string;

    uint256 internal constant CREATION_BLOCK = 100;
    bytes32 internal constant OTHER_ROOT = bytes32(uint256(0x0ddba11));

    MockCensusValidator internal census;
    bytes32 internal root;

    function setUp() public override {
        super.setUp();
        census = new MockCensusValidator();
        root = fixture.readBytes32(".census_root");
        vm.roll(CREATION_BLOCK);
    }

    // --- helpers ---------------------------------------------------------------

    function _offchainDynamic(bytes32 r) internal pure returns (DAVINCITypes.Census memory c) {
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1;
        c.censusRoot = r;
        c.censusURI = "https://example.com/census-v1.jsonl";
    }

    function _onchain(address addr) internal pure returns (DAVINCITypes.Census memory c) {
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1;
        c.contractAddress = addr;
        c.censusURI = "https://example.com/census-indexer";
    }

    function _create(DAVINCITypes.Census memory c) internal returns (bytes31) {
        return _newProcess(block.timestamp, DURATION, MAX_VOTERS, _ballotMode(), c);
    }

    function _setCensus(bytes31 pid, DAVINCITypes.Census memory c) internal {
        vm.prank(ORGANIZER);
        registry.setProcessCensus(pid, c);
    }

    /// @dev The fixture's first batch on the genesis of a process with this origin.
    function _dynamic(DAVINCITypes.CensusOrigin origin) internal view returns (bytes32 genesis, Transition memory t) {
        string memory p = string.concat(".dynamic[", vm.toString(uint256(origin) - 2), "]");
        assertEq(fixture.readUint(string.concat(p, ".census_origin")), uint256(origin), "fixture origin");
        genesis = fixture.readBytes32(string.concat(p, ".genesis_root"));
        t = _transitionAt(string.concat(p, ".transition"));
    }

    function _offchainTransition() internal view returns (Transition memory t) {
        (, t) = _dynamic(DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_DYNAMIC_V1);
    }

    function _onchainTransition() internal view returns (Transition memory t) {
        (, t) = _dynamic(DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1);
    }

    /// @dev Publics with the census root register set to r (BE integer). Neither the blob
    ///      digest nor the opening points cover it, so the openings still verify.
    function _withCensusRoot(Transition memory t, bytes32 r) internal pure returns (Transition memory) {
        _setReg32(t.publicValues, 20, PublicsLib.reverse32(r));
        return t;
    }

    function _assertSettled(bytes31 pid, Transition memory t) internal view {
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.latestStateRoot, t.rootAfter);
        assertEq(p.batchNumber, 1);
        assertEq(p.votersCount, 3);
    }

    // --- origin 2: creation ----------------------------------------------------

    function test_Offchain_CreateStoresRootAndGenesis() public {
        DAVINCITypes.Census memory c = _offchainDynamic(root);
        (bytes32 genesis,) = _dynamic(c.censusOrigin);
        bytes31 pid = _create(c);
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(uint256(p.census.censusOrigin), 2);
        assertEq(p.census.censusRoot, root);
        assertEq(p.latestStateRoot, genesis, "go-sdk genesis");
        assertEq(p.latestStateRoot, registry.genesisRoot(pid, _ballotMode(), _encKey(), c.censusOrigin));
    }

    function test_Offchain_RevertWhen_CreateWithZeroRoot() public {
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _create(_offchainDynamic(bytes32(0)));
    }

    function test_Offchain_RevertWhen_CreateWithEmptyURI() public {
        DAVINCITypes.Census memory c = _offchainDynamic(root);
        c.censusURI = "";
        vm.expectRevert(IProcessRegistry.InvalidCensusURI.selector);
        _create(c);
    }

    // --- origin 2: setProcessCensus --------------------------------------------

    function test_SetProcessCensus_UpdatesRootAndURI() public {
        bytes31 pid = _create(_offchainDynamic(root));
        bytes32 genesis = registry.getProcess(pid).latestStateRoot;
        DAVINCITypes.Census memory c = _offchainDynamic(OTHER_ROOT);
        c.censusURI = "https://example.com/census-v2.jsonl";

        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.CensusUpdated(pid, OTHER_ROOT, c.censusURI);
        vm.prank(ORGANIZER);
        uint256 g = gasleft();
        registry.setProcessCensus(pid, c);
        emit log_named_uint("setProcessCensus gas:", g - gasleft());

        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.census.censusRoot, OTHER_ROOT);
        assertEq(p.census.censusURI, c.censusURI);
        assertEq(uint256(p.census.censusOrigin), 2);
        assertEq(p.latestStateRoot, genesis);
    }

    function test_SetProcessCensus_WhilePaused() public {
        bytes31 pid = _create(_offchainDynamic(root));
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        _setCensus(pid, _offchainDynamic(OTHER_ROOT));
        assertEq(registry.getProcess(pid).census.censusRoot, OTHER_ROOT);
    }

    function test_SetProcessCensus_RevertWhen_NotOrganizer() public {
        bytes31 pid = _create(_offchainDynamic(root));
        vm.prank(address(0xBEEF));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        registry.setProcessCensus(pid, _offchainDynamic(OTHER_ROOT));
    }

    function test_SetProcessCensus_RevertWhen_OriginNotUpdatable() public {
        bytes31 static_ = _create(_census());
        DAVINCITypes.Census memory csp = _census();
        csp.censusOrigin = DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1;
        csp.censusRoot = bytes32(uint256(uint160(address(0xC5C5))));
        bytes31 cspPid = _create(csp);
        bytes31 onchain = _create(_onchain(address(census)));

        bytes31[3] memory pids = [static_, cspPid, onchain];
        for (uint256 i = 0; i < pids.length; i++) {
            DAVINCITypes.Census memory c = registry.getProcess(pids[i]).census;
            c.censusRoot = OTHER_ROOT;
            vm.prank(ORGANIZER);
            vm.expectRevert(IProcessRegistry.CensusNotUpdatable.selector);
            registry.setProcessCensus(pids[i], c);
        }
    }

    function test_SetProcessCensus_RevertWhen_OriginChanges() public {
        bytes31 pid = _create(_offchainDynamic(root));
        DAVINCITypes.Census memory c = _offchainDynamic(OTHER_ROOT);
        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_OFFCHAIN_STATIC_V1;
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidCensusOrigin.selector);
        registry.setProcessCensus(pid, c);

        c.censusOrigin = DAVINCITypes.CensusOrigin.MERKLE_TREE_ONCHAIN_DYNAMIC_V1;
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidCensusOrigin.selector);
        registry.setProcessCensus(pid, c);
    }

    function test_SetProcessCensus_RevertWhen_ContractAddress() public {
        bytes31 pid = _create(_offchainDynamic(root));
        DAVINCITypes.Census memory c = _offchainDynamic(OTHER_ROOT);
        c.contractAddress = address(census);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidCensusAddress.selector);
        registry.setProcessCensus(pid, c);
    }

    function test_SetProcessCensus_RevertWhen_ZeroRoot() public {
        bytes31 pid = _create(_offchainDynamic(root));
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        registry.setProcessCensus(pid, _offchainDynamic(bytes32(0)));
    }

    function test_SetProcessCensus_RevertWhen_EmptyURI() public {
        bytes31 pid = _create(_offchainDynamic(root));
        DAVINCITypes.Census memory c = _offchainDynamic(OTHER_ROOT);
        c.censusURI = "";
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidCensusURI.selector);
        registry.setProcessCensus(pid, c);
    }

    function test_SetProcessCensus_RevertWhen_Ended() public {
        bytes31 pid = _create(_offchainDynamic(root));
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.ENDED);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessCensus(pid, _offchainDynamic(OTHER_ROOT));
    }

    function test_SetProcessCensus_RevertWhen_Canceled() public {
        bytes31 pid = _create(_offchainDynamic(root));
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        registry.setProcessCensus(pid, _offchainDynamic(OTHER_ROOT));
    }

    function test_SetProcessCensus_RevertWhen_PastEndTime() public {
        bytes31 pid = _create(_offchainDynamic(root));
        vm.warp(block.timestamp + DURATION);
        vm.prank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        registry.setProcessCensus(pid, _offchainDynamic(OTHER_ROOT));
    }

    function test_SetProcessCensus_RevertWhen_UnknownProcess() public {
        DAVINCITypes.Census memory c = _offchainDynamic(OTHER_ROOT);
        vm.startPrank(ORGANIZER);
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        registry.setProcessCensus(bytes31(0), c);
        vm.expectRevert(IProcessRegistry.UnknownProcessIdPrefix.selector);
        registry.setProcessCensus(bytes31(uint248(1)), c);
        vm.expectRevert(IProcessRegistry.ProcessNotFound.selector);
        registry.setProcessCensus(bytes31(fixture.readBytes(".process_id")), c);
        vm.stopPrank();
    }

    // --- origin 2: settlement --------------------------------------------------

    function test_Offchain_SettlesWithUpdatedRoot() public {
        bytes31 pid = _create(_offchainDynamic(OTHER_ROOT));
        Transition memory t = _offchainTransition();
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, t);

        _setCensus(pid, _offchainDynamic(root));
        _submit(pid, t);
        _assertSettled(pid, t);
    }

    function test_Offchain_RevertWhen_StaleRoot() public {
        bytes31 pid = _create(_offchainDynamic(root));
        _setCensus(pid, _offchainDynamic(OTHER_ROOT));
        Transition memory t = _offchainTransition();
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, t);

        // The same batch proven against the current root settles.
        _submit(pid, _withCensusRoot(t, OTHER_ROOT));
        _assertSettled(pid, t);
    }

    // --- origin 3: creation ----------------------------------------------------

    function test_Onchain_CreateStoresContractRoot() public {
        census.setRoot(uint256(root));
        DAVINCITypes.Census memory c = _onchain(address(census));
        c.censusRoot = OTHER_ROOT; // informational only: replaced by the contract's root
        (bytes32 genesis,) = _dynamic(c.censusOrigin);
        bytes31 pid = _create(c);
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(uint256(p.census.censusOrigin), 3);
        assertEq(p.census.contractAddress, address(census));
        assertEq(p.census.censusRoot, root);
        assertEq(p.creationBlock, CREATION_BLOCK);
        assertEq(p.latestStateRoot, genesis, "go-sdk genesis");
        assertEq(p.latestStateRoot, registry.genesisRoot(pid, _ballotMode(), _encKey(), c.censusOrigin));
    }

    function test_Onchain_CreateWithEmptyCensus() public {
        bytes31 pid = _create(_onchain(address(census)));
        assertEq(registry.getProcess(pid).census.censusRoot, bytes32(0));
    }

    function test_Onchain_RevertWhen_NoCode() public {
        vm.expectRevert(IProcessRegistry.InvalidCensusAddress.selector);
        _create(_onchain(address(0)));
        vm.expectRevert(IProcessRegistry.InvalidCensusAddress.selector);
        _create(_onchain(address(0xBEEF)));
    }

    function test_Onchain_RevertWhen_NotACensusContract() public {
        address notACensus = address(new MockZiskVerifier());
        vm.expectRevert(IProcessRegistry.InvalidCensusAddress.selector);
        _create(_onchain(notACensus));
    }

    function test_Onchain_RevertWhen_AllowAnyValidRoot() public {
        DAVINCITypes.Census memory c = _onchain(address(census));
        c.onchainAllowAnyValidRoot = true;
        vm.expectRevert(IProcessRegistry.InvalidCensusConfig.selector);
        _create(c);
    }

    function test_RevertWhen_ContractAddressOnOtherOrigins() public {
        DAVINCITypes.Census memory c = _offchainDynamic(root);
        c.contractAddress = address(census);
        vm.expectRevert(IProcessRegistry.InvalidCensusAddress.selector);
        _create(c);

        c = _census();
        c.contractAddress = address(census);
        vm.expectRevert(IProcessRegistry.InvalidCensusAddress.selector);
        _create(c);

        c.censusOrigin = DAVINCITypes.CensusOrigin.CSP_EDDSA_BABYJUBJUB_V1;
        c.censusRoot = bytes32(uint256(uint160(address(0xC5C5))));
        vm.expectRevert(IProcessRegistry.InvalidCensusAddress.selector);
        _create(c);
    }

    function test_Onchain_RevertWhen_CensusRootCallFails() public {
        BadCensusValidator bad = new BadCensusValidator();
        BadCensusValidator.Mode[4] memory modes = [
            BadCensusValidator.Mode.Revert,
            BadCensusValidator.Mode.Short,
            BadCensusValidator.Mode.Empty,
            BadCensusValidator.Mode.Burn
        ];
        for (uint256 i = 0; i < modes.length; i++) {
            bad.setModes(BadCensusValidator.Mode.Ok, modes[i]);
            DAVINCITypes.Census memory c = _onchain(address(bad));
            DAVINCITypes.BallotMode memory m = _ballotMode();
            DAVINCITypes.EncryptionKey memory k = _encKey();
            vm.prank(ORGANIZER);
            vm.expectRevert(IProcessRegistry.InvalidCensusAddress.selector);
            registry.newProcess{gas: 5_000_000}(
                DAVINCITypes.ProcessStatus.READY, block.timestamp, DURATION, MAX_VOTERS, m, c, "", k, _noDkg()
            );
        }
    }

    function test_Onchain_RevertWhen_EmptyURI() public {
        DAVINCITypes.Census memory c = _onchain(address(census));
        c.censusURI = "";
        vm.expectRevert(IProcessRegistry.InvalidCensusURI.selector);
        _create(c);
    }

    // --- origin 3: settlement --------------------------------------------------

    function test_Onchain_SettlesWithCurrentRoot() public {
        census.setRoot(uint256(root));
        bytes31 pid = _create(_onchain(address(census)));
        vm.roll(CREATION_BLOCK + 10);
        Transition memory t = _onchainTransition();
        vm.blobhashes(t.versionedHashes);
        uint256 g = gasleft();
        registry.submitStateTransition(pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs);
        emit log_named_uint("submitStateTransition gas, origin 3, 1 blob (mock verifier):", g - gasleft());
        _assertSettled(pid, t);
    }

    function test_Onchain_SettlesWithRootAddedAfterCreation() public {
        bytes31 pid = _create(_onchain(address(census)));
        vm.roll(CREATION_BLOCK + 5);
        census.setRoot(uint256(root));
        vm.roll(CREATION_BLOCK + 10);
        Transition memory t = _onchainTransition();
        _submit(pid, t);
        _assertSettled(pid, t);
    }

    function test_Onchain_SettlesWithHistoricalRoot() public {
        census.setRoot(uint256(root));
        bytes31 pid = _create(_onchain(address(census)));
        vm.roll(CREATION_BLOCK + 5);
        census.setRoot(uint256(OTHER_ROOT));
        vm.roll(CREATION_BLOCK + 10);
        assertEq(census.getRootBlockNumber(uint256(root)), CREATION_BLOCK + 5);
        Transition memory t = _onchainTransition();
        _submit(pid, t);
        _assertSettled(pid, t);
    }

    function test_Onchain_SettlesWithRootReplacedInCreationBlock() public {
        // rbn == creationBlock: the root was current in the creation block.
        census.setRoot(uint256(root));
        bytes31 pid = _create(_onchain(address(census)));
        census.setRoot(uint256(OTHER_ROOT));
        vm.roll(CREATION_BLOCK + 10);
        assertEq(census.getRootBlockNumber(uint256(root)), CREATION_BLOCK);
        Transition memory t = _onchainTransition();
        _submit(pid, t);
        _assertSettled(pid, t);
    }

    /// @dev An unknown root answers rbn = 0 and must not settle.
    function test_Onchain_RevertWhen_UnknownRoot() public {
        census.setRoot(uint256(OTHER_ROOT));
        bytes31 pid = _create(_onchain(address(census)));
        vm.roll(CREATION_BLOCK + 10);
        assertEq(census.getRootBlockNumber(uint256(root)), 0);
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, _onchainTransition());
    }

    /// @dev With creationBlock == 0, rbn < creationBlock can never fire: only the rbn == 0
    ///      guard rejects the unknown root.
    function test_Onchain_RevertWhen_UnknownRootCreatedAtBlockZero() public {
        vm.roll(0);
        bytes31 pid = _create(_onchain(address(census)));
        assertEq(registry.getProcess(pid).creationBlock, 0);
        vm.roll(10);
        assertEq(census.getRootBlockNumber(uint256(root)), 0);
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, _onchainTransition());
    }

    function test_Onchain_RevertWhen_RootBlockCallFails() public {
        BadCensusValidator bad = new BadCensusValidator();
        bytes31 pid = _create(_onchain(address(bad)));
        vm.roll(CREATION_BLOCK + 10);
        BadCensusValidator.Mode[3] memory modes =
            [BadCensusValidator.Mode.Revert, BadCensusValidator.Mode.Short, BadCensusValidator.Mode.Empty];
        for (uint256 i = 0; i < modes.length; i++) {
            bad.setModes(modes[i], BadCensusValidator.Mode.Ok);
            vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
            _submit(pid, _onchainTransition());
        }

        // The same contract answering block.number settles.
        bad.setModes(BadCensusValidator.Mode.Ok, BadCensusValidator.Mode.Ok);
        Transition memory t = _onchainTransition();
        _submit(pid, t);
        _assertSettled(pid, t);
    }

    function test_Onchain_RootBlockCallGasIsCapped() public {
        BadCensusValidator bad = new BadCensusValidator();
        bytes31 pid = _create(_onchain(address(bad)));
        bad.setModes(BadCensusValidator.Mode.Burn, BadCensusValidator.Mode.Ok);
        vm.roll(CREATION_BLOCK + 10);
        Transition memory t = _onchainTransition();
        vm.blobhashes(t.versionedHashes);
        uint256 g = gasleft();
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        registry.submitStateTransition{gas: 5_000_000}(
            pid, t.publicValues, t.proofBytes, t.commitments, t.ys, t.kzgProofs
        );
        uint256 used = g - gasleft();
        emit log_named_uint("submitStateTransition gas, census call burns its cap:", used);
        assertLt(used, 250_000, "census call not capped");
    }

    function test_Onchain_RevertWhen_UnknownRootOnEmptyCensus() public {
        bytes31 pid = _create(_onchain(address(census)));
        vm.roll(CREATION_BLOCK + 10);
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, _onchainTransition());
    }

    function test_Onchain_RevertWhen_ContractAnswersZero() public {
        // rbn = 0 is rejected even for the contract's own current root.
        census.setRoot(uint256(root));
        census.force(uint256(root), 0);
        bytes31 pid = _create(_onchain(address(census)));
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, _onchainTransition());
    }

    function test_Onchain_RevertWhen_RootReplacedBeforeCreation() public {
        vm.roll(CREATION_BLOCK - 20);
        census.setRoot(uint256(root));
        vm.roll(CREATION_BLOCK - 10);
        census.setRoot(uint256(OTHER_ROOT));
        vm.roll(CREATION_BLOCK);
        bytes31 pid = _create(_onchain(address(census)));
        vm.roll(CREATION_BLOCK + 10);
        assertEq(census.getRootBlockNumber(uint256(root)), CREATION_BLOCK - 10);
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, _onchainTransition());
    }

    function test_Onchain_RevertWhen_RootBlockInFuture() public {
        bytes31 pid = _create(_onchain(address(census)));
        vm.roll(CREATION_BLOCK + 10);
        census.force(uint256(root), block.number + 1);
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, _onchainTransition());

        // Boundary: rbn == block.number settles.
        census.force(uint256(root), block.number);
        Transition memory t = _onchainTransition();
        _submit(pid, t);
        _assertSettled(pid, t);
    }

    function test_Onchain_RevertWhen_CensusRootByteOrderSwapped() public {
        // The guest publishes LE limbs; the census contract is asked for the BE integer.
        census.setRoot(uint256(PublicsLib.reverse32(root)));
        bytes31 pid = _create(_onchain(address(census)));
        vm.roll(CREATION_BLOCK + 10);
        vm.expectRevert(IProcessRegistry.InvalidCensusRoot.selector);
        _submit(pid, _onchainTransition());
    }

    function test_Onchain_UsesBatchRootNotStoredRoot() public {
        // The root stored at creation plays no part in settlement.
        census.setRoot(uint256(root));
        bytes31 pid = _create(_onchain(address(census)));
        vm.roll(CREATION_BLOCK + 5);
        census.setRoot(uint256(OTHER_ROOT));
        vm.roll(CREATION_BLOCK + 10);
        Transition memory t = _withCensusRoot(_onchainTransition(), OTHER_ROOT);
        _submit(pid, t);
        _assertSettled(pid, t);
        assertEq(registry.getProcess(pid).census.censusRoot, root);
    }
}
