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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/ellipticp256"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

func VerifyAssessorAssignFundamental(t *testing.T) {
	//
	//
	voterset := FreshAssessorAssign([]*Assessor{})
	assert.Panics(t, func() { voterset.AdvanceNominatorUrgency(1) })

	voterset = FreshAssessorAssign(nil)
	assert.Panics(t, func() { voterset.AdvanceNominatorUrgency(1) })

	assert.EqualValues(t, voterset, voterset.Duplicate())
	assert.False(t, voterset.OwnsLocation([]byte("REDACTED")))
	idx, val := voterset.ObtainViaLocation([]byte("REDACTED"))
	assert.EqualValues(t, -1, idx)
	assert.Nil(t, val)
	location, val := voterset.ObtainViaOrdinal(-100)
	assert.Nil(t, location)
	assert.Nil(t, val)
	location, val = voterset.ObtainViaOrdinal(0)
	assert.Nil(t, location)
	assert.Nil(t, val)
	location, val = voterset.ObtainViaOrdinal(100)
	assert.Nil(t, location)
	assert.Nil(t, val)
	assert.Zero(t, voterset.Extent())
	assert.Equal(t, int64(0), voterset.SumBallotingPotency())
	assert.Nil(t, voterset.ObtainNominator())
	assert.Equal(t, []byte{
		0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4,
		0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95,
		0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55,
	}, voterset.Digest())
	//
	val = arbitraryAssessor(voterset.SumBallotingPotency())
	assert.NoError(t, voterset.ReviseUsingModifyAssign([]*Assessor{val}))

	assert.True(t, voterset.OwnsLocation(val.Location))
	idx, _ = voterset.ObtainViaLocation(val.Location)
	assert.EqualValues(t, 0, idx)
	location, _ = voterset.ObtainViaOrdinal(0)
	assert.Equal(t, []byte(val.Location), location)
	assert.Equal(t, 1, voterset.Extent())
	assert.Equal(t, val.BallotingPotency, voterset.SumBallotingPotency())
	assert.NotNil(t, voterset.Digest())
	assert.NotPanics(t, func() { voterset.AdvanceNominatorUrgency(1) })
	assert.Equal(t, val.Location, voterset.ObtainNominator().Location)

	//
	val = arbitraryAssessor(voterset.SumBallotingPotency())
	assert.NoError(t, voterset.ReviseUsingModifyAssign([]*Assessor{val}))
	_, val = voterset.ObtainViaLocation(val.Location)
	val.BallotingPotency += 100
	nominatorUrgency := val.NominatorUrgency

	val.NominatorUrgency = 0
	assert.NoError(t, voterset.ReviseUsingModifyAssign([]*Assessor{val}))
	_, val = voterset.ObtainViaLocation(val.Location)
	assert.Equal(t, nominatorUrgency, val.NominatorUrgency)
}

func Testvalidset_Certifyfundamental(t *testing.T) {
	val, _ := ArbitraryAssessor(false, 1)
	flawedItem := &Assessor{}
	valid2, _ := ArbitraryAssessor(false, 1)

	verifyScenarios := []struct {
		values AssessorAssign
		err  bool
		msg  string
	}{
		{
			values: AssessorAssign{},
			err:  true,
			msg:  "REDACTED",
		},
		{
			values: AssessorAssign{
				Assessors: []*Assessor{},
			},
			err: true,
			msg: "REDACTED",
		},
		{
			values: AssessorAssign{
				Assessors: []*Assessor{val},
			},
			err: true,
			msg: "REDACTED",
		},
		{
			values: AssessorAssign{
				Assessors: []*Assessor{flawedItem},
			},
			err: true,
			msg: "REDACTED",
		},
		{
			values: AssessorAssign{
				Assessors: []*Assessor{val},
				Nominator:   val,
			},
			err: false,
			msg: "REDACTED",
		},
		{
			values: AssessorAssign{
				Assessors: []*Assessor{val},
				Nominator:   valid2,
			},
			err: true,
			msg: FaultNominatorNegationInsideValues.Error(),
		},
	}

	for _, tc := range verifyScenarios {
		err := tc.values.CertifyFundamental()
		if tc.err {
			if assert.Error(t, err) {
				assert.Equal(t, tc.msg, err.Error())
			}
		} else {
			assert.NoError(t, err)
		}
	}
}

func VerifyDuplicate(t *testing.T) {
	voterset := arbitraryAssessorAssign(10)
	votersetDigest := voterset.Digest()
	if len(votersetDigest) == 0 {
		t.Fatalf("REDACTED")
	}

	votersetDuplicate := voterset.Duplicate()
	votersetDuplicateDigest := votersetDuplicate.Digest()

	if !bytes.Equal(votersetDigest, votersetDuplicateDigest) {
		t.Fatalf("REDACTED", votersetDigest, votersetDuplicateDigest)
	}
}

func Testvalidset_Nominatorurgencydigest(t *testing.T) {
	voterset := FreshAssessorAssign(nil)
	assert.Equal(t, []byte(nil), voterset.NominatorUrgencyDigest())

	voterset = arbitraryAssessorAssign(3)
	assert.NotNil(t, voterset.NominatorUrgencyDigest())

	//
	bz, err := voterset.TowardSchema()
	assert.NoError(t, err)
	votersetSchema, err := AssessorAssignOriginatingSchema(bz)
	assert.NoError(t, err)
	assert.Equal(t, voterset.NominatorUrgencyDigest(), votersetSchema.NominatorUrgencyDigest())

	//
	votersetDuplicate := voterset.Duplicate()
	assert.Equal(t, voterset.NominatorUrgencyDigest(), votersetDuplicate.NominatorUrgencyDigest())

	//
	voterset.AdvanceNominatorUrgency(1)
	assert.Equal(t, voterset.Digest(), votersetDuplicate.Digest())
	assert.NotEqual(t, voterset.NominatorUrgencyDigest(), votersetDuplicate.NominatorUrgencyDigest())

	//
	votersetDuplicate2 := voterset.Duplicate()
	votersetDuplicate2.Assessors[1].NominatorUrgency = -voterset.Assessors[1].NominatorUrgency * 10
	assert.NotEqual(t, votersetDuplicate2.Assessors[1].NominatorUrgency, voterset.Assessors[1].NominatorUrgency)
	assert.NotEqual(t, voterset.NominatorUrgencyDigest(), votersetDuplicate2.NominatorUrgencyDigest())
}

//
func VerifyAdvanceNominatorUrgencyAffirmativeMultiples(t *testing.T) {
	voterset := FreshAssessorAssign([]*Assessor{
		freshAssessor([]byte("REDACTED"), 1000),
		freshAssessor([]byte("REDACTED"), 300),
		freshAssessor([]byte("REDACTED"), 330),
	})

	assert.Panics(t, func() { voterset.AdvanceNominatorUrgency(-1) })
	assert.Panics(t, func() { voterset.AdvanceNominatorUrgency(0) })
	voterset.AdvanceNominatorUrgency(1)
}

func AssessmentAssessorAssignDuplicate(b *testing.B) {

	voterset := FreshAssessorAssign([]*Assessor{})
	for i := 0; i < 1000; i++ {
		privateToken := edwards25519.ProducePrivateToken()
		publicToken := privateToken.PublicToken()
		val := FreshAssessor(publicToken, 10)
		err := voterset.ReviseUsingModifyAssign([]*Assessor{val})
		if err != nil {
			panic("REDACTED")
		}
	}

	for b.Loop() {
		voterset.Duplicate()
	}
}

//

