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
	"github.com/valkyrieworks/security/hashing"
	"github.com/valkyrieworks/utils/automaticfile"
	"github.com/valkyrieworks/utils/log"
	enginetypes "github.com/valkyrieworks/kinds"
	cttime "github.com/valkyrieworks/kinds/moment"
)

const (
	walTestFlushInterval = time.Duration(100) * time.Millisecond
)

func TestWALTruncate(t *testing.T) {
	walDir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(walDir)

	walFile := filepath.Join(walDir, "REDACTED")

	//
	//
	//
	//
	wal, err := NewWAL(walFile,
		automaticfile.GroupHeadSizeLimit(4096),
		automaticfile.GroupCheckDuration(1*time.Millisecond),
	)
	require.NoError(t, err)
	wal.SetLogger(log.TestingLogger())
	err = wal.Start()
	require.NoError(t, err)
	defer func() {
		if err := wal.Stop(); err != nil {
			t.Error(err)
		}
		//
		//
		wal.Wait()
	}()

	//
	//
	//
	err = WALGenerateNBlocks(t, wal.Group(), 60, getConfig(t))
	require.NoError(t, err)

	time.Sleep(1 * time.Millisecond) //

	if err := wal.FlushAndSync(); err != nil {
		t.Error(err)
	}

	h := int64(50)
	gr, found, err := wal.SearchForEndHeight(h, &WALSearchOptions{})
	assert.NoError(t, err, "REDACTED", h)
	assert.True(t, found, "REDACTED", h)
	assert.NotNil(t, gr)
	defer gr.Close()

	dec := NewWALDecoder(gr)
	msg, err := dec.Decode()
	assert.NoError(t, err, "REDACTED")
	rs, ok := msg.Msg.(enginetypes.EventDataRoundState)
	assert.True(t, ok, "REDACTED")
	assert.Equal(t, rs.Height, h+1, "REDACTED")
}

func TestWALEncoderDecoder(t *testing.T) {
	now := cttime.Now()
	msgs := []TimedWALMessage{
		{Time: now, Msg: EndHeightMessage{0}},
		{Time: now, Msg: timeoutInfo{Duration: time.Second, Height: 1, Round: 1, Step: kinds.RoundStepPropose}},
		{Time: now, Msg: enginetypes.EventDataRoundState{Height: 1, Round: 1, Step: "REDACTED"}},
	}

	b := new(bytes.Buffer)

	for _, msg := range msgs {

		b.Reset()

		enc := NewWALEncoder(b)
		err := enc.Encode(&msg)
		require.NoError(t, err)

		dec := NewWALDecoder(b)
		decoded, err := dec.Decode()
		require.NoError(t, err)
		assert.Equal(t, msg.Time.UTC(), decoded.Time)
		assert.Equal(t, msg.Msg, decoded.Msg)
	}
}

func TestWALWrite(t *testing.T) {
	walDir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(walDir)
	walFile := filepath.Join(walDir, "REDACTED")

	wal, err := NewWAL(walFile)
	require.NoError(t, err)
	err = wal.Start()
	require.NoError(t, err)
	defer func() {
		if err := wal.Stop(); err != nil {
			t.Error(err)
		}
		//
		//
		wal.Wait()
	}()

	//
	msg := &BlockPartMessage{
		Height: 1,
		Round:  1,
		Part: &enginetypes.Part{
			Index: 1,
			Bytes: make([]byte, 1),
			Proof: hashing.Proof{
				Total:    1,
				Index:    1,
				LeafHash: make([]byte, maxMsgSizeBytes-30),
			},
		},
	}

	err = wal.Write(msgInfo{
		Msg: msg,
	})
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "REDACTED")
	}
}

func TestWALSearchForEndHeight(t *testing.T) {
	walBody, err := WALWithNBlocks(t, 6, getConfig(t))
	if err != nil {
		t.Fatal(err)
	}
	walFile := tempWALWithData(walBody)

	wal, err := NewWAL(walFile)
	require.NoError(t, err)
	wal.SetLogger(log.TestingLogger())

	h := int64(3)
	gr, found, err := wal.SearchForEndHeight(h, &WALSearchOptions{})
	assert.NoError(t, err, "REDACTED", h)
	assert.True(t, found, "REDACTED", h)
	assert.NotNil(t, gr)
	defer gr.Close()

	dec := NewWALDecoder(gr)
	msg, err := dec.Decode()
	assert.NoError(t, err, "REDACTED")
	rs, ok := msg.Msg.(enginetypes.EventDataRoundState)
	assert.True(t, ok, "REDACTED")
	assert.Equal(t, rs.Height, h+1, "REDACTED")
}

func TestWALPeriodicSync(t *testing.T) {
	walDir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(walDir)

	walFile := filepath.Join(walDir, "REDACTED")
	wal, err := NewWAL(walFile, automaticfile.GroupCheckDuration(1*time.Millisecond))
	require.NoError(t, err)

	wal.SetFlushInterval(walTestFlushInterval)
	wal.SetLogger(log.TestingLogger())

	//
	err = WALGenerateNBlocks(t, wal.Group(), 5, getConfig(t))
	require.NoError(t, err)

	//
	assert.NotZero(t, wal.Group().Buffered())

	require.NoError(t, wal.Start())
	defer func() {
		if err := wal.Stop(); err != nil {
			t.Error(err)
		}
		wal.Wait()
	}()

	time.Sleep(walTestFlushInterval + (10 * time.Millisecond))

	//
	assert.Zero(t, wal.Group().Buffered())

	h := int64(4)
	gr, found, err := wal.SearchForEndHeight(h, &WALSearchOptions{})
	assert.NoError(t, err, "REDACTED", h)
	assert.True(t, found, "REDACTED", h)
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

func nBytes(n int) []byte {
	buf := make([]byte, n)
	n, _ = rand.Read(buf)
	return buf[:n]
}

func benchmarkWalDecode(b *testing.B, n int) {
	//

	buf := new(bytes.Buffer)
	enc := NewWALEncoder(buf)

	data := nBytes(n)
	if err := enc.Encode(&TimedWALMessage{Msg: data, Time: time.Now().Round(time.Second).UTC()}); err != nil {
		b.Error(err)
	}

	encoded := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(encoded)
		dec := NewWALDecoder(buf)
		if _, err := dec.Decode(); err != nil {
			b.Fatal(err)
		}
	}
	b.ReportAllocs()
}

func BenchmarkWalDecode512B(b *testing.B) {
	benchmarkWalDecode(b, 512)
}

func BenchmarkWalDecode10KB(b *testing.B) {
	benchmarkWalDecode(b, 10*1024)
}

func BenchmarkWalDecode100KB(b *testing.B) {
	benchmarkWalDecode(b, 100*1024)
}

func BenchmarkWalDecode1MB(b *testing.B) {
	benchmarkWalDecode(b, 1024*1024)
}

func BenchmarkWalDecode10MB(b *testing.B) {
	benchmarkWalDecode(b, 10*1024*1024)
}

func BenchmarkWalDecode100MB(b *testing.B) {
	benchmarkWalDecode(b, 100*1024*1024)
}

func BenchmarkWalDecode1GB(b *testing.B) {
	benchmarkWalDecode(b, 1024*1024*1024)
}
