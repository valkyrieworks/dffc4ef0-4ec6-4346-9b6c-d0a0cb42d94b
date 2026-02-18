package chainconnect

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"

	chainproto "github.com/valkyrieworks/schema/consensuscore/chainconnect"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"
	txpoolsims "github.com/valkyrieworks/txpool/simulations"
	"github.com/valkyrieworks/p2p"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

var settings *cfg.Settings

func randomOriginPaper(countRatifiers int, randomEnergy bool, minimumEnergy int64) (*kinds.OriginPaper, []kinds.PrivateRatifier) {
	ratifiers := make([]kinds.OriginRatifier, countRatifiers)
	privateRatifiers := make([]kinds.PrivateRatifier, countRatifiers)
	for i := 0; i < countRatifiers; i++ {
		val, privateValue := kinds.RandomRatifier(randomEnergy, minimumEnergy)
		ratifiers[i] = kinds.OriginRatifier{
			PublicKey: val.PublicKey,
			Energy:  val.PollingEnergy,
		}
		privateRatifiers[i] = privateValue
	}
	sort.Sort(kinds.PrivateRatifiersByLocation(privateRatifiers))

	constParallel := kinds.StandardAgreementOptions()
	constParallel.Iface.BallotPluginsActivateLevel = 1
	return &kinds.OriginPaper{
		OriginMoment:     engineclock.Now(),
		LedgerUID:         verify.StandardVerifyLedgerUID,
		Ratifiers:      ratifiers,
		AgreementOptions: constParallel,
	}, privateRatifiers
}

type HandlerCouple struct {
	handler *FaultyHandler
	app     gateway.ApplicationLinks
}

func newHandler(
	t *testing.T,
	tracer log.Tracer,
	generatePaper *kinds.OriginPaper,
	privateValues []kinds.PrivateRatifier,
	maximumLedgerLevel int64,
	invalidData ...int64,
) HandlerCouple {
	if len(privateValues) != 1 {
		panic("REDACTED")
	}
	var invalidLedger int64 = 0
	if len(invalidData) > 0 {
		invalidLedger = invalidData[0]
	}

	app := iface.NewRootSoftware()
	cc := gateway.NewNativeCustomerOriginator(app)
	gatewayApplication := gateway.NewApplicationLinks(cc, gateway.NoopStats())
	err := gatewayApplication.Begin()
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

	ledgerStore := dbm.NewMemoryStore()
	statusStore := dbm.NewMemoryStore()
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	ledgerDepot := depot.NewLedgerDepot(ledgerStore)

	status, err := statusDepot.ImportFromStoreOrOriginPaper(generatePaper)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}

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

	//
	//
	//
	ledgerAlign := true
	db := dbm.NewMemoryStore()
	statusDepot = sm.NewDepot(db, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	ledgerExecute := sm.NewLedgerRunner(statusDepot, log.VerifyingTracer(), gatewayApplication.Agreement(),
		mp, sm.EmptyProofDepository{}, ledgerDepot)
	if err = statusDepot.Persist(status); err != nil {
		panic(err)
	}

	//
	viewedExtensionEndorse := &kinds.ExpandedEndorse{}

	publicKey, err := privateValues[0].FetchPublicKey()
	if err != nil {
		panic(err)
	}
	address := publicKey.Location()
	idx, _ := status.Ratifiers.FetchByLocation(address)

	//
	for ledgerLevel := int64(1); ledgerLevel <= maximumLedgerLevel; ledgerLevel++ {
		ballotAdditionIsActivated := generatePaper.AgreementOptions.Iface.BallotPluginsActivated(ledgerLevel)

		finalExtensionEndorse := viewedExtensionEndorse.Replicate()

		thisLedger, err := status.CreateLedger(ledgerLevel, nil, finalExtensionEndorse.ToEndorse(), nil, status.Ratifiers.Recommender.Location)
		require.NoError(t, err)

		thisSegments, err := thisLedger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
		require.NoError(t, err)
		ledgerUID := kinds.LedgerUID{Digest: thisLedger.Digest(), SegmentAssignHeading: thisSegments.Heading()}

		//
		ballot, err := kinds.CreateBallot(
			privateValues[0],
			thisLedger.LedgerUID,
			idx,
			thisLedger.Level,
			0,
			engineproto.PreendorseKind,
			ledgerUID,
			time.Now(),
		)
		if err != nil {
			panic(err)
		}
		viewedExtensionEndorse = &kinds.ExpandedEndorse{
			Level:             ballot.Level,
			Cycle:              ballot.Cycle,
			LedgerUID:            ledgerUID,
			ExpandedEndorsements: []kinds.ExpandedEndorseSignature{ballot.ExpandedEndorseSignature()},
		}

		status, err = ledgerExecute.ExecuteLedger(status, ledgerUID, thisLedger)
		if err != nil {
			panic(fmt.Errorf("REDACTED", err))
		}

		persistAccurateBallotPlugins := ledgerLevel != invalidLedger
		if persistAccurateBallotPlugins == ballotAdditionIsActivated {
			ledgerDepot.PersistLedgerWithExpandedEndorse(thisLedger, thisSegments, viewedExtensionEndorse)
		} else {
			ledgerDepot.PersistLedger(thisLedger, thisSegments, viewedExtensionEndorse.ToEndorse())
		}
	}

	r := NewHandler(ledgerAlign, false, status.Clone(), ledgerExecute, ledgerDepot, nil, 0, NoopStats())
	bcodeHandler := NewFaultyHandler(invalidLedger, r)
	bcodeHandler.AssignTracer(tracer.With("REDACTED", "REDACTED"))

	return HandlerCouple{bcodeHandler, gatewayApplication}
}

