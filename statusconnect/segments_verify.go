package statusconnect

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/p2p"
)

func configureSegmentBuffer(t *testing.T) (*segmentBuffer, func()) {
	mirror := &mirror{
		Level:   3,
		Layout:   1,
		Segments:   5,
		Digest:     []byte{7},
		Metainfo: nil,
	}
	buffer, err := newSegmentBuffer(mirror, "REDACTED")
	require.NoError(t, err)
	shutdown := func() {
		err := buffer.End()
		require.NoError(t, err)
	}
	return buffer, shutdown
}

func Verifysegmentqueue_Tempfolder(t *testing.T) {
	mirror := &mirror{
		Level:   3,
		Layout:   1,
		Segments:   5,
		Digest:     []byte{7},
		Metainfo: nil,
	}
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(dir)
	buffer, err := newSegmentBuffer(mirror, dir)
	require.NoError(t, err)

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	assert.Len(t, entries, 1)

	err = buffer.End()
	require.NoError(t, err)

	entries, err = os.ReadDir(dir)
	require.NoError(t, err)
	assert.Len(t, entries, 0)
}

func VerifySegmentBuffer(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	//
	appended, err := buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.True(t, appended)

	//
	appended, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 4, Segment: []byte{3, 1, 4}})
	require.NoError(t, err)
	assert.True(t, appended)

	//
	appended, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.False(t, appended)

	appended, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 4, Segment: []byte{3, 1, 4}})
	require.NoError(t, err)
	assert.False(t, appended)

	//
	appended, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 3, Segment: []byte{3, 1, 3}})
	require.NoError(t, err)
	assert.True(t, appended)

	appended, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 2, Segment: []byte{3, 1, 2}})
	require.NoError(t, err)
	assert.True(t, appended)

	appended, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 1, Segment: []byte{3, 1, 1}})
	require.NoError(t, err)
	assert.True(t, appended)

	//
	for i := 0; i < 5; i++ {
		c, err := buffer.Following()
		require.NoError(t, err)
		assert.Equal(t, &segment{Level: 3, Layout: 1, Ordinal: uint32(i), Segment: []byte{3, 1, byte(i)}}, c)
	}
	_, err = buffer.Following()
	require.Error(t, err)
	assert.Equal(t, errDone, err)

	//
	appended, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.False(t, appended)

	//
	err = buffer.End()
	require.NoError(t, err)
	appended, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.False(t, appended)

	//
	err = buffer.End()
	require.NoError(t, err)
}

func Verifysegmentqueue_Append_Segmentfaults(t *testing.T) {
	verifyscenarios := map[string]struct {
		segment *segment
	}{
		"REDACTED":     {nil},
		"REDACTED":      {&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: nil}},
		"REDACTED":  {&segment{Level: 9, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}}},
		"REDACTED":  {&segment{Level: 3, Layout: 9, Ordinal: 0, Segment: []byte{3, 1, 0}}},
		"REDACTED": {&segment{Level: 3, Layout: 1, Ordinal: 5, Segment: []byte{3, 1, 0}}},
	}
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			buffer, shutdown := configureSegmentBuffer(t)
			defer shutdown()
			_, err := buffer.Add(tc.segment)
			require.Error(t, err)
		})
	}
}

func Verifysegmentqueue_Assign(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	for i := uint32(0); i < buffer.Volume(); i++ {
		ordinal, err := buffer.Assign()
		require.NoError(t, err)
		assert.EqualValues(t, i, ordinal)
	}

	_, err := buffer.Assign()
	require.Error(t, err)
	assert.Equal(t, errDone, err)

	for i := uint32(0); i < buffer.Volume(); i++ {
		_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: i, Segment: []byte{byte(i)}})
		require.NoError(t, err)
	}

	//
	err = buffer.Drop(2)
	require.NoError(t, err)

	ordinal, err := buffer.Assign()
	require.NoError(t, err)
	assert.EqualValues(t, 2, ordinal)
	_, err = buffer.Assign()
	require.Error(t, err)
	assert.Equal(t, errDone, err)

	//
	err = buffer.Drop(2)
	require.NoError(t, err)
	err = buffer.End()
	require.NoError(t, err)
	_, err = buffer.Assign()
	require.Error(t, err)
	assert.Equal(t, errDone, err)
}

