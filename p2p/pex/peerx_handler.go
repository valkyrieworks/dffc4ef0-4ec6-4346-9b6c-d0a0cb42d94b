package pex

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/componentindex"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
)

type Node = p2p.Node

const (
	//
	PeerxConduit = byte(0x00)

	//
	//
	//
	maximumLocatorExtent = 256

	//
	//
	maximumSignalExtent = maximumLocatorExtent * maximumFetchPreference

	//
	fallbackAssureNodesSpan = 30 * time.Second

	//

	//
	minimumMomentAmongExplores = 2 * time.Minute

	//
	exploreNodeSpan = 30 * time.Second

	maximumEndeavorsTowardCall = 16 //

	//
	//
	//
	tendencyTowardPreferFreshNodes = 30 //

	//
	fallbackProhibitMoment = 24 * time.Hour
)

type faultMaximumEndeavorsTowardCall struct{}

func (e faultMaximumEndeavorsTowardCall) Failure() string {
	return fmt.Sprintf("REDACTED", maximumEndeavorsTowardCall)
}

type faultExcessivelyPrematureTowardCall struct {
	retreatInterval time.Duration
	finalCalled      time.Time
}

func (e faultExcessivelyPrematureTowardCall) Failure() string {
	return fmt.Sprintf(
		"REDACTED",
		e.retreatInterval, e.finalCalled, time.Since(e.finalCalled))
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
type Handler struct {
	p2p.FoundationHandler

	register              LocationRegister
	settings            *HandlerSettings
	assureNodesChnl     chan struct{} //
	assureNodesSpan time.Duration //

	//
	solicitsRelayed         *componentindex.CNIndex //
	finalAcceptedSolicits *componentindex.CNIndex //

	germLocations []*p2p.NetworkLocator

	endeavorsTowardCall sync.Map //

	//
	exploreNodeInsights map[p2p.ID]exploreNodeDetails
}

func (r *Handler) minimumAcceptSolicitDuration() time.Duration {
	//
	//
	return r.assureNodesSpan / 3
}

//
type HandlerSettings struct {
	//
	OriginStyle bool

	//
	//
	//
	GermDetachPauseSpan time.Duration

	//
	EnduringNodesMaximumCallSpan time.Duration

	//
	//
	Origins []string
}

type _attemptscustody struct {
	numeral     int
	finalCalled time.Time
}

//
func FreshHandler(b LocationRegister, settings *HandlerSettings) *Handler {
	r := &Handler{
		register:                 b,
		settings:               settings,
		assureNodesChnl:        make(chan struct{}),
		assureNodesSpan:    fallbackAssureNodesSpan,
		solicitsRelayed:         componentindex.FreshCNIndex(),
		finalAcceptedSolicits: componentindex.FreshCNIndex(),
		exploreNodeInsights:       make(map[p2p.ID]exploreNodeDetails),
	}
	r.FoundationHandler = *p2p.FreshFoundationHandler("REDACTED", r)
	return r
}

//
func (r *Handler) UponInitiate() error {
	err := r.register.Initiate()
	if err != nil && err != facility.FaultEarlierInitiated {
		return err
	}

	countLive, germLocations, err := r.inspectOrigins()
	if err != nil {
		return err
	} else if countLive == 0 && r.register.Blank() {
		return errors.New("REDACTED")
	}

	r.germLocations = germLocations

	//
	//
	if r.settings.OriginStyle {
		go r.exploreNodesProcedure()
	} else {
		go r.assureNodesProcedure()
	}
	return nil
}

//
func (r *Handler) UponHalt() {
	if err := r.register.Halt(); err != nil {
		r.Tracer.Failure("REDACTED", "REDACTED", err)
	}
}

//
func (r *Handler) ObtainConduits() []*link.ConduitDefinition {
	return []*link.ConduitDefinition{
		{
			ID:                  PeerxConduit,
			Urgency:            1,
			TransmitStagingVolume:   10,
			ObtainSignalVolume: maximumSignalExtent,
			SignalKind:         &tmpfabric.Signal{},
		},
	}
}

//
//
func (r *Handler) AppendNode(p Node) {
	if p.EqualsOutgoing() {
		//
		//
		//
		if r.register.RequireExtraLocations() {
			r.SolicitLocations(p)
		}
	} else {
		//
		location, err := p.PeerDetails().NetworkLocator()
		if err != nil {
			r.Tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", p)
			return
		}

		//
		src := location

		//
		//
		err = r.register.AppendLocator(location, src)
		r.reportFaultLocationRegister(err)
	}
}

//
func (r *Handler) DiscardNode(p Node, _ any) {
	id := string(p.ID())
	r.solicitsRelayed.Erase(id)
	r.finalAcceptedSolicits.Erase(id)
}

func (r *Handler) reportFaultLocationRegister(err error) {
	if err != nil {
		switch err.(type) {
		case FaultLocationRegisterVoidLocation:
			r.Tracer.Failure("REDACTED", "REDACTED", err)
		default:
			//
			r.Tracer.Diagnose("REDACTED", "REDACTED", err)
		}
	}
}

//
func (r *Handler) Accept(e p2p.Wrapper) {
	r.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.ConduitUUID, "REDACTED", e.Signal)

	switch msg := e.Signal.(type) {
	case *tmpfabric.PeerxSolicit:

		//
		//
		//
		//

		//
		//
		if r.settings.OriginStyle && !e.Src.EqualsOutgoing() {
			id := string(e.Src.ID())
			v := r.finalAcceptedSolicits.Get(id)
			if v != nil {
				//
				//
				return
			}
			r.finalAcceptedSolicits.Set(id, time.Now())

			//
			r.TransmitLocations(e.Src, r.register.FetchPreferenceUsingTendency(tendencyTowardPreferFreshNodes))
			go func() {
				//
				e.Src.PurgeHalt()
				r.Router.HaltNodeSmoothly(e.Src)
			}()

		} else {
			//
			if err := r.acceptSolicit(e.Src); err != nil {
				r.Router.HaltNodeForeachFailure(e.Src, err)
				r.register.LabelFlawed(e.Src.PortLocation(), fallbackProhibitMoment)
				return
			}
			r.TransmitLocations(e.Src, r.register.FetchPreference())
		}

	case *tmpfabric.PeerxLocations:
		//
		locations, err := p2p.NetworkLocatorsOriginatingSchema(msg.Locations)
		if err != nil {
			r.Router.HaltNodeForeachFailure(e.Src, err)
			r.register.LabelFlawed(e.Src.PortLocation(), fallbackProhibitMoment)
			return
		}
		err = r.AcceptLocations(locations, e.Src)
		if err != nil {
			r.Router.HaltNodeForeachFailure(e.Src, err)
			if err == FaultUnpromptedCatalog {
				r.register.LabelFlawed(e.Src.PortLocation(), fallbackProhibitMoment)
			}
			return
		}

	default:
		r.Tracer.Failure(fmt.Sprintf("REDACTED", msg))
	}
}

