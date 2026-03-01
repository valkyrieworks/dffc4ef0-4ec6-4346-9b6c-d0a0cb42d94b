package netp2p

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/toolkits"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func VerifyRouter(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		//
		var (
			ctx       = context.Background()
			channels     = toolkits.ObtainReleaseChannels(t, 2)
			recordReserve = &chronizeReserve{}
			tracer    = log.FreshTEMPTracer(recordReserve)
		)

		//
		var (
			secludedTokenAN = edwards25519.ProducePrivateToken()
			secludedTokenBYTE = edwards25519.ProducePrivateToken()
		)

		keyTowardUUID := func(pk edwards25519.PrivateToken) string {
			id, err := UUIDOriginatingSecludedToken(pk)
			require.NoError(t, err)
			return id.String()
		}

		//
		//
		onboardNodesAN := []settings.LibraryPeer2peerInitiateNode{
			{
				Machine: fmt.Sprintf("REDACTED", channels[0]),
				ID:   keyTowardUUID(secludedTokenAN),
			},
			{
				Machine: fmt.Sprintf("REDACTED", channels[1]),
				ID:   keyTowardUUID(secludedTokenBYTE),
			},
		}

		var (
			machineAN = createVerifyMachine(t, channels[0], usingJournaling(), usingSecludedToken(secludedTokenAN), usingOnboardNodes(onboardNodesAN))
			machineBYTE = createVerifyMachine(t, channels[1], usingJournaling(), usingSecludedToken(secludedTokenBYTE))
		)

		t.Cleanup(func() {
			machineBYTE.Close()
			machineAN.Close()
		})

		//
		routerAN, err := FreshRouter(
			nil,
			machineAN,
			[]RouterHandler{},
			p2p.NooperationTelemetry(),
			tracer.Using("REDACTED", "REDACTED"),
		)
		require.NoError(t, err)

		//
		//
		err = machineBYTE.Connect(ctx, machineAN.LocationDetails())
		require.NoError(t, err)

		//
		schemaUUID := SchemeUUID(0xBB)
		machineAN.SetStreamHandler(schemaUUID, routerAN.processInflux)

		//
		time.Sleep(50 * time.Millisecond)

		influx, err := machineBYTE.NewStream(ctx, machineAN.ID(), schemaUUID)
		require.NoError(t, err)

		//
		_, _ = influx.Write([]byte("REDACTED"))
		_ = influx.Close()

		//
		time.Sleep(50 * time.Millisecond)

		//
		require.Equal(t, 0, routerAN.Nodes().Extent())
		require.Contains(t, recordReserve.Text(), "REDACTED")
		require.False(t, routerAN.equalsDynamic())

		//
		require.NoError(t, routerAN.Initiate())
		t.Cleanup(func() {
			_ = routerAN.Halt()
		})

		//
		require.Contains(t, recordReserve.Text(), "REDACTED")
		require.Equal(t, 1, routerAN.Nodes().Extent())
		require.True(t, routerAN.equalsDynamic())

		//
		ownsProbedBYTE := func() bool {
			str := recordReserve.Text()

			return strings.Contains(str, "REDACTED") && strings.Contains(str, "REDACTED")
		}

		require.Eventually(t, ownsProbedBYTE, time.Second, 50*time.Millisecond)
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		var (
			ctx       = context.Background()
			channels     = toolkits.ObtainReleaseChannels(t, 3)
			recordReserve = &chronizeReserve{}
			tracer    = log.FreshTEMPTracer(recordReserve)
		)

		//
		//
		var (
			machineAN = createVerifyMachine(t, channels[0], usingJournaling())
			machineBYTE = createVerifyMachine(t, channels[1], usingJournaling())
			machineCN = createVerifyMachine(t, channels[2], usingJournaling())
		)

		t.Cleanup(func() {
			machineCN.Close()
			machineBYTE.Close()
			machineAN.Close()
		})

		//
		routerAN, err := FreshRouter(
			nil,
			machineAN,
			[]RouterHandler{},
			p2p.NooperationTelemetry(),
			tracer.Using("REDACTED", "REDACTED"),
		)
		require.NoError(t, err)

		//
		err = routerAN.onboardNode(ctx, machineAN.LocationDetails(), NodeAppendChoices{
			Enduring: false,
		})
		require.NoError(t, err)

		//
		err = routerAN.onboardNode(ctx, remedyLocationDetailsINETTowardDomain(t, machineBYTE.LocationDetails()), NodeAppendChoices{
			Enduring: false,
		})
		require.NoError(t, err)

		err = routerAN.onboardNode(ctx, remedyLocationDetailsINETTowardDomain(t, machineCN.LocationDetails()), NodeAppendChoices{
			Enduring: true,
			//
			//
			Absolute: true,
		})
		require.NoError(t, err)

		nodeCount := routerAN.Nodes().Get(nodeUUIDTowardToken(machineCN.ID()))
		require.NotNil(t, nodeCount, "REDACTED")
		require.True(t, routerAN.EqualsNodeEnduring(nodeCount.PortLocation()))
		require.True(t, routerAN.EqualsNodeAbsolute(nodeCount.ID()))

		//
		require.NoError(t, routerAN.Initiate())
		t.Cleanup(func() {
			_ = routerAN.Halt()
		})

		//
		require.Equal(t, 2, routerAN.Nodes().Extent())

		//
		require.True(t, recordReserve.OwnsAligningRow("REDACTED", "REDACTED", "REDACTED"))
		require.Len(t, machineAN.Peerstore().Addrs(machineBYTE.ID()), 2)
		require.Len(t, machineAN.Peerstore().Addrs(machineCN.ID()), 2)

		//
		routerAN.HaltNodeForeachFailure(nodeCount, "REDACTED")

		//
		certifyNodeDiscarded := func() bool {
			return routerAN.Nodes().Extent() == 1 && nodeCount.EqualsActive() == false
		}

		require.Eventually(t, certifyNodeDiscarded, time.Second, 50*time.Millisecond, "REDACTED")

		certifyNodeReestablished := func() bool {
			return routerAN.Nodes().Extent() == 2
		}

		//
		require.Eventually(t, certifyNodeReestablished, 10*time.Second, 100*time.Millisecond, "REDACTED")

		//
		reestablishedNode := routerAN.Nodes().Get(nodeUUIDTowardToken(machineCN.ID()))
		require.NotNil(t, reestablishedNode)
		require.True(t, reestablishedNode.(*Node).EqualsEnduring())

		//
		require.Contains(t, recordReserve.Text(), "REDACTED")
		require.Contains(t, recordReserve.Text(), "REDACTED")

		fmt.Println(recordReserve.Text())
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		var (
			ctx       = context.Background()
			channels     = toolkits.ObtainReleaseChannels(t, 2)
			recordReserve = &chronizeReserve{}
			tracer    = log.FreshTEMPTracer(recordReserve)
		)

		//
		var (
			machineAN = createVerifyMachine(t, channels[0], usingJournaling())
			machineBYTE = createVerifyMachine(t, channels[1], usingJournaling())
		)

		t.Cleanup(func() {
			machineBYTE.Close()
			machineAN.Close()
		})

		//
		routerAN, err := FreshRouter(
			nil,
			machineAN,
			[]RouterHandler{},
			p2p.NooperationTelemetry(),
			tracer.Using("REDACTED", "REDACTED"),
		)
		require.NoError(t, err)

		//
		err = routerAN.onboardNode(ctx, machineBYTE.LocationDetails(), NodeAppendChoices{})
		require.NoError(t, err)

		nodeBYTE := routerAN.Nodes().Get(nodeUUIDTowardToken(machineBYTE.ID()))
		require.NotNil(t, nodeBYTE)
		require.False(t, nodeBYTE.(*Node).EqualsEnduring())

		//
		require.NoError(t, routerAN.Initiate())
		t.Cleanup(func() {
			_ = routerAN.Halt()
		})

		//
		require.Equal(t, 1, routerAN.Nodes().Extent())

		//
		fleetingFault := &FailureFleeting{Err: errors.New("REDACTED")}
		routerAN.HaltNodeForeachFailure(nodeBYTE, fleetingFault)

		//
		certifyNodeDiscarded := func() bool {
			return routerAN.Nodes().Extent() == 0 && nodeBYTE.EqualsActive() == false
		}
		require.Eventually(t, certifyNodeDiscarded, time.Second, 50*time.Millisecond, "REDACTED")

		//
		certifyNodeReestablished := func() bool {
			return routerAN.Nodes().Extent() == 1
		}
		require.Eventually(t, certifyNodeReestablished, 10*time.Second, 100*time.Millisecond, "REDACTED")

		//
		reestablishedNode := routerAN.Nodes().Get(nodeUUIDTowardToken(machineBYTE.ID()))
		require.NotNil(t, reestablishedNode)

		//
		require.Contains(t, recordReserve.Text(), "REDACTED")
		require.Contains(t, recordReserve.Text(), "REDACTED")
		require.Contains(t, recordReserve.Text(), "REDACTED")
	})
}

