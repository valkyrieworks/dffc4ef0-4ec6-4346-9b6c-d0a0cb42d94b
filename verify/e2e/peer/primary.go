package primary

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/node"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	strongmindflags "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/cli/switches"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile"
	adelegate "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/delegate"
	airpc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/rpc"
	dbs "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/depot/db"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/peer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	rpchandler "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/node"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/app"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
)

var tracer = log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))

//
func primary() {
	if len(os.Args) != 2 {
		fmt.Printf("REDACTED", os.Args[0])
		return
	}
	settingsRecord := "REDACTED"
	if len(os.Args) == 2 {
		settingsRecord = os.Args[1]
	}

	if err := run(settingsRecord); err != nil {
		tracer.Failure(err.Error())
		os.Exit(1)
	}
}

//
func run(settingsRecord string) error {
	cfg, err := FetchSettings(settingsRecord)
	if err != nil {
		return err
	}

	//
	if cfg.PrivateItemDaemon != "REDACTED" {
		if err = initiateEndorser(cfg); err != nil {
			return err
		}
		if cfg.Scheme == string(e2e.SchemeIntrinsic) || cfg.Scheme == string(e2e.SchemeIntrinsicLinkChronize) {
			time.Sleep(1 * time.Second)
		}
	}

	//
	switch cfg.Scheme {
	case "REDACTED", "REDACTED":
		err = initiateApplication(cfg)
	case string(e2e.SchemeIntrinsic), string(e2e.SchemeIntrinsicLinkChronize):
		if cfg.Style == string(e2e.StyleAgile) {
			err = initiateAgileCustomer(cfg)
		} else {
			err = initiatePeer(cfg)
		}
	default:
		err = fmt.Errorf("REDACTED", cfg.Scheme)
	}
	if err != nil {
		return err
	}

	//
	for {
		time.Sleep(1 * time.Hour)
	}
}

//
func initiateApplication(cfg *Settings) error {
	app, err := app.FreshPlatform(cfg.App())
	if err != nil {
		return err
	}
	node, err := node.FreshDaemon(cfg.Overhear, cfg.Scheme, app)
	if err != nil {
		return err
	}
	err = node.Initiate()
	if err != nil {
		return err
	}
	tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", cfg.Overhear, cfg.Scheme))
	return nil
}

//
//
//
//
func initiatePeer(cfg *Settings) error {
	app, err := app.FreshPlatform(cfg.App())
	if err != nil {
		return err
	}

	strongmindsettings, peerTracer, peerToken, err := configurePeer()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	var customerOriginator delegate.CustomerOriginator
	if cfg.Scheme == string(e2e.SchemeIntrinsicLinkChronize) {
		customerOriginator = delegate.FreshLinkChronizeRegionalCustomerOriginator(app)
		peerTracer.Details("REDACTED")
	} else {
		customerOriginator = delegate.FreshRegionalCustomerOriginator(app)
		peerTracer.Details("REDACTED")
	}

	n, err := peer.FreshPeer(strongmindsettings,
		privatevalue.FetchEitherProduceRecordPRV(strongmindsettings.PrivateAssessorTokenRecord(), strongmindsettings.PrivateAssessorStatusRecord()),
		peerToken,
		customerOriginator,
		peer.FallbackInaugurationPaperSupplierMethod(strongmindsettings),
		settings.FallbackDatastoreSupplier,
		peer.FallbackTelemetrySupplier(strongmindsettings.Telemetry),
		peerTracer,
	)
	if err != nil {
		return err
	}
	return n.Initiate()
}

