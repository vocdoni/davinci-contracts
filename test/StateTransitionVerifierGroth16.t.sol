// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        11540529015229420508441040255455111908497518547236688677125823343202283604502;
    uint256 public constant ROOT_HASH_AFTER =
        1955711634442266218442623404753454420621019669225137434447114896701269920059;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 5;

    bytes public zkp =
        hex"2ec28afc568946445b8188ba7a9eed8c3503b57acf585e6803690a02f02da72f2ab7fe9a2ed0fa8158997b6c61771738727e8529242bec80448300ab29b679080f714ddabca5e637b0ea857ed01d6d31e34f65f47b9ea9c9caebaa5808f7e2b627b3035936c38b340889f60434f4ba6eb1208f68f85441afef15932785bca8b42475061108d4fed129aaf7654a1325f7351b0e9c6819f62d97f922964df7aa9327363d908548e015634ee6f95ed824be52b742a8d55e31255d152bca6cb10fb00305489c14f01784fb1d98eb4864ba6bc24dd245f06baffce8d597ea1e5a21be248ccce1880be3274f1059d05dc5f6f7bf53e52f6896ad498608e0e60ebef972276db33e5a454c61ec2fde9ceb0d33862f43e93c9ed753fa6ea8e16f3e192b8c2add78e00adac6c0ffd91bfe8e20b7e0e5cfb4a3bc566799ad371ba6e2a2cd2f05bb9fcf7ff49f2380559831d4d0f092d0e76032d26b6de16e3631dcce9d1e782c34b777dab451675bed9512ff95a34e59bf2f036319aeba17fcc61c0320c341";

    bytes public encodedInputs =
        hex"1983b53c0e048008bd702af3060684911a3866f56e08f74881143cdf854cca160452e4c03c0464b0d490346235926028e4f8708b358f059229596a7feceac93b00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000005";

    uint256[2] public AR = [
        21150118609413006174209871243847185094450593556773573746860220359769519400751,
        19322229853685838388557871280391010030567649080609042044142439169286846445832
    ];
    uint256[2][2] public BS = [
        [6984883778672987409100878733434326649989746471686296835522897717399656063670,17956489829784319477931669113655480936886283209154636946655674747891027323060],
        [
            16490025525303954583030194987107726114652992073360787428386675562395018439315,
            17736035739041410156766272531304160251539146118182729044548820971642437046192
        ]
    ];
    uint256[2] public KRS = [
        1366273914767353983980592316959689850861059469919347819617775504110185488830,
        16532035174636670997625987548207266292273435327758231256013834303783724579186
    ];

    uint256[2] public commitments = [
        17834024518427615858487297545819435874536408483860269102108676330603010403212,
        19388447091546001858184580473798747276252716935013095003469060050537656012079
    ];

    uint256[2] public commitmentPok = [
        2593067615883282157221040446073215086126861288810602487072692225439168274040,
        19994907635879571237265884712172422860328618123648884555233106592923826307905
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