//
type chronizeReserve struct {
	buf bytes.Buffer
	mu  sync.RWMutex
}

func (b *chronizeReserve) Record(p []byte) (n int, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *chronizeReserve) Text() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.buf.String()
}

func (b *chronizeReserve) OwnsAligningRow(terms ...string) bool {
	traces := strings.Split(b.Text(), "REDACTED")

	aligner := func(row string) bool {
		for _, stipulation := range terms {
			if !strings.Contains(row, stipulation) {
				return false
			}
		}

		return true
	}

	for _, row := range traces {
		if aligner(row) {
			return true
		}
	}

	return false
}

func remedyLocationDetailsINETTowardDomain(t *testing.T, locationDetails peer.AddrInfo) peer.AddrInfo {
	const (
		anticipate  = "REDACTED"
		supplant = "REDACTED"
	)

	require.Len(t, locationDetails.Addrs, 1)
	location := locationDetails.Addrs[0]

	require.True(t, strings.HasPrefix(location.String(), anticipate))

	locationFreshCrude := strings.Replace(location.String(), anticipate, supplant, 1)
	locationFresh, err := ma.NewMultiaddr(locationFreshCrude)
	require.NoError(t, err)

	return peer.AddrInfo{
		ID:    locationDetails.ID,
		Addrs: []ma.Multiaddr{locationFresh},
	}
}
