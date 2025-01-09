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
	block := big.NewInt(233803)
	res, err := sdk.SignatureByBlock(block)
	ast.NoError(err)
	ast.NotNil(res)
	result := SignResponse{}
	err = json.Unmarshal((*res.(*interface{})).([]byte), &result)
	ast.NoError(err)
	t.Log("signature:", result.Signature)
}

func TestGetStakerDelegation(t *testing.T) {
	ast := assert.New(t)
	sdk, err := NewFinalitySDK("localhost:9000")
	ast.NoError(err)
	ast.NotNil(sdk)
	res, err := sdk.StakerDelegationByAddress("bbn1tkescl5t5j4n3386ztek9smadtr2y3fka5u44m")
	ast.NoError(err)
	ast.NotNil(res)
	result := StakerDelegationResponse{}
	err = json.Unmarshal((*res.(*interface{})).([]byte), &result)
	ast.NoError(err)
	t.Log("amount:", result.Amount)
}
