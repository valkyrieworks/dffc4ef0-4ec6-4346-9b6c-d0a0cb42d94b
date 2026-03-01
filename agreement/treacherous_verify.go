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
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/proof"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	strongmindcons "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/agreement"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//

//
func VerifyTreacherousPreballotAmbiguity(t *testing.T) {
	ctx := t.Context()

	const nthAssessors = 4
	const treacherousPeer = 0
	const preballotAltitude = int64(2)
	verifyAlias := "REDACTED"
	metronomeMethod := freshSimulateMetronomeMethod(true)
	applicationMethod := freshTokvalDepot

	producePaper, privateItems := arbitraryOriginPaper(nthAssessors, false, 30, nil)
	css := make([]*Status, nthAssessors)

	for i := 0; i < nthAssessors; i++ {
		tracer := agreementTracer().Using("REDACTED", "REDACTED", "REDACTED", i)
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

		ledgerDatastore := dbm.FreshMemoryDatastore()
		ledgerDepot := depot.FreshLedgerDepot(ledgerDatastore)

		mtx := new(commitchronize.Exclusion)
		//
		delegateApplicationLinkConnection := delegate.FreshApplicationLinkAgreement(abcicustomer.FreshRegionalCustomer(mtx, app), delegate.NooperationTelemetry())
		delegateApplicationLinkMemory := delegate.FreshApplicationLinkTxpool(abcicustomer.FreshRegionalCustomer(mtx, app), delegate.NooperationTelemetry())

		//
		txpool := txpooll.FreshCNCatalogTxpool(settings.Txpool,
			delegateApplicationLinkMemory,
			status.FinalLedgerAltitude,
			txpooll.UsingPriorInspect(sm.TransferPriorInspect(status)),
			txpooll.UsingRelayInspect(sm.TransferRelayInspect(status)))

		if thatSettings.Agreement.PauseForeachTrans() {
			txpool.ActivateTransAccessible()
		}

		//
		proofDatastore := dbm.FreshMemoryDatastore()
		incidentpool, err := proof.FreshHub(proofDatastore, statusDepot, ledgerDepot)
		require.NoError(t, err)
		incidentpool.AssignTracer(tracer.Using("REDACTED", "REDACTED"))

		//
		ledgerExecute := sm.FreshLedgerHandler(statusDepot, log.VerifyingTracer(), delegateApplicationLinkConnection, txpool, incidentpool, ledgerDepot)
		cs := FreshStatus(thatSettings.Agreement, status, ledgerExecute, ledgerDepot, txpool, incidentpool)
		cs.AssignTracer(cs.Tracer)
		//
		pv := privateItems[i]
		cs.AssignPrivateAssessor(pv)

		incidentChannel := kinds.FreshIncidentPipeline()
		incidentChannel.AssignTracer(log.VerifyingTracer().Using("REDACTED", "REDACTED"))
		err = incidentChannel.Initiate()
		require.NoError(t, err)
		cs.AssignIncidentChannel(incidentChannel)

		cs.AssignDeadlineMetronome(metronomeMethod())
		cs.AssignTracer(tracer)

		css[i] = cs
	}

	//
	engines := make([]*Handler, nthAssessors)
	ledgersSubscriptions := make([]kinds.Listening, 0)
	incidentPipes := make([]*kinds.IncidentChannel, nthAssessors)
	for i := 0; i < nthAssessors; i++ {
		engines[i] = FreshHandler(css[i], true) //
		engines[i].AssignTracer(css[i].Tracer)

		//
		incidentPipes[i] = css[i].incidentChannel
		engines[i].AssignIncidentChannel(incidentPipes[i])

		ledgersUnder, err := incidentPipes[i].Listen(context.Background(), verifyListener, kinds.IncidentInquireFreshLedger, 100)
		require.NoError(t, err)
		ledgersSubscriptions = append(ledgersSubscriptions, ledgersUnder)

		if css[i].status.FinalLedgerAltitude == 0 { //
			err = css[i].ledgerExecute.Depot().Persist(css[i].status)
			require.NoError(t, err)
		}
	}
	//
	p2p.CreateAssociatedRouters(settings.P2P, nthAssessors, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", engines[i])
		s.AssignTracer(engines[i].connectionSTR.Tracer.Using("REDACTED", "REDACTED"))
		return s
	}, p2p.Connect2routers)

	//
	bcs := css[treacherousPeer]

	//
	bcs.performPreballot = func(altitude int64, iteration int32) {
		//
		if altitude == preballotAltitude {
			bcs.Tracer.Details("REDACTED")
			preballot1, err := bcs.attestBallot(commitchema.PreballotKind, bcs.NominationLedger.Digest(), bcs.NominationLedgerFragments.Heading(), nil)
			require.NoError(t, err)
			preballot2, err := bcs.attestBallot(commitchema.PreballotKind, nil, kinds.FragmentAssignHeading{}, nil)
			require.NoError(t, err)
			nodeCatalog := engines[treacherousPeer].Router.Nodes().Duplicate()
			bcs.Tracer.Details("REDACTED", "REDACTED", nodeCatalog)
			//
			for i, node := range nodeCatalog {
				if i < len(nodeCatalog)/2 {
					bcs.Tracer.Details("REDACTED", "REDACTED", preballot1, "REDACTED", node)
					node.Transmit(p2p.Wrapper{
						Signal:   &strongmindcons.Ballot{Ballot: preballot1.TowardSchema()},
						ConduitUUID: BallotConduit,
					})
				} else {
					bcs.Tracer.Details("REDACTED", "REDACTED", preballot2, "REDACTED", node)
					node.Transmit(p2p.Wrapper{
						Signal:   &strongmindcons.Ballot{Ballot: preballot2.TowardSchema()},
						ConduitUUID: BallotConduit,
					})
				}
			}
		} else {
			bcs.Tracer.Details("REDACTED")
			bcs.fallbackPerformPreballot(altitude, iteration)
		}
	}

	//
	//
	//
	idleNominator := css[1]

	idleNominator.resolveNomination = func(altitude int64, iteration int32) {
		idleNominator.Tracer.Details("REDACTED")
		if idleNominator.privateAssessor == nil {
			panic("REDACTED")
		}

		var addnEndorse *kinds.ExpandedEndorse
		switch {
		case idleNominator.Altitude == idleNominator.status.PrimaryAltitude:
			//
			//
			addnEndorse = &kinds.ExpandedEndorse{}
		case idleNominator.FinalEndorse.OwnsCoupleTrinityPreponderance():
			//
			verAltitudeArgument := kinds.IfaceParameters{BallotAdditionsActivateAltitude: altitude}
			addnEndorse = idleNominator.FinalEndorse.CreateExpandedEndorse(verAltitudeArgument)
		default: //
			idleNominator.Tracer.Failure("REDACTED")
			return
		}

		//
		addnEndorse.ExpandedNotations[len(addnEndorse.ExpandedNotations)-1] = kinds.FreshExpandedEndorseSignatureMissing()

		if idleNominator.privateAssessorPublicToken == nil {
			//
			//
			idleNominator.Tracer.Failure(fmt.Sprintf("REDACTED", FaultPublicTokenEqualsNegationAssign))
			return
		}
		nominatorLocation := idleNominator.privateAssessorPublicToken.Location()

		ledger, err := idleNominator.ledgerExecute.GenerateNominationLedger(
			ctx, idleNominator.Altitude, idleNominator.status, addnEndorse, nominatorLocation)
		require.NoError(t, err)
		ledgerFragments, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
		require.NoError(t, err)

		//
		//
		if err := idleNominator.wal.PurgeAlsoChronize(); err != nil {
			idleNominator.Tracer.Failure("REDACTED")
		}

		//
		itemLedgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: ledgerFragments.Heading()}
		nomination := kinds.FreshNomination(altitude, iteration, idleNominator.SoundIteration, itemLedgerUUID)
		p := nomination.TowardSchema()
		if err := idleNominator.privateAssessor.AttestNomination(idleNominator.status.SuccessionUUID, p); err == nil {
			nomination.Notation = p.Notation

			//
			idleNominator.transmitIntrinsicSignal(signalDetails{&NominationSignal{nomination}, "REDACTED"})
			for i := 0; i < int(ledgerFragments.Sum()); i++ {
				fragment := ledgerFragments.ObtainFragment(i)
				idleNominator.transmitIntrinsicSignal(signalDetails{&LedgerFragmentSignal{idleNominator.Altitude, idleNominator.Iteration, fragment}, "REDACTED"})
			}
			idleNominator.Tracer.Details("REDACTED", "REDACTED", altitude, "REDACTED", iteration, "REDACTED", nomination)
			idleNominator.Tracer.Diagnose(fmt.Sprintf("REDACTED", ledger))
		} else if !idleNominator.reenactStyle {
			idleNominator.Tracer.Failure("REDACTED", "REDACTED", altitude, "REDACTED", iteration, "REDACTED", err)
		}
	}

	//
	for i := 0; i < nthAssessors; i++ {
		s := engines[i].connectionSTR.ObtainStatus()
		engines[i].RouterTowardAgreement(s, false)
	}
	defer haltAgreementNetwork(log.VerifyingTracer(), engines, incidentPipes)

	//
	//
	proofOriginatingEveryAssessor := make([]kinds.Proof, nthAssessors)

	wg := new(sync.WaitGroup)
	for i := 0; i < nthAssessors; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for msg := range ledgersSubscriptions[i].Out() {
				ledger := msg.Data().(kinds.IncidentDataFreshLedger).Ledger
				if len(ledger.Proof.Proof) != 0 {
					proofOriginatingEveryAssessor[i] = ledger.Proof.Proof[0]
					return
				}
			}
		}(i)
	}

	complete := make(chan struct{})
	go func() {
		wg.Wait()
		close(complete)
	}()

	publickey, err := bcs.privateAssessor.ObtainPublicToken()
	require.NoError(t, err)

	select {
	case <-complete:
		for idx, ev := range proofOriginatingEveryAssessor {
			if assert.NotNil(t, ev, idx) {
				ev, ok := ev.(*kinds.ReplicatedBallotProof)
				assert.True(t, ok)
				assert.Equal(t, publickey.Location(), ev.BallotAN.AssessorLocation)
				assert.Equal(t, preballotAltitude, ev.Altitude())
			}
		}
	case <-time.After(20 * time.Second):
		t.Fatalf("REDACTED")
	}
}

