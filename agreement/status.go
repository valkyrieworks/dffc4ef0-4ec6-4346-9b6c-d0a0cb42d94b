package agreement

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sort"
	"time"

	"github.com/cosmos/gogoproto/proto"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	controlkinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	strongmindincidents "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/incidents"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/abort"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

var signalStagingExtent = 1000

//
type signalDetails struct {
	Msg    Signal `json:"msg"`
	NodeUUID p2p.ID  `json:"node_token"`
}

//
type deadlineDetails struct {
	Interval time.Duration         `json:"interval"`
	Altitude   int64                 `json:"altitude"`
	Iteration    int32                 `json:"iteration"`
	Phase     controlkinds.IterationPhaseKind `json:"phase"`
}

func (ti *deadlineDetails) Text() string {
	return fmt.Sprintf("REDACTED", ti.Interval, ti.Altitude, ti.Iteration, ti.Phase)
}

//
type transferObserver interface {
	TransAccessible() <-chan struct{}
}

//
type proofHub interface {
	//
	DiscloseDiscordantBallots(ballotAN, ballotBYTE *kinds.Ballot)
}

//
//
//
//
type Status struct {
	facility.FoundationFacility

	//
	settings        *cfg.AgreementSettings
	privateAssessor kinds.PrivateAssessor //

	//
	ledgerDepot sm.LedgerDepot

	//
	ledgerExecute *sm.LedgerHandler

	//
	transferObserver transferObserver

	//
	//
	incidentpool proofHub

	//
	mtx commitchronize.ReadwriteExclusion
	controlkinds.IterationStatus
	status sm.Status //
	//
	//
	privateAssessorPublicToken security.PublicToken

	//
	//
	nodeSignalStaging     chan signalDetails
	intrinsicSignalStaging chan signalDetails
	deadlineMetronome    DeadlineMetronome

	//
	//
	metricsSignalStaging chan signalDetails

	//
	//
	incidentChannel *kinds.IncidentChannel

	//
	//
	wal          WAL
	reenactStyle   bool //
	performJournalOvertake bool //

	//
	nthPhases int

	//
	resolveNomination func(altitude int64, iteration int32)
	performPreballot      func(altitude int64, iteration int32)
	assignNomination    func(nomination *kinds.Nomination) error

	//
	complete chan struct{}

	//
	//
	incidentctl strongmindincidents.IncidentRouter

	//
	telemetry *Telemetry

	//
	inactiveStatusChronizeAltitude int64
}

//
type StatusSelection func(*Status)

//
func FreshStatus(
	settings *cfg.AgreementSettings,
	status sm.Status,
	ledgerExecute *sm.LedgerHandler,
	ledgerDepot sm.LedgerDepot,
	transferObserver transferObserver,
	incidentpool proofHub,
	choices ...StatusSelection,
) *Status {
	cs := &Status{
		settings:           settings,
		ledgerExecute:        ledgerExecute,
		ledgerDepot:       ledgerDepot,
		transferObserver:       transferObserver,
		nodeSignalStaging:     make(chan signalDetails, signalStagingExtent),
		intrinsicSignalStaging: make(chan signalDetails, signalStagingExtent),
		deadlineMetronome:    FreshDeadlineMetronome(),
		metricsSignalStaging:    make(chan signalDetails, signalStagingExtent),
		complete:             make(chan struct{}),
		performJournalOvertake:     true,
		wal:              voidJournal{},
		incidentpool:           incidentpool,
		incidentctl:             strongmindincidents.FreshIncidentRouter(),
		telemetry:          NooperationTelemetry(),
	}
	for _, selection := range choices {
		selection(cs)
	}
	//
	cs.resolveNomination = cs.fallbackResolveNomination
	cs.performPreballot = cs.fallbackPerformPreballot
	cs.assignNomination = cs.fallbackAssignNomination

	//
	if status.FinalLedgerAltitude > 0 {
		//
		//
		//
		//
		//
		if cs.inactiveStatusChronizeAltitude != 0 {
			cs.rebuildObservedEndorse(status)
		} else {
			cs.rebuildFinalEndorse(status)
		}
	}

	cs.reviseTowardStatus(status)

	//

	cs.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", cs)

	return cs
}

//
func (cs *Status) AssignTracer(l log.Tracer) {
	cs.Tracer = l
	cs.deadlineMetronome.AssignTracer(l)
}

//
func (cs *Status) AssignIncidentChannel(b *kinds.IncidentChannel) {
	cs.incidentChannel = b
	cs.ledgerExecute.AssignIncidentChannel(b)
}

//
func StatusTelemetry(telemetry *Telemetry) StatusSelection {
	return func(cs *Status) { cs.telemetry = telemetry }
}

//
//
func InactiveStatusChronizeAltitude(altitude int64) StatusSelection {
	return func(cs *Status) { cs.inactiveStatusChronizeAltitude = altitude }
}

//
func (cs *Status) Text() string {
	//
	return "REDACTED"
}

//
func (cs *Status) ObtainStatus() sm.Status {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return cs.status.Duplicate()
}

//
//
func (cs *Status) ObtainFinalAltitude() int64 {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return cs.Altitude - 1
}

//
//
func (cs *Status) ObtainIterationStatus() *controlkinds.IterationStatus {
	cs.mtx.RLock()
	rs := cs.obtainIterationStatus()
	cs.mtx.RUnlock()
	return &rs
}

//
//
func (cs *Status) obtainIterationStatus() controlkinds.IterationStatus {
	return cs.IterationStatus //
}

//
func (cs *Status) ObtainIterationStatusJSN() ([]byte, error) {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return strongmindjson.Serialize(cs.IterationStatus)
}

//
func (cs *Status) ObtainIterationStatusPlainJSN() ([]byte, error) {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return strongmindjson.Serialize(cs.IterationStatusPlain())
}

//
func (cs *Status) ObtainAssessors() (int64, []*kinds.Assessor) {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return cs.status.FinalLedgerAltitude, cs.status.Assessors.Duplicate().Assessors
}

//
//
func (cs *Status) AssignPrivateAssessor(private kinds.PrivateAssessor) {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	cs.privateAssessor = private

	if err := cs.revisePrivateAssessorPublicToken(); err != nil {
		cs.Tracer.Failure("REDACTED", "REDACTED", err)
	}
}

//
//
func (cs *Status) AssignDeadlineMetronome(deadlineMetronome DeadlineMetronome) {
	cs.mtx.Lock()
	cs.deadlineMetronome = deadlineMetronome
	cs.mtx.Unlock()
}

//
func (cs *Status) FetchEndorse(altitude int64) *kinds.Endorse {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()

	if altitude == cs.ledgerDepot.Altitude() {
		return cs.ledgerDepot.FetchObservedEndorse(altitude)
	}

	return cs.ledgerDepot.FetchLedgerEndorse(altitude)
}

//
//
func (cs *Status) UponInitiate() error {
	//
	//
	if _, ok := cs.wal.(voidJournal); ok {
		if err := cs.fetchJournalRecord(); err != nil {
			return err
		}
	}

	//
	//
	//
	//
	//
	if err := cs.deadlineMetronome.Initiate(); err != nil {
		return err
	}

	//
	//
	if cs.performJournalOvertake {
		remedyEndeavored := false

	Cycle:
		for {
			err := cs.overtakeReenact(cs.Altitude)
			switch {
			case err == nil:
				break Cycle

			case !EqualsDataImpairmentFailure(err):
				cs.Tracer.Failure("REDACTED", "REDACTED", err)
				break Cycle

			case remedyEndeavored:
				return err
			}

			cs.Tracer.Failure("REDACTED", "REDACTED", err)

			//
			if err := cs.wal.Halt(); err != nil {
				return err
			}

			remedyEndeavored = true

			//
			taintedRecord := fmt.Sprintf("REDACTED", cs.settings.JournalRecord())
			if err := strongos.DuplicateRecord(cs.settings.JournalRecord(), taintedRecord); err != nil {
				return err
			}

			cs.Tracer.Diagnose("REDACTED", "REDACTED", cs.settings.JournalRecord(), "REDACTED", taintedRecord)

			//
			if err := remedyJournalRecord(taintedRecord, cs.settings.JournalRecord()); err != nil {
				cs.Tracer.Failure("REDACTED", "REDACTED", err)
				return err
			}

			cs.Tracer.Details("REDACTED")

			//
			if err := cs.fetchJournalRecord(); err != nil {
				return err
			}
		}
	}

	if err := cs.incidentctl.Initiate(); err != nil {
		return err
	}

	//
	if err := cs.inspectDuplicateNotatingPeril(cs.Altitude); err != nil {
		return err
	}

	//
	go cs.acceptProcedure(0)

	//
	//
	rs := cs.ObtainIterationStatus()
	cs.timelineCycle0(rs)

	return nil
}

//
//
func (cs *Status) initiateThreads(maximumPhases int) {
	err := cs.deadlineMetronome.Initiate()
	if err != nil {
		cs.Tracer.Failure("REDACTED", "REDACTED", err)
		return
	}

	go cs.acceptProcedure(maximumPhases)
}

//
func (cs *Status) fetchJournalRecord() error {
	wal, err := cs.UnsealJournal(cs.settings.JournalRecord())
	if err != nil {
		cs.Tracer.Failure("REDACTED", "REDACTED", err)
		return err
	}

	cs.wal = wal
	return nil
}

