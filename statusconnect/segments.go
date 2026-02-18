package statusconnect

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	engineconnect "github.com/valkyrieworks/utils/align"
	"github.com/valkyrieworks/p2p"
)

//
var errDone = errors.New("REDACTED")

//
type segment struct {
	Level uint64
	Layout uint32
	Ordinal  uint32
	Segment  []byte
	Emitter p2p.ID
}

//
//
//
type segmentBuffer struct {
	engineconnect.Lock
	mirror       *mirror                  //
	dir            string                     //
	segmentEntries     map[uint32]string          //
	segmentEmitters   map[uint32]p2p.ID          //
	segmentAssigned map[uint32]bool            //
	segmentYielded  map[uint32]bool            //
	observers        map[uint32][]chan<- uint32 //
}

//
//
func newSegmentBuffer(mirror *mirror, temporaryFolder string) (*segmentBuffer, error) {
	dir, err := os.MkdirTemp(temporaryFolder, "REDACTED")
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	if mirror.Segments == 0 {
		return nil, errors.New("REDACTED")
	}
	return &segmentBuffer{
		mirror:       mirror,
		dir:            dir,
		segmentEntries:     make(map[uint32]string, mirror.Segments),
		segmentEmitters:   make(map[uint32]p2p.ID, mirror.Segments),
		segmentAssigned: make(map[uint32]bool, mirror.Segments),
		segmentYielded:  make(map[uint32]bool, mirror.Segments),
		observers:        make(map[uint32][]chan<- uint32),
	}, nil
}

//
func (q *segmentBuffer) Add(segment *segment) (bool, error) {
	if segment == nil || segment.Segment == nil {
		return false, errors.New("REDACTED")
	}
	q.Lock()
	defer q.Unlock()
	if q.mirror == nil {
		return false, nil //
	}
	if segment.Level != q.mirror.Level {
		return false, fmt.Errorf("REDACTED", segment.Level, q.mirror.Level)
	}
	if segment.Layout != q.mirror.Layout {
		return false, fmt.Errorf("REDACTED", segment.Layout, q.mirror.Layout)
	}
	if segment.Ordinal >= q.mirror.Segments {
		return false, fmt.Errorf("REDACTED", segment.Ordinal)
	}
	if q.segmentEntries[segment.Ordinal] != "REDACTED" {
		return false, nil
	}

	route := filepath.Join(q.dir, strconv.FormatUint(uint64(segment.Ordinal), 10))
	err := os.WriteFile(route, segment.Segment, 0o600)
	if err != nil {
		return false, fmt.Errorf("REDACTED", segment.Ordinal, route, err)
	}
	q.segmentEntries[segment.Ordinal] = route
	q.segmentEmitters[segment.Ordinal] = segment.Emitter

	//
	for _, observer := range q.observers[segment.Ordinal] {
		observer <- segment.Ordinal
		close(observer)
	}
	delete(q.observers, segment.Ordinal)

	return true, nil
}

//
//
func (q *segmentBuffer) Assign() (uint32, error) {
	q.Lock()
	defer q.Unlock()
	if q.mirror == nil {
		return 0, errDone
	}
	if uint32(len(q.segmentAssigned)) >= q.mirror.Segments {
		return 0, errDone
	}
	for i := uint32(0); i < q.mirror.Segments; i++ {
		if !q.segmentAssigned[i] {
			q.segmentAssigned[i] = true
			return i, nil
		}
	}
	return 0, errDone
}

//
func (q *segmentBuffer) End() error {
	q.Lock()
	defer q.Unlock()
	if q.mirror == nil {
		return nil
	}
	for _, observers := range q.observers {
		for _, observer := range observers {
			close(observer)
		}
	}
	q.observers = nil
	q.mirror = nil
	err := os.RemoveAll(q.dir)
	if err != nil {
		return fmt.Errorf("REDACTED", q.dir, err)
	}
	return nil
}

//
//
//
func (q *segmentBuffer) Drop(ordinal uint32) error {
	q.Lock()
	defer q.Unlock()
	return q.drop(ordinal)
}

