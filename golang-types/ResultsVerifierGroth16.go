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
	ABI: "[{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50611b268061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063233ace111461005c57806344f636921461008f5780635f89feef146100af5780638a3ae438146100c4578063b8e72af6146100d7575b600080fd5b6040517f189351fed43a79f7db8aa1bc8177f60e8d3886850402d97e48a55067404a29bd81526020015b60405180910390f35b6100a261009d366004611705565b6100ea565b6040516100869190611722565b6100c26100bd366004611765565b610147565b005b6100c26100d23660046117a1565b610466565b6100c26100e5366004611821565b6106e6565b6100f261163c565b61010582358360015b6020020135610764565b81526101236060830135604084013560a08501356080860135610859565b6020830152604082015261013d60c08301358360076100fb565b6060820152919050565b61014f61165a565b60008061016285825b6020020135610b50565b9092509050600080808061017e60408a013560208b0135610bf6565b929650909450925090506000806101968b6003610158565b915091506000806101a68c610ddc565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f09d6a74c774261deff5f52a5ce365f9adb36b083c049d21cf39247d675e731eb6101008e01527f2ddb8b0f89bcdabeb4ebeaa105a8efb330e4cec7cfc5283013cc48270d4ab0bd6101208e01527f2cec4efc684f9a93f7bb903fe3cfd444a01db116ff37b6d625fa81a2fe48615a6101408e01527f0ebb338461fc69245c3ccf93fd004d6fb2e3bbe8cf3da8a758d494b34a908a2a6101608e01527f2fba622298414ec2a021acc932d8db5fdc05229bf0cb31d7e8f8cf1f40ad06256101808e01527f069dee094ea916a101ec6af0b54a89019d44daef82601ae24de30b344cf2f14e6101a08e01527f18ff2046eaf0dc5e75c4e66299ec6e511fff55d37c02dcb0730e3f71f80d6b306101c08e01527f1bc0bed39160dd1a4f9c929bfaf8178247f08137d4e0c3d0274efc0442b66ade6101e08e01527f20c0b7806f4824e3bf2b67004c132de842f12e830259f90a2700742ddfd852d16102008e01527f19c8a7a1c77ea30ba9abb8db7faaad0c95380c7980961d5c0c7e5f6f93708d1d6102208e01526102408d018290526102608d018190527f12b30a062a192396b7b4e2762590840d572cc653c6114cda5b0d6d2bf967c9286102808e01527f019260ee38c64fd9502d7b2e11fdac1a6807ee38f1e7c2df7ba17be21989cfc76102a08e01527f0a0e65a91b921235722f3987242c55de6a81b406ccc2541d3a731c950f0973ae6102c08e01527f0841ded6f41559e9b52ad8c1f600ee1e9870908b3ad7c690eba88d245d51029e6102e08e01529092509050600061041b611679565b6020816103008f60085afa915081158061043757508051600114155b1561045557604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050505050505050565b60008061047283610ddc565b9150915060006040516101008682377f09d6a74c774261deff5f52a5ce365f9adb36b083c049d21cf39247d675e731eb6101008201527f2ddb8b0f89bcdabeb4ebeaa105a8efb330e4cec7cfc5283013cc48270d4ab0bd6101208201527f2cec4efc684f9a93f7bb903fe3cfd444a01db116ff37b6d625fa81a2fe48615a6101408201527f0ebb338461fc69245c3ccf93fd004d6fb2e3bbe8cf3da8a758d494b34a908a2a6101608201527f2fba622298414ec2a021acc932d8db5fdc05229bf0cb31d7e8f8cf1f40ad06256101808201527f069dee094ea916a101ec6af0b54a89019d44daef82601ae24de30b344cf2f14e6101a08201527f18ff2046eaf0dc5e75c4e66299ec6e511fff55d37c02dcb0730e3f71f80d6b306101c08201527f1bc0bed39160dd1a4f9c929bfaf8178247f08137d4e0c3d0274efc0442b66ade6101e08201527f20c0b7806f4824e3bf2b67004c132de842f12e830259f90a2700742ddfd852d16102008201527f19c8a7a1c77ea30ba9abb8db7faaad0c95380c7980961d5c0c7e5f6f93708d1d61022082015283610240820152826102608201527f12b30a062a192396b7b4e2762590840d572cc653c6114cda5b0d6d2bf967c9286102808201527f019260ee38c64fd9502d7b2e11fdac1a6807ee38f1e7c2df7ba17be21989cfc76102a08201527f0a0e65a91b921235722f3987242c55de6a81b406ccc2541d3a731c950f0973ae6102c08201527f0841ded6f41559e9b52ad8c1f600ee1e9870908b3ad7c690eba88d245d51029e6102e08201526020816103008360085afa9051169050806106df57604051631ff3747d60e21b815260040160405180910390fd5b5050505050565b60006106f2858561131d565b50509050306001600160a01b0316638a3ae43882610710868661134e565b6040518363ffffffff1660e01b815260040161072d929190611892565b60006040518083038186803b15801561074557600080fd5b505afa158015610759573d6000803e3d6000fd5b505050505050505050565b6000600080516020611ab1833981519152831015806107915750600080516020611ab18339815191528210155b156107af57604051631ff3747d60e21b815260040160405180910390fd5b821580156107bb575081155b156107c857506000610853565b6000610807600080516020611ab18339815191526003600080516020611ab183398151915287600080516020611ab1833981519152898a090908611369565b905080830361081c575050600182901b610853565b610825816113cd565b8303610838575050600182811b17610853565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b600080600080516020611ab1833981519152861015806108875750600080516020611ab18339815191528510155b806108a05750600080516020611ab18339815191528410155b806108b95750600080516020611ab18339815191528310155b156108d757604051631ff3747d60e21b815260040160405180910390fd5b828486881717176000036108f057506000905080610b47565b60008080600080516020611ab183398151915261091c6003600080516020611ab1833981519152611906565b600080516020611ab18339815191528a8c090990506000600080516020611ab18339815191528a600080516020611ab18339815191528c8d090990506000600080516020611ab18339815191528a600080516020611ab18339815191528c8d09099050600080516020611ab183398151915280600080516020611ab18339815191528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610a15600080516020611ab183398151915280600080516020611ab18339815191528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775086113cd565b9350505050600080610a66600080516020611ab183398151915280610a3c57610a3c6118f0565b600080516020611ab1833981519152858609600080516020611ab183398151915287880908611369565b9050610ab3600080516020611ab18339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611ab1833981519152848808096113e6565b15915050610ac2838383611430565b90935091508683148015610ad557508186145b15610aff5780610ae6576000610ae9565b60025b60ff1660028a901b176000179450879350610b43565b610b08836113cd565b87148015610b1d5750610b1a826113cd565b86145b156108385780610b2e576000610b31565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b60008082600003610b6657506000928392509050565b600183811c925080841614600080516020611ab18339815191528310610b9f57604051631ff3747d60e21b815260040160405180910390fd5b610bdc600080516020611ab18339815191526003600080516020611ab183398151915286600080516020611ab18339815191528889090908611369565b91508015610bf057610bed826113cd565b91505b50915091565b600080808085158015610c07575084155b15610c1d57506000925082915081905080610dd3565b600286811c94508593506001808816149080881614600080516020611ab183398151915286101580610c5d5750600080516020611ab18339815191528510155b15610c7b57604051631ff3747d60e21b815260040160405180910390fd5b6000600080516020611ab1833981519152610ca56003600080516020611ab1833981519152611906565b600080516020611ab1833981519152888a090990506000600080516020611ab183398151915288600080516020611ab18339815191528a8b090990506000600080516020611ab183398151915288600080516020611ab18339815191528a8b09099050600080516020611ab183398151915280600080516020611ab18339815191528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089650610d9e600080516020611ab183398151915280600080516020611ab18339815191528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775086113cd565b9550610dab878786611430565b90975095508415610dcd57610dbf876113cd565b9650610dca866113cd565b95505b50505050505b92959194509250565b6000806000600190506040516040810160007f1f242b1654e3e3e80ef44321c54786bff0b936b389109cd8eae76d8145675be183527f117f221e951d7bca0d8509cfcb63f047d10da8dcb3c55dbb5f80636e7b586c1960208401527f293b1dd16e64287259f2fb646aee5a48d69adc0964ea1dc391905af7b36c89ca82527f131247bddcd1e1190c950e4372ca2fac0c9408553bcf58ed6ec5f76bc6006229602083015286359050806040830152600080516020611ad183398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f0273d65546c6e30428a263f8223491b278e2c83815c14b5e1814d24bfeabfa0782527f25a81aa6775148553ae62c19cf54de712b195a5ab3a6e285122e9de0b5bd81b4602083015260208701359050806040830152600080516020611ad183398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f2b93cf43be853f4cfef5546f178213c328418becce71b97c6efcfc9ea1862eed82527f1cad005be5db66915ed934d5e61ef3c03fc951f71a9fb15b097f0190e234d508602083015260408701359050806040830152600080516020611ad183398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f1d5890d30a4daebd963a5de8c9a8d5508088b664fbf3eeb5bd8fee5fac0a9af682527f220aa23957720d3c786bf8aeb6119a91e212371ebcc9f49ae3b269b4ee891ccf602083015260608701359050806040830152600080516020611ad183398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f0ea67e89f51052acfee3c8d08713859936cf207af909fc078d8916158693f95a82527f1a2f960800d4925c2b97abdbdcb9d1c8d10ca523c44c1055bce4abfc5d966c1d602083015260808701359050806040830152600080516020611ad183398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f22b1de0907a58bf0bdd6b1250ce23408364667df410c6571ebbb0fb621f5517282527f2508ac7601b972b6f6e6bded3dff76d7f0dd2be2f6b574e94e5e780a805eff08602083015260a08701359050806040830152600080516020611ad183398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f25aa76862733859c9e019b52a7801ba029da276f6e736464b7de72a492f4189782527f260f4775e6947446fad731b270344a56c8c1ebe9d28e699678965a371d494906602083015260c08701359050806040830152600080516020611ad183398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f101a5187f71f30529c65f0e11d652337a2829461e8d8eda9952d9617cb20fc8f82527f0524269174503b69bf1eb5c157cce5b58688e0a57de5e726e86757d7339fcb63602083015260e08701359050806040830152600080516020611ad183398151915281108416935060408260608460075afa8416935060408360808560065afa7f10b6bc43ef92ca3d9d490ebfdbc3e77e23b794ff05c49817d17e533dd039d74483527f0cc70306c6ce9c3ec3e1e6166a735659bc215593594158f44b5892d075c409fc60208401526101008801356040808501829052600080516020611ad183398151915290911091909516169390508160608160075afa831692505060408160808360065afa81516020909201519194509092501680610bf05760405163a54f8e2760e01b815260040160405180910390fd5b611325611697565b61132d6116b6565b6113356116b6565b611341848601866119bc565b9250925092509250925092565b6113566116d4565b61136282840184611a47565b9392505050565b6000611395827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611577565b905081600080516020611ab1833981519152828309146113c857604051631ff3747d60e21b815260040160405180910390fd5b919050565b600080516020611ab18339815191529081900681030690565b600080611413837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611577565b905082600080516020611ab1833981519152828309149392505050565b60008080611462600080516020611ab183398151915280878809600080516020611ab1833981519152898a0908611369565b9050831561147657611473816113cd565b90505b6114c1600080516020611ab18339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611ab1833981519152848a0809611369565b9250600080516020611ab18339815191526114ed600080516020611ab1833981519152600286096115dc565b86099150600080516020611ab183398151915261151a600080516020611ab18339815191528485096113cd565b600080516020611ab183398151915285860908861415806115505750600080516020611ab1833981519152808385096002098514155b1561156e57604051631ff3747d60e21b815260040160405180910390fd5b50935093915050565b6000806040516020815260208082015260206040820152846060820152836080820152600080516020611ab183398151915260a082015260208160c08360055afa9051925090508061085157604051631ff3747d60e21b815260040160405180910390fd5b6000611608827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611577565b9050600080516020611ab18339815191528183096001146113c857604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b80610100810183101561085357600080fd5b6000610100828403121561171857600080fd5b61136283836116f3565b60808101818360005b600481101561174a57815183526020928301929091019060010161172b565b50505092915050565b80610120810183101561085357600080fd5b6000806101a0838503121561177957600080fd5b608083018481111561178a57600080fd5b8392506117978582611753565b9150509250929050565b60008061022083850312156117b557600080fd5b6117bf84846116f3565b91506117cf846101008501611753565b90509250929050565b60008083601f8401126117ea57600080fd5b50813567ffffffffffffffff81111561180257600080fd5b60208301915083602082850101111561181a57600080fd5b9250929050565b6000806000806040858703121561183757600080fd5b843567ffffffffffffffff81111561184e57600080fd5b61185a878288016117d8565b909550935050602085013567ffffffffffffffff81111561187a57600080fd5b611886878288016117d8565b95989497509550505050565b6102208101818460005b60088110156118bb57815183526020928301929091019060010161189c565b50505061010082018360005b60098110156118e65781518352602092830192909101906001016118c7565b5050509392505050565b634e487b7160e01b600052601260045260246000fd5b8181038181111561085357634e487b7160e01b600052601160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561195e57634e487b7160e01b600052604160045260246000fd5b604052919050565b600082601f83011261197757600080fd5b60006119836040611927565b905080604084018581111561199757600080fd5b845b818110156119b1578035835260209283019201611999565b509195945050505050565b600080600061018084860312156119d257600080fd5b600085601f8601126119e2578081fd5b806101006119ef81611927565b9250829150860187811115611a0357600080fd5b865b81811015611a1d578035845260209384019301611a05565b50819550611a2b8882611966565b9450505050611a3e856101408601611966565b90509250925092565b60006101208284031215611a5a57600080fd5b600083601f840112611a6a578081fd5b80610120611a7781611927565b9250829150840185811115611a8b57600080fd5b845b81811015611aa5578035845260209384019301611a8d565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a264697066735822122067603d17fff7d24c9ee9badde00bf0d1abd544e524ca67c26dcb993d11db8b2364736f6c634300081c0033",
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

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _ResultsVerifierGroth16.Contract.CompressProof(&_ResultsVerifierGroth16.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) VerifyCompressedProof(compressedProof [4]*big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyCompressedProof(&_ResultsVerifierGroth16.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5f89feef.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyCompressedProof(&_ResultsVerifierGroth16.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8a3ae438.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x8a3ae438.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) VerifyProof(proof [8]*big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof(&_ResultsVerifierGroth16.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x8a3ae438.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) VerifyProof(proof [8]*big.Int, input [9]*big.Int) error {
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
