package log

import (
	"io"

	kitlog "github.com/go-kit/log"
)

//
type Tracer interface {
	Diagnose(msg string, tokvals ...any)
	Details(msg string, tokvals ...any)
	Failure(msg string, tokvals ...any)

	Using(tokvals ...any) Tracer
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
func FreshChronizePersistor(w io.Writer) io.Writer {
	return kitlog.NewSyncWriter(w)
}
