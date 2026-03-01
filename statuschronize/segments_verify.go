package statuschronize

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

func configureSegmentStaging(t *testing.T) (*segmentStaging, func()) {
	image := &image{
		Altitude:   3,
		Layout:   1,
		Segments:   5,
		Digest:     []byte{7},
		Attributes: nil,
	}
	staging, err := freshSegmentStaging(image, "REDACTED")
	require.NoError(t, err)
	deconfigure := func() {
		err := staging.Shutdown()
		require.NoError(t, err)
	}
	return staging, deconfigure
}

func Verifysegmentqueue_Temppath(t *testing.T) {
	image := &image{
		Altitude:   3,
		Layout:   1,
		Segments:   5,
		Digest:     []byte{7},
		Attributes: nil,
	}
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(dir)
	staging, err := freshSegmentStaging(image, dir)
	require.NoError(t, err)

	records, err := os.ReadDir(dir)
	require.NoError(t, err)
	assert.Len(t, records, 1)

	err = staging.Shutdown()
	require.NoError(t, err)

	records, err = os.ReadDir(dir)
	require.NoError(t, err)
	assert.Len(t, records, 0)
}

func VerifySegmentStaging(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	//
	appended, err := staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.True(t, appended)

	//
	appended, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 4, Segment: []byte{3, 1, 4}})
	require.NoError(t, err)
	assert.True(t, appended)

	//
	appended, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.False(t, appended)

	appended, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 4, Segment: []byte{3, 1, 4}})
	require.NoError(t, err)
	assert.False(t, appended)

	//
	appended, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 3, Segment: []byte{3, 1, 3}})
	require.NoError(t, err)
	assert.True(t, appended)

	appended, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 2, Segment: []byte{3, 1, 2}})
	require.NoError(t, err)
	assert.True(t, appended)

	appended, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 1, Segment: []byte{3, 1, 1}})
	require.NoError(t, err)
	assert.True(t, appended)

	//
	for i := 0; i < 5; i++ {
		c, err := staging.Following()
		require.NoError(t, err)
		assert.Equal(t, &segment{Altitude: 3, Layout: 1, Ordinal: uint32(i), Segment: []byte{3, 1, byte(i)}}, c)
	}
	_, err = staging.Following()
	require.Error(t, err)
	assert.Equal(t, faultComplete, err)

	//
	appended, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.False(t, appended)

	//
	err = staging.Shutdown()
	require.NoError(t, err)
	appended, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.False(t, appended)

	//
	err = staging.Shutdown()
	require.NoError(t, err)
}

func Verifysegmentqueue_Append_Segmentfaults(t *testing.T) {
	verifycases := map[string]struct {
		segment *segment
	}{
		"REDACTED":     {nil},
		"REDACTED":      {&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: nil}},
		"REDACTED":  {&segment{Altitude: 9, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}}},
		"REDACTED":  {&segment{Altitude: 3, Layout: 9, Ordinal: 0, Segment: []byte{3, 1, 0}}},
		"REDACTED": {&segment{Altitude: 3, Layout: 1, Ordinal: 5, Segment: []byte{3, 1, 0}}},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			staging, deconfigure := configureSegmentStaging(t)
			defer deconfigure()
			_, err := staging.Add(tc.segment)
			require.Error(t, err)
		})
	}
}

func Verifysegmentqueue_Assign(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	for i := uint32(0); i < staging.Extent(); i++ {
		ordinal, err := staging.Assign()
		require.NoError(t, err)
		assert.EqualValues(t, i, ordinal)
	}

	_, err := staging.Assign()
	require.Error(t, err)
	assert.Equal(t, faultComplete, err)

	for i := uint32(0); i < staging.Extent(); i++ {
		_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: i, Segment: []byte{byte(i)}})
		require.NoError(t, err)
	}

	//
	err = staging.Eject(2)
	require.NoError(t, err)

	ordinal, err := staging.Assign()
	require.NoError(t, err)
	assert.EqualValues(t, 2, ordinal)
	_, err = staging.Assign()
	require.Error(t, err)
	assert.Equal(t, faultComplete, err)

	//
	err = staging.Eject(2)
	require.NoError(t, err)
	err = staging.Shutdown()
	require.NoError(t, err)
	_, err = staging.Assign()
	require.Error(t, err)
	assert.Equal(t, faultComplete, err)
}

