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

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault/ed25519"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/intrinsic/verify"
	engineseed "github.com/valkyrieworks/utils/random"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

//
func configureVerifyScenario(t *testing.T) (func(t *testing.T), dbm.DB, sm.Status) {
	t.Helper()
	dismantleBelow, statusStore, status, _ := configureVerifyScenarioWithDepot(t)
	return dismantleBelow, statusStore, status
}

//
func configureVerifyScenarioWithDepot(t *testing.T) (func(t *testing.T), dbm.DB, sm.Status, sm.Depot) {
	t.Helper()
	settings := verify.RestoreVerifyOrigin("REDACTED")
	storeKind := dbm.OriginKind(settings.StoreOrigin)
	statusStore, err := dbm.NewStore("REDACTED", storeKind, settings.StoreFolder())
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	require.NoError(t, err)
	status, err := statusDepot.ImportFromStoreOrOriginEntry(settings.OriginEntry())
	require.NoError(t, err, "REDACTED")
	err = statusDepot.Persist(status)
	require.NoError(t, err)

	dismantleBelow := func(t *testing.T) {
		t.Helper()
		os.RemoveAll(settings.OriginFolder)
	}

	return dismantleBelow, statusStore, status, statusDepot
}

//
func VerifyStatusClone(t *testing.T) {
	dismantleBelow, _, status := configureVerifyScenario(t)
	defer dismantleBelow(t)
	affirm := assert.New(t)

	statusClone := status.Clone()

	assert.True(status.Matches(statusClone),
		fmt.Sprintf("REDACTED",
			statusClone, status))

	statusClone.FinalLedgerLevel++
	statusClone.FinalRatifiers = status.Ratifiers
	assert.False(status.Matches(statusClone), fmt.Sprintf(`REDACTEDe
REDACTED`, status))
}

//
func VerifyCreateOriginStatusNullRatifiers(t *testing.T) {
	doc := kinds.OriginPaper{
		LedgerUID:    "REDACTED",
		Ratifiers: nil,
	}
	require.Nil(t, doc.CertifyAndFinished())
	status, err := sm.CreateOriginStatus(&doc)
	require.Nil(t, err)
	require.Equal(t, 0, len(status.Ratifiers.Ratifiers))
	require.Equal(t, 0, len(status.FollowingRatifiers.Ratifiers))
}

//
func VerifyStatusPersistImport(t *testing.T) {
	dismantleBelow, statusStore, status := configureVerifyScenario(t)
	defer dismantleBelow(t)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	affirm := assert.New(t)

	status.FinalLedgerLevel++
	status.FinalRatifiers = status.Ratifiers
	err := statusDepot.Persist(status)
	require.NoError(t, err)

	retrievedStatus, err := statusDepot.Import()
	require.NoError(t, err)
	assert.True(status.Matches(retrievedStatus),
		fmt.Sprintf("REDACTED",
			retrievedStatus, status))
}

//
func VerifyCompleteLedgerRepliesPersistFetch1(t *testing.T) {
	dismantleBelow, statusStore, status := configureVerifyScenario(t)
	defer dismantleBelow(t)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	affirm := assert.New(t)

	status.FinalLedgerLevel++

	//
	ledger, err := createLedger(status, 2, new(kinds.Endorse))
	require.NoError(t, err)

	ifaceReplies := new(iface.ReplyCompleteLedger)
	dtrans := make([]*iface.InvokeTransferOutcome, 2)
	ifaceReplies.TransOutcomes = dtrans

	ifaceReplies.TransOutcomes[0] = &iface.InvokeTransferOutcome{Data: []byte("REDACTED"), Events: nil}
	ifaceReplies.TransOutcomes[1] = &iface.InvokeTransferOutcome{Data: []byte("REDACTED"), Log: "REDACTED", Events: nil}
	ifaceReplies.RatifierRefreshes = []iface.RatifierModify{
		kinds.Tm2schema.NewRatifierModify(ed25519.GeneratePrivateKey().PublicKey(), 10),
	}

	ifaceReplies.ApplicationDigest = make([]byte, 1)

	err = statusDepot.PersistCompleteLedgerReply(ledger.Level, ifaceReplies)
	require.NoError(t, err)
	retrievedIfaceReplies, err := statusDepot.ImportCompleteLedgerReply(ledger.Level)
	assert.NoError(err)
	assert.Equal(ifaceReplies, retrievedIfaceReplies)
}

