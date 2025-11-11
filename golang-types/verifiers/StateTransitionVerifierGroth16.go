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
	Bin: "0x6080604052348015600e575f5ffd5b506122848061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063233ace1114610059578063b1c3a00e1461008c578063b8e72af6146100ae578063dc37819a146100c3578063df078bef146100d6575b5f5ffd5b6040517f4ed124a36cbbe36aaa24027f876c307cc8754fda785e545ff8028c8f655cd0bd81526020015b60405180910390f35b61009f61009a366004611c7f565b6100e9565b60405161008393929190611cc5565b6100c16100bc366004611d6b565b610171565b005b6100c16100d1366004611de8565b6101f3565b6100c16100e4366004611e3d565b610760565b6100f1611ba7565b6100f9611bc5565b5f61010d86358760015b6020020135610c0d565b835261012b6060870135604088013560a089013560808a0135610cfa565b6020850152604084015261014560c0870135876007610103565b60608401526101578535866001610103565b82526101668435856001610103565b905093509350939050565b5f5f5f61017e8787610fd3565b9250925092505f61018f8686611004565b505060405163df078bef60e01b8152909150309063df078bef906101bd908790879087908790600401611ec7565b5f6040518083038186803b1580156101d3575f5ffd5b505afa1580156101e5573d5f5f3e3d5ffd5b505050505050505050505050565b6101fb611bc5565b610203611be3565b61020b611c01565b61021b865f5b602002013561101b565b602084015282525f8061022d8761101b565b60408051600580825260c08201909252929450909250606091906020820160a080368337019050509050602080820190604089018237608060a0890160208301375084516020808701516040515f51602061222f5f395f51905f529361029893909291869101611f55565b604051602081830303815290604052805190602001205f1c6102ba9190611faa565b865284518452602085810151858201527f19bae03f064c3302de9729c402b6dad6eab9d8c0701b758e4cb2cc65dba212186040808701919091527f1d37016c6213a6c77689b13f9ff0b1fdcf163c24387d888d911780aaf26bceeb60608701527f091886f508431246e4073ffcaf20614a37bbed64d2cde561e0a419863410be1b60808701527f1a1826b73ee8709c349b11258792d9088061df84d2d99c3f66a37e7bdc98039860a087015260c0860185905260e086018490527f0c536af3c56b5d7c8851c6a854a2623f4d6875788fb332f6358c498fa23c83ab6101008701527f21b6635b53e8828c91c9c8a516e4e7c527784d491a20b8ccd39a9ea882a5c2736101208701527f09c2db07ff53e3130e303c5b768fe4f66c5d9220bed11cd6f69015792fcb19e16101408701527f046c5da340186f902183a29d2577f1a3fbbe7c1208effc057a4cecbfd21b18ff610160870152515f91816101808860085afa90511690508061043f576040516351d49ff760e11b815260040160405180910390fd5b505050505f5f61045a895f6004811061021157610211611e8b565b90925090505f80808061047560408e013560208f01356110ba565b929650909450925090505f8061048c8f6003610211565b915091505f5f61049d8e8e8e61128e565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f02dd67761f93373c97af86fb49a9f0b0c1124328fe97d13db237b10bde03a0046101008e01527f0b98177e50cc831cdf3af35a41018b5e4f46f3518685f0973182b347630616be6101208e01527f2ad948734beb1dc0ef810d0f0d76339d3aa6c4fc2d1ad727556dc23e87a9b5296101408e01527f05a73787a72d3454ce138d8a4546c32bdf20e46fe3a858a7b64470f64c0256d16101608e01527f0e97ea97b55d39b32cefe13894dc247663b28d1c9e31881e6c4da9a1b2d76d3e6101808e01527f2262e43f45c681f4a06e0f3bd91d26c46e97a62cca8af71182d9a32dde6255b16101a08e01527f0aa6ec6221a30067f2c182d9a9a000ce5151fc6464d17fbe9b46799160ebcacb6101c08e01527f0be563b993c7d7a2bc76c6ca5f238a9fc85b3509dd03b1e13fa8118843a1493d6101e08e01527f0714f0448087194c075e107363747245c65b389319282b4f379baf037062c0a56102008e01527f289d164651ad45364a290636090c756f027b49a00917ce696b8b65cc8811fc546102208e01526102408d018290526102608d018190527f267f7cb83b3a4422aea0eccbed9b899d4f0a74b151654faa1008eec4ce7f2e056102808e01527f1b501d0a9a60ece8e3b13638f58a101963ec7b2f622fd04dc6e010a73c81eec96102a08e01527f1208fa2755810e2a880f771a68564cf2ca7953f4c820bbc1ca9cab17295b740b6102c08e01527f2c85e6080b3c6d2cf9691db692e1c787bcbb94fde8c7f8b9fd4c7a4511504c396102e08e015290925090505f610711611bc5565b6020816103008f60085afa915081158061072d57508051600114155b1561074b57604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b610768611bc5565b60408051600580825260c082019092526060916020820160a080368337019050509050602080820190604085018237608060a085016020830137506040515f51602061222f5f395f51905f52906107cb908735906020808a013591869101611f55565b604051602081830303815290604052805190602001205f1c6107ed9190611faa565b8252604080515f918782377f19bae03f064c3302de9729c402b6dad6eab9d8c0701b758e4cb2cc65dba2121860408201527f1d37016c6213a6c77689b13f9ff0b1fdcf163c24387d888d911780aaf26bceeb60608201527f091886f508431246e4073ffcaf20614a37bbed64d2cde561e0a419863410be1b60808201527f1a1826b73ee8709c349b11258792d9088061df84d2d99c3f66a37e7bdc98039860a082015260408660c08301377f0c536af3c56b5d7c8851c6a854a2623f4d6875788fb332f6358c498fa23c83ab6101008201527f21b6635b53e8828c91c9c8a516e4e7c527784d491a20b8ccd39a9ea882a5c2736101208201527f09c2db07ff53e3130e303c5b768fe4f66c5d9220bed11cd6f69015792fcb19e16101408201527f046c5da340186f902183a29d2577f1a3fbbe7c1208effc057a4cecbfd21b18ff6101608201526020816101808360085afa905116905080610962576040516351d49ff760e11b815260040160405180910390fd5b5f5f61099786868a60028060200260405190810160405280929190826002602002808284375f9201919091525061128e915050565b915091506040516101008a82377f02dd67761f93373c97af86fb49a9f0b0c1124328fe97d13db237b10bde03a0046101008201527f0b98177e50cc831cdf3af35a41018b5e4f46f3518685f0973182b347630616be6101208201527f2ad948734beb1dc0ef810d0f0d76339d3aa6c4fc2d1ad727556dc23e87a9b5296101408201527f05a73787a72d3454ce138d8a4546c32bdf20e46fe3a858a7b64470f64c0256d16101608201527f0e97ea97b55d39b32cefe13894dc247663b28d1c9e31881e6c4da9a1b2d76d3e6101808201527f2262e43f45c681f4a06e0f3bd91d26c46e97a62cca8af71182d9a32dde6255b16101a08201527f0aa6ec6221a30067f2c182d9a9a000ce5151fc6464d17fbe9b46799160ebcacb6101c08201527f0be563b993c7d7a2bc76c6ca5f238a9fc85b3509dd03b1e13fa8118843a1493d6101e08201527f0714f0448087194c075e107363747245c65b389319282b4f379baf037062c0a56102008201527f289d164651ad45364a290636090c756f027b49a00917ce696b8b65cc8811fc5461022082015282610240820152816102608201527f267f7cb83b3a4422aea0eccbed9b899d4f0a74b151654faa1008eec4ce7f2e056102808201527f1b501d0a9a60ece8e3b13638f58a101963ec7b2f622fd04dc6e010a73c81eec96102a08201527f1208fa2755810e2a880f771a68564cf2ca7953f4c820bbc1ca9cab17295b740b6102c08201527f2c85e6080b3c6d2cf9691db692e1c787bcbb94fde8c7f8b9fd4c7a4511504c396102e08201526020816103008360085afa905116925082610c0257604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b5f5f51602061220f5f395f51905f5283101580610c3757505f51602061220f5f395f51905f528210155b15610c5557604051631ff3747d60e21b815260040160405180910390fd5b82158015610c61575081155b15610c6d57505f610cf4565b5f610ca85f51602061220f5f395f51905f5260035f51602061220f5f395f51905f52875f51602061220f5f395f51905f52898a0909086118f1565b9050808303610cbd575050600182901b610cf4565b610cc681611953565b8303610cd9575050600182811b17610cf4565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b5f5f5f51602061220f5f395f51905f5286101580610d2557505f51602061220f5f395f51905f528510155b80610d3d57505f51602061220f5f395f51905f528410155b80610d5557505f51602061220f5f395f51905f528310155b15610d7357604051631ff3747d60e21b815260040160405180910390fd5b828486881717175f03610d8a57505f905080610fca565b5f80805f51602061220f5f395f51905f52610db360035f51602061220f5f395f51905f52611fc9565b5f51602061220f5f395f51905f528a8c090990505f5f51602061220f5f395f51905f528a5f51602061220f5f395f51905f528c8d090990505f5f51602061220f5f395f51905f528a5f51602061220f5f395f51905f528c8d090990505f51602061220f5f395f51905f52805f51602061220f5f395f51905f528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610ea15f51602061220f5f395f51905f52805f51602061220f5f395f51905f528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611953565b93505050505f5f610eee5f51602061220f5f395f51905f5280610ec657610ec6611f96565b5f51602061220f5f395f51905f528586095f51602061220f5f395f51905f52878809086118f1565b9050610f395f51602061220f5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f51602061220f5f395f51905f528488080961196b565b15915050610f488383836119b3565b90935091508683148015610f5b57508186145b15610f835780610f6b575f610f6e565b60025b60ff1660028a901b175f179450879350610fc6565b610f8c83611953565b87148015610fa15750610f9e82611953565b86145b15610cd95780610fb1575f610fb4565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b610fdb611c20565b610fe3611be3565b610feb611be3565b610ff78486018661206b565b9250925092509250925092565b61100c611c3f565b606080610ff784860186612156565b5f5f825f0361102e57505f928392509050565b600183811c9250808416145f51602061220f5f395f51905f52831061106657604051631ff3747d60e21b815260040160405180910390fd5b6110a05f51602061220f5f395f51905f5260035f51602061220f5f395f51905f52865f51602061220f5f395f51905f5288890909086118f1565b915080156110b4576110b182611953565b91505b50915091565b5f808080851580156110ca575084155b156110df57505f925082915081905080611285565b600286811c945085935060018088161490808816145f51602061220f5f395f51905f528610158061111d57505f51602061220f5f395f51905f528510155b1561113b57604051631ff3747d60e21b815260040160405180910390fd5b5f5f51602061220f5f395f51905f5261116260035f51602061220f5f395f51905f52611fc9565b5f51602061220f5f395f51905f52888a090990505f5f51602061220f5f395f51905f52885f51602061220f5f395f51905f528a8b090990505f5f51602061220f5f395f51905f52885f51602061220f5f395f51905f528a8b090990505f51602061220f5f395f51905f52805f51602061220f5f395f51905f528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50896506112505f51602061220f5f395f51905f52805f51602061220f5f395f51905f528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611953565b955061125d8787866119b3565b9097509550841561127f5761127187611953565b965061127c86611953565b95505b50505050505b92959194509250565b5f5f5f60019050604051604081015f7f10f1d11e274d64701a85df5ca5a13336fde3ddaba2792dbca67d209c60fcd82483527f2654010447cf95ca6e6c5a497919fc77d10a3b3c7b1a8533e39d978be980bbf26020840152865182526020870151602083015260408360808560065afa841693507f01b7c74d463f4d8326ff271ef7e57bfb45ec24a2d0222653ef2ceb3e1009002d82527f118243029732123691b8a2d69dee481f38442c0b6343526c298e4e017a0d79b06020830152883590508060408301525f51602061222f5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f05457ac6eabe79a8c2c989677fbae3fd4f338bae226b13a122742acfb709642082527f2c0ddd418b9096d4e58b4b8f600d236b7c9afc7e3a2d02c2bfbe6f016bcfe1816020830152602089013590508060408301525f51602061222f5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f2d99f5dcfb7ec9a4bf211454f81edefe393d19f5da898d03c70d7849299b291c82527f1c7195993edd6b40d3d38d04ad1c17350983d14eb17c73f1942331cb05e638976020830152604089013590508060408301525f51602061222f5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f18963cc31a66004487366f222ab5d8396c950187a575b6cda6406a6b40648f8882527f093defc115280579003a8d4e70a177b647a4fda6a059d5aa55b1f323e369783c6020830152606089013590508060408301525f51602061222f5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f101e320cf33452653c5ca920e8884d0e5b149ddb9c0d2a211faf172889fab8e982527f2de0850f3cfd63a1aa5598fde108e1dcc4354c52efdbf025ca9036dfb4a3ec2f6020830152608089013590508060408301525f51602061222f5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f295c964c644d13e7f6813d76e16b775f9c726bf17eaa0d6af2371b5fae87846f82527f0a03aba407948aeb6a46d7f6758c7f3713d3a46a1d1cfa9e92ed684efabce389602083015260a089013590508060408301525f51602061222f5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f07c498bb61c3ab7739845e2a016818f9d6e96bc15edf5d97d0acf1aab36041cd82527f16bce5cfb17edbcf4629e59e51a0b7c0355b5ae83824281834edf15e7b3e42b0602083015260c089013590508060408301525f51602061222f5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f1cf4d8a408889f66c354513e396fc82e45591cdc6677dc83b6bc0f3275fbcbcc82527f0705911c521ba02e2cae6dc6e1dae2aad798933e41c81fde4afd2541dbe7930b602083015260e089013590508060408301525f51602061222f5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f1dc8324a0700df851e57ad6b0ea3a727fc7356802c785105a1439a4594057e9f82527f123a87346f0545a81016fbc4d7f53a31451b98e8ebee08c8ca42ee04b25b21c3602083015261010089013590508060408301525f51602061222f5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f120bf67aa4e9ca7ad4671fd982ca6183b6af8ef256aae56486caa3391158b7c182527eaee2d9826cd0c90efc8b61acf46db2bef1aefed8cc3ca8d52bb2e90ecc1388602083015261012089013590508060408301525f51602061222f5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa7f01cb7dfd001d8660640343fc0052a2c4e99c12bfe43de2bc80d39fb2e5b86bcf83527f1a63e89df986e45f6b6df2d0b9f9cfe2ac943552fd5c2a31d6dd26ad517d68bb6020840152885160408085018290525f51602061222f5f395f51905f5290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806118e85760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b5f61191c827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611ae6565b9050815f51602061220f5f395f51905f528283091461194e57604051631ff3747d60e21b815260040160405180910390fd5b919050565b5f51602061220f5f395f51905f529081900681030690565b5f5f611997837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611ae6565b9050825f51602061220f5f395f51905f52828309149392505050565b5f80806119e25f51602061220f5f395f51905f52808788095f51602061220f5f395f51905f52898a09086118f1565b905083156119f6576119f381611953565b90505b611a3f5f51602061220f5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f51602061220f5f395f51905f52848a08096118f1565b92505f51602061220f5f395f51905f52611a695f51602061220f5f395f51905f5260028609611b49565b860991505f51602061220f5f395f51905f52611a945f51602061220f5f395f51905f52848509611953565b5f51602061220f5f395f51905f528586090886141580611ac857505f51602061220f5f395f51905f52808385096002098514155b156118e857604051631ff3747d60e21b815260040160405180910390fd5b5f5f60405160208152602080820152602060408201528460608201528360808201525f51602061220f5f395f51905f5260a082015260208160c08360055afa90519250905080610cf257604051631ff3747d60e21b815260040160405180910390fd5b5f611b74827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611ae6565b90505f51602061220f5f395f51905f5281830960011461194e57604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b604051806101400160405280600a906020820280368337509192915050565b806101008101831015610cf4575f5ffd5b8060408101831015610cf4575f5ffd5b5f5f5f6101808486031215611c92575f5ffd5b611c9c8585611c5e565b9250611cac856101008601611c6f565b9150611cbc856101408601611c6f565b90509250925092565b60c0810181855f5b6004811015611cec578151835260209283019290910190600101611ccd565b50505060808201845f5b6001811015611d15578151835260209283019290910190600101611cf6565b5050508260a0830152949350505050565b5f5f83601f840112611d36575f5ffd5b50813567ffffffffffffffff811115611d4d575f5ffd5b602083019150836020828501011115611d64575f5ffd5b9250929050565b5f5f5f5f60408587031215611d7e575f5ffd5b843567ffffffffffffffff811115611d94575f5ffd5b611da087828801611d26565b909550935050602085013567ffffffffffffffff811115611dbf575f5ffd5b611dcb87828801611d26565b95989497509550505050565b806101408101831015610cf4575f5ffd5b5f5f5f5f6102008587031215611dfc575f5ffd5b6080850186811115611e0c575f5ffd5b85945060a0860187811115611e1f575f5ffd5b909350359150611e328660c08701611dd7565b905092959194509250565b5f5f5f5f6102c08587031215611e51575f5ffd5b611e5b8686611c5e565b9350611e6b866101008701611c6f565b9250611e7b866101408701611c6f565b9150611e32866101808701611dd7565b634e487b7160e01b5f52603260045260245ffd5b805f5b6002811015611ec1578151845260209384019390910190600101611ea2565b50505050565b6102c0810181865f5b6008811015611eef578151835260209283019290910190600101611ed0565b505050611f00610100830186611e9f565b611f0e610140830185611e9f565b6101808201835f5b600a811015611f35578151835260209283019290910190600101611f16565b50505095945050505050565b634e487b7160e01b5f52604160045260245ffd5b8381528260208201525f604082018351602085015f5b82811015611f89578151845260209384019390910190600101611f6b565b5091979650505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f82611fc457634e487b7160e01b5f52601260045260245ffd5b500690565b81810381811115610cf457634e487b7160e01b5f52601160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561201157612011611f41565b604052919050565b5f82601f830112612028575f5ffd5b5f6120336040611fe8565b9050806040840185811115612046575f5ffd5b845b81811015612060578035835260209283019201612048565b509195945050505050565b5f5f5f610180848603121561207e575f5ffd5b5f85601f86011261208d575f5ffd5b505f8061010061209c81611fe8565b92508291508601878111156120af575f5ffd5b865b818110156120c95780358452602093840193016120b1565b508195506120d78882612019565b9450505050611cbc856101408601612019565b5f82601f8301126120f9575f5ffd5b813567ffffffffffffffff81111561211357612113611f41565b612126601f8201601f1916602001611fe8565b81815284602083860101111561213a575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f6101808486031215612169575f5ffd5b5f85601f860112612178575f5ffd5b505f8061014061218781611fe8565b925082915086018781111561219a575f5ffd5b865b818110156121b457803584526020938401930161219c565b5090945035905067ffffffffffffffff8111156121cf575f5ffd5b6121db868287016120ea565b92505061016084013567ffffffffffffffff8111156121f8575f5ffd5b612204868287016120ea565b915050925092509256fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220a5126208e4789ff14522b588e1b18365c39ebc5a87b3b96ea9d7a80e5c3153df64736f6c634300081c0033",
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
