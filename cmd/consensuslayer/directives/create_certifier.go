package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	abcijson "github.com/valkyrieworks/utils/text"
	"github.com/valkyrieworks/authkey"
)

//
//
var GenValidatorCmd = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	Run:     genValidator,
}

func genValidator(*cobra.Command, []string) {
	pv := authkey.GenFilePV("REDACTED", "REDACTED")
	jsbz, err := abcijson.Marshal(pv)
	if err != nil {
		panic(err)
	}
	fmt.Printf(`REDACTEDv
REDACTED`, string(jsbz))
}
