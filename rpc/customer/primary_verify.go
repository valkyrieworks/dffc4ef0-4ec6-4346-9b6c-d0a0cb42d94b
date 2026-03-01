package cust_test

import (
	"os"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	nm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/peer"
	rpcoverify "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/verify"
)

var peer *nm.Peer

func VerifyPrimary(m *testing.M) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}

	app := statedepot.FreshEnduringPlatform(dir)
	//
	//
	peer = rpcoverify.InitiateStrongmind(app)

	cipher := m.Run()

	//
	rpcoverify.HaltStrongmind(peer)
	_ = os.RemoveAll(dir)
	os.Exit(cipher)
}
