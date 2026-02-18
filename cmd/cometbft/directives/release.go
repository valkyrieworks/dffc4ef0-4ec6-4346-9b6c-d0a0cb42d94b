package directives

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/release"
)

//
var ReleaseCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Run: func(cmd *cobra.Command, args []string) {
		cometRelease := release.TMCoreSemaphoreRev
		if release.TMGitEndorseDigest != "REDACTED" {
			cometRelease += "REDACTED" + release.TMGitEndorseDigest
		}

		if detailed {
			items, _ := json.MarshalIndent(struct {
				CometBFT      string `json:"cometbft"`
				Iface          string `json:"iface"`
				LedgerProtocol uint64 `json:"ledger_protocol"`
				P2PProtocol   uint64 `json:"p2p_protocol"`
			}{
				CometBFT:      cometRelease,
				Iface:          release.IfaceSemaphoreRev,
				LedgerProtocol: release.LedgerProtocol,
				P2PProtocol:   release.P2PProtocol,
			}, "REDACTED", "REDACTED")
			fmt.Println(string(items))
		} else {
			fmt.Println(cometRelease)
		}
	},
}

func init() {
	ReleaseCommand.Flags().BoolVarP(&detailed, "REDACTED", "REDACTED", false, "REDACTED")
}
