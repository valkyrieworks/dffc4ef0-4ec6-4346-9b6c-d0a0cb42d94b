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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/autosave"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongmindkinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

const (
	journalVerifyPurgeDuration = time.Duration(100) * time.Millisecond
)

func VerifyJournalShorten(t *testing.T) {
	journalPath, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(journalPath)

	journalRecord := filepath.Join(journalPath, "REDACTED")

	//
	//
	//
	//
	wal, err := FreshJournal(journalRecord,
		autosave.ClusterLeadingExtentThreshold(4096),
		autosave.ClusterInspectInterval(1*time.Millisecond),
	)
	require.NoError(t, err)
	wal.AssignTracer(log.VerifyingTracer())
	err = wal.Initiate()
	require.NoError(t, err)
	defer func() {
		if err := wal.Halt(); err != nil {
			t.Error(err)
		}
		//
		//
		wal.Await()
	}()

	//
	//
	//
	err = JournalComposeNTHLedgers(t, wal.Cluster(), 60, obtainSettings(t))
	require.NoError(t, err)

	time.Sleep(1 * time.Millisecond) //

	if err := wal.PurgeAlsoChronize(); err != nil {
		t.Error(err)
	}

	h := int64(50)
	gr, detected, err := wal.LookupForeachTerminateAltitude(h, &JournalLookupChoices{})
	assert.NoError(t, err, "REDACTED", h)
	assert.True(t, detected, "REDACTED", h)
	assert.NotNil(t, gr)
	defer gr.Close()

	dec := FreshJournalDeserializer(gr)
	msg, err := dec.Deserialize()
	assert.NoError(t, err, "REDACTED")
	rs, ok := msg.Msg.(strongmindkinds.IncidentDataIterationStatus)
	assert.True(t, ok, "REDACTED")
	assert.Equal(t, rs.Altitude, h+1, "REDACTED")
}

func VerifyJournalSerializerDeserializer(t *testing.T) {
	now := committime.Now()
	signals := []ScheduledJournalSignal{
		{Moment: now, Msg: TerminateAltitudeSignal{0}},
		{Moment: now, Msg: deadlineDetails{Interval: time.Second, Altitude: 1, Iteration: 1, Phase: kinds.IterationPhaseNominate}},
		{Moment: now, Msg: strongmindkinds.IncidentDataIterationStatus{Altitude: 1, Iteration: 1, Phase: "REDACTED"}},
	}

	b := new(bytes.Buffer)

	for _, msg := range signals {

		b.Reset()

		enc := FreshJournalSerializer(b)
		err := enc.Serialize(&msg)
		require.NoError(t, err)

		dec := FreshJournalDeserializer(b)
		deserialized, err := dec.Deserialize()
		require.NoError(t, err)
		assert.Equal(t, msg.Moment.UTC(), deserialized.Moment)
		assert.Equal(t, msg.Msg, deserialized.Msg)
	}
}

func VerifyJournalPersist(t *testing.T) {
	journalPath, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(journalPath)
	journalRecord := filepath.Join(journalPath, "REDACTED")

	wal, err := FreshJournal(journalRecord)
	require.NoError(t, err)
	err = wal.Initiate()
	require.NoError(t, err)
	defer func() {
		if err := wal.Halt(); err != nil {
			t.Error(err)
		}
		//
		//
		wal.Await()
	}()

	//
	msg := &LedgerFragmentSignal{
		Altitude: 1,
		Iteration:  1,
		Fragment: &strongmindkinds.Fragment{
			Ordinal: 1,
			Octets: make([]byte, 1),
			Attestation: hashmap.Attestation{
				Sum:    1,
				Ordinal:    1,
				NodeDigest: make([]byte, maximumSignalExtentOctets-30),
			},
		},
	}

	err = wal.Persist(signalDetails{
		Msg: msg,
	})
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "REDACTED")
	}
}

