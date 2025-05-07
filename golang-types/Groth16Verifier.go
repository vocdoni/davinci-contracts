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

// Groth16VerifierMetaData contains all meta data concerning the Groth16Verifier contract.
var Groth16VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"input\",\"type\":\"uint256[4]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"input\",\"type\":\"uint256[4]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50611eac8061001f6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063b1c3a00e14610051578063b8e72af61461007c578063d148354c14610091578063e33974a8146100a4575b600080fd5b61006461005f36600461195a565b6100b7565b604051610073939291906119cc565b60405180910390f35b61008f61008a366004611a5b565b610140565b005b61008f61009f366004611add565b6101c1565b61008f6100b2366004611b39565b610663565b6100bf61189f565b6100c76118bd565b60006100dc86358760015b6020020135610bcb565b83526100fa6060870135604088013560a089013560808a0135610cc0565b6020850152604084015261011460c08701358760076100d2565b606084015261012685358660016100d2565b825261013584358560016100d2565b905093509350939050565b600080600061014f8787610fb7565b919450925090503063d148354c8484846101698a8a610fe8565b6040518563ffffffff1660e01b81526004016101889493929190611bbc565b60006040518083038186803b1580156101a057600080fd5b505afa1580156101b4573d6000803e3d6000fd5b5050505050505050505050565b6101c96118bd565b6040805160028082526060808301845292602083019080368337019050509050602081016040848101823750604051600080516020611e578339815191529061021e908735906020808a013591869101611c1b565b6040516020818303038152906040528051906020012060001c6102419190611c74565b8252604080516000918782377f2e88a9ead99f02ac4e3830eb4eca71c10101fd10de0e62019888e612716f93cb60408201527f2b1d878c768b64f6af034d9c9aeaccbc648725f9f9750e667efd249fc5dfab2960608201527e93eef58e7e91356ec1ce9a9c9567f355b86e56f736421889d4f15c4609654360808201527f20b10cdff6866c461bcb46515085d9428a2816ee3ee2031ca9619ad704d7b48a60a082015260408660c08301377f03186ad4bd167bc83beb0efc4cd4038815bb00a9fdf42ee311e034e8b54ee2a26101008201527f2be3829ff4252709e7c31196e6a5d9a41d3672c0eedc478883aa5533753ff1276101208201527f159270b78a0023cf20f7c97d98fb43561f6ac8eac954a229b6b3bc1e27839c6d6101408201527f204e9db68a76e0fdfa92cd46ee5aef9cd2d221653ba72f4756e5b058b87818b76101608201526020816101808360085afa9051169050806103b6576040516351d49ff760e11b815260040160405180910390fd5b6000806103ed86868a6002806020026040519081016040528092919082600260200280828437600092019190915250611003915050565b915091506040516101008a82377f1427461566558d8ca033b9c1baec0acf0bffcb8807742a10cb9c9cc6b400424e6101008201527f2f6c649377c0ac1ce0eaf3e4edd6f5173f194c7b8f2343542ed556fb2a93ff106101208201527f2b53268bfc91702d841693608a48dd567214f11ffbda0929095ce64103a305d86101408201527f0ed1f1b4b38d13148d34651a4f31b8d7ef87063c3779eebbc73ada577f8ca5816101608201527f2caac348b0278a2c8e6a6339c268989077718d0b264abd71eb538bcab8b5eb0b6101808201527f01d1846b70a55a4808e39c6260c72c64a744ed5facaf0f30e205f9c547cdf90b6101a08201527f1f9e366f8e419e6de5099234facd7bc318f1b040348f281123d9c1c0995a9e406101c08201527f068a15ddb3a14c6b5c737f92c685d869fe77ea7640b00e151365b0544c5f5fa26101e08201527f05401fc811ab7a7d5fa6aeb56ffa9c153f2674b6fbc4f21a18b9a966788894766102008201527f14732630c51c3c89c521fe8740a8483a9ea703b104d16881486d79d905f35aa161022082015282610240820152816102608201527f21ea42774d55f78e26601cea14871a4d6740f7932b9e6f62123c5ba4ab5f067b6102808201527f159b91eb80c20172421b1adfc2171eb2fc94df89b9677fffc2a9589a5525c59f6102a08201527f2771b5d13d999763e65d196ef22cd69c389ba903034dee73642cf6b1a45464646102c08201527f1eb07a01786302343fb7e1d6f63959bf871bf8a7b0e1d1b5836f89d6d1f09b7a6102e08201526020816103008360085afa90511692508261065857604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b61066b6118bd565b6106736118db565b61067b6118f9565b61068c8660005b6020020135611349565b6020840152825260008061069f87611349565b6040805160028082526060808301845294965092945091906020830190803683370190505090506020810160408881018237508451602080870151604051600080516020611e57833981519152936106fc93909291869101611c1b565b6040516020818303038152906040528051906020012060001c61071f9190611c74565b865284518452602085810151858201527f2e88a9ead99f02ac4e3830eb4eca71c10101fd10de0e62019888e612716f93cb6040808701919091527f2b1d878c768b64f6af034d9c9aeaccbc648725f9f9750e667efd249fc5dfab2960608701527e93eef58e7e91356ec1ce9a9c9567f355b86e56f736421889d4f15c4609654360808701527f20b10cdff6866c461bcb46515085d9428a2816ee3ee2031ca9619ad704d7b48a60a087015260c0860185905260e086018490527f03186ad4bd167bc83beb0efc4cd4038815bb00a9fdf42ee311e034e8b54ee2a26101008701527f2be3829ff4252709e7c31196e6a5d9a41d3672c0eedc478883aa5533753ff1276101208701527f159270b78a0023cf20f7c97d98fb43561f6ac8eac954a229b6b3bc1e27839c6d6101408701527f204e9db68a76e0fdfa92cd46ee5aef9cd2d221653ba72f4756e5b058b87818b761016087015251600091816101808860085afa9051169050806108a4576040516351d49ff760e11b815260040160405180910390fd5b505050506000806108c18960006004811061068257610682611b83565b909250905060008080806108dd60408e013560208f01356113ef565b929650909450925090506000806108f58f6003610682565b915091506000806109078e8e8e611003565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f1427461566558d8ca033b9c1baec0acf0bffcb8807742a10cb9c9cc6b400424e6101008e01527f2f6c649377c0ac1ce0eaf3e4edd6f5173f194c7b8f2343542ed556fb2a93ff106101208e01527f2b53268bfc91702d841693608a48dd567214f11ffbda0929095ce64103a305d86101408e01527f0ed1f1b4b38d13148d34651a4f31b8d7ef87063c3779eebbc73ada577f8ca5816101608e01527f2caac348b0278a2c8e6a6339c268989077718d0b264abd71eb538bcab8b5eb0b6101808e01527f01d1846b70a55a4808e39c6260c72c64a744ed5facaf0f30e205f9c547cdf90b6101a08e01527f1f9e366f8e419e6de5099234facd7bc318f1b040348f281123d9c1c0995a9e406101c08e01527f068a15ddb3a14c6b5c737f92c685d869fe77ea7640b00e151365b0544c5f5fa26101e08e01527f05401fc811ab7a7d5fa6aeb56ffa9c153f2674b6fbc4f21a18b9a966788894766102008e01527f14732630c51c3c89c521fe8740a8483a9ea703b104d16881486d79d905f35aa16102208e01526102408d018290526102608d018190527f21ea42774d55f78e26601cea14871a4d6740f7932b9e6f62123c5ba4ab5f067b6102808e01527f159b91eb80c20172421b1adfc2171eb2fc94df89b9677fffc2a9589a5525c59f6102a08e01527f2771b5d13d999763e65d196ef22cd69c389ba903034dee73642cf6b1a45464646102c08e01527f1eb07a01786302343fb7e1d6f63959bf871bf8a7b0e1d1b5836f89d6d1f09b7a6102e08e015290925090506000610b7c6118bd565b6020816103008f60085afa9150811580610b9857508051600114155b15610bb657604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b6000600080516020611e3783398151915283101580610bf85750600080516020611e378339815191528210155b15610c1657604051631ff3747d60e21b815260040160405180910390fd5b82158015610c22575081155b15610c2f57506000610cba565b6000610c6e600080516020611e378339815191526003600080516020611e3783398151915287600080516020611e37833981519152898a0909086115d5565b9050808303610c83575050600182901b610cba565b610c8c81611639565b8303610c9f575050600182811b17610cba565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b600080600080516020611e3783398151915286101580610cee5750600080516020611e378339815191528510155b80610d075750600080516020611e378339815191528410155b80610d205750600080516020611e378339815191528310155b15610d3e57604051631ff3747d60e21b815260040160405180910390fd5b82848688171717600003610d5757506000905080610fae565b60008080600080516020611e37833981519152610d836003600080516020611e37833981519152611c96565b600080516020611e378339815191528a8c090990506000600080516020611e378339815191528a600080516020611e378339815191528c8d090990506000600080516020611e378339815191528a600080516020611e378339815191528c8d09099050600080516020611e3783398151915280600080516020611e378339815191528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610e7c600080516020611e3783398151915280600080516020611e378339815191528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611639565b9350505050600080610ecd600080516020611e3783398151915280610ea357610ea3611c5e565b600080516020611e37833981519152858609600080516020611e37833981519152878809086115d5565b9050610f1a600080516020611e378339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611e3783398151915284880809611652565b15915050610f2983838361169c565b90935091508683148015610f3c57508186145b15610f665780610f4d576000610f50565b60025b60ff1660028a901b176000179450879350610faa565b610f6f83611639565b87148015610f845750610f8182611639565b86145b15610c9f5780610f95576000610f98565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b610fbf611918565b610fc76118db565b610fcf6118db565b610fdb84860186611d4c565b9250925092509250925092565b610ff061189f565b610ffc82840184611dce565b9392505050565b6000806000600190506040516040810160007f10580fb50381cc06e7691279cb54a4766645fa5ca68a5fc876976872ab54ff3b83527f0fd1d2b5b0f0f172efb124975df0511f527f554ca04c27796c888b7768455f6c6020840152865182526020870151602083015260408360808560065afa841693507f2afb9f53137aa3344c90d903c6035948e2615367f58d410c4922e59007d58f8882527f1b7fcf6fe6cdcf13a1b285da4f06aa7d2a73939a72c5bc1d84cb47db6b5ed812602083015288359050806040830152600080516020611e5783398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f03e24176a4b1efdbc9995fa9d620107843814eb0f6a4bf108a3850f45227896b82527f0ec5853a3a1aa90988e02d63ce645ba77134ca6c5ceb43fcc4f727bef8186385602083015260208901359050806040830152600080516020611e5783398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f1efb933304dca8cbc998cdcbda42d55e4b3dc7d6322bdfb74366117f0f784cc382527f0afb80dcd03ab6b290fe1c8380ec84806ae09f7d21fb2ee57d845ded2d3dcd23602083015260408901359050806040830152600080516020611e5783398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f23ce24189142de2e5a241058ab58aeaadbca21397055434c73ef0dde52bd4f3e82527f143a9019642556dbd4a6934e4182105ccf9aaa0deee773d54d792c8da65cc3ad602083015260608901359050806040830152600080516020611e5783398151915281108416935060408260608460075afa8416935060408360808560065afa7f128c80e86122ebab3afdb91b9a52590e26a95e8eb4f2ce228ebd6f4837aae9d283527f2260337c46ad7589b1d0a815c879276d7a89b5e0e783662ab7989c1bcb8e8a0e602084015288516040808501829052600080516020611e5783398151915290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806113405760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b6000808260000361135f57506000928392509050565b600183811c925080841614600080516020611e37833981519152831061139857604051631ff3747d60e21b815260040160405180910390fd5b6113d5600080516020611e378339815191526003600080516020611e3783398151915286600080516020611e3783398151915288890909086115d5565b915080156113e9576113e682611639565b91505b50915091565b600080808085158015611400575084155b15611416575060009250829150819050806115cc565b600286811c94508593506001808816149080881614600080516020611e37833981519152861015806114565750600080516020611e378339815191528510155b1561147457604051631ff3747d60e21b815260040160405180910390fd5b6000600080516020611e3783398151915261149e6003600080516020611e37833981519152611c96565b600080516020611e37833981519152888a090990506000600080516020611e3783398151915288600080516020611e378339815191528a8b090990506000600080516020611e3783398151915288600080516020611e378339815191528a8b09099050600080516020611e3783398151915280600080516020611e378339815191528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089650611597600080516020611e3783398151915280600080516020611e378339815191528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611639565b95506115a487878661169c565b909750955084156115c6576115b887611639565b96506115c386611639565b95505b50505050505b92959194509250565b6000611601827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526117da565b905081600080516020611e378339815191528283091461163457604051631ff3747d60e21b815260040160405180910390fd5b919050565b600080516020611e378339815191529081900681030690565b60008061167f837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526117da565b905082600080516020611e37833981519152828309149392505050565b600080806116ce600080516020611e3783398151915280878809600080516020611e37833981519152898a09086115d5565b905083156116e2576116df81611639565b90505b61172d600080516020611e378339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611e37833981519152848a08096115d5565b9250600080516020611e37833981519152611759600080516020611e378339815191526002860961183f565b86099150600080516020611e37833981519152611786600080516020611e37833981519152848509611639565b600080516020611e3783398151915285860908861415806117bc5750600080516020611e37833981519152808385096002098514155b1561134057604051631ff3747d60e21b815260040160405180910390fd5b6000806040516020815260208082015260206040820152846060820152836080820152600080516020611e3783398151915260a082015260208160c08360055afa90519250905080610cb857604051631ff3747d60e21b815260040160405180910390fd5b600061186b827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd456117da565b9050600080516020611e3783398151915281830960011461163457604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b806101008101831015610cba57600080fd5b8060408101831015610cba57600080fd5b6000806000610180848603121561197057600080fd5b61197a8585611937565b925061198a856101008601611949565b915061199a856101408601611949565b90509250925092565b8060005b60048110156119c65781518452602093840193909101906001016119a7565b50505050565b60c081016119da82866119a3565b608082018460005b6001811015611a015781518352602092830192909101906001016119e2565b5050508260a0830152949350505050565b60008083601f840112611a2457600080fd5b50813567ffffffffffffffff811115611a3c57600080fd5b602083019150836020828501011115611a5457600080fd5b9250929050565b60008060008060408587031215611a7157600080fd5b843567ffffffffffffffff811115611a8857600080fd5b611a9487828801611a12565b909550935050602085013567ffffffffffffffff811115611ab457600080fd5b611ac087828801611a12565b95989497509550505050565b8060808101831015610cba57600080fd5b6000806000806102008587031215611af457600080fd5b611afe8686611937565b9350611b0e866101008701611949565b9250611b1e866101408701611949565b9150611b2e866101808701611acc565b905092959194509250565b6000806000806101408587031215611b5057600080fd5b611b5a8686611acc565b935060a0850186811115611b6d57600080fd5b608086019350359150611b2e8660c08701611acc565b634e487b7160e01b600052603260045260246000fd5b8060005b60028110156119c6578151845260209384019390910190600101611b9d565b6102008101818660005b6008811015611be5578151835260209283019290910190600101611bc6565b505050611bf6610100830186611b99565b611c04610140830185611b99565b611c126101808301846119a3565b95945050505050565b83815282602082015260006040820183516020850160005b82811015611c51578151845260209384019390910190600101611c33565b5091979650505050505050565b634e487b7160e01b600052601260045260246000fd5b600082611c9157634e487b7160e01b600052601260045260246000fd5b500690565b81810381811115610cba57634e487b7160e01b600052601160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611cee57634e487b7160e01b600052604160045260246000fd5b604052919050565b600082601f830112611d0757600080fd5b6000611d136040611cb7565b9050806040840185811115611d2757600080fd5b845b81811015611d41578035835260209283019201611d29565b509195945050505050565b60008060006101808486031215611d6257600080fd5b600085601f860112611d72578081fd5b80610100611d7f81611cb7565b9250829150860187811115611d9357600080fd5b865b81811015611dad578035845260209384019301611d95565b50819550611dbb8882611cf6565b945050505061199a856101408601611cf6565b600060808284031215611de057600080fd5b600083601f840112611df0578081fd5b80611dfb6080611cb7565b90508091506080840185811115611e1157600080fd5b845b81811015611e2b578035845260209384019301611e13565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212201f5b1870834c64e02fb1df762d647134daea88fd803bef7e43ac7bcb68bf4cba64736f6c634300081c0033",
}

