package agreement

import (
	"errors"
	"fmt"
)

var (
	FaultVoidSignal                    = errors.New("REDACTED")
	FaultNodeStatusAltitudeRelapse     = errors.New("REDACTED")
	FaultNodeStatusUnfitInitiateMoment     = errors.New("REDACTED")
	FaultEndorseAssemblyNegationFulfilled            = errors.New("REDACTED")
	FaultVoidPrivateAssessor              = errors.New("REDACTED")
	FaultNominationLackingPriorEndorse = errors.New("REDACTED")
)

//
var (
	FaultUnfitNominationSigning   = errors.New("REDACTED")
	FaultUnfitNominationPolicyIteration    = errors.New("REDACTED")
	FaultAppendingBallot                 = errors.New("REDACTED")
	FaultSigningDetectedInsideElapsedLedgers = errors.New("REDACTED")
	FaultPublicTokenEqualsNegationAssign             = errors.New("REDACTED")
	FaultNominationExcessivelyMultipleFragments       = errors.New("REDACTED")
)

//
var (
	FaultCertification      = errors.New("REDACTED")
	FaultEarlierComprised = errors.New("REDACTED")
	FaultAltitudeBreach       = errors.New("REDACTED")
)

type FaultAgreementSignalNegationIdentified struct {
	Signal any
}

func (e FaultAgreementSignalNegationIdentified) Failure() string {
	return fmt.Sprintf("REDACTED", e.Signal)
}

type FaultRefuseSignalOverrun struct {
	Err error
}

func (e FaultRefuseSignalOverrun) Failure() string {
	return fmt.Sprintf("REDACTED", e.Err.Error())
}

func (e FaultRefuseSignalOverrun) Disclose() error {
	return e.Err
}

type FaultUnfitBallot struct {
	Rationale string
}

func (e FaultUnfitBallot) Failure() string {
	return "REDACTED" + e.Rationale
}
