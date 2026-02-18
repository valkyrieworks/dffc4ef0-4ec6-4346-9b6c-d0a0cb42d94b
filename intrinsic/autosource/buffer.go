package autosource

import (
	"container/list"
	"errors"
	"sync"
)

//
//
type Buffer struct {
	catalog *list.List
	mu   sync.RWMutex
}

var ErrUrgency = errors.New("REDACTED")

//
func NewBuffer() *Buffer {
	return &Buffer{
		catalog: list.New(),
		mu:   sync.RWMutex{},
	}
}

func (q *Buffer) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.catalog.Len()
}

func (q *Buffer) Propel(item any) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.catalog.PushBack(item)
}

func (q *Buffer) Pop() (any, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	f := q.catalog.Front()
	if f == nil {
		return nil, false
	}

	item := f.Value

	q.catalog.Remove(f)

	return item, true
}

//
//
//
type UrgencyBuffer struct {
	urgencies           int
	tiers               []*Buffer
	greatestNotEmptyLayer int
	mu                   sync.Mutex

	//
	//
	itemsAccessible chan struct{}
}

func NewUrgencyBuffer(urgencies int) *UrgencyBuffer {
	if urgencies <= 0 {
		urgencies = 1
	}

	buffers := make([]*Buffer, 0, urgencies)
	for i := 0; i < urgencies; i++ {
		buffers = append(buffers, NewBuffer())
	}

	return &UrgencyBuffer{
		urgencies:           urgencies,
		tiers:               buffers,
		greatestNotEmptyLayer: -1,
		mu:                   sync.Mutex{},
		itemsAccessible:      make(chan struct{}, 1),
	}
}

func (q *UrgencyBuffer) Propel(item any, urgency int) error {
	if urgency < 1 || urgency > q.urgencies {
		return ErrUrgency
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	idx := urgency - 1

	q.tiers[idx].Propel(item)

	if idx > q.greatestNotEmptyLayer {
		q.greatestNotEmptyLayer = idx
	}

	q.alertItemsAccessible()

	return nil
}

//
//
func (q *UrgencyBuffer) alertItemsAccessible() {
	//
	//
	//
	//
	select {
	case q.itemsAccessible <- struct{}{}:
	default:
	}
}

func (q *UrgencyBuffer) Pop() (any, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	//
	for i := q.greatestNotEmptyLayer; i >= 0; i-- {
		if v, ok := q.tiers[i].Pop(); ok {
			q.modifyGreatestNotEmpty(i)
			return v, ok
		}
	}

	return nil, false
}

//
func (q *UrgencyBuffer) modifyGreatestNotEmpty(finalLayer int) {
	//
	if q.tiers[finalLayer].Len() > 0 {
		return
	}

	//
	q.greatestNotEmptyLayer = finalLayer - 1
	for q.greatestNotEmptyLayer >= 0 && q.tiers[q.greatestNotEmptyLayer].Len() == 0 {
		q.greatestNotEmptyLayer--
	}
}

//
//
func (q *UrgencyBuffer) WaitForItems() <-chan struct{} {
	return q.itemsAccessible
}