func Verifysegmentqueue_Eject(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	//
	_, err := staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: []byte{byte(0)}})
	require.NoError(t, err)
	_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 1, Segment: []byte{byte(1)}})
	require.NoError(t, err)
	_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 2, Segment: []byte{byte(2)}})
	require.NoError(t, err)

	c, err := staging.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 0, c.Ordinal)
	c, err = staging.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 1, c.Ordinal)

	//
	//
	err = staging.Eject(0)
	require.NoError(t, err)
	appended, err := staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: []byte{byte(0)}})
	require.NoError(t, err)
	assert.True(t, appended)
	c, err = staging.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 0, c.Ordinal)
	c, err = staging.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 2, c.Ordinal)

	//
	for i := uint32(0); i < staging.Extent(); i++ {
		err := staging.Eject(i)
		require.NoError(t, err)
	}
	for i := uint32(0); i < staging.Extent(); i++ {
		_, err := staging.Assign()
		require.NoError(t, err)
		_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: i, Segment: []byte{byte(i)}})
		require.NoError(t, err)
		c, err = staging.Following()
		require.NoError(t, err)
		assert.EqualValues(t, i, c.Ordinal)
	}

	//
	err = staging.Eject(99)
	require.NoError(t, err)

	//
	err = staging.Eject(3)
	require.NoError(t, err)
	err = staging.Eject(1)
	require.NoError(t, err)

	ordinal, err := staging.Assign()
	require.NoError(t, err)
	assert.EqualValues(t, 1, ordinal)
	ordinal, err = staging.Assign()
	require.NoError(t, err)
	assert.EqualValues(t, 3, ordinal)

	appended, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 3, Segment: []byte{3}})
	require.NoError(t, err)
	assert.True(t, appended)
	appended, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 1, Segment: []byte{1}})
	require.NoError(t, err)
	assert.True(t, appended)

	segment, err := staging.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 1, segment.Ordinal)

	segment, err = staging.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 3, segment.Ordinal)

	_, err = staging.Following()
	require.Error(t, err)
	assert.Equal(t, faultComplete, err)

	//
	err = staging.Shutdown()
	require.NoError(t, err)
	err = staging.Eject(2)
	require.NoError(t, err)
}

func Verifysegmentqueue_Discardsender(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	//
	originators := []p2p.ID{"REDACTED", "REDACTED", "REDACTED"}
	for i := uint32(0); i < staging.Extent(); i++ {
		_, err := staging.Assign()
		require.NoError(t, err)
		_, err = staging.Add(&segment{
			Altitude: 3,
			Layout: 1,
			Ordinal:  i,
			Segment:  []byte{byte(i)},
			Originator: originators[int(i)%len(originators)],
		})
		require.NoError(t, err)
	}

	//
	for i := uint32(0); i < 3; i++ {
		_, err := staging.Following()
		require.NoError(t, err)
	}

	//
	err := staging.EjectOriginator("REDACTED")
	require.NoError(t, err)
	_, err = staging.Assign()
	assert.Equal(t, faultComplete, err)

	//
	//
	err = staging.EjectOriginator("REDACTED")
	require.NoError(t, err)
	ordinal, err := staging.Assign()
	require.NoError(t, err)
	assert.EqualValues(t, 4, ordinal)
	_, err = staging.Assign()
	assert.Equal(t, faultComplete, err)
}

func Verifysegmentqueue_Getoriginator(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	_, err := staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: []byte{1}, Originator: p2p.ID("REDACTED")})
	require.NoError(t, err)
	_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 1, Segment: []byte{2}, Originator: p2p.ID("REDACTED")})
	require.NoError(t, err)

	assert.EqualValues(t, "REDACTED", staging.ObtainOriginator(0))
	assert.EqualValues(t, "REDACTED", staging.ObtainOriginator(1))
	assert.EqualValues(t, "REDACTED", staging.ObtainOriginator(2))

	//
	segment, err := staging.Following()
	require.NoError(t, err)
	require.NotNil(t, segment)
	require.EqualValues(t, 0, segment.Ordinal)
	assert.EqualValues(t, "REDACTED", staging.ObtainOriginator(0))
}

func Verifysegmentqueue_Following(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	//
	streamFollowing := make(chan *segment, 10)
	go func() {
		for {
			c, err := staging.Following()
			if err == faultComplete {
				close(streamFollowing)
				break
			}
			require.NoError(t, err)
			streamFollowing <- c
		}
	}()

	assert.Empty(t, streamFollowing)
	_, err := staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 1, Segment: []byte{3, 1, 1}, Originator: p2p.ID("REDACTED")})
	require.NoError(t, err)
	select {
	case <-streamFollowing:
		assert.Fail(t, "REDACTED")
	default:
	}

	_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}, Originator: p2p.ID("REDACTED")})
	require.NoError(t, err)

	assert.Equal(t,
		&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}, Originator: p2p.ID("REDACTED")},
		<-streamFollowing)
	assert.Equal(t,
		&segment{Altitude: 3, Layout: 1, Ordinal: 1, Segment: []byte{3, 1, 1}, Originator: p2p.ID("REDACTED")},
		<-streamFollowing)

	_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 4, Segment: []byte{3, 1, 4}, Originator: p2p.ID("REDACTED")})
	require.NoError(t, err)
	select {
	case <-streamFollowing:
		assert.Fail(t, "REDACTED")
	default:
	}

	_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 2, Segment: []byte{3, 1, 2}, Originator: p2p.ID("REDACTED")})
	require.NoError(t, err)
	_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 3, Segment: []byte{3, 1, 3}, Originator: p2p.ID("REDACTED")})
	require.NoError(t, err)

	assert.Equal(t,
		&segment{Altitude: 3, Layout: 1, Ordinal: 2, Segment: []byte{3, 1, 2}, Originator: p2p.ID("REDACTED")},
		<-streamFollowing)
	assert.Equal(t,
		&segment{Altitude: 3, Layout: 1, Ordinal: 3, Segment: []byte{3, 1, 3}, Originator: p2p.ID("REDACTED")},
		<-streamFollowing)
	assert.Equal(t,
		&segment{Altitude: 3, Layout: 1, Ordinal: 4, Segment: []byte{3, 1, 4}, Originator: p2p.ID("REDACTED")},
		<-streamFollowing)

	_, ok := <-streamFollowing
	assert.False(t, ok, "REDACTED")

	//
	_, err = staging.Following()
	assert.Equal(t, faultComplete, err)
}

