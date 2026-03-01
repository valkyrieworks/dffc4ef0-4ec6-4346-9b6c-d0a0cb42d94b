package txpool

import (
	"errors"
	"fmt"
)

//
var FaultTransferNegationDetected = errors.New("REDACTED")

//
var FaultTransferInsideStash = errors.New("REDACTED")

//
//
var FaultReinspectComplete = errors.New("REDACTED")

//
//
type FaultTransferExcessivelyAmple struct {
	Max    int
	Existing int
}

func (e FaultTransferExcessivelyAmple) Failure() string {
	return fmt.Sprintf("REDACTED", e.Max, e.Existing)
}

//
//
type FaultTxpoolEqualsComplete struct {
	CountTrans      int
	MaximumTrans      int
	TransOctets    int64
	MaximumTransOctets int64
	ReinspectComplete bool
}

func (e FaultTxpoolEqualsComplete) Failure() string {
	return fmt.Sprintf(
		"REDACTED",
		e.CountTrans,
		e.MaximumTrans,
		e.TransOctets,
		e.MaximumTransOctets,
	)
}

//
type FaultAnteInspect struct {
	Err error
}

func (e FaultAnteInspect) Failure() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

func (e FaultAnteInspect) Disclose() error {
	return e.Err
}

//
func EqualsPriorInspectFailure(err error) bool {
	return errors.As(err, &FaultAnteInspect{})
}

type FaultApplicationLinkTxpool struct {
	Err error
}

func (e FaultApplicationLinkTxpool) Failure() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

func (e FaultApplicationLinkTxpool) Disclose() error {
	return e.Err
}

type FaultPurgeApplicationLink struct {
	Err error
}

func (e FaultPurgeApplicationLink) Failure() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

func (e FaultPurgeApplicationLink) Disclose() error {
	return e.Err
}
