//

package statuschronize

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
		Chronizing: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
	}
}

func NooperationTelemetry() *Telemetry {
	return &Telemetry{
		Chronizing: discard.NewGauge(),
	}
}
