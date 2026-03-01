//
//
package netp2p

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	cryptography "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	ma "github.com/multiformats/go-multiaddr"
)

//
//
//
type Machine struct {
	host.Machine

	//
	initiateNodes []OnboardNode

	tracer log.Tracer

	nodeBreakdownReactors []func(id peer.ID, err error)
}

//
type OnboardNode struct {
	LocationDetails      peer.AddrInfo
	Secluded       bool
	Enduring    bool
	Absolute bool
}

//
//
const CarrierQuicprotocol = "REDACTED"

//
func FreshMachine(settings *settings.Peer2peerSettings, peerToken cryptography.PrivateToken, tracer log.Tracer) (*Machine, error) {
	if !settings.LibraryPeer2peerActivated() {
		return nil, fmt.Errorf("REDACTED")
	}

	secludedToken, err := secludedTokenOriginatingUniverseToken(peerToken)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	overhearLocation, err := LocatorTowardVariedLocation(settings.OverhearLocation, CarrierQuicprotocol)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", settings.OverhearLocation, err)
	}

	initiateNodes, err := OnboardNodesOriginatingSettings(settings)
	switch {
	case err != nil:
		return nil, fmt.Errorf("REDACTED", err)
	case len(initiateNodes) == 0:
		tracer.Details("REDACTED")
	}

	//
	//
	choices := []libp2p.Option{
		libp2p.Identity(secludedToken),
		libp2p.ListenAddrs(overhearLocation),
		libp2p.UserAgent("REDACTED"),
		libp2p.Ping(true),
		libp2p.Transport(quic.NewTransport),
	}

	if settings.LibraryPeer2peerSettings.DeactivateAssetAdministrator {
		choices = append(choices, libp2p.ResourceManager(&network.NullResourceManager{}))
	}

	//
	if settings.OutsideLocation != "REDACTED" {
		outsideLocation, err := LocatorTowardVariedLocation(settings.OutsideLocation, CarrierQuicprotocol)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", settings.OutsideLocation, err)
		}

		choices = append(choices, usingLocatorBuilder(outsideLocation))
	}

	machine, err := libp2p.New(choices...)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return &Machine{
		Machine:           machine,
		initiateNodes: initiateNodes,
		tracer:         tracer,
	}, nil
}

func (h *Machine) LocationDetails() peer.AddrInfo {
	return peer.AddrInfo{ID: h.ID(), Addrs: h.Addrs()}
}

func (h *Machine) InitiateNodes() []OnboardNode {
	return h.initiateNodes
}

func (h *Machine) Tracer() log.Tracer {
	return h.tracer
}

//
//
func (h *Machine) Ping(ctx context.Context, locationDetails peer.AddrInfo) (time.Duration, error) {
	res := <-ping.Ping(ctx, h, locationDetails.ID)

	return res.RTT, res.Error
}

func (h *Machine) AppendNodeBreakdownProcessor(processor func(id peer.ID, err error)) {
	h.nodeBreakdownReactors = append(h.nodeBreakdownReactors, processor)
}

//
//
func (h *Machine) RelayNodeBreakdown(id peer.ID, err error) {
	for _, processor := range h.nodeBreakdownReactors {
		go processor(id, err)
	}
}

func (h *Machine) variedLocationTxtViaUUID(id peer.ID) string {
	return variedLocationTxt(h.Peerstore().Addrs(id))
}

func OnboardNodesOriginatingSettings(settings *settings.Peer2peerSettings) ([]OnboardNode, error) {
	nodes := make([]OnboardNode, 0, len(settings.LibraryPeer2peerSettings.InitiateNodes))

	//
	stash := make(map[peer.ID]struct{})

	for _, bp := range settings.LibraryPeer2peerSettings.InitiateNodes {
		location, err := LocationDetailsOriginatingMachineAlsoUUID(bp.Machine, bp.ID)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", bp.Machine, bp.ID, err)
		}

		if _, ok := stash[location.ID]; ok {
			continue
		}

		nodes = append(nodes, OnboardNode{
			LocationDetails:      location,
			Secluded:       bp.Secluded,
			Enduring:    bp.Enduring,
			Absolute: bp.Absolute,
		})

		stash[location.ID] = struct{}{}
	}

	return nodes, nil
}

//
func EqualsDomainLocation(location ma.Multiaddr) bool {
	for _, a := range location {
		cipher := a.Protocol().Code

		if cipher == ma.P_DNS || cipher == ma.P_DNS4 || cipher == ma.P_DNS6 || cipher == ma.P_DNSADDR {
			return true
		}
	}

	return false
}

func variedLocationTxt(locations []ma.Multiaddr) string {
	if len(locations) == 0 {
		return "REDACTED"
	}

	fragments := make([]string, len(locations))
	for i, location := range locations {
		fragments[i] = location.String()
	}

	return strings.Join(fragments, "REDACTED")
}
