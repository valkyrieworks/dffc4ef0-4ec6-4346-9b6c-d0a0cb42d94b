package kinds

import (
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/digits"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//

//
//
type NodeIterationStatus struct {
	Altitude int64         `json:"altitude"` //
	Iteration  int32         `json:"iteration"`  //
	Phase   IterationPhaseKind `json:"phase"`   //

	//
	InitiateMoment time.Time `json:"initiate_moment"`

	//
	Nomination                   bool                `json:"nomination"`
	NominationLedgerFragmentAssignHeading kinds.FragmentAssignHeading `json:"nomination_ledger_fragment_assign_heading"`
	NominationLedgerFragments         *digits.DigitSeries      `json:"nomination_ledger_fragments"`
	//
	NominationPolicyIteration int32 `json:"nomination_policy_iteration"`

	//
	NominationPolicy     *digits.DigitSeries `json:"nomination_policy"`
	Preballots        *digits.DigitSeries `json:"preballots"`          //
	Preendorsements      *digits.DigitSeries `json:"preendorsements"`        //
	FinalEndorseIteration int32          `json:"final_endorse_iteration"` //
	FinalEndorse      *digits.DigitSeries `json:"final_endorse"`       //

	//
	OvertakeEndorseIteration int32 `json:"overtake_endorse_iteration"`

	//
	OvertakeEndorse *digits.DigitSeries `json:"overtake_endorse"`
}

//
func (prs NodeIterationStatus) Text() string {
	return prs.TextFormatted("REDACTED")
}

//
func (prs NodeIterationStatus) TextFormatted(format string) string {
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTED)
REDACTEDv
REDACTEDv
REDACTED)
REDACTED)
REDACTED`,
		format, prs.Altitude, prs.Iteration, prs.Phase, prs.InitiateMoment,
		format, prs.NominationLedgerFragmentAssignHeading, prs.NominationLedgerFragments,
		format, prs.NominationPolicy, prs.NominationPolicyIteration,
		format, prs.Preballots,
		format, prs.Preendorsements,
		format, prs.FinalEndorse, prs.FinalEndorseIteration,
		format, prs.OvertakeEndorse, prs.OvertakeEndorseIteration,
		format)
}
