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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
	fabricscheme "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
)

var cfg *settings.Peer2peerSettings

func initialize() {
	cfg = settings.FallbackPeer2peerSettings()
	cfg.PeerxHandler = true
	cfg.PermitReplicatedINET = true
}

type NodeArtifact struct {
	Material proto.Message
	Tally  int
}

type VerifyHandler struct {
	FoundationHandler

	mtx          commitchronize.Exclusion
	conduits     []*link.ConduitDefinition
	reportSignals  bool
	artifactsTally  int
	artifactsAccepted map[byte][]NodeArtifact
}

func FreshVerifyHandler(conduits []*link.ConduitDefinition, reportSignals bool) *VerifyHandler {
	tr := &VerifyHandler{
		conduits:     conduits,
		reportSignals:  reportSignals,
		artifactsAccepted: make(map[byte][]NodeArtifact),
	}
	tr.FoundationHandler = *FreshFoundationHandler("REDACTED", tr)
	tr.AssignTracer(log.VerifyingTracer())
	return tr
}

func (tr *VerifyHandler) ObtainConduits() []*link.ConduitDefinition {
	return tr.conduits
}

func (tr *VerifyHandler) AppendNode(Node) {}

func (tr *VerifyHandler) DiscardNode(Node, any) {}

func (tr *VerifyHandler) Accept(e Wrapper) {
	if tr.reportSignals {
		tr.mtx.Lock()
		defer tr.mtx.Unlock()
		fmt.Printf("REDACTED", e.ConduitUUID, e.Signal)
		tr.artifactsAccepted[e.ConduitUUID] = append(tr.artifactsAccepted[e.ConduitUUID], NodeArtifact{Material: e.Signal, Tally: tr.artifactsTally})
		tr.artifactsTally++
	}
}

func (tr *VerifyHandler) obtainArtifacts(chnlUUID byte) []NodeArtifact {
	tr.mtx.Lock()
	defer tr.mtx.Unlock()
	return tr.artifactsAccepted[chnlUUID]
}

//

//
//
func CreateRouterDuo(initializeRouter func(int, *Router) *Router) (*Router, *Router) {
	//
	routers := CreateAssociatedRouters(cfg, 2, initializeRouter, Connect2routers)
	return routers[0], routers[1]
}

func initializeRouterMethod(_ int, sw *Router) *Router {
	sw.AssignLocationRegister(&LocationRegisterSimulate{
		Locations:    make(map[string]struct{}),
		MineLocations: make(map[string]struct{}),
	})

	//
	sw.AppendHandler("REDACTED", FreshVerifyHandler([]*link.ConduitDefinition{
		{ID: byte(0x00), Urgency: 10, SignalKind: &fabricscheme.Signal{}},
		{ID: byte(0x01), Urgency: 10, SignalKind: &fabricscheme.Signal{}},
	}, true))
	sw.AppendHandler("REDACTED", FreshVerifyHandler([]*link.ConduitDefinition{
		{ID: byte(0x02), Urgency: 10, SignalKind: &fabricscheme.Signal{}},
		{ID: byte(0x03), Urgency: 10, SignalKind: &fabricscheme.Signal{}},
	}, true))

	return sw
}

