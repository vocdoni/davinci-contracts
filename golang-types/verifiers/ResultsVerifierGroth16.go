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
	Bin: "0x60808060405234601557611e77908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461167a575080635d26278e14610c8b57806360e583461461037f578063b1c3a00e1461026e5763b8e72af614610055575f80fd5b346102365760403660031901126102365760043567ffffffffffffffff8111610236576100869036906004016116b2565b60243567ffffffffffffffff8111610236576100a69036906004016116b2565b61010093846040516100b882826116e0565b3690376040938480516100cb82826116e0565b3690378480516100db82826116e0565b3690378101610180828203126102365780601f8301121561023657845195610105610100886116e0565b8201868282116102365783905b82821061025e5750509061014061012c8261013394611bda565b9301611bda565b916101209081865161014582826116e0565b36903784019080858303126102365781601f860112156102365785519461016e610120876116e0565b8591810192831161023657905b82821061023a57505050303b15610236578351633072c1a360e11b8152945f600487015b6008821061022057505050906101bd6101c89261010487019061173a565b61014485019061173a565b5f61018484015b6009821061020a575050505f826102a481305afa90811561020157506101f3575080f35b6101ff91505f906116e0565b005b513d5f823e3d90fd5b60208060019285518152019301910190916101cf565b602080600192855181520193019101909161019f565b5f80fd5b813581526020918201910161017b565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610112565b3461023657610180366003190112610236573661010411610236573661014411610236573661018411610236576040516102a96080826116e0565b60803682376040519060206102be81846116e0565b803684376102d0602435600435611905565b82526102e660843560a4356044356064356119a6565b8284015260408301526102fd60e43560c435611905565b60608301526103126101243561010435611905565b83526103246101643561014435611905565b9060405192835f905b60048210610369575050506080830193905f945b600186106103545760c0858560a0820152f35b81806001928551815201930195019491610341565b825181529183019160019190910190830161032d565b34610236576102a036600319011261023657366101041161023657366101441161023657366101841161023657366102a4116102365760405160206103c481836116e0565b803683375f516020611e225f395f51905f5260405182810190610400816103f2610124356101043586611702565b03601f1981018352826116e0565b519020068252604051604061010482377e5e05be57e171c3a347db635d092b7e68c39800f835f0413afb63112ec5288760408201527f2179ccc912b796ac01bee1ad052d04c1505c90e70b148dd087152a7aec140bf560608201527f0643313911eee539d5d85553653b1aab35cc2bce090475777323511c83c92fa360808201527f28948810f3674b1a43264025dcaa55ab0ebf02f878c710924df9355e42542b2660a0820152604061014460c08301377f11c8aae23950e89df078ca38154ec65d898a2e0bcd75ce597a09ca0ff0104d246101008201527f0283e55585eeeff03dc42c3b3253765606aa7517414e3f4ab3cc2d11e13f082e6101208201527f1fbd305985e9603404e1a4fa733822c14b01f6cdb3bdb688fdbc5f1c3b1841976101408201527f0a710e3954904abb6e6b58c8b96f1df3a309cbc1bf9b355434db2b3873b8581d61016082015281816101808160085afa90511615610c7c5760408051929061056f90846116e0565b61010482845b6101448310610c6c5750505060405190604082017f1973319dc2c60b1809ffee7d137baa5ce915c873ecda66348811eb0b7164d05e83525f516020611e225f395f51905f52604085808601977f10d6e7b4103ec74a51963af7965c4fb75597f1aca06be92530a2604581b9a22989528051855201519260608601938452818660808160065afa947f0bbc6a73ea810759c14873c10055f4c99a889286bb5dc60e24310173d979bd5382527f0fb84d242b096873f6f292c35d22d14f0b1ee3390542fdbe56257f17b4f31f2b85527f1a68f8ca362c02e265c1bbc85e8b38984e087e47d0ad334389c62fbd1d0ddfde60016101843560808a0198818a5287878760608160075afa92101616858a60808160065afa16167f053ef158f34cac51dd997ad04d4710f76b07b4b4cf70e1637d58164f6b47e06f84527f06b43fddfa8cb29417f9109da64c1875b8b2503a539e5a2ee1381ff40359228f87526101a4359081895286868660608160075afa92101616848960808160065afa167f216331ae90f74ba87032e90005bb94715dcb9726482a80d6bc795c6be9183b1b84527f0dec6faac2bc3f692c33b0743c96257f7c7bbf60ea10f31005a7aa2d17cfac5587526101c4359081895286868660608160075afa92101616848960808160065afa167f1647346e2176160420b49a409059225dec433eab7249bb018ad84182cbdf71ec84527f1c1f123b0faf9f3ad0e7fac0c98fe90dae5a7ff331cf73c547eda1c2ee4a2ab387526101e4359081895286868660608160075afa92101616848960808160065afa167f22b9920dc31273977a34876de223438bb7b9cf998c2f2a20830135c5e8968bac84527f2802e5e9551b410b41c658627501988e6cdba9ee3d91a3c41554b216e1374e748752610204359081895286868660608160075afa92101616848960808160065afa167f13e23be2d0d10d9d9a522740ff9f16ebf3a640d6457c86893c785ee69d421d1684527f0545bd05fa0164555981fa368b88770974cc5d0caab416206f9b64720391a2db8752610224359081895286868660608160075afa92101616848960808160065afa167f24aac08ac31677a9327371e8608c59d80f436a23ce999370854664ba930f475084527f0543a83130dc1a1209f562a0e64aee11d3cd5ba73cbec42637eac197c8cea3308752610244359081895286868660608160075afa92101616848960808160065afa167f06b312c446d4a4b093e93a0365442b9ce68ed2a3ce8f48a535d5eabb4753d3f784527e3207319db3da4e43eafc290ca1b0be5c0270532c0e626584c8c5cb31b8c5278752610264359081895286868660608160075afa92101616848960808160065afa167f0cc107e0c46d7cc4c9748996de600cca2b69737e6b7d1c1b824b53da79188a9284527f29538badc1a6e3dc55ea28cc058864befdc57142e5bee37dc12a189aa072bf518752610284359081895286868660608160075afa92101616848960808160065afa16957f035884d3b482d6c0fd153b085db75f70f23542758fff345c7b46a839b4126f088452525180955260608160075afa9210161660408260808160065afa16905192519015610c5d5760405192610100600485377f19e805e891035c3e515c4849adb5fc70a193fb55299ff8170f74e49dcce9c0876101008501527f10f5d2fe5800355674050513d6f9ad809fb74252fb668793e697e6a53fe62b1c6101208501527f188c60f6ffbf01fa70d95b24c80e2b3e93ef216f38130959ecfef9e9059160956101408501527f09b6742c38fd756588d401751fc56341a7b04d4eda4818cb614377aaeb5990dc6101608501527f2e06add4f4f5c8e99d198c33e68cb7e94b0eba3abf3c373da956dd06483ca6eb6101808501527f140a14898c366bcccdf13e8d60be7bb3ca08fb0f3cdab1762539618892efa9576101a08501527f14b2c0b679a7bf23f61db0ce848885b1d23bb2a20587b2f708d33817c97095246101c08501527f2560c8068242f8b26bb0917ded1a3491b359f163a15ffd74590cf51d3de9f6da6101e08501527f0982914d4f35129270478f9f66594c2d5ce0ef86762b78bcedc7e468972cdd6c6102008501527f093050fb59cbb7712fe1cd41a3fd7443cac3da58cee330cb04314be4b24b67c96102208501526102408401526102608301527f24e375ab5edde07add353aabc9bfb17724014824072d0f13c44eae1c43693a4f6102808301527f09bdd0ee79664950b6b6a5c03495fa53f1628caaf1b6ee06e5eea7372278026c6102a08301527f1e69e4e8b120f69492993e4788222fd7faf8c8796e162b91f044f828d61fd1396102c08301527f2ceb6098ad916f7b32289d408e7854ae887df3067670baa27c2611d5e19b86fc6102e0830152816103008160085afa90511615610c4e57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191849101610575565b6351d49ff760e11b5f5260045ffd5b34610236576101e03660031901126102365736608411610236573660a41161023657366101e411610236576040516020610cc581836116e0565b803683376040805190610cd881836116e0565b8036833761030093815192610ced86856116e0565b85368537610cfc608435611761565b868301528152610d0d60a435611761565b905f516020611e225f395f51905f5283516103f2610d378a87015189519283918d83019586611702565b5190200684528251865286830151878701527e5e05be57e171c3a347db635d092b7e68c39800f835f0413afb63112ec52887858701527f2179ccc912b796ac01bee1ad052d04c1505c90e70b148dd087152a7aec140bf560608701527f0643313911eee539d5d85553653b1aab35cc2bce090475777323511c83c92fa360808701527f28948810f3674b1a43264025dcaa55ab0ebf02f878c710924df9355e42542b2660a087015260c086015260e08501527f11c8aae23950e89df078ca38154ec65d898a2e0bcd75ce597a09ca0ff0104d246101008501527f0283e55585eeeff03dc42c3b3253765606aa7517414e3f4ab3cc2d11e13f082e6101208501527f1fbd305985e9603404e1a4fa733822c14b01f6cdb3bdb688fdbc5f1c3b1841976101408501527f0a710e3954904abb6e6b58c8b96f1df3a309cbc1bf9b355434db2b3873b8581d610160850152825185816101808760085afa90511615610c7c578490610ea6600435611761565b610eb76024959295356044356117cc565b9291909389610ec7606435611761565b9890977f1a68f8ca362c02e265c1bbc85e8b38984e087e47d0ad334389c62fbd1d0ddfde83519b8c937f1973319dc2c60b1809ffee7d137baa5ce915c873ecda66348811eb0b7164d05e85527f10d6e7b4103ec74a51963af7965c4fb75597f1aca06be92530a2604581b9a229828601528051868601520151606084019081527f035884d3b482d6c0fd153b085db75f70f23542758fff345c7b46a839b4126f08858560808160065afa957f0bbc6a73ea810759c14873c10055f4c99a889286bb5dc60e24310173d979bd53818701527f0fb84d242b096873f6f292c35d22d14f0b1ee3390542fdbe56257f17b4f31f2b8352600160c4356080880198818a525f516020611e225f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f053ef158f34cac51dd997ad04d4710f76b07b4b4cf70e1637d58164f6b47e06f828801527f06b43fddfa8cb29417f9109da64c1875b8b2503a539e5a2ee1381ff40359228f845260e435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f216331ae90f74ba87032e90005bb94715dcb9726482a80d6bc795c6be9183b1b828801527f0dec6faac2bc3f692c33b0743c96257f7c7bbf60ea10f31005a7aa2d17cfac55845261010435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1647346e2176160420b49a409059225dec433eab7249bb018ad84182cbdf71ec828801527f1c1f123b0faf9f3ad0e7fac0c98fe90dae5a7ff331cf73c547eda1c2ee4a2ab3845261012435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f22b9920dc31273977a34876de223438bb7b9cf998c2f2a20830135c5e8968bac828801527f2802e5e9551b410b41c658627501988e6cdba9ee3d91a3c41554b216e1374e74845261014435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f13e23be2d0d10d9d9a522740ff9f16ebf3a640d6457c86893c785ee69d421d16828801527f0545bd05fa0164555981fa368b88770974cc5d0caab416206f9b64720391a2db845261016435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f24aac08ac31677a9327371e8608c59d80f436a23ce999370854664ba930f4750828801527f0543a83130dc1a1209f562a0e64aee11d3cd5ba73cbec42637eac197c8cea330845261018435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f06b312c446d4a4b093e93a0365442b9ce68ed2a3ce8f48a535d5eabb4753d3f7828801527e3207319db3da4e43eafc290ca1b0be5c0270532c0e626584c8c5cb31b8c52784526101a435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0cc107e0c46d7cc4c9748996de600cca2b69737e6b7d1c1b824b53da79188a92828801527f29538badc1a6e3dc55ea28cc058864befdc57142e5bee37dc12a189aa072bf5184526101c435908189525f516020611e225f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611e225f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610c5d578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f19e805e891035c3e515c4849adb5fc70a193fb55299ff8170f74e49dcce9c0876101008501527f10f5d2fe5800355674050513d6f9ad809fb74252fb668793e697e6a53fe62b1c6101208501527f188c60f6ffbf01fa70d95b24c80e2b3e93ef216f38130959ecfef9e9059160956101408501527f09b6742c38fd756588d401751fc56341a7b04d4eda4818cb614377aaeb5990dc6101608501527f2e06add4f4f5c8e99d198c33e68cb7e94b0eba3abf3c373da956dd06483ca6eb6101808501527f140a14898c366bcccdf13e8d60be7bb3ca08fb0f3cdab1762539618892efa9576101a08501527f14b2c0b679a7bf23f61db0ce848885b1d23bb2a20587b2f708d33817c97095246101c08501527f2560c8068242f8b26bb0917ded1a3491b359f163a15ffd74590cf51d3de9f6da6101e08501527f0982914d4f35129270478f9f66594c2d5ce0ef86762b78bcedc7e468972cdd6c6102008501527f093050fb59cbb7712fe1cd41a3fd7443cac3da58cee330cb04314be4b24b67c96102208501526102408401526102608301527f24e375ab5edde07add353aabc9bfb17724014824072d0f13c44eae1c43693a4f6102808301527f09bdd0ee79664950b6b6a5c03495fa53f1628caaf1b6ee06e5eea7372278026c6102a08301527f1e69e4e8b120f69492993e4788222fd7faf8c8796e162b91f044f828d61fd1396102c08301527f2ceb6098ad916f7b32289d408e7854ae887df3067670baa27c2611d5e19b86fc6102e08301525192839161165584846116e0565b8336843760085afa1590811561166d575b50610c4e57005b6001915051141581611666565b34610236575f36600319011261023657807fb590ba65ecac3aa3c4cc2acdacaa4401c607720beb50c12ea0dd6d23b64e87a060209252f35b9181601f840112156102365782359167ffffffffffffffff8311610236576020838186019501011161023657565b90601f8019910116810190811067ffffffffffffffff82111761024a57604052565b909160409282526020820152016060516080905f5b8181106117245750505090565b8251845260209384019390920191600101611717565b905f905b6002821061174b57505050565b602080600192855181520193019101909161173e565b80156117c5578060011c915f516020611e025f395f51905f52831015610c4e576001806117a45f516020611e025f395f51905f5260038188818180090908611c22565b9316146117ad57565b905f516020611e025f395f51905f5280910681030690565b505f905f90565b8015806118fd575b6118f1578060021c92825f516020611e025f395f51905f5285108015906118da575b610c4e5784815f516020611e025f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816118a49d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611c45565b809291600180829616146118b6575050565b5f516020611e025f395f51905f528093945080929550809106810306930681030690565b505f516020611e025f395f51905f528110156117f6565b50505f905f905f905f90565b5081156117d4565b905f516020611e025f395f51905f52821080159061198f575b610c4e57811580611987575b6119815761194e5f516020611e025f395f51905f5260038185818180090908611c22565b81810361195d57505060011b90565b5f516020611e025f395f51905f52809106810306145f14610c4e57600190811b1790565b50505f90565b50801561192a565b505f516020611e025f395f51905f5281101561191e565b919093925f516020611e025f395f51905f528310801590611bc3575b8015611bac575b8015611b95575b610c4e578082868517171715611b8a57908291611aed5f516020611e025f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611e025f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611ac781808b80098187800908611c22565b8408095f516020611e025f395f51905f52611ae182611d99565b80091415958691611c45565b929080821480611b81575b15611b1f5750505050905f14611b175760ff60025b169060021b179190565b60ff5f611b0d565b5f516020611e025f395f51905f52809106810306149182611b62575b505015610c4e5760019115611b5a5760ff60025b169060021b17179190565b60ff5f611b4f565b5f516020611e025f395f51905f52919250819006810306145f80611b3b565b50838314611af8565b50505090505f905f90565b505f516020611e025f395f51905f528110156119d0565b505f516020611e025f395f51905f528210156119c9565b505f516020611e025f395f51905f528510156119c2565b9080601f8301121561023657604080519290611bf690846116e0565b82906040810192831161023657905b828210611c125750505090565b8135815260209182019101611c05565b90611c2c82611d99565b915f516020611e025f395f51905f5283800903610c4e57565b915f516020611e025f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c9d93969496611c8f82808a8009818a800908611c22565b90611d8d575b860809611c22565b925f516020611e025f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611e025f395f51905f5260a083015260208260c08160055afa91519115610c4e575f516020611e025f395f51905f52826001920903610c4e575f516020611e025f395f51905f52908209925f516020611e025f395f51905f528080808780090681030681878009081490811591611d6e575b50610c4e57565b90505f516020611e025f395f51905f528084860960020914155f611d67565b81809106810306611c95565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611e025f395f51905f5260a083015260208260c08160055afa91519115610c4e5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220db99c6be2fdde81e4228242d99899055b9743b7165618f09db194bf158c6936264736f6c634300081c0033",
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
