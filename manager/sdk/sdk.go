package sdk

import (
	"math/big"
	"net/rpc"

	_rpc "github.com/eniac-x-labs/manta-relayer/manager/rpc"
	"github.com/ethereum/go-ethereum/log"
)

type FinalitySDK struct {
	*rpc.Client
}

type SdkResponse struct {
	Signature []byte `json:"signature"`
	Message   string `json:"message"`
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
