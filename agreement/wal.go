package agreement

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"path/filepath"
	"time"

	"github.com/cosmos/gogoproto/proto"

	automatic "github.com/valkyrieworks/utils/autofile"
	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/utils/log"
	cometos "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/utils/daemon"
	cometconnect "github.com/valkyrieworks/schema/consensuscore/agreement"
	cometfaults "github.com/valkyrieworks/kinds/faults"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

const (
	//
	maximumMessageVolumeOctets = maximumMessageVolume + 24

	//
	journalStandardPurgeCadence = 2 * time.Second
)

//
//

//
type ScheduledJournalSignal struct {
	Time time.Time  `json:"moment"`
	Msg  JournalSignal `json:"msg"`
}

//
//
type TerminateLevelSignal struct {
	Level int64 `json:"level"`
}

type JournalSignal any

func init() {
	cometjson.EnrollKind(messageDetails{}, "REDACTED")
	cometjson.EnrollKind(deadlineDetails{}, "REDACTED")
	cometjson.EnrollKind(TerminateLevelSignal{}, "REDACTED")
}

//
//

//
type WAL interface {
	Record(JournalSignal) error
	RecordAlign(JournalSignal) error
	PurgeAndAlign() error

	ScanForTerminateLevel(level int64, options *JournalScanSettings) (rd io.ReadCloser, located bool, err error)

	//
	Begin() error
	Halt() error
	Wait()
}

//
//
//
//
//
type RootJournal struct {
	daemon.RootDaemon

	cluster *automatic.Cluster

	enc *JournalSerializer

	purgeTimer   *time.Ticker
	purgeCadence time.Duration
}

var _ WAL = &RootJournal{}

//
//
func NewJournal(journalEntry string, clusterSettings ...func(*automatic.Cluster)) (*RootJournal, error) {
	err := cometos.AssureFolder(filepath.Dir(journalEntry), 0o700)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	cluster, err := automatic.AccessCluster(journalEntry, clusterSettings...)
	if err != nil {
		return nil, err
	}
	wal := &RootJournal{
		cluster:         cluster,
		enc:           NewJournalSerializer(cluster),
		purgeCadence: journalStandardPurgeCadence,
	}
	wal.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", wal)
	return wal, nil
}

//
func (wal *RootJournal) AssignPurgeCadence(i time.Duration) {
	wal.purgeCadence = i
}

func (wal *RootJournal) Cluster() *automatic.Cluster {
	return wal.cluster
}

func (wal *RootJournal) AssignTracer(l log.Tracer) {
	wal.Tracer = l
	wal.cluster.AssignTracer(l)
}

func (wal *RootJournal) OnBegin() error {
	volume, err := wal.cluster.Front.Volume()
	if err != nil {
		return err
	} else if volume == 0 {
		if err := wal.RecordAlign(TerminateLevelSignal{0}); err != nil {
			return err
		}
	}
	err = wal.cluster.Begin()
	if err != nil {
		return err
	}
	wal.purgeTimer = time.NewTicker(wal.purgeCadence)
	go wal.handlePurgeTicks()
	return nil
}

func (wal *RootJournal) handlePurgeTicks() {
	for {
		select {
		case <-wal.purgeTimer.C:
			if err := wal.PurgeAndAlign(); err != nil {
				wal.Tracer.Fault("REDACTED", "REDACTED", err)
			}
		case <-wal.Exit():
			return
		}
	}
}

//
//
func (wal *RootJournal) PurgeAndAlign() error {
	return wal.cluster.PurgeAndAlign()
}

