// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { ResultsVerifierGroth16 } from "../src/verifiers/ResultsVerifierGroth16.sol";

contract ResultsVerifierGroth16Test is Test {
    uint256 public constant STATE_ROOT = 7961306345560008268868495476665083487633676217546477646432696512080076649619;
    uint256[8] public results = [35,36,36,46,55,0,0,0];

    bytes public zkp =
        hex"1c4ab76936f3dc3518b83445231af31409e440ae65388786d72709f61720e84520254801c9160f9e006eb95910746080a3bcc2a2149d2588258fcd97e4d853d61847d112c11b9e9c6315b673d40026cc0162ab0fb21e6e9509614739ed4e779b0acba5e7734a7b6940bd1069d17dd752163a7e1cccd8c30e46e471979a9b26e80948598b02148e1a088bd1a2b85906f99ca40fd1a92a37c13e22e41cafe9bdf2213f2d2ae823915fbfa4b4e4eb390f1bba56ee963f38a35d8ba9a89aed044ed10a477913d37e54158aedbb469b2ae72c743e9b7387a99051f169c532d7a153c52ddcb0969f2a4ab7799f438a82452d11ce7eac50754ef67394f329f8138f9fd806948bf47f4cbc3151729ce5545711af2534cdccc79ce10b423189e98d0496762c57b02cf80452a807e34d83b029945b99d3f2555f4c967a20a89a415edc31d52c0361f44d05774b4d205064825a0ae337646293e9a8fa6c3a5128884a4ac8ea1f9face353bec75ca86911fa618df8a9cc61c60dcee1df117d17f96ceab555ad";

    bytes public encodedInputs =
        hex"1199f090a5032e8d2ea0cc67145c1652c37c96b1caf28d9f2051253a383cb493000000000000000000000000000000000000000000000000000000000000002300000000000000000000000000000000000000000000000000000000000000240000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000002e0000000000000000000000000000000000000000000000000000000000000037000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        12796772299288104050454121345809068151808322416182504978333164528699260856389,
        14539881469935086232192894719894227055247735593738646621266239297099019080662
    ];
    uint256[2][2] public BS = [
        [
            10982397478199812170565155550063424882214903645376764280964391742599620097947,
            4882943468019130147718399329183768999107975786206278114799420941657261680360
        ],
        [
            4198646628989972715946129732272866723557155543781580869800195662048054066674,
            15037947103679278063802743383702618682173002776263010628289640746535149129425
        ]
    ];
    uint256[2] public KRS = [
        4649410273251739558856650323107573559425582909939696438375599934740608668613,
        20744003308609319479389994276921187179864737271467779803180310934763762327512
    ];

    uint256[2] public commitments = [
        2976336391462217926564099624008927511456170359281042037893849781367021540982,
        20056696952013355680980384485249794795965491394902598425761715986038366417365
    ];

    uint256[2] public commitmentPok = [
        19907741934591976748335085443953393583292857641816873161579552115923775834346,
        14303820218479964605400048155725597059907239063044458417058305397754252580269
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
