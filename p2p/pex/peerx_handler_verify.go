package pex

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/simulate"
	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
)

var cfg *settings.Peer2peerSettings

func initialize() {
	cfg = settings.FallbackPeer2peerSettings()
	cfg.PeerxHandler = true
	cfg.PermitReplicatedINET = true
}

func VerifyPeerxHandlerFundamental(t *testing.T) {
	r, register := generateHandler(&HandlerSettings{})
	defer deconfigureHandler(register)

	assert.NotNil(t, r)
	assert.NotEmpty(t, r.ObtainConduits())
}

func VerifyPeerxHandlerAppendDiscardNode(t *testing.T) {
	r, register := generateHandler(&HandlerSettings{})
	defer deconfigureHandler(register)

	extent := register.Extent()
	node := p2p.GenerateUnpredictableNode(false)

	r.AppendNode(node)
	assert.Equal(t, extent+1, register.Extent())

	r.DiscardNode(node, "REDACTED")

	outgoingNode := p2p.GenerateUnpredictableNode(true)

	r.AppendNode(outgoingNode)
	assert.Equal(t, extent+1, register.Extent(), "REDACTED")

	r.DiscardNode(outgoingNode, "REDACTED")
}

//
//
//
//
//
//
//
//
//
//
func VerifyPeerxHandlerActive(t *testing.T) {
	N := 3
	routers := make([]*p2p.Router, N)

	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	ledgers := make([]LocationRegister, N)
	tracer := log.VerifyingTracer()

	//
	for i := 0; i < N; i++ {
		routers[i] = p2p.CreateRouter(cfg, i, func(i int, sw *p2p.Router) *p2p.Router {
			ledgers[i] = FreshLocationRegister(filepath.Join(dir, fmt.Sprintf("REDACTED", i)), false)
			ledgers[i].AssignTracer(tracer.Using("REDACTED", i))
			sw.AssignLocationRegister(ledgers[i])

			sw.AssignTracer(tracer.Using("REDACTED", i))

			r := FreshHandler(ledgers[i], &HandlerSettings{})
			r.AssignTracer(tracer.Using("REDACTED", i))
			r.AssignAssureNodesSpan(250 * time.Millisecond)
			sw.AppendHandler("REDACTED", r)

			return sw
		})
	}

	appendAnotherPeerLocationTowardLocationRegister := func(routerPosition, anotherRouterPosition int) {
		location := routers[anotherRouterPosition].NetworkLocator()
		err := ledgers[routerPosition].AppendLocator(location, location)
		require.NoError(t, err)
	}

	appendAnotherPeerLocationTowardLocationRegister(0, 1)
	appendAnotherPeerLocationTowardLocationRegister(1, 0)
	appendAnotherPeerLocationTowardLocationRegister(2, 1)

	for _, sw := range routers {
		err := sw.Initiate() //
		require.Nil(t, err)
	}

	affirmNodesUsingDeadline(t, routers, 10*time.Millisecond, 10*time.Second, N-1)

	//
	for _, s := range routers {
		err := s.Halt()
		require.NoError(t, err)
	}
}

func VerifyPeerxHandlerAccept(t *testing.T) {
	r, register := generateHandler(&HandlerSettings{})
	defer deconfigureHandler(register)

	node := p2p.GenerateUnpredictableNode(false)

	//
	r.SolicitLocations(node)

	extent := register.Extent()
	msg := &tmpfabric.PeerxLocations{Locations: []tmpfabric.NetworkLocator{node.PortLocation().TowardSchema()}}
	r.Accept(p2p.Wrapper{ConduitUUID: PeerxConduit, Src: node, Signal: msg})
	assert.Equal(t, extent+1, register.Extent())

	r.Accept(p2p.Wrapper{ConduitUUID: PeerxConduit, Src: node, Signal: &tmpfabric.PeerxSolicit{}})
}

