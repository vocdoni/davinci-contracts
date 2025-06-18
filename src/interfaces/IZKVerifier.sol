// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

interface IZKVerifier {
    /**
     * @notice Verify a ZK proof.
     * @notice There is no return value. If the function does not revert, the proof was successfully verified.
     * @param proof The proof points.
     * @param input The public input elements.
     */
    function verifyProof(bytes calldata proof, bytes calldata input) external view;

    /**
     * @notice Returns the hash of the proving key.
     * @return The hash of the proving key.
     * @dev This is used to verify that the correct proving key is being used.
     */
    function provingKeyHash() external pure returns (bytes32);
}
