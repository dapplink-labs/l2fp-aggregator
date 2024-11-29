package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	common2 "github.com/eniac-x-labs/manta-relayer/common"
	"github.com/eniac-x-labs/manta-relayer/common/retry"
)

const (
	defaultDialTimeout  = 5 * time.Second
	defaultDialAttempts = 5
)

type RollupClient interface {
	OutputAtBlock(ctx context.Context, blockNum uint64) (*OutputResponse, error)
	SyncStatus(ctx context.Context) (*common2.Status, error)
	Version(ctx context.Context) (string, error)
	Close()
}

type clnt struct {
	rpc RPC
}

func DialEthClient(ctx context.Context, rpcUrl string) (RollupClient, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()

	bOff := retry.Exponential()
	rpcClient, err := retry.Do(ctx, defaultDialAttempts, bOff, func() (*rpc.Client, error) {
		if !IsURLAvailable(rpcUrl) {
			return nil, fmt.Errorf("address unavailable (%s)", rpcUrl)
		}

		client, err := rpc.DialContext(ctx, rpcUrl)
		if err != nil {
			return nil, fmt.Errorf("failed to dial address (%s): %w", rpcUrl, err)
		}

		return client, nil
	})

	if err != nil {
		return nil, err
	}

	return &clnt{rpc: NewRPC(rpcClient)}, nil
}

type Bytes32 [32]byte

func (b *Bytes32) UnmarshalJSON(text []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(b), text, b[:])
}

func (b *Bytes32) UnmarshalText(text []byte) error {
	return hexutil.UnmarshalFixedText("Bytes32", text, b[:])
}

func (b Bytes32) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

func (b Bytes32) String() string {
	return hexutil.Encode(b[:])
}

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (b Bytes32) TerminalString() string {
	return fmt.Sprintf("%x..%x", b[:3], b[29:])
}

type OutputResponse struct {
	Version               Bytes32            `json:"version"`
	OutputRoot            Bytes32            `json:"outputRoot"`
	BlockRef              common2.L2BlockRef `json:"blockRef"`
	WithdrawalStorageRoot common.Hash        `json:"withdrawalStorageRoot"`
	StateRoot             common.Hash        `json:"stateRoot"`
	Status                *common2.Status    `json:"syncStatus"`
}

func (c *clnt) OutputAtBlock(ctx context.Context, blockNum uint64) (*OutputResponse, error) {
	var output *OutputResponse
	err := c.rpc.CallContext(ctx, &output, "optimism_outputAtBlock", hexutil.Uint64(blockNum))
	return output, err
}

func (c *clnt) SyncStatus(ctx context.Context) (*common2.Status, error) {
	var output *common2.Status
	err := c.rpc.CallContext(ctx, &output, "optimism_syncStatus")
	return output, err
}

func (c *clnt) Version(ctx context.Context) (string, error) {
	var output string
	err := c.rpc.CallContext(ctx, &output, "optimism_version")
	return output, err
}

// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
//
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.
func (c *clnt) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var hex hexutil.Bytes
	err := c.rpc.CallContext(ctx, &hex, "eth_call", toCallArg(msg), toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

func (c *clnt) Close() {
	c.rpc.Close()
}

type RPC interface {
	Close()
	CallContext(ctx context.Context, result any, method string, args ...any) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
}

type rpcClient struct {
	rpc *rpc.Client
}

func NewRPC(client *rpc.Client) RPC {
	return &rpcClient{client}
}

func (c *rpcClient) Close() {
	c.rpc.Close()
}

func (c *rpcClient) CallContext(ctx context.Context, result any, method string, args ...any) error {
	err := c.rpc.CallContext(ctx, result, method, args...)
	return err
}

func (c *rpcClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	err := c.rpc.BatchCallContext(ctx, b)
	return err
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	return rpc.BlockNumber(number.Int64()).String()
}

func IsURLAvailable(address string) bool {
	u, err := url.Parse(address)
	if err != nil {
		return false
	}
	addr := u.Host
	if u.Port() == "" {
		switch u.Scheme {
		case "http", "ws":
			addr += ":80"
		case "https", "wss":
			addr += ":443"
		default:
			return true
		}
	}
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["input"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}

func DialEthClientWithTimeout(ctx context.Context, url string, disableHTTP2 bool) (
	*ethclient.Client, error) {
	ctxt, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()
	if strings.HasPrefix(url, "http") {
		httpClient := new(http.Client)
		if disableHTTP2 {
			log.Debug("Disabled HTTP/2 support in  eth client")
			httpClient.Transport = &http.Transport{
				TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
			}
		}
		rpcClient, err := rpc.DialHTTPWithClient(url, httpClient)
		if err != nil {
			return nil, err
		}
		return ethclient.NewClient(rpcClient), nil
	}
	return ethclient.DialContext(ctxt, url)
}
