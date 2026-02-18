package agreement

import (
	"time"

	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"
)

var pulseTockBufferVolume = 10

//
//
//
type DeadlineTimer interface {
	Begin() error
	Halt() error
	Chan() <-chan deadlineDetails       //
	SequenceDeadline(ti deadlineDetails) //

	AssignTracer(log.Tracer)
}

//
//
//
//
//
type deadlineTimer struct {
	daemon.RootDaemon

	clockEnabled bool
	clock       *time.Timer
	pulseChan    chan deadlineDetails //
	tockChan    chan deadlineDetails //
}

//
func NewDeadlineTimer() DeadlineTimer {
	tt := &deadlineTimer{
		clock: time.NewTimer(0),
		//
		//
		clockEnabled: true,
		pulseChan:    make(chan deadlineDetails, pulseTockBufferVolume),
		tockChan:    make(chan deadlineDetails, pulseTockBufferVolume),
	}
	tt.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", tt)
	tt.haltClock() //
	return tt
}

//
func (t *deadlineTimer) OnBegin() error {
	go t.deadlineProcedure()

	return nil
}

//
func (t *deadlineTimer) OnHalt() {
	t.RootDaemon.OnHalt()
}

//
func (t *deadlineTimer) Chan() <-chan deadlineDetails {
	return t.tockChan
}

//
//
//
func (t *deadlineTimer) SequenceDeadline(ti deadlineDetails) {
	t.pulseChan <- ti
}

//

//
func (t *deadlineTimer) haltClock() {
	if !t.clockEnabled {
		return
	}
	//
	if !t.clock.Stop() {
		<-t.clock.C
	}
	t.clockEnabled = false
}

//
//
//
//
//
func (t *deadlineTimer) deadlineProcedure() {
	t.Tracer.Diagnose("REDACTED")
	var ti deadlineDetails
	for {
		select {
		case newti := <-t.pulseChan:
			t.Tracer.Diagnose("REDACTED", "REDACTED", ti, "REDACTED", newti)

			//
			if newti.Level < ti.Level {
				continue
			} else if newti.Level == ti.Level {
				if newti.Cycle < ti.Cycle {
					continue
				} else if newti.Cycle == ti.Cycle {
					if ti.Phase > 0 && newti.Phase <= ti.Phase {
						continue
					}
				}
			}

			//
			t.haltClock()

			//
			//
			ti = newti
			t.clock.Reset(ti.Period)
			t.clockEnabled = true

			t.Tracer.Diagnose("REDACTED", "REDACTED", ti.Period, "REDACTED", ti.Level, "REDACTED", ti.Cycle, "REDACTED", ti.Phase)
		case <-t.clock.C:
			t.clockEnabled = false
			t.Tracer.Details("REDACTED", "REDACTED", ti.Period, "REDACTED", ti.Level, "REDACTED", ti.Cycle, "REDACTED", ti.Phase)
			//
			//
			//
			//
			go func(toi deadlineDetails) { t.tockChan <- toi }(ti)
		case <-t.Exit():
			t.haltClock()
			return
		}
	}
}
