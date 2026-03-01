package report_typ_test

import (
	"bytes"
	stderr "errors"
	"fmt"
	"strings"
	"testing"

	"github.com/pkg/errors"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

func VerifyMonitoringTracer(t *testing.T) {
	var buf bytes.Buffer

	tracer := log.FreshTempjsonTracerNegativeTimestamp(&buf)

	tracer1 := log.FreshLoggingTracer(tracer)
	faultone := errors.New("REDACTED")
	fault2 := errors.New("REDACTED")
	tracer1.Using("REDACTED", faultone).Details("REDACTED", "REDACTED", fault2)

	desire := strings.ReplaceAll(
		strings.ReplaceAll(
			"REDACTED"+
				fmt.Sprintf("REDACTED", faultone)+
				"REDACTED"+
				fmt.Sprintf("REDACTED", fault2)+
				"REDACTED",
			"REDACTED", "REDACTED",
		), "REDACTED", "REDACTED")
	possess := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(buf.String()), "REDACTED", "REDACTED"), "REDACTED", "REDACTED")
	if desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}

	buf.Reset()

	tracer.Using(
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

	tracer.Using("REDACTED", "REDACTED").Using("REDACTED", "REDACTED").Details("REDACTED", "REDACTED", "REDACTED")

	desire = "REDACTED"
	possess = strings.TrimSpace(buf.String())
	if desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}
}
