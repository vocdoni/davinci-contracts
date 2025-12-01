// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { TestHelpers } from "./TestHelpers.t.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test, TestHelpers {
    uint256[8] public proof = [STATETRANITION_PROOF_AR[0], STATETRANITION_PROOF_AR[1], STATETRANITION_PROOF_BS[0][0], STATETRANITION_PROOF_BS[0][1], STATETRANITION_PROOF_BS[1][0], STATETRANITION_PROOF_BS[1][1], STATETRANITION_PROOF_KRS[0], STATETRANITION_PROOF_KRS[1]];

    StateTransitionVerifierGroth16 public stv;

    function setUp() public {
        stv = new StateTransitionVerifierGroth16();
    }

    function test_Verify_OK() public view {
        uint256[10] memory input = [
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
        ];
        stv.verifyProof(proof, STATETRANITION_PROOF_COMMITMENTS, STATETRANITION_PROOF_COMMITMENTSPOK, input);
    }

    function test_Verify_OK_ABIEncoded() public view {
        stv.verifyProof(STATETRANSITION_ABI_PROOF, stateTransitionInputs());
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(
            STATETRANSITION_ABI_PROOF
        );
        uint256[10] memory inputBad = [
            ROOT_HASH_BEFORE,
            ROOT_HASH_AFTER_BAD,
            VOTERS_COUNT,
            OVERWRITTEN_VOTES_COUNT,
            CENSUS_ROOT,
            BLOB_EVALUATION_POINT_Z,
            BLOB_EVALUATION_POINT_Y_L1,
            BLOB_EVALUATION_POINT_Y_L2,
            BLOB_EVALUATION_POINT_Y_L3,
            BLOB_EVALUATION_POINT_Y_L4
        ];
        vm.expectRevert();
        stv.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
    }

    function test_Encode_Proof() public view {
        bytes memory encodedProof = encodeProof(proof, STATETRANITION_PROOF_COMMITMENTS, STATETRANITION_PROOF_COMMITMENTSPOK);
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
            _commitments[0] != STATETRANITION_PROOF_COMMITMENTS[0] ||
            _commitments[1] != STATETRANITION_PROOF_COMMITMENTS[1] ||
            _commitmentPok[0] != STATETRANITION_PROOF_COMMITMENTSPOK[0] ||
            _commitmentPok[1] != STATETRANITION_PROOF_COMMITMENTSPOK[1]
        ) {
            revert();
        }
    }

    function test_Decode_Inputs() public pure {
        (uint256[10] memory inputs, bytes memory blobCommitment, bytes memory blobProof) = decodeStateTransitionInputs(
            stateTransitionInputs()
        );
        if (
            inputs[0] != ROOT_HASH_BEFORE ||
            inputs[1] != ROOT_HASH_AFTER ||
            inputs[2] != VOTERS_COUNT ||
            inputs[3] != OVERWRITTEN_VOTES_COUNT ||
            inputs[4] != CENSUS_ROOT ||
            inputs[5] != BLOB_EVALUATION_POINT_Z ||
            inputs[6] != BLOB_EVALUATION_POINT_Y_L1 ||
            inputs[7] != BLOB_EVALUATION_POINT_Y_L2 ||
            inputs[8] != BLOB_EVALUATION_POINT_Y_L3 ||
            inputs[9] != BLOB_EVALUATION_POINT_Y_L4 ||
            keccak256(blobCommitment) != keccak256(BLOB_COMMITMENT) ||
            keccak256(blobProof) != keccak256(BLOB_PROOF)
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
