package sdk

import (
	"math/big"
	"net/rpc"

	"github.com/ethereum/go-ethereum/log"

	_rpc "github.com/dapplink-labs/l2fp-aggregator/manager/rpc"
)

type FinalitySDK struct {
	*rpc.Client
}

type SignResponse struct {
	Signature []byte `json:"signature"`
	Message   string `json:"message"`
}

type StakerDelegationResponse struct {
	Amount  uint64 `json:"amount"`
	Message string `json:"message"`
}

func NewFinalitySDK(addr string) (_rpc.FinalityInterface, error) {
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Error("rpc Dial failed", "err", err)
		return nil, err
	}
	return &FinalitySDK{client}, nil
}

func (f *FinalitySDK) SignatureByBlock(block *big.Int) (interface{}, error) {
	var res interface{}
	err := f.Call("FinalityRpcServer.Finality", _rpc.FinalityRequest{
		BlockNumber: block,
	}, &res)
	return &res, err
}

func (f *FinalitySDK) StakerDelegationByAddress(address string) (interface{}, error) {
	var res interface{}
	err := f.Call("FinalityRpcServer.Staker", _rpc.StakerDelegationRequest{
		Address: address,
	}, &res)
	return &res, err
}
