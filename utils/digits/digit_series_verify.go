package digits

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
)

var (
	blank16digits = "REDACTED"
	blank64digits = blank16digits + blank16digits + blank16digits + blank16digits
	complete16digits  = "REDACTED"
	complete64digits  = complete16digits + complete16digits + complete16digits + complete16digits
)

func arbitraryDigitSeries(digits int) *DigitSeries {
	src := commitrand.Octets((digits + 7) / 8)
	originPositionTowardDigit := func(i int) bool {
		return src[i/8]&(1<<uint(i%8)) > 0
	}
	return FreshDigitSeriesOriginatingProc(digits, originPositionTowardDigit)
}

func VerifyAlso(t *testing.T) {
	bA1 := arbitraryDigitSeries(51)
	bA2 := arbitraryDigitSeries(31)
	bA3 := bA1.And(bA2)

	var byteVoid *DigitSeries
	require.Equal(t, byteVoid.And(bA1), (*DigitSeries)(nil))
	require.Equal(t, bA1.And(nil), (*DigitSeries)(nil))
	require.Equal(t, byteVoid.And(nil), (*DigitSeries)(nil))

	if bA3.Digits != 31 {
		t.Error("REDACTED", bA3.Digits)
	}
	if len(bA3.Components) != len(bA2.Components) {
		t.Error("REDACTED")
	}
	for i := 0; i < bA3.Digits; i++ {
		anticipated := bA1.ObtainOrdinal(i) && bA2.ObtainOrdinal(i)
		if bA3.ObtainOrdinal(i) != anticipated {
			t.Error("REDACTED", i, bA1.ObtainOrdinal(i), bA2.ObtainOrdinal(i), bA3.ObtainOrdinal(i))
		}
	}
}

func VerifyEither(t *testing.T) {
	bA1 := arbitraryDigitSeries(57)
	bA2 := arbitraryDigitSeries(31)
	bA3 := bA1.Or(bA2)

	byteVoid := (*DigitSeries)(nil)
	require.Equal(t, byteVoid.Or(bA1), bA1)
	require.Equal(t, bA1.Or(nil), bA1)
	require.Equal(t, byteVoid.Or(nil), (*DigitSeries)(nil))

	if bA3.Digits != 57 {
		t.Error("REDACTED")
	}
	if len(bA3.Components) != len(bA1.Components) {
		t.Error("REDACTED")
	}
	for i := 0; i < bA3.Digits; i++ {
		anticipated := bA1.ObtainOrdinal(i) || bA2.ObtainOrdinal(i)
		if bA3.ObtainOrdinal(i) != anticipated {
			t.Error("REDACTED", i, bA1.ObtainOrdinal(i), bA2.ObtainOrdinal(i), bA3.ObtainOrdinal(i))
		}
	}
	if bA3.fetchCountSuccessPositions() == 0 {
		t.Error("REDACTED" +
			"REDACTED")
	}
}

func VerifyUnder(t *testing.T) {
	verifyScenarios := []struct {
		initializeBYA        string
		deductingBYA string
		anticipatedBYA    string
	}{
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
		{"REDACTED", "REDACTED", "REDACTED"},
	}
	for _, tc := range verifyScenarios {
		var bA *DigitSeries
		err := json.Unmarshal([]byte(tc.initializeBYA), &bA)
		require.Nil(t, err)

		var o *DigitSeries
		err = json.Unmarshal([]byte(tc.deductingBYA), &o)
		require.Nil(t, err)

		got, _ := json.Marshal(bA.Sub(o))
		require.Equal(
			t,
			tc.anticipatedBYA,
			string(got),
			"REDACTED",
			tc.initializeBYA,
			tc.deductingBYA,
			tc.anticipatedBYA,
		)
	}
}

