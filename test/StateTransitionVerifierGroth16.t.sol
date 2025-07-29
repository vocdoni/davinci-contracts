// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        13531415782221581052397236423765203891258985052064429347355472164469722820300;
    uint256 public constant ROOT_HASH_AFTER =
        7961306345560008268868495476665083487633676217546477646432696512080076649619;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 5;

    bytes public zkp =
        hex"2584c45f793130cd312f863c09e5655ea174b3a97eec316a8f50a4d289a95ca60edba9bc0ae1f60b7c7708e1337a18f05d341a73ea30250c6de0739a6499d47106473ea902aa32fcbc81d18dd1a8c0ff28f038e3f8de86ae665c0c1c9ccdc4f917f4c2398784d81c3915b50a4721de76d1656a691e65d8df6ee25caea68c75b80c1759410f6e42af4af6dbc959dde83d86e919acdd73133bbcb2430ff5c91b3b22a3771d43ad380e1425f551c0c1cdc7ede36f6f439e083a89d424d66d1b4b352173196bd78a235d42bbc969c1a34f489eca8a38b22b05d7ee1071ef828bcfff13ed917d389d49ba80feb5d9168f8e4c61c465e137b72dd9a705063a3e16974f0ed555a514e10e2801d51887fea533670bd93c1e26d5cd7de68778b7fbfb91a701be42514d05a818671ef6ae3f9a15b8deda48a1bbcc5d109996a7e9d740b9c3141389b2eb8422443330bd46c515ab545b26f11302f16aabe78bddbfbbbbb0aa0864a196d64ad676e23b05451e56eb1f36aba8ea39b92ae4b626593d5d0971e8";

    bytes public encodedInputs =
        hex"1dea828a7820db09993535c3cb8b6202b9f7bed6dab77754729fbd8b14efaecc1199f090a5032e8d2ea0cc67145c1652c37c96b1caf28d9f2051253a383cb49300000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000005";

    uint256[2] public AR = [
        16970154526373530574891719885706603583340627935741036845150124562481311997094,
        6720490852100854561899407185126765900129996510853792137741120141385482753137
    ];
    uint256[2][2] public BS = [
        [
            2839755697883991841430601970102734707010305963535167924018286064087782507769,
            10835646691001098407388954518377398718300036021557935320743229114182087243192
        ],
        [
            5469007674935540133625508935311106373589271406987041291208057300924525779771,
            15667455020170836965675768443211289303536477745147126924441770964251116784437
        ]
    ];
    uint256[2] public KRS = [
        15129686866769237607640609291636108157918751165457307937887463633641649459199,
        9013691006610349355278713676536380355296344835612122020214865772218196334415
    ];

    uint256[2] public commitments = [
        6709309403993030790187359753551490145531086176173058986642287680600959324583,
        788471498017101085351895164863093667432203664246658128843378837897684761027
    ];

    uint256[2] public commitmentPok = [
        9080777428818879386275474018659839180023175477405188947897012876033459466410,
        3796302742865419775678982616840413498330999222162777889423528932917811704296
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
        assertEq(_encodedInputs, encodedInputs);
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
