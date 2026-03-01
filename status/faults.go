package status

import (
	"errors"
	"fmt"
)

type (
	FaultUnfitLedger error
	FaultDelegateApplicationLink error

	FaultUnfamiliarLedger struct {
		Altitude int64
	}

	FaultLedgerDigestDiscrepancy struct {
		BaseDigest []byte
		PlatformDigest  []byte
		Altitude   int64
	}

	FaultApplicationLedgerAltitudeExcessivelySuperior struct {
		BaseAltitude int64
		ApplicationAltitude  int64
	}

	FaultApplicationLedgerAltitudeExcessivelyInferior struct {
		ApplicationAltitude int64
		DepotFoundation int64
	}

	FaultFinalStatusDiscrepancy struct {
		Altitude int64
		Base   []byte
		App    []byte
	}

	FaultStatusDiscrepancy struct {
		Got      *Status
		Anticipated *Status
	}

	FaultNegativeItemAssignForeachAltitude struct {
		Altitude int64
	}

	FaultNegativeAgreementParametersForeachAltitude struct {
		Altitude int64
	}

	FaultNegativeIfaceRepliesForeachAltitude struct {
		Altitude int64
	}

	FaultIfaceReplyReplyDecodeForeachAltitude struct {
		Altitude int64
	}

	FaultIfaceReplyTaintedEitherBlueprintAlterationForeachAltitude struct {
		Err    error
		Altitude int64
	}
)

func (e FaultUnfamiliarLedger) Failure() string {
	return fmt.Sprintf("REDACTED", e.Altitude)
}

func (e FaultLedgerDigestDiscrepancy) Failure() string {
	return fmt.Sprintf(
		"REDACTED",
		e.PlatformDigest,
		e.BaseDigest,
		e.Altitude,
	)
}

func (e FaultApplicationLedgerAltitudeExcessivelySuperior) Failure() string {
	return fmt.Sprintf("REDACTED", e.ApplicationAltitude, e.BaseAltitude)
}

func (e FaultApplicationLedgerAltitudeExcessivelyInferior) Failure() string {
	return fmt.Sprintf("REDACTED", e.ApplicationAltitude, e.DepotFoundation)
}

func (e FaultFinalStatusDiscrepancy) Failure() string {
	return fmt.Sprintf(
		"REDACTED",
		e.Altitude,
		e.Base,
		e.App,
	)
}

func (e FaultStatusDiscrepancy) Failure() string {
	return fmt.Sprintf(
		"REDACTED",
		e.Got,
		e.Anticipated,
	)
}

func (e FaultNegativeItemAssignForeachAltitude) Failure() string {
	return fmt.Sprintf("REDACTED", e.Altitude)
}

func (e FaultNegativeAgreementParametersForeachAltitude) Failure() string {
	return fmt.Sprintf("REDACTED", e.Altitude)
}

func (e FaultNegativeIfaceRepliesForeachAltitude) Failure() string {
	return fmt.Sprintf("REDACTED", e.Altitude)
}

func (e FaultIfaceReplyReplyDecodeForeachAltitude) Failure() string {
	return fmt.Sprintf("REDACTED", e.Altitude)
}

func (e FaultIfaceReplyTaintedEitherBlueprintAlterationForeachAltitude) Failure() string {
	return fmt.Sprintf("REDACTED", e.Altitude)
}

func (e FaultIfaceReplyTaintedEitherBlueprintAlterationForeachAltitude) Disclose() error {
	return e.Err
}

var FaultCulminateLedgerRepliesNegationStored = errors.New("REDACTED")
