package diagnose

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/cli"
	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
)

var exportDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDe
REDACTEDd
REDACTEDs
REDACTED`,
	Args: cobra.ExactArgs(1),
	RunE: exportDirectiveProcessor,
}

func initialize() {
	exportDirective.Flags().UintVar(
		&recurrence,
		markerRecurrence,
		30,
		"REDACTED",
	)

	exportDirective.Flags().StringVar(
		&analyzerLocation,
		markerAnalyzerLocation,
		"REDACTED",
		"REDACTED",
	)
}

func exportDirectiveProcessor(_ *cobra.Command, arguments []string) error {
	outputPath := arguments[0]
	if outputPath == "REDACTED" {
		return errors.New("REDACTED")
	}

	if recurrence == 0 {
		return errors.New("REDACTED")
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		if err := os.Mkdir(outputPath, os.ModePerm); err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}

	rpc, err := rpchttpsvc.New(peerRemoteLocation, "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	domain := viper.GetString(cli.DomainMarker)
	setting := cfg.FallbackSettings()
	setting = setting.AssignOrigin(domain)
	cfg.AssureOrigin(setting.OriginPath)

	exportDiagnoseData(outputPath, setting, rpc)

	metronome := time.NewTicker(time.Duration(recurrence) * time.Second)
	for range metronome.C {
		exportDiagnoseData(outputPath, setting, rpc)
	}

	return nil
}

func exportDiagnoseData(outputPath string, setting *cfg.Settings, rpc *rpchttpsvc.Httpsvc) {
	initiate := time.Now().UTC()

	scratchPath, err := os.MkdirTemp(outputPath, "REDACTED")
	if err != nil {
		tracer.Failure("REDACTED", "REDACTED", scratchPath, "REDACTED", err)
		return
	}
	defer os.RemoveAll(scratchPath)

	tracer.Details("REDACTED")
	if err := exportCondition(rpc, scratchPath, "REDACTED"); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
		return
	}

	tracer.Details("REDACTED")
	if err := exportNetworkDetails(rpc, scratchPath, "REDACTED"); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
		return
	}

	tracer.Details("REDACTED")
	if err := exportAgreementStatus(rpc, scratchPath, "REDACTED"); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
		return
	}

	tracer.Details("REDACTED")
	if err := duplicateJournal(setting, scratchPath); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
		return
	}

	if analyzerLocation != "REDACTED" {
		tracer.Details("REDACTED")
		if err := exportAnalysis(scratchPath, analyzerLocation, "REDACTED", 2); err != nil {
			tracer.Failure("REDACTED", "REDACTED", err)
			return
		}

		tracer.Details("REDACTED")
		if err := exportAnalysis(scratchPath, analyzerLocation, "REDACTED", 2); err != nil {
			tracer.Failure("REDACTED", "REDACTED", err)
			return
		}
	}

	outputRecord := filepath.Join(outputPath, fmt.Sprintf("REDACTED", initiate.Format(time.RFC3339)))
	if err := compressPath(scratchPath, outputRecord); err != nil {
		tracer.Failure("REDACTED", "REDACTED", outputRecord, "REDACTED", err)
	}
}
