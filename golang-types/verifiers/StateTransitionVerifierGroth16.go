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
	Bin: "0x60808060405234601557611dfe908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461157b575080639ed57f3214610cb9578063b1c3a00e14610ba8578063b8e72af614610a315763e2dd9f4714610055575f80fd5b34610a2e576101c0366003190112610a2e5736608411610a2e573660a411610a2e57366101c411610a2e57604051602061008f81836115e1565b803683376040918251926100a381856115e1565b80368537610300938151926100b886856115e1565b853685376100c7608435611a05565b8684015282526100d860a435611a05565b905f516020611da95f395f51905f5285516100f387826115e1565b6001815288610104818301601f198a01368237378551906101328a610124818a0151938b5194859384019687611603565b03601f1981018352826115e1565b5190200683528351865286840151878701527f21b2ab490951a680218ed69eadb0d93956ef17090ee96a53f29fcf1fc2c3e38e858701527f1c90797f06c7b56d0117065d76cc1e05f386f44afff8b05af24ee4fc0dfd9ae160608701527f28e4e6168112d2d3290281416d9bc4377b9226593c25fb1ba9886dc0038877a060808701527f2746167425c67a5cb3f8a5e200988360b739f3d18d91c35aabf4f0e3e775dea360a087015260c086015260e08501527e621c2a2edeef44e908990c402c25b8bcf1818b530afb91ad4b08c819b6dd126101008501527f24724f3331e9da6ad69c7e2efdee71d31791c9b37bedd2625d6cb593f2fe51b06101208501527f2cedaeb50afc4317585964202b9b27508f605e77db1704637321fdbdf7f0b3c06101408501527f041b2172102d1e575d86f51c9e89b5ef09ef0339b4efb1c30a82e02c6c4e83fa610160850152825185816101808760085afa90511615610a1f578490876102a2600435611a05565b866102b4602497939735604435611a70565b9491929390956102c5606435611a05565b999098507e7d9e1cdfc49efc8da8f1d624e4aa8765c274054cafa2124e9df40dfad4444b83519b8c937f0b2b4b06d576f194660299caa9e5a38c0071958fcc0367c7e6cd03379842cc2c85527f13cc9e79ba2cbbfe18b197db38e5301a59b9086cc9f344e640b6d6f47a9156a4828601528051868601520151606084019081527f1bafbb974b76979abd38f8eb79d90bec4b7b19abfeab4b93415062ede06f94df858560808160065afa957f1d2e3d6160d499a4c5899baa135a8e5889c2d6c8ced309e153936ff3263c8080818701527f231336c274fee0ba6796a305270875ab7122620fc424a613273a58da1d15868d8352600160c4356080880198818a525f516020611da95f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f0f34e6ef22cc248de40ea8fd79837ec3350477c1e2f9a8daf8bbe5fadd77e02d828801527f07b19027776625ba289a433a028ce73173128085eecda4a8dd938a5ffb1739d6845260e435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f021fff26816e8ba78d14ef51adea5359998cc72d5cf584d2b1d125e9e24164b0828801527f09b8cb3f6e86acf05da89be0d1421f05ca97938c35b51be23f4f4146929d0bd4845261010435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f20bee48b92be7d802fa5530d8e7cd2ee453967a132847e3aaf19591c1319aacf828801527f21d1fb4037e965ec7fba85ad672cab071224e77dac2fef27e66e508bac2cf7e0845261012435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1f5d9e93b499ec2ecd769c14e6e0389120a72aab7d3d1ad5c0724e5cf698c3bf828801527f1616ed3fbae614f90b5b66e59616573c896e241bbfb47fe376f2788b28fa7443845261014435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2547eebcc3dc489ef504791e92df090eedc05bb186e0990e2f92e4116d41d485828801527f2a7e1661b4eb1bc509756cc3cd167f20ab72b63e349a9a7d2cc0283d57ccdeba845261016435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0459c1a9453aa6ed9ee445dc4c2f3db3f90581163f84ac0de2847aa7a9f4c87a828801527f07c78bb83d0d899cd952724065482cd4c1d371f774da8ec785eea2c3d296ecd3845261018435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f10a9137e39e9b302c9498600f7461e567c5acd20914cf83fff2ffd94fa51798b828801527f23e9fb78ca1e198b7ebda70311acace9c6a6c74cb81489811eedce0b55172cc584526101a435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611da95f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610a10578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f0dc5dad3c1defa3cb29182f4ab4677587eb6c4438c87737a7a76980fba22c00e6101008501527f2ed56587ddec52e7b068e4a1f04d0328f36c731c45eb5671f811280df4148cda6101208501527f237376c54f192cff33fc74edfcebddcb198790a34879bb1823a1d593ad1b867e6101408501527f20a528a1ef1b2f3fe86ff7144703fe8f6eaaff301a5cb868c50945d6da303eb26101608501527f08aec0ccb99d27f6a21d6022210611dcbb8b3894c8d06dc6184e5bf38df38bcd6101808501527f290cf47b0ca7d5ec89485841e1aa4321c94e44d1c5f48a3ce38fbd06c8b220f86101a08501527f179cf6368f8c16be829f9d4a0dee5cd53aaff398c6303b2f2943b16c50efdbd76101c08501527f243f20531c949138f4c45e37672b998d0df8702656566c507297ab7c0026a3996101e08501527f2723165bc08776a8be83f08c55fed45b0581c5258dee289fadb397aa2d09b3ec6102008501527f063b4f27fddd8d27a18e7d7d4c691d8d8248bc1a3a5f9fa530d2293bd1598ab46102208501526102408401526102608301527f1f30819bf3d0d9151334c46c109a9f58035343d3b363cea41d782930a2ddf1786102808301527f17239b8c705deb2c5455cacaee468bcf3ecd951a1459b8d531e40e3a6de6a1326102a08301527f27af660252f751bccf3f46815a5974767927b5901e2c66275f154c4d3e5468e96102c08301527f19b283aefc8e3bb529d08a4ea2602fee23c3d4712261420784d8b8e152bd5dbb6102e0830152519283916109db84846115e1565b8336843760085afa15908115610a03575b506109f45780f35b631ff3747d60e21b8152600490fd5b600191505114155f6109ec565b63a54f8e2760e01b8f5260048ffd5b6351d49ff760e11b8752600487fd5b80fd5b5034610ba4576040366003190112610ba45760043567ffffffffffffffff8111610ba457610a639036906004016115b3565b60243567ffffffffffffffff8111610ba457610a839036906004016115b3565b90926101009384604051610a9782826115e1565b369037604093848051610aaa82826115e1565b369037848051610aba82826115e1565b36903782019461018083870312610ba457610aef610ad88785611960565b96610140610ae8828588016119a9565b95016119a9565b93818651610afd82826115e1565b36903782019082820312610ba457610b1491611960565b90610b1d6119f1565b50610b266119f1565b50303b15610ba457610b7492610b5e610b6992610b53875198634f6abf9960e11b8a5260048a019061163d565b610104880190611664565b610144860190611664565b61018484019061163d565b5f8261028481305afa908115610b9b5750610b8d575080f35b610b9991505f906115e1565b005b513d5f823e3d90fd5b5f80fd5b34610ba457610180366003190112610ba4573661010411610ba4573661014411610ba4573661018411610ba457604051610be36080826115e1565b6080368237604051906020610bf881846115e1565b80368437610c0a60243560043561168b565b8252610c2060843560a43560443560643561172c565b828401526040830152610c3760e43560c43561168b565b6060830152610c4c610124356101043561168b565b8352610c5e610164356101443561168b565b9060405192835f905b60048210610ca3575050506080830193905f945b60018610610c8e5760c0858560a0820152f35b81806001928551815201930195019491610c7b565b8251815291830191600191909101908301610c67565b34610ba457610280366003190112610ba4573661010411610ba4573661014411610ba4573661018411610ba4573661028411610ba4576040516020610cfe81836115e1565b803683376040915f516020611da95f395f51905f528351610d1f85826115e1565b60018152836101c4818301601f198801368237378451610d4e8161012487820194610124356101043587611603565b51902006815282518361010482377f21b2ab490951a680218ed69eadb0d93956ef17090ee96a53f29fcf1fc2c3e38e848201527f1c90797f06c7b56d0117065d76cc1e05f386f44afff8b05af24ee4fc0dfd9ae160608201527f28e4e6168112d2d3290281416d9bc4377b9226593c25fb1ba9886dc0038877a060808201527f2746167425c67a5cb3f8a5e200988360b739f3d18d91c35aabf4f0e3e775dea360a08201528361014460c08301377e621c2a2edeef44e908990c402c25b8bcf1818b530afb91ad4b08c819b6dd126101008201527f24724f3331e9da6ad69c7e2efdee71d31791c9b37bedd2625d6cb593f2fe51b06101208201527f2cedaeb50afc4317585964202b9b27508f605e77db1704637321fdbdf7f0b3c06101408201527f041b2172102d1e575d86f51c9e89b5ef09ef0339b4efb1c30a82e02c6c4e83fa61016082015282816101808160085afa9051161561156c57825192610eb76040856115e1565b61010483855b6101448310611548575050508051918183017f0b2b4b06d576f194660299caa9e5a38c0071958fcc0367c7e6cd03379842cc2c84525f516020611da95f395f51905f528386808701987f13cc9e79ba2cbbfe18b197db38e5301a59b9086cc9f344e640b6d6f47a9156a48a528051855201519260608701938452818760808160065afa947f1d2e3d6160d499a4c5899baa135a8e5889c2d6c8ced309e153936ff3263c808082527f231336c274fee0ba6796a305270875ab7122620fc424a613273a58da1d15868d85527e7d9e1cdfc49efc8da8f1d624e4aa8765c274054cafa2124e9df40dfad4444b60016101843560808b0198818a5287878760608160075afa92101616858b60808160065afa16167f0f34e6ef22cc248de40ea8fd79837ec3350477c1e2f9a8daf8bbe5fadd77e02d84527f07b19027776625ba289a433a028ce73173128085eecda4a8dd938a5ffb1739d687526101a4359081895286868660608160075afa92101616848a60808160065afa167f021fff26816e8ba78d14ef51adea5359998cc72d5cf584d2b1d125e9e24164b084527f09b8cb3f6e86acf05da89be0d1421f05ca97938c35b51be23f4f4146929d0bd487526101c4359081895286868660608160075afa92101616848a60808160065afa167f20bee48b92be7d802fa5530d8e7cd2ee453967a132847e3aaf19591c1319aacf84527f21d1fb4037e965ec7fba85ad672cab071224e77dac2fef27e66e508bac2cf7e087526101e4359081895286868660608160075afa92101616848a60808160065afa167f1f5d9e93b499ec2ecd769c14e6e0389120a72aab7d3d1ad5c0724e5cf698c3bf84527f1616ed3fbae614f90b5b66e59616573c896e241bbfb47fe376f2788b28fa74438752610204359081895286868660608160075afa92101616848a60808160065afa167f2547eebcc3dc489ef504791e92df090eedc05bb186e0990e2f92e4116d41d48584527f2a7e1661b4eb1bc509756cc3cd167f20ab72b63e349a9a7d2cc0283d57ccdeba8752610224359081895286868660608160075afa92101616848a60808160065afa167f0459c1a9453aa6ed9ee445dc4c2f3db3f90581163f84ac0de2847aa7a9f4c87a84527f07c78bb83d0d899cd952724065482cd4c1d371f774da8ec785eea2c3d296ecd38752610244359081895286868660608160075afa92101616848a60808160065afa167f10a9137e39e9b302c9498600f7461e567c5acd20914cf83fff2ffd94fa51798b84527f23e9fb78ca1e198b7ebda70311acace9c6a6c74cb81489811eedce0b55172cc58752610264359081895286868660608160075afa92101616848a60808160065afa16957f1bafbb974b76979abd38f8eb79d90bec4b7b19abfeab4b93415062ede06f94df8452525180955260608160075afa92101616818360808160065afa16915193519115611539575192610100600485377f0dc5dad3c1defa3cb29182f4ab4677587eb6c4438c87737a7a76980fba22c00e6101008501527f2ed56587ddec52e7b068e4a1f04d0328f36c731c45eb5671f811280df4148cda6101208501527f237376c54f192cff33fc74edfcebddcb198790a34879bb1823a1d593ad1b867e6101408501527f20a528a1ef1b2f3fe86ff7144703fe8f6eaaff301a5cb868c50945d6da303eb26101608501527f08aec0ccb99d27f6a21d6022210611dcbb8b3894c8d06dc6184e5bf38df38bcd6101808501527f290cf47b0ca7d5ec89485841e1aa4321c94e44d1c5f48a3ce38fbd06c8b220f86101a08501527f179cf6368f8c16be829f9d4a0dee5cd53aaff398c6303b2f2943b16c50efdbd76101c08501527f243f20531c949138f4c45e37672b998d0df8702656566c507297ab7c0026a3996101e08501527f2723165bc08776a8be83f08c55fed45b0581c5258dee289fadb397aa2d09b3ec6102008501527f063b4f27fddd8d27a18e7d7d4c691d8d8248bc1a3a5f9fa530d2293bd1598ab46102208501526102408401526102608301527f1f30819bf3d0d9151334c46c109a9f58035343d3b363cea41d782930a2ddf1786102808301527f17239b8c705deb2c5455cacaee468bcf3ecd951a1459b8d531e40e3a6de6a1326102a08301527f27af660252f751bccf3f46815a5974767927b5901e2c66275f154c4d3e5468e96102c08301527f19b283aefc8e3bb529d08a4ea2602fee23c3d4712261420784d8b8e152bd5dbb6102e0830152816103008160085afa9051161561152a57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191859101610ebd565b634e487b7160e01b5f52604160045260245ffd5b6351d49ff760e11b5f5260045ffd5b34610ba4575f366003190112610ba457807f38bcaf9462ec80f8dde0b7b2086693acad293f371c70dd0ca92f0bb7f6809c4060209252f35b9181601f84011215610ba45782359167ffffffffffffffff8311610ba45760208381860195010111610ba457565b90601f8019910116810190811067ffffffffffffffff82111761155857604052565b9091604092825260208201520160208251919201905f5b8181106116275750505090565b825184526020938401939092019160010161161a565b905f905b6008821061164e57505050565b6020806001928551815201930191019091611641565b905f905b6002821061167557505050565b6020806001928551815201930191019091611668565b905f516020611d895f395f51905f528210801590611715575b61152a5781158061170d575b611707576116d45f516020611d895f395f51905f5260038185818180090908611ba9565b8181036116e357505060011b90565b5f516020611d895f395f51905f52809106810306145f1461152a57600190811b1790565b50505f90565b5080156116b0565b505f516020611d895f395f51905f528110156116a4565b919093925f516020611d895f395f51905f528310801590611949575b8015611932575b801561191b575b61152a578082868517171715611910579082916118735f516020611d895f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611d895f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161184d81808b80098187800908611ba9565b8408095f516020611d895f395f51905f5261186782611d20565b80091415958691611bcc565b929080821480611907575b156118a55750505050905f1461189d5760ff60025b169060021b179190565b60ff5f611893565b5f516020611d895f395f51905f528091068103061491826118e8575b50501561152a57600191156118e05760ff60025b169060021b17179190565b60ff5f6118d5565b5f516020611d895f395f51905f52919250819006810306145f806118c1565b5083831461187e565b50505090505f905f90565b505f516020611d895f395f51905f52811015611756565b505f516020611d895f395f51905f5282101561174f565b505f516020611d895f395f51905f52851015611748565b9080601f83011215610ba4576040519161197c610100846115e1565b82906101008101928311610ba457905b8282106119995750505090565b813581526020918201910161198c565b9080601f83011215610ba4576040805192906119c590846115e1565b829060408101928311610ba457905b8282106119e15750505090565b81358152602091820191016119d4565b60405190611a006020836115e1565b5f8252565b8015611a69578060011c915f516020611d895f395f51905f5283101561152a57600180611a485f516020611d895f395f51905f5260038188818180090908611ba9565b931614611a5157565b905f516020611d895f395f51905f5280910681030690565b505f905f90565b801580611ba1575b611b95578060021c92825f516020611d895f395f51905f528510801590611b7e575b61152a5784815f516020611d895f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4481611b489d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611bcc565b80929160018082961614611b5a575050565b5f516020611d895f395f51905f528093945080929550809106810306930681030690565b505f516020611d895f395f51905f52811015611a9a565b50505f905f905f905f90565b508115611a78565b90611bb382611d20565b915f516020611d895f395f51905f528380090361152a57565b915f516020611d895f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c2493969496611c1682808a8009818a800908611ba9565b90611d14575b860809611ba9565b925f516020611d895f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611d895f395f51905f5260a083015260208260c08160055afa9151911561152a575f516020611d895f395f51905f5282600192090361152a575f516020611d895f395f51905f52908209925f516020611d895f395f51905f528080808780090681030681878009081490811591611cf5575b5061152a57565b90505f516020611d895f395f51905f528084860960020914155f611cee565b81809106810306611c1c565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611d895f395f51905f5260a083015260208260c08160055afa9151911561152a5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220ff8e68d4320f1d589f90fa56581d33854159ab21f0b57023f038306159d5d13964736f6c634300081c0033",
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
