package chainchronize

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"

	chainchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/chainchronize"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	tpmocks "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool/simulations"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

var settings *cfg.Settings

func arbitraryOriginPaper(countAssessors int, arbitraryPotency bool, minimumPotency int64) (*kinds.OriginPaper, []kinds.PrivateAssessor) {
	assessors := make([]kinds.OriginAssessor, countAssessors)
	privateAssessors := make([]kinds.PrivateAssessor, countAssessors)
	for i := 0; i < countAssessors; i++ {
		val, privateItem := kinds.ArbitraryAssessor(arbitraryPotency, minimumPotency)
		assessors[i] = kinds.OriginAssessor{
			PublicToken: val.PublicToken,
			Potency:  val.BallotingPotency,
		}
		privateAssessors[i] = privateItem
	}
	sort.Sort(kinds.PrivateAssessorsViaLocation(privateAssessors))

	agreementParallel := kinds.FallbackAgreementSettings()
	agreementParallel.Iface.BallotAdditionsActivateAltitude = 1
	return &kinds.OriginPaper{
		OriginMoment:     committime.Now(),
		SuccessionUUID:         verify.FallbackVerifySuccessionUUID,
		Assessors:      assessors,
		AgreementSettings: agreementParallel,
	}, privateAssessors
}

type HandlerCouple struct {
	handler *TreacherousHandler
	app     delegate.PlatformLinks
}

type handlerChoices struct {
	taintedLedger          int64
	everyMissingAddnEndorseLedger int64
	certainBallotMultiples  bool
}

type handlerSelection func(*handlerChoices)

func usingTaintedLedger(altitude int64) handlerSelection {
	return func(o *handlerChoices) {
		o.taintedLedger = altitude
	}
}

func usingEveryMissingAddnEndorseLedger(altitude int64) handlerSelection {
	return func(o *handlerChoices) {
		o.everyMissingAddnEndorseLedger = altitude
	}
}

func usingCertainBallotMultiples() handlerSelection {
	return func(o *handlerChoices) {
		o.certainBallotMultiples = true
	}
}

