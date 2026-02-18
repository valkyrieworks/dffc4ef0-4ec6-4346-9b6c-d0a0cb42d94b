package p2p

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/log"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/p2p/link"
	p2pproto "github.com/valkyrieworks/schema/consensuscore/p2p"
)

var cfg *settings.P2PSettings

func init() {
	cfg = settings.StandardP2PSettings()
	cfg.PexHandler = true
	cfg.PermitReplicatedIP = true
}

type NodeSignal struct {
	Payloads proto.Message
	Tally  int
}

type VerifyHandler struct {
	RootHandler

	mtx          engineconnect.Lock
	streams     []*link.StreamDefinition
	traceSignals  bool
	noticesTally  int
	noticesAccepted map[byte][]NodeSignal
}

func NewVerifyHandler(streams []*link.StreamDefinition, traceSignals bool) *VerifyHandler {
	tr := &VerifyHandler{
		streams:     streams,
		traceSignals:  traceSignals,
		noticesAccepted: make(map[byte][]NodeSignal),
	}
	tr.RootHandler = *NewRootHandler("REDACTED", tr)
	tr.AssignTracer(log.VerifyingTracer())
	return tr
}

func (tr *VerifyHandler) FetchStreams() []*link.StreamDefinition {
	return tr.streams
}

func (tr *VerifyHandler) AppendNode(Node) {}

func (tr *VerifyHandler) DeleteNode(Node, any) {}

func (tr *VerifyHandler) Accept(e Packet) {
	if tr.traceSignals {
		tr.mtx.Lock()
		defer tr.mtx.Unlock()
		fmt.Printf("REDACTED", e.StreamUID, e.Signal)
		tr.noticesAccepted[e.StreamUID] = append(tr.noticesAccepted[e.StreamUID], NodeSignal{Payloads: e.Signal, Tally: tr.noticesTally})
		tr.noticesTally++
	}
}

func (tr *VerifyHandler) fetchNotices(chanUID byte) []NodeSignal {
	tr.mtx.Lock()
	defer tr.mtx.Unlock()
	return tr.noticesAccepted[chanUID]
}

//

//
//
func CreateRouterCouple(initRouter func(int, *Router) *Router) (*Router, *Router) {
	//
	routers := CreateLinkedRouters(cfg, 2, initRouter, Connect2routers)
	return routers[0], routers[1]
}

func initRouterFunction(_ int, sw *Router) *Router {
	sw.CollectionAddressLedger(&AddressLedgerEmulate{
		Locations:    make(map[string]struct{}),
		OurLocations: make(map[string]struct{}),
	})

	//
	sw.AppendHandler("REDACTED", NewVerifyHandler([]*link.StreamDefinition{
		{ID: byte(0x00), Urgency: 10, SignalKind: &p2pproto.Signal{}},
		{ID: byte(0x01), Urgency: 10, SignalKind: &p2pproto.Signal{}},
	}, true))
	sw.AppendHandler("REDACTED", NewVerifyHandler([]*link.StreamDefinition{
		{ID: byte(0x02), Urgency: 10, SignalKind: &p2pproto.Signal{}},
		{ID: byte(0x03), Urgency: 10, SignalKind: &p2pproto.Signal{}},
	}, true))

	return sw
}

