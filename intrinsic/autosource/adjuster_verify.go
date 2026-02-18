package autosource

import (
	"testing"
	"time"

	"github.com/valkyrieworks/utils/log"
	"github.com/stretchr/testify/require"
)

func VerifyVelocityWaitperiodAdjuster(t *testing.T) {
	//
	//
	const (
		min                 = 4
		max                 = 10
		limitQuantile = 90.0
		limitWaitperiod    = 100 * time.Millisecond
		eraPeriod       = time.Second

		bufferLimit = 10
	)

	tracer := log.VerifyingTracer()

	adjuster := NewVelocityWaitperiodAdjuster(
		min,
		max,
		limitQuantile,
		limitWaitperiod,
		eraPeriod,
		tracer,
	)

	countOperators := min

	for ordinal, tt := range []struct {
		waitperiodsMillis        []int
		bufferSize           int
		anticipatedVerdict   uint8
		anticipatedCountOperators int
	}{
		{
			waitperiodsMillis:        []int{},
			bufferSize:           5,
			anticipatedVerdict:   MustRemain,
			anticipatedCountOperators: min,
		},
		{
			//
			waitperiodsMillis:        []int{200},
			bufferSize:           5,
			anticipatedVerdict:   MustRemain,
			anticipatedCountOperators: min,
		},
		{
			waitperiodsMillis:        []int{50, 50, 50},
			bufferSize:           5,
			anticipatedVerdict:   MustSize,
			anticipatedCountOperators: 5,
		},
		{
			waitperiodsMillis:        []int{50, 50, 50},
			bufferSize:           5,
			anticipatedVerdict:   MustSize,
			anticipatedCountOperators: 6,
		},
		{
			waitperiodsMillis:        []int{50, 50, 50, 80},
			bufferSize:           5,
			anticipatedVerdict:   MustSize,
			anticipatedCountOperators: 7,
		},
		{
			waitperiodsMillis:        []int{50, 50, 50, 80},
			anticipatedVerdict:   MustSize,
			anticipatedCountOperators: 8,
		},
		{
			waitperiodsMillis:        []int{50, 50, 50, 80, 90, 90},
			bufferSize:           5,
			anticipatedVerdict:   MustSize,
			anticipatedCountOperators: 9,
		},
		{
			waitperiodsMillis:        []int{50, 50, 50},
			bufferSize:           5,
			anticipatedVerdict:   MustReduce,
			anticipatedCountOperators: 8,
		},
		{
			waitperiodsMillis:        []int{50, 50, 50, 80, 90, 90, 95, 99},
			bufferSize:           5,
			anticipatedVerdict:   MustSize,
			anticipatedCountOperators: 9,
		},
		{
			//
			waitperiodsMillis:        []int{50, 50, 50},
			bufferSize:           7,
			anticipatedVerdict:   MustSize,
			anticipatedCountOperators: 10,
		},
		{
			//
			//
			waitperiodsMillis:        []int{50, 50, 50, 80, 90, 90, 95, 99, 100, 120, 130, 150},
			bufferSize:           8,
			anticipatedVerdict:   MustReduce,
			anticipatedCountOperators: 9,
		},
	} {
		//
		//
		for _, waitperiodMillis := range tt.waitperiodsMillis {
			lt := time.Duration(waitperiodMillis) * time.Millisecond
			adjuster.Monitor(lt)
		}

		verdict := adjuster.Determine(countOperators, tt.bufferSize, bufferLimit)
		switch verdict {
		case MustSize:
			countOperators++
		case MustReduce:
			countOperators--
		case MustRemain:
			//
		}

		//
		//
		require.Equal(t, tt.anticipatedVerdict, verdict, "REDACTED", ordinal)
		require.Equal(t, tt.anticipatedCountOperators, countOperators, "REDACTED", ordinal)
	}

}
