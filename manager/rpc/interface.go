package rpc

import (
	"math/big"
)

type FinalityInterface interface {
	StateByBlock(L2BlockNumber *big.Int) (interface{}, error)
}

type DRNGRpcInterface interface {
	Finality(req FinalityRequest, reply *interface{}) error
}
