package selfpool

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"testing"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/toolkits"
	"github.com/stretchr/testify/require"
)

func VerifyStaging(t *testing.T) {
	q := FreshStaging()

	q.Propel(1)
	q.Propel(2)
	q.Propel(3)

	require.Equal(t, 3, q.Len())

	pop := func(anticipated int) {
		v, ok := q.Pop()
		require.True(t, ok)
		require.Equal(t, anticipated, v)
	}

	pop(1)
	pop(2)

	q.Propel(4)

	pop(3)
	pop(4)

	require.Equal(t, 0, q.Len())

	_, ok := q.Pop()
	require.False(t, ok)
}

func VerifyUrgencyStaging(t *testing.T) {
	const (
		repetitions = 100_000
		urgencies = 10
	)

	t.Run("REDACTED", func(t *testing.T) {
		//
		staging := FreshUrgencyStaging(urgencies)

		//
		entries := produceUnpredictableData(repetitions, urgencies)

		//
		intervals := []time.Duration{}

		for _, record := range entries {
			now := time.Now()

			err := staging.Propel(record.datum, int(record.urgency))
			if err != nil {
				t.Fatalf("REDACTED", err)
			}

			intervals = append(intervals, time.Since(now))
		}

		//
		toolkits.RecordIntervalMetrics(t, "REDACTED", intervals)

		t.Run("REDACTED", func(t *testing.T) {
			utilized := 0
			intervals := []time.Duration{}

			finalUtilized := time.Now()

			for {
				_, ok := staging.Pop()
				if !ok {
					break
				}

				intervals = append(intervals, time.Since(finalUtilized))

				utilized++
				if utilized == len(entries) {
					break
				}

				finalUtilized = time.Now()
			}

			toolkits.RecordIntervalMetrics(t, "REDACTED", intervals)
		})
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		staging := FreshUrgencyStaging(urgencies)

		//
		entries := produceUnpredictableData(repetitions, urgencies)

		//
		propelIntervals := make([]time.Duration, 0, repetitions)
		utilizeIntervals := make([]time.Duration, 0, repetitions)
		utilizedItems := make([]string, 0, repetitions)

		wg := sync.WaitGroup{}
		wg.Add(2)

		initiate := time.Now()

		go func() {
			defer wg.Done()
			for _, record := range entries {
				now := time.Now()
				err := staging.Propel(record.datum, int(record.urgency))
				if err != nil {
					//
					panic(err)
				}

				propelIntervals = append(propelIntervals, time.Since(now))
			}
		}()

		go func() {
			defer wg.Done()

			utilized := 0
			finalUtilized := time.Now()

			for {
				datum, ok := staging.Pop()
				if !ok {
					time.Sleep(10 * time.Millisecond)
					finalUtilized = time.Now()
					continue
				}

				utilizeIntervals = append(utilizeIntervals, time.Since(finalUtilized))
				utilizedItems = append(utilizedItems, datum.(string))
				utilized++

				if utilized == repetitions {
					return
				}

				finalUtilized = time.Now()
			}
		}()

		wg.Wait()

		//
		t.Logf("REDACTED", time.Since(initiate))
		toolkits.RecordIntervalMetrics(t, "REDACTED", propelIntervals)
		toolkits.RecordIntervalMetrics(t, "REDACTED", utilizeIntervals)

		//
		veritableItems := make(map[string]struct{}, len(utilizedItems))
		for _, datum := range utilizedItems {
			veritableItems[datum] = struct{}{}
		}

		for _, record := range entries {
			if _, ok := veritableItems[record.datum]; !ok {
				t.Fatalf("REDACTED", record.datum)
			}
		}
	})
}

type verifyData struct {
	urgency uint64
	datum    string
}

func produceUnpredictableData(tally int, urgencies uint64) []verifyData {
	out := []verifyData{}

	for i := 0; i < tally; i++ {
		out = append(out, verifyData{
			urgency: 1 + (rand.Uint64() % urgencies),
			datum:    fmt.Sprintf("REDACTED", i),
		})
	}

	return out
}
