package p2p

import (
	"context"
	"fmt"
	"net"
	"time"

	"golang.org/x/net/netutil"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"
	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
)

const (
	fallbackCallDeadline      = time.Second
	fallbackRefineDeadline    = 5 * time.Second
	fallbackNegotiationDeadline = 3 * time.Second
)

//
type INETDecrypter interface {
	SearchINETLocation(context.Context, string) ([]net.IPAddr, error)
}

//
//
type embrace struct {
	networkLocation  *NetworkLocator
	link     net.Conn
	peerDetails PeerDetails
	err      error
}

//
//
//
//
//
//
type nodeSettings struct {
	chnlDescriptions     []*link.ConduitDefinition
	uponNodeFailure func(Node, any)
	outgoing    bool
	//
	//
	//
	equalsEnduring  func(*NetworkLocator) bool
	enginesViaChnl  map[byte]Handler
	signalKindViaChnlUUID map[byte]proto.Message
	telemetry       *Telemetry
	mlc           *telemetryTagStash
}

//
//
//
type Carrier interface {
	//
	NetworkLocator() NetworkLocator

	//
	Embrace(nodeSettings) (Node, error)

	//
	Call(NetworkLocator, nodeSettings) (Node, error)

	//
	Sanitize(Node)
}

//
//
type carrierDuration interface {
	Shutdown() error
	Overhear(NetworkLocator) error
}

//
//
//
type LinkRefineMethod func(LinkAssign, net.Conn, []net.IP) error

//
//
func LinkReplicatedINETRefine() LinkRefineMethod {
	return func(cs LinkAssign, c net.Conn, ips []net.IP) error {
		for _, ip := range ips {
			if cs.OwnsINET(ip) {
				return FaultDeclined{
					link:        c,
					err:         fmt.Errorf("REDACTED", ip),
					equalsReplicated: true,
				}
			}
		}

		return nil
	}
}

//
//
type MultiplexCarrierSelection func(*MultiplexCarrier)

//
func MultiplexCarrierLinkCriteria(
	criteria ...LinkRefineMethod,
) MultiplexCarrierSelection {
	return func(mt *MultiplexCarrier) { mt.linkCriteria = criteria }
}

//
//
func MultiplexCarrierRefineDeadline(
	deadline time.Duration,
) MultiplexCarrierSelection {
	return func(mt *MultiplexCarrier) { mt.refineDeadline = deadline }
}

//
//
func MultiplexCarrierDecrypter(decrypter INETDecrypter) MultiplexCarrierSelection {
	return func(mt *MultiplexCarrier) { mt.decrypter = decrypter }
}

//
//
func MultiplexCarrierMaximumArrivingLinkages(n int) MultiplexCarrierSelection {
	return func(mt *MultiplexCarrier) { mt.maximumArrivingLinkages = n }
}

//
//
type MultiplexCarrier struct {
	networkLocation                NetworkLocator
	observer               net.Listener
	maximumArrivingLinkages int //

	approvechnl chan embrace
	terminatechnl  chan struct{}

	//
	links       LinkAssign
	linkCriteria []LinkRefineMethod

	callDeadline      time.Duration
	refineDeadline    time.Duration
	negotiationDeadline time.Duration
	peerDetails         PeerDetails
	peerToken          PeerToken
	decrypter         INETDecrypter

	//
	//
	//
	moduleSettings link.ModuleLinkSettings
}

//
var (
	_ Carrier          = (*MultiplexCarrier)(nil)
	_ carrierDuration = (*MultiplexCarrier)(nil)
)

//
func FreshMultiplexCarrier(
	peerDetails PeerDetails,
	peerToken PeerToken,
	moduleSettings link.ModuleLinkSettings,
) *MultiplexCarrier {
	return &MultiplexCarrier{
		approvechnl:          make(chan embrace),
		terminatechnl:           make(chan struct{}),
		callDeadline:      fallbackCallDeadline,
		refineDeadline:    fallbackRefineDeadline,
		negotiationDeadline: fallbackNegotiationDeadline,
		moduleSettings:          moduleSettings,
		peerDetails:         peerDetails,
		peerToken:          peerToken,
		links:            FreshLinkAssign(),
		decrypter:         net.DefaultResolver,
	}
}

//
func (mt *MultiplexCarrier) NetworkLocator() NetworkLocator {
	return mt.networkLocation
}

//
func (mt *MultiplexCarrier) Embrace(cfg nodeSettings) (Node, error) {
	select {
	//
	//
	case a := <-mt.approvechnl:
		if a.err != nil {
			return nil, a.err
		}

		cfg.outgoing = false

		return mt.encloseNode(a.link, a.peerDetails, cfg, a.networkLocation), nil
	case <-mt.terminatechnl:
		return nil, FaultCarrierTerminated{}
	}
}

