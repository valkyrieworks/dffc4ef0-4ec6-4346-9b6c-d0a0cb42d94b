package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
)

//
var DisplayAssessorDirective = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    displayAssessor,
}

func displayAssessor(*cobra.Command, []string) error {
	tokenRecordRoute := settings.PrivateAssessorTokenRecord()
	if !strongos.RecordPresent(tokenRecordRoute) {
		return fmt.Errorf("REDACTED", tokenRecordRoute)
	}

	pv := privatevalue.FetchRecordPRV(tokenRecordRoute, settings.PrivateAssessorStatusRecord())

	publicToken, err := pv.ObtainPublicToken()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	bz, err := strongmindjson.Serialize(publicToken)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	fmt.Println(string(bz))
	return nil
}
