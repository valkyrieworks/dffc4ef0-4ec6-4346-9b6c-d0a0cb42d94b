package p2p

import (
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/componentindex"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"

	consensuslink "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
)

//

const telemetryMetronomeInterval = 10 * time.Second

//
type Node interface {
	facility.Facility
	PurgeHalt()

	ID() ID               //
	DistantINET() net.IP     //
	DistantLocation() net.Addr //

	EqualsOutgoing() bool   //
	EqualsEnduring() bool //

	ShutdownLink() error //

	PeerDetails() PeerDetails //
	Condition() consensuslink.LinkageCondition
	PortLocation() *NetworkLocator //

	Transmit(Wrapper) bool
	AttemptTransmit(Wrapper) bool

	Set(string, any)
	Get(string) any

	AssignDeletionUnsuccessful()
	ObtainDeletionUnsuccessful() bool
}

//

//
type nodeLink struct {
	outgoing   bool
	enduring bool
	link       net.Conn //

	portLocation *NetworkLocator

	//
	ip net.IP
}

func freshNodeLink(
	outgoing, enduring bool,
	link net.Conn,
	portLocation *NetworkLocator,
) nodeLink {
	return nodeLink{
		outgoing:   outgoing,
		enduring: enduring,
		link:       link,
		portLocation: portLocation,
	}
}

//
//
func (pc nodeLink) ID() ID {
	return PublicTokenTowardUUID(pc.link.(*consensuslink.CredentialLinkage).DistantPublicToken())
}

//
func (pc nodeLink) DistantINET() net.IP {
	if pc.ip != nil {
		return pc.ip
	}

	machine, _, err := net.SplitHostPort(pc.link.RemoteAddr().String())
	if err != nil {
		panic(err)
	}

	ips, err := net.LookupIP(machine)
	if err != nil {
		panic(err)
	}

	pc.ip = ips[0]

	return pc.ip
}

//
//
//
type node struct {
	facility.FoundationFacility

	//
	nodeLink
	multilink *consensuslink.ModuleLinkage

	//
	//
	//
	peerDetails PeerDetails
	conduits []byte

	//
	Data *componentindex.CNIndex

	telemetry *Telemetry
	mlc     *telemetryTagStash

	//
	deletionEffortUnsuccessful bool
}

type NodeSelection func(*node)

func freshNode(
	pc nodeLink,
	moduleSettings consensuslink.ModuleLinkSettings,
	peerDetails PeerDetails,
	enginesViaChnl map[byte]Handler,
	signalKindViaChnlUUID map[byte]proto.Message,
	chnlDescriptions []*consensuslink.ConduitDefinition,
	uponNodeFailure func(Node, any),
	mlc *telemetryTagStash,
	choices ...NodeSelection,
) *node {
	p := &node{
		nodeLink: pc,
		peerDetails: peerDetails,
		conduits: peerDetails.(FallbackPeerDetails).Conduits,
		Data:     componentindex.FreshCNIndex(),
		telemetry:  NooperationTelemetry(),
		mlc:      mlc,
	}

	p.multilink = generateModuleLinkage(
		pc.link,
		p,
		enginesViaChnl,
		signalKindViaChnlUUID,
		chnlDescriptions,
		uponNodeFailure,
		moduleSettings,
	)
	p.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", p)
	for _, selection := range choices {
		selection(p)
	}

	return p
}

//
func (p *node) Text() string {
	if p.outgoing {
		return fmt.Sprintf("REDACTED", p.multilink, p.ID())
	}

	return fmt.Sprintf("REDACTED", p.multilink, p.ID())
}

//
//

//
func (p *node) AssignTracer(l log.Tracer) {
	p.Tracer = l
	p.multilink.AssignTracer(l)
}

//
func (p *node) UponInitiate() error {
	if err := p.FoundationFacility.UponInitiate(); err != nil {
		return err
	}

	if err := p.multilink.Initiate(); err != nil {
		return err
	}

	go p.telemetryMonitor()
	return nil
}

//
//
//
//
func (p *node) PurgeHalt() {
	p.multilink.PurgeHalt() //
}

//
func (p *node) UponHalt() {
	if err := p.multilink.Halt(); err != nil { //
		p.Tracer.Diagnose("REDACTED", "REDACTED", err)
	}
}

//
//

//
func (p *node) ID() ID {
	return p.peerDetails.ID()
}

//
func (p *node) EqualsOutgoing() bool {
	return p.outgoing
}

//
func (p *node) EqualsEnduring() bool {
	return p.enduring
}

//
func (p *node) PeerDetails() PeerDetails {
	return p.peerDetails
}

//
//
//
//
func (p *node) PortLocation() *NetworkLocator {
	return p.portLocation
}

//
func (p *node) Condition() consensuslink.LinkageCondition {
	return p.multilink.Condition()
}

//
//
//
//
func (p *node) Transmit(e Wrapper) bool {
	return p.transmit(e.ConduitUUID, e.Signal, p.multilink.Transmit)
}

