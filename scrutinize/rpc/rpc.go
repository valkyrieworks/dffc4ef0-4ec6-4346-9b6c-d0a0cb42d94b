package rpc

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/cors"

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/rpc/core"
	"github.com/valkyrieworks/rpc/jsonrpc/host"
	"github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/status/transordinal"
)

//
type Host struct {
	Address    string //
	Manager http.Handler
	Tracer  log.Tracer
	Settings  *settings.RPCSettings
}

//
func Paths(cfg settings.RPCSettings, s status.Depot, bs status.LedgerDepot, transferidx transordinal.TransOrdinaler, ledgeridx ordinaler.LedgerOrdinaler, tracer log.Tracer) core.PathsIndex { //
	env := &core.Context{
		Settings:           cfg,
		LedgerOrdinaler:     ledgeridx,
		TransOrdinaler:        transferidx,
		StatusDepot:       s,
		LedgerDepot:       bs,
		AgreementHandler: waitAlignValidatorImpl{},
		Tracer:           tracer,
	}
	return core.PathsIndex{
		"REDACTED":       host.NewRPCFunction(env.LedgerchainDetails, "REDACTED"),
		"REDACTED": host.NewRPCFunction(env.AgreementOptions, "REDACTED"),
		"REDACTED":            host.NewRPCFunction(env.Ledger, "REDACTED"),
		"REDACTED":    host.NewRPCFunction(env.LedgerByDigest, "REDACTED"),
		"REDACTED":    host.NewRPCFunction(env.LedgerOutcomes, "REDACTED"),
		"REDACTED":           host.NewRPCFunction(env.Endorse, "REDACTED"),
		"REDACTED":           host.NewRPCFunction(env.Heading, "REDACTED"),
		"REDACTED":   host.NewRPCFunction(env.HeadingByDigest, "REDACTED"),
		"REDACTED":       host.NewRPCFunction(env.Ratifiers, "REDACTED"),
		"REDACTED":               host.NewRPCFunction(env.Tx, "REDACTED"),
		"REDACTED":        host.NewRPCFunction(env.TransferScan, "REDACTED"),
		"REDACTED":     host.NewRPCFunction(env.LedgerScan, "REDACTED"),
	}
}

//
//
//
func Manager(rpcSettings *settings.RPCSettings, paths core.PathsIndex, tracer log.Tracer) http.Handler {
	mux := http.NewServeMux()
	wmTracer := tracer.With("REDACTED", "REDACTED")
	wm := host.NewWebchannelOverseer(paths,
		host.ScanCeiling(rpcSettings.MaximumContentOctets))
	wm.AssignTracer(wmTracer)
	mux.HandleFunc("REDACTED", wm.WebchannelManager)

	host.EnrollRPCRoutines(mux, paths, tracer)
	var originManager http.Handler = mux
	if rpcSettings.IsCorsActivated() {
		originManager = appendCORSManager(rpcSettings, mux)
	}
	return originManager
}

func appendCORSManager(rpcSettings *settings.RPCSettings, h http.Handler) http.Handler {
	corsInterceptor := cors.New(cors.Options{
		AllowedOrigins: rpcSettings.CORSPermittedSources,
		AllowedMethods: rpcSettings.CORSPermittedTechniques,
		AllowedHeaders: rpcSettings.CORSPermittedHeadings,
	})
	h = corsInterceptor.Handler(h)
	return h
}

type waitAlignValidatorImpl struct{}

func (waitAlignValidatorImpl) WaitAlign() bool {
	return false
}

//
//
func (srv *Host) AcceptAndHost(ctx context.Context) error {
	observer, err := host.Observe(srv.Address, srv.Settings.MaximumAccessLinks)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		observer.Close()
	}()
	return host.Attend(observer, srv.Manager, srv.Tracer, hostRPCSettings(srv.Settings))
}

//
//
func (srv *Host) ObserveAndAttendTLS(ctx context.Context, tokenEntry, keyEntry string) error {
	observer, err := host.Observe(srv.Address, srv.Settings.MaximumAccessLinks)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		observer.Close()
	}()
	return host.AttendTLS(observer, srv.Manager, tokenEntry, keyEntry, srv.Tracer, hostRPCSettings(srv.Settings))
}

func hostRPCSettings(r *settings.RPCSettings) *host.Settings {
	cfg := host.StandardSettings()
	cfg.MaximumContentOctets = r.MaximumContentOctets
	cfg.MaximumHeadingOctets = r.MaximumHeadingOctets
	//
	//
	//
	if cfg.RecordDeadline <= r.DeadlineMulticastTransEndorse {
		cfg.RecordDeadline = r.DeadlineMulticastTransEndorse + 1*time.Second
	}
	return cfg
}