// Groth16VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use Groth16VerifierMetaData.ABI instead.
var Groth16VerifierABI = Groth16VerifierMetaData.ABI

// Groth16VerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Groth16VerifierMetaData.Bin instead.
var Groth16VerifierBin = Groth16VerifierMetaData.Bin

// DeployGroth16Verifier deploys a new Ethereum contract, binding an instance of Groth16Verifier to it.
func DeployGroth16Verifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Groth16Verifier, error) {
	parsed, err := Groth16VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Groth16VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Groth16Verifier{Groth16VerifierCaller: Groth16VerifierCaller{contract: contract}, Groth16VerifierTransactor: Groth16VerifierTransactor{contract: contract}, Groth16VerifierFilterer: Groth16VerifierFilterer{contract: contract}}, nil
}

// Groth16Verifier is an auto generated Go binding around an Ethereum contract.
type Groth16Verifier struct {
	Groth16VerifierCaller     // Read-only binding to the contract
	Groth16VerifierTransactor // Write-only binding to the contract
	Groth16VerifierFilterer   // Log filterer for contract events
}

// Groth16VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type Groth16VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Groth16VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Groth16VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Groth16VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Groth16VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Groth16VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Groth16VerifierSession struct {
	Contract     *Groth16Verifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Groth16VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Groth16VerifierCallerSession struct {
	Contract *Groth16VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Groth16VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Groth16VerifierTransactorSession struct {
	Contract     *Groth16VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Groth16VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type Groth16VerifierRaw struct {
	Contract *Groth16Verifier // Generic contract binding to access the raw methods on
}

// Groth16VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Groth16VerifierCallerRaw struct {
	Contract *Groth16VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// Groth16VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Groth16VerifierTransactorRaw struct {
	Contract *Groth16VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGroth16Verifier creates a new instance of Groth16Verifier, bound to a specific deployed contract.
func NewGroth16Verifier(address common.Address, backend bind.ContractBackend) (*Groth16Verifier, error) {
	contract, err := bindGroth16Verifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Groth16Verifier{Groth16VerifierCaller: Groth16VerifierCaller{contract: contract}, Groth16VerifierTransactor: Groth16VerifierTransactor{contract: contract}, Groth16VerifierFilterer: Groth16VerifierFilterer{contract: contract}}, nil
}

// NewGroth16VerifierCaller creates a new read-only instance of Groth16Verifier, bound to a specific deployed contract.
func NewGroth16VerifierCaller(address common.Address, caller bind.ContractCaller) (*Groth16VerifierCaller, error) {
	contract, err := bindGroth16Verifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Groth16VerifierCaller{contract: contract}, nil
}

// NewGroth16VerifierTransactor creates a new write-only instance of Groth16Verifier, bound to a specific deployed contract.
func NewGroth16VerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*Groth16VerifierTransactor, error) {
	contract, err := bindGroth16Verifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Groth16VerifierTransactor{contract: contract}, nil
}

// NewGroth16VerifierFilterer creates a new log filterer instance of Groth16Verifier, bound to a specific deployed contract.
func NewGroth16VerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*Groth16VerifierFilterer, error) {
	contract, err := bindGroth16Verifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Groth16VerifierFilterer{contract: contract}, nil
}

