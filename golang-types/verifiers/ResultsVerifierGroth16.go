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
	Bin: "0x6080604052348015600e575f5ffd5b506120c38061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063233ace11146100595780635d26278e1461008c57806360e58346146100a1578063b1c3a00e146100b4578063b8e72af6146100d6575b5f5ffd5b6040517f8d0cee3417f4cff975938a3e3b3a13d0e7b55bce16e6a85eb7b9c7cfb0dd4ee581526020015b60405180910390f35b61009f61009a366004611b73565b6100e9565b005b61009f6100af366004611be9565b61061c565b6100c76100c2366004611c37565b610a90565b60405161008393929190611c7d565b61009f6100e4366004611d23565b610b18565b6100f1611aab565b6100f9611ac9565b610101611ae7565b610111865f5b6020020135610b93565b602084015282525f8061012387610b93565b85516020808801516040519496509294506060935f51602061206e5f395f51905f529361015593929091869101611da3565b604051602081830303815290604052805190602001205f1c6101779190611df8565b865284518452602085810151858201527f1e70ab38629f5091b9e2f7d72e931df3a285af6b2dff8d602f9b7ec0b8f5af106040808701919091527f2e800f8970e7a46ee0132beecdeb739edfff8e373ecbf4068a1b7ae7d17b34fc60608701527f2968f4ca2db982db1b5a16762a88f675300f6e47be3a3f19bc4bbbdc42a08f0c60808701527f2014567dc6005885e09299cd19edf0606146c4f459fc321e1064de1a279877e260a087015260c0860185905260e086018490527f2b12e8ed1596ee4c70865cb4c3e419ef4858310f62188e8fdce45f1e8b1c4d0c6101008701527f1d81472f5650baced744698692994b59855e2e3f0a7bff9004b2f434c985b0886101208701527e18a8e18a55f23b496e11dd2bbe68fbd301642c0ea7765f0ae33cd135cd23556101408701527f03b8fa16a50162cd5334d3b4f1eb38f49814d21722ce6774efb7c0f2b449bc63610160870152515f91816101808860085afa9051169050806102fb576040516351d49ff760e11b815260040160405180910390fd5b505050505f5f610316895f6004811061010757610107611d8f565b90925090505f80808061033160408e013560208f0135610c32565b929650909450925090505f806103488f6003610107565b915091505f5f6103598e8e8e610e06565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f2167539e87fb7b1884dcae8b54954bdb7ec02d36b1ec8ccefb063151b246e0456101008e01527f08f359e93f4b8ad59f17b0851f63f63f58909ad4f668a49701c66435149d59fb6101208e01527f15229d2bee0319b7788f85517363b4825d05080c4f80996af5fb46f8368dcda66101408e01527f09514d166b708ffb95997005e9cbce7cc620b27317609236c5893210286670806101608e01527f16976d04c68c156f76cdf58700d71a261c07f1677c56c640a17948af62d24c596101808e01527f233b312bf47011f1e6759e0d93161be46ac4fb5f2c0070d9e5c97c54221cd2936101a08e01527f1514e031e0ef5e8cbe9aa455877dc5f039deb39446e26a917bb09ffb53455a636101c08e01527f2d3788e0f54f19aa6602a0fe3eb620460276da8bd625d7087e0ec9e27ad08c006101e08e01527f1eab641f675c11fc79264f5fc92216816438f6613a274983ee3ae7c0c01e7d116102008e01527f1484dd9cb4945159d105f0a02b2e0b5029d757f1012dd9b52fcb38837319aec16102208e01526102408d018290526102608d018190527f0ff2c0fe8c8f9bae6970ddd1682a3d7203dbc2c2fa4da33527957380e54977fb6102808e01527f2b06c1da47fb4ad6bca422847059d08e40b7a0355b400a9a35e3abb2f9e3e9386102a08e01527f2c4d7a7b5b4c3ca5ce1738a505f6228cb6cf7aa54740387812fe13816f691e3d6102c08e01527f2a793b258a478fb8be8a5a79e5c1fcae60abd2ca4d13af8550045beaa4974bfe6102e08e015290925090505f6105cd611aab565b6020816103008f60085afa91508115806105e957508051600114155b1561060757604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b610624611aab565b6040516060905f51602061206e5f395f51905f529061064f908735906020808a013591869101611da3565b604051602081830303815290604052805190602001205f1c6106719190611df8565b8252604080515f918782377f1e70ab38629f5091b9e2f7d72e931df3a285af6b2dff8d602f9b7ec0b8f5af1060408201527f2e800f8970e7a46ee0132beecdeb739edfff8e373ecbf4068a1b7ae7d17b34fc60608201527f2968f4ca2db982db1b5a16762a88f675300f6e47be3a3f19bc4bbbdc42a08f0c60808201527f2014567dc6005885e09299cd19edf0606146c4f459fc321e1064de1a279877e260a082015260408660c08301377f2b12e8ed1596ee4c70865cb4c3e419ef4858310f62188e8fdce45f1e8b1c4d0c6101008201527f1d81472f5650baced744698692994b59855e2e3f0a7bff9004b2f434c985b0886101208201527e18a8e18a55f23b496e11dd2bbe68fbd301642c0ea7765f0ae33cd135cd23556101408201527f03b8fa16a50162cd5334d3b4f1eb38f49814d21722ce6774efb7c0f2b449bc636101608201526020816101808360085afa9051169050806107e5576040516351d49ff760e11b815260040160405180910390fd5b5f5f61081a86868a60028060200260405190810160405280929190826002602002808284375f92019190915250610e06915050565b915091506040516101008a82377f2167539e87fb7b1884dcae8b54954bdb7ec02d36b1ec8ccefb063151b246e0456101008201527f08f359e93f4b8ad59f17b0851f63f63f58909ad4f668a49701c66435149d59fb6101208201527f15229d2bee0319b7788f85517363b4825d05080c4f80996af5fb46f8368dcda66101408201527f09514d166b708ffb95997005e9cbce7cc620b27317609236c5893210286670806101608201527f16976d04c68c156f76cdf58700d71a261c07f1677c56c640a17948af62d24c596101808201527f233b312bf47011f1e6759e0d93161be46ac4fb5f2c0070d9e5c97c54221cd2936101a08201527f1514e031e0ef5e8cbe9aa455877dc5f039deb39446e26a917bb09ffb53455a636101c08201527f2d3788e0f54f19aa6602a0fe3eb620460276da8bd625d7087e0ec9e27ad08c006101e08201527f1eab641f675c11fc79264f5fc92216816438f6613a274983ee3ae7c0c01e7d116102008201527f1484dd9cb4945159d105f0a02b2e0b5029d757f1012dd9b52fcb38837319aec161022082015282610240820152816102608201527f0ff2c0fe8c8f9bae6970ddd1682a3d7203dbc2c2fa4da33527957380e54977fb6102808201527f2b06c1da47fb4ad6bca422847059d08e40b7a0355b400a9a35e3abb2f9e3e9386102a08201527f2c4d7a7b5b4c3ca5ce1738a505f6228cb6cf7aa54740387812fe13816f691e3d6102c08201527f2a793b258a478fb8be8a5a79e5c1fcae60abd2ca4d13af8550045beaa4974bfe6102e08201526020816103008360085afa905116925082610a8557604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b610a98611b06565b610aa0611aab565b5f610ab486358760015b60200201356113e3565b8352610ad26060870135604088013560a089013560808a01356114d0565b60208501526040840152610aec60c0870135876007610aaa565b6060840152610afe8535866001610aaa565b8252610b0d8435856001610aaa565b905093509350939050565b5f5f5f610b2587876117a9565b91945092509050306360e58346848484610b3f8a8a6117da565b6040518563ffffffff1660e01b8152600401610b5e9493929190611e3f565b5f6040518083038186803b158015610b74575f5ffd5b505afa158015610b86573d5f5f3e3d5ffd5b5050505050505050505050565b5f5f825f03610ba657505f928392509050565b600183811c9250808416145f51602061204e5f395f51905f528310610bde57604051631ff3747d60e21b815260040160405180910390fd5b610c185f51602061204e5f395f51905f5260035f51602061204e5f395f51905f52865f51602061204e5f395f51905f5288890909086117f5565b91508015610c2c57610c2982611857565b91505b50915091565b5f80808085158015610c42575084155b15610c5757505f925082915081905080610dfd565b600286811c945085935060018088161490808816145f51602061204e5f395f51905f5286101580610c9557505f51602061204e5f395f51905f528510155b15610cb357604051631ff3747d60e21b815260040160405180910390fd5b5f5f51602061204e5f395f51905f52610cda60035f51602061204e5f395f51905f52611eb9565b5f51602061204e5f395f51905f52888a090990505f5f51602061204e5f395f51905f52885f51602061204e5f395f51905f528a8b090990505f5f51602061204e5f395f51905f52885f51602061204e5f395f51905f528a8b090990505f51602061204e5f395f51905f52805f51602061204e5f395f51905f528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089650610dc85f51602061204e5f395f51905f52805f51602061204e5f395f51905f528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611857565b9550610dd587878661186f565b90975095508415610df757610de987611857565b9650610df486611857565b95505b50505050505b92959194509250565b5f5f5f60019050604051604081015f7f1193e9399bd62b9865dd899657f4d0aaa2397148a76eb82db3474a8dccf3dac583527f0ef6e9d9c46319e186e39a6ad78e276bf414510fe2b5e117abdde766bb8e0f4c6020840152865182526020870151602083015260408360808560065afa841693507f15a4f907e7f99758aae0a80eb3df876cd888d7a865ab00f244370d725a3e723d82527f238a7b5bdabebcf328c8ae09902a5affbc714ad647bf0864195858a56350cf736020830152883590508060408301525f51602061206e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f251d7f8421335096c1fc24a302aa46c09cc59e0b9f91744c60e67d08b310d79e82527f07396642ad920b6be1f111549de73cc4b2f65cfc89082551f328c29ee3405ee06020830152602089013590508060408301525f51602061206e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f1e653ce7e35a3ee852e6bc174cf0ecf838115d7e374080d5e9b7d1627c6b621382527f2cfc03e7eb3f920d72386c11dd6ab408842fa6716affab6970b014341525bdf46020830152604089013590508060408301525f51602061206e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f10e38ab995dbabf79eff2e63b53f0a13ebc9ff386eaa120ef977fd1495f68f9682527f04f6bdf5021e244500f42e3b07a902499358729e659915eb48bb32d27d1b6ed26020830152606089013590508060408301525f51602061206e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f10fab557ee393db60ce06b5f92a62210a6b03464e90118a2716ea0614adb943182527f1d1c3e4abc06d2755ec72d8e6ff2fd22171373cf86fe13aa5d7d2fb5328613e96020830152608089013590508060408301525f51602061206e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f039562ccee96a4fea9a381069d5697b1caa14223e681a6dc58f40c61d704fb3d82527f2bc1b13d7f48f5be5a25f3545e2870b49b4c8d6179a30e48faa8c4a2ebd6f91c602083015260a089013590508060408301525f51602061206e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f2549bcda100d37d522410bdd4f94d6bf8c8bfc69cd999c4f1bbcc0eb6a4c813f82527f21a5bd87328b7003de2a72905f61f97a5e534de1218f7912bda6841f9763f906602083015260c089013590508060408301525f51602061206e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f1b9ebbe9712f64282aa886f48db9642ef093e1ef541aba654041a5b87c17951082527f1a1f98ac3fd8306de20675321b649eb285bfd2bdfbde15ef012cf3faf6da4965602083015260e089013590508060408301525f51602061206e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f2a92e7c3bd7ff9c2014caf87cc6e05b079b43ef2c1fb087ff10f5732b54cae7782527f2efa57ccbd11eca5b9cfb7be13774a5939770f327b28c016179dea5b07941ab5602083015261010089013590508060408301525f51602061206e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa7f2315bda2bbee86b12bad6f041152c42e7e373916a0aa05b2895724c4a18babea83527f2d6ec887548913bbe55d2663774ceefcc550a1225a5ece7fe75b763677971abf6020840152885160408085018290525f51602061206e5f395f51905f5290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806113da5760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b5f5f51602061204e5f395f51905f528310158061140d57505f51602061204e5f395f51905f528210155b1561142b57604051631ff3747d60e21b815260040160405180910390fd5b82158015611437575081155b1561144357505f6114ca565b5f61147e5f51602061204e5f395f51905f5260035f51602061204e5f395f51905f52875f51602061204e5f395f51905f52898a0909086117f5565b9050808303611493575050600182901b6114ca565b61149c81611857565b83036114af575050600182811b176114ca565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b5f5f5f51602061204e5f395f51905f52861015806114fb57505f51602061204e5f395f51905f528510155b8061151357505f51602061204e5f395f51905f528410155b8061152b57505f51602061204e5f395f51905f528310155b1561154957604051631ff3747d60e21b815260040160405180910390fd5b828486881717175f0361156057505f9050806117a0565b5f80805f51602061204e5f395f51905f5261158960035f51602061204e5f395f51905f52611eb9565b5f51602061204e5f395f51905f528a8c090990505f5f51602061204e5f395f51905f528a5f51602061204e5f395f51905f528c8d090990505f5f51602061204e5f395f51905f528a5f51602061204e5f395f51905f528c8d090990505f51602061204e5f395f51905f52805f51602061204e5f395f51905f528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50894506116775f51602061204e5f395f51905f52805f51602061204e5f395f51905f528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611857565b93505050505f5f6116c45f51602061204e5f395f51905f528061169c5761169c611de4565b5f51602061204e5f395f51905f528586095f51602061204e5f395f51905f52878809086117f5565b905061170f5f51602061204e5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f51602061204e5f395f51905f52848808096119a2565b1591505061171e83838361186f565b9093509150868314801561173157508186145b156117595780611741575f611744565b60025b60ff1660028a901b175f17945087935061179c565b61176283611857565b87148015611777575061177482611857565b86145b156114af5780611787575f61178a565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b6117b1611b24565b6117b9611ac9565b6117c1611ac9565b6117cd84860186611f67565b9250925092509250925092565b6117e2611b43565b6117ee82840184611fe6565b9392505050565b5f611820827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526119ea565b9050815f51602061204e5f395f51905f528283091461185257604051631ff3747d60e21b815260040160405180910390fd5b919050565b5f51602061204e5f395f51905f529081900681030690565b5f808061189e5f51602061204e5f395f51905f52808788095f51602061204e5f395f51905f52898a09086117f5565b905083156118b2576118af81611857565b90505b6118fb5f51602061204e5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f51602061204e5f395f51905f52848a08096117f5565b92505f51602061204e5f395f51905f526119255f51602061204e5f395f51905f5260028609611a4d565b860991505f51602061204e5f395f51905f526119505f51602061204e5f395f51905f52848509611857565b5f51602061204e5f395f51905f52858609088614158061198457505f51602061204e5f395f51905f52808385096002098514155b156113da57604051631ff3747d60e21b815260040160405180910390fd5b5f5f6119ce837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526119ea565b9050825f51602061204e5f395f51905f52828309149392505050565b5f5f60405160208152602080820152602060408201528460608201528360808201525f51602061204e5f395f51905f5260a082015260208160c08360055afa905192509050806114c857604051631ff3747d60e21b815260040160405180910390fd5b5f611a78827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd456119ea565b90505f51602061204e5f395f51905f5281830960011461185257604051631ff3747d60e21b815260040160405180910390fd5b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b8061012081018310156114ca575f5ffd5b5f5f5f5f6101e08587031215611b87575f5ffd5b6080850186811115611b97575f5ffd5b85945060a0860187811115611baa575f5ffd5b909350359150611bbd8660c08701611b62565b905092959194509250565b8061010081018310156114ca575f5ffd5b80604081018310156114ca575f5ffd5b5f5f5f5f6102a08587031215611bfd575f5ffd5b611c078686611bc8565b9350611c17866101008701611bd9565b9250611c27866101408701611bd9565b9150611bbd866101808701611b62565b5f5f5f6101808486031215611c4a575f5ffd5b611c548585611bc8565b9250611c64856101008601611bd9565b9150611c74856101408601611bd9565b90509250925092565b60c0810181855f5b6004811015611ca4578151835260209283019290910190600101611c85565b50505060808201845f5b6001811015611ccd578151835260209283019290910190600101611cae565b5050508260a0830152949350505050565b5f5f83601f840112611cee575f5ffd5b50813567ffffffffffffffff811115611d05575f5ffd5b602083019150836020828501011115611d1c575f5ffd5b9250929050565b5f5f5f5f60408587031215611d36575f5ffd5b843567ffffffffffffffff811115611d4c575f5ffd5b611d5887828801611cde565b909550935050602085013567ffffffffffffffff811115611d77575f5ffd5b611d8387828801611cde565b95989497509550505050565b634e487b7160e01b5f52603260045260245ffd5b8381528260208201525f604082018351602085015f5b82811015611dd7578151845260209384019390910190600101611db9565b5091979650505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f82611e1257634e487b7160e01b5f52601260045260245ffd5b500690565b805f5b6002811015611e39578151845260209384019390910190600101611e1a565b50505050565b6102a0810181865f5b6008811015611e67578151835260209283019290910190600101611e48565b505050611e78610100830186611e17565b611e86610140830185611e17565b6101808201835f5b6009811015611ead578151835260209283019290910190600101611e8e565b50505095945050505050565b818103818111156114ca57634e487b7160e01b5f52601160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611f0d57634e487b7160e01b5f52604160045260245ffd5b604052919050565b5f82601f830112611f24575f5ffd5b5f611f2f6040611ed8565b9050806040840185811115611f42575f5ffd5b845b81811015611f5c578035835260209283019201611f44565b509195945050505050565b5f5f5f6101808486031215611f7a575f5ffd5b5f85601f860112611f89575f5ffd5b505f80610100611f9881611ed8565b9250829150860187811115611fab575f5ffd5b865b81811015611fc5578035845260209384019301611fad565b50819550611fd38882611f15565b9450505050611c74856101408601611f15565b5f6101208284031215611ff7575f5ffd5b5f83601f840112612006575f5ffd5b505f8061012061201581611ed8565b9250829150840185811115612028575f5ffd5b845b8181101561204257803584526020938401930161202a565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212201b6f8c15232182557c5ac089626ee141a899830889e04dd1860baec7f50b72d564736f6c634300081c0033",
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
