package autosource

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"testing"
	"time"

	"github.com/valkyrieworks/verify/utilities"
	"github.com/stretchr/testify/require"
)

func VerifyBuffer(t *testing.T) {
	q := NewBuffer()

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

func VerifyUrgencyBuffer(t *testing.T) {
	const (
		repetitions = 100_000
		urgencies = 10
	)

	t.Run("REDACTED", func(t *testing.T) {
		//
		buffer := NewUrgencyBuffer(urgencies)

		//
		entries := generateArbitraryData(repetitions, urgencies)

		//
		periods := []time.Duration{}

		for _, item := range entries {
			now := time.Now()

			err := buffer.Propel(item.item, int(item.urgency))
			if err != nil {
				t.Fatalf("REDACTED", err)
			}

			periods = append(periods, time.Since(now))
		}

		//
		utilities.TracePeriodMetrics(t, "REDACTED", periods)

		t.Run("REDACTED", func(t *testing.T) {
			ingested := 0
			periods := []time.Duration{}

			finalIngested := time.Now()

			for {
				_, ok := buffer.Pop()
				if !ok {
					break
				}

				periods = append(periods, time.Since(finalIngested))

				ingested++
				if ingested == len(entries) {
					break
				}

				finalIngested = time.Now()
			}

			utilities.TracePeriodMetrics(t, "REDACTED", periods)
		})
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		buffer := NewUrgencyBuffer(urgencies)

		//
		entries := generateArbitraryData(repetitions, urgencies)

		//
		propelPeriods := make([]time.Duration, 0, repetitions)
		ingestPeriods := make([]time.Duration, 0, repetitions)
		ingestedItems := make([]string, 0, repetitions)

		wg := sync.WaitGroup{}
		wg.Add(2)

		begin := time.Now()

		go func() {
			defer wg.Done()
			for _, item := range entries {
				now := time.Now()
				err := buffer.Propel(item.item, int(item.urgency))
				if err != nil {
					//
					panic(err)
				}

				propelPeriods = append(propelPeriods, time.Since(now))
			}
		}()

		go func() {
			defer wg.Done()

			ingested := 0
			finalIngested := time.Now()

			for {
				item, ok := buffer.Pop()
				if !ok {
					time.Sleep(10 * time.Millisecond)
					finalIngested = time.Now()
					continue
				}

				ingestPeriods = append(ingestPeriods, time.Since(finalIngested))
				ingestedItems = append(ingestedItems, item.(string))
				ingested++

				if ingested == repetitions {
					return
				}

				finalIngested = time.Now()
			}
		}()

		wg.Wait()

		//
		t.Logf("REDACTED", time.Since(begin))
		utilities.TracePeriodMetrics(t, "REDACTED", propelPeriods)
		utilities.TracePeriodMetrics(t, "REDACTED", ingestPeriods)

		//
		factualItems := make(map[string]struct{}, len(ingestedItems))
		for _, item := range ingestedItems {
			factualItems[item] = struct{}{}
		}

		for _, item := range entries {
			if _, ok := factualItems[item.item]; !ok {
				t.Fatalf("REDACTED", item.item)
			}
		}
	})
}

type verifyData struct {
	urgency uint64
	item    string
}

func generateArbitraryData(tally int, urgencies uint64) []verifyData {
	out := []verifyData{}

	for i := 0; i < tally; i++ {
		out = append(out, verifyData{
			urgency: 1 + (rand.Uint64() % urgencies),
			item:    fmt.Sprintf("REDACTED", i),
		})
	}

	return out
}
