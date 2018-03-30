package main

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Order struct {
	Uuid   string `json:"Uuid"`
	Price  int    `json:"Price"`
	Num    int    `json:"Num"`
	SumPay int    `json:"SumPay"`
}

type ChaincodeTest struct {
}

func (c *ChaincodeTest) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("init success"))
}

func (c *ChaincodeTest) create(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("create Error. ")
	}

	price, err := strconv.Atoi(args[0])
	if err != nil {
	}

	num, err := strconv.Atoi(args[1])
	if err != nil {
	}

	order := new(Order)

	Uuid, _ := uuid.NewV4()
	order.Uuid = Uuid.String()
	order.Price = price
	order.Num = num
	order.SumPay = price * num

	data, err := json.Marshal(order)
	if err != nil {
	}

	stub.PutState(order.Uuid, data)

	return shim.Success([]byte(order.Uuid))
}

func (c *ChaincodeTest) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	var resp pb.Response

	switch function {
	case "create":
		{
			resp = c.create(stub, args)
		}

	default:
		{
			return shim.Error("Invoke Error.  eg: create.")
		}
	}

	return resp

}

func main() {
	err := shim.Start(new(ChaincodeTest))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
