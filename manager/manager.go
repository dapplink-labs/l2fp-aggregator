package manager

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/eniac-x-labs/manta-relayer/config"
	"github.com/eniac-x-labs/manta-relayer/manager/router"
	"github.com/eniac-x-labs/manta-relayer/manager/rpc"
	"github.com/eniac-x-labs/manta-relayer/manager/types"
	"github.com/eniac-x-labs/manta-relayer/store"
	"github.com/eniac-x-labs/manta-relayer/synchronizer"
	"github.com/eniac-x-labs/manta-relayer/ws/server"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/pkg/slices"
)

var (
	errNotEnoughSignNode = errors.New("not enough available nodes to sign state")
	errNotEnoughVoteNode = errors.New("not enough available nodes to vote state")
)

type Manager struct {
	wg          sync.WaitGroup
	done        chan struct{}
	log         log.Logger
	db          *store.Storage
	wsServer    server.IWebsocketManager
	NodeMembers []string
	httpAddr    string
	httpServer  *http.Server

	ctx     context.Context
	stopped atomic.Bool

	l1ChainID  uint64
	privateKey *ecdsa.PrivateKey
	from       common.Address
	l1Client   *ethclient.Client

	signTimeout time.Duration

	synchronizer *synchronizer.Synchronizer
	txMsgChan    chan store.TxMessage
}

func NewFinalityManager(ctx context.Context, db *store.Storage, wsServer server.IWebsocketManager, cfg *config.Config, shutdown context.CancelCauseFunc, logger log.Logger, priv *ecdsa.PrivateKey) (*Manager, error) {
	nodeMemberS := strings.Split(cfg.Manager.NodeMembers, ",")

	service := NewFinalityService(db, logger)
	if cfg.Manager.SdkRpc != "" {
		go rpc.NewAndStartFinalityRpcServer(ctx, cfg.Manager.SdkRpc, service)
	}
	txMsgChan := make(chan store.TxMessage, 100)
	synchronizer, err := synchronizer.NewSynchronizer(ctx, cfg, db, shutdown, logger, txMsgChan)
	if err != nil {
		return nil, err
	}

	return &Manager{
		done:         make(chan struct{}),
		log:          logger,
		db:           db,
		wsServer:     wsServer,
		NodeMembers:  nodeMemberS,
		ctx:          ctx,
		privateKey:   priv,
		from:         crypto.PubkeyToAddress(priv.PublicKey),
		signTimeout:  cfg.Manager.SignTimeout,
		synchronizer: synchronizer,
		txMsgChan:    txMsgChan,
	}, nil
}

func (m *Manager) Start(ctx context.Context) error {
	waitNodeTicker := time.NewTicker(3 * time.Second)
	var done bool
	for !done {
		select {
		case <-waitNodeTicker.C:
			availableNodes := m.availableNodes(m.NodeMembers)
			if len(availableNodes) < len(m.NodeMembers) {
				m.log.Warn("wait node to connect", "availableNodesNum", len(availableNodes), "connectedNodeNum", len(m.NodeMembers))
				continue
			} else {
				done = true
				break
			}
		}
	}

	registry := router.NewRegistry(m)
	r := gin.Default()
	registry.Register(r)

	var s *http.Server
	s = &http.Server{
		Addr:    m.httpAddr,
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error("api server starts failed", err)
		}
	}()
	m.httpServer = s
	go m.synchronizer.Start()

	m.wg.Add(1)
	go m.work()
	m.log.Info("manager is starting......")
	return nil
}

func (m *Manager) Stop(ctx context.Context) error {
	close(m.done)
	if err := m.httpServer.Shutdown(ctx); err != nil {
		m.log.Error("http server forced to shutdown", "err", err)
		return err
	}
	if err := m.synchronizer.Close(); err != nil {
		m.log.Error("synchronizer server forced to shutdown", "err", err)
		return err
	}
	m.stopped.Store(true)
	m.log.Info("Server exiting")
	return nil
}

