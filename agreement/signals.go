package agreement

import (
	"fmt"

	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
	"github.com/cosmos/gogoproto/proto"

	controlkinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/digits"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	strongmindcons "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/agreement"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
func SignalTowardSchema(msg Signal) (proto.Message, error) {
	if msg == nil {
		return nil, FaultVoidSignal
	}
	var pb proto.Message

	switch msg := msg.(type) {
	case *FreshIterationPhaseSignal:
		pb = &strongmindcons.FreshIterationPhase{
			Altitude:                msg.Altitude,
			Iteration:                 msg.Iteration,
			Phase:                  uint32(msg.Phase),
			MomentsBecauseInitiateMoment: msg.MomentsBecauseInitiateMoment,
			FinalEndorseIteration:       msg.FinalEndorseIteration,
		}

	case *FreshSoundLedgerSignal:
		bufferFragmentAssignHeading := msg.LedgerFragmentAssignHeading.TowardSchema()
		bufferDigits := msg.LedgerFragments.TowardSchema()
		pb = &strongmindcons.FreshSoundLedger{
			Altitude:             msg.Altitude,
			Iteration:              msg.Iteration,
			LedgerFragmentAssignHeading: bufferFragmentAssignHeading,
			LedgerFragments:         bufferDigits,
			EqualsEndorse:           msg.EqualsEndorse,
		}

	case *NominationSignal:
		pbP := msg.Nomination.TowardSchema()
		pb = &strongmindcons.Nomination{
			Nomination: *pbP,
		}

	case *NominationPolicySignal:
		bufferDigits := msg.NominationPolicy.TowardSchema()
		pb = &strongmindcons.NominationPolicy{
			Altitude:           msg.Altitude,
			NominationPolicyIteration: msg.NominationPolicyIteration,
			NominationPolicy:      *bufferDigits,
		}

	case *LedgerFragmentSignal:
		fragments, err := msg.Fragment.TowardSchema()
		if err != nil {
			return nil, strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err}
		}
		pb = &strongmindcons.LedgerFragment{
			Altitude: msg.Altitude,
			Iteration:  msg.Iteration,
			Fragment:   *fragments,
		}

	case *BallotSignal:
		ballot := msg.Ballot.TowardSchema()
		pb = &strongmindcons.Ballot{
			Ballot: ballot,
		}

	case *OwnsBallotSignal:
		pb = &strongmindcons.OwnsBallot{
			Altitude: msg.Altitude,
			Iteration:  msg.Iteration,
			Kind:   msg.Kind,
			Ordinal:  msg.Ordinal,
		}

	case *BallotAssignMajor23signal:
		bi := msg.LedgerUUID.TowardSchema()
		pb = &strongmindcons.BallotAssignMajor23{
			Altitude:  msg.Altitude,
			Iteration:   msg.Iteration,
			Kind:    msg.Kind,
			LedgerUUID: bi,
		}

	case *BallotAssignDigitsSignal:
		bi := msg.LedgerUUID.TowardSchema()
		digits := msg.Ballots.TowardSchema()

		vsb := &strongmindcons.BallotAssignDigits{
			Altitude:  msg.Altitude,
			Iteration:   msg.Iteration,
			Kind:    msg.Kind,
			LedgerUUID: bi,
		}

		if digits != nil {
			vsb.Ballots = *digits
		}

		pb = vsb

	default:
		return nil, FaultAgreementSignalNegationIdentified{msg}
	}

	return pb, nil
}

//
func SignalOriginatingSchema(p proto.Message) (Signal, error) {
	if p == nil {
		return nil, FaultVoidSignal
	}
	var pb Signal

	switch msg := p.(type) {
	case *strongmindcons.FreshIterationPhase:
		rs, err := strongarithmetic.SecureTransformOctet(int64(msg.Phase))
		//
		if err != nil {
			return nil, FaultRefuseSignalOverrun{err}
		}
		pb = &FreshIterationPhaseSignal{
			Altitude:                msg.Altitude,
			Iteration:                 msg.Iteration,
			Phase:                  controlkinds.IterationPhaseKind(rs),
			MomentsBecauseInitiateMoment: msg.MomentsBecauseInitiateMoment,
			FinalEndorseIteration:       msg.FinalEndorseIteration,
		}
	case *strongmindcons.FreshSoundLedger:
		bufferFragmentAssignHeading, err := kinds.FragmentAssignHeadingOriginatingSchema(&msg.LedgerFragmentAssignHeading)
		if err != nil {
			return nil, strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err}
		}

		bufferDigits := new(digits.DigitSeries)
		bufferDigits.OriginatingSchema(msg.LedgerFragments)

		pb = &FreshSoundLedgerSignal{
			Altitude:             msg.Altitude,
			Iteration:              msg.Iteration,
			LedgerFragmentAssignHeading: *bufferFragmentAssignHeading,
			LedgerFragments:         bufferDigits,
			EqualsEndorse:           msg.EqualsEndorse,
		}
	case *strongmindcons.Nomination:
		pbP, err := kinds.NominationOriginatingSchema(&msg.Nomination)
		if err != nil {
			return nil, strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err}
		}

		pb = &NominationSignal{
			Nomination: pbP,
		}
	case *strongmindcons.NominationPolicy:
		bufferDigits := new(digits.DigitSeries)
		bufferDigits.OriginatingSchema(&msg.NominationPolicy)
		pb = &NominationPolicySignal{
			Altitude:           msg.Altitude,
			NominationPolicyIteration: msg.NominationPolicyIteration,
			NominationPolicy:      bufferDigits,
		}
	case *strongmindcons.LedgerFragment:
		fragments, err := kinds.FragmentOriginatingSchema(&msg.Fragment)
		if err != nil {
			return nil, strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err}
		}
		pb = &LedgerFragmentSignal{
			Altitude: msg.Altitude,
			Iteration:  msg.Iteration,
			Fragment:   fragments,
		}
	case *strongmindcons.Ballot:
		//
		//
		ballot, err := kinds.BallotOriginatingSchema(msg.Ballot)
		if err != nil {
			return nil, strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err}
		}

		pb = &BallotSignal{
			Ballot: ballot,
		}
	case *strongmindcons.OwnsBallot:
		pb = &OwnsBallotSignal{
			Altitude: msg.Altitude,
			Iteration:  msg.Iteration,
			Kind:   msg.Kind,
			Ordinal:  msg.Ordinal,
		}
	case *strongmindcons.BallotAssignMajor23:
		bi, err := kinds.LedgerUUIDOriginatingSchema(&msg.LedgerUUID)
		if err != nil {
			return nil, strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err}
		}
		pb = &BallotAssignMajor23signal{
			Altitude:  msg.Altitude,
			Iteration:   msg.Iteration,
			Kind:    msg.Kind,
			LedgerUUID: *bi,
		}
	case *strongmindcons.BallotAssignDigits:
		bi, err := kinds.LedgerUUIDOriginatingSchema(&msg.LedgerUUID)
		if err != nil {
			return nil, strongminderrors.FaultSignalTowardSchema{SignalAlias: "REDACTED", Err: err}
		}
		digits := new(digits.DigitSeries)
		digits.OriginatingSchema(&msg.Ballots)

		pb = &BallotAssignDigitsSignal{
			Altitude:  msg.Altitude,
			Iteration:   msg.Iteration,
			Kind:    msg.Kind,
			LedgerUUID: *bi,
			Ballots:   digits,
		}
	default:
		return nil, FaultAgreementSignalNegationIdentified{msg}
	}

	if err := pb.CertifyFundamental(); err != nil {
		return nil, err
	}

	return pb, nil
}

