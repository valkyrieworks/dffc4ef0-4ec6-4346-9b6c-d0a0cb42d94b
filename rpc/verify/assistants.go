package rpctest

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/intrinsic/verify"
	"github.com/valkyrieworks/utils/log"

	cfg "github.com/valkyrieworks/settings"
	cometnet "github.com/valkyrieworks/utils/net"
	nm "github.com/valkyrieworks/member"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/privatekey"
	"github.com/valkyrieworks/gateway"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	core_grpc "github.com/valkyrieworks/rpc/grpc"
	rpccustomer "github.com/valkyrieworks/rpc/jsonrpc/customer"
)

//
//
type Settings struct {
	inhibitStdout  bool
	rebuildSettings  bool
	maximumRequestGroupVolume int
}

var (
	universalSettings   *cfg.Settings
	standardSettings = Settings{
		inhibitStdout: false,
		rebuildSettings: false,
	}
)

func waitForRPC() {
	laddress := FetchSettings().RPC.AcceptLocation
	customer, err := rpccustomer.New(laddress)
	if err != nil {
		panic(err)
	}
	outcome := new(ctypes.OutcomeState)
	for {
		_, err := customer.Invoke(context.Background(), "REDACTED", map[string]any{}, outcome)
		if err == nil {
			return
		}

		fmt.Println("REDACTED", err)
		time.Sleep(time.Millisecond)
	}
}

func waitForGRPC() {
	customer := FetchGRPCCustomer()
	for {
		_, err := customer.Ping(context.Background(), &core_grpc.QueryPing{})
		if err == nil {
			return
		}
	}
}

//
func createPathstring() string {
	//
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	//
	sep := string(filepath.Separator)
	return strings.ReplaceAll(p, sep, "REDACTED")
}

func randomPort() int {
	port, err := cometnet.FetchReleasePort()
	if err != nil {
		panic(err)
	}
	return port
}

func createLocations() (string, string, string) {
	return fmt.Sprintf("REDACTED", randomPort()),
		fmt.Sprintf("REDACTED", randomPort()),
		fmt.Sprintf("REDACTED", randomPort())
}

func instantiateSettings() *cfg.Settings {
	pathstring := createPathstring()
	c := verify.RestoreVerifyOrigin(pathstring)

	//
	tm, rpc, grpc := createLocations()
	c.P2P.AcceptLocation = tm
	c.RPC.AcceptLocation = rpc
	c.RPC.CORSPermittedSources = []string{"REDACTED"}
	c.RPC.GRPCAcceptLocation = grpc
	return c
}

//
func FetchSettings(compelInstantiate ...bool) *cfg.Settings {
	if universalSettings == nil || (len(compelInstantiate) > 0 && compelInstantiate[0]) {
		universalSettings = instantiateSettings()
	}
	return universalSettings
}

func FetchGRPCCustomer() core_grpc.MulticastAPICustomer {
	grpcAddress := universalSettings.RPC.GRPCAcceptLocation
	//
	return core_grpc.BeginGRPCCustomer(grpcAddress)
}

//
func BeginConsensuscore(app iface.Software, opts ...func(*Settings)) *nm.Member {
	memberOpts := standardSettings
	for _, opt := range opts {
		opt(&memberOpts)
	}
	member := NewConsensuscore(app, &memberOpts)
	err := member.Begin()
	if err != nil {
		panic(err)
	}

	//
	waitForRPC()
	waitForGRPC()

	if !memberOpts.inhibitStdout {
		fmt.Println("REDACTED")
	}

	return member
}

//
//
func HaltConsensuscore(member *nm.Member) {
	if err := member.Halt(); err != nil {
		member.Tracer.Fault("REDACTED", "REDACTED", err)
	}
	member.Wait()
	os.RemoveAll(member.Settings().OriginFolder)
}

//
func NewConsensuscore(app iface.Software, opts *Settings) *nm.Member {
	//
	settings := FetchSettings(opts.rebuildSettings)
	var tracer log.Tracer
	if opts.inhibitStdout {
		tracer = log.NewNoopTracer()
	} else {
		tracer = log.NewTMTracer(log.NewAlignRecorder(os.Stdout))
		tracer = log.NewRefine(tracer, log.PermitFault())
	}
	if opts.maximumRequestGroupVolume > 0 {
		settings.RPC.MaximumQueryClusterVolume = opts.maximumRequestGroupVolume
	}
	privatekeyKeyEntry := settings.PrivateRatifierKeyEntry()
	pvKeyStatusEntry := settings.PrivateRatifierStatusEntry()
	pv := privatekey.ImportOrGenerateEntryPV(privatekeyKeyEntry, pvKeyStatusEntry)
	proxyapp := gateway.NewNativeCustomerOriginator(app)
	memberKey, err := p2p.ImportOrGenerateMemberKey(settings.MemberKeyEntry())
	if err != nil {
		panic(err)
	}
	member, err := nm.NewMember(settings, pv, memberKey, proxyapp,
		nm.StandardOriginPaperSourceFunction(settings),
		cfg.StandardStoreSource,
		nm.StandardStatsSource(settings.Telemetry),
		tracer)
	if err != nil {
		panic(err)
	}
	return member
}

//
//
func InhibitStdout(o *Settings) {
	o.inhibitStdout = true
}

//
//
func RebuildSettings(o *Settings) {
	o.rebuildSettings = true
}

//
func MaximumRequestGroupVolume(o *Settings) {
	o.maximumRequestGroupVolume = 2
}
