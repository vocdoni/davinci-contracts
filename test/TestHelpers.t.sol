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

    function decodeStateTransitionInputs(bytes memory _encodedInputs) internal pure returns (uint256[8] memory) {
        return abi.decode(_encodedInputs, (uint256[8]));
    }

    function encodeStateTransitionInputs(uint256[8] memory inputs) public pure returns (bytes memory) {
        return abi.encode(inputs);
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
                    BLOBS_COMMITMENT_L1,
                    BLOBS_COMMITMENT_L2,
                    BLOBS_COMMITMENT_L3
                ]
            );
    }

    function decodedResultsProof()
        internal
        pure
        returns (uint256[8] memory proofArr, uint256[2] memory commitmentsArr, uint256[2] memory commitmentPokArr)
    {
        return abi.decode(RESULTS_ABI_PROOF, (uint256[8], uint256[2], uint256[2]));
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

    function decodedProof()
        internal
        pure
        returns (uint256[8] memory proofArr, uint256[2] memory commitmentsArr, uint256[2] memory commitmentPokArr)
    {
        return abi.decode(STATETRANSITION_ABI_PROOF, (uint256[8], uint256[2], uint256[2]));
    }
}
