package member

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	_ "net/http/pprof" //

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/chainconnect"
	cfg "github.com/valkyrieworks/settings"
	cs "github.com/valkyrieworks/agreement"
	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/proof"
	"github.com/valkyrieworks/statusconnect"

	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/rapid"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/p2p/pex"
	"github.com/valkyrieworks/privatekey"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/status/ordinaler/ledger"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"

	_ "github.com/lib/pq" //
)

const readHeadingDeadline = 10 * time.Second

//
//
//
type OriginPaperSource func() (*kinds.OriginPaper, error)

//
//
func StandardOriginPaperSourceFunction(settings *cfg.Settings) OriginPaperSource {
	return func() (*kinds.OriginPaper, error) {
		return kinds.OriginPaperFromEntry(settings.OriginEntry())
	}
}

//
type Source func(*cfg.Settings, log.Tracer) (*Member, error)

//
//
//
func StandardNewMember(settings *cfg.Settings, tracer log.Tracer) (*Member, error) {
	memberKey, err := p2p.ImportOrGenerateMemberKey(settings.MemberKeyEntry())
	if err != nil {
		return nil, fmt.Errorf("REDACTED", settings.MemberKeyEntry(), err)
	}

	return NewMember(settings,
		privatekey.ImportOrGenerateEntryPV(settings.PrivateRatifierKeyEntry(), settings.PrivateRatifierStatusEntry()),
		memberKey,
		gateway.StandardCustomerOriginator(settings.GatewayApplication, settings.Iface, settings.StoreFolder()),
		StandardOriginPaperSourceFunction(settings),
		cfg.StandardStoreSource,
		StandardStatsSource(settings.Telemetry),
		tracer,
	)
}

//
type StatsSource func(ledgerUID string) (*cs.Stats, *p2p.Stats, *txpool.Stats, *sm.Stats, *gateway.Stats, *chainconnect.Stats, *statusconnect.Stats)

//
//
func StandardStatsSource(settings *cfg.TelemetrySettings) StatsSource {
	return func(ledgerUID string) (*cs.Stats, *p2p.Stats, *txpool.Stats, *sm.Stats, *gateway.Stats, *chainconnect.Stats, *statusconnect.Stats) {
		if settings.Monitorstats {
			return cs.MonitorstatsStats(settings.Scope, "REDACTED", ledgerUID),
				p2p.MonitorstatsStats(settings.Scope, "REDACTED", ledgerUID),
				txpool.MonitorstatsStats(settings.Scope, "REDACTED", ledgerUID),
				sm.MonitorstatsStats(settings.Scope, "REDACTED", ledgerUID),
				gateway.MonitorstatsStats(settings.Scope, "REDACTED", ledgerUID),
				chainconnect.MonitorstatsStats(settings.Scope, "REDACTED", ledgerUID),
				statusconnect.MonitorstatsStats(settings.Scope, "REDACTED", ledgerUID)
		}
		return cs.NoopStats(), p2p.NoopStats(), txpool.NoopStats(), sm.NoopStats(), gateway.NoopStats(), chainconnect.NoopStats(), statusconnect.NoopStats()
	}
}

type ledgerAlignHandler interface {
	Activate(sm.Status) error
}

//

func initDSz(settings *cfg.Settings, storeSource cfg.StoreSource) (ledgerDepot *depot.LedgerDepot, statusStore dbm.DB, err error) {
	var ledgerDepotStore dbm.DB
	ledgerDepotStore, err = storeSource(&cfg.StoreContext{ID: "REDACTED", Settings: settings})
	if err != nil {
		return
	}
	ledgerDepot = depot.NewLedgerDepot(ledgerDepotStore)

	statusStore, err = storeSource(&cfg.StoreContext{ID: "REDACTED", Settings: settings})
	if err != nil {
		return
	}

	return
}

