package simulate

import (
	"context"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/supplier"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

type lifelessSimulate struct {
	successionUUID string
}

//
func FreshLifelessSimulate(successionUUID string) supplier.Supplier {
	return &lifelessSimulate{successionUUID: successionUUID}
}

func (p *lifelessSimulate) SuccessionUUID() string { return p.successionUUID }

func (p *lifelessSimulate) Text() string { return "REDACTED" }

func (p *lifelessSimulate) AgileLedger(context.Context, int64) (*kinds.AgileLedger, error) {
	return nil, supplier.FaultNegativeReply
}

func (p *lifelessSimulate) NotifyProof(context.Context, kinds.Proof) error {
	return supplier.FaultNegativeReply
}
