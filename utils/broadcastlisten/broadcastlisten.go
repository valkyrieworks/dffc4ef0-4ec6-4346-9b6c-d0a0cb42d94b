//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
package broadcastlisten

import (
	"context"
	"errors"
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

type procedure int

const (
	sub procedure = iota
	pub
	unlisten
	halt
)

var (
	//
	//
	FaultListeningNegationDetected = errors.New("REDACTED")

	//
	//
	FaultEarlierListened = errors.New("REDACTED")
)

//
//
//
//
//
//
type Inquire interface {
	Aligns(incidents map[string][]string) (bool, error)
	Text() string
}

type cmd struct {
	op procedure

	//
	inquire        Inquire
	listening *Listening
	customerUUID     string

	//
	msg    any
	incidents map[string][]string
}

//
//
type Daemon struct {
	facility.FoundationFacility

	actions    chan cmd
	actionsCeiling int

	//
	//
	mtx           commitchronize.ReadwriteExclusion
	feeds map[string]map[string]struct{} //
}

//
type Selection func(*Daemon)

//
//
//
func FreshDaemon(choices ...Selection) *Daemon {
	s := &Daemon{
		feeds: make(map[string]map[string]struct{}),
	}
	s.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", s)

	for _, selection := range choices {
		selection(s)
	}

	//
	s.actions = make(chan cmd, s.actionsCeiling)

	return s
}

//
//
//
//
func ReserveVolume(cap int) Selection {
	return func(s *Daemon) {
		if cap > 0 {
			s.actionsCeiling = cap
		}
	}
}

//
func (s *Daemon) ReserveVolume() int {
	return s.actionsCeiling
}

//
//
//
//
//
//
//
//
func (s *Daemon) Listen(
	ctx context.Context,
	customerUUID string,
	inquire Inquire,
	outputVolume ...int,
) (*Listening, error) {
	outputCeiling := 1
	if len(outputVolume) > 0 {
		if outputVolume[0] <= 0 {
			panic("REDACTED")
		}
		outputCeiling = outputVolume[0]
	}

	return s.listen(ctx, customerUUID, inquire, outputCeiling)
}

//
//
//
func (s *Daemon) ListenUncached(ctx context.Context, customerUUID string, inquire Inquire) (*Listening, error) {
	return s.listen(ctx, customerUUID, inquire, 0)
}

func (s *Daemon) listen(ctx context.Context, customerUUID string, inquire Inquire, outputVolume int) (*Listening, error) {
	s.mtx.RLock()
	customerFeeds, ok := s.feeds[customerUUID]
	if ok {
		_, ok = customerFeeds[inquire.Text()]
	}
	s.mtx.RUnlock()
	if ok {
		return nil, FaultEarlierListened
	}

	listening := FreshListening(outputVolume)
	select {
	case s.actions <- cmd{op: sub, customerUUID: customerUUID, inquire: inquire, listening: listening}:
		s.mtx.Lock()
		if _, ok = s.feeds[customerUUID]; !ok {
			s.feeds[customerUUID] = make(map[string]struct{})
		}
		s.feeds[customerUUID][inquire.Text()] = struct{}{}
		s.mtx.Unlock()
		return listening, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-s.Exit():
		return nil, errors.New("REDACTED")
	}
}

//
//
//
func (s *Daemon) Unlisten(ctx context.Context, customerUUID string, inquire Inquire) error {
	s.mtx.RLock()
	customerFeeds, ok := s.feeds[customerUUID]
	if ok {
		_, ok = customerFeeds[inquire.Text()]
	}
	s.mtx.RUnlock()
	if !ok {
		return FaultListeningNegationDetected
	}

	select {
	case s.actions <- cmd{op: unlisten, customerUUID: customerUUID, inquire: inquire}:
		s.mtx.Lock()
		delete(customerFeeds, inquire.Text())
		if len(customerFeeds) == 0 {
			delete(s.feeds, customerUUID)
		}
		s.mtx.Unlock()
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-s.Exit():
		return nil
	}
}

//
//
func (s *Daemon) UnlistenEvery(ctx context.Context, customerUUID string) error {
	s.mtx.RLock()
	_, ok := s.feeds[customerUUID]
	s.mtx.RUnlock()
	if !ok {
		return FaultListeningNegationDetected
	}

	select {
	case s.actions <- cmd{op: unlisten, customerUUID: customerUUID}:
		s.mtx.Lock()
		delete(s.feeds, customerUUID)
		s.mtx.Unlock()
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-s.Exit():
		return nil
	}
}

//
func (s *Daemon) CountCustomers() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.feeds)
}