// bindGroth16Verifier binds a generic wrapper to an already deployed contract.
func bindGroth16Verifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Groth16VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Groth16Verifier *Groth16VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Groth16Verifier.Contract.Groth16VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Groth16Verifier *Groth16VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Groth16Verifier.Contract.Groth16VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Groth16Verifier *Groth16VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Groth16Verifier.Contract.Groth16VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Groth16Verifier *Groth16VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Groth16Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Groth16Verifier *Groth16VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Groth16Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Groth16Verifier *Groth16VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Groth16Verifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_Groth16Verifier *Groth16VerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	var out []interface{}
	err := _Groth16Verifier.contract.Call(opts, &out, "compressProof", proof, commitments, commitmentPok)

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
func (_Groth16Verifier *Groth16VerifierSession) CompressProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _Groth16Verifier.Contract.CompressProof(&_Groth16Verifier.CallOpts, proof, commitments, commitmentPok)
}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_Groth16Verifier *Groth16VerifierCallerSession) CompressProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _Groth16Verifier.Contract.CompressProof(&_Groth16Verifier.CallOpts, proof, commitments, commitmentPok)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe33974a8.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[4] input) view returns()
func (_Groth16Verifier *Groth16VerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [4]*big.Int) error {
	var out []interface{}
	err := _Groth16Verifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, compressedCommitments, compressedCommitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe33974a8.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[4] input) view returns()