func VerifyNominatorOption1(t *testing.T) {
	voterset := FreshAssessorAssign([]*Assessor{
		freshAssessor([]byte("REDACTED"), 1000),
		freshAssessor([]byte("REDACTED"), 300),
		freshAssessor([]byte("REDACTED"), 330),
	})
	var nominators []string
	for i := 0; i < 99; i++ {
		val := voterset.ObtainNominator()
		nominators = append(nominators, string(val.Location))
		voterset.AdvanceNominatorUrgency(1)
	}
	anticipated := "REDACTED" +
		"REDACTED" +
		"REDACTED" +
		"REDACTED" +
		"REDACTED"
	if anticipated != strings.Join(nominators, "REDACTED") {
		t.Errorf("REDACTED", anticipated, strings.Join(nominators, "REDACTED"))
	}
}

func VerifyNominatorOption2(t *testing.T) {
	location0 := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	location1 := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	location2 := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}

	//
	item0, valid1, valid2 := freshAssessor(location0, 100), freshAssessor(location1, 100), freshAssessor(location2, 100)
	itemCatalog := []*Assessor{item0, valid1, valid2}
	values := FreshAssessorAssign(itemCatalog)
	for i := 0; i < len(itemCatalog)*5; i++ {
		ii := (i) % len(itemCatalog)
		item := values.ObtainNominator()
		if !bytes.Equal(item.Location, itemCatalog[ii].Location) {
			t.Fatalf("REDACTED", i, itemCatalog[ii].Location, item.Location)
		}
		values.AdvanceNominatorUrgency(1)
	}

	//
	*valid2 = *freshAssessor(location2, 400)
	values = FreshAssessorAssign(itemCatalog)
	//
	item := values.ObtainNominator()
	if !bytes.Equal(item.Location, location2) {
		t.Fatalf("REDACTED", item.Location)
	}
	values.AdvanceNominatorUrgency(1)
	item = values.ObtainNominator()
	if !bytes.Equal(item.Location, location0) {
		t.Fatalf("REDACTED", item.Location)
	}

	//
	*valid2 = *freshAssessor(location2, 401)
	values = FreshAssessorAssign(itemCatalog)
	item = values.ObtainNominator()
	if !bytes.Equal(item.Location, location2) {
		t.Fatalf("REDACTED", item.Location)
	}
	values.AdvanceNominatorUrgency(1)
	item = values.ObtainNominator()
	if !bytes.Equal(item.Location, location2) {
		t.Fatalf("REDACTED", item.Location)
	}
	values.AdvanceNominatorUrgency(1)
	item = values.ObtainNominator()
	if !bytes.Equal(item.Location, location0) {
		t.Fatalf("REDACTED", item.Location)
	}

	//
	item0, valid1, valid2 = freshAssessor(location0, 4), freshAssessor(location1, 5), freshAssessor(location2, 3)
	itemCatalog = []*Assessor{item0, valid1, valid2}
	nominateNumber := make([]int, 3)
	values = FreshAssessorAssign(itemCatalog)
	N := 1
	for i := 0; i < 120*N; i++ {
		item := values.ObtainNominator()
		ii := item.Location[19]
		nominateNumber[ii]++
		values.AdvanceNominatorUrgency(1)
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

func VerifyNominatorOption3(t *testing.T) {
	values := []*Assessor{
		freshAssessor([]byte("REDACTED"), 1),
		freshAssessor([]byte("REDACTED"), 1),
		freshAssessor([]byte("REDACTED"), 1),
		freshAssessor([]byte("REDACTED"), 1),
	}

	for i := 0; i < 4; i++ {
		pk := edwards25519.ProducePrivateToken().PublicToken()
		values[i].PublicToken = pk
		values[i].Location = pk.Location()
	}
	sort.Sort(AssessorsViaLocator(values))
	voterset := FreshAssessorAssign(values)
	nominatorSequence := make([]*Assessor, 4)
	for i := 0; i < 4; i++ {
		nominatorSequence[i] = voterset.ObtainNominator()
		voterset.AdvanceNominatorUrgency(1)
	}

	//
	//
	//
	var (
		i int
		j int32
	)
	for ; i < 10000; i++ {
		got := voterset.ObtainNominator().Location
		anticipated := nominatorSequence[j%4].Location
		if !bytes.Equal(got, anticipated) {
			t.Fatalf("REDACTED", got, anticipated, i, j)
		}

		//
		b := voterset.towardOctets()
		voterset = voterset.originatingOctets(b)

		estimated := voterset.ObtainNominator() //
		if i != 0 {
			if !bytes.Equal(got, estimated.Location) {
				t.Fatalf(
					"REDACTED",
					got,
					estimated.Location,
					i,
					j,
				)
			}
		}

		//
		multiples := int32(1)
		mod := (commitrand.Int() % 5) + 1
		if commitrand.Int()%mod > 0 {
			//
			multiples = (commitrand.Int31n() % 4) + 1
		}
		voterset.AdvanceNominatorUrgency(multiples)

		j += multiples
	}
}

func freshAssessor(location []byte, potency int64) *Assessor {
	return &Assessor{Location: location, BallotingPotency: potency}
}

func arbitraryPublicToken() security.PublicToken {
	publicToken := make(edwards25519.PublicToken, edwards25519.PublicTokenExtent)
	copy(publicToken, commitrand.Octets(32))
	return edwards25519.PublicToken(commitrand.Octets(32))
}

func arbitraryAssessor(sumBallotingPotency int64) *Assessor {
	//
	//
	val := FreshAssessor(arbitraryPublicToken(), int64(commitrand.Uint64n()%uint64(MaximumSumBallotingPotency-sumBallotingPotency)))
	val.NominatorUrgency = commitrand.Int64n() % (MaximumSumBallotingPotency - sumBallotingPotency)
	return val
}

func arbitraryAssessorAssign(countAssessors int) *AssessorAssign {
	assessors := make([]*Assessor, countAssessors)
	sumBallotingPotency := int64(0)
	for i := 0; i < countAssessors; i++ {
		assessors[i] = arbitraryAssessor(sumBallotingPotency)
		sumBallotingPotency += assessors[i].BallotingPotency
	}
	return FreshAssessorAssign(assessors)
}

func (values *AssessorAssign) towardOctets() []byte {
	buffervs, err := values.TowardSchema()
	if err != nil {
		panic(err)
	}

	bz, err := buffervs.Serialize()
	if err != nil {
		panic(err)
	}

	return bz
}

func (values *AssessorAssign) originatingOctets(b []byte) *AssessorAssign {
	buffervs := new(commitchema.AssessorAssign)
	err := buffervs.Decode(b)
	if err != nil {
		//
		panic(err)
	}

	vs, err := AssessorAssignOriginatingSchema(buffervs)
	if err != nil {
		panic(err)
	}

	return vs
}

//

func VerifyAssessorAssignSumBallotingPotencyAlarmsUponOverrun(t *testing.T) {
	//
	//
	mustAlarm := func() {
		FreshAssessorAssign([]*Assessor{
			{Location: []byte("REDACTED"), BallotingPotency: math.MaxInt64, NominatorUrgency: 0},
			{Location: []byte("REDACTED"), BallotingPotency: math.MaxInt64, NominatorUrgency: 0},
			{Location: []byte("REDACTED"), BallotingPotency: math.MaxInt64, NominatorUrgency: 0},
		})
	}

	assert.Panics(t, mustAlarm)
}

func VerifyAssessorAssignOriginatingSchemaYieldsFailureUponOverrun(t *testing.T) {
	//
	publicToken := edwards25519.ProducePrivateToken().PublicToken()
	keySchema, err := cryptocode.PublicTokenTowardSchema(publicToken)
	require.NoError(t, err)

	schemaValues := &commitchema.AssessorAssign{
		Assessors: []*commitchema.Assessor{
			{Location: publicToken.Location(), PublicToken: keySchema, BallotingPotency: math.MaxInt64, NominatorUrgency: 0},
			{Location: publicToken.Location(), PublicToken: keySchema, BallotingPotency: math.MaxInt64, NominatorUrgency: 0},
		},
		Nominator: &commitchema.Assessor{Location: publicToken.Location(), PublicToken: keySchema, BallotingPotency: math.MaxInt64, NominatorUrgency: 0},
	}

	_, err = AssessorAssignOriginatingSchema(schemaValues)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "REDACTED")
}

