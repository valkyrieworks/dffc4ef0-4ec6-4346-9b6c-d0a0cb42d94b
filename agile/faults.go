package agile

import (
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
type FaultAgedHeadlineLapsed struct {
	At  time.Time
	Now time.Time
}

func (e FaultAgedHeadlineLapsed) Failure() string {
	return fmt.Sprintf("REDACTED", e.At, e.Now)
}

//
//
type FaultFreshItemAssignCannotExistReliable struct {
	Rationale kinds.FaultNegationAmpleBallotingPotencyNotated
}

func (e FaultFreshItemAssignCannotExistReliable) Failure() string {
	return fmt.Sprintf("REDACTED", e.Rationale)
}

//
//
type FaultUnfitHeadline struct {
	Rationale error
}

func (e FaultUnfitHeadline) Failure() string {
	return fmt.Sprintf("REDACTED", e.Rationale)
}

//
//
var FaultUnsuccessfulHeadlineIntersectAlluding = errors.New("REDACTED" +
	"REDACTED" +
	"REDACTED")

//
//
type FaultValidationUnsuccessful struct {
	Originating   int64
	To     int64
	Rationale error
}

//
func (e FaultValidationUnsuccessful) Disclose() error {
	return e.Rationale
}

func (e FaultValidationUnsuccessful) Failure() string {
	return fmt.Sprintf("REDACTED", e.Originating, e.To, e.Rationale)
}

//
//
var FaultAgileCustomerOnslaught = errors.New(`REDACTED.
REDACTED.
REDACTED.
REDACTED`,
)

//
//
var FaultNegativeAttestors = errors.New("REDACTED")

//
type FaultDiscordantHeadings struct {
	Ledger        *kinds.AgileLedger
	AttestorPosition int
}

func (e FaultDiscordantHeadings) Failure() string {
	return fmt.Sprintf(
		"REDACTED",
		e.Ledger.Digest(), e.AttestorPosition)
}

//
//
//
//
type FaultNominatorUrgenciesDeviate struct {
	AttestorDigest  []byte
	AttestorPosition int
	LeadingDigest  []byte
}

func (e FaultNominatorUrgenciesDeviate) Failure() string {
	return fmt.Sprintf(
		"REDACTED",
		e.AttestorPosition, e.AttestorDigest, e.LeadingDigest)
}

//

//
//
type faultFlawedAttestor struct {
	Rationale       error
	AttestorPosition int
}

func (e faultFlawedAttestor) Failure() string {
	return fmt.Sprintf("REDACTED", e.AttestorPosition, e.Rationale.Error())
}

var faultNegativeDeviation = errors.New(
	"REDACTED",
)
