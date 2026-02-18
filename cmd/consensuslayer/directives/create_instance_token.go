package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	ctsystem "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/p2p"
)

//
//
var GenNodeKeyCmd = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    genNodeKey,
}

func genNodeKey(*cobra.Command, []string) error {
	nodeKeyFile := config.NodeKeyFile()
	if ctsystem.FileExists(nodeKeyFile) {
		return fmt.Errorf("REDACTED", nodeKeyFile)
	}

	nodeKey, err := p2p.LoadOrGenNodeKey(nodeKeyFile)
	if err != nil {
		return err
	}
	fmt.Println(nodeKey.ID())
	return nil
}
