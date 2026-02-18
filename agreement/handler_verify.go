package agreement

import (
	"context"
	"fmt"
	"os"
	"path"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	abciend "github.com/valkyrieworks/iface/customer"
	"github.com/valkyrieworks/iface/instance/objectdepot"
	iface "github.com/valkyrieworks/iface/kinds"
	cfg "github.com/valkyrieworks/settings"
	cskinds "github.com/valkyrieworks/agreement/kinds"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/utils/bits"
	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"
	engineconnect "github.com/valkyrieworks/utils/align"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/p2p"
	p2pemulator "github.com/valkyrieworks/p2p/emulate"
	cometconnect "github.com/valkyrieworks/schema/consensuscore/agreement"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	statemulators "github.com/valkyrieworks/status/simulations"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	cometfaults "github.com/valkyrieworks/kinds/faults"
)

//
//

var standardVerifyTime = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

func beginAgreementNet(t *testing.T, css []*Status, n int) (
	[]*Handler,
	[]kinds.Enrollment,
	[]*kinds.EventBus,
) {
	handlers := make([]*Handler, n)
	ledgersEnrollments := make([]kinds.Enrollment, 0)
	eventBuses := make([]*kinds.EventBus, n)
	for i := 0; i < n; i++ {
		/*)
*/
		handlers[i] = NewHandler(css[i], true) //
		handlers[i].AssignTracer(css[i].Tracer)

		//
		eventBuses[i] = css[i].eventBus
		handlers[i].AssignEventBus(eventBuses[i])

		ledgersSubtract, err := eventBuses[i].Enrol(context.Background(), verifyEnrollee, kinds.EventInquireNewLedger)
		require.NoError(t, err)
		ledgersEnrollments = append(ledgersEnrollments, ledgersSubtract)

		if css[i].status.FinalLedgerLevel == 0 { //
			if err := css[i].ledgerExecute.Depot().Persist(css[i].status); err != nil {
				t.Error(err)
			}
		}
	}
	//
	p2p.CreateLinkedRouters(settings.P2P, n, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", handlers[i])
		s.AssignTracer(handlers[i].connectS.Tracer.With("REDACTED", "REDACTED"))
		return s
	}, p2p.Connect2routers)

	//
	//
	//
	//
	for i := 0; i < n; i++ {
		s := handlers[i].connectS.FetchStatus()
		handlers[i].RouterToAgreement(s, false)
	}
	return handlers, ledgersEnrollments, eventBuses
}

func haltAgreementNet(tracer log.Tracer, handlers []*Handler, eventBuses []*kinds.EventBus) {
	tracer.Details("REDACTED", "REDACTED", len(handlers))
	for i, r := range handlers {
		tracer.Details("REDACTED", "REDACTED", i)
		if err := r.Router.Halt(); err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	for i, b := range eventBuses {
		tracer.Details("REDACTED", "REDACTED", i)
		if err := b.Halt(); err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	tracer.Details("REDACTED", "REDACTED", len(handlers))
}

//
func VerifyHandlerSimple(t *testing.T) {
	N := 4
	css, sanitize := randomAgreementNet(t, N, "REDACTED", newEmulateTimerFunction(true), newObjectDepot)
	defer sanitize()
	handlers, ledgersEnrollments, eventBuses := beginAgreementNet(t, css, N)
	defer haltAgreementNet(log.VerifyingTracer(), handlers, eventBuses)
	//
	deadlineWaitCluster(N, func(j int) {
		<-ledgersEnrollments[j].Out()
	})
}

//
func VerifyHandlerWithProof(t *testing.T) {
	nRatifiers := 4
	verifyLabel := "REDACTED"
	timerFunction := newEmulateTimerFunction(true)
	applicationFunction := newObjectDepot

	//
	//
	//

	generatePaper, privateValues := randomOriginPaper(nRatifiers, false, 30, nil)
	css := make([]*Status, nRatifiers)
	tracer := agreementTracer()
	for i := 0; i < nRatifiers; i++ {
		statusStore := dbm.NewMemoryStore() //
		statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
			DropIfaceReplies: false,
		})
		status, _ := statusDepot.ImportFromStoreOrOriginPaper(generatePaper)
		thisSettings := RestoreSettings(fmt.Sprintf("REDACTED", verifyLabel, i))
		defer os.RemoveAll(thisSettings.OriginFolder)
		assureFolder(path.Dir(thisSettings.Agreement.JournalEntry()), 0o700) //
		app := applicationFunction()
		values := kinds.Tm2schema.RatifierRefreshes(status.Ratifiers)
		_, err := app.InitSeries(context.Background(), &iface.QueryInitSeries{Ratifiers: values})
		require.NoError(t, err)

		pv := privateValues[i]
		//
		//

		ledgerStore := dbm.NewMemoryStore()
		ledgerDepot := depot.NewLedgerDepot(ledgerStore)

		mtx := new(engineconnect.Lock)
		memplStats := txpool.NoopStats()
		//
		gatewayApplicationLinkConnect := gateway.NewApplicationLinkAgreement(abciend.NewNativeCustomer(mtx, app), gateway.NoopStats())
		gatewayApplicationLinkMemory := gateway.NewApplicationLinkTxpool(abciend.NewNativeCustomer(mtx, app), gateway.NoopStats())

		//
		txpool := txpool.NewCCatalogTxpool(settings.Txpool,
			gatewayApplicationLinkMemory,
			status.FinalLedgerLevel,
			txpool.WithStats(memplStats),
			txpool.WithPreInspect(sm.TransferPreInspect(status)),
			txpool.WithSubmitInspect(sm.TransferSubmitInspect(status)))

		if thisSettings.Agreement.WaitForTrans() {
			txpool.ActivateTransAccessible()
		}

		//
		//
		vIdx := (i + 1) % nRatifiers
		ev, err := kinds.NewEmulateReplicatedBallotProofWithRatifier(1, standardVerifyTime, privateValues[vIdx], generatePaper.LedgerUID)
		require.NoError(t, err)
		eventpool := &statemulators.ProofDepository{}
		eventpool.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)
		eventpool.On("REDACTED", mock.AnythingOfType("REDACTED")).Return([]kinds.Proof{
			ev,
		}, int64(len(ev.Octets())))
		eventpool.On("REDACTED", mock.AnythingOfType("REDACTED"), mock.AnythingOfType("REDACTED")).Return()

		evpool2 := sm.EmptyProofDepository{}

		//
		ledgerExecute := sm.NewLedgerRunner(statusDepot, log.VerifyingTracer(), gatewayApplicationLinkConnect, txpool, eventpool, ledgerDepot)
		cs := NewStatus(thisSettings.Agreement, status, ledgerExecute, ledgerDepot, txpool, evpool2)
		cs.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
		cs.CollectionPrivateRatifier(pv)

		eventBus := kinds.NewEventBus()
		eventBus.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
		err = eventBus.Begin()
		require.NoError(t, err)
		cs.AssignEventBus(eventBus)

		cs.CollectionDeadlineTimer(timerFunction())
		cs.AssignTracer(tracer.With("REDACTED", i, "REDACTED", "REDACTED"))

		css[i] = cs
	}

	handlers, ledgersEnrollments, eventBuses := beginAgreementNet(t, css, nRatifiers)
	defer haltAgreementNet(log.VerifyingTracer(), handlers, eventBuses)

	//
	for i := 0; i < nRatifiers; i++ {
		deadlineWaitCluster(nRatifiers, func(j int) {
			msg := <-ledgersEnrollments[j].Out()
			ledger := msg.Data().(kinds.EventDataNewLedger).Ledger
			assert.Len(t, ledger.Proof.Proof, 1)
		})
	}
}

