package directives

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/examine"
	"github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/locator/record"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
)

//
var InspectCmd = &cobra.Command{
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

	RunE: runInspect,
}

func init() {
	InspectCmd.Flags().
		String("REDACTED",
			config.RPC.ListenAddress, "REDACTED")
	InspectCmd.Flags().
		String("REDACTED",
			config.DBBackend, "REDACTED")
	InspectCmd.Flags().
		String("REDACTED", config.DBPath, "REDACTED")
}

func runInspect(cmd *cobra.Command, _ []string) error {
	ctx, cancel := context.WithCancel(cmd.Context())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		cancel()
	}()

	blockStoreDB, err := cfg.DefaultDBProvider(&cfg.DBContext{ID: "REDACTED", Config: config})
	if err != nil {
		return err
	}
	blockStore := depot.NewBlockStore(blockStoreDB)
	defer blockStore.Close()

	stateDB, err := cfg.DefaultDBProvider(&cfg.DBContext{ID: "REDACTED", Config: config})
	if err != nil {
		return err
	}
	stateStore := status.NewStore(stateDB, status.StoreOptions{DiscardABCIResponses: false})
	defer stateStore.Close()

	genDoc, err := kinds.GenesisDocFromFile(config.GenesisFile())
	if err != nil {
		return err
	}
	txIndexer, blockIndexer, err := record.IndexerFromConfig(config, cfg.DefaultDBProvider, genDoc.ChainID)
	if err != nil {
		return err
	}
	ins := examine.New(config.RPC, blockStore, stateStore, txIndexer, blockIndexer)

	logger.Info("REDACTED")
	return ins.Run(ctx)
}
