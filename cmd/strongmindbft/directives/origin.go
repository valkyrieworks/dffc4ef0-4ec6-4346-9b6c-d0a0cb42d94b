package directives

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/cli"
	strongmindflags "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/cli/switches"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

var (
	settings = cfg.FallbackSettings()
	tracer = log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))
)

func initialize() {
	enrollSwitchesOriginDirective(OriginDirective)
}

func enrollSwitchesOriginDirective(cmd *cobra.Command) {
	cmd.PersistentFlags().String("REDACTED", settings.RecordStratum, "REDACTED")
}

//
//
func AnalyzeSettings(cmd *cobra.Command) (*cfg.Settings, error) {
	setting := cfg.FallbackSettings()
	err := viper.Unmarshal(setting)
	if err != nil {
		return nil, err
	}

	var domain string

	strongmindDomain := os.Getenv("REDACTED")
	tempDomain := os.Getenv("REDACTED")

	switch {
	case strongmindDomain != "REDACTED":
		domain = strongmindDomain

	case tempDomain != "REDACTED":
		//
		domain = tempDomain
		tracer.Failure("REDACTED")

	default:
		var err error
		domain, err = cmd.Flags().GetString(cli.DomainMarker)
		if err != nil {
			return nil, err
		}
	}

	setting.OriginPath = domain

	setting.AssignOrigin(setting.OriginPath)
	cfg.AssureOrigin(setting.OriginPath)
	if err := setting.CertifyFundamental(); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	if advisories := setting.InspectObsolete(); len(advisories) > 0 {
		for _, advisory := range advisories {
			tracer.Details("REDACTED", "REDACTED", advisory)
		}
	}
	return setting, nil
}

//
var OriginDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	PersistentPreRunE: func(cmd *cobra.Command, arguments []string) (err error) {
		if cmd.Name() == EditionDirective.Name() {
			return nil
		}

		settings, err = AnalyzeSettings(cmd)
		if err != nil {
			return err
		}

		if settings.RecordLayout == cfg.RecordLayoutJSN {
			tracer = log.FreshTempjsonTracer(log.FreshChronizePersistor(os.Stdout))
		}

		tracer, err = strongmindflags.AnalyzeRecordStratum(settings.RecordStratum, tracer, cfg.FallbackRecordStratum)
		if err != nil {
			return err
		}

		if viper.GetBool(cli.LoggingMarker) {
			tracer = log.FreshLoggingTracer(tracer)
		}

		tracer = tracer.Using("REDACTED", "REDACTED")
		return nil
	},
}
