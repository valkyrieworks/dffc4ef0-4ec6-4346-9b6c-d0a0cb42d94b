package ordinalizer

import (
	"context"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//

//
type LedgerOrdinalizer interface {
	//
	//
	Has(altitude int64) (bool, error)

	//
	Ordinal(kinds.IncidentDataFreshLedgerIncidents) error

	//
	//
	Lookup(ctx context.Context, q *inquire.Inquire) ([]int64, error)

	AssignTracer(l log.Tracer)
}
