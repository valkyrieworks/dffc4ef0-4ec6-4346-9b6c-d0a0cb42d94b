package bits

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	engineseed "github.com/valkyrieworks/utils/random"
)

var (
	void16bits = "REDACTED"
	void64bits = void16bits + void16bits + void16bits + void16bits
	complete16bits  = "REDACTED"
	complete64bits  = complete16bits + complete16bits + complete16bits + complete16bits
)

func randomBitList(bits int) *BitList {
	src := engineseed.Octets((bits + 7) / 8)
	sourceOrdinalToBit := func(i int) bool {
		return src[i/8]&(1<<uint(i%8)) > 0
	}
	return NewBitListFromFn(bits, sourceOrdinalToBit)
}

func VerifyAnd(t *testing.T) {
	bA1 := randomBitList(51)
	bA2 := randomBitList(31)
	bA3 := bA1.And(bA2)

	var byteNull *BitList
	require.Equal(t, byteNull.And(bA1), (*BitList)(nil))
	require.Equal(t, bA1.And(nil), (*BitList)(nil))
	require.Equal(t, byteNull.And(nil), (*BitList)(nil))

	if bA3.Bits != 31 {
		t.Error("REDACTED", bA3.Bits)
	}
	if len(bA3.Elements) != len(bA2.Elements) {
		t.Error("REDACTED")
	}
	for i := 0; i < bA3.Bits; i++ {
		anticipated := bA1.FetchOrdinal(i) && bA2.FetchOrdinal(i)
		if bA3.FetchOrdinal(i) != anticipated {
			t.Error("REDACTED", i, bA1.FetchOrdinal(i), bA2.FetchOrdinal(i), bA3.FetchOrdinal(i))
		}
	}
}

func VerifyOr(t *testing.T) {
	bA1 := randomBitList(57)
	bA2 := randomBitList(31)
	bA3 := bA1.Or(bA2)

	byteNull := (*BitList)(nil)
	require.Equal(t, byteNull.Or(bA1), bA1)
	require.Equal(t, bA1.Or(nil), bA1)
	require.Equal(t, byteNull.Or(nil), (*BitList)(nil))

	if bA3.Bits != 57 {
		t.Error("REDACTED")
	}
	if len(bA3.Elements) != len(bA1.Elements) {
		t.Error("REDACTED")
	}
	for i := 0; i < bA3.Bits; i++ {
		anticipated := bA1.FetchOrdinal(i) || bA2.FetchOrdinal(i)
		if bA3.FetchOrdinal(i) != anticipated {
			t.Error("REDACTED", i, bA1.FetchOrdinal(i), bA2.FetchOrdinal(i), bA3.FetchOrdinal(i))
		}
	}
	if bA3.fetchCountTrueOrdinals() == 0 {
		t.Error("REDACTED" +
			"REDACTED")
	}
}

func VerifySubtract(t *testing.T) {
	verifyScenarios := []struct {
		initBA        string
		deductingBA string
		anticipatedBA    string
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
		var bA *BitList
		err := json.Unmarshal([]byte(tc.initBA), &bA)
		require.Nil(t, err)

		var o *BitList
		err = json.Unmarshal([]byte(tc.deductingBA), &o)
		require.Nil(t, err)

		got, _ := json.Marshal(bA.Sub(o))
		require.Equal(
			t,
			tc.anticipatedBA,
			string(got),
			"REDACTED",
			tc.initBA,
			tc.deductingBA,
			tc.anticipatedBA,
		)
	}
}

func VerifySelectArbitrary(t *testing.T) {
	verifyScenarios := []struct {
		bA string
		ok bool
	}{
		{"REDACTED", false},
		{"REDACTED", true},
		{"REDACTED" + void16bits + "REDACTED", false},
		{"REDACTED" + void16bits + "REDACTED", true},
		{"REDACTED" + void16bits + "REDACTED", true},
		{"REDACTED" + void16bits + "REDACTED", true},
		{"REDACTED" + void64bits + "REDACTED", false},
		{"REDACTED" + void64bits + "REDACTED", true},
		{"REDACTED" + void64bits + "REDACTED", true},
		{"REDACTED" + void64bits + "REDACTED", true},
		{"REDACTED" + void64bits + "REDACTED", true},
	}
	for _, tc := range verifyScenarios {
		var bitArr *BitList
		err := json.Unmarshal([]byte(tc.bA), &bitArr)
		require.NoError(t, err)
		_, ok := bitArr.SelectArbitrary()
		require.Equal(t, tc.ok, ok, "REDACTED", tc.bA)
	}
}

