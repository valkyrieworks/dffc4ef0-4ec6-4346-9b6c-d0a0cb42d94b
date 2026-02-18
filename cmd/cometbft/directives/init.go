package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	cfg "github.com/valkyrieworks/settings"
	cometos "github.com/valkyrieworks/utils/os"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/privatekey"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

//
var InitEntriesCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	RunE:  initEntries,
}

func initEntries(*cobra.Command, []string) error {
	return initEntriesWithSettings(settings)
}

func initEntriesWithSettings(settings *cfg.Settings) error {
	//
	privateValueKeyEntry := settings.PrivateRatifierKeyEntry()
	privateValueStatusEntry := settings.PrivateRatifierStatusEntry()
	var pv *privatekey.EntryPV
	if cometos.EntryPresent(privateValueKeyEntry) {
		pv = privatekey.ImportEntryPrivatekey(privateValueKeyEntry, privateValueStatusEntry)
		tracer.Details("REDACTED", "REDACTED", privateValueKeyEntry,
			"REDACTED", privateValueStatusEntry)
	} else {
		pv = privatekey.GenerateEntryPrivatekey(privateValueKeyEntry, privateValueStatusEntry)
		pv.Persist()
		tracer.Details("REDACTED", "REDACTED", privateValueKeyEntry,
			"REDACTED", privateValueStatusEntry)
	}

	memberKeyEntry := settings.MemberKeyEntry()
	if cometos.EntryPresent(memberKeyEntry) {
		tracer.Details("REDACTED", "REDACTED", memberKeyEntry)
	} else {
		if _, err := p2p.ImportOrGenerateMemberKey(memberKeyEntry); err != nil {
			return err
		}
		tracer.Details("REDACTED", "REDACTED", memberKeyEntry)
	}

	//
	generateEntry := settings.OriginEntry()
	if cometos.EntryPresent(generateEntry) {
		tracer.Details("REDACTED", "REDACTED", generateEntry)
	} else {
		generatePaper := kinds.OriginPaper{
			LedgerUID:         fmt.Sprintf("REDACTED", engineseed.Str(6)),
			OriginMoment:     engineclock.Now(),
			AgreementOptions: kinds.StandardAgreementOptions(),
		}
		publicKey, err := pv.FetchPublicKey()
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		generatePaper.Ratifiers = []kinds.OriginRatifier{{
			Location: publicKey.Location(),
			PublicKey:  publicKey,
			Energy:   10,
		}}

		if err := generatePaper.PersistAs(generateEntry); err != nil {
			return err
		}
		tracer.Details("REDACTED", "REDACTED", generateEntry)
	}

	return nil
}
