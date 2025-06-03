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

// IProcessRegistryBallotMode is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryBallotMode struct {
	CostFromWeight  bool
	ForceUniqueness bool
	MaxCount        uint8
	CostExponent    uint8
	MaxValue        *big.Int
	MinValue        *big.Int
	MaxTotalCost    *big.Int
	MinTotalCost    *big.Int
}

// IProcessRegistryCensus is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryCensus struct {
	CensusOrigin uint8
	MaxVotes     *big.Int
	CensusRoot   [32]byte
	CensusURI    string
}

// IProcessRegistryEncryptionKey is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryEncryptionKey struct {
	X *big.Int
	Y *big.Int
}

// IProcessRegistryProcess is an auto generated low-level Go binding around an user-defined struct.
type IProcessRegistryProcess struct {
	Status             uint8
	OrganizationId     common.Address
	EncryptionKey      IProcessRegistryEncryptionKey
	LatestStateRoot    *big.Int
	Result             []*big.Int
	StartTime          *big.Int
	Duration           *big.Int
	VoteCount          *big.Int
	VoteOverwriteCount *big.Int
	MetadataURI        string
	BallotMode         IProcessRegistryBallotMode
	Census             IProcessRegistryCensus
}

// ProcessRegistryMetaData contains all meta data concerning the ProcessRegistry contract.
var ProcessRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_chainID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_organizationRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidCensus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStateRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOrganizationAdministrator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrganizationNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"}],\"name\":\"CensusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProcessCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProcessDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"name\":\"ProcessResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"ProcessStateRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ProcessStatusChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CENSUS_ORIGIN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STATUS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"}],\"name\":\"getProcess\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteOverwriteCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forceUniqueness\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"maxCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"internalType\":\"structIProcessRegistry.Process\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSTVerifierVKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forceUniqueness\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"maxCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"initStateRoot\",\"type\":\"uint256\"}],\"name\":\"newProcess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"organizationRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"processes\",\"outputs\":[{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"organizationId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.EncryptionKey\",\"name\":\"encryptionKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"latestStateRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteOverwriteCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"costFromWeight\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forceUniqueness\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"maxCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"costExponent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalCost\",\"type\":\"uint256\"}],\"internalType\":\"structIProcessRegistry.BallotMode\",\"name\":\"ballotMode\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"enumIProcessRegistry.CensusOrigin\",\"name\":\"censusOrigin\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxVotes\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIProcessRegistry.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"setProcessCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setProcessDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setProcessResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIProcessRegistry.ProcessStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"setProcessStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"processId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"submitStateTransition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161258938038061258983398101604081905261002f916100b1565b600261003b8582610238565b50600180546001600160a01b039485166001600160a01b031991821617909155600380549385169382169390931790925560048054919093169116179055506102f6565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b03811681146100ac57600080fd5b919050565b600080600080608085870312156100c757600080fd5b84516001600160401b038111156100dd57600080fd5b8501601f810187136100ee57600080fd5b80516001600160401b038111156101075761010761007f565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101355761013561007f565b60405281815282820160200189101561014d57600080fd5b60005b8281101561016c57602081850181015183830182015201610150565b5060006020838301015280965050505061018860208601610095565b925061019660408601610095565b91506101a460608601610095565b905092959194509250565b600181811c908216806101c357607f821691505b6020821081036101e357634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561023357806000526020600020601f840160051c810160208510156102105750805b601f840160051c820191505b81811015610230576000815560010161021c565b50505b505050565b81516001600160401b038111156102515761025161007f565b6102658161025f84546101af565b846101e9565b6020601f82116001811461029957600083156102815750848201515b600019600385901b1c1916600184901b178455610230565b600084815260208120601f198516915b828110156102c957878501518255602094850194600190920191016102a9565b50848210156102e75786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b612284806103056000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806392bc2669116100a2578063adc879e911610071578063adc879e91461024f578063b79938e914610264578063c718c01f14610277578063d71d072e1461028a578063f9aa44991461029d57600080fd5b806392bc2669146101f6578063992bc45b146102095780639c18deaf14610229578063a5b110de1461023c57600080fd5b806356a6f1e2116100de57806356a6f1e21461019a5780636bae04ea146101af57806372c628ef146101c2578063848df540146101ca57600080fd5b806302cc919a146101105780630535fece1461012f57806309706392146101595780634c0acc5614610184575b600080fd5b610118600981565b60405160ff90911681526020015b60405180910390f35b61014261013d3660046116d5565b6102a5565b6040516101269b9a9998979695949392919061180b565b60035461016c906001600160a01b031681565b6040516001600160a01b039091168152602001610126565b61018c6104e7565b604051908152602001610126565b6101ad6101a83660046118ba565b61055a565b005b6101ad6101bd3660046118fe565b610729565b610118600481565b6001546101e190600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610126565b6101ad6102043660046119ca565b610956565b61021c6102173660046116d5565b610cb1565b6040516101269190611aed565b60045461016c906001600160a01b031681565b6101ad61024a366004611be1565b610fb4565b610257611189565b6040516101269190611c60565b60015461016c906001600160a01b031681565b6101ad610285366004611c7a565b611217565b6101ad610298366004611be1565b611397565b61018c61152e565b6000602081815291815260409081902080548251808401909352600182015483526002820154938301939093526003810154600582015460068301546007840154600885015460098601805460ff8a169961010090046001600160a01b03169897919061031190611c9c565b80601f016020809104026020016040519081016040528092919081815260200182805461033d90611c9c565b801561038a5780601f1061035f5761010080835404028352916020019161038a565b820191906000526020600020905b81548152906001019060200180831161036d57829003601f168201915b5050604080516101008082018352600a88015460ff808216151584529181048216151560208401526201000081048216838501526301000000900481166060830152600b880154608080840191909152600c89015460a0840152600d89015460c0840152600e89015460e08401528351908101909352600f88018054979892979296509294509192508391166009811115610427576104276116ee565b6009811115610438576104386116ee565b8152602001600182015481526020016002820154815260200160038201805461046090611c9c565b80601f016020809104026020016040519081016040528092919081815260200182805461048c90611c9c565b80156104d95780601f106104ae576101008083540402835291602001916104d9565b820191906000526020600020905b8154815290600101906020018083116104bc57829003601f168201915b50505050508152505090508b565b6003546040805163233ace1160e01b815290516000926001600160a01b03169163233ace119160048083019260209291908290030181865afa158015610531573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105559190611cd0565b905090565b6000828152602081905260409020805461010090046001600160a01b03168061059657604051634d36eb6960e01b815260040160405180910390fd5b815460ff168060048111156105ad576105ad6116ee565b8460048111156105bf576105bf6116ee565b14806105de5750600484818111156105d9576105d96116ee565b60ff16115b80610618575060008160048111156105f8576105f86116ee565b1415801561061857506003816004811115610615576106156116ee565b14155b15610636576040516307a92f1960e51b815260040160405180910390fd5b60015460405163c1af6e0360e01b81526001600160a01b0384811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa158015610686573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106aa9190611cfa565b6106c75760405163f8b857d360e01b815260040160405180910390fd5b82548490849060ff191660018360048111156106e5576106e56116ee565b0217905550847fac0c4085a30bc70d0d8893ee5d6466ac9c5f03e27fd7292dcef128a610e7c1908560405161071a9190611d17565b60405180910390a25050505050565b6107366060820182611d2b565b90506000036107585760405163111127cb60e01b815260040160405180910390fd5b604081013560000361077d5760405163111127cb60e01b815260040160405180910390fd5b6000828152602081905260409020805461010090046001600160a01b0316806107b957604051634d36eb6960e01b815260040160405180910390fd5b60015460405163c1af6e0360e01b81526001600160a01b0383811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa158015610809573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061082d9190611cfa565b61084a5760405163f8b857d360e01b815260040160405180910390fd5b6010820154602084013510156108735760405163111127cb60e01b815260040160405180910390fd5b6000825460ff16600481111561088b5761088b6116ee565b141580156108af57506003825460ff1660048111156108ac576108ac6116ee565b14155b156108cd576040516307a92f1960e51b815260040160405180910390fd5b60208301356010830155604083013560118301556108ee6060840184611d2b565b60128401916108fe919083611dd7565b50837f35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b60408501356109336060870187611d2b565b87602001356040516109489493929190611ec0565b60405180910390a250505050565b60015460405163c1af6e0360e01b81526001600160a01b0386811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa1580156109a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ca9190611cfa565b6109e75760405163f8b857d360e01b815260040160405180910390fd5b6000838152602081905260409020805461010090046001600160a01b031615610a2357604051635040c4d360e11b815260040160405180910390fd5b610a3360608a0160408b01611efa565b60ff16600003610a565760405163ac38930b60e01b815260040160405180910390fd5b610a6660608a0160408b01611efa565b60ff16896080013511610a8c5760405163b1911d8b60e01b815260040160405180910390fd5b60048c81811115610a9f57610a9f6116ee565b60ff161180610add575060008c6004811115610abd57610abd6116ee565b14158015610add575060038c6004811115610ada57610ada6116ee565b14155b15610afb576040516307a92f1960e51b815260040160405180910390fd5b4260008c9003610b0957809b505b808c1015610b2a57604051632ca4094f60e21b815260040160405180910390fd5b80610b358c8e611f2d565b11610b5357604051637616640160e01b815260040160405180910390fd5b6009610b6260208b018b611f4d565b6009811115610b7357610b736116ee565b60ff161115610b955760405163111127cb60e01b815260040160405180910390fd5b81548d90839060ff19166001836004811115610bb357610bb36116ee565b0217905550600582018c9055600682018b90558154610100600160a81b0319166101006001600160a01b0388160217825583356001830155602084013560028301556003820183905560098201610c0b888a83611dd7565b5089600a8301610c1b8282611f84565b50899050600f8301610c2d8282612046565b505060018054600160a01b900463ffffffff16906014610c4c836120de565b91906101000a81548163ffffffff021916908363ffffffff16021790555050336001600160a01b0316857fada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa60405160405180910390a350505050505050505050505050565b610cb9611583565b60008281526020819052604090819020815161018081019092528054829060ff166004811115610ceb57610ceb6116ee565b6004811115610cfc57610cfc6116ee565b8152815461010090046001600160a01b03166020808301919091526040805180820182526001850154815260028501548184015281840152600384015460608401526004840180548251818502810185019093528083526080909401939192909190830182828015610d8d57602002820191906000526020600020905b815481526020019060010190808311610d79575b5050505050815260200160058201548152602001600682015481526020016007820154815260200160088201548152602001600982018054610dce90611c9c565b80601f0160208091040260200160405190810160405280929190818152602001828054610dfa90611c9c565b8015610e475780601f10610e1c57610100808354040283529160200191610e47565b820191906000526020600020905b815481529060010190602001808311610e2a57829003601f168201915b5050509183525050604080516101008082018352600a85015460ff8082161515845291810482161515602084810191909152620100008204831684860152630100000090910482166060840152600b860154608080850191909152600c87015460a0850152600d87015460c0850152600e87015460e08501529085019290925282519182018352600f85018054939094019391928391166009811115610eef57610eef6116ee565b6009811115610f0057610f006116ee565b81526020016001820154815260200160028201548152602001600382018054610f2890611c9c565b80601f0160208091040260200160405190810160405280929190818152602001828054610f5490611c9c565b8015610fa15780601f10610f7657610100808354040283529160200191610fa1565b820191906000526020600020905b815481529060010190602001808311610f8457829003601f168201915b5050509190925250505090525092915050565b6000858152602081905260409020805461010090046001600160a01b0316610fef57604051634d36eb6960e01b815260040160405180910390fd5b60048054604051635c73957b60e11b81526001600160a01b039091169163b8e72af69161102491899189918991899101612103565b60006040518083038186803b15801561103c57600080fd5b505afa158015611050573d6000803e3d6000fd5b506000925061106491505083850185612166565b600383015481519192501461108c57604051630b6fac0360e41b815260040160405180910390fd5b600061109a600160096121e5565b67ffffffffffffffff8111156110b2576110b2611d72565b6040519080825280602002602001820160405280156110db578160200160208202803683370190505b50905060015b6009811015611130578281600981106110fc576110fc6121cf565b60200201518261110d6001846121e5565b8151811061111d5761111d6121cf565b60209081029190910101526001016110e1565b5080516111469060048501906020840190611675565b50877fa21d6aca26cc870bcd3da1cc0d0b0f22572550a7d04fac698fc08b67b92c43528260405161117791906121f8565b60405180910390a25050505050505050565b6002805461119690611c9c565b80601f01602080910402602001604051908101604052809291908181526020018280546111c290611c9c565b801561120f5780601f106111e45761010080835404028352916020019161120f565b820191906000526020600020905b8154815290600101906020018083116111f257829003601f168201915b505050505081565b42811161123757604051637616640160e01b815260040160405180910390fd5b6000828152602081905260409020805461010090046001600160a01b03168061127357604051634d36eb6960e01b815260040160405180910390fd5b60015460405163c1af6e0360e01b81526001600160a01b0383811660048301523360248301529091169063c1af6e0390604401602060405180830381865afa1580156112c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112e79190611cfa565b6113045760405163f8b857d360e01b815260040160405180910390fd5b6000825460ff16600481111561131c5761131c6116ee565b1415801561134057506003825460ff16600481111561133d5761133d6116ee565b14155b1561135e576040516307a92f1960e51b815260040160405180910390fd5b6006820183905560405183815284907f0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f4690602001610948565b6000858152602081905260409020805461010090046001600160a01b03166113d257604051634d36eb6960e01b815260040160405180910390fd5b6000815460ff1660048111156113ea576113ea6116ee565b14611408576040516307a92f1960e51b815260040160405180910390fd5b600354604051635c73957b60e11b81526001600160a01b039091169063b8e72af69061143e908890889088908890600401612103565b60006040518083038186803b15801561145657600080fd5b505afa15801561146a573d6000803e3d6000fd5b506000925061147e9150508385018561220b565b60038301548151919250146114a657604051630b6fac0360e41b815260040160405180910390fd5b6020810151600383015560408101516007830180546000906114c9908490611f2d565b909155505060608101516008830180546000906114e7908490611f2d565b909155505060208082015160405190815288917fd25cbd3fa66107efb1230dc4f49f939d7b6b57ba10ab0ec87ba39040c075cf20910160405180910390a250505050505050565b6000600460009054906101000a90046001600160a01b03166001600160a01b031663233ace116040518163ffffffff1660e01b8152600401602060405180830381865afa158015610531573d6000803e3d6000fd5b604080516101808101909152806000815260200160006001600160a01b031681526020016115c4604051806040016040528060008152602001600081525090565b81526020016000815260200160608152602001600081526020016000815260200160008152602001600081526020016060815260200161164c604051806101000160405280600015158152602001600015158152602001600060ff168152602001600060ff168152602001600081526020016000815260200160008152602001600081525090565b815260408051608081018252600080825260208281018290529282015260608082015291015290565b8280548282559060005260206000209081019282156116b0579160200282015b828111156116b0578251825591602001919060010190611695565b506116bc9291506116c0565b5090565b5b808211156116bc57600081556001016116c1565b6000602082840312156116e757600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b60058110611714576117146116ee565b9052565b6000815180845260005b8181101561173e57602081850181015186830182015201611722565b506000602082860101526020601f19601f83011685010191505092915050565b80511515825260208101511515602083015260ff60408201511660408301526060810151611791606084018260ff169052565b506080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b60008151600a81106117d2576117d26116ee565b8084525060208201516020840152604082015160408401526060820151608060608501526118036080850182611718565b949350505050565b611815818d611704565b6001600160a01b038b16602082015261183b604082018b80518252602090810151910152565b8860808201528760a08201528660c08201528560e0820152846101008201526102606101208201526000611873610260830186611718565b61188161014084018661175e565b82810361024084015261189481856117be565b9e9d5050505050505050505050505050565b8035600581106118b557600080fd5b919050565b600080604083850312156118cd57600080fd5b823591506118dd602084016118a6565b90509250929050565b6000608082840312156118f857600080fd5b50919050565b6000806040838503121561191157600080fd5b82359150602083013567ffffffffffffffff81111561192f57600080fd5b61193b858286016118e6565b9150509250929050565b600061010082840312156118f857600080fd5b60008083601f84011261196a57600080fd5b50813567ffffffffffffffff81111561198257600080fd5b60208301915083602082850101111561199a57600080fd5b9250929050565b80356001600160a01b03811681146118b557600080fd5b6000604082840312156118f857600080fd5b60008060008060008060008060008060006102408c8e0312156119ec57600080fd5b6119f58c6118a6565b9a5060208c0135995060408c01359850611a128d60608e01611945565b97506101608c013567ffffffffffffffff811115611a2f57600080fd5b611a3b8e828f016118e6565b9750506101808c013567ffffffffffffffff811115611a5957600080fd5b611a658e828f01611958565b9097509550611a7990506101a08d016119a1565b93506101c08c01359250611a918d6101e08e016119b8565b915060006102208d01359050809150509295989b509295989b9093969950565b600081518084526020840193506020830160005b82811015611ae3578151865260209586019590910190600101611ac5565b5093949350505050565b60208152611aff602082018351611704565b60006020830151611b1b60408401826001600160a01b03169052565b506040830151805160608401526020810151608084015250606083015160a0830152608083015161028060c0840152611b586102a0840182611ab1565b905060a084015160e084015260c084015161010084015260e0840151610120840152610100840151610140840152610120840151601f1984830301610160850152611ba38282611718565b915050610140840151611bba61018085018261175e565b50610160840151838203601f1901610280850152611bd882826117be565b95945050505050565b600080600080600060608688031215611bf957600080fd5b85359450602086013567ffffffffffffffff811115611c1757600080fd5b611c2388828901611958565b909550935050604086013567ffffffffffffffff811115611c4357600080fd5b611c4f88828901611958565b969995985093965092949392505050565b602081526000611c736020830184611718565b9392505050565b60008060408385031215611c8d57600080fd5b50508035926020909101359150565b600181811c90821680611cb057607f821691505b6020821081036118f857634e487b7160e01b600052602260045260246000fd5b600060208284031215611ce257600080fd5b5051919050565b8015158114611cf757600080fd5b50565b600060208284031215611d0c57600080fd5b8151611c7381611ce9565b60208101611d258284611704565b92915050565b6000808335601e19843603018112611d4257600080fd5b83018035915067ffffffffffffffff821115611d5d57600080fd5b60200191503681900382131561199a57600080fd5b634e487b7160e01b600052604160045260246000fd5b601f821115611dd257806000526020600020601f840160051c81016020851015611daf5750805b601f840160051c820191505b81811015611dcf5760008155600101611dbb565b50505b505050565b67ffffffffffffffff831115611def57611def611d72565b611e0383611dfd8354611c9c565b83611d88565b6000601f841160018114611e375760008515611e1f5750838201355b600019600387901b1c1916600186901b178355611dcf565b600083815260209020601f19861690835b82811015611e685786850135825560209485019460019092019101611e48565b5086821015611e855760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b848152606060208201526000611eda606083018587611e97565b905082604083015295945050505050565b60ff81168114611cf757600080fd5b600060208284031215611f0c57600080fd5b8135611c7381611eeb565b634e487b7160e01b600052601160045260246000fd5b80820180821115611d2557611d25611f17565b600a8110611cf757600080fd5b600060208284031215611f5f57600080fd5b8135611c7381611f40565b60008135611d2581611ce9565b60008135611d2581611eeb565b8135611f8f81611ce9565b815490151560ff1660ff1991909116178155611fca611fb060208401611f6a565b82805461ff00191691151560081b61ff0016919091179055565b611ff1611fd960408401611f77565b825462ff0000191660109190911b62ff000016178255565b61201a61200060608401611f77565b825463ff000000191660189190911b63ff00000016178255565b6080820135600182015560a0820135600282015560c0820135600382015560e090910135600490910155565b813561205181611f40565b600a8110612061576120616116ee565b815460ff191660ff9190911617815560208201356001820155604082013560028201556060820135601e193684900301811261209c57600080fd5b8201803567ffffffffffffffff8111156120b557600080fd5b6020820191508036038213156120ca57600080fd5b6120d8818360038601611dd7565b50505050565b600063ffffffff821663ffffffff81036120fa576120fa611f17565b60010192915050565b604081526000612117604083018688611e97565b828103602084015261212a818587611e97565b979650505050505050565b604051601f8201601f1916810167ffffffffffffffff8111828210171561215e5761215e611d72565b604052919050565b6000610120828403121561217957600080fd5b600083601f840112612189578081fd5b8061012061219681612135565b92508291508401858111156121aa57600080fd5b845b818110156121c45780358452602093840193016121ac565b509095945050505050565b634e487b7160e01b600052603260045260246000fd5b81810381811115611d2557611d25611f17565b602081526000611c736020830184611ab1565b60006080828403121561221d57600080fd5b600083601f84011261222d578081fd5b806122386080612135565b905080915060808401858111156121aa57600080fdfea2646970667358221220cf11078e162e29af48a6fd5601d16be44a28d44b949029ab3fd53ba6a482c6d464736f6c634300081c0033",
}

// ProcessRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProcessRegistryMetaData.ABI instead.
var ProcessRegistryABI = ProcessRegistryMetaData.ABI

// ProcessRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProcessRegistryMetaData.Bin instead.
var ProcessRegistryBin = ProcessRegistryMetaData.Bin

// DeployProcessRegistry deploys a new Ethereum contract, binding an instance of ProcessRegistry to it.
func DeployProcessRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _chainID string, _organizationRegistryAddress common.Address, _stVerifier common.Address, _rVerifier common.Address) (common.Address, *types.Transaction, *ProcessRegistry, error) {
	parsed, err := ProcessRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProcessRegistryBin), backend, _chainID, _organizationRegistryAddress, _stVerifier, _rVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProcessRegistry{ProcessRegistryCaller: ProcessRegistryCaller{contract: contract}, ProcessRegistryTransactor: ProcessRegistryTransactor{contract: contract}, ProcessRegistryFilterer: ProcessRegistryFilterer{contract: contract}}, nil
}

// ProcessRegistry is an auto generated Go binding around an Ethereum contract.
type ProcessRegistry struct {
	ProcessRegistryCaller     // Read-only binding to the contract
	ProcessRegistryTransactor // Write-only binding to the contract
	ProcessRegistryFilterer   // Log filterer for contract events
}

// ProcessRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProcessRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProcessRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProcessRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProcessRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProcessRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProcessRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProcessRegistrySession struct {
	Contract     *ProcessRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProcessRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProcessRegistryCallerSession struct {
	Contract *ProcessRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProcessRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProcessRegistryTransactorSession struct {
	Contract     *ProcessRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProcessRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProcessRegistryRaw struct {
	Contract *ProcessRegistry // Generic contract binding to access the raw methods on
}

// ProcessRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProcessRegistryCallerRaw struct {
	Contract *ProcessRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProcessRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProcessRegistryTransactorRaw struct {
	Contract *ProcessRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProcessRegistry creates a new instance of ProcessRegistry, bound to a specific deployed contract.
func NewProcessRegistry(address common.Address, backend bind.ContractBackend) (*ProcessRegistry, error) {
	contract, err := bindProcessRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistry{ProcessRegistryCaller: ProcessRegistryCaller{contract: contract}, ProcessRegistryTransactor: ProcessRegistryTransactor{contract: contract}, ProcessRegistryFilterer: ProcessRegistryFilterer{contract: contract}}, nil
}

// NewProcessRegistryCaller creates a new read-only instance of ProcessRegistry, bound to a specific deployed contract.
func NewProcessRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProcessRegistryCaller, error) {
	contract, err := bindProcessRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryCaller{contract: contract}, nil
}

// NewProcessRegistryTransactor creates a new write-only instance of ProcessRegistry, bound to a specific deployed contract.
func NewProcessRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProcessRegistryTransactor, error) {
	contract, err := bindProcessRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryTransactor{contract: contract}, nil
}

// NewProcessRegistryFilterer creates a new log filterer instance of ProcessRegistry, bound to a specific deployed contract.
func NewProcessRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProcessRegistryFilterer, error) {
	contract, err := bindProcessRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryFilterer{contract: contract}, nil
}

