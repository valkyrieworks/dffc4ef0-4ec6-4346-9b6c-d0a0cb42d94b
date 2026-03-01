package netp2p

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/toolkits"
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
		conduitSample = byte(0xaa)
		conduitDivider = byte(0xbb)
		schemaSample   = SchemeUUID(conduitSample)
		schemaDivider   = SchemeUUID(conduitDivider)
	)

	//
	channels := toolkits.ObtainReleaseChannels(t, 2)

	//
	machine1 := createVerifyMachine(t, channels[0], usingJournaling())
	machine2 := createVerifyMachine(t, channels[1], usingJournaling(), usingOnboardNodes([]settings.LibraryPeer2peerInitiateNode{
		{
			//
			Machine: fmt.Sprintf("REDACTED", channels[0]),
			ID:   machine1.ID().String(),
		},
	}))

	relateOnboardNodes(t, ctx, machine2, machine2.InitiateNodes())

	t.Logf("REDACTED", machine1.LocationDetails())
	t.Logf("REDACTED", machine2.LocationDetails())

	t.Cleanup(func() {
		machine2.Close()
		machine1.Close()
	})

	//
	type wrapper struct {
		scheme protocol.ID
		originator   corepeer.ID
		acceptor corepeer.ID
		signal  string
	}

	wrappers := []wrapper{}
	mu := sync.Mutex{}

	//
	processor := func(influx network.Stream) {
		var (
			link     = influx.Conn()
			acceptor = link.LocalPeer()
			originator   = link.RemotePeer()
		)

		if link.ConnState().Transport != CarrierQuicprotocol {
			t.Fatalf("REDACTED", link.ConnState().Transport)
			return
		}

		content, err := InfluxFetchShutdown(influx)
		if err != nil {
			t.Fatalf("REDACTED", originator, err)
			return
		}

		msg := &kinds.Solicit{}
		require.NoError(t, msg.Decode(content))
		require.NotNil(t, msg.ObtainReverberate())

		e := wrapper{
			scheme: influx.Protocol(),
			originator:   originator,
			acceptor: acceptor,
			signal:  msg.ObtainReverberate().ObtainArtifact(),
		}

		recordArtifact := e.signal
		if len(recordArtifact) > 64 {
			recordArtifact = recordArtifact[:64] + "REDACTED"
		}

		t.Logf(
			"REDACTED",
			e.originator.String(),
			e.acceptor.String(),
			e.scheme,
			recordArtifact,
		)

		mu.Lock()
		defer mu.Unlock()

		wrappers = append(wrappers, e)
	}

	machine1.SetStreamHandler(schemaSample, processor)
	machine1.SetStreamHandler(schemaDivider, processor)

	machine2.SetStreamHandler(schemaSample, processor)
	machine2.SetStreamHandler(schemaDivider, processor)

	//
	machine1peer2, err := FreshNode(machine1, machine2.LocationDetails(), p2p.NooperationTelemetry(), false, false, false)
	require.NoError(t, err, "REDACTED")
	require.NoError(t, machine1peer2.Initiate(), "REDACTED")

	machine2peer1, err := FreshNode(machine2, machine1.LocationDetails(), p2p.NooperationTelemetry(), false, false, false)
	require.NoError(t, err, "REDACTED")
	require.NoError(t, machine2peer1.Initiate(), "REDACTED")

	t.Logf("REDACTED", machine1peer2.ID())
	t.Logf("REDACTED", machine2peer1.ID())

	//
	//
	extendedTxt := strings.Repeat("REDACTED", 300*1024)

	//
	dispatch1 := machine1peer2.Transmit(p2p.Wrapper{
		ConduitUUID: conduitSample,
		Signal:   kinds.TowardSolicitReverberate("REDACTED"),
	})

	dispatch2 := machine2peer1.Transmit(p2p.Wrapper{
		ConduitUUID: conduitDivider,
		Signal:   kinds.TowardSolicitReverberate("REDACTED"),
	})

	dispatch3 := machine1peer2.AttemptTransmit(p2p.Wrapper{
		ConduitUUID: conduitDivider,
		Signal:   kinds.TowardSolicitReverberate(extendedTxt),
	})

	//
	//
	require.True(t, dispatch1, "REDACTED")
	require.True(t, dispatch2, "REDACTED")
	require.True(t, dispatch3, "REDACTED")

	//
	pause := func() bool {
		mu.Lock()
		defer mu.Unlock()
		return len(wrappers) == 3
	}

	require.Eventually(t, pause, 500*time.Millisecond, 50*time.Millisecond)

	//
	anticipatedWrappers := []wrapper{
		{
			scheme: schemaSample,
			originator:   machine1.ID(),
			acceptor: machine2.ID(),
			signal:  "REDACTED",
		},
		{
			scheme: schemaDivider,
			originator:   machine2.ID(),
			acceptor: machine1.ID(),
			signal:  "REDACTED",
		},
		{
			scheme: schemaDivider,
			originator:   machine1.ID(),
			acceptor: machine2.ID(),
			signal:  extendedTxt,
		},
	}

	require.ElementsMatch(t, anticipatedWrappers, wrappers)

	t.Run("REDACTED", func(t *testing.T) {
		//
		ctx, abort := context.WithTimeout(context.Background(), 5*time.Second)
		defer abort()

		//
		rtt, err := machine1.Ping(ctx, machine2.LocationDetails())

		//
		require.NoError(t, err)
		require.NotZero(t, rtt)

		t.Logf("REDACTED", rtt.String())
	})
}

