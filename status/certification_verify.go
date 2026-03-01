package status_test

import (
	"testing"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	tpmocks "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool/simulations"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

const certificationVerifiesHaltAltitude int64 = 10

func VerifyCertifyLedgerHeadline(t *testing.T) {
	delegatePlatform := freshVerifyApplication()
	require.NoError(t, delegatePlatform.Initiate())
	defer delegatePlatform.Halt() //

	status, statusDatastore, privateItems := createStatus(3, 1)
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

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())

	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.VerifyingTracer(),
		delegatePlatform.Agreement(),
		mp,
		sm.VoidProofHub{},
		ledgerDepot,
	)
	finalEndorse := &kinds.Endorse{}
	var finalAddnEndorse *kinds.ExpandedEndorse

	//
	incorrectDigest := tenderminthash.Sum([]byte("REDACTED"))
	incorrectEdition1 := status.Edition.Agreement
	incorrectEdition1.Ledger += 2
	incorrectEdition2 := status.Edition.Agreement
	incorrectEdition2.App += 2

	//
	verifyScenarios := []struct {
		alias          string
		distortLedger func(ledger *kinds.Ledger)
	}{
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.Edition = incorrectEdition1 }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.Edition = incorrectEdition2 }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.SuccessionUUID = "REDACTED" }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.Altitude += 10 }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.Moment = ledger.Moment.Add(-time.Second * 1) }},

		{"REDACTED", func(ledger *kinds.Ledger) { ledger.FinalLedgerUUID.FragmentAssignHeading.Sum += 10 }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.FinalEndorseDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.DataDigest = incorrectDigest }},

		{"REDACTED", func(ledger *kinds.Ledger) { ledger.AssessorsDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.FollowingAssessorsDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.AgreementDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.PlatformDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.FinalOutcomesDigest = incorrectDigest }},

		{"REDACTED", func(ledger *kinds.Ledger) { ledger.ProofDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.NominatorLocation = edwards25519.ProducePrivateToken().PublicToken().Location() }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.NominatorLocation = []byte("REDACTED") }},
	}

	//
	for altitude := int64(1); altitude < certificationVerifiesHaltAltitude; altitude++ {
		/**
s
*/
		for _, tc := range verifyScenarios {
			ledger, err := createLedger(status, altitude, finalEndorse)
			require.NoError(t, err)
			tc.distortLedger(ledger)
			err = ledgerExecute.CertifyLedger(status, ledger)
			require.Error(t, err, tc.alias)
		}

		/**
s
*/
		var err error
		status, _, finalAddnEndorse, err = createAlsoEndorseValidLedger(
			status, altitude, finalEndorse, status.Assessors.ObtainNominator().Location, ledgerExecute, privateItems, nil)
		require.NoError(t, err, "REDACTED", altitude)
		finalEndorse = finalAddnEndorse.TowardEndorse()
	}

	followingAltitude := certificationVerifiesHaltAltitude
	ledger, err := createLedger(status, followingAltitude, finalEndorse)
	require.NoError(t, err)
	status.PrimaryAltitude = followingAltitude + 1
	err = ledgerExecute.CertifyLedger(status, ledger)
	require.Error(t, err, "REDACTED")
	assert.Contains(t, err.Error(), "REDACTED")
}

