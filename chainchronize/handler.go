package chainchronize

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/netp2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	chainchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/chainchronize"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
const ChainchronizeConduit = byte(0x40)

const (
	fallbackDurationConditionRevise      = 10 * time.Second
	aggregateStyleIntrinsicConditionRevise = 1 * time.Second

	fallbackDurationRouterTowardAgreement = 1 * time.Second

	//
	durationAttemptChronize = 10 * time.Millisecond
)

type agreementHandler interface {
	//
	//
	RouterTowardAgreement(status sm.Status, omitJournal bool)
}

type txpoolHandler interface {
	//
	ActivateInsideOutputTrans()
}

type nodeFailure struct {
	err    error
	nodeUUID p2p.ID
}

func (e nodeFailure) Failure() string {
	return fmt.Sprintf("REDACTED", e.nodeUUID, e.err.Error())
}

//
type Handler struct {
	p2p.FoundationHandler

	//
	primaryStatus sm.Status

	//
	//
	activated *atomic.Bool

	//
	//
	//
	aggregateStyleActivated bool

	ledgerExecute     *sm.LedgerHandler
	depot         sm.LedgerDepot
	hub          *LedgerHub
	regionalLocation     security.Location
	hubProcedureGroup sync.WaitGroup

	solicitsStream <-chan LedgerSolicit
	faultsStream   <-chan nodeFailure

	//
	durationRouterTowardAgreement time.Duration

	//
	durationConditionRevise time.Duration

	telemetry *Telemetry
}

//
func FreshHandler(
	activated bool,
	aggregateStyle bool,
	status sm.Status,
	ledgerExecute *sm.LedgerHandler,
	depot *depot.LedgerDepot,
	regionalLocation security.Location,
	inactiveStatusChronizeAltitude int64,
	telemetry *Telemetry,
) *Handler {
	depotAltitude := depot.Altitude()
	if depotAltitude == 0 {
		//
		//
		//
		//
		//
		//
		depotAltitude = inactiveStatusChronizeAltitude
	}

	if status.FinalLedgerAltitude != depotAltitude {
		panic(fmt.Sprintf(
			"REDACTED",
			status.FinalLedgerAltitude,
			depotAltitude,
		))
	}

	//
	//
	solicitsStream := make(chan LedgerSolicit)

	const volume = 1000                      //
	faultsStream := make(chan nodeFailure, volume) //

	initiateAltitude := depotAltitude + 1
	if initiateAltitude == 1 {
		initiateAltitude = status.PrimaryAltitude
	}
	hub := FreshLedgerHub(initiateAltitude, solicitsStream, faultsStream)

	activatedMarker := &atomic.Bool{}
	activatedMarker.Store(activated)

	durationConditionRevise := fallbackDurationConditionRevise
	if aggregateStyle {
		durationConditionRevise = aggregateStyleIntrinsicConditionRevise
	}

	r := &Handler{
		primaryStatus:              status,
		ledgerExecute:                 ledgerExecute,
		depot:                     depot,
		hub:                      hub,
		activated:                   activatedMarker,
		aggregateStyleActivated:       aggregateStyle,
		regionalLocation:                 regionalLocation,
		solicitsStream:                solicitsStream,
		faultsStream:                  faultsStream,
		telemetry:                   telemetry,
		durationRouterTowardAgreement: fallbackDurationRouterTowardAgreement,
		durationConditionRevise:      durationConditionRevise,
	}

	r.FoundationHandler = *p2p.FreshFoundationHandler("REDACTED", r)

	return r
}

//
func (r *Handler) AssignTracer(l log.Tracer) {
	r.Tracer = l
	r.hub.Tracer = l
}

//
func (r *Handler) UponInitiate() error {
	//
	if !r.activated.Load() {
		return nil
	}

	return r.executeHub(false)
}

