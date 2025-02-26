// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {console} from "forge-std/console.sol";
import {stdError} from "forge-std/StdError.sol";
import {TendermintZKVerifier} from "../src/TendermintZKVerifier.sol";

contract TendermintZKVerifierTest is Test {
    
    // uint256 constant public ROOT_HASH_BEFORE = 2773072458553873709201164114201768243066148425063652934942402475964298215834;
    // uint256 constant public ROOT_HASH_AFTER = 3033651216511111400175039953981700650715262298138585644694383935611028748527;
    // uint256 constant public ROOT_HASH_AFTER_BAD = 4033651216511111400175039953981700650715262298138585644694383935611028748527;
    // uint256 constant public NUM_NEW_VOTES = 5;
    // uint256 constant public NUM_OVERWRITES = 0;

    uint256[2] public AR = [
        20222927412171929846340323185531585839100270540213228660527904900187000592853,
        3178734298704144059539150876351635505988609316794635995019677492637812113175
    ];
    uint256[2][2] public BS = [
        [
            7605439213418656917659879524804083292387798897804916381428950879086070795072,
            3653689718424821676369663629502171612324365455333125016800197816797237482851
        ],
        [
            1584725744552237632465300955017020398815489340871347784582529319036386689856,
            9591792697590990967115349874365751514589658199007804119987783469834178912520
        ]
    ];
    uint256[2] public KRS = [
        4287489626916526198565203942872201229280824880379861773678216014159081541113,
        21265835527296322046893052336739508021791250052513348109549531531423238777668
    ];

    uint256[8] public proof = [
        AR[0], AR[1],
        BS[0][0], BS[0][1], BS[1][0], BS[1][1],
        KRS[0], KRS[1]
    ];

    uint256[2] public commitments = [
        21403487034752347650012951979726008418183987569978399883282167509882435011382,
        2927529764782243306466112172429128846140111075079615581336908625977979713718
    ];

    uint256[2] public commitmentPok = [
        6640085190073830525988067786465408723512768158067254055676682769326485437677,
        152980611164891181481406554278041451465839241854908358301492087390485542954
    ];

    bytes public zkp = "2cb5c5736bdc882456f13091cad38d78e2845a53d92b2920a86d909ff2b6d9d5070719901f1d8f0befb5def04011cb666c353a1844388b79deb5a5e38e07bf1710d086ab174664a15ceec8472550ea6981a40cc6a62fe593b3f23f08c8e90b400813ea43c349f7d3b34fb10ba504ddca6c1710a015c52104887d4538617215630380ec48c9d86cc641998efa9a07b032b2120d2bc777894eeb6c54bddb0643401534c324b342670ffa1f965bdee5cec6f50f01463daa1673ec084cde5d63e108097aa214f3a0d6e84c337a7f9fa6028263d854f6560f5978b66856d0337a8df92f04094f5fa0b60d2596efded1c6ce2eb4297a57275cf012a4a6d8e5ecaeeb442f51f1c18dfb80785c53f4b043aa1cd8e8032e959339dad036aaae7183f003360678ec521937d1772355db164326859c8955149a79dea53b8b644ac15e8c38b60eae27b047f65f0aaec8a52b0149b487cdf8c56c7126014fca10bac88f227ced0056957e3d1a9d41dece1644e1b8662ef730cf30813ccb7fa720a3b23caaf82a";

    // uint256[4] public input = [
    //     ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES
    // ];

    uint256[3] public input = [1,1,1];
    
    TendermintZKVerifier public tendermintZKVerifier;

    function setUp() public {
        tendermintZKVerifier = new TendermintZKVerifier();
    }

    function test_Verify_OK() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) =  abi.decode(zkp, (uint256[8], uint256[2], uint256[2]));
        
        console.logString("Decoded");
        
        for (uint i = 0; i < 8; i++) {
            console.logUint(proof[i]);
        }
        for (uint i = 0; i < 2; i++) {
            console.logUint(commitments[i]);
        }
        for (uint i = 0; i < 2; i++) {
            console.logUint(commitmentPok[i]);
        }

        console.logString("Encoded");
        
        for (uint i = 0; i < 8; i++) {
            console.logUint(_proof[i]);
        }
        for (uint i = 0; i < 2; i++) {
            console.logUint(_commitments[i]);
        }
        for (uint i = 0; i < 2; i++) {
            console.logUint(_commitmentPok[i]);
        }
        //tendermintZKVerifier.verifyProof(proof, commitments, commitmentPok, input);
        tendermintZKVerifier.verifyProof(_proof, _commitments, _commitmentPok, input);
    }

    // function test_Verify_Fail() public {
    //     uint256[4] memory inputBad = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER_BAD, NUM_NEW_VOTES, NUM_OVERWRITES];
    //     vm.expectRevert();
    // tendermintZKVerifier.verifyProof(proof, inputBad);
    // }
}