package directives

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
)

func Verify_Restartevery(t *testing.T) {
	settings := cfg.VerifySettings()
	dir := t.TempDir()
	settings.AssignOrigin(dir)
	cfg.AssureOrigin(dir)
	require.NoError(t, initializeRecordsUsingSettings(settings))
	pv := privatevalue.FetchRecordPRV(settings.PrivateAssessorTokenRecord(), settings.PrivateAssessorStatusRecord())
	pv.FinalAttestStatus.Altitude = 10
	pv.Persist()
	require.NoError(t, restoreEvery(settings.DatastorePath(), settings.P2P.LocationRegisterRecord(), settings.PrivateAssessorTokenRecord(),
		settings.PrivateAssessorStatusRecord(), tracer))
	require.DirExists(t, settings.DatastorePath())
	require.NoFileExists(t, filepath.Join(settings.DatastorePath(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.DatastorePath(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.DatastorePath(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.DatastorePath(), "REDACTED"))
	require.FileExists(t, settings.PrivateAssessorStatusRecord())
	pv = privatevalue.FetchRecordPRV(settings.PrivateAssessorTokenRecord(), settings.PrivateAssessorStatusRecord())
	require.Equal(t, int64(0), pv.FinalAttestStatus.Altitude)
}

func Verify_Restartstatus(t *testing.T) {
	settings := cfg.VerifySettings()
	dir := t.TempDir()
	settings.AssignOrigin(dir)
	cfg.AssureOrigin(dir)
	require.NoError(t, initializeRecordsUsingSettings(settings))
	pv := privatevalue.FetchRecordPRV(settings.PrivateAssessorTokenRecord(), settings.PrivateAssessorStatusRecord())
	pv.FinalAttestStatus.Altitude = 10
	pv.Persist()
	require.NoError(t, restoreStatus(settings.DatastorePath(), tracer))
	require.DirExists(t, settings.DatastorePath())
	require.NoFileExists(t, filepath.Join(settings.DatastorePath(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.DatastorePath(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.DatastorePath(), "REDACTED"))
	require.NoFileExists(t, filepath.Join(settings.DatastorePath(), "REDACTED"))
	require.FileExists(t, settings.PrivateAssessorStatusRecord())
	pv = privatevalue.FetchRecordPRV(settings.PrivateAssessorTokenRecord(), settings.PrivateAssessorStatusRecord())
	//
	require.Equal(t, int64(10), pv.FinalAttestStatus.Altitude)
}
