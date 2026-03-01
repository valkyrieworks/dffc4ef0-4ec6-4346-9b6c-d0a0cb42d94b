package status_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	abcicustomerfakes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer/simulations"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	abcistubs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	tpmocks "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool/simulations"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	fakes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate/simulations"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

var (
	successionUUID             = "REDACTED"
	verifyFragmentExtent uint32 = kinds.LedgerFragmentExtentOctets
)

func VerifyExecuteLedger(t *testing.T) {
	app := &verifyApplication{}
	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.Nil(t, err)
	defer delegatePlatform.Halt() //

	status, statusDatastore, _ := createStatus(1, 1)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())

	mp := &tpmocks.Txpool{}
	mp.On("REDACTED").Return()
	mp.On("REDACTED").Return()
	mp.On("REDACTED", mock.Anything).Return(nil)
	mp.On("REDACTED",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(nil)
	ledgerExecute := sm.FreshLedgerHandler(statusDepot, log.VerifyingTracer(), delegatePlatform.Agreement(),
		mp, sm.VoidProofHub{}, ledgerDepot)

	ledger, err := createLedger(status, 1, new(kinds.Endorse))
	require.NoError(t, err)
	bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
	require.NoError(t, err)
	ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}

	status, err = ledgerExecute.ExecuteLedger(status, ledgerUUID, ledger)
	require.Nil(t, err)

	//
	assert.EqualValues(t, 1, status.Edition.Agreement.App, "REDACTED")
}

//
//
//
//
func VerifyCulminateLedgerResolvedFinalEndorse(t *testing.T) {
	app := &verifyApplication{}
	foundationMoment := time.Now()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	status, statusDatastore, privateItems := createStatus(7, 1)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	missingSignature := kinds.FreshExpandedEndorseSignatureMissing()

	verifyScenarios := []struct {
		alias             string
		missingEndorseSignatures map[int]bool
	}{
		{"REDACTED", map[int]bool{}},
		{"REDACTED", map[int]bool{1: true}},
		{"REDACTED", map[int]bool{1: true, 3: true}},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.alias, func(t *testing.T) {
			ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
			incidentpool := &simulations.ProofHub{}
			incidentpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, 0)
			incidentpool.On("REDACTED", mock.Anything, mock.Anything).Return()
			incidentpool.On("REDACTED", mock.Anything).Return(nil)
			mp := &tpmocks.Txpool{}
			mp.On("REDACTED").Return()
			mp.On("REDACTED").Return()
			mp.On("REDACTED", mock.Anything).Return(nil)
			mp.On("REDACTED",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).Return(nil)

			incidentPipeline := kinds.FreshIncidentPipeline()
			require.NoError(t, incidentPipeline.Initiate())

			ledgerExecute := sm.FreshLedgerHandler(statusDepot, log.FreshNooperationTracer(), delegatePlatform.Agreement(), mp, incidentpool, ledgerDepot)
			status, _, finalEndorse, err := createAlsoEndorseValidLedger(status, 1, new(kinds.Endorse), status.FollowingAssessors.Assessors[0].Location, ledgerExecute, privateItems, nil)
			require.NoError(t, err)

			for idx, equalsMissing := range tc.missingEndorseSignatures {
				if equalsMissing {
					finalEndorse.ExpandedNotations[idx] = missingSignature
				}
			}

			//
			ledger, err := createLedger(status, 2, finalEndorse.TowardEndorse())
			require.NoError(t, err)
			bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
			require.NoError(t, err)
			ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}
			_, err = ledgerExecute.ExecuteLedger(status, ledgerUUID, ledger)
			require.NoError(t, err)
			require.True(t, app.FinalMoment.After(foundationMoment))

			//
			for i, v := range app.EndorseBallots {
				_, missing := tc.missingEndorseSignatures[i]
				assert.Equal(t, !missing, v.LedgerUuidMarker != commitchema.LedgerUUIDMarkerMissing)
			}
		})
	}
}

