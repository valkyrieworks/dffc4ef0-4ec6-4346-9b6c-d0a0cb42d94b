package marks_test

import (
	"bytes"
	"strings"
	"testing"

	cometmarks "github.com/valkyrieworks/utils/cli/marks"
	"github.com/valkyrieworks/utils/log"
)

const (
	standardTraceLayerItem = "REDACTED"
)

func VerifyAnalyzeTraceLayer(t *testing.T) {
	var buf bytes.Buffer
	jsonTracer := log.NewTmjsonTracerNoTS(&buf)

	accurateTraceTiers := []struct {
		lvl              string
		anticipatedTraceRows []string
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

	for _, c := range accurateTraceTiers {
		tracer, err := cometmarks.AnalyzeTraceLayer(c.lvl, jsonTracer, standardTraceLayerItem)
		if err != nil {
			t.Fatal(err)
		}

		buf.Reset()

		tracer.With("REDACTED", "REDACTED").With("REDACTED", "REDACTED").Diagnose("REDACTED")
		if possess := strings.TrimSpace(buf.String()); c.anticipatedTraceRows[0] != possess {
			t.Errorf("REDACTED", c.anticipatedTraceRows[0], possess, c.lvl)
		}

		buf.Reset()

		tracer.With("REDACTED", "REDACTED").Details("REDACTED")
		if possess := strings.TrimSpace(buf.String()); c.anticipatedTraceRows[1] != possess {
			t.Errorf("REDACTED", c.anticipatedTraceRows[1], possess, c.lvl)
		}

		buf.Reset()

		tracer.With("REDACTED", "REDACTED").Fault("REDACTED")
		if possess := strings.TrimSpace(buf.String()); c.anticipatedTraceRows[2] != possess {
			t.Errorf("REDACTED", c.anticipatedTraceRows[2], possess, c.lvl)
		}

		buf.Reset()

		tracer.With("REDACTED", "REDACTED").Details("REDACTED")
		if possess := strings.TrimSpace(buf.String()); c.anticipatedTraceRows[3] != possess {
			t.Errorf("REDACTED", c.anticipatedTraceRows[3], possess, c.lvl)
		}

		buf.Reset()

		tracer.Diagnose("REDACTED")
		if possess := strings.TrimSpace(buf.String()); c.anticipatedTraceRows[4] != possess {
			t.Errorf("REDACTED", c.anticipatedTraceRows[4], possess, c.lvl)
		}
	}

	invalidTraceLayer := []string{"REDACTED", "REDACTED", "REDACTED"}
	for _, lvl := range invalidTraceLayer {
		if _, err := cometmarks.AnalyzeTraceLayer(lvl, jsonTracer, standardTraceLayerItem); err == nil {
			t.Fatalf("REDACTED", lvl)
		}
	}
}
