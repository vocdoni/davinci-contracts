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
	Bin: "0x6080604052348015600e575f5ffd5b506122858061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063233ace1114610059578063b1c3a00e1461008c578063b8e72af6146100ae578063dc37819a146100c3578063df078bef146100d6575b5f5ffd5b6040517f6faf7e5a81fdf3779b989e9751e35285d06ff157f8abf33d388ffbd3e6609d2d81526020015b60405180910390f35b61009f61009a366004611c80565b6100e9565b60405161008393929190611cc6565b6100c16100bc366004611d6c565b610171565b005b6100c16100d1366004611de9565b6101f3565b6100c16100e4366004611e3e565b610760565b6100f1611ba8565b6100f9611bc6565b5f61010d86358760015b6020020135610c0d565b835261012b6060870135604088013560a089013560808a0135610cfa565b6020850152604084015261014560c0870135876007610103565b60608401526101578535866001610103565b82526101668435856001610103565b905093509350939050565b5f5f5f61017e8787610fd3565b9250925092505f61018f8686611004565b505060405163df078bef60e01b8152909150309063df078bef906101bd908790879087908790600401611ec8565b5f6040518083038186803b1580156101d3575f5ffd5b505afa1580156101e5573d5f5f3e3d5ffd5b505050505050505050505050565b6101fb611bc6565b610203611be4565b61020b611c02565b61021b865f5b602002013561101b565b602084015282525f8061022d8761101b565b60408051600580825260c08201909252929450909250606091906020820160a080368337019050509050602080820190604089018237608060c0890160208301375084516020808701516040515f5160206122305f395f51905f529361029893909291869101611f56565b604051602081830303815290604052805190602001205f1c6102ba9190611fab565b865284518452602085810151858201527f1c140cab0c82e4de4ff1ae2a41d0d5474349ea75bc496416aeb53ef277568d5c6040808701919091527f2fbb9940188ad89db824074a507716d495968a3148c5ab405b799323232207e460608701527f2cc70ec706367e61dea6508b30c03cf3d99ba7c35d82ec84d93e602c0a4e266160808701527f0a99efc8954591f22ca29d17b076d341809ba3569ff0a057fd10997dc7a1f9da60a087015260c0860185905260e086018490527f2322047bf491718c5d8ace9e157b34b88030edcec41064b8ea17b82d911534d26101008701527f1814467f5972b801be3036801e563b4369dd4eda8bb638d5da325420d56f0b756101208701527f10f9f7ca8dde6bbfbc5541a81da08f95d1f1a103e496d261615009bdee3135a26101408701527f01beed4f7d5904b2a91bfd1caa80f2c75faac63252aa5b2eccd61d55473a1502610160870152515f91816101808860085afa90511690508061043f576040516351d49ff760e11b815260040160405180910390fd5b505050505f5f61045a895f6004811061021157610211611e8c565b90925090505f80808061047560408e013560208f01356110ba565b929650909450925090505f8061048c8f6003610211565b915091505f5f61049d8e8e8e61128e565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f166e801f097a494b5665f4f0a531f5fb3198931040f8a483a7e0ceea0a5937196101008e01527f1befdc3a5a30a4a4924a55eaa58c1774bb94eba606b7f6892b3034d34ecbd8aa6101208e01527f123311b16924a0c6b2206e3fea972ee635ff2c00b34e24c2e04cf80e398b31db6101408e01527f20c5aae3d9826ba98debb88ccd1fe5885c16e1066cfce9c9da8b1e75ea45c2cf6101608e01527f205bd38a217126712286531ce1831b9753cf59b9d8e39040f4d9fbf4781cc9e96101808e01527f06fd7033dc1adb79c4912a2d7c8f56145fb2bfd0c2713d8f49620f3c37045ad96101a08e01527f2cb1bb00bf4d39008f9f017c4071643c359a7dc45277551cb4eae7bfb0af2ef16101c08e01527f07d026129c1c5b6ceaed6d79d4481b509ace22bef6581fc6a5ca57b1023ac5936101e08e01527f2bc32b68f3d6a51599950755901f7b97d211c5535de7eebc3de7e7e482f425b56102008e01527f24ca2c78271c501adc900bf3d5aa38bc79022b4c5aad08ee5e1624b0b148ddac6102208e01526102408d018290526102608d018190527f1f8137a25041250243aab7c4376f5378fdd8582667540a5e7080c671464972e46102808e01527f126cc5d05e9d3d029f174fa09c063edbb5a30b6ae4a65e7fb98aa781b94eb51a6102a08e01527f0241f6eabbb636ff354e52e9c46a3267a1237d8d5f2a03d228eca6791ea92ff76102c08e01527f0daa3ef90e9a8311edbc67d06c3ba89bbf57a58e48a00175dff129efbed4bd036102e08e015290925090505f610711611bc6565b6020816103008f60085afa915081158061072d57508051600114155b1561074b57604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b610768611bc6565b60408051600580825260c082019092526060916020820160a080368337019050509050602080820190604085018237608060c085016020830137506040515f5160206122305f395f51905f52906107cb908735906020808a013591869101611f56565b604051602081830303815290604052805190602001205f1c6107ed9190611fab565b8252604080515f918782377f1c140cab0c82e4de4ff1ae2a41d0d5474349ea75bc496416aeb53ef277568d5c60408201527f2fbb9940188ad89db824074a507716d495968a3148c5ab405b799323232207e460608201527f2cc70ec706367e61dea6508b30c03cf3d99ba7c35d82ec84d93e602c0a4e266160808201527f0a99efc8954591f22ca29d17b076d341809ba3569ff0a057fd10997dc7a1f9da60a082015260408660c08301377f2322047bf491718c5d8ace9e157b34b88030edcec41064b8ea17b82d911534d26101008201527f1814467f5972b801be3036801e563b4369dd4eda8bb638d5da325420d56f0b756101208201527f10f9f7ca8dde6bbfbc5541a81da08f95d1f1a103e496d261615009bdee3135a26101408201527f01beed4f7d5904b2a91bfd1caa80f2c75faac63252aa5b2eccd61d55473a15026101608201526020816101808360085afa905116905080610962576040516351d49ff760e11b815260040160405180910390fd5b5f5f61099786868a60028060200260405190810160405280929190826002602002808284375f9201919091525061128e915050565b915091506040516101008a82377f166e801f097a494b5665f4f0a531f5fb3198931040f8a483a7e0ceea0a5937196101008201527f1befdc3a5a30a4a4924a55eaa58c1774bb94eba606b7f6892b3034d34ecbd8aa6101208201527f123311b16924a0c6b2206e3fea972ee635ff2c00b34e24c2e04cf80e398b31db6101408201527f20c5aae3d9826ba98debb88ccd1fe5885c16e1066cfce9c9da8b1e75ea45c2cf6101608201527f205bd38a217126712286531ce1831b9753cf59b9d8e39040f4d9fbf4781cc9e96101808201527f06fd7033dc1adb79c4912a2d7c8f56145fb2bfd0c2713d8f49620f3c37045ad96101a08201527f2cb1bb00bf4d39008f9f017c4071643c359a7dc45277551cb4eae7bfb0af2ef16101c08201527f07d026129c1c5b6ceaed6d79d4481b509ace22bef6581fc6a5ca57b1023ac5936101e08201527f2bc32b68f3d6a51599950755901f7b97d211c5535de7eebc3de7e7e482f425b56102008201527f24ca2c78271c501adc900bf3d5aa38bc79022b4c5aad08ee5e1624b0b148ddac61022082015282610240820152816102608201527f1f8137a25041250243aab7c4376f5378fdd8582667540a5e7080c671464972e46102808201527f126cc5d05e9d3d029f174fa09c063edbb5a30b6ae4a65e7fb98aa781b94eb51a6102a08201527f0241f6eabbb636ff354e52e9c46a3267a1237d8d5f2a03d228eca6791ea92ff76102c08201527f0daa3ef90e9a8311edbc67d06c3ba89bbf57a58e48a00175dff129efbed4bd036102e08201526020816103008360085afa905116925082610c0257604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b5f5f5160206122105f395f51905f5283101580610c3757505f5160206122105f395f51905f528210155b15610c5557604051631ff3747d60e21b815260040160405180910390fd5b82158015610c61575081155b15610c6d57505f610cf4565b5f610ca85f5160206122105f395f51905f5260035f5160206122105f395f51905f52875f5160206122105f395f51905f52898a0909086118f2565b9050808303610cbd575050600182901b610cf4565b610cc681611954565b8303610cd9575050600182811b17610cf4565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b5f5f5f5160206122105f395f51905f5286101580610d2557505f5160206122105f395f51905f528510155b80610d3d57505f5160206122105f395f51905f528410155b80610d5557505f5160206122105f395f51905f528310155b15610d7357604051631ff3747d60e21b815260040160405180910390fd5b828486881717175f03610d8a57505f905080610fca565b5f80805f5160206122105f395f51905f52610db360035f5160206122105f395f51905f52611fca565b5f5160206122105f395f51905f528a8c090990505f5f5160206122105f395f51905f528a5f5160206122105f395f51905f528c8d090990505f5f5160206122105f395f51905f528a5f5160206122105f395f51905f528c8d090990505f5160206122105f395f51905f52805f5160206122105f395f51905f528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610ea15f5160206122105f395f51905f52805f5160206122105f395f51905f528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611954565b93505050505f5f610eee5f5160206122105f395f51905f5280610ec657610ec6611f97565b5f5160206122105f395f51905f528586095f5160206122105f395f51905f52878809086118f2565b9050610f395f5160206122105f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f5160206122105f395f51905f528488080961196c565b15915050610f488383836119b4565b90935091508683148015610f5b57508186145b15610f835780610f6b575f610f6e565b60025b60ff1660028a901b175f179450879350610fc6565b610f8c83611954565b87148015610fa15750610f9e82611954565b86145b15610cd95780610fb1575f610fb4565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b610fdb611c21565b610fe3611be4565b610feb611be4565b610ff78486018661206c565b9250925092509250925092565b61100c611c40565b606080610ff784860186612157565b5f5f825f0361102e57505f928392509050565b600183811c9250808416145f5160206122105f395f51905f52831061106657604051631ff3747d60e21b815260040160405180910390fd5b6110a05f5160206122105f395f51905f5260035f5160206122105f395f51905f52865f5160206122105f395f51905f5288890909086118f2565b915080156110b4576110b182611954565b91505b50915091565b5f808080851580156110ca575084155b156110df57505f925082915081905080611285565b600286811c945085935060018088161490808816145f5160206122105f395f51905f528610158061111d57505f5160206122105f395f51905f528510155b1561113b57604051631ff3747d60e21b815260040160405180910390fd5b5f5f5160206122105f395f51905f5261116260035f5160206122105f395f51905f52611fca565b5f5160206122105f395f51905f52888a090990505f5f5160206122105f395f51905f52885f5160206122105f395f51905f528a8b090990505f5f5160206122105f395f51905f52885f5160206122105f395f51905f528a8b090990505f5160206122105f395f51905f52805f5160206122105f395f51905f528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e50896506112505f5160206122105f395f51905f52805f5160206122105f395f51905f528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611954565b955061125d8787866119b4565b9097509550841561127f5761127187611954565b965061127c86611954565b95505b50505050505b92959194509250565b5f5f5f60019050604051604081015f7f2b0f908f6af4a924cce5dd69d0db54a0530547d649378be97369cc2a7f12125b83527f0fd8bcaad7d28069de6620a6817c5c1f0751c1844469a595f512205bfa7d494e6020840152865182526020870151602083015260408360808560065afa841693507f12b76f317d9009eda11266f648146a5f07bdce6de84a6971cabf88cff03764aa82527f193d493034f3b7106b9266fd44d1a9e03a63bcbf70aff1819f9df914b0f1b0036020830152883590508060408301525f5160206122305f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f2651b907b5cf153726d01ea9d4a4d1810b44c8a4927a313419da03f6599d435082527f2929ba32c056cf6e7456f4b7de19c2e3ca7246eb8500217ec43fa4575d4c44fb6020830152602089013590508060408301525f5160206122305f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f02bcb5760dd4347e9f69d0155c64ce2ee73bd1ea6cb509829d8fbe890c87690782527f14d7887914a60649f898def706292e9a9612b9f225563745602fede82bc54e2f6020830152604089013590508060408301525f5160206122305f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f1803853366c081025c008a16df6cbbe1d7a100f632de9f4a86b7bb2dee507ed882527f048ea72d8fe8e2099551423cbbec8d3a7b3614f7e3a3b01fc16a0dbaf9781f3b6020830152606089013590508060408301525f5160206122305f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f0a06c2839668685f72f3948a3e2df680ec42f8c5790a9ec53959b6086ddc7ac382527f148e35b3b99cc26f96341ce4459a0e2d1377a1b23deb7cad6a657f570ec0b5516020830152608089013590508060408301525f5160206122305f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f042d5f42a35b08c818de2f26e6fa902a36705810d001b71d0c1d4de6034f852882527f1afb04b304108af56ad3fe492616d9a9bde49301dab5818f6d5d8ba968070140602083015260a089013590508060408301525f5160206122305f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f2697b686496a1c8c52323cb8fc960dcaa21d09192a3a166df5488ba38470793c82527f09f0f8b395198280124dead1867510244c8630d44a6e8005a447ec034f1f4d66602083015260c089013590508060408301525f5160206122305f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f206797702bdf8f7b75eb63583efa8b24f7e4528828a189e6c7dedb8b9c091cd082527f0e010582fb9ed633c54835184eff34329594ce2412c1951d40c72809e997557e602083015260e089013590508060408301525f5160206122305f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f10b8a9407f87677ca992a5fdcab076dbc3c5d9aa44b5442ba1aa1cfd4f97ce3c82527f20747e2b1d10812647d2f0237de903abb6a642a9e4d0336c5ac003e8ba691a99602083015261010089013590508060408301525f5160206122305f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f1c0ea08075cd060a095a5dc8503aa43f8eb8f688f7411187157fb4c095c2348582527f131bef32098f8b695bf3bd9eb7e999c71082c10280a01eaf8d4293e546dc0834602083015261012089013590508060408301525f5160206122305f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa7f1337a7b52bb259bacba9a02e3e7129b82ad220e36f7852588cb2b67f66b2d2f083527f14a4b5286af4e41f84d4b6f03ce7ff90ac1d5957999c36de3a3296dfcf1526896020840152885160408085018290525f5160206122305f395f51905f5290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806118e95760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b5f61191d827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611ae7565b9050815f5160206122105f395f51905f528283091461194f57604051631ff3747d60e21b815260040160405180910390fd5b919050565b5f5160206122105f395f51905f529081900681030690565b5f5f611998837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611ae7565b9050825f5160206122105f395f51905f52828309149392505050565b5f80806119e35f5160206122105f395f51905f52808788095f5160206122105f395f51905f52898a09086118f2565b905083156119f7576119f481611954565b90505b611a405f5160206122105f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f5160206122105f395f51905f52848a08096118f2565b92505f5160206122105f395f51905f52611a6a5f5160206122105f395f51905f5260028609611b4a565b860991505f5160206122105f395f51905f52611a955f5160206122105f395f51905f52848509611954565b5f5160206122105f395f51905f528586090886141580611ac957505f5160206122105f395f51905f52808385096002098514155b156118e957604051631ff3747d60e21b815260040160405180910390fd5b5f5f60405160208152602080820152602060408201528460608201528360808201525f5160206122105f395f51905f5260a082015260208160c08360055afa90519250905080610cf257604051631ff3747d60e21b815260040160405180910390fd5b5f611b75827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611ae7565b90505f5160206122105f395f51905f5281830960011461194f57604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b604051806101400160405280600a906020820280368337509192915050565b806101008101831015610cf4575f5ffd5b8060408101831015610cf4575f5ffd5b5f5f5f6101808486031215611c93575f5ffd5b611c9d8585611c5f565b9250611cad856101008601611c70565b9150611cbd856101408601611c70565b90509250925092565b60c0810181855f5b6004811015611ced578151835260209283019290910190600101611cce565b50505060808201845f5b6001811015611d16578151835260209283019290910190600101611cf7565b5050508260a0830152949350505050565b5f5f83601f840112611d37575f5ffd5b50813567ffffffffffffffff811115611d4e575f5ffd5b602083019150836020828501011115611d65575f5ffd5b9250929050565b5f5f5f5f60408587031215611d7f575f5ffd5b843567ffffffffffffffff811115611d95575f5ffd5b611da187828801611d27565b909550935050602085013567ffffffffffffffff811115611dc0575f5ffd5b611dcc87828801611d27565b95989497509550505050565b806101408101831015610cf4575f5ffd5b5f5f5f5f6102008587031215611dfd575f5ffd5b6080850186811115611e0d575f5ffd5b85945060a0860187811115611e20575f5ffd5b909350359150611e338660c08701611dd8565b905092959194509250565b5f5f5f5f6102c08587031215611e52575f5ffd5b611e5c8686611c5f565b9350611e6c866101008701611c70565b9250611e7c866101408701611c70565b9150611e33866101808701611dd8565b634e487b7160e01b5f52603260045260245ffd5b805f5b6002811015611ec2578151845260209384019390910190600101611ea3565b50505050565b6102c0810181865f5b6008811015611ef0578151835260209283019290910190600101611ed1565b505050611f01610100830186611ea0565b611f0f610140830185611ea0565b6101808201835f5b600a811015611f36578151835260209283019290910190600101611f17565b50505095945050505050565b634e487b7160e01b5f52604160045260245ffd5b8381528260208201525f604082018351602085015f5b82811015611f8a578151845260209384019390910190600101611f6c565b5091979650505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f82611fc557634e487b7160e01b5f52601260045260245ffd5b500690565b81810381811115610cf457634e487b7160e01b5f52601160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561201257612012611f42565b604052919050565b5f82601f830112612029575f5ffd5b5f6120346040611fe9565b9050806040840185811115612047575f5ffd5b845b81811015612061578035835260209283019201612049565b509195945050505050565b5f5f5f610180848603121561207f575f5ffd5b5f85601f86011261208e575f5ffd5b505f8061010061209d81611fe9565b92508291508601878111156120b0575f5ffd5b865b818110156120ca5780358452602093840193016120b2565b508195506120d8888261201a565b9450505050611cbd85610140860161201a565b5f82601f8301126120fa575f5ffd5b813567ffffffffffffffff81111561211457612114611f42565b612127601f8201601f1916602001611fe9565b81815284602083860101111561213b575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f610180848603121561216a575f5ffd5b5f85601f860112612179575f5ffd5b505f8061014061218881611fe9565b925082915086018781111561219b575f5ffd5b865b818110156121b557803584526020938401930161219d565b5090945035905067ffffffffffffffff8111156121d0575f5ffd5b6121dc868287016120eb565b92505061016084013567ffffffffffffffff8111156121f9575f5ffd5b612205868287016120eb565b915050925092509256fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a264697066735822122066cd26f20595a8bca435fc1e0443fac208fbba23e13b6c9c8acf2dfac23ea5b964736f6c634300081c0033",
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
