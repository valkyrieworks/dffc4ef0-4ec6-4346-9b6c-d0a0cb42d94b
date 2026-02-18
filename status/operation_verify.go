package status_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	ifaceclientmocks "github.com/valkyrieworks/iface/customer/simulations"
	iface "github.com/valkyrieworks/iface/kinds"
	abciemulators "github.com/valkyrieworks/iface/kinds/simulations"
	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"
	txpoolsims "github.com/valkyrieworks/txpool/simulations"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/gateway"
	omocks "github.com/valkyrieworks/gateway/simulations"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/simulations"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
	"github.com/valkyrieworks/release"
)

var (
	ledgerUID             = "REDACTED"
	verifySegmentVolume uint32 = kinds.LedgerSegmentVolumeOctets
)

func VerifyExecuteLedger(t *testing.T) {
	app := &verifyApplication{}
	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.Nil(t, err)
	defer gatewayApplication.Halt() //

	status, statusStore, _ := createStatus(1, 1)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())

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
	ledgerExecute := sm.NewLedgerRunner(statusDepot, log.VerifyingTracer(), gatewayApplication.Agreement(),
		mp, sm.EmptyProofDepository{}, ledgerDepot)

	ledger, err := createLedger(status, 1, new(kinds.Endorse))
	require.NoError(t, err)
	bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
	require.NoError(t, err)
	ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}

	status, err = ledgerExecute.ExecuteLedger(status, ledgerUID, ledger)
	require.Nil(t, err)

	//
	assert.EqualValues(t, 1, status.Release.Agreement.App, "REDACTED")
}

//
//
//
//
func VerifyCompleteLedgerResolvedFinalEndorse(t *testing.T) {
	app := &verifyApplication{}
	rootTime := time.Now()
	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	status, statusStore, privateValues := createStatus(7, 1)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	missingSignature := kinds.NewExpandedEndorseSignatureMissing()

	verifyScenarios := []struct {
		label             string
		missingEndorseAutographs map[int]bool
	}{
		{"REDACTED", map[int]bool{}},
		{"REDACTED", map[int]bool{1: true}},
		{"REDACTED", map[int]bool{1: true, 3: true}},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.label, func(t *testing.T) {
			ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
			eventpool := &simulations.ProofDepository{}
			eventpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, 0)
			eventpool.On("REDACTED", mock.Anything, mock.Anything).Return()
			eventpool.On("REDACTED", mock.Anything).Return(nil)
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
				mock.Anything,
				mock.Anything).Return(nil)

			eventBus := kinds.NewEventBus()
			require.NoError(t, eventBus.Begin())

			ledgerExecute := sm.NewLedgerRunner(statusDepot, log.NewNoopTracer(), gatewayApplication.Agreement(), mp, eventpool, ledgerDepot)
			status, _, finalEndorse, err := createAndEndorseValidLedger(status, 1, new(kinds.Endorse), status.FollowingRatifiers.Ratifiers[0].Location, ledgerExecute, privateValues, nil)
			require.NoError(t, err)

			for idx, isMissing := range tc.missingEndorseAutographs {
				if isMissing {
					finalEndorse.ExpandedEndorsements[idx] = missingSignature
				}
			}

			//
			ledger, err := createLedger(status, 2, finalEndorse.ToEndorse())
			require.NoError(t, err)
			bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
			require.NoError(t, err)
			ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}
			_, err = ledgerExecute.ExecuteLedger(status, ledgerUID, ledger)
			require.NoError(t, err)
			require.True(t, app.FinalTime.After(rootTime))

			//
			for i, v := range app.EndorseBallots {
				_, missing := tc.missingEndorseAutographs[i]
				assert.Equal(t, !missing, v.LedgerUidMark != engineproto.LedgerUIDMarkMissing)
			}
		})
	}
}

