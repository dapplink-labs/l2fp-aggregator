package finality

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eniac-x-labs/manta-relayer/client"
	"github.com/eniac-x-labs/manta-relayer/config"
	"github.com/eniac-x-labs/manta-relayer/database"
	"github.com/eniac-x-labs/manta-relayer/manager/types"
	common2 "github.com/eniac-x-labs/manta-relayer/node/common"
	"github.com/eniac-x-labs/manta-relayer/sign"
	wsclient "github.com/eniac-x-labs/manta-relayer/ws/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"sync"
	"sync/atomic"
	"time"
)

type Node struct {
	wg         sync.WaitGroup
	done       chan struct{}
	log        log.Logger
	db         *database.DB
	privateKey *ecdsa.PrivateKey
	from       common.Address

	ctx      context.Context
	cancel   context.CancelFunc
	stopChan chan struct{}
	stopped  atomic.Bool

	rollupClient client.RollupClient
	wsClient     *wsclient.WSClients
	keyPairs     *sign.KeyPair

	signTimeout     time.Duration
	outputInterval  time.Duration
	signRequestChan chan tdtypes.RPCRequest
}

func NewFinalityNode(ctx context.Context, db *database.DB, privKey *ecdsa.PrivateKey, keyPairs *sign.KeyPair, cfg *config.Config, logger log.Logger) (*Node, error) {

	from := crypto.PubkeyToAddress(privKey.PublicKey)

	pubkey := crypto.CompressPubkey(&privKey.PublicKey)
	pubkeyHex := hex.EncodeToString(pubkey)
	logger.Info(fmt.Sprintf("pub key is (%s) \n", pubkeyHex))

	wsClient, err := wsclient.NewWSClient(cfg.Node.WsAddr, "/ws", privKey, pubkeyHex)
	if err != nil {
		return nil, err
	}

	rollupClient, err := client.DialEthClient(ctx, cfg.Node.RollupRpc)
	if err != nil {
		return nil, err
	}

	return &Node{
		wg:              sync.WaitGroup{},
		done:            make(chan struct{}),
		stopChan:        make(chan struct{}),
		log:             logger,
		db:              db,
		privateKey:      privKey,
		from:            from,
		ctx:             ctx,
		rollupClient:    rollupClient,
		wsClient:        wsClient,
		keyPairs:        keyPairs,
		signRequestChan: make(chan tdtypes.RPCRequest, 100),
		signTimeout:     cfg.Node.SignTimeout,
		outputInterval:  cfg.Node.OutputInterval,
	}, nil
}

func (n *Node) Start(ctx context.Context) error {
	n.wg.Add(2)
	go n.ProcessMessage()
	go n.sign()
	return nil
}

func (n *Node) Stop(ctx context.Context) error {
	n.cancel()
	close(n.done)
	n.wg.Wait()
	n.stopped.Store(true)
	return nil
}

func (n *Node) Stopped() bool {
	return n.stopped.Load()
}

func (n *Node) sign() {
	defer n.wg.Done()

	n.log.Info("start to sign message")

	go func() {
		defer func() {
			n.log.Info("exit sign process")
		}()
		for {
			select {
			case <-n.stopChan:
				return
			case req := <-n.signRequestChan:
				var resId = req.ID.(tdtypes.JSONRPCStringID).String()
				n.log.Info(fmt.Sprintf("dealing resId (%s) ", resId))

				var nodeSignRequest types.NodeSignRequest
				rawMsg := json.RawMessage{}
				nodeSignRequest.RequestBody = &rawMsg

				if err := json.Unmarshal(req.Params, &nodeSignRequest); err != nil {
					n.log.Error("failed to unmarshal ask request")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					if err := n.wsClient.SendMsg(RpcResponse); err != nil {
						n.log.Error("failed to send msg to manager", "err", err)
					}
					continue
				}
				var requestBody types.SignStateRequest
				if err := json.Unmarshal(rawMsg, &requestBody); err != nil {
					n.log.Error("failed to unmarshal asker's params request body")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					if err := n.wsClient.SendMsg(RpcResponse); err != nil {
						n.log.Error("failed to send msg to manager", "err", err)
					}
					continue
				}
				if len(requestBody.StateRoot.String()) == 0 || requestBody.L2BlockNumber.Uint64() <= 0 {
					n.log.Error("state root and l2 block number must not be nil or negative")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", "state root and l2 block number must not be nil or negative")
					if err := n.wsClient.SendMsg(RpcResponse); err != nil {
						n.log.Error("failed to send msg to manager", "err", err)
					}
					continue
				}
				nodeSignRequest.RequestBody = requestBody

				go n.handleSign(req.ID.(tdtypes.JSONRPCStringID), nodeSignRequest)

			}
		}
	}()
}

