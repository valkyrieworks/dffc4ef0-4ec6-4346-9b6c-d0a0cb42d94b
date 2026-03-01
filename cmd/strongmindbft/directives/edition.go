package directives

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

//
var EditionDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Run: func(cmd *cobra.Command, arguments []string) {
		strongmindEdition := edition.TEMPBaseSemaphoreEdtn
		if edition.TEMPSourceEndorseDigest != "REDACTED" {
			strongmindEdition += "REDACTED" + edition.TEMPSourceEndorseDigest
		}

		if detailed {
			items, _ := json.MarshalIndent(struct {
				StrongSFT      string `json:"strongmindbft"`
				Iface          string `json:"iface"`
				LedgerScheme uint64 `json:"ledger_scheme"`
				Peer2peerScheme   uint64 `json:"peer2peer_scheme"`
			}{
				StrongSFT:      strongmindEdition,
				Iface:          edition.IfaceSemaphoreEdtn,
				LedgerScheme: edition.LedgerScheme,
				Peer2peerScheme:   edition.Peer2peerScheme,
			}, "REDACTED", "REDACTED")
			fmt.Println(string(items))
		} else {
			fmt.Println(strongmindEdition)
		}
	},
}

func initialize() {
	EditionDirective.Flags().BoolVarP(&detailed, "REDACTED", "REDACTED", false, "REDACTED")
}
