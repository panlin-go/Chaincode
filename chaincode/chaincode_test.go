package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

const (
	uuid = "chaincode-1"
)

func TestCreateOrder(t *testing.T) {
	fmt.Println("Entering TestCreateLoanApplication")

	stub := shim.NewMockStub("mockStub", new(ChaincodeTest))
	if stub == nil {
		t.Fatalf("MockStub creation failed")
	}

	fmt.Printf("%#v\n", stub)

	//simolation chaincodeinterface Init
	Args := [][]byte{
		{'p', 'a', 'n'},
		{'l', 'i', 'n'}}

	respone := stub.MockInit(uuid, Args)
	fmt.Printf("init resp: %#v\n", respone.String())

	//invoke
	respone = stub.MockInvoke(uuid, Args)
	fmt.Printf("invoke resp: %#v\n", respone.String())

	//get args
	args := stub.GetArgs()
	for i := 0; i < len(args); i++ {
		fmt.Printf("args %d : %s\n", i, string(args[i]))
	}

	//composite key = str1+str2
	attr := make([]string, 10)

	attr[0] = "panlin"
	attr[1] = "lijie"

	str, err := stub.CreateCompositeKey("xxx", attr)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("string: %s\n", str)

}
