package rpc

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/cors"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/node"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
)

//
type Daemon struct {
	Location    string //
	Processor http.Handler
	Tracer  log.Tracer
	Settings  *settings.RemoteSettings
}

//
func Paths(cfg settings.RemoteSettings, s status.Depot, bs status.LedgerDepot, transoffset transferordinal.TransferOrdinalizer, ldgoffset ordinalizer.LedgerOrdinalizer, tracer log.Tracer) base.PathsIndex { //
	env := &base.Context{
		Settings:           cfg,
		LedgerOrdinalizer:     ldgoffset,
		TransferOrdinalizer:        transoffset,
		StatusDepot:       s,
		LedgerDepot:       bs,
		AgreementHandler: pauseChronizeValidatorImplementation{},
		Tracer:           tracer,
	}
	return base.PathsIndex{
		"REDACTED":       node.FreshRemoteMethod(env.LedgerchainDetails, "REDACTED"),
		"REDACTED": node.FreshRemoteMethod(env.AgreementSettings, "REDACTED"),
		"REDACTED":            node.FreshRemoteMethod(env.Ledger, "REDACTED"),
		"REDACTED":    node.FreshRemoteMethod(env.LedgerViaDigest, "REDACTED"),
		"REDACTED":    node.FreshRemoteMethod(env.LedgerOutcomes, "REDACTED"),
		"REDACTED":           node.FreshRemoteMethod(env.Endorse, "REDACTED"),
		"REDACTED":           node.FreshRemoteMethod(env.Heading, "REDACTED"),
		"REDACTED":   node.FreshRemoteMethod(env.HeadingViaDigest, "REDACTED"),
		"REDACTED":       node.FreshRemoteMethod(env.Assessors, "REDACTED"),
		"REDACTED":               node.FreshRemoteMethod(env.Tx, "REDACTED"),
		"REDACTED":        node.FreshRemoteMethod(env.TransferLookup, "REDACTED"),
		"REDACTED":     node.FreshRemoteMethod(env.LedgerLookup, "REDACTED"),
	}
}

//
//
//
func Processor(remoteSettings *settings.RemoteSettings, paths base.PathsIndex, tracer log.Tracer) http.Handler {
	mux := http.NewServeMux()
	watermarkTracer := tracer.Using("REDACTED", "REDACTED")
	wm := node.FreshWebterminalAdministrator(paths,
		node.RetrieveThreshold(remoteSettings.MaximumContentOctets))
	wm.AssignTracer(watermarkTracer)
	mux.HandleFunc("REDACTED", wm.WebterminalProcessor)

	node.EnrollRemoteRoutines(mux, paths, tracer)
	var originProcessor http.Handler = mux
	if remoteSettings.EqualsCrossoriginActivated() {
		originProcessor = appendCrossoriginProcessor(remoteSettings, mux)
	}
	return originProcessor
}

func appendCrossoriginProcessor(remoteSettings *settings.RemoteSettings, h http.Handler) http.Handler {
	crossoriginIntermediary := cors.New(cors.Options{
		AllowedOrigins: remoteSettings.CrossoriginPermittedSources,
		AllowedMethods: remoteSettings.CrossoriginPermittedApproaches,
		AllowedHeaders: remoteSettings.CrossoriginPermittedHeadings,
	})
	h = crossoriginIntermediary.Handler(h)
	return h
}

type pauseChronizeValidatorImplementation struct{}

func (pauseChronizeValidatorImplementation) AwaitChronize() bool {
	return false
}

//
//
func (srv *Daemon) OverhearAlsoAttend(ctx context.Context) error {
	observer, err := node.Overhear(srv.Location, srv.Settings.MaximumInitiateLinks)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		observer.Close()
	}()
	return node.Attend(observer, srv.Processor, srv.Tracer, daemonRemoteSettings(srv.Settings))
}

//
//
func (srv *Daemon) OverhearAlsoAttendTransportsec(ctx context.Context, licenseRecord, tokenRecord string) error {
	observer, err := node.Overhear(srv.Location, srv.Settings.MaximumInitiateLinks)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		observer.Close()
	}()
	return node.AttendTransportsec(observer, srv.Processor, licenseRecord, tokenRecord, srv.Tracer, daemonRemoteSettings(srv.Settings))
}

func daemonRemoteSettings(r *settings.RemoteSettings) *node.Settings {
	cfg := node.FallbackSettings()
	cfg.MaximumContentOctets = r.MaximumContentOctets
	cfg.MaximumHeadingOctets = r.MaximumHeadingOctets
	//
	//
	//
	if cfg.PersistDeadline <= r.DeadlineMulticastTransferEndorse {
		cfg.PersistDeadline = r.DeadlineMulticastTransferEndorse + 1*time.Second
	}
	return cfg
}
