package rpcoverify

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	nm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/peer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	base_grps "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/grps"
	customeriface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/customer"
)

//
//
type Choices struct {
	quashStandardemission  bool
	rebuildSettings  bool
	maximumRequestClusterExtent int
}

var (
	universalSettings   *cfg.Settings
	fallbackChoices = Choices{
		quashStandardemission: false,
		rebuildSettings: false,
	}
)

func pauseForeachRemote() {
	localaddr := FetchSettings().RPC.OverhearLocation
	customer, err := customeriface.New(localaddr)
	if err != nil {
		panic(err)
	}
	outcome := new(ktypes.OutcomeCondition)
	for {
		_, err := customer.Invocation(context.Background(), "REDACTED", map[string]any{}, outcome)
		if err == nil {
			return
		}

		fmt.Println("REDACTED", err)
		time.Sleep(time.Millisecond)
	}
}

func pauseForeachGRPS() {
	customer := FetchGRPSCustomer()
	for {
		_, err := customer.Ping(context.Background(), &base_grps.SolicitPing{})
		if err == nil {
			return
		}
	}
}

//
func createFilepath() string {
	//
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	//
	sep := string(filepath.Separator)
	return strings.ReplaceAll(p, sep, "REDACTED")
}

func arbitraryChannel() int {
	channel, err := strongmindnet.ObtainLiberateChannel()
	if err != nil {
		panic(err)
	}
	return channel
}

func createLocations() (string, string, string) {
	return fmt.Sprintf("REDACTED", arbitraryChannel()),
		fmt.Sprintf("REDACTED", arbitraryChannel()),
		fmt.Sprintf("REDACTED", arbitraryChannel())
}

func generateSettings() *cfg.Settings {
	filepath := createFilepath()
	c := verify.RestoreVerifyOrigin(filepath)

	//
	tm, rpc, grps := createLocations()
	c.P2P.OverhearLocation = tm
	c.RPC.OverhearLocation = rpc
	c.RPC.CrossoriginPermittedSources = []string{"REDACTED"}
	c.RPC.GRPSOverhearLocation = grps
	return c
}

//
func FetchSettings(compelGenerate ...bool) *cfg.Settings {
	if universalSettings == nil || (len(compelGenerate) > 0 && compelGenerate[0]) {
		universalSettings = generateSettings()
	}
	return universalSettings
}

func FetchGRPSCustomer() base_grps.MulticastAPICustomer {
	grpsLocation := universalSettings.RPC.GRPSOverhearLocation
	//
	return base_grps.InitiateGRPSCustomer(grpsLocation)
}

//
func InitiateStrongmind(app iface.Platform, choices ...func(*Choices)) *nm.Peer {
	peerOptions := fallbackChoices
	for _, opt := range choices {
		opt(&peerOptions)
	}
	peer := FreshStrongmind(app, &peerOptions)
	err := peer.Initiate()
	if err != nil {
		panic(err)
	}

	//
	pauseForeachRemote()
	pauseForeachGRPS()

	if !peerOptions.quashStandardemission {
		fmt.Println("REDACTED")
	}

	return peer
}

//
//
func HaltStrongmind(peer *nm.Peer) {
	if err := peer.Halt(); err != nil {
		peer.Tracer.Failure("REDACTED", "REDACTED", err)
	}
	peer.Pause()
	os.RemoveAll(peer.Settings().OriginPath)
}

//
func FreshStrongmind(app iface.Platform, choices *Choices) *nm.Peer {
	//
	settings := FetchSettings(choices.rebuildSettings)
	var tracer log.Tracer
	if choices.quashStandardemission {
		tracer = log.FreshNooperationTracer()
	} else {
		tracer = log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))
		tracer = log.FreshRefine(tracer, log.PermitFailure())
	}
	if choices.maximumRequestClusterExtent > 0 {
		settings.RPC.MaximumSolicitClusterExtent = choices.maximumRequestClusterExtent
	}
	prvTokenRecord := settings.PrivateAssessorTokenRecord()
	prvTokenStatusRecord := settings.PrivateAssessorStatusRecord()
	pv := privatevalue.FetchEitherProduceRecordPRV(prvTokenRecord, prvTokenStatusRecord)
	proxyapp := delegate.FreshRegionalCustomerOriginator(app)
	peerToken, err := p2p.FetchEitherProducePeerToken(settings.PeerTokenRecord())
	if err != nil {
		panic(err)
	}
	peer, err := nm.FreshPeer(settings, pv, peerToken, proxyapp,
		nm.FallbackInaugurationPaperSupplierMethod(settings),
		cfg.FallbackDatastoreSupplier,
		nm.FallbackTelemetrySupplier(settings.Telemetry),
		tracer)
	if err != nil {
		panic(err)
	}
	return peer
}

//
//
func QuashStandardemission(o *Choices) {
	o.quashStandardemission = true
}

//
//
func RebuildSettings(o *Choices) {
	o.rebuildSettings = true
}

//
func MaximumRequestClusterExtent(o *Choices) {
	o.maximumRequestClusterExtent = 2
}
