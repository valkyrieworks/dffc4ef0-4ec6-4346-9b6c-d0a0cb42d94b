package netpeer

import (
	"context"
	"fmt"
	"math/rand"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/p2p"
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
	daemon.RootDaemon

	memberDetails p2p.MemberDetails //

	machine    *Machine
	nodeCollection *NodeCollection

	handlers *handlerCollection

	stats *p2p.Stats

	//
	//
	//
	enabled atomic.Bool
}

//
//
type RouterHandler struct {
	p2p.Handler
	Label string
}

const MaximumReestablishRetreat = 5 * time.Minute

var _ p2p.Toggeler = (*Router)(nil)

var ErrUnacceptedNodeLayout = errors.New("REDACTED")

//
func NewRouter(
	memberDetails p2p.MemberDetails,
	machine *Machine,
	handlers []RouterHandler,
	stats *p2p.Stats,
	tracer log.Tracer,
) (*Router, error) {
	s := &Router{
		memberDetails: memberDetails,

		machine:    machine,
		nodeCollection: NewNodeCollection(machine, stats, tracer),

		stats: stats,

		enabled: atomic.Bool{},
	}

	root := daemon.NewRootDaemon(tracer, "REDACTED", s)
	s.RootDaemon = *root

	s.handlers = newHandlerCollection(s)

	for _, item := range handlers {
		if err := s.handlers.Add(item.Handler, item.Label); err != nil {
			return nil, errors.Wrapf(err, "REDACTED", item.Label)
		}
	}

	return s, nil
}

//
//
//

func (s *Router) OnBegin() error {
	s.Tracer.Details("REDACTED")

	ctx, revoke := context.WithTimeout(context.Background(), 10*time.Second)
	defer revoke()

	protocolManager := func(protocolUID protocol.ID) {
		s.machine.SetStreamHandler(protocolUID, s.processInflux)
	}

	//
	err := s.handlers.Begin(protocolManager)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	s.machine.AppendNodeBreakdownManager(func(id peer.ID, err error) {
		key := nodeUIDToKey(id)
		node := s.nodeCollection.Get(key)
		s.HaltNodeForFault(node, err)
	})

	//
	onboardNodes := s.machine.OnboardNodes()

	s.Tracer.Details("REDACTED", "REDACTED", len(onboardNodes))

	for _, bp := range onboardNodes {
		opts := NodeAppendSettings{
			Internal:       bp.Internal,
			Durable:    bp.Durable,
			Absolute: bp.Absolute,
			OnPriorBegin: s.handlers.InitNode,
			OnAfterBegin:  s.handlers.AppendNode,
			OnBeginErrored: s.handlers.DeleteNode,
		}

		err := s.onboardNode(ctx, bp.AddressDetails, opts)
		if err != nil {
			s.Tracer.Fault("REDACTED", "REDACTED", bp.AddressDetails.String(), "REDACTED", err)
			go s.reestablishNode(bp.AddressDetails, MaximumReestablishRetreat, opts)
			continue
		}
	}

	s.enabled.Store(true)

	return nil
}

