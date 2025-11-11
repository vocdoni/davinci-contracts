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
	Bin: "0x6080604052348015600e575f5ffd5b506122818061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063233ace1114610059578063b1c3a00e1461008c578063b8e72af6146100ae578063dc37819a146100c3578063df078bef146100d6575b5f5ffd5b6040517f7fa5475263018a09b457daf45caaecf30b2545305db604f494c45f900a54c81181526020015b60405180910390f35b61009f61009a366004611c7c565b6100e9565b60405161008393929190611cc2565b6100c16100bc366004611d68565b610171565b005b6100c16100d1366004611de5565b6101f3565b6100c16100e4366004611e3a565b61075e565b6100f1611ba4565b6100f9611bc2565b5f61010d86358760015b6020020135610c09565b835261012b6060870135604088013560a089013560808a0135610cf6565b6020850152604084015261014560c0870135876007610103565b60608401526101578535866001610103565b82526101668435856001610103565b905093509350939050565b5f5f5f61017e8787610fcf565b9250925092505f61018f8686611000565b505060405163df078bef60e01b8152909150309063df078bef906101bd908790879087908790600401611ec4565b5f6040518083038186803b1580156101d3575f5ffd5b505afa1580156101e5573d5f5f3e3d5ffd5b505050505050505050505050565b6101fb611bc2565b610203611be0565b61020b611bfe565b61021b865f5b6020020135611017565b602084015282525f8061022d87611017565b60408051600580825260c08201909252929450909250606091906020820160a080368337019050509050602080820190604089018237608060a0890160208301375084516020808701516040515f51602061222c5f395f51905f529361029893909291869101611f52565b604051602081830303815290604052805190602001205f1c6102ba9190611fa7565b865284518452602085810151858201527f0d5488f8f9d7178a089446059d7307960c83102ee7b543273314cce224ff39956040808701919091527f16211e1b5f65f888f262117ecd85a4b6307e3701758683c245c1a48fea1901b460608701527f155a235e10546dc0cd96707e44b1fa2a74a48b67943e2b1edee2c52cc56d99ac60808701527f296cc946ba490faa5d35c8abd4d62c83396632ecb6426b4a3c5ad8834dc371b160a087015260c0860185905260e086018490527f0f606aeb6dc61925748db9ea874fba6d594e7ea13b4806ac2bbf5691e00c995a6101008701527f06fc5e10b7a64accbb2920533e49dddf6d1c1e763e77efd2e849ccbc1348b0426101208701527f03739bf17a7ac46bc2c271a4a2505134a6cd3c4aea758bebff3ae598c9d509af6101408701527f1f51c44dd848dae5bfbf8e8f0af877413ea6d6ae8c059e47c0f51d70426e97e4610160870152515f91816101808860085afa90511690508061043f576040516351d49ff760e11b815260040160405180910390fd5b505050505f5f61045a895f6004811061021157610211611e88565b90925090505f80808061047560408e013560208f01356110b6565b929650909450925090505f8061048c8f6003610211565b915091505f5f61049d8e8e8e61128a565b8b8d5260208d018b905260408d0189905260608d018a905260808d0187905260a08d0188905260c08d0186905260e08d018590527f158cecbb15c159d7e17f68eb4cb1e1403e8f28192355663d2a03a3309ca8e9f16101008e01527ea827d7806ae7c749b4aac9af99f0e5c1dbacca628582721d076989231ceb2b6101208e01527f2277ec685c856882c717b7469231072ce696ed8944ed65c3c56277f4a3a7920f6101408e01527f2a6d609e8ee9457474ee9f79e4c9c46ff3013b7a2072579111f0eb2c19de667c6101608e01527f18fbd18fcdc5525eb45876953ec0730e626a7db97e6966b3bb440cde2656daa36101808e01527f27336bc574ec08e319780a043c87980491b3ed11ef3a2d59464533a30e013d826101a08e01527f0f6b721013c3c00927ec20346c88d6f6fa32696fffb0aa86c3526228efebf3606101c08e01527f135c8dd4c17db3567b12d088770211848d83f405dd227179ca25abb65f5ec0136101e08e01527f1770aac9a216c0ca38d400b7ce62e2a0f9f9d39c5577c2b0e23f9fcce1e030a96102008e01527f0af354f4c4a40b853e92f5649bd1ee0af9fd04337f9bed2ed411db9aa5a60eb66102208e01526102408d018290526102608d018190527f2491c0d8c72d3cb2742167a5c478b500579993779adc785fa95dc5c68a1a61656102808e01527e3c934820a9c85885bddef031e3c1ad4c39cdaf2088f6d50d2356082df33b686102a08e01527f29dd5174ae39f8ff6c7ef73ddfe1adf4a3df96a40522849ccb0886564be4d5966102c08e01527f2b530c82e7be3d9f2e4c80c45235a270d67253fd0ebbefcf29e95e53f82d7e706102e08e015290925090505f61070f611bc2565b6020816103008f60085afa915081158061072b57508051600114155b1561074957604051631ff3747d60e21b815260040160405180910390fd5b50505050505050505050505050505050505050565b610766611bc2565b60408051600580825260c082019092526060916020820160a080368337019050509050602080820190604085018237608060a085016020830137506040515f51602061222c5f395f51905f52906107c9908735906020808a013591869101611f52565b604051602081830303815290604052805190602001205f1c6107eb9190611fa7565b8252604080515f918782377f0d5488f8f9d7178a089446059d7307960c83102ee7b543273314cce224ff399560408201527f16211e1b5f65f888f262117ecd85a4b6307e3701758683c245c1a48fea1901b460608201527f155a235e10546dc0cd96707e44b1fa2a74a48b67943e2b1edee2c52cc56d99ac60808201527f296cc946ba490faa5d35c8abd4d62c83396632ecb6426b4a3c5ad8834dc371b160a082015260408660c08301377f0f606aeb6dc61925748db9ea874fba6d594e7ea13b4806ac2bbf5691e00c995a6101008201527f06fc5e10b7a64accbb2920533e49dddf6d1c1e763e77efd2e849ccbc1348b0426101208201527f03739bf17a7ac46bc2c271a4a2505134a6cd3c4aea758bebff3ae598c9d509af6101408201527f1f51c44dd848dae5bfbf8e8f0af877413ea6d6ae8c059e47c0f51d70426e97e46101608201526020816101808360085afa905116905080610960576040516351d49ff760e11b815260040160405180910390fd5b5f5f61099586868a60028060200260405190810160405280929190826002602002808284375f9201919091525061128a915050565b915091506040516101008a82377f158cecbb15c159d7e17f68eb4cb1e1403e8f28192355663d2a03a3309ca8e9f16101008201527ea827d7806ae7c749b4aac9af99f0e5c1dbacca628582721d076989231ceb2b6101208201527f2277ec685c856882c717b7469231072ce696ed8944ed65c3c56277f4a3a7920f6101408201527f2a6d609e8ee9457474ee9f79e4c9c46ff3013b7a2072579111f0eb2c19de667c6101608201527f18fbd18fcdc5525eb45876953ec0730e626a7db97e6966b3bb440cde2656daa36101808201527f27336bc574ec08e319780a043c87980491b3ed11ef3a2d59464533a30e013d826101a08201527f0f6b721013c3c00927ec20346c88d6f6fa32696fffb0aa86c3526228efebf3606101c08201527f135c8dd4c17db3567b12d088770211848d83f405dd227179ca25abb65f5ec0136101e08201527f1770aac9a216c0ca38d400b7ce62e2a0f9f9d39c5577c2b0e23f9fcce1e030a96102008201527f0af354f4c4a40b853e92f5649bd1ee0af9fd04337f9bed2ed411db9aa5a60eb661022082015282610240820152816102608201527f2491c0d8c72d3cb2742167a5c478b500579993779adc785fa95dc5c68a1a61656102808201527e3c934820a9c85885bddef031e3c1ad4c39cdaf2088f6d50d2356082df33b686102a08201527f29dd5174ae39f8ff6c7ef73ddfe1adf4a3df96a40522849ccb0886564be4d5966102c08201527f2b530c82e7be3d9f2e4c80c45235a270d67253fd0ebbefcf29e95e53f82d7e706102e08201526020816103008360085afa905116925082610bfe57604051631ff3747d60e21b815260040160405180910390fd5b505050505050505050565b5f5f51602061220c5f395f51905f5283101580610c3357505f51602061220c5f395f51905f528210155b15610c5157604051631ff3747d60e21b815260040160405180910390fd5b82158015610c5d575081155b15610c6957505f610cf0565b5f610ca45f51602061220c5f395f51905f5260035f51602061220c5f395f51905f52875f51602061220c5f395f51905f52898a0909086118ee565b9050808303610cb9575050600182901b610cf0565b610cc281611950565b8303610cd5575050600182811b17610cf0565b604051631ff3747d60e21b815260040160405180910390fd5b505b92915050565b5f5f5f51602061220c5f395f51905f5286101580610d2157505f51602061220c5f395f51905f528510155b80610d3957505f51602061220c5f395f51905f528410155b80610d5157505f51602061220c5f395f51905f528310155b15610d6f57604051631ff3747d60e21b815260040160405180910390fd5b828486881717175f03610d8657505f905080610fc6565b5f80805f51602061220c5f395f51905f52610daf60035f51602061220c5f395f51905f52611fc6565b5f51602061220c5f395f51905f528a8c090990505f5f51602061220c5f395f51905f528a5f51602061220c5f395f51905f528c8d090990505f5f51602061220c5f395f51905f528a5f51602061220c5f395f51905f528c8d090990505f51602061220c5f395f51905f52805f51602061220c5f395f51905f528c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610e9d5f51602061220c5f395f51905f52805f51602061220c5f395f51905f528e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611950565b93505050505f5f610eea5f51602061220c5f395f51905f5280610ec257610ec2611f93565b5f51602061220c5f395f51905f528586095f51602061220c5f395f51905f52878809086118ee565b9050610f355f51602061220c5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f51602061220c5f395f51905f5284880809611968565b15915050610f448383836119b0565b90935091508683148015610f5757508186145b15610f7f5780610f67575f610f6a565b60025b60ff1660028a901b175f179450879350610fc2565b610f8883611950565b87148015610f9d5750610f9a82611950565b86145b15610cd55780610fad575f610fb0565b60025b60ff1660028a901b1760011794508793505b5050505b94509492505050565b610fd7611c1d565b610fdf611be0565b610fe7611be0565b610ff384860186612068565b9250925092509250925092565b611008611c3c565b606080610ff384860186612153565b5f5f825f0361102a57505f928392509050565b600183811c9250808416145f51602061220c5f395f51905f52831061106257604051631ff3747d60e21b815260040160405180910390fd5b61109c5f51602061220c5f395f51905f5260035f51602061220c5f395f51905f52865f51602061220c5f395f51905f5288890909086118ee565b915080156110b0576110ad82611950565b91505b50915091565b5f808080851580156110c6575084155b156110db57505f925082915081905080611281565b600286811c945085935060018088161490808816145f51602061220c5f395f51905f528610158061111957505f51602061220c5f395f51905f528510155b1561113757604051631ff3747d60e21b815260040160405180910390fd5b5f5f51602061220c5f395f51905f5261115e60035f51602061220c5f395f51905f52611fc6565b5f51602061220c5f395f51905f52888a090990505f5f51602061220c5f395f51905f52885f51602061220c5f395f51905f528a8b090990505f5f51602061220c5f395f51905f52885f51602061220c5f395f51905f528a8b090990505f51602061220c5f395f51905f52805f51602061220c5f395f51905f528a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508965061124c5f51602061220c5f395f51905f52805f51602061220c5f395f51905f528c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611950565b95506112598787866119b0565b9097509550841561127b5761126d87611950565b965061127886611950565b95505b50505050505b92959194509250565b5f5f5f60019050604051604081015f7f19c8bcbb8ed7050f62e10739c8ebf7038a082247936415bc5903c864848c1bec83527f1fb57080e9f02eed82a1a9dc0d223219273138d5ff9414427265ac792ae11c8b6020840152865182526020870151602083015260408360808560065afa841693507f0358db00af07acd5f75371f6087f1497d058b60a67c06c06875be6d7be5b72b482527f23e175ee982f9eabb473ffc1867d87b8f82963122480d8b4d69a30d9e80bda156020830152883590508060408301525f51602061222c5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f1ebc5cffc99cd925c5953afb55c8f3d895bbb27f2fc013be7e9c0a85a0a9065682527f191ec4139fd065656b5e1169e20fe0fbb5dde1d044f56cae5c30af465c7d869d6020830152602089013590508060408301525f51602061222c5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f24f148662c1da59035257e7dec99ee75d9718b5395fd7c888567adc3bce8796b82527f0d644482a4d3a1f2fbac1cd57c6085a907a61ae3a029fb3dc16a34f3f830bc066020830152604089013590508060408301525f51602061222c5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f07a5d21e173bb6d2b626c20633a15d69857fe665e7377aee3acb12867fc81c1882527f2617b6ee8dad11d2df5bf2f9ef9fa4f300bcf56458e2a588e0c1beca634c61966020830152606089013590508060408301525f51602061222c5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f0e15c43490d87d7aa9acca7f375a4de8a265b54e76beed0f72de2a7dd7ec160682527f2a9ece89d1c833fdfce0bd07d7b6a2524e4526eed90148d350f4fbf51e6ddb8d6020830152608089013590508060408301525f51602061222c5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f28ca2ba26a990ac38f865d664a272fc8ce63da8b4d2fa508563b143306267ef782527f16cd001f536da909d68057e8d9c6d50ccb54931433f77368424ce72238b8dc46602083015260a089013590508060408301525f51602061222c5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f219de828fdebfcabcafd4050f07d7b7d6f2e198b5e78e1850b1c474586805c0b82527f2928d216ba8b0e950fe550dcaf98709d2145ead6ecfd233720f6c7ff86cbbba9602083015260c089013590508060408301525f51602061222c5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f05adba0956d80dbd1581e64b5e53fbcc9fbf03d7d82727bbdbf5a676b9b85dd582527f2503e5bc239f0139094e3bf0cdb565a24c2fe432d0124ae1e6cf4b76422a6d51602083015260e089013590508060408301525f51602061222c5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f09954b5c9475271966c8415a9727c798764facb06b4c99596e143d1125f15eac82527f2c11a4985fdeb3ef5d8952a98003f29a6e4276823565f5245f8d984dac628ebd602083015261010089013590508060408301525f51602061222c5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa841693507f25a35843e2b1056509fbdb734ded9b85e4b023d8f2dc25461967636208f8f44082527f0a77030450a10fc0f2fa6f469dfda92e4593f5647b7fac4e21b0dc3a87817418602083015261012089013590508060408301525f51602061222c5f395f51905f5281108416935060408260608460075afa8416935060408360808560065afa7f019c94f217e49da9918e5e4ffa41dbfcecf09a0be5822b4cf6eb493970c5284c83527f06cd61bb378b17c5d746749c0965f311dc97500e224cfe0687260296798922946020840152885160408085018290525f51602061222c5f395f51905f5290911091909516169390508160608160075afa831692505060408160808360065afa815160209092015191945090925016806118e55760405163a54f8e2760e01b815260040160405180910390fd5b50935093915050565b5f611919827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611ae3565b9050815f51602061220c5f395f51905f528283091461194b57604051631ff3747d60e21b815260040160405180910390fd5b919050565b5f51602061220c5f395f51905f529081900681030690565b5f5f611994837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611ae3565b9050825f51602061220c5f395f51905f52828309149392505050565b5f80806119df5f51602061220c5f395f51905f52808788095f51602061220c5f395f51905f52898a09086118ee565b905083156119f3576119f081611950565b90505b611a3c5f51602061220c5f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea45f51602061220c5f395f51905f52848a08096118ee565b92505f51602061220c5f395f51905f52611a665f51602061220c5f395f51905f5260028609611b46565b860991505f51602061220c5f395f51905f52611a915f51602061220c5f395f51905f52848509611950565b5f51602061220c5f395f51905f528586090886141580611ac557505f51602061220c5f395f51905f52808385096002098514155b156118e557604051631ff3747d60e21b815260040160405180910390fd5b5f5f60405160208152602080820152602060408201528460608201528360808201525f51602061220c5f395f51905f5260a082015260208160c08360055afa90519250905080610cee57604051631ff3747d60e21b815260040160405180910390fd5b5f611b71827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611ae3565b90505f51602061220c5f395f51905f5281830960011461194b57604051631ff3747d60e21b815260040160405180910390fd5b60405180608001604052806004906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6040518061030001604052806018906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b604051806101400160405280600a906020820280368337509192915050565b806101008101831015610cf0575f5ffd5b8060408101831015610cf0575f5ffd5b5f5f5f6101808486031215611c8f575f5ffd5b611c998585611c5b565b9250611ca9856101008601611c6c565b9150611cb9856101408601611c6c565b90509250925092565b60c0810181855f5b6004811015611ce9578151835260209283019290910190600101611cca565b50505060808201845f5b6001811015611d12578151835260209283019290910190600101611cf3565b5050508260a0830152949350505050565b5f5f83601f840112611d33575f5ffd5b50813567ffffffffffffffff811115611d4a575f5ffd5b602083019150836020828501011115611d61575f5ffd5b9250929050565b5f5f5f5f60408587031215611d7b575f5ffd5b843567ffffffffffffffff811115611d91575f5ffd5b611d9d87828801611d23565b909550935050602085013567ffffffffffffffff811115611dbc575f5ffd5b611dc887828801611d23565b95989497509550505050565b806101408101831015610cf0575f5ffd5b5f5f5f5f6102008587031215611df9575f5ffd5b6080850186811115611e09575f5ffd5b85945060a0860187811115611e1c575f5ffd5b909350359150611e2f8660c08701611dd4565b905092959194509250565b5f5f5f5f6102c08587031215611e4e575f5ffd5b611e588686611c5b565b9350611e68866101008701611c6c565b9250611e78866101408701611c6c565b9150611e2f866101808701611dd4565b634e487b7160e01b5f52603260045260245ffd5b805f5b6002811015611ebe578151845260209384019390910190600101611e9f565b50505050565b6102c0810181865f5b6008811015611eec578151835260209283019290910190600101611ecd565b505050611efd610100830186611e9c565b611f0b610140830185611e9c565b6101808201835f5b600a811015611f32578151835260209283019290910190600101611f13565b50505095945050505050565b634e487b7160e01b5f52604160045260245ffd5b8381528260208201525f604082018351602085015f5b82811015611f86578151845260209384019390910190600101611f68565b5091979650505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f82611fc157634e487b7160e01b5f52601260045260245ffd5b500690565b81810381811115610cf057634e487b7160e01b5f52601160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561200e5761200e611f3e565b604052919050565b5f82601f830112612025575f5ffd5b5f6120306040611fe5565b9050806040840185811115612043575f5ffd5b845b8181101561205d578035835260209283019201612045565b509195945050505050565b5f5f5f610180848603121561207b575f5ffd5b5f85601f86011261208a575f5ffd5b505f8061010061209981611fe5565b92508291508601878111156120ac575f5ffd5b865b818110156120c65780358452602093840193016120ae565b508195506120d48882612016565b9450505050611cb9856101408601612016565b5f82601f8301126120f6575f5ffd5b813567ffffffffffffffff81111561211057612110611f3e565b612123601f8201601f1916602001611fe5565b818152846020838601011115612137575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f6101808486031215612166575f5ffd5b5f85601f860112612175575f5ffd5b505f8061014061218481611fe5565b9250829150860187811115612197575f5ffd5b865b818110156121b1578035845260209384019301612199565b5090945035905067ffffffffffffffff8111156121cc575f5ffd5b6121d8868287016120e7565b92505061016084013567ffffffffffffffff8111156121f5575f5ffd5b612201868287016120e7565b915050925092509256fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a26469706673582212202b943941a3f8be4e1741058962d1a2525680181b29b43c141ec918fe30a075bb64736f6c634300081c0033",
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
