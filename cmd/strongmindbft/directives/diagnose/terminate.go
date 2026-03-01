package diagnose

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/cli"
	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
)

var terminateDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDa
REDACTED,
REDACTEDa
REDACTED.

REDACTED:
REDACTED`,
	Args: cobra.ExactArgs(2),
	RunE: terminateDirectiveProcessor,
}

func terminateDirectiveProcessor(_ *cobra.Command, arguments []string) error {
	pid, err := strconv.Atoi(arguments[0])
	if err != nil {
		return err
	}

	outputRecord := arguments[1]
	if outputRecord == "REDACTED" {
		return errors.New("REDACTED")
	}

	rpc, err := rpchttpsvc.New(peerRemoteLocation, "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	domain := viper.GetString(cli.DomainMarker)
	setting := cfg.FallbackSettings()
	setting = setting.AssignOrigin(domain)
	cfg.AssureOrigin(setting.OriginPath)

	//
	//
	scratchPath, err := os.MkdirTemp(os.TempDir(), "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	defer os.RemoveAll(scratchPath)

	tracer.Details("REDACTED")
	if err := exportCondition(rpc, scratchPath, "REDACTED"); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	if err := exportNetworkDetails(rpc, scratchPath, "REDACTED"); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	if err := exportAgreementStatus(rpc, scratchPath, "REDACTED"); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	if err := duplicateJournal(setting, scratchPath); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	if err := duplicateSettings(domain, scratchPath); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	if err := terminateRoutine(pid, scratchPath); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	return compressPath(scratchPath, outputRecord)
}

//
//
//
//
//
func terminateRoutine(pid int, dir string) error {
	//
	//
	//
	cmd := exec.Command("REDACTED", "REDACTED", fmt.Sprintf("REDACTED", pid)) //

	outputRecord, err := os.Create(filepath.Join(dir, "REDACTED"))
	if err != nil {
		return err
	}
	defer outputRecord.Close()

	cmd.Stdout = outputRecord
	cmd.Stderr = outputRecord

	if err := cmd.Start(); err != nil {
		return err
	}

	//
	go func() {
		//
		//
		p, err := os.FindProcess(pid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "REDACTED", err)
		} else if err = p.Signal(syscall.SIGABRT); err != nil {
			fmt.Fprintf(os.Stderr, "REDACTED", err)
		}

		//
		//
		//
		//
		time.Sleep(5 * time.Second)

		if err := cmd.Process.Kill(); err != nil {
			fmt.Fprintf(os.Stderr, "REDACTED", err)
		}
	}()

	if err := cmd.Wait(); err != nil {
		//
		if _, ok := err.(*exec.ExitError); !ok {
			return err
		}
	}

	return nil
}