func initiateAgileCustomer(cfg *Settings) error {
	strongmindsettings, peerTracer, _, err := configurePeer()
	if err != nil {
		return err
	}

	datastoreEnv := &settings.DatastoreScope{ID: "REDACTED", Settings: strongmindsettings}
	agileDatastore, err := settings.FallbackDatastoreSupplier(datastoreEnv)
	if err != nil {
		return err
	}

	suppliers := remoteTerminals(strongmindsettings.P2P.EnduringNodes)

	c, err := agile.FreshHttpsvcCustomer(
		context.Background(),
		cfg.SuccessionUUID,
		agile.RelianceChoices{
			Cycle: strongmindsettings.StatusChronize.RelianceSpan,
			Altitude: strongmindsettings.StatusChronize.RelianceAltitude,
			Digest:   strongmindsettings.StatusChronize.RelianceDigestOctets(),
		},
		suppliers[0],
		suppliers[1:],
		dbs.New(agileDatastore, "REDACTED"),
		agile.Tracer(peerTracer),
	)
	if err != nil {
		return err
	}

	remotecfg := rpchandler.FallbackSettings()
	remotecfg.MaximumContentOctets = strongmindsettings.RPC.MaximumContentOctets
	remotecfg.MaximumHeadingOctets = strongmindsettings.RPC.MaximumHeadingOctets
	remotecfg.MaximumInitiateLinks = strongmindsettings.RPC.MaximumInitiateLinks
	//
	//
	//
	if remotecfg.PersistDeadline <= strongmindsettings.RPC.DeadlineMulticastTransferEndorse {
		remotecfg.PersistDeadline = strongmindsettings.RPC.DeadlineMulticastTransferEndorse + 1*time.Second
	}

	p, err := adelegate.FreshDelegate(c, strongmindsettings.RPC.OverhearLocation, suppliers[0], remotecfg, peerTracer,
		airpc.TokenRouteProc(airpc.FallbackHashmapTokenRouteProc()))
	if err != nil {
		return err
	}

	tracer.Details("REDACTED", "REDACTED", strongmindsettings.RPC.OverhearLocation)
	if err := p.OverhearAlsoAttend(); err != http.ErrServerClosed {
		//
		tracer.Failure("REDACTED", "REDACTED", err)
	}

	return nil
}

//
func initiateEndorser(cfg *Settings) error {
	recordPRV := privatevalue.FetchRecordPRV(cfg.PrivateItemToken, cfg.PrivateItemStatus)

	scheme, location := strongmindnet.SchemeAlsoLocation(cfg.PrivateItemDaemon)
	var callProc privatevalue.PortCaller
	switch scheme {
	case "REDACTED":
		callProc = privatevalue.CallStreamProc(location, 3*time.Second, edwards25519.ProducePrivateToken())
	case "REDACTED":
		callProc = privatevalue.CallPosixProc(location)
	default:
		return fmt.Errorf("REDACTED", scheme)
	}

	gateway := privatevalue.FreshEndorserCallerGateway(tracer, callProc,
		privatevalue.EndorserCallerGatewayReissuePauseDuration(1*time.Second),
		privatevalue.EndorserCallerGatewayLinkAttempts(100))
	err := privatevalue.FreshEndorserDaemon(gateway, cfg.SuccessionUUID, recordPRV).Initiate()
	if err != nil {
		return err
	}
	tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", cfg.PrivateItemDaemon))
	return nil
}

func configurePeer() (*settings.Settings, log.Tracer, *p2p.PeerToken, error) {
	var strongmindsettings *settings.Settings

	domain := os.Getenv("REDACTED")
	if domain == "REDACTED" {
		return nil, nil, nil, errors.New("REDACTED")
	}

	viper.AddConfigPath(filepath.Join(domain, "REDACTED"))
	viper.SetConfigName("REDACTED")

	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, nil, err
	}

	strongmindsettings = settings.FallbackSettings()

	if err := viper.Unmarshal(strongmindsettings); err != nil {
		return nil, nil, nil, err
	}

	strongmindsettings.AssignOrigin(domain)

	if err := strongmindsettings.CertifyFundamental(); err != nil {
		return nil, nil, nil, fmt.Errorf("REDACTED", err)
	}

	if strongmindsettings.RecordLayout == settings.RecordLayoutJSN {
		tracer = log.FreshTempjsonTracer(log.FreshChronizePersistor(os.Stdout))
	}

	peerTracer, err := strongmindflags.AnalyzeRecordStratum(strongmindsettings.RecordStratum, tracer, settings.FallbackRecordStratum)
	if err != nil {
		return nil, nil, nil, err
	}

	peerTracer = peerTracer.Using("REDACTED", "REDACTED")

	peerToken, err := p2p.FetchEitherProducePeerToken(strongmindsettings.PeerTokenRecord())
	if err != nil {
		return nil, nil, nil, fmt.Errorf("REDACTED", strongmindsettings.PeerTokenRecord(), err)
	}

	return strongmindsettings, peerTracer, peerToken, nil
}

//
//
func remoteTerminals(nodes string) []string {
	arr := strings.Split(nodes, "REDACTED")
	terminals := make([]string, len(arr))
	for i, v := range arr {
		webrouteText := strings.SplitAfter(v, "REDACTED")[1]
		machineAlias := strings.Split(webrouteText, "REDACTED")[0]
		//
		channel := 26657
		remoteGateway := "REDACTED" + machineAlias + "REDACTED" + fmt.Sprint(channel)
		terminals[i] = remoteGateway
	}
	return terminals
}
