package types

import "github.com/dapplink-labs/l2fp-aggregator/sign"

type SignService interface {
	SignMsgBatch(request SignMsgRequest) (*sign.G1Point, *sign.G2Point, error)
}