//
func VerifyCulminateLedgerAssessors(t *testing.T) {
	app := &verifyApplication{}
	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	status, statusDatastore, _ := createStatus(2, 2)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	previousDigest := status.FinalLedgerUUID.Digest
	previousFragments := kinds.FragmentAssignHeading{}
	previousLedgerUUID := kinds.LedgerUUID{Digest: previousDigest, FragmentAssignHeading: previousFragments}

	var (
		now        = committime.Now()
		endorseSignature0 = kinds.ExpandedEndorseSignature{
			EndorseSignature: kinds.EndorseSignature{
				LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
				AssessorLocation: status.Assessors.Assessors[0].Location,
				Timestamp:        now,
				Notation:        []byte("REDACTED"),
			},
			Addition:          []byte("REDACTED"),
			AdditionNotation: []byte("REDACTED"),
		}

		endorseSignature1 = kinds.ExpandedEndorseSignature{
			EndorseSignature: kinds.EndorseSignature{
				LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
				AssessorLocation: status.Assessors.Assessors[1].Location,
				Timestamp:        now,
				Notation:        []byte("REDACTED"),
			},
			Addition:          []byte("REDACTED"),
			AdditionNotation: []byte("REDACTED"),
		}
		missingSignature = kinds.FreshExpandedEndorseSignatureMissing()
	)

	verifyScenarios := []struct {
		description                     string
		finalEndorseSignatures           []kinds.ExpandedEndorseSignature
		anticipatedMissingAssessors []int
		mustPossessMoment           bool
	}{
		{"REDACTED", []kinds.ExpandedEndorseSignature{endorseSignature0, endorseSignature1}, []int{}, true},
		{"REDACTED", []kinds.ExpandedEndorseSignature{endorseSignature0, missingSignature}, []int{1}, true},
		{"REDACTED", []kinds.ExpandedEndorseSignature{missingSignature, missingSignature}, []int{0, 1}, false},
	}

	for _, tc := range verifyScenarios {
		finalEndorse := &kinds.ExpandedEndorse{
			Altitude:             1,
			LedgerUUID:            previousLedgerUUID,
			ExpandedNotations: tc.finalEndorseSignatures,
		}

		//
		ledger, err := createLedger(status, 2, finalEndorse.TowardEndorse())
		require.NoError(t, err)

		_, err = sm.InvokeEndorseLedger(delegatePlatform.Agreement(), ledger, log.VerifyingTracer(), statusDepot, 1)
		require.NoError(t, err, tc.description)
		require.True(t,
			!tc.mustPossessMoment ||
				app.FinalMoment.Equal(now) || app.FinalMoment.After(now),
			"REDACTED", tc.description, app.FinalMoment, now,
		)

		//
		ctr := 0
		for i, v := range app.EndorseBallots {
			if ctr < len(tc.anticipatedMissingAssessors) &&
				tc.anticipatedMissingAssessors[ctr] == i {

				assert.Equal(t, v.LedgerUuidMarker, commitchema.LedgerUUIDMarkerMissing)
				ctr++
			} else {
				assert.NotEqual(t, v.LedgerUuidMarker, commitchema.LedgerUUIDMarkerMissing)
			}
		}
	}
}

