package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dapplink-labs/bbn-relayer/node/common"
	"github.com/dapplink-labs/bbn-relayer/sign"
	"sync"
	"time"

	"github.com/influxdata/influxdb/pkg/slices"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"

	"github.com/dapplink-labs/bbn-relayer/manager/types"
	"github.com/dapplink-labs/bbn-relayer/ws/server"
)

func (m *Manager) sign(ctx types.Context, request interface{}, method types.Method) (types.SignMsgResponse, error) {
	respChan := make(chan server.ResponseMsg)
	stopChan := make(chan struct{})

	if err := m.wsServer.RegisterResChannel(ctx.RequestId(), respChan, stopChan); err != nil {
		m.log.Error("failed to register response channel at signing step", "err", err)
		return types.SignMsgResponse{}, err
	}
	m.log.Info("Registered ResChannel with requestID", "requestID", ctx.RequestId())

	errSendChan := make(chan struct{})
	responseNodes := make(map[string]struct{})
	var validSignResponse types.SignMsgResponse
	var g2Point *sign.G2Point
	var g2Points []*sign.G2Point
	var g1Point *sign.G1Point
	var g1Points []*sign.G1Point

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		cctx, cancel := context.WithTimeout(context.Background(), m.signTimeout)
		defer func() {
			m.log.Info("exit signing process")
			cancel()
			close(stopChan)
			wg.Done()
		}()
		for {
			select {
			case <-errSendChan:
				return
			case resp := <-respChan:
				m.log.Info(fmt.Sprintf("signed response: %s", resp.RpcResponse.String()), "node", resp.SourceNode)
				if !slices.ExistsIgnoreCase(m.NodeMembers, resp.SourceNode) { // ignore the message which the sender should not be involved in approver set
					continue
				}
				func() {
					defer func() {
						responseNodes[resp.SourceNode] = struct{}{}
					}()
					if resp.RpcResponse.Error != nil {
						m.log.Error("Unrecognized error code",
							"err_code", resp.RpcResponse.Error.Code,
							"err_data", resp.RpcResponse.Error.Data,
							"err_message", resp.RpcResponse.Error.Message)
						return
					} else {
						var signResponse types.SignMsgResponse
						if err := tmjson.Unmarshal(resp.RpcResponse.Result, &signResponse); err != nil {
							m.log.Error("failed to unmarshal sign response", "err", err)
							return
						}

						if signResponse.Vote != uint8(common.AgreeVote) {
							return
						}
						dG2Point, err := g2Point.Deserialize(signResponse.G2Point)
						if err != nil {
							m.log.Error("failed to deserialize g2Point", "err", err)
							return
						}

						dSign, err := g1Point.Deserialize(signResponse.Signature)
						if err != nil {
							m.log.Error("failed to deserialize signature", "err", err)
							return
						}
						g2Points = append(g2Points, dG2Point)
						g1Points = append(g1Points, dSign)
						return
					}
				}()

			case <-cctx.Done():
				m.log.Warn("wait for signature timeout")
				return
			default:
				if len(responseNodes) == len(m.NodeMembers) {
					m.log.Info("received all signing responses")
					aSign, _ := aggregateSignaturesAndG2Point(g1Points, g2Points)
					if aSign != nil {
						validSignResponse = types.SignMsgResponse{
							Signature: aSign.Serialize(),
						}
					}
					return
				}
			}
		}
	}()

	m.sendToNodes(ctx, request, method, errSendChan)
	wg.Wait()

	return validSignResponse, nil
}

func (m *Manager) sendToNodes(ctx types.Context, request interface{}, method types.Method, errSendChan chan struct{}) {
	nodes := ctx.AvailableNodes()
	nodeRequest := types.NodeSignRequest{
		Timestamp:   time.Now().UnixMilli(),
		Nodes:       ctx.Approvers(),
		RequestBody: request,
	}
	requestBz, err := json.Marshal(nodeRequest)
	if err != nil {
		m.log.Error("failed to json marshal node request", "err", err)
		errSendChan <- struct{}{}
		return
	}

	rpcRequest := tmtypes.NewRPCRequest(tmtypes.JSONRPCStringID(ctx.RequestId()), method.String(), requestBz)
	for _, node := range nodes {
		go func(node string, request tmtypes.RPCRequest) {
			if err := m.wsServer.SendMsg(
				server.RequestMsg{
					RpcRequest: request,
					TargetNode: node,
				}); err != nil {
				m.log.Error("failed to send sign request to nodes", "err", err)
				errSendChan <- struct{}{}
				return
			}
		}(node, rpcRequest)
	}

}

func aggregateSignaturesAndG2Point(signatures []*sign.G1Point, points []*sign.G2Point) (*sign.G1Point, *sign.G2Point) {
	if len(signatures) == 0 {
		return nil, nil
	}
	var aggSig *sign.G1Point
	var g2Point *sign.G2Point

	for _, sig := range signatures {
		if aggSig == nil {
			aggSig = sig.Clone()
		} else {
			aggSig.Add(sig)
		}
	}

	for _, point := range points {
		if g2Point == nil {
			g2Point = point.Clone()
		} else {
			g2Point.Add(point)
		}
	}

	return aggSig, g2Point
}
