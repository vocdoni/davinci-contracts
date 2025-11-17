// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 3521621988802002318964005439092080066299578336508168327546544274269862942207;
    uint256[8] public results = [38,27,31,40,43,0,0,0];

    bytes public zkp =
        hex"108234851b33f21350c393d69e715e271e094df428ee33d230a2544a9c5ec8830277e616826d2c621f84870dd4dbf33c305d7fb03e3504916b0fdf9070960a6b08d24caa523ff0656e3302624ce4c06c970fb26e0bb840c317055b972c0b652d109ef4dc0199dd71c6e2bed91b9f6f31db9d369e03979e4b17eb1ef737940617083e406618446daa26a7aea207487c86201a6577715255517b92c23122b07e390e9aa87cfb3bd7aeb2ae22f7247580d2e0eef6b4d72f06c201a79bf9bc548ffc0f5035df170ca66867ab15c9310c1f96cb156b5b1a53880ff54059b8158e1754190a2af85a93d80ed07a764ee73e5ded58d140e5bb9441f69fd8915ba9d1b6b2155e9806ae050dde7a18b1ee1d6e59eed92ce5078991f94e5c5b6e2e66acb5451852249589e4acf68c5febabbb7c7100ceefe4a44e9d6952dccdebf0c9afacf911afc41382447213a6117beb4f66f4d944de6df075565f9a0df08ca579401c1d0511a058689c081315116cae46c20a03aac0e2cbd11ab68ce646f888c3bcd394";

    bytes public encodedInputs =
        hex"07c92adb69a073e49ea2072e41e4cf7cf97ce16b718f77b2f5e7e300fd38d5ff0000000000000000000000000000000000000000000000000000000000000026000000000000000000000000000000000000000000000000000000000000001b000000000000000000000000000000000000000000000000000000000000001f0000000000000000000000000000000000000000000000000000000000000028000000000000000000000000000000000000000000000000000000000000002b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [7467058175101193075579914496640452811439592031508033561170199174833623713923,1116468506389258256254412204529285418627982982505132556835824942734164494955];
    uint256[2][2] public BS = [
        [3990069796844792230075169501045171923308761130675441400514032861467653137709,7517857371032739297235270906411316705898095458006490497904852425779086362135],
        [3728491770918790382690182885108739833291986664955253938977747649996921011769,6605637191019251154682775197464935752114573208343418748936660255106854588412]
    ];
    uint256[2] public KRS = [6926412300983139875564580189011055947371627129284352627956764262473938114388,11325786254181707943927238381962466692251429975232616515974810027999620085426];
    uint256[2] public commitments = [9665702689868572369150725406313179013078209470522994058553800567394189227333,11000642319732614361039547448429285254571083162828417368944743149787618979065];
    uint256[2] public commitmentPok = [7999869930493458295048937164091527582232806375233523785346930894985945422877,2292707305925023661483177287141133877224598486820036179303293641587035001748];

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
