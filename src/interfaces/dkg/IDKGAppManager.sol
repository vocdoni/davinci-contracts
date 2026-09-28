// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {DKGTypes} from "./DKGTypes.sol";

/**
 * @title IDKGAppManager
 * @notice Vendored subset of davinci-dkg `src/interfaces/IDKGAppManager.sol`. ABI-identical to upstream.
 */
interface IDKGAppManager {
    error InvalidApplication();
    error ApplicationAlreadyExists();
    error InvalidSchnorrProof();
    error PointNotInSubgroup();
    error InvalidEpoch();
    error InvalidPhase();
    error InvalidOrganizerSecret();
    error InvalidPolicy();
    error AlreadyRevealed();
    error PoolExhausted();

    /// @notice Register an application against a Live epoch and claim the epoch's
    ///         next pool key. In `OrganizerLocked` mode the Schnorr PoP proves
    ///         knowledge of `sk_org` and `PK_aid = P_j + PK_org`; in `Automatic`
    ///         mode the key and Schnorr arguments are ignored and `PK_aid = P_j`.
    ///         Organizer key and PoP are in the DKG's reduced (a = -1) form.
    function registerApplication(
        bytes12 epochId,
        bytes32 aid,
        DKGTypes.AppPolicy calldata policy,
        uint256 pkOrgX,
        uint256 pkOrgY,
        uint256 schnorrAx,
        uint256 schnorrAy,
        uint256 schnorrZ
    ) external;

    /// @notice Publish `sk_org` for an `OrganizerLocked` application, once.
    ///         Permissionless — the contract checks `sk·G == PK_org`.
    function revealOrganizerSecret(bytes12 epochId, bytes32 aid, uint256 organizerSecret) external;

    function getApplication(bytes12 epochId, bytes32 aid) external view returns (DKGTypes.Application memory);

    /// @notice `PK_aid` of a registered application in reduced form: `P_j` for
    ///         `Automatic`, `P_j + PK_org` for `OrganizerLocked`. Reverts
    ///         `InvalidApplication` for an unknown aid.
    function getApplicationKey(bytes12 epochId, bytes32 aid) external view returns (uint256 x, uint256 y);

    /// @notice `PK_org` of a registered application — `(0, 1)` for Automatic ones,
    ///         zero for an unknown aid.
    function getOrganizerPK(bytes12 epochId, bytes32 aid) external view returns (uint256, uint256);
}
