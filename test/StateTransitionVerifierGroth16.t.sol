// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        14424036259012750574058040874268427628643176097624202352033583206326205773957;
    uint256 public constant ROOT_HASH_AFTER =
        5244826649985884941716637507837432287554322849727016690652948098167443117049;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 5;

    bytes public zkp =
        hex"2fae5436ab7f722800bf8f04d0141f8bd8bb45fe6f697fb17de992fc213c38990ab27a48bf7d7707b87043ecc0e3c49f1cbe05437145285f7dec429d5a923506053e36ba3d95b0d8854d1cc8c61d2e26484e8c127e84c07edb8e7fc5ded4e73e05df7d9f586322f7be883b3cd98efd5ebe2d976269f67a50ccc83629983c61b70796813e0681f073711a783b65ea906ce5a091f99756f1dbfc36448e4b957a4f1fef874f0919eef257a3e20065ccbeb54f93353b4e4697b59f836fcd10329f9218a4b5117979d61dec2526b7a2c65e6f24d687140c42224dffbb3af67230d12b28c200ca8db1fbf553ca5100aa717e21be7fd01018797c98222b0193ae342c532d3064567ee8d80219a2d9f8f9f5ae6d9c0ab2fe0b3bf2fdecca83059d9096c9286122c449fbf3746250995ffaa4a76bb027db97ae89cab1cd92b933c3e8470a0b1538acb5269c96454f9f5ce37661596f332f1838ab051bbc309124a81b136a2e8029dba69e96d8d9e6546f1160425a62d2b280bb7dd45912a17863f8d9680c";

    bytes public encodedInputs =
        hex"1fe3b71818c0557f2f40ec198c8cf176b54d37e9b6af21f71ae2fc13ca2578850b987778b1b51747b94a15831946b0eb2431ad7c2a3608c733419986189553f900000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000005";

    uint256[2] public AR = [
        21566716493276039647444112558325306066520096128004776965305975737876622358681,
        4838471237699945778227309883730549429794554219183860603782926813282830529798
    ];
    uint256[2][2] public BS = [
        [
            2371486476271027567524812685561301807115132198521865285060204917610436159294,
            2656438152595044640425053385076959456425871741913321787588850867424825467319
        ],
        [
            3432108997280379375874571068765509280117088210076159316224546436975210887759,
            14444908621114374326317661009043804764863117078183468387156759899277280583570
        ]
    ];
    uint256[2] public KRS = [
        11146520971822818661939912202312492479925371817053136697732095499503504773419,
        18435287734729140442622957654700750181306263094719920919599368287952953420883
    ];

    uint256[2] public commitments = [
        20439579351911577482316546479515148181301916227456051804859693800773603333833,
        18264138059930940374872528219350738108428714864098287036500863080933255300874
    ];

    uint256[2] public commitmentPok = [
        5012936276759932075413024397439620700254171511210897790686824966193837445994,
        21032836352497505827363153360237653259634076194158047454227906364102999107596
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
