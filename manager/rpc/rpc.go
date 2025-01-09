package rpc

import (
	"context"
	"math/big"
	"net"
	"net/rpc"

	"github.com/ethereum/go-ethereum/log"
)

type FinalityRequest struct {
	BlockNumber *big.Int `json:"l2_block_number"`
}

type StakerDelegationRequest struct {
	Address string `json:"address"`
}

type FinalityRpcServer struct {
	FinalityInterface
}

func NewAndStartFinalityRpcServer(ctx context.Context, address string, finality FinalityInterface) {
	if err := rpc.Register(&FinalityRpcServer{
		FinalityInterface: finality,
	}); err != nil {
		log.Error("RpcServer Register failed", "err", err)
		return
	}
	log.Debug("RpcServer Register finished")

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("RpcServer Listen failed", "err", err, "address", address)
		return
	}
	log.Debug("RpcServer listen address finished", "address", address)

	for {
		select {
		case <-ctx.Done():
			listener.Close()
			log.Info("finality rpc listener closed successfully")
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Error("RpcServer listener.Accept failed", "err", err)
			}

			go rpc.ServeConn(conn)
		}
	}
}

func (s *FinalityRpcServer) Finality(req FinalityRequest, reply *interface{}) error {
	var err error
	*reply, err = s.SignatureByBlock(req.BlockNumber)
	if err != nil {
		return err
	}

	return nil
}

func (s *FinalityRpcServer) Staker(req StakerDelegationRequest, reply *interface{}) error {
	var err error
	*reply, err = s.StakerDelegationByAddress(req.Address)
	if err != nil {
		return err
	}

	return nil
}
