package peer

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

	bc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/chainchronize"
	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	cs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/proof"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	tendermintpubsub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/netp2p"
	txpooll "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/pex"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	remotecore "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base"
	grpcshell "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/grps"
	rpchandler "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/node"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal/nothing"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/statuschronize"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"

	_ "net/http/pprof" //
)

//
//
type Peer struct {
	facility.FoundationFacility

	//
	settings        *cfg.Settings
	inaugurationPaper    *kinds.OriginPaper   //
	privateAssessor kinds.PrivateAssessor //

	//
	carrier   p2p.Carrier
	sw          p2p.Router //
	peerDetails    p2p.PeerDetails
	peerToken     *p2p.PeerToken //
	equalsObserving bool

	//
	incidentPipeline          *kinds.IncidentChannel //
	statusDepot        sm.Depot
	ledgerDepot        *depot.LedgerDepot //
	bchainHandler         p2p.Handler       //
	txpoolHandler    pauseChronizeHandler   //
	txpool           txpooll.Txpool
	statusChronize         bool                    //
	statusChronizeHandler  *statuschronize.Handler      //
	statusChronizeSupplier statuschronize.StatusSupplier //
	statusChronizeInauguration  sm.Status                //
	agreementStatus    *cs.Status               //
	agreementHandler  *cs.Handler             //
	proofHub      *proof.Hub          //
	delegatePlatform          delegate.PlatformLinks          //
	remoteObservers      []net.Listener          //
	transferOrdinalizer         transferordinal.TransferOrdinalizer
	ledgerOrdinalizer      ordinalizer.LedgerOrdinalizer
	ordinalizerFacility    *transferordinal.OrdinalizerFacility
	titanDaemon     *http.Server
	profilerDaemon          *http.Server
}

type pauseChronizeHandler interface {
	p2p.Handler
	//
	AwaitChronize() bool
}

//
type Selection func(*Peer)

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
func BespokeEngines(engines map[string]p2p.Handler) Selection {
	return func(n *Peer) {
		for alias, handler := range engines {
			if extantHandler, ok := n.sw.Handler(alias); ok {
				n.sw.Log().Details("REDACTED",
					"REDACTED", alias, "REDACTED", extantHandler, "REDACTED", handler)
				n.sw.DiscardHandler(alias, extantHandler)
			}
			n.sw.AppendHandler(alias, handler)

			//
			//
			//
			//
			ni, ok := n.peerDetails.(p2p.FallbackPeerDetails)
			if !ok {
				n.Tracer.Failure("REDACTED")
				continue
			}

			mp, ok := n.carrier.(*p2p.MultiplexCarrier)
			if !ok {
				n.Tracer.Failure("REDACTED")
				continue
			}

			for _, chnlDescription := range handler.ObtainConduits() {
				if ni.OwnsConduit(chnlDescription.ID) {
					continue
				}

				ni.Conduits = append(ni.Conduits, chnlDescription.ID)
				mp.AppendConduit(chnlDescription.ID)
			}

			n.peerDetails = ni
		}
	}
}

//
//
//
func StatusSupplier(statusSupplier statuschronize.StatusSupplier) Selection {
	return func(n *Peer) {
		n.statusChronizeSupplier = statusSupplier
	}
}

//
//
//
//
//
func OnboardStatus(ctx context.Context, settings *cfg.Settings, datastoreSupplier cfg.DatastoreSupplier, altitude uint64, platformDigest []byte) error {
	return OnboardStatusUsingProduceSupplier(ctx, settings, datastoreSupplier, FallbackInaugurationPaperSupplierMethod(settings), altitude, platformDigest)
}

