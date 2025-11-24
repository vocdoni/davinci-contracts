// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 3002110162893369572337896915733180539392271285963756713877989564646742292824;
    uint256[8] public results = [60, 27, 32, 36, 53, 0, 0, 0];

    bytes public zkp =
        hex"05ca2c09dcc334557003b061bf248e42aa63dc99aebb056cd650b6dc3e3598ac06c1fde80f3f9cbec4b4ef7c40515a5387d1a6df8ea9221a8377c7ad29cc88621af8aab5bf16c1867f76e34688ada2602a108fb5f1f106b90be876898542423c2e4ac3cd8cf8901b8795c16312774f5ff6d58b6989cb846fb428229663e75ad92cfc02f3e14c1b3469ff2dcc94a1cf0381f45be8d8e6969725e0de3c1928164316db93d51bba4466090e35fa4c022c74f9bd9cd4a4dcdcf82fd50d1d8dd855c516a59c77fd86ec48faa40e8f2f55c1e05257bbe5798c602b88083a54da2fe2ca0a8f41b7a2feb905ae12ad6360efda192c34c2229cd936807ab834462d179bf916bd3e70a70266e2e1490c887159f37d1b80fc615b8f415a9f5e1efaf55af2740e696ac8f42daf5ec1567d2e5de0614b2ca54d20fbd6d79178d55779211f77ac1a54944159e00b44c1fc34782a372f54a8022451432bb036bc22f76b536a7e550d2c349a6eb6e4883a69ca261ae13cddc32e3b346fc63d7707ca06b74f470113";

    bytes public encodedInputs =
        hex"06a32256cffcebf6d1b5cc77122368f5a1ff31550fe5897ca53a75968c8ff958000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000000000000000000000000000000000000000001b000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000240000000000000000000000000000000000000000000000000000000000000035000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        2618771292729350976935376365241543206163775670059861835509223906170631461036,
        3056630973141044249049811849449943075314164747805361582386017038325655570530
    ];
    uint256[2][2] public BS = [
        [
            12199490331983275528612034635349955044690610601100929957819259479330939617852,
            20938489099796518919418358217958627304820457840925455418472355330396063750873
        ],
        [
            20347031176474181566564287938660158864546031334440598963483759287788395566659,
            10338842478120002777594810079563454132739766901083862359922787284577437636037
        ]
    ];
    uint256[2] public KRS = [
        10243492341883560567762045395671772723454698417790589919851535859531413185226,
        4776241180444132426965032762057554694110354844528683355652929863280426982393
    ];

    uint256[2] public commitments = [
        10285247709450643711088050752531800617020510612648855539080014836526239904372,
        6518635824784545924065967299220927812983917455223544913833921467265430026156
    ];

    uint256[2] public commitmentPok = [
        11909572436927129071253036827987077542469491031338624344826989023214313307733,
        5958171356734121587761975129364390494077076695303098837760534157744706879763
    ];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];
    uint256[9] public inputs = [
        STATE_ROOT,
        results[0],
        results[1],
        results[2],
        results[3],
        results[4],
        results[5],
        results[6],
        results[7]
    ];

    ResultsVerifierGroth16 public rv;

    function setUp() public {
        rv = new ResultsVerifierGroth16();
    }

    function test_Verify_OK() public view {
        rv.verifyProof(proof, commitments, commitmentPok, inputs);
    }

    function test_Verify_OK_ABIEncoded() public view {
        rv.verifyProof(zkp, encodedInputs);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[9] memory inputBad = [
            STATE_ROOT,
            results[0],
            results[1],
            results[2],
            results[3],
            results[4],
            results[5],
            results[6],
            999 // Bad input
        ];
        vm.expectRevert();
        rv.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
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
        bytes memory _encodedInputs = encodeInputs(
            [STATE_ROOT, results[0], results[1], results[2], results[3], results[4], results[5], results[6], results[7]]
        );
        if (keccak256(_encodedInputs) != keccak256(encodedInputs)) {
            revert();
        }
    }

    function test_Decode_Inputs() public view {
        uint256[9] memory _inputs = decodeInputs(encodedInputs);
        if (
            _inputs[0] != STATE_ROOT ||
            _inputs[1] != results[0] ||
            _inputs[2] != results[1] ||
            _inputs[3] != results[2] ||
            _inputs[4] != results[3] ||
            _inputs[5] != results[4] ||
            _inputs[6] != results[5] ||
            _inputs[7] != results[6] ||
            _inputs[8] != results[7]
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

    function decodeInputs(bytes memory _encodedInputs) public pure returns (uint256[9] memory) {
        uint256[9] memory _inputs = abi.decode(_encodedInputs, (uint256[9]));
        return _inputs;
    }

    function encodeInputs(uint256[9] memory _inputs) public pure returns (bytes memory) {
        return abi.encode(_inputs);
    }
}
