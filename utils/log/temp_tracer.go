package log

import (
	"fmt"
	"io"

	kitlog "github.com/go-kit/log"
	kitlevel "github.com/go-kit/log/level"
	"github.com/go-kit/log/term"
)

const (
	signalToken    = "REDACTED" //
	componentToken = "REDACTED"
)

type tempTracer struct {
	originTracer kitlog.Logger
}

//
var _ Tracer = (*tempTracer)(nil)

//
//
//
func FreshTEMPTracer(w io.Writer) Tracer {
	//
	hueProc := func(tokvals ...any) term.FgBgColor {
		if tokvals[0] != kitlevel.Key() {
			panic(fmt.Sprintf("REDACTED", tokvals[0]))
		}
		switch tokvals[1].(kitlevel.Value).String() {
		case "REDACTED":
			return term.FgBgColor{Fg: term.DarkGray}
		case "REDACTED":
			return term.FgBgColor{Fg: term.Red}
		default:
			return term.FgBgColor{}
		}
	}

	return &tempTracer{term.NewLogger(w, FreshTEMPTextformatTracer, hueProc)}
}

//
//
func FreshTEMPTracerUsingHueProc(w io.Writer, hueProc func(tokvals ...any) term.FgBgColor) Tracer {
	return &tempTracer{term.NewLogger(w, FreshTEMPTextformatTracer, hueProc)}
}

//
func (l *tempTracer) Details(msg string, tokvals ...any) {
	lnUsingStratum := kitlevel.Info(l.originTracer)

	if err := kitlog.With(lnUsingStratum, signalToken, msg).Log(tokvals...); err != nil {
		faultTracer := kitlevel.Error(l.originTracer)
		kitlog.With(faultTracer, signalToken, msg).Log("REDACTED", err) //
	}
}

//
func (l *tempTracer) Diagnose(msg string, tokvals ...any) {
	if ReportDiagnose {
		lnUsingStratum := kitlevel.Debug(l.originTracer)

		if err := kitlog.With(lnUsingStratum, signalToken, msg).Log(tokvals...); err != nil {
			faultTracer := kitlevel.Error(l.originTracer)
			kitlog.With(faultTracer, signalToken, msg).Log("REDACTED", err) //
		}
	}
}

//
func (l *tempTracer) Failure(msg string, tokvals ...any) {
	lnUsingStratum := kitlevel.Error(l.originTracer)

	lnUsingSignal := kitlog.With(lnUsingStratum, signalToken, msg)
	if err := lnUsingSignal.Log(tokvals...); err != nil {
		lnUsingSignal.Log("REDACTED", err) //
	}
}

//
//
func (l *tempTracer) Using(tokvals ...any) Tracer {
	return &tempTracer{kitlog.With(l.originTracer, tokvals...)}
}