func (s *Router) OnHalt() {
	s.Tracer.Details("REDACTED")

	s.handlers.Halt()
	s.nodeCollection.DeleteAll(NodeDeletionSettings{Cause: "REDACTED"})

	if err := s.machine.Network().Close(); err != nil {
		s.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	if err := s.machine.Peerstore().Close(); err != nil {
		s.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	s.enabled.Store(false)
}

func (s *Router) MemberDetails() p2p.MemberDetails {
	return s.memberDetails
}

func (s *Router) Log() log.Tracer {
	return s.Tracer
}

//
//
//

func (s *Router) Handler(label string) (p2p.Handler, bool) {
	return s.handlers.FetchByLabel(label)
}

//
//
func (s *Router) AppendHandler(label string, handler p2p.Handler) p2p.Handler {
	//
	s.traceUnexecuted("REDACTED")

	return nil
}

func (s *Router) DeleteHandler(_ string, _ p2p.Handler) {
	//
	s.traceUnexecuted("REDACTED")
}

//
//
//

func (s *Router) Nodes() p2p.IDXNodeCollection {
	return s.nodeCollection
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
	s.traceUnexecuted("REDACTED")

	return 0
}

func (s *Router) AppendDurableNodes(locations []string) error    { return ErrUnacceptedNodeLayout }
func (s *Router) AppendInternalNodeIDXDatastore(ids []string) error       { return ErrUnacceptedNodeLayout }
func (s *Router) AppendAbsoluteNodeIDXDatastore(ids []string) error { return ErrUnacceptedNodeLayout }

func (s *Router) CallNodeWithLocation(_ *p2p.NetLocation) error {
	//
	s.traceUnexecuted("REDACTED")

	return nil
}

func (s *Router) CallNodesAsync(nodes []string) error {
	s.traceUnexecuted("REDACTED", "REDACTED", nodes)

	return nil
}

func (s *Router) HaltNodeSmoothly(_ p2p.Node) {
	//
	s.traceUnexecuted("REDACTED")
}

func (s *Router) HaltNodeForFault(node p2p.Node, cause any) {
	//
	p, ok := node.(*Node)
	if !ok {
		return
	}

	pid := p.ID()

	deletionOpts := NodeDeletionSettings{
		Cause:      cause,
		OnAfterHalt: s.handlers.DeleteNode,
	}

	if err := s.nodeCollection.Delete(pid, deletionOpts); err != nil {
		s.Tracer.Fault("REDACTED", "REDACTED", pid, "REDACTED", err)
		return
	}

	//
	//
	if err := s.machine.Network().ClosePeer(p.addressDetails.ID); err != nil {
		//
		s.Tracer.Fault("REDACTED", "REDACTED", pid, "REDACTED", err)
	}

	//
	mustReestablish := false

	if p.IsDurable() {
		mustReestablish = true
		s.Tracer.Diagnose("REDACTED", "REDACTED", pid, "REDACTED", cause)
	} else if errTemporary, ok := TemporaryFaultFromAny(cause); ok {
		mustReestablish = true
		s.Tracer.Diagnose("REDACTED", "REDACTED", pid, "REDACTED", errTemporary.Err)
	}

	if !mustReestablish {
		return
	}

	go s.reestablishNode(p.AddressDetails(), MaximumReestablishRetreat, NodeAppendSettings{
		Durable:    p.IsDurable(),
		Absolute: p.IsAbsolute(),
		Internal:       p.IsInternal(),
		OnPriorBegin: s.handlers.InitNode,
		OnAfterBegin:  s.handlers.AppendNode,
		OnBeginErrored: s.handlers.DeleteNode,
	})
}

func (s *Router) IsCallingOrCurrentLocation(address *p2p.NetLocation) bool {
	s.traceUnexecuted("REDACTED")
	return false
}

func (s *Router) IsNodeDurable(netAddress *p2p.NetLocation) bool {
	p := s.nodeCollection.Get(netAddress.ID)
	if p == nil {
		return false
	}

	return p.(*Node).IsDurable()
}

func (s *Router) IsNodeAbsolute(id p2p.ID) bool {
	p := s.nodeCollection.Get(id)
	if p == nil {
		return false
	}

	return p.(*Node).IsAbsolute()
}

func (s *Router) StampNodeAsSound(_ p2p.Node) {
	//
	s.traceUnexecuted("REDACTED")
}

//
//
//

func (s *Router) Multicast(e p2p.Packet) chan bool {
	s.Tracer.Diagnose("REDACTED", "REDACTED", e.StreamUID)

	e.Signal = newPreSerializedSignal(e.Signal)

	var wg sync.WaitGroup
	successChannel := make(chan bool, s.nodeCollection.Volume())

	s.nodeCollection.ForEach(func(p p2p.Node) {
		wg.Add(1)

		go func(p p2p.Node) {
			defer wg.Done()

			success := p.Transmit(e)
			select {
			case successChannel <- success:
			default:
				//
				//
			}
		}(p)
	})

	go func() {
		wg.Wait()
		close(successChannel)
	}()

	return successChannel
}

func (s *Router) MulticastAsync(e p2p.Packet) {
	s.Tracer.Diagnose("REDACTED", "REDACTED", e.StreamUID)

	e.Signal = newPreSerializedSignal(e.Signal)

	s.nodeCollection.ForEach(func(p p2p.Node) {
		go p.Transmit(e)
	})
}

func (s *Router) AttemptMulticast(e p2p.Packet) {
	s.Tracer.Diagnose("REDACTED", "REDACTED", e.StreamUID)

	e.Signal = newPreSerializedSignal(e.Signal)

	s.nodeCollection.ForEach(func(p p2p.Node) {
		go p.AttemptTransmit(e)
	})
}

func (s *Router) traceUnexecuted(procedure string, kv ...any) {
	s.Tracer.Details(
		"REDACTED",
		append(kv, "REDACTED", procedure)...,
	)
}

func (s *Router) processInflux(influx network.Stream) {
	var (
		nodeUID     = influx.Conn().RemotePeer()
		protocolUID = influx.Protocol()
	)

	if !s.isEnabled() {
		s.Log().Diagnose(
			"REDACTED",
			"REDACTED", nodeUID.String(),
			"REDACTED", protocolUID,
		)
		_ = influx.Reset()
		return
	}

	defer func() {
		if r := recover(); r != nil {
			s.Tracer.Fault(
				"REDACTED",
				"REDACTED", nodeUID.String(),
				"REDACTED", protocolUID,
				"REDACTED", r,
				"REDACTED", string(debug.Stack()),
			)
			_ = influx.Reset()
		}
	}()

	//
	schema, handler, err := s.handlers.fetchHandlerWithProtocol(protocolUID)
	if err != nil {
		//
		s.Tracer.Fault("REDACTED", "REDACTED", protocolUID)
		_ = influx.Reset()
		return
	}

	//
	shipment, err := InfluxFetchEnd(influx)
	if err != nil {
		s.Tracer.Fault("REDACTED", "REDACTED", protocolUID, "REDACTED", err)
		return
	}

	msg, err := unserializeSchema(schema.definition, shipment)
	if err != nil {
		s.Tracer.Fault("REDACTED", "REDACTED", protocolUID, "REDACTED", err)
		return
	}

	//
	node, err := s.decipherNode(nodeUID)
	if err != nil {
		s.Tracer.Fault("REDACTED", "REDACTED", nodeUID.String(), "REDACTED", err)
		return
	}

	var (
		//
		nodeStr     = nodeUID.String()
		signalKind = schemaKindLabel(msg)
		shipmentSize  = float64(len(shipment))
		tags      = []string{
			"REDACTED", nodeStr,
			"REDACTED", fmt.Sprintf("REDACTED", schema.definition.ID),
		}
	)

	s.stats.NodeAcceptOctetsSum.With(tags...).Add(shipmentSize)
	s.stats.SignalAcceptOctetsSum.With("REDACTED", signalKind).Add(shipmentSize)

	s.Tracer.Diagnose(
		"REDACTED",
		"REDACTED", nodeUID.String(),
		"REDACTED", protocolUID,
		"REDACTED", signalKind,
		"REDACTED", shipmentSize,
	)

	packet := p2p.Packet{
		Src:       node,
		StreamUID: schema.definition.ID,
		Signal:   msg,
	}

	urgency := schema.definition.Urgency

	s.handlers.Accept(handler.label, signalKind, packet, urgency)
}

func (s *Router) decipherNode(id peer.ID) (p2p.Node, error) {
	key := nodeUIDToKey(id)

	//
	if node := s.nodeCollection.Get(key); node != nil {
		return node, nil
	}

	addressDetails := s.machine.Peerstore().PeerInfo(id)
	if len(addressDetails.Addrs) == 0 {
		return nil, errors.New("REDACTED")
	}

	//
	opts := NodeAppendSettings{
		Internal:       false,
		Durable:    false,
		Absolute: false,
		OnPriorBegin: s.handlers.InitNode,
		OnAfterBegin:  s.handlers.AppendNode,
		OnBeginErrored: s.handlers.DeleteNode,
	}

	node, err := s.nodeCollection.Add(addressDetails, opts)
	switch {
	case errors.Is(err, ErrNodePresent):
		//
		if p := s.nodeCollection.Get(key); p != nil {
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
func (s *Router) onboardNode(ctx context.Context, addressDetails peer.AddrInfo, opts NodeAppendSettings) error {
	if addressDetails.ID == s.machine.ID() {
		s.Tracer.Details("REDACTED")
		return nil
	}

	pid := addressDetails.ID.String()

	s.Tracer.Details(
		"REDACTED",
		"REDACTED", pid,
		"REDACTED", addressDetails.String(),
		"REDACTED", opts.Durable,
		"REDACTED", opts.Absolute,
		"REDACTED", opts.Internal,
	)

	if err := s.machine.Connect(ctx, addressDetails); err != nil {
		return errors.Wrap(err, "REDACTED")
	}

	if _, err := s.nodeCollection.Add(addressDetails, opts); err != nil {
		return errors.Wrap(err, "REDACTED")
	}

	for _, address := range addressDetails.Addrs {
		//
		//
		if IsDNSAddress(address) {
			s.machine.Peerstore().AddAddr(addressDetails.ID, address, peerstore.PermanentAddrTTL)
		}
	}

	//
	//
	locations := s.machine.multipleAddressStrByUID(addressDetails.ID)

	s.Tracer.Details("REDACTED", "REDACTED", pid, "REDACTED", locations)

	go s.pingNode(addressDetails)

	return nil
}

//
//
func (s *Router) reestablishNode(addressDetails peer.AddrInfo, retreatMaximum time.Duration, opts NodeAppendSettings) {
	defer func() {
		if r := recover(); r != nil {
			s.Tracer.Fault("REDACTED", "REDACTED", r)
		}
	}()

	retreat := 1 * time.Second
	pause := func() {
		variance := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(retreat + variance)

		retreat *= 2
		if retreatMaximum > 0 && retreat > retreatMaximum {
			retreat = retreatMaximum
		}
	}

	var (
		ctx   = network.WithDialPeerTimeout(context.Background(), 3*time.Second)
		pid   = addressDetails.ID.String()
		begin = time.Now()
	)

	for {
		if !s.isEnabled() {
			return
		}

		s.Tracer.Details(
			"REDACTED",
			"REDACTED", pid,
			"REDACTED", opts.Internal,
			"REDACTED", opts.Durable,
			"REDACTED", opts.Absolute,
		)

		//
		if err := s.machine.Connect(ctx, addressDetails); err != nil {
			s.Tracer.Fault(
				"REDACTED",
				"REDACTED", pid,
				"REDACTED", err,
				"REDACTED", retreat.String(),
			)

			pause()
			continue
		}

		//
		_, err := s.nodeCollection.Add(addressDetails, opts)
		if err != nil && !errors.Is(err, ErrNodePresent) {
			s.Tracer.Fault(
				"REDACTED",
				"REDACTED", pid,
				"REDACTED", err,
				"REDACTED", retreat.String(),
			)
			pause()
			continue
		}

		var (
			passed   = time.Since(begin)
			locations = s.machine.multipleAddressStrByUID(addressDetails.ID)
		)

		s.Tracer.Details(
			"REDACTED",
			"REDACTED", pid,
			"REDACTED", locations,
			"REDACTED", passed.String(),
		)

		go s.pingNode(addressDetails)

		return
	}
}

//
//
func (s *Router) pingNode(addressDetails peer.AddrInfo) {
	const deadline = 5 * time.Second

	ctx, revoke := context.WithTimeout(context.Background(), deadline)
	defer revoke()

	var (
		pid       = addressDetails.ID.String()
		locations = s.machine.multipleAddressStrByUID(addressDetails.ID)
	)

	rtt, err := s.machine.Ping(ctx, addressDetails)
	if err != nil {
		s.Tracer.Fault("REDACTED", "REDACTED", pid, "REDACTED", locations, "REDACTED", err)
		return
	}

	s.Tracer.Details("REDACTED", "REDACTED", pid, "REDACTED", locations, "REDACTED", rtt.String())
}

func (s *Router) isEnabled() bool {
	return s.enabled.Load()
}