func VerifyCertifyLedgerEndorse(t *testing.T) {
	delegatePlatform := freshVerifyApplication()
	require.NoError(t, delegatePlatform.Initiate())
	defer delegatePlatform.Halt() //

	status, statusDatastore, privateItems := createStatus(1, 1)
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

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())

	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.VerifyingTracer(),
		delegatePlatform.Agreement(),
		mp,
		sm.VoidProofHub{},
		ledgerDepot,
	)
	finalEndorse := &kinds.Endorse{}
	var finalAddnEndorse *kinds.ExpandedEndorse
	incorrectSignaturesEndorse := &kinds.Endorse{Altitude: 1}
	flawedPrivateItem := kinds.FreshSimulatePRV()

	for altitude := int64(1); altitude < certificationVerifiesHaltAltitude; altitude++ {
		nominatorLocation := status.Assessors.ObtainNominator().Location
		if altitude > 1 {
			/**
e
*/
			//
			idx, _ := status.Assessors.ObtainViaLocation(nominatorLocation)
			incorrectAltitudeBallot := kinds.CreateBallotNegativeFailure(
				t,
				privateItems[nominatorLocation.Text()],
				successionUUID,
				idx,
				altitude,
				0,
				2,
				status.FinalLedgerUUID,
				time.Now(),
			)
			incorrectAltitudeEndorse := &kinds.Endorse{
				Altitude:     incorrectAltitudeBallot.Altitude,
				Iteration:      incorrectAltitudeBallot.Iteration,
				LedgerUUID:    status.FinalLedgerUUID,
				Notations: []kinds.EndorseSignature{incorrectAltitudeBallot.EndorseSignature()},
			}
			ledger, err := createLedger(status, altitude, incorrectAltitudeEndorse)
			require.NoError(t, err)
			err = ledgerExecute.CertifyLedger(status, ledger)
			_, equalsFaultUnfitEndorseAltitude := err.(faults.FaultUnfitEndorseAltitude)
			require.True(t, equalsFaultUnfitEndorseAltitude, "REDACTED", altitude, err)

			/**
)
*/
			_, err = createLedger(status, altitude, incorrectSignaturesEndorse)
			require.Error(t, err)
			require.ErrorContains(t, err, "REDACTED")
		}

		/**
s
*/
		var err error
		var ledgerUUID kinds.LedgerUUID
		status, ledgerUUID, finalAddnEndorse, err = createAlsoEndorseValidLedger(
			status,
			altitude,
			finalEndorse,
			nominatorLocation,
			ledgerExecute,
			privateItems,
			nil,
		)
		require.NoError(t, err, "REDACTED", altitude)
		finalEndorse = finalAddnEndorse.TowardEndorse()

		/**
t
*/
		idx, _ := status.Assessors.ObtainViaLocation(nominatorLocation)
		validBallot := kinds.CreateBallotNegativeFailure(
			t,
			privateItems[nominatorLocation.Text()],
			successionUUID,
			idx,
			altitude,
			0,
			commitchema.PreendorseKind,
			ledgerUUID,
			time.Now(),
		)

		bpvaluePublicToken, err := flawedPrivateItem.ObtainPublicToken()
		require.NoError(t, err)

		flawedBallot := &kinds.Ballot{
			AssessorLocation: bpvaluePublicToken.Location(),
			AssessorOrdinal:   0,
			Altitude:           altitude,
			Iteration:            0,
			Timestamp:        committime.Now(),
			Kind:             commitchema.PreendorseKind,
			LedgerUUID:          ledgerUUID,
		}

		g := validBallot.TowardSchema()
		b := flawedBallot.TowardSchema()

		err = flawedPrivateItem.AttestBallot(successionUUID, g)
		require.NoError(t, err, "REDACTED", altitude)
		err = flawedPrivateItem.AttestBallot(successionUUID, b)
		require.NoError(t, err, "REDACTED", altitude)

		validBallot.Notation, flawedBallot.Notation = g.Notation, b.Notation

		incorrectSignaturesEndorse = &kinds.Endorse{
			Altitude:     validBallot.Altitude,
			Iteration:      validBallot.Iteration,
			LedgerUUID:    ledgerUUID,
			Notations: []kinds.EndorseSignature{validBallot.EndorseSignature(), flawedBallot.EndorseSignature()},
		}
	}
}

func VerifyCertifyLedgerProof(t *testing.T) {
	delegatePlatform := freshVerifyApplication()
	require.NoError(t, delegatePlatform.Initiate())
	defer delegatePlatform.Halt() //

	status, statusDatastore, privateItems := createStatus(4, 1)
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	fallbackProofMoment := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

	incidentpool := &simulations.ProofHub{}
	incidentpool.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)
	incidentpool.On("REDACTED", mock.AnythingOfType("REDACTED"), mock.AnythingOfType("REDACTED")).Return()
	incidentpool.On("REDACTED", mock.AnythingOfType("REDACTED"), mock.AnythingOfType("REDACTED")).Return(
		[]iface.Malpractice{})

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
	status.AgreementSettings.Proof.MaximumOctets = 1000
	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())

	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.VerifyingTracer(),
		delegatePlatform.Agreement(),
		mp,
		incidentpool,
		ledgerDepot,
	)
	finalEndorse := &kinds.Endorse{}
	var finalAddnEndorse *kinds.ExpandedEndorse

	for altitude := int64(1); altitude < certificationVerifiesHaltAltitude; altitude++ {
		nominatorLocation := status.Assessors.ObtainNominator().Location
		maximumOctetsProof := status.AgreementSettings.Proof.MaximumOctets
		if altitude > 1 {
			/**
s
*/
			proof := make([]kinds.Proof, 0)
			var prevailingOctets int64
			//
			for prevailingOctets <= maximumOctetsProof {
				freshOccurence, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(altitude, time.Now(),
					privateItems[nominatorLocation.Text()], successionUUID)
				require.NoError(t, err)
				proof = append(proof, freshOccurence)
				prevailingOctets += int64(len(freshOccurence.Octets()))
			}
			ledger, err := status.CreateLedger(altitude, verify.CreateNTHTrans(altitude, 10), finalEndorse, proof, nominatorLocation)
			require.NoError(t, err)

			err = ledgerExecute.CertifyLedger(status, ledger)
			if assert.Error(t, err) {
				_, ok := err.(*kinds.FaultProofOverrun)
				require.True(t, ok, "REDACTED", altitude, err)
			}
		}

		/**
s
*/
		proof := make([]kinds.Proof, 0)
		var prevailingOctets int64
		//
		for {
			freshOccurence, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(altitude, fallbackProofMoment,
				privateItems[nominatorLocation.Text()], successionUUID)
			require.NoError(t, err)
			prevailingOctets += int64(len(freshOccurence.Octets()))
			if prevailingOctets >= maximumOctetsProof {
				break
			}
			proof = append(proof, freshOccurence)
		}

		var err error
		status, _, finalAddnEndorse, err = createAlsoEndorseValidLedger(
			status,
			altitude,
			finalEndorse,
			nominatorLocation,
			ledgerExecute,
			privateItems,
			proof,
		)
		require.NoError(t, err, "REDACTED", altitude)
		finalEndorse = finalAddnEndorse.TowardEndorse()

	}
}

