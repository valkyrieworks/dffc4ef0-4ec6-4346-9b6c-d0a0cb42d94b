package status_test

import (
	"testing"
	"time"

	"github.com/valkyrieworks/kinds/faults"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"
	txpoolsims "github.com/valkyrieworks/txpool/simulations"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/simulations"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

const verificationVerifiesHaltLevel int64 = 10

func VerifyCertifyLedgerHeading(t *testing.T) {
	gatewayApplication := newVerifyApplication()
	require.NoError(t, gatewayApplication.Begin())
	defer gatewayApplication.Halt() //

	status, statusStore, privateValues := createStatus(3, 1)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	mp := &txpoolsims.Txpool{}
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

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())

	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.VerifyingTracer(),
		gatewayApplication.Agreement(),
		mp,
		sm.EmptyProofDepository{},
		ledgerDepot,
	)
	finalEndorse := &kinds.Endorse{}
	var finalExtensionEndorse *kinds.ExpandedEndorse

	//
	incorrectDigest := comethash.Sum([]byte("REDACTED"))
	incorrectEdition1 := status.Release.Agreement
	incorrectEdition1.Ledger += 2
	incorrectEdition2 := status.Release.Agreement
	incorrectEdition2.App += 2

	//
	verifyScenarios := []struct {
		label          string
		distortLedger func(ledger *kinds.Ledger)
	}{
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.Release = incorrectEdition1 }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.Release = incorrectEdition2 }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.LedgerUID = "REDACTED" }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.Level += 10 }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.Time = ledger.Time.Add(-time.Second * 1) }},

		{"REDACTED", func(ledger *kinds.Ledger) { ledger.FinalLedgerUID.SegmentAssignHeading.Sum += 10 }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.FinalEndorseDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.DataDigest = incorrectDigest }},

		{"REDACTED", func(ledger *kinds.Ledger) { ledger.RatifiersDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.FollowingRatifiersDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.AgreementDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.ApplicationDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.FinalOutcomesDigest = incorrectDigest }},

		{"REDACTED", func(ledger *kinds.Ledger) { ledger.ProofDigest = incorrectDigest }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.RecommenderLocation = ed25519.GeneratePrivateKey().PublicKey().Location() }},
		{"REDACTED", func(ledger *kinds.Ledger) { ledger.RecommenderLocation = []byte("REDACTED") }},
	}

	//
	for level := int64(1); level < verificationVerifiesHaltLevel; level++ {
		/**
s
*/
		for _, tc := range verifyScenarios {
			ledger, err := createLedger(status, level, finalEndorse)
			require.NoError(t, err)
			tc.distortLedger(ledger)
			err = ledgerExecute.CertifyLedger(status, ledger)
			require.Error(t, err, tc.label)
		}

		/**
s
*/
		var err error
		status, _, finalExtensionEndorse, err = createAndEndorseValidLedger(
			status, level, finalEndorse, status.Ratifiers.FetchRecommender().Location, ledgerExecute, privateValues, nil)
		require.NoError(t, err, "REDACTED", level)
		finalEndorse = finalExtensionEndorse.ToEndorse()
	}

	followingLevel := verificationVerifiesHaltLevel
	ledger, err := createLedger(status, followingLevel, finalEndorse)
	require.NoError(t, err)
	status.PrimaryLevel = followingLevel + 1
	err = ledgerExecute.CertifyLedger(status, ledger)
	require.Error(t, err, "REDACTED")
	assert.Contains(t, err.Error(), "REDACTED")
}

