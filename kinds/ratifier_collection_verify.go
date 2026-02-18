package kinds

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strings"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/vault/secp256k1"
	cometmath "github.com/valkyrieworks/utils/math"
	engineseed "github.com/valkyrieworks/utils/random"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

func VerifyRatifierCollectionSimple(t *testing.T) {
	//
	//
	rset := NewRatifierCollection([]*Ratifier{})
	assert.Panics(t, func() { rset.AugmentRecommenderUrgency(1) })

	rset = NewRatifierCollection(nil)
	assert.Panics(t, func() { rset.AugmentRecommenderUrgency(1) })

	assert.EqualValues(t, rset, rset.Clone())
	assert.False(t, rset.HasLocation([]byte("REDACTED")))
	idx, val := rset.FetchByLocation([]byte("REDACTED"))
	assert.EqualValues(t, -1, idx)
	assert.Nil(t, val)
	address, val := rset.FetchByOrdinal(-100)
	assert.Nil(t, address)
	assert.Nil(t, val)
	address, val = rset.FetchByOrdinal(0)
	assert.Nil(t, address)
	assert.Nil(t, val)
	address, val = rset.FetchByOrdinal(100)
	assert.Nil(t, address)
	assert.Nil(t, val)
	assert.Zero(t, rset.Volume())
	assert.Equal(t, int64(0), rset.SumPollingEnergy())
	assert.Nil(t, rset.FetchRecommender())
	assert.Equal(t, []byte{
		0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4,
		0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95,
		0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55,
	}, rset.Digest())
	//
	val = randomRatifier(rset.SumPollingEnergy())
	assert.NoError(t, rset.ModifyWithAlterCollection([]*Ratifier{val}))

	assert.True(t, rset.HasLocation(val.Location))
	idx, _ = rset.FetchByLocation(val.Location)
	assert.EqualValues(t, 0, idx)
	address, _ = rset.FetchByOrdinal(0)
	assert.Equal(t, []byte(val.Location), address)
	assert.Equal(t, 1, rset.Volume())
	assert.Equal(t, val.PollingEnergy, rset.SumPollingEnergy())
	assert.NotNil(t, rset.Digest())
	assert.NotPanics(t, func() { rset.AugmentRecommenderUrgency(1) })
	assert.Equal(t, val.Location, rset.FetchRecommender().Location)

	//
	val = randomRatifier(rset.SumPollingEnergy())
	assert.NoError(t, rset.ModifyWithAlterCollection([]*Ratifier{val}))
	_, val = rset.FetchByLocation(val.Location)
	val.PollingEnergy += 100
	recommenderUrgency := val.RecommenderUrgency

	val.RecommenderUrgency = 0
	assert.NoError(t, rset.ModifyWithAlterCollection([]*Ratifier{val}))
	_, val = rset.FetchByLocation(val.Location)
	assert.Equal(t, recommenderUrgency, val.RecommenderUrgency)
}

func Verifyratifierset_Verifybasic(t *testing.T) {
	val, _ := RandomRatifier(false, 1)
	flawedValue := &Ratifier{}
	value2, _ := RandomRatifier(false, 1)

	verifyScenarios := []struct {
		values RatifierAssign
		err  bool
		msg  string
	}{
		{
			values: RatifierAssign{},
			err:  true,
			msg:  "REDACTED",
		},
		{
			values: RatifierAssign{
				Ratifiers: []*Ratifier{},
			},
			err: true,
			msg: "REDACTED",
		},
		{
			values: RatifierAssign{
				Ratifiers: []*Ratifier{val},
			},
			err: true,
			msg: "REDACTED",
		},
		{
			values: RatifierAssign{
				Ratifiers: []*Ratifier{flawedValue},
			},
			err: true,
			msg: "REDACTED",
		},
		{
			values: RatifierAssign{
				Ratifiers: []*Ratifier{val},
				Recommender:   val,
			},
			err: false,
			msg: "REDACTED",
		},
		{
			values: RatifierAssign{
				Ratifiers: []*Ratifier{val},
				Recommender:   value2,
			},
			err: true,
			msg: ErrRecommenderNoInValues.Error(),
		},
	}

	for _, tc := range verifyScenarios {
		err := tc.values.CertifySimple()
		if tc.err {
			if assert.Error(t, err) {
				assert.Equal(t, tc.msg, err.Error())
			}
		} else {
			assert.NoError(t, err)
		}
	}
}

func VerifyClone(t *testing.T) {
	rset := randomRatifierCollection(10)
	rsetDigest := rset.Digest()
	if len(rsetDigest) == 0 {
		t.Fatalf("REDACTED")
	}

	rsetClone := rset.Clone()
	rsetCloneDigest := rsetClone.Digest()

	if !bytes.Equal(rsetDigest, rsetCloneDigest) {
		t.Fatalf("REDACTED", rsetDigest, rsetCloneDigest)
	}
}

func Verifyratifierset_Recommenderurgencyhash(t *testing.T) {
	rset := NewRatifierCollection(nil)
	assert.Equal(t, []byte(nil), rset.RecommenderUrgencyDigest())

	rset = randomRatifierCollection(3)
	assert.NotNil(t, rset.RecommenderUrgencyDigest())

	//
	bz, err := rset.ToSchema()
	assert.NoError(t, err)
	rsetSchema, err := RatifierCollectionFromSchema(bz)
	assert.NoError(t, err)
	assert.Equal(t, rset.RecommenderUrgencyDigest(), rsetSchema.RecommenderUrgencyDigest())

	//
	rsetClone := rset.Clone()
	assert.Equal(t, rset.RecommenderUrgencyDigest(), rsetClone.RecommenderUrgencyDigest())

	//
	rset.AugmentRecommenderUrgency(1)
	assert.Equal(t, rset.Digest(), rsetClone.Digest())
	assert.NotEqual(t, rset.RecommenderUrgencyDigest(), rsetClone.RecommenderUrgencyDigest())

	//
	rsetCopy2 := rset.Clone()
	rsetCopy2.Ratifiers[1].RecommenderUrgency = -rset.Ratifiers[1].RecommenderUrgency * 10
	assert.NotEqual(t, rsetCopy2.Ratifiers[1].RecommenderUrgency, rset.Ratifiers[1].RecommenderUrgency)
	assert.NotEqual(t, rset.RecommenderUrgencyDigest(), rsetCopy2.RecommenderUrgencyDigest())
}

//
func VerifyAugmentRecommenderUrgencyAffirmativeInstances(t *testing.T) {
	rset := NewRatifierCollection([]*Ratifier{
		newRatifier([]byte("REDACTED"), 1000),
		newRatifier([]byte("REDACTED"), 300),
		newRatifier([]byte("REDACTED"), 330),
	})

	assert.Panics(t, func() { rset.AugmentRecommenderUrgency(-1) })
	assert.Panics(t, func() { rset.AugmentRecommenderUrgency(0) })
	rset.AugmentRecommenderUrgency(1)
}

func CriterionRatifierCollectionClone(b *testing.B) {

	rset := NewRatifierCollection([]*Ratifier{})
	for i := 0; i < 1000; i++ {
		privateKey := ed25519.GeneratePrivateKey()
		publicKey := privateKey.PublicKey()
		val := NewRatifier(publicKey, 10)
		err := rset.ModifyWithAlterCollection([]*Ratifier{val})
		if err != nil {
			panic("REDACTED")
		}
	}

	for b.Loop() {
		rset.Clone()
	}
}

//

func VerifyRecommenderChoice1(t *testing.T) {
	rset := NewRatifierCollection([]*Ratifier{
		newRatifier([]byte("REDACTED"), 1000),
		newRatifier([]byte("REDACTED"), 300),
		newRatifier([]byte("REDACTED"), 330),
	})
	var recommenders []string
	for i := 0; i < 99; i++ {
		val := rset.FetchRecommender()
		recommenders = append(recommenders, string(val.Location))
		rset.AugmentRecommenderUrgency(1)
	}
	anticipated := "REDACTED" +
		"REDACTED" +
		"REDACTED" +
		"REDACTED" +
		"REDACTED"
	if anticipated != strings.Join(recommenders, "REDACTED") {
		t.Errorf("REDACTED", anticipated, strings.Join(recommenders, "REDACTED"))
	}
}