//
func (s *Daemon) CountCustomerFeeds(customerUUID string) int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.feeds[customerUUID])
}

//
//
func (s *Daemon) Broadcast(ctx context.Context, msg any) error {
	return s.BroadcastUsingIncidents(ctx, msg, make(map[string][]string))
}

//
//
//
func (s *Daemon) BroadcastUsingIncidents(ctx context.Context, msg any, incidents map[string][]string) error {
	select {
	case s.actions <- cmd{op: pub, msg: msg, incidents: incidents}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-s.Exit():
		return nil
	}
}

//
func (s *Daemon) UponHalt() {
	s.actions <- cmd{op: halt}
}

//
type status struct {
	//
	feeds map[string]map[string]*Listening
	//
	inquiries map[string]*inquireAdditionPointerTally
}

//
//
type inquireAdditionPointerTally struct {
	q        Inquire
	pointerTally int
}

//
func (s *Daemon) UponInitiate() error {
	go s.cycle(status{
		feeds: make(map[string]map[string]*Listening),
		inquiries:       make(map[string]*inquireAdditionPointerTally),
	})
	return nil
}

//
func (s *Daemon) UponRestore() error {
	return nil
}

func (s *Daemon) cycle(status status) {
cycle:
	for cmd := range s.actions {
		switch cmd.op {
		case unlisten:
			if cmd.inquire != nil {
				status.discard(cmd.customerUUID, cmd.inquire.Text(), FaultUnlistened)
			} else {
				status.discardCustomer(cmd.customerUUID, FaultUnlistened)
			}
		case halt:
			status.discardEvery(nil)
			break cycle
		case sub:
			status.add(cmd.customerUUID, cmd.inquire, cmd.listening)
		case pub:
			if err := status.transmit(cmd.msg, cmd.incidents); err != nil {
				s.Tracer.Failure("REDACTED", "REDACTED", err)
			}
		}
	}
}

func (status *status) add(customerUUID string, q Inquire, listening *Listening) {
	queueTxt := q.Text()

	//
	if _, ok := status.feeds[queueTxt]; !ok {
		status.feeds[queueTxt] = make(map[string]*Listening)
	}
	//
	status.feeds[queueTxt][customerUUID] = listening

	//
	if _, ok := status.inquiries[queueTxt]; !ok {
		status.inquiries[queueTxt] = &inquireAdditionPointerTally{q: q, pointerTally: 0}
	}
	//
	status.inquiries[queueTxt].pointerTally++
}

func (status *status) discard(customerUUID string, queueTxt string, rationale error) {
	customerFeeds, ok := status.feeds[queueTxt]
	if !ok {
		return
	}

	listening, ok := customerFeeds[customerUUID]
	if !ok {
		return
	}

	listening.abort(rationale)

	//
	//
	delete(status.feeds[queueTxt], customerUUID)
	if len(status.feeds[queueTxt]) == 0 {
		delete(status.feeds, queueTxt)
	}

	//
	status.inquiries[queueTxt].pointerTally--
	//
	if status.inquiries[queueTxt].pointerTally == 0 {
		delete(status.inquiries, queueTxt)
	}
}

func (status *status) discardCustomer(customerUUID string, rationale error) {
	for queueTxt, customerFeeds := range status.feeds {
		if _, ok := customerFeeds[customerUUID]; ok {
			status.discard(customerUUID, queueTxt, rationale)
		}
	}
}

func (status *status) discardEvery(rationale error) {
	for queueTxt, customerFeeds := range status.feeds {
		for customerUUID := range customerFeeds {
			status.discard(customerUUID, queueTxt, rationale)
		}
	}
}

func (status *status) transmit(msg any, incidents map[string][]string) error {
	for queueTxt, customerFeeds := range status.feeds {
		q := status.inquiries[queueTxt].q

		align, err := q.Aligns(incidents)
		if err != nil {
			return fmt.Errorf("REDACTED", q.Text(), err)
		}

		if align {
			for customerUUID, listening := range customerFeeds {
				if cap(listening.out) == 0 {
					//
					listening.out <- FreshArtifact(msg, incidents)
				} else {
					//
					select {
					case listening.out <- FreshArtifact(msg, incidents):
					default:
						status.discard(customerUUID, queueTxt, FaultOutputBelongingVolume)
					}
				}
			}
		}
	}

	return nil
}
