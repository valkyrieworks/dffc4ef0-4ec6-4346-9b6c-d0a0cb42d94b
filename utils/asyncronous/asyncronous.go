package asyncronous

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

//
//

//
//
//
type Activity func(i int) (val any, cancel bool, err error)

type ActivityOutcome struct {
	Datum any
	Failure error
}

type ActivityOutcomeChnl <-chan ActivityOutcome

type activityOutcomeOKAY struct {
	ActivityOutcome
	OK bool
}

type ActivityOutcomeAssign struct {
	chz     []ActivityOutcomeChnl
	outcomes []activityOutcomeOKAY
}

func freshActivityOutcomeAssign(chz []ActivityOutcomeChnl) *ActivityOutcomeAssign {
	return &ActivityOutcomeAssign{
		chz:     chz,
		outcomes: make([]activityOutcomeOKAY, len(chz)),
	}
}

func (trs *ActivityOutcomeAssign) Conduits() []ActivityOutcomeChnl {
	return trs.chz
}

func (trs *ActivityOutcomeAssign) NewestOutcome(ordinal int) (ActivityOutcome, bool) {
	if len(trs.outcomes) <= ordinal {
		return ActivityOutcome{}, false
	}
	outcomeOKAY := trs.outcomes[ordinal]
	return outcomeOKAY.ActivityOutcome, outcomeOKAY.OK
}

//
//
func (trs *ActivityOutcomeAssign) Harvest() *ActivityOutcomeAssign {
	for i := 0; i < len(trs.outcomes); i++ {
		trchnl := trs.chz[i]
		select {
		case outcome, ok := <-trchnl:
			if ok {
				//
				trs.outcomes[i] = activityOutcomeOKAY{
					ActivityOutcome: outcome,
					OK:         true,
				}
			}
			//
			//
			//
		default:
			//
		}
	}
	return trs
}

//
//
func (trs *ActivityOutcomeAssign) Pause() *ActivityOutcomeAssign {
	for i := 0; i < len(trs.outcomes); i++ {
		trchnl := trs.chz[i]
		outcome, ok := <-trchnl
		if ok {
			//
			trs.outcomes[i] = activityOutcomeOKAY{
				ActivityOutcome: outcome,
				OK:         true,
			}
		}
		//
		//
		//
	}
	return trs
}

//
//
func (trs *ActivityOutcomeAssign) InitialDatum() any {
	for _, outcome := range trs.outcomes {
		if outcome.Datum != nil {
			return outcome.Datum
		}
	}
	return nil
}

//
//
func (trs *ActivityOutcomeAssign) InitialFailure() error {
	for _, outcome := range trs.outcomes {
		if outcome.Failure != nil {
			return outcome.Failure
		}
	}
	return nil
}

//
//

//
//
//
//
//
func Concurrent(activities ...Activity) (trs *ActivityOutcomeAssign, ok bool) {
	activityOutcomeChnlz := make([]ActivityOutcomeChnl, len(activities)) //
	activityCompleteChnl := make(chan bool, len(activities))         //
	countAlarms := new(int32)                           //

	//
	ok = true

	//
	//
	//
	for i, activity := range activities {
		activityOutcomeChnl := make(chan ActivityOutcome, 1) //
		activityOutcomeChnlz[i] = activityOutcomeChnl
		go func(i int, activity Activity, activityOutcomeChnl chan ActivityOutcome) {
			//
			defer func() {
				if pnk := recover(); pnk != nil {
					atomic.AddInt32(countAlarms, 1)
					//
					const extent = 64 << 10
					buf := make([]byte, extent)
					buf = buf[:runtime.Stack(buf, false)]
					activityOutcomeChnl <- ActivityOutcome{nil, fmt.Errorf("REDACTED", pnk, buf)}
					//
					close(activityOutcomeChnl)
					//
					activityCompleteChnl <- false
				}
			}()
			//
			val, cancel, err := activity(i)
			//
			//
			activityOutcomeChnl <- ActivityOutcome{val, err}
			//
			close(activityOutcomeChnl)
			//
			activityCompleteChnl <- cancel
		}(i, activity, activityOutcomeChnl)
	}

	//
	//
	for i := 0; i < len(activities); i++ {
		cancel := <-activityCompleteChnl
		if cancel {
			ok = false
			break
		}
	}

	//
	//
	ok = ok && (atomic.LoadInt32(countAlarms) == 0)

	return freshActivityOutcomeAssign(activityOutcomeChnlz).Harvest(), ok
}