type verifyOptions struct {
	pk             edwards25519.PrivateToken
	initiateNodes []settings.LibraryPeer2peerInitiateNode
	activateJournaling  bool
}

type verifySelection func(*verifyOptions)

func usingJournaling() verifySelection {
	return func(choices *verifyOptions) { choices.activateJournaling = true }
}

func usingOnboardNodes(initiateNodes []settings.LibraryPeer2peerInitiateNode) verifySelection {
	return func(choices *verifyOptions) { choices.initiateNodes = initiateNodes }
}

func usingSecludedToken(pk edwards25519.PrivateToken) verifySelection {
	return func(choices *verifyOptions) { choices.pk = pk }
}

func createVerifyMachine(t *testing.T, channel int, choices ...verifySelection) *Machine {
	t.Helper()

	optionsItem := &verifyOptions{
		pk:             edwards25519.ProducePrivateToken(),
		initiateNodes: []settings.LibraryPeer2peerInitiateNode{},
		activateJournaling:  false,
	}

	for _, opt := range choices {
		opt(optionsItem)
	}

	//
	settings := settings.FallbackPeer2peerSettings()
	settings.OriginPath = t.TempDir()
	settings.OverhearLocation = fmt.Sprintf("REDACTED", channel)
	settings.OutsideLocation = fmt.Sprintf("REDACTED", channel)

	settings.LibraryPeer2peerSettings.Activated = true
	settings.LibraryPeer2peerSettings.DeactivateAssetAdministrator = true
	settings.LibraryPeer2peerSettings.InitiateNodes = optionsItem.initiateNodes

	tracer := log.FreshNooperationTracer()
	if optionsItem.activateJournaling {
		tracer = log.VerifyingTracer()
	}

	machine, err := FreshMachine(settings, optionsItem.pk, tracer)
	require.NoError(t, err)

	return machine
}

func relateOnboardNodes(t *testing.T, ctx context.Context, h *Machine, nodes []OnboardNode) {
	require.NotEmpty(t, nodes, "REDACTED")

	for _, node := range nodes {
		//
		if h.ID().String() == node.LocationDetails.ID.String() {
			continue
		}

		h.tracer.Details("REDACTED", "REDACTED", node.LocationDetails.ID.String())

		err := h.Connect(ctx, node.LocationDetails)
		require.NoError(t, err, "REDACTED", "REDACTED", node.LocationDetails.ID.String())
	}
}

