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

	automatic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/autosave"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	strongmindcons "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/agreement"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

const (
	//
	maximumSignalExtentOctets = maximumSignalExtent + 24

	//
	journalFallbackPurgeDuration = 2 * time.Second
)

//
//

//
type ScheduledJournalSignal struct {
	Moment time.Time  `json:"moment"`
	Msg  JournalSignal `json:"msg"`
}

//
//
type TerminateAltitudeSignal struct {
	Altitude int64 `json:"altitude"`
}

type JournalSignal any

func initialize() {
	strongmindjson.EnrollKind(signalDetails{}, "REDACTED")
	strongmindjson.EnrollKind(deadlineDetails{}, "REDACTED")
	strongmindjson.EnrollKind(TerminateAltitudeSignal{}, "REDACTED")
}

//
//

//
type WAL interface {
	Record(JournalSignal) error
	RecordChronize(JournalSignal) error
	PurgeAlsoChronize() error

	LookupForeachTerminateAltitude(altitude int64, choices *JournalLookupChoices) (rd io.ReadCloser, detected bool, err error)

	//
	Initiate() error
	Halt() error
	Pause()
}

//
//
//
//
//
type FoundationJournal struct {
	facility.FoundationFacility

	cluster *automatic.Cluster

	enc *JournalSerializer

	purgeMetronome   *time.Ticker
	purgeDuration time.Duration
}

var _ WAL = &FoundationJournal{}

//
//
func FreshJournal(journalRecord string, clusterChoices ...func(*automatic.Cluster)) (*FoundationJournal, error) {
	err := strongos.AssurePath(filepath.Dir(journalRecord), 0o700)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	cluster, err := automatic.InitiateCluster(journalRecord, clusterChoices...)
	if err != nil {
		return nil, err
	}
	wal := &FoundationJournal{
		cluster:         cluster,
		enc:           FreshJournalSerializer(cluster),
		purgeDuration: journalFallbackPurgeDuration,
	}
	wal.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", wal)
	return wal, nil
}

//
func (wal *FoundationJournal) AssignPurgeDuration(i time.Duration) {
	wal.purgeDuration = i
}

func (wal *FoundationJournal) Cluster() *automatic.Cluster {
	return wal.cluster
}

func (wal *FoundationJournal) AssignTracer(l log.Tracer) {
	wal.Tracer = l
	wal.cluster.AssignTracer(l)
}

func (wal *FoundationJournal) UponInitiate() error {
	extent, err := wal.cluster.Leading.Extent()
	if err != nil {
		return err
	} else if extent == 0 {
		if err := wal.RecordChronize(TerminateAltitudeSignal{0}); err != nil {
			return err
		}
	}
	err = wal.cluster.Initiate()
	if err != nil {
		return err
	}
	wal.purgeMetronome = time.NewTicker(wal.purgeDuration)
	go wal.handlePurgeCounts()
	return nil
}

func (wal *FoundationJournal) handlePurgeCounts() {
	for {
		select {
		case <-wal.purgeMetronome.C:
			if err := wal.PurgeAlsoChronize(); err != nil {
				wal.Tracer.Failure("REDACTED", "REDACTED", err)
			}
		case <-wal.Exit():
			return
		}
	}
}

//
//
func (wal *FoundationJournal) PurgeAlsoChronize() error {
	return wal.cluster.PurgeAlsoChronize()
}

