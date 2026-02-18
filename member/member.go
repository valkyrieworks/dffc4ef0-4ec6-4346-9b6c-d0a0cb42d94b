package member

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"

	bc "github.com/valkyrieworks/chainconnect"
	cfg "github.com/valkyrieworks/settings"
	cs "github.com/valkyrieworks/agreement"
	"github.com/valkyrieworks/proof"
	"github.com/valkyrieworks/rapid"

	"github.com/valkyrieworks/utils/log"
	cometbroadcast "github.com/valkyrieworks/utils/broadcast"
	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/netpeer"
	txpool "github.com/valkyrieworks/txpool"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/p2p/pex"
	"github.com/valkyrieworks/gateway"
	rpcbase "github.com/valkyrieworks/rpc/core"
	grpccore "github.com/valkyrieworks/rpc/grpc"
	rpchost "github.com/valkyrieworks/rpc/jsonrpc/host"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/status/transordinal/void"
	"github.com/valkyrieworks/statusconnect"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
	"github.com/valkyrieworks/release"

	_ "net/http/pprof" //
)

//
//
type Member struct {
	daemon.RootDaemon

	//
	settings        *cfg.Settings
	originPaper    *kinds.OriginPaper   //
	privateRatifier kinds.PrivateRatifier //

	//
	carrier   p2p.Carrier
	sw          p2p.Toggeler //
	memberDetails    p2p.MemberDetails
	memberKey     *p2p.MemberKey //
	isObserving bool

	//
	eventBus          *kinds.EventBus //
	statusDepot        sm.Depot
	ledgerDepot        *depot.LedgerDepot //
	bcodeHandler         p2p.Handler       //
	txpoolHandler    waitAlignHandler   //
	txpool           txpool.Txpool
	statusAlign         bool                    //
	statusAlignHandler  *statusconnect.Handler      //
	statusAlignSource statusconnect.StatusSource //
	statusAlignOrigin  sm.Status                //
	agreementStatus    *cs.Status               //
	agreementHandler  *cs.Handler             //
	proofDepository      *proof.Depository          //
	gatewayApplication          gateway.ApplicationLinks          //
	rpcObservers      []net.Listener          //
	transOrdinaler         transordinal.TransOrdinaler
	ledgerOrdinaler      ordinaler.LedgerOrdinaler
	ordinalerDaemon    *transordinal.OrdinalerDaemon
	monitorstatsSvc     *http.Server
	pprofSvc          *http.Server
}

type waitAlignHandler interface {
	p2p.Handler
	//
	WaitAlign() bool
}

//
type Setting func(*Member)

//
//
//
//
//
//
//
//
//
//
//
//
func BespokeHandlers(handlers map[string]p2p.Handler) Setting {
	return func(n *Member) {
		for label, handler := range handlers {
			if presentHandler, ok := n.sw.Handler(label); ok {
				n.sw.Log().Details("REDACTED",
					"REDACTED", label, "REDACTED", presentHandler, "REDACTED", handler)
				n.sw.DeleteHandler(label, presentHandler)
			}
			n.sw.AppendHandler(label, handler)

			//
			//
			//
			//
			ni, ok := n.memberDetails.(p2p.StandardMemberDetails)
			if !ok {
				n.Tracer.Fault("REDACTED")
				continue
			}

			mp, ok := n.carrier.(*p2p.MulticastCarrier)
			if !ok {
				n.Tracer.Fault("REDACTED")
				continue
			}

			for _, chanNote := range handler.FetchStreams() {
				if ni.HasConduit(chanNote.ID) {
					continue
				}

				ni.Streams = append(ni.Streams, chanNote.ID)
				mp.AppendConduit(chanNote.ID)
			}

			n.memberDetails = ni
		}
	}
}

//
//
//
func StatusSource(statusSource statusconnect.StatusSource) Setting {
	return func(n *Member) {
		n.statusAlignSource = statusSource
	}
}

//
//
//
//
//
func OnboardStatus(ctx context.Context, settings *cfg.Settings, storeSource cfg.StoreSource, level uint64, applicationDigest []byte) error {
	return OnboardStatusWithGenerateSource(ctx, settings, storeSource, StandardOriginPaperSourceFunction(settings), level, applicationDigest)
}

