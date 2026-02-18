package p2p

import (
	"fmt"
	"net"
	"time"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/utils/log"
	cometnet "github.com/valkyrieworks/utils/net"
	engineseed "github.com/valkyrieworks/utils/random"

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/p2p/link"
)

const verifyChan = 0x01

//

type emulateMemberDetails struct {
	address *NetLocation
}

func (ni emulateMemberDetails) ID() ID                           { return ni.address.ID }
func (ni emulateMemberDetails) NetLocation() (*NetLocation, error) { return ni.address, nil }
func (ni emulateMemberDetails) Certify() error                  { return nil }
func (ni emulateMemberDetails) HarmoniousWith(MemberDetails) error    { return nil }

func AppendNodeToRouterNodeCollection(sw *Router, node Node) {
	sw.nodes.Add(node) //
}

func InstantiateArbitraryNode(outgoing bool) Node {
	address, netAddress := InstantiateForwardableAddress()
	p := &node{
		nodeLink: nodeLink{
			outgoing:   outgoing,
			socketAddress: netAddress,
		},
		memberDetails: emulateMemberDetails{netAddress},
		mconn:    &link.MLinkage{},
		stats:  NoopStats(),
	}
	p.AssignTracer(log.VerifyingTracer().With("REDACTED", address))
	return p
}