//
func VerifyCulminateLedgerMalpractice(t *testing.T) {
	app := &verifyApplication{}
	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	status, statusDatastore, privateItems := createStatus(1, 1)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	fallbackProofMoment := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	privateItem := privateItems[status.Assessors.Assessors[0].Location.Text()]
	ledgerUUID := createLedgerUUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
	heading := &kinds.Heading{
		Edition:            strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 1},
		SuccessionUUID:            status.SuccessionUUID,
		Altitude:             10,
		Moment:               fallbackProofMoment,
		FinalLedgerUUID:        ledgerUUID,
		FinalEndorseDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
		DataDigest:           security.CHARArbitraryOctets(tenderminthash.Extent),
		AssessorsDigest:     status.Assessors.Digest(),
		FollowingAssessorsDigest: status.Assessors.Digest(),
		AgreementDigest:      security.CHARArbitraryOctets(tenderminthash.Extent),
		PlatformDigest:            security.CHARArbitraryOctets(tenderminthash.Extent),
		FinalOutcomesDigest:    security.CHARArbitraryOctets(tenderminthash.Extent),
		ProofDigest:       security.CHARArbitraryOctets(tenderminthash.Extent),
		NominatorLocation:    security.CHARArbitraryOctets(security.LocatorExtent),
	}

	//
	dve, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(3, fallbackProofMoment, privateItem, status.SuccessionUUID)
	require.NoError(t, err)
	dve.AssessorPotency = 1000
	agilecustomerattackevidence := &kinds.AgileCustomerOnslaughtProof{
		DiscordantLedger: &kinds.AgileLedger{
			NotatedHeading: &kinds.NotatedHeading{
				Heading: heading,
				Endorse: &kinds.Endorse{
					Altitude:  10,
					LedgerUUID: createLedgerUUID(heading.Digest(), 100, []byte("REDACTED")),
					Notations: []kinds.EndorseSignature{{
						LedgerUUIDMarker:      kinds.LedgerUUIDMarkerVoid,
						AssessorLocation: security.LocatorDigest([]byte("REDACTED")),
						Timestamp:        fallbackProofMoment,
						Notation:        security.CHARArbitraryOctets(kinds.MaximumNotationExtent),
					}},
				},
			},
			AssessorAssign: status.Assessors,
		},
		SharedAltitude:        8,
		TreacherousAssessors: []*kinds.Assessor{status.Assessors.Assessors[0]},
		SumBallotingPotency:    12,
		Timestamp:           fallbackProofMoment,
	}

	ev := []kinds.Proof{dve, agilecustomerattackevidence}

	ifaceMegabyte := []iface.Malpractice{
		{
			Kind:             iface.Malfunctionkind_REPLICATED_BALLOT,
			Altitude:           3,
			Moment:             fallbackProofMoment,
			Assessor:        kinds.Temp2buffer.Assessor(status.Assessors.Assessors[0]),
			SumBallotingPotency: 10,
		},
		{
			Kind:             iface.Malfunctionkind_AGILE_CUSTOMER_ONSLAUGHT,
			Altitude:           8,
			Moment:             fallbackProofMoment,
			Assessor:        kinds.Temp2buffer.Assessor(status.Assessors.Assessors[0]),
			SumBallotingPotency: 12,
		},
	}

	incidentpool := &simulations.ProofHub{}
	incidentpool.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(ev, int64(100))
	incidentpool.On("REDACTED", mock.AnythingOfType("REDACTED"), mock.AnythingOfType("REDACTED")).Return()
	incidentpool.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)
	mp := &tpmocks.Txpool{}
	mp.On("REDACTED").Return()
	mp.On("REDACTED").Return()
	mp.On("REDACTED", mock.Anything).Return(nil)
	mp.On("REDACTED",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(nil)

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())

	ledgerExecute := sm.FreshLedgerHandler(statusDepot, log.VerifyingTracer(), delegatePlatform.Agreement(),
		mp, incidentpool, ledgerDepot)

	ledger, err := createLedger(status, 1, new(kinds.Endorse))
	require.NoError(t, err)
	ledger.Proof = kinds.ProofData{Proof: ev}
	ledger.ProofDigest = ledger.Proof.Digest()
	bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
	require.NoError(t, err)

	ledgerUUID = kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}

	_, err = ledgerExecute.ExecuteLedger(status, ledgerUUID, ledger)
	require.NoError(t, err)

	//
	assert.Equal(t, ifaceMegabyte, app.Malpractice)
}

