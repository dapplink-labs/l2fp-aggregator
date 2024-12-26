package node

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/dapplink-labs/l2fp-aggregator/bindings/bls"
	"github.com/dapplink-labs/l2fp-aggregator/bindings/finality"
	"github.com/dapplink-labs/l2fp-aggregator/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	common3 "github.com/dapplink-labs/l2fp-aggregator/common"
	"github.com/dapplink-labs/l2fp-aggregator/config"
	"github.com/dapplink-labs/l2fp-aggregator/manager/types"
	common2 "github.com/dapplink-labs/l2fp-aggregator/node/common"
	"github.com/dapplink-labs/l2fp-aggregator/sign"
	"github.com/dapplink-labs/l2fp-aggregator/store"
	"github.com/dapplink-labs/l2fp-aggregator/synchronizer"
	wsclient "github.com/dapplink-labs/l2fp-aggregator/ws/client"
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
	synchronizer     *synchronizer.BabylonSynchronizer
	txMsgChan        chan store.TxMessage
}

func NewFinalityNode(ctx context.Context, db *store.Storage, privKey *ecdsa.PrivateKey, keyPairs *sign.KeyPair, shouldRegist bool, cfg *config.Config, logger log.Logger, shutdown context.CancelCauseFunc) (*Node, error) {
	from := crypto.PubkeyToAddress(privKey.PublicKey)

	pubkey := crypto.CompressPubkey(&privKey.PublicKey)
	pubkeyHex := hex.EncodeToString(pubkey)
	logger.Info(fmt.Sprintf("pub key is (%s) \n", pubkeyHex))
	if shouldRegist {
		tx, err := registerOperator(ctx, cfg, privKey, pubkeyHex, keyPairs)
		if err != nil {
			logger.Error("failed to register operator", "err", err)
			return nil, err
		}
		logger.Info("success to register operator", "tx_hash", tx.Hash())
	}

	wsClient, err := wsclient.NewWSClient(cfg.Node.WsAddr, "/ws", privKey, pubkeyHex)
	if err != nil {
		return nil, err
	}
	txMsgChan := make(chan store.TxMessage, 100)
	synchronizer, err := synchronizer.NewBabylonSynchronizer(ctx, cfg, db, shutdown, logger, txMsgChan)
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
	height, err := n.db.GetBabylonScannedHeight()
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
				height, err := n.db.GetBabylonScannedHeight()
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
			byteData := crypto.Keccak256Hash(bCFP)
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
			byteData := crypto.Keccak256Hash(bCBD)
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
			byteData := crypto.Keccak256Hash(bCPR)
			bSign = n.keyPairs.SignMessage(byteData)
			n.log.Info("success to sign CommitPubRandListMsg", "signature", bSign.String())
		} else {
			return nil, nil
		}
	}
	return bSign, nil
}

func registerOperator(ctx context.Context, cfg *config.Config, priKey *ecdsa.PrivateKey, node string, keyPairs *sign.KeyPair) (*types2.Transaction, error) {
	ethCli, err := client.DialEthClientWithTimeout(ctx, cfg.EthRpc, false)
	if err != nil {
		return nil, err
	}
	frmContract, err := finality.NewFinalityRelayerManager(common.HexToAddress(cfg.Contracts.FrmContractAddress), ethCli)
	if err != nil {
		return nil, err
	}
	bar, err := bls.NewBLSApkRegistry(common.HexToAddress(cfg.Contracts.BarContactAddress), ethCli)
	if err != nil {
		return nil, err
	}
	fParsed, err := finality.FinalityRelayerManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	rawFrmContract := bind.NewBoundContract(
		common.HexToAddress(cfg.Contracts.FrmContractAddress), *fParsed, ethCli, ethCli,
		ethCli,
	)
	bParsed, err := bls.BLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	rawBarContract := bind.NewBoundContract(
		common.HexToAddress(cfg.Contracts.BarContactAddress), *bParsed, ethCli, ethCli,
		ethCli,
	)

	topts, err := client.NewTransactOpts(ctx, cfg.EthChainID, priKey)
	if err != nil {
		return nil, err
	}

	nodeAddr := crypto.PubkeyToAddress(priKey.PublicKey)
	latestBlock, err := ethCli.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	cOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(latestBlock)),
		From:        nodeAddr,
	}

	msg, err := bar.PubkeyRegistrationMessageHash(cOpts, nodeAddr)
	if err != nil {
		return nil, err
	}

	sigMsg := new(bn254.G1Affine).ScalarMultiplication(sign.NewG1Point(msg.X, msg.Y).G1Affine, keyPairs.PrivKey.BigInt(new(big.Int)))

	res, err := sign.VerifySigHashedToCurve(sigMsg, keyPairs.GetPubKeyG2().G2Affine, sign.NewG1Point(msg.X, msg.Y).G1Affine)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("local verify result", res)

	params := bls.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: bls.BN254G1Point{
			X: sigMsg.X.BigInt(new(big.Int)),
			Y: sigMsg.Y.BigInt(new(big.Int)),
		},
		PubkeyG1: bls.BN254G1Point{
			X: keyPairs.GetPubKeyG1().X.BigInt(new(big.Int)),
			Y: keyPairs.GetPubKeyG1().Y.BigInt(new(big.Int)),
		},
		PubkeyG2: bls.BN254G2Point{
			X: [2]*big.Int{keyPairs.GetPubKeyG2().X.A1.BigInt(new(big.Int)), keyPairs.GetPubKeyG2().X.A0.BigInt(new(big.Int))},
			Y: [2]*big.Int{keyPairs.GetPubKeyG2().Y.A1.BigInt(new(big.Int)), keyPairs.GetPubKeyG2().Y.A0.BigInt(new(big.Int))},
		},
	}

	regBlsTx, err := bar.RegisterBLSPublicKey(topts, nodeAddr, params, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to craft RegisterBLSPublicKey transaction, err: %v", err)
	}
	fRegBlsTx, err := rawBarContract.RawTransact(topts, regBlsTx.Data())
	if err != nil {
		return nil, fmt.Errorf("failed to raw RegisterBLSPublicKey transaction, err: %v", err)
	}
	err = ethCli.SendTransaction(ctx, fRegBlsTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send RegisterBLSPublicKey transaction, err: %v", err)
	}

	_, err = client.GetTransactionReceipt(ctx, ethCli, fRegBlsTx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get RegisterBLSPublicKey transaction receipt, err: %v, tx_hash: %v", err, fRegBlsTx.Hash().String())
	}

	regOTx, err := frmContract.RegisterOperator(topts, node)
	if err != nil {
		return nil, fmt.Errorf("failed to craft RegisterOperator transaction, err: %v", err)
	}
	fRegOTx, err := rawFrmContract.RawTransact(topts, regOTx.Data())
	if err != nil {
		return nil, fmt.Errorf("failed to raw RegisterOperator transaction, err: %v", err)
	}
	err = ethCli.SendTransaction(ctx, fRegOTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send RegisterOperator transaction, err: %v", err)
	}
	_, err = client.GetTransactionReceipt(ctx, ethCli, fRegOTx.Hash())
	if err != nil {
		return nil, fmt.Errorf("failed to get RegisterOperator transaction receipt, err: %v, tx_hash: %v", err, fRegOTx.Hash().String())
	}

	return fRegOTx, nil
}
