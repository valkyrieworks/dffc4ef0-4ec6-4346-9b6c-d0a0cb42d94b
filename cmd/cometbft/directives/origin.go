package directives

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/cli"
	cometmarks "github.com/valkyrieworks/utils/cli/marks"
	"github.com/valkyrieworks/utils/log"
)

var (
	settings = cfg.StandardSettings()
	tracer = log.NewTMTracer(log.NewAlignRecorder(os.Stdout))
)

func init() {
	enrollOptionsOriginCommand(OriginCommand)
}

func enrollOptionsOriginCommand(cmd *cobra.Command) {
	cmd.PersistentFlags().String("REDACTED", settings.TraceLayer, "REDACTED")
}

//
//
func AnalyzeSettings(cmd *cobra.Command) (*cfg.Settings, error) {
	cfg := cfg.StandardSettings()
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	var home string

	cometHome := os.Getenv("REDACTED")
	tmHome := os.Getenv("REDACTED")

	switch {
	case cometHome != "REDACTED":
		home = cometHome

	case tmHome != "REDACTED":
		//
		home = tmHome
		tracer.Fault("REDACTED")

	default:
		var err error
		home, err = cmd.Flags().GetString(cli.HomeMark)
		if err != nil {
			return nil, err
		}
	}

	cfg.OriginFolder = home

	cfg.AssignOrigin(cfg.OriginFolder)
	cfg.AssureOrigin(cfg.OriginFolder)
	if err := cfg.CertifySimple(); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	if cautions := cfg.InspectObsolete(); len(cautions) > 0 {
		for _, caution := range cautions {
			tracer.Details("REDACTED", "REDACTED", caution)
		}
	}
	return cfg, nil
}

//
var OriginCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if cmd.Name() == ReleaseCommand.Name() {
			return nil
		}

		settings, err = AnalyzeSettings(cmd)
		if err != nil {
			return err
		}

		if settings.TraceLayout == cfg.TraceLayoutJSON {
			tracer = log.NewTmjsonTracer(log.NewAlignRecorder(os.Stdout))
		}

		tracer, err = cometmarks.AnalyzeTraceLayer(settings.TraceLayer, tracer, cfg.StandardTraceLayer)
		if err != nil {
			return err
		}

		if viper.GetBool(cli.TrackMark) {
			tracer = log.NewLoggingTracer(tracer)
		}

		tracer = tracer.With("REDACTED", "REDACTED")
		return nil
	},
}