//
func VerifyCompleteLedgerRatifiers(t *testing.T) {
	app := &verifyApplication{}
	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	status, statusStore, _ := createStatus(2, 2)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	previousDigest := status.FinalLedgerUID.Digest
	previousSections := kinds.SegmentAssignHeading{}
	previousLedgerUID := kinds.LedgerUID{Digest: previousDigest, SegmentAssignHeading: previousSections}

	var (
		now        = engineclock.Now()
		endorseSignature0 = kinds.ExpandedEndorseSignature{
			EndorseSignature: kinds.EndorseSignature{
				LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
				RatifierLocation: status.Ratifiers.Ratifiers[0].Location,
				Timestamp:        now,
				Autograph:        []byte("REDACTED"),
			},
			Addition:          []byte("REDACTED"),
			AdditionAutograph: []byte("REDACTED"),
		}

		endorseSignature1 = kinds.ExpandedEndorseSignature{
			EndorseSignature: kinds.EndorseSignature{
				LedgerUIDMark:      kinds.LedgerUIDMarkEndorse,
				RatifierLocation: status.Ratifiers.Ratifiers[1].Location,
				Timestamp:        now,
				Autograph:        []byte("REDACTED"),
			},
			Addition:          []byte("REDACTED"),
			AdditionAutograph: []byte("REDACTED"),
		}
		missingSignature = kinds.NewExpandedEndorseSignatureMissing()
	)

	verifyScenarios := []struct {
		note                     string
		finalEndorseAutographs           []kinds.ExpandedEndorseSignature
		anticipatedMissingRatifiers []int
		mustPossessTime           bool
	}{
		{"REDACTED", []kinds.ExpandedEndorseSignature{endorseSignature0, endorseSignature1}, []int{}, true},
		{"REDACTED", []kinds.ExpandedEndorseSignature{endorseSignature0, missingSignature}, []int{1}, true},
		{"REDACTED", []kinds.ExpandedEndorseSignature{missingSignature, missingSignature}, []int{0, 1}, false},
	}

	for _, tc := range verifyScenarios {
		finalEndorse := &kinds.ExpandedEndorse{
			Level:             1,
			LedgerUID:            previousLedgerUID,
			ExpandedEndorsements: tc.finalEndorseAutographs,
		}

		//
		ledger, err := createLedger(status, 2, finalEndorse.ToEndorse())
		require.NoError(t, err)

		_, err = sm.InvokeEndorseLedger(gatewayApplication.Agreement(), ledger, log.VerifyingTracer(), statusDepot, 1)
		require.NoError(t, err, tc.note)
		require.True(t,
			!tc.mustPossessTime ||
				app.FinalTime.Equal(now) || app.FinalTime.After(now),
			"REDACTED", tc.note, app.FinalTime, now,
		)

		//
		ctr := 0
		for i, v := range app.EndorseBallots {
			if ctr < len(tc.anticipatedMissingRatifiers) &&
				tc.anticipatedMissingRatifiers[ctr] == i {

				assert.Equal(t, v.LedgerUidMark, engineproto.LedgerUIDMarkMissing)
				ctr++
			} else {
				assert.NotEqual(t, v.LedgerUidMark, engineproto.LedgerUIDMarkMissing)
			}
		}
	}
}

