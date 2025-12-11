// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ResultsVerifierGroth16MetaData contains all meta data concerning the ResultsVerifierGroth16 contract.
var ResultsVerifierGroth16MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557611e7b908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461167e575080635d26278e14610c8d57806360e583461461037f578063b1c3a00e1461026e5763b8e72af614610055575f80fd5b346102365760403660031901126102365760043567ffffffffffffffff8111610236576100869036906004016116b6565b60243567ffffffffffffffff8111610236576100a69036906004016116b6565b61010093846040516100b882826116e4565b3690376040938480516100cb82826116e4565b3690378480516100db82826116e4565b3690378101610180828203126102365780601f8301121561023657845195610105610100886116e4565b8201868282116102365783905b82821061025e5750509061014061012c8261013394611bde565b9301611bde565b916101209081865161014582826116e4565b36903784019080858303126102365781601f860112156102365785519461016e610120876116e4565b8591810192831161023657905b82821061023a57505050303b15610236578351633072c1a360e11b8152945f600487015b6008821061022057505050906101bd6101c89261010487019061173e565b61014485019061173e565b5f61018484015b6009821061020a575050505f826102a481305afa90811561020157506101f3575080f35b6101ff91505f906116e4565b005b513d5f823e3d90fd5b60208060019285518152019301910190916101cf565b602080600192855181520193019101909161019f565b5f80fd5b813581526020918201910161017b565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610112565b3461023657610180366003190112610236573661010411610236573661014411610236573661018411610236576040516102a96080826116e4565b60803682376040519060206102be81846116e4565b803684376102d0602435600435611909565b82526102e660843560a4356044356064356119aa565b8284015260408301526102fd60e43560c435611909565b60608301526103126101243561010435611909565b83526103246101643561014435611909565b9060405192835f905b60048210610369575050506080830193905f945b600186106103545760c0858560a0820152f35b81806001928551815201930195019491610341565b825181529183019160019190910190830161032d565b34610236576102a036600319011261023657366101041161023657366101441161023657366101841161023657366102a4116102365760405160206103c481836116e4565b803683375f516020611e265f395f51905f5260405182810190610400816103f2610124356101043586611706565b03601f1981018352826116e4565b519020068252604051604061010482377f1130a6954a65ff3b4f282b1b3c5e16fb8d50c629f3c0576fc132988856495b5360408201527f2414c69ae892022b5423a30a93a27ef6641b8f9ac0df47c7a03437faf057a65960608201527f26c6aa8e9bc65fecada4eeebd1da7397d64b149e0be7f6f9fe895371b8b7b0df60808201527f0edf4f42bce5997864d32f8dd64a652d0aa9b2bf2797f7feb5eb60448bd127a160a0820152604061014460c08301377f012ba85e6411b7f069d1a23abc02719ccc008690612f32810057727071287ad26101008201527f296ff26d81e3d06d0fe1ee0807b357ece4902f7905abd71470cbecf0b05fd3d06101208201527f24a96b50c1cb621dd08ff1a739930b94e464e8a19149c0e459e86127c42a04e56101408201527f2da6d70ce908208c105dbbc340fa67401f5aea45ee04af156e2918a66bdb8f1661016082015281816101808160085afa90511615610c7e5760408051929061057090846116e4565b61010482845b6101448310610c6e5750505060405190604082017f06ff080d10e29a0f481dd0594addf732726dfa453179484dc81b4ff09cbdf86983525f516020611e265f395f51905f52604085808601977f1c9420d026b48bbd1de34903e506328c20224b702ec30670dd678234d80c436e89528051855201519260608601938452818660808160065afa947f19e1ec208a3f6c783b9272afe0cd2f3dba6b65d16a78cea0a05582440f411a9982527f2f843afcd2d1d364dc6835af76b74a3d61d1adb05d4c70d0e7daf1b0dbb9b82785527f1d399eb2e4260149fe79ee474ee5c71fdb849605606cd1bd3929b857135a74d160016101843560808a0198818a5287878760608160075afa92101616858a60808160065afa16167f2150086177878788cd454f01542f99a12af2fee0978d0f3c1c21cb2d475d046484527f27fd8102aeb1af0eb0f17a7e223244400f9ffaf32dda54c46c6de0c090b7b94b87526101a4359081895286868660608160075afa92101616848960808160065afa167f2170e721dc467f2dcd074b6dba21de5b1990b77f4334bc4953cebc2ecc77e83c84527f27130af6de00980852d653d469f1bc917dcff82c701aa0a247da34c8a7f232d187526101c4359081895286868660608160075afa92101616848960808160065afa167f1a370bfb4d9933d6bae8c65c4f41c1e4abf5e821ee472a5457c21b9328feae0c84527f06ac7394c4cf7456cc60e9da7d5a945e992ed0f7e4979aaa2abc22efbc9fc4c887526101e4359081895286868660608160075afa92101616848960808160065afa167f0d6725204f08a7318747880cc67e561ad2332d45a7232bdb8f09f757bff6fe5784527f2892e7493d2d56b90678fa529c8770a29a927aff7474b72a0e39c8db24d37a4f8752610204359081895286868660608160075afa92101616848960808160065afa167f28cf0719114007b55c52a71f0b3a1c2fb60781bb251766e04c996e5337413c2e84527f2701266463ff031d6927268f8798f7334b122f57da1123e5f601de175780ebd18752610224359081895286868660608160075afa92101616848960808160065afa167f0d6325e0cd52ee00aa1ec11ee038ce8c8607867778b17d3a1a1d22423725213484527f194efb11ab312fc89dfe9dc37d5a70c0de408de29891b140ddcb08b69b351ef48752610244359081895286868660608160075afa92101616848960808160065afa167f1053ea377fb03c51dcc8d4223d7b797c38d9a012313f0bae5d063add31009e4284527f0f3cc22d48f4c554f9f75856265543c2ed6f63397c26146786c0f6c72466e1f08752610264359081895286868660608160075afa92101616848960808160065afa167f01d1cf1bef078f0d9cd0bc4cf24034b389ea7e294856227dd5862e02987c026884527f07174e93bcca1980f209c404100f706747073da63dffbc6fc33498941b2ce14a8752610284359081895286868660608160075afa92101616848960808160065afa16957f297b9192cf377d9605ead8a1d09770bac7ccd521755b277a136ca5898e3156558452525180955260608160075afa9210161660408260808160065afa16905192519015610c5f5760405192610100600485377f09db1119a07ed79120dfaee28718f3e81cb504976ccb6c30214f83b598b9dc146101008501527f273a0e652d0ef4f8d3026d3db01c7406fb655ad69f7b127d963ec73b77fe51a46101208501527f2c806327f060d009b15feb81a511eb70e35e13051b32c9bb938f9d340e86bd746101408501527f087a66fab2082f62ab2bf0f36c4567488fc3db1487e262b4db9eecab52f57f616101608501527f0cb18dcdfd66083c47d4597cdd7f6d7d99115d97ff97b1cb2f62e266adee15d36101808501527f01ed124b6574d18301a94f683c4447538a091bd8cb8f22024c6452176a6420da6101a08501527f209fce5fd35738c10c170c45fb6633c86f6b4580dd69e2f38b76e1eaf4642fae6101c08501527f0e0bcb17a68fcb6ead6c0baccbfad4c507f6cd5a0a82ec8ed47248d92a1decca6101e08501527f14f0e892d8d3e91a1faa6eb7ecb7ec7a8843523c5a36b60c54cc6e2ab58a5b776102008501527f0721ad350acfac4e8ce26c326d142fda577e97b3815757197b0f6234b2d3d8aa6102208501526102408401526102608301527f23a1dbbdb025c9cef8e1a4eb04b0c3b2f6e063ca604cbeb0dc3ce59c800540ad6102808301527f1542860ca9c276bc73b143373cb7b868ef1e551f199efab4784b7d14cac9ebba6102a08301527f1882e97d1c74908c29fc04d1a81a6ef490b035d4d47acca24c4e6aaa7c2ccccc6102c08301527f0f0be2dcc6408c923f83a13319f860a94f111fe98799945021a3f49f5664741e6102e0830152816103008160085afa90511615610c5057005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191849101610576565b6351d49ff760e11b5f5260045ffd5b34610236576101e03660031901126102365736608411610236573660a41161023657366101e411610236576040516020610cc781836116e4565b803683376040805190610cda81836116e4565b8036833761030093815192610cef86856116e4565b85368537610cfe608435611765565b868301528152610d0f60a435611765565b905f516020611e265f395f51905f5283516103f2610d398a87015189519283918d83019586611706565b5190200684528251865286830151878701527f1130a6954a65ff3b4f282b1b3c5e16fb8d50c629f3c0576fc132988856495b53858701527f2414c69ae892022b5423a30a93a27ef6641b8f9ac0df47c7a03437faf057a65960608701527f26c6aa8e9bc65fecada4eeebd1da7397d64b149e0be7f6f9fe895371b8b7b0df60808701527f0edf4f42bce5997864d32f8dd64a652d0aa9b2bf2797f7feb5eb60448bd127a160a087015260c086015260e08501527f012ba85e6411b7f069d1a23abc02719ccc008690612f32810057727071287ad26101008501527f296ff26d81e3d06d0fe1ee0807b357ece4902f7905abd71470cbecf0b05fd3d06101208501527f24a96b50c1cb621dd08ff1a739930b94e464e8a19149c0e459e86127c42a04e56101408501527f2da6d70ce908208c105dbbc340fa67401f5aea45ee04af156e2918a66bdb8f16610160850152825185816101808760085afa90511615610c7e578490610ea9600435611765565b610eba6024959295356044356117d0565b9291909389610eca606435611765565b9890977f1d399eb2e4260149fe79ee474ee5c71fdb849605606cd1bd3929b857135a74d183519b8c937f06ff080d10e29a0f481dd0594addf732726dfa453179484dc81b4ff09cbdf86985527f1c9420d026b48bbd1de34903e506328c20224b702ec30670dd678234d80c436e828601528051868601520151606084019081527f297b9192cf377d9605ead8a1d09770bac7ccd521755b277a136ca5898e315655858560808160065afa957f19e1ec208a3f6c783b9272afe0cd2f3dba6b65d16a78cea0a05582440f411a99818701527f2f843afcd2d1d364dc6835af76b74a3d61d1adb05d4c70d0e7daf1b0dbb9b8278352600160c4356080880198818a525f516020611e265f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f2150086177878788cd454f01542f99a12af2fee0978d0f3c1c21cb2d475d0464828801527f27fd8102aeb1af0eb0f17a7e223244400f9ffaf32dda54c46c6de0c090b7b94b845260e435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2170e721dc467f2dcd074b6dba21de5b1990b77f4334bc4953cebc2ecc77e83c828801527f27130af6de00980852d653d469f1bc917dcff82c701aa0a247da34c8a7f232d1845261010435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1a370bfb4d9933d6bae8c65c4f41c1e4abf5e821ee472a5457c21b9328feae0c828801527f06ac7394c4cf7456cc60e9da7d5a945e992ed0f7e4979aaa2abc22efbc9fc4c8845261012435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0d6725204f08a7318747880cc67e561ad2332d45a7232bdb8f09f757bff6fe57828801527f2892e7493d2d56b90678fa529c8770a29a927aff7474b72a0e39c8db24d37a4f845261014435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f28cf0719114007b55c52a71f0b3a1c2fb60781bb251766e04c996e5337413c2e828801527f2701266463ff031d6927268f8798f7334b122f57da1123e5f601de175780ebd1845261016435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0d6325e0cd52ee00aa1ec11ee038ce8c8607867778b17d3a1a1d224237252134828801527f194efb11ab312fc89dfe9dc37d5a70c0de408de29891b140ddcb08b69b351ef4845261018435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1053ea377fb03c51dcc8d4223d7b797c38d9a012313f0bae5d063add31009e42828801527f0f3cc22d48f4c554f9f75856265543c2ed6f63397c26146786c0f6c72466e1f084526101a435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f01d1cf1bef078f0d9cd0bc4cf24034b389ea7e294856227dd5862e02987c0268828801527f07174e93bcca1980f209c404100f706747073da63dffbc6fc33498941b2ce14a84526101c435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611e265f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610c5f578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f09db1119a07ed79120dfaee28718f3e81cb504976ccb6c30214f83b598b9dc146101008501527f273a0e652d0ef4f8d3026d3db01c7406fb655ad69f7b127d963ec73b77fe51a46101208501527f2c806327f060d009b15feb81a511eb70e35e13051b32c9bb938f9d340e86bd746101408501527f087a66fab2082f62ab2bf0f36c4567488fc3db1487e262b4db9eecab52f57f616101608501527f0cb18dcdfd66083c47d4597cdd7f6d7d99115d97ff97b1cb2f62e266adee15d36101808501527f01ed124b6574d18301a94f683c4447538a091bd8cb8f22024c6452176a6420da6101a08501527f209fce5fd35738c10c170c45fb6633c86f6b4580dd69e2f38b76e1eaf4642fae6101c08501527f0e0bcb17a68fcb6ead6c0baccbfad4c507f6cd5a0a82ec8ed47248d92a1decca6101e08501527f14f0e892d8d3e91a1faa6eb7ecb7ec7a8843523c5a36b60c54cc6e2ab58a5b776102008501527f0721ad350acfac4e8ce26c326d142fda577e97b3815757197b0f6234b2d3d8aa6102208501526102408401526102608301527f23a1dbbdb025c9cef8e1a4eb04b0c3b2f6e063ca604cbeb0dc3ce59c800540ad6102808301527f1542860ca9c276bc73b143373cb7b868ef1e551f199efab4784b7d14cac9ebba6102a08301527f1882e97d1c74908c29fc04d1a81a6ef490b035d4d47acca24c4e6aaa7c2ccccc6102c08301527f0f0be2dcc6408c923f83a13319f860a94f111fe98799945021a3f49f5664741e6102e08301525192839161165984846116e4565b8336843760085afa15908115611671575b50610c5057005b600191505114158161166a565b34610236575f36600319011261023657807f9f0fdda496eec3ca9d71938acd7dd585ff09da6e811df8d652a80a12b623f7c760209252f35b9181601f840112156102365782359167ffffffffffffffff8311610236576020838186019501011161023657565b90601f8019910116810190811067ffffffffffffffff82111761024a57604052565b909160409282526020820152016060516080905f5b8181106117285750505090565b825184526020938401939092019160010161171b565b905f905b6002821061174f57505050565b6020806001928551815201930191019091611742565b80156117c9578060011c915f516020611e065f395f51905f52831015610c50576001806117a85f516020611e065f395f51905f5260038188818180090908611c26565b9316146117b157565b905f516020611e065f395f51905f5280910681030690565b505f905f90565b801580611901575b6118f5578060021c92825f516020611e065f395f51905f5285108015906118de575b610c505784815f516020611e065f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816118a89d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611c49565b809291600180829616146118ba575050565b5f516020611e065f395f51905f528093945080929550809106810306930681030690565b505f516020611e065f395f51905f528110156117fa565b50505f905f905f905f90565b5081156117d8565b905f516020611e065f395f51905f528210801590611993575b610c505781158061198b575b611985576119525f516020611e065f395f51905f5260038185818180090908611c26565b81810361196157505060011b90565b5f516020611e065f395f51905f52809106810306145f14610c5057600190811b1790565b50505f90565b50801561192e565b505f516020611e065f395f51905f52811015611922565b919093925f516020611e065f395f51905f528310801590611bc7575b8015611bb0575b8015611b99575b610c50578082868517171715611b8e57908291611af15f516020611e065f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611e065f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611acb81808b80098187800908611c26565b8408095f516020611e065f395f51905f52611ae582611d9d565b80091415958691611c49565b929080821480611b85575b15611b235750505050905f14611b1b5760ff60025b169060021b179190565b60ff5f611b11565b5f516020611e065f395f51905f52809106810306149182611b66575b505015610c505760019115611b5e5760ff60025b169060021b17179190565b60ff5f611b53565b5f516020611e065f395f51905f52919250819006810306145f80611b3f565b50838314611afc565b50505090505f905f90565b505f516020611e065f395f51905f528110156119d4565b505f516020611e065f395f51905f528210156119cd565b505f516020611e065f395f51905f528510156119c6565b9080601f8301121561023657604080519290611bfa90846116e4565b82906040810192831161023657905b828210611c165750505090565b8135815260209182019101611c09565b90611c3082611d9d565b915f516020611e065f395f51905f5283800903610c5057565b915f516020611e065f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611ca193969496611c9382808a8009818a800908611c26565b90611d91575b860809611c26565b925f516020611e065f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611e065f395f51905f5260a083015260208260c08160055afa91519115610c50575f516020611e065f395f51905f52826001920903610c50575f516020611e065f395f51905f52908209925f516020611e065f395f51905f528080808780090681030681878009081490811591611d72575b50610c5057565b90505f516020611e065f395f51905f528084860960020914155f611d6b565b81809106810306611c99565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611e065f395f51905f5260a083015260208260c08160055afa91519115610c505756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212205c189bb037085fab80b30358fc30c5d6b80d926424d594f1910a4b764b9d5fd164736f6c634300081c0033",
}