func Verifysegmentqueue_Drop(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	//
	_, err := buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: []byte{byte(0)}})
	require.NoError(t, err)
	_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 1, Segment: []byte{byte(1)}})
	require.NoError(t, err)
	_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 2, Segment: []byte{byte(2)}})
	require.NoError(t, err)

	c, err := buffer.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 0, c.Ordinal)
	c, err = buffer.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 1, c.Ordinal)

	//
	//
	err = buffer.Drop(0)
	require.NoError(t, err)
	appended, err := buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: []byte{byte(0)}})
	require.NoError(t, err)
	assert.True(t, appended)
	c, err = buffer.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 0, c.Ordinal)
	c, err = buffer.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 2, c.Ordinal)

	//
	for i := uint32(0); i < buffer.Volume(); i++ {
		err := buffer.Drop(i)
		require.NoError(t, err)
	}
	for i := uint32(0); i < buffer.Volume(); i++ {
		_, err := buffer.Assign()
		require.NoError(t, err)
		_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: i, Segment: []byte{byte(i)}})
		require.NoError(t, err)
		c, err = buffer.Following()
		require.NoError(t, err)
		assert.EqualValues(t, i, c.Ordinal)
	}

	//
	err = buffer.Drop(99)
	require.NoError(t, err)

	//
	err = buffer.Drop(3)
	require.NoError(t, err)
	err = buffer.Drop(1)
	require.NoError(t, err)

	ordinal, err := buffer.Assign()
	require.NoError(t, err)
	assert.EqualValues(t, 1, ordinal)
	ordinal, err = buffer.Assign()
	require.NoError(t, err)
	assert.EqualValues(t, 3, ordinal)

	appended, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 3, Segment: []byte{3}})
	require.NoError(t, err)
	assert.True(t, appended)
	appended, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 1, Segment: []byte{1}})
	require.NoError(t, err)
	assert.True(t, appended)

	segment, err := buffer.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 1, segment.Ordinal)

	segment, err = buffer.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 3, segment.Ordinal)

	_, err = buffer.Following()
	require.Error(t, err)
	assert.Equal(t, errDone, err)

	//
	err = buffer.End()
	require.NoError(t, err)
	err = buffer.Drop(2)
	require.NoError(t, err)
}

func Verifysegmentqueue_Dropeitter(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	//
	emitters := []p2p.ID{"REDACTED", "REDACTED", "REDACTED"}
	for i := uint32(0); i < buffer.Volume(); i++ {
		_, err := buffer.Assign()
		require.NoError(t, err)
		_, err = buffer.Add(&segment{
			Level: 3,
			Layout: 1,
			Ordinal:  i,
			Segment:  []byte{byte(i)},
			Emitter: emitters[int(i)%len(emitters)],
		})
		require.NoError(t, err)
	}

	//
	for i := uint32(0); i < 3; i++ {
		_, err := buffer.Following()
		require.NoError(t, err)
	}

	//
	err := buffer.DropEmitter("REDACTED")
	require.NoError(t, err)
	_, err = buffer.Assign()
	assert.Equal(t, errDone, err)

	//
	//
	err = buffer.DropEmitter("REDACTED")
	require.NoError(t, err)
	ordinal, err := buffer.Assign()
	require.NoError(t, err)
	assert.EqualValues(t, 4, ordinal)
	_, err = buffer.Assign()
	assert.Equal(t, errDone, err)
}

func Verifysegmentqueue_Getemitter(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	_, err := buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: []byte{1}, Emitter: p2p.ID("REDACTED")})
	require.NoError(t, err)
	_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 1, Segment: []byte{2}, Emitter: p2p.ID("REDACTED")})
	require.NoError(t, err)

	assert.EqualValues(t, "REDACTED", buffer.FetchEmitter(0))
	assert.EqualValues(t, "REDACTED", buffer.FetchEmitter(1))
	assert.EqualValues(t, "REDACTED", buffer.FetchEmitter(2))

	//
	segment, err := buffer.Following()
	require.NoError(t, err)
	require.NotNil(t, segment)
	require.EqualValues(t, 0, segment.Ordinal)
	assert.EqualValues(t, "REDACTED", buffer.FetchEmitter(0))
}

