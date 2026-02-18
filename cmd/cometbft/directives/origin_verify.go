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

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/cli"
	cometos "github.com/valkyrieworks/utils/os"
)

//
func flushSettings(t *testing.T, dir string) {
	os.Clearenv()
	err := os.RemoveAll(dir)
	require.NoError(t, err)

	viper.Reset()
	settings = cfg.StandardSettings()
}

//
func verifyOriginCommand() *cobra.Command {
	originCommand := &cobra.Command{
		Use:               OriginCommand.Use,
		PersistentPreRunE: OriginCommand.PersistentPreRunE,
		Run:               func(cmd *cobra.Command, args []string) {},
	}
	enrollOptionsOriginCommand(originCommand)
	var l string
	originCommand.PersistentFlags().String("REDACTED", l, "REDACTED")
	return originCommand
}

func verifyConfigure(t *testing.T, origin string, args []string, env map[string]string) error {
	flushSettings(t, origin)

	originCommand := verifyOriginCommand()
	cmd := cli.ArrangeRootCommand(originCommand, "REDACTED", origin)

	//
	args = append([]string{originCommand.Use}, args...)
	return cli.ExecuteWithArgs(cmd, args, env)
}

func VerifyOriginHome(t *testing.T) {
	tempFolder := os.TempDir()
	origin := filepath.Join(tempFolder, "REDACTED")
	newOrigin := filepath.Join(tempFolder, "REDACTED")
	defer flushSettings(t, origin)
	defer flushSettings(t, newOrigin)

	scenarios := []struct {
		args []string
		env  map[string]string
		origin string
	}{
		{nil, nil, origin},
		{[]string{"REDACTED", newOrigin}, nil, newOrigin},
		{nil, map[string]string{"REDACTED": newOrigin}, newOrigin}, //
		{nil, map[string]string{"REDACTED": newOrigin}, newOrigin},
	}

	for i, tc := range scenarios {
		idxString := "REDACTED" + strconv.Itoa(i)

		err := verifyConfigure(t, origin, tc.args, tc.env)
		require.Nil(t, err, idxString)

		assert.Equal(t, tc.origin, settings.OriginFolder, idxString)
		assert.Equal(t, tc.origin, settings.P2P.OriginFolder, idxString)
		assert.Equal(t, tc.origin, settings.Agreement.OriginFolder, idxString)
		assert.Equal(t, tc.origin, settings.Txpool.OriginFolder, idxString)
	}
}

func VerifyOriginOptionsContext(t *testing.T) {
	//
	standards := cfg.StandardSettings()
	standardTraceTier := standards.TraceLayer

	scenarios := []struct {
		args     []string
		env      map[string]string
		traceLayer string
	}{
		{[]string{"REDACTED", "REDACTED"}, nil, standardTraceTier},                  //
		{[]string{"REDACTED", "REDACTED"}, nil, "REDACTED"},                  //
		{nil, map[string]string{"REDACTED": "REDACTED"}, standardTraceTier},        //
		{nil, map[string]string{"REDACTED": "REDACTED"}, standardTraceTier},  //
		{nil, map[string]string{"REDACTED": "REDACTED"}, standardTraceTier},  //
		{nil, map[string]string{"REDACTED": "REDACTED"}, standardTraceTier},       //
		{nil, map[string]string{"REDACTED": "REDACTED"}, standardTraceTier}, //
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED"},       //
	}

	for i, tc := range scenarios {
		idxString := strconv.Itoa(i)
		origin := filepath.Join(os.TempDir(), "REDACTED"+idxString)
		idxString = "REDACTED" + idxString
		defer flushSettings(t, origin)
		err := verifyConfigure(t, origin, tc.args, tc.env)
		require.Nil(t, err, idxString)

		assert.Equal(t, tc.traceLayer, settings.TraceLayer, idxString)
	}
}

func VerifyOriginSettings(t *testing.T) {
	//
	notStandardTraceTier := "REDACTED"
	cvalues := map[string]string{
		"REDACTED": notStandardTraceTier,
	}

	scenarios := []struct {
		args []string
		env  map[string]string

		traceTier string
	}{
		{nil, nil, notStandardTraceTier},                                           //
		{[]string{"REDACTED"}, nil, "REDACTED"},                    //
		{nil, map[string]string{"REDACTED": "REDACTED"}, notStandardTraceTier}, //
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED"},      //
	}

	for i, tc := range scenarios {
		idxString := strconv.Itoa(i)
		origin := filepath.Join(os.TempDir(), "REDACTED"+idxString)
		idxString = "REDACTED" + idxString
		defer flushSettings(t, origin)
		//
		settingsEntryRoute := filepath.Join(origin, "REDACTED")
		err := cometos.AssureFolder(settingsEntryRoute, 0o700)
		require.Nil(t, err)

		//
		//
		err = RecordSettingsValues(settingsEntryRoute, cvalues)
		require.Nil(t, err)

		originCommand := verifyOriginCommand()
		cmd := cli.ArrangeRootCommand(originCommand, "REDACTED", origin)

		//
		tc.args = append([]string{originCommand.Use}, tc.args...)
		err = cli.ExecuteWithArgs(cmd, tc.args, tc.env)
		require.Nil(t, err, idxString)

		assert.Equal(t, tc.traceTier, settings.TraceLayer, idxString)
	}
}

//
//
func RecordSettingsValues(dir string, values map[string]string) error {
	data := "REDACTED"
	for k, v := range values {
		data += fmt.Sprintf("REDACTED", k, v)
	}
	centry := filepath.Join(dir, "REDACTED")
	return os.WriteFile(centry, []byte(data), 0o600)
}
