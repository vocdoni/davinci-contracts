// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

abstract contract TestHelpers {
    uint256 public constant ROOT_HASH_BEFORE =
        6980206406614621291864198316968348419717519918519760483937482600927519745732;
    uint256 public constant ROOT_HASH_AFTER =
        3521621988802002318964005439092080066299578336508168327546544274269862942207;
    uint256 public constant ROOT_HASH_AFTER_BAD =
        16148210138652248397115915945807375483400710386723928656748369930999699296491;
    uint256 public constant NUM_NEW_VOTES = 5;
    uint256 public constant NUM_OVERWRITES = 5;
    uint256 public constant CENSUS_ROOT = 9308290556862861573687918542897313508232635323659582158781603992249330983808;
    uint256 public constant BLOB_EVALUATION_POINT_Z =
        599167127392922207579004858769981027405123358559235195094224755125660299289;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L1 = 13638905466418695812;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L2 = 15436936371128404306;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L3 = 4649852911866324234;
    uint256 public constant BLOB_EVALUATION_POINT_Y_L4 = 4920605713834212623;

    bytes public constant BLOB_COMMITMENT =
        hex"92d34aa730267b086182746e1a0880859a7e4bf69d5667316eb2d33deae5966a55d1b8b9422fdbf07215323125fc2237";
    bytes public constant BLOB_PROOF =
        hex"917240dacf3c97d4bcdd1b74d76e580058c7d75feb6c8bfe109c73152b2aca537c4f2be8152c456354f608bf912af436";

    bytes32 public constant BLOB_VERSIONEDHASH = hex"01da6cff0db21ff16065c057b2597c3a22b02180c64168fdf6a2892daa8852d4";

    /// @dev KZG precompile address as per EIP-4844
    address public constant KZG_PRECOMPILE = address(0x0A);
    /// @dev Number of field elements per blob (as per EIP-4844)
    uint256 public constant FIELD_ELEMENTS_PER_BLOB = 4096;
    /// @dev The modulus used in the BLS signature scheme.
    uint256 public constant BLS_MODULUS = 52435875175126190479447740508185965837690552500527637822603658699938581184513;

    bytes public constant stateTransitionZKProof =
        hex"2530c0de108f2181ad3bd8b88b38c2eb458e4255644bfa34a14674b7c7944f2e1b9aec3a723ae2d6831ecc7152a272abf62a1e679a7a8fd9c1b5ea20c9809c7b18a40053c854925dec7b91e6238c36c5729da31fc8eb990b97affe2d5e219214215387239451d3363ee24e23391ad120bc5b179c4bb111e2c865767dd8725576011c8aecde19989c19342076536998b292684c39625bdfe532246b06ee2417bd00f145e6770afe7ba03ffdf41bc11531184fd6e3a99ca2d0bef92050b3068a71259c3a0a755cd81293743cbb99281dc69fb1645d9494dd9fe23826ae363b8a2822f40ddc8c318230a4de3b1271bca53b1d9a4de7bb311f588902797ba1d1e8fc26ec59c46ddd15bc2295b656a97d99a290150020036571d47c4cbfa040fece7b2a57ec923302b6c60a4513827bc598828cd85f2f964599fec923652b424ab9130a2a1ed4c87e007a5bf5ccc4062c0a039a49f288673df57af00f4475ceb03bec1b8bc238ee00ef111dee6074ad4126e2cdaa0f2531fa53e802818193c1dc5ceb";

    bytes public constant stateTransitionZKProof_Invalid =
        hex"23157cb0f5de6132aef9a69670ded82d32a7117d70d2641e21f1dc3c7950abd91b4abed9fc10894c52473e6ce0f883bbe48b53e5d8819832fa1e70642f40618824efcc5d80a2e710cdf443f8887209be3ace0c6b697efaf6c6fd37d9961d8cb610677d1a2d2f9fc402f5c182e8e7a7a6a4f4d70a2ea33c1b249530b61a5cea950879a6fefd5096f64070cd5e5bd9354c8397039522b5b2ce0ef70a3af0e5cc1a106457c18e4549cb4bb4c0b2ad6d61446fa4b1a8438eeb1f8a302fbd82d346db0111b856b981c93eeca59d74704d5485e8e96239a92271da88ba52a31ce190f8148595e0c6572fb9072b38a63df6b684793dfe5bb23e1c639cb647ad6efcefe41b64ca52189124df74e92651c34b3b6703ea11e8a05e79a3a4b8eca9787f9ab1093be0e7f9f98f60bda25760cb57ed74ae20413cf3073082d31825f5fe3d880c12d479518aa464ca3038691ef2fa74b1ecb46f32b0d3bece08dbd292d2d87815250ff61e8312580f9af186cb729ca616d4095222b1a4616ebc8f93547de1b9e9";

    uint256[8] public results = [38, 27, 31, 40, 43, 0, 0, 0];

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
