package log

import (
	"fmt"

	"github.com/pkg/errors"
)

//
//
//
//
//
//
//
func NewLoggingTracer(following Tracer) Tracer {
	return &loggingTracer{
		following: following,
	}
}

type pileTracker interface {
	error
	PileTrack() errors.StackTrace
}

type loggingTracer struct {
	following Tracer
}

func (l *loggingTracer) Details(msg string, keyvalues ...any) {
	l.following.Details(msg, layoutFaults(keyvalues)...)
}

func (l *loggingTracer) Diagnose(msg string, keyvalues ...any) {
	if TraceDiagnose {
		l.following.Diagnose(msg, layoutFaults(keyvalues)...)
	}
}

func (l *loggingTracer) Fault(msg string, keyvalues ...any) {
	l.following.Fault(msg, layoutFaults(keyvalues)...)
}

func (l *loggingTracer) With(keyvalues ...any) Tracer {
	return &loggingTracer{following: l.following.With(layoutFaults(keyvalues)...)}
}

func layoutFaults(keyvalues []any) []any {
	newKeyvalues := make([]any, len(keyvalues))
	copy(newKeyvalues, keyvalues)
	for i := 0; i < len(newKeyvalues)-1; i += 2 {
		if err, ok := newKeyvalues[i+1].(pileTracker); ok {
			newKeyvalues[i+1] = trackedFault{err}
		}
	}
	return newKeyvalues
}

//
//
type trackedFault struct {
	encapsulated pileTracker
}

var _ pileTracker = trackedFault{}

func (t trackedFault) PileTrack() errors.StackTrace {
	return t.encapsulated.PileTrack()
}

func (t trackedFault) Origin() error {
	return t.encapsulated
}

func (t trackedFault) Fault() string {
	return fmt.Sprintf("REDACTED", t.encapsulated)
}
