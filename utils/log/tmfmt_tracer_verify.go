package trace_t_test

import (
	"bytes"
	"errors"
	"io"
	"math"
	"regexp"
	"testing"

	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/assert"

	"github.com/valkyrieworks/utils/log"
)

func VerifyTMFmtTracer(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	tracer := log.NewTMFmtTracer(buf)

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

func CriterionTMFmtTracerBasic(b *testing.B) {
	criterionExecutorKittrace(b, log.NewTMFmtTracer(io.Discard), rootSignal)
}

func CriterionTMFmtTracerSituational(b *testing.B) {
	criterionExecutorKittrace(b, log.NewTMFmtTracer(io.Discard), withSignal)
}

func VerifyTMFmtTracerParallelism(t *testing.T) {
	t.Parallel()
	verifyParallelism(t, log.NewTMFmtTracer(io.Discard), 10000)
}

func criterionExecutorKittrace(b *testing.B, tracer kitlog.Logger, f func(kitlog.Logger)) {
	lc := kitlog.With(tracer, "REDACTED", "REDACTED")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(lc)
	}
}

var (
	rootSignal = func(tracer kitlog.Logger) { tracer.Log("REDACTED", "REDACTED") }          //
	withSignal = func(tracer kitlog.Logger) { kitlog.With(tracer, "REDACTED", "REDACTED").Log("REDACTED", "REDACTED") } //
)

//

func verifyParallelism(t *testing.T, tracer kitlog.Logger, sum int) {
	n := int(math.Sqrt(float64(sum)))
	allocate := sum / n

	errC := make(chan error, n)

	for i := 0; i < n; i++ {
		go func() {
			errC <- junk(tracer, allocate)
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errC
		if err != nil {
			t.Fatalf("REDACTED", err)
		}
	}
}

func junk(tracer kitlog.Logger, tally int) error {
	for i := 0; i < tally; i++ {
		err := tracer.Log("REDACTED", i)
		if err != nil {
			return err
		}
	}
	return nil
}

type myindex map[int]int

func (m myindex) String() string { return "REDACTED" }
