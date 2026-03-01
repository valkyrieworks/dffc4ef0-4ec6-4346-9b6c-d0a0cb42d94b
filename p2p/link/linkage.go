package link

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"reflect"
	"runtime/debug"
	"sync/atomic"
	"time"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	stream "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/throughput"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/clock"
	tmpfabric "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/p2p"
)

const (
	fallbackMaximumPacketSignalWorkloadExtent = 1024

	countClusterPacketArtifacts = 10
	minimumFetchReserveExtent  = 1024
	minimumPersistReserveExtent = 65536
	reviseMetrics        = 2 * time.Second

	//
	//
	//
	fallbackPurgeRegulate = 10 * time.Millisecond

	fallbackTransmitStagingVolume   = 1
	fallbackObtainReserveVolume  = 4096
	fallbackObtainArtifactVolume = 22020096      //
	fallbackTransmitFrequency            = int64(512000) //
	fallbackObtainFrequency            = int64(512000) //
	fallbackTransmitDeadline         = 10 * time.Second
	fallbackPingDuration        = 60 * time.Second
	fallbackPongDeadline         = 45 * time.Second
)

type (
	acceptClbkMethod func(chnlUUID byte, signalOctets []byte)
	failureClbkMethod   func(any)
)

/**
.

f
.

n
.
n
.

:

}
}

s
e
.

e
.

.
*/
type ModuleLinkage struct {
	facility.FoundationFacility

	link          net.Conn
	areaLinkFetcher *bufio.Reader
	areaLinkPersistor *bufio.Writer
	transmitOverseer   *stream.Overseer
	obtainOverseer   *stream.Overseer
	transmit          chan struct{}
	pong          chan struct{}
	conduits      []*Conduit
	conduitsOffset   map[byte]*Conduit
	uponAccept     acceptClbkMethod
	uponFailure       failureClbkMethod
	faulted       uint32
	settings        ModuleLinkSettings

	//
	//
	exitTransmitProcedure chan struct{}
	completeTransmitProcedure chan struct{}

	//
	exitObtainProcedure chan struct{}

	//
	//
	haltMutex commitchronize.Exclusion

	purgeClock *clock.RegulateClock //
	pingClock  *time.Ticker         //

	//
	pongClock     *time.Timer
	pongDeadlineChnl chan bool //

	chnlMetricsClock *time.Ticker //

	spawned time.Time //

	_maxartifactsize int
}

//
type ModuleLinkSettings struct {
	TransmitFrequency int64 `mapstructure:"transmit_frequency"`
	ObtainFrequency int64 `mapstructure:"obtain_frequency"`

	//
	MaximumPacketSignalWorkloadExtent int `mapstructure:"maximum_packet_signal_workload_extent"`

	//
	PurgeRegulate time.Duration `mapstructure:"purge_regulate"`

	//
	PingDuration time.Duration `mapstructure:"ping_duration"`

	//
	PongDeadline time.Duration `mapstructure:"pong_deadline"`

	//
	VerifyRandomize       bool                   `mapstructure:"verify_randomize"`
	VerifyRandomizeSettings *settings.RandomizeLinkSettings `mapstructure:"verify_randomize_settings"`
}

//
func FallbackModuleLinkSettings() ModuleLinkSettings {
	return ModuleLinkSettings{
		TransmitFrequency:                fallbackTransmitFrequency,
		ObtainFrequency:                fallbackObtainFrequency,
		MaximumPacketSignalWorkloadExtent: fallbackMaximumPacketSignalWorkloadExtent,
		PurgeRegulate:           fallbackPurgeRegulate,
		PingDuration:            fallbackPingDuration,
		PongDeadline:             fallbackPongDeadline,
	}
}

//
func FreshModuleLinkage(
	link net.Conn,
	chnlDescriptions []*ConduitDefinition,
	uponAccept acceptClbkMethod,
	uponFailure failureClbkMethod,
) *ModuleLinkage {
	return FreshModuleLinkageUsingSettings(
		link,
		chnlDescriptions,
		uponAccept,
		uponFailure,
		FallbackModuleLinkSettings())
}

