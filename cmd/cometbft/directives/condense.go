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

var CondenseGoLayerStoreCommand = &cobra.Command{
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
		if settings.StoreOrigin != "REDACTED" {
			return errors.New("REDACTED")
		}

		condenseGoLayerDSz(settings.OriginFolder, tracer)
		return nil
	},
}

func condenseGoLayerDSz(originFolder string, tracer log.Tracer) {
	storeLabels := []string{"REDACTED", "REDACTED"}
	o := &opt.Options{
		DisableSeeksCompaction: true,
	}
	wg := sync.WaitGroup{}

	for _, storeLabel := range storeLabels {

		wg.Add(1)
		go func() {
			defer wg.Done()
			storeRoute := filepath.Join(originFolder, "REDACTED", storeLabel+"REDACTED")
			depot, err := leveldb.OpenFile(storeRoute, o)
			if err != nil {
				tracer.Fault("REDACTED", "REDACTED", storeRoute, "REDACTED", err)
				return
			}
			defer depot.Close()

			tracer.Details("REDACTED", "REDACTED", storeRoute)

			err = depot.CompactRange(util.Range{Start: nil, Limit: nil})
			if err != nil {
				tracer.Fault("REDACTED", "REDACTED", storeRoute, "REDACTED", err)
			}
		}()
	}
	wg.Wait()
}
