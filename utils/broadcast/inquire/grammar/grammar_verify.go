package syntax_test

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/valkyrieworks/utils/broadcast/inquire/grammar"
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
		{"REDACTED", []grammar.Symbol{grammar.TAmount, grammar.TAmount}},
		{"REDACTED", []grammar.Symbol{grammar.TAmount, grammar.TAmount}},

		//
		{"REDACTED", []grammar.Symbol{grammar.TMarker, grammar.TMarker}},
		{"REDACTED", []grammar.Symbol{grammar.TMarker, grammar.TMarker}},
		{"REDACTED", []grammar.Symbol{grammar.TMarker, grammar.TMarker}},

		//
		{"REDACTED", []grammar.Symbol{grammar.TString, grammar.TMarker, grammar.TString, grammar.TString}},
		{"REDACTED", []grammar.Symbol{grammar.TString}},

		//
		{"REDACTED", []grammar.Symbol{
			grammar.TLt, grammar.TLeq, grammar.TEq, grammar.TGt, grammar.TGeq,
		}},

		//
		{"REDACTED", []grammar.Symbol{grammar.TMarker, grammar.TAnd, grammar.TMarker}},
		{"REDACTED", []grammar.Symbol{grammar.TMarker, grammar.TIncludes, grammar.TString}},
		{"REDACTED", []grammar.Symbol{grammar.TMarker, grammar.TPresent}},
		{"REDACTED", []grammar.Symbol{grammar.TMarker, grammar.TAnd}},

		//
		{"REDACTED", []grammar.Symbol{grammar.TTime}},

		//
		{"REDACTED", []grammar.Symbol{grammar.TDate}},
	}

	for _, verify := range verifies {
		s := grammar.NewAnalyzer(strings.NewReader(verify.influx))
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
		s := grammar.NewAnalyzer(strings.NewReader(verify.influx))
		if err := s.Following(); err == nil {
			t.Errorf("REDACTED", s.Symbol(), s.Content())
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
			qstring := q.String()
			r, err := grammar.Analyze(qstring)
			if err != nil {
				t.Errorf("REDACTED", qstring, err)
			}
			if rstring := r.String(); rstring != qstring {
				t.Errorf("REDACTED", qstring, rstring)
			}
		}
	}
}
