package agreement

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	cometfaults "github.com/valkyrieworks/kinds/faults"

	cskinds "github.com/valkyrieworks/agreement/kinds"
	"github.com/valkyrieworks/utils/bits"
	cometsignals "github.com/valkyrieworks/utils/events"
	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/utils/log"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/p2p"
	cometconnect "github.com/valkyrieworks/schema/consensuscore/agreement"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

const (
	StatusStream       = byte(0x20)
	DataStream        = byte(0x21)
	BallotStream        = byte(0x22)
	BallotAssignBitsStream = byte(0x23)

	maximumMessageVolume = 1048576 //

	ledgersToInputToTransformSoundNode = 10000
	ballotsToInputToTransformSoundNode  = 10000
)

//

//
type Handler struct {
	p2p.RootHandler //

	connectS *Status

	waitAlign atomic.Bool
	eventBus *kinds.EventBus

	rsMutex         engineconnect.ReadwriteLock
	rs            cskinds.DurationStatus //
	primaryLevel atomic.Int64

	agreementOptions atomic.Pointer[kinds.AgreementOptions] //

	Stats *Stats
}

type HandlerSetting func(*Handler)

//
func NewHandler(agreementStatus *Status, waitAlign bool, options ...HandlerSetting) *Handler {
	connectReader := &Handler{
		connectS:          agreementStatus,
		waitAlign:      atomic.Bool{},
		rs:            agreementStatus.fetchDurationStatus(),
		primaryLevel: atomic.Int64{},
		Stats:       NoopStats(),
	}
	//
	options := agreementStatus.status.AgreementOptions
	connectReader.agreementOptions.Store(&options)
	connectReader.primaryLevel.Store(agreementStatus.status.PrimaryLevel)
	connectReader.RootHandler = *p2p.NewRootHandler("REDACTED", connectReader)
	if waitAlign {
		connectReader.waitAlign.Store(true)
	}

	for _, setting := range options {
		setting(connectReader)
	}

	return connectReader
}

//
//
func (connectReader *Handler) OnBegin() error {
	if connectReader.WaitAlign() {
		connectReader.Tracer.Details("REDACTED")
	}

	//
	go connectReader.nodeMetricsProcess()

	connectReader.enrollToMulticastEvents()

	if !connectReader.WaitAlign() {
		err := connectReader.connectS.Begin()
		if err != nil {
			return err
		}
	}

	return nil
}

