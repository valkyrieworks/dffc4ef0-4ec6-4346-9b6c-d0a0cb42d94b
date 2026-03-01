package txpool

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	schemaspace "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/txpool"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/pkg/errors"
)

//
type ApplicationHandler struct {
	p2p.FoundationHandler
	settings  *settings.TxpoolSettings
	txpool *ApplicationTxpool

	ctx       context.Context
	abortContext context.CancelFunc

	routedUpon           atomic.Bool
	pauseForeachRoutingUponChnl chan struct{}
}

func FreshApplicationHandler(
	settings *settings.TxpoolSettings,
	txpool *ApplicationTxpool,
	pauseForeachChronize bool,
) *ApplicationHandler {
	ctx, abortContext := context.WithCancel(context.Background())

	r := &ApplicationHandler{
		settings:               settings,
		txpool:              txpool,
		ctx:                  ctx,
		abortContext:            abortContext,
		routedUpon:           atomic.Bool{},
		pauseForeachRoutingUponChnl: nil,
	}

	r.FoundationHandler = *p2p.FreshFoundationHandler("REDACTED", r)

	if pauseForeachChronize {
		r.routedUpon.Store(false)
		r.pauseForeachRoutingUponChnl = make(chan struct{})
	} else {
		r.routedUpon.Store(true)
	}

	return r
}

//
func (r *ApplicationHandler) UponInitiate() error {
	if !r.routedUpon.Load() {
		r.Tracer.Details("REDACTED")
	}

	if !r.settings.Multicast {
		r.Tracer.Details("REDACTED")
		return nil
	}

	go func() {
		defer func() {
			if p := recover(); p != nil {
				r.Tracer.Failure("REDACTED", "REDACTED", p)
			}
		}()

		//
		//
		maximumClusterExtentOctets := r.settings.MaximumTransferOctets
		if r.settings.MaximumClusterOctets > 0 {
			maximumClusterExtentOctets = r.settings.MaximumClusterOctets
		}

		r.multicastTransfersCluster(r.ctx, maximumClusterExtentOctets)

		r.Tracer.Details("REDACTED")
	}()

	return nil
}

func (r *ApplicationHandler) UponHalt() {
	if !r.activated() {
		return
	}

	//
	r.abortContext()
}

//
func (r *ApplicationHandler) ObtainConduits() []*p2p.ConduitDefinition {
	greatestTransfer := make([]byte, r.settings.MaximumTransferOctets)
	clusterSignal := schemaspace.Signal{
		Sum: &schemaspace.Artifact_Trans{
			Txs: &schemaspace.Txs{Txs: [][]byte{greatestTransfer}},
		},
	}

	return []*p2p.ConduitDefinition{
		{
			ID:                  TxpoolConduit,
			Urgency:            5,
			ObtainSignalVolume: clusterSignal.Extent(),
			SignalKind:         &schemaspace.Signal{},
		},
	}
}

//
func (r *ApplicationHandler) AwaitChronize() bool {
	return !r.activated()
}

//
func (r *ApplicationHandler) ActivateInsideOutputTrans() {
	if !r.routedUpon.CompareAndSwap(false, true) {
		//
		return
	}

	r.Tracer.Details("REDACTED")
	close(r.pauseForeachRoutingUponChnl)
}

func (r *ApplicationHandler) Accept(e p2p.Wrapper) {
	if !r.activated() {
		r.Tracer.Diagnose("REDACTED")
		return
	}

	nodeUUID := e.Src.ID()

	txs, err := transOriginatingWrapper(e)
	if err != nil {
		r.Tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", nodeUUID)
		//
		return
	}

	r.txpool.telemetry.ClusterExtent.With("REDACTED", "REDACTED").Observe(float64(len(txs)))

	for _, tx := range txs {
		r.appendTransfer(nodeUUID, tx)
	}
}

func (r *ApplicationHandler) appendTransfer(nodeUUID p2p.ID, tx kinds.Tx) {
	err := r.txpool.AppendTransfer(tx)
	if err == nil {
		//
		return
	}

	transferDigest := transferDigest(tx)
	switch {
	case errors.Is(err, FaultObservedTransfer):
		r.Tracer.Diagnose("REDACTED", "REDACTED", transferDigest, "REDACTED", nodeUUID)
	case errors.As(err, &FaultTransferExcessivelyAmple{}):
		r.Tracer.Diagnose("REDACTED", "REDACTED", err, "REDACTED", transferDigest, "REDACTED", nodeUUID)
	default:
		r.Tracer.Details("REDACTED", "REDACTED", err, "REDACTED", transferDigest, "REDACTED", nodeUUID)
	}
}

//
//
//
//
func (r *ApplicationHandler) multicastTransfersCluster(ctx context.Context, maximumClusterExtentOctets int) {
	//
	//
	influx := r.txpool.TransferInflux(ctx)

	for txs := range influx {
		clusters := segmentTrans(txs, maximumClusterExtentOctets)
		for _, txs := range clusters {
			r.multicast(txs)
		}
	}
}

func (r *ApplicationHandler) multicast(txs kinds.Txs) {
	r.txpool.telemetry.ClusterExtent.With("REDACTED", "REDACTED").Observe(float64(len(txs)))

	r.Router.MulticastAsyncronous(p2p.Wrapper{
		Signal:   &schemaspace.Txs{Txs: txs.TowardSegmentBelongingOctets()},
		ConduitUUID: TxpoolConduit,
	})
}

func (r *ApplicationHandler) activated() bool {
	return r.routedUpon.Load()
}

func transOriginatingWrapper(e p2p.Wrapper) ([]kinds.Tx, error) {
	msg, ok := e.Signal.(*schemaspace.Txs)
	if !ok {
		return nil, fmt.Errorf("REDACTED", e.Signal)
	}

	transCrude := msg.ObtainTrans()
	switch len(transCrude) {
	case 0:
		return nil, fmt.Errorf("REDACTED")
	case 1:
		//
		return []kinds.Tx{kinds.Tx(transCrude[0])}, nil
	default:
		txs := make([]kinds.Tx, len(transCrude))
		for i, tx := range transCrude {
			txs[i] = kinds.Tx(tx)
		}
		return txs, nil
	}
}

//
//
//
//
//
func segmentTrans(txs kinds.Txs, maximumClusterExtentOctets int) []kinds.Txs {
	//
	if len(txs) == 0 {
		return nil
	}

	segments := []kinds.Txs{}

	finalSegmentExtentOctets := 0
	finalSegment := kinds.Txs{}

	for _, tx := range txs {
		transferExtentOctets := len(tx)

		//
		if (finalSegmentExtentOctets + transferExtentOctets) > maximumClusterExtentOctets {
			//
			//
			if len(finalSegment) > 0 {
				segments = append(segments, finalSegment)
			}

			//
			finalSegment = kinds.Txs{}
			finalSegmentExtentOctets = 0
		}

		finalSegment = append(finalSegment, tx)
		finalSegmentExtentOctets += transferExtentOctets
	}

	//
	if len(finalSegment) > 0 {
		segments = append(segments, finalSegment)
	}

	return segments
}

func transferDigest(tx kinds.Tx) string {
	return fmt.Sprintf("REDACTED", tx.Digest())
}
