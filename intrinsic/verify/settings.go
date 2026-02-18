package verify

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/valkyrieworks/settings"
	cometos "github.com/valkyrieworks/utils/os"
)

func RestoreVerifyOrigin(verifyLabel string) *settings.Settings {
	return RestoreVerifyRootWithSeriesUID(verifyLabel, "REDACTED")
}

func RestoreVerifyRootWithSeriesUID(verifyLabel string, ledgerUID string) *settings.Settings {
	//
	originFolder, err := os.MkdirTemp("REDACTED", fmt.Sprintf("REDACTED", ledgerUID, verifyLabel))
	if err != nil {
		panic(err)
	}

	settings.AssureOrigin(originFolder)

	rootSettings := settings.StandardRootSettings()
	originEntryRoute := filepath.Join(originFolder, rootSettings.Origin)
	privateKeyEntryRoute := filepath.Join(originFolder, rootSettings.PrivateRatifierKey)
	privateStatusEntryRoute := filepath.Join(originFolder, rootSettings.PrivateRatifierStatus)

	if !cometos.EntryPresent(originEntryRoute) {
		if ledgerUID == "REDACTED" {
			ledgerUID = StandardVerifyLedgerUID
		}
		verifyOrigin := fmt.Sprintf(verifyOriginFmt, ledgerUID)
		cometos.ShouldRecordEntry(originEntryRoute, []byte(verifyOrigin), 0o644)
	}
	//
	cometos.ShouldRecordEntry(privateKeyEntryRoute, []byte(verifyPrivateRatifierKey), 0o644)
	cometos.ShouldRecordEntry(privateStatusEntryRoute, []byte(verifyPrivateRatifierStatus), 0o644)

	settings := settings.VerifySettings().AssignOrigin(originFolder)
	return settings
}

var verifyOriginFmt = `REDACTED{
REDACTED,
REDACTED,
REDACTED,
REDACTED{
REDACTED{
REDACTED,
REDACTED,
REDACTED"
REDACTED,
REDACTED{
REDACTED,
REDACTED,
REDACTED"
REDACTED,
REDACTED{
REDACTED[
REDACTED"
REDACTED]
REDACTED,
REDACTED{
REDACTED"
REDACTED,
REDACTED}
REDACTED,
REDACTED[
REDACTED{
REDACTED{
REDACTED,
REDACTED"
REDACTED,
REDACTED,
REDACTED"
REDACTED}
REDACTED,
REDACTED"
REDACTED`

var verifyPrivateRatifierKey = `REDACTED{
REDACTED,
REDACTED{
REDACTED,
REDACTED"
REDACTED,
REDACTED{
REDACTED,
REDACTED"
REDACTED}
REDACTED`

var verifyPrivateRatifierStatus = `REDACTED{
REDACTED,
REDACTED,
REDACTED0
REDACTED`
