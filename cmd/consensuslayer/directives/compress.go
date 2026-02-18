package directives

import (
	"errors"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"

	"github.com/valkyrieworks/utils/log"
)

var CompactGoLevelDBCmd = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	Long: `
REDACTED 
REDACTED 
REDACTEDr
REDACTED.

REDACTED.
REDACTED`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if config.DBBackend != "REDACTED" {
			return errors.New("REDACTED")
		}

		compactGoLevelDBs(config.RootDir, logger)
		return nil
	},
}

func compactGoLevelDBs(rootDir string, logger log.Logger) {
	dbNames := []string{"REDACTED", "REDACTED"}
	o := &opt.Options{
		DisableSeeksCompaction: true,
	}
	wg := sync.WaitGroup{}

	for _, dbName := range dbNames {

		wg.Add(1)
		go func() {
			defer wg.Done()
			dbPath := filepath.Join(rootDir, "REDACTED", dbName+"REDACTED")
			store, err := leveldb.OpenFile(dbPath, o)
			if err != nil {
				logger.Error("REDACTED", "REDACTED", dbPath, "REDACTED", err)
				return
			}
			defer store.Close()

			logger.Info("REDACTED", "REDACTED", dbPath)

			err = store.CompactRange(util.Range{Start: nil, Limit: nil})
			if err != nil {
				logger.Error("REDACTED", "REDACTED", dbPath, "REDACTED", err)
			}
		}()
	}
	wg.Wait()
}
