package directives

import (
	"github.com/spf13/cobra"

	"github.com/valkyrieworks/agreement"
)

//
var ResimulateCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Run: func(cmd *cobra.Command, args []string) {
		agreement.ExecuteResimulateEntry(settings.RootSettings, settings.Agreement, false)
	},
}

//
//
var ResimulateTerminalCommand = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	Run: func(cmd *cobra.Command, args []string) {
		agreement.ExecuteResimulateEntry(settings.RootSettings, settings.Agreement, true)
	},
}
