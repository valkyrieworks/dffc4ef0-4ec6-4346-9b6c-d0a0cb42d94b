package autofile

import (
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	engineseed "github.com/valkyrieworks/utils/random"
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
	automaticEntryEndDuration = 1000 * time.Millisecond
	automaticEntryModes       = os.FileMode(0o600)
)

//
//
//
//
type AutomaticEntry struct {
	ID   string
	Route string

	endTimer      *time.Ticker
	endTimerHaltchan chan struct{} //
	hupchan             chan os.Signal

	mtx  sync.Mutex
	entry *os.File
}

//
//
//
func AccessAutomaticEntry(route string) (*AutomaticEntry, error) {
	var err error
	route, err = filepath.Abs(route)
	if err != nil {
		return nil, err
	}
	af := &AutomaticEntry{
		ID:               engineseed.Str(12) + "REDACTED" + route,
		Route:             route,
		endTimer:      time.NewTicker(automaticEntryEndDuration),
		endTimerHaltchan: make(chan struct{}),
	}
	if err := af.accessEntry(); err != nil {
		af.End()
		return nil, err
	}

	//
	af.hupchan = make(chan os.Signal, 1)
	signal.Notify(af.hupchan, syscall.SIGHUP)
	go func() {
		for range af.hupchan {
			_ = af.endEntry()
		}
	}()

	go af.endEntryProcedure()

	return af, nil
}

//
//
func (af *AutomaticEntry) End() error {
	af.endTimer.Stop()
	close(af.endTimerHaltchan)
	if af.hupchan != nil {
		close(af.hupchan)
	}
	return af.endEntry()
}

func (af *AutomaticEntry) endEntryProcedure() {
	for {
		select {
		case <-af.endTimer.C:
			_ = af.endEntry()
		case <-af.endTimerHaltchan:
			return
		}
	}
}

func (af *AutomaticEntry) endEntry() (err error) {
	af.mtx.Lock()
	defer af.mtx.Unlock()

	entry := af.entry
	if entry == nil {
		return nil
	}

	af.entry = nil
	return entry.Close()
}

//
//
//
//
func (af *AutomaticEntry) Record(b []byte) (n int, err error) {
	af.mtx.Lock()
	defer af.mtx.Unlock()

	if af.entry == nil {
		if err = af.accessEntry(); err != nil {
			return
		}
	}

	n, err = af.entry.Write(b)
	return
}

//
//
//
//
func (af *AutomaticEntry) Align() error {
	af.mtx.Lock()
	defer af.mtx.Unlock()

	if af.entry == nil {
		if err := af.accessEntry(); err != nil {
			return err
		}
	}
	return af.entry.Sync()
}

func (af *AutomaticEntry) accessEntry() error {
	entry, err := os.OpenFile(af.Route, os.O_RDWR|os.O_CREATE|os.O_APPEND, automaticEntryModes)
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
	af.entry = entry
	return nil
}

//
//
//
func (af *AutomaticEntry) Volume() (int64, error) {
	af.mtx.Lock()
	defer af.mtx.Unlock()

	if af.entry == nil {
		if err := af.accessEntry(); err != nil {
			return -1, err
		}
	}

	status, err := af.entry.Stat()
	if err != nil {
		return -1, err
	}
	return status.Size(), nil
}
