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

// StateTransitionVerifierGroth16MetaData contains all meta data concerning the StateTransitionVerifierGroth16 contract.
var StateTransitionVerifierGroth16MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"input\",\"type\":\"uint256[4]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"input\",\"type\":\"uint256[4]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50611ee38061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063233ace111461005c578063b1c3a00e1461008f578063b8e72af6146100b1578063d148354c146100c6578063e33974a8146100d9575b600080fd5b6040517fdfa73b0982a5f73a021421fac39e7096f10fe0d78193d66130fa2a861099ee3181526020015b60405180910390f35b6100a261009d366004611991565b6100ec565b60405161008693929190611a03565b6100c46100bf366004611a92565b610175565b005b6100c46100d4366004611b14565b6101f6565b6100c46100e7366004611b70565b610699565b6100f46118d6565b6100fc6118f4565b600061011186358760015b6020020135610c02565b835261012f6060870135604088013560a089013560808a0135610cf7565b6020850152604084015261014960c0870135876007610107565b606084015261015b8535866001610107565b825261016a8435856001610107565b905093509350939050565b60008060006101848787610fee565b919450925090503063d148354c84848461019e8a8a61101f565b6040518563ffffffff1660e01b81526004016101bd9493929190611bf3565b60006040518083038186803b1580156101d557600080fd5b505afa1580156101e9573d6000803e3d6000fd5b5050505050505050505050565b6101fe6118f4565b6040805160028082526060808301845292602083019080368337019050509050602081016040848101823750604051600080516020611e8e83398151915290610253908735906020808a013591869101611c52565b6040516020818303038152906040528051906020012060001c6102769190611cab565b8252604080516000918782377f0130b9a60483e81094a515fd7ae2ce24c2553b54138b2e5bd83487d1bdb87db260408201527f226161531add140908b5c7e0970f272db8964795d987d25949a9ae3f84167b4a60608201527f216c7e0bc61fab8ba2dce8db14e3b664d9c9cf0db3e4d918e995695544a5de9060808201527f1ef5893104a98624a5f7c1f0b2bcc1495cb44f71f99e45036b64367c223d874160a082015260408660c08301377f170ccf986f57703ae803d2ddf1195ea00891d95e92a06720519a1985417b9b736101008201527f103c9252d54b100696b99d74594b4a934e1a79c8422f03e4b4b12faad84fe59c6101208201527f181cbb600f66514e3f3f762afac31b1d158c414cd82cd7a7045465285210b4426101408201527f19c98a2c04ba100f2e9283263f1a7c169a78bccfd447f48caa1941bd42d7c9dc6101608201526020816101808360085afa9051169050806103ec576040516351d49ff760e11b815260040160405180910390fd5b60008061042386868a600280602002604051908101604052809291908260026020028082843760009201919091525061103a915050565b915091506040516101008a82377f1f7a57446dca8c1e4cb2c635c9edd6f82135ec4ea5f4740c698daae96eefe6e36101008201527f141f6a8e0972175e7ff667de9191091f684080f68327be23a86341ecd95908f86101208201527f0bc96f4b68f9645dca1774bf2201a2aa0e65adee331b5425a799cc8ebbfd810b6101408201527f2dd50714b41b0efe6f635f481b472668fe0c7754d0c4c5ada42e3c4ccd2f35b36101608201527f04384061a929248d0627c0cb79dc28af4fb9ec5848d2d19a85869ecd39f77b936101808201527f152970895883e8130a6823d3103637b60a5949def9d849811474e414f54b423a6101a08201527f2cade224ef780b5711d57378615b114ebb3d45202ab12c323deb8174163c7ddc6101c08201527f1efe2d8238d9fd0f340b6d666954e98bb3aae31b7b14cf97f59d3821bed07e466101e08201527f061e0fc1a642a67a598a8d7272d281bc98bc6257cc90c1813ae1845e970d520c6102008201527f1ef4814ab5620ee121d93e6cbfe30be7eb56f1d5fa086222838ac09d86824d8461022082015282610240820152816102608201527f24cefd73b9957d01ea01235010fb9cbd5eb040e85e2e43a25d456a800cdfe4976102808201527f2bff99f34f8ba0b14099d22d5c36827e568621195f1276a7fb1a3ae98ea5597b6102a08201527f1da36783d688ec57cd9dc7dfd6a6b20e6054b0e3a3bdb5218449e5c7915c624c6102c08201527f124115b995eecf55ec99fadd8bddfbfbe98cc02cf4e0d37fbfd98225f8fc59d26102e08201526020816103008360085afa90511692508261068e57604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b6106a16118f4565b6106a9611912565b6106b1611930565b6106c28660005b6020020135611380565b602084015282526000806106d587611380565b6040805160028082526060808301845294965092945091906020830190803683370190505090506020810160408881018237508451602080870151604051600080516020611e8e8339815191529361073293909291869101611c52565b6040516020818303038152906040528051906020012060001c6107559190611cab565b865284518452602085810151858201527f0130b9a60483e81094a515fd7ae2ce24c2553b54138b2e5bd83487d1bdb87db26040808701919091527f226161531add140908b5c7e0970f272db8964795d987d25949a9ae3f84167b4a60608701527f216c7e0bc61fab8ba2dce8db14e3b664d9c9cf0db3e4d918e995695544a5de9060808701527f1ef5893104a98624a5f7c1f0b2bcc1495cb44f71f99e45036b64367c223d874160a087015260c0860185905260e086018490527f170ccf986f57703ae803d2ddf1195ea00891d95e92a06720519a1985417b9b736101008701527f103c9252d54b100696b99d74594b4a934e1a79c8422f03e4b4b12faad84fe59c6101208701527f181cbb600f66514e3f3f762afac31b1d158c414cd82cd7a7045465285210b4426101408701527f19c98a2c04ba100f2e9283263f1a7c169a78bccfd447f48caa1941bd42d7c9dc61016087015251600091816101808860085afa9051169050806108db576040516351d49ff760e11b815260040160405180910390fd5b505050506000806108f8896000600481106106b8576106b8611bba565b9092509050600080808061091460408e013560208f0135611426565b9296509094509250905060008061092c8f60036106b8565b9150915060008061093e8e8e8e61103a565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f1f7a57446dca8c1e4cb2c635c9edd6f82135ec4ea5f4740c698daae96eefe6e36101008e01527f141f6a8e0972175e7ff667de9191091f684080f68327be23a86341ecd95908f86101208e01527f0bc96f4b68f9645dca1774bf2201a2aa0e65adee331b5425a799cc8ebbfd810b6101408e01527f2dd50714b41b0efe6f635f481b472668fe0c7754d0c4c5ada42e3c4ccd2f35b36101608e01527f04384061a929248d0627c0cb79dc28af4fb9ec5848d2d19a85869ecd39f77b936101808e01527f152970895883e8130a6823d3103637b60a5949def9d849811474e414f54b423a6101a08e01527f2cade224ef780b5711d57378615b114ebb3d45202ab12c323deb8174163c7ddc6101c08e01527f1efe2d8238d9fd0f340b6d666954e98bb3aae31b7b14cf97f59d3821bed07e466101e08e01527f061e0fc1a642a67a598a8d7272d281bc98bc6257cc90c1813ae1845e970d520c6102008e01527f1ef4814ab5620ee121d93e6cbfe30be7eb56f1d5fa086222838ac09d86824d846102208e01526102408d018290526102608d018190527f24cefd73b9957d01ea01235010fb9cbd5eb040e85e2e43a25d456a800cdfe4976102808e01527f2bff99f34f8ba0b14099d22d5c36827e568621195f1276a7fb1a3ae98ea5597b6102a08e01527f1da36783d688ec57cd9dc7dfd6a6b20e6054b0e3a3bdb5218449e5c7915c624c6102c08e01527f124115b995eecf55ec99fadd8bddfbfbe98cc02cf4e0d37fbfd98225f8fc59d26102e08e015290925090506000610bb36118f4565b6020816103008f60085afa9150811580610bcf57508051600114155b15610bed57604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b6000600080516020611e6e83398151915283101580610c2f5750600080516020611e6e8339815191528210155b15610c4d57604051631ff3747d60e21b815260040160405180910390fd5b82158015610c59575081155b15610c6657506000610cf1565b6000610ca5600080516020611e6e8339815191526003600080516020611e6e83398151915287600080516020611e6e833981519152898a09090861160c565b9050808303610cba575050600182901b610cf1565b610cc381611670565b8303610cd6575050600182811b17610cf1565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b600080600080516020611e6e83398151915286101580610d255750600080516020611e6e8339815191528510155b80610d3e5750600080516020611e6e8339815191528410155b80610d575750600080516020611e6e8339815191528310155b15610d7557604051631ff3747d60e21b815260040160405180910390fd5b82848688171717600003610d8e57506000905080610fe5565b60008080600080516020611e6e833981519152610dba6003600080516020611e6e833981519152611ccd565b600080516020611e6e8339815191528a8c090990506000600080516020611e6e8339815191528a600080516020611e6e8339815191528c8d090990506000600080516020611e6e8339815191528a600080516020611e6e8339815191528c8d09099050600080516020611e6e83398151915280600080516020611e6e8339815191528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610eb3600080516020611e6e83398151915280600080516020611e6e8339815191528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611670565b9350505050600080610f04600080516020611e6e83398151915280610eda57610eda611c95565b600080516020611e6e833981519152858609600080516020611e6e8339815191528788090861160c565b9050610f51600080516020611e6e8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611e6e83398151915284880809611689565b15915050610f608383836116d3565b90935091508683148015610f7357508186145b15610f9d5780610f84576000610f87565b60025b60ff1660028a901b176000179450879350610fe1565b610fa683611670565b87148015610fbb5750610fb882611670565b86145b15610cd65780610fcc576000610fcf565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b610ff661194f565b610ffe611912565b611006611912565b61101284860186611d83565b9250925092509250925092565b6110276118d6565b61103382840184611e05565b9392505050565b6000806000600190506040516040810160007f2d69ca0c5e57e9b45d0e13ae879ff2b3e5c300bcffd3723732e3c3afbcbf4e9383527f1c5f562d7bc1042dbd38d349559db2f77f09f050b1ccd0566664c68dee4c06586020840152865182526020870151602083015260408360808560065afa841693507f126ac41ab51d92b2868f8c6e18e71922b8998f28dff87b7528f2eec611f3d7e582527f2291bac2a098eabccc4aac785f583a7a38ecbea5924f93824a213738ec9fde82602083015288359050806040830152600080516020611e8e83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f010602b5f9e9ce25be2823fbd5007a4aa615d2d397f7b929aa891c470bf7244082527f0e46815fc9a9e0013f13461e956c9431527271023e93c2393795e062ddeca3b3602083015260208901359050806040830152600080516020611e8e83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f2f70c993fe60bfafc067514dd1f60438ec8e84a0f384f3ffa1fdce386b3d38bb82527f15f2222fc2a85fa561706d5d0c0af2c28492949908593231b95dd466c14c1553602083015260408901359050806040830152600080516020611e8e83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f0662325b763579516e7d03067df47454ca7591cd2f3026751fc91136aff95b1182527f283e2dd692a570727bde33b1c88fa985ba4cb40b1b310f7d067c9cb54e547f98602083015260608901359050806040830152600080516020611e8e83398151915281108416935060408260608460075afa8416935060408360808560065afa7f2b041e67c03cc18e5cda151e4bb87679f0ea6401a0565fbafee028da97ba438683527f1a76021f5f1f512fc54a0b499a522e70f8d0b3c486cb51fac74f9e086dc22c5c602084015288516040808501829052600080516020611e8e83398151915290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806113775760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b6000808260000361139657506000928392509050565b600183811c925080841614600080516020611e6e83398151915283106113cf57604051631ff3747d60e21b815260040160405180910390fd5b61140c600080516020611e6e8339815191526003600080516020611e6e83398151915286600080516020611e6e833981519152888909090861160c565b915080156114205761141d82611670565b91505b50915091565b600080808085158015611437575084155b1561144d57506000925082915081905080611603565b600286811c94508593506001808816149080881614600080516020611e6e8339815191528610158061148d5750600080516020611e6e8339815191528510155b156114ab57604051631ff3747d60e21b815260040160405180910390fd5b6000600080516020611e6e8339815191526114d56003600080516020611e6e833981519152611ccd565b600080516020611e6e833981519152888a090990506000600080516020611e6e83398151915288600080516020611e6e8339815191528a8b090990506000600080516020611e6e83398151915288600080516020611e6e8339815191528a8b09099050600080516020611e6e83398151915280600080516020611e6e8339815191528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50896506115ce600080516020611e6e83398151915280600080516020611e6e8339815191528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611670565b95506115db8787866116d3565b909750955084156115fd576115ef87611670565b96506115fa86611670565b95505b50505050505b92959194509250565b6000611638827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611811565b905081600080516020611e6e8339815191528283091461166b57604051631ff3747d60e21b815260040160405180910390fd5b919050565b600080516020611e6e8339815191529081900681030690565b6000806116b6837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611811565b905082600080516020611e6e833981519152828309149392505050565b60008080611705600080516020611e6e83398151915280878809600080516020611e6e833981519152898a090861160c565b905083156117195761171681611670565b90505b611764600080516020611e6e8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611e6e833981519152848a080961160c565b9250600080516020611e6e833981519152611790600080516020611e6e83398151915260028609611876565b86099150600080516020611e6e8339815191526117bd600080516020611e6e833981519152848509611670565b600080516020611e6e83398151915285860908861415806117f35750600080516020611e6e833981519152808385096002098514155b1561137757604051631ff3747d60e21b815260040160405180910390fd5b6000806040516020815260208082015260206040820152846060820152836080820152600080516020611e6e83398151915260a082015260208160c08360055afa90519250905080610cef57604051631ff3747d60e21b815260040160405180910390fd5b60006118a2827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611811565b9050600080516020611e6e83398151915281830960011461166b57604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b806101008101831015610cf157600080fd5b8060408101831015610cf157600080fd5b600080600061018084860312156119a757600080fd5b6119b1858561196e565b92506119c1856101008601611980565b91506119d1856101408601611980565b90509250925092565b8060005b60048110156119fd5781518452602093840193909101906001016119de565b50505050565b60c08101611a1182866119da565b608082018460005b6001811015611a38578151835260209283019290910190600101611a19565b5050508260a0830152949350505050565b60008083601f840112611a5b57600080fd5b50813567ffffffffffffffff811115611a7357600080fd5b602083019150836020828501011115611a8b57600080fd5b9250929050565b60008060008060408587031215611aa857600080fd5b843567ffffffffffffffff811115611abf57600080fd5b611acb87828801611a49565b909550935050602085013567ffffffffffffffff811115611aeb57600080fd5b611af787828801611a49565b95989497509550505050565b8060808101831015610cf157600080fd5b6000806000806102008587031215611b2b57600080fd5b611b35868661196e565b9350611b45866101008701611980565b9250611b55866101408701611980565b9150611b65866101808701611b03565b905092959194509250565b6000806000806101408587031215611b8757600080fd5b611b918686611b03565b935060a0850186811115611ba457600080fd5b608086019350359150611b658660c08701611b03565b634e487b7160e01b600052603260045260246000fd5b8060005b60028110156119fd578151845260209384019390910190600101611bd4565b6102008101818660005b6008811015611c1c578151835260209283019290910190600101611bfd565b505050611c2d610100830186611bd0565b611c3b610140830185611bd0565b611c496101808301846119da565b95945050505050565b83815282602082015260006040820183516020850160005b82811015611c88578151845260209384019390910190600101611c6a565b5091979650505050505050565b634e487b7160e01b600052601260045260246000fd5b600082611cc857634e487b7160e01b600052601260045260246000fd5b500690565b81810381811115610cf157634e487b7160e01b600052601160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611d2557634e487b7160e01b600052604160045260246000fd5b604052919050565b600082601f830112611d3e57600080fd5b6000611d4a6040611cee565b9050806040840185811115611d5e57600080fd5b845b81811015611d78578035835260209283019201611d60565b509195945050505050565b60008060006101808486031215611d9957600080fd5b600085601f860112611da9578081fd5b80610100611db681611cee565b9250829150860187811115611dca57600080fd5b865b81811015611de4578035845260209384019301611dcc565b50819550611df28882611d2d565b94505050506119d1856101408601611d2d565b600060808284031215611e1757600080fd5b600083601f840112611e27578081fd5b80611e326080611cee565b90508091506080840185811115611e4857600080fd5b845b81811015611e62578035845260209384019301611e4a565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212206f4f67485bbf5154c915cb2b74692ceb2b9bae6735d28528c9eb083285f1976b64736f6c634300081c0033",
}

