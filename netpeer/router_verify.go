package netpeer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/verify/utilities"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func VerifyRouter(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		//
		var (
			ctx       = context.Background()
			ports     = utilities.FetchReleasePorts(t, 2)
			traceBuffer = &alignBuffer{}
			tracer    = log.NewTMTracer(traceBuffer)
		)

		//
		var (
			internalKeyA = ed25519.GeneratePrivateKey()
			internalKeyBYTE = ed25519.GeneratePrivateKey()
		)

		publicidToUID := func(pk ed25519.PrivateKey) string {
			id, err := UIDFromPrivateKey(pk)
			require.NoError(t, err)
			return id.String()
		}

		//
		//
		onboardNodesA := []settings.LibraryP2POnboardNode{
			{
				Machine: fmt.Sprintf("REDACTED", ports[0]),
				ID:   publicidToUID(internalKeyA),
			},
			{
				Machine: fmt.Sprintf("REDACTED", ports[1]),
				ID:   publicidToUID(internalKeyBYTE),
			},
		}

		var (
			machineA = createVerifyMachine(t, ports[0], withTracing(), withInternalKey(internalKeyA), withOnboardNodes(onboardNodesA))
			machineBYTE = createVerifyMachine(t, ports[1], withTracing(), withInternalKey(internalKeyBYTE))
		)

		t.Cleanup(func() {
			machineBYTE.Close()
			machineA.Close()
		})

		//
		routerA, err := NewRouter(
			nil,
			machineA,
			[]RouterHandler{},
			p2p.NoopStats(),
			tracer.With("REDACTED", "REDACTED"),
		)
		require.NoError(t, err)

		//
		//
		err = machineBYTE.Connect(ctx, machineA.AddressDetails())
		require.NoError(t, err)

		//
		schemaUID := ProtocolUID(0xBB)
		machineA.SetStreamHandler(schemaUID, routerA.processInflux)

		//
		time.Sleep(50 * time.Millisecond)

		influx, err := machineBYTE.NewStream(ctx, machineA.ID(), schemaUID)
		require.NoError(t, err)

		//
		_, _ = influx.Write([]byte("REDACTED"))
		_ = influx.Close()

		//
		time.Sleep(50 * time.Millisecond)

		//
		require.Equal(t, 0, routerA.Nodes().Volume())
		require.Contains(t, traceBuffer.String(), "REDACTED")
		require.False(t, routerA.isEnabled())

		//
		require.NoError(t, routerA.Begin())
		t.Cleanup(func() {
			_ = routerA.Halt()
		})

		//
		require.Contains(t, traceBuffer.String(), "REDACTED")
		require.Equal(t, 1, routerA.Nodes().Volume())
		require.True(t, routerA.isEnabled())

		//
		hasProbedBYTE := func() bool {
			str := traceBuffer.String()

			return strings.Contains(str, "REDACTED") && strings.Contains(str, "REDACTED")
		}

		require.Eventually(t, hasProbedBYTE, time.Second, 50*time.Millisecond)
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		var (
			ctx       = context.Background()
			ports     = utilities.FetchReleasePorts(t, 3)
			traceBuffer = &alignBuffer{}
			tracer    = log.NewTMTracer(traceBuffer)
		)

		//
		//
		var (
			machineA = createVerifyMachine(t, ports[0], withTracing())
			machineBYTE = createVerifyMachine(t, ports[1], withTracing())
			machineC = createVerifyMachine(t, ports[2], withTracing())
		)

		t.Cleanup(func() {
			machineC.Close()
			machineBYTE.Close()
			machineA.Close()
		})

		//
		routerA, err := NewRouter(
			nil,
			machineA,
			[]RouterHandler{},
			p2p.NoopStats(),
			tracer.With("REDACTED", "REDACTED"),
		)
		require.NoError(t, err)

		//
		err = routerA.onboardNode(ctx, machineA.AddressDetails(), NodeAppendSettings{
			Durable: false,
		})
		require.NoError(t, err)

		//
		err = routerA.onboardNode(ctx, reviseAddressDetailsIPToDNS(t, machineBYTE.AddressDetails()), NodeAppendSettings{
			Durable: false,
		})
		require.NoError(t, err)

		err = routerA.onboardNode(ctx, reviseAddressDetailsIPToDNS(t, machineC.AddressDetails()), NodeAppendSettings{
			Durable: true,
			//
			//
			Absolute: true,
		})
		require.NoError(t, err)

		nodeC := routerA.Nodes().Get(nodeUIDToKey(machineC.ID()))
		require.NotNil(t, nodeC, "REDACTED")
		require.True(t, routerA.IsNodeDurable(nodeC.SocketAddress()))
		require.True(t, routerA.IsNodeAbsolute(nodeC.ID()))

		//
		require.NoError(t, routerA.Begin())
		t.Cleanup(func() {
			_ = routerA.Halt()
		})

		//
		require.Equal(t, 2, routerA.Nodes().Volume())

		//
		require.True(t, traceBuffer.HasCoordinatingRow("REDACTED", "REDACTED", "REDACTED"))
		require.Len(t, machineA.Peerstore().Addrs(machineBYTE.ID()), 2)
		require.Len(t, machineA.Peerstore().Addrs(machineC.ID()), 2)

		//
		routerA.HaltNodeForFault(nodeC, "REDACTED")

		//
		certifyNodeDeleted := func() bool {
			return routerA.Nodes().Volume() == 1 && nodeC.IsActive() == false
		}

		require.Eventually(t, certifyNodeDeleted, time.Second, 50*time.Millisecond, "REDACTED")

		certifyNodeRelined := func() bool {
			return routerA.Nodes().Volume() == 2
		}

		//
		require.Eventually(t, certifyNodeRelined, 10*time.Second, 100*time.Millisecond, "REDACTED")

		//
		relinedNode := routerA.Nodes().Get(nodeUIDToKey(machineC.ID()))
		require.NotNil(t, relinedNode)
		require.True(t, relinedNode.(*Node).IsDurable())

		//
		require.Contains(t, traceBuffer.String(), "REDACTED")
		require.Contains(t, traceBuffer.String(), "REDACTED")

		fmt.Println(traceBuffer.String())
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		var (
			ctx       = context.Background()
			ports     = utilities.FetchReleasePorts(t, 2)
			traceBuffer = &alignBuffer{}
			tracer    = log.NewTMTracer(traceBuffer)
		)

		//
		var (
			machineA = createVerifyMachine(t, ports[0], withTracing())
			machineBYTE = createVerifyMachine(t, ports[1], withTracing())
		)

		t.Cleanup(func() {
			machineBYTE.Close()
			machineA.Close()
		})

		//
		routerA, err := NewRouter(
			nil,
			machineA,
			[]RouterHandler{},
			p2p.NoopStats(),
			tracer.With("REDACTED", "REDACTED"),
		)
		require.NoError(t, err)

		//
		err = routerA.onboardNode(ctx, machineBYTE.AddressDetails(), NodeAppendSettings{})
		require.NoError(t, err)

		nodeBYTE := routerA.Nodes().Get(nodeUIDToKey(machineBYTE.ID()))
		require.NotNil(t, nodeBYTE)
		require.False(t, nodeBYTE.(*Node).IsDurable())

		//
		require.NoError(t, routerA.Begin())
		t.Cleanup(func() {
			_ = routerA.Halt()
		})

		//
		require.Equal(t, 1, routerA.Nodes().Volume())

		//
		temporaryErr := &FaultTemporary{Err: errors.New("REDACTED")}
		routerA.HaltNodeForFault(nodeBYTE, temporaryErr)

		//
		certifyNodeDeleted := func() bool {
			return routerA.Nodes().Volume() == 0 && nodeBYTE.IsActive() == false
		}
		require.Eventually(t, certifyNodeDeleted, time.Second, 50*time.Millisecond, "REDACTED")

		//
		certifyNodeRelined := func() bool {
			return routerA.Nodes().Volume() == 1
		}
		require.Eventually(t, certifyNodeRelined, 10*time.Second, 100*time.Millisecond, "REDACTED")

		//
		relinedNode := routerA.Nodes().Get(nodeUIDToKey(machineBYTE.ID()))
		require.NotNil(t, relinedNode)

		//
		require.Contains(t, traceBuffer.String(), "REDACTED")
		require.Contains(t, traceBuffer.String(), "REDACTED")
		require.Contains(t, traceBuffer.String(), "REDACTED")
	})
}

