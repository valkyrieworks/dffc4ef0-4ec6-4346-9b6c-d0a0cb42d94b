package proof_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/proof"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/proof/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	machinestubs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

const (
	fallbackBallotingPotency = 10
)

func Verifyagilecustomerexploit_Insane(t *testing.T) {
	const (
		altitude       int64 = 10
		sharedAltitude int64 = 4
		sumValues          = 10
		byzantineValues            = 4
	)
	onslaughtMoment := fallbackProofMoment.Add(1 * time.Hour)
	//
	ev, reliable, shared := createInsaneProof(
		t, altitude, sharedAltitude, sumValues, byzantineValues, sumValues-byzantineValues, fallbackProofMoment, onslaughtMoment)
	require.NoError(t, ev.CertifyFundamental())

	//
	err := proof.ValidateAgileCustomerOnslaught(ev, shared.NotatedHeading, reliable.NotatedHeading, shared.AssessorAssign,
		fallbackProofMoment.Add(2*time.Hour), 3*time.Hour)
	assert.NoError(t, err)

	//
	err = proof.ValidateAgileCustomerOnslaught(ev, shared.NotatedHeading, ev.DiscordantLedger.NotatedHeading, shared.AssessorAssign,
		fallbackProofMoment.Add(2*time.Hour), 3*time.Hour)
	assert.Error(t, err)

	//
	ev.SumBallotingPotency = 1 * fallbackBallotingPotency
	err = proof.ValidateAgileCustomerOnslaught(ev, shared.NotatedHeading, reliable.NotatedHeading, shared.AssessorAssign,
		fallbackProofMoment.Add(2*time.Hour), 3*time.Hour)
	assert.Error(t, err)

	//
	ev, reliable, shared = createInsaneProof(
		t, altitude, sharedAltitude, sumValues, byzantineValues-1, sumValues-byzantineValues, fallbackProofMoment, onslaughtMoment)
	err = proof.ValidateAgileCustomerOnslaught(ev, shared.NotatedHeading, reliable.NotatedHeading, shared.AssessorAssign,
		fallbackProofMoment.Add(2*time.Hour), 3*time.Hour)
	assert.Error(t, err)
}

