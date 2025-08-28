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
	Bin: "0x6080604052348015600f57600080fd5b506121718061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063233ace111461005c5780635d26278e1461008f57806360e58346146100a4578063b1c3a00e146100b7578063b8e72af6146100d9575b600080fd5b6040517f8e647e94d2814554a3a240fca4a32df8a9e3a2714e06ad1f41410af6e29bc42f81526020015b60405180910390f35b6100a261009d366004611bf0565b6100ec565b005b6100a26100b2366004611c6d565b61062b565b6100ca6100c5366004611cbe565b610aa5565b60405161008693929190611d07565b6100a26100e7366004611db3565b610b2e565b6100f4611b27565b6100fc611b45565b610104611b63565b6101158660005b6020020135610baf565b6020840152825260008061012887610baf565b855160208088015160405194965092945060609360008051602061211c8339815191529361015b93929091869101611e3a565b6040516020818303038152906040528051906020012060001c61017e9190611e93565b865284518452602085810151858201527f1e1c5fb595b688d8f325a564a0167da9dc2f6a973b56b730229a13ddc1463ea26040808701919091527f115380cccc293ae7ea076338004ecea8dac4c5d8f7f41a535909e6c21193b46c60608701527f19be1e09ea61acfb1cf0a5ec997a346baff5f14305f1965fdc5dcd0c50dc683b60808701527f13ac4a4abaa3a1325fe099afcac34028d93eb5001d16c0ad247f088e740f73b260a087015260c0860185905260e086018490527f2c62363b7adc1209b370844b009a350355b5bd29039a2c67337dbc5c558ab95e6101008701527f083bd74b4f2da45d48dda9dc6f97cf5b78d43dceeb1d282a17e8a90a764f11766101208701527f1bd9fb4ee02a4b5a5182e9e39bb8e0c1531cbde65f2445661addc9532a0e52ab6101408701527f229443f5b802ee3d8c881e940790de97a4e0b0b38f1807a4ad656ecd1500ed0761016087015251600091816101808860085afa905116905080610304576040516351d49ff760e11b815260040160405180910390fd5b505050506000806103218960006004811061010b5761010b611e24565b9092509050600080808061033d60408e013560208f0135610c55565b929650909450925090506000806103558f600361010b565b915091506000806103678e8e8e610e3b565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f0243e06ad0608fb131cd8077e28d76ccdb9b720a2b6b9bf4558d635cae58ea716101008e01527f04be9de2a61b3685451fb986fa243562916e538a159ec93cb5c3b22461b24e7e6101208e01527f2bf73a1e9297ead977a1e5e933d3205e8e0d3122a55f5769a291904d6bcb71416101408e01527f03daa11852d5e73f07ce1c2963cff495f57006bcada7fb32576f4b523ffbe3c06101608e01527f19929209b7fdb77dc66e05be7c5963231a8b72f84ba421695936a7e2f34dac296101808e01527f18dc1c5abd4263f8b3360f5a5637131db8c30eea14241db05fe6534ba8b4ad046101a08e01527f26efccb2cae8670da27578eaaf73d8a1c0e821a966cf119540d10021c5e3b2dc6101c08e01527f1ec42cdeb7662441d0cc66ffc1e6e51e12daca18f40480dc8e14eb9b452c885a6101e08e01527f203fba5fe7ea43ee444b0d94773a4831ae9b664442d6821cbda35321273bafc66102008e01527f081f604ab8477dba95dfce2cc6ece1307a50a5cdf84439d02d1b0e461d5a6ca16102208e01526102408d018290526102608d018190527f27e5bacad5b6a86c362af10f1f71869039477f78a6aedf07ca9efcb4c41d48856102808e01527f234b5f78a112318940f36eb98d525fc552a63c1241ec7dec7a260c0bff6565146102a08e01527f0fbd874a42d587fd0381193abe8cae674aa021befee6eed4e25aab93c99d65b96102c08e01527f18a35805879d4a486700044b97f485f9c9a86272836b9c64336ca8074f7794f76102e08e0152909250905060006105dc611b27565b6020816103008f60085afa91508115806105f857508051600114155b1561061657604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b610633611b27565b60405160609060008051602061211c8339815191529061065f908735906020808a013591869101611e3a565b6040516020818303038152906040528051906020012060001c6106829190611e93565b8252604080516000918782377f1e1c5fb595b688d8f325a564a0167da9dc2f6a973b56b730229a13ddc1463ea260408201527f115380cccc293ae7ea076338004ecea8dac4c5d8f7f41a535909e6c21193b46c60608201527f19be1e09ea61acfb1cf0a5ec997a346baff5f14305f1965fdc5dcd0c50dc683b60808201527f13ac4a4abaa3a1325fe099afcac34028d93eb5001d16c0ad247f088e740f73b260a082015260408660c08301377f2c62363b7adc1209b370844b009a350355b5bd29039a2c67337dbc5c558ab95e6101008201527f083bd74b4f2da45d48dda9dc6f97cf5b78d43dceeb1d282a17e8a90a764f11766101208201527f1bd9fb4ee02a4b5a5182e9e39bb8e0c1531cbde65f2445661addc9532a0e52ab6101408201527f229443f5b802ee3d8c881e940790de97a4e0b0b38f1807a4ad656ecd1500ed076101608201526020816101808360085afa9051169050806107f8576040516351d49ff760e11b815260040160405180910390fd5b60008061082f86868a6002806020026040519081016040528092919082600260200280828437600092019190915250610e3b915050565b915091506040516101008a82377f0243e06ad0608fb131cd8077e28d76ccdb9b720a2b6b9bf4558d635cae58ea716101008201527f04be9de2a61b3685451fb986fa243562916e538a159ec93cb5c3b22461b24e7e6101208201527f2bf73a1e9297ead977a1e5e933d3205e8e0d3122a55f5769a291904d6bcb71416101408201527f03daa11852d5e73f07ce1c2963cff495f57006bcada7fb32576f4b523ffbe3c06101608201527f19929209b7fdb77dc66e05be7c5963231a8b72f84ba421695936a7e2f34dac296101808201527f18dc1c5abd4263f8b3360f5a5637131db8c30eea14241db05fe6534ba8b4ad046101a08201527f26efccb2cae8670da27578eaaf73d8a1c0e821a966cf119540d10021c5e3b2dc6101c08201527f1ec42cdeb7662441d0cc66ffc1e6e51e12daca18f40480dc8e14eb9b452c885a6101e08201527f203fba5fe7ea43ee444b0d94773a4831ae9b664442d6821cbda35321273bafc66102008201527f081f604ab8477dba95dfce2cc6ece1307a50a5cdf84439d02d1b0e461d5a6ca161022082015282610240820152816102608201527f27e5bacad5b6a86c362af10f1f71869039477f78a6aedf07ca9efcb4c41d48856102808201527f234b5f78a112318940f36eb98d525fc552a63c1241ec7dec7a260c0bff6565146102a08201527f0fbd874a42d587fd0381193abe8cae674aa021befee6eed4e25aab93c99d65b96102c08201527f18a35805879d4a486700044b97f485f9c9a86272836b9c64336ca8074f7794f76102e08201526020816103008360085afa905116925082610a9a57604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b610aad611b82565b610ab5611b27565b6000610aca86358760015b6020020135611425565b8352610ae86060870135604088013560a089013560808a013561151a565b60208501526040840152610b0260c0870135876007610ac0565b6060840152610b148535866001610ac0565b8252610b238435856001610ac0565b905093509350939050565b6000806000610b3d8787611811565b91945092509050306360e58346848484610b578a8a611842565b6040518563ffffffff1660e01b8152600401610b769493929190611ede565b60006040518083038186803b158015610b8e57600080fd5b505afa158015610ba2573d6000803e3d6000fd5b5050505050505050505050565b60008082600003610bc557506000928392509050565b600183811c9250808416146000805160206120fc8339815191528310610bfe57604051631ff3747d60e21b815260040160405180910390fd5b610c3b6000805160206120fc83398151915260036000805160206120fc833981519152866000805160206120fc833981519152888909090861185d565b91508015610c4f57610c4c826118c1565b91505b50915091565b600080808085158015610c66575084155b15610c7c57506000925082915081905080610e32565b600286811c945085935060018088161490808816146000805160206120fc83398151915286101580610cbc57506000805160206120fc8339815191528510155b15610cda57604051631ff3747d60e21b815260040160405180910390fd5b60006000805160206120fc833981519152610d0460036000805160206120fc833981519152611f5a565b6000805160206120fc833981519152888a0909905060006000805160206120fc833981519152886000805160206120fc8339815191528a8b0909905060006000805160206120fc833981519152886000805160206120fc8339815191528a8b090990506000805160206120fc833981519152806000805160206120fc8339815191528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089650610dfd6000805160206120fc833981519152806000805160206120fc8339815191528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775086118c1565b9550610e0a8787866118da565b90975095508415610e2c57610e1e876118c1565b9650610e29866118c1565b95505b50505050505b92959194509250565b6000806000600190506040516040810160007f279ea2a5532d7bd04b28b0f0fa211206e0a569a38eaedd3d39703285e2f414fc83527f304d08de70635dc5f9a000939699fd072aa9807b109fc2afa724e269b9bae1286020840152865182526020870151602083015260408360808560065afa841693507f09440714037f2d643a0afaa9c35497b3ffaadd1f29b5b3367c837e9b31b7ba4a82527f2fbb60ccf35b2b42b0fc683e91e43a5902ebc500923b144c2dac68853a92a12d60208301528835905080604083015260008051602061211c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f28d9a01d840999bfec2f5f88dcb296d4d6b5ede5b5dd10f89c236655cc97d12d82527f10507e0e18de2f93a8c68b59d8afe652acfb6855aae0119993af93e36593966360208301526020890135905080604083015260008051602061211c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f29587b1491de5a7b923362a9c1c9f8df90c440e58028dcccd4cea36e4aec478082527f05cc22b89b9e4100f21e06828b01cbba3de74e31d9772f2092fad0e25063151c60208301526040890135905080604083015260008051602061211c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f217131232c0d66474d58e63c4d003d6e412517611f03aa4bb9967117eff6efc782527f2e2b153fae0001a60c2a8a9435348e25f4587787aa8dbcc58c53c4a08f68467f60208301526060890135905080604083015260008051602061211c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f2cec6878bada538058198e47db87b3c6e83a0571152c1436f60b494d7b3dd5e382527f179f833adc3cb1386ce1aa43a1b5bf86e883b3438d9d365aebe619af3010172b60208301526080890135905080604083015260008051602061211c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f28af68e36135ae3acb29a81fab853828886adff9c4ec35e606ae25fd6df8cf8c82527f2f1a4fb5495efb32c9440ede57d79485aa1581ec764aa9b17139cc837d31f1b4602083015260a0890135905080604083015260008051602061211c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f0ea7cfd61b26692251ce1dca66b58a8778ef9f67bc355bd81e3d08bdd7becf9282527f0d97336b2a05e90236f8b6d18e96eb42249334383bfd04d1d44445c279a7dbb0602083015260c0890135905080604083015260008051602061211c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f2f2539ff07b0206c12fd251e686e58a85eb631ba6142e03faa4664d0d5ed51e682527f070702f766e47fe1010c1b911e28cbc3ad539081638cbef48197e7615af49b0b602083015260e0890135905080604083015260008051602061211c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f0559de020763b4e3bd11d8e825c9d3e3f569d89507596ee5c46c930477ce686e82527f0170ceaa1b60dd2b921b3442ac85b812d44bd5c71789a4921a20a095acba7d1e6020830152610100890135905080604083015260008051602061211c83398151915281108416935060408260608460075afa8416935060408360808560065afa7f260bdc3affbf1944659b70532095950305281d7b053b7e8b806e987ed9280f6883527f2f360856a493088e23c9c53f67de21752f00692a8ec35e0affdd0afa22eb0dc860208401528851604080850182905260008051602061211c83398151915290911091909516169390508160608160075afa831692505060408160808360065afa8151602090920151919450909250168061141c5760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b60006000805160206120fc8339815191528310158061145257506000805160206120fc8339815191528210155b1561147057604051631ff3747d60e21b815260040160405180910390fd5b8215801561147c575081155b1561148957506000611514565b60006114c86000805160206120fc83398151915260036000805160206120fc833981519152876000805160206120fc833981519152898a09090861185d565b90508083036114dd575050600182901b611514565b6114e6816118c1565b83036114f9575050600182811b17611514565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b6000806000805160206120fc8339815191528610158061154857506000805160206120fc8339815191528510155b8061156157506000805160206120fc8339815191528410155b8061157a57506000805160206120fc8339815191528310155b1561159857604051631ff3747d60e21b815260040160405180910390fd5b828486881717176000036115b157506000905080611808565b600080806000805160206120fc8339815191526115dd60036000805160206120fc833981519152611f5a565b6000805160206120fc8339815191528a8c0909905060006000805160206120fc8339815191528a6000805160206120fc8339815191528c8d0909905060006000805160206120fc8339815191528a6000805160206120fc8339815191528c8d090990506000805160206120fc833981519152806000805160206120fc8339815191528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50894506116d66000805160206120fc833981519152806000805160206120fc8339815191528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775086118c1565b93505050506000806117276000805160206120fc833981519152806116fd576116fd611e7d565b6000805160206120fc8339815191528586096000805160206120fc8339815191528788090861185d565b90506117746000805160206120fc8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea46000805160206120fc83398151915284880809611a18565b159150506117838383836118da565b9093509150868314801561179657508186145b156117c057806117a75760006117aa565b60025b60ff1660028a901b176000179450879350611804565b6117c9836118c1565b871480156117de57506117db826118c1565b86145b156114f957806117ef5760006117f2565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b611819611ba0565b611821611b45565b611829611b45565b61183584860186612010565b9250925092509250925092565b61184a611bbf565b61185682840184612092565b9392505050565b6000611889827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611a62565b9050816000805160206120fc833981519152828309146118bc57604051631ff3747d60e21b815260040160405180910390fd5b919050565b6000805160206120fc8339815191529081900681030690565b6000808061190c6000805160206120fc833981519152808788096000805160206120fc833981519152898a090861185d565b905083156119205761191d816118c1565b90505b61196b6000805160206120fc8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea46000805160206120fc833981519152848a080961185d565b92506000805160206120fc8339815191526119976000805160206120fc83398151915260028609611ac7565b860991506000805160206120fc8339815191526119c46000805160206120fc8339815191528485096118c1565b6000805160206120fc83398151915285860908861415806119fa57506000805160206120fc833981519152808385096002098514155b1561141c57604051631ff3747d60e21b815260040160405180910390fd5b600080611a45837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611a62565b9050826000805160206120fc833981519152828309149392505050565b60008060405160208152602080820152602060408201528460608201528360808201526000805160206120fc83398151915260a082015260208160c08360055afa9051925090508061151257604051631ff3747d60e21b815260040160405180910390fd5b6000611af3827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611a62565b90506000805160206120fc8339815191528183096001146118bc57604051631ff3747d60e21b815260040160405180910390fd5b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b80610120810183101561151457600080fd5b6000806000806101e08587031215611c0757600080fd5b6080850186811115611c1857600080fd5b85945060a0860187811115611c2c57600080fd5b909350359150611c3f8660c08701611bde565b905092959194509250565b80610100810183101561151457600080fd5b806040810183101561151457600080fd5b6000806000806102a08587031215611c8457600080fd5b611c8e8686611c4a565b9350611c9e866101008701611c5c565b9250611cae866101408701611c5c565b9150611c3f866101808701611bde565b60008060006101808486031215611cd457600080fd5b611cde8585611c4a565b9250611cee856101008601611c5c565b9150611cfe856101408601611c5c565b90509250925092565b60c08101818560005b6004811015611d2f578151835260209283019290910190600101611d10565b505050608082018460005b6001811015611d59578151835260209283019290910190600101611d3a565b5050508260a0830152949350505050565b60008083601f840112611d7c57600080fd5b50813567ffffffffffffffff811115611d9457600080fd5b602083019150836020828501011115611dac57600080fd5b9250929050565b60008060008060408587031215611dc957600080fd5b843567ffffffffffffffff811115611de057600080fd5b611dec87828801611d6a565b909550935050602085013567ffffffffffffffff811115611e0c57600080fd5b611e1887828801611d6a565b95989497509550505050565b634e487b7160e01b600052603260045260246000fd5b83815282602082015260006040820183516020850160005b82811015611e70578151845260209384019390910190600101611e52565b5091979650505050505050565b634e487b7160e01b600052601260045260246000fd5b600082611eb057634e487b7160e01b600052601260045260246000fd5b500690565b8060005b6002811015611ed8578151845260209384019390910190600101611eb9565b50505050565b6102a08101818660005b6008811015611f07578151835260209283019290910190600101611ee8565b505050611f18610100830186611eb5565b611f26610140830185611eb5565b61018082018360005b6009811015611f4e578151835260209283019290910190600101611f2f565b50505095945050505050565b8181038181111561151457634e487b7160e01b600052601160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611fb257634e487b7160e01b600052604160045260246000fd5b604052919050565b600082601f830112611fcb57600080fd5b6000611fd76040611f7b565b9050806040840185811115611feb57600080fd5b845b81811015612005578035835260209283019201611fed565b509195945050505050565b6000806000610180848603121561202657600080fd5b600085601f860112612036578081fd5b8061010061204381611f7b565b925082915086018781111561205757600080fd5b865b81811015612071578035845260209384019301612059565b5081955061207f8882611fba565b9450505050611cfe856101408601611fba565b600061012082840312156120a557600080fd5b600083601f8401126120b5578081fd5b806101206120c281611f7b565b92508291508401858111156120d657600080fd5b845b818110156120f05780358452602093840193016120d8565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212209d065047b941ff446a20218a6e56ab32d482ac1b02827a3c0d3524d16deb722364736f6c634300081c0033",
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
