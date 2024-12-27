package node

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/cometbft/cometbft/rpc/client/http"
	types2 "github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/l2fp-aggregator/common/bigint"
)

type BabylonHeaderTraversal struct {
	babylonClient *http.HTTP

	latestHeader        *types2.Header
	lastTraversedHeader *types2.Header

	blockConfirmationDepth *big.Int
}

func NewBabylonHeaderTraversal(babylonClient *http.HTTP, fromHeader *types2.Header, confDepth *big.Int) *BabylonHeaderTraversal {
	return &BabylonHeaderTraversal{
		babylonClient:          babylonClient,
		lastTraversedHeader:    fromHeader,
		blockConfirmationDepth: confDepth,
	}
}

func (f *BabylonHeaderTraversal) LatestHeader() *types2.Header {
	return f.latestHeader
}

func (f *BabylonHeaderTraversal) LastTraversedHeader() *types2.Header {
	return f.lastTraversedHeader
}

func (f *BabylonHeaderTraversal) NextHeaders(maxSize uint64) ([]types2.Header, error) {
	ctx := context.Background()
	status, err := f.babylonClient.Status(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to query latest block: %w", err)
	} else if &status.SyncInfo.LatestBlockHeight == nil {
		return nil, fmt.Errorf("latest header unreported")
	} else {
		block, err := f.babylonClient.Block(ctx, &status.SyncInfo.LatestBlockHeight)
		if err != nil {
			return nil, fmt.Errorf("unable to query latest block: %w", err)
		}
		f.latestHeader = &block.Block.Header
	}
	latestHeaderJson, _ := json.Marshal(status.SyncInfo.LatestBlockHeight)
	log.Info("header traversal db latest header: ", "info", string(latestHeaderJson))

	endHeight := new(big.Int).Sub(big.NewInt(status.SyncInfo.LatestBlockHeight), f.blockConfirmationDepth)
	if endHeight.Sign() < 0 {
		// No blocks with the provided confirmation depth available
		return nil, nil
	}

	lastTraversedHeaderJson, _ := json.Marshal(f.lastTraversedHeader)
	log.Info("babylon header traversal last traversed deader to json: ", "info", string(lastTraversedHeaderJson))

	if f.lastTraversedHeader != nil {
		cmp := big.NewInt(f.lastTraversedHeader.Height).Cmp(endHeight)
		if cmp == 0 {
			return nil, nil
		} else if cmp > 0 {
			return nil, ErrHeaderTraversalAheadOfProvider
		}
	}

	nextHeight := bigint.Zero
	if f.lastTraversedHeader != nil {
		nextHeight = new(big.Int).Add(big.NewInt(f.lastTraversedHeader.Height), bigint.One)
	}

	endHeight = bigint.Clamp(nextHeight, endHeight, maxSize)
	headers, err := f.BlockHeadersByRange(ctx, nextHeight, endHeight)
	if err != nil {
		return nil, fmt.Errorf("error querying blocks by range: %w", err)
	}
	if len(headers) == 0 {
		return nil, nil
	}
	err = f.checkHeaderListByHash(f.lastTraversedHeader, headers)
	if err != nil {
		log.Error("next headers check blockList by hash", "error", err)
		return nil, err
	}

	numHeaders := len(headers)
	if numHeaders == 0 {
		return nil, nil
	} else if f.lastTraversedHeader != nil && headers[0].LastBlockID.Hash.String() != f.lastTraversedHeader.Hash().String() {
		log.Error("Err header traversal and provider mismatched state", "parentHash = ", headers[0].LastBlockID.Hash.String(), "hash", f.lastTraversedHeader.Hash().String())
		return nil, ErrHeaderTraversalAndProviderMismatchedState
	}
	f.lastTraversedHeader = &headers[numHeaders-1]
	return headers, nil
}

func (f *BabylonHeaderTraversal) checkHeaderListByHash(dbLatestHeader *types2.Header, headerList []types2.Header) error {
	if len(headerList) == 0 {
		return nil
	}
	if len(headerList) == 1 {
		return nil
	}
	// check input and db
	// input first ParentHash = dbLatestHeader.Hash
	if dbLatestHeader != nil && headerList[0].LastBlockID.Hash.String() != dbLatestHeader.Hash().String() {
		log.Error("check header list by hash", "parentHash = ", headerList[0].LastBlockID.Hash.String(), "hash", dbLatestHeader.Hash().String())
		return ErrHeaderTraversalCheckHeaderByHashDelDbData
	}

	// check input
	for i := 1; i < len(headerList); i++ {
		if headerList[i].LastBlockID.Hash.String() != headerList[i-1].Hash().String() {
			return fmt.Errorf("check header list by hash: block parent hash not equal parent block hash")
		}
	}
	return nil
}

func (f *BabylonHeaderTraversal) ChangeLastTraversedHeaderByDelAfter(dbLatestHeader *types2.Header) {
	f.lastTraversedHeader = dbLatestHeader
}

func (f *BabylonHeaderTraversal) BlockHeadersByRange(ctx context.Context, nextHeight *big.Int, endHeight *big.Int) ([]types2.Header, error) {
	var headers []types2.Header
	for blockHeight := nextHeight.Int64(); blockHeight <= endHeight.Int64(); blockHeight++ {
		block, err := f.babylonClient.Block(ctx, &blockHeight)
		if err != nil {
			return nil, fmt.Errorf("failed to get block, height = %v , err = %v", blockHeight, err)
		}
		headers = append(headers, block.Block.Header)
	}
	return headers, nil
}
