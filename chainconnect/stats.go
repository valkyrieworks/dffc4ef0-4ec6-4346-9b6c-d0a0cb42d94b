package chainconnect

import (
	"github.com/valkyrieworks/kinds"
	"github.com/go-kit/kit/metrics"
)

const (
	//
	//
	StatsComponent = "REDACTED"
)

//

//
type Stats struct {
	//
	Aligning metrics.Gauge
	//
	CountTrans metrics.Gauge
	//
	SumTrans metrics.Gauge
	//
	LedgerVolumeOctets metrics.Gauge
	//
	NewestLedgerLevel metrics.Gauge
}

func (m *Stats) logLedgerStats(ledger *kinds.Ledger) {
	m.CountTrans.Set(float64(len(ledger.Txs)))
	m.SumTrans.Add(float64(len(ledger.Txs)))
	m.LedgerVolumeOctets.Set(float64(ledger.Volume()))
	m.NewestLedgerLevel.Set(float64(ledger.Level))
}
