package pex

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/valkyrieworks/utils/cmap"
	cometmath "github.com/valkyrieworks/utils/math"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/p2p/link"
	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
)

type Node = p2p.Node

const (
	//
	PexConduit = byte(0x00)

	//
	//
	//
	maximumLocationVolume = 256

	//
	//
	maximumMessageVolume = maximumLocationVolume * maximumFetchPreference

	//
	standardAssureNodesDuration = 30 * time.Second

	//

	//
	minimumTimeAmongScans = 2 * time.Minute

	//
	scanNodeDuration = 30 * time.Second

	maximumTriesToCall = 16 //

	//
	//
	//
	tendencyToChooseNewNodes = 30 //

	//
	standardProhibitTime = 24 * time.Hour
)

type errMaximumTriesToCall struct{}

func (e errMaximumTriesToCall) Fault() string {
	return fmt.Sprintf("REDACTED", maximumTriesToCall)
}

type errTooPrematureToCall struct {
	retreatPeriod time.Duration
	finalCalled      time.Time
}

func (e errTooPrematureToCall) Fault() string {
	return fmt.Sprintf(
		"REDACTED",
		e.retreatPeriod, e.finalCalled, time.Since(e.finalCalled))
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
	p2p.RootHandler

	registry              AddressLedger
	settings            *HandlerSettings
	assureNodesChan     chan struct{} //
	assureNodesDuration time.Duration //

	//
	queriesRelayed         *cmap.CIndex //
	finalAcceptedQueries *cmap.CIndex //

	sourceLocations []*p2p.NetLocation

	triesToCall sync.Map //

	//
	scanNodeDetails map[p2p.ID]scanNodeDetails
}

func (r *Handler) minimumAcceptQueryCadence() time.Duration {
	//
	//
	return r.assureNodesDuration / 3
}

//
type HandlerSettings struct {
	//
	OriginStyle bool

	//
	//
	//
	SourceDetachWaitDuration time.Duration

	//
	DurableNodesMaximumCallDuration time.Duration

	//
	//
	Origins []string
}

type _attemptstodial struct {
	amount     int
	finalCalled time.Time
}

//
func NewHandler(b AddressLedger, settings *HandlerSettings) *Handler {
	r := &Handler{
		registry:                 b,
		settings:               settings,
		assureNodesChan:        make(chan struct{}),
		assureNodesDuration:    standardAssureNodesDuration,
		queriesRelayed:         cmap.NewCIndex(),
		finalAcceptedQueries: cmap.NewCIndex(),
		scanNodeDetails:       make(map[p2p.ID]scanNodeDetails),
	}
	r.RootHandler = *p2p.NewRootHandler("REDACTED", r)
	return r
}

//
func (r *Handler) OnBegin() error {
	err := r.registry.Begin()
	if err != nil && err != daemon.ErrYetLaunched {
		return err
	}

	countActive, sourceLocations, err := r.inspectOrigins()
	if err != nil {
		return err
	} else if countActive == 0 && r.registry.Empty() {
		return errors.New("REDACTED")
	}

	r.sourceLocations = sourceLocations

	//
	//
	if r.settings.OriginStyle {
		go r.scanNodesProcedure()
	} else {
		go r.assureNodesProcedure()
	}
	return nil
}

//
func (r *Handler) OnHalt() {
	if err := r.registry.Halt(); err != nil {
		r.Tracer.Fault("REDACTED", "REDACTED", err)
	}
}

//
func (r *Handler) FetchStreams() []*link.StreamDefinition {
	return []*link.StreamDefinition{
		{
			ID:                  PexConduit,
			Urgency:            1,
			TransmitBufferVolume:   10,
			AcceptSignalVolume: maximumMessageVolume,
			SignalKind:         &tmp2p.Signal{},
		},
	}
}

//
//
func (r *Handler) AppendNode(p Node) {
	if p.IsOutgoing() {
		//
		//
		//
		if r.registry.RequireAdditionalLocations() {
			r.QueryLocations(p)
		}
	} else {
		//
		address, err := p.MemberDetails().NetLocation()
		if err != nil {
			r.Tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", p)
			return
		}

		//
		src := address

		//
		//
		err = r.registry.AppendLocation(address, src)
		r.traceErrAddressRegistry(err)
	}
}

//
func (r *Handler) DeleteNode(p Node, _ any) {
	id := string(p.ID())
	r.queriesRelayed.Erase(id)
	r.finalAcceptedQueries.Erase(id)
}

