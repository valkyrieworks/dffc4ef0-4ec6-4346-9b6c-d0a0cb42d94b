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

	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
)

//
const BlankNetworkLocator = "REDACTED"

//
//
type NetworkLocator struct {
	ID   ID     `json:"id"`
	IP   net.IP `json:"ip"`
	Channel uint16 `json:"channel"`
}

//
//
func UUIDLocationText(id ID, schemeMachineChannel string) string {
	machineChannel := discardSchemeConditionalSpecified(schemeMachineChannel)
	return fmt.Sprintf("REDACTED", id, machineChannel)
}

//
//
//
//
//
func FreshNetworkLocator(id ID, location net.Addr) *NetworkLocator {
	tcpsocketLocation, ok := location.(*net.TCPAddr)
	if !ok {
		if flag.Lookup("REDACTED") == nil { //
			panic(fmt.Sprintf("REDACTED", location))
		}
		//
		networkLocation := FreshNetworkLocatorINETChannel(net.IP("REDACTED"), 0)
		networkLocation.ID = id
		return networkLocation
	}

	if err := certifyUUID(id); err != nil {
		panic(fmt.Sprintf("REDACTED", id, err, location))
	}

	ip := tcpsocketLocation.IP
	channel := uint16(tcpsocketLocation.Port)
	na := FreshNetworkLocatorINETChannel(ip, channel)
	na.ID = id
	return na
}

//
//
//
//
func FreshNetworkLocatorText(location string) (*NetworkLocator, error) {
	locationLackingScheme := discardSchemeConditionalSpecified(location)
	spl := strings.Split(locationLackingScheme, "REDACTED")
	if len(spl) != 2 {
		return nil, FaultNetworkLocatorNegativeUUID{location}
	}

	//
	if err := certifyUUID(ID(spl[0])); err != nil {
		return nil, FaultNetworkLocatorUnfit{locationLackingScheme, err}
	}
	var id ID
	id, locationLackingScheme = ID(spl[0]), spl[1]

	//
	machine, channelTxt, err := net.SplitHostPort(locationLackingScheme)
	if err != nil {
		return nil, FaultNetworkLocatorUnfit{locationLackingScheme, err}
	}
	if len(machine) == 0 {
		return nil, FaultNetworkLocatorUnfit{
			locationLackingScheme,
			errors.New("REDACTED"),
		}
	}

	ip := net.ParseIP(machine)
	if ip == nil {
		ips, err := net.LookupIP(machine)
		if err != nil {
			return nil, FaultNetworkLocatorSearch{machine, err}
		}
		ip = ips[0]
	}

	channel, err := strconv.ParseUint(channelTxt, 10, 16)
	if err != nil {
		return nil, FaultNetworkLocatorUnfit{channelTxt, err}
	}

	na := FreshNetworkLocatorINETChannel(ip, uint16(channel))
	na.ID = id
	return na, nil
}

//
//
func FreshNetworkLocatorTexts(locations []string) ([]*NetworkLocator, []error) {
	networkLocations := make([]*NetworkLocator, 0)
	errors := make([]error, 0)
	for _, location := range locations {
		networkLocation, err := FreshNetworkLocatorText(location)
		if err != nil {
			errors = append(errors, err)
		} else {
			networkLocations = append(networkLocations, networkLocation)
		}
	}
	return networkLocations, errors
}

//
//
func FreshNetworkLocatorINETChannel(ip net.IP, channel uint16) *NetworkLocator {
	return &NetworkLocator{
		IP:   ip,
		Channel: channel,
	}
}

//
func NetworkLocatorOriginatingSchema(pb tmpfabric.NetworkLocator) (*NetworkLocator, error) {
	ip := net.ParseIP(pb.IP)
	if ip == nil {
		return nil, fmt.Errorf("REDACTED", pb.IP)
	}
	if pb.Channel >= 1<<16 {
		return nil, fmt.Errorf("REDACTED", pb.Channel)
	}
	return &NetworkLocator{
		ID:   ID(pb.ID),
		IP:   ip,
		Channel: uint16(pb.Channel),
	}, nil
}

//
func NetworkLocatorsOriginatingSchema(pbs []tmpfabric.NetworkLocator) ([]*NetworkLocator, error) {
	nas := make([]*NetworkLocator, 0, len(pbs))
	for _, pb := range pbs {
		na, err := NetworkLocatorOriginatingSchema(pb)
		if err != nil {
			return nil, err
		}
		nas = append(nas, na)
	}
	return nas, nil
}

//
func NetworkLocatorsTowardSchema(nas []*NetworkLocator) []tmpfabric.NetworkLocator {
	pbs := make([]tmpfabric.NetworkLocator, 0, len(nas))
	for _, na := range nas {
		if na != nil {
			pbs = append(pbs, na.TowardSchema())
		}
	}
	return pbs
}

//
func (na *NetworkLocator) TowardSchema() tmpfabric.NetworkLocator {
	return tmpfabric.NetworkLocator{
		ID:   string(na.ID),
		IP:   na.IP.String(),
		Channel: uint32(na.Channel),
	}
}

//
//
func (na *NetworkLocator) Matches(another any) bool {
	if o, ok := another.(*NetworkLocator); ok {
		return na.Text() == o.Text()
	}
	return false
}

//
func (na *NetworkLocator) Identical(another any) bool {
	if o, ok := another.(*NetworkLocator); ok {
		if na.CallText() == o.CallText() {
			return true
		}
		if na.ID != "REDACTED" && na.ID == o.ID {
			return true
		}
	}
	return false
}