//
//
//
//
//
func VerifyTreacherousDiscordantItemsUsingSegment(t *testing.T) {
	N := 4
	tracer := agreementTracer().Using("REDACTED", "REDACTED")

	ctx := t.Context()

	app := freshTokvalDepot
	css, sanitize := arbitraryAgreementNetwork(t, N, "REDACTED", freshSimulateMetronomeMethod(false), app)
	defer sanitize()

	//
	metronome := FreshDeadlineMetronome()
	metronome.AssignTracer(css[0].Tracer)
	css[0].AssignDeadlineMetronome(metronome)

	routers := make([]*p2p.Router, N)
	peer2peerTracer := tracer.Using("REDACTED", "REDACTED")
	for i := 0; i < N; i++ {
		routers[i] = p2p.CreateRouter(
			settings.P2P,
			i,
			func(i int, sw *p2p.Router) *p2p.Router {
				return sw
			})
		routers[i].AssignTracer(peer2peerTracer.Using("REDACTED", i))
	}

	ledgersSubscriptions := make([]kinds.Listening, N)
	engines := make([]p2p.Handler, N)
	for i := 0; i < N; i++ {

		//
		attestTxpool(css[i].transferObserver).ActivateTransAccessible()
		//
		if i == 0 {
			//
			//
			css[i].privateAssessor.(kinds.SimulatePRV).DeactivateVerifications()
			j := i
			css[i].resolveNomination = func(altitude int64, iteration int32) {
				treacherousResolveNominationMethod(ctx, t, altitude, iteration, css[j], routers[j])
			}
			//
			//
			css[i].performPreballot = func(altitude int64, iteration int32) {}
		}

		incidentChannel := css[i].incidentChannel
		incidentChannel.AssignTracer(tracer.Using("REDACTED", "REDACTED", "REDACTED", i))

		var err error
		ledgersSubscriptions[i], err = incidentChannel.Listen(context.Background(), verifyListener, kinds.IncidentInquireFreshLedger)
		require.NoError(t, err)

		connectionReader := FreshHandler(css[i], true) //
		connectionReader.AssignTracer(tracer.Using("REDACTED", i))
		connectionReader.AssignIncidentChannel(incidentChannel)

		var connectionRandidx p2p.Handler = connectionReader

		//
		if i == 0 {
			connectionRandidx = FreshTreacherousHandler(connectionReader)
		}

		engines[i] = connectionRandidx
		err = css[i].ledgerExecute.Depot().Persist(css[i].status) //
		require.NoError(t, err)
	}

	defer func() {
		for _, r := range engines {
			if rr, ok := r.(*TreacherousHandler); ok {
				err := rr.handler.Router.Halt()
				require.NoError(t, err)
			} else {
				err := r.(*Handler).Router.Halt()
				require.NoError(t, err)
			}
		}
	}()

	p2p.CreateAssociatedRouters(settings.P2P, N, func(i int, s *p2p.Router) *p2p.Router {
		//
		routers[i].AppendHandler("REDACTED", engines[i])
		return routers[i]
	}, func(sws []*p2p.Router, i, j int) {
		//
		if i != 0 {
			return
		}
		p2p.Connect2routers(sws, i, j)
	})

	//
	//
	for i := 1; i < N; i++ {
		cr := engines[i].(*Handler)
		cr.RouterTowardAgreement(cr.connectionSTR.ObtainStatus(), false)
	}

	//
	byzantineReader := engines[0].(*TreacherousHandler)
	s := byzantineReader.handler.connectionSTR.ObtainStatus()
	byzantineReader.handler.RouterTowardAgreement(s, false)

	//
	//
	//
	nodes := routers[0].Nodes().Duplicate()

	//
	ordinal0 := obtainRouterOrdinal(routers, nodes[0])

	//
	ordinal1 := obtainRouterOrdinal(routers, nodes[1])
	ordinal2 := obtainRouterOrdinal(routers, nodes[2])
	p2p.Connect2routers(routers, ordinal1, ordinal2)

	//
	<-ledgersSubscriptions[ordinal2].Out()

	t.Log("REDACTED")
	p2p.Connect2routers(routers, ordinal0, ordinal1)
	p2p.Connect2routers(routers, ordinal0, ordinal2)

	//
	//
	wg := new(sync.WaitGroup)
	for i := 1; i < N-1; i++ {
		wg.Add(1)
		go func(j int) {
			<-ledgersSubscriptions[j].Out()
			wg.Done()
		}(i)
	}

	complete := make(chan struct{})
	go func() {
		wg.Wait()
		close(complete)
	}()

	pulse := time.NewTicker(time.Second * 10)
	select {
	case <-complete:
	case <-pulse.C:
		for i, handler := range engines {
			t.Logf("REDACTED", i)
			t.Logf("REDACTED", handler)
		}
		t.Fatalf("REDACTED")
	}
}

