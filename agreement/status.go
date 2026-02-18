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

	cfg "github.com/valkyrieworks/settings"
	cskinds "github.com/valkyrieworks/agreement/kinds"
	"github.com/valkyrieworks/vault"
	cometsignals "github.com/valkyrieworks/utils/events"
	"github.com/valkyrieworks/utils/abort"
	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/utils/log"
	cometmath "github.com/valkyrieworks/utils/math"
	cometos "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/p2p"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
	cometfaults "github.com/valkyrieworks/kinds/faults"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

var messageBufferVolume = 1000

//
type messageDetails struct {
	Msg    Signal `json:"msg"`
	NodeUID p2p.ID  `json:"node_key"`
}

//
type deadlineDetails struct {
	Period time.Duration         `json:"period"`
	Level   int64                 `json:"level"`
	Cycle    int32                 `json:"duration"`
	Phase     cskinds.DurationPhaseKind `json:"phase"`
}

func (ti *deadlineDetails) String() string {
	return fmt.Sprintf("REDACTED", ti.Period, ti.Level, ti.Cycle, ti.Phase)
}

//
type transferAlerter interface {
	TransAccessible() <-chan struct{}
}

//
type proofDepository interface {
	//
	NotifyClashingBallots(ballotA, ballotBYTE *kinds.Ballot)
}

//
//
//
//
type Status struct {
	daemon.RootDaemon

	//
	settings        *cfg.AgreementSettings
	privateRatifier kinds.PrivateRatifier //

	//
	ledgerDepot sm.LedgerDepot

	//
	ledgerExecute *sm.LedgerRunner

	//
	transferAlerter transferAlerter

	//
	//
	eventpool proofDepository

	//
	mtx engineconnect.ReadwriteLock
	cskinds.DurationStatus
	status sm.Status //
	//
	//
	privateRatifierPublicKey vault.PublicKey

	//
	//
	nodeMessageBuffer     chan messageDetails
	intrinsicMessageBuffer chan messageDetails
	deadlineTimer    DeadlineTimer

	//
	//
	metricsMessageBuffer chan messageDetails

	//
	//
	eventBus *kinds.EventBus

	//
	//
	wal          WAL
	resimulateStyle   bool //
	executeJournalOvertake bool //

	//
	nPhases int

	//
	determineNomination func(level int64, duration int32)
	doPreballot      func(level int64, duration int32)
	collectionNomination    func(nomination *kinds.Nomination) error

	//
	done chan struct{}

	//
	//
	evsw cometsignals.EventRouter

	//
	stats *Stats

	//
	inactiveStatusAlignLevel int64
}

//
type StatusSetting func(*Status)

//
func NewStatus(
	settings *cfg.AgreementSettings,
	status sm.Status,
	ledgerExecute *sm.LedgerRunner,
	ledgerDepot sm.LedgerDepot,
	transferAlerter transferAlerter,
	eventpool proofDepository,
	options ...StatusSetting,
) *Status {
	cs := &Status{
		settings:           settings,
		ledgerExecute:        ledgerExecute,
		ledgerDepot:       ledgerDepot,
		transferAlerter:       transferAlerter,
		nodeMessageBuffer:     make(chan messageDetails, messageBufferVolume),
		intrinsicMessageBuffer: make(chan messageDetails, messageBufferVolume),
		deadlineTimer:    NewDeadlineTimer(),
		metricsMessageBuffer:    make(chan messageDetails, messageBufferVolume),
		done:             make(chan struct{}),
		executeJournalOvertake:     true,
		wal:              nullJournal{},
		eventpool:           eventpool,
		evsw:             cometsignals.NewEventRouter(),
		stats:          NoopStats(),
	}
	for _, setting := range options {
		setting(cs)
	}
	//
	cs.determineNomination = cs.standardDetermineNomination
	cs.doPreballot = cs.standardDoPreballot
	cs.collectionNomination = cs.standardCollectionNomination

	//
	if status.FinalLedgerLevel > 0 {
		//
		//
		//
		//
		//
		if cs.inactiveStatusAlignLevel != 0 {
			cs.reassembleViewedEndorse(status)
		} else {
			cs.reassembleFinalEndorse(status)
		}
	}

	cs.modifyToStatus(status)

	//

	cs.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", cs)

	return cs
}

//
func (cs *Status) AssignTracer(l log.Tracer) {
	cs.Tracer = l
	cs.deadlineTimer.AssignTracer(l)
}

//
func (cs *Status) AssignEventBus(b *kinds.EventBus) {
	cs.eventBus = b
	cs.ledgerExecute.AssignEventBus(b)
}

//
func StatusStats(stats *Stats) StatusSetting {
	return func(cs *Status) { cs.stats = stats }
}

//
//
func InactiveStatusAlignLevel(level int64) StatusSetting {
	return func(cs *Status) { cs.inactiveStatusAlignLevel = level }
}

//
func (cs *Status) String() string {
	//
	return "REDACTED"
}

//
func (cs *Status) FetchStatus() sm.Status {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return cs.status.Clone()
}

//
//
func (cs *Status) FetchFinalLevel() int64 {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return cs.Level - 1
}

//
//
func (cs *Status) FetchDurationStatus() *cskinds.DurationStatus {
	cs.mtx.RLock()
	rs := cs.fetchDurationStatus()
	cs.mtx.RUnlock()
	return &rs
}

//
//
func (cs *Status) fetchDurationStatus() cskinds.DurationStatus {
	return cs.DurationStatus //
}

//
func (cs *Status) FetchEpochStatusJSON() ([]byte, error) {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return cometjson.Serialize(cs.DurationStatus)
}

//
func (cs *Status) FetchEpochStatusBasicJSON() ([]byte, error) {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return cometjson.Serialize(cs.EpochStatusBasic())
}

//
func (cs *Status) FetchRatifiers() (int64, []*kinds.Ratifier) {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return cs.status.FinalLedgerLevel, cs.status.Ratifiers.Clone().Ratifiers
}

//
//
func (cs *Status) CollectionPrivateRatifier(private kinds.PrivateRatifier) {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	cs.privateRatifier = private

	if err := cs.modifyPrivateRatifierPublicKey(); err != nil {
		cs.Tracer.Fault("REDACTED", "REDACTED", err)
	}
}

//
//
func (cs *Status) CollectionDeadlineTimer(deadlineTimer DeadlineTimer) {
	cs.mtx.Lock()
	cs.deadlineTimer = deadlineTimer
	cs.mtx.Unlock()
}

//
func (cs *Status) ImportEndorse(level int64) *kinds.Endorse {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()

	if level == cs.ledgerDepot.Level() {
		return cs.ledgerDepot.ImportViewedEndorse(level)
	}

	return cs.ledgerDepot.ImportLedgerEndorse(level)
}

//
//
func (cs *Status) OnBegin() error {
	//
	//
	if _, ok := cs.wal.(nullJournal); ok {
		if err := cs.importJournalEntry(); err != nil {
			return err
		}
	}

	//
	//
	//
	//
	//
	if err := cs.deadlineTimer.Begin(); err != nil {
		return err
	}

	//
	//
	if cs.executeJournalOvertake {
		remediateEndeavored := false

	Cycle:
		for {
			err := cs.overtakeResimulate(cs.Level)
			switch {
			case err == nil:
				break Cycle

			case !IsDataImpairmentFault(err):
				cs.Tracer.Fault("REDACTED", "REDACTED", err)
				break Cycle

			case remediateEndeavored:
				return err
			}

			cs.Tracer.Fault("REDACTED", "REDACTED", err)

			//
			if err := cs.wal.Halt(); err != nil {
				return err
			}

			remediateEndeavored = true

			//
			taintedEntry := fmt.Sprintf("REDACTED", cs.settings.JournalEntry())
			if err := cometos.CloneEntry(cs.settings.JournalEntry(), taintedEntry); err != nil {
				return err
			}

			cs.Tracer.Diagnose("REDACTED", "REDACTED", cs.settings.JournalEntry(), "REDACTED", taintedEntry)

			//
			if err := remediateJournalEntry(taintedEntry, cs.settings.JournalEntry()); err != nil {
				cs.Tracer.Fault("REDACTED", "REDACTED", err)
				return err
			}

			cs.Tracer.Details("REDACTED")

			//
			if err := cs.importJournalEntry(); err != nil {
				return err
			}
		}
	}

	if err := cs.evsw.Begin(); err != nil {
		return err
	}

	//
	if err := cs.inspectRepeatAttestingHazard(cs.Level); err != nil {
		return err
	}

	//
	go cs.acceptProcedure(0)

	//
	//
	rs := cs.FetchDurationStatus()
	cs.sequenceEpoch0(rs)

	return nil
}

//
//
func (cs *Status) beginProcedures(maximumPhases int) {
	err := cs.deadlineTimer.Begin()
	if err != nil {
		cs.Tracer.Fault("REDACTED", "REDACTED", err)
		return
	}

	go cs.acceptProcedure(maximumPhases)
}

//
func (cs *Status) importJournalEntry() error {
	wal, err := cs.AccessJournal(cs.settings.JournalEntry())
	if err != nil {
		cs.Tracer.Fault("REDACTED", "REDACTED", err)
		return err
	}

	cs.wal = wal
	return nil
}

