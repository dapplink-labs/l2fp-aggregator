package store

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractEvent struct {
	BlockHeight     uint64         `json:"block_height"`
	ContractAddress common.Address `json:"contract_address"`
	TransactionHash common.Hash    `json:"transaction_hash"`
	LogIndex        uint64         `json:"log_index"`
	EventSignature  common.Hash    `json:"event_signature"`
	RLPLog          *types.Log     `json:"rlp_log"`
	Timestamp       uint64         `json:"timestamp"`
}

func ContractEventFromLog(log *types.Log, timestamp uint64) ContractEvent {
	eventSig := common.Hash{}
	if len(log.Topics) > 0 {
		eventSig = log.Topics[0]
	}
	return ContractEvent{
		BlockHeight:     log.BlockNumber,
		TransactionHash: log.TxHash,
		ContractAddress: log.Address,
		EventSignature:  eventSig,
		LogIndex:        uint64(log.Index),
		RLPLog:          log,
		Timestamp:       timestamp,
	}
}

func (s *Storage) SetContractEvent(event ContractEvent) error {
	bz, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return s.db.Put(getContractEventKey(event.TransactionHash.Bytes()), bz, nil)
}

func (s *Storage) SetContractEvents(events []ContractEvent) error {
	for _, event := range events {
		err := s.SetContractEvent(event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) GetContractEvent(txHash []byte) (bool, ContractEvent) {
	ceb, err := s.db.Get(getContractEventKey(txHash), nil)
	if err != nil {
		return handleError2(ContractEvent{}, err)
	}
	var ce ContractEvent
	if err = json.Unmarshal(ceb, &ce); err != nil {
		return false, ContractEvent{}
	}
	return true, ce
}
