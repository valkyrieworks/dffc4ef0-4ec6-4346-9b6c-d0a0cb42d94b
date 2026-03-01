package tenderminthash

import (
	"bytes"
	"crypto/sha256"
	"strings"
	"testing"
)

var receiver any

var multipleSegments = []struct {
	alias string
	in   [][]byte
	desire [32]byte
}{
	{
		alias: "REDACTED",
		in:   [][]byte{[]byte("REDACTED"), []byte("REDACTED")},
		desire: sha256.Sum256(nil),
	},
	{
		alias: "REDACTED",
		in:   [][]byte{[]byte("REDACTED"), []byte("REDACTED"), []byte("REDACTED")},
		desire: sha256.Sum256([]byte("REDACTED")),
	},
	{
		alias: "REDACTED",
		in:   [][]byte{bytes.Repeat([]byte("REDACTED"), 1<<10), []byte("REDACTED"), bytes.Repeat([]byte("REDACTED"), 100)},
		desire: sha256.Sum256([]byte(strings.Repeat("REDACTED", 1<<10) + "REDACTED" + strings.Repeat("REDACTED", 100))),
	},
}

func AssessmentHash256series(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, tt := range multipleSegments {
			got := TotalMultiple(tt.in[0], tt.in[1:]...)
			if !bytes.Equal(got, tt.desire[:]) {
				b.Fatalf("REDACTED", tt.alias, got, tt.desire)
			}
			receiver = got
		}
	}

	if receiver == nil {
		b.Fatal("REDACTED")
	}

	receiver = nil
}
