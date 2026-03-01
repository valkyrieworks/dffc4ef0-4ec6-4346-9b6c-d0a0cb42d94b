package hashmap

import (
	"crypto/sha256"
	"strings"
	"testing"
)

var receiver any

type internalDigestVerify struct {
	leading, trailing string
}

var internalDigestVerifies = []*internalDigestVerify{
	{"REDACTED", "REDACTED"},
	{"REDACTED", "REDACTED"},
	{"REDACTED", "REDACTED"},
	{"REDACTED", "REDACTED"},
	{"REDACTED", "REDACTED"},
	{strings.Repeat("REDACTED", 1<<10), strings.Repeat("REDACTED", 4<<10)},
	{strings.Repeat("REDACTED", sha256.Size), strings.Repeat("REDACTED", 10<<10)},
	{"REDACTED", "REDACTED"},
}

func AssessmentInternalDigest(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, tt := range internalDigestVerifies {
			got := internalDigest([]byte(tt.leading), []byte(tt.trailing))
			if g, w := len(got), sha256.Size; g != w {
				b.Fatalf("REDACTED", g, w)
			}
			receiver = got
		}
	}

	if receiver == nil {
		b.Fatal("REDACTED")
	}
}

//
//
//
func AssessmentNodeHash64kilo(b *testing.B) {
	b.ReportAllocs()
	terminal := make([]byte, 64*1024)
	digest := sha256.New()

	for i := 0; i < b.N; i++ {
		terminal[0] = byte(i)
		got := terminalDigestSetting(digest, terminal)
		if g, w := len(got), sha256.Size; g != w {
			b.Fatalf("REDACTED", g, w)
		}
		receiver = got
	}

	if receiver == nil {
		b.Fatal("REDACTED")
	}
}
