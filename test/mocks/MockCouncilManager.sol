// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {ICouncilManager} from "../../src/interfaces/council/ICouncilManager.sol";
import {ICouncilManagerErrors as E} from "../../src/interfaces/council/ICouncilManagerErrors.sol";

/**
 * @dev The CouncilManager's adapter surface (ICouncilManager) with the checks, revert order
 *      and request-id derivation of vocdoni/davinci-dkg-council `solidity/src/CouncilManager.sol`.
 *      Ceremonies, authorization and plaintexts come from test setters instead of
 *      signed actions, dealings, partials and combines. A ciphertext half is checked canonical,
 *      on the TE curve and not the identity; the prime-subgroup check is left to the real
 *      manager's suite.
 *      Test setters: newCeremony, setPhase, allowAdapter, authorizeCreator, setPlaintext.
 */
contract MockCouncilManager is ICouncilManager {
    uint256 internal constant P = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    uint256 internal constant A = 168700;
    uint256 internal constant D = 168696;
    uint256 internal constant MAX_FIELDS = 16;

    /// @dev Values of the manager's `Phase` (CouncilTypes.sol).
    enum Phase {
        None,
        Registration,
        Dealing,
        Live,
        Aborted
    }

    struct Ceremony {
        Phase phase;
        uint256 pkX;
        uint256 pkY;
        bytes32[] requestIds;
        mapping(address => bool) adapters;
        mapping(address => bool) creators;
    }

    struct Request {
        bytes12 cid;
        address adapter;
        bytes31 processId;
        uint8 fieldCount; // 0 until submitted
        address creator;
        uint16 completedBitmap;
        uint256[4][16] cts;
        uint64[16] plaintexts;
    }

    mapping(bytes12 => Ceremony) internal ceremonies;
    mapping(bytes32 => bytes32) internal bindings; // keccak(adapter, processId) => requestId
    mapping(bytes32 => Request) internal requests;

    // --- test setters ------------------------------------------------------------

    function newCeremony(bytes12 cid, uint256 pkX, uint256 pkY) external {
        Ceremony storage c = ceremonies[cid];
        c.phase = Phase.Live;
        (c.pkX, c.pkY) = (pkX, pkY);
    }

    function setPhase(bytes12 cid, Phase phase) external {
        ceremonies[cid].phase = phase;
    }

    function allowAdapter(bytes12 cid, address adapter) external {
        ceremonies[cid].adapters[adapter] = true;
    }

    function authorizeCreator(bytes12 cid, address creator) external {
        ceremonies[cid].creators[creator] = true;
    }

    /// @dev Marks `field` combined with `value`, as a combine would.
    function setPlaintext(bytes32 requestId, uint8 field, uint64 value) external {
        Request storage r = requests[requestId];
        require(field < r.fieldCount, "field");
        r.plaintexts[field] = value;
        r.completedBitmap |= uint16(1 << field);
    }

    // --- ICouncilManager ----------------------------------------------------------------

    function requestIdFor(bytes12 cid, address adapter, bytes31 processId) public view returns (bytes32) {
        return keccak256(
            abi.encode(
                keccak256("davinci-dkg-council/v1/request"), block.chainid, address(this), cid, adapter, processId
            )
        );
    }

    function bindProcess(bytes12 cid, bytes31 processId, address creator)
        external
        returns (bytes32 requestId, uint256 pkX, uint256 pkY)
    {
        Ceremony storage c = _existing(cid);
        if (c.phase != Phase.Live) revert E.WrongPhase();
        if (!c.adapters[msg.sender]) revert E.NotAllowedAdapter();
        if (!c.creators[creator]) revert E.NotAuthorizedCreator();
        bytes32 key = keccak256(abi.encode(msg.sender, processId));
        if (bindings[key] != bytes32(0)) revert E.AlreadyBound();
        requestId = requestIdFor(cid, msg.sender, processId);
        bindings[key] = requestId;
        Request storage r = requests[requestId];
        (r.cid, r.adapter, r.processId, r.creator) = (cid, msg.sender, processId, creator);
        c.requestIds.push(requestId);
        (pkX, pkY) = (c.pkX, c.pkY);
    }

    function submitRequest(bytes12 cid, bytes31 processId, uint256[4][] calldata cts)
        external
        returns (bytes32 requestId)
    {
        Ceremony storage c = _existing(cid);
        requestId = bindings[keccak256(abi.encode(msg.sender, processId))];
        if (requestId == bytes32(0)) revert E.UnknownBinding();
        Request storage r = requests[requestId];
        if (r.cid != cid) revert E.UnknownBinding();
        if (c.phase != Phase.Live) revert E.WrongPhase();
        if (r.fieldCount != 0) revert E.AlreadyRequested();
        uint256 count = cts.length;
        if (count == 0 || count > MAX_FIELDS) revert E.BadFieldCount();
        for (uint256 k; k < count; ++k) {
            uint256[4] calldata ct = cts[k];
            _requirePoint(ct[0], ct[1]);
            _requirePoint(ct[2], ct[3]);
            r.cts[k] = ct;
        }
        r.fieldCount = uint8(count);
    }

    function getPlaintexts(bytes32 requestId) external view returns (bool ready, uint256[] memory values) {
        Request storage r = _bound(requestId);
        uint256 count = r.fieldCount;
        values = new uint256[](count);
        for (uint256 k; k < count; ++k) {
            values[k] = r.plaintexts[k];
        }
        ready = count != 0 && r.completedBitmap == (1 << count) - 1;
    }

    function getRequest(bytes32 requestId)
        external
        view
        returns (bytes12 cid, uint8 fieldCount, uint16 completedBitmap, uint16 partialBitmap, uint256[4][] memory cts)
    {
        Request storage r = _bound(requestId);
        fieldCount = r.fieldCount;
        cts = new uint256[4][](fieldCount);
        for (uint256 k; k < fieldCount; ++k) {
            cts[k] = r.cts[k];
        }
        return (r.cid, fieldCount, r.completedBitmap, 0, cts);
    }

    function getBinding(address adapter, bytes31 processId)
        external
        view
        returns (bytes12 cid, bytes32 requestId, bool requested)
    {
        requestId = bindings[keccak256(abi.encode(adapter, processId))];
        if (requestId == bytes32(0)) revert E.UnknownBinding();
        Request storage r = requests[requestId];
        return (r.cid, requestId, r.fieldCount != 0);
    }

    function getPublicKey(bytes12 cid) external view returns (uint256 x, uint256 y) {
        Ceremony storage c = _existing(cid);
        if (c.phase != Phase.Live) revert E.WrongPhase();
        return (c.pkX, c.pkY);
    }

    // --- extras the real manager also has --------------------------------------------

    function getRequestIds(bytes12 cid) external view returns (bytes32[] memory) {
        return _existing(cid).requestIds;
    }

    /// @dev The creator of a request (the real manager's getRequestOrigin, third value).
    function bindingCreator(address adapter, bytes31 processId) external view returns (address) {
        return requests[bindings[keccak256(abi.encode(adapter, processId))]].creator;
    }

    // --- internals ---------------------------------------------------------------------

    function _existing(bytes12 cid) internal view returns (Ceremony storage c) {
        c = ceremonies[cid];
        if (c.phase == Phase.None) revert E.UnknownCeremony();
    }

    function _bound(bytes32 requestId) internal view returns (Request storage r) {
        r = requests[requestId];
        if (r.cid == bytes12(0)) revert E.UnknownRequest();
    }

    /// @dev Canonical, on a·x² + y² = 1 + d·x²·y², not the identity (0, 1).
    function _requirePoint(uint256 x, uint256 y) internal pure {
        if (x >= P || y >= P) revert E.NonCanonical();
        uint256 x2 = mulmod(x, x, P);
        uint256 y2 = mulmod(y, y, P);
        if (addmod(mulmod(A, x2, P), y2, P) != addmod(1, mulmod(D, mulmod(x2, y2, P), P), P)) {
            revert E.InvalidPoint();
        }
        if (x == 0 && y == 1) revert E.InvalidPoint();
    }
}
