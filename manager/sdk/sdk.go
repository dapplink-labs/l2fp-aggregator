package sdk

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/rpc"

	_rpc "github.com/eniac-x-labs/manta-relayer/manager/rpc"
	"github.com/ethereum/go-ethereum/log"
)

type FinalitySDK struct {
	*rpc.Client
}

type SdkResponse struct {
	StateRoot   *common.Hash `json:"state_root"`
	IsFinalized uint8        `json:"is_finalized"`
	Message     string       `json:"message"`
}

func NewFinalitySDK(addr string) (_rpc.FinalityInterface, error) {
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Error("rpc Dial failed", "err", err)
		return nil, err
	}
	return &FinalitySDK{client}, nil
}

func (f *FinalitySDK) StateByBlock(block *big.Int) (interface{}, error) {
	var res interface{}
	err := f.Call("FinalityRpcServer.Finality", _rpc.FinalityRequest{
		L2BlockNumber: block,
	}, &res)
	return &res, err
}