func VerifyRecommenderChoice2(t *testing.T) {
	location0 := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	location1 := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	location2 := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}

	//
	node0, value1, value2 := newRatifier(location0, 100), newRatifier(location1, 100), newRatifier(location2, 100)
	valueCatalog := []*Ratifier{node0, value1, value2}
	values := NewRatifierCollection(valueCatalog)
	for i := 0; i < len(valueCatalog)*5; i++ {
		ii := (i) % len(valueCatalog)
		nomination := values.FetchRecommender()
		if !bytes.Equal(nomination.Location, valueCatalog[ii].Location) {
			t.Fatalf("REDACTED", i, valueCatalog[ii].Location, nomination.Location)
		}
		values.AugmentRecommenderUrgency(1)
	}

	//
	*value2 = *newRatifier(location2, 400)
	values = NewRatifierCollection(valueCatalog)
	//
	nomination := values.FetchRecommender()
	if !bytes.Equal(nomination.Location, location2) {
		t.Fatalf("REDACTED", nomination.Location)
	}
	values.AugmentRecommenderUrgency(1)
	nomination = values.FetchRecommender()
	if !bytes.Equal(nomination.Location, location0) {
		t.Fatalf("REDACTED", nomination.Location)
	}

	//
	*value2 = *newRatifier(location2, 401)
	values = NewRatifierCollection(valueCatalog)
	nomination = values.FetchRecommender()
	if !bytes.Equal(nomination.Location, location2) {
		t.Fatalf("REDACTED", nomination.Location)
	}
	values.AugmentRecommenderUrgency(1)
	nomination = values.FetchRecommender()
	if !bytes.Equal(nomination.Location, location2) {
		t.Fatalf("REDACTED", nomination.Location)
	}
	values.AugmentRecommenderUrgency(1)
	nomination = values.FetchRecommender()
	if !bytes.Equal(nomination.Location, location0) {
		t.Fatalf("REDACTED", nomination.Location)
	}

	//
	node0, value1, value2 = newRatifier(location0, 4), newRatifier(location1, 5), newRatifier(location2, 3)
	valueCatalog = []*Ratifier{node0, value1, value2}
	nominateNumber := make([]int, 3)
	values = NewRatifierCollection(valueCatalog)
	N := 1
	for i := 0; i < 120*N; i++ {
		nomination := values.FetchRecommender()
		ii := nomination.Location[19]
		nominateNumber[ii]++
		values.AugmentRecommenderUrgency(1)
	}

	if nominateNumber[0] != 40*N {
		t.Fatalf(
			"REDACTED",
			40*N,
			120*N,
			nominateNumber[0],
			120*N,
		)
	}
	if nominateNumber[1] != 50*N {
		t.Fatalf(
			"REDACTED",
			50*N,
			120*N,
			nominateNumber[1],
			120*N,
		)
	}
	if nominateNumber[2] != 30*N {
		t.Fatalf(
			"REDACTED",
			30*N,
			120*N,
			nominateNumber[2],
			120*N,
		)
	}
}

func VerifyRecommenderChoice3(t *testing.T) {
	values := []*Ratifier{
		newRatifier([]byte("REDACTED"), 1),
		newRatifier([]byte("REDACTED"), 1),
		newRatifier([]byte("REDACTED"), 1),
		newRatifier([]byte("REDACTED"), 1),
	}

	for i := 0; i < 4; i++ {
		pk := ed25519.GeneratePrivateKey().PublicKey()
		values[i].PublicKey = pk
		values[i].Location = pk.Location()
	}
	sort.Sort(RatifiersByLocation(values))
	rset := NewRatifierCollection(values)
	recommenderSequence := make([]*Ratifier, 4)
	for i := 0; i < 4; i++ {
		recommenderSequence[i] = rset.FetchRecommender()
		rset.AugmentRecommenderUrgency(1)
	}

	//
	//
	//
	var (
		i int
		j int32
	)
	for ; i < 10000; i++ {
		got := rset.FetchRecommender().Location
		anticipated := recommenderSequence[j%4].Location
		if !bytes.Equal(got, anticipated) {
			t.Fatalf("REDACTED", got, anticipated, i, j)
		}

		//
		b := rset.toOctets()
		rset = rset.fromOctets(b)

		derived := rset.FetchRecommender() //
		if i != 0 {
			if !bytes.Equal(got, derived.Location) {
				t.Fatalf(
					"REDACTED",
					got,
					derived.Location,
					i,
					j,
				)
			}
		}

		//
		instances := int32(1)
		mod := (engineseed.Int() % 5) + 1
		if engineseed.Int()%mod > 0 {
			//
			instances = (engineseed.Int31() % 4) + 1
		}
		rset.AugmentRecommenderUrgency(instances)

		j += instances
	}
}

func newRatifier(location []byte, energy int64) *Ratifier {
	return &Ratifier{Location: location, PollingEnergy: energy}
}

func randomPublicKey() vault.PublicKey {
	publicKey := make(ed25519.PublicKey, ed25519.PublicKeyVolume)
	copy(publicKey, engineseed.Octets(32))
	return ed25519.PublicKey(engineseed.Octets(32))
}

func randomRatifier(sumPollingEnergy int64) *Ratifier {
	//
	//
	val := NewRatifier(randomPublicKey(), int64(engineseed.Uint64()%uint64(MaximumSumPollingEnergy-sumPollingEnergy)))
	val.RecommenderUrgency = engineseed.Int64() % (MaximumSumPollingEnergy - sumPollingEnergy)
	return val
}

func randomRatifierCollection(countRatifiers int) *RatifierAssign {
	ratifiers := make([]*Ratifier, countRatifiers)
	sumPollingEnergy := int64(0)
	for i := 0; i < countRatifiers; i++ {
		ratifiers[i] = randomRatifier(sumPollingEnergy)
		sumPollingEnergy += ratifiers[i].PollingEnergy
	}
	return NewRatifierCollection(ratifiers)
}

func (values *RatifierAssign) toOctets() []byte {
	schemas, err := values.ToSchema()
	if err != nil {
		panic(err)
	}

	bz, err := schemas.Serialize()
	if err != nil {
		panic(err)
	}

	return bz
}

func (values *RatifierAssign) fromOctets(b []byte) *RatifierAssign {
	schemas := new(engineproto.RatifierAssign)
	err := schemas.Unserialize(b)
	if err != nil {
		//
		panic(err)
	}

	vs, err := RatifierCollectionFromSchema(schemas)
	if err != nil {
		panic(err)
	}

	return vs
}

//

func VerifyRatifierCollectionSumPollingEnergyAlarmsOnOverload(t *testing.T) {
	//
	//
	mustAlarm := func() {
		NewRatifierCollection([]*Ratifier{
			{Location: []byte("REDACTED"), PollingEnergy: math.MaxInt64, RecommenderUrgency: 0},
			{Location: []byte("REDACTED"), PollingEnergy: math.MaxInt64, RecommenderUrgency: 0},
			{Location: []byte("REDACTED"), PollingEnergy: math.MaxInt64, RecommenderUrgency: 0},
		})
	}

	assert.Panics(t, mustAlarm)
}

func VerifyRatifierCollectionFromSchemaYieldsFaultOnOverload(t *testing.T) {
	//
	publicKey := ed25519.GeneratePrivateKey().PublicKey()
	publicidSchema, err := cryptocode.PublicKeyToSchema(publicKey)
	require.NoError(t, err)

	schemaValues := &engineproto.RatifierAssign{
		Ratifiers: []*engineproto.Ratifier{
			{Location: publicKey.Location(), PublicKey: publicidSchema, PollingEnergy: math.MaxInt64, RecommenderUrgency: 0},
			{Location: publicKey.Location(), PublicKey: publicidSchema, PollingEnergy: math.MaxInt64, RecommenderUrgency: 0},
		},
		Recommender: &engineproto.Ratifier{Location: publicKey.Location(), PublicKey: publicidSchema, PollingEnergy: math.MaxInt64, RecommenderUrgency: 0},
	}

	_, err = RatifierCollectionFromSchema(schemaValues)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "REDACTED")
}

