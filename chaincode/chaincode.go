package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ChaincodeTest struct {
}

func (c *ChaincodeTest) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("init success"))
}

func (c *ChaincodeTest) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("invoke success"))
}

func CreateOrder(stub shim.ChaincodeStubInterface) {
	stub.PutState("panlin", []byte("aaaa"))
}

func main() {
	err := shim.Start(new(ChaincodeTest))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
