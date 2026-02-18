package diagnostics

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
	httpendpoint "github.com/valkyrieworks/rpc/requester/rest"
)

var killCmd = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDa
REDACTED,
REDACTEDa
REDACTED.

REDACTED:
REDACTED`,
	Args: cobra.ExactArgs(2),
	RunE: killCmdHandler,
}

func killCmdHandler(_ *cobra.Command, args []string) error {
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	outFile := args[1]
	if outFile == "REDACTED" {
		return errors.New("REDACTED")
	}

	rpc, err := httpendpoint.New(nodeRPCAddr, "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	home := viper.GetString(cli.HomeFlag)
	conf := cfg.DefaultConfig()
	conf = conf.SetRoot(home)
	cfg.EnsureRoot(conf.RootDir)

	//
	//
	tmpDir, err := os.MkdirTemp(os.TempDir(), "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	defer os.RemoveAll(tmpDir)

	logger.Info("REDACTED")
	if err := dumpStatus(rpc, tmpDir, "REDACTED"); err != nil {
		return err
	}

	logger.Info("REDACTED")
	if err := dumpNetInfo(rpc, tmpDir, "REDACTED"); err != nil {
		return err
	}

	logger.Info("REDACTED")
	if err := dumpConsensusState(rpc, tmpDir, "REDACTED"); err != nil {
		return err
	}

	logger.Info("REDACTED")
	if err := copyWAL(conf, tmpDir); err != nil {
		return err
	}

	logger.Info("REDACTED")
	if err := copyConfig(home, tmpDir); err != nil {
		return err
	}

	logger.Info("REDACTED")
	if err := killProc(pid, tmpDir); err != nil {
		return err
	}

	logger.Info("REDACTED")
	return zipDir(tmpDir, outFile)
}

//
//
//
//
//
func killProc(pid int, dir string) error {
	//
	//
	//
	cmd := exec.Command("REDACTED", "REDACTED", fmt.Sprintf("REDACTED", pid)) //

	outFile, err := os.Create(filepath.Join(dir, "REDACTED"))
	if err != nil {
		return err
	}
	defer outFile.Close()

	cmd.Stdout = outFile
	cmd.Stderr = outFile

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
