package agreement

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"

	controlkinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/digits"
	strongmindincidents "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/incidents"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	strongmindcons "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/agreement"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

const (
	StatusConduit       = byte(0x20)
	DataConduit        = byte(0x21)
	BallotConduit        = byte(0x22)
	BallotAssignDigitsConduit = byte(0x23)

	maximumSignalExtent = 1048576 //

	ledgersTowardInputTowardTransformValidNode = 10000
	ballotsTowardInputTowardTransformValidNode  = 10000
)

//

//
type Handler struct {
	p2p.FoundationHandler //

	connectionSTR *Status

	awaitChronize atomic.Bool
	incidentChannel *kinds.IncidentChannel

	resultsMutex         commitchronize.ReadwriteExclusion
	rs            controlkinds.IterationStatus //
	primaryAltitude atomic.Int64

	agreementParameters atomic.Pointer[kinds.AgreementSettings] //

	Telemetry *Telemetry
}

type HandlerSelection func(*Handler)

//
func FreshHandler(agreementStatus *Status, awaitChronize bool, choices ...HandlerSelection) *Handler {
	connectionReader := &Handler{
		connectionSTR:          agreementStatus,
		awaitChronize:      atomic.Bool{},
		rs:            agreementStatus.obtainIterationStatus(),
		primaryAltitude: atomic.Int64{},
		Telemetry:       NooperationTelemetry(),
	}
	//
	parameters := agreementStatus.status.AgreementSettings
	connectionReader.agreementParameters.Store(&parameters)
	connectionReader.primaryAltitude.Store(agreementStatus.status.PrimaryAltitude)
	connectionReader.FoundationHandler = *p2p.FreshFoundationHandler("REDACTED", connectionReader)
	if awaitChronize {
		connectionReader.awaitChronize.Store(true)
	}

	for _, selection := range choices {
		selection(connectionReader)
	}

	return connectionReader
}

//
//
func (connectionReader *Handler) UponInitiate() error {
	if connectionReader.AwaitChronize() {
		connectionReader.Tracer.Details("REDACTED")
	}

	//
	go connectionReader.nodeMetricsProcedure()

	connectionReader.listenTowardMulticastIncidents()

	if !connectionReader.AwaitChronize() {
		err := connectionReader.connectionSTR.Initiate()
		if err != nil {
			return err
		}
	}

	return nil
}

