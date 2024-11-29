package database

import (
	"context"
	"fmt"
	"github.com/eniac-x-labs/manta-relayer/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

func TestConnectDatabase(t *testing.T) {
	ctx := context.Background()
	cfg := config.DB{
		DbHost:     "127.0.0.1",
		DbPort:     5432,
		DbName:     "finality_node",
		DbUser:     "postgres",
		DbPassword: "",
	}
	db, err := NewDB(ctx, cfg)
	require.Nil(t, err)
	err = db.ExecuteSQLMigration("../migrations")
	require.Nil(t, err)
}

func TestStateRootStore(t *testing.T) {
	ctx := context.Background()
	cfg := config.DB{
		DbHost:     "127.0.0.1",
		DbPort:     5432,
		DbName:     "finality_node",
		DbUser:     "postgres",
		DbPassword: "",
	}
	db, err := NewDB(ctx, cfg)
	require.Nil(t, err)
	var stateRoots []StateRoot
	fakeStateRoot := StateRoot{
		GUID:        uuid.New(),
		L2BlockNum:  big.NewInt(1000),
		FinalizedL1: big.NewInt(100),
		SafeL1:      big.NewInt(99),
		FinalizedL2: big.NewInt(900),
		SafeL2:      big.NewInt(999),
		StateRoot:   common.HexToHash("0x1234567890"),
		Timestamp:   time.Now().Unix(),
	}
	stateRoots = append(stateRoots, fakeStateRoot)
	err = db.StateRoot.StoreStateRoot(stateRoots)
	require.Nil(t, err)
	sR, err := db.StateRoot.LatestStateRoot()
	require.Nil(t, err)
	t.Log(fmt.Sprintf("get state root successfully, root: %s", sR.StateRoot))
}