func (r *Handler) traceErrAddressRegistry(err error) {
	if err != nil {
		switch err.(type) {
		case ErrAddressRegistryNullAddress:
			r.Tracer.Fault("REDACTED", "REDACTED", err)
		default:
			//
			r.Tracer.Diagnose("REDACTED", "REDACTED", err)
		}
	}
}

//
func (r *Handler) Accept(e p2p.Packet) {
	r.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.StreamUID, "REDACTED", e.Signal)

	switch msg := e.Signal.(type) {
	case *tmp2p.PexQuery:

		//
		//
		//
		//

		//
		//
		if r.settings.OriginStyle && !e.Src.IsOutgoing() {
			id := string(e.Src.ID())
			v := r.finalAcceptedQueries.Get(id)
			if v != nil {
				//
				//
				return
			}
			r.finalAcceptedQueries.Set(id, time.Now())

			//
			r.TransmitLocations(e.Src, r.registry.FetchPreferenceWithTendency(tendencyToChooseNewNodes))
			go func() {
				//
				e.Src.PurgeHalt()
				r.Router.HaltNodeSmoothly(e.Src)
			}()

		} else {
			//
			if err := r.acceptQuery(e.Src); err != nil {
				r.Router.HaltNodeForFault(e.Src, err)
				r.registry.StampFlawed(e.Src.SocketAddress(), standardProhibitTime)
				return
			}
			r.TransmitLocations(e.Src, r.registry.FetchPreference())
		}

	case *tmp2p.PexLocations:
		//
		locations, err := p2p.NetAddressesFromSchema(msg.Locations)
		if err != nil {
			r.Router.HaltNodeForFault(e.Src, err)
			r.registry.StampFlawed(e.Src.SocketAddress(), standardProhibitTime)
			return
		}
		err = r.AcceptLocations(locations, e.Src)
		if err != nil {
			r.Router.HaltNodeForFault(e.Src, err)
			if err == ErrUninvitedCatalog {
				r.registry.StampFlawed(e.Src.SocketAddress(), standardProhibitTime)
			}
			return
		}

	default:
		r.Tracer.Fault(fmt.Sprintf("REDACTED", msg))
	}
}

//
func (r *Handler) acceptQuery(src Node) error {
	id := string(src.ID())
	v := r.finalAcceptedQueries.Get(id)
	if v == nil {
		//
		finalAccepted := time.Time{}
		r.finalAcceptedQueries.Set(id, finalAccepted)
		return nil
	}

	finalAccepted := v.(time.Time)
	if finalAccepted.Equal(time.Time{}) {
		//
		finalAccepted = time.Now()
		r.finalAcceptedQueries.Set(id, finalAccepted)
		return nil
	}

	now := time.Now()
	minimumCadence := r.minimumAcceptQueryCadence()
	if now.Sub(finalAccepted) < minimumCadence {
		return fmt.Errorf(
			"REDACTED",
			src.ID(),
			finalAccepted,
			now,
			minimumCadence,
		)
	}
	r.finalAcceptedQueries.Set(id, now)
	return nil
}

//
//
func (r *Handler) QueryLocations(p Node) {
	id := string(p.ID())
	if r.queriesRelayed.Has(id) {
		return
	}
	r.Tracer.Diagnose("REDACTED", "REDACTED", p)
	r.queriesRelayed.Set(id, struct{}{})
	p.Transmit(p2p.Packet{
		StreamUID: PexConduit,
		Signal:   &tmp2p.PexQuery{},
	})
}

//
//
//
func (r *Handler) AcceptLocations(locations []*p2p.NetLocation, src Node) error {
	id := string(src.ID())
	if !r.queriesRelayed.Has(id) {
		return ErrUninvitedCatalog
	}
	r.queriesRelayed.Erase(id)

	originAddress, err := src.MemberDetails().NetLocation()
	if err != nil {
		return err
	}

	for _, netAddress := range locations {
		//
		err = r.registry.AppendLocation(netAddress, originAddress)
		if err != nil {
			r.traceErrAddressRegistry(err)
			//
			//
			continue
		}
	}

	//
	for _, sourceAddress := range r.sourceLocations {
		if sourceAddress.Matches(originAddress) {
			select {
			case r.assureNodesChan <- struct{}{}:
			default:
			}
			break
		}
	}

	return nil
}

//
func (r *Handler) TransmitLocations(p Node, netLocations []*p2p.NetLocation) {
	e := p2p.Packet{
		StreamUID: PexConduit,
		Signal:   &tmp2p.PexLocations{Locations: p2p.NetAddressesToSchema(netLocations)},
	}
	p.Transmit(e)
}

//
func (r *Handler) CollectionAssureNodesDuration(d time.Duration) {
	r.assureNodesDuration = d
}

