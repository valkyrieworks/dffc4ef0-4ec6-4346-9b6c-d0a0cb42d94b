package directives

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/privatekey"
)

func Verify_Restoreall(t *testing.T) {
	settings := cfg.VerifySettings()
	dir := t.TempDir()
	settings.AssignOrigin(dir)
	cfg.AssureOrigin(dir)
	require.NoError(t, initEntriesWithSettings(settings))
	pv := privatekey.ImportEntryPrivatekey(settings.PrivateRatifierKeyEntry(), settings.PrivateRatifierStatusEntry())
	pv.FinalAttestStatus.Level = 10
	pv.Persist()
	require.NoError(t, restoreAll(settings.StoreFolder(), settings.P2P.AddressLedgerEntry(), settings.PrivateRatifierKeyEntry(),
		settings.PrivateRatifierStatusEntry(), tracer))
	require.DirExists(t, settings.StoreFolder())
	require.NoFileExists(t, filepath.Join(settings.StoreFolder(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.StoreFolder(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.StoreFolder(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.StoreFolder(), "REDACTED"))
	require.FileExists(t, settings.PrivateRatifierStatusEntry())
	pv = privatekey.ImportEntryPrivatekey(settings.PrivateRatifierKeyEntry(), settings.PrivateRatifierStatusEntry())
	require.Equal(t, int64(0), pv.FinalAttestStatus.Level)
}

func Verify_Restorestatus(t *testing.T) {
	settings := cfg.VerifySettings()
	dir := t.TempDir()
	settings.AssignOrigin(dir)
	cfg.AssureOrigin(dir)
	require.NoError(t, initEntriesWithSettings(settings))
	pv := privatekey.ImportEntryPrivatekey(settings.PrivateRatifierKeyEntry(), settings.PrivateRatifierStatusEntry())
	pv.FinalAttestStatus.Level = 10
	pv.Persist()
	require.NoError(t, restoreStatus(settings.StoreFolder(), tracer))
	require.DirExists(t, settings.StoreFolder())
	require.NoFileExists(t, filepath.Join(settings.StoreFolder(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.StoreFolder(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.StoreFolder(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.StoreFolder(), "REDACTED"))
	require.FileExists(t, settings.PrivateRatifierStatusEntry())
	pv = privatekey.ImportEntryPrivatekey(settings.PrivateRatifierKeyEntry(), settings.PrivateRatifierStatusEntry())
	//
	require.Equal(t, int64(10), pv.FinalAttestStatus.Level)
}
