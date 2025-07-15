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
	Bin: "0x6080604052348015600f57600080fd5b50611ee98061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063233ace111461005c578063b1c3a00e1461008f578063b8e72af6146100b1578063d148354c146100c6578063e33974a8146100d9575b600080fd5b6040517f7d051009ee4bdde7dc37085b5eeca36e18fe9350c689ca6cdd5e537705fc5efe81526020015b60405180910390f35b6100a261009d366004611997565b6100ec565b60405161008693929190611a09565b6100c46100bf366004611a98565b610175565b005b6100c46100d4366004611b1a565b6101f6565b6100c46100e7366004611b76565b61069c565b6100f46118dc565b6100fc6118fa565b600061011186358760015b6020020135610c08565b835261012f6060870135604088013560a089013560808a0135610cfd565b6020850152604084015261014960c0870135876007610107565b606084015261015b8535866001610107565b825261016a8435856001610107565b905093509350939050565b60008060006101848787610ff4565b919450925090503063d148354c84848461019e8a8a611025565b6040518563ffffffff1660e01b81526004016101bd9493929190611bf9565b60006040518083038186803b1580156101d557600080fd5b505afa1580156101e9573d6000803e3d6000fd5b5050505050505050505050565b6101fe6118fa565b6040805160018082528183019092526060916020808301908036833701905050905060208082019060408501823750604051600080516020611e9483398151915290610256908735906020808a013591869101611c58565b6040516020818303038152906040528051906020012060001c6102799190611cb1565b8252604080516000918782377f3015a7844f93ffc868339ed133711db1e2f607bfaaf7c1d0ac69f46406de3ed960408201527f05420fddd408e12baebc9b71a2dc2bda13d46441141f986d8482a1ff778a2d1d60608201527f12949be1f4ecb86dcbaf648977452ac9ce9c0319dd6a5c3aadb7bc57c99d064a60808201527f23715dc774469021e45801227bcf7d0b979c32d4e023be3d6737d79a353508ca60a082015260408660c08301377f2941b46d10b2ceb33fec211da693cb4eee86d16bb092695eb313f85a8a1d4d596101008201527f1446b12ca9c28c64895ecd1c50a12c5f47251544e37ac0037c33c45a659b6e326101208201527f2f9f0251bc29a050c473fa576ecef0d072388d45883eb9d85edebd6def79c1d76101408201527f05eebf0cce919a2d90faa16fe4308f834bc4ba20af1088433050a6748ef5381d6101608201526020816101808360085afa9051169050806103ef576040516351d49ff760e11b815260040160405180910390fd5b60008061042686868a6002806020026040519081016040528092919082600260200280828437600092019190915250611040915050565b915091506040516101008a82377f14126a2627081476380525f8c47422f16da067ca72052ca6e6d090f5d4161bda6101008201527f0bad53b001700ab3d94f31adf720466bcf03b686e76a91a056d1990f261e6c696101208201527f0ea6d3a4517a1cf7847c09b300538f1a92d5489b73476847549ba85e27c582106101408201527f2c0364f790cfbf0af6119c3945eb8b1650559aabe96dec3c2725b89eb1df2a866101608201527f065500fe2e1b2043e4be23ddbb26c1f86016749045b7b3995762b6c0f20cb1896101808201527f154edc337988eacbc206168aa0dfb93c1e62bf66a4d8918d4fa09172f8ca8ce46101a08201527f0cf64179b832af7cded840b22edbdce7e74be8e19990ea07ac2747ad628e1e8c6101c08201527f0ebe51c4e4e5bff9914c8e9d97af4896adbaaa98f053940491b978e59e695bb26101e08201527f2d986de9b7a778b5ba9a4feb955d6f58d46d24ea0c01dd006f4cc0af0da7ccac6102008201527f29dbd0cfa1105f82fe4fe882c13d68b3e7b5ab0a996d06f456e5a8e879f2927361022082015282610240820152816102608201527f2a15c11f0ea27d2f5bdf3d100fe769cb01e325d038f06ece7bcdec7c7f296f1c6102808201527f16fd4771be9232658b0f51e6c975161a912cdf5b9c144f359b51bca90c0721506102a08201527f1b9eef24c8818a223c0662a68b479b0faf803e7371da862ce3b3225eb0b2381a6102c08201527f2ffe7b0551a62bc6d27a9c1268bbf1017edaf701c3782eea6b39c39c800f470f6102e08201526020816103008360085afa90511692508261069157604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b6106a46118fa565b6106ac611918565b6106b4611936565b6106c58660005b6020020135611386565b602084015282526000806106d887611386565b6040805160018082528183019092529294509092506060919060208083019080368337019050509050602080820190604089018237508451602080870151604051600080516020611e948339815191529361073893909291869101611c58565b6040516020818303038152906040528051906020012060001c61075b9190611cb1565b865284518452602085810151858201527f3015a7844f93ffc868339ed133711db1e2f607bfaaf7c1d0ac69f46406de3ed96040808701919091527f05420fddd408e12baebc9b71a2dc2bda13d46441141f986d8482a1ff778a2d1d60608701527f12949be1f4ecb86dcbaf648977452ac9ce9c0319dd6a5c3aadb7bc57c99d064a60808701527f23715dc774469021e45801227bcf7d0b979c32d4e023be3d6737d79a353508ca60a087015260c0860185905260e086018490527f2941b46d10b2ceb33fec211da693cb4eee86d16bb092695eb313f85a8a1d4d596101008701527f1446b12ca9c28c64895ecd1c50a12c5f47251544e37ac0037c33c45a659b6e326101208701527f2f9f0251bc29a050c473fa576ecef0d072388d45883eb9d85edebd6def79c1d76101408701527f05eebf0cce919a2d90faa16fe4308f834bc4ba20af1088433050a6748ef5381d61016087015251600091816101808860085afa9051169050806108e1576040516351d49ff760e11b815260040160405180910390fd5b505050506000806108fe896000600481106106bb576106bb611bc0565b9092509050600080808061091a60408e013560208f013561142c565b929650909450925090506000806109328f60036106bb565b915091506000806109448e8e8e611040565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f14126a2627081476380525f8c47422f16da067ca72052ca6e6d090f5d4161bda6101008e01527f0bad53b001700ab3d94f31adf720466bcf03b686e76a91a056d1990f261e6c696101208e01527f0ea6d3a4517a1cf7847c09b300538f1a92d5489b73476847549ba85e27c582106101408e01527f2c0364f790cfbf0af6119c3945eb8b1650559aabe96dec3c2725b89eb1df2a866101608e01527f065500fe2e1b2043e4be23ddbb26c1f86016749045b7b3995762b6c0f20cb1896101808e01527f154edc337988eacbc206168aa0dfb93c1e62bf66a4d8918d4fa09172f8ca8ce46101a08e01527f0cf64179b832af7cded840b22edbdce7e74be8e19990ea07ac2747ad628e1e8c6101c08e01527f0ebe51c4e4e5bff9914c8e9d97af4896adbaaa98f053940491b978e59e695bb26101e08e01527f2d986de9b7a778b5ba9a4feb955d6f58d46d24ea0c01dd006f4cc0af0da7ccac6102008e01527f29dbd0cfa1105f82fe4fe882c13d68b3e7b5ab0a996d06f456e5a8e879f292736102208e01526102408d018290526102608d018190527f2a15c11f0ea27d2f5bdf3d100fe769cb01e325d038f06ece7bcdec7c7f296f1c6102808e01527f16fd4771be9232658b0f51e6c975161a912cdf5b9c144f359b51bca90c0721506102a08e01527f1b9eef24c8818a223c0662a68b479b0faf803e7371da862ce3b3225eb0b2381a6102c08e01527f2ffe7b0551a62bc6d27a9c1268bbf1017edaf701c3782eea6b39c39c800f470f6102e08e015290925090506000610bb96118fa565b6020816103008f60085afa9150811580610bd557508051600114155b15610bf357604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b6000600080516020611e7483398151915283101580610c355750600080516020611e748339815191528210155b15610c5357604051631ff3747d60e21b815260040160405180910390fd5b82158015610c5f575081155b15610c6c57506000610cf7565b6000610cab600080516020611e748339815191526003600080516020611e7483398151915287600080516020611e74833981519152898a090908611612565b9050808303610cc0575050600182901b610cf7565b610cc981611676565b8303610cdc575050600182811b17610cf7565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b600080600080516020611e7483398151915286101580610d2b5750600080516020611e748339815191528510155b80610d445750600080516020611e748339815191528410155b80610d5d5750600080516020611e748339815191528310155b15610d7b57604051631ff3747d60e21b815260040160405180910390fd5b82848688171717600003610d9457506000905080610feb565b60008080600080516020611e74833981519152610dc06003600080516020611e74833981519152611cd3565b600080516020611e748339815191528a8c090990506000600080516020611e748339815191528a600080516020611e748339815191528c8d090990506000600080516020611e748339815191528a600080516020611e748339815191528c8d09099050600080516020611e7483398151915280600080516020611e748339815191528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610eb9600080516020611e7483398151915280600080516020611e748339815191528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611676565b9350505050600080610f0a600080516020611e7483398151915280610ee057610ee0611c9b565b600080516020611e74833981519152858609600080516020611e7483398151915287880908611612565b9050610f57600080516020611e748339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611e748339815191528488080961168f565b15915050610f668383836116d9565b90935091508683148015610f7957508186145b15610fa35780610f8a576000610f8d565b60025b60ff1660028a901b176000179450879350610fe7565b610fac83611676565b87148015610fc15750610fbe82611676565b86145b15610cdc5780610fd2576000610fd5565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b610ffc611955565b611004611918565b61100c611918565b61101884860186611d89565b9250925092509250925092565b61102d6118dc565b61103982840184611e0b565b9392505050565b6000806000600190506040516040810160007f220aa929812cd5153eb1e0eff6383d271e35fe0bdc1b05d3b3b56502a64861ad83527f1be69e3da6d0751ea666e831a266e5d80e44c3bb74ed0ac01f56803ee1e484886020840152865182526020870151602083015260408360808560065afa841693507f17ae48190efadf2d375c27f0e7596ae390bfdd952fde9496ce20927fa80f70a582527f03d7ee9bed605c7fc3a3aab4a6c01e1aa741f4f6d15cacc11593149e0f81eed6602083015288359050806040830152600080516020611e9483398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f1b00e3d4bc6143b9da3ef5178fe99bae576a1ae33df5b9a3aa00694eed3afb2082527f2a88546aa42a856c3b4b5676f2ec6fb2b674004df07de4917981340ecaf4665c602083015260208901359050806040830152600080516020611e9483398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f1bc913ffd4f05b49a32c1021bd0c97f0d31c1dd155b8df39bd8a65855278530e82527f227578cea63c2b28ec38734c95faa106d1c8bc7c3981f3f9f5e8b73a9494ef85602083015260408901359050806040830152600080516020611e9483398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f1d86268edadc5136f4024ccba4df39f1334bf926e6a1fabedca22f4e4e4396e282527f202d128bfa069837c8168e799c5185a4c8bb26486ffcf4283f91f1983d13ac46602083015260608901359050806040830152600080516020611e9483398151915281108416935060408260608460075afa8416935060408360808560065afa7f2491a22385718e35e5fafe4833abe7d0dbc3ecf4111e559687528157c4bf674083527f2b5d1c3951f31e4892a6a4d4507a7e5e4e2187dbe8efb1dee40a36182e1544f6602084015288516040808501829052600080516020611e9483398151915290911091909516169390508160608160075afa831692505060408160808360065afa8151602090920151919450909250168061137d5760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b6000808260000361139c57506000928392509050565b600183811c925080841614600080516020611e7483398151915283106113d557604051631ff3747d60e21b815260040160405180910390fd5b611412600080516020611e748339815191526003600080516020611e7483398151915286600080516020611e748339815191528889090908611612565b915080156114265761142382611676565b91505b50915091565b60008080808515801561143d575084155b1561145357506000925082915081905080611609565b600286811c94508593506001808816149080881614600080516020611e74833981519152861015806114935750600080516020611e748339815191528510155b156114b157604051631ff3747d60e21b815260040160405180910390fd5b6000600080516020611e748339815191526114db6003600080516020611e74833981519152611cd3565b600080516020611e74833981519152888a090990506000600080516020611e7483398151915288600080516020611e748339815191528a8b090990506000600080516020611e7483398151915288600080516020611e748339815191528a8b09099050600080516020611e7483398151915280600080516020611e748339815191528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50896506115d4600080516020611e7483398151915280600080516020611e748339815191528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611676565b95506115e18787866116d9565b90975095508415611603576115f587611676565b965061160086611676565b95505b50505050505b92959194509250565b600061163e827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611817565b905081600080516020611e748339815191528283091461167157604051631ff3747d60e21b815260040160405180910390fd5b919050565b600080516020611e748339815191529081900681030690565b6000806116bc837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611817565b905082600080516020611e74833981519152828309149392505050565b6000808061170b600080516020611e7483398151915280878809600080516020611e74833981519152898a0908611612565b9050831561171f5761171c81611676565b90505b61176a600080516020611e748339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611e74833981519152848a0809611612565b9250600080516020611e74833981519152611796600080516020611e748339815191526002860961187c565b86099150600080516020611e748339815191526117c3600080516020611e74833981519152848509611676565b600080516020611e7483398151915285860908861415806117f95750600080516020611e74833981519152808385096002098514155b1561137d57604051631ff3747d60e21b815260040160405180910390fd5b6000806040516020815260208082015260206040820152846060820152836080820152600080516020611e7483398151915260a082015260208160c08360055afa90519250905080610cf557604051631ff3747d60e21b815260040160405180910390fd5b60006118a8827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611817565b9050600080516020611e7483398151915281830960011461167157604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b806101008101831015610cf757600080fd5b8060408101831015610cf757600080fd5b600080600061018084860312156119ad57600080fd5b6119b78585611974565b92506119c7856101008601611986565b91506119d7856101408601611986565b90509250925092565b8060005b6004811015611a035781518452602093840193909101906001016119e4565b50505050565b60c08101611a1782866119e0565b608082018460005b6001811015611a3e578151835260209283019290910190600101611a1f565b5050508260a0830152949350505050565b60008083601f840112611a6157600080fd5b50813567ffffffffffffffff811115611a7957600080fd5b602083019150836020828501011115611a9157600080fd5b9250929050565b60008060008060408587031215611aae57600080fd5b843567ffffffffffffffff811115611ac557600080fd5b611ad187828801611a4f565b909550935050602085013567ffffffffffffffff811115611af157600080fd5b611afd87828801611a4f565b95989497509550505050565b8060808101831015610cf757600080fd5b6000806000806102008587031215611b3157600080fd5b611b3b8686611974565b9350611b4b866101008701611986565b9250611b5b866101408701611986565b9150611b6b866101808701611b09565b905092959194509250565b6000806000806101408587031215611b8d57600080fd5b611b978686611b09565b935060a0850186811115611baa57600080fd5b608086019350359150611b6b8660c08701611b09565b634e487b7160e01b600052603260045260246000fd5b8060005b6002811015611a03578151845260209384019390910190600101611bda565b6102008101818660005b6008811015611c22578151835260209283019290910190600101611c03565b505050611c33610100830186611bd6565b611c41610140830185611bd6565b611c4f6101808301846119e0565b95945050505050565b83815282602082015260006040820183516020850160005b82811015611c8e578151845260209384019390910190600101611c70565b5091979650505050505050565b634e487b7160e01b600052601260045260246000fd5b600082611cce57634e487b7160e01b600052601260045260246000fd5b500690565b81810381811115610cf757634e487b7160e01b600052601160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611d2b57634e487b7160e01b600052604160045260246000fd5b604052919050565b600082601f830112611d4457600080fd5b6000611d506040611cf4565b9050806040840185811115611d6457600080fd5b845b81811015611d7e578035835260209283019201611d66565b509195945050505050565b60008060006101808486031215611d9f57600080fd5b600085601f860112611daf578081fd5b80610100611dbc81611cf4565b9250829150860187811115611dd057600080fd5b865b81811015611dea578035845260209384019301611dd2565b50819550611df88882611d33565b94505050506119d7856101408601611d33565b600060808284031215611e1d57600080fd5b600083601f840112611e2d578081fd5b80611e386080611cf4565b90508091506080840185811115611e4e57600080fd5b845b81811015611e68578035845260209384019301611e50565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220e8e7f1607c4b85af4d03c108bc9b593337ac8739fb73e3e95f59abdd3ca22a8e64736f6c634300081c0033",
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
