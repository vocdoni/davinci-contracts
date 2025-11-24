// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { TestHelpers } from "./TestHelpers.t.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test, TestHelpers {
    uint256[2] public AR = [
        4579420157231911784964202373132776012336161807136104787020948288813701866933,
        14920430514456108787722407652316251876343953536221764801998793348747353418249
    ];
    uint256[2][2] public BS = [
        [
            19189805169944668669556052971383810812122136623144153950620952453939755398546,
            20758673655803804563919822603120705259834715681020212050226179998040274112088
        ],
        [
            13372889941645865729257104070193754065630859438749225570725180618775193790028,
            6350632104835825638006224825105889936572254911831187334664884041743238221815
        ]
    ];
    uint256[2] public KRS = [
        1450007305552418548704388167822650206126597198961540624468425709256481223170,
        1285798123240657209313199623079649355230757187107621425965425237257357155228
    ];

    uint256[2] public commitments = [
        1844664545958878781001959220571109888530626708801953169192869296864036898070,
        3887145990917029470600540643219114194700460468289610204550951311703752189954
    ];

    uint256[2] public commitmentPok = [
        10837860982295682743045677768042985917843229244233326821781100140220415860989,
        18028336241621685322163939597823486900555359699241746491850812600011955855496
    ];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    StateTransitionVerifierGroth16 public stv;

    function setUp() public {
        stv = new StateTransitionVerifierGroth16();
    }

    function test_Verify_OK() public view {
        uint256[10] memory input = [
            ROOT_HASH_BEFORE,
            ROOT_HASH_AFTER,
            NUM_NEW_VOTES,
            NUM_OVERWRITES,
            CENSUS_ROOT,
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
        uint256[10] memory inputBad = [
            ROOT_HASH_BEFORE,
            ROOT_HASH_AFTER_BAD,
            NUM_NEW_VOTES,
            NUM_OVERWRITES,
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
        (uint256[10] memory inputs, bytes memory blobCommitment, bytes memory blobProof) = decodeStateTransitionInputs(
            stateTransitionInputs()
        );
        if (
            inputs[0] != ROOT_HASH_BEFORE ||
            inputs[1] != ROOT_HASH_AFTER ||
            inputs[2] != NUM_NEW_VOTES ||
            inputs[3] != NUM_OVERWRITES ||
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