func VerifyNoLedgerReply(t *testing.T) {
	settings = verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	generatePaper, privateValues := randomOriginPaper(1, false, 30)

	maximumLedgerLevel := int64(65)

	handlerCouples := make([]HandlerCouple, 2)

	handlerCouples[0] = newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, maximumLedgerLevel)
	handlerCouples[1] = newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, 0)

	p2p.CreateLinkedRouters(settings.P2P, 2, func(i int, s *p2p.Router) *p2p.Router {
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
		level   int64
		existing bool
	}{
		{maximumLedgerLevel + 2, false},
		{10, true},
		{1, true},
		{100, false},
	}

	for !handlerCouples[1].handler.depository.IsSeizedUp() {
		time.Sleep(10 * time.Millisecond)
	}

	assert.Equal(t, maximumLedgerLevel, handlerCouples[0].handler.depot.Level())

	for _, tt := range verifies {
		ledger := handlerCouples[1].handler.depot.ImportLedger(tt.level)
		if tt.existing {
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
	defer os.RemoveAll(settings.OriginFolder)
	generatePaper, privateValues := randomOriginPaper(1, false, 30)

	maximumLedgerLevel := int64(148)

	//
	anotherGeneratePaper, anotherPrivateValues := randomOriginPaper(1, false, 30)
	anotherLedger := newHandler(t, log.VerifyingTracer(), anotherGeneratePaper, anotherPrivateValues, maximumLedgerLevel)

	defer func() {
		err := anotherLedger.handler.Halt()
		require.Error(t, err)
		err = anotherLedger.app.Halt()
		require.NoError(t, err)
	}()

	handlerCouples := make([]HandlerCouple, 4)

	handlerCouples[0] = newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, maximumLedgerLevel)
	handlerCouples[1] = newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, 0)
	handlerCouples[2] = newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, 0)
	handlerCouples[3] = newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, 0)

	routers := p2p.CreateLinkedRouters(settings.P2P, 4, func(i int, s *p2p.Router) *p2p.Router {
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
		seizedUp := true
		for _, r := range handlerCouples {
			if !r.handler.depository.IsSeizedUp() {
				seizedUp = false
			}
		}
		if seizedUp {
			break
		}
	}

	//
	assert.Equal(t, 3, handlerCouples[1].handler.Router.Nodes().Volume())

	//
	//
	handlerCouples[3].handler.depot = anotherLedger.handler.depot

	finalHandlerCouple := newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, 0)
	handlerCouples = append(handlerCouples, finalHandlerCouple)

	routers = append(routers, p2p.CreateLinkedRouters(settings.P2P, 1, func(i int, s *p2p.Router) *p2p.Router {
		s.AppendHandler("REDACTED", handlerCouples[len(handlerCouples)-1].handler)
		return s
	}, p2p.Connect2routers)...)

	for i := 0; i < len(handlerCouples)-1; i++ {
		p2p.Connect2routers(routers, i, len(handlerCouples)-1)
	}

	for !finalHandlerCouple.handler.depository.IsSeizedUp() && finalHandlerCouple.handler.Router.Nodes().Volume() != 0 {
		time.Sleep(1 * time.Second)
	}

	assert.True(t, finalHandlerCouple.handler.Router.Nodes().Volume() < len(handlerCouples)-1)
}

