package types

type SignService interface {
	SignMsgBatch(request SignMsgRequest) ([]byte, error)
}
