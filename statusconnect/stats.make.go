//

package statusconnect

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
		Aligning: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
	}
}

func NoopStats() *Stats {
	return &Stats{
		Aligning: discard.NewGauge(),
	}
}