func VerifyAverageRecommenderUrgency(t *testing.T) {
	//
	tcs := []struct {
		vs   RatifierAssign
		desire int64
	}{
		0: {RatifierAssign{Ratifiers: []*Ratifier{{RecommenderUrgency: 0}, {RecommenderUrgency: 0}, {RecommenderUrgency: 0}}}, 0},
		1: {
			RatifierAssign{
				Ratifiers: []*Ratifier{{RecommenderUrgency: math.MaxInt64}, {RecommenderUrgency: 0}, {RecommenderUrgency: 0}},
			}, math.MaxInt64 / 3,
		},
		2: {
			RatifierAssign{
				Ratifiers: []*Ratifier{{RecommenderUrgency: math.MaxInt64}, {RecommenderUrgency: 0}},
			}, math.MaxInt64 / 2,
		},
		3: {
			RatifierAssign{
				Ratifiers: []*Ratifier{{RecommenderUrgency: math.MaxInt64}, {RecommenderUrgency: math.MaxInt64}},
			}, math.MaxInt64,
		},
		4: {
			RatifierAssign{
				Ratifiers: []*Ratifier{{RecommenderUrgency: math.MinInt64}, {RecommenderUrgency: math.MinInt64}},
			}, math.MinInt64,
		},
	}
	for i, tc := range tcs {
		got := tc.vs.calculateAverageRecommenderUrgency()
		assert.Equal(t, tc.desire, got, "REDACTED", i)
	}
}

func VerifyCalculatingInAugmentRecommenderUrgency(t *testing.T) {
	//
	//
	//
	tcs := []struct {
		vs    RatifierAssign
		instances int32
		avg   int64
	}{
		0: {
			RatifierAssign{
				Ratifiers: []*Ratifier{
					{Location: []byte("REDACTED"), RecommenderUrgency: 1},
					{Location: []byte("REDACTED"), RecommenderUrgency: 2},
					{Location: []byte("REDACTED"), RecommenderUrgency: 3},
				},
			},
			1, 2,
		},
		1: {
			RatifierAssign{
				Ratifiers: []*Ratifier{
					{Location: []byte("REDACTED"), RecommenderUrgency: 10},
					{Location: []byte("REDACTED"), RecommenderUrgency: -10},
					{Location: []byte("REDACTED"), RecommenderUrgency: 1},
				},
			},
			//
			//
			11,
			0, //
		},
		2: {
			RatifierAssign{
				Ratifiers: []*Ratifier{
					{Location: []byte("REDACTED"), RecommenderUrgency: 100},
					{Location: []byte("REDACTED"), RecommenderUrgency: -10},
					{Location: []byte("REDACTED"), RecommenderUrgency: 1},
				},
			},
			1, 91 / 3,
		},
	}
	for i, tc := range tcs {
		//
		newRset := tc.vs.CloneAugmentRecommenderUrgency(tc.instances)
		for _, val := range tc.vs.Ratifiers {
			_, refreshedValue := newRset.FetchByLocation(val.Location)
			assert.Equal(t, refreshedValue.RecommenderUrgency, val.RecommenderUrgency-tc.avg, "REDACTED", i)
		}
	}
}

func VerifyCalculatingInAugmentRecommenderUrgencyWithPollingEnergy(t *testing.T) {
	//
	//
	//
	vp0 := int64(10)
	vp1 := int64(1)
	vp2 := int64(1)
	sum := vp0 + vp1 + vp2
	avg := (vp0 + vp1 + vp2 - sum) / 3
	values := RatifierAssign{Ratifiers: []*Ratifier{
		{Location: []byte{0}, RecommenderUrgency: 0, PollingEnergy: vp0},
		{Location: []byte{1}, RecommenderUrgency: 0, PollingEnergy: vp1},
		{Location: []byte{2}, RecommenderUrgency: 0, PollingEnergy: vp2},
	}}
	tcs := []struct {
		values                  *RatifierAssign
		desireRecommenderUrgencies []int64
		instances                 int32
		desireRecommender          *Ratifier
	}{
		0: {
			values.Clone(),
			[]int64{
				//
				0 + vp0 - sum - avg, //
				0 + vp1,
				0 + vp2,
			},
			1,
			values.Ratifiers[0],
		},
		1: {
			values.Clone(),
			[]int64{
				(0 + vp0 - sum) + vp0 - sum - avg, //
				(0 + vp1) + vp1,
				(0 + vp2) + vp2,
			},
			2,
			values.Ratifiers[0],
		}, //
		2: {
			values.Clone(),
			[]int64{
				0 + 3*(vp0-sum) - avg, //
				0 + 3*vp1,
				0 + 3*vp2,
			},
			3,
			values.Ratifiers[0],
		},
		3: {
			values.Clone(),
			[]int64{
				0 + 4*(vp0-sum), //
				0 + 4*vp1,
				0 + 4*vp2,
			},
			4,
			values.Ratifiers[0],
		},
		4: {
			values.Clone(),
			[]int64{
				0 + 4*(vp0-sum) + vp0, //
				0 + 5*vp1 - sum,       //
				0 + 5*vp2,
			},
			5,
			values.Ratifiers[1],
		},
		5: {
			values.Clone(),
			[]int64{
				0 + 6*vp0 - 5*sum, //
				0 + 6*vp1 - sum,   //
				0 + 6*vp2,
			},
			6,
			values.Ratifiers[0],
		},
		6: {
			values.Clone(),
			[]int64{
				0 + 7*vp0 - 6*sum, //
				0 + 7*vp1 - sum,   //
				0 + 7*vp2,
			},
			7,
			values.Ratifiers[0],
		},
		7: {
			values.Clone(),
			[]int64{
				0 + 8*vp0 - 7*sum, //
				0 + 8*vp1 - sum,
				0 + 8*vp2,
			},
			8,
			values.Ratifiers[0],
		},
		8: {
			values.Clone(),
			[]int64{
				0 + 9*vp0 - 7*sum,
				0 + 9*vp1 - sum,
				0 + 9*vp2 - sum,
			}, //
			9,
			values.Ratifiers[2],
		},
		9: {
			values.Clone(),
			[]int64{
				0 + 10*vp0 - 8*sum, //
				0 + 10*vp1 - sum,   //
				0 + 10*vp2 - sum,
			}, //
			10,
			values.Ratifiers[0],
		},
		10: {
			values.Clone(),
			[]int64{
				0 + 11*vp0 - 9*sum,
				0 + 11*vp1 - sum, //
				0 + 11*vp2 - sum,
			}, //
			11,
			values.Ratifiers[0],
		},
	}
	for i, tc := range tcs {
		tc.values.AugmentRecommenderUrgency(tc.instances)

		assert.Equal(t, tc.desireRecommender.Location, tc.values.FetchRecommender().Location,
			"REDACTED",
			i)

		for valueIdx, val := range tc.values.Ratifiers {
			assert.Equal(t,
				tc.desireRecommenderUrgencies[valueIdx],
				val.RecommenderUrgency,
				"REDACTED",
				i,
				valueIdx)
		}
	}
}

