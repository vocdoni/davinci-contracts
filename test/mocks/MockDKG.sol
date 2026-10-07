// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {DKGTypes} from "../../src/interfaces/dkg/DKGTypes.sol";

/// @dev Real reduced-form (a = -1) BabyJubJub affine math for the mock: the complete
///      twisted Edwards addition law of davinci-dkg's BabyJubJub.sol, with the modexp
///      precompile for the one inversion per add. Slow but exact; test-only.
library MockBjj {
    uint256 internal constant Q = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    uint256 internal constant D = 12181644023421730124874158521699555681764249180949974110617291017600649128846;
    uint256 internal constant GX = 9671717474070082183213120605117400219616337014328744928644933853176787189663;
    uint256 internal constant GY = 16950150798460657717958625567821834550301663161624707787222815936182638968203;

    function inv(uint256 a) private view returns (uint256 r) {
        (bool ok, bytes memory out) = address(0x05).staticcall(abi.encode(32, 32, 32, a, Q - 2, Q));
        require(ok && out.length == 32, "modexp");
        r = abi.decode(out, (uint256));
    }

    /// @dev Complete affine addition: x3 = (x1y2 + y1x2)/(1 + d x1x2y1y2),
    ///      y3 = (y1y2 + x1x2)/(1 - d x1x2y1y2)  (a = -1).
    function add(uint256 x1, uint256 y1, uint256 x2, uint256 y2) internal view returns (uint256 x3, uint256 y3) {
        uint256 dxy = mulmod(D, mulmod(mulmod(x1, x2, Q), mulmod(y1, y2, Q), Q), Q);
        x3 = mulmod(addmod(mulmod(x1, y2, Q), mulmod(y1, x2, Q), Q), inv(addmod(1, dxy, Q)), Q);
        y3 = mulmod(addmod(mulmod(y1, y2, Q), mulmod(x1, x2, Q), Q), inv(addmod(1, Q - dxy, Q)), Q);
    }

    function mulBase(uint256 s) internal view returns (uint256 x, uint256 y) {
        (x, y) = (0, 1);
        (uint256 px, uint256 py) = (GX, GY);
        while (s > 0) {
            if (s & 1 == 1) (x, y) = add(x, y, px, py);
            (px, py) = add(px, py, px, py);
            s >>= 1;
        }
    }

    /// @dev -x² + y² == 1 + d x²y², coordinates canonical.
    function isOnCurve(uint256 x, uint256 y) internal pure returns (bool) {
        if (x >= Q || y >= Q) return false;
        uint256 x2 = mulmod(x, x, Q);
        uint256 y2 = mulmod(y, y, Q);
        return addmod(Q - x2, y2, Q) == addmod(1, mulmod(D, mulmod(x2, y2, Q), Q), Q);
    }
}

/**
 * @dev One contract playing both the DKGManager and the DKGAppManager (appManager()
 *      returns itself). Matches the vendored dkg interfaces' ABI without inheriting them
 *      (the two error sets would collide) and uses real reduced-form BabyJubJub math for
 *      the application key and the reveal check; only the Schnorr PoP is skipped. Like
 *      the real DKGAppManager, it only registers ids bound to the registrant.
 *      Test setters: newEpoch, setLive, setPoolNext, setPlaintext.
 */
