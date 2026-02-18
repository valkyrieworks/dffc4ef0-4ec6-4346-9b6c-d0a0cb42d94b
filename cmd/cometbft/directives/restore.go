package directives

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/utils/log"
	cometos "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/privatekey"
)

//
//
var RestoreAllCommand = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    restoreAllCommand,
}

var retainAddressLedger bool

//
var RestoreStatusCommand = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		settings, err = AnalyzeSettings(cmd)
		if err != nil {
			return err
		}

		return restoreStatus(settings.StoreFolder(), tracer)
	},
}

func init() {
	RestoreAllCommand.Flags().BoolVar(&retainAddressLedger, "REDACTED", false, "REDACTED")
}

//
var RestorePrivateRatifierCommand = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    restorePrivateRatifier,
}

//
//
func restoreAllCommand(cmd *cobra.Command, _ []string) (err error) {
	settings, err = AnalyzeSettings(cmd)
	if err != nil {
		return err
	}

	return restoreAll(
		settings.StoreFolder(),
		settings.P2P.AddressLedgerEntry(),
		settings.PrivateRatifierKeyEntry(),
		settings.PrivateRatifierStatusEntry(),
		tracer,
	)
}

//
//
func restorePrivateRatifier(cmd *cobra.Command, _ []string) (err error) {
	settings, err = AnalyzeSettings(cmd)
	if err != nil {
		return err
	}

	restoreEntryPrivatekey(settings.PrivateRatifierKeyEntry(), settings.PrivateRatifierStatusEntry(), tracer)
	return nil
}

//
func restoreAll(storeFolder, addressLedgerEntry, privateValueKeyEntry, privateValueStatusEntry string, tracer log.Tracer) error {
	if retainAddressLedger {
		tracer.Details("REDACTED")
	} else {
		deleteAddressLedger(addressLedgerEntry, tracer)
	}

	if err := os.RemoveAll(storeFolder); err == nil {
		tracer.Details("REDACTED", "REDACTED", storeFolder)
	} else {
		tracer.Fault("REDACTED", "REDACTED", storeFolder, "REDACTED", err)
	}

	if err := cometos.AssureFolder(storeFolder, 0o700); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
	}

	//
	restoreEntryPrivatekey(privateValueKeyEntry, privateValueStatusEntry, tracer)
	return nil
}

//
func restoreStatus(storeFolder string, tracer log.Tracer) error {
	ledgerstore := filepath.Join(storeFolder, "REDACTED")
	status := filepath.Join(storeFolder, "REDACTED")
	wal := filepath.Join(storeFolder, "REDACTED")
	proof := filepath.Join(storeFolder, "REDACTED")
	transOrdinal := filepath.Join(storeFolder, "REDACTED")

	if cometos.EntryPresent(ledgerstore) {
		if err := os.RemoveAll(ledgerstore); err == nil {
			tracer.Details("REDACTED", "REDACTED", ledgerstore)
		} else {
			tracer.Fault("REDACTED", "REDACTED", ledgerstore, "REDACTED", err)
		}
	}

	if cometos.EntryPresent(status) {
		if err := os.RemoveAll(status); err == nil {
			tracer.Details("REDACTED", "REDACTED", status)
		} else {
			tracer.Fault("REDACTED", "REDACTED", status, "REDACTED", err)
		}
	}

	if cometos.EntryPresent(wal) {
		if err := os.RemoveAll(wal); err == nil {
			tracer.Details("REDACTED", "REDACTED", wal)
		} else {
			tracer.Fault("REDACTED", "REDACTED", wal, "REDACTED", err)
		}
	}

	if cometos.EntryPresent(proof) {
		if err := os.RemoveAll(proof); err == nil {
			tracer.Details("REDACTED", "REDACTED", proof)
		} else {
			tracer.Fault("REDACTED", "REDACTED", proof, "REDACTED", err)
		}
	}

	if cometos.EntryPresent(transOrdinal) {
		if err := os.RemoveAll(transOrdinal); err == nil {
			tracer.Details("REDACTED", "REDACTED", transOrdinal)
		} else {
			tracer.Fault("REDACTED", "REDACTED", transOrdinal, "REDACTED", err)
		}
	}

	if err := cometos.AssureFolder(storeFolder, 0o700); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
	}
	return nil
}

func restoreEntryPrivatekey(privateValueKeyEntry, privateValueStatusEntry string, tracer log.Tracer) {
	if _, err := os.Stat(privateValueKeyEntry); err == nil {
		pv := privatekey.ImportEntryPrivatekeyEmptyStatus(privateValueKeyEntry, privateValueStatusEntry)
		pv.Restore()
		tracer.Details(
			"REDACTED",
			"REDACTED", privateValueKeyEntry,
			"REDACTED", privateValueStatusEntry,
		)
	} else {
		pv := privatekey.GenerateEntryPrivatekey(privateValueKeyEntry, privateValueStatusEntry)
		pv.Persist()
		tracer.Details(
			"REDACTED",
			"REDACTED", privateValueKeyEntry,
			"REDACTED", privateValueStatusEntry,
		)
	}
}

func deleteAddressLedger(addressLedgerEntry string, tracer log.Tracer) {
	if err := os.Remove(addressLedgerEntry); err == nil {
		tracer.Details("REDACTED", "REDACTED", addressLedgerEntry)
	} else if !os.IsNotExist(err) {
		tracer.Details("REDACTED", "REDACTED", addressLedgerEntry, "REDACTED", err)
	}
}