//
//

func treacherousResolveNominationMethod(ctx context.Context, t *testing.T, altitude int64, iteration int32, cs *Status, sw *p2p.Router) {
	//
	//

	//
	ledger1, err := cs.generateNominationLedger(ctx)
	require.NoError(t, err)
	ledgerFragments1, err := ledger1.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)
	policyIteration, itemLedgerUUID := cs.SoundIteration, kinds.LedgerUUID{Digest: ledger1.Digest(), FragmentAssignHeading: ledgerFragments1.Heading()}
	item1 := kinds.FreshNomination(altitude, iteration, policyIteration, itemLedgerUUID)
	p1 := item1.TowardSchema()
	if err := cs.privateAssessor.AttestNomination(cs.status.SuccessionUUID, p1); err != nil {
		t.Error(err)
	}

	item1.Notation = p1.Notation

	//
	dispatchTransScope(t, cs, 0, 1)

	//
	ledger2, err := cs.generateNominationLedger(ctx)
	require.NoError(t, err)
	ledgerFragments2, err := ledger2.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)
	policyIteration, itemLedgerUUID = cs.SoundIteration, kinds.LedgerUUID{Digest: ledger2.Digest(), FragmentAssignHeading: ledgerFragments2.Heading()}
	item2 := kinds.FreshNomination(altitude, iteration, policyIteration, itemLedgerUUID)
	p2 := item2.TowardSchema()
	if err := cs.privateAssessor.AttestNomination(cs.status.SuccessionUUID, p2); err != nil {
		t.Error(err)
	}

	item2.Notation = p2.Notation

	ledger1hash := ledger1.Digest()
	ledger2hash := ledger2.Digest()

	//
	nodes := sw.Nodes().Duplicate()
	t.Logf("REDACTED", len(nodes))
	for i, node := range nodes {
		if i < len(nodes)/2 {
			go transmitNominationAlsoFragments(altitude, iteration, cs, node, item1, ledger1hash, ledgerFragments1)
		} else {
			go transmitNominationAlsoFragments(altitude, iteration, cs, node, item2, ledger2hash, ledgerFragments2)
		}
	}
}

