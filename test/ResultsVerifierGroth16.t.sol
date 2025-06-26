// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 20672007792813623327193303115591878510251320530597528837663039205365355109159;
    uint256[8] public results = [53, 36, 50, 57, 40, 0, 0, 0];

    bytes public zkp =
        hex"184d0f61c6961066a3233c1779afe78a457591d98a72530661ef4fadb063051009b58c0e1a7af79186bf7ea98ee93f2ed2bd152a48d8e30267781d4f24cfdf5d299f94197ee127f165a723d0464b2edbff6aeca22313958c1b594b5f7e10e9e3159d0ff3262a2c06393e3e771efda55362a91ca400f23d69cd5206fb8c4e39590e1239226cd7e28a043318555692f31db718e6917fcfdf89caa2bb2678fb480f1be879f372983f2d2c9f007dabb138b75f3861fc59bebfc9b286446308a31afd136f7ac1ed707883d376432bd61b947194366241269bde21da1106004e03eec10a0b2df8b03d8346bf943d6bb2f6cc1819160bcdfcd4eadf9cfe675034c315841f4355b1a1718d9cb0171f8e51f83b3ee77e7070b1ae63c69b2dcee099fb9acd198dd0dd2c074480aa4e588f58ffc1e027b4e7dd6fdb44d861911ac42a8b0cf62870f092b2043dd915c135e2bb5f21eed53f8c0f65451e8948051ebcd9a668ab15a90a5620010ea16b53293c1724078b53aaa843b75a6190b1223f409bc9eab6";

    bytes public encodedInputs =
        hex"2db3f118858947af9b60f5052330ac6a3dbf8aec405406e17abf4f977063db2700000000000000000000000000000000000000000000000000000000000000350000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000003200000000000000000000000000000000000000000000000000000000000000390000000000000000000000000000000000000000000000000000000000000028000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        10991661752209924012953641889762801111029974460791013883566199087959254435088,
        4391581580690794377615171515745468396528611760600926003297199675676229099357
    ];
    uint256[2][2] public BS = [
        [
            18826777621033649281361859389273870838665964499070151847790615649172969286115,
            9776074890900258097243867987429043837147854150227354702953498381893629786457
        ],
        [
            6364577454974231504917881049083725854651616073894373652589588520613847975951,
            12623197105420011370421024795678783831593843754236328228745327961555106732797
        ]
    ];
    uint256[2] public KRS = [
        8790911388601758161636170140540405701584391440837029272671721175097979170497,
        4542881086757873371231399108464886865607020626107187544067506265457135129988
    ];

    uint256[2] public commitments = [
        14140668496773439097432170087658799841002616181780075731612041806599375526605,
        11558388176740488323187281990233810501838957495439205577579009527844983999734
    ];

    uint256[2] public commitmentPok = [
        18292061188608610526594391296285860959901830725712497593909880868366413359275,
        9797238313585450608043038243828817078136630293131406713340529073235419589302
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
