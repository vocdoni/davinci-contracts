// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {ICouncilManager} from "./interfaces/council/ICouncilManager.sol";
import {IDkgResultsAdapter} from "./interfaces/IDkgResultsAdapter.sol";
import {DAVINCITypes} from "./libraries/DAVINCITypes.sol";

/**
 * @title CouncilAdapter
 * @notice The registry's seam to a Council manager (invite-only threshold DKG): it binds
 *         each COUNCIL-mode process to the ceremony its DKGParams name, hands the ceremony
 *         key to the registry, submits the final accumulator as one decryption request and
 *         reads the combined plaintexts back. Created by the ProcessRegistry constructor, so
 *         `registry` is fixed to its creator. A ceremony's organizer must allow this adapter
 *         and authorize the process creator on the manager before a process can bind.
 * @dev Council stores and returns circomlib (TE) points, the registry's form, so unlike
 *      DavinciDKGAdapter nothing is converted. The registry stores the request id as the
 *      process's dkgAid; the manager keys requests by the process id, which this adapter
 *      maps back.
 */
contract CouncilAdapter is IDkgResultsAdapter {
    /// @notice A bound process: its ceremony, the number of fields submitted (0 until the
    ///         request) and the registry process id.
    struct Binding {
        bytes12 ceremonyId;
        uint8 fieldCount;
        bytes31 processId;
    }

    /// @notice The ProcessRegistry that created this adapter; the only allowed caller of
    ///         the state-changing functions.
    address public immutable registry;
    /// @notice The Council manager (ceremonies, bindings, requests, combines).
    ICouncilManager public immutable manager;
    /// @notice Bindings by request id.
    mapping(bytes32 requestId => Binding) public bindings;

    error NotRegistry();
    /// @dev Selector-identical to IProcessRegistry.InvalidKeyMode and InvalidDKGParams, so
    ///      the revert decodes against the registry ABI when it bubbles out of newProcess.
    error InvalidKeyMode();
    error InvalidDKGParams();
    /// @notice Council has no organizer secret to reveal.
    error UnsupportedKeyMode();
    /// @notice No request was submitted under this request id and ceremony.
    error UnknownRequest();
    /// @notice The manager answered for a different request.
    error RequestMismatch();
    /// @notice plaintexts covers whole requests only: first == 0, count == fieldCount.
    error InvalidFieldRange();

    modifier onlyRegistry() {
        if (msg.sender != registry) revert NotRegistry();
        _;
    }

    constructor(address councilManager) {
        registry = msg.sender;
        manager = ICouncilManager(councilManager);
    }

    /// @notice Binds `processId`, created by `creator`, to the ceremony in `dkg.epochId` and
    ///         returns that ceremony id, the request id and the ceremony key in TE. The
    ///         manager requires a Live ceremony, this adapter allowed and `creator`
    ///         authorized.
    function register(bytes31 processId, address creator, DAVINCITypes.DKGParams calldata dkg)
        external
        onlyRegistry
        returns (bytes12 cid, bytes32 requestId, uint256 pkX, uint256 pkY)
    {
        if (dkg.mode != DAVINCITypes.KeyMode.COUNCIL) revert InvalidKeyMode();
        cid = dkg.epochId;
        if (
            cid == bytes12(0) || dkg.orgPKx != 0 || dkg.orgPKy != 0 || dkg.popAx != 0 || dkg.popAy != 0 || dkg.popZ != 0
        ) revert InvalidDKGParams();
        (requestId, pkX, pkY) = manager.bindProcess(cid, processId, creator);
        bindings[requestId] = Binding({ceremonyId: cid, fieldCount: 0, processId: processId});
    }

    /// @notice Submits the active ciphertexts (TE, [c1x, c1y, c2x, c2y] each) of the process
    ///         bound under `requestId` as one request. Council numbers fields within the
    ///         request, so the first index is always 0.
    function submit(bytes12 cid, bytes32 requestId, uint256[4][] calldata cts)
        external
        onlyRegistry
        returns (uint16 firstIndex)
    {
        Binding storage b = bindings[requestId];
        bytes31 processId = b.processId;
        if (processId == bytes31(0) || b.ceremonyId != cid) revert UnknownRequest();
        // The registry submits at most 16 fields (numFields); the manager rejects more.
        b.fieldCount = uint8(cts.length);
        if (manager.submitRequest(cid, processId, cts) != requestId) revert RequestMismatch();
        return 0;
    }

    /// @notice The combined plaintexts of the request, in field order. Only the whole
    ///         request can be read (first == 0, count == its field count), and ready is
    ///         false until every field is combined, so the registry never finalizes a
    ///         partial vector.
    function plaintexts(bytes12 cid, bytes32 requestId, uint16 first, uint16 count)
        external
        view
        returns (bool ready, uint256[] memory values)
    {
        Binding memory b = bindings[requestId];
        if (b.fieldCount == 0 || b.ceremonyId != cid) revert UnknownRequest();
        if (first != 0 || count != b.fieldCount) revert InvalidFieldRange();
        (ready, values) = manager.getPlaintexts(requestId);
        if (ready && values.length != count) revert RequestMismatch();
    }

    /// @notice Council processes have no organizer key; always reverts.
    function reveal(bytes12, bytes32, uint256) external pure {
        revert UnsupportedKeyMode();
    }
}
