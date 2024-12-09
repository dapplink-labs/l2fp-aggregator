package node

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	common3 "github.com/eniac-x-labs/manta-relayer/common"
	"github.com/eniac-x-labs/manta-relayer/config"
	"github.com/eniac-x-labs/manta-relayer/manager/types"
	common2 "github.com/eniac-x-labs/manta-relayer/node/common"
	"github.com/eniac-x-labs/manta-relayer/sign"
	"github.com/eniac-x-labs/manta-relayer/store"
	"github.com/eniac-x-labs/manta-relayer/synchronizer"
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
	db         *store.Storage
	privateKey *ecdsa.PrivateKey
	from       common.Address

	ctx      context.Context
	cancel   context.CancelFunc
	stopChan chan struct{}
	stopped  atomic.Bool

	wsClient *wsclient.WSClients
	keyPairs *sign.KeyPair

	signTimeout      time.Duration
	waitScanInterval time.Duration
	signRequestChan  chan tdtypes.RPCRequest
	synchronizer     *synchronizer.Synchronizer
	txMsgChan        chan store.TxMessage
}

func NewFinalityNode(ctx context.Context, db *store.Storage, privKey *ecdsa.PrivateKey, keyPairs *sign.KeyPair, cfg *config.Config, logger log.Logger, shutdown context.CancelCauseFunc) (*Node, error) {

	from := crypto.PubkeyToAddress(privKey.PublicKey)

	pubkey := crypto.CompressPubkey(&privKey.PublicKey)
	pubkeyHex := hex.EncodeToString(pubkey)
	logger.Info(fmt.Sprintf("pub key is (%s) \n", pubkeyHex))

	wsClient, err := wsclient.NewWSClient(cfg.Node.WsAddr, "/ws", privKey, pubkeyHex)
	if err != nil {
		return nil, err
	}
	txMsgChan := make(chan store.TxMessage, 100)
	synchronizer, err := synchronizer.NewSynchronizer(ctx, cfg, db, shutdown, logger, txMsgChan)
	if err != nil {
		return nil, err
	}

	return &Node{
		wg:               sync.WaitGroup{},
		done:             make(chan struct{}),
		stopChan:         make(chan struct{}),
		log:              logger,
		db:               db,
		privateKey:       privKey,
		from:             from,
		ctx:              ctx,
		wsClient:         wsClient,
		keyPairs:         keyPairs,
		signRequestChan:  make(chan tdtypes.RPCRequest, 100),
		signTimeout:      cfg.Node.SignTimeout,
		waitScanInterval: cfg.Node.WaitScanInterval,
		synchronizer:     synchronizer,
		txMsgChan:        txMsgChan,
	}, nil
}

func (n *Node) Start(ctx context.Context) error {
	n.wg.Add(3)
	go n.ProcessMessage()
	go n.sign()
	go n.synchronizer.Start()
	go n.work()
	return nil
}

func (n *Node) Stop(ctx context.Context) error {
	n.cancel()
	close(n.done)
	n.wg.Wait()
	n.synchronizer.Close()
	n.stopped.Store(true)
	return nil
}

func (n *Node) Stopped() bool {
	return n.stopped.Load()
}