//
//
//
func (wal *RootJournal) OnHalt() {
	wal.purgeTimer.Stop()
	if err := wal.PurgeAndAlign(); err != nil {
		wal.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	if err := wal.cluster.Halt(); err != nil {
		wal.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	wal.cluster.End()
}

//
//
func (wal *RootJournal) Wait() {
	wal.cluster.Wait()
}

//
//
//
func (wal *RootJournal) Record(msg JournalSignal) error {
	if wal == nil {
		return nil
	}

	if err := wal.enc.Serialize(&ScheduledJournalSignal{engineclock.Now(), msg}); err != nil {
		wal.Tracer.Fault("REDACTED",
			"REDACTED", err, "REDACTED", msg)
		return err
	}

	return nil
}

//
//
//
func (wal *RootJournal) RecordAlign(msg JournalSignal) error {
	if wal == nil {
		return nil
	}

	if err := wal.Record(msg); err != nil {
		return err
	}

	if err := wal.PurgeAndAlign(); err != nil {
		wal.Tracer.Fault(`REDACTED.
REDACTED`,
			"REDACTED", err)
		return err
	}

	return nil
}

//
type JournalScanSettings struct {
	//
	BypassDataImpairmentFaults bool
}

//
//
//
//
//
func (wal *RootJournal) ScanForTerminateLevel(
	level int64,
	options *JournalScanSettings,
) (rd io.ReadCloser, located bool, err error) {
	var (
		msg *ScheduledJournalSignal
		gr  *automatic.ClusterScanner
	)
	finalLevelLocated := int64(-1)

	//
	//
	min, max := wal.cluster.MinimumOrdinal(), wal.cluster.MaximumOrdinal()
	wal.Tracer.Details("REDACTED", "REDACTED", level, "REDACTED", min, "REDACTED", max)
	for ordinal := max; ordinal >= min; ordinal-- {
		gr, err = wal.cluster.NewScanner(ordinal)
		if err != nil {
			return nil, false, err
		}

		dec := NewJournalParser(gr)
		for {
			msg, err = dec.Parse()
			if err == io.EOF {
				//
				if finalLevelLocated > 0 && finalLevelLocated < level {
					gr.End()
					return nil, false, nil
				}
				//
				break
			}
			if options.BypassDataImpairmentFaults && IsDataImpairmentFault(err) {
				wal.Tracer.Fault("REDACTED", "REDACTED", err)
				//
				continue
			} else if err != nil {
				gr.End()
				return nil, false, err
			}

			if m, ok := msg.Msg.(TerminateLevelSignal); ok {
				finalLevelLocated = m.Level
				if m.Level == level { //
					wal.Tracer.Details("REDACTED", "REDACTED", level, "REDACTED", ordinal)
					return gr, true, nil
				}
			}
		}
		gr.End()
	}

	return nil, false, nil
}

//
//
//
type JournalSerializer struct {
	wr io.Writer
}

//
func NewJournalSerializer(wr io.Writer) *JournalSerializer {
	return &JournalSerializer{wr}
}

//
//
//
func (enc *JournalSerializer) Serialize(v *ScheduledJournalSignal) error {
	pbMessage, err := JournalToSchema(v.Msg)
	if err != nil {
		return err
	}
	pv := cometconnect.ScheduledJournalSignal{
		Time: v.Time,
		Msg:  pbMessage,
	}

	data, err := proto.Marshal(&pv)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	crc := crc32.Checksum(data, crc32c)
	extent := uint32(len(data))
	if extent > maximumMessageVolumeOctets {
		return fmt.Errorf("REDACTED", extent, maximumMessageVolumeOctets)
	}
	sumExtent := 8 + int(extent)

	msg := make([]byte, sumExtent)
	binary.BigEndian.PutUint32(msg[0:4], crc)
	binary.BigEndian.PutUint32(msg[4:8], extent)
	copy(msg[8:], data)

	_, err = enc.wr.Write(msg)
	return err
}

//
func IsDataImpairmentFault(err error) bool {
	_, ok := err.(DataImpairmentFault)
	return ok
}

//
type DataImpairmentFault struct {
	origin error
}

func (e DataImpairmentFault) Fault() string {
	return fmt.Sprintf("REDACTED", e.origin)
}

func (e DataImpairmentFault) Origin() error {
	return e.origin
}

//
//
//
//
//
type JournalParser struct {
	rd io.Reader
}

//
func NewJournalParser(rd io.Reader) *JournalParser {
	return &JournalParser{rd}
}

//
func (dec *JournalParser) Parse() (*ScheduledJournalSignal, error) {
	b := make([]byte, 4)

	_, err := dec.rd.Read(b)
	if errors.Is(err, io.EOF) {
		return nil, err
	}
	if err != nil {
		return nil, DataImpairmentFault{fmt.Errorf("REDACTED", err)}
	}
	crc := binary.BigEndian.Uint32(b)

	b = make([]byte, 4)
	_, err = dec.rd.Read(b)
	if err != nil {
		return nil, DataImpairmentFault{fmt.Errorf("REDACTED", err)}
	}
	extent := binary.BigEndian.Uint32(b)

	if extent > maximumMessageVolumeOctets {
		return nil, DataImpairmentFault{fmt.Errorf(
			"REDACTED",
			extent,
			maximumMessageVolumeOctets)}
	}

	data := make([]byte, extent)
	n, err := dec.rd.Read(data)
	if err != nil {
		return nil, DataImpairmentFault{fmt.Errorf("REDACTED", err, n, extent)}
	}

	//
	factualCRC := crc32.Checksum(data, crc32c)
	if factualCRC != crc {
		return nil, DataImpairmentFault{fmt.Errorf("REDACTED", crc, factualCRC)}
	}

	res := new(cometconnect.ScheduledJournalSignal)
	err = proto.Unmarshal(data, res)
	if err != nil {
		return nil, DataImpairmentFault{fmt.Errorf("REDACTED", err)}
	}

	journalMessage, err := JournalFromSchema(res.Msg)
	if err != nil {
		return nil, DataImpairmentFault{cometfaults.ErrMessageFromSchema{SignalLabel: "REDACTED", Err: err}}
	}
	tMessageJournal := &ScheduledJournalSignal{
		Time: res.Time,
		Msg:  journalMessage,
	}

	return tMessageJournal, err
}

type nullJournal struct{}

var _ WAL = nullJournal{}

func (nullJournal) Record(JournalSignal) error     { return nil }
func (nullJournal) RecordAlign(JournalSignal) error { return nil }
func (nullJournal) PurgeAndAlign() error        { return nil }
func (nullJournal) ScanForTerminateLevel(int64, *JournalScanSettings) (rd io.ReadCloser, located bool, err error) {
	return nil, false, nil
}
func (nullJournal) Begin() error { return nil }
func (nullJournal) Halt() error  { return nil }
func (nullJournal) Wait()        {}