func instantiateAndBeginGatewayApplicationLinks(customerOriginator gateway.CustomerOriginator, tracer log.Tracer, stats *gateway.Stats) (gateway.ApplicationLinks, error) {
	gatewayApplication := gateway.NewApplicationLinks(customerOriginator, stats)
	gatewayApplication.AssignTracer(tracer.With("REDACTED", "REDACTED"))
	if err := gatewayApplication.Begin(); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	return gatewayApplication, nil
}

func instantiateAndBeginEventBus(tracer log.Tracer) (*kinds.EventBus, error) {
	eventBus := kinds.NewEventBus()
	eventBus.AssignTracer(tracer.With("REDACTED", "REDACTED"))
	if err := eventBus.Begin(); err != nil {
		return nil, err
	}
	return eventBus, nil
}

func instantiateAndBeginOrdinalerDaemon(
	settings *cfg.Settings,
	ledgerUID string,
	storeSource cfg.StoreSource,
	eventBus *kinds.EventBus,
	tracer log.Tracer,
) (*transordinal.OrdinalerDaemon, transordinal.TransOrdinaler, ordinaler.LedgerOrdinaler, error) {
	var (
		transOrdinaler    transordinal.TransOrdinaler
		ledgerOrdinaler ordinaler.LedgerOrdinaler
	)

	transOrdinaler, ledgerOrdinaler, allOrdinalersDeactivated, err := ledger.OrdinalerFromSettingsWithDeactivatedOrdinalers(settings, storeSource, ledgerUID)
	if err != nil {
		return nil, nil, nil, err
	}
	if allOrdinalersDeactivated {
		return nil, transOrdinaler, ledgerOrdinaler, nil
	}

	transOrdinaler.AssignTracer(tracer.With("REDACTED", "REDACTED"))
	ledgerOrdinaler.AssignTracer(tracer.With("REDACTED", "REDACTED"))

	ordinalerDaemon := transordinal.NewOrdinalerDaemon(transOrdinaler, ledgerOrdinaler, eventBus, false)
	ordinalerDaemon.AssignTracer(tracer.With("REDACTED", "REDACTED"))
	if err := ordinalerDaemon.Begin(); err != nil {
		return nil, nil, nil, err
	}

	return ordinalerDaemon, transOrdinaler, ledgerOrdinaler, nil
}