//
//
//
//
//
func OnboardStatusWithGenerateSource(ctx context.Context, settings *cfg.Settings, storeSource cfg.StoreSource, generateSource OriginPaperSource, level uint64, applicationDigest []byte) (err error) {
	tracer := log.NewTMTracer(log.NewAlignRecorder(os.Stdout))
	if ctx == nil {
		ctx = context.Background()
	}

	if settings == nil {
		tracer.Details("REDACTED")
		settings = cfg.StandardSettings()
	}

	if storeSource == nil {
		storeSource = cfg.StandardStoreSource
	}
	ledgerDepot, statusStore, err := initDSz(settings, storeSource)

	defer func() {
		if derr := ledgerDepot.End(); derr != nil {
			tracer.Fault("REDACTED", "REDACTED", derr)
			//
			err = derr
		}
	}()

	if err != nil {
		return err
	}

	if !ledgerDepot.IsEmpty() {
		return fmt.Errorf("REDACTED")
	}

	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: settings.Archival.DropIfaceReplies,
	})

	defer func() {
		if derr := statusDepot.End(); derr != nil {
			tracer.Fault("REDACTED", "REDACTED", derr)
			//
			err = derr
		}
	}()
	status, err := statusDepot.Import()
	if err != nil {
		return err
	}

	if !status.IsEmpty() {
		return fmt.Errorf("REDACTED")
	}

	generateStatus, _, err := ImportStatusFromStoreOrOriginPaperSource(statusStore, generateSource)
	if err != nil {
		return err
	}

	statusSource, err := statusconnect.NewRapidCustomerStatusSource(
		ctx,
		generateStatus.LedgerUID, generateStatus.Release, generateStatus.PrimaryLevel,
		settings.StatusAlign.RPCHosts, rapid.ValidateOptions{
			Duration: settings.StatusAlign.RelianceDuration,
			Level: settings.StatusAlign.RelianceLevel,
			Digest:   settings.StatusAlign.RelianceDigestOctets(),
		}, tracer.With("REDACTED", "REDACTED"))
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	status, err = statusSource.Status(ctx, level)
	if err != nil {
		return err
	}
	if applicationDigest == nil {
		tracer.Details("REDACTED")
	} else if !bytes.Equal(applicationDigest, status.ApplicationDigest) {
		if err := ledgerDepot.End(); err != nil {
			tracer.Fault("REDACTED", err)
		}
		if err := statusDepot.End(); err != nil {
			tracer.Fault("REDACTED", err)
		}
		return fmt.Errorf("REDACTED", status.ApplicationDigest, applicationDigest)

	}

	endorse, err := statusSource.Endorse(ctx, level)
	if err != nil {
		return err
	}

	if err = statusDepot.Onboard(status); err != nil {
		return err
	}

	err = ledgerDepot.PersistViewedEndorse(status.FinalLedgerLevel, endorse)
	if err != nil {
		return err
	}

	//
	//
	//
	//
	err = statusDepot.CollectionInactiveStatusAlignLevel(status.FinalLedgerLevel)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return err
}

//

//
func NewMember(
	settings *cfg.Settings,
	privateRatifier kinds.PrivateRatifier,
	memberKey *p2p.MemberKey,
	customerOriginator gateway.CustomerOriginator,
	originPaperSource OriginPaperSource,
	storeSource cfg.StoreSource,
	statsSource StatsSource,
	tracer log.Tracer,
	options ...Setting,
) (*Member, error) {
	return NewMemberWithContext(context.TODO(), settings, privateRatifier,
		memberKey, customerOriginator, originPaperSource, storeSource,
		statsSource, tracer, options...)
}

