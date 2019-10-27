package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"../chaincode"
	".."
)

func main() {
	var _ cartransfer.CarTransfer = (*chaincode.CarTransferCC)(nil)

	err := shim.Start(new(chaincode.CarTransferCC))
	if err != nil {
		fmt.Printf("Error in chaincode process: %s", err)
	}
}
