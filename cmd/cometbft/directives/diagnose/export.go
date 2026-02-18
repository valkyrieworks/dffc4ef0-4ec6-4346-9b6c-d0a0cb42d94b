package diagnose

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/cli"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
)

var exportCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDe
REDACTEDd
REDACTEDs
REDACTED`,
	Args: cobra.ExactArgs(1),
	RunE: exportCommandManager,
}

func init() {
	exportCommand.Flags().UintVar(
		&recurrence,
		markRecurrence,
		30,
		"REDACTED",
	)

	exportCommand.Flags().StringVar(
		&profAddress,
		markProfAddress,
		"REDACTED",
		"REDACTED",
	)
}

func exportCommandManager(_ *cobra.Command, args []string) error {
	outFolder := args[0]
	if outFolder == "REDACTED" {
		return errors.New("REDACTED")
	}

	if recurrence == 0 {
		return errors.New("REDACTED")
	}

	if _, err := os.Stat(outFolder); os.IsNotExist(err) {
		if err := os.Mkdir(outFolder, os.ModePerm); err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}

	rpc, err := rpchttp.New(memberRPCAddress, "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	home := viper.GetString(cli.HomeMark)
	cfg := cfg.StandardSettings()
	cfg = cfg.AssignOrigin(home)
	cfg.AssureOrigin(cfg.OriginFolder)

	exportDiagnoseData(outFolder, cfg, rpc)

	timer := time.NewTicker(time.Duration(recurrence) * time.Second)
	for range timer.C {
		exportDiagnoseData(outFolder, cfg, rpc)
	}

	return nil
}

func exportDiagnoseData(outFolder string, cfg *cfg.Settings, rpc *rpchttp.HTTP) {
	begin := time.Now().UTC()

	tempFolder, err := os.MkdirTemp(outFolder, "REDACTED")
	if err != nil {
		tracer.Fault("REDACTED", "REDACTED", tempFolder, "REDACTED", err)
		return
	}
	defer os.RemoveAll(tempFolder)

	tracer.Details("REDACTED")
	if err := exportState(rpc, tempFolder, "REDACTED"); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
		return
	}

	tracer.Details("REDACTED")
	if err := exportNetDetails(rpc, tempFolder, "REDACTED"); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
		return
	}

	tracer.Details("REDACTED")
	if err := exportAgreementStatus(rpc, tempFolder, "REDACTED"); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
		return
	}

	tracer.Details("REDACTED")
	if err := cloneJournal(cfg, tempFolder); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
		return
	}

	if profAddress != "REDACTED" {
		tracer.Details("REDACTED")
		if err := exportBlueprint(tempFolder, profAddress, "REDACTED", 2); err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
			return
		}

		tracer.Details("REDACTED")
		if err := exportBlueprint(tempFolder, profAddress, "REDACTED", 2); err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
			return
		}
	}

	outEntry := filepath.Join(outFolder, fmt.Sprintf("REDACTED", begin.Format(time.RFC3339)))
	if err := zipFolder(tempFolder, outEntry); err != nil {
		tracer.Fault("REDACTED", "REDACTED", outEntry, "REDACTED", err)
	}
}
