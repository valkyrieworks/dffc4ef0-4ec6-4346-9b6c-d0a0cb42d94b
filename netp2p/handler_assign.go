package netp2p

import (
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/selfpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/libp2p/go-libp2p/core/protocol"
)

//
type handlerAssign struct {
	routerReference *Router

	engines []handlerElement

	//
	handlerIdentifiers map[string]int

	//
	schemes map[protocol.ID]handlerScheme
}

//
type handlerElement struct {
	p2p.Handler
	alias          string
	subscriberStaging *selfpool.Hub[awaitingWrapper]
}

//
type handlerScheme struct {
	handlerUUID  int
	definition *p2p.ConduitDefinition
}

//
type awaitingWrapper struct {
	p2p.Wrapper
	signalKind string
	appendedLocated     time.Time
}

func freshHandlerAssign(routerReference *Router) *handlerAssign {
	return &handlerAssign{
		routerReference: routerReference,

		engines:     []handlerElement{},
		handlerIdentifiers: make(map[string]int),
		schemes:    make(map[protocol.ID]handlerScheme),
	}
}

//
//
func (rs *handlerAssign) Add(handler p2p.Handler, alias string) error {
	followingUUID := len(rs.engines)

	if _, ok := rs.handlerIdentifiers[alias]; ok {
		return fmt.Errorf("REDACTED", alias)
	}

	//
	for i := range handler.ObtainConduits() {
		var (
			conduitDefinition = handler.ObtainConduits()[i]
			schemeUUID        = SchemeUUID(conduitDefinition.ID)
		)

		if _, ok := rs.schemes[schemeUUID]; ok {
			return fmt.Errorf("REDACTED", schemeUUID)
		}

		rs.schemes[schemeUUID] = handlerScheme{
			handlerUUID:  followingUUID,
			definition: conduitDefinition,
		}
	}

	rs.engines = append(rs.engines, handlerElement{
		Handler:       handler,
		alias:          alias,
		subscriberStaging: rs.freshHandlerUrgencyStaging(followingUUID, alias),
	})

	//
	rs.handlerIdentifiers[alias] = followingUUID

	rs.routerReference.Tracer.Details("REDACTED", "REDACTED", alias)

	return nil
}

//
func (rs *handlerAssign) Initiate(everySchemeReact func(protocol.ID)) error {
	for _, handler := range rs.engines {
		rs.routerReference.Tracer.Details("REDACTED", "REDACTED", handler.alias)
		handler.AssignRouter(rs.routerReference)

		if err := handler.Initiate(); err != nil {
			return fmt.Errorf("REDACTED", handler.alias, err)
		}

		handler.subscriberStaging.Initiate()
	}

	for schemeUUID := range rs.schemes {
		everySchemeReact(schemeUUID)
	}

	return nil
}

func (rs *handlerAssign) Halt() {
	for _, handler := range rs.engines {
		handler.subscriberStaging.Halt()

		rs.routerReference.Tracer.Details("REDACTED", "REDACTED", handler.alias)
		if err := handler.Halt(); err != nil {
			rs.routerReference.Tracer.Failure("REDACTED", "REDACTED", handler.alias, "REDACTED", err)
		}
	}
}

func (rs *handlerAssign) InitializeNode(node *Node) {
	for _, handler := range rs.engines {
		handler.InitializeNode(node)
	}
}

func (rs *handlerAssign) AppendNode(node *Node) {
	for _, handler := range rs.engines {
		handler.AppendNode(node)
	}
}

func (rs *handlerAssign) DiscardNode(node *Node, rationale any) {
	for _, handler := range rs.engines {
		handler.DiscardNode(node, rationale)
	}
}

func (rs *handlerAssign) ObtainViaAlias(alias string) (p2p.Handler, bool) {
	idx, ok := rs.handlerIdentifiers[alias]
	if !ok {
		return nil, false
	}

	return rs.engines[idx].Handler, true
}

