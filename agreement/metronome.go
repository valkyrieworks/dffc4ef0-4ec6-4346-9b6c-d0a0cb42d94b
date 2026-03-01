package agreement

import (
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

var pulseChimeReserveExtent = 10

//
//
//
type DeadlineMetronome interface {
	Initiate() error
	Halt() error
	Channel() <-chan deadlineDetails       //
	TimelineDeadline(ti deadlineDetails) //

	AssignTracer(log.Tracer)
}

//
//
//
//
//
type deadlineMetronome struct {
	facility.FoundationFacility

	clockDynamic bool
	clock       *time.Timer
	pulseChannel    chan deadlineDetails //
	chimeChannel    chan deadlineDetails //
}

//
func FreshDeadlineMetronome() DeadlineMetronome {
	tt := &deadlineMetronome{
		clock: time.NewTimer(0),
		//
		//
		clockDynamic: true,
		pulseChannel:    make(chan deadlineDetails, pulseChimeReserveExtent),
		chimeChannel:    make(chan deadlineDetails, pulseChimeReserveExtent),
	}
	tt.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", tt)
	tt.haltClock() //
	return tt
}

//
func (t *deadlineMetronome) UponInitiate() error {
	go t.deadlineProcedure()

	return nil
}

//
func (t *deadlineMetronome) UponHalt() {
	t.FoundationFacility.UponHalt()
}

//
func (t *deadlineMetronome) Channel() <-chan deadlineDetails {
	return t.chimeChannel
}

//
//
//
func (t *deadlineMetronome) TimelineDeadline(ti deadlineDetails) {
	t.pulseChannel <- ti
}

//

//
func (t *deadlineMetronome) haltClock() {
	if !t.clockDynamic {
		return
	}
	//
	if !t.clock.Stop() {
		<-t.clock.C
	}
	t.clockDynamic = false
}

//
//
//
//
//
func (t *deadlineMetronome) deadlineProcedure() {
	t.Tracer.Diagnose("REDACTED")
	var ti deadlineDetails
	for {
		select {
		case newtimer := <-t.pulseChannel:
			t.Tracer.Diagnose("REDACTED", "REDACTED", ti, "REDACTED", newtimer)

			//
			if newtimer.Altitude < ti.Altitude {
				continue
			} else if newtimer.Altitude == ti.Altitude {
				if newtimer.Iteration < ti.Iteration {
					continue
				} else if newtimer.Iteration == ti.Iteration {
					if ti.Phase > 0 && newtimer.Phase <= ti.Phase {
						continue
					}
				}
			}

			//
			t.haltClock()

			//
			//
			ti = newtimer
			t.clock.Reset(ti.Interval)
			t.clockDynamic = true

			t.Tracer.Diagnose("REDACTED", "REDACTED", ti.Interval, "REDACTED", ti.Altitude, "REDACTED", ti.Iteration, "REDACTED", ti.Phase)
		case <-t.clock.C:
			t.clockDynamic = false
			t.Tracer.Details("REDACTED", "REDACTED", ti.Interval, "REDACTED", ti.Altitude, "REDACTED", ti.Iteration, "REDACTED", ti.Phase)
			//
			//
			//
			//
			go func(toi deadlineDetails) { t.chimeChannel <- toi }(ti)
		case <-t.Exit():
			t.haltClock()
			return
		}
	}
}
