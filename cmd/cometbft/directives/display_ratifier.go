package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	cometjson "github.com/valkyrieworks/utils/json"
	cometos "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/privatekey"
)

//
var DisplayRatifierCommand = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    displayRatifier,
}

func displayRatifier(*cobra.Command, []string) error {
	keyEntryRoute := settings.PrivateRatifierKeyEntry()
	if !cometos.EntryPresent(keyEntryRoute) {
		return fmt.Errorf("REDACTED", keyEntryRoute)
	}

	pv := privatekey.ImportEntryPrivatekey(keyEntryRoute, settings.PrivateRatifierStatusEntry())

	publicKey, err := pv.FetchPublicKey()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	bz, err := cometjson.Serialize(publicKey)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	fmt.Println(string(bz))
	return nil
}