func VerifyMedianNominatorUrgency(t *testing.T) {
	//
	tcs := []struct {
		vs   AssessorAssign
		desire int64
	}{
		0: {AssessorAssign{Assessors: []*Assessor{{NominatorUrgency: 0}, {NominatorUrgency: 0}, {NominatorUrgency: 0}}}, 0},
		1: {
			AssessorAssign{
				Assessors: []*Assessor{{NominatorUrgency: math.MaxInt64}, {NominatorUrgency: 0}, {NominatorUrgency: 0}},
			}, math.MaxInt64 / 3,
		},
		2: {
			AssessorAssign{
				Assessors: []*Assessor{{NominatorUrgency: math.MaxInt64}, {NominatorUrgency: 0}},
			}, math.MaxInt64 / 2,
		},
		3: {
			AssessorAssign{
				Assessors: []*Assessor{{NominatorUrgency: math.MaxInt64}, {NominatorUrgency: math.MaxInt64}},
			}, math.MaxInt64,
		},
		4: {
			AssessorAssign{
				Assessors: []*Assessor{{NominatorUrgency: math.MinInt64}, {NominatorUrgency: math.MinInt64}},
			}, math.MinInt64,
		},
	}
	for i, tc := range tcs {
		got := tc.vs.calculateMedianNominatorUrgency()
		assert.Equal(t, tc.desire, got, "REDACTED", i)
	}
}

func VerifyMediatingInsideAdvanceNominatorUrgency(t *testing.T) {
	//
	//
	//
	tcs := []struct {
		vs    AssessorAssign
		multiples int32
		avg   int64
	}{
		0: {
			AssessorAssign{
				Assessors: []*Assessor{
					{Location: []byte("REDACTED"), NominatorUrgency: 1},
					{Location: []byte("REDACTED"), NominatorUrgency: 2},
					{Location: []byte("REDACTED"), NominatorUrgency: 3},
				},
			},
			1, 2,
		},
		1: {
			AssessorAssign{
				Assessors: []*Assessor{
					{Location: []byte("REDACTED"), NominatorUrgency: 10},
					{Location: []byte("REDACTED"), NominatorUrgency: -10},
					{Location: []byte("REDACTED"), NominatorUrgency: 1},
				},
			},
			//
			//
			11,
			0, //
		},
		2: {
			AssessorAssign{
				Assessors: []*Assessor{
					{Location: []byte("REDACTED"), NominatorUrgency: 100},
					{Location: []byte("REDACTED"), NominatorUrgency: -10},
					{Location: []byte("REDACTED"), NominatorUrgency: 1},
				},
			},
			1, 91 / 3,
		},
	}
	for i, tc := range tcs {
		//
		freshVoterset := tc.vs.DuplicateAdvanceNominatorUrgency(tc.multiples)
		for _, val := range tc.vs.Assessors {
			_, revisedItem := freshVoterset.ObtainViaLocation(val.Location)
			assert.Equal(t, revisedItem.NominatorUrgency, val.NominatorUrgency-tc.avg, "REDACTED", i)
		}
	}
}

func VerifyMediatingInsideAdvanceNominatorUrgencyUsingBallotingPotency(t *testing.T) {
	//
	//
	//
	vp0 := int64(10)
	vp1 := int64(1)
	vp2 := int64(1)
	sum := vp0 + vp1 + vp2
	avg := (vp0 + vp1 + vp2 - sum) / 3
	values := AssessorAssign{Assessors: []*Assessor{
		{Location: []byte{0}, NominatorUrgency: 0, BallotingPotency: vp0},
		{Location: []byte{1}, NominatorUrgency: 0, BallotingPotency: vp1},
		{Location: []byte{2}, NominatorUrgency: 0, BallotingPotency: vp2},
	}}
	tcs := []struct {
		values                  *AssessorAssign
		desireNominatorUrgencys []int64
		multiples                 int32
		desireNominator          *Assessor
	}{
		0: {
			values.Duplicate(),
			[]int64{
				//
				0 + vp0 - sum - avg, //
				0 + vp1,
				0 + vp2,
			},
			1,
			values.Assessors[0],
		},
		1: {
			values.Duplicate(),
			[]int64{
				(0 + vp0 - sum) + vp0 - sum - avg, //
				(0 + vp1) + vp1,
				(0 + vp2) + vp2,
			},
			2,
			values.Assessors[0],
		}, //
		2: {
			values.Duplicate(),
			[]int64{
				0 + 3*(vp0-sum) - avg, //
				0 + 3*vp1,
				0 + 3*vp2,
			},
			3,
			values.Assessors[0],
		},
		3: {
			values.Duplicate(),
			[]int64{
				0 + 4*(vp0-sum), //
				0 + 4*vp1,
				0 + 4*vp2,
			},
			4,
			values.Assessors[0],
		},
		4: {
			values.Duplicate(),
			[]int64{
				0 + 4*(vp0-sum) + vp0, //
				0 + 5*vp1 - sum,       //
				0 + 5*vp2,
			},
			5,
			values.Assessors[1],
		},
		5: {
			values.Duplicate(),
			[]int64{
				0 + 6*vp0 - 5*sum, //
				0 + 6*vp1 - sum,   //
				0 + 6*vp2,
			},
			6,
			values.Assessors[0],
		},
		6: {
			values.Duplicate(),
			[]int64{
				0 + 7*vp0 - 6*sum, //
				0 + 7*vp1 - sum,   //
				0 + 7*vp2,
			},
			7,
			values.Assessors[0],
		},
		7: {
			values.Duplicate(),
			[]int64{
				0 + 8*vp0 - 7*sum, //
				0 + 8*vp1 - sum,
				0 + 8*vp2,
			},
			8,
			values.Assessors[0],
		},
		8: {
			values.Duplicate(),
			[]int64{
				0 + 9*vp0 - 7*sum,
				0 + 9*vp1 - sum,
				0 + 9*vp2 - sum,
			}, //
			9,
			values.Assessors[2],
		},
		9: {
			values.Duplicate(),
			[]int64{
				0 + 10*vp0 - 8*sum, //
				0 + 10*vp1 - sum,   //
				0 + 10*vp2 - sum,
			}, //
			10,
			values.Assessors[0],
		},
		10: {
			values.Duplicate(),
			[]int64{
				0 + 11*vp0 - 9*sum,
				0 + 11*vp1 - sum, //
				0 + 11*vp2 - sum,
			}, //
			11,
			values.Assessors[0],
		},
	}
	for i, tc := range tcs {
		tc.values.AdvanceNominatorUrgency(tc.multiples)

		assert.Equal(t, tc.desireNominator.Location, tc.values.ObtainNominator().Location,
			"REDACTED",
			i)

		for itemOffset, val := range tc.values.Assessors {
			assert.Equal(t,
				tc.desireNominatorUrgencys[itemOffset],
				val.NominatorUrgency,
				"REDACTED",
				i,
				itemOffset)
		}
	}
}

