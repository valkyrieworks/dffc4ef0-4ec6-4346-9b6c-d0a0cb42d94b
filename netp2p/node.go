package netp2p

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/net/swarm"
)

//
//
//
type Node struct {
	facility.FoundationFacility

	machine *Machine

	//
	//
	//
	//
	//
	locationDetails peer.AddrInfo

	networkLocation *p2p.NetworkLocator

	//
	equalsSecluded       bool
	equalsEnduring    bool
	equalsAbsolute bool

	telemetry *p2p.Telemetry
}

var _ p2p.Node = (*Node)(nil)

func FreshNode(
	machine *Machine,
	locationDetails peer.AddrInfo,
	telemetry *p2p.Telemetry,
	equalsSecluded, equalsEnduring, equalsAbsolute bool,
) (*Node, error) {
	networkLocation, err := networkLocatorOriginatingNode(locationDetails)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	p := &Node{
		machine:     machine,
		locationDetails: locationDetails,
		networkLocation:  networkLocation,

		equalsSecluded:       equalsSecluded,
		equalsEnduring:    equalsEnduring,
		equalsAbsolute: equalsAbsolute,

		telemetry: telemetry,
	}

	tracer := machine.Tracer().Using("REDACTED", locationDetails.ID.String())

	p.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", p)
	p.AssignTracer(tracer)

	return p, nil
}

func (p *Node) Text() string {
	return fmt.Sprintf("REDACTED", p.ID())
}

func (p *Node) ID() p2p.ID {
	return nodeUUIDTowardToken(p.locationDetails.ID)
}

func (p *Node) PortLocation() *p2p.NetworkLocator {
	return p.networkLocation
}

//
//
func (p *Node) LocationDetails() peer.AddrInfo {
	return p.locationDetails
}

func (p *Node) Get(key string) any {
	v, err := p.machine.Peerstore().Get(p.locationDetails.ID, key)
	if err != nil {
		return nil
	}

	return v
}

func (p *Node) Set(key string, datum any) {
	//
	p.machine.Peerstore().Put(p.locationDetails.ID, key, datum)
}

func (p *Node) EqualsEnduring() bool {
	return p.equalsEnduring
}

func (p *Node) EqualsSecluded() bool {
	//
	return p.equalsSecluded
}

func (p *Node) EqualsAbsolute() bool {
	return p.equalsAbsolute
}

//
func (p *Node) Transmit(e p2p.Wrapper) bool {
	if err := p.transmit(e); err != nil {
		p.Tracer.Failure("REDACTED", "REDACTED", e.ConduitUUID, "REDACTED", "REDACTED", "REDACTED", err)
		p.processTransmitFault(err)
		return false
	}

	return true
}

func (p *Node) AttemptTransmit(e p2p.Wrapper) bool {
	//
	if err := p.transmit(e); err != nil {
		p.Tracer.Failure("REDACTED", "REDACTED", e.ConduitUUID, "REDACTED", "REDACTED", "REDACTED", err)
		p.processTransmitFault(err)
		return false
	}

	return true
}

func (p *Node) ShutdownLink() error {
	return p.machine.Network().ClosePeer(p.locationDetails.ID)
}

func (p *Node) transmit(e p2p.Wrapper) (err error) {
	var (
		nodeUUID     = p.locationDetails.ID
		schemeUUID = SchemeUUID(e.ConduitUUID)
	)

	content, err := serializeSchema(e.Signal)
	if err != nil {
		return err
	}

	var (
		nodeUUIDTxt    = nodeUUID.String()
		signalKind  = schemaKindAlias(e.Signal)
		contentLength   = float64(len(content))
		indicatorTags = []string{
			"REDACTED", nodeUUIDTxt,
			"REDACTED", fmt.Sprintf("REDACTED", e.ConduitUUID),
		}

		nodeTransmitStagingExtent = p.telemetry.NodeTransmitStagingExtent.With("REDACTED", nodeUUIDTxt)
	)

	nodeTransmitStagingExtent.Add(1)

	ctx, abort := context.WithTimeout(context.Background(), DeadlineInflux)
	defer abort()

	initiate := time.Now()

	defer func() {
		nodeTransmitStagingExtent.Add(-1)

		if err != nil {
			return
		}

		p.telemetry.NodeTransmitOctetsSum.With(indicatorTags...).Add(contentLength)
		p.telemetry.ArtifactTransmitOctetsSum.With("REDACTED", signalKind).Add(contentLength)

		p.Tracer.Diagnose(
			"REDACTED",
			"REDACTED", schemeUUID,
			"REDACTED", nodeUUIDTxt,
			"REDACTED", time.Since(initiate).String(),
		)
	}()

	//
	s, err := p.machine.NewStream(ctx, nodeUUID, schemeUUID)
	if err != nil {
		return fmt.Errorf("REDACTED", schemeUUID, err)
	}

	return InfluxRecordShutdown(s, content)
}

func (p *Node) processTransmitFault(err error) {
	switch {
	case err == nil:
		return
	case errors.Is(err, swarm.ErrAllDialsFailed), errors.Is(err, swarm.ErrNoGoodAddresses):
		p.machine.RelayNodeBreakdown(p.locationDetails.ID, err)
	}
}

//
//
//
func (p *Node) PeerDetails() p2p.PeerDetails {
	return p2p.FallbackPeerDetails{
		FallbackPeerUUID: p.ID(),
		OverhearLocation:    p.networkLocation.CallText(),
	}
}

//
func (p *Node) DistantINET() net.IP {
	return p.networkLocation.IP
}

//
func (p *Node) DistantLocation() net.Addr {
	return &net.TCPAddr{
		IP:   p.networkLocation.IP,
		Port: int(p.networkLocation.Channel),
	}
}

//
func (*Node) EqualsOutgoing() bool { return true }

//
//
func (*Node) Condition() link.LinkageCondition { return link.LinkageCondition{} }

func (*Node) PurgeHalt()             {}
func (*Node) AssignDeletionUnsuccessful()      {}
func (*Node) ObtainDeletionUnsuccessful() bool { return false }
