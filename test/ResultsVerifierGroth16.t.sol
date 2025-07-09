// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 2992761243883063123620743517261721469964951289365933052491486976045184248823;
    uint256[8] public results = [23, 36, 38, 55, 33, 0, 0, 0];

    bytes public zkp =
        hex"2cbd14d18ed1e7ae904842e1cc0b8fac12e6211392794f27e019b198c8667b602d83ce77a182550ac69d124180fc0664b361211869707542ac359cc1ea3cf7cd2d06d51b5cdb5b8dc4d2a9402bdec565ae45ec08bcec9cecfe01807e049b5dfd2b563d00fd4ef607fd65a1c82a503cc15013ff836073220d8b6c02e4c4ac3d6e1a6c0097cc576956414169aac6c7e84daa5bf9faac07540c60337dda132c794e18db06472efa5b5c654427a9063fdcde7ec9546ad10483a46e0447ded5f9cc9716bd19b90fc1496c43cd2fe18f05af580f57f105283a991aeda069e008132e441cc3878c1badb57f9329196b35c13b81e4175501ea638d1cdb7fee672e100d2b1ba697939ea73551eb6dbce4377f50a6e4e20f188c362589ba23bbed6c36235202c565cd10d2551037b242d8df89e7c160bba28409dabf81ced41d0b2257a63a17e99dbbdcae0c2f6fc086f81e70e6735af53395fc3882f6000c091b158c1bf225736e8705d16904fd030bebc04446edc741cf41e46457a9c63b6e1acaf488ce";

    bytes public encodedInputs =
        hex"069dd7c4222d0d9d59910d579fd511f91857626506ea8456aaf8fe201da0ebf700000000000000000000000000000000000000000000000000000000000000170000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000002600000000000000000000000000000000000000000000000000000000000000370000000000000000000000000000000000000000000000000000000000000021000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        20235843117503320300973842346998433773889476415573771254995124423610759281504,
        20586960136722942585166762245304640559532157721536771035482556432741811681229
    ];
    uint256[2][2] public BS = [
        [
            20366150078305053515211137861547145387157750853251048474825477950834644573693,
            19601822369855042464712743962379581168933250131842627201459342792698836565358
        ],
        [
            11950957638632604839480892705159424948608161254425536062526811110210875980110,
            11242491202766509863141305955298813170957486139024442075979721260811563945111
        ]
    ];
    uint256[2] public KRS = [
        10284994296982976325453235187540919721315232769610857264935211618738294173252,
        13010230451027485549664177959147027165855941286583391556306301386989851643179
    ];

    uint256[2] public commitments = [
        12506789668020057311901008349999462271063468695176738964599296971360662070098,
        1253397173869489449987584929193363207532499730239479879708426068535885145658
    ];

    uint256[2] public commitmentPok = [
        10815959522435266553410151593929182472491395622049042404805037805608104631282,
        16939525642334025426337770080010264078081631846788590122854205312380121155790
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
