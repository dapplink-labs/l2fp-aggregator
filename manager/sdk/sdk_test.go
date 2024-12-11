package sdk

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBabyMsgByBlock(t *testing.T) {
	ast := assert.New(t)
	sdk, err := NewFinalitySDK("localhost:9000")
	ast.NoError(err)
	ast.NotNil(sdk)
	block := big.NewInt(31)
	res, err := sdk.SignatureByBlock(block)
	ast.NoError(err)
	ast.NotNil(res)
	result := SdkResponse{}
	err = json.Unmarshal((*res.(*interface{})).([]byte), &result)
	ast.NoError(err)
	t.Log("signature:", result.Signature)
}
