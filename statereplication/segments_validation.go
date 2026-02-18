package statereplication

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/p2p"
)

func setupChunkQueue(t *testing.T) (*chunkQueue, func()) {
	snapshot := &snapshot{
		Height:   3,
		Format:   1,
		Chunks:   5,
		Hash:     []byte{7},
		Metadata: nil,
	}
	queue, err := newChunkQueue(snapshot, "REDACTED")
	require.NoError(t, err)
	teardown := func() {
		err := queue.Close()
		require.NoError(t, err)
	}
	return queue, teardown
}

func TestNewChunkQueue_TempDir(t *testing.T) {
	snapshot := &snapshot{
		Height:   3,
		Format:   1,
		Chunks:   5,
		Hash:     []byte{7},
		Metadata: nil,
	}
	dir, err := os.MkdirTemp("REDACTED", "REDACTED")
	require.NoError(t, err)
	defer os.RemoveAll(dir)
	queue, err := newChunkQueue(snapshot, dir)
	require.NoError(t, err)

	files, err := os.ReadDir(dir)
	require.NoError(t, err)
	assert.Len(t, files, 1)

	err = queue.Close()
	require.NoError(t, err)

	files, err = os.ReadDir(dir)
	require.NoError(t, err)
	assert.Len(t, files, 0)
}

func TestChunkQueue(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	//
	added, err := queue.Add(&chunk{Height: 3, Format: 1, Index: 0, Chunk: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.True(t, added)

	//
	added, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 4, Chunk: []byte{3, 1, 4}})
	require.NoError(t, err)
	assert.True(t, added)

	//
	added, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 0, Chunk: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.False(t, added)

	added, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 4, Chunk: []byte{3, 1, 4}})
	require.NoError(t, err)
	assert.False(t, added)

	//
	added, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 3, Chunk: []byte{3, 1, 3}})
	require.NoError(t, err)
	assert.True(t, added)

	added, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 2, Chunk: []byte{3, 1, 2}})
	require.NoError(t, err)
	assert.True(t, added)

	added, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 1, Chunk: []byte{3, 1, 1}})
	require.NoError(t, err)
	assert.True(t, added)

	//
	for i := 0; i < 5; i++ {
		c, err := queue.Next()
		require.NoError(t, err)
		assert.Equal(t, &chunk{Height: 3, Format: 1, Index: uint32(i), Chunk: []byte{3, 1, byte(i)}}, c)
	}
	_, err = queue.Next()
	require.Error(t, err)
	assert.Equal(t, errDone, err)

	//
	added, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 0, Chunk: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.False(t, added)

	//
	err = queue.Close()
	require.NoError(t, err)
	added, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 0, Chunk: []byte{3, 1, 0}})
	require.NoError(t, err)
	assert.False(t, added)

	//
	err = queue.Close()
	require.NoError(t, err)
}

func TestChunkQueue_Add_ChunkErrors(t *testing.T) {
	testcases := map[string]struct {
		chunk *chunk
	}{
		"REDACTED":     {nil},
		"REDACTED":      {&chunk{Height: 3, Format: 1, Index: 0, Chunk: nil}},
		"REDACTED":  {&chunk{Height: 9, Format: 1, Index: 0, Chunk: []byte{3, 1, 0}}},
		"REDACTED":  {&chunk{Height: 3, Format: 9, Index: 0, Chunk: []byte{3, 1, 0}}},
		"REDACTED": {&chunk{Height: 3, Format: 1, Index: 5, Chunk: []byte{3, 1, 0}}},
	}
	for name, tc := range testcases {

		t.Run(name, func(t *testing.T) {
			queue, teardown := setupChunkQueue(t)
			defer teardown()
			_, err := queue.Add(tc.chunk)
			require.Error(t, err)
		})
	}
}