func VerifySecureAppend(t *testing.T) {
	f := func(a, b int64) bool {
		c, overload := secureAppend(a, b)
		return overload || (!overload && c == a+b)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func VerifySecureAppendTruncate(t *testing.T) {
	assert.EqualValues(t, math.MaxInt64, secureAppendTruncate(math.MaxInt64, 10))
	assert.EqualValues(t, math.MaxInt64, secureAppendTruncate(math.MaxInt64, math.MaxInt64))
	assert.EqualValues(t, math.MinInt64, secureAppendTruncate(math.MinInt64, -10))
}

func VerifySecureSubtractTruncate(t *testing.T) {
	assert.EqualValues(t, math.MinInt64, secureSubtractTruncate(math.MinInt64, 10))
	assert.EqualValues(t, 0, secureSubtractTruncate(math.MinInt64, math.MinInt64))
	assert.EqualValues(t, math.MinInt64, secureSubtractTruncate(math.MinInt64, math.MaxInt64))
	assert.EqualValues(t, math.MaxInt64, secureSubtractTruncate(math.MaxInt64, -10))
}

//

func VerifyEmptyCollection(t *testing.T) {
	var valueCatalog []*Ratifier
	valueCollection := NewRatifierCollection(valueCatalog)
	assert.Panics(t, func() { valueCollection.AugmentRecommenderUrgency(1) })
	assert.Panics(t, func() { valueCollection.ReadjustUrgencies(100) })
	assert.Panics(t, func() { valueCollection.displaceByAverageRecommenderUrgency() })
	assert.Panics(t, func() { assert.Zero(t, calculateMaximumMinimumUrgencyVary(valueCollection)) })
	valueCollection.FetchRecommender()

	//
	v1 := newRatifier([]byte("REDACTED"), 100)
	v2 := newRatifier([]byte("REDACTED"), 100)
	valueCatalog = []*Ratifier{v1, v2}
	assert.NoError(t, valueCollection.ModifyWithAlterCollection(valueCatalog))
	validateRatifierCollection(t, valueCollection)

	//
	v1 = newRatifier([]byte("REDACTED"), 0)
	v2 = newRatifier([]byte("REDACTED"), 0)
	removeCatalog := []*Ratifier{v1, v2}
	assert.Error(t, valueCollection.ModifyWithAlterCollection(removeCatalog))

	//
	assert.Error(t, valueCollection.ModifyWithAlterCollection(removeCatalog))
}

func VerifyRefreshesForNewRatifierCollection(t *testing.T) {
	v1 := newRatifier([]byte("REDACTED"), 100)
	v2 := newRatifier([]byte("REDACTED"), 100)
	valueCatalog := []*Ratifier{v1, v2}
	valueCollection := NewRatifierCollection(valueCatalog)
	validateRatifierCollection(t, valueCollection)

	//
	v111 := newRatifier([]byte("REDACTED"), 100)
	v112 := newRatifier([]byte("REDACTED"), 123)
	v113 := newRatifier([]byte("REDACTED"), 234)
	valueCatalog = []*Ratifier{v111, v112, v113}
	assert.Panics(t, func() { NewRatifierCollection(valueCatalog) })

	//
	v1 = newRatifier([]byte("REDACTED"), 0)
	v2 = newRatifier([]byte("REDACTED"), 22)
	v3 := newRatifier([]byte("REDACTED"), 33)
	valueCatalog = []*Ratifier{v1, v2, v3}
	assert.Panics(t, func() { NewRatifierCollection(valueCatalog) })

	//
	v1 = newRatifier([]byte("REDACTED"), 10)
	v2 = newRatifier([]byte("REDACTED"), -20)
	v3 = newRatifier([]byte("REDACTED"), 30)
	valueCatalog = []*Ratifier{v1, v2, v3}
	assert.Panics(t, func() { NewRatifierCollection(valueCatalog) })
}

type verifyValue struct {
	label  string
	energy int64
}

func arrangement(valueCatalog []verifyValue) []verifyValue {
	if len(valueCatalog) == 0 {
		return nil
	}
	modeCatalog := make([]verifyValue, len(valueCatalog))
	mode := engineseed.Mode(len(valueCatalog))
	for i, v := range mode {
		modeCatalog[v] = valueCatalog[i]
	}
	return modeCatalog
}

func instantiateNewRatifierCatalog(verifyValueCatalog []verifyValue) []*Ratifier {
	valueCatalog := make([]*Ratifier, 0, len(verifyValueCatalog))
	for _, val := range verifyValueCatalog {
		valueCatalog = append(valueCatalog, newRatifier([]byte(val.label), val.energy))
	}
	return valueCatalog
}

func instantiateNewRatifierCollection(verifyValueCatalog []verifyValue) *RatifierAssign {
	return NewRatifierCollection(instantiateNewRatifierCatalog(verifyValueCatalog))
}

func valueCollectionSumRecommenderUrgency(valueCollection *RatifierAssign) int64 {
	sum := int64(0)
	for _, val := range valueCollection.Ratifiers {
		//
		sum = secureAppendTruncate(sum, val.RecommenderUrgency)
	}
	return sum
}

func validateRatifierCollection(t *testing.T, valueCollection *RatifierAssign) {
	//
	assert.Equal(t, len(valueCollection.Ratifiers), cap(valueCollection.Ratifiers))

	//
	tvp := valueCollection.sumPollingEnergy
	err := valueCollection.modifySumPollingEnergy()
	require.NoError(t, err)
	anticipatedTvp := valueCollection.SumPollingEnergy()
	assert.Equal(t, anticipatedTvp, tvp,
		"REDACTED", anticipatedTvp, tvp, valueCollection)

	//
	valuesTally := int64(len(valueCollection.Ratifiers))
	tpp := valueCollectionSumRecommenderUrgency(valueCollection)
	assert.True(t, tpp < valuesTally && tpp > -valuesTally,
		"REDACTED", valuesTally, valuesTally, tpp)

	//
	distance := calculateMaximumMinimumUrgencyVary(valueCollection)
	assert.True(t, distance <= UrgencySpanVolumeCoefficient*tvp,
		"REDACTED", UrgencySpanVolumeCoefficient*tvp, distance)
}

func toVerifyValueCatalog(valueCatalog []*Ratifier) []verifyValue {
	verifyCatalog := make([]verifyValue, len(valueCatalog))
	for i, val := range valueCatalog {
		verifyCatalog[i].label = string(val.Location)
		verifyCatalog[i].energy = val.PollingEnergy
	}
	return verifyCatalog
}

func verifyValueCollection(nValues int, energy int64) []verifyValue {
	values := make([]verifyValue, nValues)
	for i := 0; i < nValues; i++ {
		values[i] = verifyValue{fmt.Sprintf("REDACTED", i+1), energy}
	}
	return values
}

type valueCollectionErrVerifyScenario struct {
	beginValues  []verifyValue
	modifyValues []verifyValue
}

func performValueCollectionErrVerifyScenario(t *testing.T, idx int, tt valueCollectionErrVerifyScenario) {
	//
	valueCollection := instantiateNewRatifierCollection(tt.beginValues)
	valueCollectionClone := valueCollection.Clone()
	valueCatalog := instantiateNewRatifierCatalog(tt.modifyValues)
	valueCatalogClone := ratifierCatalogClone(valueCatalog)
	err := valueCollection.ModifyWithAlterCollection(valueCatalog)

	//
	assert.Error(t, err, "REDACTED", idx)
	assert.Equal(t, valueCollection, valueCollectionClone, "REDACTED", idx)

	//
	assert.Equal(t, valueCatalog, valueCatalogClone, "REDACTED", idx)
}

func VerifyValueCollectionRefreshesReplicatedRecords(t *testing.T) {
	verifyScenarios := []valueCollectionErrVerifyScenario{
		//
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 22}},
		},
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 22}},
		},
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 12}},
		},
		{ //
			verifyValueCollection(3, 10),
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 12}},
		},

		//
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", 0}, {"REDACTED", 0}},
		},
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", 0}, {"REDACTED", 0}},
		},
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", 0}, {"REDACTED", 0}, {"REDACTED", 0}},
		},
		{ //
			verifyValueCollection(3, 10),
			[]verifyValue{{"REDACTED", 0}, {"REDACTED", 0}, {"REDACTED", 0}},
		},

		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", 0}, {"REDACTED", 20}, {"REDACTED", 30}},
		},
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", 0}, {"REDACTED", 20}, {"REDACTED", 30}, {"REDACTED", 0}},
		},
		{ //
			verifyValueCollection(3, 10),
			[]verifyValue{{"REDACTED", 0}, {"REDACTED", 5}, {"REDACTED", 20}, {"REDACTED", 30}, {"REDACTED", 0}},
		},
	}

	for i, tt := range verifyScenarios {
		performValueCollectionErrVerifyScenario(t, i, tt)
	}
}

