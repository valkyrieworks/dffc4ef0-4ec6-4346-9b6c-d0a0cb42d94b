package kinds

import (
	"fmt"
	"time"

	"github.com/valkyrieworks/utils/bits"
	"github.com/valkyrieworks/kinds"
)

//

//
//
type NodeDurationStatus struct {
	Level int64         `json:"level"` //
	Cycle  int32         `json:"duration"`  //
	Phase   DurationPhaseKind `json:"phase"`   //

	//
	BeginTime time.Time `json:"begin_time"`

	//
	Nomination                   bool                `json:"nomination"`
	NominationLedgerSegmentAssignHeading kinds.SegmentAssignHeading `json:"nomination_ledger_segment_collection_heading"`
	NominationLedgerSegments         *bits.BitList      `json:"nomination_ledger_segments"`
	//
	NominationPOLDuration int32 `json:"nomination_pol_epoch"`

	//
	NominationPOL     *bits.BitList `json:"nomination_pol"`
	Preballots        *bits.BitList `json:"preballots"`          //
	Preendorsements      *bits.BitList `json:"preendorsements"`        //
	FinalEndorseDuration int32          `json:"final_endorse_epoch"` //
	FinalEndorse      *bits.BitList `json:"final_endorse"`       //

	//
	OvertakeEndorseDuration int32 `json:"overtake_endorse_epoch"`

	//
	OvertakeEndorse *bits.BitList `json:"overtake_endorse"`
}

//
func (prs NodeDurationStatus) String() string {
	return prs.StringIndented("REDACTED")
}

//
func (prs NodeDurationStatus) StringIndented(indent string) string {
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTED)
REDACTEDv
REDACTEDv
REDACTED)
REDACTED)
REDACTED`,
		indent, prs.Level, prs.Cycle, prs.Phase, prs.BeginTime,
		indent, prs.NominationLedgerSegmentAssignHeading, prs.NominationLedgerSegments,
		indent, prs.NominationPOL, prs.NominationPOLDuration,
		indent, prs.Preballots,
		indent, prs.Preendorsements,
		indent, prs.FinalEndorse, prs.FinalEndorseDuration,
		indent, prs.OvertakeEndorse, prs.OvertakeEndorseDuration,
		indent)
}
