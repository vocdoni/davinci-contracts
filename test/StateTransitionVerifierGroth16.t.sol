// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { TestHelpers } from "./TestHelpers.t.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test, TestHelpers {
    uint256[2] public AR = [
        3814216515795987304425536225738282214029140770405826833670205149905004929569,
        202919102913338390746368395023430198738544407915768597387248569680763414954
    ];
    uint256[2][2] public BS = [
        [
            3085217683851624874572489424510466884830826468295439930305666786085039345795,
            14015798750296761045037332208417144969347191778185846849471509950310617924776
        ],
        [
            1393105205277051362900736523413933661953784030994053699268672329907648405707,
            2076801956973884389476767914964940326682848543248715239384134424639052823201
        ]
    ];
    uint256[2] public KRS = [
        19283135182921776480761614663756925720200504521362709277553488470944584230511,
        12727195824737883589367200800442034951311006585401605732932457739509730268031
    ];

    uint256[2] public commitments = [
        12333879902195505304413058401218737582248800433653263354065103799058872076534,
        1396049335338488636819468598792448300617682374309797936689178871545065860993
    ];

    uint256[2] public commitmentPok = [
        16895532588526411029313615673529068826170856786148716600736643288362606652302,
        4593762621572348507191867426123472751786865384460733196614883907539831601359
    ];

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

    function test_Decode_Inputs() public view {
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
