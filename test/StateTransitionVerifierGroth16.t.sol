// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { TestHelpers } from "./TestHelpers.t.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test, TestHelpers {
    uint256[2] public AR = [16821715178840840010892626430339508771016478604624368433428533629029610442542,12486171747568638521436034082922672340800029125892437664344067230870987119739];
    uint256[2][2] public BS = [
        [11145273543394870856453188271128080984772592160473742252644548533585700819476,15073905004599226830723769206496134595497691471246689583573915738305897583990],
        [502743393330190083286856071889872880420097829691848689011348302863014303677,426292576433938025205741738608094531209815498482222725480951279053974178417]
    ];
    uint256[2] public KRS = [17011604122933596794694590088776900835322036015143446466806262738302613031464,15809843204291845503295719170577272173663599265499454552806477828222436370684];
    uint256[2] public commitments = [17605483704596221426977260580699441727749524076885295716915316826890106228347,19152488088795007058617697574732346239647544293953316609824002368177217124627];
    uint256[2] public commitmentPok = [4597548851566745673753323113523989027666149648679005935596369930070879386604,12459379127365388220456249246602053014934428333995447479926171432613016526059];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    StateTransitionVerifierGroth16 public stv;

    function setUp() public {
        stv = new StateTransitionVerifierGroth16();
    }

    function test_Verify_OK() public view {
        uint256[9] memory input = [
            ROOT_HASH_BEFORE,
            ROOT_HASH_AFTER,
            NUM_NEW_VOTES,
            NUM_OVERWRITES,
            BLOB_EVALUATION_POINT_Z,
            BLOB_EVALUATION_POINT_Y_L1,
            BLOB_EVALUATION_POINT_Y_L2,
            BLOB_EVALUATION_POINT_Y_L3,
            BLOB_EVALUATION_POINT_Y_L4
        ];
        stv.verifyProof(proof, commitments, commitmentPok, input);
    }

    function test_Verify_OK_ABIEncoded() public view {
        stv.verifyProof(stateTransitionZKProof, stateTransitionInputs());
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(
            stateTransitionZKProof
        );
        uint256[9] memory inputBad = [
            ROOT_HASH_BEFORE,
            ROOT_HASH_AFTER_BAD,
            NUM_NEW_VOTES,
            NUM_OVERWRITES,
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
        bytes memory encodedProof = encodeProof(proof, commitments, commitmentPok);
        if (keccak256(encodedProof) != keccak256(stateTransitionZKProof)) {
            revert();
        }
    }

    function test_Decode_Proof() public view {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(
            stateTransitionZKProof
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
            _commitments[0] != commitments[0] ||
            _commitments[1] != commitments[1] ||
            _commitmentPok[0] != commitmentPok[0] ||
            _commitmentPok[1] != commitmentPok[1]
        ) {
            revert();
        }
    }

    function test_Decode_Inputs() public pure {
        (uint256[9] memory inputs, bytes memory blobCommitment, bytes memory blobProof) = decodeStateTransitionInputs(
            stateTransitionInputs()
        );
        if (
            inputs[0] != ROOT_HASH_BEFORE ||
            inputs[1] != ROOT_HASH_AFTER ||
            inputs[2] != NUM_NEW_VOTES ||
            inputs[3] != NUM_OVERWRITES ||
            inputs[4] != BLOB_EVALUATION_POINT_Z ||
            inputs[5] != BLOB_EVALUATION_POINT_Y_L1 ||
            inputs[6] != BLOB_EVALUATION_POINT_Y_L2 ||
            inputs[7] != BLOB_EVALUATION_POINT_Y_L3 ||
            inputs[8] != BLOB_EVALUATION_POINT_Y_L4 ||
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
