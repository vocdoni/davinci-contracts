// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

abstract contract TestHelpers {
    uint256 public constant ROOT_HASH_BEFORE =
        15331516407712863826696478307862449443970234474805325024220140359692581265418;
    uint256 public constant ROOT_HASH_AFTER =
        4273098974177774955897510947763663722585826160465176782764856851114362357749;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;
    uint256 public constant CENSUS_ROOT = 
        13029601842640297299988411423808295458337223470560016337807635638197351586928;
    uint256 public constant BLOB_EVALUATION_POINT_Z =
        1733415104459172530752502972670505734309631578562008385156125852337774017909;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L1 = 15152863485526121165;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L2 = 981881010312593211;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L3 = 9657939671070314833;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L4 = 1088533931541091291;

    bytes public constant BLOB_COMMITMENT =
        hex"acaee880ed71ef51ff395497006107ec7db05e5735336d1150f73a3edd95925e99a879346d05b42ff047c80fcf4df850";
    bytes public constant BLOB_PROOF =
        hex"90aaea42cc124df560ea804659cad7808aafdfa75dcef2735a46f936da25f3178cf3338df72da9a183371e1b4920a840";

    bytes public constant stateTransitionZKProof =
        hex"2408f06929dbc6c05d5dc916ac97a6a378d5bb5ff48047ae8d54ef355bad5b5b0e3a9cc6474824a8ce88954f605668a3330ff040d744f2071add1a9bc9187425142870756235f05373fea46c3f4b6cfadb415610a4e41ba4d6a48627f67cd32302e083bb83669aceccf10f46010dc57af452d5a5520f9bb2fe5c7fe667402e4c232d4f6728168dc7bdb6468a3d3f2409004ddd720f60c6ee0ed7683d3464ba7e16329dc5accfe103492591c49dc6a2e73d586335ccf293ccf4d1cc68cec024eb1da4f722bd103bc642b5db12183f164aabc7eb5ae131f8e0745b524c07bb726d136a35032cfb2e1a8224cf415fcf806943aa2528552a93d40e92ec3c2af1ca690a921734adfcea4181a22669929384325efdb9de82a71b36995d218d48f4b0ac22fe1340dc047a70921df1dea06152fa2d78f12bf5294a526566d5c599e885d3158649be707c6286ab1509c9f7821b26fda1b59ff06e8fca1bbca7be77641a431491159a425b1eb55a2ff040fb2dde1b153d7f8b04c59a9c814a462d3dd0b03c";

    bytes public constant stateTransitionZKProof_Invalid =
        hex"23823aa8b274657807ec3b991e1a380bdc8788b247a694c42504d37cb1357b1217b4d0591e7b35f8b81cc964fdb73e1799988f09218551fa22d42df66838fb1113e797c5d862c96446c74f2c88ae80b596b66fc6735ce48d4cda62393e0fde44270489b86397b032072f0487c8179fcd59f4d7ff083f24be533970a1440347c71ce0655b83ab5c4024d8a8526abff73d9ba1fce2d282bc03db6fb62ee8c4af6c1a85a1821ad50b286c704497a481665bcd46c4b3b07dc7963b848a4048c21d9d179c7840aebee03dafc718d4e6b25d07e6d3d628ba0cc80f67a4284ac5b2084002401535732eb37a65da6965018b70c5df2836f88cc24f9048989d28892c1a9e1f34aa34b7368bc639a53194dfa0459905ff4c5e3f2e5d6ea8f59b612f747fc60e868adada940931e9c30b8ed302301ddc533b7a2bbb203c8a420c8d7f22c2942ea5e837f887be989a4c6b84d95ff788dde5297f497c1249b49111c9b9f03c2c0780deff25465b1e5dbefcbdb2ebf48826a4139d95033d631dcd1b793b870332";

    uint256[8] public results = [43, 40, 30, 39, 22, 0, 0, 0];

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
