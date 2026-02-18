package basetypes

import (
	"encoding/json"
	"time"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/p2p"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

//
type OutcomeLedgerchainDetails struct {
	FinalLevel int64              `json:"final_level"`
	LedgerMetadata []*kinds.LedgerMeta `json:"ledger_metadata"`
}

//
type OutcomeOrigin struct {
	Origin *kinds.OriginPaper `json:"origin"`
}

//
//
//
//
type OutcomeOriginSegment struct {
	SegmentAmount int    `json:"segment"`
	SumSegments int    `json:"sum"`
	Data        string `json:"data"`
}

//
type OutcomeLedger struct {
	LedgerUID kinds.LedgerUID `json:"ledger_uid"`
	Ledger   *kinds.Ledger  `json:"ledger"`
}

//
type OutcomeHeading struct {
	Heading *kinds.Heading `json:"heading"`
}

//
type OutcomeEndorse struct {
	kinds.AttestedHeading `json:"attested_heading"`
	NormativeEndorse    bool `json:"standard"`
}

//
type OutcomeLedgerOutcomes struct {
	Level                int64                     `json:"level"`
	TransOutcomes            []*iface.InvokeTransferOutcome      `json:"trans_outcomes"`
	CompleteLedgerEvents   []iface.Event              `json:"complete_ledger_events"`
	RatifierRefreshes      []iface.RatifierModify    `json:"ratifier_refreshes"`
	AgreementArgumentRefreshes *engineproto.AgreementOptions `json:"agreement_argument_refreshes"`
	ApplicationDigest               []byte                    `json:"application_digest"`
}

//
//
func NewOutcomeEndorse(heading *kinds.Heading, endorse *kinds.Endorse,
	standard bool,
) *OutcomeEndorse {
	return &OutcomeEndorse{
		AttestedHeading: kinds.AttestedHeading{
			Heading: heading,
			Endorse: endorse,
		},
		NormativeEndorse: standard,
	}
}

//
type AlignDetails struct {
	NewestLedgerDigest   octets.HexOctets `json:"newest_ledger_digest"`
	NewestApplicationDigest     octets.HexOctets `json:"newest_application_digest"`
	NewestLedgerLevel int64          `json:"newest_ledger_level"`
	NewestLedgerTime   time.Time      `json:"newest_ledger_time"`

	OldestLedgerDigest   octets.HexOctets `json:"oldest_ledger_digest"`
	OldestApplicationDigest     octets.HexOctets `json:"oldest_application_digest"`
	OldestLedgerLevel int64          `json:"oldest_ledger_level"`
	OldestLedgerTime   time.Time      `json:"oldest_ledger_time"`

	TrappingUp bool `json:"trapping_up"`
}

//
type RatifierDetails struct {
	Location     octets.HexOctets `json:"location"`
	PublicKey      vault.PublicKey  `json:"public_key"`
	PollingEnergy int64          `json:"polling_energy"`
}

//
type OutcomeState struct {
	MemberDetails      p2p.StandardMemberDetails `json:"member_details"`
	AlignDetails      AlignDetails            `json:"align_details"`
	RatifierDetails RatifierDetails       `json:"ratifier_details"`
}

//
func (s *OutcomeState) TransferOrdinalActivated() bool {
	if s == nil {
		return false
	}
	return s.MemberDetails.Another.TransOrdinal == "REDACTED"
}

//
type OutcomeNetDetails struct {
	Observing bool     `json:"observing"`
	Observers []string `json:"observers"`
	NNodes    int      `json:"n_nodes"`
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
	MemberDetails         p2p.StandardMemberDetails  `json:"member_details"`
	IsOutgoing       bool                 `json:"is_outgoing"`
	LinkageState p2p.LinkageState `json:"linkage_state"`
	DistantIP         string               `json:"distant_ip"`
}

//
type OutcomeRatifiers struct {
	LedgerLevel int64              `json:"ledger_level"`
	Ratifiers  []*kinds.Ratifier `json:"ratifiers"`
	//
	Number int `json:"tally"`
	//
	Sum int `json:"sum"`
}

//
type OutcomeAgreementOptions struct {
	LedgerLevel     int64                 `json:"ledger_level"`
	AgreementOptions kinds.AgreementOptions `json:"agreement_options"`
}

//
//
type OutcomeExportAgreementStatus struct {
	EpochStatus json.RawMessage `json:"duration_status"`
	Nodes      []NodeStatusDetails `json:"nodes"`
}

//
type NodeStatusDetails struct {
	MemberLocation string          `json:"member_location"`
	NodeStatus   json.RawMessage `json:"node_status"`
}

//
type OutcomeAgreementStatus struct {
	EpochStatus json.RawMessage `json:"duration_status"`
}

//
type OutcomeMulticastTransfer struct {
	Code      uint32         `json:"code"`
	Data      octets.HexOctets `json:"data"`
	Log       string         `json:"log"`
	Codex string         `json:"codex"`

	Digest octets.HexOctets `json:"digest"`
}

//
type OutcomeMulticastTransferEndorse struct {
	InspectTransfer  iface.ReplyInspectTransfer `json:"inspect_transfer"`
	TransOutcome iface.InvokeTransferOutcome    `json:"transfer_outcome"`
	Digest     octets.HexOctets       `json:"digest"`
	Level   int64                `json:"level"`
}

//
type OutcomeInspectTransfer struct {
	iface.ReplyInspectTransfer
}

//
type OutcomeTransfer struct {
	Digest     octets.HexOctets    `json:"digest"`
	Level   int64             `json:"level"`
	Ordinal    uint32            `json:"ordinal"`
	TransOutcome iface.InvokeTransferOutcome `json:"transfer_outcome"`
	Tx       kinds.Tx          `json:"tx"`
	Attestation    kinds.TransferEvidence     `json:"evidence,omitempty"`
}

//
type OutcomeTransferScan struct {
	Txs        []*OutcomeTransfer `json:"txs"`
	SumNumber int         `json:"sum_number"`
}

//
type OutcomeLedgerScan struct {
	Ledgers     []*OutcomeLedger `json:"ledgers"`
	SumNumber int            `json:"sum_number"`
}

//
type OutcomeUnattestedTrans struct {
	Number      int        `json:"n_trans"`
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
	OutcomeRiskyPurgeTxpool struct{}
	OutcomeRiskyBlueprint      struct{}
	OutcomeEnrol          struct{}
	OutcomeDeenroll        struct{}
	OutcomeVitality             struct{}
)

//
type OutcomeEvent struct {
	Inquire  string              `json:"inquire"`
	Data   kinds.TMEventData   `json:"data"`
	Events map[string][]string `json:"events"`
}
