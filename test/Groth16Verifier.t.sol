// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { Groth16Verifier } from "../src/Groth16Verifier.sol";

contract Groth16VerifierTest is Test {
    uint256 public constant ROOT_HASH_AFTER =
        15037863611910112978170348844338176055167222321310792864682330036565985067068;
    uint256 public constant ROOT_HASH_BEFORE =
        17279952791970455231980792301299095313810117315352690795312020434419193683212;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;

    bytes public zkp =
        hex"28941cc89af4c1cc7eef0d8aa0898d2118d44f251f28b0dd2a863ef9981d8a97267eb23ca9579278bf978e79f83aa52d916a897b212964da1316a5e22de80b81010ce045c105d6ee15ea5e74725d146eac9c80493889a9fac3cd021d550e002314b24ebc81c21e465e471332b125c6a3e4ca102d426e5d5c53c26a9a610e0b8e18ab6f1f409ec3a6369f52a86e1f6e583cb62795838288b84df03733a688d84d0dcd1373011eec61fbb82575e38b616a028b8c4b01b3757cf7f2687f34e514630792bdb8bba1b25a4dc34924901e74eb836c01425c2ecdc3ce1be301353f1dbd0946d21d8988f345b8ac25de1023ceb94956dee673c4c82115cad30723166a071eda569cc330346a5c0d2ca649fb7abae8990e9674d10aea7bfbd8da5b51683f11b314e3f73818fd4f38671b566d525931a98642649919498b88086acdc34aa9233d269e53d34c2df0ede8f1a6f11a10df0b6841da0e1b0e34c16cc589a937c0105ba383d71aba7b0800fc7f06c1fb482c337fbe309fdf8afb9061567a1a1f52";

    bytes public encodedInputs =
        hex"213f211206637c602aa9d7251fbb300386cedf03477121489740a42eaad6003c26341b4fcbafc76387ca184fcbfd1fe0b5b4db9e4ecff2707f6f779707fb1d0c00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        18354205966123680099263405597349980669361136475226504693276884704468920011415,
        17411741122606530320118758087551481722378545584847840797835544054129821748097
    ];
    uint256[2][2] public BS = [
        [
            475062885106282671968094522514747686807922528886182873448622012010710564899,
            9361299167546044512552097216399206435293192471244251357447339186847023500174
        ],
        [
            11158406150483635968627465022143141995025381712452145153371371043393771198541,
            6242404913554521007731247209895608616428823714325364263123708678632355730531
        ]
    ];
    uint256[2] public KRS = [
        3425459021990155705276123846536755710319779137546322924012376039437421714877,
        4195945094839285801328849237466672178419428762594124460031008170621405653511
    ];

    uint256[2] public commitments = [
        13955155894112910657499097609176065145422475864082293665588620147325670615103,
        8005728231380901914311303002845245391374691757383602436176988436144003959465
    ];

    uint256[2] public commitmentPok = [
        15938993906226399477953388844429762057956899155382747564773851119740019881920,
        7398917199287773597714814105548244473211793445967546431663791470462319075154
    ];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    Groth16Verifier public groth16Verifier;

    function setUp() public {
        groth16Verifier = new Groth16Verifier();
    }

    function test_Verify_OK() public view {
        uint256[4] memory input = [ROOT_HASH_AFTER, ROOT_HASH_BEFORE, NUM_NEW_VOTES, NUM_OVERWRITES];
        groth16Verifier.verifyProof(proof, commitments, commitmentPok, input);
    }

    function test_Verify_OK_ABIEncoded() public view {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory input = decodeInputs(encodedInputs);
        groth16Verifier.verifyProof(_proof, _commitments, _commitmentPok, input);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory inputBad = [ROOT_HASH_AFTER_BAD, ROOT_HASH_BEFORE, NUM_NEW_VOTES, NUM_OVERWRITES];
        vm.expectRevert();
        groth16Verifier.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
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

    function test_Encode_Inputs() public pure {
        bytes memory _encodedInputs = encodeInputs([ROOT_HASH_AFTER, ROOT_HASH_BEFORE, NUM_NEW_VOTES, NUM_OVERWRITES]);
        if (keccak256(_encodedInputs) != keccak256(_encodedInputs)) {
            revert();
        }
    }

    function test_Decode_Inputs() public view {
        uint256[4] memory inputs = decodeInputs(encodedInputs);
        if (
            inputs[0] != ROOT_HASH_AFTER ||
            inputs[1] != ROOT_HASH_BEFORE ||
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
