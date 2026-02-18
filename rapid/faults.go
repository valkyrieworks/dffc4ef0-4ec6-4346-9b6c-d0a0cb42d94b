package rapid

import (
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/kinds"
)

//
//
//
type ErrAgedHeadingLapsed struct {
	At  time.Time
	Now time.Time
}

func (e ErrAgedHeadingLapsed) Fault() string {
	return fmt.Sprintf("REDACTED", e.At, e.Now)
}

//
//
type ErrNewValueCollectionCannotBeValidated struct {
	Cause kinds.ErrNoSufficientPollingEnergyAttested
}

func (e ErrNewValueCollectionCannotBeValidated) Fault() string {
	return fmt.Sprintf("REDACTED", e.Cause)
}

//
//
type ErrCorruptHeading struct {
	Cause error
}

func (e ErrCorruptHeading) Fault() string {
	return fmt.Sprintf("REDACTED", e.Cause)
}

//
//
var ErrErroredHeadingIntersectPointing = errors.New("REDACTED" +
	"REDACTED" +
	"REDACTED")

//
//
type ErrValidationErrored struct {
	From   int64
	To     int64
	Cause error
}

//
func (e ErrValidationErrored) Disclose() error {
	return e.Cause
}

func (e ErrValidationErrored) Fault() string {
	return fmt.Sprintf("REDACTED", e.From, e.To, e.Cause)
}

//
//
var ErrRapidCustomerAssault = errors.New(`REDACTED.
REDACTED.
REDACTED.
REDACTED`,
)

//
//
var ErrNoAttestors = errors.New("REDACTED")

//
type ErrClashingHeadings struct {
	Ledger        *kinds.RapidLedger
	AttestorOrdinal int
}

func (e ErrClashingHeadings) Fault() string {
	return fmt.Sprintf(
		"REDACTED",
		e.Ledger.Digest(), e.AttestorOrdinal)
}

//
//
//
//
type ErrRecommenderUrgenciesDeviate struct {
	AttestorDigest  []byte
	AttestorOrdinal int
	LeadingDigest  []byte
}

func (e ErrRecommenderUrgenciesDeviate) Fault() string {
	return fmt.Sprintf(
		"REDACTED",
		e.AttestorOrdinal, e.AttestorDigest, e.LeadingDigest)
}

//

//
//
type errFlawedAttestor struct {
	Cause       error
	AttestorOrdinal int
}

func (e errFlawedAttestor) Fault() string {
	return fmt.Sprintf("REDACTED", e.AttestorOrdinal, e.Cause.Error())
}

var errNoDeviation = errors.New(
	"REDACTED",
)