func Verifytest_Insaneexploitagainststate(t *testing.T) {
	const (
		altitude       int64 = 10
		sharedAltitude int64 = 4
		sumValues          = 10
		byzantineValues            = 4
	)
	onslaughtMoment := fallbackProofMoment.Add(1 * time.Hour)
	//
	ev, reliable, shared := createInsaneProof(
		t, altitude, sharedAltitude, sumValues, byzantineValues, sumValues-byzantineValues, fallbackProofMoment, onslaughtMoment)

	//
	status := sm.Status{
		FinalLedgerMoment:   fallbackProofMoment.Add(2 * time.Hour),
		FinalLedgerAltitude: altitude + 1,
		AgreementSettings: *kinds.FallbackAgreementSettings(),
	}
	statusDepot := &machinestubs.Depot{}
	statusDepot.On("REDACTED", sharedAltitude).Return(shared.AssessorAssign, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", sharedAltitude).Return(&kinds.LedgerSummary{Heading: *shared.Heading})
	ledgerDepot.On("REDACTED", altitude).Return(&kinds.LedgerSummary{Heading: *reliable.Heading})
	ledgerDepot.On("REDACTED", sharedAltitude).Return(shared.Endorse)
	ledgerDepot.On("REDACTED", altitude).Return(reliable.Endorse)
	hub, err := proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	hub.AssignTracer(log.VerifyingTracer())

	occurenceCatalog := kinds.ProofCatalog{ev}
	//
	assert.NoError(t, hub.InspectProof(occurenceCatalog))

	//
	awaitingProofs, _ := hub.AwaitingProof(status.AgreementSettings.Proof.MaximumOctets)
	assert.Equal(t, 1, len(awaitingProofs))
	assert.Equal(t, ev, awaitingProofs[0])

	//
	//
	ev.TreacherousAssessors = ev.TreacherousAssessors[:1]
	t.Log(occurenceCatalog)
	assert.Error(t, hub.InspectProof(occurenceCatalog))
	//
	ev.TreacherousAssessors = ev.ObtainTreacherousAssessors(shared.AssessorAssign, reliable.NotatedHeading)

	//
	occurenceCatalog = kinds.ProofCatalog{ev, ev}
	hub, err = proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	assert.Error(t, hub.InspectProof(occurenceCatalog))

	//
	ev.Timestamp = fallbackProofMoment.Add(1 * time.Minute)
	hub, err = proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	assert.Error(t, hub.AppendProof(ev))
	ev.Timestamp = fallbackProofMoment

	//
	ev.SumBallotingPotency = 1
	hub, err = proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	assert.Error(t, hub.AppendProof(ev))
	ev.SumBallotingPotency = shared.AssessorAssign.SumBallotingPotency()
}

func Verifytest_Insaneexploit_Byzantineassessorpubkeyswapredirectsabcimisbehavior(t *testing.T) {
	const (
		altitude       int64 = 10
		sharedAltitude int64 = 4
		sumValues          = 10
		byzantineValues            = 4
	)
	onslaughtMoment := fallbackProofMoment.Add(1 * time.Hour)

	ev, reliable, shared := createInsaneProof(
		t, altitude, sharedAltitude, sumValues, byzantineValues, sumValues-byzantineValues, fallbackProofMoment, onslaughtMoment)

	require.Len(t, ev.TreacherousAssessors, byzantineValues)
	require.GreaterOrEqual(t, shared.AssessorAssign.Extent(), byzantineValues*2)
	anotherValues := shared.AssessorAssign.Assessors[byzantineValues : byzantineValues+byzantineValues]

	initialByzantine := ev.TreacherousAssessors
	adjustedByzantine := make([]*kinds.Assessor, len(initialByzantine))
	for i := range initialByzantine {
		initial := initialByzantine[i]
		another := anotherValues[i]
		adjustedByzantine[i] = &kinds.Assessor{
			Location:          initial.Location,
			PublicToken:           another.PublicToken,
			BallotingPotency:      initial.BallotingPotency,
			NominatorUrgency: initial.NominatorUrgency,
		}
	}
	ev.TreacherousAssessors = adjustedByzantine

	status := sm.Status{
		FinalLedgerMoment:   fallbackProofMoment.Add(2 * time.Hour),
		FinalLedgerAltitude: altitude + 1,
		AgreementSettings: *kinds.FallbackAgreementSettings(),
	}
	statusDepot := &machinestubs.Depot{}
	statusDepot.On("REDACTED", sharedAltitude).Return(shared.AssessorAssign, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", sharedAltitude).Return(&kinds.LedgerSummary{Heading: *shared.Heading})
	ledgerDepot.On("REDACTED", altitude).Return(&kinds.LedgerSummary{Heading: *reliable.Heading})
	ledgerDepot.On("REDACTED", sharedAltitude).Return(shared.Endorse)
	ledgerDepot.On("REDACTED", altitude).Return(reliable.Endorse)

	hub, err := proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	hub.AssignTracer(log.VerifyingTracer())

	err = hub.AppendProof(ev)
	require.Error(t, err)
	require.Contains(t, err.Error(), "REDACTED")
}

func Verifytest_Relayinsaneexploit(t *testing.T) {
	const (
		peerAltitude   int64 = 8
		onslaughtAltitude int64 = 10
		sharedAltitude int64 = 4
		sumValues          = 10
		byzantineValues            = 5
	)
	onslaughtMoment := fallbackProofMoment.Add(1 * time.Hour)

	//
	ev, reliable, shared := createInsaneProof(
		t, onslaughtAltitude, sharedAltitude, sumValues, byzantineValues, sumValues-byzantineValues, fallbackProofMoment, onslaughtMoment)

	//
	status := sm.Status{
		FinalLedgerMoment:   fallbackProofMoment.Add(2 * time.Hour),
		FinalLedgerAltitude: peerAltitude,
		AgreementSettings: *kinds.FallbackAgreementSettings(),
	}

	//
	reliable.Altitude = status.FinalLedgerAltitude
	reliable.Moment = status.FinalLedgerMoment

	statusDepot := &machinestubs.Depot{}
	statusDepot.On("REDACTED", sharedAltitude).Return(shared.AssessorAssign, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", sharedAltitude).Return(&kinds.LedgerSummary{Heading: *shared.Heading})
	ledgerDepot.On("REDACTED", peerAltitude).Return(&kinds.LedgerSummary{Heading: *reliable.Heading})
	ledgerDepot.On("REDACTED", onslaughtAltitude).Return(nil)
	ledgerDepot.On("REDACTED", sharedAltitude).Return(shared.Endorse)
	ledgerDepot.On("REDACTED", peerAltitude).Return(reliable.Endorse)
	ledgerDepot.On("REDACTED").Return(peerAltitude)
	hub, err := proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, ledgerDepot)
	require.NoError(t, err)

	//
	assert.NoError(t, hub.InspectProof(kinds.ProofCatalog{ev}))

	//
	agedLedgerDepot := &simulations.LedgerDepot{}
	agedHeadline := reliable.Heading
	agedHeadline.Moment = fallbackProofMoment
	agedLedgerDepot.On("REDACTED", sharedAltitude).Return(&kinds.LedgerSummary{Heading: *shared.Heading})
	agedLedgerDepot.On("REDACTED", peerAltitude).Return(&kinds.LedgerSummary{Heading: *agedHeadline})
	agedLedgerDepot.On("REDACTED", onslaughtAltitude).Return(nil)
	agedLedgerDepot.On("REDACTED", sharedAltitude).Return(shared.Endorse)
	agedLedgerDepot.On("REDACTED", peerAltitude).Return(reliable.Endorse)
	agedLedgerDepot.On("REDACTED").Return(peerAltitude)
	require.Equal(t, fallbackProofMoment, agedLedgerDepot.FetchLedgerSummary(peerAltitude).Heading.Moment)

	hub, err = proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, agedLedgerDepot)
	require.NoError(t, err)
	assert.Error(t, hub.InspectProof(kinds.ProofCatalog{ev}))
}

func Verifyagilecustomerexploit_Ambiguity(t *testing.T) {
	discordantValues, discordantPrivateValues := kinds.ArbitraryAssessorAssign(5, 10)
	reliableHeading := createHeadingUnpredictable(10)

	discordantHeadline := createHeadingUnpredictable(10)
	discordantHeadline.AssessorsDigest = discordantValues.Digest()

	reliableHeading.AssessorsDigest = discordantHeadline.AssessorsDigest
	reliableHeading.FollowingAssessorsDigest = discordantHeadline.FollowingAssessorsDigest
	reliableHeading.AgreementDigest = discordantHeadline.AgreementDigest
	reliableHeading.PlatformDigest = discordantHeadline.PlatformDigest
	reliableHeading.FinalOutcomesDigest = discordantHeadline.FinalOutcomesDigest

	//
	//
	ledgerUUID := createLedgerUUID(discordantHeadline.Digest(), 1000, []byte("REDACTED"))
	ballotAssign := kinds.FreshBallotAssign(proofSuccessionUUID, 10, 1, commitchema.AttestedSignalKind(2), discordantValues)
	endorse, err := verify.CreateEndorseOriginatingBallotAssign(ledgerUUID, ballotAssign, discordantPrivateValues[:4], fallbackProofMoment)
	require.NoError(t, err)
	ev := &kinds.AgileCustomerOnslaughtProof{
		DiscordantLedger: &kinds.AgileLedger{
			NotatedHeading: &kinds.NotatedHeading{
				Heading: discordantHeadline,
				Endorse: endorse,
			},
			AssessorAssign: discordantValues,
		},
		SharedAltitude:        10,
		TreacherousAssessors: discordantValues.Assessors[:4],
		SumBallotingPotency:    50,
		Timestamp:           fallbackProofMoment,
	}

	reliableLedgerUUID := createLedgerUUID(reliableHeading.Digest(), 1000, []byte("REDACTED"))
	reliableBallotAssign := kinds.FreshBallotAssign(proofSuccessionUUID, 10, 1, commitchema.AttestedSignalKind(2), discordantValues)
	reliableEndorse, err := verify.CreateEndorseOriginatingBallotAssign(reliableLedgerUUID, reliableBallotAssign, discordantPrivateValues, fallbackProofMoment)
	require.NoError(t, err)
	reliableNotatedHeadline := &kinds.NotatedHeading{
		Heading: reliableHeading,
		Endorse: reliableEndorse,
	}

	//
	err = proof.ValidateAgileCustomerOnslaught(ev, reliableNotatedHeadline, reliableNotatedHeadline, discordantValues,
		fallbackProofMoment.Add(1*time.Minute), 2*time.Hour)
	assert.NoError(t, err)

	//
	err = proof.ValidateAgileCustomerOnslaught(ev, reliableNotatedHeadline, ev.DiscordantLedger.NotatedHeading, discordantValues,
		fallbackProofMoment.Add(1*time.Minute), 2*time.Hour)
	assert.Error(t, err)

	//
	//
	ev.DiscordantLedger.FollowingAssessorsDigest = security.CHARArbitraryOctets(tenderminthash.Extent)
	err = proof.ValidateAgileCustomerOnslaught(ev, reliableNotatedHeadline, reliableNotatedHeadline, nil,
		fallbackProofMoment.Add(1*time.Minute), 2*time.Hour)
	assert.Error(t, err)
	//
	ev.DiscordantLedger.FollowingAssessorsDigest = reliableHeading.FollowingAssessorsDigest

	status := sm.Status{
		FinalLedgerMoment:   fallbackProofMoment.Add(1 * time.Minute),
		FinalLedgerAltitude: 11,
		AgreementSettings: *kinds.FallbackAgreementSettings(),
	}
	statusDepot := &machinestubs.Depot{}
	statusDepot.On("REDACTED", int64(10)).Return(discordantValues, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", int64(10)).Return(&kinds.LedgerSummary{Heading: *reliableHeading})
	ledgerDepot.On("REDACTED", int64(10)).Return(reliableEndorse)

	hub, err := proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	hub.AssignTracer(log.VerifyingTracer())

	occurenceCatalog := kinds.ProofCatalog{ev}
	err = hub.InspectProof(occurenceCatalog)
	assert.NoError(t, err)

	awaitingProofs, _ := hub.AwaitingProof(status.AgreementSettings.Proof.MaximumOctets)
	assert.Equal(t, 1, len(awaitingProofs))
}

func Verifyagilecustomerexploit_Forgetfulness(t *testing.T) {
	discordantValues, discordantPrivateValues := kinds.ArbitraryAssessorAssign(5, 10)

	discordantHeadline := createHeadingUnpredictable(10)
	discordantHeadline.AssessorsDigest = discordantValues.Digest()
	reliableHeading := createHeadingUnpredictable(10)
	reliableHeading.AssessorsDigest = discordantHeadline.AssessorsDigest
	reliableHeading.FollowingAssessorsDigest = discordantHeadline.FollowingAssessorsDigest
	reliableHeading.PlatformDigest = discordantHeadline.PlatformDigest
	reliableHeading.AgreementDigest = discordantHeadline.AgreementDigest
	reliableHeading.FinalOutcomesDigest = discordantHeadline.FinalOutcomesDigest

	//
	//
	ledgerUUID := createLedgerUUID(discordantHeadline.Digest(), 1000, []byte("REDACTED"))
	ballotAssign := kinds.FreshBallotAssign(proofSuccessionUUID, 10, 0, commitchema.AttestedSignalKind(2), discordantValues)
	endorse, err := verify.CreateEndorseOriginatingBallotAssign(ledgerUUID, ballotAssign, discordantPrivateValues, fallbackProofMoment)
	require.NoError(t, err)
	ev := &kinds.AgileCustomerOnslaughtProof{
		DiscordantLedger: &kinds.AgileLedger{
			NotatedHeading: &kinds.NotatedHeading{
				Heading: discordantHeadline,
				Endorse: endorse,
			},
			AssessorAssign: discordantValues,
		},
		SharedAltitude:        10,
		TreacherousAssessors: nil, //
		SumBallotingPotency:    50,
		Timestamp:           fallbackProofMoment,
	}

	reliableLedgerUUID := createLedgerUUID(reliableHeading.Digest(), 1000, []byte("REDACTED"))
	reliableBallotAssign := kinds.FreshBallotAssign(proofSuccessionUUID, 10, 1, commitchema.AttestedSignalKind(2), discordantValues)
	reliableEndorse, err := verify.CreateEndorseOriginatingBallotAssign(reliableLedgerUUID, reliableBallotAssign, discordantPrivateValues, fallbackProofMoment)
	require.NoError(t, err)
	reliableNotatedHeadline := &kinds.NotatedHeading{
		Heading: reliableHeading,
		Endorse: reliableEndorse,
	}

	//
	err = proof.ValidateAgileCustomerOnslaught(ev, reliableNotatedHeadline, reliableNotatedHeadline, discordantValues,
		fallbackProofMoment.Add(1*time.Minute), 2*time.Hour)
	assert.NoError(t, err)

	//
	err = proof.ValidateAgileCustomerOnslaught(ev, reliableNotatedHeadline, ev.DiscordantLedger.NotatedHeading, discordantValues,
		fallbackProofMoment.Add(1*time.Minute), 2*time.Hour)
	assert.Error(t, err)

	status := sm.Status{
		FinalLedgerMoment:   fallbackProofMoment.Add(1 * time.Minute),
		FinalLedgerAltitude: 11,
		AgreementSettings: *kinds.FallbackAgreementSettings(),
	}
	statusDepot := &machinestubs.Depot{}
	statusDepot.On("REDACTED", int64(10)).Return(discordantValues, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", int64(10)).Return(&kinds.LedgerSummary{Heading: *reliableHeading})
	ledgerDepot.On("REDACTED", int64(10)).Return(reliableEndorse)

	hub, err := proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	hub.AssignTracer(log.VerifyingTracer())

	occurenceCatalog := kinds.ProofCatalog{ev}
	err = hub.InspectProof(occurenceCatalog)
	assert.NoError(t, err)

	awaitingProofs, _ := hub.AwaitingProof(status.AgreementSettings.Proof.MaximumOctets)
	assert.Equal(t, 1, len(awaitingProofs))
}

type ballotData struct {
	ballot1 *kinds.Ballot
	ballot2 *kinds.Ballot
	sound bool
}

func VerifyValidateReplicatedBallotProof(t *testing.T) {
	val := kinds.FreshSimulatePRV()
	valid2 := kinds.FreshSimulatePRV()
	itemAssign := kinds.FreshAssessorAssign([]*kinds.Assessor{val.DeriveWithinAssessor(1)})

	ledgerUUID := createLedgerUUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
	ledgerUuid2 := createLedgerUUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
	ledgerUuid3 := createLedgerUUID([]byte("REDACTED"), 10000, []byte("REDACTED"))
	ledgerUuid4 := createLedgerUUID([]byte("REDACTED"), 10000, []byte("REDACTED"))

	const successionUUID = "REDACTED"

	ballot1 := kinds.CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 2, 1, ledgerUUID, fallbackProofMoment)

	v1 := ballot1.TowardSchema()
	err := val.AttestBallot(successionUUID, v1)
	require.NoError(t, err)
	flawedBallot := kinds.CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 2, 1, ledgerUUID, fallbackProofMoment)
	bv := flawedBallot.TowardSchema()
	err = valid2.AttestBallot(successionUUID, bv)
	require.NoError(t, err)

	ballot1.Notation = v1.Notation
	flawedBallot.Notation = bv.Notation

	scenarios := []ballotData{
		{ballot1, kinds.CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 2, 1, ledgerUuid2, fallbackProofMoment), true}, //
		{ballot1, kinds.CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 2, 1, ledgerUuid3, fallbackProofMoment), true},
		{ballot1, kinds.CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 2, 1, ledgerUuid4, fallbackProofMoment), true},
		{ballot1, kinds.CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 2, 1, ledgerUUID, fallbackProofMoment), false},     //
		{ballot1, kinds.CreateBallotNegativeFailure(t, val, "REDACTED", 0, 10, 2, 1, ledgerUuid2, fallbackProofMoment), false}, //
		{ballot1, kinds.CreateBallotNegativeFailure(t, val, successionUUID, 0, 11, 2, 1, ledgerUuid2, fallbackProofMoment), false},    //
		{ballot1, kinds.CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 3, 1, ledgerUuid2, fallbackProofMoment), false},    //
		{ballot1, kinds.CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 2, 2, ledgerUuid2, fallbackProofMoment), false},    //
		{ballot1, kinds.CreateBallotNegativeFailure(t, valid2, successionUUID, 0, 10, 2, 1, ledgerUuid2, fallbackProofMoment), false},   //
		//
		{ballot1, kinds.CreateBallotNegativeFailure(t, val, successionUUID, 0, 10, 2, 1, ledgerUuid2, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)), true},
		{ballot1, flawedBallot, false}, //
	}

	require.NoError(t, err)
	for _, c := range scenarios {
		ev := &kinds.ReplicatedBallotProof{
			BallotAN:            c.ballot1,
			BallotBYTE:            c.ballot2,
			AssessorPotency:   1,
			SumBallotingPotency: 1,
			Timestamp:        fallbackProofMoment,
		}
		if c.sound {
			assert.Nil(t, proof.ValidateReplicatedBallot(ev, successionUUID, itemAssign), "REDACTED")
		} else {
			assert.NotNil(t, proof.ValidateReplicatedBallot(ev, successionUUID, itemAssign), "REDACTED")
		}
	}

	//
	validOccurence, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(10, fallbackProofMoment, val, successionUUID)
	require.NoError(t, err)
	validOccurence.AssessorPotency = 1
	validOccurence.SumBallotingPotency = 1
	flawedOccurence, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(10, fallbackProofMoment, val, successionUUID)
	require.NoError(t, err)
	flawedMomentOccurence, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(10, fallbackProofMoment.Add(1*time.Minute), val, successionUUID)
	require.NoError(t, err)
	flawedMomentOccurence.AssessorPotency = 1
	flawedMomentOccurence.SumBallotingPotency = 1
	status := sm.Status{
		SuccessionUUID:         successionUUID,
		FinalLedgerMoment:   fallbackProofMoment.Add(1 * time.Minute),
		FinalLedgerAltitude: 11,
		AgreementSettings: *kinds.FallbackAgreementSettings(),
	}
	statusDepot := &machinestubs.Depot{}
	statusDepot.On("REDACTED", int64(10)).Return(itemAssign, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", int64(10)).Return(&kinds.LedgerSummary{Heading: kinds.Heading{Moment: fallbackProofMoment}})

	hub, err := proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, ledgerDepot)
	require.NoError(t, err)

	occurenceCatalog := kinds.ProofCatalog{validOccurence}
	err = hub.InspectProof(occurenceCatalog)
	assert.NoError(t, err)

	//
	occurenceCatalog = kinds.ProofCatalog{flawedOccurence}
	err = hub.InspectProof(occurenceCatalog)
	assert.Error(t, err)

	//
	occurenceCatalog = kinds.ProofCatalog{flawedMomentOccurence}
	err = hub.InspectProof(occurenceCatalog)
	assert.Error(t, err)
}

