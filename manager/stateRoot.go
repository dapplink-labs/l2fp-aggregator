package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/eniac-x-labs/manta-relayer/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func (m *Manager) FetchNextOutputInfo(ctx context.Context) (*client.OutputResponse, bool, error) {
	cCtx, cancel := context.WithTimeout(ctx, m.networkTimeout)
	defer cancel()
	callOpts := &bind.CallOpts{
		From:    m.from,
		Context: cCtx,
	}
	nextCheckpointBlock, err := m.l2ooContract.NextBlockNumber(callOpts)
	if err != nil {
		m.log.Error("finality manager unable to get next block number", "err", err)
		return nil, false, err
	}
	LatestSR, err := m.db.StateRoot.LatestStateRoot()
	if err != nil {
		m.log.Error("fail to get latest state root", "err", err)
		return nil, false, err
	}
	if LatestSR != nil {
		if nextCheckpointBlock.Cmp(LatestSR.L2BlockNum) < 0 {
			m.log.Info("next checkpoint block is less than latest checked l2 block")
			return nil, false, errCheckedNextBlock
		}
	}

	//Fetch the current L2 heads
	//cCtx, cancel = context.WithTimeout(ctx, m.networkTimeout)
	//defer cancel()
	//status, err := n.rollupClient.SyncStatus(cCtx)
	//if err != nil {
	//	n.log.Error("proposer unable to get sync status", "err", err)
	//	return nil, false, err
	//}
	//currentBlockNumber := new(big.Int).SetUint64(status.FinalizedL2.Number)

	// Ensure that we do not submit a block in the future
	//if currentBlockNumber.Cmp(nextCheckpointBlock) < 0 {
	//	n.log.Info("proposer submission interval has not elapsed", "currentBlockNumber", currentBlockNumber, "nextBlockNumber", nextCheckpointBlock)
	//	return nil, false, nil
	//}

	return m.fetchOuput(ctx, nextCheckpointBlock)
}

func (m *Manager) fetchOuput(ctx context.Context, block *big.Int) (*client.OutputResponse, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, m.networkTimeout)
	defer cancel()
	output, err := m.rollupClient.OutputAtBlock(ctx, block.Uint64())
	if err != nil {
		m.log.Error(fmt.Sprintf("failed to fetch output at block %d: %w", block.Uint64(), err))
		return nil, false, err
	}

	if output.BlockRef.Number != block.Uint64() { // sanity check, e.g. in case of bad RPC caching
		m.log.Error(fmt.Sprintf("invalid blockNumber: next blockNumber is %v, blockNumber of block is %v", block.Uint64(), output.BlockRef.Number))
		return nil, false, errors.New("invalid blockNumber")
	}

	return output, true, nil
}
