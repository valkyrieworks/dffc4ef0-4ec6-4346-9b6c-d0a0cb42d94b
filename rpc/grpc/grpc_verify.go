package baseg_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/iface/instance/objectdepot"
	core_grpc "github.com/valkyrieworks/rpc/grpc"
	rpctest "github.com/valkyrieworks/rpc/verify"
)

func VerifyMain(m *testing.M) {
	//
	app := objectdepot.NewInRamSoftware()
	member := rpctest.BeginConsensuscore(app)

	code := m.Run()

	//
	rpctest.HaltConsensuscore(member)
	os.Exit(code)
}

func VerifyMulticastTransfer(t *testing.T) {
	res, err := rpctest.FetchGRPCCustomer().MulticastTransfer(
		context.Background(),
		&core_grpc.QueryMulticastTransfer{Tx: objectdepot.NewTransfer("REDACTED", "REDACTED")},
	)
	require.NoError(t, err)
	require.EqualValues(t, 0, res.InspectTransfer.Code)
	require.EqualValues(t, 0, res.TransOutcome.Code)
}