//
func VerifyCompleteLedgerMalpractice(t *testing.T) {
	app := &verifyApplication{}
	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	status, statusStore, privateValues := createStatus(1, 1)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	standardProofTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	privateValue := privateValues[status.Ratifiers.Ratifiers[0].Location.String()]
	ledgerUID := createLedgerUID([]byte("REDACTED"), 1000, []byte("REDACTED"))
	heading := &kinds.Heading{
		Release:            cometrelease.Agreement{Ledger: release.LedgerProtocol, App: 1},
		LedgerUID:            status.LedgerUID,
		Level:             10,
		Time:               standardProofTime,
		FinalLedgerUID:        ledgerUID,
		FinalEndorseDigest:     vault.CRandomOctets(comethash.Volume),
		DataDigest:           vault.CRandomOctets(comethash.Volume),
		RatifiersDigest:     status.Ratifiers.Digest(),
		FollowingRatifiersDigest: status.Ratifiers.Digest(),
		AgreementDigest:      vault.CRandomOctets(comethash.Volume),
		ApplicationDigest:            vault.CRandomOctets(comethash.Volume),
		FinalOutcomesDigest:    vault.CRandomOctets(comethash.Volume),
		ProofDigest:       vault.CRandomOctets(comethash.Volume),
		RecommenderLocation:    vault.CRandomOctets(vault.LocationVolume),
	}

	//
	dve, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(3, standardProofTime, privateValue, status.LedgerUID)
	require.NoError(t, err)
	dve.RatifierEnergy = 1000
	rce := &kinds.RapidCustomerAssaultProof{
		ClashingLedger: &kinds.RapidLedger{
			AttestedHeading: &kinds.AttestedHeading{
				Heading: heading,
				Endorse: &kinds.Endorse{
					Level:  10,
					LedgerUID: createLedgerUID(heading.Digest(), 100, []byte("REDACTED")),
					Endorsements: []kinds.EndorseSignature{{
						LedgerUIDMark:      kinds.LedgerUIDMarkNull,
						RatifierLocation: vault.LocationDigest([]byte("REDACTED")),
						Timestamp:        standardProofTime,
						Autograph:        vault.CRandomOctets(kinds.MaximumAutographVolume),
					}},
				},
			},
			RatifierAssign: status.Ratifiers,
		},
		SharedLevel:        8,
		FaultyRatifiers: []*kinds.Ratifier{status.Ratifiers.Ratifiers[0]},
		SumPollingEnergy:    12,
		Timestamp:           standardProofTime,
	}

	ev := []kinds.Proof{dve, rce}

	ifaceMegabyte := []iface.Malpractice{
		{
			Kind:             iface.Misconductkind_REPLICATED_BALLOT,
			Level:           3,
			Time:             standardProofTime,
			Ratifier:        kinds.Tm2schema.Ratifier(status.Ratifiers.Ratifiers[0]),
			SumPollingEnergy: 10,
		},
		{
			Kind:             iface.Misconductkind_RAPID_CUSTOMER_ASSAULT,
			Level:           8,
			Time:             standardProofTime,
			Ratifier:        kinds.Tm2schema.Ratifier(status.Ratifiers.Ratifiers[0]),
			SumPollingEnergy: 12,
		},
	}

	eventpool := &simulations.ProofDepository{}
	eventpool.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(ev, int64(100))
	eventpool.On("REDACTED", mock.AnythingOfType("REDACTED"), mock.AnythingOfType("REDACTED")).Return()
	eventpool.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)
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

	ledgerExecute := sm.NewLedgerRunner(statusDepot, log.VerifyingTracer(), gatewayApplication.Agreement(),
		mp, eventpool, ledgerDepot)

	ledger, err := createLedger(status, 1, new(kinds.Endorse))
	require.NoError(t, err)
	ledger.Proof = kinds.ProofData{Proof: ev}
	ledger.ProofDigest = ledger.Proof.Digest()
	bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
	require.NoError(t, err)

	ledgerUID = kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}

	_, err = ledgerExecute.ExecuteLedger(status, ledgerUID, ledger)
	require.NoError(t, err)

	//
	assert.Equal(t, ifaceMegabyte, app.Malpractice)
}