//
func (cs *Status) OnHalt() {
	if err := cs.evsw.Halt(); err != nil {
		cs.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	if err := cs.deadlineTimer.Halt(); err != nil {
		cs.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	//
}

//
//
//
func (cs *Status) Wait() {
	<-cs.done
}

//
//
func (cs *Status) AccessJournal(journalEntry string) (WAL, error) {
	wal, err := NewJournal(journalEntry)
	if err != nil {
		cs.Tracer.Fault("REDACTED", "REDACTED", journalEntry, "REDACTED", err)
		return nil, err
	}

	wal.AssignTracer(cs.Tracer.With("REDACTED", journalEntry))

	if err := wal.Begin(); err != nil {
		cs.Tracer.Fault("REDACTED", "REDACTED", err)
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
func (cs *Status) AppendBallot(ballot *kinds.Ballot, nodeUID p2p.ID) (appended bool, err error) {
	if nodeUID == "REDACTED" {
		cs.intrinsicMessageBuffer <- messageDetails{&BallotSignal{ballot}, "REDACTED"}
	} else {
		cs.nodeMessageBuffer <- messageDetails{&BallotSignal{ballot}, nodeUID}
	}

	//
	return false, nil
}

//
func (cs *Status) CollectionNomination(nomination *kinds.Nomination, nodeUID p2p.ID) error {
	if nodeUID == "REDACTED" {
		cs.intrinsicMessageBuffer <- messageDetails{&NominationSignal{nomination}, "REDACTED"}
	} else {
		cs.nodeMessageBuffer <- messageDetails{&NominationSignal{nomination}, nodeUID}
	}

	//
	return nil
}

//
func (cs *Status) AppendNominationLedgerSegment(level int64, duration int32, segment *kinds.Segment, nodeUID p2p.ID) error {
	if nodeUID == "REDACTED" {
		cs.intrinsicMessageBuffer <- messageDetails{&LedgerSegmentSignal{level, duration, segment}, "REDACTED"}
	} else {
		cs.nodeMessageBuffer <- messageDetails{&LedgerSegmentSignal{level, duration, segment}, nodeUID}
	}

	//
	return nil
}

//
func (cs *Status) CollectionNominationAndLedger(
	nomination *kinds.Nomination,
	ledger *kinds.Ledger, //
	segments *kinds.SegmentCollection,
	nodeUID p2p.ID,
) error {
	//
	if err := cs.CollectionNomination(nomination, nodeUID); err != nil {
		return err
	}

	for i := 0; i < int(segments.Sum()); i++ {
		segment := segments.FetchSegment(i)
		if err := cs.AppendNominationLedgerSegment(nomination.Level, nomination.Cycle, segment, nodeUID); err != nil {
			return err
		}
	}

	return nil
}

//
//

func (cs *Status) modifyLevel(level int64) {
	cs.stats.Level.Set(float64(level))
	cs.Level = level
}

func (cs *Status) modifyEpochPhase(duration int32, phase cskinds.DurationPhaseKind) {
	if !cs.resimulateStyle {
		if duration != cs.Cycle || duration == 0 && phase == cskinds.EpochPhaseNewEpoch {
			cs.stats.StampDuration(cs.Cycle, cs.BeginTime)
		}
		if cs.Phase != phase {
			cs.stats.StampPhase(cs.Phase)
		}
	}
	cs.Cycle = duration
	cs.Phase = phase
}

//
func (cs *Status) sequenceEpoch0(rs *cskinds.DurationStatus) {
	//
	pausePeriod := rs.BeginTime.Sub(engineclock.Now())
	cs.sequenceDeadline(pausePeriod, rs.Level, 0, cskinds.DurationPhaseNewLevel)
}

//
func (cs *Status) sequenceDeadline(period time.Duration, level int64, duration int32, phase cskinds.DurationPhaseKind) {
	cs.deadlineTimer.SequenceDeadline(deadlineDetails{period, level, duration, phase})
}

//
func (cs *Status) transmitIntrinsicSignal(mi messageDetails) {
	select {
	case cs.intrinsicMessageBuffer <- mi:
	default:
		//
		//
		//
		//
		cs.Tracer.Diagnose("REDACTED")
		go func() { cs.intrinsicMessageBuffer <- mi }()
	}
}

//
//
//
//
func (cs *Status) reassembleViewedEndorse(status sm.Status) {
	ballots, err := cs.ballotsFromViewedEndorse(status)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	cs.FinalEndorse = ballots
}

//
//
//
//
func (cs *Status) reassembleFinalEndorse(status sm.Status) {
	pluginsActivated := status.AgreementOptions.Iface.BallotPluginsActivated(status.FinalLedgerLevel)
	if !pluginsActivated {
		cs.reassembleViewedEndorse(status)
		return
	}
	ballots, err := cs.ballotsFromExpandedEndorse(status)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	cs.FinalEndorse = ballots
}

func (cs *Status) ballotsFromExpandedEndorse(status sm.Status) (*kinds.BallotCollection, error) {
	ec := cs.ledgerDepot.ImportLedgerExpandedEndorse(status.FinalLedgerLevel)
	if ec == nil {
		return nil, fmt.Errorf("REDACTED", status.FinalLedgerLevel)
	}
	if ec.Level != status.FinalLedgerLevel {
		return nil, fmt.Errorf("REDACTED",
			ec.Level, status.FinalLedgerLevel)
	}
	vs := ec.ToExpandedBallotCollection(status.LedgerUID, status.FinalRatifiers)
	if !vs.HasDualThirdsBulk() {
		return nil, ErrEndorseAssemblyNotFulfilled
	}
	return vs, nil
}

func (cs *Status) ballotsFromViewedEndorse(status sm.Status) (*kinds.BallotCollection, error) {
	endorse := cs.ledgerDepot.ImportViewedEndorse(status.FinalLedgerLevel)
	if endorse == nil {
		endorse = cs.ledgerDepot.ImportLedgerEndorse(status.FinalLedgerLevel)
	}
	if endorse == nil {
		return nil, fmt.Errorf("REDACTED", status.FinalLedgerLevel)
	}
	if endorse.Level != status.FinalLedgerLevel {
		return nil, fmt.Errorf("REDACTED",
			endorse.Level, status.FinalLedgerLevel)
	}
	vs := endorse.ToBallotCollection(status.LedgerUID, status.FinalRatifiers)
	if !vs.HasDualThirdsBulk() {
		return nil, ErrEndorseAssemblyNotFulfilled
	}
	return vs, nil
}

//
//
func (cs *Status) modifyToStatus(status sm.Status) {
	if cs.EndorseEpoch > -1 && 0 < cs.Level && cs.Level != status.FinalLedgerLevel {
		panic(fmt.Sprintf(
			"REDACTED",
			cs.Level, status.FinalLedgerLevel,
		))
	}

	if !cs.status.IsEmpty() {
		if cs.status.FinalLedgerLevel > 0 && cs.status.FinalLedgerLevel+1 != cs.Level {
			//
			//
			panic(fmt.Sprintf(
				"REDACTED",
				cs.status.FinalLedgerLevel+1, cs.Level,
			))
		}
		if cs.status.FinalLedgerLevel > 0 && cs.Level == cs.status.PrimaryLevel {
			panic(fmt.Sprintf(
				"REDACTED",
				cs.status.FinalLedgerLevel, cs.status.PrimaryLevel,
			))
		}

		//
		//
		//
		//
		//
		if status.FinalLedgerLevel <= cs.status.FinalLedgerLevel {
			cs.Tracer.Diagnose(
				"REDACTED",
				"REDACTED", status.FinalLedgerLevel+1,
				"REDACTED", cs.status.FinalLedgerLevel+1,
			)
			cs.newPhase()
			return
		}
	}

	//
	ratifiers := status.Ratifiers

	switch {
	case status.FinalLedgerLevel == 0: //
		cs.FinalEndorse = (*kinds.BallotCollection)(nil)
	case cs.EndorseEpoch > -1 && cs.Ballots != nil: //
		if !cs.Ballots.Preendorsements(cs.EndorseEpoch).HasDualThirdsBulk() {
			panic(fmt.Sprintf(
				"REDACTED",
				status.FinalLedgerLevel, cs.EndorseEpoch, cs.Ballots.Preendorsements(cs.EndorseEpoch),
			))
		}

		cs.FinalEndorse = cs.Ballots.Preendorsements(cs.EndorseEpoch)

	case cs.FinalEndorse == nil:
		//
		//
		panic(fmt.Sprintf(
			"REDACTED",
			status.FinalLedgerLevel+1,
		))
	}

	//
	level := status.FinalLedgerLevel + 1
	if level == 1 {
		level = status.PrimaryLevel
	}

	//
	cs.modifyLevel(level)
	cs.modifyEpochPhase(0, cskinds.DurationPhaseNewLevel)

	if cs.EndorseTime.IsZero() {
		//
		//
		//
		//
		//
		cs.BeginTime = cs.settings.Endorse(engineclock.Now())
	} else {
		cs.BeginTime = cs.settings.Endorse(cs.EndorseTime)
	}

	cs.Ratifiers = ratifiers
	cs.Nomination = nil
	cs.NominationLedger = nil
	cs.NominationLedgerSegments = nil
	cs.LatchedEpoch = -1
	cs.LatchedLedger = nil
	cs.LatchedLedgerSegments = nil
	cs.SoundEpoch = -1
	cs.SoundLedger = nil
	cs.SoundLedgerSegments = nil
	if status.AgreementOptions.Iface.BallotPluginsActivated(level) {
		cs.Ballots = cskinds.NewExpandedLevelBallotCollection(status.LedgerUID, level, ratifiers)
	} else {
		cs.Ballots = cskinds.NewLevelBallotCollection(status.LedgerUID, level, ratifiers)
	}
	cs.EndorseEpoch = -1
	cs.FinalRatifiers = status.FinalRatifiers
	cs.ActivatedDeadlinePreendorse = false

	cs.status = status

	//
	cs.newPhase()
	cs.newAgreementOptions()
}

func (cs *Status) newPhase() {
	rs := cs.EpochStatusEvent()
	if err := cs.wal.Record(rs); err != nil {
		cs.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	cs.nPhases++

	//
	if cs.eventBus != nil {
		if err := cs.eventBus.BroadcastEventNewEpochPhase(rs); err != nil {
			cs.Tracer.Fault("REDACTED", "REDACTED", err)
		}

		cs.evsw.TriggerEvent(kinds.EventNewDurationPhase, cs.DurationStatus)
	}
}

//
//
func (cs *Status) newAgreementOptions() {
	cs.evsw.TriggerEvent(kinds.EventNewAgreementOptions, cs.status.AgreementOptions)
}

//
//

//
//
//
//
//
func (cs *Status) acceptProcedure(maximumPhases int) {
	onQuit := func(cs *Status) {
		//
		//
		//

		//
		if err := cs.wal.Halt(); err != nil {
			cs.Tracer.Fault("REDACTED", "REDACTED", err)
		}

		cs.wal.Wait()
		close(cs.done)
	}

	defer func() {
		if r := recover(); r != nil {
			cs.Tracer.Fault("REDACTED", "REDACTED", r, "REDACTED", string(debug.Stack()))
			//
			//
			//
			//
			//
			//
			//
			//
			onQuit(cs)
		}
	}()

	for {
		if maximumPhases > 0 {
			if cs.nPhases >= maximumPhases {
				cs.Tracer.Diagnose("REDACTED")
				cs.nPhases = 0
				return
			}
		}

		rs := cs.DurationStatus
		var mi messageDetails

		select {
		case <-cs.transferAlerter.TransAccessible():
			cs.processTransAccessible()

		case mi = <-cs.nodeMessageBuffer:
			cs.Tracer.Diagnose("REDACTED", "REDACTED", mi.NodeUID)

			if err := cs.wal.Record(mi); err != nil {
				cs.Tracer.Fault("REDACTED", "REDACTED", err)
			}

			//
			//
			cs.processMessage(mi)

		case mi = <-cs.intrinsicMessageBuffer:
			cs.Tracer.Diagnose("REDACTED", "REDACTED", mi.NodeUID)

			err := cs.wal.RecordAlign(mi) //
			if err != nil {
				panic(fmt.Sprintf(
					"REDACTED",
					mi, err,
				))
			}

			if _, ok := mi.Msg.(*BallotSignal); ok {
				//
				//
				//
				//
				abort.Abort() //
			}

			//
			cs.processMessage(mi)

		case ti := <-cs.deadlineTimer.Chan(): //
			if err := cs.wal.Record(ti); err != nil {
				cs.Tracer.Fault("REDACTED", "REDACTED", err)
			}

			//
			//
			cs.processDeadline(ti, rs)

		case <-cs.Exit():
			onQuit(cs)
			return
		}
	}
}

//
func (cs *Status) processMessage(mi messageDetails) {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	var (
		appended bool
		err   error
	)

	msg, nodeUID := mi.Msg, mi.NodeUID

	cs.Tracer.Diagnose("REDACTED", "REDACTED", string(nodeUID), "REDACTED", msg)

	switch msg := msg.(type) {
	case *NominationSignal:
		//
		//
		err = cs.collectionNomination(msg.Nomination)

	case *LedgerSegmentSignal:
		//
		appended, err = cs.appendNominationLedgerSegment(msg, nodeUID)

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
		if appended && cs.NominationLedgerSegments.IsFinished() {
			cs.processFinishedNomination(msg.Level)
		}
		if appended {
			cs.metricsMessageBuffer <- mi
		}

		if err != nil && msg.Cycle != cs.Cycle {
			cs.Tracer.Diagnose(
				"REDACTED",
				"REDACTED", cs.Level,
				"REDACTED", cs.Cycle,
				"REDACTED", msg.Cycle,
			)
			err = nil
		}

	case *BallotSignal:
		//
		//
		appended, err = cs.attemptAppendBallot(msg.Ballot, nodeUID)
		if appended {
			cs.metricsMessageBuffer <- mi
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

	default:
		cs.Tracer.Fault("REDACTED", "REDACTED", fmt.Sprintf("REDACTED", msg))
		return
	}

	if err != nil {
		cs.Tracer.Fault(
			"REDACTED",
			"REDACTED", cs.Level,
			"REDACTED", cs.Cycle,
			"REDACTED", nodeUID,
			"REDACTED", fmt.Sprintf("REDACTED", msg),
			"REDACTED", err,
		)
	}
}

func (cs *Status) processDeadline(ti deadlineDetails, rs cskinds.DurationStatus) {
	cs.Tracer.Diagnose("REDACTED", "REDACTED", ti.Period, "REDACTED", ti.Level, "REDACTED", ti.Cycle, "REDACTED", ti.Phase)

	//
	if ti.Level != rs.Level || ti.Cycle < rs.Cycle || (ti.Cycle == rs.Cycle && ti.Phase < rs.Phase) {
		cs.Tracer.Diagnose("REDACTED", "REDACTED", rs.Level, "REDACTED", rs.Cycle, "REDACTED", rs.Phase)
		return
	}

	//
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	switch ti.Phase {
	case cskinds.DurationPhaseNewLevel:
		//
		//
		cs.joinNewEpoch(ti.Level, 0)

	case cskinds.EpochPhaseNewEpoch:
		cs.joinNominate(ti.Level, ti.Cycle)

	case cskinds.DurationPhaseNominate:
		if err := cs.eventBus.BroadcastEventDeadlineNominate(cs.EpochStatusEvent()); err != nil {
			cs.Tracer.Fault("REDACTED", "REDACTED", err)
		}

		cs.joinPreballot(ti.Level, ti.Cycle)

	case cskinds.DurationPhasePreballotWait:
		if err := cs.eventBus.BroadcastEventDeadlineWait(cs.EpochStatusEvent()); err != nil {
			cs.Tracer.Fault("REDACTED", "REDACTED", err)
		}

		cs.joinPreendorse(ti.Level, ti.Cycle)

	case cskinds.DurationPhasePreendorseWait:
		if err := cs.eventBus.BroadcastEventDeadlineWait(cs.EpochStatusEvent()); err != nil {
			cs.Tracer.Fault("REDACTED", "REDACTED", err)
		}

		cs.issuePreendorseDeadlineStats(ti.Cycle)
		cs.joinPreendorse(ti.Level, ti.Cycle)
		cs.joinNewEpoch(ti.Level, ti.Cycle+1)

	default:
		panic(cometfaults.ErrCorruptField{Field: "REDACTED"})
	}
}

func (cs *Status) processTransAccessible() {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	//
	if cs.Cycle != 0 {
		return
	}

	switch cs.Phase {
	case cskinds.DurationPhaseNewLevel: //
		if cs.requireAttestationLedger(cs.Level) {
			//
			return
		}

		//
		deadlineEndorse := cs.BeginTime.Sub(engineclock.Now()) + 1*time.Millisecond
		cs.sequenceDeadline(deadlineEndorse, cs.Level, 0, cskinds.EpochPhaseNewEpoch)

	case cskinds.EpochPhaseNewEpoch: //
		cs.joinNominate(cs.Level, 0)
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
func (cs *Status) joinNewEpoch(level int64, duration int32) {
	tracer := cs.Tracer.With("REDACTED", level, "REDACTED", duration)

	if cs.Level != level || duration < cs.Cycle || (cs.Cycle == duration && cs.Phase != cskinds.DurationPhaseNewLevel) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase),
		)
		return
	}

	if now := engineclock.Now(); cs.BeginTime.After(now) {
		tracer.Diagnose("REDACTED", "REDACTED", cs.BeginTime, "REDACTED", now)
	}

	priorLevel, priorEpoch, priorPhase := cs.Level, cs.Cycle, cs.Phase

	//
	ratifiers := cs.Ratifiers
	if cs.Cycle < duration {
		ratifiers = ratifiers.Clone()
		ratifiers.AugmentRecommenderUrgency(cometmath.SecureSubtractInt32(duration, cs.Cycle))
	}

	//
	//
	//
	cs.modifyEpochPhase(duration, cskinds.EpochPhaseNewEpoch)
	cs.Ratifiers = ratifiers
	//
	//
	nominationLocation := ratifiers.FetchRecommender().PublicKey.Location()
	if duration != 0 {
		tracer.Details("REDACTED", "REDACTED", nominationLocation)
		cs.Nomination = nil
		cs.NominationLedger = nil
		cs.NominationLedgerSegments = nil
	}

	tracer.Diagnose("REDACTED",
		"REDACTED", log.NewIdleFormat("REDACTED", priorLevel, priorEpoch, priorPhase),
		"REDACTED", nominationLocation,
	)

	if duration > 0 && !cs.resimulateStyle {
		cs.stats.StampDurationAugmented(priorPhase)
	}

	cs.Ballots.CollectionEpoch(cometmath.SecureAppendInt32(duration, 1)) //
	cs.ActivatedDeadlinePreendorse = false

	if err := cs.eventBus.BroadcastEventNewEpoch(cs.NewEpochEvent()); err != nil {
		cs.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	//
	//
	//
	waitForTrans := cs.settings.WaitForTrans() && duration == 0 && !cs.requireAttestationLedger(level)
	if waitForTrans {
		if cs.settings.GenerateEmptyLedgersCadence > 0 {
			cs.sequenceDeadline(cs.settings.GenerateEmptyLedgersCadence, level, duration,
				cskinds.EpochPhaseNewEpoch)
		}
	} else {
		cs.joinNominate(level, duration)
	}
}

//
//
func (cs *Status) requireAttestationLedger(level int64) bool {
	if level == cs.status.PrimaryLevel {
		return true
	}

	finalLedgerMeta := cs.ledgerDepot.ImportLedgerMeta(level - 1)
	if finalLedgerMeta == nil {
		//
		cs.Tracer.Details("REDACTED", "REDACTED", level, "REDACTED", cs.status.PrimaryLevel)
		return true
	}

	return !bytes.Equal(cs.status.ApplicationDigest, finalLedgerMeta.Heading.ApplicationDigest)
}

//
//
//
//
//
//
func (cs *Status) joinNominate(level int64, duration int32) {
	tracer := cs.Tracer.With("REDACTED", level, "REDACTED", duration)

	if cs.Level != level || duration < cs.Cycle || (cs.Cycle == duration && cskinds.DurationPhaseNominate <= cs.Phase) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase),
		)
		return
	}

	tracer.Diagnose("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase))

	defer func() {
		//
		cs.modifyEpochPhase(duration, cskinds.DurationPhaseNominate)
		cs.newPhase()

		//
		//
		//
		if cs.isNominationFinished() {
			cs.joinPreballot(level, cs.Cycle)
		}
	}()

	//
	cs.sequenceDeadline(cs.settings.Nominate(duration), level, duration, cskinds.DurationPhaseNominate)

	//
	if cs.privateRatifier == nil {
		tracer.Diagnose("REDACTED")
		return
	}

	tracer.Diagnose("REDACTED")

	if cs.privateRatifierPublicKey == nil {
		//
		//
		tracer.Fault("REDACTED", "REDACTED", ErrPublicKeyIsNotCollection)
		return
	}

	location := cs.privateRatifierPublicKey.Location()

	//
	if !cs.Ratifiers.HasLocation(location) {
		tracer.Diagnose("REDACTED", "REDACTED", location, "REDACTED", cs.Ratifiers)
		return
	}

	if cs.isRecommender(location) {
		tracer.Diagnose("REDACTED", "REDACTED", location)
		cs.determineNomination(level, duration)
	} else {
		tracer.Diagnose("REDACTED", "REDACTED", cs.Ratifiers.FetchRecommender().Location)
	}
}

func (cs *Status) isRecommender(location []byte) bool {
	return bytes.Equal(cs.Ratifiers.FetchRecommender().Location, location)
}

func (cs *Status) standardDetermineNomination(level int64, duration int32) {
	var ledger *kinds.Ledger
	var ledgerSegments *kinds.SegmentCollection

	//
	if cs.SoundLedger != nil {
		//
		ledger, ledgerSegments = cs.SoundLedger, cs.SoundLedgerSegments
	} else {
		//
		var err error
		ledger, err = cs.instantiateNominationLedger(context.TODO())
		if err != nil {
			cs.Tracer.Fault("REDACTED", "REDACTED", err)
			return
		} else if ledger == nil {
			panic("REDACTED")
		}
		cs.stats.NominationInstantiateTally.Add(1)
		ledgerSegments, err = ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
		if err != nil {
			cs.Tracer.Fault("REDACTED", "REDACTED", err)
			return
		}
	}

	//
	//
	if err := cs.wal.PurgeAndAlign(); err != nil {
		cs.Tracer.Fault("REDACTED")
	}

	//
	nominationLedgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: ledgerSegments.Heading()}
	nomination := kinds.NewNomination(level, duration, cs.SoundEpoch, nominationLedgerUID)
	p := nomination.ToSchema()
	if err := cs.privateRatifier.AttestNomination(cs.status.LedgerUID, p); err == nil {
		nomination.Autograph = p.Autograph

		//
		cs.transmitIntrinsicSignal(messageDetails{&NominationSignal{nomination}, "REDACTED"})

		for i := 0; i < int(ledgerSegments.Sum()); i++ {
			segment := ledgerSegments.FetchSegment(i)
			cs.transmitIntrinsicSignal(messageDetails{&LedgerSegmentSignal{cs.Level, cs.Cycle, segment}, "REDACTED"})
		}

		cs.Tracer.Diagnose("REDACTED", "REDACTED", level, "REDACTED", duration, "REDACTED", nomination)
	} else if !cs.resimulateStyle {
		cs.Tracer.Fault("REDACTED", "REDACTED", level, "REDACTED", duration, "REDACTED", err)
	}
}

//
//
func (cs *Status) isNominationFinished() bool {
	if cs.Nomination == nil || cs.NominationLedger == nil {
		return false
	}
	//
	//
	if cs.Nomination.POLDuration < 0 {
		return true
	}
	//
	return cs.Ballots.Preballots(cs.Nomination.POLDuration).HasDualThirdsBulk()
}

//
//
//
//
//
//
//
func (cs *Status) instantiateNominationLedger(ctx context.Context) (*kinds.Ledger, error) {
	if cs.privateRatifier == nil {
		return nil, ErrNullPrivateRatifier
	}

	//
	var finalExtensionEndorse *kinds.ExpandedEndorse
	switch {
	case cs.Level == cs.status.PrimaryLevel:
		//
		//
		finalExtensionEndorse = &kinds.ExpandedEndorse{}

	case cs.FinalEndorse.HasDualThirdsBulk():
		//
		finalExtensionEndorse = cs.FinalEndorse.CreateExpandedEndorse(cs.status.AgreementOptions.Iface)

	default: //
		return nil, ErrNominationLackingPrecedingEndorse
	}

	if cs.privateRatifierPublicKey == nil {
		//
		//
		return nil, fmt.Errorf("REDACTED", ErrPublicKeyIsNotCollection)
	}

	recommenderAddress := cs.privateRatifierPublicKey.Location()

	ret, err := cs.ledgerExecute.InstantiateNominationLedger(ctx, cs.Level, cs.status, finalExtensionEndorse, recommenderAddress)
	if err != nil {
		panic(err)
	}
	return ret, nil
}

//
//
//
//
func (cs *Status) joinPreballot(level int64, duration int32) {
	tracer := cs.Tracer.With("REDACTED", level, "REDACTED", duration)

	if cs.Level != level || duration < cs.Cycle || (cs.Cycle == duration && cskinds.EpochPhasePreballot <= cs.Phase) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase),
		)
		return
	}

	defer func() {
		//
		cs.modifyEpochPhase(duration, cskinds.EpochPhasePreballot)
		cs.newPhase()
	}()

	tracer.Diagnose("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase))

	//
	cs.doPreballot(level, duration)

	//
	//
}

func (cs *Status) standardDoPreballot(level int64, duration int32) {
	tracer := cs.Tracer.With("REDACTED", level, "REDACTED", duration)

	//
	if cs.LatchedLedger != nil {
		tracer.Diagnose("REDACTED")
		cs.attestAppendBallot(engineproto.PreballotKind, cs.LatchedLedger.Digest(), cs.LatchedLedgerSegments.Heading(), nil)
		return
	}

	//
	if cs.NominationLedger == nil {
		tracer.Diagnose("REDACTED")
		cs.attestAppendBallot(engineproto.PreballotKind, nil, kinds.SegmentAssignHeading{}, nil)
		return
	}

	//
	err := cs.ledgerExecute.CertifyLedger(cs.status, cs.NominationLedger)
	if err != nil {
		//
		tracer.Fault("REDACTED",
			"REDACTED", err)
		cs.attestAppendBallot(engineproto.PreballotKind, nil, kinds.SegmentAssignHeading{}, nil)
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
	isApplicationSound, err := cs.ledgerExecute.HandleNomination(cs.NominationLedger, cs.status)
	if err != nil {
		panic(fmt.Sprintf(
			"REDACTED", err,
		))
	}
	cs.stats.StampNominationHandled(isApplicationSound)

	//
	if !isApplicationSound {
		tracer.Fault("REDACTED"+
			"REDACTED", "REDACTED", err)
		cs.attestAppendBallot(engineproto.PreballotKind, nil, kinds.SegmentAssignHeading{}, nil)
		return
	}

	//
	//
	//
	tracer.Diagnose("REDACTED")
	cs.attestAppendBallot(engineproto.PreballotKind, cs.NominationLedger.Digest(), cs.NominationLedgerSegments.Heading(), nil)
}

//
func (cs *Status) joinPreballotWait(level int64, duration int32) {
	tracer := cs.Tracer.With("REDACTED", level, "REDACTED", duration)

	if cs.Level != level || duration < cs.Cycle || (cs.Cycle == duration && cskinds.DurationPhasePreballotWait <= cs.Phase) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase),
		)
		return
	}

	if !cs.Ballots.Preballots(duration).HasDualThirdsAny() {
		panic(fmt.Sprintf(
			"REDACTED",
			level, duration,
		))
	}

	tracer.Diagnose("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase))

	defer func() {
		//
		cs.modifyEpochPhase(duration, cskinds.DurationPhasePreballotWait)
		cs.newPhase()
	}()

	//
	cs.sequenceDeadline(cs.settings.Preballot(duration), level, duration, cskinds.DurationPhasePreballotWait)
}

