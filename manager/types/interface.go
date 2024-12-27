package types

type SignService interface {
	SignMsgBatch(request SignMsgRequest) (*SignResult, error)
}
