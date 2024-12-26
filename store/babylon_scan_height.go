package store

import (
	"encoding/binary"
)

func (s *Storage) UpdateBabylonHeight(height uint64) error {
	heightBz := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBz, height)
	return s.db.Put(getBabylonScannedHeightKey(), heightBz, nil)
}

func (s *Storage) GetBabylonScannedHeight() (uint64, error) {
	bz, err := s.db.Get(getBabylonScannedHeightKey(), nil)
	if err != nil {
		return handleError(uint64(0), err)
	}
	return binary.BigEndian.Uint64(bz), nil
}

func (s *Storage) ResetBabylonScanHeight(height uint64) error {
	heightBz := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBz, height)
	return s.db.Put(getBabylonScannedHeightKey(), heightBz, nil)
}
