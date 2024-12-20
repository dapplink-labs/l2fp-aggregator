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

// IBLSApkRegistryPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// IBLSApkRegistryStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryStakeTotals struct {
	TotalBtcStaking   *big.Int
	TotalMantaStaking *big.Int
}

// BLSApkRegistryMetaData contains all meta data concerning the BLSApkRegistry contract.
var BLSApkRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"PUBKEY_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addOrRemoverBlsRegisterWhitelist\",\"inputs\":[{\"name\":\"register\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAdd\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"apkHistory\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsRegisterWhitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.FinalityNonSingerAndSignature\",\"components\":[{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalBtcStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalMantaStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.StakeTotals\",\"components\":[{\"name\":\"totalBtcStaking\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalMantaStaking\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentApk\",\"inputs\":[],\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalityRelayerManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFinalityRelayerManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_finalityRelayerManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_relayerManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorToPubkey\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorToPubkeyHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyHashToOperator\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerBLSPublicKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"pubkeyRegistrationMessageHash\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"relayerManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewPubkeyRegistration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAdded\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemoved\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b61240b806100d96000396000f3fe608060405234801561001057600080fd5b506004361061012b5760003560e01c80639feab859116100ad578063d8cf98ca11610071578063d8cf98ca14610345578063de29fac014610358578063e8bb9ae614610378578063ea349348146103a1578063f2fde38b146103b457600080fd5b80639feab859146102c4578063a05274fc146102f9578063bf79ce581461030c578063c0c53b8b1461031f578063c71d36b31461033257600080fd5b80633682a450116100f45780633682a4501461020f578063626e0db514610224578063715018a6146102655780637ff81a871461026d5780638da5cb5b1461028057600080fd5b8062a1f4cb146101305780631070de84146101715780631499b662146101a4578063171f1d5b146101b25780632faea61c146101dc575b600080fd5b61015761013e366004611d49565b6004602052600090815260409020805460019091015482565b604080519283526020830191909152015b60405180910390f35b61018461017f366004611ebe565b6103c7565b604080518351815260209384015193810193909352820152606001610168565b600554600654610157919082565b6101c56101c0366004611fe9565b61068a565b604080519215158352901515602083015201610168565b6101ff6101ea366004611d49565b60086020526000908152604090205460ff1681565b6040519015158152602001610168565b61022261021d366004611d49565b610814565b005b61023761023236600461203a565b6108ab565b6040805167ffffffffffffffff19909416845263ffffffff9283166020850152911690820152606001610168565b6102226108e8565b61018461027b366004611d49565b6108fc565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b039091168152602001610168565b6102eb7f2bd82124057f0913bc3b772ce7b83e8057c1ad1f3510fc83778be20f10ec5de681565b604051908152602001610168565b6000546102ac906001600160a01b031681565b6102eb61031a366004612053565b6109cb565b61022261032d3660046120b0565b610df9565b6001546102ac906001600160a01b031681565b610222610353366004611d49565b610f3b565b6102eb610366366004611d49565b60026020526000908152604090205481565b6102ac61038636600461203a565b6003602052600090815260409020546001600160a01b031681565b6102226103af3660046120f3565b610fd2565b6102226103c2366004611d49565b6110b5565b604080518082019091526000808252602082015260004363ffffffff16841061045d5760405162461bcd60e51b815260206004820152603c60248201527f424c535369676e6174757265436865636b65722e636865636b5369676e61747560448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b0000000060648201526084015b60405180910390fd5b604080518082019091526000808252602082018190526060905b855151811015610527576104b98660000151828151811061049a5761049a61212f565b6020026020010151805160009081526020918201519091526040902090565b8282815181106104cb576104cb61212f565b60200260200101818152505061051d610500876000015183815181106104f3576104f361212f565b60200260200101516110f3565b60408051808201909152600554815260065460208201529061118e565b9250600101610477565b5060008061053f898589602001518a6040015161068a565b91509150816105c25760405162461bcd60e51b815260206004820152604360248201527f424c535369676e6174757265436865636b65722e636865636b5369676e61747560448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610454565b806106355760405162461bcd60e51b815260206004820152603960248201527f424c535369676e6174757265436865636b65722e636865636b5369676e61747560448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610454565b6000888460405160200161064a929190612145565b60408051601f19818403018152828252805160209182012083830190925260608b0151835260808b01519083015290975095505050505050935093915050565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106106d2576106d261212f565b60200201518951600160200201518a602001516000600281106106f7576106f761212f565b60200201518b602001516001600281106107135761071361212f565b602090810291909101518c518d8301516040516107709a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6107939190612181565b90506108066107ac6107a5888461122b565b869061118e565b6107b46112b4565b6107fc6107ed856107e7604080518082018252600080825260209182015281518083019092526001825260029082015290565b9061122b565b6107f68c611374565b9061118e565b886201d4c0611403565b909890975095505050505050565b6000546001600160a01b0316331461083e5760405162461bcd60e51b8152600401610454906121a3565b6000610849826108fc565b5090506108558161161d565b6001600160a01b038216600081815260026020908152604091829020548251938452908301527f1012e33559dd6c66686972872c5e7d4953b63b46a144935dd149c29819b67d5b91015b60405180910390a15050565b600781815481106108bb57600080fd5b600091825260209091200154604081901b915063ffffffff600160c01b8204811691600160e01b90041683565b6108f06117fd565b6108fa6000611858565b565b60408051808201909152600080825260208201526001600160a01b03821660008181526004602090815260408083208151808301835281548152600190910154818401529383526002909152812054909190806109c15760405162461bcd60e51b815260206004820152603e60248201527f424c5341706b52656769737472792e676574526567697374657265645075626b60448201527f65793a206f70657261746f72206973206e6f74207265676973746572656400006064820152608401610454565b9094909350915050565b6001546000906001600160a01b031633146109f85760405162461bcd60e51b815260040161045490612226565b6000610a26610a0f36869003860160408701612298565b805160009081526020918201519091526040902090565b90507fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb58103610aad576040805162461bcd60e51b815260206004820152602481019190915260008051602061239683398151915260448201527f4b65793a2063616e6e6f74207265676973746572207a65726f207075626b65796064820152608401610454565b6001600160a01b03851660009081526002602052604090205415610b375760405162461bcd60e51b8152602060048201526047602482015260008051602061239683398151915260448201527f4b65793a206f70657261746f7220616c72656164792072656769737465726564606482015266207075626b657960c81b608482015260a401610454565b6000818152600360205260409020546001600160a01b031615610bbb5760405162461bcd60e51b8152602060048201526042602482015260008051602061239683398151915260448201527f4b65793a207075626c6963206b657920616c7265616479207265676973746572606482015261195960f21b608482015260a401610454565b604080516000917f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000191610c14918835916020808b0135928b01359160608c01359160808d019160c08e01918d35918e82013591016122b4565b6040516020818303038152906040528051906020012060001c610c379190612181565b9050610cc5610c64610c55836107e7368a90038a0160408b01612298565b6107f636899003890189612298565b610c6c6112b4565b610cae610c9f856107e7604080518082018252600080825260209182015281518083019092526001825260029082015290565b6107f6368a90038a018a612298565b610cc0368a90038a0160808b016122f6565b6118c9565b610d605760405162461bcd60e51b815260206004820152606c602482015260008051602061239683398151915260448201527f4b65793a2065697468657220746865204731207369676e61747572652069732060648201527f77726f6e672c206f7220473120616e642047322070726976617465206b65792060848201526b0c8de40dcdee840dac2e8c6d60a31b60a482015260c401610454565b6001600160a01b03861660008181526004602090815260408083208982018035825560608b01356001909201919091556002835281842087905586845260039092529182902080546001600160a01b0319168417905590517fe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba382804191610de89160808a0190612312565b60405180910390a250949350505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff16600081158015610e3f5750825b905060008267ffffffffffffffff166001148015610e5c5750303b155b905081158015610e6a575080155b15610e885760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610eb257845460ff60401b1916600160401b1785555b610ebb88611858565b600080546001600160a01b03808a166001600160a01b03199283161790925560018054928916929091169190911790558315610f3157845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6000546001600160a01b03163314610f655760405162461bcd60e51b8152600401610454906121a3565b6000610f70826108fc565b509050610f84610f7f826110f3565b61161d565b6001600160a01b038216600081815260026020908152604091829020548251938452908301527fd7fc6178490520705355c7c3941d247da7890efa312e3366fe195e176c0e9e80910161089f565b6001546001600160a01b03163314610ffc5760405162461bcd60e51b815260040161045490612226565b6001600160a01b03821661108a5760405162461bcd60e51b815260206004820152604960248201527f424c5341706b52656769737472792e6164644f7252656d6f766572426c73526560448201527f67697374657257686974656c6973743a206f70657261746f722061646472657360648201526873206973207a65726f60b81b608482015260a401610454565b6001600160a01b03919091166000908152600860205260409020805460ff1916911515919091179055565b6110bd6117fd565b6001600160a01b0381166110e757604051631e4fbdf760e01b815260006004820152602401610454565b6110f081611858565b50565b6040805180820190915260008082526020820152815115801561111857506020820151155b15611136575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206123b683398151915284602001516111699190612181565b611181906000805160206123b6833981519152612352565b905292915050565b919050565b60408051808201909152600080825260208201526111aa611c58565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080806111e557fe5b50806112235760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610454565b505092915050565b6040805180820190915260008082526020820152611247611c76565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808061127657fe5b50806112235760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610454565b6112bc611c94565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806113a46000805160206123b683398151915286612181565b90505b6113b081611b2d565b90935091506000805160206123b683398151915282830983036113e9576040805180820190915290815260208101919091529392505050565b6000805160206123b68339815191526001820890506113a7565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190611435611cb9565b60005b60028110156115f057600061144e82600661236b565b90508482600281106114625761146261212f565b60200201515183611474836000612382565b600c81106114845761148461212f565b602002015284826002811061149b5761149b61212f565b602002015160200151838260016114b29190612382565b600c81106114c2576114c261212f565b60200201528382600281106114d9576114d961212f565b60200201515151836114ec836002612382565b600c81106114fc576114fc61212f565b60200201528382600281106115135761151361212f565b602002015151600160200201518361152c836003612382565b600c811061153c5761153c61212f565b60200201528382600281106115535761155361212f565b60200201516020015160006002811061156e5761156e61212f565b60200201518361157f836004612382565b600c811061158f5761158f61212f565b60200201528382600281106115a6576115a661212f565b6020020151602001516001600281106115c1576115c161212f565b6020020151836115d2836005612382565b600c81106115e2576115e261212f565b602002015250600101611438565b506115f9611cd8565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b604080518082019091526000808252602082015260075460008190036116ab5760405162461bcd60e51b815260206004820152603760248201527f424c5341706b52656769737472792e5f70726f6365737341706b55706461746560448201527f3a2071756f72756d20646f6573206e6f742065786973740000000000000000006064820152608401610454565b60408051808201909152600554815260065460208201526116cc908461118e565b805160058190556020808301805160065560009283525190526040812091935060076116f9600185612352565b815481106117095761170961212f565b6000918252602090912001805490915063ffffffff438116600160c01b90920416036117485780546001600160c01b031916604083901c1781556117f6565b805463ffffffff438116600160e01b8181026001600160e01b039485161785556040805160608101825267ffffffffffffffff198816815260208101938452600081830181815260078054600181018255925291517fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688909101805495519251871690940291909516600160c01b026001600160e01b0319949094169490911c93909317919091179092161790555b5050505050565b3361182f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146108fa5760405163118cdaa760e01b8152336004820152602401610454565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6040805180820182528581526020808201859052825180840190935285835282018390526000916118f8611cb9565b60005b6002811015611ab357600061191182600661236b565b90508482600281106119255761192561212f565b60200201515183611937836000612382565b600c81106119475761194761212f565b602002015284826002811061195e5761195e61212f565b602002015160200151838260016119759190612382565b600c81106119855761198561212f565b602002015283826002811061199c5761199c61212f565b60200201515151836119af836002612382565b600c81106119bf576119bf61212f565b60200201528382600281106119d6576119d661212f565b60200201515160016020020151836119ef836003612382565b600c81106119ff576119ff61212f565b6020020152838260028110611a1657611a1661212f565b602002015160200151600060028110611a3157611a3161212f565b602002015183611a42836004612382565b600c8110611a5257611a5261212f565b6020020152838260028110611a6957611a6961212f565b602002015160200151600160028110611a8457611a8461212f565b602002015183611a95836005612382565b600c8110611aa557611aa561212f565b6020020152506001016118fb565b50611abc611cd8565b60006020826101808560086107d05a03fa90508080611ad757fe5b5080611b1d5760405162461bcd60e51b81526020600482015260156024820152741c185a5c9a5b99cb5bdc18dbd9194b59985a5b1959605a1b6044820152606401610454565b5051151598975050505050505050565b600080806000805160206123b683398151915260036000805160206123b6833981519152866000805160206123b6833981519152888909090890506000611ba3827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206123b6833981519152611baf565b91959194509092505050565b600080611bba611cd8565b611bc2611cf6565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280611bff57fe5b5082611c4d5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610454565b505195945050505050565b60405180608001604052806004906020820280368337509192915050565b60405180606001604052806003906020820280368337509192915050565b6040518060400160405280611ca7611d14565b8152602001611cb4611d14565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b80356001600160a01b038116811461118957600080fd5b600060208284031215611d5b57600080fd5b611d6482611d32565b9392505050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611da457611da4611d6b565b60405290565b60405160a0810167ffffffffffffffff81118282101715611da457611da4611d6b565b604051601f8201601f1916810167ffffffffffffffff81118282101715611df657611df6611d6b565b604052919050565b600060408284031215611e1057600080fd5b611e18611d81565b9050813581526020820135602082015292915050565b600082601f830112611e3f57600080fd5b611e47611d81565b806040840185811115611e5957600080fd5b845b81811015611e73578035845260209384019301611e5b565b509095945050505050565b600060808284031215611e9057600080fd5b611e98611d81565b9050611ea48383611e2e565b8152611eb38360408401611e2e565b602082015292915050565b600080600060608486031215611ed357600080fd5b83359250602080850135925060408086013567ffffffffffffffff80821115611efb57600080fd5b90870190610120828a031215611f1057600080fd5b611f18611daa565b823582811115611f2757600080fd5b8301601f81018b13611f3857600080fd5b803583811115611f4a57611f4a611d6b565b611f58878260051b01611dcd565b818152878101945060069190911b82018701908c821115611f7857600080fd5b918701915b81831015611f9e57611f8f8d84611dfe565b85529387019391860191611f7d565b835250611faf90508a848701611e7e565b85820152611fc08a60a08501611dfe565b604082015260e08301356060820152610100830135608082015280955050505050509250925092565b600080600080610120858703121561200057600080fd5b843593506120118660208701611dfe565b92506120208660608701611e7e565b915061202f8660e08701611dfe565b905092959194509250565b60006020828403121561204c57600080fd5b5035919050565b600080600083850361016081121561206a57600080fd5b61207385611d32565b9350610100601f198201121561208857600080fd5b602085019250604061011f19820112156120a157600080fd5b50610120840190509250925092565b6000806000606084860312156120c557600080fd5b6120ce84611d32565b92506120dc60208501611d32565b91506120ea60408501611d32565b90509250925092565b6000806040838503121561210657600080fd5b61210f83611d32565b91506020830135801515811461212457600080fd5b809150509250929050565b634e487b7160e01b600052603260045260246000fd5b8281526000602080830184516020860160005b8281101561217457815184529284019290840190600101612158565b5091979650505050505050565b60008261219e57634e487b7160e01b600052601260045260246000fd5b500690565b6020808252605c908201527f424c5341706b52656769737472792e6f6e6c7946696e616c69747952656c617960408201527f65724d616e616765723a2063616c6c6572206973206e6f742066696e616c697460608201527f792072656c61796572206d616e6167657220636f6e7472616374732000000000608082015260a00190565b6020808252604c908201527f424c5341706b52656769737472792e6f6e6c7952656c617965724d616e61676560408201527f723a2063616c6c6572206973206e6f74207468652072656c61796572206d616e60608201526b61676572206164647265737360a01b608082015260a00190565b6000604082840312156122aa57600080fd5b611d648383611dfe565b888152876020820152866040820152856060820152604085608083013760408460c0830137610100810192909252610120820152610140019695505050505050565b60006080828403121561230857600080fd5b611d648383611e7e565b823581526020808401359082015260c0810160408381840137604080840160808401379392505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156123655761236561233c565b92915050565b80820281158282048414176123655761236561233c565b808201808211156123655761236561233c56fe424c5341706b52656769737472792e7265676973746572424c535075626c696330644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a264697066735822122059a2dfbd16f31848b02aa7c0362d2427558d58acaa8165ebf6b3e2b9254c9e9964736f6c63430008190033",
}

// BLSApkRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use BLSApkRegistryMetaData.ABI instead.
var BLSApkRegistryABI = BLSApkRegistryMetaData.ABI

// BLSApkRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BLSApkRegistryMetaData.Bin instead.
var BLSApkRegistryBin = BLSApkRegistryMetaData.Bin

// DeployBLSApkRegistry deploys a new Ethereum contract, binding an instance of BLSApkRegistry to it.
func DeployBLSApkRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BLSApkRegistry, error) {
	parsed, err := BLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BLSApkRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BLSApkRegistry{BLSApkRegistryCaller: BLSApkRegistryCaller{contract: contract}, BLSApkRegistryTransactor: BLSApkRegistryTransactor{contract: contract}, BLSApkRegistryFilterer: BLSApkRegistryFilterer{contract: contract}}, nil
}

// BLSApkRegistry is an auto generated Go binding around an Ethereum contract.
type BLSApkRegistry struct {
	BLSApkRegistryCaller     // Read-only binding to the contract
	BLSApkRegistryTransactor // Write-only binding to the contract
	BLSApkRegistryFilterer   // Log filterer for contract events
}

// BLSApkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSApkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSApkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSApkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSApkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSApkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSApkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSApkRegistrySession struct {
	Contract     *BLSApkRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BLSApkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSApkRegistryCallerSession struct {
	Contract *BLSApkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BLSApkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSApkRegistryTransactorSession struct {
	Contract     *BLSApkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BLSApkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSApkRegistryRaw struct {
	Contract *BLSApkRegistry // Generic contract binding to access the raw methods on
}

// BLSApkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSApkRegistryCallerRaw struct {
	Contract *BLSApkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BLSApkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSApkRegistryTransactorRaw struct {
	Contract *BLSApkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLSApkRegistry creates a new instance of BLSApkRegistry, bound to a specific deployed contract.
func NewBLSApkRegistry(address common.Address, backend bind.ContractBackend) (*BLSApkRegistry, error) {
	contract, err := bindBLSApkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistry{BLSApkRegistryCaller: BLSApkRegistryCaller{contract: contract}, BLSApkRegistryTransactor: BLSApkRegistryTransactor{contract: contract}, BLSApkRegistryFilterer: BLSApkRegistryFilterer{contract: contract}}, nil
}

// NewBLSApkRegistryCaller creates a new read-only instance of BLSApkRegistry, bound to a specific deployed contract.
func NewBLSApkRegistryCaller(address common.Address, caller bind.ContractCaller) (*BLSApkRegistryCaller, error) {
	contract, err := bindBLSApkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryCaller{contract: contract}, nil
}

// NewBLSApkRegistryTransactor creates a new write-only instance of BLSApkRegistry, bound to a specific deployed contract.
func NewBLSApkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSApkRegistryTransactor, error) {
	contract, err := bindBLSApkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryTransactor{contract: contract}, nil
}

// NewBLSApkRegistryFilterer creates a new log filterer instance of BLSApkRegistry, bound to a specific deployed contract.
func NewBLSApkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSApkRegistryFilterer, error) {
	contract, err := bindBLSApkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryFilterer{contract: contract}, nil
}

// bindBLSApkRegistry binds a generic wrapper to an already deployed contract.
func bindBLSApkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSApkRegistry *BLSApkRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSApkRegistry.Contract.BLSApkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSApkRegistry *BLSApkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.BLSApkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSApkRegistry *BLSApkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.BLSApkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSApkRegistry *BLSApkRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSApkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSApkRegistry *BLSApkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSApkRegistry *BLSApkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.contract.Transact(opts, method, params...)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistryCaller) PUBKEYREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "PUBKEY_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistrySession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _BLSApkRegistry.Contract.PUBKEYREGISTRATIONTYPEHASH(&_BLSApkRegistry.CallOpts)
}

// PUBKEYREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0x9feab859.
//
// Solidity: function PUBKEY_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) PUBKEYREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _BLSApkRegistry.Contract.PUBKEYREGISTRATIONTYPEHASH(&_BLSApkRegistry.CallOpts)
}

// ApkHistory is a free data retrieval call binding the contract method 0x626e0db5.
//
// Solidity: function apkHistory(uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_BLSApkRegistry *BLSApkRegistryCaller) ApkHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "apkHistory", arg0)

	outstruct := new(struct {
		ApkHash               [24]byte
		UpdateBlockNumber     uint32
		NextUpdateBlockNumber uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ApkHash = *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)
	outstruct.UpdateBlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.NextUpdateBlockNumber = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// ApkHistory is a free data retrieval call binding the contract method 0x626e0db5.
//
// Solidity: function apkHistory(uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_BLSApkRegistry *BLSApkRegistrySession) ApkHistory(arg0 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _BLSApkRegistry.Contract.ApkHistory(&_BLSApkRegistry.CallOpts, arg0)
}

// ApkHistory is a free data retrieval call binding the contract method 0x626e0db5.
//
// Solidity: function apkHistory(uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) ApkHistory(arg0 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _BLSApkRegistry.Contract.ApkHistory(&_BLSApkRegistry.CallOpts, arg0)
}

