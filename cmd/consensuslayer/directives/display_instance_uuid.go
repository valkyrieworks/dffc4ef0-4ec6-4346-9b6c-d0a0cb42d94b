package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/p2p"
)

//
var ShowNodeIDCmd = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    showNodeID,
}

func showNodeID(*cobra.Command, []string) error {
	nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
	if err != nil {
		return err
	}

	fmt.Println(nodeKey.ID())
	return nil
}
