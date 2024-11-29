package sdk

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestGetFinalityStateByBlock(t *testing.T) {
	ast := assert.New(t)
	sdk, err := NewFinalitySDK("localhost:9000")
	ast.NoError(err)
	ast.NotNil(sdk)
	block := big.NewInt(1000)
	res, err := sdk.StateByBlock(block)
	ast.NoError(err)
	ast.NotNil(res)
	result := SdkResponse{}
	err = json.Unmarshal((*res.(*interface{})).([]byte), &result)
	ast.NoError(err)
	t.Log("state root:", result.StateRoot)
	t.Log("is finalize:", result.IsFinalized)
	t.Log("err", result.Message)
}
