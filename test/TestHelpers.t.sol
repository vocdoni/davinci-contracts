// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

abstract contract TestHelpers {
    uint256 public constant ROOT_HASH_BEFORE =
        16980121122017757910180031700420307374330809959175493496324824819030387103629;
    uint256 public constant ROOT_HASH_AFTER =
        16190605552028828311968978748997920696068614572573554591582206967422535362166;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;
    uint256 public constant CENSUS_ROOT = 5635465933697407994856253033999012801458042285248466599268679150980884016331;
    uint256 public constant BLOB_EVALUATION_POINT_Z =
        922219588160357198618336950274993316083825129697523780768115598500162165797;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L1 = 6730448159929965864;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L2 = 14515305029396022914;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L3 = 15132511764768392958;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L4 = 8071604173454021060;

    bytes public constant BLOB_COMMITMENT =
        hex"b9938b99e188bb368cbc85779769746dc57b6b59da7ff185716dd571b3376b308f130e4fe453dfe68574d2493860903e";
    bytes public constant BLOB_PROOF =
        hex"973ae57991e4a3dfc2687d24b16d7233c5f2cb1cd4f69da9b30253ddd793a2dcbe1d391743463d4205a2611423ac5bf9";

    bytes public constant stateTransitionZKProof =
        hex"09fd51e8e09aebcd29cb38fa3c120182718595879e6d9133fa7238bb9af03b481d9204792d82049ae46e275488ab33284e4824c0c6ff27299417b6b08992e215108bc427924ee617f853f4c3753476424c506de4528218c7c765d0657052477b2dff96bba80a4f543500c0e2caaed6c55cf50eafd3647027b947307e92a289d1122783a1481026e9b753153ae1b6bebb669fbc4b144925f943c3e7055cd6d9681c746c82b490c7d621ca8a965ead2f3fb6ffafecb41e0d6d46772f7b5cfb2c6324ed1b3fdc1fdf92315d7da3ff492f295e80d19d24c3e3f501235c9a674d7576053cb8f5f9d94c963fdcc911399084558a64165a29b75d7e1d373a45dfd401b70a6080dc6302dfebf5aafdaf981f8c76af7d16026737b73867022d98548c972300b48448bde7942f7817ecb3e59c88e47bafcc97a5fd2524be3850d851ac5347002b128b9c823fb5302098df055086a33efd75647c7cbe234ed0a5e6e68fdecd28350bd000649063b3d1fb92ef9a875ce9afa6ca505efeec069d9a7913ffd032";

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

    function stateTransitionInputs() internal view returns (bytes memory) {
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
