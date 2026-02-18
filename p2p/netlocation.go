//
//
//

package p2p

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
)

//
const EmptyNetLocation = "REDACTED"

//
//
type NetLocation struct {
	ID   ID     `json:"id"`
	IP   net.IP `json:"ip"`
	Port uint16 `json:"port"`
}

//
//
func UIDLocationString(id ID, protocolMachinePort string) string {
	machinePort := deleteProtocolIfSpecified(protocolMachinePort)
	return fmt.Sprintf("REDACTED", id, machinePort)
}

//
//
//
//
//
func NewNetLocation(id ID, address net.Addr) *NetLocation {
	tcpAddress, ok := address.(*net.TCPAddr)
	if !ok {
		if flag.Lookup("REDACTED") == nil { //
			panic(fmt.Sprintf("REDACTED", address))
		}
		//
		netAddress := NewNetLocationIPPort(net.IP("REDACTED"), 0)
		netAddress.ID = id
		return netAddress
	}

	if err := certifyUID(id); err != nil {
		panic(fmt.Sprintf("REDACTED", id, err, address))
	}

	ip := tcpAddress.IP
	port := uint16(tcpAddress.Port)
	na := NewNetLocationIPPort(ip, port)
	na.ID = id
	return na
}

//
//
//
//
func NewNetLocationString(address string) (*NetLocation, error) {
	addressLackingProtocol := deleteProtocolIfSpecified(address)
	spl := strings.Split(addressLackingProtocol, "REDACTED")
	if len(spl) != 2 {
		return nil, ErrNetLocationNoUID{address}
	}

	//
	if err := certifyUID(ID(spl[0])); err != nil {
		return nil, ErrNetLocationCorrupt{addressLackingProtocol, err}
	}
	var id ID
	id, addressLackingProtocol = ID(spl[0]), spl[1]

	//
	machine, portStr, err := net.SplitHostPort(addressLackingProtocol)
	if err != nil {
		return nil, ErrNetLocationCorrupt{addressLackingProtocol, err}
	}
	if len(machine) == 0 {
		return nil, ErrNetLocationCorrupt{
			addressLackingProtocol,
			errors.New("REDACTED"),
		}
	}

	ip := net.ParseIP(machine)
	if ip == nil {
		ips, err := net.LookupIP(machine)
		if err != nil {
			return nil, ErrNetLocationSearch{machine, err}
		}
		ip = ips[0]
	}

	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return nil, ErrNetLocationCorrupt{portStr, err}
	}

	na := NewNetLocationIPPort(ip, uint16(port))
	na.ID = id
	return na, nil
}

//
//
func NewNetLocationStrings(locations []string) ([]*NetLocation, []error) {
	netLocations := make([]*NetLocation, 0)
	faults := make([]error, 0)
	for _, address := range locations {
		netAddress, err := NewNetLocationString(address)
		if err != nil {
			faults = append(faults, err)
		} else {
			netLocations = append(netLocations, netAddress)
		}
	}
	return netLocations, faults
}

//
//
func NewNetLocationIPPort(ip net.IP, port uint16) *NetLocation {
	return &NetLocation{
		IP:   ip,
		Port: port,
	}
}

//
func NetLocationFromSchema(pb tmp2p.NetLocation) (*NetLocation, error) {
	ip := net.ParseIP(pb.IP)
	if ip == nil {
		return nil, fmt.Errorf("REDACTED", pb.IP)
	}
	if pb.Port >= 1<<16 {
		return nil, fmt.Errorf("REDACTED", pb.Port)
	}
	return &NetLocation{
		ID:   ID(pb.ID),
		IP:   ip,
		Port: uint16(pb.Port),
	}, nil
}

//
func NetAddressesFromSchema(pbs []tmp2p.NetLocation) ([]*NetLocation, error) {
	nas := make([]*NetLocation, 0, len(pbs))
	for _, pb := range pbs {
		na, err := NetLocationFromSchema(pb)
		if err != nil {
			return nil, err
		}
		nas = append(nas, na)
	}
	return nas, nil
}

//
func NetAddressesToSchema(nas []*NetLocation) []tmp2p.NetLocation {
	pbs := make([]tmp2p.NetLocation, 0, len(nas))
	for _, na := range nas {
		if na != nil {
			pbs = append(pbs, na.ToSchema())
		}
	}
	return pbs
}

//
func (na *NetLocation) ToSchema() tmp2p.NetLocation {
	return tmp2p.NetLocation{
		ID:   string(na.ID),
		IP:   na.IP.String(),
		Port: uint32(na.Port),
	}
}

//
//
func (na *NetLocation) Matches(another any) bool {
	if o, ok := another.(*NetLocation); ok {
		return na.String() == o.String()
	}
	return false
}

//
func (na *NetLocation) Identical(another any) bool {
	if o, ok := another.(*NetLocation); ok {
		if na.CallString() == o.CallString() {
			return true
		}
		if na.ID != "REDACTED" && na.ID == o.ID {
			return true
		}
	}
	return false
}

//
func (na *NetLocation) String() string {
	if na == nil {
		return EmptyNetLocation
	}

	addressStr := na.CallString()
	if na.ID != "REDACTED" {
		addressStr = UIDLocationString(na.ID, addressStr)
	}

	return addressStr
}

