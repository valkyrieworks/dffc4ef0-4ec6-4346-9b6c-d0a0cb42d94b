package p2p

import (
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/utils/cmap"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"

	cmtconn "github.com/valkyrieworks/p2p/link"
)

//

const statsTimerPeriod = 10 * time.Second

//
type Node interface {
	daemon.Daemon
	PurgeHalt()

	ID() ID               //
	DistantIP() net.IP     //
	DistantAddress() net.Addr //

	IsOutgoing() bool   //
	IsDurable() bool //

	EndLink() error //

	MemberDetails() MemberDetails //
	Status() cmtconn.LinkageState
	SocketAddress() *NetLocation //

	Transmit(Packet) bool
	AttemptTransmit(Packet) bool

	Set(string, any)
	Get(string) any

	CollectionDeletionErrored()
	FetchDeletionErrored() bool
}

//

//
type nodeLink struct {
	outgoing   bool
	durable bool
	link       net.Conn //

	socketAddress *NetLocation

	//
	ip net.IP
}

func newNodeLink(
	outgoing, durable bool,
	link net.Conn,
	socketAddress *NetLocation,
) nodeLink {
	return nodeLink{
		outgoing:   outgoing,
		durable: durable,
		link:       link,
		socketAddress: socketAddress,
	}
}

//
//
func (pc nodeLink) ID() ID {
	return PublicKeyToUID(pc.link.(*cmtconn.TokenLinkage).DistantPublicKey())
}

//
func (pc nodeLink) DistantIP() net.IP {
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
	daemon.RootDaemon

	//
	nodeLink
	mconn *cmtconn.MLinkage

	//
	//
	//
	memberDetails MemberDetails
	streams []byte

	//
	Data *cmap.CIndex

	stats *Stats
	mlc     *statsTagRepository

	//
	deletionEndeavorErrored bool
}

type NodeSetting func(*node)

func newNode(
	pc nodeLink,
	mSettings cmtconn.MLinkSettings,
	memberDetails MemberDetails,
	handlersByChan map[byte]Handler,
	messageKindByChanUID map[byte]proto.Message,
	chanTraits []*cmtconn.StreamDefinition,
	onNodeFault func(Node, any),
	mlc *statsTagRepository,
	options ...NodeSetting,
) *node {
	p := &node{
		nodeLink: pc,
		memberDetails: memberDetails,
		streams: memberDetails.(StandardMemberDetails).Streams,
		Data:     cmap.NewCIndex(),
		stats:  NoopStats(),
		mlc:      mlc,
	}

	p.mconn = instantiateMLinkage(
		pc.link,
		p,
		handlersByChan,
		messageKindByChanUID,
		chanTraits,
		onNodeFault,
		mSettings,
	)
	p.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", p)
	for _, setting := range options {
		setting(p)
	}

	return p
}

//
func (p *node) String() string {
	if p.outgoing {
		return fmt.Sprintf("REDACTED", p.mconn, p.ID())
	}

	return fmt.Sprintf("REDACTED", p.mconn, p.ID())
}

//
//

//
func (p *node) AssignTracer(l log.Tracer) {
	p.Tracer = l
	p.mconn.AssignTracer(l)
}

//
func (p *node) OnBegin() error {
	if err := p.RootDaemon.OnBegin(); err != nil {
		return err
	}

	if err := p.mconn.Begin(); err != nil {
		return err
	}

	go p.statsMonitor()
	return nil
}

//
//
//
//
func (p *node) PurgeHalt() {
	p.mconn.PurgeHalt() //
}

//
func (p *node) OnHalt() {
	if err := p.mconn.Halt(); err != nil { //
		p.Tracer.Diagnose("REDACTED", "REDACTED", err)
	}
}

//
//

//
func (p *node) ID() ID {
	return p.memberDetails.ID()
}

//
func (p *node) IsOutgoing() bool {
	return p.outgoing
}

//
func (p *node) IsDurable() bool {
	return p.durable
}

//
func (p *node) MemberDetails() MemberDetails {
	return p.memberDetails
}

//
//
//
//
func (p *node) SocketAddress() *NetLocation {
	return p.socketAddress
}

//
func (p *node) Status() cmtconn.LinkageState {
	return p.mconn.Status()
}

//
//
//
//
func (p *node) Transmit(e Packet) bool {
	return p.transmit(e.StreamUID, e.Signal, p.mconn.Transmit)
}