func doGreeting(
	ctx context.Context,
	statusDepot sm.Depot,
	status sm.Status,
	ledgerDepot sm.LedgerDepot,
	generatePaper *kinds.OriginPaper,
	eventBus kinds.LedgerEventBroadcaster,
	gatewayApplication gateway.ApplicationLinks,
	agreementTracer log.Tracer,
) error {
	greeter := cs.NewGreeter(statusDepot, status, ledgerDepot, generatePaper)
	greeter.AssignTracer(agreementTracer)
	greeter.AssignEventBus(eventBus)
	if err := greeter.GreetingWithContext(ctx, gatewayApplication); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

func traceMemberLaunchDetails(status sm.Status, publicKey vault.PublicKey, tracer, agreementTracer log.Tracer) {
	//
	tracer.Details("REDACTED",
		"REDACTED", release.TMCoreSemaphoreRev,
		"REDACTED", release.IfaceSemaphoreRev,
		"REDACTED", release.LedgerProtocol,
		"REDACTED", release.P2PProtocol,
		"REDACTED", release.TMGitEndorseDigest,
	)

	//
	if status.Release.Agreement.Ledger != release.LedgerProtocol {
		tracer.Details("REDACTED",
			"REDACTED", release.LedgerProtocol,
			"REDACTED", status.Release.Agreement.Ledger,
		)
	}

	address := publicKey.Location()
	//
	if status.Ratifiers.HasLocation(address) {
		agreementTracer.Details("REDACTED", "REDACTED", address, "REDACTED", publicKey)
	} else {
		agreementTracer.Details("REDACTED", "REDACTED", address, "REDACTED", publicKey)
	}
}

func solelyRatifierIsWe(status sm.Status, nativeAddress vault.Location) bool {
	if status.Ratifiers.Volume() > 1 {
		return false
	}
	valueAddress, _ := status.Ratifiers.FetchByOrdinal(0)
	return bytes.Equal(nativeAddress, valueAddress)
}

//
func instantiateTxpoolAndTxpoolHandler(
	settings *cfg.Settings,
	gatewayApplication gateway.ApplicationLinks,
	status sm.Status,
	waitAlign bool,
	memplStats *txpool.Stats,
	tracer log.Tracer,
) (txpool.Txpool, waitAlignHandler) {
	tracer = tracer.With("REDACTED", "REDACTED")

	switch settings.Txpool.Kind {
	//
	case cfg.TxpoolKindOverflow, "REDACTED":
		mp := txpool.NewCCatalogTxpool(
			settings.Txpool,
			gatewayApplication.Txpool(),
			status.FinalLedgerLevel,
			txpool.WithStats(memplStats),
			txpool.WithPreInspect(sm.TransferPreInspect(status)),
			txpool.WithSubmitInspect(sm.TransferSubmitInspect(status)),
		)
		mp.AssignTracer(tracer)
		handler := txpool.NewHandler(
			settings.Txpool,
			mp,
			waitAlign,
		)
		if settings.Agreement.WaitForTrans() {
			mp.ActivateTransAccessible()
		}
		handler.AssignTracer(tracer)

		return mp, handler
	case cfg.TxpoolKindNoop:
		//
		//
		return &txpool.NoopTxpool{}, txpool.NewNoopTxpoolHandler()
	case cfg.TxpoolKindApplication:
		mp := txpool.NewApplicationTxpool(
			settings.Txpool,
			gatewayApplication.Txpool(),
			txpool.WithMorningTracer(tracer),
			txpool.WithMorningStats(memplStats),
		)
		handler := txpool.NewApplicationHandler(settings.Txpool, mp, waitAlign)
		handler.AssignTracer(tracer)

		return mp, handler
	default:
		panic(fmt.Sprintf("REDACTED", settings.Txpool.Kind))
	}
}

func instantiateProofHandler(settings *cfg.Settings, storeSource cfg.StoreSource,
	statusDepot sm.Depot, ledgerDepot *depot.LedgerDepot, tracer log.Tracer,
) (*proof.Handler, *proof.Depository, error) {
	proofStore, err := storeSource(&cfg.StoreContext{ID: "REDACTED", Settings: settings})
	if err != nil {
		return nil, nil, err
	}
	proofTracer := tracer.With("REDACTED", "REDACTED")
	proofDepository, err := proof.NewDepository(proofStore, statusDepot, ledgerDepot)
	if err != nil {
		return nil, nil, err
	}
	proofHandler := proof.NewHandler(proofDepository)
	proofHandler.AssignTracer(proofTracer)
	return proofHandler, proofDepository, nil
}

func instantiateChainconnectHandler(
	activated bool,
	settings *cfg.Settings,
	status sm.Status,
	ledgerExecute *sm.LedgerRunner,
	ledgerDepot *depot.LedgerDepot,
	nativeAddress vault.Location,
	inactiveStatusAlignLevel int64,
	tracer log.Tracer,
	stats *chainconnect.Stats,
) (p2p.Handler, error) {
	release := settings.LedgerAlign.Release
	if release != "REDACTED" {
		return nil, fmt.Errorf("REDACTED", release)
	}

	bcodeHandler := chainconnect.NewHandler(
		activated,
		settings.LedgerAlign.ReplicaStyle,
		status.Clone(),
		ledgerExecute,
		ledgerDepot,
		nativeAddress,
		inactiveStatusAlignLevel,
		stats,
	)

	bcodeHandler.AssignTracer(tracer.With("REDACTED", "REDACTED"))

	return bcodeHandler, nil
}

func instantiateAgreementHandler(settings *cfg.Settings,
	status sm.Status,
	ledgerExecute *sm.LedgerRunner,
	ledgerDepot sm.LedgerDepot,
	txpool txpool.Txpool,
	proofDepository *proof.Depository,
	privateRatifier kinds.PrivateRatifier,
	csStats *cs.Stats,
	waitAlign bool,
	eventBus *kinds.EventBus,
	agreementTracer log.Tracer,
	inactiveStatusAlignLevel int64,
) (*cs.Handler, *cs.Status) {
	agreementStatus := cs.NewStatus(
		settings.Agreement,
		status.Clone(),
		ledgerExecute,
		ledgerDepot,
		txpool,
		proofDepository,
		cs.StatusStats(csStats),
		cs.InactiveStatusAlignLevel(inactiveStatusAlignLevel),
	)
	agreementStatus.AssignTracer(agreementTracer)
	if privateRatifier != nil {
		agreementStatus.CollectionPrivateRatifier(privateRatifier)
	}
	agreementHandler := cs.NewHandler(agreementStatus, waitAlign, cs.HandlerStats(csStats))
	agreementHandler.AssignTracer(agreementTracer)
	//
	//
	agreementHandler.AssignEventBus(eventBus)
	return agreementHandler, agreementStatus
}

func instantiateCometCarrierWithRouter(
	settings *cfg.Settings,
	memberDetails p2p.MemberDetails,
	memberKey *p2p.MemberKey,
	gatewayApplication gateway.ApplicationLinks,
	txpoolHandler p2p.Handler,
	bcodeHandler p2p.Handler,
	statusAlignHandler *statusconnect.Handler,
	agreementHandler *cs.Handler,
	proofHandler *proof.Handler,
	p2pStats *p2p.Stats,
	tracer log.Tracer,
) (p2p.Carrier, *p2p.Router) {
	carrier, nodeScreens := instantiateCometCarrier(settings, memberDetails, memberKey, gatewayApplication)

	sw := instantiateCometRouter(
		settings,
		carrier,
		p2pStats,
		nodeScreens,
		txpoolHandler,
		bcodeHandler,
		statusAlignHandler,
		agreementHandler,
		proofHandler,
		memberDetails,
		memberKey,
		tracer,
	)

	return carrier, sw
}

func instantiateCometCarrier(
	settings *cfg.Settings,
	memberDetails p2p.MemberDetails,
	memberKey *p2p.MemberKey,
	gatewayApplication gateway.ApplicationLinks,
) (
	*p2p.MulticastCarrier,
	[]p2p.NodeRefineFunction,
) {
	var (
		mLinkSettings = p2p.MLinkSettings(settings.P2P)
		carrier   = p2p.NewMultiplexCarrier(memberDetails, *memberKey, mLinkSettings)
		linkScreens = []p2p.LinkRefineFunction{}
		nodeScreens = []p2p.NodeRefineFunction{}
	)

	if !settings.P2P.PermitReplicatedIP {
		linkScreens = append(linkScreens, p2p.LinkReplicatedIPRefine())
	}

	//
	//
	if settings.RefineNodes {
		linkScreens = append(
			linkScreens,
			//
			func(_ p2p.LinkCollection, c net.Conn, _ []net.IP) error {
				res, err := gatewayApplication.Inquire().Inquire(context.TODO(), &iface.QueryInquire{
					Route: fmt.Sprintf("REDACTED", c.RemoteAddr().String()),
				})
				if err != nil {
					return err
				}
				if res.IsErr() {
					return fmt.Errorf("REDACTED", res)
				}

				return nil
			},
		)

		nodeScreens = append(
			nodeScreens,
			//
			func(_ p2p.IDXNodeCollection, p p2p.Node) error {
				res, err := gatewayApplication.Inquire().Inquire(context.TODO(), &iface.QueryInquire{
					Route: fmt.Sprintf("REDACTED", p.ID()),
				})
				if err != nil {
					return err
				}
				if res.IsErr() {
					return fmt.Errorf("REDACTED", res)
				}

				return nil
			},
		)
	}

	p2p.MulticastCarrierLinkScreens(linkScreens...)(carrier)

	//
	max := settings.P2P.MaximumCountIncomingNodes + len(divideAndShaveEmpty(settings.P2P.AbsoluteNodeIDXDatastore, "REDACTED", "REDACTED"))
	p2p.MulticastCarrierMaximumIncomingLinkages(max)(carrier)

	return carrier, nodeScreens
}

func instantiateCometRouter(
	settings *cfg.Settings,
	carrier p2p.Carrier,
	p2pStats *p2p.Stats,
	nodeScreens []p2p.NodeRefineFunction,
	txpoolHandler p2p.Handler,
	bcodeHandler p2p.Handler,
	statusAlignHandler *statusconnect.Handler,
	agreementHandler *cs.Handler,
	proofHandler *proof.Handler,
	memberDetails p2p.MemberDetails,
	memberKey *p2p.MemberKey,
	p2pTracer log.Tracer,
) *p2p.Router {
	sw := p2p.NewRouter(
		settings.P2P,
		carrier,
		p2p.WithStats(p2pStats),
		p2p.RouterNodeScreens(nodeScreens...),
	)
	sw.AssignTracer(p2pTracer)
	if settings.Txpool.Kind != cfg.TxpoolKindNoop {
		sw.AppendHandler("REDACTED", txpoolHandler)
	}
	sw.AppendHandler("REDACTED", bcodeHandler)
	sw.AppendHandler("REDACTED", agreementHandler)
	sw.AppendHandler("REDACTED", proofHandler)
	sw.AppendHandler("REDACTED", statusAlignHandler)

	sw.CollectionMemberDetails(memberDetails)
	sw.CollectionMemberKey(memberKey)

	p2pTracer.Details("REDACTED", "REDACTED", memberKey.ID(), "REDACTED", settings.MemberKeyEntry())
	return sw
}

func instantiateAddressRegistryAndCollectionOnRouter(
	settings *cfg.Settings,
	sw *p2p.Router,
	p2pTracer log.Tracer,
	memberKey *p2p.MemberKey,
) (pex.AddressLedger, error) {
	addressRegistry := pex.NewAddressRegistry(settings.P2P.AddressLedgerEntry(), settings.P2P.AddressLedgerPrecise)
	addressRegistry.AssignTracer(p2pTracer.With("REDACTED", settings.P2P.AddressLedgerEntry()))

	//
	if settings.P2P.OutsideLocation != "REDACTED" {
		address, err := p2p.NewNetLocationString(p2p.UIDLocationString(memberKey.ID(), settings.P2P.OutsideLocation))
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		addressRegistry.AppendOurLocation(address)
	}
	if settings.P2P.AcceptLocation != "REDACTED" {
		address, err := p2p.NewNetLocationString(p2p.UIDLocationString(memberKey.ID(), settings.P2P.AcceptLocation))
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		addressRegistry.AppendOurLocation(address)
	}

	sw.CollectionAddressRegistry(addressRegistry)

	return addressRegistry, nil
}

func instantiatePEXHandlerAndAppendToRouter(
	addressRegistry pex.AddressLedger,
	settings *cfg.Settings,
	sw p2p.Toggeler,
	tracer log.Tracer,
) *pex.Handler {
	cfg := &pex.HandlerSettings{
		Origins:    divideAndShaveEmpty(settings.P2P.Origins, "REDACTED", "REDACTED"),
		OriginStyle: settings.P2P.OriginStyle,
		//
		//
		//
		//
		//
		SourceDetachWaitDuration:     28 * time.Hour,
		DurableNodesMaximumCallDuration: settings.P2P.DurableNodesMaximumCallDuration,
	}

	//
	pexHandler := pex.NewHandler(addressRegistry, cfg)
	pexHandler.AssignTracer(tracer.With("REDACTED", "REDACTED"))

	sw.AppendHandler("REDACTED", pexHandler)

	return pexHandler
}

//
func beginStatusAlign(
	ssR *statusconnect.Handler,
	bcR ledgerAlignHandler,
	statusSource statusconnect.StatusSource,
	settings *cfg.StatusAlignSettings,
	statusDepot sm.Depot,
	ledgerDepot *depot.LedgerDepot,
	status sm.Status,
) error {
	ssR.Tracer.Details("REDACTED")

	if statusSource == nil {
		var err error
		ctx, revoke := context.WithTimeout(context.Background(), 10*time.Second)
		defer revoke()
		statusSource, err = statusconnect.NewRapidCustomerStatusSource(
			ctx,
			status.LedgerUID, status.Release, status.PrimaryLevel,
			settings.RPCHosts, rapid.ValidateOptions{
				Duration: settings.RelianceDuration,
				Level: settings.RelianceLevel,
				Digest:   settings.RelianceDigestOctets(),
			}, ssR.Tracer.With("REDACTED", "REDACTED"))
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}

	go func() {
		status, endorse, err := ssR.Align(statusSource, settings.DetectionTime)
		if err != nil {
			ssR.Tracer.Fault("REDACTED", "REDACTED", err)
			return
		}
		err = statusDepot.Onboard(status)
		if err != nil {
			ssR.Tracer.Fault("REDACTED", "REDACTED", err)
			return
		}
		err = ledgerDepot.PersistViewedEndorse(status.FinalLedgerLevel, endorse)
		if err != nil {
			ssR.Tracer.Fault("REDACTED", "REDACTED", err)
			return
		}

		if err := bcR.Activate(status); err != nil {
			ssR.Tracer.Fault("REDACTED", "REDACTED", err)
			return
		}
	}()
	return nil
}

//

var originPaperKey = []byte("REDACTED")

//
//
//
func ImportStatusFromStoreOrOriginPaperSource(
	statusStore dbm.DB,
	originPaperSource OriginPaperSource,
) (sm.Status, *kinds.OriginPaper, error) {
	//
	generatePaper, err := importOriginPaper(statusStore)
	if err != nil {
		generatePaper, err = originPaperSource()
		if err != nil {
			return sm.Status{}, nil, err
		}

		err = generatePaper.CertifyAndFinished()
		if err != nil {
			return sm.Status{}, nil, fmt.Errorf("REDACTED", err)
		}
		//
		//
		if err := persistOriginPaper(statusStore, generatePaper); err != nil {
			return sm.Status{}, nil, err
		}
	}
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	status, err := statusDepot.ImportFromStoreOrOriginPaper(generatePaper)
	if err != nil {
		return sm.Status{}, nil, err
	}
	return status, generatePaper, nil
}

//
func importOriginPaper(db dbm.DB) (*kinds.OriginPaper, error) {
	b, err := db.Get(originPaperKey)
	if err != nil {
		panic(err)
	}
	if len(b) == 0 {
		return nil, errors.New("REDACTED")
	}
	var generatePaper *kinds.OriginPaper
	err = cometjson.Unserialize(b, &generatePaper)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err, b))
	}
	return generatePaper, nil
}