func VerifyInspectRouterToAgreementFinalLevelNil(t *testing.T) {
	const maximumLedgerLevel = int64(45)

	settings = verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	generatePaper, privateValues := randomOriginPaper(1, false, 30)

	handlerCouples := make([]HandlerCouple, 1, 2)
	handlerCouples[0] = newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, 0)
	handlerCouples[0].handler.cadenceRouterToAgreement = 50 * time.Millisecond
	defer func() {
		for _, r := range handlerCouples {
			err := r.handler.Halt()
			require.NoError(t, err)
			err = r.app.Halt()
			require.NoError(t, err)
		}
	}()

	handlerCouples = append(handlerCouples, newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, maximumLedgerLevel))

	var routers []*p2p.Router
	for _, r := range handlerCouples {
		routers = append(routers, p2p.CreateLinkedRouters(settings.P2P, 1, func(i int, s *p2p.Router) *p2p.Router {
			s.AppendHandler("REDACTED", r.handler)
			return s
		}, p2p.Connect2routers)...)
	}

	time.Sleep(60 * time.Millisecond)

	//
	p2p.Connect2routers(routers, 0, 1)

	beginMoment := time.Now()
	for {
		time.Sleep(20 * time.Millisecond)
		seizedUp := true
		for _, r := range handlerCouples {
			if !r.handler.depository.IsSeizedUp() {
				seizedUp = false
				break
			}
		}
		if seizedUp {
			break
		}
		if time.Since(beginMoment) > 90*time.Second {
			msg := "REDACTED"
			for i, r := range handlerCouples {
				h, p, lr := r.handler.depository.FetchStatus()
				c := r.handler.depository.IsSeizedUp()
				msg += fmt.Sprintf("REDACTED", i, h, p, lr, c)
			}
			require.Fail(t, msg)
		}
	}

	//
	//
	//
	const maximumVary = 3
	for _, r := range handlerCouples {
		assert.GreaterOrEqual(t, r.handler.depot.Level(), maximumLedgerLevel-maximumVary)
	}
}

func ExpandedEndorseFabricAssister(t *testing.T, maximumLedgerLevel int64, activateBallotAdditionAt int64, corruptLedgerLevelAt int64) {
	settings = verify.RestoreVerifyOrigin("REDACTED")
	defer os.RemoveAll(settings.OriginFolder)
	generatePaper, privateValues := randomOriginPaper(1, false, 30)
	generatePaper.AgreementOptions.Iface.BallotPluginsActivateLevel = activateBallotAdditionAt

	handlerCouples := make([]HandlerCouple, 1, 2)
	handlerCouples[0] = newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, 0)
	handlerCouples[0].handler.cadenceRouterToAgreement = 50 * time.Millisecond
	defer func() {
		for _, r := range handlerCouples {
			err := r.handler.Halt()
			require.NoError(t, err)
			err = r.app.Halt()
			require.NoError(t, err)
		}
	}()

	handlerCouples = append(handlerCouples, newHandler(t, log.VerifyingTracer(), generatePaper, privateValues, maximumLedgerLevel, corruptLedgerLevelAt))

	var routers []*p2p.Router
	for _, r := range handlerCouples {
		routers = append(routers, p2p.CreateLinkedRouters(settings.P2P, 1, func(i int, s *p2p.Router) *p2p.Router {
			s.AppendHandler("REDACTED", r.handler)
			return s
		}, p2p.Connect2routers)...)
	}

	time.Sleep(60 * time.Millisecond)

	//
	p2p.Connect2routers(routers, 0, 1)

	beginMoment := time.Now()
	for {
		time.Sleep(20 * time.Millisecond)
		//
		require.False(t, handlerCouples[0].handler.depository.IsSeizedUp(), "REDACTED")
		//
		if time.Since(beginMoment) > 5*time.Second {
			assert.Equal(t, 0, handlerCouples[0].handler.Router.Nodes().Volume(), "REDACTED")
			assert.Equal(t, 0, handlerCouples[1].handler.Router.Nodes().Volume(), "REDACTED")
			break
		}
	}
}

//
func VerifyInspectExpandedEndorseSurplus(t *testing.T) {
	const maximumLedgerLevel = 10
	const activateBallotAddition = 5
	const corruptLedgerLevel = 3

	ExpandedEndorseFabricAssister(t, maximumLedgerLevel, activateBallotAddition, corruptLedgerLevel)
}

