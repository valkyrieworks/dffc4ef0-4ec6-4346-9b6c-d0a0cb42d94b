package statuschronize

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

//
var faultComplete = errors.New("REDACTED")

//
type segment struct {
	Altitude uint64
	Layout uint32
	Ordinal  uint32
	Segment  []byte
	Originator p2p.ID
}

//
//
//
type segmentStaging struct {
	commitchronize.Exclusion
	image       *image                  //
	dir            string                     //
	segmentRecords     map[uint32]string          //
	segmentOriginators   map[uint32]p2p.ID          //
	segmentAssigned map[uint32]bool            //
	segmentYielded  map[uint32]bool            //
	observers        map[uint32][]chan<- uint32 //
}

//
//
func freshSegmentStaging(image *image, transientPath string) (*segmentStaging, error) {
	dir, err := os.MkdirTemp(transientPath, "REDACTED")
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	if image.Segments == 0 {
		return nil, errors.New("REDACTED")
	}
	return &segmentStaging{
		image:       image,
		dir:            dir,
		segmentRecords:     make(map[uint32]string, image.Segments),
		segmentOriginators:   make(map[uint32]p2p.ID, image.Segments),
		segmentAssigned: make(map[uint32]bool, image.Segments),
		segmentYielded:  make(map[uint32]bool, image.Segments),
		observers:        make(map[uint32][]chan<- uint32),
	}, nil
}

//
func (q *segmentStaging) Add(segment *segment) (bool, error) {
	if segment == nil || segment.Segment == nil {
		return false, errors.New("REDACTED")
	}
	q.Lock()
	defer q.Unlock()
	if q.image == nil {
		return false, nil //
	}
	if segment.Altitude != q.image.Altitude {
		return false, fmt.Errorf("REDACTED", segment.Altitude, q.image.Altitude)
	}
	if segment.Layout != q.image.Layout {
		return false, fmt.Errorf("REDACTED", segment.Layout, q.image.Layout)
	}
	if segment.Ordinal >= q.image.Segments {
		return false, fmt.Errorf("REDACTED", segment.Ordinal)
	}
	if q.segmentRecords[segment.Ordinal] != "REDACTED" {
		return false, nil
	}

	route := filepath.Join(q.dir, strconv.FormatUint(uint64(segment.Ordinal), 10))
	err := os.WriteFile(route, segment.Segment, 0o600)
	if err != nil {
		return false, fmt.Errorf("REDACTED", segment.Ordinal, route, err)
	}
	q.segmentRecords[segment.Ordinal] = route
	q.segmentOriginators[segment.Ordinal] = segment.Originator

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
func (q *segmentStaging) Assign() (uint32, error) {
	q.Lock()
	defer q.Unlock()
	if q.image == nil {
		return 0, faultComplete
	}
	if uint32(len(q.segmentAssigned)) >= q.image.Segments {
		return 0, faultComplete
	}
	for i := uint32(0); i < q.image.Segments; i++ {
		if !q.segmentAssigned[i] {
			q.segmentAssigned[i] = true
			return i, nil
		}
	}
	return 0, faultComplete
}

//
func (q *segmentStaging) Shutdown() error {
	q.Lock()
	defer q.Unlock()
	if q.image == nil {
		return nil
	}
	for _, observers := range q.observers {
		for _, observer := range observers {
			close(observer)
		}
	}
	q.observers = nil
	q.image = nil
	err := os.RemoveAll(q.dir)
	if err != nil {
		return fmt.Errorf("REDACTED", q.dir, err)
	}
	return nil
}

//
//
//
func (q *segmentStaging) Eject(ordinal uint32) error {
	q.Lock()
	defer q.Unlock()
	return q.eject(ordinal)
}

//
func (q *segmentStaging) eject(ordinal uint32) error {
	if q.image == nil {
		return nil
	}
	route := q.segmentRecords[ordinal]
	if route == "REDACTED" {
		return nil
	}
	err := os.Remove(route)
	if err != nil {
		return fmt.Errorf("REDACTED", ordinal, err)
	}
	delete(q.segmentRecords, ordinal)
	delete(q.segmentYielded, ordinal)
	delete(q.segmentAssigned, ordinal)
	return nil
}

//
//
func (q *segmentStaging) EjectOriginator(nodeUUID p2p.ID) error {
	q.Lock()
	defer q.Unlock()

	for ordinal, originator := range q.segmentOriginators {
		if originator == nodeUUID && !q.segmentYielded[ordinal] {
			err := q.eject(ordinal)
			if err != nil {
				return err
			}
			delete(q.segmentOriginators, ordinal)
		}
	}
	return nil
}

//
func (q *segmentStaging) ObtainOriginator(ordinal uint32) p2p.ID {
	q.Lock()
	defer q.Unlock()
	return q.segmentOriginators[ordinal]
}

//
func (q *segmentStaging) Has(ordinal uint32) bool {
	q.Lock()
	defer q.Unlock()
	return q.segmentRecords[ordinal] != "REDACTED"
}

//
//
func (q *segmentStaging) fetch(ordinal uint32) (*segment, error) {
	route, ok := q.segmentRecords[ordinal]
	if !ok {
		return nil, nil
	}
	content, err := os.ReadFile(route)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", ordinal, err)
	}
	return &segment{
		Altitude: q.image.Altitude,
		Layout: q.image.Layout,
		Ordinal:  ordinal,
		Segment:  content,
		Originator: q.segmentOriginators[ordinal],
	}, nil
}

//
//
func (q *segmentStaging) Following() (*segment, error) {
	q.Lock()
	var segment *segment
	ordinal, err := q.followingActive()
	if err == nil {
		segment, err = q.fetch(ordinal)
		if err == nil {
			q.segmentYielded[ordinal] = true
		}
	}
	q.Unlock()
	if segment != nil || err != nil {
		return segment, err
	}

	select {
	case _, ok := <-q.AwaitForeach(ordinal):
		if !ok {
			return nil, faultComplete //
		}
	case <-time.After(segmentDeadline):
		return nil, faultDeadline
	}

	q.Lock()
	defer q.Unlock()
	segment, err = q.fetch(ordinal)
	if err != nil {
		return nil, err
	}
	q.segmentYielded[ordinal] = true
	return segment, nil
}

//
//
func (q *segmentStaging) followingActive() (uint32, error) {
	if q.image == nil {
		return 0, faultComplete
	}
	for i := uint32(0); i < q.image.Segments; i++ {
		if !q.segmentYielded[i] {
			return i, nil
		}
	}
	return 0, faultComplete
}

//
func (q *segmentStaging) Reissue(ordinal uint32) {
	q.Lock()
	defer q.Unlock()
	delete(q.segmentYielded, ordinal)
}

//
func (q *segmentStaging) ReissueEvery() {
	q.Lock()
	defer q.Unlock()
	q.segmentYielded = make(map[uint32]bool)
}

//
func (q *segmentStaging) Extent() uint32 {
	q.Lock()
	defer q.Unlock()
	if q.image == nil {
		return 0
	}
	return q.image.Segments
}

//
//
//
func (q *segmentStaging) AwaitForeach(ordinal uint32) <-chan uint32 {
	q.Lock()
	defer q.Unlock()
	ch := make(chan uint32, 1)
	switch {
	case q.image == nil:
		close(ch)
	case ordinal >= q.image.Segments:
		close(ch)
	case q.segmentRecords[ordinal] != "REDACTED":
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
