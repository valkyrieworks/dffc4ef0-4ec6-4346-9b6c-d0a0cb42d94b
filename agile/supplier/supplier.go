package supplier

import (
	"context"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
type Supplier interface {
	//
	SuccessionUUID() string

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
	AgileLedger(ctx context.Context, altitude int64) (*kinds.AgileLedger, error)

	//
	NotifyProof(context.Context, kinds.Proof) error
}