func (na *NetLocation) CallString() string {
	if na == nil {
		return "REDACTED"
	}
	return net.JoinHostPort(
		na.IP.String(),
		strconv.FormatUint(uint64(na.Port), 10),
	)
}

//
func (na *NetLocation) Call() (net.Conn, error) {
	link, err := net.Dial("REDACTED", na.CallString())
	if err != nil {
		return nil, err
	}
	return link, nil
}

//
func (na *NetLocation) CallDeadline(deadline time.Duration) (net.Conn, error) {
	link, err := net.DialTimeout("REDACTED", na.CallString(), deadline)
	if err != nil {
		return nil, err
	}
	return link, nil
}

//
func (na *NetLocation) Forwardable() bool {
	if err := na.Sound(); err != nil {
		return false
	}
	//
	return !na.Rfc1918() && !na.Rfc3927() && !na.Rfc4862() && !na.Rfc4193() && !na.Rfc4843() && !na.Native()
}

//
//
func (na *NetLocation) Sound() error {
	if err := certifyUID(na.ID); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if na.IP == nil {
		return errors.New("REDACTED")
	}
	if na.IP.IsUnspecified() || na.Rfc3849() || na.IP.Equal(net.IPv4bcast) {
		return errors.New("REDACTED")
	}
	return nil
}

//
//
func (na *NetLocation) HasUID() bool {
	return string(na.ID) != "REDACTED"
}

//
func (na *NetLocation) Native() bool {
	return na.IP.IsLoopback() || zero4.Contains(na.IP)
}

//
func (na *NetLocation) AccessibilityTo(o *NetLocation) int {
	const (
		Inaccessible = 0
		Standard     = iota
		Teredo
		Ipv6weak
		Ipv4
		Ipv6strong
	)
	switch {
	case !na.Forwardable():
		return Inaccessible
	case na.Rfc4380():
		switch {
		case !o.Forwardable():
			return Standard
		case o.Rfc4380():
			return Teredo
		case o.IP.To4() != nil:
			return Ipv4
		default: //
			return Ipv6weak
		}
	case na.IP.To4() != nil:
		if o.Forwardable() && o.IP.To4() != nil {
			return Ipv4
		}
		return Standard
	default: //
		var channeled bool
		//
		if o.Rfc3964() || o.Rfc6052() || o.Rfc6145() {
			channeled = true
		}
		switch {
		case !o.Forwardable():
			return Standard
		case o.Rfc4380():
			return Teredo
		case o.IP.To4() != nil:
			return Ipv4
		case channeled:
			//
			return Ipv6weak
		}
		return Ipv6strong
	}
}

//
//
//
//
//
//
//
//
//
//
var (
	rfc1918_10  = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(8, 32)}
	rfc1918_192 = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(16, 32)}
	rfc1918_172 = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(12, 32)}
	rfc3849     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(32, 128)}
	rfc3927     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(16, 32)}
	rfc3964     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(16, 128)}
	rfc4193     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(7, 128)}
	rfc4380     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(32, 128)}
	rfc4843     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(28, 128)}
	rfc4862     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(64, 128)}
	rfc6052     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(96, 128)}
	rfc6145     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(96, 128)}
	zero4       = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(8, 32)}
)

//
//
//
//
//
//
//
//
//
//
//
var onionCatNet = ipNet("REDACTED", 48, 128)

//
//
//
func ipNet(ip string, ones, bits int) net.IPNet {
	return net.IPNet{IP: net.ParseIP(ip), Mask: net.CIDRMask(ones, bits)}
}

func (na *NetLocation) Rfc1918() bool {
	return rfc1918_10.Contains(na.IP) ||
		rfc1918_192.Contains(na.IP) ||
		rfc1918_172.Contains(na.IP)
}
func (na *NetLocation) Rfc3849() bool     { return rfc3849.Contains(na.IP) }
func (na *NetLocation) Rfc3927() bool     { return rfc3927.Contains(na.IP) }
func (na *NetLocation) Rfc3964() bool     { return rfc3964.Contains(na.IP) }
func (na *NetLocation) Rfc4193() bool     { return rfc4193.Contains(na.IP) }
func (na *NetLocation) Rfc4380() bool     { return rfc4380.Contains(na.IP) }
func (na *NetLocation) Rfc4843() bool     { return rfc4843.Contains(na.IP) }
func (na *NetLocation) Rfc4862() bool     { return rfc4862.Contains(na.IP) }
func (na *NetLocation) Rfc6052() bool     { return rfc6052.Contains(na.IP) }
func (na *NetLocation) Rfc6145() bool     { return rfc6145.Contains(na.IP) }
func (na *NetLocation) OnionCatTor() bool { return onionCatNet.Contains(na.IP) }

func deleteProtocolIfSpecified(address string) string {
	if strings.Contains(address, "REDACTED") {
		return strings.Split(address, "REDACTED")[1]
	}
	return address
}

func certifyUID(id ID) error {
	if len(id) == 0 {
		return errors.New("REDACTED")
	}
	uidOctets, err := hex.DecodeString(string(id))
	if err != nil {
		return err
	}
	if len(uidOctets) != UIDOctetExtent {
		return fmt.Errorf("REDACTED", len(uidOctets), UIDOctetExtent)
	}
	return nil
}