//
func FreshModuleLinkageUsingSettings(
	link net.Conn,
	chnlDescriptions []*ConduitDefinition,
	uponAccept acceptClbkMethod,
	uponFailure failureClbkMethod,
	settings ModuleLinkSettings,
) *ModuleLinkage {
	if settings.PongDeadline >= settings.PingDuration {
		panic("REDACTED")
	}

	multilink := &ModuleLinkage{
		link:          link,
		areaLinkFetcher: bufio.NewReaderSize(link, minimumFetchReserveExtent),
		areaLinkPersistor: bufio.NewWriterSize(link, minimumPersistReserveExtent),
		transmitOverseer:   stream.New(0, 0),
		obtainOverseer:   stream.New(0, 0),
		transmit:          make(chan struct{}, 1),
		pong:          make(chan struct{}, 1),
		uponAccept:     uponAccept,
		uponFailure:       uponFailure,
		settings:        settings,
		spawned:       time.Now(),
	}

	//
	conduitsOffset := map[byte]*Conduit{}
	conduits := []*Conduit{}

	for _, description := range chnlDescriptions {
		conduit := freshConduit(multilink, *description)
		conduitsOffset[conduit.description.ID] = conduit
		conduits = append(conduits, conduit)
	}
	multilink.conduits = conduits
	multilink.conduitsOffset = conduitsOffset

	multilink.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", multilink)

	//
	multilink._maxartifactsize = multilink.maximumPacketSignalExtent()

	return multilink
}

func (c *ModuleLinkage) AssignTracer(l log.Tracer) {
	c.FoundationFacility.AssignTracer(l)
	for _, ch := range c.conduits {
		ch.AssignTracer(l)
	}
}

//
func (c *ModuleLinkage) UponInitiate() error {
	if err := c.FoundationFacility.UponInitiate(); err != nil {
		return err
	}
	c.purgeClock = clock.FreshRegulateClock("REDACTED", c.settings.PurgeRegulate)
	c.pingClock = time.NewTicker(c.settings.PingDuration)
	c.pongDeadlineChnl = make(chan bool, 1)
	c.chnlMetricsClock = time.NewTicker(reviseMetrics)
	c.exitTransmitProcedure = make(chan struct{})
	c.completeTransmitProcedure = make(chan struct{})
	c.exitObtainProcedure = make(chan struct{})
	go c.transmitProcedure()
	go c.obtainProcedure()
	return nil
}

//
//
//
func (c *ModuleLinkage) haltUtilities() (earlierHalted bool) {
	c.haltMutex.Lock()
	defer c.haltMutex.Unlock()

	select {
	case <-c.exitTransmitProcedure:
		//
		return true
	default:
	}

	select {
	case <-c.exitObtainProcedure:
		//
		return true
	default:
	}

	c.FoundationFacility.UponHalt()
	c.purgeClock.Halt()
	c.pingClock.Stop()
	c.chnlMetricsClock.Stop()

	//
	close(c.exitObtainProcedure)
	close(c.exitTransmitProcedure)
	return false
}

//
//
//
//
func (c *ModuleLinkage) PurgeHalt() {
	if c.haltUtilities() {
		return
	}

	//
	{
		//
		//
		<-c.completeTransmitProcedure

		//
		//
		//
		w := protocolio.FreshSeparatedPersistor(c.areaLinkPersistor)
		eof := c.transmitFewPacketArtifacts(w)
		for !eof {
			eof = c.transmitFewPacketArtifacts(w)
		}
		c.purge()

		//
	}

	c.link.Close()

	//
	//
	//
	//

	//
}

//
func (c *ModuleLinkage) UponHalt() {
	if c.haltUtilities() {
		return
	}

	c.link.Close()

	//
	//
	//
	//
}

func (c *ModuleLinkage) Text() string {
	return fmt.Sprintf("REDACTED", c.link.RemoteAddr())
}

func (c *ModuleLinkage) purge() {
	c.Tracer.Diagnose("REDACTED", "REDACTED", c)
	err := c.areaLinkPersistor.Flush()
	if err != nil {
		c.Tracer.Diagnose("REDACTED", "REDACTED", err)
	}
}

//
func (c *ModuleLinkage) _restore() {
	if r := recover(); r != nil {
		c.Tracer.Failure("REDACTED", "REDACTED", r, "REDACTED", string(debug.Stack()))
		c.haltForeachFailure(fmt.Errorf("REDACTED", r))
	}
}

