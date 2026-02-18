package source

import (
	"context"

	"github.com/valkyrieworks/kinds"
)

//
//
type Source interface {
	//
	LedgerUID() string

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	RapidLedger(ctx context.Context, level int64) (*kinds.RapidLedger, error)

	//
	NotifyProof(context.Context, kinds.Proof) error
}