//
//
//
//
//
//
func (cs *Status) joinPreendorse(level int64, duration int32) {
	tracer := cs.Tracer.With("REDACTED", level, "REDACTED", duration)

	if cs.Level != level || duration < cs.Cycle || (cs.Cycle == duration && cskinds.EpochPhasePreendorse <= cs.Phase) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase),
		)
		return
	}

	tracer.Diagnose("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase))

	defer func() {
		//
		cs.modifyEpochPhase(duration, cskinds.EpochPhasePreendorse)
		cs.newPhase()
	}()

	//
	ledgerUID, ok := cs.Ballots.Preballots(duration).DualThirdsBulk()

	//
	if !ok {
		if cs.LatchedLedger != nil {
			tracer.Diagnose("REDACTED")
		} else {
			tracer.Diagnose("REDACTED")
		}

		cs.attestAppendBallot(engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, nil)
		return
	}

	//
	if err := cs.eventBus.BroadcastEventPolka(cs.EpochStatusEvent()); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
	}

	//
	polEpoch, _ := cs.Ballots.POLDetails()
	if polEpoch < duration {
		panic(fmt.Sprintf("REDACTED", duration, polEpoch))
	}

	//
	if len(ledgerUID.Digest) == 0 {
		if cs.LatchedLedger == nil {
			tracer.Diagnose("REDACTED")
		} else {
			tracer.Diagnose("REDACTED")
			cs.LatchedEpoch = -1
			cs.LatchedLedger = nil
			cs.LatchedLedgerSegments = nil

			if err := cs.eventBus.BroadcastEventRelease(cs.EpochStatusEvent()); err != nil {
				tracer.Fault("REDACTED", "REDACTED", err)
			}
		}

		cs.attestAppendBallot(engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, nil)
		return
	}

	//

	//
	if cs.LatchedLedger.DigestsTo(ledgerUID.Digest) {
		tracer.Diagnose("REDACTED")
		cs.LatchedEpoch = duration

		if err := cs.eventBus.BroadcastEventResecure(cs.EpochStatusEvent()); err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
		}

		cs.attestAppendBallot(engineproto.PreendorseKind, ledgerUID.Digest, ledgerUID.SegmentAssignHeading, cs.LatchedLedger)
		return
	}

	//
	if cs.NominationLedger.DigestsTo(ledgerUID.Digest) {
		tracer.Diagnose("REDACTED", "REDACTED", ledgerUID.Digest)

		//
		if err := cs.ledgerExecute.CertifyLedger(cs.status, cs.NominationLedger); err != nil {
			panic(fmt.Sprintf("REDACTED", err))
		}

		cs.LatchedEpoch = duration
		cs.LatchedLedger = cs.NominationLedger
		cs.LatchedLedgerSegments = cs.NominationLedgerSegments

		if err := cs.eventBus.BroadcastEventSecure(cs.EpochStatusEvent()); err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
		}

		cs.attestAppendBallot(engineproto.PreendorseKind, ledgerUID.Digest, ledgerUID.SegmentAssignHeading, cs.NominationLedger)
		return
	}

	//
	//
	//
	tracer.Diagnose("REDACTED", "REDACTED", ledgerUID)

	cs.LatchedEpoch = -1
	cs.LatchedLedger = nil
	cs.LatchedLedgerSegments = nil

	if !cs.NominationLedgerSegments.HasHeading(ledgerUID.SegmentAssignHeading) {
		cs.NominationLedger = nil
		cs.NominationLedgerSegments = kinds.NewSegmentCollectionFromHeading(ledgerUID.SegmentAssignHeading)
	}

	if err := cs.eventBus.BroadcastEventRelease(cs.EpochStatusEvent()); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
	}

	cs.attestAppendBallot(engineproto.PreendorseKind, nil, kinds.SegmentAssignHeading{}, nil)
}

