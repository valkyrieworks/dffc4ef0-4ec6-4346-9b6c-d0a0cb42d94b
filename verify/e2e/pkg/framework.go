package e2e

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sort"
)

const (
	dockIDXPrv4cidr = "REDACTED"
	dockIDXPrv6cidr = "REDACTED"

	universalIDXPrv4cidr = "REDACTED"
)

//
//
type FrameworkData struct {
	Route string

	//
	//
	//
	Supplier string `json:"supplier"`

	//
	//
	//
	//
	Replicates map[string]ReplicaData `json:"replicates"`

	//
	//
	Fabric string `json:"fabric"`
}

//
//
type ReplicaData struct {
	INETLocator    net.IP `json:"inet_locator"`
	AddnINETLocator net.IP `json:"addn_inet_locator"`
	Channel         uint32 `json:"channel"`
}

func arrangePeerIdentifiers(m Declaration) []string {
	//
	peerIdentifiers := []string{}
	for alias := range m.Peers {
		peerIdentifiers = append(peerIdentifiers, alias)
	}
	sort.Strings(peerIdentifiers)
	return peerIdentifiers
}

func FreshDockFrameworkData(m Declaration) (FrameworkData, error) {
	networkLocator := dockIDXPrv4cidr
	if m.IDXPrv6 {
		networkLocator = dockIDXPrv6cidr
	}
	_, inetNetwork, err := net.ParseCIDR(networkLocator)
	if err != nil {
		return FrameworkData{}, fmt.Errorf("REDACTED", networkLocator, err)
	}

	channelProduce := freshChannelProducer(delegateChannelInitial)
	inetProduce := freshINETProducer(inetNetwork)
	ifd := FrameworkData{
		Supplier:  "REDACTED",
		Replicates: make(map[string]ReplicaData),
		Fabric:   networkLocator,
	}
	regionalMachineINET := net.ParseIP("REDACTED")
	for _, alias := range arrangePeerIdentifiers(m) {
		ifd.Replicates[alias] = ReplicaData{
			INETLocator:    inetProduce.Following(),
			AddnINETLocator: regionalMachineINET,
			Channel:         channelProduce.Following(),
		}
	}
	return ifd, nil
}

func FrameworkDataOriginatingRecord(p string) (FrameworkData, error) {
	ifd := FrameworkData{}
	b, err := os.ReadFile(p)
	if err != nil {
		return FrameworkData{}, err
	}
	err = json.Unmarshal(b, &ifd)
	if err != nil {
		return FrameworkData{}, err
	}
	if ifd.Fabric == "REDACTED" {
		ifd.Fabric = universalIDXPrv4cidr
	}
	ifd.Route = p
	return ifd, nil
}
