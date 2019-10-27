package chaincode_test

import (
	"cartransfer/chaincode"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	aliceJSON  = `{"Id":"1", "Name":"Alice"}`
	ownersJSON = "[" + aliceJSON + "]"
)

func TestInit(t *testing.T) {
	stub := shim.NewMockStub("cartransfer", new(chaincode.CarTransferCC))
	if assert.NotNil(t, stub) {
		res := stub.MockInit(util.GenerateUUID(), nil)
		assert.True(t, res.Status < shim.ERRORTHRESHOLD)
	}
}

func TestInvoke(t *testing.T) {
	stub := shim.NewMockStub("cartransfer", new(chaincode.CarTransferCC))
	if assert.NotNil(t, stub) {
		return
	}

	if !assert.True(t, stub.MockInit(util.GenerateUUID(), nil).Status < shim.ERRORTHRESHOLD) {
		return
	}

	// Invoke() checks
	if !assert.True(
		t,
		stub.MockInvoke(
			util.GenerateUUID(),
			getBytes("AddOwner", aliceJSON)).Status < shim.ERRORTHRESHOLD) {
		return
	}
}


func getBytes(function string, args ...string) [][]byte {
	bytes := make([][]byte, 0, len(args) + 1)
	bytes = append(bytes, []byte(function))
	for _, s := range args {
		bytes = append(bytes, []byte(s))
	}
	return bytes
}