// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

abstract contract TestHelpers {
    uint256 public constant ROOT_HASH_BEFORE =
        7032340283086585880861647366443810304707403891235619959310912991801659997688;
    uint256 public constant ROOT_HASH_AFTER =
        7384572664550429695003907929599672995136478284639697609692654896457909561352;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        4033651216511111400175039953981700650715262298138585644694383935611028748527;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 0;
    uint256 public constant CENSUS_ROOT = 18565236114904902810648657293231472124579609911361453612213297312660844343981;
    uint256 public constant BLOB_EVALUATION_POINT_Z =
        38242756523272807551714356037080441478028494296701068869010837704451898429;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L1 = 371762577589005684;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L2 = 15381718359988807054;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L3 = 36765840227053040;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L4 = 4199814078534886736;

    bytes public constant BLOB_COMMITMENT =
        hex"a3e62e064a88d4f48951d0d75e2a1cd0d3a6259ecd44e2bd36446a21875f5cc23a2facf97a47494c646b2e5c60325d3f";
    bytes public constant BLOB_PROOF =
        hex"a9ab3d53ad96649bc82af65625341e947ad41313bd83cce44ef1af385c2691b09e189b05334958c972b327f7a4bf04f5";

    bytes public constant stateTransitionZKProof =
        hex"12f940778b7efa614b2c23bfbd1308b9d3413ea712639c9d54941d195f460f360a857375635ee68bde4b4402c075206246ffaec8cd65acf80d79afe20363df8c04c59d54f14c616f2a03d1264cb01f3a0acfdfe2e813faeeb62de077c9ab98ff152b4daf2004da578d83093eeac06852895fe88c3a2f6b3b55c7634922cbf8bd1354b7e62a90693a18f622c239dc3e27e46e0b90963cab12140598392deabf5f249fe7f6aa237b617c6c1d2b8103afd218fde03684fc4f31c5b17482bf0930390fdf182b0b4549d1f79373483fd754238774a885d2af17fe0e23906dc2c417b015830bdf12f3d388e6ebc2b51a442d98aebbc3fd41278064a3e1f45130a70eae238cfef5118e241c72616a64ed2775d321d09892af64131cfa312e51071580381ee26e1b3d88c9f61f7284cf7cf3ed95d83bfe27c1ed8c97dfb2006d7cd0a359169ef56ce6dd2ec18cc79d512ca70f1593f711cdba124c977d06ba90517cba0e1a6274e9c8730d63c9578e41a396723e9f26bd7ba877caeed2d17c9747c04999";

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
