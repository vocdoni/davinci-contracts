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

// ResultsVerifierGroth16MetaData contains all meta data concerning the ResultsVerifierGroth16 contract.
var ResultsVerifierGroth16MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[9]\",\"name\":\"input\",\"type\":\"uint256[9]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557611e79908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461167c575080635d26278e14610c8c57806360e583461461037f578063b1c3a00e1461026e5763b8e72af614610055575f80fd5b346102365760403660031901126102365760043567ffffffffffffffff8111610236576100869036906004016116b4565b60243567ffffffffffffffff8111610236576100a69036906004016116b4565b61010093846040516100b882826116e2565b3690376040938480516100cb82826116e2565b3690378480516100db82826116e2565b3690378101610180828203126102365780601f8301121561023657845195610105610100886116e2565b8201868282116102365783905b82821061025e5750509061014061012c8261013394611bdc565b9301611bdc565b916101209081865161014582826116e2565b36903784019080858303126102365781601f860112156102365785519461016e610120876116e2565b8591810192831161023657905b82821061023a57505050303b15610236578351633072c1a360e11b8152945f600487015b6008821061022057505050906101bd6101c89261010487019061173c565b61014485019061173c565b5f61018484015b6009821061020a575050505f826102a481305afa90811561020157506101f3575080f35b6101ff91505f906116e2565b005b513d5f823e3d90fd5b60208060019285518152019301910190916101cf565b602080600192855181520193019101909161019f565b5f80fd5b813581526020918201910161017b565b634e487b7160e01b5f52604160045260245ffd5b8135815260209182019101610112565b3461023657610180366003190112610236573661010411610236573661014411610236573661018411610236576040516102a96080826116e2565b60803682376040519060206102be81846116e2565b803684376102d0602435600435611907565b82526102e660843560a4356044356064356119a8565b8284015260408301526102fd60e43560c435611907565b60608301526103126101243561010435611907565b83526103246101643561014435611907565b9060405192835f905b60048210610369575050506080830193905f945b600186106103545760c0858560a0820152f35b81806001928551815201930195019491610341565b825181529183019160019190910190830161032d565b34610236576102a036600319011261023657366101041161023657366101441161023657366101841161023657366102a4116102365760405160206103c481836116e2565b803683375f516020611e245f395f51905f5260405182810190610400816103f2610124356101043586611704565b03601f1981018352826116e2565b519020068252604051604061010482377f09ddd27feefb6c023e839129fa943b31d68b893aceab9fddad2d1ba0c109eb6960408201527f0e75d1ba244e4dc3f0920d2e25b381d0bddedee811f9a9c8413f1e99f1df735b60608201527f2ac04b5c087f24ff052155e21769eaaa258a0c04e567115b0132c971c5e43acc60808201527f16418313c3293b318fb0ac3f0225841f52eb6c7f05c0f1937434fff70e707ff260a0820152604061014460c08301377f2c38ce00c3b929eb49faee654b41737dadc4955ff3f1b6217385ed71986a33286101008201527f0546151c4ef9d8f9b23dd5e16cf3451267bc005f5034418b20a95aaa050440336101208201527f2f5014ada8044c975f48a5dc4b495a29da4a7b69d6bdce7068c8886e93d2de716101408201527f2ed4d5366f0d5be9bc5f2c5115f3de39512a781c3b2bc239f5694f65942a3bb461016082015281816101808160085afa90511615610c7d5760408051929061057090846116e2565b61010482845b6101448310610c6d5750505060405190604082017f05574e80ce1ca3eac89b2258e9992dc9fb93794fed201e39bf9882d20a02743183525f516020611e245f395f51905f52604085808601977f16badad6f18757c316b87c55d2bac966845e73544f7bbc25efec0dc24e62a7e389528051855201519260608601938452818660808160065afa947f080589e6a8dccf65c88fcf4af97bfa338a78bc8bf4bbab9d0ce8a6f262c8913782527f1b254d58abe1b851d232c2d833528296b0ec330842b7953e04d72eb31d60fb9585527e4fa179270b6fba9cb57b8701174fac69097c05f6ca65cbbcde38df26c6e45160016101843560808a0198818a5287878760608160075afa92101616858a60808160065afa16167f2de5ced8808ddc533d94af120ac0cc662c719034d455bed024becc28ea2a45cb84527f2a5947a1f6041ba5004c8652e6851c1add9c59dbebf65fc3b0b60aa680d27ace87526101a4359081895286868660608160075afa92101616848960808160065afa167f1bfca0e810e1b22a32ec77709f6cd535baec2256f2e5443ee54dc605e2dcfd1684527f220bc874a83eb6c31b86ff0ee2d52d9fcf69d2cbc61b1657b66d891ad726f55b87526101c4359081895286868660608160075afa92101616848960808160065afa167f078627d8c01cb1b633a7f9f5a203869619c240ea65d9978c9c1d88c529dd1bda84527f0e80bc3778f8029e1a00a9168e819cdb7603303867f1c9836d6a089d18e98afa87526101e4359081895286868660608160075afa92101616848960808160065afa167f1081f50a584ebb3f9baf138534d26f79f960553705d00272c4a6d25048b76f2e84527f16d8c778241a9e067d2176146b0490f1607f643df2be5dc2550f8439aa53f7ae8752610204359081895286868660608160075afa92101616848960808160065afa167f2b729ac370f73e1cf654a384dbf3242450609fd78f09b7d8a527dc246602119484527f2acdd14139e8154d80316218b9aadde07fe819a55b44ad9c6631365e7d0eff7a8752610224359081895286868660608160075afa92101616848960808160065afa167f26d1c3f0ba4bf331e827bc30439d13cdacb334b7235bc3cd76be3448f9b4898a84527f0d93eb6739bf8645d56dc48b39fb8e7d00fa560f16b137a3c03aaf5065222ebb8752610244359081895286868660608160075afa92101616848960808160065afa167f1da07187014ac69f585f284650979ab39b4401e503f3615d3cc16dcc12e9bbab84527f2d6dc91f66e8391fa1d9424c02ffa4ac206d9dbff13e81a584471b26d3e1bd248752610264359081895286868660608160075afa92101616848960808160065afa167f024ce63e778a68b4892c05e7a19a91dbbffc31fbbaa51eefe593d0b83742bb6584527f127dd2dd6bfccca1cf33fc681751540cd0f54af1df8dbd95f4b641140a16a91b8752610284359081895286868660608160075afa92101616848960808160065afa16957f28217303dacdc05cfbcefa4468da9a9b2981987caf637e5b31e8cb8e085813338452525180955260608160075afa9210161660408260808160065afa16905192519015610c5e5760405192610100600485377f04d613d06b5a22c324976f0f6a216bc0f582c42bd6dc3ff85ccaaa6c8836a45e6101008501527f2cc4673c36df19dbc29da663ab5cb91b8845edcd8d4faa3832f11bce68191a2d6101208501527f1c54423463bb05e5b1cb4858307e05dcd7973b48e7c9746c14554f1e3ddc87716101408501527f203163b0201fbf8f183c9bc62f6dd2e305b577491a115305a103a76d40a79a876101608501527f0f76c6aed2085df5e8fcf3c3830a38c1fca0bf7859457198fbff2fe328f1edff6101808501527f2a6463d5b386aa18475945e056d97d0f0ad40636487e27232d36694f08f25eb76101a08501527f26eec16458e0c63586f17a31b05b39de58cf2094562824052b3b3ca91057556a6101c08501527f2bb0c64beacd5a3b46424048bda9bbe7801ab44d15e2fca5774699fadba5668e6101e08501527f27d507df29f94126f2c898b8fbebe6f1903f9f4b038819227048c29fbaafa2ec6102008501527f191e001612f65dc338d3ad86f7a563953ee2a3f457f531031da0760b0d56fbcb6102208501526102408401526102608301527f18e204fe0489cba105b05f21bf9e6cf6e715cf7f1d5f3dec1918f70c726566c36102808301527f2c1ea700992d88fce9d1790f666d992962f9645f5aa701180336a924d7aa818e6102a08301527f1c1cd14a520f9c0b614a3b87443ba6c6f0a92a5a4f5ef40433dd61a9a8be10876102c08301527f06be7a34b436398781cca782735c094c91ce90fdeb5e789ce12b3cc92fb4f8886102e0830152816103008160085afa90511615610c4f57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191849101610576565b6351d49ff760e11b5f5260045ffd5b34610236576101e03660031901126102365736608411610236573660a41161023657366101e411610236576040516020610cc681836116e2565b803683376040805190610cd981836116e2565b8036833761030093815192610cee86856116e2565b85368537610cfd608435611763565b868301528152610d0e60a435611763565b905f516020611e245f395f51905f5283516103f2610d388a87015189519283918d83019586611704565b5190200684528251865286830151878701527f09ddd27feefb6c023e839129fa943b31d68b893aceab9fddad2d1ba0c109eb69858701527f0e75d1ba244e4dc3f0920d2e25b381d0bddedee811f9a9c8413f1e99f1df735b60608701527f2ac04b5c087f24ff052155e21769eaaa258a0c04e567115b0132c971c5e43acc60808701527f16418313c3293b318fb0ac3f0225841f52eb6c7f05c0f1937434fff70e707ff260a087015260c086015260e08501527f2c38ce00c3b929eb49faee654b41737dadc4955ff3f1b6217385ed71986a33286101008501527f0546151c4ef9d8f9b23dd5e16cf3451267bc005f5034418b20a95aaa050440336101208501527f2f5014ada8044c975f48a5dc4b495a29da4a7b69d6bdce7068c8886e93d2de716101408501527f2ed4d5366f0d5be9bc5f2c5115f3de39512a781c3b2bc239f5694f65942a3bb4610160850152825185816101808760085afa90511615610c7d578490610ea8600435611763565b610eb96024959295356044356117ce565b9291909389610ec9606435611763565b9890977e4fa179270b6fba9cb57b8701174fac69097c05f6ca65cbbcde38df26c6e45183519b8c937f05574e80ce1ca3eac89b2258e9992dc9fb93794fed201e39bf9882d20a02743185527f16badad6f18757c316b87c55d2bac966845e73544f7bbc25efec0dc24e62a7e3828601528051868601520151606084019081527f28217303dacdc05cfbcefa4468da9a9b2981987caf637e5b31e8cb8e08581333858560808160065afa957f080589e6a8dccf65c88fcf4af97bfa338a78bc8bf4bbab9d0ce8a6f262c89137818701527f1b254d58abe1b851d232c2d833528296b0ec330842b7953e04d72eb31d60fb958352600160c4356080880198818a525f516020611e245f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f2de5ced8808ddc533d94af120ac0cc662c719034d455bed024becc28ea2a45cb828801527f2a5947a1f6041ba5004c8652e6851c1add9c59dbebf65fc3b0b60aa680d27ace845260e435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1bfca0e810e1b22a32ec77709f6cd535baec2256f2e5443ee54dc605e2dcfd16828801527f220bc874a83eb6c31b86ff0ee2d52d9fcf69d2cbc61b1657b66d891ad726f55b845261010435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f078627d8c01cb1b633a7f9f5a203869619c240ea65d9978c9c1d88c529dd1bda828801527f0e80bc3778f8029e1a00a9168e819cdb7603303867f1c9836d6a089d18e98afa845261012435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1081f50a584ebb3f9baf138534d26f79f960553705d00272c4a6d25048b76f2e828801527f16d8c778241a9e067d2176146b0490f1607f643df2be5dc2550f8439aa53f7ae845261014435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2b729ac370f73e1cf654a384dbf3242450609fd78f09b7d8a527dc2466021194828801527f2acdd14139e8154d80316218b9aadde07fe819a55b44ad9c6631365e7d0eff7a845261016435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f26d1c3f0ba4bf331e827bc30439d13cdacb334b7235bc3cd76be3448f9b4898a828801527f0d93eb6739bf8645d56dc48b39fb8e7d00fa560f16b137a3c03aaf5065222ebb845261018435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1da07187014ac69f585f284650979ab39b4401e503f3615d3cc16dcc12e9bbab828801527f2d6dc91f66e8391fa1d9424c02ffa4ac206d9dbff13e81a584471b26d3e1bd2484526101a435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f024ce63e778a68b4892c05e7a19a91dbbffc31fbbaa51eefe593d0b83742bb65828801527f127dd2dd6bfccca1cf33fc681751540cd0f54af1df8dbd95f4b641140a16a91b84526101c435908189525f516020611e245f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611e245f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610c5e578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f04d613d06b5a22c324976f0f6a216bc0f582c42bd6dc3ff85ccaaa6c8836a45e6101008501527f2cc4673c36df19dbc29da663ab5cb91b8845edcd8d4faa3832f11bce68191a2d6101208501527f1c54423463bb05e5b1cb4858307e05dcd7973b48e7c9746c14554f1e3ddc87716101408501527f203163b0201fbf8f183c9bc62f6dd2e305b577491a115305a103a76d40a79a876101608501527f0f76c6aed2085df5e8fcf3c3830a38c1fca0bf7859457198fbff2fe328f1edff6101808501527f2a6463d5b386aa18475945e056d97d0f0ad40636487e27232d36694f08f25eb76101a08501527f26eec16458e0c63586f17a31b05b39de58cf2094562824052b3b3ca91057556a6101c08501527f2bb0c64beacd5a3b46424048bda9bbe7801ab44d15e2fca5774699fadba5668e6101e08501527f27d507df29f94126f2c898b8fbebe6f1903f9f4b038819227048c29fbaafa2ec6102008501527f191e001612f65dc338d3ad86f7a563953ee2a3f457f531031da0760b0d56fbcb6102208501526102408401526102608301527f18e204fe0489cba105b05f21bf9e6cf6e715cf7f1d5f3dec1918f70c726566c36102808301527f2c1ea700992d88fce9d1790f666d992962f9645f5aa701180336a924d7aa818e6102a08301527f1c1cd14a520f9c0b614a3b87443ba6c6f0a92a5a4f5ef40433dd61a9a8be10876102c08301527f06be7a34b436398781cca782735c094c91ce90fdeb5e789ce12b3cc92fb4f8886102e08301525192839161165784846116e2565b8336843760085afa1590811561166f575b50610c4f57005b6001915051141581611668565b34610236575f36600319011261023657807f43f36fcf45c2220e6c29360fc8ac79c70e42d0a597b3130d35a8e826a1b795d260209252f35b9181601f840112156102365782359167ffffffffffffffff8311610236576020838186019501011161023657565b90601f8019910116810190811067ffffffffffffffff82111761024a57604052565b909160409282526020820152016060516080905f5b8181106117265750505090565b8251845260209384019390920191600101611719565b905f905b6002821061174d57505050565b6020806001928551815201930191019091611740565b80156117c7578060011c915f516020611e045f395f51905f52831015610c4f576001806117a65f516020611e045f395f51905f5260038188818180090908611c24565b9316146117af57565b905f516020611e045f395f51905f5280910681030690565b505f905f90565b8015806118ff575b6118f3578060021c92825f516020611e045f395f51905f5285108015906118dc575b610c4f5784815f516020611e045f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd44816118a69d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611c47565b809291600180829616146118b8575050565b5f516020611e045f395f51905f528093945080929550809106810306930681030690565b505f516020611e045f395f51905f528110156117f8565b50505f905f905f905f90565b5081156117d6565b905f516020611e045f395f51905f528210801590611991575b610c4f57811580611989575b611983576119505f516020611e045f395f51905f5260038185818180090908611c24565b81810361195f57505060011b90565b5f516020611e045f395f51905f52809106810306145f14610c4f57600190811b1790565b50505f90565b50801561192c565b505f516020611e045f395f51905f52811015611920565b919093925f516020611e045f395f51905f528310801590611bc5575b8015611bae575b8015611b97575b610c4f578082868517171715611b8c57908291611aef5f516020611e045f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611e045f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611ac981808b80098187800908611c24565b8408095f516020611e045f395f51905f52611ae382611d9b565b80091415958691611c47565b929080821480611b83575b15611b215750505050905f14611b195760ff60025b169060021b179190565b60ff5f611b0f565b5f516020611e045f395f51905f52809106810306149182611b64575b505015610c4f5760019115611b5c5760ff60025b169060021b17179190565b60ff5f611b51565b5f516020611e045f395f51905f52919250819006810306145f80611b3d565b50838314611afa565b50505090505f905f90565b505f516020611e045f395f51905f528110156119d2565b505f516020611e045f395f51905f528210156119cb565b505f516020611e045f395f51905f528510156119c4565b9080601f8301121561023657604080519290611bf890846116e2565b82906040810192831161023657905b828210611c145750505090565b8135815260209182019101611c07565b90611c2e82611d9b565b915f516020611e045f395f51905f5283800903610c4f57565b915f516020611e045f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c9f93969496611c9182808a8009818a800908611c24565b90611d8f575b860809611c24565b925f516020611e045f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611e045f395f51905f5260a083015260208260c08160055afa91519115610c4f575f516020611e045f395f51905f52826001920903610c4f575f516020611e045f395f51905f52908209925f516020611e045f395f51905f528080808780090681030681878009081490811591611d70575b50610c4f57565b90505f516020611e045f395f51905f528084860960020914155f611d69565b81809106810306611c97565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611e045f395f51905f5260a083015260208260c08160055afa91519115610c4f5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220fe601195f61ca6989f1a8920e4016d945151116ba070a59df21881ed6792f91a64736f6c634300081c0033",
}