func transmitNominationAlsoFragments(
	altitude int64,
	iteration int32,
	cs *Status,
	node p2p.Node,
	nomination *kinds.Nomination,
	ledgerDigest []byte,
	fragments *kinds.FragmentAssign,
) {
	//
	node.Transmit(p2p.Wrapper{
		ConduitUUID: DataConduit,
		Signal:   &strongmindcons.Nomination{Nomination: *nomination.TowardSchema()},
	})

	//
	for i := 0; i < int(fragments.Sum()); i++ {
		fragment := fragments.ObtainFragment(i)
		pp, err := fragment.TowardSchema()
		if err != nil {
			panic(err) //
		}
		node.Transmit(p2p.Wrapper{
			ConduitUUID: DataConduit,
			Signal: &strongmindcons.LedgerFragment{
				Altitude: altitude, //
				Iteration:  iteration,  //
				Fragment:   *pp,
			},
		})
	}

	//
	cs.mtx.Lock()
	preballot, _ := cs.attestBallot(commitchema.PreballotKind, ledgerDigest, fragments.Heading(), nil)
	preendorse, _ := cs.attestBallot(commitchema.PreendorseKind, ledgerDigest, fragments.Heading(), nil)
	cs.mtx.Unlock()
	node.Transmit(p2p.Wrapper{
		ConduitUUID: BallotConduit,
		Signal:   &strongmindcons.Ballot{Ballot: preballot.TowardSchema()},
	})
	node.Transmit(p2p.Wrapper{
		ConduitUUID: BallotConduit,
		Signal:   &strongmindcons.Ballot{Ballot: preendorse.TowardSchema()},
	})
}

