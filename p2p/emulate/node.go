package emulate

import (
	"net"

	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/p2p/link"
)

type Node struct {
	*daemon.RootDaemon
	ip                   net.IP
	id                   p2p.ID
	address                 *p2p.NetLocation
	kv                   map[string]any
	Outgoing, Durable bool
}

//
//
func NewNode(ip net.IP) *Node {
	var netAddress *p2p.NetLocation
	if ip == nil {
		_, netAddress = p2p.InstantiateForwardableAddress()
	} else {
		netAddress = p2p.NewNetLocationIPPort(ip, 26656)
	}
	memberKey := p2p.MemberKey{PrivateKey: ed25519.GeneratePrivateKey()}
	netAddress.ID = memberKey.ID()
	mp := &Node{
		ip:   ip,
		id:   memberKey.ID(),
		address: netAddress,
		kv:   make(map[string]any),
	}
	mp.RootDaemon = daemon.NewRootDaemon(nil, "REDACTED", mp)
	if err := mp.Begin(); err != nil {
		panic(err)
	}
	return mp
}

func (mp *Node) PurgeHalt()                  { mp.Halt() } //
func (mp *Node) AttemptTransmit(_ p2p.Packet) bool { return true }
func (mp *Node) Transmit(_ p2p.Packet) bool    { return true }
func (mp *Node) MemberDetails() p2p.MemberDetails {
	return p2p.StandardMemberDetails{
		StandardMemberUID: mp.address.ID,
		ObserveAddress:    mp.address.CallString(),
	}
}
func (mp *Node) Status() link.LinkageState { return link.LinkageState{} }
func (mp *Node) ID() p2p.ID                    { return mp.id }
func (mp *Node) IsOutgoing() bool              { return mp.Outgoing }
func (mp *Node) IsDurable() bool            { return mp.Durable }
func (mp *Node) Get(key string) any {
	if item, ok := mp.kv[key]; ok {
		return item
	}
	return nil
}

func (mp *Node) Set(key string, item any) {
	mp.kv[key] = item
}
func (mp *Node) DistantIP() net.IP            { return mp.ip }
func (mp *Node) SocketAddress() *p2p.NetLocation { return mp.address }
func (mp *Node) DistantAddress() net.Addr        { return &net.TCPAddr{IP: mp.ip, Port: 8800} }
func (mp *Node) EndLink() error            { return nil }
func (mp *Node) CollectionDeletionErrored()           {}
func (mp *Node) FetchDeletionErrored() bool      { return false }
