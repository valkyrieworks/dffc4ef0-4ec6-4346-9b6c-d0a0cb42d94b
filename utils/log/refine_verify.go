package report_typ_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

func VerifySundryTiers(t *testing.T) {
	verifyScenarios := []struct {
		alias    string
		permitted log.Selection
		desire    string
	}{
		{
			"REDACTED",
			log.PermitEvery(),
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
			log.PermitFailure(),
			strings.Join([]string{
				"REDACTED",
			}, "REDACTED"),
		},
		{
			"REDACTED",
			log.PermitNil(),
			"REDACTED",
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.alias, func(t *testing.T) {
			var buf bytes.Buffer
			tracer := log.FreshRefine(log.FreshTempjsonTracerNegativeTimestamp(&buf), tc.permitted)

			tracer.Diagnose("REDACTED", "REDACTED", "REDACTED")
			tracer.Details("REDACTED", "REDACTED", "REDACTED")
			tracer.Failure("REDACTED", "REDACTED", "REDACTED")

			if desire, possess := tc.desire, strings.TrimSpace(buf.String()); desire != possess {
				t.Errorf("REDACTED", desire, possess)
			}
		})
	}
}

func VerifyStratumEnv(t *testing.T) {
	var buf bytes.Buffer

	tracer := log.FreshTempjsonTracerNegativeTimestamp(&buf)
	tracer = log.FreshRefine(tracer, log.PermitFailure())
	tracer = tracer.Using("REDACTED", "REDACTED")

	tracer.Failure("REDACTED", "REDACTED", "REDACTED")

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

func VerifySundryPermitUsing(t *testing.T) {
	var buf bytes.Buffer

	tracer := log.FreshTempjsonTracerNegativeTimestamp(&buf)

	tracer1 := log.FreshRefine(tracer, log.PermitFailure(), log.PermitDetailsUsing("REDACTED", "REDACTED"))
	tracer1.Using("REDACTED", "REDACTED").Details("REDACTED", "REDACTED", "REDACTED")

	desire := "REDACTED"
	possess := strings.TrimSpace(buf.String())
	if desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}

	buf.Reset()

	tracer2 := log.FreshRefine(
		tracer,
		log.PermitFailure(),
		log.PermitDetailsUsing("REDACTED", "REDACTED"),
		log.PermitNilUsing("REDACTED", "REDACTED"),
	)

	tracer2.Using("REDACTED", "REDACTED", "REDACTED", "REDACTED").Details("REDACTED", "REDACTED", "REDACTED")
	if desire, possess := "REDACTED", strings.TrimSpace(buf.String()); desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}

	buf.Reset()

	tracer3 := log.FreshRefine(
		tracer,
		log.PermitFailure(),
		log.PermitDetailsUsing("REDACTED", "REDACTED"),
		log.PermitNilUsing("REDACTED", "REDACTED"),
	)

	tracer3.Using("REDACTED", "REDACTED").Using("REDACTED", "REDACTED").Details("REDACTED", "REDACTED", "REDACTED")

	desire = "REDACTED"
	possess = strings.TrimSpace(buf.String())
	if desire != possess {
		t.Errorf("REDACTED", desire, possess)
	}
}