func VerifySelectUnpredictable(t *testing.T) {
	verifyScenarios := []struct {
		bA string
		ok bool
	}{
		{"REDACTED", false},
		{"REDACTED", true},
		{"REDACTED" + blank16digits + "REDACTED", false},
		{"REDACTED" + blank16digits + "REDACTED", true},
		{"REDACTED" + blank16digits + "REDACTED", true},
		{"REDACTED" + blank16digits + "REDACTED", true},
		{"REDACTED" + blank64digits + "REDACTED", false},
		{"REDACTED" + blank64digits + "REDACTED", true},
		{"REDACTED" + blank64digits + "REDACTED", true},
		{"REDACTED" + blank64digits + "REDACTED", true},
		{"REDACTED" + blank64digits + "REDACTED", true},
	}
	for _, tc := range verifyScenarios {
		var digitList *DigitSeries
		err := json.Unmarshal([]byte(tc.bA), &digitList)
		require.NoError(t, err)
		_, ok := digitList.SelectArbitrary()
		require.Equal(t, tc.ok, ok, "REDACTED", tc.bA)
	}
}

func VerifyFetchCountSuccessPositions(t *testing.T) {
	type scenario struct {
		Influx          string
		AnticipatedOutcome int
	}
	verifyScenarios := []scenario{
		{"REDACTED", 3},
		{"REDACTED", 0},
		{"REDACTED", 6},
		{"REDACTED", 9},
	}
	countAuthenticVerifyScenarios := len(verifyScenarios)
	for i := 0; i < countAuthenticVerifyScenarios; i++ {
		verifyScenarios = append(verifyScenarios, scenario{verifyScenarios[i].Influx + "REDACTED", verifyScenarios[i].AnticipatedOutcome + 1})
		verifyScenarios = append(verifyScenarios, scenario{complete64digits + verifyScenarios[i].Influx, verifyScenarios[i].AnticipatedOutcome + 64})
		verifyScenarios = append(verifyScenarios, scenario{blank64digits + verifyScenarios[i].Influx, verifyScenarios[i].AnticipatedOutcome})
	}

	for _, tc := range verifyScenarios {
		var digitList *DigitSeries
		err := json.Unmarshal([]byte("REDACTED"+tc.Influx+"REDACTED"), &digitList)
		require.NoError(t, err)
		outcome := digitList.fetchCountSuccessPositions()
		require.Equal(t, tc.AnticipatedOutcome, outcome, "REDACTED", tc.Influx, tc.AnticipatedOutcome, outcome)
		outcome = digitList.Not().fetchCountSuccessPositions()
		require.Equal(t, digitList.Digits-outcome, digitList.fetchCountSuccessPositions())
	}
}

func VerifyFetchCountSuccessPositionsUnfitStatuses(t *testing.T) {
	verifyScenarios := []struct {
		alias string
		bA1  *DigitSeries
		exp  int
	}{
		{"REDACTED", &DigitSeries{}, 0},
		{"REDACTED", &DigitSeries{Digits: 0, Components: nil}, 0},
		{"REDACTED", &DigitSeries{Digits: 0, Components: make([]uint64, 0)}, 0},
		{"REDACTED", nil, 0},
		{"REDACTED", FreshDigitCollection(10), 0},
		{"REDACTED", &DigitSeries{Digits: 0, Components: make([]uint64, 5)}, 0},
		{"REDACTED", &DigitSeries{Digits: 200, Components: make([]uint64, 1)}, 0},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.alias, func(t *testing.T) {
			n := tc.bA1.fetchCountSuccessPositions()
			require.Equal(t, n, tc.exp)
		})
	}
}

