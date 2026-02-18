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

	"github.com/valkyrieworks/settings"
	stream "github.com/valkyrieworks/utils/pace"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/protoio"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/utils/clock"
	tmp2p "github.com/valkyrieworks/schema/consensuscore/p2p"
)

const (
	standardMaximumPackageMessageShipmentVolume = 1024

	countClusterPackageNotices = 10
	minimumReadFrameVolume  = 1024
	minimumRecordFrameVolume = 65536
	modifyMetrics        = 2 * time.Second

	//
	//
	//
	standardPurgeRegulate = 10 * time.Millisecond

	standardTransmitBufferAbility   = 1
	standardReceiveFrameAbility  = 4096
	standardReceiveSignalAbility = 22020096      //
	standardTransmitRatio            = int64(512000) //
	standardReceiveRatio            = int64(512000) //
	standardTransmitDeadline         = 10 * time.Second
	standardPingCadence        = 60 * time.Second
	standardPongDeadline         = 45 * time.Second
)

type (
	acceptCallbackfnFunction func(chanUID byte, messageOctets []byte)
	faultCallbackfnFunction   func(any)
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
type MLinkage struct {
	daemon.RootDaemon

	link          net.Conn
	imageLinkScanner *bufio.Reader
	imageLinkRecorder *bufio.Writer
	transmitAuditor   *stream.Auditor
	acceptAuditor   *stream.Auditor
	transmit          chan struct{}
	pong          chan struct{}
	streams      []*Conduit
	streamsIdx   map[byte]*Conduit
	onAccept     acceptCallbackfnFunction
	onFault       faultCallbackfnFunction
	failed       uint32
	settings        MLinkSettings

	//
	//
	exitTransmitProcedure chan struct{}
	doneTransmitProcedure chan struct{}

	//
	exitReceiveProcedure chan struct{}

	//
	//
	haltMutex engineconnect.Lock

	purgeClock *clock.RegulateClock //
	pingClock  *time.Ticker         //

	//
	pongClock     *time.Timer
	pongDeadlineChan chan bool //

	chanMetricsClock *time.Ticker //

	spawned time.Time //

	_maxpacketmsgsize int
}

//
type MLinkSettings struct {
	TransmitRatio int64 `mapstructure:"transmit_ratio"`
	ReceiveRatio int64 `mapstructure:"receive_ratio"`

	//
	MaximumPackageMessageShipmentVolume int `mapstructure:"maximum_package_message_shipment_volume"`

	//
	PurgeRegulate time.Duration `mapstructure:"purge_regulate"`

	//
	PingCadence time.Duration `mapstructure:"ping_cadence"`

	//
	PongDeadline time.Duration `mapstructure:"pong_deadline"`

	//
	VerifyRandomize       bool                   `mapstructure:"verify_randomize"`
	VerifyRandomizeSettings *settings.RandomizeLinkSettings `mapstructure:"verify_randomize_settings"`
}

//
func StandardMLinkSettings() MLinkSettings {
	return MLinkSettings{
		TransmitRatio:                standardTransmitRatio,
		ReceiveRatio:                standardReceiveRatio,
		MaximumPackageMessageShipmentVolume: standardMaximumPackageMessageShipmentVolume,
		PurgeRegulate:           standardPurgeRegulate,
		PingCadence:            standardPingCadence,
		PongDeadline:             standardPongDeadline,
	}
}

//
func NewMLinkage(
	link net.Conn,
	chanTraits []*StreamDefinition,
	onAccept acceptCallbackfnFunction,
	onFault faultCallbackfnFunction,
) *MLinkage {
	return NewMLinkageWithSettings(
		link,
		chanTraits,
		onAccept,
		onFault,
		StandardMLinkSettings())
}

//
func NewMLinkageWithSettings(
	link net.Conn,
	chanTraits []*StreamDefinition,
	onAccept acceptCallbackfnFunction,
	onFault faultCallbackfnFunction,
	settings MLinkSettings,
) *MLinkage {
	if settings.PongDeadline >= settings.PingCadence {
		panic("REDACTED")
	}

	mconn := &MLinkage{
		link:          link,
		imageLinkScanner: bufio.NewReaderSize(link, minimumReadFrameVolume),
		imageLinkRecorder: bufio.NewWriterSize(link, minimumRecordFrameVolume),
		transmitAuditor:   stream.New(0, 0),
		acceptAuditor:   stream.New(0, 0),
		transmit:          make(chan struct{}, 1),
		pong:          make(chan struct{}, 1),
		onAccept:     onAccept,
		onFault:       onFault,
		settings:        settings,
		spawned:       time.Now(),
	}

	//
	streamsIdx := map[byte]*Conduit{}
	streams := []*Conduit{}

	for _, note := range chanTraits {
		conduit := newConduit(mconn, *note)
		streamsIdx[conduit.note.ID] = conduit
		streams = append(streams, conduit)
	}
	mconn.streams = streams
	mconn.streamsIdx = streamsIdx

	mconn.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", mconn)

	//
	mconn._maxpacketmsgsize = mconn.maximumPackageMessageVolume()

	return mconn
}

func (c *MLinkage) AssignTracer(l log.Tracer) {
	c.RootDaemon.AssignTracer(l)
	for _, ch := range c.streams {
		ch.AssignTracer(l)
	}
}

//
func (c *MLinkage) OnBegin() error {
	if err := c.RootDaemon.OnBegin(); err != nil {
		return err
	}
	c.purgeClock = clock.NewRegulateClock("REDACTED", c.settings.PurgeRegulate)
	c.pingClock = time.NewTicker(c.settings.PingCadence)
	c.pongDeadlineChan = make(chan bool, 1)
	c.chanMetricsClock = time.NewTicker(modifyMetrics)
	c.exitTransmitProcedure = make(chan struct{})
	c.doneTransmitProcedure = make(chan struct{})
	c.exitReceiveProcedure = make(chan struct{})
	go c.transmitProcedure()
	go c.receiveProcedure()
	return nil
}

//
//
//
func (c *MLinkage) haltIfaces() (yetCeased bool) {
	c.haltMutex.Lock()
	defer c.haltMutex.Unlock()

	select {
	case <-c.exitTransmitProcedure:
		//
		return true
	default:
	}

	select {
	case <-c.exitReceiveProcedure:
		//
		return true
	default:
	}

	c.RootDaemon.OnHalt()
	c.purgeClock.Halt()
	c.pingClock.Stop()
	c.chanMetricsClock.Stop()

	//
	close(c.exitReceiveProcedure)
	close(c.exitTransmitProcedure)
	return false
}

//
//
//
//
func (c *MLinkage) PurgeHalt() {
	if c.haltIfaces() {
		return
	}

	//
	{
		//
		//
		<-c.doneTransmitProcedure

		//
		//
		//
		w := protoio.NewSeparatedRecorder(c.imageLinkRecorder)
		eof := c.transmitSomePackageNotices(w)
		for !eof {
			eof = c.transmitSomePackageNotices(w)
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
func (c *MLinkage) OnHalt() {
	if c.haltIfaces() {
		return
	}

	c.link.Close()

	//
	//
	//
	//
}

func (c *MLinkage) String() string {
	return fmt.Sprintf("REDACTED", c.link.RemoteAddr())
}

func (c *MLinkage) purge() {
	c.Tracer.Diagnose("REDACTED", "REDACTED", c)
	err := c.imageLinkRecorder.Flush()
	if err != nil {
		c.Tracer.Diagnose("REDACTED", "REDACTED", err)
	}
}

//
func (c *MLinkage) _recoup() {
	if r := recover(); r != nil {
		c.Tracer.Fault("REDACTED", "REDACTED", r, "REDACTED", string(debug.Stack()))
		c.haltForFault(fmt.Errorf("REDACTED", r))
	}
}

func (c *MLinkage) haltForFault(r any) {
	if err := c.Halt(); err != nil {
		c.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	if atomic.CompareAndSwapUint32(&c.failed, 0, 1) {
		if c.onFault != nil {
			c.onFault(r)
		}
	}
}

//
func (c *MLinkage) Transmit(chanUID byte, messageOctets []byte) bool {
	if !c.IsActive() {
		return false
	}

	c.Tracer.Diagnose("REDACTED", "REDACTED", chanUID, "REDACTED", c, "REDACTED", log.NewIdleFormat("REDACTED", messageOctets))

	//
	conduit, ok := c.streamsIdx[chanUID]
	if !ok {
		c.Tracer.Fault(fmt.Sprintf("REDACTED", chanUID))
		return false
	}

	success := conduit.transmitOctets(messageOctets)
	if success {
		//
		select {
		case c.transmit <- struct{}{}:
		default:
		}
	} else {
		c.Tracer.Diagnose("REDACTED", "REDACTED", chanUID, "REDACTED", c, "REDACTED", log.NewIdleFormat("REDACTED", messageOctets))
	}
	return success
}

//
//
func (c *MLinkage) AttemptTransmit(chanUID byte, messageOctets []byte) bool {
	if !c.IsActive() {
		return false
	}

	c.Tracer.Diagnose("REDACTED", "REDACTED", chanUID, "REDACTED", c, "REDACTED", log.NewIdleFormat("REDACTED", messageOctets))

	//
	conduit, ok := c.streamsIdx[chanUID]
	if !ok {
		c.Tracer.Fault(fmt.Sprintf("REDACTED", chanUID))
		return false
	}

	ok = conduit.attemptTransmitOctets(messageOctets)
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
func (c *MLinkage) MayTransmit(chanUID byte) bool {
	if !c.IsActive() {
		return false
	}

	conduit, ok := c.streamsIdx[chanUID]
	if !ok {
		c.Tracer.Fault(fmt.Sprintf("REDACTED", chanUID))
		return false
	}
	return conduit.mayTransmit()
}

//
func (c *MLinkage) transmitProcedure() {
	defer c._recoup()

	schemaRecorder := protoio.NewSeparatedRecorder(c.imageLinkRecorder)

FOR_CYCLE:
	for {
		var _n int
		var err error
	Preference:
		select {
		case <-c.purgeClock.Ch:
			//
			//
			c.purge()
		case <-c.chanMetricsClock.C:
			for _, conduit := range c.streams {
				conduit.modifyMetrics()
			}
		case <-c.pingClock.C:
			c.Tracer.Diagnose("REDACTED")
			_n, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePing{}))
			if err != nil {
				c.Tracer.Fault("REDACTED", "REDACTED", err)
				break Preference
			}
			c.transmitAuditor.Modify(_n)
			c.Tracer.Diagnose("REDACTED", "REDACTED", c.settings.PongDeadline)
			c.pongClock = time.AfterFunc(c.settings.PongDeadline, func() {
				select {
				case c.pongDeadlineChan <- true:
				default:
				}
			})
			c.purge()
		case deadline := <-c.pongDeadlineChan:
			if deadline {
				c.Tracer.Diagnose("REDACTED")
				err = errors.New("REDACTED")
			} else {
				c.haltPongClock()
			}
		case <-c.pong:
			c.Tracer.Diagnose("REDACTED")
			_n, err = schemaRecorder.RecordMessage(shouldEnclosePackage(&tmp2p.PackagePong{}))
			if err != nil {
				c.Tracer.Fault("REDACTED", "REDACTED", err)
				break Preference
			}
			c.transmitAuditor.Modify(_n)
			c.purge()
		case <-c.exitTransmitProcedure:
			break FOR_CYCLE
		case <-c.transmit:
			//
			eof := c.transmitSomePackageNotices(schemaRecorder)
			if !eof {
				//
				select {
				case c.transmit <- struct{}{}:
				default:
				}
			}
		}

		if !c.IsActive() {
			break FOR_CYCLE
		}
		if err != nil {
			c.Tracer.Fault("REDACTED", "REDACTED", c, "REDACTED", err)
			c.haltForFault(err)
			break FOR_CYCLE
		}
	}

	//
	c.haltPongClock()
	close(c.doneTransmitProcedure)
}

//
//
func (c *MLinkage) transmitSomePackageNotices(w protoio.Recorder) bool {
	//
	//
	//
	c.transmitAuditor.Ceiling(c._maxpacketmsgsize, c.settings.TransmitRatio, true)

	//
	return c.transmitClusterPackageNotices(w, countClusterPackageNotices)
}

//
func (c *MLinkage) transmitClusterPackageNotices(w protoio.Recorder, clusterVolume int) bool {
	//
	sumOctetsInscribed := 0
	defer func() {
		if sumOctetsInscribed > 0 {
			c.transmitAuditor.Modify(sumOctetsInscribed)
		}
	}()
	for i := 0; i < clusterVolume; i++ {
		conduit := chooseConduitToGossipOn(c.streams)
		//
		if conduit == nil {
			return true
		}
		octetsInscribed, err := c.transmitPackageMessageOnConduit(w, conduit)
		if err {
			return true
		}
		sumOctetsInscribed += octetsInscribed
	}
	return false
}

//
//
//
//
func chooseConduitToGossipOn(streams []*Conduit) *Conduit {
	//
	//
	var minimumProportion float32 = math.MaxFloat32
	var minimumConduit *Conduit
	for _, conduit := range streams {
		//
		//
		if !conduit.isTransmitAwaiting() {
			continue
		}
		//
		//
		//
		proportion := float32(conduit.latelyRelayed) / float32(conduit.note.Urgency)
		if proportion < minimumProportion {
			minimumProportion = proportion
			minimumConduit = conduit
		}
	}
	return minimumConduit
}

//
func (c *MLinkage) transmitPackageMessageOnConduit(w protoio.Recorder, transmitConduit *Conduit) (int, bool) {
	//
	n, err := transmitConduit.recordPackageMessageTo(w)
	if err != nil {
		c.Tracer.Fault("REDACTED", "REDACTED", err)
		c.haltForFault(err)
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
func (c *MLinkage) receiveProcedure() {
	defer c._recoup()

	schemaScanner := protoio.NewSeparatedScanner(c.imageLinkScanner, c._maxpacketmsgsize)

FOR_CYCLE:
	for {
		//
		c.acceptAuditor.Ceiling(c._maxpacketmsgsize, atomic.LoadInt64(&c.settings.ReceiveRatio), true)

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
		var package tmp2p.Package

		_n, err := schemaScanner.ScanMessage(&package)
		c.acceptAuditor.Modify(_n)
		if err != nil {
			//
			//
			select {
			case <-c.exitReceiveProcedure:
				break FOR_CYCLE
			default:
			}

			if c.IsActive() {
				if err == io.EOF {
					c.Tracer.Details("REDACTED", "REDACTED", c)
				} else {
					c.Tracer.Diagnose("REDACTED", "REDACTED", c, "REDACTED", err)
				}
				c.haltForFault(err)
			}
			break FOR_CYCLE
		}

		//
		switch pkt := package.Sum.(type) {
		case *tmp2p.Package_Packageping:
			//
			//
			c.Tracer.Diagnose("REDACTED")
			select {
			case c.pong <- struct{}{}:
			default:
				//
			}
		case *tmp2p.Package_Packagepong:
			c.Tracer.Diagnose("REDACTED")
			select {
			case c.pongDeadlineChan <- false:
			default:
				//
			}
		case *tmp2p.Package_Messagedata:
			conduitUID := byte(pkt.PackageMessage.StreamUID)
			conduit, ok := c.streamsIdx[conduitUID]
			if pkt.PackageMessage.StreamUID < 0 || pkt.PackageMessage.StreamUID > math.MaxUint8 || !ok || conduit == nil {
				err := fmt.Errorf("REDACTED", pkt.PackageMessage.StreamUID)
				c.Tracer.Diagnose("REDACTED", "REDACTED", c, "REDACTED", err)
				c.haltForFault(err)
				break FOR_CYCLE
			}

			messageOctets, err := conduit.receivePackageMessage(*pkt.PackageMessage)
			if err != nil {
				if c.IsActive() {
					c.Tracer.Diagnose("REDACTED", "REDACTED", c, "REDACTED", err)
					c.haltForFault(err)
				}
				break FOR_CYCLE
			}
			if messageOctets != nil {
				c.Tracer.Diagnose("REDACTED", "REDACTED", conduitUID, "REDACTED", messageOctets)
				//
				c.onAccept(conduitUID, messageOctets)
			}
		default:
			err := fmt.Errorf("REDACTED", reflect.TypeOf(package))
			c.Tracer.Fault("REDACTED", "REDACTED", c, "REDACTED", err)
			c.haltForFault(err)
			break FOR_CYCLE
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
func (c *MLinkage) haltPongClock() {
	if c.pongClock != nil {
		_ = c.pongClock.Stop()
		c.pongClock = nil
	}
}

//
func (c *MLinkage) maximumPackageMessageVolume() int {
	bz, err := proto.Marshal(shouldEnclosePackage(&tmp2p.PackageMessage{
		StreamUID: 0x01,
		EOF:       true,
		Data:      make([]byte, c.settings.MaximumPackageMessageShipmentVolume),
	}))
	if err != nil {
		panic(err)
	}
	return len(bz)
}

type LinkageState struct {
	Period    time.Duration
	TransmitAuditor stream.Status
	ReceiveAuditor stream.Status
	Streams    []ConduitState
}

type ConduitState struct {
	ID                byte
	TransmitBufferVolume int
	TransmitBufferVolume     int
	Urgency          int
	LatelyRelayed      int64
}

func (c *MLinkage) Status() LinkageState {
	var state LinkageState
	state.Period = time.Since(c.spawned)
	state.TransmitAuditor = c.transmitAuditor.Status()
	state.ReceiveAuditor = c.acceptAuditor.Status()
	state.Streams = make([]ConduitState, len(c.streams))
	for i, conduit := range c.streams {
		state.Streams[i] = ConduitState{
			ID:                conduit.note.ID,
			TransmitBufferVolume: cap(conduit.transmitBuffer),
			TransmitBufferVolume:     int(atomic.LoadInt32(&conduit.transmitBufferVolume)),
			Urgency:          conduit.note.Urgency,
			LatelyRelayed:      atomic.LoadInt64(&conduit.latelyRelayed),
		}
	}
	return state
}

//

type StreamDefinition struct {
	ID                  byte
	Urgency            int
	TransmitBufferVolume   int
	AcceptBufferVolume  int
	AcceptSignalVolume int
	SignalKind         proto.Message
}

func (chanNote StreamDefinition) PopulateStandards() (populated StreamDefinition) {
	if chanNote.TransmitBufferVolume == 0 {
		chanNote.TransmitBufferVolume = standardTransmitBufferAbility
	}
	if chanNote.AcceptBufferVolume == 0 {
		chanNote.AcceptBufferVolume = standardReceiveFrameAbility
	}
	if chanNote.AcceptSignalVolume == 0 {
		chanNote.AcceptSignalVolume = standardReceiveSignalAbility
	}
	populated = chanNote
	return
}

//
//
type Conduit struct {
	link          *MLinkage
	note          StreamDefinition
	transmitBuffer     chan []byte
	transmitBufferVolume int32 //
	accepting       []byte
	dispatching       []byte
	latelyRelayed  int64 //

	followingPackageMessage           *tmp2p.PackageMessage
	followingP2pAdapterPackageMessage *tmp2p.Package_Messagedata
	followingPackage              *tmp2p.Package

	maximumPackageMessageShipmentVolume int

	Tracer log.Tracer
}

func newConduit(link *MLinkage, note StreamDefinition) *Conduit {
	note = note.PopulateStandards()
	if note.Urgency <= 0 {
		panic("REDACTED")
	}
	return &Conduit{
		link:                    link,
		note:                    note,
		transmitBuffer:               make(chan []byte, note.TransmitBufferVolume),
		accepting:                 make([]byte, 0, note.AcceptBufferVolume),
		followingPackageMessage:           &tmp2p.PackageMessage{StreamUID: int32(note.ID)},
		followingP2pAdapterPackageMessage: &tmp2p.Package_Messagedata{},
		followingPackage:              &tmp2p.Package{},
		maximumPackageMessageShipmentVolume: link.settings.MaximumPackageMessageShipmentVolume,
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
	case ch.transmitBuffer <- octets:
		atomic.AddInt32(&ch.transmitBufferVolume, 1)
		return true
	case <-time.After(standardTransmitDeadline):
		return false
	}
}

//
//
//
func (ch *Conduit) attemptTransmitOctets(octets []byte) bool {
	select {
	case ch.transmitBuffer <- octets:
		atomic.AddInt32(&ch.transmitBufferVolume, 1)
		return true
	default:
		return false
	}
}

//
func (ch *Conduit) importTransmitBufferVolume() (volume int) {
	return int(atomic.LoadInt32(&ch.transmitBufferVolume))
}

//
//
func (ch *Conduit) mayTransmit() bool {
	return ch.importTransmitBufferVolume() < standardTransmitBufferAbility
}

//
//
//
func (ch *Conduit) isTransmitAwaiting() bool {
	if len(ch.dispatching) == 0 {
		if len(ch.transmitBuffer) == 0 {
			return false
		}
		ch.dispatching = <-ch.transmitBuffer
	}
	return true
}

//
//
func (ch *Conduit) modifyFollowingPackage() {
	maximumVolume := ch.maximumPackageMessageShipmentVolume
	if len(ch.dispatching) <= maximumVolume {
		ch.followingPackageMessage.Data = ch.dispatching
		ch.followingPackageMessage.EOF = true
		ch.dispatching = nil
		atomic.AddInt32(&ch.transmitBufferVolume, -1) //
	} else {
		ch.followingPackageMessage.Data = ch.dispatching[:maximumVolume]
		ch.followingPackageMessage.EOF = false
		ch.dispatching = ch.dispatching[maximumVolume:]
	}

	ch.followingP2pAdapterPackageMessage.PackageMessage = ch.followingPackageMessage
	ch.followingPackage.Sum = ch.followingP2pAdapterPackageMessage
}

//
//
func (ch *Conduit) recordPackageMessageTo(w protoio.Recorder) (n int, err error) {
	ch.modifyFollowingPackage()
	n, err = w.RecordMessage(ch.followingPackage)
	if err != nil {
		return 0, err
	}
	atomic.AddInt64(&ch.latelyRelayed, int64(n))
	return n, nil
}

//
//
//
func (ch *Conduit) receivePackageMessage(package tmp2p.PackageMessage) ([]byte, error) {
	ch.Tracer.Diagnose("REDACTED", "REDACTED", ch.link, "REDACTED", package)
	receiveLimit, receiveAccepted := ch.note.AcceptSignalVolume, len(ch.accepting)+len(package.Data)
	if receiveLimit < receiveAccepted {
		return nil, fmt.Errorf("REDACTED", receiveLimit, receiveAccepted)
	}
	ch.accepting = append(ch.accepting, package.Data...)
	if package.EOF {
		messageOctets := ch.accepting

		//
		//
		//
		//
		ch.accepting = ch.accepting[:0] //
		return messageOctets, nil
	}
	return nil, nil
}

//
//
func (ch *Conduit) modifyMetrics() {
	//
	//
	atomic.StoreInt64(&ch.latelyRelayed, int64(float64(atomic.LoadInt64(&ch.latelyRelayed))*0.8))
}

//
//

//
func shouldEnclosePackage(pb proto.Message) *tmp2p.Package {
	msg := &tmp2p.Package{}
	shouldEnclosePackageToward(pb, msg)
	return msg
}

func shouldEnclosePackageToward(pb proto.Message, dst *tmp2p.Package) {
	switch pb := pb.(type) {
	case *tmp2p.PackagePing:
		dst.Sum = &tmp2p.Package_Packageping{
			PackagePing: pb,
		}
	case *tmp2p.PackagePong:
		dst.Sum = &tmp2p.Package_Packagepong{
			PackagePong: pb,
		}
	case *tmp2p.PackageMessage:
		dst.Sum = &tmp2p.Package_Messagedata{
			PackageMessage: pb,
		}
	default:
		panic(fmt.Errorf("REDACTED", pb))
	}
}
