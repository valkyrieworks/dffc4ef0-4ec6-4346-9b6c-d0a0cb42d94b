package p2p

import (
	"net"

	engineconnect "github.com/valkyrieworks/utils/align"
)

//
type LinkCollection interface {
	Has(net.Conn) bool
	HasIP(net.IP) bool
	Set(net.Conn, []net.IP)
	Delete(net.Conn)
	DeleteAddress(net.Addr)
}

type linkCollectionItem struct {
	link net.Conn
	ips  []net.IP
}

type linkCollection struct {
	engineconnect.ReadwriteLock

	links map[string]linkCollectionItem
}

//
func NewLinkCollection() LinkCollection {
	return &linkCollection{
		links: map[string]linkCollectionItem{},
	}
}

func (cs *linkCollection) Has(c net.Conn) bool {
	cs.RLock()
	defer cs.RUnlock()

	_, ok := cs.links[c.RemoteAddr().String()]

	return ok
}

func (cs *linkCollection) HasIP(ip net.IP) bool {
	cs.RLock()
	defer cs.RUnlock()

	for _, c := range cs.links {
		for _, recognized := range c.ips {
			if recognized.Equal(ip) {
				return true
			}
		}
	}

	return false
}

func (cs *linkCollection) Delete(c net.Conn) {
	cs.Lock()
	defer cs.Unlock()

	delete(cs.links, c.RemoteAddr().String())
}

func (cs *linkCollection) DeleteAddress(address net.Addr) {
	cs.Lock()
	defer cs.Unlock()

	delete(cs.links, address.String())
}

func (cs *linkCollection) Set(c net.Conn, ips []net.IP) {
	cs.Lock()
	defer cs.Unlock()

	cs.links[c.RemoteAddr().String()] = linkCollectionItem{
		link: c,
		ips:  ips,
	}
}