func (_Groth16Verifier *Groth16VerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [4]*big.Int) error {
	return _Groth16Verifier.Contract.VerifyCompressedProof(&_Groth16Verifier.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe33974a8.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[4] input) view returns()
func (_Groth16Verifier *Groth16VerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [4]*big.Int) error {
	return _Groth16Verifier.Contract.VerifyCompressedProof(&_Groth16Verifier.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_Groth16Verifier *Groth16VerifierCaller) VerifyProof(opts *bind.CallOpts, _proof []byte, _input []byte) error {
	var out []interface{}
	err := _Groth16Verifier.contract.Call(opts, &out, "verifyProof", _proof, _input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_Groth16Verifier *Groth16VerifierSession) VerifyProof(_proof []byte, _input []byte) error {
	return _Groth16Verifier.Contract.VerifyProof(&_Groth16Verifier.CallOpts, _proof, _input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_Groth16Verifier *Groth16VerifierCallerSession) VerifyProof(_proof []byte, _input []byte) error {
	return _Groth16Verifier.Contract.VerifyProof(&_Groth16Verifier.CallOpts, _proof, _input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xd148354c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[4] input) view returns()
func (_Groth16Verifier *Groth16VerifierCaller) VerifyProof0(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [4]*big.Int) error {
	var out []interface{}
	err := _Groth16Verifier.contract.Call(opts, &out, "verifyProof0", proof, commitments, commitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xd148354c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[4] input) view returns()
func (_Groth16Verifier *Groth16VerifierSession) VerifyProof0(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [4]*big.Int) error {
	return _Groth16Verifier.Contract.VerifyProof0(&_Groth16Verifier.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xd148354c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[4] input) view returns()
func (_Groth16Verifier *Groth16VerifierCallerSession) VerifyProof0(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [4]*big.Int) error {
	return _Groth16Verifier.Contract.VerifyProof0(&_Groth16Verifier.CallOpts, proof, commitments, commitmentPok, input)
}
