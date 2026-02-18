package main

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

	"github.com/valkyrieworks/iface/host"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/vault/ed25519"
	cometmarks "github.com/valkyrieworks/utils/cli/marks"
	"github.com/valkyrieworks/utils/log"
	cometnet "github.com/valkyrieworks/utils/net"
	"github.com/valkyrieworks/rapid"
	rapidgateway "github.com/valkyrieworks/rapid/gateway"
	lrpc "github.com/valkyrieworks/rapid/rpc"
	dbs "github.com/valkyrieworks/rapid/depot/db"
	"github.com/valkyrieworks/member"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/privatekey"
	"github.com/valkyrieworks/gateway"
	rpchost "github.com/valkyrieworks/rpc/jsonrpc/host"
	"github.com/valkyrieworks/verify/e2e/app"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
)

var tracer = log.NewTMTracer(log.NewAlignRecorder(os.Stdout))

//
func main() {
	if len(os.Args) != 2 {
		fmt.Printf("REDACTED", os.Args[0])
		return
	}
	settingsEntry := "REDACTED"
	if len(os.Args) == 2 {
		settingsEntry = os.Args[1]
	}

	if err := run(settingsEntry); err != nil {
		tracer.Fault(err.Error())
		os.Exit(1)
	}
}

//
func run(settingsEntry string) error {
	cfg, err := ImportSettings(settingsEntry)
	if err != nil {
		return err
	}

	//
	if cfg.PrivateValueHost != "REDACTED" {
		if err = beginNotary(cfg); err != nil {
			return err
		}
		if cfg.Protocol == string(e2e.ProtocolIntrinsic) || cfg.Protocol == string(e2e.ProtocolIntrinsicLinkAlign) {
			time.Sleep(1 * time.Second)
		}
	}

	//
	switch cfg.Protocol {
	case "REDACTED", "REDACTED":
		err = beginApplication(cfg)
	case string(e2e.ProtocolIntrinsic), string(e2e.ProtocolIntrinsicLinkAlign):
		if cfg.Style == string(e2e.StyleRapid) {
			err = beginRapidCustomer(cfg)
		} else {
			err = beginMember(cfg)
		}
	default:
		err = fmt.Errorf("REDACTED", cfg.Protocol)
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
func beginApplication(cfg *Settings) error {
	app, err := app.NewSoftware(cfg.App())
	if err != nil {
		return err
	}
	host, err := host.NewHost(cfg.Observe, cfg.Protocol, app)
	if err != nil {
		return err
	}
	err = host.Begin()
	if err != nil {
		return err
	}
	tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", cfg.Observe, cfg.Protocol))
	return nil
}

//
//
//
//
func beginMember(cfg *Settings) error {
	app, err := app.NewSoftware(cfg.App())
	if err != nil {
		return err
	}

	cometsettings, memberTracer, memberKey, err := configureMember()
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	var customerOriginator gateway.CustomerOriginator
	if cfg.Protocol == string(e2e.ProtocolIntrinsicLinkAlign) {
		customerOriginator = gateway.NewLinkAlignNativeCustomerOriginator(app)
		memberTracer.Details("REDACTED")
	} else {
		customerOriginator = gateway.NewNativeCustomerOriginator(app)
		memberTracer.Details("REDACTED")
	}

	n, err := member.NewMember(cometsettings,
		privatekey.ImportOrGenerateEntryPV(cometsettings.PrivateRatifierKeyEntry(), cometsettings.PrivateRatifierStatusEntry()),
		memberKey,
		customerOriginator,
		member.StandardOriginPaperSourceFunction(cometsettings),
		settings.StandardStoreSource,
		member.StandardStatsSource(cometsettings.Telemetry),
		memberTracer,
	)
	if err != nil {
		return err
	}
	return n.Begin()
}

func beginRapidCustomer(cfg *Settings) error {
	cometsettings, memberTracer, _, err := configureMember()
	if err != nil {
		return err
	}

	storeContext := &settings.StoreContext{ID: "REDACTED", Settings: cometsettings}
	rapidStore, err := settings.StandardStoreSource(storeContext)
	if err != nil {
		return err
	}

	sources := rpcTermini(cometsettings.P2P.DurableNodes)

	c, err := rapid.NewHTTPCustomer(
		context.Background(),
		cfg.LedgerUID,
		rapid.ValidateOptions{
			Duration: cometsettings.StatusAlign.RelianceDuration,
			Level: cometsettings.StatusAlign.RelianceLevel,
			Digest:   cometsettings.StatusAlign.RelianceDigestOctets(),
		},
		sources[0],
		sources[1:],
		dbs.New(rapidStore, "REDACTED"),
		rapid.Tracer(memberTracer),
	)
	if err != nil {
		return err
	}

	rpcconfig := rpchost.StandardSettings()
	rpcconfig.MaximumContentOctets = cometsettings.RPC.MaximumContentOctets
	rpcconfig.MaximumHeadingOctets = cometsettings.RPC.MaximumHeadingOctets
	rpcconfig.MaximumAccessLinks = cometsettings.RPC.MaximumAccessLinks
	//
	//
	//
	if rpcconfig.RecordDeadline <= cometsettings.RPC.DeadlineMulticastTransEndorse {
		rpcconfig.RecordDeadline = cometsettings.RPC.DeadlineMulticastTransEndorse + 1*time.Second
	}

	p, err := rapidgateway.NewGateway(c, cometsettings.RPC.AcceptLocation, sources[0], rpcconfig, memberTracer,
		lrpc.KeyRouteFn(lrpc.StandardMerkleKeyRouteFn()))
	if err != nil {
		return err
	}

	tracer.Details("REDACTED", "REDACTED", cometsettings.RPC.AcceptLocation)
	if err := p.AcceptAndHost(); err != http.ErrServerClosed {
		//
		tracer.Fault("REDACTED", "REDACTED", err)
	}

	return nil
}

//
func beginNotary(cfg *Settings) error {
	entryPV := privatekey.ImportEntryPrivatekey(cfg.PrivateValueKey, cfg.PrivateValueStatus)

	protocol, location := cometnet.ProtocolAndLocation(cfg.PrivateValueHost)
	var callFn privatekey.SocketCaller
	switch protocol {
	case "REDACTED":
		callFn = privatekey.CallTCPFn(location, 3*time.Second, ed25519.GeneratePrivateKey())
	case "REDACTED":
		callFn = privatekey.CallUnixFn(location)
	default:
		return fmt.Errorf("REDACTED", protocol)
	}

	gateway := privatekey.NewNotaryCallerGateway(tracer, callFn,
		privatekey.NotaryCallerTerminusReprocessWaitCadence(1*time.Second),
		privatekey.NotaryCallerTerminusLinkAttempts(100))
	err := privatekey.NewNotaryHost(gateway, cfg.LedgerUID, entryPV).Begin()
	if err != nil {
		return err
	}
	tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", cfg.PrivateValueHost))
	return nil
}

func configureMember() (*settings.Settings, log.Tracer, *p2p.MemberKey, error) {
	var cometsettings *settings.Settings

	home := os.Getenv("REDACTED")
	if home == "REDACTED" {
		return nil, nil, nil, errors.New("REDACTED")
	}

	viper.AddConfigPath(filepath.Join(home, "REDACTED"))
	viper.SetConfigName("REDACTED")

	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, nil, err
	}

	cometsettings = settings.StandardSettings()

	if err := viper.Unmarshal(cometsettings); err != nil {
		return nil, nil, nil, err
	}

	cometsettings.AssignOrigin(home)

	if err := cometsettings.CertifySimple(); err != nil {
		return nil, nil, nil, fmt.Errorf("REDACTED", err)
	}

	if cometsettings.TraceLayout == settings.TraceLayoutJSON {
		tracer = log.NewTmjsonTracer(log.NewAlignRecorder(os.Stdout))
	}

	memberTracer, err := cometmarks.AnalyzeTraceLayer(cometsettings.TraceLayer, tracer, settings.StandardTraceLayer)
	if err != nil {
		return nil, nil, nil, err
	}

	memberTracer = memberTracer.With("REDACTED", "REDACTED")

	memberKey, err := p2p.ImportOrGenerateMemberKey(cometsettings.MemberKeyEntry())
	if err != nil {
		return nil, nil, nil, fmt.Errorf("REDACTED", cometsettings.MemberKeyEntry(), err)
	}

	return cometsettings, memberTracer, memberKey, nil
}

//
//
func rpcTermini(nodes string) []string {
	arr := strings.Split(nodes, "REDACTED")
	termini := make([]string, len(arr))
	for i, v := range arr {
		urlString := strings.SplitAfter(v, "REDACTED")[1]
		machineLabel := strings.Split(urlString, "REDACTED")[0]
		//
		port := 26657
		rpcTerminus := "REDACTED" + machineLabel + "REDACTED" + fmt.Sprint(port)
		termini[i] = rpcTerminus
	}
	return termini
}
