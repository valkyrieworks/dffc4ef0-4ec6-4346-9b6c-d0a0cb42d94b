package status_test

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
func configureVerifyInstance(t *testing.T) (func(t *testing.T), dbm.DB, sm.Status) {
	t.Helper()
	dismantleDepressed, statusDatastore, status, _ := configureVerifyInstanceUsingDepot(t)
	return dismantleDepressed, statusDatastore, status
}

//
func configureVerifyInstanceUsingDepot(t *testing.T) (func(t *testing.T), dbm.DB, sm.Status, sm.Depot) {
	t.Helper()
	settings := verify.RestoreVerifyOrigin("REDACTED")
	datastoreKind := dbm.OriginKind(settings.DatastoreRepository)
	statusDatastore, err := dbm.FreshDatastore("REDACTED", datastoreKind, settings.DatastorePath())
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	require.NoError(t, err)
	status, err := statusDepot.FetchOriginatingDatastoreEitherInaugurationRecord(settings.InaugurationRecord())
	require.NoError(t, err, "REDACTED")
	err = statusDepot.Persist(status)
	require.NoError(t, err)

	dismantleDepressed := func(t *testing.T) {
		t.Helper()
		os.RemoveAll(settings.OriginPath)
	}

	return dismantleDepressed, statusDatastore, status, statusDepot
}

//
func VerifyStatusDuplicate(t *testing.T) {
	dismantleDepressed, _, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)
	affirm := assert.New(t)

	statusDuplicate := status.Duplicate()

	assert.True(status.Matches(statusDuplicate),
		fmt.Sprintf("REDACTED",
			statusDuplicate, status))

	statusDuplicate.FinalLedgerAltitude++
	statusDuplicate.FinalAssessors = status.Assessors
	assert.False(status.Matches(statusDuplicate), fmt.Sprintf(`REDACTEDe
REDACTED`, status))
}

//
func VerifyCreateInaugurationStatusVoidAssessors(t *testing.T) {
	doc := kinds.OriginPaper{
		SuccessionUUID:    "REDACTED",
		Assessors: nil,
	}
	require.Nil(t, doc.CertifyAlsoFinish())
	status, err := sm.CreateInaugurationStatus(&doc)
	require.Nil(t, err)
	require.Equal(t, 0, len(status.Assessors.Assessors))
	require.Equal(t, 0, len(status.FollowingAssessors.Assessors))
}

//
func VerifyStatusPersistFetch(t *testing.T) {
	dismantleDepressed, statusDatastore, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	affirm := assert.New(t)

	status.FinalLedgerAltitude++
	status.FinalAssessors = status.Assessors
	err := statusDepot.Persist(status)
	require.NoError(t, err)

	retrievedStatus, err := statusDepot.Fetch()
	require.NoError(t, err)
	assert.True(status.Matches(retrievedStatus),
		fmt.Sprintf("REDACTED",
			retrievedStatus, status))
}

//
func VerifyCulminateLedgerRepliesPersistFetch1(t *testing.T) {
	dismantleDepressed, statusDatastore, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	affirm := assert.New(t)

	status.FinalLedgerAltitude++

	//
	ledger, err := createLedger(status, 2, new(kinds.Endorse))
	require.NoError(t, err)

	ifaceReplies := new(iface.ReplyCulminateLedger)
	dtrans := make([]*iface.InvokeTransferOutcome, 2)
	ifaceReplies.TransferOutcomes = dtrans

	ifaceReplies.TransferOutcomes[0] = &iface.InvokeTransferOutcome{Data: []byte("REDACTED"), Incidents: nil}
	ifaceReplies.TransferOutcomes[1] = &iface.InvokeTransferOutcome{Data: []byte("REDACTED"), Log: "REDACTED", Incidents: nil}
	ifaceReplies.AssessorRevisions = []iface.AssessorRevise{
		kinds.Temp2buffer.FreshAssessorRevise(edwards25519.ProducePrivateToken().PublicToken(), 10),
	}

	ifaceReplies.PlatformDigest = make([]byte, 1)

	err = statusDepot.PersistCulminateLedgerReply(ledger.Altitude, ifaceReplies)
	require.NoError(t, err)
	retrievedIfaceReplies, err := statusDepot.FetchCulminateLedgerReply(ledger.Altitude)
	assert.NoError(err)
	assert.Equal(ifaceReplies, retrievedIfaceReplies)
}

