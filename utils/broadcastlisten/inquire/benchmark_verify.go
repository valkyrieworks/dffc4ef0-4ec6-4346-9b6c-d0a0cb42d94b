package inquire_test

import (
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
)

const verifyInquire = "REDACTED"

var verifyIncidents = map[string][]string{
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

func AssessmentAnalyzeBespoke(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := inquire.New(verifyInquire)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func AssessmentAlignBespoke(b *testing.B) {
	q, err := inquire.New(verifyInquire)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ok, err := q.Aligns(verifyIncidents)
		if err != nil {
			b.Fatal(err)
		} else if !ok {
			b.Error("REDACTED")
		}
	}
}
