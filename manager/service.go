package manager

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/l2fp-aggregator/manager/sdk"
	"github.com/dapplink-labs/l2fp-aggregator/store"
)

var NotVerifiedBlock = errors.New("the manager has not verified this block")

type FinalityService struct {
	db  *store.Storage
	log log.Logger
}

func NewFinalityService(db *store.Storage, logger log.Logger) *FinalityService {
	return &FinalityService{
		db:  db,
		log: logger,
	}
}

func (f *FinalityService) SignatureByBlock(BlockNumber *big.Int) (interface{}, error) {
	signature, err := f.db.GetSignature(BlockNumber.Int64())
	if err != nil {
		f.log.Error("failed to get signature by block number", "err", err)
		return nil, err
	}
	var bRre []byte
	if signature.Data == nil {
		f.log.Warn("the manager has not verified this block", "blockNumber", BlockNumber.Int64())
		bRre, err = json.Marshal(sdk.SdkResponse{
			Signature: nil,
			Message:   NotVerifiedBlock.Error(),
		})
	} else {
		bRre, err = json.Marshal(sdk.SdkResponse{
			Signature: signature.Data,
			Message:   "successful",
		})
	}
	if err != nil {
		f.log.Error("failed to marshal sdk response", "err", err)
		return nil, err
	}
	return bRre, nil
}