//
func (cs *Status) UponHalt() {
	if err := cs.incidentctl.Halt(); err != nil {
		cs.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	if err := cs.deadlineMetronome.Halt(); err != nil {
		cs.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	//
}

//
//
//
func (cs *Status) Await() {
	<-cs.complete
}

//
//
func (cs *Status) UnsealJournal(journalRecord string) (WAL, error) {
	wal, err := FreshJournal(journalRecord)
	if err != nil {
		cs.Tracer.Failure("REDACTED", "REDACTED", journalRecord, "REDACTED", err)
		return nil, err
	}

	wal.AssignTracer(cs.Tracer.Using("REDACTED", journalRecord))

	if err := wal.Initiate(); err != nil {
		cs.Tracer.Failure("REDACTED", "REDACTED", err)
		return nil, err
	}

	return wal, nil
}

//
//
//
//
//
//

//
func (cs *Status) AppendBallot(ballot *kinds.Ballot, nodeUUID p2p.ID) (appended bool, err error) {
	if nodeUUID == "REDACTED" {
		cs.intrinsicSignalStaging <- signalDetails{&BallotSignal{ballot}, "REDACTED"}
	} else {
		cs.nodeSignalStaging <- signalDetails{&BallotSignal{ballot}, nodeUUID}
	}

	//
	return false, nil
}

//
func (cs *Status) AssignNomination(nomination *kinds.Nomination, nodeUUID p2p.ID) error {
	if nodeUUID == "REDACTED" {
		cs.intrinsicSignalStaging <- signalDetails{&NominationSignal{nomination}, "REDACTED"}
	} else {
		cs.nodeSignalStaging <- signalDetails{&NominationSignal{nomination}, nodeUUID}
	}

	//
	return nil
}

//
func (cs *Status) AppendNominationLedgerFragment(altitude int64, iteration int32, fragment *kinds.Fragment, nodeUUID p2p.ID) error {
	if nodeUUID == "REDACTED" {
		cs.intrinsicSignalStaging <- signalDetails{&LedgerFragmentSignal{altitude, iteration, fragment}, "REDACTED"}
	} else {
		cs.nodeSignalStaging <- signalDetails{&LedgerFragmentSignal{altitude, iteration, fragment}, nodeUUID}
	}

	//
	return nil
}

//
func (cs *Status) AssignNominationAlsoLedger(
	nomination *kinds.Nomination,
	ledger *kinds.Ledger, //
	fragments *kinds.FragmentAssign,
	nodeUUID p2p.ID,
) error {
	//
	if err := cs.AssignNomination(nomination, nodeUUID); err != nil {
		return err
	}

	for i := 0; i < int(fragments.Sum()); i++ {
		fragment := fragments.ObtainFragment(i)
		if err := cs.AppendNominationLedgerFragment(nomination.Altitude, nomination.Iteration, fragment, nodeUUID); err != nil {
			return err
		}
	}

	return nil
}

//
//

func (cs *Status) reviseAltitude(altitude int64) {
	cs.telemetry.Altitude.Set(float64(altitude))
	cs.Altitude = altitude
}

func (cs *Status) reviseIterationPhase(iteration int32, phase controlkinds.IterationPhaseKind) {
	if !cs.reenactStyle {
		if iteration != cs.Iteration || iteration == 0 && phase == controlkinds.IterationPhaseFreshIteration {
			cs.telemetry.LabelIteration(cs.Iteration, cs.InitiateMoment)
		}
		if cs.Phase != phase {
			cs.telemetry.LabelPhase(cs.Phase)
		}
	}
	cs.Iteration = iteration
	cs.Phase = phase
}

//
func (cs *Status) timelineCycle0(rs *controlkinds.IterationStatus) {
	//
	snoozeInterval := rs.InitiateMoment.Sub(committime.Now())
	cs.timelineDeadline(snoozeInterval, rs.Altitude, 0, controlkinds.IterationPhaseFreshAltitude)
}

//
func (cs *Status) timelineDeadline(interval time.Duration, altitude int64, iteration int32, phase controlkinds.IterationPhaseKind) {
	cs.deadlineMetronome.TimelineDeadline(deadlineDetails{interval, altitude, iteration, phase})
}

//
func (cs *Status) transmitIntrinsicSignal(mi signalDetails) {
	select {
	case cs.intrinsicSignalStaging <- mi:
	default:
		//
		//
		//
		//
		cs.Tracer.Diagnose("REDACTED")
		go func() { cs.intrinsicSignalStaging <- mi }()
	}
}

//
//
//
//
func (cs *Status) rebuildObservedEndorse(status sm.Status) {
	ballots, err := cs.ballotsOriginatingObservedEndorse(status)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	cs.FinalEndorse = ballots
}

//
//
//
//
func (cs *Status) rebuildFinalEndorse(status sm.Status) {
	additionsActivated := status.AgreementSettings.Iface.BallotAdditionsActivated(status.FinalLedgerAltitude)
	if !additionsActivated {
		cs.rebuildObservedEndorse(status)
		return
	}
	ballots, err := cs.ballotsOriginatingExpandedEndorse(status)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	cs.FinalEndorse = ballots
}

func (cs *Status) ballotsOriginatingExpandedEndorse(status sm.Status) (*kinds.BallotAssign, error) {
	ec := cs.ledgerDepot.FetchLedgerExpandedEndorse(status.FinalLedgerAltitude)
	if ec == nil {
		return nil, fmt.Errorf("REDACTED", status.FinalLedgerAltitude)
	}
	if ec.Altitude != status.FinalLedgerAltitude {
		return nil, fmt.Errorf("REDACTED",
			ec.Altitude, status.FinalLedgerAltitude)
	}
	vs := ec.TowardExpandedBallotAssign(status.SuccessionUUID, status.FinalAssessors)
	if !vs.OwnsCoupleTrinityPreponderance() {
		return nil, FaultEndorseAssemblyNegationFulfilled
	}
	return vs, nil
}

func (cs *Status) ballotsOriginatingObservedEndorse(status sm.Status) (*kinds.BallotAssign, error) {
	endorse := cs.ledgerDepot.FetchObservedEndorse(status.FinalLedgerAltitude)
	if endorse == nil {
		endorse = cs.ledgerDepot.FetchLedgerEndorse(status.FinalLedgerAltitude)
	}
	if endorse == nil {
		return nil, fmt.Errorf("REDACTED", status.FinalLedgerAltitude)
	}
	if endorse.Altitude != status.FinalLedgerAltitude {
		return nil, fmt.Errorf("REDACTED",
			endorse.Altitude, status.FinalLedgerAltitude)
	}
	vs := endorse.TowardBallotAssign(status.SuccessionUUID, status.FinalAssessors)
	if !vs.OwnsCoupleTrinityPreponderance() {
		return nil, FaultEndorseAssemblyNegationFulfilled
	}
	return vs, nil
}

//
//
func (cs *Status) reviseTowardStatus(status sm.Status) {
	if cs.EndorseIteration > -1 && 0 < cs.Altitude && cs.Altitude != status.FinalLedgerAltitude {
		panic(fmt.Sprintf(
			"REDACTED",
			cs.Altitude, status.FinalLedgerAltitude,
		))
	}

	if !cs.status.EqualsBlank() {
		if cs.status.FinalLedgerAltitude > 0 && cs.status.FinalLedgerAltitude+1 != cs.Altitude {
			//
			//
			panic(fmt.Sprintf(
				"REDACTED",
				cs.status.FinalLedgerAltitude+1, cs.Altitude,
			))
		}
		if cs.status.FinalLedgerAltitude > 0 && cs.Altitude == cs.status.PrimaryAltitude {
			panic(fmt.Sprintf(
				"REDACTED",
				cs.status.FinalLedgerAltitude, cs.status.PrimaryAltitude,
			))
		}

		//
		//
		//
		//
		//
		if status.FinalLedgerAltitude <= cs.status.FinalLedgerAltitude {
			cs.Tracer.Diagnose(
				"REDACTED",
				"REDACTED", status.FinalLedgerAltitude+1,
				"REDACTED", cs.status.FinalLedgerAltitude+1,
			)
			cs.freshPhase()
			return
		}
	}

	//
	assessors := status.Assessors

	switch {
	case status.FinalLedgerAltitude == 0: //
		cs.FinalEndorse = (*kinds.BallotAssign)(nil)
	case cs.EndorseIteration > -1 && cs.Ballots != nil: //
		if !cs.Ballots.Preendorsements(cs.EndorseIteration).OwnsCoupleTrinityPreponderance() {
			panic(fmt.Sprintf(
				"REDACTED",
				status.FinalLedgerAltitude, cs.EndorseIteration, cs.Ballots.Preendorsements(cs.EndorseIteration),
			))
		}

		cs.FinalEndorse = cs.Ballots.Preendorsements(cs.EndorseIteration)

	case cs.FinalEndorse == nil:
		//
		//
		panic(fmt.Sprintf(
			"REDACTED",
			status.FinalLedgerAltitude+1,
		))
	}

	//
	altitude := status.FinalLedgerAltitude + 1
	if altitude == 1 {
		altitude = status.PrimaryAltitude
	}

	//
	cs.reviseAltitude(altitude)
	cs.reviseIterationPhase(0, controlkinds.IterationPhaseFreshAltitude)

	if cs.EndorseMoment.IsZero() {
		//
		//
		//
		//
		//
		cs.InitiateMoment = cs.settings.Endorse(committime.Now())
	} else {
		cs.InitiateMoment = cs.settings.Endorse(cs.EndorseMoment)
	}

	cs.Assessors = assessors
	cs.Nomination = nil
	cs.NominationLedger = nil
	cs.NominationLedgerFragments = nil
	cs.SecuredIteration = -1
	cs.SecuredLedger = nil
	cs.SecuredLedgerFragments = nil
	cs.SoundIteration = -1
	cs.SoundLedger = nil
	cs.SoundLedgerFragments = nil
	if status.AgreementSettings.Iface.BallotAdditionsActivated(altitude) {
		cs.Ballots = controlkinds.FreshExpandedAltitudeBallotAssign(status.SuccessionUUID, altitude, assessors)
	} else {
		cs.Ballots = controlkinds.FreshAltitudeBallotAssign(status.SuccessionUUID, altitude, assessors)
	}
	cs.EndorseIteration = -1
	cs.FinalAssessors = status.FinalAssessors
	cs.ActivatedDeadlinePreendorse = false

	cs.status = status

	//
	cs.freshPhase()
	cs.freshAgreementParameters()
}

func (cs *Status) freshPhase() {
	rs := cs.IterationStatusIncident()
	if err := cs.wal.Persist(rs); err != nil {
		cs.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	cs.nthPhases++

	//
	if cs.incidentChannel != nil {
		if err := cs.incidentChannel.BroadcastIncidentFreshIterationPhase(rs); err != nil {
			cs.Tracer.Failure("REDACTED", "REDACTED", err)
		}

		cs.incidentctl.TriggerIncident(kinds.IncidentFreshIterationPhase, cs.IterationStatus)
	}
}

//
//
func (cs *Status) freshAgreementParameters() {
	cs.incidentctl.TriggerIncident(kinds.IncidentFreshAgreementParameters, cs.status.AgreementSettings)
}

//
//

//
//
//
//
//
func (cs *Status) acceptProcedure(maximumPhases int) {
	uponQuit := func(cs *Status) {
		//
		//
		//

		//
		if err := cs.wal.Halt(); err != nil {
			cs.Tracer.Failure("REDACTED", "REDACTED", err)
		}

		cs.wal.Await()
		close(cs.complete)
	}

	defer func() {
		if r := recover(); r != nil {
			cs.Tracer.Failure("REDACTED", "REDACTED", r, "REDACTED", string(debug.Stack()))
			//
			//
			//
			//
			//
			//
			//
			//
			uponQuit(cs)
		}
	}()

	for {
		if maximumPhases > 0 {
			if cs.nthPhases >= maximumPhases {
				cs.Tracer.Diagnose("REDACTED")
				cs.nthPhases = 0
				return
			}
		}

		rs := cs.IterationStatus
		var mi signalDetails

		select {
		case <-cs.transferObserver.TransAccessible():
			cs.processTransAccessible()

		case mi = <-cs.nodeSignalStaging:
			cs.Tracer.Diagnose("REDACTED", "REDACTED", mi.NodeUUID)

			if err := cs.wal.Persist(mi); err != nil {
				cs.Tracer.Failure("REDACTED", "REDACTED", err)
			}

			//
			//
			cs.processSignal(mi)

		case mi = <-cs.intrinsicSignalStaging:
			cs.Tracer.Diagnose("REDACTED", "REDACTED", mi.NodeUUID)

			persistJournal := true

			//
			if _, ok := mi.Msg.(*absorbAttestedLedgerSolicit); ok {
				persistJournal = false
			}

			if persistJournal {
				//
				if err := cs.wal.PersistChronize(mi); err != nil {
					panic(fmt.Errorf(
						"REDACTED",
						mi, err,
					))
				}
			}

			if _, ok := mi.Msg.(*BallotSignal); ok {
				//
				//
				//
				//
				abort.Mishap() //
			}

			//
			cs.processSignal(mi)

		case ti := <-cs.deadlineMetronome.Channel(): //
			if err := cs.wal.Persist(ti); err != nil {
				cs.Tracer.Failure("REDACTED", "REDACTED", err)
			}

			//
			//
			cs.processDeadline(ti, rs)

		case <-cs.Exit():
			uponQuit(cs)
			return
		}
	}
}

//
func (cs *Status) processSignal(mi signalDetails) {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	var (
		appended bool
		err   error
	)

	msg, nodeUUID := mi.Msg, mi.NodeUUID

	cs.Tracer.Diagnose("REDACTED", "REDACTED", string(nodeUUID), "REDACTED", msg)

	switch msg := msg.(type) {
	case *NominationSignal:
		//
		//
		err = cs.assignNomination(msg.Nomination)

	case *LedgerFragmentSignal:
		//
		appended, err = cs.appendNominationLedgerFragment(msg, nodeUUID)

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
		cs.mtx.Unlock()

		cs.mtx.Lock()
		if appended && cs.NominationLedgerFragments.EqualsFinish() {
			cs.processFinishNomination(msg.Altitude)
		}
		if appended {
			cs.metricsSignalStaging <- mi
		}

		if err != nil && msg.Iteration != cs.Iteration {
			cs.Tracer.Diagnose(
				"REDACTED",
				"REDACTED", cs.Altitude,
				"REDACTED", cs.Iteration,
				"REDACTED", msg.Iteration,
			)
			err = nil
		}

	case *BallotSignal:
		//
		//
		appended, err = cs.attemptAppendBallot(msg.Ballot, nodeUUID)
		if appended {
			cs.metricsSignalStaging <- mi
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

		//
		//
		//

	case *absorbAttestedLedgerSolicit:
		cs.processAbsorbAttestedLedgerSolicit(msg)
	default:
		cs.Tracer.Failure("REDACTED", "REDACTED", fmt.Sprintf("REDACTED", msg))
		return
	}

	if err != nil {
		cs.Tracer.Failure(
			"REDACTED",
			"REDACTED", cs.Altitude,
			"REDACTED", cs.Iteration,
			"REDACTED", nodeUUID,
			"REDACTED", fmt.Sprintf("REDACTED", msg),
			"REDACTED", err,
		)
	}
}

func (cs *Status) processDeadline(ti deadlineDetails, rs controlkinds.IterationStatus) {
	cs.Tracer.Diagnose("REDACTED", "REDACTED", ti.Interval, "REDACTED", ti.Altitude, "REDACTED", ti.Iteration, "REDACTED", ti.Phase)

	//
	if ti.Altitude != rs.Altitude || ti.Iteration < rs.Iteration || (ti.Iteration == rs.Iteration && ti.Phase < rs.Phase) {
		cs.Tracer.Diagnose("REDACTED", "REDACTED", rs.Altitude, "REDACTED", rs.Iteration, "REDACTED", rs.Phase)
		return
	}

	//
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	switch ti.Phase {
	case controlkinds.IterationPhaseFreshAltitude:
		//
		//
		cs.joinFreshIteration(ti.Altitude, 0)

	case controlkinds.IterationPhaseFreshIteration:
		cs.joinNominate(ti.Altitude, ti.Iteration)

	case controlkinds.IterationPhaseNominate:
		if err := cs.incidentChannel.BroadcastIncidentDeadlineNominate(cs.IterationStatusIncident()); err != nil {
			cs.Tracer.Failure("REDACTED", "REDACTED", err)
		}

		cs.joinPreballot(ti.Altitude, ti.Iteration)

	case controlkinds.IterationPhasePreballotAwait:
		if err := cs.incidentChannel.BroadcastIncidentDeadlinePause(cs.IterationStatusIncident()); err != nil {
			cs.Tracer.Failure("REDACTED", "REDACTED", err)
		}

		cs.joinPreendorse(ti.Altitude, ti.Iteration)

	case controlkinds.IterationPhasePreendorseAwait:
		if err := cs.incidentChannel.BroadcastIncidentDeadlinePause(cs.IterationStatusIncident()); err != nil {
			cs.Tracer.Failure("REDACTED", "REDACTED", err)
		}

		cs.relayPreendorseDeadlineTelemetry(ti.Iteration)
		cs.joinPreendorse(ti.Altitude, ti.Iteration)
		cs.joinFreshIteration(ti.Altitude, ti.Iteration+1)

	default:
		panic(strongminderrors.FaultUnfitAttribute{Attribute: "REDACTED"})
	}
}

func (cs *Status) processTransAccessible() {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	//
	if cs.Iteration != 0 {
		return
	}

	switch cs.Phase {
	case controlkinds.IterationPhaseFreshAltitude: //
		if cs.requireAttestationLedger(cs.Altitude) {
			//
			return
		}

		//
		deadlineEndorse := cs.InitiateMoment.Sub(committime.Now()) + 1*time.Millisecond
		cs.timelineDeadline(deadlineEndorse, cs.Altitude, 0, controlkinds.IterationPhaseFreshIteration)

	case controlkinds.IterationPhaseFreshIteration: //
		cs.joinNominate(cs.Altitude, 0)
	}
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
//
//
func (cs *Status) joinFreshIteration(altitude int64, iteration int32) {
	tracer := cs.Tracer.Using("REDACTED", altitude, "REDACTED", iteration)

	if cs.Altitude != altitude || iteration < cs.Iteration || (cs.Iteration == iteration && cs.Phase != controlkinds.IterationPhaseFreshAltitude) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase),
		)
		return
	}

	if now := committime.Now(); cs.InitiateMoment.After(now) {
		tracer.Diagnose("REDACTED", "REDACTED", cs.InitiateMoment, "REDACTED", now)
	}

	priorAltitude, priorIteration, priorPhase := cs.Altitude, cs.Iteration, cs.Phase

	//
	assessors := cs.Assessors
	if cs.Iteration < iteration {
		assessors = assessors.Duplicate()
		assessors.AdvanceNominatorUrgency(strongarithmetic.SecureUnderInteger32(iteration, cs.Iteration))
	}

	//
	//
	//
	cs.reviseIterationPhase(iteration, controlkinds.IterationPhaseFreshIteration)
	cs.Assessors = assessors
	//
	//
	itemLocation := assessors.ObtainNominator().PublicToken.Location()
	if iteration != 0 {
		tracer.Details("REDACTED", "REDACTED", itemLocation)
		cs.Nomination = nil
		cs.NominationLedger = nil
		cs.NominationLedgerFragments = nil
	}

	tracer.Diagnose("REDACTED",
		"REDACTED", log.FreshIdleFormat("REDACTED", priorAltitude, priorIteration, priorPhase),
		"REDACTED", itemLocation,
	)

	if iteration > 0 && !cs.reenactStyle {
		cs.telemetry.LabelIterationAdvanced(priorPhase)
	}

	cs.Ballots.AssignIteration(strongarithmetic.SecureAppendInteger32(iteration, 1)) //
	cs.ActivatedDeadlinePreendorse = false

	if err := cs.incidentChannel.BroadcastIncidentFreshIteration(cs.FreshIterationIncident()); err != nil {
		cs.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	//
	//
	//
	pauseForeachTrans := cs.settings.PauseForeachTrans() && iteration == 0 && !cs.requireAttestationLedger(altitude)
	if pauseForeachTrans {
		if cs.settings.GenerateVoidLedgersDuration > 0 {
			cs.timelineDeadline(cs.settings.GenerateVoidLedgersDuration, altitude, iteration,
				controlkinds.IterationPhaseFreshIteration)
		}
	} else {
		cs.joinNominate(altitude, iteration)
	}
}

//
//
func (cs *Status) requireAttestationLedger(altitude int64) bool {
	if altitude == cs.status.PrimaryAltitude {
		return true
	}

	finalLedgerSummary := cs.ledgerDepot.FetchLedgerSummary(altitude - 1)
	if finalLedgerSummary == nil {
		//
		cs.Tracer.Details("REDACTED", "REDACTED", altitude, "REDACTED", cs.status.PrimaryAltitude)
		return true
	}

	return !bytes.Equal(cs.status.PlatformDigest, finalLedgerSummary.Heading.PlatformDigest)
}

//
//
//
//
//
//
func (cs *Status) joinNominate(altitude int64, iteration int32) {
	tracer := cs.Tracer.Using("REDACTED", altitude, "REDACTED", iteration)

	if cs.Altitude != altitude || iteration < cs.Iteration || (cs.Iteration == iteration && controlkinds.IterationPhaseNominate <= cs.Phase) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase),
		)
		return
	}

	tracer.Diagnose("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase))

	defer func() {
		//
		cs.reviseIterationPhase(iteration, controlkinds.IterationPhaseNominate)
		cs.freshPhase()

		//
		//
		//
		if cs.equalsNominationFinish() {
			cs.joinPreballot(altitude, cs.Iteration)
		}
	}()

	//
	cs.timelineDeadline(cs.settings.Nominate(iteration), altitude, iteration, controlkinds.IterationPhaseNominate)

	//
	if cs.privateAssessor == nil {
		tracer.Diagnose("REDACTED")
		return
	}

	tracer.Diagnose("REDACTED")

	if cs.privateAssessorPublicToken == nil {
		//
		//
		tracer.Failure("REDACTED", "REDACTED", FaultPublicTokenEqualsNegationAssign)
		return
	}

	location := cs.privateAssessorPublicToken.Location()

	//
	if !cs.Assessors.OwnsLocation(location) {
		tracer.Diagnose("REDACTED", "REDACTED", location, "REDACTED", cs.Assessors)
		return
	}

	if cs.equalsNominator(location) {
		tracer.Diagnose("REDACTED", "REDACTED", location)
		cs.resolveNomination(altitude, iteration)
	} else {
		tracer.Diagnose("REDACTED", "REDACTED", cs.Assessors.ObtainNominator().Location)
	}
}

func (cs *Status) equalsNominator(location []byte) bool {
	return bytes.Equal(cs.Assessors.ObtainNominator().Location, location)
}

func (cs *Status) fallbackResolveNomination(altitude int64, iteration int32) {
	var ledger *kinds.Ledger
	var ledgerFragments *kinds.FragmentAssign

	//
	if cs.SoundLedger != nil {
		//
		ledger, ledgerFragments = cs.SoundLedger, cs.SoundLedgerFragments
	} else {
		//
		var err error
		ledger, err = cs.generateNominationLedger(context.TODO())
		if err != nil {
			cs.Tracer.Failure("REDACTED", "REDACTED", err)
			return
		} else if ledger == nil {
			panic("REDACTED")
		}
		cs.telemetry.NominationGenerateTally.Add(1)
		ledgerFragments, err = ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
		if err != nil {
			cs.Tracer.Failure("REDACTED", "REDACTED", err)
			return
		}
	}

	//
	//
	if err := cs.wal.PurgeAlsoChronize(); err != nil {
		cs.Tracer.Failure("REDACTED")
	}

	//
	itemLedgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: ledgerFragments.Heading()}
	nomination := kinds.FreshNomination(altitude, iteration, cs.SoundIteration, itemLedgerUUID)
	p := nomination.TowardSchema()
	if err := cs.privateAssessor.AttestNomination(cs.status.SuccessionUUID, p); err == nil {
		nomination.Notation = p.Notation

		//
		cs.transmitIntrinsicSignal(signalDetails{&NominationSignal{nomination}, "REDACTED"})

		for i := 0; i < int(ledgerFragments.Sum()); i++ {
			fragment := ledgerFragments.ObtainFragment(i)
			cs.transmitIntrinsicSignal(signalDetails{&LedgerFragmentSignal{cs.Altitude, cs.Iteration, fragment}, "REDACTED"})
		}

		cs.Tracer.Diagnose("REDACTED", "REDACTED", altitude, "REDACTED", iteration, "REDACTED", nomination)
	} else if !cs.reenactStyle {
		cs.Tracer.Failure("REDACTED", "REDACTED", altitude, "REDACTED", iteration, "REDACTED", err)
	}
}