func VerifySecureAppend(t *testing.T) {
	f := func(a, b int64) bool {
		c, overrun := secureAppend(a, b)
		return overrun || (!overrun && c == a+b)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func VerifySecureAppendRestrict(t *testing.T) {
	assert.EqualValues(t, math.MaxInt64, secureAppendRestrict(math.MaxInt64, 10))
	assert.EqualValues(t, math.MaxInt64, secureAppendRestrict(math.MaxInt64, math.MaxInt64))
	assert.EqualValues(t, math.MinInt64, secureAppendRestrict(math.MinInt64, -10))
}

func VerifySecureUnderRestrict(t *testing.T) {
	assert.EqualValues(t, math.MinInt64, secureUnderRestrict(math.MinInt64, 10))
	assert.EqualValues(t, 0, secureUnderRestrict(math.MinInt64, math.MinInt64))
	assert.EqualValues(t, math.MinInt64, secureUnderRestrict(math.MinInt64, math.MaxInt64))
	assert.EqualValues(t, math.MaxInt64, secureUnderRestrict(math.MaxInt64, -10))
}

//

func VerifyBlankAssign(t *testing.T) {
	var itemCatalog []*Assessor
	itemAssign := FreshAssessorAssign(itemCatalog)
	assert.Panics(t, func() { itemAssign.AdvanceNominatorUrgency(1) })
	assert.Panics(t, func() { itemAssign.RecalibrateUrgencies(100) })
	assert.Panics(t, func() { itemAssign.relocateViaMedianNominatorUrgency() })
	assert.Panics(t, func() { assert.Zero(t, calculateMaximumMinimumUrgencyVariance(itemAssign)) })
	itemAssign.ObtainNominator()

	//
	v1 := freshAssessor([]byte("REDACTED"), 100)
	v2 := freshAssessor([]byte("REDACTED"), 100)
	itemCatalog = []*Assessor{v1, v2}
	assert.NoError(t, itemAssign.ReviseUsingModifyAssign(itemCatalog))
	validateAssessorAssign(t, itemAssign)

	//
	v1 = freshAssessor([]byte("REDACTED"), 0)
	v2 = freshAssessor([]byte("REDACTED"), 0)
	removeCatalog := []*Assessor{v1, v2}
	assert.Error(t, itemAssign.ReviseUsingModifyAssign(removeCatalog))

	//
	assert.Error(t, itemAssign.ReviseUsingModifyAssign(removeCatalog))
}

func VerifyRevisionsForeachFreshAssessorAssign(t *testing.T) {
	v1 := freshAssessor([]byte("REDACTED"), 100)
	v2 := freshAssessor([]byte("REDACTED"), 100)
	itemCatalog := []*Assessor{v1, v2}
	itemAssign := FreshAssessorAssign(itemCatalog)
	validateAssessorAssign(t, itemAssign)

	//
	ver111 := freshAssessor([]byte("REDACTED"), 100)
	ver112 := freshAssessor([]byte("REDACTED"), 123)
	ver113 := freshAssessor([]byte("REDACTED"), 234)
	itemCatalog = []*Assessor{ver111, ver112, ver113}
	assert.Panics(t, func() { FreshAssessorAssign(itemCatalog) })

	//
	v1 = freshAssessor([]byte("REDACTED"), 0)
	v2 = freshAssessor([]byte("REDACTED"), 22)
	v3 := freshAssessor([]byte("REDACTED"), 33)
	itemCatalog = []*Assessor{v1, v2, v3}
	assert.Panics(t, func() { FreshAssessorAssign(itemCatalog) })

	//
	v1 = freshAssessor([]byte("REDACTED"), 10)
	v2 = freshAssessor([]byte("REDACTED"), -20)
	v3 = freshAssessor([]byte("REDACTED"), 30)
	itemCatalog = []*Assessor{v1, v2, v3}
	assert.Panics(t, func() { FreshAssessorAssign(itemCatalog) })
}

type verifyItem struct {
	alias  string
	potency int64
}

func arrangement(itemCatalog []verifyItem) []verifyItem {
	if len(itemCatalog) == 0 {
		return nil
	}
	modeCatalog := make([]verifyItem, len(itemCatalog))
	mode := commitrand.Mode(len(itemCatalog))
	for i, v := range mode {
		modeCatalog[v] = itemCatalog[i]
	}
	return modeCatalog
}

func generateFreshAssessorCatalog(verifyItemCatalog []verifyItem) []*Assessor {
	itemCatalog := make([]*Assessor, 0, len(verifyItemCatalog))
	for _, val := range verifyItemCatalog {
		itemCatalog = append(itemCatalog, freshAssessor([]byte(val.alias), val.potency))
	}
	return itemCatalog
}

func generateFreshAssessorAssign(verifyItemCatalog []verifyItem) *AssessorAssign {
	return FreshAssessorAssign(generateFreshAssessorCatalog(verifyItemCatalog))
}

func itemAssignSumNominatorUrgency(itemAssign *AssessorAssign) int64 {
	sum := int64(0)
	for _, val := range itemAssign.Assessors {
		//
		sum = secureAppendRestrict(sum, val.NominatorUrgency)
	}
	return sum
}

func validateAssessorAssign(t *testing.T, itemAssign *AssessorAssign) {
	//
	assert.Equal(t, len(itemAssign.Assessors), cap(itemAssign.Assessors))

	//
	tvp := itemAssign.sumBallotingPotency
	err := itemAssign.reviseSumBallotingPotency()
	require.NoError(t, err)
	anticipatedTvoter := itemAssign.SumBallotingPotency()
	assert.Equal(t, anticipatedTvoter, tvp,
		"REDACTED", anticipatedTvoter, tvp, itemAssign)

	//
	valuesTally := int64(len(itemAssign.Assessors))
	tpp := itemAssignSumNominatorUrgency(itemAssign)
	assert.True(t, tpp < valuesTally && tpp > -valuesTally,
		"REDACTED", valuesTally, valuesTally, tpp)

	//
	spread := calculateMaximumMinimumUrgencyVariance(itemAssign)
	assert.True(t, spread <= UrgencyFrameworkExtentElement*tvp,
		"REDACTED", UrgencyFrameworkExtentElement*tvp, spread)
}

func towardVerifyItemCatalog(itemCatalog []*Assessor) []verifyItem {
	verifyCatalog := make([]verifyItem, len(itemCatalog))
	for i, val := range itemCatalog {
		verifyCatalog[i].alias = string(val.Location)
		verifyCatalog[i].potency = val.BallotingPotency
	}
	return verifyCatalog
}

func verifyItemAssign(nthValues int, potency int64) []verifyItem {
	values := make([]verifyItem, nthValues)
	for i := 0; i < nthValues; i++ {
		values[i] = verifyItem{fmt.Sprintf("REDACTED", i+1), potency}
	}
	return values
}

type itemAssignFaultVerifyInstance struct {
	initiateValues  []verifyItem
	reviseValues []verifyItem
}

func performItemAssignFaultVerifyInstance(t *testing.T, idx int, tt itemAssignFaultVerifyInstance) {
	//
	itemAssign := generateFreshAssessorAssign(tt.initiateValues)
	itemAssignDuplicate := itemAssign.Duplicate()
	itemCatalog := generateFreshAssessorCatalog(tt.reviseValues)
	itemCatalogDuplicate := assessorCatalogDuplicate(itemCatalog)
	err := itemAssign.ReviseUsingModifyAssign(itemCatalog)

	//
	assert.Error(t, err, "REDACTED", idx)
	assert.Equal(t, itemAssign, itemAssignDuplicate, "REDACTED", idx)

	//
	assert.Equal(t, itemCatalog, itemCatalogDuplicate, "REDACTED", idx)
}

func VerifyItemAssignRevisionsReplicatedListings(t *testing.T) {
	verifyScenarios := []itemAssignFaultVerifyInstance{
		//
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 22}},
		},
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 22}},
		},
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 12}},
		},
		{ //
			verifyItemAssign(3, 10),
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 12}},
		},

		//
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", 0}, {"REDACTED", 0}},
		},
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", 0}, {"REDACTED", 0}},
		},
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", 0}, {"REDACTED", 0}, {"REDACTED", 0}},
		},
		{ //
			verifyItemAssign(3, 10),
			[]verifyItem{{"REDACTED", 0}, {"REDACTED", 0}, {"REDACTED", 0}},
		},

		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", 0}, {"REDACTED", 20}, {"REDACTED", 30}},
		},
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", 0}, {"REDACTED", 20}, {"REDACTED", 30}, {"REDACTED", 0}},
		},
		{ //
			verifyItemAssign(3, 10),
			[]verifyItem{{"REDACTED", 0}, {"REDACTED", 5}, {"REDACTED", 20}, {"REDACTED", 30}, {"REDACTED", 0}},
		},
	}

	for i, tt := range verifyScenarios {
		performItemAssignFaultVerifyInstance(t, i, tt)
	}
}

