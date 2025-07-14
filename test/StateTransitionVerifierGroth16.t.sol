// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        5540002152505115677895513852514986637835803378092572379307423469474254349764;
    uint256 public constant ROOT_HASH_AFTER =
        11187963620570868951372678928135497405170546783119497388495578896002282867940;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;

    bytes public zkp =
        hex"11f664630c1f811156665400ce2e163469b445d8b831c64b4529368e7904b6b60d68eea422950d76e139f3f74316f0152613f0da19d9787271b8534210f029482443ea282e0702139c7ca31dbd3500592254caaa48c83c941e2d3d9841c6823f13b7cfdb060b920b72963cc04c88aa0cdf1022f450a56bdefd365f8671bc278212eebecb93cf16ab145bc39bda72f7cc7f46bbd37afb1df7ec8f8827a7541dc619eac5e80acc8392522f7b3e8f04fe49e3fc3511d5a569648cfa1454efaae2fc1c30d76825ad65208eea4f5bd324a3b5d5fca2fed46395a4810aa0ba7edb22ad0bccca00295b7ef4ce64610fe7102161c5070553318699442df5a80ba04f3b210eeb32f50f3f9d4642b3ff903c797d61129491294e940211145e1df69cb059be26ef32e5d966f153d4643ed8c655e9a18b2446ddacb1fe606feb47c521228bf92488cc65e8b36098fc26e05501b4c4b5f4523f1cd8c267dcfe4cf1c16aa5dbb30c0a3f6fdda0a5d38edb7765deca42beac677d6aab46a1da5baf85dc6f29636c";

    bytes public encodedInputs =
        hex"0c3f87b496c4768f3dec047b277b18e7a5eed8c66e211e45a5a11f6831f2e5c418bc29bac03a0d1a8daa0554b51e8ff61bde2eb32e660d4e258419b20fd2cce400000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        8124655648797117815894964507170268036634444540169842227913763945637040404150,
        6065466167023139183922486976086098303896886725607080137638361105401688107336
    ];
    uint256[2][2] public BS = [
        [
            16403257394228005796180613147401395434629991757901488495595845489687102128703,
            8918711702295245036457235593371850780255335690708194143499971632806246492034
        ],
        [
            8563457696157212262499987743209393073137748663609851837387447908991851830726,
            11722629327615001238934939265351955974321994191012405793538836584743098311420
        ]
    ];
    uint256[2] public KRS = [
        12751055102707720046108667950522281871768558690076189331932568222820361511597,
        5337272292748215567183642147200984055447181396900716418436481510648052661025
    ];

    uint256[2] public commitments = [
        6747940634498776536886527223368825980969490713229830521593849283749398731198,
        17610515978686398375207148113236454639347481816516225852466380274829738806265
    ];

    uint256[2] public commitmentPok = [
        16524964453523052129890103734925638587319081155274372441689107827397564423091,
        5445860479561011476807954470491214330986399268340387557200642329662531920748
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