//
func NewMemberWithContext(
	ctx context.Context,
	settings *cfg.Settings,
	privateRatifier kinds.PrivateRatifier,
	memberKey *p2p.MemberKey,
	customerOriginator gateway.CustomerOriginator,
	originPaperSource OriginPaperSource,
	storeSource cfg.StoreSource,
	statsSource StatsSource,
	tracer log.Tracer,
	options ...Setting,
) (*Member, error) {
	ledgerDepot, statusStore, err := initDSz(settings, storeSource)
	if err != nil {
		return nil, err
	}

	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: settings.Archival.DropIfaceReplies,
	})

	status, generatePaper, err := ImportStatusFromStoreOrOriginPaperSource(statusStore, originPaperSource)
	if err != nil {
		return nil, err
	}

	csStats, p2pStats, memplStats, machineStats, ifaceStats, szStats, ssStats := statsSource(generatePaper.LedgerUID)

	//
	gatewayApplication, err := instantiateAndBeginGatewayApplicationLinks(customerOriginator, tracer, ifaceStats)
	if err != nil {
		return nil, err
	}

	//
	//
	//
	//
	eventBus, err := instantiateAndBeginEventBus(tracer)
	if err != nil {
		return nil, err
	}

	ordinalerDaemon, transOrdinaler, ledgerOrdinaler, err := instantiateAndBeginOrdinalerDaemon(settings,
		generatePaper.LedgerUID, storeSource, eventBus, tracer)
	if err != nil {
		return nil, err
	}

	//
	//
	if settings.PrivateRatifierAcceptAddress != "REDACTED" {
		//
		privateRatifier, err = instantiateAndBeginPrivateRatifierSocketCustomer(settings.PrivateRatifierAcceptAddress, generatePaper.LedgerUID, tracer)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	publicKey, err := privateRatifier.FetchPublicKey()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	nativeAddress := publicKey.Location()

	//
	statusAlign := settings.StatusAlign.Activate && !solelyRatifierIsWe(status, nativeAddress)
	if statusAlign && status.FinalLedgerLevel > 0 {
		tracer.Details("REDACTED")
		statusAlign = false
	}

	//
	//
	agreementTracer := tracer.With("REDACTED", "REDACTED")
	if !statusAlign {
		if err := doGreeting(ctx, statusDepot, status, ledgerDepot, generatePaper, eventBus, gatewayApplication, agreementTracer); err != nil {
			return nil, err
		}

		//
		//
		//
		status, err = statusDepot.Import()
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	traceMemberLaunchDetails(status, publicKey, tracer, agreementTracer)

	//
	//
	ledgerAlign := !solelyRatifierIsWe(status, nativeAddress)
	waitAlign := statusAlign || ledgerAlign

	if settings.LedgerAlign.ReplicaStyle {
		if status.Ratifiers.HasLocation(nativeAddress) {
			tracer.Fault("REDACTED")
			settings.LedgerAlign.ReplicaStyle = false
		} else {
			tracer.Details("REDACTED")
		}
	}

	//
	//
	txpoolWaitAlign := waitAlign && !settings.LedgerAlign.ReplicaStyle
	txpool, txpoolHandler := instantiateTxpoolAndTxpoolHandler(settings, gatewayApplication, status, txpoolWaitAlign, memplStats, tracer)

	proofHandler, proofDepository, err := instantiateProofHandler(settings, storeSource, statusDepot, ledgerDepot, tracer)
	if err != nil {
		return nil, err
	}

	//
	ledgerExecute := sm.NewLedgerRunner(
		statusDepot,
		tracer.With("REDACTED", "REDACTED"),
		gatewayApplication.Agreement(),
		txpool,
		proofDepository,
		ledgerDepot,
		sm.LedgerRunnerWithStats(machineStats),
	)

	inactiveStatusAlignLevel := int64(0)
	if ledgerDepot.Level() == 0 {
		inactiveStatusAlignLevel, err = ledgerExecute.Depot().FetchInactiveStatusAlignLevel()
		if err != nil && err.Error() != "REDACTED" {
			panic(fmt.Sprintf("REDACTED", err, status.FinalLedgerLevel))
		}
	}

	//
	activateLedgerAlign := ledgerAlign && !statusAlign

	bcodeHandler, err := instantiateChainconnectHandler(
		activateLedgerAlign,
		settings,
		status,
		ledgerExecute,
		ledgerDepot,
		nativeAddress,
		inactiveStatusAlignLevel,
		tracer,
		szStats,
	)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	//
	agreementWaitAlign := waitAlign || settings.LedgerAlign.ReplicaStyle
	agreementHandler, agreementStatus := instantiateAgreementHandler(
		settings, status, ledgerExecute, ledgerDepot, txpool, proofDepository,
		privateRatifier, csStats, agreementWaitAlign, eventBus, agreementTracer, inactiveStatusAlignLevel,
	)

	err = statusDepot.CollectionInactiveStatusAlignLevel(0)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	//
	//
	//
	//
	statusAlignHandler := statusconnect.NewHandler(
		*settings.StatusAlign,
		gatewayApplication.Mirror(),
		gatewayApplication.Inquire(),
		ssStats,
	)
	statusAlignHandler.AssignTracer(tracer.With("REDACTED", "REDACTED"))

	//
	employCometConnectivity := !settings.P2P.LibraryP2PActivated()

	if settings.P2P.PexHandler && !employCometConnectivity {
		settings.P2P.PexHandler = false
		tracer.Details("REDACTED")
	}

	memberDetails, err := createMemberDetails(settings, memberKey, transOrdinaler, generatePaper, status)
	if err != nil {
		return nil, err
	}

	var (
		carrier p2p.Carrier
		sw        p2p.Toggeler
		p2pTracer = tracer.With("REDACTED", "REDACTED")
	)

	//
	if employCometConnectivity {
		cometCarrier, toggeler := instantiateCometCarrierWithRouter(
			settings,
			memberDetails,
			memberKey,
			gatewayApplication,
			txpoolHandler,
			bcodeHandler,
			statusAlignHandler,
			agreementHandler,
			proofHandler,
			p2pStats,
			p2pTracer,
		)

		err = toggeler.AppendDurableNodes(divideAndShaveEmpty(settings.P2P.DurableNodes, "REDACTED", "REDACTED"))
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		err = toggeler.AppendAbsoluteNodeIDXDatastore(divideAndShaveEmpty(settings.P2P.AbsoluteNodeIDXDatastore, "REDACTED", "REDACTED"))
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		addressRegistry, err := instantiateAddressRegistryAndCollectionOnRouter(settings, toggeler, p2pTracer, memberKey)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		//
		if settings.P2P.PexHandler {
			_ = instantiatePEXHandlerAndAppendToRouter(addressRegistry, settings, toggeler, tracer)
		}

		//
		addressRegistry.AppendInternalIDXDatastore(divideAndShaveEmpty(settings.P2P.PrivateNodeIDXDatastore, "REDACTED", "REDACTED"))

		carrier = cometCarrier
		sw = toggeler
	} else {
		p2pTracer.Details("REDACTED")

		handlers := []netpeer.RouterHandler{
			{Label: "REDACTED", Handler: txpoolHandler},
			{Label: "REDACTED", Handler: bcodeHandler},
			{Label: "REDACTED", Handler: agreementHandler},
			{Label: "REDACTED", Handler: proofHandler},
			{Label: "REDACTED", Handler: statusAlignHandler},
		}

		//
		if settings.Txpool.Kind == cfg.TxpoolKindNoop {
			handlers = handlers[1:]
		}

		machine, err := netpeer.NewMachine(settings.P2P, memberKey.PrivateKey, p2pTracer)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		sw, err = netpeer.NewRouter(memberDetails, machine, handlers, p2pStats, p2pTracer)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	member := &Member{
		settings:        settings,
		originPaper:    generatePaper,
		privateRatifier: privateRatifier,

		carrier: carrier,
		sw:        sw,

		memberDetails: memberDetails,
		memberKey:  memberKey,

		statusDepot:       statusDepot,
		ledgerDepot:       ledgerDepot,
		bcodeHandler:        bcodeHandler,
		txpoolHandler:   txpoolHandler,
		txpool:          txpool,
		agreementStatus:   agreementStatus,
		agreementHandler: agreementHandler,
		statusAlignHandler: statusAlignHandler,
		statusAlign:        statusAlign,
		statusAlignOrigin: status, //
		proofDepository:     proofDepository,
		gatewayApplication:         gatewayApplication,
		transOrdinaler:        transOrdinaler,
		ordinalerDaemon:   ordinalerDaemon,
		ledgerOrdinaler:     ledgerOrdinaler,
		eventBus:         eventBus,
	}

	member.RootDaemon = *daemon.NewRootDaemon(tracer, "REDACTED", member)

	for _, setting := range options {
		setting(member)
	}

	return member, nil
}

//
func (n *Member) OnBegin() error {
	now := engineclock.Now()
	generateTime := n.originPaper.OriginMoment
	if generateTime.After(now) {
		n.Tracer.Details("REDACTED", "REDACTED", generateTime)
		time.Sleep(generateTime.Sub(now))
	}

	//
	if n.settings.RPC.IsPprofActivated() {
		n.pprofSvc = n.beginPprofHost()
	}

	//
	if n.settings.Telemetry.IsMonitorstatsActivated() {
		n.monitorstatsSvc = n.beginMonitorstatsHost()
	}

	//
	//
	if n.settings.RPC.AcceptLocation != "REDACTED" {
		observers, err := n.beginRPC()
		if err != nil {
			return err
		}
		n.rpcObservers = observers
	}

	//
	address, err := p2p.NewNetLocationString(p2p.UIDLocationString(n.memberKey.ID(), n.settings.P2P.AcceptLocation))
	if err != nil {
		return err
	}

	if mp, ok := n.carrier.(*p2p.MulticastCarrier); ok {
		if err := mp.Observe(*address); err != nil {
			return err
		}
	}

	//
	err = n.sw.Begin()
	if err != nil {
		return err
	}

	n.isObserving = true

	//
	err = n.sw.CallNodesAsync(divideAndShaveEmpty(n.settings.P2P.DurableNodes, "REDACTED", "REDACTED"))
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if n.statusAlign {
		bcR, ok := n.bcodeHandler.(ledgerAlignHandler)
		if !ok {
			return fmt.Errorf("REDACTED")
		}
		err := beginStatusAlign(n.statusAlignHandler, bcR, n.statusAlignSource,
			n.settings.StatusAlign, n.statusDepot, n.ledgerDepot, n.statusAlignOrigin)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}

	return nil
}

//
func (n *Member) OnHalt() {
	n.RootDaemon.OnHalt()

	n.Tracer.Details("REDACTED")

	//
	if err := n.eventBus.Halt(); err != nil {
		n.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	if n.ordinalerDaemon != nil {
		if err := n.ordinalerDaemon.Halt(); err != nil {
			n.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	//
	if err := n.sw.Halt(); err != nil {
		n.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	if mp, ok := n.carrier.(*p2p.MulticastCarrier); ok {
		if err := mp.End(); err != nil {
			n.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}

	n.isObserving = false

	//
	for _, l := range n.rpcObservers {
		n.Tracer.Details("REDACTED", "REDACTED", l)
		if err := l.Close(); err != nil {
			n.Tracer.Fault("REDACTED", "REDACTED", l, "REDACTED", err)
		}
	}

	if pvsc, ok := n.privateRatifier.(daemon.Daemon); ok {
		if err := pvsc.Halt(); err != nil {
			n.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}

	if n.monitorstatsSvc != nil {
		if err := n.monitorstatsSvc.Shutdown(context.Background()); err != nil {
			//
			n.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	if n.pprofSvc != nil {
		if err := n.pprofSvc.Shutdown(context.Background()); err != nil {
			n.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	if n.ledgerDepot != nil {
		n.Tracer.Details("REDACTED")
		if err := n.ledgerDepot.End(); err != nil {
			n.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	if n.statusDepot != nil {
		n.Tracer.Details("REDACTED")
		if err := n.statusDepot.End(); err != nil {
			n.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	if n.proofDepository != nil {
		n.Tracer.Details("REDACTED")
		if err := n.ProofDepository().End(); err != nil {
			n.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
}

//
func (n *Member) SetupRPC() (*rpcbase.Context, error) {
	publicKey, err := n.privateRatifier.FetchPublicKey()
	if publicKey == nil || err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	rpcCoreContext := rpcbase.Context{
		GatewayApplicationInquire:   n.gatewayApplication.Inquire(),
		GatewayApplicationTxpool: n.gatewayApplication.Txpool(),

		StatusDepot:     n.statusDepot,
		LedgerDepot:     n.ledgerDepot,
		ProofDepository:   n.proofDepository,
		AgreementStatus: n.agreementStatus,
		P2PNodes:       n.sw,
		P2PCarrier:   n,
		PublicKey:         publicKey,

		GeneratePaper:           n.originPaper,
		TransOrdinaler:        n.transOrdinaler,
		LedgerOrdinaler:     n.ledgerOrdinaler,
		AgreementHandler: n.agreementHandler,
		TxpoolHandler:   n.txpoolHandler,
		EventBus:         n.eventBus,
		Txpool:          n.txpool,
		IsReplicaStyle:   n.settings.LedgerAlign.ReplicaStyle,

		Tracer: n.Tracer.With("REDACTED", "REDACTED"),

		Settings: *n.settings.RPC,
	}
	if err := rpcCoreContext.InitOriginSegments(); err != nil {
		return nil, err
	}
	return &rpcCoreContext, nil
}

func (n *Member) beginRPC() ([]net.Listener, error) {
	env, err := n.SetupRPC()
	if err != nil {
		return nil, err
	}

	observeLocations := divideAndShaveEmpty(n.settings.RPC.AcceptLocation, "REDACTED", "REDACTED")
	paths := env.FetchPaths()

	if n.settings.RPC.Risky {
		env.AppendRiskyPaths(paths)
	}

	settings := rpchost.StandardSettings()
	settings.MaximumQueryClusterVolume = n.settings.RPC.MaximumQueryClusterVolume
	settings.MaximumContentOctets = n.settings.RPC.MaximumContentOctets
	settings.MaximumHeadingOctets = n.settings.RPC.MaximumHeadingOctets
	settings.MaximumAccessLinks = n.settings.RPC.MaximumAccessLinks
	//
	//
	//
	if settings.RecordDeadline <= n.settings.RPC.DeadlineMulticastTransEndorse {
		settings.RecordDeadline = n.settings.RPC.DeadlineMulticastTransEndorse + 1*time.Second
	}

	//
	observers := make([]net.Listener, len(observeLocations))
	for i, acceptAddress := range observeLocations {
		mux := http.NewServeMux()
		rpcTracer := n.Tracer.With("REDACTED", "REDACTED")
		wmTracer := rpcTracer.With("REDACTED", "REDACTED")
		wm := rpchost.NewWebchannelOverseer(paths,
			rpchost.OnDetach(func(distantAddress string) {
				err := n.eventBus.DeenrollAll(context.Background(), distantAddress)
				if err != nil && err != cometbroadcast.ErrEnrollmentNegateLocated {
					wmTracer.Fault("REDACTED", "REDACTED", distantAddress, "REDACTED", err)
				}
			}),
			rpchost.ScanCeiling(settings.MaximumContentOctets),
			rpchost.RecordChannelAbility(n.settings.RPC.WebSocketRecordBufferVolume),
		)
		wm.AssignTracer(wmTracer)
		mux.HandleFunc("REDACTED", wm.WebchannelManager)
		rpchost.EnrollRPCRoutines(mux, paths, rpcTracer)
		observer, err := rpchost.Observe(
			acceptAddress,
			settings.MaximumAccessLinks,
		)
		if err != nil {
			return nil, err
		}

		var originManager http.Handler = mux
		if n.settings.RPC.IsCorsActivated() {
			corsInterceptor := cors.New(cors.Options{
				AllowedOrigins: n.settings.RPC.CORSPermittedSources,
				AllowedMethods: n.settings.RPC.CORSPermittedTechniques,
				AllowedHeaders: n.settings.RPC.CORSPermittedHeadings,
			})
			originManager = corsInterceptor.Handler(mux)
		}
		if n.settings.RPC.IsTLSActivated() {
			go func() {
				if err := rpchost.AttendTLS(
					observer,
					originManager,
					n.settings.RPC.TokenEntry(),
					n.settings.RPC.KeyEntry(),
					rpcTracer,
					settings,
				); err != nil {
					n.Tracer.Fault("REDACTED", "REDACTED", err)
				}
			}()
		} else {
			go func() {
				if err := rpchost.Attend(
					observer,
					originManager,
					rpcTracer,
					settings,
				); err != nil {
					n.Tracer.Fault("REDACTED", "REDACTED", err)
				}
			}()
		}

		observers[i] = observer
	}

	//
	grpcObserveAddress := n.settings.RPC.GRPCAcceptLocation
	if grpcObserveAddress != "REDACTED" {
		settings := rpchost.StandardSettings()
		settings.MaximumContentOctets = n.settings.RPC.MaximumContentOctets
		settings.MaximumHeadingOctets = n.settings.RPC.MaximumHeadingOctets
		//
		settings.MaximumAccessLinks = n.settings.RPC.GRPCMaximumAccessLinkages
		//
		//
		//
		if settings.RecordDeadline <= n.settings.RPC.DeadlineMulticastTransEndorse {
			settings.RecordDeadline = n.settings.RPC.DeadlineMulticastTransEndorse + 1*time.Second
		}
		observer, err := rpchost.Observe(grpcObserveAddress, settings.MaximumAccessLinks)
		if err != nil {
			return nil, err
		}
		go func() {
			//
			if err := grpccore.BeginGRPCHost(env, observer); err != nil {
				n.Tracer.Fault("REDACTED", "REDACTED", err)
			}
		}()
		observers = append(observers, observer)

	}

	return observers, nil
}

//
//
func (n *Member) beginMonitorstatsHost() *http.Server {
	srv := &http.Server{
		Addr: n.settings.Telemetry.MonitorstatsObserveAddress,
		Handler: promhttp.InstrumentMetricHandler(
			prometheus.DefaultRegisterer, promhttp.HandlerFor(
				prometheus.DefaultGatherer,
				promhttp.HandlerOpts{MaxRequestsInFlight: n.settings.Telemetry.MaximumAccessLinks},
			),
		),
		ReadHeaderTimeout: readHeadingDeadline,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			//
			n.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}()
	return srv
}

//
func (n *Member) beginPprofHost() *http.Server {
	srv := &http.Server{
		Addr:              n.settings.RPC.PprofAcceptLocation,
		Handler:           nil,
		ReadHeaderTimeout: readHeadingDeadline,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			//
			n.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}()
	return srv
}

//
func (n *Member) Router() p2p.Toggeler {
	return n.sw
}

//
func (n *Member) LedgerDepot() *depot.LedgerDepot {
	return n.ledgerDepot
}

//
func (n *Member) AgreementHandler() *cs.Handler {
	return n.agreementHandler
}

//
func (n *Member) TxpoolHandler() p2p.Handler {
	return n.txpoolHandler
}

//
func (n *Member) Txpool() txpool.Txpool {
	return n.txpool
}

//
func (n *Member) ProofDepository() *proof.Depository {
	return n.proofDepository
}

//
func (n *Member) EventBus() *kinds.EventBus {
	return n.eventBus
}

//
//
func (n *Member) PrivateRatifier() kinds.PrivateRatifier {
	return n.privateRatifier
}

//
func (n *Member) OriginPaper() *kinds.OriginPaper {
	return n.originPaper
}

//
func (n *Member) GatewayApplication() gateway.ApplicationLinks {
	return n.gatewayApplication
}

//
func (n *Member) Settings() *cfg.Settings {
	return n.settings
}

//

func (n *Member) Observers() []string {
	return []string{
		fmt.Sprintf("REDACTED", n.settings.P2P.OutsideLocation),
	}
}

func (n *Member) IsObserving() bool {
	return n.isObserving
}

//
func (n *Member) MemberDetails() p2p.MemberDetails {
	return n.memberDetails
}

func createMemberDetails(
	settings *cfg.Settings,
	memberKey *p2p.MemberKey,
	transOrdinaler transordinal.TransOrdinaler,
	generatePaper *kinds.OriginPaper,
	status sm.Status,
) (p2p.StandardMemberDetails, error) {
	transferOrdinalerState := "REDACTED"
	if _, ok := transOrdinaler.(*void.TransOrdinal); ok {
		transferOrdinalerState = "REDACTED"
	}

	memberDetails := p2p.StandardMemberDetails{
		ProtocolRelease: p2p.NewProtocolRelease(
			release.P2PProtocol, //
			status.Release.Agreement.Ledger,
			status.Release.Agreement.App,
		),
		StandardMemberUID: memberKey.ID(),
		Fabric:       generatePaper.LedgerUID,
		Release:       release.TMCoreSemaphoreRev,
		Streams: []byte{
			bc.ChainconnectStream,
			cs.StatusStream, cs.DataStream, cs.BallotStream, cs.BallotAssignBitsStream,
			txpool.TxpoolConduit,
			proof.ProofConduit,
			statusconnect.MirrorStream, statusconnect.SegmentStream,
		},
		Moniker: settings.Moniker,
		Another: p2p.StandardMemberDetailsAnother{
			TransOrdinal:    transferOrdinalerState,
			RPCLocation: settings.RPC.AcceptLocation,
		},
	}

	if settings.P2P.PexHandler {
		memberDetails.Streams = append(memberDetails.Streams, pex.PexConduit)
	}

	lAddress := settings.P2P.OutsideLocation

	if lAddress == "REDACTED" {
		lAddress = settings.P2P.AcceptLocation
	}

	memberDetails.ObserveAddress = lAddress

	err := memberDetails.Certify()
	return memberDetails, err
}
