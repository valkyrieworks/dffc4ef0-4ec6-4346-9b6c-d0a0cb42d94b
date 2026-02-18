package agreement

import (
	"fmt"

	cometfaults "github.com/valkyrieworks/kinds/faults"
	"github.com/cosmos/gogoproto/proto"

	cskinds "github.com/valkyrieworks/agreement/kinds"
	"github.com/valkyrieworks/utils/bits"
	cometmath "github.com/valkyrieworks/utils/math"
	"github.com/valkyrieworks/p2p"
	cometconnect "github.com/valkyrieworks/schema/consensuscore/agreement"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//
//
func MessageToSchema(msg Signal) (proto.Message, error) {
	if msg == nil {
		return nil, ErrNullSignal
	}
	var pb proto.Message

	switch msg := msg.(type) {
	case *NewDurationPhaseSignal:
		pb = &cometconnect.NewDurationPhase{
			Level:                msg.Level,
			Cycle:                 msg.Cycle,
			Phase:                  uint32(msg.Phase),
			MomentsSinceBeginTime: msg.MomentsSinceBeginTime,
			FinalEndorseDuration:       msg.FinalEndorseDuration,
		}

	case *NewSoundLedgerSignal:
		pbSegmentAssignHeading := msg.LedgerSegmentAssignHeading.ToSchema()
		pbBits := msg.LedgerSegments.ToSchema()
		pb = &cometconnect.NewSoundLedger{
			Level:             msg.Level,
			Cycle:              msg.Cycle,
			LedgerSegmentAssignHeading: pbSegmentAssignHeading,
			LedgerSegments:         pbBits,
			IsEndorse:           msg.IsEndorse,
		}

	case *NominationSignal:
		pbP := msg.Nomination.ToSchema()
		pb = &cometconnect.Nomination{
			Nomination: *pbP,
		}

	case *NominationPOLSignal:
		pbBits := msg.NominationPOL.ToSchema()
		pb = &cometconnect.NominationPOL{
			Level:           msg.Level,
			NominationPolDuration: msg.NominationPOLDuration,
			NominationPol:      *pbBits,
		}

	case *LedgerSegmentSignal:
		segments, err := msg.Segment.ToSchema()
		if err != nil {
			return nil, cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err}
		}
		pb = &cometconnect.LedgerSegment{
			Level: msg.Level,
			Cycle:  msg.Cycle,
			Segment:   *segments,
		}

	case *BallotSignal:
		ballot := msg.Ballot.ToSchema()
		pb = &cometconnect.Ballot{
			Ballot: ballot,
		}

	case *HasBallotSignal:
		pb = &cometconnect.HasBallot{
			Level: msg.Level,
			Cycle:  msg.Cycle,
			Kind:   msg.Kind,
			Ordinal:  msg.Ordinal,
		}

	case *BallotAssignMaj23signal:
		bi := msg.LedgerUID.ToSchema()
		pb = &cometconnect.BallotAssignMaj23{
			Level:  msg.Level,
			Cycle:   msg.Cycle,
			Kind:    msg.Kind,
			LedgerUID: bi,
		}

	case *BallotAssignBitsSignal:
		bi := msg.LedgerUID.ToSchema()
		bits := msg.Ballots.ToSchema()

		vsb := &cometconnect.BallotAssignBits{
			Level:  msg.Level,
			Cycle:   msg.Cycle,
			Kind:    msg.Kind,
			LedgerUID: bi,
		}

		if bits != nil {
			vsb.Ballots = *bits
		}

		pb = vsb

	default:
		return nil, ErrAgreementSignalNotIdentified{msg}
	}

	return pb, nil
}