func VerifyHandleNomination(t *testing.T) {
	const level = 2
	txs := verify.CreateNTrans(level, 10)

	tracer := log.NewNoopTracer()
	app := &abciemulators.Software{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyHandleNomination{Status: iface.Responseprocessnomination_ALLOW}, nil)

	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	status, statusStore, privateValues := createStatus(1, level)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	eventBus := kinds.NewEventBus()
	err = eventBus.Begin()
	require.NoError(t, err)

	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		tracer,
		gatewayApplication.Agreement(),
		new(txpoolsims.Txpool),
		sm.EmptyProofDepository{},
		ledgerDepot,
	)

	ledger0, err := createLedger(status, level-1, new(kinds.Endorse))
	require.NoError(t, err)
	finalEndorseSignature := []kinds.EndorseSignature{}
	sectionCollection, err := ledger0.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)
	ledgerUID := kinds.LedgerUID{Digest: ledger0.Digest(), SegmentAssignHeading: sectionCollection.Heading()}
	ballotDetails := []iface.BallotDetails{}
	for _, privateValue := range privateValues {
		pk, err := privateValue.FetchPublicKey()
		require.NoError(t, err)
		idx, _ := status.Ratifiers.FetchByLocation(pk.Location())
		ballot := kinds.CreateBallotNoFault(t, privateValue, ledger0.LedgerUID, idx, level-1, 0, 2, ledgerUID, time.Now())
		address := pk.Location()
		ballotDetails = append(ballotDetails,
			iface.BallotDetails{
				LedgerUidMark: engineproto.LedgerUIDMarkEndorse,
				Ratifier: iface.Ratifier{
					Location: address,
					Energy:   1000,
				},
			})
		finalEndorseSignature = append(finalEndorseSignature, ballot.EndorseSignature())
	}

	ledger1, err := createLedger(status, level, &kinds.Endorse{
		Level:     level - 1,
		Endorsements: finalEndorseSignature,
	})
	require.NoError(t, err)

	ledger1.Txs = txs

	anticipatedRec := &iface.QueryHandleNomination{
		Txs:         ledger1.Txs.ToSegmentOfOctets(),
		Digest:        ledger1.Digest(),
		Level:      ledger1.Level,
		Time:        ledger1.Time,
		Malpractice: ledger1.Proof.Proof.ToIface(),
		NominatedFinalEndorse: iface.EndorseDetails{
			Cycle: 0,
			Ballots: ballotDetails,
		},
		FollowingRatifiersDigest: ledger1.FollowingRatifiersDigest,
		RecommenderLocation:    ledger1.RecommenderLocation,
	}

	allowLedger, err := ledgerExecute.HandleNomination(ledger1, status)
	require.NoError(t, err)
	require.True(t, allowLedger)
	app.AssertExpectations(t)
	app.AssertCalled(t, "REDACTED", context.TODO(), anticipatedRec)
}

