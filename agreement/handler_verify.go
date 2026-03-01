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

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	controlkinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement/kinds"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/digits"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	nodestub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/simulate"
	strongmindcons "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/agreement"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	machinestubs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
)

//
//

var fallbackVerifyMoment = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

func initiateAgreementNetwork(t *testing.T, css []*Status, n int) (
	[]*Handler,
	[]kinds.Listening,
	[]*kinds.IncidentChannel,
) {
	engines := make([]*Handler, n)
	ledgersSubscriptions := make([]kinds.Listening, 0)
	incidentPipes := make([]*kinds.IncidentChannel, n)
	for i := 0; i < n; i++ {
		/*)
*/
		engines[i] = FreshHandler(css[i], true) //
		engines[i].AssignTracer(css[i].Tracer)

		//
		incidentPipes[i] = css[i].incidentChannel
		engines[i].AssignIncidentChannel(incidentPipes[i])

		ledgersUnder, err := incidentPipes[i].Listen(context.Background(), verifyListener, kinds.IncidentInquireFreshLedger)
		require.NoError(t, err)
		ledgersSubscriptions = append(ledgersSubscriptions, ledgersUnder)

		if css[i].status.FinalLedgerAltitude == 0 { //
			if err := css[i].ledgerExecute.Depot().Persist(css[i].status); err != nil {
				t.Error(err)
			}
		}
	}
	//
	p2p.CreateAssociatedRouters(settings.P2P, n, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", engines[i])
		s.AssignTracer(engines[i].connectionSTR.Tracer.Using("REDACTED", "REDACTED"))
		return s
	}, p2p.Connect2routers)

	//
	//
	//
	//
	for i := 0; i < n; i++ {
		s := engines[i].connectionSTR.ObtainStatus()
		engines[i].RouterTowardAgreement(s, false)
	}
	return engines, ledgersSubscriptions, incidentPipes
}

