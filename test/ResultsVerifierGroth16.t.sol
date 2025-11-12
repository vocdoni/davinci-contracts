// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 19723816014752875800279306456496578536046296876388232350336782237768121417505;
    uint256[8] public results = [31, 28, 42, 41, 48, 0, 0, 0];

    bytes public zkp =
        hex"30171a422b25e70e0707be4fccffa7638a4c4cefc7e8c2e5625182aed3bda8d62e522daaafba69d96b0939a7229eb2bccf92696759a918ab388db28da665749904c0791c56532e7f8bf8e0dfe9ac16abfef3c476700a7c7cd137426c1c34976013e912297ba7f085d18e257891dca0960883f037509ccd93ff1a963c3d542dbd24d0ffc967592007d2e055cb3a9391802e65f8574678ecd9583045cfdc2881e22316ea8aa41936ada7d834546fdb274e92bbad1d3c640b4190f62126b6d5cceb0c31d3592f9fa60e51e1ad2e01ea70a2aef1b87923dd92cbe404cc19f9c76c8a2cf50262055abac1a0f624f5b8f958ea2c51c0b9e5820bb143840d6cbc4c0b5a22f679979ac66b13134ed83a22a5b25a052bbd72843e6450d91c4a639645514402920440fb98ae4c80851fe5f43983e759601749d7ffd2c7a408b886f45587410cee512815a3168796b2f6e8f15cb58a2a79e0426a8c6de0dcc2987727781d461d6f7c167707e8e9d855e1d024a5cf82042cc0df7fa5e6e9601c1c5e1920a716";

    bytes public encodedInputs =
        hex"2b9b48c4c4fcec686d3144c6c3a6ea1172251bd6a43d915596485f19a4c57f21000000000000000000000000000000000000000000000000000000000000001f000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002a00000000000000000000000000000000000000000000000000000000000000290000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        21751835443792206817743884760833787228258571007278070660101464903993828157654,
        20951587674424957468341124912404862519155878053713857723217099777740915438745
    ];
    uint256[2][2] public BS = [
        [
            2149321906048059706043880905463891679766229649045475566468481151472887961440,
            9005744838989978706834600054575500457682665559213049434016780190128371150269
        ],
        [
            16652532113623035602426977791711646886834884224248171207946128260870324126178,
            15871439082238825208866892266095288925657833170009172554758536222528612191467
        ]
    ];
    uint256[2] public KRS = [
        5515788362103117034173768695177154998397366143777310477692845563160107117706,
        20334659314665777627850590386068661690134953840022707742001926480908430084954
    ];

    uint256[2] public commitments = [
        15814120428326165602389706827044584199056193061705973658794617260695848571204,
        1162614727542345218178700669108184462101321416248720237849961050893658195777
    ];

    uint256[2] public commitmentPok = [
        5848823906547057492104467462170259767766656047334576429470440104719331892550,
        13314049055306371249662984277454939602793917684714572907325858493349085882134
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