//
//
func (cs *Status) equalsNominationFinish() bool {
	if cs.Nomination == nil || cs.NominationLedger == nil {
		return false
	}
	//
	//
	if cs.Nomination.PolicyIteration < 0 {
		return true
	}
	//
	return cs.Ballots.Preballots(cs.Nomination.PolicyIteration).OwnsCoupleTrinityPreponderance()
}

//
//
//
//
//
//
//
func (cs *Status) generateNominationLedger(ctx context.Context) (*kinds.Ledger, error) {
	if cs.privateAssessor == nil {
		return nil, FaultVoidPrivateAssessor
	}

	//
	var finalAddnEndorse *kinds.ExpandedEndorse
	switch {
	case cs.Altitude == cs.status.PrimaryAltitude:
		//
		//
		finalAddnEndorse = &kinds.ExpandedEndorse{}

	case cs.FinalEndorse.OwnsCoupleTrinityPreponderance():
		//
		finalAddnEndorse = cs.FinalEndorse.CreateExpandedEndorse(cs.status.AgreementSettings.Iface)

	default: //
		return nil, FaultNominationLackingPrecedingEndorse
	}

	if cs.privateAssessorPublicToken == nil {
		//
		//
		return nil, fmt.Errorf("REDACTED", FaultPublicTokenEqualsNegationAssign)
	}

	nominatorLocation := cs.privateAssessorPublicToken.Location()

	ret, err := cs.ledgerExecute.GenerateNominationLedger(ctx, cs.Altitude, cs.status, finalAddnEndorse, nominatorLocation)
	if err != nil {
		panic(err)
	}
	return ret, nil
}

