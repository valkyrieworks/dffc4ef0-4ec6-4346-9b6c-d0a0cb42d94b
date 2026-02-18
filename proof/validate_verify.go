package proofs_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/proof"
	"github.com/valkyrieworks/proof/simulations"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	sm "github.com/valkyrieworks/status"
	smemulators "github.com/valkyrieworks/status/simulations"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

const (
	standardPollingEnergy = 10
)

func Testvalidaterapidcustomerassault_Erratic(t *testing.T) {
	const (
		level       int64 = 10
		sharedLevel int64 = 4
		sumValues          = 10
		byzValues            = 4
	)
	assaultTime := standardProofTime.Add(1 * time.Hour)
	//
	ev, validated, shared := createErraticProof(
		t, level, sharedLevel, sumValues, byzValues, sumValues-byzValues, standardProofTime, assaultTime)
	require.NoError(t, ev.CertifySimple())

	//
	err := proof.ValidateRapidCustomerAssault(ev, shared.AttestedHeading, validated.AttestedHeading, shared.RatifierAssign,
		standardProofTime.Add(2*time.Hour), 3*time.Hour)
	assert.NoError(t, err)

	//
	err = proof.ValidateRapidCustomerAssault(ev, shared.AttestedHeading, ev.ClashingLedger.AttestedHeading, shared.RatifierAssign,
		standardProofTime.Add(2*time.Hour), 3*time.Hour)
	assert.Error(t, err)

	//
	ev.SumPollingEnergy = 1 * standardPollingEnergy
	err = proof.ValidateRapidCustomerAssault(ev, shared.AttestedHeading, validated.AttestedHeading, shared.RatifierAssign,
		standardProofTime.Add(2*time.Hour), 3*time.Hour)
	assert.Error(t, err)

	//
	ev, validated, shared = createErraticProof(
		t, level, sharedLevel, sumValues, byzValues-1, sumValues-byzValues, standardProofTime, assaultTime)
	err = proof.ValidateRapidCustomerAssault(ev, shared.AttestedHeading, validated.AttestedHeading, shared.RatifierAssign,
		standardProofTime.Add(2*time.Hour), 3*time.Hour)
	assert.Error(t, err)
}