func VerifyHandleNomination(t *testing.T) {
	const altitude = 2
	txs := verify.CreateNTHTrans(altitude, 10)

	tracer := log.FreshNooperationTracer()
	app := &abcistubs.Platform{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{Condition: iface.Responseexecuteitem_EMBRACE}, nil)

	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	status, statusDatastore, privateItems := createStatus(1, altitude)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	incidentPipeline := kinds.FreshIncidentPipeline()
	err = incidentPipeline.Initiate()
	require.NoError(t, err)

	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		tracer,
		delegatePlatform.Agreement(),
		new(tpmocks.Txpool),
		sm.VoidProofHub{},
		ledgerDepot,
	)

	ledger0, err := createLedger(status, altitude-1, new(kinds.Endorse))
	require.NoError(t, err)
	finalEndorseSignature := []kinds.EndorseSignature{}
	fragmentAssign, err := ledger0.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)
	ledgerUUID := kinds.LedgerUUID{Digest: ledger0.Digest(), FragmentAssignHeading: fragmentAssign.Heading()}
	ballotDetails := []iface.BallotDetails{}
	for _, privateItem := range privateItems {
		pk, err := privateItem.ObtainPublicToken()
		require.NoError(t, err)
		idx, _ := status.Assessors.ObtainViaLocation(pk.Location())
		ballot := kinds.CreateBallotNegativeFailure(t, privateItem, ledger0.SuccessionUUID, idx, altitude-1, 0, 2, ledgerUUID, time.Now())
		location := pk.Location()
		ballotDetails = append(ballotDetails,
			iface.BallotDetails{
				LedgerUuidMarker: commitchema.LedgerUUIDMarkerEndorse,
				Assessor: iface.Assessor{
					Location: location,
					Potency:   1000,
				},
			})
		finalEndorseSignature = append(finalEndorseSignature, ballot.EndorseSignature())
	}

	ledger1, err := createLedger(status, altitude, &kinds.Endorse{
		Altitude:     altitude - 1,
		Notations: finalEndorseSignature,
	})
	require.NoError(t, err)

	ledger1.Txs = txs

	anticipatedRecperpage := &iface.SolicitHandleNomination{
		Txs:         ledger1.Txs.TowardSegmentBelongingOctets(),
		Digest:        ledger1.Digest(),
		Altitude:      ledger1.Altitude,
		Moment:        ledger1.Moment,
		Malpractice: ledger1.Proof.Proof.TowardIface(),
		ItemizedFinalEndorse: iface.EndorseDetails{
			Iteration: 0,
			Ballots: ballotDetails,
		},
		FollowingAssessorsDigest: ledger1.FollowingAssessorsDigest,
		NominatorLocation:    ledger1.NominatorLocation,
	}

	embraceLedger, err := ledgerExecute.HandleNomination(ledger1, status)
	require.NoError(t, err)
	require.True(t, embraceLedger)
	app.AssertExpectations(t)
	app.AssertCalled(t, "REDACTED", context.TODO(), anticipatedRecperpage)
}

