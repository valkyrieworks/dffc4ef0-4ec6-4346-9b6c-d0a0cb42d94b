package main

import (
	"os"
	"path/filepath"

	cmd "github.com/valkyrieworks/cmd/cometbft/directives"
	"github.com/valkyrieworks/cmd/cometbft/directives/diagnose"
	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/cli"
	nm "github.com/valkyrieworks/member"
)

func main() {
	originCommand := cmd.OriginCommand
	originCommand.AddCommand(
		cmd.GenerateRatifierCommand,
		cmd.InitEntriesCommand,
		cmd.RapidCommand,
		cmd.ResimulateCommand,
		cmd.ResimulateTerminalCommand,
		cmd.RestoreAllCommand,
		cmd.RestorePrivateRatifierCommand,
		cmd.RestoreStatusCommand,
		cmd.DisplayRatifierCommand,
		cmd.VerifychainEntriesCommand,
		cmd.DisplayMemberUIDCommand,
		cmd.ReOrdinalEventCommand,
		cmd.GenerateMemberKeyCommand,
		cmd.ReleaseCommand,
		cmd.RevertStatusCommand,
		cmd.CondenseGoLayerStoreCommand,
		cmd.ScrutinizeCommand,
		diagnose.DiagnoseCommand,
		cli.NewFinalizationCommand(originCommand, true),
	)

	//
	//
	//
	//
	//
	//
	//
	//
	memberFunction := nm.StandardNewMember

	//
	originCommand.AddCommand(cmd.NewExecuteMemberCommand(memberFunction))

	cmd := cli.ArrangeRootCommand(originCommand, "REDACTED", os.ExpandEnv(filepath.Join("REDACTED", cfg.StandardCometFolder)))
	if err := cmd.Perform(); err != nil {
		panic(err)
	}
}