func haltAgreementNetwork(tracer log.Tracer, engines []*Handler, incidentPipes []*kinds.IncidentChannel) {
	tracer.Details("REDACTED", "REDACTED", len(engines))
	for i, r := range engines {
		tracer.Details("REDACTED", "REDACTED", i)
		if err := r.Router.Halt(); err != nil {
			tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	for i, b := range incidentPipes {
		tracer.Details("REDACTED", "REDACTED", i)
		if err := b.Halt(); err != nil {
			tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	tracer.Details("REDACTED", "REDACTED", len(engines))
}

//
func VerifyHandlerFundamental(t *testing.T) {
	N := 4
	css, sanitize := arbitraryAgreementNetwork(t, N, "REDACTED", freshSimulateMetronomeMethod(true), freshTokvalDepot)
	defer sanitize()
	engines, ledgersSubscriptions, incidentPipes := initiateAgreementNetwork(t, css, N)
	defer haltAgreementNetwork(log.VerifyingTracer(), engines, incidentPipes)
	//
	deadlinePauseCluster(N, func(j int) {
		<-ledgersSubscriptions[j].Out()
	})
}

//
func VerifyHandlerUsingProof(t *testing.T) {
	nthAssessors := 4
	verifyAlias := "REDACTED"
	metronomeMethod := freshSimulateMetronomeMethod(true)
	applicationMethod := freshTokvalDepot

	//
	//
	//

	producePaper, privateItems := arbitraryOriginPaper(nthAssessors, false, 30, nil)
	css := make([]*Status, nthAssessors)
	tracer := agreementTracer()
	for i := 0; i < nthAssessors; i++ {
		statusDatastore := dbm.FreshMemoryDatastore() //
		statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
			EjectIfaceReplies: false,
		})
		status, _ := statusDepot.FetchOriginatingDatastoreEitherOriginPaper(producePaper)
		thatSettings := RestoreSettings(fmt.Sprintf("REDACTED", verifyAlias, i))
		defer os.RemoveAll(thatSettings.OriginPath)
		assurePath(path.Dir(thatSettings.Agreement.JournalRecord()), 0o700) //
		app := applicationMethod()
		values := kinds.Temp2buffer.AssessorRevisions(status.Assessors)
		_, err := app.InitializeSuccession(context.Background(), &iface.SolicitInitializeSuccession{Assessors: values})
		require.NoError(t, err)

		pv := privateItems[i]
		//
		//

		ledgerDatastore := dbm.FreshMemoryDatastore()
		ledgerDepot := depot.FreshLedgerDepot(ledgerDatastore)

		mtx := new(commitchronize.Exclusion)
		txpoollTelemetry := txpooll.NooperationTelemetry()
		//
		delegateApplicationLinkConnection := delegate.FreshApplicationLinkAgreement(abcicustomer.FreshRegionalCustomer(mtx, app), delegate.NooperationTelemetry())
		delegateApplicationLinkMemory := delegate.FreshApplicationLinkTxpool(abcicustomer.FreshRegionalCustomer(mtx, app), delegate.NooperationTelemetry())

		//
		txpool := txpooll.FreshCNCatalogTxpool(settings.Txpool,
			delegateApplicationLinkMemory,
			status.FinalLedgerAltitude,
			txpooll.UsingTelemetry(txpoollTelemetry),
			txpooll.UsingPriorInspect(sm.TransferPriorInspect(status)),
			txpooll.UsingRelayInspect(sm.TransferRelayInspect(status)))

		if thatSettings.Agreement.PauseForeachTrans() {
			txpool.ActivateTransAccessible()
		}

		//
		//
		verOffset := (i + 1) % nthAssessors
		ev, err := kinds.FreshSimulateReplicatedBallotProofUsingAssessor(1, fallbackVerifyMoment, privateItems[verOffset], producePaper.SuccessionUUID)
		require.NoError(t, err)
		incidentpool := &machinestubs.ProofHub{}
		incidentpool.On("REDACTED", mock.AnythingOfType("REDACTED")).Return(nil)
		incidentpool.On("REDACTED", mock.AnythingOfType("REDACTED")).Return([]kinds.Proof{
			ev,
		}, int64(len(ev.Octets())))
		incidentpool.On("REDACTED", mock.AnythingOfType("REDACTED"), mock.AnythingOfType("REDACTED")).Return()

		incidentpool2 := sm.VoidProofHub{}

		//
		ledgerExecute := sm.FreshLedgerHandler(statusDepot, log.VerifyingTracer(), delegateApplicationLinkConnection, txpool, incidentpool, ledgerDepot)
		cs := FreshStatus(thatSettings.Agreement, status, ledgerExecute, ledgerDepot, txpool, incidentpool2)
		cs.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
		cs.AssignPrivateAssessor(pv)

		incidentChannel := kinds.FreshIncidentPipeline()
		incidentChannel.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
		err = incidentChannel.Initiate()
		require.NoError(t, err)
		cs.AssignIncidentChannel(incidentChannel)

		cs.AssignDeadlineMetronome(metronomeMethod())
		cs.AssignTracer(tracer.Using("REDACTED", i, "REDACTED", "REDACTED"))

		css[i] = cs
	}

	engines, ledgersSubscriptions, incidentPipes := initiateAgreementNetwork(t, css, nthAssessors)
	defer haltAgreementNetwork(log.VerifyingTracer(), engines, incidentPipes)

	//
	for i := 0; i < nthAssessors; i++ {
		deadlinePauseCluster(nthAssessors, func(j int) {
			msg := <-ledgersSubscriptions[j].Out()
			ledger := msg.Data().(kinds.IncidentDataFreshLedger).Ledger
			assert.Len(t, ledger.Proof.Proof, 1)
		})
	}
}

//

//
func VerifyHandlerGeneratesLedgerWheneverBlankLedgersNo(t *testing.T) {
	N := 4
	css, sanitize := arbitraryAgreementNetwork(t, N, "REDACTED", freshSimulateMetronomeMethod(true), freshTokvalDepot,
		func(c *cfg.Settings) {
			c.Agreement.GenerateVoidLedgers = false
		})
	defer sanitize()
	engines, ledgersSubscriptions, incidentPipes := initiateAgreementNetwork(t, css, N)
	defer haltAgreementNetwork(log.VerifyingTracer(), engines, incidentPipes)

	//
	if err := attestTxpool(css[3].transferObserver).InspectTransfer(statedepot.FreshTransferOriginatingUUID(1), func(reply *iface.ReplyInspectTransfer) {
		require.False(t, reply.EqualsFault())
	}, txpooll.TransferDetails{}); err != nil {
		t.Error(err)
	}

	//
	deadlinePauseCluster(N, func(j int) {
		<-ledgersSubscriptions[j].Out()
	})
}

func VerifyHandlerAcceptExecutesNegationAlarmConditionAppendNodeNotyetOccurredInvokedStill(t *testing.T) {
	N := 1
	css, sanitize := arbitraryAgreementNetwork(t, N, "REDACTED", freshSimulateMetronomeMethod(true), freshTokvalDepot)
	defer sanitize()
	engines, _, incidentPipes := initiateAgreementNetwork(t, css, N)
	defer haltAgreementNetwork(log.VerifyingTracer(), engines, incidentPipes)

	var (
		handler = engines[0]
		node    = nodestub.FreshNode(nil)
	)

	handler.InitializeNode(node)

	//
	assert.NotPanics(t, func() {
		handler.Accept(p2p.Wrapper{
			ConduitUUID: StatusConduit,
			Src:       node,
			Signal: &strongmindcons.OwnsBallot{
				Altitude: 1,
				Iteration:  1,
				Ordinal:  1,
				Kind:   commitchema.PreballotKind,
			},
		})
		handler.AppendNode(node)
	})
}

func VerifyHandlerAcceptAlarmsConditionInitializeNodeNotyetOccurredInvokedStill(t *testing.T) {
	N := 1
	css, sanitize := arbitraryAgreementNetwork(t, N, "REDACTED", freshSimulateMetronomeMethod(true), freshTokvalDepot)
	defer sanitize()
	engines, _, incidentPipes := initiateAgreementNetwork(t, css, N)
	defer haltAgreementNetwork(log.VerifyingTracer(), engines, incidentPipes)

	var (
		handler = engines[0]
		node    = nodestub.FreshNode(nil)
	)

	//

	//
	assert.Panics(t, func() {
		handler.Accept(p2p.Wrapper{
			ConduitUUID: StatusConduit,
			Src:       node,
			Signal: &strongmindcons.OwnsBallot{
				Altitude: 1,
				Iteration:  1,
				Ordinal:  1,
				Kind:   commitchema.PreballotKind,
			},
		})
	})
}

//
//
func VerifyRouterTowardAgreementBallotAdditions(t *testing.T) {
	for _, verifyInstance := range []struct {
		alias                  string
		persistedAltitude          int64
		primaryMandatoryAltitude int64
		encompassAdditions     bool
		mustAlarm           bool
	}{
		{
			alias:                  "REDACTED",
			primaryMandatoryAltitude: 0,
			persistedAltitude:          2,
			encompassAdditions:     false,
			mustAlarm:           false,
		},
		{
			alias:                  "REDACTED",
			primaryMandatoryAltitude: 2,
			persistedAltitude:          2,
			encompassAdditions:     false,
			mustAlarm:           true,
		},
		{
			alias:                  "REDACTED",
			primaryMandatoryAltitude: 3,
			persistedAltitude:          2,
			encompassAdditions:     false,
			mustAlarm:           false,
		},
		{
			alias:                  "REDACTED",
			primaryMandatoryAltitude: 1,
			persistedAltitude:          2,
			encompassAdditions:     false,
			mustAlarm:           true,
		},
		{
			alias:                  "REDACTED",
			primaryMandatoryAltitude: 1,
			persistedAltitude:          2,
			encompassAdditions:     true,
			mustAlarm:           false,
		},
	} {
		t.Run(verifyInstance.alias, func(t *testing.T) {
			ctx := t.Context()

			cs, vs := arbitraryStatus(1)
			assessor := vs[0]
			assessor.Altitude = verifyInstance.persistedAltitude

			cs.status.FinalLedgerAltitude = verifyInstance.persistedAltitude
			cs.status.FinalAssessors = cs.status.Assessors.Duplicate()
			cs.status.AgreementSettings.Iface.BallotAdditionsActivateAltitude = verifyInstance.primaryMandatoryAltitude

			itemLedger, err := cs.generateNominationLedger(ctx)
			require.NoError(t, err)

			//
			cs.Altitude = verifyInstance.persistedAltitude + 1
			itemLedger.Altitude = verifyInstance.persistedAltitude
			ledgerFragments, err := itemLedger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
			require.NoError(t, err)

			var ballotAssign *kinds.BallotAssign
			if verifyInstance.encompassAdditions {
				ballotAssign = kinds.FreshExpandedBallotAssign(cs.status.SuccessionUUID, verifyInstance.persistedAltitude, 0, commitchema.PreendorseKind, cs.status.Assessors)
			} else {
				ballotAssign = kinds.FreshBallotAssign(cs.status.SuccessionUUID, verifyInstance.persistedAltitude, 0, commitchema.PreendorseKind, cs.status.Assessors)
			}
			notatedBallot := attestBallot(assessor, commitchema.PreendorseKind, itemLedger.Digest(), ledgerFragments.Heading(), verifyInstance.encompassAdditions)

			var verAltitude int64
			if verifyInstance.encompassAdditions {
				require.NotNil(t, notatedBallot.AdditionNotation)
				verAltitude = verifyInstance.persistedAltitude
			} else {
				require.Nil(t, notatedBallot.Addition)
				require.Nil(t, notatedBallot.AdditionNotation)
			}

			appended, err := ballotAssign.AppendBallot(notatedBallot)
			require.NoError(t, err)
			require.True(t, appended)

			verAltitudeArgument := kinds.IfaceParameters{BallotAdditionsActivateAltitude: verAltitude}
			if verifyInstance.encompassAdditions {
				cs.ledgerDepot.PersistLedgerUsingExpandedEndorse(itemLedger, ledgerFragments, ballotAssign.CreateExpandedEndorse(verAltitudeArgument))
			} else {
				cs.ledgerDepot.PersistLedger(itemLedger, ledgerFragments, ballotAssign.CreateExpandedEndorse(verAltitudeArgument).TowardEndorse())
			}
			handler := FreshHandler(
				cs,
				true,
			)

			if verifyInstance.mustAlarm {
				assert.Panics(t, func() {
					handler.RouterTowardAgreement(cs.status, false)
				})
			} else {
				handler.RouterTowardAgreement(cs.status, false)
			}
		})
	}
}

//
func VerifyHandlerEntriesBallotsAlsoLedgerFragments(t *testing.T) {
	N := 4
	css, sanitize := arbitraryAgreementNetwork(t, N, "REDACTED", freshSimulateMetronomeMethod(true), freshTokvalDepot)
	defer sanitize()
	engines, ledgersSubscriptions, incidentPipes := initiateAgreementNetwork(t, css, N)
	defer haltAgreementNetwork(log.VerifyingTracer(), engines, incidentPipes)

	//
	deadlinePauseCluster(N, func(j int) {
		<-ledgersSubscriptions[j].Out()
	})

	//
	node := engines[1].Router.Nodes().Duplicate()[0]
	//
	ps := node.Get(kinds.NodeStatusToken).(*NodeStatus)

	assert.Equal(t, true, ps.BallotsRelayed() > 0, "REDACTED")
	assert.Equal(t, true, ps.LedgerFragmentsRelayed() > 0, "REDACTED")
}

//
//

func VerifyHandlerBallotingPotencyModify(t *testing.T) {
	nthValues := 4
	tracer := log.VerifyingTracer()
	css, sanitize := arbitraryAgreementNetwork(
		t,
		nthValues,
		"REDACTED",
		freshSimulateMetronomeMethod(true),
		freshEnduringTokvalDepot)
	defer sanitize()
	engines, ledgersSubscriptions, incidentPipes := initiateAgreementNetwork(t, css, nthValues)
	defer haltAgreementNetwork(tracer, engines, incidentPipes)

	//
	dynamicValues := make(map[string]struct{})
	for i := 0; i < nthValues; i++ {
		publicToken, err := css[i].privateAssessor.ObtainPublicToken()
		require.NoError(t, err)
		location := publicToken.Location()
		dynamicValues[string(location)] = struct{}{}
	}

	//
	deadlinePauseCluster(nthValues, func(j int) {
		<-ledgersSubscriptions[j].Out()
	})

	//
	tracer.Diagnose("REDACTED")

	assessor1keyToken, err := css[0].privateAssessor.ObtainPublicToken()
	require.NoError(t, err)

	assessor1keyTokenIface, err := cryptocode.PublicTokenTowardSchema(assessor1keyToken)
	require.NoError(t, err)
	reviseAssessorTransfer := statedepot.CreateItemAssignModifyTransfer(assessor1keyTokenIface, 25)
	precedingSumBallotingPotency := css[0].ObtainIterationStatus().FinalAssessors.SumBallotingPotency()

	pauseForeachAlsoCertifyLedger(t, nthValues, dynamicValues, ledgersSubscriptions, css, reviseAssessorTransfer)
	pauseForeachAlsoCertifyLedgerUsingTransfer(t, nthValues, dynamicValues, ledgersSubscriptions, css, reviseAssessorTransfer)
	pauseForeachAlsoCertifyLedger(t, nthValues, dynamicValues, ledgersSubscriptions, css)
	pauseForeachAlsoCertifyLedger(t, nthValues, dynamicValues, ledgersSubscriptions, css)

	if css[0].ObtainIterationStatus().FinalAssessors.SumBallotingPotency() == precedingSumBallotingPotency {
		t.Fatalf(
			"REDACTED",
			precedingSumBallotingPotency,
			css[0].ObtainIterationStatus().FinalAssessors.SumBallotingPotency())
	}

	reviseAssessorTransfer = statedepot.CreateItemAssignModifyTransfer(assessor1keyTokenIface, 2)
	precedingSumBallotingPotency = css[0].ObtainIterationStatus().FinalAssessors.SumBallotingPotency()

	pauseForeachAlsoCertifyLedger(t, nthValues, dynamicValues, ledgersSubscriptions, css, reviseAssessorTransfer)
	pauseForeachAlsoCertifyLedgerUsingTransfer(t, nthValues, dynamicValues, ledgersSubscriptions, css, reviseAssessorTransfer)
	pauseForeachAlsoCertifyLedger(t, nthValues, dynamicValues, ledgersSubscriptions, css)
	pauseForeachAlsoCertifyLedger(t, nthValues, dynamicValues, ledgersSubscriptions, css)

	if css[0].ObtainIterationStatus().FinalAssessors.SumBallotingPotency() == precedingSumBallotingPotency {
		t.Fatalf(
			"REDACTED",
			precedingSumBallotingPotency,
			css[0].ObtainIterationStatus().FinalAssessors.SumBallotingPotency())
	}

	reviseAssessorTransfer = statedepot.CreateItemAssignModifyTransfer(assessor1keyTokenIface, 26)
	precedingSumBallotingPotency = css[0].ObtainIterationStatus().FinalAssessors.SumBallotingPotency()

	pauseForeachAlsoCertifyLedger(t, nthValues, dynamicValues, ledgersSubscriptions, css, reviseAssessorTransfer)
	pauseForeachAlsoCertifyLedgerUsingTransfer(t, nthValues, dynamicValues, ledgersSubscriptions, css, reviseAssessorTransfer)
	pauseForeachAlsoCertifyLedger(t, nthValues, dynamicValues, ledgersSubscriptions, css)
	pauseForeachAlsoCertifyLedger(t, nthValues, dynamicValues, ledgersSubscriptions, css)

	if css[0].ObtainIterationStatus().FinalAssessors.SumBallotingPotency() == precedingSumBallotingPotency {
		t.Fatalf(
			"REDACTED",
			precedingSumBallotingPotency,
			css[0].ObtainIterationStatus().FinalAssessors.SumBallotingPotency())
	}
}

func VerifyHandlerAssessorAssignModifications(t *testing.T) {
	nthNodes := 7
	nthValues := 4
	css, _, _, sanitize := arbitraryAgreementNetworkUsingNodes(
		t,
		nthValues,
		nthNodes,
		"REDACTED",
		freshSimulateMetronomeMethod(true),
		freshEnduringTokvalDepotUsingRoute)

	defer sanitize()
	tracer := log.VerifyingTracer()

	engines, ledgersSubscriptions, incidentPipes := initiateAgreementNetwork(t, css, nthNodes)
	defer haltAgreementNetwork(tracer, engines, incidentPipes)

	//
	dynamicValues := make(map[string]struct{})
	for i := 0; i < nthValues; i++ {
		publicToken, err := css[i].privateAssessor.ObtainPublicToken()
		require.NoError(t, err)
		dynamicValues[string(publicToken.Location())] = struct{}{}
	}

	//
	deadlinePauseCluster(nthNodes, func(j int) {
		<-ledgersSubscriptions[j].Out()
	})

	t.Run("REDACTED", func(t *testing.T) {
		freshAssessorPublicToken1, err := css[nthValues].privateAssessor.ObtainPublicToken()
		assert.NoError(t, err)
		itemPublicToken1iface, err := cryptocode.PublicTokenTowardSchema(freshAssessorPublicToken1)
		assert.NoError(t, err)
		freshAssessorTransfer1 := statedepot.CreateItemAssignModifyTransfer(itemPublicToken1iface, verifyMinimumPotency)

		//
		//
		//
		pauseForeachAlsoCertifyLedger(t, nthNodes, dynamicValues, ledgersSubscriptions, css, freshAssessorTransfer1)

		//
		//
		pauseForeachAlsoCertifyLedgerUsingTransfer(t, nthNodes, dynamicValues, ledgersSubscriptions, css, freshAssessorTransfer1)

		//
		//
		pauseForeachAlsoCertifyLedger(t, nthNodes, dynamicValues, ledgersSubscriptions, css)

		//
		dynamicValues[string(freshAssessorPublicToken1.Location())] = struct{}{}

		//
		//
		pauseForeachLedgerUsingRevisedValuesAlsoCertifyThat(t, nthNodes, dynamicValues, ledgersSubscriptions, css)
	})

	t.Run("REDACTED", func(t *testing.T) {
		reviseAssessorPublicToken1, err := css[nthValues].privateAssessor.ObtainPublicToken()
		require.NoError(t, err)
		revisePublicToken1iface, err := cryptocode.PublicTokenTowardSchema(reviseAssessorPublicToken1)
		require.NoError(t, err)
		reviseAssessorTransfer1 := statedepot.CreateItemAssignModifyTransfer(revisePublicToken1iface, 25)
		precedingSumBallotingPotency := css[nthValues].ObtainIterationStatus().FinalAssessors.SumBallotingPotency()

		pauseForeachAlsoCertifyLedger(t, nthNodes, dynamicValues, ledgersSubscriptions, css, reviseAssessorTransfer1)
		pauseForeachAlsoCertifyLedgerUsingTransfer(t, nthNodes, dynamicValues, ledgersSubscriptions, css, reviseAssessorTransfer1)
		pauseForeachAlsoCertifyLedger(t, nthNodes, dynamicValues, ledgersSubscriptions, css)
		pauseForeachLedgerUsingRevisedValuesAlsoCertifyThat(t, nthNodes, dynamicValues, ledgersSubscriptions, css)

		if css[nthValues].ObtainIterationStatus().FinalAssessors.SumBallotingPotency() == precedingSumBallotingPotency {
			t.Errorf(
				"REDACTED",
				precedingSumBallotingPotency,
				css[nthValues].ObtainIterationStatus().FinalAssessors.SumBallotingPotency())
		}
	})

	freshAssessorPublicToken2, err := css[nthValues+1].privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	freshAssessor2iface, err := cryptocode.PublicTokenTowardSchema(freshAssessorPublicToken2)
	require.NoError(t, err)
	freshAssessorTransfer2 := statedepot.CreateItemAssignModifyTransfer(freshAssessor2iface, verifyMinimumPotency)

	freshAssessorPublicToken3, err := css[nthValues+2].privateAssessor.ObtainPublicToken()
	require.NoError(t, err)
	freshAssessor3iface, err := cryptocode.PublicTokenTowardSchema(freshAssessorPublicToken3)
	require.NoError(t, err)
	freshAssessorTransfer3 := statedepot.CreateItemAssignModifyTransfer(freshAssessor3iface, verifyMinimumPotency)

	t.Run("REDACTED", func(t *testing.T) {
		pauseForeachAlsoCertifyLedger(t, nthNodes, dynamicValues, ledgersSubscriptions, css, freshAssessorTransfer2, freshAssessorTransfer3)
		pauseForeachAlsoCertifyLedgerUsingTransfer(t, nthNodes, dynamicValues, ledgersSubscriptions, css, freshAssessorTransfer2, freshAssessorTransfer3)
		pauseForeachAlsoCertifyLedger(t, nthNodes, dynamicValues, ledgersSubscriptions, css)
		dynamicValues[string(freshAssessorPublicToken2.Location())] = struct{}{}
		dynamicValues[string(freshAssessorPublicToken3.Location())] = struct{}{}
		pauseForeachLedgerUsingRevisedValuesAlsoCertifyThat(t, nthNodes, dynamicValues, ledgersSubscriptions, css)
	})

	t.Run("REDACTED", func(t *testing.T) {
		discardAssessorTransfer2 := statedepot.CreateItemAssignModifyTransfer(freshAssessor2iface, 0)
		discardAssessorTransfer3 := statedepot.CreateItemAssignModifyTransfer(freshAssessor3iface, 0)

		pauseForeachAlsoCertifyLedger(t, nthNodes, dynamicValues, ledgersSubscriptions, css, discardAssessorTransfer2, discardAssessorTransfer3)
		pauseForeachAlsoCertifyLedgerUsingTransfer(t, nthNodes, dynamicValues, ledgersSubscriptions, css, discardAssessorTransfer2, discardAssessorTransfer3)
		pauseForeachAlsoCertifyLedger(t, nthNodes, dynamicValues, ledgersSubscriptions, css)
		delete(dynamicValues, string(freshAssessorPublicToken2.Location()))
		delete(dynamicValues, string(freshAssessorPublicToken3.Location()))
		pauseForeachLedgerUsingRevisedValuesAlsoCertifyThat(t, nthNodes, dynamicValues, ledgersSubscriptions, css)
	})
}

//
func VerifyHandlerUsingDeadlineEndorse(t *testing.T) {
	N := 4
	css, sanitize := arbitraryAgreementNetwork(t, N, "REDACTED", freshSimulateMetronomeMethod(false), freshTokvalDepot)
	defer sanitize()
	//
	for i := 0; i < N; i++ {
		css[i].settings.OmitDeadlineEndorse = false
	}

	engines, ledgersSubscriptions, incidentPipes := initiateAgreementNetwork(t, css, N-1)
	defer haltAgreementNetwork(log.VerifyingTracer(), engines, incidentPipes)

	//
	deadlinePauseCluster(N-1, func(j int) {
		<-ledgersSubscriptions[j].Out()
	})
}

func pauseForeachAlsoCertifyLedger(
	t *testing.T,
	n int,
	dynamicValues map[string]struct{},
	ledgersSubscriptions []kinds.Listening,
	css []*Status,
	txs ...[]byte,
) {
	deadlinePauseCluster(n, func(j int) {
		css[j].Tracer.Diagnose("REDACTED")
		msg := <-ledgersSubscriptions[j].Out()
		freshLedger := msg.Data().(kinds.IncidentDataFreshLedger).Ledger
		css[j].Tracer.Diagnose("REDACTED", "REDACTED", freshLedger.Altitude)
		err := certifyLedger(freshLedger, dynamicValues)
		require.NoError(t, err)

		//
		for _, tx := range txs {
			err := attestTxpool(css[j].transferObserver).InspectTransfer(tx, func(reply *iface.ReplyInspectTransfer) {
				require.False(t, reply.EqualsFault())
				fmt.Println(reply)
			}, txpooll.TransferDetails{})
			require.NoError(t, err)
		}
	})
}

func pauseForeachAlsoCertifyLedgerUsingTransfer(
	t *testing.T,
	n int,
	dynamicValues map[string]struct{},
	ledgersSubscriptions []kinds.Listening,
	css []*Status,
	txs ...[]byte,
) {
	deadlinePauseCluster(n, func(j int) {
		ntransactions := 0
	LEDGER_TRANSFER_CYCLE:
		for {
			css[j].Tracer.Diagnose("REDACTED", "REDACTED", ntransactions)
			msg := <-ledgersSubscriptions[j].Out()
			freshLedger := msg.Data().(kinds.IncidentDataFreshLedger).Ledger
			css[j].Tracer.Diagnose("REDACTED", "REDACTED", freshLedger.Altitude)
			err := certifyLedger(freshLedger, dynamicValues)
			require.NoError(t, err)

			//
			//
			//
			for _, tx := range freshLedger.Txs {
				assert.EqualValues(t, txs[ntransactions], tx)
				ntransactions++
			}

			if ntransactions == len(txs) {
				break LEDGER_TRANSFER_CYCLE
			}
		}
	})
}

func pauseForeachLedgerUsingRevisedValuesAlsoCertifyThat(
	t *testing.T,
	n int,
	revisedValues map[string]struct{},
	ledgersSubscriptions []kinds.Listening,
	css []*Status,
) {
	deadlinePauseCluster(n, func(j int) {
		var freshLedger *kinds.Ledger
	Cycle:
		for {
			css[j].Tracer.Diagnose("REDACTED")
			msg := <-ledgersSubscriptions[j].Out()
			freshLedger = msg.Data().(kinds.IncidentDataFreshLedger).Ledger
			if freshLedger.FinalEndorse.Extent() == len(revisedValues) {
				css[j].Tracer.Diagnose("REDACTED", "REDACTED", freshLedger.Altitude)
				break Cycle
			}
			css[j].Tracer.Diagnose(
				"REDACTED",
				"REDACTED", freshLedger.Altitude, "REDACTED", freshLedger.FinalEndorse.Extent(), "REDACTED", len(revisedValues),
			)
		}

		err := certifyLedger(freshLedger, revisedValues)
		assert.Nil(t, err)
	})
}

//
func certifyLedger(ledger *kinds.Ledger, dynamicValues map[string]struct{}) error {
	if ledger.FinalEndorse.Extent() != len(dynamicValues) {
		return fmt.Errorf(
			"REDACTED",
			ledger.FinalEndorse.Extent(),
			len(dynamicValues))
	}

	for _, endorseSignature := range ledger.FinalEndorse.Notations {
		if _, ok := dynamicValues[string(endorseSignature.AssessorLocation)]; !ok {
			return fmt.Errorf("REDACTED", endorseSignature.AssessorLocation)
		}
	}
	return nil
}

func deadlinePauseCluster(n int, f func(int)) {
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(j int) {
			f(j)
			wg.Done()
		}(i)
	}

	complete := make(chan struct{})
	go func() {
		wg.Wait()
		close(complete)
	}()

	//
	//
	deadline := time.Second * 20

	select {
	case <-complete:
	case <-time.After(deadline):
		panic("REDACTED")
	}
}

//
//

func VerifyFreshIterationPhaseSignalCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		anticipateFault              bool
		signalIteration           int32
		signalFinalEndorseIteration int32
		signalAltitude          int64
		verifyAlias               string
		signalPhase            controlkinds.IterationPhaseKind
	}{
		{false, 0, 0, 0, "REDACTED", controlkinds.IterationPhaseFreshAltitude},
		{true, -1, 0, 0, "REDACTED", controlkinds.IterationPhaseFreshAltitude},
		{true, 0, 0, -1, "REDACTED", controlkinds.IterationPhaseFreshAltitude},
		{true, 0, 0, 0, "REDACTED", controlkinds.IterationPhaseEndorse + 1},
		//
		{false, 0, 0, 1, "REDACTED", controlkinds.IterationPhaseFreshAltitude},
		{false, 0, -1, 2, "REDACTED", controlkinds.IterationPhaseFreshAltitude},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			signal := FreshIterationPhaseSignal{
				Altitude:          tc.signalAltitude,
				Iteration:           tc.signalIteration,
				Phase:            tc.signalPhase,
				FinalEndorseIteration: tc.signalFinalEndorseIteration,
			}

			err := signal.CertifyFundamental()
			if tc.anticipateFault {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func VerifyFreshIterationPhaseSignalCertifyAltitude(t *testing.T) {
	primaryAltitude := int64(10)
	verifyScenarios := []struct {
		anticipateFault              bool
		signalFinalEndorseIteration int32
		signalAltitude          int64
		verifyAlias               string
	}{
		{false, 0, 11, "REDACTED"},
		{true, 0, -1, "REDACTED"},
		{true, 0, 0, "REDACTED"},
		{true, 0, 10, "REDACTED"},
		{true, -1, 11, "REDACTED"},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			signal := FreshIterationPhaseSignal{
				Altitude:          tc.signalAltitude,
				Iteration:           0,
				Phase:            controlkinds.IterationPhaseFreshAltitude,
				FinalEndorseIteration: tc.signalFinalEndorseIteration,
			}

			err := signal.CertifyAltitude(primaryAltitude)
			if tc.anticipateFault {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func VerifyFreshSoundLedgerSignalCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		distortProc func(*FreshSoundLedgerSignal)
		expirationFault     string
	}{
		{func(msg *FreshSoundLedgerSignal) {}, "REDACTED"},
		{func(msg *FreshSoundLedgerSignal) { msg.Altitude = -1 }, strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}.Failure()},
		{func(msg *FreshSoundLedgerSignal) { msg.Iteration = -1 }, strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}.Failure()},
		{
			func(msg *FreshSoundLedgerSignal) { msg.LedgerFragmentAssignHeading.Sum = 2 },
			"REDACTED",
		},
		{
			func(msg *FreshSoundLedgerSignal) {
				msg.LedgerFragmentAssignHeading.Sum = 0
				msg.LedgerFragments = digits.FreshDigitCollection(0)
			},
			strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}.Failure(),
		},
		{
			func(msg *FreshSoundLedgerSignal) { msg.LedgerFragments = digits.FreshDigitCollection(int(kinds.MaximumLedgerFragmentsTally) + 1) },
			"REDACTED",
		},
		{
			func(msg *FreshSoundLedgerSignal) { msg.LedgerFragments.Components = nil },
			"REDACTED",
		},
		{
			func(msg *FreshSoundLedgerSignal) { msg.LedgerFragments.Digits = 500 },
			"REDACTED",
		},
	}

	for i, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			msg := &FreshSoundLedgerSignal{
				Altitude: 1,
				Iteration:  0,
				LedgerFragmentAssignHeading: kinds.FragmentAssignHeading{
					Sum: 1,
				},
				LedgerFragments: digits.FreshDigitCollection(1),
			}

			tc.distortProc(msg)
			err := msg.CertifyFundamental()
			if tc.expirationFault != "REDACTED" && assert.Error(t, err) {
				assert.Contains(t, err.Error(), tc.expirationFault)
			}
		})
	}
}

func VerifyNominationPolicySignalCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		distortProc func(*NominationPolicySignal)
		expirationFault     string
	}{
		{func(msg *NominationPolicySignal) {}, "REDACTED"},
		{func(msg *NominationPolicySignal) { msg.Altitude = -1 }, strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}.Failure()},
		{func(msg *NominationPolicySignal) { msg.NominationPolicyIteration = -1 }, strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}.Failure()},
		{func(msg *NominationPolicySignal) { msg.NominationPolicy = digits.FreshDigitCollection(0) }, strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}.Failure()},
		{
			func(msg *NominationPolicySignal) { msg.NominationPolicy = digits.FreshDigitCollection(kinds.MaximumBallotsTally + 1) },
			"REDACTED",
		},
		{
			func(msg *NominationPolicySignal) { msg.NominationPolicy.Components = nil },
			"REDACTED",
		},
		{
			func(msg *NominationPolicySignal) { msg.NominationPolicy.Digits = 500 },
			"REDACTED",
		},
	}

	for i, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			msg := &NominationPolicySignal{
				Altitude:           1,
				NominationPolicyIteration: 1,
				NominationPolicy:      digits.FreshDigitCollection(1),
			}

			tc.distortProc(msg)
			err := msg.CertifyFundamental()
			if tc.expirationFault != "REDACTED" && assert.Error(t, err) {
				assert.Contains(t, err.Error(), tc.expirationFault)
			}
		})
	}
}

