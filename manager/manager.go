package manager

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eniac-x-labs/manta-relayer/bindings"
	"github.com/eniac-x-labs/manta-relayer/client"
	"github.com/eniac-x-labs/manta-relayer/config"
	"github.com/eniac-x-labs/manta-relayer/database"
	"github.com/eniac-x-labs/manta-relayer/manager/router"
	"github.com/eniac-x-labs/manta-relayer/manager/rpc"
	"github.com/eniac-x-labs/manta-relayer/manager/types"
	"github.com/eniac-x-labs/manta-relayer/ws/server"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/pkg/slices"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	errCheckedNextBlock  = errors.New("next checkpoint block is less than latest checked l2 block")
	errNotEnoughSignNode = errors.New("not enough available nodes to sign state")
	errNotEnoughVoteNode = errors.New("not enough available nodes to vote state")
)

type Manager struct {
	wg          sync.WaitGroup
	done        chan struct{}
	log         log.Logger
	db          *database.DB
	wsServer    server.IWebsocketManager
	NodeMembers []string
	httpAddr    string
	httpServer  *http.Server

	ctx     context.Context
	stopped atomic.Bool

	l1ChainID       uint64
	privateKey      *ecdsa.PrivateKey
	from            common.Address
	l2ooContract    *bindings.L2OutputOracle
	msmContract     *bindings.MantaServiceManager
	msmContractAddr common.Address
	msmABI          *abi.ABI
	l1Client        *ethclient.Client
	rollupClient    client.RollupClient

	networkTimeout time.Duration
	pollInterval   time.Duration
	signTimeout    time.Duration
}

func NewFinalityManager(ctx context.Context, db *database.DB, wsServer server.IWebsocketManager, cfg *config.Config, logger log.Logger, priv *ecdsa.PrivateKey) (*Manager, error) {
	nodeMemberS := strings.Split(cfg.Manager.NodeMembers, ",")

	l1Client, err := client.DialEthClientWithTimeout(ctx, cfg.Manager.L1EthRpc, false)
	if err != nil {
		return nil, err
	}
	l2ooContract, err := bindings.NewL2OutputOracle(common.HexToAddress(cfg.L2ooContractAddress), l1Client)
	if err != nil {
		return nil, err
	}
	msmContract, err := bindings.NewMantaServiceManager(common.HexToAddress(cfg.MsmContractAddress), l1Client)
	if err != nil {
		return nil, err
	}
	parsed, err := bindings.MantaServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	rollupClient, err := client.DialEthClient(ctx, cfg.Manager.RollupRpc)
	if err != nil {
		return nil, err
	}

	service := NewFinalityService(db, logger)
	if cfg.Manager.SdkRpc != "" {
		go rpc.NewAndStartFinalityRpcServer(ctx, cfg.Manager.SdkRpc, service)
	}

	return &Manager{
		done:            make(chan struct{}),
		log:             logger,
		db:              db,
		wsServer:        wsServer,
		NodeMembers:     nodeMemberS,
		ctx:             ctx,
		privateKey:      priv,
		from:            crypto.PubkeyToAddress(priv.PublicKey),
		l2ooContract:    l2ooContract,
		msmContract:     msmContract,
		msmContractAddr: common.HexToAddress(cfg.MsmContractAddress),
		msmABI:          parsed,
		l1ChainID:       cfg.L1ChainID,
		l1Client:        l1Client,
		rollupClient:    rollupClient,
		networkTimeout:  cfg.Manager.NetworkTimeout,
		pollInterval:    cfg.Manager.PollInterval,
		signTimeout:     cfg.Manager.SignTimeout,
	}, nil
}

func (m *Manager) Start(ctx context.Context) error {
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

	m.wg.Add(1)
	go m.loop()
	m.log.Info("manager is starting......")
	return nil
}

func (m *Manager) Stop(ctx context.Context) error {
	close(m.done)
	if err := m.httpServer.Shutdown(ctx); err != nil {
		m.log.Error("Server forced to shutdown", "err", err)
		return err
	}
	m.stopped.Store(true)
	m.log.Info("Server exiting")
	return nil
}

func (m *Manager) Stopped() bool {
	return m.stopped.Load()
}

