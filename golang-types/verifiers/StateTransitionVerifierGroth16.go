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
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506121fc8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063233ace11146100595780635d26278e1461008c57806360e58346146100a1578063b1c3a00e146100b4578063b8e72af6146100d6575b5f5ffd5b6040517fe04b7d8ff99455e516ad34c8d7fd158a2c5165ed87611fb474783a2c813457c881526020015b60405180910390f35b61009f61009a366004611be7565b6100e9565b005b61009f6100af366004611c5d565b610656565b6100c76100c2366004611cab565b610b03565b60405161008393929190611cf1565b61009f6100e4366004611d97565b610b8b565b6100f1611b1f565b6100f9611b3d565b610101611b5b565b610111865f5b6020020135610c0d565b602084015282525f8061012387610c0d565b60408051600580825260c08201909252929450909250606091906020820160a080368337019050509050602080820190604089018237608060a0890160208301375084516020808701516040515f5160206121a75f395f51905f529361018e93909291869101611e2b565b604051602081830303815290604052805190602001205f1c6101b09190611e80565b865284518452602085810151858201527f198be4751111bd381484a726650193fe895b3ab818b1e71bd53681302ca350dd6040808701919091527f0137f6c9e6e471a4fd33a0ee39bdc938e0410bc69c5573cf5e994225d41873a760608701527f23706f25e923cf98e94d386fc4b894bdfa9bcd53cbf3f4082240a9aa15065f5260808701527f12829491faed424b8876880e23f99b71378851ac86002d1d03b46c1ebb27c10e60a087015260c0860185905260e086018490527f06f2b0b691006851752891ed07dd1237bc870e9eb0eac467121883d8a01ebb096101008701527f301a5c21478cd20dcaffed3e2d1f9303fdb54c29f916019bb75e22bc957ef40c6101208701527f13f38284930926dced82c1940e2fd8260ca26fc7ccb81b449274adc03059e9af6101408701527f2e64ffe35fd4dc682a9bf4cccba4c1a51540665743efc71382ecfb0d0bf926c2610160870152515f91816101808860085afa905116905080610335576040516351d49ff760e11b815260040160405180910390fd5b505050505f5f610350895f6004811061010757610107611e03565b90925090505f80808061036b60408e013560208f0135610cac565b929650909450925090505f806103828f6003610107565b915091505f5f6103938e8e8e610e80565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f23bc01cf9248e5d80dc46b28067fe7e46f0cf3730edadc49828344fb87cdf5ce6101008e01527f1ab0dc99a7e3b3b09d8352c8045eba851ca8a093d4d57f305b6bdd6e7076ae8f6101208e01527f1f6aa33abe93e1b5bfb3f4d5343c831074a2a82ba4fd8156c225179ca83b6fb26101408e01527f1c4fe88e4f31c231ea7ca4b2d58db699311c06d51f7bf646b9006399077ac6bb6101608e01527f066e2d1345f346803283bd403e305b6099ff08e068ab94df60dca31afda03d796101808e01527f27d6ba49e1f199f48169359db9732f6ce3a0fc06823421744bfefd7f38bb935f6101a08e01527f2788c3892fcc3499d769d7ea6e8e55791324653f7454065fd53d5db3d81d13e66101c08e01527f01592a8f04864f82e3a6e8ec8c44755b21e3ae9e5c753596c44a0d0c5c23e7296101e08e01527f16b08f2c7accbf2920f806f3ccf534d71f815006cdd7d21e484d98d02d309ee56102008e01527f0979a834b5ed14f30681a0cb7a949aedb4b550bf8971e9ea05e6266e69bde6976102208e01526102408d018290526102608d018190527f083838ba92dd5f49d21e6ee03782111eeffa2b35c13349bd003890ddeeb783dd6102808e01527f2674c0778f6dbcb96d0ce934d4d960118179406c85de70a85a728cdc4f593b8e6102a08e01527f2061c0fda84e37bcff52071528d4c50536d09150f09d14a3dd471d41236c078b6102c08e01527f193619d2cd4c00b1a086ca697e226390e72de990ee1bae7c7c1efdfb9f5a7b9d6102e08e015290925090505f610607611b1f565b6020816103008f60085afa915081158061062357508051600114155b1561064157604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b61065e611b1f565b60408051600580825260c082019092526060916020820160a080368337019050509050602080820190604085018237608060a085016020830137506040515f5160206121a75f395f51905f52906106c1908735906020808a013591869101611e2b565b604051602081830303815290604052805190602001205f1c6106e39190611e80565b8252604080515f918782377f198be4751111bd381484a726650193fe895b3ab818b1e71bd53681302ca350dd60408201527f0137f6c9e6e471a4fd33a0ee39bdc938e0410bc69c5573cf5e994225d41873a760608201527f23706f25e923cf98e94d386fc4b894bdfa9bcd53cbf3f4082240a9aa15065f5260808201527f12829491faed424b8876880e23f99b71378851ac86002d1d03b46c1ebb27c10e60a082015260408660c08301377f06f2b0b691006851752891ed07dd1237bc870e9eb0eac467121883d8a01ebb096101008201527f301a5c21478cd20dcaffed3e2d1f9303fdb54c29f916019bb75e22bc957ef40c6101208201527f13f38284930926dced82c1940e2fd8260ca26fc7ccb81b449274adc03059e9af6101408201527f2e64ffe35fd4dc682a9bf4cccba4c1a51540665743efc71382ecfb0d0bf926c26101608201526020816101808360085afa905116905080610858576040516351d49ff760e11b815260040160405180910390fd5b5f5f61088d86868a60028060200260405190810160405280929190826002602002808284375f92019190915250610e80915050565b915091506040516101008a82377f23bc01cf9248e5d80dc46b28067fe7e46f0cf3730edadc49828344fb87cdf5ce6101008201527f1ab0dc99a7e3b3b09d8352c8045eba851ca8a093d4d57f305b6bdd6e7076ae8f6101208201527f1f6aa33abe93e1b5bfb3f4d5343c831074a2a82ba4fd8156c225179ca83b6fb26101408201527f1c4fe88e4f31c231ea7ca4b2d58db699311c06d51f7bf646b9006399077ac6bb6101608201527f066e2d1345f346803283bd403e305b6099ff08e068ab94df60dca31afda03d796101808201527f27d6ba49e1f199f48169359db9732f6ce3a0fc06823421744bfefd7f38bb935f6101a08201527f2788c3892fcc3499d769d7ea6e8e55791324653f7454065fd53d5db3d81d13e66101c08201527f01592a8f04864f82e3a6e8ec8c44755b21e3ae9e5c753596c44a0d0c5c23e7296101e08201527f16b08f2c7accbf2920f806f3ccf534d71f815006cdd7d21e484d98d02d309ee56102008201527f0979a834b5ed14f30681a0cb7a949aedb4b550bf8971e9ea05e6266e69bde69761022082015282610240820152816102608201527f083838ba92dd5f49d21e6ee03782111eeffa2b35c13349bd003890ddeeb783dd6102808201527f2674c0778f6dbcb96d0ce934d4d960118179406c85de70a85a728cdc4f593b8e6102a08201527f2061c0fda84e37bcff52071528d4c50536d09150f09d14a3dd471d41236c078b6102c08201527f193619d2cd4c00b1a086ca697e226390e72de990ee1bae7c7c1efdfb9f5a7b9d6102e08201526020816103008360085afa905116925082610af857604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b610b0b611b7a565b610b13611b1f565b5f610b2786358760015b602002013561145b565b8352610b456060870135604088013560a089013560808a0135611548565b60208501526040840152610b5f60c0870135876007610b1d565b6060840152610b718535866001610b1d565b8252610b808435856001610b1d565b905093509350939050565b5f5f5f610b988787611821565b9250925092505f610ba98686611852565b5050604051633072c1a360e11b815290915030906360e5834690610bd7908790879087908790600401611ec7565b5f6040518083038186803b158015610bed575f5ffd5b505afa158015610bff573d5f5f3e3d5ffd5b505050505050505050505050565b5f5f825f03610c2057505f928392509050565b600183811c9250808416145f5160206121875f395f51905f528310610c5857604051631ff3747d60e21b815260040160405180910390fd5b610c925f5160206121875f395f51905f5260035f5160206121875f395f51905f52865f5160206121875f395f51905f528889090908611869565b91508015610ca657610ca3826118cb565b91505b50915091565b5f80808085158015610cbc575084155b15610cd157505f925082915081905080610e77565b600286811c945085935060018088161490808816145f5160206121875f395f51905f5286101580610d0f57505f5160206121875f395f51905f528510155b15610d2d57604051631ff3747d60e21b815260040160405180910390fd5b5f5f5160206121875f395f51905f52610d5460035f5160206121875f395f51905f52611f41565b5f5160206121875f395f51905f52888a090990505f5f5160206121875f395f51905f52885f5160206121875f395f51905f528a8b090990505f5f5160206121875f395f51905f52885f5160206121875f395f51905f528a8b090990505f5160206121875f395f51905f52805f5160206121875f395f51905f528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089650610e425f5160206121875f395f51905f52805f5160206121875f395f51905f528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775086118cb565b9550610e4f8787866118e3565b90975095508415610e7157610e63876118cb565b9650610e6e866118cb565b95505b50505050505b92959194509250565b5f5f5f60019050604051604081015f7f22895f79fef987b9217fd93eaf9429852eb7fdcf5c161e0f398987c8fcd7d61f83527f128eb9bbf4546b8c77c0cabb47f9cb0880e0b17ea18199c0df81d0717bcc05c26020840152865182526020870151602083015260408360808560065afa841693507f06a530644f0087b89ee5a83aadfabdc5d926692c12e242c9964b7cb178466fad82527f29d2622265a47fec2b7e6f428000b498dad261f1cd24bc4c2a17da5526cc58116020830152883590508060408301525f5160206121a75f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f20e96f9ee48a7b677b9474b128624b35ac7b52d3a972e2902398f8149cb321bb82527f2692429e2865cd100554d326490c27ad37fba619cc52fc2e912e8d1f9609c0ac6020830152602089013590508060408301525f5160206121a75f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f1eb866e1ea8ab654975abec658f3d31cd706d122e91349f7a001decbc2b8d97282527f20f71708f085846891cf6517a206e74a6d73ba7411e0633a6e8f8fbbf38f2f666020830152604089013590508060408301525f5160206121a75f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f0dd8d99c404f20954c61cadaa3131be4d543731798a5ca3f385e267d866f1b7582527eea73434a733b63610cd582e3e9a4dfec7154289c7870e9c3861a461b31d9286020830152606089013590508060408301525f5160206121a75f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f016d07e8e6c596a6a2421e8802a2ae03a6ba649aa36615ab864aeb39c5efee9082527f035457bcb0ad95fc8d861282c4839cf64d0970343cb54d5bc014a229dadeea8d6020830152608089013590508060408301525f5160206121a75f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f21a8b061a196dd0ec1417a2bb4daa3a3864c799e3b2fdf056b19435a310d1a5f82527ebed16216b87c9aa0b5f22d660c631f6444a469436ae0920ec41fb3842c6791602083015260a089013590508060408301525f5160206121a75f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f13051dc89f172be1d3f40ec652701dd01e56c89ef603a0b2367302dc959a85b482527f16090247df3d9d1ecbb70014f698ca810017b9f7feab398b7ecb6da468178a72602083015260c089013590508060408301525f5160206121a75f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f2380a4eebf4842b42654d7b7a1b672f9520a2a8a979623731a129a4f7d3e6d1282527f124d4dcb9643566b54e12ad51d844cbf05bdb60c6947ad021964c54ceebb8a34602083015260e089013590508060408301525f5160206121a75f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f094782a11c577a5a0ea737edcf778fce5408f6e6a2f8868307ae0237fa14f87182527f03a8689137dc7779adc13e0edd75b35f6c7b01e280e16a01d5b7f603d7cdc1d4602083015261010089013590508060408301525f5160206121a75f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa7f1a4c1061fab8487011fab24441e93ddb6be4cdd41c0101cd528bb36bb982fc2083527f28d088242ac449c6a7b859bf2b85b4f9b2f3ec4775fe71e794fd95bef998db1f6020840152885160408085018290525f5160206121a75f395f51905f5290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806114525760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b5f5f5160206121875f395f51905f528310158061148557505f5160206121875f395f51905f528210155b156114a357604051631ff3747d60e21b815260040160405180910390fd5b821580156114af575081155b156114bb57505f611542565b5f6114f65f5160206121875f395f51905f5260035f5160206121875f395f51905f52875f5160206121875f395f51905f52898a090908611869565b905080830361150b575050600182901b611542565b611514816118cb565b8303611527575050600182811b17611542565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b5f5f5f5160206121875f395f51905f528610158061157357505f5160206121875f395f51905f528510155b8061158b57505f5160206121875f395f51905f528410155b806115a357505f5160206121875f395f51905f528310155b156115c157604051631ff3747d60e21b815260040160405180910390fd5b828486881717175f036115d857505f905080611818565b5f80805f5160206121875f395f51905f5261160160035f5160206121875f395f51905f52611f41565b5f5160206121875f395f51905f528a8c090990505f5f5160206121875f395f51905f528a5f5160206121875f395f51905f528c8d090990505f5f5160206121875f395f51905f528a5f5160206121875f395f51905f528c8d090990505f5160206121875f395f51905f52805f5160206121875f395f51905f528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50894506116ef5f5160206121875f395f51905f52805f5160206121875f395f51905f528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e775086118cb565b93505050505f5f61173c5f5160206121875f395f51905f528061171457611714611e6c565b5f5160206121875f395f51905f528586095f5160206121875f395f51905f5287880908611869565b90506117875f5160206121875f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f5160206121875f395f51905f5284880809611a16565b159150506117968383836118e3565b909350915086831480156117a957508186145b156117d157806117b9575f6117bc565b60025b60ff1660028a901b175f179450879350611814565b6117da836118cb565b871480156117ef57506117ec826118cb565b86145b1561152757806117ff575f611802565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b611829611b98565b611831611b3d565b611839611b3d565b61184584860186611fe3565b9250925092509250925092565b61185a611bb7565b606080611845848601866120ce565b5f611894827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611a5e565b9050815f5160206121875f395f51905f52828309146118c657604051631ff3747d60e21b815260040160405180910390fd5b919050565b5f5160206121875f395f51905f529081900681030690565b5f80806119125f5160206121875f395f51905f52808788095f5160206121875f395f51905f52898a0908611869565b9050831561192657611923816118cb565b90505b61196f5f5160206121875f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f5160206121875f395f51905f52848a0809611869565b92505f5160206121875f395f51905f526119995f5160206121875f395f51905f5260028609611ac1565b860991505f5160206121875f395f51905f526119c45f5160206121875f395f51905f528485096118cb565b5f5160206121875f395f51905f5285860908861415806119f857505f5160206121875f395f51905f52808385096002098514155b1561145257604051631ff3747d60e21b815260040160405180910390fd5b5f5f611a42837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611a5e565b9050825f5160206121875f395f51905f52828309149392505050565b5f5f60405160208152602080820152602060408201528460608201528360808201525f5160206121875f395f51905f5260a082015260208160c08360055afa9051925090508061154057604051631ff3747d60e21b815260040160405180910390fd5b5f611aec827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611a5e565b90505f5160206121875f395f51905f528183096001146118c657604051631ff3747d60e21b815260040160405180910390fd5b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b806101208101831015611542575f5ffd5b5f5f5f5f6101e08587031215611bfb575f5ffd5b6080850186811115611c0b575f5ffd5b85945060a0860187811115611c1e575f5ffd5b909350359150611c318660c08701611bd6565b905092959194509250565b806101008101831015611542575f5ffd5b8060408101831015611542575f5ffd5b5f5f5f5f6102a08587031215611c71575f5ffd5b611c7b8686611c3c565b9350611c8b866101008701611c4d565b9250611c9b866101408701611c4d565b9150611c31866101808701611bd6565b5f5f5f6101808486031215611cbe575f5ffd5b611cc88585611c3c565b9250611cd8856101008601611c4d565b9150611ce8856101408601611c4d565b90509250925092565b60c0810181855f5b6004811015611d18578151835260209283019290910190600101611cf9565b50505060808201845f5b6001811015611d41578151835260209283019290910190600101611d22565b5050508260a0830152949350505050565b5f5f83601f840112611d62575f5ffd5b50813567ffffffffffffffff811115611d79575f5ffd5b602083019150836020828501011115611d90575f5ffd5b9250929050565b5f5f5f5f60408587031215611daa575f5ffd5b843567ffffffffffffffff811115611dc0575f5ffd5b611dcc87828801611d52565b909550935050602085013567ffffffffffffffff811115611deb575f5ffd5b611df787828801611d52565b95989497509550505050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52604160045260245ffd5b8381528260208201525f604082018351602085015f5b82811015611e5f578151845260209384019390910190600101611e41565b5091979650505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f82611e9a57634e487b7160e01b5f52601260045260245ffd5b500690565b805f5b6002811015611ec1578151845260209384019390910190600101611ea2565b50505050565b6102a0810181865f5b6008811015611eef578151835260209283019290910190600101611ed0565b505050611f00610100830186611e9f565b611f0e610140830185611e9f565b6101808201835f5b6009811015611f35578151835260209283019290910190600101611f16565b50505095945050505050565b8181038181111561154257634e487b7160e01b5f52601160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611f8957611f89611e17565b604052919050565b5f82601f830112611fa0575f5ffd5b5f611fab6040611f60565b9050806040840185811115611fbe575f5ffd5b845b81811015611fd8578035835260209283019201611fc0565b509195945050505050565b5f5f5f6101808486031215611ff6575f5ffd5b5f85601f860112612005575f5ffd5b505f8061010061201481611f60565b9250829150860187811115612027575f5ffd5b865b81811015612041578035845260209384019301612029565b5081955061204f8882611f91565b9450505050611ce8856101408601611f91565b5f82601f830112612071575f5ffd5b813567ffffffffffffffff81111561208b5761208b611e17565b61209e601f8201601f1916602001611f60565b8181528460208386010111156120b2575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f61016084860312156120e1575f5ffd5b5f85601f8601126120f0575f5ffd5b505f806101206120ff81611f60565b9250829150860187811115612112575f5ffd5b865b8181101561212c578035845260209384019301612114565b5090945035905067ffffffffffffffff811115612147575f5ffd5b61215386828701612062565b92505061014084013567ffffffffffffffff811115612170575f5ffd5b61217c86828701612062565b915050925092509256fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212208729ff3dad15401b4274e0810ee5f55e520693f99bd8757331e3d941a60c679164736f6c634300081c0033",
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5d26278e.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[9] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, compressedCommitments, compressedCommitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5d26278e.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[9] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [9]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyCompressedProof(&_StateTransitionVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5d26278e.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[9] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [9]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyCompressedProof(&_StateTransitionVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[9] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyProof", proof, commitments, commitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[9] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [9]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[9] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [9]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyProof0(opts *bind.CallOpts, _proof []byte, _input []byte) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyProof0", _proof, _input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyProof0(_proof []byte, _input []byte) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof0(&_StateTransitionVerifierGroth16.CallOpts, _proof, _input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyProof0(_proof []byte, _input []byte) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof0(&_StateTransitionVerifierGroth16.CallOpts, _proof, _input)
}