//

//
func VerifyHandlerGeneratesLedgerWhenEmptyLedgersFalse(t *testing.T) {
	N := 4
	css, sanitize := randomAgreementNet(t, N, "REDACTED", newEmulateTimerFunction(true), newObjectDepot,
		func(c *cfg.Settings) {
			c.Agreement.GenerateEmptyLedgers = false
		})
	defer sanitize()
	handlers, ledgersEnrollments, eventBuses := beginAgreementNet(t, css, N)
	defer haltAgreementNet(log.VerifyingTracer(), handlers, eventBuses)

	//
	if err := affirmTxpool(css[3].transferAlerter).InspectTransfer(objectdepot.NewTransferFromUID(1), func(reply *iface.ReplyInspectTransfer) {
		require.False(t, reply.IsErr())
	}, txpool.TransferDetails{}); err != nil {
		t.Error(err)
	}

	//
	deadlineWaitCluster(N, func(j int) {
		<-ledgersEnrollments[j].Out()
	})
}

func VerifyHandlerAcceptDoesNotAlarmIfAppendNodeHasntExistedInvokedYet(t *testing.T) {
	N := 1
	css, sanitize := randomAgreementNet(t, N, "REDACTED", newEmulateTimerFunction(true), newObjectDepot)
	defer sanitize()
	handlers, _, eventBuses := beginAgreementNet(t, css, N)
	defer haltAgreementNet(log.VerifyingTracer(), handlers, eventBuses)

	var (
		handler = handlers[0]
		node    = p2pemulator.NewNode(nil)
	)

	handler.InitNode(node)

	//
	assert.NotPanics(t, func() {
		handler.Accept(p2p.Packet{
			StreamUID: StatusStream,
			Src:       node,
			Signal: &cometconnect.HasBallot{
				Level: 1,
				Cycle:  1,
				Ordinal:  1,
				Kind:   engineproto.PreballotKind,
			},
		})
		handler.AppendNode(node)
	})
}

func VerifyHandlerAcceptAlarmsIfInitNodeHasntExistedInvokedYet(t *testing.T) {
	N := 1
	css, sanitize := randomAgreementNet(t, N, "REDACTED", newEmulateTimerFunction(true), newObjectDepot)
	defer sanitize()
	handlers, _, eventBuses := beginAgreementNet(t, css, N)
	defer haltAgreementNet(log.VerifyingTracer(), handlers, eventBuses)

	var (
		handler = handlers[0]
		node    = p2pemulator.NewNode(nil)
	)

	//

	//
	assert.Panics(t, func() {
		handler.Accept(p2p.Packet{
			StreamUID: StatusStream,
			Src:       node,
			Signal: &cometconnect.HasBallot{
				Level: 1,
				Cycle:  1,
				Ordinal:  1,
				Kind:   engineproto.PreballotKind,
			},
		})
	})
}