//
//
//
func (wal *FoundationJournal) UponHalt() {
	wal.purgeMetronome.Stop()
	if err := wal.PurgeAlsoChronize(); err != nil {
		wal.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	if err := wal.cluster.Halt(); err != nil {
		wal.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	wal.cluster.Shutdown()
}

//
//
func (wal *FoundationJournal) Pause() {
	wal.cluster.Pause()
}

//
//
//
func (wal *FoundationJournal) Record(msg JournalSignal) error {
	if wal == nil {
		return nil
	}

	if err := wal.enc.Serialize(&ScheduledJournalSignal{committime.Now(), msg}); err != nil {
		wal.Tracer.Failure("REDACTED",
			"REDACTED", err, "REDACTED", msg)
		return err
	}

	return nil
}

//
//
//
func (wal *FoundationJournal) RecordChronize(msg JournalSignal) error {
	if wal == nil {
		return nil
	}

	if err := wal.Record(msg); err != nil {
		return err
	}

	if err := wal.PurgeAlsoChronize(); err != nil {
		wal.Tracer.Failure(`REDACTED.
REDACTED`,
			"REDACTED", err)
		return err
	}

	return nil
}

//
type JournalLookupChoices struct {
	//
	BypassDataImpairmentFaults bool
}

//
//
//
//
//
func (wal *FoundationJournal) LookupForeachTerminateAltitude(
	altitude int64,
	choices *JournalLookupChoices,
) (rd io.ReadCloser, detected bool, err error) {
	var (
		msg *ScheduledJournalSignal
		gr  *automatic.ClusterFetcher
	)
	finalAltitudeDetected := int64(-1)

	//
	//
	min, max := wal.cluster.MinimumOrdinal(), wal.cluster.MaximumOrdinal()
	wal.Tracer.Details("REDACTED", "REDACTED", altitude, "REDACTED", min, "REDACTED", max)
	for ordinal := max; ordinal >= min; ordinal-- {
		gr, err = wal.cluster.FreshFetcher(ordinal)
		if err != nil {
			return nil, false, err
		}

		dec := FreshJournalDeserializer(gr)
		for {
			msg, err = dec.Deserialize()
			if err == io.EOF {
				//
				if finalAltitudeDetected > 0 && finalAltitudeDetected < altitude {
					gr.Shutdown()
					return nil, false, nil
				}
				//
				break
			}
			if choices.BypassDataImpairmentFaults && EqualsDataImpairmentFailure(err) {
				wal.Tracer.Failure("REDACTED", "REDACTED", err)
				//
				continue
			} else if err != nil {
				gr.Shutdown()
				return nil, false, err
			}

			if m, ok := msg.Msg.(TerminateAltitudeSignal); ok {
				finalAltitudeDetected = m.Altitude
				if m.Altitude == altitude { //
					wal.Tracer.Details("REDACTED", "REDACTED", altitude, "REDACTED", ordinal)
					return gr, true, nil
				}
			}
		}
		gr.Shutdown()
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
func FreshJournalSerializer(wr io.Writer) *JournalSerializer {
	return &JournalSerializer{wr}
}

//
//
//
func (enc *JournalSerializer) Serialize(v *ScheduledJournalSignal) error {
	bufferSignal, err := JournalTowardSchema(v.Msg)
	if err != nil {
		return err
	}
	pv := strongmindcons.ScheduledJournalSignal{
		Moment: v.Moment,
		Msg:  bufferSignal,
	}

	data, err := proto.Marshal(&pv)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	crc := crc32.Checksum(data, checksum32c)
	magnitude := uint32(len(data))
	if magnitude > maximumSignalExtentOctets {
		return fmt.Errorf("REDACTED", magnitude, maximumSignalExtentOctets)
	}
	sumMagnitude := 8 + int(magnitude)

	msg := make([]byte, sumMagnitude)
	binary.BigEndian.PutUint32(msg[0:4], crc)
	binary.BigEndian.PutUint32(msg[4:8], magnitude)
	copy(msg[8:], data)

	_, err = enc.wr.Write(msg)
	return err
}

//
func EqualsDataImpairmentFailure(err error) bool {
	_, ok := err.(DataImpairmentFailure)
	return ok
}

//
type DataImpairmentFailure struct {
	reason error
}

func (e DataImpairmentFailure) Failure() string {
	return fmt.Sprintf("REDACTED", e.reason)
}

func (e DataImpairmentFailure) Reason() error {
	return e.reason
}

//
//
//
//
//
type JournalDeserializer struct {
	rd io.Reader
}

//
func FreshJournalDeserializer(rd io.Reader) *JournalDeserializer {
	return &JournalDeserializer{rd}
}

//
func (dec *JournalDeserializer) Deserialize() (*ScheduledJournalSignal, error) {
	b := make([]byte, 4)

	_, err := dec.rd.Read(b)
	if errors.Is(err, io.EOF) {
		return nil, err
	}
	if err != nil {
		return nil, DataImpairmentFailure{fmt.Errorf("REDACTED", err)}
	}
	crc := binary.BigEndian.Uint32(b)

	b = make([]byte, 4)
	_, err = dec.rd.Read(b)
	if err != nil {
		return nil, DataImpairmentFailure{fmt.Errorf("REDACTED", err)}
	}
	magnitude := binary.BigEndian.Uint32(b)

	if magnitude > maximumSignalExtentOctets {
		return nil, DataImpairmentFailure{fmt.Errorf(
			"REDACTED",
			magnitude,
			maximumSignalExtentOctets)}
	}

	data := make([]byte, magnitude)
	n, err := dec.rd.Read(data)
	if err != nil {
		return nil, DataImpairmentFailure{fmt.Errorf("REDACTED", err, n, magnitude)}
	}

	//
	preciseChecksum := crc32.Checksum(data, checksum32c)
	if preciseChecksum != crc {
		return nil, DataImpairmentFailure{fmt.Errorf("REDACTED", crc, preciseChecksum)}
	}

	res := new(strongmindcons.ScheduledJournalSignal)
	err = proto.Unmarshal(data, res)
	if err != nil {
		return nil, DataImpairmentFailure{fmt.Errorf("REDACTED", err)}
	}

	journalSignal, err := JournalOriginatingSchema(res.Msg)
	if err != nil {
		return nil, DataImpairmentFailure{strongminderrors.FaultSignalOriginatingSchema{SignalAlias: "REDACTED", Err: err}}
	}
	tempSignalJournal := &ScheduledJournalSignal{
		Moment: res.Moment,
		Msg:  journalSignal,
	}

	return tempSignalJournal, err
}

type voidJournal struct{}

var _ WAL = voidJournal{}

func (voidJournal) Record(JournalSignal) error     { return nil }
func (voidJournal) RecordChronize(JournalSignal) error { return nil }
func (voidJournal) PurgeAlsoChronize() error        { return nil }
func (voidJournal) LookupForeachTerminateAltitude(int64, *JournalLookupChoices) (rd io.ReadCloser, detected bool, err error) {
	return nil, false, nil
}
func (voidJournal) Initiate() error { return nil }
func (voidJournal) Halt() error  { return nil }
func (voidJournal) Pause()        {}
