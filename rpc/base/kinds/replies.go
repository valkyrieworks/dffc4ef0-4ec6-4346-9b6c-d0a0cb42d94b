package basetypes

import (
	"encoding/json"
	"time"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type OutcomeLedgerchainDetails struct {
	FinalAltitude int64              `json:"final_altitude"`
	LedgerMetadata []*kinds.LedgerSummary `json:"ledger_metadata"`
}

//
type OutcomeInauguration struct {
	Inauguration *kinds.OriginPaper `json:"inauguration"`
}

//
//
//
//
type OutcomeInaugurationSegment struct {
	SegmentNumeral int    `json:"segment"`
	SumSegments int    `json:"sum"`
	Data        string `json:"data"`
}

//
type OutcomeLedger struct {
	LedgerUUID kinds.LedgerUUID `json:"ledger_uuid"`
	Ledger   *kinds.Ledger  `json:"ledger"`
}

//
type OutcomeHeadline struct {
	Heading *kinds.Heading `json:"heading"`
}

//
type OutcomeEndorse struct {
	kinds.NotatedHeading `json:"notated_heading"`
	StandardEndorse    bool `json:"standard"`
}

//
type OutcomeLedgerOutcomes struct {
	Altitude                int64                     `json:"altitude"`
	TransOutcomes            []*iface.InvokeTransferOutcome      `json:"trans_outcomes"`
	CulminateLedgerIncidents   []iface.Incident              `json:"culminate_ledger_incidents"`
	AssessorRevisions      []iface.AssessorRevise    `json:"assessor_revisions"`
	AgreementArgumentRevisions *commitchema.AgreementSettings `json:"agreement_argument_revisions"`
	PlatformDigest               []byte                    `json:"application_digest"`
}

//
//
func FreshOutcomeEndorse(heading *kinds.Heading, endorse *kinds.Endorse,
	standard bool,
) *OutcomeEndorse {
	return &OutcomeEndorse{
		NotatedHeading: kinds.NotatedHeading{
			Heading: heading,
			Endorse: endorse,
		},
		StandardEndorse: standard,
	}
}

//
type ChronizeDetails struct {
	NewestLedgerDigest   octets.HexadecimalOctets `json:"newest_ledger_digest"`
	NewestApplicationDigest     octets.HexadecimalOctets `json:"newest_application_digest"`
	NewestLedgerAltitude int64          `json:"newest_ledger_altitude"`
	NewestLedgerMoment   time.Time      `json:"newest_ledger_moment"`

	InitialLedgerDigest   octets.HexadecimalOctets `json:"initial_ledger_digest"`
	InitialApplicationDigest     octets.HexadecimalOctets `json:"initial_application_digest"`
	InitialLedgerAltitude int64          `json:"initial_ledger_altitude"`
	InitialLedgerMoment   time.Time      `json:"initial_ledger_moment"`

	ObtainingAscend bool `json:"obtaining_ascend"`
}

//
type AssessorDetails struct {
	Location     octets.HexadecimalOctets `json:"location"`
	PublicToken      security.PublicToken  `json:"public_token"`
	BallotingPotency int64          `json:"balloting_potency"`
}

//
type OutcomeCondition struct {
	PeerDetails      p2p.FallbackPeerDetails `json:"peer_details"`
	ChronizeDetails      ChronizeDetails            `json:"chronize_details"`
	AssessorDetails AssessorDetails       `json:"assessor_details"`
}

//
func (s *OutcomeCondition) TransferPositionActivated() bool {
	if s == nil {
		return false
	}
	return s.PeerDetails.Another.TransferOrdinal == "REDACTED"
}

//
type OutcomeNetworkDetails struct {
	Observing bool     `json:"observing"`
	Observers []string `json:"observers"`
	NTHNodes    int      `json:"nth_nodes"`
	Nodes     []Node   `json:"nodes"`
}

//
type OutcomeCallOrigins struct {
	Log string `json:"log"`
}

//
type OutcomeCallNodes struct {
	Log string `json:"log"`
}

//
type Node struct {
	PeerDetails         p2p.FallbackPeerDetails  `json:"peer_details"`
	EqualsOutgoing       bool                 `json:"equals_outgoing"`
	LinkageCondition p2p.LinkageCondition `json:"linkage_condition"`
	DistantINET         string               `json:"distant_inet"`
}

//
type OutcomeAssessors struct {
	LedgerAltitude int64              `json:"ledger_altitude"`
	Assessors  []*kinds.Assessor `json:"assessors"`
	//
	Tally int `json:"tally"`
	//
	Sum int `json:"sum"`
}

//
type OutcomeAgreementParameters struct {
	LedgerAltitude     int64                 `json:"ledger_altitude"`
	AgreementSettings kinds.AgreementSettings `json:"agreement_parameters"`
}

//
//
type OutcomeExportAgreementStatus struct {
	IterationStatus json.RawMessage `json:"iteration_status"`
	Nodes      []NodeStatusDetails `json:"nodes"`
}

//
type NodeStatusDetails struct {
	PeerLocator string          `json:"peer_locator"`
	NodeStatus   json.RawMessage `json:"node_status"`
}

//
type OutcomeAgreementStatus struct {
	IterationStatus json.RawMessage `json:"iteration_status"`
}

//
type OutcomeMulticastTransfer struct {
	Cipher      uint32         `json:"cipher"`
	Data      octets.HexadecimalOctets `json:"data"`
	Log       string         `json:"log"`
	Codeset string         `json:"codeset"`

	Digest octets.HexadecimalOctets `json:"digest"`
}

//
type OutcomeMulticastTransferEndorse struct {
	InspectTransfer  iface.ReplyInspectTransfer `json:"inspect_transfer"`
	TransferOutcome iface.InvokeTransferOutcome    `json:"transfer_outcome"`
	Digest     octets.HexadecimalOctets       `json:"digest"`
	Altitude   int64                `json:"altitude"`
}

//
type OutcomeInspectTransfer struct {
	iface.ReplyInspectTransfer
}

//
type OutcomeTransfer struct {
	Digest     octets.HexadecimalOctets    `json:"digest"`
	Altitude   int64             `json:"altitude"`
	Ordinal    uint32            `json:"ordinal"`
	TransferOutcome iface.InvokeTransferOutcome `json:"transfer_outcome"`
	Tx       kinds.Tx          `json:"tx"`
	Attestation    kinds.TransferAttestation     `json:"attestation,omitempty"`
}

//
type OutcomeTransferLookup struct {
	Txs        []*OutcomeTransfer `json:"txs"`
	SumTally int         `json:"sum_tally"`
}

//
type OutcomeLedgerLookup struct {
	Ledgers     []*OutcomeLedger `json:"ledgers"`
	SumTally int            `json:"sum_tally"`
}

//
type OutcomePendingTrans struct {
	Tally      int        `json:"nth_trans"`
	Sum      int        `json:"sum"`
	SumOctets int64      `json:"sum_octets"`
	Txs        []kinds.Tx `json:"txs"`
}

//
type OutcomeIfaceDetails struct {
	Reply iface.ReplyDetails `json:"reply"`
}

//
type OutcomeIfaceInquire struct {
	Reply iface.ReplyInquire `json:"reply"`
}

//
type OutcomeMulticastProof struct {
	Digest []byte `json:"digest"`
}

//
type (
	OutcomeInsecurePurgeTxpool struct{}
	OutcomeInsecureAnalysis      struct{}
	OutcomeListen          struct{}
	OutcomeUnlisten        struct{}
	OutcomeVitality             struct{}
)

//
type OutcomeIncident struct {
	Inquire  string              `json:"inquire"`
	Data   kinds.TEMPIncidentData   `json:"data"`
	Incidents map[string][]string `json:"incidents"`
}
