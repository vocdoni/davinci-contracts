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
	Bin: "0x6080604052348015600e575f5ffd5b506122838061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063233ace1114610059578063b1c3a00e1461008c578063b8e72af6146100ae578063dc37819a146100c3578063df078bef146100d6575b5f5ffd5b6040517f037d0059bf0a6c47974904c57de1b2559183125bf0c26ea94efb503c164bf8da81526020015b60405180910390f35b61009f61009a366004611c7e565b6100e9565b60405161008393929190611cc4565b6100c16100bc366004611d6a565b610171565b005b6100c16100d1366004611de7565b6101f3565b6100c16100e4366004611e3c565b61075f565b6100f1611ba6565b6100f9611bc4565b5f61010d86358760015b6020020135610c0b565b835261012b6060870135604088013560a089013560808a0135610cf8565b6020850152604084015261014560c0870135876007610103565b60608401526101578535866001610103565b82526101668435856001610103565b905093509350939050565b5f5f5f61017e8787610fd1565b9250925092505f61018f8686611002565b505060405163df078bef60e01b8152909150309063df078bef906101bd908790879087908790600401611ec6565b5f6040518083038186803b1580156101d3575f5ffd5b505afa1580156101e5573d5f5f3e3d5ffd5b505050505050505050505050565b6101fb611bc4565b610203611be2565b61020b611c00565b61021b865f5b6020020135611019565b602084015282525f8061022d87611019565b60408051600580825260c08201909252929450909250606091906020820160a080368337019050509050602080820190604089018237608060c0890160208301375084516020808701516040515f51602061222e5f395f51905f529361029893909291869101611f54565b604051602081830303815290604052805190602001205f1c6102ba9190611fa9565b865284518452602085810151858201527f08480e99474689ed895e0f83b2c28bc234b7c86bb7c5cf322c2cc29362fb48ad6040808701919091527f06170922dfcf55190a392089203c49b03a237f16c28b822a27acba4f2af24eb460608701527f0d5c11ef1d50d15154301c075f3b7893e2fb98d3ed90b15ac96f51552e5c36ac60808701527f071e1f97309edb22456d11f96e2c6718c7d525bfd35fb13125d854bf927a4a4e60a087015260c0860185905260e086018490527f131443a740e68b693e0b637c8cf26808f65425688ffc3ff9a10abdd47a3bfd906101008701527f1a78fae6054a26f8a53f21ec8c056d4d7eef33931f728103e40a766e81256c526101208701527f0eccc78a9e3b7fcc9d605d6e8df148ee93f85550c25dded759a6d35503cae1eb6101408701527f1f29f1a675c809bc3c705bedff15fbb9bddc5aad42d64ae858ab8c300dc96ff0610160870152515f91816101808860085afa90511690508061043f576040516351d49ff760e11b815260040160405180910390fd5b505050505f5f61045a895f6004811061021157610211611e8a565b90925090505f80808061047560408e013560208f01356110b8565b929650909450925090505f8061048c8f6003610211565b915091505f5f61049d8e8e8e61128c565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f1ad227eb2f384e9075cc8ed2018117556ca911e50fb867ba5958f1838833d4fd6101008e01527f25a7fed94b8b55821dcec4da0f0f38217b4307eb275f054de637d8736b96beb46101208e01527f100cfe015e2bff9e7592d4e475dbdc4b72b49bfb4c7cd48bf5abb2d15a371fdd6101408e01527f23043b3c22a4f942c457bc902c3c84f429346d9aeda87031948091a363d38db76101608e01527f264858f744d06905279fd6899b94ffb20bc3650dbf31b28f5e9b3d07c1f591a16101808e01527f2d3f129bf66575378f39effb18d0353a3582ccb27679f27dcbe6ce31fb56bf366101a08e01527f225f334f71fda2c050fabf3fe222568f134e00efdb6fec322f799cdc646266c66101c08e01527f05b90cfdc2e85b5777b23d7e37cac69ef33e85055021e423fd838574a87485646101e08e01527f07cac0db8e4604ec86e5b568f91dcf4e6ae645f437c548e07af2ab319aa375de6102008e01527f139553512b424d5db0fb8a6a81c48da0175433b34dbfb7909d20c89f1761d02a6102208e01526102408d018290526102608d018190527f08e3127cddee1e6633c965f4a96a2bf9604fcc0c6e394940688c66baad53e7b96102808e01527e796bf3d2b960e7699e4d44a325ce5405527b23ded7926746817113f9b6b3846102a08e01527f23a3f24661e90135bfda3b65a2e8a1b7bc4b3be9229dd3db008215c6a2ff3e376102c08e01527f02ce17e16de452fdf75d8a97c065075bb5a269cc1addaf930311c0665610ec1c6102e08e015290925090505f610710611bc4565b6020816103008f60085afa915081158061072c57508051600114155b1561074a57604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b610767611bc4565b60408051600580825260c082019092526060916020820160a080368337019050509050602080820190604085018237608060c085016020830137506040515f51602061222e5f395f51905f52906107ca908735906020808a013591869101611f54565b604051602081830303815290604052805190602001205f1c6107ec9190611fa9565b8252604080515f918782377f08480e99474689ed895e0f83b2c28bc234b7c86bb7c5cf322c2cc29362fb48ad60408201527f06170922dfcf55190a392089203c49b03a237f16c28b822a27acba4f2af24eb460608201527f0d5c11ef1d50d15154301c075f3b7893e2fb98d3ed90b15ac96f51552e5c36ac60808201527f071e1f97309edb22456d11f96e2c6718c7d525bfd35fb13125d854bf927a4a4e60a082015260408660c08301377f131443a740e68b693e0b637c8cf26808f65425688ffc3ff9a10abdd47a3bfd906101008201527f1a78fae6054a26f8a53f21ec8c056d4d7eef33931f728103e40a766e81256c526101208201527f0eccc78a9e3b7fcc9d605d6e8df148ee93f85550c25dded759a6d35503cae1eb6101408201527f1f29f1a675c809bc3c705bedff15fbb9bddc5aad42d64ae858ab8c300dc96ff06101608201526020816101808360085afa905116905080610961576040516351d49ff760e11b815260040160405180910390fd5b5f5f61099686868a60028060200260405190810160405280929190826002602002808284375f9201919091525061128c915050565b915091506040516101008a82377f1ad227eb2f384e9075cc8ed2018117556ca911e50fb867ba5958f1838833d4fd6101008201527f25a7fed94b8b55821dcec4da0f0f38217b4307eb275f054de637d8736b96beb46101208201527f100cfe015e2bff9e7592d4e475dbdc4b72b49bfb4c7cd48bf5abb2d15a371fdd6101408201527f23043b3c22a4f942c457bc902c3c84f429346d9aeda87031948091a363d38db76101608201527f264858f744d06905279fd6899b94ffb20bc3650dbf31b28f5e9b3d07c1f591a16101808201527f2d3f129bf66575378f39effb18d0353a3582ccb27679f27dcbe6ce31fb56bf366101a08201527f225f334f71fda2c050fabf3fe222568f134e00efdb6fec322f799cdc646266c66101c08201527f05b90cfdc2e85b5777b23d7e37cac69ef33e85055021e423fd838574a87485646101e08201527f07cac0db8e4604ec86e5b568f91dcf4e6ae645f437c548e07af2ab319aa375de6102008201527f139553512b424d5db0fb8a6a81c48da0175433b34dbfb7909d20c89f1761d02a61022082015282610240820152816102608201527f08e3127cddee1e6633c965f4a96a2bf9604fcc0c6e394940688c66baad53e7b96102808201527e796bf3d2b960e7699e4d44a325ce5405527b23ded7926746817113f9b6b3846102a08201527f23a3f24661e90135bfda3b65a2e8a1b7bc4b3be9229dd3db008215c6a2ff3e376102c08201527f02ce17e16de452fdf75d8a97c065075bb5a269cc1addaf930311c0665610ec1c6102e08201526020816103008360085afa905116925082610c0057604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b5f5f51602061220e5f395f51905f5283101580610c3557505f51602061220e5f395f51905f528210155b15610c5357604051631ff3747d60e21b815260040160405180910390fd5b82158015610c5f575081155b15610c6b57505f610cf2565b5f610ca65f51602061220e5f395f51905f5260035f51602061220e5f395f51905f52875f51602061220e5f395f51905f52898a0909086118f0565b9050808303610cbb575050600182901b610cf2565b610cc481611952565b8303610cd7575050600182811b17610cf2565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b5f5f5f51602061220e5f395f51905f5286101580610d2357505f51602061220e5f395f51905f528510155b80610d3b57505f51602061220e5f395f51905f528410155b80610d5357505f51602061220e5f395f51905f528310155b15610d7157604051631ff3747d60e21b815260040160405180910390fd5b828486881717175f03610d8857505f905080610fc8565b5f80805f51602061220e5f395f51905f52610db160035f51602061220e5f395f51905f52611fc8565b5f51602061220e5f395f51905f528a8c090990505f5f51602061220e5f395f51905f528a5f51602061220e5f395f51905f528c8d090990505f5f51602061220e5f395f51905f528a5f51602061220e5f395f51905f528c8d090990505f51602061220e5f395f51905f52805f51602061220e5f395f51905f528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610e9f5f51602061220e5f395f51905f52805f51602061220e5f395f51905f528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611952565b93505050505f5f610eec5f51602061220e5f395f51905f5280610ec457610ec4611f95565b5f51602061220e5f395f51905f528586095f51602061220e5f395f51905f52878809086118f0565b9050610f375f51602061220e5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f51602061220e5f395f51905f528488080961196a565b15915050610f468383836119b2565b90935091508683148015610f5957508186145b15610f815780610f69575f610f6c565b60025b60ff1660028a901b175f179450879350610fc4565b610f8a83611952565b87148015610f9f5750610f9c82611952565b86145b15610cd75780610faf575f610fb2565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b610fd9611c1f565b610fe1611be2565b610fe9611be2565b610ff58486018661206a565b9250925092509250925092565b61100a611c3e565b606080610ff584860186612155565b5f5f825f0361102c57505f928392509050565b600183811c9250808416145f51602061220e5f395f51905f52831061106457604051631ff3747d60e21b815260040160405180910390fd5b61109e5f51602061220e5f395f51905f5260035f51602061220e5f395f51905f52865f51602061220e5f395f51905f5288890909086118f0565b915080156110b2576110af82611952565b91505b50915091565b5f808080851580156110c8575084155b156110dd57505f925082915081905080611283565b600286811c945085935060018088161490808816145f51602061220e5f395f51905f528610158061111b57505f51602061220e5f395f51905f528510155b1561113957604051631ff3747d60e21b815260040160405180910390fd5b5f5f51602061220e5f395f51905f5261116060035f51602061220e5f395f51905f52611fc8565b5f51602061220e5f395f51905f52888a090990505f5f51602061220e5f395f51905f52885f51602061220e5f395f51905f528a8b090990505f5f51602061220e5f395f51905f52885f51602061220e5f395f51905f528a8b090990505f51602061220e5f395f51905f52805f51602061220e5f395f51905f528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508965061124e5f51602061220e5f395f51905f52805f51602061220e5f395f51905f528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611952565b955061125b8787866119b2565b9097509550841561127d5761126f87611952565b965061127a86611952565b95505b50505050505b92959194509250565b5f5f5f60019050604051604081015f7f0abda91b93553fab611cf205f43c4929e13a29664c6f6b1cb87887cf7a97110983527f182dea6934668a65e4acb3f18ea50e200b4ca22d442481345ffb3b278390b36d6020840152865182526020870151602083015260408360808560065afa841693507f18cc51b23cf3b98e5c65b2a986b2cb28514cfeda4b10b0f2a65fcbf9ae2043ff82527f091a495b3fde8e4359617a338fe3d16a4df8247a2fcca7ea63b115441228abd36020830152883590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f2280b99a402eefd18fbec5c574ecc59541426f9d871d5d8645a3ccd8471ceb5982527f04bfcb74ee48d715ef2c565d838a694697ac445f10d92a5f43409ebdd33e081e6020830152602089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f18702af999fdfb084ab374ed64b8ee09935b2fe45a31a2beb44ed4d780bc63ac82527f1f337daff338e14a7763cd02b766d27f7d4baa07922c5a8aafe91d22b111761a6020830152604089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f172545bf99daf78b18f069c42b0beded5407c5905ce473a4a06ba3aab21a1e8582527f2fedd2c2905a78068d502c39b494c19c94b2f8b1c67f566116cce2166eb0fdc76020830152606089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f09888c24abfb0f110a79c44e64c06e2d250d1b887585f3e90aa6aa4ff17072c982527f18d69bd6bd7b75a7fb04db5180c7a34529c87336370bdb323353a6ed31aa59b86020830152608089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f12981dc688e37c182b36ce1bae3710f2c10fb2d103847bb56d6559b7de44fe7682527f194a3701487e0f227a8ada3e61985569d92de569146aeaac2269566a5d2d987d602083015260a089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f01443c797a7fdf0739def6a1e97af8a4e069da0b27d9adf17074aa20de5fd73d82527f1a2105196f3a41f1cb3a476950c3a1e8dfec97cf35fbd0be7a10c66dce36d485602083015260c089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f044ebcc29152c4e640d8314673c87b65890f7ad31f8b466b2fee5e8987492f2e82527f089349e5de9d84107a7737206a19199a17fbee86f4f2c762112b45fc195dada7602083015260e089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f189111ce67d6a0884adf5102e1a952a33f299ef4bffd6989f4f1f6f5b2d8820482527f073c3a024edffdbd09e1441588f9a270342b4bfad0439e4916d9cb31963ccf36602083015261010089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f22666bc12092e505973668a6291c13d34527a4a1cc97ad746f66d29af69be51982527f12c46ffd30b663df813cdda4b434160156a5eb86961560579df481cfb42a6d22602083015261012089013590508060408301525f51602061222e5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa7f07ae2b84cd2ad8f4cde91d68ebf1f07bf362b7adbe0f071093927096c6f73d6b83527f1bf427bc1f30cfb006f16a548bd7e0b7644a8d1ed9f095aaf4a74c70ef0ee1e56020840152885160408085018290525f51602061222e5f395f51905f5290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806118e75760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b5f61191b827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611ae5565b9050815f51602061220e5f395f51905f528283091461194d57604051631ff3747d60e21b815260040160405180910390fd5b919050565b5f51602061220e5f395f51905f529081900681030690565b5f5f611996837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611ae5565b9050825f51602061220e5f395f51905f52828309149392505050565b5f80806119e15f51602061220e5f395f51905f52808788095f51602061220e5f395f51905f52898a09086118f0565b905083156119f5576119f281611952565b90505b611a3e5f51602061220e5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f51602061220e5f395f51905f52848a08096118f0565b92505f51602061220e5f395f51905f52611a685f51602061220e5f395f51905f5260028609611b48565b860991505f51602061220e5f395f51905f52611a935f51602061220e5f395f51905f52848509611952565b5f51602061220e5f395f51905f528586090886141580611ac757505f51602061220e5f395f51905f52808385096002098514155b156118e757604051631ff3747d60e21b815260040160405180910390fd5b5f5f60405160208152602080820152602060408201528460608201528360808201525f51602061220e5f395f51905f5260a082015260208160c08360055afa90519250905080610cf057604051631ff3747d60e21b815260040160405180910390fd5b5f611b73827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611ae5565b90505f51602061220e5f395f51905f5281830960011461194d57604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b604051806101400160405280600a906020820280368337509192915050565b806101008101831015610cf2575f5ffd5b8060408101831015610cf2575f5ffd5b5f5f5f6101808486031215611c91575f5ffd5b611c9b8585611c5d565b9250611cab856101008601611c6e565b9150611cbb856101408601611c6e565b90509250925092565b60c0810181855f5b6004811015611ceb578151835260209283019290910190600101611ccc565b50505060808201845f5b6001811015611d14578151835260209283019290910190600101611cf5565b5050508260a0830152949350505050565b5f5f83601f840112611d35575f5ffd5b50813567ffffffffffffffff811115611d4c575f5ffd5b602083019150836020828501011115611d63575f5ffd5b9250929050565b5f5f5f5f60408587031215611d7d575f5ffd5b843567ffffffffffffffff811115611d93575f5ffd5b611d9f87828801611d25565b909550935050602085013567ffffffffffffffff811115611dbe575f5ffd5b611dca87828801611d25565b95989497509550505050565b806101408101831015610cf2575f5ffd5b5f5f5f5f6102008587031215611dfb575f5ffd5b6080850186811115611e0b575f5ffd5b85945060a0860187811115611e1e575f5ffd5b909350359150611e318660c08701611dd6565b905092959194509250565b5f5f5f5f6102c08587031215611e50575f5ffd5b611e5a8686611c5d565b9350611e6a866101008701611c6e565b9250611e7a866101408701611c6e565b9150611e31866101808701611dd6565b634e487b7160e01b5f52603260045260245ffd5b805f5b6002811015611ec0578151845260209384019390910190600101611ea1565b50505050565b6102c0810181865f5b6008811015611eee578151835260209283019290910190600101611ecf565b505050611eff610100830186611e9e565b611f0d610140830185611e9e565b6101808201835f5b600a811015611f34578151835260209283019290910190600101611f15565b50505095945050505050565b634e487b7160e01b5f52604160045260245ffd5b8381528260208201525f604082018351602085015f5b82811015611f88578151845260209384019390910190600101611f6a565b5091979650505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f82611fc357634e487b7160e01b5f52601260045260245ffd5b500690565b81810381811115610cf257634e487b7160e01b5f52601160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561201057612010611f40565b604052919050565b5f82601f830112612027575f5ffd5b5f6120326040611fe7565b9050806040840185811115612045575f5ffd5b845b8181101561205f578035835260209283019201612047565b509195945050505050565b5f5f5f610180848603121561207d575f5ffd5b5f85601f86011261208c575f5ffd5b505f8061010061209b81611fe7565b92508291508601878111156120ae575f5ffd5b865b818110156120c85780358452602093840193016120b0565b508195506120d68882612018565b9450505050611cbb856101408601612018565b5f82601f8301126120f8575f5ffd5b813567ffffffffffffffff81111561211257612112611f40565b612125601f8201601f1916602001611fe7565b818152846020838601011115612139575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f6101808486031215612168575f5ffd5b5f85601f860112612177575f5ffd5b505f8061014061218681611fe7565b9250829150860187811115612199575f5ffd5b865b818110156121b357803584526020938401930161219b565b5090945035905067ffffffffffffffff8111156121ce575f5ffd5b6121da868287016120e9565b92505061016084013567ffffffffffffffff8111156121f7575f5ffd5b612203868287016120e9565b915050925092509256fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212200cbadece0a5dfa93c261112f07442f917c251e8755256b55dd86038722b8222964736f6c634300081c0033",
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
