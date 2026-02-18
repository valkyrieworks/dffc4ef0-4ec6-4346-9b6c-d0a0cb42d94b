package txpool

import (
	"errors"
	"fmt"
)

//
var ErrTransferNegateLocated = errors.New("REDACTED")

//
var ErrTransferInRepository = errors.New("REDACTED")

//
//
var ErrRevalidateComplete = errors.New("REDACTED")

//
//
type ErrTransferTooBulky struct {
	Max    int
	Factual int
}

func (e ErrTransferTooBulky) Fault() string {
	return fmt.Sprintf("REDACTED", e.Max, e.Factual)
}

//
//
type ErrTxpoolIsComplete struct {
	CountTrans      int
	MaximumTrans      int
	TransOctets    int64
	MaximumTransOctets int64
	RevalidateComplete bool
}

func (e ErrTxpoolIsComplete) Fault() string {
	return fmt.Sprintf(
		"REDACTED",
		e.CountTrans,
		e.MaximumTrans,
		e.TransOctets,
		e.MaximumTransOctets,
	)
}

//
type ErrPreInspect struct {
	Err error
}

func (e ErrPreInspect) Fault() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

func (e ErrPreInspect) Disclose() error {
	return e.Err
}

//
func IsPreInspectFault(err error) bool {
	return errors.As(err, &ErrPreInspect{})
}

type ErrApplicationLinkTxpool struct {
	Err error
}

func (e ErrApplicationLinkTxpool) Fault() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

func (e ErrApplicationLinkTxpool) Disclose() error {
	return e.Err
}

type ErrPurgeApplicationLink struct {
	Err error
}

func (e ErrPurgeApplicationLink) Fault() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

func (e ErrPurgeApplicationLink) Disclose() error {
	return e.Err
}
