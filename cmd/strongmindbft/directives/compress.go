package directives

import (
	"errors"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

var CompressProceedStratumDatastoreDirective = &cobra.Command{
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
	RunE: func(cmd *cobra.Command, arguments []string) error {
		if settings.DatastoreRepository != "REDACTED" {
			return errors.New("REDACTED")
		}

		compressProceedStratumDeltaBytes(settings.OriginPath, tracer)
		return nil
	},
}

func compressProceedStratumDeltaBytes(originPath string, tracer log.Tracer) {
	datastoreIdentifiers := []string{"REDACTED", "REDACTED"}
	o := &opt.Options{
		DisableSeeksCompaction: true,
	}
	wg := sync.WaitGroup{}

	for _, datastoreAlias := range datastoreIdentifiers {

		wg.Add(1)
		go func() {
			defer wg.Done()
			datastoreRoute := filepath.Join(originPath, "REDACTED", datastoreAlias+"REDACTED")
			depot, err := leveldb.OpenFile(datastoreRoute, o)
			if err != nil {
				tracer.Failure("REDACTED", "REDACTED", datastoreRoute, "REDACTED", err)
				return
			}
			defer depot.Close()

			tracer.Details("REDACTED", "REDACTED", datastoreRoute)

			err = depot.CompactRange(util.Range{Start: nil, Limit: nil})
			if err != nil {
				tracer.Failure("REDACTED", "REDACTED", datastoreRoute, "REDACTED", err)
			}
		}()
	}
	wg.Wait()
}
