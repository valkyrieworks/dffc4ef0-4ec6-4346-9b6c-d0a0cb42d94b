package report_typ_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

func VerifyTracerRecordsTheirFaults(t *testing.T) {
	var buf bytes.Buffer

	tracer := log.FreshTEMPTracer(&buf)
	tracer.Details("REDACTED", "REDACTED", "REDACTED")
	msg := strings.TrimSpace(buf.String())
	if !strings.Contains(msg, "REDACTED") {
		t.Errorf("REDACTED", msg)
	}
}

func VerifyDetails(t *testing.T) {
	var areaDetails bytes.Buffer

	l := log.FreshTEMPTracer(&areaDetails)
	l.Details("REDACTED",
		"REDACTED", 42,
		"REDACTED", "REDACTED",
		"REDACTED", []byte("REDACTED"))

	msg := strings.TrimSpace(areaDetails.String())

	//
	//
	acceptedartifact := strings.Split(msg, "REDACTED")[1]

	const anticipatedartifact = `REDACTEDr
REDACTED"
REDACTED`
	if strings.EqualFold(acceptedartifact, anticipatedartifact) {
		t.Fatalf("REDACTED", acceptedartifact, anticipatedartifact)
	}
}

func VerifyDiagnose(t *testing.T) {
	var areaDiagnose bytes.Buffer

	ld := log.FreshTEMPTracer(&areaDiagnose)
	ld.Diagnose("REDACTED",
		"REDACTED", 42,
		"REDACTED", "REDACTED",
		"REDACTED", []byte("REDACTED"))

	msg := strings.TrimSpace(areaDiagnose.String())

	//
	//
	acceptedartifact := strings.Split(msg, "REDACTED")[1]

	const anticipatedartifact = `REDACTEDr
REDACTED"
REDACTED`
	if strings.EqualFold(acceptedartifact, anticipatedartifact) {
		t.Fatalf("REDACTED", acceptedartifact, anticipatedartifact)
	}
}

func VerifyFailure(t *testing.T) {
	var areaFault bytes.Buffer

	le := log.FreshTEMPTracer(&areaFault)
	le.Failure("REDACTED",
		"REDACTED", 42,
		"REDACTED", "REDACTED",
		"REDACTED", []byte("REDACTED"))

	msg := strings.TrimSpace(areaFault.String())

	//
	//
	acceptedartifact := strings.Split(msg, "REDACTED")[1]

	const anticipatedartifact = `REDACTEDr
REDACTED"
REDACTED`
	if strings.EqualFold(acceptedartifact, anticipatedartifact) {
		t.Fatalf("REDACTED", acceptedartifact, anticipatedartifact)
	}
}

func AssessmentTEMPTracerPlain(b *testing.B) {
	assessmentExecutor(b, log.FreshTEMPTracer(io.Discard), foundationDetailsArtifact)
}

func AssessmentTEMPTracerSituational(b *testing.B) {
	assessmentExecutor(b, log.FreshTEMPTracer(io.Discard), usingDetailsArtifact)
}

func assessmentExecutor(b *testing.B, tracer log.Tracer, f func(log.Tracer)) {
	lc := tracer.Using("REDACTED", "REDACTED")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(lc)
	}
}

var (
	foundationDetailsArtifact = func(tracer log.Tracer) { tracer.Details("REDACTED", "REDACTED", "REDACTED") }
	usingDetailsArtifact = func(tracer log.Tracer) { tracer.Using("REDACTED", "REDACTED").Details("REDACTED", "REDACTED", "REDACTED") }
)
