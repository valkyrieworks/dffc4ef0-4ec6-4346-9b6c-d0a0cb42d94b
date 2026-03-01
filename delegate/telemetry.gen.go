//

package delegate

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func TitanTelemetry(scope string, tagsAlsoItems ...string) *Telemetry {
	tags := []string{}
	for i := 0; i < len(tagsAlsoItems); i += 2 {
		tags = append(tags, tagsAlsoItems[i])
	}
	return &Telemetry{
		ProcedureScheduleMoments: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: []float64{.0001, .0004, .002, .009, .02, .1, .65, 2, 6, 25},
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAlsoItems...),
	}
}

func NooperationTelemetry() *Telemetry {
	return &Telemetry{
		ProcedureScheduleMoments: discard.NewHistogram(),
	}
}