//
func VerifyCulminateLedgerRepliesPersistFetch2(t *testing.T) {
	dismantleDepressed, statusDatastore, _ := configureVerifyInstance(t)
	defer dismantleDepressed(t)
	affirm := assert.New(t)

	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	scenarios := [...]struct {
		//
		//
		appended    []*iface.InvokeTransferOutcome
		anticipated []*iface.InvokeTransferOutcome
	}{
		0: {
			nil,
			nil,
		},
		1: {
			[]*iface.InvokeTransferOutcome{
				{Cipher: 32, Data: []byte("REDACTED"), Log: "REDACTED"},
			},
			[]*iface.InvokeTransferOutcome{
				{Cipher: 32, Data: []byte("REDACTED")},
			},
		},
		2: {
			[]*iface.InvokeTransferOutcome{
				{Cipher: 383},
				{
					Data: []byte("REDACTED"),
					Incidents: []iface.Incident{
						{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED"}}},
						{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED"}}},
					},
				},
			},
			[]*iface.InvokeTransferOutcome{
				{Cipher: 383, Data: nil},
				{Cipher: 0, Data: []byte("REDACTED"), Incidents: []iface.Incident{
					{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED"}}},
					{Kind: "REDACTED", Properties: []iface.IncidentProperty{{Key: "REDACTED", Datum: "REDACTED"}}},
				}},
			},
		},
		3: {
			nil,
			nil,
		},
		4: {
			[]*iface.InvokeTransferOutcome{nil},
			nil,
		},
	}

	//
	for i := range scenarios {
		h := int64(i + 1)
		res, err := statusDepot.FetchCulminateLedgerReply(h)
		assert.Error(err, "REDACTED", i, res)
	}

	//
	for i, tc := range scenarios {
		h := int64(i + 1) //
		replies := &iface.ReplyCulminateLedger{
			TransferOutcomes: tc.appended,
			PlatformDigest:   []byte(fmt.Sprintf("REDACTED", h)),
		}
		err := statusDepot.PersistCulminateLedgerReply(h, replies)
		require.NoError(t, err)
	}

	//
	for i, tc := range scenarios {
		h := int64(i + 1)
		res, err := statusDepot.FetchCulminateLedgerReply(h)
		if assert.NoError(err, "REDACTED", i) {
			t.Log(res)
			replies := &iface.ReplyCulminateLedger{
				TransferOutcomes: tc.anticipated,
				PlatformDigest:   []byte(fmt.Sprintf("REDACTED", h)),
			}
			assert.Equal(sm.TransferOutcomesDigest(replies.TransferOutcomes), sm.TransferOutcomesDigest(res.TransferOutcomes), "REDACTED", i)
		}
	}
}

//
func VerifyAssessorPlainPersistFetch(t *testing.T) {
	dismantleDepressed, statusDatastore, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)
	affirm := assert.New(t)

	statusstore := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	//
	_, err := statusstore.FetchAssessors(0)
	assert.IsType(sm.FaultNegativeItemAssignForeachAltitude{}, err, "REDACTED")

	//
	v, err := statusstore.FetchAssessors(1)
	assert.Nil(err, "REDACTED")
	assert.Equal(v.Digest(), status.Assessors.Digest(), "REDACTED")

	//
	v, err = statusstore.FetchAssessors(2)
	assert.Nil(err, "REDACTED")
	assert.Equal(v.Digest(), status.FollowingAssessors.Digest(), "REDACTED")

	//
	status.FinalLedgerAltitude++
	followingAltitude := status.FinalLedgerAltitude + 1
	err = statusstore.Persist(status)
	require.NoError(t, err)
	vp0, err := statusstore.FetchAssessors(followingAltitude + 0)
	assert.Nil(err, "REDACTED")
	vp1, err := statusstore.FetchAssessors(followingAltitude + 1)
	assert.Nil(err, "REDACTED")
	assert.Equal(vp0.Digest(), status.Assessors.Digest(), "REDACTED")
	assert.Equal(vp1.Digest(), status.FollowingAssessors.Digest(), "REDACTED")
}

//
func VerifySingleAssessorModificationsPersistFetch(t *testing.T) {
	dismantleDepressed, statusDatastore, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	//
	alterationElevations := []int64{1, 2, 4, 5, 10, 15, 16, 17, 20}
	N := len(alterationElevations)

	//
	//
	utmostAltitude := alterationElevations[N-1] + 5
	alterationPosition := 0
	_, val := status.Assessors.ObtainViaOrdinal(0)
	potency := val.BallotingPotency
	var err error
	var assessorRevisions []*kinds.Assessor
	for i := int64(1); i < utmostAltitude; i++ {
		//
		if alterationPosition < len(alterationElevations) && i == alterationElevations[alterationPosition] {
			alterationPosition++
			potency++
		}
		heading, ledgerUUID, replies := createHeadlineFragmentsRepliesItemPotencyAlteration(status, potency)
		assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions(replies.AssessorRevisions)
		require.NoError(t, err)
		status, err = sm.ReviseStatus(status, ledgerUUID, &heading, replies, assessorRevisions)
		require.NoError(t, err)
		err = statusDepot.Persist(status)
		require.NoError(t, err)
	}

	//
	verifyScenarios := make([]int64, utmostAltitude)
	alterationPosition = 0
	potency = val.BallotingPotency
	for i := int64(1); i < utmostAltitude+1; i++ {
		//
		//
		if alterationPosition < len(alterationElevations) && i == alterationElevations[alterationPosition]+1 {
			alterationPosition++
			potency++
		}
		verifyScenarios[i-1] = potency
	}

	for i, potency := range verifyScenarios {
		v, err := statusDepot.FetchAssessors(int64(i + 1 + 1)) //
		assert.Nil(t, err, fmt.Sprintf("REDACTED", i))
		assert.Equal(t, v.Extent(), 1, "REDACTED", v.Extent())
		_, val := v.ObtainViaOrdinal(0)

		assert.Equal(t, val.BallotingPotency, potency, fmt.Sprintf(`REDACTEDt
REDACTED`, i))
	}
}

func VerifyNominatorRecurrence(t *testing.T) {
	//
	verifyScenarios := []struct {
		strengths []int64
	}{
		//
		{[]int64{1, 1}},
		{[]int64{1, 2}},
		{[]int64{1, 100}},
		{[]int64{5, 5}},
		{[]int64{5, 100}},
		{[]int64{50, 50}},
		{[]int64{50, 100}},
		{[]int64{1, 1000}},

		//
		{[]int64{1, 1, 1}},
		{[]int64{1, 2, 3}},
		{[]int64{1, 2, 3}},
		{[]int64{1, 1, 10}},
		{[]int64{1, 1, 100}},
		{[]int64{1, 10, 100}},
		{[]int64{1, 1, 1000}},
		{[]int64{1, 10, 1000}},
		{[]int64{1, 100, 1000}},

		//
		{[]int64{1, 1, 1, 1}},
		{[]int64{1, 2, 3, 4}},
		{[]int64{1, 1, 1, 10}},
		{[]int64{1, 1, 1, 100}},
		{[]int64{1, 1, 1, 1000}},
		{[]int64{1, 1, 10, 100}},
		{[]int64{1, 1, 10, 1000}},
		{[]int64{1, 1, 100, 1000}},
		{[]int64{1, 10, 100, 1000}},
	}

	for instanceCount, verifyInstance := range verifyScenarios {
		//
		//
		for i := 0; i < 5; i++ {
			itemAssign := produceItemAssignUsingStrengths(verifyInstance.strengths)
			verifyNominatorRecurrence(t, instanceCount, itemAssign)
		}
	}

	//
	maximumValues := 100
	maximumPotency := 1000
	nthVerifyScenarios := 5
	for i := 0; i < nthVerifyScenarios; i++ {
		N := commitrand.Int()%maximumValues + 1
		values := make([]*kinds.Assessor, N)
		sumBallotPotency := int64(0)
		for j := 0; j < N; j++ {
			//
			ballotPotency := int64(commitrand.Int()%maximumPotency) + 1
			sumBallotPotency += ballotPotency
			privateItem := kinds.FreshSimulatePRV()
			publicToken, err := privateItem.ObtainPublicToken()
			require.NoError(t, err)
			val := kinds.FreshAssessor(publicToken, ballotPotency)
			val.NominatorUrgency = commitrand.Int64n()
			values[j] = val
		}
		itemAssign := kinds.FreshAssessorAssign(values)
		itemAssign.RecalibrateUrgencies(sumBallotPotency)
		verifyNominatorRecurrence(t, i, itemAssign)
	}
}

//
func produceItemAssignUsingStrengths(strengths []int64) *kinds.AssessorAssign {
	extent := len(strengths)
	values := make([]*kinds.Assessor, extent)
	sumBallotPotency := int64(0)
	for i := 0; i < extent; i++ {
		sumBallotPotency += strengths[i]
		val := kinds.FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), strengths[i])
		val.NominatorUrgency = commitrand.Int64n()
		values[i] = val
	}
	itemAssign := kinds.FreshAssessorAssign(values)
	itemAssign.RecalibrateUrgencies(sumBallotPotency)
	return itemAssign
}

