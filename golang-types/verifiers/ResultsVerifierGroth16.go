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
	Bin: "0x6080604052348015600f57600080fd5b5061216f8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063233ace111461005c5780635d26278e1461008f57806360e58346146100a4578063b1c3a00e146100b7578063b8e72af6146100d9575b600080fd5b6040517f03e08181622212463d232b19acc793cacdc6e621c7ac24ec2cd55743a08f163681526020015b60405180910390f35b6100a261009d366004611bee565b6100ec565b005b6100a26100b2366004611c6b565b61062a565b6100ca6100c5366004611cbc565b610aa3565b60405161008693929190611d05565b6100a26100e7366004611db1565b610b2c565b6100f4611b25565b6100fc611b43565b610104611b61565b6101158660005b6020020135610bad565b6020840152825260008061012887610bad565b855160208088015160405194965092945060609360008051602061211a8339815191529361015b93929091869101611e38565b6040516020818303038152906040528051906020012060001c61017e9190611e91565b865284518452602085810151858201527f2c022fac242aba025eff55de38aab89085160c3a1c61d2f7b96ad93eb12893eb6040808701919091527f06fb294f7d92a37e12cc7c5460e990ff7b8bd4dbaee7d0b14515665cf091f7f660608701527f27ca60a3924c9991badb98107ed2d2b477d6692ed0a76306e72b4711d19110af60808701527e06c139f180152da3e6494f90fd275ef1799cccad519b76d7330c85720f0f4e60a087015260c0860185905260e086018490527f048308ed5e482ad7c8e0d4aeb6fa5467958ed3c0552278d4bc254c0c706d98f16101008701527f176100a40dd054106aa3b43f1a86dd14bc0c3abfa3b6abdd3f656c23afb7e4236101208701527f1affb35e77861d88c79098469b8b6d9ca898a0c1a04561abadfa720d93f0a0896101408701527f0b90ac9e80640caa08bf911d52216839b67e4494f968c7b05842263d851c182661016087015251600091816101808860085afa905116905080610303576040516351d49ff760e11b815260040160405180910390fd5b505050506000806103208960006004811061010b5761010b611e22565b9092509050600080808061033c60408e013560208f0135610c53565b929650909450925090506000806103548f600361010b565b915091506000806103668e8e8e610e39565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f298abf67816594b8cda829f40efb28be3eaab70f9d1cc4287d3e977f846f441a6101008e01527f2fc7db374811a4fe02a6646e136daa50edcea537c409998a6a2ba6796a8a52a66101208e01527f0fd61f9e1b6c763699da89445cac8bae6fcf48ea5ebe79b3286d2cf17c03a3b16101408e01527f0b81146dbec9183ac71ae4b13ee2c82259582c18d0f0dfc9ed0a0671a84c8a306101608e01527f222406ba9ccb6c23fc5558c0139d381db3831b8432834177cd2d8f30d26d21926101808e01527f201e4eab04c8daa253cb76270e9729ef79701d3a524ae847b26fbdf6f80d38716101a08e01527f1881db17fdd9c012842ac96467ebfb43d4a48cf553ac2d79c269246a0766fa486101c08e01527f2aba5e2d8348e812c8255b343a44ee56c8e5540c9e89bcb81a3cd7dbbfe1822c6101e08e01527f0e8d14afe89b7aa09b293af7c21c7a8223fadd28ae70dd9f58d78ea667f0960a6102008e01527f09e2801be32540e89e244692169f30e3e4d56feb454d8127031cb440eab91bb36102208e01526102408d018290526102608d018190527f21f1561d2c93080d8123184c1324bbfac61c7745cc91675d6ad8211b42cf22e96102808e01527f1abe61aee2871d26af42bec18adc48bc2e60013c1c4769669e67289b4e44dc6c6102a08e01527f2de9ed8320e7054880782942c7c73e6e60597bf2e2cceabb276efd7bad44d6e26102c08e01527f0aa88e1ed87584c8fcad5451b717911cb87d9df6f7e0c63106a6de05748c897f6102e08e0152909250905060006105db611b25565b6020816103008f60085afa91508115806105f757508051600114155b1561061557604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b610632611b25565b60405160609060008051602061211a8339815191529061065e908735906020808a013591869101611e38565b6040516020818303038152906040528051906020012060001c6106819190611e91565b8252604080516000918782377f2c022fac242aba025eff55de38aab89085160c3a1c61d2f7b96ad93eb12893eb60408201527f06fb294f7d92a37e12cc7c5460e990ff7b8bd4dbaee7d0b14515665cf091f7f660608201527f27ca60a3924c9991badb98107ed2d2b477d6692ed0a76306e72b4711d19110af60808201527e06c139f180152da3e6494f90fd275ef1799cccad519b76d7330c85720f0f4e60a082015260408660c08301377f048308ed5e482ad7c8e0d4aeb6fa5467958ed3c0552278d4bc254c0c706d98f16101008201527f176100a40dd054106aa3b43f1a86dd14bc0c3abfa3b6abdd3f656c23afb7e4236101208201527f1affb35e77861d88c79098469b8b6d9ca898a0c1a04561abadfa720d93f0a0896101408201527f0b90ac9e80640caa08bf911d52216839b67e4494f968c7b05842263d851c18266101608201526020816101808360085afa9051169050806107f6576040516351d49ff760e11b815260040160405180910390fd5b60008061082d86868a6002806020026040519081016040528092919082600260200280828437600092019190915250610e39915050565b915091506040516101008a82377f298abf67816594b8cda829f40efb28be3eaab70f9d1cc4287d3e977f846f441a6101008201527f2fc7db374811a4fe02a6646e136daa50edcea537c409998a6a2ba6796a8a52a66101208201527f0fd61f9e1b6c763699da89445cac8bae6fcf48ea5ebe79b3286d2cf17c03a3b16101408201527f0b81146dbec9183ac71ae4b13ee2c82259582c18d0f0dfc9ed0a0671a84c8a306101608201527f222406ba9ccb6c23fc5558c0139d381db3831b8432834177cd2d8f30d26d21926101808201527f201e4eab04c8daa253cb76270e9729ef79701d3a524ae847b26fbdf6f80d38716101a08201527f1881db17fdd9c012842ac96467ebfb43d4a48cf553ac2d79c269246a0766fa486101c08201527f2aba5e2d8348e812c8255b343a44ee56c8e5540c9e89bcb81a3cd7dbbfe1822c6101e08201527f0e8d14afe89b7aa09b293af7c21c7a8223fadd28ae70dd9f58d78ea667f0960a6102008201527f09e2801be32540e89e244692169f30e3e4d56feb454d8127031cb440eab91bb361022082015282610240820152816102608201527f21f1561d2c93080d8123184c1324bbfac61c7745cc91675d6ad8211b42cf22e96102808201527f1abe61aee2871d26af42bec18adc48bc2e60013c1c4769669e67289b4e44dc6c6102a08201527f2de9ed8320e7054880782942c7c73e6e60597bf2e2cceabb276efd7bad44d6e26102c08201527f0aa88e1ed87584c8fcad5451b717911cb87d9df6f7e0c63106a6de05748c897f6102e08201526020816103008360085afa905116925082610a9857604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b610aab611b80565b610ab3611b25565b6000610ac886358760015b6020020135611423565b8352610ae66060870135604088013560a089013560808a0135611518565b60208501526040840152610b0060c0870135876007610abe565b6060840152610b128535866001610abe565b8252610b218435856001610abe565b905093509350939050565b6000806000610b3b878761180f565b91945092509050306360e58346848484610b558a8a611840565b6040518563ffffffff1660e01b8152600401610b749493929190611edc565b60006040518083038186803b158015610b8c57600080fd5b505afa158015610ba0573d6000803e3d6000fd5b5050505050505050505050565b60008082600003610bc357506000928392509050565b600183811c9250808416146000805160206120fa8339815191528310610bfc57604051631ff3747d60e21b815260040160405180910390fd5b610c396000805160206120fa83398151915260036000805160206120fa833981519152866000805160206120fa833981519152888909090861185b565b91508015610c4d57610c4a826118bf565b91505b50915091565b600080808085158015610c64575084155b15610c7a57506000925082915081905080610e30565b600286811c945085935060018088161490808816146000805160206120fa83398151915286101580610cba57506000805160206120fa8339815191528510155b15610cd857604051631ff3747d60e21b815260040160405180910390fd5b60006000805160206120fa833981519152610d0260036000805160206120fa833981519152611f58565b6000805160206120fa833981519152888a0909905060006000805160206120fa833981519152886000805160206120fa8339815191528a8b0909905060006000805160206120fa833981519152886000805160206120fa8339815191528a8b090990506000805160206120fa833981519152806000805160206120fa8339815191528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089650610dfb6000805160206120fa833981519152806000805160206120fa8339815191528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775086118bf565b9550610e088787866118d8565b90975095508415610e2a57610e1c876118bf565b9650610e27866118bf565b95505b50505050505b92959194509250565b6000806000600190506040516040810160007f290a9b4db9d7bc523a665a6da45a1c1304a534c7bffefde14416ab2a0c882eaa83527f107381100460435384037229e1502d588aace41779daef1c6dc54a5c977948366020840152865182526020870151602083015260408360808560065afa841693507f1bbb05b7ef4cec589613f6688fe38d74e6749a3e99c8cf1c0d81781c8e0bcbce82527f26f82a40da2c1a65183a219a3d77a28ac25b03a9d8c89d3e08e3d5b32f03a1df60208301528835905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f1f98b68718d1cb545d6d6bacf3aa1d8506c988e2cddbac7ab580e795df35e4b382527f28786c5521d38f6bd903589e8904983fd448e3e8e7dbc2d48b9d4c0dd87f696360208301526020890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f0a31c92dcb049654c5206a452ddf6938537d87a2242ec038b45d9714fe79691d82527f16a4c6fbd3604bc3e8344a7162e439497b4360faeb513821828efb04b83971a360208301526040890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f2b8684f5cd64a4df2329b8db50b65699f30d35763a358c04fafe80850e60608182527f06e7a99f54243ab3ad4ca1ff9d2b64c8e78f62ae5b2240b12fbfe709e5c8bc0760208301526060890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f230e05f6ea8045ed54163240e406cfb2f5e1178abef93c7014557fb230784e8182527f13bea0cb75b2e655676306be43a0763d94f2a02a872293d894940cf4e2aa391360208301526080890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f18fee19fbd9afdcdf00248ddf784a1f69ae20b04a067b55f870aa143d6d09d5682527f2142eaa31078014bcc97ab7850275bc414cdbaf92c1a83ccadb60515187eef9e602083015260a0890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f1e8fc5f0fdbf176ad7a22e3866245e98a7b6c10b534c2b07c6dd83932f7b9c1282527f035b71485d24124abc180399e0cdccffd710439c91ccfca833b29bda55b31415602083015260c0890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f2b160837af7e88625f7a0fcf029ab6e8b90cea7797a7f31cbecdcfd9598a4aab82527f134a2dee742035d7810d1922071e87e36d0335fef532a546fc2292716f2aa064602083015260e0890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f032d81c7656919ed8168227903604daf8c956fae3e66adf7160ec00132e207ea82527f2b3a9c5b15e936872c8fc2a68a9a44c7cab2b38cde1a2c966ffd9526d1cad9676020830152610100890135905080604083015260008051602061211a83398151915281108416935060408260608460075afa8416935060408360808560065afa7f0e132e1f634325febddc213e5d4ec875b426f8abc529bf5d5c39a5daa023139d83527f13f1fab0c76c08720af583480b012cd89eda832fec7daf8d15e742fc08c1bdc860208401528851604080850182905260008051602061211a83398151915290911091909516169390508160608160075afa831692505060408160808360065afa8151602090920151919450909250168061141a5760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b60006000805160206120fa8339815191528310158061145057506000805160206120fa8339815191528210155b1561146e57604051631ff3747d60e21b815260040160405180910390fd5b8215801561147a575081155b1561148757506000611512565b60006114c66000805160206120fa83398151915260036000805160206120fa833981519152876000805160206120fa833981519152898a09090861185b565b90508083036114db575050600182901b611512565b6114e4816118bf565b83036114f7575050600182811b17611512565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b6000806000805160206120fa8339815191528610158061154657506000805160206120fa8339815191528510155b8061155f57506000805160206120fa8339815191528410155b8061157857506000805160206120fa8339815191528310155b1561159657604051631ff3747d60e21b815260040160405180910390fd5b828486881717176000036115af57506000905080611806565b600080806000805160206120fa8339815191526115db60036000805160206120fa833981519152611f58565b6000805160206120fa8339815191528a8c0909905060006000805160206120fa8339815191528a6000805160206120fa8339815191528c8d0909905060006000805160206120fa8339815191528a6000805160206120fa8339815191528c8d090990506000805160206120fa833981519152806000805160206120fa8339815191528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50894506116d46000805160206120fa833981519152806000805160206120fa8339815191528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775086118bf565b93505050506000806117256000805160206120fa833981519152806116fb576116fb611e7b565b6000805160206120fa8339815191528586096000805160206120fa8339815191528788090861185b565b90506117726000805160206120fa8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea46000805160206120fa83398151915284880809611a16565b159150506117818383836118d8565b9093509150868314801561179457508186145b156117be57806117a55760006117a8565b60025b60ff1660028a901b176000179450879350611802565b6117c7836118bf565b871480156117dc57506117d9826118bf565b86145b156114f757806117ed5760006117f0565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b611817611b9e565b61181f611b43565b611827611b43565b6118338486018661200e565b9250925092509250925092565b611848611bbd565b61185482840184612090565b9392505050565b6000611887827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611a60565b9050816000805160206120fa833981519152828309146118ba57604051631ff3747d60e21b815260040160405180910390fd5b919050565b6000805160206120fa8339815191529081900681030690565b6000808061190a6000805160206120fa833981519152808788096000805160206120fa833981519152898a090861185b565b9050831561191e5761191b816118bf565b90505b6119696000805160206120fa8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea46000805160206120fa833981519152848a080961185b565b92506000805160206120fa8339815191526119956000805160206120fa83398151915260028609611ac5565b860991506000805160206120fa8339815191526119c26000805160206120fa8339815191528485096118bf565b6000805160206120fa83398151915285860908861415806119f857506000805160206120fa833981519152808385096002098514155b1561141a57604051631ff3747d60e21b815260040160405180910390fd5b600080611a43837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611a60565b9050826000805160206120fa833981519152828309149392505050565b60008060405160208152602080820152602060408201528460608201528360808201526000805160206120fa83398151915260a082015260208160c08360055afa9051925090508061151057604051631ff3747d60e21b815260040160405180910390fd5b6000611af1827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611a60565b90506000805160206120fa8339815191528183096001146118ba57604051631ff3747d60e21b815260040160405180910390fd5b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b80610120810183101561151257600080fd5b6000806000806101e08587031215611c0557600080fd5b6080850186811115611c1657600080fd5b85945060a0860187811115611c2a57600080fd5b909350359150611c3d8660c08701611bdc565b905092959194509250565b80610100810183101561151257600080fd5b806040810183101561151257600080fd5b6000806000806102a08587031215611c8257600080fd5b611c8c8686611c48565b9350611c9c866101008701611c5a565b9250611cac866101408701611c5a565b9150611c3d866101808701611bdc565b60008060006101808486031215611cd257600080fd5b611cdc8585611c48565b9250611cec856101008601611c5a565b9150611cfc856101408601611c5a565b90509250925092565b60c08101818560005b6004811015611d2d578151835260209283019290910190600101611d0e565b505050608082018460005b6001811015611d57578151835260209283019290910190600101611d38565b5050508260a0830152949350505050565b60008083601f840112611d7a57600080fd5b50813567ffffffffffffffff811115611d9257600080fd5b602083019150836020828501011115611daa57600080fd5b9250929050565b60008060008060408587031215611dc757600080fd5b843567ffffffffffffffff811115611dde57600080fd5b611dea87828801611d68565b909550935050602085013567ffffffffffffffff811115611e0a57600080fd5b611e1687828801611d68565b95989497509550505050565b634e487b7160e01b600052603260045260246000fd5b83815282602082015260006040820183516020850160005b82811015611e6e578151845260209384019390910190600101611e50565b5091979650505050505050565b634e487b7160e01b600052601260045260246000fd5b600082611eae57634e487b7160e01b600052601260045260246000fd5b500690565b8060005b6002811015611ed6578151845260209384019390910190600101611eb7565b50505050565b6102a08101818660005b6008811015611f05578151835260209283019290910190600101611ee6565b505050611f16610100830186611eb3565b611f24610140830185611eb3565b61018082018360005b6009811015611f4c578151835260209283019290910190600101611f2d565b50505095945050505050565b8181038181111561151257634e487b7160e01b600052601160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611fb057634e487b7160e01b600052604160045260246000fd5b604052919050565b600082601f830112611fc957600080fd5b6000611fd56040611f79565b9050806040840185811115611fe957600080fd5b845b81811015612003578035835260209283019201611feb565b509195945050505050565b6000806000610180848603121561202457600080fd5b600085601f860112612034578081fd5b8061010061204181611f79565b925082915086018781111561205557600080fd5b865b8181101561206f578035845260209384019301612057565b5081955061207d8882611fb8565b9450505050611cfc856101408601611fb8565b600061012082840312156120a357600080fd5b600083601f8401126120b3578081fd5b806101206120c081611f79565b92508291508401858111156120d457600080fd5b845b818110156120ee5780358452602093840193016120d6565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212202615e20772919c7dfbebed55c5ee0a6691cbcac6db700a47a21d73d1291f08d264736f6c634300081c0033",
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