func VerifyValueCollectionRefreshesOverloads(t *testing.T) {
	maximumVP := MaximumSumPollingEnergy
	verifyScenarios := []valueCollectionErrVerifyScenario{
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", math.MaxInt64}},
		},
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", math.MaxInt64}},
		},
		{ //
			verifyValueCollection(1, maximumVP),
			[]verifyValue{{"REDACTED", math.MaxInt64}},
		},
		{ //
			verifyValueCollection(1, maximumVP-1),
			[]verifyValue{{"REDACTED", 5}},
		},
		{ //
			verifyValueCollection(2, maximumVP/3),
			[]verifyValue{{"REDACTED", maximumVP / 2}},
		},
		{ //
			verifyValueCollection(1, maximumVP),
			[]verifyValue{{"REDACTED", maximumVP}},
		},
	}

	for i, tt := range verifyScenarios {
		performValueCollectionErrVerifyScenario(t, i, tt)
	}
}

func VerifyValueCollectionRefreshesAnotherFaults(t *testing.T) {
	verifyScenarios := []valueCollectionErrVerifyScenario{
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", -123}},
		},
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", -123}},
		},
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", 0}},
		},
		{ //
			[]verifyValue{{"REDACTED", 10}, {"REDACTED", 20}, {"REDACTED", 30}},
			[]verifyValue{{"REDACTED", 0}, {"REDACTED", 0}, {"REDACTED", 0}},
		},
	}

	for i, tt := range verifyScenarios {
		performValueCollectionErrVerifyScenario(t, i, tt)
	}
}

func VerifyValueCollectionRefreshesSimpleVerifiesPerform(t *testing.T) {
	valueCollectionRefreshesSimpleVerifies := []struct {
		beginValues    []verifyValue
		modifyValues   []verifyValue
		anticipatedValues []verifyValue
	}{
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{},
			verifyValueCollection(2, 10),
		},
		{ //
			verifyValueCollection(2, 10),
			[]verifyValue{{"REDACTED", 22}, {"REDACTED", 11}},
			[]verifyValue{{"REDACTED", 22}, {"REDACTED", 11}},
		},
		{ //
			[]verifyValue{{"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyValue{{"REDACTED", 40}, {"REDACTED", 30}},
			[]verifyValue{{"REDACTED", 40}, {"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
		},
		{ //
			[]verifyValue{{"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyValue{{"REDACTED", 30}},
			[]verifyValue{{"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
		},
		{ //
			[]verifyValue{{"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyValue{{"REDACTED", 30}},
			[]verifyValue{{"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
		},
		{ //
			[]verifyValue{{"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyValue{{"REDACTED", 0}},
			[]verifyValue{{"REDACTED", 30}, {"REDACTED", 10}},
		},
	}

	for i, tt := range valueCollectionRefreshesSimpleVerifies {
		//
		valueCollection := instantiateNewRatifierCollection(tt.beginValues)
		valueCatalog := instantiateNewRatifierCatalog(tt.modifyValues)
		err := valueCollection.ModifyWithAlterCollection(valueCatalog)
		assert.NoError(t, err, "REDACTED", i)

		valueCatalogClone := ratifierCatalogClone(valueCollection.Ratifiers)
		//
		//
		//
		if len(valueCatalog) > 0 {
			valueCatalog[0].PollingEnergy++
			assert.Equal(t, toVerifyValueCatalog(valueCatalogClone), toVerifyValueCatalog(valueCollection.Ratifiers), "REDACTED", i)

		}

		//
		assert.Equal(t, tt.anticipatedValues, toVerifyValueCatalog(valueCollection.Ratifiers), "REDACTED", i)
		validateRatifierCollection(t, valueCollection)
	}
}

//
func VerifyValueCollectionRefreshesSequenceAutonomyVerifiesPerform(t *testing.T) {
	//
	//
	//
	valueCollectionRefreshesSequenceVerifies := []struct {
		beginValues  []verifyValue
		modifyValues []verifyValue
	}{
		0: { //
			[]verifyValue{{"REDACTED", 40}, {"REDACTED", 30}, {"REDACTED", 10}, {"REDACTED", 10}},
			[]verifyValue{{"REDACTED", 44}, {"REDACTED", 33}, {"REDACTED", 22}, {"REDACTED", 11}},
		},

		1: { //
			[]verifyValue{{"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyValue{{"REDACTED", 30}, {"REDACTED", 40}, {"REDACTED", 50}, {"REDACTED", 60}},
		},

		2: { //
			[]verifyValue{{"REDACTED", 40}, {"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyValue{{"REDACTED", 0}, {"REDACTED", 0}, {"REDACTED", 0}},
		},

		3: { //
			[]verifyValue{{"REDACTED", 40}, {"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyValue{{"REDACTED", 0}, {"REDACTED", 0}, {"REDACTED", 22}, {"REDACTED", 50}, {"REDACTED", 44}},
		},
	}

	for i, tt := range valueCollectionRefreshesSequenceVerifies {
		//
		valueCollection := instantiateNewRatifierCollection(tt.beginValues)
		valueCollectionClone := valueCollection.Clone()
		valueCatalog := instantiateNewRatifierCatalog(tt.modifyValues)
		assert.NoError(t, valueCollectionClone.ModifyWithAlterCollection(valueCatalog))

		//
		valueCollectionExpiration := valueCollectionClone.Clone()

		//
		n := len(tt.modifyValues)
		maximumCountModes := cometmath.MinimumInteger(20, n*n)
		for j := 0; j < maximumCountModes; j++ {
			//
			valueCollectionClone := valueCollection.Clone()
			valueCatalog := instantiateNewRatifierCatalog(arrangement(tt.modifyValues))

			//
			assert.NoError(t, valueCollectionClone.ModifyWithAlterCollection(valueCatalog),
				"REDACTED", i, valueCatalog)
			validateRatifierCollection(t, valueCollectionClone)

			//
			assert.Equal(t, valueCollectionClone, valueCollectionExpiration,
				"REDACTED", i, valueCatalog)
		}
	}
}

//
//
func VerifyValueCollectionExecuteRefreshesVerifiesPerform(t *testing.T) {
	valueCollectionRefreshesSimpleVerifies := []struct {
		beginValues    []verifyValue
		modifyValues   []verifyValue
		anticipatedValues []verifyValue
	}{
		//
		0: { //
			[]verifyValue{{"REDACTED", 44}, {"REDACTED", 55}},
			[]verifyValue{{"REDACTED", 11}},
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 44}, {"REDACTED", 55}},
		},
		1: { //
			[]verifyValue{{"REDACTED", 44}, {"REDACTED", 55}},
			[]verifyValue{{"REDACTED", 66}},
			[]verifyValue{{"REDACTED", 44}, {"REDACTED", 55}, {"REDACTED", 66}},
		},
		2: { //
			[]verifyValue{{"REDACTED", 44}, {"REDACTED", 66}},
			[]verifyValue{{"REDACTED", 55}},
			[]verifyValue{{"REDACTED", 44}, {"REDACTED", 55}, {"REDACTED", 66}},
		},
		3: { //
			[]verifyValue{{"REDACTED", 44}, {"REDACTED", 66}, {"REDACTED", 99}},
			[]verifyValue{{"REDACTED", 55}, {"REDACTED", 77}, {"REDACTED", 88}},
			[]verifyValue{{"REDACTED", 44}, {"REDACTED", 55}, {"REDACTED", 66}, {"REDACTED", 77}, {"REDACTED", 88}, {"REDACTED", 99}},
		},
		//
		4: { //
			[]verifyValue{{"REDACTED", 111}, {"REDACTED", 22}},
			[]verifyValue{{"REDACTED", 11}},
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 22}},
		},
		5: { //
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 222}},
			[]verifyValue{{"REDACTED", 22}},
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 22}},
		},
		6: { //
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 222}, {"REDACTED", 33}},
			[]verifyValue{{"REDACTED", 22}},
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 33}},
		},
		7: { //
			[]verifyValue{{"REDACTED", 111}, {"REDACTED", 222}, {"REDACTED", 333}},
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 33}},
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 33}},
		},
		//
		8: {
			[]verifyValue{{"REDACTED", 111}, {"REDACTED", 22}},
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 33}, {"REDACTED", 44}},
			[]verifyValue{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 33}, {"REDACTED", 44}},
		},
	}

	for i, tt := range valueCollectionRefreshesSimpleVerifies {
		//
		valueCollection := instantiateNewRatifierCollection(tt.beginValues)

		//
		valueCatalog := instantiateNewRatifierCatalog(tt.modifyValues)
		valueCollection.executeRefreshes(valueCatalog)

		//
		assert.Equal(t, toVerifyValueCatalog(valueCollection.Ratifiers), tt.anticipatedValues, "REDACTED", i)
	}
}

