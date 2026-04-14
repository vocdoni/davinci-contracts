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

// StateTransitionVerifierGroth16MetaData contains all meta data concerning the StateTransitionVerifierGroth16 contract.
var StateTransitionVerifierGroth16MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"},{\"internalType\":\"uint256[8]\",\"name\":\"input\",\"type\":\"uint256[8]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[8]\",\"name\":\"input\",\"type\":\"uint256[8]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557611e02908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461157f575080639ed57f3214610cbb578063b1c3a00e14610baa578063b8e72af614610a335763e2dd9f4714610055575f80fd5b34610a30576101c0366003190112610a305736608411610a30573660a411610a3057366101c411610a3057604051602061008f81836115e5565b803683376040918251926100a381856115e5565b80368537610300938151926100b886856115e5565b853685376100c7608435611a09565b8684015282526100d860a435611a09565b905f516020611dad5f395f51905f5285516100f387826115e5565b6001815288610104818301601f198a01368237378551906101328a610124818a0151938b5194859384019687611607565b03601f1981018352826115e5565b5190200683528351865286840151878701527f05308f084f095d57a0168e0c1f5c4bec01a57686b54538012381077931b26122858701527f1c6e519eb5104bee7d271211178d24444eb23327b8e8173c02e524f79b51614660608701527f0660abefed7b9419420bf99fb6018d753b9cdc2a0f525369e49a98e3df19b8f960808701527f20ca63ebb727707cff29ce8dc1afdc6134ed6d33bddb0334c0396d9edf95dfe360a087015260c086015260e08501527f11830abf8fa5e7b38fb18592d32c14d9c672427e2a439c08d92b8e3b44adb85b6101008501527f0af036ef978692f1e9b18ce71d40de61da9c854e65b9a493020149f04b8384846101208501527f220dc6c610831ef767165d79e1553b2c6f3efd4e9ffc86c01cc31f44e068c13d6101408501527f02ec16e696f5680cc8f403b95f782505577413227e58cd5b74e1ee737fe6a333610160850152825185816101808760085afa90511615610a21578490876102a3600435611a09565b866102b5602497939735604435611a74565b9491929390956102c6606435611a09565b999098507f07e40b34d4db90e63f99e40a823155a693bd72dc739450717a1d179d3b2aaa6183519b8c937f1b12f72775b54389573dcee940a146f7b763e1633140ce83ecc2a3eb37fc0a3585527f2b978a2c1daa5cca056239eb20bad12da6e76b0fa110f53acd1a4a0187666504828601528051868601520151606084019081527f109cc9cc77001b165d5feb542b71935d5d064091c69a9335dd32adc934719818858560808160065afa957f1c4c179d363ca07e79df2eb1486cb5acd36cfcf9438b7efabc3a3b2db6ba9536818701527f08233c066ea88b37dd4e8d1168546dc3304b6b5594f044c1040a1882b009b96c8352600160c4356080880198818a525f516020611dad5f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f052418bf3f53e095b83d6571f4045a820406cb6c0252a83a3838099a6460583a828801527f0b6be2dd2d63f24c145fa877ccc9f590782e9e08f167969de390eee8ec8caf1e845260e435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0d388e5f5366f4946e093215d2b6ebdf9ada366264e9b321c1710122577669b6828801527f129166154275cd693cae4018f3c1da323bf2479052a5afd8211937575c208544845261010435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f07ef1f411756616515268dff6f656185038be88c4c69a8bd86b2ea84b51e2d57828801527f0bad4e9d07dd7e0ba82d0e8607df6d3e428b3a5a5cc8c75880edce364da7ebfd845261012435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f044d5efaabceedba95fe9d4932e9469d3f448120ab82006078e3a467b216a700828801527f05ab3170a8113a47196f8cf2f505f3c06e0c082640facf6b3ce2e835f0ddb3e2845261014435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f14d6c9671b244e632cecf5022d3331224f6d957e9cd8dab371fecbb6e241cc23828801527f0861501c80d06be6b6ce1a0ddfad65bf3adcb5509ef8910053c6c357498b3db6845261016435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1bf743cfb1cce06d6deff8a916ff8e16707d56528d21aacedd5dd6938985d37c828801527f02fa7fccd7d6852bae03096c65e755439dbe390d8728d670bbf2a52e2a25da5e845261018435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2eac335b52fc8c6c34bbc322a29aba897113f326ce36d8d4467813a5a1386871828801527f07efa7eb29d1bbd343fef155f1ff221c91b5a1ee26482dd416d237b93be43b7384526101a435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611dad5f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610a12578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f1aa8830546092d5005f8ac38154e887c6c627b92c64dddf936007bf56b0ac56c6101008501527f2e09f094e1475e159b03aa882025585daf025347e9ed5ac276d968d75d3a5e2f6101208501527f1bde56adef8bfa4c181bbc097867454db41c59284a3609e5cde057fb4db1ab376101408501527f07678fcb2b1040e20bd7b604ab90d1c874bea166781adf35716eaf839e6989ba6101608501527f2331fda05e9c971e2d05701cb406e63b63dc1c1fe4e0d404cfd9a42f19f27d576101808501527f0f5f095e950964b28e57aded55ae716369c8e213db3d783c0d4b1594bb4e53aa6101a08501527f048488267ad5f22c59da274e8e2ae1dbbcc6c04ef8c6ff4d13581df22d459eb66101c08501527f155407bd8edbbb29aebade4884bd1b8c66c49f8806d2307e5b83d831004e47086101e08501527f02f2a48c98fa9cde4a7025079f39ccf8a6bd894d8ec4d2bf1635097a44fc60666102008501527f02ca7afe9085ac23766d19cbde5ea581f61dac9c4d6160f6bea52f102a015cf36102208501526102408401526102608301527f072e94a35c73cdbee40f7ec7093ddeb482f98e5d164e03020a6354e65e54a1c66102808301527f2a6ce154fc4ffc138ad4551da548e7743bd0ec1edb75f4b4528a21e39d990d886102a08301527f07c0fa19c07ac630f622529a5cb5f58eb546394b2e1dc8188efbf91a96be621e6102c08301527f2e53374be87b4b10b181c153b594d45a650b82e7a045944a74da05e741a133e16102e0830152519283916109dd84846115e5565b8336843760085afa15908115610a05575b506109f65780f35b631ff3747d60e21b8152600490fd5b600191505114155f6109ee565b63a54f8e2760e01b8f5260048ffd5b6351d49ff760e11b8752600487fd5b80fd5b5034610ba6576040366003190112610ba65760043567ffffffffffffffff8111610ba657610a659036906004016115b7565b60243567ffffffffffffffff8111610ba657610a859036906004016115b7565b90926101009384604051610a9982826115e5565b369037604093848051610aac82826115e5565b369037848051610abc82826115e5565b36903782019461018083870312610ba657610af1610ada8785611964565b96610140610aea828588016119ad565b95016119ad565b93818651610aff82826115e5565b36903782019082820312610ba657610b1691611964565b90610b1f6119f5565b50610b286119f5565b50303b15610ba657610b7692610b60610b6b92610b55875198634f6abf9960e11b8a5260048a0190611641565b610104880190611668565b610144860190611668565b610184840190611641565b5f8261028481305afa908115610b9d5750610b8f575080f35b610b9b91505f906115e5565b005b513d5f823e3d90fd5b5f80fd5b34610ba657610180366003190112610ba6573661010411610ba6573661014411610ba6573661018411610ba657604051610be56080826115e5565b6080368237604051906020610bfa81846115e5565b80368437610c0c60243560043561168f565b8252610c2260843560a435604435606435611730565b828401526040830152610c3960e43560c43561168f565b6060830152610c4e610124356101043561168f565b8352610c60610164356101443561168f565b9060405192835f905b60048210610ca5575050506080830193905f945b60018610610c905760c0858560a0820152f35b81806001928551815201930195019491610c7d565b8251815291830191600191909101908301610c69565b34610ba657610280366003190112610ba6573661010411610ba6573661014411610ba6573661018411610ba6573661028411610ba6576040516020610d0081836115e5565b803683376040915f516020611dad5f395f51905f528351610d2185826115e5565b60018152836101c4818301601f198801368237378451610d508161012487820194610124356101043587611607565b51902006815282518361010482377f05308f084f095d57a0168e0c1f5c4bec01a57686b54538012381077931b26122848201527f1c6e519eb5104bee7d271211178d24444eb23327b8e8173c02e524f79b51614660608201527f0660abefed7b9419420bf99fb6018d753b9cdc2a0f525369e49a98e3df19b8f960808201527f20ca63ebb727707cff29ce8dc1afdc6134ed6d33bddb0334c0396d9edf95dfe360a08201528361014460c08301377f11830abf8fa5e7b38fb18592d32c14d9c672427e2a439c08d92b8e3b44adb85b6101008201527f0af036ef978692f1e9b18ce71d40de61da9c854e65b9a493020149f04b8384846101208201527f220dc6c610831ef767165d79e1553b2c6f3efd4e9ffc86c01cc31f44e068c13d6101408201527f02ec16e696f5680cc8f403b95f782505577413227e58cd5b74e1ee737fe6a33361016082015282816101808160085afa9051161561157057825192610eba6040856115e5565b61010483855b610144831061154c575050508051918183017f1b12f72775b54389573dcee940a146f7b763e1633140ce83ecc2a3eb37fc0a3584525f516020611dad5f395f51905f528386808701987f2b978a2c1daa5cca056239eb20bad12da6e76b0fa110f53acd1a4a01876665048a528051855201519260608701938452818760808160065afa947f1c4c179d363ca07e79df2eb1486cb5acd36cfcf9438b7efabc3a3b2db6ba953682527f08233c066ea88b37dd4e8d1168546dc3304b6b5594f044c1040a1882b009b96c85527f07e40b34d4db90e63f99e40a823155a693bd72dc739450717a1d179d3b2aaa6160016101843560808b0198818a5287878760608160075afa92101616858b60808160065afa16167f052418bf3f53e095b83d6571f4045a820406cb6c0252a83a3838099a6460583a84527f0b6be2dd2d63f24c145fa877ccc9f590782e9e08f167969de390eee8ec8caf1e87526101a4359081895286868660608160075afa92101616848a60808160065afa167f0d388e5f5366f4946e093215d2b6ebdf9ada366264e9b321c1710122577669b684527f129166154275cd693cae4018f3c1da323bf2479052a5afd8211937575c20854487526101c4359081895286868660608160075afa92101616848a60808160065afa167f07ef1f411756616515268dff6f656185038be88c4c69a8bd86b2ea84b51e2d5784527f0bad4e9d07dd7e0ba82d0e8607df6d3e428b3a5a5cc8c75880edce364da7ebfd87526101e4359081895286868660608160075afa92101616848a60808160065afa167f044d5efaabceedba95fe9d4932e9469d3f448120ab82006078e3a467b216a70084527f05ab3170a8113a47196f8cf2f505f3c06e0c082640facf6b3ce2e835f0ddb3e28752610204359081895286868660608160075afa92101616848a60808160065afa167f14d6c9671b244e632cecf5022d3331224f6d957e9cd8dab371fecbb6e241cc2384527f0861501c80d06be6b6ce1a0ddfad65bf3adcb5509ef8910053c6c357498b3db68752610224359081895286868660608160075afa92101616848a60808160065afa167f1bf743cfb1cce06d6deff8a916ff8e16707d56528d21aacedd5dd6938985d37c84527f02fa7fccd7d6852bae03096c65e755439dbe390d8728d670bbf2a52e2a25da5e8752610244359081895286868660608160075afa92101616848a60808160065afa167f2eac335b52fc8c6c34bbc322a29aba897113f326ce36d8d4467813a5a138687184527f07efa7eb29d1bbd343fef155f1ff221c91b5a1ee26482dd416d237b93be43b738752610264359081895286868660608160075afa92101616848a60808160065afa16957f109cc9cc77001b165d5feb542b71935d5d064091c69a9335dd32adc9347198188452525180955260608160075afa92101616818360808160065afa1691519351911561153d575192610100600485377f1aa8830546092d5005f8ac38154e887c6c627b92c64dddf936007bf56b0ac56c6101008501527f2e09f094e1475e159b03aa882025585daf025347e9ed5ac276d968d75d3a5e2f6101208501527f1bde56adef8bfa4c181bbc097867454db41c59284a3609e5cde057fb4db1ab376101408501527f07678fcb2b1040e20bd7b604ab90d1c874bea166781adf35716eaf839e6989ba6101608501527f2331fda05e9c971e2d05701cb406e63b63dc1c1fe4e0d404cfd9a42f19f27d576101808501527f0f5f095e950964b28e57aded55ae716369c8e213db3d783c0d4b1594bb4e53aa6101a08501527f048488267ad5f22c59da274e8e2ae1dbbcc6c04ef8c6ff4d13581df22d459eb66101c08501527f155407bd8edbbb29aebade4884bd1b8c66c49f8806d2307e5b83d831004e47086101e08501527f02f2a48c98fa9cde4a7025079f39ccf8a6bd894d8ec4d2bf1635097a44fc60666102008501527f02ca7afe9085ac23766d19cbde5ea581f61dac9c4d6160f6bea52f102a015cf36102208501526102408401526102608301527f072e94a35c73cdbee40f7ec7093ddeb482f98e5d164e03020a6354e65e54a1c66102808301527f2a6ce154fc4ffc138ad4551da548e7743bd0ec1edb75f4b4528a21e39d990d886102a08301527f07c0fa19c07ac630f622529a5cb5f58eb546394b2e1dc8188efbf91a96be621e6102c08301527f2e53374be87b4b10b181c153b594d45a650b82e7a045944a74da05e741a133e16102e0830152816103008160085afa9051161561152e57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191859101610ec0565b634e487b7160e01b5f52604160045260245ffd5b6351d49ff760e11b5f5260045ffd5b34610ba6575f366003190112610ba657807fba6e46003838f8d4f21e45a0beb72501331b60ebdb6cca104eb75e574d469ad760209252f35b9181601f84011215610ba65782359167ffffffffffffffff8311610ba65760208381860195010111610ba657565b90601f8019910116810190811067ffffffffffffffff82111761155c57604052565b9091604092825260208201520160208251919201905f5b81811061162b5750505090565b825184526020938401939092019160010161161e565b905f905b6008821061165257505050565b6020806001928551815201930191019091611645565b905f905b6002821061167957505050565b602080600192855181520193019101909161166c565b905f516020611d8d5f395f51905f528210801590611719575b61152e57811580611711575b61170b576116d85f516020611d8d5f395f51905f5260038185818180090908611bad565b8181036116e757505060011b90565b5f516020611d8d5f395f51905f52809106810306145f1461152e57600190811b1790565b50505f90565b5080156116b4565b505f516020611d8d5f395f51905f528110156116a8565b919093925f516020611d8d5f395f51905f52831080159061194d575b8015611936575b801561191f575b61152e578082868517171715611914579082916118775f516020611d8d5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611d8d5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161185181808b80098187800908611bad565b8408095f516020611d8d5f395f51905f5261186b82611d24565b80091415958691611bd0565b92908082148061190b575b156118a95750505050905f146118a15760ff60025b169060021b179190565b60ff5f611897565b5f516020611d8d5f395f51905f528091068103061491826118ec575b50501561152e57600191156118e45760ff60025b169060021b17179190565b60ff5f6118d9565b5f516020611d8d5f395f51905f52919250819006810306145f806118c5565b50838314611882565b50505090505f905f90565b505f516020611d8d5f395f51905f5281101561175a565b505f516020611d8d5f395f51905f52821015611753565b505f516020611d8d5f395f51905f5285101561174c565b9080601f83011215610ba65760405191611980610100846115e5565b82906101008101928311610ba657905b82821061199d5750505090565b8135815260209182019101611990565b9080601f83011215610ba6576040805192906119c990846115e5565b829060408101928311610ba657905b8282106119e55750505090565b81358152602091820191016119d8565b60405190611a046020836115e5565b5f8252565b8015611a6d578060011c915f516020611d8d5f395f51905f5283101561152e57600180611a4c5f516020611d8d5f395f51905f5260038188818180090908611bad565b931614611a5557565b905f516020611d8d5f395f51905f5280910681030690565b505f905f90565b801580611ba5575b611b99578060021c92825f516020611d8d5f395f51905f528510801590611b82575b61152e5784815f516020611d8d5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4481611b4c9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611bd0565b80929160018082961614611b5e575050565b5f516020611d8d5f395f51905f528093945080929550809106810306930681030690565b505f516020611d8d5f395f51905f52811015611a9e565b50505f905f905f905f90565b508115611a7c565b90611bb782611d24565b915f516020611d8d5f395f51905f528380090361152e57565b915f516020611d8d5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c2893969496611c1a82808a8009818a800908611bad565b90611d18575b860809611bad565b925f516020611d8d5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611d8d5f395f51905f5260a083015260208260c08160055afa9151911561152e575f516020611d8d5f395f51905f5282600192090361152e575f516020611d8d5f395f51905f52908209925f516020611d8d5f395f51905f528080808780090681030681878009081490811591611cf9575b5061152e57565b90505f516020611d8d5f395f51905f528084860960020914155f611cf2565b81809106810306611c20565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611d8d5f395f51905f5260a083015260208260c08160055afa9151911561152e5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212202fb3e97aaa9830cd444c463cab4e829033bcf63a4d01399f3405ad3167d572b864736f6c634300081c0033",
}

