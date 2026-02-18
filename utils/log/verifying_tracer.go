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
	return VerifyingTracerWithResult(os.Stdout)
}

//
//
//
//
//
//
func VerifyingTracerWithResult(w io.Writer) Tracer {
	if _verifyingtracer != nil {
		return _verifyingtracer
	}

	if testing.Verbose() {
		_verifyingtracer = NewTMTracer(NewAlignRecorder(w))
	} else {
		_verifyingtracer = NewNoopTracer()
	}

	return _verifyingtracer
}

//
//
func VerifyingTracerWithHueFn(hueFn func(keyvalues ...any) term.FgBgColor) Tracer {
	if _verifyingtracer != nil {
		return _verifyingtracer
	}

	if testing.Verbose() {
		_verifyingtracer = NewTMTracerWithHueFn(NewAlignRecorder(os.Stdout), hueFn)
	} else {
		_verifyingtracer = NewNoopTracer()
	}

	return _verifyingtracer
}
