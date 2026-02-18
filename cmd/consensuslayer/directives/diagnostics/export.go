package diagnostics

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
	httpendpoint "github.com/valkyrieworks/rpc/requester/rest"
)

var dumpCmd = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDe
REDACTEDd
REDACTEDs
REDACTED`,
	Args: cobra.ExactArgs(1),
	RunE: dumpCmdHandler,
}

func init() {
	dumpCmd.Flags().UintVar(
		&frequency,
		flagFrequency,
		30,
		"REDACTED",
	)

	dumpCmd.Flags().StringVar(
		&profAddr,
		flagProfAddr,
		"REDACTED",
		"REDACTED",
	)
}

func dumpCmdHandler(_ *cobra.Command, args []string) error {
	outDir := args[0]
	if outDir == "REDACTED" {
		return errors.New("REDACTED")
	}

	if frequency == 0 {
		return errors.New("REDACTED")
	}

	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		if err := os.Mkdir(outDir, os.ModePerm); err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}

	rpc, err := httpendpoint.New(nodeRPCAddr, "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	home := viper.GetString(cli.HomeFlag)
	conf := cfg.DefaultConfig()
	conf = conf.SetRoot(home)
	cfg.EnsureRoot(conf.RootDir)

	dumpDebugData(outDir, conf, rpc)

	ticker := time.NewTicker(time.Duration(frequency) * time.Second)
	for range ticker.C {
		dumpDebugData(outDir, conf, rpc)
	}

	return nil
}

func dumpDebugData(outDir string, conf *cfg.Config, rpc *httpendpoint.HTTP) {
	start := time.Now().UTC()

	tmpDir, err := os.MkdirTemp(outDir, "REDACTED")
	if err != nil {
		logger.Error("REDACTED", "REDACTED", tmpDir, "REDACTED", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	logger.Info("REDACTED")
	if err := dumpStatus(rpc, tmpDir, "REDACTED"); err != nil {
		logger.Error("REDACTED", "REDACTED", err)
		return
	}

	logger.Info("REDACTED")
	if err := dumpNetInfo(rpc, tmpDir, "REDACTED"); err != nil {
		logger.Error("REDACTED", "REDACTED", err)
		return
	}

	logger.Info("REDACTED")
	if err := dumpConsensusState(rpc, tmpDir, "REDACTED"); err != nil {
		logger.Error("REDACTED", "REDACTED", err)
		return
	}

	logger.Info("REDACTED")
	if err := copyWAL(conf, tmpDir); err != nil {
		logger.Error("REDACTED", "REDACTED", err)
		return
	}

	if profAddr != "REDACTED" {
		logger.Info("REDACTED")
		if err := dumpProfile(tmpDir, profAddr, "REDACTED", 2); err != nil {
			logger.Error("REDACTED", "REDACTED", err)
			return
		}

		logger.Info("REDACTED")
		if err := dumpProfile(tmpDir, profAddr, "REDACTED", 2); err != nil {
			logger.Error("REDACTED", "REDACTED", err)
			return
		}
	}

	outFile := filepath.Join(outDir, fmt.Sprintf("REDACTED", start.Format(time.RFC3339)))
	if err := zipDir(tmpDir, outFile); err != nil {
		logger.Error("REDACTED", "REDACTED", outFile, "REDACTED", err)
	}
}
