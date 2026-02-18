package netpeer

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/p2p/link"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
)

//
//
//
type Node struct {
	daemon.RootDaemon

	machine *Machine

	//
	//
	//
	//
	//
	addressDetails peer.AddrInfo

	netAddress *p2p.NetLocation

	//
	isInternal       bool
	isDurable    bool
	isAbsolute bool

	stats *p2p.Stats
}

var _ p2p.Node = (*Node)(nil)

func NewNode(
	machine *Machine,
	addressDetails peer.AddrInfo,
	stats *p2p.Stats,
	isInternal, isDurable, isAbsolute bool,
) (*Node, error) {
	netAddress, err := netLocationFromNode(addressDetails)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	p := &Node{
		machine:     machine,
		addressDetails: addressDetails,
		netAddress:  netAddress,

		isInternal:       isInternal,
		isDurable:    isDurable,
		isAbsolute: isAbsolute,

		stats: stats,
	}

	tracer := machine.Tracer().With("REDACTED", addressDetails.ID.String())

	p.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", p)
	p.AssignTracer(tracer)

	return p, nil
}

func (p *Node) String() string {
	return fmt.Sprintf("REDACTED", p.ID())
}

func (p *Node) ID() p2p.ID {
	return nodeUIDToKey(p.addressDetails.ID)
}

func (p *Node) SocketAddress() *p2p.NetLocation {
	return p.netAddress
}

//
//
func (p *Node) AddressDetails() peer.AddrInfo {
	return p.addressDetails
}

func (p *Node) Get(key string) any {
	v, err := p.machine.Peerstore().Get(p.addressDetails.ID, key)
	if err != nil {
		return nil
	}

	return v
}

func (p *Node) Set(key string, item any) {
	//
	p.machine.Peerstore().Put(p.addressDetails.ID, key, item)
}

func (p *Node) IsDurable() bool {
	return p.isDurable
}

func (p *Node) IsInternal() bool {
	//
	return p.isInternal
}

func (p *Node) IsAbsolute() bool {
	return p.isAbsolute
}

//
func (p *Node) Transmit(e p2p.Packet) bool {
	if err := p.transmit(e); err != nil {
		p.Tracer.Fault("REDACTED", "REDACTED", e.StreamUID, "REDACTED", "REDACTED", "REDACTED", err)
		p.processTransmitErr(err)
		return false
	}

	return true
}

func (p *Node) AttemptTransmit(e p2p.Packet) bool {
	//
	if err := p.transmit(e); err != nil {
		p.Tracer.Fault("REDACTED", "REDACTED", e.StreamUID, "REDACTED", "REDACTED", "REDACTED", err)
		p.processTransmitErr(err)
		return false
	}

	return true
}

func (p *Node) EndLink() error {
	return p.machine.Network().ClosePeer(p.addressDetails.ID)
}

func (p *Node) transmit(e p2p.Packet) (err error) {
	var (
		nodeUID     = p.addressDetails.ID
		protocolUID = ProtocolUID(e.StreamUID)
	)

	shipment, err := serializeSchema(e.Signal)
	if err != nil {
		return err
	}

	var (
		nodeUIDStr    = nodeUID.String()
		signalKind  = schemaKindLabel(e.Signal)
		shipmentSize   = float64(len(shipment))
		indicatorTags = []string{
			"REDACTED", nodeUIDStr,
			"REDACTED", fmt.Sprintf("REDACTED", e.StreamUID),
		}

		//
		awaitingSignalsTally = p.stats.NodeAwaitingTransmitOctets.With("REDACTED", nodeUIDStr)
	)

	awaitingSignalsTally.Add(1)

	ctx, revoke := context.WithTimeout(context.Background(), DeadlineInflux)
	defer revoke()

	begin := time.Now()

	defer func() {
		awaitingSignalsTally.Add(-1)

		if err != nil {
			return
		}

		p.stats.NodeTransmitOctetsSum.With(indicatorTags...).Add(shipmentSize)
		p.stats.SignalTransmitOctetsSum.With("REDACTED", signalKind).Add(shipmentSize)

		p.Tracer.Diagnose(
			"REDACTED",
			"REDACTED", protocolUID,
			"REDACTED", nodeUIDStr,
			"REDACTED", time.Since(begin).String(),
		)
	}()

	//
	s, err := p.machine.NewStream(ctx, nodeUID, protocolUID)
	if err != nil {
		return fmt.Errorf("REDACTED", protocolUID, err)
	}

	return InfluxRecordEnd(s, shipment)
}

func (p *Node) processTransmitErr(err error) {
	switch {
	case err == nil:
		return
	case errors.Is(err, swarm.ErrAllDialsFailed), errors.Is(err, swarm.ErrNoGoodAddresses):
		p.machine.IssueNodeBreakdown(p.addressDetails.ID, err)
	}
}

//
//
//
func (p *Node) MemberDetails() p2p.MemberDetails {
	return p2p.StandardMemberDetails{
		StandardMemberUID: p.ID(),
		ObserveAddress:    p.netAddress.CallString(),
	}
}

//
func (p *Node) DistantIP() net.IP {
	return p.netAddress.IP
}

//
func (p *Node) DistantAddress() net.Addr {
	return &net.TCPAddr{
		IP:   p.netAddress.IP,
		Port: int(p.netAddress.Port),
	}
}

//
func (*Node) IsOutgoing() bool { return true }

//
//
func (*Node) Status() link.LinkageState { return link.LinkageState{} }

func (*Node) PurgeHalt()             {}
func (*Node) CollectionDeletionErrored()      {}
func (*Node) FetchDeletionErrored() bool { return false }
