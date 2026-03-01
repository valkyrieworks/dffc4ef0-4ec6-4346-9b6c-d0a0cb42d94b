package p2p

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/componentindex"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
)

const (
	//
	//
	callGeneratorDurationMillis = 3000

	//
	//
	reestablishEndeavors = 20
	reestablishDuration = 5 * time.Second

	//
	//
	reestablishRearUnactivatedEndeavors    = 10
	reestablishRearUnactivatedFoundationMoments = 3
)

//
//
func ModuleLinkSettings(cfg *settings.Peer2peerSettings) link.ModuleLinkSettings {
	moduleSettings := link.FallbackModuleLinkSettings()
	moduleSettings.PurgeRegulate = cfg.PurgeRegulateDeadline
	moduleSettings.TransmitFrequency = cfg.TransmitFrequency
	moduleSettings.ObtainFrequency = cfg.ObtainFrequency
	moduleSettings.MaximumPacketSignalWorkloadExtent = cfg.MaximumPacketSignalWorkloadExtent
	moduleSettings.VerifyRandomize = cfg.VerifyRandomize
	moduleSettings.VerifyRandomizeSettings = cfg.VerifyRandomizeSettings
	return moduleSettings
}

//

//
//
type LocationRegister interface {
	AppendLocator(location *NetworkLocator, src *NetworkLocator) error
	AppendSecludedIDXDstore([]string)
	AppendMineLocator(*NetworkLocator)
	MineLocator(*NetworkLocator) bool
	LabelValid(ID)
	DiscardLocator(*NetworkLocator)
	OwnsLocation(*NetworkLocator) bool
	Persist()
}

//
//
type NodeRefineMethod func(IDXNodeAssign, Node) error

//

//
//
//
//
type Router struct {
	facility.FoundationFacility

	settings        *settings.Peer2peerSettings
	engines      map[string]Handler
	chnlDescriptions       []*link.ConduitDefinition
	enginesViaChnl  map[byte]Handler
	signalKindViaChnlUUID map[byte]proto.Message
	nodes         *NodeAssign
	calling       *componentindex.CNIndex
	reestablishing  *componentindex.CNIndex
	peerDetails      PeerDetails //
	peerToken       *PeerToken //
	locationRegister      LocationRegister
	//
	enduringNodesLocations []*NetworkLocator
	absoluteNodeIDXDstore map[ID]struct{}

	carrier Carrier

	refineDeadline time.Duration
	nodeCriteria   []NodeRefineMethod

	rng *arbitrary.Arbitrary //

	telemetry *Telemetry
	mlc     *telemetryTagStash
}

var _ Router = (*Router)(nil)

//
func (sw *Router) NetworkLocator() *NetworkLocator {
	location := sw.carrier.NetworkLocator()
	return &location
}

//
type RouterSelection func(*Router)

//
func FreshRouter(
	cfg *settings.Peer2peerSettings,
	carrier Carrier,
	choices ...RouterSelection,
) *Router {
	sw := &Router{
		settings:               cfg,
		engines:             make(map[string]Handler),
		chnlDescriptions:              make([]*link.ConduitDefinition, 0),
		enginesViaChnl:         make(map[byte]Handler),
		signalKindViaChnlUUID:        make(map[byte]proto.Message),
		nodes:                FreshNodeAssign(),
		calling:              componentindex.FreshCNIndex(),
		reestablishing:         componentindex.FreshCNIndex(),
		telemetry:              NooperationTelemetry(),
		carrier:            carrier,
		refineDeadline:        fallbackRefineDeadline,
		enduringNodesLocations: make([]*NetworkLocator, 0),
		absoluteNodeIDXDstore: make(map[ID]struct{}),
		mlc:                  freshTelemetryTagStash(),
	}

	//
	sw.rng = arbitrary.FreshArbitrary()

	sw.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", sw)

	for _, selection := range choices {
		selection(sw)
	}

	return sw
}

