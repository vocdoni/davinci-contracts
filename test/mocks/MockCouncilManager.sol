// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {ICouncilManager} from "../../src/interfaces/council/ICouncilManager.sol";
import {ICouncilManagerErrors as E} from "../../src/interfaces/council/ICouncilManagerErrors.sol";

/**
 * @dev The CouncilManager's adapter surface (ICouncilManager, protocol v2) with the checks,
 *      revert order and request-id derivation of vocdoni/davinci-dkg-council
 *      `solidity/src/CouncilManager.sol`. Ceremonies, authorization and plaintexts come from
 *      test setters instead of signed actions, dealings, partials and combines. A ciphertext
 *      half is checked canonical, on the TE curve and not the identity; the prime-subgroup
 *      check is left to the real manager's suite. The decryption gate is the §8.7 predicate
 *      over the ceremony's policy, and setPlaintext (a combine) is refused while it is closed.
 *      Test setters: newCeremony, setPhase, setDecryptionPolicy, openDecryption,
 *      allowAdapter, authorizeCreator, setPlaintext. Extras of the real manager's read
 *      surface: getPolicy (decryption fields), getRequestIds, requestCts, bindingCreator.
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

    /// @dev Values of the manager's `PhaseMode` (protocol §2.3).
    enum PhaseMode {
        Manual,
        Scheduled
    }

    /// @dev The manager's openDecryption reverts; they never reach the adapter.
    error WrongMode();
    error AlreadyOpen();

    /// @dev The manager's `PhasePolicyView` (architecture §1.2).
    struct PhasePolicyView {
        uint8 registrationMode;
        uint8 decryptionMode;
        uint64 dealingDuration;
        uint64 decryptionOpenAt;
        uint64 manualDecryptionFallbackAt;
        uint64 manualOpenedAt;
        bool decryptionOpen;
        bool scheduledRegistrationCloseDue;
    }

    struct Ceremony {
        Phase phase;
        PhaseMode decryptionMode;
        uint64 decryptionOpenAt; // Scheduled only
        uint64 manualDecryptionFallbackAt; // Manual only, 0 = no fallback
        uint64 manualOpenedAt; // 0 until openDecryption
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

    /// @dev A Live ceremony whose decryption is already open: Scheduled, opening now.
    function newCeremony(bytes12 cid, uint256 pkX, uint256 pkY) external {
        Ceremony storage c = ceremonies[cid];
        c.phase = Phase.Live;
        (c.pkX, c.pkY) = (pkX, pkY);
        (c.decryptionMode, c.decryptionOpenAt) = (PhaseMode.Scheduled, uint64(block.timestamp));
    }

    /// @dev Replaces the decryption policy (the real one is fixed at createCeremony) and
    ///      clears a manual opening.
    function setDecryptionPolicy(bytes12 cid, PhaseMode mode, uint64 openAt, uint64 fallbackAt) external {
        Ceremony storage c = ceremonies[cid];
        (c.decryptionMode, c.decryptionOpenAt, c.manualDecryptionFallbackAt, c.manualOpenedAt) =
        (mode, openAt, fallbackAt, 0);
    }

    /// @dev The organizer's openDecryption, without the signed action.
    function openDecryption(bytes12 cid) external {
        Ceremony storage c = _existing(cid);
        if (c.phase != Phase.Live) revert E.WrongPhase();
        if (c.decryptionMode != PhaseMode.Manual) revert WrongMode();
        if (isDecryptionOpen(cid)) revert AlreadyOpen();
        c.manualOpenedAt = uint64(block.timestamp);
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

    /// @dev Marks `field` combined with `value`, as a combine would: only once the gate is open.
    function setPlaintext(bytes32 requestId, uint8 field, uint64 value) external {
        Request storage r = requests[requestId];
        require(field < r.fieldCount, "field");
        if (!isDecryptionOpen(r.cid)) revert E.DecryptionNotOpen();
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

    function getRequestMeta(bytes32 requestId)
        external
        view
        returns (bytes12 cid, uint8 fieldCount, uint16 completedBitmap, uint16 partialBitmap)
    {
        Request storage r = _bound(requestId);
        return (r.cid, r.fieldCount, r.completedBitmap, 0);
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

    /// @dev Protocol §8.7.
    function isDecryptionOpen(bytes12 cid) public view returns (bool) {
        Ceremony storage c = _existing(cid);
        if (c.phase != Phase.Live) return false;
        if (c.decryptionMode == PhaseMode.Scheduled) return block.timestamp >= c.decryptionOpenAt;
        return
            c.manualOpenedAt != 0
                || (c.manualDecryptionFallbackAt != 0 && block.timestamp >= c.manualDecryptionFallbackAt);
    }

    // --- extras the real manager also has --------------------------------------------

    /// @dev The decryption half of the real view; the registration fields read 0.
    function getPolicy(bytes12 cid) external view returns (PhasePolicyView memory v) {
        Ceremony storage c = _existing(cid);
        v.decryptionMode = uint8(c.decryptionMode);
        v.decryptionOpenAt = c.decryptionOpenAt;
        v.manualDecryptionFallbackAt = c.manualDecryptionFallbackAt;
        v.manualOpenedAt = c.manualOpenedAt;
        v.decryptionOpen = isDecryptionOpen(cid);
    }

    function getRequestIds(bytes12 cid) external view returns (bytes32[] memory) {
        return _existing(cid).requestIds;
    }

    /// @dev The submitted ciphertexts in full (the real manager stores them compressed and
    ///      serves getRequestCompressed).
    function requestCts(bytes32 requestId) external view returns (uint256[4][] memory cts) {
        Request storage r = _bound(requestId);
        cts = new uint256[4][](r.fieldCount);
        for (uint256 k; k < cts.length; ++k) {
            cts[k] = r.cts[k];
        }
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
