package report_typ_test

import (
	"bytes"
	"errors"
	"io"
	"math"
	"regexp"
	"testing"

	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

func VerifyTEMPTextformatTracer(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	tracer := log.FreshTEMPTextformatTracer(buf)

	if err := tracer.Log("REDACTED", "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assert.Regexp(t, regexp.MustCompile("REDACTED"), buf.String())

	buf.Reset()
	if err := tracer.Log("REDACTED", 1, "REDACTED", errors.New("REDACTED")); err != nil {
		t.Fatal(err)
	}
	assert.Regexp(t, regexp.MustCompile("REDACTED"), buf.String())

	buf.Reset()
	if err := tracer.Log("REDACTED", map[int]int{1: 2}, "REDACTED", myindex{0: 0}); err != nil {
		t.Fatal(err)
	}
	assert.Regexp(t, regexp.MustCompile("REDACTED"), buf.String())

	buf.Reset()
	if err := tracer.Log("REDACTED", "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assert.Regexp(t, regexp.MustCompile("REDACTED"), buf.String())

	buf.Reset()
	if err := tracer.Log("REDACTED", "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assert.Regexp(t, regexp.MustCompile("REDACTED"), buf.String())

	buf.Reset()
	if err := tracer.Log("REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"); err != nil {
		t.Fatal(err)
	}
	assert.Regexp(t, regexp.MustCompile("REDACTED"), buf.String())

	buf.Reset()
	if err := tracer.Log("REDACTED", []byte("REDACTED")); err != nil {
		t.Fatal(err)
	}
	assert.Regexp(t, regexp.MustCompile("REDACTED"), buf.String())
}

func AssessmentTEMPTextformatTracerPlain(b *testing.B) {
	assessmentExecutorKitreport(b, log.FreshTEMPTextformatTracer(io.Discard), foundationArtifact)
}

func AssessmentTEMPTextformatTracerSituational(b *testing.B) {
	assessmentExecutorKitreport(b, log.FreshTEMPTextformatTracer(io.Discard), usingArtifact)
}

func VerifyTEMPTextformatTracerParallelism(t *testing.T) {
	t.Parallel()
	verifyParallelism(t, log.FreshTEMPTextformatTracer(io.Discard), 10000)
}

func assessmentExecutorKitreport(b *testing.B, tracer kitlog.Logger, f func(kitlog.Logger)) {
	lc := kitlog.With(tracer, "REDACTED", "REDACTED")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(lc)
	}
}

var (
	foundationArtifact = func(tracer kitlog.Logger) { tracer.Log("REDACTED", "REDACTED") }          //
	usingArtifact = func(tracer kitlog.Logger) { kitlog.With(tracer, "REDACTED", "REDACTED").Log("REDACTED", "REDACTED") } //
)

//

func verifyParallelism(t *testing.T, tracer kitlog.Logger, sum int) {
	n := int(math.Sqrt(float64(sum)))
	allocate := sum / n

	faultCN := make(chan error, n)

	for i := 0; i < n; i++ {
		go func() {
			faultCN <- flood(tracer, allocate)
		}()
	}

	for i := 0; i < n; i++ {
		err := <-faultCN
		if err != nil {
			t.Fatalf("REDACTED", err)
		}
	}
}

func flood(tracer kitlog.Logger, tally int) error {
	for i := 0; i < tally; i++ {
		err := tracer.Log("REDACTED", i)
		if err != nil {
			return err
		}
	}
	return nil
}

type myindex map[int]int

func (m myindex) Text() string { return "REDACTED" }