// ResultsVerifierGroth16ABI is the input ABI used to generate the binding from.
// Deprecated: Use ResultsVerifierGroth16MetaData.ABI instead.
var ResultsVerifierGroth16ABI = ResultsVerifierGroth16MetaData.ABI

// ResultsVerifierGroth16Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ResultsVerifierGroth16MetaData.Bin instead.
var ResultsVerifierGroth16Bin = ResultsVerifierGroth16MetaData.Bin

// DeployResultsVerifierGroth16 deploys a new Ethereum contract, binding an instance of ResultsVerifierGroth16 to it.
func DeployResultsVerifierGroth16(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ResultsVerifierGroth16, error) {
	parsed, err := ResultsVerifierGroth16MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ResultsVerifierGroth16Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ResultsVerifierGroth16{ResultsVerifierGroth16Caller: ResultsVerifierGroth16Caller{contract: contract}, ResultsVerifierGroth16Transactor: ResultsVerifierGroth16Transactor{contract: contract}, ResultsVerifierGroth16Filterer: ResultsVerifierGroth16Filterer{contract: contract}}, nil
}

// ResultsVerifierGroth16 is an auto generated Go binding around an Ethereum contract.
type ResultsVerifierGroth16 struct {
	ResultsVerifierGroth16Caller     // Read-only binding to the contract
	ResultsVerifierGroth16Transactor // Write-only binding to the contract
	ResultsVerifierGroth16Filterer   // Log filterer for contract events
}

