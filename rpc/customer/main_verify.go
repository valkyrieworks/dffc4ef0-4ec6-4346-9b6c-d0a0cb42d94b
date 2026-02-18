package agent_test

import (
	"os"
	"testing"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	nm "github.com/valkyrieworks/member"
	rpctest "github.com/valkyrieworks/rpc/verify"
)

var member *nm.Member

func VerifyMain(m *testing.M) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}

	app := objectdepot.NewDurableSoftware(dir)
	//
	//
	member = rpctest.BeginConsensuscore(app)

	code := m.Run()

	//
	rpctest.HaltConsensuscore(member)
	_ = os.RemoveAll(dir)
	os.Exit(code)
}
