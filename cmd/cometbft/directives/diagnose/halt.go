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

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/cli"
	rpchttp "github.com/valkyrieworks/rpc/customer/http"
)

var haltCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDa
REDACTED,
REDACTEDa
REDACTED.

REDACTED:
REDACTED`,
	Args: cobra.ExactArgs(2),
	RunE: haltCommandManager,
}

func haltCommandManager(_ *cobra.Command, args []string) error {
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	outEntry := args[1]
	if outEntry == "REDACTED" {
		return errors.New("REDACTED")
	}

	rpc, err := rpchttp.New(memberRPCAddress, "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	home := viper.GetString(cli.HomeMark)
	cfg := cfg.StandardSettings()
	cfg = cfg.AssignOrigin(home)
	cfg.AssureOrigin(cfg.OriginFolder)

	//
	//
	tempFolder, err := os.MkdirTemp(os.TempDir(), "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	defer os.RemoveAll(tempFolder)

	tracer.Details("REDACTED")
	if err := exportState(rpc, tempFolder, "REDACTED"); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	if err := exportNetDetails(rpc, tempFolder, "REDACTED"); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	if err := exportAgreementStatus(rpc, tempFolder, "REDACTED"); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	if err := cloneJournal(cfg, tempFolder); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	if err := cloneSettings(home, tempFolder); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	if err := haltProc(pid, tempFolder); err != nil {
		return err
	}

	tracer.Details("REDACTED")
	return zipFolder(tempFolder, outEntry)
}

//
//
//
//
//
func haltProc(pid int, dir string) error {
	//
	//
	//
	cmd := exec.Command("REDACTED", "REDACTED", fmt.Sprintf("REDACTED", pid)) //

	outEntry, err := os.Create(filepath.Join(dir, "REDACTED"))
	if err != nil {
		return err
	}
	defer outEntry.Close()

	cmd.Stdout = outEntry
	cmd.Stderr = outEntry

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
