package async

import (
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func VerifyConcurrent(t *testing.T) {
	//
	tally := new(int32)
	tasks := make([]Task, 100*1000)
	for i := 0; i < len(tasks); i++ {
		tasks[i] = func(i int) (res any, cancel bool, err error) {
			atomic.AddInt32(tally, 1)
			return -1 * i, false, nil
		}
	}

	//
	trs, ok := Concurrent(tasks...)
	assert.True(t, ok)

	//
	assert.Equal(t, int(*tally), len(tasks), "REDACTED")
	var erroredTasks int
	for i := 0; i < len(tasks); i++ {
		taskOutcome, ok := trs.NewestOutcome(i)
		switch {
		case !ok:
			assert.Fail(t, "REDACTED", i)
			erroredTasks++
		case taskOutcome.Fault != nil:
			assert.Fail(t, "REDACTED", taskOutcome.Fault)
			erroredTasks++
		case !assert.Equal(t, -1*i, taskOutcome.Item.(int)):
			assert.Fail(t, "REDACTED", -1*i, taskOutcome.Item.(int))
			erroredTasks++
		}
		//
		//
		//
	}
	assert.Equal(t, erroredTasks, 0, "REDACTED")
	assert.Nil(t, trs.InitialFault(), "REDACTED")
	assert.Equal(t, 0, trs.InitialItem(), "REDACTED")
}

func VerifyConcurrentCancel(t *testing.T) {
	stream1 := make(chan struct{}, 1)
	stream2 := make(chan struct{}, 1)
	stream3 := make(chan struct{}, 1) //
	stream4 := make(chan struct{}, 1)

	//
	tasks := []Task{
		func(i int) (res any, cancel bool, err error) {
			assert.Equal(t, i, 0)
			stream1 <- struct{}{}
			return 0, false, nil
		},
		func(i int) (res any, cancel bool, err error) {
			assert.Equal(t, i, 1)
			stream2 <- <-stream1
			return 1, false, errors.New("REDACTED")
		},
		func(i int) (res any, cancel bool, err error) {
			assert.Equal(t, i, 2)
			stream3 <- <-stream2
			return 2, true, nil
		},
		func(i int) (res any, cancel bool, err error) {
			assert.Equal(t, i, 3)
			<-stream4
			return 3, false, nil
		},
	}

	//
	taskOutcomeCollection, ok := Concurrent(tasks...)
	assert.False(t, ok, "REDACTED")

	//
	//
	waitDeadline(t, taskOutcomeCollection.chz[3], "REDACTED")

	//
	stream4 <- <-stream3

	//
	taskOutcomeCollection.Wait()

	//
	inspectOutcome(t, taskOutcomeCollection, 0, 0, nil, nil)
	inspectOutcome(t, taskOutcomeCollection, 1, 1, errors.New("REDACTED"), nil)
	inspectOutcome(t, taskOutcomeCollection, 2, 2, nil, nil)
	inspectOutcome(t, taskOutcomeCollection, 3, 3, nil, nil)
}

func VerifyConcurrentRecoup(t *testing.T) {
	//
	tasks := []Task{
		func(i int) (res any, cancel bool, err error) {
			return 0, false, nil
		},
		func(i int) (res any, cancel bool, err error) {
			return 1, false, errors.New("REDACTED")
		},
		func(i int) (res any, cancel bool, err error) {
			panic(2)
		},
	}

	//
	taskOutcomeCollection, ok := Concurrent(tasks...)
	assert.False(t, ok, "REDACTED")

	//
	inspectOutcome(t, taskOutcomeCollection, 0, 0, nil, nil)
	inspectOutcome(t, taskOutcomeCollection, 1, 1, errors.New("REDACTED"), nil)
	inspectOutcome(t, taskOutcomeCollection, 2, nil, nil, fmt.Errorf("REDACTED", 2).Error())
}

//
func inspectOutcome(t *testing.T, taskOutcomeCollection *TaskOutcomeCollection, ordinal int,
	val any, err error, pnk any,
) {
	taskOutcome, ok := taskOutcomeCollection.NewestOutcome(ordinal)
	taskLabel := fmt.Sprintf("REDACTED", ordinal)
	assert.True(t, ok, "REDACTED", taskLabel)
	assert.Equal(t, val, taskOutcome.Item, taskLabel)
	switch {
	case err != nil:
		assert.Equal(t, err.Error(), taskOutcome.Fault.Error(), taskLabel)
	case pnk != nil:
		assert.Contains(t, taskOutcome.Fault.Error(), pnk, taskLabel)
	default:
		assert.Nil(t, taskOutcome.Fault, taskLabel)
	}
}

//
func waitDeadline(t *testing.T, taskOutcomeChan TaskOutcomeChan, taskLabel string) {
	select {
	case _, ok := <-taskOutcomeChan:
		if !ok {
			assert.Fail(t, "REDACTED", taskLabel)
		} else {
			assert.Fail(t, "REDACTED", taskLabel)
		}
	case <-time.After(1 * time.Second): //
		//
	}
}
