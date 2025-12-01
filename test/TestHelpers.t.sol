// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { TestInputs } from "./TestInputs.t.sol";

abstract contract TestHelpers is TestInputs {
    /// @dev KZG precompile address as per EIP-4844
    address public constant KZG_PRECOMPILE = address(0x0A);
    /// @dev Number of field elements per blob (as per EIP-4844)
    uint256 public constant FIELD_ELEMENTS_PER_BLOB = 4096;
    /// @dev The modulus used in the BLS signature scheme.
    uint256 public constant BLS_MODULUS = 52435875175126190479447740508185965837690552500527637822603658699938581184513;

    function decodeStateTransitionInputs(
        bytes memory _encodedInputs
    ) internal pure returns (uint256[10] memory, bytes memory, bytes memory) {
        return abi.decode(_encodedInputs, (uint256[10], bytes, bytes));
    }

    function encodeStateTransitionInputs(
        uint256[10] memory inputs,
        bytes memory blobCommitment,
        bytes memory blobProof
    ) public pure returns (bytes memory) {
        return abi.encode(inputs, blobCommitment, blobProof);
    }

    function stateTransitionInputs() internal pure returns (bytes memory) {
        return
            encodeStateTransitionInputs(
                [
                    ROOT_HASH_BEFORE,
                    ROOT_HASH_AFTER,
                    VOTERS_COUNT,
                    OVERWRITTEN_VOTES_COUNT,
                    CENSUS_ROOT,
                    BLOB_EVALUATION_POINT_Z,
                    BLOB_EVALUATION_POINT_Y_L1,
                    BLOB_EVALUATION_POINT_Y_L2,
                    BLOB_EVALUATION_POINT_Y_L3,
                    BLOB_EVALUATION_POINT_Y_L4
                ],
                BLOB_COMMITMENT,
                BLOB_PROOF
            );
    }

    function resultsInputs() internal view returns (bytes memory) {
        return
            abi.encode(
                ROOT_HASH_AFTER, // After state root
                FINAL_RESULTS[0],
                FINAL_RESULTS[1],
                FINAL_RESULTS[2],
                FINAL_RESULTS[3],
                FINAL_RESULTS[4],
                FINAL_RESULTS[5],
                FINAL_RESULTS[6],
                FINAL_RESULTS[7]
            );
    }
}