func VerifyCertifyRatifierRefreshes(t *testing.T) {
	publickey1 := ed25519.GeneratePrivateKey().PublicKey()
	publickey2 := ed25519.GeneratePrivateKey().PublicKey()
	pk1, err := cryptocode.PublicKeyToSchema(publickey1)
	assert.NoError(t, err)
	pk2, err := cryptocode.PublicKeyToSchema(publickey2)
	assert.NoError(t, err)

	standardRatifierOptions := kinds.RatifierOptions{PublicKeyKinds: []string{kinds.IfacePublicKeyKindEd25519}}

	verifyScenarios := []struct {
		label string

		ifaceRefreshes     []iface.RatifierModify
		ratifierOptions kinds.RatifierOptions

		mustErr bool
	}{
		{
			"REDACTED",
			[]iface.RatifierModify{{PublicKey: pk2, Energy: 20}},
			standardRatifierOptions,
			false,
		},
		{
			"REDACTED",
			[]iface.RatifierModify{{PublicKey: pk1, Energy: 20}},
			standardRatifierOptions,
			false,
		},
		{
			"REDACTED",
			[]iface.RatifierModify{{PublicKey: pk2, Energy: 0}},
			standardRatifierOptions,
			false,
		},
		{
			"REDACTED",
			[]iface.RatifierModify{{PublicKey: pk2, Energy: -100}},
			standardRatifierOptions,
			true,
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.label, func(t *testing.T) {
			err := sm.CertifyRatifierRefreshes(tc.ifaceRefreshes, tc.ratifierOptions)
			if tc.mustErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func VerifyModifyRatifiers(t *testing.T) {
	publickey1 := ed25519.GeneratePrivateKey().PublicKey()
	value1 := kinds.NewRatifier(publickey1, 10)
	publickey2 := ed25519.GeneratePrivateKey().PublicKey()
	value2 := kinds.NewRatifier(publickey2, 20)

	pk, err := cryptocode.PublicKeyToSchema(publickey1)
	require.NoError(t, err)
	pk2, err := cryptocode.PublicKeyToSchema(publickey2)
	require.NoError(t, err)

	verifyScenarios := []struct {
		label string

		ongoingCollection  *kinds.RatifierAssign
		ifaceRefreshes []iface.RatifierModify

		ensuingCollection *kinds.RatifierAssign
		mustErr    bool
	}{
		{
			"REDACTED",
			kinds.NewRatifierCollection([]*kinds.Ratifier{value1}),
			[]iface.RatifierModify{{PublicKey: pk2, Energy: 20}},
			kinds.NewRatifierCollection([]*kinds.Ratifier{value1, value2}),
			false,
		},
		{
			"REDACTED",
			kinds.NewRatifierCollection([]*kinds.Ratifier{value1}),
			[]iface.RatifierModify{{PublicKey: pk, Energy: 20}},
			kinds.NewRatifierCollection([]*kinds.Ratifier{kinds.NewRatifier(publickey1, 20)}),
			false,
		},
		{
			"REDACTED",
			kinds.NewRatifierCollection([]*kinds.Ratifier{value1, value2}),
			[]iface.RatifierModify{{PublicKey: pk2, Energy: 0}},
			kinds.NewRatifierCollection([]*kinds.Ratifier{value1}),
			false,
		},
		{
			"REDACTED",
			kinds.NewRatifierCollection([]*kinds.Ratifier{value1}),
			[]iface.RatifierModify{{PublicKey: pk2, Energy: 0}},
			kinds.NewRatifierCollection([]*kinds.Ratifier{value1}),
			true,
		},
	}

	for _, tc := range verifyScenarios {

		t.Run(tc.label, func(t *testing.T) {
			refreshes, err := kinds.Schema2tm.RatifierRefreshes(tc.ifaceRefreshes)
			assert.NoError(t, err)
			err = tc.ongoingCollection.ModifyWithAlterCollection(refreshes)
			if tc.mustErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				require.Equal(t, tc.ensuingCollection.Volume(), tc.ongoingCollection.Volume())

				assert.Equal(t, tc.ensuingCollection.SumPollingEnergy(), tc.ongoingCollection.SumPollingEnergy())

				assert.Equal(t, tc.ensuingCollection.Ratifiers[0].Location, tc.ongoingCollection.Ratifiers[0].Location)
				if tc.ensuingCollection.Volume() > 1 {
					assert.Equal(t, tc.ensuingCollection.Ratifiers[1].Location, tc.ongoingCollection.Ratifiers[1].Location)
				}
			}
		})
	}
}

//
func VerifyCompleteLedgerRatifierRefreshes(t *testing.T) {
	app := &verifyApplication{}
	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	status, statusStore, _ := createStatus(1, 1)
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
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(kinds.Txs{})

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.VerifyingTracer(),
		gatewayApplication.Agreement(),
		mp,
		sm.EmptyProofDepository{},
		ledgerDepot,
	)

	eventBus := kinds.NewEventBus()
	err = eventBus.Begin()
	require.NoError(t, err)
	defer eventBus.Halt() //

	ledgerExecute.AssignEventBus(eventBus)

	refreshesSubtract, err := eventBus.Enrol(
		context.Background(),
		"REDACTED",
		kinds.EventInquireRatifierCollectionRefreshes,
	)
	require.NoError(t, err)

	ledger, err := createLedger(status, 1, new(kinds.Endorse))
	require.NoError(t, err)
	bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
	require.NoError(t, err)
	ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}

	publickey := ed25519.GeneratePrivateKey().PublicKey()
	pk, err := cryptocode.PublicKeyToSchema(publickey)
	require.NoError(t, err)
	app.RatifierRefreshes = []iface.RatifierModify{
		{PublicKey: pk, Energy: 10},
	}

	status, err = ledgerExecute.ExecuteLedger(status, ledgerUID, ledger)
	require.NoError(t, err)
	//
	if assert.Equal(t, status.Ratifiers.Volume()+1, status.FollowingRatifiers.Volume()) {
		idx, _ := status.FollowingRatifiers.FetchByLocation(publickey.Location())
		if idx < 0 {
			t.Fatalf("REDACTED", publickey.Location(), status.FollowingRatifiers)
		}
	}

	//
	select {
	case msg := <-refreshesSubtract.Out():
		event, ok := msg.Data().(kinds.EventDataRatifierCollectionRefreshes)
		require.True(t, ok, "REDACTED", msg.Data())
		if assert.NotEmpty(t, event.RatifierRefreshes) {
			assert.Equal(t, publickey, event.RatifierRefreshes[0].PublicKey)
			assert.EqualValues(t, 10, event.RatifierRefreshes[0].PollingEnergy)
		}
	case <-refreshesSubtract.Revoked():
		t.Fatalf("REDACTED", refreshesSubtract.Err())
	case <-time.After(1 * time.Second):
		t.Fatal("REDACTED")
	}
}

//
//
func VerifyCompleteLedgerRatifierRefreshesEnsuingInEmptyCollection(t *testing.T) {
	app := &verifyApplication{}
	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	status, statusStore, _ := createStatus(1, 1)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.VerifyingTracer(),
		gatewayApplication.Agreement(),
		new(txpoolsims.Txpool),
		sm.EmptyProofDepository{},
		ledgerDepot,
	)

	ledger, err := createLedger(status, 1, new(kinds.Endorse))
	require.NoError(t, err)
	bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
	require.NoError(t, err)
	ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}

	vp, err := cryptocode.PublicKeyToSchema(status.Ratifiers.Ratifiers[0].PublicKey)
	require.NoError(t, err)
	//
	app.RatifierRefreshes = []iface.RatifierModify{
		{PublicKey: vp, Energy: 0},
	}

	assert.NotPanics(t, func() { status, err = ledgerExecute.ExecuteLedger(status, ledgerUID, ledger) })
	assert.Error(t, err)
	assert.NotEmpty(t, status.FollowingRatifiers.Ratifiers)
}