//
//
//
//
//
func OnboardStatusUsingProduceSupplier(ctx context.Context, settings *cfg.Settings, datastoreSupplier cfg.DatastoreSupplier, produceSupplier InaugurationPaperSupplier, altitude uint64, platformDigest []byte) (err error) {
	tracer := log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))
	if ctx == nil {
		ctx = context.Background()
	}

	if settings == nil {
		tracer.Details("REDACTED")
		settings = cfg.FallbackSettings()
	}

	if datastoreSupplier == nil {
		datastoreSupplier = cfg.FallbackDatastoreSupplier
	}
	ledgerDepot, statusDatastore, err := initializeDeltaBytes(settings, datastoreSupplier)

	defer func() {
		if efault := ledgerDepot.Shutdown(); efault != nil {
			tracer.Failure("REDACTED", "REDACTED", efault)
			//
			err = efault
		}
	}()

	if err != nil {
		return err
	}

	if !ledgerDepot.EqualsBlank() {
		return fmt.Errorf("REDACTED")
	}

	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: settings.Repository.EjectIfaceReplies,
	})

	defer func() {
		if efault := statusDepot.Shutdown(); efault != nil {
			tracer.Failure("REDACTED", "REDACTED", efault)
			//
			err = efault
		}
	}()
	status, err := statusDepot.Fetch()
	if err != nil {
		return err
	}

	if !status.EqualsBlank() {
		return fmt.Errorf("REDACTED")
	}

	produceStatus, _, err := FetchStatusOriginatingDatastoreEitherInaugurationPaperSupplier(statusDatastore, produceSupplier)
	if err != nil {
		return err
	}

	statusSupplier, err := statuschronize.FreshAgileCustomerStatusSupplier(
		ctx,
		produceStatus.SuccessionUUID, produceStatus.Edition, produceStatus.PrimaryAltitude,
		settings.StatusChronize.RemoteHosts, agile.RelianceChoices{
			Cycle: settings.StatusChronize.RelianceSpan,
			Altitude: settings.StatusChronize.RelianceAltitude,
			Digest:   settings.StatusChronize.RelianceDigestOctets(),
		}, tracer.Using("REDACTED", "REDACTED"))
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	status, err = statusSupplier.Status(ctx, altitude)
	if err != nil {
		return err
	}
	if platformDigest == nil {
		tracer.Details("REDACTED")
	} else if !bytes.Equal(platformDigest, status.PlatformDigest) {
		if err := ledgerDepot.Shutdown(); err != nil {
			tracer.Failure("REDACTED", err)
		}
		if err := statusDepot.Shutdown(); err != nil {
			tracer.Failure("REDACTED", err)
		}
		return fmt.Errorf("REDACTED", status.PlatformDigest, platformDigest)

	}

	endorse, err := statusSupplier.Endorse(ctx, altitude)
	if err != nil {
		return err
	}

	if err = statusDepot.Onboard(status); err != nil {
		return err
	}

	err = ledgerDepot.PersistObservedEndorse(status.FinalLedgerAltitude, endorse)
	if err != nil {
		return err
	}

	//
	//
	//
	//
	err = statusDepot.AssignInactiveStatusChronizeAltitude(status.FinalLedgerAltitude)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return err
}

//

//
func FreshPeer(
	settings *cfg.Settings,
	privateAssessor kinds.PrivateAssessor,
	peerToken *p2p.PeerToken,
	customerOriginator delegate.CustomerOriginator,
	inaugurationPaperSupplier InaugurationPaperSupplier,
	datastoreSupplier cfg.DatastoreSupplier,
	telemetrySupplier TelemetrySupplier,
	tracer log.Tracer,
	choices ...Selection,
) (*Peer, error) {
	return FreshPeerUsingEnv(context.TODO(), settings, privateAssessor,
		peerToken, customerOriginator, inaugurationPaperSupplier, datastoreSupplier,
		telemetrySupplier, tracer, choices...)
}

