//

package agreement

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
		Altitude: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		AssessorFinalAttestedAltitude: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		Cycles: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		IterationIntervalMoments: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, tags).With(tagsAlsoItems...),
		Assessors: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		AssessorsPotency: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		AssessorPotency: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		AssessorOmittedLedgers: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		AbsentAssessors: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		AbsentAssessorsPotency: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		TreacherousAssessors: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		TreacherousAssessorsPotency: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		LedgerDurationMoments: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
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
		LedgerExtentOctets: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		SuccessionExtentOctets: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
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
		RatifiedAltitude: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		LedgerFragments: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		ReplicatedLedgerFragment: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		ReplicatedBallot: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		PhaseIntervalMoments: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",

			Buckets: stdprometheus.ExponentialBucketsRange(0.1, 100, 8),
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		LedgerMulticastFragmentsAccepted: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		AssemblyPreballotDeferral: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		AssemblyPreendorseDeferral: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		CompletePreballotDeferral: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		PreendorsementsTallied: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		PreendorsementsPledgingFraction: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		BallotAdditionAcceptTally: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		NominationAcceptTally: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		NominationGenerateTally: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, tags).With(tagsAlsoItems...),
		IterationBallotingPotencyRatio: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		TardyBallots: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		NodeAltitude: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
		IterationAdvanceSum: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: scope,
			Subsystem: TelemetryComponent,
			Name:      "REDACTED",
			Help:      "REDACTED",
		}, append(tags, "REDACTED")).With(tagsAlsoItems...),
	}
}

func NooperationTelemetry() *Telemetry {
	return &Telemetry{
		Altitude:                      discard.NewGauge(),
		AssessorFinalAttestedAltitude:   discard.NewGauge(),
		Cycles:                      discard.NewGauge(),
		IterationIntervalMoments:        discard.NewHistogram(),
		Assessors:                  discard.NewGauge(),
		AssessorsPotency:             discard.NewGauge(),
		AssessorPotency:              discard.NewGauge(),
		AssessorOmittedLedgers:       discard.NewGauge(),
		AbsentAssessors:           discard.NewGauge(),
		AbsentAssessorsPotency:      discard.NewGauge(),
		TreacherousAssessors:         discard.NewGauge(),
		TreacherousAssessorsPotency:    discard.NewGauge(),
		LedgerDurationMoments:        discard.NewHistogram(),
		CountTrans:                      discard.NewGauge(),
		LedgerExtentOctets:              discard.NewGauge(),
		SuccessionExtentOctets:              discard.NewCounter(),
		SumTrans:                    discard.NewGauge(),
		RatifiedAltitude:             discard.NewGauge(),
		LedgerFragments:                  discard.NewCounter(),
		ReplicatedLedgerFragment:          discard.NewCounter(),
		ReplicatedBallot:               discard.NewCounter(),
		PhaseIntervalMoments:         discard.NewHistogram(),
		LedgerMulticastFragmentsAccepted:    discard.NewCounter(),
		AssemblyPreballotDeferral:          discard.NewGauge(),
		AssemblyPreendorseDeferral:        discard.NewGauge(),
		CompletePreballotDeferral:            discard.NewGauge(),
		PreendorsementsTallied:           discard.NewGauge(),
		PreendorsementsPledgingFraction: discard.NewGauge(),
		BallotAdditionAcceptTally:   discard.NewCounter(),
		NominationAcceptTally:        discard.NewCounter(),
		NominationGenerateTally:         discard.NewCounter(),
		IterationBallotingPotencyRatio:     discard.NewGauge(),
		TardyBallots:                   discard.NewCounter(),
		NodeAltitude:                  discard.NewGauge(),
		IterationAdvanceSum:         discard.NewCounter(),
	}
}
