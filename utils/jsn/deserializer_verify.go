package jsn__test

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
)

func VerifyDecode(t *testing.T) {
	int64null := (*int64)(nil)
	str := "REDACTED"
	txtReference := &str
	schemaVoid := (*Schema)(nil)
	i32 := int32(32)
	i64 := int64(64)

	verifycases := map[string]struct {
		jsn  string
		datum any
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
		"REDACTED":       {"REDACTED", int64null, false},
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
		"REDACTED":                 {"REDACTED", Car{Rotations: 4}, false},
		"REDACTED":             {"REDACTED", &Car{Rotations: 4}, false},
		"REDACTED":           {"REDACTED", Carriage(&Car{Rotations: 4}), false},
		"REDACTED":                {"REDACTED", Vessel{Navigate: true}, false},
		"REDACTED":            {"REDACTED", &Vessel{Navigate: true}, false},
		"REDACTED":          {"REDACTED", Carriage(Vessel{Navigate: true}), false},
		"REDACTED":       {"REDACTED", Car{}, true},
		"REDACTED": {"REDACTED", Carriage(&Car{}), true},
		"REDACTED":               {"REDACTED", Car{}, true},
		"REDACTED":           {"REDACTED", &Car{}, true},
		"REDACTED":         {"REDACTED", Carriage(&Car{}), true},
		"REDACTED":          {"REDACTED", CommonToken{1, 2, 3, 4, 5, 6, 7, 8}, false},
		"REDACTED":           {"REDACTED", SecludedToken{1, 2, 3, 4, 5, 6, 7, 8}, true},
		"REDACTED":        {"REDACTED", Carriage(&Car{}), true},
		"REDACTED": {
			"REDACTED",
			Labels{JSNAlias: "REDACTED", ExcludeBlank: "REDACTED", Labels: &Labels{JSNAlias: "REDACTED"}},
			false,
		},
		"REDACTED": {
			"REDACTED",
			&Labels{JSNAlias: "REDACTED", ExcludeBlank: "REDACTED"},
			false,
		},
		"REDACTED": {"REDACTED", Labels{}, false},
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
			Schema{
				Flag: true, Float64: 3.14, Integer32: 32, Int64n: 64, Int64reference: &i64,
				Text: "REDACTED", TextReferenceReference: &txtReference, Octets: []byte{1, 2, 3},
				Moment: time.Date(2020, 6, 2, 16, 5, 13, 4346374, time.UTC),
				Car:  &Car{Rotations: 4}, Vessel: Vessel{Navigate: true}, Carriages: []Carriage{
					Carriage(&Car{Rotations: 4}),
					Carriage(Vessel{Navigate: true}),
				},
				Offspring: &Schema{Flag: false, Text: "REDACTED"},
			},
			false,
		},
		"REDACTED": {`REDACTED[
REDACTED,
REDACTED}
REDACTED`, Schema{}, true},
		"REDACTED":  {"REDACTED", schemaVoid, false},
		"REDACTED":     {"REDACTED", BespokeDatum{}, false},
		"REDACTED":       {"REDACTED", &BespokeReference{Datum: "REDACTED"}, false},
		"REDACTED": {"REDACTED", BespokeReference{Datum: "REDACTED"}, false},
		"REDACTED":     {"REDACTED", Schema{}, true},
	}
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			//
			//
			objective := reflect.New(reflect.TypeOf(tc.datum)).Interface()
			err := jsn.Decode([]byte(tc.jsn), objective)
			if tc.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			//
			existing := reflect.ValueOf(objective).Elem().Interface()
			assert.Equal(t, tc.datum, existing)
		})
	}
}
