package nothing

import (
	"context"
	"errors"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
)

var _ transferordinal.TransferOrdinalizer = (*TransferOrdinal)(nil)

//
type TransferOrdinal struct{}

//
func (txi *TransferOrdinal) Get(_ []byte) (*iface.TransferOutcome, error) {
	return nil, errors.New("REDACTED")
}

//
func (txi *TransferOrdinal) AppendCluster(_ *transferordinal.Cluster) error {
	return nil
}

//
func (txi *TransferOrdinal) Ordinal(_ *iface.TransferOutcome) error {
	return nil
}

func (txi *TransferOrdinal) Lookup(_ context.Context, _ *inquire.Inquire) ([]*iface.TransferOutcome, error) {
	return []*iface.TransferOutcome{}, nil
}

func (txi *TransferOrdinal) AssignTracer(log.Tracer) {

}
