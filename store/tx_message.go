package store

import (
	"encoding/json"
)

type TxMessage struct {
	BlockHeight     uint64 `json:"block_height"`
	TransactionHash []byte `json:"transaction_hash"`
	Type            string `json:"type"`
	Data            []byte `json:"data"`
	Timestamp       int64  `json:"timestamp"`
}

func (s *Storage) SetTxMessage(msg TxMessage) error {
	bz, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return s.db.Put(getTxMessageKey(msg.TransactionHash), bz, nil)
}

func (s *Storage) SetTxMessages(msgs []TxMessage) error {
	for _, msg := range msgs {
		err := s.SetTxMessage(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) GetTxMessage(txHash []byte) (bool, TxMessage) {
	tmb, err := s.db.Get(getTxMessageKey(txHash), nil)
	if err != nil {
		return handleError2(TxMessage{}, err)
	}
	var tm TxMessage
	if err = json.Unmarshal(tmb, &tm); err != nil {
		return false, TxMessage{}
	}
	return true, tm
}
