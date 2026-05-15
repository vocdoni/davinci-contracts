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
	ABI: "[{\"inputs\":[],\"name\":\"CommitmentInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provingKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[1]\",\"name\":\"compressedCommitments\",\"type\":\"uint256[1]\"},{\"internalType\":\"uint256\",\"name\":\"compressedCommitmentPok\",\"type\":\"uint256\"},{\"internalType\":\"uint256[8]\",\"name\":\"input\",\"type\":\"uint256[8]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitments\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"commitmentPok\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[8]\",\"name\":\"input\",\"type\":\"uint256[8]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60808060405234601557611dfe908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c908163233ace111461157b575080639ed57f3214610cb9578063b1c3a00e14610ba8578063b8e72af614610a315763e2dd9f4714610055575f80fd5b34610a2e576101c0366003190112610a2e5736608411610a2e573660a411610a2e57366101c411610a2e57604051602061008f81836115e1565b803683376040918251926100a381856115e1565b80368537610300938151926100b886856115e1565b853685376100c7608435611a05565b8684015282526100d860a435611a05565b905f516020611da95f395f51905f5285516100f387826115e1565b6001815288610104818301601f198a01368237378551906101328a610124818a0151938b5194859384019687611603565b03601f1981018352826115e1565b5190200683528351865286840151878701527f0399d309b609c308d0b5b44bf26f386038097e7f6f2cc7c042178e114bae7532858701527f076e7df5029c8dbffcb8d5c61d6cb1127579a160dcda2c8ae6ac5c964e9714c160608701527f058b5cefcb2d7bb8a441fc99b927c804ddb6009f8816461f2f368ea650b23d5860808701527f0af41e75d7c8afbccb2d911a04213523375d130da858649d3ab787812f1befc260a087015260c086015260e08501527f03f9f68a4e152f4c6b52317822146b7ba88ce42a45c69766a118e4934e2727426101008501527f2c3f776d51b23ad2c12fc91e716cb3c20b56e3fcf89846d8620f38e2e046623e6101208501527f0475715d9ecec423ed21fc4f389074e3dd8559fe4bbaae2bc947773eee8d9f786101408501527f25a7d84b2115594342d64b0b1adb9f3bfde6916849e334ce50a6b61ef91ef936610160850152825185816101808760085afa90511615610a1f578490876102a3600435611a05565b866102b5602497939735604435611a70565b9491929390956102c6606435611a05565b999098507f2da52ba7c21ad6c8264523cf92aad2997eec89fbe65cb8add30f68812dc5757d83519b8c937f2defa9cca267e87d3aa0ed81ad0d40adc02f82c496bd566051f8011412b38efe85527f0e34574cd477146db5f8a3ed17476ae7bcb52fd3b1382dd5306dd1dc73d9d8be828601528051868601520151606084019081527f05b601470063e116b1360cad9c1aa6326a3d82d2748ec616266f541b813defb3858560808160065afa957f29e7a40cf04aeea37655d03140a8fb30a7d5c9fbe15f374bafb6ba48c3633fb2818701527f20a913527d48e88caa5fbd0afa82f7c7800e440ddab59974c6846dcb5e156a848352600160c4356080880198818a525f516020611da95f395f51905f5284808b016060828d0160075afa92101616828860808160065afa16167f254a65aa1f32410d1a48a4de84223fc880cbf3549b4a39d6e78897e3765b3638828801527f15ec0dad4440fe69f463429d37a164e7703c6281ba4be271c5f8eadabcbf0f41845260e435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f23710f807c4d32ddf5ba2fa04bf9819324e4a193317adc39ddd80313826123d9828801527f123cddba2edc5ca0190a547ce1d4df7498a4d00367f2d20d90626cc119c38037845261010435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f2c3cca18074f3e04179c2c3cba3e46c82f8fa6548a04cd66daf6abbb26cd39bf828801527f1d329b5922904931f14acf061586570a091afb622bf0e112c757391afbe40892845261012435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f091eff56ec0fafad6068156070dbaaa55a0360687ff3e544de82809233c7fbba828801527f25d8557db85219b665cacdbf9c25b437420a2e5990e3305c48901a52b3903ed8845261014435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f1fa6a81909a65fd813f283980b2ecc9ef08c4bb8ab0d124fc3410323d594f369828801527f0e3cd7a8df0ad5a59b2426d2dfcc4e683412b0a61a7b0cbf57a12039f224ea74845261016435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167e1739da6fd50ebf29ea9b6e69c8193b63b0c2de3b88dffc5edc4e0cd1bb98a7828801527f1f39424f91a05caa74f1ae6dbe454ac9b91edd627456fb6c739482f2b8362460845261018435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa167f188ac5dae9bfad73c137d2ee1077ff8b1268bb06914b8b914708f6d8273e4828828801527f038e8fbeae99b34d8587e1ae59874491f1a80c09f74a68238520b42d6c80685484526101a435908189525f516020611da95f395f51905f5283808a016060828c0160075afa92101616818760808160065afa1695015252518092525f516020611da95f395f51905f528c8b606082808301920160075afa921016168a8960808160065afa16988c89519901519915610a10578b528b8b0152888a01526060890152608088015260a087015260c086015260e08501527f1a455a8b4a7f7bd3090154dc80432b771edab1130feba2997070cdf9612f86db6101008501527f1f23f4849650096b6e7bcabbe9e8caf2257eabe3e0dec2c8328467061a3dbf906101208501527f05860d0c5f0fb773831ae7a34d7729c8797097c3786432193cb14574c1df00c56101408501527f05e307545f3062695d5291e83b9e93c115500b056a2b0ed4c4f7860c52b849fe6101608501527f2f437ec49e88bf48b5b7e2edb741c2917c182be53c169f6ab74bc6d9a91656106101808501527f30365c4a0575fb12d2b103124f3083cf672047fb873a58d8d0038dbcab6bc2926101a08501527f2401e81e1227e211abc40ca339b4c8c6c2b7a65ebca627123db2c0ced2671bb16101c08501527f1a9fd719682ac62bbfca44d3e13c366b5525dcf2c4933ca0df8fd4ef0d1c0e326101e08501527f2280b35aefb0712bc3a165658e646aa3c275aeb7cdc5b0b7c0ca97d6d621dae36102008501527f0662f988ae72122abb73eff5aaacacda1197228f5e3863db638d0590c6b17ea56102208501526102408401526102608301527f2ff4777ef2f0d3aed1eb13cce8f901d981d1f0083e22459cecf3f25048be082f6102808301527ea988b85883362da08ae1f26cd329ffa0865a6da4c8024cd1c3efea90436bb46102a08301527f1be6a06698df8fc509c35a3763e68ac9ae1b609eabb5dc61830a80fdfa624ad36102c08301527f0c58811ba25d4fd2c48c88db9eb31cfb68355a79e1a53583e47244cb5b4a84f06102e0830152519283916109db84846115e1565b8336843760085afa15908115610a03575b506109f45780f35b631ff3747d60e21b8152600490fd5b600191505114155f6109ec565b63a54f8e2760e01b8f5260048ffd5b6351d49ff760e11b8752600487fd5b80fd5b5034610ba4576040366003190112610ba45760043567ffffffffffffffff8111610ba457610a639036906004016115b3565b60243567ffffffffffffffff8111610ba457610a839036906004016115b3565b90926101009384604051610a9782826115e1565b369037604093848051610aaa82826115e1565b369037848051610aba82826115e1565b36903782019461018083870312610ba457610aef610ad88785611960565b96610140610ae8828588016119a9565b95016119a9565b93818651610afd82826115e1565b36903782019082820312610ba457610b1491611960565b90610b1d6119f1565b50610b266119f1565b50303b15610ba457610b7492610b5e610b6992610b53875198634f6abf9960e11b8a5260048a019061163d565b610104880190611664565b610144860190611664565b61018484019061163d565b5f8261028481305afa908115610b9b5750610b8d575080f35b610b9991505f906115e1565b005b513d5f823e3d90fd5b5f80fd5b34610ba457610180366003190112610ba4573661010411610ba4573661014411610ba4573661018411610ba457604051610be36080826115e1565b6080368237604051906020610bf881846115e1565b80368437610c0a60243560043561168b565b8252610c2060843560a43560443560643561172c565b828401526040830152610c3760e43560c43561168b565b6060830152610c4c610124356101043561168b565b8352610c5e610164356101443561168b565b9060405192835f905b60048210610ca3575050506080830193905f945b60018610610c8e5760c0858560a0820152f35b81806001928551815201930195019491610c7b565b8251815291830191600191909101908301610c67565b34610ba457610280366003190112610ba4573661010411610ba4573661014411610ba4573661018411610ba4573661028411610ba4576040516020610cfe81836115e1565b803683376040915f516020611da95f395f51905f528351610d1f85826115e1565b60018152836101c4818301601f198801368237378451610d4e8161012487820194610124356101043587611603565b51902006815282518361010482377f0399d309b609c308d0b5b44bf26f386038097e7f6f2cc7c042178e114bae7532848201527f076e7df5029c8dbffcb8d5c61d6cb1127579a160dcda2c8ae6ac5c964e9714c160608201527f058b5cefcb2d7bb8a441fc99b927c804ddb6009f8816461f2f368ea650b23d5860808201527f0af41e75d7c8afbccb2d911a04213523375d130da858649d3ab787812f1befc260a08201528361014460c08301377f03f9f68a4e152f4c6b52317822146b7ba88ce42a45c69766a118e4934e2727426101008201527f2c3f776d51b23ad2c12fc91e716cb3c20b56e3fcf89846d8620f38e2e046623e6101208201527f0475715d9ecec423ed21fc4f389074e3dd8559fe4bbaae2bc947773eee8d9f786101408201527f25a7d84b2115594342d64b0b1adb9f3bfde6916849e334ce50a6b61ef91ef93661016082015282816101808160085afa9051161561156c57825192610eb86040856115e1565b61010483855b6101448310611548575050508051918183017f2defa9cca267e87d3aa0ed81ad0d40adc02f82c496bd566051f8011412b38efe84525f516020611da95f395f51905f528386808701987f0e34574cd477146db5f8a3ed17476ae7bcb52fd3b1382dd5306dd1dc73d9d8be8a528051855201519260608701938452818760808160065afa947f29e7a40cf04aeea37655d03140a8fb30a7d5c9fbe15f374bafb6ba48c3633fb282527f20a913527d48e88caa5fbd0afa82f7c7800e440ddab59974c6846dcb5e156a8485527f2da52ba7c21ad6c8264523cf92aad2997eec89fbe65cb8add30f68812dc5757d60016101843560808b0198818a5287878760608160075afa92101616858b60808160065afa16167f254a65aa1f32410d1a48a4de84223fc880cbf3549b4a39d6e78897e3765b363884527f15ec0dad4440fe69f463429d37a164e7703c6281ba4be271c5f8eadabcbf0f4187526101a4359081895286868660608160075afa92101616848a60808160065afa167f23710f807c4d32ddf5ba2fa04bf9819324e4a193317adc39ddd80313826123d984527f123cddba2edc5ca0190a547ce1d4df7498a4d00367f2d20d90626cc119c3803787526101c4359081895286868660608160075afa92101616848a60808160065afa167f2c3cca18074f3e04179c2c3cba3e46c82f8fa6548a04cd66daf6abbb26cd39bf84527f1d329b5922904931f14acf061586570a091afb622bf0e112c757391afbe4089287526101e4359081895286868660608160075afa92101616848a60808160065afa167f091eff56ec0fafad6068156070dbaaa55a0360687ff3e544de82809233c7fbba84527f25d8557db85219b665cacdbf9c25b437420a2e5990e3305c48901a52b3903ed88752610204359081895286868660608160075afa92101616848a60808160065afa167f1fa6a81909a65fd813f283980b2ecc9ef08c4bb8ab0d124fc3410323d594f36984527f0e3cd7a8df0ad5a59b2426d2dfcc4e683412b0a61a7b0cbf57a12039f224ea748752610224359081895286868660608160075afa92101616848a60808160065afa167e1739da6fd50ebf29ea9b6e69c8193b63b0c2de3b88dffc5edc4e0cd1bb98a784527f1f39424f91a05caa74f1ae6dbe454ac9b91edd627456fb6c739482f2b83624608752610244359081895286868660608160075afa92101616848a60808160065afa167f188ac5dae9bfad73c137d2ee1077ff8b1268bb06914b8b914708f6d8273e482884527f038e8fbeae99b34d8587e1ae59874491f1a80c09f74a68238520b42d6c8068548752610264359081895286868660608160075afa92101616848a60808160065afa16957f05b601470063e116b1360cad9c1aa6326a3d82d2748ec616266f541b813defb38452525180955260608160075afa92101616818360808160065afa16915193519115611539575192610100600485377f1a455a8b4a7f7bd3090154dc80432b771edab1130feba2997070cdf9612f86db6101008501527f1f23f4849650096b6e7bcabbe9e8caf2257eabe3e0dec2c8328467061a3dbf906101208501527f05860d0c5f0fb773831ae7a34d7729c8797097c3786432193cb14574c1df00c56101408501527f05e307545f3062695d5291e83b9e93c115500b056a2b0ed4c4f7860c52b849fe6101608501527f2f437ec49e88bf48b5b7e2edb741c2917c182be53c169f6ab74bc6d9a91656106101808501527f30365c4a0575fb12d2b103124f3083cf672047fb873a58d8d0038dbcab6bc2926101a08501527f2401e81e1227e211abc40ca339b4c8c6c2b7a65ebca627123db2c0ced2671bb16101c08501527f1a9fd719682ac62bbfca44d3e13c366b5525dcf2c4933ca0df8fd4ef0d1c0e326101e08501527f2280b35aefb0712bc3a165658e646aa3c275aeb7cdc5b0b7c0ca97d6d621dae36102008501527f0662f988ae72122abb73eff5aaacacda1197228f5e3863db638d0590c6b17ea56102208501526102408401526102608301527f2ff4777ef2f0d3aed1eb13cce8f901d981d1f0083e22459cecf3f25048be082f6102808301527ea988b85883362da08ae1f26cd329ffa0865a6da4c8024cd1c3efea90436bb46102a08301527f1be6a06698df8fc509c35a3763e68ac9ae1b609eabb5dc61830a80fdfa624ad36102c08301527f0c58811ba25d4fd2c48c88db9eb31cfb68355a79e1a53583e47244cb5b4a84f06102e0830152816103008160085afa9051161561152a57005b631ff3747d60e21b5f5260045ffd5b63a54f8e2760e01b5f5260045ffd5b8235815291810191859101610ebe565b634e487b7160e01b5f52604160045260245ffd5b6351d49ff760e11b5f5260045ffd5b34610ba4575f366003190112610ba457807fd8d1400e7922b6eb7652caf13546156b915a9565cfff6130365be103bedd6e6860209252f35b9181601f84011215610ba45782359167ffffffffffffffff8311610ba45760208381860195010111610ba457565b90601f8019910116810190811067ffffffffffffffff82111761155857604052565b9091604092825260208201520160208251919201905f5b8181106116275750505090565b825184526020938401939092019160010161161a565b905f905b6008821061164e57505050565b6020806001928551815201930191019091611641565b905f905b6002821061167557505050565b6020806001928551815201930191019091611668565b905f516020611d895f395f51905f528210801590611715575b61152a5781158061170d575b611707576116d45f516020611d895f395f51905f5260038185818180090908611ba9565b8181036116e357505060011b90565b5f516020611d895f395f51905f52809106810306145f1461152a57600190811b1790565b50505f90565b5080156116b0565b505f516020611d895f395f51905f528110156116a4565b919093925f516020611d895f395f51905f528310801590611949575b8015611932575b801561191b575b61152a578082868517171715611910579082916118735f516020611d895f395f51905f5280808080888180808f9d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd448f839290839109099d8e0981848181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089a09818c8181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306945f516020611d895f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea48161184d81808b80098187800908611ba9565b8408095f516020611d895f395f51905f5261186782611d20565b80091415958691611bcc565b929080821480611907575b156118a55750505050905f1461189d5760ff60025b169060021b179190565b60ff5f611893565b5f516020611d895f395f51905f528091068103061491826118e8575b50501561152a57600191156118e05760ff60025b169060021b17179190565b60ff5f6118d5565b5f516020611d895f395f51905f52919250819006810306145f806118c1565b5083831461187e565b50505090505f905f90565b505f516020611d895f395f51905f52811015611756565b505f516020611d895f395f51905f5282101561174f565b505f516020611d895f395f51905f52851015611748565b9080601f83011215610ba4576040519161197c610100846115e1565b82906101008101928311610ba457905b8282106119995750505090565b813581526020918201910161198c565b9080601f83011215610ba4576040805192906119c590846115e1565b829060408101928311610ba457905b8282106119e15750505090565b81358152602091820191016119d4565b60405190611a006020836115e1565b5f8252565b8015611a69578060011c915f516020611d895f395f51905f5283101561152a57600180611a485f516020611d895f395f51905f5260038188818180090908611ba9565b931614611a5157565b905f516020611d895f395f51905f5280910681030690565b505f905f90565b801580611ba1575b611b95578060021c92825f516020611d895f395f51905f528510801590611b7e575b61152a5784815f516020611d895f395f51905f5280808080808080807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4481611b489d8d0909998a0981898181800909087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750806810306936002808a16149509818a8181800909087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508611bcc565b80929160018082961614611b5a575050565b5f516020611d895f395f51905f528093945080929550809106810306930681030690565b505f516020611d895f395f51905f52811015611a9a565b50505f905f905f905f90565b508115611a78565b90611bb382611d20565b915f516020611d895f395f51905f528380090361152a57565b915f516020611d895f395f51905f527f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea481611c2493969496611c1682808a8009818a800908611ba9565b90611d14575b860809611ba9565b925f516020611d895f395f51905f52600285096040519060208252602080830152602060408301528060608301527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4560808301525f516020611d895f395f51905f5260a083015260208260c08160055afa9151911561152a575f516020611d895f395f51905f5282600192090361152a575f516020611d895f395f51905f52908209925f516020611d895f395f51905f528080808780090681030681878009081490811591611cf5575b5061152a57565b90505f516020611d895f395f51905f528084860960020914155f611cee565b81809106810306611c1c565b9060405191602083526020808401526020604084015260608301527f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260808301525f516020611d895f395f51905f5260a083015260208260c08160055afa9151911561152a5756fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4730644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001a2646970667358221220332dd4666c7c35088eda0f26dd6bd52d352b4ac95bef09a17b5a1f898cf7154164736f6c634300081c0033",
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe2dd9f47.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, compressedCommitments, compressedCommitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe2dd9f47.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [8]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyCompressedProof(&_StateTransitionVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xe2dd9f47.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[1] compressedCommitments, uint256 compressedCommitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, compressedCommitments [1]*big.Int, compressedCommitmentPok *big.Int, input [8]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyCompressedProof(&_StateTransitionVerifierGroth16.CallOpts, compressedProof, compressedCommitments, compressedCommitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x9ed57f32.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Caller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _StateTransitionVerifierGroth16.contract.Call(opts, &out, "verifyProof", proof, commitments, commitmentPok, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x9ed57f32.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16Session) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [8]*big.Int) error {
	return _StateTransitionVerifierGroth16.Contract.VerifyProof(&_StateTransitionVerifierGroth16.CallOpts, proof, commitments, commitmentPok, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x9ed57f32.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] commitments, uint256[2] commitmentPok, uint256[8] input) view returns()
func (_StateTransitionVerifierGroth16 *StateTransitionVerifierGroth16CallerSession) VerifyProof(proof [8]*big.Int, commitments [2]*big.Int, commitmentPok [2]*big.Int, input [8]*big.Int) error {
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
