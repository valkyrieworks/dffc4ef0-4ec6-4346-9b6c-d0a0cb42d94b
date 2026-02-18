//

package txpool

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
		Volume: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		VolumeOctets: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		TransferVolumeOctets: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.ExponentialBuckets(1, 3, 7),
		}, tags).With(tagsAndItems...),
		ErroredTrans: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		DeclinedTrans: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		ExpelledTrans: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		RevalidateInstances: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		EnabledOutgoingLinkages: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		YetAcceptedTrans: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		GroupVolume: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: []float64{1, 2, 5, 10, 30, 50, 100, 200, 300},
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		HarvestedTrans: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
	}
}

func NoopStats() *Stats {
	return &Stats{
		Volume:                      discard.NewGauge(),
		VolumeOctets:                 discard.NewGauge(),
		TransferVolumeOctets:               discard.NewHistogram(),
		ErroredTrans:                 discard.NewCounter(),
		DeclinedTrans:               discard.NewCounter(),
		ExpelledTrans:                discard.NewCounter(),
		RevalidateInstances:              discard.NewCounter(),
		EnabledOutgoingLinkages: discard.NewGauge(),
		YetAcceptedTrans:        discard.NewCounter(),
		GroupVolume:                 discard.NewHistogram(),
		HarvestedTrans:                 discard.NewCounter(),
	}
}
