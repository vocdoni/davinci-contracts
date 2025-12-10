// SPDX-License-Identifier: MIT

pragma solidity ^0.8.28;

import { IZKVerifier } from "../interfaces/IZKVerifier.sol";
import { Verifier as StateTransitionVerifierBaseGroth16 } from "./statetransition_vkey.sol";

contract StateTransitionVerifierGroth16 is IZKVerifier, StateTransitionVerifierBaseGroth16 {
    /// @inheritdoc IZKVerifier
    function verifyProof(bytes calldata _proof, bytes calldata _input) external view override {
        (uint256[8] memory proof, uint256[2] memory commitments, uint256[2] memory commitmentPok) = _decodeProof(
            _proof
        );
        (uint256[8] memory zkinput, , ) = _decodeInput(_input);
        this.verifyProof(proof, commitments, commitmentPok, zkinput);
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

    function _decodeInput(
        bytes calldata encodedInputs
    ) internal pure returns (uint256[8] memory, bytes memory, bytes memory) {
        uint256[8] memory inputs = abi.decode(encodedInputs, (uint256[8]));
        return (inputs, "", "");
    }
}
