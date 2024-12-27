package contracts

import (
	"context"
	"github.com/dapplink-labs/l2fp-aggregator/bindings/finality"
	"github.com/dapplink-labs/l2fp-aggregator/store"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type FinalityRelayerManager struct {
	FRMAbi    *abi.ABI
	FRMFilter *finality.FinalityRelayerManagerFilterer
	FRMCtx    context.Context
	log       log.Logger
}

func NewFinalityRelayerManager(log log.Logger) (*FinalityRelayerManager, error) {
	FinalityRelayerManagerAbi, err := finality.FinalityRelayerManagerMetaData.GetAbi()
	if err != nil {
		log.Error("get delegate manager abi fail", "err", err)
		return nil, err
	}

	FinalityRelayerManagerUnpack, err := finality.NewFinalityRelayerManagerFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("new delegation manager fail", "err", err)
		return nil, err
	}

	return &FinalityRelayerManager{
		FRMAbi:    FinalityRelayerManagerAbi,
		FRMFilter: FinalityRelayerManagerUnpack,
		FRMCtx:    context.Background(),
		log:       log,
	}, nil
}

func (frm *FinalityRelayerManager) ProcessFinalityRelayerManagerEvent(db *store.Storage, event store.ContractEvent) error {

	var operatorRegistered *store.OperatorRegistered
	header, err := db.GetEthBlockHeader(int64(event.BlockHeight))
	if err != nil {
		frm.log.Error("ProcessDelegationEvent db Blocks BlockHeader by BlockHash fail", "err", err)
		return err
	}

	if event.EventSignature.String() == frm.FRMAbi.Events["OperatorRegistered"].ID.String() {
		operatorRegisteredEvent, err := frm.FRMFilter.ParseOperatorRegistered(*event.RLPLog)
		if err != nil {
			frm.log.Error("parse operator registered fail", "err", err)
			return err
		}
		log.Info("parse operator registered success",
			"operator", operatorRegisteredEvent.Operator.String())

		operatorRegistered = &store.OperatorRegistered{
			BlockNumber: header.Number,
			TxHash:      event.TransactionHash,
			Operator:    operatorRegisteredEvent.Operator,
			NodeUrl:     operatorRegisteredEvent.NodeUrl,
			Timestamp:   event.Timestamp,
		}
	}

	if operatorRegistered != nil {
		if err = db.SetOperatorRegisteredEvent(*operatorRegistered); err != nil {
			return err
		}
		frm.log.Info("store operator registered success")
	}

	return nil
}
