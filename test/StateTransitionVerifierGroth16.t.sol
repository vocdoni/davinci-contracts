// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { TestHelpers } from "./TestHelpers.t.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test, TestHelpers {
    uint256[2] public AR = [16821715178840840010892626430339508771016478604624368433428533629029610442542,12486171747568638521436034082922672340800029125892437664344067230870987119739];
    uint256[2][2] public BS = [
        [11145273543394870856453188271128080984772592160473742252644548533585700819476,15073905004599226830723769206496134595497691471246689583573915738305897583990],
        [502743393330190083286856071889872880420097829691848689011348302863014303677,426292576433938025205741738608094531209815498482222725480951279053974178417]
    ];
    uint256[2] public KRS = [17011604122933596794694590088776900835322036015143446466806262738302613031464,15809843204291845503295719170577272173663599265499454552806477828222436370684];
    uint256[2] public commitments = [17605483704596221426977260580699441727749524076885295716915316826890106228347,19152488088795007058617697574732346239647544293953316609824002368177217124627];
    uint256[2] public commitmentPok = [4597548851566745673753323113523989027666149648679005935596369930070879386604,12459379127365388220456249246602053014934428333995447479926171432613016526059];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    StateTransitionVerifierGroth16 public stv;

    function setUp() public {
        stv = new StateTransitionVerifierGroth16();
    }

    function test_Verify_OK() public view {
        uint256[10] memory input = [
            ROOT_HASH_BEFORE,
            ROOT_HASH_AFTER,
            NUM_NEW_VOTES,
            NUM_OVERWRITES,
            CENSUS_ROOT,
            BLOB_EVALUATION_POINT_Z,
            BLOB_EVALUATION_POINT_Y_L1,
            BLOB_EVALUATION_POINT_Y_L2,
            BLOB_EVALUATION_POINT_Y_L3,
            BLOB_EVALUATION_POINT_Y_L4
        ];
        stv.verifyProof(proof, commitments, commitmentPok, input);
    }

    function test_Verify_OK_ABIEncoded() public view {
        stv.verifyProof(stateTransitionZKProof, stateTransitionInputs());
    }

    function test_raw_sequencer_proof() public view {
        bytes
            memory test_proof = hex"055165a7bd817c672203c656fafb63d83ebfe6f20189405472675c85eaca1827139a14d8509942c9d50e9ee594b31c0cdf84d2a69da3b8133f2a1a2264b88ed91def08ac4e4c0a9ba227ef7f2c2847695359d10ecceb39de1b19cd7830e8db2e0922d53cc9dd962519bcabd6b8e1ac0a333e38753a77613c54af314ba5f2073c0991d9490592ec26ae9ae54590302a82b8619c64bce364d74f906a3dcd905bb00dbf29f43d56baa1849533ff5374d69037c37706f3cd5fdc7ab958ae034a67921e8028970394a6d8800945e7df30dcdfbf7562e66c48511b237128970b36367818c786b27d116690189388d5d23022e8bcb3bc30c50470b7c59ab3669eb8472509caa1cd77ed221638ccf6b3b1288ef4bed9b4da680b6f3444270430fbc131e103346593da1d2a5150d4301221a3099e58644113fee8ade86d8a7778a26056f71c2d0f5117ec3639bb720ab2e3c619a029964b194ac28ada1d075b3ebf52b16f0f45fc59092b8e563c6206caf851c3e26fb00f1369acc3491bbb5fe9df4103bc";
        bytes
            memory test_inputs = hex"19797a8cebe81c4f77491e79587e4801c082c267711b2b733b64f90d43798d8b22e1dcc956da63f65cddf43716a4e43e7ecd3da6f9218904c123a98fef6bba960000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000000020fd18d9eb24a571a75a38b1987eb2b0a237ca69ace798cd5cf69938f7ace68500b3890afda67259fa51887a9eaac449943bf2fd31f63e3e6b78732a545f4409000000000000000000000000000000000000000000000000245a6912eb42ae2500000000000000000000000000000000000000000000000069e8cca8ac340c640000000000000000000000000000000000000000000000000ef872126e0d136100000000000000000000000000000000000000000000000061a95c9993440404000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000003085a9b32d6f7f1b91d137ee1d584603d5d70555543a3548640cc899b6cab55d69e7333a3758c9ce64ccc9ac18b2d55534000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030902d64389d4d5b9e4ccc30e756a023938c5018da1bea0645caf8c6c0a8e6d6cb697bd6a75624ba618e623e1d4625306600000000000000000000000000000000";
        stv.verifyProof(test_proof, test_inputs);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(
            stateTransitionZKProof
        );
        uint256[10] memory inputBad = [
            ROOT_HASH_BEFORE,
            ROOT_HASH_AFTER_BAD,
            NUM_NEW_VOTES,
            NUM_OVERWRITES,
            CENSUS_ROOT,
            BLOB_EVALUATION_POINT_Z,
            BLOB_EVALUATION_POINT_Y_L1,
            BLOB_EVALUATION_POINT_Y_L2,
            BLOB_EVALUATION_POINT_Y_L3,
            BLOB_EVALUATION_POINT_Y_L4
        ];
        vm.expectRevert();
        stv.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
    }

    function test_Encode_Proof() public view {
        bytes memory encodedProof = encodeProof(proof, commitments, commitmentPok);
        if (keccak256(encodedProof) != keccak256(stateTransitionZKProof)) {
            revert();
        }
    }

    function test_Decode_Proof() public view {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(
            stateTransitionZKProof
        );
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

    function test_Decode_Inputs() public pure {
        (uint256[10] memory inputs, bytes memory blobCommitment, bytes memory blobProof) = decodeStateTransitionInputs(
            stateTransitionInputs()
        );
        if (
            inputs[0] != ROOT_HASH_BEFORE ||
            inputs[1] != ROOT_HASH_AFTER ||
            inputs[2] != NUM_NEW_VOTES ||
            inputs[3] != NUM_OVERWRITES ||
            inputs[4] != CENSUS_ROOT ||
            inputs[5] != BLOB_EVALUATION_POINT_Z ||
            inputs[6] != BLOB_EVALUATION_POINT_Y_L1 ||
            inputs[7] != BLOB_EVALUATION_POINT_Y_L2 ||
            inputs[8] != BLOB_EVALUATION_POINT_Y_L3 ||
            inputs[9] != BLOB_EVALUATION_POINT_Y_L4 ||
            keccak256(blobCommitment) != keccak256(BLOB_COMMITMENT) ||
            keccak256(blobProof) != keccak256(BLOB_PROOF)
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
}
