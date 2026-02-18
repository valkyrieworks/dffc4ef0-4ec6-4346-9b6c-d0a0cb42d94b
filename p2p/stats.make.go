//

package p2p

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
		Nodes: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		NodeAcceptOctetsSum: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAndItems...),
		NodeTransmitOctetsSum: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAndItems...),
		NodeAwaitingTransmitOctets: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		CountTrans: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		SignalAcceptOctetsSum: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		SignalTransmitOctetsSum: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		SignalsAccepted: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAndItems...),
		SignalsHandlerInJourney: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAndItems...),
		SignalsHandlerAwaitingPeriod: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAndItems...),
		SignalHandlerAcceptPeriod: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED", "REDACTED")).With(tagsAndItems...),
		SignalHandlerBufferParallelism: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
	}
}

func NoopStats() *Stats {
	return &Stats{
		Nodes:                          discard.NewGauge(),
		NodeAcceptOctetsSum:          discard.NewCounter(),
		NodeTransmitOctetsSum:             discard.NewCounter(),
		NodeAwaitingTransmitOctets:           discard.NewGauge(),
		CountTrans:                         discard.NewGauge(),
		SignalAcceptOctetsSum:       discard.NewCounter(),
		SignalTransmitOctetsSum:          discard.NewCounter(),
		SignalsAccepted:               discard.NewCounter(),
		SignalsHandlerInJourney:        discard.NewGauge(),
		SignalsHandlerAwaitingPeriod: discard.NewHistogram(),
		SignalHandlerAcceptPeriod:  discard.NewHistogram(),
		SignalHandlerBufferParallelism: discard.NewGauge(),
	}
}