func VerifyFetchOrdinalSuccessPosition(t *testing.T) {
	type scenario struct {
		Influx          string
		N              int
		AnticipatedOutcome int
	}
	verifyScenarios := []scenario{
		//
		{"REDACTED", 0, 0},
		{"REDACTED", 1, 2},
		{"REDACTED", 2, 4},
		{"REDACTED", 1, -1},         //
		{"REDACTED", 5, 5},          //
		{"REDACTED", 9, -1}, //

		//
		{"REDACTED", 7, -1}, //
		{"REDACTED", 0, -1}, //
		{"REDACTED", 49, 49},  //
		{"REDACTED", 1, -1},                               //
		{"REDACTED", 63, 63},  //
		{"REDACTED", 64, 64},  //
		{"REDACTED", 100, -1}, //

		//
		{"REDACTED", 99, 99}, //

		//
		{"REDACTED", 3, -1}, //
	}

	countAuthenticVerifyScenarios := len(verifyScenarios)
	//
	for i := 0; i < countAuthenticVerifyScenarios; i++ {
		anticipatedOutcome := verifyScenarios[i].AnticipatedOutcome
		if anticipatedOutcome != -1 {
			anticipatedOutcome += 64
		}
		verifyScenarios = append(verifyScenarios, scenario{blank64digits + verifyScenarios[i].Influx, verifyScenarios[i].N, anticipatedOutcome})
	}

	for _, tc := range verifyScenarios {
		var digitList *DigitSeries
		err := json.Unmarshal([]byte("REDACTED"+tc.Influx+"REDACTED"), &digitList)
		require.NoError(t, err)

		//
		outcome := digitList.fetchOrdinalSuccessPosition(tc.N)

		require.Equal(t, tc.AnticipatedOutcome, outcome, "REDACTED", tc.Influx, tc.N, tc.AnticipatedOutcome, outcome)
	}
}

func VerifyOctets(t *testing.T) {
	bA := FreshDigitCollection(4)
	bA.AssignOrdinal(0, true)
	inspect := func(bA *DigitSeries, bz []byte) {
		if !bytes.Equal(bA.Octets(), bz) {
			panic(fmt.Sprintf("REDACTED", bz, bA.Octets()))
		}
	}
	inspect(bA, []byte{0x01})
	bA.AssignOrdinal(3, true)
	inspect(bA, []byte{0x09})

	bA = FreshDigitCollection(9)
	inspect(bA, []byte{0x00, 0x00})
	bA.AssignOrdinal(7, true)
	inspect(bA, []byte{0x80, 0x00})
	bA.AssignOrdinal(8, true)
	inspect(bA, []byte{0x80, 0x01})

	bA = FreshDigitCollection(16)
	inspect(bA, []byte{0x00, 0x00})
	bA.AssignOrdinal(7, true)
	inspect(bA, []byte{0x80, 0x00})
	bA.AssignOrdinal(8, true)
	inspect(bA, []byte{0x80, 0x01})
	bA.AssignOrdinal(9, true)
	inspect(bA, []byte{0x80, 0x03})

	bA = FreshDigitCollection(4)
	bA.Components = nil
	require.False(t, bA.AssignOrdinal(1, true))
}

func VerifyBlankComplete(t *testing.T) {
	ns := []int{47, 123}
	for _, n := range ns {
		bA := FreshDigitCollection(n)
		if !bA.EqualsBlank() {
			t.Fatal("REDACTED")
		}
		for i := 0; i < n; i++ {
			bA.AssignOrdinal(i, true)
		}
		if !bA.EqualsComplete() {
			t.Fatal("REDACTED")
		}
	}
}

func VerifyReviseEverAlarms(_ *testing.T) {
	freshArbitraryDigitSeries := func(n int) *DigitSeries {
		ba := arbitraryDigitSeries(n)
		return ba
	}
	couples := []struct {
		a, b *DigitSeries
	}{
		{nil, nil},
		{freshArbitraryDigitSeries(10), freshArbitraryDigitSeries(12)},
		{freshArbitraryDigitSeries(23), freshArbitraryDigitSeries(23)},
		{freshArbitraryDigitSeries(37), nil},
		{nil, FreshDigitCollection(10)},
	}

	for _, duo := range couples {
		a, b := duo.a, duo.b
		a.Revise(b)
		b.Revise(a)
	}
}

func VerifyFreshDigitSeriesEverFailuresUponMinuses(_ *testing.T) {
	digitCatalog := []int{-127, -128, -1 << 31}
	for _, digits := range digitCatalog {
		_ = FreshDigitCollection(digits)
	}
}