//
func (cs *Status) joinPreendorseWait(level int64, duration int32) {
	tracer := cs.Tracer.With("REDACTED", level, "REDACTED", duration)

	if cs.Level != level || duration < cs.Cycle || (cs.Cycle == duration && cs.ActivatedDeadlinePreendorse) {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", cs.ActivatedDeadlinePreendorse,
			"REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle),
		)
		return
	}

	if !cs.Ballots.Preendorsements(duration).HasDualThirdsAny() {
		panic(fmt.Sprintf(
			"REDACTED",
			level, duration,
		))
	}

	tracer.Diagnose("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase))

	defer func() {
		//
		cs.ActivatedDeadlinePreendorse = true
		cs.newPhase()
	}()

	//
	cs.sequenceDeadline(cs.settings.Preendorse(duration), level, duration, cskinds.DurationPhasePreendorseWait)
}

//
func (cs *Status) joinEndorse(level int64, endorseEpoch int32) {
	tracer := cs.Tracer.With("REDACTED", level, "REDACTED", endorseEpoch)

	if cs.Level != level || cskinds.DurationPhaseEndorse <= cs.Phase {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase),
		)
		return
	}

	tracer.Diagnose("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase))

	defer func() {
		//
		//
		cs.modifyEpochPhase(cs.Cycle, cskinds.DurationPhaseEndorse)
		cs.EndorseEpoch = endorseEpoch
		cs.EndorseTime = engineclock.Now()
		cs.newPhase()

		//
		cs.attemptCompleteEndorse(level)
	}()

	ledgerUID, ok := cs.Ballots.Preendorsements(endorseEpoch).DualThirdsBulk()
	if !ok {
		panic("REDACTED")
	}

	//
	//
	//
	if cs.LatchedLedger.DigestsTo(ledgerUID.Digest) {
		tracer.Diagnose("REDACTED", "REDACTED", ledgerUID.Digest)
		cs.NominationLedger = cs.LatchedLedger
		cs.NominationLedgerSegments = cs.LatchedLedgerSegments
	}

	//
	if !cs.NominationLedger.DigestsTo(ledgerUID.Digest) {
		if !cs.NominationLedgerSegments.HasHeading(ledgerUID.SegmentAssignHeading) {
			tracer.Details(
				"REDACTED",
				"REDACTED", log.NewIdleLedgerDigest(cs.NominationLedger),
				"REDACTED", ledgerUID.Digest,
			)

			//
			//
			cs.NominationLedger = nil
			cs.NominationLedgerSegments = kinds.NewSegmentCollectionFromHeading(ledgerUID.SegmentAssignHeading)

			if err := cs.eventBus.BroadcastEventSoundLedger(cs.EpochStatusEvent()); err != nil {
				tracer.Fault("REDACTED", "REDACTED", err)
			}

			cs.evsw.TriggerEvent(kinds.EventSoundLedger, cs.DurationStatus)
		}
	}
}

