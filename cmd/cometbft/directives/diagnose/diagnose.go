package diagnose

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/utils/log"
)

var (
	memberRPCAddress string
	profAddress    string
	recurrence   uint

	markMemberRPCAddress = "REDACTED"
	markProfAddress    = "REDACTED"
	markRecurrence   = "REDACTED"

	tracer = log.NewTMTracer(log.NewAlignRecorder(os.Stdout))
)

//
//
var DiagnoseCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
}

func init() {
	DiagnoseCommand.PersistentFlags().SortFlags = true
	DiagnoseCommand.PersistentFlags().StringVar(
		&memberRPCAddress,
		markMemberRPCAddress,
		"REDACTED",
		"REDACTED",
	)

	DiagnoseCommand.AddCommand(haltCommand)
	DiagnoseCommand.AddCommand(exportCommand)
}