//
//
func (connectionReader *Handler) UponHalt() {
	connectionReader.cancelOriginatingMulticastIncidents()
	if err := connectionReader.connectionSTR.Halt(); err != nil {
		connectionReader.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	if !connectionReader.AwaitChronize() {
		connectionReader.connectionSTR.Await()
	}
}

//
//
func (connectionReader *Handler) RouterTowardAgreement(status sm.Status, omitJournal bool) {
	connectionReader.Tracer.Details("REDACTED")

	//
	func() {
		//
		connectionReader.connectionSTR.mtx.Lock()
		defer connectionReader.connectionSTR.mtx.Unlock()
		//
		if status.FinalLedgerAltitude > 0 {
			connectionReader.connectionSTR.rebuildFinalEndorse(status)
		}

		//
		//
		connectionReader.connectionSTR.reviseTowardStatus(status)
	}()

	//
	connectionReader.awaitChronize.Store(false)

	if omitJournal {
		connectionReader.connectionSTR.performJournalOvertake = false
	}

	//
	err := connectionReader.connectionSTR.Initiate()
	if err != nil {
		panic(fmt.Sprintf(`REDACTEDv

REDACTED:
REDACTEDv

REDACTED:
REDACTED`, err, connectionReader.connectionSTR, connectionReader))
	}
}

//
func (connectionReader *Handler) ObtainConduits() []*p2p.ConduitDefinition {
	//
	return []*p2p.ConduitDefinition{
		{
			ID:                  StatusConduit,
			Urgency:            6,
			TransmitStagingVolume:   100,
			ObtainSignalVolume: maximumSignalExtent,
			SignalKind:         &strongmindcons.Signal{},
		},
		{
			ID: DataConduit, //
			//
			Urgency:            10,
			TransmitStagingVolume:   100,
			ObtainReserveVolume:  50 * 4096,
			ObtainSignalVolume: maximumSignalExtent,
			SignalKind:         &strongmindcons.Signal{},
		},
		{
			ID:                  BallotConduit,
			Urgency:            7,
			TransmitStagingVolume:   100,
			ObtainReserveVolume:  100 * 100,
			ObtainSignalVolume: maximumSignalExtent,
			SignalKind:         &strongmindcons.Signal{},
		},
		{
			ID:                  BallotAssignDigitsConduit,
			Urgency:            1,
			TransmitStagingVolume:   2,
			ObtainReserveVolume:  1024,
			ObtainSignalVolume: maximumSignalExtent,
			SignalKind:         &strongmindcons.Signal{},
		},
	}
}

//
func (connectionReader *Handler) InitializeNode(node p2p.Node) p2p.Node {
	nodeStatus := FreshNodeStatus(node).AssignTracer(connectionReader.Tracer)
	node.Set(kinds.NodeStatusToken, nodeStatus)
	return node
}

//
//
func (connectionReader *Handler) AppendNode(node p2p.Node) {
	if !connectionReader.EqualsActive() {
		return
	}

	nodeStatus, ok := node.Get(kinds.NodeStatusToken).(*NodeStatus)
	if !ok {
		panic(fmt.Sprintf("REDACTED", node))
	}
	//
	go connectionReader.multicastDataProcedure(node, nodeStatus)
	go connectionReader.multicastBallotsProcedure(node, nodeStatus)
	go connectionReader.inquireMajor23task(node, nodeStatus)

	//
	//
	if !connectionReader.AwaitChronize() {
		connectionReader.transmitFreshIterationPhaseSignal(node)
	}
}

//
func (connectionReader *Handler) DiscardNode(p2p.Node, any) {
	if !connectionReader.EqualsActive() {
		return
	}
	//
	//
	//
	//
	//
	//
}

//
//
//
//
//
//
func (connectionReader *Handler) Accept(e p2p.Wrapper) {
	if !connectionReader.EqualsActive() {
		connectionReader.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.ConduitUUID)
		return
	}
	msg, err := SignalOriginatingSchema(e.Signal)
	if err != nil {
		connectionReader.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", e.ConduitUUID, "REDACTED", err)
		connectionReader.Router.HaltNodeForeachFailure(e.Src, err)
		return
	}

	if err = msg.CertifyFundamental(); err != nil {
		connectionReader.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", e.Signal, "REDACTED", err)
		connectionReader.Router.HaltNodeForeachFailure(e.Src, err)
		return
	}

	connectionReader.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.ConduitUUID, "REDACTED", msg)

	//
	ps, ok := e.Src.Get(kinds.NodeStatusToken).(*NodeStatus)
	if !ok {
		panic(fmt.Sprintf("REDACTED", e.Src))
	}

	switch e.ConduitUUID {
	case StatusConduit:
		switch msg := msg.(type) {
		case *FreshIterationPhaseSignal:
			primaryAltitude := connectionReader.primaryAltitude.Load()
			if err = msg.CertifyAltitude(primaryAltitude); err != nil {
				connectionReader.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", msg, "REDACTED", err)
				connectionReader.Router.HaltNodeForeachFailure(e.Src, err)
				return
			}
			ps.ExecuteFreshIterationPhaseSignal(msg)
			connectionReader.connectionSTR.metricsSignalStaging <- signalDetails{msg, e.Src.ID()}
		case *FreshSoundLedgerSignal:
			ps.ExecuteFreshSoundLedgerSignal(msg)
		case *OwnsBallotSignal:
			ps.ExecuteOwnsBallotSignal(msg)
		case *BallotAssignMajor23signal:
			rs := connectionReader.obtainIterationStatus()
			altitude, ballots := rs.Altitude, rs.Ballots
			if altitude != msg.Altitude {
				return
			}
			//
			err := ballots.AssignNodeMajor23(msg.Iteration, msg.Kind, ps.node.ID(), msg.LedgerUUID)
			if err != nil {
				connectionReader.Router.HaltNodeForeachFailure(e.Src, err)
				return
			}
			//
			//
			var mineBallots *digits.DigitCollection
			switch msg.Kind {
			case commitchema.PreballotKind:
				mineBallots = ballots.Preballots(msg.Iteration).DigitCollectionViaLedgerUUID(msg.LedgerUUID)
			case commitchema.PreendorseKind:
				mineBallots = ballots.Preendorsements(msg.Iteration).DigitCollectionViaLedgerUUID(msg.LedgerUUID)
			default:
				panic("REDACTED")
			}
			exSignal := &strongmindcons.BallotAssignDigits{
				Altitude:  msg.Altitude,
				Iteration:   msg.Iteration,
				Kind:    msg.Kind,
				LedgerUUID: msg.LedgerUUID.TowardSchema(),
			}
			if ballots := mineBallots.TowardSchema(); ballots != nil {
				exSignal.Ballots = *ballots
			}
			e.Src.AttemptTransmit(p2p.Wrapper{
				ConduitUUID: BallotAssignDigitsConduit,
				Signal:   exSignal,
			})
		default:
			connectionReader.Tracer.Failure(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
		}

	case DataConduit:
		if connectionReader.AwaitChronize() {
			connectionReader.Tracer.Details("REDACTED", "REDACTED", msg)
			return
		}
		switch msg := msg.(type) {
		case *NominationSignal:
			parameters := connectionReader.agreementParameters.Load()
			maximumOctets := parameters.Ledger.MaximumOctets
			if err := msg.Nomination.CertifyLedgerExtent(maximumOctets); err != nil {
				connectionReader.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", msg.Nomination.Altitude)
				connectionReader.Router.HaltNodeForeachFailure(e.Src, FaultNominationExcessivelyMultipleFragments)
				return
			}

			ps.AssignOwnsNomination(msg.Nomination)
			connectionReader.connectionSTR.nodeSignalStaging <- signalDetails{msg, e.Src.ID()}
		case *NominationPolicySignal:
			ps.ExecuteNominationPolicySignal(msg)
		case *LedgerFragmentSignal:
			ps.AssignOwnsNominationLedgerFragment(msg.Altitude, msg.Iteration, int(msg.Fragment.Ordinal))
			connectionReader.Telemetry.LedgerFragments.With("REDACTED", string(e.Src.ID())).Add(1)
			connectionReader.connectionSTR.nodeSignalStaging <- signalDetails{msg, e.Src.ID()}
		default:
			connectionReader.Tracer.Failure(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
		}

	case BallotConduit:
		if connectionReader.AwaitChronize() {
			connectionReader.Tracer.Details("REDACTED", "REDACTED", msg)
			return
		}
		switch msg := msg.(type) {
		case *BallotSignal:
			rs := connectionReader.obtainIterationStatus()

			altitude, itemExtent, finalEndorseExtent := rs.Altitude, rs.Assessors.Extent(), rs.FinalEndorse.Extent()
			ps.AssignOwnsBallotOriginatingNode(msg.Ballot, altitude, itemExtent, finalEndorseExtent)

			connectionReader.connectionSTR.nodeSignalStaging <- signalDetails{msg, e.Src.ID()}

		default:
			//
			connectionReader.Tracer.Failure(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
		}

	case BallotAssignDigitsConduit:
		if connectionReader.AwaitChronize() {
			connectionReader.Tracer.Details("REDACTED", "REDACTED", msg)
			return
		}
		switch msg := msg.(type) {
		case *BallotAssignDigitsSignal:
			rs := connectionReader.obtainIterationStatus()

			altitude, ballots := rs.Altitude, rs.Ballots

			if altitude == msg.Altitude {
				var mineBallots *digits.DigitCollection
				switch msg.Kind {
				case commitchema.PreballotKind:
					mineBallots = ballots.Preballots(msg.Iteration).DigitCollectionViaLedgerUUID(msg.LedgerUUID)
				case commitchema.PreendorseKind:
					mineBallots = ballots.Preendorsements(msg.Iteration).DigitCollectionViaLedgerUUID(msg.LedgerUUID)
				default:
					panic("REDACTED")
				}
				ps.ExecuteBallotAssignDigitsSignal(msg, mineBallots)
			} else {
				ps.ExecuteBallotAssignDigitsSignal(msg, nil)
			}
		default:
			//
			connectionReader.Tracer.Failure(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
		}

	default:
		connectionReader.Tracer.Failure(fmt.Sprintf("REDACTED", e.ConduitUUID))
	}
}

//
func (connectionReader *Handler) AssignIncidentChannel(b *kinds.IncidentChannel) {
	connectionReader.incidentChannel = b
	connectionReader.connectionSTR.AssignIncidentChannel(b)
}

//
func (connectionReader *Handler) AwaitChronize() bool {
	return connectionReader.awaitChronize.Load()
}

//
func (connectionReader *Handler) AbsorbAttestedLedger(ledger AbsorbNominee) error {
	return connectionReader.connectionSTR.AbsorbAttestedLedger(ledger)
}

//

//
//
//
func (connectionReader *Handler) listenTowardMulticastIncidents() {
	const listener = "REDACTED"
	err := connectionReader.connectionSTR.incidentctl.AppendObserverForeachIncident(
		listener,
		kinds.IncidentFreshIterationPhase,
		func(data strongmindincidents.IncidentData) {
			rs := data.(controlkinds.IterationStatus)

			//
			connectionReader.reviseIterationStatus(&rs)

			connectionReader.multicastFreshIterationPhaseSignal(&rs)
		},
	)
	if err != nil {
		connectionReader.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	err = connectionReader.connectionSTR.incidentctl.AppendObserverForeachIncident(
		listener,
		kinds.IncidentSoundLedger,
		func(data strongmindincidents.IncidentData) {
			rs := data.(controlkinds.IterationStatus)

			//
			connectionReader.reviseIterationStatus(&rs)

			connectionReader.multicastFreshSoundLedgerSignal(&rs)
		},
	)
	if err != nil {
		connectionReader.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	err = connectionReader.connectionSTR.incidentctl.AppendObserverForeachIncident(
		listener,
		kinds.IncidentBallot,
		func(data strongmindincidents.IncidentData) {
			connectionReader.multicastOwnsBallotSignal(data.(*kinds.Ballot))

			//
			//
			//
			//
			rs := connectionReader.connectionSTR.obtainIterationStatus()
			connectionReader.reviseIterationStatus(&rs)
		},
	)
	if err != nil {
		connectionReader.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	err = connectionReader.connectionSTR.incidentctl.AppendObserverForeachIncident(
		listener,
		kinds.IncidentFreshAgreementParameters,
		func(data strongmindincidents.IncidentData) {
			agreementParameters := data.(kinds.AgreementSettings)

			//
			connectionReader.reviseAgreementParameters(agreementParameters)
		},
	)
	if err != nil {
		connectionReader.Tracer.Failure("REDACTED", "REDACTED", err)
	}
}

//
func (connectionReader *Handler) reviseAgreementParameters(agreementParameters kinds.AgreementSettings) {
	parameters := agreementParameters //
	connectionReader.agreementParameters.Store(&parameters)
}

//
func (connectionReader *Handler) reviseIterationStatus(rs *controlkinds.IterationStatus) {
	connectionReader.resultsMutex.Lock()
	connectionReader.rs = *rs //
	connectionReader.resultsMutex.Unlock()
}

func (connectionReader *Handler) cancelOriginatingMulticastIncidents() {
	const listener = "REDACTED"
	connectionReader.connectionSTR.incidentctl.DiscardObserver(listener)
}

func (connectionReader *Handler) multicastFreshIterationPhaseSignal(rs *controlkinds.IterationStatus) {
	numbersSignal := createIterationPhaseSignal(rs)
	go func() {
		connectionReader.Router.MulticastAsyncronous(p2p.Wrapper{
			ConduitUUID: StatusConduit,
			Signal:   numbersSignal,
		})
	}()
}

func (connectionReader *Handler) multicastFreshSoundLedgerSignal(rs *controlkinds.IterationStatus) {
	psh := rs.NominationLedgerFragments.Heading()
	controlSignal := &strongmindcons.FreshSoundLedger{
		Altitude:             rs.Altitude,
		Iteration:              rs.Iteration,
		LedgerFragmentAssignHeading: psh.TowardSchema(),
		LedgerFragments:         rs.NominationLedgerFragments.DigitCollection().TowardSchema(),
		EqualsEndorse:           rs.Phase == controlkinds.IterationPhaseEndorse,
	}
	go func() {
		connectionReader.Router.MulticastAsyncronous(p2p.Wrapper{
			ConduitUUID: StatusConduit,
			Signal:   controlSignal,
		})
	}()
}

//
func (connectionReader *Handler) multicastOwnsBallotSignal(ballot *kinds.Ballot) {
	msg := &strongmindcons.OwnsBallot{
		Altitude: ballot.Altitude,
		Iteration:  ballot.Iteration,
		Kind:   ballot.Kind,
		Ordinal:  ballot.AssessorOrdinal,
	}

	go func() {
		connectionReader.Router.AttemptMulticast(p2p.Wrapper{
			ConduitUUID: StatusConduit,
			Signal:   msg,
		})
	}()
	/**
.
{
)
{
)
}
)
{
?
{
,
,
}
)
{
h
?
.
}
}
*/
}

func createIterationPhaseSignal(rs *controlkinds.IterationStatus) (numbersSignal *strongmindcons.FreshIterationPhase) {
	numbersSignal = &strongmindcons.FreshIterationPhase{
		Altitude:                rs.Altitude,
		Iteration:                 rs.Iteration,
		Phase:                  uint32(rs.Phase),
		MomentsBecauseInitiateMoment: int64(time.Since(rs.InitiateMoment).Seconds()),
		FinalEndorseIteration:       rs.FinalEndorse.ObtainIteration(),
	}
	return
}

func (connectionReader *Handler) transmitFreshIterationPhaseSignal(node p2p.Node) {
	rs := connectionReader.obtainIterationStatus()
	numbersSignal := createIterationPhaseSignal(&rs)
	node.Transmit(p2p.Wrapper{
		ConduitUUID: StatusConduit,
		Signal:   numbersSignal,
	})
}

func (connectionReader *Handler) obtainIterationStatus() controlkinds.IterationStatus {
	connectionReader.resultsMutex.RLock()
	defer connectionReader.resultsMutex.RUnlock()
	return connectionReader.rs
}

//
//

func (connectionReader *Handler) multicastDataProcedure(node p2p.Node, ps *NodeStatus) {
	tracer := connectionReader.Tracer.Using("REDACTED", node)

EXTERNAL_CYCLE:
	for {
		//
		if !node.EqualsActive() || !connectionReader.EqualsActive() {
			return
		}
		rs := connectionReader.obtainIterationStatus()
		prs := ps.ObtainIterationStatus()

		//
		//
		//
		//

		if fragment, proceedCycle := selectFragmentTowardTransmit(tracer, connectionReader.connectionSTR.ledgerDepot, &rs, ps, prs); fragment != nil {
			//
			//
			if ps.TransmitFragmentAssignOwnsFragment(fragment, prs) || proceedCycle {
				continue EXTERNAL_CYCLE
			}
		} else if proceedCycle {
			//
			continue EXTERNAL_CYCLE
		}

		//
		//
		//
		//

		altitudeIterationAlign := (rs.Altitude == prs.Altitude) && (rs.Iteration == prs.Iteration)
		nominationTowardTransmit := rs.Nomination != nil && !prs.Nomination

		if altitudeIterationAlign && nominationTowardTransmit {
			ps.TransmitNominationAssignOwnsNomination(tracer, &rs, prs)
			continue EXTERNAL_CYCLE
		}

		//
		time.Sleep(connectionReader.connectionSTR.settings.NodeMulticastSnoozeInterval)
	}
}

func (connectionReader *Handler) multicastBallotsProcedure(node p2p.Node, ps *NodeStatus) {
	tracer := connectionReader.Tracer.Using("REDACTED", node)

	//
	dormant := 0

EXTERNAL_CYCLE:
	for {
		//
		if !node.EqualsActive() || !connectionReader.EqualsActive() {
			return
		}
		rs := connectionReader.obtainIterationStatus()
		prs := ps.ObtainIterationStatus()

		switch dormant {
		case 1: //
			dormant = 2
		case 2: //
			dormant = 0
		}

		if ballot := selectBallotTowardTransmit(tracer, connectionReader.connectionSTR, &rs, ps, prs); ballot != nil {
			if ps.transmitBallotAssignOwnsBallot(ballot) {
				continue EXTERNAL_CYCLE
			}
			tracer.Diagnose("REDACTED",
				"REDACTED", prs.Altitude,
				"REDACTED", ballot,
			)
		}

		switch dormant {
		case 0:
			//
			dormant = 1
			tracer.Diagnose("REDACTED", "REDACTED", rs.Altitude, "REDACTED", prs.Altitude,
				"REDACTED", rs.Ballots.Preballots(rs.Iteration).DigitCollection(), "REDACTED", prs.Preballots,
				"REDACTED", rs.Ballots.Preendorsements(rs.Iteration).DigitCollection(), "REDACTED", prs.Preendorsements)
		case 2:
			//
			dormant = 1
		}

		time.Sleep(connectionReader.connectionSTR.settings.NodeMulticastSnoozeInterval)
	}
}

//
//
func (connectionReader *Handler) inquireMajor23task(node p2p.Node, ps *NodeStatus) {
EXTERNAL_CYCLE:
	for {
		//
		if !node.EqualsActive() || !connectionReader.EqualsActive() {
			return
		}

		//
		{
			rs := connectionReader.obtainIterationStatus()
			prs := ps.ObtainIterationStatus()
			if rs.Altitude == prs.Altitude {
				if major23, ok := rs.Ballots.Preballots(prs.Iteration).CoupleTrinityPreponderance(); ok {

					node.AttemptTransmit(p2p.Wrapper{
						ConduitUUID: StatusConduit,
						Signal: &strongmindcons.BallotAssignMajor23{
							Altitude:  prs.Altitude,
							Iteration:   prs.Iteration,
							Kind:    commitchema.PreballotKind,
							LedgerUUID: major23.TowardSchema(),
						},
					})
					time.Sleep(connectionReader.connectionSTR.settings.NodeInquireMajor23dormantInterval)
				}
			}
		}

		//
		{
			rs := connectionReader.obtainIterationStatus()
			prs := ps.ObtainIterationStatus()
			if rs.Altitude == prs.Altitude {
				if major23, ok := rs.Ballots.Preendorsements(prs.Iteration).CoupleTrinityPreponderance(); ok {
					node.AttemptTransmit(p2p.Wrapper{
						ConduitUUID: StatusConduit,
						Signal: &strongmindcons.BallotAssignMajor23{
							Altitude:  prs.Altitude,
							Iteration:   prs.Iteration,
							Kind:    commitchema.PreendorseKind,
							LedgerUUID: major23.TowardSchema(),
						},
					})
					time.Sleep(connectionReader.connectionSTR.settings.NodeInquireMajor23dormantInterval)
				}
			}
		}

		//
		{
			rs := connectionReader.obtainIterationStatus()
			prs := ps.ObtainIterationStatus()
			if rs.Altitude == prs.Altitude && prs.NominationPolicyIteration >= 0 {
				if major23, ok := rs.Ballots.Preballots(prs.NominationPolicyIteration).CoupleTrinityPreponderance(); ok {

					node.AttemptTransmit(p2p.Wrapper{
						ConduitUUID: StatusConduit,
						Signal: &strongmindcons.BallotAssignMajor23{
							Altitude:  prs.Altitude,
							Iteration:   prs.NominationPolicyIteration,
							Kind:    commitchema.PreballotKind,
							LedgerUUID: major23.TowardSchema(),
						},
					})
					time.Sleep(connectionReader.connectionSTR.settings.NodeInquireMajor23dormantInterval)
				}
			}
		}

		//
		//

		//
		{
			prs := ps.ObtainIterationStatus()
			if prs.OvertakeEndorseIteration != -1 && prs.Altitude > 0 && prs.Altitude <= connectionReader.connectionSTR.ledgerDepot.Altitude() &&
				prs.Altitude >= connectionReader.connectionSTR.ledgerDepot.Foundation() {
				if endorse := connectionReader.connectionSTR.FetchEndorse(prs.Altitude); endorse != nil {
					node.AttemptTransmit(p2p.Wrapper{
						ConduitUUID: StatusConduit,
						Signal: &strongmindcons.BallotAssignMajor23{
							Altitude:  prs.Altitude,
							Iteration:   endorse.Iteration,
							Kind:    commitchema.PreendorseKind,
							LedgerUUID: endorse.LedgerUUID.TowardSchema(),
						},
					})
					time.Sleep(connectionReader.connectionSTR.settings.NodeInquireMajor23dormantInterval)
				}
			}
		}

		time.Sleep(connectionReader.connectionSTR.settings.NodeInquireMajor23dormantInterval)

		continue EXTERNAL_CYCLE
	}
}

//
//
//
func selectFragmentTowardTransmit(
	tracer log.Tracer,
	ledgerDepot sm.LedgerDepot,
	rs *controlkinds.IterationStatus,
	ps *NodeStatus,
	prs *controlkinds.NodeIterationStatus,
) (*kinds.Fragment, bool) {
	//
	if rs.NominationLedgerFragments.OwnsHeading(prs.NominationLedgerFragmentAssignHeading) {
		if ordinal, ok := rs.NominationLedgerFragments.DigitCollection().Sub(prs.NominationLedgerFragments.Duplicate()).SelectArbitrary(); ok {
			fragment := rs.NominationLedgerFragments.ObtainFragment(ordinal)
			//
			return fragment, true
		}
	}

	//
	ledgerDepotFoundation := ledgerDepot.Foundation()
	if ledgerDepotFoundation > 0 &&
		0 < prs.Altitude && prs.Altitude < rs.Altitude &&
		prs.Altitude >= ledgerDepotFoundation {
		altitudeTracer := tracer.Using("REDACTED", prs.Altitude)

		//
		if prs.NominationLedgerFragments == nil {
			ledgerSummary := ledgerDepot.FetchLedgerSummary(prs.Altitude)
			if ledgerSummary == nil {
				altitudeTracer.Failure("REDACTED",
					"REDACTED", ledgerDepotFoundation, "REDACTED", ledgerDepot.Altitude())
				return nil, false
			}
			ps.InitializeNominationLedgerFragments(ledgerSummary.LedgerUUID.FragmentAssignHeading)
			//
			return nil, true //
		}
		fragment := selectFragmentForeachOvertake(altitudeTracer, rs, prs, ledgerDepot)
		if fragment != nil {
			//
			return fragment, false
		}
	}

	return nil, false
}

func selectFragmentForeachOvertake(
	tracer log.Tracer,
	rs *controlkinds.IterationStatus,
	prs *controlkinds.NodeIterationStatus,
	ledgerDepot sm.LedgerDepot,
) *kinds.Fragment {
	ordinal, ok := prs.NominationLedgerFragments.Not().SelectArbitrary()
	if !ok {
		return nil
	}
	//
	ledgerSummary := ledgerDepot.FetchLedgerSummary(prs.Altitude)
	if ledgerSummary == nil {
		tracer.Failure("REDACTED", "REDACTED", rs.Altitude,
			"REDACTED", ledgerDepot.Foundation(), "REDACTED", ledgerDepot.Altitude())
		return nil
	} else if !ledgerSummary.LedgerUUID.FragmentAssignHeading.Matches(prs.NominationLedgerFragmentAssignHeading) {
		tracer.Details("REDACTED",
			"REDACTED", ledgerSummary.LedgerUUID.FragmentAssignHeading, "REDACTED", prs.NominationLedgerFragmentAssignHeading)
		return nil
	}
	//
	fragment := ledgerDepot.FetchLedgerFragment(prs.Altitude, ordinal)
	if fragment == nil {
		tracer.Failure("REDACTED", "REDACTED", ordinal,
			"REDACTED", ledgerSummary.LedgerUUID.FragmentAssignHeading, "REDACTED", prs.NominationLedgerFragmentAssignHeading)
		return nil
	}
	return fragment
}

func selectBallotTowardTransmit(
	tracer log.Tracer,
	connectionSTR *Status,
	rs *controlkinds.IterationStatus,
	ps *NodeStatus,
	prs *controlkinds.NodeIterationStatus,
) *kinds.Ballot {
	//
	if rs.Altitude == prs.Altitude {
		altitudeTracer := tracer.Using("REDACTED", prs.Altitude)
		return selectBallotPrevailingAltitude(altitudeTracer, rs, prs, ps)
	}

	//
	//
	if prs.Altitude != 0 && rs.Altitude == prs.Altitude+1 {
		if ballot := ps.SelectBallotTowardTransmit(rs.FinalEndorse); ballot != nil {
			tracer.Diagnose("REDACTED", "REDACTED", prs.Altitude)
			return ballot
		}
	}

	//
	//
	ledgerDepotFoundation := connectionSTR.ledgerDepot.Foundation()
	if ledgerDepotFoundation > 0 && prs.Altitude != 0 && rs.Altitude >= prs.Altitude+2 && prs.Altitude >= ledgerDepotFoundation {
		//
		//
		var ec *kinds.ExpandedEndorse
		var verActivated bool
		func() {
			connectionSTR.mtx.RLock()
			defer connectionSTR.mtx.RUnlock()
			verActivated = connectionSTR.status.AgreementSettings.Iface.BallotAdditionsActivated(prs.Altitude)
		}()
		if verActivated {
			ec = connectionSTR.ledgerDepot.FetchLedgerExpandedEndorse(prs.Altitude)
		} else {
			c := connectionSTR.ledgerDepot.FetchLedgerEndorse(prs.Altitude)
			if c == nil {
				return nil
			}
			ec = c.EncapsulatedExpandedEndorse()
		}
		if ec == nil {
			return nil
		}
		if ballot := ps.SelectBallotTowardTransmit(ec); ballot != nil {
			tracer.Diagnose("REDACTED", "REDACTED", prs.Altitude)
			return ballot
		}
	}
	return nil
}

func selectBallotPrevailingAltitude(
	tracer log.Tracer,
	rs *controlkinds.IterationStatus,
	prs *controlkinds.NodeIterationStatus,
	ps *NodeStatus,
) *kinds.Ballot {
	//
	if prs.Phase == controlkinds.IterationPhaseFreshAltitude {
		if ballot := ps.SelectBallotTowardTransmit(rs.FinalEndorse); ballot != nil {
			tracer.Diagnose("REDACTED")
			return ballot
		}
	}
	//
	if prs.Phase <= controlkinds.IterationPhaseNominate && prs.Iteration != -1 && prs.Iteration <= rs.Iteration && prs.NominationPolicyIteration != -1 {
		if policyPreballots := rs.Ballots.Preballots(prs.NominationPolicyIteration); policyPreballots != nil {
			if ballot := ps.SelectBallotTowardTransmit(policyPreballots); ballot != nil {
				tracer.Diagnose("REDACTED",
					"REDACTED", prs.NominationPolicyIteration)
				return ballot
			}
		}
	}
	//
	if prs.Phase <= controlkinds.IterationPhasePreballotAwait && prs.Iteration != -1 && prs.Iteration <= rs.Iteration {
		if ballot := ps.SelectBallotTowardTransmit(rs.Ballots.Preballots(prs.Iteration)); ballot != nil {
			tracer.Diagnose("REDACTED", "REDACTED", prs.Iteration)
			return ballot
		}
	}
	//
	if prs.Phase <= controlkinds.IterationPhasePreendorseAwait && prs.Iteration != -1 && prs.Iteration <= rs.Iteration {
		if ballot := ps.SelectBallotTowardTransmit(rs.Ballots.Preendorsements(prs.Iteration)); ballot != nil {
			tracer.Diagnose("REDACTED", "REDACTED", prs.Iteration)
			return ballot
		}
	}
	//
	if prs.Iteration != -1 && prs.Iteration <= rs.Iteration {
		if ballot := ps.SelectBallotTowardTransmit(rs.Ballots.Preballots(prs.Iteration)); ballot != nil {
			tracer.Diagnose("REDACTED", "REDACTED", prs.Iteration)
			return ballot
		}
	}
	//
	if prs.NominationPolicyIteration != -1 {
		if policyPreballots := rs.Ballots.Preballots(prs.NominationPolicyIteration); policyPreballots != nil {
			if ballot := ps.SelectBallotTowardTransmit(policyPreballots); ballot != nil {
				tracer.Diagnose("REDACTED",
					"REDACTED", prs.NominationPolicyIteration)
				return ballot
			}
		}
	}

	return nil
}

//

func (connectionReader *Handler) nodeMetricsProcedure() {
	for {
		if !connectionReader.EqualsActive() {
			connectionReader.Tracer.Details("REDACTED")
			return
		}

		select {
		case msg := <-connectionReader.connectionSTR.metricsSignalStaging:
			connectionReader.Tracer.Diagnose("REDACTED", "REDACTED", msg.NodeUUID)

			//
			if msg.NodeUUID == "REDACTED" {
				continue
			}

			//
			node := connectionReader.Router.Nodes().Get(msg.NodeUUID)
			if node == nil {
				connectionReader.Tracer.Diagnose("REDACTED", "REDACTED", msg.NodeUUID)
				continue
			}
			//
			ps, ok := node.Get(kinds.NodeStatusToken).(*NodeStatus)
			if !ok {
				panic(fmt.Sprintf("REDACTED", node))
			}
			switch tangibleSignal := msg.Msg.(type) {
			case *BallotSignal:
				if countBallots := ps.LogBallot(); countBallots%ballotsTowardInputTowardTransformValidNode == 0 {
					connectionReader.Router.LabelNodeLikeValid(node)
				}
			case *LedgerFragmentSignal:
				if countFragments := ps.LogLedgerFragment(); countFragments%ledgersTowardInputTowardTransformValidNode == 0 {
					connectionReader.Router.LabelNodeLikeValid(node)
				}
			case *FreshIterationPhaseSignal:
				connectionReader.Telemetry.NodeAltitude.With("REDACTED", string(msg.NodeUUID)).Set(float64(tangibleSignal.Altitude))
			}
		case <-connectionReader.connectionSTR.Exit():
			return

		case <-connectionReader.Exit():
			return
		}
	}
}

//
//
//
func (connectionReader *Handler) Text() string {
	//
	return "REDACTED" //
}

//
func (connectionReader *Handler) TextFormatted(format string) string {
	s := "REDACTED"
	s += format + "REDACTED" + connectionReader.connectionSTR.TextFormatted(format+"REDACTED") + "REDACTED"
	connectionReader.Router.Nodes().ForeachEvery(func(node p2p.Node) {
		ps, ok := node.Get(kinds.NodeStatusToken).(*NodeStatus)
		if !ok {
			panic(fmt.Sprintf("REDACTED", node))
		}
		s += format + "REDACTED" + ps.TextFormatted(format+"REDACTED") + "REDACTED"
	})
	s += format + "REDACTED"
	return s
}

//
func HandlerTelemetry(telemetry *Telemetry) HandlerSelection {
	return func(connectionReader *Handler) { connectionReader.Telemetry = telemetry }
}

//

//
//
//
//
type NodeStatus struct {
	node   p2p.Node
	tracer log.Tracer

	mtx   sync.Mutex             //
	PRS   controlkinds.NodeIterationStatus `json:"iteration_status"` //
	Metrics *nodeStatusMetrics        `json:"metrics"`       //
}

//
type nodeStatusMetrics struct {
	Ballots      int `json:"ballots"`
	LedgerFragments int `json:"ledger_fragments"`
}

func (pss nodeStatusMetrics) Text() string {
	return fmt.Sprintf("REDACTED",
		pss.Ballots, pss.LedgerFragments)
}

//
func FreshNodeStatus(node p2p.Node) *NodeStatus {
	return &NodeStatus{
		node:   node,
		tracer: log.FreshNooperationTracer(),
		PRS: controlkinds.NodeIterationStatus{
			Iteration:              -1,
			NominationPolicyIteration:   -1,
			FinalEndorseIteration:    -1,
			OvertakeEndorseIteration: -1,
		},
		Metrics: &nodeStatusMetrics{},
	}
}

//
//
func (ps *NodeStatus) AssignTracer(tracer log.Tracer) *NodeStatus {
	ps.tracer = tracer
	return ps
}

//
//
func (ps *NodeStatus) ObtainIterationStatus() *controlkinds.NodeIterationStatus {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	prs := ps.PRS //
	return &prs
}

//
func (ps *NodeStatus) SerializeJSN() ([]byte, error) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	type jsnNodeStatus NodeStatus
	return strongmindjson.Serialize((*jsnNodeStatus)(ps))
}

//
//
func (ps *NodeStatus) ObtainAltitude() int64 {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.PRS.Altitude
}

//
func (ps *NodeStatus) AssignOwnsNomination(nomination *kinds.Nomination) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.Altitude != nomination.Altitude || ps.PRS.Iteration != nomination.Iteration {
		return
	}

	if ps.PRS.Nomination {
		return
	}

	ps.PRS.Nomination = true

	//
	if ps.PRS.NominationLedgerFragments != nil {
		return
	}

	ps.PRS.NominationLedgerFragmentAssignHeading = nomination.LedgerUUID.FragmentAssignHeading
	ps.PRS.NominationLedgerFragments = digits.FreshDigitCollection(int(nomination.LedgerUUID.FragmentAssignHeading.Sum))
	ps.PRS.NominationPolicyIteration = nomination.PolicyIteration
	ps.PRS.NominationPolicy = nil //
}

//
func (ps *NodeStatus) InitializeNominationLedgerFragments(fragmentAssignHeading kinds.FragmentAssignHeading) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.NominationLedgerFragments != nil {
		return
	}

	ps.PRS.NominationLedgerFragmentAssignHeading = fragmentAssignHeading
	ps.PRS.NominationLedgerFragments = digits.FreshDigitCollection(int(fragmentAssignHeading.Sum))
}

//
func (ps *NodeStatus) AssignOwnsNominationLedgerFragment(altitude int64, iteration int32, ordinal int) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.Altitude != altitude || ps.PRS.Iteration != iteration {
		return
	}

	ps.PRS.NominationLedgerFragments.AssignOrdinal(ordinal, true)
}

//
//
//
//
//
func (ps *NodeStatus) SelectTransmitBallot(ballots kinds.BallotAssignFetcher) bool {
	if ballot := ps.SelectBallotTowardTransmit(ballots); ballot != nil {
		ps.tracer.Diagnose("REDACTED", "REDACTED", ps, "REDACTED", ballot)
		if ps.node.Transmit(p2p.Wrapper{
			ConduitUUID: BallotConduit,
			Signal: &strongmindcons.Ballot{
				Ballot: ballot.TowardSchema(),
			},
		}) {
			ps.AssignOwnsBallot(ballot)
			return true
		}
		return false
	}
	return false
}

//
//
func (ps *NodeStatus) TransmitFragmentAssignOwnsFragment(fragment *kinds.Fragment, prs *controlkinds.NodeIterationStatus) bool {
	//
	ps.tracer.Diagnose("REDACTED", "REDACTED", prs.Altitude, "REDACTED", prs.Iteration, "REDACTED", fragment.Ordinal)
	pp, err := fragment.TowardSchema()
	if err != nil {
		//
		ps.tracer.Failure("REDACTED", "REDACTED", fragment.Ordinal, "REDACTED", err)
		return false
	}
	if ps.node.Transmit(p2p.Wrapper{
		ConduitUUID: DataConduit,
		Signal: &strongmindcons.LedgerFragment{
			Altitude: prs.Altitude, //
			Iteration:  prs.Iteration,  //
			Fragment:   *pp,
		},
	}) {
		ps.AssignOwnsNominationLedgerFragment(prs.Altitude, prs.Iteration, int(fragment.Ordinal))
		return true
	}
	ps.tracer.Diagnose("REDACTED")
	return false
}

//
//
func (ps *NodeStatus) TransmitNominationAssignOwnsNomination(
	tracer log.Tracer,
	rs *controlkinds.IterationStatus,
	prs *controlkinds.NodeIterationStatus,
) {
	//
	tracer.Diagnose("REDACTED", "REDACTED", prs.Altitude, "REDACTED", prs.Iteration)
	if ps.node.Transmit(p2p.Wrapper{
		ConduitUUID: DataConduit,
		Signal:   &strongmindcons.Nomination{Nomination: *rs.Nomination.TowardSchema()},
	}) {
		//
		ps.AssignOwnsNomination(rs.Nomination)
	}

	//
	//
	//
	//
	if 0 <= rs.Nomination.PolicyIteration {
		tracer.Diagnose("REDACTED", "REDACTED", prs.Altitude, "REDACTED", prs.Iteration)
		ps.node.Transmit(p2p.Wrapper{
			ConduitUUID: DataConduit,
			Signal: &strongmindcons.NominationPolicy{
				Altitude:           rs.Altitude,
				NominationPolicyIteration: rs.Nomination.PolicyIteration,
				NominationPolicy:      *rs.Ballots.Preballots(rs.Nomination.PolicyIteration).DigitCollection().TowardSchema(),
			},
		})
	}
}

//
//
func (ps *NodeStatus) transmitBallotAssignOwnsBallot(ballot *kinds.Ballot) bool {
	ps.tracer.Diagnose("REDACTED", "REDACTED", ps, "REDACTED", ballot)
	if ps.node.Transmit(p2p.Wrapper{
		ConduitUUID: BallotConduit,
		Signal: &strongmindcons.Ballot{
			Ballot: ballot.TowardSchema(),
		},
	}) {
		ps.AssignOwnsBallot(ballot)
		return true
	}
	return false
}

//
//
//
func (ps *NodeStatus) SelectBallotTowardTransmit(ballots kinds.BallotAssignFetcher) *kinds.Ballot {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ballots.Extent() == 0 {
		return nil
	}

	altitude, iteration, ballotsKind, extent := ballots.ObtainAltitude(), ballots.ObtainIteration(), commitchema.AttestedSignalKind(ballots.Kind()), ballots.Extent()

	//
	if ballots.EqualsEndorse() {
		ps.assureOvertakeEndorseIteration(altitude, iteration, extent)
	}
	ps.assureBallotDigitCollections(altitude, extent)

	processesBallots := ps.obtainBallotDigitCollection(altitude, iteration, ballotsKind)
	if processesBallots == nil {
		return nil //
	}
	if ordinal, ok := ballots.DigitCollection().Sub(processesBallots).SelectArbitrary(); ok {
		ballot := ballots.ObtainViaOrdinal(int32(ordinal))
		if ballot == nil {
			ps.tracer.Failure("REDACTED", "REDACTED", ballots, "REDACTED", ordinal)
		}
		return ballot
	}
	return nil
}

func (ps *NodeStatus) obtainBallotDigitCollection(altitude int64, iteration int32, ballotsKind commitchema.AttestedSignalKind) *digits.DigitCollection {
	if !kinds.EqualsBallotKindSound(ballotsKind) {
		return nil
	}

	if ps.PRS.Altitude == altitude {
		if ps.PRS.Iteration == iteration {
			switch ballotsKind {
			case commitchema.PreballotKind:
				return ps.PRS.Preballots
			case commitchema.PreendorseKind:
				return ps.PRS.Preendorsements
			}
		}
		if ps.PRS.OvertakeEndorseIteration == iteration {
			switch ballotsKind {
			case commitchema.PreballotKind:
				return nil
			case commitchema.PreendorseKind:
				return ps.PRS.OvertakeEndorse
			}
		}
		if ps.PRS.NominationPolicyIteration == iteration {
			switch ballotsKind {
			case commitchema.PreballotKind:
				return ps.PRS.NominationPolicy
			case commitchema.PreendorseKind:
				return nil
			}
		}
		return nil
	}
	if ps.PRS.Altitude == altitude+1 {
		if ps.PRS.FinalEndorseIteration == iteration {
			switch ballotsKind {
			case commitchema.PreballotKind:
				return nil
			case commitchema.PreendorseKind:
				return ps.PRS.FinalEndorse
			}
		}
		return nil
	}
	return nil
}

//
func (ps *NodeStatus) assureOvertakeEndorseIteration(altitude int64, iteration int32, countAssessors int) {
	if ps.PRS.Altitude != altitude {
		return
	}
	/**
.
.
{
(
,
,
,
,
,
)
}
*/
	if ps.PRS.OvertakeEndorseIteration == iteration {
		return //
	}
	ps.PRS.OvertakeEndorseIteration = iteration
	if iteration == ps.PRS.Iteration {
		ps.PRS.OvertakeEndorse = ps.PRS.Preendorsements
	} else {
		ps.PRS.OvertakeEndorse = digits.FreshDigitCollection(countAssessors)
	}
}

//
//
//
//
func (ps *NodeStatus) AssureBallotDigitCollections(altitude int64, countAssessors int) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	ps.assureBallotDigitCollections(altitude, countAssessors)
}

func (ps *NodeStatus) assureBallotDigitCollections(altitude int64, countAssessors int) {
	switch ps.PRS.Altitude {
	case altitude:
		if ps.PRS.Preballots == nil {
			ps.PRS.Preballots = digits.FreshDigitCollection(countAssessors)
		}
		if ps.PRS.Preendorsements == nil {
			ps.PRS.Preendorsements = digits.FreshDigitCollection(countAssessors)
		}
		if ps.PRS.OvertakeEndorse == nil {
			ps.PRS.OvertakeEndorse = digits.FreshDigitCollection(countAssessors)
		}
		if ps.PRS.NominationPolicy == nil {
			ps.PRS.NominationPolicy = digits.FreshDigitCollection(countAssessors)
		}
	case altitude + 1:
		if ps.PRS.FinalEndorse == nil {
			ps.PRS.FinalEndorse = digits.FreshDigitCollection(countAssessors)
		}
	}
}

//
//
func (ps *NodeStatus) LogBallot() int {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	ps.Metrics.Ballots++

	return ps.Metrics.Ballots
}

//
//
func (ps *NodeStatus) BallotsRelayed() int {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	return ps.Metrics.Ballots
}

//
//
func (ps *NodeStatus) LogLedgerFragment() int {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	ps.Metrics.LedgerFragments++
	return ps.Metrics.LedgerFragments
}

//
func (ps *NodeStatus) LedgerFragmentsRelayed() int {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	return ps.Metrics.LedgerFragments
}

//
func (ps *NodeStatus) AssignOwnsBallot(ballot *kinds.Ballot) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	ps.assignOwnsBallot(ballot.Altitude, ballot.Iteration, ballot.Kind, ballot.AssessorOrdinal)
}