contract MockDKG {
    // Errors selector-identical to the vendored IDKGManager / IDKGAppManager ones.
    error InvalidEpoch();
    error InvalidPhase();
    error InvalidProofInput();
    error InvalidCiphertext();
    error PoolExhausted();
    error Unauthorized();
    error InvalidApplication();
    error ApplicationAlreadyExists();
    error InvalidOrganizerSecret();
    error AlreadyRevealed();

    uint32 public constant EPOCH_PREFIX = 0x00444b47;

    struct Epoch {
        bool live;
        uint8 next;
        uint256[2][] keys;
    }

    struct App {
        bool exists;
        bool locked;
        bool revealed;
        address submitter; // policy.submitters[0]; the registrant when empty
        uint16 ctCount;
        uint256 orgPKx;
        uint256 orgPKy;
        uint256 keyX; // PK_aid, reduced form
        uint256 keyY;
    }

    uint64 public epochNonce;
    uint16 public skipAtCall; // test hook: burn one index at the n-th submit (1-based)
    uint16 internal submitCalls;
    mapping(bytes12 => Epoch) internal epochs;
    mapping(bytes12 => mapping(bytes32 => App)) internal apps;
    mapping(bytes12 => mapping(bytes32 => mapping(uint16 => DKGTypes.CombinedDecryptionRecord))) internal combined;

    // --- test setters ---------------------------------------------------------

    function newEpoch(bool live, uint256[2][] memory keys) external returns (bytes12 eid) {
        eid = bytes12((uint96(EPOCH_PREFIX) << 64) | uint96(++epochNonce));
        Epoch storage e = epochs[eid];
        e.live = live;
        for (uint256 i = 0; i < keys.length; ++i) {
            e.keys.push(keys[i]);
        }
    }

    function setLive(bytes12 eid, bool live) external {
        epochs[eid].live = live;
    }

    function setPoolNext(bytes12 eid, uint8 next) external {
        epochs[eid].next = next;
    }

    function setSkipAtCall(uint16 n) external {
        skipAtCall = n;
    }

    function setPlaintext(bytes12 eid, bytes32 aid, uint16 index, uint256 plaintext) external {
        combined[eid][aid][index] = DKGTypes.CombinedDecryptionRecord(index, true, plaintext);
    }

    // --- IDKGManager ------------------------------------------------------------

    function appManager() external view returns (address) {
        return address(this);
    }

    function getPoolStatus(bytes12 eid) external view returns (uint8) {
        return epochs[eid].next;
    }

    function getPoolKey(bytes12 eid, uint8 keyIndex) external view returns (uint256, uint256) {
        Epoch storage e = epochs[eid];
        if (!e.live) revert InvalidPhase();
        if (keyIndex >= e.keys.length) revert InvalidProofInput();
        return (e.keys[keyIndex][0], e.keys[keyIndex][1]);
    }

    function submitCiphertext(bytes12 eid, bytes32 aid, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
        external
        returns (uint16)
    {
        App storage a = apps[eid][aid];
        if (!a.exists) revert InvalidApplication();
        if (msg.sender != a.submitter) revert Unauthorized();
        if ((c1x == 0 && c1y == 1) || (c2x == 0 && c2y == 1)) revert InvalidCiphertext();
        if (!MockBjj.isOnCurve(c1x, c1y) || !MockBjj.isOnCurve(c2x, c2y)) revert InvalidCiphertext();
        if (++submitCalls == skipAtCall) ++a.ctCount; // gap: someone else's ciphertext landed in between
        return ++a.ctCount; // 1-based, per (eid, aid)
    }

    function getCombinedDecryption(bytes12 eid, bytes32 aid, uint16 index)
        external
        view
        returns (DKGTypes.CombinedDecryptionRecord memory)
    {
        return combined[eid][aid][index];
    }

    // --- IDKGAppManager -----------------------------------------------------------

    function registerApplication(
        bytes12 eid,
        bytes32 aid,
        DKGTypes.AppPolicy calldata policy,
        uint256 pkOrgX,
        uint256 pkOrgY,
        uint256, // schnorrAx: PoP not checked by the mock
        uint256, // schnorrAy
        uint256 // schnorrZ
    ) external {
        Epoch storage e = epochs[eid];
        if (!e.live) revert InvalidPhase();
        // DKGAppManager._requireValidAid: a non-zero field element whose low 160 bits are
        // the registrant (aid = salt << 160 | msg.sender).
        if (aid == bytes32(0) || uint256(aid) >= MockBjj.Q || address(uint160(uint256(aid))) != msg.sender) {
            revert InvalidApplication();
        }
        App storage a = apps[eid][aid];
        if (a.exists) revert ApplicationAlreadyExists();
        if (e.next >= e.keys.length) revert PoolExhausted();
        uint256[2] storage pool = e.keys[e.next++];

        a.exists = true;
        a.submitter = policy.submitters.length > 0 ? policy.submitters[0] : msg.sender;
        (a.keyX, a.keyY) = (pool[0], pool[1]);
        if (policy.mode == DKGTypes.AppMode.OrganizerLocked) {
            a.locked = true;
            (a.orgPKx, a.orgPKy) = (pkOrgX, pkOrgY);
            (a.keyX, a.keyY) = MockBjj.add(pool[0], pool[1], pkOrgX, pkOrgY);
        }
    }

    function getApplicationKey(bytes12 eid, bytes32 aid) external view returns (uint256, uint256) {
        App storage a = apps[eid][aid];
        if (!a.exists) revert InvalidApplication();
        return (a.keyX, a.keyY);
    }

    function revealOrganizerSecret(bytes12 eid, bytes32 aid, uint256 sk) external {
        App storage a = apps[eid][aid];
        if (!a.exists || !a.locked) revert InvalidApplication();
        if (a.revealed) revert AlreadyRevealed();
        (uint256 x, uint256 y) = MockBjj.mulBase(sk);
        if (x != a.orgPKx || y != a.orgPKy) revert InvalidOrganizerSecret();
        a.revealed = true;
    }

    function revealed(bytes12 eid, bytes32 aid) external view returns (bool) {
        return apps[eid][aid].revealed;
    }

    function ctCount(bytes12 eid, bytes32 aid) external view returns (uint16) {
        return apps[eid][aid].ctCount;
    }
}