//
func RouterRefineDeadline(deadline time.Duration) RouterSelection {
	return func(sw *Router) { sw.refineDeadline = deadline }
}

//
func RouterNodeCriteria(criteria ...NodeRefineMethod) RouterSelection {
	return func(sw *Router) { sw.nodeCriteria = criteria }
}

//
func UsingTelemetry(telemetry *Telemetry) RouterSelection {
	return func(sw *Router) { sw.telemetry = telemetry }
}

//
//

//
//
func (sw *Router) AppendHandler(alias string, handler Handler) Handler {
	for _, chnlDescription := range handler.ObtainConduits() {
		chnlUUID := chnlDescription.ID
		//
		if sw.enginesViaChnl[chnlUUID] != nil {
			panic(fmt.Sprintf("REDACTED", chnlUUID, sw.enginesViaChnl[chnlUUID], handler))
		}
		sw.chnlDescriptions = append(sw.chnlDescriptions, chnlDescription)
		sw.enginesViaChnl[chnlUUID] = handler
		sw.signalKindViaChnlUUID[chnlUUID] = chnlDescription.SignalKind
	}
	sw.engines[alias] = handler
	handler.AssignRouter(sw)
	return handler
}

//
//
func (sw *Router) DiscardHandler(alias string, handler Handler) {
	for _, chnlDescription := range handler.ObtainConduits() {
		//
		for i := 0; i < len(sw.chnlDescriptions); i++ {
			if chnlDescription.ID == sw.chnlDescriptions[i].ID {
				sw.chnlDescriptions = append(sw.chnlDescriptions[:i], sw.chnlDescriptions[i+1:]...)
				break
			}
		}
		delete(sw.enginesViaChnl, chnlDescription.ID)
		delete(sw.signalKindViaChnlUUID, chnlDescription.ID)
	}
	delete(sw.engines, alias)
	handler.AssignRouter(nil)
}

//
//
func (sw *Router) Engines() map[string]Handler {
	return sw.engines
}

//
//
func (sw *Router) Handler(alias string) (Handler, bool) {
	handler, ok := sw.engines[alias]
	return handler, ok
}

//
//
func (sw *Router) AssignPeerDetails(peerDetails PeerDetails) {
	sw.peerDetails = peerDetails
}

//
//
func (sw *Router) PeerDetails() PeerDetails {
	return sw.peerDetails
}

//
//
func (sw *Router) AssignPeerToken(peerToken *PeerToken) {
	sw.peerToken = peerToken
}

//
//

//
func (sw *Router) UponInitiate() error {
	//
	for _, handler := range sw.engines {
		err := handler.Initiate()
		if err != nil {
			return fmt.Errorf("REDACTED", handler, err)
		}
	}

	//
	go sw.embraceProcedure()

	return nil
}

//
func (sw *Router) UponHalt() {
	//
	for _, p := range sw.nodes.Duplicate() {
		sw.haltAlsoDiscardNode(p, nil)
	}

	//
	sw.Tracer.Diagnose("REDACTED")
	for _, handler := range sw.engines {
		if err := handler.Halt(); err != nil {
			sw.Tracer.Failure("REDACTED", "REDACTED", handler, "REDACTED", err)
		}
	}
}

//
//

//
//
//
//
//
//
func (sw *Router) Multicast(e Wrapper) chan bool {
	sw.Tracer.Diagnose("REDACTED", "REDACTED", e.ConduitUUID)

	var wg sync.WaitGroup
	triumphChn := make(chan bool, sw.nodes.Extent())

	sw.nodes.ForeachEvery(func(p Node) {
		wg.Add(1) //
		go func(node Node) {
			defer wg.Done()
			triumph := node.Transmit(e)
			//
			select {
			case triumphChn <- triumph:
			default:
			}
		}(p)
	})

	go func() {
		wg.Wait()
		close(triumphChn)
	}()

	return triumphChn
}