func (c *ModuleLinkage) haltForeachFailure(r any) {
	if err := c.Halt(); err != nil {
		c.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	if atomic.CompareAndSwapUint32(&c.faulted, 0, 1) {
		if c.uponFailure != nil {
			c.uponFailure(r)
		}
	}
}

//
func (c *ModuleLinkage) Transmit(chnlUUID byte, signalOctets []byte) bool {
	if !c.EqualsActive() {
		return false
	}

	c.Tracer.Diagnose("REDACTED", "REDACTED", chnlUUID, "REDACTED", c, "REDACTED", log.FreshIdleFormat("REDACTED", signalOctets))

	//
	conduit, ok := c.conduitsOffset[chnlUUID]
	if !ok {
		c.Tracer.Failure(fmt.Sprintf("REDACTED", chnlUUID))
		return false
	}

	triumph := conduit.transmitOctets(signalOctets)
	if triumph {
		//
		select {
		case c.transmit <- struct{}{}:
		default:
		}
	} else {
		c.Tracer.Diagnose("REDACTED", "REDACTED", chnlUUID, "REDACTED", c, "REDACTED", log.FreshIdleFormat("REDACTED", signalOctets))
	}
	return triumph
}

//
//
func (c *ModuleLinkage) AttemptTransmit(chnlUUID byte, signalOctets []byte) bool {
	if !c.EqualsActive() {
		return false
	}

	c.Tracer.Diagnose("REDACTED", "REDACTED", chnlUUID, "REDACTED", c, "REDACTED", log.FreshIdleFormat("REDACTED", signalOctets))

	//
	conduit, ok := c.conduitsOffset[chnlUUID]
	if !ok {
		c.Tracer.Failure(fmt.Sprintf("REDACTED", chnlUUID))
		return false
	}

	ok = conduit.attemptTransmitOctets(signalOctets)
	if ok {
		//
		select {
		case c.transmit <- struct{}{}:
		default:
		}
	}

	return ok
}

//
//
func (c *ModuleLinkage) AbleTransmit(chnlUUID byte) bool {
	if !c.EqualsActive() {
		return false
	}

	conduit, ok := c.conduitsOffset[chnlUUID]
	if !ok {
		c.Tracer.Failure(fmt.Sprintf("REDACTED", chnlUUID))
		return false
	}
	return conduit.ableTransmit()
}

//
func (c *ModuleLinkage) transmitProcedure() {
	defer c._restore()

	schemaPersistor := protocolio.FreshSeparatedPersistor(c.areaLinkPersistor)

FOREACH_CYCLE:
	for {
		var _n int
		var err error
	Preference:
		select {
		case <-c.purgeClock.Ch:
			//
			//
			c.purge()
		case <-c.chnlMetricsClock.C:
			for _, conduit := range c.conduits {
				conduit.reviseMetrics()
			}
		case <-c.pingClock.C:
			c.Tracer.Diagnose("REDACTED")
			_n, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPing{}))
			if err != nil {
				c.Tracer.Failure("REDACTED", "REDACTED", err)
				break Preference
			}
			c.transmitOverseer.Revise(_n)
			c.Tracer.Diagnose("REDACTED", "REDACTED", c.settings.PongDeadline)
			c.pongClock = time.AfterFunc(c.settings.PongDeadline, func() {
				select {
				case c.pongDeadlineChnl <- true:
				default:
				}
			})
			c.purge()
		case deadline := <-c.pongDeadlineChnl:
			if deadline {
				c.Tracer.Diagnose("REDACTED")
				err = errors.New("REDACTED")
			} else {
				c.haltPongClock()
			}
		case <-c.pong:
			c.Tracer.Diagnose("REDACTED")
			_n, err = schemaPersistor.PersistSignal(shouldEnclosePacket(&tmpfabric.PacketPong{}))
			if err != nil {
				c.Tracer.Failure("REDACTED", "REDACTED", err)
				break Preference
			}
			c.transmitOverseer.Revise(_n)
			c.purge()
		case <-c.exitTransmitProcedure:
			break FOREACH_CYCLE
		case <-c.transmit:
			//
			eof := c.transmitFewPacketArtifacts(schemaPersistor)
			if !eof {
				//
				select {
				case c.transmit <- struct{}{}:
				default:
				}
			}
		}

		if !c.EqualsActive() {
			break FOREACH_CYCLE
		}
		if err != nil {
			c.Tracer.Failure("REDACTED", "REDACTED", c, "REDACTED", err)
			c.haltForeachFailure(err)
			break FOREACH_CYCLE
		}
	}

	//
	c.haltPongClock()
	close(c.completeTransmitProcedure)
}

