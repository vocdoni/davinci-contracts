// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {DKGTypes} from "./DKGTypes.sol";

/**
 * @title IDKGManager
 * @notice Vendored subset of davinci-dkg `src/interfaces/IDKGManager.sol` (commit 2338d8a),
 *         plus the public state getters (`appManager`, `epochNonce`, `EPOCH_PREFIX`) the
 *         adapter reads from `DKGManager`. ABI-identical to upstream.
 */
interface IDKGManager {
    error InvalidEpoch();
    error InvalidPhase();
    error InvalidProofInput();
    error InvalidCiphertext();
    error CiphertextAlreadySubmitted();
    error DecryptionLimitReached();
    error PoolExhausted();
    error Unauthorized();

    /// @notice The sibling DKGAppManager (public state variable on DKGManager).
    function appManager() external view returns (address);

    /// @notice Nonce of the newest epoch; epoch ids are
    ///         `bytes12((uint96(EPOCH_PREFIX) << 64) | nonce)` (DKGIdLib.computeEpochId).
    function epochNonce() external view returns (uint64);

    /// @notice The manager's epoch-id prefix (public immutable on DKGManager).
    // solhint-disable-next-line func-name-mixedcase
    function EPOCH_PREFIX() external view returns (uint32);

    /// @notice Submit a ciphertext for threshold decryption under `PK_aid`. The
    ///         index is assigned on chain (1, 2, ... per application) and returned.
    ///         Coordinates must be canonical, on-curve and non-identity.
    function submitCiphertext(bytes12 epochId, bytes32 aid, uint256 c1x, uint256 c1y, uint256 c2x, uint256 c2y)
        external
        returns (uint16 ciphertextIndex);

    function getCombinedDecryption(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex)
        external
        view
        returns (DKGTypes.CombinedDecryptionRecord memory);

    /// @notice The pool key `P_j` of a Live epoch. Reverts `InvalidPhase` before the
    ///         epoch is Live and `InvalidProofInput` for `keyIndex >= MAX_K`.
    function getPoolKey(bytes12 epochId, uint8 keyIndex) external view returns (uint256 x, uint256 y);

    /// @notice `nextIndex` is the pool key the next registration claims
    ///         (`MAX_K` once the pool is spent). Plain mapping read, never reverts.
    function getPoolStatus(bytes12 epochId) external view returns (uint8 nextIndex);

    /// @notice The pool key claimed by `aid`. Reverts `InvalidProofInput` when the
    ///         application never claimed one.
    function getAppPoolIndex(bytes12 epochId, bytes32 aid) external view returns (uint8);
}
