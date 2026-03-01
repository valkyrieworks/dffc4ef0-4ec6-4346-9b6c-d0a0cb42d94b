package supplier

import (
	"errors"
	"fmt"
)

var (
	//
	//
	FaultAltitudeExcessivelyTall = errors.New("REDACTED")
	//
	//
	//
	FaultAgileLedgerNegationDetected = errors.New("REDACTED")
	//
	//
	FaultNegativeReply = errors.New("REDACTED")
)

//
//
type FaultFlawedAgileLedger struct {
	Rationale error
}

func (e FaultFlawedAgileLedger) Failure() string {
	return fmt.Sprintf("REDACTED", e.Rationale.Error())
}
