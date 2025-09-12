// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { StateTransitionVerifierGroth16 } from "../src/verifiers/StateTransitionVerifierGroth16.sol";

contract StateTransitionVerifierGroth16Test is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        1948660668130447904216243664725338719370673611882115968663146044679451280335;
    uint256 public constant ROOT_HASH_AFTER =
        15977845708338371169793025518832793983533753933654506631846328160167915014374;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 5;
    uint256 public constant BLOB_EVALUATION_POINT_Z = 825815966275299770668920226134104406255248154364423716199071467470333673677;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L1 = 4732179056639898907;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L2 = 7489020789064866881;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L3 = 2688621184006160023;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L4 = 7157299434657619676;

    bytes public zkp =
        hex"23823aa8b274657807ec3b991e1a380bdc8788b247a694c42504d37cb1357b1217b4d0591e7b35f8b81cc964fdb73e1799988f09218551fa22d42df66838fb1113e797c5d862c96446c74f2c88ae80b596b66fc6735ce48d4cda62393e0fde44270489b86397b032072f0487c8179fcd59f4d7ff083f24be533970a1440347c71ce0655b83ab5c4024d8a8526abff73d9ba1fce2d282bc03db6fb62ee8c4af6c1a85a1821ad50b286c704497a481665bcd46c4b3b07dc7963b848a4048c21d9d179c7840aebee03dafc718d4e6b25d07e6d3d628ba0cc80f67a4284ac5b2084002401535732eb37a65da6965018b70c5df2836f88cc24f9048989d28892c1a9e1f34aa34b7368bc639a53194dfa0459905ff4c5e3f2e5d6ea8f59b612f747fc60e868adada940931e9c30b8ed302301ddc533b7a2bbb203c8a420c8d7f22c2942ea5e837f887be989a4c6b84d95ff788dde5297f497c1249b49111c9b9f03c2c0780deff25465b1e5dbefcbdb2ebf48826a4139d95033d631dcd1b793b870332";

    bytes public encodedInputs =
        hex"044ee7215bbaefbb2a9b246b2a431c911c88431d83d064d83eb524ec6c0383cf235323e3c2d30ff0c284fe50e4fa3367828ae92a9730d6c2780c129830d9c8e60000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000000501d365309d4c152ee3a7b9cf0fc7de179d5bac672495522b836136244b7800cd00000000000000000000000000000000000000000000000041ac13c7a490751b00000000000000000000000000000000000000000000000067ee58baa7ba1c41000000000000000000000000000000000000000000000000254fe6b8d046fe970000000000000000000000000000000000000000000000006353d5ea31df1edc";

    uint256[2] public AR = [
        16061044668188138093273769721630187023306453925335357727992904049901536901906,
        10722665954960676669732520268339532655144934995299210117377351416502048848657
    ];
    uint256[2][2] public BS = [
        [9003133292641820449090119475873711577894717051289514934929177744235039809092,17648218993374517554187632481762942263835874679628954833602210925059575596999],
        [
            13061233046444393291275012356169849944819979931914543512895333522051344674668,
            11996239411561111855015259802049071323177335794861324973634972190913029348765
        ]
    ];
    uint256[2] public KRS = [
        10679653612921617011433773987067444853958302367899398485328230717039288911936,
        1017850286992930589170461502400628597249460061909163961456266275272627329694
    ];

    uint256[2] public commitments = [
        14114749071540503407926069064805317691509862701015681964572430134157382025158,
        6570095728129241502151401020124981196962794266583253540888088792059238072980
    ];

    uint256[2] public commitmentPok = [
        21099523514641499772381063612476018350994455621202014222492869056996416961580,
        3393885430775383084837352311587976790895446802162328940723296737643102405426
    ];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    StateTransitionVerifierGroth16 public stv;

    function setUp() public {
        stv = new StateTransitionVerifierGroth16();
    }

    function test_Verify_OK() public view {
        uint256[9] memory input = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES, BLOB_EVALUATION_POINT_Z, BLOB_EVALUATION_POINT_Y_L1, BLOB_EVALUATION_POINT_Y_L2, BLOB_EVALUATION_POINT_Y_L3, BLOB_EVALUATION_POINT_Y_L4];
        stv.verifyProof(proof, commitments, commitmentPok, input);
    }

    function test_Verify_OK_ABIEncoded() public view {
        stv.verifyProof(zkp, encodedInputs);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[9] memory inputBad = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER_BAD, NUM_NEW_VOTES, NUM_OVERWRITES, BLOB_EVALUATION_POINT_Z, BLOB_EVALUATION_POINT_Y_L1, BLOB_EVALUATION_POINT_Y_L2, BLOB_EVALUATION_POINT_Y_L3, BLOB_EVALUATION_POINT_Y_L4];
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
        bytes memory _encodedInputs = encodeInputs([ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES, BLOB_EVALUATION_POINT_Z, BLOB_EVALUATION_POINT_Y_L1, BLOB_EVALUATION_POINT_Y_L2, BLOB_EVALUATION_POINT_Y_L3, BLOB_EVALUATION_POINT_Y_L4]);
        assertEq(_encodedInputs, encodedInputs);
    }

    function test_Decode_Inputs() public view {
        uint256[9] memory inputs = decodeInputs(encodedInputs);
        if (
            inputs[0] != ROOT_HASH_BEFORE ||
            inputs[1] != ROOT_HASH_AFTER ||
            inputs[2] != NUM_NEW_VOTES ||
            inputs[3] != NUM_OVERWRITES ||
            inputs[4] != BLOB_EVALUATION_POINT_Z ||
            inputs[5] != BLOB_EVALUATION_POINT_Y_L1 ||
            inputs[6] != BLOB_EVALUATION_POINT_Y_L2 ||
            inputs[7] != BLOB_EVALUATION_POINT_Y_L3 ||
            inputs[8] != BLOB_EVALUATION_POINT_Y_L4
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
        uint256[9] memory inputs = abi.decode(_encodedInputs, (uint256[9]));
        return inputs;
    }

    function encodeInputs(uint256[9] memory inputs) public pure returns (bytes memory) {
        return abi.encode(inputs);
    }
}
