package asyncronous

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
	activities := make([]Activity, 100*1000)
	for i := 0; i < len(activities); i++ {
		activities[i] = func(i int) (res any, cancel bool, err error) {
			atomic.AddInt32(tally, 1)
			return -1 * i, false, nil
		}
	}

	//
	trs, ok := Concurrent(activities...)
	assert.True(t, ok)

	//
	assert.Equal(t, int(*tally), len(activities), "REDACTED")
	var unsuccessfulActivities int
	for i := 0; i < len(activities); i++ {
		activityOutcome, ok := trs.NewestOutcome(i)
		switch {
		case !ok:
			assert.Fail(t, "REDACTED", i)
			unsuccessfulActivities++
		case activityOutcome.Failure != nil:
			assert.Fail(t, "REDACTED", activityOutcome.Failure)
			unsuccessfulActivities++
		case !assert.Equal(t, -1*i, activityOutcome.Datum.(int)):
			assert.Fail(t, "REDACTED", -1*i, activityOutcome.Datum.(int))
			unsuccessfulActivities++
		}
		//
		//
		//
	}
	assert.Equal(t, unsuccessfulActivities, 0, "REDACTED")
	assert.Nil(t, trs.InitialFailure(), "REDACTED")
	assert.Equal(t, 0, trs.InitialDatum(), "REDACTED")
}

func VerifyConcurrentCancel(t *testing.T) {
	stream1 := make(chan struct{}, 1)
	stream2 := make(chan struct{}, 1)
	stream3 := make(chan struct{}, 1) //
	stream4 := make(chan struct{}, 1)

	//
	activities := []Activity{
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
	activityOutcomeAssign, ok := Concurrent(activities...)
	assert.False(t, ok, "REDACTED")

	//
	//
	pauseDeadline(t, activityOutcomeAssign.chz[3], "REDACTED")

	//
	stream4 <- <-stream3

	//
	activityOutcomeAssign.Pause()

	//
	inspectOutcome(t, activityOutcomeAssign, 0, 0, nil, nil)
	inspectOutcome(t, activityOutcomeAssign, 1, 1, errors.New("REDACTED"), nil)
	inspectOutcome(t, activityOutcomeAssign, 2, 2, nil, nil)
	inspectOutcome(t, activityOutcomeAssign, 3, 3, nil, nil)
}

func VerifyConcurrentRestore(t *testing.T) {
	//
	activities := []Activity{
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
	activityOutcomeAssign, ok := Concurrent(activities...)
	assert.False(t, ok, "REDACTED")

	//
	inspectOutcome(t, activityOutcomeAssign, 0, 0, nil, nil)
	inspectOutcome(t, activityOutcomeAssign, 1, 1, errors.New("REDACTED"), nil)
	inspectOutcome(t, activityOutcomeAssign, 2, nil, nil, fmt.Errorf("REDACTED", 2).Error())
}

//
func inspectOutcome(t *testing.T, activityOutcomeAssign *ActivityOutcomeAssign, ordinal int,
	val any, err error, pnk any,
) {
	activityOutcome, ok := activityOutcomeAssign.NewestOutcome(ordinal)
	activityAlias := fmt.Sprintf("REDACTED", ordinal)
	assert.True(t, ok, "REDACTED", activityAlias)
	assert.Equal(t, val, activityOutcome.Datum, activityAlias)
	switch {
	case err != nil:
		assert.Equal(t, err.Error(), activityOutcome.Failure.Error(), activityAlias)
	case pnk != nil:
		assert.Contains(t, activityOutcome.Failure.Error(), pnk, activityAlias)
	default:
		assert.Nil(t, activityOutcome.Failure, activityAlias)
	}
}

//
func pauseDeadline(t *testing.T, activityOutcomeChnl ActivityOutcomeChnl, activityAlias string) {
	select {
	case _, ok := <-activityOutcomeChnl:
		if !ok {
			assert.Fail(t, "REDACTED", activityAlias)
		} else {
			assert.Fail(t, "REDACTED", activityAlias)
		}
	case <-time.After(1 * time.Second): //
		//
	}
}
