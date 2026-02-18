package directives

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/utils/log"
	ctsystem "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/authkey"
)

//
//
var ResetAllCmd = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    resetAllCmd,
}

var keepAddrBook bool

//
var ResetStateCmd = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		config, err = ParseConfig(cmd)
		if err != nil {
			return err
		}

		return resetState(config.DBDir(), logger)
	},
}

func init() {
	ResetAllCmd.Flags().BoolVar(&keepAddrBook, "REDACTED", false, "REDACTED")
}

//
var ResetPrivValidatorCmd = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	RunE:    resetPrivValidator,
}

//
//
func resetAllCmd(cmd *cobra.Command, _ []string) (err error) {
	config, err = ParseConfig(cmd)
	if err != nil {
		return err
	}

	return resetAll(
		config.DBDir(),
		config.P2P.AddrBookFile(),
		config.PrivValidatorKeyFile(),
		config.PrivValidatorStateFile(),
		logger,
	)
}

//
//
func resetPrivValidator(cmd *cobra.Command, _ []string) (err error) {
	config, err = ParseConfig(cmd)
	if err != nil {
		return err
	}

	resetFilePV(config.PrivValidatorKeyFile(), config.PrivValidatorStateFile(), logger)
	return nil
}

//
func resetAll(dbDir, addrBookFile, privValKeyFile, privValStateFile string, logger log.Logger) error {
	if keepAddrBook {
		logger.Info("REDACTED")
	} else {
		removeAddrBook(addrBookFile, logger)
	}

	if err := os.RemoveAll(dbDir); err == nil {
		logger.Info("REDACTED", "REDACTED", dbDir)
	} else {
		logger.Error("REDACTED", "REDACTED", dbDir, "REDACTED", err)
	}

	if err := ctsystem.EnsureDir(dbDir, 0o700); err != nil {
		logger.Error("REDACTED", "REDACTED", err)
	}

	//
	resetFilePV(privValKeyFile, privValStateFile, logger)
	return nil
}

//
func resetState(dbDir string, logger log.Logger) error {
	blockdb := filepath.Join(dbDir, "REDACTED")
	state := filepath.Join(dbDir, "REDACTED")
	wal := filepath.Join(dbDir, "REDACTED")
	evidence := filepath.Join(dbDir, "REDACTED")
	txIndex := filepath.Join(dbDir, "REDACTED")

	if ctsystem.FileExists(blockdb) {
		if err := os.RemoveAll(blockdb); err == nil {
			logger.Info("REDACTED", "REDACTED", blockdb)
		} else {
			logger.Error("REDACTED", "REDACTED", blockdb, "REDACTED", err)
		}
	}

	if ctsystem.FileExists(state) {
		if err := os.RemoveAll(state); err == nil {
			logger.Info("REDACTED", "REDACTED", state)
		} else {
			logger.Error("REDACTED", "REDACTED", state, "REDACTED", err)
		}
	}

	if ctsystem.FileExists(wal) {
		if err := os.RemoveAll(wal); err == nil {
			logger.Info("REDACTED", "REDACTED", wal)
		} else {
			logger.Error("REDACTED", "REDACTED", wal, "REDACTED", err)
		}
	}

	if ctsystem.FileExists(evidence) {
		if err := os.RemoveAll(evidence); err == nil {
			logger.Info("REDACTED", "REDACTED", evidence)
		} else {
			logger.Error("REDACTED", "REDACTED", evidence, "REDACTED", err)
		}
	}

	if ctsystem.FileExists(txIndex) {
		if err := os.RemoveAll(txIndex); err == nil {
			logger.Info("REDACTED", "REDACTED", txIndex)
		} else {
			logger.Error("REDACTED", "REDACTED", txIndex, "REDACTED", err)
		}
	}

	if err := ctsystem.EnsureDir(dbDir, 0o700); err != nil {
		logger.Error("REDACTED", "REDACTED", err)
	}
	return nil
}

func resetFilePV(privValKeyFile, privValStateFile string, logger log.Logger) {
	if _, err := os.Stat(privValKeyFile); err == nil {
		pv := authkey.LoadFilePVEmptyState(privValKeyFile, privValStateFile)
		pv.Reset()
		logger.Info(
			"REDACTED",
			"REDACTED", privValKeyFile,
			"REDACTED", privValStateFile,
		)
	} else {
		pv := authkey.GenFilePV(privValKeyFile, privValStateFile)
		pv.Save()
		logger.Info(
			"REDACTED",
			"REDACTED", privValKeyFile,
			"REDACTED", privValStateFile,
		)
	}
}

func removeAddrBook(addrBookFile string, logger log.Logger) {
	if err := os.Remove(addrBookFile); err == nil {
		logger.Info("REDACTED", "REDACTED", addrBookFile)
	} else if !os.IsNotExist(err) {
		logger.Info("REDACTED", "REDACTED", addrBookFile, "REDACTED", err)
	}
}
