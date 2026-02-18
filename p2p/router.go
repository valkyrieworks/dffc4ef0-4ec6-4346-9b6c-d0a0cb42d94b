package p2p

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/cmap"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/p2p/link"
)

const (
	//
	//
	callShufflerCadenceMilliseconds = 3000

	//
	//
	reestablishTries = 20
	reestablishCadence = 5 * time.Second

	//
	//
	reestablishRearOffsetTries    = 10
	reestablishRearOffsetRootMoments = 3
)

//
//
func MLinkSettings(cfg *settings.P2PSettings) link.MLinkSettings {
	mSettings := link.StandardMLinkSettings()
	mSettings.PurgeRegulate = cfg.PurgeRegulateDeadline
	mSettings.TransmitRatio = cfg.TransmitRatio
	mSettings.ReceiveRatio = cfg.ReceiveRatio
	mSettings.MaximumPackageMessageShipmentVolume = cfg.MaximumPackageMessageShipmentVolume
	mSettings.VerifyRandomize = cfg.VerifyRandomize
	mSettings.VerifyRandomizeSettings = cfg.VerifyRandomizeSettings
	return mSettings
}

//

//
//
type AddressLedger interface {
	AppendLocation(address *NetLocation, src *NetLocation) error
	AppendInternalIDXDatastore([]string)
	AppendOurLocation(*NetLocation)
	OurLocation(*NetLocation) bool
	StampValid(ID)
	DeleteLocation(*NetLocation)
	HasLocation(*NetLocation) bool
	Persist()
}

//
//
type NodeRefineFunction func(IDXNodeCollection, Node) error

//

//
//
//
//
type Router struct {
	daemon.RootDaemon

	settings        *settings.P2PSettings
	handlers      map[string]Handler
	chanTraits       []*link.StreamDefinition
	handlersByChan  map[byte]Handler
	messageKindByChanUID map[byte]proto.Message
	nodes         *NodeCollection
	calling       *cmap.CIndex
	reestablishing  *cmap.CIndex
	memberDetails      MemberDetails //
	memberKey       *MemberKey //
	addressRegistry      AddressLedger
	//
	durableNodesLocations []*NetLocation
	absoluteNodeIDXDatastore map[ID]struct{}

	carrier Carrier

	refineDeadline time.Duration
	nodeScreens   []NodeRefineFunction

	rng *random.Random //

	stats *Stats
	mlc     *statsTagRepository
}

var _ Toggeler = (*Router)(nil)

//
func (sw *Router) NetLocation() *NetLocation {
	address := sw.carrier.NetLocation()
	return &address
}

//
type RouterSetting func(*Router)

//
func NewRouter(
	cfg *settings.P2PSettings,
	carrier Carrier,
	options ...RouterSetting,
) *Router {
	sw := &Router{
		settings:               cfg,
		handlers:             make(map[string]Handler),
		chanTraits:              make([]*link.StreamDefinition, 0),
		handlersByChan:         make(map[byte]Handler),
		messageKindByChanUID:        make(map[byte]proto.Message),
		nodes:                NewNodeCollection(),
		calling:              cmap.NewCIndex(),
		reestablishing:         cmap.NewCIndex(),
		stats:              NoopStats(),
		carrier:            carrier,
		refineDeadline:        standardRefineDeadline,
		durableNodesLocations: make([]*NetLocation, 0),
		absoluteNodeIDXDatastore: make(map[ID]struct{}),
		mlc:                  newStatsTagRepository(),
	}

	//
	sw.rng = random.NewRandom()

	sw.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", sw)

	for _, setting := range options {
		setting(sw)
	}

	return sw
}

//
func RouterRefineDeadline(deadline time.Duration) RouterSetting {
	return func(sw *Router) { sw.refineDeadline = deadline }
}

//
func RouterNodeScreens(screens ...NodeRefineFunction) RouterSetting {
	return func(sw *Router) { sw.nodeScreens = screens }
}

//
func WithStats(stats *Stats) RouterSetting {
	return func(sw *Router) { sw.stats = stats }
}

//
//