//
func (r *Handler) acceptSolicit(src Node) error {
	id := string(src.ID())
	v := r.finalAcceptedSolicits.Get(id)
	if v == nil {
		//
		finalAccepted := time.Time{}
		r.finalAcceptedSolicits.Set(id, finalAccepted)
		return nil
	}

	finalAccepted := v.(time.Time)
	if finalAccepted.Equal(time.Time{}) {
		//
		finalAccepted = time.Now()
		r.finalAcceptedSolicits.Set(id, finalAccepted)
		return nil
	}

	now := time.Now()
	minimumDuration := r.minimumAcceptSolicitDuration()
	if now.Sub(finalAccepted) < minimumDuration {
		return fmt.Errorf(
			"REDACTED",
			src.ID(),
			finalAccepted,
			now,
			minimumDuration,
		)
	}
	r.finalAcceptedSolicits.Set(id, now)
	return nil
}

//
//
func (r *Handler) SolicitLocations(p Node) {
	id := string(p.ID())
	if r.solicitsRelayed.Has(id) {
		return
	}
	r.Tracer.Diagnose("REDACTED", "REDACTED", p)
	r.solicitsRelayed.Set(id, struct{}{})
	p.Transmit(p2p.Wrapper{
		ConduitUUID: PeerxConduit,
		Signal:   &tmpfabric.PeerxSolicit{},
	})
}

//
//
//
func (r *Handler) AcceptLocations(locations []*p2p.NetworkLocator, src Node) error {
	id := string(src.ID())
	if !r.solicitsRelayed.Has(id) {
		return FaultUnpromptedCatalog
	}
	r.solicitsRelayed.Erase(id)

	originLocation, err := src.PeerDetails().NetworkLocator()
	if err != nil {
		return err
	}

	for _, networkLocation := range locations {
		//
		err = r.register.AppendLocator(networkLocation, originLocation)
		if err != nil {
			r.reportFaultLocationRegister(err)
			//
			//
			continue
		}
	}

	//
	for _, germLocation := range r.germLocations {
		if germLocation.Matches(originLocation) {
			select {
			case r.assureNodesChnl <- struct{}{}:
			default:
			}
			break
		}
	}

	return nil
}

//
func (r *Handler) TransmitLocations(p Node, networkLocations []*p2p.NetworkLocator) {
	e := p2p.Wrapper{
		ConduitUUID: PeerxConduit,
		Signal:   &tmpfabric.PeerxLocations{Locations: p2p.NetworkLocatorsTowardSchema(networkLocations)},
	}
	p.Transmit(e)
}

//
func (r *Handler) AssignAssureNodesSpan(d time.Duration) {
	r.assureNodesSpan = d
}