func VerifyRouters(t *testing.T) {
	s1, s2 := CreateRouterCouple(initRouterFunction)
	t.Cleanup(func() {
		if err := s2.Halt(); err != nil {
			t.Error(err)
		}
		if err := s1.Halt(); err != nil {
			t.Error(err)
		}
	})

	if s1.Nodes().Volume() != 1 {
		t.Errorf("REDACTED", s1.Nodes().Volume())
	}
	if s2.Nodes().Volume() != 1 {
		t.Errorf("REDACTED", s2.Nodes().Volume())
	}

	//
	ch0signal := &p2pproto.PexLocations{
		Locations: []p2pproto.NetLocation{
			{
				ID: "REDACTED",
			},
		},
	}
	ch1signal := &p2pproto.PexLocations{
		Locations: []p2pproto.NetLocation{
			{
				ID: "REDACTED",
			},
		},
	}
	ch2signal := &p2pproto.PexLocations{
		Locations: []p2pproto.NetLocation{
			{
				ID: "REDACTED",
			},
		},
	}
	//
	//
	s1.MulticastAsync(Packet{StreamUID: byte(0x00), Signal: ch0signal})
	s1.MulticastAsync(Packet{StreamUID: byte(0x01), Signal: ch1signal})
	s1.AttemptMulticast(Packet{StreamUID: byte(0x02), Signal: ch2signal})

	fetchHandler := func(label string) *VerifyHandler {
		r, ok := s2.Handler(label)
		require.True(t, ok)

		tr, ok := r.(*VerifyHandler)
		require.True(t, ok)

		return tr
	}

	handlerFoo := fetchHandler("REDACTED")
	handlerBar := fetchHandler("REDACTED")

	affirmMessageAcceptedWithDeadline(t, ch0signal, byte(0x00), handlerFoo, 200*time.Millisecond, 5*time.Second)
	affirmMessageAcceptedWithDeadline(t, ch1signal, byte(0x01), handlerFoo, 200*time.Millisecond, 5*time.Second)
	affirmMessageAcceptedWithDeadline(t, ch2signal, byte(0x02), handlerBar, 200*time.Millisecond, 5*time.Second)
}

func affirmMessageAcceptedWithDeadline(
	t *testing.T,
	msg proto.Message,
	conduit byte,
	handler *VerifyHandler,
	inspectDuration,
	deadline time.Duration,
) {
	timer := time.NewTicker(inspectDuration)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			notices := handler.fetchNotices(conduit)
			if len(notices) == 0 {
				t.Fatalf("REDACTED", conduit, len(notices))
			}

			anticipatedOctets, err := proto.Marshal(notices[0].Payloads)
			require.NoError(t, err)
			acquiredOctets, err := proto.Marshal(msg)
			require.NoError(t, err)
			if len(notices) > 0 {
				if !bytes.Equal(anticipatedOctets, acquiredOctets) {
					t.Fatalf("REDACTED", msg, notices[0].Tally)
				}
				return
			}

		case <-time.After(deadline):
			t.Fatalf("REDACTED", conduit)
		}
	}
}

func VerifyRouterScreensOutSelf(t *testing.T) {
	s1 := CreateRouter(cfg, 1, initRouterFunction)

	//
	rp := &distantNode{PrivateKey: s1.memberKey.PrivateKey, Settings: cfg}
	rp.Begin()

	//
	err := s1.CallNodeWithLocation(rp.Address())
	if assert.Error(t, err) {
		if err, ok := err.(ErrDeclined); ok {
			if !err.IsEgo() {
				t.Errorf("REDACTED")
			}
		} else {
			t.Errorf("REDACTED")
		}
	}

	assert.True(t, s1.addressLedger.OurLocation(rp.Address()))
	assert.False(t, s1.addressLedger.HasLocation(rp.Address()))

	rp.Halt()

	affirmNoNodesAfterDeadline(t, s1, 100*time.Millisecond)
}

