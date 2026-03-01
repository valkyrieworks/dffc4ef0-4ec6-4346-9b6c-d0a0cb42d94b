package netp2p

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"

	cryptography "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

const (
	layer4udpsocket = "REDACTED"
)

func UUIDOriginatingSecludedToken(universeKEY cryptography.PrivateToken) (peer.ID, error) {
	pk, err := secludedTokenOriginatingUniverseToken(universeKEY)
	if err != nil {
		return "REDACTED", fmt.Errorf("REDACTED", err)
	}

	return peer.IDFromPrivateKey(pk)
}

//
//
//
func LocatorTowardVariedLocation(location string, carrier string) (ma.Multiaddr, error) {
	if !strings.Contains(location, "REDACTED") {
		location = "REDACTED" + location
	}

	fragments, err := url.Parse(location)
	switch {
	case err != nil:
		return nil, fmt.Errorf("REDACTED", err)
	case fragments.Hostname() == "REDACTED":
		return nil, fmt.Errorf("REDACTED")
	case fragments.Port() == "REDACTED":
		return nil, fmt.Errorf("REDACTED")
	case carrier == CarrierQuicprotocol:
		return locationTowardQuicprotocolMultilocator(fragments, layer4udpsocket)
	}

	return nil, fmt.Errorf("REDACTED", carrier)
}

func LocationDetailsOriginatingMachineAlsoUUID(machine, id string) (peer.AddrInfo, error) {
	location, err := LocatorTowardVariedLocation(machine, CarrierQuicprotocol)
	if err != nil {
		return peer.AddrInfo{}, fmt.Errorf("REDACTED", err)
	}

	nodeUUID, err := peer.Decode(id)
	if err != nil {
		return peer.AddrInfo{}, fmt.Errorf("REDACTED", err)
	}

	return peer.AddrInfo{ID: nodeUUID, Addrs: []ma.Multiaddr{location}}, nil
}

//
//
//
func locationTowardQuicprotocolMultilocator(fragments *url.URL, layer4socket string) (ma.Multiaddr, error) {
	machinename := fragments.Hostname()

	//
	var fabricSchema string
	if ip := net.ParseIP(machinename); ip != nil {
		if ip.To4() != nil {
			fabricSchema = "REDACTED"
		} else {
			fabricSchema = "REDACTED"
		}
	} else {
		//
		fabricSchema = "REDACTED"
	}

	raw := fmt.Sprintf("REDACTED", fabricSchema, machinename, layer4socket, fragments.Port(), CarrierQuicprotocol)

	return ma.NewMultiaddr(raw)
}

//
func networkLocatorOriginatingNode(locationDetails peer.AddrInfo) (*p2p.NetworkLocator, error) {
	if len(locationDetails.Addrs) == 0 {
		return nil, fmt.Errorf("REDACTED")
	}

	//
	_, inetChannel, err := manet.DialArgs(locationDetails.Addrs[0])
	if err != nil {
		return nil, err
	}

	fragments := strings.Split(inetChannel, "REDACTED")
	if len(fragments) != 2 {
		return nil, fmt.Errorf("REDACTED", inetChannel)
	}

	ip := net.ParseIP(fragments[0])
	if ip == nil {
		//
		ips, err := net.LookupIP(fragments[0])
		if err != nil || len(ips) == 0 {
			return nil, fmt.Errorf("REDACTED", fragments[0])
		}
		//
		ip = favorIDXPrivatevalue4(ips)
	}

	channel, err := strconv.ParseUint(fragments[1], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", fragments[1])
	}

	return &p2p.NetworkLocator{
		ID:   nodeUUIDTowardToken(locationDetails.ID),
		IP:   ip,
		Channel: uint16(channel),
	}, nil
}

func favorIDXPrivatevalue4(ips []net.IP) net.IP {
	for _, ip := range ips {
		if ip.To4() != nil {
			return ip
		}
	}
	return ips[0]
}