//
func (cs *Status) attemptCompleteEndorse(level int64) {
	tracer := cs.Tracer.With("REDACTED", level)

	if cs.Level != level {
		panic(fmt.Sprintf("REDACTED", cs.Level, level))
	}

	ledgerUID, ok := cs.Ballots.Preendorsements(cs.EndorseEpoch).DualThirdsBulk()
	if !ok || len(ledgerUID.Digest) == 0 {
		tracer.Fault("REDACTED")
		return
	}

	if !cs.NominationLedger.DigestsTo(ledgerUID.Digest) {
		//
		//
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.NewIdleLedgerDigest(cs.NominationLedger),
			"REDACTED", ledgerUID.Digest,
		)
		return
	}

	cs.completeEndorse(level)
}

//
func (cs *Status) completeEndorse(level int64) {
	tracer := cs.Tracer.With("REDACTED", level)

	if cs.Level != level || cs.Phase != cskinds.DurationPhaseEndorse {
		tracer.Diagnose(
			"REDACTED",
			"REDACTED", log.NewIdleFormat("REDACTED", cs.Level, cs.Cycle, cs.Phase),
		)
		return
	}

	cs.computePreballotSignalDeferralStats()

	ledgerUID, ok := cs.Ballots.Preendorsements(cs.EndorseEpoch).DualThirdsBulk()
	ledger, ledgerSegments := cs.NominationLedger, cs.NominationLedgerSegments

	if !ok {
		panic("REDACTED")
	}
	if !ledgerSegments.HasHeading(ledgerUID.SegmentAssignHeading) {
		panic("REDACTED")
	}
	if !ledger.DigestsTo(ledgerUID.Digest) {
		panic("REDACTED")
	}

	if err := cs.ledgerExecute.CertifyLedger(cs.status, ledger); err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	cs.computePreendorseSignalDeferralStats()

	tracer.Details(
		"REDACTED",
		"REDACTED", log.NewIdleLedgerDigest(ledger),
		"REDACTED", ledger.ApplicationDigest,
		"REDACTED", len(ledger.Txs),
	)
	tracer.Diagnose("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", ledger))

	abort.Abort() //

	//
	if cs.ledgerDepot.Level() < ledger.Level {
		//
		//
		viewedExpandedEndorse := cs.Ballots.Preendorsements(cs.EndorseEpoch).CreateExpandedEndorse(cs.status.AgreementOptions.Iface)
		if cs.status.AgreementOptions.Iface.BallotPluginsActivated(ledger.Level) {
			cs.ledgerDepot.PersistLedgerWithExpandedEndorse(ledger, ledgerSegments, viewedExpandedEndorse)
		} else {
			cs.ledgerDepot.PersistLedger(ledger, ledgerSegments, viewedExpandedEndorse.ToEndorse())
		}
	} else {
		//
		tracer.Diagnose("REDACTED", "REDACTED", ledger.Level)
	}

	abort.Abort() //

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
	terminateMessage := TerminateLevelSignal{level}
	if err := cs.wal.RecordAlign(terminateMessage); err != nil { //
		panic(fmt.Sprintf(
			"REDACTED",
			terminateMessage, err,
		))
	}

	abort.Abort() //

	//
	statusClone := cs.status.Clone()

	//
	//
	//
	statusClone, err := cs.ledgerExecute.ExecuteValidatedLedger(
		statusClone,
		kinds.LedgerUID{
			Digest:          ledger.Digest(),
			SegmentAssignHeading: ledgerSegments.Heading(),
		},
		ledger,
	)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}

	abort.Abort() //

	//
	cs.logStats(level, ledger)

	//
	cs.modifyToStatus(statusClone)

	abort.Abort() //

	//
	if err := cs.modifyPrivateRatifierPublicKey(); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
	}

	//
	//
	cs.sequenceEpoch0(&cs.DurationStatus)

	//
	//
	//
	//
}