func VerifyRouters(t *testing.T) {
	s1, s2 := CreateRouterDuo(initializeRouterMethod)
	t.Cleanup(func() {
		if err := s2.Halt(); err != nil {
			t.Error(err)
		}
		if err := s1.Halt(); err != nil {
			t.Error(err)
		}
	})

	if s1.Nodes().Extent() != 1 {
		t.Errorf("REDACTED", s1.Nodes().Extent())
	}
	if s2.Nodes().Extent() != 1 {
		t.Errorf("REDACTED", s2.Nodes().Extent())
	}

	//
	chnl0artifact := &fabricscheme.PeerxLocations{
		Locations: []fabricscheme.NetworkLocator{
			{
				ID: "REDACTED",
			},
		},
	}
	chnl1artifact := &fabricscheme.PeerxLocations{
		Locations: []fabricscheme.NetworkLocator{
			{
				ID: "REDACTED",
			},
		},
	}
	chnl2artifact := &fabricscheme.PeerxLocations{
		Locations: []fabricscheme.NetworkLocator{
			{
				ID: "REDACTED",
			},
		},
	}
	//
	//
	s1.MulticastAsyncronous(Wrapper{ConduitUUID: byte(0x00), Signal: chnl0artifact})
	s1.MulticastAsyncronous(Wrapper{ConduitUUID: byte(0x01), Signal: chnl1artifact})
	s1.AttemptMulticast(Wrapper{ConduitUUID: byte(0x02), Signal: chnl2artifact})

	obtainHandler := func(alias string) *VerifyHandler {
		r, ok := s2.Handler(alias)
		require.True(t, ok)

		tr, ok := r.(*VerifyHandler)
		require.True(t, ok)

		return tr
	}

	handlerSample := obtainHandler("REDACTED")
	handlerDivider := obtainHandler("REDACTED")

	affirmSignalAcceptedUsingDeadline(t, chnl0artifact, byte(0x00), handlerSample, 200*time.Millisecond, 5*time.Second)
	affirmSignalAcceptedUsingDeadline(t, chnl1artifact, byte(0x01), handlerSample, 200*time.Millisecond, 5*time.Second)
	affirmSignalAcceptedUsingDeadline(t, chnl2artifact, byte(0x02), handlerDivider, 200*time.Millisecond, 5*time.Second)
}

func affirmSignalAcceptedUsingDeadline(
	t *testing.T,
	msg proto.Message,
	conduit byte,
	handler *VerifyHandler,
	inspectSpan,
	deadline time.Duration,
) {
	metronome := time.NewTicker(inspectSpan)
	defer metronome.Stop()

	for {
		select {
		case <-metronome.C:
			signals := handler.obtainArtifacts(conduit)
			if len(signals) == 0 {
				t.Fatalf("REDACTED", conduit, len(signals))
			}

			anticipatedOctets, err := proto.Marshal(signals[0].Material)
			require.NoError(t, err)
			attainedOctets, err := proto.Marshal(msg)
			require.NoError(t, err)
			if len(signals) > 0 {
				if !bytes.Equal(anticipatedOctets, attainedOctets) {
					t.Fatalf("REDACTED", msg, signals[0].Tally)
				}
				return
			}

		case <-time.After(deadline):
			t.Fatalf("REDACTED", conduit)
		}
	}
}

func VerifyRouterCriteriaOutputSelf(t *testing.T) {
	s1 := CreateRouter(cfg, 1, initializeRouterMethod)

	//
	rp := &distantNode{PrivateToken: s1.peerToken.PrivateToken, Settings: cfg}
	rp.Initiate()

	//
	err := s1.CallNodeUsingLocator(rp.Location())
	if assert.Error(t, err) {
		if err, ok := err.(FaultDeclined); ok {
			if !err.EqualsEgo() {
				t.Errorf("REDACTED")
			}
		} else {
			t.Errorf("REDACTED")
		}
	}

	assert.True(t, s1.locationRegister.MineLocator(rp.Location()))
	assert.False(t, s1.locationRegister.OwnsLocation(rp.Location()))

	rp.Halt()

	affirmNegativeNodesSubsequentDeadline(t, s1, 100*time.Millisecond)
}

