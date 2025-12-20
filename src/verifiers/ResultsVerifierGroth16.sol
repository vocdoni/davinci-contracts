// SPDX-License-Identifier: MIT

pragma solidity ^0.8.28;

import { IZKVerifier } from "../interfaces/IZKVerifier.sol";
import { Verifier as ResultsVerifierBaseGroth16 } from "./resultsverifier_vkey.sol";

contract ResultsVerifierGroth16 is IZKVerifier, ResultsVerifierBaseGroth16 {
    error InvalidProofLength(uint256 length, uint256 expected);
    error InvalidInputLength(uint256 length, uint256 expected);

    /// @inheritdoc IZKVerifier
    function verifyProof(bytes calldata _proof, bytes calldata _input) external view override {
        (uint256[8] memory proof, uint256[2] memory commitments, uint256[2] memory commitmentPok) = _decodeProof(
            _proof
        );
        this.verifyProof(proof, commitments, commitmentPok, _decodeInput(_input));
    }

    /// @inheritdoc IZKVerifier
    function provingKeyHash() external pure override returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function _decodeProof(
        bytes calldata encodedProof
    ) internal pure returns (uint256[8] memory, uint256[2] memory, uint256[2] memory) {
        if (encodedProof.length != 384) revert InvalidProofLength(encodedProof.length, 384);
        return abi.decode(encodedProof, (uint256[8], uint256[2], uint256[2]));
    }

    function _decodeInput(bytes calldata encodedInputs) internal pure returns (uint256[9] memory) {
        if (encodedInputs.length != 288) revert InvalidInputLength(encodedInputs.length, 288);
        return abi.decode(encodedInputs, (uint256[9]));
    }
}