// bindProcessRegistry binds a generic wrapper to an already deployed contract.
func bindProcessRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProcessRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProcessRegistry *ProcessRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProcessRegistry.Contract.ProcessRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProcessRegistry *ProcessRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.ProcessRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProcessRegistry *ProcessRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.ProcessRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProcessRegistry *ProcessRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProcessRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProcessRegistry *ProcessRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProcessRegistry *ProcessRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXCENSUSORIGIN is a free data retrieval call binding the contract method 0x02cc919a.
//
// Solidity: function MAX_CENSUS_ORIGIN() view returns(uint8)
func (_ProcessRegistry *ProcessRegistryCaller) MAXCENSUSORIGIN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "MAX_CENSUS_ORIGIN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXCENSUSORIGIN is a free data retrieval call binding the contract method 0x02cc919a.
//
// Solidity: function MAX_CENSUS_ORIGIN() view returns(uint8)
func (_ProcessRegistry *ProcessRegistrySession) MAXCENSUSORIGIN() (uint8, error) {
	return _ProcessRegistry.Contract.MAXCENSUSORIGIN(&_ProcessRegistry.CallOpts)
}

// MAXCENSUSORIGIN is a free data retrieval call binding the contract method 0x02cc919a.
//
// Solidity: function MAX_CENSUS_ORIGIN() view returns(uint8)
func (_ProcessRegistry *ProcessRegistryCallerSession) MAXCENSUSORIGIN() (uint8, error) {
	return _ProcessRegistry.Contract.MAXCENSUSORIGIN(&_ProcessRegistry.CallOpts)
}

// MAXSTATUS is a free data retrieval call binding the contract method 0x72c628ef.
//
// Solidity: function MAX_STATUS() view returns(uint8)
func (_ProcessRegistry *ProcessRegistryCaller) MAXSTATUS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "MAX_STATUS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXSTATUS is a free data retrieval call binding the contract method 0x72c628ef.
//
// Solidity: function MAX_STATUS() view returns(uint8)
func (_ProcessRegistry *ProcessRegistrySession) MAXSTATUS() (uint8, error) {
	return _ProcessRegistry.Contract.MAXSTATUS(&_ProcessRegistry.CallOpts)
}