//
//
func (sw *Router) AppendHandler(label string, handler Handler) Handler {
	for _, chanNote := range handler.FetchStreams() {
		chanUID := chanNote.ID
		//
		if sw.handlersByChan[chanUID] != nil {
			panic(fmt.Sprintf("REDACTED", chanUID, sw.handlersByChan[chanUID], handler))
		}
		sw.chanTraits = append(sw.chanTraits, chanNote)
		sw.handlersByChan[chanUID] = handler
		sw.messageKindByChanUID[chanUID] = chanNote.SignalKind
	}
	sw.handlers[label] = handler
	handler.CollectionRouter(sw)
	return handler
}

//
//
func (sw *Router) DeleteHandler(label string, handler Handler) {
	for _, chanNote := range handler.FetchStreams() {
		//
		for i := 0; i < len(sw.chanTraits); i++ {
			if chanNote.ID == sw.chanTraits[i].ID {
				sw.chanTraits = append(sw.chanTraits[:i], sw.chanTraits[i+1:]...)
				break
			}
		}
		delete(sw.handlersByChan, chanNote.ID)
		delete(sw.messageKindByChanUID, chanNote.ID)
	}
	delete(sw.handlers, label)
	handler.CollectionRouter(nil)
}

//
//
func (sw *Router) Handlers() map[string]Handler {
	return sw.handlers
}

//
//
func (sw *Router) Handler(label string) (Handler, bool) {
	handler, ok := sw.handlers[label]
	return handler, ok
}

//
//
func (sw *Router) CollectionMemberDetails(memberDetails MemberDetails) {
	sw.memberDetails = memberDetails
}

//
//
func (sw *Router) MemberDetails() MemberDetails {
	return sw.memberDetails
}

//
//
func (sw *Router) CollectionMemberKey(memberKey *MemberKey) {
	sw.memberKey = memberKey
}

//
//

//
func (sw *Router) OnBegin() error {
	//
	for _, handler := range sw.handlers {
		err := handler.Begin()
		if err != nil {
			return fmt.Errorf("REDACTED", handler, err)
		}
	}

	//
	go sw.allowProcedure()

	return nil
}