func VerifyRouterNodeRefine(t *testing.T) {
	var (
		screens = []NodeRefineFunction{
			func(_ IDXNodeCollection, _ Node) error { return nil },
			func(_ IDXNodeCollection, _ Node) error { return fmt.Errorf("REDACTED") },
			func(_ IDXNodeCollection, _ Node) error { return nil },
		}
		sw = CreateRouter(
			cfg,
			1,
			initRouterFunction,
			RouterNodeScreens(screens...),
		)
	)
	err := sw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
	rp.Begin()
	t.Cleanup(rp.Halt)

	p, err := sw.carrier.Call(*rp.Address(), nodeSettings{
		chanTraits:      sw.chanTraits,
		onNodeFault:  sw.HaltNodeForFault,
		isDurable: sw.IsNodeDurable,
		handlersByChan: sw.handlersByChan,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = sw.appendNode(p)
	if err, ok := err.(ErrDeclined); ok {
		if !err.IsScreened() {
			t.Errorf("REDACTED")
		}
	} else {
		t.Errorf("REDACTED")
	}
}

func VerifyRouterNodeRefineDeadline(t *testing.T) {
	var (
		screens = []NodeRefineFunction{
			func(_ IDXNodeCollection, _ Node) error {
				time.Sleep(10 * time.Millisecond)
				return nil
			},
		}
		sw = CreateRouter(
			cfg,
			1,
			initRouterFunction,
			RouterRefineDeadline(5*time.Millisecond),
			RouterNodeScreens(screens...),
		)
	)
	err := sw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Log(err)
		}
	})

	//
	rp := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
	rp.Begin()
	defer rp.Halt()

	p, err := sw.carrier.Call(*rp.Address(), nodeSettings{
		chanTraits:      sw.chanTraits,
		onNodeFault:  sw.HaltNodeForFault,
		isDurable: sw.IsNodeDurable,
		handlersByChan: sw.handlersByChan,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = sw.appendNode(p)
	if _, ok := err.(ErrRefineDeadline); !ok {
		t.Errorf("REDACTED")
	}
}

func VerifyRouterNodeRefineReplicated(t *testing.T) {
	sw := CreateRouter(cfg, 1, initRouterFunction)
	err := sw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
	rp.Begin()
	defer rp.Halt()

	p, err := sw.carrier.Call(*rp.Address(), nodeSettings{
		chanTraits:      sw.chanTraits,
		onNodeFault:  sw.HaltNodeForFault,
		isDurable: sw.IsNodeDurable,
		handlersByChan: sw.handlersByChan,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := sw.appendNode(p); err != nil {
		t.Fatal(err)
	}

	err = sw.appendNode(p)
	if errRej, ok := err.(ErrDeclined); ok {
		if !errRej.IsReplicated() {
			t.Errorf("REDACTED", errRej)
		}
	} else {
		t.Errorf("REDACTED", err)
	}
}

func affirmNoNodesAfterDeadline(t *testing.T, sw *Router, deadline time.Duration) {
	time.Sleep(deadline)
	if sw.Nodes().Volume() != 0 {
		t.Fatalf("REDACTED", sw, sw.Nodes().Volume())
	}
}

func VerifyRouterHaltsNotDurableNodeOnFault(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	sw := CreateRouter(cfg, 1, initRouterFunction)
	err := sw.Begin()
	if err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
	rp.Begin()
	defer rp.Halt()

	p, err := sw.carrier.Call(*rp.Address(), nodeSettings{
		chanTraits:      sw.chanTraits,
		onNodeFault:  sw.HaltNodeForFault,
		isDurable: sw.IsNodeDurable,
		handlersByChan: sw.handlersByChan,
	})
	require.Nil(err)

	err = sw.appendNode(p)
	require.Nil(err)

	require.NotNil(sw.Nodes().Get(rp.ID()))

	//
	err = p.(*node).EndLink()
	require.NoError(err)

	affirmNoNodesAfterDeadline(t, sw, 100*time.Millisecond)
	assert.False(p.IsActive())
}

func VerifyRouterHaltNodeForFault(t *testing.T) {
	s := httptest.NewServer(promhttp.Handler())
	defer s.Close()

	extractStats := func() string {
		reply, err := http.Get(s.URL)
		require.NoError(t, err)
		defer reply.Body.Close()
		buf, _ := io.ReadAll(reply.Body)
		return string(buf)
	}

	scope, component, label := settings.VerifyTelemetrySettings().Scope, StatsComponent, "REDACTED"
	re := regexp.MustCompile(scope + "REDACTED" + component + "REDACTED" + label + "REDACTED")
	nodesIndicatorItem := func() float64 {
		aligns := re.FindStringSubmatch(extractStats())
		f, _ := strconv.ParseFloat(aligns[1], 64)
		return f
	}

	p2pStats := MonitorstatsStats(scope)

	//
	sw1, sw2 := CreateRouterCouple(func(i int, sw *Router) *Router {
		//
		if i == 0 {
			opt := WithStats(p2pStats)
			opt(sw)
		}
		return initRouterFunction(i, sw)
	})

	assert.Len(t, sw1.Nodes().Clone(), 1)
	assert.EqualValues(t, 1, nodesIndicatorItem())

	//
	p := sw1.Nodes().Clone()[0]
	p.Transmit(Packet{
		StreamUID: 0x1,
		Signal:   &p2pproto.Signal{},
	})

	//
	//
	t.Cleanup(func() {
		if err := sw2.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	sw1.HaltNodeForFault(p, fmt.Errorf("REDACTED"))

	require.Empty(t, len(sw1.Nodes().Clone()), 0)
	assert.EqualValues(t, 0, nodesIndicatorItem())
}

func VerifyRouterReestablishesToOutgoingDurableNode(t *testing.T) {
	sw := CreateRouter(cfg, 1, initRouterFunction)
	err := sw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
	rp.Begin()
	defer rp.Halt()

	err = sw.AppendDurableNodes([]string{rp.Address().String()})
	require.NoError(t, err)

	err = sw.CallNodeWithLocation(rp.Address())
	require.Nil(t, err)
	require.NotNil(t, sw.Nodes().Get(rp.ID()))

	p := sw.Nodes().Clone()[0]
	err = p.(*node).EndLink()
	require.NoError(t, err)

	waitUntilRouterHasAtMinimumNNodes(sw, 1)
	assert.False(t, p.IsActive())        //
	assert.Equal(t, 1, sw.Nodes().Volume()) //

	//
	rp = &distantNode{
		PrivateKey: ed25519.GeneratePrivateKey(),
		Settings:  cfg,
		//
		//
		acceptAddress: "REDACTED",
	}
	rp.Begin()
	defer rp.Halt()

	cfg := settings.StandardP2PSettings()
	cfg.VerifyCallAbort = true //
	err = sw.appendOutgoingNodeWithSettings(rp.Address(), cfg)
	require.NotNil(t, err)
	//
	waitUntilRouterHasAtMinimumNNodes(sw, 2)
	assert.Equal(t, 2, sw.Nodes().Volume())
}

func VerifyRouterReestablishesToIncomingDurableNode(t *testing.T) {
	sw := CreateRouter(cfg, 1, initRouterFunction)
	err := sw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
	rp.Begin()
	defer rp.Halt()

	err = sw.AppendDurableNodes([]string{rp.Address().String()})
	require.NoError(t, err)

	link, err := rp.Call(sw.NetLocation())
	require.NoError(t, err)
	time.Sleep(50 * time.Millisecond)
	require.NotNil(t, sw.Nodes().Get(rp.ID()))

	link.Close()

	waitUntilRouterHasAtMinimumNNodes(sw, 1)
	assert.Equal(t, 1, sw.Nodes().Volume())
}

func VerifyRouterCallNodesAsync(t *testing.T) {
	if testing.Short() {
		return
	}

	sw := CreateRouter(cfg, 1, initRouterFunction)
	err := sw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	rp := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
	rp.Begin()
	defer rp.Halt()

	err = sw.CallNodesAsync([]string{rp.Address().String()})
	require.NoError(t, err)
	time.Sleep(callShufflerCadenceMilliseconds * time.Millisecond)
	require.NotNil(t, sw.Nodes().Get(rp.ID()))
}

func waitUntilRouterHasAtMinimumNNodes(sw *Router, n int) {
	for i := 0; i < 20; i++ {
		time.Sleep(250 * time.Millisecond)
		has := sw.Nodes().Volume()
		if has >= n {
			break
		}
	}
}

func VerifyRouterCompleteReachability(t *testing.T) {
	routers := CreateLinkedRouters(cfg, 3, initRouterFunction, Connect2routers)
	defer func() {
		for _, sw := range routers {
			t.Cleanup(func() {
				if err := sw.Halt(); err != nil {
					t.Error(err)
				}
			})
		}
	}()

	for i, sw := range routers {
		if sw.Nodes().Volume() != 2 {
			t.Fatalf("REDACTED", sw.Nodes().Volume(), i)
		}
	}
}

func VerifyRouterAllowProcedure(t *testing.T) {
	cfg.MaximumCountIncomingNodes = 5

	//
	const absoluteNodesCount = 2
	var (
		absoluteNodes   = make([]*distantNode, absoluteNodesCount)
		absoluteNodeIDXDatastore = make([]string, absoluteNodesCount)
	)
	for i := 0; i < absoluteNodesCount; i++ {
		node := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
		node.Begin()
		absoluteNodes[i] = node
		absoluteNodeIDXDatastore[i] = string(node.ID())
	}

	//
	sw := CreateRouter(cfg, 1, initRouterFunction)
	err := sw.AppendAbsoluteNodeIDXDatastore(absoluteNodeIDXDatastore)
	require.NoError(t, err)
	err = sw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		err := sw.Halt()
		require.NoError(t, err)
	})

	//
	assert.Equal(t, 0, sw.Nodes().Volume())

	//
	nodes := make([]*distantNode, 0)
	for i := 0; i < cfg.MaximumCountIncomingNodes; i++ {
		node := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
		nodes = append(nodes, node)
		node.Begin()
		c, err := node.Call(sw.NetLocation())
		require.NoError(t, err)
		//
		go func(c net.Conn) {
			for {
				one := make([]byte, 1)
				_, err := c.Read(one)
				if err != nil {
					return
				}
			}
		}(c)
	}
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, cfg.MaximumCountIncomingNodes, sw.Nodes().Volume())

	//
	node := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
	node.Begin()
	link, err := node.Call(sw.NetLocation())
	require.NoError(t, err)
	//
	one := make([]byte, 1)
	_ = link.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
	_, err = link.Read(one)
	assert.Error(t, err)
	assert.Equal(t, cfg.MaximumCountIncomingNodes, sw.Nodes().Volume())
	node.Halt()

	//
	for _, node := range absoluteNodes {
		c, err := node.Call(sw.NetLocation())
		require.NoError(t, err)
		//
		go func(c net.Conn) {
			for {
				one := make([]byte, 1)
				_, err := c.Read(one)
				if err != nil {
					return
				}
			}
		}(c)
	}
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, cfg.MaximumCountIncomingNodes+absoluteNodesCount, sw.Nodes().Volume())

	for _, node := range nodes {
		node.Halt()
	}
	for _, node := range absoluteNodes {
		node.Halt()
	}
}

type faultCarrier struct {
	allowErr error
}

func (et faultCarrier) NetLocation() NetLocation {
	panic("REDACTED")
}

func (et faultCarrier) Allow(nodeSettings) (Node, error) {
	return nil, et.allowErr
}

func (faultCarrier) Call(NetLocation, nodeSettings) (Node, error) {
	panic("REDACTED")
}

func (faultCarrier) Sanitize(Node) {
	panic("REDACTED")
}

func VerifyRouterAllowProcedureFaultScenarios(t *testing.T) {
	sw := NewRouter(cfg, faultCarrier{ErrRefineDeadline{}})
	assert.NotPanics(t, func() {
		err := sw.Begin()
		require.NoError(t, err)
		err = sw.Halt()
		require.NoError(t, err)
	})

	sw = NewRouter(cfg, faultCarrier{ErrDeclined{link: nil, err: errors.New("REDACTED"), isScreened: true}})
	assert.NotPanics(t, func() {
		err := sw.Begin()
		require.NoError(t, err)
		err = sw.Halt()
		require.NoError(t, err)
	})
	//

	sw = NewRouter(cfg, faultCarrier{ErrCarrierHalted{}})
	assert.NotPanics(t, func() {
		err := sw.Begin()
		require.NoError(t, err)
		err = sw.Halt()
		require.NoError(t, err)
	})
}

//
//
type emulateHandler struct {
	*RootHandler

	//
	deleteNodeInAdvancement           uint32
	initInvokedPriorDeleteCompleted uint32
}

func (r *emulateHandler) DeleteNode(Node, any) {
	atomic.StoreUint32(&r.deleteNodeInAdvancement, 1)
	defer atomic.StoreUint32(&r.deleteNodeInAdvancement, 0)
	time.Sleep(100 * time.Millisecond)
}

func (r *emulateHandler) InitNode(node Node) Node {
	if atomic.LoadUint32(&r.deleteNodeInAdvancement) == 1 {
		atomic.StoreUint32(&r.initInvokedPriorDeleteCompleted, 1)
	}

	return node
}

func (r *emulateHandler) InitInvokedPriorDeleteCompleted() bool {
	return atomic.LoadUint32(&r.initInvokedPriorDeleteCompleted) == 1
}

//
func VerifyRouterInitNodeIsNegateInvokedPriorDeleteNode(t *testing.T) {
	//
	handler := &emulateHandler{}
	handler.RootHandler = NewRootHandler("REDACTED", handler)

	//
	sw := CreateRouter(cfg, 1, func(i int, sw *Router) *Router {
		sw.AppendHandler("REDACTED", handler)
		return sw
	})
	err := sw.Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateKey: ed25519.GeneratePrivateKey(), Settings: cfg}
	rp.Begin()
	defer rp.Halt()
	_, err = rp.Call(sw.NetLocation())
	require.NoError(t, err)

	//
	for {
		time.Sleep(20 * time.Millisecond)
		if node := sw.Nodes().Get(rp.ID()); node != nil {
			go sw.HaltNodeForFault(node, "REDACTED")
			break
		}
	}

	//
	_, err = rp.Call(sw.NetLocation())
	require.NoError(t, err)
	//
	time.Sleep(50 * time.Millisecond)

	//
	assert.False(t, handler.InitInvokedPriorDeleteCompleted())
}