func VerifyPeerxHandlerSolicitArtifactMisuse(t *testing.T) {
	r, register := generateHandler(&HandlerSettings{})
	defer deconfigureHandler(register)

	sw := generateRouterAlsoAppendEngines(r)
	sw.AssignLocationRegister(register)

	node := simulate.FreshNode(nil)
	nodeLocation := node.PortLocation()
	p2p.AppendNodeTowardRouterNodeAssign(sw, node)
	assert.True(t, sw.Nodes().Has(node.ID()))
	err := register.AppendLocator(nodeLocation, nodeLocation)
	require.NoError(t, err)
	require.True(t, register.OwnsLocation(nodeLocation))

	id := string(node.ID())

	//
	r.Accept(p2p.Wrapper{ConduitUUID: PeerxConduit, Src: node, Signal: &tmpfabric.PeerxSolicit{}})
	assert.True(t, r.finalAcceptedSolicits.Has(id))
	assert.True(t, sw.Nodes().Has(node.ID()))

	//
	r.Accept(p2p.Wrapper{ConduitUUID: PeerxConduit, Src: node, Signal: &tmpfabric.PeerxSolicit{}})
	assert.True(t, r.finalAcceptedSolicits.Has(id))
	assert.True(t, sw.Nodes().Has(node.ID()))

	//
	r.Accept(p2p.Wrapper{ConduitUUID: PeerxConduit, Src: node, Signal: &tmpfabric.PeerxSolicit{}})
	assert.False(t, r.finalAcceptedSolicits.Has(id))
	assert.False(t, sw.Nodes().Has(node.ID()))
	assert.True(t, register.EqualsProhibited(nodeLocation))
}

func VerifyPeerxHandlerLocationsArtifactMisuse(t *testing.T) {
	r, register := generateHandler(&HandlerSettings{})
	defer deconfigureHandler(register)

	sw := generateRouterAlsoAppendEngines(r)
	sw.AssignLocationRegister(register)

	node := simulate.FreshNode(nil)
	p2p.AppendNodeTowardRouterNodeAssign(sw, node)
	assert.True(t, sw.Nodes().Has(node.ID()))

	id := string(node.ID())

	//
	r.SolicitLocations(node)
	assert.True(t, r.solicitsRelayed.Has(id))
	assert.True(t, sw.Nodes().Has(node.ID()))

	msg := &tmpfabric.PeerxLocations{Locations: []tmpfabric.NetworkLocator{node.PortLocation().TowardSchema()}}

	//
	r.Accept(p2p.Wrapper{ConduitUUID: PeerxConduit, Src: node, Signal: msg})
	assert.False(t, r.solicitsRelayed.Has(id))
	assert.True(t, sw.Nodes().Has(node.ID()))

	//
	r.Accept(p2p.Wrapper{ConduitUUID: PeerxConduit, Src: node, Signal: msg})
	assert.False(t, sw.Nodes().Has(node.ID()))
	assert.True(t, register.EqualsProhibited(node.PortLocation()))
}

func VerifyInspectOrigins(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	//
	nodeRouter := verifyGenerateFallbackNode(dir, 0)
	require.Nil(t, nodeRouter.Initiate())
	nodeRouter.Halt() //

	//
	germ := verifyGenerateGerm(dir, 1, []*p2p.NetworkLocator{}, []*p2p.NetworkLocator{})

	//
	nodeRouter = verifyGenerateNodeUsingGerm(dir, 2, germ)
	require.Nil(t, nodeRouter.Initiate())
	nodeRouter.Halt() //

	//
	flawedNodeSettings := &HandlerSettings{
		Origins: []string{
			"REDACTED",
			"REDACTED",
		},
	}
	nodeRouter = verifyGenerateNodeUsingSettings(dir, 2, flawedNodeSettings)
	require.Error(t, nodeRouter.Initiate())
	nodeRouter.Halt() //

	//
	flawedNodeSettings = &HandlerSettings{
		Origins: []string{
			"REDACTED",
			"REDACTED",
			germ.NetworkLocator().Text(),
		},
	}
	nodeRouter = verifyGenerateNodeUsingSettings(dir, 2, flawedNodeSettings)
	require.Nil(t, nodeRouter.Initiate())
	nodeRouter.Halt() //
}

func VerifyPeerxHandlerEmploysOriginsConditionalRequired(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	//
	germ := verifyGenerateGerm(dir, 0, []*p2p.NetworkLocator{}, []*p2p.NetworkLocator{})
	require.Nil(t, germ.Initiate())
	defer germ.Halt() //

	//
	node := verifyGenerateNodeUsingGerm(dir, 1, germ)
	require.Nil(t, node.Initiate())
	defer node.Halt() //

	//
	affirmNodesUsingDeadline(t, []*p2p.Router{node}, 10*time.Millisecond, 3*time.Second, 1)
}