func createVerifyMachines(t *testing.T, countMachines int, choices ...verifySelection) []*Machine {
	channels := toolkits.ObtainReleaseChannels(t, countMachines)

	machines := make([]*Machine, len(channels))
	for i, channel := range channels {
		machines[i] = createVerifyMachine(t, channel, choices...)
	}

	t.Cleanup(func() {
		for _, machine := range machines {
			machine.Close()
		}
	})

	return machines
}

func VerifyOnboardNodes(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		//
		//
		pk1 := edwards25519.ProducePrivateToken()
		pk2 := edwards25519.ProducePrivateToken()

		keyUUID := func(pk edwards25519.PrivateToken) string {
			id, err := UUIDOriginatingSecludedToken(pk)
			require.NoError(t, err)
			return id.String()
		}

		//
		cfg := settings.FallbackPeer2peerSettings()
		cfg.LibraryPeer2peerSettings.InitiateNodes = []settings.LibraryPeer2peerInitiateNode{
			{Machine: "REDACTED", ID: keyUUID(pk1), Secluded: true, Enduring: false, Absolute: true},
			{Machine: "REDACTED", ID: keyUUID(pk2), Secluded: false, Enduring: true, Absolute: false},
			//
			{Machine: "REDACTED", ID: keyUUID(pk2), Secluded: false, Enduring: true, Absolute: false},
		}

		//
		initiateNodes, err := OnboardNodesOriginatingSettings(cfg)

		//
		require.NoError(t, err)
		require.Len(t, initiateNodes, 2)

		//
		require.Equal(t, keyUUID(pk1), initiateNodes[0].LocationDetails.ID.String())
		require.Len(t, initiateNodes[0].LocationDetails.Addrs, 1)
		require.True(t, initiateNodes[0].Secluded)
		require.False(t, initiateNodes[0].Enduring)
		require.True(t, initiateNodes[0].Absolute)

		//
		require.Equal(t, keyUUID(pk2), initiateNodes[1].LocationDetails.ID.String())
		require.Len(t, initiateNodes[1].LocationDetails.Addrs, 1)
		require.False(t, initiateNodes[1].Secluded)
		require.True(t, initiateNodes[1].Enduring)
		require.False(t, initiateNodes[1].Absolute)
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		cfg := settings.FallbackPeer2peerSettings()
		cfg.LibraryPeer2peerSettings.InitiateNodes = []settings.LibraryPeer2peerInitiateNode{
			{Machine: "REDACTED", ID: "REDACTED"},
		}

		//
		initiateNodes, err := OnboardNodesOriginatingSettings(cfg)

		//
		require.Error(t, err)
		require.Nil(t, initiateNodes)
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		cfg := settings.FallbackPeer2peerSettings()
		cfg.LibraryPeer2peerSettings.InitiateNodes = []settings.LibraryPeer2peerInitiateNode{
			{Machine: "REDACTED", ID: "REDACTED"},
		}

		//
		initiateNodes, err := OnboardNodesOriginatingSettings(cfg)

		//
		require.Error(t, err)
		require.Nil(t, initiateNodes)
	})
}

func VerifyEqualsDomainLocation(t *testing.T) {
	for _, tc := range []struct {
		alias   string
		raw    string
		anticipate bool
	}{
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: false,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: false,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: false,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: false,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: false,
		},
		{
			alias:   "REDACTED",
			raw:    "REDACTED",
			anticipate: true,
		},
	} {
		t.Run(tc.alias, func(t *testing.T) {
			//
			location, err := ma.NewMultiaddr(tc.raw)
			require.NoError(t, err)

			//
			got := EqualsDomainLocation(location)

			//
			require.Equal(t, tc.anticipate, got)
		})
	}
}
