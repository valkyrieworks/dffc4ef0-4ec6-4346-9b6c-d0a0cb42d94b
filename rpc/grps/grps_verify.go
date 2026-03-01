package cocode_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	base_grps "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/grps"
	rpcoverify "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/verify"
)

func VerifyPrimary(m *testing.M) {
	//
	app := statedepot.FreshInsideRamPlatform()
	peer := rpcoverify.InitiateStrongmind(app)

	cipher := m.Run()

	//
	rpcoverify.HaltStrongmind(peer)
	os.Exit(cipher)
}

func VerifyMulticastTransfer(t *testing.T) {
	res, err := rpcoverify.FetchGRPSCustomer().MulticastTransfer(
		context.Background(),
		&base_grps.SolicitMulticastTransfer{Tx: statedepot.FreshTransfer("REDACTED", "REDACTED")},
	)
	require.NoError(t, err)
	require.EqualValues(t, 0, res.InspectTransfer.Cipher)
	require.EqualValues(t, 0, res.TransferOutcome.Cipher)
}