func VerifyCertifyLedgerMoment(t *testing.T) {
	delegatePlatform := freshVerifyApplication()
	require.NoError(t, delegatePlatform.Initiate())
	defer delegatePlatform.Halt() //

	status, statusDatastore, privateItems := createStatus(3, 1)
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

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())

	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.VerifyingTracer(),
		delegatePlatform.Agreement(),
		mp,
		sm.VoidProofHub{},
		ledgerDepot,
	)
	finalEndorse := &kinds.Endorse{}
	var finalAddnEndorse *kinds.ExpandedEndorse

	//
	for altitude := int64(1); altitude < 3; altitude++ {
		var err error
		status, _, finalAddnEndorse, err = createAlsoEndorseValidLedger(
			status, altitude, finalEndorse, status.Assessors.ObtainNominator().Location, ledgerExecute, privateItems, nil)
		require.NoError(t, err, "REDACTED", altitude)
		finalEndorse = finalAddnEndorse.TowardEndorse()
	}

	t.Run("REDACTED", func(t *testing.T) {
		altitude := int64(3)
		ledger, err := createLedger(status, altitude, finalEndorse)
		require.NoError(t, err)

		//
		ledger.Moment = ledger.Moment.Add(-time.Millisecond * 10)
		err = ledgerExecute.CertifyLedger(status, ledger)

		require.ErrorContains(t, err, "REDACTED")
	})

	t.Run("REDACTED", func(t *testing.T) {
		altitude := int64(3)
		ledger, err := createLedger(status, altitude, finalEndorse)
		require.NoError(t, err)
		//
		ledger.Moment = ledger.Moment.Add(time.Second)
		err = ledgerExecute.CertifyLedger(status, ledger)
		require.Error(t, err)
		require.Contains(t, err.Error(), "REDACTED")
	})

	t.Run("REDACTED", func(t *testing.T) {
		altitude := int64(3)
		ledger, err := createLedger(status, altitude, finalEndorse)
		require.NoError(t, err)
		err = ledgerExecute.CertifyLedger(status, ledger)
		require.NoError(t, err)
	})
}

func VerifyCertifyLedgerUnfitEndorse(t *testing.T) {
	delegatePlatform := freshVerifyApplication()
	require.NoError(t, delegatePlatform.Initiate())
	defer delegatePlatform.Halt() //

	status, statusDatastore, privateItems := createStatus(3, 1)
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

	ledgerDepot := depot.FreshLedgerDepot(dbm.FreshMemoryDatastore())

	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		log.VerifyingTracer(),
		delegatePlatform.Agreement(),
		mp,
		sm.VoidProofHub{},
		ledgerDepot,
	)
	finalEndorse := &kinds.Endorse{}
	var finalAddnEndorse *kinds.ExpandedEndorse

	//
	for altitude := int64(1); altitude < 3; altitude++ {
		var err error
		status, _, finalAddnEndorse, err = createAlsoEndorseValidLedger(
			status, altitude, finalEndorse, status.Assessors.ObtainNominator().Location, ledgerExecute, privateItems, nil)
		require.NoError(t, err, "REDACTED", altitude)
		finalEndorse = finalAddnEndorse.TowardEndorse()
	}

	t.Run("REDACTED", func(t *testing.T) {
		altitude := int64(3)

		//
		unfamiliarItem := edwards25519.ProducePrivateToken()
		now := time.Now()

		unfitEndorse := &kinds.Endorse{
			Altitude:  altitude - 1,
			Iteration:   0,
			LedgerUUID: status.FinalLedgerUUID,
			Notations: []kinds.EndorseSignature{
				{
					LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
					AssessorLocation: unfamiliarItem.PublicToken().Location(),
					Timestamp:        now,
					Notation:        []byte("REDACTED"),
				},
			},
		}

		_, err := createLedger(status, altitude, unfitEndorse)
		require.Error(t, err)
		require.Contains(t, err.Error(), "REDACTED")
	})
}
