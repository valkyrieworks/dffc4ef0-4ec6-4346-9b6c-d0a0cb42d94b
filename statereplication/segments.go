package statereplication

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	ctsync "github.com/valkyrieworks/utils/alignment"
	"github.com/valkyrieworks/p2p"
)

//
var errDone = errors.New("REDACTED")

//
type chunk struct {
	Height uint64
	Format uint32
	Index  uint32
	Chunk  []byte
	Sender p2p.ID
}

//
//
//
type chunkQueue struct {
	ctsync.Mutex
	snapshot       *snapshot                  //
	dir            string                     //
	chunkFiles     map[uint32]string          //
	chunkSenders   map[uint32]p2p.ID          //
	chunkAllocated map[uint32]bool            //
	chunkReturned  map[uint32]bool            //
	waiters        map[uint32][]chan<- uint32 //
}

//
//
func newChunkQueue(snapshot *snapshot, tempDir string) (*chunkQueue, error) {
	dir, err := os.MkdirTemp(tempDir, "REDACTED")
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	if snapshot.Chunks == 0 {
		return nil, errors.New("REDACTED")
	}
	return &chunkQueue{
		snapshot:       snapshot,
		dir:            dir,
		chunkFiles:     make(map[uint32]string, snapshot.Chunks),
		chunkSenders:   make(map[uint32]p2p.ID, snapshot.Chunks),
		chunkAllocated: make(map[uint32]bool, snapshot.Chunks),
		chunkReturned:  make(map[uint32]bool, snapshot.Chunks),
		waiters:        make(map[uint32][]chan<- uint32),
	}, nil
}

//
func (q *chunkQueue) Add(chunk *chunk) (bool, error) {
	if chunk == nil || chunk.Chunk == nil {
		return false, errors.New("REDACTED")
	}
	q.Lock()
	defer q.Unlock()
	if q.snapshot == nil {
		return false, nil //
	}
	if chunk.Height != q.snapshot.Height {
		return false, fmt.Errorf("REDACTED", chunk.Height, q.snapshot.Height)
	}
	if chunk.Format != q.snapshot.Format {
		return false, fmt.Errorf("REDACTED", chunk.Format, q.snapshot.Format)
	}
	if chunk.Index >= q.snapshot.Chunks {
		return false, fmt.Errorf("REDACTED", chunk.Index)
	}
	if q.chunkFiles[chunk.Index] != "REDACTED" {
		return false, nil
	}

	path := filepath.Join(q.dir, strconv.FormatUint(uint64(chunk.Index), 10))
	err := os.WriteFile(path, chunk.Chunk, 0o600)
	if err != nil {
		return false, fmt.Errorf("REDACTED", chunk.Index, path, err)
	}
	q.chunkFiles[chunk.Index] = path
	q.chunkSenders[chunk.Index] = chunk.Sender

	//
	for _, waiter := range q.waiters[chunk.Index] {
		waiter <- chunk.Index
		close(waiter)
	}
	delete(q.waiters, chunk.Index)

	return true, nil
}

//
//
func (q *chunkQueue) Allocate() (uint32, error) {
	q.Lock()
	defer q.Unlock()
	if q.snapshot == nil {
		return 0, errDone
	}
	if uint32(len(q.chunkAllocated)) >= q.snapshot.Chunks {
		return 0, errDone
	}
	for i := uint32(0); i < q.snapshot.Chunks; i++ {
		if !q.chunkAllocated[i] {
			q.chunkAllocated[i] = true
			return i, nil
		}
	}
	return 0, errDone
}

//
func (q *chunkQueue) Close() error {
	q.Lock()
	defer q.Unlock()
	if q.snapshot == nil {
		return nil
	}
	for _, waiters := range q.waiters {
		for _, waiter := range waiters {
			close(waiter)
		}
	}
	q.waiters = nil
	q.snapshot = nil
	err := os.RemoveAll(q.dir)
	if err != nil {
		return fmt.Errorf("REDACTED", q.dir, err)
	}
	return nil
}