func createInsaneProof(
	t *testing.T,
	altitude, sharedAltitude int64,
	sumValues, byzantineValues, apparitionValues int,
	sharedMoment, onslaughtMoment time.Time,
) (ev *kinds.AgileCustomerOnslaughtProof, reliable *kinds.AgileLedger, shared *kinds.AgileLedger) {
	sharedItemAssign, sharedPrivateValues := kinds.ArbitraryAssessorAssign(sumValues, fallbackBallotingPotency)

	require.Greater(t, sumValues, byzantineValues)

	//
	byzantineItemAssign, byzantinePrivateValues := sharedItemAssign.Assessors[:byzantineValues], sharedPrivateValues[:byzantineValues]

	apparitionItemAssign, apparitionPrivateValues := kinds.ArbitraryAssessorAssign(apparitionValues, fallbackBallotingPotency)

	discordantValues := apparitionItemAssign.Duplicate()
	require.NoError(t, discordantValues.ReviseUsingModifyAssign(byzantineItemAssign))
	discordantPrivateValues := append(apparitionPrivateValues, byzantinePrivateValues...) //

	discordantPrivateValues = sequencePrivateValuesViaItemAssign(t, discordantValues, discordantPrivateValues)

	sharedHeadline := createHeadingUnpredictable(sharedAltitude)
	sharedHeadline.Moment = sharedMoment
	reliableHeading := createHeadingUnpredictable(altitude)

	discordantHeadline := createHeadingUnpredictable(altitude)
	discordantHeadline.Moment = onslaughtMoment
	discordantHeadline.AssessorsDigest = discordantValues.Digest()

	ledgerUUID := createLedgerUUID(discordantHeadline.Digest(), 1000, []byte("REDACTED"))
	ballotAssign := kinds.FreshBallotAssign(proofSuccessionUUID, altitude, 1, commitchema.AttestedSignalKind(2), discordantValues)
	endorse, err := verify.CreateEndorseOriginatingBallotAssign(ledgerUUID, ballotAssign, discordantPrivateValues, fallbackProofMoment)
	require.NoError(t, err)
	ev = &kinds.AgileCustomerOnslaughtProof{
		DiscordantLedger: &kinds.AgileLedger{
			NotatedHeading: &kinds.NotatedHeading{
				Heading: discordantHeadline,
				Endorse: endorse,
			},
			AssessorAssign: discordantValues,
		},
		SharedAltitude:        sharedAltitude,
		SumBallotingPotency:    sharedItemAssign.SumBallotingPotency(),
		TreacherousAssessors: byzantineItemAssign,
		Timestamp:           sharedMoment,
	}

	shared = &kinds.AgileLedger{
		NotatedHeading: &kinds.NotatedHeading{
			Heading: sharedHeadline,
			//
			Endorse: &kinds.Endorse{},
		},
		AssessorAssign: sharedItemAssign,
	}
	reliableLedgerUUID := createLedgerUUID(reliableHeading.Digest(), 1000, []byte("REDACTED"))
	reliableValues, privateItems := kinds.ArbitraryAssessorAssign(sumValues, fallbackBallotingPotency)
	reliableBallotAssign := kinds.FreshBallotAssign(proofSuccessionUUID, altitude, 1, commitchema.AttestedSignalKind(2), reliableValues)
	reliableEndorse, err := verify.CreateEndorseOriginatingBallotAssign(reliableLedgerUUID, reliableBallotAssign, privateItems, fallbackProofMoment)
	require.NoError(t, err)
	reliable = &kinds.AgileLedger{
		NotatedHeading: &kinds.NotatedHeading{
			Heading: reliableHeading,
			Endorse: reliableEndorse,
		},
		AssessorAssign: reliableValues,
	}
	return ev, reliable, shared
}

