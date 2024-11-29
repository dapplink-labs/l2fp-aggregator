// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// IMantaServiceManagerBatchHeader is an auto generated low-level Go binding around an user-defined struct.
type IMantaServiceManagerBatchHeader struct {
	FinalityRoot          [32]byte
	QuorumNumbers         []byte
	SignedStakeForQuorums []byte
	ReferenceBlockNumber  uint32
	OutputRoot            [32]byte
	L2BlockNumber         *big.Int
	L1BlockHash           [32]byte
	L1BlockNumber         *big.Int
}

// MantaServiceManagerMetaData contains all meta data concerning the MantaServiceManager contract.
var MantaServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_finalityAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BLOCK_STALE_MEASURE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"THRESHOLD_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchIdToBatchMetadataHash\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalityAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isBatchConfirmer\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2OutputOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumAdversaryThresholdPercentages\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumConfirmationThresholdPercentages\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumNumbersRequired\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyFinality\",\"inputs\":[{\"name\":\"batchHeader\",\"type\":\"tuple\",\"internalType\":\"structIMantaServiceManager.BatchHeader\",\"components\":[{\"name\":\"finalityRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signedStakeForQuorums\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"outputRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l1BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l1BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FinalityVerified\",\"inputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"outputRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"l1BlockHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"l1BlockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516110c73803806110c783398101604081905261002f9161011f565b6001600160a01b03808316608052811660a05261004a610051565b5050610152565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a15760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101005780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b80516001600160a01b038116811461011a57600080fd5b919050565b6000806040838503121561013257600080fd5b61013b83610103565b915061014960208401610103565b90509250929050565b60805160a051610f42610185600039600081816101e201526109130152600081816102ed01526105d60152610f426000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063a217fddf116100b8578063d547741f1161007c578063d547741f14610322578063e15234ff14610335578063e3feaa1714610352578063eccbbfc914610365578063ef02445814610385578063f2fde38b1461038d57600080fd5b8063a217fddf1461029d578063a5b7890a146102a5578063bafa9107146102c8578063c228fd08146102e8578063c4d66de81461030f57600080fd5b80634d9f15591161010a5780634d9f1559146101dd5780635e8b3f2d1461021c578063715018a6146102255780638687feae1461022d5780638da5cb5b1461025a57806391d148541461028a57600080fd5b806301ffc9a714610147578063248a9ca31461016f5780632f2ff15d1461019057806336568abe146101a55780634972134a146101b8575b600080fd5b61015a610155366004610cd1565b6103a0565b60405190151581526020015b60405180910390f35b61018261017d366004610d02565b6103d7565b604051908152602001610166565b6101a361019e366004610d37565b6103f9565b005b6101a36101b3366004610d37565b61041b565b6000546101c89063ffffffff1681565b60405163ffffffff9091168152602001610166565b6102047f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610166565b6101c861012c81565b6101a3610453565b61024d604051806040016040528060018152602001602160f81b81525081565b6040516101669190610d87565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610204565b61015a610298366004610d37565b610467565b610182600081565b61015a6102b3366004610dba565b60026020526000908152604090205460ff1681565b61024d604051806040016040528060018152602001603760f81b81525081565b6102047f000000000000000000000000000000000000000000000000000000000000000081565b6101a361031d366004610dba565b61049f565b6101a3610330366004610d37565b6105af565b61024d604051806040016040528060018152602001600081525081565b6101a3610360366004610dd5565b6105cb565b610182610373366004610e11565b60016020526000908152604090205481565b610182606481565b6101a361039b366004610dba565b610a58565b60006001600160e01b03198216637965db0b60e01b14806103d157506301ffc9a760e01b6001600160e01b03198316145b92915050565b6000908152600080516020610eed833981519152602052604090206001015490565b610402826103d7565b61040b81610a96565b6104158383610aa0565b50505050565b6001600160a01b03811633146104445760405163334bd91960e11b815260040160405180910390fd5b61044e8282610b4c565b505050565b61045b610bc8565b6104656000610c23565b565b6000918252600080516020610eed833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156104e55750825b905060008267ffffffffffffffff1660011480156105025750303b155b905081158015610510575080155b1561052e5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561055857845460ff60401b1916600160401b1785555b61056186610c23565b83156105a757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6105b8826103d7565b6105c181610a96565b6104158383610b4c565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610670576040805162461bcd60e51b81526020600482015260248101919091527f4d616e7461536572766963654d616e616765722e6f6e6c792066696e616c697460448201527f79206d616e616765722063616e2063616c6c20746869732066756e6374696f6e60648201526084015b60405180910390fd5b3233146106ed5760405162461bcd60e51b81526020600482015260516024820152600080516020610ecd83398151915260448201527f74793a2068656164657220616e64206e6f6e7369676e65722064617461206d75606482015270737420626520696e2063616c6c6461746160781b608482015260a401610667565b436106fe6080830160608401610e11565b63ffffffff161061077d5760405162461bcd60e51b815260206004820152604f6024820152600080516020610ecd83398151915260448201527f74793a20737065636966696564207265666572656e6365426c6f636b4e756d6260648201526e657220697320696e2066757475726560881b608482015260a401610667565b63ffffffff431661012c6107976080840160608501610e11565b6107a19190610e37565b63ffffffff1610156108275760405162461bcd60e51b81526020600482015260556024820152600080516020610ecd83398151915260448201527f74793a20737065636966696564207265666572656e6365426c6f636b4e756d62606482015274195c881a5cc81d1bdbc819985c881a5b881c185cdd605a1b608482015260a401610667565b6108346040820182610e62565b90506108436020830183610e62565b9050146108db5760405162461bcd60e51b81526020600482015260666024820152600080516020610ecd83398151915260448201527f74793a2071756f72756d4e756d6265727320616e64207369676e65645374616b60648201527f65466f7251756f72756d73206d757374206265206f66207468652073616d65206084820152650d8cadccee8d60d31b60a482015260c401610667565b6040516080820135602482015260a0820135604482015260c0820135606482015260e082013560848201526000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169060a40160408051601f198184030181529181526020820180516001600160e01b031663135556c960e31b1790525161096c9190610eb0565b6000604051808303816000865af19150503d80600081146109a9576040519150601f19603f3d011682016040523d82523d6000602084013e6109ae565b606091505b50509050806109ff5760405162461bcd60e51b815260206004820152601a60248201527f63616c6c2070726f706f7365204c324f7574707574206661696c0000000000006044820152606401610667565b60408051608084810135825260a0850135602083015260c08501358284015260e08501356060830152915133927fbea36b7214d7d22729389923a440c2ae7ecb5a9c0039ef0f3659f76f478a52b2928290030190a25050565b610a60610bc8565b6001600160a01b038116610a8a57604051631e4fbdf760e01b815260006004820152602401610667565b610a9381610c23565b50565b610a938133610c94565b6000600080516020610eed833981519152610abb8484610467565b610b3b576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055610af13390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506103d1565b60009150506103d1565b5092915050565b6000600080516020610eed833981519152610b678484610467565b15610b3b576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506103d1565b33610bfa7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146104655760405163118cdaa760e01b8152336004820152602401610667565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b610c9e8282610467565b610ccd5760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610667565b5050565b600060208284031215610ce357600080fd5b81356001600160e01b031981168114610cfb57600080fd5b9392505050565b600060208284031215610d1457600080fd5b5035919050565b80356001600160a01b0381168114610d3257600080fd5b919050565b60008060408385031215610d4a57600080fd5b82359150610d5a60208401610d1b565b90509250929050565b60005b83811015610d7e578181015183820152602001610d66565b50506000910152565b6020815260008251806020840152610da6816040850160208701610d63565b601f01601f19169190910160400192915050565b600060208284031215610dcc57600080fd5b610cfb82610d1b565b600060208284031215610de757600080fd5b813567ffffffffffffffff811115610dfe57600080fd5b82016101008185031215610cfb57600080fd5b600060208284031215610e2357600080fd5b813563ffffffff81168114610cfb57600080fd5b63ffffffff818116838216019080821115610b4557634e487b7160e01b600052601160045260246000fd5b6000808335601e19843603018112610e7957600080fd5b83018035915067ffffffffffffffff821115610e9457600080fd5b602001915036819003821315610ea957600080fd5b9250929050565b60008251610ec2818460208701610d63565b919091019291505056fe4d616e7461536572766963654d616e616765722e76657269667946696e616c6902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a26469706673582212201053aac7a6ce917b8db55d0d6e12429697ecd04d7ad9cbc922f0a82cd721356564736f6c63430008180033",
}

// MantaServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MantaServiceManagerMetaData.ABI instead.
var MantaServiceManagerABI = MantaServiceManagerMetaData.ABI

// MantaServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MantaServiceManagerMetaData.Bin instead.
var MantaServiceManagerBin = MantaServiceManagerMetaData.Bin

// DeployMantaServiceManager deploys a new Ethereum contract, binding an instance of MantaServiceManager to it.
func DeployMantaServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _finalityAddress common.Address, _l2OutputOracle common.Address) (common.Address, *types.Transaction, *MantaServiceManager, error) {
	parsed, err := MantaServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MantaServiceManagerBin), backend, _finalityAddress, _l2OutputOracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MantaServiceManager{MantaServiceManagerCaller: MantaServiceManagerCaller{contract: contract}, MantaServiceManagerTransactor: MantaServiceManagerTransactor{contract: contract}, MantaServiceManagerFilterer: MantaServiceManagerFilterer{contract: contract}}, nil
}

// MantaServiceManager is an auto generated Go binding around an Ethereum contract.
type MantaServiceManager struct {
	MantaServiceManagerCaller     // Read-only binding to the contract
	MantaServiceManagerTransactor // Write-only binding to the contract
	MantaServiceManagerFilterer   // Log filterer for contract events
}

// MantaServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MantaServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantaServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MantaServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantaServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MantaServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantaServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MantaServiceManagerSession struct {
	Contract     *MantaServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MantaServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MantaServiceManagerCallerSession struct {
	Contract *MantaServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MantaServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MantaServiceManagerTransactorSession struct {
	Contract     *MantaServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MantaServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MantaServiceManagerRaw struct {
	Contract *MantaServiceManager // Generic contract binding to access the raw methods on
}

// MantaServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MantaServiceManagerCallerRaw struct {
	Contract *MantaServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MantaServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MantaServiceManagerTransactorRaw struct {
	Contract *MantaServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMantaServiceManager creates a new instance of MantaServiceManager, bound to a specific deployed contract.
func NewMantaServiceManager(address common.Address, backend bind.ContractBackend) (*MantaServiceManager, error) {
	contract, err := bindMantaServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MantaServiceManager{MantaServiceManagerCaller: MantaServiceManagerCaller{contract: contract}, MantaServiceManagerTransactor: MantaServiceManagerTransactor{contract: contract}, MantaServiceManagerFilterer: MantaServiceManagerFilterer{contract: contract}}, nil
}

// NewMantaServiceManagerCaller creates a new read-only instance of MantaServiceManager, bound to a specific deployed contract.
func NewMantaServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*MantaServiceManagerCaller, error) {
	contract, err := bindMantaServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MantaServiceManagerCaller{contract: contract}, nil
}

// NewMantaServiceManagerTransactor creates a new write-only instance of MantaServiceManager, bound to a specific deployed contract.
func NewMantaServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MantaServiceManagerTransactor, error) {
	contract, err := bindMantaServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MantaServiceManagerTransactor{contract: contract}, nil
}

// NewMantaServiceManagerFilterer creates a new log filterer instance of MantaServiceManager, bound to a specific deployed contract.
func NewMantaServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MantaServiceManagerFilterer, error) {
	contract, err := bindMantaServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MantaServiceManagerFilterer{contract: contract}, nil
}

