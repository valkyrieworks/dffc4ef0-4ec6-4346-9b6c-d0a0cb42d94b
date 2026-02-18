package inquire_test

import (
	"testing"

	"github.com/valkyrieworks/utils/broadcast/inquire"
)

const verifyInquire = "REDACTED"

var verifyEvents = map[string][]string{
	"REDACTED": {
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED", "REDACTED",
	},
}

func CriterionAnalyzeBespoke(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := inquire.New(verifyInquire)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func CriterionAlignBespoke(b *testing.B) {
	q, err := inquire.New(verifyInquire)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok, err := q.Aligns(verifyEvents)
		if err != nil {
			b.Fatal(err)
		} else if !ok {
			b.Error("REDACTED")
		}
	}
}