//
//
//
//
func (cs *Status) joinPreballot(altitude int64, iteration int32) {
	tracer := cs.Tracer.Using("REDACTED", altitude, "REDACTED", iteration)

	if cs.Altitude != altitude || iteration < cs.Iteration || (cs.Iteration == iteration && controlkinds.IterationPhasePreballot <= cs.Phase) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase),
		)
		return
	}

	defer func() {
		//
		cs.reviseIterationPhase(iteration, controlkinds.IterationPhasePreballot)
		cs.freshPhase()
	}()

	tracer.Diagnose("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase))

	//
	cs.performPreballot(altitude, iteration)

	//
	//
}

func (cs *Status) fallbackPerformPreballot(altitude int64, iteration int32) {
	tracer := cs.Tracer.Using("REDACTED", altitude, "REDACTED", iteration)

	//
	if cs.SecuredLedger != nil {
		tracer.Diagnose("REDACTED")
		cs.attestAppendBallot(commitchema.PreballotKind, cs.SecuredLedger.Digest(), cs.SecuredLedgerFragments.Heading(), nil)
		return
	}

	//
	if cs.NominationLedger == nil {
		tracer.Diagnose("REDACTED")
		cs.attestAppendBallot(commitchema.PreballotKind, nil, kinds.FragmentAssignHeading{}, nil)
		return
	}

	//
	err := cs.ledgerExecute.CertifyLedger(cs.status, cs.NominationLedger)
	if err != nil {
		//
		tracer.Failure("REDACTED",
			"REDACTED", err)
		cs.attestAppendBallot(commitchema.PreballotKind, nil, kinds.FragmentAssignHeading{}, nil)
		return
	}

	/**
,
s
.

e
.
s
.
*/
	equalsApplicationSound, err := cs.ledgerExecute.HandleNomination(cs.NominationLedger, cs.status)
	if err != nil {
		panic(fmt.Sprintf(
			"REDACTED", err,
		))
	}
	cs.telemetry.LabelNominationHandled(equalsApplicationSound)

	//
	if !equalsApplicationSound {
		tracer.Failure("REDACTED"+
			"REDACTED", "REDACTED", err)
		cs.attestAppendBallot(commitchema.PreballotKind, nil, kinds.FragmentAssignHeading{}, nil)
		return
	}

	//
	//
	//
	tracer.Diagnose("REDACTED")
	cs.attestAppendBallot(commitchema.PreballotKind, cs.NominationLedger.Digest(), cs.NominationLedgerFragments.Heading(), nil)
}

//
func (cs *Status) joinPreballotPause(altitude int64, iteration int32) {
	tracer := cs.Tracer.Using("REDACTED", altitude, "REDACTED", iteration)

	if cs.Altitude != altitude || iteration < cs.Iteration || (cs.Iteration == iteration && controlkinds.IterationPhasePreballotAwait <= cs.Phase) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase),
		)
		return
	}

	if !cs.Ballots.Preballots(iteration).OwnsCoupleTrinitySome() {
		panic(fmt.Sprintf(
			"REDACTED",
			altitude, iteration,
		))
	}

	tracer.Diagnose("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase))

	defer func() {
		//
		cs.reviseIterationPhase(iteration, controlkinds.IterationPhasePreballotAwait)
		cs.freshPhase()
	}()

	//
	cs.timelineDeadline(cs.settings.Preballot(iteration), altitude, iteration, controlkinds.IterationPhasePreballotAwait)
}

