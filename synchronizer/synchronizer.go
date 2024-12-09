package synchronizer

import (
	"context"
	"fmt"
	"github.com/cometbft/cometbft/rpc/client/http"
	types2 "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"
	cTx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/eniac-x-labs/manta-relayer/common/tasks"
	"github.com/eniac-x-labs/manta-relayer/config"
	"github.com/eniac-x-labs/manta-relayer/store"
	"github.com/eniac-x-labs/manta-relayer/synchronizer/node"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

var validMsgTypes = map[string]bool{
	"/babylon.btcstaking.v1.MsgCreateFinalityProvider": true,
	"/babylon.btcstaking.v1.MsgCreateBTCDelegation":    true,
	"/babylon.finality.v1.MsgCommitPubRandList":        true,
}

type Synchronizer struct {
	client            *http.HTTP
	db                *store.Storage
	headers           []types2.Header
	latestHeader      *types2.Header
	headerTraversal   *node.HeaderTraversal
	blockStep         uint64
	startHeight       *big.Int
	confirmationDepth *big.Int
	resourceCtx       context.Context
	resourceCancel    context.CancelFunc
	tasks             tasks.Group
	log               log.Logger
	txMsgChan         chan store.TxMessage
}

func NewSynchronizer(ctx context.Context, cfg *config.Config, db *store.Storage, shutdown context.CancelCauseFunc, logger log.Logger, txMsgChan chan store.TxMessage) (*Synchronizer, error) {

	cli, err := client.NewClientFromNode(cfg.BabylonRpc)
	if err != nil {
		fmt.Printf("Error creating client: %v", err)
	}

	dbLatestHeader, err := db.GetScannedHeight()
	if err != nil {
		return nil, err
	}
	var fromHeader *types2.Header
	if dbLatestHeader != 0 {
		logger.Info("sync detected last indexed block", "number", dbLatestHeader)
		height := int64(dbLatestHeader)
		block, err := cli.Block(ctx, &height)
		if err != nil {
			logger.Info("failed to get babylon block", "height", dbLatestHeader)
		}
		fromHeader = &block.Block.Header
	} else if cfg.StartingHeight > 0 {
		logger.Info("no sync indexed state starting from supplied babylon height", "height", cfg.StartingHeight)
		block, err := cli.Block(ctx, &cfg.StartingHeight)
		if err != nil {
			return nil, fmt.Errorf("could not fetch starting block header: %w", err)
		}
		fromHeader = &block.Block.Header
	} else {
		logger.Info("no ethereum block indexed state")
	}

	headerTraversal := node.NewHeaderTraversal(cli, fromHeader, big.NewInt(0))

	resCtx, resCancel := context.WithCancel(context.Background())
	return &Synchronizer{
		client:          cli,
		blockStep:       cfg.BlockStep,
		headerTraversal: headerTraversal,
		latestHeader:    fromHeader,
		db:              db,
		resourceCtx:     resCtx,
		resourceCancel:  resCancel,
		log:             logger,
		txMsgChan:       txMsgChan,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in Synchronizer: %w", err))
		}},
	}, nil
}

func (syncer *Synchronizer) Start() error {
	tickerSyncer := time.NewTicker(time.Second * 2)
	syncer.tasks.Go(func() error {
		for range tickerSyncer.C {
			if len(syncer.headers) > 0 {
				syncer.log.Info("retrying previous batch")
			} else {
				newHeaders, err := syncer.headerTraversal.NextHeaders(syncer.blockStep)
				if err != nil {
					syncer.log.Error("error querying for headers", "err", err)
					continue
				} else if len(newHeaders) == 0 {
					syncer.log.Warn("no new headers. syncer at head?")
				} else {
					syncer.headers = newHeaders
				}
				latestHeader := syncer.headerTraversal.LatestHeader()
				if latestHeader != nil {
					syncer.log.Info("Latest header", "latestHeader Number", latestHeader.Height)
				}
			}
			err := syncer.processBatch(syncer.headers)
			if err == nil {
				syncer.headers = nil
			}
		}
		return nil
	})
	return nil
}

func (syncer *Synchronizer) processBatch(headers []types2.Header) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHeader := headers[0], headers[len(headers)-1]
	syncer.log.Info("extracting batch", "size", len(headers), "startBlock", firstHeader.Height, "endBlock", lastHeader.Height)

	headerMap := make(map[int64]*types2.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Height] = &header
	}
	blockHeaders := make([]store.BlockHeader, 0, len(headers))
	var txMessages []store.TxMessage
	for i := range headers {
		if headers[i].Hash() == nil {
			continue
		}
		bHeader := store.BlockHeader{
			Hash:       headers[i].Hash(),
			ParentHash: headers[i].LastResultsHash.Bytes(),
			Number:     headers[i].Height,
			Timestamp:  headers[i].Time.Unix(),
		}
		blockHeaders = append(blockHeaders, bHeader)

		block, err := syncer.client.Block(syncer.resourceCtx, &headers[i].Height)
		if err != nil {
			syncer.log.Error("failed to get block", "err", err, "height", headers[i].Height)
			return err
		}
		for _, transaction := range block.Block.Txs {
			var tx cTx.Tx
			tx.Unmarshal(transaction)

			for _, msg := range tx.Body.Messages {
				if validMsgTypes[msg.TypeUrl] {
					if err != nil {
						syncer.log.Error("failed to marshal event", "err", err)
						continue
					}
					txMessage := store.TxMessage{
						BlockHeight:     uint64(block.Block.Height),
						TransactionHash: transaction.Hash(),
						Type:            msg.TypeUrl,
						Data:            msg.Value,
						Timestamp:       time.Now().Unix(),
					}
					txMessages = append(txMessages, txMessage)
					syncer.txMsgChan <- txMessage
				}

			}

		}
	}

	if err := syncer.db.SetBlockHeaders(blockHeaders); err != nil {
		return err
	}
	if err := syncer.db.SetTxMessages(txMessages); err != nil {
		return err
	}
	if err := syncer.db.UpdateHeight(uint64(lastHeader.Height)); err != nil {
		return err
	}

	return nil
}

func (syncer *Synchronizer) Close() error {
	return nil
}
