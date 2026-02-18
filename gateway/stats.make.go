//

package gateway

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
		ProcedureCadenceMoments: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: []float64{.0001, .0004, .002, .009, .02, .1, .65, 2, 6, 25},
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAndItems...),
	}
}

func NoopStats() *Stats {
	return &Stats{
		ProcedureCadenceMoments: discard.NewHistogram(),
	}
}