//
//
//
//
func (p *node) AttemptTransmit(e Wrapper) bool {
	return p.transmit(e.ConduitUUID, e.Signal, p.multilink.AttemptTransmit)
}

func (p *node) transmit(chnlUUID byte, msg proto.Message, transmitMethod func(byte, []byte) bool) bool {
	if !p.EqualsActive() {
		return false
	} else if !p.ownsConduit(chnlUUID) {
		return false
	}
	measurementTagDatum := p.mlc.DatumTowardMeasurementTag(msg)
	if w, ok := msg.(Encapsulator); ok {
		msg = w.Enclose()
	}
	signalOctets, err := proto.Marshal(msg)
	if err != nil {
		p.Tracer.Failure("REDACTED", "REDACTED", err)
		return false
	}
	res := transmitMethod(chnlUUID, signalOctets)
	if res {
		tags := []string{
			"REDACTED", string(p.ID()),
			"REDACTED", fmt.Sprintf("REDACTED", chnlUUID),
		}
		p.telemetry.NodeTransmitOctetsSum.With(tags...).Add(float64(len(signalOctets)))
		p.telemetry.ArtifactTransmitOctetsSum.With("REDACTED", measurementTagDatum).Add(float64(len(signalOctets)))
	}
	return res
}

//
//
//
func (p *node) Get(key string) any {
	return p.Data.Get(key)
}

//
//
//
func (p *node) Set(key string, data any) {
	p.Data.Set(key, data)
}

//
//
func (p *node) ownsConduit(chnlUUID byte) bool {
	for _, ch := range p.conduits {
		if ch == chnlUUID {
			return true
		}
	}
	return false
}

//
func (p *node) ShutdownLink() error {
	return p.link.Close()
}

func (p *node) AssignDeletionUnsuccessful() {
	p.deletionEffortUnsuccessful = true
}

func (p *node) ObtainDeletionUnsuccessful() bool {
	return p.deletionEffortUnsuccessful
}

//
//
//

//
func (pc *nodeLink) ShutdownLink() {
	pc.link.Close()
}

//
func (p *node) DistantLocation() net.Addr {
	return p.link.RemoteAddr()
}

//
func (p *node) AbleTransmit(chnlUUID byte) bool {
	if !p.EqualsActive() {
		return false
	}
	return p.multilink.AbleTransmit(chnlUUID)
}

//

func NodeTelemetry(telemetry *Telemetry) NodeSelection {
	return func(p *node) {
		p.telemetry = telemetry
	}
}

func (p *node) telemetryMonitor() {
	telemetryMetronome := time.NewTicker(telemetryMetronomeInterval)
	defer telemetryMetronome.Stop()

	for {
		select {
		case <-telemetryMetronome.C:
			condition := p.multilink.Condition()
			var transmitStagingExtent float64
			for _, chnlCondition := range condition.Conduits {
				transmitStagingExtent += float64(chnlCondition.TransmitStagingExtent)
			}

			p.telemetry.NodeTransmitStagingExtent.With("REDACTED", string(p.ID())).Set(transmitStagingExtent)
		case <-p.Exit():
			return
		}
	}
}

//
//

func generateModuleLinkage(
	link net.Conn,
	p *node,
	enginesViaChnl map[byte]Handler,
	signalKindViaChnlUUID map[byte]proto.Message,
	chnlDescriptions []*consensuslink.ConduitDefinition,
	uponNodeFailure func(Node, any),
	settings consensuslink.ModuleLinkSettings,
) *consensuslink.ModuleLinkage {
	uponAccept := func(chnlUUID byte, signalOctets []byte) {
		handler := enginesViaChnl[chnlUUID]
		if handler == nil {
			//
			//
			panic(fmt.Sprintf("REDACTED", chnlUUID))
		}
		mt := signalKindViaChnlUUID[chnlUUID]
		msg := proto.Clone(mt)
		err := proto.Unmarshal(signalOctets, msg)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err, reflect.TypeOf(mt)))
		}
		tags := []string{
			"REDACTED", string(p.ID()),
			"REDACTED", fmt.Sprintf("REDACTED", chnlUUID),
		}
		if w, ok := msg.(Unwrapper); ok {
			msg, err = w.Disclose()
			if err != nil {
				panic(fmt.Errorf("REDACTED", err))
			}
		}
		p.telemetry.NodeAcceptOctetsSum.With(tags...).Add(float64(len(signalOctets)))
		p.telemetry.ArtifactAcceptOctetsSum.With("REDACTED", p.mlc.DatumTowardMeasurementTag(msg)).Add(float64(len(signalOctets)))
		handler.Accept(Wrapper{
			ConduitUUID: chnlUUID,
			Src:       p,
			Signal:   msg,
		})
	}

	uponFailure := func(r any) {
		uponNodeFailure(p, r)
	}

	return consensuslink.FreshModuleLinkageUsingSettings(
		link,
		chnlDescriptions,
		uponAccept,
		uponFailure,
		settings,
	)
}
