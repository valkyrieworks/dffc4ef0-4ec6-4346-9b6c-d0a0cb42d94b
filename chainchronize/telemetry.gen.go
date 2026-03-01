//

package chainchronize

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
		CountTrans: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		SumTrans: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		LedgerExtentOctets: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		NewestLedgerAltitude: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		EarlierComprisedLedgers: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		AbsorbedLedgers: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		AbsorbedLedgerInterval: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, tags).With(tagsAlsoItems...),
	}
}

func NooperationTelemetry() *Telemetry {
	return &Telemetry{
		Chronizing:               discard.NewGauge(),
		CountTrans:                discard.NewGauge(),
		SumTrans:              discard.NewGauge(),
		LedgerExtentOctets:        discard.NewGauge(),
		NewestLedgerAltitude:     discard.NewGauge(),
		EarlierComprisedLedgers: discard.NewCounter(),
		AbsorbedLedgers:        discard.NewCounter(),
		AbsorbedLedgerInterval: discard.NewHistogram(),
	}
}