//
func persistOriginPaper(db dbm.DB, generatePaper *kinds.OriginPaper) error {
	b, err := cometjson.Serialize(generatePaper)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return db.CollectionAlign(originPaperKey, b)
}

func instantiateAndBeginPrivateRatifierSocketCustomer(
	acceptAddress,
	ledgerUID string,
	tracer log.Tracer,
) (kinds.PrivateRatifier, error) {
	pve, err := privatekey.NewNotaryObserver(acceptAddress, tracer)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	pvsc, err := privatekey.NewNotaryCustomer(pve, ledgerUID)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	//
	_, err = pvsc.FetchPublicKey()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	const (
		attempts = 50 //
		deadline = 100 * time.Millisecond
	)
	pvscWithAttempts := privatekey.NewReprocessNotaryCustomer(pvsc, attempts, deadline)

	return pvscWithAttempts, nil
}

//
//
//
//
//
func divideAndShaveEmpty(s, sep, delimiters string) []string {
	if s == "REDACTED" {
		return []string{}
	}

	spl := strings.Split(s, sep)
	notEmptyStrings := make([]string, 0, len(spl))
	for i := 0; i < len(spl); i++ {
		member := strings.Trim(spl[i], delimiters)
		if member != "REDACTED" {
			notEmptyStrings = append(notEmptyStrings, member)
		}
	}
	return notEmptyStrings
}