func (cs *Status) logStats(level int64, ledger *kinds.Ledger) {
	cs.stats.Ratifiers.Set(float64(cs.Ratifiers.Volume()))
	cs.stats.RatifiersEnergy.Set(float64(cs.Ratifiers.SumPollingEnergy()))

	var (
		absentRatifiers      int
		absentRatifiersEnergy int64
	)
	//
	//
	//
	if level > cs.status.PrimaryLevel {
		//
		//
		var (
			endorseVolume = ledger.FinalEndorse.Volume()
			valueCollectionSize  = len(cs.FinalRatifiers.Ratifiers)
			location    kinds.Location
		)
		if endorseVolume != valueCollectionSize {
			panic(fmt.Sprintf("REDACTED",
				endorseVolume, valueCollectionSize, ledger.Level, ledger.FinalEndorse.Endorsements, cs.FinalRatifiers.Ratifiers))
		}

		if cs.privateRatifier != nil {
			if cs.privateRatifierPublicKey == nil {
				//
				cs.Tracer.Fault(fmt.Sprintf("REDACTED", ErrPublicKeyIsNotCollection))
			} else {
				location = cs.privateRatifierPublicKey.Location()
			}
		}

		for i, val := range cs.FinalRatifiers.Ratifiers {
			endorseSignature := ledger.FinalEndorse.Endorsements[i]
			if endorseSignature.LedgerUIDMark == kinds.LedgerUIDMarkMissing {
				absentRatifiers++
				absentRatifiersEnergy += val.PollingEnergy
			}

			if bytes.Equal(val.Location, location) {
				tag := []string{
					"REDACTED", val.Location.String(),
				}
				cs.stats.RatifierEnergy.With(tag...).Set(float64(val.PollingEnergy))
				if endorseSignature.LedgerUIDMark == kinds.LedgerUIDMarkEndorse {
					cs.stats.RatifierFinalAttestedLevel.With(tag...).Set(float64(level))
				} else {
					cs.stats.RatifierSkippedLedgers.With(tag...).Add(float64(1))
				}
			}

		}
	}
	cs.stats.AbsentRatifiers.Set(float64(absentRatifiers))
	cs.stats.AbsentRatifiersEnergy.Set(float64(absentRatifiersEnergy))

	//
	var (
		faultyRatifiersEnergy = int64(0)
		faultyRatifiersTally = int64(0)
	)
	for _, ev := range ledger.Proof.Proof {
		if dve, ok := ev.(*kinds.ReplicatedBallotProof); ok {
			if _, val := cs.Ratifiers.FetchByLocationMut(dve.BallotA.RatifierLocation); val != nil {
				faultyRatifiersTally++
				faultyRatifiersEnergy += val.PollingEnergy
			}
		}
	}
	cs.stats.FaultyRatifiers.Set(float64(faultyRatifiersTally))
	cs.stats.FaultyRatifiersEnergy.Set(float64(faultyRatifiersEnergy))

	if level > 1 {
		finalLedgerMeta := cs.ledgerDepot.ImportLedgerMeta(level - 1)
		if finalLedgerMeta != nil {
			cs.stats.LedgerCadenceMoments.Observe(
				ledger.Time.Sub(finalLedgerMeta.Heading.Time).Seconds(),
			)
		}
	}

	cs.stats.CountTrans.Set(float64(len(ledger.Txs)))
	cs.stats.SumTrans.Add(float64(len(ledger.Txs)))
	cs.stats.LedgerVolumeOctets.Set(float64(ledger.Volume()))
	cs.stats.SeriesVolumeOctets.Add(float64(ledger.Volume()))
	cs.stats.ConfirmedLevel.Set(float64(ledger.Level))
}

//

func (cs *Status) standardCollectionNomination(nomination *kinds.Nomination) error {
	//
	//
	if cs.Nomination != nil {
		return nil
	}

	//
	if nomination.Level != cs.Level || nomination.Cycle != cs.Cycle {
		return nil
	}

	//
	if nomination.POLDuration < -1 ||
		(nomination.POLDuration >= 0 && nomination.POLDuration >= nomination.Cycle) {
		return ErrCorruptNominationPOLEpoch
	}

	p := nomination.ToSchema()
	//
	publicKey := cs.Ratifiers.FetchRecommender().PublicKey
	if !publicKey.ValidateAutograph(
		kinds.NominationAttestOctets(cs.status.LedgerUID, p), nomination.Autograph,
	) {
		return ErrCorruptNominationAutograph
	}

	//
	maximumOctets := cs.status.AgreementOptions.Ledger.MaximumOctets
	if maximumOctets == -1 {
		maximumOctets = int64(kinds.MaximumLedgerVolumeOctets)
	}
	if int64(nomination.LedgerUID.SegmentAssignHeading.Sum) > (maximumOctets-1)/int64(kinds.LedgerSegmentVolumeOctets)+1 {
		return ErrNominationTooNumerousSegments
	}

	nomination.Autograph = p.Autograph
	cs.Nomination = nomination
	//
	//
	//
	if cs.NominationLedgerSegments == nil {
		cs.NominationLedgerSegments = kinds.NewSegmentCollectionFromHeading(nomination.LedgerUID.SegmentAssignHeading)
	}

	cs.Tracer.Details("REDACTED", "REDACTED", nomination, "REDACTED", publicKey.Location())
	return nil
}

