package directives

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	dbm "github.com/valkyrieworks/-db"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
)

var deleteLedger = false

func init() {
	RevertStatusCommand.Flags().BoolVar(&deleteLedger, "REDACTED", false, "REDACTED")
}

var RevertStatusCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `
REDACTED,
REDACTEDe
REDACTED.
REDACTED 
REDACTED 
REDACTEDn
REDACTED.
REDACTED`,
	RunE: func(cmd *cobra.Command, args []string) error {
		level, digest, err := RevertStatus(settings, deleteLedger)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}

		if deleteLedger {
			fmt.Printf("REDACTED", level, digest)
		} else {
			fmt.Printf("REDACTED", level, digest)
		}

		return nil
	},
}

//
//
//
func RevertStatus(settings *cfg.Settings, deleteLedger bool) (int64, []byte, error) {
	//
	ledgerDepot, statusDepot, err := importStatusAndLedgerDepot(settings)
	if err != nil {
		return -1, nil, err
	}
	defer func() {
		_ = ledgerDepot.End()
		_ = statusDepot.End()
	}()

	//
	return status.Revert(ledgerDepot, statusDepot, deleteLedger)
}

func importStatusAndLedgerDepot(settings *cfg.Settings) (*depot.LedgerDepot, status.Depot, error) {
	storeKind := dbm.OriginKind(settings.StoreOrigin)

	if !os.EntryPresent(filepath.Join(settings.StoreFolder(), "REDACTED")) {
		return nil, nil, fmt.Errorf("REDACTED", settings.StoreFolder())
	}

	//
	ledgerDepotStore, err := dbm.NewStore("REDACTED", storeKind, settings.StoreFolder())
	if err != nil {
		return nil, nil, err
	}
	ledgerDepot := depot.NewLedgerDepot(ledgerDepotStore)

	if !os.EntryPresent(filepath.Join(settings.StoreFolder(), "REDACTED")) {
		return nil, nil, fmt.Errorf("REDACTED", settings.StoreFolder())
	}

	//
	statusStore, err := dbm.NewStore("REDACTED", storeKind, settings.StoreFolder())
	if err != nil {
		return nil, nil, err
	}
	statusDepot := status.NewDepot(statusStore, status.DepotSettings{
		DropIfaceReplies: settings.Archival.DropIfaceReplies,
	})

	return ledgerDepot, statusDepot, nil
}
