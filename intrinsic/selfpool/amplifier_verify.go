package selfpool

import (
	"testing"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/stretchr/testify/require"
)

func VerifyYieldWaitstateAmplifier(t *testing.T) {
	//
	//
	const (
		min                 = 4
		max                 = 10
		limitQuantile = 90.0
		limitWaitstate    = 100 * time.Millisecond
		eraInterval       = time.Second

		stagingLimit = 10
	)

	tracer := log.VerifyingTracer()

	amplifier := FreshYieldWaitstateAmplifier(
		min,
		max,
		limitQuantile,
		limitWaitstate,
		eraInterval,
		tracer,
	)

	countLaborers := min

	for ordinal, tt := range []struct {
		waitstatesMSEC        []int
		stagingLength           int
		anticipatedVerdict   uint8
		anticipatedCountLaborers int
	}{
		{
			waitstatesMSEC:        []int{},
			stagingLength:           5,
			anticipatedVerdict:   MustRemain,
			anticipatedCountLaborers: min,
		},
		{
			//
			waitstatesMSEC:        []int{200},
			stagingLength:           5,
			anticipatedVerdict:   MustRemain,
			anticipatedCountLaborers: min,
		},
		{
			waitstatesMSEC:        []int{50, 50, 50},
			stagingLength:           5,
			anticipatedVerdict:   MustAmplify,
			anticipatedCountLaborers: 5,
		},
		{
			waitstatesMSEC:        []int{50, 50, 50},
			stagingLength:           5,
			anticipatedVerdict:   MustAmplify,
			anticipatedCountLaborers: 6,
		},
		{
			waitstatesMSEC:        []int{50, 50, 50, 80},
			stagingLength:           5,
			anticipatedVerdict:   MustAmplify,
			anticipatedCountLaborers: 7,
		},
		{
			waitstatesMSEC:        []int{50, 50, 50, 80},
			anticipatedVerdict:   MustAmplify,
			anticipatedCountLaborers: 8,
		},
		{
			waitstatesMSEC:        []int{50, 50, 50, 80, 90, 90},
			stagingLength:           5,
			anticipatedVerdict:   MustAmplify,
			anticipatedCountLaborers: 9,
		},
		{
			waitstatesMSEC:        []int{50, 50, 50},
			stagingLength:           5,
			anticipatedVerdict:   MustReduce,
			anticipatedCountLaborers: 8,
		},
		{
			waitstatesMSEC:        []int{50, 50, 50, 80, 90, 90, 95, 99},
			stagingLength:           5,
			anticipatedVerdict:   MustAmplify,
			anticipatedCountLaborers: 9,
		},
		{
			//
			waitstatesMSEC:        []int{50, 50, 50},
			stagingLength:           7,
			anticipatedVerdict:   MustAmplify,
			anticipatedCountLaborers: 10,
		},
		{
			//
			//
			waitstatesMSEC:        []int{50, 50, 50, 80, 90, 90, 95, 99, 100, 120, 130, 150},
			stagingLength:           8,
			anticipatedVerdict:   MustReduce,
			anticipatedCountLaborers: 9,
		},
	} {
		//
		//
		for _, waitstateMSEC := range tt.waitstatesMSEC {
			lt := time.Duration(waitstateMSEC) * time.Millisecond
			amplifier.Monitor(lt)
		}

		verdict := amplifier.Resolve(countLaborers, tt.stagingLength, stagingLimit)
		switch verdict {
		case MustAmplify:
			countLaborers++
		case MustReduce:
			countLaborers--
		case MustRemain:
			//
		}

		//
		//
		require.Equal(t, tt.anticipatedVerdict, verdict, "REDACTED", ordinal)
		require.Equal(t, tt.anticipatedCountLaborers, countLaborers, "REDACTED", ordinal)
	}

}
