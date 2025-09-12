// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 15977845708338371169793025518832793983533753933654506631846328160167915014374;
    uint256[8] public results = [43,40,30,39,22,0,0,0];

    bytes public zkp =
        hex"15aa1ee6ff35372b147cee8f0e121ae25ef3476e92e0cd5718ec2a0caeb7b4ef242bee8c88921a3fce5a890b39aeeeab80f322d52e3082e1555f9b14923440380e0025ec9497cdb89aed33a13a8b67d72df976d84c9220affb6f2c61c7b9a83f28c5b0dc3028254d51c7e4eff4c061503f0361936a745796c324efb512fcf23026a42cba6afe122c663f6146d62f38d972473021b1e1b2362153e58e71301f47304fbc8468b18e63ffbfd5355dbca2da61ddeec8f45ba0a373eb04d3ec7e61c2221ef2e53f722f7b85957542f757f88cfd5d588703d981ce382492e75fd8e31d03228309a0a43e02546f0de2ca6d948b0d4b25d8437301c34d63cc1bb98c1e190bf56a09d26f2822a73baa19c8e7374458743e16f3ca83cd0ccf03f07a240f4d06108f8a90b86a5f9bc63631892df28abd82b3c5edd7b36defe12e5e452ff1600ce13df7091e8425165345659abea13b425e6e7a5f6ec8b451f6d02ec1745cc3018611fb2f8d55a6b1f10f3067656b4d9bc000ff8167d794bcea6503738ce9eb";

    bytes public encodedInputs =
        hex"235323e3c2d30ff0c284fe50e4fa3367828ae92a9730d6c2780c129830d9c8e6000000000000000000000000000000000000000000000000000000000000002b0000000000000000000000000000000000000000000000000000000000000028000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000000000000000000000000000000000000000000270000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        9799147101315582616306505878295260602970356177657922463323496534912768914671,
        16360883377188704827055070623522173382394007840625338866182181232533782675512
    ];
    uint256[2][2] public BS = [
        [
            6332641622976641680707584853478046621470148577963222472029007185741279242303,
            18441803458708804147941907878240603812777580722772389559816853027906926735920
        ],
        [
            17477959867444738322904962526588601155617435664731807770582697884577088675655,
            21851898748165956416930853928693050813051748481801892771800601049002860765634
        ]
    ];
    uint256[2] public KRS = [
        15433318666899766384963409068883870091359528154196103669917759127644532171549,
        1417915734280745799140294438253588678883756994552243023119602370371604192793
    ];

    uint256[2] public commitments = [
        5409050715200222733878144727625714591714259533954134946581687574875504643917,
        2743137329977117742409113082293592573628886385045964563886568411515627499872
    ];

    uint256[2] public commitmentPok = [
        5825722439168678403054138999893769683649922124297620786344535511134880685251,
        689194456905898163095987522435865013302938301831268466669224674673633847787
    ];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];
    uint256[9] public inputs = [
        STATE_ROOT,
        results[0],
        results[1],
        results[2],
        results[3],
        results[4],
        results[5],
        results[6],
        results[7]
    ];

    ResultsVerifierGroth16 public rv;

    function setUp() public {
        rv = new ResultsVerifierGroth16();
    }

    function test_Verify_OK() public view {
        rv.verifyProof(proof, commitments, commitmentPok, inputs);
    }

    function test_Verify_OK_ABIEncoded() public view {
        rv.verifyProof(zkp, encodedInputs);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[9] memory inputBad = [
            STATE_ROOT,
            results[0],
            results[1],
            results[2],
            results[3],
            results[4],
            results[5],
            results[6],
            999 // Bad input
        ];
        vm.expectRevert();
        rv.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
    }

    function test_Encode_Proof() public view {
        bytes memory encodedProof = encodeProof(proof, commitments, commitmentPok);
        if (keccak256(encodedProof) != keccak256(zkp)) {
            revert();
        }
    }

    function test_Decode_Proof() public view {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
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

    function test_Encode_Inputs() public view {
        bytes memory _encodedInputs = encodeInputs(
            [STATE_ROOT, results[0], results[1], results[2], results[3], results[4], results[5], results[6], results[7]]
        );
        if (keccak256(_encodedInputs) != keccak256(encodedInputs)) {
            revert();
        }
    }

    function test_Decode_Inputs() public view {
        uint256[9] memory _inputs = decodeInputs(encodedInputs);
        if (
            _inputs[0] != STATE_ROOT ||
            _inputs[1] != results[0] ||
            _inputs[2] != results[1] ||
            _inputs[3] != results[2] ||
            _inputs[4] != results[3] ||
            _inputs[5] != results[4] ||
            _inputs[6] != results[5] ||
            _inputs[7] != results[6] ||
            _inputs[8] != results[7]
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

    function decodeInputs(bytes memory _encodedInputs) public pure returns (uint256[9] memory) {
        uint256[9] memory _inputs = abi.decode(_encodedInputs, (uint256[9]));
        return _inputs;
    }

    function encodeInputs(uint256[9] memory _inputs) public pure returns (bytes memory) {
        return abi.encode(_inputs);
    }
}
