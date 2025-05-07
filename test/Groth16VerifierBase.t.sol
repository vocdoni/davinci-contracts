// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { Groth16VerifierBase } from "../src/Groth16VerifierBase.sol";

contract Groth16VerifierBaseTest is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        20375616703360125164212336204194149108274950832458075705859486132212778035627;
    uint256 public constant ROOT_HASH_AFTER =
        10499356723460278798789300931775850701687590422599810173373225874452881489710;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 3;
    uint256 public constant NUM_OVERWRITES = 0;

    bytes public zkp =
        hex"18e15ca675fedf8b3bf0595227edb7fb1019866b7c758bee421534ec672601b52fb8bef9285ddbabb85b808e611d9c42923e3f610f7cb4ab894c536e26705e4728a63ee2c5b23239cf924c2c39e02ff1a73aa88e378b962c95df98336f9af1210f7898e923dd2c32a8b4d71070f0bb5287930e36e13a5917e0ccf80db974adb415b605c7edef99b41295fae42fae560c435a345ff12511fa3450e00d9c43d959174d011cc13a371a0cd663f8a3d7434f0b1e4ec57ca13d77e6a325f2314c351e26b6a7d1cff127a1131bd448f0f5f171b67fc39f9181687e66de1b902f89377e03744b938fab17e34aab9ca9bb4f9d929d18c6d293f82ce43a99ef20aac5b586190dc272796d608f6cd8867655198cc559db49c3b212152908bd3957c09133021a800434201d0e59b4bb47adff7b45bf8f9192beb283f8fab56ad02ba72657e21f96780780f62e9518076c2c4946d64de0406d1d95307d5af3d379861053134b25f62690ab6b52c61bbfc7dfd08ebd742603a51df9c31c982c1f15752fd55354";

    uint256[2] public AR = [
        11253688404014968159772466738023582338955280243113054325314952781115320304053,
        21585121792416455942855009762436828878986263224192760691969876572793277013575
    ];
    uint256[2][2] public BS = [
        [
            18386244578125142245936417410030574241458347909667625033840923832667929768225,
            6997769727411602915401821730171064614881431549175035558477351242986905972148
        ],
        [
            9820175884856959834223027075277544655488906737678871158393645421635293141337,
            10539250418377136649970618837582675240262172950076740189641926357972266267934
        ]
    ];
    uint256[2] public KRS = [
        17510612660121405871442286886251756691360659300438955779488570169411680941950,
        1562414414482302525516925440125640794822858576159351220681329994866731890054
    ];

    uint256[2] public commitments = [
        11332132251436758587594496578113083270201286933368353577478692157115333686018,
        11986319499741119447642117118045353912795545577046074224057908070650799871970
    ];

    uint256[2] public commitmentPok = [
        14287553777660503945597472106287272023510506384335607203097576359475960550219,
        17170485942162389160704339116747716169786379114770058377976135872541837579092
    ];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    Groth16VerifierBase public groth16Verifier;

    function setUp() public {
        groth16Verifier = new Groth16VerifierBase();
    }

    function test_Verify_OK() public view {
        uint256[4] memory input = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES];
        groth16Verifier.verifyProof(proof, commitments, commitmentPok, input);
    }

    function test_Verify_OK_ABIEncoded() public view {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory input = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES];
        groth16Verifier.verifyProof(_proof, _commitments, _commitmentPok, input);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory inputBad = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER_BAD, NUM_NEW_VOTES, NUM_OVERWRITES];
        vm.expectRevert();
        groth16Verifier.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
    }

    function test_Encode_Proof() public {
        bytes memory encodedProof = encodeProof(proof, commitments, commitmentPok);
        if (keccak256(encodedProof) != keccak256(zkp)) {
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