//
func (na *NetworkLocator) Text() string {
	if na == nil {
		return BlankNetworkLocator
	}

	locationTxt := na.CallText()
	if na.ID != "REDACTED" {
		locationTxt = UUIDLocationText(na.ID, locationTxt)
	}

	return locationTxt
}

func (na *NetworkLocator) CallText() string {
	if na == nil {
		return "REDACTED"
	}
	return net.JoinHostPort(
		na.IP.String(),
		strconv.FormatUint(uint64(na.Channel), 10),
	)
}

//
func (na *NetworkLocator) Call() (net.Conn, error) {
	link, err := net.Dial("REDACTED", na.CallText())
	if err != nil {
		return nil, err
	}
	return link, nil
}

//
func (na *NetworkLocator) CallDeadline(deadline time.Duration) (net.Conn, error) {
	link, err := net.DialTimeout("REDACTED", na.CallText(), deadline)
	if err != nil {
		return nil, err
	}
	return link, nil
}

//
func (na *NetworkLocator) Directable() bool {
	if err := na.Sound(); err != nil {
		return false
	}
	//
	return !na.Rfc1918() && !na.Rfc3927() && !na.Rfc4862() && !na.Rfc4193() && !na.Rfc4843() && !na.Regional()
}

//
//
func (na *NetworkLocator) Sound() error {
	if err := certifyUUID(na.ID); err != nil {
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
func (na *NetworkLocator) OwnsUUID() bool {
	return string(na.ID) != "REDACTED"
}

//
func (na *NetworkLocator) Regional() bool {
	return na.IP.IsLoopback() || null4.Contains(na.IP)
}

//
func (na *NetworkLocator) AccessibilityToward(o *NetworkLocator) int {
	const (
		Inaccessible = 0
		Fallback     = iota
		Teredo
		Ipv6fragile
		Ipv4
		Ipv6robust
	)
	switch {
	case !na.Directable():
		return Inaccessible
	case na.Rfc4380():
		switch {
		case !o.Directable():
			return Fallback
		case o.Rfc4380():
			return Teredo
		case o.IP.To4() != nil:
			return Ipv4
		default: //
			return Ipv6fragile
		}
	case na.IP.To4() != nil:
		if o.Directable() && o.IP.To4() != nil {
			return Ipv4
		}
		return Fallback
	default: //
		var channeled bool
		//
		if o.Rfc3964() || o.Rfc6052() || o.Rfc6145() {
			channeled = true
		}
		switch {
		case !o.Directable():
			return Fallback
		case o.Rfc4380():
			return Teredo
		case o.IP.To4() != nil:
			return Ipv4
		case channeled:
			//
			return Ipv6fragile
		}
		return Ipv6robust
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
	rfc1918_one0  = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(8, 32)}
	rfc1918_one92 = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(16, 32)}
	rfc1918_one72 = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(12, 32)}
	rfc3849     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(32, 128)}
	rfc3927     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(16, 32)}
	rfc3964     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(16, 128)}
	rfc4193     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(7, 128)}
	rfc4380     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(32, 128)}
	rfc4843     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(28, 128)}
	rfc4862     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(64, 128)}
	rfc6052     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(96, 128)}
	rfc6145     = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(96, 128)}
	null4       = net.IPNet{IP: net.ParseIP("REDACTED"), Mask: net.CIDRMask(8, 32)}
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
var onionConcatenateNetwork = inetNetwork("REDACTED", 48, 128)

//
//
//
func inetNetwork(ip string, unity, digits int) net.IPNet {
	return net.IPNet{IP: net.ParseIP(ip), Mask: net.CIDRMask(unity, digits)}
}

func (na *NetworkLocator) Rfc1918() bool {
	return rfc1918_one0.Contains(na.IP) ||
		rfc1918_one92.Contains(na.IP) ||
		rfc1918_one72.Contains(na.IP)
}
func (na *NetworkLocator) Rfc3849() bool     { return rfc3849.Contains(na.IP) }
func (na *NetworkLocator) Rfc3927() bool     { return rfc3927.Contains(na.IP) }
func (na *NetworkLocator) Rfc3964() bool     { return rfc3964.Contains(na.IP) }
func (na *NetworkLocator) Rfc4193() bool     { return rfc4193.Contains(na.IP) }
func (na *NetworkLocator) Rfc4380() bool     { return rfc4380.Contains(na.IP) }
func (na *NetworkLocator) Rfc4843() bool     { return rfc4843.Contains(na.IP) }
func (na *NetworkLocator) Rfc4862() bool     { return rfc4862.Contains(na.IP) }
func (na *NetworkLocator) Rfc6052() bool     { return rfc6052.Contains(na.IP) }
func (na *NetworkLocator) Rfc6145() bool     { return rfc6145.Contains(na.IP) }
func (na *NetworkLocator) OnionConcatenateTor() bool { return onionConcatenateNetwork.Contains(na.IP) }

func discardSchemeConditionalSpecified(location string) string {
	if strings.Contains(location, "REDACTED") {
		return strings.Split(location, "REDACTED")[1]
	}
	return location
}

func certifyUUID(id ID) error {
	if len(id) == 0 {
		return errors.New("REDACTED")
	}
	uuidOctets, err := hex.DecodeString(string(id))
	if err != nil {
		return err
	}
	if len(uuidOctets) != UUIDOctetMagnitude {
		return fmt.Errorf("REDACTED", len(uuidOctets), UUIDOctetMagnitude)
	}
	return nil
}
