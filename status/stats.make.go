//

package status

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func MonitorstatsStats(scope string, tagsAndItems ...string) *Stats {
	tags := []string{}
	for i := 0; i < len(tagsAndItems); i += 2 {
		tags = append(tags, tagsAndItems[i])
	}
	return &Stats{
		LedgerExecutionTime: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.LinearBuckets(1, 10, 10),
		}, tags).With(tagsAndItems...),
		AgreementArgumentRefreshes: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		RatifierCollectionRefreshes: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
	}
}

func NoopStats() *Stats {
	return &Stats{
		LedgerExecutionTime:   discard.NewHistogram(),
		AgreementArgumentRefreshes: discard.NewCounter(),
		RatifierCollectionRefreshes:   discard.NewCounter(),
	}
}
