// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Vm} from "forge-std/Vm.sol";
import {RegistryTestBase} from "./RegistryTestBase.sol";
import {IProcessRegistry} from "../src/interfaces/IProcessRegistry.sol";
import {DAVINCITypes} from "../src/libraries/DAVINCITypes.sol";

/// @dev The metadata URI and hash: set at creation, replaceable by the organizer until the
///      end, and every value on record in ProcessMetadataUpdated.
contract ProcessMetadataTest is RegistryTestBase {
    string internal constant URI_2 = "https://example.com/metadata-2.json";
    /// @dev sha256("metadata document, second version")
    bytes32 internal constant HASH_2 = 0xfb80b2c63e260040028879ea4b7e0fdb3ed57cec8d7913617ee3a687898a08bc;
    string internal constant URI_3 = "ipfs://metadata-3";
    /// @dev sha256("metadata document, third version")
    bytes32 internal constant HASH_3 = 0x7f1803b2473620648ca51e4d6f411a46b56d361da7077619328a3f3778a8c7f2;

    function test_MetadataHashConstants() public pure {
        assertEq(METADATA_HASH, sha256("metadata document"));
        assertEq(HASH_2, sha256("metadata document, second version"));
        assertEq(HASH_3, sha256("metadata document, third version"));
    }

    function _create(string memory uri, bytes32 hash) internal returns (bytes31) {
        vm.prank(ORGANIZER);
        return registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            0,
            DURATION,
            MAX_VOTERS,
            _ballotMode(),
            _census(),
            uri,
            hash,
            _encKey(),
            _noDkg()
        );
    }

    function _set(bytes31 pid, string memory uri, bytes32 hash) internal {
        vm.prank(ORGANIZER);
        registry.setProcessMetadata(pid, uri, hash);
    }

    function _assertMetadata(bytes31 pid, string memory uri, bytes32 hash) internal view {
        DAVINCITypes.Process memory p = registry.getProcess(pid);
        assertEq(p.metadataURI, uri);
        assertEq(p.metadataHash, hash);
    }

    // --- newProcess ------------------------------------------------------------

    function test_NewProcess_StoresAndEmitsMetadata() public {
        bytes31 pid = registry.getNextProcessId(ORGANIZER);
        vm.expectEmit(true, true, false, true, address(registry));
        emit IProcessRegistry.ProcessCreated(pid, ORGANIZER);
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessMetadataUpdated(pid, METADATA_URI, METADATA_HASH);
        assertEq(_create(METADATA_URI, METADATA_HASH), pid);
        _assertMetadata(pid, METADATA_URI, METADATA_HASH);
    }

    function test_NewProcess_RevertWhen_EmptyURI() public {
        vm.expectRevert(IProcessRegistry.InvalidMetadata.selector);
        _create("", METADATA_HASH);
    }

    function test_NewProcess_RevertWhen_ZeroHash() public {
        vm.expectRevert(IProcessRegistry.InvalidMetadata.selector);
        _create(METADATA_URI, bytes32(0));
    }

    // --- setProcessMetadata: allowed ---------------------------------------------

    function test_SetMetadata_WhileReady() public {
        bytes31 pid = _create(METADATA_URI, METADATA_HASH);
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessMetadataUpdated(pid, URI_2, HASH_2);
        _set(pid, URI_2, HASH_2);
        _assertMetadata(pid, URI_2, HASH_2);
    }

    function test_SetMetadata_WhilePaused() public {
        bytes31 pid = _create(METADATA_URI, METADATA_HASH);
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.PAUSED);
        vm.expectEmit(true, false, false, true, address(registry));
        emit IProcessRegistry.ProcessMetadataUpdated(pid, URI_2, HASH_2);
        _set(pid, URI_2, HASH_2);
        _assertMetadata(pid, URI_2, HASH_2);
    }

    function test_SetMetadata_BeforeStart() public {
        vm.prank(ORGANIZER);
        bytes31 pid = registry.newProcess(
            DAVINCITypes.ProcessStatus.READY,
            block.timestamp + 1 hours,
            DURATION,
            MAX_VOTERS,
            _ballotMode(),
            _census(),
            METADATA_URI,
            METADATA_HASH,
            _encKey(),
            _noDkg()
        );
        _set(pid, URI_2, HASH_2);
        _assertMetadata(pid, URI_2, HASH_2);
    }

    function test_SetMetadata_UntilTheLastSecond() public {
        bytes31 pid = _create(METADATA_URI, METADATA_HASH);
        vm.warp(registry.getProcessEndTime(pid) - 1);
        _set(pid, URI_2, HASH_2);
        _assertMetadata(pid, URI_2, HASH_2);
    }

    /// @dev The event log alone is the full history: creation, then each update, in order.
    function test_SetMetadata_EventsAreTheHistory() public {
        vm.recordLogs();
        bytes31 pid = _create(METADATA_URI, METADATA_HASH);
        _set(pid, URI_2, HASH_2);
        _set(pid, URI_3, HASH_3);
        _assertMetadata(pid, URI_3, HASH_3);

        Vm.Log[] memory logs = vm.getRecordedLogs();
        string[3] memory uris = [METADATA_URI, URI_2, URI_3];
        bytes32[3] memory hashes = [METADATA_HASH, HASH_2, HASH_3];
        uint256 n;
        for (uint256 i = 0; i < logs.length; i++) {
            if (logs[i].topics[0] != IProcessRegistry.ProcessMetadataUpdated.selector) continue;
            assertEq(logs[i].emitter, address(registry));
            assertEq(logs[i].topics[1], bytes32(pid));
            (string memory uri, bytes32 hash) = abi.decode(logs[i].data, (string, bytes32));
            assertEq(uri, uris[n]);
            assertEq(hash, hashes[n]);
            n++;
        }
        assertEq(n, 3);
    }

    // --- setProcessMetadata: refused ---------------------------------------------

    function test_SetMetadata_RevertWhen_NotOrganizer() public {
        bytes31 pid = _create(METADATA_URI, METADATA_HASH);
        vm.prank(address(0xdead));
        vm.expectRevert(IProcessRegistry.Unauthorized.selector);
        registry.setProcessMetadata(pid, URI_2, HASH_2);
        _assertMetadata(pid, METADATA_URI, METADATA_HASH);
    }

    function test_SetMetadata_RevertWhen_UnknownProcess() public {
        bytes31 pid = registry.getNextProcessId(ORGANIZER);
        vm.expectRevert(IProcessRegistry.ProcessNotFound.selector);
        _set(pid, URI_2, HASH_2);
        vm.expectRevert(IProcessRegistry.InvalidProcessId.selector);
        _set(bytes31(0), URI_2, HASH_2);
        // An id of another registry (prefix bytes 20..23 differ).
        vm.expectRevert(IProcessRegistry.UnknownProcessIdPrefix.selector);
        _set(pid ^ bytes31(uint248(0xff) << 56), URI_2, HASH_2);
    }

    function test_SetMetadata_RevertWhen_EmptyURIOrZeroHash() public {
        bytes31 pid = _create(METADATA_URI, METADATA_HASH);
        vm.expectRevert(IProcessRegistry.InvalidMetadata.selector);
        _set(pid, "", HASH_2);
        vm.expectRevert(IProcessRegistry.InvalidMetadata.selector);
        _set(pid, URI_2, bytes32(0));
        _assertMetadata(pid, METADATA_URI, METADATA_HASH);
    }

    function test_SetMetadata_RevertWhen_PastTheEnd() public {
        bytes31 pid = _create(METADATA_URI, METADATA_HASH);
        vm.warp(registry.getProcessEndTime(pid));
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        _set(pid, URI_2, HASH_2);

        // Paused does not stop the clock.
        bytes31 paused = _create(METADATA_URI, METADATA_HASH);
        vm.prank(ORGANIZER);
        registry.setProcessStatus(paused, DAVINCITypes.ProcessStatus.PAUSED);
        vm.warp(registry.getProcessEndTime(paused) + 1);
        vm.expectRevert(IProcessRegistry.InvalidTimeBounds.selector);
        _set(paused, URI_2, HASH_2);
    }

    function test_SetMetadata_RevertWhen_Ended() public {
        bytes31 pid = _create(METADATA_URI, METADATA_HASH);
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.ENDED);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        _set(pid, URI_2, HASH_2);
    }

    function test_SetMetadata_RevertWhen_Canceled() public {
        bytes31 pid = _create(METADATA_URI, METADATA_HASH);
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.CANCELED);
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        _set(pid, URI_2, HASH_2);
    }

    function test_SetMetadata_RevertWhen_Results() public {
        bytes31 pid = _create(METADATA_URI, METADATA_HASH);
        vm.prank(ORGANIZER);
        registry.setProcessStatus(pid, DAVINCITypes.ProcessStatus.ENDED);
        _warpToGraceEnd(pid);
        uint64[16] memory values;
        values[0] = 7;
        registry.setProcessResults(pid, _resultsPublics(registry.getProcess(pid).latestStateRoot, values), "");
        assertEq(uint8(registry.getProcess(pid).status), uint8(DAVINCITypes.ProcessStatus.RESULTS));
        vm.expectRevert(IProcessRegistry.InvalidStatus.selector);
        _set(pid, URI_2, HASH_2);
        _assertMetadata(pid, METADATA_URI, METADATA_HASH);
    }
}
