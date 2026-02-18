package log

import (
	"fmt"

	cometbytes "github.com/valkyrieworks/utils/octets"
)

type IdleFormat struct {
	layout string
	args   []any
}

//
//
//
func NewIdleFormat(layout string, args ...any) *IdleFormat {
	return &IdleFormat{layout, args}
}

func (l *IdleFormat) String() string {
	return fmt.Sprintf(l.layout, l.args...)
}

type IdleLedgerDigest struct {
	ledger digestible
}

type digestible interface {
	Digest() cometbytes.HexOctets
}

//
//
//
func NewIdleLedgerDigest(ledger digestible) *IdleLedgerDigest {
	return &IdleLedgerDigest{ledger}
}

func (l *IdleLedgerDigest) String() string {
	return l.ledger.Digest().String()
}