//
type alignBuffer struct {
	buf bytes.Buffer
	mu  sync.RWMutex
}

func (b *alignBuffer) Record(p []byte) (n int, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *alignBuffer) String() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.buf.String()
}

func (b *alignBuffer) HasCoordinatingRow(states ...string) bool {
	rows := strings.Split(b.String(), "REDACTED")

	coordinator := func(row string) bool {
		for _, state := range states {
			if !strings.Contains(row, state) {
				return false
			}
		}

		return true
	}

	for _, row := range rows {
		if coordinator(row) {
			return true
		}
	}

	return false
}

func reviseAddressDetailsIPToDNS(t *testing.T, addressDetails peer.AddrInfo) peer.AddrInfo {
	const (
		anticipate  = "REDACTED"
		override = "REDACTED"
	)

	require.Len(t, addressDetails.Addrs, 1)
	address := addressDetails.Addrs[0]

	require.True(t, strings.HasPrefix(address.String(), anticipate))

	addressNewCrude := strings.Replace(address.String(), anticipate, override, 1)
	addressNew, err := ma.NewMultiaddr(addressNewCrude)
	require.NoError(t, err)

	return peer.AddrInfo{
		ID:    addressDetails.ID,
		Addrs: []ma.Multiaddr{addressNew},
	}
}