func VerifyFetchCountTrueOrdinals(t *testing.T) {
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
	countInitialVerifyScenarios := len(verifyScenarios)
	for i := 0; i < countInitialVerifyScenarios; i++ {
		verifyScenarios = append(verifyScenarios, scenario{verifyScenarios[i].Influx + "REDACTED", verifyScenarios[i].AnticipatedOutcome + 1})
		verifyScenarios = append(verifyScenarios, scenario{complete64bits + verifyScenarios[i].Influx, verifyScenarios[i].AnticipatedOutcome + 64})
		verifyScenarios = append(verifyScenarios, scenario{void64bits + verifyScenarios[i].Influx, verifyScenarios[i].AnticipatedOutcome})
	}

	for _, tc := range verifyScenarios {
		var bitArr *BitList
		err := json.Unmarshal([]byte("REDACTED"+tc.Influx+"REDACTED"), &bitArr)
		require.NoError(t, err)
		outcome := bitArr.fetchCountTrueOrdinals()
		require.Equal(t, tc.AnticipatedOutcome, outcome, "REDACTED", tc.Influx, tc.AnticipatedOutcome, outcome)
		outcome = bitArr.Not().fetchCountTrueOrdinals()
		require.Equal(t, bitArr.Bits-outcome, bitArr.fetchCountTrueOrdinals())
	}
}

func VerifyFetchCountTrueOrdinalsCorruptConditions(t *testing.T) {
	verifyScenarios := []struct {
		label string
		bA1  *BitList
		exp  int
	}{
		{"REDACTED", &BitList{}, 0},
		{"REDACTED", &BitList{Bits: 0, Elements: nil}, 0},
		{"REDACTED", &BitList{Bits: 0, Elements: make([]uint64, 0)}, 0},
		{"REDACTED", nil, 0},
		{"REDACTED", NewBitList(10), 0},
		{"REDACTED", &BitList{Bits: 0, Elements: make([]uint64, 5)}, 0},
		{"REDACTED", &BitList{Bits: 200, Elements: make([]uint64, 1)}, 0},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.label, func(t *testing.T) {
			n := tc.bA1.fetchCountTrueOrdinals()
			require.Equal(t, n, tc.exp)
		})
	}
}

func VerifyFetchNthTrueOrdinal(t *testing.T) {
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

	countInitialVerifyScenarios := len(verifyScenarios)
	//
	for i := 0; i < countInitialVerifyScenarios; i++ {
		anticipatedOutcome := verifyScenarios[i].AnticipatedOutcome
		if anticipatedOutcome != -1 {
			anticipatedOutcome += 64
		}
		verifyScenarios = append(verifyScenarios, scenario{void64bits + verifyScenarios[i].Influx, verifyScenarios[i].N, anticipatedOutcome})
	}

	for _, tc := range verifyScenarios {
		var bitArr *BitList
		err := json.Unmarshal([]byte("REDACTED"+tc.Influx+"REDACTED"), &bitArr)
		require.NoError(t, err)

		//
		outcome := bitArr.fetchNthTrueOrdinal(tc.N)

		require.Equal(t, tc.AnticipatedOutcome, outcome, "REDACTED", tc.Influx, tc.N, tc.AnticipatedOutcome, outcome)
	}
}

func VerifyOctets(t *testing.T) {
	bA := NewBitList(4)
	bA.AssignOrdinal(0, true)
	inspect := func(bA *BitList, bz []byte) {
		if !bytes.Equal(bA.Octets(), bz) {
			panic(fmt.Sprintf("REDACTED", bz, bA.Octets()))
		}
	}
	inspect(bA, []byte{0x01})
	bA.AssignOrdinal(3, true)
	inspect(bA, []byte{0x09})

	bA = NewBitList(9)
	inspect(bA, []byte{0x00, 0x00})
	bA.AssignOrdinal(7, true)
	inspect(bA, []byte{0x80, 0x00})
	bA.AssignOrdinal(8, true)
	inspect(bA, []byte{0x80, 0x01})

	bA = NewBitList(16)
	inspect(bA, []byte{0x00, 0x00})
	bA.AssignOrdinal(7, true)
	inspect(bA, []byte{0x80, 0x00})
	bA.AssignOrdinal(8, true)
	inspect(bA, []byte{0x80, 0x01})
	bA.AssignOrdinal(9, true)
	inspect(bA, []byte{0x80, 0x03})

	bA = NewBitList(4)
	bA.Elements = nil
	require.False(t, bA.AssignOrdinal(1, true))
}

func VerifyEmptyComplete(t *testing.T) {
	ns := []int{47, 123}
	for _, n := range ns {
		bA := NewBitList(n)
		if !bA.IsEmpty() {
			t.Fatal("REDACTED")
		}
		for i := 0; i < n; i++ {
			bA.AssignOrdinal(i, true)
		}
		if !bA.IsComplete() {
			t.Fatal("REDACTED")
		}
	}
}