//
func verifyNominatorRecurrence(t *testing.T, instanceCount int, itemAssign *kinds.AssessorAssign) {
	N := itemAssign.Extent()
	sumPotency := itemAssign.SumBallotingPotency()

	//
	executeMultiply := 1
	iterations := int(sumPotency) * executeMultiply
	recurrences := make([]int, N)
	for i := 0; i < iterations; i++ {
		item := itemAssign.ObtainNominator()
		idx, _ := itemAssign.ObtainViaLocation(item.Location)
		recurrences[idx]++
		itemAssign.AdvanceNominatorUrgency(1)
	}

	//
	for i, recurrence := range recurrences {
		_, val := itemAssign.ObtainViaOrdinal(int32(i))
		anticipateRecurrence := int(val.BallotingPotency) * executeMultiply
		attainedRecurrence := recurrence
		abs := int(math.Abs(float64(anticipateRecurrence - attainedRecurrence)))

		//
		//
		//
		//
		restricted := N - 1
		require.True(
			t,
			abs <= restricted,
			fmt.Sprintf("REDACTED", instanceCount, i, N, attainedRecurrence, anticipateRecurrence),
		)
	}
}

//
//
func VerifyNominatorUrgencyExecutesNegationObtainRestoreTowardNull(t *testing.T) {
	dismantleDepressed, _, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)
	item1ballotingPotency := int64(10)
	assessor1keyToken := edwards25519.ProducePrivateToken().PublicToken()
	valid1 := &kinds.Assessor{Location: assessor1keyToken.Location(), PublicToken: assessor1keyToken, BallotingPotency: item1ballotingPotency}

	status.Assessors = kinds.FreshAssessorAssign([]*kinds.Assessor{valid1})
	status.FollowingAssessors = status.Assessors

	//
	assert.EqualValues(t, 0, valid1.NominatorUrgency)

	ledger, err := createLedger(status, status.FinalLedgerAltitude+1, new(kinds.Endorse))
	require.NoError(t, err)
	bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
	require.NoError(t, err)
	ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}
	ifaceReplies := &iface.ReplyCulminateLedger{}
	assessorRevisions, err := kinds.Buffer2temp.AssessorRevisions(ifaceReplies.AssessorRevisions)
	require.NoError(t, err)
	modifiedStatus, err := sm.ReviseStatus(status, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
	assert.NoError(t, err)
	presentSum := item1ballotingPotency
	//
	assert.Equal(t, 0+item1ballotingPotency-presentSum, modifiedStatus.FollowingAssessors.Assessors[0].NominatorUrgency)

	//
	item2publicToken := edwards25519.ProducePrivateToken().PublicToken()
	item2ballotingPotency := int64(100)
	fvp, err := cryptocode.PublicTokenTowardSchema(item2publicToken)
	require.NoError(t, err)

	reviseAppendItem := iface.AssessorRevise{PublicToken: fvp, Potency: item2ballotingPotency}
	assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions([]iface.AssessorRevise{reviseAppendItem})
	assert.NoError(t, err)
	modifiedStatus2, err := sm.ReviseStatus(modifiedStatus, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
	assert.NoError(t, err)

	require.Equal(t, len(modifiedStatus2.FollowingAssessors.Assessors), 2)
	_, modifiedValid1 := modifiedStatus2.FollowingAssessors.ObtainViaLocation(assessor1keyToken.Location())
	_, appendedValid2 := modifiedStatus2.FollowingAssessors.ObtainViaLocation(item2publicToken.Location())

	//
	//
	//
	//
	desireItem1urgency := int64(0)
	sumPotencySubsequent := item1ballotingPotency + item2ballotingPotency
	//
	desireItem2urgency := -(sumPotencySubsequent + (sumPotencySubsequent >> 3))
	//
	//
	avg := big.NewInt(0).Add(big.NewInt(desireItem1urgency), big.NewInt(desireItem2urgency))
	avg.Div(avg, big.NewInt(2))
	desireItem2urgency -= avg.Int64() //
	desireItem1urgency -= avg.Int64() //

	//
	desireItem1urgency += item1ballotingPotency //
	desireItem2urgency += item2ballotingPotency //
	desireItem1urgency -= sumPotencySubsequent //

	assert.Equal(t, desireItem1urgency, modifiedValid1.NominatorUrgency)
	assert.Equal(t, desireItem2urgency, appendedValid2.NominatorUrgency)

	//
	//
	modifiedBallotingExponentValid2 := int64(1)
	reviseItem := iface.AssessorRevise{PublicToken: fvp, Potency: modifiedBallotingExponentValid2}
	assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions([]iface.AssessorRevise{reviseItem})
	assert.NoError(t, err)

	//
	//
	modifiedStatus3, err := sm.ReviseStatus(modifiedStatus2, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
	assert.NoError(t, err)

	require.Equal(t, len(modifiedStatus3.FollowingAssessors.Assessors), 2)
	_, previousValid1 := modifiedStatus3.Assessors.ObtainViaLocation(assessor1keyToken.Location())
	_, previousValid2 := modifiedStatus3.Assessors.ObtainViaLocation(item2publicToken.Location())
	_, modifiedValid1 = modifiedStatus3.FollowingAssessors.ObtainViaLocation(assessor1keyToken.Location())
	_, modifiedValid2 := modifiedStatus3.FollowingAssessors.ObtainViaLocation(item2publicToken.Location())

	//
	//
	desireItem1urgency = previousValid1.NominatorUrgency
	desireItem2urgency = previousValid2.NominatorUrgency
	//
	//
	sumPotency := modifiedValid1.BallotingPotency + modifiedValid2.BallotingPotency
	spread := desireItem2urgency - desireItem1urgency
	//
	proportion := (spread + 2*sumPotency - 1) / (2 * sumPotency)
	//
	desireItem1urgency /= proportion //
	desireItem2urgency /= proportion //

	//
	//
	//
	//
	desireItem2urgency += modifiedValid2.BallotingPotency //
	desireItem1urgency += modifiedValid1.BallotingPotency //
	desireItem2urgency -= sumPotency              //

	assert.Equal(t, desireItem2urgency, modifiedValid2.NominatorUrgency)
	assert.Equal(t, desireItem1urgency, modifiedValid1.NominatorUrgency)
}

func VerifyNominatorUrgencyNominatorOptions(t *testing.T) {
	//
	//
	//
	//
	dismantleDepressed, _, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)
	item1ballotingPotency := int64(10)
	assessor1keyToken := edwards25519.ProducePrivateToken().PublicToken()
	valid1 := &kinds.Assessor{Location: assessor1keyToken.Location(), PublicToken: assessor1keyToken, BallotingPotency: item1ballotingPotency}

	//
	status.Assessors = kinds.FreshAssessorAssign([]*kinds.Assessor{valid1})
	status.FollowingAssessors = status.Assessors
	//
	assert.Equal(t, assessor1keyToken.Location(), status.Assessors.Nominator.Location)

	ledger, err := createLedger(status, status.FinalLedgerAltitude+1, new(kinds.Endorse))
	assert.NoError(t, err)
	bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
	require.NoError(t, err)
	ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}
	//
	ifaceReplies := &iface.ReplyCulminateLedger{}
	assessorRevisions, err := kinds.Buffer2temp.AssessorRevisions(ifaceReplies.AssessorRevisions)
	require.NoError(t, err)

	modifiedStatus, err := sm.ReviseStatus(status, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
	assert.NoError(t, err)

	//
	sumPotency := item1ballotingPotency
	desireItem1urgency := 0 + item1ballotingPotency - sumPotency
	assert.Equal(t, desireItem1urgency, modifiedStatus.FollowingAssessors.Assessors[0].NominatorUrgency)
	assert.Equal(t, assessor1keyToken.Location(), modifiedStatus.FollowingAssessors.Nominator.Location)

	//
	item2publicToken := edwards25519.ProducePrivateToken().PublicToken()
	fvp, err := cryptocode.PublicTokenTowardSchema(item2publicToken)
	require.NoError(t, err)
	reviseAppendItem := iface.AssessorRevise{PublicToken: fvp, Potency: item1ballotingPotency}
	assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions([]iface.AssessorRevise{reviseAppendItem})
	assert.NoError(t, err)

	modifiedStatus2, err := sm.ReviseStatus(modifiedStatus, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
	assert.NoError(t, err)

	require.Equal(t, len(modifiedStatus2.FollowingAssessors.Assessors), 2)
	assert.Equal(t, modifiedStatus2.Assessors, modifiedStatus.FollowingAssessors)

	//
	assert.Equal(t, assessor1keyToken.Location(), modifiedStatus.FollowingAssessors.Nominator.Location)
	assert.Equal(t, modifiedStatus2.Assessors.Nominator.Location, modifiedStatus2.FollowingAssessors.Nominator.Location)
	assert.Equal(t, modifiedStatus2.Assessors.Nominator.Location, assessor1keyToken.Location())
	assert.Equal(t, modifiedStatus2.FollowingAssessors.Nominator.Location, assessor1keyToken.Location())

	_, modifiedValid1 := modifiedStatus2.FollowingAssessors.ObtainViaLocation(assessor1keyToken.Location())
	_, agedValid1 := modifiedStatus2.Assessors.ObtainViaLocation(assessor1keyToken.Location())
	_, modifiedValid2 := modifiedStatus2.FollowingAssessors.ObtainViaLocation(item2publicToken.Location())

	//
	item2ballotingPotency := item1ballotingPotency
	sumPotency = item1ballotingPotency + item2ballotingPotency           //
	ver2urgencyWheneverAppendedValid2 := -(sumPotency + (sumPotency >> 3)) //
	//
	//
	medianTotal := big.NewInt(0).Add(big.NewInt(ver2urgencyWheneverAppendedValid2), big.NewInt(agedValid1.NominatorUrgency))
	avg := medianTotal.Div(medianTotal, big.NewInt(2))                   //
	anticipatedItem2urgency := ver2urgencyWheneverAppendedValid2 - avg.Int64()      //
	anticipatedItem1urgency := agedValid1.NominatorUrgency - avg.Int64() //
	//
	anticipatedItem2urgency += item2ballotingPotency //
	anticipatedItem1urgency += item1ballotingPotency //
	anticipatedItem1urgency -= sumPotency      //

	assert.EqualValues(t, anticipatedItem1urgency, modifiedValid1.NominatorUrgency)
	assert.EqualValues(
		t,
		anticipatedItem2urgency,
		modifiedValid2.NominatorUrgency,
		"REDACTED",
		modifiedValid2,
	)

	assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions(ifaceReplies.AssessorRevisions)
	require.NoError(t, err)

	modifiedStatus3, err := sm.ReviseStatus(modifiedStatus2, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
	assert.NoError(t, err)

	assert.Equal(t, modifiedStatus3.Assessors.Nominator.Location, modifiedStatus3.FollowingAssessors.Nominator.Location)

	assert.Equal(t, modifiedStatus3.Assessors, modifiedStatus2.FollowingAssessors)
	_, modifiedValid1 = modifiedStatus3.FollowingAssessors.ObtainViaLocation(assessor1keyToken.Location())
	_, modifiedValid2 = modifiedStatus3.FollowingAssessors.ObtainViaLocation(item2publicToken.Location())

	//
	assert.Equal(t, assessor1keyToken.Location(), modifiedStatus3.FollowingAssessors.Nominator.Location)

	//
	//
	anticipatedItem2urgency2 := anticipatedItem2urgency + item2ballotingPotency //
	anticipatedItem1urgency2 := anticipatedItem1urgency + item1ballotingPotency //
	anticipatedItem1urgency2 -= sumPotency                         //

	assert.EqualValues(
		t,
		anticipatedItem1urgency2,
		modifiedValid1.NominatorUrgency,
		"REDACTED",
		modifiedValid2,
	)
	assert.EqualValues(
		t,
		anticipatedItem2urgency2,
		modifiedValid2.NominatorUrgency,
		"REDACTED",
		modifiedValid2,
	)

	//
	//
	agedStatus := modifiedStatus3
	ifaceReplies = &iface.ReplyCulminateLedger{}
	assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions(ifaceReplies.AssessorRevisions)
	require.NoError(t, err)

	agedStatus, err = sm.ReviseStatus(agedStatus, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
	assert.NoError(t, err)
	anticipatedItem1urgency2 = 1
	anticipatedItem2urgency2 = -1
	anticipatedItem1urgency = -9
	anticipatedItem2urgency = 9

	for i := 0; i < 1000; i++ {
		//
		ifaceReplies := &iface.ReplyCulminateLedger{}
		assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions(ifaceReplies.AssessorRevisions)
		require.NoError(t, err)

		modifiedStatus, err := sm.ReviseStatus(agedStatus, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
		assert.NoError(t, err)
		//
		assert.NotEqual(
			t,
			modifiedStatus.Assessors.Nominator.Location,
			modifiedStatus.FollowingAssessors.Nominator.Location,
			"REDACTED",
			i,
		)
		assert.Equal(t, agedStatus.Assessors.Nominator.Location, modifiedStatus.FollowingAssessors.Nominator.Location, "REDACTED", i)

		_, modifiedValid1 = modifiedStatus.FollowingAssessors.ObtainViaLocation(assessor1keyToken.Location())
		_, modifiedValid2 = modifiedStatus.FollowingAssessors.ObtainViaLocation(item2publicToken.Location())

		if i%2 == 0 {
			assert.Equal(t, modifiedStatus.Assessors.Nominator.Location, item2publicToken.Location())
			assert.Equal(t, anticipatedItem1urgency, modifiedValid1.NominatorUrgency) //
			assert.Equal(t, anticipatedItem2urgency, modifiedValid2.NominatorUrgency) //
		} else {
			assert.Equal(t, modifiedStatus.Assessors.Nominator.Location, assessor1keyToken.Location())
			assert.Equal(t, anticipatedItem1urgency2, modifiedValid1.NominatorUrgency) //
			assert.Equal(t, anticipatedItem2urgency2, modifiedValid2.NominatorUrgency) //
		}
		//
		agedStatus = modifiedStatus
	}
}

func VerifyAmpleInaugurationAssessor(t *testing.T) {
	dismantleDepressed, _, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)

	inaugurationBallotingPotency := kinds.MaximumSumBallotingPotency / 1000
	inaugurationPublicToken := edwards25519.ProducePrivateToken().PublicToken()
	//
	inaugurationItem := &kinds.Assessor{
		Location:     inaugurationPublicToken.Location(),
		PublicToken:      inaugurationPublicToken,
		BallotingPotency: inaugurationBallotingPotency,
	}
	//
	status.Assessors = kinds.FreshAssessorAssign([]*kinds.Assessor{inaugurationItem})
	status.FollowingAssessors = status.Assessors
	require.True(t, len(status.Assessors.Assessors) == 1)

	//
	//
	agedStatus := status
	for i := 0; i < 10; i++ {
		//
		ifaceReplies := &iface.ReplyCulminateLedger{}
		assessorRevisions, err := kinds.Buffer2temp.AssessorRevisions(ifaceReplies.AssessorRevisions)
		require.NoError(t, err)

		ledger, err := createLedger(agedStatus, agedStatus.FinalLedgerAltitude+1, new(kinds.Endorse))
		require.NoError(t, err)
		bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
		require.NoError(t, err)
		ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}

		modifiedStatus, err := sm.ReviseStatus(agedStatus, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
		require.NoError(t, err)
		//
		//
		//
		assert.EqualValues(t, agedStatus.FollowingAssessors, modifiedStatus.FollowingAssessors)
		assert.EqualValues(t, 0, modifiedStatus.FollowingAssessors.Nominator.NominatorUrgency)

		agedStatus = modifiedStatus
	}
	//
	//
	//
	//
	//
	initialAppendedItemPublicToken := edwards25519.ProducePrivateToken().PublicToken()
	initialAppendedItemBallotingPotency := int64(10)
	fvp, err := cryptocode.PublicTokenTowardSchema(initialAppendedItemPublicToken)
	require.NoError(t, err)
	initialAppendedItem := iface.AssessorRevise{PublicToken: fvp, Potency: initialAppendedItemBallotingPotency}
	assessorRevisions, err := kinds.Buffer2temp.AssessorRevisions([]iface.AssessorRevise{initialAppendedItem})
	assert.NoError(t, err)
	ifaceReplies := &iface.ReplyCulminateLedger{
		AssessorRevisions: []iface.AssessorRevise{initialAppendedItem},
	}
	ledger, err := createLedger(agedStatus, agedStatus.FinalLedgerAltitude+1, new(kinds.Endorse))
	require.NoError(t, err)

	bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
	require.NoError(t, err)

	ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}
	modifiedStatus, err := sm.ReviseStatus(agedStatus, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
	require.NoError(t, err)

	finalStatus := modifiedStatus
	for i := 0; i < 200; i++ {
		//
		ifaceReplies := &iface.ReplyCulminateLedger{}
		assessorRevisions, err := kinds.Buffer2temp.AssessorRevisions(ifaceReplies.AssessorRevisions)
		require.NoError(t, err)

		ledger, err := createLedger(finalStatus, finalStatus.FinalLedgerAltitude+1, new(kinds.Endorse))
		require.NoError(t, err)

		bps, err = ledger.CreateFragmentAssign(verifyFragmentExtent)
		require.NoError(t, err)

		ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}

		modifiedStatusInternal, err := sm.ReviseStatus(finalStatus, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
		require.NoError(t, err)
		finalStatus = modifiedStatusInternal
	}
	//
	status = finalStatus

	//
	agedStatus = modifiedStatus
	_, agedInaugurationItem := agedStatus.FollowingAssessors.ObtainViaLocation(inaugurationItem.Location)
	_, freshInaugurationItem := status.FollowingAssessors.ObtainViaLocation(inaugurationItem.Location)
	_, appendedAgedItem := agedStatus.FollowingAssessors.ObtainViaLocation(initialAppendedItemPublicToken.Location())
	_, appendedFreshItem := status.FollowingAssessors.ObtainViaLocation(initialAppendedItemPublicToken.Location())
	//
	assert.True(t, agedInaugurationItem.NominatorUrgency > freshInaugurationItem.NominatorUrgency)
	assert.True(t, appendedAgedItem.NominatorUrgency < appendedFreshItem.NominatorUrgency)

	//
	for i := 0; i < 10; i++ {
		appendedPublicToken := edwards25519.ProducePrivateToken().PublicToken()
		ap, err := cryptocode.PublicTokenTowardSchema(appendedPublicToken)
		require.NoError(t, err)
		appendedItem := iface.AssessorRevise{PublicToken: ap, Potency: initialAppendedItemBallotingPotency}
		assessorRevisions, err := kinds.Buffer2temp.AssessorRevisions([]iface.AssessorRevise{appendedItem})
		assert.NoError(t, err)

		ifaceReplies := &iface.ReplyCulminateLedger{
			AssessorRevisions: []iface.AssessorRevise{appendedItem},
		}
		ledger, err := createLedger(agedStatus, agedStatus.FinalLedgerAltitude+1, new(kinds.Endorse))
		require.NoError(t, err)
		bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
		require.NoError(t, err)

		ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}
		status, err = sm.ReviseStatus(status, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
		require.NoError(t, err)
	}
	require.Equal(t, 10+2, len(status.FollowingAssessors.Assessors))

	//
	gp, err := cryptocode.PublicTokenTowardSchema(inaugurationPublicToken)
	require.NoError(t, err)
	discardInaugurationItem := iface.AssessorRevise{PublicToken: gp, Potency: 0}
	ifaceReplies = &iface.ReplyCulminateLedger{
		AssessorRevisions: []iface.AssessorRevise{discardInaugurationItem},
	}

	ledger, err = createLedger(agedStatus, agedStatus.FinalLedgerAltitude+1, new(kinds.Endorse))
	require.NoError(t, err)

	bps, err = ledger.CreateFragmentAssign(verifyFragmentExtent)
	require.NoError(t, err)

	ledgerUUID = kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}
	assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions(ifaceReplies.AssessorRevisions)
	require.NoError(t, err)
	modifiedStatus, err = sm.ReviseStatus(status, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
	require.NoError(t, err)
	//
	assert.Equal(t, 11, len(modifiedStatus.FollowingAssessors.Assessors))

	//
	//
	presentStatus := modifiedStatus
	tally := 0
	equalsNominatorUnaltered := true
	for equalsNominatorUnaltered {
		ifaceReplies := &iface.ReplyCulminateLedger{}
		assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions(ifaceReplies.AssessorRevisions)
		require.NoError(t, err)
		ledger, err = createLedger(presentStatus, presentStatus.FinalLedgerAltitude+1, new(kinds.Endorse))
		require.NoError(t, err)

		bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
		require.NoError(t, err)

		ledgerUUID = kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}
		presentStatus, err = sm.ReviseStatus(presentStatus, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
		require.NoError(t, err)
		if !bytes.Equal(presentStatus.Assessors.Nominator.Location, presentStatus.FollowingAssessors.Nominator.Location) {
			equalsNominatorUnaltered = false
		}
		tally++
	}
	modifiedStatus = presentStatus
	//
	initialNominatorAlterationAnticipatedSubsequent := 1
	assert.Equal(t, initialNominatorAlterationAnticipatedSubsequent, tally)
	//
	countValues := len(modifiedStatus.Assessors.Assessors)
	nominators := make([]*kinds.Assessor, countValues)
	for i := 0; i < 100; i++ {
		//
		ifaceReplies := &iface.ReplyCulminateLedger{}
		assessorRevisions, err := kinds.Buffer2temp.AssessorRevisions(ifaceReplies.AssessorRevisions)
		require.NoError(t, err)

		ledger, err := createLedger(modifiedStatus, modifiedStatus.FinalLedgerAltitude+1, new(kinds.Endorse))
		require.NoError(t, err)

		bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
		require.NoError(t, err)

		ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}

		modifiedStatus, err = sm.ReviseStatus(modifiedStatus, ledgerUUID, &ledger.Heading, ifaceReplies, assessorRevisions)
		require.NoError(t, err)
		if i > countValues { //
			if nominators[i%countValues] == nil {
				nominators[i%countValues] = modifiedStatus.FollowingAssessors.Nominator
			} else {
				assert.Equal(t, nominators[i%countValues], modifiedStatus.FollowingAssessors.Nominator)
			}
		}
	}
}

func VerifyDepotFetchAssessorsAdditionsNominatorUrgency(t *testing.T) {
	const itemAssignExtent = 2
	dismantleDepressed, statusDatastore, status := configureVerifyInstance(t)
	t.Cleanup(func() { dismantleDepressed(t) })
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	status.Assessors = produceItemAssign(itemAssignExtent)
	status.FollowingAssessors = status.Assessors.DuplicateAdvanceNominatorUrgency(1)
	err := statusDepot.Persist(status)
	require.NoError(t, err)

	followingAltitude := status.FinalLedgerAltitude + 1

	v0, err := statusDepot.FetchAssessors(followingAltitude)
	assert.Nil(t, err)
	accumulator0 := v0.Assessors[0].NominatorUrgency

	v1, err := statusDepot.FetchAssessors(followingAltitude + 1)
	assert.Nil(t, err)
	accumulator1 := v1.Assessors[0].NominatorUrgency

	assert.NotEqual(t, accumulator1, accumulator0, "REDACTED")
}

//
//
func VerifyMultipleAssessorModificationsPersistFetch(t *testing.T) {
	const itemAssignExtent = 7
	dismantleDepressed, statusDatastore, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	require.Equal(t, int64(0), status.FinalLedgerAltitude)
	status.Assessors = produceItemAssign(itemAssignExtent)
	status.FollowingAssessors = status.Assessors.DuplicateAdvanceNominatorUrgency(1)
	err := statusDepot.Persist(status)
	require.NoError(t, err)

	_, itemAged := status.Assessors.ObtainViaOrdinal(0)
	publickeyAged := itemAged.PublicToken
	publickey := edwards25519.ProducePrivateToken().PublicToken()

	//
	heading, ledgerUUID, replies := createHeadlineFragmentsRepliesItemPublicTokenAlteration(status, publickey)

	//
	var assessorRevisions []*kinds.Assessor
	assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions(replies.AssessorRevisions)
	require.NoError(t, err)
	status, err = sm.ReviseStatus(status, ledgerUUID, &heading, replies, assessorRevisions)
	require.Nil(t, err)
	followingAltitude := status.FinalLedgerAltitude + 1
	err = statusDepot.Persist(status)
	require.NoError(t, err)

	//
	v0, err := statusDepot.FetchAssessors(followingAltitude)
	assert.Nil(t, err)
	assert.Equal(t, itemAssignExtent, v0.Extent())
	ordinal, val := v0.ObtainViaLocation(publickeyAged.Location())
	assert.NotNil(t, val)
	if ordinal < 0 {
		t.Fatal("REDACTED")
	}

	//
	v1, err := statusDepot.FetchAssessors(followingAltitude + 1)
	assert.Nil(t, err)
	assert.Equal(t, itemAssignExtent, v1.Extent())
	ordinal, val = v1.ObtainViaLocation(publickey.Location())
	assert.NotNil(t, val)
	if ordinal < 0 {
		t.Fatal("REDACTED")
	}
}

func VerifyStatusCreateLedger(t *testing.T) {
	dismantleDepressed, _, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)

	nominatorLocator := status.Assessors.ObtainNominator().Location
	statusEdition := status.Edition.Agreement
	ledger, err := createLedger(status, 2, new(kinds.Endorse))
	require.NoError(t, err)

	//
	assert.Equal(t, statusEdition, ledger.Edition)
	assert.Equal(t, nominatorLocator, ledger.NominatorLocation)
}

//
//
func VerifyAgreementParametersModificationsPersistFetch(t *testing.T) {
	dismantleDepressed, statusDatastore, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)

	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	//
	alterationElevations := []int64{1, 2, 4, 5, 10, 15, 16, 17, 20}
	N := len(alterationElevations)

	//
	//
	parameters := make([]kinds.AgreementSettings, N+1)
	parameters[0] = status.AgreementSettings
	for i := 1; i < N+1; i++ {
		parameters[i] = *kinds.FallbackAgreementSettings()
		parameters[i].Ledger.MaximumOctets += int64(i)

	}

	//
	//
	utmostAltitude := alterationElevations[N-1] + 5
	alterationPosition := 0
	cp := parameters[alterationPosition]
	var err error
	var assessorRevisions []*kinds.Assessor
	for i := int64(1); i < utmostAltitude; i++ {
		//
		if alterationPosition < len(alterationElevations) && i == alterationElevations[alterationPosition] {
			alterationPosition++
			cp = parameters[alterationPosition]
		}
		heading, ledgerUUID, replies := createHeadlineFragmentsRepliesParameters(status, cp.TowardSchema())
		assessorRevisions, err = kinds.Buffer2temp.AssessorRevisions(replies.AssessorRevisions)
		require.NoError(t, err)
		status, err = sm.ReviseStatus(status, ledgerUUID, &heading, replies, assessorRevisions)

		require.NoError(t, err)
		err = statusDepot.Persist(status)
		require.NoError(t, err)
	}

	//
	verifyScenarios := make([]parametersAlterationVerifyInstance, utmostAltitude)
	alterationPosition = 0
	cp = parameters[alterationPosition]
	for i := int64(1); i < utmostAltitude+1; i++ {
		//
		//
		if alterationPosition < len(alterationElevations) && i == alterationElevations[alterationPosition]+1 {
			alterationPosition++
			cp = parameters[alterationPosition]
		}
		verifyScenarios[i-1] = parametersAlterationVerifyInstance{i, cp}
	}

	for _, verifyInstance := range verifyScenarios {
		p, err := statusDepot.FetchAgreementParameters(verifyInstance.altitude)
		assert.Nil(t, err, fmt.Sprintf("REDACTED", verifyInstance.altitude))
		assert.EqualValues(t, verifyInstance.parameters, p, fmt.Sprintf(`REDACTEDt
REDACTED`, verifyInstance.altitude))
	}
}

func VerifyStatusSchema(t *testing.T) {
	dismantleDepressed, _, status := configureVerifyInstance(t)
	defer dismantleDepressed(t)

	tc := []struct {
		verifyAlias string
		status    *sm.Status
		expirationPhase1 bool
		expirationPhase2 bool
	}{
		{"REDACTED", &sm.Status{}, true, false},
		{"REDACTED", nil, false, false},
		{"REDACTED", &status, true, true},
	}

	for _, tt := range tc {

		pbs, err := tt.status.TowardSchema()
		if !tt.expirationPhase1 {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err, tt.verifyAlias)
		}

		smt, err := sm.OriginatingSchema(pbs)
		if tt.expirationPhase2 {
			require.NoError(t, err, tt.verifyAlias)
			require.Equal(t, tt.status, smt, tt.verifyAlias)
		} else {
			require.Error(t, err, tt.verifyAlias)
		}
	}
}