// MAXSTATUS is a free data retrieval call binding the contract method 0x72c628ef.
//
// Solidity: function MAX_STATUS() view returns(uint8)
func (_ProcessRegistry *ProcessRegistryCallerSession) MAXSTATUS() (uint8, error) {
	return _ProcessRegistry.Contract.MAXSTATUS(&_ProcessRegistry.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(string)
func (_ProcessRegistry *ProcessRegistryCaller) ChainID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(string)
func (_ProcessRegistry *ProcessRegistrySession) ChainID() (string, error) {
	return _ProcessRegistry.Contract.ChainID(&_ProcessRegistry.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(string)
func (_ProcessRegistry *ProcessRegistryCallerSession) ChainID() (string, error) {
	return _ProcessRegistry.Contract.ChainID(&_ProcessRegistry.CallOpts)
}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,uint256,bytes32,string)))
func (_ProcessRegistry *ProcessRegistryCaller) GetProcess(opts *bind.CallOpts, processId [32]byte) (IProcessRegistryProcess, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getProcess", processId)

	if err != nil {
		return *new(IProcessRegistryProcess), err
	}

	out0 := *abi.ConvertType(out[0], new(IProcessRegistryProcess)).(*IProcessRegistryProcess)

	return out0, err

}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,uint256,bytes32,string)))
func (_ProcessRegistry *ProcessRegistrySession) GetProcess(processId [32]byte) (IProcessRegistryProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetProcess is a free data retrieval call binding the contract method 0x992bc45b.
//
// Solidity: function getProcess(bytes32 processId) view returns((uint8,address,(uint256,uint256),uint256,uint256[],uint256,uint256,uint256,uint256,string,(bool,bool,uint8,uint8,uint256,uint256,uint256,uint256),(uint8,uint256,bytes32,string)))
func (_ProcessRegistry *ProcessRegistryCallerSession) GetProcess(processId [32]byte) (IProcessRegistryProcess, error) {
	return _ProcessRegistry.Contract.GetProcess(&_ProcessRegistry.CallOpts, processId)
}

// GetRVerifierVKeyHash is a free data retrieval call binding the contract method 0xf9aa4499.
//
// Solidity: function getRVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) GetRVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getRVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRVerifierVKeyHash is a free data retrieval call binding the contract method 0xf9aa4499.
//
// Solidity: function getRVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) GetRVerifierVKeyHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.GetRVerifierVKeyHash(&_ProcessRegistry.CallOpts)
}

// GetRVerifierVKeyHash is a free data retrieval call binding the contract method 0xf9aa4499.
//
// Solidity: function getRVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetRVerifierVKeyHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.GetRVerifierVKeyHash(&_ProcessRegistry.CallOpts)
}

// GetSTVerifierVKeyHash is a free data retrieval call binding the contract method 0x4c0acc56.
//
// Solidity: function getSTVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCaller) GetSTVerifierVKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "getSTVerifierVKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSTVerifierVKeyHash is a free data retrieval call binding the contract method 0x4c0acc56.
//
// Solidity: function getSTVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistrySession) GetSTVerifierVKeyHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.GetSTVerifierVKeyHash(&_ProcessRegistry.CallOpts)
}

// GetSTVerifierVKeyHash is a free data retrieval call binding the contract method 0x4c0acc56.
//
// Solidity: function getSTVerifierVKeyHash() view returns(bytes32)
func (_ProcessRegistry *ProcessRegistryCallerSession) GetSTVerifierVKeyHash() ([32]byte, error) {
	return _ProcessRegistry.Contract.GetSTVerifierVKeyHash(&_ProcessRegistry.CallOpts)
}

// OrganizationRegistryAddress is a free data retrieval call binding the contract method 0xb79938e9.
//
// Solidity: function organizationRegistryAddress() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) OrganizationRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "organizationRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrganizationRegistryAddress is a free data retrieval call binding the contract method 0xb79938e9.
//
// Solidity: function organizationRegistryAddress() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) OrganizationRegistryAddress() (common.Address, error) {
	return _ProcessRegistry.Contract.OrganizationRegistryAddress(&_ProcessRegistry.CallOpts)
}

// OrganizationRegistryAddress is a free data retrieval call binding the contract method 0xb79938e9.
//
// Solidity: function organizationRegistryAddress() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) OrganizationRegistryAddress() (common.Address, error) {
	return _ProcessRegistry.Contract.OrganizationRegistryAddress(&_ProcessRegistry.CallOpts)
}

// ProcessCount is a free data retrieval call binding the contract method 0x848df540.
//
// Solidity: function processCount() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCaller) ProcessCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "processCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ProcessCount is a free data retrieval call binding the contract method 0x848df540.
//
// Solidity: function processCount() view returns(uint32)
func (_ProcessRegistry *ProcessRegistrySession) ProcessCount() (uint32, error) {
	return _ProcessRegistry.Contract.ProcessCount(&_ProcessRegistry.CallOpts)
}

// ProcessCount is a free data retrieval call binding the contract method 0x848df540.
//
// Solidity: function processCount() view returns(uint32)
func (_ProcessRegistry *ProcessRegistryCallerSession) ProcessCount() (uint32, error) {
	return _ProcessRegistry.Contract.ProcessCount(&_ProcessRegistry.CallOpts)
}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 voteCount, uint256 voteOverwriteCount, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census)
func (_ProcessRegistry *ProcessRegistryCaller) Processes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status             uint8
	OrganizationId     common.Address
	EncryptionKey      IProcessRegistryEncryptionKey
	LatestStateRoot    *big.Int
	StartTime          *big.Int
	Duration           *big.Int
	VoteCount          *big.Int
	VoteOverwriteCount *big.Int
	MetadataURI        string
	BallotMode         IProcessRegistryBallotMode
	Census             IProcessRegistryCensus
}, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "processes", arg0)

	outstruct := new(struct {
		Status             uint8
		OrganizationId     common.Address
		EncryptionKey      IProcessRegistryEncryptionKey
		LatestStateRoot    *big.Int
		StartTime          *big.Int
		Duration           *big.Int
		VoteCount          *big.Int
		VoteOverwriteCount *big.Int
		MetadataURI        string
		BallotMode         IProcessRegistryBallotMode
		Census             IProcessRegistryCensus
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.OrganizationId = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.EncryptionKey = *abi.ConvertType(out[2], new(IProcessRegistryEncryptionKey)).(*IProcessRegistryEncryptionKey)
	outstruct.LatestStateRoot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.VoteCount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.VoteOverwriteCount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.MetadataURI = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.BallotMode = *abi.ConvertType(out[9], new(IProcessRegistryBallotMode)).(*IProcessRegistryBallotMode)
	outstruct.Census = *abi.ConvertType(out[10], new(IProcessRegistryCensus)).(*IProcessRegistryCensus)

	return *outstruct, err

}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 voteCount, uint256 voteOverwriteCount, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census)
func (_ProcessRegistry *ProcessRegistrySession) Processes(arg0 [32]byte) (struct {
	Status             uint8
	OrganizationId     common.Address
	EncryptionKey      IProcessRegistryEncryptionKey
	LatestStateRoot    *big.Int
	StartTime          *big.Int
	Duration           *big.Int
	VoteCount          *big.Int
	VoteOverwriteCount *big.Int
	MetadataURI        string
	BallotMode         IProcessRegistryBallotMode
	Census             IProcessRegistryCensus
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// Processes is a free data retrieval call binding the contract method 0x0535fece.
//
// Solidity: function processes(bytes32 ) view returns(uint8 status, address organizationId, (uint256,uint256) encryptionKey, uint256 latestStateRoot, uint256 startTime, uint256 duration, uint256 voteCount, uint256 voteOverwriteCount, string metadataURI, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census)
func (_ProcessRegistry *ProcessRegistryCallerSession) Processes(arg0 [32]byte) (struct {
	Status             uint8
	OrganizationId     common.Address
	EncryptionKey      IProcessRegistryEncryptionKey
	LatestStateRoot    *big.Int
	StartTime          *big.Int
	Duration           *big.Int
	VoteCount          *big.Int
	VoteOverwriteCount *big.Int
	MetadataURI        string
	BallotMode         IProcessRegistryBallotMode
	Census             IProcessRegistryCensus
}, error) {
	return _ProcessRegistry.Contract.Processes(&_ProcessRegistry.CallOpts, arg0)
}

// RVerifier is a free data retrieval call binding the contract method 0x9c18deaf.
//
// Solidity: function rVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) RVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "rVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RVerifier is a free data retrieval call binding the contract method 0x9c18deaf.
//
// Solidity: function rVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) RVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.RVerifier(&_ProcessRegistry.CallOpts)
}

// RVerifier is a free data retrieval call binding the contract method 0x9c18deaf.
//
// Solidity: function rVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) RVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.RVerifier(&_ProcessRegistry.CallOpts)
}