func VerifyItemAssignRevisionsSurpluses(t *testing.T) {
	maximumVoter := MaximumSumBallotingPotency
	verifyScenarios := []itemAssignFaultVerifyInstance{
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", math.MaxInt64}},
		},
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", math.MaxInt64}},
		},
		{ //
			verifyItemAssign(1, maximumVoter),
			[]verifyItem{{"REDACTED", math.MaxInt64}},
		},
		{ //
			verifyItemAssign(1, maximumVoter-1),
			[]verifyItem{{"REDACTED", 5}},
		},
		{ //
			verifyItemAssign(2, maximumVoter/3),
			[]verifyItem{{"REDACTED", maximumVoter / 2}},
		},
		{ //
			verifyItemAssign(1, maximumVoter),
			[]verifyItem{{"REDACTED", maximumVoter}},
		},
	}

	for i, tt := range verifyScenarios {
		performItemAssignFaultVerifyInstance(t, i, tt)
	}
}

func VerifyItemAssignRevisionsAnotherFaults(t *testing.T) {
	verifyScenarios := []itemAssignFaultVerifyInstance{
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", -123}},
		},
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", -123}},
		},
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", 0}},
		},
		{ //
			[]verifyItem{{"REDACTED", 10}, {"REDACTED", 20}, {"REDACTED", 30}},
			[]verifyItem{{"REDACTED", 0}, {"REDACTED", 0}, {"REDACTED", 0}},
		},
	}

	for i, tt := range verifyScenarios {
		performItemAssignFaultVerifyInstance(t, i, tt)
	}
}

func VerifyItemAssignRevisionsFundamentalVerifiesPerform(t *testing.T) {
	itemAssignRevisionsFundamentalVerifies := []struct {
		initiateValues    []verifyItem
		reviseValues   []verifyItem
		anticipatedValues []verifyItem
	}{
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{},
			verifyItemAssign(2, 10),
		},
		{ //
			verifyItemAssign(2, 10),
			[]verifyItem{{"REDACTED", 22}, {"REDACTED", 11}},
			[]verifyItem{{"REDACTED", 22}, {"REDACTED", 11}},
		},
		{ //
			[]verifyItem{{"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyItem{{"REDACTED", 40}, {"REDACTED", 30}},
			[]verifyItem{{"REDACTED", 40}, {"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
		},
		{ //
			[]verifyItem{{"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyItem{{"REDACTED", 30}},
			[]verifyItem{{"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
		},
		{ //
			[]verifyItem{{"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyItem{{"REDACTED", 30}},
			[]verifyItem{{"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
		},
		{ //
			[]verifyItem{{"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyItem{{"REDACTED", 0}},
			[]verifyItem{{"REDACTED", 30}, {"REDACTED", 10}},
		},
	}

	for i, tt := range itemAssignRevisionsFundamentalVerifies {
		//
		itemAssign := generateFreshAssessorAssign(tt.initiateValues)
		itemCatalog := generateFreshAssessorCatalog(tt.reviseValues)
		err := itemAssign.ReviseUsingModifyAssign(itemCatalog)
		assert.NoError(t, err, "REDACTED", i)

		itemCatalogDuplicate := assessorCatalogDuplicate(itemAssign.Assessors)
		//
		//
		//
		if len(itemCatalog) > 0 {
			itemCatalog[0].BallotingPotency++
			assert.Equal(t, towardVerifyItemCatalog(itemCatalogDuplicate), towardVerifyItemCatalog(itemAssign.Assessors), "REDACTED", i)

		}

		//
		assert.Equal(t, tt.anticipatedValues, towardVerifyItemCatalog(itemAssign.Assessors), "REDACTED", i)
		validateAssessorAssign(t, itemAssign)
	}
}

//
func VerifyItemAssignRevisionsSequenceAutonomyVerifiesPerform(t *testing.T) {
	//
	//
	//
	itemAssignRevisionsSequenceVerifies := []struct {
		initiateValues  []verifyItem
		reviseValues []verifyItem
	}{
		0: { //
			[]verifyItem{{"REDACTED", 40}, {"REDACTED", 30}, {"REDACTED", 10}, {"REDACTED", 10}},
			[]verifyItem{{"REDACTED", 44}, {"REDACTED", 33}, {"REDACTED", 22}, {"REDACTED", 11}},
		},

		1: { //
			[]verifyItem{{"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyItem{{"REDACTED", 30}, {"REDACTED", 40}, {"REDACTED", 50}, {"REDACTED", 60}},
		},

		2: { //
			[]verifyItem{{"REDACTED", 40}, {"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyItem{{"REDACTED", 0}, {"REDACTED", 0}, {"REDACTED", 0}},
		},

		3: { //
			[]verifyItem{{"REDACTED", 40}, {"REDACTED", 30}, {"REDACTED", 20}, {"REDACTED", 10}},
			[]verifyItem{{"REDACTED", 0}, {"REDACTED", 0}, {"REDACTED", 22}, {"REDACTED", 50}, {"REDACTED", 44}},
		},
	}

	for i, tt := range itemAssignRevisionsSequenceVerifies {
		//
		itemAssign := generateFreshAssessorAssign(tt.initiateValues)
		itemAssignDuplicate := itemAssign.Duplicate()
		itemCatalog := generateFreshAssessorCatalog(tt.reviseValues)
		assert.NoError(t, itemAssignDuplicate.ReviseUsingModifyAssign(itemCatalog))

		//
		itemAssignExpiration := itemAssignDuplicate.Duplicate()

		//
		n := len(tt.reviseValues)
		maximumCountModes := strongarithmetic.MinimumInteger(20, n*n)
		for j := 0; j < maximumCountModes; j++ {
			//
			itemAssignDuplicate := itemAssign.Duplicate()
			itemCatalog := generateFreshAssessorCatalog(arrangement(tt.reviseValues))

			//
			assert.NoError(t, itemAssignDuplicate.ReviseUsingModifyAssign(itemCatalog),
				"REDACTED", i, itemCatalog)
			validateAssessorAssign(t, itemAssignDuplicate)

			//
			assert.Equal(t, itemAssignDuplicate, itemAssignExpiration,
				"REDACTED", i, itemCatalog)
		}
	}
}

//
//
func VerifyItemAssignExecuteRevisionsVerifiesPerform(t *testing.T) {
	itemAssignRevisionsFundamentalVerifies := []struct {
		initiateValues    []verifyItem
		reviseValues   []verifyItem
		anticipatedValues []verifyItem
	}{
		//
		0: { //
			[]verifyItem{{"REDACTED", 44}, {"REDACTED", 55}},
			[]verifyItem{{"REDACTED", 11}},
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 44}, {"REDACTED", 55}},
		},
		1: { //
			[]verifyItem{{"REDACTED", 44}, {"REDACTED", 55}},
			[]verifyItem{{"REDACTED", 66}},
			[]verifyItem{{"REDACTED", 44}, {"REDACTED", 55}, {"REDACTED", 66}},
		},
		2: { //
			[]verifyItem{{"REDACTED", 44}, {"REDACTED", 66}},
			[]verifyItem{{"REDACTED", 55}},
			[]verifyItem{{"REDACTED", 44}, {"REDACTED", 55}, {"REDACTED", 66}},
		},
		3: { //
			[]verifyItem{{"REDACTED", 44}, {"REDACTED", 66}, {"REDACTED", 99}},
			[]verifyItem{{"REDACTED", 55}, {"REDACTED", 77}, {"REDACTED", 88}},
			[]verifyItem{{"REDACTED", 44}, {"REDACTED", 55}, {"REDACTED", 66}, {"REDACTED", 77}, {"REDACTED", 88}, {"REDACTED", 99}},
		},
		//
		4: { //
			[]verifyItem{{"REDACTED", 111}, {"REDACTED", 22}},
			[]verifyItem{{"REDACTED", 11}},
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 22}},
		},
		5: { //
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 222}},
			[]verifyItem{{"REDACTED", 22}},
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 22}},
		},
		6: { //
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 222}, {"REDACTED", 33}},
			[]verifyItem{{"REDACTED", 22}},
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 33}},
		},
		7: { //
			[]verifyItem{{"REDACTED", 111}, {"REDACTED", 222}, {"REDACTED", 333}},
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 33}},
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 33}},
		},
		//
		8: {
			[]verifyItem{{"REDACTED", 111}, {"REDACTED", 22}},
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 33}, {"REDACTED", 44}},
			[]verifyItem{{"REDACTED", 11}, {"REDACTED", 22}, {"REDACTED", 33}, {"REDACTED", 44}},
		},
	}

	for i, tt := range itemAssignRevisionsFundamentalVerifies {
		//
		itemAssign := generateFreshAssessorAssign(tt.initiateValues)

		//
		itemCatalog := generateFreshAssessorCatalog(tt.reviseValues)
		itemAssign.executeRevisions(itemCatalog)

		//
		assert.Equal(t, towardVerifyItemCatalog(itemAssign.Assessors), tt.anticipatedValues, "REDACTED", i)
	}
}

