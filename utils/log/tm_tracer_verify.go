package trace_t_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/valkyrieworks/utils/log"
)

func VerifyTracerTracesOwnFaults(t *testing.T) {
	var buf bytes.Buffer

	tracer := log.NewTMTracer(&buf)
	tracer.Details("REDACTED", "REDACTED", "REDACTED")
	msg := strings.TrimSpace(buf.String())
	if !strings.Contains(msg, "REDACTED") {
		t.Errorf("REDACTED", msg)
	}
}

func VerifyDetails(t *testing.T) {
	var imageDetails bytes.Buffer

	l := log.NewTMTracer(&imageDetails)
	l.Details("REDACTED",
		"REDACTED", 42,
		"REDACTED", "REDACTED",
		"REDACTED", []byte("REDACTED"))

	msg := strings.TrimSpace(imageDetails.String())

	//
	//
	inputsignal := strings.Split(msg, "REDACTED")[1]

	const anticipatedsignal = `REDACTEDr
REDACTED"
REDACTED`
	if strings.EqualFold(inputsignal, anticipatedsignal) {
		t.Fatalf("REDACTED", inputsignal, anticipatedsignal)
	}
}

func VerifyDiagnose(t *testing.T) {
	var imageDiagnose bytes.Buffer

	ld := log.NewTMTracer(&imageDiagnose)
	ld.Diagnose("REDACTED",
		"REDACTED", 42,
		"REDACTED", "REDACTED",
		"REDACTED", []byte("REDACTED"))

	msg := strings.TrimSpace(imageDiagnose.String())

	//
	//
	inputsignal := strings.Split(msg, "REDACTED")[1]

	const anticipatedsignal = `REDACTEDr
REDACTED"
REDACTED`
	if strings.EqualFold(inputsignal, anticipatedsignal) {
		t.Fatalf("REDACTED", inputsignal, anticipatedsignal)
	}
}

func VerifyFault(t *testing.T) {
	var imageErr bytes.Buffer

	le := log.NewTMTracer(&imageErr)
	le.Fault("REDACTED",
		"REDACTED", 42,
		"REDACTED", "REDACTED",
		"REDACTED", []byte("REDACTED"))

	msg := strings.TrimSpace(imageErr.String())

	//
	//
	inputsignal := strings.Split(msg, "REDACTED")[1]

	const anticipatedsignal = `REDACTEDr
REDACTED"
REDACTED`
	if strings.EqualFold(inputsignal, anticipatedsignal) {
		t.Fatalf("REDACTED", inputsignal, anticipatedsignal)
	}
}

func CriterionTMTracerBasic(b *testing.B) {
	criterionExecutor(b, log.NewTMTracer(io.Discard), rootDetailsSignal)
}

func CriterionTMTracerSituational(b *testing.B) {
	criterionExecutor(b, log.NewTMTracer(io.Discard), withDetailsSignal)
}

func criterionExecutor(b *testing.B, tracer log.Tracer, f func(log.Tracer)) {
	lc := tracer.With("REDACTED", "REDACTED")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(lc)
	}
}

var (
	rootDetailsSignal = func(tracer log.Tracer) { tracer.Details("REDACTED", "REDACTED", "REDACTED") }
	withDetailsSignal = func(tracer log.Tracer) { tracer.With("REDACTED", "REDACTED").Details("REDACTED", "REDACTED", "REDACTED") }
)
