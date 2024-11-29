package common

type VoteMessageType uint8

const (
	DisagreeVote VoteMessageType = iota
	AgreeVote
	DidNotVote
)