// StateTransitionVerifierGroth16ABI is the input ABI used to generate the binding from.
// Deprecated: Use StateTransitionVerifierGroth16MetaData.ABI instead.
var StateTransitionVerifierGroth16ABI = StateTransitionVerifierGroth16MetaData.ABI

// StateTransitionVerifierGroth16Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateTransitionVerifierGroth16MetaData.Bin instead.
var StateTransitionVerifierGroth16Bin = StateTransitionVerifierGroth16MetaData.Bin

// DeployStateTransitionVerifierGroth16 deploys a new Ethereum contract, binding an instance of StateTransitionVerifierGroth16 to it.
func DeployStateTransitionVerifierGroth16(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StateTransitionVerifierGroth16, error) {
	parsed, err := StateTransitionVerifierGroth16MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateTransitionVerifierGroth16Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateTransitionVerifierGroth16{StateTransitionVerifierGroth16Caller: StateTransitionVerifierGroth16Caller{contract: contract}, StateTransitionVerifierGroth16Transactor: StateTransitionVerifierGroth16Transactor{contract: contract}, StateTransitionVerifierGroth16Filterer: StateTransitionVerifierGroth16Filterer{contract: contract}}, nil
}

// StateTransitionVerifierGroth16 is an auto generated Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16 struct {
	StateTransitionVerifierGroth16Caller     // Read-only binding to the contract
	StateTransitionVerifierGroth16Transactor // Write-only binding to the contract
	StateTransitionVerifierGroth16Filterer   // Log filterer for contract events
}

