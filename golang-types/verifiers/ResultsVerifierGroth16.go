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
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557611db3908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f5f3560e01c80630ea14a39146114de578063233ace11146114a4578063454c28a314610bdc5780635d26278e146101a85763b8e72af614610051575f80fd5b346101945760403660031901126101945760043567ffffffffffffffff811161019457610082903690600401611613565b60243567ffffffffffffffff8111610194576100a2903690600401611613565b61012090816040516100b48282611641565b36903782019080838303126101945781601f8401121561019457604051926100dc8285611641565b8391810192831161019457905b82821061019857505050303b156101945760405163454c28a360e01b8152610140600482015261014481018390529283919083906101648401375f6101648484010152602482015f905b6009821061017a57505050610164815f93601f80199101168101030181305afa801561016f57610161575080f35b61016d91505f90611641565b005b6040513d5f823e3d90fd5b829350602080916001939451815201930191018492610133565b5f80fd5b81358152602091820191016100e9565b34610194576101e03660031901126101945736608411610194573660a41161019457366101e4116101945760405160206101e28183611641565b8036833760408051906101f58183611641565b803683376103009381519261020a8685611641565b853685376102196084356119ba565b86830152815261022a60a4356119ba565b905f516020611d5e5f395f51905f528351888501518751908a82019283528882015261026a8161025c606082016116ba565b03601f198101835282611641565b5190200684528251865286830151878701527f298304a3901399d6cc9d84fd950893382804be563f92b894ecb20015b5fdb1c5858701527f037a3a2a5542679728288bd51a777a11609229503428c89a8597e7498692de1360608701527f11870a0152e3641b46ef3eb5ca4e43b36dfd9b1b7a8d3c3fd0ab95e0705a0b8060808701527f305567066f608e2423c3585ed83f29259a0d1e5c81f70e286513a0d43cc70e2760a087015260c086015260e08501527f1267ed180d91f6af83c0af45b72fb66fcd2dcbf2cecfa76b7b3e6d47c3bee4856101008501527f2f2b6eccd82560a4b5abeb031f53183541ab387b78ecf7cd781783a7474d9a7c6101208501527f29a03ac3e743d9420a91e8c09e66278b2de3dba6e299d5855cd913177dfcbdd16101408501527f2d69368d608314d31e50f91825382898d0c40d51bfb705c6ae49ab86ac6cc595610160850152825185816101808760085afa90511615610bcd5784906103da6004356119ba565b6103eb602495929535604435611a25565b92919093896103fb6064356119ba565b9890977f0e2f87d6f9c49226a9831458477066c3f626bb4285683f41b1be28e10bf8bf0f83519b8c937f238d9f75615471aea9a1dd8b4ac31b8bdd52640247a75e1b8a56dd68a5a2973185527f036dc7ee3ea1598590475c0700fd2b00aba35f059e8801117b6ccd4d4da05401828601528051868601520151606084019081527f1576c734b4a3986f6fbc43584156373a2f31289d18d06d9478995c4a579b4642858560808160065afa957f0c381443ea81c28574c71e316a81f1cceaccfb6b2fee33ffe2c8aef3543e531b818701527f148daf29588659c676856c2ccef1d732f8da4b0a21fc76f274d3013c3bcdf9048352600160c4356080880198818a525f516020611d5e5f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f20ac6b974973dfcb181ca2e24a79a3dc6f926809f11fa8763dad2aea6e039bd1828801527f29ce0b801022051b8ea3d251186317bbbebd4edf2875f85927b4a1507ebbf4fb845260e435908189525f516020611d5e5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f20b44e45ef3992f07d29637d7530e962b305092971f88960599660e953a8129a828801527f0d37b99ede63c5d06aab826a22c8dad9327c8a825e9778a97f0d5c0ff640fc42845261010435908189525f516020611d5e5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f251370c6ead16fd71b384261752afb7b73aaeb0d70fa3004713c4b3123f5664e828801527f2b5a72ac864f630bf2363f03041d9717c22051ed7552606ca172b8fc7ee37ce2845261012435908189525f516020611d5e5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f094fd56ccc740a0e079b3c4b955e40b6e89f1159b8a27659680c6e11b8c5a06c828801527f0265c0205ef83e2f103c71854f7874b4ef5a235f42178435224c292b1328c83c845261014435908189525f516020611d5e5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f18da4bd19d18200f7256d8dc5468c3ffb96960314e89d0fd2f0dfc6a2137fcab828801527f1a226ffdf3b3c500926c967810e80a23ea9eb6e6e3032bf5d4a1cab1f0774fdf845261016435908189525f516020611d5e5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2426cabab4e3a6e187c24d3c4d216d61b4a6715a04870b73bb22a3bbab123b68828801527f2589ec79e3e09ff9f8f4f57c341df6b792f0590f8c852bc5b6a6ceaed5588110845261018435908189525f516020611d5e5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f128a9eede0ee3b79c9d924851fdd03323e087af96a2490918cd1efb646d18f03828801527f02dade84d586c2a78f9d9f70687bbc6a2d7907c8ed83ec9b139fe3afe408455984526101a435908189525f516020611d5e5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f0f5ee05c324de2c97f881689fe61402dd7f107a7bff0b86b9b2826cce72e6099828801527f081a67b415fd8e00cb2413ade5b1089383a21a186c600c8b4d9f5686c1e79ba784526101c435908189525f516020611d5e5f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611d5e5f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610bbe578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f19a5bf158a4a0d3867b31f67b6dda477b33a1de4d446b2910db36c885bb0a5676101008501527f0658bc02f23b3a35327333f72f3cb6df5d5b7673375fb63422a9a47ab7d080a66101208501527f260fb6841ebe557b9442b00c10e7c272562518c9edaa1c858d0277cdbee40c736101408501527f29e0e4631517e6256d5c0fb34fe61623da0ae7c36f84a8f5ff9679f53bcbfe526101608501527f25f2f2d8c80a9074a219124f7280263e72dce6f4303a89cbe7b9a5b9fef9b6b56101808501527f0d7a5ebf22930951d8a91322c8870027dc8a7462dc4022c091975022eb20244d6101a08501527f1843de734b6931a9705a8f6452b3d643e21cb5ab2eb7783dac3f666c74cb932d6101c08501527f0865f16b17a9a70dbd673c8f14f741a925df9db6adc911a13c30c38e0db52cca6101e08501527f13a08e71081e636e2baeb28d6469ca007f8f3bcef20e5c604adfec2aaee410856102008501527f279ae656d03fecefa67719deb9ff85404a43fdc8d2543b2e6bd98f1d9378e57a6102208501526102408401526102608301527f067720c3b8d9a327ae2455485d23864b0e8be3d8f37f7ddc6385b2c35c07c3c06102808301527f139d93f9a9ff7a6aedf90f8024b0db44d8fbcde7a42753edc063341e0e9a39616102a08301527f01cb5b0313cef901575c5137dce7e25e0ce9de842e53b03b865955b8cbcaef2d6102c08301527f25f561f8b81a3a069ba7623e0def6871fbc6c69daac16f54340754876803bf3c6102e083015251928391610b8a8484611641565b8336843760085afa15908115610bb1575b50610ba257005b631ff3747d60e21b5f5260045ffd5b6001915051141581610b9b565b63a54f8e2760e01b5f5260045ffd5b6351d49ff760e11b5f5260045ffd5b34610194576101403660031901126101945760043567ffffffffffffffff811161019457610c0e903690600401611613565b366101441161019457610c246101808214611677565b6040908151610c338382611641565b8236823761010084019280848337602093815193610c518686611641565b8536863761014011610194575f516020611d5e5f395f51905f5282518681019084848337610c888161025c606082015f81526116ba565b5190200684528180519182377f298304a3901399d6cc9d84fd950893382804be563f92b894ecb20015b5fdb1c5828201527f037a3a2a5542679728288bd51a777a11609229503428c89a8597e7498692de1360608201527f11870a0152e3641b46ef3eb5ca4e43b36dfd9b1b7a8d3c3fd0ab95e0705a0b8060808201527f305567066f608e2423c3585ed83f29259a0d1e5c81f70e286513a0d43cc70e2760a082015281610140870160c08301377f1267ed180d91f6af83c0af45b72fb66fcd2dcbf2cecfa76b7b3e6d47c3bee4856101008201527f2f2b6eccd82560a4b5abeb031f53183541ab387b78ecf7cd781783a7474d9a7c6101208201527f29a03ac3e743d9420a91e8c09e66278b2de3dba6e299d5855cd913177dfcbdd16101408201527f2d69368d608314d31e50f91825382898d0c40d51bfb705c6ae49ab86ac6cc59561016082015284816101808160085afa90511615610bcd578051928184017f238d9f75615471aea9a1dd8b4ac31b8bdd52640247a75e1b8a56dd68a5a2973185525f516020611d5e5f395f51905f528387808801967f036dc7ee3ea1598590475c0700fd2b00aba35f059e8801117b6ccd4d4da0540188528051855201519260608801938452818860808160065afa947f0c381443ea81c28574c71e316a81f1cceaccfb6b2fee33ffe2c8aef3543e531b82527f148daf29588659c676856c2ccef1d732f8da4b0a21fc76f274d3013c3bcdf90485527f0e2f87d6f9c49226a9831458477066c3f626bb4285683f41b1be28e10bf8bf0f600160243560808c0198818a5287878760608160075afa92101616858c60808160065afa16167f20ac6b974973dfcb181ca2e24a79a3dc6f926809f11fa8763dad2aea6e039bd184527f29ce0b801022051b8ea3d251186317bbbebd4edf2875f85927b4a1507ebbf4fb87526044359081895286868660608160075afa92101616848b60808160065afa167f20b44e45ef3992f07d29637d7530e962b305092971f88960599660e953a8129a84527f0d37b99ede63c5d06aab826a22c8dad9327c8a825e9778a97f0d5c0ff640fc4287526064359081895286868660608160075afa92101616848b60808160065afa167f251370c6ead16fd71b384261752afb7b73aaeb0d70fa3004713c4b3123f5664e84527f2b5a72ac864f630bf2363f03041d9717c22051ed7552606ca172b8fc7ee37ce287526084359081895286868660608160075afa92101616848b60808160065afa167f094fd56ccc740a0e079b3c4b955e40b6e89f1159b8a27659680c6e11b8c5a06c84527f0265c0205ef83e2f103c71854f7874b4ef5a235f42178435224c292b1328c83c875260a4359081895286868660608160075afa92101616848b60808160065afa167f18da4bd19d18200f7256d8dc5468c3ffb96960314e89d0fd2f0dfc6a2137fcab84527f1a226ffdf3b3c500926c967810e80a23ea9eb6e6e3032bf5d4a1cab1f0774fdf875260c4359081895286868660608160075afa92101616848b60808160065afa167f2426cabab4e3a6e187c24d3c4d216d61b4a6715a04870b73bb22a3bbab123b6884527f2589ec79e3e09ff9f8f4f57c341df6b792f0590f8c852bc5b6a6ceaed5588110875260e4359081895286868660608160075afa92101616848b60808160065afa167f128a9eede0ee3b79c9d924851fdd03323e087af96a2490918cd1efb646d18f0384527f02dade84d586c2a78f9d9f70687bbc6a2d7907c8ed83ec9b139fe3afe40845598752610104359081895286868660608160075afa92101616848b60808160065afa167f0f5ee05c324de2c97f881689fe61402dd7f107a7bff0b86b9b2826cce72e609984527f081a67b415fd8e00cb2413ade5b1089383a21a186c600c8b4d9f5686c1e79ba78752610124359081895286868660608160075afa92101616848b60808160065afa16957f1576c734b4a3986f6fbc43584156373a2f31289d18d06d9478995c4a579b46428452525180955260608160075afa92101616818460808160065afa16925191519215610bbe5761010090519485377f19a5bf158a4a0d3867b31f67b6dda477b33a1de4d446b2910db36c885bb0a5676101008501527f0658bc02f23b3a35327333f72f3cb6df5d5b7673375fb63422a9a47ab7d080a66101208501527f260fb6841ebe557b9442b00c10e7c272562518c9edaa1c858d0277cdbee40c736101408501527f29e0e4631517e6256d5c0fb34fe61623da0ae7c36f84a8f5ff9679f53bcbfe526101608501527f25f2f2d8c80a9074a219124f7280263e72dce6f4303a89cbe7b9a5b9fef9b6b56101808501527f0d7a5ebf22930951d8a91322c8870027dc8a7462dc4022c091975022eb20244d6101a08501527f1843de734b6931a9705a8f6452b3d643e21cb5ab2eb7783dac3f666c74cb932d6101c08501527f0865f16b17a9a70dbd673c8f14f741a925df9db6adc911a13c30c38e0db52cca6101e08501527f13a08e71081e636e2baeb28d6469ca007f8f3bcef20e5c604adfec2aaee410856102008501527f279ae656d03fecefa67719deb9ff85404a43fdc8d2543b2e6bd98f1d9378e57a6102208501526102408401526102608301527f067720c3b8d9a327ae2455485d23864b0e8be3d8f37f7ddc6385b2c35c07c3c06102808301527f139d93f9a9ff7a6aedf90f8024b0db44d8fbcde7a42753edc063341e0e9a39616102a08301527f01cb5b0313cef901575c5137dce7e25e0ce9de842e53b03b865955b8cbcaef2d6102c08301527f25f561f8b81a3a069ba7623e0def6871fbc6c69daac16f54340754876803bf3c6102e0830152816103008160085afa90511615610ba257005b34610194575f3660031901126101945760206040517f6323dadb3595fca48b4cb35b1e77229d79a128a4786115f33847cddcdc941ab48152f35b346101945760203660031901126101945760043567ffffffffffffffff81116101945761150f903690600401611613565b906080604051916115208284611641565b813684376115ba6020916115486101806040519761153e868a611641565b85368a3714611677565b6115568382013582356116e5565b85526115738482013560a083013560408401356060850135611786565b84870152604086015261158e60e082013560c08301356116e5565b60608601526115a76101208201356101008301356116e5565b86526101406101608201359101356116e5565b9160405193845f905b600482106115fd57505050830193905f945b600186106115e85760c0858560a0820152f35b818060019285518152019301950194916115d5565b82518152918401916001919091019084016115c3565b9181601f840112156101945782359167ffffffffffffffff8311610194576020838186019501011161019457565b90601f8019910116810190811067ffffffffffffffff82111761166357604052565b634e487b7160e01b5f52604160045260245ffd5b1561167e57565b60405162461bcd60e51b81526020600482015260146024820152730d2dcecc2d8d2c840e0e4dedecc40d8cadccee8d60631b6044820152606490fd5b6060516080905f5b8181106116cf5750505090565b82518452602093840193909201916001016116c2565b905f516020611d3e5f395f51905f52821080159061176f575b610ba257811580611767575b6117615761172e5f516020611d3e5f395f51905f5260038185818180090908611b5e565b81810361173d57505060011b90565b5f516020611d3e5f395f51905f52809106810306145f14610ba257600190811b1790565b50505f90565b50801561170a565b505f516020611d3e5f395f51905f528110156116fe565b919093925f516020611d3e5f395f51905f5283108015906119a3575b801561198c575b8015611975575b610ba257808286851717171561196a579082916118cd5f516020611d3e5f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611d3e5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4816118a781808b80098187800908611b5e565b8408095f516020611d3e5f395f51905f526118c182611cd5565b80091415958691611b81565b929080821480611961575b156118ff5750505050905f146118f75760ff60025b169060021b179190565b60ff5f6118ed565b5f516020611d3e5f395f51905f52809106810306149182611942575b505015610ba2576001911561193a5760ff60025b169060021b17179190565b60ff5f61192f565b5f516020611d3e5f395f51905f52919250819006810306145f8061191b565b508383146118d8565b50505090505f905f90565b505f516020611d3e5f395f51905f528110156117b0565b505f516020611d3e5f395f51905f528210156117a9565b505f516020611d3e5f395f51905f528510156117a2565b8015611a1e578060011c915f516020611d3e5f395f51905f52831015610ba2576001806119fd5f516020611d3e5f395f51905f5260038188818180090908611b5e565b931614611a0657565b905f516020611d3e5f395f51905f5280910681030690565b505f905f90565b801580611b56575b611b4a578060021c92825f516020611d3e5f395f51905f528510801590611b33575b610ba25784815f516020611d3e5f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4481611afd9d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611b81565b80929160018082961614611b0f575050565b5f516020611d3e5f395f51905f528093945080929550809106810306930681030690565b505f516020611d3e5f395f51905f52811015611a4f565b50505f905f905f905f90565b508115611a2d565b90611b6882611cd5565b915f516020611d3e5f395f51905f5283800903610ba257565b915f516020611d3e5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611bd993969496611bcb82808a8009818a800908611b5e565b90611cc9575b860809611b5e565b925f516020611d3e5f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611d3e5f395f51905f5260a083015260208260c08160055afa91519115610ba2575f516020611d3e5f395f51905f52826001920903610ba2575f516020611d3e5f395f51905f52908209925f516020611d3e5f395f51905f528080808780090681030681878009081490811591611caa575b50610ba257565b90505f516020611d3e5f395f51905f528084860960020914155f611ca3565b81809106810306611bd1565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611d3e5f395f51905f5260a083015260208260c08160055afa91519115610ba25756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220719de4f4412728eebcb3a9de06929810d293c2cc007c1f108f38d820d2b257bf64736f6c634300081c0033",
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

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) CompressProof(opts *bind.CallOpts, proof []byte) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "compressProof", proof)

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
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) CompressProof(proof []byte) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _ResultsVerifierGroth16.Contract.CompressProof(&_ResultsVerifierGroth16.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x0ea14a39.
//
// Solidity: function compressProof(bytes proof) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) CompressProof(proof []byte) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _ResultsVerifierGroth16.Contract.CompressProof(&_ResultsVerifierGroth16.CallOpts, proof)
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

// VerifyProof is a free data retrieval call binding the contract method 0x454c28a3.
//
// Solidity: function verifyProof(bytes proof, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) VerifyProof(opts *bind.CallOpts, proof []byte, input [9]*big.Int) error {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x454c28a3.
//
// Solidity: function verifyProof(bytes proof, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) VerifyProof(proof []byte, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof(&_ResultsVerifierGroth16.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x454c28a3.
//
// Solidity: function verifyProof(bytes proof, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) VerifyProof(proof []byte, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof(&_ResultsVerifierGroth16.CallOpts, proof, input)
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
