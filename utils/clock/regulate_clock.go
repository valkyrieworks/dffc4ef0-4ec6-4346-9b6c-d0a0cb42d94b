package clock

import (
	"time"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

/**
.
.
s
.
*/
type RegulateClock struct {
	Alias string
	Ch   chan struct{}
	exit chan struct{}
	dur  time.Duration

	mtx   commitchronize.Exclusion
	clock *time.Timer
	equalsAssign bool
}

func FreshRegulateClock(alias string, dur time.Duration) *RegulateClock {
	ch := make(chan struct{})
	exit := make(chan struct{})
	t := &RegulateClock{Alias: alias, Ch: ch, dur: dur, exit: exit}
	t.mtx.Lock()
	t.clock = time.AfterFunc(dur, t.triggerProcedure)
	t.mtx.Unlock()
	t.clock.Stop()
	return t
}

func (t *RegulateClock) triggerProcedure() {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	select {
	case t.Ch <- struct{}{}:
		t.equalsAssign = false
	case <-t.exit:
		//
	default:
		t.clock.Reset(t.dur)
	}
}

func (t *RegulateClock) Set() {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if !t.equalsAssign {
		t.equalsAssign = true
		t.clock.Reset(t.dur)
	}
}

func (t *RegulateClock) Deassign() {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	t.equalsAssign = false
	t.clock.Stop()
}

//
//
func (t *RegulateClock) Halt() bool {
	if t == nil {
		return false
	}
	close(t.exit)
	t.mtx.Lock()
	defer t.mtx.Unlock()
	return t.clock.Stop()
}
