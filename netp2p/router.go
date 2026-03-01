package netp2p

import (
	"context"
	"fmt"
	"math/rand"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/pkg/errors"
)

//
//
//
type Router struct {
	facility.FoundationFacility

	peerDetails p2p.PeerDetails //

	machine    *Machine
	nodeAssign *NodeAssign

	engines *handlerAssign

	telemetry *p2p.Telemetry

	//
	//
	//
	dynamic atomic.Bool
}

//
//
type RouterHandler struct {
	p2p.Handler
	Alias string
}

const MaximumReestablishRetreat = 5 * time.Minute

var _ p2p.Router = (*Router)(nil)

var FaultUnservicedNodeLayout = errors.New("REDACTED")

//
func FreshRouter(
	peerDetails p2p.PeerDetails,
	machine *Machine,
	engines []RouterHandler,
	telemetry *p2p.Telemetry,
	tracer log.Tracer,
) (*Router, error) {
	s := &Router{
		peerDetails: peerDetails,

		machine:    machine,
		nodeAssign: FreshNodeAssign(machine, telemetry, tracer),

		telemetry: telemetry,

		dynamic: atomic.Bool{},
	}

	foundation := facility.FreshFoundationFacility(tracer, "REDACTED", s)
	s.FoundationFacility = *foundation

	s.engines = freshHandlerAssign(s)

	for _, record := range engines {
		if err := s.engines.Add(record.Handler, record.Alias); err != nil {
			return nil, errors.Wrapf(err, "REDACTED", record.Alias)
		}
	}

	return s, nil
}

//
//
//

func (s *Router) UponInitiate() error {
	s.Tracer.Details("REDACTED")

	ctx, abort := context.WithTimeout(context.Background(), 10*time.Second)
	defer abort()

	schemeProcessor := func(schemeUUID protocol.ID) {
		s.machine.SetStreamHandler(schemeUUID, s.processInflux)
	}

	//
	err := s.engines.Initiate(schemeProcessor)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	s.machine.AppendNodeBreakdownProcessor(func(id peer.ID, err error) {
		key := nodeUUIDTowardToken(id)
		node := s.nodeAssign.Get(key)
		s.HaltNodeForeachFailure(node, err)
	})

	//
	initiateNodes := s.machine.InitiateNodes()

	s.Tracer.Details("REDACTED", "REDACTED", len(initiateNodes))

	for _, bp := range initiateNodes {
		choices := NodeAppendChoices{
			Secluded:       bp.Secluded,
			Enduring:    bp.Enduring,
			Absolute: bp.Absolute,
			UponPriorInitiate: s.engines.InitializeNode,
			UponSubsequentInitiate:  s.engines.AppendNode,
			UponInitiateUnsuccessful: s.engines.DiscardNode,
		}

		err := s.onboardNode(ctx, bp.LocationDetails, choices)
		if err != nil {
			s.Tracer.Failure("REDACTED", "REDACTED", bp.LocationDetails.String(), "REDACTED", err)
			go s.reestablishNode(bp.LocationDetails, MaximumReestablishRetreat, choices)
			continue
		}
	}

	s.dynamic.Store(true)

	return nil
}

