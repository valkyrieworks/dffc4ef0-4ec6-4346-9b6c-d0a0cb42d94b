package agreement

import (
	"bytes"
	"crypto/rand"
	"os"
	"path/filepath"

	//
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/agreement/kinds"
	"github.com/valkyrieworks/vault/merkle"
	"github.com/valkyrieworks/utils/autofile"
	"github.com/valkyrieworks/utils/log"
	cometkinds "github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

const (
	journalVerifyPurgeCadence = time.Duration(100) * time.Millisecond
)

func VerifyJournalClip(t *testing.T) {
	journalFolder, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(journalFolder)

	journalEntry := filepath.Join(journalFolder, "REDACTED")

	//
	//
	//
	//
	wal, err := NewJournal(journalEntry,
		autofile.ClusterFrontVolumeCeiling(4096),
		autofile.ClusterInspectPeriod(1*time.Millisecond),
	)
	require.NoError(t, err)
	wal.AssignTracer(log.VerifyingTracer())
	err = wal.Begin()
	require.NoError(t, err)
	defer func() {
		if err := wal.Halt(); err != nil {
			t.Error(err)
		}
		//
		//
		wal.Wait()
	}()

	//
	//
	//
	err = JournalComposeNLedgers(t, wal.Cluster(), 60, fetchSettings(t))
	require.NoError(t, err)

	time.Sleep(1 * time.Millisecond) //

	if err := wal.PurgeAndAlign(); err != nil {
		t.Error(err)
	}

	h := int64(50)
	gr, located, err := wal.ScanForTerminateLevel(h, &JournalScanSettings{})
	assert.NoError(t, err, "REDACTED", h)
	assert.True(t, located, "REDACTED", h)
	assert.NotNil(t, gr)
	defer gr.Close()

	dec := NewJournalParser(gr)
	msg, err := dec.Parse()
	assert.NoError(t, err, "REDACTED")
	rs, ok := msg.Msg.(cometkinds.EventDataDurationStatus)
	assert.True(t, ok, "REDACTED")
	assert.Equal(t, rs.Level, h+1, "REDACTED")
}

func VerifyJournalSerializerParser(t *testing.T) {
	now := engineclock.Now()
	notices := []ScheduledJournalSignal{
		{Time: now, Msg: TerminateLevelSignal{0}},
		{Time: now, Msg: deadlineDetails{Period: time.Second, Level: 1, Cycle: 1, Phase: kinds.DurationPhaseNominate}},
		{Time: now, Msg: cometkinds.EventDataDurationStatus{Level: 1, Cycle: 1, Phase: "REDACTED"}},
	}

	b := new(bytes.Buffer)

	for _, msg := range notices {

		b.Reset()

		enc := NewJournalSerializer(b)
		err := enc.Serialize(&msg)
		require.NoError(t, err)

		dec := NewJournalParser(b)
		parsed, err := dec.Parse()
		require.NoError(t, err)
		assert.Equal(t, msg.Time.UTC(), parsed.Time)
		assert.Equal(t, msg.Msg, parsed.Msg)
	}
}

func VerifyJournalRecord(t *testing.T) {
	journalFolder, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(journalFolder)
	journalEntry := filepath.Join(journalFolder, "REDACTED")

	wal, err := NewJournal(journalEntry)
	require.NoError(t, err)
	err = wal.Begin()
	require.NoError(t, err)
	defer func() {
		if err := wal.Halt(); err != nil {
			t.Error(err)
		}
		//
		//
		wal.Wait()
	}()

	//
	msg := &LedgerSegmentSignal{
		Level: 1,
		Cycle:  1,
		Segment: &cometkinds.Segment{
			Ordinal: 1,
			Octets: make([]byte, 1),
			Attestation: merkle.Attestation{
				Sum:    1,
				Ordinal:    1,
				NodeDigest: make([]byte, maximumMessageVolumeOctets-30),
			},
		},
	}

	err = wal.Record(messageDetails{
		Msg: msg,
	})
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "REDACTED")
	}
}

