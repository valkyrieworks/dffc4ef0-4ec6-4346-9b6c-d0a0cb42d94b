package netpeer

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/verify/utilities"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/stretchr/testify/require"
)

type netpeerOnewaySettings struct {
	period        time.Duration
	transmitParallelism int
	acceptDeferral    time.Duration
}

//
//
//
func VerifyBenchmarkNetpeerOneway(t *testing.T) {
	utilities.ShieldP2PBenchmarkVerify(t)

	for _, tt := range []netpeerOnewaySettings{
		{
			period:        10 * time.Second,
			transmitParallelism: 1,
			acceptDeferral:    0,
		},
		{
			period:        10 * time.Second,
			transmitParallelism: 8,
			acceptDeferral:    0,
		},
		{
			period:        10 * time.Second,
			transmitParallelism: 1,
			acceptDeferral:    10 * time.Millisecond,
		},
		{
			period:        10 * time.Second,
			transmitParallelism: 8,
			acceptDeferral:    10 * time.Millisecond,
		},
		{
			period:        10 * time.Second,
			transmitParallelism: 16,
			acceptDeferral:    10 * time.Millisecond,
		},
	} {
		label := fmt.Sprintf("REDACTED", tt.period, tt.transmitParallelism, tt.acceptDeferral)
		t.Run(label, func(t *testing.T) {
			runtime.GC()
			verifyBenchmarkNetpeerOneway(t, tt)
			t.Log("REDACTED")
		})
	}
}

func verifyBenchmarkNetpeerOneway(t *testing.T, cfg netpeerOnewaySettings) {
	const conduitFoo = byte(0xaa)

	t.Logf("REDACTED", cfg.period.String())
	t.Logf("REDACTED", cfg.transmitParallelism)
	t.Logf("REDACTED", cfg.acceptDeferral.String())

	//
	ctx := context.Background()

	//
	ports := utilities.FetchReleasePorts(t, 2)

	host1 := createVerifyMachine(t, ports[0])
	host2 := createVerifyMachine(t, ports[1], withOnboardNodes([]settings.LibraryP2POnboardNode{
		{
			Machine: fmt.Sprintf("REDACTED", ports[0]),
			ID:   host1.ID().String(),
		},
	}))

	t.Logf("REDACTED", host1.AddressDetails().String())
	t.Logf("REDACTED", host2.AddressDetails().String())

	establishOnboardNodes(t, ctx, host2, host2.OnboardNodes())
	t.Cleanup(func() {
		host2.Close()
		host1.Close()
	})

	type log struct {
		shipment     []byte
		err         error
		acceptedAt  time.Time
		handledAt time.Time
	}

	//
	drain := make(chan log, 1_000_000)

	//

	//
	host2member1, err := NewNode(host2, host1.AddressDetails(), p2p.NoopStats(), false, false, false)
	require.NoError(t, err)
	require.NoError(t, host2member1.Begin())

	//
	//
	ctx, revoke := context.WithTimeout(ctx, cfg.period)
	defer revoke()

	var (
		begin         = time.Now()
		transmitTriumphs = atomic.Uint64{}
		transmitBreakdowns  = atomic.Uint64{}

		acceptTriumphs = atomic.Uint64{}
		acceptBreakdowns  = atomic.Uint64{}

		acceptWaitperiods = make([]time.Duration, 0, 10_000)
		handleWaitperiods = make([]time.Duration, 0, 10_000)

		waitExecution = make(chan struct{})
	)

	protocolUID := ProtocolUID(conduitFoo)
	host1.SetStreamHandler(protocolUID, func(influx network.Stream) {
		defer func() {
			if r := recover(); r != nil {
				if strings.Contains(fmt.Sprintf("REDACTED", r), "REDACTED") {
					//
					return
				}

				panic(r)
			}
		}()

		shipment, err := InfluxFetchEnd(influx)

		acceptedAt := time.Now()

		//
		if err == nil && cfg.acceptDeferral > 0 {
			time.Sleep(cfg.acceptDeferral)
		}

		handledAt := time.Now()

		//
		drain <- log{
			shipment:     shipment,
			err:         err,
			acceptedAt:  acceptedAt,
			handledAt: handledAt,
		}
	})

	transmitFunction := func() {
		currentStr := strconv.FormatInt(time.Now().UnixMicro(), 10)
		msg := kinds.ToQueryReverberate(currentStr)

		relayed := host2member1.Transmit(p2p.Packet{
			StreamUID: conduitFoo,
			Signal:   msg,
		})

		if relayed {
			transmitTriumphs.Add(1)
		} else {
			transmitBreakdowns.Add(1)
		}
	}

	conclude := func() {
		t.Logf("REDACTED", len(drain))
		close(drain)
		<-waitExecution
	}

	//
	go func() {
		for log := range drain {
			if log.err != nil {
				acceptBreakdowns.Add(1)
				continue
			}

			acceptTriumphs.Add(1)

			msg := &kinds.Query{}
			require.NoError(t, msg.Unserialize(log.shipment))
			require.NotNil(t, msg.FetchReverberate())

			i64, err := strconv.ParseInt(msg.FetchReverberate().FetchSignal(), 10, 64)
			require.NoError(t, err)

			relayedAt := time.UnixMicro(i64)

			acceptWaitperiods = append(acceptWaitperiods, log.acceptedAt.Sub(relayedAt))
			handleWaitperiods = append(handleWaitperiods, log.handledAt.Sub(relayedAt))
		}

		t.Logf("REDACTED")

		close(waitExecution)
	}()

	t.Log("REDACTED")

Cycle:
	for {
		select {
		case <-ctx.Done():
			conclude()
			break Cycle
		default:
			//
			if cfg.transmitParallelism < 2 {
				transmitFunction()
				continue
			}

			//
			wg := sync.WaitGroup{}
			wg.Add(cfg.transmitParallelism)
			for i := 0; i < cfg.transmitParallelism; i++ {
				go func() {
					defer wg.Done()
					transmitFunction()
				}()
			}
			wg.Wait()
		}
	}

	<-waitExecution

	//
	timeSeized := time.Since(begin)

	t.Logf("REDACTED", transmitTriumphs.Load()+transmitBreakdowns.Load())
	t.Logf("REDACTED", transmitTriumphs.Load(), transmitBreakdowns.Load())
	t.Logf("REDACTED", float64(transmitTriumphs.Load())/timeSeized.Seconds())

	t.Logf("REDACTED", acceptTriumphs.Load()+acceptBreakdowns.Load())
	t.Logf("REDACTED", acceptTriumphs.Load(), acceptBreakdowns.Load())
	t.Logf("REDACTED", float64(acceptTriumphs.Load())/timeSeized.Seconds())

	signalsMissing := transmitTriumphs.Load() - acceptTriumphs.Load() - acceptBreakdowns.Load()
	signalLeakageFraction := float64(signalsMissing) / float64(transmitTriumphs.Load()+transmitBreakdowns.Load()) * 100

	t.Logf("REDACTED", int64(signalsMissing), signalLeakageFraction)

	utilities.TracePeriodMetrics(t, "REDACTED", acceptWaitperiods)
	utilities.TracePeriodMetrics(t, "REDACTED", handleWaitperiods)
}