func VerifyCertifyAssessorRevisions(t *testing.T) {
	publictoken1 := edwards25519.ProducePrivateToken().PublicToken()
	publictoken2 := edwards25519.ProducePrivateToken().PublicToken()
	pk1, err := cryptocode.PublicTokenTowardSchema(publictoken1)
	assert.NoError(t, err)
	pk2, err := cryptocode.PublicTokenTowardSchema(publictoken2)
	assert.NoError(t, err)

	fallbackAssessorParameters := kinds.AssessorParameters{PublicTokenKinds: []string{kinds.IfacePublicTokenKindEdwards25519}}

	verifyScenarios := []struct {
		alias string

		ifaceRevisions     []iface.AssessorRevise
		assessorParameters kinds.AssessorParameters

		mustFault bool
	}{
		{
			"REDACTED",
			[]iface.AssessorRevise{{PublicToken: pk2, Potency: 20}},
			fallbackAssessorParameters,
			false,
		},
		{
			"REDACTED",
			[]iface.AssessorRevise{{PublicToken: pk1, Potency: 20}},
			fallbackAssessorParameters,
			false,
		},
		{
			"REDACTED",
			[]iface.AssessorRevise{{PublicToken: pk2, Potency: 0}},
			fallbackAssessorParameters,
			false,
		},
		{
			"REDACTED",
			[]iface.AssessorRevise{{PublicToken: pk2, Potency: -100}},
			fallbackAssessorParameters,
			true,
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.alias, func(t *testing.T) {
			err := sm.CertifyAssessorRevisions(tc.ifaceRevisions, tc.assessorParameters)
			if tc.mustFault {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func VerifyReviseAssessors(t *testing.T) {
	publictoken1 := edwards25519.ProducePrivateToken().PublicToken()
	valid1 := kinds.FreshAssessor(publictoken1, 10)
	publictoken2 := edwards25519.ProducePrivateToken().PublicToken()
	valid2 := kinds.FreshAssessor(publictoken2, 20)

	pk, err := cryptocode.PublicTokenTowardSchema(publictoken1)
	require.NoError(t, err)
	pk2, err := cryptocode.PublicTokenTowardSchema(publictoken2)
	require.NoError(t, err)

	verifyScenarios := []struct {
		alias string

		prevailingAssign  *kinds.AssessorAssign
		ifaceRevisions []iface.AssessorRevise

		ensuingAssign *kinds.AssessorAssign
		mustFault    bool
	}{
		{
			"REDACTED",
			kinds.FreshAssessorAssign([]*kinds.Assessor{valid1}),
			[]iface.AssessorRevise{{PublicToken: pk2, Potency: 20}},
			kinds.FreshAssessorAssign([]*kinds.Assessor{valid1, valid2}),
			false,
		},
		{
			"REDACTED",
			kinds.FreshAssessorAssign([]*kinds.Assessor{valid1}),
			[]iface.AssessorRevise{{PublicToken: pk, Potency: 20}},
			kinds.FreshAssessorAssign([]*kinds.Assessor{kinds.FreshAssessor(publictoken1, 20)}),
			false,
		},
		{
			"REDACTED",
			kinds.FreshAssessorAssign([]*kinds.Assessor{valid1, valid2}),
			[]iface.AssessorRevise{{PublicToken: pk2, Potency: 0}},
			kinds.FreshAssessorAssign([]*kinds.Assessor{valid1}),
			false,
		},
		{
			"REDACTED",
			kinds.FreshAssessorAssign([]*kinds.Assessor{valid1}),
			[]iface.AssessorRevise{{PublicToken: pk2, Potency: 0}},
			kinds.FreshAssessorAssign([]*kinds.Assessor{valid1}),
			true,
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.alias, func(t *testing.T) {
			revisions, err := kinds.Buffer2temp.AssessorRevisions(tc.ifaceRevisions)
			assert.NoError(t, err)
			err = tc.prevailingAssign.ReviseUsingModifyAssign(revisions)
			if tc.mustFault {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				require.Equal(t, tc.ensuingAssign.Extent(), tc.prevailingAssign.Extent())

				assert.Equal(t, tc.ensuingAssign.SumBallotingPotency(), tc.prevailingAssign.SumBallotingPotency())

				assert.Equal(t, tc.ensuingAssign.Assessors[0].Location, tc.prevailingAssign.Assessors[0].Location)
				if tc.ensuingAssign.Extent() > 1 {
					assert.Equal(t, tc.ensuingAssign.Assessors[1].Location, tc.prevailingAssign.Assessors[1].Location)
				}
			}
		})
	}
}

//
func VerifyCulminateLedgerAssessorRevisions(t *testing.T) {
	app := &verifyApplication{}
	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	status, statusDatastore, _ := createStatus(1, 1)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	mp := &tpmocks.Txpool{}
	mp.On("REDACTED").Return()
	mp.On("REDACTED").Return()
	mp.On("REDACTED", mock.Anything).Return(nil)
	mp.On("REDACTED",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(nil)
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(kinds.Txs{})

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.VerifyingTracer(),
		delegatePlatform.Agreement(),
		mp,
		sm.VoidProofHub{},
		ledgerDepot,
	)

	incidentPipeline := kinds.FreshIncidentPipeline()
	err = incidentPipeline.Initiate()
	require.NoError(t, err)
	defer incidentPipeline.Halt() //

	ledgerExecute.AssignIncidentChannel(incidentPipeline)

	revisionsUnder, err := incidentPipeline.Listen(
		context.Background(),
		"REDACTED",
		kinds.IncidentInquireAssessorAssignRevisions,
	)
	require.NoError(t, err)

	ledger, err := createLedger(status, 1, new(kinds.Endorse))
	require.NoError(t, err)
	bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
	require.NoError(t, err)
	ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}

	publickey := edwards25519.ProducePrivateToken().PublicToken()
	pk, err := cryptocode.PublicTokenTowardSchema(publickey)
	require.NoError(t, err)
	app.AssessorRevisions = []iface.AssessorRevise{
		{PublicToken: pk, Potency: 10},
	}

	status, err = ledgerExecute.ExecuteLedger(status, ledgerUUID, ledger)
	require.NoError(t, err)
	//
	if assert.Equal(t, status.Assessors.Extent()+1, status.FollowingAssessors.Extent()) {
		idx, _ := status.FollowingAssessors.ObtainViaLocation(publickey.Location())
		if idx < 0 {
			t.Fatalf("REDACTED", publickey.Location(), status.FollowingAssessors)
		}
	}

	//
	select {
	case msg := <-revisionsUnder.Out():
		incident, ok := msg.Data().(kinds.IncidentDataAssessorAssignRevisions)
		require.True(t, ok, "REDACTED", msg.Data())
		if assert.NotEmpty(t, incident.AssessorRevisions) {
			assert.Equal(t, publickey, incident.AssessorRevisions[0].PublicToken)
			assert.EqualValues(t, 10, incident.AssessorRevisions[0].BallotingPotency)
		}
	case <-revisionsUnder.Aborted():
		t.Fatalf("REDACTED", revisionsUnder.Err())
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

//
//
func VerifyCulminateLedgerAssessorRevisionsEnsuingInsideBlankAssign(t *testing.T) {
	app := &verifyApplication{}
	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	status, statusDatastore, _ := createStatus(1, 1)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.VerifyingTracer(),
		delegatePlatform.Agreement(),
		new(tpmocks.Txpool),
		sm.VoidProofHub{},
		ledgerDepot,
	)

	ledger, err := createLedger(status, 1, new(kinds.Endorse))
	require.NoError(t, err)
	bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
	require.NoError(t, err)
	ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}

	vp, err := cryptocode.PublicTokenTowardSchema(status.Assessors.Assessors[0].PublicToken)
	require.NoError(t, err)
	//
	app.AssessorRevisions = []iface.AssessorRevise{
		{PublicToken: vp, Potency: 0},
	}

	assert.NotPanics(t, func() { status, err = ledgerExecute.ExecuteLedger(status, ledgerUUID, ledger) })
	assert.Error(t, err)
	assert.NotEmpty(t, status.FollowingAssessors.Assessors)
}

func VerifyBlankArrangeNomination(t *testing.T) {
	const altitude = 2
	ctx := t.Context()

	app := &iface.FoundationPlatform{}
	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	status, statusDatastore, privateItems := createStatus(1, altitude)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	mp := &tpmocks.Txpool{}
	mp.On("REDACTED").Return()
	mp.On("REDACTED").Return()
	mp.On("REDACTED", mock.Anything).Return(nil)
	mp.On("REDACTED",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(nil)
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(kinds.Txs{})

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.VerifyingTracer(),
		delegatePlatform.Agreement(),
		mp,
		sm.VoidProofHub{},
		ledgerDepot,
	)
	pa, _ := status.Assessors.ObtainViaOrdinal(0)
	endorse, _, err := createSoundEndorse(altitude, kinds.LedgerUUID{}, status.Assessors, privateItems)
	require.NoError(t, err)
	_, err = ledgerExecute.GenerateNominationLedger(ctx, altitude, status, endorse, pa)
	require.NoError(t, err)
}

//
//
func VerifyArrangeNominationTransEveryComprised(t *testing.T) {
	const altitude = 2
	ctx := t.Context()

	status, statusDatastore, privateItems := createStatus(1, altitude)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	incidentpool := &simulations.ProofHub{}
	incidentpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, int64(0))

	txs := verify.CreateNTHTrans(altitude, 10)
	mp := &tpmocks.Txpool{}
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(txs[2:])

	app := &abcistubs.Platform{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{
		Txs: txs.TowardSegmentBelongingOctets(),
	}, nil)
	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.VerifyingTracer(),
		delegatePlatform.Agreement(),
		mp,
		incidentpool,
		ledgerDepot,
	)
	pa, _ := status.Assessors.ObtainViaOrdinal(0)
	endorse, _, err := createSoundEndorse(altitude, kinds.LedgerUUID{}, status.Assessors, privateItems)
	require.NoError(t, err)
	ledger, err := ledgerExecute.GenerateNominationLedger(ctx, altitude, status, endorse, pa)
	require.NoError(t, err)

	for i, tx := range ledger.Txs {
		require.Equal(t, txs[i], tx)
	}

	mp.AssertExpectations(t)
}

//
//
func VerifyArrangeNominationRearrangeTrans(t *testing.T) {
	const altitude = 2
	ctx := t.Context()

	status, statusDatastore, privateItems := createStatus(1, altitude)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	incidentpool := &simulations.ProofHub{}
	incidentpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, int64(0))

	txs := verify.CreateNTHTrans(altitude, 10)
	mp := &tpmocks.Txpool{}
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(txs)

	txs = txs[2:]
	txs = append(txs[len(txs)/2:], txs[:len(txs)/2]...)

	app := &abcistubs.Platform{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{
		Txs: txs.TowardSegmentBelongingOctets(),
	}, nil)

	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.VerifyingTracer(),
		delegatePlatform.Agreement(),
		mp,
		incidentpool,
		ledgerDepot,
	)
	pa, _ := status.Assessors.ObtainViaOrdinal(0)
	endorse, _, err := createSoundEndorse(altitude, kinds.LedgerUUID{}, status.Assessors, privateItems)
	require.NoError(t, err)
	ledger, err := ledgerExecute.GenerateNominationLedger(ctx, altitude, status, endorse, pa)
	require.NoError(t, err)
	for i, tx := range ledger.Txs {
		require.Equal(t, txs[i], tx)
	}

	mp.AssertExpectations(t)
}

