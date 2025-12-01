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
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"},{\"internalType\":\"uint256[10]\",\"name\":\"input\",\"type\":\"uint256[10]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[10]\",\"name\":\"input\",\"type\":\"uint256[10]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506122838061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063233ace1114610059578063b1c3a00e1461008c578063b8e72af6146100ae578063dc37819a146100c3578063df078bef146100d6575b5f5ffd5b6040517fbecc241f62caba8a8f143500630e1f8e43cdbbb7eb5692011ea8db114faa1c2581526020015b60405180910390f35b61009f61009a366004611c7e565b6100e9565b60405161008393929190611cc4565b6100c16100bc366004611d6a565b610171565b005b6100c16100d1366004611de7565b6101f3565b6100c16100e4366004611e3c565b61075f565b6100f1611ba6565b6100f9611bc4565b5f61010d86358760015b6020020135610c0b565b835261012b6060870135604088013560a089013560808a0135610cf8565b6020850152604084015261014560c0870135876007610103565b60608401526101578535866001610103565b82526101668435856001610103565b905093509350939050565b5f5f5f61017e8787610fd1565b9250925092505f61018f8686611002565b505060405163df078bef60e01b8152909150309063df078bef906101bd908790879087908790600401611ec6565b5f6040518083038186803b1580156101d3575f5ffd5b505afa1580156101e5573d5f5f3e3d5ffd5b505050505050505050505050565b6101fb611bc4565b610203611be2565b61020b611c00565b61021b865f5b6020020135611019565b602084015282525f8061022d87611019565b60408051600680825260e08201909252929450909250606091906020820160c0803683370190505090506020810160408881018237608060c0890160408301375084516020808701516040515f51602061222e5f395f51905f529361029793909291869101611f54565b604051602081830303815290604052805190602001205f1c6102b99190611fa9565b865284518452602085810151858201527f2bfec23d1e47d8380e418059bf140893685ac47d38cc904a1d37763396e3ad686040808701919091527f0a04458ec21eee30335c32a5c62ec2c6ae221c26b6a92303521f0eefb380325d60608701527f1458086ba2f4c201e25bd4776cdca7f3032d752e6e77e4009501c0dd9326c45060808701527f025503912ba4c8f8e8ddebecb96727ded29322ebf7fcfdafdd49494777551b2b60a087015260c0860185905260e086018490527f22312b603461adc5477f2762bc30364deeb50f35f1ca6556100b11f1199fd9896101008701527f2f8d4b81f84b71ad28090b486e8189b98a91c216c20af897b79363fe0a1ead9d6101208701527f11213ae54cfe34f4bcbcf600bd1c41e87660078080830d4a168f4feefd07c2266101408701527f205158fe75700956cac165bd9227d8c1dba0250e113fce3fbf1249979e50927a610160870152515f91816101808860085afa90511690508061043e576040516351d49ff760e11b815260040160405180910390fd5b505050505f5f610459895f6004811061021157610211611e8a565b90925090505f80808061047460408e013560208f01356110b8565b929650909450925090505f8061048b8f6003610211565b915091505f5f61049c8e8e8e61128c565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f2bfb03e649e62a6f85d1281799dcbba8ee2c9f0da8acfae0a42aeb6a74149ee46101008e01527f0ce955c97ac319f9c0a71638b6ab1335b0b5457014df1f92ea483736cf50303a6101208e01527f2a9bb33a933d524e330525e3c694718d51ab6908ac85ea3b40e67e77f4901c1e6101408e01527f1b91e4881ff9f4255fd403c89969619794c3a610b660ab1b6660a7da7a79081c6101608e01527f1e7595921ea821ff6e14d997bcbc7afaaff2baec7e5820939222672f7164df606101808e01527f25981316d4f87a82de4b20300794483c2f46670259b60343a2f67373840cf47e6101a08e01527f124a7e82338fa7a644db147f4fe9f7b38c95b1e423b06b3569cde5bfdf6eb0a56101c08e01527f09deb8d12db3771be215f60b5958683e5d5620b908a3a4fbfefd3587a1a5b1546101e08e01527f176dfcdc004b315a1600577bc3c0c05b77f7c161bb13181588a68805ebb354146102008e01527f0ebd8b4c76a653629c9834d6608691044e14d55e7ba598ae1686ee7a5b2dd26f6102208e01526102408d018290526102608d018190527f1e677e20b4ba82bb161219fbe2d1639f7d5b75e761974db9648364ba33a4561e6102808e01527f05c54f897b66086444d0e4eed709665a1cb5ef72227375c9a4003a2096587a6b6102a08e01527f037b00897e6219192a453286e2f288edca93df14e76eeb750c85ea2610e8cebc6102c08e01527f0fe63eb9d684b99cd5fd3a1387efe402737bba926e710ef1255133102d4923f46102e08e015290925090505f610710611bc4565b6020816103008f60085afa915081158061072c57508051600114155b1561074a57604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b610767611bc4565b60408051600680825260e082019092526060916020820160c0803683370190505090506020810160408481018237608060c085016040830137506040515f51602061222e5f395f51905f52906107c9908735906020808a013591869101611f54565b604051602081830303815290604052805190602001205f1c6107eb9190611fa9565b8252604080515f918782377f2bfec23d1e47d8380e418059bf140893685ac47d38cc904a1d37763396e3ad6860408201527f0a04458ec21eee30335c32a5c62ec2c6ae221c26b6a92303521f0eefb380325d60608201527f1458086ba2f4c201e25bd4776cdca7f3032d752e6e77e4009501c0dd9326c45060808201527f025503912ba4c8f8e8ddebecb96727ded29322ebf7fcfdafdd49494777551b2b60a082015260408660c08301377f22312b603461adc5477f2762bc30364deeb50f35f1ca6556100b11f1199fd9896101008201527f2f8d4b81f84b71ad28090b486e8189b98a91c216c20af897b79363fe0a1ead9d6101208201527f11213ae54cfe34f4bcbcf600bd1c41e87660078080830d4a168f4feefd07c2266101408201527f205158fe75700956cac165bd9227d8c1dba0250e113fce3fbf1249979e50927a6101608201526020816101808360085afa905116905080610960576040516351d49ff760e11b815260040160405180910390fd5b5f5f61099586868a60028060200260405190810160405280929190826002602002808284375f9201919091525061128c915050565b915091506040516101008a82377f2bfb03e649e62a6f85d1281799dcbba8ee2c9f0da8acfae0a42aeb6a74149ee46101008201527f0ce955c97ac319f9c0a71638b6ab1335b0b5457014df1f92ea483736cf50303a6101208201527f2a9bb33a933d524e330525e3c694718d51ab6908ac85ea3b40e67e77f4901c1e6101408201527f1b91e4881ff9f4255fd403c89969619794c3a610b660ab1b6660a7da7a79081c6101608201527f1e7595921ea821ff6e14d997bcbc7afaaff2baec7e5820939222672f7164df606101808201527f25981316d4f87a82de4b20300794483c2f46670259b60343a2f67373840cf47e6101a08201527f124a7e82338fa7a644db147f4fe9f7b38c95b1e423b06b3569cde5bfdf6eb0a56101c08201527f09deb8d12db3771be215f60b5958683e5d5620b908a3a4fbfefd3587a1a5b1546101e08201527f176dfcdc004b315a1600577bc3c0c05b77f7c161bb13181588a68805ebb354146102008201527f0ebd8b4c76a653629c9834d6608691044e14d55e7ba598ae1686ee7a5b2dd26f61022082015282610240820152816102608201527f1e677e20b4ba82bb161219fbe2d1639f7d5b75e761974db9648364ba33a4561e6102808201527f05c54f897b66086444d0e4eed709665a1cb5ef72227375c9a4003a2096587a6b6102a08201527f037b00897e6219192a453286e2f288edca93df14e76eeb750c85ea2610e8cebc6102c08201527f0fe63eb9d684b99cd5fd3a1387efe402737bba926e710ef1255133102d4923f46102e08201526020816103008360085afa905116925082610c0057604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b5f5f51602061220e5f395f51905f5283101580610c3557505f51602061220e5f395f51905f528210155b15610c5357604051631ff3747d60e21b815260040160405180910390fd5b82158015610c5f575081155b15610c6b57505f610cf2565b5f610ca65f51602061220e5f395f51905f5260035f51602061220e5f395f51905f52875f51602061220e5f395f51905f52898a0909086118f0565b9050808303610cbb575050600182901b610cf2565b610cc481611952565b8303610cd7575050600182811b17610cf2565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b5f5f5f51602061220e5f395f51905f5286101580610d2357505f51602061220e5f395f51905f528510155b80610d3b57505f51602061220e5f395f51905f528410155b80610d5357505f51602061220e5f395f51905f528310155b15610d7157604051631ff3747d60e21b815260040160405180910390fd5b828486881717175f03610d8857505f905080610fc8565b5f80805f51602061220e5f395f51905f52610db160035f51602061220e5f395f51905f52611fc8565b5f51602061220e5f395f51905f528a8c090990505f5f51602061220e5f395f51905f528a5f51602061220e5f395f51905f528c8d090990505f5f51602061220e5f395f51905f528a5f51602061220e5f395f51905f528c8d090990505f51602061220e5f395f51905f52805f51602061220e5f395f51905f528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610e9f5f51602061220e5f395f51905f52805f51602061220e5f395f51905f528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611952565b93505050505f5f610eec5f51602061220e5f395f51905f5280610ec457610ec4611f95565b5f51602061220e5f395f51905f528586095f51602061220e5f395f51905f52878809086118f0565b9050610f375f51602061220e5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f51602061220e5f395f51905f528488080961196a565b15915050610f468383836119b2565b90935091508683148015610f5957508186145b15610f815780610f69575f610f6c565b60025b60ff1660028a901b175f179450879350610fc4565b610f8a83611952565b87148015610f9f5750610f9c82611952565b86145b15610cd75780610faf575f610fb2565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b610fd9611c1f565b610fe1611be2565b610fe9611be2565b610ff58486018661206a565b9250925092509250925092565b61100a611c3e565b606080610ff584860186612155565b5f5f825f0361102c57505f928392509050565b600183811c9250808416145f51602061220e5f395f51905f52831061106457604051631ff3747d60e21b815260040160405180910390fd5b61109e5f51602061220e5f395f51905f5260035f51602061220e5f395f51905f52865f51602061220e5f395f51905f5288890909086118f0565b915080156110b2576110af82611952565b91505b50915091565b5f808080851580156110c8575084155b156110dd57505f925082915081905080611283565b600286811c945085935060018088161490808816145f51602061220e5f395f51905f528610158061111b57505f51602061220e5f395f51905f528510155b1561113957604051631ff3747d60e21b815260040160405180910390fd5b5f5f51602061220e5f395f51905f5261116060035f51602061220e5f395f51905f52611fc8565b5f51602061220e5f395f51905f52888a090990505f5f51602061220e5f395f51905f52885f51602061220e5f395f51905f528a8b090990505f5f51602061220e5f395f51905f52885f51602061220e5f395f51905f528a8b090990505f51602061220e5f395f51905f52805f51602061220e5f395f51905f528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508965061124e5f51602061220e5f395f51905f52805f51602061220e5f395f51905f528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611952565b955061125b8787866119b2565b9097509550841561127d5761126f87611952565b965061127a86611952565b95505b50505050505b92959194509250565b5f5f5f60019050604051604081015f7f2e92bd94db81c8538fb5f73e859cbd086e30f3b7915a2f356c62858d7a04c1b683527f121126613d0cfd74cf50fbc262f96c4ea8a665c2715f398b45a2b1a3391f851f6020840152865182526020870151602083015260408360808560065afa841693507f0aa6b248a1fbc118d5af6a5876c3726361e2ebea339974bd4d58582a40832b0982527f0c5c662c99fb6cba5518a3258515a9aa5b6956236d618f353f15e799188bafd16020830152883590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f181f15c9f749b2a1783813398981a8e85ee63d372f33a7653945ca17fa966bbc82527f1e2c5a6ce19ecf5de7f6a155fff1815bb009eff8b9aed23855e808c20c60c52d6020830152602089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f234809d2bffe20de78d1160bcd60ad80aa0514cc9dc3a9c2e041b60ceb7e7bdf82527f03a614342d1af93aaaa5cdd6255fe33c823ddb2493530545e0460ce910b400df6020830152604089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f18bbf6e26bd503aed746ade7c5821a1b4a395a2a9915987b04e3504134575a2282527f03ae4b74da0596d23922b3e0ccda7789e17abb0b53a8dc7dee5658a874777ecb6020830152606089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f2365adb235068aab327328bb4c9a6067b4738fbb5ae09d1b439e447b48884e6682527f0678c2fef6e6a27c7d6e63a26ae374d0d06aa86497cc7f7d03efb23cd401feec6020830152608089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f23db6820591bca1b95eb8fff37f2174ceba413422e322230d06344eb17f3023e82527f2fa8a4c77c7c9e40bf7f3872c1bd481357b8df8066e1a9f4bcf515dde8c3c1e8602083015260a089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f1cdad219ecb8bb0d2ea328ac67ac86fdea7e19e6bf801c908590841e4703e55882527f1c3aada55fb30351a71131db9f7e63a76fa1f638187795556aaa6d9f13cdf619602083015260c089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f049d41f08e70794c7b206d49013366c90ba1a8b1a84c3356373342ee8feb4f5982527f09458b6d4c14ef010179fdfef352b95064cf7f6114bf1e0f79c8b4b8aa2e5d7b602083015260e089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f084e30b73d50a251e5489422da44a809a2a0c070c9738e3ae6c8c40d9a44362c82527f039ac736d856789b9f7d0243a49673175fd47fe4efc3b67c6ca2291e4d45c0ad602083015261010089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f06d11a4d902d7d638206fb953b144b911d50ddfe296af682284513a894c7d08182527f283cdf56f5abe38e79352f4aea1e1a25f14b18a6707720e9f40a46e572f045e1602083015261012089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa7f09afda261e8493b3ef82bd1ec1736842b7641d63d9acaa21e5e6f0bea866ab9e83527f17e1f4dd4f5831313cbbfe3d919963d6938fa8bfb199a95ea003e028b8a9b4bf6020840152885160408085018290525f51602061222e5f395f51905f5290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806118e75760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b5f61191b827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611ae5565b9050815f51602061220e5f395f51905f528283091461194d57604051631ff3747d60e21b815260040160405180910390fd5b919050565b5f51602061220e5f395f51905f529081900681030690565b5f5f611996837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611ae5565b9050825f51602061220e5f395f51905f52828309149392505050565b5f80806119e15f51602061220e5f395f51905f52808788095f51602061220e5f395f51905f52898a09086118f0565b905083156119f5576119f281611952565b90505b611a3e5f51602061220e5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f51602061220e5f395f51905f52848a08096118f0565b92505f51602061220e5f395f51905f52611a685f51602061220e5f395f51905f5260028609611b48565b860991505f51602061220e5f395f51905f52611a935f51602061220e5f395f51905f52848509611952565b5f51602061220e5f395f51905f528586090886141580611ac757505f51602061220e5f395f51905f52808385096002098514155b156118e757604051631ff3747d60e21b815260040160405180910390fd5b5f5f60405160208152602080820152602060408201528460608201528360808201525f51602061220e5f395f51905f5260a082015260208160c08360055afa90519250905080610cf057604051631ff3747d60e21b815260040160405180910390fd5b5f611b73827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611ae5565b90505f51602061220e5f395f51905f5281830960011461194d57604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b604051806101400160405280600a906020820280368337509192915050565b806101008101831015610cf2575f5ffd5b8060408101831015610cf2575f5ffd5b5f5f5f6101808486031215611c91575f5ffd5b611c9b8585611c5d565b9250611cab856101008601611c6e565b9150611cbb856101408601611c6e565b90509250925092565b60c0810181855f5b6004811015611ceb578151835260209283019290910190600101611ccc565b50505060808201845f5b6001811015611d14578151835260209283019290910190600101611cf5565b5050508260a0830152949350505050565b5f5f83601f840112611d35575f5ffd5b50813567ffffffffffffffff811115611d4c575f5ffd5b602083019150836020828501011115611d63575f5ffd5b9250929050565b5f5f5f5f60408587031215611d7d575f5ffd5b843567ffffffffffffffff811115611d93575f5ffd5b611d9f87828801611d25565b909550935050602085013567ffffffffffffffff811115611dbe575f5ffd5b611dca87828801611d25565b95989497509550505050565b806101408101831015610cf2575f5ffd5b5f5f5f5f6102008587031215611dfb575f5ffd5b6080850186811115611e0b575f5ffd5b85945060a0860187811115611e1e575f5ffd5b909350359150611e318660c08701611dd6565b905092959194509250565b5f5f5f5f6102c08587031215611e50575f5ffd5b611e5a8686611c5d565b9350611e6a866101008701611c6e565b9250611e7a866101408701611c6e565b9150611e31866101808701611dd6565b634e487b7160e01b5f52603260045260245ffd5b805f5b6002811015611ec0578151845260209384019390910190600101611ea1565b50505050565b6102c0810181865f5b6008811015611eee578151835260209283019290910190600101611ecf565b505050611eff610100830186611e9e565b611f0d610140830185611e9e565b6101808201835f5b600a811015611f34578151835260209283019290910190600101611f15565b50505095945050505050565b634e487b7160e01b5f52604160045260245ffd5b8381528260208201525f604082018351602085015f5b82811015611f88578151845260209384019390910190600101611f6a565b5091979650505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f82611fc357634e487b7160e01b5f52601260045260245ffd5b500690565b81810381811115610cf257634e487b7160e01b5f52601160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561201057612010611f40565b604052919050565b5f82601f830112612027575f5ffd5b5f6120326040611fe7565b9050806040840185811115612045575f5ffd5b845b8181101561205f578035835260209283019201612047565b509195945050505050565b5f5f5f610180848603121561207d575f5ffd5b5f85601f86011261208c575f5ffd5b505f8061010061209b81611fe7565b92508291508601878111156120ae575f5ffd5b865b818110156120c85780358452602093840193016120b0565b508195506120d68882612018565b9450505050611cbb856101408601612018565b5f82601f8301126120f8575f5ffd5b813567ffffffffffffffff81111561211257612112611f40565b612125601f8201601f1916602001611fe7565b818152846020838601011115612139575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f6101808486031215612168575f5ffd5b5f85601f860112612177575f5ffd5b505f8061014061218681611fe7565b9250829150860187811115612199575f5ffd5b865b818110156121b357803584526020938401930161219b565b5090945035905067ffffffffffffffff8111156121ce575f5ffd5b6121da868287016120e9565b92505061016084013567ffffffffffffffff8111156121f7575f5ffd5b612203868287016120e9565b915050925092509256fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212203db8645a59c9f9a9b5b914789d54bb9e3f3495fd5d30e56d2879cb73b158a98b64736f6c634300081c0033",
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xdc37819a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[10] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [10]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, compressedCommitments, compressedCommitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xdc37819a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[10] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [10]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyCompressedProof(&_StateTransitionVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xdc37819a.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[10] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [10]*big.Int) error {
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

// VerifyProof0 is a free data retrieval call binding the contract method 0xdf078bef.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[10] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyProof0(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [10]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyProof0", proof, commitments, commitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xdf078bef.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[10] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyProof0(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [10]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof0(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xdf078bef.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[10] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyProof0(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [10]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof0(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}
