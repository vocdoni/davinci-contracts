// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {console} from "forge-std/console.sol";
import {stdError} from "forge-std/StdError.sol";
import {Groth16Verifier2} from "../src/Groth16Verifier2.sol";

contract Groth16VerifierTest is Test {
    
    // uint256 constant public ROOT_HASH_BEFORE = 2773072458553873709201164114201768243066148425063652934942402475964298215834;
    // uint256 constant public ROOT_HASH_AFTER = 3033651216511111400175039953981700650715262298138585644694383935611028748527;
    // uint256 constant public ROOT_HASH_AFTER_BAD = 4033651216511111400175039953981700650715262298138585644694383935611028748527;
    // uint256 constant public NUM_NEW_VOTES = 5;
    // uint256 constant public NUM_OVERWRITES = 0;

    uint256[2] public AR = [
        1789547992475670499119714582072356751624746211722303505554450495454326064538,
        17607690026071408010814650961065395666991958580177594424792657274425062438206
    ];
    uint256[2][2] public BS = [
        [
            1171463437421115354197253097910525673197539969978697354025567719052967164962,
            3176597828154874263226368201690634054439427106779514583514347899806539556395
        ],
        [
            8955387325562334364642107725476596604072040214261939539454671385292170625647,
            21365311841324878777594025609538204248013939158859038616351883652442082880689
        ]
    ];
    uint256[2] public KRS = [
        11022294906613436550300690733354558139842952409817973826755059992566063463316,
        5736785961929767951974648882363323865607386967916583362142768872698131395818
    ];

    uint256[8] public proof = [
        AR[0], AR[1],
        BS[0][0], BS[0][1], BS[1][0], BS[1][1],
        KRS[0], KRS[1]
    ];

    uint256[2] public commitments = [
        7321318609717622877963739851020817867366094614587137326955634146738807283115,
        15583562541033382230085852674096985608241908228890862190107747045979597505447
    ];

    uint256[2] public commitmentPok = [
        18391086158343838917709023254500434590540583586083651153223067641686855287227,
        4752067756303322069167914960130123525059396603900405780910746163331108237562
    ];

    bytes public zkp = "03f4d92834d7574c4a5cfd88aa182f86067ed2e81473428f22f5cc159f10859a26ed997170e671a8c214b64b5aab86ddb8f3bdae9a752a3739c83ab3022bb53e02970659dfb67d091c77ee1896802777bc81e4585a69933afeae4bd59f583c220705e40206472e69f382a320c608f5a39e496417ac7812ede8f185ebb15fa62b13cc91d17f6553b08db1ff890fb7b75627317b59ab96ca6dac19c72be078826f2f3c5684ee959c2aa5474c7f1dfed43374101c4a977e1796b88e4d5eec5dacb1185e65d89f7db62f10e903ebfe1e5b63367172d16fe22d20698bcf50568c7b940caee7e1c06d8b33c9cc7a579f858fced9698fca5560147b5027208c605634ea00000001102fb8302edd9a102ca2cae7906c1b71ef600425d5803b5df30f0ff8f0e09dab2273fbdc44aacd437345384e14c20a4582d1fd785dd0c2a263da7a5a245703a728a8fc62f6a8c4318505da32b73b5be6752bd2893b1815ab1a5fdc467dc081bb0a819335806b410092de90557ebdfdc3b5cd3f328d5de07193a9ee941118f0fa";

    // uint256[4] public input = [
    //     ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES
    // ];

    uint256[3] public input = [1,1,1];
    
    Groth16Verifier2 public groth16Verifier2;

    function setUp() public {
        groth16Verifier2 = new Groth16Verifier2();
    }

    function test_Verify_OK() public {
        (uint256[4] memory _proof, uint256[1] memory _commitments, uint256 _commitmentPok) =  abi.decode(zkp, (uint256[4], uint256[1], uint256));
        
        // console.logString("Decoded");
        
        // for (uint i = 0; i < 8; i++) {
        //     console.logUint(proof[i]);
        // }
        // for (uint i = 0; i < 2; i++) {
        //     console.logUint(commitments[i]);
        // }
        // for (uint i = 0; i < 2; i++) {
        //     console.logUint(commitmentPok[i]);
        // }

        // console.logString("Encoded");
        
        // for (uint i = 0; i < 8; i++) {
        //     console.logUint(_proof[i]);
        // }
        // for (uint i = 0; i < 2; i++) {
        //     console.logUint(_commitments[i]);
        // }
        // for (uint i = 0; i < 2; i++) {
        //     console.logUint(_commitmentPok[i]);
        // }
        groth16Verifier2.verifyProof(proof, commitments, commitmentPok, input);
        //groth16Verifier2.verifyProof(_proof, _commitments, _commitmentPok, input);
        //groth16Verifier2.verifyCompressedProof(proof, commitments, commitmentPok, input);
        //groth16Verifier2.verifyCompressedProof(_proof, _commitments, _commitmentPok, input);
    }

    // function test_Verify_Fail() public {
    //     uint256[4] memory inputBad = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER_BAD, NUM_NEW_VOTES, NUM_OVERWRITES];
    //     vm.expectRevert();
    // groth16Verifier2.verifyProof(proof, inputBad);
    // }
}