func (rs *handlerAssign) obtainHandlerUsingScheme(id protocol.ID) (handlerScheme, handlerElement, error) {
	scheme, ok := rs.schemes[id]
	if !ok {
		return handlerScheme{}, handlerElement{}, fmt.Errorf("REDACTED")
	}

	return scheme, rs.engines[protocol.handlerUUID], nil
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
func (rs *handlerAssign) Accept(handlerAlias, signalKind string, wrapper p2p.Wrapper, urgency int) {
	idx, ok := rs.handlerIdentifiers[handlerAlias]
	if !ok {
		rs.routerReference.Tracer.Failure("REDACTED", "REDACTED", handlerAlias)
		return
	}

	handler := rs.engines[idx]

	tags := []string{
		"REDACTED", handlerAlias,
		"REDACTED", signalKind,
	}

	//
	rs.routerReference.telemetry.SignalsAccepted.With(tags...).Add(1)
	rs.routerReference.telemetry.SignalsHandlerInsideAirtime.With(tags...).Add(1)
	now := time.Now()

	pq := awaitingWrapper{
		Wrapper:    wrapper,
		signalKind: signalKind,
		appendedLocated:     now,
	}

	err := handler.subscriberStaging.PropelUrgency(pq, urgency)
	if err != nil {
		rs.routerReference.telemetry.SignalsHandlerInsideAirtime.With(tags...).Add(-1)
		rs.routerReference.Tracer.Failure("REDACTED", "REDACTED", handlerAlias, "REDACTED", err)
	}

	rs.routerReference.Tracer.Diagnose(
		"REDACTED",
		"REDACTED", handlerAlias,
		"REDACTED", signalKind,
	)
}

func (rs *handlerAssign) acceptStaged(handlerUUID int, e awaitingWrapper) {
	handler := rs.engines[handlerUUID]

	tags := []string{
		"REDACTED", handler.alias,
		"REDACTED", e.signalKind,
	}

	//
	awaitingForeach := time.Since(e.appendedLocated)
	rs.routerReference.telemetry.SignalsHandlerAwaitingInterval.With(tags...).Observe(awaitingForeach.Seconds())

	//
	//
	//
	reportGradual := (awaitingForeach > time.Second) &&
		(awaitingForeach > 10*time.Second || e.appendedLocated.UnixMilli()%10 == 0)

	if reportGradual {
		rs.routerReference.Tracer.Details(
			"REDACTED",
			"REDACTED", handler.alias,
			"REDACTED", e.signalKind,
			"REDACTED", awaitingForeach.String(),
		)
	}

	now := time.Now()

	handler.Accept(e.Wrapper)

	momentSeized := time.Since(now)

	rs.routerReference.telemetry.SignalsHandlerInsideAirtime.With(tags...).Add(-1)
	rs.routerReference.telemetry.ArtifactHandlerAcceptInterval.With(tags...).Observe(momentSeized.Seconds())
}

//
//
//
func (rs *handlerAssign) freshHandlerUrgencyStaging(handlerUUID int, handlerAlias string) *selfpool.Hub[awaitingWrapper] {
	const (
		//
		urgencies = 10

		//
		//
		parallelHubVolume = 512
	)

	parallelismTally := rs.
		routerReference.telemetry.ArtifactHandlerStagingParallelism.
		With("REDACTED", handlerAlias)

	return selfpool.New(
		rs.freshHandlerAmplifier(handlerAlias),
		func(e awaitingWrapper) {
			rs.acceptStaged(handlerUUID, e)
		},
		parallelHubVolume,
		selfpool.UsingTracer[awaitingWrapper](rs.routerReference.Tracer),
		selfpool.UsingUponAmplify[awaitingWrapper](func() {
			parallelismTally.Add(1)
		}),
		selfpool.UsingUponReduce[awaitingWrapper](func() {
			parallelismTally.Add(-1)
		}),
		selfpool.UsingUrgencyStaging[awaitingWrapper](selfpool.FreshUrgencyStaging(urgencies)),
	)
}

func (rs *handlerAssign) freshHandlerAmplifier(handlerAlias string) *selfpool.YieldWaitstateAmplifier {
	const (
		fallbackMinimumLaborers       = 4
		fallbackMaximumLaborers       = 32
		fallbackWaitstateLimit = 100 * time.Millisecond
		waitstateQuantile       = 90.0 //

		//
		automaticAmplifyRecurrence = 200 * time.Millisecond
	)

	var (
		maximumLaborers       = fallbackMaximumLaborers
		minimumLaborers       = fallbackMinimumLaborers
		waitstateLimit = automaticAmplifyRecurrence
	)

	//
	if handlerAlias == "REDACTED" {
		minimumLaborers = 8
		maximumLaborers = 512
		waitstateLimit = 500 * time.Millisecond
	}

	return selfpool.FreshYieldWaitstateAmplifier(
		minimumLaborers,
		maximumLaborers,
		waitstateQuantile,
		waitstateLimit,
		automaticAmplifyRecurrence,
		rs.routerReference.Tracer,
	)
}
