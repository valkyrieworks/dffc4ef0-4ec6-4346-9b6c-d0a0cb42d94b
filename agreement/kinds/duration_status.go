package kinds

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/kinds"
)

//
//

//
type DurationPhaseKind uint8 //

//
const (
	DurationPhaseNewLevel     = DurationPhaseKind(0x01) //
	EpochPhaseNewEpoch      = DurationPhaseKind(0x02) //
	DurationPhaseNominate       = DurationPhaseKind(0x03) //
	EpochPhasePreballot       = DurationPhaseKind(0x04) //
	DurationPhasePreballotWait   = DurationPhaseKind(0x05) //
	EpochPhasePreendorse     = DurationPhaseKind(0x06) //
	DurationPhasePreendorseWait = DurationPhaseKind(0x07) //
	DurationPhaseEndorse        = DurationPhaseKind(0x08) //
	//

	//
)

//
func (rs DurationPhaseKind) IsSound() bool {
	return uint8(rs) >= 0x01 && uint8(rs) <= 0x08
}

//
func (rs DurationPhaseKind) String() string {
	switch rs {
	case DurationPhaseNewLevel:
		return "REDACTED"
	case EpochPhaseNewEpoch:
		return "REDACTED"
	case DurationPhaseNominate:
		return "REDACTED"
	case EpochPhasePreballot:
		return "REDACTED"
	case DurationPhasePreballotWait:
		return "REDACTED"
	case EpochPhasePreendorse:
		return "REDACTED"
	case DurationPhasePreendorseWait:
		return "REDACTED"
	case DurationPhaseEndorse:
		return "REDACTED"
	default:
		return "REDACTED" //
	}
}

//

//
//
//
type EpochStatus struct {
	Level    int64         `json:"level"` //
	Cycle     int32         `json:"epoch"`
	Phase      DurationPhaseKind `json:"phase"`
	BeginTime time.Time     `json:"begin_time"`

	//
	EndorseTime         time.Time           `json:"endorse_time"`
	Ratifiers         *kinds.RatifierAssign `json:"ratifiers"`
	Nomination           *kinds.Nomination     `json:"nomination"`
	NominationLedger      *kinds.Ledger        `json:"nomination_ledger"`
	NominationLedgerSegments *kinds.SegmentCollection      `json:"nomination_ledger_segments"`
	LatchedEpoch        int32               `json:"latched_epoch"`
	LatchedLedger        *kinds.Ledger        `json:"latched_ledger"`
	LatchedLedgerSegments   *kinds.SegmentCollection      `json:"latched_ledger_segments"`

	//
	//
	//
	//
	//
	//
	//
	//

	//
	SoundEpoch int32        `json:"sound_epoch"`
	SoundLedger *kinds.Ledger `json:"sound_ledger"` //

	//
	SoundLedgerSegments           *kinds.SegmentCollection      `json:"sound_ledger_segments"`
	Ballots                     *LevelBallotCollection      `json:"ballots"`
	EndorseEpoch               int32               `json:"endorse_epoch"` //
	FinalEndorse                *kinds.BallotCollection      `json:"final_endorse"`  //
	FinalRatifiers            *kinds.RatifierAssign `json:"final_ratifiers"`
	ActivatedDeadlinePreendorse bool                `json:"activated_deadline_preendorse"`
}

//
type EpochStatusBasic struct {
	LevelEpochPhase   string              `json:"height/epoch/phase"`
	BeginTime         time.Time           `json:"begin_time"`
	NominationLedgerDigest octets.HexOctets      `json:"nomination_ledger_digest"`
	LatchedLedgerDigest   octets.HexOctets      `json:"latched_ledger_digest"`
	SoundLedgerDigest    octets.HexOctets      `json:"sound_ledger_digest"`
	Ballots             json.RawMessage     `json:"level_ballot_assign"`
	Recommender          kinds.RatifierDetails `json:"recommender"`
}

//
func (rs *EpochStatus) EpochStatusBasic() EpochStatusBasic {
	ballotsJSON, err := rs.Ballots.SerializeJSON()
	if err != nil {
		panic(err)
	}

	address := rs.Ratifiers.FetchRecommender().Location
	idx, _ := rs.Ratifiers.FetchByLocationMut(address)

	return EpochStatusBasic{
		LevelEpochPhase:   fmt.Sprintf("REDACTED", rs.Level, rs.Cycle, rs.Phase),
		BeginTime:         rs.BeginTime,
		NominationLedgerDigest: rs.NominationLedger.Digest(),
		LatchedLedgerDigest:   rs.LatchedLedger.Digest(),
		SoundLedgerDigest:    rs.SoundLedger.Digest(),
		Ballots:             ballotsJSON,
		Recommender: kinds.RatifierDetails{
			Location: address,
			Ordinal:   idx,
		},
	}
}

//
func (rs *EpochStatus) NewEpochEvent() kinds.EventDataNewEpoch {
	address := rs.Ratifiers.FetchRecommender().Location
	idx, _ := rs.Ratifiers.FetchByLocationMut(address)

	return kinds.EventDataNewEpoch{
		Level: rs.Level,
		Cycle:  rs.Cycle,
		Phase:   rs.Phase.String(),
		Recommender: kinds.RatifierDetails{
			Location: address,
			Ordinal:   idx,
		},
	}
}

//
func (rs *EpochStatus) FinishedNominationEvent() kinds.EventDataFinishedNomination {
	//
	//
	ledgerUID := kinds.LedgerUID{
		Digest:          rs.NominationLedger.Digest(),
		SegmentAssignHeading: rs.NominationLedgerSegments.Heading(),
	}

	return kinds.EventDataFinishedNomination{
		Level:  rs.Level,
		Cycle:   rs.Cycle,
		Phase:    rs.Phase.String(),
		LedgerUID: ledgerUID,
	}
}

//
func (rs *EpochStatus) EpochStatusEvent() kinds.EventDataDurationStatus {
	return kinds.EventDataDurationStatus{
		Level: rs.Level,
		Cycle:  rs.Cycle,
		Phase:   rs.Phase.String(),
	}
}

//
func (rs *EpochStatus) String() string {
	return rs.StringIndented("REDACTED")
}

//
func (rs *EpochStatus) StringIndented(indent string) string {
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
		indent, rs.Level, rs.Cycle, rs.Phase,
		indent, rs.BeginTime,
		indent, rs.EndorseTime,
		indent, rs.Ratifiers.StringIndented(indent+"REDACTED"),
		indent, rs.Nomination,
		indent, rs.NominationLedgerSegments.StringBrief(), rs.NominationLedger.StringBrief(),
		indent, rs.LatchedEpoch,
		indent, rs.LatchedLedgerSegments.StringBrief(), rs.LatchedLedger.StringBrief(),
		indent, rs.SoundEpoch,
		indent, rs.SoundLedgerSegments.StringBrief(), rs.SoundLedger.StringBrief(),
		indent, rs.Ballots.StringIndented(indent+"REDACTED"),
		indent, rs.FinalEndorse.StringBrief(),
		indent, rs.FinalRatifiers.StringIndented(indent+"REDACTED"),
		indent)
}

//
func (rs *EpochStatus) StringBrief() string {
	return fmt.Sprintf("REDACTED",
		rs.Level, rs.Cycle, rs.Phase, rs.BeginTime)
}
