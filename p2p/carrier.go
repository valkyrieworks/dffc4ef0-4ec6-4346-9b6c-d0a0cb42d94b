package p2p

import (
	"context"
	"fmt"
	"net"
	"time"

	"golang.org/x/net/netutil"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/utils/protoio"
	"github.com/valkyrieworks/p2p/link"
	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
)

const (
	standardCallDeadline      = time.Second
	standardRefineDeadline    = 5 * time.Second
	standardGreetingDeadline = 3 * time.Second
)

//
type IPDecoder interface {
	SearchIPAddress(context.Context, string) ([]net.IPAddr, error)
}

//
//
type allow struct {
	netAddress  *NetLocation
	link     net.Conn
	memberDetails MemberDetails
	err      error
}

//
//
//
//
//
//
type nodeSettings struct {
	chanTraits     []*link.StreamDefinition
	onNodeFault func(Node, any)
	outgoing    bool
	//
	//
	//
	isDurable  func(*NetLocation) bool
	handlersByChan  map[byte]Handler
	messageKindByChanUID map[byte]proto.Message
	stats       *Stats
	mlc           *statsTagRepository
}

//
//
//
type Carrier interface {
	//
	NetLocation() NetLocation

	//
	Allow(nodeSettings) (Node, error)

	//
	Call(NetLocation, nodeSettings) (Node, error)

	//
	Sanitize(Node)
}

//
//
type carrierActuality interface {
	End() error
	Observe(NetLocation) error
}

//
//
//
type LinkRefineFunction func(LinkCollection, net.Conn, []net.IP) error

//
//
func LinkReplicatedIPRefine() LinkRefineFunction {
	return func(cs LinkCollection, c net.Conn, ips []net.IP) error {
		for _, ip := range ips {
			if cs.HasIP(ip) {
				return ErrDeclined{
					link:        c,
					err:         fmt.Errorf("REDACTED", ip),
					isReplicated: true,
				}
			}
		}

		return nil
	}
}

//
//
type MulticastCarrierSetting func(*MulticastCarrier)

//
func MulticastCarrierLinkScreens(
	screens ...LinkRefineFunction,
) MulticastCarrierSetting {
	return func(mt *MulticastCarrier) { mt.linkScreens = screens }
}

//
//
func MulticastCarrierRefineDeadline(
	deadline time.Duration,
) MulticastCarrierSetting {
	return func(mt *MulticastCarrier) { mt.refineDeadline = deadline }
}

//
//
func MulticastCarrierDecoder(decoder IPDecoder) MulticastCarrierSetting {
	return func(mt *MulticastCarrier) { mt.decoder = decoder }
}

//
//
func MulticastCarrierMaximumIncomingLinkages(n int) MulticastCarrierSetting {
	return func(mt *MulticastCarrier) { mt.maximumIncomingLinkages = n }
}

//
//
type MulticastCarrier struct {
	netAddress                NetLocation
	observer               net.Listener
	maximumIncomingLinkages int //

	acceptc chan allow
	closec  chan struct{}

	//
	links       LinkCollection
	linkScreens []LinkRefineFunction

	callDeadline      time.Duration
	refineDeadline    time.Duration
	greetingDeadline time.Duration
	memberDetails         MemberDetails
	memberKey          MemberKey
	decoder         IPDecoder

	//
	//
	//
	mSettings link.MLinkSettings
}

//
var (
	_ Carrier          = (*MulticastCarrier)(nil)
	_ carrierActuality = (*MulticastCarrier)(nil)
)

//
func NewMulticastCarrier(
	memberDetails MemberDetails,
	memberKey MemberKey,
	mSettings link.MLinkSettings,
) *MulticastCarrier {
	return &MulticastCarrier{
		acceptc:          make(chan allow),
		closec:           make(chan struct{}),
		callDeadline:      standardCallDeadline,
		refineDeadline:    standardRefineDeadline,
		greetingDeadline: standardGreetingDeadline,
		mSettings:          mSettings,
		memberDetails:         memberDetails,
		memberKey:          memberKey,
		links:            NewLinkCollection(),
		decoder:         net.DefaultResolver,
	}
}

//
func (mt *MulticastCarrier) NetLocation() NetLocation {
	return mt.netAddress
}

//
func (mt *MulticastCarrier) Allow(cfg nodeSettings) (Node, error) {
	select {
	//
	//
	case a := <-mt.acceptc:
		if a.err != nil {
			return nil, a.err
		}

		cfg.outgoing = false

		return mt.encloseNode(a.link, a.memberDetails, cfg, a.netAddress), nil
	case <-mt.closec:
		return nil, ErrCarrierHalted{}
	}
}

