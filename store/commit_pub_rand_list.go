package store

import (
	"encoding/json"
	"github.com/babylonlabs-io/babylon/x/finality/types"
)

type CommitPubRandList struct {
	CPR    types.MsgCommitPubRandList
	TxHash []byte `json:"tx_hash"`
}

func (s *Storage) SetCommitPubRandListMsg(event CommitPubRandList) error {
	bz, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return s.db.Put(getCommitPubRandListKey(event.TxHash), bz, nil)
}

func (s *Storage) GetCommitPubRandListMsg(txHash []byte) (bool, CommitPubRandList) {
	CPRB, err := s.db.Get(getCommitPubRandListKey(txHash), nil)
	if err != nil {
		return handleError2(CommitPubRandList{}, err)
	}
	var cPR CommitPubRandList
	if err = json.Unmarshal(CPRB, &cPR); err != nil {
		return false, CommitPubRandList{}
	}
	return true, cPR
}
