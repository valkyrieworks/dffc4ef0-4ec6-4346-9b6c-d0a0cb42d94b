package agreement

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement/kinds"
)

func VerifyDeadlineMetronome(t *testing.T) {
	metronome := FreshDeadlineMetronome()
	err := metronome.Initiate()
	require.NoError(t, err)
	defer func() {
		err := metronome.Halt()
		require.NoError(t, err)
	}()

	c := metronome.Conduit()
	for i := 1; i <= 10; i++ {
		altitude := int64(i)

		initiateMoment := time.Now()
		//
		negativeDeadline := deadlineDetails{Interval: -1 * time.Millisecond, Altitude: altitude, Iteration: 0, Phase: kinds.IterationPhaseFreshAltitude}
		deadline := deadlineDetails{Interval: 5 * time.Millisecond, Altitude: altitude, Iteration: 0, Phase: kinds.IterationPhaseFreshIteration}
		metronome.ArrangeDeadline(negativeDeadline)
		metronome.ArrangeDeadline(deadline)

		//
		to := <-c
		terminateMoment := time.Now()
		passedMoment := terminateMoment.Sub(initiateMoment)
		if deadline == to {
			require.True(t, passedMoment >= deadline.Interval, "REDACTED", passedMoment.Milliseconds(), initiateMoment.UnixMilli(), terminateMoment.UnixMilli())
		}
	}
}