func (n *Node) work() {
	defer n.wg.Done()
	for {
		select {
		case txMsg := <-n.txMsgChan:
			if err := n.synchronizer.ProcessNewFinalityProvider(txMsg); err != nil {
				n.log.Error("failed to process NewFinalityProvider msg", "err", err)
				continue
			}
			if err := n.synchronizer.ProcessCreateBTCDelegation(txMsg); err != nil {
				n.log.Error("failed to process CreateBTCDelegation msg", "err", err)
				continue
			}
			if err := n.synchronizer.ProcessCommitPubRandList(txMsg); err != nil {
				n.log.Error("failed to process CommitPubRandList msg", "err", err)
				continue
			}
		}
	}
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
				var requestBody types.SignMsgRequest
				if err := json.Unmarshal(rawMsg, &requestBody); err != nil {
					n.log.Error("failed to unmarshal asker's params request body")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", err.Error())
					if err := n.wsClient.SendMsg(RpcResponse); err != nil {
						n.log.Error("failed to send msg to manager", "err", err)
					}
					continue
				}
				if len(requestBody.TxHash) == 0 || requestBody.BlockNumber.Uint64() <= 0 {
					n.log.Error("tx hash and l2 block number must not be nil or negative")
					RpcResponse := tdtypes.NewRPCErrorResponse(req.ID, 201, "failed", "tx hash and l2 block number must not be nil or negative")
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
	var err error
	var bSign *sign.Signature

	requestBody := req.RequestBody.(types.SignMsgRequest)
	height, err := n.db.GetScannedHeight()
	if err != nil {
		n.log.Error("node failed to get scanned height", "err", err)
		return err
	}
	if requestBody.BlockNumber.Uint64() <= height {
		bSign, err = n.SignMessage(requestBody)
		if err != nil {
			n.log.Error("node failed to sign messages", "err", err)
			return err
		}
		if bSign != nil {
			signResponse := types.SignMsgResponse{
				L2BlockNumber: requestBody.BlockNumber,
				G2Point:       n.keyPairs.GetPubKeyG2().Serialize(),
				Signature:     bSign.Serialize(),
				Vote:          uint8(common2.AgreeVote),
			}
			RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
			n.log.Info("node agree the msg, start to send response to finality manager")

			err = n.wsClient.SendMsg(RpcResponse)
			if err != nil {
				n.log.Error("failed to sendMsg to finality manager", "err", err)
				return err
			} else {
				n.log.Info("send sign response to finality manager successfully ")
				return nil
			}
		} else {
			signResponse := types.SignMsgResponse{
				L2BlockNumber: requestBody.BlockNumber,
				G2Point:       nil,
				Signature:     nil,
				Vote:          uint8(common2.DisagreeVote),
			}
			RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
			fmt.Println()
			n.log.Info("node disagree the msg, start to send response to finality manager")

			err = n.wsClient.SendMsg(RpcResponse)
			if err != nil {
				n.log.Error("failed to sendMsg to finality manager", "err", err)
				return err
			} else {
				n.log.Info("send sign response to finality manager successfully ")
				return nil
			}
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), n.signTimeout)
		ticker := time.NewTicker(n.waitScanInterval)
		defer cancel()
		defer ticker.Stop()
		for bSign != nil {
			select {
			case <-ticker.C:
				height, err := n.db.GetScannedHeight()
				if err != nil {
					n.log.Error("node failed to get scanned height", "err", err)
					return err
				}
				if requestBody.BlockNumber.Uint64() > height {
					n.log.Warn(fmt.Sprintf("node received the task from the manager, the height is %v, but the synchronized height is %v", requestBody.BlockNumber.Uint64(), height))
					return nil
				} else {
					bSign, err = n.SignMessage(requestBody)
					if bSign != nil {
						signResponse := types.SignMsgResponse{
							L2BlockNumber: requestBody.BlockNumber,
							G2Point:       n.keyPairs.GetPubKeyG2().Serialize(),
							Signature:     bSign.Serialize(),
							Vote:          uint8(common2.AgreeVote),
						}
						RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
						n.log.Info("node agree the msg, start to send response to finality manager")

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
			case <-ctx.Done():
				n.log.Warn("sign messages timeout !")
				signResponse := types.SignMsgResponse{
					Signature:     nil,
					G2Point:       nil,
					L2BlockNumber: requestBody.BlockNumber,
					Vote:          uint8(common2.DidNotVote),
				}
				RpcResponse := tdtypes.NewRPCSuccessResponse(resId, signResponse)
				n.log.Info("node did not vote msg, start to send response to finality manager")

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
	}
	return nil
}

func (n *Node) SignMessage(requestBody types.SignMsgRequest) (*sign.Signature, error) {
	var bSign *sign.Signature
	switch requestBody.TxType {
	case common3.MsgCreateFinalityProvider:
		exist, cFP := n.db.GetCreateFinalityProviderMsg(requestBody.TxHash)
		if exist {
			bCFP, err := cFP.FP.Marshal()
			if err != nil {
				n.log.Error("failed to marshal FinalityProviderMsg", "err", err)
				return nil, err
			}
			byteData := sha256.Sum256(bCFP)
			bSign = n.keyPairs.SignMessage(byteData)
			n.log.Info("success to sign FinalityProviderMsg", "signature", bSign.String())
		} else {
			return nil, nil
		}
	case common3.MsgCreateBTCDelegation:
		exist, cBD := n.db.GetCreateBTCDelegationrMsg(requestBody.TxHash)
		if exist {
			bCBD, err := cBD.CBD.Marshal()
			if err != nil {
				n.log.Info("failed to marshal CreateBTCDelegationMsg", "err", err)
				return nil, err
			}
			byteData := sha256.Sum256(bCBD)
			bSign = n.keyPairs.SignMessage(byteData)
			n.log.Info("success to sign CreateBTCDelegationMsg", "signature", bSign.String())
		} else {
			return nil, nil
		}
	case common3.MsgCommitPubRandList:
		exist, cPR := n.db.GetCommitPubRandListMsg(requestBody.TxHash)
		if exist {
			bCPR, err := cPR.CPR.Marshal()
			if err != nil {
				n.log.Info("failed to marshal CommitPubRandListMsg", "err", err)
				return nil, err
			}
			byteData := sha256.Sum256(bCPR)
			bSign = n.keyPairs.SignMessage(byteData)
			n.log.Info("success to sign CommitPubRandListMsg", "signature", bSign.String())
		} else {
			return nil, nil
		}
	}
	return bSign, nil
}