func VerifyEmptyArrangeNomination(t *testing.T) {
	const level = 2
	ctx := t.Context()

	app := &iface.RootSoftware{}
	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	status, statusStore, privateValues := createStatus(1, level)
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
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(kinds.Txs{})

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.VerifyingTracer(),
		gatewayApplication.Agreement(),
		mp,
		sm.EmptyProofDepository{},
		ledgerDepot,
	)
	pa, _ := status.Ratifiers.FetchByOrdinal(0)
	endorse, _, err := createSoundEndorse(level, kinds.LedgerUID{}, status.Ratifiers, privateValues)
	require.NoError(t, err)
	_, err = ledgerExecute.InstantiateNominationLedger(ctx, level, status, endorse, pa)
	require.NoError(t, err)
}

//
//
func VerifyArrangeNominationTransAllEnclosed(t *testing.T) {
	const level = 2
	ctx := t.Context()

	status, statusStore, privateValues := createStatus(1, level)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	eventpool := &simulations.ProofDepository{}
	eventpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, int64(0))

	txs := verify.CreateNTrans(level, 10)
	mp := &txpoolsims.Txpool{}
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(txs[2:])

	app := &abciemulators.Software{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{
		Txs: txs.ToSegmentOfOctets(),
	}, nil)
	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.VerifyingTracer(),
		gatewayApplication.Agreement(),
		mp,
		eventpool,
		ledgerDepot,
	)
	pa, _ := status.Ratifiers.FetchByOrdinal(0)
	endorse, _, err := createSoundEndorse(level, kinds.LedgerUID{}, status.Ratifiers, privateValues)
	require.NoError(t, err)
	ledger, err := ledgerExecute.InstantiateNominationLedger(ctx, level, status, endorse, pa)
	require.NoError(t, err)

	for i, tx := range ledger.Txs {
		require.Equal(t, txs[i], tx)
	}

	mp.AssertExpectations(t)
}

//
//
func VerifyArrangeNominationRearrangeTrans(t *testing.T) {
	const level = 2
	ctx := t.Context()

	status, statusStore, privateValues := createStatus(1, level)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	eventpool := &simulations.ProofDepository{}
	eventpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, int64(0))

	txs := verify.CreateNTrans(level, 10)
	mp := &txpoolsims.Txpool{}
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(txs)

	txs = txs[2:]
	txs = append(txs[len(txs)/2:], txs[:len(txs)/2]...)

	app := &abciemulators.Software{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{
		Txs: txs.ToSegmentOfOctets(),
	}, nil)

	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.VerifyingTracer(),
		gatewayApplication.Agreement(),
		mp,
		eventpool,
		ledgerDepot,
	)
	pa, _ := status.Ratifiers.FetchByOrdinal(0)
	endorse, _, err := createSoundEndorse(level, kinds.LedgerUID{}, status.Ratifiers, privateValues)
	require.NoError(t, err)
	ledger, err := ledgerExecute.InstantiateNominationLedger(ctx, level, status, endorse, pa)
	require.NoError(t, err)
	for i, tx := range ledger.Txs {
		require.Equal(t, txs[i], tx)
	}

	mp.AssertExpectations(t)
}