//
//
func VerifyArrangeNominationFailureUponExcessivelyMultipleTrans(t *testing.T) {
	const altitude = 2
	ctx := t.Context()

	status, statusDatastore, privateItems := createStatus(1, altitude)
	//
	status.AgreementSettings.Ledger.MaximumOctets = 60 * 1024
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	incidentpool := &simulations.ProofHub{}
	incidentpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, int64(0))

	const nthAssessors = 1
	var octetsEveryTransfer int64 = 3
	maximumDataOctets := kinds.MaximumDataOctets(status.AgreementSettings.Ledger.MaximumOctets, 0, nthAssessors)
	txs := verify.CreateNTHTrans(altitude, maximumDataOctets/octetsEveryTransfer+2) //
	mp := &tpmocks.Txpool{}
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(txs)

	app := &abcistubs.Platform{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{
		Txs: txs.TowardSegmentBelongingOctets(),
	}, nil)

	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.FreshNooperationTracer(),
		delegatePlatform.Agreement(),
		mp,
		incidentpool,
		ledgerDepot,
	)
	pa, _ := status.Assessors.ObtainViaOrdinal(0)
	endorse, _, err := createSoundEndorse(altitude, kinds.LedgerUUID{}, status.Assessors, privateItems)
	require.NoError(t, err)
	ledger, err := ledgerExecute.GenerateNominationLedger(ctx, altitude, status, endorse, pa)
	require.Nil(t, ledger)
	require.ErrorContains(t, err, "REDACTED")

	mp.AssertExpectations(t)
}