type verifyVCollectionConfig struct {
	label         string
	beginValues    []verifyValue
	removedValues  []verifyValue
	refreshedValues  []verifyValue
	appendedValues    []verifyValue
	anticipatedValues []verifyValue
	expirationErr       error
}

func randomVerifyVCollectionConfig(nRoot, nAppendMaximum int) verifyVCollectionConfig {
	if nRoot <= 0 || nAppendMaximum < 0 {
		panic(fmt.Sprintf("REDACTED", nRoot, nAppendMaximum))
	}

	const maximumEnergy = 1000
	var nAged, nRemove, nModified, nAppend int

	nAged = int(engineseed.Uint()%uint(nRoot)) + 1
	if nRoot-nAged > 0 {
		nRemove = int(engineseed.Uint() % uint(nRoot-nAged))
	}
	nModified = nRoot - nAged - nRemove

	if nAppendMaximum > 0 {
		nAppend = engineseed.Int()%nAppendMaximum + 1
	}

	cfg := verifyVCollectionConfig{}

	cfg.beginValues = make([]verifyValue, nRoot)
	cfg.removedValues = make([]verifyValue, nRemove)
	cfg.appendedValues = make([]verifyValue, nAppend)
	cfg.refreshedValues = make([]verifyValue, nModified)
	cfg.anticipatedValues = make([]verifyValue, nRoot-nRemove+nAppend)

	for i := 0; i < nRoot; i++ {
		cfg.beginValues[i] = verifyValue{fmt.Sprintf("REDACTED", i), int64(engineseed.Uint()%maximumEnergy + 1)}
		if i < nAged {
			cfg.anticipatedValues[i] = cfg.beginValues[i]
		}
		if i >= nAged && i < nAged+nModified {
			cfg.refreshedValues[i-nAged] = verifyValue{fmt.Sprintf("REDACTED", i), int64(engineseed.Uint()%maximumEnergy + 1)}
			cfg.anticipatedValues[i] = cfg.refreshedValues[i-nAged]
		}
		if i >= nAged+nModified {
			cfg.removedValues[i-nAged-nModified] = verifyValue{fmt.Sprintf("REDACTED", i), 0}
		}
	}

	for i := nRoot; i < nRoot+nAppend; i++ {
		cfg.appendedValues[i-nRoot] = verifyValue{fmt.Sprintf("REDACTED", i), int64(engineseed.Uint()%maximumEnergy + 1)}
		cfg.anticipatedValues[i-nRemove] = cfg.appendedValues[i-nRoot]
	}

	sort.Sort(verifyValuesByPollingEnergy(cfg.beginValues))
	sort.Sort(verifyValuesByPollingEnergy(cfg.removedValues))
	sort.Sort(verifyValuesByPollingEnergy(cfg.refreshedValues))
	sort.Sort(verifyValuesByPollingEnergy(cfg.appendedValues))
	sort.Sort(verifyValuesByPollingEnergy(cfg.anticipatedValues))

	return cfg
}

func executeModificationsToValueCollection(t *testing.T, expirationErr error, valueCollection *RatifierAssign, valuesCatalogs ...[]verifyValue) {
	modifications := make([]verifyValue, 0)
	for _, valuesCatalog := range valuesCatalogs {
		modifications = append(modifications, valuesCatalog...)
	}
	valueCatalog := instantiateNewRatifierCatalog(modifications)
	err := valueCollection.ModifyWithAlterCollection(valueCatalog)
	if expirationErr != nil {
		assert.Equal(t, expirationErr, err)
	} else {
		assert.NoError(t, err)
	}
}

func VerifyValueCollectionModifyUrgencySequenceVerifies(t *testing.T) {
	const nMaximumPolls int32 = 5000

	verifyScenarios := []verifyVCollectionConfig{
		0: { //
			beginValues:    []verifyValue{{"REDACTED", 1000}, {"REDACTED", 1}, {"REDACTED", 1}},
			removedValues:  []verifyValue{{"REDACTED", 0}},
			refreshedValues:  []verifyValue{},
			appendedValues:    []verifyValue{},
			anticipatedValues: []verifyValue{{"REDACTED", 1}, {"REDACTED", 1}},
		},
		1: { //
			beginValues:    []verifyValue{{"REDACTED", 1000}, {"REDACTED", 10}, {"REDACTED", 1}},
			removedValues:  []verifyValue{{"REDACTED", 0}},
			refreshedValues:  []verifyValue{},
			appendedValues:    []verifyValue{},
			anticipatedValues: []verifyValue{{"REDACTED", 10}, {"REDACTED", 1}},
		},
		2: { //
			beginValues:    []verifyValue{{"REDACTED", 1000}, {"REDACTED", 2}, {"REDACTED", 1}},
			removedValues:  []verifyValue{{"REDACTED", 0}},
			refreshedValues:  []verifyValue{{"REDACTED", 1}},
			appendedValues:    []verifyValue{{"REDACTED", 50}, {"REDACTED", 40}},
			anticipatedValues: []verifyValue{{"REDACTED", 50}, {"REDACTED", 40}, {"REDACTED", 1}, {"REDACTED", 1}},
		},

		//
		//
		//
		3: randomVerifyVCollectionConfig(100, 10),

		4: randomVerifyVCollectionConfig(1000, 100),

		5: randomVerifyVCollectionConfig(10, 100),

		6: randomVerifyVCollectionConfig(100, 1000),

		7: randomVerifyVCollectionConfig(1000, 1000),
	}

	for _, cfg := range verifyScenarios {

		//
		valueCollection := instantiateNewRatifierCollection(cfg.beginValues)
		validateRatifierCollection(t, valueCollection)

		//
		validateValueCollectionModifyUrgencySequence(t, valueCollection, cfg, nMaximumPolls)
	}
}

func validateValueCollectionModifyUrgencySequence(t *testing.T, valueCollection *RatifierAssign, cfg verifyVCollectionConfig, nMaximumPolls int32) {
	//
	valueCollection.AugmentRecommenderUrgency(engineseed.Int31()%nMaximumPolls + 1)

	//
	executeModificationsToValueCollection(t, nil, valueCollection, cfg.appendedValues, cfg.refreshedValues, cfg.removedValues)

	//
	assert.Equal(t, cfg.anticipatedValues, toVerifyValueCatalog(valueCollection.Ratifiers))
	validateRatifierCollection(t, valueCollection)

	//
	//
	//
	if len(cfg.appendedValues) > 0 {
		refreshedValuesUrgencyOrdered := ratifierCatalogClone(valueCollection.Ratifiers)
		sort.Sort(ratifiersByUrgency(refreshedValuesUrgencyOrdered))

		appendedValuesUrgencySection := refreshedValuesUrgencyOrdered[:len(cfg.appendedValues)]
		sort.Sort(RatifiersByPollingEnergy(appendedValuesUrgencySection))
		assert.Equal(t, cfg.appendedValues, toVerifyValueCatalog(appendedValuesUrgencySection))

		//
		anticipatedUrgency := appendedValuesUrgencySection[0].RecommenderUrgency
		for _, val := range appendedValuesUrgencySection[1:] {
			assert.Equal(t, anticipatedUrgency, val.RecommenderUrgency)
		}
	}
}

