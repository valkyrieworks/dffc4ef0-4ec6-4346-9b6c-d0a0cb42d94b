package netp2p

import (
	"context"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyNodeAssign(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		//
		ctx := context.Background()

		machines := createVerifyMachines(t, 2)
		machineAN := machines[0]
		machineBYTE := machines[1]

		//
		err := machineAN.Connect(ctx, machineBYTE.LocationDetails())
		require.NoError(t, err)

		ps := FreshNodeAssign(machineAN, p2p.NooperationTelemetry(), log.FreshNooperationTracer())
		nodeBYTEToken := nodeUUIDTowardToken(machineBYTE.ID())

		//
		ownsNode := ps.Has(nodeBYTEToken)
		assert.False(t, ownsNode)

		attainedNode := ps.Get(nodeBYTEToken)
		assert.Nil(t, attainedNode)

		//
		nodeBYTE, err := ps.Add(machineBYTE.LocationDetails(), NodeAppendChoices{
			Secluded:       true,
			Enduring:    true,
			Absolute: true,
		})
		require.NoError(t, err)
		require.NotNil(t, nodeBYTE)

		//
		require.True(t, nodeBYTE.EqualsSecluded())
		require.True(t, nodeBYTE.EqualsEnduring())
		require.True(t, nodeBYTE.EqualsAbsolute())

		ownsNode = ps.Has(nodeBYTEToken)
		assert.True(t, ownsNode)

		attainedNode = ps.Get(nodeBYTEToken)
		require.NotNil(t, attainedNode)
		assert.Equal(t, nodeBYTEToken, attainedNode.ID())

		//
		err = ps.Discard(nodeBYTEToken, NodeDeletionChoices{Rationale: "REDACTED"})

		//
		require.NoError(t, err)

		ownsNode = ps.Has(nodeBYTEToken)
		assert.False(t, ownsNode)

		attainedNode = ps.Get(nodeBYTEToken)
		assert.Nil(t, attainedNode)
	})

	t.Run("REDACTED", func(t *testing.T) {
		t.Run("REDACTED", func(t *testing.T) {
			//
			machines := createVerifyMachines(t, 2)
			machineAN := machines[0]
			machineBYTE := machines[1]

			ps := FreshNodeAssign(machineAN, p2p.NooperationTelemetry(), log.FreshNooperationTracer())

			//
			_, err := ps.Add(peer.AddrInfo{ID: machineBYTE.ID()}, NodeAppendChoices{})

			//
			require.Error(t, err)
		})

		t.Run("REDACTED", func(t *testing.T) {
			//
			ctx := context.Background()

			machines := createVerifyMachines(t, 2)
			machineAN := machines[0]
			machineBYTE := machines[1]

			err := machineAN.Connect(ctx, machineBYTE.LocationDetails())
			require.NoError(t, err)

			ps := FreshNodeAssign(machineAN, p2p.NooperationTelemetry(), log.FreshNooperationTracer())

			//
			_, faultone := ps.Add(machineBYTE.LocationDetails(), NodeAppendChoices{})
			_, fault2 := ps.Add(machineBYTE.LocationDetails(), NodeAppendChoices{})

			//
			require.NoError(t, faultone)
			require.ErrorIs(t, fault2, FaultNodePresent)
		})

		t.Run("REDACTED", func(t *testing.T) {
			//
			machines := createVerifyMachines(t, 2)
			machineAN := machines[0]
			machineBYTE := machines[1]

			ps := FreshNodeAssign(machineAN, p2p.NooperationTelemetry(), log.FreshNooperationTracer())
			unPrevailingToken := nodeUUIDTowardToken(machineBYTE.ID())

			//
			err := ps.Discard(unPrevailingToken, NodeDeletionChoices{Rationale: "REDACTED"})

			//
			require.ErrorContains(t, err, "REDACTED")
		})
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		ctx := context.Background()
		const nodes = 6

		machines := createVerifyMachines(t, nodes)
		machineAN := machines[0]

		//
		ps := FreshNodeAssign(machineAN, p2p.NooperationTelemetry(), log.FreshNooperationTracer())

		//
		for i := 1; i < nodes; i++ {
			err := machineAN.Connect(ctx, machines[i].LocationDetails())
			require.NoError(t, err)

			_, err = ps.Add(machines[i].LocationDetails(), NodeAppendChoices{})
			require.NoError(t, err)
		}

		//
		//
		gatheredIDXDstore := make(map[p2p.ID]struct{})
		ps.ForeachEvery(func(p p2p.Node) {
			gatheredIDXDstore[p.ID()] = struct{}{}
		})

		//
		assert.Equal(t, nodes-1, ps.Extent())
		assert.Equal(t, nodes-1, len(gatheredIDXDstore))
	})

	t.Run("REDACTED", func(t *testing.T) {
		t.Run("REDACTED", func(t *testing.T) {
			//
			ctx := context.Background()

			machines := createVerifyMachines(t, 4)
			machineAN := machines[0]
			ps := FreshNodeAssign(machineAN, p2p.NooperationTelemetry(), log.FreshNooperationTracer())

			//
			unpredictableNode := ps.Unpredictable()
			assert.Nil(t, unpredictableNode)

			//
			for i := 1; i < 4; i++ {
				err := machineAN.Connect(ctx, machines[i].LocationDetails())
				require.NoError(t, err)

				_, err = ps.Add(machines[i].LocationDetails(), NodeAppendChoices{})
				require.NoError(t, err)
			}

			//
			unpredictableNode = ps.Unpredictable()
			require.NotNil(t, unpredictableNode)
			require.Contains(
				t,
				[]peer.ID{machines[1].ID(), machines[2].ID(), machines[3].ID()},
				ps.tokenTowardNodeUUID(unpredictableNode.ID()),
			)
		})

		t.Run("REDACTED", func(t *testing.T) {
			//
			ctx := context.Background()

			machines := createVerifyMachines(t, 4)
			machineAN := machines[0]
			ps := FreshNodeAssign(machineAN, p2p.NooperationTelemetry(), log.FreshNooperationTracer())

			anticipatedIDXDstore := make([]p2p.ID, 0, 3)

			for i := 1; i < 4; i++ {
				err := machineAN.Connect(ctx, machines[i].LocationDetails())
				require.NoError(t, err)

				_, err = ps.Add(machines[i].LocationDetails(), NodeAppendChoices{})
				require.NoError(t, err)

				anticipatedIDXDstore = append(anticipatedIDXDstore, nodeUUIDTowardToken(machines[i].ID()))
			}

			//
			replicated := ps.Duplicate()

			//
			require.Len(t, replicated, len(anticipatedIDXDstore))

			replicatedIDXDstore := make([]p2p.ID, len(replicated))
			for i, p := range replicated {
				replicatedIDXDstore[i] = p.ID()
			}

			//
			for i := 1; i < len(replicated); i++ {
				assert.True(t, replicated[i-1].ID() < replicated[i].ID())
			}

			assert.ElementsMatch(t, anticipatedIDXDstore, replicatedIDXDstore)
		})
	})
}
