package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"

	cometjson "github.com/valkyrieworks/utils/json"
	coregrpc "github.com/valkyrieworks/rpc/grpc"
)

var grpcAddress = "REDACTED"

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("REDACTED")
		os.Exit(1)
	}
	tx := args[1]
	transferOctets, err := hex.DecodeString(tx)
	if err != nil {
		fmt.Println("REDACTED", err)
		os.Exit(1)
	}

	//
	customerGRPC := coregrpc.BeginGRPCCustomer(grpcAddress)
	res, err := customerGRPC.MulticastTransfer(context.Background(), &coregrpc.QueryMulticastTransfer{Tx: transferOctets})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bz, err := cometjson.Serialize(res)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bz))
}