func VerifyCertifyLedgerEndorse(t *testing.T) {
	gatewayApplication := newVerifyApplication()
	require.NoError(t, gatewayApplication.Begin())
	defer gatewayApplication.Halt() //

	status, statusStore, privateValues := createStatus(1, 1)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	mp := &txpoolsims.Txpool{}
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

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())

	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.VerifyingTracer(),
		gatewayApplication.Agreement(),
		mp,
		sm.EmptyProofDepository{},
		ledgerDepot,
	)
	finalEndorse := &kinds.Endorse{}
	var finalExtensionEndorse *kinds.ExpandedEndorse
	incorrectAutographsEndorse := &kinds.Endorse{Level: 1}
	flawedPrivateValue := kinds.NewEmulatePV()

	for level := int64(1); level < verificationVerifiesHaltLevel; level++ {
		recommenderAddress := status.Ratifiers.FetchRecommender().Location
		if level > 1 {
			/**
e
*/
			//
			idx, _ := status.Ratifiers.FetchByLocation(recommenderAddress)
			incorrectLevelBallot := kinds.CreateBallotNoFault(
				t,
				privateValues[recommenderAddress.String()],
				ledgerUID,
				idx,
				level,
				0,
				2,
				status.FinalLedgerUID,
				time.Now(),
			)
			incorrectLevelEndorse := &kinds.Endorse{
				Level:     incorrectLevelBallot.Level,
				Cycle:      incorrectLevelBallot.Cycle,
				LedgerUID:    status.FinalLedgerUID,
				Endorsements: []kinds.EndorseSignature{incorrectLevelBallot.EndorseSignature()},
			}
			ledger, err := createLedger(status, level, incorrectLevelEndorse)
			require.NoError(t, err)
			err = ledgerExecute.CertifyLedger(status, ledger)
			_, isErrCorruptEndorseLevel := err.(faults.ErrCorruptEndorseLevel)
			require.True(t, isErrCorruptEndorseLevel, "REDACTED", level, err)

			/**
)
*/
			_, err = createLedger(status, level, incorrectAutographsEndorse)
			require.Error(t, err)
			require.ErrorContains(t, err, "REDACTED")
		}

		/**
s
*/
		var err error
		var ledgerUID kinds.LedgerUID
		status, ledgerUID, finalExtensionEndorse, err = createAndEndorseValidLedger(
			status,
			level,
			finalEndorse,
			recommenderAddress,
			ledgerExecute,
			privateValues,
			nil,
		)
		require.NoError(t, err, "REDACTED", level)
		finalEndorse = finalExtensionEndorse.ToEndorse()

		/**
t
*/
		idx, _ := status.Ratifiers.FetchByLocation(recommenderAddress)
		validBallot := kinds.CreateBallotNoFault(
			t,
			privateValues[recommenderAddress.String()],
			ledgerUID,
			idx,
			level,
			0,
			engineproto.PreendorseKind,
			ledgerUID,
			time.Now(),
		)

		lbvPublicKey, err := flawedPrivateValue.FetchPublicKey()
		require.NoError(t, err)

		flawedBallot := &kinds.Ballot{
			RatifierLocation: lbvPublicKey.Location(),
			RatifierOrdinal:   0,
			Level:           level,
			Cycle:            0,
			Timestamp:        engineclock.Now(),
			Kind:             engineproto.PreendorseKind,
			LedgerUID:          ledgerUID,
		}

		g := validBallot.ToSchema()
		b := flawedBallot.ToSchema()

		err = flawedPrivateValue.AttestBallot(ledgerUID, g)
		require.NoError(t, err, "REDACTED", level)
		err = flawedPrivateValue.AttestBallot(ledgerUID, b)
		require.NoError(t, err, "REDACTED", level)

		validBallot.Autograph, flawedBallot.Autograph = g.Autograph, b.Autograph

		incorrectAutographsEndorse = &kinds.Endorse{
			Level:     validBallot.Level,
			Cycle:      validBallot.Cycle,
			LedgerUID:    ledgerUID,
			Endorsements: []kinds.EndorseSignature{validBallot.EndorseSignature(), flawedBallot.EndorseSignature()},
		}
	}
}

func VerifyCertifyLedgerProof(t *testing.T) {
	gatewayApplication := newVerifyApplication()
	require.NoError(t, gatewayApplication.Begin())
	defer gatewayApplication.Halt() //

	status, statusStore, privateValues := createStatus(4, 1)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	standardProofTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

	eventpool := &simulations.ProofDepository{}
	eventpool.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)
	eventpool.On("REDACTED", mock.AnythingOfType("REDACTED"), mock.AnythingOfType("REDACTED")).Return()
	eventpool.On("REDACTED", mock.AnythingOfType("REDACTED"), mock.AnythingOfType("REDACTED")).Return(
		[]iface.Malpractice{})

	mp := &txpoolsims.Txpool{}
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
	status.AgreementOptions.Proof.MaximumOctets = 1000
	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())

	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.VerifyingTracer(),
		gatewayApplication.Agreement(),
		mp,
		eventpool,
		ledgerDepot,
	)
	finalEndorse := &kinds.Endorse{}
	var finalExtensionEndorse *kinds.ExpandedEndorse

	for level := int64(1); level < verificationVerifiesHaltLevel; level++ {
		recommenderAddress := status.Ratifiers.FetchRecommender().Location
		maximumOctetsProof := status.AgreementOptions.Proof.MaximumOctets
		if level > 1 {
			/**
s
*/
			proof := make([]kinds.Proof, 0)
			var ongoingOctets int64
			//
			for ongoingOctets <= maximumOctetsProof {
				newEvt, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(level, time.Now(),
					privateValues[recommenderAddress.String()], ledgerUID)
				require.NoError(t, err)
				proof = append(proof, newEvt)
				ongoingOctets += int64(len(newEvt.Octets()))
			}
			ledger, err := status.CreateLedger(level, verify.CreateNTrans(level, 10), finalEndorse, proof, recommenderAddress)
			require.NoError(t, err)

			err = ledgerExecute.CertifyLedger(status, ledger)
			if assert.Error(t, err) {
				_, ok := err.(*kinds.ErrProofOverload)
				require.True(t, ok, "REDACTED", level, err)
			}
		}

		/**
s
*/
		proof := make([]kinds.Proof, 0)
		var ongoingOctets int64
		//
		for {
			newEvt, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(level, standardProofTime,
				privateValues[recommenderAddress.String()], ledgerUID)
			require.NoError(t, err)
			ongoingOctets += int64(len(newEvt.Octets()))
			if ongoingOctets >= maximumOctetsProof {
				break
			}
			proof = append(proof, newEvt)
		}

		var err error
		status, _, finalExtensionEndorse, err = createAndEndorseValidLedger(
			status,
			level,
			finalEndorse,
			recommenderAddress,
			ledgerExecute,
			privateValues,
			proof,
		)
		require.NoError(t, err, "REDACTED", level)
		finalEndorse = finalExtensionEndorse.ToEndorse()

	}
}