//
//
func (connectReader *Handler) OnHalt() {
	connectReader.cancelFromMulticastEvents()
	if err := connectReader.connectS.Halt(); err != nil {
		connectReader.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	if !connectReader.WaitAlign() {
		connectReader.connectS.Wait()
	}
}

//
//
func (connectReader *Handler) RouterToAgreement(status sm.Status, omitJournal bool) {
	connectReader.Tracer.Details("REDACTED")

	//
	func() {
		//
		connectReader.connectS.mtx.Lock()
		defer connectReader.connectS.mtx.Unlock()
		//
		if status.FinalLedgerLevel > 0 {
			connectReader.connectS.reassembleFinalEndorse(status)
		}

		//
		//
		connectReader.connectS.modifyToStatus(status)
	}()

	//
	connectReader.waitAlign.Store(false)

	if omitJournal {
		connectReader.connectS.executeJournalOvertake = false
	}

	//
	err := connectReader.connectS.Begin()
	if err != nil {
		panic(fmt.Sprintf(`REDACTEDv

REDACTED:
REDACTED

REDACTED:
REDACTED`, err, connectReader.connectS, connectReader))
	}
}

//
func (connectReader *Handler) FetchStreams() []*p2p.StreamDefinition {
	//
	return []*p2p.StreamDefinition{
		{
			ID:                  StatusStream,
			Urgency:            6,
			TransmitBufferVolume:   100,
			AcceptSignalVolume: maximumMessageVolume,
			SignalKind:         &cometconnect.Signal{},
		},
		{
			ID: DataStream, //
			//
			Urgency:            10,
			TransmitBufferVolume:   100,
			AcceptBufferVolume:  50 * 4096,
			AcceptSignalVolume: maximumMessageVolume,
			SignalKind:         &cometconnect.Signal{},
		},
		{
			ID:                  BallotStream,
			Urgency:            7,
			TransmitBufferVolume:   100,
			AcceptBufferVolume:  100 * 100,
			AcceptSignalVolume: maximumMessageVolume,
			SignalKind:         &cometconnect.Signal{},
		},
		{
			ID:                  BallotAssignBitsStream,
			Urgency:            1,
			TransmitBufferVolume:   2,
			AcceptBufferVolume:  1024,
			AcceptSignalVolume: maximumMessageVolume,
			SignalKind:         &cometconnect.Signal{},
		},
	}
}

//
func (connectReader *Handler) InitNode(node p2p.Node) p2p.Node {
	nodeStatus := NewNodeStatus(node).AssignTracer(connectReader.Tracer)
	node.Set(kinds.NodeStatusKey, nodeStatus)
	return node
}

//
//
func (connectReader *Handler) AppendNode(node p2p.Node) {
	if !connectReader.IsActive() {
		return
	}

	nodeStatus, ok := node.Get(kinds.NodeStatusKey).(*NodeStatus)
	if !ok {
		panic(fmt.Sprintf("REDACTED", node))
	}
	//
	go connectReader.gossipDataProcess(node, nodeStatus)
	go connectReader.gossipBallotsProcess(node, nodeStatus)
	go connectReader.inquireMaj23process(node, nodeStatus)

	//
	//
	if !connectReader.WaitAlign() {
		connectReader.transmitNewDurationPhaseSignal(node)
	}
}

//
func (connectReader *Handler) DeleteNode(p2p.Node, any) {
	if !connectReader.IsActive() {
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
func (connectReader *Handler) Accept(e p2p.Packet) {
	if !connectReader.IsActive() {
		connectReader.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.StreamUID)
		return
	}
	msg, err := MessageFromSchema(e.Signal)
	if err != nil {
		connectReader.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", e.StreamUID, "REDACTED", err)
		connectReader.Router.HaltNodeForFault(e.Src, err)
		return
	}

	if err = msg.CertifySimple(); err != nil {
		connectReader.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", e.Signal, "REDACTED", err)
		connectReader.Router.HaltNodeForFault(e.Src, err)
		return
	}

	connectReader.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.StreamUID, "REDACTED", msg)

	//
	ps, ok := e.Src.Get(kinds.NodeStatusKey).(*NodeStatus)
	if !ok {
		panic(fmt.Sprintf("REDACTED", e.Src))
	}

	switch e.StreamUID {
	case StatusStream:
		switch msg := msg.(type) {
		case *NewDurationPhaseSignal:
			primaryLevel := connectReader.primaryLevel.Load()
			if err = msg.CertifyLevel(primaryLevel); err != nil {
				connectReader.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", msg, "REDACTED", err)
				connectReader.Router.HaltNodeForFault(e.Src, err)
				return
			}
			ps.ExecuteNewDurationPhaseSignal(msg)
			connectReader.connectS.metricsMessageBuffer <- messageDetails{msg, e.Src.ID()}
		case *NewSoundLedgerSignal:
			ps.ExecuteNewSoundLedgerSignal(msg)
		case *HasBallotSignal:
			ps.ExecuteHasBallotSignal(msg)
		case *BallotAssignMaj23signal:
			rs := connectReader.fetchDurationStatus()
			level, ballots := rs.Level, rs.Ballots
			if level != msg.Level {
				return
			}
			//
			err := ballots.AssignNodeMaj23(msg.Cycle, msg.Kind, ps.node.ID(), msg.LedgerUID)
			if err != nil {
				connectReader.Router.HaltNodeForFault(e.Src, err)
				return
			}
			//
			//
			var ourBallots *bits.BitList
			switch msg.Kind {
			case engineproto.PreballotKind:
				ourBallots = ballots.Preballots(msg.Cycle).BitListByLedgerUID(msg.LedgerUID)
			case engineproto.PreendorseKind:
				ourBallots = ballots.Preendorsements(msg.Cycle).BitListByLedgerUID(msg.LedgerUID)
			default:
				panic("REDACTED")
			}
			eMessage := &cometconnect.BallotAssignBits{
				Level:  msg.Level,
				Cycle:   msg.Cycle,
				Kind:    msg.Kind,
				LedgerUID: msg.LedgerUID.ToSchema(),
			}
			if ballots := ourBallots.ToSchema(); ballots != nil {
				eMessage.Ballots = *ballots
			}
			e.Src.AttemptTransmit(p2p.Packet{
				StreamUID: BallotAssignBitsStream,
				Signal:   eMessage,
			})
		default:
			connectReader.Tracer.Fault(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
		}

	case DataStream:
		if connectReader.WaitAlign() {
			connectReader.Tracer.Details("REDACTED", "REDACTED", msg)
			return
		}
		switch msg := msg.(type) {
		case *NominationSignal:
			options := connectReader.agreementOptions.Load()
			maximumOctets := options.Ledger.MaximumOctets
			if err := msg.Nomination.CertifyLedgerVolume(maximumOctets); err != nil {
				connectReader.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", msg.Nomination.Level)
				connectReader.Router.HaltNodeForFault(e.Src, ErrNominationTooNumerousSegments)
				return
			}

			ps.AssignHasNomination(msg.Nomination)
			connectReader.connectS.nodeMessageBuffer <- messageDetails{msg, e.Src.ID()}
		case *NominationPOLSignal:
			ps.ExecuteNominationPOLSignal(msg)
		case *LedgerSegmentSignal:
			ps.AssignHasNominationLedgerSegment(msg.Level, msg.Cycle, int(msg.Segment.Ordinal))
			connectReader.Stats.LedgerSegments.With("REDACTED", string(e.Src.ID())).Add(1)
			connectReader.connectS.nodeMessageBuffer <- messageDetails{msg, e.Src.ID()}
		default:
			connectReader.Tracer.Fault(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
		}

	case BallotStream:
		if connectReader.WaitAlign() {
			connectReader.Tracer.Details("REDACTED", "REDACTED", msg)
			return
		}
		switch msg := msg.(type) {
		case *BallotSignal:
			rs := connectReader.fetchDurationStatus()

			level, valueVolume, finalEndorseVolume := rs.Level, rs.Ratifiers.Volume(), rs.FinalEndorse.Volume()
			ps.AssignHasBallotFromNode(msg.Ballot, level, valueVolume, finalEndorseVolume)

			connectReader.connectS.nodeMessageBuffer <- messageDetails{msg, e.Src.ID()}

		default:
			//
			connectReader.Tracer.Fault(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
		}

	case BallotAssignBitsStream:
		if connectReader.WaitAlign() {
			connectReader.Tracer.Details("REDACTED", "REDACTED", msg)
			return
		}
		switch msg := msg.(type) {
		case *BallotAssignBitsSignal:
			rs := connectReader.fetchDurationStatus()

			level, ballots := rs.Level, rs.Ballots

			if level == msg.Level {
				var ourBallots *bits.BitList
				switch msg.Kind {
				case engineproto.PreballotKind:
					ourBallots = ballots.Preballots(msg.Cycle).BitListByLedgerUID(msg.LedgerUID)
				case engineproto.PreendorseKind:
					ourBallots = ballots.Preendorsements(msg.Cycle).BitListByLedgerUID(msg.LedgerUID)
				default:
					panic("REDACTED")
				}
				ps.ExecuteBallotAssignBitsSignal(msg, ourBallots)
			} else {
				ps.ExecuteBallotAssignBitsSignal(msg, nil)
			}
		default:
			//
			connectReader.Tracer.Fault(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
		}

	default:
		connectReader.Tracer.Fault(fmt.Sprintf("REDACTED", e.StreamUID))
	}
}

//
func (connectReader *Handler) AssignEventBus(b *kinds.EventBus) {
	connectReader.eventBus = b
	connectReader.connectS.AssignEventBus(b)
}

//
func (connectReader *Handler) WaitAlign() bool {
	return connectReader.waitAlign.Load()
}

//

//
//
//
func (connectReader *Handler) enrollToMulticastEvents() {
	const enrollee = "REDACTED"
	err := connectReader.connectS.evsw.AppendObserverForEvent(
		enrollee,
		kinds.EventNewDurationPhase,
		func(data cometsignals.EventData) {
			rs := data.(cskinds.DurationStatus)

			//
			connectReader.modifyDurationStatus(&rs)

			connectReader.multicastNewDurationPhaseSignal(&rs)
		},
	)
	if err != nil {
		connectReader.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	err = connectReader.connectS.evsw.AppendObserverForEvent(
		enrollee,
		kinds.EventSoundLedger,
		func(data cometsignals.EventData) {
			rs := data.(cskinds.DurationStatus)

			//
			connectReader.modifyDurationStatus(&rs)

			connectReader.multicastNewSoundLedgerSignal(&rs)
		},
	)
	if err != nil {
		connectReader.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	err = connectReader.connectS.evsw.AppendObserverForEvent(
		enrollee,
		kinds.EventBallot,
		func(data cometsignals.EventData) {
			connectReader.multicastHasBallotSignal(data.(*kinds.Ballot))

			//
			//
			//
			//
			rs := connectReader.connectS.fetchDurationStatus()
			connectReader.modifyDurationStatus(&rs)
		},
	)
	if err != nil {
		connectReader.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	err = connectReader.connectS.evsw.AppendObserverForEvent(
		enrollee,
		kinds.EventNewAgreementOptions,
		func(data cometsignals.EventData) {
			agreementOptions := data.(kinds.AgreementOptions)

			//
			connectReader.modifyAgreementOptions(agreementOptions)
		},
	)
	if err != nil {
		connectReader.Tracer.Fault("REDACTED", "REDACTED", err)
	}
}

//
func (connectReader *Handler) modifyAgreementOptions(agreementOptions kinds.AgreementOptions) {
	options := agreementOptions //
	connectReader.agreementOptions.Store(&options)
}

//
func (connectReader *Handler) modifyDurationStatus(rs *cskinds.DurationStatus) {
	connectReader.rsMutex.Lock()
	connectReader.rs = *rs //
	connectReader.rsMutex.Unlock()
}

func (connectReader *Handler) cancelFromMulticastEvents() {
	const enrollee = "REDACTED"
	connectReader.connectS.evsw.DeleteObserver(enrollee)
}

func (connectReader *Handler) multicastNewDurationPhaseSignal(rs *cskinds.DurationStatus) {
	nrsMessage := createDurationPhaseSignal(rs)
	go func() {
		connectReader.Router.MulticastAsync(p2p.Packet{
			StreamUID: StatusStream,
			Signal:   nrsMessage,
		})
	}()
}

func (connectReader *Handler) multicastNewSoundLedgerSignal(rs *cskinds.DurationStatus) {
	psh := rs.NominationLedgerSegments.Heading()
	csMessage := &cometconnect.NewSoundLedger{
		Level:             rs.Level,
		Cycle:              rs.Cycle,
		LedgerSegmentAssignHeading: psh.ToSchema(),
		LedgerSegments:         rs.NominationLedgerSegments.BitList().ToSchema(),
		IsEndorse:           rs.Phase == cskinds.DurationPhaseEndorse,
	}
	go func() {
		connectReader.Router.MulticastAsync(p2p.Packet{
			StreamUID: StatusStream,
			Signal:   csMessage,
		})
	}()
}

//
func (connectReader *Handler) multicastHasBallotSignal(ballot *kinds.Ballot) {
	msg := &cometconnect.HasBallot{
		Level: ballot.Level,
		Cycle:  ballot.Cycle,
		Kind:   ballot.Kind,
		Ordinal:  ballot.RatifierOrdinal,
	}

	go func() {
		connectReader.Router.AttemptMulticast(p2p.Packet{
			StreamUID: StatusStream,
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

func createDurationPhaseSignal(rs *cskinds.DurationStatus) (nrsMessage *cometconnect.NewDurationPhase) {
	nrsMessage = &cometconnect.NewDurationPhase{
		Level:                rs.Level,
		Cycle:                 rs.Cycle,
		Phase:                  uint32(rs.Phase),
		MomentsSinceBeginTime: int64(time.Since(rs.BeginTime).Seconds()),
		FinalEndorseDuration:       rs.FinalEndorse.FetchDuration(),
	}
	return
}

func (connectReader *Handler) transmitNewDurationPhaseSignal(node p2p.Node) {
	rs := connectReader.fetchDurationStatus()
	nrsMessage := createDurationPhaseSignal(&rs)
	node.Transmit(p2p.Packet{
		StreamUID: StatusStream,
		Signal:   nrsMessage,
	})
}

func (connectReader *Handler) fetchDurationStatus() cskinds.DurationStatus {
	connectReader.rsMutex.RLock()
	defer connectReader.rsMutex.RUnlock()
	return connectReader.rs
}

//
//

func (connectReader *Handler) gossipDataProcess(node p2p.Node, ps *NodeStatus) {
	tracer := connectReader.Tracer.With("REDACTED", node)

EXTERNAL_CYCLE:
	for {
		//
		if !node.IsActive() || !connectReader.IsActive() {
			return
		}
		rs := connectReader.fetchDurationStatus()
		prs := ps.FetchDurationStatus()

		//
		//
		//
		//

		if segment, resumeCycle := selectSegmentToTransmit(tracer, connectReader.connectS.ledgerDepot, &rs, ps, prs); segment != nil {
			//
			//
			if ps.TransmitSegmentAssignHasSegment(segment, prs) || resumeCycle {
				continue EXTERNAL_CYCLE
			}
		} else if resumeCycle {
			//
			continue EXTERNAL_CYCLE
		}

		//
		//
		//
		//

		levelDurationAlign := (rs.Level == prs.Level) && (rs.Cycle == prs.Cycle)
		nominationToTransmit := rs.Nomination != nil && !prs.Nomination

		if levelDurationAlign && nominationToTransmit {
			ps.TransmitNominationAssignHasNomination(tracer, &rs, prs)
			continue EXTERNAL_CYCLE
		}

		//
		time.Sleep(connectReader.connectS.settings.NodeGossipPausePeriod)
	}
}

func (connectReader *Handler) gossipBallotsProcess(node p2p.Node, ps *NodeStatus) {
	tracer := connectReader.Tracer.With("REDACTED", node)

	//
	dormant := 0

EXTERNAL_CYCLE:
	for {
		//
		if !node.IsActive() || !connectReader.IsActive() {
			return
		}
		rs := connectReader.fetchDurationStatus()
		prs := ps.FetchDurationStatus()

		switch dormant {
		case 1: //
			dormant = 2
		case 2: //
			dormant = 0
		}

		if ballot := selectBallotToTransmit(tracer, connectReader.connectS, &rs, ps, prs); ballot != nil {
			if ps.transmitBallotAssignHasBallot(ballot) {
				continue EXTERNAL_CYCLE
			}
			tracer.Diagnose("REDACTED",
				"REDACTED", prs.Level,
				"REDACTED", ballot,
			)
		}

		switch dormant {
		case 0:
			//
			dormant = 1
			tracer.Diagnose("REDACTED", "REDACTED", rs.Level, "REDACTED", prs.Level,
				"REDACTED", rs.Ballots.Preballots(rs.Cycle).BitList(), "REDACTED", prs.Preballots,
				"REDACTED", rs.Ballots.Preendorsements(rs.Cycle).BitList(), "REDACTED", prs.Preendorsements)
		case 2:
			//
			dormant = 1
		}

		time.Sleep(connectReader.connectS.settings.NodeGossipPausePeriod)
	}
}

//
//
func (connectReader *Handler) inquireMaj23process(node p2p.Node, ps *NodeStatus) {
EXTERNAL_CYCLE:
	for {
		//
		if !node.IsActive() || !connectReader.IsActive() {
			return
		}

		//
		{
			rs := connectReader.fetchDurationStatus()
			prs := ps.FetchDurationStatus()
			if rs.Level == prs.Level {
				if maj23, ok := rs.Ballots.Preballots(prs.Cycle).DualThirdsBulk(); ok {

					node.AttemptTransmit(p2p.Packet{
						StreamUID: StatusStream,
						Signal: &cometconnect.BallotAssignMaj23{
							Level:  prs.Level,
							Cycle:   prs.Cycle,
							Kind:    engineproto.PreballotKind,
							LedgerUID: maj23.ToSchema(),
						},
					})
					time.Sleep(connectReader.connectS.settings.NodeInquireMaj23pausePeriod)
				}
			}
		}

		//
		{
			rs := connectReader.fetchDurationStatus()
			prs := ps.FetchDurationStatus()
			if rs.Level == prs.Level {
				if maj23, ok := rs.Ballots.Preendorsements(prs.Cycle).DualThirdsBulk(); ok {
					node.AttemptTransmit(p2p.Packet{
						StreamUID: StatusStream,
						Signal: &cometconnect.BallotAssignMaj23{
							Level:  prs.Level,
							Cycle:   prs.Cycle,
							Kind:    engineproto.PreendorseKind,
							LedgerUID: maj23.ToSchema(),
						},
					})
					time.Sleep(connectReader.connectS.settings.NodeInquireMaj23pausePeriod)
				}
			}
		}

		//
		{
			rs := connectReader.fetchDurationStatus()
			prs := ps.FetchDurationStatus()
			if rs.Level == prs.Level && prs.NominationPOLDuration >= 0 {
				if maj23, ok := rs.Ballots.Preballots(prs.NominationPOLDuration).DualThirdsBulk(); ok {

					node.AttemptTransmit(p2p.Packet{
						StreamUID: StatusStream,
						Signal: &cometconnect.BallotAssignMaj23{
							Level:  prs.Level,
							Cycle:   prs.NominationPOLDuration,
							Kind:    engineproto.PreballotKind,
							LedgerUID: maj23.ToSchema(),
						},
					})
					time.Sleep(connectReader.connectS.settings.NodeInquireMaj23pausePeriod)
				}
			}
		}

		//
		//

		//
		{
			prs := ps.FetchDurationStatus()
			if prs.OvertakeEndorseDuration != -1 && prs.Level > 0 && prs.Level <= connectReader.connectS.ledgerDepot.Level() &&
				prs.Level >= connectReader.connectS.ledgerDepot.Root() {
				if endorse := connectReader.connectS.ImportEndorse(prs.Level); endorse != nil {
					node.AttemptTransmit(p2p.Packet{
						StreamUID: StatusStream,
						Signal: &cometconnect.BallotAssignMaj23{
							Level:  prs.Level,
							Cycle:   endorse.Cycle,
							Kind:    engineproto.PreendorseKind,
							LedgerUID: endorse.LedgerUID.ToSchema(),
						},
					})
					time.Sleep(connectReader.connectS.settings.NodeInquireMaj23pausePeriod)
				}
			}
		}

		time.Sleep(connectReader.connectS.settings.NodeInquireMaj23pausePeriod)

		continue EXTERNAL_CYCLE
	}
}

//
//
//
func selectSegmentToTransmit(
	tracer log.Tracer,
	ledgerDepot sm.LedgerDepot,
	rs *cskinds.DurationStatus,
	ps *NodeStatus,
	prs *cskinds.NodeDurationStatus,
) (*kinds.Segment, bool) {
	//
	if rs.NominationLedgerSegments.HasHeading(prs.NominationLedgerSegmentAssignHeading) {
		if ordinal, ok := rs.NominationLedgerSegments.BitList().Sub(prs.NominationLedgerSegments.Clone()).SelectArbitrary(); ok {
			segment := rs.NominationLedgerSegments.FetchSegment(ordinal)
			//
			return segment, true
		}
	}

	//
	ledgerDepotRoot := ledgerDepot.Root()
	if ledgerDepotRoot > 0 &&
		0 < prs.Level && prs.Level < rs.Level &&
		prs.Level >= ledgerDepotRoot {
		levelTracer := tracer.With("REDACTED", prs.Level)

		//
		if prs.NominationLedgerSegments == nil {
			ledgerMeta := ledgerDepot.ImportLedgerMeta(prs.Level)
			if ledgerMeta == nil {
				levelTracer.Fault("REDACTED",
					"REDACTED", ledgerDepotRoot, "REDACTED", ledgerDepot.Level())
				return nil, false
			}
			ps.InitNominationLedgerSegments(ledgerMeta.LedgerUID.SegmentAssignHeading)
			//
			return nil, true //
		}
		segment := selectSegmentForOvertake(levelTracer, rs, prs, ledgerDepot)
		if segment != nil {
			//
			return segment, false
		}
	}

	return nil, false
}

func selectSegmentForOvertake(
	tracer log.Tracer,
	rs *cskinds.DurationStatus,
	prs *cskinds.NodeDurationStatus,
	ledgerDepot sm.LedgerDepot,
) *kinds.Segment {
	ordinal, ok := prs.NominationLedgerSegments.Not().SelectArbitrary()
	if !ok {
		return nil
	}
	//
	ledgerMeta := ledgerDepot.ImportLedgerMeta(prs.Level)
	if ledgerMeta == nil {
		tracer.Fault("REDACTED", "REDACTED", rs.Level,
			"REDACTED", ledgerDepot.Root(), "REDACTED", ledgerDepot.Level())
		return nil
	} else if !ledgerMeta.LedgerUID.SegmentAssignHeading.Matches(prs.NominationLedgerSegmentAssignHeading) {
		tracer.Details("REDACTED",
			"REDACTED", ledgerMeta.LedgerUID.SegmentAssignHeading, "REDACTED", prs.NominationLedgerSegmentAssignHeading)
		return nil
	}
	//
	segment := ledgerDepot.ImportLedgerSegment(prs.Level, ordinal)
	if segment == nil {
		tracer.Fault("REDACTED", "REDACTED", ordinal,
			"REDACTED", ledgerMeta.LedgerUID.SegmentAssignHeading, "REDACTED", prs.NominationLedgerSegmentAssignHeading)
		return nil
	}
	return segment
}

func selectBallotToTransmit(
	tracer log.Tracer,
	connectS *Status,
	rs *cskinds.DurationStatus,
	ps *NodeStatus,
	prs *cskinds.NodeDurationStatus,
) *kinds.Ballot {
	//
	if rs.Level == prs.Level {
		levelTracer := tracer.With("REDACTED", prs.Level)
		return selectBallotOngoingLevel(levelTracer, rs, prs, ps)
	}

	//
	//
	if prs.Level != 0 && rs.Level == prs.Level+1 {
		if ballot := ps.SelectBallotToTransmit(rs.FinalEndorse); ballot != nil {
			tracer.Diagnose("REDACTED", "REDACTED", prs.Level)
			return ballot
		}
	}

	//
	//
	ledgerDepotRoot := connectS.ledgerDepot.Root()
	if ledgerDepotRoot > 0 && prs.Level != 0 && rs.Level >= prs.Level+2 && prs.Level >= ledgerDepotRoot {
		//
		//
		var ec *kinds.ExpandedEndorse
		var veActivated bool
		func() {
			connectS.mtx.RLock()
			defer connectS.mtx.RUnlock()
			veActivated = connectS.status.AgreementOptions.Iface.BallotPluginsActivated(prs.Level)
		}()
		if veActivated {
			ec = connectS.ledgerDepot.ImportLedgerExpandedEndorse(prs.Level)
		} else {
			c := connectS.ledgerDepot.ImportLedgerEndorse(prs.Level)
			if c == nil {
				return nil
			}
			ec = c.EncapsulatedExpandedEndorse()
		}
		if ec == nil {
			return nil
		}
		if ballot := ps.SelectBallotToTransmit(ec); ballot != nil {
			tracer.Diagnose("REDACTED", "REDACTED", prs.Level)
			return ballot
		}
	}
	return nil
}

func selectBallotOngoingLevel(
	tracer log.Tracer,
	rs *cskinds.DurationStatus,
	prs *cskinds.NodeDurationStatus,
	ps *NodeStatus,
) *kinds.Ballot {
	//
	if prs.Phase == cskinds.DurationPhaseNewLevel {
		if ballot := ps.SelectBallotToTransmit(rs.FinalEndorse); ballot != nil {
			tracer.Diagnose("REDACTED")
			return ballot
		}
	}
	//
	if prs.Phase <= cskinds.DurationPhaseNominate && prs.Cycle != -1 && prs.Cycle <= rs.Cycle && prs.NominationPOLDuration != -1 {
		if polPreballots := rs.Ballots.Preballots(prs.NominationPOLDuration); polPreballots != nil {
			if ballot := ps.SelectBallotToTransmit(polPreballots); ballot != nil {
				tracer.Diagnose("REDACTED",
					"REDACTED", prs.NominationPOLDuration)
				return ballot
			}
		}
	}
	//
	if prs.Phase <= cskinds.DurationPhasePreballotWait && prs.Cycle != -1 && prs.Cycle <= rs.Cycle {
		if ballot := ps.SelectBallotToTransmit(rs.Ballots.Preballots(prs.Cycle)); ballot != nil {
			tracer.Diagnose("REDACTED", "REDACTED", prs.Cycle)
			return ballot
		}
	}
	//
	if prs.Phase <= cskinds.DurationPhasePreendorseWait && prs.Cycle != -1 && prs.Cycle <= rs.Cycle {
		if ballot := ps.SelectBallotToTransmit(rs.Ballots.Preendorsements(prs.Cycle)); ballot != nil {
			tracer.Diagnose("REDACTED", "REDACTED", prs.Cycle)
			return ballot
		}
	}
	//
	if prs.Cycle != -1 && prs.Cycle <= rs.Cycle {
		if ballot := ps.SelectBallotToTransmit(rs.Ballots.Preballots(prs.Cycle)); ballot != nil {
			tracer.Diagnose("REDACTED", "REDACTED", prs.Cycle)
			return ballot
		}
	}
	//
	if prs.NominationPOLDuration != -1 {
		if polPreballots := rs.Ballots.Preballots(prs.NominationPOLDuration); polPreballots != nil {
			if ballot := ps.SelectBallotToTransmit(polPreballots); ballot != nil {
				tracer.Diagnose("REDACTED",
					"REDACTED", prs.NominationPOLDuration)
				return ballot
			}
		}
	}

	return nil
}

//

func (connectReader *Handler) nodeMetricsProcess() {
	for {
		if !connectReader.IsActive() {
			connectReader.Tracer.Details("REDACTED")
			return
		}

		select {
		case msg := <-connectReader.connectS.metricsMessageBuffer:
			connectReader.Tracer.Diagnose("REDACTED", "REDACTED", msg.NodeUID)

			//
			if msg.NodeUID == "REDACTED" {
				continue
			}

			//
			node := connectReader.Router.Nodes().Get(msg.NodeUID)
			if node == nil {
				connectReader.Tracer.Diagnose("REDACTED", "REDACTED", msg.NodeUID)
				continue
			}
			//
			ps, ok := node.Get(kinds.NodeStatusKey).(*NodeStatus)
			if !ok {
				panic(fmt.Sprintf("REDACTED", node))
			}
			switch tangibleMessage := msg.Msg.(type) {
			case *BallotSignal:
				if countBallots := ps.LogBallot(); countBallots%ballotsToInputToTransformSoundNode == 0 {
					connectReader.Router.StampNodeAsSound(node)
				}
			case *LedgerSegmentSignal:
				if countSegments := ps.LogLedgerSegment(); countSegments%ledgersToInputToTransformSoundNode == 0 {
					connectReader.Router.StampNodeAsSound(node)
				}
			case *NewDurationPhaseSignal:
				connectReader.Stats.NodeLevel.With("REDACTED", string(msg.NodeUID)).Set(float64(tangibleMessage.Level))
			}
		case <-connectReader.connectS.Exit():
			return

		case <-connectReader.Exit():
			return
		}
	}
}

//
//
//
func (connectReader *Handler) String() string {
	//
	return "REDACTED" //
}

//
func (connectReader *Handler) StringIndented(indent string) string {
	s := "REDACTED"
	s += indent + "REDACTED" + connectReader.connectS.StringIndented(indent+"REDACTED") + "REDACTED"
	connectReader.Router.Nodes().ForEach(func(node p2p.Node) {
		ps, ok := node.Get(kinds.NodeStatusKey).(*NodeStatus)
		if !ok {
			panic(fmt.Sprintf("REDACTED", node))
		}
		s += indent + "REDACTED" + ps.StringIndented(indent+"REDACTED") + "REDACTED"
	})
	s += indent + "REDACTED"
	return s
}

//
func HandlerStats(stats *Stats) HandlerSetting {
	return func(connectReader *Handler) { connectReader.Stats = stats }
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
	PRS   cskinds.NodeDurationStatus `json:"duration_status"` //
	Metrics *nodeStatusMetrics        `json:"metrics"`       //
}

//
type nodeStatusMetrics struct {
	Ballots      int `json:"ballots"`
	LedgerSegments int `json:"ledger_segments"`
}

func (pss nodeStatusMetrics) String() string {
	return fmt.Sprintf("REDACTED",
		pss.Ballots, pss.LedgerSegments)
}

//
func NewNodeStatus(node p2p.Node) *NodeStatus {
	return &NodeStatus{
		node:   node,
		tracer: log.NewNoopTracer(),
		PRS: cskinds.NodeDurationStatus{
			Cycle:              -1,
			NominationPOLDuration:   -1,
			FinalEndorseDuration:    -1,
			OvertakeEndorseDuration: -1,
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
func (ps *NodeStatus) FetchDurationStatus() *cskinds.NodeDurationStatus {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	prs := ps.PRS //
	return &prs
}

//
func (ps *NodeStatus) SerializeJSON() ([]byte, error) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	type jsonNodeStatus NodeStatus
	return cometjson.Serialize((*jsonNodeStatus)(ps))
}

//
//
func (ps *NodeStatus) FetchLevel() int64 {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return ps.PRS.Level
}

//
func (ps *NodeStatus) AssignHasNomination(nomination *kinds.Nomination) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.Level != nomination.Level || ps.PRS.Cycle != nomination.Cycle {
		return
	}

	if ps.PRS.Nomination {
		return
	}

	ps.PRS.Nomination = true

	//
	if ps.PRS.NominationLedgerSegments != nil {
		return
	}

	ps.PRS.NominationLedgerSegmentAssignHeading = nomination.LedgerUID.SegmentAssignHeading
	ps.PRS.NominationLedgerSegments = bits.NewBitList(int(nomination.LedgerUID.SegmentAssignHeading.Sum))
	ps.PRS.NominationPOLDuration = nomination.POLDuration
	ps.PRS.NominationPOL = nil //
}

//
func (ps *NodeStatus) InitNominationLedgerSegments(segmentAssignHeading kinds.SegmentAssignHeading) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.NominationLedgerSegments != nil {
		return
	}

	ps.PRS.NominationLedgerSegmentAssignHeading = segmentAssignHeading
	ps.PRS.NominationLedgerSegments = bits.NewBitList(int(segmentAssignHeading.Sum))
}

//
func (ps *NodeStatus) AssignHasNominationLedgerSegment(level int64, duration int32, ordinal int) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.Level != level || ps.PRS.Cycle != duration {
		return
	}

	ps.PRS.NominationLedgerSegments.AssignOrdinal(ordinal, true)
}

//
//
//
//
//
func (ps *NodeStatus) SelectTransmitBallot(ballots kinds.BallotAssignScanner) bool {
	if ballot := ps.SelectBallotToTransmit(ballots); ballot != nil {
		ps.tracer.Diagnose("REDACTED", "REDACTED", ps, "REDACTED", ballot)
		if ps.node.Transmit(p2p.Packet{
			StreamUID: BallotStream,
			Signal: &cometconnect.Ballot{
				Ballot: ballot.ToSchema(),
			},
		}) {
			ps.AssignHasBallot(ballot)
			return true
		}
		return false
	}
	return false
}

//
//
func (ps *NodeStatus) TransmitSegmentAssignHasSegment(segment *kinds.Segment, prs *cskinds.NodeDurationStatus) bool {
	//
	ps.tracer.Diagnose("REDACTED", "REDACTED", prs.Level, "REDACTED", prs.Cycle, "REDACTED", segment.Ordinal)
	pp, err := segment.ToSchema()
	if err != nil {
		//
		ps.tracer.Fault("REDACTED", "REDACTED", segment.Ordinal, "REDACTED", err)
		return false
	}
	if ps.node.Transmit(p2p.Packet{
		StreamUID: DataStream,
		Signal: &cometconnect.LedgerSegment{
			Level: prs.Level, //
			Cycle:  prs.Cycle,  //
			Segment:   *pp,
		},
	}) {
		ps.AssignHasNominationLedgerSegment(prs.Level, prs.Cycle, int(segment.Ordinal))
		return true
	}
	ps.tracer.Diagnose("REDACTED")
	return false
}

//
//
func (ps *NodeStatus) TransmitNominationAssignHasNomination(
	tracer log.Tracer,
	rs *cskinds.DurationStatus,
	prs *cskinds.NodeDurationStatus,
) {
	//
	tracer.Diagnose("REDACTED", "REDACTED", prs.Level, "REDACTED", prs.Cycle)
	if ps.node.Transmit(p2p.Packet{
		StreamUID: DataStream,
		Signal:   &cometconnect.Nomination{Nomination: *rs.Nomination.ToSchema()},
	}) {
		//
		ps.AssignHasNomination(rs.Nomination)
	}

	//
	//
	//
	//
	if 0 <= rs.Nomination.POLDuration {
		tracer.Diagnose("REDACTED", "REDACTED", prs.Level, "REDACTED", prs.Cycle)
		ps.node.Transmit(p2p.Packet{
			StreamUID: DataStream,
			Signal: &cometconnect.NominationPOL{
				Level:           rs.Level,
				NominationPolDuration: rs.Nomination.POLDuration,
				NominationPol:      *rs.Ballots.Preballots(rs.Nomination.POLDuration).BitList().ToSchema(),
			},
		})
	}
}

//
//
func (ps *NodeStatus) transmitBallotAssignHasBallot(ballot *kinds.Ballot) bool {
	ps.tracer.Diagnose("REDACTED", "REDACTED", ps, "REDACTED", ballot)
	if ps.node.Transmit(p2p.Packet{
		StreamUID: BallotStream,
		Signal: &cometconnect.Ballot{
			Ballot: ballot.ToSchema(),
		},
	}) {
		ps.AssignHasBallot(ballot)
		return true
	}
	return false
}

//
//
//
func (ps *NodeStatus) SelectBallotToTransmit(ballots kinds.BallotAssignScanner) *kinds.Ballot {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ballots.Volume() == 0 {
		return nil
	}

	level, duration, ballotsKind, volume := ballots.FetchLevel(), ballots.FetchDuration(), engineproto.AttestedMessageKind(ballots.Kind()), ballots.Volume()

	//
	if ballots.IsEndorse() {
		ps.assureOvertakeEndorseDuration(level, duration, volume)
	}
	ps.assureBallotBitLists(level, volume)

	psBallots := ps.fetchBallotBitList(level, duration, ballotsKind)
	if psBallots == nil {
		return nil //
	}
	if ordinal, ok := ballots.BitList().Sub(psBallots).SelectArbitrary(); ok {
		ballot := ballots.FetchByOrdinal(int32(ordinal))
		if ballot == nil {
			ps.tracer.Fault("REDACTED", "REDACTED", ballots, "REDACTED", ordinal)
		}
		return ballot
	}
	return nil
}

func (ps *NodeStatus) fetchBallotBitList(level int64, duration int32, ballotsKind engineproto.AttestedMessageKind) *bits.BitList {
	if !kinds.IsBallotKindSound(ballotsKind) {
		return nil
	}

	if ps.PRS.Level == level {
		if ps.PRS.Cycle == duration {
			switch ballotsKind {
			case engineproto.PreballotKind:
				return ps.PRS.Preballots
			case engineproto.PreendorseKind:
				return ps.PRS.Preendorsements
			}
		}
		if ps.PRS.OvertakeEndorseDuration == duration {
			switch ballotsKind {
			case engineproto.PreballotKind:
				return nil
			case engineproto.PreendorseKind:
				return ps.PRS.OvertakeEndorse
			}
		}
		if ps.PRS.NominationPOLDuration == duration {
			switch ballotsKind {
			case engineproto.PreballotKind:
				return ps.PRS.NominationPOL
			case engineproto.PreendorseKind:
				return nil
			}
		}
		return nil
	}
	if ps.PRS.Level == level+1 {
		if ps.PRS.FinalEndorseDuration == duration {
			switch ballotsKind {
			case engineproto.PreballotKind:
				return nil
			case engineproto.PreendorseKind:
				return ps.PRS.FinalEndorse
			}
		}
		return nil
	}
	return nil
}

//
func (ps *NodeStatus) assureOvertakeEndorseDuration(level int64, duration int32, countRatifiers int) {
	if ps.PRS.Level != level {
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
	if ps.PRS.OvertakeEndorseDuration == duration {
		return //
	}
	ps.PRS.OvertakeEndorseDuration = duration
	if duration == ps.PRS.Cycle {
		ps.PRS.OvertakeEndorse = ps.PRS.Preendorsements
	} else {
		ps.PRS.OvertakeEndorse = bits.NewBitList(countRatifiers)
	}
}

//
//
//
//
func (ps *NodeStatus) AssureBallotBitLists(level int64, countRatifiers int) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	ps.assureBallotBitLists(level, countRatifiers)
}

func (ps *NodeStatus) assureBallotBitLists(level int64, countRatifiers int) {
	switch ps.PRS.Level {
	case level:
		if ps.PRS.Preballots == nil {
			ps.PRS.Preballots = bits.NewBitList(countRatifiers)
		}
		if ps.PRS.Preendorsements == nil {
			ps.PRS.Preendorsements = bits.NewBitList(countRatifiers)
		}
		if ps.PRS.OvertakeEndorse == nil {
			ps.PRS.OvertakeEndorse = bits.NewBitList(countRatifiers)
		}
		if ps.PRS.NominationPOL == nil {
			ps.PRS.NominationPOL = bits.NewBitList(countRatifiers)
		}
	case level + 1:
		if ps.PRS.FinalEndorse == nil {
			ps.PRS.FinalEndorse = bits.NewBitList(countRatifiers)
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
func (ps *NodeStatus) LogLedgerSegment() int {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	ps.Metrics.LedgerSegments++
	return ps.Metrics.LedgerSegments
}

//
func (ps *NodeStatus) LedgerSegmentsRelayed() int {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	return ps.Metrics.LedgerSegments
}

//
func (ps *NodeStatus) AssignHasBallot(ballot *kinds.Ballot) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	ps.assignHasBallot(ballot.Level, ballot.Cycle, ballot.Kind, ballot.RatifierOrdinal)
}

func (ps *NodeStatus) assignHasBallot(level int64, duration int32, ballotKind engineproto.AttestedMessageKind, ordinal int32) {
	ps.tracer.Diagnose("REDACTED",
		"REDACTED",
		log.NewIdleFormat("REDACTED", ps.PRS.Level, ps.PRS.Cycle),
		"REDACTED",
		log.NewIdleFormat("REDACTED", level, duration),
		"REDACTED", ballotKind, "REDACTED", ordinal)

	//
	psBallots := ps.fetchBallotBitList(level, duration, ballotKind)
	if psBallots != nil {
		psBallots.AssignOrdinal(int(ordinal), true)
	}
}

//
func (ps *NodeStatus) AssignHasBallotFromNode(ballot *kinds.Ballot, csLevel int64, valueVolume, finalEndorseVolume int) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	ps.assureBallotBitLists(csLevel, valueVolume)
	ps.assureBallotBitLists(csLevel-1, finalEndorseVolume)
	ps.assignHasBallot(ballot.Level, ballot.Cycle, ballot.Kind, ballot.RatifierOrdinal)
}

//
func (ps *NodeStatus) ExecuteNewDurationPhaseSignal(msg *NewDurationPhaseSignal) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	//
	if ContrastHRS(msg.Level, msg.Cycle, msg.Phase, ps.PRS.Level, ps.PRS.Cycle, ps.PRS.Phase) <= 0 {
		return
	}

	//
	psLevel := ps.PRS.Level
	psDuration := ps.PRS.Cycle
	psOvertakeEndorseDuration := ps.PRS.OvertakeEndorseDuration
	psOvertakeEndorse := ps.PRS.OvertakeEndorse
	finalPreendorsements := ps.PRS.Preendorsements

	beginMoment := engineclock.Now().Add(-1 * time.Duration(msg.MomentsSinceBeginTime) * time.Second)
	ps.PRS.Level = msg.Level
	ps.PRS.Cycle = msg.Cycle
	ps.PRS.Phase = msg.Phase
	ps.PRS.BeginTime = beginMoment
	if psLevel != msg.Level || psDuration != msg.Cycle {
		ps.PRS.Nomination = false
		ps.PRS.NominationLedgerSegmentAssignHeading = kinds.SegmentAssignHeading{}
		ps.PRS.NominationLedgerSegments = nil
		ps.PRS.NominationPOLDuration = -1
		ps.PRS.NominationPOL = nil
		//
		ps.PRS.Preballots = nil
		ps.PRS.Preendorsements = nil
	}
	if psLevel == msg.Level && psDuration != msg.Cycle && msg.Cycle == psOvertakeEndorseDuration {
		//
		//
		//
		//
		ps.PRS.Preendorsements = psOvertakeEndorse
	}
	if psLevel != msg.Level {
		//
		if psLevel+1 == msg.Level && psDuration == msg.FinalEndorseDuration {
			ps.PRS.FinalEndorseDuration = msg.FinalEndorseDuration
			ps.PRS.FinalEndorse = finalPreendorsements
		} else {
			ps.PRS.FinalEndorseDuration = msg.FinalEndorseDuration
			ps.PRS.FinalEndorse = nil
		}
		//
		ps.PRS.OvertakeEndorseDuration = -1
		ps.PRS.OvertakeEndorse = nil
	}
}

//
func (ps *NodeStatus) ExecuteNewSoundLedgerSignal(msg *NewSoundLedgerSignal) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.Level != msg.Level {
		return
	}

	if ps.PRS.Cycle != msg.Cycle && !msg.IsEndorse {
		return
	}

	ps.PRS.NominationLedgerSegmentAssignHeading = msg.LedgerSegmentAssignHeading
	ps.PRS.NominationLedgerSegments = msg.LedgerSegments
}

//
func (ps *NodeStatus) ExecuteNominationPOLSignal(msg *NominationPOLSignal) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.Level != msg.Level {
		return
	}
	if ps.PRS.NominationPOLDuration != msg.NominationPOLDuration {
		return
	}

	//
	//
	ps.PRS.NominationPOL = msg.NominationPOL
}

//
func (ps *NodeStatus) ExecuteHasBallotSignal(msg *HasBallotSignal) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if ps.PRS.Level != msg.Level {
		return
	}

	ps.assignHasBallot(msg.Level, msg.Cycle, msg.Kind, msg.Ordinal)
}

//
//
//
//
//
func (ps *NodeStatus) ExecuteBallotAssignBitsSignal(msg *BallotAssignBitsSignal, ourBallots *bits.BitList) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	ballots := ps.fetchBallotBitList(msg.Level, msg.Cycle, msg.Kind)
	if ballots != nil {
		if ourBallots == nil {
			ballots.Modify(msg.Ballots)
		} else {
			anotherBallots := ballots.Sub(ourBallots)
			hasBallots := anotherBallots.Or(msg.Ballots)
			ballots.Modify(hasBallots)
		}
	}
}

//
func (ps *NodeStatus) String() string {
	return ps.StringIndented("REDACTED")
}

//
func (ps *NodeStatus) StringIndented(indent string) string {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTEDv
REDACTED`,
		indent, ps.node.ID(),
		indent, ps.PRS.StringIndented(indent+"REDACTED"),
		indent, ps.Metrics,
		indent)
}

//
//

//
type Signal interface {
	CertifySimple() error
}

func init() {
	cometjson.EnrollKind(&NewDurationPhaseSignal{}, "REDACTED")
	cometjson.EnrollKind(&NewSoundLedgerSignal{}, "REDACTED")
	cometjson.EnrollKind(&NominationSignal{}, "REDACTED")
	cometjson.EnrollKind(&NominationPOLSignal{}, "REDACTED")
	cometjson.EnrollKind(&LedgerSegmentSignal{}, "REDACTED")
	cometjson.EnrollKind(&BallotSignal{}, "REDACTED")
	cometjson.EnrollKind(&HasBallotSignal{}, "REDACTED")
	cometjson.EnrollKind(&BallotAssignMaj23signal{}, "REDACTED")
	cometjson.EnrollKind(&BallotAssignBitsSignal{}, "REDACTED")
}

//

//
//
type NewDurationPhaseSignal struct {
	Level                int64
	Cycle                 int32
	Phase                  cskinds.DurationPhaseKind
	MomentsSinceBeginTime int64
	FinalEndorseDuration       int32
}

//
func (m *NewDurationPhaseSignal) CertifySimple() error {
	if m.Level < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if m.Cycle < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if !m.Phase.IsSound() {
		return cometfaults.ErrCorruptField{Field: "REDACTED"}
	}

	//

	//
	//
	//
	if m.FinalEndorseDuration < -1 {
		return cometfaults.ErrCorruptField{Field: "REDACTED", Cause: "REDACTED"}
	}

	return nil
}

//
func (m *NewDurationPhaseSignal) CertifyLevel(primaryLevel int64) error {
	if m.Level < primaryLevel {
		return cometfaults.ErrCorruptField{
			Field:  "REDACTED",
			Cause: fmt.Sprintf("REDACTED", m.Level, primaryLevel),
		}
	}

	if m.Level == primaryLevel && m.FinalEndorseDuration != -1 {
		return cometfaults.ErrCorruptField{
			Field:  "REDACTED",
			Cause: fmt.Sprintf("REDACTED", m.FinalEndorseDuration, primaryLevel),
		}
	}

	if m.Level > primaryLevel && m.FinalEndorseDuration < 0 {
		return cometfaults.ErrCorruptField{
			Field:  "REDACTED",
			Cause: fmt.Sprintf("REDACTED", primaryLevel),
		}
	}
	return nil
}

//
func (m *NewDurationPhaseSignal) String() string {
	return fmt.Sprintf("REDACTED",
		m.Level, m.Cycle, m.Phase, m.FinalEndorseDuration)
}

//

//
//
//
type NewSoundLedgerSignal struct {
	Level             int64
	Cycle              int32
	LedgerSegmentAssignHeading kinds.SegmentAssignHeading
	LedgerSegments         *bits.BitList
	IsEndorse           bool
}

//
func (m *NewSoundLedgerSignal) CertifySimple() error {
	if m.Level < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if m.Cycle < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if err := m.LedgerSegmentAssignHeading.CertifySimple(); err != nil {
		return cometfaults.ErrIncorrectField{Field: "REDACTED", Err: err}
	}
	if err := m.LedgerSegments.CertifySimple(); err != nil {
		return cometfaults.ErrIncorrectField{Field: "REDACTED", Err: err}
	}
	if m.LedgerSegments.Volume() == 0 {
		return cometfaults.ErrMandatoryField{Field: "REDACTED"}
	}
	if m.LedgerSegments.Volume() != int(m.LedgerSegmentAssignHeading.Sum) {
		return fmt.Errorf("REDACTED",
			m.LedgerSegments.Volume(),
			m.LedgerSegmentAssignHeading.Sum)
	}
	if m.LedgerSegments.Volume() > int(kinds.MaximumLedgerSegmentsTally) {
		return fmt.Errorf("REDACTED", m.LedgerSegments.Volume(), kinds.MaximumLedgerSegmentsTally)
	}
	return nil
}

//
func (m *NewSoundLedgerSignal) String() string {
	return fmt.Sprintf("REDACTED",
		m.Level, m.Cycle, m.LedgerSegmentAssignHeading, m.LedgerSegments, m.IsEndorse)
}

//

//
type NominationSignal struct {
	Nomination *kinds.Nomination
}

//
func (m *NominationSignal) CertifySimple() error {
	return m.Nomination.CertifySimple()
}

//
//
func (m *NominationSignal) CertifyLedgerVolume(maximumLedgerVolumeOctets int64) error {
	return m.Nomination.CertifyLedgerVolume(maximumLedgerVolumeOctets)
}

//
func (m *NominationSignal) String() string {
	return fmt.Sprintf("REDACTED", m.Nomination)
}

//

//
type NominationPOLSignal struct {
	Level           int64
	NominationPOLDuration int32
	NominationPOL      *bits.BitList
}

//
func (m *NominationPOLSignal) CertifySimple() error {
	if m.Level < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if m.NominationPOLDuration < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if err := m.NominationPOL.CertifySimple(); err != nil {
		return cometfaults.ErrIncorrectField{Field: "REDACTED", Err: err}
	}
	if m.NominationPOL.Volume() == 0 {
		return cometfaults.ErrMandatoryField{Field: "REDACTED"}
	}
	if m.NominationPOL.Volume() > kinds.MaximumBallotsTally {
		return fmt.Errorf("REDACTED", m.NominationPOL.Volume(), kinds.MaximumBallotsTally)
	}
	return nil
}

//
func (m *NominationPOLSignal) String() string {
	return fmt.Sprintf("REDACTED", m.Level, m.NominationPOLDuration, m.NominationPOL)
}

//

//
type LedgerSegmentSignal struct {
	Level int64
	Cycle  int32
	Segment   *kinds.Segment
}

//
func (m *LedgerSegmentSignal) CertifySimple() error {
	if m.Level < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if m.Cycle < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if err := m.Segment.CertifySimple(); err != nil {
		return cometfaults.ErrIncorrectField{Field: "REDACTED", Err: err}
	}
	return nil
}

//
func (m *LedgerSegmentSignal) String() string {
	return fmt.Sprintf("REDACTED", m.Level, m.Cycle, m.Segment)
}

//

//
type BallotSignal struct {
	Ballot *kinds.Ballot
}

//
func (m *BallotSignal) CertifySimple() error {
	return m.Ballot.CertifySimple()
}

//
func (m *BallotSignal) String() string {
	return fmt.Sprintf("REDACTED", m.Ballot)
}

//

//
type HasBallotSignal struct {
	Level int64
	Cycle  int32
	Kind   engineproto.AttestedMessageKind
	Ordinal  int32
}

//
func (m *HasBallotSignal) CertifySimple() error {
	if m.Level < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if m.Cycle < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if !kinds.IsBallotKindSound(m.Kind) {
		return cometfaults.ErrCorruptField{Field: "REDACTED"}
	}
	if m.Ordinal < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	return nil
}

//
func (m *HasBallotSignal) String() string {
	return fmt.Sprintf("REDACTED", m.Ordinal, m.Level, m.Cycle, m.Kind)
}

//

//
type BallotAssignMaj23signal struct {
	Level  int64
	Cycle   int32
	Kind    engineproto.AttestedMessageKind
	LedgerUID kinds.LedgerUID
}

//
func (m *BallotAssignMaj23signal) CertifySimple() error {
	if m.Level < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if m.Cycle < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if !kinds.IsBallotKindSound(m.Kind) {
		return cometfaults.ErrCorruptField{Field: "REDACTED"}
	}
	if err := m.LedgerUID.CertifySimple(); err != nil {
		return cometfaults.ErrIncorrectField{Field: "REDACTED", Err: err}
	}
	return nil
}

//
func (m *BallotAssignMaj23signal) String() string {
	return fmt.Sprintf("REDACTED", m.Level, m.Cycle, m.Kind, m.LedgerUID)
}

//

//
type BallotAssignBitsSignal struct {
	Level  int64
	Cycle   int32
	Kind    engineproto.AttestedMessageKind
	LedgerUID kinds.LedgerUID
	Ballots   *bits.BitList
}

//
func (m *BallotAssignBitsSignal) CertifySimple() error {
	if m.Level < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if !kinds.IsBallotKindSound(m.Kind) {
		return cometfaults.ErrCorruptField{Field: "REDACTED"}
	}
	if err := m.LedgerUID.CertifySimple(); err != nil {
		return cometfaults.ErrIncorrectField{Field: "REDACTED", Err: err}
	}
	if err := m.Ballots.CertifySimple(); err != nil {
		return cometfaults.ErrIncorrectField{Field: "REDACTED", Err: err}
	}
	//
	if m.Ballots.Volume() > kinds.MaximumBallotsTally {
		return fmt.Errorf("REDACTED", m.Ballots.Volume(), kinds.MaximumBallotsTally)
	}
	return nil
}

//
func (m *BallotAssignBitsSignal) String() string {
	return fmt.Sprintf("REDACTED", m.Level, m.Cycle, m.Kind, m.LedgerUID, m.Ballots)
}

//

//
type HasNominationLedgerSegmentSignal struct {
	Level int64
	Cycle  int32
	Ordinal  int32
}

//
func (m *HasNominationLedgerSegmentSignal) CertifySimple() error {
	if m.Level < 1 {
		return cometfaults.ErrCorruptField{Field: "REDACTED", Cause: "REDACTED"}
	}
	if m.Cycle < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	if m.Ordinal < 0 {
		return cometfaults.ErrAdverseField{Field: "REDACTED"}
	}
	return nil
}

//
func (m *HasNominationLedgerSegmentSignal) String() string {
	return fmt.Sprintf("REDACTED", m.Ordinal, m.Level, m.Cycle)
}