//
func VerifyCompleteLedgerRepliesPersistFetch2(t *testing.T) {
	dismantleBelow, statusStore, _ := configureVerifyScenario(t)
	defer dismantleBelow(t)
	affirm := assert.New(t)

	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
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
				{Code: 32, Data: []byte("REDACTED"), Log: "REDACTED"},
			},
			[]*iface.InvokeTransferOutcome{
				{Code: 32, Data: []byte("REDACTED")},
			},
		},
		2: {
			[]*iface.InvokeTransferOutcome{
				{Code: 383},
				{
					Data: []byte("REDACTED"),
					Events: []iface.Event{
						{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED"}}},
						{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED"}}},
					},
				},
			},
			[]*iface.InvokeTransferOutcome{
				{Code: 383, Data: nil},
				{Code: 0, Data: []byte("REDACTED"), Events: []iface.Event{
					{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED"}}},
					{Kind: "REDACTED", Properties: []iface.EventProperty{{Key: "REDACTED", Item: "REDACTED"}}},
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
		res, err := statusDepot.ImportCompleteLedgerReply(h)
		assert.Error(err, "REDACTED", i, res)
	}

	//
	for i, tc := range scenarios {
		h := int64(i + 1) //
		replies := &iface.ReplyCompleteLedger{
			TransOutcomes: tc.appended,
			ApplicationDigest:   []byte(fmt.Sprintf("REDACTED", h)),
		}
		err := statusDepot.PersistCompleteLedgerReply(h, replies)
		require.NoError(t, err)
	}

	//
	for i, tc := range scenarios {
		h := int64(i + 1)
		res, err := statusDepot.ImportCompleteLedgerReply(h)
		if assert.NoError(err, "REDACTED", i) {
			t.Log(res)
			replies := &iface.ReplyCompleteLedger{
				TransOutcomes: tc.anticipated,
				ApplicationDigest:   []byte(fmt.Sprintf("REDACTED", h)),
			}
			assert.Equal(sm.TransferOutcomesDigest(replies.TransOutcomes), sm.TransferOutcomesDigest(res.TransOutcomes), "REDACTED", i)
		}
	}
}

//
func VerifyRatifierBasicPersistImport(t *testing.T) {
	dismantleBelow, statusStore, status := configureVerifyScenario(t)
	defer dismantleBelow(t)
	affirm := assert.New(t)

	statusdepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	//
	_, err := statusdepot.ImportRatifiers(0)
	assert.IsType(sm.ErrNoValueCollectionForLevel{}, err, "REDACTED")

	//
	v, err := statusdepot.ImportRatifiers(1)
	assert.Nil(err, "REDACTED")
	assert.Equal(v.Digest(), status.Ratifiers.Digest(), "REDACTED")

	//
	v, err = statusdepot.ImportRatifiers(2)
	assert.Nil(err, "REDACTED")
	assert.Equal(v.Digest(), status.FollowingRatifiers.Digest(), "REDACTED")

	//
	status.FinalLedgerLevel++
	followingLevel := status.FinalLedgerLevel + 1
	err = statusdepot.Persist(status)
	require.NoError(t, err)
	vp0, err := statusdepot.ImportRatifiers(followingLevel + 0)
	assert.Nil(err, "REDACTED")
	vp1, err := statusdepot.ImportRatifiers(followingLevel + 1)
	assert.Nil(err, "REDACTED")
	assert.Equal(vp0.Digest(), status.Ratifiers.Digest(), "REDACTED")
	assert.Equal(vp1.Digest(), status.FollowingRatifiers.Digest(), "REDACTED")
}

//
func VerifyOneRatifierModificationsPersistImport(t *testing.T) {
	dismantleBelow, statusStore, status := configureVerifyScenario(t)
	defer dismantleBelow(t)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	//
	alterLevels := []int64{1, 2, 4, 5, 10, 15, 16, 17, 20}
	N := len(alterLevels)

	//
	//
	maximumLevel := alterLevels[N-1] + 5
	alterOrdinal := 0
	_, val := status.Ratifiers.FetchByOrdinal(0)
	energy := val.PollingEnergy
	var err error
	var ratifierRefreshes []*kinds.Ratifier
	for i := int64(1); i < maximumLevel; i++ {
		//
		if alterOrdinal < len(alterLevels) && i == alterLevels[alterOrdinal] {
			alterOrdinal++
			energy++
		}
		heading, ledgerUID, replies := createHeadingSectionsRepliesValueEnergyAlter(status, energy)
		ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes(replies.RatifierRefreshes)
		require.NoError(t, err)
		status, err = sm.ModifyStatus(status, ledgerUID, &heading, replies, ratifierRefreshes)
		require.NoError(t, err)
		err = statusDepot.Persist(status)
		require.NoError(t, err)
	}

	//
	verifyScenarios := make([]int64, maximumLevel)
	alterOrdinal = 0
	energy = val.PollingEnergy
	for i := int64(1); i < maximumLevel+1; i++ {
		//
		//
		if alterOrdinal < len(alterLevels) && i == alterLevels[alterOrdinal]+1 {
			alterOrdinal++
			energy++
		}
		verifyScenarios[i-1] = energy
	}

	for i, energy := range verifyScenarios {
		v, err := statusDepot.ImportRatifiers(int64(i + 1 + 1)) //
		assert.Nil(t, err, fmt.Sprintf("REDACTED", i))
		assert.Equal(t, v.Volume(), 1, "REDACTED", v.Volume())
		_, val := v.FetchByOrdinal(0)

		assert.Equal(t, val.PollingEnergy, energy, fmt.Sprintf(`REDACTEDt
REDACTED`, i))
	}
}

func VerifyRecommenderRecurrence(t *testing.T) {
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

	for scenarioCount, verifyInstance := range verifyScenarios {
		//
		//
		for i := 0; i < 5; i++ {
			valueCollection := generateValueCollectionWithStrengths(verifyInstance.strengths)
			verifyRecommenderCount(t, scenarioCount, valueCollection)
		}
	}

	//
	maximumValues := 100
	maximumEnergy := 1000
	nVerifyScenarios := 5
	for i := 0; i < nVerifyScenarios; i++ {
		N := engineseed.Int()%maximumValues + 1
		values := make([]*kinds.Ratifier, N)
		sumBallotEnergy := int64(0)
		for j := 0; j < N; j++ {
			//
			ballotEnergy := int64(engineseed.Int()%maximumEnergy) + 1
			sumBallotEnergy += ballotEnergy
			privateValue := kinds.NewEmulatePV()
			publicKey, err := privateValue.FetchPublicKey()
			require.NoError(t, err)
			val := kinds.NewRatifier(publicKey, ballotEnergy)
			val.RecommenderUrgency = engineseed.Int64()
			values[j] = val
		}
		valueCollection := kinds.NewRatifierCollection(values)
		valueCollection.ReadjustUrgencies(sumBallotEnergy)
		verifyRecommenderCount(t, i, valueCollection)
	}
}

//
func generateValueCollectionWithStrengths(strengths []int64) *kinds.RatifierAssign {
	volume := len(strengths)
	values := make([]*kinds.Ratifier, volume)
	sumBallotEnergy := int64(0)
	for i := 0; i < volume; i++ {
		sumBallotEnergy += strengths[i]
		val := kinds.NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), strengths[i])
		val.RecommenderUrgency = engineseed.Int64()
		values[i] = val
	}
	valueCollection := kinds.NewRatifierCollection(values)
	valueCollection.ReadjustUrgencies(sumBallotEnergy)
	return valueCollection
}

//
func verifyRecommenderCount(t *testing.T, scenarioCount int, valueCollection *kinds.RatifierAssign) {
	N := valueCollection.Volume()
	sumEnergy := valueCollection.SumPollingEnergy()

	//
	runMultiply := 1
	cycles := int(sumEnergy) * runMultiply
	counts := make([]int, N)
	for i := 0; i < cycles; i++ {
		nomination := valueCollection.FetchRecommender()
		idx, _ := valueCollection.FetchByLocation(nomination.Location)
		counts[idx]++
		valueCollection.AugmentRecommenderUrgency(1)
	}

	//
	for i, count := range counts {
		_, val := valueCollection.FetchByOrdinal(int32(i))
		anticipateCount := int(val.PollingEnergy) * runMultiply
		acquiredCount := count
		abs := int(math.Abs(float64(anticipateCount - acquiredCount)))

		//
		//
		//
		//
		limited := N - 1
		require.True(
			t,
			abs <= limited,
			fmt.Sprintf("REDACTED", scenarioCount, i, N, acquiredCount, anticipateCount),
		)
	}
}

//
//
func VerifyRecommenderUrgencyDoesNotFetchRestoreToNil(t *testing.T) {
	dismantleBelow, _, status := configureVerifyScenario(t)
	defer dismantleBelow(t)
	value1votingEnergy := int64(10)
	val1publicKey := ed25519.GeneratePrivateKey().PublicKey()
	value1 := &kinds.Ratifier{Location: val1publicKey.Location(), PublicKey: val1publicKey, PollingEnergy: value1votingEnergy}

	status.Ratifiers = kinds.NewRatifierCollection([]*kinds.Ratifier{value1})
	status.FollowingRatifiers = status.Ratifiers

	//
	assert.EqualValues(t, 0, value1.RecommenderUrgency)

	ledger, err := createLedger(status, status.FinalLedgerLevel+1, new(kinds.Endorse))
	require.NoError(t, err)
	bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
	require.NoError(t, err)
	ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}
	ifaceReplies := &iface.ReplyCompleteLedger{}
	ratifierRefreshes, err := kinds.Schema2tm.RatifierRefreshes(ifaceReplies.RatifierRefreshes)
	require.NoError(t, err)
	refreshedStatus, err := sm.ModifyStatus(status, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
	assert.NoError(t, err)
	currentSum := value1votingEnergy
	//
	assert.Equal(t, 0+value1votingEnergy-currentSum, refreshedStatus.FollowingRatifiers.Ratifiers[0].RecommenderUrgency)

	//
	value2publicKey := ed25519.GeneratePrivateKey().PublicKey()
	value2votingEnergy := int64(100)
	fvp, err := cryptocode.PublicKeyToSchema(value2publicKey)
	require.NoError(t, err)

	modifyAppendValue := iface.RatifierModify{PublicKey: fvp, Energy: value2votingEnergy}
	ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes([]iface.RatifierModify{modifyAppendValue})
	assert.NoError(t, err)
	refreshedStatus2, err := sm.ModifyStatus(refreshedStatus, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
	assert.NoError(t, err)

	require.Equal(t, len(refreshedStatus2.FollowingRatifiers.Ratifiers), 2)
	_, refreshedValue1 := refreshedStatus2.FollowingRatifiers.FetchByLocation(val1publicKey.Location())
	_, appendedValue2 := refreshedStatus2.FollowingRatifiers.FetchByLocation(value2publicKey.Location())

	//
	//
	//
	//
	desireValue1priority := int64(0)
	sumEnergyAfter := value1votingEnergy + value2votingEnergy
	//
	desireValue2priority := -(sumEnergyAfter + (sumEnergyAfter >> 3))
	//
	//
	avg := big.NewInt(0).Add(big.NewInt(desireValue1priority), big.NewInt(desireValue2priority))
	avg.Div(avg, big.NewInt(2))
	desireValue2priority -= avg.Int64() //
	desireValue1priority -= avg.Int64() //

	//
	desireValue1priority += value1votingEnergy //
	desireValue2priority += value2votingEnergy //
	desireValue1priority -= sumEnergyAfter //

	assert.Equal(t, desireValue1priority, refreshedValue1.RecommenderUrgency)
	assert.Equal(t, desireValue2priority, appendedValue2.RecommenderUrgency)

	//
	//
	refreshedPollingNonceValue2 := int64(1)
	modifyValue := iface.RatifierModify{PublicKey: fvp, Energy: refreshedPollingNonceValue2}
	ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes([]iface.RatifierModify{modifyValue})
	assert.NoError(t, err)

	//
	//
	refreshedStatus3, err := sm.ModifyStatus(refreshedStatus2, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
	assert.NoError(t, err)

	require.Equal(t, len(refreshedStatus3.FollowingRatifiers.Ratifiers), 2)
	_, previousValue1 := refreshedStatus3.Ratifiers.FetchByLocation(val1publicKey.Location())
	_, previousValue2 := refreshedStatus3.Ratifiers.FetchByLocation(value2publicKey.Location())
	_, refreshedValue1 = refreshedStatus3.FollowingRatifiers.FetchByLocation(val1publicKey.Location())
	_, refreshedValue2 := refreshedStatus3.FollowingRatifiers.FetchByLocation(value2publicKey.Location())

	//
	//
	desireValue1priority = previousValue1.RecommenderUrgency
	desireValue2priority = previousValue2.RecommenderUrgency
	//
	//
	sumEnergy := refreshedValue1.PollingEnergy + refreshedValue2.PollingEnergy
	distance := desireValue2priority - desireValue1priority
	//
	proportion := (distance + 2*sumEnergy - 1) / (2 * sumEnergy)
	//
	desireValue1priority /= proportion //
	desireValue2priority /= proportion //

	//
	//
	//
	//
	desireValue2priority += refreshedValue2.PollingEnergy //
	desireValue1priority += refreshedValue1.PollingEnergy //
	desireValue2priority -= sumEnergy              //

	assert.Equal(t, desireValue2priority, refreshedValue2.RecommenderUrgency)
	assert.Equal(t, desireValue1priority, refreshedValue1.RecommenderUrgency)
}

func VerifyRecommenderUrgencyRecommenderVariants(t *testing.T) {
	//
	//
	//
	//
	dismantleBelow, _, status := configureVerifyScenario(t)
	defer dismantleBelow(t)
	value1votingEnergy := int64(10)
	val1publicKey := ed25519.GeneratePrivateKey().PublicKey()
	value1 := &kinds.Ratifier{Location: val1publicKey.Location(), PublicKey: val1publicKey, PollingEnergy: value1votingEnergy}

	//
	status.Ratifiers = kinds.NewRatifierCollection([]*kinds.Ratifier{value1})
	status.FollowingRatifiers = status.Ratifiers
	//
	assert.Equal(t, val1publicKey.Location(), status.Ratifiers.Recommender.Location)

	ledger, err := createLedger(status, status.FinalLedgerLevel+1, new(kinds.Endorse))
	assert.NoError(t, err)
	bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
	require.NoError(t, err)
	ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}
	//
	ifaceReplies := &iface.ReplyCompleteLedger{}
	ratifierRefreshes, err := kinds.Schema2tm.RatifierRefreshes(ifaceReplies.RatifierRefreshes)
	require.NoError(t, err)

	refreshedStatus, err := sm.ModifyStatus(status, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
	assert.NoError(t, err)

	//
	sumEnergy := value1votingEnergy
	desireValue1priority := 0 + value1votingEnergy - sumEnergy
	assert.Equal(t, desireValue1priority, refreshedStatus.FollowingRatifiers.Ratifiers[0].RecommenderUrgency)
	assert.Equal(t, val1publicKey.Location(), refreshedStatus.FollowingRatifiers.Recommender.Location)

	//
	value2publicKey := ed25519.GeneratePrivateKey().PublicKey()
	fvp, err := cryptocode.PublicKeyToSchema(value2publicKey)
	require.NoError(t, err)
	modifyAppendValue := iface.RatifierModify{PublicKey: fvp, Energy: value1votingEnergy}
	ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes([]iface.RatifierModify{modifyAppendValue})
	assert.NoError(t, err)

	refreshedStatus2, err := sm.ModifyStatus(refreshedStatus, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
	assert.NoError(t, err)

	require.Equal(t, len(refreshedStatus2.FollowingRatifiers.Ratifiers), 2)
	assert.Equal(t, refreshedStatus2.Ratifiers, refreshedStatus.FollowingRatifiers)

	//
	assert.Equal(t, val1publicKey.Location(), refreshedStatus.FollowingRatifiers.Recommender.Location)
	assert.Equal(t, refreshedStatus2.Ratifiers.Recommender.Location, refreshedStatus2.FollowingRatifiers.Recommender.Location)
	assert.Equal(t, refreshedStatus2.Ratifiers.Recommender.Location, val1publicKey.Location())
	assert.Equal(t, refreshedStatus2.FollowingRatifiers.Recommender.Location, val1publicKey.Location())

	_, refreshedValue1 := refreshedStatus2.FollowingRatifiers.FetchByLocation(val1publicKey.Location())
	_, agedValue1 := refreshedStatus2.Ratifiers.FetchByLocation(val1publicKey.Location())
	_, refreshedValue2 := refreshedStatus2.FollowingRatifiers.FetchByLocation(value2publicKey.Location())

	//
	value2votingEnergy := value1votingEnergy
	sumEnergy = value1votingEnergy + value2votingEnergy           //
	value2priorityWhenAppendedValue2 := -(sumEnergy + (sumEnergy >> 3)) //
	//
	//
	averageTotal := big.NewInt(0).Add(big.NewInt(value2priorityWhenAppendedValue2), big.NewInt(agedValue1.RecommenderUrgency))
	avg := averageTotal.Div(averageTotal, big.NewInt(2))                   //
	anticipatedValue2priority := value2priorityWhenAppendedValue2 - avg.Int64()      //
	anticipatedValue1priority := agedValue1.RecommenderUrgency - avg.Int64() //
	//
	anticipatedValue2priority += value2votingEnergy //
	anticipatedValue1priority += value1votingEnergy //
	anticipatedValue1priority -= sumEnergy      //

	assert.EqualValues(t, anticipatedValue1priority, refreshedValue1.RecommenderUrgency)
	assert.EqualValues(
		t,
		anticipatedValue2priority,
		refreshedValue2.RecommenderUrgency,
		"REDACTED",
		refreshedValue2,
	)

	ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes(ifaceReplies.RatifierRefreshes)
	require.NoError(t, err)

	refreshedStatus3, err := sm.ModifyStatus(refreshedStatus2, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
	assert.NoError(t, err)

	assert.Equal(t, refreshedStatus3.Ratifiers.Recommender.Location, refreshedStatus3.FollowingRatifiers.Recommender.Location)

	assert.Equal(t, refreshedStatus3.Ratifiers, refreshedStatus2.FollowingRatifiers)
	_, refreshedValue1 = refreshedStatus3.FollowingRatifiers.FetchByLocation(val1publicKey.Location())
	_, refreshedValue2 = refreshedStatus3.FollowingRatifiers.FetchByLocation(value2publicKey.Location())

	//
	assert.Equal(t, val1publicKey.Location(), refreshedStatus3.FollowingRatifiers.Recommender.Location)

	//
	//
	anticipatedValue2priority2 := anticipatedValue2priority + value2votingEnergy //
	anticipatedValue1priority2 := anticipatedValue1priority + value1votingEnergy //
	anticipatedValue1priority2 -= sumEnergy                         //

	assert.EqualValues(
		t,
		anticipatedValue1priority2,
		refreshedValue1.RecommenderUrgency,
		"REDACTED",
		refreshedValue2,
	)
	assert.EqualValues(
		t,
		anticipatedValue2priority2,
		refreshedValue2.RecommenderUrgency,
		"REDACTED",
		refreshedValue2,
	)

	//
	//
	agedStatus := refreshedStatus3
	ifaceReplies = &iface.ReplyCompleteLedger{}
	ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes(ifaceReplies.RatifierRefreshes)
	require.NoError(t, err)

	agedStatus, err = sm.ModifyStatus(agedStatus, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
	assert.NoError(t, err)
	anticipatedValue1priority2 = 1
	anticipatedValue2priority2 = -1
	anticipatedValue1priority = -9
	anticipatedValue2priority = 9

	for i := 0; i < 1000; i++ {
		//
		ifaceReplies := &iface.ReplyCompleteLedger{}
		ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes(ifaceReplies.RatifierRefreshes)
		require.NoError(t, err)

		refreshedStatus, err := sm.ModifyStatus(agedStatus, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
		assert.NoError(t, err)
		//
		assert.NotEqual(
			t,
			refreshedStatus.Ratifiers.Recommender.Location,
			refreshedStatus.FollowingRatifiers.Recommender.Location,
			"REDACTED",
			i,
		)
		assert.Equal(t, agedStatus.Ratifiers.Recommender.Location, refreshedStatus.FollowingRatifiers.Recommender.Location, "REDACTED", i)

		_, refreshedValue1 = refreshedStatus.FollowingRatifiers.FetchByLocation(val1publicKey.Location())
		_, refreshedValue2 = refreshedStatus.FollowingRatifiers.FetchByLocation(value2publicKey.Location())

		if i%2 == 0 {
			assert.Equal(t, refreshedStatus.Ratifiers.Recommender.Location, value2publicKey.Location())
			assert.Equal(t, anticipatedValue1priority, refreshedValue1.RecommenderUrgency) //
			assert.Equal(t, anticipatedValue2priority, refreshedValue2.RecommenderUrgency) //
		} else {
			assert.Equal(t, refreshedStatus.Ratifiers.Recommender.Location, val1publicKey.Location())
			assert.Equal(t, anticipatedValue1priority2, refreshedValue1.RecommenderUrgency) //
			assert.Equal(t, anticipatedValue2priority2, refreshedValue2.RecommenderUrgency) //
		}
		//
		agedStatus = refreshedStatus
	}
}

func VerifyBulkyOriginRatifier(t *testing.T) {
	dismantleBelow, _, status := configureVerifyScenario(t)
	defer dismantleBelow(t)

	originPollingEnergy := kinds.MaximumSumPollingEnergy / 1000
	originPublicKey := ed25519.GeneratePrivateKey().PublicKey()
	//
	originValue := &kinds.Ratifier{
		Location:     originPublicKey.Location(),
		PublicKey:      originPublicKey,
		PollingEnergy: originPollingEnergy,
	}
	//
	status.Ratifiers = kinds.NewRatifierCollection([]*kinds.Ratifier{originValue})
	status.FollowingRatifiers = status.Ratifiers
	require.True(t, len(status.Ratifiers.Ratifiers) == 1)

	//
	//
	agedStatus := status
	for i := 0; i < 10; i++ {
		//
		ifaceReplies := &iface.ReplyCompleteLedger{}
		ratifierRefreshes, err := kinds.Schema2tm.RatifierRefreshes(ifaceReplies.RatifierRefreshes)
		require.NoError(t, err)

		ledger, err := createLedger(agedStatus, agedStatus.FinalLedgerLevel+1, new(kinds.Endorse))
		require.NoError(t, err)
		bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
		require.NoError(t, err)
		ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}

		refreshedStatus, err := sm.ModifyStatus(agedStatus, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
		require.NoError(t, err)
		//
		//
		//
		assert.EqualValues(t, agedStatus.FollowingRatifiers, refreshedStatus.FollowingRatifiers)
		assert.EqualValues(t, 0, refreshedStatus.FollowingRatifiers.Recommender.RecommenderUrgency)

		agedStatus = refreshedStatus
	}
	//
	//
	//
	//
	//
	initialAppendedValuePublicKey := ed25519.GeneratePrivateKey().PublicKey()
	initialAppendedValuePollingEnergy := int64(10)
	fvp, err := cryptocode.PublicKeyToSchema(initialAppendedValuePublicKey)
	require.NoError(t, err)
	initialAppendedValue := iface.RatifierModify{PublicKey: fvp, Energy: initialAppendedValuePollingEnergy}
	ratifierRefreshes, err := kinds.Schema2tm.RatifierRefreshes([]iface.RatifierModify{initialAppendedValue})
	assert.NoError(t, err)
	ifaceReplies := &iface.ReplyCompleteLedger{
		RatifierRefreshes: []iface.RatifierModify{initialAppendedValue},
	}
	ledger, err := createLedger(agedStatus, agedStatus.FinalLedgerLevel+1, new(kinds.Endorse))
	require.NoError(t, err)

	bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
	require.NoError(t, err)

	ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}
	refreshedStatus, err := sm.ModifyStatus(agedStatus, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
	require.NoError(t, err)

	finalStatus := refreshedStatus
	for i := 0; i < 200; i++ {
		//
		ifaceReplies := &iface.ReplyCompleteLedger{}
		ratifierRefreshes, err := kinds.Schema2tm.RatifierRefreshes(ifaceReplies.RatifierRefreshes)
		require.NoError(t, err)

		ledger, err := createLedger(finalStatus, finalStatus.FinalLedgerLevel+1, new(kinds.Endorse))
		require.NoError(t, err)

		bps, err = ledger.CreateSegmentAssign(verifySegmentVolume)
		require.NoError(t, err)

		ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}

		refreshedStatusDeeper, err := sm.ModifyStatus(finalStatus, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
		require.NoError(t, err)
		finalStatus = refreshedStatusDeeper
	}
	//
	status = finalStatus

	//
	agedStatus = refreshedStatus
	_, agedOriginValue := agedStatus.FollowingRatifiers.FetchByLocation(originValue.Location)
	_, newOriginValue := status.FollowingRatifiers.FetchByLocation(originValue.Location)
	_, appendedAgedValue := agedStatus.FollowingRatifiers.FetchByLocation(initialAppendedValuePublicKey.Location())
	_, appendedNewValue := status.FollowingRatifiers.FetchByLocation(initialAppendedValuePublicKey.Location())
	//
	assert.True(t, agedOriginValue.RecommenderUrgency > newOriginValue.RecommenderUrgency)
	assert.True(t, appendedAgedValue.RecommenderUrgency < appendedNewValue.RecommenderUrgency)

	//
	for i := 0; i < 10; i++ {
		appendedPublicKey := ed25519.GeneratePrivateKey().PublicKey()
		ap, err := cryptocode.PublicKeyToSchema(appendedPublicKey)
		require.NoError(t, err)
		appendedValue := iface.RatifierModify{PublicKey: ap, Energy: initialAppendedValuePollingEnergy}
		ratifierRefreshes, err := kinds.Schema2tm.RatifierRefreshes([]iface.RatifierModify{appendedValue})
		assert.NoError(t, err)

		ifaceReplies := &iface.ReplyCompleteLedger{
			RatifierRefreshes: []iface.RatifierModify{appendedValue},
		}
		ledger, err := createLedger(agedStatus, agedStatus.FinalLedgerLevel+1, new(kinds.Endorse))
		require.NoError(t, err)
		bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
		require.NoError(t, err)

		ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}
		status, err = sm.ModifyStatus(status, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
		require.NoError(t, err)
	}
	require.Equal(t, 10+2, len(status.FollowingRatifiers.Ratifiers))

	//
	gp, err := cryptocode.PublicKeyToSchema(originPublicKey)
	require.NoError(t, err)
	deleteOriginValue := iface.RatifierModify{PublicKey: gp, Energy: 0}
	ifaceReplies = &iface.ReplyCompleteLedger{
		RatifierRefreshes: []iface.RatifierModify{deleteOriginValue},
	}

	ledger, err = createLedger(agedStatus, agedStatus.FinalLedgerLevel+1, new(kinds.Endorse))
	require.NoError(t, err)

	bps, err = ledger.CreateSegmentAssign(verifySegmentVolume)
	require.NoError(t, err)

	ledgerUID = kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}
	ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes(ifaceReplies.RatifierRefreshes)
	require.NoError(t, err)
	refreshedStatus, err = sm.ModifyStatus(status, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
	require.NoError(t, err)
	//
	assert.Equal(t, 11, len(refreshedStatus.FollowingRatifiers.Ratifiers))

	//
	//
	currentStatus := refreshedStatus
	tally := 0
	isRecommenderUnmodified := true
	for isRecommenderUnmodified {
		ifaceReplies := &iface.ReplyCompleteLedger{}
		ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes(ifaceReplies.RatifierRefreshes)
		require.NoError(t, err)
		ledger, err = createLedger(currentStatus, currentStatus.FinalLedgerLevel+1, new(kinds.Endorse))
		require.NoError(t, err)

		bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
		require.NoError(t, err)

		ledgerUID = kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}
		currentStatus, err = sm.ModifyStatus(currentStatus, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
		require.NoError(t, err)
		if !bytes.Equal(currentStatus.Ratifiers.Recommender.Location, currentStatus.FollowingRatifiers.Recommender.Location) {
			isRecommenderUnmodified = false
		}
		tally++
	}
	refreshedStatus = currentStatus
	//
	initialRecommenderAlterAnticipatedAfter := 1
	assert.Equal(t, initialRecommenderAlterAnticipatedAfter, tally)
	//
	countValues := len(refreshedStatus.Ratifiers.Ratifiers)
	recommenders := make([]*kinds.Ratifier, countValues)
	for i := 0; i < 100; i++ {
		//
		ifaceReplies := &iface.ReplyCompleteLedger{}
		ratifierRefreshes, err := kinds.Schema2tm.RatifierRefreshes(ifaceReplies.RatifierRefreshes)
		require.NoError(t, err)

		ledger, err := createLedger(refreshedStatus, refreshedStatus.FinalLedgerLevel+1, new(kinds.Endorse))
		require.NoError(t, err)

		bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
		require.NoError(t, err)

		ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}

		refreshedStatus, err = sm.ModifyStatus(refreshedStatus, ledgerUID, &ledger.Heading, ifaceReplies, ratifierRefreshes)
		require.NoError(t, err)
		if i > countValues { //
			if recommenders[i%countValues] == nil {
				recommenders[i%countValues] = refreshedStatus.FollowingRatifiers.Recommender
			} else {
				assert.Equal(t, recommenders[i%countValues], refreshedStatus.FollowingRatifiers.Recommender)
			}
		}
	}
}

func VerifyDepotImportRatifiersAdditionsRecommenderUrgency(t *testing.T) {
	const valueCollectionVolume = 2
	dismantleBelow, statusStore, status := configureVerifyScenario(t)
	t.Cleanup(func() { dismantleBelow(t) })
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	status.Ratifiers = generateValueCollection(valueCollectionVolume)
	status.FollowingRatifiers = status.Ratifiers.CloneAugmentRecommenderUrgency(1)
	err := statusDepot.Persist(status)
	require.NoError(t, err)

	followingLevel := status.FinalLedgerLevel + 1

	v0, err := statusDepot.ImportRatifiers(followingLevel)
	assert.Nil(t, err)
	tally0 := v0.Ratifiers[0].RecommenderUrgency

	v1, err := statusDepot.ImportRatifiers(followingLevel + 1)
	assert.Nil(t, err)
	tally1 := v1.Ratifiers[0].RecommenderUrgency

	assert.NotEqual(t, tally1, tally0, "REDACTED")
}

//
//
func VerifyNumerousRatifierModificationsPersistImport(t *testing.T) {
	const valueCollectionVolume = 7
	dismantleBelow, statusStore, status := configureVerifyScenario(t)
	defer dismantleBelow(t)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	require.Equal(t, int64(0), status.FinalLedgerLevel)
	status.Ratifiers = generateValueCollection(valueCollectionVolume)
	status.FollowingRatifiers = status.Ratifiers.CloneAugmentRecommenderUrgency(1)
	err := statusDepot.Persist(status)
	require.NoError(t, err)

	_, valueAged := status.Ratifiers.FetchByOrdinal(0)
	publickeyAged := valueAged.PublicKey
	publickey := ed25519.GeneratePrivateKey().PublicKey()

	//
	heading, ledgerUID, replies := createHeadingSectionsRepliesValuePublicKeyAlter(status, publickey)

	//
	var ratifierRefreshes []*kinds.Ratifier
	ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes(replies.RatifierRefreshes)
	require.NoError(t, err)
	status, err = sm.ModifyStatus(status, ledgerUID, &heading, replies, ratifierRefreshes)
	require.Nil(t, err)
	followingLevel := status.FinalLedgerLevel + 1
	err = statusDepot.Persist(status)
	require.NoError(t, err)

	//
	v0, err := statusDepot.ImportRatifiers(followingLevel)
	assert.Nil(t, err)
	assert.Equal(t, valueCollectionVolume, v0.Volume())
	ordinal, val := v0.FetchByLocation(publickeyAged.Location())
	assert.NotNil(t, val)
	if ordinal < 0 {
		t.Fatal("REDACTED")
	}

	//
	v1, err := statusDepot.ImportRatifiers(followingLevel + 1)
	assert.Nil(t, err)
	assert.Equal(t, valueCollectionVolume, v1.Volume())
	ordinal, val = v1.FetchByLocation(publickey.Location())
	assert.NotNil(t, val)
	if ordinal < 0 {
		t.Fatal("REDACTED")
	}
}

func VerifyStatusCreateLedger(t *testing.T) {
	dismantleBelow, _, status := configureVerifyScenario(t)
	defer dismantleBelow(t)

	recommenderLocation := status.Ratifiers.FetchRecommender().Location
	statusRelease := status.Release.Agreement
	ledger, err := createLedger(status, 2, new(kinds.Endorse))
	require.NoError(t, err)

	//
	assert.Equal(t, statusRelease, ledger.Release)
	assert.Equal(t, recommenderLocation, ledger.RecommenderLocation)
}

//
//
func VerifyAgreementOptionsModificationsPersistImport(t *testing.T) {
	dismantleBelow, statusStore, status := configureVerifyScenario(t)
	defer dismantleBelow(t)

	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	//
	alterLevels := []int64{1, 2, 4, 5, 10, 15, 16, 17, 20}
	N := len(alterLevels)

	//
	//
	options := make([]kinds.AgreementOptions, N+1)
	options[0] = status.AgreementOptions
	for i := 1; i < N+1; i++ {
		options[i] = *kinds.StandardAgreementOptions()
		options[i].Ledger.MaximumOctets += int64(i)

	}

	//
	//
	maximumLevel := alterLevels[N-1] + 5
	alterOrdinal := 0
	cp := options[alterOrdinal]
	var err error
	var ratifierRefreshes []*kinds.Ratifier
	for i := int64(1); i < maximumLevel; i++ {
		//
		if alterOrdinal < len(alterLevels) && i == alterLevels[alterOrdinal] {
			alterOrdinal++
			cp = options[alterOrdinal]
		}
		heading, ledgerUID, replies := createHeadingSectionsRepliesOptions(status, cp.ToSchema())
		ratifierRefreshes, err = kinds.Schema2tm.RatifierRefreshes(replies.RatifierRefreshes)
		require.NoError(t, err)
		status, err = sm.ModifyStatus(status, ledgerUID, &heading, replies, ratifierRefreshes)

		require.NoError(t, err)
		err = statusDepot.Persist(status)
		require.NoError(t, err)
	}

	//
	verifyScenarios := make([]optionsAlterVerifyScenario, maximumLevel)
	alterOrdinal = 0
	cp = options[alterOrdinal]
	for i := int64(1); i < maximumLevel+1; i++ {
		//
		//
		if alterOrdinal < len(alterLevels) && i == alterLevels[alterOrdinal]+1 {
			alterOrdinal++
			cp = options[alterOrdinal]
		}
		verifyScenarios[i-1] = optionsAlterVerifyScenario{i, cp}
	}

	for _, verifyInstance := range verifyScenarios {
		p, err := statusDepot.ImportAgreementOptions(verifyInstance.level)
		assert.Nil(t, err, fmt.Sprintf("REDACTED", verifyInstance.level))
		assert.EqualValues(t, verifyInstance.options, p, fmt.Sprintf(`REDACTEDt
REDACTED`, verifyInstance.level))
	}
}

func VerifyStatusSchema(t *testing.T) {
	dismantleBelow, _, status := configureVerifyScenario(t)
	defer dismantleBelow(t)

	tc := []struct {
		verifyLabel string
		status    *sm.Status
		expirationPass1 bool
		expirationPass2 bool
	}{
		{"REDACTED", &sm.Status{}, true, false},
		{"REDACTED", nil, false, false},
		{"REDACTED", &status, true, true},
	}

	for _, tt := range tc {

		pbs, err := tt.status.ToSchema()
		if !tt.expirationPass1 {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err, tt.verifyLabel)
		}

		smt, err := sm.FromSchema(pbs)
		if tt.expirationPass2 {
			require.NoError(t, err, tt.verifyLabel)
			require.Equal(t, tt.status, smt, tt.verifyLabel)
		} else {
			require.Error(t, err, tt.verifyLabel)
		}
	}
}

func VerifyMidpointTime(t *testing.T) {
	value1 := kinds.NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 30)
	value2 := kinds.NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 30)
	value3 := kinds.NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 30)

	values := kinds.NewRatifierCollection([]*kinds.Ratifier{value1, value2, value3})

	t.Run("REDACTED", func(t *testing.T) {
		now := time.Now()
		endorse := &kinds.Endorse{
			Level: 1,
			Endorsements: []kinds.EndorseSignature{
				{
					LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
					RatifierLocation: value1.Location,
					Timestamp:        now,
				},
				{
					LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
					RatifierLocation: value2.Location,
					Timestamp:        now.Add(1 * time.Minute),
				},
				{
					LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
					RatifierLocation: value3.Location,
					Timestamp:        now.Add(2 * time.Minute),
				},
			},
		}

		midpointTime, err := sm.MidpointTime(endorse, values)
		require.NoError(t, err)
		require.Equal(t, midpointTime, now.Add(1*time.Minute))
	})

	t.Run("REDACTED", func(t *testing.T) {
		unclearValue := ed25519.GeneratePrivateKey().PublicKey().Location()
		now := time.Now()
		endorse := &kinds.Endorse{
			Level: 1,
			Endorsements: []kinds.EndorseSignature{
				{
					LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
					RatifierLocation: value1.Location,
					Timestamp:        now,
				},
				{
					LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
					RatifierLocation: unclearValue,
					Timestamp:        now.Add(1 * time.Minute),
				},
			},
		}

		_, err := sm.MidpointTime(endorse, values)
		require.Error(t, err)
		require.Contains(t, err.Error(), "REDACTED")
	})

	t.Run("REDACTED", func(t *testing.T) {
		now := time.Now()
		endorse := &kinds.Endorse{
			Level: 1,
			Endorsements: []kinds.EndorseSignature{
				{
					LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
					RatifierLocation: value1.Location,
					Timestamp:        now,
				},
				{
					LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
					RatifierLocation: value2.Location,
					Timestamp:        now.Add(1 * time.Minute),
				},
			},
		}

		midpointTime, err := sm.MidpointTime(endorse, values)
		require.NoError(t, err)
		require.Equal(t, midpointTime, now)
	})
}