// ResultsVerifierGroth16Caller is an auto generated read-only Go binding around an Ethereum contract.
type ResultsVerifierGroth16Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResultsVerifierGroth16Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ResultsVerifierGroth16Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResultsVerifierGroth16Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResultsVerifierGroth16Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResultsVerifierGroth16Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResultsVerifierGroth16Session struct {
	Contract     *ResultsVerifierGroth16 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ResultsVerifierGroth16CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResultsVerifierGroth16CallerSession struct {
	Contract *ResultsVerifierGroth16Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ResultsVerifierGroth16TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResultsVerifierGroth16TransactorSession struct {
	Contract     *ResultsVerifierGroth16Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ResultsVerifierGroth16Raw is an auto generated low-level Go binding around an Ethereum contract.
type ResultsVerifierGroth16Raw struct {
	Contract *ResultsVerifierGroth16 // Generic contract binding to access the raw methods on
}

// ResultsVerifierGroth16CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResultsVerifierGroth16CallerRaw struct {
	Contract *ResultsVerifierGroth16Caller // Generic read-only contract binding to access the raw methods on
}

// ResultsVerifierGroth16TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResultsVerifierGroth16TransactorRaw struct {
	Contract *ResultsVerifierGroth16Transactor // Generic write-only contract binding to access the raw methods on
}

// NewResultsVerifierGroth16 creates a new instance of ResultsVerifierGroth16, bound to a specific deployed contract.
func NewResultsVerifierGroth16(address common.Address, backend bind.ContractBackend) (*ResultsVerifierGroth16, error) {
	contract, err := bindResultsVerifierGroth16(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResultsVerifierGroth16{ResultsVerifierGroth16Caller: ResultsVerifierGroth16Caller{contract: contract}, ResultsVerifierGroth16Transactor: ResultsVerifierGroth16Transactor{contract: contract}, ResultsVerifierGroth16Filterer: ResultsVerifierGroth16Filterer{contract: contract}}, nil
}

// NewResultsVerifierGroth16Caller creates a new read-only instance of ResultsVerifierGroth16, bound to a specific deployed contract.
func NewResultsVerifierGroth16Caller(address common.Address, caller bind.ContractCaller) (*ResultsVerifierGroth16Caller, error) {
	contract, err := bindResultsVerifierGroth16(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResultsVerifierGroth16Caller{contract: contract}, nil
}

// NewResultsVerifierGroth16Transactor creates a new write-only instance of ResultsVerifierGroth16, bound to a specific deployed contract.
func NewResultsVerifierGroth16Transactor(address common.Address, transactor bind.ContractTransactor) (*ResultsVerifierGroth16Transactor, error) {
	contract, err := bindResultsVerifierGroth16(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResultsVerifierGroth16Transactor{contract: contract}, nil
}

// NewResultsVerifierGroth16Filterer creates a new log filterer instance of ResultsVerifierGroth16, bound to a specific deployed contract.
func NewResultsVerifierGroth16Filterer(address common.Address, filterer bind.ContractFilterer) (*ResultsVerifierGroth16Filterer, error) {
	contract, err := bindResultsVerifierGroth16(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResultsVerifierGroth16Filterer{contract: contract}, nil
}

// bindResultsVerifierGroth16 binds a generic wrapper to an already deployed contract.
func bindResultsVerifierGroth16(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResultsVerifierGroth16MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResultsVerifierGroth16.Contract.ResultsVerifierGroth16Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResultsVerifierGroth16.Contract.ResultsVerifierGroth16Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResultsVerifierGroth16.Contract.ResultsVerifierGroth16Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResultsVerifierGroth16.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResultsVerifierGroth16.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResultsVerifierGroth16.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "compressProof", proof, commitments, commitmentPok)

	outstruct := new(struct {
		Compressed              [4]*big.Int
		CompressedCommitments   [1]*big.Int
		CompressedCommitmentPok *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Compressed = *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)
	outstruct.CompressedCommitments = *abi.ConvertType(out[1], new([1]*big.Int)).(*[1]*big.Int)
	outstruct.CompressedCommitmentPok = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) CompressProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _ResultsVerifierGroth16.Contract.CompressProof(&_ResultsVerifierGroth16.CallOpts, proof, commitments, commitmentPok)
}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) CompressProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _ResultsVerifierGroth16.Contract.CompressProof(&_ResultsVerifierGroth16.CallOpts, proof, commitments, commitmentPok)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) ProvingKeyHash() ([32]byte, error) {
	return _ResultsVerifierGroth16.Contract.ProvingKeyHash(&_ResultsVerifierGroth16.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) ProvingKeyHash() ([32]byte, error) {
	return _ResultsVerifierGroth16.Contract.ProvingKeyHash(&_ResultsVerifierGroth16.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5d26278e.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, compressedCommitments, compressedCommitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5d26278e.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyCompressedProof(&_ResultsVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5d26278e.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyCompressedProof(&_ResultsVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "verifyProof", proof, commitments, commitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof(&_ResultsVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof(&_ResultsVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) VerifyProof0(opts *bind.CallOpts, _proof []byte, _input []byte) error {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "verifyProof0", _proof, _input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) VerifyProof0(_proof []byte, _input []byte) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof0(&_ResultsVerifierGroth16.CallOpts, _proof, _input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) VerifyProof0(_proof []byte, _input []byte) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof0(&_ResultsVerifierGroth16.CallOpts, _proof, _input)
}