//
//
func (c *ModuleLinkage) transmitFewPacketArtifacts(w protocolio.Persistor) bool {
	//
	//
	//
	c.transmitOverseer.Threshold(c._maxartifactsize, c.settings.TransmitFrequency, true)

	//
	return c.transmitClusterPacketArtifacts(w, countClusterPacketArtifacts)
}

//
func (c *ModuleLinkage) transmitClusterPacketArtifacts(w protocolio.Persistor, clusterExtent int) bool {
	//
	sumOctetsRecorded := 0
	defer func() {
		if sumOctetsRecorded > 0 {
			c.transmitOverseer.Revise(sumOctetsRecorded)
		}
	}()
	for i := 0; i < clusterExtent; i++ {
		conduit := preferConduitTowardBroadcastUpon(c.conduits)
		//
		if conduit == nil {
			return true
		}
		octetsRecorded, err := c.transmitPacketSignalUponConduit(w, conduit)
		if err {
			return true
		}
		sumOctetsRecorded += octetsRecorded
	}
	return false
}

//
//
//
//
func preferConduitTowardBroadcastUpon(conduits []*Conduit) *Conduit {
	//
	//
	var minimalProportion float32 = math.MaxFloat32
	var minimalConduit *Conduit
	for _, conduit := range conduits {
		//
		//
		if !conduit.equalsTransmitAwaiting() {
			continue
		}
		//
		//
		//
		proportion := float32(conduit.latelyRelayed) / float32(conduit.description.Urgency)
		if proportion < minimalProportion {
			minimalProportion = proportion
			minimalConduit = conduit
		}
	}
	return minimalConduit
}

//
func (c *ModuleLinkage) transmitPacketSignalUponConduit(w protocolio.Persistor, transmitConduit *Conduit) (int, bool) {
	//
	n, err := transmitConduit.persistPacketSignalToward(w)
	if err != nil {
		c.Tracer.Failure("REDACTED", "REDACTED", err)
		c.haltForeachFailure(err)
		return n, true
	}
	//
	c.purgeClock.Set()
	return n, false
}

//
//
//
//
func (c *ModuleLinkage) obtainProcedure() {
	defer c._restore()

	schemaFetcher := protocolio.FreshSeparatedFetcher(c.areaLinkFetcher, c._maxartifactsize)

FOREACH_CYCLE:
	for {
		//
		c.obtainOverseer.Threshold(c._maxartifactsize, atomic.LoadInt64(&c.settings.ObtainFrequency), true)

		//
		/**
{
)
{
n
{
)
l
}
)
}
*/

		//
		var packet tmpfabric.Packet

		_n, err := schemaFetcher.FetchSignal(&packet)
		c.obtainOverseer.Revise(_n)
		if err != nil {
			//
			//
			select {
			case <-c.exitObtainProcedure:
				break FOREACH_CYCLE
			default:
			}

			if c.EqualsActive() {
				if err == io.EOF {
					c.Tracer.Details("REDACTED", "REDACTED", c)
				} else {
					c.Tracer.Diagnose("REDACTED", "REDACTED", c, "REDACTED", err)
				}
				c.haltForeachFailure(err)
			}
			break FOREACH_CYCLE
		}

		//
		switch pkt := packet.Sum.(type) {
		case *tmpfabric.Packet_Pingpacket:
			//
			//
			c.Tracer.Diagnose("REDACTED")
			select {
			case c.pong <- struct{}{}:
			default:
				//
			}
		case *tmpfabric.Packet_Pongpacket:
			c.Tracer.Diagnose("REDACTED")
			select {
			case c.pongDeadlineChnl <- false:
			default:
				//
			}
		case *tmpfabric.Packet_Packetsignal:
			conduitUUID := byte(pkt.PacketSignal.ConduitUUID)
			conduit, ok := c.conduitsOffset[conduitUUID]
			if pkt.PacketSignal.ConduitUUID < 0 || pkt.PacketSignal.ConduitUUID > math.MaxUint8 || !ok || conduit == nil {
				err := fmt.Errorf("REDACTED", pkt.PacketSignal.ConduitUUID)
				c.Tracer.Diagnose("REDACTED", "REDACTED", c, "REDACTED", err)
				c.haltForeachFailure(err)
				break FOREACH_CYCLE
			}

			signalOctets, err := conduit.obtainPacketSignal(*pkt.PacketSignal)
			if err != nil {
				if c.EqualsActive() {
					c.Tracer.Diagnose("REDACTED", "REDACTED", c, "REDACTED", err)
					c.haltForeachFailure(err)
				}
				break FOREACH_CYCLE
			}
			if signalOctets != nil {
				c.Tracer.Diagnose("REDACTED", "REDACTED", conduitUUID, "REDACTED", signalOctets)
				//
				c.uponAccept(conduitUUID, signalOctets)
			}
		default:
			err := fmt.Errorf("REDACTED", reflect.TypeOf(packet))
			c.Tracer.Failure("REDACTED", "REDACTED", c, "REDACTED", err)
			c.haltForeachFailure(err)
			break FOREACH_CYCLE
		}
	}

	//
	close(c.pong)
	//
	for range c.pong {
		//
	}
}

