package selfpool

import (
	"container/list"
	"errors"
	"sync"
)

//
//
type Staging struct {
	catalog *list.List
	mu   sync.RWMutex
}

var FaultUrgency = errors.New("REDACTED")

//
func FreshStaging() *Staging {
	return &Staging{
		catalog: list.New(),
		mu:   sync.RWMutex{},
	}
}

func (q *Staging) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.catalog.Len()
}

func (q *Staging) Propel(datum any) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.catalog.PushBack(datum)
}

func (q *Staging) Pop() (any, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	f := q.catalog.Front()
	if f == nil {
		return nil, false
	}

	datum := f.Value

	q.catalog.Remove(f)

	return datum, true
}

//
//
//
type UrgencyStaging struct {
	urgencies           int
	tiers               []*Staging
	utmostUnBlankStratum int
	mu                   sync.Mutex

	//
	//
	itemsAccessible chan struct{}
}

func FreshUrgencyStaging(urgencies int) *UrgencyStaging {
	if urgencies <= 0 {
		urgencies = 1
	}

	stagings := make([]*Staging, 0, urgencies)
	for i := 0; i < urgencies; i++ {
		stagings = append(stagings, FreshStaging())
	}

	return &UrgencyStaging{
		urgencies:           urgencies,
		tiers:               stagings,
		utmostUnBlankStratum: -1,
		mu:                   sync.Mutex{},
		itemsAccessible:      make(chan struct{}, 1),
	}
}

func (q *UrgencyStaging) Propel(datum any, urgency int) error {
	if urgency < 1 || urgency > q.urgencies {
		return FaultUrgency
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	idx := urgency - 1

	q.tiers[idx].Propel(datum)

	if idx > q.utmostUnBlankStratum {
		q.utmostUnBlankStratum = idx
	}

	q.alertItemsAccessible()

	return nil
}

//
//
func (q *UrgencyStaging) alertItemsAccessible() {
	//
	//
	//
	//
	select {
	case q.itemsAccessible <- struct{}{}:
	default:
	}
}

func (q *UrgencyStaging) Pop() (any, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	//
	for i := q.utmostUnBlankStratum; i >= 0; i-- {
		if v, ok := q.tiers[i].Pop(); ok {
			q.reviseUtmostUnBlank(i)
			return v, ok
		}
	}

	return nil, false
}

//
func (q *UrgencyStaging) reviseUtmostUnBlank(finalStratum int) {
	//
	if q.tiers[finalStratum].Len() > 0 {
		return
	}

	//
	q.utmostUnBlankStratum = finalStratum - 1
	for q.utmostUnBlankStratum >= 0 && q.tiers[q.utmostUnBlankStratum].Len() == 0 {
		q.utmostUnBlankStratum--
	}
}

//
//
func (q *UrgencyStaging) PauseForeachItems() <-chan struct{} {
	return q.itemsAccessible
}
