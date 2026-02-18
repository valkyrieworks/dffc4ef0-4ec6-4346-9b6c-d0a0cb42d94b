package directives

import (
	"fmt"

	"github.com/spf13/cobra"

	cfg "github.com/valkyrieworks/settings"
	ctsystem "github.com/valkyrieworks/utils/os"
	ctrng "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/authkey"
	"github.com/valkyrieworks/kinds"
	cttime "github.com/valkyrieworks/kinds/moment"
)

//
var InitFilesCmd = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	RunE:  initFiles,
}

func initFiles(*cobra.Command, []string) error {
	return initFilesWithConfig(config)
}

func initFilesWithConfig(config *cfg.Config) error {
	//
	privValKeyFile := config.PrivValidatorKeyFile()
	privValStateFile := config.PrivValidatorStateFile()
	var pv *authkey.FilePV
	if ctsystem.FileExists(privValKeyFile) {
		pv = authkey.LoadFilePV(privValKeyFile, privValStateFile)
		logger.Info("REDACTED", "REDACTED", privValKeyFile,
			"REDACTED", privValStateFile)
	} else {
		pv = authkey.GenFilePV(privValKeyFile, privValStateFile)
		pv.Save()
		logger.Info("REDACTED", "REDACTED", privValKeyFile,
			"REDACTED", privValStateFile)
	}

	nodeKeyFile := config.NodeKeyFile()
	if ctsystem.FileExists(nodeKeyFile) {
		logger.Info("REDACTED", "REDACTED", nodeKeyFile)
	} else {
		if _, err := p2p.LoadOrGenNodeKey(nodeKeyFile); err != nil {
			return err
		}
		logger.Info("REDACTED", "REDACTED", nodeKeyFile)
	}

	//
	genFile := config.GenesisFile()
	if ctsystem.FileExists(genFile) {
		logger.Info("REDACTED", "REDACTED", genFile)
	} else {
		genDoc := kinds.GenesisDoc{
			ChainID:         fmt.Sprintf("REDACTED", ctrng.Str(6)),
			GenesisTime:     cttime.Now(),
			ConsensusParams: kinds.DefaultConsensusParams(),
		}
		pubKey, err := pv.GetPubKey()
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		genDoc.Validators = []kinds.GenesisValidator{{
			Address: pubKey.Address(),
			PubKey:  pubKey,
			Power:   10,
		}}

		if err := genDoc.SaveAs(genFile); err != nil {
			return err
		}
		logger.Info("REDACTED", "REDACTED", genFile)
	}

	return nil
}