//
//
//
//
func (p *node) AttemptTransmit(e Packet) bool {
	return p.transmit(e.StreamUID, e.Signal, p.mconn.AttemptTransmit)
}

func (p *node) transmit(chanUID byte, msg proto.Message, transmitFunction func(byte, []byte) bool) bool {
	if !p.IsActive() {
		return false
	} else if !p.hasConduit(chanUID) {
		return false
	}
	indicatorTagItem := p.mlc.ItemToIndicatorTag(msg)
	if w, ok := msg.(Adapter); ok {
		msg = w.Enclose()
	}
	messageOctets, err := proto.Marshal(msg)
	if err != nil {
		p.Tracer.Fault("REDACTED", "REDACTED", err)
		return false
	}
	res := transmitFunction(chanUID, messageOctets)
	if res {
		tags := []string{
			"REDACTED", string(p.ID()),
			"REDACTED", fmt.Sprintf("REDACTED", chanUID),
		}
		p.stats.NodeTransmitOctetsSum.With(tags...).Add(float64(len(messageOctets)))
		p.stats.SignalTransmitOctetsSum.With("REDACTED", indicatorTagItem).Add(float64(len(messageOctets)))
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
func (p *node) hasConduit(chanUID byte) bool {
	for _, ch := range p.streams {
		if ch == chanUID {
			return true
		}
	}
	return false
}

//
func (p *node) EndLink() error {
	return p.link.Close()
}

func (p *node) CollectionDeletionErrored() {
	p.deletionEndeavorErrored = true
}

func (p *node) FetchDeletionErrored() bool {
	return p.deletionEndeavorErrored
}

//
//
//

//
func (pc *nodeLink) EndLink() {
	pc.link.Close()
}

//
func (p *node) DistantAddress() net.Addr {
	return p.link.RemoteAddr()
}

//
func (p *node) MayTransmit(chanUID byte) bool {
	if !p.IsActive() {
		return false
	}
	return p.mconn.MayTransmit(chanUID)
}

//

func NodeStats(stats *Stats) NodeSetting {
	return func(p *node) {
		p.stats = stats
	}
}

func (p *node) statsMonitor() {
	statsTimer := time.NewTicker(statsTimerPeriod)
	defer statsTimer.Stop()

	for {
		select {
		case <-statsTimer.C:
			state := p.mconn.Status()
			var transmitBufferVolume float64
			for _, chanState := range state.Streams {
				transmitBufferVolume += float64(chanState.TransmitBufferVolume)
			}

			p.stats.NodeAwaitingTransmitOctets.With("REDACTED", string(p.ID())).Set(transmitBufferVolume)
		case <-p.Exit():
			return
		}
	}
}

//
//

func instantiateMLinkage(
	link net.Conn,
	p *node,
	handlersByChan map[byte]Handler,
	messageKindByChanUID map[byte]proto.Message,
	chanTraits []*cmtconn.StreamDefinition,
	onNodeFault func(Node, any),
	settings cmtconn.MLinkSettings,
) *cmtconn.MLinkage {
	onAccept := func(chanUID byte, messageOctets []byte) {
		handler := handlersByChan[chanUID]
		if handler == nil {
			//
			//
			panic(fmt.Sprintf("REDACTED", chanUID))
		}
		mt := messageKindByChanUID[chanUID]
		msg := proto.Clone(mt)
		err := proto.Unmarshal(messageOctets, msg)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err, reflect.TypeOf(mt)))
		}
		tags := []string{
			"REDACTED", string(p.ID()),
			"REDACTED", fmt.Sprintf("REDACTED", chanUID),
		}
		if w, ok := msg.(Extractor); ok {
			msg, err = w.Disclose()
			if err != nil {
				panic(fmt.Errorf("REDACTED", err))
			}
		}
		p.stats.NodeAcceptOctetsSum.With(tags...).Add(float64(len(messageOctets)))
		p.stats.SignalAcceptOctetsSum.With("REDACTED", p.mlc.ItemToIndicatorTag(msg)).Add(float64(len(messageOctets)))
		handler.Accept(Packet{
			StreamUID: chanUID,
			Src:       p,
			Signal:   msg,
		})
	}

	onFault := func(r any) {
		onNodeFault(p, r)
	}

	return cmtconn.NewMLinkageWithSettings(
		link,
		chanTraits,
		onAccept,
		onFault,
		settings,
	)
}
