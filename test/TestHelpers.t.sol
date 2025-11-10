// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

abstract contract TestHelpers {
    uint256 public constant ROOT_HASH_BEFORE =
        18299372292841235153481344924842291634627461178673397032963174807282274013560;
    uint256 public constant ROOT_HASH_AFTER =
        4193942925472549488867074454787863110893414875494267714245348714942973222460;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 5;
    uint256 public constant BLOB_EVALUATION_POINT_Z =
        28184956704614141570129917478851580216356702704004469292802022185081790370;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L1 = 1861872938329520188;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L2 = 2802192535128096292;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L3 = 14239606864274154423;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L4 = 630891406967983951;

    bytes public constant BLOB_COMMITMENT =
        hex"b0f8bc221e5d2a0393cde0461e31b7a0cb70fd57b9b6d51ccbf8b92807fd5bff867874dfb44a7100b38b1868fd9657de";
    bytes public constant BLOB_PROOF =
        hex"8ce25251ee4c0af5ac3b4e244b868d3b3db06844169a0b3e3f90bc7c59b321a5014dd1ce56b258546cb28a33c381d42b";

    bytes public constant stateTransitionZKProof =
        hex"086ec5219abc982c5a8e6cd096bb646374e60184c7e39aed1601a8baedde22210072d91fd89f5636275b0541739dc3d409003b2fa64217f1f0a82b34f321e9aa06d22bdc176de2bdeca3c3fce1ee9a941193bf6cd54bfa2e70b3eb55fa5590831efca9355060e01ceba0cb937d59561dc2ca8a84d9e7b252bb0ce7b58d1744a803147837f5adc943acc272418ee4e4e5b9e718e007bba29f206cf0a9e5d120cb04976da1ecf8d800b9f44a88ce935933197d3d973fc628a5c76d3f04542dbaa12aa1de244013f7a60c6b06a3e639e19332a9bd0a37016f1a3e4d49161e2ee66f1c23566a57598e52ddaf10a08693bf27832e9d94ce869d5c4f5a54f3cbb2537f1b44ba87f35cd75fddd02918733e46020600db2bac48ac7698bb625f78bb08f6031622cbd0212ab4c25fd8bded0f0e0723cfb7ac718b9f37464131743a376f81255a8855f70b81da5cd7d29bb6948ba8b8108ab7c3fb2f19b63b0fbb77cfeb8e0a27fa3db5145182be67bfd4e826bafaf950cb24bc880a45ec9c923c374bb8cf";

    bytes public constant stateTransitionZKProof_Invalid =
        hex"23823aa8b274657807ec3b991e1a380bdc8788b247a694c42504d37cb1357b1217b4d0591e7b35f8b81cc964fdb73e1799988f09218551fa22d42df66838fb1113e797c5d862c96446c74f2c88ae80b596b66fc6735ce48d4cda62393e0fde44270489b86397b032072f0487c8179fcd59f4d7ff083f24be533970a1440347c71ce0655b83ab5c4024d8a8526abff73d9ba1fce2d282bc03db6fb62ee8c4af6c1a85a1821ad50b286c704497a481665bcd46c4b3b07dc7963b848a4048c21d9d179c7840aebee03dafc718d4e6b25d07e6d3d628ba0cc80f67a4284ac5b2084002401535732eb37a65da6965018b70c5df2836f88cc24f9048989d28892c1a9e1f34aa34b7368bc639a53194dfa0459905ff4c5e3f2e5d6ea8f59b612f747fc60e868adada940931e9c30b8ed302301ddc533b7a2bbb203c8a420c8d7f22c2942ea5e837f887be989a4c6b84d95ff788dde5297f497c1249b49111c9b9f03c2c0780deff25465b1e5dbefcbdb2ebf48826a4139d95033d631dcd1b793b870332";

    uint256[8] public results = [43, 40, 30, 39, 22, 0, 0, 0];

    function decodeStateTransitionInputs(
        bytes memory _encodedInputs
    ) internal pure returns (uint256[9] memory, bytes memory, bytes memory) {
        return abi.decode(_encodedInputs, (uint256[9], bytes, bytes));
    }

    function encodeStateTransitionInputs(
        uint256[9] memory inputs,
        bytes memory blobCommitment,
        bytes memory blobProof
    ) public pure returns (bytes memory) {
        return abi.encode(inputs, blobCommitment, blobProof);
    }

    function stateTransitionInputs() internal view returns (bytes memory) {
        return
            encodeStateTransitionInputs(
                [
                    ROOT_HASH_BEFORE,
                    ROOT_HASH_AFTER,
                    NUM_NEW_VOTES,
                    NUM_OVERWRITES,
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
