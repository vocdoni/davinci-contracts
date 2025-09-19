// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        20537488792664972079444046972623508603551134024737764656077520772428338065655;
    uint256 public constant ROOT_HASH_AFTER =
        3886101432899391880372659930436735438625666716230583631467284935225407559329;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 5;
    uint256 public constant BLOB_EVALUATION_POINT_Z = 1326229533887009909157123213737668090348869338874577013664169420100719045278;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L1 = 16261872794105072506;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L2 = 1519966843562379078;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L3 = 1690861515513695806;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L4 = 6100054667831901601;
    bytes public BLOB_COMMITMENT = hex"b550f3efaa76c4105356129177873a1651a95602e7b953d7a1a9f54393607d05b62ea1b0fd9302768edcd3ed22a43688";
    bytes public BLOB_PROOF = hex"872cdf0112e5c8ed5571b68834ed5a6b6c87c7970f71f26696157fe9c561a1d9628d8b15206599ddbc3e596261012288";

    bytes public zkp =
        hex"0b9e5e5ac2155e0dfccdc50d65730e7c03a5ab41ca1726f955ade81007bceac42ba37b28cffa631a4c3bdef83a2bb765a109bdef37f1bc87d7ce91ab757d1b7a12029e6f7cd00004475c6e8199c8728d6690642237160a7abc91b484da5ea19f06deb616c029bc86ee55ec76c1ecf9fab10c14f54fe93ff3a9a5202cf48db854010ad415448b2e2f08d8df52fc5f0001e98645737eb76b9f993ae81a1ac493fb2e083074e0c0c34af14fb92c9fd3ad496474b8881b4c3b71e6e6da4852279a3d2b169aa283ca5ea903af59e05f4d66050cfeb731e429c86b7795372aeb682dd1073eae964860ba39620bf54f4a63acc56380edef49d5d4ee2f86dc941ad2657b22ef0896630fa45b41e21aff253c7533e499d8d85fca4afaa8c83aceb20967db258204bdf43a23ba71190cab10b90b9e39febeb864094e8ae4b9aa5edbcc75130e485262206b921c7de08b1b329997fd0bb3c72d13bb668181cf414065f6828e25fcb9db37a9c4259b68b3e084c16681923e3dc0cdbeffb6e2452c3ddabd2e4a";

    bytes public encodedInputs =
        hex"2d67ce857e88697ec8e07273147771c48dd613b3414188c49dcd0b7de6fc70f7089774995efd8b027fd54b2a2949114cd45a5185789e7a38643e83a71d35d6a10000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000000502ee9e8b71739857968ff827595e586cb32dfdeca8c7cd3a8096aa3803ab6a9e000000000000000000000000000000000000000000000000e1adc725ab38cb7a000000000000000000000000000000000000000000000000151801ca84e8474600000000000000000000000000000000000000000000000017772596fc928e3e00000000000000000000000000000000000000000000000054a7bf66a2957da1000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000030b550f3efaa76c4105356129177873a1651a95602e7b953d7a1a9f54393607d05b62ea1b0fd9302768edcd3ed22a43688000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030872cdf0112e5c8ed5571b68834ed5a6b6c87c7970f71f26696157fe9c561a1d9628d8b15206599ddbc3e59626101228800000000000000000000000000000000";

    uint256[2] public AR = [
        5255254381642087975337571456439890521991844234454453936336644091465676155588,
        19738298575740500933310651504490921151381719545626107314255749214553830660986
    ];
    uint256[2][2] public BS = [
        [   8146258450249515152918402313251778446612256628084971067118032998240771547551,
            3107373871071471584734273928292080043110860130360218549345895688337324095572
        ],
        [
            471445062833941172575078748119631501009779854687540787459477316757507838971,
            20820860246196191652194916214732833993758652820630185921882981702437477784125
        ]
    ];
    uint256[2] public KRS = [
        19489390374833505282353104192358198793426298298406771904451856948619066748369,
        3276939413577742537929372270471611694077360790999588494067577444261409482107
    ];

    uint256[2] public commitments = [
        15800972568708217084517782590253837421866525167695114398551634054700425897947,
        16965298244137469982739627623621726367523236829941645408152266132150878500115
    ];

    uint256[2] public commitmentPok = [
        6460161457519228688051785894560190107390904475431178048840498540631486268046,
        17182103591069499386396298016551026626815155146233528456720982956114397965898
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
        stv.verifyProof(zkp, encodedInputs);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[9] memory inputBad = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER_BAD, NUM_NEW_VOTES, NUM_OVERWRITES, BLOB_EVALUATION_POINT_Z, BLOB_EVALUATION_POINT_Y_L1, BLOB_EVALUATION_POINT_Y_L2, BLOB_EVALUATION_POINT_Y_L3, BLOB_EVALUATION_POINT_Y_L4];
        vm.expectRevert();
        stv.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
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
            [
                ROOT_HASH_BEFORE,
                ROOT_HASH_AFTER,
                NUM_NEW_VOTES,
                NUM_OVERWRITES,
                BLOB_EVALUATION_POINT_Z,
                BLOB_EVALUATION_POINT_Y_L1,
                BLOB_EVALUATION_POINT_Y_L2,
                BLOB_EVALUATION_POINT_Y_L3,
                BLOB_EVALUATION_POINT_Y_L4
            ],
            BLOB_COMMITMENT,
            BLOB_PROOF
        );
        assertEq(_encodedInputs, encodedInputs);
    }

    function test_Decode_Inputs() public view {
        (uint256[9] memory inputs, bytes memory blobCommitment, bytes memory blobProof) = decodeInputs(encodedInputs);
        if (
            inputs[0] != ROOT_HASH_BEFORE ||
            inputs[1] != ROOT_HASH_AFTER ||
            inputs[2] != NUM_NEW_VOTES ||
            inputs[3] != NUM_OVERWRITES ||
            inputs[4] != BLOB_EVALUATION_POINT_Z ||
            inputs[5] != BLOB_EVALUATION_POINT_Y_L1 ||
            inputs[6] != BLOB_EVALUATION_POINT_Y_L2 ||
            inputs[7] != BLOB_EVALUATION_POINT_Y_L3 ||
            inputs[8] != BLOB_EVALUATION_POINT_Y_L4
        ) {
            revert();
        }
        if (keccak256(blobCommitment) != keccak256(BLOB_COMMITMENT)) {
            revert();
        }
        if (keccak256(blobProof) != keccak256(BLOB_PROOF)) {
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

    function decodeInputs(bytes memory _encodedInputs) public pure returns (uint256[9] memory, bytes memory, bytes memory) {
        (uint256[9] memory zkinput, bytes memory blobCommitment, bytes memory blobProof) = abi.decode(_encodedInputs, (uint256[9], bytes, bytes));
        return (zkinput, blobCommitment, blobProof);
    }

    function encodeInputs(uint256[9] memory inputs, bytes memory blobCommitment, bytes memory blobProof) public pure returns (bytes memory) {
        return abi.encode(inputs, blobCommitment, blobProof);
    }       
}