func VerifyLedgerFragmentSignalCertifyFundamental(t *testing.T) {
	verifyFragment := new(kinds.Fragment)
	verifyFragment.Attestation.NodeDigest = tenderminthash.Sum([]byte("REDACTED"))
	verifyScenarios := []struct {
		verifyAlias      string
		signalAltitude int64
		signalIteration  int32
		signalFragment   *kinds.Fragment
		anticipateFault     bool
	}{
		{"REDACTED", 0, 0, verifyFragment, false},
		{"REDACTED", -1, 0, verifyFragment, true},
		{"REDACTED", 0, -1, verifyFragment, true},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			signal := LedgerFragmentSignal{
				Altitude: tc.signalAltitude,
				Iteration:  tc.signalIteration,
				Fragment:   tc.signalFragment,
			}

			assert.Equal(t, tc.anticipateFault, signal.CertifyFundamental() != nil, "REDACTED")
		})
	}

	signal := LedgerFragmentSignal{Altitude: 0, Iteration: 0, Fragment: new(kinds.Fragment)}
	signal.Fragment.Ordinal = 1

	assert.Equal(t, true, signal.CertifyFundamental() != nil, "REDACTED")
}

func VerifyOwnsBallotSignalCertifyFundamental(t *testing.T) {
	const (
		soundNotatedSignalKind   commitchema.AttestedSignalKind = 0x01
		unfitNotatedSignalKind commitchema.AttestedSignalKind = 0x03
	)

	verifyScenarios := []struct {
		anticipateFault     bool
		signalIteration  int32
		signalOrdinal  int32
		signalAltitude int64
		verifyAlias      string
		signalKind   commitchema.AttestedSignalKind
	}{
		{false, 0, 0, 0, "REDACTED", soundNotatedSignalKind},
		{true, -1, 0, 0, "REDACTED", soundNotatedSignalKind},
		{true, 0, -1, 0, "REDACTED", soundNotatedSignalKind},
		{true, 0, 0, 0, "REDACTED", unfitNotatedSignalKind},
		{true, 0, 0, -1, "REDACTED", soundNotatedSignalKind},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			signal := OwnsBallotSignal{
				Altitude: tc.signalAltitude,
				Iteration:  tc.signalIteration,
				Kind:   tc.signalKind,
				Ordinal:  tc.signalOrdinal,
			}

			assert.Equal(t, tc.anticipateFault, signal.CertifyFundamental() != nil, "REDACTED")
		})
	}
}

