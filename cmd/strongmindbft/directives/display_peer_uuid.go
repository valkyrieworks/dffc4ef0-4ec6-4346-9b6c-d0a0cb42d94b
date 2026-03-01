package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

//
var DisplayPeerUUIDDirective = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    displayPeerUUID,
}

func displayPeerUUID(*cobra.Command, []string) error {
	peerToken, err := p2p.FetchPeerToken(settings.PeerTokenRecord())
	if err != nil {
		return err
	}

	fmt.Println(peerToken.ID())
	return nil
}