//
func JournalTowardSchema(msg JournalSignal) (*strongmindcons.JournalSignal, error) {
	var pb strongmindcons.JournalSignal

	switch msg := msg.(type) {
	case kinds.IncidentDataIterationStatus:
		pb = strongmindcons.JournalSignal{
			Sum: &strongmindcons.Walrecord_Incidentiterationstate{
				IncidentDataIterationStatus: &commitchema.IncidentDataIterationStatus{
					Altitude: msg.Altitude,
					Iteration:  msg.Iteration,
					Phase:   msg.Phase,
				},
			},
		}
	case signalDetails:
		consensusSignal, err := SignalTowardSchema(msg.Msg)
		if err != nil {
			return nil, err
		}
		if w, ok := consensusSignal.(p2p.Encapsulator); ok {
			consensusSignal = w.Enclose()
		}
		cm := consensusSignal.(*strongmindcons.Signal)
		pb = strongmindcons.JournalSignal{
			Sum: &strongmindcons.Walrecord_Signalinfo{
				SignalDetails: &strongmindcons.SignalDetails{
					Msg:    *cm,
					NodeUUID: string(msg.NodeUUID),
				},
			},
		}
	case deadlineDetails:
		pb = strongmindcons.JournalSignal{
			Sum: &strongmindcons.Walrecord_Alarminfo{
				DeadlineDetails: &strongmindcons.DeadlineDetails{
					Interval: msg.Interval,
					Altitude:   msg.Altitude,
					Iteration:    msg.Iteration,
					Phase:     uint32(msg.Phase),
				},
			},
		}
	case TerminateAltitudeSignal:
		pb = strongmindcons.JournalSignal{
			Sum: &strongmindcons.Walrecord_Finalheight{
				TerminateAltitude: &strongmindcons.TerminateAltitude{
					Altitude: msg.Altitude,
				},
			},
		}
	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}

	return &pb, nil
}

//
func JournalOriginatingSchema(msg *strongmindcons.JournalSignal) (JournalSignal, error) {
	if msg == nil {
		return nil, FaultVoidSignal
	}
	var pb JournalSignal

	switch msg := msg.Sum.(type) {
	case *strongmindcons.Walrecord_Incidentiterationstate:
		pb = kinds.IncidentDataIterationStatus{
			Altitude: msg.IncidentDataIterationStatus.Altitude,
			Iteration:  msg.IncidentDataIterationStatus.Iteration,
			Phase:   msg.IncidentDataIterationStatus.Phase,
		}
	case *strongmindcons.Walrecord_Signalinfo:
		um, err := msg.SignalDetails.Msg.Disclose()
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		journalSignal, err := SignalOriginatingSchema(um)
		if err != nil {
			return nil, strongminderrors.FaultSignalOriginatingSchema{SignalAlias: "REDACTED", Err: err}
		}
		pb = signalDetails{
			Msg:    journalSignal,
			NodeUUID: p2p.ID(msg.SignalDetails.NodeUUID),
		}

	case *strongmindcons.Walrecord_Alarminfo:
		tis, err := strongarithmetic.SecureTransformOctet(int64(msg.DeadlineDetails.Phase))
		//
		if err != nil {
			return nil, FaultRefuseSignalOverrun{err}
		}
		pb = deadlineDetails{
			Interval: msg.DeadlineDetails.Interval,
			Altitude:   msg.DeadlineDetails.Altitude,
			Iteration:    msg.DeadlineDetails.Iteration,
			Phase:     controlkinds.IterationPhaseKind(tis),
		}
		return pb, nil
	case *strongmindcons.Walrecord_Finalheight:
		pb := TerminateAltitudeSignal{
			Altitude: msg.TerminateAltitude.Altitude,
		}
		return pb, nil
	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
	return pb, nil
}
