package ifacec_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	abcimaster "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/node"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

func VerifyGRPS(t *testing.T) {
	app := kinds.FreshFoundationPlatform()
	countInspectTrans := 2000
	portRecord := fmt.Sprintf("REDACTED", rand.Int31n(1<<30))
	defer os.Remove(portRecord)
	port := fmt.Sprintf("REDACTED", portRecord)

	//
	node := abcimaster.FreshGRPSDaemon(port, app)
	node.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
	err := node.Initiate()
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := node.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	link, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := link.Close(); err != nil {
			t.Error(err)
		}
	})

	customer := kinds.FreshIfaceCustomer(link)

	//
	for tally := 0; tally < countInspectTrans; tally++ {
		//
		reply, err := customer.InspectTransfer(context.Background(), &kinds.SolicitInspectTransfer{Tx: []byte("REDACTED")})
		require.NoError(t, err)
		if reply.Cipher != 0 {
			t.Error("REDACTED", reply.Cipher)
		}
		t.Log("REDACTED", tally)
	}
}
