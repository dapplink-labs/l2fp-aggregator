package store

import (
	"encoding/json"
)

type Signature struct {
	BlockNumber     int64  `json:"block_number"`
	TransactionHash []byte `json:"transaction_hash"`
	Data            []byte `json:"data"`
	Timestamp       int64  `json:"timestamp"`
}

func (s *Storage) SetSignature(sign Signature) error {
	bz, err := json.Marshal(sign)
	if err != nil {
		return err
	}
	return s.db.Put(getSignatureKey(sign.BlockNumber), bz, nil)
}

func (s *Storage) GetSignature(BlockNumber int64) (Signature, error) {
	sb, err := s.db.Get(getSignatureKey(BlockNumber), nil)
	if err != nil {
		return handleError(Signature{}, err)
	}
	var sign Signature
	if err = json.Unmarshal(sb, &sign); err != nil {
		return Signature{}, err
	}
	return sign, nil
}
