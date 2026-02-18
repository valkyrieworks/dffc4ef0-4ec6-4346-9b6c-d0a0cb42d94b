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
	ctsystem "github.com/valkyrieworks/utils/os"
)

//
func clearConfig(t *testing.T, dir string) {
	os.Clearenv()
	err := os.RemoveAll(dir)
	require.NoError(t, err)

	viper.Reset()
	config = cfg.DefaultConfig()
}

//
func testRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               RootCmd.Use,
		PersistentPreRunE: RootCmd.PersistentPreRunE,
		Run:               func(cmd *cobra.Command, args []string) {},
	}
	registerFlagsRootCmd(rootCmd)
	var l string
	rootCmd.PersistentFlags().String("REDACTED", l, "REDACTED")
	return rootCmd
}

func testSetup(t *testing.T, root string, args []string, env map[string]string) error {
	clearConfig(t, root)

	rootCmd := testRootCmd()
	cmd := cli.PrepareBaseCmd(rootCmd, "REDACTED", root)

	//
	args = append([]string{rootCmd.Use}, args...)
	return cli.RunWithArgs(cmd, args, env)
}

func TestRootHome(t *testing.T) {
	tmpDir := os.TempDir()
	root := filepath.Join(tmpDir, "REDACTED")
	newRoot := filepath.Join(tmpDir, "REDACTED")
	defer clearConfig(t, root)
	defer clearConfig(t, newRoot)

	cases := []struct {
		args []string
		env  map[string]string
		root string
	}{
		{nil, nil, root},
		{[]string{"REDACTED", newRoot}, nil, newRoot},
		{nil, map[string]string{"REDACTED": newRoot}, newRoot}, //
		{nil, map[string]string{"REDACTED": newRoot}, newRoot},
	}

	for i, tc := range cases {
		idxString := "REDACTED" + strconv.Itoa(i)

		err := testSetup(t, root, tc.args, tc.env)
		require.Nil(t, err, idxString)

		assert.Equal(t, tc.root, config.RootDir, idxString)
		assert.Equal(t, tc.root, config.P2P.RootDir, idxString)
		assert.Equal(t, tc.root, config.Consensus.RootDir, idxString)
		assert.Equal(t, tc.root, config.Mempool.RootDir, idxString)
	}
}

func TestRootFlagsEnv(t *testing.T) {
	//
	defaults := cfg.DefaultConfig()
	defaultLogLvl := defaults.LogLevel

	cases := []struct {
		args     []string
		env      map[string]string
		logLevel string
	}{
		{[]string{"REDACTED", "REDACTED"}, nil, defaultLogLvl},                  //
		{[]string{"REDACTED", "REDACTED"}, nil, "REDACTED"},                  //
		{nil, map[string]string{"REDACTED": "REDACTED"}, defaultLogLvl},        //
		{nil, map[string]string{"REDACTED": "REDACTED"}, defaultLogLvl},  //
		{nil, map[string]string{"REDACTED": "REDACTED"}, defaultLogLvl},  //
		{nil, map[string]string{"REDACTED": "REDACTED"}, defaultLogLvl},       //
		{nil, map[string]string{"REDACTED": "REDACTED"}, defaultLogLvl}, //
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED"},       //
	}

	for i, tc := range cases {
		idxString := strconv.Itoa(i)
		root := filepath.Join(os.TempDir(), "REDACTED"+idxString)
		idxString = "REDACTED" + idxString
		defer clearConfig(t, root)
		err := testSetup(t, root, tc.args, tc.env)
		require.Nil(t, err, idxString)

		assert.Equal(t, tc.logLevel, config.LogLevel, idxString)
	}
}

func TestRootConfig(t *testing.T) {
	//
	nonDefaultLogLvl := "REDACTED"
	cvals := map[string]string{
		"REDACTED": nonDefaultLogLvl,
	}

	cases := []struct {
		args []string
		env  map[string]string

		logLvl string
	}{
		{nil, nil, nonDefaultLogLvl},                                           //
		{[]string{"REDACTED"}, nil, "REDACTED"},                    //
		{nil, map[string]string{"REDACTED": "REDACTED"}, nonDefaultLogLvl}, //
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED"},      //
	}

	for i, tc := range cases {
		idxString := strconv.Itoa(i)
		root := filepath.Join(os.TempDir(), "REDACTED"+idxString)
		idxString = "REDACTED" + idxString
		defer clearConfig(t, root)
		//
		configFilePath := filepath.Join(root, "REDACTED")
		err := ctsystem.EnsureDir(configFilePath, 0o700)
		require.Nil(t, err)

		//
		//
		err = WriteConfigVals(configFilePath, cvals)
		require.Nil(t, err)

		rootCmd := testRootCmd()
		cmd := cli.PrepareBaseCmd(rootCmd, "REDACTED", root)

		//
		tc.args = append([]string{rootCmd.Use}, tc.args...)
		err = cli.RunWithArgs(cmd, tc.args, tc.env)
		require.Nil(t, err, idxString)

		assert.Equal(t, tc.logLvl, config.LogLevel, idxString)
	}
}

//
//
func WriteConfigVals(dir string, vals map[string]string) error {
	data := "REDACTED"
	for k, v := range vals {
		data += fmt.Sprintf("REDACTED", k, v)
	}
	cfile := filepath.Join(dir, "REDACTED")
	return os.WriteFile(cfile, []byte(data), 0o600)
}