//
//
//
func (cs *Status) appendNominationLedgerSegment(msg *LedgerSegmentSignal, nodeUID p2p.ID) (appended bool, err error) {
	level, duration, segment := msg.Level, msg.Cycle, msg.Segment

	//
	if cs.Level != level {
		cs.Tracer.Diagnose("REDACTED", "REDACTED", level, "REDACTED", duration)
		cs.stats.LedgerGossipSegmentsAccepted.With("REDACTED", "REDACTED").Add(1)
		return false, nil
	}

	//
	if cs.NominationLedgerSegments == nil {
		cs.stats.LedgerGossipSegmentsAccepted.With("REDACTED", "REDACTED").Add(1)
		//
		//
		cs.Tracer.Diagnose(
			"REDACTED",
			"REDACTED", level,
			"REDACTED", duration,
			"REDACTED", segment.Ordinal,
			"REDACTED", nodeUID,
		)
		return false, nil
	}

	appended, err = cs.NominationLedgerSegments.AppendSegment(segment)
	if err != nil {
		if errors.Is(err, kinds.ErrSegmentCollectionCorruptAttestation) || errors.Is(err, kinds.ErrSegmentCollectionUnforeseenOrdinal) {
			cs.stats.LedgerGossipSegmentsAccepted.With("REDACTED", "REDACTED").Add(1)
		}
		return appended, err
	}

	cs.stats.LedgerGossipSegmentsAccepted.With("REDACTED", "REDACTED").Add(1)
	if !appended {
		//
		//
		cs.stats.ReplicatedLedgerSegment.Add(1)
	}

	maximumOctets := cs.status.AgreementOptions.Ledger.MaximumOctets
	if maximumOctets == -1 {
		maximumOctets = int64(kinds.MaximumLedgerVolumeOctets)
	}
	if cs.NominationLedgerSegments.OctetVolume() > maximumOctets {
		return appended, fmt.Errorf("REDACTED",
			cs.NominationLedgerSegments.OctetVolume(), maximumOctets,
		)
	}
	if appended && cs.NominationLedgerSegments.IsFinished() {
		bz, err := io.ReadAll(cs.NominationLedgerSegments.FetchScanner())
		if err != nil {
			return appended, err
		}

		pbb := new(engineproto.Ledger)
		err = proto.Unmarshal(bz, pbb)
		if err != nil {
			return appended, err
		}

		ledger, err := kinds.LedgerFromSchema(pbb)
		if err != nil {
			return appended, err
		}

		cs.NominationLedger = ledger

		//
		cs.Tracer.Details("REDACTED", "REDACTED", cs.NominationLedger.Level, "REDACTED", cs.NominationLedger.Digest())

		if err := cs.eventBus.BroadcastEventFinishedNomination(cs.FinishedNominationEvent()); err != nil {
			cs.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	return appended, nil
}

func (cs *Status) processFinishedNomination(ledgerLevel int64) {
	//
	preballots := cs.Ballots.Preballots(cs.Cycle)
	ledgerUID, hasDualThirds := preballots.DualThirdsBulk()
	if hasDualThirds && !ledgerUID.IsNil() && (cs.SoundEpoch < cs.Cycle) {
		if cs.NominationLedger.DigestsTo(ledgerUID.Digest) {
			cs.Tracer.Diagnose(
				"REDACTED",
				"REDACTED", cs.Cycle,
				"REDACTED", log.NewIdleLedgerDigest(cs.NominationLedger),
			)

			cs.SoundEpoch = cs.Cycle
			cs.SoundLedger = cs.NominationLedger
			cs.SoundLedgerSegments = cs.NominationLedgerSegments
		}
		//
		//
		//
		//
		//
	}

	if cs.Phase <= cskinds.DurationPhaseNominate && cs.isNominationFinished() {
		//
		cs.joinPreballot(ledgerLevel, cs.Cycle)
		if hasDualThirds { //
			cs.joinPreendorse(ledgerLevel, cs.Cycle)
		}
	} else if cs.Phase == cskinds.DurationPhaseEndorse {
		//
		cs.attemptCompleteEndorse(ledgerLevel)
	}
}

//
func (cs *Status) attemptAppendBallot(ballot *kinds.Ballot, nodeUID p2p.ID) (bool, error) {
	appended, err := cs.appendBallot(ballot, nodeUID)
	//
	if err != nil {
		//
		//
		//
		//
		if ballotErr, ok := err.(*kinds.ErrBallotClashingBallots); ok {
			if cs.privateRatifierPublicKey == nil {
				return false, ErrPublicKeyIsNotCollection
			}

			if bytes.Equal(ballot.RatifierLocation, cs.privateRatifierPublicKey.Location()) {
				cs.Tracer.Fault(
					"REDACTED",
					"REDACTED", ballot.Level,
					"REDACTED", ballot.Cycle,
					"REDACTED", ballot.Kind,
				)

				return appended, err
			}

			//
			cs.eventpool.NotifyClashingBallots(ballotErr.BallotA, ballotErr.BallotBYTE)
			cs.Tracer.Diagnose(
				"REDACTED",
				"REDACTED", ballotErr.BallotA,
				"REDACTED", ballotErr.BallotBYTE,
			)

			return appended, err
		} else if errors.Is(err, kinds.ErrBallotNotCertainAutograph) {
			cs.Tracer.Diagnose("REDACTED", "REDACTED", err)
		} else if errors.Is(err, kinds.ErrCorruptBallotAddition) {
			cs.Tracer.Diagnose("REDACTED")
		} else {
			//
			//
			//
			//
			//
			cs.Tracer.Details("REDACTED", "REDACTED", err)
			return appended, ErrAppendingBallot
		}
	}

	return appended, nil
}

func (cs *Status) appendBallot(ballot *kinds.Ballot, nodeUID p2p.ID) (appended bool, err error) {
	cs.Tracer.Diagnose(
		"REDACTED",
		"REDACTED", ballot.Level,
		"REDACTED", ballot.Kind,
		"REDACTED", ballot.RatifierOrdinal,
		"REDACTED", cs.Level,
		"REDACTED", len(ballot.Addition),
		"REDACTED", len(ballot.AdditionAutograph),
		"REDACTED", nodeUID,
	)

	if ballot.Level < cs.Level || (ballot.Level == cs.Level && ballot.Cycle < cs.Cycle) {
		cs.stats.StampTardyBallot(ballot.Kind)
	}

	//
	//
	if ballot.Level+1 == cs.Level && ballot.Kind == engineproto.PreendorseKind {
		if cs.Phase != cskinds.DurationPhaseNewLevel {
			//
			cs.Tracer.Diagnose("REDACTED", "REDACTED", ballot)
			return appended, err
		}

		appended, err = cs.FinalEndorse.AppendBallot(ballot)
		if !appended {
			//
			if err == nil {
				cs.stats.ReplicatedBallot.Add(1)
			}
			return appended, err
		}

		cs.Tracer.Diagnose("REDACTED", "REDACTED", cs.FinalEndorse.StringBrief())
		if err := cs.eventBus.BroadcastEventBallot(kinds.EventDataBallot{Ballot: ballot}); err != nil {
			return appended, err
		}

		cs.evsw.TriggerEvent(kinds.EventBallot, ballot)

		//
		if cs.settings.OmitDeadlineEndorse && cs.FinalEndorse.HasAll() {
			//
			//
			cs.joinNewEpoch(cs.Level, 0)
		}

		return appended, err
	}

	//
	//
	if ballot.Level != cs.Level {
		cs.Tracer.Diagnose("REDACTED", "REDACTED", ballot.Level, "REDACTED", cs.Level, "REDACTED", nodeUID)
		return appended, err
	}

	//
	extensionActivated := cs.status.AgreementOptions.Iface.BallotPluginsActivated(ballot.Level)
	if extensionActivated {
		//
		//
		//

		var mineAddress []byte
		if cs.privateRatifierPublicKey != nil {
			mineAddress = cs.privateRatifierPublicKey.Location()
		}
		//
		//
		if ballot.Kind == engineproto.PreendorseKind && !ballot.LedgerUID.IsNil() &&
			!bytes.Equal(ballot.RatifierLocation, mineAddress) { //

			//
			//
			//
			//
			_, val := cs.status.Ratifiers.FetchByOrdinal(ballot.RatifierOrdinal)
			if val == nil { //
				valuesTally := cs.status.Ratifiers.Volume()
				cs.Tracer.Details("REDACTED",
					"REDACTED", nodeUID,
					"REDACTED", ballot.RatifierOrdinal,
					"REDACTED", valuesTally)
				return appended, ErrCorruptBallot{Cause: fmt.Sprintf("REDACTED", ballot.RatifierOrdinal, valuesTally)}
			}
			if err := ballot.ValidateAddition(cs.status.LedgerUID, val.PublicKey); err != nil {
				return false, err
			}

			err := cs.ledgerExecute.ValidateBallotAddition(context.TODO(), ballot)
			cs.stats.StampBallotAdditionAccepted(err == nil)
			if err != nil {
				return false, err
			}
		}
	} else if len(ballot.Addition) > 0 || len(ballot.AdditionAutograph) > 0 {
		//
		//
		//
		//
		//
		//

		return false, fmt.Errorf("REDACTED", ballot.Level, nodeUID)

	}

	level := cs.Level
	appended, err = cs.Ballots.AppendBallot(ballot, nodeUID, extensionActivated)
	if !appended {
		//

		//
		if err == nil {
			cs.stats.ReplicatedBallot.Add(1)
		}
		return appended, err
	}
	if ballot.Cycle == cs.Cycle {
		values := cs.status.Ratifiers
		_, val := values.FetchByOrdinal(ballot.RatifierOrdinal)
		cs.stats.StampBallotAccepted(ballot.Kind, val.PollingEnergy, values.SumPollingEnergy())
	}

	if err := cs.eventBus.BroadcastEventBallot(kinds.EventDataBallot{Ballot: ballot}); err != nil {
		return appended, err
	}
	cs.evsw.TriggerEvent(kinds.EventBallot, ballot)

	switch ballot.Kind {
	case engineproto.PreballotKind:
		preballots := cs.Ballots.Preballots(ballot.Cycle)
		cs.Tracer.Diagnose("REDACTED", "REDACTED", ballot, "REDACTED", preballots.StringBrief())

		//
		if ledgerUID, ok := preballots.DualThirdsBulk(); ok {
			//
			//
			//

			//
			//
			if (cs.LatchedLedger != nil) &&
				(cs.LatchedEpoch < ballot.Cycle) &&
				(ballot.Cycle <= cs.Cycle) &&
				!cs.LatchedLedger.DigestsTo(ledgerUID.Digest) {

				cs.Tracer.Diagnose("REDACTED", "REDACTED", cs.LatchedEpoch, "REDACTED", ballot.Cycle)

				cs.LatchedEpoch = -1
				cs.LatchedLedger = nil
				cs.LatchedLedgerSegments = nil

				if err := cs.eventBus.BroadcastEventRelease(cs.EpochStatusEvent()); err != nil {
					return appended, err
				}
			}

			//
			//
			if len(ledgerUID.Digest) != 0 && (cs.SoundEpoch < ballot.Cycle) && (ballot.Cycle == cs.Cycle) {
				if cs.NominationLedger.DigestsTo(ledgerUID.Digest) {
					cs.Tracer.Diagnose("REDACTED", "REDACTED", cs.SoundEpoch, "REDACTED", ballot.Cycle)
					cs.SoundEpoch = ballot.Cycle
					cs.SoundLedger = cs.NominationLedger
					cs.SoundLedgerSegments = cs.NominationLedgerSegments
				} else {
					cs.Tracer.Diagnose(
						"REDACTED",
						"REDACTED", log.NewIdleLedgerDigest(cs.NominationLedger),
						"REDACTED", ledgerUID.Digest,
					)

					//
					cs.NominationLedger = nil
				}

				if !cs.NominationLedgerSegments.HasHeading(ledgerUID.SegmentAssignHeading) {
					cs.NominationLedgerSegments = kinds.NewSegmentCollectionFromHeading(ledgerUID.SegmentAssignHeading)
				}

				cs.evsw.TriggerEvent(kinds.EventSoundLedger, cs.DurationStatus)
				if err := cs.eventBus.BroadcastEventSoundLedger(cs.EpochStatusEvent()); err != nil {
					return appended, err
				}
			}
		}

		//
		switch {
		case cs.Cycle < ballot.Cycle && preballots.HasDualThirdsAny():
			//
			cs.joinNewEpoch(level, ballot.Cycle)

		case cs.Cycle == ballot.Cycle && cskinds.EpochPhasePreballot <= cs.Phase: //
			ledgerUID, ok := preballots.DualThirdsBulk()
			if ok && (cs.isNominationFinished() || len(ledgerUID.Digest) == 0) {
				cs.joinPreendorse(level, ballot.Cycle)
			} else if preballots.HasDualThirdsAny() {
				cs.joinPreballotWait(level, ballot.Cycle)
			}

		case cs.Nomination != nil && 0 <= cs.Nomination.POLDuration && cs.Nomination.POLDuration == ballot.Cycle:
			//
			if cs.isNominationFinished() {
				cs.joinPreballot(level, cs.Cycle)
			}
		}

	case engineproto.PreendorseKind:
		preendorsements := cs.Ballots.Preendorsements(ballot.Cycle)
		cs.Tracer.Diagnose("REDACTED",
			"REDACTED", ballot.Level,
			"REDACTED", ballot.Cycle,
			"REDACTED", ballot.RatifierLocation.String(),
			"REDACTED", ballot.Timestamp,
			"REDACTED", preendorsements.TraceString())

		ledgerUID, ok := preendorsements.DualThirdsBulk()
		if ok {
			//
			cs.joinNewEpoch(level, ballot.Cycle)
			cs.joinPreendorse(level, ballot.Cycle)

			if len(ledgerUID.Digest) != 0 {
				cs.joinEndorse(level, ballot.Cycle)
				if cs.settings.OmitDeadlineEndorse && preendorsements.HasAll() {
					cs.joinNewEpoch(cs.Level, 0)
				}
			} else {
				cs.joinPreendorseWait(level, ballot.Cycle)
			}
		} else if cs.Cycle <= ballot.Cycle && preendorsements.HasDualThirdsAny() {
			cs.joinNewEpoch(level, ballot.Cycle)
			cs.joinPreendorseWait(level, ballot.Cycle)
		}

	default:
		panic(fmt.Sprintf("REDACTED", ballot.Kind))
	}

	return appended, err
}

//
func (cs *Status) attestBallot(
	messageKind engineproto.AttestedMessageKind,
	digest []byte,
	heading kinds.SegmentAssignHeading,
	ledger *kinds.Ledger,
) (*kinds.Ballot, error) {
	//
	//
	if err := cs.wal.PurgeAndAlign(); err != nil {
		return nil, err
	}

	if cs.privateRatifierPublicKey == nil {
		return nil, ErrPublicKeyIsNotCollection
	}

	address := cs.privateRatifierPublicKey.Location()
	valueIdx, _ := cs.Ratifiers.FetchByLocation(address)

	ballot := &kinds.Ballot{
		RatifierLocation: address,
		RatifierOrdinal:   valueIdx,
		Level:           cs.Level,
		Cycle:            cs.Cycle,
		Timestamp:        cs.ballotTime(),
		Kind:             messageKind,
		LedgerUID:          kinds.LedgerUID{Digest: digest, SegmentAssignHeading: heading},
	}

	extensionActivated := cs.status.AgreementOptions.Iface.BallotPluginsActivated(ballot.Level)
	if messageKind == engineproto.PreendorseKind && !ballot.LedgerUID.IsNil() {
		//
		//
		if extensionActivated {
			ext, err := cs.ledgerExecute.ExpandBallot(context.TODO(), ballot, ledger, cs.status)
			if err != nil {
				return nil, err
			}
			ballot.Addition = ext
		}
	}

	retrievable, err := kinds.AttestAndInspectBallot(ballot, cs.privateRatifier, cs.status.LedgerUID, extensionActivated && (messageKind == engineproto.PreendorseKind))
	if err != nil && !retrievable {
		panic(fmt.Sprintf("REDACTED", ballot, err))
	}

	return ballot, err
}

func (cs *Status) ballotTime() time.Time {
	now := engineclock.Now()
	minimumBallotTime := now
	//
	const timeIota = time.Millisecond
	//
	//
	if cs.LatchedLedger != nil {
		//
		//
		minimumBallotTime = cs.LatchedLedger.Time.Add(timeIota)
	} else if cs.NominationLedger != nil {
		minimumBallotTime = cs.NominationLedger.Time.Add(timeIota)
	}

	if now.After(minimumBallotTime) {
		return now
	}
	return minimumBallotTime
}

//
//
func (cs *Status) attestAppendBallot(
	messageKind engineproto.AttestedMessageKind,
	digest []byte,
	heading kinds.SegmentAssignHeading,
	ledger *kinds.Ledger,
) {
	if cs.privateRatifier == nil { //
		return
	}

	if cs.privateRatifierPublicKey == nil {
		//
		cs.Tracer.Fault(fmt.Sprintf("REDACTED", ErrPublicKeyIsNotCollection))
		return
	}

	//
	if !cs.Ratifiers.HasLocation(cs.privateRatifierPublicKey.Location()) {
		return
	}

	//
	ballot, err := cs.attestBallot(messageKind, digest, heading, ledger)
	if err != nil {
		cs.Tracer.Fault("REDACTED", "REDACTED", cs.Level, "REDACTED", cs.Cycle, "REDACTED", ballot, "REDACTED", err)
		return
	}
	hasExtension := len(ballot.AdditionAutograph) > 0
	extensionActivated := cs.status.AgreementOptions.Iface.BallotPluginsActivated(ballot.Level)
	if ballot.Kind == engineproto.PreendorseKind && !ballot.LedgerUID.IsNil() && hasExtension != extensionActivated {
		panic(fmt.Errorf("REDACTED",
			hasExtension, extensionActivated, ballot.Level, ballot.Kind))
	}
	cs.transmitIntrinsicSignal(messageDetails{&BallotSignal{ballot}, "REDACTED"})
	cs.Tracer.Diagnose("REDACTED", "REDACTED", cs.Level, "REDACTED", cs.Cycle, "REDACTED", ballot)
}

//
//
//
func (cs *Status) modifyPrivateRatifierPublicKey() error {
	if cs.privateRatifier == nil {
		return nil
	}

	publicKey, err := cs.privateRatifier.FetchPublicKey()
	if err != nil {
		return err
	}
	cs.privateRatifierPublicKey = publicKey
	return nil
}

//
func (cs *Status) inspectRepeatAttestingHazard(level int64) error {
	if cs.privateRatifier != nil && cs.privateRatifierPublicKey != nil && cs.settings.RepeatAttestInspectLevel > 0 && level > 0 {
		valueAddress := cs.privateRatifierPublicKey.Location()
		repeatAttestInspectLevel := cs.settings.RepeatAttestInspectLevel
		if repeatAttestInspectLevel > level {
			repeatAttestInspectLevel = level
		}

		for i := int64(1); i < repeatAttestInspectLevel; i++ {
			finalEndorse := cs.ledgerDepot.ImportViewedEndorse(level - i)
			if finalEndorse != nil {
				for signatureIdx, s := range finalEndorse.Endorsements {
					if s.LedgerUIDMark == kinds.LedgerUIDMarkEndorse && bytes.Equal(s.RatifierLocation, valueAddress) {
						cs.Tracer.Details("REDACTED", "REDACTED", s, "REDACTED", signatureIdx, "REDACTED", level-i)
						return ErrAutographLocatedInElapsedLedgers
					}
				}
			}
		}
	}

	return nil
}

//
//
func (cs *Status) issuePreendorseDeadlineStats(duration int32) {
	//
	//
	sumBallotsGathered := 0
	sumPollingEnergyGathered := int64(0)

	for _, ballot := range cs.Ballots.Preendorsements(duration).Catalog() {
		sumBallotsGathered++
		_, val := cs.Ratifiers.FetchByLocation(ballot.RatifierLocation)
		if val != nil {
			sumPollingEnergyGathered += val.PollingEnergy
		}
	}

	//
	sumFeasiblePollingEnergy := cs.Ratifiers.SumPollingEnergy()
	var equityFraction float64
	if sumFeasiblePollingEnergy > 0 {
		equityFraction = float64(sumPollingEnergyGathered) / float64(sumFeasiblePollingEnergy)
	}

	//
	cs.stats.PreendorsementsTallied.Set(float64(sumBallotsGathered))
	cs.stats.PreendorsementsStakingRatio.Set(equityFraction)

	cs.Tracer.Diagnose("REDACTED",
		"REDACTED", sumBallotsGathered,
		"REDACTED", equityFraction)
}

func (cs *Status) computePreendorseSignalDeferralStats() {
	if cs.Nomination == nil {
		return
	}

	ps := cs.Ballots.Preendorsements(cs.Cycle)
	pl := ps.Catalog()

	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Timestamp.Before(pl[j].Timestamp)
	})

	var pollingEnergyViewed int64
	for _, v := range pl {
		_, val := cs.Ratifiers.FetchByLocation(v.RatifierLocation)
		pollingEnergyViewed += val.PollingEnergy
		if pollingEnergyViewed >= cs.Ratifiers.SumPollingEnergy()*2/3+1 {
			cs.stats.AssemblyPreendorseDeferral.With("REDACTED", cs.Ratifiers.FetchRecommender().Location.String()).Set(v.Timestamp.Sub(cs.Nomination.Timestamp).Seconds())
			break
		}
	}
}

