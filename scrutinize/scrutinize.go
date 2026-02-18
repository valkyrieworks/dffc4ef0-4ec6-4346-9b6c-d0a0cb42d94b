package scrutinize

import (
	"context"
	"errors"
	"net"
	"os"

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/scrutinize/rpc"
	"github.com/valkyrieworks/utils/log"
	endorsementsstrings "github.com/valkyrieworks/utils/strings"
	rpcbase "github.com/valkyrieworks/rpc/core"
	"github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/status/ordinaler/ledger"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"

	"golang.org/x/sync/errgroup"
)

var tracer = log.NewTMTracer(log.NewAlignRecorder(os.Stdout))

//
//
//
//
//
//
type Auditor struct {
	paths rpcbase.PathsIndex

	settings *settings.RPCSettings

	tracer log.Tracer

	//
	//
	ss status.Depot
	bs status.LedgerDepot
}

//
//
//
//
//
//
func New(
	cfg *settings.RPCSettings,
	bs status.LedgerDepot,
	ss status.Depot,
	transferidx transordinal.TransOrdinaler,
	ledgeridx ordinaler.LedgerOrdinaler,
) *Auditor {
	paths := rpc.Paths(*cfg, ss, bs, transferidx, ledgeridx, tracer)
	eb := kinds.NewEventBus()
	eb.AssignTracer(tracer.With("REDACTED", "REDACTED"))
	return &Auditor{
		paths: paths,
		settings: cfg,
		tracer: tracer,
		ss:     ss,
		bs:     bs,
	}
}

//
func NewFromSettings(cfg *settings.Settings) (*Auditor, error) {
	szStore, err := settings.StandardStoreSource(&settings.StoreContext{ID: "REDACTED", Settings: cfg})
	if err != nil {
		return nil, err
	}
	bs := depot.NewLedgerDepot(szStore)
	sDB, err := settings.StandardStoreSource(&settings.StoreContext{ID: "REDACTED", Settings: cfg})
	if err != nil {
		return nil, err
	}
	generatePaper, err := kinds.OriginPaperFromEntry(cfg.OriginEntry())
	if err != nil {
		return nil, err
	}
	transferidx, ledgeridx, err := ledger.OrdinalerFromSettings(cfg, settings.StandardStoreSource, generatePaper.LedgerUID)
	if err != nil {
		return nil, err
	}
	ss := status.NewDepot(sDB, status.DepotSettings{})
	return New(cfg.RPC, bs, ss, transferidx, ledgeridx), nil
}

//
//
func (ins *Auditor) Run(ctx context.Context) error {
	defer ins.bs.End()
	defer ins.ss.End()

	return beginRPCHosts(ctx, ins.settings, ins.tracer, ins.paths)
}

func beginRPCHosts(ctx context.Context, cfg *settings.RPCSettings, tracer log.Tracer, paths rpcbase.PathsIndex) error {
	g, tctx := errgroup.WithContext(ctx)
	observeLocations := endorsementsstrings.DivideAndClipEmpty(cfg.AcceptLocation, "REDACTED", "REDACTED")
	rh := rpc.Manager(cfg, paths, tracer)
	for _, observerAddress := range observeLocations {
		host := rpc.Host{
			Tracer:  tracer,
			Settings:  cfg,
			Manager: rh,
			Address:    observerAddress,
		}
		if cfg.IsTLSActivated() {
			keyEntry := cfg.KeyEntry()
			tokenEntry := cfg.TokenEntry()
			observerAddress := observerAddress
			g.Go(func() error {
				tracer.Details("REDACTED", "REDACTED", observerAddress,
					"REDACTED", tokenEntry, "REDACTED", keyEntry)
				err := host.ObserveAndAttendTLS(tctx, tokenEntry, keyEntry)
				if !errors.Is(err, net.ErrClosed) {
					return err
				}
				tracer.Details("REDACTED", "REDACTED", observerAddress)
				return nil
			})
		} else {
			observerAddress := observerAddress
			g.Go(func() error {
				tracer.Details("REDACTED", "REDACTED", observerAddress)
				err := host.AcceptAndHost(tctx)
				if !errors.Is(err, net.ErrClosed) {
					return err
				}
				tracer.Details("REDACTED", "REDACTED", observerAddress)
				return nil
			})
		}
	}
	return g.Wait()
}
