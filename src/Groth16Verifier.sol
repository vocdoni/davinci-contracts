// SPDX-License-Identifier: MIT

pragma solidity ^0.8.28;

import { IZKVerifier } from "./IZKVerifier.sol";
import { Groth16VerifierBase } from "./Groth16VerifierBase.sol";

contract Groth16Verifier is IZKVerifier, Groth16VerifierBase {
    function verifyProof(bytes calldata _proof, bytes calldata _input) external view override {
        (uint256[8] memory proof, uint256[2] memory commitments, uint256[2] memory commitmentPok) = _decodeProof(
            _proof
        );
        this.verifyProof(proof, commitments, commitmentPok, _decodeInput(_input));
    }

    function _decodeProof(
        bytes calldata encodedProof
    ) internal pure returns (uint256[8] memory, uint256[2] memory, uint256[2] memory) {
        return abi.decode(encodedProof, (uint256[8], uint256[2], uint256[2]));
    }

    function _decodeInput(bytes calldata _encodedInputs) internal pure returns (uint256[4] memory) {
        return abi.decode(_encodedInputs, (uint256[4]));
    }
}
