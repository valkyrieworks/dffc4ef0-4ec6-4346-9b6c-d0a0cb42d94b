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

	"github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p/link"
	"github.com/valkyrieworks/verify/utilities"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
)

type p2pOnewaySettings struct {
	period        time.Duration
	transmitParallelism int
	acceptDeferral    time.Duration
}

//
//
func VerifyBenchmarkP2POneway(t *testing.T) {
	utilities.ShieldP2PBenchmarkVerify(t)

	for _, tt := range []p2pOnewaySettings{
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
			verifyBenchmarkP2POneway(t, tt)
			t.Log("REDACTED")
		})
	}
}

func verifyBenchmarkP2POneway(t *testing.T, cfg p2pOnewaySettings) {
	t.Logf("REDACTED", cfg.period.String())
	t.Logf("REDACTED", cfg.transmitParallelism)
	t.Logf("REDACTED", cfg.acceptDeferral.String())

	//
	const (
		benchmarkChannelFoo = byte(0xaa)
		throughput    = 100 * (1 << 20) //
	)

	//
	p2pConfig := settings.StandardP2PSettings()
	p2pConfig.AddressLedgerPrecise = true

	//
	//
	//

	multiplexerLinkSettings := link.StandardMLinkSettings()

	//
	//
	//
	//
	//
	//
	//
	//
	//

	tracer := log.NewNoopTracer()

	//
	receiverHandler := NewBenchmarkHandler(t, benchmarkChannelFoo, cfg.acceptDeferral)

	//
	receiverRouter := instantiateRouterWithHandler(t, p2pConfig, multiplexerLinkSettings, "REDACTED", receiverHandler, tracer)
	require.NoError(t, receiverRouter.Begin())
	defer func() { require.NoError(t, receiverRouter.Halt()) }()

	receiverAddress := receiverRouter.NetLocation()
	t.Logf("REDACTED", receiverAddress.String())

	//
	proxyHandler := NewVerifyHandler(receiverHandler.FetchStreams(), false)
	proxyHandler.AssignTracer(tracer)

	emitterRouter := instantiateRouterWithHandler(t, p2pConfig, multiplexerLinkSettings, "REDACTED", proxyHandler, tracer)
	require.NoError(t, emitterRouter.Begin())
	defer func() { require.NoError(t, emitterRouter.Halt()) }()

	t.Logf("REDACTED", emitterRouter.NetLocation().String())

	//
	require.NoError(t, emitterRouter.CallNodeWithLocation(receiverAddress))
	time.Sleep(100 * time.Millisecond)

	require.Equal(t, 1, receiverRouter.Nodes().Volume(), "REDACTED")
	require.Equal(t, 1, emitterRouter.Nodes().Volume(), "REDACTED")

	emitterNodes := emitterRouter.Nodes().Clone()
	require.Len(t, emitterNodes, 1)

	receiverNode := emitterNodes[0]

	ctx, revoke := context.WithTimeout(context.Background(), cfg.period)
	defer revoke()

	var (
		begin = time.Now()

		transmitTriumphs = atomic.Uint64{}
		transmitBreakdowns  = atomic.Uint64{}

		acceptTriumphs = atomic.Uint64{}

		//
		//

		acceptWaitperiods = make([]time.Duration, 0, 10_000)
		handleWaitperiods = make([]time.Duration, 0, 10_000)

		waitExecution = make(chan struct{})
	)

	transmitFunction := func() {
		currentStr := strconv.FormatInt(time.Now().UnixMicro(), 10)
		msg := kinds.ToQueryReverberate(currentStr)

		relayed := receiverNode.Transmit(Packet{
			StreamUID: benchmarkChannelFoo,
			Signal:   msg,
		})

		if relayed {
			transmitTriumphs.Add(1)
		} else {
			transmitBreakdowns.Add(1)
		}
	}

	conclude := func() {
		t.Logf("REDACTED", len(receiverHandler.drain))
		close(receiverHandler.drain)
		<-waitExecution
	}

	go func() {
		for log := range receiverHandler.drain {
			acceptTriumphs.Add(1)

			req, ok := log.shipment.(*kinds.QueryReverberate)
			require.True(t, ok)

			msg := strings.TrimLeft(req.Signal, "REDACTED")
			i64, err := strconv.ParseInt(msg, 10, 64)
			require.NoError(t, err, "REDACTED", msg)

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

	time.Sleep(time.Second)
	<-waitExecution

	//
	timeSeized := time.Since(begin)

	t.Logf("REDACTED", transmitTriumphs.Load()+transmitBreakdowns.Load())
	t.Logf("REDACTED", transmitTriumphs.Load(), transmitBreakdowns.Load())
	t.Logf("REDACTED", float64(transmitTriumphs.Load())/timeSeized.Seconds())

	t.Logf("REDACTED", acceptTriumphs.Load())
	t.Logf("REDACTED", float64(acceptTriumphs.Load())/timeSeized.Seconds())

	signalsMissing := transmitTriumphs.Load() - acceptTriumphs.Load()
	signalLeakageFraction := float64(signalsMissing) / float64(transmitTriumphs.Load()+transmitBreakdowns.Load()) * 100

	t.Logf("REDACTED", int64(signalsMissing), signalLeakageFraction)

	utilities.TracePeriodMetrics(t, "REDACTED", acceptWaitperiods)
	utilities.TracePeriodMetrics(t, "REDACTED", handleWaitperiods)
}

//
func instantiateRouterWithHandler(
	t *testing.T,
	cfg *settings.P2PSettings,
	multiplexerLinkSettings link.MLinkSettings,
	label string,
	handler Handler,
	tracer log.Tracer,
) *Router {
	t.Helper()

	memberKey := MemberKey{
		PrivateKey: ed25519.GeneratePrivateKey(),
	}

	ports := utilities.FetchReleasePorts(t, 1)
	addressStr := fmt.Sprintf("REDACTED", ports[0])

	memberDetails := StandardMemberDetails{
		Moniker:         label,
		ProtocolRelease: standardProtocolRelease,
		StandardMemberUID:   memberKey.ID(),
		ObserveAddress:      addressStr,
		Fabric:         "REDACTED",
		Release:         "REDACTED",
		Streams:        []byte{},
	}

	address, err := NewNetLocationString(UIDLocationString(memberKey.ID(), addressStr))
	require.NoError(t, err)

	//
	carrier := NewMulticastCarrier(memberDetails, memberKey, multiplexerLinkSettings)

	//
	require.NoError(t, carrier.Observe(*address))

	//
	//
	factualAddress := NewNetLocation(memberKey.ID(), carrier.observer.Addr())
	carrier.netAddress = *factualAddress

	//
	sw := NewRouter(cfg, carrier)
	sw.AssignTracer(tracer.With("REDACTED", label))
	sw.CollectionMemberKey(&memberKey)

	//
	sw.AppendHandler(label, handler)
	//
	for ch := range sw.handlersByChan {
		memberDetails.Streams = append(memberDetails.Streams, ch)
	}

	carrier.memberDetails = memberDetails

	sw.CollectionMemberDetails(memberDetails)

	return sw
}

type BenchmarkHandler struct {
	RootHandler

	t *testing.T

	conduitUID    byte
	acceptDeferral time.Duration
	drain         chan benchmarkLog
}

type benchmarkLog struct {
	shipment     proto.Message
	acceptedAt  time.Time
	handledAt time.Time
}

func NewBenchmarkHandler(t *testing.T, conduitUID byte, acceptDeferral time.Duration) *BenchmarkHandler {
	r := &BenchmarkHandler{
		t:            t,
		conduitUID:    conduitUID,
		acceptDeferral: acceptDeferral,
		drain:         make(chan benchmarkLog, 1_000_000),
	}

	r.RootHandler = *NewRootHandler("REDACTED", r)
	r.AssignTracer(log.NewNoopTracer())

	return r
}

func (r *BenchmarkHandler) FetchStreams() []*link.StreamDefinition {
	return []*link.StreamDefinition{
		{
			ID:                  r.conduitUID,
			Urgency:            1,
			SignalKind:         &kinds.QueryReverberate{},
			TransmitBufferVolume:   1_000_000,
			AcceptBufferVolume:  100 * (1 << 20),
			AcceptSignalVolume: 1_000_000,
		},
	}
}

func (r *BenchmarkHandler) Accept(e Packet) {
	acceptedAt := time.Now()

	//
	if r.acceptDeferral > 0 {
		time.Sleep(r.acceptDeferral)
	}

	handledAt := time.Now()

	//
	r.drain <- benchmarkLog{
		shipment:     e.Signal,
		acceptedAt:  acceptedAt,
		handledAt: handledAt,
	}
}