func InstantiateForwardableAddress() (address string, netAddress *NetLocation) {
	for {
		var err error
		address = fmt.Sprintf("REDACTED",
			engineseed.Octets(20),
			engineseed.Int()%256,
			engineseed.Int()%256,
			engineseed.Int()%256,
			engineseed.Int()%256)
		netAddress, err = NewNetLocationString(address)
		if err != nil {
			panic(err)
		}
		if netAddress.Forwardable() {
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
func CreateLinkedRouters(cfg *settings.P2PSettings,
	n int,
	initRouter func(int, *Router) *Router,
	establish func([]*Router, int, int),
) []*Router {
	routers := make([]*Router, n)
	for i := 0; i < n; i++ {
		routers[i] = CreateRouter(cfg, i, initRouter)
	}

	if err := BeginRouters(routers); err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			establish(routers, i, j)
		}
	}

	return routers
}

//
//
//
func Connect2routers(routers []*Router, i, j int) {
	routerIDX := routers[i]
	routerJ := routers[j]

	c1, c2 := link.NetPipe()

	doneChan := make(chan struct{})
	go func() {
		err := routerIDX.appendNodeWithLinkage(c1)
		if err != nil {
			panic(err)
		}
		doneChan <- struct{}{}
	}()
	go func() {
		err := routerJ.appendNodeWithLinkage(c2)
		if err != nil {
			panic(err)
		}
		doneChan <- struct{}{}
	}()
	<-doneChan
	<-doneChan
}

func (sw *Router) appendNodeWithLinkage(link net.Conn) error {
	pc, err := verifyIncomingNodeLink(link, sw.settings, sw.memberKey.PrivateKey)
	if err != nil {
		if err := link.Close(); err != nil {
			sw.Tracer.Fault("REDACTED", "REDACTED", err)
		}
		return err
	}

	ni, err := greeting(link, time.Second, sw.memberDetails)
	if err != nil {
		if err := link.Close(); err != nil {
			sw.Tracer.Fault("REDACTED", "REDACTED", err)
		}
		return err
	}

	p := newNode(
		pc,
		MLinkSettings(sw.settings),
		ni,
		sw.handlersByChan,
		sw.messageKindByChanUID,
		sw.chanTraits,
		sw.HaltNodeForFault,
		sw.mlc,
	)

	if err = sw.appendNode(p); err != nil {
		pc.EndLink()
		return err
	}

	return nil
}

//
//
func BeginRouters(routers []*Router) error {
	for _, s := range routers {
		err := s.Begin() //
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateRouter(
	cfg *settings.P2PSettings,
	i int,
	initRouter func(int, *Router) *Router,
	opts ...RouterSetting,
) *Router {
	memberKey := MemberKey{
		PrivateKey: ed25519.GeneratePrivateKey(),
	}
	memberDetails := verifyMemberDetails(memberKey.ID(), fmt.Sprintf("REDACTED", i))
	address, err := NewNetLocationString(
		UIDLocationString(memberKey.ID(), memberDetails.(StandardMemberDetails).ObserveAddress),
	)
	if err != nil {
		panic(err)
	}

	t := NewMultiplexCarrier(memberDetails, memberKey, MLinkSettings(cfg))

	if err := t.Observe(*address); err != nil {
		panic(err)
	}

	//
	sw := initRouter(i, NewRouter(cfg, t, opts...))
	sw.AssignTracer(log.VerifyingTracer().With("REDACTED", i))
	sw.CollectionMemberKey(&memberKey)

	ni := memberDetails.(StandardMemberDetails)
	for ch := range sw.handlersByChan {
		ni.Streams = append(ni.Streams, ch)
	}
	memberDetails = ni

	//
	//
	t.memberDetails = memberDetails
	sw.CollectionMemberDetails(memberDetails)

	return sw
}

func verifyIncomingNodeLink(
	link net.Conn,
	settings *settings.P2PSettings,
	ourMemberPrivateKey vault.PrivateKey,
) (nodeLink, error) {
	return verifyNodeLink(link, settings, false, false, ourMemberPrivateKey, nil)
}

func verifyNodeLink(
	crudeLink net.Conn,
	cfg *settings.P2PSettings,
	outgoing, durable bool,
	ourMemberPrivateKey vault.PrivateKey,
	socketAddress *NetLocation,
) (pc nodeLink, err error) {
	link := crudeLink

	//
	if cfg.VerifyRandomize {
		//
		link = RandomizeLinkAfterFromSettings(link, 10*time.Second, cfg.VerifyRandomizeSettings)
	}

	//
	link, err = enhanceTokenLink(link, cfg.GreetingDeadline, ourMemberPrivateKey)
	if err != nil {
		return pc, fmt.Errorf("REDACTED", err)
	}

	//
	return newNodeLink(outgoing, durable, link, socketAddress), nil
}

//
//

func verifyMemberDetails(id ID, label string) MemberDetails {
	return verifyMemberDetailsWithFabric(id, label, "REDACTED")
}

func verifyMemberDetailsWithFabric(id ID, label, fabric string) MemberDetails {
	return StandardMemberDetails{
		ProtocolRelease: standardProtocolRelease,
		StandardMemberUID:   id,
		ObserveAddress:      fmt.Sprintf("REDACTED", fetchReleasePort()),
		Fabric:         fabric,
		Release:         "REDACTED",
		Streams:        []byte{verifyChan},
		Moniker:         label,
		Another: StandardMemberDetailsAnother{
			TransOrdinal:    "REDACTED",
			RPCLocation: fmt.Sprintf("REDACTED", fetchReleasePort()),
		},
	}
}

func fetchReleasePort() int {
	port, err := cometnet.FetchReleasePort()
	if err != nil {
		panic(err)
	}
	return port
}

type AddressRegistryEmulate struct {
	Locations        map[string]struct{}
	OurLocations     map[string]struct{}
	InternalLocations map[string]struct{}
}

var _ AddressLedger = (*AddressRegistryEmulate)(nil)

func (registry *AddressRegistryEmulate) AppendLocation(address *NetLocation, _ *NetLocation) error {
	registry.Locations[address.String()] = struct{}{}
	return nil
}
func (registry *AddressRegistryEmulate) AppendOurLocation(address *NetLocation) { registry.OurLocations[address.String()] = struct{}{} }
func (registry *AddressRegistryEmulate) OurLocation(address *NetLocation) bool {
	_, ok := registry.OurLocations[address.String()]
	return ok
}
func (registry *AddressRegistryEmulate) StampValid(ID) {}
func (registry *AddressRegistryEmulate) HasLocation(address *NetLocation) bool {
	_, ok := registry.Locations[address.String()]
	return ok
}

func (registry *AddressRegistryEmulate) DeleteLocation(address *NetLocation) {
	delete(registry.Locations, address.String())
}
func (registry *AddressRegistryEmulate) Persist() {}
func (registry *AddressRegistryEmulate) AppendInternalIDXDatastore(locations []string) {
	for _, address := range locations {
		registry.InternalLocations[address] = struct{}{}
	}
}
