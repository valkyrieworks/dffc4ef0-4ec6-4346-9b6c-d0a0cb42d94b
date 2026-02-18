package comethash

import (
	"bytes"
	"crypto/sha256"
	"strings"
	"testing"
)

var drain any

var numerousSegments = []struct {
	label string
	in   [][]byte
	desire [32]byte
}{
	{
		label: "REDACTED",
		in:   [][]byte{[]byte("REDACTED"), []byte("REDACTED")},
		desire: sha256.Sum256(nil),
	},
	{
		label: "REDACTED",
		in:   [][]byte{[]byte("REDACTED"), []byte("REDACTED"), []byte("REDACTED")},
		desire: sha256.Sum256([]byte("REDACTED")),
	},
	{
		label: "REDACTED",
		in:   [][]byte{bytes.Repeat([]byte("REDACTED"), 1<<10), []byte("REDACTED"), bytes.Repeat([]byte("REDACTED"), 100)},
		desire: sha256.Sum256([]byte(strings.Repeat("REDACTED", 1<<10) + "REDACTED" + strings.Repeat("REDACTED", 100))),
	},
}

func CriterionSha256numerous(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, tt := range numerousSegments {
			got := TotalNumerous(tt.in[0], tt.in[1:]...)
			if !bytes.Equal(got, tt.desire[:]) {
				b.Fatalf("REDACTED", tt.label, got, tt.desire)
			}
			drain = got
		}
	}

	if drain == nil {
		b.Fatal("REDACTED")
	}

	drain = nil
}
