package kinds

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//

//
type IterationPhaseKind uint8 //

//
const (
	IterationPhaseFreshAltitude     = IterationPhaseKind(0x01) //
	IterationPhaseFreshIteration      = IterationPhaseKind(0x02) //
	IterationPhaseNominate       = IterationPhaseKind(0x03) //
	IterationPhasePreballot       = IterationPhaseKind(0x04) //
	IterationPhasePreballotAwait   = IterationPhaseKind(0x05) //
	IterationPhasePreendorse     = IterationPhaseKind(0x06) //
	IterationPhasePreendorseAwait = IterationPhaseKind(0x07) //
	IterationPhaseEndorse        = IterationPhaseKind(0x08) //
	//

	//
)

//
func (rs IterationPhaseKind) EqualsSound() bool {
	return uint8(rs) >= 0x01 && uint8(rs) <= 0x08
}

//
func (rs IterationPhaseKind) Text() string {
	switch rs {
	case IterationPhaseFreshAltitude:
		return "REDACTED"
	case IterationPhaseFreshIteration:
		return "REDACTED"
	case IterationPhaseNominate:
		return "REDACTED"
	case IterationPhasePreballot:
		return "REDACTED"
	case IterationPhasePreballotAwait:
		return "REDACTED"
	case IterationPhasePreendorse:
		return "REDACTED"
	case IterationPhasePreendorseAwait:
		return "REDACTED"
	case IterationPhaseEndorse:
		return "REDACTED"
	default:
		return "REDACTED" //
	}
}

//

//
//
//
type IterationStatus struct {
	Altitude    int64         `json:"altitude"` //
	Iteration     int32         `json:"iteration"`
	Phase      IterationPhaseKind `json:"phase"`
	InitiateMoment time.Time     `json:"initiate_moment"`

	//
	EndorseMoment         time.Time           `json:"endorse_moment"`
	Assessors         *kinds.AssessorAssign `json:"assessors"`
	Nomination           *kinds.Nomination     `json:"nomination"`
	NominationLedger      *kinds.Ledger        `json:"nomination_ledger"`
	NominationLedgerFragments *kinds.FragmentAssign      `json:"nomination_ledger_fragments"`
	SecuredIteration        int32               `json:"secured_iteration"`
	SecuredLedger        *kinds.Ledger        `json:"secured_ledger"`
	SecuredLedgerFragments   *kinds.FragmentAssign      `json:"secured_ledger_fragments"`

	//
	//
	//
	//
	//
	//
	//
	//

	//
	SoundIteration int32        `json:"sound_iteration"`
	SoundLedger *kinds.Ledger `json:"sound_ledger"` //

	//
	SoundLedgerFragments           *kinds.FragmentAssign      `json:"sound_ledger_fragments"`
	Ballots                     *AltitudeBallotAssign      `json:"ballots"`
	EndorseIteration               int32               `json:"endorse_iteration"` //
	FinalEndorse                *kinds.BallotAssign      `json:"final_endorse"`  //
	FinalAssessors            *kinds.AssessorAssign `json:"final_assessors"`
	ActivatedDeadlinePreendorse bool                `json:"activated_deadline_preendorse"`
}

//
type IterationStatusPlain struct {
	AltitudeIterationPhase   string              `json:"altitude/cycle/phase"`
	InitiateMoment         time.Time           `json:"initiate_moment"`
	NominationLedgerDigest octets.HexadecimalOctets      `json:"nomination_ledger_digest"`
	SecuredLedgerDigest   octets.HexadecimalOctets      `json:"secured_ledger_digest"`
	SoundLedgerDigest    octets.HexadecimalOctets      `json:"sound_ledger_digest"`
	Ballots             json.RawMessage     `json:"altitude_ballot_assign"`
	Nominator          kinds.AssessorDetails `json:"nominator"`
}

//
func (rs *IterationStatus) IterationStatusPlain() IterationStatusPlain {
	ballotsJSN, err := rs.Ballots.SerializeJSN()
	if err != nil {
		panic(err)
	}

	location := rs.Assessors.ObtainNominator().Location
	idx, _ := rs.Assessors.ObtainViaLocatorMutable(location)

	return IterationStatusPlain{
		AltitudeIterationPhase:   fmt.Sprintf("REDACTED", rs.Altitude, rs.Iteration, rs.Phase),
		InitiateMoment:         rs.InitiateMoment,
		NominationLedgerDigest: rs.NominationLedger.Digest(),
		SecuredLedgerDigest:   rs.SecuredLedger.Digest(),
		SoundLedgerDigest:    rs.SoundLedger.Digest(),
		Ballots:             ballotsJSN,
		Nominator: kinds.AssessorDetails{
			Location: location,
			Ordinal:   idx,
		},
	}
}

//
func (rs *IterationStatus) FreshIterationIncident() kinds.IncidentDataFreshIteration {
	location := rs.Assessors.ObtainNominator().Location
	idx, _ := rs.Assessors.ObtainViaLocatorMutable(location)

	return kinds.IncidentDataFreshIteration{
		Altitude: rs.Altitude,
		Iteration:  rs.Iteration,
		Phase:   rs.Phase.Text(),
		Nominator: kinds.AssessorDetails{
			Location: location,
			Ordinal:   idx,
		},
	}
}

//
func (rs *IterationStatus) FinishedNominationIncident() kinds.IncidentDataFinishedNomination {
	//
	//
	ledgerUUID := kinds.LedgerUUID{
		Digest:          rs.NominationLedger.Digest(),
		FragmentAssignHeading: rs.NominationLedgerFragments.Heading(),
	}

	return kinds.IncidentDataFinishedNomination{
		Altitude:  rs.Altitude,
		Iteration:   rs.Iteration,
		Phase:    rs.Phase.Text(),
		LedgerUUID: ledgerUUID,
	}
}

//
func (rs *IterationStatus) IterationStatusIncident() kinds.IncidentDataIterationStatus {
	return kinds.IncidentDataIterationStatus{
		Altitude: rs.Altitude,
		Iteration:  rs.Iteration,
		Phase:   rs.Phase.Text(),
	}
}

//
func (rs *IterationStatus) Text() string {
	return rs.TextFormatted("REDACTED")
}

//
func (rs *IterationStatus) TextFormatted(format string) string {
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTEDv
REDACTED`,
		format, rs.Altitude, rs.Iteration, rs.Phase,
		format, rs.InitiateMoment,
		format, rs.EndorseMoment,
		format, rs.Assessors.TextFormatted(format+"REDACTED"),
		format, rs.Nomination,
		format, rs.NominationLedgerFragments.TextBrief(), rs.NominationLedger.TextBrief(),
		format, rs.SecuredIteration,
		format, rs.SecuredLedgerFragments.TextBrief(), rs.SecuredLedger.TextBrief(),
		format, rs.SoundIteration,
		format, rs.SoundLedgerFragments.TextBrief(), rs.SoundLedger.TextBrief(),
		format, rs.Ballots.TextFormatted(format+"REDACTED"),
		format, rs.FinalEndorse.TextBrief(),
		format, rs.FinalAssessors.TextFormatted(format+"REDACTED"),
		format)
}

//
func (rs *IterationStatus) TextBrief() string {
	return fmt.Sprintf("REDACTED",
		rs.Altitude, rs.Iteration, rs.Phase, rs.InitiateMoment)
}
