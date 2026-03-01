package broadcastlisten

import (
	"errors"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

var (
	//
	FaultUnlistened = errors.New("REDACTED")

	//
	//
	FaultOutputBelongingVolume = errors.New("REDACTED")
)

//
//
//
//
//
type Listening struct {
	out chan Signal

	aborted chan struct{}
	mtx      commitchronize.ReadwriteExclusion
	err      error
}

//
func FreshListening(outputVolume int) *Listening {
	return &Listening{
		out:      make(chan Signal, outputVolume),
		aborted: make(chan struct{}),
	}
}

//
//
//
func (s *Listening) Out() <-chan Signal {
	return s.out
}

//
//
func (s *Listening) Aborted() <-chan struct{} {
	return s.aborted
}

//
//
//
//
//
//
//
//
func (s *Listening) Err() error {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.err
}

func (s *Listening) abort(err error) {
	s.mtx.Lock()
	s.err = err
	s.mtx.Unlock()
	close(s.aborted)
}

//
type Signal struct {
	data   any
	incidents map[string][]string
}

func FreshArtifact(data any, incidents map[string][]string) Signal {
	return Signal{data, incidents}
}

//
func (msg Signal) Data() any {
	return msg.data
}

//
func (msg Signal) Incidents() map[string][]string {
	return msg.incidents
}
