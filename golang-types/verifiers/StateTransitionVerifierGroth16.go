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
	Bin: "0x60808060405234601557611e02908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461157f575080639ed57f3214610cbb578063b1c3a00e14610baa578063b8e72af614610a335763e2dd9f4714610055575f80fd5b34610a30576101c0366003190112610a305736608411610a30573660a411610a3057366101c411610a3057604051602061008f81836115e5565b803683376040918251926100a381856115e5565b80368537610300938151926100b886856115e5565b853685376100c7608435611a09565b8684015282526100d860a435611a09565b905f516020611dad5f395f51905f5285516100f387826115e5565b6001815288610104818301601f198a01368237378551906101328a610124818a0151938b5194859384019687611607565b03601f1981018352826115e5565b5190200683528351865286840151878701527f18ce07d3a0e0329a60a7c40a99733a6009730ee9eece59213add463c88ebb417858701527f052028973c5dffd65fa9e214c088ea7d9224dcff73c4d6f64cd4ef82cd8dce7760608701527f228fb97326df67250c01c87e0875c078123367221076dc1097330b15ca9db86260808701527f0956323fba91950d9e99b6277c0b98a3782f8cb358f8abe6272f0850faf2b58e60a087015260c086015260e08501527f2c197812dc02324b9b85c2bfb45840f3467c99ecc82fc6ce574237f0b3e992236101008501527f061a118834bedaa67d5305d3ce8dba47c343f1b39ae6fb51dd0b5bb39a464b846101208501527f188f41e3c01b988d7907ddf89bb35fc65022f7dc3ffdccde38e273439f458f396101408501527f1372a959a565478feced82efaf3a32b26d74e817f34a196d186ed9c3285d31f3610160850152825185816101808760085afa90511615610a21578490876102a3600435611a09565b866102b5602497939735604435611a74565b9491929390956102c6606435611a09565b999098507f261550015b69d48418d328ed2f4d5d4a8a0364a07a40babb1a8ffe1ad794fdee83519b8c937f2762b5e2c67298a3f79fc5c2e1ad85c915001da374bfa87a193bc911dec235f085527f10870e3d772123a7718ddb81076253d9b747562e0d1f942ebdecb6ada55f6c00828601528051868601520151606084019081527f2cb3368150c5fe73a921d3863e23d0c830d94b164c8a4fbed7a906ec77dbcd3f858560808160065afa957f062fbc3ad2566cbb55334df3d700cf3b5e34f1a0f18cecfcb64b0e7eb0ec5c6c818701527f1e27d21a3b3d3a62f294f436d58ca5c31271aca22bed0ab46c31e6869cea02b48352600160c4356080880198818a525f516020611dad5f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f068098b40c34881453b7ecbc1d55dc2d705b23a3d14bf7a1bf34db446b47c949828801527f29363081fd6a3d2b08019b410041929fe4e32ca7149f7b1c2d704bb4d40a8c57845260e435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f30006087bdb3e0230304a08ccd1b00d4c08930683fa3073cbc7fbb2ddf3aac74828801527f21bbfb874221750a364721a2b7edb9dd8106b357ec93c0ab0943b84a31748922845261010435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f16635daf71f37096a3c3812cdc7acc1802f6dc7f93914bd82f3009dc535ee157828801527f0d525466657fcd858439b3aa39557c6dd47d6285971bef257b37af9d436cfeeb845261012435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1fa015de422bc76ddf09f1c1fb4eca285ca1273fc098f6f539502c0543709f10828801527f2466ee4c5a2d0e83a63b7b9261b94aa9a8255ac3dc38f79e16f667776df350cf845261014435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f167172f32145033ae345a6afc970e58a80c7d57eed2ed1917f7565a73ced791d828801527f11562cfaa3ad9df5c0642fac9bcd0962173a55d181c5496db04467bb299416a1845261016435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f269c408d78b2e3b6d574c862c31ece5f83af87ee5251c5f7a8f4727f2114aeae828801527f2834b43342fee057a92dcb2864065f19985a362aeca2dad18b50d39228c1004f845261018435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f08d19450e3a4633830614dda9046f49e5fd081fe41ccda95e3d8d0e82dee86a3828801527f2451cd88e73f758b540bb45f55f0f4c997a0bdf3a7cf243d9ec72897d2116b7684526101a435908189525f516020611dad5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611dad5f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610a12578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f0dee0ef9c4eba8f154781ec4b11c81f53e38979078b71fc0e22eeaa1fa76f3316101008501527f0b2e88d90fa9c00ee902fc108ed875f741069a2e46f3b8311c9afdc4f4e445326101208501527f2285c039ca7b06061fa0dfdefd5d0b16201858243ef4a18760af5f4ebdcd9e2b6101408501527f2b7aa1ee070838a981605148be3d9f54b449fd3335cf055c16e10b9b0faedb8a6101608501527f0f73c7e56a29cc65c878431c1cb135a044c9c27684868115d8487de10f2760ee6101808501527f0e69b3a109bcc1cfa26dd6ed953daf7f7109a6a2d23040cab32dc0294e7aafe96101a08501527f29f93bb5e8ef1c60c8fbaf09f635d1e0c4fcfeb0e77f95bd45982776ae1695106101c08501527f037c5ba28338ee8de9ada7b565a82446389f442c2db9c791cfaa6925592a7d5b6101e08501527f0980b4a66952bbc49111d4afa4ce49741da890c8f440719cfc689c190de5cc296102008501527f06590b5e4d12960e5b548a451a82a68d95395b78073a4c518874c762d8d8fe5b6102208501526102408401526102608301527f039f9e23adcc132c46d2166557f92f4bd92f0a911001afed9aef508417d5b9f76102808301527f1337d507cb41f6d774f75e3350a6f20fd79ec255c63922c732da66ba34eaac0f6102a08301527f1e23e5141bc165333de2ab8f4661a4aa9e8fbf09a0308fe2be20b062278a0b5c6102c08301527f0a96e1aa86f1a34fc6908809025d2385f4302de43e0682fa5def3295715291a86102e0830152519283916109dd84846115e5565b8336843760085afa15908115610a05575b506109f65780f35b631ff3747d60e21b8152600490fd5b600191505114155f6109ee565b63a54f8e2760e01b8f5260048ffd5b6351d49ff760e11b8752600487fd5b80fd5b5034610ba6576040366003190112610ba65760043567ffffffffffffffff8111610ba657610a659036906004016115b7565b60243567ffffffffffffffff8111610ba657610a859036906004016115b7565b90926101009384604051610a9982826115e5565b369037604093848051610aac82826115e5565b369037848051610abc82826115e5565b36903782019461018083870312610ba657610af1610ada8785611964565b96610140610aea828588016119ad565b95016119ad565b93818651610aff82826115e5565b36903782019082820312610ba657610b1691611964565b90610b1f6119f5565b50610b286119f5565b50303b15610ba657610b7692610b60610b6b92610b55875198634f6abf9960e11b8a5260048a0190611641565b610104880190611668565b610144860190611668565b610184840190611641565b5f8261028481305afa908115610b9d5750610b8f575080f35b610b9b91505f906115e5565b005b513d5f823e3d90fd5b5f80fd5b34610ba657610180366003190112610ba6573661010411610ba6573661014411610ba6573661018411610ba657604051610be56080826115e5565b6080368237604051906020610bfa81846115e5565b80368437610c0c60243560043561168f565b8252610c2260843560a435604435606435611730565b828401526040830152610c3960e43560c43561168f565b6060830152610c4e610124356101043561168f565b8352610c60610164356101443561168f565b9060405192835f905b60048210610ca5575050506080830193905f945b60018610610c905760c0858560a0820152f35b81806001928551815201930195019491610c7d565b8251815291830191600191909101908301610c69565b34610ba657610280366003190112610ba6573661010411610ba6573661014411610ba6573661018411610ba6573661028411610ba6576040516020610d0081836115e5565b803683376040915f516020611dad5f395f51905f528351610d2185826115e5565b60018152836101c4818301601f198801368237378451610d508161012487820194610124356101043587611607565b51902006815282518361010482377f18ce07d3a0e0329a60a7c40a99733a6009730ee9eece59213add463c88ebb417848201527f052028973c5dffd65fa9e214c088ea7d9224dcff73c4d6f64cd4ef82cd8dce7760608201527f228fb97326df67250c01c87e0875c078123367221076dc1097330b15ca9db86260808201527f0956323fba91950d9e99b6277c0b98a3782f8cb358f8abe6272f0850faf2b58e60a08201528361014460c08301377f2c197812dc02324b9b85c2bfb45840f3467c99ecc82fc6ce574237f0b3e992236101008201527f061a118834bedaa67d5305d3ce8dba47c343f1b39ae6fb51dd0b5bb39a464b846101208201527f188f41e3c01b988d7907ddf89bb35fc65022f7dc3ffdccde38e273439f458f396101408201527f1372a959a565478feced82efaf3a32b26d74e817f34a196d186ed9c3285d31f361016082015282816101808160085afa9051161561157057825192610eba6040856115e5565b61010483855b610144831061154c575050508051918183017f2762b5e2c67298a3f79fc5c2e1ad85c915001da374bfa87a193bc911dec235f084525f516020611dad5f395f51905f528386808701987f10870e3d772123a7718ddb81076253d9b747562e0d1f942ebdecb6ada55f6c008a528051855201519260608701938452818760808160065afa947f062fbc3ad2566cbb55334df3d700cf3b5e34f1a0f18cecfcb64b0e7eb0ec5c6c82527f1e27d21a3b3d3a62f294f436d58ca5c31271aca22bed0ab46c31e6869cea02b485527f261550015b69d48418d328ed2f4d5d4a8a0364a07a40babb1a8ffe1ad794fdee60016101843560808b0198818a5287878760608160075afa92101616858b60808160065afa16167f068098b40c34881453b7ecbc1d55dc2d705b23a3d14bf7a1bf34db446b47c94984527f29363081fd6a3d2b08019b410041929fe4e32ca7149f7b1c2d704bb4d40a8c5787526101a4359081895286868660608160075afa92101616848a60808160065afa167f30006087bdb3e0230304a08ccd1b00d4c08930683fa3073cbc7fbb2ddf3aac7484527f21bbfb874221750a364721a2b7edb9dd8106b357ec93c0ab0943b84a3174892287526101c4359081895286868660608160075afa92101616848a60808160065afa167f16635daf71f37096a3c3812cdc7acc1802f6dc7f93914bd82f3009dc535ee15784527f0d525466657fcd858439b3aa39557c6dd47d6285971bef257b37af9d436cfeeb87526101e4359081895286868660608160075afa92101616848a60808160065afa167f1fa015de422bc76ddf09f1c1fb4eca285ca1273fc098f6f539502c0543709f1084527f2466ee4c5a2d0e83a63b7b9261b94aa9a8255ac3dc38f79e16f667776df350cf8752610204359081895286868660608160075afa92101616848a60808160065afa167f167172f32145033ae345a6afc970e58a80c7d57eed2ed1917f7565a73ced791d84527f11562cfaa3ad9df5c0642fac9bcd0962173a55d181c5496db04467bb299416a18752610224359081895286868660608160075afa92101616848a60808160065afa167f269c408d78b2e3b6d574c862c31ece5f83af87ee5251c5f7a8f4727f2114aeae84527f2834b43342fee057a92dcb2864065f19985a362aeca2dad18b50d39228c1004f8752610244359081895286868660608160075afa92101616848a60808160065afa167f08d19450e3a4633830614dda9046f49e5fd081fe41ccda95e3d8d0e82dee86a384527f2451cd88e73f758b540bb45f55f0f4c997a0bdf3a7cf243d9ec72897d2116b768752610264359081895286868660608160075afa92101616848a60808160065afa16957f2cb3368150c5fe73a921d3863e23d0c830d94b164c8a4fbed7a906ec77dbcd3f8452525180955260608160075afa92101616818360808160065afa1691519351911561153d575192610100600485377f0dee0ef9c4eba8f154781ec4b11c81f53e38979078b71fc0e22eeaa1fa76f3316101008501527f0b2e88d90fa9c00ee902fc108ed875f741069a2e46f3b8311c9afdc4f4e445326101208501527f2285c039ca7b06061fa0dfdefd5d0b16201858243ef4a18760af5f4ebdcd9e2b6101408501527f2b7aa1ee070838a981605148be3d9f54b449fd3335cf055c16e10b9b0faedb8a6101608501527f0f73c7e56a29cc65c878431c1cb135a044c9c27684868115d8487de10f2760ee6101808501527f0e69b3a109bcc1cfa26dd6ed953daf7f7109a6a2d23040cab32dc0294e7aafe96101a08501527f29f93bb5e8ef1c60c8fbaf09f635d1e0c4fcfeb0e77f95bd45982776ae1695106101c08501527f037c5ba28338ee8de9ada7b565a82446389f442c2db9c791cfaa6925592a7d5b6101e08501527f0980b4a66952bbc49111d4afa4ce49741da890c8f440719cfc689c190de5cc296102008501527f06590b5e4d12960e5b548a451a82a68d95395b78073a4c518874c762d8d8fe5b6102208501526102408401526102608301527f039f9e23adcc132c46d2166557f92f4bd92f0a911001afed9aef508417d5b9f76102808301527f1337d507cb41f6d774f75e3350a6f20fd79ec255c63922c732da66ba34eaac0f6102a08301527f1e23e5141bc165333de2ab8f4661a4aa9e8fbf09a0308fe2be20b062278a0b5c6102c08301527f0a96e1aa86f1a34fc6908809025d2385f4302de43e0682fa5def3295715291a86102e0830152816103008160085afa9051161561152e57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191859101610ec0565b634e487b7160e01b5f52604160045260245ffd5b6351d49ff760e11b5f5260045ffd5b34610ba6575f366003190112610ba657807f37fdaef3c4b55dc3a608db50bade1f868e67d55db41c15dd4665eca9588adcfe60209252f35b9181601f84011215610ba65782359167ffffffffffffffff8311610ba65760208381860195010111610ba657565b90601f8019910116810190811067ffffffffffffffff82111761155c57604052565b9091604092825260208201520160208251919201905f5b81811061162b5750505090565b825184526020938401939092019160010161161e565b905f905b6008821061165257505050565b6020806001928551815201930191019091611645565b905f905b6002821061167957505050565b602080600192855181520193019101909161166c565b905f516020611d8d5f395f51905f528210801590611719575b61152e57811580611711575b61170b576116d85f516020611d8d5f395f51905f5260038185818180090908611bad565b8181036116e757505060011b90565b5f516020611d8d5f395f51905f52809106810306145f1461152e57600190811b1790565b50505f90565b5080156116b4565b505f516020611d8d5f395f51905f528110156116a8565b919093925f516020611d8d5f395f51905f52831080159061194d575b8015611936575b801561191f575b61152e578082868517171715611914579082916118775f516020611d8d5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611d8d5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161185181808b80098187800908611bad565b8408095f516020611d8d5f395f51905f5261186b82611d24565b80091415958691611bd0565b92908082148061190b575b156118a95750505050905f146118a15760ff60025b169060021b179190565b60ff5f611897565b5f516020611d8d5f395f51905f528091068103061491826118ec575b50501561152e57600191156118e45760ff60025b169060021b17179190565b60ff5f6118d9565b5f516020611d8d5f395f51905f52919250819006810306145f806118c5565b50838314611882565b50505090505f905f90565b505f516020611d8d5f395f51905f5281101561175a565b505f516020611d8d5f395f51905f52821015611753565b505f516020611d8d5f395f51905f5285101561174c565b9080601f83011215610ba65760405191611980610100846115e5565b82906101008101928311610ba657905b82821061199d5750505090565b8135815260209182019101611990565b9080601f83011215610ba6576040805192906119c990846115e5565b829060408101928311610ba657905b8282106119e55750505090565b81358152602091820191016119d8565b60405190611a046020836115e5565b5f8252565b8015611a6d578060011c915f516020611d8d5f395f51905f5283101561152e57600180611a4c5f516020611d8d5f395f51905f5260038188818180090908611bad565b931614611a5557565b905f516020611d8d5f395f51905f5280910681030690565b505f905f90565b801580611ba5575b611b99578060021c92825f516020611d8d5f395f51905f528510801590611b82575b61152e5784815f516020611d8d5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4481611b4c9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611bd0565b80929160018082961614611b5e575050565b5f516020611d8d5f395f51905f528093945080929550809106810306930681030690565b505f516020611d8d5f395f51905f52811015611a9e565b50505f905f905f905f90565b508115611a7c565b90611bb782611d24565b915f516020611d8d5f395f51905f528380090361152e57565b915f516020611d8d5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c2893969496611c1a82808a8009818a800908611bad565b90611d18575b860809611bad565b925f516020611d8d5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611d8d5f395f51905f5260a083015260208260c08160055afa9151911561152e575f516020611d8d5f395f51905f5282600192090361152e575f516020611d8d5f395f51905f52908209925f516020611d8d5f395f51905f528080808780090681030681878009081490811591611cf9575b5061152e57565b90505f516020611d8d5f395f51905f528084860960020914155f611cf2565b81809106810306611c20565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611d8d5f395f51905f5260a083015260208260c08160055afa9151911561152e5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a264697066735822122088f97cc59551b876464c0bf80b883a6d636262de04aeb1fd21cf61f2bf60b10464736f6c634300081c0033",
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
