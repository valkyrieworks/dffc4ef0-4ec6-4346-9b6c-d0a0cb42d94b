package e2e

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sort"
)

const (
	dockerIDXIpv4range = "REDACTED"
	dockerIDXIpv6range = "REDACTED"

	universalIDXIpv4range = "REDACTED"
)

//
//
type PlatformData struct {
	Route string

	//
	//
	//
	Source string `json:"source"`

	//
	//
	//
	//
	Occurrences map[string]OccurrenceData `json:"occurrences"`

	//
	//
	Fabric string `json:"fabric"`
}

//
//
type OccurrenceData struct {
	IPLocation    net.IP `json:"ip_location"`
	ExtensionIPLocation net.IP `json:"extension_ip_location"`
	Port         uint32 `json:"port"`
}

func arrangeMemberLabels(m Declaration) []string {
	//
	memberLabels := []string{}
	for label := range m.Instances {
		memberLabels = append(memberLabels, label)
	}
	sort.Strings(memberLabels)
	return memberLabels
}

func NewDockerPlatformData(m Declaration) (PlatformData, error) {
	netLocation := dockerIDXIpv4range
	if m.IDXIpv6 {
		netLocation = dockerIDXIpv6range
	}
	_, ipNet, err := net.ParseCIDR(netLocation)
	if err != nil {
		return PlatformData{}, fmt.Errorf("REDACTED", netLocation, err)
	}

	portGenerate := newPortProducer(gatewayPortInitial)
	ipGenerate := newIPProducer(ipNet)
	ifd := PlatformData{
		Source:  "REDACTED",
		Occurrences: make(map[string]OccurrenceData),
		Fabric:   netLocation,
	}
	nativeMachineIP := net.ParseIP("REDACTED")
	for _, label := range arrangeMemberLabels(m) {
		ifd.Occurrences[label] = OccurrenceData{
			IPLocation:    ipGenerate.Following(),
			ExtensionIPLocation: nativeMachineIP,
			Port:         portGenerate.Following(),
		}
	}
	return ifd, nil
}

func PlatformDataFromEntry(p string) (PlatformData, error) {
	ifd := PlatformData{}
	b, err := os.ReadFile(p)
	if err != nil {
		return PlatformData{}, err
	}
	err = json.Unmarshal(b, &ifd)
	if err != nil {
		return PlatformData{}, err
	}
	if ifd.Fabric == "REDACTED" {
		ifd.Fabric = universalIDXIpv4range
	}
	ifd.Route = p
	return ifd, nil
}
