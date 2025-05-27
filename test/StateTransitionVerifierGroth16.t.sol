// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        15518021866372582043604022785954659526716738998075185570627351931444151175689;
    uint256 public constant ROOT_HASH_AFTER =
        17221650114163349820651399714058649394077581872287092757887500660867278649420;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;

    bytes public zkp =
        hex"0ecab3cb35e0fe048245884fd3714a38d5f252a635eee1daaefcecedbe28557d04ca7508c53b20f185cf07cee7769eb259d6b95862cd6b3bb283828e1f2aad3115d502cc34fb6f67960d0ba635f26ca472e92507b34269ab274d5b32942b53e20983e9e55869b7a6894d5741a8eb92616910cb07faffc943309c1a75b60e41652ed5cb3a891b992a9e65b617f92d90fe27df72b22026f64973537d84b759aa2c2f96d7d7cf3a22184b05c7a697ef2539cff420956944da825a50f2017fa9362d05cb4d8a84499313b004bfc497bd180bf614c1e35e96fcae757df780c7e081db1c0ac6c54f917aee5df7532b681d39d2d584f6fd4c306f04870cb070db3eddb523cfbb5ec47f2ba67e73558f057f5fa22da5e1bdea9858a0553c159d3edee8f50779f592936e47a69b718f08fdaed9a34b2660fe110e1c38607cf3fc3bd0bc4612a36b159967b420a0c0860efd43eae410620b58e2bf639c795100e899baf22926a7512d00d5f9668fb943cb61ee6ee12a036d5d565e6c57ed252a2c1edf0d9b";

    bytes public encodedInputs =
        hex"224ee39d8852a0a5b414cb2a312f9cc812b767a6d676c583b0fe17b82517ca0926131bc948cf601e6038c00507492262a56a511c85c0cb0a19ad411639b1384c00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        6690523878390319918509611247270935344081208698466174844978970979792442643837,
        2166962242191292610243539137597601380137671858915626055585063694555543153969
    ];
    uint256[2][2] public BS = [
        [
            9874927553947877295979652402631591550437856872022349896963978568639907058658,
            4303886892772926097884654255499039817666521483630969759858018479378304745829
        ],
        [
            21184132092252492149957299047417895209020312208239936825930839844634945497644,
            21525220636806930291441574321110514102010571040506320292371059359334974174765
        ]
    ];
    uint256[2] public KRS = [
        2620769365939176202866016348031386409814515811278094596815241863871051366875,
        12683800096244918802769427560548958824349330543563720026993670470796809526709
    ];

    uint256[2] public commitments = [
        16197980224318809777045985116144560269586463851287371737497825306201848473845,
        3381673314454499703629915435640441865696436233911336509982687421882380106822
    ];

    uint256[2] public commitmentPok = [
        8430366415231100058227725042285520738574705791554959158016636990365512692265,
        17483511960721826887711148106968922804607479998602155924724037380035129052571
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
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory input = decodeInputs(encodedInputs);
        stv.verifyProof(_proof, _commitments, _commitmentPok, input);
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
