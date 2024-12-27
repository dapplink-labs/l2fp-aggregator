package store

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

type EthBlockHeader struct {
	Hash       common.Hash `json:"hash"`
	ParentHash common.Hash `json:"parent_hash"`
	Number     int64       `json:"number"`
	Timestamp  int64       `json:"timestamp"`
}

func (s *Storage) SetEthBlockHeader(header EthBlockHeader) error {
	bz, err := json.Marshal(header)
	if err != nil {
		return err
	}
	return s.db.Put(getEthBlockHeaderKey(header.Number), bz, nil)
}

func (s *Storage) SetEthBlockHeaders(headers []EthBlockHeader) error {
	for _, header := range headers {
		err := s.SetEthBlockHeader(header)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) GetEthBlockHeader(number int64) (EthBlockHeader, error) {
	ehb, err := s.db.Get(getEthBlockHeaderKey(number), nil)
	if err != nil {
		return handleError(EthBlockHeader{}, err)
	}
	var eh EthBlockHeader
	if err = json.Unmarshal(ehb, &eh); err != nil {
		return EthBlockHeader{}, err
	}
	return eh, nil
}
