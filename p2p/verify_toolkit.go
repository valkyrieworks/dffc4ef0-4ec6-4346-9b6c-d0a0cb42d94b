package p2p

import (
	"fmt"
	"net"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
)

const verifyChnl = 0x01

//

type simulatePeerDetails struct {
	location *NetworkLocator
}

func (ni simulatePeerDetails) ID() ID                           { return ni.location.ID }
func (ni simulatePeerDetails) NetworkLocator() (*NetworkLocator, error) { return ni.location, nil }
func (ni simulatePeerDetails) Certify() error                  { return nil }
func (ni simulatePeerDetails) MatchedUsing(PeerDetails) error    { return nil }

func AppendNodeTowardRouterNodeAssign(sw *Router, node Node) {
	sw.nodes.Add(node) //
}

func GenerateUnpredictableNode(outgoing bool) Node {
	location, networkLocation := GenerateDirectableLocation()
	p := &node{
		nodeLink: nodeLink{
			outgoing:   outgoing,
			portLocation: networkLocation,
		},
		peerDetails: simulatePeerDetails{networkLocation},
		multilink:    &link.ModuleLinkage{},
		telemetry:  NooperationTelemetry(),
	}
	p.AssignTracer(log.VerifyingTracer().Using("REDACTED", location))
	return p
}

func GenerateDirectableLocation() (location string, networkLocation *NetworkLocator) {
	for {
		var err error
		location = fmt.Sprintf("REDACTED",
			commitrand.Octets(20),
			commitrand.Int()%256,
			commitrand.Int()%256,
			commitrand.Int()%256,
			commitrand.Int()%256)
		networkLocation, err = FreshNetworkLocatorText(location)
		if err != nil {
			panic(err)
		}
		if networkLocation.Directable() {
			break
		}
	}
	return
}

//
//

const VerifyMachine = "REDACTED"

//
//
//
//
func CreateAssociatedRouters(cfg *settings.Peer2peerSettings,
	n int,
	initializeRouter func(int, *Router) *Router,
	relate func([]*Router, int, int),
) []*Router {
	routers := make([]*Router, n)
	for i := 0; i < n; i++ {
		routers[i] = CreateRouter(cfg, i, initializeRouter)
	}

	if err := InitiateRouters(routers); err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			relate(routers, i, j)
		}
	}

	return routers
}

//
//
//
func Connect2routers(routers []*Router, i, j int) {
	routerIDX := routers[i]
	routerJTH := routers[j]

	c1, c2 := link.NetworkTube()

	completeChnl := make(chan struct{})
	go func() {
		err := routerIDX.appendNodeUsingLinkage(c1)
		if err != nil {
			panic(err)
		}
		completeChnl <- struct{}{}
	}()
	go func() {
		err := routerJTH.appendNodeUsingLinkage(c2)
		if err != nil {
			panic(err)
		}
		completeChnl <- struct{}{}
	}()
	<-completeChnl
	<-completeChnl
}

func (sw *Router) appendNodeUsingLinkage(link net.Conn) error {
	pc, err := verifyIncomingNodeLink(link, sw.settings, sw.peerToken.PrivateToken)
	if err != nil {
		if err := link.Close(); err != nil {
			sw.Tracer.Failure("REDACTED", "REDACTED", err)
		}
		return err
	}

	ni, err := negotiation(link, time.Second, sw.peerDetails)
	if err != nil {
		if err := link.Close(); err != nil {
			sw.Tracer.Failure("REDACTED", "REDACTED", err)
		}
		return err
	}

	p := freshNode(
		pc,
		ModuleLinkSettings(sw.settings),
		ni,
		sw.enginesViaChnl,
		sw.signalKindViaChnlUUID,
		sw.chnlDescriptions,
		sw.HaltNodeForeachFailure,
		sw.mlc,
	)

	if err = sw.appendNode(p); err != nil {
		pc.ShutdownLink()
		return err
	}

	return nil
}