func VerifyBallotAssignMajor23signalCertifyFundamental(t *testing.T) {
	const (
		soundNotatedSignalKind   commitchema.AttestedSignalKind = 0x01
		unfitNotatedSignalKind commitchema.AttestedSignalKind = 0x03
	)

	soundLedgerUUID := kinds.LedgerUUID{}
	unfitLedgerUUID := kinds.LedgerUUID{
		Digest: octets.HexadecimalOctets{},
		FragmentAssignHeading: kinds.FragmentAssignHeading{
			Sum: 1,
			Digest:  []byte{0},
		},
	}

	verifyScenarios := []struct {
		anticipateFault      bool
		signalIteration   int32
		signalAltitude  int64
		verifyAlias       string
		signalKind    commitchema.AttestedSignalKind
		signalLedgerUUID kinds.LedgerUUID
	}{
		{false, 0, 0, "REDACTED", soundNotatedSignalKind, soundLedgerUUID},
		{true, -1, 0, "REDACTED", soundNotatedSignalKind, soundLedgerUUID},
		{true, 0, -1, "REDACTED", soundNotatedSignalKind, soundLedgerUUID},
		{true, 0, 0, "REDACTED", unfitNotatedSignalKind, soundLedgerUUID},
		{true, 0, 0, "REDACTED", soundNotatedSignalKind, unfitLedgerUUID},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			signal := BallotAssignMajor23signal{
				Altitude:  tc.signalAltitude,
				Iteration:   tc.signalIteration,
				Kind:    tc.signalKind,
				LedgerUUID: tc.signalLedgerUUID,
			}

			assert.Equal(t, tc.anticipateFault, signal.CertifyFundamental() != nil, "REDACTED")
		})
	}
}