//
//
func VerifyRouterToAgreementBallotPlugins(t *testing.T) {
	for _, verifyInstance := range []struct {
		label                  string
		archivedLevel          int64
		primaryMandatoryLevel int64
		encompassPlugins     bool
		mustAlarm           bool
	}{
		{
			label:                  "REDACTED",
			primaryMandatoryLevel: 0,
			archivedLevel:          2,
			encompassPlugins:     false,
			mustAlarm:           false,
		},
		{
			label:                  "REDACTED",
			primaryMandatoryLevel: 2,
			archivedLevel:          2,
			encompassPlugins:     false,
			mustAlarm:           true,
		},
		{
			label:                  "REDACTED",
			primaryMandatoryLevel: 3,
			archivedLevel:          2,
			encompassPlugins:     false,
			mustAlarm:           false,
		},
		{
			label:                  "REDACTED",
			primaryMandatoryLevel: 1,
			archivedLevel:          2,
			encompassPlugins:     false,
			mustAlarm:           true,
		},
		{
			label:                  "REDACTED",
			primaryMandatoryLevel: 1,
			archivedLevel:          2,
			encompassPlugins:     true,
			mustAlarm:           false,
		},
	} {
		t.Run(verifyInstance.label, func(t *testing.T) {
			ctx := t.Context()

			cs, vs := randomStatus(1)
			ratifier := vs[0]
			ratifier.Level = verifyInstance.archivedLevel

			cs.status.FinalLedgerLevel = verifyInstance.archivedLevel
			cs.status.FinalRatifiers = cs.status.Ratifiers.Clone()
			cs.status.AgreementOptions.Iface.BallotPluginsActivateLevel = verifyInstance.primaryMandatoryLevel

			nominationLedger, err := cs.instantiateNominationLedger(ctx)
			require.NoError(t, err)

			//
			cs.Level = verifyInstance.archivedLevel + 1
			nominationLedger.Level = verifyInstance.archivedLevel
			ledgerSegments, err := nominationLedger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
			require.NoError(t, err)

			var ballotCollection *kinds.BallotCollection
			if verifyInstance.encompassPlugins {
				ballotCollection = kinds.NewExpandedBallotCollection(cs.status.LedgerUID, verifyInstance.archivedLevel, 0, engineproto.PreendorseKind, cs.status.Ratifiers)
			} else {
				ballotCollection = kinds.NewBallotCollection(cs.status.LedgerUID, verifyInstance.archivedLevel, 0, engineproto.PreendorseKind, cs.status.Ratifiers)
			}
			attestedBallot := attestBallot(ratifier, engineproto.PreendorseKind, nominationLedger.Digest(), ledgerSegments.Heading(), verifyInstance.encompassPlugins)

			var veLevel int64
			if verifyInstance.encompassPlugins {
				require.NotNil(t, attestedBallot.AdditionAutograph)
				veLevel = verifyInstance.archivedLevel
			} else {
				require.Nil(t, attestedBallot.Addition)
				require.Nil(t, attestedBallot.AdditionAutograph)
			}

			appended, err := ballotCollection.AppendBallot(attestedBallot)
			require.NoError(t, err)
			require.True(t, appended)

			veLevelArgument := kinds.IfaceOptions{BallotPluginsActivateLevel: veLevel}
			if verifyInstance.encompassPlugins {
				cs.ledgerDepot.PersistLedgerWithExpandedEndorse(nominationLedger, ledgerSegments, ballotCollection.CreateExpandedEndorse(veLevelArgument))
			} else {
				cs.ledgerDepot.PersistLedger(nominationLedger, ledgerSegments, ballotCollection.CreateExpandedEndorse(veLevelArgument).ToEndorse())
			}
			handler := NewHandler(
				cs,
				true,
			)

			if verifyInstance.mustAlarm {
				assert.Panics(t, func() {
					handler.RouterToAgreement(cs.status, false)
				})
			} else {
				handler.RouterToAgreement(cs.status, false)
			}
		})
	}
}

//
func VerifyHandlerEntriesBallotsAndLedgerSegments(t *testing.T) {
	N := 4
	css, sanitize := randomAgreementNet(t, N, "REDACTED", newEmulateTimerFunction(true), newObjectDepot)
	defer sanitize()
	handlers, ledgersEnrollments, eventBuses := beginAgreementNet(t, css, N)
	defer haltAgreementNet(log.VerifyingTracer(), handlers, eventBuses)

	//
	deadlineWaitCluster(N, func(j int) {
		<-ledgersEnrollments[j].Out()
	})

	//
	node := handlers[1].Router.Nodes().Clone()[0]
	//
	ps := node.Get(kinds.NodeStatusKey).(*NodeStatus)

	assert.Equal(t, true, ps.BallotsRelayed() > 0, "REDACTED")
	assert.Equal(t, true, ps.LedgerSegmentsRelayed() > 0, "REDACTED")
}

//
//

func VerifyHandlerPollingEnergyAlter(t *testing.T) {
	nValues := 4
	tracer := log.VerifyingTracer()
	css, sanitize := randomAgreementNet(
		t,
		nValues,
		"REDACTED",
		newEmulateTimerFunction(true),
		newDurableObjectDepot)
	defer sanitize()
	handlers, ledgersEnrollments, eventBuses := beginAgreementNet(t, css, nValues)
	defer haltAgreementNet(tracer, handlers, eventBuses)

	//
	enabledValues := make(map[string]struct{})
	for i := 0; i < nValues; i++ {
		publicKey, err := css[i].privateRatifier.FetchPublicKey()
		require.NoError(t, err)
		address := publicKey.Location()
		enabledValues[string(address)] = struct{}{}
	}

	//
	deadlineWaitCluster(nValues, func(j int) {
		<-ledgersEnrollments[j].Out()
	})

	//
	tracer.Diagnose("REDACTED")

	val1publicKey, err := css[0].privateRatifier.FetchPublicKey()
	require.NoError(t, err)

	val1publicKeyIface, err := cryptocode.PublicKeyToSchema(val1publicKey)
	require.NoError(t, err)
	modifyRatifierTransfer := objectdepot.CreateValueCollectionAlterTransfer(val1publicKeyIface, 25)
	precedingSumPollingEnergy := css[0].FetchDurationStatus().FinalRatifiers.SumPollingEnergy()

	waitForAndCertifyLedger(t, nValues, enabledValues, ledgersEnrollments, css, modifyRatifierTransfer)
	waitForAndCertifyLedgerWithTransfer(t, nValues, enabledValues, ledgersEnrollments, css, modifyRatifierTransfer)
	waitForAndCertifyLedger(t, nValues, enabledValues, ledgersEnrollments, css)
	waitForAndCertifyLedger(t, nValues, enabledValues, ledgersEnrollments, css)

	if css[0].FetchDurationStatus().FinalRatifiers.SumPollingEnergy() == precedingSumPollingEnergy {
		t.Fatalf(
			"REDACTED",
			precedingSumPollingEnergy,
			css[0].FetchDurationStatus().FinalRatifiers.SumPollingEnergy())
	}

	modifyRatifierTransfer = objectdepot.CreateValueCollectionAlterTransfer(val1publicKeyIface, 2)
	precedingSumPollingEnergy = css[0].FetchDurationStatus().FinalRatifiers.SumPollingEnergy()

	waitForAndCertifyLedger(t, nValues, enabledValues, ledgersEnrollments, css, modifyRatifierTransfer)
	waitForAndCertifyLedgerWithTransfer(t, nValues, enabledValues, ledgersEnrollments, css, modifyRatifierTransfer)
	waitForAndCertifyLedger(t, nValues, enabledValues, ledgersEnrollments, css)
	waitForAndCertifyLedger(t, nValues, enabledValues, ledgersEnrollments, css)

	if css[0].FetchDurationStatus().FinalRatifiers.SumPollingEnergy() == precedingSumPollingEnergy {
		t.Fatalf(
			"REDACTED",
			precedingSumPollingEnergy,
			css[0].FetchDurationStatus().FinalRatifiers.SumPollingEnergy())
	}

	modifyRatifierTransfer = objectdepot.CreateValueCollectionAlterTransfer(val1publicKeyIface, 26)
	precedingSumPollingEnergy = css[0].FetchDurationStatus().FinalRatifiers.SumPollingEnergy()

	waitForAndCertifyLedger(t, nValues, enabledValues, ledgersEnrollments, css, modifyRatifierTransfer)
	waitForAndCertifyLedgerWithTransfer(t, nValues, enabledValues, ledgersEnrollments, css, modifyRatifierTransfer)
	waitForAndCertifyLedger(t, nValues, enabledValues, ledgersEnrollments, css)
	waitForAndCertifyLedger(t, nValues, enabledValues, ledgersEnrollments, css)

	if css[0].FetchDurationStatus().FinalRatifiers.SumPollingEnergy() == precedingSumPollingEnergy {
		t.Fatalf(
			"REDACTED",
			precedingSumPollingEnergy,
			css[0].FetchDurationStatus().FinalRatifiers.SumPollingEnergy())
	}
}

