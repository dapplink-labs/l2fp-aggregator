package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRollupClient(t *testing.T) {
	ctx := context.Background()
	client, err := DialEthClient(ctx, "http://127.0.0.1:7545")
	require.Nil(t, err)
	status, err := client.SyncStatus(ctx)
	require.Nil(t, err)
	output, err := client.OutputAtBlock(ctx, status.SafeL2.Number)
	require.Nil(t, err)
	t.Log(fmt.Sprintf("get state root successfully, root: %s", output.StateRoot))
}
