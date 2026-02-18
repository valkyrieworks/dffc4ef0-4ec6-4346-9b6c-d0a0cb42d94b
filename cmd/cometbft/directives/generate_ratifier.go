package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/privatekey"
)

//
//
var GenerateRatifierCommand = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	Run:     generateRatifier,
}

func generateRatifier(*cobra.Command, []string) {
	pv := privatekey.GenerateEntryPrivatekey("REDACTED", "REDACTED")
	jsbz, err := cometjson.Serialize(pv)
	if err != nil {
		panic(err)
	}
	fmt.Printf(`REDACTEDv
REDACTED`, string(jsbz))
}
