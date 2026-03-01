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

func VerifyConfigureEnvironment(t *testing.T) {
	scenarios := []struct {
		arguments     []string
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
		prototype := &cobra.Command{
			Use: "REDACTED",
			RunE: func(cmd *cobra.Command, arguments []string) error {
				foo = viper.GetString("REDACTED")
				return nil
			},
		}
		prototype.Flags().String("REDACTED", "REDACTED", "REDACTED")
		cmd := ArrangeFoundationDirective(prototype, "REDACTED", "REDACTED") //
		cmd.Quit = func(int) {}

		viper.Reset()
		arguments := append([]string{cmd.Use}, tc.arguments...)
		err := ExecuteUsingArguments(cmd, arguments, tc.env)
		require.Nil(t, err, i)
		assert.Equal(t, tc.anticipated, foo, i)
	}
}

func transientPath() string {
	cnpath, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}
	return cnpath
}

func VerifyConfigureSettings(t *testing.T) {
	//
	//
	cnval1 := "REDACTED"
	cfg1 := transientPath()
	err := PersistSettingsValues(cfg1, map[string]string{"REDACTED": cnval1})
	require.Nil(t, err)

	scenarios := []struct {
		arguments        []string
		env         map[string]string
		anticipated    string
		anticipatedCouple string
	}{
		{nil, nil, "REDACTED", "REDACTED"},
		//
		{[]string{"REDACTED", "REDACTED"}, nil, "REDACTED", "REDACTED"},
		{[]string{"REDACTED", "REDACTED"}, nil, "REDACTED", "REDACTED"},
		{[]string{"REDACTED", cfg1}, nil, cnval1, "REDACTED"},
		//
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED", "REDACTED"},
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED", "REDACTED"},
		{nil, map[string]string{"REDACTED": "REDACTED"}, "REDACTED", "REDACTED"},
		{nil, map[string]string{"REDACTED": cfg1}, cnval1, "REDACTED"},
		{nil, map[string]string{"REDACTED": cfg1}, cnval1, "REDACTED"},
	}

	for idx, tc := range scenarios {
		i := strconv.Itoa(idx)
		//
		var foo, two string
		boo := &cobra.Command{
			Use: "REDACTED",
			RunE: func(cmd *cobra.Command, arguments []string) error {
				foo = viper.GetString("REDACTED")
				two = viper.GetString("REDACTED")
				return nil
			},
		}
		boo.Flags().String("REDACTED", "REDACTED", "REDACTED")
		boo.Flags().String("REDACTED", "REDACTED", "REDACTED")
		cmd := ArrangeFoundationDirective(boo, "REDACTED", "REDACTED") //
		cmd.Quit = func(int) {}

		viper.Reset()
		arguments := append([]string{cmd.Use}, tc.arguments...)
		err := ExecuteUsingArguments(cmd, arguments, tc.env)
		require.Nil(t, err, i)
		assert.Equal(t, tc.anticipated, foo, i)
		assert.Equal(t, tc.anticipatedCouple, two, i)
	}
}

type PrototypeSettings struct {
	Alias   string `mapstructure:"alias"`
	Age    int    `mapstructure:"age"`
	Idle int    `mapstructure:"idle"`
}

func VerifyConfigureDecode(t *testing.T) {
	//
	//
	cnval1, cnval2 := "REDACTED", "REDACTED"
	cfg1 := transientPath()
	err := PersistSettingsValues(cfg1, map[string]string{"REDACTED": cnval1})
	require.Nil(t, err)
	//
	cfg2 := transientPath()
	err = PersistSettingsValues(cfg2, map[string]string{"REDACTED": cnval2, "REDACTED": "REDACTED"})
	require.Nil(t, err)

	//
	foundation := PrototypeSettings{
		Alias:   "REDACTED",
		Age:    42,
		Idle: -7,
	}
	c := func(alias string, age int) PrototypeSettings {
		r := foundation
		//
		//
		r.Alias = "REDACTED"
		if alias != "REDACTED" {
			r.Alias = alias
		}
		if age != 0 {
			r.Age = age
		}
		return r
	}

	scenarios := []struct {
		arguments     []string
		env      map[string]string
		anticipated PrototypeSettings
	}{
		{nil, nil, c("REDACTED", 0)},
		//
		{[]string{"REDACTED", "REDACTED"}, nil, c("REDACTED", 0)},
		{[]string{"REDACTED", cfg1}, nil, c(cnval1, 0)},
		//
		{nil, map[string]string{"REDACTED": "REDACTED"}, c("REDACTED", 56)},
		{nil, map[string]string{"REDACTED": cfg1}, c(cnval1, 0)},
		{[]string{"REDACTED", "REDACTED"}, map[string]string{"REDACTED": cfg2}, c(cnval2, 17)},
	}

	for idx, tc := range scenarios {
		i := strconv.Itoa(idx)
		//
		cfg := foundation
		encode := &cobra.Command{
			Use: "REDACTED",
			RunE: func(cmd *cobra.Command, arguments []string) error {
				return viper.Unmarshal(&cfg)
			},
		}
		encode.Flags().String("REDACTED", "REDACTED", "REDACTED")
		//
		//
		encode.Flags().Int("REDACTED", foundation.Age, "REDACTED")
		cmd := ArrangeFoundationDirective(encode, "REDACTED", "REDACTED") //
		cmd.Quit = func(int) {}

		viper.Reset()
		arguments := append([]string{cmd.Use}, tc.arguments...)
		err := ExecuteUsingArguments(cmd, arguments, tc.env)
		require.Nil(t, err, i)
		assert.Equal(t, tc.anticipated, cfg, i)
	}
}

func VerifyConfigureLogging(t *testing.T) {
	scenarios := []struct {
		arguments     []string
		env      map[string]string
		extended     bool
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
		logging := &cobra.Command{
			Use: "REDACTED",
			RunE: func(cmd *cobra.Command, arguments []string) error {
				return fmt.Errorf("REDACTED", viper.GetBool(LoggingMarker))
			},
		}
		cmd := ArrangeFoundationDirective(logging, "REDACTED", "REDACTED") //
		cmd.Quit = func(int) {}

		viper.Reset()
		arguments := append([]string{cmd.Use}, tc.arguments...)
		standardemission, standardfailure, err := ExecuteSeizeUsingArguments(cmd, arguments, tc.env)
		require.NotNil(t, err, i)
		require.Equal(t, "REDACTED", standardemission, i)
		require.NotEqual(t, "REDACTED", standardfailure, i)
		msg := strings.Split(standardfailure, "REDACTED")
		intended := fmt.Sprintf("REDACTED", tc.anticipated)
		assert.Equal(t, intended, msg[0], i)
		t.Log(msg)
		if tc.extended && assert.True(t, len(msg) > 2, i) {
			//
			assert.Contains(t, standardfailure, "REDACTED", i)
			assert.Contains(t, standardfailure, "REDACTED", i)
		}
	}
}