//
//
//
//
//
//
func (cs *Status) joinPreendorse(altitude int64, iteration int32) {
	tracer := cs.Tracer.Using("REDACTED", altitude, "REDACTED", iteration)

	if cs.Altitude != altitude || iteration < cs.Iteration || (cs.Iteration == iteration && controlkinds.IterationPhasePreendorse <= cs.Phase) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase),
		)
		return
	}

	tracer.Diagnose("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase))

	defer func() {
		//
		cs.reviseIterationPhase(iteration, controlkinds.IterationPhasePreendorse)
		cs.freshPhase()
	}()

	//
	ledgerUUID, ok := cs.Ballots.Preballots(iteration).CoupleTrinityPreponderance()

	//
	if !ok {
		if cs.SecuredLedger != nil {
			tracer.Diagnose("REDACTED")
		} else {
			tracer.Diagnose("REDACTED")
		}

		cs.attestAppendBallot(commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, nil)
		return
	}

	//
	if err := cs.incidentChannel.BroadcastIncidentSpeck(cs.IterationStatusIncident()); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
	}

	//
	policyIteration, _ := cs.Ballots.PolicyDetails()
	if policyIteration < iteration {
		panic(fmt.Sprintf("REDACTED", iteration, policyIteration))
	}

	//
	if len(ledgerUUID.Digest) == 0 {
		if cs.SecuredLedger == nil {
			tracer.Diagnose("REDACTED")
		} else {
			tracer.Diagnose("REDACTED")
			cs.SecuredIteration = -1
			cs.SecuredLedger = nil
			cs.SecuredLedgerFragments = nil

			if err := cs.incidentChannel.BroadcastIncidentRelease(cs.IterationStatusIncident()); err != nil {
				tracer.Failure("REDACTED", "REDACTED", err)
			}
		}

		cs.attestAppendBallot(commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, nil)
		return
	}

	//

	//
	if cs.SecuredLedger.DigestsToward(ledgerUUID.Digest) {
		tracer.Diagnose("REDACTED")
		cs.SecuredIteration = iteration

		if err := cs.incidentChannel.BroadcastIncidentResecure(cs.IterationStatusIncident()); err != nil {
			tracer.Failure("REDACTED", "REDACTED", err)
		}

		cs.attestAppendBallot(commitchema.PreendorseKind, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, cs.SecuredLedger)
		return
	}

	//
	if cs.NominationLedger.DigestsToward(ledgerUUID.Digest) {
		tracer.Diagnose("REDACTED", "REDACTED", ledgerUUID.Digest)

		//
		if err := cs.ledgerExecute.CertifyLedger(cs.status, cs.NominationLedger); err != nil {
			panic(fmt.Sprintf("REDACTED", err))
		}

		cs.SecuredIteration = iteration
		cs.SecuredLedger = cs.NominationLedger
		cs.SecuredLedgerFragments = cs.NominationLedgerFragments

		if err := cs.incidentChannel.BroadcastIncidentSecure(cs.IterationStatusIncident()); err != nil {
			tracer.Failure("REDACTED", "REDACTED", err)
		}

		cs.attestAppendBallot(commitchema.PreendorseKind, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, cs.NominationLedger)
		return
	}

	//
	//
	//
	tracer.Diagnose("REDACTED", "REDACTED", ledgerUUID)

	cs.SecuredIteration = -1
	cs.SecuredLedger = nil
	cs.SecuredLedgerFragments = nil

	if !cs.NominationLedgerFragments.OwnsHeading(ledgerUUID.FragmentAssignHeading) {
		cs.NominationLedger = nil
		cs.NominationLedgerFragments = kinds.FreshFragmentAssignOriginatingHeading(ledgerUUID.FragmentAssignHeading)
	}

	if err := cs.incidentChannel.BroadcastIncidentRelease(cs.IterationStatusIncident()); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
	}

	cs.attestAppendBallot(commitchema.PreendorseKind, nil, kinds.FragmentAssignHeading{}, nil)
}

//
func (cs *Status) joinPreendorsePause(altitude int64, iteration int32) {
	tracer := cs.Tracer.Using("REDACTED", altitude, "REDACTED", iteration)

	if cs.Altitude != altitude || iteration < cs.Iteration || (cs.Iteration == iteration && cs.ActivatedDeadlinePreendorse) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", cs.ActivatedDeadlinePreendorse,
			"REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration),
		)
		return
	}

	if !cs.Ballots.Preendorsements(iteration).OwnsCoupleTrinitySome() {
		panic(fmt.Sprintf(
			"REDACTED",
			altitude, iteration,
		))
	}

	tracer.Diagnose("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase))

	defer func() {
		//
		cs.ActivatedDeadlinePreendorse = true
		cs.freshPhase()
	}()

	//
	cs.timelineDeadline(cs.settings.Preendorse(iteration), altitude, iteration, controlkinds.IterationPhasePreendorseAwait)
}

//
func (cs *Status) joinEndorse(altitude int64, endorseIteration int32) {
	tracer := cs.Tracer.Using("REDACTED", altitude, "REDACTED", endorseIteration)

	if cs.Altitude != altitude || controlkinds.IterationPhaseEndorse <= cs.Phase {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase),
		)
		return
	}

	tracer.Diagnose("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase))

	defer func() {
		//
		//
		cs.reviseIterationPhase(cs.Iteration, controlkinds.IterationPhaseEndorse)
		cs.EndorseIteration = endorseIteration
		cs.EndorseMoment = committime.Now()
		cs.freshPhase()

		//
		cs.attemptCulminateEndorse(altitude)
	}()

	ledgerUUID, ok := cs.Ballots.Preendorsements(endorseIteration).CoupleTrinityPreponderance()
	if !ok {
		panic("REDACTED")
	}

	//
	//
	//
	if cs.SecuredLedger.DigestsToward(ledgerUUID.Digest) {
		tracer.Diagnose("REDACTED", "REDACTED", ledgerUUID.Digest)
		cs.NominationLedger = cs.SecuredLedger
		cs.NominationLedgerFragments = cs.SecuredLedgerFragments
	}

	//
	if !cs.NominationLedger.DigestsToward(ledgerUUID.Digest) {
		if !cs.NominationLedgerFragments.OwnsHeading(ledgerUUID.FragmentAssignHeading) {
			tracer.Details(
				"REDACTED",
				"REDACTED", log.FreshIdleLedgerDigest(cs.NominationLedger),
				"REDACTED", ledgerUUID.Digest,
			)

			//
			//
			cs.NominationLedger = nil
			cs.NominationLedgerFragments = kinds.FreshFragmentAssignOriginatingHeading(ledgerUUID.FragmentAssignHeading)

			if err := cs.incidentChannel.BroadcastIncidentSoundLedger(cs.IterationStatusIncident()); err != nil {
				tracer.Failure("REDACTED", "REDACTED", err)
			}

			cs.incidentctl.TriggerIncident(kinds.IncidentSoundLedger, cs.IterationStatus)
		}
	}
}

//
func (cs *Status) attemptCulminateEndorse(altitude int64) {
	tracer := cs.Tracer.Using("REDACTED", altitude)

	if cs.Altitude != altitude {
		panic(fmt.Sprintf("REDACTED", cs.Altitude, altitude))
	}

	ledgerUUID, ok := cs.Ballots.Preendorsements(cs.EndorseIteration).CoupleTrinityPreponderance()
	if !ok || len(ledgerUUID.Digest) == 0 {
		tracer.Failure("REDACTED")
		return
	}

	if !cs.NominationLedger.DigestsToward(ledgerUUID.Digest) {
		//
		//
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.FreshIdleLedgerDigest(cs.NominationLedger),
			"REDACTED", ledgerUUID.Digest,
		)
		return
	}

	cs.culminateEndorse(altitude)
}

//
func (cs *Status) culminateEndorse(altitude int64) {
	tracer := cs.Tracer.Using("REDACTED", altitude)

	if cs.Altitude != altitude || cs.Phase != controlkinds.IterationPhaseEndorse {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.FreshIdleFormat("REDACTED", cs.Altitude, cs.Iteration, cs.Phase),
		)
		return
	}

	cs.computePreballotSignalDeferralTelemetry()

	ledgerUUID, ok := cs.Ballots.Preendorsements(cs.EndorseIteration).CoupleTrinityPreponderance()
	ledger, ledgerFragments := cs.NominationLedger, cs.NominationLedgerFragments

	if !ok {
		panic("REDACTED")
	}
	if !ledgerFragments.OwnsHeading(ledgerUUID.FragmentAssignHeading) {
		panic("REDACTED")
	}
	if !ledger.DigestsToward(ledgerUUID.Digest) {
		panic("REDACTED")
	}

	if err := cs.ledgerExecute.CertifyLedger(cs.status, ledger); err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	cs.computePreendorseSignalDeferralTelemetry()

	tracer.Details(
		"REDACTED",
		"REDACTED", log.FreshIdleLedgerDigest(ledger),
		"REDACTED", ledger.PlatformDigest,
		"REDACTED", len(ledger.Txs),
	)
	tracer.Diagnose("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", ledger))

	abort.Mishap() //

	//
	if cs.ledgerDepot.Altitude() < ledger.Altitude {
		//
		//
		observedExpandedEndorse := cs.Ballots.Preendorsements(cs.EndorseIteration).CreateExpandedEndorse(cs.status.AgreementSettings.Iface)
		if cs.status.AgreementSettings.Iface.BallotAdditionsActivated(ledger.Altitude) {
			cs.ledgerDepot.PersistLedgerUsingExpandedEndorse(ledger, ledgerFragments, observedExpandedEndorse)
		} else {
			cs.ledgerDepot.PersistLedger(ledger, ledgerFragments, observedExpandedEndorse.TowardEndorse())
		}
	} else {
		//
		tracer.Diagnose("REDACTED", "REDACTED", ledger.Altitude)
	}

	abort.Mishap() //

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
	terminateSignal := TerminateAltitudeSignal{altitude}
	if err := cs.wal.PersistChronize(terminateSignal); err != nil { //
		panic(fmt.Sprintf(
			"REDACTED",
			terminateSignal, err,
		))
	}

	abort.Mishap() //

	//
	statusDuplicate := cs.status.Duplicate()

	//
	//
	//
	statusDuplicate, err := cs.ledgerExecute.ExecuteAttestedLedger(
		statusDuplicate,
		kinds.LedgerUUID{
			Digest:          ledger.Digest(),
			FragmentAssignHeading: ledgerFragments.Heading(),
		},
		ledger,
	)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}

	abort.Mishap() //

	//
	cs.logTelemetry(altitude, ledger)

	//
	cs.reviseTowardStatus(statusDuplicate)

	abort.Mishap() //

	//
	if err := cs.revisePrivateAssessorPublicToken(); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
	}

	//
	//
	cs.timelineCycle0(&cs.IterationStatus)

	//
	//
	//
	//
}

