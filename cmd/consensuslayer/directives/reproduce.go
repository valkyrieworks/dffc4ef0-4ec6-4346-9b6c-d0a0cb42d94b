package directives

import (
	"github.com/spf13/cobra"

	"github.com/valkyrieworks/agreement"
)

//
var ReplayCmd = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Run: func(cmd *cobra.Command, args []string) {
		agreement.RunReplayFile(config.BaseConfig, config.Consensus, false)
	},
}

//
//
var ReplayConsoleCmd = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	Run: func(cmd *cobra.Command, args []string) {
		agreement.RunReplayFile(config.BaseConfig, config.Consensus, true)
	},
}
