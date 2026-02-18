package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	cometos "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/p2p"
)

//
//
var GenerateMemberKeyCommand = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    generateMemberKey,
}

func generateMemberKey(*cobra.Command, []string) error {
	memberKeyEntry := settings.MemberKeyEntry()
	if cometos.EntryPresent(memberKeyEntry) {
		return fmt.Errorf("REDACTED", memberKeyEntry)
	}

	memberKey, err := p2p.ImportOrGenerateMemberKey(memberKeyEntry)
	if err != nil {
		return err
	}
	fmt.Println(memberKey.ID())
	return nil
}