// StVerifier is a free data retrieval call binding the contract method 0x09706392.
//
// Solidity: function stVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCaller) StVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProcessRegistry.contract.Call(opts, &out, "stVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StVerifier is a free data retrieval call binding the contract method 0x09706392.
//
// Solidity: function stVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistrySession) StVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.StVerifier(&_ProcessRegistry.CallOpts)
}

// StVerifier is a free data retrieval call binding the contract method 0x09706392.
//
// Solidity: function stVerifier() view returns(address)
func (_ProcessRegistry *ProcessRegistryCallerSession) StVerifier() (common.Address, error) {
	return _ProcessRegistry.Contract.StVerifier(&_ProcessRegistry.CallOpts)
}

// NewProcess is a paid mutator transaction binding the contract method 0x92bc2669.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census, string metadata, address organizationId, bytes32 processId, (uint256,uint256) encryptionKey, uint256 initStateRoot) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) NewProcess(opts *bind.TransactOpts, status uint8, startTime *big.Int, duration *big.Int, ballotMode IProcessRegistryBallotMode, census IProcessRegistryCensus, metadata string, organizationId common.Address, processId [32]byte, encryptionKey IProcessRegistryEncryptionKey, initStateRoot *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "newProcess", status, startTime, duration, ballotMode, census, metadata, organizationId, processId, encryptionKey, initStateRoot)
}

// NewProcess is a paid mutator transaction binding the contract method 0x92bc2669.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census, string metadata, address organizationId, bytes32 processId, (uint256,uint256) encryptionKey, uint256 initStateRoot) returns()
func (_ProcessRegistry *ProcessRegistrySession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, ballotMode IProcessRegistryBallotMode, census IProcessRegistryCensus, metadata string, organizationId common.Address, processId [32]byte, encryptionKey IProcessRegistryEncryptionKey, initStateRoot *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, ballotMode, census, metadata, organizationId, processId, encryptionKey, initStateRoot)
}

// NewProcess is a paid mutator transaction binding the contract method 0x92bc2669.
//
// Solidity: function newProcess(uint8 status, uint256 startTime, uint256 duration, (bool,bool,uint8,uint8,uint256,uint256,uint256,uint256) ballotMode, (uint8,uint256,bytes32,string) census, string metadata, address organizationId, bytes32 processId, (uint256,uint256) encryptionKey, uint256 initStateRoot) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) NewProcess(status uint8, startTime *big.Int, duration *big.Int, ballotMode IProcessRegistryBallotMode, census IProcessRegistryCensus, metadata string, organizationId common.Address, processId [32]byte, encryptionKey IProcessRegistryEncryptionKey, initStateRoot *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.NewProcess(&_ProcessRegistry.TransactOpts, status, startTime, duration, ballotMode, census, metadata, organizationId, processId, encryptionKey, initStateRoot)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x6bae04ea.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,uint256,bytes32,string) census) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessCensus(opts *bind.TransactOpts, processId [32]byte, census IProcessRegistryCensus) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessCensus", processId, census)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x6bae04ea.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,uint256,bytes32,string) census) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessCensus(processId [32]byte, census IProcessRegistryCensus) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessCensus(&_ProcessRegistry.TransactOpts, processId, census)
}

// SetProcessCensus is a paid mutator transaction binding the contract method 0x6bae04ea.
//
// Solidity: function setProcessCensus(bytes32 processId, (uint8,uint256,bytes32,string) census) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessCensus(processId [32]byte, census IProcessRegistryCensus) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessCensus(&_ProcessRegistry.TransactOpts, processId, census)
}

// SetProcessDuration is a paid mutator transaction binding the contract method 0xc718c01f.
//
// Solidity: function setProcessDuration(bytes32 processId, uint256 _duration) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessDuration(opts *bind.TransactOpts, processId [32]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessDuration", processId, _duration)
}

// SetProcessDuration is a paid mutator transaction binding the contract method 0xc718c01f.
//
// Solidity: function setProcessDuration(bytes32 processId, uint256 _duration) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessDuration(processId [32]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessDuration(&_ProcessRegistry.TransactOpts, processId, _duration)
}

// SetProcessDuration is a paid mutator transaction binding the contract method 0xc718c01f.
//
// Solidity: function setProcessDuration(bytes32 processId, uint256 _duration) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessDuration(processId [32]byte, _duration *big.Int) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessDuration(&_ProcessRegistry.TransactOpts, processId, _duration)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0xa5b110de.
//
// Solidity: function setProcessResults(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessResults(opts *bind.TransactOpts, processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessResults", processId, proof, input)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0xa5b110de.
//
// Solidity: function setProcessResults(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessResults(processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessResults(&_ProcessRegistry.TransactOpts, processId, proof, input)
}

// SetProcessResults is a paid mutator transaction binding the contract method 0xa5b110de.
//
// Solidity: function setProcessResults(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessResults(processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessResults(&_ProcessRegistry.TransactOpts, processId, proof, input)
}

// SetProcessStatus is a paid mutator transaction binding the contract method 0x56a6f1e2.
//
// Solidity: function setProcessStatus(bytes32 processId, uint8 newStatus) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SetProcessStatus(opts *bind.TransactOpts, processId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "setProcessStatus", processId, newStatus)
}

// SetProcessStatus is a paid mutator transaction binding the contract method 0x56a6f1e2.
//
// Solidity: function setProcessStatus(bytes32 processId, uint8 newStatus) returns()
func (_ProcessRegistry *ProcessRegistrySession) SetProcessStatus(processId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessStatus(&_ProcessRegistry.TransactOpts, processId, newStatus)
}