func TestChunkQueue_Allocate(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	for i := uint32(0); i < queue.Size(); i++ {
		index, err := queue.Allocate()
		require.NoError(t, err)
		assert.EqualValues(t, i, index)
	}

	_, err := queue.Allocate()
	require.Error(t, err)
	assert.Equal(t, errDone, err)

	for i := uint32(0); i < queue.Size(); i++ {
		_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: i, Chunk: []byte{byte(i)}})
		require.NoError(t, err)
	}

	//
	err = queue.Discard(2)
	require.NoError(t, err)

	index, err := queue.Allocate()
	require.NoError(t, err)
	assert.EqualValues(t, 2, index)
	_, err = queue.Allocate()
	require.Error(t, err)
	assert.Equal(t, errDone, err)

	//
	err = queue.Discard(2)
	require.NoError(t, err)
	err = queue.Close()
	require.NoError(t, err)
	_, err = queue.Allocate()
	require.Error(t, err)
	assert.Equal(t, errDone, err)
}

func TestChunkQueue_Discard(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	//
	_, err := queue.Add(&chunk{Height: 3, Format: 1, Index: 0, Chunk: []byte{byte(0)}})
	require.NoError(t, err)
	_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 1, Chunk: []byte{byte(1)}})
	require.NoError(t, err)
	_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 2, Chunk: []byte{byte(2)}})
	require.NoError(t, err)

	c, err := queue.Next()
	require.NoError(t, err)
	assert.EqualValues(t, 0, c.Index)
	c, err = queue.Next()
	require.NoError(t, err)
	assert.EqualValues(t, 1, c.Index)

	//
	//
	err = queue.Discard(0)
	require.NoError(t, err)
	added, err := queue.Add(&chunk{Height: 3, Format: 1, Index: 0, Chunk: []byte{byte(0)}})
	require.NoError(t, err)
	assert.True(t, added)
	c, err = queue.Next()
	require.NoError(t, err)
	assert.EqualValues(t, 0, c.Index)
	c, err = queue.Next()
	require.NoError(t, err)
	assert.EqualValues(t, 2, c.Index)

	//
	for i := uint32(0); i < queue.Size(); i++ {
		err := queue.Discard(i)
		require.NoError(t, err)
	}
	for i := uint32(0); i < queue.Size(); i++ {
		_, err := queue.Allocate()
		require.NoError(t, err)
		_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: i, Chunk: []byte{byte(i)}})
		require.NoError(t, err)
		c, err = queue.Next()
		require.NoError(t, err)
		assert.EqualValues(t, i, c.Index)
	}

	//
	err = queue.Discard(99)
	require.NoError(t, err)

	//
	err = queue.Discard(3)
	require.NoError(t, err)
	err = queue.Discard(1)
	require.NoError(t, err)

	index, err := queue.Allocate()
	require.NoError(t, err)
	assert.EqualValues(t, 1, index)
	index, err = queue.Allocate()
	require.NoError(t, err)
	assert.EqualValues(t, 3, index)

	added, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 3, Chunk: []byte{3}})
	require.NoError(t, err)
	assert.True(t, added)
	added, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 1, Chunk: []byte{1}})
	require.NoError(t, err)
	assert.True(t, added)

	chunk, err := queue.Next()
	require.NoError(t, err)
	assert.EqualValues(t, 1, chunk.Index)

	chunk, err = queue.Next()
	require.NoError(t, err)
	assert.EqualValues(t, 3, chunk.Index)

	_, err = queue.Next()
	require.Error(t, err)
	assert.Equal(t, errDone, err)

	//
	err = queue.Close()
	require.NoError(t, err)
	err = queue.Discard(2)
	require.NoError(t, err)
}

