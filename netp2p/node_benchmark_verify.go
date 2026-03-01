package netp2p

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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/toolkits"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/stretchr/testify/require"
)

type netp2pSimplexSettings struct {
	interval        time.Duration
	transmitParallelism int
	acceptDeferral    time.Duration
}

//
//
//
func VerifyBenchmarkNetp2pSimplex(t *testing.T) {
	toolkits.ShieldPeer2peerBenchmarkVerify(t)

	for _, tt := range []netp2pSimplexSettings{
		{
			interval:        10 * time.Second,
			transmitParallelism: 1,
			acceptDeferral:    0,
		},
		{
			interval:        10 * time.Second,
			transmitParallelism: 8,
			acceptDeferral:    0,
		},
		{
			interval:        10 * time.Second,
			transmitParallelism: 1,
			acceptDeferral:    10 * time.Millisecond,
		},
		{
			interval:        10 * time.Second,
			transmitParallelism: 8,
			acceptDeferral:    10 * time.Millisecond,
		},
		{
			interval:        10 * time.Second,
			transmitParallelism: 16,
			acceptDeferral:    10 * time.Millisecond,
		},
	} {
		alias := fmt.Sprintf("REDACTED", tt.interval, tt.transmitParallelism, tt.acceptDeferral)
		t.Run(alias, func(t *testing.T) {
			runtime.GC()
			verifyBenchmarkNetp2pSimplex(t, tt)
			t.Log("REDACTED")
		})
	}
}

func verifyBenchmarkNetp2pSimplex(t *testing.T, cfg netp2pSimplexSettings) {
	const conduitSample = byte(0xaa)

	t.Logf("REDACTED", cfg.interval.String())
	t.Logf("REDACTED", cfg.transmitParallelism)
	t.Logf("REDACTED", cfg.acceptDeferral.String())

	//
	ctx := context.Background()

	//
	channels := toolkits.ObtainReleaseChannels(t, 2)

	machine1 := createVerifyMachine(t, channels[0])
	machine2 := createVerifyMachine(t, channels[1], usingOnboardNodes([]settings.LibraryPeer2peerInitiateNode{
		{
			Machine: fmt.Sprintf("REDACTED", channels[0]),
			ID:   machine1.ID().String(),
		},
	}))

	t.Logf("REDACTED", machine1.LocationDetails().String())
	t.Logf("REDACTED", machine2.LocationDetails().String())

	relateOnboardNodes(t, ctx, machine2, machine2.InitiateNodes())
	t.Cleanup(func() {
		machine2.Close()
		machine1.Close()
	})

	type log struct {
		content     []byte
		err         error
		acceptedLocated  time.Time
		handledLocated time.Time
	}

	//
	receiver := make(chan log, 1_000_000)

	//

	//
	machine2peer1, err := FreshNode(machine2, machine1.LocationDetails(), p2p.NooperationTelemetry(), false, false, false)
	require.NoError(t, err)
	require.NoError(t, machine2peer1.Initiate())

	//
	//
	ctx, abort := context.WithTimeout(ctx, cfg.interval)
	defer abort()

	var (
		initiate         = time.Now()
		transmitTriumphs = atomic.Uint64{}
		transmitMishaps  = atomic.Uint64{}

		acceptTriumphs = atomic.Uint64{}
		acceptMishaps  = atomic.Uint64{}

		acceptWaitstates = make([]time.Duration, 0, 10_000)
		handleWaitstates = make([]time.Duration, 0, 10_000)

		pauseExecution = make(chan struct{})
	)

	schemeUUID := SchemeUUID(conduitSample)
	machine1.SetStreamHandler(schemeUUID, func(influx network.Stream) {
		defer func() {
			if r := recover(); r != nil {
				if strings.Contains(fmt.Sprintf("REDACTED", r), "REDACTED") {
					//
					return
				}

				panic(r)
			}
		}()

		content, err := InfluxFetchShutdown(influx)

		acceptedLocated := time.Now()

		//
		if err == nil && cfg.acceptDeferral > 0 {
			time.Sleep(cfg.acceptDeferral)
		}

		handledLocated := time.Now()

		//
		receiver <- log{
			content:     content,
			err:         err,
			acceptedLocated:  acceptedLocated,
			handledLocated: handledLocated,
		}
	})

	transmitMethod := func() {
		presentTxt := strconv.FormatInt(time.Now().UnixMicro(), 10)
		msg := kinds.TowardSolicitReverberate(presentTxt)

		relayed := machine2peer1.Transmit(p2p.Wrapper{
			ConduitUUID: conduitSample,
			Signal:   msg,
		})

		if relayed {
			transmitTriumphs.Add(1)
		} else {
			transmitMishaps.Add(1)
		}
	}

	conclude := func() {
		t.Logf("REDACTED", len(receiver))
		close(receiver)
		<-pauseExecution
	}

	//
	go func() {
		for log := range receiver {
			if log.err != nil {
				acceptMishaps.Add(1)
				continue
			}

			acceptTriumphs.Add(1)

			msg := &kinds.Solicit{}
			require.NoError(t, msg.Decode(log.content))
			require.NotNil(t, msg.ObtainReverberate())

			i64, err := strconv.ParseInt(msg.ObtainReverberate().ObtainArtifact(), 10, 64)
			require.NoError(t, err)

			relayedLocated := time.UnixMicro(i64)

			acceptWaitstates = append(acceptWaitstates, log.acceptedLocated.Sub(relayedLocated))
			handleWaitstates = append(handleWaitstates, log.handledLocated.Sub(relayedLocated))
		}

		t.Logf("REDACTED")

		close(pauseExecution)
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
				transmitMethod()
				continue
			}

			//
			wg := sync.WaitGroup{}
			wg.Add(cfg.transmitParallelism)
			for i := 0; i < cfg.transmitParallelism; i++ {
				go func() {
					defer wg.Done()
					transmitMethod()
				}()
			}
			wg.Wait()
		}
	}

	<-pauseExecution

	//
	momentSeized := time.Since(initiate)

	t.Logf("REDACTED", transmitTriumphs.Load()+transmitMishaps.Load())
	t.Logf("REDACTED", transmitTriumphs.Load(), transmitMishaps.Load())
	t.Logf("REDACTED", float64(transmitTriumphs.Load())/momentSeized.Seconds())

	t.Logf("REDACTED", acceptTriumphs.Load()+acceptMishaps.Load())
	t.Logf("REDACTED", acceptTriumphs.Load(), acceptMishaps.Load())
	t.Logf("REDACTED", float64(acceptTriumphs.Load())/momentSeized.Seconds())

	signalsMislaid := transmitTriumphs.Load() - acceptTriumphs.Load() - acceptMishaps.Load()
	artifactLeakageFraction := float64(signalsMislaid) / float64(transmitTriumphs.Load()+transmitMishaps.Load()) * 100

	t.Logf("REDACTED", int64(signalsMislaid), artifactLeakageFraction)

	toolkits.RecordIntervalMetrics(t, "REDACTED", acceptWaitstates)
	toolkits.RecordIntervalMetrics(t, "REDACTED", handleWaitstates)
}
