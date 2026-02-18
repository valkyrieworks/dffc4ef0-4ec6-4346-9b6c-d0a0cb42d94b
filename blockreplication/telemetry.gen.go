//

package blockreplication

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		Syncing: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		NumTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		TotalTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		BlockSizeBytes: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		LatestBlockHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		Syncing:           discard.NewGauge(),
		NumTxs:            discard.NewGauge(),
		TotalTxs:          discard.NewGauge(),
		BlockSizeBytes:    discard.NewGauge(),
		LatestBlockHeight: discard.NewGauge(),
	}
}
