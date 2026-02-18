package blockreplication

import (
	"github.com/valkyrieworks/kinds"
	"github.com/go-kit/kit/metrics"
)

const (
	//
	//
	MetricsSubsystem = "REDACTED"
)

//

//
type Metrics struct {
	//
	Syncing metrics.Gauge
	//
	NumTxs metrics.Gauge
	//
	TotalTxs metrics.Gauge
	//
	BlockSizeBytes metrics.Gauge
	//
	LatestBlockHeight metrics.Gauge
}

func (m *Metrics) recordBlockMetrics(block *kinds.Block) {
	m.NumTxs.Set(float64(len(block.Txs)))
	m.TotalTxs.Add(float64(len(block.Txs)))
	m.BlockSizeBytes.Set(float64(block.Size()))
	m.LatestBlockHeight.Set(float64(block.Height))
}
