package log

import (
	"fmt"

	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
)

type IdleFormat struct {
	layout string
	arguments   []any
}

//
//
//
func FreshIdleFormat(layout string, arguments ...any) *IdleFormat {
	return &IdleFormat{layout, arguments}
}

func (l *IdleFormat) Text() string {
	return fmt.Sprintf(l.layout, l.arguments...)
}

type IdleLedgerDigest struct {
	ledger digestible
}

type digestible interface {
	Digest() tendermintoctets.HexadecimalOctets
}

//
//
//
func FreshIdleLedgerDigest(ledger digestible) *IdleLedgerDigest {
	return &IdleLedgerDigest{ledger}
}

func (l *IdleLedgerDigest) Text() string {
	return l.ledger.Digest().Text()
}
