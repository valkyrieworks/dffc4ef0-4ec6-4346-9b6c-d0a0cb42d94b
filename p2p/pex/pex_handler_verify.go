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

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/p2p/emulate"
	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
)

var cfg *settings.P2PSettings

func init() {
	cfg = settings.StandardP2PSettings()
	cfg.PexHandler = true
	cfg.PermitReplicatedIP = true
}

func VerifyPEXHandlerSimple(t *testing.T) {
	r, registry := instantiateHandler(&HandlerSettings{})
	defer shutdownHandler(registry)

	assert.NotNil(t, r)
	assert.NotEmpty(t, r.FetchStreams())
}

func VerifyPEXHandlerAppendDeleteNode(t *testing.T) {
	r, registry := instantiateHandler(&HandlerSettings{})
	defer shutdownHandler(registry)

	volume := registry.Volume()
	node := p2p.InstantiateArbitraryNode(false)

	r.AppendNode(node)
	assert.Equal(t, volume+1, registry.Volume())

	r.DeleteNode(node, "REDACTED")

	outgoingNode := p2p.InstantiateArbitraryNode(true)

	r.AppendNode(outgoingNode)
	assert.Equal(t, volume+1, registry.Volume(), "REDACTED")

	r.DeleteNode(outgoingNode, "REDACTED")
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
func VerifyPEXHandlerActive(t *testing.T) {
	N := 3
	routers := make([]*p2p.Router, N)

	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	registries := make([]AddressLedger, N)
	tracer := log.VerifyingTracer()

	//
	for i := 0; i < N; i++ {
		routers[i] = p2p.CreateRouter(cfg, i, func(i int, sw *p2p.Router) *p2p.Router {
			registries[i] = NewAddressRegistry(filepath.Join(dir, fmt.Sprintf("REDACTED", i)), false)
			registries[i].AssignTracer(tracer.With("REDACTED", i))
			sw.CollectionAddressLedger(registries[i])

			sw.AssignTracer(tracer.With("REDACTED", i))

			r := NewHandler(registries[i], &HandlerSettings{})
			r.AssignTracer(tracer.With("REDACTED", i))
			r.CollectionAssureNodesDuration(250 * time.Millisecond)
			sw.AppendHandler("REDACTED", r)

			return sw
		})
	}

	appendAnotherMemberAddressToAddressRegistry := func(routerOrdinal, anotherRouterOrdinal int) {
		address := routers[anotherRouterOrdinal].NetLocation()
		err := registries[routerOrdinal].AppendLocation(address, address)
		require.NoError(t, err)
	}

	appendAnotherMemberAddressToAddressRegistry(0, 1)
	appendAnotherMemberAddressToAddressRegistry(1, 0)
	appendAnotherMemberAddressToAddressRegistry(2, 1)

	for _, sw := range routers {
		err := sw.Begin() //
		require.Nil(t, err)
	}

	affirmNodesWithDeadline(t, routers, 10*time.Millisecond, 10*time.Second, N-1)

	//
	for _, s := range routers {
		err := s.Halt()
		require.NoError(t, err)
	}
}

func VerifyPEXHandlerAccept(t *testing.T) {
	r, registry := instantiateHandler(&HandlerSettings{})
	defer shutdownHandler(registry)

	node := p2p.InstantiateArbitraryNode(false)

	//
	r.QueryLocations(node)

	volume := registry.Volume()
	msg := &tmp2p.PexLocations{Locations: []tmp2p.NetLocation{node.SocketAddress().ToSchema()}}
	r.Accept(p2p.Packet{StreamUID: PexConduit, Src: node, Signal: msg})
	assert.Equal(t, volume+1, registry.Volume())

	r.Accept(p2p.Packet{StreamUID: PexConduit, Src: node, Signal: &tmp2p.PexQuery{}})
}

func VerifyPEXHandlerQuerySignalMisuse(t *testing.T) {
	r, registry := instantiateHandler(&HandlerSettings{})
	defer shutdownHandler(registry)

	sw := instantiateRouterAndAppendHandlers(r)
	sw.CollectionAddressLedger(registry)

	node := emulate.NewNode(nil)
	nodeAddress := node.SocketAddress()
	p2p.AppendNodeToRouterNodeCollection(sw, node)
	assert.True(t, sw.Nodes().Has(node.ID()))
	err := registry.AppendLocation(nodeAddress, nodeAddress)
	require.NoError(t, err)
	require.True(t, registry.HasLocation(nodeAddress))

	id := string(node.ID())

	//
	r.Accept(p2p.Packet{StreamUID: PexConduit, Src: node, Signal: &tmp2p.PexQuery{}})
	assert.True(t, r.finalAcceptedQueries.Has(id))
	assert.True(t, sw.Nodes().Has(node.ID()))

	//
	r.Accept(p2p.Packet{StreamUID: PexConduit, Src: node, Signal: &tmp2p.PexQuery{}})
	assert.True(t, r.finalAcceptedQueries.Has(id))
	assert.True(t, sw.Nodes().Has(node.ID()))

	//
	r.Accept(p2p.Packet{StreamUID: PexConduit, Src: node, Signal: &tmp2p.PexQuery{}})
	assert.False(t, r.finalAcceptedQueries.Has(id))
	assert.False(t, sw.Nodes().Has(node.ID()))
	assert.True(t, registry.IsProhibited(nodeAddress))
}

func VerifyPEXHandlerLocationsSignalMisuse(t *testing.T) {
	r, registry := instantiateHandler(&HandlerSettings{})
	defer shutdownHandler(registry)

	sw := instantiateRouterAndAppendHandlers(r)
	sw.CollectionAddressLedger(registry)

	node := emulate.NewNode(nil)
	p2p.AppendNodeToRouterNodeCollection(sw, node)
	assert.True(t, sw.Nodes().Has(node.ID()))

	id := string(node.ID())

	//
	r.QueryLocations(node)
	assert.True(t, r.queriesRelayed.Has(id))
	assert.True(t, sw.Nodes().Has(node.ID()))

	msg := &tmp2p.PexLocations{Locations: []tmp2p.NetLocation{node.SocketAddress().ToSchema()}}

	//
	r.Accept(p2p.Packet{StreamUID: PexConduit, Src: node, Signal: msg})
	assert.False(t, r.queriesRelayed.Has(id))
	assert.True(t, sw.Nodes().Has(node.ID()))

	//
	r.Accept(p2p.Packet{StreamUID: PexConduit, Src: node, Signal: msg})
	assert.False(t, sw.Nodes().Has(node.ID()))
	assert.True(t, registry.IsProhibited(node.SocketAddress()))
}

func VerifyInspectOrigins(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	//
	nodeRouter := verifyInstantiateStandardNode(dir, 0)
	require.Nil(t, nodeRouter.Begin())
	nodeRouter.Halt() //

	//
	origin := verifyInstantiateSource(dir, 1, []*p2p.NetLocation{}, []*p2p.NetLocation{})

	//
	nodeRouter = verifyInstantiateNodeWithSource(dir, 2, origin)
	require.Nil(t, nodeRouter.Begin())
	nodeRouter.Halt() //

	//
	flawedNodeSettings := &HandlerSettings{
		Origins: []string{
			"REDACTED",
			"REDACTED",
		},
	}
	nodeRouter = verifyInstantiateNodeWithSettings(dir, 2, flawedNodeSettings)
	require.Error(t, nodeRouter.Begin())
	nodeRouter.Halt() //

	//
	flawedNodeSettings = &HandlerSettings{
		Origins: []string{
			"REDACTED",
			"REDACTED",
			origin.NetLocation().String(),
		},
	}
	nodeRouter = verifyInstantiateNodeWithSettings(dir, 2, flawedNodeSettings)
	require.Nil(t, nodeRouter.Begin())
	nodeRouter.Halt() //
}

func VerifyPEXHandlerEmploysOriginsIfRequired(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	//
	origin := verifyInstantiateSource(dir, 0, []*p2p.NetLocation{}, []*p2p.NetLocation{})
	require.Nil(t, origin.Begin())
	defer origin.Halt() //

	//
	node := verifyInstantiateNodeWithSource(dir, 1, origin)
	require.Nil(t, node.Begin())
	defer node.Halt() //

	//
	affirmNodesWithDeadline(t, []*p2p.Router{node}, 10*time.Millisecond, 3*time.Second, 1)
}

func VerifyLinkageVelocityForNodeAcceptedFromSource(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	var id int
	var recognizedLocations []*p2p.NetLocation

	//
	for id = 0; id < 3; id++ {
		node := verifyInstantiateStandardNode(dir, id)
		require.NoError(t, node.Begin())
		address := node.NetLocation()
		defer node.Halt() //

		recognizedLocations = append(recognizedLocations, address)
		t.Log("REDACTED", id, address)
	}

	//
	origin := verifyInstantiateSource(dir, id, recognizedLocations, recognizedLocations)
	require.NoError(t, origin.Begin())
	defer origin.Halt() //
	t.Log("REDACTED", id, origin.NetLocation())

	//
	id++
	member := verifyInstantiateNodeWithSource(dir, id, origin)
	require.NoError(t, member.Begin())
	defer member.Halt() //
	t.Log("REDACTED", id, member.NetLocation())

	//
	affirmNodesWithDeadline(t, []*p2p.Router{member}, 10*time.Millisecond, 3*time.Second, 1)

	//
	affirmNodesWithDeadline(t, []*p2p.Router{member}, 10*time.Millisecond, 1*time.Second, 2)

	//
	//
	outgoing, incoming, calling := member.CountNodes()
	assert.LessOrEqual(t, incoming, cfg.MaximumCountIncomingNodes)
	assert.LessOrEqual(t, outgoing, cfg.MaximumCountOutgoingNodes)
	assert.LessOrEqual(t, calling, cfg.MaximumCountOutgoingNodes+cfg.MaximumCountIncomingNodes-outgoing-incoming)
}

func VerifyPEXHandlerSourceStyle(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	pexReaderSettings := &HandlerSettings{OriginStyle: true, SourceDetachWaitDuration: 10 * time.Millisecond}
	pexReader, registry := instantiateHandler(pexReaderSettings)
	defer shutdownHandler(registry)

	sw := instantiateRouterAndAppendHandlers(pexReader)
	sw.CollectionAddressLedger(registry)
	err = sw.Begin()
	require.NoError(t, err)
	defer sw.Halt() //

	assert.Zero(t, sw.Nodes().Volume())

	nodeRouter := verifyInstantiateStandardNode(dir, 1)
	require.NoError(t, nodeRouter.Begin())
	defer nodeRouter.Halt() //

	//
	pexReader.scanNodes([]*p2p.NetLocation{nodeRouter.NetLocation()})
	assert.Equal(t, 1, sw.Nodes().Volume())
	assert.True(t, sw.Nodes().Has(nodeRouter.MemberDetails().ID()))

	//
	pexReader.endeavorUnlinks()
	assert.Equal(t, 1, sw.Nodes().Volume())

	//
	time.Sleep(pexReaderSettings.SourceDetachWaitDuration + 1*time.Millisecond)

	//
	pexReader.endeavorUnlinks()
	assert.Equal(t, 0, sw.Nodes().Volume())
}

func VerifyPEXHandlerDoesNegateDetachFromDurableNodeInSourceStyle(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	pexReaderSettings := &HandlerSettings{OriginStyle: true, SourceDetachWaitDuration: 1 * time.Millisecond}
	pexReader, registry := instantiateHandler(pexReaderSettings)
	defer shutdownHandler(registry)

	sw := instantiateRouterAndAppendHandlers(pexReader)
	sw.CollectionAddressLedger(registry)
	err = sw.Begin()
	require.NoError(t, err)
	defer sw.Halt() //

	assert.Zero(t, sw.Nodes().Volume())

	nodeRouter := verifyInstantiateStandardNode(dir, 1)
	require.NoError(t, nodeRouter.Begin())
	defer nodeRouter.Halt() //

	err = sw.AppendDurableNodes([]string{nodeRouter.NetLocation().String()})
	require.NoError(t, err)

	//
	pexReader.scanNodes([]*p2p.NetLocation{nodeRouter.NetLocation()})
	assert.Equal(t, 1, sw.Nodes().Volume())
	assert.True(t, sw.Nodes().Has(nodeRouter.MemberDetails().ID()))

	//
	time.Sleep(pexReaderSettings.SourceDetachWaitDuration + 1*time.Millisecond)

	//
	pexReader.endeavorUnlinks()
	assert.Equal(t, 1, sw.Nodes().Volume())
}

func VerifyPEXHandlerCallsNodeUpToMaximumTriesInSourceStyle(t *testing.T) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	pexReader, registry := instantiateHandler(&HandlerSettings{OriginStyle: true})
	defer shutdownHandler(registry)

	sw := instantiateRouterAndAppendHandlers(pexReader)
	sw.CollectionAddressLedger(registry)
	//

	node := emulate.NewNode(nil)
	address := node.SocketAddress()

	err = registry.AppendLocation(address, address)
	require.NoError(t, err)

	assert.True(t, registry.HasLocation(address))

	//
	pexReader.triesToCall.Store(address.CallString(), _attemptstodial{maximumTriesToCall + 1, time.Now()})
	pexReader.scanNodes([]*p2p.NetLocation{address})

	assert.False(t, registry.HasLocation(address))
}

