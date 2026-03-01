package agreement

import (
	"strings"
	"time"

	"github.com/go-kit/kit/metrics"

	controlkinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement/kinds"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	//
	//
	TelemetryComponent = "REDACTED"
)

//

//
type Telemetry struct {
	//
	Altitude metrics.Gauge

	//
	AssessorFinalAttestedAltitude metrics.Gauge `metrics_labels:"assessor_location"`

	//
	Cycles metrics.Gauge

	//
	IterationIntervalMoments metrics.Histogram `metrics_buckettype:"experscope" metrics_bucketsizes:"0.1, 100, 8"`

	//
	Assessors metrics.Gauge
	//
	AssessorsPotency metrics.Gauge
	//
	AssessorPotency metrics.Gauge `metrics_labels:"assessor_location"`
	//
	AssessorOmittedLedgers metrics.Gauge `metrics_labels:"assessor_location"`
	//
	AbsentAssessors metrics.Gauge
	//
	AbsentAssessorsPotency metrics.Gauge
	//
	TreacherousAssessors metrics.Gauge
	//
	TreacherousAssessorsPotency metrics.Gauge

	//
	LedgerDurationMoments metrics.Histogram

	//
	CountTrans metrics.Gauge
	//
	LedgerExtentOctets metrics.Gauge
	//
	SuccessionExtentOctets metrics.Counter
	//
	SumTrans metrics.Gauge
	//
	RatifiedAltitude metrics.Gauge `metrics_name:"newest_ledger_altitude"`

	//
	LedgerFragments metrics.Counter `metrics_labels:"node_uuid"`

	//
	ReplicatedLedgerFragment metrics.Counter

	//
	ReplicatedBallot metrics.Counter

	//
	PhaseIntervalMoments metrics.Histogram `metrics_labels:"phase" metrics_buckettype:"experscope" metrics_bucketsizes:"0.1, 100, 8"`
	phaseInitiate           time.Time

	//
	//
	LedgerBroadcastFragmentsAccepted metrics.Counter `metrics_labels:"aligns_prevailing"`

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
	AssemblyPreballotDeferral metrics.Gauge `metrics_labels:"nominator_location"`

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
	AssemblyPreendorseDeferral metrics.Gauge `metrics_labels:"nominator_location"`

	//
	//
	//
	//
	CompletePreballotDeferral metrics.Gauge `metrics_labels:"nominator_location"`

	//
	PreendorsementsTallied metrics.Gauge

	//
	PreendorsementsPledgingFraction metrics.Gauge

	//
	//
	//
	BallotAdditionAcceptTally metrics.Counter `metrics_labels:"condition"`

	//
	//
	//
	//
	NominationAcceptTally metrics.Counter `metrics_labels:"condition"`

	//
	//
	//
	//
	NominationGenerateTally metrics.Counter

	//
	//
	//
	IterationBallotingPotencyRatio metrics.Gauge `metrics_labels:"ballot_kind"`

	//
	//
	//
	TardyBallots metrics.Counter `metrics_labels:"ballot_kind"`

	//
	//
	//
	NodeAltitude metrics.Gauge `metrics_labels:"node_uuid"`

	//
	//
	IterationAdvanceSum metrics.Counter `metrics_labels:"phase"`
}

func (m *Telemetry) LabelIterationAdvanced(phase controlkinds.IterationPhaseKind) {
	phaseAlias := strings.TrimPrefix(phase.Text(), "REDACTED")
	m.IterationAdvanceSum.With("REDACTED", phaseAlias).Add(1)
}

func (m *Telemetry) LabelNominationHandled(approved bool) {
	condition := "REDACTED"
	if !approved {
		condition = "REDACTED"
	}
	m.NominationAcceptTally.With("REDACTED", condition).Add(1)
}

func (m *Telemetry) LabelBallotAdditionAccepted(approved bool) {
	condition := "REDACTED"
	if !approved {
		condition = "REDACTED"
	}
	m.BallotAdditionAcceptTally.With("REDACTED", condition).Add(1)
}

func (m *Telemetry) LabelBallotAccepted(vt commitchema.AttestedSignalKind, potency, sumPotency int64) {
	p := float64(potency) / float64(sumPotency)
	n := kinds.AttestedSignalKindTowardBriefText(vt)
	m.IterationBallotingPotencyRatio.With("REDACTED", n).Add(p)
}

func (m *Telemetry) LabelIteration(r int32, st time.Time) {
	m.Cycles.Set(float64(r))
	iterationMoment := time.Since(st).Seconds()
	m.IterationIntervalMoments.Observe(iterationMoment)

	pvn := kinds.AttestedSignalKindTowardBriefText(commitchema.PreballotKind)
	m.IterationBallotingPotencyRatio.With("REDACTED", pvn).Set(0)

	pcn := kinds.AttestedSignalKindTowardBriefText(commitchema.PreendorseKind)
	m.IterationBallotingPotencyRatio.With("REDACTED", pcn).Set(0)
}

func (m *Telemetry) LabelTardyBallot(vt commitchema.AttestedSignalKind) {
	n := kinds.AttestedSignalKindTowardBriefText(vt)
	m.TardyBallots.With("REDACTED", n).Add(1)
}

func (m *Telemetry) LabelPhase(s controlkinds.IterationPhaseKind) {
	if !m.phaseInitiate.IsZero() {
		phaseMoment := time.Since(m.phaseInitiate).Seconds()
		phaseAlias := strings.TrimPrefix(s.Text(), "REDACTED")
		m.PhaseIntervalMoments.With("REDACTED", phaseAlias).Observe(phaseMoment)
	}
	m.phaseInitiate = time.Now()
}
