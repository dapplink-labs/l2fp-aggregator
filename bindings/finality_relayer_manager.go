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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSApkRegistryFinalityNonSingerAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryFinalityNonSingerAndSignature struct {
	NonSignerPubkeys []BN254G1Point
	ApkG2            BN254G2Point
	Sigma            BN254G1Point
	TotalBtcStake    *big.Int
	TotalMantaStake  *big.Int
}

// IFinalityRelayerManagerFinalityBatch is an auto generated low-level Go binding around an user-defined struct.
type IFinalityRelayerManagerFinalityBatch struct {
	StateRoot       [32]byte
	L2BlockNumber   *big.Int
	L1BlockHash     [32]byte
	L1BlockNumber   *big.Int
	MsgHash         [32]byte
	DisputeGameType uint32
}

// FinalityRelayerManagerMetaData contains all meta data concerning the FinalityRelayerManager contract.
var FinalityRelayerManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"VerifyFinalitySignature\",\"inputs\":[{\"name\":\"finalityBatch\",\"type\":\"tuple\",\"internalType\":\"structIFinalityRelayerManager.FinalityBatch\",\"components\":[{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l1BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l1BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"disputeGameType\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"finalityNonSingerAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.FinalityNonSingerAndSignature\",\"components\":[{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalBtcStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalMantaStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOrRemoverOperatorWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAdd\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deRegisterOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disputeGameFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isDisputeGameFactory\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_disputeGameFactory\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDisputeGameFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2OutputOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorWhitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorWhitelistManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"nodeUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nodeUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifyFinalitySig\",\"inputs\":[{\"name\":\"totalBtcStaking\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalMantaStaking\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b61114d806100d96000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063715018a61161008c578063b9a0634d11610066578063b9a0634d146101ef578063e03c8632146101f7578063f2b4e6171461021a578063f2fde38b1461022d57600080fd5b8063715018a6146101a45780638da5cb5b146101ac578063a1e4c636146101dc57600080fd5b80634a5f825a116100c85780634a5f825a146101405780634d9f1559146101535780634f5b4e151461017e5780635df459461461019157600080fd5b8063097c4af1146100ef5780632a630164146101045780634383371f1461012d575b600080fd5b6101026100fd366004610b0c565b610240565b005b60045461011890600160a01b900460ff1681565b60405190151581526020015b60405180910390f35b61010261013b366004610baa565b61031c565b61010261014e366004610bdd565b61046a565b600354610166906001600160a01b031681565b6040516001600160a01b039091168152602001610124565b600054610166906001600160a01b031681565b600254610166906001600160a01b031681565b6101026105d0565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610166565b6101026101ea366004610d95565b6105e4565b6101026108c4565b610118610205366004610ed0565b60016020526000908152604090205460ff1681565b600454610166906001600160a01b031681565b61010261023b366004610ed0565b610986565b3360009081526001602052604090205460ff166102785760405162461bcd60e51b815260040161026f90610ef2565b60405180910390fd5b6002546040516303682a4560e41b81523360048201526001600160a01b0390911690633682a45090602401600060405180830381600087803b1580156102bd57600080fd5b505af11580156102d1573d6000803e3d6000fd5b50505050336001600160a01b03167f11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f58058383604051610310929190610f6e565b60405180910390a25050565b6000546001600160a01b031633146103b55760405162461bcd60e51b815260206004820152605060248201527f53747261746567794d616e616765722e6f6e6c7946696e616c6974795768697460448201527f654c6973744d616e616765723a206e6f74207468652066696e616c697479207760648201526f3434ba32b634b9ba1036b0b730b3b2b960811b608482015260a40161026f565b6001600160a01b03821661043f5760405162461bcd60e51b815260206004820152604560248201527f46696e616c69747952656c617965724d616e616765722e6164644f706572617460448201527f6f7257686974656c6973743a206f70657261746f722061646472657373206973606482015264207a65726f60d81b608482015260a40161026f565b6001600160a01b03919091166000908152600160205260409020805460ff1916911515919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156104b05750825b905060008267ffffffffffffffff1660011480156104cd5750303b155b9050811580156104db575080155b156104f95760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561052357845460ff60401b1916600160401b1785555b61052c8a6109c4565b600280546001600160a01b03199081166001600160a01b038b81169190911790925560038054909116898316179055600480549188166001600160a81b031990921691909117600160a01b8b15150217905583156105c457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b6105d8610a35565b6105e260006109c4565b565b60025460405163041c37a160e21b815260009182916001600160a01b0390911690631070de84906106259060808901359060208a0135908990600401610fe8565b606060405180830381865afa158015610642573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066691906110a9565b6004549193509150600160a01b900460ff1661077a57600354604080518735602482015260208801356044820152908701356064820152606087013560848201526000916106f0916001600160a01b03909116908690849060a40160408051601f198184030181529190526020810180516001600160e01b031663135556c960e31b179052610a90565b9050806107745760405162461bcd60e51b815260206004820152604660248201527f5374726174656779426173652e56657269667946696e616c6974795369676e6160448201527f747572653a2070726f706f73654c324f7574707574207374617465726f6f742060648201526519985a5b195960d21b608482015260a40161026f565b50610878565b600454600090610801906001600160a01b0316858361079f60c08b0160a08c016110f1565b60405163ffffffff90911660248201528a356044820152606060648201526002608482015261060f60f31b60a482015260c40160408051601f198184030181529190526020810180516001600160e01b0316634176797b60e11b179052610a90565b9050806108765760405162461bcd60e51b815260206004820152603860248201527f5374726174656779426173652e56657269667946696e616c6974795369676e6160448201527f747572653a206372656174652067616d65206661696c65640000000000000000606482015260840161026f565b505b8151602080840151604080519384529183015281018290527f5867a1f09ebc8c9fa2b0ab07694a570b9bb77b2603f5939e40b08b76e49b94e19060600160405180910390a15050505050565b3360009081526001602052604090205460ff166108f35760405162461bcd60e51b815260040161026f90610ef2565b600254604051636c67cc6560e11b81523360048201526001600160a01b039091169063d8cf98ca90602401600060405180830381600087803b15801561093857600080fd5b505af115801561094c573d6000803e3d6000fd5b50506040513381527fb2c38c6252ee2d17f80059fb47a790e20f7bd75e7ba577685375e5484f412d739250602001905060405180910390a1565b61098e610a35565b6001600160a01b0381166109b857604051631e4fbdf760e01b81526000600482015260240161026f565b6109c1816109c4565b50565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b33610a677f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146105e25760405163118cdaa760e01b815233600482015260240161026f565b6000806000610aa0866000610aee565b905080610ad6576308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b600080855160208701888b5af1979650505050505050565b600080603f83619c4001026040850201603f5a021015949350505050565b60008060208385031215610b1f57600080fd5b823567ffffffffffffffff80821115610b3757600080fd5b818501915085601f830112610b4b57600080fd5b813581811115610b5a57600080fd5b866020828501011115610b6c57600080fd5b60209290920196919550909350505050565b80356001600160a01b0381168114610b9557600080fd5b919050565b80358015158114610b9557600080fd5b60008060408385031215610bbd57600080fd5b610bc683610b7e565b9150610bd460208401610b9a565b90509250929050565b600080600080600060a08688031215610bf557600080fd5b610bfe86610b7e565b9450610c0c60208701610b9a565b9350610c1a60408701610b7e565b9250610c2860608701610b7e565b9150610c3660808701610b7e565b90509295509295909350565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610c7b57610c7b610c42565b60405290565b60405160a0810167ffffffffffffffff81118282101715610c7b57610c7b610c42565b604051601f8201601f1916810167ffffffffffffffff81118282101715610ccd57610ccd610c42565b604052919050565b600060408284031215610ce757600080fd5b610cef610c58565b9050813581526020820135602082015292915050565b600082601f830112610d1657600080fd5b610d1e610c58565b806040840185811115610d3057600080fd5b845b81811015610d4a578035845260209384019301610d32565b509095945050505050565b600060808284031215610d6757600080fd5b610d6f610c58565b9050610d7b8383610d05565b8152610d8a8360408401610d05565b602082015292915050565b600080600083850361010080821215610dad57600080fd5b60c0821215610dbb57600080fd5b85945060c0860135915067ffffffffffffffff80831115610ddb57600080fd5b918601916101208389031215610df057600080fd5b610df8610c81565b833582811115610e0757600080fd5b8401601f81018a13610e1857600080fd5b8035602084821115610e2c57610e2c610c42565b610e3a818360051b01610ca4565b828152818101955060069290921b83018101918c831115610e5a57600080fd5b928101925b82841015610e8357610e718d85610cd5565b86528186019550604084019350610e5f565b8452610e918c888301610d55565b81850152505050610ea58960a08601610cd5565b604082015260e084810135606083015292909301356080840152509396909550939092013592915050565b600060208284031215610ee257600080fd5b610eeb82610b7e565b9392505050565b60208082526056908201527f46696e616c69747952656c617965724d616e616765722e72656769737465724f60408201527f70657261746f723a207468697320616464726573732068617665206e6f742070606082015275032b936b4b9b9b4b7b7103a37903932b3b4b9ba32b9160551b608082015260a00190565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b8060005b6002811015610fc0578151845260209384019390910190600101610fa1565b50505050565b610fd1828251610f9d565b6020810151610fe36040840182610f9d565b505050565b838152600060208460208401526040606060408501526101808401855161012060608701528181518084526101a088019150602083019350600092505b808310156110565761104282855180518252602090810151910152565b928501926001929092019190840190611025565b506020880151945061106b6080880186610fc6565b604088015180516101008901526020015161012088015260608801516101408801526080909701516101609096019590955250939695505050505050565b60008082840360608112156110bd57600080fd5b60408112156110cb57600080fd5b506110d4610c58565b835181526020808501519082015260409093015192949293505050565b60006020828403121561110357600080fd5b813563ffffffff81168114610eeb57600080fdfea264697066735822122090a214161715fb3cf640757d222fb4b786288bb0bc3c981ca0224fa818b0f55164736f6c63430008190033",
}

// FinalityRelayerManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FinalityRelayerManagerMetaData.ABI instead.
var FinalityRelayerManagerABI = FinalityRelayerManagerMetaData.ABI

// FinalityRelayerManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FinalityRelayerManagerMetaData.Bin instead.
var FinalityRelayerManagerBin = FinalityRelayerManagerMetaData.Bin

// DeployFinalityRelayerManager deploys a new Ethereum contract, binding an instance of FinalityRelayerManager to it.
func DeployFinalityRelayerManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FinalityRelayerManager, error) {
	parsed, err := FinalityRelayerManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FinalityRelayerManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FinalityRelayerManager{FinalityRelayerManagerCaller: FinalityRelayerManagerCaller{contract: contract}, FinalityRelayerManagerTransactor: FinalityRelayerManagerTransactor{contract: contract}, FinalityRelayerManagerFilterer: FinalityRelayerManagerFilterer{contract: contract}}, nil
}

// FinalityRelayerManager is an auto generated Go binding around an Ethereum contract.
type FinalityRelayerManager struct {
	FinalityRelayerManagerCaller     // Read-only binding to the contract
	FinalityRelayerManagerTransactor // Write-only binding to the contract
	FinalityRelayerManagerFilterer   // Log filterer for contract events
}

// FinalityRelayerManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalityRelayerManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalityRelayerManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalityRelayerManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalityRelayerManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalityRelayerManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalityRelayerManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalityRelayerManagerSession struct {
	Contract     *FinalityRelayerManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FinalityRelayerManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalityRelayerManagerCallerSession struct {
	Contract *FinalityRelayerManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// FinalityRelayerManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalityRelayerManagerTransactorSession struct {
	Contract     *FinalityRelayerManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// FinalityRelayerManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalityRelayerManagerRaw struct {
	Contract *FinalityRelayerManager // Generic contract binding to access the raw methods on
}

// FinalityRelayerManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalityRelayerManagerCallerRaw struct {
	Contract *FinalityRelayerManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FinalityRelayerManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalityRelayerManagerTransactorRaw struct {
	Contract *FinalityRelayerManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalityRelayerManager creates a new instance of FinalityRelayerManager, bound to a specific deployed contract.
func NewFinalityRelayerManager(address common.Address, backend bind.ContractBackend) (*FinalityRelayerManager, error) {
	contract, err := bindFinalityRelayerManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManager{FinalityRelayerManagerCaller: FinalityRelayerManagerCaller{contract: contract}, FinalityRelayerManagerTransactor: FinalityRelayerManagerTransactor{contract: contract}, FinalityRelayerManagerFilterer: FinalityRelayerManagerFilterer{contract: contract}}, nil
}

// NewFinalityRelayerManagerCaller creates a new read-only instance of FinalityRelayerManager, bound to a specific deployed contract.
func NewFinalityRelayerManagerCaller(address common.Address, caller bind.ContractCaller) (*FinalityRelayerManagerCaller, error) {
	contract, err := bindFinalityRelayerManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerCaller{contract: contract}, nil
}

// NewFinalityRelayerManagerTransactor creates a new write-only instance of FinalityRelayerManager, bound to a specific deployed contract.
func NewFinalityRelayerManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalityRelayerManagerTransactor, error) {
	contract, err := bindFinalityRelayerManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerTransactor{contract: contract}, nil
}

