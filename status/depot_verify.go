package status_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/intrinsic/verify"
	cometstatus "github.com/valkyrieworks/schema/consensuscore/status"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

func VerifyDepotImportRatifiers(t *testing.T) {
	statusStore := dbm.NewMemoryStore()
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	val, _ := kinds.RandomRatifier(true, 10)
	values := kinds.NewRatifierCollection([]*kinds.Ratifier{val})

	//
	err := sm.PersistRatifiersDetails(statusStore, 1, 1, values)
	require.NoError(t, err)
	err = sm.PersistRatifiersDetails(statusStore, 2, 1, values)
	require.NoError(t, err)
	retrievedValues, err := statusDepot.ImportRatifiers(2)
	require.NoError(t, err)
	assert.NotZero(t, retrievedValues.Volume())

	//

	err = sm.PersistRatifiersDetails(statusStore, sm.ValueCollectionMilestoneCadence, 1, values)
	require.NoError(t, err)

	retrievedValues, err = statusDepot.ImportRatifiers(sm.ValueCollectionMilestoneCadence)
	require.NoError(t, err)
	assert.NotZero(t, retrievedValues.Volume())
}

func CriterionImportRatifiers(b *testing.B) {
	const valueCollectionVolume = 100

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	storeKind := dbm.OriginKind(settings.StoreOrigin)
	statusStore, err := dbm.NewStore("REDACTED", storeKind, settings.StoreFolder())
	require.NoError(b, err)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	status, err := statusDepot.ImportFromStoreOrOriginEntry(settings.OriginEntry())
	if err != nil {
		b.Fatal(err)
	}

	status.Ratifiers = generateValueCollection(valueCollectionVolume)
	status.FollowingRatifiers = status.Ratifiers.CloneAugmentRecommenderUrgency(1)
	err = statusDepot.Persist(status)
	require.NoError(b, err)

	for i := 10; i < 10000000000; i *= 10 { //

		if err := sm.PersistRatifiersDetails(statusStore,
			int64(i), status.FinalLevelRatifiersModified, status.FollowingRatifiers); err != nil {
			b.Fatal(err)
		}

		b.Run(fmt.Sprintf("REDACTED", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, err := statusDepot.ImportRatifiers(int64(i))
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func VerifyTrimConditions(t *testing.T) {
	verifyscenarios := map[string]struct {
		createLevels             int64
		trimFrom               int64
		trimTo                 int64
		proofLimitLevel int64
		anticipateErr               bool
		anticipateValues              []int64
		anticipateOptions            []int64
		anticipateIface              []int64
	}{
		"REDACTED":      {100, 0, 5, 100, true, nil, nil, nil},
		"REDACTED":         {100, 3, 2, 2, true, nil, nil, nil},
		"REDACTED":        {100, 3, 3, 3, true, nil, nil, nil},
		"REDACTED": {100, 1, 101, 101, true, nil, nil, nil},
		"REDACTED":                    {100, 1, 100, 100, false, []int64{93, 100}, []int64{95, 100}, []int64{100}},
		"REDACTED": {
			10, 2, 8, 8, false,
			[]int64{1, 3, 8, 9, 10},
			[]int64{1, 5, 8, 9, 10},
			[]int64{1, 8, 9, 10},
		},
		"REDACTED": {
			100001, 1, 100001, 100001, false,
			[]int64{99993, 100000, 100001},
			[]int64{99995, 100001},
			[]int64{100001},
		},
		"REDACTED": {20, 1, 18, 17, false, []int64{13, 17, 18, 19, 20}, []int64{15, 18, 19, 20}, []int64{18, 19, 20}},
	}
	for label, tc := range verifyscenarios {

		t.Run(label, func(t *testing.T) {
			db := dbm.NewMemoryStore()
			statusDepot := sm.NewDepot(db, sm.DepotSettings{
				DropIfaceReplies: false,
			})
			pk := ed25519.GeneratePrivateKey().PublicKey()

			//
			//
			ratifier := &kinds.Ratifier{Location: pk.Location(), PollingEnergy: 100, PublicKey: pk}
			ratifierCollection := &kinds.RatifierAssign{
				Ratifiers: []*kinds.Ratifier{ratifier},
				Recommender:   ratifier,
			}
			valuesModified := int64(0)
			optionsModified := int64(0)

			for h := int64(1); h <= tc.createLevels; h++ {
				if valuesModified == 0 || h%10 == 2 {
					valuesModified = h + 1 //
				}
				if optionsModified == 0 || h%10 == 5 {
					optionsModified = h
				}

				status := sm.Status{
					PrimaryLevel:   1,
					FinalLedgerLevel: h - 1,
					Ratifiers:      ratifierCollection,
					FollowingRatifiers:  ratifierCollection,
					AgreementOptions: kinds.AgreementOptions{
						Ledger: kinds.LedgerOptions{MaximumOctets: 10e6},
					},
					FinalLevelRatifiersModified:      valuesModified,
					FinalLevelAgreementOptionsModified: optionsModified,
				}

				if status.FinalLedgerLevel >= 1 {
					status.FinalRatifiers = status.Ratifiers
				}

				err := statusDepot.Persist(status)
				require.NoError(t, err)

				err = statusDepot.PersistCompleteLedgerReply(h, &iface.ReplyCompleteLedger{
					TransOutcomes: []*iface.InvokeTransferOutcome{
						{Data: []byte{1}},
						{Data: []byte{2}},
						{Data: []byte{3}},
					},
					ApplicationDigest: make([]byte, 1),
				})
				require.NoError(t, err)
			}

			//
			err := statusDepot.TrimConditions(tc.trimFrom, tc.trimTo, tc.proofLimitLevel)
			if tc.anticipateErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			anticipateValues := segmentToIndex(tc.anticipateValues)
			anticipateOptions := segmentToIndex(tc.anticipateOptions)
			anticipateIface := segmentToIndex(tc.anticipateIface)

			for h := int64(1); h <= tc.createLevels; h++ {
				values, err := statusDepot.ImportRatifiers(h)
				if anticipateValues[h] {
					require.NoError(t, err, "REDACTED", h)
					require.NotNil(t, values)
				} else {
					require.Error(t, err, "REDACTED", h)
					require.Equal(t, sm.ErrNoValueCollectionForLevel{Level: h}, err)
				}

				options, err := statusDepot.ImportAgreementOptions(h)
				if anticipateOptions[h] {
					require.NoError(t, err, "REDACTED", h)
					require.NotEmpty(t, options)
				} else {
					require.Error(t, err, "REDACTED", h)
					require.Empty(t, options)
				}

				iface, err := statusDepot.ImportCompleteLedgerReply(h)
				if anticipateIface[h] {
					require.NoError(t, err, "REDACTED", h)
					require.NotNil(t, iface)
				} else {
					require.Error(t, err, "REDACTED", h)
					require.Equal(t, sm.ErrNoIfaceRepliesForLevel{Level: h}, err)
				}
			}
		})
	}
}

func VerifyTransferOutcomesDigest(t *testing.T) {
	transferOutcomes := []*iface.InvokeTransferOutcome{
		{Code: 32, Data: []byte("REDACTED"), Log: "REDACTED"},
	}

	origin := sm.TransferOutcomesDigest(transferOutcomes)

	//
	outcomes := kinds.NewOutcomes(transferOutcomes)
	assert.Equal(t, origin, outcomes.Digest())

	//
	evidence := outcomes.DemonstrateOutcome(0)
	bz, err := outcomes[0].Serialize()
	require.NoError(t, err)
	assert.NoError(t, evidence.Validate(origin, bz))
}

func segmentToIndex(s []int64) map[int64]bool {
	m := make(map[int64]bool, len(s))
	for _, i := range s {
		m[i] = true
	}
	return m
}

func VerifyFinalCompleteLedgerReplies(t *testing.T) {
	//
	t.Run("REDACTED", func(t *testing.T) {
		statusStore := dbm.NewMemoryStore()
		statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
			DropIfaceReplies: false,
		})
		replies, err := statusDepot.ImportCompleteLedgerReply(1)
		require.Error(t, err)
		require.Nil(t, replies)
		//
		reply1 := &iface.ReplyCompleteLedger{
			TransOutcomes: []*iface.InvokeTransferOutcome{
				{Code: 32, Data: []byte("REDACTED"), Log: "REDACTED"},
			},
			ApplicationDigest: make([]byte, 1),
		}
		//
		statusStore = dbm.NewMemoryStore()
		statusDepot = sm.NewDepot(statusStore, sm.DepotSettings{DropIfaceReplies: false})
		level := int64(10)
		//
		err = statusDepot.PersistCompleteLedgerReply(level, reply1)
		require.NoError(t, err)
		//
		finalReply, err := statusDepot.ImportFinalCompleteLedgerReply(level)
		require.NoError(t, err)
		//
		assert.Equal(t, finalReply, reply1)
		//
		_, err = statusDepot.ImportFinalCompleteLedgerReply(level + 1)
		assert.Error(t, err)
		//
		replies, err = statusDepot.ImportCompleteLedgerReply(level)
		require.NoError(t, err, replies)
		require.Equal(t, reply1, replies)
	})

	t.Run("REDACTED", func(t *testing.T) {
		statusStore := dbm.NewMemoryStore()
		level := int64(10)
		//
		reply2 := &iface.ReplyCompleteLedger{
			TransOutcomes: []*iface.InvokeTransferOutcome{
				{Code: 44, Data: []byte("REDACTED"), Log: "REDACTED"},
			},
		}
		//
		statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
			DropIfaceReplies: true,
		})
		//
		err := statusDepot.PersistCompleteLedgerReply(level+1, reply2)
		require.NoError(t, err)
		//
		finalReply2, err := statusDepot.ImportFinalCompleteLedgerReply(level + 1)
		require.NoError(t, err)
		//
		assert.Equal(t, reply2, finalReply2)
		//
		_, err = statusDepot.ImportCompleteLedgerReply(level + 1)
		assert.Equal(t, sm.ErrCompleteLedgerRepliesNotSustained, err)
	})
}