//
func (r *Handler) assureNodesProcedure() {
	var (
		origin   = engineseed.NewRandom()
		variance = origin.Int64count(r.assureNodesDuration.Nanoseconds())
	)

	//
	//
	//
	if r.memberHasSomeNodesOrCallingAny() {
		time.Sleep(time.Duration(variance))
	}

	//
	//
	r.assureNodes(true)

	//
	timer := time.NewTicker(r.assureNodesDuration)
	for {
		select {
		case <-timer.C:
			r.assureNodes(true)
		case <-r.assureNodesChan:
			r.assureNodes(false)
		case <-r.Exit():
			timer.Stop()
			return
		}
	}
}

//
//
//
//
//
func (r *Handler) assureNodes(assureNodesDurationPassed bool) {
	var (
		out, in, call = r.Router.CountNodes()
		countToCall     = r.Router.MaximumCountOutgoingNodes() - (out + call)
	)
	r.Tracer.Details(
		"REDACTED",
		"REDACTED", out,
		"REDACTED", in,
		"REDACTED", call,
		"REDACTED", countToCall,
	)

	if countToCall <= 0 {
		return
	}

	//
	//
	//
	newTendency := cometmath.MinimumInteger(out, 8)*10 + 10

	toCall := make(map[p2p.ID]*p2p.NetLocation)
	//
	maximumTries := countToCall * 3

	for i := 0; i < maximumTries && len(toCall) < countToCall; i++ {
		try := r.registry.SelectLocation(newTendency)
		if try == nil {
			continue
		}
		if _, chosen := toCall[try.ID]; chosen {
			continue
		}
		if r.Router.IsCallingOrCurrentLocation(try) {
			continue
		}
		//
		//
		//
		toCall[try.ID] = try
	}

	//
	for _, address := range toCall {
		go func(address *p2p.NetLocation) {
			err := r.callNode(address)
			if err != nil {
				switch err.(type) {
				case errMaximumTriesToCall, errTooPrematureToCall:
					r.Tracer.Diagnose(err.Error(), "REDACTED", address)
				default:
					r.Tracer.Diagnose(err.Error(), "REDACTED", address)
				}
			}
		}(address)
	}

	if r.registry.RequireAdditionalLocations() {
		//
		r.registry.RestoreFlawedNodes()
	}

	if r.registry.RequireAdditionalLocations() {

		//
		node := r.Router.Nodes().Arbitrary()
		if node != nil {
			r.Tracer.Details("REDACTED", "REDACTED", node)
			r.QueryLocations(node)
		}

		//
		//
		//
		if len(toCall) == 0 {
			r.Tracer.Details("REDACTED")
			r.callOrigins()
		}
	}
}

func (r *Handler) callTriesDetails(address *p2p.NetLocation) (tries int, finalCalled time.Time) {
	_tries, ok := r.triesToCall.Load(address.CallString())
	if !ok {
		return
	}
	atd := _tries.(_attemptstodial)
	return atd.amount, atd.finalCalled
}

func (r *Handler) callNode(address *p2p.NetLocation) error {
	tries, finalCalled := r.callTriesDetails(address)
	if !r.Router.IsNodeDurable(address) && tries > maximumTriesToCall {
		r.registry.StampFlawed(address, standardProhibitTime)
		return errMaximumTriesToCall{}
	}

	//
	if tries > 0 {
		variance := time.Duration(engineseed.Float64() * float64(time.Second)) //
		retreatPeriod := variance + ((1 << uint(tries)) * time.Second)
		retreatPeriod = r.maximumRetreatPeriodForNode(address, retreatPeriod)
		sinceFinalCalled := time.Since(finalCalled)
		if sinceFinalCalled < retreatPeriod {
			return errTooPrematureToCall{retreatPeriod, finalCalled}
		}
	}

	err := r.Router.CallNodeWithLocation(address)
	if err != nil {
		if _, ok := err.(p2p.ErrPresentlyCallingOrPresentLocation); ok {
			return err
		}

		stampAddressInRegistryRootedOnErr(address, r.registry, err)
		switch err.(type) {
		case p2p.ErrRouterAuthorizationBreakdown:
			//
			r.triesToCall.Delete(address.CallString())
		default:
			r.triesToCall.Store(address.CallString(), _attemptstodial{tries + 1, time.Now()})
		}
		return fmt.Errorf("REDACTED", tries+1, err)
	}

	//
	r.triesToCall.Delete(address.CallString())
	return nil
}

