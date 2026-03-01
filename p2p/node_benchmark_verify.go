package p2p

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
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/toolkits"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
)

type peer2peerSimplexSettings struct {
	interval        time.Duration
	transmitParallelism int
	acceptDeferral    time.Duration
}

//
//
func VerifyBenchmarkPeer2peerSimplex(t *testing.T) {
	toolkits.ShieldPeer2peerBenchmarkVerify(t)

	for _, tt := range []peer2peerSimplexSettings{
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
			verifyBenchmarkPeer2peerSimplex(t, tt)
			t.Log("REDACTED")
		})
	}
}

func verifyBenchmarkPeer2peerSimplex(t *testing.T, cfg peer2peerSimplexSettings) {
	t.Logf("REDACTED", cfg.interval.String())
	t.Logf("REDACTED", cfg.transmitParallelism)
	t.Logf("REDACTED", cfg.acceptDeferral.String())

	//
	const (
		benchmarkChnSample = byte(0xaa)
		capacity    = 100 * (1 << 20) //
	)

	//
	peer2peerConfig := settings.FallbackPeer2peerSettings()
	peer2peerConfig.LocationRegisterPrecise = true

	//
	//
	//

	multiplexerLinkSettings := link.FallbackModuleLinkSettings()

	//
	//
	//
	//
	//
	//
	//
	//
	//

	tracer := log.FreshNooperationTracer()

	//
	receiverHandler := FreshBenchmarkHandler(t, benchmarkChnSample, cfg.acceptDeferral)

	//
	receiverRouter := generateRouterUsingHandler(t, peer2peerConfig, multiplexerLinkSettings, "REDACTED", receiverHandler, tracer)
	require.NoError(t, receiverRouter.Initiate())
	defer func() { require.NoError(t, receiverRouter.Halt()) }()

	receiverLocation := receiverRouter.NetworkLocator()
	t.Logf("REDACTED", receiverLocation.Text())

	//
	mockHandler := FreshVerifyHandler(receiverHandler.ObtainConduits(), false)
	mockHandler.AssignTracer(tracer)

	originatorRouter := generateRouterUsingHandler(t, peer2peerConfig, multiplexerLinkSettings, "REDACTED", mockHandler, tracer)
	require.NoError(t, originatorRouter.Initiate())
	defer func() { require.NoError(t, originatorRouter.Halt()) }()

	t.Logf("REDACTED", originatorRouter.NetworkLocator().Text())

	//
	require.NoError(t, originatorRouter.CallNodeUsingLocator(receiverLocation))
	time.Sleep(100 * time.Millisecond)

	require.Equal(t, 1, receiverRouter.Nodes().Extent(), "REDACTED")
	require.Equal(t, 1, originatorRouter.Nodes().Extent(), "REDACTED")

	originatorNodes := originatorRouter.Nodes().Duplicate()
	require.Len(t, originatorNodes, 1)

	receiverNode := originatorNodes[0]

	ctx, abort := context.WithTimeout(context.Background(), cfg.interval)
	defer abort()

	var (
		initiate = time.Now()

		transmitTriumphs = atomic.Uint64{}
		transmitMishaps  = atomic.Uint64{}

		acceptTriumphs = atomic.Uint64{}

		//
		//

		acceptWaitstates = make([]time.Duration, 0, 10_000)
		handleWaitstates = make([]time.Duration, 0, 10_000)

		pauseExecution = make(chan struct{})
	)

	transmitMethod := func() {
		presentTxt := strconv.FormatInt(time.Now().UnixMicro(), 10)
		msg := kinds.TowardSolicitReverberate(presentTxt)

		relayed := receiverNode.Transmit(Wrapper{
			ConduitUUID: benchmarkChnSample,
			Signal:   msg,
		})

		if relayed {
			transmitTriumphs.Add(1)
		} else {
			transmitMishaps.Add(1)
		}
	}

	conclude := func() {
		t.Logf("REDACTED", len(receiverHandler.receiver))
		close(receiverHandler.receiver)
		<-pauseExecution
	}

	go func() {
		for log := range receiverHandler.receiver {
			acceptTriumphs.Add(1)

			req, ok := log.content.(*kinds.SolicitReverberate)
			require.True(t, ok)

			msg := strings.TrimLeft(req.Signal, "REDACTED")
			i64, err := strconv.ParseInt(msg, 10, 64)
			require.NoError(t, err, "REDACTED", msg)

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

	time.Sleep(time.Second)
	<-pauseExecution

	//
	momentSeized := time.Since(initiate)

	t.Logf("REDACTED", transmitTriumphs.Load()+transmitMishaps.Load())
	t.Logf("REDACTED", transmitTriumphs.Load(), transmitMishaps.Load())
	t.Logf("REDACTED", float64(transmitTriumphs.Load())/momentSeized.Seconds())

	t.Logf("REDACTED", acceptTriumphs.Load())
	t.Logf("REDACTED", float64(acceptTriumphs.Load())/momentSeized.Seconds())

	signalsMislaid := transmitTriumphs.Load() - acceptTriumphs.Load()
	artifactLeakageFraction := float64(signalsMislaid) / float64(transmitTriumphs.Load()+transmitMishaps.Load()) * 100

	t.Logf("REDACTED", int64(signalsMislaid), artifactLeakageFraction)

	toolkits.RecordIntervalMetrics(t, "REDACTED", acceptWaitstates)
	toolkits.RecordIntervalMetrics(t, "REDACTED", handleWaitstates)
}

//
func generateRouterUsingHandler(
	t *testing.T,
	cfg *settings.Peer2peerSettings,
	multiplexerLinkSettings link.ModuleLinkSettings,
	alias string,
	handler Handler,
	tracer log.Tracer,
) *Router {
	t.Helper()

	peerToken := PeerToken{
		PrivateToken: edwards25519.ProducePrivateToken(),
	}

	channels := toolkits.ObtainReleaseChannels(t, 1)
	locationTxt := fmt.Sprintf("REDACTED", channels[0])

	peerDetails := FallbackPeerDetails{
		Pseudonym:         alias,
		SchemeEdition: fallbackSchemeEdition,
		FallbackPeerUUID:   peerToken.ID(),
		OverhearLocation:      locationTxt,
		Fabric:         "REDACTED",
		Edition:         "REDACTED",
		Conduits:        []byte{},
	}

	location, err := FreshNetworkLocatorText(UUIDLocationText(peerToken.ID(), locationTxt))
	require.NoError(t, err)

	//
	carrier := FreshMultiplexCarrier(peerDetails, peerToken, multiplexerLinkSettings)

	//
	require.NoError(t, carrier.Overhear(*location))

	//
	//
	veritableLocation := FreshNetworkLocator(peerToken.ID(), carrier.observer.Addr())
	carrier.networkLocation = *veritableLocation

	//
	sw := FreshRouter(cfg, carrier)
	sw.AssignTracer(tracer.Using("REDACTED", alias))
	sw.AssignPeerToken(&peerToken)

	//
	sw.AppendHandler(alias, handler)
	//
	for ch := range sw.enginesViaChnl {
		peerDetails.Conduits = append(peerDetails.Conduits, ch)
	}

	carrier.peerDetails = peerDetails

	sw.AssignPeerDetails(peerDetails)

	return sw
}

type BenchmarkHandler struct {
	FoundationHandler

	t *testing.T

	conduitUUID    byte
	acceptDeferral time.Duration
	receiver         chan benchmarkLog
}

type benchmarkLog struct {
	content     proto.Message
	acceptedLocated  time.Time
	handledLocated time.Time
}

func FreshBenchmarkHandler(t *testing.T, conduitUUID byte, acceptDeferral time.Duration) *BenchmarkHandler {
	r := &BenchmarkHandler{
		t:            t,
		conduitUUID:    conduitUUID,
		acceptDeferral: acceptDeferral,
		receiver:         make(chan benchmarkLog, 1_000_000),
	}

	r.FoundationHandler = *FreshFoundationHandler("REDACTED", r)
	r.AssignTracer(log.FreshNooperationTracer())

	return r
}

func (r *BenchmarkHandler) ObtainConduits() []*link.ConduitDefinition {
	return []*link.ConduitDefinition{
		{
			ID:                  r.conduitUUID,
			Urgency:            1,
			SignalKind:         &kinds.SolicitReverberate{},
			TransmitStagingVolume:   1_000_000,
			ObtainReserveVolume:  100 * (1 << 20),
			ObtainSignalVolume: 1_000_000,
		},
	}
}

func (r *BenchmarkHandler) Accept(e Wrapper) {
	acceptedLocated := time.Now()

	//
	if r.acceptDeferral > 0 {
		time.Sleep(r.acceptDeferral)
	}

	handledLocated := time.Now()

	//
	r.receiver <- benchmarkLog{
		content:     e.Signal,
		acceptedLocated:  acceptedLocated,
		handledLocated: handledLocated,
	}
}
