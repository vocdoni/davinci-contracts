// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        16282774127260957557439015348635445656767673488688041267615841575933035731354;
    uint256 public constant ROOT_HASH_AFTER =
        13999159323556783713340982889683090407637562075598756749526547570965285300122;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;

    bytes public zkp =
        hex"2e688a61c5924ca3464fce45d4b1b1639b31e3c7fc6e4e74244fd6af74ea57522e9aea56c395612a89f15176942a5f5a264d4125e00c024f725e00fe9d52765a226cd50609972adb62b8c71a5b1aa60b6dc298075b44203d1cecdf3fe43c7a171e586cc471d3934312167ded54fab40cdf81e61abd3608c830ae730759e2c87f0728ad0f1277c067e26ea451d337abc81abfb83cf1c972f2d01f14a87ffbe27a1f5253ae069b804783f59ef87db09ef7df1a45b93adc28f0dcc3d8cddcd90ab20ef68f7c1c26efedcfc7f897abc69a34a0dde812ac895baf2c69b1af02493ecb2e5e1f8d4ce0f18bc3e863970eaef46154e94c8f4f2b112a7a636c5283a2fe790741f6373905e5c54bfa57f0b70104f485f24b694baa249491f0a308734548bb02c9b350371a51c8db2b1abe0e01da678da7e826474481ffef66b1dcebad58a82d56fcfbcdaeca5a963a3b510a09bd6cd17afce730f213b729bba8d1da84d982209d534b8a7b2e19f402935ffc64a3897273f82379a8a670a42bc0b00ae157cc";

    bytes public encodedInputs =
        hex"23ffb93b6e4fb1b2098ceff1e0f20ec383761619d80f1d80ea21c652c738d99a1ef33e4eac84b3192df54859cf12c0a632f66912dbd1eee0ebc9d1dcb8bde79a00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        20991098206484595174095275617464791995657970944631916178407615875871687726930,
        21080102830604035478949797516223037152081833231184740081921684617959081932378
    ];
    uint256[2][2] public BS = [
        [
            15570926569568666395838260283250864520198628224121511621164666713260565494295,
            13725618683940815579464988764257096923950100675882000278250879709839091746943
        ],
        [
            3238058231136079061403837127294804495067796854805944870942721385332372923002,
            14167157302066432444191625404897139034131309392246532885940746937603911649970
        ]
    ];
    uint256[2] public KRS = [
        6768014553826951770579587661025634756542616873721497214783927943999617842891,
        20972692422504923139281148756018980671518573245472786286187530887864529387129
    ];

    uint256[2] public commitments = [
        3282734317697063605081642239025504272840109462145212864710103764402442684603,
        1260999532381800154131183638378553939895605722757747963982473359731838507176
    ];

    uint256[2] public commitmentPok = [
        20507773062504857616434131837637700095790883183305385620023541503658543077762,
        14751981025361284018242499959790194274043221045359177562958814531184631044044
    ];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    StateTransitionVerifierGroth16 public stv;

    function setUp() public {
        stv = new StateTransitionVerifierGroth16();
    }

    function test_Verify_OK() public view {
        uint256[4] memory input = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES];
        stv.verifyProof(proof, commitments, commitmentPok, input);
    }

    function test_Verify_OK_ABIEncoded() public view {
        stv.verifyProof(zkp, encodedInputs);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory inputBad = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER_BAD, NUM_NEW_VOTES, NUM_OVERWRITES];
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
        bytes memory _encodedInputs = encodeInputs([ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES]);
        if (keccak256(_encodedInputs) != keccak256(encodedInputs)) {
            revert();
        }
    }

    function test_Decode_Inputs() public view {
        uint256[4] memory inputs = decodeInputs(encodedInputs);
        if (
            inputs[0] != ROOT_HASH_BEFORE ||
            inputs[1] != ROOT_HASH_AFTER ||
            inputs[2] != NUM_NEW_VOTES ||
            inputs[3] != NUM_OVERWRITES
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

    function decodeInputs(bytes memory _encodedInputs) public pure returns (uint256[4] memory) {
        uint256[4] memory inputs = abi.decode(_encodedInputs, (uint256[4]));
        return inputs;
    }

    function encodeInputs(uint256[4] memory inputs) public pure returns (bytes memory) {
        return abi.encode(inputs);
    }
}
