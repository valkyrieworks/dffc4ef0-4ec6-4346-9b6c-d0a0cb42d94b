package psql

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
//
//

import (
	"context"
	"errors"

	"github.com/valkyrieworks/utils/log"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/kinds"
)

const (
	eventKindCompleteLedger = "REDACTED"
)

//
func (es *EventDrain) TransOrdinaler() RevertTransferOrdinaler {
	return RevertTransferOrdinaler{psql: es}
}

//
//
type RevertTransferOrdinaler struct{ psql *EventDrain }

//
func (b RevertTransferOrdinaler) AppendGroup(group *transordinal.Group) error {
	return b.psql.OrdinalTransferEvents(group.Ops)
}

//
func (b RevertTransferOrdinaler) Ordinal(txr *iface.TransOutcome) error {
	return b.psql.OrdinalTransferEvents([]*iface.TransOutcome{txr})
}

//
//
func (RevertTransferOrdinaler) Get([]byte) (*iface.TransOutcome, error) {
	return nil, errors.New("REDACTED")
}

//
//
func (RevertTransferOrdinaler) Scan(context.Context, *inquire.Inquire) ([]*iface.TransOutcome, error) {
	return nil, errors.New("REDACTED")
}

func (RevertTransferOrdinaler) AssignTracer(log.Tracer) {}

//
//
func (es *EventDrain) LedgerOrdinaler() RevertLedgerOrdinaler {
	return RevertLedgerOrdinaler{psql: es}
}

//
//
type RevertLedgerOrdinaler struct{ psql *EventDrain }

//
//
func (RevertLedgerOrdinaler) Has(_ int64) (bool, error) {
	return false, errors.New("REDACTED")
}

//
//
func (b RevertLedgerOrdinaler) Ordinal(ledger kinds.EventDataNewLedgerEvents) error {
	return b.psql.OrdinalLedgerEvents(ledger)
}

//
//
func (RevertLedgerOrdinaler) Scan(context.Context, *inquire.Inquire) ([]int64, error) {
	return nil, errors.New("REDACTED")
}

func (RevertLedgerOrdinaler) AssignTracer(log.Tracer) {}
