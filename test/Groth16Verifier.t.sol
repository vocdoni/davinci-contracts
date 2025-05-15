// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { Groth16Verifier } from "../src/Groth16Verifier.sol";

contract Groth16VerifierTest is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        18588877429853608149719955156994158953324877258620908236720089071601357088214;
    uint256 public constant ROOT_HASH_AFTER =
        975414945923616221792865859054025364019163639071210742046394818834082831994;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;

    bytes public zkp =
        hex"0b7040fa918a73fc425a389ee2d6621f582cb0e39ecde00a8f7bf3869f0bdbfa260c95c56603a628f9074e82a1032901a269ab095048a32523531460d03736df0a3f6aaf4e991a38a8431ab393c36bc29e32b87c10f0bb35a5911cbf111f270c01d0f11d0102640a05582d209fb955ea53b530c0abc5203099e28e96fdb784852a0de089e14c41292dbf7af6e34f8422ae275932d7ea01471fbbfacd2e69beec05f7faad672a539282192b4e6ae6b3c08cf865e3904edd564d936c43eadb141b2e79619b41b31753c444b9a01563fb5307abca632682e69dfb600550b78ccdcf042f2490f89161d61e8e8616e2bed435afbbe583babb62545008f90976a25cc2146770e09fad564a591b46b12de00ca7da3dfe498e0af0cd943b39842b33e1f701183f0c4ca7bc26c5eb1d29b453e73e19b7e9f64974c99d712864c5b78216f70bc863836978b5e6413acabcb4961d63f162c6f3f47e8a6b21bc419be7dc768a13388826cc95cf1ab63c3298d058046b2a46760a9a783702c6a9eeb7e287794d";

    bytes public encodedInputs =
        hex"022810b72b3e8c4fbe85150856ce804b0a47eab2a64ec90136277c84c5dc1e7a2918ee88f8ccdbb40c38c0c709de736b94c03909155117b5be90de45b736e5d600000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        5173776672751209281262297117153416716779721983222591036322217403123373956090,
        17210124092999983480338087416621658436639718058555779948175894056483822188255
    ];
    uint256[2][2] public BS = [
        [
            4635176162294470083118041433381193449974070364997689296062819620507311089420,
            821481140871496002821682157387128423573873423543600976022287002404889527429
        ],
        [
            19021658360760316715360489718634837589268671726670452596805030365721854066412,
            2699705579438642064499029658071162493981678923093768543361112913614883984411
        ]
    ];
    uint256[2] public KRS = [
        21020853184774783234872331383550460401533712963966522436446649848924874788303,
        1892545577655709238016579147815581741864794534652919151346865421417080577218
    ];

    uint256[2] public commitments = [
        9229021270772356449068025267993728632029556912900337044274877453418383008247,
        495152319749891753423247786203894207247292674294035972260693277826101679863
    ];

    uint256[2] public commitmentPok = [
        5329497563120400863608971716974581572918357904886160074830421877239771330186,
        8693827242236123118943307574123644454918715497848307163787962839732767390029
    ];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    Groth16Verifier public groth16Verifier;

    function setUp() public {
        groth16Verifier = new Groth16Verifier();
    }

    function test_Verify_OK() public view {
        uint256[4] memory input = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES];
        groth16Verifier.verifyProof(proof, commitments, commitmentPok, input);
    }

    function test_Verify_OK_ABIEncoded() public view {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory input = decodeInputs(encodedInputs);
        groth16Verifier.verifyProof(_proof, _commitments, _commitmentPok, input);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory inputBad = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER_BAD, NUM_NEW_VOTES, NUM_OVERWRITES];
        vm.expectRevert();
        groth16Verifier.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
    }

    function test_Encode_Proof() public {
        bytes memory encodedProof = encodeProof(proof, commitments, commitmentPok);
        if (keccak256(encodedProof) != keccak256(zkp)) {
            revert();
        }
    }

    function test_Decode_Proof() public {
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

    function test_Encode_Inputs() public {
        bytes memory _encodedInputs = encodeInputs([ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES]);
        if (keccak256(_encodedInputs) != keccak256(encodedInputs)) {
            revert();
        }
    }

    function test_Decode_Inputs() public {
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
