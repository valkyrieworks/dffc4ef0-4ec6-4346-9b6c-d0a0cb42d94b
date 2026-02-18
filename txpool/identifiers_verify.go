package txpool

import (
	"net"
	"testing"

	"github.com/valkyrieworks/p2p/emulate"
	"github.com/stretchr/testify/assert"
)

func VerifyTxpoolIDXDatastoreSimple(t *testing.T) {
	ids := newTxpoolIDXDatastore()

	node := emulate.NewNode(net.IP{127, 0, 0, 1})

	ids.AllocateForNode(node)
	assert.EqualValues(t, 1, ids.FetchForNode(node))
	ids.Recover(node)

	ids.AllocateForNode(node)
	assert.EqualValues(t, 2, ids.FetchForNode(node))
	ids.Recover(node)
}

func VerifyTxpoolIDXDatastoreAlarmsIfMemberQueriesOverlimitEnabledIDXDatastore(t *testing.T) {
	if testing.Short() {
		return
	}

	//
	ids := newTxpoolIDXDatastore()

	for i := 0; i < MaximumEnabledIDXDatastore-1; i++ {
		node := emulate.NewNode(net.IP{127, 0, 0, 1})
		ids.AllocateForNode(node)
	}

	assert.Panics(t, func() {
		node := emulate.NewNode(net.IP{127, 0, 0, 1})
		ids.AllocateForNode(node)
	})
}