//
func VerifyInspectExpandedEndorseAbsent(t *testing.T) {
	const maximumLedgerLevel = 10
	const activateBallotAddition = 5
	const corruptLedgerLevel = 8

	ExpandedEndorseFabricAssister(t, maximumLedgerLevel, activateBallotAddition, corruptLedgerLevel)
}

//
//
//
//
type FaultyHandler struct {
	*Handler
	taintedLedger int64
}

func NewFaultyHandler(corruptLedger int64, connectReader *Handler) *FaultyHandler {
	return &FaultyHandler{
		Handler:        connectReader,
		taintedLedger: corruptLedger,
	}
}

//
//
//
func (bcR *FaultyHandler) answerToNode(msg *chainproto.LedgerQuery, src p2p.Node) (buffered bool) {
	ledger := bcR.depot.ImportLedger(msg.Level)
	if ledger == nil {
		bcR.Tracer.Details("REDACTED", "REDACTED", src, "REDACTED", msg.Level)
		return src.AttemptTransmit(p2p.Packet{
			StreamUID: ChainconnectStream,
			Signal:   &chainproto.NoLedgerReply{Level: msg.Level},
		})
	}

	status, err := bcR.ledgerExecute.Depot().Import()
	if err != nil {
		bcR.Tracer.Fault("REDACTED", "REDACTED", err)
		return false
	}
	var extensionEndorse *kinds.ExpandedEndorse
	ballotAdditionActivated := status.AgreementOptions.Iface.BallotPluginsActivated(msg.Level)
	invalidLedger := bcR.taintedLedger == msg.Level
	if ballotAdditionActivated && !invalidLedger || !ballotAdditionActivated && invalidLedger {
		extensionEndorse = bcR.depot.ImportLedgerExpandedEndorse(msg.Level)
		if extensionEndorse == nil {
			bcR.Tracer.Fault("REDACTED", "REDACTED", ledger)
			return false
		}
	}

	bl, err := ledger.ToSchema()
	if err != nil {
		bcR.Tracer.Fault("REDACTED", "REDACTED", err)
		return false
	}

	return src.AttemptTransmit(p2p.Packet{
		StreamUID: ChainconnectStream,
		Signal: &chainproto.LedgerReply{
			Ledger:     bl,
			ExtensionEndorse: extensionEndorse.ToSchema(),
		},
	})
}

//
//
func (bcR *FaultyHandler) Accept(e p2p.Packet) {
	if err := CertifyMessage(e.Signal); err != nil {
		bcR.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", e.Signal, "REDACTED", err)
		bcR.Router.HaltNodeForFault(e.Src, err)
		return
	}

	bcR.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", e.StreamUID, "REDACTED", e.Signal)

	switch msg := e.Signal.(type) {
	case *chainproto.LedgerQuery:
		bcR.answerToNode(msg, e.Src)
	case *chainproto.LedgerReply:
		bi, err := kinds.LedgerFromSchema(msg.Ledger)
		if err != nil {
			bcR.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", e.Signal, "REDACTED", err)
			bcR.Router.HaltNodeForFault(e.Src, err)
			return
		}
		var extensionEndorse *kinds.ExpandedEndorse
		if msg.ExtensionEndorse != nil {
			var err error
			extensionEndorse, err = kinds.ExpandedEndorseFromSchema(msg.ExtensionEndorse)
			if err != nil {
				bcR.Tracer.Fault("REDACTED",
					"REDACTED", e.Src,
					"REDACTED", err)
				bcR.Router.HaltNodeForFault(e.Src, err)
				return
			}
		}

		if err := bcR.depository.AppendLedger(e.Src.ID(), bi, extensionEndorse, msg.Ledger.Volume()); err != nil {
			bcR.Tracer.Fault("REDACTED", "REDACTED", e.Src, "REDACTED", err)
		}
	case *chainproto.StatusQuery:
		//
		e.Src.AttemptTransmit(p2p.Packet{
			StreamUID: ChainconnectStream,
			Signal: &chainproto.StatusReply{
				Level: bcR.depot.Level(),
				Root:   bcR.depot.Root(),
			},
		})
	case *chainproto.StatusReply:
		//
		bcR.depository.AssignNodeScope(e.Src.ID(), msg.Root, msg.Level)
	case *chainproto.NoLedgerReply:
		bcR.Tracer.Diagnose("REDACTED", "REDACTED", e.Src, "REDACTED", msg.Level)
		bcR.depository.ReworkQueryFrom(msg.Level, e.Src.ID())
	default:
		bcR.Tracer.Fault(fmt.Sprintf("REDACTED", reflect.TypeOf(msg)))
	}
}
