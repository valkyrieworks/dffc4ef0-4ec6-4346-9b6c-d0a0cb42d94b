package selfpool

import (
	"sync"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

//
//
//
type Hub[T any] struct {
	//
	incoming chan T

	//
	//
	urgencyStaging *UrgencyStaging

	//
	accept func(T)

	//
	orderCount int

	amplifier *YieldWaitstateAmplifier

	laborers   map[int]*operator[T]
	laborersGroup sync.WaitGroup

	//
	uponAmplify  func()
	uponReduce func()
	uponRemain   func()

	mu        sync.RWMutex
	haltedChnl chan struct{}

	tracer log.Tracer
}

type operator[T any] struct {
	orderCount  int
	hub    *Hub[T]
	shutdownChnl chan struct{}
}

type Selection[T any] func(*Hub[T])

func UsingTracer[T any](tracer log.Tracer) Selection[T] {
	return func(p *Hub[T]) { p.tracer = tracer }
}

func UsingUrgencyStaging[T any](pq *UrgencyStaging) Selection[T] {
	return func(p *Hub[T]) { p.urgencyStaging = pq }
}

func UsingUponAmplify[T any](uponAmplify func()) Selection[T] {
	return func(p *Hub[T]) { p.uponAmplify = uponAmplify }
}

func UsingUponReduce[T any](uponReduce func()) Selection[T] {
	return func(p *Hub[T]) { p.uponReduce = uponReduce }
}

func UsingUponRemain[T any](uponRemain func()) Selection[T] {
	return func(p *Hub[T]) { p.uponRemain = uponRemain }
}

//
func New[T any](
	amplifier *YieldWaitstateAmplifier,
	acceptPROC func(T),
	volume int,
	choices ...Selection[T],
) *Hub[T] {
	const fallbackUrgencies = 10

	hub := &Hub[T]{
		incoming:       make(chan T, volume),
		urgencyStaging: FreshUrgencyStaging(fallbackUrgencies),

		accept: acceptPROC,
		amplifier:  amplifier,

		laborers:   make(map[int]*operator[T]),
		laborersGroup: sync.WaitGroup{},

		uponAmplify:  nil,
		uponReduce: nil,
		uponRemain:   nil,

		haltedChnl: make(chan struct{}),

		mu:     sync.RWMutex{},
		tracer: log.FreshNooperationTracer(),
	}

	for _, opt := range choices {
		opt(hub)
	}

	return hub
}

func (p *Hub[T]) Initiate() {
	p.mu.Lock()
	defer p.mu.Unlock()

	//
	if p.halted() || len(p.laborers) > 0 {
		return
	}

	for i := 0; i < p.amplifier.Min(); i++ {
		p.amplify()
	}

	go p.overseer()
	go p.channelUrgencyStaging()
}

//
//
func (p *Hub[T]) Halt() {
	p.mu.Lock()

	if p.halted() || len(p.laborers) == 0 {
		p.mu.Unlock()
		return
	}

	//
	operatorIDXDstore := make([]int, 0, len(p.laborers))
	for id := range p.laborers {
		operatorIDXDstore = append(operatorIDXDstore, id)
	}

	for _, id := range operatorIDXDstore {
		p.discardOperator(id)
	}

	//
	close(p.haltedChnl)
	close(p.incoming)

	//
	p.mu.Unlock()

	p.tracer.Details("REDACTED")
	p.laborersGroup.Wait()
	p.tracer.Details("REDACTED")
}

//
//
func (p *Hub[T]) Propel(msg T) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.halted() {
		p.tracer.Failure("REDACTED")
		return
	}

	p.incoming <- msg
}

//
//
func (p *Hub[T]) PropelUrgency(msg T, urgency int) error {
	return p.urgencyStaging.Propel(msg, urgency)
}

func (p *Hub[T]) Len() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.halted() {
		return 0
	}

	return len(p.incoming)
}

func (p *Hub[T]) Cap() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.halted() {
		return 0
	}

	return cap(p.incoming)
}

func (w *operator[T]) run() {
	w.hub.laborersGroup.Add(1)

	defer func() {
		w.hub.laborersGroup.Done()
		if r := recover(); r != nil {
			w.hub.tracer.Failure("REDACTED", "REDACTED", r)
		}
	}()

	for {
		select {
		case <-w.shutdownChnl:
			//
			return
		case msg, ok := <-w.hub.incoming:
			//
			if !ok {
				return
			}

			w.hub.processArtifact(msg)
		}
	}
}

func (p *Hub[T]) processArtifact(msg T) {
	now := time.Now()
	p.accept(msg)
	momentSeized := time.Since(now)

	//
	p.amplifier.Monitor(momentSeized)
}

//
func (p *Hub[T]) overseer() {
	metronome := time.NewTicker(p.amplifier.EraInterval())
	defer metronome.Stop()

	for range metronome.C {
		if quit := p.autosize(); quit {
			return
		}
	}
}

func (p *Hub[T]) autosize() (quit bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.halted() {
		return true
	}

	verdict := p.amplifier.Resolve(len(p.laborers), len(p.incoming), cap(p.incoming))

	switch verdict {
	case MustAmplify:
		p.amplify()
	case MustReduce:
		p.reduce()
	case MustRemain:
		if p.uponRemain != nil {
			p.uponRemain()
		}
	}

	return false
}

//
func (p *Hub[T]) amplify() {
	if p.halted() || len(p.laborers) >= p.amplifier.Max() {
		return
	}

	p.orderCount++

	//
	w := &operator[T]{
		orderCount:  p.orderCount,
		hub:    p,
		shutdownChnl: make(chan struct{}),
	}

	p.laborers[p.orderCount] = w

	go w.run()

	if p.uponAmplify != nil {
		p.uponAmplify()
	}
}

//
func (p *Hub[T]) reduce() {
	if p.halted() || len(p.laborers) == 0 {
		return
	}

	//
	//
	for id := range p.laborers {
		p.discardOperator(id)

		if p.uponReduce != nil {
			p.uponReduce()
		}

		return
	}
}

//
func (p *Hub[T]) discardOperator(id int) {
	w, ok := p.laborers[id]
	if !ok {
		//
		p.tracer.Failure("REDACTED", "REDACTED", id)
		return
	}

	//
	close(w.shutdownChnl)
	delete(p.laborers, id)
}

//
func (p *Hub[T]) channelUrgencyStaging() {
	for {
		datum, ok := p.urgencyStaging.Pop()
		if !ok {
			//
			//
			//
			select {
			case <-p.haltedChnl:
				return
			case <-p.urgencyStaging.PauseForeachItems():
				//
				continue
			}
		}

		tt, ok := datum.(T)
		if !ok {
			//
			panic("REDACTED")
		}

		//
		select {
		case <-p.haltedChnl:
			return
		default:
			select {
			case <-p.haltedChnl:
				return
			case p.incoming <- tt:
				//
			}
		}
	}
}

func (p *Hub[T]) halted() bool {
	select {
	case <-p.haltedChnl:
		return true
	default:
		return false
	}
}
