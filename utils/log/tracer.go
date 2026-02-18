package log

import (
	"io"

	kitlog "github.com/go-kit/log"
)

//
type Tracer interface {
	Diagnose(msg string, keyvalues ...any)
	Details(msg string, keyvalues ...any)
	Fault(msg string, keyvalues ...any)

	With(keyvalues ...any) Tracer
}

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
func NewAlignRecorder(w io.Writer) io.Writer {
	return kitlog.NewSyncWriter(w)
}
