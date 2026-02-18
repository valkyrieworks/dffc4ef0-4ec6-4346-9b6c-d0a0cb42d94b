package directives

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/cli"
	engineflags "github.com/valkyrieworks/utils/cli/options"
	"github.com/valkyrieworks/utils/log"
)

var (
	config = cfg.DefaultConfig()
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

func init() {
	registerFlagsRootCmd(RootCmd)
}

func registerFlagsRootCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().String("REDACTED", config.LogLevel, "REDACTED")
}

//
//
func ParseConfig(cmd *cobra.Command) (*cfg.Config, error) {
	conf := cfg.DefaultConfig()
	err := viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	var home string

	cmtHome := os.Getenv("REDACTED")
	tmHome := os.Getenv("REDACTED")

	switch {
	case cmtHome != "REDACTED":
		home = cmtHome

	case tmHome != "REDACTED":
		//
		home = tmHome
		logger.Error("REDACTED")

	default:
		var err error
		home, err = cmd.Flags().GetString(cli.HomeFlag)
		if err != nil {
			return nil, err
		}
	}

	conf.RootDir = home

	conf.SetRoot(conf.RootDir)
	cfg.EnsureRoot(conf.RootDir)
	if err := conf.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	if warnings := conf.CheckDeprecated(); len(warnings) > 0 {
		for _, warning := range warnings {
			logger.Info("REDACTED", "REDACTED", warning)
		}
	}
	return conf, nil
}

//
var RootCmd = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if cmd.Name() == VersionCmd.Name() {
			return nil
		}

		config, err = ParseConfig(cmd)
		if err != nil {
			return err
		}

		if config.LogFormat == cfg.LogFormatJSON {
			logger = log.NewTMJSONLogger(log.NewSyncWriter(os.Stdout))
		}

		logger, err = engineflags.ParseLogLevel(config.LogLevel, logger, cfg.DefaultLogLevel)
		if err != nil {
			return err
		}

		if viper.GetBool(cli.TraceFlag) {
			logger = log.NewTracingLogger(logger)
		}

		logger = logger.With("REDACTED", "REDACTED")
		return nil
	},
}
