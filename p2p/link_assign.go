package p2p

import (
	"net"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
type LinkAssign interface {
	Has(net.Conn) bool
	OwnsINET(net.IP) bool
	Set(net.Conn, []net.IP)
	Discard(net.Conn)
	DiscardLocation(net.Addr)
}

type linkAssignElement struct {
	link net.Conn
	ips  []net.IP
}

type linkAssign struct {
	commitchronize.ReadwriteExclusion

	links map[string]linkAssignElement
}

//
func FreshLinkAssign() LinkAssign {
	return &linkAssign{
		links: map[string]linkAssignElement{},
	}
}

func (cs *linkAssign) Has(c net.Conn) bool {
	cs.RLock()
	defer cs.RUnlock()

	_, ok := cs.links[c.RemoteAddr().String()]

	return ok
}

func (cs *linkAssign) OwnsINET(ip net.IP) bool {
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

func (cs *linkAssign) Discard(c net.Conn) {
	cs.Lock()
	defer cs.Unlock()

	delete(cs.links, c.RemoteAddr().String())
}

func (cs *linkAssign) DiscardLocation(location net.Addr) {
	cs.Lock()
	defer cs.Unlock()

	delete(cs.links, location.String())
}

func (cs *linkAssign) Set(c net.Conn, ips []net.IP) {
	cs.Lock()
	defer cs.Unlock()

	cs.links[c.RemoteAddr().String()] = linkAssignElement{
		link: c,
		ips:  ips,
	}
}
