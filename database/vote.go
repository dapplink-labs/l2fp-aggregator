package database

import (
	"database/sql"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
	"time"
)

type Vote struct {
	GUID       uuid.UUID `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	L2BlockNum *big.Int  `gorm:"serializer:u256;column:l2_block_num" db:"l2_block_num" json:"l2_block_num" form:"l2_block_num"`
	Node       string    `gorm:"column:node" db:"node" json:"node" form:"node"`
	Signature  []byte    `gorm:"column:signature" db:"signature" json:"signature" form:"signature"`
	Result     uint8     `gorm:"column:result" db:"result" json:"result" form:"result"`
	Timestamp  int64     `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (Vote) TableName() string {
	return "vote"
}

type VoteDB interface {
	VoteView
	BuildVote(block *big.Int, node string, signature []byte, res uint8) []Vote
	StoreVote([]Vote) error
}

type VoteView interface {
	LatestVote() (*Node, error)
	GetSignatureByStateRoot(stateRoot common.Hash) (string, error)
}

type voteDB struct {
	gorm *gorm.DB
}

func NewVoteDB(db *gorm.DB) VoteDB {
	return &voteDB{gorm: db}
}

func (v voteDB) BuildVote(block *big.Int, node string, signature []byte, res uint8) []Vote {
	var votes []Vote

	vote := Vote{
		GUID:       uuid.New(),
		L2BlockNum: block,
		Node:       node,
		Signature:  signature,
		Result:     res,
		Timestamp:  time.Now().Unix(),
	}
	votes = append(votes, vote)

	return votes
}

func (v voteDB) StoreVote(votes []Vote) error {
	result := v.gorm.CreateInBatches(&votes, len(votes))
	return result.Error
}

func (v voteDB) LatestVote() (*Node, error) {
	var node Node
	result := v.gorm.Table("node").Order("timestamp desc").Take(&node)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &node, nil
}

func (v voteDB) GetSignatureByStateRoot(stateRoot common.Hash) (string, error) {
	var signature sql.NullString
	err := v.gorm.Table("node").Where("state_root = ?", stateRoot.String()).Select("signature").Row().Scan(&signature)

	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	if signature.Valid {
		return signature.String, nil
	} else {
		return "", nil
	}
}
