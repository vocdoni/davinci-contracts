// SPDX-License-Identifier: MIT

pragma solidity ^0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as ResultsVerifierBaseGroth16} from "./resultsverifier_vkey.sol";

contract ResultsVerifierGroth16 is IZKVerifier, ResultsVerifierBaseGroth16 {
    /// @inheritdoc IZKVerifier
    function verifyProof(bytes calldata _proof, bytes calldata _input) external view override {
        ResultsVerifierBaseGroth16(address(this)).verifyProof(_proof, _decodeInput(_input));
    }

    /// @inheritdoc IZKVerifier
    function provingKeyHash() external pure override returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function _decodeInput(bytes calldata encodedInputs) internal pure returns (uint256[9] memory) {
        return abi.decode(encodedInputs, (uint256[9]));
    }
}