// StateTransitionVerifierGroth16ABI is the input ABI used to generate the binding from.
// Deprecated: Use StateTransitionVerifierGroth16MetaData.ABI instead.
var StateTransitionVerifierGroth16ABI = StateTransitionVerifierGroth16MetaData.ABI

// StateTransitionVerifierGroth16Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateTransitionVerifierGroth16MetaData.Bin instead.
var StateTransitionVerifierGroth16Bin = StateTransitionVerifierGroth16MetaData.Bin

// DeployStateTransitionVerifierGroth16 deploys a new Ethereum contract, binding an instance of StateTransitionVerifierGroth16 to it.
func DeployStateTransitionVerifierGroth16(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StateTransitionVerifierGroth16, error) {
	parsed, err := StateTransitionVerifierGroth16MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateTransitionVerifierGroth16Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateTransitionVerifierGroth16{StateTransitionVerifierGroth16Caller: StateTransitionVerifierGroth16Caller{contract: contract}, StateTransitionVerifierGroth16Transactor: StateTransitionVerifierGroth16Transactor{contract: contract}, StateTransitionVerifierGroth16Filterer: StateTransitionVerifierGroth16Filterer{contract: contract}}, nil
}

// StateTransitionVerifierGroth16 is an auto generated Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16 struct {
	StateTransitionVerifierGroth16Caller     // Read-only binding to the contract
	StateTransitionVerifierGroth16Transactor // Write-only binding to the contract
	StateTransitionVerifierGroth16Filterer   // Log filterer for contract events
}

