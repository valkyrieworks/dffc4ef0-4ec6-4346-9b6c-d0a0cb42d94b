package source

import (
	"errors"
	"fmt"
)

var (
	//
	//
	ErrLevelTooElevated = errors.New("REDACTED")
	//
	//
	//
	ErrRapidLedgerNegateLocated = errors.New("REDACTED")
	//
	//
	ErrNoReply = errors.New("REDACTED")
)

//
//
type ErrFlawedRapidLedger struct {
	Cause error
}

func (e ErrFlawedRapidLedger) Fault() string {
	return fmt.Sprintf("REDACTED", e.Cause.Error())
}
