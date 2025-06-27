// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        9959529712980071391776980725980743246464286289977480485264915201285808726643;
    uint256 public constant ROOT_HASH_AFTER =
        12682003531705660610074706570328067259249936297413194117693986188674465425187;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 5;

    bytes public zkp =
        hex"005699b8baf41495cabc2f0fb63afc4361babe681ee2549f93822ef048fd8e5f0e6c314ce9cdc7efa0c1e5c11f2d3cc9dfaa037c0dc1cfaaee04093c551bf8f828151dbe684f08e922289420cef5e4d1207aa6451dd27618f5dc548a1da2ae59188cd0cc991ac8a5c11f4307f02f8c9cd7a6ebed7d811006053a324a3b6902a62737e479a360a13c70d97bec2ce756ceebda5449d278075c9a5ed9cb775794740fd97d58037e5ed76a14363d13c3406bf4cabc1f8b7fc968cfdea693fb1e799a03d965be7e447c30a5d4a5c743b3be7ab1466b1665954b82cb64cf2006f9249100cb2b09eef198db726779b852b353cdcbc63507991de67a20f4df2f70e011300a03db667a6f0c4af1a9e74ff046f35846dfab0bf3cb74280c4e9aa91fed35362ea9fb9f1b7c7d6db01504eb794acb3ee3c89fccd7790f58e3af1f8e6b9b00772cebbe9327f28bedbb32d1626651bef57d9ec9cf648dde45e206e495f2ce89010e04a637078ca3cea563333742b9173537a9a7584bcd1c6ab9615399a3676b92";

    bytes public encodedInputs =
        hex"1604e4e0b1828996b4d2d7a9859b8b252e52afdd7d9dc33e1d9bac795edc22731c09c2770746b03396a1c304f2c0b3001be50c46444a7f73845f9add41da672300000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000005";

    uint256[2] public AR = [
        153009795080681871689687386681049957364277728013004608799763545137453305439,
        6523539622311151738072161950628292072614435930568885704598083322590030199032
    ];
    uint256[2][2] public BS = [
        [
            18129823015709916658269276137117648851324893444362357665849873863607050219097,
            11104308034260418496077562096453936871720199256461150968646602913146273858214
        ],
        [
            17738954560836470006226496931350113295943659890721962680142085343527136892020,
            7168963632942483427764540589135722772371119091304364244483770674877604723098
        ]
    ];
    uint256[2] public KRS = [
        1741046570875123063600833194172798190802268152843101030007060521551853593745,
        358966997046156716272430600952184339275483364745787000443337259689871151408
    ];

    uint256[2] public commitments = [
        4529943272285288308273974867207418474058194618007884221046403473089227470134,
        21106724816637009397440005765432975584029284575950621939151508875202193457271
    ];

    uint256[2] public commitmentPok = [
        20318289697011648941585020960599540149047070564548187827382195283254141946113,
        6340594441910518830644378238430502581535600116002536380339375944417099082642
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
