package netpeer

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"

	cmcrypto "github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/p2p"
	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

const (
	transportudp = "REDACTED"
)

func UIDFromPrivateKey(cosmosPublicid cmcrypto.PrivateKey) (peer.ID, error) {
	pk, err := internalKeyFromCosmosKey(cosmosPublicid)
	if err != nil {
		return "REDACTED", fmt.Errorf("REDACTED", err)
	}

	return peer.IDFromPrivateKey(pk)
}

//
//
//
func LocationToMultipleAddress(address string, carrier string) (ma.Multiaddr, error) {
	if !strings.Contains(address, "REDACTED") {
		address = "REDACTED" + address
	}

	segments, err := url.Parse(address)
	switch {
	case err != nil:
		return nil, fmt.Errorf("REDACTED", err)
	case segments.Hostname() == "REDACTED":
		return nil, fmt.Errorf("REDACTED")
	case segments.Port() == "REDACTED":
		return nil, fmt.Errorf("REDACTED")
	case carrier == CarrierQUIC:
		return addressToQuicMultiaddress(segments, transportudp)
	}

	return nil, fmt.Errorf("REDACTED", carrier)
}

func AddressDetailsFromMachineAndUID(machine, id string) (peer.AddrInfo, error) {
	address, err := LocationToMultipleAddress(machine, CarrierQUIC)
	if err != nil {
		return peer.AddrInfo{}, fmt.Errorf("REDACTED", err)
	}

	nodeUID, err := peer.Decode(id)
	if err != nil {
		return peer.AddrInfo{}, fmt.Errorf("REDACTED", err)
	}

	return peer.AddrInfo{ID: nodeUID, Addrs: []ma.Multiaddr{address}}, nil
}

//
//
//
func addressToQuicMultiaddress(segments *url.URL, transport string) (ma.Multiaddr, error) {
	hostlabel := segments.Hostname()

	//
	var fabricSchema string
	if ip := net.ParseIP(hostlabel); ip != nil {
		if ip.To4() != nil {
			fabricSchema = "REDACTED"
		} else {
			fabricSchema = "REDACTED"
		}
	} else {
		//
		fabricSchema = "REDACTED"
	}

	raw := fmt.Sprintf("REDACTED", fabricSchema, hostlabel, transport, segments.Port(), CarrierQUIC)

	return ma.NewMultiaddr(raw)
}

//
func netLocationFromNode(addressDetails peer.AddrInfo) (*p2p.NetLocation, error) {
	if len(addressDetails.Addrs) == 0 {
		return nil, fmt.Errorf("REDACTED")
	}

	//
	_, ipPort, err := manet.DialArgs(addressDetails.Addrs[0])
	if err != nil {
		return nil, err
	}

	segments := strings.Split(ipPort, "REDACTED")
	if len(segments) != 2 {
		return nil, fmt.Errorf("REDACTED", ipPort)
	}

	ip := net.ParseIP(segments[0])
	if ip == nil {
		//
		ips, err := net.LookupIP(segments[0])
		if err != nil || len(ips) == 0 {
			return nil, fmt.Errorf("REDACTED", segments[0])
		}
		//
		ip = favorIDXPv4(ips)
	}

	port, err := strconv.ParseUint(segments[1], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", segments[1])
	}

	return &p2p.NetLocation{
		ID:   nodeUIDToKey(addressDetails.ID),
		IP:   ip,
		Port: uint16(port),
	}, nil
}

func favorIDXPv4(ips []net.IP) net.IP {
	for _, ip := range ips {
		if ip.To4() != nil {
			return ip
		}
	}
	return ips[0]
}
