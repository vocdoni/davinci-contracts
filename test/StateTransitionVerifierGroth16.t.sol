// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        10633697268432566100547111228192675718879714909099079783139311163602236074767;
    uint256 public constant ROOT_HASH_AFTER =
        9856495895231824317321181173471266818915032605823360987155589090285852538736;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;

    bytes public zkp =
        hex"078ae3e6c6c404cb161927a4dc711893798845cccc3b226ec56d0da22e16e1cb2ea60a232e5028dd56fbbf1beb512cfa933552d76ce4eb481e643b645c641e022226f0c2536331c01e39e71dc4ed1d8d761408d5aa94206174a9af6a0405589013031e68c4a1de5007d3915249f7f685d90d9a4235943c90785675a2e8d7d35a29be5d62953b0ee9629e5535077c9fd46e5bd2989922308ee5dc0ad7dab48d8c21537b7a5eb721cb271622974c96c000992c53b7d5872c16082eacee81bbdd9c17d329bd56549866e695a25cbcab95d1a61f7124a751b7dbc1dc1fc9b6ea92af2727e2440e22d8f263298d520bd6059094e50cefcb4d4dc7dc26c99a62deb4bc12d7bf39651303fdb7d56edd404c159b68baa249ddac050b220dedd75a0a736816b60bcf0e6ee111aef390daf2a76574d00a3ecc49912e7626c4a31b0dcc01921dd5a77472474f9df1dea6620d54e6ca4b366178995c911ff03cdececfc344841048e3d00072b27542d626d1326d1e51f039b182f31edcbd066434eeb4d838a0";

    bytes public encodedInputs =
        hex"17827599204d66e404f804ab4a1d7f1f8a86af5196d5d2a718b2a3e13ffbeb0f15ca9437fcfa27637465824011e905ff484485742595c25132d2c027994bb77000000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        3411587753163207834951403808383211390818213421667342181797361459606509248971,
        21099757613522408094588140894207705825288132597959893002020495008710101573122
    ];
    uint256[2][2] public BS = [
        [
            15447438698427245363943083334387791136935484153842763957170038170102163396752,
            8599454541209051600737325743653870666591202452982092435059162522718882485082
        ],
        [
            18881172254422666462072999261995266665249153664159965343476327264668196572556,
            15073824523513236901906351300296159880897262690917692651283903755645290798492
        ]
    ];
    uint256[2] public KRS = [
        10776288324205162110415420934998745209718826237604705749344492659100311261871,
        17710669759713201661277520743562448238878613548924564783526005366327359747260
    ];

    uint256[2] public commitments = [
        8522823174339708665450125365251808609089834963971045305212521302336196997992,
        10272530336060286419328009060608179880461086346421550679790483691018165944722
    ];

    uint256[2] public commitmentPok = [
        13494566764741185434303428838013205024553838187966007895073922824589579338884,
        7365790870133117761931271418371584834680374442055793538726056983152484169888
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
