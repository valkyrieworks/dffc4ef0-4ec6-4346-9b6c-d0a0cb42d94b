//

package p2p

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
		Nodes: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		NodeAcceptOctetsSum: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAlsoItems...),
		NodeTransmitOctetsSum: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAlsoItems...),
		NodeTransmitStagingExtent: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		CountTrans: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		ArtifactAcceptOctetsSum: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		ArtifactTransmitOctetsSum: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		SignalsAccepted: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAlsoItems...),
		SignalsHandlerInsideAirtime: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAlsoItems...),
		SignalsHandlerAwaitingInterval: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAlsoItems...),
		ArtifactHandlerAcceptInterval: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAlsoItems...),
		ArtifactHandlerStagingParallelism: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
	}
}

func NooperationTelemetry() *Telemetry {
	return &Telemetry{
		Nodes:                          discard.NewGauge(),
		NodeAcceptOctetsSum:          discard.NewCounter(),
		NodeTransmitOctetsSum:             discard.NewCounter(),
		NodeTransmitStagingExtent:              discard.NewGauge(),
		CountTrans:                         discard.NewGauge(),
		ArtifactAcceptOctetsSum:       discard.NewCounter(),
		ArtifactTransmitOctetsSum:          discard.NewCounter(),
		SignalsAccepted:               discard.NewCounter(),
		SignalsHandlerInsideAirtime:        discard.NewGauge(),
		SignalsHandlerAwaitingInterval: discard.NewHistogram(),
		ArtifactHandlerAcceptInterval:  discard.NewHistogram(),
		ArtifactHandlerStagingParallelism: discard.NewGauge(),
	}
}