//
func (q *segmentBuffer) drop(ordinal uint32) error {
	if q.mirror == nil {
		return nil
	}
	route := q.segmentEntries[ordinal]
	if route == "REDACTED" {
		return nil
	}
	err := os.Remove(route)
	if err != nil {
		return fmt.Errorf("REDACTED", ordinal, err)
	}
	delete(q.segmentEntries, ordinal)
	delete(q.segmentYielded, ordinal)
	delete(q.segmentAssigned, ordinal)
	return nil
}

//
//
func (q *segmentBuffer) DropEmitter(nodeUID p2p.ID) error {
	q.Lock()
	defer q.Unlock()

	for ordinal, emitter := range q.segmentEmitters {
		if emitter == nodeUID && !q.segmentYielded[ordinal] {
			err := q.drop(ordinal)
			if err != nil {
				return err
			}
			delete(q.segmentEmitters, ordinal)
		}
	}
	return nil
}

//
func (q *segmentBuffer) FetchEmitter(ordinal uint32) p2p.ID {
	q.Lock()
	defer q.Unlock()
	return q.segmentEmitters[ordinal]
}

//
func (q *segmentBuffer) Has(ordinal uint32) bool {
	q.Lock()
	defer q.Unlock()
	return q.segmentEntries[ordinal] != "REDACTED"
}

//
//
func (q *segmentBuffer) import(ordinal uint32) (*segment, error) {
	route, ok := q.segmentEntries[ordinal]
	if !ok {
		return nil, nil
	}
	content, err := os.ReadFile(route)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", ordinal, err)
	}
	return &segment{
		Level: q.mirror.Level,
		Layout: q.mirror.Layout,
		Ordinal:  ordinal,
		Segment:  content,
		Emitter: q.segmentEmitters[ordinal],
	}, nil
}

//
//
func (q *segmentBuffer) Following() (*segment, error) {
	q.Lock()
	var segment *segment
	ordinal, err := q.followingUp()
	if err == nil {
		segment, err = q.import(ordinal)
		if err == nil {
			q.segmentYielded[ordinal] = true
		}
	}
	q.Unlock()
	if segment != nil || err != nil {
		return segment, err
	}

	select {
	case _, ok := <-q.WaitFor(ordinal):
		if !ok {
			return nil, errDone //
		}
	case <-time.After(segmentDeadline):
		return nil, errDeadline
	}

	q.Lock()
	defer q.Unlock()
	segment, err = q.import(ordinal)
	if err != nil {
		return nil, err
	}
	q.segmentYielded[ordinal] = true
	return segment, nil
}

//
//
func (q *segmentBuffer) followingUp() (uint32, error) {
	if q.mirror == nil {
		return 0, errDone
	}
	for i := uint32(0); i < q.mirror.Segments; i++ {
		if !q.segmentYielded[i] {
			return i, nil
		}
	}
	return 0, errDone
}

//
func (q *segmentBuffer) Reprocess(ordinal uint32) {
	q.Lock()
	defer q.Unlock()
	delete(q.segmentYielded, ordinal)
}

//
func (q *segmentBuffer) ReprocessAll() {
	q.Lock()
	defer q.Unlock()
	q.segmentYielded = make(map[uint32]bool)
}

//
func (q *segmentBuffer) Volume() uint32 {
	q.Lock()
	defer q.Unlock()
	if q.mirror == nil {
		return 0
	}
	return q.mirror.Segments
}

//
//
//
func (q *segmentBuffer) WaitFor(ordinal uint32) <-chan uint32 {
	q.Lock()
	defer q.Unlock()
	ch := make(chan uint32, 1)
	switch {
	case q.mirror == nil:
		close(ch)
	case ordinal >= q.mirror.Segments:
		close(ch)
	case q.segmentEntries[ordinal] != "REDACTED":
		ch <- ordinal
		close(ch)
	default:
		if q.observers[ordinal] == nil {
			q.observers[ordinal] = make([]chan<- uint32, 0)
		}
		q.observers[ordinal] = append(q.observers[ordinal], ch)
	}
	return ch
}
