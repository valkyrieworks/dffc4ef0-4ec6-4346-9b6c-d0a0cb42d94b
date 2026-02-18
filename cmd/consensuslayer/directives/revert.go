package directives

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	dbm "github.com/valkyrieworks/-db"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
)

var removeBlock = false

func init() {
	RollbackStateCmd.Flags().BoolVar(&removeBlock, "REDACTED", false, "REDACTED")
}

var RollbackStateCmd = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `
REDACTED,
REDACTEDe
REDACTED.
REDACTED 
REDACTED 
REDACTEDn
REDACTED.
REDACTED`,
	RunE: func(cmd *cobra.Command, args []string) error {
		height, hash, err := RollbackState(config, removeBlock)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}

		if removeBlock {
			fmt.Printf("REDACTED", height, hash)
		} else {
			fmt.Printf("REDACTED", height, hash)
		}

		return nil
	},
}

//
//
//
func RollbackState(config *cfg.Config, removeBlock bool) (int64, []byte, error) {
	//
	blockStore, stateStore, err := loadStateAndBlockStore(config)
	if err != nil {
		return -1, nil, err
	}
	defer func() {
		_ = blockStore.Close()
		_ = stateStore.Close()
	}()

	//
	return status.Rollback(blockStore, stateStore, removeBlock)
}

func loadStateAndBlockStore(config *cfg.Config) (*depot.BlockStore, status.Store, error) {
	dbType := dbm.BackendType(config.DBBackend)

	if !os.FileExists(filepath.Join(config.DBDir(), "REDACTED")) {
		return nil, nil, fmt.Errorf("REDACTED", config.DBDir())
	}

	//
	blockStoreDB, err := dbm.NewDB("REDACTED", dbType, config.DBDir())
	if err != nil {
		return nil, nil, err
	}
	blockStore := depot.NewBlockStore(blockStoreDB)

	if !os.FileExists(filepath.Join(config.DBDir(), "REDACTED")) {
		return nil, nil, fmt.Errorf("REDACTED", config.DBDir())
	}

	//
	stateDB, err := dbm.NewDB("REDACTED", dbType, config.DBDir())
	if err != nil {
		return nil, nil, err
	}
	stateStore := status.NewStore(stateDB, status.StoreOptions{
		DiscardABCIResponses: config.Storage.DiscardABCIResponses,
	})

	return blockStore, stateStore, nil
}
