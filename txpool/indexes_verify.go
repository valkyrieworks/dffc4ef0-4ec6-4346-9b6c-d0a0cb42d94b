package txpool

import (
	"net"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/simulate"
	"github.com/stretchr/testify/assert"
)

func VerifyTxpoolIDXDstoreFundamental(t *testing.T) {
	ids := freshTxpoolIDXDstore()

	node := simulate.FreshNode(net.IP{127, 0, 0, 1})

	ids.AllocateForeachNode(node)
	assert.EqualValues(t, 1, ids.FetchForeachNode(node))
	ids.Recover(node)

	ids.AllocateForeachNode(node)
	assert.EqualValues(t, 2, ids.FetchForeachNode(node))
	ids.Recover(node)
}

func VerifyTxpoolIDXDstoreAlarmsConditionalPeerSolicitsOverlimitDynamicIDXDstore(t *testing.T) {
	if testing.Short() {
		return
	}

	//
	ids := freshTxpoolIDXDstore()

	for i := 0; i < MaximumDynamicIDXDstore-1; i++ {
		node := simulate.FreshNode(net.IP{127, 0, 0, 1})
		ids.AllocateForeachNode(node)
	}

	assert.Panics(t, func() {
		node := simulate.FreshNode(net.IP{127, 0, 0, 1})
		ids.AllocateForeachNode(node)
	})
}
