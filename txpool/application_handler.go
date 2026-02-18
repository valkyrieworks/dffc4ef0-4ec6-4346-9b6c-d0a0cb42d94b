package txpool

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/p2p"
	protomemory "github.com/valkyrieworks/schema/consensuscore/txpool"
	"github.com/valkyrieworks/kinds"
	"github.com/pkg/errors"
)

//
type ApplicationHandler struct {
	p2p.RootHandler
	settings  *settings.TxpoolSettings
	txpool *ApplicationTxpool

	ctx       context.Context
	revokeCtx context.CancelFunc

	divertedOn           atomic.Bool
	waitForDivertingOnChan chan struct{}
}

func NewApplicationHandler(
	settings *settings.TxpoolSettings,
	txpool *ApplicationTxpool,
	waitAlign bool,
) *ApplicationHandler {
	ctx, revokeCtx := context.WithCancel(context.Background())

	r := &ApplicationHandler{
		settings:               settings,
		txpool:              txpool,
		ctx:                  ctx,
		revokeCtx:            revokeCtx,
		divertedOn:           atomic.Bool{},
		waitForDivertingOnChan: nil,
	}

	r.RootHandler = *p2p.NewRootHandler("REDACTED", r)

	if waitAlign {
		r.divertedOn.Store(false)
		r.waitForDivertingOnChan = make(chan struct{})
	} else {
		r.divertedOn.Store(true)
	}

	return r
}

//
func (r *ApplicationHandler) OnBegin() error {
	if !r.divertedOn.Load() {
		r.Tracer.Details("REDACTED")
	}

	if !r.settings.Multicast {
		r.Tracer.Details("REDACTED")
		return nil
	}

	go func() {
		defer func() {
			if p := recover(); p != nil {
				r.Tracer.Fault("REDACTED", "REDACTED", p)
			}
		}()

		//
		//
		maximumGroupVolumeOctets := r.settings.MaximumTransferOctets
		if r.settings.MaximumClusterOctets > 0 {
			maximumGroupVolumeOctets = r.settings.MaximumClusterOctets
		}

		r.multicastTransfersGroup(r.ctx, maximumGroupVolumeOctets)

		r.Tracer.Details("REDACTED")
	}()

	return nil
}

func (r *ApplicationHandler) OnHalt() {
	if !r.activated() {
		return
	}

	//
	r.revokeCtx()
}

//
func (r *ApplicationHandler) FetchStreams() []*p2p.StreamDefinition {
	greatestTransfer := make([]byte, r.settings.MaximumTransferOctets)
	groupMessage := protomemory.Signal{
		Sum: &protomemory.Signal_Trans{
			Txs: &protomemory.Txs{Txs: [][]byte{greatestTransfer}},
		},
	}

	return []*p2p.StreamDefinition{
		{
			ID:                  TxpoolConduit,
			Urgency:            5,
			AcceptSignalVolume: groupMessage.Volume(),
			SignalKind:         &protomemory.Signal{},
		},
	}
}

//
func (r *ApplicationHandler) WaitAlign() bool {
	return !r.activated()
}

//
func (r *ApplicationHandler) ActivateInOutTrans() {
	if !r.divertedOn.CompareAndSwap(false, true) {
		//
		return
	}

	r.Tracer.Details("REDACTED")
	close(r.waitForDivertingOnChan)
}

func (r *ApplicationHandler) Accept(e p2p.Packet) {
	if !r.activated() {
		r.Tracer.Diagnose("REDACTED")
		return
	}

	nodeUID := e.Src.ID()

	txs, err := transFromPacket(e)
	if err != nil {
		r.Tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", nodeUID)
		//
		return
	}

	r.txpool.stats.GroupVolume.With("REDACTED", "REDACTED").Observe(float64(len(txs)))

	for _, tx := range txs {
		r.embedTransfer(nodeUID, tx)
	}
}

func (r *ApplicationHandler) embedTransfer(nodeUID p2p.ID, tx kinds.Tx) {
	err := r.txpool.EmbedTransfer(tx)
	if err == nil {
		//
		return
	}

	transferDigest := transferDigest(tx)
	switch {
	case errors.Is(err, ErrViewedTransfer):
		r.Tracer.Diagnose("REDACTED", "REDACTED", transferDigest, "REDACTED", nodeUID)
	case errors.As(err, &ErrTransferTooBulky{}):
		r.Tracer.Diagnose("REDACTED", "REDACTED", err, "REDACTED", transferDigest, "REDACTED", nodeUID)
	default:
		r.Tracer.Details("REDACTED", "REDACTED", err, "REDACTED", transferDigest, "REDACTED", nodeUID)
	}
}

//
//
//
//
func (r *ApplicationHandler) multicastTransfersGroup(ctx context.Context, maximumGroupVolumeOctets int) {
	//
	//
	influx := r.txpool.TransferFlow(ctx)

	for txs := range influx {
		groupings := segmentTrans(txs, maximumGroupVolumeOctets)
		for _, txs := range groupings {
			r.multicast(txs)
		}
	}
}

func (r *ApplicationHandler) multicast(txs kinds.Txs) {
	r.txpool.stats.GroupVolume.With("REDACTED", "REDACTED").Observe(float64(len(txs)))

	r.Router.MulticastAsync(p2p.Packet{
		Signal:   &protomemory.Txs{Txs: txs.ToSegmentOfOctets()},
		StreamUID: TxpoolConduit,
	})
}

func (r *ApplicationHandler) activated() bool {
	return r.divertedOn.Load()
}

func transFromPacket(e p2p.Packet) ([]kinds.Tx, error) {
	msg, ok := e.Signal.(*protomemory.Txs)
	if !ok {
		return nil, fmt.Errorf("REDACTED", e.Signal)
	}

	transCrude := msg.FetchTrans()
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
func segmentTrans(txs kinds.Txs, maximumGroupVolumeOctets int) []kinds.Txs {
	//
	if len(txs) == 0 {
		return nil
	}

	segments := []kinds.Txs{}

	finalSegmentVolumeOctets := 0
	finalSegment := kinds.Txs{}

	for _, tx := range txs {
		transferVolumeOctets := len(tx)

		//
		if (finalSegmentVolumeOctets + transferVolumeOctets) > maximumGroupVolumeOctets {
			//
			//
			if len(finalSegment) > 0 {
				segments = append(segments, finalSegment)
			}

			//
			finalSegment = kinds.Txs{}
			finalSegmentVolumeOctets = 0
		}

		finalSegment = append(finalSegment, tx)
		finalSegmentVolumeOctets += transferVolumeOctets
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