// NewFinalityRelayerManagerFilterer creates a new log filterer instance of FinalityRelayerManager, bound to a specific deployed contract.
func NewFinalityRelayerManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalityRelayerManagerFilterer, error) {
	contract, err := bindFinalityRelayerManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerFilterer{contract: contract}, nil
}

// bindFinalityRelayerManager binds a generic wrapper to an already deployed contract.
func bindFinalityRelayerManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FinalityRelayerManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalityRelayerManager *FinalityRelayerManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalityRelayerManager.Contract.FinalityRelayerManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalityRelayerManager *FinalityRelayerManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.FinalityRelayerManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalityRelayerManager *FinalityRelayerManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.FinalityRelayerManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalityRelayerManager *FinalityRelayerManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalityRelayerManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.contract.Transact(opts, method, params...)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) BlsApkRegistry() (common.Address, error) {
	return _FinalityRelayerManager.Contract.BlsApkRegistry(&_FinalityRelayerManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _FinalityRelayerManager.Contract.BlsApkRegistry(&_FinalityRelayerManager.CallOpts)
}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) DisputeGameFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "disputeGameFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) DisputeGameFactory() (common.Address, error) {
	return _FinalityRelayerManager.Contract.DisputeGameFactory(&_FinalityRelayerManager.CallOpts)
}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) DisputeGameFactory() (common.Address, error) {
	return _FinalityRelayerManager.Contract.DisputeGameFactory(&_FinalityRelayerManager.CallOpts)
}

// IsDisputeGameFactory is a free data retrieval call binding the contract method 0x2a630164.
//
// Solidity: function isDisputeGameFactory() view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) IsDisputeGameFactory(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "isDisputeGameFactory")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDisputeGameFactory is a free data retrieval call binding the contract method 0x2a630164.
//
// Solidity: function isDisputeGameFactory() view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) IsDisputeGameFactory() (bool, error) {
	return _FinalityRelayerManager.Contract.IsDisputeGameFactory(&_FinalityRelayerManager.CallOpts)
}

// IsDisputeGameFactory is a free data retrieval call binding the contract method 0x2a630164.
//
// Solidity: function isDisputeGameFactory() view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) IsDisputeGameFactory() (bool, error) {
	return _FinalityRelayerManager.Contract.IsDisputeGameFactory(&_FinalityRelayerManager.CallOpts)
}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) L2OutputOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "l2OutputOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) L2OutputOracle() (common.Address, error) {
	return _FinalityRelayerManager.Contract.L2OutputOracle(&_FinalityRelayerManager.CallOpts)
}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) L2OutputOracle() (common.Address, error) {
	return _FinalityRelayerManager.Contract.L2OutputOracle(&_FinalityRelayerManager.CallOpts)
}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) OperatorWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "operatorWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) OperatorWhitelist(arg0 common.Address) (bool, error) {
	return _FinalityRelayerManager.Contract.OperatorWhitelist(&_FinalityRelayerManager.CallOpts, arg0)
}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) OperatorWhitelist(arg0 common.Address) (bool, error) {
	return _FinalityRelayerManager.Contract.OperatorWhitelist(&_FinalityRelayerManager.CallOpts, arg0)
}