func freshHandler(
	t *testing.T,
	tracer log.Tracer,
	producePaper *kinds.OriginPaper,
	privateItems []kinds.PrivateAssessor,
	maximumLedgerAltitude int64,
	choices ...handlerSelection,
) HandlerCouple {
	if len(privateItems) != 1 {
		panic("REDACTED")
	}

	var choices handlerChoices
	for _, opt := range choices {
		opt(&choices)
	}

	app := iface.FreshFoundationPlatform()
	cc := delegate.FreshRegionalCustomerOriginator(app)
	delegatePlatform := delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
	err := delegatePlatform.Initiate()
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	ledgerDatastore := dbm.FreshMemoryDatastore()
	statusDatastore := dbm.FreshMemoryDatastore()
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	ledgerDepot := depot.FreshLedgerDepot(ledgerDatastore)

	status, err := statusDepot.FetchOriginatingDatastoreEitherOriginPaper(producePaper)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	mp := &tpmocks.Txpool{}
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

	//
	//
	//
	ledgerChronize := true
	db := dbm.FreshMemoryDatastore()
	statusDepot = sm.FreshDepot(db, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	ledgerExecute := sm.FreshLedgerHandler(statusDepot, log.VerifyingTracer(), delegatePlatform.Agreement(),
		mp, sm.VoidProofHub{}, ledgerDepot)
	if err = statusDepot.Persist(status); err != nil {
		panic(err)
	}

	//
	observedAddnEndorse := &kinds.ExpandedEndorse{}

	publicToken, err := privateItems[0].ObtainPublicToken()
	if err != nil {
		panic(err)
	}
	location := publicToken.Location()
	idx, _ := status.Assessors.ObtainViaLocation(location)

	//
	for ledgerAltitude := int64(1); ledgerAltitude <= maximumLedgerAltitude; ledgerAltitude++ {
		ballotAdditionEqualsActivated := producePaper.AgreementSettings.Iface.BallotAdditionsActivated(ledgerAltitude)

		finalAddnEndorse := observedAddnEndorse.Replicate()

		thatLedger, err := status.CreateLedger(ledgerAltitude, nil, finalAddnEndorse.TowardEndorse(), nil, status.Assessors.Nominator.Location)
		require.NoError(t, err)

		thatFragments, err := thatLedger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
		require.NoError(t, err)
		ledgerUUID := kinds.LedgerUUID{Digest: thatLedger.Digest(), FragmentAssignHeading: thatFragments.Heading()}

		ballotMoment := time.Now()
		if choices.certainBallotMultiples {
			//
			//
			ballotMoment = producePaper.OriginMoment.Add(time.Duration(ledgerAltitude) * time.Second)
		}

		//
		ballot, err := kinds.CreateBallot(
			privateItems[0],
			thatLedger.SuccessionUUID,
			idx,
			thatLedger.Altitude,
			0,
			commitchema.PreendorseKind,
			ledgerUUID,
			ballotMoment,
		)
		if err != nil {
			panic(err)
		}
		observedAddnEndorse = &kinds.ExpandedEndorse{
			Altitude:             ballot.Altitude,
			Iteration:              ballot.Iteration,
			LedgerUUID:            ledgerUUID,
			ExpandedNotations: []kinds.ExpandedEndorseSignature{ballot.ExpandedEndorseSignature()},
		}

		status, err = ledgerExecute.ExecuteLedger(status, ledgerUUID, thatLedger)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}

		persistPreciseBallotAdditions := ledgerAltitude != choices.taintedLedger
		if persistPreciseBallotAdditions == ballotAdditionEqualsActivated {
			ledgerDepot.PersistLedgerUsingExpandedEndorse(thatLedger, thatFragments, observedAddnEndorse)
		} else {
			ledgerDepot.PersistLedger(thatLedger, thatFragments, observedAddnEndorse.TowardEndorse())
		}
	}

	r := FreshHandler(ledgerChronize, false, status.Duplicate(), ledgerExecute, ledgerDepot, nil, 0, NooperationTelemetry())
	bchainHandler := FreshTreacherousHandler(r)
	bchainHandler.taintedLedger = choices.taintedLedger
	bchainHandler.missingAddnEndorseLedger = choices.everyMissingAddnEndorseLedger
	bchainHandler.AssignTracer(tracer.Using("REDACTED", "REDACTED"))

	return HandlerCouple{bchainHandler, delegatePlatform}
}

func VerifyNegativeLedgerReply(t *testing.T) {
	settings = verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	producePaper, privateItems := arbitraryOriginPaper(1, false, 30)

	maximumLedgerAltitude := int64(65)

	handlerCouples := make([]HandlerCouple, 2)

	handlerCouples[0] = freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, maximumLedgerAltitude)
	handlerCouples[1] = freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, 0)

	p2p.CreateAssociatedRouters(settings.P2P, 2, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", handlerCouples[i].handler)
		return s
	}, p2p.Connect2routers)

	defer func() {
		for _, r := range handlerCouples {
			err := r.handler.Halt()
			require.NoError(t, err)
			err = r.app.Halt()
			require.NoError(t, err)
		}
	}()

	verifies := []struct {
		altitude   int64
		present bool
	}{
		{maximumLedgerAltitude + 2, false},
		{10, true},
		{1, true},
		{100, false},
	}

	for !handlerCouples[1].handler.hub.EqualsSeizedActive() {
		time.Sleep(10 * time.Millisecond)
	}

	assert.Equal(t, maximumLedgerAltitude, handlerCouples[0].handler.depot.Altitude())

	for _, tt := range verifies {
		ledger := handlerCouples[1].handler.depot.FetchLedger(tt.altitude)
		if tt.present {
			assert.True(t, ledger != nil)
		} else {
			assert.True(t, ledger == nil)
		}
	}
}

