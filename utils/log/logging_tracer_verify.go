package trace_t_test

import (
	"bytes"
	stderr "errors"
	"fmt"
	"strings"
	"testing"

	"github.com/pkg/errors"

	"github.com/valkyrieworks/utils/log"
)

func VerifyLoggingTracer(t *testing.T) {
	var buf bytes.Buffer

	tracer := log.NewTmjsonTracerNoTS(&buf)

	tracer1 := log.NewLoggingTracer(tracer)
	fault1 := errors.New("REDACTED")
	err2 := errors.New("REDACTED")
	tracer1.With("REDACTED", fault1).Details("REDACTED", "REDACTED", err2)

	desire := strings.ReplaceAll(
		strings.ReplaceAll(
			"REDACTED"+
				fmt.Sprintf("REDACTED", fault1)+
				"REDACTED"+
				fmt.Sprintf("REDACTED", err2)+
				"REDACTED",
			"REDACTED", "REDACTED",
		), "REDACTED", "REDACTED")
	possess := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(buf.String()), "REDACTED", "REDACTED"), "REDACTED", "REDACTED")
	if desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}

	buf.Reset()

	tracer.With(
		"REDACTED", stderr.New("REDACTED"),
	).Details(
		"REDACTED", "REDACTED", stderr.New("REDACTED"),
	)

	desire = "REDACTED" +
		"REDACTED" +
		"REDACTED" +
		"REDACTED"
	possess = strings.TrimSpace(buf.String())
	if desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}

	buf.Reset()

	tracer.With("REDACTED", "REDACTED").With("REDACTED", "REDACTED").Details("REDACTED", "REDACTED", "REDACTED")

	desire = "REDACTED"
	possess = strings.TrimSpace(buf.String())
	if desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}
}
