// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

abstract contract TestHelpers {
    uint256 public constant ROOT_HASH_BEFORE =
        16148210138652248397115915945807375483400710386723928656748369930999699296491;
    uint256 public constant ROOT_HASH_AFTER =
        19723816014752875800279306456496578536046296876388232350336782237768121417505;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        16148210138652248397115915945807375483400710386723928656748369930999699296491;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 5;
    uint256 public constant CENSUS_ROOT = 9308290556862861573687918542897313508232635323659582158781603992249330983808;
    uint256 public constant BLOB_EVALUATION_POINT_Z =
        258831730294192976160839780809862072644845493999173215524632449490950491783;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L1 = 1972749316818219786;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L2 = 16559905186189800966;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L3 = 8158107273131462392;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L4 = 2221210225551586153;

    bytes public constant BLOB_COMMITMENT =
        hex"a8a4d78fe15776833f87428bc093a0f1c3684605fe64cbe791a372b2bfe9f349c92a3d9719688a0ecafeb49ed9edeebd";
    bytes public constant BLOB_PROOF =
        hex"a8bc33f0adb47d3f9e3273f9f68ef96e414055d094c104b2290f47803581bd04fd33b52a7fc920985aa3af1fec58c49c";

    bytes public constant stateTransitionZKProof =
        hex"00e53c8167049f28eab4d9352b59e1379e357e4a8cd7b070ddee85f4c767baa1001a7110fcdbf19e2ba7bc794c0504fd21f95ebe2a8a97bda8eb6ddf653cf226267ccd3f08fefb8b90c5621753b2f7ef403172663984c5c4bb27bbf6b971cde5220db7bd2efd0de9c3208f439aa96f52c3bda840f0fb85a3ac676e4fe1a74a001b348ca610ad440bbdfda60d0e803233a001f6ee07baee0444124e9f082e10641bf6a437676dc5eee04391e2d4a1e582c2e4d44047826a254193daedaba3233b2f88cc0117eb857fc39484b658336f8d0578270b8bd8f32bc06e4245243f8e3e08860b64057c5f0da39cae5700cbab4dc785c86a334d21ebec37c860f20ffd7a249f1c54ff9e5127c74e89e799e770f85a17cf3072f4e7c04f7bf592a1b109ad19114e6304e3adc1663569832a251ffa2bf3b5f340d388af2dc7a7790f1a8c6d173c081ceb82feda29e2c4feb7c7019d8bd56c4bea503658f5f87711aaa2612812bfca2a38897d7a1c87f7cf95d78f9bb003d61e8c2a6499d21becc216399d85";

    bytes public constant stateTransitionZKProof_Invalid =
        hex"1a43b4c5ea8d187af5e5cb4dd1042f61519e8692f28fd57254b1b2c6d554114a0f90b51d98df08ed81b1bc64b63672e124e9f8c40f8ee885ee59390067d7f4831669a863100be1f0a34a9e4be66a4521260228aefb37aa0c537f0002b6a9658b19133bd7b15dc5cf35597fbeee3c186292de14d0ed19d12e4361de962ef45d201b24a92e9a702397551e2d615e654f9c1f35829a56926e55d04dfdfc2ef2dbed0b611f38cfd232d55a26f176736aed292bbbd99a2a5a2a19a8c9f2f9035da7a8050d76ed25f2e75b1da056dfb91353068245d621e5e42ba5d2cb2383dc5a29d50fdc8f7b046a8ebecc09ca193f74379240e333b6a474f7014ed38095c0d9d6340bb808966513f8b4d66ba399ce75306ec4da498993cdf21988e3df928038797b2023dd1a091e3f974ddcd14a179e9df8922e40e9fdf081e50c5893aeb713773503358c02e9797970cfbce6cd9f6033bde84c0ef7dc395b23f82b86a0fb13bd1408deaeac150d8d0e687275bbd3c95f9d270181ce689bdb28f5c413ed509ded67";

    uint256[8] public results = [31, 28, 42, 41, 48, 0, 0, 0];

    function decodeStateTransitionInputs(
        bytes memory _encodedInputs
    ) internal pure returns (uint256[10] memory, bytes memory, bytes memory) {
        return abi.decode(_encodedInputs, (uint256[10], bytes, bytes));
    }

    function encodeStateTransitionInputs(
        uint256[10] memory inputs,
        bytes memory blobCommitment,
        bytes memory blobProof
    ) public pure returns (bytes memory) {
        return abi.encode(inputs, blobCommitment, blobProof);
    }

    function stateTransitionInputs() internal pure returns (bytes memory) {
        return
            encodeStateTransitionInputs(
                [
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
                ],
                BLOB_COMMITMENT,
                BLOB_PROOF
            );
    }

    function resultsInputs() internal view returns (bytes memory) {
        return
            abi.encode(
                ROOT_HASH_AFTER, // After state root
                results[0],
                results[1],
                results[2],
                results[3],
                results[4],
                results[5],
                results[6],
                results[7]
            );
    }
}
