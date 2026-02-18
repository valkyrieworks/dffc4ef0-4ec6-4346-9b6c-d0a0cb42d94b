package proofs_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/proof"
	"github.com/valkyrieworks/proof/simulations"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	sm "github.com/valkyrieworks/status"
	smemulators "github.com/valkyrieworks/status/simulations"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

func VerifyMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

const proofSeriesUID = "REDACTED"

var (
	standardProofTime           = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	standardProofMaximumOctets int64 = 1000
)

func VerifyProofDepositorySimple(t *testing.T) {
	var (
		level     = int64(1)
		statusDepot = &smemulators.Depot{}
		proofStore = dbm.NewMemoryStore()
		ledgerDepot = &simulations.LedgerDepot{}
	)

	valueCollection, privateValues := kinds.RandomRatifierCollection(1, 10)

	ledgerDepot.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(
		&kinds.LedgerMeta{Heading: kinds.Heading{Time: standardProofTime}},
	)
	statusDepot.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(valueCollection, nil)
	statusDepot.On("REDACTED").Return(instantiateStatus(level+1, valueCollection), nil)

	depository, err := proof.NewDepository(proofStore, statusDepot, ledgerDepot)
	require.NoError(t, err)
	depository.AssignTracer(log.VerifyingTracer())

	//
	evs, volume := depository.AwaitingProof(standardProofMaximumOctets)
	assert.Equal(t, 0, len(evs))
	assert.Zero(t, volume)

	ev, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(level, standardProofTime, privateValues[0], proofSeriesUID)
	require.NoError(t, err)

	//
	evtAppended := make(chan struct{})
	go func() {
		<-depository.ProofWaitChan()
		close(evtAppended)
	}()

	//
	assert.NoError(t, depository.AppendProof(ev))

	select {
	case <-evtAppended:
	case <-time.After(5 * time.Second):
		t.Fatal("REDACTED")
	}

	following := depository.ProofHead()
	assert.Equal(t, ev, following.Item.(kinds.Proof))

	const proofOctets int64 = 372
	evs, volume = depository.AwaitingProof(proofOctets)
	assert.Equal(t, 1, len(evs))
	assert.Equal(t, proofOctets, volume) //

	//
	assert.NoError(t, depository.AppendProof(ev))
	evs, _ = depository.AwaitingProof(standardProofMaximumOctets)
	assert.Equal(t, 1, len(evs))
}

