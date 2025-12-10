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
	Bin: "0x60808060405234601557611e77908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461167a575080635d26278e14610c8b57806360e583461461037f578063b1c3a00e1461026e5763b8e72af614610055575f80fd5b346102365760403660031901126102365760043567ffffffffffffffff8111610236576100869036906004016116b2565b60243567ffffffffffffffff8111610236576100a69036906004016116b2565b61010093846040516100b882826116e0565b3690376040938480516100cb82826116e0565b3690378480516100db82826116e0565b3690378101610180828203126102365780601f8301121561023657845195610105610100886116e0565b8201868282116102365783905b82821061025e5750509061014061012c8261013394611bda565b9301611bda565b916101209081865161014582826116e0565b36903784019080858303126102365781601f860112156102365785519461016e610120876116e0565b8591810192831161023657905b82821061023a57505050303b15610236578351633072c1a360e11b8152945f600487015b6008821061022057505050906101bd6101c89261010487019061173a565b61014485019061173a565b5f61018484015b6009821061020a575050505f826102a481305afa90811561020157506101f3575080f35b6101ff91505f906116e0565b005b513d5f823e3d90fd5b60208060019285518152019301910190916101cf565b602080600192855181520193019101909161019f565b5f80fd5b813581526020918201910161017b565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610112565b3461023657610180366003190112610236573661010411610236573661014411610236573661018411610236576040516102a96080826116e0565b60803682376040519060206102be81846116e0565b803684376102d0602435600435611905565b82526102e660843560a4356044356064356119a6565b8284015260408301526102fd60e43560c435611905565b60608301526103126101243561010435611905565b83526103246101643561014435611905565b9060405192835f905b60048210610369575050506080830193905f945b600186106103545760c0858560a0820152f35b81806001928551815201930195019491610341565b825181529183019160019190910190830161032d565b34610236576102a036600319011261023657366101041161023657366101441161023657366101841161023657366102a4116102365760405160206103c481836116e0565b803683375f516020611e225f395f51905f5260405182810190610400816103f2610124356101043586611702565b03601f1981018352826116e0565b519020068252604051604061010482377f1018ce80e4a47ad132eef29aedad7e8041468258af65c2cef547caeffa4cd83660408201527f218054bf0d9b7dc3b5199a4bf94b94ab2dcb725490fb5e6290a95886e600d1c460608201527f06f66718dfb94fc45a0caace18a77655125b2afce631fe64a1da993dd771ddc760808201527f2a72428d1884469a43e265c8a533d736d7bd81d4914b7ee684a69398d0766b7260a0820152604061014460c08301377f2c5fd7d3d25b66b3c95ff1a499004323521abda847feaf82d44a3b1586b026dc6101008201527f0641e168a1e38b2a46d3020753dd9c1d73523b2dc271dcaf364bad151c1d5c3b6101208201527ed44d66b01dd16358f3615ea8c40b4368dd94b7dd13e7af7470662f411d26736101408201527f241da06f6083a33cf3a3a127e0719825e1813f34d209b7a3d49cc72e3d373e0361016082015281816101808160085afa90511615610c7c5760408051929061056f90846116e0565b61010482845b6101448310610c6c5750505060405190604082017f19652b20df436d75efeaacf8e130d9245b109383e3514481759bb706283392ee83525f516020611e225f395f51905f52604085808601977f1b7349b995fdb19d77c4379c08566d05495064c0741c623b174c1c47fce9dae489528051855201519260608601938452818660808160065afa947f1094b1f2cb247b270095233ac46fc949e0840288a7b3f2149ffb0050496e207382527f2c411bfc92499e667658661a69816c3d4a5b397969aa3302002b265982e2076185527f0c262194d7fd4887f0aaad1b6db4ee003a8fa3636ed46dd264b98afb2b0996a960016101843560808a0198818a5287878760608160075afa92101616858a60808160065afa16167f1a6ed5b8aa05e5d119798e66e324c60d983790d8648e67a1abc5e7895c84101b84527f205294af93a5f1bdc6e4258743470bc8878feaaf7de6ce6b92755b77a752a0f387526101a4359081895286868660608160075afa92101616848960808160065afa167f055146318abb3cdf8757df6f77306214997161a0dde308ccf7f993351692c21e84527f2ab14e5ca24a3ae5474a9eda06784c793e2d812b281403f8ee2f503a9ed4da1a87526101c4359081895286868660608160075afa92101616848960808160065afa167f03fdf1eca536b87ec65c0e050603d0ebead282fa4ee357c6813e6fbf1eee243284527f1a0f89d9bc7ee6f0ac2c7c2d92f040ff4ca96ae175e0dccf9193f705a6ce8ef387526101e4359081895286868660608160075afa92101616848960808160065afa167f04de319bf65dbe7da31d6dcca2668b60ec369133f746efee56f3bdd3519428e284527f242ad5f856a3e7640295d2939472ec8160efd1765e76b09bc04f38f4b33230648752610204359081895286868660608160075afa92101616848960808160065afa167f0c29c85eed3353ad191e841294306ecfbaaa1834efea0e2380d6ae2d47da8d3984527f23c7b8ebdc147384e9e01c62662e63de10fb95a5755f193d3f6a868372bdffac8752610224359081895286868660608160075afa92101616848960808160065afa167f0b75bb2724c4563d36eec772e78bf197aff42d057f0b95ed7cfd8a5ebb75f4f384527f028c964cfc9c2215b46a78ff70df652ef2a1a144339bb93bd7814138f73db4028752610244359081895286868660608160075afa92101616848960808160065afa167f0ccfd62b121396090dd45f3a8f6a9557987be48a3381a7740a45f4aabe06835b84527f03055784ebbf1f76f017939ee8f224c5fc745dc986c4426cef614e65ec74a7de8752610264359081895286868660608160075afa92101616848960808160065afa167f0e5d84a1da7db22651e26056eda29e0dd4ff6d9b04ff135a91fe6f82049f4b8684527f0ca6e7a0dc784e7c1d9fd54c443bedd176d7bfe8da810d35218278f9e7862a248752610284359081895286868660608160075afa92101616848960808160065afa16957f297609b60d930f563e8e58f1f5888b72fec09b1ea432276618ee7d15aee413e68452525180955260608160075afa9210161660408260808160065afa16905192519015610c5d5760405192610100600485377f2f181d33c17dfc6023989cf6d65a174a77b2a880955f9021d405e02b456c2e576101008501527f1a4be67ab666b1a09f5730f7c8dffc6a065f7b320fb56e3bf029c04f398625c26101208501527f2cc67d09e3c74a8685a490ed79abe050e3600a5edcd680f3f185956ceac61ce06101408501527f1e0afe2df114db61b08f10ca00dc02a4f01f8ac5485d8b1fd71b77f02defbf4b6101608501527f2a35fdcb3f6e8218229e06bf46e44be46ecf8465bc52acd6701cadca6c781a976101808501527f27eb5b4569eb95e5300d72086e5df8ab10bdeb11a5a9c9e1aac414ad5adc04756101a08501527e4cf95d8cf61bd3b5490c929960bb39afddcdd06a4784c4621acc7cfca695d06101c08501527f1bc799382455d5e6493246a5a22a5688ad0072640764cc17dc8a1d642d83afb46101e08501527f0caf346721bc5738deb9436dba2f2c7c4c2e0a52a3691ddb82b733bfbd888bbd6102008501527f2666c401d6b843983a92a864972f2e4df33777389686e40399c2b1a45514592c6102208501526102408401526102608301527f290442bec8261100de9c35e4d301a289ddd78a3f40e34d5a0d4bf313f7cf48846102808301527f0400a7770c5682521f2bdae030d0fc7f9eab4489a3bdc3c2130d0cbdf8f386646102a08301527f3023784f8620b6ca7cbe233be10458711b09b2a60a0879df82acbb784bb704486102c08301527f1c31525fcd312a76086fdd2c9ac0a6ff57ed3d0a74618ed5bcb8954fbba512b76102e0830152816103008160085afa90511615610c4e57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191849101610575565b6351d49ff760e11b5f5260045ffd5b34610236576101e03660031901126102365736608411610236573660a41161023657366101e411610236576040516020610cc581836116e0565b803683376040805190610cd881836116e0565b8036833761030093815192610ced86856116e0565b85368537610cfc608435611761565b868301528152610d0d60a435611761565b905f516020611e225f395f51905f5283516103f2610d378a87015189519283918d83019586611702565b5190200684528251865286830151878701527f1018ce80e4a47ad132eef29aedad7e8041468258af65c2cef547caeffa4cd836858701527f218054bf0d9b7dc3b5199a4bf94b94ab2dcb725490fb5e6290a95886e600d1c460608701527f06f66718dfb94fc45a0caace18a77655125b2afce631fe64a1da993dd771ddc760808701527f2a72428d1884469a43e265c8a533d736d7bd81d4914b7ee684a69398d0766b7260a087015260c086015260e08501527f2c5fd7d3d25b66b3c95ff1a499004323521abda847feaf82d44a3b1586b026dc6101008501527f0641e168a1e38b2a46d3020753dd9c1d73523b2dc271dcaf364bad151c1d5c3b6101208501527ed44d66b01dd16358f3615ea8c40b4368dd94b7dd13e7af7470662f411d26736101408501527f241da06f6083a33cf3a3a127e0719825e1813f34d209b7a3d49cc72e3d373e03610160850152825185816101808760085afa90511615610c7c578490610ea6600435611761565b610eb76024959295356044356117cc565b9291909389610ec7606435611761565b9890977f0c262194d7fd4887f0aaad1b6db4ee003a8fa3636ed46dd264b98afb2b0996a983519b8c937f19652b20df436d75efeaacf8e130d9245b109383e3514481759bb706283392ee85527f1b7349b995fdb19d77c4379c08566d05495064c0741c623b174c1c47fce9dae4828601528051868601520151606084019081527f297609b60d930f563e8e58f1f5888b72fec09b1ea432276618ee7d15aee413e6858560808160065afa957f1094b1f2cb247b270095233ac46fc949e0840288a7b3f2149ffb0050496e2073818701527f2c411bfc92499e667658661a69816c3d4a5b397969aa3302002b265982e207618352600160c4356080880198818a525f516020611e225f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f1a6ed5b8aa05e5d119798e66e324c60d983790d8648e67a1abc5e7895c84101b828801527f205294af93a5f1bdc6e4258743470bc8878feaaf7de6ce6b92755b77a752a0f3845260e435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f055146318abb3cdf8757df6f77306214997161a0dde308ccf7f993351692c21e828801527f2ab14e5ca24a3ae5474a9eda06784c793e2d812b281403f8ee2f503a9ed4da1a845261010435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f03fdf1eca536b87ec65c0e050603d0ebead282fa4ee357c6813e6fbf1eee2432828801527f1a0f89d9bc7ee6f0ac2c7c2d92f040ff4ca96ae175e0dccf9193f705a6ce8ef3845261012435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f04de319bf65dbe7da31d6dcca2668b60ec369133f746efee56f3bdd3519428e2828801527f242ad5f856a3e7640295d2939472ec8160efd1765e76b09bc04f38f4b3323064845261014435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0c29c85eed3353ad191e841294306ecfbaaa1834efea0e2380d6ae2d47da8d39828801527f23c7b8ebdc147384e9e01c62662e63de10fb95a5755f193d3f6a868372bdffac845261016435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0b75bb2724c4563d36eec772e78bf197aff42d057f0b95ed7cfd8a5ebb75f4f3828801527f028c964cfc9c2215b46a78ff70df652ef2a1a144339bb93bd7814138f73db402845261018435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0ccfd62b121396090dd45f3a8f6a9557987be48a3381a7740a45f4aabe06835b828801527f03055784ebbf1f76f017939ee8f224c5fc745dc986c4426cef614e65ec74a7de84526101a435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0e5d84a1da7db22651e26056eda29e0dd4ff6d9b04ff135a91fe6f82049f4b86828801527f0ca6e7a0dc784e7c1d9fd54c443bedd176d7bfe8da810d35218278f9e7862a2484526101c435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611e225f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610c5d578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f2f181d33c17dfc6023989cf6d65a174a77b2a880955f9021d405e02b456c2e576101008501527f1a4be67ab666b1a09f5730f7c8dffc6a065f7b320fb56e3bf029c04f398625c26101208501527f2cc67d09e3c74a8685a490ed79abe050e3600a5edcd680f3f185956ceac61ce06101408501527f1e0afe2df114db61b08f10ca00dc02a4f01f8ac5485d8b1fd71b77f02defbf4b6101608501527f2a35fdcb3f6e8218229e06bf46e44be46ecf8465bc52acd6701cadca6c781a976101808501527f27eb5b4569eb95e5300d72086e5df8ab10bdeb11a5a9c9e1aac414ad5adc04756101a08501527e4cf95d8cf61bd3b5490c929960bb39afddcdd06a4784c4621acc7cfca695d06101c08501527f1bc799382455d5e6493246a5a22a5688ad0072640764cc17dc8a1d642d83afb46101e08501527f0caf346721bc5738deb9436dba2f2c7c4c2e0a52a3691ddb82b733bfbd888bbd6102008501527f2666c401d6b843983a92a864972f2e4df33777389686e40399c2b1a45514592c6102208501526102408401526102608301527f290442bec8261100de9c35e4d301a289ddd78a3f40e34d5a0d4bf313f7cf48846102808301527f0400a7770c5682521f2bdae030d0fc7f9eab4489a3bdc3c2130d0cbdf8f386646102a08301527f3023784f8620b6ca7cbe233be10458711b09b2a60a0879df82acbb784bb704486102c08301527f1c31525fcd312a76086fdd2c9ac0a6ff57ed3d0a74618ed5bcb8954fbba512b76102e08301525192839161165584846116e0565b8336843760085afa1590811561166d575b50610c4e57005b6001915051141581611666565b34610236575f36600319011261023657807f6e2fe44cbdb5b75c383e4d7a2e14dd89fe0b8771ff0bfa2edb9518730a11fe9660209252f35b9181601f840112156102365782359167ffffffffffffffff8311610236576020838186019501011161023657565b90601f8019910116810190811067ffffffffffffffff82111761024a57604052565b909160409282526020820152016060516080905f5b8181106117245750505090565b8251845260209384019390920191600101611717565b905f905b6002821061174b57505050565b602080600192855181520193019101909161173e565b80156117c5578060011c915f516020611e025f395f51905f52831015610c4e576001806117a45f516020611e025f395f51905f5260038188818180090908611c22565b9316146117ad57565b905f516020611e025f395f51905f5280910681030690565b505f905f90565b8015806118fd575b6118f1578060021c92825f516020611e025f395f51905f5285108015906118da575b610c4e5784815f516020611e025f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816118a49d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611c45565b809291600180829616146118b6575050565b5f516020611e025f395f51905f528093945080929550809106810306930681030690565b505f516020611e025f395f51905f528110156117f6565b50505f905f905f905f90565b5081156117d4565b905f516020611e025f395f51905f52821080159061198f575b610c4e57811580611987575b6119815761194e5f516020611e025f395f51905f5260038185818180090908611c22565b81810361195d57505060011b90565b5f516020611e025f395f51905f52809106810306145f14610c4e57600190811b1790565b50505f90565b50801561192a565b505f516020611e025f395f51905f5281101561191e565b919093925f516020611e025f395f51905f528310801590611bc3575b8015611bac575b8015611b95575b610c4e578082868517171715611b8a57908291611aed5f516020611e025f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611e025f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611ac781808b80098187800908611c22565b8408095f516020611e025f395f51905f52611ae182611d99565b80091415958691611c45565b929080821480611b81575b15611b1f5750505050905f14611b175760ff60025b169060021b179190565b60ff5f611b0d565b5f516020611e025f395f51905f52809106810306149182611b62575b505015610c4e5760019115611b5a5760ff60025b169060021b17179190565b60ff5f611b4f565b5f516020611e025f395f51905f52919250819006810306145f80611b3b565b50838314611af8565b50505090505f905f90565b505f516020611e025f395f51905f528110156119d0565b505f516020611e025f395f51905f528210156119c9565b505f516020611e025f395f51905f528510156119c2565b9080601f8301121561023657604080519290611bf690846116e0565b82906040810192831161023657905b828210611c125750505090565b8135815260209182019101611c05565b90611c2c82611d99565b915f516020611e025f395f51905f5283800903610c4e57565b915f516020611e025f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c9d93969496611c8f82808a8009818a800908611c22565b90611d8d575b860809611c22565b925f516020611e025f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611e025f395f51905f5260a083015260208260c08160055afa91519115610c4e575f516020611e025f395f51905f52826001920903610c4e575f516020611e025f395f51905f52908209925f516020611e025f395f51905f528080808780090681030681878009081490811591611d6e575b50610c4e57565b90505f516020611e025f395f51905f528084860960020914155f611d67565b81809106810306611c95565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611e025f395f51905f5260a083015260208260c08160055afa91519115610c4e5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220f6bb875205d4fdab49d6ebcfef184d6bd57382fa8fa19b959518f7012db6891e64736f6c634300081c0033",
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