//
func (c *ModuleLinkage) haltPongClock() {
	if c.pongClock != nil {
		_ = c.pongClock.Stop()
		c.pongClock = nil
	}
}

//
func (c *ModuleLinkage) maximumPacketSignalExtent() int {
	bz, err := proto.Marshal(shouldEnclosePacket(&tmpfabric.PacketSignal{
		ConduitUUID: 0x01,
		EOF:       true,
		Data:      make([]byte, c.settings.MaximumPacketSignalWorkloadExtent),
	}))
	if err != nil {
		panic(err)
	}
	return len(bz)
}

type LinkageCondition struct {
	Interval    time.Duration
	TransmitOverseer stream.Condition
	ObtainOverseer stream.Condition
	Conduits    []ConduitCondition
}

type ConduitCondition struct {
	ID                byte
	TransmitStagingVolume int
	TransmitStagingExtent     int
	Urgency          int
	LatelyRelayed      int64
}

func (c *ModuleLinkage) Condition() LinkageCondition {
	var condition LinkageCondition
	condition.Interval = time.Since(c.spawned)
	condition.TransmitOverseer = c.transmitOverseer.Condition()
	condition.ObtainOverseer = c.obtainOverseer.Condition()
	condition.Conduits = make([]ConduitCondition, len(c.conduits))
	for i, conduit := range c.conduits {
		condition.Conduits[i] = ConduitCondition{
			ID:                conduit.description.ID,
			TransmitStagingVolume: cap(conduit.transmitStaging),
			TransmitStagingExtent:     int(atomic.LoadInt32(&conduit.transmitStagingExtent)),
			Urgency:          conduit.description.Urgency,
			LatelyRelayed:      atomic.LoadInt64(&conduit.latelyRelayed),
		}
	}
	return condition
}

//

type ConduitDefinition struct {
	ID                  byte
	Urgency            int
	TransmitStagingVolume   int
	ObtainReserveVolume  int
	ObtainSignalVolume int
	SignalKind         proto.Message
}

func (chnlDescription ConduitDefinition) PopulatePreset() (populated ConduitDefinition) {
	if chnlDescription.TransmitStagingVolume == 0 {
		chnlDescription.TransmitStagingVolume = fallbackTransmitStagingVolume
	}
	if chnlDescription.ObtainReserveVolume == 0 {
		chnlDescription.ObtainReserveVolume = fallbackObtainReserveVolume
	}
	if chnlDescription.ObtainSignalVolume == 0 {
		chnlDescription.ObtainSignalVolume = fallbackObtainArtifactVolume
	}
	populated = chnlDescription
	return
}

//
//
type Conduit struct {
	link          *ModuleLinkage
	description          ConduitDefinition
	transmitStaging     chan []byte
	transmitStagingExtent int32 //
	accepting       []byte
	relaying       []byte
	latelyRelayed  int64 //

	followingPacketSignal           *tmpfabric.PacketSignal
	followingPeer2peerEncapsulatorPacketSignal *tmpfabric.Packet_Packetsignal
	followingPacket              *tmpfabric.Packet

	maximumPacketSignalWorkloadExtent int

	Tracer log.Tracer
}

func freshConduit(link *ModuleLinkage, description ConduitDefinition) *Conduit {
	description = description.PopulatePreset()
	if description.Urgency <= 0 {
		panic("REDACTED")
	}
	return &Conduit{
		link:                    link,
		description:                    description,
		transmitStaging:               make(chan []byte, description.TransmitStagingVolume),
		accepting:                 make([]byte, 0, description.ObtainReserveVolume),
		followingPacketSignal:           &tmpfabric.PacketSignal{ConduitUUID: int32(description.ID)},
		followingPeer2peerEncapsulatorPacketSignal: &tmpfabric.Packet_Packetsignal{},
		followingPacket:              &tmpfabric.Packet{},
		maximumPacketSignalWorkloadExtent: link.settings.MaximumPacketSignalWorkloadExtent,
	}
}