//
//
//
//
//
func VerifyPEXHandlerSourceStylePurgeHalt(t *testing.T) {
	N := 2
	routers := make([]*p2p.Router, N)

	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.Nil(t, err)
	defer os.RemoveAll(dir)

	registries := make([]AddressLedger, N)
	tracer := log.VerifyingTracer()

	//
	for i := 0; i < N; i++ {
		routers[i] = p2p.CreateRouter(cfg, i, func(i int, sw *p2p.Router) *p2p.Router {
			registries[i] = NewAddressRegistry(filepath.Join(dir, fmt.Sprintf("REDACTED", i)), false)
			registries[i].AssignTracer(tracer.With("REDACTED", i))
			sw.CollectionAddressLedger(registries[i])

			sw.AssignTracer(tracer.With("REDACTED", i))

			settings := &HandlerSettings{}
			if i == 0 {
				//
				settings = &HandlerSettings{OriginStyle: true}
			}
			r := NewHandler(registries[i], settings)
			r.AssignTracer(tracer.With("REDACTED", i))
			r.CollectionAssureNodesDuration(250 * time.Millisecond)
			sw.AppendHandler("REDACTED", r)

			return sw
		})
	}

	for _, sw := range routers {
		err := sw.Begin() //
		require.Nil(t, err)
	}

	handler := routers[0].Handlers()["REDACTED"].(*Handler)
	nodeUID := routers[1].MemberDetails().ID()

	err = routers[1].CallNodeWithLocation(routers[0].NetLocation())
	assert.NoError(t, err)

	//
	//
	//
	for i := 0; i < 1000; i++ {
		v := handler.finalAcceptedQueries.Get(string(nodeUID))
		if v != nil {
			break
		}
		time.Sleep(time.Millisecond)
	}

	//
	//
	nodes := routers[0].Nodes().Clone()
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

func VerifyPEXHandlerDoesNegateAppendInternalNodesToAddressRegistry(t *testing.T) {
	node := p2p.InstantiateArbitraryNode(false)

	pexReader, registry := instantiateHandler(&HandlerSettings{})
	registry.AppendInternalIDXDatastore([]string{string(node.MemberDetails().ID())})
	defer shutdownHandler(registry)

	//
	pexReader.QueryLocations(node)

	volume := registry.Volume()
	msg := &tmp2p.PexLocations{Locations: []tmp2p.NetLocation{node.SocketAddress().ToSchema()}}
	pexReader.Accept(p2p.Packet{
		StreamUID: PexConduit,
		Src:       node,
		Signal:   msg,
	})
	assert.Equal(t, volume, registry.Volume())

	pexReader.AppendNode(node)
	assert.Equal(t, volume, registry.Volume())
}

func VerifyPEXHandlerCallNode(t *testing.T) {
	pexReader, registry := instantiateHandler(&HandlerSettings{})
	defer shutdownHandler(registry)

	sw := instantiateRouterAndAppendHandlers(pexReader)
	sw.CollectionAddressLedger(registry)

	node := emulate.NewNode(nil)
	address := node.SocketAddress()

	assert.Equal(t, 0, pexReader.TriesToCall(address))

	//
	err := pexReader.callNode(address)
	require.Error(t, err)

	assert.Equal(t, 1, pexReader.TriesToCall(address))

	//
	err = pexReader.callNode(address)
	require.Error(t, err)

	//
	assert.Equal(t, 1, pexReader.TriesToCall(address))

	if !testing.Short() {
		time.Sleep(3 * time.Second)

		//
		err = pexReader.callNode(address)
		require.Error(t, err)

		assert.Equal(t, 2, pexReader.TriesToCall(address))
	}
}

func affirmNodesWithDeadline(
	t *testing.T,
	routers []*p2p.Router,
	inspectDuration, deadline time.Duration,
	nNodes int,
) {
	var (
		timer    = time.NewTicker(inspectDuration)
		deadlineChan = time.After(deadline)
	)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			//
			allValid := true
			for _, s := range routers {
				outgoing, incoming, _ := s.CountNodes()
				if outgoing+incoming < nNodes {
					allValid = false
					break
				}
			}
			if allValid {
				return
			}
		case <-deadlineChan:
			countNodesStr := "REDACTED"
			for i, s := range routers {
				outgoing, incoming, _ := s.CountNodes()
				countNodesStr += fmt.Sprintf("REDACTED", i, outgoing, incoming)
			}
			t.Errorf(
				"REDACTED",
				nNodes, countNodesStr,
			)
			return
		}
	}
}