// ResultsVerifierGroth16ABI is the input ABI used to generate the binding from.
// Deprecated: Use ResultsVerifierGroth16MetaData.ABI instead.
var ResultsVerifierGroth16ABI = ResultsVerifierGroth16MetaData.ABI

// ResultsVerifierGroth16Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ResultsVerifierGroth16MetaData.Bin instead.
var ResultsVerifierGroth16Bin = ResultsVerifierGroth16MetaData.Bin

// DeployResultsVerifierGroth16 deploys a new Ethereum contract, binding an instance of ResultsVerifierGroth16 to it.
func DeployResultsVerifierGroth16(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ResultsVerifierGroth16, error) {
	parsed, err := ResultsVerifierGroth16MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ResultsVerifierGroth16Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ResultsVerifierGroth16{ResultsVerifierGroth16Caller: ResultsVerifierGroth16Caller{contract: contract}, ResultsVerifierGroth16Transactor: ResultsVerifierGroth16Transactor{contract: contract}, ResultsVerifierGroth16Filterer: ResultsVerifierGroth16Filterer{contract: contract}}, nil
}

// ResultsVerifierGroth16 is an auto generated Go binding around an Ethereum contract.
type ResultsVerifierGroth16 struct {
	ResultsVerifierGroth16Caller     // Read-only binding to the contract
	ResultsVerifierGroth16Transactor // Write-only binding to the contract
	ResultsVerifierGroth16Filterer   // Log filterer for contract events
}

