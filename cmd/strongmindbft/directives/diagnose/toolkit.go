package diagnose

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
)

//
//
func exportCondition(rpc *rpchttpsvc.Httpsvc, dir, recordname string) error {
	condition, err := rpc.Condition(context.Background())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return persistStatusJSNTowardRecord(condition, dir, recordname)
}

//
//
func exportNetworkDetails(rpc *rpchttpsvc.Httpsvc, dir, recordname string) error {
	networkDetails, err := rpc.NetworkDetails(context.Background())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return persistStatusJSNTowardRecord(networkDetails, dir, recordname)
}

//
//
func exportAgreementStatus(rpc *rpchttpsvc.Httpsvc, dir, recordname string) error {
	consensusExport, err := rpc.ExportAgreementStatus(context.Background())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return persistStatusJSNTowardRecord(consensusExport, dir, recordname)
}

//
//
func duplicateJournal(setting *cfg.Settings, dir string) error {
	journalRoute := setting.Agreement.JournalRecord()
	journalRecord := filepath.Base(journalRoute)

	return duplicateRecord(journalRoute, filepath.Join(dir, journalRecord))
}

//
//
func duplicateSettings(domain, dir string) error {
	settingsRecord := "REDACTED"
	settingsRoute := filepath.Join(domain, "REDACTED", settingsRecord)

	return duplicateRecord(settingsRoute, filepath.Join(dir, settingsRecord))
}

func exportAnalysis(dir, location, analysis string, diagnose int) error {
	gateway := fmt.Sprintf("REDACTED", location, analysis, diagnose)

	//
	reply, err := http.Get(gateway)
	if err != nil {
		return fmt.Errorf("REDACTED", analysis, err)
	}
	defer reply.Body.Close()

	content, err := io.ReadAll(reply.Body)
	if err != nil {
		return fmt.Errorf("REDACTED", analysis, err)
	}

	return os.WriteFile(path.Join(dir, fmt.Sprintf("REDACTED", analysis)), content, 0o600)
}
