// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        1751626999407588979472546341264113545911838590760664666761189598937796389237;
    uint256 public constant ROOT_HASH_AFTER =
        17835852697581905354318759729521767940424740307802761986573785697616889279907;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;

    bytes public zkp =
        hex"177c9222ef15afca7ea204431a34fac072c034af245fb766402717af9fec8ebd0d53a546d47ac08ee2966d752a3be857e8e75c40374ced6fdeb697e82cc8fb9e207735c103a374cb3a361fc739c69019994da4659d714348854b4b69d2d78751291d6764ee1bded4484ac9768729fe5cc67608acd5e77cd085a6f929345bd8fd17f9e9a5349f7dc13742646d64063fe2605a556612b94a787e43de3e3010cc292be1f4d5fb505826a7669c56b8d3b38174a14ff87d5cbcba68379db53f2772c00abdf4756c4759353381bfa7ef7f63770f12eb2bafc4a9640c2086b6c2a42e0d24e60ea1acd697bc249445159aeeec75340f6d33d142f92436260fcc134cfc2613f88a50d67f1f704db53f19200a06c0351fd565ee3d13566a2df8dd7f0d62500225b17b01b872291b84b34a2638ccbbc4a8f34a1f2323218a5266750a34915d11047979da1df7c706f6a0eaaee20e387882828a4682cc5229cc6d0fa396683c05202fd6899b19e2c49c0f73522a05f11c1997e818ffa8f35a02a7f5a8d61962";

    bytes public encodedInputs =
        hex"03df62c0662d3d85fa8dc8073e9d8a2b59c8388cfd75862efe3b2e4255a92975276ebc21494819763f77dd398f166812e2243905ffb8dd705ad0ca9023a275a300000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        10623293150231023787837932013817088497845976979958358049193679687167662395069,
        6027856035679259243519616626173512873878669783494886303663574718578334170014
    ];
    uint256[2][2] public BS = [
        [
            14684636951582417689329636898168362762403127740460244691344081577937305306961,
            18596778957736671857640811445792172141627209400899166183964572258820935112957
        ],
        [
            10844752997376796964880349326498605040667845772712491430138777659036323073065,
            19848682873699285127996286125748833592979649443095542213093534235328919335616
        ]
    ];
    uint256[2] public KRS = [
        4858749772901234033745377295932363743194655342872725570701554501980392795661,
        16689738357098879345835108078346621645103942735589082540983300879199236848678
    ];

    uint256[2] public commitments = [
        9033076815527816869470013500639345737757328860825645361475623869433129034320,
        971223963921343936538598606207276822633511391736765741181224849232083849565
    ];

    uint256[2] public commitmentPok = [
        7697224210606558177907503326864390716112866449997781470200446960926402701372,
        2318433514987715024233695425523400671605089109399818000999418356624480672098
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