func VerifyLinkageVelocityForeachNodeAcceptedOriginatingGerm(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	var id int
	var recognizedLocations []*p2p.NetworkLocator

	//
	for id = 0; id < 3; id++ {
		node := verifyGenerateFallbackNode(dir, id)
		require.NoError(t, node.Initiate())
		location := node.NetworkLocator()
		defer node.Halt() //

		recognizedLocations = append(recognizedLocations, location)
		t.Log("REDACTED", id, location)
	}

	//
	germ := verifyGenerateGerm(dir, id, recognizedLocations, recognizedLocations)
	require.NoError(t, germ.Initiate())
	defer germ.Halt() //
	t.Log("REDACTED", id, germ.NetworkLocator())

	//
	id++
	peer := verifyGenerateNodeUsingGerm(dir, id, germ)
	require.NoError(t, peer.Initiate())
	defer peer.Halt() //
	t.Log("REDACTED", id, peer.NetworkLocator())

	//
	affirmNodesUsingDeadline(t, []*p2p.Router{peer}, 10*time.Millisecond, 3*time.Second, 1)

	//
	affirmNodesUsingDeadline(t, []*p2p.Router{peer}, 10*time.Millisecond, 1*time.Second, 2)

	//
	//
	outgoing, incoming, calling := peer.CountNodes()
	assert.LessOrEqual(t, incoming, cfg.MaximumCountIncomingNodes)
	assert.LessOrEqual(t, outgoing, cfg.MaximumCountOutgoingNodes)
	assert.LessOrEqual(t, calling, cfg.MaximumCountOutgoingNodes+cfg.MaximumCountIncomingNodes-outgoing-incoming)
}

func VerifyPeerxHandlerGermStyle(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	peerxReaderSettings := &HandlerSettings{OriginStyle: true, GermDetachPauseSpan: 10 * time.Millisecond}
	peerxReader, register := generateHandler(peerxReaderSettings)
	defer deconfigureHandler(register)

	sw := generateRouterAlsoAppendEngines(peerxReader)
	sw.AssignLocationRegister(register)
	err = sw.Initiate()
	require.NoError(t, err)
	defer sw.Halt() //

	assert.Zero(t, sw.Nodes().Extent())

	nodeRouter := verifyGenerateFallbackNode(dir, 1)
	require.NoError(t, nodeRouter.Initiate())
	defer nodeRouter.Halt() //

	//
	peerxReader.exploreNodes([]*p2p.NetworkLocator{nodeRouter.NetworkLocator()})
	assert.Equal(t, 1, sw.Nodes().Extent())
	assert.True(t, sw.Nodes().Has(nodeRouter.PeerDetails().ID()))

	//
	peerxReader.effortUnlinks()
	assert.Equal(t, 1, sw.Nodes().Extent())

	//
	time.Sleep(peerxReaderSettings.GermDetachPauseSpan + 1*time.Millisecond)

	//
	peerxReader.effortUnlinks()
	assert.Equal(t, 0, sw.Nodes().Extent())
}

func VerifyPeerxHandlerExecutesNegationDetachOriginatingEnduringNodeInsideGermStyle(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	peerxReaderSettings := &HandlerSettings{OriginStyle: true, GermDetachPauseSpan: 1 * time.Millisecond}
	peerxReader, register := generateHandler(peerxReaderSettings)
	defer deconfigureHandler(register)

	sw := generateRouterAlsoAppendEngines(peerxReader)
	sw.AssignLocationRegister(register)
	err = sw.Initiate()
	require.NoError(t, err)
	defer sw.Halt() //

	assert.Zero(t, sw.Nodes().Extent())

	nodeRouter := verifyGenerateFallbackNode(dir, 1)
	require.NoError(t, nodeRouter.Initiate())
	defer nodeRouter.Halt() //

	err = sw.AppendEnduringNodes([]string{nodeRouter.NetworkLocator().Text()})
	require.NoError(t, err)

	//
	peerxReader.exploreNodes([]*p2p.NetworkLocator{nodeRouter.NetworkLocator()})
	assert.Equal(t, 1, sw.Nodes().Extent())
	assert.True(t, sw.Nodes().Has(nodeRouter.PeerDetails().ID()))

	//
	time.Sleep(peerxReaderSettings.GermDetachPauseSpan + 1*time.Millisecond)

	//
	peerxReader.effortUnlinks()
	assert.Equal(t, 1, sw.Nodes().Extent())
}

func VerifyPeerxHandlerCallsNodeAscendTowardMaximumEndeavorsInsideGermStyle(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	peerxReader, register := generateHandler(&HandlerSettings{OriginStyle: true})
	defer deconfigureHandler(register)

	sw := generateRouterAlsoAppendEngines(peerxReader)
	sw.AssignLocationRegister(register)
	//

	node := simulate.FreshNode(nil)
	location := node.PortLocation()

	err = register.AppendLocator(location, location)
	require.NoError(t, err)

	assert.True(t, register.OwnsLocation(location))

	//
	peerxReader.endeavorsTowardCall.Store(location.CallText(), _attemptscustody{maximumEndeavorsTowardCall + 1, time.Now()})
	peerxReader.exploreNodes([]*p2p.NetworkLocator{location})

	assert.False(t, register.OwnsLocation(location))
}