//
func (sw *Router) OnHalt() {
	//
	for _, p := range sw.nodes.Clone() {
		sw.haltAndDeleteNode(p, nil)
	}

	//
	sw.Tracer.Diagnose("REDACTED")
	for _, handler := range sw.handlers {
		if err := handler.Halt(); err != nil {
			sw.Tracer.Fault("REDACTED", "REDACTED", handler, "REDACTED", err)
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
func (sw *Router) Multicast(e Packet) chan bool {
	sw.Tracer.Diagnose("REDACTED", "REDACTED", e.StreamUID)

	var wg sync.WaitGroup
	successChannel := make(chan bool, sw.nodes.Volume())

	sw.nodes.ForEach(func(p Node) {
		wg.Add(1) //
		go func(node Node) {
			defer wg.Done()
			success := node.Transmit(e)
			//
			select {
			case successChannel <- success:
			default:
			}
		}(p)
	})

	go func() {
		wg.Wait()
		close(successChannel)
	}()

	return successChannel
}

//
//
//
//
//
func (sw *Router) MulticastAsync(e Packet) {
	sw.Tracer.Diagnose("REDACTED", "REDACTED", e.StreamUID)

	sw.nodes.ForEach(func(p Node) {
		go func(node Node) {
			success := node.Transmit(e)
			_ = success
		}(p)
	})
}

//
//
//
//
//
//
func (sw *Router) AttemptMulticast(e Packet) {
	sw.nodes.ForEach(func(p Node) {
		go func(node Node) {
			node.AttemptTransmit(e)
		}(p)
	})
}

//
//
func (sw *Router) CountNodes() (outgoing, incoming, calling int) {
	sw.nodes.ForEach(func(node Node) {
		if node.IsOutgoing() && !sw.IsNodeAbsolute(node.ID()) {
			outgoing++
		} else if !sw.IsNodeAbsolute(node.ID()) {
			incoming++
		}
	})
	calling = sw.calling.Volume()
	return
}

func (sw *Router) IsNodeAbsolute(id ID) bool {
	_, ok := sw.absoluteNodeIDXDatastore[id]
	return ok
}

//
func (sw *Router) MaximumCountOutgoingNodes() int {
	return sw.settings.MaximumCountOutgoingNodes
}

//
func (sw *Router) Nodes() IDXNodeCollection {
	return sw.nodes
}

//
//
//
func (sw *Router) HaltNodeForFault(node Node, cause any) {
	if !node.IsActive() {
		return
	}

	sw.Tracer.Fault("REDACTED", "REDACTED", node, "REDACTED", cause)
	sw.haltAndDeleteNode(node, cause)

	if node.IsDurable() {
		var address *NetLocation
		if node.IsOutgoing() { //
			address = node.SocketAddress()
		} else { //
			var err error
			address, err = node.MemberDetails().NetLocation()
			if err != nil {
				sw.Tracer.Fault("REDACTED",
					"REDACTED", node, "REDACTED", err)
				return
			}
		}
		go sw.reestablishToNode(address)
	}
}

//
//
func (sw *Router) HaltNodeSmoothly(node Node) {
	sw.Tracer.Details("REDACTED")
	sw.haltAndDeleteNode(node, nil)
}

func (sw *Router) haltAndDeleteNode(node Node, cause any) {
	//
	//
	if err := node.Halt(); err != nil {
		sw.Tracer.Fault("REDACTED", "REDACTED", node.ID(), "REDACTED", err)
		return
	}

	sw.carrier.Sanitize(node)
	for _, handler := range sw.handlers {
		handler.DeleteNode(node, cause)
	}

	//
	//
	//
	//
	if !sw.nodes.Delete(node) {
		//
		//
		sw.Tracer.Diagnose("REDACTED", "REDACTED", node.ID())
		return
	}

	sw.stats.Nodes.Add(float64(-1))
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
func (sw *Router) reestablishToNode(address *NetLocation) {
	if sw.reestablishing.Has(string(address.ID)) {
		return
	}
	sw.reestablishing.Set(string(address.ID), address)
	defer sw.reestablishing.Erase(string(address.ID))

	begin := time.Now()
	sw.Tracer.Details("REDACTED", "REDACTED", address)

	for i := 0; i < reestablishTries; i++ {
		if !sw.IsActive() {
			return
		}

		err := sw.CallNodeWithLocation(address)
		if err == nil {
			return //
		} else if _, ok := err.(ErrPresentlyCallingOrCurrentLocation); ok {
			return
		}

		sw.Tracer.Details("REDACTED", "REDACTED", i, "REDACTED", err, "REDACTED", address)
		//
		sw.arbitraryPause(reestablishCadence)
		continue
	}

	sw.Tracer.Fault("REDACTED",
		"REDACTED", address, "REDACTED", time.Since(begin))
	for i := 1; i <= reestablishRearOffsetTries; i++ {
		if !sw.IsActive() {
			return
		}

		//
		pauseCadenceMoments := math.Pow(reestablishRearOffsetRootMoments, float64(i))
		sw.arbitraryPause(time.Duration(pauseCadenceMoments) * time.Second)

		err := sw.CallNodeWithLocation(address)
		if err == nil {
			return //
		} else if _, ok := err.(ErrPresentlyCallingOrCurrentLocation); ok {
			return
		}
		sw.Tracer.Details("REDACTED", "REDACTED", i, "REDACTED", err, "REDACTED", address)
	}
	sw.Tracer.Fault("REDACTED", "REDACTED", address, "REDACTED", time.Since(begin))
}

//
func (sw *Router) CollectionAddressRegistry(addressRegistry AddressLedger) {
	sw.addressRegistry = addressRegistry
}

//
//
func (sw *Router) StampNodeAsSound(node Node) {
	if sw.addressRegistry != nil {
		sw.addressRegistry.StampValid(node.ID())
	}
}

//
//

type internalAddress interface {
	InternalAddress() bool
}

func isInternalAddress(err error) bool {
	te, ok := err.(internalAddress)
	return ok && te.InternalAddress()
}

//
//
//
//
//
func (sw *Router) CallNodesAsync(nodes []string) error {
	netLocations, faults := NewNetLocationStrings(nodes)
	//
	for _, err := range faults {
		sw.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	//
	for _, err := range faults {
		if _, ok := err.(ErrNetLocationSearch); ok {
			continue
		}
		return err
	}
	sw.callNodesAsync(netLocations)
	return nil
}

func (sw *Router) callNodesAsync(netLocations []*NetLocation) {
	ourAddress := sw.NetLocation()

	//
	//
	//
	//
	if sw.addressRegistry != nil {
		//
		for _, netAddress := range netLocations {
			//
			if !netAddress.Identical(ourAddress) {
				if err := sw.addressRegistry.AppendLocation(netAddress, ourAddress); err != nil {
					if isInternalAddress(err) {
						sw.Tracer.Diagnose("REDACTED", "REDACTED", err)
					} else {
						sw.Tracer.Fault("REDACTED", "REDACTED", err)
					}
				}
			}
		}
		//
		//
		sw.addressRegistry.Persist()
	}

	//
	mode := sw.rng.Mode(len(netLocations))
	for i := 0; i < len(mode); i++ {
		go func(i int) {
			j := mode[i]
			address := netLocations[j]

			if address.Identical(ourAddress) {
				sw.Tracer.Diagnose("REDACTED", "REDACTED", address, "REDACTED", ourAddress)
				return
			}

			sw.arbitraryPause(0)

			err := sw.CallNodeWithLocation(address)
			if err != nil {
				switch err.(type) {
				case ErrRouterEstablishToEgo, ErrRouterReplicatedNodeUID, ErrPresentlyCallingOrCurrentLocation:
					sw.Tracer.Diagnose("REDACTED", "REDACTED", err)
				default:
					sw.Tracer.Fault("REDACTED", "REDACTED", err)
				}
			}
		}(i)
	}
}

//
//
//
//
func (sw *Router) CallNodeWithLocation(address *NetLocation) error {
	if sw.IsCallingOrCurrentLocation(address) {
		return ErrPresentlyCallingOrCurrentLocation{address.String()}
	}

	sw.calling.Set(string(address.ID), address)
	defer sw.calling.Erase(string(address.ID))

	return sw.appendOutgoingNodeWithSettings(address, sw.settings)
}

//
func (sw *Router) arbitraryPause(cadence time.Duration) {
	r := time.Duration(sw.rng.Int64count(callShufflerCadenceMilliseconds)) * time.Millisecond
	time.Sleep(r + cadence)
}

//
//
func (sw *Router) IsCallingOrCurrentLocation(address *NetLocation) bool {
	return sw.calling.Has(string(address.ID)) ||
		sw.nodes.Has(address.ID) ||
		(!sw.settings.PermitReplicatedIP && sw.nodes.HasIP(address.IP))
}

//
//
//
func (sw *Router) AppendDurableNodes(locations []string) error {
	sw.Tracer.Details("REDACTED", "REDACTED", locations)
	netLocations, faults := NewNetLocationStrings(locations)
	//
	for _, err := range faults {
		sw.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	//
	for _, err := range faults {
		if _, ok := err.(ErrNetLocationSearch); ok {
			continue
		}
		return err
	}
	sw.durableNodesLocations = netLocations
	return nil
}

func (sw *Router) AppendAbsoluteNodeIDXDatastore(ids []string) error {
	sw.Tracer.Details("REDACTED", "REDACTED", ids)
	for i, id := range ids {
		err := certifyUID(ID(id))
		if err != nil {
			return fmt.Errorf("REDACTED", i, err)
		}
		sw.absoluteNodeIDXDatastore[ID(id)] = struct{}{}
	}
	return nil
}

func (sw *Router) AppendInternalNodeIDXDatastore(ids []string) error {
	soundIDXDatastore := make([]string, 0, len(ids))
	for i, id := range ids {
		err := certifyUID(ID(id))
		if err != nil {
			return fmt.Errorf("REDACTED", i, err)
		}
		soundIDXDatastore = append(soundIDXDatastore, id)
	}

	sw.addressRegistry.AppendInternalIDXDatastore(soundIDXDatastore)

	return nil
}

func (sw *Router) IsNodeDurable(na *NetLocation) bool {
	for _, pa := range sw.durableNodesLocations {
		if pa.Matches(na) {
			return true
		}
	}
	return false
}

func (sw *Router) Log() log.Tracer {
	return sw.Tracer
}

func (sw *Router) allowProcedure() {
	for {
		p, err := sw.carrier.Allow(nodeSettings{
			chanTraits:       sw.chanTraits,
			onNodeFault:   sw.HaltNodeForFault,
			handlersByChan:  sw.handlersByChan,
			messageKindByChanUID: sw.messageKindByChanUID,
			stats:       sw.stats,
			mlc:           sw.mlc,
			isDurable:  sw.IsNodeDurable,
		})
		if err != nil {
			switch err := err.(type) {
			case ErrDeclined:
				if err.IsEgo() {
					//
					//
					address := err.Address()
					sw.addressRegistry.DeleteLocation(&address)
					sw.addressRegistry.AppendOurLocation(&address)
				}

				sw.Tracer.Details(
					"REDACTED",
					"REDACTED", err,
					"REDACTED", sw.nodes.Volume(),
				)

				continue
			case ErrRefineDeadline:
				sw.Tracer.Fault(
					"REDACTED",
					"REDACTED", err,
				)

				continue
			case ErrCarrierHalted:
				sw.Tracer.Fault(
					"REDACTED",
					"REDACTED", sw.nodes.Volume(),
				)
			default:
				sw.Tracer.Fault(
					"REDACTED",
					"REDACTED", err,
					"REDACTED", sw.nodes.Volume(),
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

		if !sw.IsNodeAbsolute(p.MemberDetails().ID()) {
			//
			_, in, _ := sw.CountNodes()
			if in >= sw.settings.MaximumCountIncomingNodes {
				sw.Tracer.Details(
					"REDACTED",
					"REDACTED", p.SocketAddress(),
					"REDACTED", in,
					"REDACTED", sw.settings.MaximumCountIncomingNodes,
				)

				sw.carrier.Sanitize(p)

				continue
			}

		}

		if err := sw.appendNode(p); err != nil {
			sw.carrier.Sanitize(p)
			if p.IsActive() {
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
func (sw *Router) appendOutgoingNodeWithSettings(
	address *NetLocation,
	cfg *settings.P2PSettings,
) error {
	sw.Tracer.Diagnose("REDACTED", "REDACTED", address)

	//
	if cfg.VerifyCallAbort {
		go sw.reestablishToNode(address)
		return fmt.Errorf("REDACTED")
	}

	p, err := sw.carrier.Call(*address, nodeSettings{
		chanTraits:       sw.chanTraits,
		onNodeFault:   sw.HaltNodeForFault,
		isDurable:  sw.IsNodeDurable,
		handlersByChan:  sw.handlersByChan,
		messageKindByChanUID: sw.messageKindByChanUID,
		stats:       sw.stats,
		mlc:           sw.mlc,
	})
	if err != nil {
		if e, ok := err.(ErrDeclined); ok {
			if e.IsEgo() {
				//
				//
				sw.addressRegistry.DeleteLocation(address)
				sw.addressRegistry.AppendOurLocation(address)

				return err
			}
		}

		//
		//
		if sw.IsNodeDurable(address) {
			go sw.reestablishToNode(address)
		}

		return err
	}

	if err := sw.appendNode(p); err != nil {
		sw.carrier.Sanitize(p)
		if p.IsActive() {
			_ = p.Halt()
		}
		return err
	}

	return nil
}

func (sw *Router) refineNode(p Node) error {
	//
	if sw.nodes.Has(p.ID()) {
		return ErrDeclined{id: p.ID(), isReplicated: true}
	}

	faultc := make(chan error, len(sw.nodeScreens))

	for _, f := range sw.nodeScreens {
		go func(f NodeRefineFunction, p Node, faultc chan<- error) {
			faultc <- f(sw.nodes, p)
		}(f, p, faultc)
	}

	for i := 0; i < cap(faultc); i++ {
		select {
		case err := <-faultc:
			if err != nil {
				return ErrDeclined{id: p.ID(), err: err, isScreened: true}
			}
		case <-time.After(sw.refineDeadline):
			return ErrRefineDeadline{}
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

	p.AssignTracer(sw.Tracer.With("REDACTED", p.SocketAddress()))

	//
	//
	if !sw.IsActive() {
		//
		sw.Tracer.Fault("REDACTED", "REDACTED", p)
		return nil
	}

	//
	for _, handler := range sw.handlers {
		p = handler.InitNode(p)
	}

	//
	//
	//
	err := p.Begin()
	if err != nil {
		//
		sw.Tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", p)
		return err
	}

	//
	//
	//
	if err := sw.nodes.Add(p); err != nil {
		var errNodeDeletion ErrNodeDeletion
		if errors.As(err, &errNodeDeletion) {
			sw.Tracer.Fault("REDACTED",
				"REDACTED", "REDACTED",
				"REDACTED", p.ID())
		}
		return err
	}
	sw.stats.Nodes.Add(float64(1))

	//
	for _, handler := range sw.handlers {
		handler.AppendNode(p)
	}

	sw.Tracer.Diagnose("REDACTED", "REDACTED", p)

	return nil
}
