package netpeer

import (
	"context"
	"net"
	"testing"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Verifymember_Memberinfo(t *testing.T) {
	//
	ctx := context.Background()

	hosts := createVerifyHosts(t, 3)
	machineA := hosts[0]
	machineBYTE := hosts[1]
	machineC := hosts[2]

	//
	err := machineA.Connect(ctx, machineBYTE.AddressDetails())
	require.NoError(t, err)

	//
	err = machineC.Connect(ctx, machineA.AddressDetails())
	require.NoError(t, err)

	nodeBYTE, err := NewNode(machineA, machineBYTE.AddressDetails(), p2p.NoopStats(), false, false, false)
	require.NoError(t, err)

	nodeC, err := NewNode(machineA, machineC.AddressDetails(), p2p.NoopStats(), false, false, false)
	require.NoError(t, err)

	//
	//
	memberDetails, ok := nodeBYTE.MemberDetails().(p2p.StandardMemberDetails)
	require.True(t, ok, "REDACTED")
	assert.Equal(t, nodeBYTE.ID(), memberDetails.StandardMemberUID)
	assert.NotEmpty(t, memberDetails.ObserveAddress)

	assert.True(t, nodeBYTE.DistantIP().Equal(net.IPv4(127, 0, 0, 1)))
	assert.True(t, nodeBYTE.IsOutgoing())

	distantAddress, ok := nodeBYTE.DistantAddress().(*net.TCPAddr)
	require.True(t, ok)
	assert.NotZero(t, distantAddress.Port)

	//
	_, ok = nodeC.MemberDetails().(p2p.StandardMemberDetails)
	require.True(t, ok)
	assert.True(t, nodeC.IsOutgoing(), "REDACTED")
}

func Verifymember_Netinfotraversal(t *testing.T) {
	//
	//
	ctx := context.Background()

	hosts := createVerifyHosts(t, 4)
	machineA := hosts[0]

	ps := NewNodeCollection(machineA, p2p.NoopStats(), log.NewNoopTracer())

	for i := 1; i < len(hosts); i++ {
		err := machineA.Connect(ctx, hosts[i].AddressDetails())
		require.NoError(t, err)

		_, err = ps.Add(hosts[i].AddressDetails(), NodeAppendSettings{})
		require.NoError(t, err)
	}

	//
	tally := 0

	ps.ForEach(func(node p2p.Node) {
		//
		_, ok := node.MemberDetails().(p2p.StandardMemberDetails)
		require.True(t, ok, "REDACTED")
		_ = node.IsOutgoing()
		_ = node.DistantIP().String()
		_ = node.Status()

		tally++
	})

	assert.Equal(t, 3, tally)
}