func VerifyJournalLookupForeachTerminateAltitude(t *testing.T) {
	journalContent, err := JournalUsingNTHLedgers(t, 6, obtainSettings(t))
	if err != nil {
		t.Fatal(err)
	}
	journalRecord := transientJournalUsingData(journalContent)

	wal, err := FreshJournal(journalRecord)
	require.NoError(t, err)
	wal.AssignTracer(log.VerifyingTracer())

	h := int64(3)
	gr, detected, err := wal.LookupForeachTerminateAltitude(h, &JournalLookupChoices{})
	assert.NoError(t, err, "REDACTED", h)
	assert.True(t, detected, "REDACTED", h)
	assert.NotNil(t, gr)
	defer gr.Close()

	dec := FreshJournalDeserializer(gr)
	msg, err := dec.Deserialize()
	assert.NoError(t, err, "REDACTED")
	rs, ok := msg.Msg.(strongmindkinds.IncidentDataIterationStatus)
	assert.True(t, ok, "REDACTED")
	assert.Equal(t, rs.Altitude, h+1, "REDACTED")
}

func VerifyJournalRecurrentChronize(t *testing.T) {
	journalPath, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(journalPath)

	journalRecord := filepath.Join(journalPath, "REDACTED")
	wal, err := FreshJournal(journalRecord, autosave.ClusterInspectInterval(1*time.Millisecond))
	require.NoError(t, err)

	wal.AssignPurgeDuration(journalVerifyPurgeDuration)
	wal.AssignTracer(log.VerifyingTracer())

	//
	err = JournalComposeNTHLedgers(t, wal.Cluster(), 5, obtainSettings(t))
	require.NoError(t, err)

	//
	assert.NotZero(t, wal.Cluster().Cached())

	require.NoError(t, wal.Initiate())
	defer func() {
		if err := wal.Halt(); err != nil {
			t.Error(err)
		}
		wal.Await()
	}()

	time.Sleep(journalVerifyPurgeDuration + (10 * time.Millisecond))

	//
	assert.Zero(t, wal.Cluster().Cached())

	h := int64(4)
	gr, detected, err := wal.LookupForeachTerminateAltitude(h, &JournalLookupChoices{})
	assert.NoError(t, err, "REDACTED", h)
	assert.True(t, detected, "REDACTED", h)
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

func nthOctets(n int) []byte {
	buf := make([]byte, n)
	n, _ = rand.Read(buf)
	return buf[:n]
}

func assessmentJournalDeserialize(b *testing.B, n int) {
	//

	buf := new(bytes.Buffer)
	enc := FreshJournalSerializer(buf)

	data := nthOctets(n)
	if err := enc.Serialize(&ScheduledJournalSignal{Msg: data, Moment: time.Now().Round(time.Second).UTC()}); err != nil {
		b.Error(err)
	}

	serialized := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(serialized)
		dec := FreshJournalDeserializer(buf)
		if _, err := dec.Deserialize(); err != nil {
			b.Fatal(err)
		}
	}
	b.ReportAllocs()
}

func AssessmentJournalDeserialize512b(b *testing.B) {
	assessmentJournalDeserialize(b, 512)
}

func AssessmentJournalDeserialize10kb(b *testing.B) {
	assessmentJournalDeserialize(b, 10*1024)
}

func AssessmentJournalDeserialize100kb(b *testing.B) {
	assessmentJournalDeserialize(b, 100*1024)
}

func AssessmentJournalDeserialize1mb(b *testing.B) {
	assessmentJournalDeserialize(b, 1024*1024)
}

func AssessmentJournalDeserialize10mb(b *testing.B) {
	assessmentJournalDeserialize(b, 10*1024*1024)
}

func AssessmentJournalDeserialize100mb(b *testing.B) {
	assessmentJournalDeserialize(b, 100*1024*1024)
}

func AssessmentJournalDeserialize1gb(b *testing.B) {
	assessmentJournalDeserialize(b, 1024*1024*1024)
}