//

//

//

//

func createHeadingUnpredictable(altitude int64) *kinds.Heading {
	return &kinds.Heading{
		Edition:            strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 1},
		SuccessionUUID:            proofSuccessionUUID,
		Altitude:             altitude,
		Moment:               fallbackProofMoment,
		FinalLedgerUUID:        createLedgerUUID([]byte("REDACTED"), 1000, []byte("REDACTED")),
		FinalEndorseDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
		DataDigest:           security.CHARArbitraryOctets(tenderminthash.Extent),
		AssessorsDigest:     security.CHARArbitraryOctets(tenderminthash.Extent),
		FollowingAssessorsDigest: security.CHARArbitraryOctets(tenderminthash.Extent),
		AgreementDigest:      security.CHARArbitraryOctets(tenderminthash.Extent),
		PlatformDigest:            security.CHARArbitraryOctets(tenderminthash.Extent),
		FinalOutcomesDigest:    security.CHARArbitraryOctets(tenderminthash.Extent),
		ProofDigest:       security.CHARArbitraryOctets(tenderminthash.Extent),
		NominatorLocation:    security.CHARArbitraryOctets(security.LocatorExtent),
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

func sequencePrivateValuesViaItemAssign(
	t *testing.T, values *kinds.AssessorAssign, privateItems []kinds.PrivateAssessor,
) []kinds.PrivateAssessor {
	emission := make([]kinds.PrivateAssessor, len(privateItems))
	for idx, v := range values.Assessors {
		for _, p := range privateItems {
			publicToken, err := p.ObtainPublicToken()
			require.NoError(t, err)
			if bytes.Equal(v.Location, publicToken.Location()) {
				emission[idx] = p
				break
			}
		}
		require.NotEmpty(t, emission[idx])
	}
	return emission
}
