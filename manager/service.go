package manager

import (
	"encoding/json"
	"errors"
	"github.com/eniac-x-labs/manta-relayer/database"
	"github.com/eniac-x-labs/manta-relayer/manager/sdk"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

var NotVerifiedBlock = errors.New("the manager has not verified this block")

type FinalityService struct {
	db  *database.DB
	log log.Logger
}

func NewFinalityService(db *database.DB, logger log.Logger) *FinalityService {
	return &FinalityService{
		db:  db,
		log: logger,
	}
}

func (f *FinalityService) StateByBlock(L2BlockNumber *big.Int) (interface{}, error) {
	state, err := f.db.StateRoot.StateRootByL2Block(L2BlockNumber)
	if err != nil {
		f.log.Error("failed to get state root by l2 block number", "err", err)
		return nil, err
	}
	var bRre []byte
	if state == nil {
		f.log.Warn("the manager has not verified this block", "blockNum", L2BlockNumber)
		bRre, err = json.Marshal(sdk.SdkResponse{
			StateRoot:   nil,
			IsFinalized: 0,
			Message:     NotVerifiedBlock.Error(),
		})
	} else {
		bRre, err = json.Marshal(sdk.SdkResponse{
			StateRoot:   &state.StateRoot,
			IsFinalized: state.IsFinalized,
			Message:     "successful",
		})
	}
	if err != nil {
		f.log.Error("failed to marshal sdk response", "err", err)
		return nil, err
	}
	return bRre, nil
}
