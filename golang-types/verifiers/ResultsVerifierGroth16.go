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
	Bin: "0x60808060405234601557611e7b908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461167e575080635d26278e14610c8d57806360e583461461037f578063b1c3a00e1461026e5763b8e72af614610055575f80fd5b346102365760403660031901126102365760043567ffffffffffffffff8111610236576100869036906004016116b6565b60243567ffffffffffffffff8111610236576100a69036906004016116b6565b61010093846040516100b882826116e4565b3690376040938480516100cb82826116e4565b3690378480516100db82826116e4565b3690378101610180828203126102365780601f8301121561023657845195610105610100886116e4565b8201868282116102365783905b82821061025e5750509061014061012c8261013394611bde565b9301611bde565b916101209081865161014582826116e4565b36903784019080858303126102365781601f860112156102365785519461016e610120876116e4565b8591810192831161023657905b82821061023a57505050303b15610236578351633072c1a360e11b8152945f600487015b6008821061022057505050906101bd6101c89261010487019061173e565b61014485019061173e565b5f61018484015b6009821061020a575050505f826102a481305afa90811561020157506101f3575080f35b6101ff91505f906116e4565b005b513d5f823e3d90fd5b60208060019285518152019301910190916101cf565b602080600192855181520193019101909161019f565b5f80fd5b813581526020918201910161017b565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610112565b3461023657610180366003190112610236573661010411610236573661014411610236573661018411610236576040516102a96080826116e4565b60803682376040519060206102be81846116e4565b803684376102d0602435600435611909565b82526102e660843560a4356044356064356119aa565b8284015260408301526102fd60e43560c435611909565b60608301526103126101243561010435611909565b83526103246101643561014435611909565b9060405192835f905b60048210610369575050506080830193905f945b600186106103545760c0858560a0820152f35b81806001928551815201930195019491610341565b825181529183019160019190910190830161032d565b34610236576102a036600319011261023657366101041161023657366101441161023657366101841161023657366102a4116102365760405160206103c481836116e4565b803683375f516020611e265f395f51905f5260405182810190610400816103f2610124356101043586611706565b03601f1981018352826116e4565b519020068252604051604061010482377f222a11e5a1fa455473a7d8ee37037b7360f26184362f37931965f306f9fb02a460408201527f2c3841a672fd7ebd288ffc216ed99f9f948602cc2b598b060a2c5c38051b249b60608201527f20a4faaf8edd4f1ca3a79fdc4549aa3de0497adc3277b662aa02231cd31b70c260808201527f254eddb67c04de0e884152db574a42fa772ef20f18ee78c53f80fd77c014739760a0820152604061014460c08301377f1197eee61cc440c73efb850b21936d39424f00466970e4312c8610a25b2d1bb96101008201527f2f58382774fd44cd86350ec4739ac86a4c88c25d389629d7cdf1f4d238d8400a6101208201527f1ef12d9c41366f9e89041c590ad60025322c377a17fbd52a6042e39ce3b6e9836101408201527f222ea6108d7d4dbc652743ebab55fa4cba8ec4313889097aa749c4fe94226cbb61016082015281816101808160085afa90511615610c7e5760408051929061057090846116e4565b61010482845b6101448310610c6e5750505060405190604082017f2b11d2aad66b2e5de7bc5c32dea8d7c6e6f61b91fe0274879914fba5b9a4fa2e83525f516020611e265f395f51905f52604085808601977f085ed1f31aee200e0e4fe1f2e4dce72625a6a0f443fab336ac52d36288d94fc889528051855201519260608601938452818660808160065afa947f22bc9d1bd5bd55ba95b4705ef3915f3ae54f106cc341d68fa3d8ddb1413c719682527f057cc358f348d3d428b2ed1f7597d5ac6270964c04292f938bf2881768b1298885527f2e6d4b0d9931eab376671ee5c9e8956a92e71ee9db8a8e7a28cb24e73c8b8a7260016101843560808a0198818a5287878760608160075afa92101616858a60808160065afa16167f0c44fbf8a54d7e1be9dccec9c0055e5cb443543559d01141f0c2fc1b24bddfc984527f2625162083823a129368b22f4976ee21ba246283f409cca35147a0daa082a9dc87526101a4359081895286868660608160075afa92101616848960808160065afa167f2d817d224dba2ca8ae799a71fd04d3eae7e317c4716064126e78eadf1bddc19384527f07360180b1e6209c27e0f0ab9aaf8a1cfc82b580ccba300159cceca7e1e0435a87526101c4359081895286868660608160075afa92101616848960808160065afa167f084d7c3a02f36da14bb43c7f6a2496d526bbd8022cfc93cdd60dda08414137b484527f2c052811f54cee8f3d1e678d0830c59725b92d18e1a78008789335fc04b42c7287526101e4359081895286868660608160075afa92101616848960808160065afa167f01c000ed078e0e50fb92cddc5c7f63f795b7930ad01c4e0fac6977602311b2a684527f12661b14b3517a2037010a4e59b62efa6c7665a279d90c7f2162c27ae30b28378752610204359081895286868660608160075afa92101616848960808160065afa167f1930f6e6c400e96eaf3fd146d05ce4189ea36a77dfd20895ecd05e1aa590d6e284527f2555538e1bb4d818d6e2c3f37440f2114172e879b1feabb026f3f5eed9a614d38752610224359081895286868660608160075afa92101616848960808160065afa167f052cecd0a6c307d8fa3c1c4e093b30aa79cd8753bb7b6ad569e803e20311198384527f158c8e0e1754b73c255823b606690f73f5204af05a24cbd300cabf4091559ca08752610244359081895286868660608160075afa92101616848960808160065afa167f2169dfcfc704758ce2eedc72026f89e2500534c97cb451252caf3da17aa7b26b84527f2524c400945827a371220bedea45ceb55f78f9e0253fbe19c396b55793b7a5558752610264359081895286868660608160075afa92101616848960808160065afa167f25dd3501e409e81e21dc8bb9b8b19422403525555179db5c3eb57e0b7ad2eba284527f1771134d12ef003c358932c33b3b191f89e04e1307a3181c5415b490afcfe0ee8752610284359081895286868660608160075afa92101616848960808160065afa16957f131592a6da4b7dd846d1e4315e44b4e0c23e5c032f60c4a94825b4fd2292ec578452525180955260608160075afa9210161660408260808160065afa16905192519015610c5f5760405192610100600485377f130a3b67f543deda846445d6230f06dbcc4fc9230d34f91a4b1faad0efe61a8a6101008501527f10e55a50c7b08a64e18b7bde27ec4110fe7fb12202121e6ff101944c875337146101208501527f27bca9fe60d54b8da5015075996a8c1809d5147158087f42e6b066a1141f51776101408501527f1b75effaed932252e6501d6431f3903cd63e860385d61a30a674af708bb976616101608501527f0e414ca47ed0d06cea28bb089a5c55a8f2b279595f96720bfddcdaa56719edda6101808501527f0b60dc69b88f82c08f96b4cfa6b9a2529850f6dd7c665caee80756132e412bf46101a08501527f0b5cbe03d20f03e13c4091fbcddce607c89ce811700287e0ce57f7bf7673c9716101c08501527f217b18acf2565725c2ca9da0ed7d01d58e4d7755cc0d65e2df8be4308782291f6101e08501527f0e39b086089c4a1bd56a58c895756919edabd8a5db92c37d0820c45f1d7404506102008501527f22ac99990ad6874151104f5ceb45277f14c37b1c62ef0396e5c76b9378d740cc6102208501526102408401526102608301527f050b4460d2b42db249c2a912daa47cdbd1a63d9d3f802869aba28520d98aaa346102808301527f2cc154a7a501fb364878767578b3e84476fb3b15e1e4f1f36c41e6e3c4a4f0066102a08301527f021624c1c2e20436e57f00c869b4f4aae66e8f16607b9027306b9abd91bab7156102c08301527f253dffdc41beb8cb4a01c3068dc7a60501047d6003fef6f81a418f3dd70ae7bf6102e0830152816103008160085afa90511615610c5057005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191849101610576565b6351d49ff760e11b5f5260045ffd5b34610236576101e03660031901126102365736608411610236573660a41161023657366101e411610236576040516020610cc781836116e4565b803683376040805190610cda81836116e4565b8036833761030093815192610cef86856116e4565b85368537610cfe608435611765565b868301528152610d0f60a435611765565b905f516020611e265f395f51905f5283516103f2610d398a87015189519283918d83019586611706565b5190200684528251865286830151878701527f222a11e5a1fa455473a7d8ee37037b7360f26184362f37931965f306f9fb02a4858701527f2c3841a672fd7ebd288ffc216ed99f9f948602cc2b598b060a2c5c38051b249b60608701527f20a4faaf8edd4f1ca3a79fdc4549aa3de0497adc3277b662aa02231cd31b70c260808701527f254eddb67c04de0e884152db574a42fa772ef20f18ee78c53f80fd77c014739760a087015260c086015260e08501527f1197eee61cc440c73efb850b21936d39424f00466970e4312c8610a25b2d1bb96101008501527f2f58382774fd44cd86350ec4739ac86a4c88c25d389629d7cdf1f4d238d8400a6101208501527f1ef12d9c41366f9e89041c590ad60025322c377a17fbd52a6042e39ce3b6e9836101408501527f222ea6108d7d4dbc652743ebab55fa4cba8ec4313889097aa749c4fe94226cbb610160850152825185816101808760085afa90511615610c7e578490610ea9600435611765565b610eba6024959295356044356117d0565b9291909389610eca606435611765565b9890977f2e6d4b0d9931eab376671ee5c9e8956a92e71ee9db8a8e7a28cb24e73c8b8a7283519b8c937f2b11d2aad66b2e5de7bc5c32dea8d7c6e6f61b91fe0274879914fba5b9a4fa2e85527f085ed1f31aee200e0e4fe1f2e4dce72625a6a0f443fab336ac52d36288d94fc8828601528051868601520151606084019081527f131592a6da4b7dd846d1e4315e44b4e0c23e5c032f60c4a94825b4fd2292ec57858560808160065afa957f22bc9d1bd5bd55ba95b4705ef3915f3ae54f106cc341d68fa3d8ddb1413c7196818701527f057cc358f348d3d428b2ed1f7597d5ac6270964c04292f938bf2881768b129888352600160c4356080880198818a525f516020611e265f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f0c44fbf8a54d7e1be9dccec9c0055e5cb443543559d01141f0c2fc1b24bddfc9828801527f2625162083823a129368b22f4976ee21ba246283f409cca35147a0daa082a9dc845260e435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2d817d224dba2ca8ae799a71fd04d3eae7e317c4716064126e78eadf1bddc193828801527f07360180b1e6209c27e0f0ab9aaf8a1cfc82b580ccba300159cceca7e1e0435a845261010435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f084d7c3a02f36da14bb43c7f6a2496d526bbd8022cfc93cdd60dda08414137b4828801527f2c052811f54cee8f3d1e678d0830c59725b92d18e1a78008789335fc04b42c72845261012435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f01c000ed078e0e50fb92cddc5c7f63f795b7930ad01c4e0fac6977602311b2a6828801527f12661b14b3517a2037010a4e59b62efa6c7665a279d90c7f2162c27ae30b2837845261014435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1930f6e6c400e96eaf3fd146d05ce4189ea36a77dfd20895ecd05e1aa590d6e2828801527f2555538e1bb4d818d6e2c3f37440f2114172e879b1feabb026f3f5eed9a614d3845261016435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f052cecd0a6c307d8fa3c1c4e093b30aa79cd8753bb7b6ad569e803e203111983828801527f158c8e0e1754b73c255823b606690f73f5204af05a24cbd300cabf4091559ca0845261018435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2169dfcfc704758ce2eedc72026f89e2500534c97cb451252caf3da17aa7b26b828801527f2524c400945827a371220bedea45ceb55f78f9e0253fbe19c396b55793b7a55584526101a435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f25dd3501e409e81e21dc8bb9b8b19422403525555179db5c3eb57e0b7ad2eba2828801527f1771134d12ef003c358932c33b3b191f89e04e1307a3181c5415b490afcfe0ee84526101c435908189525f516020611e265f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611e265f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610c5f578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f130a3b67f543deda846445d6230f06dbcc4fc9230d34f91a4b1faad0efe61a8a6101008501527f10e55a50c7b08a64e18b7bde27ec4110fe7fb12202121e6ff101944c875337146101208501527f27bca9fe60d54b8da5015075996a8c1809d5147158087f42e6b066a1141f51776101408501527f1b75effaed932252e6501d6431f3903cd63e860385d61a30a674af708bb976616101608501527f0e414ca47ed0d06cea28bb089a5c55a8f2b279595f96720bfddcdaa56719edda6101808501527f0b60dc69b88f82c08f96b4cfa6b9a2529850f6dd7c665caee80756132e412bf46101a08501527f0b5cbe03d20f03e13c4091fbcddce607c89ce811700287e0ce57f7bf7673c9716101c08501527f217b18acf2565725c2ca9da0ed7d01d58e4d7755cc0d65e2df8be4308782291f6101e08501527f0e39b086089c4a1bd56a58c895756919edabd8a5db92c37d0820c45f1d7404506102008501527f22ac99990ad6874151104f5ceb45277f14c37b1c62ef0396e5c76b9378d740cc6102208501526102408401526102608301527f050b4460d2b42db249c2a912daa47cdbd1a63d9d3f802869aba28520d98aaa346102808301527f2cc154a7a501fb364878767578b3e84476fb3b15e1e4f1f36c41e6e3c4a4f0066102a08301527f021624c1c2e20436e57f00c869b4f4aae66e8f16607b9027306b9abd91bab7156102c08301527f253dffdc41beb8cb4a01c3068dc7a60501047d6003fef6f81a418f3dd70ae7bf6102e08301525192839161165984846116e4565b8336843760085afa15908115611671575b50610c5057005b600191505114158161166a565b34610236575f36600319011261023657807f0e408af9c063d217680deb42b0570387d9ca2e6f9a87710af091f3491cc45c9960209252f35b9181601f840112156102365782359167ffffffffffffffff8311610236576020838186019501011161023657565b90601f8019910116810190811067ffffffffffffffff82111761024a57604052565b909160409282526020820152016060516080905f5b8181106117285750505090565b825184526020938401939092019160010161171b565b905f905b6002821061174f57505050565b6020806001928551815201930191019091611742565b80156117c9578060011c915f516020611e065f395f51905f52831015610c50576001806117a85f516020611e065f395f51905f5260038188818180090908611c26565b9316146117b157565b905f516020611e065f395f51905f5280910681030690565b505f905f90565b801580611901575b6118f5578060021c92825f516020611e065f395f51905f5285108015906118de575b610c505784815f516020611e065f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816118a89d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611c49565b809291600180829616146118ba575050565b5f516020611e065f395f51905f528093945080929550809106810306930681030690565b505f516020611e065f395f51905f528110156117fa565b50505f905f905f905f90565b5081156117d8565b905f516020611e065f395f51905f528210801590611993575b610c505781158061198b575b611985576119525f516020611e065f395f51905f5260038185818180090908611c26565b81810361196157505060011b90565b5f516020611e065f395f51905f52809106810306145f14610c5057600190811b1790565b50505f90565b50801561192e565b505f516020611e065f395f51905f52811015611922565b919093925f516020611e065f395f51905f528310801590611bc7575b8015611bb0575b8015611b99575b610c50578082868517171715611b8e57908291611af15f516020611e065f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611e065f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611acb81808b80098187800908611c26565b8408095f516020611e065f395f51905f52611ae582611d9d565b80091415958691611c49565b929080821480611b85575b15611b235750505050905f14611b1b5760ff60025b169060021b179190565b60ff5f611b11565b5f516020611e065f395f51905f52809106810306149182611b66575b505015610c505760019115611b5e5760ff60025b169060021b17179190565b60ff5f611b53565b5f516020611e065f395f51905f52919250819006810306145f80611b3f565b50838314611afc565b50505090505f905f90565b505f516020611e065f395f51905f528110156119d4565b505f516020611e065f395f51905f528210156119cd565b505f516020611e065f395f51905f528510156119c6565b9080601f8301121561023657604080519290611bfa90846116e4565b82906040810192831161023657905b828210611c165750505090565b8135815260209182019101611c09565b90611c3082611d9d565b915f516020611e065f395f51905f5283800903610c5057565b915f516020611e065f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611ca193969496611c9382808a8009818a800908611c26565b90611d91575b860809611c26565b925f516020611e065f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611e065f395f51905f5260a083015260208260c08160055afa91519115610c50575f516020611e065f395f51905f52826001920903610c50575f516020611e065f395f51905f52908209925f516020611e065f395f51905f528080808780090681030681878009081490811591611d72575b50610c5057565b90505f516020611e065f395f51905f528084860960020914155f611d6b565b81809106810306611c99565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611e065f395f51905f5260a083015260208260c08160055afa91519115610c505756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220cbbf26a311ec3345a663d7afcbf885cbea1907432a33e8d9d30357dbb104123964736f6c634300081c0033",
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