func VerifyCompleteLedgerRestoreUtilizingPastIfaceReplies(t *testing.T) {
	var (
		level              int64 = 10
		finalIfaceReplyKey       = []byte("REDACTED")
		memoryStore                     = dbm.NewMemoryStore()
		cp                        = kinds.StandardAgreementOptions().ToSchema()
		pastReply                = cometstatus.IfaceRepliesDetails{
			PastIfaceReplies: &cometstatus.PastIfaceReplies{
				InitiateLedger: &cometstatus.AnswerInitiateLedger{
					Events: []iface.Event{{
						Kind: "REDACTED",
						Properties: []iface.EventProperty{{
							Key:   "REDACTED",
							Item: "REDACTED",
						}},
					}},
				},
				DispatchTrans: []*iface.InvokeTransferOutcome{{
					Events: []iface.Event{{
						Kind: "REDACTED",
						Properties: []iface.EventProperty{{
							Key:   "REDACTED",
							Item: "REDACTED",
						}},
					}},
				}},
				TerminateLedger: &cometstatus.AnswerTerminateLedger{
					AgreementArgumentRefreshes: &cp,
				},
			},
			Level: level,
		}
	)
	bz, err := pastReply.Serialize()
	require.NoError(t, err)
	//
	require.NoError(t, memoryStore.Set(finalIfaceReplyKey, bz))
	statusDepot := sm.NewDepot(memoryStore, sm.DepotSettings{DropIfaceReplies: false})
	reply, err := statusDepot.ImportFinalCompleteLedgerReply(level)
	require.NoError(t, err)
	require.Equal(t, reply.AgreementArgumentRefreshes, &cp)
	require.Equal(t, len(reply.Events), len(pastReply.PastIfaceReplies.InitiateLedger.Events))
	require.Equal(t, reply.TransOutcomes[0], pastReply.PastIfaceReplies.DispatchTrans[0])
}

func VerifyIntegerTransform(t *testing.T) {
	x := int64(10)
	b := sm.Int64toOctets(x)
	require.Equal(t, x, sm.Int64fromOctets(b))
}
