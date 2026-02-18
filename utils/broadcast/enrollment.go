package broadcast

import (
	"errors"

	engineconnect "github.com/valkyrieworks/utils/align"
)

var (
	//
	ErrDeactivated = errors.New("REDACTED")

	//
	//
	ErrOutOfAbility = errors.New("REDACTED")
)

//
//
//
//
//
type Enrollment struct {
	out chan Signal

	revoked chan struct{}
	mtx      engineconnect.ReadwriteLock
	err      error
}

//
func NewEnrollment(outVolume int) *Enrollment {
	return &Enrollment{
		out:      make(chan Signal, outVolume),
		revoked: make(chan struct{}),
	}
}

//
//
//
func (s *Enrollment) Out() <-chan Signal {
	return s.out
}

//
//
func (s *Enrollment) Revoked() <-chan struct{} {
	return s.revoked
}

//
//
//
//
//
//
//
//
func (s *Enrollment) Err() error {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.err
}

func (s *Enrollment) revoke(err error) {
	s.mtx.Lock()
	s.err = err
	s.mtx.Unlock()
	close(s.revoked)
}

//
type Signal struct {
	data   any
	events map[string][]string
}

func NewSignal(data any, events map[string][]string) Signal {
	return Signal{data, events}
}

//
func (msg Signal) Data() any {
	return msg.data
}

//
func (msg Signal) Events() map[string][]string {
	return msg.events
}
