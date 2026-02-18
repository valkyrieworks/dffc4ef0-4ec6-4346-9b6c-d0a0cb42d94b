package json__test

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/utils/json"
)

func VerifyUnserialize(t *testing.T) {
	i64null := (*int64)(nil)
	str := "REDACTED"
	strPointer := &str
	structNull := (*Struct)(nil)
	i32 := int32(32)
	i64 := int64(64)

	verifyscenarios := map[string]struct {
		json  string
		item any
		err   bool
	}{
		"REDACTED":           {"REDACTED", true, false},
		"REDACTED":          {"REDACTED", false, false},
		"REDACTED":             {"REDACTED", float32(3.14), false},
		"REDACTED":             {"REDACTED", float64(3.14), false},
		"REDACTED":               {"REDACTED", int32(32), false},
		"REDACTED":        {"REDACTED", int32(32), true},
		"REDACTED":           {"REDACTED", &i32, false},
		"REDACTED":               {"REDACTED", int64(64), false},
		"REDACTED":         {"REDACTED", int64(64), true},
		"REDACTED":        {"REDACTED", int64(64), true},
		"REDACTED":           {"REDACTED", &i64, false},
		"REDACTED":       {"REDACTED", i64null, false},
		"REDACTED":              {"REDACTED", "REDACTED", false},
		"REDACTED":        {"REDACTED", "REDACTED", true},
		"REDACTED":          {"REDACTED", &str, false},
		"REDACTED":          {"REDACTED", []byte{1, 2, 3}, false},
		"REDACTED":         {"REDACTED", [][]byte{{1, 2, 3}}, false},
		"REDACTED":         {"REDACTED", []int32{1, 2, 3}, false},
		"REDACTED":         {"REDACTED", []int64{1, 2, 3}, false},
		"REDACTED":  {"REDACTED", []int64{1, 2, 3}, true},
		"REDACTED":     {"REDACTED", []*int64{&i64}, false},
		"REDACTED":   {"REDACTED", []int64(nil), false},
		"REDACTED":    {"REDACTED", []int64(nil), false},
		"REDACTED":          {"REDACTED", [3]byte{1, 2, 3}, false},
		"REDACTED":    {"REDACTED", [4]byte{1, 2, 3, 4}, true},
		"REDACTED":    {"REDACTED", [2]byte{1, 2}, true},
		"REDACTED":         {"REDACTED", [3]int32{1, 2, 3}, false},
		"REDACTED":         {"REDACTED", [3]int64{1, 2, 3}, false},
		"REDACTED":  {"REDACTED", [3]int64{1, 2, 3}, true},
		"REDACTED":   {"REDACTED", [4]int64{1, 2, 3, 4}, true},
		"REDACTED":   {"REDACTED", [2]int64{1, 2}, true},
		"REDACTED":           {"REDACTED", map[string][]byte{"REDACTED": {1, 2, 3}}, false},
		"REDACTED":           {"REDACTED", map[string]int32{"REDACTED": 1, "REDACTED": 2}, false},
		"REDACTED":           {"REDACTED", map[string]int64{"REDACTED": 1, "REDACTED": 2}, false},
		"REDACTED":     {"REDACTED", map[string]int64{}, false},
		"REDACTED":      {"REDACTED", map[string]int64(nil), false},
		"REDACTED":         {"REDACTED", map[int]int{}, true},
		"REDACTED":                {"REDACTED", time.Date(2020, 6, 3, 17, 35, 30, 0, time.UTC), false},
		"REDACTED":        {"REDACTED", time.Time{}, true},
		"REDACTED":         {"REDACTED", time.Time{}, true},
		"REDACTED":                 {"REDACTED", Car{Revolutions: 4}, false},
		"REDACTED":             {"REDACTED", &Car{Revolutions: 4}, false},
		"REDACTED":           {"REDACTED", Automobile(&Car{Revolutions: 4}), false},
		"REDACTED":                {"REDACTED", Vessel{Navigate: true}, false},
		"REDACTED":            {"REDACTED", &Vessel{Navigate: true}, false},
		"REDACTED":          {"REDACTED", Automobile(Vessel{Navigate: true}), false},
		"REDACTED":       {"REDACTED", Car{}, true},
		"REDACTED": {"REDACTED", Automobile(&Car{}), true},
		"REDACTED":               {"REDACTED", Car{}, true},
		"REDACTED":           {"REDACTED", &Car{}, true},
		"REDACTED":         {"REDACTED", Automobile(&Car{}), true},
		"REDACTED":          {"REDACTED", PublicKey{1, 2, 3, 4, 5, 6, 7, 8}, false},
		"REDACTED":           {"REDACTED", InternalKey{1, 2, 3, 4, 5, 6, 7, 8}, true},
		"REDACTED":        {"REDACTED", Automobile(&Car{}), true},
		"REDACTED": {
			"REDACTED",
			Markers{JSONLabel: "REDACTED", IgnoreEmpty: "REDACTED", Markers: &Markers{JSONLabel: "REDACTED"}},
			false,
		},
		"REDACTED": {
			"REDACTED",
			&Markers{JSONLabel: "REDACTED", IgnoreEmpty: "REDACTED"},
			false,
		},
		"REDACTED": {"REDACTED", Markers{}, false},
		"REDACTED": {
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
REDACTED,
REDACTED"
REDACTED`,
			Struct{
				Bool: true, Float64: 3.14, Int32: 32, Int64: 64, Int64pointer: &i64,
				String: "REDACTED", StringPointerPointer: &strPointer, Octets: []byte{1, 2, 3},
				Time: time.Date(2020, 6, 2, 16, 5, 13, 4346374, time.UTC),
				Car:  &Car{Revolutions: 4}, Vessel: Vessel{Navigate: true}, Automobiles: []Automobile{
					Automobile(&Car{Revolutions: 4}),
					Automobile(Vessel{Navigate: true}),
				},
				Offspring: &Struct{Bool: false, String: "REDACTED"},
			},
			false,
		},
		"REDACTED": {`REDACTED[
REDACTED,
REDACTED}
REDACTED`, Struct{}, true},
		"REDACTED":  {"REDACTED", structNull, false},
		"REDACTED":     {"REDACTED", BespokeItem{}, false},
		"REDACTED":       {"REDACTED", &BespokePointer{Item: "REDACTED"}, false},
		"REDACTED": {"REDACTED", BespokePointer{Item: "REDACTED"}, false},
		"REDACTED":     {"REDACTED", Struct{}, true},
	}
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			//
			//
			objective := reflect.New(reflect.TypeOf(tc.item)).Interface()
			err := json.Unserialize([]byte(tc.json), objective)
			if tc.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			//
			factual := reflect.ValueOf(objective).Elem().Interface()
			assert.Equal(t, tc.item, factual)
		})
	}
}
