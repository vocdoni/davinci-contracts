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
	Bin: "0x6080604052348015600f57600080fd5b50611ee08061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063233ace111461005c578063b1c3a00e1461008f578063b8e72af6146100b1578063d148354c146100c6578063e33974a8146100d9575b600080fd5b6040517f65ded4dc5bdc5a70bacd99e3bf2d8763c938d414a06c74173d5b2cf34b3ef7ee81526020015b60405180910390f35b6100a261009d36600461198e565b6100ec565b60405161008693929190611a00565b6100c46100bf366004611a8f565b610175565b005b6100c46100d4366004611b11565b6101f6565b6100c46100e7366004611b6d565b610698565b6100f46118d3565b6100fc6118f1565b600061011186358760015b6020020135610c00565b835261012f6060870135604088013560a089013560808a0135610cf5565b6020850152604084015261014960c0870135876007610107565b606084015261015b8535866001610107565b825261016a8435856001610107565b905093509350939050565b60008060006101848787610fec565b919450925090503063d148354c84848461019e8a8a61101d565b6040518563ffffffff1660e01b81526004016101bd9493929190611bf0565b60006040518083038186803b1580156101d557600080fd5b505afa1580156101e9573d6000803e3d6000fd5b5050505050505050505050565b6101fe6118f1565b6040805160028082526060808301845292602083019080368337019050509050602081016040848101823750604051600080516020611e8b83398151915290610253908735906020808a013591869101611c4f565b6040516020818303038152906040528051906020012060001c6102769190611ca8565b8252604080516000918782377f17f99d996509d3a1877f08a6c648707591207af6a8285e4cd07f8c5b9952fb6460408201527f099af4a793091aa57110def2ad1285be0f4059b26776f7d9bb54cdefc92388e560608201527f1016099853aa7b3e4db622c890ccf4efd8614fe255813ca2e49b7b8cf87b896160808201527f2ed40b2d824d367863f3a67e23fe38c9a8b2761298e1445211e40935985921ef60a082015260408660c08301377f1c798bca6aa96598854c891d9b4261380530942261619de249feaef8e77f91cc6101008201527f23fd4ad25d5dcd9ed55fe8559326d0afc54da78702591664bc8040a82775317a6101208201527f20e84e96fac27d035529c5d57320175ecf761b92a960ab582b7e15335e2643986101408201527f038923dd4108538532043d765d162d957ae46c2dbf7b059eea81c397ca3e47936101608201526020816101808360085afa9051169050806103ec576040516351d49ff760e11b815260040160405180910390fd5b60008061042386868a6002806020026040519081016040528092919082600260200280828437600092019190915250611038915050565b915091506040516101008a82377f0e9be5ee6f87c6535e964bef13c108b84a0921291a01dba0272c0d0811ebd4d66101008201527f035ae78bdfd140971c1d6da0597fc08b5811163a783a0cb71c7b5074387119716101208201527f0fbfb73a6e7566195ce5fc75886354712107ed83958429f178c801d4d4a82a056101408201527f04999d0dcc4affdbc549a46f70eae2971fd22af7cdba2c65ee8c78d13c8cb1e46101608201527e0ff331b9243d4d9f3c1a5debd8e2dd2add40e39d6a0204bad8f47048a77edd6101808201527f0fa0944ed513ffc15e4a8a4e5834ba43a810964ee97f88de0838c490a8b00ae66101a08201527f2e4d13d2b457ea657ab3b792edf02275a1f97b25267efc4e053a25cc60f0495c6101c08201527f0f10f2fb4d9754fba377fa460ac9e754a2d85a78e2169d6284975521f4fc4bff6101e08201527f104c8570ccfc156fcdcca39255d7d85d4dd2eb7400f8d3add325cbb7c6d2cc286102008201527f089da7b303348577700bac6953ac610cd3306efa0736c71a103b99b251d85ef961022082015282610240820152816102608201527f1b817c4559f76b5a8bbf6e4da5279448df6a2185e9c47e4de805fd095edd767f6102808201527f0c00c5314d9792b2e9cd5a49820307e3f3cfbc83a4733abbffd1c057446f2c476102a08201527f1e60d678fa467f38f26eca0769bc594e6d11294cb566482c306ef3fa80cba8a06102c08201527f115f2dbee0ffa074fa84feae3737b617d11f2b3690ce2774d70f7a29bf6bf6ed6102e08201526020816103008360085afa90511692508261068d57604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b6106a06118f1565b6106a861190f565b6106b061192d565b6106c18660005b602002013561137d565b602084015282526000806106d48761137d565b6040805160028082526060808301845294965092945091906020830190803683370190505090506020810160408881018237508451602080870151604051600080516020611e8b8339815191529361073193909291869101611c4f565b6040516020818303038152906040528051906020012060001c6107549190611ca8565b865284518452602085810151858201527f17f99d996509d3a1877f08a6c648707591207af6a8285e4cd07f8c5b9952fb646040808701919091527f099af4a793091aa57110def2ad1285be0f4059b26776f7d9bb54cdefc92388e560608701527f1016099853aa7b3e4db622c890ccf4efd8614fe255813ca2e49b7b8cf87b896160808701527f2ed40b2d824d367863f3a67e23fe38c9a8b2761298e1445211e40935985921ef60a087015260c0860185905260e086018490527f1c798bca6aa96598854c891d9b4261380530942261619de249feaef8e77f91cc6101008701527f23fd4ad25d5dcd9ed55fe8559326d0afc54da78702591664bc8040a82775317a6101208701527f20e84e96fac27d035529c5d57320175ecf761b92a960ab582b7e15335e2643986101408701527f038923dd4108538532043d765d162d957ae46c2dbf7b059eea81c397ca3e479361016087015251600091816101808860085afa9051169050806108da576040516351d49ff760e11b815260040160405180910390fd5b505050506000806108f7896000600481106106b7576106b7611bb7565b9092509050600080808061091360408e013560208f0135611423565b9296509094509250905060008061092b8f60036106b7565b9150915060008061093d8e8e8e611038565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f0e9be5ee6f87c6535e964bef13c108b84a0921291a01dba0272c0d0811ebd4d66101008e01527f035ae78bdfd140971c1d6da0597fc08b5811163a783a0cb71c7b5074387119716101208e01527f0fbfb73a6e7566195ce5fc75886354712107ed83958429f178c801d4d4a82a056101408e01527f04999d0dcc4affdbc549a46f70eae2971fd22af7cdba2c65ee8c78d13c8cb1e46101608e01527e0ff331b9243d4d9f3c1a5debd8e2dd2add40e39d6a0204bad8f47048a77edd6101808e01527f0fa0944ed513ffc15e4a8a4e5834ba43a810964ee97f88de0838c490a8b00ae66101a08e01527f2e4d13d2b457ea657ab3b792edf02275a1f97b25267efc4e053a25cc60f0495c6101c08e01527f0f10f2fb4d9754fba377fa460ac9e754a2d85a78e2169d6284975521f4fc4bff6101e08e01527f104c8570ccfc156fcdcca39255d7d85d4dd2eb7400f8d3add325cbb7c6d2cc286102008e01527f089da7b303348577700bac6953ac610cd3306efa0736c71a103b99b251d85ef96102208e01526102408d018290526102608d018190527f1b817c4559f76b5a8bbf6e4da5279448df6a2185e9c47e4de805fd095edd767f6102808e01527f0c00c5314d9792b2e9cd5a49820307e3f3cfbc83a4733abbffd1c057446f2c476102a08e01527f1e60d678fa467f38f26eca0769bc594e6d11294cb566482c306ef3fa80cba8a06102c08e01527f115f2dbee0ffa074fa84feae3737b617d11f2b3690ce2774d70f7a29bf6bf6ed6102e08e015290925090506000610bb16118f1565b6020816103008f60085afa9150811580610bcd57508051600114155b15610beb57604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b6000600080516020611e6b83398151915283101580610c2d5750600080516020611e6b8339815191528210155b15610c4b57604051631ff3747d60e21b815260040160405180910390fd5b82158015610c57575081155b15610c6457506000610cef565b6000610ca3600080516020611e6b8339815191526003600080516020611e6b83398151915287600080516020611e6b833981519152898a090908611609565b9050808303610cb8575050600182901b610cef565b610cc18161166d565b8303610cd4575050600182811b17610cef565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b600080600080516020611e6b83398151915286101580610d235750600080516020611e6b8339815191528510155b80610d3c5750600080516020611e6b8339815191528410155b80610d555750600080516020611e6b8339815191528310155b15610d7357604051631ff3747d60e21b815260040160405180910390fd5b82848688171717600003610d8c57506000905080610fe3565b60008080600080516020611e6b833981519152610db86003600080516020611e6b833981519152611cca565b600080516020611e6b8339815191528a8c090990506000600080516020611e6b8339815191528a600080516020611e6b8339815191528c8d090990506000600080516020611e6b8339815191528a600080516020611e6b8339815191528c8d09099050600080516020611e6b83398151915280600080516020611e6b8339815191528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610eb1600080516020611e6b83398151915280600080516020611e6b8339815191528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750861166d565b9350505050600080610f02600080516020611e6b83398151915280610ed857610ed8611c92565b600080516020611e6b833981519152858609600080516020611e6b83398151915287880908611609565b9050610f4f600080516020611e6b8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611e6b83398151915284880809611686565b15915050610f5e8383836116d0565b90935091508683148015610f7157508186145b15610f9b5780610f82576000610f85565b60025b60ff1660028a901b176000179450879350610fdf565b610fa48361166d565b87148015610fb95750610fb68261166d565b86145b15610cd45780610fca576000610fcd565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b610ff461194c565b610ffc61190f565b61100461190f565b61101084860186611d80565b9250925092509250925092565b6110256118d3565b61103182840184611e02565b9392505050565b6000806000600190506040516040810160007f1001c94d5433271edc60a846c243413ae240571dbea43a1762cc9844c3291caf83527f12d7facd78c75023648376bf93f62e2d26a8df4f07ebbf43494736a9287e6f206020840152865182526020870151602083015260408360808560065afa841693507f1b6245c2f20be71b1b97ceaa83586d36fe89ea1a1ee8f66abaeaacfbc4e6030982527f189899d667c67b9c9013672978cd595422b6bdb104f0466138c8aa97537d8a4f602083015288359050806040830152600080516020611e8b83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f22e1228889a8e26941b0acdc110dd37d97efbefa05f1f3f7e3d77921950f46dc82527f1a1fc800fc32501759f04cde450487512b6464269219e2c0c65db16500836838602083015260208901359050806040830152600080516020611e8b83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f21f3647f66bbeb662cb673b0720db4f7ab45a2e7d62328216f3ad1ea63a6c59382527f148bc83185a702d7eb99b0d703c3af85c07243e589135cf4104bb9387c6b8c78602083015260408901359050806040830152600080516020611e8b83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f25bf3f522122c0c08b21ec7bbc6ea8d8ffc50c881b2bd7595332fd590163aa7282527f1cc0556c84bb4778aa91571121d6749fbcb39bc3c27a31d8697f54a702334341602083015260608901359050806040830152600080516020611e8b83398151915281108416935060408260608460075afa8416935060408360808560065afa7f1420acfedc44cc438d3ecc746e28412aaac5020284b3eb868ec669c4ccf985d483527ecab036688ad665979b2f71ca49a6e891f8e7b59b8513c2278f562a65415d3d602084015288516040808501829052600080516020611e8b83398151915290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806113745760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b6000808260000361139357506000928392509050565b600183811c925080841614600080516020611e6b83398151915283106113cc57604051631ff3747d60e21b815260040160405180910390fd5b611409600080516020611e6b8339815191526003600080516020611e6b83398151915286600080516020611e6b8339815191528889090908611609565b9150801561141d5761141a8261166d565b91505b50915091565b600080808085158015611434575084155b1561144a57506000925082915081905080611600565b600286811c94508593506001808816149080881614600080516020611e6b8339815191528610158061148a5750600080516020611e6b8339815191528510155b156114a857604051631ff3747d60e21b815260040160405180910390fd5b6000600080516020611e6b8339815191526114d26003600080516020611e6b833981519152611cca565b600080516020611e6b833981519152888a090990506000600080516020611e6b83398151915288600080516020611e6b8339815191528a8b090990506000600080516020611e6b83398151915288600080516020611e6b8339815191528a8b09099050600080516020611e6b83398151915280600080516020611e6b8339815191528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50896506115cb600080516020611e6b83398151915280600080516020611e6b8339815191528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750861166d565b95506115d88787866116d0565b909750955084156115fa576115ec8761166d565b96506115f78661166d565b95505b50505050505b92959194509250565b6000611635827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5261180e565b905081600080516020611e6b8339815191528283091461166857604051631ff3747d60e21b815260040160405180910390fd5b919050565b600080516020611e6b8339815191529081900681030690565b6000806116b3837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5261180e565b905082600080516020611e6b833981519152828309149392505050565b60008080611702600080516020611e6b83398151915280878809600080516020611e6b833981519152898a0908611609565b90508315611716576117138161166d565b90505b611761600080516020611e6b8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611e6b833981519152848a0809611609565b9250600080516020611e6b83398151915261178d600080516020611e6b83398151915260028609611873565b86099150600080516020611e6b8339815191526117ba600080516020611e6b83398151915284850961166d565b600080516020611e6b83398151915285860908861415806117f05750600080516020611e6b833981519152808385096002098514155b1561137457604051631ff3747d60e21b815260040160405180910390fd5b6000806040516020815260208082015260206040820152846060820152836080820152600080516020611e6b83398151915260a082015260208160c08360055afa90519250905080610ced57604051631ff3747d60e21b815260040160405180910390fd5b600061189f827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4561180e565b9050600080516020611e6b83398151915281830960011461166857604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b806101008101831015610cef57600080fd5b8060408101831015610cef57600080fd5b600080600061018084860312156119a457600080fd5b6119ae858561196b565b92506119be85610100860161197d565b91506119ce85610140860161197d565b90509250925092565b8060005b60048110156119fa5781518452602093840193909101906001016119db565b50505050565b60c08101611a0e82866119d7565b608082018460005b6001811015611a35578151835260209283019290910190600101611a16565b5050508260a0830152949350505050565b60008083601f840112611a5857600080fd5b50813567ffffffffffffffff811115611a7057600080fd5b602083019150836020828501011115611a8857600080fd5b9250929050565b60008060008060408587031215611aa557600080fd5b843567ffffffffffffffff811115611abc57600080fd5b611ac887828801611a46565b909550935050602085013567ffffffffffffffff811115611ae857600080fd5b611af487828801611a46565b95989497509550505050565b8060808101831015610cef57600080fd5b6000806000806102008587031215611b2857600080fd5b611b32868661196b565b9350611b4286610100870161197d565b9250611b5286610140870161197d565b9150611b62866101808701611b00565b905092959194509250565b6000806000806101408587031215611b8457600080fd5b611b8e8686611b00565b935060a0850186811115611ba157600080fd5b608086019350359150611b628660c08701611b00565b634e487b7160e01b600052603260045260246000fd5b8060005b60028110156119fa578151845260209384019390910190600101611bd1565b6102008101818660005b6008811015611c19578151835260209283019290910190600101611bfa565b505050611c2a610100830186611bcd565b611c38610140830185611bcd565b611c466101808301846119d7565b95945050505050565b83815282602082015260006040820183516020850160005b82811015611c85578151845260209384019390910190600101611c67565b5091979650505050505050565b634e487b7160e01b600052601260045260246000fd5b600082611cc557634e487b7160e01b600052601260045260246000fd5b500690565b81810381811115610cef57634e487b7160e01b600052601160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611d2257634e487b7160e01b600052604160045260246000fd5b604052919050565b600082601f830112611d3b57600080fd5b6000611d476040611ceb565b9050806040840185811115611d5b57600080fd5b845b81811015611d75578035835260209283019201611d5d565b509195945050505050565b60008060006101808486031215611d9657600080fd5b600085601f860112611da6578081fd5b80610100611db381611ceb565b9250829150860187811115611dc757600080fd5b865b81811015611de1578035845260209384019301611dc9565b50819550611def8882611d2a565b94505050506119ce856101408601611d2a565b600060808284031215611e1457600080fd5b600083601f840112611e24578081fd5b80611e2f6080611ceb565b90508091506080840185811115611e4557600080fd5b845b81811015611e5f578035845260209384019301611e47565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212208857002d20824c45c583a6ec6e74e4fac5d6e4dc8e0938f762b334cfd6607c5664736f6c634300081c0033",
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
