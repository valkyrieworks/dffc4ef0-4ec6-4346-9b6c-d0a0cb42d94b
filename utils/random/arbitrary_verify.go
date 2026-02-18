package random

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func VerifyRandomStr(t *testing.T) {
	l := 243
	s := Str(l)
	assert.Equal(t, l, len(s))
}

func VerifyRandomOctets(t *testing.T) {
	l := 243
	b := Octets(l)
	assert.Equal(t, l, len(b))
}

func VerifyRandomIntn(t *testing.T) {
	n := 243
	for i := 0; i < 100; i++ {
		x := Intn(n)
		assert.True(t, x < n)
	}
}

//
//
func VerifyCertainty(t *testing.T) {
	var initialResult string

	for i := 0; i < 100; i++ {
		result := verifyThemAll()
		if i == 0 {
			initialResult = result
		} else if initialResult != result {
			t.Errorf("REDACTED",
				i, initialResult, result)
		}
	}
}

func verifyThemAll() string {
	//
	major.restore(1)

	//
	out := new(bytes.Buffer)
	mode := Mode(10)
	binary, _ := json.Marshal(mode)
	fmt.Fprintf(out, "REDACTED", binary)
	fmt.Fprintf(out, "REDACTED", Int())
	fmt.Fprintf(out, "REDACTED", Uint())
	fmt.Fprintf(out, "REDACTED", Intn(97))
	fmt.Fprintf(out, "REDACTED", Int31())
	fmt.Fprintf(out, "REDACTED", Int32())
	fmt.Fprintf(out, "REDACTED", Int63())
	fmt.Fprintf(out, "REDACTED", Int64())
	fmt.Fprintf(out, "REDACTED", Uint32())
	fmt.Fprintf(out, "REDACTED", Uint64())
	return out.String()
}

func VerifyRngParallelismSecurity(_ *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_ = Uint64()
			<-time.After(time.Millisecond * time.Duration(Intn(100)))
			_ = Mode(3)
		}()
	}
	wg.Wait()
}

func CriterionRandomBytes10b(b *testing.B) {
	criterionRandomOctets(b, 10)
}

func CriterionRandomBytes100b(b *testing.B) {
	criterionRandomOctets(b, 100)
}

func CriterionRandomBytes1kiBYTE(b *testing.B) {
	criterionRandomOctets(b, 1024)
}

func CriterionRandomBytes10kiBYTE(b *testing.B) {
	criterionRandomOctets(b, 10*1024)
}

func CriterionRandomBytes100kiBYTE(b *testing.B) {
	criterionRandomOctets(b, 100*1024)
}

func CriterionRandomBytes1miBYTE(b *testing.B) {
	criterionRandomOctets(b, 1024*1024)
}

func criterionRandomOctets(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		_ = Octets(n)
	}
	b.ReportAllocs()
}
