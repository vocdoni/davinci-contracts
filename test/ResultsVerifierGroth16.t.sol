// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 848094410278823325666743560735506188625075380389348078732430222085993157699;
    // uint256 public constant ROOT_HASH_AFTER =
    //     17221650114163349820651399714058649394077581872287092757887500660867278649420;
    // uint256 public constant ROOT_HASH_AFTER_BAD =
    //     4033651216511111400175039953981700650715262298138585644694383935611028748527;
    // uint256 public constant NUM_NEW_VOTES = 5;
    // uint256 public constant NUM_OVERWRITES = 0;
    uint256[8] public results = [34, 43, 53, 54, 39, 0, 0, 0];

    bytes public zkp =
        hex"05a86dbf85917bfc66f6f14bcf657727336a160ac5d9b93d9459cbfd77626f5406c8735848fce63aea7d4ecab604910b0fd5e933009c57b0e6cbf62186a3b59e076f542c6bd7ed0135fa00712fa42f5cfbe732ed234a0a018778b21f8d07dcac12d05fd657a345232aa76bef4a59045a8e150805d7624652bb43f6e946da5e5009da5799af8c624ca93e968bbac7cbe9f88a3d583da37a718c4f1b4a977c05220e9ed21954dad6d23ff43d4e87f52f07c7c61475ba57d71ca5b2f1792ba255342d4393ce8820b73cb0a940402e4f62496500e57c437815341a184c633f9c7144026fe47f6729cf06e693eafbaf52da300529fead83d29ad58bb1e589355980db07bdf65fdbf964724e701735ec446cea024aef67a9ab2b8f4c7b52c7a3bbdca81bfe1c66b0601d688010f87de208e015a97fca68dcf1d6fa6728d51f9a37310d1db54e116a9b33f33b730effc5871b06a4482430f4ca406eec53806ad16a0d990709be94fafc539d09f11fd394c7471507a236baa7ef178cace12680014a3c0f";

    bytes public encodedInputs =
        hex"01e00122079a4f490caa16c23494b9323afe1347d9f507796146f7fa0cd654430000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000002b000000000000000000000000000000000000000000000000000000000000003500000000000000000000000000000000000000000000000000000000000000360000000000000000000000000000000000000000000000000000000000000027000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        2559152003567097676316411795198882359963344899624775768693209184796481318740,
        3068042585446984573067560437949981829641236615196609808894052094598619313566
    ];
    uint256[2][2] public BS = [
        [
            8509796908533565590414684930771419624731157470256424436925734786475904097872,
            3362890908561291071592037173181884097239615433507365346910843129676951968940
        ],
        [
            6612991766068464561495849680004841547408853560332343214761387044477784446260,
            4456592892662497794948153547332934625044266539597669174464746999740377072930
        ]
    ];
    uint256[2] public KRS = [
        20473477064385061026792487775572474710530571620298806368013501761047989350724,
        1102322754301593557070639315935211549293026673236343302003455261391351021787
    ];

    uint256[2] public commitments = [
        3501824449288209160331585076274396823285601485286021250891200154160374471848,
        12661422083588674991149031249329810647859543559683099283663662364290117480717
    ];

    uint256[2] public commitmentPok = [
        13437410733400705531114207559136546151185643617708927034834716963142963432857,
        3183406911975741188850774099879276101141441570347825191722910713085992516623
    ];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    ResultsVerifierGroth16 public rv;

    function setUp() public {
        rv = new ResultsVerifierGroth16();
    }

    // function test_Verify_OK() public view {
    //     uint256[4] memory input = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES];
    //     rv.verifyProof(proof, commitments, commitmentPok, input);
    // }

    function test_Verify_OK_ABIEncoded() public view {
        // (uint256[8] memory _proof, , ) = decodeProof(zkp);
        // uint256[4] memory input = decodeInputs(encodedInputs);
        rv.verifyProof(zkp, encodedInputs);
    }

    // function test_Verify_Fail() public {
    //     (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
    //     uint256[4] memory inputBad = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER_BAD, NUM_NEW_VOTES, NUM_OVERWRITES];
    //     vm.expectRevert();
    //     rv.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
    // }

    // function test_Encode_Proof() public view {
    //     bytes memory encodedProof = encodeProof(proof, commitments, commitmentPok);
    //     if (keccak256(encodedProof) != keccak256(zkp)) {
    //         revert();
    //     }
    // }

    // function test_Decode_Proof() public view {
    //     (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
    //     if (
    //         _proof[0] != proof[0] ||
    //         _proof[1] != proof[1] ||
    //         _proof[2] != proof[2] ||
    //         _proof[3] != proof[3] ||
    //         _proof[4] != proof[4] ||
    //         _proof[5] != proof[5] ||
    //         _proof[6] != proof[6] ||
    //         _proof[7] != proof[7] ||
    //         _commitments[0] != commitments[0] ||
    //         _commitments[1] != commitments[1] ||
    //         _commitmentPok[0] != commitmentPok[0] ||
    //         _commitmentPok[1] != commitmentPok[1]
    //     ) {
    //         revert();
    //     }
    // }

    // function test_Encode_Inputs() public view {
    //     bytes memory _encodedInputs = encodeInputs([ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES]);
    //     if (keccak256(_encodedInputs) != keccak256(encodedInputs)) {
    //         revert();
    //     }
    // }

    // function test_Decode_Inputs() public view {
    //     uint256[9] memory inputs = decodeInputs(encodedInputs);
    //     if (
    //         inputs[0] != ROOT_HASH_BEFORE ||
    //         inputs[1] != ROOT_HASH_AFTER ||
    //         inputs[2] != NUM_NEW_VOTES ||
    //         inputs[3] != NUM_OVERWRITES
    //     ) {
    //         revert();
    //     }
    // }

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
        uint256[9] memory inputs = abi.decode(_encodedInputs, (uint256[9]));
        return inputs;
    }

    function encodeInputs(uint256[9] memory inputs) public pure returns (bytes memory) {
        return abi.encode(inputs);
    }
}
