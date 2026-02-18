package agreement

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/agreement/kinds"
)

func TestTimeoutTicker(t *testing.T) {
	ticker := NewTimeoutTicker()
	err := ticker.Start()
	require.NoError(t, err)
	defer func() {
		err := ticker.Stop()
		require.NoError(t, err)
	}()

	c := ticker.Chan()
	for i := 1; i <= 10; i++ {
		height := int64(i)

		startTime := time.Now()
		//
		negTimeout := timeoutInfo{Duration: -1 * time.Millisecond, Height: height, Round: 0, Step: kinds.RoundStepNewHeight}
		timeout := timeoutInfo{Duration: 5 * time.Millisecond, Height: height, Round: 0, Step: kinds.RoundStepNewRound}
		ticker.ScheduleTimeout(negTimeout)
		ticker.ScheduleTimeout(timeout)

		//
		to := <-c
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		if timeout == to {
			require.True(t, elapsedTime >= timeout.Duration, "REDACTED", elapsedTime.Milliseconds(), startTime.UnixMilli(), endTime.UnixMilli())
		}
	}
}