//
//
//
//
//
func VerifyPeerxHandlerGermStylePurgeHalt(t *testing.T) {
	N := 2
	routers := make([]*p2p.Router, N)

	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	ledgers := make([]LocationRegister, N)
	tracer := log.VerifyingTracer()

	//
	for i := 0; i < N; i++ {
		routers[i] = p2p.CreateRouter(cfg, i, func(i int, sw *p2p.Router) *p2p.Router {
			ledgers[i] = FreshLocationRegister(filepath.Join(dir, fmt.Sprintf("REDACTED", i)), false)
			ledgers[i].AssignTracer(tracer.Using("REDACTED", i))
			sw.AssignLocationRegister(ledgers[i])

			sw.AssignTracer(tracer.Using("REDACTED", i))

			settings := &HandlerSettings{}
			if i == 0 {
				//
				settings = &HandlerSettings{OriginStyle: true}
			}
			r := FreshHandler(ledgers[i], settings)
			r.AssignTracer(tracer.Using("REDACTED", i))
			r.AssignAssureNodesSpan(250 * time.Millisecond)
			sw.AppendHandler("REDACTED", r)

			return sw
		})
	}

	for _, sw := range routers {
		err := sw.Initiate() //
		require.Nil(t, err)
	}

	handler := routers[0].Engines()["REDACTED"].(*Handler)
	nodeUUID := routers[1].PeerDetails().ID()

	err = routers[1].CallNodeUsingLocator(routers[0].NetworkLocator())
	assert.NoError(t, err)

	//
	//
	//
	for i := 0; i < 1000; i++ {
		v := handler.finalAcceptedSolicits.Get(string(nodeUUID))
		if v != nil {
			break
		}
		time.Sleep(time.Millisecond)
	}

	//
	//
	nodes := routers[0].Nodes().Duplicate()
	for _, node := range nodes {
		err := node.Halt()
		require.NoError(t, err)
	}

	//
	for _, s := range routers {
		err := s.Halt()
		require.NoError(t, err)
	}
}

func VerifyPeerxHandlerExecutesNegationAppendSecludedNodesTowardLocationRegister(t *testing.T) {
	node := p2p.GenerateUnpredictableNode(false)

	peerxReader, register := generateHandler(&HandlerSettings{})
	register.AppendSecludedIDXDstore([]string{string(node.PeerDetails().ID())})
	defer deconfigureHandler(register)

	//
	peerxReader.SolicitLocations(node)

	extent := register.Extent()
	msg := &tmpfabric.PeerxLocations{Locations: []tmpfabric.NetworkLocator{node.PortLocation().TowardSchema()}}
	peerxReader.Accept(p2p.Wrapper{
		ConduitUUID: PeerxConduit,
		Src:       node,
		Signal:   msg,
	})
	assert.Equal(t, extent, register.Extent())

	peerxReader.AppendNode(node)
	assert.Equal(t, extent, register.Extent())
}

func VerifyPeerxHandlerCallNode(t *testing.T) {
	peerxReader, register := generateHandler(&HandlerSettings{})
	defer deconfigureHandler(register)

	sw := generateRouterAlsoAppendEngines(peerxReader)
	sw.AssignLocationRegister(register)

	node := simulate.FreshNode(nil)
	location := node.PortLocation()

	assert.Equal(t, 0, peerxReader.EndeavorsTowardCall(location))

	//
	err := peerxReader.callNode(location)
	require.Error(t, err)

	assert.Equal(t, 1, peerxReader.EndeavorsTowardCall(location))

	//
	err = peerxReader.callNode(location)
	require.Error(t, err)

	//
	assert.Equal(t, 1, peerxReader.EndeavorsTowardCall(location))

	if !testing.Short() {
		time.Sleep(3 * time.Second)

		//
		err = peerxReader.callNode(location)
		require.Error(t, err)

		assert.Equal(t, 2, peerxReader.EndeavorsTowardCall(location))
	}
}