func VerifyModifyNeverAlarms(_ *testing.T) {
	newRandomBitList := func(n int) *BitList {
		ba := randomBitList(n)
		return ba
	}
	couples := []struct {
		a, b *BitList
	}{
		{nil, nil},
		{newRandomBitList(10), newRandomBitList(12)},
		{newRandomBitList(23), newRandomBitList(23)},
		{newRandomBitList(37), nil},
		{nil, NewBitList(10)},
	}

	for _, couple := range couples {
		a, b := couple.a, couple.b
		a.Modify(b)
		b.Modify(a)
	}
}

func VerifyNewBitListNeverCollapsesOnMinuses(_ *testing.T) {
	bitCatalog := []int{-127, -128, -1 << 31}
	for _, bits := range bitCatalog {
		_ = NewBitList(bits)
	}
}

func VerifyJSONSerializeUnserialize(t *testing.T) {
	bA1 := NewBitList(0)

	bA2 := NewBitList(1)

	bA3 := NewBitList(1)
	bA3.AssignOrdinal(0, true)

	bA4 := NewBitList(5)
	bA4.AssignOrdinal(0, true)
	bA4.AssignOrdinal(1, true)

	verifyScenarios := []struct {
		bA           *BitList
		serializedBA string
	}{
		{nil, "REDACTED"},
		{bA1, "REDACTED"},
		{bA2, "REDACTED"},
		{bA3, "REDACTED"},
		{bA4, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.bA.String(), func(t *testing.T) {
			bz, err := json.Marshal(tc.bA)
			require.NoError(t, err)

			assert.Equal(t, tc.serializedBA, string(bz))

			var unserializedBA *BitList
			err = json.Unmarshal(bz, &unserializedBA)
			require.NoError(t, err)

			if tc.bA == nil {
				require.Nil(t, unserializedBA)
			} else {
				require.NotNil(t, unserializedBA)
				assert.EqualValues(t, tc.bA.Bits, unserializedBA.Bits)
				if assert.EqualValues(t, tc.bA.String(), unserializedBA.String()) {
					assert.EqualValues(t, tc.bA.Elements, unserializedBA.Elements)
				}
			}
		})
	}
}

func VerifyBitListSchemaImage(t *testing.T) {
	verifyScenarios := []struct {
		msg     string
		bA1     *BitList
		expirationPass bool
	}{
		{"REDACTED", &BitList{}, true},
		{"REDACTED", NewBitList(1), true},
		{"REDACTED", NewBitList(2), true},
		{"REDACTED", NewBitList(-1), false},
	}
	for _, tc := range verifyScenarios {
		schemaBA := tc.bA1.ToSchema()
		ba := new(BitList)
		ba.FromSchema(schemaBA)
		if tc.expirationPass {
			require.Equal(t, tc.bA1, ba, tc.msg)
		} else {
			require.NotEqual(t, tc.bA1, ba, tc.msg)
		}
	}
}

func VerifyBitListCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		label    string
		bA1     *BitList
		expirationPass bool
	}{
		{"REDACTED", &BitList{}, true},
		{"REDACTED", &BitList{Bits: 0, Elements: nil}, true},
		{"REDACTED", &BitList{Bits: 0, Elements: make([]uint64, 0)}, true},
		{"REDACTED", nil, true},
		{"REDACTED", NewBitList(10), true},
		{"REDACTED", &BitList{Bits: 0, Elements: make([]uint64, 5)}, false},
		{"REDACTED", &BitList{Bits: 200, Elements: make([]uint64, 1)}, false},
	}
	for _, tc := range verifyScenarios {
		t.Run(tc.label, func(t *testing.T) {
			err := tc.bA1.CertifySimple()
			require.Equal(t, err == nil, tc.expirationPass)
		})
	}
}

//
//
func VerifyUnserializeJSONDoesntCollapseOnNilBits(t *testing.T) {
	type ordinalDataset struct {
		BitList *BitList `json:"ba"`
		Ordinal    int       `json:"i"`
	}

	ic := new(ordinalDataset)
	binary := []byte("REDACTED")
	err := json.Unmarshal(binary, ic)
	require.NoError(t, err)
	require.Equal(t, ic.BitList, &BitList{Bits: 0, Elements: nil})
}

func CriterionSelectArbitraryBitList(b *testing.B) {
	//
	criterionBitListStr := "REDACTED"
	var bitArr *BitList
	err := json.Unmarshal([]byte("REDACTED"+criterionBitListStr+"REDACTED"), &bitArr)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = bitArr.SelectArbitrary()
	}
}
