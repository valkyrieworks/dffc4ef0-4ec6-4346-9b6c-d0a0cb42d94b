package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	abcijson "github.com/valkyrieworks/utils/text"
	ctsystem "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/authkey"
)

//
var ShowValidatorCmd = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    showValidator,
}

func showValidator(*cobra.Command, []string) error {
	keyFilePath := config.PrivValidatorKeyFile()
	if !ctsystem.FileExists(keyFilePath) {
		return fmt.Errorf("REDACTED", keyFilePath)
	}

	pv := authkey.LoadFilePV(keyFilePath, config.PrivValidatorStateFile())

	pubKey, err := pv.GetPubKey()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	bz, err := abcijson.Marshal(pubKey)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	fmt.Println(string(bz))
	return nil
}
