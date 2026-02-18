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
package broadcast

import (
	"context"
	"errors"
	"fmt"

	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
)

type process int

const (
	sub process = iota
	pub
	unreg
	terminate
)

var (
	//
	//
	ErrEnrollmentNegateLocated = errors.New("REDACTED")

	//
	//
	ErrYetActivated = errors.New("REDACTED")
)

//
//
//
//
//
//
type Inquire interface {
	Aligns(events map[string][]string) (bool, error)
	String() string
}

type cmd struct {
	op process

	//
	inquire        Inquire
	enrollment *Enrollment
	customerUID     string

	//
	msg    any
	events map[string][]string
}

//
//
type Host struct {
	daemon.RootDaemon

	orders    chan cmd
	ordersCeiling int

	//
	//
	mtx           engineconnect.ReadwriteLock
	registrations map[string]map[string]struct{} //
}

//
type Setting func(*Host)

//
//
//
func NewHost(options ...Setting) *Host {
	s := &Host{
		registrations: make(map[string]map[string]struct{}),
	}
	s.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", s)

	for _, setting := range options {
		setting(s)
	}

	//
	s.orders = make(chan cmd, s.ordersCeiling)

	return s
}

//
//
//
//
func BufferVolume(cap int) Setting {
	return func(s *Host) {
		if cap > 0 {
			s.ordersCeiling = cap
		}
	}
}

//
func (s *Host) BufferVolume() int {
	return s.ordersCeiling
}

//
//
//
//
//
//
//
//
func (s *Host) Enrol(
	ctx context.Context,
	customerUID string,
	inquire Inquire,
	outVolume ...int,
) (*Enrollment, error) {
	outCeiling := 1
	if len(outVolume) > 0 {
		if outVolume[0] <= 0 {
			panic("REDACTED")
		}
		outCeiling = outVolume[0]
	}

	return s.enrol(ctx, customerUID, inquire, outCeiling)
}

//
//
//
func (s *Host) EnrolUnbuffered(ctx context.Context, customerUID string, inquire Inquire) (*Enrollment, error) {
	return s.enrol(ctx, customerUID, inquire, 0)
}

func (s *Host) enrol(ctx context.Context, customerUID string, inquire Inquire, outVolume int) (*Enrollment, error) {
	s.mtx.RLock()
	customerRegistrations, ok := s.registrations[customerUID]
	if ok {
		_, ok = customerRegistrations[inquire.String()]
	}
	s.mtx.RUnlock()
	if ok {
		return nil, ErrYetActivated
	}

	enrollment := NewEnrollment(outVolume)
	select {
	case s.orders <- cmd{op: sub, customerUID: customerUID, inquire: inquire, enrollment: enrollment}:
		s.mtx.Lock()
		if _, ok = s.registrations[customerUID]; !ok {
			s.registrations[customerUID] = make(map[string]struct{})
		}
		s.registrations[customerUID][inquire.String()] = struct{}{}
		s.mtx.Unlock()
		return enrollment, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-s.Exit():
		return nil, errors.New("REDACTED")
	}
}

//
//
//
func (s *Host) Deenroll(ctx context.Context, customerUID string, inquire Inquire) error {
	s.mtx.RLock()
	customerRegistrations, ok := s.registrations[customerUID]
	if ok {
		_, ok = customerRegistrations[inquire.String()]
	}
	s.mtx.RUnlock()
	if !ok {
		return ErrEnrollmentNegateLocated
	}

	select {
	case s.orders <- cmd{op: unreg, customerUID: customerUID, inquire: inquire}:
		s.mtx.Lock()
		delete(customerRegistrations, inquire.String())
		if len(customerRegistrations) == 0 {
			delete(s.registrations, customerUID)
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
func (s *Host) DeenrollAll(ctx context.Context, customerUID string) error {
	s.mtx.RLock()
	_, ok := s.registrations[customerUID]
	s.mtx.RUnlock()
	if !ok {
		return ErrEnrollmentNegateLocated
	}

	select {
	case s.orders <- cmd{op: unreg, customerUID: customerUID}:
		s.mtx.Lock()
		delete(s.registrations, customerUID)
		s.mtx.Unlock()
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-s.Exit():
		return nil
	}
}

//
func (s *Host) CountAgents() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.registrations)
}

//
func (s *Host) CountCustomerRegistrations(customerUID string) int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.registrations[customerUID])
}