//
//

type TreacherousHandler struct {
	facility.Facility
	handler *Handler
}

func FreshTreacherousHandler(connectionReader *Handler) *TreacherousHandler {
	return &TreacherousHandler{
		Facility: connectionReader,
		handler: connectionReader,
	}
}

func (br *TreacherousHandler) AssignRouter(s p2p.Router)              { br.handler.AssignRouter(s) }
func (br *TreacherousHandler) ObtainConduits() []*p2p.ConduitDefinition { return br.handler.ObtainConduits() }
func (br *TreacherousHandler) AppendNode(node p2p.Node) {
	if !br.handler.EqualsActive() {
		return
	}

	//
	nodeStatus := FreshNodeStatus(node).AssignTracer(br.handler.Tracer)
	node.Set(kinds.NodeStatusToken, nodeStatus)

	//
	//
	if !br.handler.AwaitChronize() {
		br.handler.transmitFreshIterationPhaseSignal(node)
	}
}

func (br *TreacherousHandler) DiscardNode(node p2p.Node, rationale any) {
	br.handler.DiscardNode(node, rationale)
}

func (br *TreacherousHandler) Accept(e p2p.Wrapper) {
	br.handler.Accept(e)
}

func (br *TreacherousHandler) InitializeNode(node p2p.Node) p2p.Node { return node }