// OperatorWhitelistManager is a free data retrieval call binding the contract method 0x4f5b4e15.
//
// Solidity: function operatorWhitelistManager() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) OperatorWhitelistManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "operatorWhitelistManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorWhitelistManager is a free data retrieval call binding the contract method 0x4f5b4e15.
//
// Solidity: function operatorWhitelistManager() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) OperatorWhitelistManager() (common.Address, error) {
	return _FinalityRelayerManager.Contract.OperatorWhitelistManager(&_FinalityRelayerManager.CallOpts)
}

// OperatorWhitelistManager is a free data retrieval call binding the contract method 0x4f5b4e15.
//
// Solidity: function operatorWhitelistManager() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) OperatorWhitelistManager() (common.Address, error) {
	return _FinalityRelayerManager.Contract.OperatorWhitelistManager(&_FinalityRelayerManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) Owner() (common.Address, error) {
	return _FinalityRelayerManager.Contract.Owner(&_FinalityRelayerManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) Owner() (common.Address, error) {
	return _FinalityRelayerManager.Contract.Owner(&_FinalityRelayerManager.CallOpts)
}

// VerifyFinalitySignature is a paid mutator transaction binding the contract method 0xa1e4c636.
//
// Solidity: function VerifyFinalitySignature((bytes32,uint256,bytes32,uint256,bytes32,uint32) finalityBatch, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) finalityNonSingerAndSignature, uint256 minGas) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) VerifyFinalitySignature(opts *bind.TransactOpts, finalityBatch IFinalityRelayerManagerFinalityBatch, finalityNonSingerAndSignature IBLSApkRegistryFinalityNonSingerAndSignature, minGas *big.Int) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "VerifyFinalitySignature", finalityBatch, finalityNonSingerAndSignature, minGas)
}

// VerifyFinalitySignature is a paid mutator transaction binding the contract method 0xa1e4c636.
//
// Solidity: function VerifyFinalitySignature((bytes32,uint256,bytes32,uint256,bytes32,uint32) finalityBatch, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) finalityNonSingerAndSignature, uint256 minGas) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) VerifyFinalitySignature(finalityBatch IFinalityRelayerManagerFinalityBatch, finalityNonSingerAndSignature IBLSApkRegistryFinalityNonSingerAndSignature, minGas *big.Int) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.VerifyFinalitySignature(&_FinalityRelayerManager.TransactOpts, finalityBatch, finalityNonSingerAndSignature, minGas)
}

// VerifyFinalitySignature is a paid mutator transaction binding the contract method 0xa1e4c636.
//
// Solidity: function VerifyFinalitySignature((bytes32,uint256,bytes32,uint256,bytes32,uint32) finalityBatch, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) finalityNonSingerAndSignature, uint256 minGas) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) VerifyFinalitySignature(finalityBatch IFinalityRelayerManagerFinalityBatch, finalityNonSingerAndSignature IBLSApkRegistryFinalityNonSingerAndSignature, minGas *big.Int) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.VerifyFinalitySignature(&_FinalityRelayerManager.TransactOpts, finalityBatch, finalityNonSingerAndSignature, minGas)
}

// AddOrRemoverOperatorWhitelist is a paid mutator transaction binding the contract method 0x4383371f.
//
// Solidity: function addOrRemoverOperatorWhitelist(address operator, bool isAdd) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) AddOrRemoverOperatorWhitelist(opts *bind.TransactOpts, operator common.Address, isAdd bool) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "addOrRemoverOperatorWhitelist", operator, isAdd)
}

