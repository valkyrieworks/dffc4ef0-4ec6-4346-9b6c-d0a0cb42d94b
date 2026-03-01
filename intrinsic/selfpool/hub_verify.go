package selfpool

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func VerifyHub(t *testing.T) {
	tracer := log.VerifyingTracer()

	t.Run("REDACTED", func(t *testing.T) {
		const (
			verifyInterval = 10 * time.Second

			minimumLaborers = 4
			maximumLaborers = 10

			//
			//
			waitstateQuantile = 80.0
			waitstateLimit  = 20 * time.Millisecond

			//
			//
			waitstateMaximumVariance = waitstateLimit / 10

			//
			eraInterval = waitstateLimit * 10
		)

		//
		//
		amplifier := FreshEfficiencyWaitstateAmplifier(
			minimumLaborers,
			maximumLaborers,
			waitstateQuantile,
			waitstateLimit,
			eraInterval,
			tracer,
		)

		tracer.Details(amplifier.Text())

		//
		var (
			sized, lessened, remained atomic.Int64
			signalsBroadcasted      = atomic.Int64{}
			signalsUtilized       = atomic.Int64{}
			stagingVolume          = 1024
			shutdownVerify              = make(chan struct{})
			hub                   *Hub[time.Duration]
		)

		subscriber := func(waitstate time.Duration) {
			time.Sleep(waitstate)

			sum := signalsUtilized.Add(1)

			if sum%200 == 0 {
				qs := hub.Len()
				if qs > 0 {
					tracer.Details("REDACTED", "REDACTED", qs)
				}
			}
		}

		hub = New(
			amplifier,
			subscriber,
			stagingVolume,
			UsingTracer[time.Duration](tracer),
			UsingUponAmplify[time.Duration](func() { sized.Add(1) }),
			UsingUponReduce[time.Duration](func() { lessened.Add(1) }),
			UsingUponRemain[time.Duration](func() { remained.Add(1) }),
		)

		//
		clock := time.NewTimer(verifyInterval)
		defer clock.Stop()

		go func() {
			defer close(shutdownVerify)

			for {
				select {
				case <-clock.C:
					return
				default:
					//
					subscriberDeferral := time.Duration(
						rand.Uint64() % uint64(waitstateLimit+waitstateMaximumVariance),
					)

					hub.Propel(subscriberDeferral)
					signalsBroadcasted.Add(1)

					//
					time.Sleep(subscriberDeferral / 8)
				}
			}
		}()

		//
		//
		tracer.Details("REDACTED", "REDACTED", verifyInterval)
		hub.Initiate()

		//
		tracer.Details("REDACTED")
		<-shutdownVerify

		//
		tracer.Details("REDACTED")
		hub.Halt()

		//
		t.Logf("REDACTED", signalsBroadcasted.Load())
		t.Logf("REDACTED", signalsUtilized.Load())
		t.Logf("REDACTED", sized.Load())
		t.Logf("REDACTED", lessened.Load())
		t.Logf("REDACTED", remained.Load())

		//
		variation := float64(signalsBroadcasted.Load()) * 0.2

		require.InDelta(t, signalsUtilized.Load(), signalsBroadcasted.Load(), variation, "REDACTED")
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		//
		amplifier := FreshEfficiencyWaitstateAmplifier(
			4,
			8,
			90.0,
			10*time.Millisecond,
			20*time.Millisecond,
			tracer,
		)

		tracer.Details(amplifier.Text())

		//
		mu := sync.Mutex{}
		outcomes := []string{}

		subscriber := func(msg string) {
			mu.Lock()
			defer mu.Unlock()
			outcomes = append(outcomes, msg)
		}

		hub := New(
			amplifier,
			subscriber,
			1024,
			UsingTracer[string](tracer),
			UsingUrgencyStaging[string](FreshUrgencyStaging(10)),
		)

		hub.Initiate()
		defer hub.Halt()

		const elementsTally = 1000

		for i := 0; i < elementsTally; i++ {
			var (
				//
				urgency = 1 + i%10
				datum    = fmt.Sprintf("REDACTED", i)
			)

			//
			//
			go hub.PropelUrgency(datum, urgency)
		}

		anticipate := func() bool {
			mu.Lock()
			defer mu.Unlock()

			return len(outcomes) == elementsTally
		}

		assert.Eventually(t, anticipate, 2*time.Second, 500*time.Millisecond)
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		//
		amplifier := FreshEfficiencyWaitstateAmplifier(
			1,
			1,
			90.0,
			10*time.Millisecond,
			100*time.Millisecond, tracer,
		)

		var (
			utilized          atomic.Int64
			ledgerSubscriber     = make(chan struct{})
			subscriberReleased atomic.Bool
		)

		subscriber := func(_ int) {
			if !subscriberReleased.Load() {
				<-ledgerSubscriber
				subscriberReleased.Store(true)
			}
			utilized.Add(1)
		}

		hub := New(
			amplifier,
			subscriber,
			1,
			UsingTracer[int](tracer),
			UsingUrgencyStaging[int](FreshUrgencyStaging(10)),
		)

		hub.Initiate()
		defer hub.Halt()

		//
		//
		const sumElements = 100
		for i := 1; i <= sumElements; i++ {
			require.NoError(t, hub.PropelUrgency(i, 1))
		}

		time.Sleep(50 * time.Millisecond)
		close(ledgerSubscriber)

		//
		require.Eventually(t, func() bool {
			return utilized.Load() == sumElements
		}, 5*time.Second, 10*time.Millisecond,
			"REDACTED", sumElements, utilized.Load())
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		//
		amplifier := FreshEfficiencyWaitstateAmplifier(2, 4, 90.0, 10*time.Millisecond, 50*time.Millisecond, tracer)

		var utilized atomic.Int64

		subscriber := func(_ int) {
			utilized.Add(1)
		}

		hub := New(
			amplifier,
			subscriber,
			10,
			UsingTracer[int](tracer),
			UsingUrgencyStaging[int](FreshUrgencyStaging(10)),
		)

		hub.Initiate()
		defer hub.Halt()

		//
		//
		for i := 0; i < 50; i++ {
			require.NoError(t, hub.PropelUrgency(i, 1))
		}

		require.Eventually(t, func() bool {
			return utilized.Load() == 50
		}, 2*time.Second, 10*time.Millisecond)

		//
		for i := 50; i < 100; i++ {
			require.NoError(t, hub.PropelUrgency(i, 1))
		}

		//
		require.Eventually(t, func() bool {
			return utilized.Load() == 100
		}, 2*time.Second, 10*time.Millisecond,
			"REDACTED", utilized.Load())
	})

}
