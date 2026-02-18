//

package agreement

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
		Level: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		RatifierFinalAttestedLevel: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		Iterations: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		DurationPeriodMoments: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, tags).With(tagsAndItems...),
		Ratifiers: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		RatifiersEnergy: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		RatifierEnergy: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		RatifierSkippedLedgers: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		AbsentRatifiers: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		AbsentRatifiersEnergy: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		FaultyRatifiers: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		FaultyRatifiersEnergy: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		LedgerCadenceMoments: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		CountTrans: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		LedgerVolumeOctets: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		SeriesVolumeOctets: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		SumTrans: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		ConfirmedLevel: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		LedgerSegments: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		ReplicatedLedgerSegment: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		ReplicatedBallot: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		PhasePeriodMoments: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		LedgerGossipSegmentsAccepted: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		AssemblyPreballotDeferral: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		AssemblyPreendorseDeferral: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		CompletePreballotDeferral: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		PreendorsementsTallied: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		PreendorsementsStakingRatio: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		BallotAdditionAcceptTally: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		NominationAcceptTally: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		NominationInstantiateTally: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAndItems...),
		DurationPollingEnergyFraction: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		TardyBallots: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		NodeLevel: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
		DurationAugmentSum: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: StatsComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAndItems...),
	}
}

func NoopStats() *Stats {
	return &Stats{
		Level:                      discard.NewGauge(),
		RatifierFinalAttestedLevel:   discard.NewGauge(),
		Iterations:                      discard.NewGauge(),
		DurationPeriodMoments:        discard.NewHistogram(),
		Ratifiers:                  discard.NewGauge(),
		RatifiersEnergy:             discard.NewGauge(),
		RatifierEnergy:              discard.NewGauge(),
		RatifierSkippedLedgers:       discard.NewGauge(),
		AbsentRatifiers:           discard.NewGauge(),
		AbsentRatifiersEnergy:      discard.NewGauge(),
		FaultyRatifiers:         discard.NewGauge(),
		FaultyRatifiersEnergy:    discard.NewGauge(),
		LedgerCadenceMoments:        discard.NewHistogram(),
		CountTrans:                      discard.NewGauge(),
		LedgerVolumeOctets:              discard.NewGauge(),
		SeriesVolumeOctets:              discard.NewCounter(),
		SumTrans:                    discard.NewGauge(),
		ConfirmedLevel:             discard.NewGauge(),
		LedgerSegments:                  discard.NewCounter(),
		ReplicatedLedgerSegment:          discard.NewCounter(),
		ReplicatedBallot:               discard.NewCounter(),
		PhasePeriodMoments:         discard.NewHistogram(),
		LedgerGossipSegmentsAccepted:    discard.NewCounter(),
		AssemblyPreballotDeferral:          discard.NewGauge(),
		AssemblyPreendorseDeferral:        discard.NewGauge(),
		CompletePreballotDeferral:            discard.NewGauge(),
		PreendorsementsTallied:           discard.NewGauge(),
		PreendorsementsStakingRatio: discard.NewGauge(),
		BallotAdditionAcceptTally:   discard.NewCounter(),
		NominationAcceptTally:        discard.NewCounter(),
		NominationInstantiateTally:         discard.NewCounter(),
		DurationPollingEnergyFraction:     discard.NewGauge(),
		TardyBallots:                   discard.NewCounter(),
		NodeLevel:                  discard.NewGauge(),
		DurationAugmentSum:         discard.NewCounter(),
	}
}
