package status_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	strongstatus "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/status"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyDepotFetchAssessors(t *testing.T) {
	statusDatastore := dbm.FreshMemoryDatastore()
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	val, _ := kinds.ArbitraryAssessor(true, 10)
	values := kinds.FreshAssessorAssign([]*kinds.Assessor{val})

	//
	err := sm.PersistAssessorsDetails(statusDatastore, 1, 1, values)
	require.NoError(t, err)
	err = sm.PersistAssessorsDetails(statusDatastore, 2, 1, values)
	require.NoError(t, err)
	retrievedValues, err := statusDepot.FetchAssessors(2)
	require.NoError(t, err)
	assert.NotZero(t, retrievedValues.Extent())

	//

	err = sm.PersistAssessorsDetails(statusDatastore, sm.ItemAssignMilestoneDuration, 1, values)
	require.NoError(t, err)

	retrievedValues, err = statusDepot.FetchAssessors(sm.ItemAssignMilestoneDuration)
	require.NoError(t, err)
	assert.NotZero(t, retrievedValues.Extent())
}

func AssessmentFetchAssessors(b *testing.B) {
	const itemAssignExtent = 100

	settings := verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	datastoreKind := dbm.OriginKind(settings.DatastoreRepository)
	statusDatastore, err := dbm.FreshDatastore("REDACTED", datastoreKind, settings.DatastorePath())
	require.NoError(b, err)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	status, err := statusDepot.FetchOriginatingDatastoreEitherInaugurationRecord(settings.InaugurationRecord())
	if err != nil {
		b.Fatal(err)
	}

	status.Assessors = produceItemAssign(itemAssignExtent)
	status.FollowingAssessors = status.Assessors.DuplicateAdvanceNominatorUrgency(1)
	err = statusDepot.Persist(status)
	require.NoError(b, err)

	for i := 10; i < 10000000000; i *= 10 { //

		if err := sm.PersistAssessorsDetails(statusDatastore,
			int64(i), status.FinalAltitudeAssessorsAltered, status.FollowingAssessors); err != nil {
			b.Fatal(err)
		}

		b.Run(fmt.Sprintf("REDACTED", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, err := statusDepot.FetchAssessors(int64(i))
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func VerifyTrimStatuses(t *testing.T) {
	verifycases := map[string]struct {
		createElevations             int64
		trimOriginating               int64
		trimToward                 int64
		proofLimitAltitude int64
		anticipateFault               bool
		anticipateValues              []int64
		anticipateParameters            []int64
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
	for alias, tc := range verifycases {

		t.Run(alias, func(t *testing.T) {
			db := dbm.FreshMemoryDatastore()
			statusDepot := sm.FreshDepot(db, sm.DepotChoices{
				EjectIfaceReplies: false,
			})
			pk := edwards25519.ProducePrivateToken().PublicToken()

			//
			//
			assessor := &kinds.Assessor{Location: pk.Location(), BallotingPotency: 100, PublicToken: pk}
			assessorAssign := &kinds.AssessorAssign{
				Assessors: []*kinds.Assessor{assessor},
				Nominator:   assessor,
			}
			valuesAltered := int64(0)
			parametersAltered := int64(0)

			for h := int64(1); h <= tc.createElevations; h++ {
				if valuesAltered == 0 || h%10 == 2 {
					valuesAltered = h + 1 //
				}
				if parametersAltered == 0 || h%10 == 5 {
					parametersAltered = h
				}

				status := sm.Status{
					PrimaryAltitude:   1,
					FinalLedgerAltitude: h - 1,
					Assessors:      assessorAssign,
					FollowingAssessors:  assessorAssign,
					AgreementSettings: kinds.AgreementSettings{
						Ledger: kinds.LedgerParameters{MaximumOctets: 10e6},
					},
					FinalAltitudeAssessorsAltered:      valuesAltered,
					FinalAltitudeAgreementParametersAltered: parametersAltered,
				}

				if status.FinalLedgerAltitude >= 1 {
					status.FinalAssessors = status.Assessors
				}

				err := statusDepot.Persist(status)
				require.NoError(t, err)

				err = statusDepot.PersistCulminateLedgerReply(h, &iface.ReplyCulminateLedger{
					TransferOutcomes: []*iface.InvokeTransferOutcome{
						{Data: []byte{1}},
						{Data: []byte{2}},
						{Data: []byte{3}},
					},
					PlatformDigest: make([]byte, 1),
				})
				require.NoError(t, err)
			}

			//
			err := statusDepot.TrimStatuses(tc.trimOriginating, tc.trimToward, tc.proofLimitAltitude)
			if tc.anticipateFault {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			anticipateValues := sectionTowardIndex(tc.anticipateValues)
			anticipateParameters := sectionTowardIndex(tc.anticipateParameters)
			anticipateIface := sectionTowardIndex(tc.anticipateIface)

			for h := int64(1); h <= tc.createElevations; h++ {
				values, err := statusDepot.FetchAssessors(h)
				if anticipateValues[h] {
					require.NoError(t, err, "REDACTED", h)
					require.NotNil(t, values)
				} else {
					require.Error(t, err, "REDACTED", h)
					require.Equal(t, sm.FaultNegativeItemAssignForeachAltitude{Altitude: h}, err)
				}

				parameters, err := statusDepot.FetchAgreementParameters(h)
				if anticipateParameters[h] {
					require.NoError(t, err, "REDACTED", h)
					require.NotEmpty(t, parameters)
				} else {
					require.Error(t, err, "REDACTED", h)
					require.Empty(t, parameters)
				}

				iface, err := statusDepot.FetchCulminateLedgerReply(h)
				if anticipateIface[h] {
					require.NoError(t, err, "REDACTED", h)
					require.NotNil(t, iface)
				} else {
					require.Error(t, err, "REDACTED", h)
					require.Equal(t, sm.FaultNegativeIfaceRepliesForeachAltitude{Altitude: h}, err)
				}
			}
		})
	}
}

func VerifyTransferOutcomesDigest(t *testing.T) {
	transferOutcomes := []*iface.InvokeTransferOutcome{
		{Cipher: 32, Data: []byte("REDACTED"), Log: "REDACTED"},
	}

	origin := sm.TransferOutcomesDigest(transferOutcomes)

	//
	outcomes := kinds.FreshOutcomes(transferOutcomes)
	assert.Equal(t, origin, outcomes.Digest())

	//
	attestation := outcomes.AscertainOutcome(0)
	bz, err := outcomes[0].Serialize()
	require.NoError(t, err)
	assert.NoError(t, attestation.Validate(origin, bz))
}

func sectionTowardIndex(s []int64) map[int64]bool {
	m := make(map[int64]bool, len(s))
	for _, i := range s {
		m[i] = true
	}
	return m
}

func VerifyFinalCulminateLedgerReplies(t *testing.T) {
	//
	t.Run("REDACTED", func(t *testing.T) {
		statusDatastore := dbm.FreshMemoryDatastore()
		statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
			EjectIfaceReplies: false,
		})
		replies, err := statusDepot.FetchCulminateLedgerReply(1)
		require.Error(t, err)
		require.Nil(t, replies)
		//
		reply1 := &iface.ReplyCulminateLedger{
			TransferOutcomes: []*iface.InvokeTransferOutcome{
				{Cipher: 32, Data: []byte("REDACTED"), Log: "REDACTED"},
			},
			PlatformDigest: make([]byte, 1),
		}
		//
		statusDatastore = dbm.FreshMemoryDatastore()
		statusDepot = sm.FreshDepot(statusDatastore, sm.DepotChoices{EjectIfaceReplies: false})
		altitude := int64(10)
		//
		err = statusDepot.PersistCulminateLedgerReply(altitude, reply1)
		require.NoError(t, err)
		//
		finalReply, err := statusDepot.FetchFinalCulminateLedgerReply(altitude)
		require.NoError(t, err)
		//
		assert.Equal(t, finalReply, reply1)
		//
		_, err = statusDepot.FetchFinalCulminateLedgerReply(altitude + 1)
		assert.Error(t, err)
		//
		replies, err = statusDepot.FetchCulminateLedgerReply(altitude)
		require.NoError(t, err, replies)
		require.Equal(t, reply1, replies)
	})

	t.Run("REDACTED", func(t *testing.T) {
		statusDatastore := dbm.FreshMemoryDatastore()
		altitude := int64(10)
		//
		reply2 := &iface.ReplyCulminateLedger{
			TransferOutcomes: []*iface.InvokeTransferOutcome{
				{Cipher: 44, Data: []byte("REDACTED"), Log: "REDACTED"},
			},
		}
		//
		statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
			EjectIfaceReplies: true,
		})
		//
		err := statusDepot.PersistCulminateLedgerReply(altitude+1, reply2)
		require.NoError(t, err)
		//
		finalReply2, err := statusDepot.FetchFinalCulminateLedgerReply(altitude + 1)
		require.NoError(t, err)
		//
		assert.Equal(t, reply2, finalReply2)
		//
		_, err = statusDepot.FetchCulminateLedgerReply(altitude + 1)
		assert.Equal(t, sm.FaultCulminateLedgerRepliesNegationStored, err)
	})
}

func VerifyCulminateLedgerRecuperationApplyingHeritageIfaceReplies(t *testing.T) {
	var (
		altitude              int64 = 10
		finalIfaceReplyToken       = []byte("REDACTED")
		memoryDatastore                     = dbm.FreshMemoryDatastore()
		cp                        = kinds.FallbackAgreementSettings().TowardSchema()
		heritageAnswer                = strongstatus.IfaceRepliesDetails{
			HeritageIfaceReplies: &strongstatus.HeritageIfaceReplies{
				InitiateLedger: &strongstatus.ReplyInitiateLedger{
					Incidents: []iface.Incident{{
						Kind: "REDACTED",
						Properties: []iface.IncidentProperty{{
							Key:   "REDACTED",
							Datum: "REDACTED",
						}},
					}},
				},
				DispatchTrans: []*iface.InvokeTransferOutcome{{
					Incidents: []iface.Incident{{
						Kind: "REDACTED",
						Properties: []iface.IncidentProperty{{
							Key:   "REDACTED",
							Datum: "REDACTED",
						}},
					}},
				}},
				TerminateLedger: &strongstatus.ReplyTerminateLedger{
					AgreementArgumentRevisions: &cp,
				},
			},
			Altitude: altitude,
		}
	)
	bz, err := heritageAnswer.Serialize()
	require.NoError(t, err)
	//
	require.NoError(t, memoryDatastore.Set(finalIfaceReplyToken, bz))
	statusDepot := sm.FreshDepot(memoryDatastore, sm.DepotChoices{EjectIfaceReplies: false})
	reply, err := statusDepot.FetchFinalCulminateLedgerReply(altitude)
	require.NoError(t, err)
	require.Equal(t, reply.AgreementArgumentRevisions, &cp)
	require.Equal(t, len(reply.Incidents), len(heritageAnswer.HeritageIfaceReplies.InitiateLedger.Incidents))
	require.Equal(t, reply.TransferOutcomes[0], heritageAnswer.HeritageIfaceReplies.DispatchTrans[0])
}

func VerifyIntegerAdaptation(t *testing.T) {
	x := int64(10)
	b := sm.Integer64towOctets(x)
	require.Equal(t, x, sm.Integer64fromOctets(b))
}