func VerifyRouterNodeRefine(t *testing.T) {
	var (
		criteria = []NodeRefineMethod{
			func(_ IDXNodeAssign, _ Node) error { return nil },
			func(_ IDXNodeAssign, _ Node) error { return fmt.Errorf("REDACTED") },
			func(_ IDXNodeAssign, _ Node) error { return nil },
		}
		sw = CreateRouter(
			cfg,
			1,
			initializeRouterMethod,
			RouterNodeCriteria(criteria...),
		)
	)
	err := sw.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
	rp.Initiate()
	t.Cleanup(rp.Halt)

	p, err := sw.carrier.Call(*rp.Location(), nodeSettings{
		chnlDescriptions:      sw.chnlDescriptions,
		uponNodeFailure:  sw.HaltNodeForeachFailure,
		equalsEnduring: sw.EqualsNodeEnduring,
		enginesViaChnl: sw.enginesViaChnl,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = sw.appendNode(p)
	if err, ok := err.(FaultDeclined); ok {
		if !err.EqualsScreened() {
			t.Errorf("REDACTED")
		}
	} else {
		t.Errorf("REDACTED")
	}
}

func VerifyRouterNodeRefineDeadline(t *testing.T) {
	var (
		criteria = []NodeRefineMethod{
			func(_ IDXNodeAssign, _ Node) error {
				time.Sleep(10 * time.Millisecond)
				return nil
			},
		}
		sw = CreateRouter(
			cfg,
			1,
			initializeRouterMethod,
			RouterRefineDeadline(5*time.Millisecond),
			RouterNodeCriteria(criteria...),
		)
	)
	err := sw.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Log(err)
		}
	})

	//
	rp := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
	rp.Initiate()
	defer rp.Halt()

	p, err := sw.carrier.Call(*rp.Location(), nodeSettings{
		chnlDescriptions:      sw.chnlDescriptions,
		uponNodeFailure:  sw.HaltNodeForeachFailure,
		equalsEnduring: sw.EqualsNodeEnduring,
		enginesViaChnl: sw.enginesViaChnl,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = sw.appendNode(p)
	if _, ok := err.(FaultRefineDeadline); !ok {
		t.Errorf("REDACTED")
	}
}

func VerifyRouterNodeRefineReplicated(t *testing.T) {
	sw := CreateRouter(cfg, 1, initializeRouterMethod)
	err := sw.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
	rp.Initiate()
	defer rp.Halt()

	p, err := sw.carrier.Call(*rp.Location(), nodeSettings{
		chnlDescriptions:      sw.chnlDescriptions,
		uponNodeFailure:  sw.HaltNodeForeachFailure,
		equalsEnduring: sw.EqualsNodeEnduring,
		enginesViaChnl: sw.enginesViaChnl,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := sw.appendNode(p); err != nil {
		t.Fatal(err)
	}

	err = sw.appendNode(p)
	if faultNack, ok := err.(FaultDeclined); ok {
		if !faultNack.EqualsReplicated() {
			t.Errorf("REDACTED", faultNack)
		}
	} else {
		t.Errorf("REDACTED", err)
	}
}

func affirmNegativeNodesSubsequentDeadline(t *testing.T, sw *Router, deadline time.Duration) {
	time.Sleep(deadline)
	if sw.Nodes().Extent() != 0 {
		t.Fatalf("REDACTED", sw, sw.Nodes().Extent())
	}
}

func VerifyRouterHaltsUnEnduringNodeUponFailure(t *testing.T) {
	affirm, demand := assert.New(t), require.New(t)

	sw := CreateRouter(cfg, 1, initializeRouterMethod)
	err := sw.Initiate()
	if err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
	rp.Initiate()
	defer rp.Halt()

	p, err := sw.carrier.Call(*rp.Location(), nodeSettings{
		chnlDescriptions:      sw.chnlDescriptions,
		uponNodeFailure:  sw.HaltNodeForeachFailure,
		equalsEnduring: sw.EqualsNodeEnduring,
		enginesViaChnl: sw.enginesViaChnl,
	})
	require.Nil(err)

	err = sw.appendNode(p)
	require.Nil(err)

	require.NotNil(sw.Nodes().Get(rp.ID()))

	//
	err = p.(*node).ShutdownLink()
	require.NoError(err)

	affirmNegativeNodesSubsequentDeadline(t, sw, 100*time.Millisecond)
	assert.False(p.EqualsActive())
}

func VerifyRouterHaltNodeForeachFailure(t *testing.T) {
	s := httptest.NewServer(promhttp.Handler())
	defer s.Close()

	extractTelemetry := func() string {
		reply, err := http.Get(s.URL)
		require.NoError(t, err)
		defer reply.Body.Close()
		buf, _ := io.ReadAll(reply.Body)
		return string(buf)
	}

	scope, component, alias := settings.VerifyTelemetrySettings().Scope, TelemetryComponent, "REDACTED"
	re := regexp.MustCompile(scope + "REDACTED" + component + "REDACTED" + alias + "REDACTED")
	nodesIndicatorDatum := func() float64 {
		aligns := re.FindStringSubmatch(extractTelemetry())
		f, _ := strconv.ParseFloat(aligns[1], 64)
		return f
	}

	peer2peerTelemetry := TitanTelemetry(scope)

	//
	sw1, sw2 := CreateRouterDuo(func(i int, sw *Router) *Router {
		//
		if i == 0 {
			opt := UsingTelemetry(peer2peerTelemetry)
			opt(sw)
		}
		return initializeRouterMethod(i, sw)
	})

	assert.Len(t, sw1.Nodes().Duplicate(), 1)
	assert.EqualValues(t, 1, nodesIndicatorDatum())

	//
	p := sw1.Nodes().Duplicate()[0]
	p.Transmit(Wrapper{
		ConduitUUID: 0x1,
		Signal:   &fabricscheme.Signal{},
	})

	//
	//
	t.Cleanup(func() {
		if err := sw2.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	sw1.HaltNodeForeachFailure(p, fmt.Errorf("REDACTED"))

	require.Empty(t, len(sw1.Nodes().Duplicate()), 0)
	assert.EqualValues(t, 0, nodesIndicatorDatum())
}

func VerifyRouterReestablishesTowardOutgoingEnduringNode(t *testing.T) {
	sw := CreateRouter(cfg, 1, initializeRouterMethod)
	err := sw.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
	rp.Initiate()
	defer rp.Halt()

	err = sw.AppendEnduringNodes([]string{rp.Location().Text()})
	require.NoError(t, err)

	err = sw.CallNodeUsingLocator(rp.Location())
	require.Nil(t, err)
	require.NotNil(t, sw.Nodes().Get(rp.ID()))

	p := sw.Nodes().Duplicate()[0]
	err = p.(*node).ShutdownLink()
	require.NoError(t, err)

	pauseTillRouterOwnsLocatedMinimumNTHNodes(sw, 1)
	assert.False(t, p.EqualsActive())        //
	assert.Equal(t, 1, sw.Nodes().Extent()) //

	//
	rp = &distantNode{
		PrivateToken: edwards25519.ProducePrivateToken(),
		Settings:  cfg,
		//
		//
		overhearLocation: "REDACTED",
	}
	rp.Initiate()
	defer rp.Halt()

	setting := settings.FallbackPeer2peerSettings()
	setting.VerifyCallMishap = true //
	err = sw.appendOutgoingNodeUsingSettings(rp.Location(), setting)
	require.NotNil(t, err)
	//
	pauseTillRouterOwnsLocatedMinimumNTHNodes(sw, 2)
	assert.Equal(t, 2, sw.Nodes().Extent())
}

func VerifyRouterReestablishesTowardIncomingEnduringNode(t *testing.T) {
	sw := CreateRouter(cfg, 1, initializeRouterMethod)
	err := sw.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
	rp.Initiate()
	defer rp.Halt()

	err = sw.AppendEnduringNodes([]string{rp.Location().Text()})
	require.NoError(t, err)

	link, err := rp.Call(sw.NetworkLocator())
	require.NoError(t, err)
	time.Sleep(50 * time.Millisecond)
	require.NotNil(t, sw.Nodes().Get(rp.ID()))

	link.Close()

	pauseTillRouterOwnsLocatedMinimumNTHNodes(sw, 1)
	assert.Equal(t, 1, sw.Nodes().Extent())
}

func VerifyRouterCallNodesAsyncronous(t *testing.T) {
	if testing.Short() {
		return
	}

	sw := CreateRouter(cfg, 1, initializeRouterMethod)
	err := sw.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	rp := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
	rp.Initiate()
	defer rp.Halt()

	err = sw.CallNodesAsyncronous([]string{rp.Location().Text()})
	require.NoError(t, err)
	time.Sleep(callGeneratorDurationMillis * time.Millisecond)
	require.NotNil(t, sw.Nodes().Get(rp.ID()))
}

func pauseTillRouterOwnsLocatedMinimumNTHNodes(sw *Router, n int) {
	for i := 0; i < 20; i++ {
		time.Sleep(250 * time.Millisecond)
		has := sw.Nodes().Extent()
		if has >= n {
			break
		}
	}
}

func VerifyRouterCompleteAccessibility(t *testing.T) {
	routers := CreateAssociatedRouters(cfg, 3, initializeRouterMethod, Connect2routers)
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
		if sw.Nodes().Extent() != 2 {
			t.Fatalf("REDACTED", sw.Nodes().Extent(), i)
		}
	}
}

func VerifyRouterEmbraceProcedure(t *testing.T) {
	cfg.MaximumCountIncomingNodes = 5

	//
	const absoluteNodesCount = 2
	var (
		absoluteNodes   = make([]*distantNode, absoluteNodesCount)
		absoluteNodeIDXDstore = make([]string, absoluteNodesCount)
	)
	for i := 0; i < absoluteNodesCount; i++ {
		node := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
		node.Initiate()
		absoluteNodes[i] = node
		absoluteNodeIDXDstore[i] = string(node.ID())
	}

	//
	sw := CreateRouter(cfg, 1, initializeRouterMethod)
	err := sw.AppendAbsoluteNodeIDXDstore(absoluteNodeIDXDstore)
	require.NoError(t, err)
	err = sw.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		err := sw.Halt()
		require.NoError(t, err)
	})

	//
	assert.Equal(t, 0, sw.Nodes().Extent())

	//
	nodes := make([]*distantNode, 0)
	for i := 0; i < cfg.MaximumCountIncomingNodes; i++ {
		node := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
		nodes = append(nodes, node)
		node.Initiate()
		c, err := node.Call(sw.NetworkLocator())
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
	assert.Equal(t, cfg.MaximumCountIncomingNodes, sw.Nodes().Extent())

	//
	node := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
	node.Initiate()
	link, err := node.Call(sw.NetworkLocator())
	require.NoError(t, err)
	//
	one := make([]byte, 1)
	_ = link.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
	_, err = link.Read(one)
	assert.Error(t, err)
	assert.Equal(t, cfg.MaximumCountIncomingNodes, sw.Nodes().Extent())
	node.Halt()

	//
	for _, node := range absoluteNodes {
		c, err := node.Call(sw.NetworkLocator())
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
	assert.Equal(t, cfg.MaximumCountIncomingNodes+absoluteNodesCount, sw.Nodes().Extent())

	for _, node := range nodes {
		node.Halt()
	}
	for _, node := range absoluteNodes {
		node.Halt()
	}
}

type failureCarrier struct {
	embraceFault error
}

func (et failureCarrier) NetworkLocator() NetworkLocator {
	panic("REDACTED")
}

func (et failureCarrier) Embrace(nodeSettings) (Node, error) {
	return nil, et.embraceFault
}

func (failureCarrier) Call(NetworkLocator, nodeSettings) (Node, error) {
	panic("REDACTED")
}

func (failureCarrier) Sanitize(Node) {
	panic("REDACTED")
}

func VerifyRouterEmbraceProcedureFailureScenarios(t *testing.T) {
	sw := FreshRouter(cfg, failureCarrier{FaultRefineDeadline{}})
	assert.NotPanics(t, func() {
		err := sw.Initiate()
		require.NoError(t, err)
		err = sw.Halt()
		require.NoError(t, err)
	})

	sw = FreshRouter(cfg, failureCarrier{FaultDeclined{link: nil, err: errors.New("REDACTED"), equalsScreened: true}})
	assert.NotPanics(t, func() {
		err := sw.Initiate()
		require.NoError(t, err)
		err = sw.Halt()
		require.NoError(t, err)
	})
	//

	sw = FreshRouter(cfg, failureCarrier{FaultCarrierTerminated{}})
	assert.NotPanics(t, func() {
		err := sw.Initiate()
		require.NoError(t, err)
		err = sw.Halt()
		require.NoError(t, err)
	})
}

//
//
type simulateHandler struct {
	*FoundationHandler

	//
	discardNodeInsideOnward           uint32
	initializeInvokedPriorDiscardConcluded uint32
}

func (r *simulateHandler) DiscardNode(Node, any) {
	atomic.StoreUint32(&r.discardNodeInsideOnward, 1)
	defer atomic.StoreUint32(&r.discardNodeInsideOnward, 0)
	time.Sleep(100 * time.Millisecond)
}

func (r *simulateHandler) InitializeNode(node Node) Node {
	if atomic.LoadUint32(&r.discardNodeInsideOnward) == 1 {
		atomic.StoreUint32(&r.initializeInvokedPriorDiscardConcluded, 1)
	}

	return node
}

func (r *simulateHandler) InitializeInvokedPriorDiscardConcluded() bool {
	return atomic.LoadUint32(&r.initializeInvokedPriorDiscardConcluded) == 1
}

//
func VerifyRouterInitializeNodeEqualsNegationInvokedPriorDiscardNode(t *testing.T) {
	//
	handler := &simulateHandler{}
	handler.FoundationHandler = FreshFoundationHandler("REDACTED", handler)

	//
	sw := CreateRouter(cfg, 1, func(i int, sw *Router) *Router {
		sw.AppendHandler("REDACTED", handler)
		return sw
	})
	err := sw.Initiate()
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := sw.Halt(); err != nil {
			t.Error(err)
		}
	})

	//
	rp := &distantNode{PrivateToken: edwards25519.ProducePrivateToken(), Settings: cfg}
	rp.Initiate()
	defer rp.Halt()
	_, err = rp.Call(sw.NetworkLocator())
	require.NoError(t, err)

	//
	for {
		time.Sleep(20 * time.Millisecond)
		if node := sw.Nodes().Get(rp.ID()); node != nil {
			go sw.HaltNodeForeachFailure(node, "REDACTED")
			break
		}
	}

	//
	_, err = rp.Call(sw.NetworkLocator())
	require.NoError(t, err)
	//
	time.Sleep(50 * time.Millisecond)

	//
	assert.False(t, handler.InitializeInvokedPriorDiscardConcluded())
}

func createRouterForeachAssessment(b *testing.B) *Router {
	b.Helper()
	s1, s2 := CreateRouterDuo(initializeRouterMethod)
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

func AssessmentRouterMulticast(b *testing.B) {
	sw := createRouterForeachAssessment(b)
	chnlSignal := &fabricscheme.PeerxLocations{
		Locations: []fabricscheme.NetworkLocator{
			{
				ID: "REDACTED",
			},
		},
	}

	b.ResetTimer()

	//
	for i := 0; i < b.N; i++ {
		chnlUUID := byte(i % 4)
		sw.MulticastAsyncronous(Wrapper{ConduitUUID: chnlUUID, Signal: chnlSignal})
	}
}

func AssessmentRouterAttemptMulticast(b *testing.B) {
	sw := createRouterForeachAssessment(b)
	chnlSignal := &fabricscheme.PeerxLocations{
		Locations: []fabricscheme.NetworkLocator{
			{
				ID: "REDACTED",
			},
		},
	}

	b.ResetTimer()

	//
	for i := 0; i < b.N; i++ {
		chnlUUID := byte(i % 4)
		sw.AttemptMulticast(Wrapper{ConduitUUID: chnlUUID, Signal: chnlSignal})
	}
}

func VerifyRouterDeletionFault(t *testing.T) {
	sw1, sw2 := CreateRouterDuo(initializeRouterMethod)

	require.Len(t, sw1.Nodes().Duplicate(), 1)
	p := sw1.Nodes().Duplicate()[0]

	sw2.HaltNodeForeachFailure(p, fmt.Errorf("REDACTED"))

	assert.Equal(t, sw2.nodes.Add(p).Error(), FaultNodeDeletion{}.Failure())
}