func (r *Handler) executeHub(statusChronized bool) error {
	if err := r.hub.Initiate(); err != nil {
		return err
	}

	//
	run := func(fn func()) {
		r.hubProcedureGroup.Add(1)
		go func() {
			defer r.hubProcedureGroup.Done()
			fn()
		}()
	}

	//
	run(func() {
		metronome := time.NewTicker(r.durationConditionRevise)
		defer metronome.Stop()
		r.hubSignalsProcedure(metronome)
	})

	if r.aggregateStyleActivated {
		//
		ledgerAbsorber, err := r.obtainLedgerAbsorber()
		if err != nil {
			return err
		}

		run(func() {
			r.ledgerAbsorberProcedure(ledgerAbsorber)
		})

		return nil
	}

	//
	run(func() { r.hubProcedure(statusChronized) })

	return nil
}

//
func (r *Handler) Activate(status sm.Status) error {
	if !r.activated.CompareAndSwap(false, true) {
		return FaultEarlierActivated
	}

	r.primaryStatus = status
	r.hub.altitude = status.FinalLedgerAltitude + 1

	return r.executeHub(true)
}

//
func (r *Handler) UponHalt() {
	if !r.activated.Load() {
		return
	}

	if err := r.hub.Halt(); err != nil {
		r.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	r.hubProcedureGroup.Wait()
}

//
func (r *Handler) ObtainConduits() []*p2p.ConduitDefinition {
	return []*p2p.ConduitDefinition{
		{
			ID:                  ChainchronizeConduit,
			Urgency:            5,
			TransmitStagingVolume:   1000,
			ObtainReserveVolume:  50 * 4096,
			ObtainSignalVolume: MaximumSignalExtent,
			SignalKind:         &chainchema.Signal{},
		},
	}
}

//
func (r *Handler) AppendNode(node p2p.Node) {
	node.Transmit(p2p.Wrapper{
		ConduitUUID: ChainchronizeConduit,
		Signal: &chainchema.ConditionReply{
			Foundation:   r.depot.Foundation(),
			Altitude: r.depot.Altitude(),
		},
	})
	//

	//
	//
}

//
func (r *Handler) DiscardNode(node p2p.Node, _ any) {
	r.hub.DiscardNode(node.ID())
}

//
//
func (r *Handler) replyTowardNode(msg *chainchema.LedgerSolicit, src p2p.Node) {
	ledger := r.depot.FetchLedger(msg.Altitude)
	if ledger == nil {
		r.Tracer.Details("REDACTED", "REDACTED", src, "REDACTED", msg.Altitude)
		src.AttemptTransmit(p2p.Wrapper{
			ConduitUUID: ChainchronizeConduit,
			Signal:   &chainchema.NegativeLedgerReply{Altitude: msg.Altitude},
		})

		return
	}

	status, err := r.ledgerExecute.Depot().Fetch()
	if err != nil {
		r.Tracer.Failure("REDACTED", "REDACTED", err)
		return
	}

	var addnEndorse *kinds.ExpandedEndorse
	if status.AgreementSettings.Iface.BallotAdditionsActivated(msg.Altitude) {
		addnEndorse = r.depot.FetchLedgerExpandedEndorse(msg.Altitude)
		if addnEndorse == nil {
			r.Tracer.Failure("REDACTED", "REDACTED", ledger)
			return
		}
	}

	bl, err := ledger.TowardSchema()
	if err != nil {
		r.Tracer.Failure("REDACTED", "REDACTED", err)
		return
	}

	src.AttemptTransmit(p2p.Wrapper{
		ConduitUUID: ChainchronizeConduit,
		Signal: &chainchema.LedgerReply{
			Ledger:     bl,
			AddnEndorse: addnEndorse.TowardSchema(),
		},
	})
}

func (r *Handler) processNodeReply(msg *chainchema.LedgerReply, src p2p.Node) {
	bi, err := kinds.LedgerOriginatingSchema(msg.Ledger)
	if err != nil {
		r.Tracer.Failure("REDACTED", "REDACTED", src, "REDACTED", msg, "REDACTED", err)
		r.haltNodeForeachFailure(src, err)
		return
	}

	var addnEndorse *kinds.ExpandedEndorse
	if msg.AddnEndorse != nil {
		addnEndorse, err = kinds.ExpandedEndorseOriginatingSchema(msg.AddnEndorse)
		if err != nil {
			r.Tracer.Failure("REDACTED", "REDACTED", src, "REDACTED", err)
			r.haltNodeForeachFailure(src, err)
			return
		}
	}

	if err := r.hub.AppendLedger(src.ID(), bi, addnEndorse, msg.Ledger.Extent()); err != nil {
		r.Tracer.Failure("REDACTED", "REDACTED", src, "REDACTED", err)
	}
}

//
func (r *Handler) Accept(e p2p.Wrapper) {
	if err := CertifySignal(e.Signal); err != nil {
		r.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", e.Signal, "REDACTED", err)
		r.haltNodeForeachFailure(e.Src, err)
		return
	}

	r.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.ConduitUUID, "REDACTED", e.Signal)

	switch msg := e.Signal.(type) {
	case *chainchema.LedgerSolicit:
		//
		r.replyTowardNode(msg, e.Src)
	case *chainchema.LedgerReply:
		//
		go r.processNodeReply(msg, e.Src)
	case *chainchema.ConditionSolicit:
		//
		e.Src.AttemptTransmit(p2p.Wrapper{
			ConduitUUID: ChainchronizeConduit,
			Signal: &chainchema.ConditionReply{
				Altitude: r.depot.Altitude(),
				Foundation:   r.depot.Foundation(),
			},
		})
	case *chainchema.ConditionReply:
		//
		r.hub.AssignNodeScope(e.Src.ID(), msg.Foundation, msg.Altitude)
	case *chainchema.NegativeLedgerReply:
		r.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", msg.Altitude)
		r.hub.ReiterateSolicitOriginating(msg.Altitude, e.Src.ID())
	default:
		r.Tracer.Failure(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
	}
}

func (r *Handler) regionalPeerLedgersThatSuccession(status sm.Status) bool {
	_, val := status.Assessors.ObtainViaLocation(r.regionalLocation)
	if val == nil {
		return false
	}
	sum := status.Assessors.SumBallotingPotency()
	return val.BallotingPotency >= sum/3
}

//
//
func (r *Handler) hubProcedure(statusChronized bool) {
	r.Tracer.Details("REDACTED", "REDACTED", statusChronized)

	r.telemetry.Chronizing.Set(1)
	defer r.telemetry.Chronizing.Set(0)

	attemptChronizeMetronome := time.NewTicker(durationAttemptChronize)
	defer attemptChronizeMetronome.Stop()

	routerTowardAgreementMetronome := time.NewTicker(r.durationRouterTowardAgreement)
	defer routerTowardAgreementMetronome.Stop()

	var (
		successionUUID                    = r.primaryStatus.SuccessionUUID
		status                      = r.primaryStatus
		primaryEndorseOwnsAdditions = status.FinalLedgerAltitude > 0 &&
			r.depot.FetchLedgerExpandedEndorse(status.FinalLedgerAltitude) != nil

		actedHandleStream = make(chan struct{}, 1)

		//
		ledgersChronized = 0
		finalCentury  = time.Now()
		finalFrequency     = 0.0
	)

FOREACH_CYCLE:
	for {
		select {
		case <-r.Exit():
			break FOREACH_CYCLE
		case <-r.hub.Exit():
			break FOREACH_CYCLE
		case <-routerTowardAgreementMetronome.C:
			altitude, countAwaiting, extentSolicitors := r.hub.ObtainCondition()
			outgoing, incoming, _ := r.Router.CountNodes()

			r.Tracer.Diagnose(
				"REDACTED",
				"REDACTED", countAwaiting,
				"REDACTED", extentSolicitors,
				"REDACTED", outgoing,
				"REDACTED", incoming,
				"REDACTED", status.FinalLedgerAltitude,
			)

			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			//
			absentAddition := true
			if status.FinalLedgerAltitude == 0 ||
				!status.AgreementSettings.Iface.BallotAdditionsActivated(status.FinalLedgerAltitude) ||
				ledgersChronized > 0 ||
				primaryEndorseOwnsAdditions {
				absentAddition = false
			}

			//
			if absentAddition {
				r.Tracer.Details(
					"REDACTED",
					"REDACTED", altitude,
					"REDACTED", status.FinalLedgerAltitude,
					"REDACTED", status.PrimaryAltitude,
					"REDACTED", r.hub.MaximumNodeAltitude(),
				)
				continue FOREACH_CYCLE
			}

			//
			if !r.hub.EqualsSeizedActive() && !r.regionalPeerLedgersThatSuccession(status) {
				continue FOREACH_CYCLE
			}

			r.Tracer.Details("REDACTED", "REDACTED", altitude)
			if err := r.hub.Halt(); err != nil {
				r.Tracer.Failure("REDACTED", "REDACTED", err)
			}

			memoryReader, present := r.Router.Handler("REDACTED")
			if present {
				if memoryReader, ok := memoryReader.(txpoolHandler); ok {
					memoryReader.ActivateInsideOutputTrans()
				}
			}

			connectionReader, present := r.Router.Handler("REDACTED")
			if present {
				if connectionReader, ok := connectionReader.(agreementHandler); ok {
					connectionReader.RouterTowardAgreement(status, ledgersChronized > 0 || statusChronized)
				}
			}

			break FOREACH_CYCLE
		case <-attemptChronizeMetronome.C:
			select {
			case actedHandleStream <- struct{}{}:
			default:
			}
		case <-actedHandleStream:
			//
			//
			//
			//
			//
			//
			//

			//
			initial, ordinal, addnEndorse := r.hub.GlanceCoupleLedgers()
			if initial == nil || ordinal == nil {
				//
				//
				continue FOREACH_CYCLE
			}
			//
			if status.FinalLedgerAltitude > 0 && status.FinalLedgerAltitude+1 != initial.Altitude {
				//
				panic(fmt.Errorf("REDACTED", status.FinalLedgerAltitude+1, initial.Altitude))
			}
			if initial.Altitude+1 != ordinal.Altitude {
				//
				panic(fmt.Errorf("REDACTED", status.FinalLedgerAltitude, initial.Altitude))
			}

			//
			//
			//
			//
			if !r.EqualsActive() || !r.hub.EqualsActive() {
				break FOREACH_CYCLE
			}
			//
			actedHandleStream <- struct{}{}

			initialFragments, err := initial.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
			if err != nil {
				r.Tracer.Failure("REDACTED", "REDACTED", initial.Altitude, "REDACTED", err.Error())
				break FOREACH_CYCLE
			}

			initialFragmentAssignHeading := initialFragments.Heading()
			initialUUID := kinds.LedgerUUID{Digest: initial.Digest(), FragmentAssignHeading: initialFragmentAssignHeading}

			//
			//
			//
			//
			//
			err = status.Assessors.ValidateEndorseAgile(successionUUID, initialUUID, initial.Altitude, ordinal.FinalEndorse)

			if err == nil {
				//
				err = r.ledgerExecute.CertifyLedger(status, initial)
			}

			//
			existingAddnEndorse := addnEndorse != nil
			additionsActivated := status.AgreementSettings.Iface.BallotAdditionsActivated(initial.Altitude)
			if existingAddnEndorse != additionsActivated {
				err = fmt.Errorf("REDACTED"+
					"REDACTED",
					initial.Altitude, existingAddnEndorse, additionsActivated,
				)
			}
			if err == nil && additionsActivated {
				//
				err = addnEndorse.AssureAdditions(true)
			}
			if err == nil && additionsActivated {
				//
				err = status.Assessors.ValidateEndorseAgile(successionUUID, initialUUID, initial.Altitude, addnEndorse.TowardEndorse())
			}

			if err != nil {
				r.processCertificationBreakdown(initial, ordinal, err)
				continue FOREACH_CYCLE
			}

			r.hub.ExtractSolicit()

			//
			if additionsActivated {
				r.depot.PersistLedgerUsingExpandedEndorse(initial, initialFragments, addnEndorse)
			} else {
				//
				//
				//
				//
				r.depot.PersistLedger(initial, initialFragments, ordinal.FinalEndorse)
			}

			//
			//
			status, err = r.ledgerExecute.ExecuteAttestedLedger(status, initialUUID, initial)
			if err != nil {
				//
				panic(fmt.Sprintf("REDACTED", initial.Altitude, initial.Digest(), err))
			}

			r.telemetry.logLedgerTelemetry(initial)
			ledgersChronized++

			if ledgersChronized%100 == 0 {
				finalFrequency = 0.9*finalFrequency + 0.1*(100/time.Since(finalCentury).Seconds())
				finalCentury = time.Now()
				r.Tracer.Details(
					"REDACTED",
					"REDACTED", r.hub.altitude,
					"REDACTED", r.hub.MaximumNodeAltitude(),
					"REDACTED", finalFrequency,
				)
			}

			continue FOREACH_CYCLE
		}
	}
}

func (r *Handler) hubSignalsProcedure(conditionReviseMetronome *time.Ticker) {
	for {
		select {
		case <-r.Exit():
			return
		case <-r.hub.Exit():
			return
		case solicit := <-r.solicitsStream:
			//
			node := r.Router.Nodes().Get(solicit.NodeUUID)
			if node == nil {
				continue
			}

			staged := node.AttemptTransmit(p2p.Wrapper{
				ConduitUUID: ChainchronizeConduit,
				Signal:   &chainchema.LedgerSolicit{Altitude: solicit.Altitude},
			})

			if !staged {
				r.Tracer.Diagnose("REDACTED", "REDACTED", node.ID(), "REDACTED", solicit.Altitude)
			}
		case err := <-r.faultsStream:
			//
			if node := r.Router.Nodes().Get(err.nodeUUID); node != nil {
				r.haltNodeForeachFailure(node, err.err)
			}
		case <-conditionReviseMetronome.C:
			//
			r.Router.MulticastAsyncronous(p2p.Wrapper{
				ConduitUUID: ChainchronizeConduit,
				Signal:   &chainchema.ConditionSolicit{},
			})
		}
	}
}

func (r *Handler) processCertificationBreakdown(ledgerAN, ledgerBYTE *kinds.Ledger, err error) {
	r.Tracer.Failure("REDACTED", "REDACTED", ledgerAN.Altitude, "REDACTED", ledgerAN.Digest(), "REDACTED", err)

	err = FaultHandlerCertification{Err: err}

	idA := r.hub.DiscardNodeAlsoReiterateEveryNodeSolicits(ledgerAN.Altitude)
	if nodeAN := r.Router.Nodes().Get(idA); nodeAN != nil {
		//
		//
		r.haltNodeForeachFailure(nodeAN, err)
	}

	idB := r.hub.DiscardNodeAlsoReiterateEveryNodeSolicits(ledgerBYTE.Altitude)
	if idA == idB {
		return
	}

	if nodeBYTE := r.Router.Nodes().Get(idB); nodeBYTE != nil {
		//
		//
		r.haltNodeForeachFailure(nodeBYTE, err)
	}
}

func (r *Handler) haltNodeForeachFailure(node p2p.Node, err error) {
	if r.aggregateStyleActivated && mustExistReestablished(err) {
		err = &netp2p.FailureFleeting{Err: err}
	}

	r.Router.HaltNodeForeachFailure(node, err)
}

//
//
//
//
func mustExistReestablished(err error) bool {
	//
	return errors.Is(err, FaultNodeDeadline)
}
