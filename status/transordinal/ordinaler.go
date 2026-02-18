package transordinal

import (
	"context"
	"errors"

	"github.com/valkyrieworks/utils/log"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/broadcast/inquire"
)

//

//

//
type TransOrdinaler interface {
	//
	AppendGroup(b *Group) error

	//
	Ordinal(outcome *iface.TransOutcome) error

	//
	//
	Get(digest []byte) (*iface.TransOutcome, error)

	//
	Scan(ctx context.Context, q *inquire.Inquire) ([]*iface.TransOutcome, error)

	//
	AssignTracer(l log.Tracer)
}

//
//
type Group struct {
	Ops []*iface.TransOutcome
}

//
func NewGroup(n int64) *Group {
	return &Group{
		Ops: make([]*iface.TransOutcome, n),
	}
}

//
func (b *Group) Add(outcome *iface.TransOutcome) error {
	b.Ops[outcome.Ordinal] = outcome
	return nil
}

//
func (b *Group) Volume() int {
	return len(b.Ops)
}

//
var FaultEmptyDigest = errors.New("REDACTED")
