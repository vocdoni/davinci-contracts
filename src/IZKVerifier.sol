// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

interface IZKVerifier {
    /**
     * @notice InvalidProof error is emitted when a ZK proof is invalid.
     */
    error InvalidProof();

    /**
     * @notice Verify a ZK proof.
     * @notice There is no return value. If the function does not revert, the proof was successfully verified.
     * @param proof The proof points.
     * @param input The public input elements.
     */
    function verifyProof(bytes calldata proof, bytes calldata input) external view;
}
