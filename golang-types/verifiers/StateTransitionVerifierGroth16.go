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
	Bin: "0x60808060405234601557611dfe908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461157b575080639ed57f3214610cb9578063b1c3a00e14610ba8578063b8e72af614610a315763e2dd9f4714610055575f80fd5b34610a2e576101c0366003190112610a2e5736608411610a2e573660a411610a2e57366101c411610a2e57604051602061008f81836115e1565b803683376040918251926100a381856115e1565b80368537610300938151926100b886856115e1565b853685376100c7608435611a05565b8684015282526100d860a435611a05565b905f516020611da95f395f51905f5285516100f387826115e1565b6001815288610104818301601f198a01368237378551906101328a610124818a0151938b5194859384019687611603565b03601f1981018352826115e1565b5190200683528351865286840151878701527f129ea85df19445ba9e718bfc798832ec19fc54d6871ee992ee3644efe60361c0858701527f2eeee6c432f64a24931ad776307a61b99bab74b41c1c621b7c8df2ca0359c7bd60608701527f2bc629c54341a945ca3ec97522d445b5c7618c17ef06e822c9e007c3b24895e660808701527ec1d36a45205652c210cec488e2c0225d692544bfe895e4dea0041c71d7a18960a087015260c086015260e08501527f1a4ef41410348a77b7640f4c5e491410026ce6f14096d871a8331e57cc696f206101008501527f0cb03eb14330efcdc5a721900747b87232de16ec461bc8e61239939e39deed386101208501527f2ce3280e12a6069c70aa22955a5837eb207c418c279a9e7cad7cd0f4e6f8cfaa6101408501527f0155da331b524e59bbc31093bbb893dc7bc49566c4eee7f4f65e4a9e375993e6610160850152825185816101808760085afa90511615610a1f578490876102a2600435611a05565b866102b4602497939735604435611a70565b9491929390956102c5606435611a05565b999098507f0c4a62450775f19c01e7429ca32d5e7b05d57a81347b7aa7dcb66e0f1e433b3a83519b8c937f1d51a2cebd061916b84b001d2be570f4f4ddf347ff63449aa74639384996f58d85527f0220b2619e95a34954edaf5bb2cef83a8c30927a8e166c282970534793fdf73e828601528051868601520151606084019081527f125a5906e995d8b74e6d1c37a3a1721b58685d2e8ee2658b3f03e4f29cd4d54b858560808160065afa957f0629cf10d24fce1656b615f53d79eddee0705e223157bae926949e2c46c2c74f818701527f0aec8680b8b2d499ffd34f18aa16e7ee1dcf578191ec4e0e743adcb8da5780c48352600160c4356080880198818a525f516020611da95f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f11ba9a7dd0fc2b6460f65e25fdae8bf29c6ea110ab41cbca6f183a084bcddea7828801527f1a0e182a16f03235233a0bdf69e6cc62d0b64bef1f7ea5e06dc53bfd76b2cb60845260e435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f19b7a440300b6fad49c2932a7de19ac45bd0c0e54fb5fc6f84a8852fb3d1bf3c828801527f296c7c6496a65377db64b25396844a9f284eff7de4a52be20b9d1313e73cf3a7845261010435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1fc55b137d7f9f532b18026f3f36d7395fb12d1c30584133e310cef908696320828801527f192322785e18de98d9052e4ed2041859b7cc4c1daed7bc35b720e76006c05de8845261012435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f23b303bbc6d5d587270831635ae99a798b2daf7057fda73ff66ddda09e16ba83828801527eb1ff9fe1618b0c68cfa03c3484e6ad358feefb12e2e7b59f77aa2233c2fd8e845261014435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f07d064634eef00186a74b29eb92b1601b7f6bcfece8a7e6557b30c79e73a1d1c828801527f0da6793eb2498691be04464e0574016a52b679673abba4dcea4fb2e2c9a109e2845261016435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f139848a2a7c875a5b8834f97c536129ee22574aafb7ee24659feb3257c596dc6828801527f24d9a60d2c279ea55bfd179537a212b847b1076a95408bbf7c06e9f6cfa9a428845261018435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f08aa82d8ff4e425230ff25ee9c6633af86d3e3078c2ede53f4accb464baba410828801527f0e74ab57a507dc4717a3bddefeef61dbd0fdc664ef8b2480027f34f5b7eeac7784526101a435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611da95f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610a10578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f2405ca43b82e73302dc922162c5d06891235ae11a88461d7b9835b01c5c06aa56101008501527f2fbb411ac5f9e75b110fc9f8fd4ba389580a9a253c8c59f83a9bd7684c5b522e6101208501527f15e078edd0756ce3f96c7bca9d28bd2a8ad0e17ceb3287904f572cb88839ebd46101408501527f1fbf6495e01dfac5f4241e3fe068ec38b2b88653856939166b3a4e8de54cd4866101608501527f1f302f91957fb0acfbfe150cc040a0c9b6375a7d90e632b86bec96c7631492406101808501527f17c613ebc236baf72e453f81c76879a89c879e9fc9d93a94fecb4d68ea67c2626101a08501527f0fac29aa7d7e813379ad38c124c57dcedc88ada161c9e391c110a26d07412c976101c08501527f099a1db894a80ac96b7242026e0a80b7c40bad4064ea763948c6ce2a3b7b9e976101e08501527f1d0ba55682d8b0fa1ba6eed361d49bbabd28c6c2a7a354d1e08f522d326fb13b6102008501527f2225659acd03b38ab4167a0d653e817b00b3a420bc5393817881df3ff89f3f916102208501526102408401526102608301527f29609df80be2dab2fe3a262b1a0fd377107200c9a3cdc3eaf9fec484dd2f30f76102808301527f2b68ddc9fd8defbe00f4b8ca76434f8ab731da07ad29e62c6f89c4e1d09589e06102a08301527f06ff4f7edb850015ebf5c4a86f668dd748c92731167093a2cd89ce2336fde72e6102c08301527f1057eecfe6cd48bbd604b90ab35daf5bb7c66c3c5432d90de97f16f922799ecd6102e0830152519283916109db84846115e1565b8336843760085afa15908115610a03575b506109f45780f35b631ff3747d60e21b8152600490fd5b600191505114155f6109ec565b63a54f8e2760e01b8f5260048ffd5b6351d49ff760e11b8752600487fd5b80fd5b5034610ba4576040366003190112610ba45760043567ffffffffffffffff8111610ba457610a639036906004016115b3565b60243567ffffffffffffffff8111610ba457610a839036906004016115b3565b90926101009384604051610a9782826115e1565b369037604093848051610aaa82826115e1565b369037848051610aba82826115e1565b36903782019461018083870312610ba457610aef610ad88785611960565b96610140610ae8828588016119a9565b95016119a9565b93818651610afd82826115e1565b36903782019082820312610ba457610b1491611960565b90610b1d6119f1565b50610b266119f1565b50303b15610ba457610b7492610b5e610b6992610b53875198634f6abf9960e11b8a5260048a019061163d565b610104880190611664565b610144860190611664565b61018484019061163d565b5f8261028481305afa908115610b9b5750610b8d575080f35b610b9991505f906115e1565b005b513d5f823e3d90fd5b5f80fd5b34610ba457610180366003190112610ba4573661010411610ba4573661014411610ba4573661018411610ba457604051610be36080826115e1565b6080368237604051906020610bf881846115e1565b80368437610c0a60243560043561168b565b8252610c2060843560a43560443560643561172c565b828401526040830152610c3760e43560c43561168b565b6060830152610c4c610124356101043561168b565b8352610c5e610164356101443561168b565b9060405192835f905b60048210610ca3575050506080830193905f945b60018610610c8e5760c0858560a0820152f35b81806001928551815201930195019491610c7b565b8251815291830191600191909101908301610c67565b34610ba457610280366003190112610ba4573661010411610ba4573661014411610ba4573661018411610ba4573661028411610ba4576040516020610cfe81836115e1565b803683376040915f516020611da95f395f51905f528351610d1f85826115e1565b60018152836101c4818301601f198801368237378451610d4e8161012487820194610124356101043587611603565b51902006815282518361010482377f129ea85df19445ba9e718bfc798832ec19fc54d6871ee992ee3644efe60361c0848201527f2eeee6c432f64a24931ad776307a61b99bab74b41c1c621b7c8df2ca0359c7bd60608201527f2bc629c54341a945ca3ec97522d445b5c7618c17ef06e822c9e007c3b24895e660808201527ec1d36a45205652c210cec488e2c0225d692544bfe895e4dea0041c71d7a18960a08201528361014460c08301377f1a4ef41410348a77b7640f4c5e491410026ce6f14096d871a8331e57cc696f206101008201527f0cb03eb14330efcdc5a721900747b87232de16ec461bc8e61239939e39deed386101208201527f2ce3280e12a6069c70aa22955a5837eb207c418c279a9e7cad7cd0f4e6f8cfaa6101408201527f0155da331b524e59bbc31093bbb893dc7bc49566c4eee7f4f65e4a9e375993e661016082015282816101808160085afa9051161561156c57825192610eb76040856115e1565b61010483855b6101448310611548575050508051918183017f1d51a2cebd061916b84b001d2be570f4f4ddf347ff63449aa74639384996f58d84525f516020611da95f395f51905f528386808701987f0220b2619e95a34954edaf5bb2cef83a8c30927a8e166c282970534793fdf73e8a528051855201519260608701938452818760808160065afa947f0629cf10d24fce1656b615f53d79eddee0705e223157bae926949e2c46c2c74f82527f0aec8680b8b2d499ffd34f18aa16e7ee1dcf578191ec4e0e743adcb8da5780c485527f0c4a62450775f19c01e7429ca32d5e7b05d57a81347b7aa7dcb66e0f1e433b3a60016101843560808b0198818a5287878760608160075afa92101616858b60808160065afa16167f11ba9a7dd0fc2b6460f65e25fdae8bf29c6ea110ab41cbca6f183a084bcddea784527f1a0e182a16f03235233a0bdf69e6cc62d0b64bef1f7ea5e06dc53bfd76b2cb6087526101a4359081895286868660608160075afa92101616848a60808160065afa167f19b7a440300b6fad49c2932a7de19ac45bd0c0e54fb5fc6f84a8852fb3d1bf3c84527f296c7c6496a65377db64b25396844a9f284eff7de4a52be20b9d1313e73cf3a787526101c4359081895286868660608160075afa92101616848a60808160065afa167f1fc55b137d7f9f532b18026f3f36d7395fb12d1c30584133e310cef90869632084527f192322785e18de98d9052e4ed2041859b7cc4c1daed7bc35b720e76006c05de887526101e4359081895286868660608160075afa92101616848a60808160065afa167f23b303bbc6d5d587270831635ae99a798b2daf7057fda73ff66ddda09e16ba8384527eb1ff9fe1618b0c68cfa03c3484e6ad358feefb12e2e7b59f77aa2233c2fd8e8752610204359081895286868660608160075afa92101616848a60808160065afa167f07d064634eef00186a74b29eb92b1601b7f6bcfece8a7e6557b30c79e73a1d1c84527f0da6793eb2498691be04464e0574016a52b679673abba4dcea4fb2e2c9a109e28752610224359081895286868660608160075afa92101616848a60808160065afa167f139848a2a7c875a5b8834f97c536129ee22574aafb7ee24659feb3257c596dc684527f24d9a60d2c279ea55bfd179537a212b847b1076a95408bbf7c06e9f6cfa9a4288752610244359081895286868660608160075afa92101616848a60808160065afa167f08aa82d8ff4e425230ff25ee9c6633af86d3e3078c2ede53f4accb464baba41084527f0e74ab57a507dc4717a3bddefeef61dbd0fdc664ef8b2480027f34f5b7eeac778752610264359081895286868660608160075afa92101616848a60808160065afa16957f125a5906e995d8b74e6d1c37a3a1721b58685d2e8ee2658b3f03e4f29cd4d54b8452525180955260608160075afa92101616818360808160065afa16915193519115611539575192610100600485377f2405ca43b82e73302dc922162c5d06891235ae11a88461d7b9835b01c5c06aa56101008501527f2fbb411ac5f9e75b110fc9f8fd4ba389580a9a253c8c59f83a9bd7684c5b522e6101208501527f15e078edd0756ce3f96c7bca9d28bd2a8ad0e17ceb3287904f572cb88839ebd46101408501527f1fbf6495e01dfac5f4241e3fe068ec38b2b88653856939166b3a4e8de54cd4866101608501527f1f302f91957fb0acfbfe150cc040a0c9b6375a7d90e632b86bec96c7631492406101808501527f17c613ebc236baf72e453f81c76879a89c879e9fc9d93a94fecb4d68ea67c2626101a08501527f0fac29aa7d7e813379ad38c124c57dcedc88ada161c9e391c110a26d07412c976101c08501527f099a1db894a80ac96b7242026e0a80b7c40bad4064ea763948c6ce2a3b7b9e976101e08501527f1d0ba55682d8b0fa1ba6eed361d49bbabd28c6c2a7a354d1e08f522d326fb13b6102008501527f2225659acd03b38ab4167a0d653e817b00b3a420bc5393817881df3ff89f3f916102208501526102408401526102608301527f29609df80be2dab2fe3a262b1a0fd377107200c9a3cdc3eaf9fec484dd2f30f76102808301527f2b68ddc9fd8defbe00f4b8ca76434f8ab731da07ad29e62c6f89c4e1d09589e06102a08301527f06ff4f7edb850015ebf5c4a86f668dd748c92731167093a2cd89ce2336fde72e6102c08301527f1057eecfe6cd48bbd604b90ab35daf5bb7c66c3c5432d90de97f16f922799ecd6102e0830152816103008160085afa9051161561152a57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191859101610ebd565b634e487b7160e01b5f52604160045260245ffd5b6351d49ff760e11b5f5260045ffd5b34610ba4575f366003190112610ba457807f3a3bbd3ab2325ee3d971b25f5c29bb3bed1c9d44e311cdd6d5b4600a392d9ecb60209252f35b9181601f84011215610ba45782359167ffffffffffffffff8311610ba45760208381860195010111610ba457565b90601f8019910116810190811067ffffffffffffffff82111761155857604052565b9091604092825260208201520160208251919201905f5b8181106116275750505090565b825184526020938401939092019160010161161a565b905f905b6008821061164e57505050565b6020806001928551815201930191019091611641565b905f905b6002821061167557505050565b6020806001928551815201930191019091611668565b905f516020611d895f395f51905f528210801590611715575b61152a5781158061170d575b611707576116d45f516020611d895f395f51905f5260038185818180090908611ba9565b8181036116e357505060011b90565b5f516020611d895f395f51905f52809106810306145f1461152a57600190811b1790565b50505f90565b5080156116b0565b505f516020611d895f395f51905f528110156116a4565b919093925f516020611d895f395f51905f528310801590611949575b8015611932575b801561191b575b61152a578082868517171715611910579082916118735f516020611d895f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611d895f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161184d81808b80098187800908611ba9565b8408095f516020611d895f395f51905f5261186782611d20565b80091415958691611bcc565b929080821480611907575b156118a55750505050905f1461189d5760ff60025b169060021b179190565b60ff5f611893565b5f516020611d895f395f51905f528091068103061491826118e8575b50501561152a57600191156118e05760ff60025b169060021b17179190565b60ff5f6118d5565b5f516020611d895f395f51905f52919250819006810306145f806118c1565b5083831461187e565b50505090505f905f90565b505f516020611d895f395f51905f52811015611756565b505f516020611d895f395f51905f5282101561174f565b505f516020611d895f395f51905f52851015611748565b9080601f83011215610ba4576040519161197c610100846115e1565b82906101008101928311610ba457905b8282106119995750505090565b813581526020918201910161198c565b9080601f83011215610ba4576040805192906119c590846115e1565b829060408101928311610ba457905b8282106119e15750505090565b81358152602091820191016119d4565b60405190611a006020836115e1565b5f8252565b8015611a69578060011c915f516020611d895f395f51905f5283101561152a57600180611a485f516020611d895f395f51905f5260038188818180090908611ba9565b931614611a5157565b905f516020611d895f395f51905f5280910681030690565b505f905f90565b801580611ba1575b611b95578060021c92825f516020611d895f395f51905f528510801590611b7e575b61152a5784815f516020611d895f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4481611b489d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611bcc565b80929160018082961614611b5a575050565b5f516020611d895f395f51905f528093945080929550809106810306930681030690565b505f516020611d895f395f51905f52811015611a9a565b50505f905f905f905f90565b508115611a78565b90611bb382611d20565b915f516020611d895f395f51905f528380090361152a57565b915f516020611d895f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c2493969496611c1682808a8009818a800908611ba9565b90611d14575b860809611ba9565b925f516020611d895f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611d895f395f51905f5260a083015260208260c08160055afa9151911561152a575f516020611d895f395f51905f5282600192090361152a575f516020611d895f395f51905f52908209925f516020611d895f395f51905f528080808780090681030681878009081490811591611cf5575b5061152a57565b90505f516020611d895f395f51905f528084860960020914155f611cee565b81809106810306611c1c565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611d895f395f51905f5260a083015260208260c08160055afa9151911561152a5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220821e7433112f0ed47efdc17d5ecd65a5783050fabd5e0142fc809419bff696b564736f6c634300081c0033",
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
