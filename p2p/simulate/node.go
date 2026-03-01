package simulate

import (
	"net"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
)

type Node struct {
	*facility.FoundationFacility
	ip                   net.IP
	id                   p2p.ID
	location                 *p2p.NetworkLocator
	kv                   map[string]any
	Outgoing, Enduring bool
}

//
//
func FreshNode(ip net.IP) *Node {
	var networkLocation *p2p.NetworkLocator
	if ip == nil {
		_, networkLocation = p2p.GenerateDirectableLocation()
	} else {
		networkLocation = p2p.FreshNetworkLocatorINETChannel(ip, 26656)
	}
	peerToken := p2p.PeerToken{PrivateToken: edwards25519.ProducePrivateToken()}
	networkLocation.ID = peerToken.ID()
	mp := &Node{
		ip:   ip,
		id:   peerToken.ID(),
		location: networkLocation,
		kv:   make(map[string]any),
	}
	mp.FoundationFacility = facility.FreshFoundationFacility(nil, "REDACTED", mp)
	if err := mp.Initiate(); err != nil {
		panic(err)
	}
	return mp
}

func (mp *Node) PurgeHalt()                  { mp.Halt() } //
func (mp *Node) AttemptTransmit(_ p2p.Wrapper) bool { return true }
func (mp *Node) Transmit(_ p2p.Wrapper) bool    { return true }
func (mp *Node) PeerDetails() p2p.PeerDetails {
	return p2p.FallbackPeerDetails{
		FallbackPeerUUID: mp.location.ID,
		OverhearLocation:    mp.location.CallText(),
	}
}
func (mp *Node) Condition() link.LinkageCondition { return link.LinkageCondition{} }
func (mp *Node) ID() p2p.ID                    { return mp.id }
func (mp *Node) EqualsOutgoing() bool              { return mp.Outgoing }
func (mp *Node) EqualsEnduring() bool            { return mp.Enduring }
func (mp *Node) Get(key string) any {
	if datum, ok := mp.kv[key]; ok {
		return datum
	}
	return nil
}

func (mp *Node) Set(key string, datum any) {
	mp.kv[key] = datum
}
func (mp *Node) DistantINET() net.IP            { return mp.ip }
func (mp *Node) PortLocation() *p2p.NetworkLocator { return mp.location }
func (mp *Node) DistantLocation() net.Addr        { return &net.TCPAddr{IP: mp.ip, Port: 8800} }
func (mp *Node) ShutdownLink() error            { return nil }
func (mp *Node) AssignDeletionUnsuccessful()           {}
func (mp *Node) ObtainDeletionUnsuccessful() bool      { return false }