func VerifyNewRatifierCollectionFromPresentRatifiers(t *testing.T) {
	volume := 5
	values := make([]*Ratifier, volume)
	for i := 0; i < volume; i++ {
		pv := NewEmulatePV()
		values[i] = pv.RetrieveTowardRatifier(int64(i + 1))
	}
	valueCollection := NewRatifierCollection(values)
	valueCollection.AugmentRecommenderUrgency(5)

	newValueCollection := NewRatifierCollection(valueCollection.Ratifiers)
	assert.NotEqual(t, valueCollection, newValueCollection)

	presentValueCollection, err := RatifierCollectionFromPresentRatifiers(valueCollection.Ratifiers)
	assert.NoError(t, err)
	assert.Equal(t, valueCollection, presentValueCollection)
	assert.Equal(t, valueCollection.CloneAugmentRecommenderUrgency(3), presentValueCollection.CloneAugmentRecommenderUrgency(3))
}

func VerifyValueCollectionModifyOverloadAssociated(t *testing.T) {
	verifyScenarios := []verifyVCollectionConfig{
		{
			label:         "REDACTED",
			beginValues:    []verifyValue{{"REDACTED", MaximumSumPollingEnergy - 1}, {"REDACTED", 1}},
			refreshedValues:  []verifyValue{{"REDACTED", MaximumSumPollingEnergy - 1}, {"REDACTED", 1}},
			anticipatedValues: []verifyValue{{"REDACTED", MaximumSumPollingEnergy - 1}, {"REDACTED", 1}},
			expirationErr:       nil,
		},
		{
			//
			//
			label:         "REDACTED",
			beginValues:    []verifyValue{{"REDACTED", MaximumSumPollingEnergy - 1}, {"REDACTED", 1}},
			refreshedValues:  []verifyValue{{"REDACTED", MaximumSumPollingEnergy/2 - 1}, {"REDACTED", MaximumSumPollingEnergy / 2}},
			anticipatedValues: []verifyValue{{"REDACTED", MaximumSumPollingEnergy / 2}, {"REDACTED", MaximumSumPollingEnergy/2 - 1}},
			expirationErr:       nil,
		},
		{
			label:         "REDACTED",
			beginValues:    []verifyValue{{"REDACTED", MaximumSumPollingEnergy - 2}, {"REDACTED", 1}, {"REDACTED", 1}},
			removedValues:  []verifyValue{{"REDACTED", 0}},
			appendedValues:    []verifyValue{{"REDACTED", MaximumSumPollingEnergy - 2}},
			anticipatedValues: []verifyValue{{"REDACTED", MaximumSumPollingEnergy - 2}, {"REDACTED", 1}, {"REDACTED", 1}},
			expirationErr:       nil,
		},
		{
			label: "REDACTED",
			beginValues: []verifyValue{
				{"REDACTED", MaximumSumPollingEnergy / 4},
				{"REDACTED", MaximumSumPollingEnergy / 4},
				{"REDACTED", MaximumSumPollingEnergy / 4},
				{"REDACTED", MaximumSumPollingEnergy / 4},
			},
			removedValues: []verifyValue{{"REDACTED", 0}},
			refreshedValues: []verifyValue{
				{"REDACTED", MaximumSumPollingEnergy/2 - 2}, {"REDACTED", MaximumSumPollingEnergy/2 - 3}, {"REDACTED", 2},
			},
			appendedValues: []verifyValue{{"REDACTED", 3}},
			anticipatedValues: []verifyValue{
				{"REDACTED", MaximumSumPollingEnergy/2 - 2}, {"REDACTED", MaximumSumPollingEnergy/2 - 3}, {"REDACTED", 3}, {"REDACTED", 2},
			},
			expirationErr: nil,
		},
		{
			label: "REDACTED",
			beginValues: []verifyValue{
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
			},
			refreshedValues: []verifyValue{
				{"REDACTED", MaximumSumPollingEnergy},
				{"REDACTED", MaximumSumPollingEnergy},
				{"REDACTED", MaximumSumPollingEnergy},
				{"REDACTED", MaximumSumPollingEnergy},
				{"REDACTED", MaximumSumPollingEnergy},
				{"REDACTED", MaximumSumPollingEnergy},
				{"REDACTED", MaximumSumPollingEnergy},
				{"REDACTED", MaximumSumPollingEnergy},
				{"REDACTED", 8},
			},
			anticipatedValues: []verifyValue{
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
				{"REDACTED", 1},
			},
			expirationErr: ErrSumPollingEnergyOverload,
		},
	}

	for _, tt := range verifyScenarios {
		t.Run(tt.label, func(t *testing.T) {
			valueCollection := instantiateNewRatifierCollection(tt.beginValues)
			validateRatifierCollection(t, valueCollection)

			//
			executeModificationsToValueCollection(t, tt.expirationErr, valueCollection, tt.appendedValues, tt.refreshedValues, tt.removedValues)

			//
			assert.Equal(t, tt.anticipatedValues, toVerifyValueCatalog(valueCollection.Ratifiers))
			validateRatifierCollection(t, valueCollection)
		})
	}
}

func VerifySecureMultiply(t *testing.T) {
	verifyScenarios := []struct {
		a        int64
		b        int64
		c        int64
		overload bool
	}{
		0: {0, 0, 0, false},
		1: {1, 0, 0, false},
		2: {2, 3, 6, false},
		3: {2, -3, -6, false},
		4: {-2, -3, 6, false},
		5: {-2, 3, -6, false},
		6: {math.MaxInt64, 1, math.MaxInt64, false},
		7: {math.MaxInt64 / 2, 2, math.MaxInt64 - 1, false},
		8: {math.MaxInt64 / 2, 3, 0, true},
		9: {math.MaxInt64, 2, 0, true},
	}

	for i, tc := range verifyScenarios {
		c, overload := secureMultiply(tc.a, tc.b)
		assert.Equal(t, tc.c, c, "REDACTED", i)
		assert.Equal(t, tc.overload, overload, "REDACTED", i)
	}
}

