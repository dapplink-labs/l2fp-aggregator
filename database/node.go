package database

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Node struct {
	GUID      uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	StateRoot common.Hash `gorm:"column:state_root;serializer:bytes" db:"state_root" json:"state_root" form:"state_root"`
	Signature []byte      `gorm:"column:signature" db:"signature" json:"signature" form:"signature"`
	Vote      uint8       `gorm:"column:vote" db:"vote" json:"vote" form:"vote"`
	Timestamp int64       `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (Node) TableName() string {
	return "node"
}

type NodeDB interface {
	NodeView
	BuildNode(root common.Hash, signature []byte, vote uint8) []Node
	StoreNode([]Node) error
}

type NodeView interface {
	LatestNode() (*Node, error)
	GetNodeSignResByStateRoot(stateRoot common.Hash) (*Node, error)
}

type nodeDB struct {
	gorm *gorm.DB
}

func NewNodeDB(db *gorm.DB) NodeDB {
	return &nodeDB{gorm: db}
}

func (n nodeDB) BuildNode(root common.Hash, signature []byte, vote uint8) []Node {
	var nodes []Node

	node := Node{
		GUID:      uuid.New(),
		StateRoot: root,
		Signature: signature,
		Vote:      vote,
		Timestamp: time.Now().Unix(),
	}
	nodes = append(nodes, node)

	return nodes
}

func (n nodeDB) StoreNode(nodes []Node) error {
	result := n.gorm.CreateInBatches(&nodes, len(nodes))
	return result.Error
}

func (n nodeDB) LatestNode() (*Node, error) {
	var node Node
	result := n.gorm.Table("node").Order("timestamp desc").Take(&node)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &node, nil
}

func (n nodeDB) GetNodeSignResByStateRoot(stateRoot common.Hash) (*Node, error) {
	var nodeRes Node
	err := n.gorm.Table("node").Where("state_root = ?", stateRoot.String()).Take(&nodeRes).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &nodeRes, nil
}