func VerifyJSNSerializeDecode(t *testing.T) {
	bA1 := FreshDigitCollection(0)

	bA2 := FreshDigitCollection(1)

	bA3 := FreshDigitCollection(1)
	bA3.AssignOrdinal(0, true)

	bA4 := FreshDigitCollection(5)
	bA4.AssignOrdinal(0, true)
	bA4.AssignOrdinal(1, true)

	verifyScenarios := []struct {
		bA           *DigitSeries
		serializedBYA string
	}{
		{nil, "REDACTED"},
		{bA1, "REDACTED"},
		{bA2, "REDACTED"},
		{bA3, "REDACTED"},
		{bA4, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.bA.Text(), func(t *testing.T) {
			bz, err := json.Marshal(tc.bA)
			require.NoError(t, err)

			assert.Equal(t, tc.serializedBYA, string(bz))

			var deserializedBYA *DigitSeries
			err = json.Unmarshal(bz, &deserializedBYA)
			require.NoError(t, err)

			if tc.bA == nil {
				require.Nil(t, deserializedBYA)
			} else {
				require.NotNil(t, deserializedBYA)
				assert.EqualValues(t, tc.bA.Digits, deserializedBYA.Digits)
				if assert.EqualValues(t, tc.bA.Text(), deserializedBYA.Text()) {
					assert.EqualValues(t, tc.bA.Components, deserializedBYA.Components)
				}
			}
		})
	}
}

func VerifyDigitSeriesSchemaArea(t *testing.T) {
	verifyScenarios := []struct {
		msg     string
		bA1     *DigitSeries
		expirationPhrase bool
	}{
		{"REDACTED", &DigitSeries{}, true},
		{"REDACTED", FreshDigitCollection(1), true},
		{"REDACTED", FreshDigitCollection(2), true},
		{"REDACTED", FreshDigitCollection(-1), false},
	}
	for _, tc := range verifyScenarios {
		schemaBYA := tc.bA1.TowardSchema()
		ba := new(DigitSeries)
		ba.OriginatingSchema(schemaBYA)
		if tc.expirationPhrase {
			require.Equal(t, tc.bA1, ba, tc.msg)
		} else {
			require.NotEqual(t, tc.bA1, ba, tc.msg)
		}
	}
}

func VerifyDigitSeriesCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		alias    string
		bA1     *DigitSeries
		expirationPhrase bool
	}{
		{"REDACTED", &DigitSeries{}, true},
		{"REDACTED", &DigitSeries{Digits: 0, Components: nil}, true},
		{"REDACTED", &DigitSeries{Digits: 0, Components: make([]uint64, 0)}, true},
		{"REDACTED", nil, true},
		{"REDACTED", FreshDigitCollection(10), true},
		{"REDACTED", &DigitSeries{Digits: 0, Components: make([]uint64, 5)}, false},
		{"REDACTED", &DigitSeries{Digits: 200, Components: make([]uint64, 1)}, false},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.alias, func(t *testing.T) {
			err := tc.bA1.CertifyFundamental()
			require.Equal(t, err == nil, tc.expirationPhrase)
		})
	}
}

//
//
func VerifyDecodeJSNNotCollapseUponNullDigits(t *testing.T) {
	type positionDataset struct {
		DigitSeries *DigitSeries `json:"ba"`
		Ordinal    int       `json:"i"`
	}

	ic := new(positionDataset)
	chunk := []byte("REDACTED")
	err := json.Unmarshal(chunk, ic)
	require.NoError(t, err)
	require.Equal(t, ic.DigitSeries, &DigitSeries{Digits: 0, Components: nil})
}

func AssessmentSelectUnpredictableDigitSeries(b *testing.B) {
	//
	assessmentDigitSeriesTxt := "REDACTED"
	var digitList *DigitSeries
	err := json.Unmarshal([]byte("REDACTED"+assessmentDigitSeriesTxt+"REDACTED"), &digitList)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = digitList.SelectArbitrary()
	}
}
