package agreement

import (
	"strings"
	"time"

	"github.com/go-kit/kit/metrics"

	statetypes "github.com/valkyrieworks/agreement/kinds"
	ctschema "github.com/valkyrieworks/schema/consensuscore/kinds"
	kinds "github.com/valkyrieworks/kinds"
)

const (
	//
	//
	MetricsSubsystem = "REDACTED"
)

//

//
type Metrics struct {
	//
	Height metrics.Gauge

	//
	ValidatorLastSignedHeight metrics.Gauge `metrics_labels:"certifier_location"`

	//
	Rounds metrics.Gauge

	//
	RoundDurationSeconds metrics.Histogram `metrics_buckettype:"boundrange" metrics_bucketsizes:"0.1, 100, 8"`

	//
	Validators metrics.Gauge
	//
	ValidatorsPower metrics.Gauge
	//
	ValidatorPower metrics.Gauge `metrics_labels:"certifier_location"`
	//
	ValidatorMissedBlocks metrics.Gauge `metrics_labels:"certifier_location"`
	//
	MissingValidators metrics.Gauge
	//
	MissingValidatorsPower metrics.Gauge
	//
	ByzantineValidators metrics.Gauge
	//
	ByzantineValidatorsPower metrics.Gauge

	//
	BlockIntervalSeconds metrics.Histogram

	//
	NumTxs metrics.Gauge
	//
	BlockSizeBytes metrics.Gauge
	//
	ChainSizeBytes metrics.Counter
	//
	TotalTxs metrics.Gauge
	//
	CommittedHeight metrics.Gauge `metrics_name:"newest_record_altitude"`

	//
	BlockParts metrics.Counter `metrics_labels:"node_uuid"`

	//
	DuplicateBlockPart metrics.Counter

	//
	DuplicateVote metrics.Counter

	//
	StepDurationSeconds metrics.Histogram `metrics_labels:"phase" metrics_buckettype:"boundrange" metrics_bucketsizes:"0.1, 100, 8"`
	stepStart           time.Time

	//
	//
	BlockGossipPartsReceived metrics.Counter `metrics_labels:"aligns_present"`

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
	QuorumPrevoteDelay metrics.Gauge `metrics_labels:"nominator_location"`

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
	QuorumPrecommitDelay metrics.Gauge `metrics_labels:"nominator_location"`

	//
	//
	//
	//
	FullPrevoteDelay metrics.Gauge `metrics_labels:"nominator_location"`

	//
	PrecommitsCounted metrics.Gauge

	//
	PrecommitsStakingPercentage metrics.Gauge

	//
	//
	//
	VoteExtensionReceiveCount metrics.Counter `metrics_labels:"condition"`

	//
	//
	//
	//
	ProposalReceiveCount metrics.Counter `metrics_labels:"condition"`

	//
	//
	//
	//
	ProposalCreateCount metrics.Counter

	//
	//
	//
	RoundVotingPowerPercent metrics.Gauge `metrics_labels:"ballot_kind"`

	//
	//
	//
	LateVotes metrics.Counter `metrics_labels:"ballot_kind"`

	//
	//
	//
	PeerHeight metrics.Gauge `metrics_labels:"node_uuid"`

	//
	//
	RoundIncrementTotal metrics.Counter `metrics_labels:"phase"`
}

func (m *Metrics) MarkRoundIncremented(step statetypes.RoundStepType) {
	stepName := strings.TrimPrefix(step.String(), "REDACTED")
	m.RoundIncrementTotal.With("REDACTED", stepName).Add(1)
}

func (m *Metrics) MarkProposalProcessed(accepted bool) {
	status := "REDACTED"
	if !accepted {
		status = "REDACTED"
	}
	m.ProposalReceiveCount.With("REDACTED", status).Add(1)
}

func (m *Metrics) MarkVoteExtensionReceived(accepted bool) {
	status := "REDACTED"
	if !accepted {
		status = "REDACTED"
	}
	m.VoteExtensionReceiveCount.With("REDACTED", status).Add(1)
}

func (m *Metrics) MarkVoteReceived(vt ctschema.SignedMsgType, power, totalPower int64) {
	p := float64(power) / float64(totalPower)
	n := kinds.SignedMsgTypeToShortString(vt)
	m.RoundVotingPowerPercent.With("REDACTED", n).Add(p)
}

func (m *Metrics) MarkRound(r int32, st time.Time) {
	m.Rounds.Set(float64(r))
	roundTime := time.Since(st).Seconds()
	m.RoundDurationSeconds.Observe(roundTime)

	pvn := kinds.SignedMsgTypeToShortString(ctschema.PrevoteType)
	m.RoundVotingPowerPercent.With("REDACTED", pvn).Set(0)

	pcn := kinds.SignedMsgTypeToShortString(ctschema.PrecommitType)
	m.RoundVotingPowerPercent.With("REDACTED", pcn).Set(0)
}

func (m *Metrics) MarkLateVote(vt ctschema.SignedMsgType) {
	n := kinds.SignedMsgTypeToShortString(vt)
	m.LateVotes.With("REDACTED", n).Add(1)
}

func (m *Metrics) MarkStep(s statetypes.RoundStepType) {
	if !m.stepStart.IsZero() {
		stepTime := time.Since(m.stepStart).Seconds()
		stepName := strings.TrimPrefix(s.String(), "REDACTED")
		m.StepDurationSeconds.With("REDACTED", stepName).Observe(stepTime)
	}
	m.stepStart = time.Now()
}
