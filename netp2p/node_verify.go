package netp2p

import (
	"context"
	"net"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Mockpeer_Nodeinfo(t *testing.T) {
	//
	ctx := context.Background()

	machines := createVerifyMachines(t, 3)
	machineAN := machines[0]
	machineBYTE := machines[1]
	machineCN := machines[2]

	//
	err := machineAN.Connect(ctx, machineBYTE.LocationDetails())
	require.NoError(t, err)

	//
	err = machineCN.Connect(ctx, machineAN.LocationDetails())
	require.NoError(t, err)

	nodeBYTE, err := FreshNode(machineAN, machineBYTE.LocationDetails(), p2p.NooperationTelemetry(), false, false, false)
	require.NoError(t, err)

	nodeCount, err := FreshNode(machineAN, machineCN.LocationDetails(), p2p.NooperationTelemetry(), false, false, false)
	require.NoError(t, err)

	//
	//
	peerDetails, ok := nodeBYTE.PeerDetails().(p2p.FallbackPeerDetails)
	require.True(t, ok, "REDACTED")
	assert.Equal(t, nodeBYTE.ID(), peerDetails.FallbackPeerUUID)
	assert.NotEmpty(t, peerDetails.OverhearLocation)

	assert.True(t, nodeBYTE.DistantINET().Equal(net.IPv4(127, 0, 0, 1)))
	assert.True(t, nodeBYTE.EqualsOutgoing())

	distantLocation, ok := nodeBYTE.DistantLocation().(*net.TCPAddr)
	require.True(t, ok)
	assert.NotZero(t, distantLocation.Port)

	//
	_, ok = nodeCount.PeerDetails().(p2p.FallbackPeerDetails)
	require.True(t, ok)
	assert.True(t, nodeCount.EqualsOutgoing(), "REDACTED")
}

func Mockpeer_Networkinfoiteration(t *testing.T) {
	//
	//
	ctx := context.Background()

	machines := createVerifyMachines(t, 4)
	machineAN := machines[0]

	ps := FreshNodeAssign(machineAN, p2p.NooperationTelemetry(), log.FreshNooperationTracer())

	for i := 1; i < len(machines); i++ {
		err := machineAN.Connect(ctx, machines[i].LocationDetails())
		require.NoError(t, err)

		_, err = ps.Add(machines[i].LocationDetails(), NodeAppendChoices{})
		require.NoError(t, err)
	}

	//
	tally := 0

	ps.ForeachEvery(func(node p2p.Node) {
		//
		_, ok := node.PeerDetails().(p2p.FallbackPeerDetails)
		require.True(t, ok, "REDACTED")
		_ = node.EqualsOutgoing()
		_ = node.DistantINET().String()
		_ = node.Condition()

		tally++
	})

	assert.Equal(t, 3, tally)
}