//
func VerifyDeclineBulkyItems(t *testing.T) {
	ctx := t.Context()

	n := 2
	css, sanitize := arbitraryAgreementNetwork(t, n, "REDACTED", freshSimulateMetronomeMethod(false), freshTokvalDepot)
	defer sanitize()

	routers := make([]*p2p.Router, n)
	peer2peerTracer := agreementTracer().Using("REDACTED", "REDACTED")
	for i := 0; i < n; i++ {
		routers[i] = p2p.CreateRouter(
			settings.P2P,
			i,
			func(_ int, sw *p2p.Router) *p2p.Router {
				return sw
			})
		routers[i].AssignTracer(peer2peerTracer.Using("REDACTED", i))
	}

	engines := make([]p2p.Handler, n)
	for i := 0; i < n; i++ {
		connectionReader := FreshHandler(css[i], false)
		defer func() { require.NoError(t, connectionReader.Halt()) }()

		connectionReader.AssignTracer(agreementTracer().Using("REDACTED", i))
		engines[i] = connectionReader
	}

	p2p.CreateAssociatedRouters(settings.P2P, n, func(i int, _ *p2p.Router) *p2p.Router {
		routers[i].AppendHandler("REDACTED", engines[i])
		return routers[i]
	}, p2p.Connect2routers)

	nodes := routers[0].Nodes().Duplicate()
	objectiveNode := nodes[0]

	altitude := int64(1)
	iteration := int32(0)
	cs := css[0]

	ledger, err := cs.generateNominationLedger(ctx)
	require.NoError(t, err)

	ledgerFragments, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(t, err)

	//
	itemLedgerUUID := kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: ledgerFragments.Heading()}
	itemLedgerUUID.FragmentAssignHeading.Sum = 4294967295

	nomination := kinds.FreshNomination(altitude, iteration, -1, itemLedgerUUID)
	p := nomination.TowardSchema()
	if err := cs.privateAssessor.AttestNomination(cs.status.SuccessionUUID, p); err != nil {
		t.Error(err)
	}
	nomination.Notation = p.Notation

	triumph := objectiveNode.Transmit(p2p.Wrapper{
		ConduitUUID: DataConduit,
		Signal:   &strongmindcons.Nomination{Nomination: *nomination.TowardSchema()},
	})
	require.True(t, triumph)

	select {
	case e := <-css[1].nodeSignalStaging:
		//
		//
		if _, acceptedNomination := e.Msg.(*NominationSignal); acceptedNomination {
			assert.Fail(t, "REDACTED")
			return
		}
		//
		//
		assert.Fail(t, "REDACTED")
	case <-ctx.Done():
	case <-time.After(500 * time.Millisecond):
		//
		//
	}
}
