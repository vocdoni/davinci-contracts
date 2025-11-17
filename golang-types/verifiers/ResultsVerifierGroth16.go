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
	Bin: "0x6080604052348015600e575f5ffd5b506120c58061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063233ace11146100595780635d26278e1461008c57806360e58346146100a1578063b1c3a00e146100b4578063b8e72af6146100d6575b5f5ffd5b6040517fe33f5ca1f797e839cc3a4d40825b95db5faa8ed80554b428b9cb85e6f19d348a81526020015b60405180910390f35b61009f61009a366004611b75565b6100e9565b005b61009f6100af366004611beb565b61061d565b6100c76100c2366004611c39565b610a92565b60405161008393929190611c7f565b61009f6100e4366004611d25565b610b1a565b6100f1611aad565b6100f9611acb565b610101611ae9565b610111865f5b6020020135610b95565b602084015282525f8061012387610b95565b85516020808801516040519496509294506060935f5160206120705f395f51905f529361015593929091869101611da5565b604051602081830303815290604052805190602001205f1c6101779190611dfa565b865284518452602085810151858201527f1dd2d737ab46c2194a5a95bd34f2e953d3aa8abd0b39187391bcf74a9f471f4a6040808701919091527f1252c38c6d66d28edfb9a6bb816b24d6a4ea37eeb59c5ed7710019d1dc618f3460608701527f02f17c14384620127c4b3bba0f609a3845982b7af70062e1d9fbf8b20476ecb060808701527f207e2afe2d247bcb8d4f032ae37f35d6a6eeae3a0a155e5972569b3c81d11a1560a087015260c0860185905260e086018490527f21575cd429d060aaed57f941f5ad767db3761711e8f970c166fa8c986a0bbc686101008701527f2db56c68ef5f6031af05aae69e105e145e3934f9a54442da393d78985bdcf5676101208701527f17e916d35cbc18bad9ec08c9c2e3566c79d27074434e1214e9ca3fb040b4b2346101408701527f227b33619cdfa0a21e1a005a170015dc17d2d37faa2c544bdf8c7f3d32be442a610160870152515f91816101808860085afa9051169050806102fc576040516351d49ff760e11b815260040160405180910390fd5b505050505f5f610317895f6004811061010757610107611d91565b90925090505f80808061033260408e013560208f0135610c34565b929650909450925090505f806103498f6003610107565b915091505f5f61035a8e8e8e610e08565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f22f5fb0f00bc29881b9b0e1bbc5f055f581c4af4af0431208b537e8581d2e8d16101008e01527f0e0b10eb1e7cfa779792f8ef7c5d72907884b50d07879b36a3a0ac0c8fa50cb56101208e01527f01ac60926c9ab278d919a8d6be74d07a71a76f34e533c5173dc8a20732883a656101408e01527f0b1e64a83bbf3e7c460dbd717ca69e605d65e25adcea8629f03aff318cda08ba6101608e01527f2d348f40c9f808b1a15017628a67e2f3453ecea5c9af46fa68c6da4981e13dba6101808e01527f0f2b6c5c14e04b01ffcd2d952169bfadd6105fb076de2536c7d5e553fdc11ad66101a08e01527f21a68f5d5c2edacd1ec3fc53a68a3af01a85c0e3e6edf486795c10aeca2aff606101c08e01527f21fe42bef58082093a7acfa9fbdb76fc986a1670ce1339a11a731f4eb068cfc86101e08e01527f1a6ce968ceba9ceaa969f7032b3d1d059c37e222be9ffd6639974ccd8061a07c6102008e01527f10e8ab55d4642a601726e44a4ab1289331d5a9a979069c8bec903d482bdf6d466102208e01526102408d018290526102608d018190527f286f777cab86b1a273eddbff655bd579d965962ae85dd87799dd548a8f845bbf6102808e01527f2c742a2d2db39d912270d5a19dfe417d804bf23e4957e5adddb4e3ae1d9440606102a08e01527f0ee12caa65c102aee85548d0693630d4fb44b001667db3d4941b0c58f8d339386102c08e01527f2e84812a2a2d0ba99eac0708058df57dfb4ee79817bf2b2b6ac9861620d95a446102e08e015290925090505f6105ce611aad565b6020816103008f60085afa91508115806105ea57508051600114155b1561060857604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b610625611aad565b6040516060905f5160206120705f395f51905f5290610650908735906020808a013591869101611da5565b604051602081830303815290604052805190602001205f1c6106729190611dfa565b8252604080515f918782377f1dd2d737ab46c2194a5a95bd34f2e953d3aa8abd0b39187391bcf74a9f471f4a60408201527f1252c38c6d66d28edfb9a6bb816b24d6a4ea37eeb59c5ed7710019d1dc618f3460608201527f02f17c14384620127c4b3bba0f609a3845982b7af70062e1d9fbf8b20476ecb060808201527f207e2afe2d247bcb8d4f032ae37f35d6a6eeae3a0a155e5972569b3c81d11a1560a082015260408660c08301377f21575cd429d060aaed57f941f5ad767db3761711e8f970c166fa8c986a0bbc686101008201527f2db56c68ef5f6031af05aae69e105e145e3934f9a54442da393d78985bdcf5676101208201527f17e916d35cbc18bad9ec08c9c2e3566c79d27074434e1214e9ca3fb040b4b2346101408201527f227b33619cdfa0a21e1a005a170015dc17d2d37faa2c544bdf8c7f3d32be442a6101608201526020816101808360085afa9051169050806107e7576040516351d49ff760e11b815260040160405180910390fd5b5f5f61081c86868a60028060200260405190810160405280929190826002602002808284375f92019190915250610e08915050565b915091506040516101008a82377f22f5fb0f00bc29881b9b0e1bbc5f055f581c4af4af0431208b537e8581d2e8d16101008201527f0e0b10eb1e7cfa779792f8ef7c5d72907884b50d07879b36a3a0ac0c8fa50cb56101208201527f01ac60926c9ab278d919a8d6be74d07a71a76f34e533c5173dc8a20732883a656101408201527f0b1e64a83bbf3e7c460dbd717ca69e605d65e25adcea8629f03aff318cda08ba6101608201527f2d348f40c9f808b1a15017628a67e2f3453ecea5c9af46fa68c6da4981e13dba6101808201527f0f2b6c5c14e04b01ffcd2d952169bfadd6105fb076de2536c7d5e553fdc11ad66101a08201527f21a68f5d5c2edacd1ec3fc53a68a3af01a85c0e3e6edf486795c10aeca2aff606101c08201527f21fe42bef58082093a7acfa9fbdb76fc986a1670ce1339a11a731f4eb068cfc86101e08201527f1a6ce968ceba9ceaa969f7032b3d1d059c37e222be9ffd6639974ccd8061a07c6102008201527f10e8ab55d4642a601726e44a4ab1289331d5a9a979069c8bec903d482bdf6d4661022082015282610240820152816102608201527f286f777cab86b1a273eddbff655bd579d965962ae85dd87799dd548a8f845bbf6102808201527f2c742a2d2db39d912270d5a19dfe417d804bf23e4957e5adddb4e3ae1d9440606102a08201527f0ee12caa65c102aee85548d0693630d4fb44b001667db3d4941b0c58f8d339386102c08201527f2e84812a2a2d0ba99eac0708058df57dfb4ee79817bf2b2b6ac9861620d95a446102e08201526020816103008360085afa905116925082610a8757604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b610a9a611b08565b610aa2611aad565b5f610ab686358760015b60200201356113e5565b8352610ad46060870135604088013560a089013560808a01356114d2565b60208501526040840152610aee60c0870135876007610aac565b6060840152610b008535866001610aac565b8252610b0f8435856001610aac565b905093509350939050565b5f5f5f610b2787876117ab565b91945092509050306360e58346848484610b418a8a6117dc565b6040518563ffffffff1660e01b8152600401610b609493929190611e41565b5f6040518083038186803b158015610b76575f5ffd5b505afa158015610b88573d5f5f3e3d5ffd5b5050505050505050505050565b5f5f825f03610ba857505f928392509050565b600183811c9250808416145f5160206120505f395f51905f528310610be057604051631ff3747d60e21b815260040160405180910390fd5b610c1a5f5160206120505f395f51905f5260035f5160206120505f395f51905f52865f5160206120505f395f51905f5288890909086117f7565b91508015610c2e57610c2b82611859565b91505b50915091565b5f80808085158015610c44575084155b15610c5957505f925082915081905080610dff565b600286811c945085935060018088161490808816145f5160206120505f395f51905f5286101580610c9757505f5160206120505f395f51905f528510155b15610cb557604051631ff3747d60e21b815260040160405180910390fd5b5f5f5160206120505f395f51905f52610cdc60035f5160206120505f395f51905f52611ebb565b5f5160206120505f395f51905f52888a090990505f5f5160206120505f395f51905f52885f5160206120505f395f51905f528a8b090990505f5f5160206120505f395f51905f52885f5160206120505f395f51905f528a8b090990505f5160206120505f395f51905f52805f5160206120505f395f51905f528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089650610dca5f5160206120505f395f51905f52805f5160206120505f395f51905f528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611859565b9550610dd7878786611871565b90975095508415610df957610deb87611859565b9650610df686611859565b95505b50505050505b92959194509250565b5f5f5f60019050604051604081015f7f0aa8bb909f118e5e4e3c3dc3e89aed474cebd51de85fa49ab6b365a2617dab3f83527f08dcf65136d7c456420199cf436a0779b964112131a81f8484e1901b4c22239f6020840152865182526020870151602083015260408360808560065afa841693507f052ea9a4be49fa45739b3eebacb7b817456d74c6a887b36686662da43a2f972b82527f25e7d03358507d01732e90c60538f53000cf8e9c253b8d6aca8a31e42dc99db96020830152883590508060408301525f5160206120705f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f07d44661f2dee3868df1485aae5fcbdcd1eb82ddedbaa5012f122a9c5f79a4af82527f0578aeeee33235bc096dddc669f2d3c5d9ced038aa0acf0ea32d25f982e1a7966020830152602089013590508060408301525f5160206120705f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f0c6381dfdc349ef815a9ae4aa4aa7ed9b043953ab3a320b524822e1084005ce682527f107fcb5b4ea1a91d3f24ae4c40cdd8da272a131dd3bae45393a9c0b8be3d8a746020830152604089013590508060408301525f5160206120705f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f22a69f0ac5723696ab42865b6167b6cd064333d9e2c98d40e78815db5a749a9f82527f04ebc77d3c4a3e20122adc6f08f19174f30ae6facfb2642013c2ad922be43ce96020830152606089013590508060408301525f5160206120705f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f2940af5fff761c4c03ad244734d638068cb14394fd5f96c531bf08754b724baa82527f238c952451425908482823ad53ff41557d24aa9779612099d32d53c9c10984566020830152608089013590508060408301525f5160206120705f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f0e98ac5e49c22ee36382444c3d8d3ba146db8b1f63ffa9154d9fca6fc86dd4ce82527f1703022f1bb3b812fd6f82f3631f7d1851a42a0073086b8e7f597fb4cde5783d602083015260a089013590508060408301525f5160206120705f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f28afe3fad233ae7dd0a554c4deb8b8b2f2e270328dcf531d6aaf4345c4dd658d82527f2589c9b89b6946297400fe35437c0c5dc738dd09dd0848ad7765b7855bd156e0602083015260c089013590508060408301525f5160206120705f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f0571cb2117e23e90595e1ddfc3ab1199bfbec6255be35579a3fc87b352054fa682527f0943bd119ab33aa4540d8e86c2d272d59da000927dddd86789f841bb3b821cc9602083015260e089013590508060408301525f5160206120705f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f0cd0a418c6cd6a3e0182690fc3ea5d2ce163752a52ee1ff33b1350014c96807e82527f0370eaedb6d8aa5979e4cd732073e23eeae9538d2950393a9b739342164341bf602083015261010089013590508060408301525f5160206120705f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa7f0bfdbe6bcf952a6225d765c751234d6a9efdd733e40f3c91609c86f11449a9f283527f0b4cb2ac16bec4711ab861cd128f44a8322c980496c1910eba7c025cccc194516020840152885160408085018290525f5160206120705f395f51905f5290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806113dc5760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b5f5f5160206120505f395f51905f528310158061140f57505f5160206120505f395f51905f528210155b1561142d57604051631ff3747d60e21b815260040160405180910390fd5b82158015611439575081155b1561144557505f6114cc565b5f6114805f5160206120505f395f51905f5260035f5160206120505f395f51905f52875f5160206120505f395f51905f52898a0909086117f7565b9050808303611495575050600182901b6114cc565b61149e81611859565b83036114b1575050600182811b176114cc565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b5f5f5f5160206120505f395f51905f52861015806114fd57505f5160206120505f395f51905f528510155b8061151557505f5160206120505f395f51905f528410155b8061152d57505f5160206120505f395f51905f528310155b1561154b57604051631ff3747d60e21b815260040160405180910390fd5b828486881717175f0361156257505f9050806117a2565b5f80805f5160206120505f395f51905f5261158b60035f5160206120505f395f51905f52611ebb565b5f5160206120505f395f51905f528a8c090990505f5f5160206120505f395f51905f528a5f5160206120505f395f51905f528c8d090990505f5f5160206120505f395f51905f528a5f5160206120505f395f51905f528c8d090990505f5160206120505f395f51905f52805f5160206120505f395f51905f528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50894506116795f5160206120505f395f51905f52805f5160206120505f395f51905f528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611859565b93505050505f5f6116c65f5160206120505f395f51905f528061169e5761169e611de6565b5f5160206120505f395f51905f528586095f5160206120505f395f51905f52878809086117f7565b90506117115f5160206120505f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f5160206120505f395f51905f52848808096119a4565b15915050611720838383611871565b9093509150868314801561173357508186145b1561175b5780611743575f611746565b60025b60ff1660028a901b175f17945087935061179e565b61176483611859565b87148015611779575061177682611859565b86145b156114b15780611789575f61178c565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b6117b3611b26565b6117bb611acb565b6117c3611acb565b6117cf84860186611f69565b9250925092509250925092565b6117e4611b45565b6117f082840184611fe8565b9392505050565b5f611822827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526119ec565b9050815f5160206120505f395f51905f528283091461185457604051631ff3747d60e21b815260040160405180910390fd5b919050565b5f5160206120505f395f51905f529081900681030690565b5f80806118a05f5160206120505f395f51905f52808788095f5160206120505f395f51905f52898a09086117f7565b905083156118b4576118b181611859565b90505b6118fd5f5160206120505f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f5160206120505f395f51905f52848a08096117f7565b92505f5160206120505f395f51905f526119275f5160206120505f395f51905f5260028609611a4f565b860991505f5160206120505f395f51905f526119525f5160206120505f395f51905f52848509611859565b5f5160206120505f395f51905f52858609088614158061198657505f5160206120505f395f51905f52808385096002098514155b156113dc57604051631ff3747d60e21b815260040160405180910390fd5b5f5f6119d0837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526119ec565b9050825f5160206120505f395f51905f52828309149392505050565b5f5f60405160208152602080820152602060408201528460608201528360808201525f5160206120505f395f51905f5260a082015260208160c08360055afa905192509050806114ca57604051631ff3747d60e21b815260040160405180910390fd5b5f611a7a827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd456119ec565b90505f5160206120505f395f51905f5281830960011461185457604051631ff3747d60e21b815260040160405180910390fd5b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b8061012081018310156114cc575f5ffd5b5f5f5f5f6101e08587031215611b89575f5ffd5b6080850186811115611b99575f5ffd5b85945060a0860187811115611bac575f5ffd5b909350359150611bbf8660c08701611b64565b905092959194509250565b8061010081018310156114cc575f5ffd5b80604081018310156114cc575f5ffd5b5f5f5f5f6102a08587031215611bff575f5ffd5b611c098686611bca565b9350611c19866101008701611bdb565b9250611c29866101408701611bdb565b9150611bbf866101808701611b64565b5f5f5f6101808486031215611c4c575f5ffd5b611c568585611bca565b9250611c66856101008601611bdb565b9150611c76856101408601611bdb565b90509250925092565b60c0810181855f5b6004811015611ca6578151835260209283019290910190600101611c87565b50505060808201845f5b6001811015611ccf578151835260209283019290910190600101611cb0565b5050508260a0830152949350505050565b5f5f83601f840112611cf0575f5ffd5b50813567ffffffffffffffff811115611d07575f5ffd5b602083019150836020828501011115611d1e575f5ffd5b9250929050565b5f5f5f5f60408587031215611d38575f5ffd5b843567ffffffffffffffff811115611d4e575f5ffd5b611d5a87828801611ce0565b909550935050602085013567ffffffffffffffff811115611d79575f5ffd5b611d8587828801611ce0565b95989497509550505050565b634e487b7160e01b5f52603260045260245ffd5b8381528260208201525f604082018351602085015f5b82811015611dd9578151845260209384019390910190600101611dbb565b5091979650505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f82611e1457634e487b7160e01b5f52601260045260245ffd5b500690565b805f5b6002811015611e3b578151845260209384019390910190600101611e1c565b50505050565b6102a0810181865f5b6008811015611e69578151835260209283019290910190600101611e4a565b505050611e7a610100830186611e19565b611e88610140830185611e19565b6101808201835f5b6009811015611eaf578151835260209283019290910190600101611e90565b50505095945050505050565b818103818111156114cc57634e487b7160e01b5f52601160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611f0f57634e487b7160e01b5f52604160045260245ffd5b604052919050565b5f82601f830112611f26575f5ffd5b5f611f316040611eda565b9050806040840185811115611f44575f5ffd5b845b81811015611f5e578035835260209283019201611f46565b509195945050505050565b5f5f5f6101808486031215611f7c575f5ffd5b5f85601f860112611f8b575f5ffd5b505f80610100611f9a81611eda565b9250829150860187811115611fad575f5ffd5b865b81811015611fc7578035845260209384019301611faf565b50819550611fd58882611f17565b9450505050611c76856101408601611f17565b5f6101208284031215611ff9575f5ffd5b5f83601f840112612008575f5ffd5b505f8061012061201781611eda565b925082915084018581111561202a575f5ffd5b845b8181101561204457803584526020938401930161202c565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212203f953ab72b3cbe78b4be4a86902c68b957d21ee19d0770408d48d433c937869564736f6c634300081c0033",
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
