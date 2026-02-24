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
	Bin: "0x60808060405234601557611e79908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461167c575080635d26278e14610c8c57806360e583461461037f578063b1c3a00e1461026e5763b8e72af614610055575f80fd5b346102365760403660031901126102365760043567ffffffffffffffff8111610236576100869036906004016116b4565b60243567ffffffffffffffff8111610236576100a69036906004016116b4565b61010093846040516100b882826116e2565b3690376040938480516100cb82826116e2565b3690378480516100db82826116e2565b3690378101610180828203126102365780601f8301121561023657845195610105610100886116e2565b8201868282116102365783905b82821061025e5750509061014061012c8261013394611bdc565b9301611bdc565b916101209081865161014582826116e2565b36903784019080858303126102365781601f860112156102365785519461016e610120876116e2565b8591810192831161023657905b82821061023a57505050303b15610236578351633072c1a360e11b8152945f600487015b6008821061022057505050906101bd6101c89261010487019061173c565b61014485019061173c565b5f61018484015b6009821061020a575050505f826102a481305afa90811561020157506101f3575080f35b6101ff91505f906116e2565b005b513d5f823e3d90fd5b60208060019285518152019301910190916101cf565b602080600192855181520193019101909161019f565b5f80fd5b813581526020918201910161017b565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610112565b3461023657610180366003190112610236573661010411610236573661014411610236573661018411610236576040516102a96080826116e2565b60803682376040519060206102be81846116e2565b803684376102d0602435600435611907565b82526102e660843560a4356044356064356119a8565b8284015260408301526102fd60e43560c435611907565b60608301526103126101243561010435611907565b83526103246101643561014435611907565b9060405192835f905b60048210610369575050506080830193905f945b600186106103545760c0858560a0820152f35b81806001928551815201930195019491610341565b825181529183019160019190910190830161032d565b34610236576102a036600319011261023657366101041161023657366101441161023657366101841161023657366102a4116102365760405160206103c481836116e2565b803683375f516020611e245f395f51905f5260405182810190610400816103f2610124356101043586611704565b03601f1981018352826116e2565b519020068252604051604061010482377f14f5ad23bfdf29ab296cc57b77b199afb5bea45f9e2dc54c9fd11c065cdf724f60408201527f2bf625f36f0bba282046b1f3d58960da310eb5241346ea7288ec1b540ac22cf560608201527f0a014ee8bd337ef459a79b569747547a62807b896c29178298d01d68c149866160808201527f20a97623c8f485acf26b13c517108623bd8d7acad5bed376e7a9f406c9de19aa60a0820152604061014460c08301377f254f834178ded679c437b2e052a547609e5287addfb1afc77611fc7ec85dd0d16101008201527f08e1c08d5a781716746179cf17a8ebe2f4b45546cce6e70c4a8b1d2afe5fc4e96101208201527f2fd9c9f121637c05ddde6f4bd965a75466ce9835f7ce7f7762a5153ff95d49aa6101408201527f045333a340ac47d5c799253216ad77e1ec24fb4e785ba8b750a52050fb33164961016082015281816101808160085afa90511615610c7d5760408051929061057090846116e2565b61010482845b6101448310610c6d5750505060405190604082017f1139b7c1b51cea906a3feeae3255bc8e68e1bef84a71c668bf7569d9f581e90583525f516020611e245f395f51905f52604085808601977f084341e10d030b5ad643f873e3cbb6ba60a9876b5ff6fc8cf6fd6f2547eb3fcc89528051855201519260608601938452818660808160065afa947f17d1fb3b3fb2ecb9c892eb8fb4ff930ea024932bdfaf0ba712181280f9734af182527f2e0af6e3055e43b5bde4a6ce7a06bd622c2743ca1b37aff0b700cd8abc5c308b85527f16ab1bfb804f22e70deb30858c4a3de9d0322241212612efe919264c63317fd460016101843560808a0198818a5287878760608160075afa92101616858a60808160065afa16167f04e63fac635f1db0a75b3ac9f1ab2cb57119f1fa68fb811f23f29dfa677f50b584527f02c244ee0d0ce0fadabb58a20feba03e23fdf29415f328b09e2b4ceca742ff6187526101a4359081895286868660608160075afa92101616848960808160065afa167f0a1b29f21c1a8478436fa6351c19a81bf5860acaab41ba75293a955d9110978b84527f1fda26a590bf4c83fade057ccd4b9bb3eb554a56b5c75a6f93c05f48e64c0b3e87526101c4359081895286868660608160075afa92101616848960808160065afa167f14947a1e32f5f0e59a223d64fcdb96b0a9a3219510b8854c2190e02fce22fb6e84527f2a9b8c17b0b5b6360d0171e97d51bc542f548cf034afed40d23849314e10e2e287526101e4359081895286868660608160075afa92101616848960808160065afa167f0aeb586f5d31b1a7b62b828776f329780eb01026ed2b729c2ab41e304a10d4cb84527f2e709d6f77a352e49ade095639638b10e1a33c27145a88fbac6c110c236b205b8752610204359081895286868660608160075afa92101616848960808160065afa167f08ce833b2dbd34cb8a4b00eaa45fe963d24cc944655b811d719c1d9583b4cd1d84527f28c63c36fb110ab2f256c5913d7e7d8017cba906c430b45d19d71cebdf0272268752610224359081895286868660608160075afa92101616848960808160065afa167f2b771270bdc4661d0cb6107b47a52f79e334785d41ca5f066c8944b4fd95f21984527f02fc865e5a17e8d38d82eadb62f7ee715d6004d46d9fe06be546f356ffbb5dd38752610244359081895286868660608160075afa92101616848960808160065afa167f115be55d78a138e0cc8215a42d0011d2a8a2d781ad16874cb4afda39908c4b0c84527f0c8a39b25f66b912a58ea652cb9bf2c8149c8d51e438dd0159eee8b8b151ef998752610264359081895286868660608160075afa92101616848960808160065afa167f118e76c9f149834d8975e7e1e3486a668cd4e89482d3a59d6f17d45484f6251884527f2c78ab2ec3ba68f93116cf48477b740632664681f3482a81b64cf352fdf7edc58752610284359081895286868660608160075afa92101616848960808160065afa16957f25ca340bd4dc5a0629c02fe442b3cdc7135c22829933a5f7619bc4804745812d8452525180955260608160075afa9210161660408260808160065afa16905192519015610c5e5760405192610100600485377f13fd1dfd8781fe3bdb30f4e390a4716e4fadbf64763c88a9fb56b16c40117db36101008501527f15f3d4660c4a69d77b07fa91e84f5aa8b94126c93752988235f8a937171d36b86101208501527f0f2aab6e68ef0ea52b15742aa2772b47622dbfa1a7b7dd93cb322c5d469273cc6101408501527f273e7d3120556d9330dec620de3099b236484cef7136ad9f15b389174398a1de6101608501527f2413a48ec2fa419ed4181f34f111c06baf8dbdeaf4e1defc55d6a1fb0576a8646101808501527f1ea205ecaed0fc03f580a43223bdc715355ad3eff5410ef23e38e60c1ee84ebd6101a08501527f070c7681a0a0afc103d7a1f9c5e0aae92d0f69e06b8d788c433d9658a70b74f06101c08501527f2186619a41c57807574dbc85151946e7b1fcfeeaaf2a0197ada080fa1e45dd9e6101e08501527f276c9a9426abbf40a134b8b309f58a52212bf43f7f97fbdabc8a960815292f126102008501527f0510e091d41f64779ae660caee8b24f8292e285bccf04a38d8c3c0cbc642df496102208501526102408401526102608301527f02f80747a8cd73f5a51f44bf710dc73f3a03d37e671f625e5446ef862900c7946102808301527f1f4156641b4ba54b3eff74c6297ec1bdc2cebb65ede72c7999c5cef7e11557e86102a08301527e448785732e9cbfd68980bbc7a1432cd48839fb8f9f76ed84adba7f45ca91f46102c08301527f034d395cea82a76fe06b1ad9b29a0328ae0a6b57700d47b576f781f8f2538a176102e0830152816103008160085afa90511615610c4f57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191849101610576565b6351d49ff760e11b5f5260045ffd5b34610236576101e03660031901126102365736608411610236573660a41161023657366101e411610236576040516020610cc681836116e2565b803683376040805190610cd981836116e2565b8036833761030093815192610cee86856116e2565b85368537610cfd608435611763565b868301528152610d0e60a435611763565b905f516020611e245f395f51905f5283516103f2610d388a87015189519283918d83019586611704565b5190200684528251865286830151878701527f14f5ad23bfdf29ab296cc57b77b199afb5bea45f9e2dc54c9fd11c065cdf724f858701527f2bf625f36f0bba282046b1f3d58960da310eb5241346ea7288ec1b540ac22cf560608701527f0a014ee8bd337ef459a79b569747547a62807b896c29178298d01d68c149866160808701527f20a97623c8f485acf26b13c517108623bd8d7acad5bed376e7a9f406c9de19aa60a087015260c086015260e08501527f254f834178ded679c437b2e052a547609e5287addfb1afc77611fc7ec85dd0d16101008501527f08e1c08d5a781716746179cf17a8ebe2f4b45546cce6e70c4a8b1d2afe5fc4e96101208501527f2fd9c9f121637c05ddde6f4bd965a75466ce9835f7ce7f7762a5153ff95d49aa6101408501527f045333a340ac47d5c799253216ad77e1ec24fb4e785ba8b750a52050fb331649610160850152825185816101808760085afa90511615610c7d578490610ea8600435611763565b610eb96024959295356044356117ce565b9291909389610ec9606435611763565b9890977f16ab1bfb804f22e70deb30858c4a3de9d0322241212612efe919264c63317fd483519b8c937f1139b7c1b51cea906a3feeae3255bc8e68e1bef84a71c668bf7569d9f581e90585527f084341e10d030b5ad643f873e3cbb6ba60a9876b5ff6fc8cf6fd6f2547eb3fcc828601528051868601520151606084019081527f25ca340bd4dc5a0629c02fe442b3cdc7135c22829933a5f7619bc4804745812d858560808160065afa957f17d1fb3b3fb2ecb9c892eb8fb4ff930ea024932bdfaf0ba712181280f9734af1818701527f2e0af6e3055e43b5bde4a6ce7a06bd622c2743ca1b37aff0b700cd8abc5c308b8352600160c4356080880198818a525f516020611e245f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f04e63fac635f1db0a75b3ac9f1ab2cb57119f1fa68fb811f23f29dfa677f50b5828801527f02c244ee0d0ce0fadabb58a20feba03e23fdf29415f328b09e2b4ceca742ff61845260e435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0a1b29f21c1a8478436fa6351c19a81bf5860acaab41ba75293a955d9110978b828801527f1fda26a590bf4c83fade057ccd4b9bb3eb554a56b5c75a6f93c05f48e64c0b3e845261010435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f14947a1e32f5f0e59a223d64fcdb96b0a9a3219510b8854c2190e02fce22fb6e828801527f2a9b8c17b0b5b6360d0171e97d51bc542f548cf034afed40d23849314e10e2e2845261012435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0aeb586f5d31b1a7b62b828776f329780eb01026ed2b729c2ab41e304a10d4cb828801527f2e709d6f77a352e49ade095639638b10e1a33c27145a88fbac6c110c236b205b845261014435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f08ce833b2dbd34cb8a4b00eaa45fe963d24cc944655b811d719c1d9583b4cd1d828801527f28c63c36fb110ab2f256c5913d7e7d8017cba906c430b45d19d71cebdf027226845261016435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2b771270bdc4661d0cb6107b47a52f79e334785d41ca5f066c8944b4fd95f219828801527f02fc865e5a17e8d38d82eadb62f7ee715d6004d46d9fe06be546f356ffbb5dd3845261018435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f115be55d78a138e0cc8215a42d0011d2a8a2d781ad16874cb4afda39908c4b0c828801527f0c8a39b25f66b912a58ea652cb9bf2c8149c8d51e438dd0159eee8b8b151ef9984526101a435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f118e76c9f149834d8975e7e1e3486a668cd4e89482d3a59d6f17d45484f62518828801527f2c78ab2ec3ba68f93116cf48477b740632664681f3482a81b64cf352fdf7edc584526101c435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611e245f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610c5e578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f13fd1dfd8781fe3bdb30f4e390a4716e4fadbf64763c88a9fb56b16c40117db36101008501527f15f3d4660c4a69d77b07fa91e84f5aa8b94126c93752988235f8a937171d36b86101208501527f0f2aab6e68ef0ea52b15742aa2772b47622dbfa1a7b7dd93cb322c5d469273cc6101408501527f273e7d3120556d9330dec620de3099b236484cef7136ad9f15b389174398a1de6101608501527f2413a48ec2fa419ed4181f34f111c06baf8dbdeaf4e1defc55d6a1fb0576a8646101808501527f1ea205ecaed0fc03f580a43223bdc715355ad3eff5410ef23e38e60c1ee84ebd6101a08501527f070c7681a0a0afc103d7a1f9c5e0aae92d0f69e06b8d788c433d9658a70b74f06101c08501527f2186619a41c57807574dbc85151946e7b1fcfeeaaf2a0197ada080fa1e45dd9e6101e08501527f276c9a9426abbf40a134b8b309f58a52212bf43f7f97fbdabc8a960815292f126102008501527f0510e091d41f64779ae660caee8b24f8292e285bccf04a38d8c3c0cbc642df496102208501526102408401526102608301527f02f80747a8cd73f5a51f44bf710dc73f3a03d37e671f625e5446ef862900c7946102808301527f1f4156641b4ba54b3eff74c6297ec1bdc2cebb65ede72c7999c5cef7e11557e86102a08301527e448785732e9cbfd68980bbc7a1432cd48839fb8f9f76ed84adba7f45ca91f46102c08301527f034d395cea82a76fe06b1ad9b29a0328ae0a6b57700d47b576f781f8f2538a176102e08301525192839161165784846116e2565b8336843760085afa1590811561166f575b50610c4f57005b6001915051141581611668565b34610236575f36600319011261023657807f90eaac8e23203b98b51f0470517654d24c269d3ffe72f0afb703e2ca9710103c60209252f35b9181601f840112156102365782359167ffffffffffffffff8311610236576020838186019501011161023657565b90601f8019910116810190811067ffffffffffffffff82111761024a57604052565b909160409282526020820152016060516080905f5b8181106117265750505090565b8251845260209384019390920191600101611719565b905f905b6002821061174d57505050565b6020806001928551815201930191019091611740565b80156117c7578060011c915f516020611e045f395f51905f52831015610c4f576001806117a65f516020611e045f395f51905f5260038188818180090908611c24565b9316146117af57565b905f516020611e045f395f51905f5280910681030690565b505f905f90565b8015806118ff575b6118f3578060021c92825f516020611e045f395f51905f5285108015906118dc575b610c4f5784815f516020611e045f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816118a69d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611c47565b809291600180829616146118b8575050565b5f516020611e045f395f51905f528093945080929550809106810306930681030690565b505f516020611e045f395f51905f528110156117f8565b50505f905f905f905f90565b5081156117d6565b905f516020611e045f395f51905f528210801590611991575b610c4f57811580611989575b611983576119505f516020611e045f395f51905f5260038185818180090908611c24565b81810361195f57505060011b90565b5f516020611e045f395f51905f52809106810306145f14610c4f57600190811b1790565b50505f90565b50801561192c565b505f516020611e045f395f51905f52811015611920565b919093925f516020611e045f395f51905f528310801590611bc5575b8015611bae575b8015611b97575b610c4f578082868517171715611b8c57908291611aef5f516020611e045f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611e045f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611ac981808b80098187800908611c24565b8408095f516020611e045f395f51905f52611ae382611d9b565b80091415958691611c47565b929080821480611b83575b15611b215750505050905f14611b195760ff60025b169060021b179190565b60ff5f611b0f565b5f516020611e045f395f51905f52809106810306149182611b64575b505015610c4f5760019115611b5c5760ff60025b169060021b17179190565b60ff5f611b51565b5f516020611e045f395f51905f52919250819006810306145f80611b3d565b50838314611afa565b50505090505f905f90565b505f516020611e045f395f51905f528110156119d2565b505f516020611e045f395f51905f528210156119cb565b505f516020611e045f395f51905f528510156119c4565b9080601f8301121561023657604080519290611bf890846116e2565b82906040810192831161023657905b828210611c145750505090565b8135815260209182019101611c07565b90611c2e82611d9b565b915f516020611e045f395f51905f5283800903610c4f57565b915f516020611e045f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c9f93969496611c9182808a8009818a800908611c24565b90611d8f575b860809611c24565b925f516020611e045f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611e045f395f51905f5260a083015260208260c08160055afa91519115610c4f575f516020611e045f395f51905f52826001920903610c4f575f516020611e045f395f51905f52908209925f516020611e045f395f51905f528080808780090681030681878009081490811591611d70575b50610c4f57565b90505f516020611e045f395f51905f528084860960020914155f611d69565b81809106810306611c97565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611e045f395f51905f5260a083015260208260c08160055afa91519115610c4f5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212200431386956d0501551a31c63561b5adb62e2a525e5b625e1ef36459447b6f0d864736f6c634300081c0033",
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