//
func verifyInstantiateNodeWithSettings(dir string, id int, settings *HandlerSettings) *p2p.Router {
	node := p2p.CreateRouter(
		cfg,
		id,
		func(i int, sw *p2p.Router) *p2p.Router {
			registry := NewAddressRegistry(filepath.Join(dir, fmt.Sprintf("REDACTED", id)), false)
			registry.AssignTracer(log.VerifyingTracer().With("REDACTED", id))
			sw.CollectionAddressLedger(registry)

			r := NewHandler(
				registry,
				settings,
			)
			r.AssignTracer(log.VerifyingTracer().With("REDACTED", id))
			sw.AppendHandler("REDACTED", r)
			return sw
		},
	)
	return node
}

//
func verifyInstantiateStandardNode(dir string, id int) *p2p.Router {
	return verifyInstantiateNodeWithSettings(dir, id, &HandlerSettings{})
}

//
//
func verifyInstantiateSource(dir string, id int, recognizedLocations, originLocations []*p2p.NetLocation) *p2p.Router {
	origin := p2p.CreateRouter(
		cfg,
		id,
		func(i int, sw *p2p.Router) *p2p.Router {
			registry := NewAddressRegistry(filepath.Join(dir, "REDACTED"), false)
			registry.AssignTracer(log.VerifyingTracer())
			for j := 0; j < len(recognizedLocations); j++ {
				registry.AppendLocation(recognizedLocations[j], originLocations[j]) //
				registry.StampValid(recognizedLocations[j].ID)
			}
			sw.CollectionAddressLedger(registry)

			sw.AssignTracer(log.VerifyingTracer())

			r := NewHandler(registry, &HandlerSettings{})
			r.AssignTracer(log.VerifyingTracer())
			sw.AppendHandler("REDACTED", r)
			return sw
		},
	)
	return origin
}

