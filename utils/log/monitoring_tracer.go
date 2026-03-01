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
func FreshLoggingTracer(following Tracer) Tracer {
	return &monitoringTracer{
		following: following,
	}
}

type pylonMonitor interface {
	error
	PylonLogging() errors.StackTrace
}

type monitoringTracer struct {
	following Tracer
}

func (l *monitoringTracer) Details(msg string, tokvals ...any) {
	l.following.Details(msg, layoutFaults(tokvals)...)
}

func (l *monitoringTracer) Diagnose(msg string, tokvals ...any) {
	if ReportDiagnose {
		l.following.Diagnose(msg, layoutFaults(tokvals)...)
	}
}

func (l *monitoringTracer) Failure(msg string, tokvals ...any) {
	l.following.Failure(msg, layoutFaults(tokvals)...)
}

func (l *monitoringTracer) Using(tokvals ...any) Tracer {
	return &monitoringTracer{following: l.following.Using(layoutFaults(tokvals)...)}
}

func layoutFaults(tokvals []any) []any {
	freshTokvals := make([]any, len(tokvals))
	copy(freshTokvals, tokvals)
	for i := 0; i < len(freshTokvals)-1; i += 2 {
		if err, ok := freshTokvals[i+1].(pylonMonitor); ok {
			freshTokvals[i+1] = monitoredFailure{err}
		}
	}
	return freshTokvals
}

//
//
type monitoredFailure struct {
	encapsulated pylonMonitor
}

var _ pylonMonitor = monitoredFailure{}

func (t monitoredFailure) PylonLogging() errors.StackTrace {
	return t.encapsulated.PylonLogging()
}

func (t monitoredFailure) Reason() error {
	return t.encapsulated
}

func (t monitoredFailure) Failure() string {
	return fmt.Sprintf("REDACTED", t.encapsulated)
}