// BlsRegisterWhitelist is a free data retrieval call binding the contract method 0x2faea61c.
//
// Solidity: function blsRegisterWhitelist(address ) view returns(bool)
func (_BLSApkRegistry *BLSApkRegistryCaller) BlsRegisterWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "blsRegisterWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlsRegisterWhitelist is a free data retrieval call binding the contract method 0x2faea61c.
//
// Solidity: function blsRegisterWhitelist(address ) view returns(bool)
func (_BLSApkRegistry *BLSApkRegistrySession) BlsRegisterWhitelist(arg0 common.Address) (bool, error) {
	return _BLSApkRegistry.Contract.BlsRegisterWhitelist(&_BLSApkRegistry.CallOpts, arg0)
}

// BlsRegisterWhitelist is a free data retrieval call binding the contract method 0x2faea61c.
//
// Solidity: function blsRegisterWhitelist(address ) view returns(bool)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) BlsRegisterWhitelist(arg0 common.Address) (bool, error) {
	return _BLSApkRegistry.Contract.BlsRegisterWhitelist(&_BLSApkRegistry.CallOpts, arg0)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x1070de84.
//
// Solidity: function checkSignatures(bytes32 msgHash, uint256 referenceBlockNumber, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) params) view returns((uint256,uint256), bytes32)
func (_BLSApkRegistry *BLSApkRegistryCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, referenceBlockNumber *big.Int, params IBLSApkRegistryFinalityNonSingerAndSignature) (IBLSApkRegistryStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "checkSignatures", msgHash, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSApkRegistryStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSApkRegistryStakeTotals)).(*IBLSApkRegistryStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x1070de84.
//
// Solidity: function checkSignatures(bytes32 msgHash, uint256 referenceBlockNumber, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) params) view returns((uint256,uint256), bytes32)
func (_BLSApkRegistry *BLSApkRegistrySession) CheckSignatures(msgHash [32]byte, referenceBlockNumber *big.Int, params IBLSApkRegistryFinalityNonSingerAndSignature) (IBLSApkRegistryStakeTotals, [32]byte, error) {
	return _BLSApkRegistry.Contract.CheckSignatures(&_BLSApkRegistry.CallOpts, msgHash, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x1070de84.
//
// Solidity: function checkSignatures(bytes32 msgHash, uint256 referenceBlockNumber, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) params) view returns((uint256,uint256), bytes32)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) CheckSignatures(msgHash [32]byte, referenceBlockNumber *big.Int, params IBLSApkRegistryFinalityNonSingerAndSignature) (IBLSApkRegistryStakeTotals, [32]byte, error) {
	return _BLSApkRegistry.Contract.CheckSignatures(&_BLSApkRegistry.CallOpts, msgHash, referenceBlockNumber, params)
}

// CurrentApk is a free data retrieval call binding the contract method 0x1499b662.
//
// Solidity: function currentApk() view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistryCaller) CurrentApk(opts *bind.CallOpts) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "currentApk")

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentApk is a free data retrieval call binding the contract method 0x1499b662.
//
// Solidity: function currentApk() view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistrySession) CurrentApk() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _BLSApkRegistry.Contract.CurrentApk(&_BLSApkRegistry.CallOpts)
}

// CurrentApk is a free data retrieval call binding the contract method 0x1499b662.
//
// Solidity: function currentApk() view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) CurrentApk() (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _BLSApkRegistry.Contract.CurrentApk(&_BLSApkRegistry.CallOpts)
}

