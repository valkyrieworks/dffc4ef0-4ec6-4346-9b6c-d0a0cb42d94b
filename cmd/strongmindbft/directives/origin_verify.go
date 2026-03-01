package directives

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/cli"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
)

//
func flushSettings(t *testing.T, dir string) {
	os.Clearenv()
	err := os.RemoveAll(dir)
	require.NoError(t, err)

	viper.Reset()
	settings = cfg.FallbackSettings()
}

//
func verifyOriginDirective() *cobra.Command {
	originDirective := &cobra.Command{
		Use:               OriginDirective.Use,
		PersistentPreRunE: OriginDirective.PersistentPreRunE,
		Run:               func(cmd *cobra.Command, arguments []string) {},
	}
	enrollSwitchesOriginDirective(originDirective)
	var l string
	originDirective.PersistentFlags().String("REDACTED", l, "REDACTED")
	return originDirective
}

func verifyConfigure(t *testing.T, origin string, arguments []string, env map[string]string) error {
	flushSettings(t, origin)

	originDirective := verifyOriginDirective()
	cmd := cli.ArrangeFoundationDirective(originDirective, "REDACTED", origin)

	//
	arguments = append([]string{originDirective.Use}, arguments...)
	return cli.ExecuteUsingArguments(cmd, arguments, env)
}

func VerifyOriginDomain(t *testing.T) {
	scratchPath := os.TempDir()
	origin := filepath.Join(scratchPath, "REDACTED")
	freshOrigin := filepath.Join(scratchPath, "REDACTED")
	defer flushSettings(t, origin)
	defer flushSettings(t, freshOrigin)

	scenarios := []struct {
		arguments []string
		env  map[string]string
		origin string
	}{
		{nil, nil, origin},
		{[]string{"REDACTED", freshOrigin}, nil, freshOrigin},
		{nil, map[string]string{"REDACTED": freshOrigin}, freshOrigin}, //
		{nil, map[string]string{"REDACTED": freshOrigin}, freshOrigin},
	}

	for i, tc := range scenarios {
		offsetText := "REDACTED" + strconv.Itoa(i)

		err := verifyConfigure(t, origin, tc.arguments, tc.env)
		require.Nil(t, err, offsetText)

		assert.Equal(t, tc.origin, settings.OriginPath, offsetText)
		assert.Equal(t, tc.origin, settings.P2P.OriginPath, offsetText)
		assert.Equal(t, tc.origin, settings.Agreement.OriginPath, offsetText)
		assert.Equal(t, tc.origin, settings.Txpool.OriginPath, offsetText)
	}
}

func VerifyOriginSwitchesContext(t *testing.T) {
	//
	preset := cfg.FallbackSettings()
	fallbackRecordLayer := preset.RecordStratum

	scenarios := []struct {
		arguments     []string
		env      map[string]string
		recordStratum string
	}{
		{[]string{"REDACTED", "REDACTED"}, nil, fallbackRecordLayer},                  //
		{[]string{"REDACTED", "REDACTED"}, nil, "REDACTED"},                  //
		{nil, map[string]string{"REDACTED": "REDACTED"}, fallbackRecordLayer},        //
		{nil, map[string]string{"REDACTED": "REDACTED"}, fallbackRecordLayer},  //
		{nil, map[string]string{"REDACTED": "REDACTED"}, fallbackRecordLayer},  //
		{nil, map[string]string{"REDACTED": "REDACTED"}, fallbackRecordLayer},       //
		{nil, map[string]string{"REDACTED": "REDACTED"}, fallbackRecordLayer}, //
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED"},       //
	}

	for i, tc := range scenarios {
		offsetText := strconv.Itoa(i)
		origin := filepath.Join(os.TempDir(), "REDACTED"+offsetText)
		offsetText = "REDACTED" + offsetText
		defer flushSettings(t, origin)
		err := verifyConfigure(t, origin, tc.arguments, tc.env)
		require.Nil(t, err, offsetText)

		assert.Equal(t, tc.recordStratum, settings.RecordStratum, offsetText)
	}
}

func VerifyOriginSettings(t *testing.T) {
	//
	unFallbackRecordLayer := "REDACTED"
	strongmindvalues := map[string]string{
		"REDACTED": unFallbackRecordLayer,
	}

	scenarios := []struct {
		arguments []string
		env  map[string]string

		recordLayer string
	}{
		{nil, nil, unFallbackRecordLayer},                                           //
		{[]string{"REDACTED"}, nil, "REDACTED"},                    //
		{nil, map[string]string{"REDACTED": "REDACTED"}, unFallbackRecordLayer}, //
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED"},      //
	}

	for i, tc := range scenarios {
		offsetText := strconv.Itoa(i)
		origin := filepath.Join(os.TempDir(), "REDACTED"+offsetText)
		offsetText = "REDACTED" + offsetText
		defer flushSettings(t, origin)
		//
		settingsRecordRoute := filepath.Join(origin, "REDACTED")
		err := strongos.AssurePath(settingsRecordRoute, 0o700)
		require.Nil(t, err)

		//
		//
		err = PersistSettingsValues(settingsRecordRoute, strongmindvalues)
		require.Nil(t, err)

		originDirective := verifyOriginDirective()
		cmd := cli.ArrangeFoundationDirective(originDirective, "REDACTED", origin)

		//
		tc.arguments = append([]string{originDirective.Use}, tc.arguments...)
		err = cli.ExecuteUsingArguments(cmd, tc.arguments, tc.env)
		require.Nil(t, err, offsetText)

		assert.Equal(t, tc.recordLayer, settings.RecordStratum, offsetText)
	}
}

//
//
func PersistSettingsValues(dir string, values map[string]string) error {
	data := "REDACTED"
	for k, v := range values {
		data += fmt.Sprintf("REDACTED", k, v)
	}
	strongmindfile := filepath.Join(dir, "REDACTED")
	return os.WriteFile(strongmindfile, []byte(data), 0o600)
}
