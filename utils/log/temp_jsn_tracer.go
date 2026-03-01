package log

import (
	"io"

	kitlog "github.com/go-kit/log"
)

//
//
//
//
func FreshTempjsonTracer(w io.Writer) Tracer {
	tracer := kitlog.NewJSONLogger(w)
	tracer = kitlog.With(tracer, "REDACTED", kitlog.DefaultTimestampUTC)
	return &tempTracer{tracer}
}

//
//
func FreshTempjsonTracerNegativeTimestamp(w io.Writer) Tracer {
	tracer := kitlog.NewJSONLogger(w)
	return &tempTracer{tracer}
}