//
func (mt *MulticastCarrier) Call(
	address NetLocation,
	cfg nodeSettings,
) (Node, error) {
	c, err := address.CallDeadline(mt.callDeadline)
	if err != nil {
		return nil, err
	}

	if mt.mSettings.VerifyRandomize {
		//
		c = RandomizeLinkAfterFromSettings(c, 10*time.Second, mt.mSettings.VerifyRandomizeSettings)
	}

	//
	if err := mt.refineLink(c); err != nil {
		return nil, err
	}

	tokenLink, memberDetails, err := mt.enhance(c, &address)
	if err != nil {
		return nil, err
	}

	cfg.outgoing = true

	p := mt.encloseNode(tokenLink, memberDetails, cfg, &address)

	return p, nil
}

//
func (mt *MulticastCarrier) End() error {
	close(mt.closec)

	if mt.observer != nil {
		return mt.observer.Close()
	}

	return nil
}

//
func (mt *MulticastCarrier) Observe(address NetLocation) error {
	ln, err := net.Listen("REDACTED", address.CallString())
	if err != nil {
		return err
	}

	if mt.maximumIncomingLinkages > 0 {
		ln = netutil.LimitListener(ln, mt.maximumIncomingLinkages)
	}

	mt.netAddress = address
	mt.observer = ln

	go mt.allowNodes()

	return nil
}

//
//
//
//
func (mt *MulticastCarrier) AppendConduit(chanUID byte) {
	if ni, ok := mt.memberDetails.(StandardMemberDetails); ok {
		if !ni.HasConduit(chanUID) {
			ni.Streams = append(ni.Streams, chanUID)
		}
		mt.memberDetails = ni
	}
}

func (mt *MulticastCarrier) allowNodes() {
	for {
		c, err := mt.observer.Accept()
		if err != nil {
			//
			select {
			case _, ok := <-mt.closec:
				if !ok {
					return
				}
			default:
				//
			}

			mt.acceptc <- allow{err: err}
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
					err := ErrDeclined{
						link:          c,
						err:           fmt.Errorf("REDACTED", r),
						isAuthBreakdown: true,
					}
					select {
					case mt.acceptc <- allow{err: err}:
					case <-mt.closec:
						//
						_ = c.Close()
						return
					}
				}
			}()

			var (
				memberDetails   MemberDetails
				tokenLink *link.TokenLinkage
				netAddress    *NetLocation
			)

			err := mt.refineLink(c)
			if err == nil {
				tokenLink, memberDetails, err = mt.enhance(c, nil)
				if err == nil {
					address := c.RemoteAddr()
					id := PublicKeyToUID(tokenLink.DistantPublicKey())
					netAddress = NewNetLocation(id, address)
				}
			}

			select {
			case mt.acceptc <- allow{netAddress, tokenLink, memberDetails, err}:
				//
			case <-mt.closec:
				//
				_ = c.Close()
				return
			}
		}(c)
	}
}

//
//
func (mt *MulticastCarrier) Sanitize(p Node) {
	mt.links.DeleteAddress(p.DistantAddress())
	_ = p.EndLink()
}

func (mt *MulticastCarrier) sanitize(c net.Conn) error {
	mt.links.Delete(c)

	return c.Close()
}

func (mt *MulticastCarrier) refineLink(c net.Conn) (err error) {
	defer func() {
		if err != nil {
			_ = c.Close()
		}
	}()

	//
	if mt.links.Has(c) {
		return ErrDeclined{link: c, isReplicated: true}
	}

	//
	ips, err := decipherIDXPs(mt.decoder, c)
	if err != nil {
		return err
	}

	faultc := make(chan error, len(mt.linkScreens))

	for _, f := range mt.linkScreens {
		go func(f LinkRefineFunction, c net.Conn, ips []net.IP, faultc chan<- error) {
			faultc <- f(mt.links, c, ips)
		}(f, c, ips, faultc)
	}

	for i := 0; i < cap(faultc); i++ {
		select {
		case err := <-faultc:
			if err != nil {
				return ErrDeclined{link: c, err: err, isScreened: true}
			}
		case <-time.After(mt.refineDeadline):
			return ErrRefineDeadline{}
		}
	}

	mt.links.Set(c, ips)

	return nil
}

