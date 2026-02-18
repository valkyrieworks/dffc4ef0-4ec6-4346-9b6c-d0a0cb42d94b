package diagnose

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	cfg "github.com/valkyrieworks/settings"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
)

//
//
func exportState(rpc *rpchttp.HTTP, dir, filename string) error {
	state, err := rpc.Status(context.Background())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return recordStatusJSONToEntry(state, dir, filename)
}

//
//
func exportNetDetails(rpc *rpchttp.HTTP, dir, filename string) error {
	netDetails, err := rpc.NetDetails(context.Background())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return recordStatusJSONToEntry(netDetails, dir, filename)
}

//
//
func exportAgreementStatus(rpc *rpchttp.HTTP, dir, filename string) error {
	constExport, err := rpc.ExportAgreementStatus(context.Background())
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return recordStatusJSONToEntry(constExport, dir, filename)
}

//
//
func cloneJournal(cfg *cfg.Settings, dir string) error {
	journalRoute := cfg.Agreement.JournalEntry()
	journalEntry := filepath.Base(journalRoute)

	return cloneEntry(journalRoute, filepath.Join(dir, journalEntry))
}

//
//
func cloneSettings(home, dir string) error {
	settingsEntry := "REDACTED"
	settingsRoute := filepath.Join(home, "REDACTED", settingsEntry)

	return cloneEntry(settingsRoute, filepath.Join(dir, settingsEntry))
}

func exportBlueprint(dir, address, blueprint string, diagnose int) error {
	gateway := fmt.Sprintf("REDACTED", address, blueprint, diagnose)

	//
	reply, err := http.Get(gateway)
	if err != nil {
		return fmt.Errorf("REDACTED", blueprint, err)
	}
	defer reply.Body.Close()

	content, err := io.ReadAll(reply.Body)
	if err != nil {
		return fmt.Errorf("REDACTED", blueprint, err)
	}

	return os.WriteFile(path.Join(dir, fmt.Sprintf("REDACTED", blueprint)), content, 0o600)
}
