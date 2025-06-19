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
	Bin: "0x6080604052348015600f57600080fd5b50611ee18061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063233ace111461005c578063b1c3a00e1461008f578063b8e72af6146100b1578063d148354c146100c6578063e33974a8146100d9575b600080fd5b6040517f0ab94b77ce41a2271f131485c39a923944a455736cf95527c1475f111a3e08c081526020015b60405180910390f35b6100a261009d36600461198f565b6100ec565b60405161008693929190611a01565b6100c46100bf366004611a90565b610175565b005b6100c46100d4366004611b12565b6101f6565b6100c46100e7366004611b6e565b610698565b6100f46118d4565b6100fc6118f2565b600061011186358760015b6020020135610c00565b835261012f6060870135604088013560a089013560808a0135610cf5565b6020850152604084015261014960c0870135876007610107565b606084015261015b8535866001610107565b825261016a8435856001610107565b905093509350939050565b60008060006101848787610fec565b919450925090503063d148354c84848461019e8a8a61101d565b6040518563ffffffff1660e01b81526004016101bd9493929190611bf1565b60006040518083038186803b1580156101d557600080fd5b505afa1580156101e9573d6000803e3d6000fd5b5050505050505050505050565b6101fe6118f2565b6040805160028082526060808301845292602083019080368337019050509050602081016040848101823750604051600080516020611e8c83398151915290610253908735906020808a013591869101611c50565b6040516020818303038152906040528051906020012060001c6102769190611ca9565b8252604080516000918782377f1c5e636ab199b196437fcad1654ce5a196a1d59a195fecf8b819d8a54cbe5d4260408201527f07d98d3ae028929c68203241a5a731672743e863e09a99962087887775566eba60608201527f25a28d65581a33d35e0ebc98bf5e683ae4a5068c485d89cae7f5d3074cfc98c060808201527e3d5ead19ed732ce622f65adf7eef64ba32f91337a60190665b88b49ae4793d60a082015260408660c08301377f1a6f87b7a35d7994f132fd716464bef29eaa151776e69fd79a6013a70c680dd56101008201527f2dd62bfd0b59232896c83a6e06c67caa4daff7a809c6761438f7d3176cfdc12c6101208201527f2218d54a19369218be0d47f5825a319debb60b343c6623dc4059734af0e37d9c6101408201527f1711042f565505f015c9062409251b910d79713a58e32fb93dec9a7c48bd7f556101608201526020816101808360085afa9051169050806103eb576040516351d49ff760e11b815260040160405180910390fd5b60008061042286868a6002806020026040519081016040528092919082600260200280828437600092019190915250611038915050565b915091506040516101008a82377f1a921844c1b0c9f2038d8fe36364d0eef82c3bd8591c895a7c789fafc00c92426101008201527f0a228ca9daf73089eb5a38ea00e8f764b6a04d0ec220cc507cf98416ecdb7a876101208201527f2e5d46de967b2c8f41c72b297db185b95999ed937b3acd6e2ed3c3f6ebcad5a76101408201527f2b2b2705d6ada3339f99af58806d5d896660a323ddbe365abb8acc3f477a41216101608201527f1267c7e73a62392799c065270d11f2b36863b1549a97483c6548ade85e54049e6101808201527f2ddb7282342319a7ebaebff6ef0bbe7bcf782bb8706cfde54ad5a77e7802e0cb6101a08201527f2a4ca7ec1ca0fecbc020a7df769171436d9213555ec2601da8238ee99e0dc3d16101c08201527f251b98882b9b4a43a56179ccda2744da9691182d9dc59a7a1858747a737192c66101e08201527f1cea811a9c83cc11f773ad30f470b5713539d2cf312d79bb17bf4613319f7bc06102008201527f165b7041d6a2d58a5db878182ca8959954300627607553a803f4e0bbfa89463b61022082015282610240820152816102608201527f0324ac2d60368ca9ffff186218706375aa9743d399844c1c6a2a20b9bf1bc3976102808201527f10a74dd9cc92fe6b496ac74129305dde3e89ac224b3b44dc7403a3f9546408ad6102a08201527f153cdfa35f77d08687673bd735c00f2955be166ad922e2ec277cfcd170a370916102c08201527f1ed2b4b47f62404d7f8c94f3d37231354113432d13a0f0f005dddf0ac8b5d4e86102e08201526020816103008360085afa90511692508261068d57604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b6106a06118f2565b6106a8611910565b6106b061192e565b6106c18660005b602002013561137e565b602084015282526000806106d48761137e565b6040805160028082526060808301845294965092945091906020830190803683370190505090506020810160408881018237508451602080870151604051600080516020611e8c8339815191529361073193909291869101611c50565b6040516020818303038152906040528051906020012060001c6107549190611ca9565b865284518452602085810151858201527f1c5e636ab199b196437fcad1654ce5a196a1d59a195fecf8b819d8a54cbe5d426040808701919091527f07d98d3ae028929c68203241a5a731672743e863e09a99962087887775566eba60608701527f25a28d65581a33d35e0ebc98bf5e683ae4a5068c485d89cae7f5d3074cfc98c060808701527e3d5ead19ed732ce622f65adf7eef64ba32f91337a60190665b88b49ae4793d60a087015260c0860185905260e086018490527f1a6f87b7a35d7994f132fd716464bef29eaa151776e69fd79a6013a70c680dd56101008701527f2dd62bfd0b59232896c83a6e06c67caa4daff7a809c6761438f7d3176cfdc12c6101208701527f2218d54a19369218be0d47f5825a319debb60b343c6623dc4059734af0e37d9c6101408701527f1711042f565505f015c9062409251b910d79713a58e32fb93dec9a7c48bd7f5561016087015251600091816101808860085afa9051169050806108d9576040516351d49ff760e11b815260040160405180910390fd5b505050506000806108f6896000600481106106b7576106b7611bb8565b9092509050600080808061091260408e013560208f0135611424565b9296509094509250905060008061092a8f60036106b7565b9150915060008061093c8e8e8e611038565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f1a921844c1b0c9f2038d8fe36364d0eef82c3bd8591c895a7c789fafc00c92426101008e01527f0a228ca9daf73089eb5a38ea00e8f764b6a04d0ec220cc507cf98416ecdb7a876101208e01527f2e5d46de967b2c8f41c72b297db185b95999ed937b3acd6e2ed3c3f6ebcad5a76101408e01527f2b2b2705d6ada3339f99af58806d5d896660a323ddbe365abb8acc3f477a41216101608e01527f1267c7e73a62392799c065270d11f2b36863b1549a97483c6548ade85e54049e6101808e01527f2ddb7282342319a7ebaebff6ef0bbe7bcf782bb8706cfde54ad5a77e7802e0cb6101a08e01527f2a4ca7ec1ca0fecbc020a7df769171436d9213555ec2601da8238ee99e0dc3d16101c08e01527f251b98882b9b4a43a56179ccda2744da9691182d9dc59a7a1858747a737192c66101e08e01527f1cea811a9c83cc11f773ad30f470b5713539d2cf312d79bb17bf4613319f7bc06102008e01527f165b7041d6a2d58a5db878182ca8959954300627607553a803f4e0bbfa89463b6102208e01526102408d018290526102608d018190527f0324ac2d60368ca9ffff186218706375aa9743d399844c1c6a2a20b9bf1bc3976102808e01527f10a74dd9cc92fe6b496ac74129305dde3e89ac224b3b44dc7403a3f9546408ad6102a08e01527f153cdfa35f77d08687673bd735c00f2955be166ad922e2ec277cfcd170a370916102c08e01527f1ed2b4b47f62404d7f8c94f3d37231354113432d13a0f0f005dddf0ac8b5d4e86102e08e015290925090506000610bb16118f2565b6020816103008f60085afa9150811580610bcd57508051600114155b15610beb57604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b6000600080516020611e6c83398151915283101580610c2d5750600080516020611e6c8339815191528210155b15610c4b57604051631ff3747d60e21b815260040160405180910390fd5b82158015610c57575081155b15610c6457506000610cef565b6000610ca3600080516020611e6c8339815191526003600080516020611e6c83398151915287600080516020611e6c833981519152898a09090861160a565b9050808303610cb8575050600182901b610cef565b610cc18161166e565b8303610cd4575050600182811b17610cef565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b600080600080516020611e6c83398151915286101580610d235750600080516020611e6c8339815191528510155b80610d3c5750600080516020611e6c8339815191528410155b80610d555750600080516020611e6c8339815191528310155b15610d7357604051631ff3747d60e21b815260040160405180910390fd5b82848688171717600003610d8c57506000905080610fe3565b60008080600080516020611e6c833981519152610db86003600080516020611e6c833981519152611ccb565b600080516020611e6c8339815191528a8c090990506000600080516020611e6c8339815191528a600080516020611e6c8339815191528c8d090990506000600080516020611e6c8339815191528a600080516020611e6c8339815191528c8d09099050600080516020611e6c83398151915280600080516020611e6c8339815191528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610eb1600080516020611e6c83398151915280600080516020611e6c8339815191528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750861166e565b9350505050600080610f02600080516020611e6c83398151915280610ed857610ed8611c93565b600080516020611e6c833981519152858609600080516020611e6c8339815191528788090861160a565b9050610f4f600080516020611e6c8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611e6c83398151915284880809611687565b15915050610f5e8383836116d1565b90935091508683148015610f7157508186145b15610f9b5780610f82576000610f85565b60025b60ff1660028a901b176000179450879350610fdf565b610fa48361166e565b87148015610fb95750610fb68261166e565b86145b15610cd45780610fca576000610fcd565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b610ff461194d565b610ffc611910565b611004611910565b61101084860186611d81565b9250925092509250925092565b6110256118d4565b61103182840184611e03565b9392505050565b6000806000600190506040516040810160007f28d2770b23a8307a30fc52c3f55204117de458f220ae928bc9199bf8d13ff4ec83527f14fa3f9fd0b0d0d239e121a1db64e749d24ab9ce70f41f3022d74a1c89ca85c16020840152865182526020870151602083015260408360808560065afa841693507f0416847a71d55eb4e925d521e64638084804afdc0d12cec531b1e3b36a54f24682527f16d5654dfef2af3ccec76550292107e375287b9ac918842b43ca50d08b4b41de602083015288359050806040830152600080516020611e8c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f22802735f44626e8c2bfed5de04c6820ae8be41926245c46fefd95cb69d5278882527f2278503f4e92e194d337e44ebd5066f6e2dd2818853e63903d9e8609b3decb8e602083015260208901359050806040830152600080516020611e8c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f127dfdab346c2ac5369b144bf6bcbe854b530617e5f0a95c2ca17f82bc61264882527f25965e1943aaa1df1c4b99efea91d907ca3d77e8b2ef86414d8da8099bd4699a602083015260408901359050806040830152600080516020611e8c83398151915281108416935060408260608460075afa8416935060408360808560065afa841693507f2312187ec83fdf2f03589aa6cc3fc39be54f672c3bab4415b76e6a143c40a81182527f12f8d7f21cf62ec537fffa349f0306700354fc4795bc02808e3ca28779b999e5602083015260608901359050806040830152600080516020611e8c83398151915281108416935060408260608460075afa8416935060408360808560065afa7f28a0a82250ae577b5b549ace3365a76ede50c4b59072a8f60807901b244c503383527f1559c5db8eced77b84a4ec25a1a1222e3ccacef78054690740f4c1de0ec8bf9e602084015288516040808501829052600080516020611e8c83398151915290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806113755760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b6000808260000361139457506000928392509050565b600183811c925080841614600080516020611e6c83398151915283106113cd57604051631ff3747d60e21b815260040160405180910390fd5b61140a600080516020611e6c8339815191526003600080516020611e6c83398151915286600080516020611e6c833981519152888909090861160a565b9150801561141e5761141b8261166e565b91505b50915091565b600080808085158015611435575084155b1561144b57506000925082915081905080611601565b600286811c94508593506001808816149080881614600080516020611e6c8339815191528610158061148b5750600080516020611e6c8339815191528510155b156114a957604051631ff3747d60e21b815260040160405180910390fd5b6000600080516020611e6c8339815191526114d36003600080516020611e6c833981519152611ccb565b600080516020611e6c833981519152888a090990506000600080516020611e6c83398151915288600080516020611e6c8339815191528a8b090990506000600080516020611e6c83398151915288600080516020611e6c8339815191528a8b09099050600080516020611e6c83398151915280600080516020611e6c8339815191528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50896506115cc600080516020611e6c83398151915280600080516020611e6c8339815191528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750861166e565b95506115d98787866116d1565b909750955084156115fb576115ed8761166e565b96506115f88661166e565b95505b50505050505b92959194509250565b6000611636827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5261180f565b905081600080516020611e6c8339815191528283091461166957604051631ff3747d60e21b815260040160405180910390fd5b919050565b600080516020611e6c8339815191529081900681030690565b6000806116b4837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5261180f565b905082600080516020611e6c833981519152828309149392505050565b60008080611703600080516020611e6c83398151915280878809600080516020611e6c833981519152898a090861160a565b90508315611717576117148161166e565b90505b611762600080516020611e6c8339815191527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea4600080516020611e6c833981519152848a080961160a565b9250600080516020611e6c83398151915261178e600080516020611e6c83398151915260028609611874565b86099150600080516020611e6c8339815191526117bb600080516020611e6c83398151915284850961166e565b600080516020611e6c83398151915285860908861415806117f15750600080516020611e6c833981519152808385096002098514155b1561137557604051631ff3747d60e21b815260040160405180910390fd5b6000806040516020815260208082015260206040820152846060820152836080820152600080516020611e6c83398151915260a082015260208160c08360055afa90519250905080610ced57604051631ff3747d60e21b815260040160405180910390fd5b60006118a0827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4561180f565b9050600080516020611e6c83398151915281830960011461166957604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b806101008101831015610cef57600080fd5b8060408101831015610cef57600080fd5b600080600061018084860312156119a557600080fd5b6119af858561196c565b92506119bf85610100860161197e565b91506119cf85610140860161197e565b90509250925092565b8060005b60048110156119fb5781518452602093840193909101906001016119dc565b50505050565b60c08101611a0f82866119d8565b608082018460005b6001811015611a36578151835260209283019290910190600101611a17565b5050508260a0830152949350505050565b60008083601f840112611a5957600080fd5b50813567ffffffffffffffff811115611a7157600080fd5b602083019150836020828501011115611a8957600080fd5b9250929050565b60008060008060408587031215611aa657600080fd5b843567ffffffffffffffff811115611abd57600080fd5b611ac987828801611a47565b909550935050602085013567ffffffffffffffff811115611ae957600080fd5b611af587828801611a47565b95989497509550505050565b8060808101831015610cef57600080fd5b6000806000806102008587031215611b2957600080fd5b611b33868661196c565b9350611b4386610100870161197e565b9250611b5386610140870161197e565b9150611b63866101808701611b01565b905092959194509250565b6000806000806101408587031215611b8557600080fd5b611b8f8686611b01565b935060a0850186811115611ba257600080fd5b608086019350359150611b638660c08701611b01565b634e487b7160e01b600052603260045260246000fd5b8060005b60028110156119fb578151845260209384019390910190600101611bd2565b6102008101818660005b6008811015611c1a578151835260209283019290910190600101611bfb565b505050611c2b610100830186611bce565b611c39610140830185611bce565b611c476101808301846119d8565b95945050505050565b83815282602082015260006040820183516020850160005b82811015611c86578151845260209384019390910190600101611c68565b5091979650505050505050565b634e487b7160e01b600052601260045260246000fd5b600082611cc657634e487b7160e01b600052601260045260246000fd5b500690565b81810381811115610cef57634e487b7160e01b600052601160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611d2357634e487b7160e01b600052604160045260246000fd5b604052919050565b600082601f830112611d3c57600080fd5b6000611d486040611cec565b9050806040840185811115611d5c57600080fd5b845b81811015611d76578035835260209283019201611d5e565b509195945050505050565b60008060006101808486031215611d9757600080fd5b600085601f860112611da7578081fd5b80610100611db481611cec565b9250829150860187811115611dc857600080fd5b865b81811015611de2578035845260209384019301611dca565b50819550611df08882611d2b565b94505050506119cf856101408601611d2b565b600060808284031215611e1557600080fd5b600083601f840112611e25578081fd5b80611e306080611cec565b90508091506080840185811115611e4657600080fd5b845b81811015611e60578035845260209384019301611e48565b50909594505050505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212200a35767bc4acb608e7460283287b634d9de998146bd6a593f686a854111b26d864736f6c634300081c0033",
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