//
func MessageFromSchema(p proto.Message) (Signal, error) {
	if p == nil {
		return nil, ErrNullSignal
	}
	var pb Signal

	switch msg := p.(type) {
	case *cometconnect.NewDurationPhase:
		rs, err := cometmath.SecureTransformUint8(int64(msg.Phase))
		//
		if err != nil {
			return nil, ErrRefuseSignalOverload{err}
		}
		pb = &NewDurationPhaseSignal{
			Level:                msg.Level,
			Cycle:                 msg.Cycle,
			Phase:                  cskinds.DurationPhaseKind(rs),
			MomentsSinceBeginTime: msg.MomentsSinceBeginTime,
			FinalEndorseDuration:       msg.FinalEndorseDuration,
		}
	case *cometconnect.NewSoundLedger:
		pbSegmentAssignHeading, err := kinds.SegmentAssignHeadingFromSchema(&msg.LedgerSegmentAssignHeading)
		if err != nil {
			return nil, cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err}
		}

		pbBits := new(bits.BitList)
		pbBits.FromSchema(msg.LedgerSegments)

		pb = &NewSoundLedgerSignal{
			Level:             msg.Level,
			Cycle:              msg.Cycle,
			LedgerSegmentAssignHeading: *pbSegmentAssignHeading,
			LedgerSegments:         pbBits,
			IsEndorse:           msg.IsEndorse,
		}
	case *cometconnect.Nomination:
		pbP, err := kinds.NominationFromSchema(&msg.Nomination)
		if err != nil {
			return nil, cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err}
		}

		pb = &NominationSignal{
			Nomination: pbP,
		}
	case *cometconnect.NominationPOL:
		pbBits := new(bits.BitList)
		pbBits.FromSchema(&msg.NominationPol)
		pb = &NominationPOLSignal{
			Level:           msg.Level,
			NominationPOLDuration: msg.NominationPolDuration,
			NominationPOL:      pbBits,
		}
	case *cometconnect.LedgerSegment:
		segments, err := kinds.SegmentFromSchema(&msg.Segment)
		if err != nil {
			return nil, cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err}
		}
		pb = &LedgerSegmentSignal{
			Level: msg.Level,
			Cycle:  msg.Cycle,
			Segment:   segments,
		}
	case *cometconnect.Ballot:
		//
		//
		ballot, err := kinds.BallotFromSchema(msg.Ballot)
		if err != nil {
			return nil, cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err}
		}

		pb = &BallotSignal{
			Ballot: ballot,
		}
	case *cometconnect.HasBallot:
		pb = &HasBallotSignal{
			Level: msg.Level,
			Cycle:  msg.Cycle,
			Kind:   msg.Kind,
			Ordinal:  msg.Ordinal,
		}
	case *cometconnect.BallotAssignMaj23:
		bi, err := kinds.LedgerUIDFromSchema(&msg.LedgerUID)
		if err != nil {
			return nil, cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err}
		}
		pb = &BallotAssignMaj23signal{
			Level:  msg.Level,
			Cycle:   msg.Cycle,
			Kind:    msg.Kind,
			LedgerUID: *bi,
		}
	case *cometconnect.BallotAssignBits:
		bi, err := kinds.LedgerUIDFromSchema(&msg.LedgerUID)
		if err != nil {
			return nil, cometfaults.ErrMessageToSchema{SignalLabel: "REDACTED", Err: err}
		}
		bits := new(bits.BitList)
		bits.FromSchema(&msg.Ballots)

		pb = &BallotAssignBitsSignal{
			Level:  msg.Level,
			Cycle:   msg.Cycle,
			Kind:    msg.Kind,
			LedgerUID: *bi,
			Ballots:   bits,
		}
	default:
		return nil, ErrAgreementSignalNotIdentified{msg}
	}

	if err := pb.CertifySimple(); err != nil {
		return nil, err
	}

	return pb, nil
}

//
func JournalToSchema(msg JournalSignal) (*cometconnect.JournalSignal, error) {
	var pb cometconnect.JournalSignal

	switch msg := msg.(type) {
	case kinds.EventDataDurationStatus:
		pb = cometconnect.JournalSignal{
			Sum: &cometconnect.Journalsignal_Signaldatadurationstate{
				EventDataDurationStatus: &engineproto.EventDataDurationStatus{
					Level: msg.Level,
					Cycle:  msg.Cycle,
					Phase:   msg.Phase,
				},
			},
		}
	case messageDetails:
		constMessage, err := MessageToSchema(msg.Msg)
		if err != nil {
			return nil, err
		}
		if w, ok := constMessage.(p2p.Adapter); ok {
			constMessage = w.Enclose()
		}
		cm := constMessage.(*cometconnect.Signal)
		pb = cometconnect.JournalSignal{
			Sum: &cometconnect.Journalsignal_Signaldetails{
				MessageDetails: &cometconnect.MessageDetails{
					Msg:    *cm,
					NodeUID: string(msg.NodeUID),
				},
			},
		}
	case deadlineDetails:
		pb = cometconnect.JournalSignal{
			Sum: &cometconnect.Journalsignal_Deadlinedetails{
				DeadlineDetails: &cometconnect.DeadlineDetails{
					Period: msg.Period,
					Level:   msg.Level,
					Cycle:    msg.Cycle,
					Phase:     uint32(msg.Phase),
				},
			},
		}
	case TerminateLevelSignal:
		pb = cometconnect.JournalSignal{
			Sum: &cometconnect.Journalsignal_Finallayer{
				TerminateLevel: &cometconnect.TerminateLevel{
					Level: msg.Level,
				},
			},
		}
	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}

	return &pb, nil
}

//
func JournalFromSchema(msg *cometconnect.JournalSignal) (JournalSignal, error) {
	if msg == nil {
		return nil, ErrNullSignal
	}
	var pb JournalSignal

	switch msg := msg.Sum.(type) {
	case *cometconnect.Journalsignal_Signaldatadurationstate:
		pb = kinds.EventDataDurationStatus{
			Level: msg.EventDataDurationStatus.Level,
			Cycle:  msg.EventDataDurationStatus.Cycle,
			Phase:   msg.EventDataDurationStatus.Phase,
		}
	case *cometconnect.Journalsignal_Signaldetails:
		um, err := msg.MessageDetails.Msg.Disclose()
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		journalMessage, err := MessageFromSchema(um)
		if err != nil {
			return nil, cometfaults.ErrMessageFromSchema{SignalLabel: "REDACTED", Err: err}
		}
		pb = messageDetails{
			Msg:    journalMessage,
			NodeUID: p2p.ID(msg.MessageDetails.NodeUID),
		}

	case *cometconnect.Journalsignal_Deadlinedetails:
		tis, err := cometmath.SecureTransformUint8(int64(msg.DeadlineDetails.Phase))
		//
		if err != nil {
			return nil, ErrRefuseSignalOverload{err}
		}
		pb = deadlineDetails{
			Period: msg.DeadlineDetails.Period,
			Level:   msg.DeadlineDetails.Level,
			Cycle:    msg.DeadlineDetails.Cycle,
			Phase:     cskinds.DurationPhaseKind(tis),
		}
		return pb, nil
	case *cometconnect.Journalsignal_Finallayer:
		pb := TerminateLevelSignal{
			Level: msg.TerminateLevel.Level,
		}
		return pb, nil
	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
	return pb, nil
}
