package clock

import (
	"time"

	engineconnect "github.com/valkyrieworks/utils/align"
)

/**
.
.
s
.
*/
type RegulateClock struct {
	Label string
	Ch   chan struct{}
	exit chan struct{}
	dur  time.Duration

	mtx   engineconnect.Lock
	clock *time.Timer
	isCollection bool
}

func NewRegulateClock(label string, dur time.Duration) *RegulateClock {
	ch := make(chan struct{})
	exit := make(chan struct{})
	t := &RegulateClock{Label: label, Ch: ch, dur: dur, exit: exit}
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
		t.isCollection = false
	case <-t.exit:
		//
	default:
		t.clock.Reset(t.dur)
	}
}

func (t *RegulateClock) Set() {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if !t.isCollection {
		t.isCollection = true
		t.clock.Reset(t.dur)
	}
}

func (t *RegulateClock) Clear() {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	t.isCollection = false
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
