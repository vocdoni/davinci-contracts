// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {IDKGManager} from "./interfaces/dkg/IDKGManager.sol";
import {IDKGAppManager} from "./interfaces/dkg/IDKGAppManager.sol";
import {DKGTypes} from "./interfaces/dkg/DKGTypes.sol";
import {DAVINCITypes} from "./libraries/DAVINCITypes.sol";
import {BjjFormLib} from "./libraries/BjjFormLib.sol";

/**
 * @title DavinciDKGAdapter
 * @notice The registry's seam to the davinci-dkg contracts: it registers one DKG
 *         application per DKG-mode process, is its sole ciphertext submitter, converts
 *         BabyJubJub points between the circomlib form (DAVINCI) and the reduced form
 *         (DKG), and reads combined plaintexts back. Created by the ProcessRegistry
 *         constructor, so `registry` is fixed to its creator.
 */
contract DavinciDKGAdapter {
    /// @notice The ProcessRegistry that created this adapter; the only allowed caller
    ///         of the state-changing functions.
    address public immutable registry;
    /// @notice The davinci-dkg manager (epochs, pool keys, ciphertexts, combines).
    IDKGManager public immutable manager;
    /// @notice The davinci-dkg application manager (registration, organizer reveal).
    IDKGAppManager public immutable appManager;

    /// @dev MAX_K in the DKG: pool keys per epoch.
    uint8 private constant MAX_POOL_KEYS = 16;
    /// @dev How many epochs registrationEpoch scans backwards from the newest.
    uint64 private constant EPOCH_SCAN = 8;

    error NotRegistry();
    error NoLiveEpoch();
    error NonContiguousIndex();

    modifier onlyRegistry() {
        if (msg.sender != registry) revert NotRegistry();
        _;
    }

    constructor(address dkgManager) {
        registry = msg.sender;
        manager = IDKGManager(dkgManager);
        appManager = IDKGAppManager(IDKGManager(dkgManager).appManager());
    }

    /// @notice The DKG application id of a process, `salt << 160 | address(this)` with
    ///         salt the top 92 bits of keccak(chainid, registry, pid). The DKG only lets an
    ///         account register ids whose low 160 bits are its own address, so these ids
    ///         are this adapter's alone: nobody can take a process's id ahead of it. Never
    ///         0 and below 2^252 < Q; deterministic and collision-free across registries.
    function aidFor(bytes31 processId) public view returns (bytes32) {
        uint256 salt = uint256(keccak256(abi.encode(block.chainid, registry, processId))) >> 164;
        return bytes32((salt << 160) | uint256(uint160(address(this))));
    }

    /// @notice The newest Live epoch with a free pool key, scanning backwards at most
    ///         EPOCH_SCAN epochs from the manager's newest. Reverts NoLiveEpoch when
    ///         none qualifies.
    function registrationEpoch() public view returns (bytes12) {
        uint64 nonce = manager.epochNonce();
        uint96 prefix = uint96(manager.EPOCH_PREFIX()) << 64;
        for (uint64 i = 0; i < EPOCH_SCAN && nonce > i; ++i) {
            bytes12 eid = bytes12(prefix | uint96(nonce - i));
            uint8 next = manager.getPoolStatus(eid);
            if (next >= MAX_POOL_KEYS) continue; // pool spent
            // getPoolKey reverts InvalidPhase unless the epoch is Live.
            try manager.getPoolKey(eid, next) returns (uint256, uint256) {
                return eid;
            } catch {}
        }
        revert NoLiveEpoch();
    }

    /// @notice Registers the DKG application of `processId` and returns its epoch, aid
    ///         and encryption key `PK_aid` converted to circomlib form. LOCKED mode uses
    ///         the caller-chosen epoch (the Schnorr PoP binds it); AUTOMATIC picks
    ///         registrationEpoch().
    function register(bytes31 processId, DAVINCITypes.DKGParams calldata dkg)
        external
        onlyRegistry
        returns (bytes12 eid, bytes32 aid, uint256 teX, uint256 teY)
    {
        aid = aidFor(processId);
        bool locked = dkg.mode == DAVINCITypes.KeyMode.DKG_LOCKED;
        eid = locked ? dkg.epochId : registrationEpoch();

        address[] memory submitters = new address[](1);
        submitters[0] = address(this);
        DKGTypes.AppPolicy memory policy = DKGTypes.AppPolicy({
            mode: locked ? DKGTypes.AppMode.OrganizerLocked : DKGTypes.AppMode.Automatic,
            openSubmission: false,
            submitters: submitters,
            maxCiphertexts: 16,
            notBeforeBlock: 0,
            notAfterBlock: 0,
            decryptNotBefore: 0,
            decryptNotAfter: 0
        });
        appManager.registerApplication(eid, aid, policy, dkg.orgPKx, dkg.orgPKy, dkg.popAx, dkg.popAy, dkg.popZ);

        (uint256 rx, uint256 ry) = appManager.getApplicationKey(eid, aid);
        teX = BjjFormLib.toTE(rx);
        teY = ry;
    }

    /// @notice Submits ElGamal ciphertexts (circomlib coordinates, [c1x, c1y, c2x, c2y]
    ///         each) to the DKG, converting to the reduced form. Returns the on-chain
    ///         index of the first one and requires the rest to follow contiguously:
    ///         finalize reads [first, first + count), so a gap would silently map the
    ///         wrong plaintexts to fields.
    function submit(bytes12 eid, bytes32 aid, uint256[4][] calldata ciphertextsTE)
        external
        onlyRegistry
        returns (uint16 firstIndex)
    {
        for (uint256 i = 0; i < ciphertextsTE.length; ++i) {
            uint256[4] calldata ct = ciphertextsTE[i];
            uint16 idx =
                manager.submitCiphertext(eid, aid, BjjFormLib.toRTE(ct[0]), ct[1], BjjFormLib.toRTE(ct[2]), ct[3]);
            if (i == 0) firstIndex = idx;
            else if (idx != firstIndex + i) revert NonContiguousIndex();
        }
    }

    /// @notice Combined plaintexts of ciphertexts [first, first + count). ready is false
    ///         (and values incomplete) until every one has a completed combine.
    function plaintexts(bytes12 eid, bytes32 aid, uint16 first, uint16 count)
        external
        view
        returns (bool ready, uint256[] memory values)
    {
        values = new uint256[](count);
        for (uint16 i = 0; i < count; ++i) {
            DKGTypes.CombinedDecryptionRecord memory rec = manager.getCombinedDecryption(eid, aid, first + i);
            if (!rec.completed) return (false, values);
            values[i] = rec.plaintext;
        }
        ready = true;
    }

    /// @notice Forwards an organizer-secret reveal for a locked application. The DKG
    ///         checks sk·G == PK_org.
    function reveal(bytes12 eid, bytes32 aid, uint256 sk) external onlyRegistry {
        appManager.revealOrganizerSecret(eid, aid, sk);
    }
}