//
func FreshPeerUsingEnv(
	ctx context.Context,
	settings *cfg.Settings,
	privateAssessor kinds.PrivateAssessor,
	peerToken *p2p.PeerToken,
	customerOriginator delegate.CustomerOriginator,
	inaugurationPaperSupplier InaugurationPaperSupplier,
	datastoreSupplier cfg.DatastoreSupplier,
	telemetrySupplier TelemetrySupplier,
	tracer log.Tracer,
	choices ...Selection,
) (*Peer, error) {
	ledgerDepot, statusDatastore, err := initializeDeltaBytes(settings, datastoreSupplier)
	if err != nil {
		return nil, err
	}

	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: settings.Repository.EjectIfaceReplies,
	})

	status, producePaper, err := FetchStatusOriginatingDatastoreEitherInaugurationPaperSupplier(statusDatastore, inaugurationPaperSupplier)
	if err != nil {
		return nil, err
	}

	controlTelemetry, peer2peerTelemetry, txpoollTelemetry, machineTelemetry, ifaceTelemetry, bytesTelemetry, sstoreTelemetry := telemetrySupplier(producePaper.SuccessionUUID)

	//
	delegatePlatform, err := generateAlsoInitiateDelegateApplicationLinks(customerOriginator, tracer, ifaceTelemetry)
	if err != nil {
		return nil, err
	}

	//
	//
	//
	//
	incidentPipeline, err := generateAlsoInitiateIncidentPipeline(tracer)
	if err != nil {
		return nil, err
	}

	ordinalizerFacility, transferOrdinalizer, ledgerOrdinalizer, err := generateAlsoInitiateOrdinalizerFacility(settings,
		producePaper.SuccessionUUID, datastoreSupplier, incidentPipeline, tracer)
	if err != nil {
		return nil, err
	}

	//
	//
	if settings.PrivateAssessorOverhearLocation != "REDACTED" {
		//
		privateAssessor, err = generateAlsoInitiatePrivateAssessorPortCustomer(settings.PrivateAssessorOverhearLocation, producePaper.SuccessionUUID, tracer)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	publicToken, err := privateAssessor.ObtainPublicToken()
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	regionalLocation := publicToken.Location()

	//
	statusChronize := settings.StatusChronize.Activate && !solelyAssessorEqualsWe(status, regionalLocation)
	if statusChronize && status.FinalLedgerAltitude > 0 {
		tracer.Details("REDACTED")
		statusChronize = false
	}

	//
	//
	agreementTracer := tracer.Using("REDACTED", "REDACTED")
	if !statusChronize {
		if err := conductNegotiation(ctx, statusDepot, status, ledgerDepot, producePaper, incidentPipeline, delegatePlatform, agreementTracer); err != nil {
			return nil, err
		}

		//
		//
		//
		status, err = statusDepot.Fetch()
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	reportPeerLaunchDetails(status, publicToken, tracer, agreementTracer)

	//
	var (
		//
		//
		activateLedgerChronize = !solelyAssessorEqualsWe(status, regionalLocation) && !statusChronize

		//
		//
		agreementPauseForeachChronize = statusChronize || (activateLedgerChronize && !settings.LedgerChronize.AggregateStyle)

		//
		//
		txpoolPauseForeachChronize = statusChronize || (activateLedgerChronize && !settings.LedgerChronize.AggregateStyle)
	)

	if settings.LedgerChronize.AggregateStyle {
		tracer.Details("REDACTED")
	}

	//
	txpool, txpoolHandler := generateTxpoolAlsoTxpoolHandler(settings, delegatePlatform, status, txpoolPauseForeachChronize, txpoollTelemetry, tracer)

	proofHandler, proofHub, err := generateProofHandler(settings, datastoreSupplier, statusDepot, ledgerDepot, tracer)
	if err != nil {
		return nil, err
	}

	//
	ledgerExecute := sm.FreshLedgerHandler(
		statusDepot,
		tracer.Using("REDACTED", "REDACTED"),
		delegatePlatform.Agreement(),
		txpool,
		proofHub,
		ledgerDepot,
		sm.LedgerHandlerUsingTelemetry(machineTelemetry),
	)

	inactiveStatusChronizeAltitude := int64(0)
	if ledgerDepot.Altitude() == 0 {
		inactiveStatusChronizeAltitude, err = ledgerExecute.Depot().ObtainInactiveStatusChronizeAltitude()
		if err != nil && err.Error() != "REDACTED" {
			panic(fmt.Sprintf("REDACTED", err, status.FinalLedgerAltitude))
		}
	}

	bchainHandler, err := generateChainchronizeHandler(
		activateLedgerChronize,
		settings,
		status,
		ledgerExecute,
		ledgerDepot,
		regionalLocation,
		inactiveStatusChronizeAltitude,
		tracer,
		bytesTelemetry,
	)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	agreementHandler, agreementStatus := generateAgreementHandler(
		settings, status, ledgerExecute, ledgerDepot, txpool, proofHub,
		privateAssessor, controlTelemetry, agreementPauseForeachChronize, incidentPipeline, agreementTracer, inactiveStatusChronizeAltitude,
	)

	err = statusDepot.AssignInactiveStatusChronizeAltitude(0)
	if err != nil {
		panic(fmt.Sprintf("REDACTED", err))
	}
	//
	//
	//
	//
	statusChronizeHandler := statuschronize.FreshHandler(
		*settings.StatusChronize,
		delegatePlatform.Image(),
		delegatePlatform.Inquire(),
		sstoreTelemetry,
	)
	statusChronizeHandler.AssignTracer(tracer.Using("REDACTED", "REDACTED"))

	//
	utilizeStrongNetting := !settings.P2P.LibraryPeer2peerActivated()

	if settings.P2P.PeerxHandler && !utilizeStrongNetting {
		settings.P2P.PeerxHandler = false
		tracer.Details("REDACTED")
	}

	peerDetails, err := createPeerDetails(settings, peerToken, transferOrdinalizer, producePaper, status)
	if err != nil {
		return nil, err
	}

	var (
		carrier p2p.Carrier
		sw        p2p.Router
		peer2peerTracer = tracer.Using("REDACTED", "REDACTED")
	)

	//
	if utilizeStrongNetting {
		strongCarrier, selector := generateStrongCarrierUsingRouter(
			settings,
			peerDetails,
			peerToken,
			delegatePlatform,
			txpoolHandler,
			bchainHandler,
			statusChronizeHandler,
			agreementHandler,
			proofHandler,
			peer2peerTelemetry,
			peer2peerTracer,
		)

		err = selector.AppendEnduringNodes(partitionAlsoShaveBlank(settings.P2P.EnduringNodes, "REDACTED", "REDACTED"))
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		err = selector.AppendAbsoluteNodeIDXDstore(partitionAlsoShaveBlank(settings.P2P.AbsoluteNodeIDXDstore, "REDACTED", "REDACTED"))
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		locationRegister, err := generateLocationRegisterAlsoAssignUponRouter(settings, selector, peer2peerTracer, peerToken)
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
		if settings.P2P.PeerxHandler {
			_ = generatePeerxHandlerAlsoAppendTowardRouter(locationRegister, settings, selector, tracer)
		}

		//
		locationRegister.AppendSecludedIDXDstore(partitionAlsoShaveBlank(settings.P2P.SecludedNodeIDXDstore, "REDACTED", "REDACTED"))

		carrier = strongCarrier
		sw = selector
	} else {
		peer2peerTracer.Details("REDACTED")

		engines := []netp2p.RouterHandler{
			{Alias: "REDACTED", Handler: txpoolHandler},
			{Alias: "REDACTED", Handler: bchainHandler},
			{Alias: "REDACTED", Handler: agreementHandler},
			{Alias: "REDACTED", Handler: proofHandler},
			{Alias: "REDACTED", Handler: statusChronizeHandler},
		}

		//
		if settings.Txpool.Kind == cfg.TxpoolKindNooperation {
			engines = engines[1:]
		}

		machine, err := netp2p.FreshMachine(settings.P2P, peerToken.PrivateToken, peer2peerTracer)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		sw, err = netp2p.FreshRouter(peerDetails, machine, engines, peer2peerTelemetry, peer2peerTracer)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	peer := &Peer{
		settings:        settings,
		inaugurationPaper:    producePaper,
		privateAssessor: privateAssessor,

		carrier: carrier,
		sw:        sw,

		peerDetails: peerDetails,
		peerToken:  peerToken,

		statusDepot:       statusDepot,
		ledgerDepot:       ledgerDepot,
		bchainHandler:        bchainHandler,
		txpoolHandler:   txpoolHandler,
		txpool:          txpool,
		agreementStatus:   agreementStatus,
		agreementHandler: agreementHandler,
		statusChronizeHandler: statusChronizeHandler,
		statusChronize:        statusChronize,
		statusChronizeInauguration: status, //
		proofHub:     proofHub,
		delegatePlatform:         delegatePlatform,
		transferOrdinalizer:        transferOrdinalizer,
		ordinalizerFacility:   ordinalizerFacility,
		ledgerOrdinalizer:     ledgerOrdinalizer,
		incidentPipeline:         incidentPipeline,
	}

	peer.FoundationFacility = *facility.FreshFoundationFacility(tracer, "REDACTED", peer)

	for _, selection := range choices {
		selection(peer)
	}

	return peer, nil
}

//
func (n *Peer) UponInitiate() error {
	now := committime.Now()
	produceMoment := n.inaugurationPaper.OriginMoment
	if produceMoment.After(now) {
		n.Tracer.Details("REDACTED", "REDACTED", produceMoment)
		time.Sleep(produceMoment.Sub(now))
	}

	//
	if n.settings.RPC.EqualsProfilerActivated() {
		n.profilerDaemon = n.initiateProfilerDaemon()
	}

	//
	if n.settings.Telemetry.EqualsTitanActivated() {
		n.titanDaemon = n.initiateTitanDaemon()
	}

	//
	//
	if n.settings.RPC.OverhearLocation != "REDACTED" {
		observers, err := n.initiateRemote()
		if err != nil {
			return err
		}
		n.remoteObservers = observers
	}

	//
	location, err := p2p.FreshNetworkLocatorText(p2p.UUIDLocationText(n.peerToken.ID(), n.settings.P2P.OverhearLocation))
	if err != nil {
		return err
	}

	if mp, ok := n.carrier.(*p2p.MultiplexCarrier); ok {
		if err := mp.Overhear(*location); err != nil {
			return err
		}
	}

	//
	err = n.sw.Initiate()
	if err != nil {
		return err
	}

	n.equalsObserving = true

	//
	err = n.sw.CallNodesAsyncronous(partitionAlsoShaveBlank(n.settings.P2P.EnduringNodes, "REDACTED", "REDACTED"))
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if n.statusChronize {
		bcR, ok := n.bchainHandler.(ledgerChronizeHandler)
		if !ok {
			return fmt.Errorf("REDACTED")
		}
		err := initiateStatusChronize(n.statusChronizeHandler, bcR, n.statusChronizeSupplier,
			n.settings.StatusChronize, n.statusDepot, n.ledgerDepot, n.statusChronizeInauguration)
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}

	return nil
}

//
func (n *Peer) UponHalt() {
	n.FoundationFacility.UponHalt()

	n.Tracer.Details("REDACTED")

	//
	if err := n.incidentPipeline.Halt(); err != nil {
		n.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	if n.ordinalizerFacility != nil {
		if err := n.ordinalizerFacility.Halt(); err != nil {
			n.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	//
	if err := n.sw.Halt(); err != nil {
		n.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	if mp, ok := n.carrier.(*p2p.MultiplexCarrier); ok {
		if err := mp.Shutdown(); err != nil {
			n.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}

	n.equalsObserving = false

	//
	for _, l := range n.remoteObservers {
		n.Tracer.Details("REDACTED", "REDACTED", l)
		if err := l.Close(); err != nil {
			n.Tracer.Failure("REDACTED", "REDACTED", l, "REDACTED", err)
		}
	}

	if prvtcs, ok := n.privateAssessor.(facility.Facility); ok {
		if err := prvtcs.Halt(); err != nil {
			n.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}

	if n.titanDaemon != nil {
		if err := n.titanDaemon.Shutdown(context.Background()); err != nil {
			//
			n.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	if n.profilerDaemon != nil {
		if err := n.profilerDaemon.Shutdown(context.Background()); err != nil {
			n.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	if n.ledgerDepot != nil {
		n.Tracer.Details("REDACTED")
		if err := n.ledgerDepot.Shutdown(); err != nil {
			n.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	if n.statusDepot != nil {
		n.Tracer.Details("REDACTED")
		if err := n.statusDepot.Shutdown(); err != nil {
			n.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	if n.proofHub != nil {
		n.Tracer.Details("REDACTED")
		if err := n.ProofHub().Shutdown(); err != nil {
			n.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
}

//
func (n *Peer) SetupRemote() (*remotecore.Context, error) {
	publicToken, err := n.privateAssessor.ObtainPublicToken()
	if publicToken == nil || err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	remoteBaseEnvironment := remotecore.Context{
		DelegateApplicationInquire:   n.delegatePlatform.Inquire(),
		DelegateApplicationTxpool: n.delegatePlatform.Txpool(),

		StatusDepot:     n.statusDepot,
		LedgerDepot:     n.ledgerDepot,
		ProofHub:   n.proofHub,
		AgreementStatus: n.agreementStatus,
		Peer2peerNodes:       n.sw,
		Peer2peerCarrier:   n,
		PublicToken:         publicToken,

		ProducePaper:           n.inaugurationPaper,
		TransferOrdinalizer:        n.transferOrdinalizer,
		LedgerOrdinalizer:     n.ledgerOrdinalizer,
		AgreementHandler: n.agreementHandler,
		TxpoolHandler:   n.txpoolHandler,
		IncidentChannel:         n.incidentPipeline,
		Txpool:          n.txpool,
		EqualsAggregateStyle:   n.settings.LedgerChronize.AggregateStyle,

		Tracer: n.Tracer.Using("REDACTED", "REDACTED"),

		Settings: *n.settings.RPC,
	}
	if err := remoteBaseEnvironment.InitializeInaugurationSegments(); err != nil {
		return nil, err
	}
	return &remoteBaseEnvironment, nil
}

func (n *Peer) initiateRemote() ([]net.Listener, error) {
	env, err := n.SetupRemote()
	if err != nil {
		return nil, err
	}

	overhearLocations := partitionAlsoShaveBlank(n.settings.RPC.OverhearLocation, "REDACTED", "REDACTED")
	paths := env.ObtainPaths()

	if n.settings.RPC.Insecure {
		env.AppendInsecurePaths(paths)
	}

	settings := rpchandler.FallbackSettings()
	settings.MaximumSolicitClusterExtent = n.settings.RPC.MaximumSolicitClusterExtent
	settings.MaximumContentOctets = n.settings.RPC.MaximumContentOctets
	settings.MaximumHeadingOctets = n.settings.RPC.MaximumHeadingOctets
	settings.MaximumInitiateLinks = n.settings.RPC.MaximumInitiateLinks
	//
	//
	//
	if settings.PersistDeadline <= n.settings.RPC.DeadlineMulticastTransferEndorse {
		settings.PersistDeadline = n.settings.RPC.DeadlineMulticastTransferEndorse + 1*time.Second
	}

	//
	observers := make([]net.Listener, len(overhearLocations))
	for i, overhearLocation := range overhearLocations {
		mux := http.NewServeMux()
		remoteTracer := n.Tracer.Using("REDACTED", "REDACTED")
		watermarkTracer := remoteTracer.Using("REDACTED", "REDACTED")
		wm := rpchandler.FreshWebterminalAdministrator(paths,
			rpchandler.UponDetach(func(distantLocation string) {
				err := n.incidentPipeline.UnlistenEvery(context.Background(), distantLocation)
				if err != nil && err != tendermintpubsub.FaultListeningNegationDetected {
					watermarkTracer.Failure("REDACTED", "REDACTED", distantLocation, "REDACTED", err)
				}
			}),
			rpchandler.RetrieveThreshold(settings.MaximumContentOctets),
			rpchandler.PersistChnVolume(n.settings.RPC.InternetPortRecordReserveExtent),
		)
		wm.AssignTracer(watermarkTracer)
		mux.HandleFunc("REDACTED", wm.WebterminalProcessor)
		rpchandler.EnrollRemoteRoutines(mux, paths, remoteTracer)
		observer, err := rpchandler.Overhear(
			overhearLocation,
			settings.MaximumInitiateLinks,
		)
		if err != nil {
			return nil, err
		}

		var originProcessor http.Handler = mux
		if n.settings.RPC.EqualsCrossoriginActivated() {
			crossoriginIntermediary := cors.New(cors.Options{
				AllowedOrigins: n.settings.RPC.CrossoriginPermittedSources,
				AllowedMethods: n.settings.RPC.CrossoriginPermittedApproaches,
				AllowedHeaders: n.settings.RPC.CrossoriginPermittedHeadings,
			})
			originProcessor = crossoriginIntermediary.Handler(mux)
		}
		if n.settings.RPC.EqualsTransportsecActivated() {
			go func() {
				if err := rpchandler.AttendTransportsec(
					observer,
					originProcessor,
					n.settings.RPC.LicenseRecord(),
					n.settings.RPC.TokenRecord(),
					remoteTracer,
					settings,
				); err != nil {
					n.Tracer.Failure("REDACTED", "REDACTED", err)
				}
			}()
		} else {
			go func() {
				if err := rpchandler.Attend(
					observer,
					originProcessor,
					remoteTracer,
					settings,
				); err != nil {
					n.Tracer.Failure("REDACTED", "REDACTED", err)
				}
			}()
		}

		observers[i] = observer
	}

	//
	grpsOverhearLocation := n.settings.RPC.GRPSOverhearLocation
	if grpsOverhearLocation != "REDACTED" {
		settings := rpchandler.FallbackSettings()
		settings.MaximumContentOctets = n.settings.RPC.MaximumContentOctets
		settings.MaximumHeadingOctets = n.settings.RPC.MaximumHeadingOctets
		//
		settings.MaximumInitiateLinks = n.settings.RPC.GRPSMaximumUnlockLinkages
		//
		//
		//
		if settings.PersistDeadline <= n.settings.RPC.DeadlineMulticastTransferEndorse {
			settings.PersistDeadline = n.settings.RPC.DeadlineMulticastTransferEndorse + 1*time.Second
		}
		observer, err := rpchandler.Overhear(grpsOverhearLocation, settings.MaximumInitiateLinks)
		if err != nil {
			return nil, err
		}
		go func() {
			//
			if err := grpcshell.InitiateGRPSDaemon(env, observer); err != nil {
				n.Tracer.Failure("REDACTED", "REDACTED", err)
			}
		}()
		observers = append(observers, observer)

	}

	return observers, nil
}

//
//
func (n *Peer) initiateTitanDaemon() *http.Server {
	srv := &http.Server{
		Addr: n.settings.Telemetry.TitanOverhearLocation,
		Handler: promhttp.InstrumentMetricHandler(
			prometheus.DefaultRegisterer, promhttp.HandlerFor(
				prometheus.DefaultGatherer,
				promhttp.HandlerOpts{MaxRequestsInFlight: n.settings.Telemetry.MaximumInitiateLinks},
			),
		),
		ReadHeaderTimeout: fetchHeadlineDeadline,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			//
			n.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}()
	return srv
}

//
func (n *Peer) initiateProfilerDaemon() *http.Server {
	srv := &http.Server{
		Addr:              n.settings.RPC.ProfilerOverhearLocation,
		Handler:           nil,
		ReadHeaderTimeout: fetchHeadlineDeadline,
	}
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			//
			n.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}()
	return srv
}

//
func (n *Peer) Router() p2p.Router {
	return n.sw
}

//
func (n *Peer) LedgerDepot() *depot.LedgerDepot {
	return n.ledgerDepot
}

//
func (n *Peer) AgreementHandler() *cs.Handler {
	return n.agreementHandler
}

//
func (n *Peer) TxpoolHandler() p2p.Handler {
	return n.txpoolHandler
}

//
func (n *Peer) Txpool() txpooll.Txpool {
	return n.txpool
}

//
func (n *Peer) ProofHub() *proof.Hub {
	return n.proofHub
}

//
func (n *Peer) IncidentChannel() *kinds.IncidentChannel {
	return n.incidentPipeline
}

//
//
func (n *Peer) PrivateAssessor() kinds.PrivateAssessor {
	return n.privateAssessor
}

//
func (n *Peer) OriginPaper() *kinds.OriginPaper {
	return n.inaugurationPaper
}

//
func (n *Peer) DelegateApplication() delegate.PlatformLinks {
	return n.delegatePlatform
}

//
func (n *Peer) Settings() *cfg.Settings {
	return n.settings
}

//

func (n *Peer) Observers() []string {
	return []string{
		fmt.Sprintf("REDACTED", n.settings.P2P.OutsideLocation),
	}
}

func (n *Peer) EqualsObserving() bool {
	return n.equalsObserving
}

//
func (n *Peer) PeerDetails() p2p.PeerDetails {
	return n.peerDetails
}

func createPeerDetails(
	settings *cfg.Settings,
	peerToken *p2p.PeerToken,
	transferOrdinalizer transferordinal.TransferOrdinalizer,
	producePaper *kinds.OriginPaper,
	status sm.Status,
) (p2p.FallbackPeerDetails, error) {
	transferOrdinalizerCondition := "REDACTED"
	if _, ok := transferOrdinalizer.(*nothing.TransferOrdinal); ok {
		transferOrdinalizerCondition = "REDACTED"
	}

	peerDetails := p2p.FallbackPeerDetails{
		SchemeEdition: p2p.FreshSchemeEdition(
			edition.Peer2peerScheme, //
			status.Edition.Agreement.Ledger,
			status.Edition.Agreement.App,
		),
		FallbackPeerUUID: peerToken.ID(),
		Fabric:       producePaper.SuccessionUUID,
		Edition:       edition.TEMPBaseSemaphoreEdtn,
		Conduits: []byte{
			bc.ChainchronizeConduit,
			cs.StatusConduit, cs.DataConduit, cs.BallotConduit, cs.BallotAssignDigitsConduit,
			txpooll.TxpoolConduit,
			proof.ProofConduit,
			statuschronize.ImageConduit, statuschronize.SegmentConduit,
		},
		Pseudonym: settings.Pseudonym,
		Another: p2p.FallbackPeerDetailsAnother{
			TransferOrdinal:    transferOrdinalizerCondition,
			RemoteLocator: settings.RPC.OverhearLocation,
		},
	}

	if settings.P2P.PeerxHandler {
		peerDetails.Conduits = append(peerDetails.Conduits, pex.PeerxConduit)
	}

	lnLocation := settings.P2P.OutsideLocation

	if lnLocation == "REDACTED" {
		lnLocation = settings.P2P.OverhearLocation
	}

	peerDetails.OverhearLocation = lnLocation

	err := peerDetails.Certify()
	return peerDetails, err
}