// bindMantaServiceManager binds a generic wrapper to an already deployed contract.
func bindMantaServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MantaServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MantaServiceManager *MantaServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MantaServiceManager.Contract.MantaServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MantaServiceManager *MantaServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.MantaServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MantaServiceManager *MantaServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.MantaServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MantaServiceManager *MantaServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MantaServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MantaServiceManager *MantaServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MantaServiceManager *MantaServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.contract.Transact(opts, method, params...)
}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint32)
func (_MantaServiceManager *MantaServiceManagerCaller) BLOCKSTALEMEASURE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "BLOCK_STALE_MEASURE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint32)
func (_MantaServiceManager *MantaServiceManagerSession) BLOCKSTALEMEASURE() (uint32, error) {
	return _MantaServiceManager.Contract.BLOCKSTALEMEASURE(&_MantaServiceManager.CallOpts)
}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint32)
func (_MantaServiceManager *MantaServiceManagerCallerSession) BLOCKSTALEMEASURE() (uint32, error) {
	return _MantaServiceManager.Contract.BLOCKSTALEMEASURE(&_MantaServiceManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MantaServiceManager *MantaServiceManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MantaServiceManager *MantaServiceManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MantaServiceManager.Contract.DEFAULTADMINROLE(&_MantaServiceManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MantaServiceManager *MantaServiceManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MantaServiceManager.Contract.DEFAULTADMINROLE(&_MantaServiceManager.CallOpts)
}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint256)
func (_MantaServiceManager *MantaServiceManagerCaller) THRESHOLDDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "THRESHOLD_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint256)
func (_MantaServiceManager *MantaServiceManagerSession) THRESHOLDDENOMINATOR() (*big.Int, error) {
	return _MantaServiceManager.Contract.THRESHOLDDENOMINATOR(&_MantaServiceManager.CallOpts)
}

// THRESHOLDDENOMINATOR is a free data retrieval call binding the contract method 0xef024458.
//
// Solidity: function THRESHOLD_DENOMINATOR() view returns(uint256)
func (_MantaServiceManager *MantaServiceManagerCallerSession) THRESHOLDDENOMINATOR() (*big.Int, error) {
	return _MantaServiceManager.Contract.THRESHOLDDENOMINATOR(&_MantaServiceManager.CallOpts)
}

// BatchId is a free data retrieval call binding the contract method 0x4972134a.
//
// Solidity: function batchId() view returns(uint32)
func (_MantaServiceManager *MantaServiceManagerCaller) BatchId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "batchId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BatchId is a free data retrieval call binding the contract method 0x4972134a.
//
// Solidity: function batchId() view returns(uint32)
func (_MantaServiceManager *MantaServiceManagerSession) BatchId() (uint32, error) {
	return _MantaServiceManager.Contract.BatchId(&_MantaServiceManager.CallOpts)
}

// BatchId is a free data retrieval call binding the contract method 0x4972134a.
//
// Solidity: function batchId() view returns(uint32)
func (_MantaServiceManager *MantaServiceManagerCallerSession) BatchId() (uint32, error) {
	return _MantaServiceManager.Contract.BatchId(&_MantaServiceManager.CallOpts)
}

// BatchIdToBatchMetadataHash is a free data retrieval call binding the contract method 0xeccbbfc9.
//
// Solidity: function batchIdToBatchMetadataHash(uint32 ) view returns(bytes32)
func (_MantaServiceManager *MantaServiceManagerCaller) BatchIdToBatchMetadataHash(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "batchIdToBatchMetadataHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchIdToBatchMetadataHash is a free data retrieval call binding the contract method 0xeccbbfc9.
//
// Solidity: function batchIdToBatchMetadataHash(uint32 ) view returns(bytes32)
func (_MantaServiceManager *MantaServiceManagerSession) BatchIdToBatchMetadataHash(arg0 uint32) ([32]byte, error) {
	return _MantaServiceManager.Contract.BatchIdToBatchMetadataHash(&_MantaServiceManager.CallOpts, arg0)
}

// BatchIdToBatchMetadataHash is a free data retrieval call binding the contract method 0xeccbbfc9.
//
// Solidity: function batchIdToBatchMetadataHash(uint32 ) view returns(bytes32)
func (_MantaServiceManager *MantaServiceManagerCallerSession) BatchIdToBatchMetadataHash(arg0 uint32) ([32]byte, error) {
	return _MantaServiceManager.Contract.BatchIdToBatchMetadataHash(&_MantaServiceManager.CallOpts, arg0)
}

// FinalityAddress is a free data retrieval call binding the contract method 0xc228fd08.
//
// Solidity: function finalityAddress() view returns(address)
func (_MantaServiceManager *MantaServiceManagerCaller) FinalityAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "finalityAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FinalityAddress is a free data retrieval call binding the contract method 0xc228fd08.
//
// Solidity: function finalityAddress() view returns(address)
func (_MantaServiceManager *MantaServiceManagerSession) FinalityAddress() (common.Address, error) {
	return _MantaServiceManager.Contract.FinalityAddress(&_MantaServiceManager.CallOpts)
}

// FinalityAddress is a free data retrieval call binding the contract method 0xc228fd08.
//
// Solidity: function finalityAddress() view returns(address)
func (_MantaServiceManager *MantaServiceManagerCallerSession) FinalityAddress() (common.Address, error) {
	return _MantaServiceManager.Contract.FinalityAddress(&_MantaServiceManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MantaServiceManager *MantaServiceManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MantaServiceManager *MantaServiceManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MantaServiceManager.Contract.GetRoleAdmin(&_MantaServiceManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MantaServiceManager *MantaServiceManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MantaServiceManager.Contract.GetRoleAdmin(&_MantaServiceManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MantaServiceManager *MantaServiceManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MantaServiceManager *MantaServiceManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MantaServiceManager.Contract.HasRole(&_MantaServiceManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MantaServiceManager *MantaServiceManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MantaServiceManager.Contract.HasRole(&_MantaServiceManager.CallOpts, role, account)
}

// IsBatchConfirmer is a free data retrieval call binding the contract method 0xa5b7890a.
//
// Solidity: function isBatchConfirmer(address ) view returns(bool)
func (_MantaServiceManager *MantaServiceManagerCaller) IsBatchConfirmer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "isBatchConfirmer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchConfirmer is a free data retrieval call binding the contract method 0xa5b7890a.
//
// Solidity: function isBatchConfirmer(address ) view returns(bool)
func (_MantaServiceManager *MantaServiceManagerSession) IsBatchConfirmer(arg0 common.Address) (bool, error) {
	return _MantaServiceManager.Contract.IsBatchConfirmer(&_MantaServiceManager.CallOpts, arg0)
}

// IsBatchConfirmer is a free data retrieval call binding the contract method 0xa5b7890a.
//
// Solidity: function isBatchConfirmer(address ) view returns(bool)
func (_MantaServiceManager *MantaServiceManagerCallerSession) IsBatchConfirmer(arg0 common.Address) (bool, error) {
	return _MantaServiceManager.Contract.IsBatchConfirmer(&_MantaServiceManager.CallOpts, arg0)
}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_MantaServiceManager *MantaServiceManagerCaller) L2OutputOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "l2OutputOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_MantaServiceManager *MantaServiceManagerSession) L2OutputOracle() (common.Address, error) {
	return _MantaServiceManager.Contract.L2OutputOracle(&_MantaServiceManager.CallOpts)
}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_MantaServiceManager *MantaServiceManagerCallerSession) L2OutputOracle() (common.Address, error) {
	return _MantaServiceManager.Contract.L2OutputOracle(&_MantaServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MantaServiceManager *MantaServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MantaServiceManager *MantaServiceManagerSession) Owner() (common.Address, error) {
	return _MantaServiceManager.Contract.Owner(&_MantaServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MantaServiceManager *MantaServiceManagerCallerSession) Owner() (common.Address, error) {
	return _MantaServiceManager.Contract.Owner(&_MantaServiceManager.CallOpts)
}

// QuorumAdversaryThresholdPercentages is a free data retrieval call binding the contract method 0x8687feae.
//
// Solidity: function quorumAdversaryThresholdPercentages() view returns(bytes)
func (_MantaServiceManager *MantaServiceManagerCaller) QuorumAdversaryThresholdPercentages(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "quorumAdversaryThresholdPercentages")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QuorumAdversaryThresholdPercentages is a free data retrieval call binding the contract method 0x8687feae.
//
// Solidity: function quorumAdversaryThresholdPercentages() view returns(bytes)
func (_MantaServiceManager *MantaServiceManagerSession) QuorumAdversaryThresholdPercentages() ([]byte, error) {
	return _MantaServiceManager.Contract.QuorumAdversaryThresholdPercentages(&_MantaServiceManager.CallOpts)
}

// QuorumAdversaryThresholdPercentages is a free data retrieval call binding the contract method 0x8687feae.
//
// Solidity: function quorumAdversaryThresholdPercentages() view returns(bytes)
func (_MantaServiceManager *MantaServiceManagerCallerSession) QuorumAdversaryThresholdPercentages() ([]byte, error) {
	return _MantaServiceManager.Contract.QuorumAdversaryThresholdPercentages(&_MantaServiceManager.CallOpts)
}

// QuorumConfirmationThresholdPercentages is a free data retrieval call binding the contract method 0xbafa9107.
//
// Solidity: function quorumConfirmationThresholdPercentages() view returns(bytes)
func (_MantaServiceManager *MantaServiceManagerCaller) QuorumConfirmationThresholdPercentages(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "quorumConfirmationThresholdPercentages")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QuorumConfirmationThresholdPercentages is a free data retrieval call binding the contract method 0xbafa9107.
//
// Solidity: function quorumConfirmationThresholdPercentages() view returns(bytes)
func (_MantaServiceManager *MantaServiceManagerSession) QuorumConfirmationThresholdPercentages() ([]byte, error) {
	return _MantaServiceManager.Contract.QuorumConfirmationThresholdPercentages(&_MantaServiceManager.CallOpts)
}

// QuorumConfirmationThresholdPercentages is a free data retrieval call binding the contract method 0xbafa9107.
//
// Solidity: function quorumConfirmationThresholdPercentages() view returns(bytes)
func (_MantaServiceManager *MantaServiceManagerCallerSession) QuorumConfirmationThresholdPercentages() ([]byte, error) {
	return _MantaServiceManager.Contract.QuorumConfirmationThresholdPercentages(&_MantaServiceManager.CallOpts)
}

// QuorumNumbersRequired is a free data retrieval call binding the contract method 0xe15234ff.
//
// Solidity: function quorumNumbersRequired() view returns(bytes)
func (_MantaServiceManager *MantaServiceManagerCaller) QuorumNumbersRequired(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "quorumNumbersRequired")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// QuorumNumbersRequired is a free data retrieval call binding the contract method 0xe15234ff.
//
// Solidity: function quorumNumbersRequired() view returns(bytes)
func (_MantaServiceManager *MantaServiceManagerSession) QuorumNumbersRequired() ([]byte, error) {
	return _MantaServiceManager.Contract.QuorumNumbersRequired(&_MantaServiceManager.CallOpts)
}

// QuorumNumbersRequired is a free data retrieval call binding the contract method 0xe15234ff.
//
// Solidity: function quorumNumbersRequired() view returns(bytes)
func (_MantaServiceManager *MantaServiceManagerCallerSession) QuorumNumbersRequired() ([]byte, error) {
	return _MantaServiceManager.Contract.QuorumNumbersRequired(&_MantaServiceManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MantaServiceManager *MantaServiceManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MantaServiceManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MantaServiceManager *MantaServiceManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MantaServiceManager.Contract.SupportsInterface(&_MantaServiceManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MantaServiceManager *MantaServiceManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MantaServiceManager.Contract.SupportsInterface(&_MantaServiceManager.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MantaServiceManager *MantaServiceManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MantaServiceManager *MantaServiceManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.GrantRole(&_MantaServiceManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MantaServiceManager *MantaServiceManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.GrantRole(&_MantaServiceManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_MantaServiceManager *MantaServiceManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_MantaServiceManager *MantaServiceManagerSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.Initialize(&_MantaServiceManager.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_MantaServiceManager *MantaServiceManagerTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.Initialize(&_MantaServiceManager.TransactOpts, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MantaServiceManager *MantaServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantaServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MantaServiceManager *MantaServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _MantaServiceManager.Contract.RenounceOwnership(&_MantaServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MantaServiceManager *MantaServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MantaServiceManager.Contract.RenounceOwnership(&_MantaServiceManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MantaServiceManager *MantaServiceManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MantaServiceManager *MantaServiceManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.RenounceRole(&_MantaServiceManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MantaServiceManager *MantaServiceManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.RenounceRole(&_MantaServiceManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MantaServiceManager *MantaServiceManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MantaServiceManager *MantaServiceManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.RevokeRole(&_MantaServiceManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MantaServiceManager *MantaServiceManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.RevokeRole(&_MantaServiceManager.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MantaServiceManager *MantaServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MantaServiceManager *MantaServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.TransferOwnership(&_MantaServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MantaServiceManager *MantaServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.TransferOwnership(&_MantaServiceManager.TransactOpts, newOwner)
}

// VerifyFinality is a paid mutator transaction binding the contract method 0xe3feaa17.
//
// Solidity: function verifyFinality((bytes32,bytes,bytes,uint32,bytes32,uint256,bytes32,uint256) batchHeader) returns()
func (_MantaServiceManager *MantaServiceManagerTransactor) VerifyFinality(opts *bind.TransactOpts, batchHeader IMantaServiceManagerBatchHeader) (*types.Transaction, error) {
	return _MantaServiceManager.contract.Transact(opts, "verifyFinality", batchHeader)
}

// VerifyFinality is a paid mutator transaction binding the contract method 0xe3feaa17.
//
// Solidity: function verifyFinality((bytes32,bytes,bytes,uint32,bytes32,uint256,bytes32,uint256) batchHeader) returns()
func (_MantaServiceManager *MantaServiceManagerSession) VerifyFinality(batchHeader IMantaServiceManagerBatchHeader) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.VerifyFinality(&_MantaServiceManager.TransactOpts, batchHeader)
}

// VerifyFinality is a paid mutator transaction binding the contract method 0xe3feaa17.
//
// Solidity: function verifyFinality((bytes32,bytes,bytes,uint32,bytes32,uint256,bytes32,uint256) batchHeader) returns()
func (_MantaServiceManager *MantaServiceManagerTransactorSession) VerifyFinality(batchHeader IMantaServiceManagerBatchHeader) (*types.Transaction, error) {
	return _MantaServiceManager.Contract.VerifyFinality(&_MantaServiceManager.TransactOpts, batchHeader)
}

// MantaServiceManagerFinalityVerifiedIterator is returned from FilterFinalityVerified and is used to iterate over the raw logs and unpacked data for FinalityVerified events raised by the MantaServiceManager contract.
type MantaServiceManagerFinalityVerifiedIterator struct {
	Event *MantaServiceManagerFinalityVerified // Event containing the contract specifics and raw log

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
func (it *MantaServiceManagerFinalityVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaServiceManagerFinalityVerified)
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
		it.Event = new(MantaServiceManagerFinalityVerified)
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
func (it *MantaServiceManagerFinalityVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaServiceManagerFinalityVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaServiceManagerFinalityVerified represents a FinalityVerified event raised by the MantaServiceManager contract.
type MantaServiceManagerFinalityVerified struct {
	Proposer      common.Address
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
	L1BlockHash   [32]byte
	L1BlockNumber *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinalityVerified is a free log retrieval operation binding the contract event 0xbea36b7214d7d22729389923a440c2ae7ecb5a9c0039ef0f3659f76f478a52b2.
//
// Solidity: event FinalityVerified(address indexed proposer, bytes32 outputRoot, uint256 l2BlockNumber, bytes32 l1BlockHash, uint256 l1BlockNumber)
func (_MantaServiceManager *MantaServiceManagerFilterer) FilterFinalityVerified(opts *bind.FilterOpts, proposer []common.Address) (*MantaServiceManagerFinalityVerifiedIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _MantaServiceManager.contract.FilterLogs(opts, "FinalityVerified", proposerRule)
	if err != nil {
		return nil, err
	}
	return &MantaServiceManagerFinalityVerifiedIterator{contract: _MantaServiceManager.contract, event: "FinalityVerified", logs: logs, sub: sub}, nil
}

// WatchFinalityVerified is a free log subscription operation binding the contract event 0xbea36b7214d7d22729389923a440c2ae7ecb5a9c0039ef0f3659f76f478a52b2.
//
// Solidity: event FinalityVerified(address indexed proposer, bytes32 outputRoot, uint256 l2BlockNumber, bytes32 l1BlockHash, uint256 l1BlockNumber)
func (_MantaServiceManager *MantaServiceManagerFilterer) WatchFinalityVerified(opts *bind.WatchOpts, sink chan<- *MantaServiceManagerFinalityVerified, proposer []common.Address) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _MantaServiceManager.contract.WatchLogs(opts, "FinalityVerified", proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaServiceManagerFinalityVerified)
				if err := _MantaServiceManager.contract.UnpackLog(event, "FinalityVerified", log); err != nil {
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

// ParseFinalityVerified is a log parse operation binding the contract event 0xbea36b7214d7d22729389923a440c2ae7ecb5a9c0039ef0f3659f76f478a52b2.
//
// Solidity: event FinalityVerified(address indexed proposer, bytes32 outputRoot, uint256 l2BlockNumber, bytes32 l1BlockHash, uint256 l1BlockNumber)
func (_MantaServiceManager *MantaServiceManagerFilterer) ParseFinalityVerified(log types.Log) (*MantaServiceManagerFinalityVerified, error) {
	event := new(MantaServiceManagerFinalityVerified)
	if err := _MantaServiceManager.contract.UnpackLog(event, "FinalityVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MantaServiceManager contract.
type MantaServiceManagerInitializedIterator struct {
	Event *MantaServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *MantaServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaServiceManagerInitialized)
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
		it.Event = new(MantaServiceManagerInitialized)
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
func (it *MantaServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaServiceManagerInitialized represents a Initialized event raised by the MantaServiceManager contract.
type MantaServiceManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MantaServiceManager *MantaServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*MantaServiceManagerInitializedIterator, error) {

	logs, sub, err := _MantaServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MantaServiceManagerInitializedIterator{contract: _MantaServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MantaServiceManager *MantaServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MantaServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _MantaServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaServiceManagerInitialized)
				if err := _MantaServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MantaServiceManager *MantaServiceManagerFilterer) ParseInitialized(log types.Log) (*MantaServiceManagerInitialized, error) {
	event := new(MantaServiceManagerInitialized)
	if err := _MantaServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MantaServiceManager contract.
type MantaServiceManagerOwnershipTransferredIterator struct {
	Event *MantaServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MantaServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaServiceManagerOwnershipTransferred)
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
		it.Event = new(MantaServiceManagerOwnershipTransferred)
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
func (it *MantaServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the MantaServiceManager contract.
type MantaServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MantaServiceManager *MantaServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MantaServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MantaServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MantaServiceManagerOwnershipTransferredIterator{contract: _MantaServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MantaServiceManager *MantaServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MantaServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MantaServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaServiceManagerOwnershipTransferred)
				if err := _MantaServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MantaServiceManager *MantaServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*MantaServiceManagerOwnershipTransferred, error) {
	event := new(MantaServiceManagerOwnershipTransferred)
	if err := _MantaServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaServiceManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MantaServiceManager contract.
type MantaServiceManagerRoleAdminChangedIterator struct {
	Event *MantaServiceManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MantaServiceManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaServiceManagerRoleAdminChanged)
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
		it.Event = new(MantaServiceManagerRoleAdminChanged)
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
func (it *MantaServiceManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaServiceManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaServiceManagerRoleAdminChanged represents a RoleAdminChanged event raised by the MantaServiceManager contract.
type MantaServiceManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MantaServiceManager *MantaServiceManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MantaServiceManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MantaServiceManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MantaServiceManagerRoleAdminChangedIterator{contract: _MantaServiceManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MantaServiceManager *MantaServiceManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MantaServiceManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MantaServiceManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaServiceManagerRoleAdminChanged)
				if err := _MantaServiceManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MantaServiceManager *MantaServiceManagerFilterer) ParseRoleAdminChanged(log types.Log) (*MantaServiceManagerRoleAdminChanged, error) {
	event := new(MantaServiceManagerRoleAdminChanged)
	if err := _MantaServiceManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaServiceManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MantaServiceManager contract.
type MantaServiceManagerRoleGrantedIterator struct {
	Event *MantaServiceManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *MantaServiceManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaServiceManagerRoleGranted)
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
		it.Event = new(MantaServiceManagerRoleGranted)
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
func (it *MantaServiceManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaServiceManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaServiceManagerRoleGranted represents a RoleGranted event raised by the MantaServiceManager contract.
type MantaServiceManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaServiceManager *MantaServiceManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MantaServiceManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaServiceManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MantaServiceManagerRoleGrantedIterator{contract: _MantaServiceManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaServiceManager *MantaServiceManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MantaServiceManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaServiceManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaServiceManagerRoleGranted)
				if err := _MantaServiceManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaServiceManager *MantaServiceManagerFilterer) ParseRoleGranted(log types.Log) (*MantaServiceManagerRoleGranted, error) {
	event := new(MantaServiceManagerRoleGranted)
	if err := _MantaServiceManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantaServiceManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MantaServiceManager contract.
type MantaServiceManagerRoleRevokedIterator struct {
	Event *MantaServiceManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MantaServiceManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantaServiceManagerRoleRevoked)
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
		it.Event = new(MantaServiceManagerRoleRevoked)
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
func (it *MantaServiceManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantaServiceManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantaServiceManagerRoleRevoked represents a RoleRevoked event raised by the MantaServiceManager contract.
type MantaServiceManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaServiceManager *MantaServiceManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MantaServiceManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaServiceManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MantaServiceManagerRoleRevokedIterator{contract: _MantaServiceManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaServiceManager *MantaServiceManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MantaServiceManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MantaServiceManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantaServiceManagerRoleRevoked)
				if err := _MantaServiceManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MantaServiceManager *MantaServiceManagerFilterer) ParseRoleRevoked(log types.Log) (*MantaServiceManagerRoleRevoked, error) {
	event := new(MantaServiceManagerRoleRevoked)
	if err := _MantaServiceManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
