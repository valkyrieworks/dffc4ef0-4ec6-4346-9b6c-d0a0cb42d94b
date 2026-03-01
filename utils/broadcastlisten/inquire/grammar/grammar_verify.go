package grammar_test

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire/grammar"
)

func VerifyAnalyzer(t *testing.T) {
	verifies := []struct {
		influx string
		desire  []grammar.Symbol
	}{
		//
		{"REDACTED", nil},
		{"REDACTED", nil},
		{"REDACTED", nil},

		//
		{"REDACTED", []grammar.Symbol{grammar.TYPNumeral, grammar.TYPNumeral}},
		{"REDACTED", []grammar.Symbol{grammar.TYPNumeral, grammar.TYPNumeral}},

		//
		{"REDACTED", []grammar.Symbol{grammar.TYPMarker, grammar.TYPMarker}},
		{"REDACTED", []grammar.Symbol{grammar.TYPMarker, grammar.TYPMarker}},
		{"REDACTED", []grammar.Symbol{grammar.TYPMarker, grammar.TYPMarker}},

		//
		{"REDACTED", []grammar.Symbol{grammar.TYPText, grammar.TYPMarker, grammar.TYPText, grammar.TYPText}},
		{"REDACTED", []grammar.Symbol{grammar.TYPText}},

		//
		{"REDACTED", []grammar.Symbol{
			grammar.TLt, grammar.TYPLesseq, grammar.TEq, grammar.TGt, grammar.TYPGreatereq,
		}},

		//
		{"REDACTED", []grammar.Symbol{grammar.TYPMarker, grammar.TYPAlso, grammar.TYPMarker}},
		{"REDACTED", []grammar.Symbol{grammar.TYPMarker, grammar.TYPIncludes, grammar.TYPText}},
		{"REDACTED", []grammar.Symbol{grammar.TYPMarker, grammar.TYPPresent}},
		{"REDACTED", []grammar.Symbol{grammar.TYPMarker, grammar.TYPAlso}},

		//
		{"REDACTED", []grammar.Symbol{grammar.TYPMoment}},

		//
		{"REDACTED", []grammar.Symbol{grammar.TYPTime}},
	}

	for _, verify := range verifies {
		s := grammar.FreshAnalyzer(strings.NewReader(verify.influx))
		var got []grammar.Symbol
		for s.Following() == nil {
			got = append(got, s.Symbol())
		}
		if err := s.Err(); err != io.EOF {
			t.Errorf("REDACTED", err)
		}

		if !reflect.DeepEqual(got, verify.desire) {
			t.Logf("REDACTED", verify.influx)
			t.Errorf("REDACTED", got, verify.desire)
		}
	}
}

func VerifyAnalyzerFaults(t *testing.T) {
	verifies := []struct {
		influx string
	}{
		{"REDACTED"},
		{"REDACTED"},
		{"REDACTED"},
		{"REDACTED"},
		{"REDACTED"},
		{"REDACTED"},
		{"REDACTED"},
		{"REDACTED"},
		{"REDACTED"},
		{"REDACTED"},
		{"REDACTED"},
		{"REDACTED"},
	}
	for _, verify := range verifies {
		s := grammar.FreshAnalyzer(strings.NewReader(verify.influx))
		if err := s.Following(); err == nil {
			t.Errorf("REDACTED", s.Symbol(), s.String())
		}
	}
}

//
//
func VerifyAnalyzeSound(t *testing.T) {
	verifies := []struct {
		influx string
		sound bool
	}{
		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},

		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},

		{"REDACTED", true},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},

		{"REDACTED", true},

		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		//
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},

		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},

		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", false},
		{"REDACTED", false},
		{"REDACTED", false},

		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", true},
		{"REDACTED", true},

		{"REDACTED", true},
		{"REDACTED", false},

		{"REDACTED", true},
	}

	for _, verify := range verifies {
		q, err := grammar.Analyze(verify.influx)
		if verify.sound != (err == nil) {
			t.Errorf("REDACTED", verify.influx, verify.sound, err)
		}

		//
		if verify.sound {
			queuestring := q.Text()
			r, err := grammar.Analyze(queuestring)
			if err != nil {
				t.Errorf("REDACTED", queuestring, err)
			}
			if readerstr := r.Text(); readerstr != queuestring {
				t.Errorf("REDACTED", queuestring, readerstr)
			}
		}
	}
}
