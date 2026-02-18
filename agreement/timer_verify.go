package agreement

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/agreement/kinds"
)

func VerifyDeadlineTimer(t *testing.T) {
	timer := NewDeadlineTimer()
	err := timer.Begin()
	require.NoError(t, err)
	defer func() {
		err := timer.Halt()
		require.NoError(t, err)
	}()

	c := timer.Chan()
	for i := 1; i <= 10; i++ {
		level := int64(i)

		beginMoment := time.Now()
		//
		negativeDeadline := deadlineDetails{Period: -1 * time.Millisecond, Level: level, Cycle: 0, Phase: kinds.DurationPhaseNewLevel}
		deadline := deadlineDetails{Period: 5 * time.Millisecond, Level: level, Cycle: 0, Phase: kinds.EpochPhaseNewEpoch}
		timer.SequenceDeadline(negativeDeadline)
		timer.SequenceDeadline(deadline)

		//
		to := <-c
		terminateTime := time.Now()
		durationTime := terminateTime.Sub(beginMoment)
		if deadline == to {
			require.True(t, durationTime >= deadline.Period, "REDACTED", durationTime.Milliseconds(), beginMoment.UnixMilli(), terminateTime.UnixMilli())
		}
	}
}
