package directives

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/scrutinize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/ledger"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
var ScrutinizeDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `
REDACTEDg
REDACTED.

REDACTEDe
REDACTED.
REDACTEDT
REDACTED.
REDACTED`,

	RunE: executeScrutinize,
}

func initialize() {
	ScrutinizeDirective.Flags().
		String("REDACTED",
			settings.RPC.OverhearLocation, "REDACTED")
	ScrutinizeDirective.Flags().
		String("REDACTED",
			settings.DatastoreOrigin, "REDACTED")
	ScrutinizeDirective.Flags().
		String("REDACTED", settings.DatastoreRoute, "REDACTED")
}

func executeScrutinize(cmd *cobra.Command, _ []string) error {
	ctx, abort := context.WithCancel(cmd.Context())
	defer abort()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		abort()
	}()

	ledgerDepotDatastore, err := cfg.FallbackDatastoreSupplier(&cfg.DatastoreScope{ID: "REDACTED", Settings: settings})
	if err != nil {
		return err
	}
	ledgerDepot := depot.FreshLedgerDepot(ledgerDepotDatastore)
	defer ledgerDepot.Shutdown()

	statusDatastore, err := cfg.FallbackDatastoreSupplier(&cfg.DatastoreScope{ID: "REDACTED", Settings: settings})
	if err != nil {
		return err
	}
	statusDepot := status.FreshDepot(statusDatastore, status.DepotChoices{EjectIfaceReplies: false})
	defer statusDepot.Shutdown()

	producePaper, err := kinds.InaugurationPaperOriginatingRecord(settings.InaugurationRecord())
	if err != nil {
		return err
	}
	transferOrdinalizer, ledgerOrdinalizer, err := ledger.OrdinalizerOriginatingSettings(settings, cfg.FallbackDatastoreSupplier, producePaper.SuccessionUUID)
	if err != nil {
		return err
	}
	ins := scrutinize.New(settings.RPC, ledgerDepot, statusDepot, transferOrdinalizer, ledgerOrdinalizer)

	tracer.Details("REDACTED")
	return ins.Run(ctx)
}