type verifyVERAssignConfig struct {
	alias         string
	initiateValues    []verifyItem
	removedValues  []verifyItem
	revisedValues  []verifyItem
	appendedValues    []verifyItem
	anticipatedValues []verifyItem
	expirationFault       error
}

func arbitraryVerifyVERAssignConfig(nthFoundation, nthAppendMaximum int) verifyVERAssignConfig {
	if nthFoundation <= 0 || nthAppendMaximum < 0 {
		panic(fmt.Sprintf("REDACTED", nthFoundation, nthAppendMaximum))
	}

	const maximumPotency = 1000
	var nthAged, nthRemove, nthAltered, nthAppend int

	nthAged = int(commitrand.Uintn()%uint(nthFoundation)) + 1
	if nthFoundation-nthAged > 0 {
		nthRemove = int(commitrand.Uintn() % uint(nthFoundation-nthAged))
	}
	nthAltered = nthFoundation - nthAged - nthRemove

	if nthAppendMaximum > 0 {
		nthAppend = commitrand.Int()%nthAppendMaximum + 1
	}

	cfg := verifyVERAssignConfig{}

	cfg.initiateValues = make([]verifyItem, nthFoundation)
	cfg.removedValues = make([]verifyItem, nthRemove)
	cfg.appendedValues = make([]verifyItem, nthAppend)
	cfg.revisedValues = make([]verifyItem, nthAltered)
	cfg.anticipatedValues = make([]verifyItem, nthFoundation-nthRemove+nthAppend)

	for i := 0; i < nthFoundation; i++ {
		cfg.initiateValues[i] = verifyItem{fmt.Sprintf("REDACTED", i), int64(commitrand.Uintn()%maximumPotency + 1)}
		if i < nthAged {
			cfg.anticipatedValues[i] = cfg.initiateValues[i]
		}
		if i >= nthAged && i < nthAged+nthAltered {
			cfg.revisedValues[i-nthAged] = verifyItem{fmt.Sprintf("REDACTED", i), int64(commitrand.Uintn()%maximumPotency + 1)}
			cfg.anticipatedValues[i] = cfg.revisedValues[i-nthAged]
		}
		if i >= nthAged+nthAltered {
			cfg.removedValues[i-nthAged-nthAltered] = verifyItem{fmt.Sprintf("REDACTED", i), 0}
		}
	}

	for i := nthFoundation; i < nthFoundation+nthAppend; i++ {
		cfg.appendedValues[i-nthFoundation] = verifyItem{fmt.Sprintf("REDACTED", i), int64(commitrand.Uintn()%maximumPotency + 1)}
		cfg.anticipatedValues[i-nthRemove] = cfg.appendedValues[i-nthFoundation]
	}

	sort.Sort(verifyValuesViaBallotingPotency(cfg.initiateValues))
	sort.Sort(verifyValuesViaBallotingPotency(cfg.removedValues))
	sort.Sort(verifyValuesViaBallotingPotency(cfg.revisedValues))
	sort.Sort(verifyValuesViaBallotingPotency(cfg.appendedValues))
	sort.Sort(verifyValuesViaBallotingPotency(cfg.anticipatedValues))

	return cfg
}

func executeModificationsTowardItemAssign(t *testing.T, expirationFault error, itemAssign *AssessorAssign, valuesCatalogs ...[]verifyItem) {
	modifications := make([]verifyItem, 0)
	for _, valuesCatalog := range valuesCatalogs {
		modifications = append(modifications, valuesCatalog...)
	}
	itemCatalog := generateFreshAssessorCatalog(modifications)
	err := itemAssign.ReviseUsingModifyAssign(itemCatalog)
	if expirationFault != nil {
		assert.Equal(t, expirationFault, err)
	} else {
		assert.NoError(t, err)
	}
}

func VerifyItemAssignReviseUrgencySequenceVerifies(t *testing.T) {
	const nthMaximumReferendums int32 = 5000

	verifyScenarios := []verifyVERAssignConfig{
		0: { //
			initiateValues:    []verifyItem{{"REDACTED", 1000}, {"REDACTED", 1}, {"REDACTED", 1}},
			removedValues:  []verifyItem{{"REDACTED", 0}},
			revisedValues:  []verifyItem{},
			appendedValues:    []verifyItem{},
			anticipatedValues: []verifyItem{{"REDACTED", 1}, {"REDACTED", 1}},
		},
		1: { //
			initiateValues:    []verifyItem{{"REDACTED", 1000}, {"REDACTED", 10}, {"REDACTED", 1}},
			removedValues:  []verifyItem{{"REDACTED", 0}},
			revisedValues:  []verifyItem{},
			appendedValues:    []verifyItem{},
			anticipatedValues: []verifyItem{{"REDACTED", 10}, {"REDACTED", 1}},
		},
		2: { //
			initiateValues:    []verifyItem{{"REDACTED", 1000}, {"REDACTED", 2}, {"REDACTED", 1}},
			removedValues:  []verifyItem{{"REDACTED", 0}},
			revisedValues:  []verifyItem{{"REDACTED", 1}},
			appendedValues:    []verifyItem{{"REDACTED", 50}, {"REDACTED", 40}},
			anticipatedValues: []verifyItem{{"REDACTED", 50}, {"REDACTED", 40}, {"REDACTED", 1}, {"REDACTED", 1}},
		},

		//
		//
		//
		3: arbitraryVerifyVERAssignConfig(100, 10),

		4: arbitraryVerifyVERAssignConfig(1000, 100),

		5: arbitraryVerifyVERAssignConfig(10, 100),

		6: arbitraryVerifyVERAssignConfig(100, 1000),

		7: arbitraryVerifyVERAssignConfig(1000, 1000),
	}

	for _, cfg := range verifyScenarios {

		//
		itemAssign := generateFreshAssessorAssign(cfg.initiateValues)
		validateAssessorAssign(t, itemAssign)

		//
		validateItemAssignReviseUrgencySequence(t, itemAssign, cfg, nthMaximumReferendums)
	}
}

