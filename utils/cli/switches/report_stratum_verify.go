package switches_test

import (
	"bytes"
	"strings"
	"testing"

	strongmindflags "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/cli/switches"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

const (
	fallbackReportStratumDatum = "REDACTED"
)

func VerifyAnalyzeReportStratum(t *testing.T) {
	var buf bytes.Buffer
	jsnTracer := log.FreshTempjsonTracerNegativeTimestamp(&buf)

	preciseReportTiers := []struct {
		lvl              string
		anticipatedReportTraces []string
	}{
		{"REDACTED", []string{
			"REDACTED", //
			"REDACTED",
			"REDACTED",
			"REDACTED", //
			"REDACTED",
		}},

		{"REDACTED", []string{
			"REDACTED",
			"REDACTED",
			"REDACTED",
			"REDACTED",
			"REDACTED",
		}},

		{"REDACTED", []string{
			"REDACTED",
			"REDACTED",
			"REDACTED",
			"REDACTED",
			"REDACTED",
		}},
	}

	for _, c := range preciseReportTiers {
		tracer, err := strongmindflags.AnalyzeRecordStratum(c.lvl, jsnTracer, fallbackReportStratumDatum)
		if err != nil {
			t.Fatal(err)
		}

		buf.Reset()

		tracer.Using("REDACTED", "REDACTED").Using("REDACTED", "REDACTED").Diagnose("REDACTED")
		if possess := strings.TrimSpace(buf.String()); c.anticipatedReportTraces[0] != possess {
			t.Errorf("REDACTED", c.anticipatedReportTraces[0], possess, c.lvl)
		}

		buf.Reset()

		tracer.Using("REDACTED", "REDACTED").Details("REDACTED")
		if possess := strings.TrimSpace(buf.String()); c.anticipatedReportTraces[1] != possess {
			t.Errorf("REDACTED", c.anticipatedReportTraces[1], possess, c.lvl)
		}

		buf.Reset()

		tracer.Using("REDACTED", "REDACTED").Failure("REDACTED")
		if possess := strings.TrimSpace(buf.String()); c.anticipatedReportTraces[2] != possess {
			t.Errorf("REDACTED", c.anticipatedReportTraces[2], possess, c.lvl)
		}

		buf.Reset()

		tracer.Using("REDACTED", "REDACTED").Details("REDACTED")
		if possess := strings.TrimSpace(buf.String()); c.anticipatedReportTraces[3] != possess {
			t.Errorf("REDACTED", c.anticipatedReportTraces[3], possess, c.lvl)
		}

		buf.Reset()

		tracer.Diagnose("REDACTED")
		if possess := strings.TrimSpace(buf.String()); c.anticipatedReportTraces[4] != possess {
			t.Errorf("REDACTED", c.anticipatedReportTraces[4], possess, c.lvl)
		}
	}

	impreciseReportStratum := []string{"REDACTED", "REDACTED", "REDACTED"}
	for _, lvl := range impreciseReportStratum {
		if _, err := strongmindflags.AnalyzeRecordStratum(lvl, jsnTracer, fallbackReportStratumDatum); err == nil {
			t.Fatalf("REDACTED", lvl)
		}
	}
}
