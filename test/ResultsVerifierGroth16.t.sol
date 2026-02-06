// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {TestHelpers} from "./TestHelpers.t.sol";
import {ResultsVerifierGroth16} from "../src/verifiers/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test, TestHelpers {
    uint256[8] public proof;
    uint256[2] public proofCommitments;
    uint256[2] public proofCommitmentPok;
    uint256[9] public inputs = [
        ROOT_HASH_AFTER,
        FINAL_RESULTS[0],
        FINAL_RESULTS[1],
        FINAL_RESULTS[2],
        FINAL_RESULTS[3],
        FINAL_RESULTS[4],
        FINAL_RESULTS[5],
        FINAL_RESULTS[6],
        FINAL_RESULTS[7]
    ];

    ResultsVerifierGroth16 public rv;

    function setUp() public {
        rv = new ResultsVerifierGroth16();
        (proof, proofCommitments, proofCommitmentPok) = decodedResultsProof();
    }

    function test_Verify_OK() public view {
        rv.verifyProof(proof, proofCommitments, proofCommitmentPok, inputs);
    }

    function test_Verify_OK_ABIEncoded() public view {
        rv.verifyProof(RESULTS_ABI_PROOF, RESULTS_ABI_INPUTS);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) =
            decodeProof(RESULTS_ABI_PROOF);
        uint256[9] memory inputBad = [
            ROOT_HASH_AFTER,
            FINAL_RESULTS[0],
            FINAL_RESULTS[1],
            FINAL_RESULTS[2],
            FINAL_RESULTS[3],
            FINAL_RESULTS[4],
            FINAL_RESULTS[5],
            FINAL_RESULTS[6],
            999 // Bad input
        ];
        vm.expectRevert();
        rv.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
    }

    function test_Encode_Proof() public view {
        bytes memory encodedProof = encodeProof(proof, proofCommitments, proofCommitmentPok);
        if (keccak256(encodedProof) != keccak256(RESULTS_ABI_PROOF)) {
            revert();
        }
    }

    function test_Decode_Proof() public view {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) =
            decodeProof(RESULTS_ABI_PROOF);
        if (
            _proof[0] != proof[0] || _proof[1] != proof[1] || _proof[2] != proof[2] || _proof[3] != proof[3]
                || _proof[4] != proof[4] || _proof[5] != proof[5] || _proof[6] != proof[6] || _proof[7] != proof[7]
                || _commitments[0] != proofCommitments[0] || _commitments[1] != proofCommitments[1]
                || _commitmentPok[0] != proofCommitmentPok[0] || _commitmentPok[1] != proofCommitmentPok[1]
        ) {
            revert();
        }
    }

    function test_Encode_Inputs() public view {
        bytes memory _encodedInputs = encodeInputs(
            [
                ROOT_HASH_AFTER,
                FINAL_RESULTS[0],
                FINAL_RESULTS[1],
                FINAL_RESULTS[2],
                FINAL_RESULTS[3],
                FINAL_RESULTS[4],
                FINAL_RESULTS[5],
                FINAL_RESULTS[6],
                FINAL_RESULTS[7]
            ]
        );
        if (keccak256(_encodedInputs) != keccak256(RESULTS_ABI_INPUTS)) {
            revert();
        }
    }

    function test_Decode_Inputs() public view {
        uint256[9] memory _inputs = decodeInputs(RESULTS_ABI_INPUTS);
        if (
            _inputs[0] != ROOT_HASH_AFTER || _inputs[1] != FINAL_RESULTS[0] || _inputs[2] != FINAL_RESULTS[1]
                || _inputs[3] != FINAL_RESULTS[2] || _inputs[4] != FINAL_RESULTS[3] || _inputs[5] != FINAL_RESULTS[4]
                || _inputs[6] != FINAL_RESULTS[5] || _inputs[7] != FINAL_RESULTS[6] || _inputs[8] != FINAL_RESULTS[7]
        ) {
            revert();
        }
    }

    function decodeProof(bytes memory encodedProof)
        public
        pure
        returns (uint256[8] memory, uint256[2] memory, uint256[2] memory)
    {
        return abi.decode(encodedProof, (uint256[8], uint256[2], uint256[2]));
    }

    function encodeProof(uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentsPok)
        public
        pure
        returns (bytes memory)
    {
        return abi.encode(_proof, _commitments, _commitmentsPok);
    }

    function decodeInputs(bytes memory _encodedInputs) public pure returns (uint256[9] memory) {
        uint256[9] memory _inputs = abi.decode(_encodedInputs, (uint256[9]));
        return _inputs;
    }

    function encodeInputs(uint256[9] memory _inputs) public pure returns (bytes memory) {
        return abi.encode(_inputs);
    }
}