//
//
//
//
//
func VerifyFlawedLedgerHaltsNode(t *testing.T) {
	settings = verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	producePaper, privateItems := arbitraryOriginPaper(1, false, 30)

	maximumLedgerAltitude := int64(148)

	//
	anotherProducePaper, anotherPrivateItems := arbitraryOriginPaper(1, false, 30)
	anotherSuccession := freshHandler(t, log.VerifyingTracer(), anotherProducePaper, anotherPrivateItems, maximumLedgerAltitude)

	defer func() {
		err := anotherSuccession.handler.Halt()
		require.Error(t, err)
		err = anotherSuccession.app.Halt()
		require.NoError(t, err)
	}()

	handlerCouples := make([]HandlerCouple, 4)

	handlerCouples[0] = freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, maximumLedgerAltitude)
	handlerCouples[1] = freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, 0)
	handlerCouples[2] = freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, 0)
	handlerCouples[3] = freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, 0)

	routers := p2p.CreateAssociatedRouters(settings.P2P, 4, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", handlerCouples[i].handler)
		return s
	}, p2p.Connect2routers)

	defer func() {
		for _, r := range handlerCouples {
			err := r.handler.Halt()
			require.NoError(t, err)

			err = r.app.Halt()
			require.NoError(t, err)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		seizedActive := true
		for _, r := range handlerCouples {
			if !r.handler.hub.EqualsSeizedActive() {
				seizedActive = false
			}
		}
		if seizedActive {
			break
		}
	}

	//
	assert.Equal(t, 3, handlerCouples[1].handler.Router.Nodes().Extent())

	//
	//
	handlerCouples[3].handler.depot = anotherSuccession.handler.depot

	finalHandlerCouple := freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, 0)
	handlerCouples = append(handlerCouples, finalHandlerCouple)

	routers = append(routers, p2p.CreateAssociatedRouters(settings.P2P, 1, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", handlerCouples[len(handlerCouples)-1].handler)
		return s
	}, p2p.Connect2routers)...)

	for i := 0; i < len(handlerCouples)-1; i++ {
		p2p.Connect2routers(routers, i, len(handlerCouples)-1)
	}

	for !finalHandlerCouple.handler.hub.EqualsSeizedActive() && finalHandlerCouple.handler.Router.Nodes().Extent() != 0 {
		time.Sleep(1 * time.Second)
	}

	assert.True(t, finalHandlerCouple.handler.Router.Nodes().Extent() < len(handlerCouples)-1)
}

func VerifyInspectRouterTowardAgreementFinalAltitudeNull(t *testing.T) {
	const maximumLedgerAltitude = int64(45)

	settings = verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	producePaper, privateItems := arbitraryOriginPaper(1, false, 30)

	handlerCouples := make([]HandlerCouple, 1, 2)
	handlerCouples[0] = freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, 0)
	handlerCouples[0].handler.durationRouterTowardAgreement = 50 * time.Millisecond
	defer func() {
		for _, r := range handlerCouples {
			err := r.handler.Halt()
			require.NoError(t, err)
			err = r.app.Halt()
			require.NoError(t, err)
		}
	}()

	handlerCouples = append(handlerCouples, freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, maximumLedgerAltitude))

	var routers []*p2p.Router
	for _, r := range handlerCouples {
		routers = append(routers, p2p.CreateAssociatedRouters(settings.P2P, 1, func(i int, s *p2p.Router) *p2p.Router {
			s.AppendHandler("REDACTED", r.handler)
			return s
		}, p2p.Connect2routers)...)
	}

	time.Sleep(60 * time.Millisecond)

	//
	p2p.Connect2routers(routers, 0, 1)

	initiateMoment := time.Now()
	for {
		time.Sleep(20 * time.Millisecond)
		seizedActive := true
		for _, r := range handlerCouples {
			if !r.handler.hub.EqualsSeizedActive() {
				seizedActive = false
				break
			}
		}
		if seizedActive {
			break
		}
		if time.Since(initiateMoment) > 90*time.Second {
			msg := "REDACTED"
			for i, r := range handlerCouples {
				h, p, lr := r.handler.hub.ObtainCondition()
				c := r.handler.hub.EqualsSeizedActive()
				msg += fmt.Sprintf("REDACTED", i, h, p, lr, c)
			}
			require.Fail(t, msg)
		}
	}

	//
	//
	//
	const maximumVariance = 3
	for _, r := range handlerCouples {
		assert.GreaterOrEqual(t, r.handler.depot.Altitude(), maximumLedgerAltitude-maximumVariance)
	}
}

