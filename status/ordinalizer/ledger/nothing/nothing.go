package nothing

import (
	"context"
	"errors"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

var _ ordinalizer.LedgerOrdinalizer = (*PreventerOrdinalizer)(nil)

//
type PreventerOrdinalizer struct{}

func (idx *PreventerOrdinalizer) Has(int64) (bool, error) {
	return false, errors.New("REDACTED")
}

func (idx *PreventerOrdinalizer) Ordinal(kinds.IncidentDataFreshLedgerIncidents) error {
	return nil
}

func (idx *PreventerOrdinalizer) Lookup(context.Context, *inquire.Inquire) ([]int64, error) {
	return []int64{}, nil
}

func (idx *PreventerOrdinalizer) AssignTracer(log.Tracer) {
}
