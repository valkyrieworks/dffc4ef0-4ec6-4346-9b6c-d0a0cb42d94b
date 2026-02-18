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

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/schema/consensuscore/p2p"

	"github.com/valkyrieworks/settings"
	cmtconn "github.com/valkyrieworks/p2p/link"
)

func VerifyNodeSimple(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	//
	rp := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
	rp.Begin()
	t.Cleanup(rp.Halt)

	p, err := instantiateOutgoingNodeAndExecuteGreeting(rp.Address(), cfg, cmtconn.StandardMLinkSettings())
	require.Nil(err)

	err = p.Begin()
	require.Nil(err)
	t.Cleanup(func() {
		if err := p.Halt(); err != nil {
			t.Error(err)
		}
	})

	assert.True(p.IsActive())
	assert.True(p.IsOutgoing())
	assert.False(p.IsDurable())
	p.durable = true
	assert.True(p.IsDurable())
	assert.Equal(rp.Address().CallString(), p.DistantAddress().String())
	assert.Equal(rp.ID(), p.ID())
}

func VerifyNodeTransmit(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	settings := cfg

	//
	rp := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: settings}
	rp.Begin()
	t.Cleanup(rp.Halt)

	p, err := instantiateOutgoingNodeAndExecuteGreeting(rp.Address(), settings, cmtconn.StandardMLinkSettings())
	require.Nil(err)

	err = p.Begin()
	require.Nil(err)

	t.Cleanup(func() {
		if err := p.Halt(); err != nil {
			t.Error(err)
		}
	})

	assert.True(p.MayTransmit(verifyChan))
	assert.True(p.Transmit(Packet{StreamUID: verifyChan, Signal: &p2p.Signal{}}))
}

func instantiateOutgoingNodeAndExecuteGreeting(
	address *NetLocation,
	settings *settings.P2PSettings,
	mSettings cmtconn.MLinkSettings,
) (*node, error) {
	chanTraits := []*cmtconn.StreamDefinition{
		{ID: verifyChan, Urgency: 1},
	}
	handlersByChan := map[byte]Handler{verifyChan: NewVerifyHandler(chanTraits, true)}
	messageKindByChanUID := map[byte]proto.Message{
		verifyChan: &p2p.Signal{},
	}
	pk := ed25519.GeneratePrivateKey()
	pc, err := verifyOutgoingNodeLink(address, settings, false, pk)
	if err != nil {
		return nil, err
	}
	deadline := 1 * time.Second
	ourMemberDetails := verifyMemberDetails(address.ID, "REDACTED")
	nodeMemberDetails, err := greeting(pc.link, deadline, ourMemberDetails)
	if err != nil {
		return nil, err
	}

	p := newNode(pc, mSettings, nodeMemberDetails, handlersByChan, messageKindByChanUID, chanTraits, func(p Node, r any) {}, newStatsTagRepository())
	p.AssignTracer(log.VerifyingTracer().With("REDACTED", address))
	return p, nil
}

func verifyCall(address *NetLocation, cfg *settings.P2PSettings) (net.Conn, error) {
	if cfg.VerifyCallAbort {
		return nil, fmt.Errorf("REDACTED")
	}

	link, err := address.CallDeadline(cfg.CallDeadline)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func verifyOutgoingNodeLink(
	address *NetLocation,
	settings *settings.P2PSettings,
	durable bool,
	ourMemberPrivateKey vault.PrivateKey,
) (nodeLink, error) {
	var pc nodeLink
	link, err := verifyCall(address, settings)
	if err != nil {
		return pc, fmt.Errorf("REDACTED", err)
	}

	pc, err = verifyNodeLink(link, settings, true, durable, ourMemberPrivateKey, address)
	if err != nil {
		if cerr := link.Close(); cerr != nil {
			return pc, fmt.Errorf("REDACTED", cerr.Error(), err)
		}
		return pc, err
	}

	//
	if address.ID != pc.ID() {
		if cerr := link.Close(); cerr != nil {
			return pc, fmt.Errorf("REDACTED", cerr.Error(), err)
		}
		return pc, ErrRouterAuthorizationBreakdown{address, pc.ID()}
	}

	return pc, nil
}

type distantNode struct {
	PrivateKey    vault.PrivateKey
	Settings     *settings.P2PSettings
	address       *NetLocation
	streams   octets.HexOctets
	acceptAddress string
	observer   net.Listener
}

func (rp *distantNode) Address() *NetLocation {
	return rp.address
}

func (rp *distantNode) ID() ID {
	return PublicKeyToUID(rp.PrivateKey.PublicKey())
}

func (rp *distantNode) Begin() {
	if rp.acceptAddress == "REDACTED" {
		rp.acceptAddress = "REDACTED"
	}

	l, e := net.Listen("REDACTED", rp.acceptAddress) //
	if e != nil {
		golog.Fatalf("REDACTED", e)
	}
	rp.observer = l
	rp.address = NewNetLocation(PublicKeyToUID(rp.PrivateKey.PublicKey()), l.Addr())
	if rp.streams == nil {
		rp.streams = []byte{verifyChan}
	}
	go rp.allow()
}

func (rp *distantNode) Halt() {
	rp.observer.Close()
}

func (rp *distantNode) Call(address *NetLocation) (net.Conn, error) {
	link, err := address.CallDeadline(1 * time.Second)
	if err != nil {
		return nil, err
	}
	pc, err := verifyIncomingNodeLink(link, rp.Settings, rp.PrivateKey)
	if err != nil {
		return nil, err
	}
	_, err = greeting(pc.link, time.Second, rp.memberDetails())
	if err != nil {
		return nil, err
	}
	return link, err
}

func (rp *distantNode) allow() {
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

		pc, err := verifyIncomingNodeLink(link, rp.Settings, rp.PrivateKey)
		if err != nil {
			golog.Fatalf("REDACTED", err)
		}

		_, err = greeting(pc.link, time.Second, rp.memberDetails())
		if err != nil {
			golog.Fatalf("REDACTED", err)
		}

		links = append(links, link)
	}
}

func (rp *distantNode) memberDetails() MemberDetails {
	return StandardMemberDetails{
		ProtocolRelease: standardProtocolRelease,
		StandardMemberUID:   rp.Address().ID,
		ObserveAddress:      rp.observer.Addr().String(),
		Fabric:         "REDACTED",
		Release:         "REDACTED",
		Streams:        rp.streams,
		Moniker:         "REDACTED",
	}
}
