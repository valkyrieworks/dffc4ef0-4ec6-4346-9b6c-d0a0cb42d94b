package netpeer

import (
	"context"
	"testing"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyNodeCollection(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		//
		ctx := context.Background()

		hosts := createVerifyHosts(t, 2)
		machineA := hosts[0]
		machineBYTE := hosts[1]

		//
		err := machineA.Connect(ctx, machineBYTE.AddressDetails())
		require.NoError(t, err)

		ps := NewNodeCollection(machineA, p2p.NoopStats(), log.NewNoopTracer())
		nodeBYTEKey := nodeUIDToKey(machineBYTE.ID())

		//
		hasNode := ps.Has(nodeBYTEKey)
		assert.False(t, hasNode)

		acquiredNode := ps.Get(nodeBYTEKey)
		assert.Nil(t, acquiredNode)

		//
		nodeBYTE, err := ps.Add(machineBYTE.AddressDetails(), NodeAppendSettings{
			Internal:       true,
			Durable:    true,
			Absolute: true,
		})
		require.NoError(t, err)
		require.NotNil(t, nodeBYTE)

		//
		require.True(t, nodeBYTE.IsInternal())
		require.True(t, nodeBYTE.IsDurable())
		require.True(t, nodeBYTE.IsAbsolute())

		hasNode = ps.Has(nodeBYTEKey)
		assert.True(t, hasNode)

		acquiredNode = ps.Get(nodeBYTEKey)
		require.NotNil(t, acquiredNode)
		assert.Equal(t, nodeBYTEKey, acquiredNode.ID())

		//
		err = ps.Delete(nodeBYTEKey, NodeDeletionSettings{Cause: "REDACTED"})

		//
		require.NoError(t, err)

		hasNode = ps.Has(nodeBYTEKey)
		assert.False(t, hasNode)

		acquiredNode = ps.Get(nodeBYTEKey)
		assert.Nil(t, acquiredNode)
	})

	t.Run("REDACTED", func(t *testing.T) {
		t.Run("REDACTED", func(t *testing.T) {
			//
			hosts := createVerifyHosts(t, 2)
			machineA := hosts[0]
			machineBYTE := hosts[1]

			ps := NewNodeCollection(machineA, p2p.NoopStats(), log.NewNoopTracer())

			//
			_, err := ps.Add(peer.AddrInfo{ID: machineBYTE.ID()}, NodeAppendSettings{})

			//
			require.Error(t, err)
		})

		t.Run("REDACTED", func(t *testing.T) {
			//
			ctx := context.Background()

			hosts := createVerifyHosts(t, 2)
			machineA := hosts[0]
			machineBYTE := hosts[1]

			err := machineA.Connect(ctx, machineBYTE.AddressDetails())
			require.NoError(t, err)

			ps := NewNodeCollection(machineA, p2p.NoopStats(), log.NewNoopTracer())

			//
			_, fault1 := ps.Add(machineBYTE.AddressDetails(), NodeAppendSettings{})
			_, err2 := ps.Add(machineBYTE.AddressDetails(), NodeAppendSettings{})

			//
			require.NoError(t, fault1)
			require.ErrorIs(t, err2, ErrNodePresent)
		})

		t.Run("REDACTED", func(t *testing.T) {
			//
			hosts := createVerifyHosts(t, 2)
			machineA := hosts[0]
			machineBYTE := hosts[1]

			ps := NewNodeCollection(machineA, p2p.NoopStats(), log.NewNoopTracer())
			notExistingKey := nodeUIDToKey(machineBYTE.ID())

			//
			err := ps.Delete(notExistingKey, NodeDeletionSettings{Cause: "REDACTED"})

			//
			require.ErrorContains(t, err, "REDACTED")
		})
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		ctx := context.Background()
		const nodes = 6

		hosts := createVerifyHosts(t, nodes)
		machineA := hosts[0]

		//
		ps := NewNodeCollection(machineA, p2p.NoopStats(), log.NewNoopTracer())

		//
		for i := 1; i < nodes; i++ {
			err := machineA.Connect(ctx, hosts[i].AddressDetails())
			require.NoError(t, err)

			_, err = ps.Add(hosts[i].AddressDetails(), NodeAppendSettings{})
			require.NoError(t, err)
		}

		//
		//
		gatheredIDXDatastore := make(map[p2p.ID]struct{})
		ps.ForEach(func(p p2p.Node) {
			gatheredIDXDatastore[p.ID()] = struct{}{}
		})

		//
		assert.Equal(t, nodes-1, ps.Volume())
		assert.Equal(t, nodes-1, len(gatheredIDXDatastore))
	})

	t.Run("REDACTED", func(t *testing.T) {
		t.Run("REDACTED", func(t *testing.T) {
			//
			ctx := context.Background()

			hosts := createVerifyHosts(t, 4)
			machineA := hosts[0]
			ps := NewNodeCollection(machineA, p2p.NoopStats(), log.NewNoopTracer())

			//
			arbitraryNode := ps.Arbitrary()
			assert.Nil(t, arbitraryNode)

			//
			for i := 1; i < 4; i++ {
				err := machineA.Connect(ctx, hosts[i].AddressDetails())
				require.NoError(t, err)

				_, err = ps.Add(hosts[i].AddressDetails(), NodeAppendSettings{})
				require.NoError(t, err)
			}

			//
			arbitraryNode = ps.Arbitrary()
			require.NotNil(t, arbitraryNode)
			require.Contains(
				t,
				[]peer.ID{hosts[1].ID(), hosts[2].ID(), hosts[3].ID()},
				ps.keyToNodeUID(arbitraryNode.ID()),
			)
		})

		t.Run("REDACTED", func(t *testing.T) {
			//
			ctx := context.Background()

			hosts := createVerifyHosts(t, 4)
			machineA := hosts[0]
			ps := NewNodeCollection(machineA, p2p.NoopStats(), log.NewNoopTracer())

			anticipatedIDXDatastore := make([]p2p.ID, 0, 3)

			for i := 1; i < 4; i++ {
				err := machineA.Connect(ctx, hosts[i].AddressDetails())
				require.NoError(t, err)

				_, err = ps.Add(hosts[i].AddressDetails(), NodeAppendSettings{})
				require.NoError(t, err)

				anticipatedIDXDatastore = append(anticipatedIDXDatastore, nodeUIDToKey(hosts[i].ID()))
			}

			//
			replicated := ps.Clone()

			//
			require.Len(t, replicated, len(anticipatedIDXDatastore))

			replicatedIDXDatastore := make([]p2p.ID, len(replicated))
			for i, p := range replicated {
				replicatedIDXDatastore[i] = p.ID()
			}

			//
			for i := 1; i < len(replicated); i++ {
				assert.True(t, replicated[i-1].ID() < replicated[i].ID())
			}

			assert.ElementsMatch(t, anticipatedIDXDatastore, replicatedIDXDatastore)
		})
	})
}