//
//
func VerifyArrangeNominationFaultOnTooNumerousTrans(t *testing.T) {
	const level = 2
	ctx := t.Context()

	status, statusStore, privateValues := createStatus(1, level)
	//
	status.AgreementOptions.Ledger.MaximumOctets = 60 * 1024
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	eventpool := &simulations.ProofDepository{}
	eventpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, int64(0))

	const nRatifiers = 1
	var octetsEachTransfer int64 = 3
	maximumDataOctets := kinds.MaximumDataOctets(status.AgreementOptions.Ledger.MaximumOctets, 0, nRatifiers)
	txs := verify.CreateNTrans(level, maximumDataOctets/octetsEachTransfer+2) //
	mp := &txpoolsims.Txpool{}
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(txs)

	app := &abciemulators.Software{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{
		Txs: txs.ToSegmentOfOctets(),
	}, nil)

	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.NewNoopTracer(),
		gatewayApplication.Agreement(),
		mp,
		eventpool,
		ledgerDepot,
	)
	pa, _ := status.Ratifiers.FetchByOrdinal(0)
	endorse, _, err := createSoundEndorse(level, kinds.LedgerUID{}, status.Ratifiers, privateValues)
	require.NoError(t, err)
	ledger, err := ledgerExecute.InstantiateNominationLedger(ctx, level, status, endorse, pa)
	require.Nil(t, ledger)
	require.ErrorContains(t, err, "REDACTED")

	mp.AssertExpectations(t)
}

//
//
//
func VerifyArrangeNominationNumberEncodingBurden(t *testing.T) {
	const level = 2
	ctx := t.Context()

	status, statusStore, privateValues := createStatus(1, level)
	//
	var octetsEachTransfer int64 = 4
	const nRatifiers = 1
	notDataVolume := 5000 - kinds.MaximumDataOctets(5000, 0, nRatifiers)
	status.AgreementOptions.Ledger.MaximumOctets = octetsEachTransfer*1024 + notDataVolume
	maximumDataOctets := kinds.MaximumDataOctets(status.AgreementOptions.Ledger.MaximumOctets, 0, nRatifiers)

	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	eventpool := &simulations.ProofDepository{}
	eventpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, int64(0))

	txs := verify.CreateNTrans(level, maximumDataOctets/octetsEachTransfer)
	mp := &txpoolsims.Txpool{}
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(txs)

	app := &abciemulators.Software{}
	app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{
		Txs: txs.ToSegmentOfOctets(),
	}, nil)

	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.NewNoopTracer(),
		gatewayApplication.Agreement(),
		mp,
		eventpool,
		ledgerDepot,
	)
	pa, _ := status.Ratifiers.FetchByOrdinal(0)
	endorse, _, err := createSoundEndorse(level, kinds.LedgerUID{}, status.Ratifiers, privateValues)
	require.NoError(t, err)
	ledger, err := ledgerExecute.InstantiateNominationLedger(ctx, level, status, endorse, pa)
	require.Nil(t, ledger)
	require.ErrorContains(t, err, "REDACTED")

	mp.AssertExpectations(t)
}

