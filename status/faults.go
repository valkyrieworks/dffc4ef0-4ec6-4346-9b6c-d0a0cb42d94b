package status

import (
	"errors"
	"fmt"
)

type (
	ErrCorruptLedger error
	ErrGatewayApplicationLink error

	ErrUnclearLedger struct {
		Level int64
	}

	ErrLedgerDigestDiscrepancy struct {
		CoreDigest []byte
		ApplicationDigest  []byte
		Level   int64
	}

	ErrApplicationLedgerLevelTooSuperior struct {
		CoreLevel int64
		ApplicationLevel  int64
	}

	ErrApplicationLedgerLevelTooInferior struct {
		ApplicationLevel int64
		DepotRoot int64
	}

	ErrFinalStatusDiscrepancy struct {
		Level int64
		Core   []byte
		App    []byte
	}

	ErrStatusDiscrepancy struct {
		Got      *Status
		Anticipated *Status
	}

	ErrNoValueCollectionForLevel struct {
		Level int64
	}

	ErrNoAgreementOptionsForLevel struct {
		Level int64
	}

	ErrNoIfaceRepliesForLevel struct {
		Level int64
	}

	ErrIfaceReplyReplyUnserializeForLevel struct {
		Level int64
	}

	ErrIfaceReplyTaintedOrBlueprintAlterForLevel struct {
		Err    error
		Level int64
	}
)

func (e ErrUnclearLedger) Fault() string {
	return fmt.Sprintf("REDACTED", e.Level)
}

func (e ErrLedgerDigestDiscrepancy) Fault() string {
	return fmt.Sprintf(
		"REDACTED",
		e.ApplicationDigest,
		e.CoreDigest,
		e.Level,
	)
}

func (e ErrApplicationLedgerLevelTooSuperior) Fault() string {
	return fmt.Sprintf("REDACTED", e.ApplicationLevel, e.CoreLevel)
}

func (e ErrApplicationLedgerLevelTooInferior) Fault() string {
	return fmt.Sprintf("REDACTED", e.ApplicationLevel, e.DepotRoot)
}

func (e ErrFinalStatusDiscrepancy) Fault() string {
	return fmt.Sprintf(
		"REDACTED",
		e.Level,
		e.Core,
		e.App,
	)
}

func (e ErrStatusDiscrepancy) Fault() string {
	return fmt.Sprintf(
		"REDACTED",
		e.Got,
		e.Anticipated,
	)
}

func (e ErrNoValueCollectionForLevel) Fault() string {
	return fmt.Sprintf("REDACTED", e.Level)
}

func (e ErrNoAgreementOptionsForLevel) Fault() string {
	return fmt.Sprintf("REDACTED", e.Level)
}

func (e ErrNoIfaceRepliesForLevel) Fault() string {
	return fmt.Sprintf("REDACTED", e.Level)
}

func (e ErrIfaceReplyReplyUnserializeForLevel) Fault() string {
	return fmt.Sprintf("REDACTED", e.Level)
}

func (e ErrIfaceReplyTaintedOrBlueprintAlterForLevel) Fault() string {
	return fmt.Sprintf("REDACTED", e.Level)
}

func (e ErrIfaceReplyTaintedOrBlueprintAlterForLevel) Disclose() error {
	return e.Err
}

var ErrCompleteLedgerRepliesNotSustained = errors.New("REDACTED")