func (ch *Conduit) AssignTracer(l log.Tracer) {
	ch.Tracer = l
}

//
//
//
func (ch *Conduit) transmitOctets(octets []byte) bool {
	select {
	case ch.transmitStaging <- octets:
		atomic.AddInt32(&ch.transmitStagingExtent, 1)
		return true
	case <-time.After(fallbackTransmitDeadline):
		return false
	}
}

//
//
//
func (ch *Conduit) attemptTransmitOctets(octets []byte) bool {
	select {
	case ch.transmitStaging <- octets:
		atomic.AddInt32(&ch.transmitStagingExtent, 1)
		return true
	default:
		return false
	}
}

//
func (ch *Conduit) fetchTransmitStagingExtent() (extent int) {
	return int(atomic.LoadInt32(&ch.transmitStagingExtent))
}

//
//
func (ch *Conduit) ableTransmit() bool {
	return ch.fetchTransmitStagingExtent() < fallbackTransmitStagingVolume
}

//
//
//
func (ch *Conduit) equalsTransmitAwaiting() bool {
	if len(ch.relaying) == 0 {
		if len(ch.transmitStaging) == 0 {
			return false
		}
		ch.relaying = <-ch.transmitStaging
	}
	return true
}

//
//
func (ch *Conduit) reviseFollowingPacket() {
	maximumExtent := ch.maximumPacketSignalWorkloadExtent
	if len(ch.relaying) <= maximumExtent {
		ch.followingPacketSignal.Data = ch.relaying
		ch.followingPacketSignal.EOF = true
		ch.relaying = nil
		atomic.AddInt32(&ch.transmitStagingExtent, -1) //
	} else {
		ch.followingPacketSignal.Data = ch.relaying[:maximumExtent]
		ch.followingPacketSignal.EOF = false
		ch.relaying = ch.relaying[maximumExtent:]
	}

	ch.followingPeer2peerEncapsulatorPacketSignal.PacketSignal = ch.followingPacketSignal
	ch.followingPacket.Sum = ch.followingPeer2peerEncapsulatorPacketSignal
}

//
//
func (ch *Conduit) persistPacketSignalToward(w protocolio.Persistor) (n int, err error) {
	ch.reviseFollowingPacket()
	n, err = w.PersistSignal(ch.followingPacket)
	if err != nil {
		return 0, err
	}
	atomic.AddInt64(&ch.latelyRelayed, int64(n))
	return n, nil
}

//
//
//
func (ch *Conduit) obtainPacketSignal(packet tmpfabric.PacketSignal) ([]byte, error) {
	ch.Tracer.Diagnose("REDACTED", "REDACTED", ch.link, "REDACTED", packet)
	obtainCeiling, obtainAccepted := ch.description.ObtainSignalVolume, len(ch.accepting)+len(packet.Data)
	if obtainCeiling < obtainAccepted {
		return nil, fmt.Errorf("REDACTED", obtainCeiling, obtainAccepted)
	}
	ch.accepting = append(ch.accepting, packet.Data...)
	if packet.EOF {
		signalOctets := ch.accepting

		//
		//
		//
		//
		ch.accepting = ch.accepting[:0] //
		return signalOctets, nil
	}
	return nil, nil
}

//
//
func (ch *Conduit) reviseMetrics() {
	//
	//
	atomic.StoreInt64(&ch.latelyRelayed, int64(float64(atomic.LoadInt64(&ch.latelyRelayed))*0.8))
}

//
//

//
func shouldEnclosePacket(pb proto.Message) *tmpfabric.Packet {
	msg := &tmpfabric.Packet{}
	shouldEnclosePacketWithin(pb, msg)
	return msg
}

func shouldEnclosePacketWithin(pb proto.Message, dst *tmpfabric.Packet) {
	switch pb := pb.(type) {
	case *tmpfabric.PacketPing:
		dst.Sum = &tmpfabric.Packet_Pingpacket{
			PacketPing: pb,
		}
	case *tmpfabric.PacketPong:
		dst.Sum = &tmpfabric.Packet_Pongpacket{
			PacketPong: pb,
		}
	case *tmpfabric.PacketSignal:
		dst.Sum = &tmpfabric.Packet_Packetsignal{
			PacketSignal: pb,
		}
	default:
		panic(fmt.Errorf("REDACTED", pb))
	}
}