// FinalityRelayerManager is a free data retrieval call binding the contract method 0xa05274fc.
//
// Solidity: function finalityRelayerManager() view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCaller) FinalityRelayerManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "finalityRelayerManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FinalityRelayerManager is a free data retrieval call binding the contract method 0xa05274fc.
//
// Solidity: function finalityRelayerManager() view returns(address)
func (_BLSApkRegistry *BLSApkRegistrySession) FinalityRelayerManager() (common.Address, error) {
	return _BLSApkRegistry.Contract.FinalityRelayerManager(&_BLSApkRegistry.CallOpts)
}

// FinalityRelayerManager is a free data retrieval call binding the contract method 0xa05274fc.
//
// Solidity: function finalityRelayerManager() view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) FinalityRelayerManager() (common.Address, error) {
	return _BLSApkRegistry.Contract.FinalityRelayerManager(&_BLSApkRegistry.CallOpts)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_BLSApkRegistry *BLSApkRegistryCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "getRegisteredPubkey", operator)

	if err != nil {
		return *new(BN254G1Point), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_BLSApkRegistry *BLSApkRegistrySession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _BLSApkRegistry.Contract.GetRegisteredPubkey(&_BLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _BLSApkRegistry.Contract.GetRegisteredPubkey(&_BLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address ) view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistryCaller) OperatorToPubkey(opts *bind.CallOpts, arg0 common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "operatorToPubkey", arg0)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address ) view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistrySession) OperatorToPubkey(arg0 common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _BLSApkRegistry.Contract.OperatorToPubkey(&_BLSApkRegistry.CallOpts, arg0)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address ) view returns(uint256 X, uint256 Y)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) OperatorToPubkey(arg0 common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _BLSApkRegistry.Contract.OperatorToPubkey(&_BLSApkRegistry.CallOpts, arg0)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistryCaller) OperatorToPubkeyHash(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "operatorToPubkeyHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistrySession) OperatorToPubkeyHash(arg0 common.Address) ([32]byte, error) {
	return _BLSApkRegistry.Contract.OperatorToPubkeyHash(&_BLSApkRegistry.CallOpts, arg0)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) OperatorToPubkeyHash(arg0 common.Address) ([32]byte, error) {
	return _BLSApkRegistry.Contract.OperatorToPubkeyHash(&_BLSApkRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BLSApkRegistry *BLSApkRegistrySession) Owner() (common.Address, error) {
	return _BLSApkRegistry.Contract.Owner(&_BLSApkRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) Owner() (common.Address, error) {
	return _BLSApkRegistry.Contract.Owner(&_BLSApkRegistry.CallOpts)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCaller) PubkeyHashToOperator(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "pubkeyHashToOperator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_BLSApkRegistry *BLSApkRegistrySession) PubkeyHashToOperator(arg0 [32]byte) (common.Address, error) {
	return _BLSApkRegistry.Contract.PubkeyHashToOperator(&_BLSApkRegistry.CallOpts, arg0)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) PubkeyHashToOperator(arg0 [32]byte) (common.Address, error) {
	return _BLSApkRegistry.Contract.PubkeyHashToOperator(&_BLSApkRegistry.CallOpts, arg0)
}

// RelayerManager is a free data retrieval call binding the contract method 0xc71d36b3.
//
// Solidity: function relayerManager() view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCaller) RelayerManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "relayerManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerManager is a free data retrieval call binding the contract method 0xc71d36b3.
//
// Solidity: function relayerManager() view returns(address)
func (_BLSApkRegistry *BLSApkRegistrySession) RelayerManager() (common.Address, error) {
	return _BLSApkRegistry.Contract.RelayerManager(&_BLSApkRegistry.CallOpts)
}

// RelayerManager is a free data retrieval call binding the contract method 0xc71d36b3.
//
// Solidity: function relayerManager() view returns(address)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) RelayerManager() (common.Address, error) {
	return _BLSApkRegistry.Contract.RelayerManager(&_BLSApkRegistry.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_BLSApkRegistry *BLSApkRegistryCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _BLSApkRegistry.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_BLSApkRegistry *BLSApkRegistrySession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _BLSApkRegistry.Contract.TrySignatureAndApkVerification(&_BLSApkRegistry.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_BLSApkRegistry *BLSApkRegistryCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _BLSApkRegistry.Contract.TrySignatureAndApkVerification(&_BLSApkRegistry.CallOpts, msgHash, apk, apkG2, sigma)
}

// AddOrRemoverBlsRegisterWhitelist is a paid mutator transaction binding the contract method 0xea349348.
//
// Solidity: function addOrRemoverBlsRegisterWhitelist(address register, bool isAdd) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactor) AddOrRemoverBlsRegisterWhitelist(opts *bind.TransactOpts, register common.Address, isAdd bool) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "addOrRemoverBlsRegisterWhitelist", register, isAdd)
}

// AddOrRemoverBlsRegisterWhitelist is a paid mutator transaction binding the contract method 0xea349348.
//
// Solidity: function addOrRemoverBlsRegisterWhitelist(address register, bool isAdd) returns()
func (_BLSApkRegistry *BLSApkRegistrySession) AddOrRemoverBlsRegisterWhitelist(register common.Address, isAdd bool) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.AddOrRemoverBlsRegisterWhitelist(&_BLSApkRegistry.TransactOpts, register, isAdd)
}

// AddOrRemoverBlsRegisterWhitelist is a paid mutator transaction binding the contract method 0xea349348.
//
// Solidity: function addOrRemoverBlsRegisterWhitelist(address register, bool isAdd) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) AddOrRemoverBlsRegisterWhitelist(register common.Address, isAdd bool) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.AddOrRemoverBlsRegisterWhitelist(&_BLSApkRegistry.TransactOpts, register, isAdd)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address operator) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "deregisterOperator", operator)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address operator) returns()
func (_BLSApkRegistry *BLSApkRegistrySession) DeregisterOperator(operator common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.DeregisterOperator(&_BLSApkRegistry.TransactOpts, operator)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xd8cf98ca.
//
// Solidity: function deregisterOperator(address operator) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) DeregisterOperator(operator common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.DeregisterOperator(&_BLSApkRegistry.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _initialOwner, address _finalityRelayerManager, address _relayerManager) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _finalityRelayerManager common.Address, _relayerManager common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "initialize", _initialOwner, _finalityRelayerManager, _relayerManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _initialOwner, address _finalityRelayerManager, address _relayerManager) returns()
func (_BLSApkRegistry *BLSApkRegistrySession) Initialize(_initialOwner common.Address, _finalityRelayerManager common.Address, _relayerManager common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.Initialize(&_BLSApkRegistry.TransactOpts, _initialOwner, _finalityRelayerManager, _relayerManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _initialOwner, address _finalityRelayerManager, address _relayerManager) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) Initialize(_initialOwner common.Address, _finalityRelayerManager common.Address, _relayerManager common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.Initialize(&_BLSApkRegistry.TransactOpts, _initialOwner, _finalityRelayerManager, _relayerManager)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistryTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "registerBLSPublicKey", operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistrySession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.RegisterBLSPublicKey(&_BLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32)
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.RegisterBLSPublicKey(&_BLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address operator) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "registerOperator", operator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address operator) returns()
func (_BLSApkRegistry *BLSApkRegistrySession) RegisterOperator(operator common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.RegisterOperator(&_BLSApkRegistry.TransactOpts, operator)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3682a450.
//
// Solidity: function registerOperator(address operator) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) RegisterOperator(operator common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.RegisterOperator(&_BLSApkRegistry.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BLSApkRegistry *BLSApkRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BLSApkRegistry *BLSApkRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.RenounceOwnership(&_BLSApkRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.RenounceOwnership(&_BLSApkRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BLSApkRegistry *BLSApkRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.TransferOwnership(&_BLSApkRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BLSApkRegistry *BLSApkRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BLSApkRegistry.Contract.TransferOwnership(&_BLSApkRegistry.TransactOpts, newOwner)
}

// BLSApkRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BLSApkRegistry contract.
type BLSApkRegistryInitializedIterator struct {
	Event *BLSApkRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *BLSApkRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegistryInitialized)
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
		it.Event = new(BLSApkRegistryInitialized)
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
func (it *BLSApkRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegistryInitialized represents a Initialized event raised by the BLSApkRegistry contract.
type BLSApkRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BLSApkRegistry *BLSApkRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*BLSApkRegistryInitializedIterator, error) {

	logs, sub, err := _BLSApkRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryInitializedIterator{contract: _BLSApkRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BLSApkRegistry *BLSApkRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BLSApkRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _BLSApkRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegistryInitialized)
				if err := _BLSApkRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BLSApkRegistry *BLSApkRegistryFilterer) ParseInitialized(log types.Log) (*BLSApkRegistryInitialized, error) {
	event := new(BLSApkRegistryInitialized)
	if err := _BLSApkRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSApkRegistryNewPubkeyRegistrationIterator is returned from FilterNewPubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewPubkeyRegistration events raised by the BLSApkRegistry contract.
type BLSApkRegistryNewPubkeyRegistrationIterator struct {
	Event *BLSApkRegistryNewPubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *BLSApkRegistryNewPubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegistryNewPubkeyRegistration)
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
		it.Event = new(BLSApkRegistryNewPubkeyRegistration)
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
func (it *BLSApkRegistryNewPubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegistryNewPubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegistryNewPubkeyRegistration represents a NewPubkeyRegistration event raised by the BLSApkRegistry contract.
type BLSApkRegistryNewPubkeyRegistration struct {
	Operator common.Address
	PubkeyG1 BN254G1Point
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPubkeyRegistration is a free log retrieval operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_BLSApkRegistry *BLSApkRegistryFilterer) FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*BLSApkRegistryNewPubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BLSApkRegistry.contract.FilterLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryNewPubkeyRegistrationIterator{contract: _BLSApkRegistry.contract, event: "NewPubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewPubkeyRegistration is a free log subscription operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_BLSApkRegistry *BLSApkRegistryFilterer) WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *BLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BLSApkRegistry.contract.WatchLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegistryNewPubkeyRegistration)
				if err := _BLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
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

// ParseNewPubkeyRegistration is a log parse operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_BLSApkRegistry *BLSApkRegistryFilterer) ParseNewPubkeyRegistration(log types.Log) (*BLSApkRegistryNewPubkeyRegistration, error) {
	event := new(BLSApkRegistryNewPubkeyRegistration)
	if err := _BLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSApkRegistryOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the BLSApkRegistry contract.
type BLSApkRegistryOperatorAddedIterator struct {
	Event *BLSApkRegistryOperatorAdded // Event containing the contract specifics and raw log

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
func (it *BLSApkRegistryOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegistryOperatorAdded)
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
		it.Event = new(BLSApkRegistryOperatorAdded)
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
func (it *BLSApkRegistryOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegistryOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegistryOperatorAdded represents a OperatorAdded event raised by the BLSApkRegistry contract.
type BLSApkRegistryOperatorAdded struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x1012e33559dd6c66686972872c5e7d4953b63b46a144935dd149c29819b67d5b.
//
// Solidity: event OperatorAdded(address operator, bytes32 operatorId)
func (_BLSApkRegistry *BLSApkRegistryFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*BLSApkRegistryOperatorAddedIterator, error) {

	logs, sub, err := _BLSApkRegistry.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryOperatorAddedIterator{contract: _BLSApkRegistry.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x1012e33559dd6c66686972872c5e7d4953b63b46a144935dd149c29819b67d5b.
//
// Solidity: event OperatorAdded(address operator, bytes32 operatorId)
func (_BLSApkRegistry *BLSApkRegistryFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *BLSApkRegistryOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _BLSApkRegistry.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegistryOperatorAdded)
				if err := _BLSApkRegistry.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0x1012e33559dd6c66686972872c5e7d4953b63b46a144935dd149c29819b67d5b.
//
// Solidity: event OperatorAdded(address operator, bytes32 operatorId)
func (_BLSApkRegistry *BLSApkRegistryFilterer) ParseOperatorAdded(log types.Log) (*BLSApkRegistryOperatorAdded, error) {
	event := new(BLSApkRegistryOperatorAdded)
	if err := _BLSApkRegistry.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSApkRegistryOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the BLSApkRegistry contract.
type BLSApkRegistryOperatorRemovedIterator struct {
	Event *BLSApkRegistryOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *BLSApkRegistryOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegistryOperatorRemoved)
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
		it.Event = new(BLSApkRegistryOperatorRemoved)
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
func (it *BLSApkRegistryOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegistryOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegistryOperatorRemoved represents a OperatorRemoved event raised by the BLSApkRegistry contract.
type BLSApkRegistryOperatorRemoved struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0xd7fc6178490520705355c7c3941d247da7890efa312e3366fe195e176c0e9e80.
//
// Solidity: event OperatorRemoved(address operator, bytes32 operatorId)
func (_BLSApkRegistry *BLSApkRegistryFilterer) FilterOperatorRemoved(opts *bind.FilterOpts) (*BLSApkRegistryOperatorRemovedIterator, error) {

	logs, sub, err := _BLSApkRegistry.contract.FilterLogs(opts, "OperatorRemoved")
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryOperatorRemovedIterator{contract: _BLSApkRegistry.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0xd7fc6178490520705355c7c3941d247da7890efa312e3366fe195e176c0e9e80.
//
// Solidity: event OperatorRemoved(address operator, bytes32 operatorId)
func (_BLSApkRegistry *BLSApkRegistryFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *BLSApkRegistryOperatorRemoved) (event.Subscription, error) {

	logs, sub, err := _BLSApkRegistry.contract.WatchLogs(opts, "OperatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegistryOperatorRemoved)
				if err := _BLSApkRegistry.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0xd7fc6178490520705355c7c3941d247da7890efa312e3366fe195e176c0e9e80.
//
// Solidity: event OperatorRemoved(address operator, bytes32 operatorId)
func (_BLSApkRegistry *BLSApkRegistryFilterer) ParseOperatorRemoved(log types.Log) (*BLSApkRegistryOperatorRemoved, error) {
	event := new(BLSApkRegistryOperatorRemoved)
	if err := _BLSApkRegistry.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSApkRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BLSApkRegistry contract.
type BLSApkRegistryOwnershipTransferredIterator struct {
	Event *BLSApkRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BLSApkRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSApkRegistryOwnershipTransferred)
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
		it.Event = new(BLSApkRegistryOwnershipTransferred)
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
func (it *BLSApkRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSApkRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSApkRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the BLSApkRegistry contract.
type BLSApkRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BLSApkRegistry *BLSApkRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BLSApkRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BLSApkRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BLSApkRegistryOwnershipTransferredIterator{contract: _BLSApkRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BLSApkRegistry *BLSApkRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BLSApkRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BLSApkRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSApkRegistryOwnershipTransferred)
				if err := _BLSApkRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BLSApkRegistry *BLSApkRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*BLSApkRegistryOwnershipTransferred, error) {
	event := new(BLSApkRegistryOwnershipTransferred)
	if err := _BLSApkRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