func (m *Manager) Stopped() bool {
	return m.stopped.Load()
}

func (m *Manager) work() {
	defer m.wg.Done()

	for {
		select {
		case txMsg := <-m.txMsgChan:
			var request types.SignMsgRequest
			var signature []byte

			request.BlockNumber = big.NewInt(int64(txMsg.BlockHeight))
			request.TxType = txMsg.Type
			request.TxHash = txMsg.TransactionHash

			if sig, err := m.db.GetSignature(request.BlockNumber.Int64()); len(sig.Data) > 0 {
				if err != nil {
					m.log.Error("failed to get signature by tx hash", "tx_hash", hexutil.Encode(request.TxHash), "err", err)
					continue
				}
				signature = sig.Data
				m.log.Info("get stored signature ", "tx_hash", hexutil.Encode(request.TxHash), "sig", sig)
			} else {
				signature, err = m.SignMsgBatch(request)
				if errors.Is(err, errNotEnoughSignNode) || errors.Is(err, errNotEnoughVoteNode) {
					continue
				} else if err != nil {
					m.log.Error("failed to sign msg", "err", err)
					continue
				}
				m.log.Info("success to sign msg", "txHash", hexutil.Encode(request.TxHash), "signature", hexutil.Encode(signature), "block_number", request.BlockNumber.Int64())
				if err = m.db.SetSignature(store.Signature{
					BlockNumber:     request.BlockNumber.Int64(),
					TransactionHash: request.TxHash,
					Data:            signature,
					Timestamp:       time.Now().Unix(),
				}); err != nil {
					m.log.Error("failed to store signature", "err", err)
					continue
				}
			}

			//tx, err := m.craftTx(ctx, data, m.msmContractAddr)
			//if err != nil {
			//	m.log.Error("failed to craft transaction options", "err", err)
			//	break
			//}
			//
			//err = m.l1Client.SendTransaction(ctx, tx)
			//if err != nil {
			//	m.log.Error("failed to send verify finality tx", "err", err)
			//	break
			//}
			//
			//receipt, err := getTransactionReceipt(ctx, m.l1Client, tx.Hash())
			//if err != nil {
			//	m.log.Error("failed to get verify finality transaction receipt", "err", err)
			//	break
			//}
			//m.log.Info("success to send verify finality transaction", "tx_hash", receipt.TxHash.String())

		case <-m.done:
			return
		}
	}
}

func (m *Manager) SignMsgBatch(request types.SignMsgRequest) ([]byte, error) {
	m.log.Info("received sign request", "tx_type", request.TxType, "block_number", request.BlockNumber.Uint64(), "tx_hash", hexutil.Encode(request.TxHash))

	availableNodes := m.availableNodes(m.NodeMembers)
	if len(availableNodes) < len(m.NodeMembers) {
		return nil, errNotEnoughSignNode
	}

	ctx := types.NewContext().
		WithAvailableNodes(availableNodes).
		WithRequestId(randomRequestId())

	var resp types.SignMsgResponse
	var signErr error
	resp, signErr = m.sign(ctx, request, types.SignMsgBatch)
	if signErr != nil {
		return nil, signErr
	}
	if resp.Signature == nil {
		return nil, errNotEnoughVoteNode
	}

	return resp.Signature, nil
}

func (m *Manager) availableNodes(nodeMembers []string) []string {
	aliveNodes := m.wsServer.AliveNodes()
	log.Info("check available nodes", "expected", fmt.Sprintf("%v", nodeMembers), "alive nodes", fmt.Sprintf("%v", aliveNodes))

	availableNodes := make([]string, 0)
	for _, n := range aliveNodes {
		if slices.ExistsIgnoreCase(nodeMembers, n) {
			availableNodes = append(availableNodes, n)
		}
	}
	return availableNodes
}

func randomRequestId() string {
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	return time.Now().Format("20060102150405") + code
}