func VerifyBallotAssignDigitsSignalCertifyFundamental(t *testing.T) {
	verifyScenarios := []struct {
		distortProc func(*BallotAssignDigitsSignal)
		expirationFault     string
	}{
		{func(msg *BallotAssignDigitsSignal) {}, "REDACTED"},
		{func(msg *BallotAssignDigitsSignal) { msg.Altitude = -1 }, strongminderrors.FaultAdverseAttribute{Attribute: "REDACTED"}.Failure()},
		{func(msg *BallotAssignDigitsSignal) { msg.Kind = 0x03 }, strongminderrors.FaultUnfitAttribute{Attribute: "REDACTED"}.Failure()},
		{func(msg *BallotAssignDigitsSignal) {
			msg.LedgerUUID = kinds.LedgerUUID{
				Digest: octets.HexadecimalOctets{},
				FragmentAssignHeading: kinds.FragmentAssignHeading{
					Sum: 1,
					Digest:  []byte{0},
				},
			}
		}, "REDACTED"},
		{
			func(msg *BallotAssignDigitsSignal) { msg.Ballots = digits.FreshDigitCollection(kinds.MaximumBallotsTally + 1) },
			"REDACTED",
		},
		{
			func(msg *BallotAssignDigitsSignal) { msg.Ballots.Components = nil },
			"REDACTED",
		},
		{
			func(msg *BallotAssignDigitsSignal) { msg.Ballots.Digits = 500 },
			"REDACTED",
		},
	}

	for i, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			msg := &BallotAssignDigitsSignal{
				Altitude:  1,
				Iteration:   0,
				Kind:    0x01,
				Ballots:   digits.FreshDigitCollection(1),
				LedgerUUID: kinds.LedgerUUID{},
			}

			tc.distortProc(msg)
			err := msg.CertifyFundamental()
			if tc.expirationFault != "REDACTED" && assert.Error(t, err) {
				assert.Contains(t, err.Error(), tc.expirationFault)
			}
		})
	}
}