func VerifyJournalScanForTerminateLevel(t *testing.T) {
	journalContent, err := JournalWithNLedgers(t, 6, fetchSettings(t))
	if err != nil {
		t.Fatal(err)
	}
	journalEntry := tempJournalWithData(journalContent)

	wal, err := NewJournal(journalEntry)
	require.NoError(t, err)
	wal.AssignTracer(log.VerifyingTracer())

	h := int64(3)
	gr, located, err := wal.ScanForTerminateLevel(h, &JournalScanSettings{})
	assert.NoError(t, err, "REDACTED", h)
	assert.True(t, located, "REDACTED", h)
	assert.NotNil(t, gr)
	defer gr.Close()

	dec := NewJournalParser(gr)
	msg, err := dec.Parse()
	assert.NoError(t, err, "REDACTED")
	rs, ok := msg.Msg.(cometkinds.EventDataDurationStatus)
	assert.True(t, ok, "REDACTED")
	assert.Equal(t, rs.Level, h+1, "REDACTED")
}

func VerifyJournalIntermittentAlign(t *testing.T) {
	journalFolder, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(journalFolder)

	journalEntry := filepath.Join(journalFolder, "REDACTED")
	wal, err := NewJournal(journalEntry, autofile.ClusterInspectPeriod(1*time.Millisecond))
	require.NoError(t, err)

	wal.AssignPurgeCadence(journalVerifyPurgeCadence)
	wal.AssignTracer(log.VerifyingTracer())

	//
	err = JournalComposeNLedgers(t, wal.Cluster(), 5, fetchSettings(t))
	require.NoError(t, err)

	//
	assert.NotZero(t, wal.Cluster().Cached())

	require.NoError(t, wal.Begin())
	defer func() {
		if err := wal.Halt(); err != nil {
			t.Error(err)
		}
		wal.Wait()
	}()

	time.Sleep(journalVerifyPurgeCadence + (10 * time.Millisecond))

	//
	assert.Zero(t, wal.Cluster().Cached())

	h := int64(4)
	gr, located, err := wal.ScanForTerminateLevel(h, &JournalScanSettings{})
	assert.NoError(t, err, "REDACTED", h)
	assert.True(t, located, "REDACTED", h)
	assert.NotNil(t, gr)
	if gr != nil {
		gr.Close()
	}
}

/**
e

{
{
(
,
,
)
)
}
*/

func nOctets(n int) []byte {
	buf := make([]byte, n)
	n, _ = rand.Read(buf)
	return buf[:n]
}

func criterionJournalParse(b *testing.B, n int) {
	//

	buf := new(bytes.Buffer)
	enc := NewJournalSerializer(buf)

	data := nOctets(n)
	if err := enc.Serialize(&ScheduledJournalSignal{Msg: data, Time: time.Now().Round(time.Second).UTC()}); err != nil {
		b.Error(err)
	}

	serialized := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(serialized)
		dec := NewJournalParser(buf)
		if _, err := dec.Parse(); err != nil {
			b.Fatal(err)
		}
	}
	b.ReportAllocs()
}

func CriterionJournalParse512b(b *testing.B) {
	criterionJournalParse(b, 512)
}

func CriterionJournalParse10kb(b *testing.B) {
	criterionJournalParse(b, 10*1024)
}

func CriterionJournalParse100kb(b *testing.B) {
	criterionJournalParse(b, 100*1024)
}

func CriterionJournalParse1mb(b *testing.B) {
	criterionJournalParse(b, 1024*1024)
}

func CriterionJournalParse10mb(b *testing.B) {
	criterionJournalParse(b, 10*1024*1024)
}

func CriterionJournalParse100mb(b *testing.B) {
	criterionJournalParse(b, 100*1024*1024)
}

func CriterionJournalParse1gb(b *testing.B) {
	criterionJournalParse(b, 1024*1024*1024)
}