// ResultsVerifierGroth16Caller is an auto generated read-only Go binding around an Ethereum contract.
type ResultsVerifierGroth16Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResultsVerifierGroth16Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ResultsVerifierGroth16Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResultsVerifierGroth16Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResultsVerifierGroth16Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResultsVerifierGroth16Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResultsVerifierGroth16Session struct {
	Contract     *ResultsVerifierGroth16 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ResultsVerifierGroth16CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResultsVerifierGroth16CallerSession struct {
	Contract *ResultsVerifierGroth16Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ResultsVerifierGroth16TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResultsVerifierGroth16TransactorSession struct {
	Contract     *ResultsVerifierGroth16Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ResultsVerifierGroth16Raw is an auto generated low-level Go binding around an Ethereum contract.
type ResultsVerifierGroth16Raw struct {
	Contract *ResultsVerifierGroth16 // Generic contract binding to access the raw methods on
}

// ResultsVerifierGroth16CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResultsVerifierGroth16CallerRaw struct {
	Contract *ResultsVerifierGroth16Caller // Generic read-only contract binding to access the raw methods on
}

// ResultsVerifierGroth16TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResultsVerifierGroth16TransactorRaw struct {
	Contract *ResultsVerifierGroth16Transactor // Generic write-only contract binding to access the raw methods on
}

// NewResultsVerifierGroth16 creates a new instance of ResultsVerifierGroth16, bound to a specific deployed contract.
func NewResultsVerifierGroth16(address common.Address, backend bind.ContractBackend) (*ResultsVerifierGroth16, error) {
	contract, err := bindResultsVerifierGroth16(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResultsVerifierGroth16{ResultsVerifierGroth16Caller: ResultsVerifierGroth16Caller{contract: contract}, ResultsVerifierGroth16Transactor: ResultsVerifierGroth16Transactor{contract: contract}, ResultsVerifierGroth16Filterer: ResultsVerifierGroth16Filterer{contract: contract}}, nil
}

// NewResultsVerifierGroth16Caller creates a new read-only instance of ResultsVerifierGroth16, bound to a specific deployed contract.
func NewResultsVerifierGroth16Caller(address common.Address, caller bind.ContractCaller) (*ResultsVerifierGroth16Caller, error) {
	contract, err := bindResultsVerifierGroth16(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResultsVerifierGroth16Caller{contract: contract}, nil
}

// NewResultsVerifierGroth16Transactor creates a new write-only instance of ResultsVerifierGroth16, bound to a specific deployed contract.
func NewResultsVerifierGroth16Transactor(address common.Address, transactor bind.ContractTransactor) (*ResultsVerifierGroth16Transactor, error) {
	contract, err := bindResultsVerifierGroth16(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResultsVerifierGroth16Transactor{contract: contract}, nil
}

// NewResultsVerifierGroth16Filterer creates a new log filterer instance of ResultsVerifierGroth16, bound to a specific deployed contract.
func NewResultsVerifierGroth16Filterer(address common.Address, filterer bind.ContractFilterer) (*ResultsVerifierGroth16Filterer, error) {
	contract, err := bindResultsVerifierGroth16(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResultsVerifierGroth16Filterer{contract: contract}, nil
}

// bindResultsVerifierGroth16 binds a generic wrapper to an already deployed contract.
func bindResultsVerifierGroth16(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResultsVerifierGroth16MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResultsVerifierGroth16.Contract.ResultsVerifierGroth16Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResultsVerifierGroth16.Contract.ResultsVerifierGroth16Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResultsVerifierGroth16.Contract.ResultsVerifierGroth16Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResultsVerifierGroth16.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResultsVerifierGroth16.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResultsVerifierGroth16.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "compressProof", proof, commitments, commitmentPok)

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
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) CompressProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _ResultsVerifierGroth16.Contract.CompressProof(&_ResultsVerifierGroth16.CallOpts, proof, commitments, commitmentPok)
}

// CompressProof is a free data retrieval call binding the contract method 0xb1c3a00e.
//
// Solidity: function compressProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok) view returns(uint256[4] compressed, uint256[1] compressedCommitments, uint256 compressedCommitmentPok)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) CompressProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int) (struct {
	Compressed              [4]*big.Int
	CompressedCommitments   [1]*big.Int
	CompressedCommitmentPok *big.Int
}, error) {
	return _ResultsVerifierGroth16.Contract.CompressProof(&_ResultsVerifierGroth16.CallOpts, proof, commitments, commitmentPok)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) ProvingKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "provingKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) ProvingKeyHash() ([32]byte, error) {
	return _ResultsVerifierGroth16.Contract.ProvingKeyHash(&_ResultsVerifierGroth16.CallOpts)
}