func (s *Router) UponHalt() {
	s.Tracer.Details("REDACTED")

	s.engines.Halt()
	s.nodeAssign.DiscardEvery(NodeDeletionChoices{Rationale: "REDACTED"})

	if err := s.machine.Network().Close(); err != nil {
		s.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	if err := s.machine.Peerstore().Close(); err != nil {
		s.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	s.dynamic.Store(false)
}

func (s *Router) PeerDetails() p2p.PeerDetails {
	return s.peerDetails
}

func (s *Router) Log() log.Tracer {
	return s.Tracer
}

//
//
//

func (s *Router) Handler(alias string) (p2p.Handler, bool) {
	return s.engines.ObtainViaAlias(alias)
}

//
//
func (s *Router) AppendHandler(alias string, handler p2p.Handler) p2p.Handler {
	//
	s.recordUndeveloped("REDACTED")

	return nil
}

func (s *Router) DiscardHandler(_ string, _ p2p.Handler) {
	//
	s.recordUndeveloped("REDACTED")
}

//
//
//

func (s *Router) Nodes() p2p.IDXNodeAssign {
	return s.nodeAssign
}

func (s *Router) CountNodes() (outgoing, incoming, calling int) {
	for _, c := range s.machine.Network().Conns() {
		switch c.Stat().Direction {
		case network.DirInbound:
			incoming++
		case network.DirOutbound:
			outgoing++
		}
	}

	//

	return outgoing, incoming, calling
}

func (s *Router) MaximumCountOutgoingNodes() int {
	//
	s.recordUndeveloped("REDACTED")

	return 0
}

func (s *Router) AppendEnduringNodes(locations []string) error    { return FaultUnservicedNodeLayout }
func (s *Router) AppendSecludedNodeIDXDstore(ids []string) error       { return FaultUnservicedNodeLayout }
func (s *Router) AppendAbsoluteNodeIDXDstore(ids []string) error { return FaultUnservicedNodeLayout }

func (s *Router) CallNodeUsingLocator(_ *p2p.NetworkLocator) error {
	//
	s.recordUndeveloped("REDACTED")

	return nil
}

func (s *Router) CallNodesAsyncronous(nodes []string) error {
	s.recordUndeveloped("REDACTED", "REDACTED", nodes)

	return nil
}

func (s *Router) HaltNodeSmoothly(_ p2p.Node) {
	//
	s.recordUndeveloped("REDACTED")
}

func (s *Router) HaltNodeForeachFailure(node p2p.Node, rationale any) {
	//
	p, ok := node.(*Node)
	if !ok {
		return
	}

	pid := p.ID()

	deletionOptions := NodeDeletionChoices{
		Rationale:      rationale,
		UponSubsequentHalt: s.engines.DiscardNode,
	}

	if err := s.nodeAssign.Discard(pid, deletionOptions); err != nil {
		s.Tracer.Failure("REDACTED", "REDACTED", pid, "REDACTED", err)
		return
	}

	//
	//
	if err := s.machine.Network().ClosePeer(p.locationDetails.ID); err != nil {
		//
		s.Tracer.Failure("REDACTED", "REDACTED", pid, "REDACTED", err)
	}

	//
	mustReestablish := false

	if p.EqualsEnduring() {
		mustReestablish = true
		s.Tracer.Diagnose("REDACTED", "REDACTED", pid, "REDACTED", rationale)
	} else if faultFleeting, ok := FleetingFailureOriginatingSome(rationale); ok {
		mustReestablish = true
		s.Tracer.Diagnose("REDACTED", "REDACTED", pid, "REDACTED", faultFleeting.Err)
	}

	if !mustReestablish {
		return
	}

	go s.reestablishNode(p.LocationDetails(), MaximumReestablishRetreat, NodeAppendChoices{
		Enduring:    p.EqualsEnduring(),
		Absolute: p.EqualsAbsolute(),
		Secluded:       p.EqualsSecluded(),
		UponPriorInitiate: s.engines.InitializeNode,
		UponSubsequentInitiate:  s.engines.AppendNode,
		UponInitiateUnsuccessful: s.engines.DiscardNode,
	})
}

func (s *Router) EqualsCallingEitherCurrentLocator(location *p2p.NetworkLocator) bool {
	s.recordUndeveloped("REDACTED")
	return false
}

func (s *Router) EqualsNodeEnduring(networkLocation *p2p.NetworkLocator) bool {
	p := s.nodeAssign.Get(networkLocation.ID)
	if p == nil {
		return false
	}

	return p.(*Node).EqualsEnduring()
}

func (s *Router) EqualsNodeAbsolute(id p2p.ID) bool {
	p := s.nodeAssign.Get(id)
	if p == nil {
		return false
	}

	return p.(*Node).EqualsAbsolute()
}

func (s *Router) LabelNodeLikeValid(_ p2p.Node) {
	//
	s.recordUndeveloped("REDACTED")
}

//
//
//

func (s *Router) Multicast(e p2p.Wrapper) chan bool {
	s.Tracer.Diagnose("REDACTED", "REDACTED", e.ConduitUUID)

	e.Signal = freshPriorSerializedArtifact(e.Signal)

	var wg sync.WaitGroup
	triumphChn := make(chan bool, s.nodeAssign.Extent())

	s.nodeAssign.ForeachEvery(func(p p2p.Node) {
		wg.Add(1)

		go func(p p2p.Node) {
			defer wg.Done()

			triumph := p.Transmit(e)
			select {
			case triumphChn <- triumph:
			default:
				//
				//
			}
		}(p)
	})

	go func() {
		wg.Wait()
		close(triumphChn)
	}()

	return triumphChn
}

func (s *Router) MulticastAsyncronous(e p2p.Wrapper) {
	s.Tracer.Diagnose("REDACTED", "REDACTED", e.ConduitUUID)

	e.Signal = freshPriorSerializedArtifact(e.Signal)

	s.nodeAssign.ForeachEvery(func(p p2p.Node) {
		go p.Transmit(e)
	})
}

func (s *Router) AttemptMulticast(e p2p.Wrapper) {
	s.Tracer.Diagnose("REDACTED", "REDACTED", e.ConduitUUID)

	e.Signal = freshPriorSerializedArtifact(e.Signal)

	s.nodeAssign.ForeachEvery(func(p p2p.Node) {
		go p.AttemptTransmit(e)
	})
}

func (s *Router) recordUndeveloped(procedure string, kv ...any) {
	s.Tracer.Details(
		"REDACTED",
		append(kv, "REDACTED", procedure)...,
	)
}

func (s *Router) processInflux(influx network.Stream) {
	var (
		nodeUUID     = influx.Conn().RemotePeer()
		schemeUUID = influx.Protocol()
	)

	if !s.equalsDynamic() {
		s.Log().Diagnose(
			"REDACTED",
			"REDACTED", nodeUUID.String(),
			"REDACTED", schemeUUID,
		)
		_ = influx.Reset()
		return
	}

	defer func() {
		if r := recover(); r != nil {
			s.Tracer.Failure(
				"REDACTED",
				"REDACTED", nodeUUID.String(),
				"REDACTED", schemeUUID,
				"REDACTED", r,
				"REDACTED", string(debug.Stack()),
			)
			_ = influx.Reset()
		}
	}()

	//
	schema, handler, err := s.engines.obtainHandlerUsingScheme(schemeUUID)
	if err != nil {
		//
		s.Tracer.Failure("REDACTED", "REDACTED", schemeUUID)
		_ = influx.Reset()
		return
	}

	//
	content, err := InfluxFetchShutdown(influx)
	if err != nil {
		s.Tracer.Failure("REDACTED", "REDACTED", schemeUUID, "REDACTED", err)
		return
	}

	msg, err := decodeSchema(schema.definition, content)
	if err != nil {
		s.Tracer.Failure("REDACTED", "REDACTED", schemeUUID, "REDACTED", err)
		return
	}

	//
	node, err := s.decipherNode(nodeUUID)
	if err != nil {
		s.Tracer.Failure("REDACTED", "REDACTED", nodeUUID.String(), "REDACTED", err)
		return
	}

	var (
		//
		nodeTxt     = nodeUUID.String()
		signalKind = schemaKindAlias(msg)
		contentLength  = float64(len(content))
		tags      = []string{
			"REDACTED", nodeTxt,
			"REDACTED", fmt.Sprintf("REDACTED", schema.definition.ID),
		}
	)

	s.telemetry.NodeAcceptOctetsSum.With(tags...).Add(contentLength)
	s.telemetry.ArtifactAcceptOctetsSum.With("REDACTED", signalKind).Add(contentLength)

	s.Tracer.Diagnose(
		"REDACTED",
		"REDACTED", nodeUUID.String(),
		"REDACTED", schemeUUID,
		"REDACTED", signalKind,
		"REDACTED", contentLength,
	)

	wrapper := p2p.Wrapper{
		Src:       node,
		ConduitUUID: schema.definition.ID,
		Signal:   msg,
	}

	urgency := schema.definition.Urgency

	s.engines.Accept(handler.alias, signalKind, wrapper, urgency)
}

func (s *Router) decipherNode(id peer.ID) (p2p.Node, error) {
	key := nodeUUIDTowardToken(id)

	//
	if node := s.nodeAssign.Get(key); node != nil {
		return node, nil
	}

	locationDetails := s.machine.Peerstore().PeerInfo(id)
	if len(locationDetails.Addrs) == 0 {
		return nil, errors.New("REDACTED")
	}

	//
	choices := NodeAppendChoices{
		Secluded:       false,
		Enduring:    false,
		Absolute: false,
		UponPriorInitiate: s.engines.InitializeNode,
		UponSubsequentInitiate:  s.engines.AppendNode,
		UponInitiateUnsuccessful: s.engines.DiscardNode,
	}

	node, err := s.nodeAssign.Add(locationDetails, choices)
	switch {
	case errors.Is(err, FaultNodePresent):
		//
		if p := s.nodeAssign.Get(key); p != nil {
			return p, nil
		}

		//
		return nil, errors.Wrap(err, "REDACTED")
	case err != nil:
		return nil, errors.Wrap(err, "REDACTED")
	default:
		return node, nil
	}
}

//
func (s *Router) onboardNode(ctx context.Context, locationDetails peer.AddrInfo, choices NodeAppendChoices) error {
	if locationDetails.ID == s.machine.ID() {
		s.Tracer.Details("REDACTED")
		return nil
	}

	pid := locationDetails.ID.String()

	s.Tracer.Details(
		"REDACTED",
		"REDACTED", pid,
		"REDACTED", locationDetails.String(),
		"REDACTED", choices.Enduring,
		"REDACTED", choices.Absolute,
		"REDACTED", choices.Secluded,
	)

	if err := s.machine.Connect(ctx, locationDetails); err != nil {
		return errors.Wrap(err, "REDACTED")
	}

	if _, err := s.nodeAssign.Add(locationDetails, choices); err != nil {
		return errors.Wrap(err, "REDACTED")
	}

	for _, location := range locationDetails.Addrs {
		//
		//
		if EqualsDomainLocation(location) {
			s.machine.Peerstore().AddAddr(locationDetails.ID, location, peerstore.PermanentAddrTTL)
		}
	}

	//
	//
	locators := s.machine.variedLocationTxtViaUUID(locationDetails.ID)

	s.Tracer.Details("REDACTED", "REDACTED", pid, "REDACTED", locators)

	go s.pingNode(locationDetails)

	return nil
}

//
//
func (s *Router) reestablishNode(locationDetails peer.AddrInfo, retreatMaximum time.Duration, choices NodeAppendChoices) {
	defer func() {
		if r := recover(); r != nil {
			s.Tracer.Failure("REDACTED", "REDACTED", r)
		}
	}()

	retreat := 1 * time.Second
	snooze := func() {
		variation := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(retreat + variation)

		retreat *= 2
		if retreatMaximum > 0 && retreat > retreatMaximum {
			retreat = retreatMaximum
		}
	}

	var (
		ctx   = network.WithDialPeerTimeout(context.Background(), 3*time.Second)
		pid   = locationDetails.ID.String()
		initiate = time.Now()
	)

	for {
		if !s.equalsDynamic() {
			return
		}

		s.Tracer.Details(
			"REDACTED",
			"REDACTED", pid,
			"REDACTED", choices.Secluded,
			"REDACTED", choices.Enduring,
			"REDACTED", choices.Absolute,
		)

		//
		if err := s.machine.Connect(ctx, locationDetails); err != nil {
			s.Tracer.Failure(
				"REDACTED",
				"REDACTED", pid,
				"REDACTED", err,
				"REDACTED", retreat.String(),
			)

			snooze()
			continue
		}

		//
		_, err := s.nodeAssign.Add(locationDetails, choices)
		if err != nil && !errors.Is(err, FaultNodePresent) {
			s.Tracer.Failure(
				"REDACTED",
				"REDACTED", pid,
				"REDACTED", err,
				"REDACTED", retreat.String(),
			)
			snooze()
			continue
		}

		var (
			passed   = time.Since(initiate)
			locators = s.machine.variedLocationTxtViaUUID(locationDetails.ID)
		)

		s.Tracer.Details(
			"REDACTED",
			"REDACTED", pid,
			"REDACTED", locators,
			"REDACTED", passed.String(),
		)

		go s.pingNode(locationDetails)

		return
	}
}

//
//
func (s *Router) pingNode(locationDetails peer.AddrInfo) {
	const deadline = 5 * time.Second

	ctx, abort := context.WithTimeout(context.Background(), deadline)
	defer abort()

	var (
		pid       = locationDetails.ID.String()
		locators = s.machine.variedLocationTxtViaUUID(locationDetails.ID)
	)

	rtt, err := s.machine.Ping(ctx, locationDetails)
	if err != nil {
		s.Tracer.Failure("REDACTED", "REDACTED", pid, "REDACTED", locators, "REDACTED", err)
		return
	}

	s.Tracer.Details("REDACTED", "REDACTED", pid, "REDACTED", locators, "REDACTED", rtt.String())
}

func (s *Router) equalsDynamic() bool {
	return s.dynamic.Load()
}
