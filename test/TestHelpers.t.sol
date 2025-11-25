// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

abstract contract TestHelpers {
    uint256 public constant ROOT_HASH_BEFORE =
        2570213685245022812819400607675929817022907648534268011919552608733582516600;
    uint256 public constant ROOT_HASH_AFTER =
        3002110162893369572337896915733180539392271285963756713877989564646742292824;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        2570213685245022812819400607675929817022907648534268011919552608733582516600;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 5;
    uint256 public constant CENSUS_ROOT = 8771235498615420563818426219944666757032730329862969307286274163034500165479;
    uint256 public constant BLOB_EVALUATION_POINT_Z =
        858662983223037549198120547542366930951127181346753043663999487464474691567;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L1 = 8878257946003448033;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L2 = 14165000158920093496;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L3 = 5326521082045595877;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L4 = 2674253753957999066;

    bytes public constant BLOB_COMMITMENT =
        hex"a010e021b299db8dffc2de9f77dbd1818002da86f2a733aa2cc608eab53a4cf58ce60744282b561f74780d662f755b34";
    bytes public constant BLOB_PROOF =
        hex"b6ddbfb1b41ba0eec428185b77865f814d65c313a308ce06c65caf22fd45978b78e075e614237983eb09215e4c575615";

    bytes32 public constant BLOB_VERSIONEDHASH = hex"01dba6ec0ba00974bef74e9522dfe72ae9a78b72f78358b8787f3e9e23d834e9";

    /// @dev KZG precompile address as per EIP-4844
    address public constant KZG_PRECOMPILE = address(0x0A);
    /// @dev Number of field elements per blob (as per EIP-4844)
    uint256 public constant FIELD_ELEMENTS_PER_BLOB = 4096;
    /// @dev The modulus used in the BLS signature scheme.
    uint256 public constant BLS_MODULUS = 52435875175126190479447740508185965837690552500527637822603658699938581184513;

    bytes public constant stateTransitionZKProof =
        hex"0a1fdc26234b29e04c8715eceffdd4f73301e190ec3d2c1e2ecf8f99bedd25b520fcaa1659eb71ef03c68e549f516d04101733bf17043136e435ffd863786a092a6d0b79ab2427988df14daf598dd5fab211a9804996d31868610d7adcfab5922de4fe300acde384c2f9853144332831b37ba2a3ba32642ee2dc32c088560e581d90c9983d7cf56c4f5703e0b479daee66e73dd2a35912ca558f6d7136a5264c0e0a5494a3dda044fa2c0e8aace5997c3c84b1737cd47db77d6358fa2b74cff70334acd029b838ee935c286fa5d7878fcbb6eeeb1bc1bda1a9c6874edef8b20202d7bc67129a3b4b9201a8d201cdc524415f3681cdfec1206112e4302e9cdb9c04140b0acc5a90e2693a5133deee0c30ad9d8379e7c75b787f7223dd24e7c51608980bf22e69481df0b75af7a92fb8359eb64fe06e2aaaae03e63a8bdf2a280217f6030e2883f76f56ef412f13a8d92f1164477e4d91721f541a7a19cad1e0fd27dbad3cbda7c253846670a43bba3514e244a68a307b39b44e06338bea036088";

    bytes public constant stateTransitionZKProof_Invalid =
        hex"1d78e4c0d6c84a00624ca4ca26829778fc62ac12c605602bc869c30c9a61eaf42f797dca4ef35d064652e9ef8cbe32e8316fc86fb32be5c7f9f02df33d819dd108db3dbbbbd433234c29db7abc0fdea6a8ec1c054e1481f01b883b82141fe99e1b375622b92fe030267d3b72b3d05453a0327b4d8a085309c98f6ddef386b8c9025ab33ab6c4101070deb5b3ab280ebeb1a97266e6259d41a7513ce18353939714f764ab416ff9c15660cb8e36bb18385f858b7517fc9ea1d6d6472af7d519252180f314df0a31b7c1d76cfd682001727ab9891b0e90b3e2dae80f6bd39ff6c80792969faf7e7accd77f6366c186ad9d6231a02fe45ec99630a178b739284a0709f1d1692e785ef907107de477746c54cc011cc414b36c05d85e0e8039147d19151b851b6399d1aa2be1ddcb58ebde72484c0c404c06c6295d92883919687ad8199b924f9ffc44c79de71edb95e2ea3441864762d80382b74e5bf8744ab206ed17f46fe4bef4d5a0688dd9fdde42baf0d1fc59e58c89c63084ac23813882d685";

    uint256[8] public results = [60, 27, 32, 36, 53, 0, 0, 0];

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
