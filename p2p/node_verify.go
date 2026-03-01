package p2p

import (
	"fmt"
	golog "log"
	"net"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	consensuslink "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
)

func VerifyNodeFundamental(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	//
	rp := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
	rp.Initiate()
	t.Cleanup(rp.Halt)

	p, err := generateOutgoingNodeAlsoExecuteNegotiation(rp.Location(), cfg, consensuslink.FallbackModuleLinkSettings())
	require.Nil(err)

	err = p.Initiate()
	require.Nil(err)
	t.Cleanup(func() {
		if err := p.Halt(); err != nil {
			t.Error(err)
		}
	})

	assert.True(p.EqualsActive())
	assert.True(p.EqualsOutgoing())
	assert.False(p.EqualsEnduring())
	p.enduring = true
	assert.True(p.EqualsEnduring())
	assert.Equal(rp.Location().CallText(), p.DistantLocation().String())
	assert.Equal(rp.ID(), p.ID())
}

func VerifyNodeTransmit(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	settings := cfg

	//
	rp := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: settings}
	rp.Initiate()
	t.Cleanup(rp.Halt)

	p, err := generateOutgoingNodeAlsoExecuteNegotiation(rp.Location(), settings, consensuslink.FallbackModuleLinkSettings())
	require.Nil(err)

	err = p.Initiate()
	require.Nil(err)

	t.Cleanup(func() {
		if err := p.Halt(); err != nil {
			t.Error(err)
		}
	})

	assert.True(p.AbleTransmit(verifyChnl))
	assert.True(p.Transmit(Wrapper{ConduitUUID: verifyChnl, Signal: &p2p.Signal{}}))
}

func generateOutgoingNodeAlsoExecuteNegotiation(
	location *NetworkLocator,
	settings *settings.Peer2peerSettings,
	moduleSettings consensuslink.ModuleLinkSettings,
) (*node, error) {
	chnlDescriptions := []*consensuslink.ConduitDefinition{
		{ID: verifyChnl, Urgency: 1},
	}
	enginesViaChnl := map[byte]Handler{verifyChnl: FreshVerifyHandler(chnlDescriptions, true)}
	signalKindViaChnlUUID := map[byte]proto.Message{
		verifyChnl: &p2p.Signal{},
	}
	pk := edwards25519.ProducePrivateToken()
	pc, err := verifyOutgoingNodeLink(location, settings, false, pk)
	if err != nil {
		return nil, err
	}
	deadline := 1 * time.Second
	minePeerDetails := verifyPeerDetails(location.ID, "REDACTED")
	nodePeerDetails, err := negotiation(pc.link, deadline, minePeerDetails)
	if err != nil {
		return nil, err
	}

	p := freshNode(pc, moduleSettings, nodePeerDetails, enginesViaChnl, signalKindViaChnlUUID, chnlDescriptions, func(p Node, r any) {}, freshTelemetryTagStash())
	p.AssignTracer(log.VerifyingTracer().Using("REDACTED", location))
	return p, nil
}

func verifyCall(location *NetworkLocator, cfg *settings.Peer2peerSettings) (net.Conn, error) {
	if cfg.VerifyCallMishap {
		return nil, fmt.Errorf("REDACTED")
	}

	link, err := location.CallDeadline(cfg.CallDeadline)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func verifyOutgoingNodeLink(
	location *NetworkLocator,
	settings *settings.Peer2peerSettings,
	enduring bool,
	minePeerPrivateToken security.PrivateToken,
) (nodeLink, error) {
	var pc nodeLink
	link, err := verifyCall(location, settings)
	if err != nil {
		return pc, fmt.Errorf("REDACTED", err)
	}

	pc, err = verifyNodeLink(link, settings, true, enduring, minePeerPrivateToken, location)
	if err != nil {
		if checkfault := link.Close(); checkfault != nil {
			return pc, fmt.Errorf("REDACTED", checkfault.Error(), err)
		}
		return pc, err
	}

	//
	if location.ID != pc.ID() {
		if checkfault := link.Close(); checkfault != nil {
			return pc, fmt.Errorf("REDACTED", checkfault.Error(), err)
		}
		return pc, FaultRouterAuthorizationBreakdown{location, pc.ID()}
	}

	return pc, nil
}

type distantNode struct {
	PrivateToken    security.PrivateToken
	Settings     *settings.Peer2peerSettings
	location       *NetworkLocator
	conduits   octets.HexadecimalOctets
	overhearLocation string
	observer   net.Listener
}

func (rp *distantNode) Location() *NetworkLocator {
	return rp.location
}

func (rp *distantNode) ID() ID {
	return PublicTokenTowardUUID(rp.PrivateToken.PublicToken())
}

func (rp *distantNode) Initiate() {
	if rp.overhearLocation == "REDACTED" {
		rp.overhearLocation = "REDACTED"
	}

	l, e := net.Listen("REDACTED", rp.overhearLocation) //
	if e != nil {
		golog.Fatalf("REDACTED", e)
	}
	rp.observer = l
	rp.location = FreshNetworkLocator(PublicTokenTowardUUID(rp.PrivateToken.PublicToken()), l.Addr())
	if rp.conduits == nil {
		rp.conduits = []byte{verifyChnl}
	}
	go rp.embrace()
}

func (rp *distantNode) Halt() {
	rp.observer.Close()
}

func (rp *distantNode) Call(location *NetworkLocator) (net.Conn, error) {
	link, err := location.CallDeadline(1 * time.Second)
	if err != nil {
		return nil, err
	}
	pc, err := verifyIncomingNodeLink(link, rp.Settings, rp.PrivateToken)
	if err != nil {
		return nil, err
	}
	_, err = negotiation(pc.link, time.Second, rp.peerDetails())
	if err != nil {
		return nil, err
	}
	return link, err
}

func (rp *distantNode) embrace() {
	links := []net.Conn{}

	for {
		link, err := rp.observer.Accept()
		if err != nil {
			golog.Printf("REDACTED", err)
			for _, link := range links {
				_ = link.Close()
			}
			return
		}

		pc, err := verifyIncomingNodeLink(link, rp.Settings, rp.PrivateToken)
		if err != nil {
			golog.Fatalf("REDACTED", err)
		}

		_, err = negotiation(pc.link, time.Second, rp.peerDetails())
		if err != nil {
			golog.Fatalf("REDACTED", err)
		}

		links = append(links, link)
	}
}

func (rp *distantNode) peerDetails() PeerDetails {
	return FallbackPeerDetails{
		SchemeEdition: fallbackSchemeEdition,
		FallbackPeerUUID:   rp.Location().ID,
		OverhearLocation:      rp.observer.Addr().String(),
		Fabric:         "REDACTED",
		Edition:         "REDACTED",
		Conduits:        rp.conduits,
		Pseudonym:         "REDACTED",
	}
}
