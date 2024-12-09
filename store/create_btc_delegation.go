package store

import (
	"encoding/json"
	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
)

type CreateBTCDelegation struct {
	CBD    types.MsgCreateBTCDelegation
	TxHash []byte `json:"tx_hash"`
}

func (s *Storage) SetCreateBTCDelegationMsg(msg CreateBTCDelegation) error {
	bz, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return s.db.Put(getCreateBTCDelegationKey(msg.TxHash), bz, nil)
}

func (s *Storage) GetCreateBTCDelegationrMsg(txHash []byte) (bool, CreateBTCDelegation) {
	cBDB, err := s.db.Get(getCreateBTCDelegationKey(txHash), nil)
	if err != nil {
		return handleError2(CreateBTCDelegation{}, err)
	}
	var cBD CreateBTCDelegation
	if err = json.Unmarshal(cBDB, &cBD); err != nil {
		return false, CreateBTCDelegation{}
	}
	return true, cBD
}
