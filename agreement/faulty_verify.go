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

	dbm "github.com/valkyrieworks/-db"

	abciend "github.com/valkyrieworks/iface/customer"
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/proof"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/gateway"

	"github.com/valkyrieworks/p2p"
	cometconnect "github.com/valkyrieworks/schema/consensuscore/agreement"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
)

//
//

//
func VerifyFaultyPreballotAmbiguity(t *testing.T) {
	ctx := t.Context()

	const nRatifiers = 4
	const faultyMember = 0
	const preballotLevel = int64(2)
	verifyLabel := "REDACTED"
	timerFunction := newEmulateTimerFunction(true)
	applicationFunction := newObjectDepot

	generatePaper, privateValues := randomOriginPaper(nRatifiers, false, 30, nil)
	css := make([]*Status, nRatifiers)

	for i := 0; i < nRatifiers; i++ {
		tracer := agreementTracer().With("REDACTED", "REDACTED", "REDACTED", i)
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

		ledgerStore := dbm.NewMemoryStore()
		ledgerDepot := depot.NewLedgerDepot(ledgerStore)

		mtx := new(engineconnect.Lock)
		//
		gatewayApplicationLinkConnect := gateway.NewApplicationLinkAgreement(abciend.NewNativeCustomer(mtx, app), gateway.NoopStats())
		gatewayApplicationLinkMemory := gateway.NewApplicationLinkTxpool(abciend.NewNativeCustomer(mtx, app), gateway.NoopStats())

		//
		txpool := txpool.NewCCatalogTxpool(settings.Txpool,
			gatewayApplicationLinkMemory,
			status.FinalLedgerLevel,
			txpool.WithPreInspect(sm.TransferPreInspect(status)),
			txpool.WithSubmitInspect(sm.TransferSubmitInspect(status)))

		if thisSettings.Agreement.WaitForTrans() {
			txpool.ActivateTransAccessible()
		}

		//
		proofStore := dbm.NewMemoryStore()
		eventpool, err := proof.NewDepository(proofStore, statusDepot, ledgerDepot)
		require.NoError(t, err)
		eventpool.AssignTracer(tracer.With("REDACTED", "REDACTED"))

		//
		ledgerExecute := sm.NewLedgerRunner(statusDepot, log.VerifyingTracer(), gatewayApplicationLinkConnect, txpool, eventpool, ledgerDepot)
		cs := NewStatus(thisSettings.Agreement, status, ledgerExecute, ledgerDepot, txpool, eventpool)
		cs.AssignTracer(cs.Tracer)
		//
		pv := privateValues[i]
		cs.CollectionPrivateRatifier(pv)

		eventBus := kinds.NewEventBus()
		eventBus.AssignTracer(log.VerifyingTracer().With("REDACTED", "REDACTED"))
		err = eventBus.Begin()
		require.NoError(t, err)
		cs.AssignEventBus(eventBus)

		cs.CollectionDeadlineTimer(timerFunction())
		cs.AssignTracer(tracer)

		css[i] = cs
	}

	//
	handlers := make([]*Handler, nRatifiers)
	ledgersEnrollments := make([]kinds.Enrollment, 0)
	eventBuses := make([]*kinds.EventBus, nRatifiers)
	for i := 0; i < nRatifiers; i++ {
		handlers[i] = NewHandler(css[i], true) //
		handlers[i].AssignTracer(css[i].Tracer)

		//
		eventBuses[i] = css[i].eventBus
		handlers[i].AssignEventBus(eventBuses[i])

		ledgersSubtract, err := eventBuses[i].Enrol(context.Background(), verifyEnrollee, kinds.EventInquireNewLedger, 100)
		require.NoError(t, err)
		ledgersEnrollments = append(ledgersEnrollments, ledgersSubtract)

		if css[i].status.FinalLedgerLevel == 0 { //
			err = css[i].ledgerExecute.Depot().Persist(css[i].status)
			require.NoError(t, err)
		}
	}
	//
	p2p.CreateLinkedRouters(settings.P2P, nRatifiers, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", handlers[i])
		s.AssignTracer(handlers[i].connectS.Tracer.With("REDACTED", "REDACTED"))
		return s
	}, p2p.Connect2routers)

	//
	bcs := css[faultyMember]

	//
	bcs.doPreballot = func(level int64, duration int32) {
		//
		if level == preballotLevel {
			bcs.Tracer.Details("REDACTED")
			preballot1, err := bcs.attestBallot(engineproto.PreballotKind, bcs.NominationLedger.Digest(), bcs.NominationLedgerSegments.Heading(), nil)
			require.NoError(t, err)
			preballot2, err := bcs.attestBallot(engineproto.PreballotKind, nil, kinds.SegmentAssignHeading{}, nil)
			require.NoError(t, err)
			nodeCatalog := handlers[faultyMember].Router.Nodes().Clone()
			bcs.Tracer.Details("REDACTED", "REDACTED", nodeCatalog)
			//
			for i, node := range nodeCatalog {
				if i < len(nodeCatalog)/2 {
					bcs.Tracer.Details("REDACTED", "REDACTED", preballot1, "REDACTED", node)
					node.Transmit(p2p.Packet{
						Signal:   &cometconnect.Ballot{Ballot: preballot1.ToSchema()},
						StreamUID: BallotStream,
					})
				} else {
					bcs.Tracer.Details("REDACTED", "REDACTED", preballot2, "REDACTED", node)
					node.Transmit(p2p.Packet{
						Signal:   &cometconnect.Ballot{Ballot: preballot2.ToSchema()},
						StreamUID: BallotStream,
					})
				}
			}
		} else {
			bcs.Tracer.Details("REDACTED")
			bcs.standardDoPreballot(level, duration)
		}
	}

	//
	//
	//
	idleRecommender := css[1]

	idleRecommender.determineNomination = func(level int64, duration int32) {
		idleRecommender.Tracer.Details("REDACTED")
		if idleRecommender.privateRatifier == nil {
			panic("REDACTED")
		}

		var extensionEndorse *kinds.ExpandedEndorse
		switch {
		case idleRecommender.Level == idleRecommender.status.PrimaryLevel:
			//
			//
			extensionEndorse = &kinds.ExpandedEndorse{}
		case idleRecommender.FinalEndorse.HasDualThirdsBulk():
			//
			veLevelArgument := kinds.IfaceOptions{BallotPluginsActivateLevel: level}
			extensionEndorse = idleRecommender.FinalEndorse.CreateExpandedEndorse(veLevelArgument)
		default: //
			idleRecommender.Tracer.Fault("REDACTED")
			return
		}

		//
		extensionEndorse.ExpandedEndorsements[len(extensionEndorse.ExpandedEndorsements)-1] = kinds.NewExpandedEndorseSignatureMissing()

		if idleRecommender.privateRatifierPublicKey == nil {
			//
			//
			idleRecommender.Tracer.Fault(fmt.Sprintf("REDACTED", ErrPublicKeyIsNotCollection))
			return
		}
		recommenderAddress := idleRecommender.privateRatifierPublicKey.Location()

		ledger, err := idleRecommender.ledgerExecute.InstantiateNominationLedger(
			ctx, idleRecommender.Level, idleRecommender.status, extensionEndorse, recommenderAddress)
		require.NoError(t, err)
		ledgerSegments, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
		require.NoError(t, err)

		//
		//
		if err := idleRecommender.wal.PurgeAndAlign(); err != nil {
			idleRecommender.Tracer.Fault("REDACTED")
		}

		//
		nominationLedgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: ledgerSegments.Heading()}
		nomination := kinds.NewNomination(level, duration, idleRecommender.SoundEpoch, nominationLedgerUID)
		p := nomination.ToSchema()
		if err := idleRecommender.privateRatifier.AttestNomination(idleRecommender.status.LedgerUID, p); err == nil {
			nomination.Autograph = p.Autograph

			//
			idleRecommender.transmitIntrinsicSignal(messageDetails{&NominationSignal{nomination}, "REDACTED"})
			for i := 0; i < int(ledgerSegments.Sum()); i++ {
				segment := ledgerSegments.FetchSegment(i)
				idleRecommender.transmitIntrinsicSignal(messageDetails{&LedgerSegmentSignal{idleRecommender.Level, idleRecommender.Cycle, segment}, "REDACTED"})
			}
			idleRecommender.Tracer.Details("REDACTED", "REDACTED", level, "REDACTED", duration, "REDACTED", nomination)
			idleRecommender.Tracer.Diagnose(fmt.Sprintf("REDACTED", ledger))
		} else if !idleRecommender.resimulateStyle {
			idleRecommender.Tracer.Fault("REDACTED", "REDACTED", level, "REDACTED", duration, "REDACTED", err)
		}
	}

	//
	for i := 0; i < nRatifiers; i++ {
		s := handlers[i].connectS.FetchStatus()
		handlers[i].RouterToAgreement(s, false)
	}
	defer haltAgreementNet(log.VerifyingTracer(), handlers, eventBuses)

	//
	//
	proofFromEachRatifier := make([]kinds.Proof, nRatifiers)

	wg := new(sync.WaitGroup)
	for i := 0; i < nRatifiers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for msg := range ledgersEnrollments[i].Out() {
				ledger := msg.Data().(kinds.EventDataNewLedger).Ledger
				if len(ledger.Proof.Proof) != 0 {
					proofFromEachRatifier[i] = ledger.Proof.Proof[0]
					return
				}
			}
		}(i)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	publickey, err := bcs.privateRatifier.FetchPublicKey()
	require.NoError(t, err)

	select {
	case <-done:
		for idx, ev := range proofFromEachRatifier {
			if assert.NotNil(t, ev, idx) {
				ev, ok := ev.(*kinds.ReplicatedBallotProof)
				assert.True(t, ok)
				assert.Equal(t, publickey.Location(), ev.BallotA.RatifierLocation)
				assert.Equal(t, preballotLevel, ev.Level())
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
func VerifyFaultyClashingNominationsWithSegment(t *testing.T) {
	N := 4
	tracer := agreementTracer().With("REDACTED", "REDACTED")

	ctx := t.Context()

	app := newObjectDepot
	css, sanitize := randomAgreementNet(t, N, "REDACTED", newEmulateTimerFunction(false), app)
	defer sanitize()

	//
	timer := NewDeadlineTimer()
	timer.AssignTracer(css[0].Tracer)
	css[0].CollectionDeadlineTimer(timer)

	routers := make([]*p2p.Router, N)
	p2pTracer := tracer.With("REDACTED", "REDACTED")
	for i := 0; i < N; i++ {
		routers[i] = p2p.CreateRouter(
			settings.P2P,
			i,
			func(i int, sw *p2p.Router) *p2p.Router {
				return sw
			})
		routers[i].AssignTracer(p2pTracer.With("REDACTED", i))
	}

	ledgersEnrollments := make([]kinds.Enrollment, N)
	handlers := make([]p2p.Handler, N)
	for i := 0; i < N; i++ {

		//
		affirmTxpool(css[i].transferAlerter).ActivateTransAccessible()
		//
		if i == 0 {
			//
			//
			css[i].privateRatifier.(kinds.EmulatePV).DeactivateValidations()
			j := i
			css[i].determineNomination = func(level int64, duration int32) {
				faultyDetermineNominationFunction(ctx, t, level, duration, css[j], routers[j])
			}
			//
			//
			css[i].doPreballot = func(level int64, duration int32) {}
		}

		eventBus := css[i].eventBus
		eventBus.AssignTracer(tracer.With("REDACTED", "REDACTED", "REDACTED", i))

		var err error
		ledgersEnrollments[i], err = eventBus.Enrol(context.Background(), verifyEnrollee, kinds.EventInquireNewLedger)
		require.NoError(t, err)

		connectReader := NewHandler(css[i], true) //
		connectReader.AssignTracer(tracer.With("REDACTED", i))
		connectReader.AssignEventBus(eventBus)

		var connectRI p2p.Handler = connectReader

		//
		if i == 0 {
			connectRI = NewFaultyHandler(connectReader)
		}

		handlers[i] = connectRI
		err = css[i].ledgerExecute.Depot().Persist(css[i].status) //
		require.NoError(t, err)
	}

	defer func() {
		for _, r := range handlers {
			if rr, ok := r.(*FaultyHandler); ok {
				err := rr.handler.Router.Halt()
				require.NoError(t, err)
			} else {
				err := r.(*Handler).Router.Halt()
				require.NoError(t, err)
			}
		}
	}()

	p2p.CreateLinkedRouters(settings.P2P, N, func(i int, s *p2p.Router) *p2p.Router {
		//
		routers[i].AppendHandler("REDACTED", handlers[i])
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
		cr := handlers[i].(*Handler)
		cr.RouterToAgreement(cr.connectS.FetchStatus(), false)
	}

	//
	byzReader := handlers[0].(*FaultyHandler)
	s := byzReader.handler.connectS.FetchStatus()
	byzReader.handler.RouterToAgreement(s, false)

	//
	//
	//
	nodes := routers[0].Nodes().Clone()

	//
	idx0 := fetchRouterOrdinal(routers, nodes[0])

	//
	idx1 := fetchRouterOrdinal(routers, nodes[1])
	idx2 := fetchRouterOrdinal(routers, nodes[2])
	p2p.Connect2routers(routers, idx1, idx2)

	//
	<-ledgersEnrollments[idx2].Out()

	t.Log("REDACTED")
	p2p.Connect2routers(routers, idx0, idx1)
	p2p.Connect2routers(routers, idx0, idx2)

	//
	//
	wg := new(sync.WaitGroup)
	for i := 1; i < N-1; i++ {
		wg.Add(1)
		go func(j int) {
			<-ledgersEnrollments[j].Out()
			wg.Done()
		}(i)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	pulse := time.NewTicker(time.Second * 10)
	select {
	case <-done:
	case <-pulse.C:
		for i, handler := range handlers {
			t.Logf("REDACTED", i)
			t.Logf("REDACTED", handler)
		}
		t.Fatalf("REDACTED")
	}
}

//
//

func faultyDetermineNominationFunction(ctx context.Context, t *testing.T, level int64, duration int32, cs *Status, sw *p2p.Router) {
	//
	//

	//
	ledger1, err := cs.instantiateNominationLedger(ctx)
	require.NoError(t, err)
	ledgerSegments1, err := ledger1.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)
	polEpoch, nominationLedgerUID := cs.SoundEpoch, kinds.LedgerUID{Digest: ledger1.Digest(), SegmentAssignHeading: ledgerSegments1.Heading()}
	nomination1 := kinds.NewNomination(level, duration, polEpoch, nominationLedgerUID)
	p1 := nomination1.ToSchema()
	if err := cs.privateRatifier.AttestNomination(cs.status.LedgerUID, p1); err != nil {
		t.Error(err)
	}

	nomination1.Autograph = p1.Autograph

	//
	transferTransScope(t, cs, 0, 1)

	//
	ledger2, err := cs.instantiateNominationLedger(ctx)
	require.NoError(t, err)
	ledgerSegments2, err := ledger2.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)
	polEpoch, nominationLedgerUID = cs.SoundEpoch, kinds.LedgerUID{Digest: ledger2.Digest(), SegmentAssignHeading: ledgerSegments2.Heading()}
	nomination2 := kinds.NewNomination(level, duration, polEpoch, nominationLedgerUID)
	p2 := nomination2.ToSchema()
	if err := cs.privateRatifier.AttestNomination(cs.status.LedgerUID, p2); err != nil {
		t.Error(err)
	}

	nomination2.Autograph = p2.Autograph

	ledger1digest := ledger1.Digest()
	ledger2digest := ledger2.Digest()

	//
	nodes := sw.Nodes().Clone()
	t.Logf("REDACTED", len(nodes))
	for i, node := range nodes {
		if i < len(nodes)/2 {
			go transmitNominationAndSegments(level, duration, cs, node, nomination1, ledger1digest, ledgerSegments1)
		} else {
			go transmitNominationAndSegments(level, duration, cs, node, nomination2, ledger2digest, ledgerSegments2)
		}
	}
}

func transmitNominationAndSegments(
	level int64,
	duration int32,
	cs *Status,
	node p2p.Node,
	nomination *kinds.Nomination,
	ledgerDigest []byte,
	segments *kinds.SegmentCollection,
) {
	//
	node.Transmit(p2p.Packet{
		StreamUID: DataStream,
		Signal:   &cometconnect.Nomination{Nomination: *nomination.ToSchema()},
	})

	//
	for i := 0; i < int(segments.Sum()); i++ {
		segment := segments.FetchSegment(i)
		pp, err := segment.ToSchema()
		if err != nil {
			panic(err) //
		}
		node.Transmit(p2p.Packet{
			StreamUID: DataStream,
			Signal: &cometconnect.LedgerSegment{
				Level: level, //
				Cycle:  duration,  //
				Segment:   *pp,
			},
		})
	}

	//
	cs.mtx.Lock()
	preballot, _ := cs.attestBallot(engineproto.PreballotKind, ledgerDigest, segments.Heading(), nil)
	preendorse, _ := cs.attestBallot(engineproto.PreendorseKind, ledgerDigest, segments.Heading(), nil)
	cs.mtx.Unlock()
	node.Transmit(p2p.Packet{
		StreamUID: BallotStream,
		Signal:   &cometconnect.Ballot{Ballot: preballot.ToSchema()},
	})
	node.Transmit(p2p.Packet{
		StreamUID: BallotStream,
		Signal:   &cometconnect.Ballot{Ballot: preendorse.ToSchema()},
	})
}

//
//

type FaultyHandler struct {
	daemon.Daemon
	handler *Handler
}

func NewFaultyHandler(connectReader *Handler) *FaultyHandler {
	return &FaultyHandler{
		Daemon: connectReader,
		handler: connectReader,
	}
}

func (br *FaultyHandler) CollectionRouter(s p2p.Toggeler)              { br.handler.CollectionRouter(s) }
func (br *FaultyHandler) FetchStreams() []*p2p.StreamDefinition { return br.handler.FetchStreams() }
func (br *FaultyHandler) AppendNode(node p2p.Node) {
	if !br.handler.IsActive() {
		return
	}

	//
	nodeStatus := NewNodeStatus(node).AssignTracer(br.handler.Tracer)
	node.Set(kinds.NodeStatusKey, nodeStatus)

	//
	//
	if !br.handler.WaitAlign() {
		br.handler.transmitNewDurationPhaseSignal(node)
	}
}

func (br *FaultyHandler) DeleteNode(node p2p.Node, cause any) {
	br.handler.DeleteNode(node, cause)
}

func (br *FaultyHandler) Accept(e p2p.Packet) {
	br.handler.Accept(e)
}

func (br *FaultyHandler) InitNode(node p2p.Node) p2p.Node { return node }

//
func VerifyDeclineExcessiveNominations(t *testing.T) {
	ctx := t.Context()

	n := 2
	css, sanitize := randomAgreementNet(t, n, "REDACTED", newEmulateTimerFunction(false), newObjectDepot)
	defer sanitize()

	routers := make([]*p2p.Router, n)
	p2pTracer := agreementTracer().With("REDACTED", "REDACTED")
	for i := 0; i < n; i++ {
		routers[i] = p2p.CreateRouter(
			settings.P2P,
			i,
			func(_ int, sw *p2p.Router) *p2p.Router {
				return sw
			})
		routers[i].AssignTracer(p2pTracer.With("REDACTED", i))
	}

	handlers := make([]p2p.Handler, n)
	for i := 0; i < n; i++ {
		connectReader := NewHandler(css[i], false)
		defer func() { require.NoError(t, connectReader.Halt()) }()

		connectReader.AssignTracer(agreementTracer().With("REDACTED", i))
		handlers[i] = connectReader
	}

	p2p.CreateLinkedRouters(settings.P2P, n, func(i int, _ *p2p.Router) *p2p.Router {
		routers[i].AppendHandler("REDACTED", handlers[i])
		return routers[i]
	}, p2p.Connect2routers)

	nodes := routers[0].Nodes().Clone()
	objectiveNode := nodes[0]

	level := int64(1)
	duration := int32(0)
	cs := css[0]

	ledger, err := cs.instantiateNominationLedger(ctx)
	require.NoError(t, err)

	ledgerSegments, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	require.NoError(t, err)

	//
	nominationLedgerUID := kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: ledgerSegments.Heading()}
	nominationLedgerUID.SegmentAssignHeading.Sum = 4294967295

	nomination := kinds.NewNomination(level, duration, -1, nominationLedgerUID)
	p := nomination.ToSchema()
	if err := cs.privateRatifier.AttestNomination(cs.status.LedgerUID, p); err != nil {
		t.Error(err)
	}
	nomination.Autograph = p.Autograph

	success := objectiveNode.Transmit(p2p.Packet{
		StreamUID: DataStream,
		Signal:   &cometconnect.Nomination{Nomination: *nomination.ToSchema()},
	})
	require.True(t, success)

	select {
	case e := <-css[1].nodeMessageBuffer:
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
