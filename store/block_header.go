package store

import (
	"encoding/json"
)

type BlockHeader struct {
	Hash       []byte `json:"hash"`
	ParentHash []byte `json:"parent_hash"`
	Number     int64  `json:"number"`
	Timestamp  int64  `json:"timestamp"`
}

func (s *Storage) SetBlockHeader(header BlockHeader) error {
	bz, err := json.Marshal(header)
	if err != nil {
		return err
	}
	return s.db.Put(getBlockHeaderKey(header.Number), bz, nil)
}

func (s *Storage) SetBlockHeaders(headers []BlockHeader) error {
	for _, header := range headers {
		err := s.SetBlockHeader(header)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) GetBlockHeader(number int64) (bool, BlockHeader) {
	bhz, err := s.db.Get(getBlockHeaderKey(number), nil)
	if err != nil {
		return handleError2(BlockHeader{}, err)
	}
	var bh BlockHeader
	if err = json.Unmarshal(bhz, &bh); err != nil {
		return false, BlockHeader{}
	}
	return true, bh
}