func VerifyCertifyLedgerTime(t *testing.T) {
	gatewayApplication := newVerifyApplication()
	require.NoError(t, gatewayApplication.Begin())
	defer gatewayApplication.Halt() //

	status, statusStore, privateValues := createStatus(3, 1)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	mp := &txpoolsims.Txpool{}
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

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())

	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.VerifyingTracer(),
		gatewayApplication.Agreement(),
		mp,
		sm.EmptyProofDepository{},
		ledgerDepot,
	)
	finalEndorse := &kinds.Endorse{}
	var finalExtensionEndorse *kinds.ExpandedEndorse

	//
	for level := int64(1); level < 3; level++ {
		var err error
		status, _, finalExtensionEndorse, err = createAndEndorseValidLedger(
			status, level, finalEndorse, status.Ratifiers.FetchRecommender().Location, ledgerExecute, privateValues, nil)
		require.NoError(t, err, "REDACTED", level)
		finalEndorse = finalExtensionEndorse.ToEndorse()
	}

	t.Run("REDACTED", func(t *testing.T) {
		level := int64(3)
		ledger, err := createLedger(status, level, finalEndorse)
		require.NoError(t, err)

		//
		ledger.Time = ledger.Time.Add(-time.Millisecond * 10)
		err = ledgerExecute.CertifyLedger(status, ledger)

		require.ErrorContains(t, err, "REDACTED")
	})

	t.Run("REDACTED", func(t *testing.T) {
		level := int64(3)
		ledger, err := createLedger(status, level, finalEndorse)
		require.NoError(t, err)
		//
		ledger.Time = ledger.Time.Add(time.Second)
		err = ledgerExecute.CertifyLedger(status, ledger)
		require.Error(t, err)
		require.Contains(t, err.Error(), "REDACTED")
	})

	t.Run("REDACTED", func(t *testing.T) {
		level := int64(3)
		ledger, err := createLedger(status, level, finalEndorse)
		require.NoError(t, err)
		err = ledgerExecute.CertifyLedger(status, ledger)
		require.NoError(t, err)
	})
}

func VerifyCertifyLedgerCorruptEndorse(t *testing.T) {
	gatewayApplication := newVerifyApplication()
	require.NoError(t, gatewayApplication.Begin())
	defer gatewayApplication.Halt() //

	status, statusStore, privateValues := createStatus(3, 1)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	mp := &txpoolsims.Txpool{}
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

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())

	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.VerifyingTracer(),
		gatewayApplication.Agreement(),
		mp,
		sm.EmptyProofDepository{},
		ledgerDepot,
	)
	finalEndorse := &kinds.Endorse{}
	var finalExtensionEndorse *kinds.ExpandedEndorse

	//
	for level := int64(1); level < 3; level++ {
		var err error
		status, _, finalExtensionEndorse, err = createAndEndorseValidLedger(
			status, level, finalEndorse, status.Ratifiers.FetchRecommender().Location, ledgerExecute, privateValues, nil)
		require.NoError(t, err, "REDACTED", level)
		finalEndorse = finalExtensionEndorse.ToEndorse()
	}

	t.Run("REDACTED", func(t *testing.T) {
		level := int64(3)

		//
		unclearValue := ed25519.GeneratePrivateKey()
		now := time.Now()

		corruptEndorse := &kinds.Endorse{
			Level:  level - 1,
			Cycle:   0,
			LedgerUID: status.FinalLedgerUID,
			Endorsements: []kinds.EndorseSignature{
				{
					LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
					RatifierLocation: unclearValue.PublicKey().Location(),
					Timestamp:        now,
					Autograph:        []byte("REDACTED"),
				},
			},
		}

		_, err := createLedger(status, level, corruptEndorse)
		require.Error(t, err)
		require.Contains(t, err.Error(), "REDACTED")
	})
}
