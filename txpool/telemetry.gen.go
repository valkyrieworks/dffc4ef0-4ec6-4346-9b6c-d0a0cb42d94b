//

package txpool

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
		Extent: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		ExtentOctets: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		TransferExtentOctets: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.ExponentialBuckets(1, 3, 7),
		}, tags).With(tagsAlsoItems...),
		UnsuccessfulTrans: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		DeclinedTrans: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		ExpelledTrans: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		ReinspectMultiples: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		DynamicOutgoingLinkages: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		EarlierAcceptedTrans: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		ClusterExtent: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: []float64{1, 2, 5, 10, 30, 50, 100, 200, 300},
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		HarvestedTrans: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
	}
}

func NooperationTelemetry() *Telemetry {
	return &Telemetry{
		Extent:                      discard.NewGauge(),
		ExtentOctets:                 discard.NewGauge(),
		TransferExtentOctets:               discard.NewHistogram(),
		UnsuccessfulTrans:                 discard.NewCounter(),
		DeclinedTrans:               discard.NewCounter(),
		ExpelledTrans:                discard.NewCounter(),
		ReinspectMultiples:              discard.NewCounter(),
		DynamicOutgoingLinkages: discard.NewGauge(),
		EarlierAcceptedTrans:        discard.NewCounter(),
		ClusterExtent:                 discard.NewHistogram(),
		HarvestedTrans:                 discard.NewCounter(),
	}
}