func VerifySerializeJSNNodeStatus(t *testing.T) {
	ps := FreshNodeStatus(nil)
	data, err := jsn.Serialize(ps)
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

func VerifyBallotSignalCertifyFundamental(t *testing.T) {
	_, vss := arbitraryStatus(2)

	arbitraryOctets := commitrand.Octets(tenderminthash.Extent)
	ledgerUUID := kinds.LedgerUUID{
		Digest: arbitraryOctets,
		FragmentAssignHeading: kinds.FragmentAssignHeading{
			Sum: 1,
			Digest:  arbitraryOctets,
		},
	}
	ballot := attestBallot(vss[1], commitchema.PreendorseKind, arbitraryOctets, ledgerUUID.FragmentAssignHeading, true)

	verifyScenarios := []struct {
		distortProc func(*BallotSignal)
		expirationFault     string
	}{
		{func(_ *BallotSignal) {}, "REDACTED"},
		{func(msg *BallotSignal) { msg.Ballot.AssessorOrdinal = -1 }, "REDACTED"},
		//
		{func(msg *BallotSignal) { msg.Ballot.AssessorOrdinal = 1000 }, "REDACTED"},
	}

	for i, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", i), func(t *testing.T) {
			msg := &BallotSignal{ballot}

			tc.distortProc(msg)
			err := msg.CertifyFundamental()
			if tc.expirationFault != "REDACTED" && assert.Error(t, err) { //
				assert.Contains(t, err.Error(), tc.expirationFault)
			}
		})
	}
}

