package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

const (
	Uuid = "chaincode-1"
)

func Show(stub *shim.MockStub) {
	//	fmt.Printf("channelId: %s\n", stub.GetChannelID())
	fmt.Printf("TxId: %s\n", stub.GetTxID())
}

func TestCreateOrder(t *testing.T) {
	stub := shim.NewMockStub("mockStub", new(ChaincodeTest))
	if stub == nil {
		t.Fatalf("MockStub creation failed")
	}

	fmt.Printf("%#v\n", stub)

	//simolation chaincodeinterface Init
	respone := stub.MockInit(Uuid, nil)
	fmt.Printf("init resp: %#v\n", respone.String())

	//invoke
	Args := make([][]byte, 3)
	Args[0] = []byte("create")
	Args[1] = []byte("234")
	Args[2] = []byte("10")

	respone = stub.MockInvoke(Uuid, Args)
	fmt.Printf("invoke resp: %#v\n", respone.String())

	//
	data, err := stub.GetState(string(respone.GetPayload()))
	if err != nil {
	}

	fmt.Printf("%#v\n", string(data))

	//get args
	args := stub.GetArgs()
	for i := 0; i < len(args); i++ {
		fmt.Printf("args %d : %s\n", i, string(args[i]))
	}

	//show
	Show(stub)
}