// StateTransitionVerifierGroth16Caller is an auto generated read-only Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransitionVerifierGroth16Transactor is an auto generated write-only Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransitionVerifierGroth16Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateTransitionVerifierGroth16Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransitionVerifierGroth16Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateTransitionVerifierGroth16Session struct {
	Contract     *StateTransitionVerifierGroth16 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// StateTransitionVerifierGroth16CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateTransitionVerifierGroth16CallerSession struct {
	Contract *StateTransitionVerifierGroth16Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// StateTransitionVerifierGroth16TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateTransitionVerifierGroth16TransactorSession struct {
	Contract     *StateTransitionVerifierGroth16Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// StateTransitionVerifierGroth16Raw is an auto generated low-level Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16Raw struct {
	Contract *StateTransitionVerifierGroth16 // Generic contract binding to access the raw methods on
}

// StateTransitionVerifierGroth16CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16CallerRaw struct {
	Contract *StateTransitionVerifierGroth16Caller // Generic read-only contract binding to access the raw methods on
}

// StateTransitionVerifierGroth16TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16TransactorRaw struct {
	Contract *StateTransitionVerifierGroth16Transactor // Generic write-only contract binding to access the raw methods on
}

// NewStateTransitionVerifierGroth16 creates a new instance of StateTransitionVerifierGroth16, bound to a specific deployed contract.
func NewStateTransitionVerifierGroth16(address common.Address, backend bind.ContractBackend) (*StateTransitionVerifierGroth16, error) {
	contract, err := bindStateTransitionVerifierGroth16(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateTransitionVerifierGroth16{StateTransitionVerifierGroth16Caller: StateTransitionVerifierGroth16Caller{contract: contract}, StateTransitionVerifierGroth16Transactor: StateTransitionVerifierGroth16Transactor{contract: contract}, StateTransitionVerifierGroth16Filterer: StateTransitionVerifierGroth16Filterer{contract: contract}}, nil
}

// NewStateTransitionVerifierGroth16Caller creates a new read-only instance of StateTransitionVerifierGroth16, bound to a specific deployed contract.
func NewStateTransitionVerifierGroth16Caller(address common.Address, caller bind.ContractCaller) (*StateTransitionVerifierGroth16Caller, error) {
	contract, err := bindStateTransitionVerifierGroth16(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateTransitionVerifierGroth16Caller{contract: contract}, nil
}

// NewStateTransitionVerifierGroth16Transactor creates a new write-only instance of StateTransitionVerifierGroth16, bound to a specific deployed contract.
func NewStateTransitionVerifierGroth16Transactor(address common.Address, transactor bind.ContractTransactor) (*StateTransitionVerifierGroth16Transactor, error) {
	contract, err := bindStateTransitionVerifierGroth16(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateTransitionVerifierGroth16Transactor{contract: contract}, nil
}

// NewStateTransitionVerifierGroth16Filterer creates a new log filterer instance of StateTransitionVerifierGroth16, bound to a specific deployed contract.
func NewStateTransitionVerifierGroth16Filterer(address common.Address, filterer bind.ContractFilterer) (*StateTransitionVerifierGroth16Filterer, error) {
	contract, err := bindStateTransitionVerifierGroth16(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateTransitionVerifierGroth16Filterer{contract: contract}, nil
}

// bindStateTransitionVerifierGroth16 binds a generic wrapper to an already deployed contract.
func bindStateTransitionVerifierGroth16(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StateTransitionVerifierGroth16MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateTransitionVerifierGroth16.Contract.StateTransitionVerifierGroth16Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateTransitionVerifierGroth16.Contract.StateTransitionVerifierGroth16Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateTransitionVerifierGroth16.Contract.StateTransitionVerifierGroth16Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateTransitionVerifierGroth16.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateTransitionVerifierGroth16.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateTransitionVerifierGroth16.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "compressProof", proof, commitments, commitmentPok)

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
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) CompressProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _StateTransitionVerifierGroth16.Contract.CompressProof(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok)
}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) CompressProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _StateTransitionVerifierGroth16.Contract.CompressProof(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) ProvingKeyHash() ([32]byte, error) {
	return _StateTransitionVerifierGroth16.Contract.ProvingKeyHash(&_StateTransitionVerifierGroth16.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) ProvingKeyHash() ([32]byte, error) {
	return _StateTransitionVerifierGroth16.Contract.ProvingKeyHash(&_StateTransitionVerifierGroth16.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe2dd9f47.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, compressedCommitments, compressedCommitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe2dd9f47.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [8]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyCompressedProof(&_StateTransitionVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe2dd9f47.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [8]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyCompressedProof(&_StateTransitionVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x9ed57f32.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyProof", proof, commitments, commitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x9ed57f32.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [8]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x9ed57f32.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [8]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyProof0(opts *bind.CallOpts, _proof []byte, _input []byte) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyProof0", _proof, _input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyProof0(_proof []byte, _input []byte) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof0(&_StateTransitionVerifierGroth16.CallOpts, _proof, _input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyProof0(_proof []byte, _input []byte) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof0(&_StateTransitionVerifierGroth16.CallOpts, _proof, _input)
}
