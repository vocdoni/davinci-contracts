// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 1685650003863957963309641629986376901572311260206043595114454660212433218452;
    // uint256 public constant ROOT_HASH_AFTER =
    //     17221650114163349820651399714058649394077581872287092757887500660867278649420;
    // uint256 public constant ROOT_HASH_AFTER_BAD =
    //     4033651216511111400175039953981700650715262298138585644694383935611028748527;
    // uint256 public constant NUM_NEW_VOTES = 5;
    // uint256 public constant NUM_OVERWRITES = 0;
    uint256[8] public results = [1000, 1010, 1020, 1030, 1040, 1050, 1060, 1070];

    bytes public zkp =
        hex"096583c1bd8d923a89497fcf7bf66891409633c18935dc56078c750091518401101683623d4fc6c9cc0834bd69672e03c5a111b8cafcbdd1433bec8f14dd85772c4750946030958bbbbec7a6c592002cb265af0d346b60370ade3bc8afc9caa407126d8ee3954db0f5bf4555c49b8102e97755ffd6a43689ddbf61e5f35aaa070d99f3d270c6d2acfedf8ad6a8362be8dc4d14cd6a038a2660fc4a2b3db9af3b10c31d171db478fab094fd81615fe456edbf3674c7e007dc6c12ee31d24910e11de1f4eaaa5ad3813f9f0f286e4cc550ce500a6fe7d047cf9fcfa9c66dc8d3921321c72c942bcb0f75e0d7cc017a1a66fe44e01907e1ac73b7f41ec2f8962a280000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    bytes public encodedInputs =
        hex"03ba0b499f549a4bb53358592b3578cccb45156798cde3fca12b562e9afdeb9400000000000000000000000000000000000000000000000000000000000003e800000000000000000000000000000000000000000000000000000000000003f200000000000000000000000000000000000000000000000000000000000003fc00000000000000000000000000000000000000000000000000000000000004060000000000000000000000000000000000000000000000000000000000000410000000000000000000000000000000000000000000000000000000000000041a0000000000000000000000000000000000000000000000000000000000000424000000000000000000000000000000000000000000000000000000000000042e";

    // uint256[2] public AR = [
    //     6690523878390319918509611247270935344081208698466174844978970979792442643837,
    //     2166962242191292610243539137597601380137671858915626055585063694555543153969
    // ];
    // uint256[2][2] public BS = [
    //     [
    //         9874927553947877295979652402631591550437856872022349896963978568639907058658,
    //         4303886892772926097884654255499039817666521483630969759858018479378304745829
    //     ],
    //     [
    //         21184132092252492149957299047417895209020312208239936825930839844634945497644,
    //         21525220636806930291441574321110514102010571040506320292371059359334974174765
    //     ]
    // ];
    // uint256[2] public KRS = [
    //     2620769365939176202866016348031386409814515811278094596815241863871051366875,
    //     12683800096244918802769427560548958824349330543563720026993670470796809526709
    // ];

    // uint256[2] public commitments = [
    //     16197980224318809777045985116144560269586463851287371737497825306201848473845,
    //     3381673314454499703629915435640441865696436233911336509982687421882380106822
    // ];

    // uint256[2] public commitmentPok = [
    //     8430366415231100058227725042285520738574705791554959158016636990365512692265,
    //     17483511960721826887711148106968922804607479998602155924724037380035129052571
    // ];

    // uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    ResultsVerifierGroth16 public rv;

    function setUp() public {
        rv = new ResultsVerifierGroth16();
    }

    // function test_Verify_OK() public view {
    //     uint256[4] memory input = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES];
    //     rv.verifyProof(proof, commitments, commitmentPok, input);
    // }

    function test_Verify_OK_ABIEncoded() public view {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory input = decodeInputs(encodedInputs);
        rv.verifyProof(_proof, _commitments, _commitmentPok, input);
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
    //     uint256[4] memory inputs = decodeInputs(encodedInputs);
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

    function decodeInputs(bytes memory _encodedInputs) public pure returns (uint256[4] memory) {
        uint256[4] memory inputs = abi.decode(_encodedInputs, (uint256[4]));
        return inputs;
    }

    function encodeInputs(uint256[4] memory inputs) public pure returns (bytes memory) {
        return abi.encode(inputs);
    }
}