//
//
//
//
//
func (sw *Router) MulticastAsyncronous(e Wrapper) {
	sw.Tracer.Diagnose("REDACTED", "REDACTED", e.ConduitUUID)

	sw.nodes.ForeachEvery(func(p Node) {
		go func(node Node) {
			triumph := node.Transmit(e)
			_ = triumph
		}(p)
	})
}

//
//
//
//
//
//
func (sw *Router) AttemptMulticast(e Wrapper) {
	sw.nodes.ForeachEvery(func(p Node) {
		go func(node Node) {
			node.AttemptTransmit(e)
		}(p)
	})
}

//
//
func (sw *Router) CountNodes() (outgoing, incoming, calling int) {
	sw.nodes.ForeachEvery(func(node Node) {
		if node.EqualsOutgoing() && !sw.EqualsNodeAbsolute(node.ID()) {
			outgoing++
		} else if !sw.EqualsNodeAbsolute(node.ID()) {
			incoming++
		}
	})
	calling = sw.calling.Extent()
	return
}

func (sw *Router) EqualsNodeAbsolute(id ID) bool {
	_, ok := sw.absoluteNodeIDXDstore[id]
	return ok
}

//
func (sw *Router) MaximumCountOutgoingNodes() int {
	return sw.settings.MaximumCountOutgoingNodes
}

//
func (sw *Router) Nodes() IDXNodeAssign {
	return sw.nodes
}

//
//
//
func (sw *Router) HaltNodeForeachFailure(node Node, rationale any) {
	if !node.EqualsActive() {
		return
	}

	sw.Tracer.Failure("REDACTED", "REDACTED", node, "REDACTED", rationale)
	sw.haltAlsoDiscardNode(node, rationale)

	if node.EqualsEnduring() {
		var location *NetworkLocator
		if node.EqualsOutgoing() { //
			location = node.PortLocation()
		} else { //
			var err error
			location, err = node.PeerDetails().NetworkLocator()
			if err != nil {
				sw.Tracer.Failure("REDACTED",
					"REDACTED", node, "REDACTED", err)
				return
			}
		}
		go sw.reestablishTowardNode(location)
	}
}

//
//
func (sw *Router) HaltNodeSmoothly(node Node) {
	sw.Tracer.Details("REDACTED")
	sw.haltAlsoDiscardNode(node, nil)
}