func (cs *Status) logTelemetry(altitude int64, ledger *kinds.Ledger) {
	cs.telemetry.Assessors.Set(float64(cs.Assessors.Extent()))
	cs.telemetry.AssessorsPotency.Set(float64(cs.Assessors.SumBallotingPotency()))

	var (
		absentAssessors      int
		absentAssessorsPotency int64
	)
	//
	//
	//
	if altitude > cs.status.PrimaryAltitude {
		//
		//
		var (
			endorseExtent = ledger.FinalEndorse.Extent()
			itemAssignLength  = len(cs.FinalAssessors.Assessors)
			location    kinds.Location
		)
		if endorseExtent != itemAssignLength {
			panic(fmt.Sprintf("REDACTED",
				endorseExtent, itemAssignLength, ledger.Altitude, ledger.FinalEndorse.Notations, cs.FinalAssessors.Assessors))
		}

		if cs.privateAssessor != nil {
			if cs.privateAssessorPublicToken == nil {
				//
				cs.Tracer.Failure(fmt.Sprintf("REDACTED", FaultPublicTokenEqualsNegationAssign))
			} else {
				location = cs.privateAssessorPublicToken.Location()
			}
		}

		for i, val := range cs.FinalAssessors.Assessors {
			endorseSignature := ledger.FinalEndorse.Notations[i]
			if endorseSignature.LedgerUUIDMarker == kinds.LedgerUUIDMarkerMissing {
				absentAssessors++
				absentAssessorsPotency += val.BallotingPotency
			}

			if bytes.Equal(val.Location, location) {
				tag := []string{
					"REDACTED", val.Location.Text(),
				}
				cs.telemetry.AssessorPotency.With(tag...).Set(float64(val.BallotingPotency))
				if endorseSignature.LedgerUUIDMarker == kinds.LedgerUUIDMarkerEndorse {
					cs.telemetry.AssessorFinalAttestedAltitude.With(tag...).Set(float64(altitude))
				} else {
					cs.telemetry.AssessorOmittedLedgers.With(tag...).Add(float64(1))
				}
			}

		}
	}
	cs.telemetry.AbsentAssessors.Set(float64(absentAssessors))
	cs.telemetry.AbsentAssessorsPotency.Set(float64(absentAssessorsPotency))

	//
	var (
		treacherousAssessorsPotency = int64(0)
		treacherousAssessorsTally = int64(0)
	)
	for _, ev := range ledger.Proof.Proof {
		if dve, ok := ev.(*kinds.ReplicatedBallotProof); ok {
			if _, val := cs.Assessors.ObtainViaLocationAlterable(dve.BallotAN.AssessorLocation); val != nil {
				treacherousAssessorsTally++
				treacherousAssessorsPotency += val.BallotingPotency
			}
		}
	}
	cs.telemetry.TreacherousAssessors.Set(float64(treacherousAssessorsTally))
	cs.telemetry.TreacherousAssessorsPotency.Set(float64(treacherousAssessorsPotency))

	if altitude > 1 {
		finalLedgerSummary := cs.ledgerDepot.FetchLedgerSummary(altitude - 1)
		if finalLedgerSummary != nil {
			cs.telemetry.LedgerDurationMoments.Observe(
				ledger.Moment.Sub(finalLedgerSummary.Heading.Moment).Seconds(),
			)
		}
	}

	cs.telemetry.CountTrans.Set(float64(len(ledger.Txs)))
	cs.telemetry.SumTrans.Add(float64(len(ledger.Txs)))
	cs.telemetry.LedgerExtentOctets.Set(float64(ledger.Extent()))
	cs.telemetry.SuccessionExtentOctets.Add(float64(ledger.Extent()))
	cs.telemetry.RatifiedAltitude.Set(float64(ledger.Altitude))
}

//

func (cs *Status) fallbackAssignNomination(nomination *kinds.Nomination) error {
	//
	//
	if cs.Nomination != nil {
		return nil
	}

	//
	if nomination.Altitude != cs.Altitude || nomination.Iteration != cs.Iteration {
		return nil
	}

	//
	if nomination.PolicyIteration < -1 ||
		(nomination.PolicyIteration >= 0 && nomination.PolicyIteration >= nomination.Iteration) {
		return FaultUnfitNominationPolicyIteration
	}

	p := nomination.TowardSchema()
	//
	publicToken := cs.Assessors.ObtainNominator().PublicToken
	if !publicToken.ValidateNotation(
		kinds.NominationAttestOctets(cs.status.SuccessionUUID, p), nomination.Notation,
	) {
		return FaultUnfitNominationNotation
	}

	//
	maximumOctets := cs.status.AgreementSettings.Ledger.MaximumOctets
	if maximumOctets == -1 {
		maximumOctets = int64(kinds.MaximumLedgerExtentOctets)
	}
	if int64(nomination.LedgerUUID.FragmentAssignHeading.Sum) > (maximumOctets-1)/int64(kinds.LedgerFragmentExtentOctets)+1 {
		return FaultNominationExcessivelyMultipleFragments
	}

	nomination.Notation = p.Notation
	cs.Nomination = nomination
	//
	//
	//
	if cs.NominationLedgerFragments == nil {
		cs.NominationLedgerFragments = kinds.FreshFragmentAssignOriginatingHeading(nomination.LedgerUUID.FragmentAssignHeading)
	}

	cs.Tracer.Details("REDACTED", "REDACTED", nomination, "REDACTED", publicToken.Location())
	return nil
}

