package async

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
type Task func(i int) (val any, cancel bool, err error)

type TaskOutcome struct {
	Item any
	Fault error
}

type TaskOutcomeChan <-chan TaskOutcome

type taskOutcomeOK struct {
	TaskOutcome
	OK bool
}

type TaskOutcomeCollection struct {
	chz     []TaskOutcomeChan
	outcomes []taskOutcomeOK
}

func newTaskOutcomeCollection(chz []TaskOutcomeChan) *TaskOutcomeCollection {
	return &TaskOutcomeCollection{
		chz:     chz,
		outcomes: make([]taskOutcomeOK, len(chz)),
	}
}

func (trs *TaskOutcomeCollection) Streams() []TaskOutcomeChan {
	return trs.chz
}

func (trs *TaskOutcomeCollection) NewestOutcome(ordinal int) (TaskOutcome, bool) {
	if len(trs.outcomes) <= ordinal {
		return TaskOutcome{}, false
	}
	outcomeOK := trs.outcomes[ordinal]
	return outcomeOK.TaskOutcome, outcomeOK.OK
}

//
//
func (trs *TaskOutcomeCollection) Harvest() *TaskOutcomeCollection {
	for i := 0; i < len(trs.outcomes); i++ {
		trchan := trs.chz[i]
		select {
		case outcome, ok := <-trchan:
			if ok {
				//
				trs.outcomes[i] = taskOutcomeOK{
					TaskOutcome: outcome,
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
func (trs *TaskOutcomeCollection) Wait() *TaskOutcomeCollection {
	for i := 0; i < len(trs.outcomes); i++ {
		trchan := trs.chz[i]
		outcome, ok := <-trchan
		if ok {
			//
			trs.outcomes[i] = taskOutcomeOK{
				TaskOutcome: outcome,
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
func (trs *TaskOutcomeCollection) InitialItem() any {
	for _, outcome := range trs.outcomes {
		if outcome.Item != nil {
			return outcome.Item
		}
	}
	return nil
}

//
//
func (trs *TaskOutcomeCollection) InitialFault() error {
	for _, outcome := range trs.outcomes {
		if outcome.Fault != nil {
			return outcome.Fault
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
func Concurrent(tasks ...Task) (trs *TaskOutcomeCollection, ok bool) {
	taskOutcomeChz := make([]TaskOutcomeChan, len(tasks)) //
	taskDoneChan := make(chan bool, len(tasks))         //
	countAlarms := new(int32)                           //

	//
	ok = true

	//
	//
	//
	for i, task := range tasks {
		taskOutcomeChan := make(chan TaskOutcome, 1) //
		taskOutcomeChz[i] = taskOutcomeChan
		go func(i int, task Task, taskOutcomeChan chan TaskOutcome) {
			//
			defer func() {
				if pnk := recover(); pnk != nil {
					atomic.AddInt32(countAlarms, 1)
					//
					const volume = 64 << 10
					buf := make([]byte, volume)
					buf = buf[:runtime.Stack(buf, false)]
					taskOutcomeChan <- TaskOutcome{nil, fmt.Errorf("REDACTED", pnk, buf)}
					//
					close(taskOutcomeChan)
					//
					taskDoneChan <- false
				}
			}()
			//
			val, cancel, err := task(i)
			//
			//
			taskOutcomeChan <- TaskOutcome{val, err}
			//
			close(taskOutcomeChan)
			//
			taskDoneChan <- cancel
		}(i, task, taskOutcomeChan)
	}

	//
	//
	for i := 0; i < len(tasks); i++ {
		cancel := <-taskDoneChan
		if cancel {
			ok = false
			break
		}
	}

	//
	//
	ok = ok && (atomic.LoadInt32(countAlarms) == 0)

	return newTaskOutcomeCollection(taskOutcomeChz).Harvest(), ok
}
