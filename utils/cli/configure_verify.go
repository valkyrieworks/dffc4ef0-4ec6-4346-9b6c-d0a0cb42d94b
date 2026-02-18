package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyConfigureContext(t *testing.T) {
	scenarios := []struct {
		args     []string
		env      map[string]string
		anticipated string
	}{
		{nil, nil, "REDACTED"},
		{[]string{"REDACTED", "REDACTED"}, nil, "REDACTED"},
		//
		{nil, nil, "REDACTED"},
		//
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED"},
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED"},
		//
		{
			[]string{"REDACTED", "REDACTED"},
			map[string]string{"REDACTED": "REDACTED"},
			"REDACTED",
		},
	}

	for idx, tc := range scenarios {
		i := strconv.Itoa(idx)
		//
		var foo string
		demo := &cobra.Command{
			Use: "REDACTED",
			RunE: func(cmd *cobra.Command, args []string) error {
				foo = viper.GetString("REDACTED")
				return nil
			},
		}
		demo.Flags().String("REDACTED", "REDACTED", "REDACTED")
		cmd := ArrangeRootCommand(demo, "REDACTED", "REDACTED") //
		cmd.Quit = func(int) {}

		viper.Reset()
		args := append([]string{cmd.Use}, tc.args...)
		err := ExecuteWithArgs(cmd, args, tc.env)
		require.Nil(t, err, i)
		assert.Equal(t, tc.anticipated, foo, i)
	}
}

func temporaryFolder() string {
	cdir, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}
	return cdir
}

func VerifyConfigureSettings(t *testing.T) {
	//
	//
	citem1 := "REDACTED"
	settings1 := temporaryFolder()
	err := RecordSettingsValues(settings1, map[string]string{"REDACTED": citem1})
	require.Nil(t, err)

	scenarios := []struct {
		args        []string
		env         map[string]string
		anticipated    string
		anticipatedDual string
	}{
		{nil, nil, "REDACTED", "REDACTED"},
		//
		{[]string{"REDACTED", "REDACTED"}, nil, "REDACTED", "REDACTED"},
		{[]string{"REDACTED", "REDACTED"}, nil, "REDACTED", "REDACTED"},
		{[]string{"REDACTED", settings1}, nil, citem1, "REDACTED"},
		//
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED", "REDACTED"},
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED", "REDACTED"},
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED", "REDACTED"},
		{nil, map[string]string{"REDACTED": settings1}, citem1, "REDACTED"},
		{nil, map[string]string{"REDACTED": settings1}, citem1, "REDACTED"},
	}

	for idx, tc := range scenarios {
		i := strconv.Itoa(idx)
		//
		var foo, two string
		boo := &cobra.Command{
			Use: "REDACTED",
			RunE: func(cmd *cobra.Command, args []string) error {
				foo = viper.GetString("REDACTED")
				two = viper.GetString("REDACTED")
				return nil
			},
		}
		boo.Flags().String("REDACTED", "REDACTED", "REDACTED")
		boo.Flags().String("REDACTED", "REDACTED", "REDACTED")
		cmd := ArrangeRootCommand(boo, "REDACTED", "REDACTED") //
		cmd.Quit = func(int) {}

		viper.Reset()
		args := append([]string{cmd.Use}, tc.args...)
		err := ExecuteWithArgs(cmd, args, tc.env)
		require.Nil(t, err, i)
		assert.Equal(t, tc.anticipated, foo, i)
		assert.Equal(t, tc.anticipatedDual, two, i)
	}
}

type DemoSettings struct {
	Label   string `mapstructure:"label"`
	Age    int    `mapstructure:"age"`
	Idle int    `mapstructure:"idle"`
}

func VerifyConfigureUnserialize(t *testing.T) {
	//
	//
	citem1, citem2 := "REDACTED", "REDACTED"
	settings1 := temporaryFolder()
	err := RecordSettingsValues(settings1, map[string]string{"REDACTED": citem1})
	require.Nil(t, err)
	//
	settings2 := temporaryFolder()
	err = RecordSettingsValues(settings2, map[string]string{"REDACTED": citem2, "REDACTED": "REDACTED"})
	require.Nil(t, err)

	//
	root := DemoSettings{
		Label:   "REDACTED",
		Age:    42,
		Idle: -7,
	}
	c := func(label string, age int) DemoSettings {
		r := root
		//
		//
		r.Label = "REDACTED"
		if label != "REDACTED" {
			r.Label = label
		}
		if age != 0 {
			r.Age = age
		}
		return r
	}

	scenarios := []struct {
		args     []string
		env      map[string]string
		anticipated DemoSettings
	}{
		{nil, nil, c("REDACTED", 0)},
		//
		{[]string{"REDACTED", "REDACTED"}, nil, c("REDACTED", 0)},
		{[]string{"REDACTED", settings1}, nil, c(citem1, 0)},
		//
		{nil, map[string]string{"REDACTED": "REDACTED"}, c("REDACTED", 56)},
		{nil, map[string]string{"REDACTED": settings1}, c(citem1, 0)},
		{[]string{"REDACTED", "REDACTED"}, map[string]string{"REDACTED": settings2}, c(citem2, 17)},
	}

	for idx, tc := range scenarios {
		i := strconv.Itoa(idx)
		//
		cfg := root
		serial := &cobra.Command{
			Use: "REDACTED",
			RunE: func(cmd *cobra.Command, args []string) error {
				return viper.Unmarshal(&cfg)
			},
		}
		serial.Flags().String("REDACTED", "REDACTED", "REDACTED")
		//
		//
		serial.Flags().Int("REDACTED", root.Age, "REDACTED")
		cmd := ArrangeRootCommand(serial, "REDACTED", "REDACTED") //
		cmd.Quit = func(int) {}

		viper.Reset()
		args := append([]string{cmd.Use}, tc.args...)
		err := ExecuteWithArgs(cmd, args, tc.env)
		require.Nil(t, err, i)
		assert.Equal(t, tc.anticipated, cfg, i)
	}
}

func VerifyConfigureTrack(t *testing.T) {
	scenarios := []struct {
		args     []string
		env      map[string]string
		lengthy     bool
		anticipated string
	}{
		{nil, nil, false, "REDACTED"},
		{[]string{"REDACTED"}, nil, true, "REDACTED"},
		{[]string{"REDACTED"}, nil, false, "REDACTED"},
		{nil, map[string]string{"REDACTED": "REDACTED"}, true, "REDACTED"},
	}

	for idx, tc := range scenarios {
		i := strconv.Itoa(idx)
		//
		track := &cobra.Command{
			Use: "REDACTED",
			RunE: func(cmd *cobra.Command, args []string) error {
				return fmt.Errorf("REDACTED", viper.GetBool(TrackMark))
			},
		}
		cmd := ArrangeRootCommand(track, "REDACTED", "REDACTED") //
		cmd.Quit = func(int) {}

		viper.Reset()
		args := append([]string{cmd.Use}, tc.args...)
		stdout, stderr, err := RunSeizeWithArgs(cmd, args, tc.env)
		require.NotNil(t, err, i)
		require.Equal(t, "REDACTED", stdout, i)
		require.NotEqual(t, "REDACTED", stderr, i)
		msg := strings.Split(stderr, "REDACTED")
		sought := fmt.Sprintf("REDACTED", tc.anticipated)
		assert.Equal(t, sought, msg[0], i)
		t.Log(msg)
		if tc.lengthy && assert.True(t, len(msg) > 2, i) {
			//
			assert.Contains(t, stderr, "REDACTED", i)
			assert.Contains(t, stderr, "REDACTED", i)
		}
	}
}
