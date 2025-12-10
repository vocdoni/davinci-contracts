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
	Bin: "0x60808060405234601557611e72908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace11146115ad575080639ed57f3214610cea578063b1c3a00e14610bd9578063b8e72af614610a325763e2dd9f4714610055575f80fd5b34610a2f576101c0366003190112610a2f5736608411610a2f573660a411610a2f57366101c411610a2f57604051602061008f8183611613565b803683376040918251926100a38185611613565b80368537610300938151926100b88685611613565b853685376100c7608435611a79565b8684015282526100d860a435611a79565b905f516020611e1d5f395f51905f5285516100f38782611613565b6001815288610104818301601f198a01368237378551906101328a610124818a0151938b5194859384019687611635565b03601f198101835282611613565b5190200683528351865286840151878701527f0e201f948325683a00cb2b504d0e1aa7d8d2dcc6562d5df255f564744c4eae6f858701527f16b2417e724e53fc00ffea0fc36e83f927d46b494aee8a26236acda11862e43060608701527f1be5289618579bc4d060540694d5af09b9ca0c7478797f6effee3f49c732757860808701527f04b31f0baa24c244dae7044a9272531bbf2cd5b5ecdcf22f0084687ec34e37c360a087015260c086015260e08501527f2e46e9eba835696819b6ec2d833042ab80df99ec9f2cb357d7340e96594cd7306101008501527f19d1588db166747ccbb2e16ca79e378f46d0f7bc2b49646c7070ac5b0460b0296101208501527e50c0d9645d908fa4e838191f7debe929b7b5ad7c364d2b99ee8c81c4b7094c6101408501527f0d5d75cc6eabb5e199699e67010439a7f96e330a8ab5b66fdc58639514b4d374610160850152825185816101808760085afa90511615610a20578490876102a2600435611a79565b866102b4602497939735604435611ae4565b9491929390956102c5606435611a79565b999098507f2590fa6590c183df3d9d60739fe230b89e4e3a061409a39dc2822092eb3446f183519b8c937f16f7c293fe8b9c9ef6a56367df54b1d12eb367b2896a17beb90d8f4c9c5cef8d85527f15771c94333ae94634dda38efdf83d08f2e8ef5cdade1ddbf3c496cae6876ef2828601528051868601520151606084019081527f29709e4d3dd7094789e2c3d922ea5dae3ecdf96f48116f3ec5d5602f860f31fb858560808160065afa957f16708c4b60a458f981206636d98cb151275675da10a31f7e80c905c37a261867818701527f27775cafa656a945c5d3624d5b03c0e3b30be203bcaf86cbfc7de5549af085588352600160c4356080880198818a525f516020611e1d5f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f1c470a0230579ed45a3a7a01f18183b905e119bdf85b92435f13f7fc7850761f828801527f23f9ed943c5efac6d38fa1c9b1e68ed4580a8b33554e098bd39ada1c1613e107845260e435908189525f516020611e1d5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0c6a28b44385ab7e895fc2e38a5639dbb199f7f45692ad74627409e59867ac47828801527f031244326f497f43a8657ec4d8cb6705453b4e32752e3bcb4bdeba786b2eb640845261010435908189525f516020611e1d5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f03ce00d2d03c5938a86ef52dd73c0b0d4b260a11a0ed9f48021f523fae545f5d828801527f2738863475c366d84b27a1ac49d19697017276def448fdc17e4325ddcd6568e2845261012435908189525f516020611e1d5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1c2f8310e99eac6beee16281c781e0a0342569f3e3548a1f4e5921e40c55f99d828801527f1ec9a8d32ca33d3b8f7cdb5807c9c10bfaa5c8e4a20dafc1e7f585e26472a9de845261014435908189525f516020611e1d5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f05fa9a70bb794b570dd173c5983375b4ca49d668068575a4ffce6aa54d4a0dbd828801527f159e9744bcd5f3435d59b5e8adb0e966c2ec847c6fe8619a7053bf771cd9108e845261016435908189525f516020611e1d5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f275f76b2af6a49ceb18db9d53e5da22584cfa89014e5d02f5d6a91ceb3c9bbbb828801527f02f7dc711eda15f822a83bc5d764dfd18b03114b4e75934e76ff341024eecdef845261018435908189525f516020611e1d5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f125fd00e0323b8d1da48b126f8a21ef1c7c39396e1fc4f7f6b170126d430644c828801527f0aa5b2621b01b5f80e68997186da8e0af4dae4cf0c3ab9c274df8094160090ff84526101a435908189525f516020611e1d5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611e1d5f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610a11578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f11e98fdab73defabb10104a30763f01ef7cf993913816277fc2aad2de45c859c6101008501527f2fd78bf77b94285f502c7c33ddf8d34ee338264470a666d46809c8f5c19e631c6101208501527f1401f13fa0490de4680a2df7d6ef7ebd4068981b13af6601a40e699cb47317fb6101408501527f13211bf0599166ef1c5a025da306bded2d5f105bcf409bb9c717026b95bb63506101608501527f04373dad598fa352c4aca4f4049adb63937e804ac0bb8c08325a705dc44d93e76101808501527f08a8665c04db93b632430b6f031392db8464fe4e50e61f093464f0a26adcafe96101a08501527f2bf97737b5f4a2929ec01900b27bc75ec657fb208f04c66b7444dfa5caf999cc6101c08501527f2cc112b07d30f40ed69c23711ca51605797748933e2efd10dff8995762848e396101e08501527f01e022c4d7d69f9a4a9e1a5c2beebe0e2823dff866033df508ac4fafe2a2a8746102008501527f2047e808302da92312405e756d583af36f898175a0b8c3f2410a2ecf86cd0b376102208501526102408401526102608301527f2ad6415a374c68d778090cdf50feced19ecfd84a95b9d3b4e91486e277e196996102808301527f2290d0a2d8c994b707e69721651935ff63ecf78eacc62b94ef4345f17ef16d436102a08301527f224af2a913eb01386814607e82751b36eb9fdc6d61b83a19523f0ad0e9116a836102c08301527f1d65e675774e61d51b5147f0c5452973c22a2d0ccc333fe8a170a6b691bda6796102e0830152519283916109dc8484611613565b8336843760085afa15908115610a04575b506109f55780f35b631ff3747d60e21b8152600490fd5b600191505114155f6109ed565b63a54f8e2760e01b8f5260048ffd5b6351d49ff760e11b8752600487fd5b80fd5b5034610bd5576040366003190112610bd55760043567ffffffffffffffff8111610bd557610a649036906004016115e5565b60243567ffffffffffffffff8111610bd557610a849036906004016115e5565b90926101009081604051610a988282611613565b369037604093848051610aab8282611613565b369037848051610abb8282611613565b36903781019461018082870312610bd557610af0610ad98784611992565b96610140610ae9828787016119db565b94016119db565b93838651610afe8282611613565b369037810161014082820312610bd557610b188183611992565b9382013567ffffffffffffffff8111610bd55781610b37918401611a23565b5061012082013567ffffffffffffffff8111610bd557610b579201611a23565b50303b15610bd557610ba592610b8f610b9a92610b84875198634f6abf9960e11b8a5260048a019061166f565b610104880190611696565b610144860190611696565b61018484019061166f565b5f8261028481305afa908115610bcc5750610bbe575080f35b610bca91505f90611613565b005b513d5f823e3d90fd5b5f80fd5b34610bd557610180366003190112610bd5573661010411610bd5573661014411610bd5573661018411610bd557604051610c14608082611613565b6080368237604051906020610c298184611613565b80368437610c3b6024356004356116bd565b8252610c5160843560a43560443560643561175e565b828401526040830152610c6860e43560c4356116bd565b6060830152610c7d61012435610104356116bd565b8352610c8f61016435610144356116bd565b9060405192835f905b60048210610cd4575050506080830193905f945b60018610610cbf5760c0858560a0820152f35b81806001928551815201930195019491610cac565b8251815291830191600191909101908301610c98565b34610bd557610280366003190112610bd5573661010411610bd5573661014411610bd5573661018411610bd5573661028411610bd5576040516020610d2f8183611613565b803683376040915f516020611e1d5f395f51905f528351610d508582611613565b60018152836101c4818301601f198801368237378451610d7f8161012487820194610124356101043587611635565b51902006815282518361010482377f0e201f948325683a00cb2b504d0e1aa7d8d2dcc6562d5df255f564744c4eae6f848201527f16b2417e724e53fc00ffea0fc36e83f927d46b494aee8a26236acda11862e43060608201527f1be5289618579bc4d060540694d5af09b9ca0c7478797f6effee3f49c732757860808201527f04b31f0baa24c244dae7044a9272531bbf2cd5b5ecdcf22f0084687ec34e37c360a08201528361014460c08301377f2e46e9eba835696819b6ec2d833042ab80df99ec9f2cb357d7340e96594cd7306101008201527f19d1588db166747ccbb2e16ca79e378f46d0f7bc2b49646c7070ac5b0460b0296101208201527e50c0d9645d908fa4e838191f7debe929b7b5ad7c364d2b99ee8c81c4b7094c6101408201527f0d5d75cc6eabb5e199699e67010439a7f96e330a8ab5b66fdc58639514b4d37461016082015282816101808160085afa9051161561159e57825192610ee8604085611613565b61010483855b610144831061157a575050508051918183017f16f7c293fe8b9c9ef6a56367df54b1d12eb367b2896a17beb90d8f4c9c5cef8d84525f516020611e1d5f395f51905f528386808701987f15771c94333ae94634dda38efdf83d08f2e8ef5cdade1ddbf3c496cae6876ef28a528051855201519260608701938452818760808160065afa947f16708c4b60a458f981206636d98cb151275675da10a31f7e80c905c37a26186782527f27775cafa656a945c5d3624d5b03c0e3b30be203bcaf86cbfc7de5549af0855885527f2590fa6590c183df3d9d60739fe230b89e4e3a061409a39dc2822092eb3446f160016101843560808b0198818a5287878760608160075afa92101616858b60808160065afa16167f1c470a0230579ed45a3a7a01f18183b905e119bdf85b92435f13f7fc7850761f84527f23f9ed943c5efac6d38fa1c9b1e68ed4580a8b33554e098bd39ada1c1613e10787526101a4359081895286868660608160075afa92101616848a60808160065afa167f0c6a28b44385ab7e895fc2e38a5639dbb199f7f45692ad74627409e59867ac4784527f031244326f497f43a8657ec4d8cb6705453b4e32752e3bcb4bdeba786b2eb64087526101c4359081895286868660608160075afa92101616848a60808160065afa167f03ce00d2d03c5938a86ef52dd73c0b0d4b260a11a0ed9f48021f523fae545f5d84527f2738863475c366d84b27a1ac49d19697017276def448fdc17e4325ddcd6568e287526101e4359081895286868660608160075afa92101616848a60808160065afa167f1c2f8310e99eac6beee16281c781e0a0342569f3e3548a1f4e5921e40c55f99d84527f1ec9a8d32ca33d3b8f7cdb5807c9c10bfaa5c8e4a20dafc1e7f585e26472a9de8752610204359081895286868660608160075afa92101616848a60808160065afa167f05fa9a70bb794b570dd173c5983375b4ca49d668068575a4ffce6aa54d4a0dbd84527f159e9744bcd5f3435d59b5e8adb0e966c2ec847c6fe8619a7053bf771cd9108e8752610224359081895286868660608160075afa92101616848a60808160065afa167f275f76b2af6a49ceb18db9d53e5da22584cfa89014e5d02f5d6a91ceb3c9bbbb84527f02f7dc711eda15f822a83bc5d764dfd18b03114b4e75934e76ff341024eecdef8752610244359081895286868660608160075afa92101616848a60808160065afa167f125fd00e0323b8d1da48b126f8a21ef1c7c39396e1fc4f7f6b170126d430644c84527f0aa5b2621b01b5f80e68997186da8e0af4dae4cf0c3ab9c274df8094160090ff8752610264359081895286868660608160075afa92101616848a60808160065afa16957f29709e4d3dd7094789e2c3d922ea5dae3ecdf96f48116f3ec5d5602f860f31fb8452525180955260608160075afa92101616818360808160065afa1691519351911561156b575192610100600485377f11e98fdab73defabb10104a30763f01ef7cf993913816277fc2aad2de45c859c6101008501527f2fd78bf77b94285f502c7c33ddf8d34ee338264470a666d46809c8f5c19e631c6101208501527f1401f13fa0490de4680a2df7d6ef7ebd4068981b13af6601a40e699cb47317fb6101408501527f13211bf0599166ef1c5a025da306bded2d5f105bcf409bb9c717026b95bb63506101608501527f04373dad598fa352c4aca4f4049adb63937e804ac0bb8c08325a705dc44d93e76101808501527f08a8665c04db93b632430b6f031392db8464fe4e50e61f093464f0a26adcafe96101a08501527f2bf97737b5f4a2929ec01900b27bc75ec657fb208f04c66b7444dfa5caf999cc6101c08501527f2cc112b07d30f40ed69c23711ca51605797748933e2efd10dff8995762848e396101e08501527f01e022c4d7d69f9a4a9e1a5c2beebe0e2823dff866033df508ac4fafe2a2a8746102008501527f2047e808302da92312405e756d583af36f898175a0b8c3f2410a2ecf86cd0b376102208501526102408401526102608301527f2ad6415a374c68d778090cdf50feced19ecfd84a95b9d3b4e91486e277e196996102808301527f2290d0a2d8c994b707e69721651935ff63ecf78eacc62b94ef4345f17ef16d436102a08301527f224af2a913eb01386814607e82751b36eb9fdc6d61b83a19523f0ad0e9116a836102c08301527f1d65e675774e61d51b5147f0c5452973c22a2d0ccc333fe8a170a6b691bda6796102e0830152816103008160085afa9051161561155c57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191859101610eee565b634e487b7160e01b5f52604160045260245ffd5b6351d49ff760e11b5f5260045ffd5b34610bd5575f366003190112610bd557807f95c27ede6d7d7b092ad108b5ba29f66fdeb69be78613da14c1994e94073370c060209252f35b9181601f84011215610bd55782359167ffffffffffffffff8311610bd55760208381860195010111610bd557565b90601f8019910116810190811067ffffffffffffffff82111761158a57604052565b9091604092825260208201520160208251919201905f5b8181106116595750505090565b825184526020938401939092019160010161164c565b905f905b6008821061168057505050565b6020806001928551815201930191019091611673565b905f905b600282106116a757505050565b602080600192855181520193019101909161169a565b905f516020611dfd5f395f51905f528210801590611747575b61155c5781158061173f575b611739576117065f516020611dfd5f395f51905f5260038185818180090908611c1d565b81810361171557505060011b90565b5f516020611dfd5f395f51905f52809106810306145f1461155c57600190811b1790565b50505f90565b5080156116e2565b505f516020611dfd5f395f51905f528110156116d6565b919093925f516020611dfd5f395f51905f52831080159061197b575b8015611964575b801561194d575b61155c578082868517171715611942579082916118a55f516020611dfd5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611dfd5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161187f81808b80098187800908611c1d565b8408095f516020611dfd5f395f51905f5261189982611d94565b80091415958691611c40565b929080821480611939575b156118d75750505050905f146118cf5760ff60025b169060021b179190565b60ff5f6118c5565b5f516020611dfd5f395f51905f5280910681030614918261191a575b50501561155c57600191156119125760ff60025b169060021b17179190565b60ff5f611907565b5f516020611dfd5f395f51905f52919250819006810306145f806118f3565b508383146118b0565b50505090505f905f90565b505f516020611dfd5f395f51905f52811015611788565b505f516020611dfd5f395f51905f52821015611781565b505f516020611dfd5f395f51905f5285101561177a565b9080601f83011215610bd557604051916119ae61010084611613565b82906101008101928311610bd557905b8282106119cb5750505090565b81358152602091820191016119be565b9080601f83011215610bd5576040805192906119f79084611613565b829060408101928311610bd557905b828210611a135750505090565b8135815260209182019101611a06565b81601f82011215610bd55780359067ffffffffffffffff821161158a5760405192611a58601f8401601f191660200185611613565b82845260208383010111610bd557815f926020809301838601378301015290565b8015611add578060011c915f516020611dfd5f395f51905f5283101561155c57600180611abc5f516020611dfd5f395f51905f5260038188818180090908611c1d565b931614611ac557565b905f516020611dfd5f395f51905f5280910681030690565b505f905f90565b801580611c15575b611c09578060021c92825f516020611dfd5f395f51905f528510801590611bf2575b61155c5784815f516020611dfd5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4481611bbc9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611c40565b80929160018082961614611bce575050565b5f516020611dfd5f395f51905f528093945080929550809106810306930681030690565b505f516020611dfd5f395f51905f52811015611b0e565b50505f905f905f905f90565b508115611aec565b90611c2782611d94565b915f516020611dfd5f395f51905f528380090361155c57565b915f516020611dfd5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c9893969496611c8a82808a8009818a800908611c1d565b90611d88575b860809611c1d565b925f516020611dfd5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611dfd5f395f51905f5260a083015260208260c08160055afa9151911561155c575f516020611dfd5f395f51905f5282600192090361155c575f516020611dfd5f395f51905f52908209925f516020611dfd5f395f51905f528080808780090681030681878009081490811591611d69575b5061155c57565b90505f516020611dfd5f395f51905f528084860960020914155f611d62565b81809106810306611c90565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611dfd5f395f51905f5260a083015260208260c08160055afa9151911561155c5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212209a6d9f7f196aaea0a0bd49d025d361a4338e83b766d7eb68148f59f291b65f0a64736f6c634300081c0033",
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
