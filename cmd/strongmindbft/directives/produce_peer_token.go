package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

//
//
var ProducePeerTokenDirective = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    producePeerToken,
}

func producePeerToken(*cobra.Command, []string) error {
	peerTokenRecord := settings.PeerTokenRecord()
	if strongos.RecordPresent(peerTokenRecord) {
		return fmt.Errorf("REDACTED", peerTokenRecord)
	}

	peerToken, err := p2p.FetchEitherProducePeerToken(peerTokenRecord)
	if err != nil {
		return err
	}
	fmt.Println(peerToken.ID())
	return nil
}
