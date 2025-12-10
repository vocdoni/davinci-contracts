// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { TestHelpers } from "./TestHelpers.t.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test, TestHelpers {
    uint256[8] public proof;
    uint256[2] public proofCommitments;
    uint256[2] public proofCommitmentPok;

    StateTransitionVerifierGroth16 public stv;

    function setUp() public {
        stv = new StateTransitionVerifierGroth16();
        (proof, proofCommitments, proofCommitmentPok) = decodedProof();
    }

    function test_Verify_OK() public view {
        uint256[8] memory input = [
            ROOT_HASH_BEFORE,
            ROOT_HASH_AFTER,
            VOTERS_COUNT,
            OVERWRITTEN_VOTES_COUNT,
            CENSUS_ROOT,
            BLOBS_COMMITMENT_L1,
            BLOBS_COMMITMENT_L2,
            BLOBS_COMMITMENT_L3
        ];
        stv.verifyProof(proof, proofCommitments, proofCommitmentPok, input);
    }

    function test_Verify_OK_ABIEncoded() public view {
        stv.verifyProof(STATETRANSITION_ABI_PROOF, stateTransitionInputs());
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(
            STATETRANSITION_ABI_PROOF
        );
        uint256[8] memory inputBad = [
            ROOT_HASH_BEFORE,
            ROOT_HASH_AFTER_BAD,
            VOTERS_COUNT,
            OVERWRITTEN_VOTES_COUNT,
            CENSUS_ROOT,
            BLOBS_COMMITMENT_L1,
            BLOBS_COMMITMENT_L2,
            BLOBS_COMMITMENT_L3
        ];
        vm.expectRevert();
        stv.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
    }

    function test_Encode_Proof() public view {
        bytes memory encodedProof = encodeProof(proof, proofCommitments, proofCommitmentPok);
        if (keccak256(encodedProof) != keccak256(STATETRANSITION_ABI_PROOF)) {
            revert();
        }
    }

    function test_Decode_Proof() public view {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(
            STATETRANSITION_ABI_PROOF
        );
        if (
            _proof[0] != proof[0] ||
            _proof[1] != proof[1] ||
            _proof[2] != proof[2] ||
            _proof[3] != proof[3] ||
            _proof[4] != proof[4] ||
            _proof[5] != proof[5] ||
            _proof[6] != proof[6] ||
            _proof[7] != proof[7] ||
            _commitments[0] != proofCommitments[0] ||
            _commitments[1] != proofCommitments[1] ||
            _commitmentPok[0] != proofCommitmentPok[0] ||
            _commitmentPok[1] != proofCommitmentPok[1]
        ) {
            revert();
        }
    }

    function test_Decode_Inputs() public pure {
        uint256[8] memory inputs = decodeStateTransitionInputs(stateTransitionInputs());
        if (
            inputs[0] != ROOT_HASH_BEFORE ||
            inputs[1] != ROOT_HASH_AFTER ||
            inputs[2] != VOTERS_COUNT ||
            inputs[3] != OVERWRITTEN_VOTES_COUNT ||
            inputs[4] != CENSUS_ROOT ||
            inputs[5] != BLOBS_COMMITMENT_L1 ||
            inputs[6] != BLOBS_COMMITMENT_L2 ||
            inputs[7] != BLOBS_COMMITMENT_L3
        ) {
            revert();
        }
    }

    function decodeProof(
        bytes memory encodedProof
    ) public pure returns (uint256[8] memory, uint256[2] memory, uint256[2] memory) {
        return abi.decode(encodedProof, (uint256[8], uint256[2], uint256[2]));
    }

    function encodeProof(
        uint256[8] memory _proof,
        uint256[2] memory _commitments,
        uint256[2] memory _commitmentsPok
    ) public pure returns (bytes memory) {
        return abi.encode(_proof, _commitments, _commitmentsPok);
    }
}