func (sw *Router) haltAlsoDiscardNode(node Node, rationale any) {
	//
	//
	if err := node.Halt(); err != nil {
		sw.Tracer.Failure("REDACTED", "REDACTED", node.ID(), "REDACTED", err)
		return
	}

	sw.carrier.Sanitize(node)
	for _, handler := range sw.engines {
		handler.DiscardNode(node, rationale)
	}

	//
	//
	//
	//
	if !sw.nodes.Discard(node) {
		//
		//
		sw.Tracer.Diagnose("REDACTED", "REDACTED", node.ID())
		return
	}

	sw.telemetry.Nodes.Add(float64(-1))
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
func (sw *Router) reestablishTowardNode(location *NetworkLocator) {
	if sw.reestablishing.Has(string(location.ID)) {
		return
	}
	sw.reestablishing.Set(string(location.ID), location)
	defer sw.reestablishing.Erase(string(location.ID))

	initiate := time.Now()
	sw.Tracer.Details("REDACTED", "REDACTED", location)

	for i := 0; i < reestablishEndeavors; i++ {
		if !sw.EqualsActive() {
			return
		}

		err := sw.CallNodeUsingLocator(location)
		if err == nil {
			return //
		} else if _, ok := err.(FaultPresentlyCallingEitherPresentLocator); ok {
			return
		}

		sw.Tracer.Details("REDACTED", "REDACTED", i, "REDACTED", err, "REDACTED", location)
		//
		sw.unpredictableSnooze(reestablishDuration)
		continue
	}

	sw.Tracer.Failure("REDACTED",
		"REDACTED", location, "REDACTED", time.Since(initiate))
	for i := 1; i <= reestablishRearUnactivatedEndeavors; i++ {
		if !sw.EqualsActive() {
			return
		}

		//
		snoozeDurationMoments := math.Pow(reestablishRearUnactivatedFoundationMoments, float64(i))
		sw.unpredictableSnooze(time.Duration(snoozeDurationMoments) * time.Second)

		err := sw.CallNodeUsingLocator(location)
		if err == nil {
			return //
		} else if _, ok := err.(FaultPresentlyCallingEitherPresentLocator); ok {
			return
		}
		sw.Tracer.Details("REDACTED", "REDACTED", i, "REDACTED", err, "REDACTED", location)
	}
	sw.Tracer.Failure("REDACTED", "REDACTED", location, "REDACTED", time.Since(initiate))
}

//
func (sw *Router) AssignLocationRegister(locationRegister LocationRegister) {
	sw.locationRegister = locationRegister
}

//
//
func (sw *Router) LabelNodeLikeValid(node Node) {
	if sw.locationRegister != nil {
		sw.locationRegister.LabelValid(node.ID())
	}
}

//
//

type secludedLocation interface {
	SecludedLocation() bool
}

func equalsSecludedLocation(err error) bool {
	te, ok := err.(secludedLocation)
	return ok && te.SecludedLocation()
}

//
//
//
//
//
func (sw *Router) CallNodesAsyncronous(nodes []string) error {
	networkLocations, errors := FreshNetworkLocatorTexts(nodes)
	//
	for _, err := range errors {
		sw.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	//
	for _, err := range errors {
		if _, ok := err.(FaultNetworkLocatorSearch); ok {
			continue
		}
		return err
	}
	sw.callNodesAsyncronous(networkLocations)
	return nil
}

func (sw *Router) callNodesAsyncronous(networkLocations []*NetworkLocator) {
	mineLocation := sw.NetworkLocator()

	//
	//
	//
	//
	if sw.locationRegister != nil {
		//
		for _, networkLocation := range networkLocations {
			//
			if !networkLocation.Identical(mineLocation) {
				if err := sw.locationRegister.AppendLocator(networkLocation, mineLocation); err != nil {
					if equalsSecludedLocation(err) {
						sw.Tracer.Diagnose("REDACTED", "REDACTED", err)
					} else {
						sw.Tracer.Failure("REDACTED", "REDACTED", err)
					}
				}
			}
		}
		//
		//
		sw.locationRegister.Persist()
	}

	//
	mode := sw.rng.Mode(len(networkLocations))
	for i := 0; i < len(mode); i++ {
		go func(i int) {
			j := mode[i]
			location := networkLocations[j]

			if location.Identical(mineLocation) {
				sw.Tracer.Diagnose("REDACTED", "REDACTED", location, "REDACTED", mineLocation)
				return
			}

			sw.unpredictableSnooze(0)

			err := sw.CallNodeUsingLocator(location)
			if err != nil {
				switch err.(type) {
				case FaultRouterRelateTowardEgo, FaultRouterReplicatedNodeUUID, FaultPresentlyCallingEitherPresentLocator:
					sw.Tracer.Diagnose("REDACTED", "REDACTED", err)
				default:
					sw.Tracer.Failure("REDACTED", "REDACTED", err)
				}
			}
		}(i)
	}
}

//
//
//
//
func (sw *Router) CallNodeUsingLocator(location *NetworkLocator) error {
	if sw.EqualsCallingEitherCurrentLocator(location) {
		return FaultPresentlyCallingEitherPresentLocator{location.Text()}
	}

	sw.calling.Set(string(location.ID), location)
	defer sw.calling.Erase(string(location.ID))

	return sw.appendOutgoingNodeUsingSettings(location, sw.settings)
}

//
func (sw *Router) unpredictableSnooze(duration time.Duration) {
	r := time.Duration(sw.rng.Int63num(callGeneratorDurationMillis)) * time.Millisecond
	time.Sleep(r + duration)
}

//
//
func (sw *Router) EqualsCallingEitherCurrentLocator(location *NetworkLocator) bool {
	return sw.calling.Has(string(location.ID)) ||
		sw.nodes.Has(location.ID) ||
		(!sw.settings.PermitReplicatedINET && sw.nodes.OwnsINET(location.IP))
}

//
//
//
func (sw *Router) AppendEnduringNodes(locations []string) error {
	sw.Tracer.Details("REDACTED", "REDACTED", locations)
	networkLocations, errors := FreshNetworkLocatorTexts(locations)
	//
	for _, err := range errors {
		sw.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	//
	for _, err := range errors {
		if _, ok := err.(FaultNetworkLocatorSearch); ok {
			continue
		}
		return err
	}
	sw.enduringNodesLocations = networkLocations
	return nil
}

func (sw *Router) AppendAbsoluteNodeIDXDstore(ids []string) error {
	sw.Tracer.Details("REDACTED", "REDACTED", ids)
	for i, id := range ids {
		err := certifyUUID(ID(id))
		if err != nil {
			return fmt.Errorf("REDACTED", i, err)
		}
		sw.absoluteNodeIDXDstore[ID(id)] = struct{}{}
	}
	return nil
}

func (sw *Router) AppendSecludedNodeIDXDstore(ids []string) error {
	soundIDXDstore := make([]string, 0, len(ids))
	for i, id := range ids {
		err := certifyUUID(ID(id))
		if err != nil {
			return fmt.Errorf("REDACTED", i, err)
		}
		soundIDXDstore = append(soundIDXDstore, id)
	}

	sw.locationRegister.AppendSecludedIDXDstore(soundIDXDstore)

	return nil
}

func (sw *Router) EqualsNodeEnduring(na *NetworkLocator) bool {
	for _, pa := range sw.enduringNodesLocations {
		if pa.Matches(na) {
			return true
		}
	}
	return false
}

func (sw *Router) Log() log.Tracer {
	return sw.Tracer
}

func (sw *Router) embraceProcedure() {
	for {
		p, err := sw.carrier.Embrace(nodeSettings{
			chnlDescriptions:       sw.chnlDescriptions,
			uponNodeFailure:   sw.HaltNodeForeachFailure,
			enginesViaChnl:  sw.enginesViaChnl,
			signalKindViaChnlUUID: sw.signalKindViaChnlUUID,
			telemetry:       sw.telemetry,
			mlc:           sw.mlc,
			equalsEnduring:  sw.EqualsNodeEnduring,
		})
		if err != nil {
			switch err := err.(type) {
			case FaultDeclined:
				if err.EqualsEgo() {
					//
					//
					location := err.Location()
					sw.locationRegister.DiscardLocator(&location)
					sw.locationRegister.AppendMineLocator(&location)
				}

				sw.Tracer.Details(
					"REDACTED",
					"REDACTED", err,
					"REDACTED", sw.nodes.Extent(),
				)

				continue
			case FaultRefineDeadline:
				sw.Tracer.Failure(
					"REDACTED",
					"REDACTED", err,
				)

				continue
			case FaultCarrierTerminated:
				sw.Tracer.Failure(
					"REDACTED",
					"REDACTED", sw.nodes.Extent(),
				)
			default:
				sw.Tracer.Failure(
					"REDACTED",
					"REDACTED", err,
					"REDACTED", sw.nodes.Extent(),
				)
				//
				//
				//
				//
				//
				panic(fmt.Errorf("REDACTED", err))
			}

			break
		}

		if !sw.EqualsNodeAbsolute(p.PeerDetails().ID()) {
			//
			_, in, _ := sw.CountNodes()
			if in >= sw.settings.MaximumCountIncomingNodes {
				sw.Tracer.Details(
					"REDACTED",
					"REDACTED", p.PortLocation(),
					"REDACTED", in,
					"REDACTED", sw.settings.MaximumCountIncomingNodes,
				)

				sw.carrier.Sanitize(p)

				continue
			}

		}

		if err := sw.appendNode(p); err != nil {
			sw.carrier.Sanitize(p)
			if p.EqualsActive() {
				_ = p.Halt()
			}
			sw.Tracer.Details(
				"REDACTED",
				"REDACTED", err,
				"REDACTED", p.ID(),
			)
		}
	}
}

//
//
//
//
//
func (sw *Router) appendOutgoingNodeUsingSettings(
	location *NetworkLocator,
	cfg *settings.Peer2peerSettings,
) error {
	sw.Tracer.Diagnose("REDACTED", "REDACTED", location)

	//
	if cfg.VerifyCallMishap {
		go sw.reestablishTowardNode(location)
		return fmt.Errorf("REDACTED")
	}

	p, err := sw.carrier.Call(*location, nodeSettings{
		chnlDescriptions:       sw.chnlDescriptions,
		uponNodeFailure:   sw.HaltNodeForeachFailure,
		equalsEnduring:  sw.EqualsNodeEnduring,
		enginesViaChnl:  sw.enginesViaChnl,
		signalKindViaChnlUUID: sw.signalKindViaChnlUUID,
		telemetry:       sw.telemetry,
		mlc:           sw.mlc,
	})
	if err != nil {
		if e, ok := err.(FaultDeclined); ok {
			if e.EqualsEgo() {
				//
				//
				sw.locationRegister.DiscardLocator(location)
				sw.locationRegister.AppendMineLocator(location)

				return err
			}
		}

		//
		//
		if sw.EqualsNodeEnduring(location) {
			go sw.reestablishTowardNode(location)
		}

		return err
	}

	if err := sw.appendNode(p); err != nil {
		sw.carrier.Sanitize(p)
		if p.EqualsActive() {
			_ = p.Halt()
		}
		return err
	}

	return nil
}

func (sw *Router) refineNode(p Node) error {
	//
	if sw.nodes.Has(p.ID()) {
		return FaultDeclined{id: p.ID(), equalsReplicated: true}
	}

	faultchnl := make(chan error, len(sw.nodeCriteria))

	for _, f := range sw.nodeCriteria {
		go func(f NodeRefineMethod, p Node, faultchnl chan<- error) {
			faultchnl <- f(sw.nodes, p)
		}(f, p, faultchnl)
	}

	for i := 0; i < cap(faultchnl); i++ {
		select {
		case err := <-faultchnl:
			if err != nil {
				return FaultDeclined{id: p.ID(), err: err, equalsScreened: true}
			}
		case <-time.After(sw.refineDeadline):
			return FaultRefineDeadline{}
		}
	}

	return nil
}

//
//
func (sw *Router) appendNode(p Node) error {
	if err := sw.refineNode(p); err != nil {
		return err
	}

	p.AssignTracer(sw.Tracer.Using("REDACTED", p.PortLocation()))

	//
	//
	if !sw.EqualsActive() {
		//
		sw.Tracer.Failure("REDACTED", "REDACTED", p)
		return nil
	}

	//
	for _, handler := range sw.engines {
		p = handler.InitializeNode(p)
	}

	//
	//
	//
	err := p.Initiate()
	if err != nil {
		//
		sw.Tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", p)
		return err
	}

	//
	//
	//
	if err := sw.nodes.Add(p); err != nil {
		var faultNodeDeletion FaultNodeDeletion
		if errors.As(err, &faultNodeDeletion) {
			sw.Tracer.Failure("REDACTED",
				"REDACTED", "REDACTED",
				"REDACTED", p.ID())
		}
		return err
	}
	sw.telemetry.Nodes.Add(float64(1))

	//
	for _, handler := range sw.engines {
		handler.AppendNode(p)
	}

	sw.Tracer.Diagnose("REDACTED", "REDACTED", p)

	return nil
}
