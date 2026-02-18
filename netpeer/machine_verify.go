package netpeer

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/verify/utilities"
	"github.com/libp2p/go-libp2p/core/network"
	corepeer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func VerifyMachine(t *testing.T) {
	//
	ctx := context.Background()

	//
	var (
		conduitFoo = byte(0xaa)
		conduitBar = byte(0xbb)
		schemaFoo   = ProtocolUID(conduitFoo)
		schemaBar   = ProtocolUID(conduitBar)
	)

	//
	ports := utilities.FetchReleasePorts(t, 2)

	//
	host1 := createVerifyMachine(t, ports[0], withTracing())
	host2 := createVerifyMachine(t, ports[1], withTracing(), withOnboardNodes([]settings.LibraryP2POnboardNode{
		{
			//
			Machine: fmt.Sprintf("REDACTED", ports[0]),
			ID:   host1.ID().String(),
		},
	}))

	establishOnboardNodes(t, ctx, host2, host2.OnboardNodes())

	t.Logf("REDACTED", host1.AddressDetails())
	t.Logf("REDACTED", host2.AddressDetails())

	t.Cleanup(func() {
		host2.Close()
		host1.Close()
	})

	//
	type packet struct {
		protocol protocol.ID
		emitter   corepeer.ID
		subscriber corepeer.ID
		signal  string
	}

	wrappers := []packet{}
	mu := sync.Mutex{}

	//
	manager := func(influx network.Stream) {
		var (
			link     = influx.Conn()
			subscriber = link.LocalPeer()
			emitter   = link.RemotePeer()
		)

		if link.ConnState().Transport != CarrierQUIC {
			t.Fatalf("REDACTED", link.ConnState().Transport)
			return
		}

		shipment, err := InfluxFetchEnd(influx)
		if err != nil {
			t.Fatalf("REDACTED", emitter, err)
			return
		}

		msg := &kinds.Query{}
		require.NoError(t, msg.Unserialize(shipment))
		require.NotNil(t, msg.FetchReverberate())

		e := packet{
			protocol: influx.Protocol(),
			emitter:   emitter,
			subscriber: subscriber,
			signal:  msg.FetchReverberate().FetchSignal(),
		}

		traceSignal := e.signal
		if len(traceSignal) > 64 {
			traceSignal = traceSignal[:64] + "REDACTED"
		}

		t.Logf(
			"REDACTED",
			e.emitter.String(),
			e.subscriber.String(),
			e.protocol,
			traceSignal,
		)

		mu.Lock()
		defer mu.Unlock()

		wrappers = append(wrappers, e)
	}

	host1.SetStreamHandler(schemaFoo, manager)
	host1.SetStreamHandler(schemaBar, manager)

	host2.SetStreamHandler(schemaFoo, manager)
	host2.SetStreamHandler(schemaBar, manager)

	//
	host1member2, err := NewNode(host1, host2.AddressDetails(), p2p.NoopStats(), false, false, false)
	require.NoError(t, err, "REDACTED")
	require.NoError(t, host1member2.Begin(), "REDACTED")

	host2member1, err := NewNode(host2, host1.AddressDetails(), p2p.NoopStats(), false, false, false)
	require.NoError(t, err, "REDACTED")
	require.NoError(t, host2member1.Begin(), "REDACTED")

	t.Logf("REDACTED", host1member2.ID())
	t.Logf("REDACTED", host2member1.ID())

	//
	//
	lengthyStr := strings.Repeat("REDACTED", 300*1024)

	//
	transmit1 := host1member2.Transmit(p2p.Packet{
		StreamUID: conduitFoo,
		Signal:   kinds.ToQueryReverberate("REDACTED"),
	})

	transmit2 := host2member1.Transmit(p2p.Packet{
		StreamUID: conduitBar,
		Signal:   kinds.ToQueryReverberate("REDACTED"),
	})

	transmit3 := host1member2.AttemptTransmit(p2p.Packet{
		StreamUID: conduitBar,
		Signal:   kinds.ToQueryReverberate(lengthyStr),
	})

	//
	//
	require.True(t, transmit1, "REDACTED")
	require.True(t, transmit2, "REDACTED")
	require.True(t, transmit3, "REDACTED")

	//
	wait := func() bool {
		mu.Lock()
		defer mu.Unlock()
		return len(wrappers) == 3
	}

	require.Eventually(t, wait, 500*time.Millisecond, 50*time.Millisecond)

	//
	anticipatedWrappers := []packet{
		{
			protocol: schemaFoo,
			emitter:   host1.ID(),
			subscriber: host2.ID(),
			signal:  "REDACTED",
		},
		{
			protocol: schemaBar,
			emitter:   host2.ID(),
			subscriber: host1.ID(),
			signal:  "REDACTED",
		},
		{
			protocol: schemaBar,
			emitter:   host1.ID(),
			subscriber: host2.ID(),
			signal:  lengthyStr,
		},
	}

	require.ElementsMatch(t, anticipatedWrappers, wrappers)

	t.Run("REDACTED", func(t *testing.T) {
		//
		ctx, revoke := context.WithTimeout(context.Background(), 5*time.Second)
		defer revoke()

		//
		rtt, err := host1.Ping(ctx, host2.AddressDetails())

		//
		require.NoError(t, err)
		require.NotZero(t, rtt)

		t.Logf("REDACTED", rtt.String())
	})
}

type verifyOpts struct {
	pk             ed25519.PrivateKey
	onboardNodes []settings.LibraryP2POnboardNode
	activateTracing  bool
}

type verifySetting func(*verifyOpts)

