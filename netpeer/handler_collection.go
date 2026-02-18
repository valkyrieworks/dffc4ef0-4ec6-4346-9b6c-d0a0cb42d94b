package netpeer

import (
	"fmt"
	"time"

	"github.com/valkyrieworks/intrinsic/autosource"
	"github.com/valkyrieworks/p2p"
	"github.com/libp2p/go-libp2p/core/protocol"
)

//
type handlerCollection struct {
	routerReference *Router

	handlers []handlerElement

	//
	handlerLabels map[string]int

	//
	protocols map[protocol.ID]handlerProtocol
}

//
type handlerElement struct {
	p2p.Handler
	label          string
	subscriberBuffer *autosource.Depository[awaitingPacket]
}

//
type handlerProtocol struct {
	handlerUID  int
	definition *p2p.StreamDefinition
}

//
type awaitingPacket struct {
	p2p.Packet
	signalKind string
	appendedAt     time.Time
}

func newHandlerCollection(routerReference *Router) *handlerCollection {
	return &handlerCollection{
		routerReference: routerReference,

		handlers:     []handlerElement{},
		handlerLabels: make(map[string]int),
		protocols:    make(map[protocol.ID]handlerProtocol),
	}
}

//
//
func (rs *handlerCollection) Add(handler p2p.Handler, label string) error {
	followingUID := len(rs.handlers)

	if _, ok := rs.handlerLabels[label]; ok {
		return fmt.Errorf("REDACTED", label)
	}

	//
	for i := range handler.FetchStreams() {
		var (
			conduitDefinition = handler.FetchStreams()[i]
			protocolUID        = ProtocolUID(conduitDefinition.ID)
		)

		if _, ok := rs.protocols[protocolUID]; ok {
			return fmt.Errorf("REDACTED", protocolUID)
		}

		rs.protocols[protocolUID] = handlerProtocol{
			handlerUID:  followingUID,
			definition: conduitDefinition,
		}
	}

	rs.handlers = append(rs.handlers, handlerElement{
		Handler:       handler,
		label:          label,
		subscriberBuffer: rs.newHandlerUrgencyBuffer(followingUID, label),
	})

	//
	rs.handlerLabels[label] = followingUID

	rs.routerReference.Tracer.Details("REDACTED", "REDACTED", label)

	return nil
}

//
func (rs *handlerCollection) Begin(eachProtocolResponse func(protocol.ID)) error {
	for _, handler := range rs.handlers {
		rs.routerReference.Tracer.Details("REDACTED", "REDACTED", handler.label)
		handler.CollectionRouter(rs.routerReference)

		if err := handler.Begin(); err != nil {
			return fmt.Errorf("REDACTED", handler.label, err)
		}

		handler.subscriberBuffer.Begin()
	}

	for protocolUID := range rs.protocols {
		eachProtocolResponse(protocolUID)
	}

	return nil
}

func (rs *handlerCollection) Halt() {
	for _, handler := range rs.handlers {
		handler.subscriberBuffer.Halt()

		rs.routerReference.Tracer.Details("REDACTED", "REDACTED", handler.label)
		if err := handler.Halt(); err != nil {
			rs.routerReference.Tracer.Fault("REDACTED", "REDACTED", handler.label, "REDACTED", err)
		}
	}
}

func (rs *handlerCollection) InitNode(node *Node) {
	for _, handler := range rs.handlers {
		handler.InitNode(node)
	}
}

func (rs *handlerCollection) AppendNode(node *Node) {
	for _, handler := range rs.handlers {
		handler.AppendNode(node)
	}
}

func (rs *handlerCollection) DeleteNode(node *Node, cause any) {
	for _, handler := range rs.handlers {
		handler.DeleteNode(node, cause)
	}
}

func (rs *handlerCollection) FetchByLabel(label string) (p2p.Handler, bool) {
	idx, ok := rs.handlerLabels[label]
	if !ok {
		return nil, false
	}

	return rs.handlers[idx].Handler, true
}

func (rs *handlerCollection) fetchHandlerWithProtocol(id protocol.ID) (handlerProtocol, handlerElement, error) {
	protocol, ok := rs.protocols[id]
	if !ok {
		return handlerProtocol{}, handlerElement{}, fmt.Errorf("REDACTED")
	}

	return protocol, rs.handlers[protocol.handlerUID], nil
}