//
//
//
func (q *chunkQueue) Discard(index uint32) error {
	q.Lock()
	defer q.Unlock()
	return q.discard(index)
}

//
func (q *chunkQueue) discard(index uint32) error {
	if q.snapshot == nil {
		return nil
	}
	path := q.chunkFiles[index]
	if path == "REDACTED" {
		return nil
	}
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("REDACTED", index, err)
	}
	delete(q.chunkFiles, index)
	delete(q.chunkReturned, index)
	delete(q.chunkAllocated, index)
	return nil
}

//
//
func (q *chunkQueue) DiscardSender(peerID p2p.ID) error {
	q.Lock()
	defer q.Unlock()

	for index, sender := range q.chunkSenders {
		if sender == peerID && !q.chunkReturned[index] {
			err := q.discard(index)
			if err != nil {
				return err
			}
			delete(q.chunkSenders, index)
		}
	}
	return nil
}

//
func (q *chunkQueue) GetSender(index uint32) p2p.ID {
	q.Lock()
	defer q.Unlock()
	return q.chunkSenders[index]
}

//
func (q *chunkQueue) Has(index uint32) bool {
	q.Lock()
	defer q.Unlock()
	return q.chunkFiles[index] != "REDACTED"
}

//
//
func (q *chunkQueue) load(index uint32) (*chunk, error) {
	path, ok := q.chunkFiles[index]
	if !ok {
		return nil, nil
	}
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", index, err)
	}
	return &chunk{
		Height: q.snapshot.Height,
		Format: q.snapshot.Format,
		Index:  index,
		Chunk:  body,
		Sender: q.chunkSenders[index],
	}, nil
}

//
//
func (q *chunkQueue) Next() (*chunk, error) {
	q.Lock()
	var chunk *chunk
	index, err := q.nextUp()
	if err == nil {
		chunk, err = q.load(index)
		if err == nil {
			q.chunkReturned[index] = true
		}
	}
	q.Unlock()
	if chunk != nil || err != nil {
		return chunk, err
	}

	select {
	case _, ok := <-q.WaitFor(index):
		if !ok {
			return nil, errDone //
		}
	case <-time.After(chunkTimeout):
		return nil, errTimeout
	}

	q.Lock()
	defer q.Unlock()
	chunk, err = q.load(index)
	if err != nil {
		return nil, err
	}
	q.chunkReturned[index] = true
	return chunk, nil
}

//
//
func (q *chunkQueue) nextUp() (uint32, error) {
	if q.snapshot == nil {
		return 0, errDone
	}
	for i := uint32(0); i < q.snapshot.Chunks; i++ {
		if !q.chunkReturned[i] {
			return i, nil
		}
	}
	return 0, errDone
}

//
func (q *chunkQueue) Retry(index uint32) {
	q.Lock()
	defer q.Unlock()
	delete(q.chunkReturned, index)
}

//
func (q *chunkQueue) RetryAll() {
	q.Lock()
	defer q.Unlock()
	q.chunkReturned = make(map[uint32]bool)
}

//
func (q *chunkQueue) Size() uint32 {
	q.Lock()
	defer q.Unlock()
	if q.snapshot == nil {
		return 0
	}
	return q.snapshot.Chunks
}

//
//
//
func (q *chunkQueue) WaitFor(index uint32) <-chan uint32 {
	q.Lock()
	defer q.Unlock()
	ch := make(chan uint32, 1)
	switch {
	case q.snapshot == nil:
		close(ch)
	case index >= q.snapshot.Chunks:
		close(ch)
	case q.chunkFiles[index] != "REDACTED":
		ch <- index
		close(ch)
	default:
		if q.waiters[index] == nil {
			q.waiters[index] = make([]chan<- uint32, 0)
		}
		q.waiters[index] = append(q.waiters[index], ch)
	}
	return ch
}