func VerifyHandlerRatifierCollectionModifications(t *testing.T) {
	nNodes := 7
	nValues := 4
	css, _, _, sanitize := randomAgreementNetWithNodes(
		t,
		nValues,
		nNodes,
		"REDACTED",
		newEmulateTimerFunction(true),
		newDurableObjectDepotWithRoute)

	defer sanitize()
	tracer := log.VerifyingTracer()

	handlers, ledgersEnrollments, eventBuses := beginAgreementNet(t, css, nNodes)
	defer haltAgreementNet(tracer, handlers, eventBuses)

	//
	enabledValues := make(map[string]struct{})
	for i := 0; i < nValues; i++ {
		publicKey, err := css[i].privateRatifier.FetchPublicKey()
		require.NoError(t, err)
		enabledValues[string(publicKey.Location())] = struct{}{}
	}

	//
	deadlineWaitCluster(nNodes, func(j int) {
		<-ledgersEnrollments[j].Out()
	})

	t.Run("REDACTED", func(t *testing.T) {
		newRatifierPublicKey1, err := css[nValues].privateRatifier.FetchPublicKey()
		assert.NoError(t, err)
		valuePublicKey1abci, err := cryptocode.PublicKeyToSchema(newRatifierPublicKey1)
		assert.NoError(t, err)
		newRatifierTrans1 := objectdepot.CreateValueCollectionAlterTransfer(valuePublicKey1abci, verifyMinimumEnergy)

		//
		//
		//
		waitForAndCertifyLedger(t, nNodes, enabledValues, ledgersEnrollments, css, newRatifierTrans1)

		//
		//
		waitForAndCertifyLedgerWithTransfer(t, nNodes, enabledValues, ledgersEnrollments, css, newRatifierTrans1)

		//
		//
		waitForAndCertifyLedger(t, nNodes, enabledValues, ledgersEnrollments, css)

		//
		enabledValues[string(newRatifierPublicKey1.Location())] = struct{}{}

		//
		//
		waitForLedgerWithRefreshedValuesAndCertifyIt(t, nNodes, enabledValues, ledgersEnrollments, css)
	})

	t.Run("REDACTED", func(t *testing.T) {
		modifyRatifierPublicKey1, err := css[nValues].privateRatifier.FetchPublicKey()
		require.NoError(t, err)
		modifyPublicKey1abci, err := cryptocode.PublicKeyToSchema(modifyRatifierPublicKey1)
		require.NoError(t, err)
		modifyRatifierTrans1 := objectdepot.CreateValueCollectionAlterTransfer(modifyPublicKey1abci, 25)
		precedingSumPollingEnergy := css[nValues].FetchDurationStatus().FinalRatifiers.SumPollingEnergy()

		waitForAndCertifyLedger(t, nNodes, enabledValues, ledgersEnrollments, css, modifyRatifierTrans1)
		waitForAndCertifyLedgerWithTransfer(t, nNodes, enabledValues, ledgersEnrollments, css, modifyRatifierTrans1)
		waitForAndCertifyLedger(t, nNodes, enabledValues, ledgersEnrollments, css)
		waitForLedgerWithRefreshedValuesAndCertifyIt(t, nNodes, enabledValues, ledgersEnrollments, css)

		if css[nValues].FetchDurationStatus().FinalRatifiers.SumPollingEnergy() == precedingSumPollingEnergy {
			t.Errorf(
				"REDACTED",
				precedingSumPollingEnergy,
				css[nValues].FetchDurationStatus().FinalRatifiers.SumPollingEnergy())
		}
	})

	newRatifierPublicKey2, err := css[nValues+1].privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	newVal2abci, err := cryptocode.PublicKeyToSchema(newRatifierPublicKey2)
	require.NoError(t, err)
	newRatifierTrans2 := objectdepot.CreateValueCollectionAlterTransfer(newVal2abci, verifyMinimumEnergy)

	newRatifierPublicKey3, err := css[nValues+2].privateRatifier.FetchPublicKey()
	require.NoError(t, err)
	newVal3abci, err := cryptocode.PublicKeyToSchema(newRatifierPublicKey3)
	require.NoError(t, err)
	newRatifierTrans3 := objectdepot.CreateValueCollectionAlterTransfer(newVal3abci, verifyMinimumEnergy)

	t.Run("REDACTED", func(t *testing.T) {
		waitForAndCertifyLedger(t, nNodes, enabledValues, ledgersEnrollments, css, newRatifierTrans2, newRatifierTrans3)
		waitForAndCertifyLedgerWithTransfer(t, nNodes, enabledValues, ledgersEnrollments, css, newRatifierTrans2, newRatifierTrans3)
		waitForAndCertifyLedger(t, nNodes, enabledValues, ledgersEnrollments, css)
		enabledValues[string(newRatifierPublicKey2.Location())] = struct{}{}
		enabledValues[string(newRatifierPublicKey3.Location())] = struct{}{}
		waitForLedgerWithRefreshedValuesAndCertifyIt(t, nNodes, enabledValues, ledgersEnrollments, css)
	})

	t.Run("REDACTED", func(t *testing.T) {
		deleteRatifierTrans2 := objectdepot.CreateValueCollectionAlterTransfer(newVal2abci, 0)
		deleteRatifierTrans3 := objectdepot.CreateValueCollectionAlterTransfer(newVal3abci, 0)

		waitForAndCertifyLedger(t, nNodes, enabledValues, ledgersEnrollments, css, deleteRatifierTrans2, deleteRatifierTrans3)
		waitForAndCertifyLedgerWithTransfer(t, nNodes, enabledValues, ledgersEnrollments, css, deleteRatifierTrans2, deleteRatifierTrans3)
		waitForAndCertifyLedger(t, nNodes, enabledValues, ledgersEnrollments, css)
		delete(enabledValues, string(newRatifierPublicKey2.Location()))
		delete(enabledValues, string(newRatifierPublicKey3.Location()))
		waitForLedgerWithRefreshedValuesAndCertifyIt(t, nNodes, enabledValues, ledgersEnrollments, css)
	})
}

