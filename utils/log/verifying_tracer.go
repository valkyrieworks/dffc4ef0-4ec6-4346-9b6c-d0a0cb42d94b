package log

import (
	"io"
	"os"
	"testing"

	"github.com/go-kit/log/term"
)

//
var _verifyingtracer Tracer

//
//
//
//
//
//
func VerifyingTracer() Tracer {
	return VerifyingTracerUsingEmission(os.Stdout)
}

//
//
//
//
//
//
func VerifyingTracerUsingEmission(w io.Writer) Tracer {
	if _verifyingtracer != nil {
		return _verifyingtracer
	}

	if testing.Verbose() {
		_verifyingtracer = FreshTEMPTracer(FreshChronizePersistor(w))
	} else {
		_verifyingtracer = FreshNooperationTracer()
	}

	return _verifyingtracer
}

//
//
func VerifyingTracerUsingHueProc(hueProc func(tokvals ...any) term.FgBgColor) Tracer {
	if _verifyingtracer != nil {
		return _verifyingtracer
	}

	if testing.Verbose() {
		_verifyingtracer = FreshTEMPTracerUsingHueProc(FreshChronizePersistor(os.Stdout), hueProc)
	} else {
		_verifyingtracer = FreshNooperationTracer()
	}

	return _verifyingtracer
}
