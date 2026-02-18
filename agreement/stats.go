package agreement

import (
	"strings"
	"time"

	"github.com/go-kit/kit/metrics"

	cskinds "github.com/valkyrieworks/agreement/kinds"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	kinds "github.com/valkyrieworks/kinds"
)

const (
	//
	//
	StatsComponent = "REDACTED"
)

//

//
type Stats struct {
	//
	Level metrics.Gauge

	//
	RatifierFinalAttestedLevel metrics.Gauge `metrics_labels:"ratifier_location"`

	//
	Iterations metrics.Gauge

	//
	DurationPeriodMoments metrics.Histogram `metrics_buckettype:"exprange" metrics_bucketsizes:"0.1, 100, 8"`

	//
	Ratifiers metrics.Gauge
	//
	RatifiersEnergy metrics.Gauge
	//
	RatifierEnergy metrics.Gauge `metrics_labels:"ratifier_location"`
	//
	RatifierSkippedLedgers metrics.Gauge `metrics_labels:"ratifier_location"`
	//
	AbsentRatifiers metrics.Gauge
	//
	AbsentRatifiersEnergy metrics.Gauge
	//
	FaultyRatifiers metrics.Gauge
	//
	FaultyRatifiersEnergy metrics.Gauge

	//
	LedgerCadenceMoments metrics.Histogram

	//
	CountTrans metrics.Gauge
	//
	LedgerVolumeOctets metrics.Gauge
	//
	SeriesVolumeOctets metrics.Counter
	//
	SumTrans metrics.Gauge
	//
	ConfirmedLevel metrics.Gauge `metrics_name:"newest_ledger_level"`

	//
	LedgerSegments metrics.Counter `metrics_labels:"node_uid"`

	//
	ReplicatedLedgerSegment metrics.Counter

	//
	ReplicatedBallot metrics.Counter

	//
	PhasePeriodMoments metrics.Histogram `metrics_labels:"phase" metrics_buckettype:"exprange" metrics_bucketsizes:"0.1, 100, 8"`
	phaseBegin           time.Time

	//
	//
	LedgerGossipSegmentsAccepted metrics.Counter `metrics_labels:"aligns_ongoing"`

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	AssemblyPreballotDeferral metrics.Gauge `metrics_labels:"recommender_location"`

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	AssemblyPreendorseDeferral metrics.Gauge `metrics_labels:"recommender_location"`

	//
	//
	//
	//
	CompletePreballotDeferral metrics.Gauge `metrics_labels:"recommender_location"`

	//
	PreendorsementsTallied metrics.Gauge

	//
	PreendorsementsStakingRatio metrics.Gauge

	//
	//
	//
	BallotAdditionAcceptTally metrics.Counter `metrics_labels:"state"`

	//
	//
	//
	//
	NominationAcceptTally metrics.Counter `metrics_labels:"state"`

	//
	//
	//
	//
	NominationInstantiateTally metrics.Counter

	//
	//
	//
	DurationPollingEnergyFraction metrics.Gauge `metrics_labels:"ballot_kind"`

	//
	//
	//
	TardyBallots metrics.Counter `metrics_labels:"ballot_kind"`

	//
	//
	//
	NodeLevel metrics.Gauge `metrics_labels:"node_uid"`

	//
	//
	DurationAugmentSum metrics.Counter `metrics_labels:"phase"`
}

func (m *Stats) StampDurationAugmented(phase cskinds.DurationPhaseKind) {
	phaseLabel := strings.TrimPrefix(phase.String(), "REDACTED")
	m.DurationAugmentSum.With("REDACTED", phaseLabel).Add(1)
}

func (m *Stats) StampNominationHandled(approved bool) {
	state := "REDACTED"
	if !approved {
		state = "REDACTED"
	}
	m.NominationAcceptTally.With("REDACTED", state).Add(1)
}

func (m *Stats) StampBallotAdditionAccepted(approved bool) {
	state := "REDACTED"
	if !approved {
		state = "REDACTED"
	}
	m.BallotAdditionAcceptTally.With("REDACTED", state).Add(1)
}

func (m *Stats) StampBallotAccepted(vt engineproto.AttestedMessageKind, energy, sumEnergy int64) {
	p := float64(energy) / float64(sumEnergy)
	n := kinds.AttestedMessageKindToBriefString(vt)
	m.DurationPollingEnergyFraction.With("REDACTED", n).Add(p)
}

func (m *Stats) StampDuration(r int32, st time.Time) {
	m.Iterations.Set(float64(r))
	durationTime := time.Since(st).Seconds()
	m.DurationPeriodMoments.Observe(durationTime)

	pvn := kinds.AttestedMessageKindToBriefString(engineproto.PreballotKind)
	m.DurationPollingEnergyFraction.With("REDACTED", pvn).Set(0)

	pcn := kinds.AttestedMessageKindToBriefString(engineproto.PreendorseKind)
	m.DurationPollingEnergyFraction.With("REDACTED", pcn).Set(0)
}

func (m *Stats) StampTardyBallot(vt engineproto.AttestedMessageKind) {
	n := kinds.AttestedMessageKindToBriefString(vt)
	m.TardyBallots.With("REDACTED", n).Add(1)
}

func (m *Stats) StampPhase(s cskinds.DurationPhaseKind) {
	if !m.phaseBegin.IsZero() {
		phaseTime := time.Since(m.phaseBegin).Seconds()
		phaseLabel := strings.TrimPrefix(s.String(), "REDACTED")
		m.PhasePeriodMoments.With("REDACTED", phaseLabel).Observe(phaseTime)
	}
	m.phaseBegin = time.Now()
}
