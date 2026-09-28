// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/**
 * @title BjjFormLib
 * @notice BabyJubJub chart change between the circomlib twisted Edwards form
 *         (a = 168700, DAVINCI's ballots and state leaves) and the reduced form
 *         (a = -1, gnark and the davinci-dkg contracts). The map scales x by a
 *         square root of -a/168700 and leaves y unchanged:
 *         x_rte = x_te · K mod Q, x_te = x_rte · K_INV mod Q. It fixes the
 *         identity (0, 1) and maps circomlib's B8 generator to the DKG's G.
 */
library BjjFormLib {
    /// @dev BN254 scalar field prime, the base field of BabyJubJub.
    uint256 internal constant Q = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
    /// @dev sqrt(-168700^-1 · 1) scale factor: K^2 = -168700 mod Q.
    uint256 internal constant K = 15527681003928902128179717624703512672403908117992798440346960750464748824729;
    /// @dev K^-1 mod Q.
    uint256 internal constant K_INV = 1911982854305225074381251344103329931637610209014896889891168275855466657090;

    /// @notice circomlib x-coordinate -> reduced (gnark/DKG) x-coordinate.
    function toRTE(uint256 xTE) internal pure returns (uint256) {
        return mulmod(xTE, K, Q);
    }

    /// @notice Reduced (gnark/DKG) x-coordinate -> circomlib x-coordinate.
    function toTE(uint256 xRTE) internal pure returns (uint256) {
        return mulmod(xRTE, K_INV, Q);
    }
}