//
//
//
func (cs *Status) appendNominationLedgerFragment(msg *LedgerFragmentSignal, nodeUUID p2p.ID) (appended bool, err error) {
	altitude, iteration, fragment := msg.Altitude, msg.Iteration, msg.Fragment

	//
	if cs.Altitude != altitude {
		cs.Tracer.Diagnose("REDACTED", "REDACTED", altitude, "REDACTED", iteration)
		cs.telemetry.LedgerMulticastFragmentsAccepted.With("REDACTED", "REDACTED").Add(1)
		return false, nil
	}

	//
	if cs.NominationLedgerFragments == nil {
		cs.telemetry.LedgerMulticastFragmentsAccepted.With("REDACTED", "REDACTED").Add(1)
		//
		//
		cs.Tracer.Diagnose(
			"REDACTED",
			"REDACTED", altitude,
			"REDACTED", iteration,
			"REDACTED", fragment.Ordinal,
			"REDACTED", nodeUUID,
		)
		return false, nil
	}

	appended, err = cs.NominationLedgerFragments.AppendFragment(fragment)
	if err != nil {
		if errors.Is(err, kinds.FaultFragmentAssignUnfitAttestation) || errors.Is(err, kinds.FaultFragmentAssignUnforeseenOrdinal) {
			cs.telemetry.LedgerMulticastFragmentsAccepted.With("REDACTED", "REDACTED").Add(1)
		}
		return appended, err
	}

	cs.telemetry.LedgerMulticastFragmentsAccepted.With("REDACTED", "REDACTED").Add(1)
	if !appended {
		//
		//
		cs.telemetry.ReplicatedLedgerFragment.Add(1)
	}

	maximumOctets := cs.status.AgreementSettings.Ledger.MaximumOctets
	if maximumOctets == -1 {
		maximumOctets = int64(kinds.MaximumLedgerExtentOctets)
	}
	if cs.NominationLedgerFragments.OctetExtent() > maximumOctets {
		return appended, fmt.Errorf("REDACTED",
			cs.NominationLedgerFragments.OctetExtent(), maximumOctets,
		)
	}
	if appended && cs.NominationLedgerFragments.EqualsFinish() {
		bz, err := io.ReadAll(cs.NominationLedgerFragments.ObtainFetcher())
		if err != nil {
			return appended, err
		}

		pbb := new(commitchema.Ledger)
		err = proto.Unmarshal(bz, pbb)
		if err != nil {
			return appended, err
		}

		ledger, err := kinds.LedgerOriginatingSchema(pbb)
		if err != nil {
			return appended, err
		}

		cs.NominationLedger = ledger

		//
		cs.Tracer.Details("REDACTED", "REDACTED", cs.NominationLedger.Altitude, "REDACTED", cs.NominationLedger.Digest())

		if err := cs.incidentChannel.BroadcastIncidentFinishNomination(cs.FinishNominationIncident()); err != nil {
			cs.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	return appended, nil
}

func (cs *Status) processFinishNomination(ledgerAltitude int64) {
	//
	preballots := cs.Ballots.Preballots(cs.Iteration)
	ledgerUUID, ownsCoupleTrinity := preballots.CoupleTrinityPreponderance()
	if ownsCoupleTrinity && !ledgerUUID.EqualsNull() && (cs.SoundIteration < cs.Iteration) {
		if cs.NominationLedger.DigestsToward(ledgerUUID.Digest) {
			cs.Tracer.Diagnose(
				"REDACTED",
				"REDACTED", cs.Iteration,
				"REDACTED", log.FreshIdleLedgerDigest(cs.NominationLedger),
			)

			cs.SoundIteration = cs.Iteration
			cs.SoundLedger = cs.NominationLedger
			cs.SoundLedgerFragments = cs.NominationLedgerFragments
		}
		//
		//
		//
		//
		//
	}

	if cs.Phase <= controlkinds.IterationPhaseNominate && cs.equalsNominationFinish() {
		//
		cs.joinPreballot(ledgerAltitude, cs.Iteration)
		if ownsCoupleTrinity { //
			cs.joinPreendorse(ledgerAltitude, cs.Iteration)
		}
	} else if cs.Phase == controlkinds.IterationPhaseEndorse {
		//
		cs.attemptCulminateEndorse(ledgerAltitude)
	}
}

//
func (cs *Status) attemptAppendBallot(ballot *kinds.Ballot, nodeUUID p2p.ID) (bool, error) {
	appended, err := cs.appendBallot(ballot, nodeUUID)
	//
	if err != nil {
		//
		//
		//
		//
		if ballotFault, ok := err.(*kinds.FaultBallotDiscordantBallots); ok {
			if cs.privateAssessorPublicToken == nil {
				return false, FaultPublicTokenEqualsNegationAssign
			}

			if bytes.Equal(ballot.AssessorLocation, cs.privateAssessorPublicToken.Location()) {
				cs.Tracer.Failure(
					"REDACTED",
					"REDACTED", ballot.Altitude,
					"REDACTED", ballot.Iteration,
					"REDACTED", ballot.Kind,
				)

				return appended, err
			}

			//
			cs.incidentpool.DiscloseDiscordantBallots(ballotFault.BallotAN, ballotFault.BallotBYTE)
			cs.Tracer.Diagnose(
				"REDACTED",
				"REDACTED", ballotFault.BallotAN,
				"REDACTED", ballotFault.BallotBYTE,
			)

			return appended, err
		} else if errors.Is(err, kinds.FaultBallotUnCertainNotation) {
			cs.Tracer.Diagnose("REDACTED", "REDACTED", err)
		} else if errors.Is(err, kinds.FaultUnfitBallotAddition) {
			cs.Tracer.Diagnose("REDACTED")
		} else {
			//
			//
			//
			//
			//
			cs.Tracer.Details("REDACTED", "REDACTED", err)
			return appended, FaultAppendingBallot
		}
	}

	return appended, nil
}

func (cs *Status) appendBallot(ballot *kinds.Ballot, nodeUUID p2p.ID) (appended bool, err error) {
	cs.Tracer.Diagnose(
		"REDACTED",
		"REDACTED", ballot.Altitude,
		"REDACTED", ballot.Kind,
		"REDACTED", ballot.AssessorOrdinal,
		"REDACTED", cs.Altitude,
		"REDACTED", len(ballot.Addition),
		"REDACTED", len(ballot.AdditionNotation),
		"REDACTED", nodeUUID,
	)

	if ballot.Altitude < cs.Altitude || (ballot.Altitude == cs.Altitude && ballot.Iteration < cs.Iteration) {
		cs.telemetry.LabelTardyBallot(ballot.Kind)
	}

	//
	//
	if ballot.Altitude+1 == cs.Altitude && ballot.Kind == commitchema.PreendorseKind {
		if cs.Phase != controlkinds.IterationPhaseFreshAltitude {
			//
			cs.Tracer.Diagnose("REDACTED", "REDACTED", ballot)
			return appended, err
		}

		appended, err = cs.FinalEndorse.AppendBallot(ballot)
		if !appended {
			//
			if err == nil {
				cs.telemetry.ReplicatedBallot.Add(1)
			}
			return appended, err
		}

		cs.Tracer.Diagnose("REDACTED", "REDACTED", cs.FinalEndorse.TextBrief())
		if err := cs.incidentChannel.BroadcastIncidentBallot(kinds.IncidentDataBallot{Ballot: ballot}); err != nil {
			return appended, err
		}

		cs.incidentctl.TriggerIncident(kinds.IncidentBallot, ballot)

		//
		if cs.settings.OmitDeadlineEndorse && cs.FinalEndorse.OwnsEvery() {
			//
			//
			cs.joinFreshIteration(cs.Altitude, 0)
		}

		return appended, err
	}

	//
	//
	if ballot.Altitude != cs.Altitude {
		cs.Tracer.Diagnose("REDACTED", "REDACTED", ballot.Altitude, "REDACTED", cs.Altitude, "REDACTED", nodeUUID)
		return appended, err
	}

	//
	addnActivated := cs.status.AgreementSettings.Iface.BallotAdditionsActivated(ballot.Altitude)
	if addnActivated {
		//
		//
		//

		var mineLocation []byte
		if cs.privateAssessorPublicToken != nil {
			mineLocation = cs.privateAssessorPublicToken.Location()
		}
		//
		//
		if ballot.Kind == commitchema.PreendorseKind && !ballot.LedgerUUID.EqualsNull() &&
			!bytes.Equal(ballot.AssessorLocation, mineLocation) { //

			//
			//
			//
			//
			_, val := cs.status.Assessors.ObtainViaOrdinal(ballot.AssessorOrdinal)
			if val == nil { //
				valuesTally := cs.status.Assessors.Extent()
				cs.Tracer.Details("REDACTED",
					"REDACTED", nodeUUID,
					"REDACTED", ballot.AssessorOrdinal,
					"REDACTED", valuesTally)
				return appended, FaultUnfitBallot{Rationale: fmt.Sprintf("REDACTED", ballot.AssessorOrdinal, valuesTally)}
			}
			if err := ballot.ValidateAddition(cs.status.SuccessionUUID, val.PublicToken); err != nil {
				return false, err
			}

			err := cs.ledgerExecute.ValidateBallotAddition(context.TODO(), ballot)
			cs.telemetry.LabelBallotAdditionAccepted(err == nil)
			if err != nil {
				return false, err
			}
		}
	} else if len(ballot.Addition) > 0 || len(ballot.AdditionNotation) > 0 {
		//
		//
		//
		//
		//
		//

		return false, fmt.Errorf("REDACTED", ballot.Altitude, nodeUUID)

	}

	altitude := cs.Altitude
	appended, err = cs.Ballots.AppendBallot(ballot, nodeUUID, addnActivated)
	if !appended {
		//

		//
		if err == nil {
			cs.telemetry.ReplicatedBallot.Add(1)
		}
		return appended, err
	}
	if ballot.Iteration == cs.Iteration {
		values := cs.status.Assessors
		_, val := values.ObtainViaOrdinal(ballot.AssessorOrdinal)
		cs.telemetry.LabelBallotAccepted(ballot.Kind, val.BallotingPotency, values.SumBallotingPotency())
	}

	if err := cs.incidentChannel.BroadcastIncidentBallot(kinds.IncidentDataBallot{Ballot: ballot}); err != nil {
		return appended, err
	}
	cs.incidentctl.TriggerIncident(kinds.IncidentBallot, ballot)

	switch ballot.Kind {
	case commitchema.PreballotKind:
		preballots := cs.Ballots.Preballots(ballot.Iteration)
		cs.Tracer.Diagnose("REDACTED", "REDACTED", ballot, "REDACTED", preballots.TextBrief())

		//
		if ledgerUUID, ok := preballots.CoupleTrinityPreponderance(); ok {
			//
			//
			//

			//
			//
			if (cs.SecuredLedger != nil) &&
				(cs.SecuredIteration < ballot.Iteration) &&
				(ballot.Iteration <= cs.Iteration) &&
				!cs.SecuredLedger.DigestsToward(ledgerUUID.Digest) {

				cs.Tracer.Diagnose("REDACTED", "REDACTED", cs.SecuredIteration, "REDACTED", ballot.Iteration)

				cs.SecuredIteration = -1
				cs.SecuredLedger = nil
				cs.SecuredLedgerFragments = nil

				if err := cs.incidentChannel.BroadcastIncidentRelease(cs.IterationStatusIncident()); err != nil {
					return appended, err
				}
			}

			//
			//
			if len(ledgerUUID.Digest) != 0 && (cs.SoundIteration < ballot.Iteration) && (ballot.Iteration == cs.Iteration) {
				if cs.NominationLedger.DigestsToward(ledgerUUID.Digest) {
					cs.Tracer.Diagnose("REDACTED", "REDACTED", cs.SoundIteration, "REDACTED", ballot.Iteration)
					cs.SoundIteration = ballot.Iteration
					cs.SoundLedger = cs.NominationLedger
					cs.SoundLedgerFragments = cs.NominationLedgerFragments
				} else {
					cs.Tracer.Diagnose(
						"REDACTED",
						"REDACTED", log.FreshIdleLedgerDigest(cs.NominationLedger),
						"REDACTED", ledgerUUID.Digest,
					)

					//
					cs.NominationLedger = nil
				}

				if !cs.NominationLedgerFragments.OwnsHeading(ledgerUUID.FragmentAssignHeading) {
					cs.NominationLedgerFragments = kinds.FreshFragmentAssignOriginatingHeading(ledgerUUID.FragmentAssignHeading)
				}

				cs.incidentctl.TriggerIncident(kinds.IncidentSoundLedger, cs.IterationStatus)
				if err := cs.incidentChannel.BroadcastIncidentSoundLedger(cs.IterationStatusIncident()); err != nil {
					return appended, err
				}
			}
		}

		//
		switch {
		case cs.Iteration < ballot.Iteration && preballots.OwnsCoupleTrinitySome():
			//
			cs.joinFreshIteration(altitude, ballot.Iteration)

		case cs.Iteration == ballot.Iteration && controlkinds.IterationPhasePreballot <= cs.Phase: //
			ledgerUUID, ok := preballots.CoupleTrinityPreponderance()
			if ok && (cs.equalsNominationFinish() || len(ledgerUUID.Digest) == 0) {
				cs.joinPreendorse(altitude, ballot.Iteration)
			} else if preballots.OwnsCoupleTrinitySome() {
				cs.joinPreballotPause(altitude, ballot.Iteration)
			}

		case cs.Nomination != nil && 0 <= cs.Nomination.PolicyIteration && cs.Nomination.PolicyIteration == ballot.Iteration:
			//
			if cs.equalsNominationFinish() {
				cs.joinPreballot(altitude, cs.Iteration)
			}
		}

	case commitchema.PreendorseKind:
		preendorsements := cs.Ballots.Preendorsements(ballot.Iteration)
		cs.Tracer.Diagnose("REDACTED",
			"REDACTED", ballot.Altitude,
			"REDACTED", ballot.Iteration,
			"REDACTED", ballot.AssessorLocation.Text(),
			"REDACTED", ballot.Timestamp,
			"REDACTED", preendorsements.RecordText())

		ledgerUUID, ok := preendorsements.CoupleTrinityPreponderance()
		if ok {
			//
			cs.joinFreshIteration(altitude, ballot.Iteration)
			cs.joinPreendorse(altitude, ballot.Iteration)

			if len(ledgerUUID.Digest) != 0 {
				cs.joinEndorse(altitude, ballot.Iteration)
				if cs.settings.OmitDeadlineEndorse && preendorsements.OwnsEvery() {
					cs.joinFreshIteration(cs.Altitude, 0)
				}
			} else {
				cs.joinPreendorsePause(altitude, ballot.Iteration)
			}
		} else if cs.Iteration <= ballot.Iteration && preendorsements.OwnsCoupleTrinitySome() {
			cs.joinFreshIteration(altitude, ballot.Iteration)
			cs.joinPreendorsePause(altitude, ballot.Iteration)
		}

	default:
		panic(fmt.Sprintf("REDACTED", ballot.Kind))
	}

	return appended, err
}

//
func (cs *Status) attestBallot(
	signalKind commitchema.AttestedSignalKind,
	digest []byte,
	heading kinds.FragmentAssignHeading,
	ledger *kinds.Ledger,
) (*kinds.Ballot, error) {
	//
	//
	if err := cs.wal.PurgeAlsoChronize(); err != nil {
		return nil, err
	}

	if cs.privateAssessorPublicToken == nil {
		return nil, FaultPublicTokenEqualsNegationAssign
	}

	location := cs.privateAssessorPublicToken.Location()
	itemOffset, _ := cs.Assessors.ObtainViaLocation(location)

	ballot := &kinds.Ballot{
		AssessorLocation: location,
		AssessorOrdinal:   itemOffset,
		Altitude:           cs.Altitude,
		Iteration:            cs.Iteration,
		Timestamp:        cs.ballotMoment(),
		Kind:             signalKind,
		LedgerUUID:          kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: heading},
	}

	addnActivated := cs.status.AgreementSettings.Iface.BallotAdditionsActivated(ballot.Altitude)
	if signalKind == commitchema.PreendorseKind && !ballot.LedgerUUID.EqualsNull() {
		//
		//
		if addnActivated {
			ext, err := cs.ledgerExecute.BroadenBallot(context.TODO(), ballot, ledger, cs.status)
			if err != nil {
				return nil, err
			}
			ballot.Addition = ext
		}
	}

	retrievable, err := kinds.AttestAlsoInspectBallot(ballot, cs.privateAssessor, cs.status.SuccessionUUID, addnActivated && (signalKind == commitchema.PreendorseKind))
	if err != nil && !retrievable {
		panic(fmt.Sprintf("REDACTED", ballot, err))
	}

	return ballot, err
}

func (cs *Status) ballotMoment() time.Time {
	now := committime.Now()
	minimumBallotMoment := now
	//
	const momentEnumerate = time.Millisecond
	//
	//
	if cs.SecuredLedger != nil {
		//
		//
		minimumBallotMoment = cs.SecuredLedger.Moment.Add(momentEnumerate)
	} else if cs.NominationLedger != nil {
		minimumBallotMoment = cs.NominationLedger.Moment.Add(momentEnumerate)
	}

	if now.After(minimumBallotMoment) {
		return now
	}
	return minimumBallotMoment
}

//
//
func (cs *Status) attestAppendBallot(
	signalKind commitchema.AttestedSignalKind,
	digest []byte,
	heading kinds.FragmentAssignHeading,
	ledger *kinds.Ledger,
) {
	if cs.privateAssessor == nil { //
		return
	}

	if cs.privateAssessorPublicToken == nil {
		//
		cs.Tracer.Failure(fmt.Sprintf("REDACTED", FaultPublicTokenEqualsNegationAssign))
		return
	}

	//
	if !cs.Assessors.OwnsLocation(cs.privateAssessorPublicToken.Location()) {
		return
	}

	//
	ballot, err := cs.attestBallot(signalKind, digest, heading, ledger)
	if err != nil {
		cs.Tracer.Failure("REDACTED", "REDACTED", cs.Altitude, "REDACTED", cs.Iteration, "REDACTED", ballot, "REDACTED", err)
		return
	}
	ownsAddn := len(ballot.AdditionNotation) > 0
	addnActivated := cs.status.AgreementSettings.Iface.BallotAdditionsActivated(ballot.Altitude)
	if ballot.Kind == commitchema.PreendorseKind && !ballot.LedgerUUID.EqualsNull() && ownsAddn != addnActivated {
		panic(fmt.Errorf("REDACTED",
			ownsAddn, addnActivated, ballot.Altitude, ballot.Kind))
	}
	cs.transmitIntrinsicSignal(signalDetails{&BallotSignal{ballot}, "REDACTED"})
	cs.Tracer.Diagnose("REDACTED", "REDACTED", cs.Altitude, "REDACTED", cs.Iteration, "REDACTED", ballot)
}

//
//
//
func (cs *Status) revisePrivateAssessorPublicToken() error {
	if cs.privateAssessor == nil {
		return nil
	}

	publicToken, err := cs.privateAssessor.ObtainPublicToken()
	if err != nil {
		return err
	}
	cs.privateAssessorPublicToken = publicToken
	return nil
}

//
func (cs *Status) inspectDuplicateNotatingPeril(altitude int64) error {
	if cs.privateAssessor != nil && cs.privateAssessorPublicToken != nil && cs.settings.DuplicateAttestInspectAltitude > 0 && altitude > 0 {
		itemLocation := cs.privateAssessorPublicToken.Location()
		duplicateAttestInspectAltitude := cs.settings.DuplicateAttestInspectAltitude
		if duplicateAttestInspectAltitude > altitude {
			duplicateAttestInspectAltitude = altitude
		}

		for i := int64(1); i < duplicateAttestInspectAltitude; i++ {
			finalEndorse := cs.ledgerDepot.FetchObservedEndorse(altitude - i)
			if finalEndorse != nil {
				for signatureOffset, s := range finalEndorse.Notations {
					if s.LedgerUUIDMarker == kinds.LedgerUUIDMarkerEndorse && bytes.Equal(s.AssessorLocation, itemLocation) {
						cs.Tracer.Details("REDACTED", "REDACTED", s, "REDACTED", signatureOffset, "REDACTED", altitude-i)
						return FaultNotationDetectedInsideElapsedLedgers
					}
				}
			}
		}
	}

	return nil
}

//
//
func (cs *Status) relayPreendorseDeadlineTelemetry(iteration int32) {
	//
	//
	sumBallotsGathered := 0
	sumBallotingPotencyGathered := int64(0)

	for _, ballot := range cs.Ballots.Preendorsements(iteration).Catalog() {
		sumBallotsGathered++
		_, val := cs.Assessors.ObtainViaLocation(ballot.AssessorLocation)
		if val != nil {
			sumBallotingPotencyGathered += val.BallotingPotency
		}
	}

	//
	sumFeasibleBallotingPotency := cs.Assessors.SumBallotingPotency()
	var equityFraction float64
	if sumFeasibleBallotingPotency > 0 {
		equityFraction = float64(sumBallotingPotencyGathered) / float64(sumFeasibleBallotingPotency)
	}

	//
	cs.telemetry.PreendorsementsTallied.Set(float64(sumBallotsGathered))
	cs.telemetry.PreendorsementsPledgingFraction.Set(equityFraction)

	cs.Tracer.Diagnose("REDACTED",
		"REDACTED", sumBallotsGathered,
		"REDACTED", equityFraction)
}

func (cs *Status) computePreendorseSignalDeferralTelemetry() {
	if cs.Nomination == nil {
		return
	}

	ps := cs.Ballots.Preendorsements(cs.Iteration)
	pl := ps.Catalog()

	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Timestamp.Before(pl[j].Timestamp)
	})

	var ballotingPotencyObserved int64
	for _, v := range pl {
		_, val := cs.Assessors.ObtainViaLocation(v.AssessorLocation)
		ballotingPotencyObserved += val.BallotingPotency
		if ballotingPotencyObserved >= cs.Assessors.SumBallotingPotency()*2/3+1 {
			cs.telemetry.AssemblyPreendorseDeferral.With("REDACTED", cs.Assessors.ObtainNominator().Location.Text()).Set(v.Timestamp.Sub(cs.Nomination.Timestamp).Seconds())
			break
		}
	}
}

