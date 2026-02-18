package void

import (
	"context"
	"errors"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/kinds"
)

var _ ordinaler.LedgerOrdinaler = (*ImpedimentOrdinaler)(nil)

//
type ImpedimentOrdinaler struct{}

func (idx *ImpedimentOrdinaler) Has(int64) (bool, error) {
	return false, errors.New("REDACTED")
}

func (idx *ImpedimentOrdinaler) Ordinal(kinds.EventDataNewLedgerEvents) error {
	return nil
}

func (idx *ImpedimentOrdinaler) Scan(context.Context, *inquire.Inquire) ([]int64, error) {
	return []int64{}, nil
}

func (idx *ImpedimentOrdinaler) AssignTracer(log.Tracer) {
}