// ProvingKeyHash is a free data retrieval call binding the contract method 0x233ace11.
//
// Solidity: function provingKeyHash() pure returns(bytes32)
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) ProvingKeyHash() ([32]byte, error) {
	return _ResultsVerifierGroth16.Contract.ProvingKeyHash(&_ResultsVerifierGroth16.CallOpts)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5d26278e.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, compressedCommitments, compressedCommitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5d26278e.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyCompressedProof(&_ResultsVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0x5d26278e.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyCompressedProof(&_ResultsVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [9]*big.Int) error {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "verifyProof", proof, commitments, commitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof(&_ResultsVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x60e58346.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[9] input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [9]*big.Int) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof(&_ResultsVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Caller) VerifyProof0(opts *bind.CallOpts, _proof []byte, _input []byte) error {
	var out []interface{}
	err := _ResultsVerifierGroth16.contract.Call(opts, &out, "verifyProof0", _proof, _input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16Session) VerifyProof0(_proof []byte, _input []byte) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof0(&_ResultsVerifierGroth16.CallOpts, _proof, _input)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0xb8e72af6.
//
// Solidity: function verifyProof(bytes _proof, bytes _input) view returns()
func (_ResultsVerifierGroth16 *ResultsVerifierGroth16CallerSession) VerifyProof0(_proof []byte, _input []byte) error {
	return _ResultsVerifierGroth16.Contract.VerifyProof0(&_ResultsVerifierGroth16.CallOpts, _proof, _input)
}