func (cs *Status) computePreballotSignalDeferralTelemetry() {
	if cs.Nomination == nil {
		return
	}

	ps := cs.Ballots.Preballots(cs.Iteration)
	pl := ps.Catalog()

	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Timestamp.Before(pl[j].Timestamp)
	})

	var ballotingPotencyObserved int64
	for _, v := range pl {
		_, val := cs.Assessors.ObtainViaLocationAlterable(v.AssessorLocation)
		ballotingPotencyObserved += val.BallotingPotency
		if ballotingPotencyObserved >= cs.Assessors.SumBallotingPotency()*2/3+1 {
			cs.telemetry.AssemblyPreballotDeferral.With("REDACTED", cs.Assessors.ObtainNominator().Location.Text()).Set(v.Timestamp.Sub(cs.Nomination.Timestamp).Seconds())
			break
		}
	}
	if ps.OwnsEvery() {
		cs.telemetry.CompletePreballotDeferral.With("REDACTED", cs.Assessors.ObtainNominator().Location.Text()).Set(pl[len(pl)-1].Timestamp.Sub(cs.Nomination.Timestamp).Seconds())
	}
}

//

func ContrastHours(h1 int64, r1 int32, s1 controlkinds.IterationPhaseKind, h2 int64, r2 int32, s2 controlkinds.IterationPhaseKind) int {
	if h1 < h2 {
		return -1
	} else if h1 > h2 {
		return 1
	}
	if r1 < r2 {
		return -1
	} else if r1 > r2 {
		return 1
	}
	if s1 < s2 {
		return -1
	} else if s1 > s2 {
		return 1
	}
	return 0
}

//
//
func remedyJournalRecord(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	var (
		dec = FreshJournalDeserializer(in)
		enc = FreshJournalSerializer(out)
	)

	//
	for {
		msg, err := dec.Deserialize()
		if err != nil {
			break
		}

		err = enc.Serialize(msg)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}

	return nil
}
