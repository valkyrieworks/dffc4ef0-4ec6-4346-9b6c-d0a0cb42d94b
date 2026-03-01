//

package status

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
		LedgerHandlingMoment: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.LinearBuckets(1, 10, 10),
		}, tags).With(tagsAlsoItems...),
		AgreementArgumentRevisions: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		AssessorAssignRevisions: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
	}
}

func NooperationTelemetry() *Telemetry {
	return &Telemetry{
		LedgerHandlingMoment:   discard.NewHistogram(),
		AgreementArgumentRevisions: discard.NewCounter(),
		AssessorAssignRevisions:   discard.NewCounter(),
	}
}