func TestChunkQueue_DiscardSender(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	//
	senders := []p2p.ID{"REDACTED", "REDACTED", "REDACTED"}
	for i := uint32(0); i < queue.Size(); i++ {
		_, err := queue.Allocate()
		require.NoError(t, err)
		_, err = queue.Add(&chunk{
			Height: 3,
			Format: 1,
			Index:  i,
			Chunk:  []byte{byte(i)},
			Sender: senders[int(i)%len(senders)],
		})
		require.NoError(t, err)
	}

	//
	for i := uint32(0); i < 3; i++ {
		_, err := queue.Next()
		require.NoError(t, err)
	}

	//
	err := queue.DiscardSender("REDACTED")
	require.NoError(t, err)
	_, err = queue.Allocate()
	assert.Equal(t, errDone, err)

	//
	//
	err = queue.DiscardSender("REDACTED")
	require.NoError(t, err)
	index, err := queue.Allocate()
	require.NoError(t, err)
	assert.EqualValues(t, 4, index)
	_, err = queue.Allocate()
	assert.Equal(t, errDone, err)
}

func TestChunkQueue_GetSender(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	_, err := queue.Add(&chunk{Height: 3, Format: 1, Index: 0, Chunk: []byte{1}, Sender: p2p.ID("REDACTED")})
	require.NoError(t, err)
	_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 1, Chunk: []byte{2}, Sender: p2p.ID("REDACTED")})
	require.NoError(t, err)

	assert.EqualValues(t, "REDACTED", queue.GetSender(0))
	assert.EqualValues(t, "REDACTED", queue.GetSender(1))
	assert.EqualValues(t, "REDACTED", queue.GetSender(2))

	//
	chunk, err := queue.Next()
	require.NoError(t, err)
	require.NotNil(t, chunk)
	require.EqualValues(t, 0, chunk.Index)
	assert.EqualValues(t, "REDACTED", queue.GetSender(0))
}

func TestChunkQueue_Next(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	//
	chNext := make(chan *chunk, 10)
	go func() {
		for {
			c, err := queue.Next()
			if err == errDone {
				close(chNext)
				break
			}
			require.NoError(t, err)
			chNext <- c
		}
	}()

	assert.Empty(t, chNext)
	_, err := queue.Add(&chunk{Height: 3, Format: 1, Index: 1, Chunk: []byte{3, 1, 1}, Sender: p2p.ID("REDACTED")})
	require.NoError(t, err)
	select {
	case <-chNext:
		assert.Fail(t, "REDACTED")
	default:
	}

	_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 0, Chunk: []byte{3, 1, 0}, Sender: p2p.ID("REDACTED")})
	require.NoError(t, err)

	assert.Equal(t,
		&chunk{Height: 3, Format: 1, Index: 0, Chunk: []byte{3, 1, 0}, Sender: p2p.ID("REDACTED")},
		<-chNext)
	assert.Equal(t,
		&chunk{Height: 3, Format: 1, Index: 1, Chunk: []byte{3, 1, 1}, Sender: p2p.ID("REDACTED")},
		<-chNext)

	_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 4, Chunk: []byte{3, 1, 4}, Sender: p2p.ID("REDACTED")})
	require.NoError(t, err)
	select {
	case <-chNext:
		assert.Fail(t, "REDACTED")
	default:
	}

	_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 2, Chunk: []byte{3, 1, 2}, Sender: p2p.ID("REDACTED")})
	require.NoError(t, err)
	_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 3, Chunk: []byte{3, 1, 3}, Sender: p2p.ID("REDACTED")})
	require.NoError(t, err)

	assert.Equal(t,
		&chunk{Height: 3, Format: 1, Index: 2, Chunk: []byte{3, 1, 2}, Sender: p2p.ID("REDACTED")},
		<-chNext)
	assert.Equal(t,
		&chunk{Height: 3, Format: 1, Index: 3, Chunk: []byte{3, 1, 3}, Sender: p2p.ID("REDACTED")},
		<-chNext)
	assert.Equal(t,
		&chunk{Height: 3, Format: 1, Index: 4, Chunk: []byte{3, 1, 4}, Sender: p2p.ID("REDACTED")},
		<-chNext)

	_, ok := <-chNext
	assert.False(t, ok, "REDACTED")

	//
	_, err = queue.Next()
	assert.Equal(t, errDone, err)
}