//
//
//
func VerifyArrangeNominationTallyEncodingMargin(t *testing.T) {
	const altitude = 2
	ctx := t.Context()

	status, statusDatastore, privateItems := createStatus(1, altitude)
	//
	var octetsEveryTransfer int64 = 4
	const nthAssessors = 1
	unDataExtent := 5000 - kinds.MaximumDataOctets(5000, 0, nthAssessors)
	status.AgreementSettings.Ledger.MaximumOctets = octetsEveryTransfer*1024 + unDataExtent
	maximumDataOctets := kinds.MaximumDataOctets(status.AgreementSettings.Ledger.MaximumOctets, 0, nthAssessors)

	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	incidentpool := &simulations.ProofHub{}
	incidentpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, int64(0))

	txs := verify.CreateNTHTrans(altitude, maximumDataOctets/octetsEveryTransfer)
	mp := &tpmocks.Txpool{}
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(txs)

	app := &abcistubs.Platform{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{
		Txs: txs.TowardSegmentBelongingOctets(),
	}, nil)

	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.FreshNooperationTracer(),
		delegatePlatform.Agreement(),
		mp,
		incidentpool,
		ledgerDepot,
	)
	pa, _ := status.Assessors.ObtainViaOrdinal(0)
	endorse, _, err := createSoundEndorse(altitude, kinds.LedgerUUID{}, status.Assessors, privateItems)
	require.NoError(t, err)
	ledger, err := ledgerExecute.GenerateNominationLedger(ctx, altitude, status, endorse, pa)
	require.Nil(t, ledger)
	require.ErrorContains(t, err, "REDACTED")

	mp.AssertExpectations(t)
}

//
//
func VerifyArrangeNominationFailureUponArrangeNominationFailure(t *testing.T) {
	const altitude = 2
	ctx := t.Context()

	status, statusDatastore, privateItems := createStatus(1, altitude)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})

	incidentpool := &simulations.ProofHub{}
	incidentpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, int64(0))

	txs := verify.CreateNTHTrans(altitude, 10)
	mp := &tpmocks.Txpool{}
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(txs)

	cm := &abcicustomerfakes.Customer{}
	cm.On("REDACTED", mock.Anything).Return()
	cm.On("REDACTED").Return(nil)
	cm.On("REDACTED").Return(nil)
	cm.On("REDACTED", mock.Anything, mock.Anything).Return(nil, errors.New("REDACTED")).Once()
	cm.On("REDACTED").Return(nil)
	cc := &fakes.CustomerOriginator{}
	cc.On("REDACTED").Return(cm, nil)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	require.NoError(t, err)
	defer delegatePlatform.Halt() //

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.FreshNooperationTracer(),
		delegatePlatform.Agreement(),
		mp,
		incidentpool,
		ledgerDepot,
	)
	pa, _ := status.Assessors.ObtainViaOrdinal(0)
	endorse, _, err := createSoundEndorse(altitude, kinds.LedgerUUID{}, status.Assessors, privateItems)
	require.NoError(t, err)
	ledger, err := ledgerExecute.GenerateNominationLedger(ctx, altitude, status, endorse, pa)
	require.Nil(t, ledger)
	require.ErrorContains(t, err, "REDACTED")

	mp.AssertExpectations(t)
}