func validateItemAssignReviseUrgencySequence(t *testing.T, itemAssign *AssessorAssign, cfg verifyVERAssignConfig, nthMaximumReferendums int32) {
	//
	itemAssign.AdvanceNominatorUrgency(commitrand.Int31n()%nthMaximumReferendums + 1)

	//
	executeModificationsTowardItemAssign(t, nil, itemAssign, cfg.appendedValues, cfg.revisedValues, cfg.removedValues)

	//
	assert.Equal(t, cfg.anticipatedValues, towardVerifyItemCatalog(itemAssign.Assessors))
	validateAssessorAssign(t, itemAssign)

	//
	//
	//
	if len(cfg.appendedValues) > 0 {
		revisedValuesUrgencyArranged := assessorCatalogDuplicate(itemAssign.Assessors)
		sort.Sort(assessorsViaUrgency(revisedValuesUrgencyArranged))

		appendedValuesUrgencySegment := revisedValuesUrgencyArranged[:len(cfg.appendedValues)]
		sort.Sort(AssessorsViaBallotingPotency(appendedValuesUrgencySegment))
		assert.Equal(t, cfg.appendedValues, towardVerifyItemCatalog(appendedValuesUrgencySegment))

		//
		anticipatedUrgency := appendedValuesUrgencySegment[0].NominatorUrgency
		for _, val := range appendedValuesUrgencySegment[1:] {
			assert.Equal(t, anticipatedUrgency, val.NominatorUrgency)
		}
	}
}

func VerifyFreshAssessorAssignOriginatingCurrentAssessors(t *testing.T) {
	extent := 5
	values := make([]*Assessor, extent)
	for i := 0; i < extent; i++ {
		pv := FreshSimulatePRV()
		values[i] = pv.DeriveWithinAssessor(int64(i + 1))
	}
	itemAssign := FreshAssessorAssign(values)
	itemAssign.AdvanceNominatorUrgency(5)

	freshItemAssign := FreshAssessorAssign(itemAssign.Assessors)
	assert.NotEqual(t, itemAssign, freshItemAssign)

	currentItemAssign, err := AssessorAssignOriginatingCurrentAssessors(itemAssign.Assessors)
	assert.NoError(t, err)
	assert.Equal(t, itemAssign, currentItemAssign)
	assert.Equal(t, itemAssign.DuplicateAdvanceNominatorUrgency(3), currentItemAssign.DuplicateAdvanceNominatorUrgency(3))
}