//
func VerifyAppendLapsedProof(t *testing.T) {
	var (
		val                 = kinds.NewEmulatePV()
		level              = int64(30)
		statusDepot          = bootstrapRatifierStatus(val, level)
		proofStore          = dbm.NewMemoryStore()
		ledgerDepot          = &simulations.LedgerDepot{}
		lapsedProofTime = time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
		lapsedLevel       = int64(2)
	)

	ledgerDepot.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(func(h int64) *kinds.LedgerMeta {
		if h == level || h == lapsedLevel {
			return &kinds.LedgerMeta{Heading: kinds.Heading{Time: standardProofTime}}
		}
		return &kinds.LedgerMeta{Heading: kinds.Heading{Time: lapsedProofTime}}
	})

	depository, err := proof.NewDepository(proofStore, statusDepot, ledgerDepot)
	require.NoError(t, err)

	verifyScenarios := []struct {
		evtLevel      int64
		evtTime        time.Time
		expirationErr        bool
		evtSummary string
	}{
		{level, standardProofTime, false, "REDACTED"},
		{lapsedLevel, standardProofTime, false, "REDACTED"},
		{level - 1, lapsedProofTime, false, "REDACTED"},
		{
			lapsedLevel - 1, lapsedProofTime, true,
			"REDACTED",
		},
		{level, standardProofTime.Add(1 * time.Minute), true, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.evtSummary, func(t *testing.T) {
			ev, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(tc.evtLevel, tc.evtTime, val, proofSeriesUID)
			require.NoError(t, err)
			err = depository.AppendProof(ev)
			if tc.expirationErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func VerifyNotifyClashingBallots(t *testing.T) {
	var level int64 = 10

	depository, pv := standardVerifyDepository(t, level)
	val := kinds.NewRatifier(pv.PrivateKey.PublicKey(), 10)
	ev, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(level+1, standardProofTime, pv, proofSeriesUID)
	require.NoError(t, err)

	depository.NotifyClashingBallots(ev.BallotA, ev.BallotBYTE)

	//
	depository.NotifyClashingBallots(ev.BallotA, ev.BallotBYTE)

	//
	evtCatalog, evtVolume := depository.AwaitingProof(standardProofMaximumOctets)
	require.Empty(t, evtCatalog)
	require.Zero(t, evtVolume)

	following := depository.ProofHead()
	require.Nil(t, following)

	//
	status := depository.Status()
	status.FinalLedgerLevel++
	status.FinalLedgerTime = ev.Time()
	status.FinalRatifiers = kinds.NewRatifierCollection([]*kinds.Ratifier{val})
	depository.Modify(status, []kinds.Proof{})

	//
	evtCatalog, _ = depository.AwaitingProof(standardProofMaximumOctets)
	require.Equal(t, []kinds.Proof{ev}, evtCatalog)

	following = depository.ProofHead()
	require.NotNil(t, following)
}

func VerifyProofDepositoryModify(t *testing.T) {
	level := int64(21)
	depository, val := standardVerifyDepository(t, level)
	status := depository.Status()

	//
	trimmedEvt, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(1, standardProofTime.Add(1*time.Minute),
		val, proofSeriesUID)
	require.NoError(t, err)
	err = depository.AppendProof(trimmedEvt)
	require.NoError(t, err)
	ev, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(level, standardProofTime.Add(21*time.Minute),
		val, proofSeriesUID)
	require.NoError(t, err)
	finalExtensionEndorse := createExtensionEndorse(level, val.PrivateKey.PublicKey().Location())
	ledger := kinds.CreateLedger(level+1, []kinds.Tx{}, finalExtensionEndorse.ToEndorse(), []kinds.Proof{ev})
	//
	status.FinalLedgerLevel = level + 1
	status.FinalLedgerTime = standardProofTime.Add(22 * time.Minute)
	err = depository.InspectProof(kinds.ProofCatalog{ev})
	require.NoError(t, err)

	depository.Modify(status, ledger.Proof.Proof)
	//
	evtCatalog, evtVolume := depository.AwaitingProof(standardProofMaximumOctets)
	assert.Empty(t, evtCatalog)
	assert.Zero(t, evtVolume)

	//
	err = depository.InspectProof(kinds.ProofCatalog{ev})
	if assert.Error(t, err) {
		assert.Equal(t, "REDACTED", err.(*kinds.ErrCorruptProof).Cause.Error())
	}
}

func VerifyValidateAwaitingProofSucceeds(t *testing.T) {
	var level int64 = 1
	depository, val := standardVerifyDepository(t, level)
	ev, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(level, standardProofTime.Add(1*time.Minute),
		val, proofSeriesUID)
	require.NoError(t, err)
	err = depository.AppendProof(ev)
	require.NoError(t, err)

	err = depository.InspectProof(kinds.ProofCatalog{ev})
	assert.NoError(t, err)
}

func VerifyValidateReplicatedProofErrors(t *testing.T) {
	var level int64 = 1
	depository, val := standardVerifyDepository(t, level)
	ev, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(level, standardProofTime.Add(1*time.Minute),
		val, proofSeriesUID)
	require.NoError(t, err)
	err = depository.InspectProof(kinds.ProofCatalog{ev, ev})
	if assert.Error(t, err) {
		assert.Equal(t, "REDACTED", err.(*kinds.ErrCorruptProof).Cause.Error())
	}
}

//
//
func VerifyRapidCustomerAssaultProofDuration(t *testing.T) {
	var (
		level       int64 = 100
		sharedLevel int64 = 90
	)

	ev, validated, shared := createErraticProof(t, level, sharedLevel,
		10, 5, 5, standardProofTime, standardProofTime.Add(1*time.Hour))

	status := sm.Status{
		FinalLedgerTime:   standardProofTime.Add(2 * time.Hour),
		FinalLedgerLevel: 110,
		AgreementOptions: *kinds.StandardAgreementOptions(),
	}
	statusDepot := &smemulators.Depot{}
	statusDepot.On("REDACTED", level).Return(validated.RatifierAssign, nil)
	statusDepot.On("REDACTED", sharedLevel).Return(shared.RatifierAssign, nil)
	statusDepot.On("REDACTED").Return(status, nil)
	ledgerDepot := &simulations.LedgerDepot{}
	ledgerDepot.On("REDACTED", level).Return(&kinds.LedgerMeta{Heading: *validated.Heading})
	ledgerDepot.On("REDACTED", sharedLevel).Return(&kinds.LedgerMeta{Heading: *shared.Heading})
	ledgerDepot.On("REDACTED", level).Return(validated.Endorse)
	ledgerDepot.On("REDACTED", sharedLevel).Return(shared.Endorse)

	depository, err := proof.NewDepository(dbm.NewMemoryStore(), statusDepot, ledgerDepot)
	require.NoError(t, err)
	depository.AssignTracer(log.VerifyingTracer())

	err = depository.AppendProof(ev)
	assert.NoError(t, err)

	digest := ev.Digest()

	require.NoError(t, depository.AppendProof(ev))
	require.NoError(t, depository.AppendProof(ev))

	awaitingEvt, _ := depository.AwaitingProof(status.AgreementOptions.Proof.MaximumOctets)
	require.Equal(t, 1, len(awaitingEvt))
	require.Equal(t, ev, awaitingEvt[0])

	require.NoError(t, depository.InspectProof(awaitingEvt))
	require.Equal(t, ev, awaitingEvt[0])

	status.FinalLedgerLevel++
	status.FinalLedgerTime = status.FinalLedgerTime.Add(1 * time.Minute)
	depository.Modify(status, awaitingEvt)
	require.Equal(t, digest, awaitingEvt[0].Digest())

	outstandingEvt, _ := depository.AwaitingProof(status.AgreementOptions.Proof.MaximumOctets)
	require.Empty(t, outstandingEvt)

	//
	require.Error(t, depository.InspectProof(kinds.ProofCatalog{ev}))
	require.NoError(t, depository.AppendProof(ev))

	outstandingEvt, _ = depository.AwaitingProof(status.AgreementOptions.Proof.MaximumOctets)
	require.Empty(t, outstandingEvt)
}

//
//
func VerifyRestoreAwaitingProof(t *testing.T) {
	level := int64(10)
	val := kinds.NewEmulatePV()
	valueLocation := val.PrivateKey.PublicKey().Location()
	proofStore := dbm.NewMemoryStore()
	statusDepot := bootstrapRatifierStatus(val, level)
	status, err := statusDepot.Import()
	require.NoError(t, err)
	ledgerDepot, err := bootstrapLedgerDepot(dbm.NewMemoryStore(), status, valueLocation)
	require.NoError(t, err)
	//
	depository, err := proof.NewDepository(proofStore, statusDepot, ledgerDepot)
	require.NoError(t, err)
	depository.AssignTracer(log.VerifyingTracer())
	validProof, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(level,
		standardProofTime.Add(10*time.Minute), val, proofSeriesUID)
	require.NoError(t, err)
	lapsedProof, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(int64(1),
		standardProofTime.Add(1*time.Minute), val, proofSeriesUID)
	require.NoError(t, err)
	err = depository.AppendProof(validProof)
	require.NoError(t, err)
	err = depository.AppendProof(lapsedProof)
	require.NoError(t, err)

	//
	newStatusDepot := &smemulators.Depot{}
	newStatusDepot.On("REDACTED").Return(sm.Status{
		FinalLedgerTime:   standardProofTime.Add(25 * time.Minute),
		FinalLedgerLevel: level + 15,
		AgreementOptions: kinds.AgreementOptions{
			Ledger: kinds.LedgerOptions{
				MaximumOctets: 22020096,
				MaximumFuel:   -1,
			},
			Proof: kinds.ProofOptions{
				MaximumDurationCountLedgers: 20,
				MaximumDurationPeriod:  20 * time.Minute,
				MaximumOctets:        standardProofMaximumOctets,
			},
		},
	}, nil)
	newDepository, err := proof.NewDepository(proofStore, newStatusDepot, ledgerDepot)
	assert.NoError(t, err)
	evtCatalog, _ := newDepository.AwaitingProof(standardProofMaximumOctets)
	assert.Equal(t, 1, len(evtCatalog))
	following := newDepository.ProofHead()
	assert.Equal(t, validProof, following.Item.(kinds.Proof))
}

func bootstrapStatusFromRatifierCollection(valueCollection *kinds.RatifierAssign, level int64) sm.Depot {
	statusStore := dbm.NewMemoryStore()
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	status := sm.Status{
		LedgerUID:                     proofSeriesUID,
		PrimaryLevel:               1,
		FinalLedgerLevel:             level,
		FinalLedgerTime:               standardProofTime,
		Ratifiers:                  valueCollection,
		FollowingRatifiers:              valueCollection.CloneAugmentRecommenderUrgency(1),
		FinalRatifiers:              valueCollection,
		FinalLevelRatifiersModified: 1,
		AgreementOptions: kinds.AgreementOptions{
			Ledger: kinds.LedgerOptions{
				MaximumOctets: 22020096,
				MaximumFuel:   -1,
			},
			Proof: kinds.ProofOptions{
				MaximumDurationCountLedgers: 20,
				MaximumDurationPeriod:  20 * time.Minute,
				MaximumOctets:        1000,
			},
		},
	}

	//
	for i := int64(0); i <= level; i++ {
		status.FinalLedgerLevel = i
		if err := statusDepot.Persist(status); err != nil {
			panic(err)
		}
	}

	return statusDepot
}

func bootstrapRatifierStatus(privateValue kinds.PrivateRatifier, level int64) sm.Depot {
	publicKey, _ := privateValue.FetchPublicKey()
	ratifier := &kinds.Ratifier{Location: publicKey.Location(), PollingEnergy: 10, PublicKey: publicKey}

	//
	valueCollection := &kinds.RatifierAssign{
		Ratifiers: []*kinds.Ratifier{ratifier},
		Recommender:   ratifier,
	}

	return bootstrapStatusFromRatifierCollection(valueCollection, level)
}

//
//
func bootstrapLedgerDepot(db dbm.DB, status sm.Status, valueAddress []byte) (*depot.LedgerDepot, error) {
	ledgerDepot := depot.NewLedgerDepot(db)

	for i := int64(1); i <= status.FinalLedgerLevel; i++ {
		finalEndorse := createExtensionEndorse(i-1, valueAddress)
		ledger, err := status.CreateLedger(i, verify.CreateNTrans(i, 1), finalEndorse.ToEndorse(), nil, status.Ratifiers.Recommender.Location)
		if err != nil {
			return nil, err
		}
		ledger.Time = standardProofTime.Add(time.Duration(i) * time.Minute)
		ledger.Release = cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 1}
		sectionCollection, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
		if err != nil {
			return nil, err
		}

		viewedEndorse := createExtensionEndorse(i, valueAddress)
		ledgerDepot.PersistLedgerWithExpandedEndorse(ledger, sectionCollection, viewedEndorse)
	}

	return ledgerDepot, nil
}

func createExtensionEndorse(level int64, valueAddress []byte) *kinds.ExpandedEndorse {
	return &kinds.ExpandedEndorse{
		Level: level,
		ExpandedEndorsements: []kinds.ExpandedEndorseSignature{{
			EndorseSignature: kinds.EndorseSignature{
				LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
				RatifierLocation: valueAddress,
				Timestamp:        standardProofTime,
				Autograph:        []byte("REDACTED"),
			},
			AdditionAutograph: []byte("REDACTED"),
		}},
	}
}

func standardVerifyDepository(t *testing.T, level int64) (*proof.Depository, kinds.EmulatePV) {
	t.Helper()
	val := kinds.NewEmulatePV()
	valueLocation := val.PrivateKey.PublicKey().Location()
	proofStore := dbm.NewMemoryStore()
	statusDepot := bootstrapRatifierStatus(val, level)
	status, _ := statusDepot.Import()
	ledgerDepot, err := bootstrapLedgerDepot(dbm.NewMemoryStore(), status, valueLocation)
	require.NoError(t, err)
	depository, err := proof.NewDepository(proofStore, statusDepot, ledgerDepot)
	if err != nil {
		panic("REDACTED")
	}
	depository.AssignTracer(log.VerifyingTracer())
	return depository, val
}

func instantiateStatus(level int64, valueCollection *kinds.RatifierAssign) sm.Status {
	return sm.Status{
		LedgerUID:         proofSeriesUID,
		FinalLedgerLevel: level,
		FinalLedgerTime:   standardProofTime,
		Ratifiers:      valueCollection,
		AgreementOptions: *kinds.StandardAgreementOptions(),
	}
}
