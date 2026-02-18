package merkle

import (
	"crypto/sha256"
	"strings"
	"testing"
)

var drain any

type deeperDigestVerify struct {
	left, correct string
}

var deeperDigestVerifies = []*deeperDigestVerify{
	{"REDACTED", "REDACTED"},
	{"REDACTED", "REDACTED"},
	{"REDACTED", "REDACTED"},
	{"REDACTED", "REDACTED"},
	{"REDACTED", "REDACTED"},
	{strings.Repeat("REDACTED", 1<<10), strings.Repeat("REDACTED", 4<<10)},
	{strings.Repeat("REDACTED", sha256.Size), strings.Repeat("REDACTED", 10<<10)},
	{"REDACTED", "REDACTED"},
}

func CriterionDeeperDigest(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, tt := range deeperDigestVerifies {
			got := deeperDigest([]byte(tt.left), []byte(tt.correct))
			if g, w := len(got), sha256.Size; g != w {
				b.Fatalf("REDACTED", g, w)
			}
			drain = got
		}
	}

	if drain == nil {
		b.Fatal("REDACTED")
	}
}

//
//
//
func CriterionElementDigest64kb(b *testing.B) {
	b.ReportAllocs()
	element := make([]byte, 64*1024)
	digest := sha256.New()

	for i := 0; i < b.N; i++ {
		element[0] = byte(i)
		got := elementDigestOption(digest, element)
		if g, w := len(got), sha256.Size; g != w {
			b.Fatalf("REDACTED", g, w)
		}
		drain = got
	}

	if drain == nil {
		b.Fatal("REDACTED")
	}
}
