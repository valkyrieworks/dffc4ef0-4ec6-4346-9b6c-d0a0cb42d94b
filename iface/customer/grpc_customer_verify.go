package abcic_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	ifaceservice "github.com/valkyrieworks/iface/host"
	"github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/log"
)

func VerifyGRPC(t *testing.T) {
	app := kinds.NewRootSoftware()
	countInspectTrans := 2000
	socketEntry := fmt.Sprintf("REDACTED", rand.Int31n(1<<30))
	defer os.Remove(socketEntry)
	socket := fmt.Sprintf("REDACTED", socketEntry)

	//
	host := ifaceservice.NewGRPCHost(socket, app)
	host.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
	err := host.Begin()
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := host.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	link, err := grpc.NewClient(socket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := link.Close(); err != nil {
			t.Error(err)
		}
	})

	customer := kinds.NewIfaceCustomer(link)

	//
	for tally := 0; tally < countInspectTrans; tally++ {
		//
		reply, err := customer.InspectTransfer(context.Background(), &kinds.QueryInspectTransfer{Tx: []byte("REDACTED")})
		require.NoError(t, err)
		if reply.Code != 0 {
			t.Error("REDACTED", reply.Code)
		}
		t.Log("REDACTED", tally)
	}
}
