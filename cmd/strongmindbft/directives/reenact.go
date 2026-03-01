package directives

import (
	"github.com/spf13/cobra"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement"
)

//
var ReenactDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Run: func(cmd *cobra.Command, arguments []string) {
		agreement.ExecuteReenactRecord(settings.FoundationSettings, settings.Agreement, false)
	},
}

//
//
var ReenactTerminalDirective = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	Run: func(cmd *cobra.Command, arguments []string) {
		agreement.ExecuteReenactRecord(settings.FoundationSettings, settings.Agreement, true)
	},
}
