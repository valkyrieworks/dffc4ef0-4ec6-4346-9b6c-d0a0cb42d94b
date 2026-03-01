package sqls

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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	incidentKindCulminateLedger = "REDACTED"
)

//
func (es *IncidentReceiver) TransferOrdinalizer() RetrofitTransferOrdinalizer {
	return RetrofitTransferOrdinalizer{sqls: es}
}

//
//
type RetrofitTransferOrdinalizer struct{ sqls *IncidentReceiver }

//
func (b RetrofitTransferOrdinalizer) AppendCluster(cluster *transferordinal.Cluster) error {
	return b.sqls.PositionTransferIncidents(cluster.Ops)
}

//
func (b RetrofitTransferOrdinalizer) Ordinal(txr *iface.TransferOutcome) error {
	return b.sqls.PositionTransferIncidents([]*iface.TransferOutcome{txr})
}

//
//
func (RetrofitTransferOrdinalizer) Get([]byte) (*iface.TransferOutcome, error) {
	return nil, errors.New("REDACTED")
}

//
//
func (RetrofitTransferOrdinalizer) Lookup(context.Context, *inquire.Inquire) ([]*iface.TransferOutcome, error) {
	return nil, errors.New("REDACTED")
}

func (RetrofitTransferOrdinalizer) AssignTracer(log.Tracer) {}

//
//
func (es *IncidentReceiver) LedgerOrdinalizer() RetrofitLedgerOrdinalizer {
	return RetrofitLedgerOrdinalizer{sqls: es}
}

//
//
type RetrofitLedgerOrdinalizer struct{ sqls *IncidentReceiver }

//
//
func (RetrofitLedgerOrdinalizer) Has(_ int64) (bool, error) {
	return false, errors.New("REDACTED")
}

//
//
func (b RetrofitLedgerOrdinalizer) Ordinal(ledger kinds.IncidentDataFreshLedgerIncidents) error {
	return b.sqls.PositionLedgerIncidents(ledger)
}

//
//
func (RetrofitLedgerOrdinalizer) Lookup(context.Context, *inquire.Inquire) ([]int64, error) {
	return nil, errors.New("REDACTED")
}

func (RetrofitLedgerOrdinalizer) AssignTracer(log.Tracer) {}