func affirmNodesUsingDeadline(
	t *testing.T,
	routers []*p2p.Router,
	inspectSpan, deadline time.Duration,
	nthNodes int,
) {
	var (
		metronome    = time.NewTicker(inspectSpan)
		deadlineChnl = time.After(deadline)
	)
	defer metronome.Stop()

	for {
		select {
		case <-metronome.C:
			//
			everyValid := true
			for _, s := range routers {
				outgoing, incoming, _ := s.CountNodes()
				if outgoing+incoming < nthNodes {
					everyValid = false
					break
				}
			}
			if everyValid {
				return
			}
		case <-deadlineChnl:
			countNodesTxt := "REDACTED"
			for i, s := range routers {
				outgoing, incoming, _ := s.CountNodes()
				countNodesTxt += fmt.Sprintf("REDACTED", i, outgoing, incoming)
			}
			t.Errorf(
				"REDACTED",
				nthNodes, countNodesTxt,
			)
			return
		}
	}
}

//
func verifyGenerateNodeUsingSettings(dir string, id int, settings *HandlerSettings) *p2p.Router {
	node := p2p.CreateRouter(
		cfg,
		id,
		func(i int, sw *p2p.Router) *p2p.Router {
			register := FreshLocationRegister(filepath.Join(dir, fmt.Sprintf("REDACTED", id)), false)
			register.AssignTracer(log.VerifyingTracer().Using("REDACTED", id))
			sw.AssignLocationRegister(register)

			r := FreshHandler(
				register,
				settings,
			)
			r.AssignTracer(log.VerifyingTracer().Using("REDACTED", id))
			sw.AppendHandler("REDACTED", r)
			return sw
		},
	)
	return node
}

//
func verifyGenerateFallbackNode(dir string, id int) *p2p.Router {
	return verifyGenerateNodeUsingSettings(dir, id, &HandlerSettings{})
}

//
//
func verifyGenerateGerm(dir string, id int, recognizedLocations, originLocations []*p2p.NetworkLocator) *p2p.Router {
	germ := p2p.CreateRouter(
		cfg,
		id,
		func(i int, sw *p2p.Router) *p2p.Router {
			register := FreshLocationRegister(filepath.Join(dir, "REDACTED"), false)
			register.AssignTracer(log.VerifyingTracer())
			for j := 0; j < len(recognizedLocations); j++ {
				register.AppendLocator(recognizedLocations[j], originLocations[j]) //
				register.LabelValid(recognizedLocations[j].ID)
			}
			sw.AssignLocationRegister(register)

			sw.AssignTracer(log.VerifyingTracer())

			r := FreshHandler(register, &HandlerSettings{})
			r.AssignTracer(log.VerifyingTracer())
			sw.AppendHandler("REDACTED", r)
			return sw
		},
	)
	return germ
}

//
//
func verifyGenerateNodeUsingGerm(dir string, id int, germ *p2p.Router) *p2p.Router {
	setting := &HandlerSettings{
		Origins: []string{germ.NetworkLocator().Text()},
	}
	return verifyGenerateNodeUsingSettings(dir, id, setting)
}

func generateHandler(setting *HandlerSettings) (r *Handler, register LocationRegister) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}
	register = FreshLocationRegister(filepath.Join(dir, "REDACTED"), true)
	register.AssignTracer(log.VerifyingTracer())

	r = FreshHandler(register, setting)
	r.AssignTracer(log.VerifyingTracer())
	return
}

func deconfigureHandler(register LocationRegister) {
	//
	err := os.RemoveAll(filepath.Dir(register.(*locationRegister).RecordRoute()))
	if err != nil {
		panic(err)
	}
}

func generateRouterAlsoAppendEngines(engines ...p2p.Handler) *p2p.Router {
	sw := p2p.CreateRouter(cfg, 0, func(i int, sw *p2p.Router) *p2p.Router { return sw })
	sw.AssignTracer(log.VerifyingTracer())
	for _, r := range engines {
		sw.AppendHandler(r.Text(), r)
		r.AssignRouter(sw)
	}
	return sw
}

func VerifyPeerxArrays(t *testing.T) {
	location := tmpfabric.NetworkLocator{
		ID:   "REDACTED",
		IP:   "REDACTED",
		Channel: 9090,
	}

	verifyScenarios := []struct {
		verifyAlias string
		msg      proto.Message
		expirationOctets string
	}{
		{"REDACTED", &tmpfabric.PeerxSolicit{}, "REDACTED"},
		{"REDACTED", &tmpfabric.PeerxLocations{Locations: []tmpfabric.NetworkLocator{location}}, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		w := tc.msg.(p2p.Encapsulator).Enclose()
		bz, err := proto.Marshal(w)
		require.NoError(t, err)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyAlias)
	}
}