func (cs *Status) computePreballotSignalDeferralStats() {
	if cs.Nomination == nil {
		return
	}

	ps := cs.Ballots.Preballots(cs.Cycle)
	pl := ps.Catalog()

	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Timestamp.Before(pl[j].Timestamp)
	})

	var pollingEnergyViewed int64
	for _, v := range pl {
		_, val := cs.Ratifiers.FetchByLocationMut(v.RatifierLocation)
		pollingEnergyViewed += val.PollingEnergy
		if pollingEnergyViewed >= cs.Ratifiers.SumPollingEnergy()*2/3+1 {
			cs.stats.AssemblyPreballotDeferral.With("REDACTED", cs.Ratifiers.FetchRecommender().Location.String()).Set(v.Timestamp.Sub(cs.Nomination.Timestamp).Seconds())
			break
		}
	}
	if ps.HasAll() {
		cs.stats.CompletePreballotDeferral.With("REDACTED", cs.Ratifiers.FetchRecommender().Location.String()).Set(pl[len(pl)-1].Timestamp.Sub(cs.Nomination.Timestamp).Seconds())
	}
}

//

func ContrastHRS(h1 int64, r1 int32, s1 cskinds.DurationPhaseKind, h2 int64, r2 int32, s2 cskinds.DurationPhaseKind) int {
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
func remediateJournalEntry(src, dst string) error {
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
		dec = NewJournalParser(in)
		enc = NewJournalSerializer(out)
	)

	//
	for {
		msg, err := dec.Parse()
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
