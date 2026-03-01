package param_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
)

func assureRecords(t *testing.T, originPath string, records ...string) {
	for _, f := range records {
		p := filepath.Join(originPath, f)
		_, err := os.Stat(p)
		assert.NoError(t, err, p)
	}
}

func VerifyAssureOrigin(t *testing.T) {
	demand := require.New(t)

	//
	scratchPath, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(err)
	defer os.RemoveAll(scratchPath)

	//
	settings.AssureOrigin(scratchPath)

	//
	data, err := os.ReadFile(filepath.Join(scratchPath, settings.FallbackSettingsPath, settings.FallbackSettingsRecordAlias))
	require.Nil(err)

	affirmSoundSettings(t, string(data))

	assureRecords(t, scratchPath, "REDACTED")
}

func VerifyAssureVerifyOrigin(t *testing.T) {
	demand := require.New(t)

	//
	cfg := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(cfg.OriginPath)
	originPath := cfg.OriginPath

	//
	data, err := os.ReadFile(filepath.Join(originPath, settings.FallbackSettingsPath, settings.FallbackSettingsRecordAlias))
	require.Nil(err)

	affirmSoundSettings(t, string(data))

	//
	foundationSettings := settings.FallbackFoundationSettings()
	assureRecords(t, originPath, settings.FallbackDataPath, foundationSettings.Inauguration, foundationSettings.PrivateAssessorToken, foundationSettings.PrivateAssessorStatus)
}

func affirmSoundSettings(t *testing.T, settingsRecord string) {
	t.Helper()
	//
	components := []string{
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
	}
	for _, e := range components {
		assert.Contains(t, settingsRecord, e)
	}
}
