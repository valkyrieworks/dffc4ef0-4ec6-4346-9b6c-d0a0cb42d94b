package diagnostics

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/utils/log"
)

var (
	nodeRPCAddr string
	profAddr    string
	frequency   uint

	flagNodeRPCAddr = "REDACTED"
	flagProfAddr    = "REDACTED"
	flagFrequency   = "REDACTED"

	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

//
//
var DebugCmd = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
}

func init() {
	DebugCmd.PersistentFlags().SortFlags = true
	DebugCmd.PersistentFlags().StringVar(
		&nodeRPCAddr,
		flagNodeRPCAddr,
		"REDACTED",
		"REDACTED",
	)

	DebugCmd.AddCommand(killCmd)
	DebugCmd.AddCommand(dumpCmd)
}
