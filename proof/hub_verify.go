package proof_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/proof"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/proof/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongmindedition "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/edition"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	machinestubs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

func VerifyPrimary(m *testing.M) {
	cipher := m.Run()
	os.Exit(cipher)
}

const proofSuccessionUUID = "REDACTED"

var (
	fallbackProofMoment           = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	fallbackProofMaximumOctets int64 = 1000
)

func VerifyProofHubFundamental(t *testing.T) {
	var (
		altitude     = int64(1)
		statusDepot = &machinestubs.Depot{}
		proofDatastore = dbm.FreshMemoryDatastore()
		ledgerDepot = &simulations.LedgerDepot{}
	)

	itemAssign, privateItems := kinds.ArbitraryAssessorAssign(1, 10)

	ledgerDepot.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(
		&kinds.LedgerSummary{Heading: kinds.Heading{Moment: fallbackProofMoment}},
	)
	statusDepot.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(itemAssign, nil)
	statusDepot.On("REDACTED").Return(generateStatus(altitude+1, itemAssign), nil)

	hub, err := proof.FreshHub(proofDatastore, statusDepot, ledgerDepot)
	require.NoError(t, err)
	hub.AssignTracer(log.VerifyingTracer())

	//
	evs, extent := hub.AwaitingProof(fallbackProofMaximumOctets)
	assert.Equal(t, 0, len(evs))
	assert.Zero(t, extent)

	ev, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(altitude, fallbackProofMoment, privateItems[0], proofSuccessionUUID)
	require.NoError(t, err)

	//
	occurenceAppended := make(chan struct{})
	go func() {
		<-hub.ProofPauseChnl()
		close(occurenceAppended)
	}()

	//
	assert.NoError(t, hub.AppendProof(ev))

	select {
	case <-occurenceAppended:
	case <-time.After(5 * time.Second):
		t.Fatal("REDACTED")
	}

	following := hub.ProofLeading()
	assert.Equal(t, ev, following.Datum.(kinds.Proof))

	const proofOctets int64 = 372
	evs, extent = hub.AwaitingProof(proofOctets)
	assert.Equal(t, 1, len(evs))
	assert.Equal(t, proofOctets, extent) //

	//
	assert.NoError(t, hub.AppendProof(ev))
	evs, _ = hub.AwaitingProof(fallbackProofMaximumOctets)
	assert.Equal(t, 1, len(evs))
}

