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
	Bin: "0x6080604052348015600f57600080fd5b5061216f8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063233ace111461005c5780635d26278e1461008f57806360e58346146100a4578063b1c3a00e146100b7578063b8e72af6146100d9575b600080fd5b6040517f0e9d0a473c7dbf7cb257035b1a98321975d8c22039180b6d9d279225eb2c6db481526020015b60405180910390f35b6100a261009d366004611bee565b6100ec565b005b6100a26100b2366004611c6b565b61062a565b6100ca6100c5366004611cbc565b610aa3565b60405161008693929190611d05565b6100a26100e7366004611db1565b610b2c565b6100f4611b25565b6100fc611b43565b610104611b61565b6101158660005b6020020135610bad565b6020840152825260008061012887610bad565b855160208088015160405194965092945060609360008051602061211a8339815191529361015b93929091869101611e38565b6040516020818303038152906040528051906020012060001c61017e9190611e91565b865284518452602085810151858201527f06a741fbbce7642e3011b12d629a0c9aee9ec5779b798aca0eb454a57cd9dc1d6040808701919091527f19061c7ccf9acb61a022b2696f76260e8ee056746ceff60290fa09f74614bd9460608701527f19c020bda171cd66ac6f2670bd667133bade3dfc1edd2857700e98cd188e2e8060808701527f13e7151a5e24eb283c7715b0d350488e3364a3c72484b28cdc9f3b7d54b64ed260a087015260c0860185905260e086018490527f091fc30d5bf24e283f4801f3da49113c036375ebbedc447237d395d5f0acc9446101008701527f22fb4247ab46cb853fd8fee4379fe8d44efc166611293ead6b122b976ac4d95e6101208701527f1e9ea7ed2cea231e05777ac43c3f9c46590dbb8184f3c25fbcd2fae89f2b49216101408701527f2d43fff77ca6952fcd99ab70ddb5bf5a247485e271ece03a26c9ac9b1e7a58be61016087015251600091816101808860085afa905116905080610304576040516351d49ff760e11b815260040160405180910390fd5b505050506000806103218960006004811061010b5761010b611e22565b9092509050600080808061033d60408e013560208f0135610c53565b929650909450925090506000806103558f600361010b565b915091506000806103678e8e8e610e39565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f0aa57ebd2053ce01b39ffb365672232fba50025d0ad15219563f40b37750bb0f6101008e01527f2956ec974ff824c922f13d1622188ed21ec9b31ba1845b6d6613c7f74ca672c96101208e01527f2c277f360d00681aa5f8282ca2285c5bf9c0e0d49de71bad468f8ca8baa597e46101408e01527f0a56e530454353718e60248e2f78a672ab1aa2104d185fce798f5824877a8fd86101608e01527f23e5265c67cf1f1541526aa457d8c75823534eac1ea67d58751d4ed371b195766101808e01527f2fb9465406e3859fcd0234cfcc41ceffd5a312f684077c8e92ccb5987d86537f6101a08e01527f157a2741b166f1d0758764447778826a2bc2e50c88c271865b0008ba785e3c1a6101c08e01527f1ce6d16697d0db238ab3bb91b05b90085331579970049f6519a40e16b086c9706101e08e01527e68fcdeed6e958c715deaa07b03d09db7a8a0b275e26dd52cf8cb033f911e0a6102008e01527f1da148119e5c0a8844aa0d7a986649916f4e57e3b0bd50d0ef8aaa8f04faf50a6102208e01526102408d018290526102608d018190527f2922093ffc53214999d05144efa247a8ece844190d0904f32f5e083f60f150086102808e01527f02cac33f664cb478bb875cae49334a6f30e9e83a9215fbbd488ee89fe80d74b56102a08e01527f301efada081ac8d5047dce0c57f8d0c21f60ebf7a1141faf8b83068a1c4eb76d6102c08e01527f253cf9e0d817c0d5439a57cde9dbd856bd4264e42e36795a0f16a582418284ee6102e08e0152909250905060006105db611b25565b6020816103008f60085afa91508115806105f757508051600114155b1561061557604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b610632611b25565b60405160609060008051602061211a8339815191529061065e908735906020808a013591869101611e38565b6040516020818303038152906040528051906020012060001c6106819190611e91565b8252604080516000918782377f06a741fbbce7642e3011b12d629a0c9aee9ec5779b798aca0eb454a57cd9dc1d60408201527f19061c7ccf9acb61a022b2696f76260e8ee056746ceff60290fa09f74614bd9460608201527f19c020bda171cd66ac6f2670bd667133bade3dfc1edd2857700e98cd188e2e8060808201527f13e7151a5e24eb283c7715b0d350488e3364a3c72484b28cdc9f3b7d54b64ed260a082015260408660c08301377f091fc30d5bf24e283f4801f3da49113c036375ebbedc447237d395d5f0acc9446101008201527f22fb4247ab46cb853fd8fee4379fe8d44efc166611293ead6b122b976ac4d95e6101208201527f1e9ea7ed2cea231e05777ac43c3f9c46590dbb8184f3c25fbcd2fae89f2b49216101408201527f2d43fff77ca6952fcd99ab70ddb5bf5a247485e271ece03a26c9ac9b1e7a58be6101608201526020816101808360085afa9051169050806107f7576040516351d49ff760e11b815260040160405180910390fd5b60008061082e86868a6002806020026040519081016040528092919082600260200280828437600092019190915250610e39915050565b915091506040516101008a82377f0aa57ebd2053ce01b39ffb365672232fba50025d0ad15219563f40b37750bb0f6101008201527f2956ec974ff824c922f13d1622188ed21ec9b31ba1845b6d6613c7f74ca672c96101208201527f2c277f360d00681aa5f8282ca2285c5bf9c0e0d49de71bad468f8ca8baa597e46101408201527f0a56e530454353718e60248e2f78a672ab1aa2104d185fce798f5824877a8fd86101608201527f23e5265c67cf1f1541526aa457d8c75823534eac1ea67d58751d4ed371b195766101808201527f2fb9465406e3859fcd0234cfcc41ceffd5a312f684077c8e92ccb5987d86537f6101a08201527f157a2741b166f1d0758764447778826a2bc2e50c88c271865b0008ba785e3c1a6101c08201527f1ce6d16697d0db238ab3bb91b05b90085331579970049f6519a40e16b086c9706101e08201527e68fcdeed6e958c715deaa07b03d09db7a8a0b275e26dd52cf8cb033f911e0a6102008201527f1da148119e5c0a8844aa0d7a986649916f4e57e3b0bd50d0ef8aaa8f04faf50a61022082015282610240820152816102608201527f2922093ffc53214999d05144efa247a8ece844190d0904f32f5e083f60f150086102808201527f02cac33f664cb478bb875cae49334a6f30e9e83a9215fbbd488ee89fe80d74b56102a08201527f301efada081ac8d5047dce0c57f8d0c21f60ebf7a1141faf8b83068a1c4eb76d6102c08201527f253cf9e0d817c0d5439a57cde9dbd856bd4264e42e36795a0f16a582418284ee6102e08201526020816103008360085afa905116925082610a9857604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b610aab611b80565b610ab3611b25565b6000610ac886358760015b6020020135611423565b8352610ae66060870135604088013560a089013560808a0135611518565b60208501526040840152610b0060c0870135876007610abe565b6060840152610b128535866001610abe565b8252610b218435856001610abe565b905093509350939050565b6000806000610b3b878761180f565b91945092509050306360e58346848484610b558a8a611840565b6040518563ffffffff1660e01b8152600401610b749493929190611edc565b60006040518083038186803b158015610b8c57600080fd5b505afa158015610ba0573d6000803e3d6000fd5b5050505050505050505050565b60008082600003610bc357506000928392509050565b600183811c9250808416146000805160206120fa8339815191528310610bfc57604051631ff3747d60e21b815260040160405180910390fd5b610c396000805160206120fa83398151915260036000805160206120fa833981519152866000805160206120fa833981519152888909090861185b565b91508015610c4d57610c4a826118bf565b91505b50915091565b600080808085158015610c64575084155b15610c7a57506000925082915081905080610e30565b600286811c945085935060018088161490808816146000805160206120fa83398151915286101580610cba57506000805160206120fa8339815191528510155b15610cd857604051631ff3747d60e21b815260040160405180910390fd5b60006000805160206120fa833981519152610d0260036000805160206120fa833981519152611f58565b6000805160206120fa833981519152888a0909905060006000805160206120fa833981519152886000805160206120fa8339815191528a8b0909905060006000805160206120fa833981519152886000805160206120fa8339815191528a8b090990506000805160206120fa833981519152806000805160206120fa8339815191528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089650610dfb6000805160206120fa833981519152806000805160206120fa8339815191528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775086118bf565b9550610e088787866118d8565b90975095508415610e2a57610e1c876118bf565b9650610e27866118bf565b95505b50505050505b92959194509250565b6000806000600190506040516040810160007f0ea085ee5ee1358f209a2595d6d14ed5e383709c2005d0469e17fb56ab232fe283527f042c5367b963cef1442df24412fed764f6867494c98213dc6a9249a6510519c26020840152865182526020870151602083015260408360808560065afa841693507f12007da0643c9c66a960c82f2b355bbf7478aa51d571e1b3bcccdfe67362654982527f261f74ba74998db7c077ad205db0fb19d0a48b94829d65e45548f324125cb6b760208301528835905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f10077912c131f04d989481c8436d4f8e378590982716c3c7efc50b35ce8b928082527f1937ecbcd357b02696da5c757b52ceb8ede29a42aff7d18d76e4b2b9000d16c760208301526020890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f2effb31312709bbd46232d6533f62b36b3cb5e22fd5d1a99e0404005df2bfd7a82527f034cda548880a3db412427b25832d79ee4a449bc8e2257ce8179c3f18d9c128360208301526040890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f22417a017a1a2288512926384021ac97be806a3606e5eae067e7d79210abb70582527f0280204e879a8628b240387957d7ce870d7eb4a49732d54234560192122529ab60208301526060890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f1d93016b6c387f3093152b48ef99133999e95b2e5eec63fe2c5dbdee8c1937e982527f29508496aca19e1ae9de7b3e31093208bbb740287ccfe870a870688919fa846560208301526080890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f1f100508def2645e56bf86396666c9d25add60572d34c8506639bf379e80008682527f126abd22a92522d256e3576d343e7358be8db4cbb841b6749ac1b3efdcf1ba33602083015260a0890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f0302e12d4c6d41475496ca65a3b5e32ffd32d86aae20e95035811807122ddda382527f07be4757dbe04b4827df7e423778e82c0d8c05ddb1420e300d8e090585588275602083015260c0890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f289d767a6cc91d6a5ded86fde40e309ed98259e35c13e9567f190c19d9ddaa5d82527f050e944e6cd17276aeb16d7fa38fbcd4276e4eff599b6d9affc87545fd27371b602083015260e0890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f057511123a37e76bc20d9f67a8e3b0126c181079c209439a3211c4660722f11782527f300b086e1164b3856d9621d715d858a55bbfe9da3cb99ef8ee17584aa9f05a6b6020830152610100890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa7f1902623881ab243a84d0a9bf149d32c90d0053f9d27f0c369b1b9715a15a012b83527f0dc8384eafd9fb424148d60c1e3b6008b6a81ca121d4eb6c8f53b35e93adf3eb60208401528851604080850182905260008051602061211a83398151915290911091909516169390508160608160075afa831692505060408160808360065afa8151602090920151919450909250168061141a5760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b60006000805160206120fa8339815191528310158061145057506000805160206120fa8339815191528210155b1561146e57604051631ff3747d60e21b815260040160405180910390fd5b8215801561147a575081155b1561148757506000611512565b60006114c66000805160206120fa83398151915260036000805160206120fa833981519152876000805160206120fa833981519152898a09090861185b565b90508083036114db575050600182901b611512565b6114e4816118bf565b83036114f7575050600182811b17611512565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b6000806000805160206120fa8339815191528610158061154657506000805160206120fa8339815191528510155b8061155f57506000805160206120fa8339815191528410155b8061157857506000805160206120fa8339815191528310155b1561159657604051631ff3747d60e21b815260040160405180910390fd5b828486881717176000036115af57506000905080611806565b600080806000805160206120fa8339815191526115db60036000805160206120fa833981519152611f58565b6000805160206120fa8339815191528a8c0909905060006000805160206120fa8339815191528a6000805160206120fa8339815191528c8d0909905060006000805160206120fa8339815191528a6000805160206120fa8339815191528c8d090990506000805160206120fa833981519152806000805160206120fa8339815191528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50894506116d46000805160206120fa833981519152806000805160206120fa8339815191528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775086118bf565b93505050506000806117256000805160206120fa833981519152806116fb576116fb611e7b565b6000805160206120fa8339815191528586096000805160206120fa8339815191528788090861185b565b90506117726000805160206120fa8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea46000805160206120fa83398151915284880809611a16565b159150506117818383836118d8565b9093509150868314801561179457508186145b156117be57806117a55760006117a8565b60025b60ff1660028a901b176000179450879350611802565b6117c7836118bf565b871480156117dc57506117d9826118bf565b86145b156114f757806117ed5760006117f0565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b611817611b9e565b61181f611b43565b611827611b43565b6118338486018661200e565b9250925092509250925092565b611848611bbd565b61185482840184612090565b9392505050565b6000611887827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611a60565b9050816000805160206120fa833981519152828309146118ba57604051631ff3747d60e21b815260040160405180910390fd5b919050565b6000805160206120fa8339815191529081900681030690565b6000808061190a6000805160206120fa833981519152808788096000805160206120fa833981519152898a090861185b565b9050831561191e5761191b816118bf565b90505b6119696000805160206120fa8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea46000805160206120fa833981519152848a080961185b565b92506000805160206120fa8339815191526119956000805160206120fa83398151915260028609611ac5565b860991506000805160206120fa8339815191526119c26000805160206120fa8339815191528485096118bf565b6000805160206120fa83398151915285860908861415806119f857506000805160206120fa833981519152808385096002098514155b1561141a57604051631ff3747d60e21b815260040160405180910390fd5b600080611a43837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611a60565b9050826000805160206120fa833981519152828309149392505050565b60008060405160208152602080820152602060408201528460608201528360808201526000805160206120fa83398151915260a082015260208160c08360055afa9051925090508061151057604051631ff3747d60e21b815260040160405180910390fd5b6000611af1827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611a60565b90506000805160206120fa8339815191528183096001146118ba57604051631ff3747d60e21b815260040160405180910390fd5b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b80610120810183101561151257600080fd5b6000806000806101e08587031215611c0557600080fd5b6080850186811115611c1657600080fd5b85945060a0860187811115611c2a57600080fd5b909350359150611c3d8660c08701611bdc565b905092959194509250565b80610100810183101561151257600080fd5b806040810183101561151257600080fd5b6000806000806102a08587031215611c8257600080fd5b611c8c8686611c48565b9350611c9c866101008701611c5a565b9250611cac866101408701611c5a565b9150611c3d866101808701611bdc565b60008060006101808486031215611cd257600080fd5b611cdc8585611c48565b9250611cec856101008601611c5a565b9150611cfc856101408601611c5a565b90509250925092565b60c08101818560005b6004811015611d2d578151835260209283019290910190600101611d0e565b505050608082018460005b6001811015611d57578151835260209283019290910190600101611d38565b5050508260a0830152949350505050565b60008083601f840112611d7a57600080fd5b50813567ffffffffffffffff811115611d9257600080fd5b602083019150836020828501011115611daa57600080fd5b9250929050565b60008060008060408587031215611dc757600080fd5b843567ffffffffffffffff811115611dde57600080fd5b611dea87828801611d68565b909550935050602085013567ffffffffffffffff811115611e0a57600080fd5b611e1687828801611d68565b95989497509550505050565b634e487b7160e01b600052603260045260246000fd5b83815282602082015260006040820183516020850160005b82811015611e6e578151845260209384019390910190600101611e50565b5091979650505050505050565b634e487b7160e01b600052601260045260246000fd5b600082611eae57634e487b7160e01b600052601260045260246000fd5b500690565b8060005b6002811015611ed6578151845260209384019390910190600101611eb7565b50505050565b6102a08101818660005b6008811015611f05578151835260209283019290910190600101611ee6565b505050611f16610100830186611eb3565b611f24610140830185611eb3565b61018082018360005b6009811015611f4c578151835260209283019290910190600101611f2d565b50505095945050505050565b8181038181111561151257634e487b7160e01b600052601160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611fb057634e487b7160e01b600052604160045260246000fd5b604052919050565b600082601f830112611fc957600080fd5b6000611fd56040611f79565b9050806040840185811115611fe957600080fd5b845b81811015612003578035835260209283019201611feb565b509195945050505050565b6000806000610180848603121561202457600080fd5b600085601f860112612034578081fd5b8061010061204181611f79565b925082915086018781111561205557600080fd5b865b8181101561206f578035845260209384019301612057565b5081955061207d8882611fb8565b9450505050611cfc856101408601611fb8565b600061012082840312156120a357600080fd5b600083601f8401126120b3578081fd5b806101206120c081611f79565b92508291508401858111156120d457600080fd5b845b818110156120ee5780358452602093840193016120d6565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212206c7accc2b43fcc11de303a789671025d3e956b7c78b9876f66c10ea07741ce6664736f6c634300081c0033",
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