func Testvalidate_Erraticassaultonstate(t *testing.T) {
	const (
		level       int64 = 10
		sharedLevel int64 = 4
		sumValues          = 10
		byzValues            = 4
	)
	assaultTime := standardProofTime.Add(1 * time.Hour)
	//
	ev, validated, shared := createErraticProof(
		t, level, sharedLevel, sumValues, byzValues, sumValues-byzValues, standardProofTime, assaultTime)

	//
	status := sm.Status{
		FinalLedgerTime:   standardProofTime.Add(2 * time.Hour),
		FinalLedgerLevel: level + 1,
		AgreementOptions: *kinds.StandardAgreementOptions(),
	}
	statusDepot := &smemulators.Depot{}
	statusDepot.On("REDACTED", sharedLevel).Return(shared.RatifierAssign, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", sharedLevel).Return(&kinds.LedgerMeta{Heading: *shared.Heading})
	ledgerDepot.On("REDACTED", level).Return(&kinds.LedgerMeta{Heading: *validated.Heading})
	ledgerDepot.On("REDACTED", sharedLevel).Return(shared.Endorse)
	ledgerDepot.On("REDACTED", level).Return(validated.Endorse)
	depository, err := proof.NewDepository(dbm.NewMemoryStore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	depository.AssignTracer(log.VerifyingTracer())

	evtCatalog := kinds.ProofCatalog{ev}
	//
	assert.NoError(t, depository.InspectProof(evtCatalog))

	//
	awaitingEvidences, _ := depository.AwaitingProof(status.AgreementOptions.Proof.MaximumOctets)
	assert.Equal(t, 1, len(awaitingEvidences))
	assert.Equal(t, ev, awaitingEvidences[0])

	//
	//
	ev.FaultyRatifiers = ev.FaultyRatifiers[:1]
	t.Log(evtCatalog)
	assert.Error(t, depository.InspectProof(evtCatalog))
	//
	ev.FaultyRatifiers = ev.FetchFaultyRatifiers(shared.RatifierAssign, validated.AttestedHeading)

	//
	evtCatalog = kinds.ProofCatalog{ev, ev}
	depository, err = proof.NewDepository(dbm.NewMemoryStore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	assert.Error(t, depository.InspectProof(evtCatalog))

	//
	ev.Timestamp = standardProofTime.Add(1 * time.Minute)
	depository, err = proof.NewDepository(dbm.NewMemoryStore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	assert.Error(t, depository.AppendProof(ev))
	ev.Timestamp = standardProofTime

	//
	ev.SumPollingEnergy = 1
	depository, err = proof.NewDepository(dbm.NewMemoryStore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	assert.Error(t, depository.AppendProof(ev))
	ev.SumPollingEnergy = shared.RatifierAssign.SumPollingEnergy()
}

func Testvalidate_Transmiterraticassault(t *testing.T) {
	const (
		memberLevel   int64 = 8
		assaultLevel int64 = 10
		sharedLevel int64 = 4
		sumValues          = 10
		byzValues            = 5
	)
	assaultTime := standardProofTime.Add(1 * time.Hour)

	//
	ev, validated, shared := createErraticProof(
		t, assaultLevel, sharedLevel, sumValues, byzValues, sumValues-byzValues, standardProofTime, assaultTime)

	//
	status := sm.Status{
		FinalLedgerTime:   standardProofTime.Add(2 * time.Hour),
		FinalLedgerLevel: memberLevel,
		AgreementOptions: *kinds.StandardAgreementOptions(),
	}

	//
	validated.Level = status.FinalLedgerLevel
	validated.Time = status.FinalLedgerTime

	statusDepot := &smemulators.Depot{}
	statusDepot.On("REDACTED", sharedLevel).Return(shared.RatifierAssign, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", sharedLevel).Return(&kinds.LedgerMeta{Heading: *shared.Heading})
	ledgerDepot.On("REDACTED", memberLevel).Return(&kinds.LedgerMeta{Heading: *validated.Heading})
	ledgerDepot.On("REDACTED", assaultLevel).Return(nil)
	ledgerDepot.On("REDACTED", sharedLevel).Return(shared.Endorse)
	ledgerDepot.On("REDACTED", memberLevel).Return(validated.Endorse)
	ledgerDepot.On("REDACTED").Return(memberLevel)
	depository, err := proof.NewDepository(dbm.NewMemoryStore(), statusDepot, ledgerDepot)
	require.NoError(t, err)

	//
	assert.NoError(t, depository.InspectProof(kinds.ProofCatalog{ev}))

	//
	agedLedgerDepot := &simulations.LedgerDepot{}
	agedHeading := validated.Heading
	agedHeading.Time = standardProofTime
	agedLedgerDepot.On("REDACTED", sharedLevel).Return(&kinds.LedgerMeta{Heading: *shared.Heading})
	agedLedgerDepot.On("REDACTED", memberLevel).Return(&kinds.LedgerMeta{Heading: *agedHeading})
	agedLedgerDepot.On("REDACTED", assaultLevel).Return(nil)
	agedLedgerDepot.On("REDACTED", sharedLevel).Return(shared.Endorse)
	agedLedgerDepot.On("REDACTED", memberLevel).Return(validated.Endorse)
	agedLedgerDepot.On("REDACTED").Return(memberLevel)
	require.Equal(t, standardProofTime, agedLedgerDepot.ImportLedgerMeta(memberLevel).Heading.Time)

	depository, err = proof.NewDepository(dbm.NewMemoryStore(), statusDepot, agedLedgerDepot)
	require.NoError(t, err)
	assert.Error(t, depository.InspectProof(kinds.ProofCatalog{ev}))
}

func Testvalidaterapidcustomerassault_Ambiguity(t *testing.T) {
	clashingValues, clashingPrivateValues := kinds.RandomRatifierCollection(5, 10)
	validatedHeading := createHeadingArbitrary(10)

	clashingHeading := createHeadingArbitrary(10)
	clashingHeading.RatifiersDigest = clashingValues.Digest()

	validatedHeading.RatifiersDigest = clashingHeading.RatifiersDigest
	validatedHeading.FollowingRatifiersDigest = clashingHeading.FollowingRatifiersDigest
	validatedHeading.AgreementDigest = clashingHeading.AgreementDigest
	validatedHeading.ApplicationDigest = clashingHeading.ApplicationDigest
	validatedHeading.FinalOutcomesDigest = clashingHeading.FinalOutcomesDigest

	//
	//
	ledgerUID := createLedgerUID(clashingHeading.Digest(), 1000, []byte("REDACTED"))
	ballotCollection := kinds.NewBallotCollection(proofSeriesUID, 10, 1, engineproto.AttestedMessageKind(2), clashingValues)
	endorse, err := verify.CreateEndorseFromBallotCollection(ledgerUID, ballotCollection, clashingPrivateValues[:4], standardProofTime)
	require.NoError(t, err)
	ev := &kinds.RapidCustomerAssaultProof{
		ClashingLedger: &kinds.RapidLedger{
			AttestedHeading: &kinds.AttestedHeading{
				Heading: clashingHeading,
				Endorse: endorse,
			},
			RatifierAssign: clashingValues,
		},
		SharedLevel:        10,
		FaultyRatifiers: clashingValues.Ratifiers[:4],
		SumPollingEnergy:    50,
		Timestamp:           standardProofTime,
	}

	validatedLedgerUID := createLedgerUID(validatedHeading.Digest(), 1000, []byte("REDACTED"))
	validatedBallotCollection := kinds.NewBallotCollection(proofSeriesUID, 10, 1, engineproto.AttestedMessageKind(2), clashingValues)
	validatedEndorse, err := verify.CreateEndorseFromBallotCollection(validatedLedgerUID, validatedBallotCollection, clashingPrivateValues, standardProofTime)
	require.NoError(t, err)
	validatedAttestedHeading := &kinds.AttestedHeading{
		Heading: validatedHeading,
		Endorse: validatedEndorse,
	}

	//
	err = proof.ValidateRapidCustomerAssault(ev, validatedAttestedHeading, validatedAttestedHeading, clashingValues,
		standardProofTime.Add(1*time.Minute), 2*time.Hour)
	assert.NoError(t, err)

	//
	err = proof.ValidateRapidCustomerAssault(ev, validatedAttestedHeading, ev.ClashingLedger.AttestedHeading, clashingValues,
		standardProofTime.Add(1*time.Minute), 2*time.Hour)
	assert.Error(t, err)

	//
	//
	ev.ClashingLedger.FollowingRatifiersDigest = vault.CRandomOctets(comethash.Volume)
	err = proof.ValidateRapidCustomerAssault(ev, validatedAttestedHeading, validatedAttestedHeading, nil,
		standardProofTime.Add(1*time.Minute), 2*time.Hour)
	assert.Error(t, err)
	//
	ev.ClashingLedger.FollowingRatifiersDigest = validatedHeading.FollowingRatifiersDigest

	status := sm.Status{
		FinalLedgerTime:   standardProofTime.Add(1 * time.Minute),
		FinalLedgerLevel: 11,
		AgreementOptions: *kinds.StandardAgreementOptions(),
	}
	statusDepot := &smemulators.Depot{}
	statusDepot.On("REDACTED", int64(10)).Return(clashingValues, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", int64(10)).Return(&kinds.LedgerMeta{Heading: *validatedHeading})
	ledgerDepot.On("REDACTED", int64(10)).Return(validatedEndorse)

	depository, err := proof.NewDepository(dbm.NewMemoryStore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	depository.AssignTracer(log.VerifyingTracer())

	evtCatalog := kinds.ProofCatalog{ev}
	err = depository.InspectProof(evtCatalog)
	assert.NoError(t, err)

	awaitingEvidences, _ := depository.AwaitingProof(status.AgreementOptions.Proof.MaximumOctets)
	assert.Equal(t, 1, len(awaitingEvidences))
}

func Testvalidaterapidcustomerassault_Forgetfulness(t *testing.T) {
	clashingValues, clashingPrivateValues := kinds.RandomRatifierCollection(5, 10)

	clashingHeading := createHeadingArbitrary(10)
	clashingHeading.RatifiersDigest = clashingValues.Digest()
	validatedHeading := createHeadingArbitrary(10)
	validatedHeading.RatifiersDigest = clashingHeading.RatifiersDigest
	validatedHeading.FollowingRatifiersDigest = clashingHeading.FollowingRatifiersDigest
	validatedHeading.ApplicationDigest = clashingHeading.ApplicationDigest
	validatedHeading.AgreementDigest = clashingHeading.AgreementDigest
	validatedHeading.FinalOutcomesDigest = clashingHeading.FinalOutcomesDigest

	//
	//
	ledgerUID := createLedgerUID(clashingHeading.Digest(), 1000, []byte("REDACTED"))
	ballotCollection := kinds.NewBallotCollection(proofSeriesUID, 10, 0, engineproto.AttestedMessageKind(2), clashingValues)
	endorse, err := verify.CreateEndorseFromBallotCollection(ledgerUID, ballotCollection, clashingPrivateValues, standardProofTime)
	require.NoError(t, err)
	ev := &kinds.RapidCustomerAssaultProof{
		ClashingLedger: &kinds.RapidLedger{
			AttestedHeading: &kinds.AttestedHeading{
				Heading: clashingHeading,
				Endorse: endorse,
			},
			RatifierAssign: clashingValues,
		},
		SharedLevel:        10,
		FaultyRatifiers: nil, //
		SumPollingEnergy:    50,
		Timestamp:           standardProofTime,
	}

	validatedLedgerUID := createLedgerUID(validatedHeading.Digest(), 1000, []byte("REDACTED"))
	validatedBallotCollection := kinds.NewBallotCollection(proofSeriesUID, 10, 1, engineproto.AttestedMessageKind(2), clashingValues)
	validatedEndorse, err := verify.CreateEndorseFromBallotCollection(validatedLedgerUID, validatedBallotCollection, clashingPrivateValues, standardProofTime)
	require.NoError(t, err)
	validatedAttestedHeading := &kinds.AttestedHeading{
		Heading: validatedHeading,
		Endorse: validatedEndorse,
	}

	//
	err = proof.ValidateRapidCustomerAssault(ev, validatedAttestedHeading, validatedAttestedHeading, clashingValues,
		standardProofTime.Add(1*time.Minute), 2*time.Hour)
	assert.NoError(t, err)

	//
	err = proof.ValidateRapidCustomerAssault(ev, validatedAttestedHeading, ev.ClashingLedger.AttestedHeading, clashingValues,
		standardProofTime.Add(1*time.Minute), 2*time.Hour)
	assert.Error(t, err)

	status := sm.Status{
		FinalLedgerTime:   standardProofTime.Add(1 * time.Minute),
		FinalLedgerLevel: 11,
		AgreementOptions: *kinds.StandardAgreementOptions(),
	}
	statusDepot := &smemulators.Depot{}
	statusDepot.On("REDACTED", int64(10)).Return(clashingValues, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", int64(10)).Return(&kinds.LedgerMeta{Heading: *validatedHeading})
	ledgerDepot.On("REDACTED", int64(10)).Return(validatedEndorse)

	depository, err := proof.NewDepository(dbm.NewMemoryStore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	depository.AssignTracer(log.VerifyingTracer())

	evtCatalog := kinds.ProofCatalog{ev}
	err = depository.InspectProof(evtCatalog)
	assert.NoError(t, err)

	awaitingEvidences, _ := depository.AwaitingProof(status.AgreementOptions.Proof.MaximumOctets)
	assert.Equal(t, 1, len(awaitingEvidences))
}

type ballotData struct {
	vote1 *kinds.Ballot
	ballot2 *kinds.Ballot
	sound bool
}

func VerifyValidateReplicatedBallotProof(t *testing.T) {
	val := kinds.NewEmulatePV()
	value2 := kinds.NewEmulatePV()
	valueCollection := kinds.NewRatifierCollection([]*kinds.Ratifier{val.RetrieveTowardRatifier(1)})

	ledgerUID := createLedgerUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
	ledgerUidtwo := createLedgerUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
	ledgerID3 := createLedgerUID([]byte("REDACTED"), 10000, []byte("REDACTED"))
	ledgerID4 := createLedgerUID([]byte("REDACTED"), 10000, []byte("REDACTED"))

	const ledgerUID = "REDACTED"

	vote1 := kinds.CreateBallotNoFault(t, val, ledgerUID, 0, 10, 2, 1, ledgerUID, standardProofTime)

	v1 := vote1.ToSchema()
	err := val.AttestBallot(ledgerUID, v1)
	require.NoError(t, err)
	flawedBallot := kinds.CreateBallotNoFault(t, val, ledgerUID, 0, 10, 2, 1, ledgerUID, standardProofTime)
	bv := flawedBallot.ToSchema()
	err = value2.AttestBallot(ledgerUID, bv)
	require.NoError(t, err)

	vote1.Autograph = v1.Autograph
	flawedBallot.Autograph = bv.Autograph

	scenarios := []ballotData{
		{vote1, kinds.CreateBallotNoFault(t, val, ledgerUID, 0, 10, 2, 1, ledgerUidtwo, standardProofTime), true}, //
		{vote1, kinds.CreateBallotNoFault(t, val, ledgerUID, 0, 10, 2, 1, ledgerID3, standardProofTime), true},
		{vote1, kinds.CreateBallotNoFault(t, val, ledgerUID, 0, 10, 2, 1, ledgerID4, standardProofTime), true},
		{vote1, kinds.CreateBallotNoFault(t, val, ledgerUID, 0, 10, 2, 1, ledgerUID, standardProofTime), false},     //
		{vote1, kinds.CreateBallotNoFault(t, val, "REDACTED", 0, 10, 2, 1, ledgerUidtwo, standardProofTime), false}, //
		{vote1, kinds.CreateBallotNoFault(t, val, ledgerUID, 0, 11, 2, 1, ledgerUidtwo, standardProofTime), false},    //
		{vote1, kinds.CreateBallotNoFault(t, val, ledgerUID, 0, 10, 3, 1, ledgerUidtwo, standardProofTime), false},    //
		{vote1, kinds.CreateBallotNoFault(t, val, ledgerUID, 0, 10, 2, 2, ledgerUidtwo, standardProofTime), false},    //
		{vote1, kinds.CreateBallotNoFault(t, value2, ledgerUID, 0, 10, 2, 1, ledgerUidtwo, standardProofTime), false},   //
		//
		{vote1, kinds.CreateBallotNoFault(t, val, ledgerUID, 0, 10, 2, 1, ledgerUidtwo, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)), true},
		{vote1, flawedBallot, false}, //
	}

	require.NoError(t, err)
	for _, c := range scenarios {
		ev := &kinds.ReplicatedBallotProof{
			BallotA:            c.vote1,
			BallotBYTE:            c.ballot2,
			RatifierEnergy:   1,
			SumPollingEnergy: 1,
			Timestamp:        standardProofTime,
		}
		if c.sound {
			assert.Nil(t, proof.ValidateReplicatedBallot(ev, ledgerUID, valueCollection), "REDACTED")
		} else {
			assert.NotNil(t, proof.ValidateReplicatedBallot(ev, ledgerUID, valueCollection), "REDACTED")
		}
	}

	//
	validEvt, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(10, standardProofTime, val, ledgerUID)
	require.NoError(t, err)
	validEvt.RatifierEnergy = 1
	validEvt.SumPollingEnergy = 1
	flawedEvt, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(10, standardProofTime, val, ledgerUID)
	require.NoError(t, err)
	flawedTimeEvt, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(10, standardProofTime.Add(1*time.Minute), val, ledgerUID)
	require.NoError(t, err)
	flawedTimeEvt.RatifierEnergy = 1
	flawedTimeEvt.SumPollingEnergy = 1
	status := sm.Status{
		LedgerUID:         ledgerUID,
		FinalLedgerTime:   standardProofTime.Add(1 * time.Minute),
		FinalLedgerLevel: 11,
		AgreementOptions: *kinds.StandardAgreementOptions(),
	}
	statusDepot := &smemulators.Depot{}
	statusDepot.On("REDACTED", int64(10)).Return(valueCollection, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", int64(10)).Return(&kinds.LedgerMeta{Heading: kinds.Heading{Time: standardProofTime}})

	depository, err := proof.NewDepository(dbm.NewMemoryStore(), statusDepot, ledgerDepot)
	require.NoError(t, err)

	evtCatalog := kinds.ProofCatalog{validEvt}
	err = depository.InspectProof(evtCatalog)
	assert.NoError(t, err)

	//
	evtCatalog = kinds.ProofCatalog{flawedEvt}
	err = depository.InspectProof(evtCatalog)
	assert.Error(t, err)

	//
	evtCatalog = kinds.ProofCatalog{flawedTimeEvt}
	err = depository.InspectProof(evtCatalog)
	assert.Error(t, err)
}

func createErraticProof(
	t *testing.T,
	level, sharedLevel int64,
	sumValues, byzValues, specterValues int,
	sharedTime, assaultTime time.Time,
) (ev *kinds.RapidCustomerAssaultProof, validated *kinds.RapidLedger, shared *kinds.RapidLedger) {
	sharedValueCollection, sharedPrivateValues := kinds.RandomRatifierCollection(sumValues, standardPollingEnergy)

	require.Greater(t, sumValues, byzValues)

	//
	byzValueCollection, byzPrivateValues := sharedValueCollection.Ratifiers[:byzValues], sharedPrivateValues[:byzValues]

	specterValueCollection, specterPrivateValues := kinds.RandomRatifierCollection(specterValues, standardPollingEnergy)

	clashingValues := specterValueCollection.Clone()
	require.NoError(t, clashingValues.ModifyWithAlterCollection(byzValueCollection))
	clashingPrivateValues := append(specterPrivateValues, byzPrivateValues...) //

	clashingPrivateValues = sequencePrivateValuesByValueCollection(t, clashingValues, clashingPrivateValues)

	sharedHeading := createHeadingArbitrary(sharedLevel)
	sharedHeading.Time = sharedTime
	validatedHeading := createHeadingArbitrary(level)

	clashingHeading := createHeadingArbitrary(level)
	clashingHeading.Time = assaultTime
	clashingHeading.RatifiersDigest = clashingValues.Digest()

	ledgerUID := createLedgerUID(clashingHeading.Digest(), 1000, []byte("REDACTED"))
	ballotCollection := kinds.NewBallotCollection(proofSeriesUID, level, 1, engineproto.AttestedMessageKind(2), clashingValues)
	endorse, err := verify.CreateEndorseFromBallotCollection(ledgerUID, ballotCollection, clashingPrivateValues, standardProofTime)
	require.NoError(t, err)
	ev = &kinds.RapidCustomerAssaultProof{
		ClashingLedger: &kinds.RapidLedger{
			AttestedHeading: &kinds.AttestedHeading{
				Heading: clashingHeading,
				Endorse: endorse,
			},
			RatifierAssign: clashingValues,
		},
		SharedLevel:        sharedLevel,
		SumPollingEnergy:    sharedValueCollection.SumPollingEnergy(),
		FaultyRatifiers: byzValueCollection,
		Timestamp:           sharedTime,
	}

	shared = &kinds.RapidLedger{
		AttestedHeading: &kinds.AttestedHeading{
			Heading: sharedHeading,
			//
			Endorse: &kinds.Endorse{},
		},
		RatifierAssign: sharedValueCollection,
	}
	validatedLedgerUID := createLedgerUID(validatedHeading.Digest(), 1000, []byte("REDACTED"))
	validatedValues, privateValues := kinds.RandomRatifierCollection(sumValues, standardPollingEnergy)
	validatedBallotCollection := kinds.NewBallotCollection(proofSeriesUID, level, 1, engineproto.AttestedMessageKind(2), validatedValues)
	validatedEndorse, err := verify.CreateEndorseFromBallotCollection(validatedLedgerUID, validatedBallotCollection, privateValues, standardProofTime)
	require.NoError(t, err)
	validated = &kinds.RapidLedger{
		AttestedHeading: &kinds.AttestedHeading{
			Heading: validatedHeading,
			Endorse: validatedEndorse,
		},
		RatifierAssign: validatedValues,
	}
	return ev, validated, shared
}

//

//

//

//

func createHeadingArbitrary(level int64) *kinds.Heading {
	return &kinds.Heading{
		Release:            cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 1},
		LedgerUID:            proofSeriesUID,
		Level:             level,
		Time:               standardProofTime,
		FinalLedgerUID:        createLedgerUID([]byte("REDACTED"), 1000, []byte("REDACTED")),
		FinalEndorseDigest:     vault.CRandomOctets(comethash.Volume),
		DataDigest:           vault.CRandomOctets(comethash.Volume),
		RatifiersDigest:     vault.CRandomOctets(comethash.Volume),
		FollowingRatifiersDigest: vault.CRandomOctets(comethash.Volume),
		AgreementDigest:      vault.CRandomOctets(comethash.Volume),
		ApplicationDigest:            vault.CRandomOctets(comethash.Volume),
		FinalOutcomesDigest:    vault.CRandomOctets(comethash.Volume),
		ProofDigest:       vault.CRandomOctets(comethash.Volume),
		RecommenderLocation:    vault.CRandomOctets(vault.LocationVolume),
	}
}

func createLedgerUID(digest []byte, sectionCollectionVolume uint32, sectionCollectionDigest []byte) kinds.LedgerUID {
	var (
		h   = make([]byte, comethash.Volume)
		psH = make([]byte, comethash.Volume)
	)
	copy(h, digest)
	copy(psH, sectionCollectionDigest)
	return kinds.LedgerUID{
		Digest: h,
		SegmentAssignHeading: kinds.SegmentAssignHeading{
			Sum: sectionCollectionVolume,
			Digest:  psH,
		},
	}
}

func sequencePrivateValuesByValueCollection(
	t *testing.T, values *kinds.RatifierAssign, privateValues []kinds.PrivateRatifier,
) []kinds.PrivateRatifier {
	result := make([]kinds.PrivateRatifier, len(privateValues))
	for idx, v := range values.Ratifiers {
		for _, p := range privateValues {
			publicKey, err := p.FetchPublicKey()
			require.NoError(t, err)
			if bytes.Equal(v.Location, publicKey.Location()) {
				result[idx] = p
				break
			}
		}
		require.NotEmpty(t, result[idx])
	}
	return result
}