// AddOrRemoverOperatorWhitelist is a paid mutator transaction binding the contract method 0x4383371f.
//
// Solidity: function addOrRemoverOperatorWhitelist(address operator, bool isAdd) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) AddOrRemoverOperatorWhitelist(operator common.Address, isAdd bool) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.AddOrRemoverOperatorWhitelist(&_FinalityRelayerManager.TransactOpts, operator, isAdd)
}

// AddOrRemoverOperatorWhitelist is a paid mutator transaction binding the contract method 0x4383371f.
//
// Solidity: function addOrRemoverOperatorWhitelist(address operator, bool isAdd) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) AddOrRemoverOperatorWhitelist(operator common.Address, isAdd bool) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.AddOrRemoverOperatorWhitelist(&_FinalityRelayerManager.TransactOpts, operator, isAdd)
}

// DeRegisterOperator is a paid mutator transaction binding the contract method 0xb9a0634d.
//
// Solidity: function deRegisterOperator() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) DeRegisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "deRegisterOperator")
}

// DeRegisterOperator is a paid mutator transaction binding the contract method 0xb9a0634d.
//
// Solidity: function deRegisterOperator() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) DeRegisterOperator() (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.DeRegisterOperator(&_FinalityRelayerManager.TransactOpts)
}

// DeRegisterOperator is a paid mutator transaction binding the contract method 0xb9a0634d.
//
// Solidity: function deRegisterOperator() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) DeRegisterOperator() (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.DeRegisterOperator(&_FinalityRelayerManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x4a5f825a.
//
// Solidity: function initialize(address _initialOwner, bool _isDisputeGameFactory, address _blsApkRegistry, address _l2OutputOracle, address _disputeGameFactory) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _isDisputeGameFactory bool, _blsApkRegistry common.Address, _l2OutputOracle common.Address, _disputeGameFactory common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "initialize", _initialOwner, _isDisputeGameFactory, _blsApkRegistry, _l2OutputOracle, _disputeGameFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x4a5f825a.
//
// Solidity: function initialize(address _initialOwner, bool _isDisputeGameFactory, address _blsApkRegistry, address _l2OutputOracle, address _disputeGameFactory) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) Initialize(_initialOwner common.Address, _isDisputeGameFactory bool, _blsApkRegistry common.Address, _l2OutputOracle common.Address, _disputeGameFactory common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.Initialize(&_FinalityRelayerManager.TransactOpts, _initialOwner, _isDisputeGameFactory, _blsApkRegistry, _l2OutputOracle, _disputeGameFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x4a5f825a.
//
// Solidity: function initialize(address _initialOwner, bool _isDisputeGameFactory, address _blsApkRegistry, address _l2OutputOracle, address _disputeGameFactory) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) Initialize(_initialOwner common.Address, _isDisputeGameFactory bool, _blsApkRegistry common.Address, _l2OutputOracle common.Address, _disputeGameFactory common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.Initialize(&_FinalityRelayerManager.TransactOpts, _initialOwner, _isDisputeGameFactory, _blsApkRegistry, _l2OutputOracle, _disputeGameFactory)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x097c4af1.
//
// Solidity: function registerOperator(string nodeUrl) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) RegisterOperator(opts *bind.TransactOpts, nodeUrl string) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "registerOperator", nodeUrl)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x097c4af1.
//
// Solidity: function registerOperator(string nodeUrl) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) RegisterOperator(nodeUrl string) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.RegisterOperator(&_FinalityRelayerManager.TransactOpts, nodeUrl)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x097c4af1.
//
// Solidity: function registerOperator(string nodeUrl) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) RegisterOperator(nodeUrl string) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.RegisterOperator(&_FinalityRelayerManager.TransactOpts, nodeUrl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.RenounceOwnership(&_FinalityRelayerManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.RenounceOwnership(&_FinalityRelayerManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.TransferOwnership(&_FinalityRelayerManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.TransferOwnership(&_FinalityRelayerManager.TransactOpts, newOwner)
}

// FinalityRelayerManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerInitializedIterator struct {
	Event *FinalityRelayerManagerInitialized // Event containing the contract specifics and raw log

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
func (it *FinalityRelayerManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalityRelayerManagerInitialized)
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
		it.Event = new(FinalityRelayerManagerInitialized)
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
func (it *FinalityRelayerManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalityRelayerManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalityRelayerManagerInitialized represents a Initialized event raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*FinalityRelayerManagerInitializedIterator, error) {

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerInitializedIterator{contract: _FinalityRelayerManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FinalityRelayerManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _FinalityRelayerManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalityRelayerManagerInitialized)
				if err := _FinalityRelayerManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseInitialized(log types.Log) (*FinalityRelayerManagerInitialized, error) {
	event := new(FinalityRelayerManagerInitialized)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalityRelayerManagerOperatorDeRegisteredIterator is returned from FilterOperatorDeRegistered and is used to iterate over the raw logs and unpacked data for OperatorDeRegistered events raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOperatorDeRegisteredIterator struct {
	Event *FinalityRelayerManagerOperatorDeRegistered // Event containing the contract specifics and raw log

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
func (it *FinalityRelayerManagerOperatorDeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalityRelayerManagerOperatorDeRegistered)
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
		it.Event = new(FinalityRelayerManagerOperatorDeRegistered)
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
func (it *FinalityRelayerManagerOperatorDeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalityRelayerManagerOperatorDeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalityRelayerManagerOperatorDeRegistered represents a OperatorDeRegistered event raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOperatorDeRegistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeRegistered is a free log retrieval operation binding the contract event 0xb2c38c6252ee2d17f80059fb47a790e20f7bd75e7ba577685375e5484f412d73.
//
// Solidity: event OperatorDeRegistered(address operator)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterOperatorDeRegistered(opts *bind.FilterOpts) (*FinalityRelayerManagerOperatorDeRegisteredIterator, error) {

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "OperatorDeRegistered")
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerOperatorDeRegisteredIterator{contract: _FinalityRelayerManager.contract, event: "OperatorDeRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeRegistered is a free log subscription operation binding the contract event 0xb2c38c6252ee2d17f80059fb47a790e20f7bd75e7ba577685375e5484f412d73.
//
// Solidity: event OperatorDeRegistered(address operator)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) WatchOperatorDeRegistered(opts *bind.WatchOpts, sink chan<- *FinalityRelayerManagerOperatorDeRegistered) (event.Subscription, error) {

	logs, sub, err := _FinalityRelayerManager.contract.WatchLogs(opts, "OperatorDeRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalityRelayerManagerOperatorDeRegistered)
				if err := _FinalityRelayerManager.contract.UnpackLog(event, "OperatorDeRegistered", log); err != nil {
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

// ParseOperatorDeRegistered is a log parse operation binding the contract event 0xb2c38c6252ee2d17f80059fb47a790e20f7bd75e7ba577685375e5484f412d73.
//
// Solidity: event OperatorDeRegistered(address operator)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseOperatorDeRegistered(log types.Log) (*FinalityRelayerManagerOperatorDeRegistered, error) {
	event := new(FinalityRelayerManagerOperatorDeRegistered)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "OperatorDeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalityRelayerManagerOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOperatorRegisteredIterator struct {
	Event *FinalityRelayerManagerOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *FinalityRelayerManagerOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalityRelayerManagerOperatorRegistered)
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
		it.Event = new(FinalityRelayerManagerOperatorRegistered)
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
func (it *FinalityRelayerManagerOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalityRelayerManagerOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalityRelayerManagerOperatorRegistered represents a OperatorRegistered event raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOperatorRegistered struct {
	Operator common.Address
	NodeUrl  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f5805.
//
// Solidity: event OperatorRegistered(address indexed operator, string nodeUrl)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*FinalityRelayerManagerOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerOperatorRegisteredIterator{contract: _FinalityRelayerManager.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f5805.
//
// Solidity: event OperatorRegistered(address indexed operator, string nodeUrl)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *FinalityRelayerManagerOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FinalityRelayerManager.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalityRelayerManagerOperatorRegistered)
				if err := _FinalityRelayerManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f5805.
//
// Solidity: event OperatorRegistered(address indexed operator, string nodeUrl)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseOperatorRegistered(log types.Log) (*FinalityRelayerManagerOperatorRegistered, error) {
	event := new(FinalityRelayerManagerOperatorRegistered)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalityRelayerManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOwnershipTransferredIterator struct {
	Event *FinalityRelayerManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FinalityRelayerManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalityRelayerManagerOwnershipTransferred)
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
		it.Event = new(FinalityRelayerManagerOwnershipTransferred)
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
func (it *FinalityRelayerManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalityRelayerManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalityRelayerManagerOwnershipTransferred represents a OwnershipTransferred event raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FinalityRelayerManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerOwnershipTransferredIterator{contract: _FinalityRelayerManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FinalityRelayerManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinalityRelayerManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalityRelayerManagerOwnershipTransferred)
				if err := _FinalityRelayerManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseOwnershipTransferred(log types.Log) (*FinalityRelayerManagerOwnershipTransferred, error) {
	event := new(FinalityRelayerManagerOwnershipTransferred)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalityRelayerManagerVerifyFinalitySigIterator is returned from FilterVerifyFinalitySig and is used to iterate over the raw logs and unpacked data for VerifyFinalitySig events raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerVerifyFinalitySigIterator struct {
	Event *FinalityRelayerManagerVerifyFinalitySig // Event containing the contract specifics and raw log

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
func (it *FinalityRelayerManagerVerifyFinalitySigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalityRelayerManagerVerifyFinalitySig)
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
		it.Event = new(FinalityRelayerManagerVerifyFinalitySig)
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
func (it *FinalityRelayerManagerVerifyFinalitySigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalityRelayerManagerVerifyFinalitySigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalityRelayerManagerVerifyFinalitySig represents a VerifyFinalitySig event raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerVerifyFinalitySig struct {
	TotalBtcStaking     *big.Int
	TotalMantaStaking   *big.Int
	SignatoryRecordHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterVerifyFinalitySig is a free log retrieval operation binding the contract event 0x5867a1f09ebc8c9fa2b0ab07694a570b9bb77b2603f5939e40b08b76e49b94e1.
//
// Solidity: event VerifyFinalitySig(uint256 totalBtcStaking, uint256 totalMantaStaking, bytes32 signatoryRecordHash)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterVerifyFinalitySig(opts *bind.FilterOpts) (*FinalityRelayerManagerVerifyFinalitySigIterator, error) {

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "VerifyFinalitySig")
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerVerifyFinalitySigIterator{contract: _FinalityRelayerManager.contract, event: "VerifyFinalitySig", logs: logs, sub: sub}, nil
}

// WatchVerifyFinalitySig is a free log subscription operation binding the contract event 0x5867a1f09ebc8c9fa2b0ab07694a570b9bb77b2603f5939e40b08b76e49b94e1.
//
// Solidity: event VerifyFinalitySig(uint256 totalBtcStaking, uint256 totalMantaStaking, bytes32 signatoryRecordHash)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) WatchVerifyFinalitySig(opts *bind.WatchOpts, sink chan<- *FinalityRelayerManagerVerifyFinalitySig) (event.Subscription, error) {

	logs, sub, err := _FinalityRelayerManager.contract.WatchLogs(opts, "VerifyFinalitySig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalityRelayerManagerVerifyFinalitySig)
				if err := _FinalityRelayerManager.contract.UnpackLog(event, "VerifyFinalitySig", log); err != nil {
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

// ParseVerifyFinalitySig is a log parse operation binding the contract event 0x5867a1f09ebc8c9fa2b0ab07694a570b9bb77b2603f5939e40b08b76e49b94e1.
//
// Solidity: event VerifyFinalitySig(uint256 totalBtcStaking, uint256 totalMantaStaking, bytes32 signatoryRecordHash)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseVerifyFinalitySig(log types.Log) (*FinalityRelayerManagerVerifyFinalitySig, error) {
	event := new(FinalityRelayerManagerVerifyFinalitySig)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "VerifyFinalitySig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