func Verifysegmentqueue_Following_Shutdown(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	//
	_, err := staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 1, Segment: []byte{3, 1, 1}})
	require.NoError(t, err)
	err = staging.Shutdown()
	require.NoError(t, err)

	_, err = staging.Following()
	assert.Equal(t, faultComplete, err)
}

func Verifysegmentqueue_Reissue(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	//
	for i := uint32(0); i < staging.Extent(); i++ {
		_, err := staging.Assign()
		require.NoError(t, err)
		_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: i, Segment: []byte{byte(i)}})
		require.NoError(t, err)
		_, err = staging.Following()
		require.NoError(t, err)
	}

	//
	staging.Reissue(3)
	staging.Reissue(1)

	_, err := staging.Assign()
	assert.Equal(t, faultComplete, err)

	segment, err := staging.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 1, segment.Ordinal)

	segment, err = staging.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 3, segment.Ordinal)

	_, err = staging.Following()
	assert.Equal(t, faultComplete, err)
}

func Verifysegmentqueue_Reissueall(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	//
	for i := uint32(0); i < staging.Extent(); i++ {
		_, err := staging.Assign()
		require.NoError(t, err)
		_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: i, Segment: []byte{byte(i)}})
		require.NoError(t, err)
		_, err = staging.Following()
		require.NoError(t, err)
	}

	_, err := staging.Following()
	assert.Equal(t, faultComplete, err)

	staging.ReissueEvery()

	_, err = staging.Assign()
	assert.Equal(t, faultComplete, err)

	for i := uint32(0); i < staging.Extent(); i++ {
		segment, err := staging.Following()
		require.NoError(t, err)
		assert.EqualValues(t, i, segment.Ordinal)
	}

	_, err = staging.Following()
	assert.Equal(t, faultComplete, err)
}

func Verifysegmentqueue_Extent(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	assert.EqualValues(t, 5, staging.Extent())

	err := staging.Shutdown()
	require.NoError(t, err)
	assert.EqualValues(t, 0, staging.Extent())
}

func Verifysegmentqueue_Await(t *testing.T) {
	staging, deconfigure := configureSegmentStaging(t)
	defer deconfigure()

	awaitForeach1 := staging.AwaitForeach(1)
	awaitForeach4 := staging.AwaitForeach(4)

	//
	_, err := staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 0, Segment: []byte{3, 1, 0}})
	require.NoError(t, err)
	_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 2, Segment: []byte{3, 1, 2}})
	require.NoError(t, err)
	select {
	case <-awaitForeach1:
		require.Fail(t, "REDACTED")
	case <-awaitForeach4:
		require.Fail(t, "REDACTED")
	default:
	}

	//
	_, err = staging.Add(&segment{Altitude: 3, Layout: 1, Ordinal: 1, Segment: []byte{3, 1, 1}})
	require.NoError(t, err)
	assert.EqualValues(t, 1, <-awaitForeach1)
	_, ok := <-awaitForeach1
	assert.False(t, ok)
	select {
	case <-awaitForeach4:
		require.Fail(t, "REDACTED")
	default:
	}

	//
	//
	c, err := staging.Following()
	require.NoError(t, err)
	assert.EqualValues(t, 0, c.Ordinal)

	w := staging.AwaitForeach(0)
	assert.EqualValues(t, 0, <-w)
	_, ok = <-w
	assert.False(t, ok)

	w = staging.AwaitForeach(1)
	assert.EqualValues(t, 1, <-w)
	_, ok = <-w
	assert.False(t, ok)

	//
	//
	err = staging.Shutdown()
	require.NoError(t, err)
	_, ok = <-awaitForeach4
	assert.False(t, ok)

	w = staging.AwaitForeach(3)
	_, ok = <-w
	assert.False(t, ok)
}
