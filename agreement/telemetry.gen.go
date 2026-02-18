//

package agreement

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		Height: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		ValidatorLastSignedHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		Rounds: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		RoundDurationSeconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, labels).With(labelsAndValues...),
		Validators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		ValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		ValidatorPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		ValidatorMissedBlocks: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		MissingValidators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		MissingValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		ByzantineValidators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		ByzantineValidatorsPower: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		BlockIntervalSeconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		NumTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		BlockSizeBytes: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		ChainSizeBytes: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		TotalTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		CommittedHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		BlockParts: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		DuplicateBlockPart: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		DuplicateVote: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		StepDurationSeconds: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		BlockGossipPartsReceived: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		QuorumPrevoteDelay: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		QuorumPrecommitDelay: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		FullPrevoteDelay: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		PrecommitsCounted: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		PrecommitsStakingPercentage: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		VoteExtensionReceiveCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		ProposalReceiveCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		ProposalCreateCount: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, labels).With(labelsAndValues...),
		RoundVotingPowerPercent: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		LateVotes: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		PeerHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
		RoundIncrementTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(labels, "REDACTED")).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		Height:                      discard.NewGauge(),
		ValidatorLastSignedHeight:   discard.NewGauge(),
		Rounds:                      discard.NewGauge(),
		RoundDurationSeconds:        discard.NewHistogram(),
		Validators:                  discard.NewGauge(),
		ValidatorsPower:             discard.NewGauge(),
		ValidatorPower:              discard.NewGauge(),
		ValidatorMissedBlocks:       discard.NewGauge(),
		MissingValidators:           discard.NewGauge(),
		MissingValidatorsPower:      discard.NewGauge(),
		ByzantineValidators:         discard.NewGauge(),
		ByzantineValidatorsPower:    discard.NewGauge(),
		BlockIntervalSeconds:        discard.NewHistogram(),
		NumTxs:                      discard.NewGauge(),
		BlockSizeBytes:              discard.NewGauge(),
		ChainSizeBytes:              discard.NewCounter(),
		TotalTxs:                    discard.NewGauge(),
		CommittedHeight:             discard.NewGauge(),
		BlockParts:                  discard.NewCounter(),
		DuplicateBlockPart:          discard.NewCounter(),
		DuplicateVote:               discard.NewCounter(),
		StepDurationSeconds:         discard.NewHistogram(),
		BlockGossipPartsReceived:    discard.NewCounter(),
		QuorumPrevoteDelay:          discard.NewGauge(),
		QuorumPrecommitDelay:        discard.NewGauge(),
		FullPrevoteDelay:            discard.NewGauge(),
		PrecommitsCounted:           discard.NewGauge(),
		PrecommitsStakingPercentage: discard.NewGauge(),
		VoteExtensionReceiveCount:   discard.NewCounter(),
		ProposalReceiveCount:        discard.NewCounter(),
		ProposalCreateCount:         discard.NewCounter(),
		RoundVotingPowerPercent:     discard.NewGauge(),
		LateVotes:                   discard.NewCounter(),
		PeerHeight:                  discard.NewGauge(),
		RoundIncrementTotal:         discard.NewCounter(),
	}
}