//
//
func (s *Host) Broadcast(ctx context.Context, msg any) error {
	return s.BroadcastWithEvents(ctx, msg, make(map[string][]string))
}

//
//
//
func (s *Host) BroadcastWithEvents(ctx context.Context, msg any, events map[string][]string) error {
	select {
	case s.orders <- cmd{op: pub, msg: msg, events: events}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-s.Exit():
		return nil
	}
}

//
func (s *Host) OnHalt() {
	s.orders <- cmd{op: terminate}
}

//
type status struct {
	//
	registrations map[string]map[string]*Enrollment
	//
	inquiries map[string]*inquirePlusReferenceNumber
}

//
//
type inquirePlusReferenceNumber struct {
	q        Inquire
	referenceNumber int
}

//
func (s *Host) OnBegin() error {
	go s.cycle(status{
		registrations: make(map[string]map[string]*Enrollment),
		inquiries:       make(map[string]*inquirePlusReferenceNumber),
	})
	return nil
}

//
func (s *Host) OnRestore() error {
	return nil
}

func (s *Host) cycle(status status) {
cycle:
	for cmd := range s.orders {
		switch cmd.op {
		case unreg:
			if cmd.inquire != nil {
				status.delete(cmd.customerUID, cmd.inquire.String(), ErrDeactivated)
			} else {
				status.deleteCustomer(cmd.customerUID, ErrDeactivated)
			}
		case terminate:
			status.deleteAll(nil)
			break cycle
		case sub:
			status.add(cmd.customerUID, cmd.inquire, cmd.enrollment)
		case pub:
			if err := status.transmit(cmd.msg, cmd.events); err != nil {
				s.Tracer.Fault("REDACTED", "REDACTED", err)
			}
		}
	}
}

func (status *status) add(customerUID string, q Inquire, enrollment *Enrollment) {
	qStr := q.String()

	//
	if _, ok := status.registrations[qStr]; !ok {
		status.registrations[qStr] = make(map[string]*Enrollment)
	}
	//
	status.registrations[qStr][customerUID] = enrollment

	//
	if _, ok := status.inquiries[qStr]; !ok {
		status.inquiries[qStr] = &inquirePlusReferenceNumber{q: q, referenceNumber: 0}
	}
	//
	status.inquiries[qStr].referenceNumber++
}

func (status *status) delete(customerUID string, qStr string, cause error) {
	customerRegistrations, ok := status.registrations[qStr]
	if !ok {
		return
	}

	enrollment, ok := customerRegistrations[customerUID]
	if !ok {
		return
	}

	enrollment.revoke(cause)

	//
	//
	delete(status.registrations[qStr], customerUID)
	if len(status.registrations[qStr]) == 0 {
		delete(status.registrations, qStr)
	}

	//
	status.inquiries[qStr].referenceNumber--
	//
	if status.inquiries[qStr].referenceNumber == 0 {
		delete(status.inquiries, qStr)
	}
}

func (status *status) deleteCustomer(customerUID string, cause error) {
	for qStr, customerRegistrations := range status.registrations {
		if _, ok := customerRegistrations[customerUID]; ok {
			status.delete(customerUID, qStr, cause)
		}
	}
}

func (status *status) deleteAll(cause error) {
	for qStr, customerRegistrations := range status.registrations {
		for customerUID := range customerRegistrations {
			status.delete(customerUID, qStr, cause)
		}
	}
}

func (status *status) transmit(msg any, events map[string][]string) error {
	for qStr, customerRegistrations := range status.registrations {
		q := status.inquiries[qStr].q

		align, err := q.Aligns(events)
		if err != nil {
			return fmt.Errorf("REDACTED", q.String(), err)
		}

		if align {
			for customerUID, enrollment := range customerRegistrations {
				if cap(enrollment.out) == 0 {
					//
					enrollment.out <- NewSignal(msg, events)
				} else {
					//
					select {
					case enrollment.out <- NewSignal(msg, events):
					default:
						status.delete(customerUID, qStr, ErrOutOfAbility)
					}
				}
			}
		}
	}

	return nil
}