//
func VerifyHandlerWithDeadlineEndorse(t *testing.T) {
	N := 4
	css, sanitize := randomAgreementNet(t, N, "REDACTED", newEmulateTimerFunction(false), newObjectDepot)
	defer sanitize()
	//
	for i := 0; i < N; i++ {
		css[i].settings.OmitDeadlineEndorse = false
	}

	handlers, ledgersEnrollments, eventBuses := beginAgreementNet(t, css, N-1)
	defer haltAgreementNet(log.VerifyingTracer(), handlers, eventBuses)

	//
	deadlineWaitCluster(N-1, func(j int) {
		<-ledgersEnrollments[j].Out()
	})
}

func waitForAndCertifyLedger(
	t *testing.T,
	n int,
	enabledValues map[string]struct{},
	ledgersEnrollments []kinds.Enrollment,
	css []*Status,
	txs ...[]byte,
) {
	deadlineWaitCluster(n, func(j int) {
		css[j].Tracer.Diagnose("REDACTED")
		msg := <-ledgersEnrollments[j].Out()
		newLedger := msg.Data().(kinds.EventDataNewLedger).Ledger
		css[j].Tracer.Diagnose("REDACTED", "REDACTED", newLedger.Level)
		err := certifyLedger(newLedger, enabledValues)
		require.NoError(t, err)

		//
		for _, tx := range txs {
			err := affirmTxpool(css[j].transferAlerter).InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) {
				require.False(t, reply.IsErr())
				fmt.Println(reply)
			}, txpool.TransferDetails{})
			require.NoError(t, err)
		}
	})
}

func waitForAndCertifyLedgerWithTransfer(
	t *testing.T,
	n int,
	enabledValues map[string]struct{},
	ledgersEnrollments []kinds.Enrollment,
	css []*Status,
	txs ...[]byte,
) {
	deadlineWaitCluster(n, func(j int) {
		ntrans := 0
	LEDGER_TRANSFER_CYCLE:
		for {
			css[j].Tracer.Diagnose("REDACTED", "REDACTED", ntrans)
			msg := <-ledgersEnrollments[j].Out()
			newLedger := msg.Data().(kinds.EventDataNewLedger).Ledger
			css[j].Tracer.Diagnose("REDACTED", "REDACTED", newLedger.Level)
			err := certifyLedger(newLedger, enabledValues)
			require.NoError(t, err)

			//
			//
			//
			for _, tx := range newLedger.Txs {
				assert.EqualValues(t, txs[ntrans], tx)
				ntrans++
			}

			if ntrans == len(txs) {
				break LEDGER_TRANSFER_CYCLE
			}
		}
	})
}

func waitForLedgerWithRefreshedValuesAndCertifyIt(
	t *testing.T,
	n int,
	refreshedValues map[string]struct{},
	ledgersEnrollments []kinds.Enrollment,
	css []*Status,
) {
	deadlineWaitCluster(n, func(j int) {
		var newLedger *kinds.Ledger
	Cycle:
		for {
			css[j].Tracer.Diagnose("REDACTED")
			msg := <-ledgersEnrollments[j].Out()
			newLedger = msg.Data().(kinds.EventDataNewLedger).Ledger
			if newLedger.FinalEndorse.Volume() == len(refreshedValues) {
				css[j].Tracer.Diagnose("REDACTED", "REDACTED", newLedger.Level)
				break Cycle
			}
			css[j].Tracer.Diagnose(
				"REDACTED",
				"REDACTED", newLedger.Level, "REDACTED", newLedger.FinalEndorse.Volume(), "REDACTED", len(refreshedValues),
			)
		}

		err := certifyLedger(newLedger, refreshedValues)
		assert.Nil(t, err)
	})
}