//
func (mt *MultiplexCarrier) Call(
	location NetworkLocator,
	cfg nodeSettings,
) (Node, error) {
	c, err := location.CallDeadline(mt.callDeadline)
	if err != nil {
		return nil, err
	}

	if mt.moduleSettings.VerifyRandomize {
		//
		c = RandomizeLinkSubsequentOriginatingSettings(c, 10*time.Second, mt.moduleSettings.VerifyRandomizeSettings)
	}

	//
	if err := mt.refineLink(c); err != nil {
		return nil, err
	}

	credentialLink, peerDetails, err := mt.modernize(c, &location)
	if err != nil {
		return nil, err
	}

	cfg.outgoing = true

	p := mt.encloseNode(credentialLink, peerDetails, cfg, &location)

	return p, nil
}

//
func (mt *MultiplexCarrier) Shutdown() error {
	close(mt.terminatechnl)

	if mt.observer != nil {
		return mt.observer.Close()
	}

	return nil
}

//
func (mt *MultiplexCarrier) Overhear(location NetworkLocator) error {
	ln, err := net.Listen("REDACTED", location.CallText())
	if err != nil {
		return err
	}

	if mt.maximumArrivingLinkages > 0 {
		ln = netutil.LimitListener(ln, mt.maximumArrivingLinkages)
	}

	mt.networkLocation = location
	mt.observer = ln

	go mt.embraceNodes()

	return nil
}

//
//
//
//
func (mt *MultiplexCarrier) AppendConduit(chnlUUID byte) {
	if ni, ok := mt.peerDetails.(FallbackPeerDetails); ok {
		if !ni.OwnsConduit(chnlUUID) {
			ni.Conduits = append(ni.Conduits, chnlUUID)
		}
		mt.peerDetails = ni
	}
}

func (mt *MultiplexCarrier) embraceNodes() {
	for {
		c, err := mt.observer.Accept()
		if err != nil {
			//
			select {
			case _, ok := <-mt.terminatechnl:
				if !ok {
					return
				}
			default:
				//
			}

			mt.approvechnl <- embrace{err: err}
			return
		}

		//
		//
		//
		//
		//
		go func(c net.Conn) {
			defer func() {
				if r := recover(); r != nil {
					err := FaultDeclined{
						link:          c,
						err:           fmt.Errorf("REDACTED", r),
						equalsAuthBreakdown: true,
					}
					select {
					case mt.approvechnl <- embrace{err: err}:
					case <-mt.terminatechnl:
						//
						_ = c.Close()
						return
					}
				}
			}()

			var (
				peerDetails   PeerDetails
				credentialLink *link.CredentialLinkage
				networkLocation    *NetworkLocator
			)

			err := mt.refineLink(c)
			if err == nil {
				credentialLink, peerDetails, err = mt.modernize(c, nil)
				if err == nil {
					location := c.RemoteAddr()
					id := PublicTokenTowardUUID(credentialLink.DistantPublicToken())
					networkLocation = FreshNetworkLocator(id, location)
				}
			}

			select {
			case mt.approvechnl <- embrace{networkLocation, credentialLink, peerDetails, err}:
				//
			case <-mt.terminatechnl:
				//
				_ = c.Close()
				return
			}
		}(c)
	}
}

//
//
func (mt *MultiplexCarrier) Sanitize(p Node) {
	mt.links.DiscardLocation(p.DistantLocation())
	_ = p.ShutdownLink()
}

func (mt *MultiplexCarrier) sanitize(c net.Conn) error {
	mt.links.Discard(c)

	return c.Close()
}

func (mt *MultiplexCarrier) refineLink(c net.Conn) (err error) {
	defer func() {
		if err != nil {
			_ = c.Close()
		}
	}()

	//
	if mt.links.Has(c) {
		return FaultDeclined{link: c, equalsReplicated: true}
	}

	//
	ips, err := decipherIDXProcesses(mt.decrypter, c)
	if err != nil {
		return err
	}

	faultchnl := make(chan error, len(mt.linkCriteria))

	for _, f := range mt.linkCriteria {
		go func(f LinkRefineMethod, c net.Conn, ips []net.IP, faultchnl chan<- error) {
			faultchnl <- f(mt.links, c, ips)
		}(f, c, ips, faultchnl)
	}

	for i := 0; i < cap(faultchnl); i++ {
		select {
		case err := <-faultchnl:
			if err != nil {
				return FaultDeclined{link: c, err: err, equalsScreened: true}
			}
		case <-time.After(mt.refineDeadline):
			return FaultRefineDeadline{}
		}
	}

	mt.links.Set(c, ips)

	return nil
}

