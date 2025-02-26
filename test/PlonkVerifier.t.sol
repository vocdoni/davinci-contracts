// // SPDX-License-Identifier: AGPL-3.0-or-later
// pragma solidity ^0.8.28;

// import {Test} from "forge-std/Test.sol";
// import {stdError} from "forge-std/StdError.sol";
// import {PlonkVerifier} from "../src/PlonkVerifier.sol";

// contract PlonkVerifierTest is Test {
//     // uint256 constant public ROOT_HASH_BEFORE = 2773072458553873709201164114201768243066148425063652934942402475964298215834;
//     // uint256 constant public ROOT_HASH_AFTER = 3033651216511111400175039953981700650715262298138585644694383935611028748527;
//     // uint256 constant public ROOT_HASH_AFTER_BAD = 4033651216511111400175039953981700650715262298138585644694383935611028748527;
//     // uint256 constant public NUM_NEW_VOTES = 5;
//     // uint256 constant public NUM_OVERWRITES = 0;

//     uint256[2] public AR = [
//         9519678411605436787658670087077116491899754926428117655678117685164774229990,
//         21467954838735231046509182721841306829800223967945975783326097976885379450608
//     ];
//     uint256[2][2] public BS = [
//         [
//             17813478590478232823486772662768485355237048345942876883292455784178807559939,
//             9370956913068583502230029815337620505908528701617174548049567574663591546892
//         ],
//         [
//             3622826544768847512029118664683092888566251157802021845624731170276667781942,
//             19368290908046627694949334779290155248067850734951107148139303237629827970543
//         ]
//     ];
//     uint256[2] public KRS = [
//         14706374278606958553512155557243443628476649075882370917757278326758691755453,
//         438890703671815274607282916296260186994864516265250512164331793183329587044
//     ];

//     // uint256[8] public proof = [
//     //     AR[0], AR[1],
//     //     BS[0][0], BS[0][1], BS[1][0], BS[1][1],
//     //     KRS[0], KRS[1]
//     // ];

//     uint256[2] public commitments = [0, 0];
//     uint256[2] public commitmentsPok = [0, 0];
//     uint256[] public input = [69];
//     bytes public proof = "2d4a305df1dacf9b25779080756fe6fd88a768e0f7e80b571c43879674bcf3b20bcb9f2417bc995dcc4ed49c8f947dcf7527493f0854fc3f7eccc17ab2307bb9003e2430af4410fb4720c55a2d53ea9db2f77be25be624b5964d71caa5b1b29a10645ab103dcef8a30a628d15b68abc2fbef384aa44b63d66b6b300216a4d86a01b7dfee7907378d0b2bf253ce22e7cd878eda789b80ad70b6f09f838704959a14964c3a1ed00950a66360fd4f7c853580f47c1da3b211f0ae423cef34418bdc1556012b2ee699a1826ab98f96ff11c8e7ed1f9151d189e98cded1afc480f0a623bdae6d27c32c44980321a5095312ab960d9d5ac7a3d63ed3f46811c4e6f9652ee1902602d04f9d370c99bab5c1384bbf4a9945a39d17d3aa50cf3ca4f0e5811badfd382c16fa718f0bae2486af0bebd675e8b70e1b232694eebf1874335ffe2e831a5e8f50f276e013464157dde1fde6ff3a8f362ea280c5b7a909f3bcd1692cd4e2d3e675f02098ced383b51b3677f16757ccec02289c6c999ea8fcb22ea000268dbe5b84628505d8cea8f2d78c8e0b4dc967aa636ebaffc9c68b6c8d204c1e099972fea6d6ddd6461d84e744cc344f310bbde98a761532dab0920672d0512a5c3e4d251271f4673076d43c1e69980f9c793afc4f4a70606d4cc8fd0e0282071cbc4c1967720aad008ee50d2aa857beb50cf56a69e5f1bb8b4fa79fb0a4d207551901467f6ff873a6e07e910447648fee789d01143d9d505a1b197d49e0f7223e7bfc18841a22ddad977415c88f22dc55fdbc3a21ccc6ec4509329950ae65133a22a2ab544ff741fe22d8438aefaa01e87642f4108a8167776e698755c075243e9d92983883ec7610ada9cbbcecfde58b50b5ea3e10e9dcc3b465225fda5718bb3b040286bf08da8a4c39574ae034fc62e702713e0ad6cdd396578a9462ed14b0b402f856b05144a549542785865849d0fe60b922646355641253d5c4a5a1252962d84894844e085bfcbeda86b3c1eced941ba8ac3ca0fa3c9fe74fb69d481ee561e76ff6bb497f5b402afb6e2c6dbfcc1e55e7c6ad2b157690c2c740760205440f7789a7d8af9e590e1eabe1a65475400f1e7b6d3dbfc1880ad9b7be002d053d12e1e32f1f251d51f8a86cf37948dbf12baf70cfc905364448277f63be9b0b495bac4b3f086c6352675746f1fa5ace304c906362cb7cefbd696bb420aad8";

//     // uint256[4] public input = [
//     //     ROOT_HASH_BEFORE, ROOT_HASH_AFTER, NUM_NEW_VOTES, NUM_OVERWRITES
//     // ];
    
//     PlonkVerifier public plonkVerifier;

//     function setUp() public {
//         plonkVerifier = new PlonkVerifier();
//     }

//     function test_Verify_OK() public {
//         // plonkVerifier.verifyProof(proof, commitments, commitmentsPok, input);
//         plonkVerifier.Verify(proof, input);
//     }

//     // function test_Verify_Fail() public {
//     //     uint256[4] memory inputBad = [ROOT_HASH_BEFORE, ROOT_HASH_AFTER_BAD, NUM_NEW_VOTES, NUM_OVERWRITES];
//     //     vm.expectRevert();
//     //     plonkVerifier.verifyProof(proof, inputBad);
//     // }
// }