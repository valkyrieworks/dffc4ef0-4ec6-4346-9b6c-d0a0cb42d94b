package log

import (
	"fmt"
	"io"

	kitlog "github.com/go-kit/log"
	kitlevel "github.com/go-kit/log/level"
	"github.com/go-kit/log/term"
)

const (
	messageKey    = "REDACTED" //
	componentKey = "REDACTED"
)

type tmTracer struct {
	originTracer kitlog.Logger
}

//
var _ Tracer = (*tmTracer)(nil)

//
//
//
func NewTMTracer(w io.Writer) Tracer {
	//
	hueFn := func(keyvalues ...any) term.FgBgColor {
		if keyvalues[0] != kitlevel.Key() {
			panic(fmt.Sprintf("REDACTED", keyvalues[0]))
		}
		switch keyvalues[1].(kitlevel.Value).String() {
		case "REDACTED":
			return term.FgBgColor{Fg: term.DarkGray}
		case "REDACTED":
			return term.FgBgColor{Fg: term.Red}
		default:
			return term.FgBgColor{}
		}
	}

	return &tmTracer{term.NewLogger(w, NewTMFmtTracer, hueFn)}
}

//
//
func NewTMTracerWithHueFn(w io.Writer, hueFn func(keyvalues ...any) term.FgBgColor) Tracer {
	return &tmTracer{term.NewLogger(w, NewTMFmtTracer, hueFn)}
}

//
func (l *tmTracer) Details(msg string, keyvalues ...any) {
	lWithLayer := kitlevel.Info(l.originTracer)

	if err := kitlog.With(lWithLayer, messageKey, msg).Log(keyvalues...); err != nil {
		errTracer := kitlevel.Error(l.originTracer)
		kitlog.With(errTracer, messageKey, msg).Log("REDACTED", err) //
	}
}

//
func (l *tmTracer) Diagnose(msg string, keyvalues ...any) {
	if TraceDiagnose {
		lWithLayer := kitlevel.Debug(l.originTracer)

		if err := kitlog.With(lWithLayer, messageKey, msg).Log(keyvalues...); err != nil {
			errTracer := kitlevel.Error(l.originTracer)
			kitlog.With(errTracer, messageKey, msg).Log("REDACTED", err) //
		}
	}
}

//
func (l *tmTracer) Fault(msg string, keyvalues ...any) {
	lWithLayer := kitlevel.Error(l.originTracer)

	lWithMessage := kitlog.With(lWithLayer, messageKey, msg)
	if err := lWithMessage.Log(keyvalues...); err != nil {
		lWithMessage.Log("REDACTED", err) //
	}
}

//
//
func (l *tmTracer) With(keyvalues ...any) Tracer {
	return &tmTracer{kitlog.With(l.originTracer, keyvalues...)}
}