func VerifyItemAssignReviseOverrunAssociated(t *testing.T) {
	verifyScenarios := []verifyVERAssignConfig{
		{
			alias:         "REDACTED",
			initiateValues:    []verifyItem{{"REDACTED", MaximumSumBallotingPotency - 1}, {"REDACTED", 1}},
			revisedValues:  []verifyItem{{"REDACTED", MaximumSumBallotingPotency - 1}, {"REDACTED", 1}},
			anticipatedValues: []verifyItem{{"REDACTED", MaximumSumBallotingPotency - 1}, {"REDACTED", 1}},
			expirationFault:       nil,
		},
		{
			//
			//
			alias:         "REDACTED",
			initiateValues:    []verifyItem{{"REDACTED", MaximumSumBallotingPotency - 1}, {"REDACTED", 1}},
			revisedValues:  []verifyItem{{"REDACTED", MaximumSumBallotingPotency/2 - 1}, {"REDACTED", MaximumSumBallotingPotency / 2}},
			anticipatedValues: []verifyItem{{"REDACTED", MaximumSumBallotingPotency / 2}, {"REDACTED", MaximumSumBallotingPotency/2 - 1}},
			expirationFault:       nil,
		},
		{
			alias:         "REDACTED",
			initiateValues:    []verifyItem{{"REDACTED", MaximumSumBallotingPotency - 2}, {"REDACTED", 1}, {"REDACTED", 1}},
			removedValues:  []verifyItem{{"REDACTED", 0}},
			appendedValues:    []verifyItem{{"REDACTED", MaximumSumBallotingPotency - 2}},
			anticipatedValues: []verifyItem{{"REDACTED", MaximumSumBallotingPotency - 2}, {"REDACTED", 1}, {"REDACTED", 1}},
			expirationFault:       nil,
		},
		{
			alias: "REDACTED",
			initiateValues: []verifyItem{
				{"REDACTED", MaximumSumBallotingPotency / 4},
				{"REDACTED", MaximumSumBallotingPotency / 4},
				{"REDACTED", MaximumSumBallotingPotency / 4},
				{"REDACTED", MaximumSumBallotingPotency / 4},
			},
			removedValues: []verifyItem{{"REDACTED", 0}},
			revisedValues: []verifyItem{
				{"REDACTED", MaximumSumBallotingPotency/2 - 2}, {"REDACTED", MaximumSumBallotingPotency/2 - 3}, {"REDACTED", 2},
			},
			appendedValues: []verifyItem{{"REDACTED", 3}},
			anticipatedValues: []verifyItem{
				{"REDACTED", MaximumSumBallotingPotency/2 - 2}, {"REDACTED", MaximumSumBallotingPotency/2 - 3}, {"REDACTED", 3}, {"REDACTED", 2},
			},
			expirationFault: nil,
		},
		{
			alias: "REDACTED",
			initiateValues: []verifyItem{
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
			revisedValues: []verifyItem{
				{"REDACTED", MaximumSumBallotingPotency},
				{"REDACTED", MaximumSumBallotingPotency},
				{"REDACTED", MaximumSumBallotingPotency},
				{"REDACTED", MaximumSumBallotingPotency},
				{"REDACTED", MaximumSumBallotingPotency},
				{"REDACTED", MaximumSumBallotingPotency},
				{"REDACTED", MaximumSumBallotingPotency},
				{"REDACTED", MaximumSumBallotingPotency},
				{"REDACTED", 8},
			},
			anticipatedValues: []verifyItem{
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
			expirationFault: FaultSumBallotingPotencyOverrun,
		},
	}

	for _, tt := range verifyScenarios {
		t.Run(tt.alias, func(t *testing.T) {
			itemAssign := generateFreshAssessorAssign(tt.initiateValues)
			validateAssessorAssign(t, itemAssign)

			//
			executeModificationsTowardItemAssign(t, tt.expirationFault, itemAssign, tt.appendedValues, tt.revisedValues, tt.removedValues)

			//
			assert.Equal(t, tt.anticipatedValues, towardVerifyItemCatalog(itemAssign.Assessors))
			validateAssessorAssign(t, itemAssign)
		})
	}
}

func VerifySecureMultiply(t *testing.T) {
	verifyScenarios := []struct {
		a        int64
		b        int64
		c        int64
		overrun bool
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
		c, overrun := secureMultiply(tc.a, tc.b)
		assert.Equal(t, tc.c, c, "REDACTED", i)
		assert.Equal(t, tc.overrun, overrun, "REDACTED", i)
	}
}

func VerifyAssessorAssignSchemaBuffer(t *testing.T) {
	validset, _ := ArbitraryAssessorAssign(10, 100)
	validset2, _ := ArbitraryAssessorAssign(10, 100)
	validset2.Assessors[0] = &Assessor{}

	validset3, _ := ArbitraryAssessorAssign(10, 100)
	validset3.Nominator = nil

	validset4, _ := ArbitraryAssessorAssign(10, 100)
	validset4.Nominator = &Assessor{}

	verifyScenarios := []struct {
		msg      string
		v1       *AssessorAssign
		expirationPhase1 bool
		expirationPhase2 bool
	}{
		{"REDACTED", validset, true, true},
		{"REDACTED", validset2, false, false},
		{"REDACTED", validset3, false, false},
		{"REDACTED", validset4, false, false},
		{"REDACTED", &AssessorAssign{}, true, false},
		{"REDACTED", nil, true, false},
	}
	for _, tc := range verifyScenarios {
		schemaItemAssign, err := tc.v1.TowardSchema()
		if tc.expirationPhase1 {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}

		itemAssign, err := AssessorAssignOriginatingSchema(schemaItemAssign)
		if tc.expirationPhase2 {
			require.NoError(t, err, tc.msg)
			require.EqualValues(t, tc.v1, itemAssign, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}

//
//
type assessorsViaUrgency []*Assessor

func (validz assessorsViaUrgency) Len() int {
	return len(validz)
}

func (validz assessorsViaUrgency) Inferior(i, j int) bool {
	if validz[i].NominatorUrgency < validz[j].NominatorUrgency {
		return true
	}
	if validz[i].NominatorUrgency > validz[j].NominatorUrgency {
		return false
	}
	return bytes.Compare(validz[i].Location, validz[j].Location) < 0
}

func (validz assessorsViaUrgency) Exchange(i, j int) {
	validz[i], validz[j] = validz[j], validz[i]
}

//

type verifyValuesViaBallotingPotency []verifyItem

func (titems verifyValuesViaBallotingPotency) Len() int {
	return len(titems)
}

func (titems verifyValuesViaBallotingPotency) Inferior(i, j int) bool {
	if titems[i].potency == titems[j].potency {
		return bytes.Compare([]byte(titems[i].alias), []byte(titems[j].alias)) == -1
	}
	return titems[i].potency > titems[j].potency
}

func (titems verifyValuesViaBallotingPotency) Exchange(i, j int) {
	titems[i], titems[j] = titems[j], titems[i]
}

//
//
func AssessmentRevisions(b *testing.B) {
	const (
		n = 100
		m = 2000
	)
	//
	vs := make([]*Assessor, n)
	for j := 0; j < n; j++ {
		vs[j] = freshAssessor([]byte(fmt.Sprintf("REDACTED", j)), 100)
	}
	itemAssign := FreshAssessorAssign(vs)
	l := len(itemAssign.Assessors)

	//
	freshItemCatalog := make([]*Assessor, m)
	for j := 0; j < m; j++ {
		freshItemCatalog[j] = freshAssessor([]byte(fmt.Sprintf("REDACTED", j+l)), 1000)
	}

	for b.Loop() {
		//
		itemAssignDuplicate := itemAssign.Duplicate()
		assert.NoError(b, itemAssignDuplicate.ReviseUsingModifyAssign(freshItemCatalog))
	}
}

func VerifyValidateEndorseUsingUnfitNominatorToken(t *testing.T) {
	vs := &AssessorAssign{
		Assessors: []*Assessor{{}, {}},
	}
	endorse := &Endorse{
		Altitude:     100,
		Notations: []EndorseSignature{{}, {}},
	}
	var bid LedgerUUID
	cid := "REDACTED"
	err := vs.ValidateEndorse(cid, bid, 100, endorse)
	assert.Error(t, err)
}

func VerifyValidateEndorseUniqueUsingUnfitNotations(t *testing.T) {
	vs := &AssessorAssign{
		Assessors: []*Assessor{{}, {}},
	}
	endorse := &Endorse{
		Altitude:     100,
		Notations: []EndorseSignature{{}, {}},
	}
	cid := "REDACTED"
	ballotingPotencyRequired := vs.SumBallotingPotency() * 2 / 3

	//
	bypass := func(c EndorseSignature) bool { return c.LedgerUUIDMarker == LedgerUUIDMarkerMissing }

	//
	tally := func(c EndorseSignature) bool { return c.LedgerUUIDMarker == LedgerUUIDMarkerEndorse }

	err := validateEndorseUnique(cid, vs, endorse, ballotingPotencyRequired, bypass, tally, true, true, nil)
	require.Error(t, err)

	stash := FreshSigningStash()
	err = validateEndorseUnique(cid, vs, endorse, ballotingPotencyRequired, bypass, tally, true, true, stash)
	require.Error(t, err)
	require.Equal(t, 0, stash.Len())
}

func Testvalidset_Allequaltokenkinds(t *testing.T) {
	verifyScenarios := []struct {
		values     *AssessorAssign
		identicalKind bool
	}{
		{
			values:     FreshAssessorAssign([]*Assessor{}),
			identicalKind: true,
		},
		{
			values:     arbitraryAssessorAssign(1),
			identicalKind: true,
		},
		{
			values:     arbitraryAssessorAssign(2),
			identicalKind: true,
		},
		{
			values:     FreshAssessorAssign([]*Assessor{arbitraryAssessor(100), FreshAssessor(ellipticp256.ProducePrivateToken().PublicToken(), 200)}),
			identicalKind: false,
		},
	}

	for i, tc := range verifyScenarios {
		if tc.identicalKind {
			assert.True(t, tc.values.EveryTokensPossessIdenticalKind(), "REDACTED", i)
		} else {
			assert.False(t, tc.values.EveryTokensPossessIdenticalKind(), "REDACTED", i)
		}
	}
}

func Testvalidset_Totalsafevotingpotency(t *testing.T) {
	verifyScenarios := []struct {
		alias          string
		assessors    []*Assessor
		anticipatedPotency int64
		anticipateFailure   bool
		failureIncludes string
	}{
		{
			alias: "REDACTED",
			assessors: []*Assessor{
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 100),
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 200),
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 300),
			},
			anticipatedPotency: 600,
			anticipateFailure:   false,
		},
		{
			alias:          "REDACTED",
			assessors:    []*Assessor{},
			anticipatedPotency: 0,
			anticipateFailure:   false,
		},
		{
			alias:          "REDACTED",
			assessors:    nil,
			anticipatedPotency: 0,
			anticipateFailure:   false,
		},
		{
			alias: "REDACTED",
			assessors: []*Assessor{
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 1000),
			},
			anticipatedPotency: 1000,
			anticipateFailure:   false,
		},
		{
			alias: "REDACTED",
			assessors: []*Assessor{
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), MaximumSumBallotingPotency),
			},
			anticipatedPotency: MaximumSumBallotingPotency,
			anticipateFailure:   false,
		},
		{
			alias: "REDACTED",
			assessors: []*Assessor{
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), MaximumSumBallotingPotency-100),
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 100),
			},
			anticipatedPotency: MaximumSumBallotingPotency,
			anticipateFailure:   false,
		},
		{
			alias: "REDACTED",
			assessors: []*Assessor{
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), MaximumSumBallotingPotency/2+1),
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), MaximumSumBallotingPotency/2+1),
			},
			anticipatedPotency: 0,
			anticipateFailure:   true,
			failureIncludes: "REDACTED",
		},
		{
			alias: "REDACTED",
			assessors: []*Assessor{
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), MaximumSumBallotingPotency),
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 1),
			},
			anticipatedPotency: 0,
			anticipateFailure:   true,
			failureIncludes: "REDACTED",
		},
		{
			alias: "REDACTED",
			assessors: []*Assessor{
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), math.MaxInt64/2),
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), math.MaxInt64/2),
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 100),
			},
			anticipatedPotency: 0,
			anticipateFailure:   true,
			failureIncludes: "REDACTED",
		},
		{
			alias: "REDACTED",
			assessors: []*Assessor{
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 100),
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 0),
				FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 200),
			},
			anticipatedPotency: 300,
			anticipateFailure:   false,
		},
		{
			alias: "REDACTED",
			assessors: func() []*Assessor {
				values := make([]*Assessor, 100)
				for i := 0; i < 100; i++ {
					values[i] = FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 1000)
				}
				return values
			}(),
			anticipatedPotency: 100000,
			anticipateFailure:   false,
		},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.alias, func(t *testing.T) {
			//
			itemAssign := &AssessorAssign{
				Assessors: tc.assessors,
			}

			//
			sumPotency, err := itemAssign.SumBallotingPotencySecure()

			//
			if tc.anticipateFailure {
				require.Error(t, err, "REDACTED")
				if tc.failureIncludes != "REDACTED" {
					require.Contains(t, err.Error(), tc.failureIncludes,
						"REDACTED")
				}
				require.Equal(t, tc.anticipatedPotency, sumPotency,
					"REDACTED", tc.anticipatedPotency)
			} else {
				require.NoError(t, err, "REDACTED", err)
				require.Equal(t, tc.anticipatedPotency, sumPotency,
					"REDACTED", tc.anticipatedPotency)
			}
		})
	}
}
