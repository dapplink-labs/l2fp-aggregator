package rpc

import (
	"math/big"
)

type FinalityInterface interface {
	SignatureByBlock(L2BlockNumber *big.Int) (interface{}, error)
}

type DRNGRpcInterface interface {
	Finality(req FinalityRequest, reply *interface{}) error
}