//
//
//
func VerifyHandlerAgreementParametersRevise(t *testing.T) {
	N := 4

	//
	reviseLocatedAltitude := int64(5)
	freshMaximumOctets := int64(5000000)
	freshMaximumFuel := int64(50000000)

	//
	applicationMethod := func() iface.Platform {
		tokvalApplication := statedepot.FreshInsideRamPlatform()
		return &agreementParametersRevisingApplication{
			Platform:    tokvalApplication,
			reviseLocatedAltitude: reviseLocatedAltitude,
			freshMaximumOctets:    freshMaximumOctets,
			freshMaximumFuel:      freshMaximumFuel,
			t:              t,
		}
	}

	css, sanitize := arbitraryAgreementNetwork(
		t,
		N,
		"REDACTED",
		freshSimulateMetronomeMethod(true),
		applicationMethod,
	)
	defer sanitize()

	engines, ledgersSubscriptions, incidentPipes := initiateAgreementNetwork(t, css, N)
	defer haltAgreementNetwork(log.VerifyingTracer(), engines, incidentPipes)

	//
	primaryParameters := engines[0].agreementParameters.Load()
	t.Logf("REDACTED", primaryParameters.Ledger.MaximumOctets, primaryParameters.Ledger.MaximumFuel)

	//
	deadlinePauseCluster(N, func(j int) {
		<-ledgersSubscriptions[j].Out()
	})

	//
	//
	time.Sleep(10 * time.Millisecond)

	//
	parametersSubsequentLedger1 := engines[0].agreementParameters.Load()
	assert.Equal(t, *primaryParameters, *parametersSubsequentLedger1, "REDACTED")

	//
	deadlinePauseCluster(N, func(j int) {
		for {
			msg := <-ledgersSubscriptions[j].Out()
			incident := msg.Data().(kinds.IncidentDataFreshLedger)
			if incident.Ledger.Altitude == reviseLocatedAltitude {
				break
			}
		}
	})

	//
	//
	time.Sleep(10 * time.Millisecond)

	//
	for i := range N {
		handlerParameters := engines[i].agreementParameters.Load()

		assert.Equal(t, freshMaximumOctets, handlerParameters.Ledger.MaximumOctets, "REDACTED", i)
		assert.Equal(t, freshMaximumFuel, handlerParameters.Ledger.MaximumFuel, "REDACTED", i)
		assert.NotEqual(t, primaryParameters.Ledger.MaximumOctets, handlerParameters.Ledger.MaximumOctets, "REDACTED", i)
	}

	//
	deadlinePauseCluster(N, func(j int) {
		<-ledgersSubscriptions[j].Out()
	})

	//
	//
	time.Sleep(10 * time.Millisecond)

	//
	ultimateParameters := engines[0].agreementParameters.Load()
	assert.Equal(t, freshMaximumOctets, ultimateParameters.Ledger.MaximumOctets, "REDACTED")
}

//
//
type agreementParametersRevisingApplication struct {
	iface.Platform
	reviseLocatedAltitude int64
	freshMaximumOctets    int64
	freshMaximumFuel      int64
	t              *testing.T
}

func (app *agreementParametersRevisingApplication) CulminateLedger(
	ctx context.Context,
	req *iface.SolicitCulminateLedger,
) (*iface.ReplyCulminateLedger, error) {
	//
	reply, err := app.Platform.CulminateLedger(ctx, req)
	if err != nil {
		return nil, err
	}

	//
	if req.Altitude == app.reviseLocatedAltitude {
		reply.AgreementArgumentRevisions = &commitchema.AgreementSettings{
			Ledger: &commitchema.LedgerParameters{
				MaximumOctets: app.freshMaximumOctets,
				MaximumFuel:   app.freshMaximumFuel,
			},
		}
		app.t.Log("REDACTED", "REDACTED", req.Altitude, "REDACTED", app.reviseLocatedAltitude)
	}

	return reply, nil
}

//