func ExpandedEndorseFabricAssistant(t *testing.T, maximumLedgerAltitude int64, activateBallotAdditionLocated int64, choices ...handlerSelection) {
	settings = verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginPath)
	producePaper, privateItems := arbitraryOriginPaper(1, false, 30)
	producePaper.AgreementSettings.Iface.BallotAdditionsActivateAltitude = activateBallotAdditionLocated

	handlerCouples := make([]HandlerCouple, 1, 2)
	handlerCouples[0] = freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, 0)
	handlerCouples[0].handler.durationRouterTowardAgreement = 50 * time.Millisecond
	defer func() {
		for _, r := range handlerCouples {
			err := r.handler.Halt()
			require.NoError(t, err)
			err = r.app.Halt()
			require.NoError(t, err)
		}
	}()

	handlerCouples = append(handlerCouples, freshHandler(t, log.VerifyingTracer(), producePaper, privateItems, maximumLedgerAltitude, choices...))

	var routers []*p2p.Router
	for _, r := range handlerCouples {
		routers = append(routers, p2p.CreateAssociatedRouters(settings.P2P, 1, func(i int, s *p2p.Router) *p2p.Router {
			s.AppendHandler("REDACTED", r.handler)
			return s
		}, p2p.Connect2routers)...)
	}

	time.Sleep(60 * time.Millisecond)

	//
	p2p.Connect2routers(routers, 0, 1)

	initiateMoment := time.Now()
	for {
		time.Sleep(20 * time.Millisecond)
		//
		require.False(t, handlerCouples[0].handler.hub.EqualsSeizedActive(), "REDACTED")
		//
		if time.Since(initiateMoment) > 5*time.Second {
			assert.Equal(t, 0, handlerCouples[0].handler.Router.Nodes().Extent(), "REDACTED")
			assert.Equal(t, 0, handlerCouples[1].handler.Router.Nodes().Extent(), "REDACTED")
			break
		}
	}
}

//
func VerifyInspectExpandedEndorseSurplus(t *testing.T) {
	const maximumLedgerAltitude = 10
	const activateBallotAddition = 5
	const unfitLedgerAltitude = 3

	ExpandedEndorseFabricAssistant(t, maximumLedgerAltitude, activateBallotAddition, usingTaintedLedger(unfitLedgerAltitude))
}

//
func VerifyInspectExpandedEndorseAbsent(t *testing.T) {
	const maximumLedgerAltitude = 10
	const activateBallotAddition = 5
	const unfitLedgerAltitude = 8

	ExpandedEndorseFabricAssistant(t, maximumLedgerAltitude, activateBallotAddition, usingTaintedLedger(unfitLedgerAltitude))
}

//
//
func VerifyInspectExpandedEndorseEveryMissing(t *testing.T) {
	const maximumLedgerAltitude = 10
	const activateBallotAddition = 1
	const everyMissingAltitude = 5

	ExpandedEndorseFabricAssistant(t, maximumLedgerAltitude, activateBallotAddition, usingEveryMissingAddnEndorseLedger(everyMissingAltitude))
}

//
//
//
//
type TreacherousHandler struct {
	*Handler
	taintedLedger       int64
	missingAddnEndorseLedger int64
}

func FreshTreacherousHandler(connectionReader *Handler) *TreacherousHandler {
	return &TreacherousHandler{
		Handler: connectionReader,
	}
}