//
//
func InitiateRouters(routers []*Router) error {
	for _, s := range routers {
		err := s.Initiate() //
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateRouter(
	cfg *settings.Peer2peerSettings,
	i int,
	initializeRouter func(int, *Router) *Router,
	choices ...RouterSelection,
) *Router {
	peerToken := PeerToken{
		PrivateToken: edwards25519.ProducePrivateToken(),
	}
	peerDetails := verifyPeerDetails(peerToken.ID(), fmt.Sprintf("REDACTED", i))
	location, err := FreshNetworkLocatorText(
		UUIDLocationText(peerToken.ID(), peerDetails.(FallbackPeerDetails).OverhearLocation),
	)
	if err != nil {
		panic(err)
	}

	t := FreshMultiplexCarrier(peerDetails, peerToken, ModuleLinkSettings(cfg))

	if err := t.Overhear(*location); err != nil {
		panic(err)
	}

	//
	sw := initializeRouter(i, FreshRouter(cfg, t, choices...))
	sw.AssignTracer(log.VerifyingTracer().Using("REDACTED", i))
	sw.AssignPeerToken(&peerToken)

	ni := peerDetails.(FallbackPeerDetails)
	for ch := range sw.enginesViaChnl {
		ni.Conduits = append(ni.Conduits, ch)
	}
	peerDetails = ni

	//
	//
	t.peerDetails = peerDetails
	sw.AssignPeerDetails(peerDetails)

	return sw
}

func verifyIncomingNodeLink(
	link net.Conn,
	settings *settings.Peer2peerSettings,
	minePeerPrivateToken security.PrivateToken,
) (nodeLink, error) {
	return verifyNodeLink(link, settings, false, false, minePeerPrivateToken, nil)
}

func verifyNodeLink(
	crudeLink net.Conn,
	cfg *settings.Peer2peerSettings,
	outgoing, enduring bool,
	minePeerPrivateToken security.PrivateToken,
	portLocation *NetworkLocator,
) (pc nodeLink, err error) {
	link := crudeLink

	//
	if cfg.VerifyRandomize {
		//
		link = RandomizeLinkSubsequentOriginatingSettings(link, 10*time.Second, cfg.VerifyRandomizeSettings)
	}

	//
	link, err = modernizeCredentialLink(link, cfg.NegotiationDeadline, minePeerPrivateToken)
	if err != nil {
		return pc, fmt.Errorf("REDACTED", err)
	}

	//
	return freshNodeLink(outgoing, enduring, link, portLocation), nil
}

//
//

func verifyPeerDetails(id ID, alias string) PeerDetails {
	return verifyPeerDetailsUsingFabric(id, alias, "REDACTED")
}

func verifyPeerDetailsUsingFabric(id ID, alias, fabric string) PeerDetails {
	return FallbackPeerDetails{
		SchemeEdition: fallbackSchemeEdition,
		FallbackPeerUUID:   id,
		OverhearLocation:      fmt.Sprintf("REDACTED", obtainLiberateChannel()),
		Fabric:         fabric,
		Edition:         "REDACTED",
		Conduits:        []byte{verifyChnl},
		Pseudonym:         alias,
		Another: FallbackPeerDetailsAnother{
			TransferOrdinal:    "REDACTED",
			RemoteLocator: fmt.Sprintf("REDACTED", obtainLiberateChannel()),
		},
	}
}

func obtainLiberateChannel() int {
	channel, err := strongmindnet.ObtainLiberateChannel()
	if err != nil {
		panic(err)
	}
	return channel
}

type LocationRegisterSimulate struct {
	Locations        map[string]struct{}
	MineLocations     map[string]struct{}
	SecludedLocations map[string]struct{}
}

var _ LocationRegister = (*LocationRegisterSimulate)(nil)

func (register *LocationRegisterSimulate) AppendLocator(location *NetworkLocator, _ *NetworkLocator) error {
	register.Locations[location.Text()] = struct{}{}
	return nil
}
func (register *LocationRegisterSimulate) AppendMineLocator(location *NetworkLocator) { register.MineLocations[location.Text()] = struct{}{} }
func (register *LocationRegisterSimulate) MineLocator(location *NetworkLocator) bool {
	_, ok := register.MineLocations[location.Text()]
	return ok
}
func (register *LocationRegisterSimulate) LabelValid(ID) {}
func (register *LocationRegisterSimulate) OwnsLocation(location *NetworkLocator) bool {
	_, ok := register.Locations[location.Text()]
	return ok
}

func (register *LocationRegisterSimulate) DiscardLocator(location *NetworkLocator) {
	delete(register.Locations, location.Text())
}
func (register *LocationRegisterSimulate) Persist() {}
func (register *LocationRegisterSimulate) AppendSecludedIDXDstore(locations []string) {
	for _, location := range locations {
		register.SecludedLocations[location] = struct{}{}
	}
}
