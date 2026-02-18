package ordinaler

import (
	"context"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	"github.com/valkyrieworks/kinds"
)

//

//
type LedgerOrdinaler interface {
	//
	//
	Has(level int64) (bool, error)

	//
	Ordinal(kinds.EventDataNewLedgerEvents) error

	//
	//
	Scan(ctx context.Context, q *inquire.Inquire) ([]int64, error)

	AssignTracer(l log.Tracer)
}
