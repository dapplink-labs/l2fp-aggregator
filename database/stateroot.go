package database

import (
	"errors"
	"github.com/eniac-x-labs/manta-relayer/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
	"time"
)

type StateRoot struct {
	GUID          uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	L2BlockNum    *big.Int    `gorm:"serializer:u256;column:l2_block_num" db:"l2_block_num" json:"l2_block_num" form:"l2_block_num"`
	CurrentL1     *big.Int    `gorm:"serializer:u256;column:current_l1" db:"current_l1" json:"current_l1" form:"current_l1"`
	CurrentL1Hash common.Hash `gorm:"column:current_l1_hash;serializer:bytes" db:"current_l1_hash" json:"current_l1_hash" form:"current_l1_hash"`
	FinalizedL1   *big.Int    `gorm:"serializer:u256;column:finalized_l1" db:"finalized_l1" json:"finalized_l1" form:"finalized_l1"`
	SafeL1        *big.Int    `gorm:"serializer:u256;column:safe_l1" db:"safe_l1" json:"safe_l1" form:"safe_l1"`
	FinalizedL2   *big.Int    `gorm:"serializer:u256;column:finalized_l2" db:"finalized_l2" json:"finalized_l2" form:"finalized_l2"`
	SafeL2        *big.Int    `gorm:"serializer:u256;column:safe_l2" db:"safe_l2" json:"safe_l2" form:"safe_l2"`
	StateRoot     common.Hash `gorm:"column:state_root;serializer:bytes" db:"state_root" json:"state_root" form:"state_root"`
	Signature     []byte      `gorm:"column:signature" db:"signature" json:"signature" form:"signature"`
	IsFinalized   uint8       `gorm:"column:is_finalized" db:"is_finalized" json:"is_finalized" form:"is_finalized"`
	Timestamp     int64       `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (StateRoot) TableName() string {
	return "state_root"
}

type StateRootDB interface {
	StateRootView
	BuildStateRoot(output *client.OutputResponse) []StateRoot
	StoreStateRoot([]StateRoot) error
	UpdateSignatureByStateRoot(stateRoot common.Hash, signature []byte) error
}

type StateRootView interface {
	StateRootByL2Block(L2block *big.Int) (*StateRoot, error)
	LatestStateRoot() (*StateRoot, error)
	GetSignatureByStateRoot(stateRoot common.Hash) ([]byte, error)
}

type stateRootDB struct {
	gorm *gorm.DB
}

func NewStateRootDB(db *gorm.DB) StateRootDB {
	return &stateRootDB{gorm: db}
}

func (s stateRootDB) BuildStateRoot(output *client.OutputResponse) []StateRoot {
	var stateroots []StateRoot

	stateroot := StateRoot{
		GUID:          uuid.New(),
		L2BlockNum:    big.NewInt(int64(output.BlockRef.Number)),
		CurrentL1:     big.NewInt(int64(output.Status.CurrentL1.Number)),
		CurrentL1Hash: output.Status.CurrentL1.Hash,
		FinalizedL1:   big.NewInt(int64(output.Status.FinalizedL1.Number)),
		SafeL1:        big.NewInt(int64(output.Status.SafeL1.Number)),
		FinalizedL2:   big.NewInt(int64(output.Status.FinalizedL2.Number)),
		SafeL2:        big.NewInt(int64(output.Status.SafeL2.Number)),
		StateRoot:     output.StateRoot,
		Signature:     nil,
		IsFinalized:   0,
		Timestamp:     time.Now().Unix(),
	}
	stateroots = append(stateroots, stateroot)

	return stateroots
}

func (s stateRootDB) StoreStateRoot(stateRoots []StateRoot) error {
	result := s.gorm.CreateInBatches(&stateRoots, len(stateRoots))
	return result.Error
}

func (s stateRootDB) StateRootByL2Block(L2block *big.Int) (*StateRoot, error) {
	var stateRoot *StateRoot
	result := s.gorm.Table("state_root").Where("l2_block_num = ?", L2block.Uint64()).Take(&stateRoot)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return stateRoot, nil
}

func (s stateRootDB) LatestStateRoot() (*StateRoot, error) {
	var stateRoot StateRoot
	result := s.gorm.Table("state_root").Order("timestamp desc").Take(&stateRoot)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &stateRoot, nil
}

func (s stateRootDB) GetSignatureByStateRoot(stateRoot common.Hash) ([]byte, error) {
	var signature []byte
	err := s.gorm.Table("state_root").Where("state_root = ?", stateRoot.String()).Select("signature").Row().Scan(&signature)
	if err != nil {
		return nil, err
	}

	if signature != nil {
		return signature, nil
	} else {
		return nil, nil
	}
}

func (s stateRootDB) UpdateSignatureByStateRoot(stateRoot common.Hash, signature []byte) error {
	result := s.gorm.Table("state_root").Where("state_root = ?", stateRoot.String()).Updates(map[string]interface{}{"signature": signature})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