//
//
func verifyInstantiateNodeWithSource(dir string, id int, origin *p2p.Router) *p2p.Router {
	cfg := &HandlerSettings{
		Origins: []string{origin.NetLocation().String()},
	}
	return verifyInstantiateNodeWithSettings(dir, id, cfg)
}

func instantiateHandler(cfg *HandlerSettings) (r *Handler, registry AddressLedger) {
	//
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	if err != nil {
		panic(err)
	}
	registry = NewAddressRegistry(filepath.Join(dir, "REDACTED"), true)
	registry.AssignTracer(log.VerifyingTracer())

	r = NewHandler(registry, cfg)
	r.AssignTracer(log.VerifyingTracer())
	return
}

func shutdownHandler(registry AddressLedger) {
	//
	err := os.RemoveAll(filepath.Dir(registry.(*addressLedger).EntryRoute()))
	if err != nil {
		panic(err)
	}
}

func instantiateRouterAndAppendHandlers(handlers ...p2p.Handler) *p2p.Router {
	sw := p2p.CreateRouter(cfg, 0, func(i int, sw *p2p.Router) *p2p.Router { return sw })
	sw.AssignTracer(log.VerifyingTracer())
	for _, r := range handlers {
		sw.AppendHandler(r.String(), r)
		r.CollectionRouter(sw)
	}
	return sw
}

func VerifyPexArrays(t *testing.T) {
	address := tmp2p.NetLocation{
		ID:   "REDACTED",
		IP:   "REDACTED",
		Port: 9090,
	}

	verifyScenarios := []struct {
		verifyLabel string
		msg      proto.Message
		expirationOctets string
	}{
		{"REDACTED", &tmp2p.PexQuery{}, "REDACTED"},
		{"REDACTED", &tmp2p.PexLocations{Locations: []tmp2p.NetLocation{address}}, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		w := tc.msg.(p2p.Adapter).Enclose()
		bz, err := proto.Marshal(w)
		require.NoError(t, err)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyLabel)
	}
}