func (ps *NodeStatus) assignOwnsBallot(altitude int64, iteration int32, ballotKind commitchema.AttestedSignalKind, ordinal int32) {
	ps.tracer.Diagnose("REDACTED",
		"REDACTED",
		log.FreshIdleFormat("REDACTED", ps.PRS.Altitude, ps.PRS.Iteration),
		"REDACTED",
		log.FreshIdleFormat("REDACTED", altitude, iteration),
		"REDACTED", ballotKind, "REDACTED", ordinal)

	//
	processesBallots := ps.obtainBallotDigitCollection(altitude, iteration, ballotKind)
	if processesBallots != nil {
		processesBallots.AssignOrdinal(int(ordinal), true)
	}
}

//
func (ps *NodeStatus) AssignOwnsBallotOriginatingNode(ballot *kinds.Ballot, controlAltitude int64, itemExtent, finalEndorseExtent int) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	ps.assureBallotDigitCollections(controlAltitude, itemExtent)
	ps.assureBallotDigitCollections(controlAltitude-1, finalEndorseExtent)
	ps.assignOwnsBallot(ballot.Altitude, ballot.Iteration, ballot.Kind, ballot.AssessorOrdinal)
}

//
func (ps *NodeStatus) ExecuteFreshIterationPhaseSignal(msg *FreshIterationPhaseSignal) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	//
	if ContrastHours(msg.Altitude, msg.Iteration, msg.Phase, ps.PRS.Altitude, ps.PRS.Iteration, ps.PRS.Phase) <= 0 {
		return
	}

	//
	processesAltitude := ps.PRS.Altitude
	processesIteration := ps.PRS.Iteration
	processesOvertakeEndorseIteration := ps.PRS.OvertakeEndorseIteration
	processesOvertakeEndorse := ps.PRS.OvertakeEndorse
	finalPreendorsements := ps.PRS.Preendorsements

	initiateMoment := committime.Now().Add(-1 * time.Duration(msg.MomentsBecauseInitiateMoment) * time.Second)
	ps.PRS.Altitude = msg.Altitude
	ps.PRS.Iteration = msg.Iteration
	ps.PRS.Phase = msg.Phase
	ps.PRS.InitiateMoment = initiateMoment
	if processesAltitude != msg.Altitude || processesIteration != msg.Iteration {
		ps.PRS.Nomination = false
		ps.PRS.NominationLedgerFragmentAssignHeading = kinds.FragmentAssignHeading{}
		ps.PRS.NominationLedgerFragments = nil
		ps.PRS.NominationPolicyIteration = -1
		ps.PRS.NominationPolicy = nil
		//
		ps.PRS.Preballots = nil
		ps.PRS.Preendorsements = nil
	}
	if processesAltitude == msg.Altitude && processesIteration != msg.Iteration && msg.Iteration == processesOvertakeEndorseIteration {
		//
		//
		//
		//
		ps.PRS.Preendorsements = processesOvertakeEndorse
	}
	if processesAltitude != msg.Altitude {
		//
		if processesAltitude+1 == msg.Altitude && processesIteration == msg.FinalEndorseIteration {
			ps.PRS.FinalEndorseIteration = msg.FinalEndorseIteration
			ps.PRS.FinalEndorse = finalPreendorsements
		} else {
			ps.PRS.FinalEndorseIteration = msg.FinalEndorseIteration
			ps.PRS.FinalEndorse = nil
		}
		//
		ps.PRS.OvertakeEndorseIteration = -1
		ps.PRS.OvertakeEndorse = nil
	}
}

