//
//
package netpeer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/valkyrieworks/settings"
	cmcrypto "github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/utils/log"
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
	onboardNodes []OnboardNode

	tracer log.Tracer

	nodeBreakdownProcessors []func(id peer.ID, err error)
}

//
type OnboardNode struct {
	AddressDetails      peer.AddrInfo
	Internal       bool
	Durable    bool
	Absolute bool
}

//
//
const CarrierQUIC = "REDACTED"

//
func NewMachine(settings *settings.P2PSettings, memberKey cmcrypto.PrivateKey, tracer log.Tracer) (*Machine, error) {
	if !settings.LibraryP2PActivated() {
		return nil, fmt.Errorf("REDACTED")
	}

	internalKey, err := internalKeyFromCosmosKey(memberKey)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	acceptAddress, err := LocationToMultipleAddress(settings.AcceptLocation, CarrierQUIC)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", settings.AcceptLocation, err)
	}

	onboardNodes, err := OnboardNodesFromSettings(settings)
	switch {
	case err != nil:
		return nil, fmt.Errorf("REDACTED", err)
	case len(onboardNodes) == 0:
		tracer.Details("REDACTED")
	}

	//
	//
	opts := []libp2p.Option{
		libp2p.Identity(internalKey),
		libp2p.ListenAddrs(acceptAddress),
		libp2p.UserAgent("REDACTED"),
		libp2p.Ping(true),
		libp2p.Transport(quic.NewTransport),
	}

	if settings.LibraryP2PSettings.DeactivateAssetAdministrator {
		opts = append(opts, libp2p.ResourceManager(&network.NullResourceManager{}))
	}

	//
	if settings.OutsideLocation != "REDACTED" {
		outsideAddress, err := LocationToMultipleAddress(settings.OutsideLocation, CarrierQUIC)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", settings.OutsideLocation, err)
		}

		opts = append(opts, withLocationBuilder(outsideAddress))
	}

	machine, err := libp2p.New(opts...)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return &Machine{
		Machine:           machine,
		onboardNodes: onboardNodes,
		tracer:         tracer,
	}, nil
}

func (h *Machine) AddressDetails() peer.AddrInfo {
	return peer.AddrInfo{ID: h.ID(), Addrs: h.Addrs()}
}

func (h *Machine) OnboardNodes() []OnboardNode {
	return h.onboardNodes
}

func (h *Machine) Tracer() log.Tracer {
	return h.tracer
}

//
//
func (h *Machine) Ping(ctx context.Context, addressDetails peer.AddrInfo) (time.Duration, error) {
	res := <-ping.Ping(ctx, h, addressDetails.ID)

	return res.RTT, res.Error
}

func (h *Machine) AppendNodeBreakdownManager(manager func(id peer.ID, err error)) {
	h.nodeBreakdownProcessors = append(h.nodeBreakdownProcessors, manager)
}

//
//
func (h *Machine) IssueNodeBreakdown(id peer.ID, err error) {
	for _, manager := range h.nodeBreakdownProcessors {
		go manager(id, err)
	}
}

func (h *Machine) multipleAddressStrByUID(id peer.ID) string {
	return multipleAddressStr(h.Peerstore().Addrs(id))
}

func OnboardNodesFromSettings(settings *settings.P2PSettings) ([]OnboardNode, error) {
	nodes := make([]OnboardNode, 0, len(settings.LibraryP2PSettings.OnboardNodes))

	//
	repository := make(map[peer.ID]struct{})

	for _, bp := range settings.LibraryP2PSettings.OnboardNodes {
		address, err := AddressDetailsFromMachineAndUID(bp.Machine, bp.ID)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", bp.Machine, bp.ID, err)
		}

		if _, ok := repository[address.ID]; ok {
			continue
		}

		nodes = append(nodes, OnboardNode{
			AddressDetails:      address,
			Internal:       bp.Internal,
			Durable:    bp.Durable,
			Absolute: bp.Absolute,
		})

		repository[address.ID] = struct{}{}
	}

	return nodes, nil
}

//
func IsDNSAddress(address ma.Multiaddr) bool {
	for _, a := range address {
		code := a.Protocol().Code

		if code == ma.P_DNS || code == ma.P_DNS4 || code == ma.P_DNS6 || code == ma.P_DNSADDR {
			return true
		}
	}

	return false
}

func multipleAddressStr(locations []ma.Multiaddr) string {
	if len(locations) == 0 {
		return "REDACTED"
	}

	segments := make([]string, len(locations))
	for i, address := range locations {
		segments[i] = address.String()
	}

	return strings.Join(segments, "REDACTED")
}