// StateTransitionVerifierGroth16Caller is an auto generated read-only Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransitionVerifierGroth16Transactor is an auto generated write-only Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransitionVerifierGroth16Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateTransitionVerifierGroth16Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransitionVerifierGroth16Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateTransitionVerifierGroth16Session struct {
	Contract     *StateTransitionVerifierGroth16 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// StateTransitionVerifierGroth16CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateTransitionVerifierGroth16CallerSession struct {
	Contract *StateTransitionVerifierGroth16Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// StateTransitionVerifierGroth16TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateTransitionVerifierGroth16TransactorSession struct {
	Contract     *StateTransitionVerifierGroth16Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// StateTransitionVerifierGroth16Raw is an auto generated low-level Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16Raw struct {
	Contract *StateTransitionVerifierGroth16 // Generic contract binding to access the raw methods on
}

// StateTransitionVerifierGroth16CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16CallerRaw struct {
	Contract *StateTransitionVerifierGroth16Caller // Generic read-only contract binding to access the raw methods on
}

// StateTransitionVerifierGroth16TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateTransitionVerifierGroth16TransactorRaw struct {
	Contract *StateTransitionVerifierGroth16Transactor // Generic write-only contract binding to access the raw methods on
}

// NewStateTransitionVerifierGroth16 creates a new instance of StateTransitionVerifierGroth16, bound to a specific deployed contract.
func NewStateTransitionVerifierGroth16(address common.Address, backend bind.ContractBackend) (*StateTransitionVerifierGroth16, error) {
	contract, err := bindStateTransitionVerifierGroth16(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateTransitionVerifierGroth16{StateTransitionVerifierGroth16Caller: StateTransitionVerifierGroth16Caller{contract: contract}, StateTransitionVerifierGroth16Transactor: StateTransitionVerifierGroth16Transactor{contract: contract}, StateTransitionVerifierGroth16Filterer: StateTransitionVerifierGroth16Filterer{contract: contract}}, nil
}

// NewStateTransitionVerifierGroth16Caller creates a new read-only instance of StateTransitionVerifierGroth16, bound to a specific deployed contract.
func NewStateTransitionVerifierGroth16Caller(address common.Address, caller bind.ContractCaller) (*StateTransitionVerifierGroth16Caller, error) {
	contract, err := bindStateTransitionVerifierGroth16(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateTransitionVerifierGroth16Caller{contract: contract}, nil
}

// NewStateTransitionVerifierGroth16Transactor creates a new write-only instance of StateTransitionVerifierGroth16, bound to a specific deployed contract.
func NewStateTransitionVerifierGroth16Transactor(address common.Address, transactor bind.ContractTransactor) (*StateTransitionVerifierGroth16Transactor, error) {
	contract, err := bindStateTransitionVerifierGroth16(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateTransitionVerifierGroth16Transactor{contract: contract}, nil
}

// NewStateTransitionVerifierGroth16Filterer creates a new log filterer instance of StateTransitionVerifierGroth16, bound to a specific deployed contract.
func NewStateTransitionVerifierGroth16Filterer(address common.Address, filterer bind.ContractFilterer) (*StateTransitionVerifierGroth16Filterer, error) {
	contract, err := bindStateTransitionVerifierGroth16(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateTransitionVerifierGroth16Filterer{contract: contract}, nil
}

// bindStateTransitionVerifierGroth16 binds a generic wrapper to an already deployed contract.
func bindStateTransitionVerifierGroth16(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StateTransitionVerifierGroth16MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateTransitionVerifierGroth16.Contract.StateTransitionVerifierGroth16Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateTransitionVerifierGroth16.Contract.StateTransitionVerifierGroth16Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateTransitionVerifierGroth16.Contract.StateTransitionVerifierGroth16Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateTransitionVerifierGroth16.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateTransitionVerifierGroth16.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateTransitionVerifierGroth16.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "compressProof", proof, commitments, commitmentPok)

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
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) CompressProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _StateTransitionVerifierGroth16.Contract.CompressProof(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok)
}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) CompressProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _StateTransitionVerifierGroth16.Contract.CompressProof(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) ProvingKeyHash() ([32]byte, error) {
	return _StateTransitionVerifierGroth16.Contract.ProvingKeyHash(&_StateTransitionVerifierGroth16.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) ProvingKeyHash() ([32]byte, error) {
	return _StateTransitionVerifierGroth16.Contract.ProvingKeyHash(&_StateTransitionVerifierGroth16.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe33974a8.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[4] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [4]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, compressedCommitments, compressedCommitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe33974a8.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[4] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [4]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyCompressedProof(&_StateTransitionVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe33974a8.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[4] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [4]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyCompressedProof(&_StateTransitionVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyProof(opts *bind.CallOpts, _proof []byte, _input []byte) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyProof", _proof, _input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyProof(_proof []byte, _input []byte) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof(&_StateTransitionVerifierGroth16.CallOpts, _proof, _input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyProof(_proof []byte, _input []byte) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof(&_StateTransitionVerifierGroth16.CallOpts, _proof, _input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xd148354c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[4] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyProof0(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [4]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyProof0", proof, commitments, commitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xd148354c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[4] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyProof0(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [4]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof0(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xd148354c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[4] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyProof0(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [4]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof0(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}