//
//
//
func VerifyGenerateNominationMissingBallotAdditions(t *testing.T) {
	for _, verifyInstance := range []struct {
		alias string

		//
		altitude int64

		//
		additionActivateAltitude int64
		anticipateAlarm           bool
	}{
		{
			alias:                  "REDACTED",
			altitude:                3,
			additionActivateAltitude: 2,
			anticipateAlarm:           true,
		},
		{
			alias:                  "REDACTED",
			altitude:                3,
			additionActivateAltitude: 3,
			anticipateAlarm:           false,
		},
		{
			alias:                  "REDACTED",
			altitude:                3,
			additionActivateAltitude: 0,
			anticipateAlarm:           false,
		},
		{
			alias:                  "REDACTED",
			altitude:                3,
			additionActivateAltitude: 4,
			anticipateAlarm:           false,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			ctx := t.Context()

			app := abcistubs.FreshPlatform(t)
			if !verifyInstance.anticipateAlarm {
				app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil)
			}
			cc := delegate.FreshRegionalCustomerOriginator(app)
			delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
			err := delegatePlatform.Initiate()
			require.NoError(t, err)

			status, statusDatastore, privateItems := createStatus(1, int(verifyInstance.altitude-1))
			statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
				EjectIfaceReplies: false,
			})
			status.AgreementSettings.Iface.BallotAdditionsActivateAltitude = verifyInstance.additionActivateAltitude
			mp := &tpmocks.Txpool{}
			mp.On("REDACTED").Return()
			mp.On("REDACTED").Return()
			mp.On("REDACTED", mock.Anything).Return(nil)
			mp.On("REDACTED",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything).Return(nil)
			mp.On("REDACTED", mock.Anything, mock.Anything).Return(kinds.Txs{})

			ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())
			ledgerExecute := sm.FreshLedgerHandler(
				statusDepot,
				log.FreshNooperationTracer(),
				delegatePlatform.Agreement(),
				mp,
				sm.VoidProofHub{},
				ledgerDepot,
			)
			ledger, err := createLedger(status, verifyInstance.altitude, new(kinds.Endorse))
			require.NoError(t, err)

			bps, err := ledger.CreateFragmentAssign(verifyFragmentExtent)
			require.NoError(t, err)
			ledgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: bps.Heading()}
			pa, _ := status.Assessors.ObtainViaOrdinal(0)
			finalEndorse, _, _ := createSoundEndorse(verifyInstance.altitude-1, ledgerUUID, status.Assessors, privateItems)
			removeNotations(finalEndorse)
			if verifyInstance.anticipateAlarm {
				require.Panics(t, func() {
					_, err := ledgerExecute.GenerateNominationLedger(ctx, verifyInstance.altitude, status, finalEndorse, pa)
					require.NoError(t, err)
				})
			} else {
				_, err = ledgerExecute.GenerateNominationLedger(ctx, verifyInstance.altitude, status, finalEndorse, pa)
				require.NoError(t, err)
			}
		})
	}
}

func removeNotations(ec *kinds.ExpandedEndorse) {
	for i, endorseSignature := range ec.ExpandedNotations {
		endorseSignature.Addition = nil
		endorseSignature.AdditionNotation = nil
		ec.ExpandedNotations[i] = endorseSignature
	}
}

func createLedgerUUID(digest []byte, fragmentAssignExtent uint32, fragmentAssignDigest []byte) kinds.LedgerUUID {
	var (
		h   = make([]byte, tenderminthash.Extent)
		psH = make([]byte, tenderminthash.Extent)
	)
	copy(h, digest)
	copy(psH, fragmentAssignDigest)
	return kinds.LedgerUUID{
		Digest: h,
		FragmentAssignHeading: kinds.FragmentAssignHeading{
			Sum: fragmentAssignExtent,
			Digest:  psH,
		},
	}
}
