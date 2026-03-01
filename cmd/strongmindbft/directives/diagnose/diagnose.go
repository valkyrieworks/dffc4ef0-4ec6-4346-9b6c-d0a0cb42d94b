package diagnose

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

var (
	peerRemoteLocation string
	analyzerLocation    string
	recurrence   uint

	markerPeerRemoteLocation = "REDACTED"
	markerAnalyzerLocation    = "REDACTED"
	markerRecurrence   = "REDACTED"

	tracer = log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))
)

//
//
var DiagnoseDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
}

func initialize() {
	DiagnoseDirective.PersistentFlags().SortFlags = true
	DiagnoseDirective.PersistentFlags().StringVar(
		&peerRemoteLocation,
		markerPeerRemoteLocation,
		"REDACTED",
		"REDACTED",
	)

	DiagnoseDirective.AddCommand(terminateDirective)
	DiagnoseDirective.AddCommand(exportDirective)
}