func (m *Manager) loop() {
	defer m.wg.Done()

	ctx := m.ctx

	ticker := time.NewTicker(m.pollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			output, shouldPropose, err := m.FetchNextOutputInfo(ctx)
			if errors.Is(err, errCheckedNextBlock) {
				//continue
			} else if err != nil {
				break
			}
			if !shouldPropose {
				continue
			}

			if err = m.db.StateRoot.StoreStateRoot(m.db.StateRoot.BuildStateRoot(output)); err != nil {
				m.log.Error("failed to store state root", "err", err)
			}
			m.log.Info(fmt.Sprintf("store state root successfully, stateroot:%s, l2block:%s", output.StateRoot, output.BlockRef))

			var request types.SignStateRequest
			var signature []byte

			request.StateRoot = output.StateRoot
			request.L2BlockNumber = big.NewInt(int64(output.BlockRef.Number))
			signature, err = m.SignStateBatch(request)
			if errors.Is(err, errNotEnoughSignNode) || errors.Is(err, errNotEnoughVoteNode) {
				continue
			} else if err != nil {
				m.log.Error("failed to sign state", "err", err)
				break
			}

			m.log.Info("success to sign state", "sign", common.BytesToHash(signature))

			if err = m.db.StateRoot.UpdateSignatureByStateRoot(output.StateRoot, signature); err != nil {
				m.log.Error("failed to store signature", "block", output.BlockRef.Number, "err", err)
				break
			}

			data, err := verifyFinalityTxData(m.msmABI, output)
			if err != nil {
				m.log.Error("failed to pack verify finality tx data", "err", err)
				break
			}

			tx, err := m.craftTx(ctx, data, m.msmContractAddr)
			if err != nil {
				m.log.Error("failed to craft transaction options", "err", err)
				break
			}

			err = m.l1Client.SendTransaction(ctx, tx)
			if err != nil {
				m.log.Error("failed to send verify finality tx", "err", err)
				break
			}

			receipt, err := getTransactionReceipt(ctx, m.l1Client, tx.Hash())
			if err != nil {
				m.log.Error("failed to get verify finality transaction receipt", "err", err)
				break
			}
			m.log.Info("success to send verify finality transaction", "tx_hash", receipt.TxHash.String())

		case <-m.done:
			return
		}
	}
}

func (m *Manager) SignStateBatch(request types.SignStateRequest) ([]byte, error) {
	m.log.Info("received sign state request", "state_root", request.StateRoot, "l2_block_number", request.L2BlockNumber.Uint64())

	if sig, err := m.db.StateRoot.GetSignatureByStateRoot(request.StateRoot); len(sig) > 0 {
		if err != nil {
			m.log.Error("failed to get state root signature by state root", "state_root", request.StateRoot, "err", err)
			return nil, err
		}
		m.log.Info("get stored signature ", "state_root", request.StateRoot.String(), "sig", sig)

		response := types.SignStateResponse{
			Signature: sig,
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Error("sign state response failed to marshal !")
			return nil, err
		}
		return responseBytes, nil
	}
	availableNodes := m.availableNodes(m.NodeMembers)
	if len(availableNodes) < len(m.NodeMembers) {
		return nil, errNotEnoughSignNode
	}

	ctx := types.NewContext().
		WithAvailableNodes(availableNodes).
		WithRequestId(randomRequestId())

	var resp types.SignStateResponse
	var signErr error
	resp, signErr = m.sign(ctx, request, types.SignStateBatch)
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

func verifyFinalityTxData(abi *abi.ABI, output *client.OutputResponse) ([]byte, error) {
	batchHeader := bindings.IMantaServiceManagerBatchHeader{
		FinalityRoot:          output.OutputRoot,
		QuorumNumbers:         nil,
		SignedStakeForQuorums: nil,
		ReferenceBlockNumber:  uint32(output.Status.CurrentL1.Number),
		OutputRoot:            output.OutputRoot,
		L2BlockNumber:         big.NewInt(int64(output.BlockRef.Number)),
		L1BlockHash:           output.Status.CurrentL1.Hash,
		L1BlockNumber:         big.NewInt(int64(output.Status.CurrentL1.Number)),
	}

	return abi.Pack(
		"verifyFinality",
		batchHeader,
	)
}
