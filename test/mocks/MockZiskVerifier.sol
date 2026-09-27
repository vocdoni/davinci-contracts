// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {IZiskVerifier} from "../../src/interfaces/IZiskVerifier.sol";

/// @notice Accepts every proof. For tests and local deployments that skip PLONK proving.
contract MockZiskVerifier is IZiskVerifier {
    function verifySnarkProof(bytes32, bytes32, bytes calldata, bytes calldata) external pure {}
}

/// @notice Accepts a proof only when it is checked against one program vk and rootC,
///         the way a real proof only verifies for the program that produced it.
contract PinnedMockZiskVerifier is IZiskVerifier {
    error InvalidProof();

    bytes32 public immutable programVK;
    bytes32 public immutable rootCVadcopFinal;

    constructor(bytes32 _programVK, bytes32 _rootCVadcopFinal) {
        programVK = _programVK;
        rootCVadcopFinal = _rootCVadcopFinal;
    }

    function verifySnarkProof(bytes32 _programVK, bytes32 _rootCVadcopFinal, bytes calldata, bytes calldata)
        external
        view
    {
        if (_programVK != programVK || _rootCVadcopFinal != rootCVadcopFinal) revert InvalidProof();
    }
}