func VerifyAverageMoment(t *testing.T) {
	valid1 := kinds.FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 30)
	valid2 := kinds.FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 30)
	item3 := kinds.FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 30)

	values := kinds.FreshAssessorAssign([]*kinds.Assessor{valid1, valid2, item3})

	t.Run("REDACTED", func(t *testing.T) {
		now := time.Now()
		endorse := &kinds.Endorse{
			Altitude: 1,
			Notations: []kinds.EndorseSignature{
				{
					LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
					AssessorLocation: valid1.Location,
					Timestamp:        now,
				},
				{
					LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
					AssessorLocation: valid2.Location,
					Timestamp:        now.Add(1 * time.Minute),
				},
				{
					LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
					AssessorLocation: item3.Location,
					Timestamp:        now.Add(2 * time.Minute),
				},
			},
		}

		averageMoment, err := sm.AverageMoment(endorse, values)
		require.NoError(t, err)
		require.Equal(t, averageMoment, now.Add(1*time.Minute))
	})

	t.Run("REDACTED", func(t *testing.T) {
		unfamiliarItem := edwards25519.ProducePrivateToken().PublicToken().Location()
		now := time.Now()
		endorse := &kinds.Endorse{
			Altitude: 1,
			Notations: []kinds.EndorseSignature{
				{
					LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
					AssessorLocation: valid1.Location,
					Timestamp:        now,
				},
				{
					LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
					AssessorLocation: unfamiliarItem,
					Timestamp:        now.Add(1 * time.Minute),
				},
			},
		}

		_, err := sm.AverageMoment(endorse, values)
		require.Error(t, err)
		require.Contains(t, err.Error(), "REDACTED")
	})

	t.Run("REDACTED", func(t *testing.T) {
		now := time.Now()
		endorse := &kinds.Endorse{
			Altitude: 1,
			Notations: []kinds.EndorseSignature{
				{
					LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
					AssessorLocation: valid1.Location,
					Timestamp:        now,
				},
				{
					LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
					AssessorLocation: valid2.Location,
					Timestamp:        now.Add(1 * time.Minute),
				},
			},
		}

		averageMoment, err := sm.AverageMoment(endorse, values)
		require.NoError(t, err)
		require.Equal(t, averageMoment, now)
	})
}
