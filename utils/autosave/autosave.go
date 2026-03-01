package autosave

import (
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

/*e

t
)
{
)
}

.
.
{
)
)
}

e
)
{
)
}
*/

const (
	automaticRecordShutdownSpan = 1000 * time.Millisecond
	automaticRecordModes       = os.FileMode(0o600)
)

//
//
//
//
type AutomaticRecord struct {
	ID   string
	Route string

	shutdownMetronome      *time.Ticker
	shutdownMetronomeStopchnl chan struct{} //
	huprchnl             chan os.Signal

	mtx  sync.Mutex
	record *os.File
}

//
//
//
func UnlockAutomaticRecord(route string) (*AutomaticRecord, error) {
	var err error
	route, err = filepath.Abs(route)
	if err != nil {
		return nil, err
	}
	af := &AutomaticRecord{
		ID:               commitrand.Str(12) + "REDACTED" + route,
		Route:             route,
		shutdownMetronome:      time.NewTicker(automaticRecordShutdownSpan),
		shutdownMetronomeStopchnl: make(chan struct{}),
	}
	if err := af.unlockRecord(); err != nil {
		af.Shutdown()
		return nil, err
	}

	//
	af.huprchnl = make(chan os.Signal, 1)
	signal.Notify(af.huprchnl, syscall.SIGHUP)
	go func() {
		for range af.huprchnl {
			_ = af.shutdownRecord()
		}
	}()

	go af.shutdownRecordProcedure()

	return af, nil
}

//
//
func (af *AutomaticRecord) Shutdown() error {
	af.shutdownMetronome.Stop()
	close(af.shutdownMetronomeStopchnl)
	if af.huprchnl != nil {
		close(af.huprchnl)
	}
	return af.shutdownRecord()
}

func (af *AutomaticRecord) shutdownRecordProcedure() {
	for {
		select {
		case <-af.shutdownMetronome.C:
			_ = af.shutdownRecord()
		case <-af.shutdownMetronomeStopchnl:
			return
		}
	}
}

func (af *AutomaticRecord) shutdownRecord() (err error) {
	af.mtx.Lock()
	defer af.mtx.Unlock()

	record := af.record
	if record == nil {
		return nil
	}

	af.record = nil
	return record.Close()
}

//
//
//
//
func (af *AutomaticRecord) Record(b []byte) (n int, err error) {
	af.mtx.Lock()
	defer af.mtx.Unlock()

	if af.record == nil {
		if err = af.unlockRecord(); err != nil {
			return
		}
	}

	n, err = af.record.Write(b)
	return
}

//
//
//
//
func (af *AutomaticRecord) Chronize() error {
	af.mtx.Lock()
	defer af.mtx.Unlock()

	if af.record == nil {
		if err := af.unlockRecord(); err != nil {
			return err
		}
	}
	return af.record.Sync()
}

func (af *AutomaticRecord) unlockRecord() error {
	record, err := os.OpenFile(af.Route, os.O_RDWR|os.O_CREATE|os.O_APPEND, automaticRecordModes)
	if err != nil {
		return err
	}
	//
	//
	//
	//
	//
	//
	//
	af.record = record
	return nil
}

//
//
//
func (af *AutomaticRecord) Extent() (int64, error) {
	af.mtx.Lock()
	defer af.mtx.Unlock()

	if af.record == nil {
		if err := af.unlockRecord(); err != nil {
			return -1, err
		}
	}

	summary, err := af.record.Stat()
	if err != nil {
		return -1, err
	}
	return summary.Size(), nil
}