func (mt *MultiplexCarrier) modernize(
	c net.Conn,
	calledLocation *NetworkLocator,
) (credentialLink *link.CredentialLinkage, peerDetails PeerDetails, err error) {
	defer func() {
		if err != nil {
			_ = mt.sanitize(c)
		}
	}()

	credentialLink, err = modernizeCredentialLink(c, mt.negotiationDeadline, mt.peerToken.PrivateToken)
	if err != nil {
		return nil, nil, FaultDeclined{
			link:          c,
			err:           fmt.Errorf("REDACTED", err),
			equalsAuthBreakdown: true,
		}
	}

	//
	linkUUID := PublicTokenTowardUUID(credentialLink.DistantPublicToken())
	if calledLocation != nil {
		if calledUUID := calledLocation.ID; linkUUID != calledUUID {
			return nil, nil, FaultDeclined{
				link: c,
				id:   linkUUID,
				err: fmt.Errorf(
					"REDACTED",
					linkUUID,
					calledUUID,
				),
				equalsAuthBreakdown: true,
			}
		}
	}

	peerDetails, err = negotiation(credentialLink, mt.negotiationDeadline, mt.peerDetails)
	if err != nil {
		return nil, nil, FaultDeclined{
			link:          c,
			err:           fmt.Errorf("REDACTED", err),
			equalsAuthBreakdown: true,
		}
	}

	if err := peerDetails.Certify(); err != nil {
		return nil, nil, FaultDeclined{
			link:              c,
			err:               err,
			equalsPeerDetailsUnfit: true,
		}
	}

	//
	if linkUUID != peerDetails.ID() {
		return nil, nil, FaultDeclined{
			link: c,
			id:   linkUUID,
			err: fmt.Errorf(
				"REDACTED",
				linkUUID,
				peerDetails.ID(),
			),
			equalsAuthBreakdown: true,
		}
	}

	//
	if mt.peerDetails.ID() == peerDetails.ID() {
		return nil, nil, FaultDeclined{
			location:   *FreshNetworkLocator(peerDetails.ID(), c.RemoteAddr()),
			link:   c,
			id:     peerDetails.ID(),
			equalsEgo: true,
		}
	}

	if err := mt.peerDetails.MatchedUsing(peerDetails); err != nil {
		return nil, nil, FaultDeclined{
			link:           c,
			err:            err,
			id:             peerDetails.ID(),
			equalsUnmatched: true,
		}
	}

	return credentialLink, peerDetails, nil
}

func (mt *MultiplexCarrier) encloseNode(
	c net.Conn,
	ni PeerDetails,
	cfg nodeSettings,
	portLocation *NetworkLocator,
) Node {
	enduring := false
	if cfg.equalsEnduring != nil {
		if cfg.outgoing {
			enduring = cfg.equalsEnduring(portLocation)
		} else {
			egoIndicatedLocation, err := ni.NetworkLocator()
			if err == nil {
				enduring = cfg.equalsEnduring(egoIndicatedLocation)
			}
		}
	}

	nodeLink := freshNodeLink(
		cfg.outgoing,
		enduring,
		c,
		portLocation,
	)

	p := freshNode(
		nodeLink,
		mt.moduleSettings,
		ni,
		cfg.enginesViaChnl,
		cfg.signalKindViaChnlUUID,
		cfg.chnlDescriptions,
		cfg.uponNodeFailure,
		cfg.mlc,
		NodeTelemetry(cfg.telemetry),
	)

	return p
}

func negotiation(
	c net.Conn,
	deadline time.Duration,
	peerDetails PeerDetails,
) (PeerDetails, error) {
	if err := c.SetDeadline(time.Now().Add(deadline)); err != nil {
		return nil, err
	}

	var (
		faultchnl = make(chan error, 2)

		fabricpeerPeerDetails tmpfabric.FallbackPeerDetails
		nodePeerDetails   FallbackPeerDetails
		minePeerDetails    = peerDetails.(FallbackPeerDetails)
	)

	go func(faultchnl chan<- error, c net.Conn) {
		_, err := protocolio.FreshSeparatedPersistor(c).PersistSignal(minePeerDetails.TowardSchema())
		faultchnl <- err
	}(faultchnl, c)
	go func(faultchnl chan<- error, c net.Conn) {
		schemaFetcher := protocolio.FreshSeparatedFetcher(c, MaximumPeerDetailsExtent())
		_, err := schemaFetcher.FetchSignal(&fabricpeerPeerDetails)
		faultchnl <- err
	}(faultchnl, c)

	for i := 0; i < cap(faultchnl); i++ {
		err := <-faultchnl
		if err != nil {
			return nil, err
		}
	}

	nodePeerDetails, err := FallbackPeerDetailsOriginatingTowardSchema(&fabricpeerPeerDetails)
	if err != nil {
		return nil, err
	}

	return nodePeerDetails, c.SetDeadline(time.Time{})
}

func modernizeCredentialLink(
	c net.Conn,
	deadline time.Duration,
	privateToken security.PrivateToken,
) (*link.CredentialLinkage, error) {
	if err := c.SetDeadline(time.Now().Add(deadline)); err != nil {
		return nil, err
	}

	sc, err := link.CreateCredentialLinkage(c, privateToken)
	if err != nil {
		return nil, err
	}

	return sc, sc.AssignExpiration(time.Time{})
}

func decipherIDXProcesses(decrypter INETDecrypter, c net.Conn) ([]net.IP, error) {
	machine, _, err := net.SplitHostPort(c.RemoteAddr().String())
	if err != nil {
		return nil, err
	}

	locations, err := decrypter.SearchINETLocation(context.Background(), machine)
	if err != nil {
		return nil, err
	}

	ips := []net.IP{}

	for _, location := range locations {
		ips = append(ips, location.IP)
	}

	return ips, nil
}
