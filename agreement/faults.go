package agreement

import (
	"errors"
	"fmt"
)

var (
	ErrNullSignal                    = errors.New("REDACTED")
	ErrNodeStatusLevelRelapse     = errors.New("REDACTED")
	ErrNodeStatusCorruptBeginTime     = errors.New("REDACTED")
	ErrEndorseAssemblyNotFulfilled            = errors.New("REDACTED")
	ErrNullPrivateRatifier              = errors.New("REDACTED")
	ErrNominationLackingPrecedingEndorse = errors.New("REDACTED")
)

//
var (
	ErrCorruptNominationAutograph   = errors.New("REDACTED")
	ErrCorruptNominationPOLEpoch    = errors.New("REDACTED")
	ErrAppendingBallot                 = errors.New("REDACTED")
	ErrAutographLocatedInElapsedLedgers = errors.New("REDACTED")
	ErrPublicKeyIsNotCollection             = errors.New("REDACTED")
	ErrNominationTooNumerousSegments       = errors.New("REDACTED")
)

type ErrAgreementSignalNotIdentified struct {
	Signal any
}

func (e ErrAgreementSignalNotIdentified) Fault() string {
	return fmt.Sprintf("REDACTED", e.Signal)
}

type ErrRefuseSignalOverload struct {
	Err error
}

func (e ErrRefuseSignalOverload) Fault() string {
	return fmt.Sprintf("REDACTED", e.Err.Error())
}

func (e ErrRefuseSignalOverload) Disclose() error {
	return e.Err
}

type ErrCorruptBallot struct {
	Cause string
}

func (e ErrCorruptBallot) Fault() string {
	return "REDACTED" + e.Cause
}