// SetProcessStatus is a paid mutator transaction binding the contract method 0x56a6f1e2.
//
// Solidity: function setProcessStatus(bytes32 processId, uint8 newStatus) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SetProcessStatus(processId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SetProcessStatus(&_ProcessRegistry.TransactOpts, processId, newStatus)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xd71d072e.
//
// Solidity: function submitStateTransition(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactor) SubmitStateTransition(opts *bind.TransactOpts, processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.contract.Transact(opts, "submitStateTransition", processId, proof, input)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xd71d072e.
//
// Solidity: function submitStateTransition(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistrySession) SubmitStateTransition(processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SubmitStateTransition(&_ProcessRegistry.TransactOpts, processId, proof, input)
}

// SubmitStateTransition is a paid mutator transaction binding the contract method 0xd71d072e.
//
// Solidity: function submitStateTransition(bytes32 processId, bytes proof, bytes input) returns()
func (_ProcessRegistry *ProcessRegistryTransactorSession) SubmitStateTransition(processId [32]byte, proof []byte, input []byte) (*types.Transaction, error) {
	return _ProcessRegistry.Contract.SubmitStateTransition(&_ProcessRegistry.TransactOpts, processId, proof, input)
}

// ProcessRegistryCensusUpdatedIterator is returned from FilterCensusUpdated and is used to iterate over the raw logs and unpacked data for CensusUpdated events raised by the ProcessRegistry contract.
type ProcessRegistryCensusUpdatedIterator struct {
	Event *ProcessRegistryCensusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProcessRegistryCensusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryCensusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProcessRegistryCensusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProcessRegistryCensusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryCensusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryCensusUpdated represents a CensusUpdated event raised by the ProcessRegistry contract.
type ProcessRegistryCensusUpdated struct {
	ProcessId  [32]byte
	CensusRoot [32]byte
	CensusURI  string
	MaxVotes   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCensusUpdated is a free log retrieval operation binding the contract event 0x35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b.
//
// Solidity: event CensusUpdated(bytes32 indexed processId, bytes32 censusRoot, string censusURI, uint256 maxVotes)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterCensusUpdated(opts *bind.FilterOpts, processId [][32]byte) (*ProcessRegistryCensusUpdatedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "CensusUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryCensusUpdatedIterator{contract: _ProcessRegistry.contract, event: "CensusUpdated", logs: logs, sub: sub}, nil
}

// WatchCensusUpdated is a free log subscription operation binding the contract event 0x35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b.
//
// Solidity: event CensusUpdated(bytes32 indexed processId, bytes32 censusRoot, string censusURI, uint256 maxVotes)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchCensusUpdated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryCensusUpdated, processId [][32]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "CensusUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryCensusUpdated)
				if err := _ProcessRegistry.contract.UnpackLog(event, "CensusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCensusUpdated is a log parse operation binding the contract event 0x35947a8913e2156f19b018078c9f0667e49cb3dc24af3434a4d0b16b82675b1b.
//
// Solidity: event CensusUpdated(bytes32 indexed processId, bytes32 censusRoot, string censusURI, uint256 maxVotes)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseCensusUpdated(log types.Log) (*ProcessRegistryCensusUpdated, error) {
	event := new(ProcessRegistryCensusUpdated)
	if err := _ProcessRegistry.contract.UnpackLog(event, "CensusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessCreatedIterator is returned from FilterProcessCreated and is used to iterate over the raw logs and unpacked data for ProcessCreated events raised by the ProcessRegistry contract.
type ProcessRegistryProcessCreatedIterator struct {
	Event *ProcessRegistryProcessCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProcessRegistryProcessCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProcessRegistryProcessCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProcessRegistryProcessCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessCreated represents a ProcessCreated event raised by the ProcessRegistry contract.
type ProcessRegistryProcessCreated struct {
	ProcessId [32]byte
	Creator   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessCreated is a free log retrieval operation binding the contract event 0xada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa.
//
// Solidity: event ProcessCreated(bytes32 indexed processId, address indexed creator)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessCreated(opts *bind.FilterOpts, processId [][32]byte, creator []common.Address) (*ProcessRegistryProcessCreatedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessCreated", processIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessCreatedIterator{contract: _ProcessRegistry.contract, event: "ProcessCreated", logs: logs, sub: sub}, nil
}

// WatchProcessCreated is a free log subscription operation binding the contract event 0xada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa.
//
// Solidity: event ProcessCreated(bytes32 indexed processId, address indexed creator)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessCreated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessCreated, processId [][32]byte, creator []common.Address) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessCreated", processIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessCreated)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProcessCreated is a log parse operation binding the contract event 0xada6f87a2a16a0c9c169ca36754c5f33f7c1a973b575d068f888a549ed4faefa.
//
// Solidity: event ProcessCreated(bytes32 indexed processId, address indexed creator)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessCreated(log types.Log) (*ProcessRegistryProcessCreated, error) {
	event := new(ProcessRegistryProcessCreated)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessDurationChangedIterator is returned from FilterProcessDurationChanged and is used to iterate over the raw logs and unpacked data for ProcessDurationChanged events raised by the ProcessRegistry contract.
type ProcessRegistryProcessDurationChangedIterator struct {
	Event *ProcessRegistryProcessDurationChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProcessRegistryProcessDurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessDurationChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProcessRegistryProcessDurationChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProcessRegistryProcessDurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessDurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessDurationChanged represents a ProcessDurationChanged event raised by the ProcessRegistry contract.
type ProcessRegistryProcessDurationChanged struct {
	ProcessId [32]byte
	Duration  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessDurationChanged is a free log retrieval operation binding the contract event 0x0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f46.
//
// Solidity: event ProcessDurationChanged(bytes32 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessDurationChanged(opts *bind.FilterOpts, processId [][32]byte) (*ProcessRegistryProcessDurationChangedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessDurationChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessDurationChangedIterator{contract: _ProcessRegistry.contract, event: "ProcessDurationChanged", logs: logs, sub: sub}, nil
}

// WatchProcessDurationChanged is a free log subscription operation binding the contract event 0x0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f46.
//
// Solidity: event ProcessDurationChanged(bytes32 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessDurationChanged(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessDurationChanged, processId [][32]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessDurationChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessDurationChanged)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessDurationChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProcessDurationChanged is a log parse operation binding the contract event 0x0f759826327c668a220d576485ac38ddc4f83fbc414b984c00e79f669b649f46.
//
// Solidity: event ProcessDurationChanged(bytes32 indexed processId, uint256 duration)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessDurationChanged(log types.Log) (*ProcessRegistryProcessDurationChanged, error) {
	event := new(ProcessRegistryProcessDurationChanged)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessDurationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessResultsSetIterator is returned from FilterProcessResultsSet and is used to iterate over the raw logs and unpacked data for ProcessResultsSet events raised by the ProcessRegistry contract.
type ProcessRegistryProcessResultsSetIterator struct {
	Event *ProcessRegistryProcessResultsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProcessRegistryProcessResultsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessResultsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProcessRegistryProcessResultsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProcessRegistryProcessResultsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessResultsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessResultsSet represents a ProcessResultsSet event raised by the ProcessRegistry contract.
type ProcessRegistryProcessResultsSet struct {
	ProcessId [32]byte
	Result    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessResultsSet is a free log retrieval operation binding the contract event 0xa21d6aca26cc870bcd3da1cc0d0b0f22572550a7d04fac698fc08b67b92c4352.
//
// Solidity: event ProcessResultsSet(bytes32 indexed processId, uint256[] result)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessResultsSet(opts *bind.FilterOpts, processId [][32]byte) (*ProcessRegistryProcessResultsSetIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessResultsSet", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessResultsSetIterator{contract: _ProcessRegistry.contract, event: "ProcessResultsSet", logs: logs, sub: sub}, nil
}

// WatchProcessResultsSet is a free log subscription operation binding the contract event 0xa21d6aca26cc870bcd3da1cc0d0b0f22572550a7d04fac698fc08b67b92c4352.
//
// Solidity: event ProcessResultsSet(bytes32 indexed processId, uint256[] result)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessResultsSet(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessResultsSet, processId [][32]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessResultsSet", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessResultsSet)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessResultsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProcessResultsSet is a log parse operation binding the contract event 0xa21d6aca26cc870bcd3da1cc0d0b0f22572550a7d04fac698fc08b67b92c4352.
//
// Solidity: event ProcessResultsSet(bytes32 indexed processId, uint256[] result)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessResultsSet(log types.Log) (*ProcessRegistryProcessResultsSet, error) {
	event := new(ProcessRegistryProcessResultsSet)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessResultsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessStateRootUpdatedIterator is returned from FilterProcessStateRootUpdated and is used to iterate over the raw logs and unpacked data for ProcessStateRootUpdated events raised by the ProcessRegistry contract.
type ProcessRegistryProcessStateRootUpdatedIterator struct {
	Event *ProcessRegistryProcessStateRootUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProcessRegistryProcessStateRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessStateRootUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProcessRegistryProcessStateRootUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProcessRegistryProcessStateRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessStateRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessStateRootUpdated represents a ProcessStateRootUpdated event raised by the ProcessRegistry contract.
type ProcessRegistryProcessStateRootUpdated struct {
	ProcessId    [32]byte
	NewStateRoot *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProcessStateRootUpdated is a free log retrieval operation binding the contract event 0xd25cbd3fa66107efb1230dc4f49f939d7b6b57ba10ab0ec87ba39040c075cf20.
//
// Solidity: event ProcessStateRootUpdated(bytes32 indexed processId, uint256 newStateRoot)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessStateRootUpdated(opts *bind.FilterOpts, processId [][32]byte) (*ProcessRegistryProcessStateRootUpdatedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessStateRootUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessStateRootUpdatedIterator{contract: _ProcessRegistry.contract, event: "ProcessStateRootUpdated", logs: logs, sub: sub}, nil
}

// WatchProcessStateRootUpdated is a free log subscription operation binding the contract event 0xd25cbd3fa66107efb1230dc4f49f939d7b6b57ba10ab0ec87ba39040c075cf20.
//
// Solidity: event ProcessStateRootUpdated(bytes32 indexed processId, uint256 newStateRoot)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessStateRootUpdated(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessStateRootUpdated, processId [][32]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessStateRootUpdated", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessStateRootUpdated)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStateRootUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProcessStateRootUpdated is a log parse operation binding the contract event 0xd25cbd3fa66107efb1230dc4f49f939d7b6b57ba10ab0ec87ba39040c075cf20.
//
// Solidity: event ProcessStateRootUpdated(bytes32 indexed processId, uint256 newStateRoot)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessStateRootUpdated(log types.Log) (*ProcessRegistryProcessStateRootUpdated, error) {
	event := new(ProcessRegistryProcessStateRootUpdated)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStateRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProcessRegistryProcessStatusChangedIterator is returned from FilterProcessStatusChanged and is used to iterate over the raw logs and unpacked data for ProcessStatusChanged events raised by the ProcessRegistry contract.
type ProcessRegistryProcessStatusChangedIterator struct {
	Event *ProcessRegistryProcessStatusChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProcessRegistryProcessStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProcessRegistryProcessStatusChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProcessRegistryProcessStatusChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProcessRegistryProcessStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProcessRegistryProcessStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProcessRegistryProcessStatusChanged represents a ProcessStatusChanged event raised by the ProcessRegistry contract.
type ProcessRegistryProcessStatusChanged struct {
	ProcessId [32]byte
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessStatusChanged is a free log retrieval operation binding the contract event 0xac0c4085a30bc70d0d8893ee5d6466ac9c5f03e27fd7292dcef128a610e7c190.
//
// Solidity: event ProcessStatusChanged(bytes32 indexed processId, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) FilterProcessStatusChanged(opts *bind.FilterOpts, processId [][32]byte) (*ProcessRegistryProcessStatusChangedIterator, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.FilterLogs(opts, "ProcessStatusChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return &ProcessRegistryProcessStatusChangedIterator{contract: _ProcessRegistry.contract, event: "ProcessStatusChanged", logs: logs, sub: sub}, nil
}

// WatchProcessStatusChanged is a free log subscription operation binding the contract event 0xac0c4085a30bc70d0d8893ee5d6466ac9c5f03e27fd7292dcef128a610e7c190.
//
// Solidity: event ProcessStatusChanged(bytes32 indexed processId, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) WatchProcessStatusChanged(opts *bind.WatchOpts, sink chan<- *ProcessRegistryProcessStatusChanged, processId [][32]byte) (event.Subscription, error) {

	var processIdRule []interface{}
	for _, processIdItem := range processId {
		processIdRule = append(processIdRule, processIdItem)
	}

	logs, sub, err := _ProcessRegistry.contract.WatchLogs(opts, "ProcessStatusChanged", processIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProcessRegistryProcessStatusChanged)
				if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStatusChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProcessStatusChanged is a log parse operation binding the contract event 0xac0c4085a30bc70d0d8893ee5d6466ac9c5f03e27fd7292dcef128a610e7c190.
//
// Solidity: event ProcessStatusChanged(bytes32 indexed processId, uint8 newStatus)
func (_ProcessRegistry *ProcessRegistryFilterer) ParseProcessStatusChanged(log types.Log) (*ProcessRegistryProcessStatusChanged, error) {
	event := new(ProcessRegistryProcessStatusChanged)
	if err := _ProcessRegistry.contract.UnpackLog(event, "ProcessStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
