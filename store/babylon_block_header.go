package store

import (
	"encoding/json"
)

type BabylonBlockHeader struct {
	Hash       []byte `json:"hash"`
	ParentHash []byte `json:"parent_hash"`
	Number     int64  `json:"number"`
	Timestamp  int64  `json:"timestamp"`
}

func (s *Storage) SetBabylonBlockHeader(header BabylonBlockHeader) error {
	bz, err := json.Marshal(header)
	if err != nil {
		return err
	}
	return s.db.Put(getBabylonBlockHeaderKey(header.Number), bz, nil)
}

func (s *Storage) SetBabylonBlockHeaders(headers []BabylonBlockHeader) error {
	for _, header := range headers {
		err := s.SetBabylonBlockHeader(header)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) GetBabylonBlockHeader(number int64) (bool, BabylonBlockHeader) {
	bhz, err := s.db.Get(getBabylonBlockHeaderKey(number), nil)
	if err != nil {
		return handleError2(BabylonBlockHeader{}, err)
	}
	var bh BabylonBlockHeader
	if err = json.Unmarshal(bhz, &bh); err != nil {
		return false, BabylonBlockHeader{}
	}
	return true, bh
}
