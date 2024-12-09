package store

import (
	"encoding/binary"
)

var (
	BlockHeaderKeyPrefix         = []byte{0x01}
	TxMessageKeyPrefix           = []byte{0x02}
	ScannedHeightKeyPrefix       = []byte{0x03}
	NewFinalityProviderKeyPrefix = []byte{0x04}
	CreateBTCDelegationKeyPrefix = []byte{0x05}
	CommitPubRandListKeyPrefix   = []byte{0x06}
	SignatureKeyPrefix           = []byte{0x07}
)

func getBlockHeaderKey(number int64) []byte {
	numberBz := make([]byte, 8)
	binary.BigEndian.PutUint64(numberBz, uint64(number))
	return append(BlockHeaderKeyPrefix, numberBz...)
}

func getSignatureKey(number int64) []byte {
	numberBz := make([]byte, 8)
	binary.BigEndian.PutUint64(numberBz, uint64(number))
	return append(SignatureKeyPrefix, numberBz...)
}

func getTxMessageKey(txHash []byte) []byte {
	return append(TxMessageKeyPrefix, txHash[:]...)
}

func getNewFinalityProviderKey(txHash []byte) []byte {
	return append(NewFinalityProviderKeyPrefix, txHash[:]...)
}

func getCreateBTCDelegationKey(txHash []byte) []byte {
	return append(CreateBTCDelegationKeyPrefix, txHash[:]...)
}

func getCommitPubRandListKey(txHash []byte) []byte {
	return append(CommitPubRandListKeyPrefix, txHash[:]...)
}

func getScannedHeightKey() []byte {
	return ScannedHeightKeyPrefix
}
