package e2e

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

//
type Declaration struct {
	//
	IDXPrv6 bool `toml:"inetv6"`

	//
	PrimaryAltitude int64 `toml:"primary_altitude"`

	//
	//
	PrimaryStatus map[string]string `toml:"primary_status"`

	//
	//
	//
	//
	//
	//
	//
	//
	//
	Assessors *map[string]int64 `toml:"assessors"`

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
	//
	//
	AssessorRevisions map[string]map[string]int64 `toml:"assessor_revise"`

	//
	Peers map[string]*DeclarationPeer `toml:"peer"`

	//
	//
	TokenKind string `toml:"token_kind"`

	//
	//
	Proof int `toml:"proof"`

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
	IfaceScheme string `toml:"iface_scheme"`

	//
	//
	ArrangeNominationDeferral time.Duration `toml:"arrange_nomination_deferral"`
	HandleNominationDeferral time.Duration `toml:"handle_nomination_deferral"`
	InspectTransferDeferral         time.Duration `toml:"inspect_transfer_deferral"`
	BallotAdditionDeferral   time.Duration `toml:"ballot_addition_deferral"`
	CulminateLedgerDeferral   time.Duration `toml:"culminate_ledger_deferral"`

	//
	//
	ModernizeEdition string `toml:"modernize_edition"`

	FetchTransferExtentOctets   int `toml:"fetch_transfer_extent_octets"`
	FetchTransferClusterExtent   int `toml:"fetch_transfer_cluster_extent"`
	FetchTransferLinkages int `toml:"fetch_transfer_linkages"`
	FetchMaximumTrans        int `toml:"fetch_maximum_trans"`

	//
	RecordStratum string `toml:"record_stratum"`

	//
	RecordLayout string `toml:"record_layout"`

	//
	//
	Titan bool `toml:"titan"`

	//
	//
	LedgerMaximumOctets int64 `toml:"ledger_maximum_octets"`

	//
	//
	//
	BallotAdditionsActivateAltitude int64 `toml:"ballot_additions_activate_altitude"`

	//
	//
	//
	//
	BallotAdditionsReviseAltitude int64 `toml:"ballot_additions_revise_altitude"`

	//
	BallotAdditionExtent uint `toml:"ballot_addition_extent"`

	//
	ExploratoryMaximumBroadcastLinkagesTowardEnduringNodes    uint `toml:"exploratory_maximum_broadcast_linkages_toward_enduring_nodes"`
	ExploratoryMaximumBroadcastLinkagesTowardUnEnduringNodes uint `toml:"exploratory_maximum_broadcast_linkages_toward_un_enduring_nodes"`
}

//
type DeclarationPeer struct {
	//
	//
	//
	Style string `toml:"style"`

	//
	//
	//
	//
	//
	Edition string `toml:"edition"`

	//
	Origins []string `toml:"origins"`

	//
	//
	//
	//
	EnduringNodes []string `toml:"enduring_nodes"`

	//
	//
	Repository string `toml:"repository"`

	//
	//
	//
	//
	PrivatevalueScheme string `toml:"privatevalue_scheme"`

	//
	//
	InitiateLocated int64 `toml:"initiate_located"`

	//
	//
	LedgerChronizeEdition string `toml:"ledger_chronize_edition"`

	//
	LedgerChronizeAggregateStyle bool `toml:"ledger_chronize_aggregate_style"`

	//
	//
	//
	//
	StatusChronize bool `toml:"status_chronize"`

	//
	//
	//
	EndureDuration *uint64 `toml:"endure_duration"`

	//
	//
	ImageDuration uint64 `toml:"image_duration"`

	//
	//
	//
	PreserveLedgers uint64 `toml:"preserve_ledgers"`

	//
	//
	//
	//
	//
	//
	//
	Disrupt []string `toml:"disrupt"`

	//
	//
	//
	TransmitNegativeFetch bool `toml:"transmit_negative_fetch"`

	//
	//
	//
	UtilizeLibpeer2peer bool `toml:"utilize_libpeer2peer"`

	//
	//
	TxpoolKind string `toml:"txpool_kind"`
}

//
func (m Declaration) Persist(record string) error {
	f, err := os.Create(record)
	if err != nil {
		return fmt.Errorf("REDACTED", record, err)
	}
	return toml.NewEncoder(f).Encode(m)
}

//
func FetchDeclaration(record string) (Declaration, error) {
	declaration := Declaration{}
	_, err := toml.DecodeFile(record, &declaration)
	if err != nil {
		return declaration, fmt.Errorf("REDACTED", record, err)
	}
	return declaration, nil
}