//
func (ps *NodeStatus) ExecuteFreshSoundLedgerSignal(msg *FreshSoundLedgerSignal) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.Altitude != msg.Altitude {
		return
	}

	if ps.PRS.Iteration != msg.Iteration && !msg.EqualsEndorse {
		return
	}

	ps.PRS.NominationLedgerFragmentAssignHeading = msg.LedgerFragmentAssignHeading
	ps.PRS.NominationLedgerFragments = msg.LedgerFragments
}

//
func (ps *NodeStatus) ExecuteNominationPolicySignal(msg *NominationPolicySignal) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.Altitude != msg.Altitude {
		return
	}
	if ps.PRS.NominationPolicyIteration != msg.NominationPolicyIteration {
		return
	}

	//
	//
	ps.PRS.NominationPolicy = msg.NominationPolicy
}

//
func (ps *NodeStatus) ExecuteOwnsBallotSignal(msg *OwnsBallotSignal) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.Altitude != msg.Altitude {
		return
	}

	ps.assignOwnsBallot(msg.Altitude, msg.Iteration, msg.Kind, msg.Ordinal)
}

//
//
//
//
//
func (ps *NodeStatus) ExecuteBallotAssignDigitsSignal(msg *BallotAssignDigitsSignal, mineBallots *digits.DigitCollection) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	ballots := ps.obtainBallotDigitCollection(msg.Altitude, msg.Iteration, msg.Kind)
	if ballots != nil {
		if mineBallots == nil {
			ballots.Revise(msg.Ballots)
		} else {
			anotherBallots := ballots.Sub(mineBallots)
			ownsBallots := anotherBallots.Or(msg.Ballots)
			ballots.Revise(ownsBallots)
		}
	}
}

