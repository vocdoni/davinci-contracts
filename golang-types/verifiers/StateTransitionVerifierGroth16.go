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
	Bin: "0x60808060405234601557611e00908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461157d575080639ed57f3214610cba578063b1c3a00e14610ba9578063b8e72af614610a325763e2dd9f4714610055575f80fd5b34610a2f576101c0366003190112610a2f5736608411610a2f573660a411610a2f57366101c411610a2f57604051602061008f81836115e3565b803683376040918251926100a381856115e3565b80368537610300938151926100b886856115e3565b853685376100c7608435611a07565b8684015282526100d860a435611a07565b905f516020611dab5f395f51905f5285516100f387826115e3565b6001815288610104818301601f198a01368237378551906101328a610124818a0151938b5194859384019687611605565b03601f1981018352826115e3565b5190200683528351865286840151878701527f189a718fa2caba0d2a086ebb7dfc022d4616d96677504ea2f5b97ceb3410c845858701527f043f582c3637557bbfc332ee85adda2a6cd63d7303d86884d4432a3ae3714f6760608701527f0ddb6084f2d95aa4be0fe8ee9c04f548efcca1fa1371a6b993b9cdf9c4115dc560808701527f213674c5ea7568d4e00dfc2280faef562fbfae521973894cf54b3bf68828e22a60a087015260c086015260e08501527f02c116cb1308b1bb25e7cda527c6122e4f30abe7b5ace0df6aeda0c49c83da2c6101008501527f25f0b4a28c08fd44a425e6408dead2c9e4cdcea224d3621cc6c6b33bf4715cd36101208501527f1fa1241835d3f9b82fcb250642ca6c67c06a94918e3b638b37abaf51b43a0f5e6101408501527f0869375266d49d080d90a5146f57055ff126366776f7a14e3b3ff1f23dc9a765610160850152825185816101808760085afa90511615610a20578490876102a3600435611a07565b866102b5602497939735604435611a72565b9491929390956102c6606435611a07565b999098507f1b412492ea93330dd6305f159ba4056edd606beaca649e3e51a77b73bbaedb6683519b8c937f1f421f77004a218681e2570b79efb16179c79b62ef52858d5f492b8e18608f7585527f14b4243e1ba4a443c7f2880b9368b8bfd14f2f5c8a620b8fef2fc4fbbc9d62ed828601528051868601520151606084019081527f245b989db8c6735b1daa51cf6f3d3c53e53678b2ff90e498ac61958b5db1b909858560808160065afa957f27ad3f6dd4847989355c8b5288aeaa46e66a640381db4c328e0c4b4427f4167a818701527f1a9278f91bd4a77b5f946598652eb7fa7c207c45bb4d805522aacda1a0cd12c08352600160c4356080880198818a525f516020611dab5f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f0cce79cf22df6c1408eebf3b30780137520cb2dc348034bfb93e4f0ae8b8cc2e828801527f21f6eefec8b27c9624f875a56dc6749221652d4033ba7656f08ada6cde297920845260e435908189525f516020611dab5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f14e1da2b248004ab208335d10315388e232b734c694f28f826041e53ca4cc97c828801527f2259aea78cac63892973e962b3e1033e60cf9aa55fa1e5a68e4f11e36e360f0b845261010435908189525f516020611dab5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2b8a570326346ba98f2300e1a5dbe651bd425e3f342d9325190c6232aa2be88b828801527f2fa54c27ebc5234acd9a54035389fda8623d83116dfd86dd6b3a801fc5aa5793845261012435908189525f516020611dab5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1d626904fe7dc058737216c5992626b0d03fe6dbaa56f40dcd691c27fd1be40e828801527f0cc2371f183db76d6d14501f503ba4ddee6fd7a7cd256ebf3c8e6b190fa9c4dc845261014435908189525f516020611dab5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f27598d17586d49e6a0240170b2b48767557a437c45af7ea2c70b4c3b74332f7b828801527f06fbd14ff8e01cacd78bf913f8c79ef6ce368804831370d3b6f9d7c34ced5d16845261016435908189525f516020611dab5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0563eea11a900a51f4e3f2a0f035279495bd715e2e264ef4c50fd51b45cdd6ec828801527f11410812cae215d05680ed35999795b6c852362d9d58df503c323a74e570c5d7845261018435908189525f516020611dab5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f145ec2bc7e35b376cda3ca790ff83dcd48aa37ff9cf5bd5cb4542ef722289b1e828801527f28b1ad618dcc0c1cbd1b04cd242b4c3cbaa1c09f701c3d74507f7df592691fad84526101a435908189525f516020611dab5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611dab5f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610a11578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f2a398cd66f3879674610e10c84ec0bc81925da9c1f18871d10347863cb71178e6101008501527f1f5326b42cbba3e77c9f5d9829229bb151b055b986d0bd4c1d93870c7265984f6101208501527f0679ad394d3adb6ec487f7c4e4372efea0e76b1157b0784ffe38aafe234b9b3c6101408501527f0288f6358454d3dfde92bfc58b428287297b6b847723d05a8c4c61ae17547a866101608501527f06aa64ae52ca2f49ecfc6ed255c6a98561f7fcb29b205c34e27f1b5721755cab6101808501527f284a4c618229abf09c9cb81443c35b3a460e2da1f33aab6a76a5e2fd44214dc16101a08501527f0c49a4c66f3282042a746c301141d606b747a16f786e97b20c22e52b3bdee38d6101c08501527f0a267f746fcd38e9ac33aeb149abf39b535dbda7e794b8c69a8f69ea7a73577c6101e08501527f2778df932b1cc44eb25d1c9a7d74deff5c9ff0b76ce62fbe0c18b0dc2ea9f1396102008501527f2ae07fb6da3db6847939ecc3b7c3dd3b402e93ecee7015d927905173eb8da2a36102208501526102408401526102608301527f0a2ab642c246fe428aea69a87694e6bd0fa93cb9caa77df30eb815e4f8fe7cb56102808301527e6f3b1d2af841df426306e617b59f1c5e203dffedb3f74aafb154d6c4d2226f6102a08301527f12231bc6cb7bb13fc5b0facddff744e0ef5aa262b6317bb35cad358931a656356102c08301527f2fde75bb8f5f3fad412f755b522b58640d9f018a6efb1c6aaefa621d9848a9646102e0830152519283916109dc84846115e3565b8336843760085afa15908115610a04575b506109f55780f35b631ff3747d60e21b8152600490fd5b600191505114155f6109ed565b63a54f8e2760e01b8f5260048ffd5b6351d49ff760e11b8752600487fd5b80fd5b5034610ba5576040366003190112610ba55760043567ffffffffffffffff8111610ba557610a649036906004016115b5565b60243567ffffffffffffffff8111610ba557610a849036906004016115b5565b90926101009384604051610a9882826115e3565b369037604093848051610aab82826115e3565b369037848051610abb82826115e3565b36903782019461018083870312610ba557610af0610ad98785611962565b96610140610ae9828588016119ab565b95016119ab565b93818651610afe82826115e3565b36903782019082820312610ba557610b1591611962565b90610b1e6119f3565b50610b276119f3565b50303b15610ba557610b7592610b5f610b6a92610b54875198634f6abf9960e11b8a5260048a019061163f565b610104880190611666565b610144860190611666565b61018484019061163f565b5f8261028481305afa908115610b9c5750610b8e575080f35b610b9a91505f906115e3565b005b513d5f823e3d90fd5b5f80fd5b34610ba557610180366003190112610ba5573661010411610ba5573661014411610ba5573661018411610ba557604051610be46080826115e3565b6080368237604051906020610bf981846115e3565b80368437610c0b60243560043561168d565b8252610c2160843560a43560443560643561172e565b828401526040830152610c3860e43560c43561168d565b6060830152610c4d610124356101043561168d565b8352610c5f610164356101443561168d565b9060405192835f905b60048210610ca4575050506080830193905f945b60018610610c8f5760c0858560a0820152f35b81806001928551815201930195019491610c7c565b8251815291830191600191909101908301610c68565b34610ba557610280366003190112610ba5573661010411610ba5573661014411610ba5573661018411610ba5573661028411610ba5576040516020610cff81836115e3565b803683376040915f516020611dab5f395f51905f528351610d2085826115e3565b60018152836101c4818301601f198801368237378451610d4f8161012487820194610124356101043587611605565b51902006815282518361010482377f189a718fa2caba0d2a086ebb7dfc022d4616d96677504ea2f5b97ceb3410c845848201527f043f582c3637557bbfc332ee85adda2a6cd63d7303d86884d4432a3ae3714f6760608201527f0ddb6084f2d95aa4be0fe8ee9c04f548efcca1fa1371a6b993b9cdf9c4115dc560808201527f213674c5ea7568d4e00dfc2280faef562fbfae521973894cf54b3bf68828e22a60a08201528361014460c08301377f02c116cb1308b1bb25e7cda527c6122e4f30abe7b5ace0df6aeda0c49c83da2c6101008201527f25f0b4a28c08fd44a425e6408dead2c9e4cdcea224d3621cc6c6b33bf4715cd36101208201527f1fa1241835d3f9b82fcb250642ca6c67c06a94918e3b638b37abaf51b43a0f5e6101408201527f0869375266d49d080d90a5146f57055ff126366776f7a14e3b3ff1f23dc9a76561016082015282816101808160085afa9051161561156e57825192610eb96040856115e3565b61010483855b610144831061154a575050508051918183017f1f421f77004a218681e2570b79efb16179c79b62ef52858d5f492b8e18608f7584525f516020611dab5f395f51905f528386808701987f14b4243e1ba4a443c7f2880b9368b8bfd14f2f5c8a620b8fef2fc4fbbc9d62ed8a528051855201519260608701938452818760808160065afa947f27ad3f6dd4847989355c8b5288aeaa46e66a640381db4c328e0c4b4427f4167a82527f1a9278f91bd4a77b5f946598652eb7fa7c207c45bb4d805522aacda1a0cd12c085527f1b412492ea93330dd6305f159ba4056edd606beaca649e3e51a77b73bbaedb6660016101843560808b0198818a5287878760608160075afa92101616858b60808160065afa16167f0cce79cf22df6c1408eebf3b30780137520cb2dc348034bfb93e4f0ae8b8cc2e84527f21f6eefec8b27c9624f875a56dc6749221652d4033ba7656f08ada6cde29792087526101a4359081895286868660608160075afa92101616848a60808160065afa167f14e1da2b248004ab208335d10315388e232b734c694f28f826041e53ca4cc97c84527f2259aea78cac63892973e962b3e1033e60cf9aa55fa1e5a68e4f11e36e360f0b87526101c4359081895286868660608160075afa92101616848a60808160065afa167f2b8a570326346ba98f2300e1a5dbe651bd425e3f342d9325190c6232aa2be88b84527f2fa54c27ebc5234acd9a54035389fda8623d83116dfd86dd6b3a801fc5aa579387526101e4359081895286868660608160075afa92101616848a60808160065afa167f1d626904fe7dc058737216c5992626b0d03fe6dbaa56f40dcd691c27fd1be40e84527f0cc2371f183db76d6d14501f503ba4ddee6fd7a7cd256ebf3c8e6b190fa9c4dc8752610204359081895286868660608160075afa92101616848a60808160065afa167f27598d17586d49e6a0240170b2b48767557a437c45af7ea2c70b4c3b74332f7b84527f06fbd14ff8e01cacd78bf913f8c79ef6ce368804831370d3b6f9d7c34ced5d168752610224359081895286868660608160075afa92101616848a60808160065afa167f0563eea11a900a51f4e3f2a0f035279495bd715e2e264ef4c50fd51b45cdd6ec84527f11410812cae215d05680ed35999795b6c852362d9d58df503c323a74e570c5d78752610244359081895286868660608160075afa92101616848a60808160065afa167f145ec2bc7e35b376cda3ca790ff83dcd48aa37ff9cf5bd5cb4542ef722289b1e84527f28b1ad618dcc0c1cbd1b04cd242b4c3cbaa1c09f701c3d74507f7df592691fad8752610264359081895286868660608160075afa92101616848a60808160065afa16957f245b989db8c6735b1daa51cf6f3d3c53e53678b2ff90e498ac61958b5db1b9098452525180955260608160075afa92101616818360808160065afa1691519351911561153b575192610100600485377f2a398cd66f3879674610e10c84ec0bc81925da9c1f18871d10347863cb71178e6101008501527f1f5326b42cbba3e77c9f5d9829229bb151b055b986d0bd4c1d93870c7265984f6101208501527f0679ad394d3adb6ec487f7c4e4372efea0e76b1157b0784ffe38aafe234b9b3c6101408501527f0288f6358454d3dfde92bfc58b428287297b6b847723d05a8c4c61ae17547a866101608501527f06aa64ae52ca2f49ecfc6ed255c6a98561f7fcb29b205c34e27f1b5721755cab6101808501527f284a4c618229abf09c9cb81443c35b3a460e2da1f33aab6a76a5e2fd44214dc16101a08501527f0c49a4c66f3282042a746c301141d606b747a16f786e97b20c22e52b3bdee38d6101c08501527f0a267f746fcd38e9ac33aeb149abf39b535dbda7e794b8c69a8f69ea7a73577c6101e08501527f2778df932b1cc44eb25d1c9a7d74deff5c9ff0b76ce62fbe0c18b0dc2ea9f1396102008501527f2ae07fb6da3db6847939ecc3b7c3dd3b402e93ecee7015d927905173eb8da2a36102208501526102408401526102608301527f0a2ab642c246fe428aea69a87694e6bd0fa93cb9caa77df30eb815e4f8fe7cb56102808301527e6f3b1d2af841df426306e617b59f1c5e203dffedb3f74aafb154d6c4d2226f6102a08301527f12231bc6cb7bb13fc5b0facddff744e0ef5aa262b6317bb35cad358931a656356102c08301527f2fde75bb8f5f3fad412f755b522b58640d9f018a6efb1c6aaefa621d9848a9646102e0830152816103008160085afa9051161561152c57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191859101610ebf565b634e487b7160e01b5f52604160045260245ffd5b6351d49ff760e11b5f5260045ffd5b34610ba5575f366003190112610ba557807f503d45a44f156f8aad8d06dba42bfe0a05e75455578ef07ba15d10e0bd7ac2dc60209252f35b9181601f84011215610ba55782359167ffffffffffffffff8311610ba55760208381860195010111610ba557565b90601f8019910116810190811067ffffffffffffffff82111761155a57604052565b9091604092825260208201520160208251919201905f5b8181106116295750505090565b825184526020938401939092019160010161161c565b905f905b6008821061165057505050565b6020806001928551815201930191019091611643565b905f905b6002821061167757505050565b602080600192855181520193019101909161166a565b905f516020611d8b5f395f51905f528210801590611717575b61152c5781158061170f575b611709576116d65f516020611d8b5f395f51905f5260038185818180090908611bab565b8181036116e557505060011b90565b5f516020611d8b5f395f51905f52809106810306145f1461152c57600190811b1790565b50505f90565b5080156116b2565b505f516020611d8b5f395f51905f528110156116a6565b919093925f516020611d8b5f395f51905f52831080159061194b575b8015611934575b801561191d575b61152c578082868517171715611912579082916118755f516020611d8b5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611d8b5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161184f81808b80098187800908611bab565b8408095f516020611d8b5f395f51905f5261186982611d22565b80091415958691611bce565b929080821480611909575b156118a75750505050905f1461189f5760ff60025b169060021b179190565b60ff5f611895565b5f516020611d8b5f395f51905f528091068103061491826118ea575b50501561152c57600191156118e25760ff60025b169060021b17179190565b60ff5f6118d7565b5f516020611d8b5f395f51905f52919250819006810306145f806118c3565b50838314611880565b50505090505f905f90565b505f516020611d8b5f395f51905f52811015611758565b505f516020611d8b5f395f51905f52821015611751565b505f516020611d8b5f395f51905f5285101561174a565b9080601f83011215610ba5576040519161197e610100846115e3565b82906101008101928311610ba557905b82821061199b5750505090565b813581526020918201910161198e565b9080601f83011215610ba5576040805192906119c790846115e3565b829060408101928311610ba557905b8282106119e35750505090565b81358152602091820191016119d6565b60405190611a026020836115e3565b5f8252565b8015611a6b578060011c915f516020611d8b5f395f51905f5283101561152c57600180611a4a5f516020611d8b5f395f51905f5260038188818180090908611bab565b931614611a5357565b905f516020611d8b5f395f51905f5280910681030690565b505f905f90565b801580611ba3575b611b97578060021c92825f516020611d8b5f395f51905f528510801590611b80575b61152c5784815f516020611d8b5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4481611b4a9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611bce565b80929160018082961614611b5c575050565b5f516020611d8b5f395f51905f528093945080929550809106810306930681030690565b505f516020611d8b5f395f51905f52811015611a9c565b50505f905f905f905f90565b508115611a7a565b90611bb582611d22565b915f516020611d8b5f395f51905f528380090361152c57565b915f516020611d8b5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c2693969496611c1882808a8009818a800908611bab565b90611d16575b860809611bab565b925f516020611d8b5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611d8b5f395f51905f5260a083015260208260c08160055afa9151911561152c575f516020611d8b5f395f51905f5282600192090361152c575f516020611d8b5f395f51905f52908209925f516020611d8b5f395f51905f528080808780090681030681878009081490811591611cf7575b5061152c57565b90505f516020611d8b5f395f51905f528084860960020914155f611cf0565b81809106810306611c1e565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611d8b5f395f51905f5260a083015260208260c08160055afa9151911561152c5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220426a57ce43b157734930046782e09ca9100dc0aaf61bb40993af3bbc8f8e4c8864736f6c634300081c0033",
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
