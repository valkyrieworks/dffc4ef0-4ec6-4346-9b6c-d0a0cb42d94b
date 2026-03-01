package directives

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
)

var discardLedger = false

func initialize() {
	RevertStatusDirective.Flags().BoolVar(&discardLedger, "REDACTED", false, "REDACTED")
}

var RevertStatusDirective = &cobra.Command{
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
	RunE: func(cmd *cobra.Command, arguments []string) error {
		altitude, digest, err := RevertStatus(settings, discardLedger)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}

		if discardLedger {
			fmt.Printf("REDACTED", altitude, digest)
		} else {
			fmt.Printf("REDACTED", altitude, digest)
		}

		return nil
	},
}

//
//
//
func RevertStatus(settings *cfg.Settings, discardLedger bool) (int64, []byte, error) {
	//
	ledgerDepot, statusDepot, err := fetchStatusAlsoLedgerDepot(settings)
	if err != nil {
		return -1, nil, err
	}
	defer func() {
		_ = ledgerDepot.Shutdown()
		_ = statusDepot.Shutdown()
	}()

	//
	return status.Revert(ledgerDepot, statusDepot, discardLedger)
}

func fetchStatusAlsoLedgerDepot(settings *cfg.Settings) (*depot.LedgerDepot, status.Depot, error) {
	datastoreKind := dbm.OriginKind(settings.DatastoreOrigin)

	if !os.RecordPresent(filepath.Join(settings.DatastorePath(), "REDACTED")) {
		return nil, nil, fmt.Errorf("REDACTED", settings.DatastorePath())
	}

	//
	ledgerDepotDatastore, err := dbm.FreshDatastore("REDACTED", datastoreKind, settings.DatastorePath())
	if err != nil {
		return nil, nil, err
	}
	ledgerDepot := depot.FreshLedgerDepot(ledgerDepotDatastore)

	if !os.RecordPresent(filepath.Join(settings.DatastorePath(), "REDACTED")) {
		return nil, nil, fmt.Errorf("REDACTED", settings.DatastorePath())
	}

	//
	statusDatastore, err := dbm.FreshDatastore("REDACTED", datastoreKind, settings.DatastorePath())
	if err != nil {
		return nil, nil, err
	}
	statusDepot := status.FreshDepot(statusDatastore, status.DepotChoices{
		EjectIfaceReplies: settings.Repository.EjectIfaceReplies,
	})

	return ledgerDepot, statusDepot, nil
}
