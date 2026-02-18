package json__test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/json"
)

func VerifySerialize(t *testing.T) {
	s := "REDACTED"
	sPointer := &s
	i64 := int64(64)
	ti := time.Date(2020, 6, 2, 18, 5, 13, 4346374, time.FixedZone("REDACTED", 2*60*60))
	car := &Car{Revolutions: 4}
	vessel := Vessel{Navigate: true}

	verifyscenarios := map[string]struct {
		item  any
		result string
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
		"REDACTED":       {BespokePointer{Item: "REDACTED"}, "REDACTED"}, //
		"REDACTED":   {&BespokePointer{Item: "REDACTED"}, "REDACTED"},
		"REDACTED":     {BespokeItem{Item: "REDACTED"}, "REDACTED"},
		"REDACTED": {&BespokeItem{Item: "REDACTED"}, "REDACTED"},
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
		"REDACTED":       {Automobile(car), "REDACTED"},
		"REDACTED":         {(*Car)(nil), "REDACTED"},
		"REDACTED":            {vessel, "REDACTED"},
		"REDACTED":        {&vessel, "REDACTED"},
		"REDACTED":      {Automobile(vessel), "REDACTED"},
		"REDACTED":      {PublicKey{1, 2, 3, 4, 5, 6, 7, 8}, "REDACTED"},
		"REDACTED": {
			Markers{JSONLabel: "REDACTED", IgnoreEmpty: "REDACTED", Concealed: "REDACTED", Markers: &Markers{JSONLabel: "REDACTED"}},
			"REDACTED",
		},
		"REDACTED": {Markers{}, "REDACTED"},
		//
		//
		//
		//
		"REDACTED": {
			Struct{
				Bool: true, Float64: 3.14, Int32: 32, Int64: 64, Int64pointer: &i64,
				String: "REDACTED", StringPointerPointer: &sPointer, Octets: []byte{1, 2, 3},
				Time: ti, Car: car, Vessel: vessel, Automobiles: []Automobile{car, vessel},
				Offspring: &Struct{Bool: false, String: "REDACTED"}, internal: "REDACTED",
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
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			bz, err := json.Serialize(tc.item)
			require.NoError(t, err)
			assert.JSONEq(t, tc.result, string(bz))
		})
	}
}

func CriterionJsonSerializeStruct(b *testing.B) {
	s := "REDACTED"
	sPointer := &s
	i64 := int64(64)
	ti := time.Date(2020, 6, 2, 18, 5, 13, 4346374, time.FixedZone("REDACTED", 2*60*60))
	car := &Car{Revolutions: 4}
	vessel := Vessel{Navigate: true}
	for i := 0; i < b.N; i++ {
		_, _ = json.Serialize(Struct{
			Bool: true, Float64: 3.14, Int32: 32, Int64: 64, Int64pointer: &i64,
			String: "REDACTED", StringPointerPointer: &sPointer, Octets: []byte{1, 2, 3},
			Time: ti, Car: car, Vessel: vessel, Automobiles: []Automobile{car, vessel},
			Offspring: &Struct{Bool: false, String: "REDACTED"}, internal: "REDACTED",
		})
	}
}
