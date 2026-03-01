package chainchronize

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/go-kit/kit/metrics"
)

const (
	//
	//
	TelemetryComponent = "REDACTED"
)

//

//
type Telemetry struct {
	//
	Chronizing metrics.Gauge
	//
	CountTrans metrics.Gauge
	//
	SumTrans metrics.Gauge
	//
	LedgerExtentOctets metrics.Gauge
	//
	NewestLedgerAltitude metrics.Gauge

	//
	EarlierComprisedLedgers metrics.Counter `metrics_name:"earlier_comprised_ledgers"`

	//
	AbsorbedLedgers metrics.Counter `metrics_name:"absorbed_ledgers"`

	//
	AbsorbedLedgerInterval metrics.Histogram `metrics_buckettype:"experscope" metrics_bucketsizes:"0.1, 100, 8"`
}

func (m *Telemetry) logLedgerTelemetry(ledger *kinds.Ledger) {
	m.CountTrans.Set(float64(len(ledger.Txs)))
	m.SumTrans.Add(float64(len(ledger.Txs)))
	m.LedgerExtentOctets.Set(float64(ledger.Extent()))
	m.NewestLedgerAltitude.Set(float64(ledger.Altitude))
}