//
func certifyLedger(ledger *kinds.Ledger, enabledValues map[string]struct{}) error {
	if ledger.FinalEndorse.Volume() != len(enabledValues) {
		return fmt.Errorf(
			"REDACTED",
			ledger.FinalEndorse.Volume(),
			len(enabledValues))
	}

	for _, endorseSignature := range ledger.FinalEndorse.Endorsements {
		if _, ok := enabledValues[string(endorseSignature.RatifierLocation)]; !ok {
			return fmt.Errorf("REDACTED", endorseSignature.RatifierLocation)
		}
	}
	return nil
}

func deadlineWaitCluster(n int, f func(int)) {
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(j int) {
			f(j)
			wg.Done()
		}(i)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	//
	//
	deadline := time.Second * 20

	select {
	case <-done:
	case <-time.After(deadline):
		panic("REDACTED")
	}
}

//
//

func VerifyNewEpochPhaseSignalCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		anticipateErr              bool
		signalEpoch           int32
		signalFinalEndorseEpoch int32
		signalLevel          int64
		verifyLabel               string
		signalPhase            cskinds.DurationPhaseKind
	}{
		{false, 0, 0, 0, "REDACTED", cskinds.DurationPhaseNewLevel},
		{true, -1, 0, 0, "REDACTED", cskinds.DurationPhaseNewLevel},
		{true, 0, 0, -1, "REDACTED", cskinds.DurationPhaseNewLevel},
		{true, 0, 0, 0, "REDACTED", cskinds.DurationPhaseEndorse + 1},
		//
		{false, 0, 0, 1, "REDACTED", cskinds.DurationPhaseNewLevel},
		{false, 0, -1, 2, "REDACTED", cskinds.DurationPhaseNewLevel},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			signal := NewDurationPhaseSignal{
				Level:          tc.signalLevel,
				Cycle:           tc.signalEpoch,
				Phase:            tc.signalPhase,
				FinalEndorseDuration: tc.signalFinalEndorseEpoch,
			}

			err := signal.CertifySimple()
			if tc.anticipateErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func VerifyNewEpochPhaseSignalCertifyLevel(t *testing.T) {
	primaryLevel := int64(10)
	verifyScenarios := []struct {
		anticipateErr              bool
		signalFinalEndorseEpoch int32
		signalLevel          int64
		verifyLabel               string
	}{
		{false, 0, 11, "REDACTED"},
		{true, 0, -1, "REDACTED"},
		{true, 0, 0, "REDACTED"},
		{true, 0, 10, "REDACTED"},
		{true, -1, 11, "REDACTED"},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			signal := NewDurationPhaseSignal{
				Level:          tc.signalLevel,
				Cycle:           0,
				Phase:            cskinds.DurationPhaseNewLevel,
				FinalEndorseDuration: tc.signalFinalEndorseEpoch,
			}

			err := signal.CertifyLevel(primaryLevel)
			if tc.anticipateErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func VerifyNewSoundLedgerSignalCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		distortFn func(*NewSoundLedgerSignal)
		expirationErr     string
	}{
		{func(msg *NewSoundLedgerSignal) {}, "REDACTED"},
		{func(msg *NewSoundLedgerSignal) { msg.Level = -1 }, cometfaults.ErrAdverseField{Field: "REDACTED"}.Fault()},
		{func(msg *NewSoundLedgerSignal) { msg.Cycle = -1 }, cometfaults.ErrAdverseField{Field: "REDACTED"}.Fault()},
		{
			func(msg *NewSoundLedgerSignal) { msg.LedgerSegmentAssignHeading.Sum = 2 },
			"REDACTED",
		},
		{
			func(msg *NewSoundLedgerSignal) {
				msg.LedgerSegmentAssignHeading.Sum = 0
				msg.LedgerSegments = bits.NewBitList(0)
			},
			cometfaults.ErrMandatoryField{Field: "REDACTED"}.Fault(),
		},
		{
			func(msg *NewSoundLedgerSignal) { msg.LedgerSegments = bits.NewBitList(int(kinds.MaximumLedgerSegmentsTally) + 1) },
			"REDACTED",
		},
		{
			func(msg *NewSoundLedgerSignal) { msg.LedgerSegments.Elements = nil },
			"REDACTED",
		},
		{
			func(msg *NewSoundLedgerSignal) { msg.LedgerSegments.Bits = 500 },
			"REDACTED",
		},
	}

	for i, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			msg := &NewSoundLedgerSignal{
				Level: 1,
				Cycle:  0,
				LedgerSegmentAssignHeading: kinds.SegmentAssignHeading{
					Sum: 1,
				},
				LedgerSegments: bits.NewBitList(1),
			}

			tc.distortFn(msg)
			err := msg.CertifySimple()
			if tc.expirationErr != "REDACTED" && assert.Error(t, err) {
				assert.Contains(t, err.Error(), tc.expirationErr)
			}
		})
	}
}

func VerifyNominationPOLSignalCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		distortFn func(*NominationPOLSignal)
		expirationErr     string
	}{
		{func(msg *NominationPOLSignal) {}, "REDACTED"},
		{func(msg *NominationPOLSignal) { msg.Level = -1 }, cometfaults.ErrAdverseField{Field: "REDACTED"}.Fault()},
		{func(msg *NominationPOLSignal) { msg.NominationPOLDuration = -1 }, cometfaults.ErrAdverseField{Field: "REDACTED"}.Fault()},
		{func(msg *NominationPOLSignal) { msg.NominationPOL = bits.NewBitList(0) }, cometfaults.ErrMandatoryField{Field: "REDACTED"}.Fault()},
		{
			func(msg *NominationPOLSignal) { msg.NominationPOL = bits.NewBitList(kinds.MaximumBallotsTally + 1) },
			"REDACTED",
		},
		{
			func(msg *NominationPOLSignal) { msg.NominationPOL.Elements = nil },
			"REDACTED",
		},
		{
			func(msg *NominationPOLSignal) { msg.NominationPOL.Bits = 500 },
			"REDACTED",
		},
	}

	for i, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			msg := &NominationPOLSignal{
				Level:           1,
				NominationPOLDuration: 1,
				NominationPOL:      bits.NewBitList(1),
			}

			tc.distortFn(msg)
			err := msg.CertifySimple()
			if tc.expirationErr != "REDACTED" && assert.Error(t, err) {
				assert.Contains(t, err.Error(), tc.expirationErr)
			}
		})
	}
}