func (n *Node) handleSign(resId tdtypes.JSONRPCStringID, req types.NodeSignRequest) error {
	requestBody := req.RequestBody.(types.SignStateRequest)

	var err error
	var vote = uint8(common2.DisagreeVote)
	var output *client.OutputResponse

	res, err := n.db.Node.GetNodeSignResByStateRoot(requestBody.StateRoot)
	if err != nil {
		n.log.Error("node failed to get signature from database", "err", err)
		return err
	}
	if res != nil {
		n.log.Info("the node has signed the state root")
		signResponse := types.SignStateResponse{
			L2BlockNumber: requestBody.L2BlockNumber,
			G2Point:       n.keyPairs.GetPubKeyG2().Serialize(),
			Signature:     res.Signature,
			Vote:          res.Vote,
		}
		RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
		n.log.Info("start to send response to finality manager")

		err = n.wsClient.SendMsg(RpcResponse)
		if err != nil {
			n.log.Error("failed to sendMsg to finality manager", "err", err)
			return err
		} else {
			n.log.Info("send sign response to finality manager successfully ")
			return nil
		}
	}

	output, err = n.rollupClient.OutputAtBlock(n.ctx, requestBody.L2BlockNumber.Uint64())
	if err != nil {
		n.log.Error(fmt.Sprintf("failed to fetch output at block %d: %w", requestBody.L2BlockNumber.Uint64(), err))
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), n.signTimeout)
	ticker := time.NewTicker(n.outputInterval)
	defer cancel()
	defer ticker.Stop()
	for output == nil {
		select {
		case <-ticker.C:
			output, err = n.rollupClient.OutputAtBlock(n.ctx, requestBody.L2BlockNumber.Uint64())
			if err != nil {
				n.log.Error(fmt.Sprintf("failed to fetch output at block %d: %w", requestBody.L2BlockNumber.Uint64(), err))
				return err
			}
		case <-ctx.Done():
			n.log.Warn("sign state root timeout !")
			signResponse := types.SignStateResponse{
				Signature:     nil,
				G2Point:       nil,
				L2BlockNumber: requestBody.L2BlockNumber,
				Vote:          uint8(common2.DidNotVote),
			}
			RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
			n.log.Info("start to send response to finality manager")

			err = n.wsClient.SendMsg(RpcResponse)
			if err != nil {
				n.log.Error("failed to sendMsg to finality manager", "err", err)
				return err
			} else {
				n.log.Info("send sign response to finality manager successfully ")
				return nil
			}
		}
	}

	if output.BlockRef.Number != requestBody.L2BlockNumber.Uint64() { // sanity check, e.g. in case of bad RPC caching
		n.log.Error(fmt.Sprintf("invalid blockNumber: next blockNumber is %v, blockNumber of block is %v", requestBody.L2BlockNumber.Uint64(), output.BlockRef.Number))
		return errors.New("invalid blockNumber")
	}
	if output.StateRoot == requestBody.StateRoot {
		vote = uint8(common2.AgreeVote)
	}

	bSign := n.keyPairs.SignMessage(output.StateRoot)
	n.log.Info("success to sign state root", "signature", bSign.String())

	if err = n.db.Node.StoreNode(n.db.Node.BuildNode(requestBody.StateRoot, bSign.Serialize(), vote)); err != nil {
		n.log.Error("node failed to store signature", "err", err)
	}
	if err != nil {
		n.log.Error("sign failed ", "resId", resId, "err", err)
		var errorRes tdtypes.RPCResponse
		errorRes = tdtypes.NewRPCErrorResponse(resId, 201, "sign failed", err.Error())

		err = n.wsClient.SendMsg(errorRes)
		if err != nil {
			n.log.Error("failed to send msg to finality manager", "err", err)
			return err
		}
	}

	signResponse := types.SignStateResponse{
		Signature:     bSign.Serialize(),
		G2Point:       n.keyPairs.GetPubKeyG2().Serialize(),
		L2BlockNumber: requestBody.L2BlockNumber,
		Vote:          vote,
	}
	RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
	n.log.Info("start to send response to finality manager")

	err = n.wsClient.SendMsg(RpcResponse)
	if err != nil {
		n.log.Error("failed to sendMsg to finality manager", "err", err)
		return err
	} else {
		n.log.Info("send sign response to finality manager successfully ")
		return nil
	}
}
