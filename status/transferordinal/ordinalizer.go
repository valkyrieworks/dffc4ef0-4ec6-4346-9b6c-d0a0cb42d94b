package transferordinal

import (
	"context"
	"errors"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
)

//

//

//
type TransferOrdinalizer interface {
	//
	AppendCluster(b *Cluster) error

	//
	Ordinal(outcome *iface.TransferOutcome) error

	//
	//
	Get(digest []byte) (*iface.TransferOutcome, error)

	//
	Lookup(ctx context.Context, q *inquire.Inquire) ([]*iface.TransferOutcome, error)

	//
	AssignTracer(l log.Tracer)
}

//
//
type Cluster struct {
	Ops []*iface.TransferOutcome
}

//
func FreshCluster(n int64) *Cluster {
	return &Cluster{
		Ops: make([]*iface.TransferOutcome, n),
	}
}

//
func (b *Cluster) Add(outcome *iface.TransferOutcome) error {
	b.Ops[outcome.Ordinal] = outcome
	return nil
}

//
func (b *Cluster) Extent() int {
	return len(b.Ops)
}

//
var FailureBlankDigest = errors.New("REDACTED")
