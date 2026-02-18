package settings_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/intrinsic/verify"
)

func assureEntries(t *testing.T, originFolder string, entries ...string) {
	for _, f := range entries {
		p := filepath.Join(originFolder, f)
		_, err := os.Stat(p)
		assert.NoError(t, err, p)
	}
}

func VerifyAssureRoot(t *testing.T) {
	demand := require.New(t)

	//
	tempFolder, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(err)
	defer os.RemoveAll(tempFolder)

	//
	settings.AssureOrigin(tempFolder)

	//
	data, err := os.ReadFile(filepath.Join(tempFolder, settings.StandardSettingsFolder, settings.StandardSettingsEntryLabel))
	require.Nil(err)

	affirmSoundSettings(t, string(data))

	assureEntries(t, tempFolder, "REDACTED")
}

func VerifyAssureVerifyRoot(t *testing.T) {
	demand := require.New(t)

	//
	cfg := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(cfg.OriginFolder)
	originFolder := cfg.OriginFolder

	//
	data, err := os.ReadFile(filepath.Join(originFolder, settings.StandardSettingsFolder, settings.StandardSettingsEntryLabel))
	require.Nil(err)

	affirmSoundSettings(t, string(data))

	//
	rootSettings := settings.StandardRootSettings()
	assureEntries(t, originFolder, settings.StandardDataFolder, rootSettings.Origin, rootSettings.PrivateRatifierKey, rootSettings.PrivateRatifierStatus)
}

func affirmSoundSettings(t *testing.T, settingsEntry string) {
	t.Helper()
	//
	elements := []string{
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
	for _, e := range elements {
		assert.Contains(t, settingsEntry, e)
	}
}
