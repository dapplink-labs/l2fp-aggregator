package store

import (
	"encoding/json"
	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
)

type CreateFinalityProvider struct {
	FP     types.MsgCreateFinalityProvider
	TxHash []byte `json:"tx_hash"`
}

func (s *Storage) SetCreateFinalityProviderMsg(msg CreateFinalityProvider) error {
	bz, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return s.db.Put(getNewFinalityProviderKey(msg.TxHash), bz, nil)
}

func (s *Storage) GetCreateFinalityProviderMsg(txHash []byte) (bool, CreateFinalityProvider) {
	nFPB, err := s.db.Get(getNewFinalityProviderKey(txHash), nil)
	if err != nil {
		return handleError2(CreateFinalityProvider{}, err)
	}
	var nFP CreateFinalityProvider
	if err = json.Unmarshal(nFPB, &nFP); err != nil {
		return false, CreateFinalityProvider{}
	}
	return true, nFP
}
