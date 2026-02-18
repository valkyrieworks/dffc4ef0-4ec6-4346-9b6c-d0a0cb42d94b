package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/p2p"
)

//
var DisplayMemberUIDCommand = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    displayMemberUID,
}

func displayMemberUID(*cobra.Command, []string) error {
	memberKey, err := p2p.ImportMemberKey(settings.MemberKeyEntry())
	if err != nil {
		return err
	}

	fmt.Println(memberKey.ID())
	return nil
}
