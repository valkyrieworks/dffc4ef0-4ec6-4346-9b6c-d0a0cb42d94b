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
	IDXIpv6 bool `toml:"ipv6"`

	//
	PrimaryLevel int64 `toml:"primary_level"`

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
	Ratifiers *map[string]int64 `toml:"ratifiers"`

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
	RatifierRefreshes map[string]map[string]int64 `toml:"ratifier_modify"`

	//
	Instances map[string]*DeclarationMember `toml:"member"`

	//
	//
	KeyKind string `toml:"key_kind"`

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
	IfaceProtocol string `toml:"iface_protocol"`

	//
	//
	ArrangeNominationDeferral time.Duration `toml:"arrange_nomination_deferral"`
	HandleNominationDeferral time.Duration `toml:"handle_nomination_deferral"`
	InspectTransferDeferral         time.Duration `toml:"inspect_transfer_deferral"`
	BallotAdditionDeferral   time.Duration `toml:"ballot_addition_deferral"`
	CompleteLedgerDeferral   time.Duration `toml:"complete_ledger_deferral"`

	//
	//
	EnhanceRelease string `toml:"enhance_release"`

	ImportTransferVolumeOctets   int `toml:"import_transfer_volume_octets"`
	ImportTransferClusterVolume   int `toml:"import_transfer_cluster_volume"`
	ImportTransferLinkages int `toml:"import_transfer_linkages"`
	ImportMaximumTrans        int `toml:"import_maximum_trans"`

	//
	TraceLayer string `toml:"trace_layer"`

	//
	TraceLayout string `toml:"trace_layout"`

	//
	//
	Monitorstats bool `toml:"monitorstats"`

	//
	//
	LedgerMaximumOctets int64 `toml:"ledger_maximum_octets"`

	//
	//
	//
	BallotPluginsActivateLevel int64 `toml:"ballot_plugins_activate_level"`

	//
	//
	//
	//
	BallotPluginsModifyLevel int64 `toml:"ballot_plugins_modify_level"`

	//
	BallotAdditionVolume uint `toml:"ballot_addition_volume"`

	//
	ExploratoryMaximumGossipLinkagesToDurableNodes    uint `toml:"exploratory_maximum_gossip_linkages_to_durable_nodes"`
	ExploratoryMaximumGossipLinkagesToNotDurableNodes uint `toml:"exploratory_maximum_gossip_linkages_to_not_durable_nodes"`
}

//
type DeclarationMember struct {
	//
	//
	//
	Style string `toml:"style"`

	//
	//
	//
	//
	//
	Release string `toml:"release"`

	//
	Origins []string `toml:"origins"`

	//
	//
	//
	//
	DurableNodes []string `toml:"durable_nodes"`

	//
	//
	Datastore string `toml:"datastore"`

	//
	//
	//
	//
	PrivatekeyProtocol string `toml:"privatekey_protocol"`

	//
	//
	BeginAt int64 `toml:"begin_at"`

	//
	//
	LedgerAlignRelease string `toml:"ledger_align_release"`

	//
	//
	//
	//
	StatusAlign bool `toml:"status_align"`

	//
	//
	//
	EndureCadence *uint64 `toml:"endure_cadence"`

	//
	//
	MirrorCadence uint64 `toml:"mirror_cadence"`

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
	Vary []string `toml:"vary"`

	//
	//
	//
	TransmitNoImport bool `toml:"transmit_no_import"`

	//
	//
	//
	EmployLibp2p bool `toml:"employ_libp2p"`

	//
	//
	TxpoolKind string `toml:"txpool_kind"`
}

//
func (m Declaration) Persist(entry string) error {
	f, err := os.Create(entry)
	if err != nil {
		return fmt.Errorf("REDACTED", entry, err)
	}
	return toml.NewEncoder(f).Encode(m)
}

//
func ImportDeclaration(entry string) (Declaration, error) {
	declaration := Declaration{}
	_, err := toml.DecodeFile(entry, &declaration)
	if err != nil {
		return declaration, fmt.Errorf("REDACTED", entry, err)
	}
	return declaration, nil
}