func withTracing() verifySetting {
	return func(opts *verifyOpts) { opts.activateTracing = true }
}

func withOnboardNodes(onboardNodes []settings.LibraryP2POnboardNode) verifySetting {
	return func(opts *verifyOpts) { opts.onboardNodes = onboardNodes }
}

func withInternalKey(pk ed25519.PrivateKey) verifySetting {
	return func(opts *verifyOpts) { opts.pk = pk }
}

func createVerifyMachine(t *testing.T, port int, opts ...verifySetting) *Machine {
	t.Helper()

	optsValue := &verifyOpts{
		pk:             ed25519.GeneratePrivateKey(),
		onboardNodes: []settings.LibraryP2POnboardNode{},
		activateTracing:  false,
	}

	for _, opt := range opts {
		opt(optsValue)
	}

	//
	settings := settings.StandardP2PSettings()
	settings.OriginFolder = t.TempDir()
	settings.AcceptLocation = fmt.Sprintf("REDACTED", port)
	settings.OutsideLocation = fmt.Sprintf("REDACTED", port)

	settings.LibraryP2PSettings.Activated = true
	settings.LibraryP2PSettings.DeactivateAssetAdministrator = true
	settings.LibraryP2PSettings.OnboardNodes = optsValue.onboardNodes

	tracer := log.NewNoopTracer()
	if optsValue.activateTracing {
		tracer = log.VerifyingTracer()
	}

	machine, err := NewMachine(settings, optsValue.pk, tracer)
	require.NoError(t, err)

	return machine
}

func establishOnboardNodes(t *testing.T, ctx context.Context, h *Machine, nodes []OnboardNode) {
	require.NotEmpty(t, nodes, "REDACTED")

	for _, node := range nodes {
		//
		if h.ID().String() == node.AddressDetails.ID.String() {
			continue
		}

		h.tracer.Details("REDACTED", "REDACTED", node.AddressDetails.ID.String())

		err := h.Connect(ctx, node.AddressDetails)
		require.NoError(t, err, "REDACTED", "REDACTED", node.AddressDetails.ID.String())
	}
}

func createVerifyHosts(t *testing.T, countHosts int, opts ...verifySetting) []*Machine {
	ports := utilities.FetchReleasePorts(t, countHosts)

	hosts := make([]*Machine, len(ports))
	for i, port := range ports {
		hosts[i] = createVerifyMachine(t, port, opts...)
	}

	t.Cleanup(func() {
		for _, machine := range hosts {
			machine.Close()
		}
	})

	return hosts
}

func VerifyOnboardNodes(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		//
		//
		pk1 := ed25519.GeneratePrivateKey()
		pk2 := ed25519.GeneratePrivateKey()

		publicidUID := func(pk ed25519.PrivateKey) string {
			id, err := UIDFromPrivateKey(pk)
			require.NoError(t, err)
			return id.String()
		}

		//
		cfg := settings.StandardP2PSettings()
		cfg.LibraryP2PSettings.OnboardNodes = []settings.LibraryP2POnboardNode{
			{Machine: "REDACTED", ID: publicidUID(pk1), Internal: true, Durable: false, Absolute: true},
			{Machine: "REDACTED", ID: publicidUID(pk2), Internal: false, Durable: true, Absolute: false},
			//
			{Machine: "REDACTED", ID: publicidUID(pk2), Internal: false, Durable: true, Absolute: false},
		}

		//
		onboardNodes, err := OnboardNodesFromSettings(cfg)

		//
		require.NoError(t, err)
		require.Len(t, onboardNodes, 2)

		//
		require.Equal(t, publicidUID(pk1), onboardNodes[0].AddressDetails.ID.String())
		require.Len(t, onboardNodes[0].AddressDetails.Addrs, 1)
		require.True(t, onboardNodes[0].Internal)
		require.False(t, onboardNodes[0].Durable)
		require.True(t, onboardNodes[0].Absolute)

		//
		require.Equal(t, publicidUID(pk2), onboardNodes[1].AddressDetails.ID.String())
		require.Len(t, onboardNodes[1].AddressDetails.Addrs, 1)
		require.False(t, onboardNodes[1].Internal)
		require.True(t, onboardNodes[1].Durable)
		require.False(t, onboardNodes[1].Absolute)
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		cfg := settings.StandardP2PSettings()
		cfg.LibraryP2PSettings.OnboardNodes = []settings.LibraryP2POnboardNode{
			{Machine: "REDACTED", ID: "REDACTED"},
		}

		//
		onboardNodes, err := OnboardNodesFromSettings(cfg)

		//
		require.Error(t, err)
		require.Nil(t, onboardNodes)
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		cfg := settings.StandardP2PSettings()
		cfg.LibraryP2PSettings.OnboardNodes = []settings.LibraryP2POnboardNode{
			{Machine: "REDACTED", ID: "REDACTED"},
		}

		//
		onboardNodes, err := OnboardNodesFromSettings(cfg)

		//
		require.Error(t, err)
		require.Nil(t, onboardNodes)
	})
}

func VerifyIsDNSAddress(t *testing.T) {
	for _, tc := range []struct {
		label   string
		raw    string
		anticipate bool
	}{
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: false,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: false,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: false,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: false,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: false,
		},
		{
			label:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
	} {
		t.Run(tc.label, func(t *testing.T) {
			//
			address, err := ma.NewMultiaddr(tc.raw)
			require.NoError(t, err)

			//
			got := IsDNSAddress(address)

			//
			require.Equal(t, tc.anticipate, got)
		})
	}
}
