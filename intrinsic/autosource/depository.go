package autosource

import (
	"sync"
	"time"

	"github.com/valkyrieworks/utils/log"
)

//
//
//
type Depository[T any] struct {
	//
	incoming chan T

	//
	//
	urgencyBuffer *UrgencyBuffer

	//
	accept func(T)

	//
	seqCount int

	adjuster *VelocityWaitperiodAdjuster

	operators   map[int]*operator[T]
	operatorsGroup sync.WaitGroup

	//
	onSize  func()
	onReduce func()
	onRemain   func()

	mu        sync.RWMutex
	ceasedChan chan struct{}

	tracer log.Tracer
}

type operator[T any] struct {
	seqCount  int
	depository    *Depository[T]
	endChan chan struct{}
}

type Setting[T any] func(*Depository[T])

func WithTracer[T any](tracer log.Tracer) Setting[T] {
	return func(p *Depository[T]) { p.tracer = tracer }
}

func WithUrgencyBuffer[T any](pq *UrgencyBuffer) Setting[T] {
	return func(p *Depository[T]) { p.urgencyBuffer = pq }
}

func WithOnRatio[T any](onSize func()) Setting[T] {
	return func(p *Depository[T]) { p.onSize = onSize }
}

func WithOnReduce[T any](onReduce func()) Setting[T] {
	return func(p *Depository[T]) { p.onReduce = onReduce }
}

func WithOnRemain[T any](onRemain func()) Setting[T] {
	return func(p *Depository[T]) { p.onRemain = onRemain }
}

//
func New[T any](
	adjuster *VelocityWaitperiodAdjuster,
	acceptFN func(T),
	volume int,
	opts ...Setting[T],
) *Depository[T] {
	const standardUrgencies = 10

	depository := &Depository[T]{
		incoming:       make(chan T, volume),
		urgencyBuffer: NewUrgencyBuffer(standardUrgencies),

		accept: acceptFN,
		adjuster:  adjuster,

		operators:   make(map[int]*operator[T]),
		operatorsGroup: sync.WaitGroup{},

		onSize:  nil,
		onReduce: nil,
		onRemain:   nil,

		ceasedChan: make(chan struct{}),

		mu:     sync.RWMutex{},
		tracer: log.NewNoopTracer(),
	}

	for _, opt := range opts {
		opt(depository)
	}

	return depository
}

func (p *Depository[T]) Begin() {
	p.mu.Lock()
	defer p.mu.Unlock()

	//
	if p.ceased() || len(p.operators) > 0 {
		return
	}

	for i := 0; i < p.adjuster.Min(); i++ {
		p.size()
	}

	go p.auditor()
	go p.pipeUrgencyBuffer()
}

//
//
func (p *Depository[T]) Halt() {
	p.mu.Lock()

	if p.ceased() || len(p.operators) == 0 {
		p.mu.Unlock()
		return
	}

	//
	operatorIDXDatastore := make([]int, 0, len(p.operators))
	for id := range p.operators {
		operatorIDXDatastore = append(operatorIDXDatastore, id)
	}

	for _, id := range operatorIDXDatastore {
		p.deleteOperator(id)
	}

	//
	close(p.ceasedChan)
	close(p.incoming)

	//
	p.mu.Unlock()

	p.tracer.Details("REDACTED")
	p.operatorsGroup.Wait()
	p.tracer.Details("REDACTED")
}

//
//
func (p *Depository[T]) Propel(msg T) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.ceased() {
		p.tracer.Fault("REDACTED")
		return
	}

	p.incoming <- msg
}

//
//
func (p *Depository[T]) PropelUrgency(msg T, urgency int) error {
	return p.urgencyBuffer.Propel(msg, urgency)
}

func (p *Depository[T]) Len() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.ceased() {
		return 0
	}

	return len(p.incoming)
}

func (p *Depository[T]) Cap() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.ceased() {
		return 0
	}

	return cap(p.incoming)
}

func (w *operator[T]) run() {
	w.depository.operatorsGroup.Add(1)

	defer func() {
		w.depository.operatorsGroup.Done()
		if r := recover(); r != nil {
			w.depository.tracer.Fault("REDACTED", "REDACTED", r)
		}
	}()

	for {
		select {
		case <-w.endChan:
			//
			return
		case msg, ok := <-w.depository.incoming:
			//
			if !ok {
				return
			}

			w.depository.processSignal(msg)
		}
	}
}

func (p *Depository[T]) processSignal(msg T) {
	now := time.Now()
	p.accept(msg)
	timeSeized := time.Since(now)

	//
	p.adjuster.Monitor(timeSeized)
}

//
func (p *Depository[T]) auditor() {
	timer := time.NewTicker(p.adjuster.EraPeriod())
	defer timer.Stop()

	for range timer.C {
		if quit := p.autosize(); quit {
			return
		}
	}
}

func (p *Depository[T]) autosize() (quit bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.ceased() {
		return true
	}

	verdict := p.adjuster.Determine(len(p.operators), len(p.incoming), cap(p.incoming))

	switch verdict {
	case MustSize:
		p.size()
	case MustReduce:
		p.reduce()
	case MustRemain:
		if p.onRemain != nil {
			p.onRemain()
		}
	}

	return false
}

//
func (p *Depository[T]) size() {
	if p.ceased() || len(p.operators) >= p.adjuster.Max() {
		return
	}

	p.seqCount++

	//
	w := &operator[T]{
		seqCount:  p.seqCount,
		depository:    p,
		endChan: make(chan struct{}),
	}

	p.operators[p.seqCount] = w

	go w.run()

	if p.onSize != nil {
		p.onSize()
	}
}

//
func (p *Depository[T]) reduce() {
	if p.ceased() || len(p.operators) == 0 {
		return
	}

	//
	//
	for id := range p.operators {
		p.deleteOperator(id)

		if p.onReduce != nil {
			p.onReduce()
		}

		return
	}
}

//
func (p *Depository[T]) deleteOperator(id int) {
	w, ok := p.operators[id]
	if !ok {
		//
		p.tracer.Fault("REDACTED", "REDACTED", id)
		return
	}

	//
	close(w.endChan)
	delete(p.operators, id)
}

//
func (p *Depository[T]) pipeUrgencyBuffer() {
	for {
		item, ok := p.urgencyBuffer.Pop()
		if !ok {
			//
			//
			//
			select {
			case <-p.ceasedChan:
				return
			case <-p.urgencyBuffer.WaitForItems():
				//
				continue
			}
		}

		tt, ok := item.(T)
		if !ok {
			//
			panic("REDACTED")
		}

		//
		select {
		case <-p.ceasedChan:
			return
		default:
			select {
			case <-p.ceasedChan:
				return
			case p.incoming <- tt:
				//
			}
		}
	}
}

func (p *Depository[T]) ceased() bool {
	select {
	case <-p.ceasedChan:
		return true
	default:
		return false
	}
}