//
func (r *Handler) assureNodesProcedure() {
	var (
		germ   = commitrand.FreshArbitrary()
		variation = germ.Int63num(r.assureNodesSpan.Nanoseconds())
	)

	//
	//
	//
	if r.peerOwnsFewNodesEitherCallingSome() {
		time.Sleep(time.Duration(variation))
	}

	//
	//
	r.assureNodes(true)

	//
	metronome := time.NewTicker(r.assureNodesSpan)
	for {
		select {
		case <-metronome.C:
			r.assureNodes(true)
		case <-r.assureNodesChnl:
			r.assureNodes(false)
		case <-r.Exit():
			metronome.Stop()
			return
		}
	}
}

//
//
//
//
//
func (r *Handler) assureNodes(assureNodesSpanPassed bool) {
	var (
		out, in, call = r.Router.CountNodes()
		countTowardCall     = r.Router.MaximumCountOutgoingNodes() - (out + call)
	)
	r.Tracer.Details(
		"REDACTED",
		"REDACTED", out,
		"REDACTED", in,
		"REDACTED", call,
		"REDACTED", countTowardCall,
	)

	if countTowardCall <= 0 {
		return
	}

	//
	//
	//
	freshTendency := strongarithmetic.MinimumInteger(out, 8)*10 + 10

	towardCall := make(map[p2p.ID]*p2p.NetworkLocator)
	//
	maximumEndeavors := countTowardCall * 3

	for i := 0; i < maximumEndeavors && len(towardCall) < countTowardCall; i++ {
		try := r.register.SelectLocator(freshTendency)
		if try == nil {
			continue
		}
		if _, preferred := towardCall[try.ID]; preferred {
			continue
		}
		if r.Router.EqualsCallingEitherCurrentLocator(try) {
			continue
		}
		//
		//
		//
		towardCall[try.ID] = try
	}

	//
	for _, location := range towardCall {
		go func(location *p2p.NetworkLocator) {
			err := r.callNode(location)
			if err != nil {
				switch err.(type) {
				case faultMaximumEndeavorsTowardCall, faultExcessivelyPrematureTowardCall:
					r.Tracer.Diagnose(err.Error(), "REDACTED", location)
				default:
					r.Tracer.Diagnose(err.Error(), "REDACTED", location)
				}
			}
		}(location)
	}

	if r.register.RequireExtraLocations() {
		//
		r.register.RestoreFlawedNodes()
	}

	if r.register.RequireExtraLocations() {

		//
		node := r.Router.Nodes().Unpredictable()
		if node != nil {
			r.Tracer.Details("REDACTED", "REDACTED", node)
			r.SolicitLocations(node)
		}

		//
		//
		//
		if len(towardCall) == 0 {
			r.Tracer.Details("REDACTED")
			r.callOrigins()
		}
	}
}

func (r *Handler) callEndeavorsDetails(location *p2p.NetworkLocator) (endeavors int, finalCalled time.Time) {
	_endeavors, ok := r.endeavorsTowardCall.Load(location.CallText())
	if !ok {
		return
	}
	atd := _endeavors.(_attemptscustody)
	return atd.numeral, atd.finalCalled
}

func (r *Handler) callNode(location *p2p.NetworkLocator) error {
	endeavors, finalCalled := r.callEndeavorsDetails(location)
	if !r.Router.EqualsNodeEnduring(location) && endeavors > maximumEndeavorsTowardCall {
		r.register.LabelFlawed(location, fallbackProhibitMoment)
		return faultMaximumEndeavorsTowardCall{}
	}

	//
	if endeavors > 0 {
		variation := time.Duration(commitrand.Float64() * float64(time.Second)) //
		retreatInterval := variation + ((1 << uint(endeavors)) * time.Second)
		retreatInterval = r.maximumRetreatIntervalForeachNode(location, retreatInterval)
		becauseFinalCalled := time.Since(finalCalled)
		if becauseFinalCalled < retreatInterval {
			return faultExcessivelyPrematureTowardCall{retreatInterval, finalCalled}
		}
	}

	err := r.Router.CallNodeUsingLocator(location)
	if err != nil {
		if _, ok := err.(p2p.FaultPresentlyCallingEitherPresentLocator); ok {
			return err
		}

		labelLocationInsideRegisterRootedUponFault(location, r.register, err)
		switch err.(type) {
		case p2p.FaultRouterAuthorizationBreakdown:
			//
			r.endeavorsTowardCall.Delete(location.CallText())
		default:
			r.endeavorsTowardCall.Store(location.CallText(), _attemptscustody{endeavors + 1, time.Now()})
		}
		return fmt.Errorf("REDACTED", endeavors+1, err)
	}

	//
	r.endeavorsTowardCall.Delete(location.CallText())
	return nil
}

