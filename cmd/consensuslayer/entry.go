package main

import (
	"os"
	"path/filepath"

	cmd "github.com/valkyrieworks/cmd/consensuslayer/directives"
	"github.com/valkyrieworks/cmd/consensuslayer/directives/diagnostics"
	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/cli"
	nm "github.com/valkyrieworks/instance"
)

func main() {
	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(
		cmd.GenValidatorCmd,
		cmd.InitFilesCmd,
		cmd.LightCmd,
		cmd.ReplayCmd,
		cmd.ReplayConsoleCmd,
		cmd.ResetAllCmd,
		cmd.ResetPrivValidatorCmd,
		cmd.ResetStateCmd,
		cmd.ShowValidatorCmd,
		cmd.TestnetFilesCmd,
		cmd.ShowNodeIDCmd,
		cmd.ReIndexEventCmd,
		cmd.GenNodeKeyCmd,
		cmd.VersionCmd,
		cmd.RollbackStateCmd,
		cmd.CompactGoLevelDBCmd,
		cmd.InspectCmd,
		diagnostics.DebugCmd,
		cli.NewCompletionCmd(rootCmd, true),
	)

	//
	//
	//
	//
	//
	//
	//
	//
	nodeFunc := nm.DefaultNewNode

	//
	rootCmd.AddCommand(cmd.NewRunNodeCmd(nodeFunc))

	cmd := cli.PrepareBaseCmd(rootCmd, "REDACTED", os.ExpandEnv(filepath.Join("REDACTED", cfg.DefaultCometDir)))
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