//
func (r *Handler) maximumRetreatPeriodForNode(address *p2p.NetLocation, scheduled time.Duration) time.Duration {
	if r.settings.DurableNodesMaximumCallDuration > 0 &&
		scheduled > r.settings.DurableNodesMaximumCallDuration &&
		r.Router.IsNodeDurable(address) {
		return r.settings.DurableNodesMaximumCallDuration
	}
	return scheduled
}

//
//
//
//
//
func (r *Handler) inspectOrigins() (countActive int, netLocations []*p2p.NetLocation, err error) {
	lOrigins := len(r.settings.Origins)
	if lOrigins == 0 {
		return -1, nil, nil
	}
	netLocations, faults := p2p.NewNetLocationStrings(r.settings.Origins)
	countActive = lOrigins - len(faults)
	for _, err := range faults {
		switch e := err.(type) {
		case p2p.ErrNetLocationSearch:
			r.Tracer.Fault("REDACTED", "REDACTED", e)
		default:
			return 0, nil, fmt.Errorf("REDACTED", e)
		}
	}
	return countActive, netLocations, nil
}

//
func (r *Handler) callOrigins() {
	mode := engineseed.Mode(len(r.sourceLocations))
	//
	for _, i := range mode {
		//
		sourceAddress := r.sourceLocations[i]
		err := r.Router.CallNodeWithLocation(sourceAddress)

		switch err.(type) {
		case nil, p2p.ErrPresentlyCallingOrPresentLocation:
			return
		}
		r.Tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", sourceAddress)
	}
	//
	if len(r.sourceLocations) > 0 {
		r.Tracer.Fault("REDACTED")
	}
}

//
//
func (r *Handler) TriesToCall(address *p2p.NetLocation) int {
	lTries, endeavored := r.triesToCall.Load(address.CallString())
	if endeavored {
		return lTries.(_attemptstodial).amount
	}
	return 0
}

//

//
//
//
func (r *Handler) scanNodesProcedure() {
	//
	if len(r.sourceLocations) > 0 {
		r.callOrigins()
	} else {
		//
		r.scanNodes(r.registry.FetchPreference())
	}

	//
	timer := time.NewTicker(scanNodeDuration)

	for {
		select {
		case <-timer.C:
			r.endeavorUnlinks()
			r.scanNodes(r.registry.FetchPreference())
			r.sanitizeScanNodeDetails()
		case <-r.Exit():
			return
		}
	}
}

//
//
func (r *Handler) memberHasSomeNodesOrCallingAny() bool {
	out, in, call := r.Router.CountNodes()
	return out+in+call > 0
}

//
//
type scanNodeDetails struct {
	Address *p2p.NetLocation `json:"address"`
	//
	FinalScanned time.Time `json:"final_scanned"`
}

//
func (r *Handler) scanNodes(locations []*p2p.NetLocation) {
	now := time.Now()

	for _, address := range locations {
		nodeDetails, ok := r.scanNodeDetails[address.ID]

		//
		if ok && now.Sub(nodeDetails.FinalScanned) < minimumTimeAmongScans {
			continue
		}

		//
		r.scanNodeDetails[address.ID] = scanNodeDetails{
			Address:        address,
			FinalScanned: now,
		}

		err := r.callNode(address)
		if err != nil {
			switch err.(type) {
			case errMaximumTriesToCall, errTooPrematureToCall, p2p.ErrPresentlyCallingOrPresentLocation:
				r.Tracer.Diagnose(err.Error(), "REDACTED", address)
			default:
				r.Tracer.Diagnose(err.Error(), "REDACTED", address)
			}
			continue
		}

		node := r.Router.Nodes().Get(address.ID)
		if node != nil {
			r.QueryLocations(node)
		}
	}
}

func (r *Handler) sanitizeScanNodeDetails() {
	for id, details := range r.scanNodeDetails {
		//
		//
		//
		//
		//
		//
		if time.Since(details.FinalScanned) > 24*time.Hour {
			delete(r.scanNodeDetails, id)
		}
	}
}

//
func (r *Handler) endeavorUnlinks() {
	for _, node := range r.Router.Nodes().Clone() {
		if node.Status().Period < r.settings.SourceDetachWaitDuration {
			continue
		}
		if node.IsDurable() {
			continue
		}
		r.Router.HaltNodeSmoothly(node)
	}
}

func stampAddressInRegistryRootedOnErr(address *p2p.NetLocation, registry AddressLedger, err error) {
	//
	switch err.(type) {
	case p2p.ErrRouterAuthorizationBreakdown:
		registry.StampFlawed(address, standardProhibitTime)
	default:
		registry.StampEndeavor(address)
	}
}