//
func VerifyAppendLapsedProof(t *testing.T) {
	var (
		val                 = kinds.FreshSimulatePRV()
		altitude              = int64(30)
		statusDepot          = bootstrapAssessorStatus(val, altitude)
		proofDatastore          = dbm.FreshMemoryDatastore()
		ledgerDepot          = &simulations.LedgerDepot{}
		lapsedProofMoment = time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
		lapsedAltitude       = int64(2)
	)

	ledgerDepot.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(func(h int64) *kinds.LedgerSummary {
		if h == altitude || h == lapsedAltitude {
			return &kinds.LedgerSummary{Heading: kinds.Heading{Moment: fallbackProofMoment}}
		}
		return &kinds.LedgerSummary{Heading: kinds.Heading{Moment: lapsedProofMoment}}
	})

	hub, err := proof.FreshHub(proofDatastore, statusDepot, ledgerDepot)
	require.NoError(t, err)

	verifyScenarios := []struct {
		occurenceAltitude      int64
		occurenceMoment        time.Time
		expirationFault        bool
		occurenceCharacterization string
	}{
		{altitude, fallbackProofMoment, false, "REDACTED"},
		{lapsedAltitude, fallbackProofMoment, false, "REDACTED"},
		{altitude - 1, lapsedProofMoment, false, "REDACTED"},
		{
			lapsedAltitude - 1, lapsedProofMoment, true,
			"REDACTED",
		},
		{altitude, fallbackProofMoment.Add(1 * time.Minute), true, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.occurenceCharacterization, func(t *testing.T) {
			ev, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(tc.occurenceAltitude, tc.occurenceMoment, val, proofSuccessionUUID)
			require.NoError(t, err)
			err = hub.AppendProof(ev)
			if tc.expirationFault {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func VerifyRecordDiscordantBallots(t *testing.T) {
	var altitude int64 = 10

	hub, pv := fallbackVerifyHub(t, altitude)
	val := kinds.FreshAssessor(pv.PrivateToken.PublicToken(), 10)
	ev, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(altitude+1, fallbackProofMoment, pv, proofSuccessionUUID)
	require.NoError(t, err)

	hub.DiscloseDiscordantBallots(ev.BallotAN, ev.BallotBYTE)

	//
	hub.DiscloseDiscordantBallots(ev.BallotAN, ev.BallotBYTE)

	//
	occurenceCatalog, occurenceExtent := hub.AwaitingProof(fallbackProofMaximumOctets)
	require.Empty(t, occurenceCatalog)
	require.Zero(t, occurenceExtent)

	following := hub.ProofLeading()
	require.Nil(t, following)

	//
	status := hub.Status()
	status.FinalLedgerAltitude++
	status.FinalLedgerMoment = ev.Moment()
	status.FinalAssessors = kinds.FreshAssessorAssign([]*kinds.Assessor{val})
	hub.Revise(status, []kinds.Proof{})

	//
	occurenceCatalog, _ = hub.AwaitingProof(fallbackProofMaximumOctets)
	require.Equal(t, []kinds.Proof{ev}, occurenceCatalog)

	following = hub.ProofLeading()
	require.NotNil(t, following)
}

func VerifyProofHubRevise(t *testing.T) {
	altitude := int64(21)
	hub, val := fallbackVerifyHub(t, altitude)
	status := hub.Status()

	//
	trimmedOccurence, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(1, fallbackProofMoment.Add(1*time.Minute),
		val, proofSuccessionUUID)
	require.NoError(t, err)
	err = hub.AppendProof(trimmedOccurence)
	require.NoError(t, err)
	ev, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(altitude, fallbackProofMoment.Add(21*time.Minute),
		val, proofSuccessionUUID)
	require.NoError(t, err)
	finalAddnEndorse := createAddnEndorse(altitude, val.PrivateToken.PublicToken().Location())
	ledger := kinds.CreateLedger(altitude+1, []kinds.Tx{}, finalAddnEndorse.TowardEndorse(), []kinds.Proof{ev})
	//
	status.FinalLedgerAltitude = altitude + 1
	status.FinalLedgerMoment = fallbackProofMoment.Add(22 * time.Minute)
	err = hub.InspectProof(kinds.ProofCatalog{ev})
	require.NoError(t, err)

	hub.Revise(status, ledger.Proof.Proof)
	//
	occurenceCatalog, occurenceExtent := hub.AwaitingProof(fallbackProofMaximumOctets)
	assert.Empty(t, occurenceCatalog)
	assert.Zero(t, occurenceExtent)

	//
	err = hub.InspectProof(kinds.ProofCatalog{ev})
	if assert.Error(t, err) {
		assert.Equal(t, "REDACTED", err.(*kinds.FaultUnfitProof).Rationale.Error())
	}
}

func VerifyValidateAwaitingProofSucceeds(t *testing.T) {
	var altitude int64 = 1
	hub, val := fallbackVerifyHub(t, altitude)
	ev, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(altitude, fallbackProofMoment.Add(1*time.Minute),
		val, proofSuccessionUUID)
	require.NoError(t, err)
	err = hub.AppendProof(ev)
	require.NoError(t, err)

	err = hub.InspectProof(kinds.ProofCatalog{ev})
	assert.NoError(t, err)
}

func VerifyValidateReproducedProofCollapses(t *testing.T) {
	var altitude int64 = 1
	hub, val := fallbackVerifyHub(t, altitude)
	ev, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(altitude, fallbackProofMoment.Add(1*time.Minute),
		val, proofSuccessionUUID)
	require.NoError(t, err)
	err = hub.InspectProof(kinds.ProofCatalog{ev, ev})
	if assert.Error(t, err) {
		assert.Equal(t, "REDACTED", err.(*kinds.FaultUnfitProof).Rationale.Error())
	}
}

//
//
func VerifyAgileCustomerOnslaughtProofDuration(t *testing.T) {
	var (
		altitude       int64 = 100
		sharedAltitude int64 = 90
	)

	ev, reliable, shared := createInsaneProof(t, altitude, sharedAltitude,
		10, 5, 5, fallbackProofMoment, fallbackProofMoment.Add(1*time.Hour))

	status := sm.Status{
		FinalLedgerMoment:   fallbackProofMoment.Add(2 * time.Hour),
		FinalLedgerAltitude: 110,
		AgreementSettings: *kinds.FallbackAgreementSettings(),
	}
	statusDepot := &machinestubs.Depot{}
	statusDepot.On("REDACTED", altitude).Return(reliable.AssessorAssign, nil)
	statusDepot.On("REDACTED", sharedAltitude).Return(shared.AssessorAssign, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", altitude).Return(&kinds.LedgerSummary{Heading: *reliable.Heading})
	ledgerDepot.On("REDACTED", sharedAltitude).Return(&kinds.LedgerSummary{Heading: *shared.Heading})
	ledgerDepot.On("REDACTED", altitude).Return(reliable.Endorse)
	ledgerDepot.On("REDACTED", sharedAltitude).Return(shared.Endorse)

	hub, err := proof.FreshHub(dbm.FreshMemoryDatastore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	hub.AssignTracer(log.VerifyingTracer())

	err = hub.AppendProof(ev)
	assert.NoError(t, err)

	digest := ev.Digest()

	require.NoError(t, hub.AppendProof(ev))
	require.NoError(t, hub.AppendProof(ev))

	awaitingOccurence, _ := hub.AwaitingProof(status.AgreementSettings.Proof.MaximumOctets)
	require.Equal(t, 1, len(awaitingOccurence))
	require.Equal(t, ev, awaitingOccurence[0])

	require.NoError(t, hub.InspectProof(awaitingOccurence))
	require.Equal(t, ev, awaitingOccurence[0])

	status.FinalLedgerAltitude++
	status.FinalLedgerMoment = status.FinalLedgerMoment.Add(1 * time.Minute)
	hub.Revise(status, awaitingOccurence)
	require.Equal(t, digest, awaitingOccurence[0].Digest())

	outstandingOccurence, _ := hub.AwaitingProof(status.AgreementSettings.Proof.MaximumOctets)
	require.Empty(t, outstandingOccurence)

	//
	require.Error(t, hub.InspectProof(kinds.ProofCatalog{ev}))
	require.NoError(t, hub.AppendProof(ev))

	outstandingOccurence, _ = hub.AwaitingProof(status.AgreementSettings.Proof.MaximumOctets)
	require.Empty(t, outstandingOccurence)
}

//
//
func VerifyRestoreAwaitingProof(t *testing.T) {
	altitude := int64(10)
	val := kinds.FreshSimulatePRV()
	itemLocator := val.PrivateToken.PublicToken().Location()
	proofDatastore := dbm.FreshMemoryDatastore()
	statusDepot := bootstrapAssessorStatus(val, altitude)
	status, err := statusDepot.Fetch()
	require.NoError(t, err)
	ledgerDepot, err := bootstrapLedgerDepot(dbm.FreshMemoryDatastore(), status, itemLocator)
	require.NoError(t, err)
	//
	hub, err := proof.FreshHub(proofDatastore, statusDepot, ledgerDepot)
	require.NoError(t, err)
	hub.AssignTracer(log.VerifyingTracer())
	validProof, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(altitude,
		fallbackProofMoment.Add(10*time.Minute), val, proofSuccessionUUID)
	require.NoError(t, err)
	lapsedProof, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(int64(1),
		fallbackProofMoment.Add(1*time.Minute), val, proofSuccessionUUID)
	require.NoError(t, err)
	err = hub.AppendProof(validProof)
	require.NoError(t, err)
	err = hub.AppendProof(lapsedProof)
	require.NoError(t, err)

	//
	freshStatusDepot := &machinestubs.Depot{}
	freshStatusDepot.On("REDACTED").Return(sm.Status{
		FinalLedgerMoment:   fallbackProofMoment.Add(25 * time.Minute),
		FinalLedgerAltitude: altitude + 15,
		AgreementSettings: kinds.AgreementSettings{
			Ledger: kinds.LedgerParameters{
				MaximumOctets: 22020096,
				MaximumFuel:   -1,
			},
			Proof: kinds.ProofParameters{
				MaximumLifespanCountLedgers: 20,
				MaximumLifespanInterval:  20 * time.Minute,
				MaximumOctets:        fallbackProofMaximumOctets,
			},
		},
	}, nil)
	freshHub, err := proof.FreshHub(proofDatastore, freshStatusDepot, ledgerDepot)
	assert.NoError(t, err)
	occurenceCatalog, _ := freshHub.AwaitingProof(fallbackProofMaximumOctets)
	assert.Equal(t, 1, len(occurenceCatalog))
	following := freshHub.ProofLeading()
	assert.Equal(t, validProof, following.Datum.(kinds.Proof))
}

func bootstrapStatusOriginatingAssessorAssign(itemAssign *kinds.AssessorAssign, altitude int64) sm.Depot {
	statusDatastore := dbm.FreshMemoryDatastore()
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	status := sm.Status{
		SuccessionUUID:                     proofSuccessionUUID,
		PrimaryAltitude:               1,
		FinalLedgerAltitude:             altitude,
		FinalLedgerMoment:               fallbackProofMoment,
		Assessors:                  itemAssign,
		FollowingAssessors:              itemAssign.DuplicateAdvanceNominatorUrgency(1),
		FinalAssessors:              itemAssign,
		FinalAltitudeAssessorsAltered: 1,
		AgreementSettings: kinds.AgreementSettings{
			Ledger: kinds.LedgerParameters{
				MaximumOctets: 22020096,
				MaximumFuel:   -1,
			},
			Proof: kinds.ProofParameters{
				MaximumLifespanCountLedgers: 20,
				MaximumLifespanInterval:  20 * time.Minute,
				MaximumOctets:        1000,
			},
		},
	}

	//
	for i := int64(0); i <= altitude; i++ {
		status.FinalLedgerAltitude = i
		if err := statusDepot.Persist(status); err != nil {
			panic(err)
		}
	}

	return statusDepot
}

func bootstrapAssessorStatus(privateItem kinds.PrivateAssessor, altitude int64) sm.Depot {
	publicToken, _ := privateItem.ObtainPublicToken()
	assessor := &kinds.Assessor{Location: publicToken.Location(), BallotingPotency: 10, PublicToken: publicToken}

	//
	itemAssign := &kinds.AssessorAssign{
		Assessors: []*kinds.Assessor{assessor},
		Nominator:   assessor,
	}

	return bootstrapStatusOriginatingAssessorAssign(itemAssign, altitude)
}

//
//
func bootstrapLedgerDepot(db dbm.DB, status sm.Status, itemLocation []byte) (*depot.LedgerDepot, error) {
	ledgerDepot := depot.FreshLedgerDepot(db)

	for i := int64(1); i <= status.FinalLedgerAltitude; i++ {
		finalEndorse := createAddnEndorse(i-1, itemLocation)
		ledger, err := status.CreateLedger(i, verify.CreateNTHTrans(i, 1), finalEndorse.TowardEndorse(), nil, status.Assessors.Nominator.Location)
		if err != nil {
			return nil, err
		}
		ledger.Moment = fallbackProofMoment.Add(time.Duration(i) * time.Minute)
		ledger.Edition = strongmindedition.Agreement{Ledger: edition.LedgerScheme, App: 1}
		fragmentAssign, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
		if err != nil {
			return nil, err
		}

		observedEndorse := createAddnEndorse(i, itemLocation)
		ledgerDepot.PersistLedgerUsingExpandedEndorse(ledger, fragmentAssign, observedEndorse)
	}

	return ledgerDepot, nil
}

func createAddnEndorse(altitude int64, itemLocation []byte) *kinds.ExpandedEndorse {
	return &kinds.ExpandedEndorse{
		Altitude: altitude,
		ExpandedNotations: []kinds.ExpandedEndorseSignature{{
			EndorseSignature: kinds.EndorseSignature{
				LedgerUUIDMarker:      kinds.LedgerUUIDMarkerEndorse,
				AssessorLocation: itemLocation,
				Timestamp:        fallbackProofMoment,
				Notation:        []byte("REDACTED"),
			},
			AdditionNotation: []byte("REDACTED"),
		}},
	}
}

func fallbackVerifyHub(t *testing.T, altitude int64) (*proof.Hub, kinds.SimulatePRV) {
	t.Helper()
	val := kinds.FreshSimulatePRV()
	itemLocator := val.PrivateToken.PublicToken().Location()
	proofDatastore := dbm.FreshMemoryDatastore()
	statusDepot := bootstrapAssessorStatus(val, altitude)
	status, _ := statusDepot.Fetch()
	ledgerDepot, err := bootstrapLedgerDepot(dbm.FreshMemoryDatastore(), status, itemLocator)
	require.NoError(t, err)
	hub, err := proof.FreshHub(proofDatastore, statusDepot, ledgerDepot)
	if err != nil {
		panic("REDACTED")
	}
	hub.AssignTracer(log.VerifyingTracer())
	return hub, val
}

func generateStatus(altitude int64, itemAssign *kinds.AssessorAssign) sm.Status {
	return sm.Status{
		SuccessionUUID:         proofSuccessionUUID,
		FinalLedgerAltitude: altitude,
		FinalLedgerMoment:   fallbackProofMoment,
		Assessors:      itemAssign,
		AgreementSettings: *kinds.FallbackAgreementSettings(),
	}
}
