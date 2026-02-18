package trace_t_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/valkyrieworks/utils/log"
)

func VerifyDiverseTiers(t *testing.T) {
	verifyScenarios := []struct {
		label    string
		permitted log.Setting
		desire    string
	}{
		{
			"REDACTED",
			log.PermitAll(),
			strings.Join([]string{
				"REDACTED",
				"REDACTED",
				"REDACTED",
			}, "REDACTED"),
		},
		{
			"REDACTED",
			log.PermitDiagnose(),
			strings.Join([]string{
				"REDACTED",
				"REDACTED",
				"REDACTED",
			}, "REDACTED"),
		},
		{
			"REDACTED",
			log.PermitDetails(),
			strings.Join([]string{
				"REDACTED",
				"REDACTED",
			}, "REDACTED"),
		},
		{
			"REDACTED",
			log.PermitFault(),
			strings.Join([]string{
				"REDACTED",
			}, "REDACTED"),
		},
		{
			"REDACTED",
			log.PermitNone(),
			"REDACTED",
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.label, func(t *testing.T) {
			var buf bytes.Buffer
			tracer := log.NewRefine(log.NewTmjsonTracerNoTS(&buf), tc.permitted)

			tracer.Diagnose("REDACTED", "REDACTED", "REDACTED")
			tracer.Details("REDACTED", "REDACTED", "REDACTED")
			tracer.Fault("REDACTED", "REDACTED", "REDACTED")

			if desire, possess := tc.desire, strings.TrimSpace(buf.String()); desire != possess {
				t.Errorf("REDACTED", desire, possess)
			}
		})
	}
}

func VerifyLayerContext(t *testing.T) {
	var buf bytes.Buffer

	tracer := log.NewTmjsonTracerNoTS(&buf)
	tracer = log.NewRefine(tracer, log.PermitFault())
	tracer = tracer.With("REDACTED", "REDACTED")

	tracer.Fault("REDACTED", "REDACTED", "REDACTED")

	desire := "REDACTED"
	possess := strings.TrimSpace(buf.String())
	if desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}

	buf.Reset()
	tracer.Details("REDACTED", "REDACTED", "REDACTED")
	if desire, possess := "REDACTED", strings.TrimSpace(buf.String()); desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}
}

func VerifyDiversePermitWith(t *testing.T) {
	var buf bytes.Buffer

	tracer := log.NewTmjsonTracerNoTS(&buf)

	tracer1 := log.NewRefine(tracer, log.PermitFault(), log.PermitDetailsWith("REDACTED", "REDACTED"))
	tracer1.With("REDACTED", "REDACTED").Details("REDACTED", "REDACTED", "REDACTED")

	desire := "REDACTED"
	possess := strings.TrimSpace(buf.String())
	if desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}

	buf.Reset()

	tracer2 := log.NewRefine(
		tracer,
		log.PermitFault(),
		log.PermitDetailsWith("REDACTED", "REDACTED"),
		log.PermitNoneWith("REDACTED", "REDACTED"),
	)

	tracer2.With("REDACTED", "REDACTED", "REDACTED", "REDACTED").Details("REDACTED", "REDACTED", "REDACTED")
	if desire, possess := "REDACTED", strings.TrimSpace(buf.String()); desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}

	buf.Reset()

	tracer3 := log.NewRefine(
		tracer,
		log.PermitFault(),
		log.PermitDetailsWith("REDACTED", "REDACTED"),
		log.PermitNoneWith("REDACTED", "REDACTED"),
	)

	tracer3.With("REDACTED", "REDACTED").With("REDACTED", "REDACTED").Details("REDACTED", "REDACTED", "REDACTED")

	desire = "REDACTED"
	possess = strings.TrimSpace(buf.String())
	if desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}
}