//
func (ps *NodeStatus) Text() string {
	return ps.TextFormatted("REDACTED")
}

//
func (ps *NodeStatus) TextFormatted(format string) string {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTEDv
REDACTED`,
		format, ps.node.ID(),
		format, ps.PRS.TextFormatted(format+"REDACTED"),
		format, ps.Metrics,
		format)
}

//
//

//
type Signal interface {
	CertifyFundamental() error
}

func initialize() {
	strongmindjson.EnrollKind(&FreshIterationPhaseSignal{}, "REDACTED")
	strongmindjson.EnrollKind(&FreshSoundLedgerSignal{}, "REDACTED")
	strongmindjson.EnrollKind(&NominationSignal{}, "REDACTED")
	strongmindjson.EnrollKind(&NominationPolicySignal{}, "REDACTED")
	strongmindjson.EnrollKind(&LedgerFragmentSignal{}, "REDACTED")
	strongmindjson.EnrollKind(&BallotSignal{}, "REDACTED")
	strongmindjson.EnrollKind(&OwnsBallotSignal{}, "REDACTED")
	strongmindjson.EnrollKind(&BallotAssignMajor23signal{}, "REDACTED")
	strongmindjson.EnrollKind(&BallotAssignDigitsSignal{}, "REDACTED")
}

//

//
//
type FreshIterationPhaseSignal struct {
	Altitude                int64
	Iteration                 int32
	Phase                  controlkinds.IterationPhaseKind
	MomentsBecauseInitiateMoment int64
	FinalEndorseIteration       int32
}

//
func (m *FreshIterationPhaseSignal) CertifyFundamental() error {
	if m.Altitude < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if m.Iteration < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if !m.Phase.EqualsSound() {
		return strongminderrors.FaultUnfitAttribute{Attribute: "REDACTED"}
	}

	//

	//
	//
	//
	if m.FinalEndorseIteration < -1 {
		return strongminderrors.FaultUnfitAttribute{Attribute: "REDACTED", Rationale: "REDACTED"}
	}

	return nil
}

//
func (m *FreshIterationPhaseSignal) CertifyAltitude(primaryAltitude int64) error {
	if m.Altitude < primaryAltitude {
		return strongminderrors.FaultUnfitAttribute{
			Attribute:  "REDACTED",
			Rationale: fmt.Sprintf("REDACTED", m.Altitude, primaryAltitude),
		}
	}

	if m.Altitude == primaryAltitude && m.FinalEndorseIteration != -1 {
		return strongminderrors.FaultUnfitAttribute{
			Attribute:  "REDACTED",
			Rationale: fmt.Sprintf("REDACTED", m.FinalEndorseIteration, primaryAltitude),
		}
	}

	if m.Altitude > primaryAltitude && m.FinalEndorseIteration < 0 {
		return strongminderrors.FaultUnfitAttribute{
			Attribute:  "REDACTED",
			Rationale: fmt.Sprintf("REDACTED", primaryAltitude),
		}
	}
	return nil
}

//
func (m *FreshIterationPhaseSignal) Text() string {
	return fmt.Sprintf("REDACTED",
		m.Altitude, m.Iteration, m.Phase, m.FinalEndorseIteration)
}

//

//
//
//
type FreshSoundLedgerSignal struct {
	Altitude             int64
	Iteration              int32
	LedgerFragmentAssignHeading kinds.FragmentAssignHeading
	LedgerFragments         *digits.DigitCollection
	EqualsEndorse           bool
}

//
func (m *FreshSoundLedgerSignal) CertifyFundamental() error {
	if m.Altitude < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if m.Iteration < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if err := m.LedgerFragmentAssignHeading.CertifyFundamental(); err != nil {
		return strongminderrors.FaultIncorrectAttribute{Attribute: "REDACTED", Err: err}
	}
	if err := m.LedgerFragments.CertifyFundamental(); err != nil {
		return strongminderrors.FaultIncorrectAttribute{Attribute: "REDACTED", Err: err}
	}
	if m.LedgerFragments.Extent() == 0 {
		return strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	}
	if m.LedgerFragments.Extent() != int(m.LedgerFragmentAssignHeading.Sum) {
		return fmt.Errorf("REDACTED",
			m.LedgerFragments.Extent(),
			m.LedgerFragmentAssignHeading.Sum)
	}
	if m.LedgerFragments.Extent() > int(kinds.MaximumLedgerFragmentsTally) {
		return fmt.Errorf("REDACTED", m.LedgerFragments.Extent(), kinds.MaximumLedgerFragmentsTally)
	}
	return nil
}

//
func (m *FreshSoundLedgerSignal) Text() string {
	return fmt.Sprintf("REDACTED",
		m.Altitude, m.Iteration, m.LedgerFragmentAssignHeading, m.LedgerFragments, m.EqualsEndorse)
}

//

//
type NominationSignal struct {
	Nomination *kinds.Nomination
}

//
func (m *NominationSignal) CertifyFundamental() error {
	return m.Nomination.CertifyFundamental()
}

//
//
func (m *NominationSignal) CertifyLedgerExtent(maximumLedgerExtentOctets int64) error {
	return m.Nomination.CertifyLedgerExtent(maximumLedgerExtentOctets)
}

//
func (m *NominationSignal) Text() string {
	return fmt.Sprintf("REDACTED", m.Nomination)
}

//

//
type NominationPolicySignal struct {
	Altitude           int64
	NominationPolicyIteration int32
	NominationPolicy      *digits.DigitCollection
}

//
func (m *NominationPolicySignal) CertifyFundamental() error {
	if m.Altitude < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if m.NominationPolicyIteration < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if err := m.NominationPolicy.CertifyFundamental(); err != nil {
		return strongminderrors.FaultIncorrectAttribute{Attribute: "REDACTED", Err: err}
	}
	if m.NominationPolicy.Extent() == 0 {
		return strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	}
	if m.NominationPolicy.Extent() > kinds.MaximumBallotsTally {
		return fmt.Errorf("REDACTED", m.NominationPolicy.Extent(), kinds.MaximumBallotsTally)
	}
	return nil
}

//
func (m *NominationPolicySignal) Text() string {
	return fmt.Sprintf("REDACTED", m.Altitude, m.NominationPolicyIteration, m.NominationPolicy)
}

//

//
type LedgerFragmentSignal struct {
	Altitude int64
	Iteration  int32
	Fragment   *kinds.Fragment
}

//
func (m *LedgerFragmentSignal) CertifyFundamental() error {
	if m.Altitude < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if m.Iteration < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if err := m.Fragment.CertifyFundamental(); err != nil {
		return strongminderrors.FaultIncorrectAttribute{Attribute: "REDACTED", Err: err}
	}
	return nil
}

//
func (m *LedgerFragmentSignal) Text() string {
	return fmt.Sprintf("REDACTED", m.Altitude, m.Iteration, m.Fragment)
}

//

//
type BallotSignal struct {
	Ballot *kinds.Ballot
}

//
func (m *BallotSignal) CertifyFundamental() error {
	return m.Ballot.CertifyFundamental()
}

//
func (m *BallotSignal) Text() string {
	return fmt.Sprintf("REDACTED", m.Ballot)
}

//

//
type OwnsBallotSignal struct {
	Altitude int64
	Iteration  int32
	Kind   commitchema.AttestedSignalKind
	Ordinal  int32
}

//
func (m *OwnsBallotSignal) CertifyFundamental() error {
	if m.Altitude < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if m.Iteration < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if !kinds.EqualsBallotKindSound(m.Kind) {
		return strongminderrors.FaultUnfitAttribute{Attribute: "REDACTED"}
	}
	if m.Ordinal < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	return nil
}

//
func (m *OwnsBallotSignal) Text() string {
	return fmt.Sprintf("REDACTED", m.Ordinal, m.Altitude, m.Iteration, m.Kind)
}

//

//
type BallotAssignMajor23signal struct {
	Altitude  int64
	Iteration   int32
	Kind    commitchema.AttestedSignalKind
	LedgerUUID kinds.LedgerUUID
}

//
func (m *BallotAssignMajor23signal) CertifyFundamental() error {
	if m.Altitude < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if m.Iteration < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if !kinds.EqualsBallotKindSound(m.Kind) {
		return strongminderrors.FaultUnfitAttribute{Attribute: "REDACTED"}
	}
	if err := m.LedgerUUID.CertifyFundamental(); err != nil {
		return strongminderrors.FaultIncorrectAttribute{Attribute: "REDACTED", Err: err}
	}
	return nil
}

//
func (m *BallotAssignMajor23signal) Text() string {
	return fmt.Sprintf("REDACTED", m.Altitude, m.Iteration, m.Kind, m.LedgerUUID)
}

//

//
type BallotAssignDigitsSignal struct {
	Altitude  int64
	Iteration   int32
	Kind    commitchema.AttestedSignalKind
	LedgerUUID kinds.LedgerUUID
	Ballots   *digits.DigitCollection
}

//
func (m *BallotAssignDigitsSignal) CertifyFundamental() error {
	if m.Altitude < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if !kinds.EqualsBallotKindSound(m.Kind) {
		return strongminderrors.FaultUnfitAttribute{Attribute: "REDACTED"}
	}
	if err := m.LedgerUUID.CertifyFundamental(); err != nil {
		return strongminderrors.FaultIncorrectAttribute{Attribute: "REDACTED", Err: err}
	}
	if err := m.Ballots.CertifyFundamental(); err != nil {
		return strongminderrors.FaultIncorrectAttribute{Attribute: "REDACTED", Err: err}
	}
	//
	if m.Ballots.Extent() > kinds.MaximumBallotsTally {
		return fmt.Errorf("REDACTED", m.Ballots.Extent(), kinds.MaximumBallotsTally)
	}
	return nil
}

//
func (m *BallotAssignDigitsSignal) Text() string {
	return fmt.Sprintf("REDACTED", m.Altitude, m.Iteration, m.Kind, m.LedgerUUID, m.Ballots)
}

//

//
type OwnsNominationLedgerFragmentSignal struct {
	Altitude int64
	Iteration  int32
	Ordinal  int32
}

//
func (m *OwnsNominationLedgerFragmentSignal) CertifyFundamental() error {
	if m.Altitude < 1 {
		return strongminderrors.FaultUnfitAttribute{Attribute: "REDACTED", Rationale: "REDACTED"}
	}
	if m.Iteration < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	if m.Ordinal < 0 {
		return strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}
	}
	return nil
}

//
func (m *OwnsNominationLedgerFragmentSignal) Text() string {
	return fmt.Sprintf("REDACTED", m.Ordinal, m.Altitude, m.Iteration)
}
