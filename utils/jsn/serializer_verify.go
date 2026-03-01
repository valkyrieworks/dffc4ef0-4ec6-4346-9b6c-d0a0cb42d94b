package jsn__test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
)

func VerifySerialize(t *testing.T) {
	s := "REDACTED"
	strReference := &s
	i64 := int64(64)
	ti := time.Date(2020, 6, 2, 18, 5, 13, 4346374, time.FixedZone("REDACTED", 2*60*60))
	car := &Car{Rotations: 4}
	vessel := Vessel{Navigate: true}

	verifycases := map[string]struct {
		datum  any
		emission string
	}{
		"REDACTED":             {nil, "REDACTED"},
		"REDACTED":          {"REDACTED", "REDACTED"},
		"REDACTED":         {float32(3.14), "REDACTED"},
		"REDACTED":     {float32(-3.14), "REDACTED"},
		"REDACTED":         {float64(3.14), "REDACTED"},
		"REDACTED":     {float64(-3.14), "REDACTED"},
		"REDACTED":           {int32(32), "REDACTED"},
		"REDACTED":           {int64(64), "REDACTED"},
		"REDACTED":       {int64(-64), "REDACTED"},
		"REDACTED":       {&i64, "REDACTED"},
		"REDACTED":          {uint64(64), "REDACTED"},
		"REDACTED":            {ti, "REDACTED"},
		"REDACTED":      {time.Time{}, "REDACTED"},
		"REDACTED":        {&ti, "REDACTED"},
		"REDACTED":       {BespokeReference{Datum: "REDACTED"}, "REDACTED"}, //
		"REDACTED":   {&BespokeReference{Datum: "REDACTED"}, "REDACTED"},
		"REDACTED":     {BespokeDatum{Datum: "REDACTED"}, "REDACTED"},
		"REDACTED": {&BespokeDatum{Datum: "REDACTED"}, "REDACTED"},
		"REDACTED":       {[]int(nil), "REDACTED"},
		"REDACTED":     {[]int{}, "REDACTED"},
		"REDACTED":     {[]byte{1, 2, 3}, "REDACTED"},
		"REDACTED":     {[]int64{1, 2, 3}, "REDACTED"},
		"REDACTED": {[]*int64{&i64, nil}, "REDACTED"},
		"REDACTED":     {[3]byte{1, 2, 3}, "REDACTED"},
		"REDACTED":     {[3]int64{1, 2, 3}, "REDACTED"},
		"REDACTED":         {map[string]int64(nil), "REDACTED"}, //
		"REDACTED":       {map[string]int64{}, "REDACTED"},
		"REDACTED":       {map[string]int64{"REDACTED": 1, "REDACTED": 2, "REDACTED": 3}, "REDACTED"},
		"REDACTED":             {car, "REDACTED"},
		"REDACTED":       {*car, "REDACTED"},
		"REDACTED":       {Carriage(car), "REDACTED"},
		"REDACTED":         {(*Car)(nil), "REDACTED"},
		"REDACTED":            {vessel, "REDACTED"},
		"REDACTED":        {&vessel, "REDACTED"},
		"REDACTED":      {Carriage(vessel), "REDACTED"},
		"REDACTED":      {CommonToken{1, 2, 3, 4, 5, 6, 7, 8}, "REDACTED"},
		"REDACTED": {
			Labels{JSNAlias: "REDACTED", ExcludeBlank: "REDACTED", Concealed: "REDACTED", Labels: &Labels{JSNAlias: "REDACTED"}},
			"REDACTED",
		},
		"REDACTED": {Labels{}, "REDACTED"},
		//
		//
		//
		//
		"REDACTED": {
			Schema{
				Flag: true, Float64: 3.14, Integer32: 32, Int64n: 64, Int64reference: &i64,
				Text: "REDACTED", TextReferenceReference: &strReference, Octets: []byte{1, 2, 3},
				Moment: ti, Car: car, Vessel: vessel, Carriages: []Carriage{car, vessel},
				Offspring: &Schema{Flag: false, Text: "REDACTED"}, secluded: "REDACTED",
			},
			`REDACTED{
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED[
REDACTED,
REDACTED}
REDACTED,
REDACTED{
REDACTED,
REDACTED,
REDACTED,
REDACTEDl
REDACTED}
REDACTED`,
		},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			bz, err := jsn.Serialize(tc.datum)
			require.NoError(t, err)
			assert.JSONEq(t, tc.emission, string(bz))
		})
	}
}

func AssessmentJsnSerializeSchema(b *testing.B) {
	s := "REDACTED"
	strReference := &s
	i64 := int64(64)
	ti := time.Date(2020, 6, 2, 18, 5, 13, 4346374, time.FixedZone("REDACTED", 2*60*60))
	car := &Car{Rotations: 4}
	vessel := Vessel{Navigate: true}
	for i := 0; i < b.N; i++ {
		_, _ = jsn.Serialize(Schema{
			Flag: true, Float64: 3.14, Integer32: 32, Int64n: 64, Int64reference: &i64,
			Text: "REDACTED", TextReferenceReference: &strReference, Octets: []byte{1, 2, 3},
			Moment: ti, Car: car, Vessel: vessel, Carriages: []Carriage{car, vessel},
			Offspring: &Schema{Flag: false, Text: "REDACTED"}, secluded: "REDACTED",
		})
	}
}
