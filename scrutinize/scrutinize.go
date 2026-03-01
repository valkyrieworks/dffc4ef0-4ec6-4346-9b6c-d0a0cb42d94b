package scrutinize

import (
	"context"
	"errors"
	"net"
	"os"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/scrutinize/rpc"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	endorsementtexts "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/texts"
	remotecore "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/ledger"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"

	"golang.org/x/sync/errgroup"
)

var tracer = log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))

//
//
//
//
//
//
type Auditor struct {
	paths remotecore.PathsIndex

	settings *settings.RemoteSettings

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
	cfg *settings.RemoteSettings,
	bs status.LedgerDepot,
	ss status.Depot,
	transoffset transferordinal.TransferOrdinalizer,
	ldgoffset ordinalizer.LedgerOrdinalizer,
) *Auditor {
	paths := rpc.Paths(*cfg, ss, bs, transoffset, ldgoffset, tracer)
	eb := kinds.FreshIncidentPipeline()
	eb.AssignTracer(tracer.Using("REDACTED", "REDACTED"))
	return &Auditor{
		paths: paths,
		settings: cfg,
		tracer: tracer,
		ss:     ss,
		bs:     bs,
	}
}

//
func FreshOriginatingSettings(cfg *settings.Settings) (*Auditor, error) {
	bytesDatastore, err := settings.FallbackDatastoreSupplier(&settings.DatastoreScope{ID: "REDACTED", Settings: cfg})
	if err != nil {
		return nil, err
	}
	bs := depot.FreshLedgerDepot(bytesDatastore)
	sDB, err := settings.FallbackDatastoreSupplier(&settings.DatastoreScope{ID: "REDACTED", Settings: cfg})
	if err != nil {
		return nil, err
	}
	producePaper, err := kinds.InaugurationPaperOriginatingRecord(cfg.InaugurationRecord())
	if err != nil {
		return nil, err
	}
	transoffset, ldgoffset, err := ledger.OrdinalizerOriginatingSettings(cfg, settings.FallbackDatastoreSupplier, producePaper.SuccessionUUID)
	if err != nil {
		return nil, err
	}
	ss := status.FreshDepot(sDB, status.DepotChoices{})
	return New(cfg.RPC, bs, ss, transoffset, ldgoffset), nil
}

//
//
func (ins *Auditor) Run(ctx context.Context) error {
	defer ins.bs.Shutdown()
	defer ins.ss.Shutdown()

	return initiateRemoteNodes(ctx, ins.settings, ins.tracer, ins.paths)
}

func initiateRemoteNodes(ctx context.Context, cfg *settings.RemoteSettings, tracer log.Tracer, paths remotecore.PathsIndex) error {
	g, tempctx := errgroup.WithContext(ctx)
	overhearLocations := endorsementtexts.PartitionAlsoShortenBlank(cfg.OverhearLocation, "REDACTED", "REDACTED")
	rh := rpc.Processor(cfg, paths, tracer)
	for _, observerLocation := range overhearLocations {
		node := rpc.Daemon{
			Tracer:  tracer,
			Settings:  cfg,
			Processor: rh,
			Location:    observerLocation,
		}
		if cfg.EqualsTransportsecActivated() {
			tokenRecord := cfg.TokenRecord()
			licenseRecord := cfg.LicenseRecord()
			observerLocation := observerLocation
			g.Go(func() error {
				tracer.Details("REDACTED", "REDACTED", observerLocation,
					"REDACTED", licenseRecord, "REDACTED", tokenRecord)
				err := node.OverhearAlsoAttendTransportsec(tempctx, licenseRecord, tokenRecord)
				if !errors.Is(err, net.ErrClosed) {
					return err
				}
				tracer.Details("REDACTED", "REDACTED", observerLocation)
				return nil
			})
		} else {
			observerLocation := observerLocation
			g.Go(func() error {
				tracer.Details("REDACTED", "REDACTED", observerLocation)
				err := node.OverhearAlsoAttend(tempctx)
				if !errors.Is(err, net.ErrClosed) {
					return err
				}
				tracer.Details("REDACTED", "REDACTED", observerLocation)
				return nil
			})
		}
	}
	return g.Wait()
}
