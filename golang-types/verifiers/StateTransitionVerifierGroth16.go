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
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"},{\"internalType\":\"uint256[8]\",\"name\":\"input\",\"type\":\"uint256[8]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256[8]\",\"name\":\"input\",\"type\":\"uint256[8]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557611d53908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a391461147c578063233ace1114611442578063449ccd1e14610b91578063b8e72af614610a395763e2dd9f4714610051575f80fd5b34610a36576101c0366003190112610a365736608411610a36573660a411610a3657366101c411610a3657604051602061008b81836115df565b8036833760409182519261009f81856115df565b80368537610300938151926100b486856115df565b853685376100c360843561195a565b8684015282526100d460a43561195a565b905f516020611cfe5f395f51905f5285516100ef87826115df565b6001815288610104818301601f198a01368237378551906101388a61012a818a0151938b5194859384019687528c8401526060830190611658565b03601f1981018352826115df565b5190200683528351865286840151878701527f0edb2d5af6534a82d734f6c8c9599223561176c6e93148568279e1f726011e7c858701527f11ce9383c53d66901aa59bb23efa55a7d1566a4f8903e6295b5004a41c05da9760608701527f2978bd840f86c9fcf14202e8f2d168a9d2f2375c49eee3b39cd22d4318c62bd060808701527f2494f1da9eaa4decf1a5a2260dc62840b4e96cff8688d58ff56f8d8e3e5a3a3060a087015260c086015260e08501527f2b2fce9edc8be658347005d796bfc86ff82d0fc8daef0290c95953c1e5c9935c6101008501527f022521d8f81cf0c7cf119e45a05fdcc00055a32fa0e5e8aab713869c436b63ec6101208501527f199e42fbcbbda303f0d4a04a9cc2a418cc450d1f66520fa2cb552d261130d2d96101408501527f28d803a477737410df240d247636ca91f7171e28be001373083103488a434f6d610160850152825185816101808760085afa90511615610a27578490876102a960043561195a565b866102bb6024979397356044356119c5565b9491929390956102cc60643561195a565b999098507f1e1542f5affdeb02ee411d8358b8cb2f87e302c1c0a13f6923f9b89667a1518e83519b8c937f198afa3741fdf69bb460a8b9ef813b10b4713e3b2dc01da7de1ba4b6676b9e0f85527f1f6c2819b0794a2fd7faeece26052552e3be6f94ab970ee71297817b246f9d55828601528051868601520151606084019081527f3043c0969def31e6a445ce8eed7d2e084e4caa31d815bf0d6e834acd983e85a2858560808160065afa957f188c496ae019c47576f22616fc5e6cc08830b52d32a3eba8d9ac0cf7145fecdb818701527f297dec4673c8916d92ead3aa420b45e10a7c3e78be509ad3a00f2bf847f38f5b8352600160c4356080880198818a525f516020611cfe5f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f086afd4e275d21964198f3c8a7622e23cad7187af5278abb7b0a05f8983ffba6828801527f1c9e6f2de2b2f030473884570ffcd4a981d28997e456246280f2ce8b69a5ca9f845260e435908189525f516020611cfe5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f06fd9e9988b1d4eb30dc4f0565a8e7a2ac763dbe2397f53e6263a6fc0080eedb828801527f0305d3166d58c6988fffd2af2c15be234b1f44cacc4610037cba2b902086e71c845261010435908189525f516020611cfe5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1af4e0f4150ee350ce8165c3d3692113f6b20c5fd6f849b21a8f454c4a8fa01a828801527f21a1c0f657c9529df4c6818965b70f2d0d69c1ee07d73a457d29f7bd4edcdc7f845261012435908189525f516020611cfe5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2b5cc1b293be91d5307c6ce964f2520964f72bce116d253ed7d22dc2587adab1828801527f04149627bee7c4d42ff97a65159071a8dc2b5786490de5f402d6bdb737a94db7845261014435908189525f516020611cfe5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f02309bb2c7a547fe72ec3c407187dd213e90d7a3d46fe0381e5cfec7305a3e87828801527f1d8db0fca5701759b556f20375f3f1fd1e78003cd2d57dbf51dbec77b76b39bd845261016435908189525f516020611cfe5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f190faa9e7ecebb3fcc7e48fd3fb33b8d47e310136ac4fbb2989dd0c5a0f92837828801527f215885b35016aadad09311ee0eba9d67c7f03fec32b2576b4f4cc14c8a219348845261018435908189525f516020611cfe5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f28ff412be210b4fdac69993ba622c67dbd9f727b691c63325cfb3fe788c98748828801527f157781cf5696de1b56e40ea61d9fa14ecfa96a87bdfeb249dfebd2964a34974884526101a435908189525f516020611cfe5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611cfe5f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610a18578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f1df2f9b54952a25c5d792ebe7c3f05cdcc7dcda10ed263be9d1e54fe6afbf7936101008501527f013469a889baac39f11aa45e09ca2f00a124360f949d6c2c43e879dc6a26864f6101208501527f194647181790c1baf4305be04f26062595022ad4bb0cfa4713fe4aa50a8649236101408501527f1ecfaf5b9f0843a15cb6c0eecc19bb4c3d647637dbcc2687aea2e2be1915e3cb6101608501527f2a2997ebee31b4b2d7f6179059e038a5190f6a8cb68a7f12797f0af412bd7c016101808501527f1be4490a7b9910bc8c236e7a71301ca22057d2fc3ab63a3ab20fa38a464dfab56101a08501527f26099a7125706a5395f8c448e37ffa480dd94e5a13438fb16c029e124a5416f06101c08501527f06a699e4746d8e177d196d6f0ceb7c8377131f4ea1e734c346b922087c8fd6df6101e08501527f1f8ec2eb1e9ad31b29b63f03cf0b06d265f5e6e43cc2fe322b4dac076ccfd1456102008501527f233516ba8994453b7ea0b0d5c8adae7703c009782f4bed9b7af22183d3351bc86102208501526102408401526102608301527f122b52951678e09bd4e52cf794e55234e20985f7d061dce6c2166d79bc64cf246102808301527f2c6edb0893291528d4b998321f2a32c7832450eec23bc074e903ce43b3145c226102a08301527f1f1b430285344ffdc20eeaae385d27a1b856f9e35abbaeae4bf9e2f3eb92ef266102c08301527f22a30100945ca3c10bd916a9e8b6483dc9eeb5b3f3bdc2616ae68133cf93790e6102e0830152519283916109e384846115df565b8336843760085afa15908115610a0b575b506109fc5780f35b631ff3747d60e21b8152600490fd5b600191505114155f6109f4565b63a54f8e2760e01b8f5260048ffd5b6351d49ff760e11b8752600487fd5b80fd5b5034610b7d576040366003190112610b7d5760043567ffffffffffffffff8111610b7d57610a6b9036906004016115b1565b60243567ffffffffffffffff8111610b7d57610a8b9036906004016115b1565b6101009081604051610a9d82826115df565b3690378201908083830312610b7d5781601f84011215610b7d5760405192610ac582856115df565b83918101928311610b7d57905b828210610b8157505050303b15610b7d5760405163224e668f60e11b8152610120600482015261012481018390529283919083906101448401375f6101448484010152602482015f905b60088210610b6357505050610144815f93601f80199101168101030181305afa8015610b5857610b4a575080f35b610b5691505f906115df565b005b6040513d5f823e3d90fd5b829350602080916001939451815201930191018492610b1c565b5f80fd5b8135815260209182019101610ad2565b34610b7d57610120366003190112610b7d5760043567ffffffffffffffff8111610b7d57610bc39036906004016115b1565b3661012411610b7d57610bd96101808214611615565b6040908151610be883826115df565b8236823761010084019280848337602093815193610c0686866115df565b85368637825190610c1784836115df565b60018252866064818401601f1987013682373761014011610b7d575f516020611cfe5f395f51905f52908351610c5f8161012a8a8201948888873760608301905f8252611658565b5190200684528180519182377f0edb2d5af6534a82d734f6c8c9599223561176c6e93148568279e1f726011e7c828201527f11ce9383c53d66901aa59bb23efa55a7d1566a4f8903e6295b5004a41c05da9760608201527f2978bd840f86c9fcf14202e8f2d168a9d2f2375c49eee3b39cd22d4318c62bd060808201527f2494f1da9eaa4decf1a5a2260dc62840b4e96cff8688d58ff56f8d8e3e5a3a3060a082015281610140870160c08301377f2b2fce9edc8be658347005d796bfc86ff82d0fc8daef0290c95953c1e5c9935c6101008201527f022521d8f81cf0c7cf119e45a05fdcc00055a32fa0e5e8aab713869c436b63ec6101208201527f199e42fbcbbda303f0d4a04a9cc2a418cc450d1f66520fa2cb552d261130d2d96101408201527f28d803a477737410df240d247636ca91f7171e28be001373083103488a434f6d61016082015284816101808160085afa90511615611433578051928184017f198afa3741fdf69bb460a8b9ef813b10b4713e3b2dc01da7de1ba4b6676b9e0f85525f516020611cfe5f395f51905f528387808801967f1f6c2819b0794a2fd7faeece26052552e3be6f94ab970ee71297817b246f9d5588528051855201519260608801938452818860808160065afa947f188c496ae019c47576f22616fc5e6cc08830b52d32a3eba8d9ac0cf7145fecdb82527f297dec4673c8916d92ead3aa420b45e10a7c3e78be509ad3a00f2bf847f38f5b85527f1e1542f5affdeb02ee411d8358b8cb2f87e302c1c0a13f6923f9b89667a1518e600160243560808c0198818a5287878760608160075afa92101616858c60808160065afa16167f086afd4e275d21964198f3c8a7622e23cad7187af5278abb7b0a05f8983ffba684527f1c9e6f2de2b2f030473884570ffcd4a981d28997e456246280f2ce8b69a5ca9f87526044359081895286868660608160075afa92101616848b60808160065afa167f06fd9e9988b1d4eb30dc4f0565a8e7a2ac763dbe2397f53e6263a6fc0080eedb84527f0305d3166d58c6988fffd2af2c15be234b1f44cacc4610037cba2b902086e71c87526064359081895286868660608160075afa92101616848b60808160065afa167f1af4e0f4150ee350ce8165c3d3692113f6b20c5fd6f849b21a8f454c4a8fa01a84527f21a1c0f657c9529df4c6818965b70f2d0d69c1ee07d73a457d29f7bd4edcdc7f87526084359081895286868660608160075afa92101616848b60808160065afa167f2b5cc1b293be91d5307c6ce964f2520964f72bce116d253ed7d22dc2587adab184527f04149627bee7c4d42ff97a65159071a8dc2b5786490de5f402d6bdb737a94db7875260a4359081895286868660608160075afa92101616848b60808160065afa167f02309bb2c7a547fe72ec3c407187dd213e90d7a3d46fe0381e5cfec7305a3e8784527f1d8db0fca5701759b556f20375f3f1fd1e78003cd2d57dbf51dbec77b76b39bd875260c4359081895286868660608160075afa92101616848b60808160065afa167f190faa9e7ecebb3fcc7e48fd3fb33b8d47e310136ac4fbb2989dd0c5a0f9283784527f215885b35016aadad09311ee0eba9d67c7f03fec32b2576b4f4cc14c8a219348875260e4359081895286868660608160075afa92101616848b60808160065afa167f28ff412be210b4fdac69993ba622c67dbd9f727b691c63325cfb3fe788c9874884527f157781cf5696de1b56e40ea61d9fa14ecfa96a87bdfeb249dfebd2964a3497488752610104359081895286868660608160075afa92101616848b60808160065afa16957f3043c0969def31e6a445ce8eed7d2e084e4caa31d815bf0d6e834acd983e85a28452525180955260608160075afa92101616818460808160065afa169251915192156114245761010090519485377f1df2f9b54952a25c5d792ebe7c3f05cdcc7dcda10ed263be9d1e54fe6afbf7936101008501527f013469a889baac39f11aa45e09ca2f00a124360f949d6c2c43e879dc6a26864f6101208501527f194647181790c1baf4305be04f26062595022ad4bb0cfa4713fe4aa50a8649236101408501527f1ecfaf5b9f0843a15cb6c0eecc19bb4c3d647637dbcc2687aea2e2be1915e3cb6101608501527f2a2997ebee31b4b2d7f6179059e038a5190f6a8cb68a7f12797f0af412bd7c016101808501527f1be4490a7b9910bc8c236e7a71301ca22057d2fc3ab63a3ab20fa38a464dfab56101a08501527f26099a7125706a5395f8c448e37ffa480dd94e5a13438fb16c029e124a5416f06101c08501527f06a699e4746d8e177d196d6f0ceb7c8377131f4ea1e734c346b922087c8fd6df6101e08501527f1f8ec2eb1e9ad31b29b63f03cf0b06d265f5e6e43cc2fe322b4dac076ccfd1456102008501527f233516ba8994453b7ea0b0d5c8adae7703c009782f4bed9b7af22183d3351bc86102208501526102408401526102608301527f122b52951678e09bd4e52cf794e55234e20985f7d061dce6c2166d79bc64cf246102808301527f2c6edb0893291528d4b998321f2a32c7832450eec23bc074e903ce43b3145c226102a08301527f1f1b430285344ffdc20eeaae385d27a1b856f9e35abbaeae4bf9e2f3eb92ef266102c08301527f22a30100945ca3c10bd916a9e8b6483dc9eeb5b3f3bdc2616ae68133cf93790e6102e0830152816103008160085afa9051161561141557005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b6351d49ff760e11b5f5260045ffd5b34610b7d575f366003190112610b7d5760206040517fd429a8d4d8eeea800cfca60a1adc27f4f14e30dd693a067f8e0f4cb22f64a60a8152f35b34610b7d576020366003190112610b7d5760043567ffffffffffffffff8111610b7d576114ad9036906004016115b1565b906080604051916114be82846115df565b813684376115586020916114e6610180604051976114dc868a6115df565b85368a3714611615565b6114f4838201358235611685565b85526115118482013560a083013560408401356060850135611726565b84870152604086015261152c60e082013560c0830135611685565b6060860152611545610120820135610100830135611685565b8652610140610160820135910135611685565b9160405193845f905b6004821061159b57505050830193905f945b600186106115865760c0858560a0820152f35b81806001928551815201930195019491611573565b8251815291840191600191909101908401611561565b9181601f84011215610b7d5782359167ffffffffffffffff8311610b7d5760208381860195010111610b7d57565b90601f8019910116810190811067ffffffffffffffff82111761160157604052565b634e487b7160e01b5f52604160045260245ffd5b1561161c57565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b80516020909101905f5b81811061166f5750505090565b8251845260209384019390920191600101611662565b905f516020611cde5f395f51905f52821080159061170f575b61141557811580611707575b611701576116ce5f516020611cde5f395f51905f5260038185818180090908611afe565b8181036116dd57505060011b90565b5f516020611cde5f395f51905f52809106810306145f1461141557600190811b1790565b50505f90565b5080156116aa565b505f516020611cde5f395f51905f5281101561169e565b919093925f516020611cde5f395f51905f528310801590611943575b801561192c575b8015611915575b61141557808286851717171561190a5790829161186d5f516020611cde5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611cde5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161184781808b80098187800908611afe565b8408095f516020611cde5f395f51905f5261186182611c75565b80091415958691611b21565b929080821480611901575b1561189f5750505050905f146118975760ff60025b169060021b179190565b60ff5f61188d565b5f516020611cde5f395f51905f528091068103061491826118e2575b50501561141557600191156118da5760ff60025b169060021b17179190565b60ff5f6118cf565b5f516020611cde5f395f51905f52919250819006810306145f806118bb565b50838314611878565b50505090505f905f90565b505f516020611cde5f395f51905f52811015611750565b505f516020611cde5f395f51905f52821015611749565b505f516020611cde5f395f51905f52851015611742565b80156119be578060011c915f516020611cde5f395f51905f528310156114155760018061199d5f516020611cde5f395f51905f5260038188818180090908611afe565b9316146119a657565b905f516020611cde5f395f51905f5280910681030690565b505f905f90565b801580611af6575b611aea578060021c92825f516020611cde5f395f51905f528510801590611ad3575b6114155784815f516020611cde5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4481611a9d9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611b21565b80929160018082961614611aaf575050565b5f516020611cde5f395f51905f528093945080929550809106810306930681030690565b505f516020611cde5f395f51905f528110156119ef565b50505f905f905f905f90565b5081156119cd565b90611b0882611c75565b915f516020611cde5f395f51905f528380090361141557565b915f516020611cde5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611b7993969496611b6b82808a8009818a800908611afe565b90611c69575b860809611afe565b925f516020611cde5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611cde5f395f51905f5260a083015260208260c08160055afa91519115611415575f516020611cde5f395f51905f52826001920903611415575f516020611cde5f395f51905f52908209925f516020611cde5f395f51905f528080808780090681030681878009081490811591611c4a575b5061141557565b90505f516020611cde5f395f51905f528084860960020914155f611c43565b81809106810306611b71565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611cde5f395f51905f5260a083015260208260c08160055afa915191156114155756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a264697066735822122056e4b8d7f63477095c85cacce58f21c2de63d6cf19cc11b29e6cdfa78e8e72b464736f6c634300081c0033",
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

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) CompressProof(opts *bind.CallOpts, proof []byte) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "compressProof", proof)

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

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) CompressProof(proof []byte) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _StateTransitionVerifierGroth16.Contract.CompressProof(&_StateTransitionVerifierGroth16.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) CompressProof(proof []byte) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _StateTransitionVerifierGroth16.Contract.CompressProof(&_StateTransitionVerifierGroth16.CallOpts, proof)
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

// VerifyProof is a free data retrieval call binding the contract method 0x449ccd1e.
//
// Solidity: function verifyProof(bytes proof, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyProof(opts *bind.CallOpts, proof []byte, input [8]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x449ccd1e.
//
// Solidity: function verifyProof(bytes proof, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyProof(proof []byte, input [8]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof(&_StateTransitionVerifierGroth16.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x449ccd1e.
//
// Solidity: function verifyProof(bytes proof, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyProof(proof []byte, input [8]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof(&_StateTransitionVerifierGroth16.CallOpts, proof, input)
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