func VerifyLedgerSegmentSignalCertifySimple(t *testing.T) {
	verifySegment := new(kinds.Segment)
	verifySegment.Attestation.NodeDigest = comethash.Sum([]byte("REDACTED"))
	verifyScenarios := []struct {
		verifyLabel      string
		signalLevel int64
		signalEpoch  int32
		signalSegment   *kinds.Segment
		anticipateErr     bool
	}{
		{"REDACTED", 0, 0, verifySegment, false},
		{"REDACTED", -1, 0, verifySegment, true},
		{"REDACTED", 0, -1, verifySegment, true},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			signal := LedgerSegmentSignal{
				Level: tc.signalLevel,
				Cycle:  tc.signalEpoch,
				Segment:   tc.signalSegment,
			}

			assert.Equal(t, tc.anticipateErr, signal.CertifySimple() != nil, "REDACTED")
		})
	}

	signal := LedgerSegmentSignal{Level: 0, Cycle: 0, Segment: new(kinds.Segment)}
	signal.Segment.Ordinal = 1

	assert.Equal(t, true, signal.CertifySimple() != nil, "REDACTED")
}

func VerifyHasBallotSignalCertifySimple(t *testing.T) {
	const (
		soundAttestedMessageKind   engineproto.AttestedMessageKind = 0x01
		corruptAttestedMessageKind engineproto.AttestedMessageKind = 0x03
	)

	verifyScenarios := []struct {
		anticipateErr     bool
		signalEpoch  int32
		signalOrdinal  int32
		signalLevel int64
		verifyLabel      string
		signalKind   engineproto.AttestedMessageKind
	}{
		{false, 0, 0, 0, "REDACTED", soundAttestedMessageKind},
		{true, -1, 0, 0, "REDACTED", soundAttestedMessageKind},
		{true, 0, -1, 0, "REDACTED", soundAttestedMessageKind},
		{true, 0, 0, 0, "REDACTED", corruptAttestedMessageKind},
		{true, 0, 0, -1, "REDACTED", soundAttestedMessageKind},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			signal := HasBallotSignal{
				Level: tc.signalLevel,
				Cycle:  tc.signalEpoch,
				Kind:   tc.signalKind,
				Ordinal:  tc.signalOrdinal,
			}

			assert.Equal(t, tc.anticipateErr, signal.CertifySimple() != nil, "REDACTED")
		})
	}
}

func VerifyBallotCollectionMaj23signalCertifySimple(t *testing.T) {
	const (
		soundAttestedMessageKind   engineproto.AttestedMessageKind = 0x01
		corruptAttestedMessageKind engineproto.AttestedMessageKind = 0x03
	)

	soundLedgerUID := kinds.LedgerUID{}
	corruptLedgerUID := kinds.LedgerUID{
		Digest: octets.HexOctets{},
		SegmentAssignHeading: kinds.SegmentAssignHeading{
			Sum: 1,
			Digest:  []byte{0},
		},
	}

	verifyScenarios := []struct {
		anticipateErr      bool
		signalEpoch   int32
		signalLevel  int64
		verifyLabel       string
		signalKind    engineproto.AttestedMessageKind
		signalLedgerUID kinds.LedgerUID
	}{
		{false, 0, 0, "REDACTED", soundAttestedMessageKind, soundLedgerUID},
		{true, -1, 0, "REDACTED", soundAttestedMessageKind, soundLedgerUID},
		{true, 0, -1, "REDACTED", soundAttestedMessageKind, soundLedgerUID},
		{true, 0, 0, "REDACTED", corruptAttestedMessageKind, soundLedgerUID},
		{true, 0, 0, "REDACTED", soundAttestedMessageKind, corruptLedgerUID},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			signal := BallotAssignMaj23signal{
				Level:  tc.signalLevel,
				Cycle:   tc.signalEpoch,
				Kind:    tc.signalKind,
				LedgerUID: tc.signalLedgerUID,
			}

			assert.Equal(t, tc.anticipateErr, signal.CertifySimple() != nil, "REDACTED")
		})
	}
}

func VerifyBallotCollectionBitsSignalCertifySimple(t *testing.T) {
	verifyScenarios := []struct {
		distortFn func(*BallotAssignBitsSignal)
		expirationErr     string
	}{
		{func(msg *BallotAssignBitsSignal) {}, "REDACTED"},
		{func(msg *BallotAssignBitsSignal) { msg.Level = -1 }, cometfaults.ErrAdverseField{Field: "REDACTED"}.Fault()},
		{func(msg *BallotAssignBitsSignal) { msg.Kind = 0x03 }, cometfaults.ErrCorruptField{Field: "REDACTED"}.Fault()},
		{func(msg *BallotAssignBitsSignal) {
			msg.LedgerUID = kinds.LedgerUID{
				Digest: octets.HexOctets{},
				SegmentAssignHeading: kinds.SegmentAssignHeading{
					Sum: 1,
					Digest:  []byte{0},
				},
			}
		}, "REDACTED"},
		{
			func(msg *BallotAssignBitsSignal) { msg.Ballots = bits.NewBitList(kinds.MaximumBallotsTally + 1) },
			"REDACTED",
		},
		{
			func(msg *BallotAssignBitsSignal) { msg.Ballots.Elements = nil },
			"REDACTED",
		},
		{
			func(msg *BallotAssignBitsSignal) { msg.Ballots.Bits = 500 },
			"REDACTED",
		},
	}

	for i, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			msg := &BallotAssignBitsSignal{
				Level:  1,
				Cycle:   0,
				Kind:    0x01,
				Ballots:   bits.NewBitList(1),
				LedgerUID: kinds.LedgerUID{},
			}

			tc.distortFn(msg)
			err := msg.CertifySimple()
			if tc.expirationErr != "REDACTED" && assert.Error(t, err) {
				assert.Contains(t, err.Error(), tc.expirationErr)
			}
		})
	}
}