func VerifyRatifierCollectionSchemaBuffer(t *testing.T) {
	ratifierset, _ := RandomRatifierCollection(10, 100)
	ratifierset2, _ := RandomRatifierCollection(10, 100)
	ratifierset2.Ratifiers[0] = &Ratifier{}

	ratifierset3, _ := RandomRatifierCollection(10, 100)
	ratifierset3.Recommender = nil

	ratifierset4, _ := RandomRatifierCollection(10, 100)
	ratifierset4.Recommender = &Ratifier{}

	verifyScenarios := []struct {
		msg      string
		v1       *RatifierAssign
		expirationPass1 bool
		expirationPass2 bool
	}{
		{"REDACTED", ratifierset, true, true},
		{"REDACTED", ratifierset2, false, false},
		{"REDACTED", ratifierset3, false, false},
		{"REDACTED", ratifierset4, false, false},
		{"REDACTED", &RatifierAssign{}, true, false},
		{"REDACTED", nil, true, false},
	}
	for _, tc := range verifyScenarios {
		schemaValueCollection, err := tc.v1.ToSchema()
		if tc.expirationPass1 {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}

		valueCollection, err := RatifierCollectionFromSchema(schemaValueCollection)
		if tc.expirationPass2 {
			require.NoError(t, err, tc.msg)
			require.EqualValues(t, tc.v1, valueCollection, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

//
//
type ratifiersByUrgency []*Ratifier

func (valz ratifiersByUrgency) Len() int {
	return len(valz)
}

func (valz ratifiersByUrgency) Lower(i, j int) bool {
	if valz[i].RecommenderUrgency < valz[j].RecommenderUrgency {
		return true
	}
	if valz[i].RecommenderUrgency > valz[j].RecommenderUrgency {
		return false
	}
	return bytes.Compare(valz[i].Location, valz[j].Location) < 0
}

func (valz ratifiersByUrgency) Exchange(i, j int) {
	valz[i], valz[j] = valz[j], valz[i]
}

//

type verifyValuesByPollingEnergy []verifyValue

func (trats verifyValuesByPollingEnergy) Len() int {
	return len(trats)
}

func (trats verifyValuesByPollingEnergy) Lower(i, j int) bool {
	if trats[i].energy == trats[j].energy {
		return bytes.Compare([]byte(trats[i].label), []byte(trats[j].label)) == -1
	}
	return trats[i].energy > trats[j].energy
}

func (trats verifyValuesByPollingEnergy) Exchange(i, j int) {
	trats[i], trats[j] = trats[j], trats[i]
}

//
//
func CriterionRefreshes(b *testing.B) {
	const (
		n = 100
		m = 2000
	)
	//
	vs := make([]*Ratifier, n)
	for j := 0; j < n; j++ {
		vs[j] = newRatifier([]byte(fmt.Sprintf("REDACTED", j)), 100)
	}
	valueCollection := NewRatifierCollection(vs)
	l := len(valueCollection.Ratifiers)

	//
	newValueCatalog := make([]*Ratifier, m)
	for j := 0; j < m; j++ {
		newValueCatalog[j] = newRatifier([]byte(fmt.Sprintf("REDACTED", j+l)), 1000)
	}

	for b.Loop() {
		//
		valueCollectionClone := valueCollection.Clone()
		assert.NoError(b, valueCollectionClone.ModifyWithAlterCollection(newValueCatalog))
	}
}

func VerifyValidateEndorseWithCorruptRecommenderKey(t *testing.T) {
	vs := &RatifierAssign{
		Ratifiers: []*Ratifier{{}, {}},
	}
	endorse := &Endorse{
		Level:     100,
		Endorsements: []EndorseSignature{{}, {}},
	}
	var bid LedgerUID
	cid := "REDACTED"
	err := vs.ValidateEndorse(cid, bid, 100, endorse)
	assert.Error(t, err)
}

func VerifyValidateEndorseUniqueWithCorruptEndorsements(t *testing.T) {
	vs := &RatifierAssign{
		Ratifiers: []*Ratifier{{}, {}},
	}
	endorse := &Endorse{
		Level:     100,
		Endorsements: []EndorseSignature{{}, {}},
	}
	cid := "REDACTED"
	pollingEnergyRequired := vs.SumPollingEnergy() * 2 / 3

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUIDMark == LedgerUIDMarkMissing }

	//
	tally := func(c EndorseSignature) bool { return c.LedgerUIDMark == LedgerUIDMarkEndorse }

	err := validateEndorseUnique(cid, vs, endorse, pollingEnergyRequired, bypass, tally, true, true, nil)
	require.Error(t, err)

	repository := NewAutographRepository()
	err = validateEndorseUnique(cid, vs, endorse, pollingEnergyRequired, bypass, tally, true, true, repository)
	require.Error(t, err)
	require.Equal(t, 0, repository.Len())
}

func Verifyratifierset_Allkeyshaveidenticalkind(t *testing.T) {
	verifyScenarios := []struct {
		values     *RatifierAssign
		identicalKind bool
	}{
		{
			values:     NewRatifierCollection([]*Ratifier{}),
			identicalKind: true,
		},
		{
			values:     randomRatifierCollection(1),
			identicalKind: true,
		},
		{
			values:     randomRatifierCollection(2),
			identicalKind: true,
		},
		{
			values:     NewRatifierCollection([]*Ratifier{randomRatifier(100), NewRatifier(secp256k1.GeneratePrivateKey().PublicKey(), 200)}),
			identicalKind: false,
		},
	}

	for i, tc := range verifyScenarios {
		if tc.identicalKind {
			assert.True(t, tc.values.AllKeysPossessIdenticalKind(), "REDACTED", i)
		} else {
			assert.False(t, tc.values.AllKeysPossessIdenticalKind(), "REDACTED", i)
		}
	}
}

func Verifyratifierset_Totalpollingpowersecure(t *testing.T) {
	verifyScenarios := []struct {
		label          string
		ratifiers    []*Ratifier
		anticipatedEnergy int64
		anticipateFault   bool
		faultIncludes string
	}{
		{
			label: "REDACTED",
			ratifiers: []*Ratifier{
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 100),
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 200),
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 300),
			},
			anticipatedEnergy: 600,
			anticipateFault:   false,
		},
		{
			label:          "REDACTED",
			ratifiers:    []*Ratifier{},
			anticipatedEnergy: 0,
			anticipateFault:   false,
		},
		{
			label:          "REDACTED",
			ratifiers:    nil,
			anticipatedEnergy: 0,
			anticipateFault:   false,
		},
		{
			label: "REDACTED",
			ratifiers: []*Ratifier{
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 1000),
			},
			anticipatedEnergy: 1000,
			anticipateFault:   false,
		},
		{
			label: "REDACTED",
			ratifiers: []*Ratifier{
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), MaximumSumPollingEnergy),
			},
			anticipatedEnergy: MaximumSumPollingEnergy,
			anticipateFault:   false,
		},
		{
			label: "REDACTED",
			ratifiers: []*Ratifier{
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), MaximumSumPollingEnergy-100),
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 100),
			},
			anticipatedEnergy: MaximumSumPollingEnergy,
			anticipateFault:   false,
		},
		{
			label: "REDACTED",
			ratifiers: []*Ratifier{
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), MaximumSumPollingEnergy/2+1),
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), MaximumSumPollingEnergy/2+1),
			},
			anticipatedEnergy: 0,
			anticipateFault:   true,
			faultIncludes: "REDACTED",
		},
		{
			label: "REDACTED",
			ratifiers: []*Ratifier{
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), MaximumSumPollingEnergy),
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 1),
			},
			anticipatedEnergy: 0,
			anticipateFault:   true,
			faultIncludes: "REDACTED",
		},
		{
			label: "REDACTED",
			ratifiers: []*Ratifier{
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), math.MaxInt64/2),
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), math.MaxInt64/2),
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 100),
			},
			anticipatedEnergy: 0,
			anticipateFault:   true,
			faultIncludes: "REDACTED",
		},
		{
			label: "REDACTED",
			ratifiers: []*Ratifier{
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 100),
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 0),
				NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 200),
			},
			anticipatedEnergy: 300,
			anticipateFault:   false,
		},
		{
			label: "REDACTED",
			ratifiers: func() []*Ratifier {
				values := make([]*Ratifier, 100)
				for i := 0; i < 100; i++ {
					values[i] = NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 1000)
				}
				return values
			}(),
			anticipatedEnergy: 100000,
			anticipateFault:   false,
		},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.label, func(t *testing.T) {
			//
			valueCollection := &RatifierAssign{
				Ratifiers: tc.ratifiers,
			}

			//
			sumEnergy, err := valueCollection.SumPollingEnergySecure()

			//
			if tc.anticipateFault {
				require.Error(t, err, "REDACTED")
				if tc.faultIncludes != "REDACTED" {
					require.Contains(t, err.Error(), tc.faultIncludes,
						"REDACTED")
				}
				require.Equal(t, tc.anticipatedEnergy, sumEnergy,
					"REDACTED", tc.anticipatedEnergy)
			} else {
				require.NoError(t, err, "REDACTED", err)
				require.Equal(t, tc.anticipatedEnergy, sumEnergy,
					"REDACTED", tc.anticipatedEnergy)
			}
		})
	}
}
