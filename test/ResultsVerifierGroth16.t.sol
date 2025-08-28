// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 1955711634442266218442623404753454420621019669225137434447114896701269920059;
    uint256[8] public results = [39,46,29,22,42,0,0,0];

    bytes public zkp =
        hex"126448c676044eb81e2cd15cb9c0383596aa3450990e57744c17eb475f2bdd502f8f66a9f7610a9b97743d892fc3f79b632c9a69d7dcc08d4a1239f68fca226c08206218068b9bc835966fbfe081fa44b61ab29aa00529597a9d57d6e3e866a30dea0313b01990038b2dcb83f07fadd555dbdb73cb3a20bde2568e3858bd8a3f232bedae35631c49aee01b02806893ac86bbf8b81542b79029a374076482ed2c07c4d94adcff5b44b99ccbe083f10c87915e2cf5e932d11ae20659cd71c84e1d0c8c853624b49262a60f1c1701015769ac0d7292c36aea95f436ffc3b5e8509a0cf80390c9cef9bf7af8dd5118abbce4dcfac7fd4ad9b8a877e30e6dc9a43a0d0c40b2eb637ebbec1ddc9e679ebb3b604d3fb31ed19e1206c340484e13af263603cec4f52a87f7f4b9d8ebe2b7f455e8b1c112f82f6b267babb8c9d5e852e13a224fef3c8e634797bc3571e232466f7c27d2750b56007e764652f930a1fc51d2267da98091f93f68eb1d5a8b29bcb195a0d5dacc60b0821ecd3986e6b0d2fa72";

    bytes public encodedInputs =
        hex"0452e4c03c0464b0d490346235926028e4f8708b358f059229596a7feceac93b0000000000000000000000000000000000000000000000000000000000000027000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000001d0000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        8318818257211664858438670582831178768568438731482799780709819644134527065424,
        21512071574087245172334609493826070378708432850968890899825369309239097041516
    ];
    uint256[2][2] public BS = [
        [
            3675718913609050257909209905418333981687392274354933349249013886400189720227,
            6293530480764111188000881368685141560364386019408370903169755593115285948991
        ],
        [
            15908564534737026720076203016775042575826068154752614735456788775755717995820,
            3513991662046486714068772023930046572264854346654427948137105234815061872157
        ]
    ];
    uint256[2] public KRS = [
        5676032164034939075069529561205093939853781727050139266126031970793897349274,
        5865956863788511855196796804764137399669926569643261237092905610872810125837
    ];

    uint256[2] public commitments = [
        5542067252060261006613010562389545868932121860457667008708566198454269716022,
        1722268393044101218178930367527544425256517857283872765964027265028719829306
    ];

    uint256[2] public commitmentPok = [
        15519868919917419240300876915352563981981693124669853680107258384363151970770,
        17409913990640018646260601168071341376016185576926969061352553997364175764082
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
