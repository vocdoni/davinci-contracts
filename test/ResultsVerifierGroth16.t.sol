// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 5244826649985884941716637507837432287554322849727016690652948098167443117049;
    uint256[8] public results = [52,57,23,42,46,0,0,0];

    bytes public zkp =
        hex"0487ebaa3d2f48b09a97fa9be27d15cd945624ab328056840e844fb8458b6f0d0305178df1ffbb419bb26d05302ebf8a2161b5868f99d153491456456bfae3821ba68ec7a1c8b968a1322eb402f9f8d0d6776f16ee557f560f86efaa78c3a00a2f1d11ef7bd8492f28ce881360c70f227ce55d8eec55cc9b5ecd0be024a976332ff734345491ec7f98ef341c29b0f65ddb31d1e9f8a2961f649056f383b6c36302bd58fb2414dc5c911c215f740fe34d3349d21e9a4dfc35944d5be1dedeb1e02c66410e6938b9b64a6170b25e0a21db9df7448cdda519290c86e59462f606f90210bcfb97de66bda8af8db0e1d9e4de84892e1cb01cf7661d147977543c65142a7dd4d5a23eaf0f75ca070c83822673c71d7ce76e5cbb8f7b7e56a29b96519005ad26c999fb1bc370771e1d43d4ba1261f569d0bdad74ff48db02ff12e56f3214e341fae2089cedbace782230ea1749af3060e36e6f139331a52d63402203bd0ac45b18912576cec84274054d72f3a9553df038a1e8fb4ab9d5f5036fe79cc1";

    bytes public encodedInputs =
        hex"0b987778b1b51747b94a15831946b0eb2431ad7c2a3608c733419986189553f9000000000000000000000000000000000000000000000000000000000000003400000000000000000000000000000000000000000000000000000000000000390000000000000000000000000000000000000000000000000000000000000017000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        2049402248104077914516084402596252830972634856675548365760403081179882549005,
        1365935348077611454870287869773443334654271550287396759904635794379603895170
    ];
    uint256[2][2] public BS = [
        [
            12506728954549887434638765373781452599751235338171802889112003233557554372618,
            21310066234449629011520918970515856822101472132274983391245360107149459027507
        ],
        [
            21695475410047303062785529163276019181597901162795812411947216950777823544163,
            1239173916834602569737619661004988862989533757304107023382404736896179548640
        ]
    ];
    uint256[2] public KRS = [
        20082432740304037133594412658822527777662320707937159605950071705145510856441,
        934199561456443649894974504433649813662733031955903739033387843745811817748
    ];

    uint256[2] public commitments = [
        19219464453375023937633935686716404541370248734284219665457485505988845261200,
        2567496486649531219967882366359362169433328808830282403043328697209941487410
    ];

    uint256[2] public commitmentPok = [
        9447786632673350285385088514303454831880971734858001503737892113245809017789,
        4870059231771199291030123507815313811825828234183500691230092082262026198209
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
