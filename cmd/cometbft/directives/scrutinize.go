package directives

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/scrutinize"
	"github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/ordinaler/ledger"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
)

//
var ScrutinizeCommand = &cobra.Command{
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

func init() {
	ScrutinizeCommand.Flags().
		String("REDACTED",
			settings.RPC.AcceptLocation, "REDACTED")
	ScrutinizeCommand.Flags().
		String("REDACTED",
			settings.StoreOrigin, "REDACTED")
	ScrutinizeCommand.Flags().
		String("REDACTED", settings.StoreRoute, "REDACTED")
}

func executeScrutinize(cmd *cobra.Command, _ []string) error {
	ctx, revoke := context.WithCancel(cmd.Context())
	defer revoke()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		revoke()
	}()

	ledgerDepotStore, err := cfg.StandardStoreSource(&cfg.StoreContext{ID: "REDACTED", Settings: settings})
	if err != nil {
		return err
	}
	ledgerDepot := depot.NewLedgerDepot(ledgerDepotStore)
	defer ledgerDepot.End()

	statusStore, err := cfg.StandardStoreSource(&cfg.StoreContext{ID: "REDACTED", Settings: settings})
	if err != nil {
		return err
	}
	statusDepot := status.NewDepot(statusStore, status.DepotSettings{DropIfaceReplies: false})
	defer statusDepot.End()

	generatePaper, err := kinds.OriginPaperFromEntry(settings.OriginEntry())
	if err != nil {
		return err
	}
	transOrdinaler, ledgerOrdinaler, err := ledger.OrdinalerFromSettings(settings, cfg.StandardStoreSource, generatePaper.LedgerUID)
	if err != nil {
		return err
	}
	ins := scrutinize.New(settings.RPC, ledgerDepot, statusDepot, transOrdinaler, ledgerOrdinaler)

	tracer.Details("REDACTED")
	return ins.Run(ctx)
}