//
func (r *Handler) maximumRetreatIntervalForeachNode(location *p2p.NetworkLocator, scheduled time.Duration) time.Duration {
	if r.settings.EnduringNodesMaximumCallSpan > 0 &&
		scheduled > r.settings.EnduringNodesMaximumCallSpan &&
		r.Router.EqualsNodeEnduring(location) {
		return r.settings.EnduringNodesMaximumCallSpan
	}
	return scheduled
}

//
//
//
//
//
func (r *Handler) inspectOrigins() (countLive int, networkLocations []*p2p.NetworkLocator, err error) {
	lnOrigins := len(r.settings.Origins)
	if lnOrigins == 0 {
		return -1, nil, nil
	}
	networkLocations, errors := p2p.FreshNetworkLocatorTexts(r.settings.Origins)
	countLive = lnOrigins - len(errors)
	for _, err := range errors {
		switch e := err.(type) {
		case p2p.FaultNetworkLocatorSearch:
			r.Tracer.Failure("REDACTED", "REDACTED", e)
		default:
			return 0, nil, fmt.Errorf("REDACTED", e)
		}
	}
	return countLive, networkLocations, nil
}

//
func (r *Handler) callOrigins() {
	mode := commitrand.Mode(len(r.germLocations))
	//
	for _, i := range mode {
		//
		germLocation := r.germLocations[i]
		err := r.Router.CallNodeUsingLocator(germLocation)

		switch err.(type) {
		case nil, p2p.FaultPresentlyCallingEitherPresentLocator:
			return
		}
		r.Tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", germLocation)
	}
	//
	if len(r.germLocations) > 0 {
		r.Tracer.Failure("REDACTED")
	}
}

//
//
func (r *Handler) EndeavorsTowardCall(location *p2p.NetworkLocator) int {
	lnEndeavors, endeavored := r.endeavorsTowardCall.Load(location.CallText())
	if endeavored {
		return lnEndeavors.(_attemptscustody).numeral
	}
	return 0
}

//

//
//
//
func (r *Handler) exploreNodesProcedure() {
	//
	if len(r.germLocations) > 0 {
		r.callOrigins()
	} else {
		//
		r.exploreNodes(r.register.FetchPreference())
	}

	//
	metronome := time.NewTicker(exploreNodeSpan)

	for {
		select {
		case <-metronome.C:
			r.effortUnlinks()
			r.exploreNodes(r.register.FetchPreference())
			r.sanitizeExploreNodeInsights()
		case <-r.Exit():
			return
		}
	}
}

//
//
func (r *Handler) peerOwnsFewNodesEitherCallingSome() bool {
	out, in, call := r.Router.CountNodes()
	return out+in+call > 0
}

//
//
type exploreNodeDetails struct {
	Location *p2p.NetworkLocator `json:"location"`
	//
	FinalExplored time.Time `json:"final_explored"`
}

//
func (r *Handler) exploreNodes(locations []*p2p.NetworkLocator) {
	now := time.Now()

	for _, location := range locations {
		nodeDetails, ok := r.exploreNodeInsights[location.ID]

		//
		if ok && now.Sub(nodeDetails.FinalExplored) < minimumMomentAmongExplores {
			continue
		}

		//
		r.exploreNodeInsights[location.ID] = exploreNodeDetails{
			Location:        location,
			FinalExplored: now,
		}

		err := r.callNode(location)
		if err != nil {
			switch err.(type) {
			case faultMaximumEndeavorsTowardCall, faultExcessivelyPrematureTowardCall, p2p.FaultPresentlyCallingEitherPresentLocator:
				r.Tracer.Diagnose(err.Error(), "REDACTED", location)
			default:
				r.Tracer.Diagnose(err.Error(), "REDACTED", location)
			}
			continue
		}

		node := r.Router.Nodes().Get(location.ID)
		if node != nil {
			r.SolicitLocations(node)
		}
	}
}

func (r *Handler) sanitizeExploreNodeInsights() {
	for id, details := range r.exploreNodeInsights {
		//
		//
		//
		//
		//
		//
		if time.Since(details.FinalExplored) > 24*time.Hour {
			delete(r.exploreNodeInsights, id)
		}
	}
}

//
func (r *Handler) effortUnlinks() {
	for _, node := range r.Router.Nodes().Duplicate() {
		if node.Condition().Interval < r.settings.GermDetachPauseSpan {
			continue
		}
		if node.EqualsEnduring() {
			continue
		}
		r.Router.HaltNodeSmoothly(node)
	}
}

func labelLocationInsideRegisterRootedUponFault(location *p2p.NetworkLocator, register LocationRegister, err error) {
	//
	switch err.(type) {
	case p2p.FaultRouterAuthorizationBreakdown:
		register.LabelFlawed(location, fallbackProhibitMoment)
	default:
		register.LabelEffort(location)
	}
}