//
//
//
func (bcR *TreacherousHandler) replyTowardNode(msg *chainchema.LedgerSolicit, src p2p.Node) (staged bool) {
	ledger := bcR.depot.FetchLedger(msg.Altitude)
	if ledger == nil {
		bcR.Tracer.Details("REDACTED", "REDACTED", src, "REDACTED", msg.Altitude)
		return src.AttemptTransmit(p2p.Wrapper{
			ConduitUUID: ChainchronizeConduit,
			Signal:   &chainchema.NegativeLedgerReply{Altitude: msg.Altitude},
		})
	}

	status, err := bcR.ledgerExecute.Depot().Fetch()
	if err != nil {
		bcR.Tracer.Failure("REDACTED", "REDACTED", err)
		return false
	}
	var addnEndorse *kinds.ExpandedEndorse
	ballotAdditionActivated := status.AgreementSettings.Iface.BallotAdditionsActivated(msg.Altitude)
	impreciseLedger := bcR.taintedLedger == msg.Altitude
	if ballotAdditionActivated && !impreciseLedger || !ballotAdditionActivated && impreciseLedger {
		addnEndorse = bcR.depot.FetchLedgerExpandedEndorse(msg.Altitude)
		if addnEndorse == nil {
			bcR.Tracer.Failure("REDACTED", "REDACTED", ledger)
			return false
		}
	}

	if bcR.missingAddnEndorseLedger == msg.Altitude && addnEndorse != nil {
		missingSignatures := make([]kinds.ExpandedEndorseSignature, len(addnEndorse.ExpandedNotations))
		for i := range missingSignatures {
			missingSignatures[i] = kinds.FreshExpandedEndorseSignatureMissing()
		}
		addnEndorse = &kinds.ExpandedEndorse{
			Altitude:             addnEndorse.Altitude,
			Iteration:              addnEndorse.Iteration,
			LedgerUUID:            addnEndorse.LedgerUUID,
			ExpandedNotations: missingSignatures,
		}
	}

	bl, err := ledger.TowardSchema()
	if err != nil {
		bcR.Tracer.Failure("REDACTED", "REDACTED", err)
		return false
	}

	return src.AttemptTransmit(p2p.Wrapper{
		ConduitUUID: ChainchronizeConduit,
		Signal: &chainchema.LedgerReply{
			Ledger:     bl,
			AddnEndorse: addnEndorse.TowardSchema(),
		},
	})
}

//
//
func (bcR *TreacherousHandler) Accept(e p2p.Wrapper) {
	if err := CertifySignal(e.Signal); err != nil {
		bcR.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", e.Signal, "REDACTED", err)
		bcR.Router.HaltNodeForeachFailure(e.Src, err)
		return
	}

	bcR.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.ConduitUUID, "REDACTED", e.Signal)

	switch msg := e.Signal.(type) {
	case *chainchema.LedgerSolicit:
		bcR.replyTowardNode(msg, e.Src)
	case *chainchema.LedgerReply:
		bi, err := kinds.LedgerOriginatingSchema(msg.Ledger)
		if err != nil {
			bcR.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", e.Signal, "REDACTED", err)
			bcR.Router.HaltNodeForeachFailure(e.Src, err)
			return
		}
		var addnEndorse *kinds.ExpandedEndorse
		if msg.AddnEndorse != nil {
			var err error
			addnEndorse, err = kinds.ExpandedEndorseOriginatingSchema(msg.AddnEndorse)
			if err != nil {
				bcR.Tracer.Failure("REDACTED",
					"REDACTED", e.Src,
					"REDACTED", err)
				bcR.Router.HaltNodeForeachFailure(e.Src, err)
				return
			}
		}

		if err := bcR.hub.AppendLedger(e.Src.ID(), bi, addnEndorse, msg.Ledger.Extent()); err != nil {
			bcR.Tracer.Failure("REDACTED", "REDACTED", e.Src, "REDACTED", err)
		}
	case *chainchema.ConditionSolicit:
		//
		e.Src.AttemptTransmit(p2p.Wrapper{
			ConduitUUID: ChainchronizeConduit,
			Signal: &chainchema.ConditionReply{
				Altitude: bcR.depot.Altitude(),
				Foundation:   bcR.depot.Foundation(),
			},
		})
	case *chainchema.ConditionReply:
		//
		bcR.hub.AssignNodeScope(e.Src.ID(), msg.Foundation, msg.Altitude)
	case *chainchema.NegativeLedgerReply:
		bcR.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", msg.Altitude)
		bcR.hub.ReiterateSolicitOriginating(msg.Altitude, e.Src.ID())
	default:
		bcR.Tracer.Failure(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
	}
}
