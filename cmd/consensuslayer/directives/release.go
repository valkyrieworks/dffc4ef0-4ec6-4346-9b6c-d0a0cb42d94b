package directives

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/release"
)

//
var VersionCmd = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Run: func(cmd *cobra.Command, args []string) {
		cmtVersion := release.TMCoreSemVer
		if release.TMGitCommitHash != "REDACTED" {
			cmtVersion += "REDACTED" + release.TMGitCommitHash
		}

		if verbose {
			values, _ := json.MarshalIndent(struct {
				CometBFT      string `json:"consensuslayer"`
				ABCI          string `json:"atci"`
				BlockProtocol uint64 `json:"record_standard"`
				P2PProtocol   uint64 `json:"mesh_standard"`
			}{
				CometBFT:      cmtVersion,
				ABCI:          release.ABCISemVer,
				BlockProtocol: release.BlockProtocol,
				P2PProtocol:   release.P2PProtocol,
			}, "REDACTED", "REDACTED")
			fmt.Println(string(values))
		} else {
			fmt.Println(cmtVersion)
		}
	},
}

func init() {
	VersionCmd.Flags().BoolVarP(&verbose, "REDACTED", "REDACTED", false, "REDACTED")
}