func createRouterForCriterion(b *testing.B) *Router {
	b.Helper()
	s1, s2 := CreateRouterCouple(initRouterFunction)
	b.Cleanup(func() {
		if err := s2.Halt(); err != nil {
			b.Error(err)
		}
		if err := s1.Halt(); err != nil {
			b.Error(err)
		}
	})
	//
	time.Sleep(1 * time.Second)
	return s1
}

func CriterionRouterMulticast(b *testing.B) {
	sw := createRouterForCriterion(b)
	chanMessage := &p2pproto.PexLocations{
		Locations: []p2pproto.NetLocation{
			{
				ID: "REDACTED",
			},
		},
	}

	b.ResetTimer()

	//
	for i := 0; i < b.N; i++ {
		chanUID := byte(i % 4)
		sw.MulticastAsync(Packet{StreamUID: chanUID, Signal: chanMessage})
	}
}

func CriterionRouterAttemptMulticast(b *testing.B) {
	sw := createRouterForCriterion(b)
	chanMessage := &p2pproto.PexLocations{
		Locations: []p2pproto.NetLocation{
			{
				ID: "REDACTED",
			},
		},
	}

	b.ResetTimer()

	//
	for i := 0; i < b.N; i++ {
		chanUID := byte(i % 4)
		sw.AttemptMulticast(Packet{StreamUID: chanUID, Signal: chanMessage})
	}
}

func VerifyRouterDeletionErr(t *testing.T) {
	sw1, sw2 := CreateRouterCouple(initRouterFunction)

	require.Len(t, sw1.Nodes().Clone(), 1)
	p := sw1.Nodes().Clone()[0]

	sw2.HaltNodeForFault(p, fmt.Errorf("REDACTED"))

	assert.Equal(t, sw2.nodes.Add(p).Error(), ErrNodeDeletion{}.Fault())
}
