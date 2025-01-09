package synchronizer

import (
	"github.com/ethereum/go-ethereum/log"

	"github.com/Manta-Network/manta-fp-aggregator/store"
	"github.com/Manta-Network/manta-fp-aggregator/synchronizer/contracts"
)

type EthEventProcess struct {
	db  *store.Storage
	log log.Logger

	contractEventChan chan store.ContractEvent

	finalityRelayerManager *contracts.FinalityRelayerManager
}

func NewEthEventProcess(db *store.Storage, logger log.Logger, contractEventChan chan store.ContractEvent) (*EthEventProcess, error) {
	finalityRelayerManager, err := contracts.NewFinalityRelayerManager(logger)
	if err != nil {
		logger.Error("new finality relayer manager fail", "err", err)
		return nil, err
	}

	return &EthEventProcess{
		db:                     db,
		log:                    logger,
		contractEventChan:      contractEventChan,
		finalityRelayerManager: finalityRelayerManager,
	}, nil
}

func (e *EthEventProcess) Start() error {
	for {
		select {
		case event := <-e.contractEventChan:
			if err := e.finalityRelayerManager.ProcessFinalityRelayerManagerEvent(e.db, event); err != nil {
				e.log.Error("failed to process FinalityRelayerManager event", "err", err)
				continue
			}

		}
	}

}
