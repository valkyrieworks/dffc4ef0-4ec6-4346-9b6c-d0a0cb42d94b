package primary

import (
	"os"
	"path/filepath"

	cmd "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/cmd/strongmindbft/directives"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/cmd/strongmindbft/directives/diagnose"
	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/cli"
	nm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/peer"
)

func primary() {
	originDirective := cmd.OriginDirective
	originDirective.AddCommand(
		cmd.ProduceAssessorDirective,
		cmd.InitializeRecordsDirective,
		cmd.AgileDirective,
		cmd.ReenactDirective,
		cmd.ReenactTerminalDirective,
		cmd.RestoreEveryDirective,
		cmd.RestorePrivateAssessorDirective,
		cmd.RestoreStatusDirective,
		cmd.DisplayAssessorDirective,
		cmd.SimnetRecordsDirective,
		cmd.DisplayPeerUUIDDirective,
		cmd.AgainOrdinalIncidentDirective,
		cmd.ProducePeerTokenDirective,
		cmd.EditionDirective,
		cmd.RevertStatusDirective,
		cmd.CompressProceedStratumDatastoreDirective,
		cmd.ScrutinizeDirective,
		diagnose.DiagnoseDirective,
		cli.FreshFinalizationDirective(originDirective, true),
	)

	//
	//
	//
	//
	//
	//
	//
	//
	peerMethod := nm.FallbackFreshPeer

	//
	originDirective.AddCommand(cmd.FreshExecutePeerDirective(peerMethod))

	cmd := cli.ArrangeFoundationDirective(originDirective, "REDACTED", os.ExpandEnv(filepath.Join("REDACTED", cfg.FallbackStrongPath)))
	if err := cmd.Perform(); err != nil {
		panic(err)
	}
}
