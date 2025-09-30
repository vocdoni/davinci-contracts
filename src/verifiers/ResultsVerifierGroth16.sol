// SPDX-License-Identifier: MIT

pragma solidity ^0.8.28;

import { IZKVerifier } from "../interfaces/IZKVerifier.sol";
import { ResultsVerifierBaseGroth16 } from "./ResultsVerifierBaseGroth16.sol";

contract ResultsVerifierGroth16 is IZKVerifier, ResultsVerifierBaseGroth16 {
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
        return abi.decode(encodedProof, (uint256[8], uint256[2], uint256[2]));
    }

    function _decodeInput(bytes calldata encodedInputs) internal pure returns (uint256[9] memory) {
        return abi.decode(encodedInputs, (uint256[9])); // TODO: BLOBS
    }
}
