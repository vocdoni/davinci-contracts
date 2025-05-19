// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import { Test } from "forge-std/Test.sol";
import { stdError } from "forge-std/StdError.sol";
import { Groth16Verifier } from "../src/Groth16Verifier.sol";

contract Groth16VerifierTest is Test {
    uint256 public constant ROOT_HASH_BEFORE =
        14604444798545114219184112291507703832081104487914651663816669607190240959987;
    uint256 public constant ROOT_HASH_AFTER =
        18226585194652543680303489095944698048129925574222676985726996869319177647405;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;

    bytes public zkp =
        hex"0c07d1fa200162652c1432be2b64524fae8ea94aa523a02e04fabadb3706c7c60e3ef6ab9394aee502c94bf66fd6013c78614b80194ddb9738498b0c0fd589ef213ff6f9546e83f9ce31c051e33d8f493c3faee50641bf7bfa0c0293e6850b6428acb11796557f70555af9dd68fe492c325f3553bf4af89eb43a621440e87c4c1f165c6f9e2d9ffe9180e06938ed13dda476c139eda6e0226a3251d1f8564d32201ceee55155227d5a36d432cc6f8af6c87c92b89a7bb4804a837ac50e67930e230cdbdfb0a4d7e6e2752f53de9ac04449b862897b5db0c003a73e1b2ebab4371d7f943089cef554ed58fc50934d2197c924fdeb6b4f77e87d7ef0bcd8d3082e25833c49c60712d1955fc055a6f302769a784ce8adf719fd4aa966419c686bd82b37c3a5016405db96e644e5590ac98cbd92a4931cf444c269e3a80bbf1135ba1de48f00a7060484fdbe01dba1ba5797994b0170512cce78ed4482fa930007ba1e8fb338c102998fc2108fb434a3499dc27e8f0b718bc9ce974adee38635e571";

    bytes public encodedInputs =
        hex"2049d2a4bdcb60f7f241e4115494be9ea41e0ffa4b12f4f204c252c6b37c1df3284be1b319af8f126a739ea5abb992545b2b3238a93c969e26b91ceb5849b52d00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000";

    uint256[2] public AR = [
        5441571320796354190808079931811770545106117630057541334655448715235418359750,
        6443626853476251415754926315261050238091088708608334123169185382851077835247
    ];
    uint256[2][2] public BS = [
        [
            15039339919848555580707600264341134380984276149412000291428881228502130821988,
            18397633883486708422262961403567591395309542506249262502828138778782977260620
        ],
        [
            14061206911382478391415853331674509433695253694394400817563097027886343146802,
            14525131670501960140592102530836451329024306703989977946311507067182860702478
        ]
    ];
    uint256[2] public KRS = [
        15853669378312494074565344684128231013517119947539689492377335769368903726135,
        13342484953191266775566242048471474682550000686144613994481058503400545781806
    ];

    uint256[2] public commitments = [
        16967448456778482476623691952688263834598389460139798955636208398674061454296,
        19547979366718549396651269529046596926947691693958716034775943164427368478138
    ];

    uint256[2] public commitmentPok = [
        13520900707001416168629374769501105512942298684462342266534029569008636790714,
        13823281530440731289898885512902135975120555357072348180834401017154954978673
    ];

    uint256[8] public proof = [AR[0], AR[1], BS[0][0], BS[0][1], BS[1][0], BS[1][1], KRS[0], KRS[1]];

    Groth16Verifier public groth16Verifier;

    function setUp() public {
        groth16Verifier = new Groth16Verifier();
    }

    function test_Verify_OK() public view {
        uint256[4] memory input = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES];
        groth16Verifier.verifyProof(proof, commitments, commitmentPok, input);
    }

    function test_Verify_OK_ABIEncoded() public view {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory input = decodeInputs(encodedInputs);
        groth16Verifier.verifyProof(_proof, _commitments, _commitmentPok, input);
    }

    function test_Verify_Fail() public {
        (uint256[8] memory _proof, uint256[2] memory _commitments, uint256[2] memory _commitmentPok) = decodeProof(zkp);
        uint256[4] memory inputBad = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER_BAD, NUM_NEW_VOTES, NUM_OVERWRITES];
        vm.expectRevert();
        groth16Verifier.verifyProof(_proof, _commitments, _commitmentPok, inputBad);
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
