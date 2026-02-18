package autosource

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/valkyrieworks/utils/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyDepository(t *testing.T) {
	tracer := log.VerifyingTracer()

	t.Run("REDACTED", func(t *testing.T) {
		const (
			verifyPeriod = 10 * time.Second

			minimumOperators = 4
			maximumOperators = 10

			//
			//
			waitperiodQuantile = 80.0
			waitperiodLimit  = 20 * time.Millisecond

			//
			//
			waitperiodMaximumVary = waitperiodLimit / 10

			//
			eraPeriod = waitperiodLimit * 10
		)

		//
		//
		adjuster := NewVelocityWaitperiodAdjuster(
			minimumOperators,
			maximumOperators,
			waitperiodQuantile,
			waitperiodLimit,
			eraPeriod,
			tracer,
		)

		tracer.Details(adjuster.String())

		//
		var (
			sized, reduced, remained atomic.Int64
			signalsIssued      = atomic.Int64{}
			signalsIngested       = atomic.Int64{}
			bufferVolume          = 1024
			endVerify              = make(chan struct{})
			depository                   *Depository[time.Duration]
		)

		ingester := func(waitperiod time.Duration) {
			time.Sleep(waitperiod)

			sum := signalsIngested.Add(1)

			if sum%200 == 0 {
				qs := depository.Len()
				if qs > 0 {
					tracer.Details("REDACTED", "REDACTED", qs)
				}
			}
		}

		depository = New(
			adjuster,
			ingester,
			bufferVolume,
			WithTracer[time.Duration](tracer),
			WithOnRatio[time.Duration](func() { sized.Add(1) }),
			WithOnReduce[time.Duration](func() { reduced.Add(1) }),
			WithOnRemain[time.Duration](func() { remained.Add(1) }),
		)

		//
		clock := time.NewTimer(verifyPeriod)
		defer clock.Stop()

		go func() {
			defer close(endVerify)

			for {
				select {
				case <-clock.C:
					return
				default:
					//
					ingesterDeferral := time.Duration(
						rand.Uint64() % uint64(waitperiodLimit+waitperiodMaximumVary),
					)

					depository.Propel(ingesterDeferral)
					signalsIssued.Add(1)

					//
					time.Sleep(ingesterDeferral / 8)
				}
			}
		}()

		//
		//
		tracer.Details("REDACTED", "REDACTED", verifyPeriod)
		depository.Begin()

		//
		tracer.Details("REDACTED")
		<-endVerify

		//
		tracer.Details("REDACTED")
		depository.Halt()

		//
		t.Logf("REDACTED", signalsIssued.Load())
		t.Logf("REDACTED", signalsIngested.Load())
		t.Logf("REDACTED", sized.Load())
		t.Logf("REDACTED", reduced.Load())
		t.Logf("REDACTED", remained.Load())

		//
		variance := float64(signalsIssued.Load()) * 0.2

		require.InDelta(t, signalsIngested.Load(), signalsIssued.Load(), variance, "REDACTED")
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		//
		adjuster := NewVelocityWaitperiodAdjuster(
			4,
			8,
			90.0,
			10*time.Millisecond,
			20*time.Millisecond,
			tracer,
		)

		tracer.Details(adjuster.String())

		//
		mu := sync.Mutex{}
		outcomes := []string{}

		ingester := func(msg string) {
			mu.Lock()
			defer mu.Unlock()
			outcomes = append(outcomes, msg)
		}

		depository := New(
			adjuster,
			ingester,
			1024,
			WithTracer[string](tracer),
			WithUrgencyBuffer[string](NewUrgencyBuffer(10)),
		)

		depository.Begin()
		defer depository.Halt()

		const itemsNumber = 1000

		for i := 0; i < itemsNumber; i++ {
			var (
				//
				urgency = 1 + i%10
				item    = fmt.Sprintf("REDACTED", i)
			)

			//
			//
			go depository.PropelUrgency(item, urgency)
		}

		anticipate := func() bool {
			mu.Lock()
			defer mu.Unlock()

			return len(outcomes) == itemsNumber
		}

		assert.Eventually(t, anticipate, 2*time.Second, 500*time.Millisecond)
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		//
		adjuster := NewVelocityWaitperiodAdjuster(
			1,
			1,
			90.0,
			10*time.Millisecond,
			100*time.Millisecond, tracer,
		)

		var (
			ingested          atomic.Int64
			ledgerIngester     = make(chan struct{})
			ingesterUnrestricted atomic.Bool
		)

		ingester := func(_ int) {
			if !ingesterUnrestricted.Load() {
				<-ledgerIngester
				ingesterUnrestricted.Store(true)
			}
			ingested.Add(1)
		}

		depository := New(
			adjuster,
			ingester,
			1,
			WithTracer[int](tracer),
			WithUrgencyBuffer[int](NewUrgencyBuffer(10)),
		)

		depository.Begin()
		defer depository.Halt()

		//
		//
		const sumItems = 100
		for i := 1; i <= sumItems; i++ {
			require.NoError(t, depository.PropelUrgency(i, 1))
		}

		time.Sleep(50 * time.Millisecond)
		close(ledgerIngester)

		//
		require.Eventually(t, func() bool {
			return ingested.Load() == sumItems
		}, 5*time.Second, 10*time.Millisecond,
			"REDACTED", sumItems, ingested.Load())
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		//
		adjuster := NewVelocityWaitperiodAdjuster(2, 4, 90.0, 10*time.Millisecond, 50*time.Millisecond, tracer)

		var ingested atomic.Int64

		ingester := func(_ int) {
			ingested.Add(1)
		}

		depository := New(
			adjuster,
			ingester,
			10,
			WithTracer[int](tracer),
			WithUrgencyBuffer[int](NewUrgencyBuffer(10)),
		)

		depository.Begin()
		defer depository.Halt()

		//
		//
		for i := 0; i < 50; i++ {
			require.NoError(t, depository.PropelUrgency(i, 1))
		}

		require.Eventually(t, func() bool {
			return ingested.Load() == 50
		}, 2*time.Second, 10*time.Millisecond)

		//
		for i := 50; i < 100; i++ {
			require.NoError(t, depository.PropelUrgency(i, 1))
		}

		//
		require.Eventually(t, func() bool {
			return ingested.Load() == 100
		}, 2*time.Second, 10*time.Millisecond,
			"REDACTED", ingested.Load())
	})

}
