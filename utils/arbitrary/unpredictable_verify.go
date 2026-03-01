package arbitrary

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func VerifyArbitraryTxt(t *testing.T) {
	l := 243
	s := Str(l)
	assert.Equal(t, l, len(s))
}

func VerifyArbitraryOctets(t *testing.T) {
	l := 243
	b := Octets(l)
	assert.Equal(t, l, len(b))
}

func VerifyArbitraryIntegern(t *testing.T) {
	n := 243
	for i := 0; i < 100; i++ {
		x := Integern(n)
		assert.True(t, x < n)
	}
}

//
//
func VerifyPredictability(t *testing.T) {
	var initialEmission string

	for i := 0; i < 100; i++ {
		emission := verifyThoseEvery()
		if i == 0 {
			initialEmission = emission
		} else if initialEmission != emission {
			t.Errorf("REDACTED",
				i, initialEmission, emission)
		}
	}
}

func verifyThoseEvery() string {
	//
	majestic.restore(1)

	//
	out := new(bytes.Buffer)
	mode := Mode(10)
	chunk, _ := json.Marshal(mode)
	fmt.Fprintf(out, "REDACTED", chunk)
	fmt.Fprintf(out, "REDACTED", Int())
	fmt.Fprintf(out, "REDACTED", Uintn())
	fmt.Fprintf(out, "REDACTED", Integern(97))
	fmt.Fprintf(out, "REDACTED", Int31n())
	fmt.Fprintf(out, "REDACTED", Integer32())
	fmt.Fprintf(out, "REDACTED", Int63n())
	fmt.Fprintf(out, "REDACTED", Int64n())
	fmt.Fprintf(out, "REDACTED", Uint32n())
	fmt.Fprintf(out, "REDACTED", Uint64n())
	return out.String()
}

func VerifyRandomizerParallelismSecurity(_ *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_ = Uint64n()
			<-time.After(time.Millisecond * time.Duration(Integern(100)))
			_ = Mode(3)
		}()
	}
	wg.Wait()
}

func AssessmentArbitraryOctets10b(b *testing.B) {
	assessmentArbitraryOctets(b, 10)
}

func AssessmentArbitraryOctets100b(b *testing.B) {
	assessmentArbitraryOctets(b, 100)
}

func AssessmentArbitraryOctets1kiBYTE(b *testing.B) {
	assessmentArbitraryOctets(b, 1024)
}

func AssessmentArbitraryOctets10kiBYTE(b *testing.B) {
	assessmentArbitraryOctets(b, 10*1024)
}

func AssessmentArbitraryOctets100kiBYTE(b *testing.B) {
	assessmentArbitraryOctets(b, 100*1024)
}

func AssessmentArbitraryOctets1miBYTE(b *testing.B) {
	assessmentArbitraryOctets(b, 1024*1024)
}

func assessmentArbitraryOctets(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		_ = Octets(n)
	}
	b.ReportAllocs()
}