func (mt *MulticastCarrier) enhance(
	c net.Conn,
	calledAddress *NetLocation,
) (tokenLink *link.TokenLinkage, memberDetails MemberDetails, err error) {
	defer func() {
		if err != nil {
			_ = mt.sanitize(c)
		}
	}()

	tokenLink, err = enhanceTokenLink(c, mt.greetingDeadline, mt.memberKey.PrivateKey)
	if err != nil {
		return nil, nil, ErrDeclined{
			link:          c,
			err:           fmt.Errorf("REDACTED", err),
			isAuthBreakdown: true,
		}
	}

	//
	linkUID := PublicKeyToUID(tokenLink.DistantPublicKey())
	if calledAddress != nil {
		if calledUID := calledAddress.ID; linkUID != calledUID {
			return nil, nil, ErrDeclined{
				link: c,
				id:   linkUID,
				err: fmt.Errorf(
					"REDACTED",
					linkUID,
					calledUID,
				),
				isAuthBreakdown: true,
			}
		}
	}

	memberDetails, err = greeting(tokenLink, mt.greetingDeadline, mt.memberDetails)
	if err != nil {
		return nil, nil, ErrDeclined{
			link:          c,
			err:           fmt.Errorf("REDACTED", err),
			isAuthBreakdown: true,
		}
	}

	if err := memberDetails.Certify(); err != nil {
		return nil, nil, ErrDeclined{
			link:              c,
			err:               err,
			isMemberDetailsCorrupt: true,
		}
	}

	//
	if linkUID != memberDetails.ID() {
		return nil, nil, ErrDeclined{
			link: c,
			id:   linkUID,
			err: fmt.Errorf(
				"REDACTED",
				linkUID,
				memberDetails.ID(),
			),
			isAuthBreakdown: true,
		}
	}

	//
	if mt.memberDetails.ID() == memberDetails.ID() {
		return nil, nil, ErrDeclined{
			address:   *NewNetLocation(memberDetails.ID(), c.RemoteAddr()),
			link:   c,
			id:     memberDetails.ID(),
			isEgo: true,
		}
	}

	if err := mt.memberDetails.HarmoniousWith(memberDetails); err != nil {
		return nil, nil, ErrDeclined{
			link:           c,
			err:            err,
			id:             memberDetails.ID(),
			isDiscordant: true,
		}
	}

	return tokenLink, memberDetails, nil
}

func (mt *MulticastCarrier) encloseNode(
	c net.Conn,
	ni MemberDetails,
	cfg nodeSettings,
	socketAddress *NetLocation,
) Node {
	durable := false
	if cfg.isDurable != nil {
		if cfg.outgoing {
			durable = cfg.isDurable(socketAddress)
		} else {
			egoRegisteredAddress, err := ni.NetLocation()
			if err == nil {
				durable = cfg.isDurable(egoRegisteredAddress)
			}
		}
	}

	nodeLink := newNodeLink(
		cfg.outgoing,
		durable,
		c,
		socketAddress,
	)

	p := newNode(
		nodeLink,
		mt.mSettings,
		ni,
		cfg.handlersByChan,
		cfg.messageKindByChanUID,
		cfg.chanTraits,
		cfg.onNodeFault,
		cfg.mlc,
		NodeStats(cfg.stats),
	)

	return p
}

func greeting(
	c net.Conn,
	deadline time.Duration,
	memberDetails MemberDetails,
) (MemberDetails, error) {
	if err := c.SetDeadline(time.Now().Add(deadline)); err != nil {
		return nil, err
	}

	var (
		faultc = make(chan error, 2)

		pbpeerMemberDetails tmp2p.StandardMemberDetails
		nodeMemberDetails   StandardMemberDetails
		ourMemberDetails    = memberDetails.(StandardMemberDetails)
	)

	go func(faultc chan<- error, c net.Conn) {
		_, err := protoio.NewSeparatedRecorder(c).RecordMessage(ourMemberDetails.ToSchema())
		faultc <- err
	}(faultc, c)
	go func(faultc chan<- error, c net.Conn) {
		schemaScanner := protoio.NewSeparatedScanner(c, MaximumMemberDetailsVolume())
		_, err := schemaScanner.ScanMessage(&pbpeerMemberDetails)
		faultc <- err
	}(faultc, c)

	for i := 0; i < cap(faultc); i++ {
		err := <-faultc
		if err != nil {
			return nil, err
		}
	}

	nodeMemberDetails, err := StandardMemberDetailsFromToSchema(&pbpeerMemberDetails)
	if err != nil {
		return nil, err
	}

	return nodeMemberDetails, c.SetDeadline(time.Time{})
}

func enhanceTokenLink(
	c net.Conn,
	deadline time.Duration,
	privateKey vault.PrivateKey,
) (*link.TokenLinkage, error) {
	if err := c.SetDeadline(time.Now().Add(deadline)); err != nil {
		return nil, err
	}

	sc, err := link.CreateTokenLinkage(c, privateKey)
	if err != nil {
		return nil, err
	}

	return sc, sc.CollectionLimit(time.Time{})
}

func decipherIDXPs(decoder IPDecoder, c net.Conn) ([]net.IP, error) {
	machine, _, err := net.SplitHostPort(c.RemoteAddr().String())
	if err != nil {
		return nil, err
	}

	locations, err := decoder.SearchIPAddress(context.Background(), machine)
	if err != nil {
		return nil, err
	}

	ips := []net.IP{}

	for _, address := range locations {
		ips = append(ips, address.IP)
	}

	return ips, nil
}
