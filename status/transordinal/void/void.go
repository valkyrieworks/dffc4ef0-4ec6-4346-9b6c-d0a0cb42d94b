package void

import (
	"context"
	"errors"

	"github.com/valkyrieworks/utils/log"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	"github.com/valkyrieworks/status/transordinal"
)

var _ transordinal.TransOrdinaler = (*TransOrdinal)(nil)

//
type TransOrdinal struct{}

//
func (txi *TransOrdinal) Get(_ []byte) (*iface.TransOutcome, error) {
	return nil, errors.New("REDACTED")
}

//
func (txi *TransOrdinal) AppendGroup(_ *transordinal.Group) error {
	return nil
}

//
func (txi *TransOrdinal) Ordinal(_ *iface.TransOutcome) error {
	return nil
}

func (txi *TransOrdinal) Scan(_ context.Context, _ *inquire.Inquire) ([]*iface.TransOutcome, error) {
	return []*iface.TransOutcome{}, nil
}

func (txi *TransOrdinal) AssignTracer(log.Tracer) {

}