func TestChunkQueue_Next_Closed(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	//
	_, err := queue.Add(&chunk{Height: 3, Format: 1, Index: 1, Chunk: []byte{3, 1, 1}})
	require.NoError(t, err)
	err = queue.Close()
	require.NoError(t, err)

	_, err = queue.Next()
	assert.Equal(t, errDone, err)
}

func TestChunkQueue_Retry(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	//
	for i := uint32(0); i < queue.Size(); i++ {
		_, err := queue.Allocate()
		require.NoError(t, err)
		_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: i, Chunk: []byte{byte(i)}})
		require.NoError(t, err)
		_, err = queue.Next()
		require.NoError(t, err)
	}

	//
	queue.Retry(3)
	queue.Retry(1)

	_, err := queue.Allocate()
	assert.Equal(t, errDone, err)

	chunk, err := queue.Next()
	require.NoError(t, err)
	assert.EqualValues(t, 1, chunk.Index)

	chunk, err = queue.Next()
	require.NoError(t, err)
	assert.EqualValues(t, 3, chunk.Index)

	_, err = queue.Next()
	assert.Equal(t, errDone, err)
}

func TestChunkQueue_RetryAll(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	//
	for i := uint32(0); i < queue.Size(); i++ {
		_, err := queue.Allocate()
		require.NoError(t, err)
		_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: i, Chunk: []byte{byte(i)}})
		require.NoError(t, err)
		_, err = queue.Next()
		require.NoError(t, err)
	}

	_, err := queue.Next()
	assert.Equal(t, errDone, err)

	queue.RetryAll()

	_, err = queue.Allocate()
	assert.Equal(t, errDone, err)

	for i := uint32(0); i < queue.Size(); i++ {
		chunk, err := queue.Next()
		require.NoError(t, err)
		assert.EqualValues(t, i, chunk.Index)
	}

	_, err = queue.Next()
	assert.Equal(t, errDone, err)
}

func TestChunkQueue_Size(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	assert.EqualValues(t, 5, queue.Size())

	err := queue.Close()
	require.NoError(t, err)
	assert.EqualValues(t, 0, queue.Size())
}

func TestChunkQueue_WaitFor(t *testing.T) {
	queue, teardown := setupChunkQueue(t)
	defer teardown()

	waitFor1 := queue.WaitFor(1)
	waitFor4 := queue.WaitFor(4)

	//
	_, err := queue.Add(&chunk{Height: 3, Format: 1, Index: 0, Chunk: []byte{3, 1, 0}})
	require.NoError(t, err)
	_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 2, Chunk: []byte{3, 1, 2}})
	require.NoError(t, err)
	select {
	case <-waitFor1:
		require.Fail(t, "REDACTED")
	case <-waitFor4:
		require.Fail(t, "REDACTED")
	default:
	}

	//
	_, err = queue.Add(&chunk{Height: 3, Format: 1, Index: 1, Chunk: []byte{3, 1, 1}})
	require.NoError(t, err)
	assert.EqualValues(t, 1, <-waitFor1)
	_, ok := <-waitFor1
	assert.False(t, ok)
	select {
	case <-waitFor4:
		require.Fail(t, "REDACTED")
	default:
	}

	//
	//
	c, err := queue.Next()
	require.NoError(t, err)
	assert.EqualValues(t, 0, c.Index)

	w := queue.WaitFor(0)
	assert.EqualValues(t, 0, <-w)
	_, ok = <-w
	assert.False(t, ok)

	w = queue.WaitFor(1)
	assert.EqualValues(t, 1, <-w)
	_, ok = <-w
	assert.False(t, ok)

	//
	//
	err = queue.Close()
	require.NoError(t, err)
	_, ok = <-waitFor4
	assert.False(t, ok)

	w = queue.WaitFor(3)
	_, ok = <-w
	assert.False(t, ok)
}