func Verifysegmentqueue_Following(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	//
	chanFollowing := make(chan *segment, 10)
	go func() {
		for {
			c, err := buffer.Following()
			if err == errDone {
				close(chanFollowing)
				break
			}
			require.NoError(t, err)
			chanFollowing <- c
		}
	}()

	assert.Empty(t, chanFollowing)
	_, err := buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 1, Segment: []byte{3, 1, 1}, Emitter: p2p.ID("REDACTED")})
	require.NoError(t, err)
	select {
	case <-chanFollowing:
		assert.Fail(t, "REDACTED")
	default:
	}

	_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}, Emitter: p2p.ID("REDACTED")})
	require.NoError(t, err)

	assert.Equal(t,
		&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}, Emitter: p2p.ID("REDACTED")},
		<-chanFollowing)
	assert.Equal(t,
		&segment{Level: 3, Layout: 1, Ordinal: 1, Segment: []byte{3, 1, 1}, Emitter: p2p.ID("REDACTED")},
		<-chanFollowing)

	_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 4, Segment: []byte{3, 1, 4}, Emitter: p2p.ID("REDACTED")})
	require.NoError(t, err)
	select {
	case <-chanFollowing:
		assert.Fail(t, "REDACTED")
	default:
	}

	_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 2, Segment: []byte{3, 1, 2}, Emitter: p2p.ID("REDACTED")})
	require.NoError(t, err)
	_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 3, Segment: []byte{3, 1, 3}, Emitter: p2p.ID("REDACTED")})
	require.NoError(t, err)

	assert.Equal(t,
		&segment{Level: 3, Layout: 1, Ordinal: 2, Segment: []byte{3, 1, 2}, Emitter: p2p.ID("REDACTED")},
		<-chanFollowing)
	assert.Equal(t,
		&segment{Level: 3, Layout: 1, Ordinal: 3, Segment: []byte{3, 1, 3}, Emitter: p2p.ID("REDACTED")},
		<-chanFollowing)
	assert.Equal(t,
		&segment{Level: 3, Layout: 1, Ordinal: 4, Segment: []byte{3, 1, 4}, Emitter: p2p.ID("REDACTED")},
		<-chanFollowing)

	_, ok := <-chanFollowing
	assert.False(t, ok, "REDACTED")

	//
	_, err = buffer.Following()
	assert.Equal(t, errDone, err)
}

func Verifysegmentqueue_Following_Halted(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	//
	_, err := buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 1, Segment: []byte{3, 1, 1}})
	require.NoError(t, err)
	err = buffer.End()
	require.NoError(t, err)

	_, err = buffer.Following()
	assert.Equal(t, errDone, err)
}

func Verifysegmentqueue_Reprocess(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	//
	for i := uint32(0); i < buffer.Volume(); i++ {
		_, err := buffer.Assign()
		require.NoError(t, err)
		_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: i, Segment: []byte{byte(i)}})
		require.NoError(t, err)
		_, err = buffer.Following()
		require.NoError(t, err)
	}

	//
	buffer.Reprocess(3)
	buffer.Reprocess(1)

	_, err := buffer.Assign()
	assert.Equal(t, errDone, err)

	segment, err := buffer.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 1, segment.Ordinal)

	segment, err = buffer.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 3, segment.Ordinal)

	_, err = buffer.Following()
	assert.Equal(t, errDone, err)
}

func Verifysegmentqueue_Reprocessall(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	//
	for i := uint32(0); i < buffer.Volume(); i++ {
		_, err := buffer.Assign()
		require.NoError(t, err)
		_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: i, Segment: []byte{byte(i)}})
		require.NoError(t, err)
		_, err = buffer.Following()
		require.NoError(t, err)
	}

	_, err := buffer.Following()
	assert.Equal(t, errDone, err)

	buffer.ReprocessAll()

	_, err = buffer.Assign()
	assert.Equal(t, errDone, err)

	for i := uint32(0); i < buffer.Volume(); i++ {
		segment, err := buffer.Following()
		require.NoError(t, err)
		assert.EqualValues(t, i, segment.Ordinal)
	}

	_, err = buffer.Following()
	assert.Equal(t, errDone, err)
}

func Verifysegmentqueue_Volume(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	assert.EqualValues(t, 5, buffer.Volume())

	err := buffer.End()
	require.NoError(t, err)
	assert.EqualValues(t, 0, buffer.Volume())
}

func Verifysegmentqueue_Awaitfor(t *testing.T) {
	buffer, shutdown := configureSegmentBuffer(t)
	defer shutdown()

	waitForone := buffer.WaitFor(1)
	waitForfour := buffer.WaitFor(4)

	//
	_, err := buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}})
	require.NoError(t, err)
	_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 2, Segment: []byte{3, 1, 2}})
	require.NoError(t, err)
	select {
	case <-waitForone:
		require.Fail(t, "REDACTED")
	case <-waitForfour:
		require.Fail(t, "REDACTED")
	default:
	}

	//
	_, err = buffer.Add(&segment{Level: 3, Layout: 1, Ordinal: 1, Segment: []byte{3, 1, 1}})
	require.NoError(t, err)
	assert.EqualValues(t, 1, <-waitForone)
	_, ok := <-waitForone
	assert.False(t, ok)
	select {
	case <-waitForfour:
		require.Fail(t, "REDACTED")
	default:
	}

	//
	//
	c, err := buffer.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 0, c.Ordinal)

	w := buffer.WaitFor(0)
	assert.EqualValues(t, 0, <-w)
	_, ok = <-w
	assert.False(t, ok)

	w = buffer.WaitFor(1)
	assert.EqualValues(t, 1, <-w)
	_, ok = <-w
	assert.False(t, ok)

	//
	//
	err = buffer.End()
	require.NoError(t, err)
	_, ok = <-waitForfour
	assert.False(t, ok)

	w = buffer.WaitFor(3)
	_, ok = <-w
	assert.False(t, ok)
}