//
//
func VerifyArrangeNominationFaultOnArrangeNominationFault(t *testing.T) {
	const level = 2
	ctx := t.Context()

	status, statusStore, privateValues := createStatus(1, level)
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})

	eventpool := &simulations.ProofDepository{}
	eventpool.On("REDACTED", mock.Anything).Return([]kinds.Proof{}, int64(0))

	txs := verify.CreateNTrans(level, 10)
	mp := &txpoolsims.Txpool{}
	mp.On("REDACTED", mock.Anything, mock.Anything).Return(txs)

	cm := &ifaceclientmocks.Customer{}
	cm.On("REDACTED", mock.Anything).Return()
	cm.On("REDACTED").Return(nil)
	cm.On("REDACTED").Return(nil)
	cm.On("REDACTED", mock.Anything, mock.Anything).Return(nil, errors.New("REDACTED")).Once()
	cm.On("REDACTED").Return(nil)
	cc := &omocks.CustomerOriginator{}
	cc.On("REDACTED").Return(cm, nil)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	require.NoError(t, err)
	defer gatewayApplication.Halt() //

	ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		log.NewNoopTracer(),
		gatewayApplication.Agreement(),
		mp,
		eventpool,
		ledgerDepot,
	)
	pa, _ := status.Ratifiers.FetchByOrdinal(0)
	endorse, _, err := createSoundEndorse(level, kinds.LedgerUID{}, status.Ratifiers, privateValues)
	require.NoError(t, err)
	ledger, err := ledgerExecute.InstantiateNominationLedger(ctx, level, status, endorse, pa)
	require.Nil(t, ledger)
	require.ErrorContains(t, err, "REDACTED")

	mp.AssertExpectations(t)
}

//
//
//
func VerifyInstantiateNominationMissingBallotPlugins(t *testing.T) {
	for _, verifyInstance := range []struct {
		label string

		//
		level int64

		//
		additionActivateLevel int64
		anticipateAlarm           bool
	}{
		{
			label:                  "REDACTED",
			level:                3,
			additionActivateLevel: 2,
			anticipateAlarm:           true,
		},
		{
			label:                  "REDACTED",
			level:                3,
			additionActivateLevel: 3,
			anticipateAlarm:           false,
		},
		{
			label:                  "REDACTED",
			level:                3,
			additionActivateLevel: 0,
			anticipateAlarm:           false,
		},
		{
			label:                  "REDACTED",
			level:                3,
			additionActivateLevel: 4,
			anticipateAlarm:           false,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			ctx := t.Context()

			app := abciemulators.NewSoftware(t)
			if !verifyInstance.anticipateAlarm {
				app.On("REDACTED", mock.Anything, mock.Anything).Return(&iface.ReplyArrangeNomination{}, nil)
			}
			cc := gateway.NewNativeCustomerOriginator(app)
			gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
			err := gatewayApplication.Begin()
			require.NoError(t, err)

			status, statusStore, privateValues := createStatus(1, int(verifyInstance.level-1))
			statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
				DropIfaceReplies: false,
			})
			status.AgreementOptions.Iface.BallotPluginsActivateLevel = verifyInstance.additionActivateLevel
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
			mp.On("REDACTED", mock.Anything, mock.Anything).Return(kinds.Txs{})

			ledgerDepot := depot.NewLedgerDepot(dbm.NewMemoryStore())
			ledgerExecute := sm.NewLedgerRunner(
				statusDepot,
				log.NewNoopTracer(),
				gatewayApplication.Agreement(),
				mp,
				sm.EmptyProofDepository{},
				ledgerDepot,
			)
			ledger, err := createLedger(status, verifyInstance.level, new(kinds.Endorse))
			require.NoError(t, err)

			bps, err := ledger.CreateSegmentAssign(verifySegmentVolume)
			require.NoError(t, err)
			ledgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: bps.Heading()}
			pa, _ := status.Ratifiers.FetchByOrdinal(0)
			finalEndorse, _, _ := createSoundEndorse(verifyInstance.level-1, ledgerUID, status.Ratifiers, privateValues)
			removeEndorsements(finalEndorse)
			if verifyInstance.anticipateAlarm {
				require.Panics(t, func() {
					_, err := ledgerExecute.InstantiateNominationLedger(ctx, verifyInstance.level, status, finalEndorse, pa)
					require.NoError(t, err)
				})
			} else {
				_, err = ledgerExecute.InstantiateNominationLedger(ctx, verifyInstance.level, status, finalEndorse, pa)
				require.NoError(t, err)
			}
		})
	}
}

func removeEndorsements(ec *kinds.ExpandedEndorse) {
	for i, endorseSignature := range ec.ExpandedEndorsements {
		endorseSignature.Addition = nil
		endorseSignature.AdditionAutograph = nil
		ec.ExpandedEndorsements[i] = endorseSignature
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