//
//
//
//
//
//
//
//
//
func (rs *handlerCollection) Accept(handlerLabel, signalKind string, packet p2p.Packet, urgency int) {
	idx, ok := rs.handlerLabels[handlerLabel]
	if !ok {
		rs.routerReference.Tracer.Fault("REDACTED", "REDACTED", handlerLabel)
		return
	}

	handler := rs.handlers[idx]

	tags := []string{
		"REDACTED", handlerLabel,
		"REDACTED", signalKind,
	}

	//
	rs.routerReference.stats.SignalsAccepted.With(tags...).Add(1)
	rs.routerReference.stats.SignalsHandlerInJourney.With(tags...).Add(1)
	now := time.Now()

	pq := awaitingPacket{
		Packet:    packet,
		signalKind: signalKind,
		appendedAt:     now,
	}

	err := handler.subscriberBuffer.PropelUrgency(pq, urgency)
	if err != nil {
		rs.routerReference.stats.SignalsHandlerInJourney.With(tags...).Add(-1)
		rs.routerReference.Tracer.Fault("REDACTED", "REDACTED", handlerLabel, "REDACTED", err)
	}

	rs.routerReference.Tracer.Diagnose(
		"REDACTED",
		"REDACTED", handlerLabel,
		"REDACTED", signalKind,
	)
}

func (rs *handlerCollection) acceptBuffered(handlerUID int, e awaitingPacket) {
	handler := rs.handlers[handlerUID]

	tags := []string{
		"REDACTED", handler.label,
		"REDACTED", e.signalKind,
	}

	//
	awaitingFor := time.Since(e.appendedAt)
	rs.routerReference.stats.SignalsHandlerAwaitingPeriod.With(tags...).Observe(awaitingFor.Seconds())

	//
	//
	//
	traceGradual := (awaitingFor > time.Second) &&
		(awaitingFor > 10*time.Second || e.appendedAt.UnixMilli()%10 == 0)

	if traceGradual {
		rs.routerReference.Tracer.Details(
			"REDACTED",
			"REDACTED", handler.label,
			"REDACTED", e.signalKind,
			"REDACTED", awaitingFor.String(),
		)
	}

	now := time.Now()

	handler.Accept(e.Packet)

	timeSeized := time.Since(now)

	rs.routerReference.stats.SignalsHandlerInJourney.With(tags...).Add(-1)
	rs.routerReference.stats.SignalHandlerAcceptPeriod.With(tags...).Observe(timeSeized.Seconds())
}

//
//
//
func (rs *handlerCollection) newHandlerUrgencyBuffer(handlerUID int, handlerLabel string) *autosource.Depository[awaitingPacket] {
	const (
		//
		urgencies = 10

		//
		//
		parallelDepositoryVolume = 512
	)

	parallelismTally := rs.
		routerReference.stats.SignalHandlerBufferParallelism.
		With("REDACTED", handlerLabel)

	return autosource.New(
		rs.newHandlerAdjuster(handlerLabel),
		func(e awaitingPacket) {
			rs.acceptBuffered(handlerUID, e)
		},
		parallelDepositoryVolume,
		autosource.WithTracer[awaitingPacket](rs.routerReference.Tracer),
		autosource.WithOnRatio[awaitingPacket](func() {
			parallelismTally.Add(1)
		}),
		autosource.WithOnReduce[awaitingPacket](func() {
			parallelismTally.Add(-1)
		}),
		autosource.WithUrgencyBuffer[awaitingPacket](autosource.NewUrgencyBuffer(urgencies)),
	)
}

func (rs *handlerCollection) newHandlerAdjuster(handlerLabel string) *autosource.VelocityWaitperiodAdjuster {
	const (
		standardMinimumOperators       = 4
		standardMaximumOperators       = 32
		standardWaitperiodLimit = 100 * time.Millisecond
		waitperiodQuantile       = 90.0 //

		//
		automaticRatioRecurrence = 200 * time.Millisecond
	)

	var (
		maximumOperators       = standardMaximumOperators
		minimumOperators       = standardMinimumOperators
		waitperiodLimit = automaticRatioRecurrence
	)

	//
	if handlerLabel == "REDACTED" {
		minimumOperators = 8
		maximumOperators = 512
		waitperiodLimit = 500 * time.Millisecond
	}

	return autosource.NewVelocityWaitperiodAdjuster(
		minimumOperators,
		maximumOperators,
		waitperiodQuantile,
		waitperiodLimit,
		automaticRatioRecurrence,
		rs.routerReference.Tracer,
	)
}
