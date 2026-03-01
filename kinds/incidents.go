package kinds

import (
	"fmt"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	tendermintpubsub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	tendermintinquire "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
)

//
const (
	//
	//
	//
	//
	//
	IncidentFreshLedger            = "REDACTED"
	IncidentFreshLedgerHeading      = "REDACTED"
	IncidentFreshLedgerIncidents      = "REDACTED"
	IncidentFreshProof         = "REDACTED"
	IncidentTransfer                  = "REDACTED"
	IncidentAssessorAssignRevisions = "REDACTED"

	//
	//
	//
	IncidentFinishedNomination   = "REDACTED"
	IncidentSecure               = "REDACTED"
	IncidentFreshIteration           = "REDACTED"
	IncidentFreshIterationPhase       = "REDACTED"
	IncidentSpeck              = "REDACTED"
	IncidentResecure             = "REDACTED"
	IncidentDeadlineNominate     = "REDACTED"
	IncidentDeadlinePause        = "REDACTED"
	IncidentRelease             = "REDACTED"
	IncidentSoundLedger         = "REDACTED"
	IncidentBallot               = "REDACTED"
	IncidentFreshAgreementParameters = "REDACTED"
)

//

//
type TEMPIncidentData interface {
	//
}

func initialize() {
	strongmindjson.EnrollKind(IncidentDataFreshLedger{}, "REDACTED")
	strongmindjson.EnrollKind(IncidentDataFreshLedgerHeading{}, "REDACTED")
	strongmindjson.EnrollKind(IncidentDataFreshLedgerIncidents{}, "REDACTED")
	strongmindjson.EnrollKind(IncidentDataFreshProof{}, "REDACTED")
	strongmindjson.EnrollKind(IncidentDataTransfer{}, "REDACTED")
	strongmindjson.EnrollKind(IncidentDataIterationStatus{}, "REDACTED")
	strongmindjson.EnrollKind(IncidentDataFreshIteration{}, "REDACTED")
	strongmindjson.EnrollKind(IncidentDataFinishNomination{}, "REDACTED")
	strongmindjson.EnrollKind(IncidentDataBallot{}, "REDACTED")
	strongmindjson.EnrollKind(IncidentDataAssessorAssignRevisions{}, "REDACTED")
	strongmindjson.EnrollKind(IncidentDataText("REDACTED"), "REDACTED")
}

//
//

type IncidentDataFreshLedger struct {
	Ledger               *Ledger                     `json:"ledger"`
	LedgerUUID             LedgerUUID                    `json:"ledger_uuid"`
	OutcomeCulminateLedger iface.ReplyCulminateLedger `json:"outcome_culminate_ledger"`
}

type IncidentDataFreshLedgerHeading struct {
	Heading Heading `json:"heading"`
}

type IncidentDataFreshLedgerIncidents struct {
	Altitude int64        `json:"altitude"`
	Incidents []iface.Incident `json:"incidents"`
	CountTrans int64        `json:"count_trans,string"` //
}

type IncidentDataFreshProof struct {
	Altitude   int64    `json:"altitude"`
	Proof Proof `json:"proof"`
}

//
type IncidentDataTransfer struct {
	iface.TransferOutcome
}

//
type IncidentDataIterationStatus struct {
	Altitude int64  `json:"altitude"`
	Iteration  int32  `json:"iteration"`
	Phase   string `json:"phase"`
}

type AssessorDetails struct {
	Location Location `json:"location"`
	Ordinal   int32   `json:"ordinal"`
}

type IncidentDataFreshIteration struct {
	Altitude int64  `json:"altitude"`
	Iteration  int32  `json:"iteration"`
	Phase   string `json:"phase"`

	Nominator AssessorDetails `json:"nominator"`
}

type IncidentDataFinishNomination struct {
	Altitude int64  `json:"altitude"`
	Iteration  int32  `json:"iteration"`
	Phase   string `json:"phase"`

	LedgerUUID LedgerUUID `json:"ledger_uuid"`
}

type IncidentDataBallot struct {
	Ballot *Ballot
}

type IncidentDataText string

type IncidentDataAssessorAssignRevisions struct {
	AssessorRevisions []*Assessor `json:"assessor_revisions"`
}

//

const (
	//
	IncidentKindToken = "REDACTED"

	//
	//
	TransferDigestToken = "REDACTED"

	//
	//
	TransferAltitudeToken = "REDACTED"

	//
	LedgerAltitudeToken = "REDACTED"
)

var (
	IncidentInquireFinishNomination    = InquireForeachIncident(IncidentFinishedNomination)
	IncidentInquireSecure                = InquireForeachIncident(IncidentSecure)
	IncidentInquireFreshLedger            = InquireForeachIncident(IncidentFreshLedger)
	IncidentInquireFreshLedgerHeading      = InquireForeachIncident(IncidentFreshLedgerHeading)
	IncidentInquireFreshLedgerIncidents      = InquireForeachIncident(IncidentFreshLedgerIncidents)
	IncidentInquireFreshProof         = InquireForeachIncident(IncidentFreshProof)
	IncidentInquireFreshIteration            = InquireForeachIncident(IncidentFreshIteration)
	IncidentInquireFreshIterationPhase        = InquireForeachIncident(IncidentFreshIterationPhase)
	IncidentInquireSpeck               = InquireForeachIncident(IncidentSpeck)
	IncidentInquireResecure              = InquireForeachIncident(IncidentResecure)
	IncidentInquireDeadlineNominate      = InquireForeachIncident(IncidentDeadlineNominate)
	IncidentInquireDeadlinePause         = InquireForeachIncident(IncidentDeadlinePause)
	IncidentInquireTransfer                  = InquireForeachIncident(IncidentTransfer)
	IncidentInquireRelease              = InquireForeachIncident(IncidentRelease)
	IncidentInquireAssessorAssignRevisions = InquireForeachIncident(IncidentAssessorAssignRevisions)
	IncidentInquireSoundLedger          = InquireForeachIncident(IncidentSoundLedger)
	IncidentInquireBallot                = InquireForeachIncident(IncidentBallot)
)

func IncidentInquireTransferForeach(tx Tx) tendermintpubsub.Inquire {
	return tendermintinquire.ShouldAssemble(fmt.Sprintf("REDACTED", IncidentKindToken, IncidentTransfer, TransferDigestToken, tx.Digest()))
}

func InquireForeachIncident(incidentKind string) tendermintpubsub.Inquire {
	return tendermintinquire.ShouldAssemble(fmt.Sprintf("REDACTED", IncidentKindToken, incidentKind))
}

//
type LedgerIncidentBroadcaster interface {
	BroadcastIncidentFreshLedger(ledger IncidentDataFreshLedger) error
	BroadcastIncidentFreshLedgerHeading(heading IncidentDataFreshLedgerHeading) error
	BroadcastIncidentFreshLedgerIncidents(incidents IncidentDataFreshLedgerIncidents) error
	BroadcastIncidentFreshProof(proof IncidentDataFreshProof) error
	BroadcastIncidentTransfer(IncidentDataTransfer) error
	BroadcastIncidentAssessorAssignRevisions(IncidentDataAssessorAssignRevisions) error
}

type TransferIncidentBroadcaster interface {
	BroadcastIncidentTransfer(IncidentDataTransfer) error
}
