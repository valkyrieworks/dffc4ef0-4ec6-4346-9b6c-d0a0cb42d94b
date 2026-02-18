package log

import (
	"io"

	kitlog "github.com/go-kit/log"
)

//
//
//
//
func NewTmjsonTracer(w io.Writer) Tracer {
	tracer := kitlog.NewJSONLogger(w)
	tracer = kitlog.With(tracer, "REDACTED", kitlog.DefaultTimestampUTC)
	return &tmTracer{tracer}
}

//
//
func NewTmjsonTracerNoTS(w io.Writer) Tracer {
	tracer := kitlog.NewJSONLogger(w)
	return &tmTracer{tracer}
}
