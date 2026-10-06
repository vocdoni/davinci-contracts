// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/**
 * @title IDkgResultsAdapter
 * @notice The results half shared by the registry's key adapters (DavinciDKGAdapter,
 *         CouncilAdapter). The two leading arguments are the process's stored dkgEpochId
 *         and dkgAid: a davinci-dkg epoch and application id, or a Council ceremony id
 *         and request id. Never a process id.
 */
interface IDkgResultsAdapter {
    /// @notice Submits the active ciphertexts ([C1x, C1y, C2x, C2y], circomlib form) and
    ///         returns the index of the first one, which the registry passes back to
    ///         plaintexts.
    function submit(bytes12 dkgEpochId, bytes32 dkgAid, uint256[4][] calldata cts) external returns (uint16 firstIndex);

    /// @notice Plaintexts of ciphertexts [first, first + count); ready is false until all
    ///         of them are decrypted.
    function plaintexts(bytes12 dkgEpochId, bytes32 dkgAid, uint16 first, uint16 count)
        external
        view
        returns (bool ready, uint256[] memory values);
}
