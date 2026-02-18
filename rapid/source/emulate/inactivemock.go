package emulate

import (
	"context"

	"github.com/valkyrieworks/rapid/source"
	"github.com/valkyrieworks/kinds"
)

type inactiveEmulate struct {
	ledgerUID string
}

//
func NewInactiveEmulate(ledgerUID string) source.Source {
	return &inactiveEmulate{ledgerUID: ledgerUID}
}

func (p *inactiveEmulate) LedgerUID() string { return p.ledgerUID }

func (p *inactiveEmulate) String() string { return "REDACTED" }

func (p *inactiveEmulate) RapidLedger(context.Context, int64) (*kinds.RapidLedger, error) {
	return nil, source.ErrNoReply
}

func (p *inactiveEmulate) NotifyProof(context.Context, kinds.Proof) error {
	return source.ErrNoReply
}
