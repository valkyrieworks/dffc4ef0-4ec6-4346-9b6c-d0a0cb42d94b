package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
)

//
//
var ProduceAssessorDirective = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	Run:     produceAssessor,
}

func produceAssessor(*cobra.Command, []string) {
	pv := privatevalue.ProduceRecordPRV("REDACTED", "REDACTED")
	jsbytes, err := strongmindjson.Serialize(pv)
	if err != nil {
		panic(err)
	}
	fmt.Printf(`REDACTEDv
REDACTED`, string(jsbytes))
}