func VerifySerializeJSONNodeStatus(t *testing.T) {
	ps := NewNodeStatus(nil)
	data, err := json.Serialize(ps)
	require.NoError(t, err)
	require.JSONEq(t, `REDACTED{
REDACTED{
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED:
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTED,
REDACTEDl
REDACTED,
REDACTED{
REDACTED,
REDACTED}
REDACTED`, string(data))
}

func VerifyBallotSignalCertifySimple(t *testing.T) {
	_, vss := randomStatus(2)

	randomOctets := engineseed.Octets(comethash.Volume)
	ledgerUID := kinds.LedgerUID{
		Digest: randomOctets,
		SegmentAssignHeading: kinds.SegmentAssignHeading{
			Sum: 1,
			Digest:  randomOctets,
		},
	}
	ballot := attestBallot(vss[1], engineproto.PreendorseKind, randomOctets, ledgerUID.SegmentAssignHeading, true)

	verifyScenarios := []struct {
		distortFn func(*BallotSignal)
		expirationErr     string
	}{
		{func(_ *BallotSignal) {}, "REDACTED"},
		{func(msg *BallotSignal) { msg.Ballot.RatifierOrdinal = -1 }, "REDACTED"},
		//
		{func(msg *BallotSignal) { msg.Ballot.RatifierOrdinal = 1000 }, "REDACTED"},
	}

	for i, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			msg := &BallotSignal{ballot}

			tc.distortFn(msg)
			err := msg.CertifySimple()
			if tc.expirationErr != "REDACTED" && assert.Error(t, err) { //
				assert.Contains(t, err.Error(), tc.expirationErr)
			}
		})
	}
}

//
//
//
func VerifyHandlerAgreementOptionsModify(t *testing.T) {
	N := 4

	//
	modifyAtLevel := int64(5)
	newMaximumOctets := int64(5000000)
	newMaximumFuel := int64(50000000)

	//
	applicationFunction := func() iface.Software {
		objectApplication := objectdepot.NewInRamSoftware()
		return &agreementOptionsRefreshingApplication{
			Software:    objectApplication,
			modifyAtLevel: modifyAtLevel,
			newMaximumOctets:    newMaximumOctets,
			newMaximumFuel:      newMaximumFuel,
			t:              t,
		}
	}

	css, sanitize := randomAgreementNet(
		t,
		N,
		"REDACTED",
		newEmulateTimerFunction(true),
		applicationFunction,
	)
	defer sanitize()

	handlers, ledgersEnrollments, eventBuses := beginAgreementNet(t, css, N)
	defer haltAgreementNet(log.VerifyingTracer(), handlers, eventBuses)

	//
	primaryOptions := handlers[0].agreementOptions.Load()
	t.Logf("REDACTED", primaryOptions.Ledger.MaximumOctets, primaryOptions.Ledger.MaximumFuel)

	//
	deadlineWaitCluster(N, func(j int) {
		<-ledgersEnrollments[j].Out()
	})

	//
	//
	time.Sleep(10 * time.Millisecond)

	//
	optionsAfterLedger1 := handlers[0].agreementOptions.Load()
	assert.Equal(t, *primaryOptions, *optionsAfterLedger1, "REDACTED")

	//
	deadlineWaitCluster(N, func(j int) {
		for {
			msg := <-ledgersEnrollments[j].Out()
			event := msg.Data().(kinds.EventDataNewLedger)
			if event.Ledger.Level == modifyAtLevel {
				break
			}
		}
	})

	//
	//
	time.Sleep(10 * time.Millisecond)

	//
	for i := range N {
		handlerOptions := handlers[i].agreementOptions.Load()

		assert.Equal(t, newMaximumOctets, handlerOptions.Ledger.MaximumOctets, "REDACTED", i)
		assert.Equal(t, newMaximumFuel, handlerOptions.Ledger.MaximumFuel, "REDACTED", i)
		assert.NotEqual(t, primaryOptions.Ledger.MaximumOctets, handlerOptions.Ledger.MaximumOctets, "REDACTED", i)
	}

	//
	deadlineWaitCluster(N, func(j int) {
		<-ledgersEnrollments[j].Out()
	})

	//
	//
	time.Sleep(10 * time.Millisecond)

	//
	ultimateOptions := handlers[0].agreementOptions.Load()
	assert.Equal(t, newMaximumOctets, ultimateOptions.Ledger.MaximumOctets, "REDACTED")
}

//
//
type agreementOptionsRefreshingApplication struct {
	iface.Software
	modifyAtLevel int64
	newMaximumOctets    int64
	newMaximumFuel      int64
	t              *testing.T
}

func (app *agreementOptionsRefreshingApplication) CompleteLedger(
	ctx context.Context,
	req *iface.QueryCompleteLedger,
) (*iface.ReplyCompleteLedger, error) {
	//
	reply, err := app.Software.CompleteLedger(ctx, req)
	if err != nil {
		return nil, err
	}

	//
	if req.Level == app.modifyAtLevel {
		reply.AgreementArgumentRefreshes = &engineproto.AgreementOptions{
			Ledger: &engineproto.LedgerOptions{
				MaximumOctets: app.newMaximumOctets,
				MaximumFuel:   app.newMaximumFuel,
			},
		}
		app.t.Log("REDACTED", "REDACTED", req.Level, "REDACTED", app.modifyAtLevel)
	}

	return reply, nil
}

//
